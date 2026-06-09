# TOC Labeling Guide

This repository has two labeling paths:

1. **Slash commands in comments** (manual, explicit)
2. **Automatic labeling** (event and file-path based)

## 1) Slash Commands in Comments

Use slash commands in a normal issue/PR comment to apply or remove labels.

- Command must be on the **first line** of the comment.
- Format is usually `/command value` (or `/command` for commands with no argument).
- Commands are processed for issue comments and PR conversation comments.

### Valid commands

#### Core triage/grouping

- `/kind <value>`
  - values: `dd`, `docs`, `enhancement`, `initiative`, `meeting`, `publication`, `review`, `subproject`
- `/triage <value>`
  - values: `valid`, `needs-information`, `duplicate`, `not-planned`
- `/tag <value>`
  - values: `developer-experience`, `infrastructure`, `operational-resilience`, `security-and-compliance`, `workloads-foundation`
- `/sub <value>`
  - values: `contributor-strategy-and-advocacy`, `mentoring`, `project-reviews`
- `/pub <value>`
  - values: `blog`, `tech-paper`
  - applying one removes any other `pub/*` label (mutually exclusive)
- `/toc`
- `/level <value>`
  - values: `archived`, `graduation`, `incubation`, `sandbox`
  - applying one removes any other `level/*` label (mutually exclusive)

#### DD lifecycle

- `/dd/triage <value>`
  - values: `needs-triage`, `needs-adopters`, `needs-more-information`, `needs-security-assessment`
- `/dd/status <value>`
  - values: `ready-for-assignment`, `in-progress`, `in-comment-period`, `in-voting`, `complete`, `waiting`
- `/dd/adopters <value>`
  - values: `not-started`, `in-progress`, `complete`
- `/dd/gov-review <value>`
  - values: `not-started`, `in-progress`, `complete`
- `/dd/tech-review <value>`
  - values: `not-started`, `in-progress`, `complete`
- `/dd/sec-review <value>`
  - values: `not-started`, `in-progress`, `complete`

#### Initiatives and votes

- `/toc/initiative <value>`
  - values: `AI`
- `/init <value>`
  - values: `not-started`, `in-progress`, `complete`, `stale`
- `/vote <value>`
  - values: `open`, `closed`, `nomination`
- `/help`

#### Remove commands

- `/remove-kind <value>`
- `/remove-triage <value>`
- `/remove-tag <value>`
- `/remove-sub <value>`
- `/remove-pub <value>`
- `/remove-toc`
- `/remove-toc/initiative <value>`
- `/remove-init <value>`
- `/remove-vote`
- `/remove-dd/triage <value>`
- `/remove-review <value>`
- `/remove-level <value>`
- `/remove-help`

## 2) Automatic Labeling (Overview)

Automatic labeling runs from GitHub Actions and handles broad classification without requiring manual `/` commands.

At a high level it does the following:

- Applies labels when issues/PRs are opened or edited.
- Uses changed file paths in PRs to infer group labels (for example, TOC/TAG/subproject areas).
- Maintains helper labels like `needs-kind`, `needs-triage`, and `needs-group` when required categories are missing.
- Keeps label state clean by removing/replacing mutually exclusive labels in both command-driven and state-driven flows.

### `needs-*` helper labels

These labels are added automatically when a required category is absent, and removed automatically when the corresponding label is present — whether applied via a `/` command, the GitHub UI, or a file-path rule.

| Helper label | Removed when |
|---|---|
| `needs-triage` | Any `triage/*` label is present |
| `needs-kind` | Any `kind/*` label is present |
| `needs-group` | Any `toc`, `tag/*`, or `sub/*` label is present |
| `dd/needs-triage` | Any `dd/triage/*` label is present |

### Mutually exclusive label groups

The following label groups enforce mutual exclusivity automatically. When a label in the group is applied (by any means), conflicting labels in the same group are removed.

| Group | Labels |
|---|---|
| `contribution-agreement` | `contribution-agreement/signed` ↔ removes `contribution-agreement/unsigned` (and vice versa) |
| `level/*` | `level/archived`, `level/graduation`, `level/incubation`, `level/sandbox` — applying one removes the others |
| `pub/*` | `pub/blog`, `pub/tech-paper` — applying one removes the others |
| `triage/*` | Enforced via `/triage` command |
| `kind/*` | Enforced via `/kind` command |
| `dd/status/*` | Enforced via `/dd/status` command |
| `init/*` | Enforced via `/init` command |
| `vote/open` + `vote/closed` | Enforced via `/vote` command |

### File path rules (PR-based)

When a PR modifies files in the paths below, `needs-group` is removed and the corresponding group label is applied automatically.

| Path pattern | Label applied |
|---|---|
| `tags/*/charter.md` | `toc` |
| `tags/tag-developer-experience/*` | `tag/developer-experience` |
| `tags/tag-infrastructure/*` | `tag/infrastructure` |
| `tags/tag-operational-resilience/*` | `tag/operational-resilience` |
| `tags/tag-security-and-compliance/*` | `tag/security-and-compliance` |
| `tags/tag-workloads-foundation/*` | `tag/workloads-foundation` |
| `toc_subprojects/contributor-strategy-and-advocacy/*` | `sub/contributor-strategy-and-advocacy` |
| `toc_subprojects/mentoring/*` | `sub/mentoring` |
| `toc_subprojects/project-reviews/*` | `sub/project-reviews` |

In short: use `/` commands when you want explicit control, and rely on automatic labeling for baseline triage, path-based routing, and mutual exclusivity enforcement.
