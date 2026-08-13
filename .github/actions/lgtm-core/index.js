// @ts-nocheck
/**
 * lgtm-core
 * =========
 * Prow-style, CODEOWNERS-driven "/lgtm" + "/approve" merge gate for cncf/toc.
 *
 * This module is consumed by .github/workflows/lgtm.yml through
 * actions/github-script (`require('./.github/actions/lgtm-core/index.js')`).
 * It has NO external npm dependencies — it uses only the authenticated
 * Octokit `github` client and `core` that github-script passes in.
 *
 * Behaviour (see .github/workflows/lgtm.yml header for the full spec):
 *   - `/lgtm`    : a code owner (not the PR author) signs off on review.
 *   - `/approve` : a *different* code owner (not the author) authorises merge.
 *   - Merge gate (`lgtm/approved` commit status) turns green only when there is
 *     at least LGTM_MIN valid `/lgtm` AND at least APPROVE_MIN valid `/approve`
 *     from disjoint people, AND the union of every signer's owned paths covers
 *     every changed file in the PR (coverage mode).
 *   - `/approve` does NOT imply `/lgtm`; the two labels are independent.
 *   - A new commit resets both axes.
 *
 * Security: CODEOWNERS is always read from the PR BASE commit (trusted), never
 * from the PR head, so a PR that edits CODEOWNERS cannot self-authorise.
 */

'use strict';

// ---------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------

const STATUS_CONTEXT = 'lgtm/approved';
const LGTM_LABEL = 'lgtm';
const APPROVED_LABEL = 'approved';

// Hidden machine-readable state carried inside the sticky status comment.
const STATE_RE = /<!--\s*lgtm-state:\s*(\{[\s\S]*?\})\s*-->/;

// Only these exact commands are actioned (anchored, nothing else on the line).
const COMMAND_RE = /^\/(lgtm|approve|remove-lgtm|remove-approve)\s*$/i;

// Label colours/descriptions used only if the labels do not yet exist.
// Keep these in sync with the canonical definitions in .github/labels.yaml.
const LABEL_DEFS = {
  [LGTM_LABEL]: { color: '2da44e', description: 'A code owner has LGTM’d this PR (see the LGTM & Approve Gate)' },
  [APPROVED_LABEL]: { color: '0e8a16', description: 'A code owner has approved this PR (see the LGTM & Approve Gate)' },
};

// Candidate locations for the CODEOWNERS file, in GitHub’s resolution order.
const CODEOWNERS_PATHS = ['.github/CODEOWNERS', 'CODEOWNERS', 'docs/CODEOWNERS'];

// ---------------------------------------------------------------------------
// Config (from repo variables, injected as env by the workflow)
// ---------------------------------------------------------------------------

function readConfig() {
  const bool = (v, dflt) => (v === undefined || v === '' ? dflt : String(v).toLowerCase() === 'true');
  const int = (v, dflt) => {
    const n = parseInt(v, 10);
    return Number.isFinite(n) && n > 0 ? n : dflt;
  };
  const cfg = {
    gateEnabled: bool(process.env.LGTM_GATE_ENABLED, true),
    ownershipMode: (process.env.LGTM_OWNERSHIP_MODE || 'coverage').toLowerCase(),
    lgtmMin: int(process.env.LGTM_MIN, 1),
    approveMin: int(process.env.APPROVE_MIN, 1),
  };
  // Only "coverage" is implemented. Fail loudly rather than silently ignoring an
  // unsupported value so the variable never looks supported when it isn't.
  if (cfg.ownershipMode !== 'coverage') {
    throw new Error(`Unsupported LGTM_OWNERSHIP_MODE: ${cfg.ownershipMode} (supported: coverage)`);
  }
  return cfg;
}

// ---------------------------------------------------------------------------
// Small helpers
// ---------------------------------------------------------------------------

/** Case-insensitive GitHub login comparison. */
const eq = (a, b) => String(a || '').toLowerCase() === String(b || '').toLowerCase();

/** Tag a 403 from the org/teams API so the top level can explain the fix. */
function isOrgReadDenied(err) {
  return err && err.status === 403;
}

// ---------------------------------------------------------------------------
// CODEOWNERS parsing & matching
// ---------------------------------------------------------------------------

/**
 * Load and parse CODEOWNERS from the trusted base commit.
 * Returns an ordered array of { pattern, owners: string[] }.
 */
