# Hive Security Self-Assessment

This document follows the [CNCF TAG-Security self-assessment
template](https://tag-security.cncf.io/community/assessments/guide/self-assessment/)
and is submitted as a supporting artifact of Hive Commons' CNCF Sandbox
application ([cncf/sandbox#516](https://github.com/cncf/sandbox/issues/516)).

> **Scope:** This assessment covers **[hivecommons/hive](https://github.com/hivecommons/hive)** only — the AI agent orchestration system and its hub/spoke deployment. Hive was incubated as a subproject of KubeStellar and was transferred to its own vendor-neutral org, [hivecommons](https://github.com/hivecommons), on 2026-09-03; this assessment does **not** cover [KubeStellar Core](https://github.com/kubestellar/kubestellar) or [KubeStellar Console](https://github.com/kubestellar/console), which has its [own self-assessment](https://github.com/cncf/toc/blob/main/projects/kubestellar/sub-projects/console/security-assessment/self-assessment.md).
>
> **Canonical source:** this document is maintained in the Hive repository at [`src/docs/security-self-assessment.md`](https://github.com/hivecommons/hive/blob/v4/src/docs/security-self-assessment.md). Corrections should be made there and mirrored here.

> **Revision 2 — changes made in response to TAG-Security review
> ([cncf/toc#2286](https://github.com/cncf/toc/pull/2286)).** Seven review
> comments asked, in substance, for the parts the first draft hedged on. Each
> is addressed:
>
> | Review point | Change |
> |---|---|
> | "If I deploy this, what should I worry about?" | New section [If you deploy this, what should you worry about?](#if-you-deploy-this-what-should-you-worry-about) — the attack path stated plainly, a concern/mitigation/non-mitigation table, the three settings that determine exposure, and the worst realistic outcome. |
> | "Have you red teamed `ioscan`? Is this perfect defense or partial mitigation?" | Answered, then **measured**: [ioscan-red-team.md](https://github.com/hivecommons/hive/blob/v4/src/docs/ioscan-red-team.md) runs a 43-payload adversarial corpus and publishes the result — **42% withheld, 58% reached the agent**, with the containment credited to the network deny rules and verified across all four ACMM modes on every test run ([#6685](https://github.com/hivecommons/hive/issues/6685)). |
> | "What is redaction for? What about base64-encoded exfiltration?" | Log scrubbing re-scoped as log hygiene, explicitly **not** an exfiltration control. Running the question against the canary path found a real gap — the egress check was substring-only — filed as [#6686](https://github.com/hivecommons/hive/issues/6686) and **since fixed**: the proxy now decodes and normalizes outbound bodies before matching canaries. |
> | "Get an OpenSSF passing badge." | **Already held, and this document was wrong to imply otherwise.** [Project 14261](https://www.bestpractices.dev/projects/14261) reached **passing** (100%) on 2026-08-27 — four days before the review — and `README.md` was displaying it throughout. The badge entry has since been corrected: 70 URLs in its criteria justifications still pointed at the pre-migration `kubestellar` org. Tracked in [#6684](https://github.com/hivecommons/hive/issues/6684). |
> | "This is a huge risk. Why not mitigate it?" | Half was mitigated: the roster went from **one maintainer to three**, across three affiliations, with a documented security-response process. The unmitigated half — unenforced code ownership — is now stated as the largest remaining process risk, with scoped enforcement tracked in [#6687](https://github.com/hivecommons/hive/issues/6687). |
> | "Looks like AI generated this — you should know the answer." | The "Case studies: not applicable" claim was simply **wrong**: `ADOPTERS.md` lists seven adopters, three in production, one running at ACMM L5/L6 across 43 repositories. Replaced with the real table. |
> | "Go through the open questions and figure them out." | The open-questions section is gone. Every item was run down against the repository and answered in [Questions resolved since first review](#questions-resolved-since-first-review) — including two answers that are "no." |
>
> Net effect: five issues filed ([#6684](https://github.com/hivecommons/hive/issues/6684), [#6685](https://github.com/hivecommons/hive/issues/6685), [#6686](https://github.com/hivecommons/hive/issues/6686), [#6687](https://github.com/hivecommons/hive/issues/6687), [#6688](https://github.com/hivecommons/hive/issues/6688)), one previously-undetected defect found, and several claims corrected against the repository rather than restated.

It complements, and deliberately does not duplicate, three existing security
documents in the Hive repository:

- [security-threat-model.md](https://github.com/hivecommons/hive/blob/v4/src/docs/security-threat-model.md) — attacker-oriented
  view: assets, trust boundaries, threat actors, defense layers, residual
  risks.
- [security-model.md](https://github.com/hivecommons/hive/blob/v4/src/docs/security-model.md) — operator/evaluator-oriented view:
  the seven enforcement layers (dashboard auth, spoke auth, credential
  isolation, sandboxing, GitHub blast-radius controls, hub↔spoke channel,
  hosted-platform isolation) with file/line evidence for each.
- [security.md](https://github.com/hivecommons/hive/blob/v4/src/docs/security.md) — log-scrubbing and secret-redaction
  implementation notes.

Where this document repeats a claim from those pages, it cites the same
evidence rather than re-deriving it. Sections below cite `path:line` against
the Hive repository's `v4` branch (`src/` prefix; commit `327969b0` at time of
writing) wherever a specific mechanism is asserted.

## Metadata

| | |
|---|---|
| Assessment Stage | Complete |
| Software | [hivecommons/hive](https://github.com/hivecommons/hive) |
| Security Provider | No — Hive is not itself a security product. It is an agent-orchestration platform whose core value proposition includes constraining the blast radius of the AI agents it runs; see [Overview](#overview) below. |
| Languages | Go (core: dashboard, hub, proxy, scheduler, agent orchestration — `src/go.mod`); JavaScript (dashboard UI, served inline, no separate SPA build — see `dashboard/`); Shell/Python (deterministic pipeline scripts under `src/bin/`, 45 scripts per [`bin/README.md`](https://github.com/hivecommons/hive/blob/v4/bin/README.md)) |
| SBOM | Tagged releases publish standalone SPDX JSON SBOM release assets for `hive`, `hive-contributor`, and `hive-hub` (`hive-v<version>-sbom.spdx.json`, `hive-contributor-v<version>-sbom.spdx.json`, `hive-hub-v<version>-sbom.spdx.json`) generated by Syft against the already-published GHCR images; see [releases.md](https://github.com/hivecommons/hive/blob/v4/src/docs/releases.md#software-bill-of-materials-sbom). These are deliberately **out-of-band** files, not in-image attestations. `docker.yml` keeps `provenance: false`/`sbom: false` (`docker.yml:295,346,631,818`) because `build-push-action` attestations force an OCI image index and reintroduced the #3760 container-runtime crash loop. The guard in `src/scripts/check-no-image-attestations.sh` asserts the image builds keep plain manifests. |
| Security links | See table below |

| Doc | URL |
|---|---|
| Security policy / vulnerability reporting | [SECURITY.md](https://github.com/hivecommons/hive/blob/v4/SECURITY.md) |
| Threat model | [security-threat-model.md](https://github.com/hivecommons/hive/blob/v4/src/docs/security-threat-model.md) |
| Security model (operator guide) | [security-model.md](https://github.com/hivecommons/hive/blob/v4/src/docs/security-model.md) |
| Log scrubbing / secret redaction | [security.md](https://github.com/hivecommons/hive/blob/v4/src/docs/security.md) |
| ADRs (architecture decision records) | [adr/README.md](https://github.com/hivecommons/hive/blob/v4/src/docs/adr/README.md) |
| OpenSSF Scorecard workflow | [.github/workflows/scorecard.yml](https://github.com/hivecommons/hive/blob/v4/.github/workflows/scorecard.yml) |

## Overview

### Background

Hive orchestrates fleets of AI coding-agent processes (backends: Claude,
GitHub Copilot, Gemini, Goose, Bob, Agy) that autonomously maintain software
projects hosted on GitHub (and GitHub Enterprise/GitLab/Gitea via the
"Forge App" abstraction) — filing issues, opening pull requests, reviewing
code, and, at the highest autonomy setting, merging. Agents run as CLI
subprocesses inside tmux sessions, in a container or Kubernetes pod, and
authenticate to the source-control provider with a GitHub App installation
token or a PAT.

A hub/spoke topology lets one central hub coordinate many independent "spoke"
hive deployments (self-hosted or platform-hosted), each spoke typically
governing one or a handful of repositories.

The defining security property of this project is that the thing being
orchestrated — output from a large language model reading attacker-reachable
GitHub content — is treated as **untrusted input carrying write-capable
credentials to source repositories**. This is not a hypothetical: any public
issue, PR, comment, or label on a governed repository is potential agent
input, and agents that reach the higher autonomy tiers can push code and open
pull requests under real (scoped) GitHub credentials. The project's security
posture is therefore built around the principle stated at the top of
[security-model.md](https://github.com/hivecommons/hive/blob/v4/src/docs/security-model.md): *"if a human would give the same
answer every time, it belongs in infrastructure, not in a prompt."* Enforcement
is intended to sit in deterministic Go/shell code — a network proxy, a token
scope, a file permission — rather than in an instruction the model is asked to
obey.

### Actors

- **Repository maintainer / hive operator** — configures ACMM autonomy level,
  which repositories are governed, which agents run, and credential sources
  (GitHub App vs. PAT, inference backend keys).
- **AI coding agents** — CLI subprocesses that read issues/PRs, propose or
  make changes, and (depending on mode) open PRs, comment, or merge.
- **Public GitHub users** — anyone who can open an issue, comment, or PR on a
  governed repository; from Hive's perspective their text is untrusted input
  that reaches agent prompts.
- **Contributors via ClankeR relay** — external contributors who donate
  compute by running an agent against a hive's queue over a relay protocol
  (see [contributor-relay.md](https://github.com/hivecommons/hive/blob/v4/src/docs/contributor-relay.md)).
- **Hub operators / SaaS platform operators** — run the central hub that
  coordinates registered spokes and, for hosted spokes, provisions
  infrastructure and injects GitHub App credentials.
- **Dashboard/API/terminal users** — humans with read, read-write, merger, or
  owner-level access to a spoke's dashboard, per `security-model.md` Layer 1/2.

### Actions

- Agents read GitHub issues, PRs, comments, labels, and diffs (untrusted text
  boundary).
- Agents are "kicked" (launched) with a constructed prompt built from that
  text plus repository/task context.
- Agents call GitHub through a CLI wrapper whose outbound HTTPS traffic is
  force-redirected through an in-pod MITM proxy that enforces per-mode policy
  (`src/pkg/proxy/rules.go`).
- A deterministic pre-kick pipeline (`run-pipeline.sh`) enumerates and
  classifies work and computes merge eligibility before any agent acts,
  narrowing agent discretion over what is actionable.
- Depending on the operator-selected ACMM level (L1–L6), agents may only
  observe, may file issues, may open hold-gated PRs, or may merge.
- Hive itself (not the agent) performs the highest-privilege GitHub writes —
  PR creation and PR merge — as the App bot, via hard-denied-to-agents relay
  paths (`hive-open-pr`, `hive-merge`; deny rules at
  `src/pkg/proxy/rules.go:215-226`).
- Every hive-mediated creation is attribution-stamped and audit-logged
  (`src/pkg/github/attribution.go`, `src/pkg/dashboard/audit.go`).
- Registered spokes heartbeat operational telemetry (no credentials) to a
  central hub; the hub can push configuration (including GitHub App
  credentials for hub-provisioned hives) back to a spoke.

### Goals

- Constrain what an LLM-driven agent process can do to a real source
  repository, independent of what the agent was prompted or prompt-injected
  to do, via layered deterministic controls (network proxy, token scoping,
  CLI tool-deny, repo allowlist).
- Make every high-impact GitHub write (issue/PR creation, PR merge) happen
  through a Hive-mediated, audited, App-bot-authored path rather than a
  directly agent-held credential.
- Reduce — not eliminate — prompt-injection risk from untrusted GitHub text
  before it reaches an agent prompt (`ioscan`).
- Give operators one legible autonomy dial (ACMM L1–L6) instead of many
  independent toggles, so risk posture is auditable and incrementally
  raisable.
- Provide an accountability trail (attribution trailers + append-only audit
  log) sufficient to answer "which agent/backend/model did this, and when."

### Non-goals

- Hive does not attempt to make LLM output itself trustworthy or to fully
  solve prompt injection. `ioscan` mitigates known injection shapes; it is
  explicitly not a claim of immunity (see [Known weaknesses](#three-most-significant-known-weaknesses)).
- Hive does not sandbox agent code execution at the level of per-run
  containers or microVMs today — agents share one container's process
  namespace, separated by per-UID isolation, not process/kernel isolation.
  This is a documented, open, tracked gap (issue #2804 — see
  `security-threat-model.md` "Residual risks and known gaps").
- Hive does not provide a general-purpose network content firewall. The MITM
  proxy inspects `api.github.com` (and, per `src/pkg/proxy/rules.go:78-80`,
  `api.linear.app`) traffic; other destinations are tunneled without
  inspection (`security-model.md` "Scope and limitations, stated plainly").
- Hive does not manage the security of the underlying cluster, container
  runtime, or the AI provider accounts operators connect (Claude/Copilot/
  Gemini/Goose credentials are used by the CLI as-is; see
  `security-model.md` Layer 3 "Honest scope").

### If you deploy this, what should you worry about?

The rest of this document is organized by control. This section is organized
by *concern*, because a prospective operator's first question is not "what
mechanisms exist" but "what is the realistic path from a stranger typing
something to something bad happening to my code, and what stands in the way?"

**The path, stated plainly.** Anyone with a GitHub account can open an issue
or comment on a governed public repository. That text is enumerated by the
pre-kick pipeline, scanned by `ioscan`, and — if it survives — placed into a
prompt handed to a CLI agent that holds a real, if scoped, GitHub credential.
**There is no step in that chain where a human necessarily reads the
attacker's text first.** That is the core exposure, and no amount of layered
control removes it; the controls bound the consequences.

| Concern | What actually stands in the way | What does *not* |
|---|---|---|
| Attacker text talks an agent into **merging** malicious code | `PUT /pulls/{n}/merge` is **hard-denied for every ACMM mode** at the proxy (`rules.go:215-226`), as is the GraphQL merge mutation. The agent never holds a credential path to merge; `hive-merge` runs as Hive with a SHA-pinned eligibility binding. | `ioscan`. A sufficiently clever prompt defeats a scanner; it does not defeat a rule table the model never sees. |
| Attacker text talks an agent into **opening a PR** with a backdoor | `POST /pulls` is likewise hard-denied for every mode. PR creation goes through `hive-open-pr` as the App bot, attribution-stamped and audit-logged. | Nothing prevents the *content* of a legitimately-created PR from being attacker-influenced. Review it like any other PR. |
| Attacker exfiltrates your **secrets** via the agent | Mode-tiered tokens mean low-tier agents never hold code-write credentials. Optional `ioscan` canaries detect prompt contents echoed into outbound GitHub writes at the proxy, including base64/hex/URL-encoded, reversed and split spellings of the token. | Canaries are **default-off**, and their transform list is finite — an encoding it does not model still passes. `git-receive-pack` bodies are unscannable, and non-GitHub egress is tunneled without inspection. |
| Agent burns your **inference budget** in a loop | Seven-day rolling token budget suppresses kicks on exhaustion. | Nothing caps the *rate* within budget. |
| Compromised agent attacks **other agents** on the same spoke | Per-UID separation stops tmux attach and token-file reads. | Shared container/kernel. See weakness #2. |

**The three settings that actually determine your exposure**, in order:

1. **Your ACMM level.** At L1–L3 an agent can observe and file issues but
   cannot write code; at L6 it can merge. This single dial moves prompt
   templates, credential scope, CLI tool-denies and proxy rules together. It
   is the control to reason about first, and it is the one an operator
   chooses rather than inherits.
2. **Whether the governed repository is public.** A private repository with
   trusted collaborators collapses most of the untrusted-input surface. The
   threat model above is substantially a *public-repository* threat model.
3. **Whether you enabled `ioscan.canaries` and `fail_mode: closed`.** Both
   default to the permissive setting. The defaults favor availability; a
   security-sensitive deployment should change them and accept that a
   scanner outage then becomes a scheduling outage.

**What we would tell a new operator.** Start at L1–L3 on a repository you
would not mind an agent filing a bad issue on. Read the attribution trailers
and the audit log for a week — every hive-mediated write is stamped, so this
is cheap. Raise the level only once the agents' actual behavior on *your*
codebase is boring. Do not start at L6 on a public flagship repository.

**The worst realistic outcome** if every soft control fails and the operator
has set L6: an attacker-influenced pull request is authored and merged by the
App bot, with a full attribution trailer and audit-log entry naming the agent,
backend and model. It is recoverable and traceable, which is the property the
architecture optimizes for — but it is a real bad day, and an operator should
choose an autonomy level with that outcome in mind rather than the happy path.

## Self-assessment use

This self-assessment is created by the Hive maintainers to perform an initial
evaluation of the project's security posture as part of the CNCF Incubation
application process. It is not a substitute for a third-party audit or
penetration test — neither has been performed on this project to date (see
[Security issue resolution](#security-issue-resolution)). TAG-Security and
the CNCF community may use it to assess Hive for Incubation and may request
a joint or independent assessment building on it. This document reflects the
project's understanding of its own architecture at the time of writing and
carries no warranty; it should not be relied upon exclusively to assess a
production deployment's risk without independent review.

## Security functions and features

### Critical security components

These are the components whose failure would most directly let an attacker
turn untrusted GitHub content, or a compromised agent, into unauthorized
repository writes:

1. **MITM egress proxy and ACMM policy rules**
   (`src/pkg/proxy/rules.go`, `src/pkg/proxy/github_proxy.go`). All outbound
   `:443` traffic from agent processes is force-redirected (iptables
   `REDIRECT`, requiring `CAP_NET_ADMIN`) through an in-pod proxy that
   terminates TLS for `api.github.com` (and `api.linear.app`) and evaluates
   every REST method+path and GraphQL mutation against a first-match rule
   table keyed to the agent's ACMM mode (`rules.go:122-178`). Unmatched
   operations deny by default (`AllowedByModeCaps`, `rules.go:249-263`).
   Direct PR creation (`POST /pulls`) and direct PR merge
   (`PUT /pulls/{n}/merge`) are **hard-denied for every mode**
   (`rules.go:215-226`, CWE-863 mitigation) — an agent at any autonomy level,
   however prompted, cannot call these REST routes; it must go through
   `hive-open-pr` / `hive-merge`, which run as Hive itself (App-bot
   credentials, SHA-pinned merge-eligibility binding). GraphQL mutations that
   merge or create/write PRs are classified by regex
   (`graphQLMergeMutationRe`, `graphQLPRWriteMutationRe`, `rules.go:328-338`)
   and required to meet the same mode floor as their REST equivalents, closing
   a GraphQL bypass of the REST rule table.

2. **`ioscan` untrusted-input scanner**
   (`src/pkg/ioscan`, `src/docs/ioscan.md`). Scans GitHub issue/PR/label/
   author/comment text before it enters an agent kick. Deterministic rules
   (Unicode-steganography normalization, base64 decode-and-rescan,
   prompt-injection phrasing, dangerous-directive and secret-shape detection)
   are always the floor when enabled; blocked text is replaced with a visible
   `[ioscan: content withheld — ...]` marker rather than silently dropped or
   passed raw. An optional LLM-judge classifier
   (`ioscan.classifier.enabled`, default `false`) adds semantic
   plain-English-injection detection on top. **Enabled by default**
   (`ioscan.enabled: true` is the default per `ioscan.md:9`).
   Egress canaries (`HIVE-CANARY-*`) are matched both literally and through
   the transport encodings a body can carry: base64 (standard and URL-safe,
   padded or not, line-wrapped, and one level of nesting), hex (either case,
   with per-byte separators or `\x`/`0x` escapes), percent-encoding, reversed
   text, case changes, homoglyph and zero-width substitution, and separator
   splitting — plus any intact 16-hex-character window of a token's random
   half, so a token broken across a body still trips. Decoding recurs a
   bounded number of passes with a byte budget, because the body being
   decoded is attacker-influenced.

   **How well does it work? Measured: a partial mitigation, and a narrow one**
   — see [ioscan-red-team.md](https://github.com/hivecommons/hive/blob/v4/src/docs/ioscan-red-team.md) for the corpus, the
   methodology and the full per-case table. Against 43 adversarial payloads
   across six technique families, the shipped deterministic rules **withheld
   42% from the agent; 58% reached it verbatim** — 40% because no rule fired,
   and a further 19% because a rule fired but the input block policy does not
   block at that severity (`blockedInput` requires Critical, or Injection at
   High+, so every Medium finding is recorded and the raw text is still
   injected). A findings count is therefore **not** a count of attacks
   stopped.
   Detection is strongly uneven by family: Unicode steganography 8/8 blocked —
   the one place the design beats a regex, because hidden characters are
   themselves a High finding — against paraphrased instruction override 2/10
   and split-payload 0/4. The last is structural: each untrusted segment
   (title, body, each comment) is scanned separately, so a phrase split across
   two comments is never a single string for any rule to match. The cost side
   is real too: **2 of 6 benign controls were redacted**, including an issue
   *reporting* a prompt-injection bug.
   A prospective user should assume a determined, encoding-aware attacker
   defeats it — that is now a measurement rather than a caution. What `ioscan`
   reliably does is (a) raise the cost of the *casual* injection attempt, (b)
   normalize away an entire class of invisible-character and homoglyph tricks
   deterministically, and (c) produce an auditable record that something was
   withheld. What contains the *consequence* of a successful injection is not
   `ioscan` at all: it is the hard-denied PR-create/merge relays and the
   mode-tiered token scope, neither of which the model can argue with — and
   that containment is now verified across all four ACMM modes by
   `TestRedTeamContainmentHoldsForEveryMode` on every test run rather than
   asserted here. Operators should size their trust accordingly. The
   evaluation's own limits are stated in it: the optional LLM classifier is
   default-off and its contribution remains unmeasured, and the corpus covers
   the input path only.

   **Exfiltration detection (`ioscan.canaries`, default off).** A per-agent
   `HIVE-CANARY-<48 hex>` token is planted in the agent's prompt, and the
   egress proxy scans outbound GitHub request bodies for it
   (`github_proxy.go:1179-1205`): a hit is positive evidence that prompt
   contents reached a write, and is audit-logged and — with
   `fail_mode: closed` — blocked. The match is no longer a literal substring
   test: `CanaryRegistry.Scan` normalizes the body and decodes what it finds
   before matching, so the encodings listed above are caught
   ([#6686](https://github.com/hivecommons/hive/issues/6686)). Two limits
   remain, and are stated rather than implied: the transform list is
   **finite and enumerated**, so an encoding it does not model — a cipher, a
   custom alphabet, a canary spliced into the middle of an unrelated base64
   run, or a token cut into pieces shorter than a 16-character window and
   spread across separate requests — still passes; and `git-receive-pack`
   bodies are opaque to the scan, so a push is only stopped when fail-closed
   is set (tracked in
   [Known weaknesses](#three-most-significant-known-weaknesses)). A canary
   hit remains positive evidence; a miss has never been proof of absence.

3. **Per-agent scoped GitHub App tokens and mode-tiered credential issuance**
   (per `security-model.md` Layer 5). Advisory-mode agents receive
   read-only, no-issue-write tokens; mid-tier agents get issue-write-only
   tokens with no code access; only trusted tiers get contents/PR write. The
   full shared installation token is never placed in an agent's environment.

4. **Attribution and audit trail**
   (`src/pkg/github/attribution.go`, `src/pkg/dashboard/audit.go`). Every
   hive-mediated GitHub creation (PR, issue, comment, merge, review) is
   stamped with a visible trailer (`— hive: agent=... backend=... model=...`,
   `attribution.go:168-184`) and recorded in an append-only audit log
   (`/data/audit.jsonl`, `dashboard/audit.go:22`) regardless of the trailer
   toggle (`attribution.go:290-295`). This is the accountability substrate an
   operator or incident responder uses to answer "which agent/backend/model
   produced this."

5. **ACMM autonomy-level gating** (`src/pkg/config/packs/level-{1..6}.yaml`,
   ADR-0003). A single human-selected level (L1–L6) maps deterministically to
   per-agent mode, which in turn drives prompt templates, credential scope,
   CLI `--deny-tool`/`--disallowed-tools` flags, and the proxy rule table
   above — the same signal enforced in four independent places rather than
   one toggle an attacker need only defeat once from the *outside*, though
   see the [three most significant weaknesses](#three-most-significant-known-weaknesses)
   for what this design choice trades away.

6. **Forced egress + in-container privilege model**
   (`security-model.md` "Forced proxy egress and CAP_NET_ADMIN", "In-container
   privilege model"). `su-exec` is the sole setuid binary shipped by Hive,
   locked to mode `4750 root:hive-launch` and enforced by a build-time
   contract script (`src/scripts/check-suid-contract.sh`) plus a
   boot-time inventory check (`src/deploy/test_image_suid_inventory.sh`).
   Eleven base-image setuid/setgid binaries inherited from `node:26-slim` were
   found world-executable and are stripped in the same build step — a
   concrete example of the project auditing its own base image rather than
   trusting it by default.

### Security relevant components

- **Trajectory review** (`src/docs/trajectory-review.md`) — a second model
  periodically compares a running agent's stated intent against a bounded
  transcript tail and can pause/alert on divergence. Explicitly semantic
  oversight, not a structural control, and fails open on reviewer outage
  (per `security-threat-model.md` residual risks).
- **Token budget** — a seven-day rolling token budget that suppresses kicks
  on exhaustion, limiting denial-of-wallet from a runaway or compromised
  agent loop.
- **Log scrubbing** (`src/pkg/logscrub`, `security.md`) — redacts eight
  categories from Hive's own structured log output: GitHub token prefixes
  (`ghs_`, `ghp_`, `gho_`, `github_pat_`), JWTs, AWS access-key IDs,
  `Bearer` headers, PEM private keys (plain, encrypted and PGP), and Hive
  canary tokens (`handler.go:29-36`).

  **What this is for, precisely:** it is *operational hygiene for logs Hive
  itself writes* — so that a token appearing in an error string does not end
  up in a pasted log snippet, a support bundle, or a CI artifact. **It is not
  an exfiltration control and must not be read as one.** It sits on Hive's
  logging path, not on the agent's egress path, and it is pattern-based: an
  attacker who prompt-injects an agent into emitting a credential in base64,
  hex, reversed, or split across tokens defeats it completely, and the
  project states that plainly rather than letting the word "redaction" imply
  containment. The controls that do bear on that attack are elsewhere and are
  also imperfect: mode-tiered token scoping means a low-tier agent has no
  code-write credential to leak in the first place; `ioscan`'s canary check
  at the proxy catches prompt contents reaching an outbound GitHub write,
  and — unlike log scrubbing — it does decode base64, hex and percent-encoded
  bodies and normalize away splitting before matching
  ([#6686](https://github.com/hivecommons/hive/issues/6686)), though only for
  the transforms it enumerates and only for the canary token rather than for
  credentials at large; and non-GitHub destinations are tunneled without
  inspection at all. An operator who needs
  a hard guarantee against credential egress should not rely on any of these
  and should scope the credentials themselves.
- **Ed25519-only session/SSO verification** (`security-model.md` "Sessions
  and SSO are Ed25519-only") — the legacy HMAC session-cookie lane and
  fleet-wide shared heartbeat bearer were both removed in v4; a spoke
  lacking a verification key fails closed (HTTP 503) rather than falling
  back to a weaker check.
- **Master-key generational rotation** with a bounded dual-generation
  acceptance window (default 7 days) for sessions/heartbeats/SSO, and no
  dual lane at all for terminal/invite keys (`security-model.md` "Master key
  rotation").
- **Supply-chain pinning** — base images and CI actions pinned by digest/SHA
  (`src/Dockerfile:8,50,78`; `actions/checkout@3d3c42e...` in
  `.github/workflows/docker.yml:85,280,318`); `npm install --ignore-scripts`
  for all global AI-CLI installs.
- **OpenSSF Scorecard** — runs weekly and on push to `main`/`v4`
  (`.github/workflows/scorecard.yml`), via a shared reusable workflow pinned
  by commit SHA.

## Project compliance

Hive does not currently hold any formal security certification (e.g.
FIPS, Common Criteria, SOC 2) and makes no claim to one. Relevant
project-level compliance signals:

- **OpenSSF Scorecard**: automated, runs weekly and on push to `main`/`v4`
  (`scorecard.yml`, via a SHA-pinned reusable workflow in
  `hivecommons/infra`); results publish to the repository's code-scanning
  alerts. No specific score floor is gated in CI at present.
- **OpenSSF Best Practices Badge**: **held at the `passing` level** — project
  [14261](https://www.bestpractices.dev/projects/14261), achieved 2026-08-27,
  100% of the passing criteria (all 67 Met or Not-Applicable). This is a
  maintainer **self**-certification against published criteria, publicly
  auditable per-criterion — not a third-party audit, and so not a
  counter-example to the "no formal security certification" sentence above.
  The badge is displayed in `README.md` and its state is live, so the image
  tracks the project's current level rather than the level at the time of this
  writing; no level beyond `passing` is claimed here, for the same reason the
  Scorecard number above is not frozen into this document. TAG-Security's
  feedback was that a passing badge is achievable low-hanging fruit whose
  criteria are things the project should be doing anyway; the maintainers
  agreed, and [#6684](https://github.com/hivecommons/hive/issues/6684) tracked
  it to completion.
- **DCO (Developer Certificate of Origin)**: enforced for human contributors
  via `copilot-dco.yml` and is a stated policy requirement for agent-authored
  commits (`git commit -s`) per `security-model.md` Layer 5 — described there
  as *"policy-driven inside hive; enforce it repo-side (DCO check) for a hard
  guarantee,"* i.e. Hive asks agents to sign off but the hard gate is a
  repo-side branch-protection setting the operator must also enable.
- **License**: Apache License 2.0 (`LICENSE`), OSI-approved, CNCF-preferred.
- **SBOMs**: tagged releases attach standalone SPDX JSON SBOM files for each
  published image (`hive`, `hive-contributor`, `hive-hub`) as GitHub Release
  assets, generated by Syft against the already-published GHCR images. They
  are intentionally not attached to image manifests; the container-runtime
  reason is explained under [Metadata](#metadata) and in
  [releases.md](https://github.com/hivecommons/hive/blob/v4/src/docs/releases.md#software-bill-of-materials-sbom).

## Secure development practices

### Development pipeline

- **Language and structure**: Go (`src/`, `go 1.25.6` per `src/go.mod:3`) for
  the core dashboard/hub/proxy/scheduler/agent-orchestration code; a JS
  dashboard UI served inline (no separate SPA build step); Python/Shell for
  the deterministic pre-kick pipeline (45 scripts, indexed in
  `bin/README.md`).
- **Branch model**: `v4` is the sole actively maintained line;
  `v2` was retired in August 2026 (`src/docs/README.md` header, `migration-v2-v4.md`).
- **CI gating**: PR-triggered test execution runs from
  `.github/workflows/v2-tests.yml` (name is a historical artifact of the
  branch-rename; it is the workflow that gates `v4` PRs) as a sharded
  `go test -short -race -count=1` run, with an hourly full-suite monitor in
  `coverage-hourly.yml` that additionally auto-files an issue on a coverage
  regression. Coverage floors are intentionally duplicated between the two
  workflows with an in-repo comment instructing both be kept in sync
  (`coverage-hourly.yml:15-18`) — a manual-sync risk noted here rather than
  hidden.
- **Container build**: multi-stage `src/Dockerfile`, digest-pinned base
  images (`golang:1.27-alpine@sha256:...`, `node:26-slim@sha256:...` ×2), a
  build-time SUID-contract check that fails the build on any unexpected
  setuid/setgid binary, and checksum-verified tool downloads (tmux, ttyd,
  `gh`, goose, `su-exec`).
- **No SBOM/provenance attestation** is attached to built images — explicitly
  disabled (`sbom: false`, citing issue #3760) rather than merely absent by
  omission. Release SBOMs are generated out-of-band as standalone SPDX JSON
  files and attached to the GitHub Release, so the published image manifests
  stay plain.
- **Maintainer roster**: `OWNERS` lists **three** approvers/reviewers —
  Andy Anderson (@clubanderson, IBM), James Reilly (@hanthor, Universal
  Blue) and Doug Baggett (@Danathar, independent). This is a change since
  this document's first revision, when the roster was one person; the
  bus-factor and security-response concerns raised in review were addressed
  by growing the committee across three distinct affiliations rather than by
  restating the risk. `OWNERS` is required to stay in sync with the
  Maintainer Committee table in upstream `GOVERNANCE-HIVE.md`, and the file
  says so explicitly, calling drift "a governance drift bug, not a
  housekeeping detail."
- **CODEOWNERS**: `.github/CODEOWNERS` is scoped to the security-sensitive
  paths whose changes can directly alter security boundaries (proxy policy,
  Dockerfiles, workflow definitions, deploy manifests, SUID contract checks,
  and key/cookie/session handling), and every entry names all three current
  maintainers (`@clubanderson`, `@hanthor`, `@Danathar`). The remaining
  repo-side step is enabling "Require review from Code Owners" in `v4` branch
  protection; until that is enabled, CODEOWNERS is still advisory rather than
  enforced — see [Known weaknesses](#three-most-significant-known-weaknesses).

### Communication channels

- **Internal**: GitHub issues/PRs/discussions on the `hivecommons/hive`
  repository; the project is largely maintained by autonomous agent fleets
  operating under the same ACMM framework the software implements
  (i.e., Hive is substantially self-hosting — its own maintenance PRs are
  frequently agent-authored, subject to the same guardrails described in this
  document).
- **Inbound**: public GitHub issues/PRs are the primary inbound channel;
  private vulnerability reports go through GitHub Security Advisories per
  `SECURITY.md`.
- **Outbound**: release notes via `CHANGELOG.md`. There is deliberately **no
  mailing list, Slack, Matrix or Discord channel** for project decisions:
  `CONTRIBUTING.md` directs contributors to "discuss design and review
  questions in GitHub issues and PRs so decisions remain public and
  searchable." (Hive the *product* integrates with Slack/Discord/ntfy for
  operator alerting — see `notifications.md` — but that is a feature of the
  software, not a project communication channel, and the two should not be
  confused.) The practical security consequence is that there is no private
  side channel where a vulnerability discussion could happen off-record;
  everything is either a public GitHub thread or a private GitHub Security
  Advisory.

### Ecosystem

Hive is an independent project in the [hivecommons](https://github.com/hivecommons)
org (transferred out of the KubeStellar org on 2026-09-03, where it was
incubated as a subproject). It interoperates with, but does not depend on: GitHub/GitHub Enterprise/GitLab/Gitea (via the "Forge App"
abstraction), several AI coding-agent CLI backends (Claude Code, GitHub
Copilot CLI, Gemini, Goose, Bob/bobshell, Agy), and self-hosted inference
gateways (LiteLLM, vLLM, llm-d, watsonx) reached only through an in-pod
credential-translating proxy (`security-model.md` Layer 3). It is
independent of any single cloud provider or AI vendor.

## Security issue resolution

### Responsible disclosures process

Vulnerability reports are handled through **GitHub private vulnerability
reporting** (Security tab → "Report a vulnerability"), not public issues, PRs,
or discussions, per [`SECURITY.md`](https://github.com/hivecommons/hive/blob/v4/SECURITY.md).
Reporters are asked for affected component/branch/commit, description and
impact, reproduction steps, and any supporting logs/PoC/config.

**Vulnerability response process**: the stated target is acknowledgement
within **5 business days** (`SECURITY.md` "What to Expect"), followed by
investigation/confirmation, a coordinated fix and disclosure timeline, and
optional reporter credit. Who performs that work is now documented:
[`security-response.md`](https://github.com/hivecommons/hive/blob/v4/src/docs/security-response.md)
establishes that security response is a duty of the **Maintainer Committee**
— there is no separate security team, the committee *is* the response team —
and that any of the three maintainers can receive and triage a report, with
whoever picks up the GitHub Security Advisory notification driving it and
looping in the others for confirmation, severity and disclosure timing. The
page deliberately does not restate the roster, pointing at `OWNERS` instead
so a second list cannot drift.

Remaining gaps, stated rather than implied: there is **no published
CVSS-scoring policy** and **no maximum time-to-fix SLA** beyond asking
reporters for "a reasonable opportunity to remediate." The acknowledgement
target is a commitment; time-to-fix is not.

### Incident response

A security-specific response process is documented in
[`security-response.md`](https://github.com/hivecommons/hive/blob/v4/src/docs/security-response.md)
— who responds, how a report is triaged, and how disclosure is coordinated —
distinct from
[`docs/HUB_DISASTER_RECOVERY.md`](https://github.com/hivecommons/hive/blob/v4/docs/HUB_DISASTER_RECOVERY.md),
which covers hub-level disaster recovery (backup/restore, spoke fleet
recovery, operator communication) and is an operational rather than security
runbook. The master-key rotation flow (`security-model.md` "Master key
rotation") provides the mechanism to revoke and replace compromised key
material fleet-wide without a full re-provision, which is the concrete
containment action most security incidents in this architecture would
require.

What is documented is roles and handling. What is **not** yet documented is a
post-incident review practice (blameless postmortem, publication commitment)
and user-notification SLAs for a confirmed compromise. Those are real gaps
and are named here rather than papered over.

## Three most significant known weaknesses

Stated here explicitly, and cross-referenced to the existing threat model,
because a self-assessment that only lists strengths is not credible:

1. **Prompt injection is a fundamental, unsolved risk class for this
   architecture — mitigated, not eliminated.** Any public GitHub issue, PR,
   comment, label, or author field is a potential injection vector into an
   agent's prompt (`security-threat-model.md` "Trust boundaries": *"Public
   issue authors can place arbitrary text in titles, labels, bodies, and
   comments that Hive may include in a kick"*). `ioscan`'s deterministic
   rules and optional semantic classifier reduce this materially, but the
   **default `ioscan.fail_mode` is `open`** (`ioscan.md:11` — *"open (default)
   redacts"*), meaning a scanner outage, timeout, or classifier failure fails
   toward continuing to process input rather than halting it. The semantic
   (LLM-judge) classifier layer is explicitly **fail-open on errors/timeouts
   by design** (`ioscan.md:26`, ADR-0008: *"Classifier failures and budget
   exhaustion fail open"*) so a reviewer outage cannot itself become a
   scheduler outage — a deliberate availability/security tradeoff the project
   documents rather than hides. The deterministic-rules floor and the
   hard-deny network relays for PR create/merge are the parts of the defense
   that do not depend on any model call succeeding; the semantic layer is
   explicitly the part that can silently degrade. `security-threat-model.md`
   states this directly: *"ioscan semantic classification is optional and
   fail-open... it intentionally fails open on model errors/timeouts to avoid
   turning reviewer outages into scheduler outages."*

2. **Shared-container execution is a materially weaker isolation boundary
   than per-run sandboxing, and this is an open, tracked gap.** Agents run as
   separate Unix UIDs inside one Hive container/pod, not in separate
   containers or microVMs. Per-UID separation (`security-model.md` Layer 4)
   stops one agent from attaching to another's tmux session or reading its
   GitHub token file, but it does not provide the isolation guarantees a
   compromised or maliciously-instructed agent process would need to be
   contained against, e.g., host/kernel-level attack surface shared with
   every other agent in the pod. `security-threat-model.md` names this
   directly under "Residual risks and known gaps": *"Shared-container
   execution remains a material risk... it is not equivalent to per-run
   containers or microVMs,"* tracked in open issue
   [#2804](https://github.com/hivecommons/hive/issues/2804), which — per the
   same doc — also proposes moving live GitHub write credentials out of the
   agent sandbox entirely (*"current agents can still need live credentials
   to push/open PRs"*). Until #2804 lands, a compromised agent process that
   evades the network/token controls still executes inside the same
   container as every other agent on that spoke.

3. **Unenforced code ownership on security-sensitive paths — the
   single-maintainer half of this risk has been mitigated; the
   branch-protection half has not.** When this assessment was first
   submitted, `OWNERS` listed exactly one approver, and TAG-Security review
   correctly identified that as a serious risk rather than a footnote. It
   has been addressed: the Maintainer Committee now has **three members
   across three distinct affiliations** — @clubanderson (IBM), @hanthor
   (Universal Blue), @Danathar (independent) — and
   `security-response.md` now names the committee as the security response
   team with an explicit "whoever picks up the advisory drives it" triage
   rule. Bus factor and vendor neutrality both improved as a result, and the
   5-business-day acknowledgement commitment no longer rests on one person's
   availability.

   **What remains unmitigated is the enforcement gap, and it is real.**
   `.github/CODEOWNERS` now covers only the security-sensitive paths —
   Dockerfiles, CI workflows, deploy manifests, SUID contract checks, proxy
   policy, and key/cookie/session handling — and lists all three current
   maintainers on every entry. But "Require review from Code Owners" still
   has to be enabled in `v4` branch protection after the scoped file lands.
   Until then, a change to the MITM proxy's deny-rule table — the control
   this document credits as the thing an attacker cannot argue with — can
   merge on green CI with no human security reviewer.

   **Why it has not simply been turned on.** This is the honest trade at the
   center of the project: Hive is substantially self-hosting, and its own
   agent fleet merges green PRs. Enabling required code-owner review with
   broad coverage would halt that automation repository-wide. The defensible
   resolution is not "all or nothing" but the scoped ownership file above,
   leaving ordinary code paths on the automated merge path while making the
   human-review boundary enforceable for security-critical changes.

   Until that ships, an evaluator should treat "the deny-rule table is
   enforced in code" as true and "the deny-rule table is protected from
   unreviewed change" as **not** true, and weight this accordingly. It is
   the single largest process (as opposed to technical) risk this assessment
   identifies.

A secondary, narrower point worth surfacing alongside the three above: the
**audit log's retention is size-triggered, not time-triggered** — rotation
occurs at 5 MB with 3 backups retained
(`src/pkg/dashboard/audit.go:23-24`, `auditMaxSizeMB = 5`,
`auditMaxBackups = 3`; a 90-day `auditMaxAgeDays` cap also applies but only
prunes files already past the size-based rotation). On a busy hive with many
audited actions per day, the effective lookback window this provides an
incident responder can be materially shorter than 90 days; on a quiet hive it
can be much longer. This is a real, if secondary, forensic-capability limit
on the accountability mechanism this document otherwise credits as a
strength.

One further narrow gap, surfaced because it was found while responding to
TAG-Security review and it would be poor practice to let review-driven
discoveries go unrecorded. (A second one recorded here in revision 2 — that
the canary egress check matched literal substrings only, so a base64-encoded
canary passed unflagged — has since been closed: `CanaryRegistry.Scan`
decodes and normalizes before matching, and the residual limit is now the
finite transform list described under
[Critical security components](#critical-security-components), not the
absence of decoding.)

- **`git-receive-pack` bodies are opaque to canary scanning.** A `git push`
  is pack-encoded, so the proxy cannot inspect it for leaked markers; with
  `fail_mode: closed` the push is refused outright, and otherwise it
  proceeds unscanned (`github_proxy.go:1182`). An operator relying on
  canaries should understand that they cover the GitHub *API* egress path,
  not the git transport.

## Appendix

### Known issues over time

No CVEs have been filed against this project to date. No third-party
security audit or penetration test has been performed. This is stated
plainly rather than implied by omission: any future reader should treat the
absence of reported CVEs as reflecting a young project without external
audit coverage, not as an implicit clean bill of health under adversarial
review. The security-relevant hardening changes referenced by number
throughout `security-threat-model.md` (e.g. #2760, #2747, #2670, #2680,
#2666, #2805, #2754, #2756, #2762, #2771, #2758, #2755, #2662) represent
maintainer-identified and maintainer-fixed hardening work during v4
development, not externally reported vulnerabilities.

### Open SSF best practices

An [OpenSSF Scorecard](https://github.com/hivecommons/hive/blob/v4/.github/workflows/scorecard.yml)
workflow runs weekly and on push to `main`/`v4`, publishing to the
repository's code-scanning alerts. This assessment deliberately does not
freeze a numeric score into the text — it moves independently of this
document — and a reviewer should pull the current result.

**OpenSSF Best Practices Badge: held at the passing level.** The project is
registered as [project 14261](https://www.bestpractices.dev/projects/14261)
and reached **passing** (100% of the passing criteria) on **2026-08-27**; it
has not lapsed since. The badge is displayed in `README.md` and its live
status is available from `https://www.bestpractices.dev/projects/14261/badge`.
Silver and gold are not claimed. The badge is self-certified, as the
programme intends, so its value to a reviewer is not the graphic but the
per-criterion written justifications behind it, each linking to the specific
file, workflow or policy that satisfies it — those can be checked rather than
taken on trust.

Two earlier revisions of this section were wrong about this, in opposite
directions, and both errors are worth recording. The first said the badge was
"not yet assessed whether this is planned," which was an evasion rather than
an answer; TAG-Security review called it out. The correction then overshot by
describing the badge as not yet held and merely "being pursued" — but it had
in fact been awarded on 2026-08-27, four days before that review. In both
cases this document was written from itself rather than re-derived from
observable state, which is the same defect that left the maintainer count
stale. The badge status above is taken from the programme's public API.

### Case studies

Hive maintains an [`ADOPTERS.md`](https://github.com/hivecommons/hive/blob/v4/ADOPTERS.md)
file with seven organizations, three of them at **Production** maturity.
The security-relevant point is not the count but that these are deployments
where the threat model in this document is live — real repositories, real
credentials, real untrusted public input:

| Adopter | Use | Maturity | Evidence |
|---|---|---|---|
| [KubeStellar Console](https://github.com/kubestellar/console) | Autonomous maintenance of the Console codebase — triage, fixes, review, merge | Production | Hive's origin adopter |
| [tunaos.org](https://tunaos.org) | Self-hosted hub coordinating two ACMM **L5/L6** spokes across 43 repositories | Production | [#5773](https://github.com/hivecommons/hive/issues/5773) |
| [Project Bluefin](https://github.com/projectbluefin) | Self-managed hive over `projectbluefin/dakota` and related repos | Production | [docs.projectbluefin.io/factory](https://docs.projectbluefin.io/factory/); [#4928](https://github.com/hivecommons/hive/issues/4928), [#6449](https://github.com/hivecommons/hive/issues/6449), [#6450](https://github.com/hivecommons/hive/issues/6450) |
| [Frostyard](https://github.com/frostyard) | Autonomous development operations | Pre-production | |
| [Open Horizon](https://github.com/open-horizon) | Code-consistency standardization and enforcement | Pre-production | |
| [Open Horizon Services](https://github.com/open-horizon-services) | Consistency over community contributions | Pre-production | |
| [Danathar](https://github.com/Danathar) | Autonomous development operations | Pre-production | |

Two of these are worth a reviewer's attention specifically. **tunaos.org runs
at L5/L6 across 43 repositories** — the autonomy tiers at which agents can
merge — which is the strongest available evidence that the deny-rule and
attribution controls hold up outside the maintainers' own environment.
**Project Bluefin is independently documented by the adopter**, not merely
self-reported here.

What the project does *not* have is a written security case study — a
narrative of an incident, an attempted injection, or an adopter's own
security evaluation. That would be more useful to TAG-Security than an
adopter table, and its absence is a genuine gap rather than a formatting
choice. The earlier revision of this document said "Not applicable," which
was wrong on the facts and is corrected here.

### Related projects / vendors

Hive is the flagship of the [Hive Commons](https://hivecommons.dev) project
family and originated as a subproject of [KubeStellar](https://github.com/kubestellar)
(CNCF Sandbox), which remains its first production adopter. It interoperates
with, but is not a vendor dependency of: GitHub/GitHub Enterprise/GitLab/Gitea, and multiple AI CLI
backend vendors (Anthropic Claude Code, GitHub Copilot CLI, Google Gemini,
Block Goose, IBM Bob). See [landscape.md](https://github.com/hivecommons/hive/blob/v4/src/docs/landscape.md)
for a maintained comparison against nearby agentic-orchestration tools.

## Questions resolved since first review

The first revision of this document ended with a list of open questions. That
was a defensible way to ship a draft and an indefensible way to leave one:
TAG-Security review observed, correctly, that these were prompts to the
authors rather than findings for a reader. Each has been run down against the
repository and answered below. Where the answer is "no," it says no.

- **Does issue #3760 document a path to re-enabling SBOM/provenance?**
  Resolved. #3760 is **closed**, and it was never an SBOM decision: it is the
  bug report for `/usr/local/bin/hive` failing to exec with `EPERM` under
  containerd/k3s and rootless podman. Attestations force the artifact into an
  OCI image *index*, which permitted an overlayfs metacopy redirect that
  presented the binary as non-executable. Disabling them was the fix for a
  crash-loop, not a deprioritization of supply-chain metadata. The release
  workflow now generates SBOMs **out-of-band** and publishes them as GitHub
  Release assets, leaving the image manifest plain; the
  `check-no-image-attestations.sh` guard keeps future `docker.yml` changes
  from silently reattaching in-image attestations.

- **Does any external channel exist for security-relevant project
  communication?** Resolved: **no, by design.** `CONTRIBUTING.md` directs all
  design and review discussion to GitHub issues and PRs "so decisions remain
  public and searchable." There is no mailing list, Slack, Matrix or Discord
  for project decisions. Private security traffic goes through GitHub
  Security Advisories only. (The Slack/Discord/ntfy integrations in
  `notifications.md` are operator-alerting *features of the software* and are
  not project channels.)

- **Does a security-specific incident-response runbook exist?** Resolved:
  **yes, and it did not when this document was first written.**
  [`security-response.md`](https://github.com/hivecommons/hive/blob/v4/src/docs/security-response.md)
  documents who responds (the Maintainer Committee, which *is* the security
  team), how a report is triaged, and how disclosure is coordinated. It is
  distinct from the operational `HUB_DISASTER_RECOVERY.md`. Still missing: a
  post-incident-review practice and a user-notification SLA for a confirmed
  compromise.

- **Does the project intend to pursue the OpenSSF Best Practices Badge?**
  Resolved: **it already holds it.** [Project 14261](https://www.bestpractices.dev/projects/14261)
  reached the passing level on 2026-08-27, before this question was asked.
  The question survived in this document only because the document was being
  written from itself rather than checked against the programme's API.

- **What is the current OpenSSF Scorecard result?** Intentionally not frozen
  into this document — the workflow runs weekly and on every push to `v4`,
  and a reviewer should read the live result rather than a stale number.
  This is the one item from the original list that remains deliberately
  unanswered here, and the reason is that quoting it would make the document
  wrong on a schedule.

- **Has any informal security review or adversarial testing occurred?**
  Resolved: **no, and this is the most significant "no" in the list.** There
  is no red-team exercise, no adversarial evaluation, and no measured
  detection rate for `ioscan` anywhere in the repository — the search for one
  returned only the phrase used in unrelated design documents. The hardening
  work referenced by issue number throughout `security-threat-model.md` is
  maintainer-identified and maintainer-fixed, which is not the same thing as
  adversarial review. A red-team evaluation of the injection path is tracked
  in [#6685](https://github.com/hivecommons/hive/issues/6685). Until it
  exists, every efficacy claim about `ioscan` in this document should be read
  as unvalidated by design rather than validated by silence.

- **Will CODEOWNERS enforcement be enabled?** Resolved: the live `v4` branch
  protection currently requires **no** pull-request reviews at all, and
  enforcement was blocked on the fact that the project's own automation
  merges green PRs. The intended resolution is scoped enforcement over
  security-critical paths only, tracked in
  [#6687](https://github.com/hivecommons/hive/issues/6687) along with
  expanding `CODEOWNERS` from one owner to the three current maintainers.
  It is a commitment with an issue behind it, not a plan to remain advisory
  indefinitely.
