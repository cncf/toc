# Hive Security Self-Assessment

This document follows the [CNCF TAG-Security self-assessment
template](https://tag-security.cncf.io/community/assessments/guide/self-assessment/)
and is submitted as a required artifact of Hive's CNCF Incubation application.

> **Scope:** This assessment covers **[kubestellar/hive](https://github.com/kubestellar/hive)** only — the AI agent orchestration system and its hub/spoke deployment. It does **not** cover [KubeStellar Core](https://github.com/kubestellar/kubestellar), the multi-cluster orchestration control plane, or [KubeStellar Console](https://github.com/kubestellar/console), which has its [own self-assessment](../../console/security-assessment/self-assessment.md).
>
> **Canonical source:** this document is maintained in the Hive repository at [`src/docs/security-self-assessment.md`](https://github.com/kubestellar/hive/blob/v4/src/docs/security-self-assessment.md). Corrections should be made there and mirrored here.

It complements, and deliberately does not duplicate, three existing security
documents in the Hive repository:

- [security-threat-model.md](https://github.com/kubestellar/hive/blob/v4/src/docs/security-threat-model.md) — attacker-oriented
  view: assets, trust boundaries, threat actors, defense layers, residual
  risks.
- [security-model.md](https://github.com/kubestellar/hive/blob/v4/src/docs/security-model.md) — operator/evaluator-oriented view:
  the seven enforcement layers (dashboard auth, spoke auth, credential
  isolation, sandboxing, GitHub blast-radius controls, hub↔spoke channel,
  hosted-platform isolation) with file/line evidence for each.
- [security.md](https://github.com/kubestellar/hive/blob/v4/src/docs/security.md) — log-scrubbing and secret-redaction
  implementation notes.

Where this document repeats a claim from those pages, it cites the same
evidence rather than re-deriving it. Sections below cite `path:line` against
the Hive repository's `v4` branch (`src/` prefix; commit `327969b0` at time of
writing) wherever a specific mechanism is asserted.

## Metadata

| | |
|---|---|
| Assessment Stage | Complete |
| Software | [kubestellar/hive](https://github.com/kubestellar/hive) |
| Security Provider | No — Hive is not itself a security product. It is an agent-orchestration platform whose core value proposition includes constraining the blast radius of the AI agents it runs; see [Overview](#overview) below. |
| Languages | Go (core: dashboard, hub, proxy, scheduler, agent orchestration — `src/go.mod`); JavaScript (dashboard UI, served inline, no separate SPA build — see `dashboard/`); Shell/Python (deterministic pipeline scripts under `src/bin/`, 45 scripts per [`bin/README.md`](https://github.com/kubestellar/hive/blob/v4/bin/README.md)) |
| SBOM | Not currently generated. Container image builds explicitly disable provenance/SBOM attestations (`sbom: false` in `.github/workflows/docker.yml:134,177,349,454`, citing issue #3760 as the reason). This is a known gap — see [Open questions](#open-questions--not-yet-assessed). |
| Security links | See table below |

| Doc | URL |
|---|---|
| Security policy / vulnerability reporting | [SECURITY.md](https://github.com/kubestellar/hive/blob/v4/SECURITY.md) |
| Threat model | [security-threat-model.md](https://github.com/kubestellar/hive/blob/v4/src/docs/security-threat-model.md) |
| Security model (operator guide) | [security-model.md](https://github.com/kubestellar/hive/blob/v4/src/docs/security-model.md) |
| Log scrubbing / secret redaction | [security.md](https://github.com/kubestellar/hive/blob/v4/src/docs/security.md) |
| ADRs (architecture decision records) | [adr/README.md](https://github.com/kubestellar/hive/blob/v4/src/docs/adr/README.md) |
| OpenSSF Scorecard workflow | [.github/workflows/scorecard.yml](https://github.com/kubestellar/hive/blob/v4/.github/workflows/scorecard.yml) |

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
[security-model.md](https://github.com/kubestellar/hive/blob/v4/src/docs/security-model.md): *"if a human would give the same
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
  (see [contributor-relay.md](https://github.com/kubestellar/hive/blob/v4/src/docs/contributor-relay.md)).
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
- **Log scrubbing** (`src/pkg/logscrub`, `security.md`) — redacts recognized
  GitHub token prefixes (`ghs_`, `ghp_`, `gho_`, `github_pat_`) and JWT-shaped
  strings from Hive's own structured log output. Explicitly pattern-based,
  not a general secret scanner; does not cover agent CLI stdout/terminal
  transcripts unless routed through Hive's logger.
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

- **OpenSSF Scorecard**: automated, runs weekly (`scorecard.yml`); results
  are public via the Scorecard badge/API. No specific score floor is gated
  in CI at present.
- **DCO (Developer Certificate of Origin)**: enforced for human contributors
  via `copilot-dco.yml` and is a stated policy requirement for agent-authored
  commits (`git commit -s`) per `security-model.md` Layer 5 — described there
  as *"policy-driven inside hive; enforce it repo-side (DCO check) for a hard
  guarantee,"* i.e. Hive asks agents to sign off but the hard gate is a
  repo-side branch-protection setting the operator must also enable.
- **License**: Apache License 2.0 (`LICENSE`), OSI-approved, CNCF-preferred.
- **No formal SBOM** is currently attached to release artifacts (see
  [Metadata](#metadata) and [Open questions](#open-questions--not-yet-assessed)).

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
  omission. Not yet assessed whether #3760 documents a planned reintroduction.
- **Single approver**: `OWNERS` lists one approver and one reviewer
  (`clubanderson`) for the whole repository at time of writing — see
  [Known weaknesses](#three-most-significant-known-weaknesses) below.
- **CODEOWNERS**: `.github/CODEOWNERS` covers security-sensitive paths
  (Dockerfiles, workflows, deploy manifests, launch scripts, key/cookie code)
  but is explicitly **advisory only** — "Require review from Code Owners" is
  not enabled in branch protection, a deliberate choice documented in
  `security-model.md:155` because the repository's automation merges green
  PRs without mandatory human review. This is a policy choice the project
  states plainly rather than obscures, but it is a real gap between what
  CODEOWNERS suggests and what branch protection enforces.

### Communication channels

- **Internal**: GitHub issues/PRs/discussions on the `kubestellar/hive`
  repository; the project is largely maintained by autonomous agent fleets
  operating under the same ACMM framework the software implements
  (i.e., Hive is substantially self-hosting — its own maintenance PRs are
  frequently agent-authored, subject to the same guardrails described in this
  document).
- **Inbound**: public GitHub issues/PRs are the primary inbound channel;
  private vulnerability reports go through GitHub Security Advisories per
  `SECURITY.md`.
- **Outbound**: release notes via `CHANGELOG.md`; no separate mailing list or
  Slack channel is asserted in this assessment (not verified — see
  [Open questions](#open-questions--not-yet-assessed)).

### Ecosystem

Hive is a KubeStellar-organization project and interoperates with, but does
not depend on: GitHub/GitHub Enterprise/GitLab/Gitea (via the "Forge App"
abstraction), several AI coding-agent CLI backends (Claude Code, GitHub
Copilot CLI, Gemini, Goose, Bob/bobshell, Agy), and self-hosted inference
gateways (LiteLLM, vLLM, llm-d, watsonx) reached only through an in-pod
credential-translating proxy (`security-model.md` Layer 3). It is
independent of any single cloud provider or AI vendor.

## Security issue resolution

### Responsible disclosures process

Vulnerability reports are handled through **GitHub private vulnerability
reporting** (Security tab → "Report a vulnerability"), not public issues, PRs,
or discussions, per [`SECURITY.md`](https://github.com/kubestellar/hive/blob/v4/SECURITY.md).
Reporters are asked for affected component/branch/commit, description and
impact, reproduction steps, and any supporting logs/PoC/config.

**Vulnerability response process**: the stated target is acknowledgement
within **5 business days** (`SECURITY.md` "What to Expect"), followed by
investigation/confirmation, a coordinated fix and disclosure timeline, and
optional reporter credit. There is no published CVSS-scoring policy, no
published maximum time-to-fix SLA, and no dedicated security response team —
response capacity is effectively the single listed approver
(`OWNERS: approvers: [clubanderson]`). This is stated plainly, not implied:
see [Known weaknesses](#three-most-significant-known-weaknesses).

### Incident response

No formal, published incident-response runbook specific to a security
incident (as distinct from operational incidents) was found in this
repository at assessment time. Related but not equivalent:
[`docs/HUB_DISASTER_RECOVERY.md`](https://github.com/kubestellar/hive/blob/v4/docs/HUB_DISASTER_RECOVERY.md)
covers hub-level disaster recovery (backup/restore, spoke fleet recovery,
operator communication), and the master-key rotation flow
(`security-model.md` "Master key rotation") gives an operator a mechanism to
revoke and replace compromised key material fleet-wide without a full
re-provision. Neither is a substitute for a documented incident-response
plan naming roles, communication SLAs, and post-incident review for a
security event specifically. This gap is recorded in
[Open questions](#open-questions--not-yet-assessed).

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
   [#2804](https://github.com/kubestellar/hive/issues/2804), which — per the
   same doc — also proposes moving live GitHub write credentials out of the
   agent sandbox entirely (*"current agents can still need live credentials
   to push/open PRs"*). Until #2804 lands, a compromised agent process that
   evades the network/token controls still executes inside the same
   container as every other agent on that spoke.

3. **Single-maintainer security response capacity, with advisory-only code
   ownership enforcement.** `OWNERS` lists exactly one approver and one
   reviewer for the entire repository (`clubanderson`) — verified in
   `OWNERS` at the root of the `v4` worktree. `SECURITY.md`'s 5-business-day
   acknowledgement target and the private-disclosure process are real
   commitments, but they rest on one person's availability with no stated
   backup or escalation path, and no published SLA for time-to-fix or
   time-to-disclosure beyond "we ask that you give us a reasonable
   opportunity to remediate." Compounding this,
   `.github/CODEOWNERS` — which does cover security-sensitive paths
   (Dockerfiles, CI workflows, deploy manifests, launch scripts, key/cookie
   handling code) — is **not enforced** by branch protection today; the
   project states this is deliberate because its own automation merges green
   PRs without mandatory human review (`security-model.md:155`). The net
   effect: a change to, say, the MITM proxy's deny-rule table or the SUID
   contract script can merge on green CI without a human security reviewer
   in the loop, mitigated only by the coverage/test gates and the fact that
   much of the review burden is itself carried by the same agent-review
   guardrails this document describes (a degree of "the system watching
   itself" that is architecturally interesting but not equivalent to
   independent human review capacity). This combination — one human in the
   loop for both ongoing code review and vulnerability response — is the
   single largest institutional (as opposed to technical) risk this
   assessment identifies, and is exactly the kind of gap CNCF Incubation
   review should weigh: it is a bus-factor and response-capacity risk, not a
   code defect.

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

An [OpenSSF Scorecard](.github/workflows/scorecard.yml) workflow runs weekly
against the repository. This assessment does not reproduce the current
numeric score here (it changes over time and is available live via the
Scorecard badge/API); a reviewer should pull the current score rather than
rely on a number frozen at the time of this document's writing. The project
has not pursued OpenSSF Best Practices Badge ("CII Best Practices")
certification at time of writing — not yet assessed whether this is planned.

### Case studies

Not applicable — Hive does not currently publish named case studies of
production deployments in the Hive repository. (Not the same claim as "Hive has
no production users" — simply that no case-study document exists to cite
here.)

### Related projects / vendors

Hive is part of the [KubeStellar](https://github.com/kubestellar) CNCF
Sandbox project family. It interoperates with, but is not a vendor
dependency of: GitHub/GitHub Enterprise/GitLab/Gitea, and multiple AI CLI
backend vendors (Anthropic Claude Code, GitHub Copilot CLI, Google Gemini,
Block Goose, IBM Bob). See [landscape.md](https://github.com/kubestellar/hive/blob/v4/docs/landscape.md)
for a maintained comparison against nearby agentic-orchestration tools.

## Open questions / not yet assessed

Items this assessment could not verify from the repository alone, listed
explicitly rather than guessed at:

- Whether issue #3760 (cited as the reason SBOM/provenance attestation is
  disabled on container builds) documents a planned path to re-enabling SBOM
  generation, or is a permanent decision.
- Whether any external channel (mailing list, Slack, Matrix) exists for
  security-relevant project communication beyond GitHub issues/PRs and
  private security advisories.
- Whether a formal, security-specific incident-response runbook (as distinct
  from the disaster-recovery runbook at `docs/HUB_DISASTER_RECOVERY.md`)
  exists outside the Hive repository (e.g., in an internal-only document) —
  none was found in the public repository.
- Whether the project intends to pursue OpenSSF Best Practices Badge
  certification.
- The current numeric OpenSSF Scorecard result — intentionally not quoted
  here since it changes independently of this document; consult the live
  badge/workflow output.
- Whether any prior informal security review (e.g., a maintainer's own
  adversarial testing beyond what is documented in the ADRs and threat
  model) occurred without producing a public artifact. This assessment
  treats "no third-party audit" as accurate based on the absence of any
  audit report in the repository, but cannot rule out unpublished internal
  review.
- Whether `.github/CODEOWNERS` enforcement (branch-protection "Require
  review from Code Owners") is planned to be enabled once the project has
  more than one active human maintainer, or is intended to remain advisory
  indefinitely regardless of maintainer count.
