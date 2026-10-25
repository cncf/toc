# Hive Security Self-Assessment

Last reviewed: 2026-10-23

This document follows the [CNCF TAG-Security self-assessment template](https://tag-security.cncf.io/community/assessments/guide/self-assessment/). It supports Hive Commons' CNCF Sandbox application ([cncf/sandbox#516](https://github.com/cncf/sandbox/issues/516)).

> **Scope:** [hivecommons/hive](https://github.com/hivecommons/hive) only: the AI agent orchestration system and its hub/spoke deployment. Hive began as a KubeStellar subproject and moved to its own vendor-neutral org on 2026-09-03. KubeStellar Core and KubeStellar Console are not covered; Console has its [own self-assessment](https://github.com/cncf/toc/blob/main/projects/kubestellar/sub-projects/console/security-assessment/self-assessment.md).
>
> **Canonical source:** [`src/docs/security-self-assessment.md`](https://github.com/hivecommons/hive/blob/v5/src/docs/security-self-assessment.md) in the Hive repository. Corrections go there and are mirrored here.

## Table of contents

* [Metadata](#metadata)
  * [Security links](#security-links)
* [Overview](#overview)
  * [Background](#background)
  * [Actors](#actors)
  * [Actions](#actions)
  * [Goals](#goals)
  * [Non-goals](#non-goals)
* [Self-assessment use](#self-assessment-use)
* [Security functions and features](#security-functions-and-features)
* [Project compliance](#project-compliance)
* [Secure development practices](#secure-development-practices)
* [Security issue resolution](#security-issue-resolution)
* [Appendix](#appendix)

## Metadata

|   |   |
| -- | -- |
| Assessment Stage | Complete |
| Software | <https://github.com/hivecommons/hive> |
| Security Provider | No. Hive runs AI coding agents; its security controls exist to limit what those agents can do. |
| Languages | Go (core), JavaScript (dashboard), Shell and Python (pipeline scripts) |
| SBOM (software bill of materials) | SPDX-format JSON files for the `hive`, `hive-contributor` and `hive-hub` images are attached to each tagged [GitHub release](https://github.com/hivecommons/hive/releases). See [releases.md](https://github.com/hivecommons/hive/blob/v5/src/docs/releases.md#software-bill-of-materials-sbom). |

### Security links

| Doc | url |
| -- | -- |
| Security policy (how to report) | <https://github.com/hivecommons/hive/blob/v5/SECURITY.md> |
| Who responds and how | <https://github.com/hivecommons/hive/blob/v5/src/docs/security-response.md> |
| Threat model | <https://github.com/hivecommons/hive/blob/v5/src/docs/security-threat-model.md> |
| Security model (every control, with evidence) | <https://github.com/hivecommons/hive/blob/v5/src/docs/security-model.md> |
| Operator hardening guide | <https://github.com/hivecommons/hive/blob/v5/src/docs/securing-your-hive.md> |
| Input scanner (`ioscan`) and its defaults | <https://github.com/hivecommons/hive/blob/v5/src/docs/ioscan.md> |
| Input scanner red-team results | <https://github.com/hivecommons/hive/blob/v5/src/docs/ioscan-red-team.md> |

## Overview

Hive runs fleets of AI coding agents that maintain GitHub repositories on their own: they triage issues, write fixes, open pull requests, review code and, at the highest setting, merge. Because the agents read public text written by strangers and hold real GitHub credentials, Hive's design puts the limits on what an agent can do into infrastructure (a network proxy, token scopes, file permissions) rather than into instructions the model is asked to follow.

### Background

An *agent* here is a command-line AI tool such as Claude Code, GitHub Copilot CLI, Codex or Goose. Hive starts these tools as processes inside one container, gives each a task built from GitHub issues and pull requests, and lets them act on a repository through GitHub's API.

An operator chooses one of six autonomy levels, called **ACMM L1–L6** (Agentic Code Maturity Model). Each level maps every agent to one of four modes: *advisory* (observe and comment), *measured* (file issues), *hold-gated* (open pull requests labelled `hold` for a human) and *full* (merge on green CI). At L1–L2 no agent writes code. From L3 some agents may open pull requests. Only L6 agents may merge.

Deployments come in two shapes. A **spoke** is one Hive instance governing one or a few repositories. A **hub** is an optional central service that registers many spokes, shows fleet status and, for hosted spokes, provisions them and delivers their credentials.

Anyone with a GitHub account can write an issue or comment that an agent will later read. There is no step where a human must read that text first. This is the central risk the design works around.

### Actors

Each actor below runs separately from the others; the isolation between them is what stops a problem in one from spreading.

* **Agent processes.** Each agent runs under its own Unix user ID inside the spoke container, in its own `tmux` session, with its own GitHub token file that other agents cannot read. Agents share one container and one kernel; there is no per-agent container or virtual machine.
* **Egress proxy.** All HTTPS traffic from agent processes is forced through an in-container proxy (using `iptables`, which requires the `CAP_NET_ADMIN` capability). The proxy identifies the calling agent by the socket's user ID, not by anything the agent claims, and applies a rule table for that agent's level. If the proxy cannot be set up, the spoke refuses to start.
* **Hive core.** A Go process that schedules agents, runs the pre-kick pipeline (see Actions), serves the dashboard and performs the GitHub writes that agents are not allowed to make themselves. It holds the GitHub App key and mints the short-lived, per-agent tokens.
* **Dashboard and API.** Browser access with four roles: read, read-write, merger, owner. Sessions and hub single sign-on are verified with Ed25519 signatures (a public-key scheme); a spoke with no verification key returns HTTP 503 rather than falling back to a weaker check.
* **Hub.** Holds a master secret that never leaves it. Each spoke receives only keys derived from that secret for its own hive ID, plus the hub's public key. Hosted spokes receive their own GitHub App key and nothing belonging to other tenants.
* **Contributor relay.** An external person's machine that lends its AI subscription to a hive. It is outside the trust boundary. It authenticates with a registration token whose SHA-256 hash is stored on the hive, receives a 55-minute GitHub token scoped to one task, and its output enters the normal pull-request review path.
* **External services.** GitHub (or GitHub Enterprise) and the AI inference providers. For self-hosted inference gateways the real key lives in an in-container translator proxy and the agent receives only a placeholder.

Where credentials live: on a hosted spoke the GitHub App key and inference keys are in a Kubernetes Secret mounted read-only with mode 0440; on the hub, one 0600 file per cluster; on a self-hosted spoke, wherever the operator's `hive.yaml` points.

### Actions

1. **Select work.** Shell scripts enumerate open issues and pull requests, classify them and compute merge eligibility before any model is involved.
2. **Scan input.** `ioscan` normalizes the text (removing invisible Unicode characters and look-alike letters), undoes common encodings and checks the whole thread for known prompt-injection phrasing, dangerous instructions and secret-shaped strings. Text from strangers is withheld when any rule fires. At L1–L4 a critical finding is redacted and the task proceeds; at L5–L6 the task is blocked. A distinct canary string is planted in each prompt.
3. **Start the agent.** Hive writes a token for the agent's mode into its private file and launches the tool with the task prompt. Advisory agents get read-only tokens; measured agents can write issues but not code; hold-gated and full agents can write code. The full App token is never placed in an agent's environment.
4. **Mediate GitHub calls.** Every request passes the proxy. The rule table allows or denies each REST method and path, and each GraphQL mutation, for the agent's level; anything unmatched is denied. Creating a pull request (`POST /pulls`) and merging one (`PUT /pulls/{n}/merge`) are denied for every level.
5. **Perform privileged writes.** To open or merge a pull request, the agent drops a request file. Hive core reads the file's owner (not its contents) to identify the agent, re-checks eligibility, and performs the write as the GitHub App. Every such write carries a visible trailer naming the agent, backend and model, and is appended to an audit log.
6. **Scan output.** The proxy scans outbound GitHub request bodies for the canary, after decoding base64, hex, percent-encoding, reversed and split spellings. A hit is logged and, at L5–L6, blocked.
7. **Talk to the hub.** A spoke sends operational status (never credentials) using a per-hive bearer token. The hub replies with configuration, signed with Ed25519 and bound to the hive ID and a sequence number. By default a spoke logs a bad signature and still applies the response; rejection requires `HIVE_HEARTBEAT_VERIFY=enforce`.
8. **Upgrade.** A spoke pulls its target version from the hub and patches its own Kubernetes Deployment through a service account that can only `get` and `patch` that one object.

### Goals

* An agent cannot merge, or open a pull request directly, at any autonomy level, whatever its prompt says.
* An agent holds only the GitHub permissions its level needs.
* Every write Hive makes on an agent's behalf is attributable to a specific agent, backend and model, and is logged.
* One operator-chosen level sets prompt, token scope, tool restrictions and proxy rules together.
* Untrusted GitHub text is filtered before it reaches a prompt, and leaked prompt content is detected on the way out.

### Non-goals

* Hive does not make model output trustworthy and does not claim to stop prompt injection. `ioscan` reduces it.
* Hive does not isolate agents from each other at the container or kernel level.
* Hive does not inspect network traffic to destinations other than the GitHub API (and `api.linear.app`); other traffic is tunneled unread.
* Hive does not secure the cluster, container runtime or AI provider accounts it runs on.

The first three are areas we are investigating and may improve over time.

## Self-assessment use

This self-assessment is created by the Hive team to perform an internal analysis of the project's security. It is not intended to provide a security audit of Hive, or function as an independent assessment or attestation of Hive's security health.

This document serves to provide Hive users with an initial understanding of Hive's security, where to find existing security documentation, Hive plans for security, and general overview of Hive security practices, both for development of Hive as well as security of Hive.

This document provides the CNCF TAG-Security with an initial understanding of Hive to assist in a joint-assessment, necessary for projects under incubation. Taken together, this document and the joint-assessment serve as a cornerstone for if and when Hive seeks graduation and is preparing for a security audit.

## Security functions and features

### Critical

* **Egress proxy with per-level rule table.** Forced `iptables` redirect; deny by default; pull-request creation and merge denied for every level, for both REST and GraphQL. This is the control that bounds the damage when every other one fails.
* **ACMM level gating.** One level drives prompt templates, token mode, CLI tool-deny flags and proxy rules.
* **Per-agent, per-mode GitHub tokens.** Minted by Hive core from the App key; read-only, issues-only or code-write depending on mode.
* **Hive-mediated pull-request creation and merge.** Performed by Hive core as the App bot after re-checking eligibility; the agent's identity comes from file ownership.
* **Attribution and audit log.** Trailer on every write plus an append-only JSON log. The log rotates at 5 MB with three backups and a 90-day cap, so on a busy hive the lookback can be shorter than 90 days.
* **Ed25519-only sessions and single sign-on.** No symmetric fallback; missing key fails closed.
* **Per-user-ID agent separation and setuid contract.** `su-exec` is the only setuid binary Hive ships (mode 4750); a build-time check fails the image on any other setuid file, and inherited world-executable setuid binaries from the base image are stripped.

### Security relevant

* **`ioscan` input scanner.** On by default. It uses fixed rules to catch prompt-injection text in issues, comments and files before an agent reads them. It reads a whole issue or PR thread as one unit, so an attack split across comments is caught. It undoes common disguises (base64, hex, HTML entities, rot13) before matching, and a bug report that only quotes an attack phrase is no longer redacted. For text from strangers, any rule that fires withholds the text; the hive's own accounts and repo members get more leeway. In our own test of 63 attack samples it withheld 97%, and it wrongly redacted one of twelve harmless samples. Rules match phrasings, so a new paraphrase can still get through. From L4 up a model-based classifier also scores the text; at L5–L6 the text is withheld if that classifier cannot answer. Nobody outside the project has tested any of this yet.
* **Canaries.** On by default. Each task gets its own fake secret planted in the prompt. If that string shows up in anything the agent sends out, the agent was tricked into leaking it; Hive logs the event and, at L5–L6, blocks the request. The check undoes the encodings it knows about (base64, hex, percent, reversed, split), so a new encoding gets through. `git push` content cannot be scanned; it is refused outright only at the two strictest levels (L5–L6).
* **`fail_mode`.** `open` (redact) at L1–L4, `closed` (block) at L5–L6; an explicit setting overrides either. `closed` turns every false positive into a stalled task.
* **Trajectory review.** A second model compares an agent's stated intent with its recent transcript and can pause it. Advisory only; fails open.
* **Token budget.** A rolling budget (default seven days) stops new tasks when exhausted.
* **Log scrubbing.** Redacts token prefixes, web tokens, cloud key IDs, bearer headers, private-key blocks and canaries from Hive's own logs. This protects pasted logs and support bundles; it is not an exfiltration control and does not see agent traffic.
* **Key rotation.** The hub master key rotates by generation with a seven-day window in which both generations are accepted.
* **Supply chain.** Base images pinned by digest; GitHub Actions pinned by commit SHA; `npm install --ignore-scripts`; checksum-verified downloads for tmux, ttyd, `gh`, goose and `su-exec`. Image attestations are turned off because they forced a multi-architecture image index that crashed the binary under some runtimes ([#3760](https://github.com/hivecommons/hive/issues/3760)); SBOMs ship as release files instead.
* **One image, every backend.** The spoke image bundles all supported CLIs so an operator can switch backends by configuration. The cost is a large image surface. A slimmer `Dockerfile.contributor` exists for the relay path.

## Project compliance

Hive holds no formal certification (SOC 2, ISO 27001, FIPS). It holds the [OpenSSF Best Practices badge](https://www.bestpractices.dev/projects/14261) at the passing level, runs OpenSSF Scorecard weekly, requires a Developer Certificate of Origin sign-off on every commit, and is licensed under Apache-2.0.

## Secure development practices

### Development pipeline

* `v5` is the maintained branch. Branch protection requires eleven status checks before merge: `gate`, `DCO` (Developer Certificate of Origin sign-off), `build-and-test`, `test`, `coverage`, `golangci-lint`, `gosec`, `govulncheck`, a changelog-fragment guard, a PR-contents check and a sign-off-survives-squash check. Administrators are not exempt.
* Merge approval is handled by Hive, not by a GitHub setting. Hive's agents merge a pull request only after tests pass and Hive's own rules allow it: the autonomy level set by the operator, review by a separate reviewer agent, and holds that route risky changes to a human. A `CODEOWNERS` file names three maintainers for sensitive files (container builds, CI workflows, deploy manifests, proxy rules, key and session code). GitHub's "require code-owner review" switch is off, so those rules live in Hive rather than in GitHub. We are reviewing whether to also turn the GitHub switch on for those files.
* `OWNERS` lists five approvers: three from IBM, one from Universal Blue, one independent.
* Tests run with the race detector on every pull request; an hourly job runs the full suite and files an issue on coverage regression.
* Signed commits are not required. Container images are not signed.
* The project is enrolled in [Anthropic's OSS Scanner](https://red.anthropic.com/oss-scanner/), which builds `v5` and emails findings privately to maintainers.

### Communication channels

* **Internal.** GitHub issues and pull requests on `hivecommons/hive`. Decisions are made in those public threads. Day-to-day discussion happens on the project [Discord](https://hivecommons.dev/discord) and the [hivecommons-dev Google Group](https://groups.google.com/g/hivecommons-dev), which also announces the community meeting.
* **Inbound.** Public GitHub issues; private vulnerability reports through GitHub's "Report a vulnerability" form.
* **Outbound.** GitHub releases and `CHANGELOG.md`.

### Ecosystem

Hive works with GitHub and GitHub Enterprise (GitLab and Gitea adapters exist but are not production-ready), takes work from GitHub Issues, GitHub Projects, Linear and Jira (a core design goal is to be work-source agnostic, so any tracker can feed the same agents), with the Claude Code, GitHub Copilot, Codex, Antigravity, Goose and Bob CLIs, and with self-hosted inference gateways such as LiteLLM, vLLM, llm-d and watsonx. It is used in production by Hive Commons itself (Hive maintains the `hivecommons/hive` repository and the other Hive Commons projects), KubeStellar Console, tunaos.org and Project Bluefin (see Case studies). Hive is a CNCF Sandbox candidate ([cncf/sandbox#516](https://github.com/cncf/sandbox/issues/516)).

## Security issue resolution

### Responsible disclosures process

Reporters use GitHub private vulnerability reporting on the repository's Security tab, or contact a maintainer directly if that is unavailable ([SECURITY.md](https://github.com/hivecommons/hive/blob/v5/SECURITY.md)). The project commits to acknowledge within five business days, to coordinate a disclosure date with the reporter, to publish a GitHub Security Advisory when a fix ships, and to credit the reporter if they wish.

### Vulnerability response process

The Maintainer Committee is the security response team; whoever picks up the advisory drives it and involves the others for severity and timing ([security-response.md](https://github.com/hivecommons/hive/blob/v5/src/docs/security-response.md)). Severity is set by maintainer judgment; there is no formal severity-scoring policy. Fixes are developed in a private advisory fork. Publicly known vulnerabilities are to be fixed within 60 days. There is no time-to-fix commitment for private reports.

### Incident response

For a credential compromise the containment step is hub master-key rotation, which re-derives every spoke's keys without re-provisioning. Hub backup, restore and fleet recovery are in [HUB_DISASTER_RECOVERY.md](https://github.com/hivecommons/hive/blob/v5/docs/HUB_DISASTER_RECOVERY.md). The project has no written post-incident review practice and no commitment on how quickly users are told of a confirmed compromise.

## Appendix

### Known issues over time

No CVEs have been filed against Hive, and no outside audit or penetration test has been done yet. All security fixes so far were found by maintainers.

Fixed in 2026:

* The canary leak check matched only exact strings and missed base64-encoded leaks ([#6686](https://github.com/hivecommons/hive/issues/6686)).
* Two Unicode characters each defeated the whole hidden-text (steganography) rule family ([#6714](https://github.com/hivecommons/hive/issues/6714)).
* The input scanner's fixed rules withheld only 42% of a red-team sample set. Reading threads as one unit, a stricter line for strangers, undoing encodings and new paraphrase rules raised that to 97% ([#6685](https://github.com/hivecommons/hive/issues/6685)).
* Hub-to-spoke configuration pushes relied on transport encryption alone; they are now signed, with enforcement opt-in ([#7082](https://github.com/hivecommons/hive/issues/7082)).
* An earlier hosted-spoke design sent every tenant's App key to every spoke (removed).
* When browsers stopped trusting two public root certificates in October 2026 before Linux packages caught up, Hive's images were changed to reject them at build time, and CI now checks this ([#13230](https://github.com/hivecommons/hive/pull/13230)).

Still open:

* Agents share one container by default. An opt-in sandbox mode runs each task in its own rootless container with no credentials, but it is not the default. We are evaluating stronger isolation options, including OpenShell.
* Contributor registration tokens never expire. A contributor can replace their own token, but nothing forces rotation ([#13804](https://github.com/hivecommons/hive/issues/13804)).

### OpenSSF best practices

Passing badge held since 2026-08-27 ([project 14261](https://www.bestpractices.dev/projects/14261), 100% of passing criteria). Silver and gold are not held (15% and 22% of their criteria met). Scorecard runs weekly from [scorecard.yml](https://github.com/hivecommons/hive/blob/v5/.github/workflows/scorecard.yml).

### Case studies

* **tunaos.org** runs a self-hosted hub with two spokes at L5/L6 across 43 repositories, the levels at which agents may merge ([#5773](https://github.com/hivecommons/hive/issues/5773)).
* **Project Bluefin** runs a self-managed hive over `projectbluefin/dakota` and related repositories, documented by the adopter at [docs.projectbluefin.io/factory](https://docs.projectbluefin.io/factory/).
* **KubeStellar Console** is maintained end to end by Hive agents: triage, fixes, review and merge.

Four further adopters are at pre-production maturity ([ADOPTERS.md](https://github.com/hivecommons/hive/blob/v5/ADOPTERS.md)). No adopter has published a security evaluation of Hive.

### Related projects / vendors

Prospective users ask how Hive differs from GitHub Agentic Workflows, OpenAI Symphony, Fullsend, and single-agent services such as the GitHub Copilot coding agent. Those tools are built around agent runs tied to a GitHub Actions event, a single issue or a single repository. Hive runs long-lived fleets across many repositories and enforces write limits at the network boundary, whichever agent CLI is in use. A dimension-by-dimension comparison is in [landscape.md](https://github.com/hivecommons/hive/blob/v5/src/docs/landscape.md).