async function loadRules(github, owner, repo, baseRef) {
  let text = null;
  for (const path of CODEOWNERS_PATHS) {
    try {
      const res = await github.rest.repos.getContent({ owner, repo, path, ref: baseRef });
      if (res.data && res.data.content) {
        text = Buffer.from(res.data.content, 'base64').toString('utf8');
        break;
      }
    } catch (e) {
      if (e.status !== 404) throw e;
    }
  }
  if (text === null) return [];
  return parseCodeowners(text);
}

function parseCodeowners(text) {
  const rules = [];
  for (const raw of text.split(/\r?\n/)) {
    const line = raw.replace(/#.*$/, '').trim();
    if (!line) continue;
    const parts = line.split(/\s+/);
    const pattern = parts.shift();
    const owners = parts.filter(Boolean);
    if (!pattern) continue;
    rules.push({ pattern, owners, re: codeownersToRegExp(pattern) });
  }
  return rules;
}

/**
 * Convert a CODEOWNERS glob into a RegExp.
 * Mirrors GitHub’s behaviour closely enough for our patterns:
 *   - leading "/" anchors at the repo root, otherwise matches at any depth;
 *   - "*" matches within a path segment, "**" crosses segments;
 *   - a directory pattern also matches everything beneath it.
 */
function codeownersToRegExp(pattern) {
  let p = pattern.trim();
  const anchored = p.startsWith('/');
  if (anchored) p = p.slice(1);
  if (p.endsWith('/')) p = p.slice(0, -1);

  let escaped = p.replace(/[.+^${}()|[\]\\]/g, '\\$&');
  escaped = escaped.replace(/\*\*/g, '\u0000'); // globstar placeholder
  escaped = escaped.replace(/\*/g, '[^/]*');
  escaped = escaped.replace(/\u0000/g, '.*');

  const head = anchored ? '^' : '(?:^|.*/)';
  // Match the path itself or anything nested under it (directory ownership).
  return new RegExp(head + escaped + '(?:/.*)?$');
}

/** Owners of a single path, using GitHub’s last-match-wins rule. */
function ownersForPath(rules, path) {
  let owners = [];
  for (const rule of rules) {
    if (rule.re.test(path)) owners = rule.owners;
  }
  return owners;
}

// ---------------------------------------------------------------------------
// Ownership / team resolution
// ---------------------------------------------------------------------------

/**
 * Does an owner token (e.g. "@cncf/cncf-toc", "@user", or a bare "user")
 * grant ownership to `user`? Team lookups are cached per (team, user).
 */
async function ownerTokenMatchesUser(github, org, token, user, cache) {
  const team = token.match(/^@([^/]+)\/(.+)$/);
  if (team) {
    const teamOrg = team[1];
    const slug = team[2];
    const key = `${teamOrg}/${slug}#${user.toLowerCase()}`;
    if (cache.has(key)) return cache.get(key);
    let member = false;
    try {
      const res = await github.rest.teams.getMembershipForUserInOrg({
        org: teamOrg,
        team_slug: slug,
        username: user,
      });
      member = !!res.data && res.data.state === 'active';
    } catch (e) {
      if (isOrgReadDenied(e)) throw e; // surfaced loudly at the top level
      member = false; // 404 => not a member
    }
    cache.set(key, member);
    return member;
  }
  // Individual owner, with or without a leading "@".
  const login = token.replace(/^@/, '');
  return eq(login, user);
}

/** Is `user` an owner of at least one of the given files? */
async function userOwnsAnyFile(github, org, user, files, rules, cache) {
  for (const file of files) {
    const owners = ownersForPath(rules, file);
    for (const token of owners) {
      if (await ownerTokenMatchesUser(github, org, token, user, cache)) return true;
    }
  }
  return false;
}

/** Keep only the pooled logins that still own at least one changed file. */
async function filterOwners(github, org, users, files, rules, cache) {
  const kept = [];
  for (const u of users) {
    if (await userOwnsAnyFile(github, org, u, files, rules, cache)) kept.push(u);
  }
  return kept;
}

/**
 * Coverage: every changed file must be owned by at least one pooled signer.
 * Returns { covered, missing: [{ file, owners }] }.
 */
async function computeCoverage(github, org, pool, files, rules, cache) {
  const missing = [];
  for (const file of files) {
    const owners = ownersForPath(rules, file);
    let ok = false;
    for (const u of pool) {
      for (const token of owners) {
        if (await ownerTokenMatchesUser(github, org, token, u, cache)) {
          ok = true;
          break;
        }
      }
      if (ok) break;
    }
    if (!ok) missing.push({ file, owners });
  }
  return { covered: missing.length === 0, missing };
}

// ---------------------------------------------------------------------------
// PR data helpers
// ---------------------------------------------------------------------------

async function listChangedFiles(github, owner, repo, prNumber) {
  const files = await github.paginate(github.rest.pulls.listFiles, {
    owner,
    repo,
    pull_number: prNumber,
    per_page: 100,
  });
  return files.map((f) => f.filename);
}

// ---------------------------------------------------------------------------
// Labels
// ---------------------------------------------------------------------------

async function ensureLabel(github, owner, repo, name) {
  try {
    await github.rest.issues.getLabel({ owner, repo, name });
  } catch (e) {
    if (e.status === 404) {
      const def = LABEL_DEFS[name] || { color: 'ededed', description: '' };
      try {
        await github.rest.issues.createLabel({ owner, repo, name, color: def.color, description: def.description });
      } catch (_) {
        /* raced with another job — ignore */
      }
    }
  }
}

async function setLabel(github, owner, repo, issue_number, name, shouldHave) {
  if (shouldHave) {
    await ensureLabel(github, owner, repo, name);
    try {
      await github.rest.issues.addLabels({ owner, repo, issue_number, labels: [name] });
    } catch (_) {
      /* already present — ignore */
    }
  } else {
    try {
      await github.rest.issues.removeLabel({ owner, repo, issue_number, name });
    } catch (e) {
      if (e.status !== 404) throw e;
    }
  }
}

// ---------------------------------------------------------------------------
// Sticky state comment
// ---------------------------------------------------------------------------

async function findStickyComment(github, owner, repo, issue_number) {
  const comments = await github.paginate(github.rest.issues.listComments, {
    owner,
    repo,
    issue_number,
    per_page: 100,
  });
  // If duplicates ever exist, keep the newest so we don't leave a stale status.
  return comments.filter((c) => STATE_RE.test(c.body || '')).pop() || null;
}

async function readState(github, owner, repo, issue_number) {
  const sticky = await findStickyComment(github, owner, repo, issue_number);
  const base = { commentId: null, sha: null, lgtm: [], approve: [] };
  if (!sticky) return base;
  base.commentId = sticky.id;
  const m = (sticky.body || '').match(STATE_RE);
  if (m) {
    try {
      const parsed = JSON.parse(m[1]);
      base.sha = parsed.sha || null;
      base.lgtm = Array.isArray(parsed.lgtm) ? parsed.lgtm : [];
      base.approve = Array.isArray(parsed.approve) ? parsed.approve : [];
    } catch (_) {
      /* corrupt state — start fresh */
    }
  }
  return base;
}

async function writeState(github, owner, repo, issue_number, state, body) {
  if (state.commentId) {
    await github.rest.issues.updateComment({ owner, repo, comment_id: state.commentId, body });
  } else {
    const res = await github.rest.issues.createComment({ owner, repo, issue_number, body });
    state.commentId = res.data.id;
  }
}

function renderStickyBody(cfg, state, view) {
  const json = JSON.stringify({ sha: state.sha, lgtm: state.lgtm, approve: state.approve });
  const list = (arr) => (arr.length ? arr.map((u) => `@${u}`).join(', ') : '_none_');

  const lines = [];
  lines.push('### 🔖 LGTM / Approve status');
  lines.push('');
  if (!cfg.gateEnabled) {
    lines.push('> ⚠️ The merge gate is currently **disabled** (`LGTM_GATE_ENABLED=false`); this is informational only.');
    lines.push('');
  }
  lines.push(`- **LGTM** (${state.lgtm.length}/${cfg.lgtmMin}): ${list(state.lgtm)} ${view.hasLgtm ? '✅' : '❌'}`);
  lines.push(`- **Approve** (${state.approve.length}/${cfg.approveMin}): ${list(state.approve)} ${view.hasApprove ? '✅' : '❌'}`);
  if (view.cov.covered) {
    lines.push('- **Coverage**: all changed files owned by a signer ✅');
  } else {
    lines.push('- **Coverage**: ❌ still need a code owner for:');
    for (const m of view.cov.missing.slice(0, 15)) {
      const need = m.owners.length ? m.owners.join(', ') : '_no CODEOWNERS entry_';
      lines.push(`  - \`${m.file}\` → ${need}`);
    }
    if (view.cov.missing.length > 15) lines.push(`  - …and ${view.cov.missing.length - 15} more`);
  }
  lines.push('');
  lines.push(`**Merge gate:** ${view.pass ? '🟢 satisfied' : '🔴 not satisfied'}`);
  lines.push('');
  lines.push('<sub>`/lgtm` and `/approve` must come from two different code owners (not the PR author). A new commit resets both. Use `/remove-lgtm` or `/remove-approve` to revoke.</sub>');
  lines.push('');
  lines.push(`<!-- lgtm-state: ${json} -->`);
  return lines.join('\n');
}

// ---------------------------------------------------------------------------
// Feedback to the commenter
// ---------------------------------------------------------------------------

const REJECTIONS = {
  self: 'you cannot `/lgtm` or `/approve` your own pull request.',
  separation: 'the same person cannot both `/lgtm` and `/approve`; those must come from two different code owners.',
  'not-owner': 'you are not a CODEOWNER for any of the files changed in this PR, so your sign-off cannot be counted.',
  'cannot-remove': 'only the person who added a sign-off (or the PR author) can remove it.',
};

async function react(github, owner, repo, commentId, content) {
  try {
    await github.rest.reactions.createForIssueComment({ owner, repo, comment_id: commentId, content });
  } catch (_) {
    /* reactions are best-effort */
  }
}

async function reject(github, owner, repo, prNumber, commentId, actor, reasonKey) {
  await react(github, owner, repo, commentId, '-1');
  const reason = REJECTIONS[reasonKey] || 'this command could not be processed.';
  await github.rest.issues.createComment({
    owner,
    repo,
    issue_number: prNumber,
    body: `@${actor} ${reason}`,
  });
}

// ---------------------------------------------------------------------------
// Core reconcile
// ---------------------------------------------------------------------------

async function recompute({ github, context, cfg, pr, changedFiles, rules, teamCache, state }) {
  const { owner, repo } = context.repo;
  const org = owner;

  // Defensively drop signers who no longer own any changed file, and keep the
  // two axes disjoint.
  state.lgtm = await filterOwners(github, org, state.lgtm, changedFiles, rules, teamCache);
  state.approve = await filterOwners(github, org, state.approve, changedFiles, rules, teamCache);
  state.approve = state.approve.filter((a) => !state.lgtm.some((l) => eq(l, a)));
  state.sha = pr.head.sha;

  const pool = [...state.lgtm, ...state.approve];
  const cov = await computeCoverage(github, org, pool, changedFiles, rules, teamCache);

  const hasLgtm = state.lgtm.length >= cfg.lgtmMin;
  const hasApprove = state.approve.length >= cfg.approveMin;
  const pass = cfg.gateEnabled ? hasLgtm && hasApprove && cov.covered : true;

  await setLabel(github, owner, repo, pr.number, LGTM_LABEL, state.lgtm.length > 0);
  await setLabel(github, owner, repo, pr.number, APPROVED_LABEL, state.approve.length > 0);

  const description = buildStatusDescription(cfg, state, hasLgtm, hasApprove, cov, pass);
  await github.rest.repos.createCommitStatus({
    owner,
    repo,
    sha: pr.head.sha,
    state: pass ? 'success' : 'failure',
    context: STATUS_CONTEXT,
    description: description.slice(0, 140),
  });

  const body = renderStickyBody(cfg, state, { hasLgtm, hasApprove, cov, pass });
  await writeState(github, owner, repo, pr.number, state, body);
}

function buildStatusDescription(cfg, state, hasLgtm, hasApprove, cov, pass) {
  if (!cfg.gateEnabled) return 'Gate disabled (LGTM_GATE_ENABLED=false)';
  if (pass) return `Approved: ${state.lgtm.length} LGTM, ${state.approve.length} approve`;
  const missing = [];
  if (!hasLgtm) missing.push(`${state.lgtm.length}/${cfg.lgtmMin} lgtm`);
  if (!hasApprove) missing.push(`${state.approve.length}/${cfg.approveMin} approve`);
  if (!cov.covered) missing.push(`${cov.missing.length} file(s) uncovered`);
  return `Needs: ${missing.join(', ')}`;
}

// ---------------------------------------------------------------------------
// Entry points (called from the workflow)
// ---------------------------------------------------------------------------

/** issue_comment handler: process /lgtm, /approve and their removals. */
async function handleComment({ github, context, core }) {
  const cfg = readConfig();
  const { owner, repo } = context.repo;
  const issue = context.payload.issue;
  const comment = context.payload.comment;

  if (!issue || !issue.pull_request) return; // issues, not PRs
  if (comment.user && comment.user.type === 'Bot') return; // no bot loops

  const match = (comment.body || '').trim().match(COMMAND_RE);
  if (!match) return;

  const cmd = match[1].toLowerCase();
  const axis = cmd.endsWith('lgtm') ? 'lgtm' : 'approve';
  const isAdd = cmd === 'lgtm' || cmd === 'approve';
  const actor = comment.user.login;
  const prNumber = issue.number;

  try {
    const { data: pr } = await github.rest.pulls.get({ owner, repo, pull_number: prNumber });
    if (pr.state !== 'open') return;

    const changedFiles = await listChangedFiles(github, owner, repo, prNumber);
    const rules = await loadRules(github, owner, repo, pr.base.sha);
    const teamCache = new Map();

    const state = await readState(github, owner, repo, prNumber);
    // A new head commit invalidates all prior sign-offs.
    if (state.sha && state.sha !== pr.head.sha) {
      state.lgtm = [];
      state.approve = [];
    }

    if (isAdd) {
      if (eq(actor, pr.user.login)) return reject(github, owner, repo, prNumber, comment.id, actor, 'self');

      const other = axis === 'lgtm' ? 'approve' : 'lgtm';
      if (state[other].some((u) => eq(u, actor))) {
        return reject(github, owner, repo, prNumber, comment.id, actor, 'separation');
      }

      const owns = await userOwnsAnyFile(github, owner, actor, changedFiles, rules, teamCache);
      if (!owns) return reject(github, owner, repo, prNumber, comment.id, actor, 'not-owner');

      if (!state[axis].some((u) => eq(u, actor))) state[axis].push(actor);
      await react(github, owner, repo, comment.id, '+1');
    } else {
      const canRemove = state[axis].some((u) => eq(u, actor)) || eq(actor, pr.user.login);
      if (!canRemove) return reject(github, owner, repo, prNumber, comment.id, actor, 'cannot-remove');
      state[axis] = state[axis].filter((u) => !eq(u, actor));
      if (eq(actor, pr.user.login)) state[axis] = []; // author resets the whole axis
      await react(github, owner, repo, comment.id, '+1');
    }

    await recompute({ github, context, cfg, pr, changedFiles, rules, teamCache, state });
  } catch (e) {
    handleFatal(core, e);
  }
}

/** pull_request_target handler: init on open/reopen, reset on synchronize. */
async function handleLifecycle({ github, context, core }) {
  const cfg = readConfig();
  const { owner, repo } = context.repo;
  const pr = context.payload.pull_request;
  const action = context.payload.action;

  try {
    const changedFiles = await listChangedFiles(github, owner, repo, pr.number);
    const rules = await loadRules(github, owner, repo, pr.base.sha);
    const teamCache = new Map();
    const state = await readState(github, owner, repo, pr.number);

    if (action === 'synchronize') {
      const changed = state.sha && state.sha !== pr.head.sha;
      state.lgtm = [];
      state.approve = [];
      await recompute({ github, context, cfg, pr, changedFiles, rules, teamCache, state });
      if (changed) {
        await github.rest.issues.createComment({
          owner,
          repo,
          issue_number: pr.number,
          body: 'ℹ️ New commit pushed — previous `/lgtm` and `/approve` sign-offs were cleared. Code owners will need to re-sign.',
        });
      }
      return;
    }

    // opened / reopened: publish the initial (failing) gate + sticky comment.
    await recompute({ github, context, cfg, pr, changedFiles, rules, teamCache, state });
  } catch (e) {
    handleFatal(core, e);
  }
}

function handleFatal(core, err) {
  if (isOrgReadDenied(err)) {
    core.setFailed(
      'Team membership lookup was denied (403). The token in SLASH_COMMANDS_PAT must have ' +
        'organization "Members: Read" permission to resolve CODEOWNERS teams.'
    );
    return;
  }
  core.setFailed(err && err.stack ? err.stack : String(err));
}

module.exports = { handleComment, handleLifecycle };
