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
  - values: `dd`, `docs`, `enhancement`, `initiative`, `meeting`, `review`, `subproject`
- `/triage <value>`
  - values: `valid`, `needs-information`, `duplicate`, `not-planned`
- `/tag <value>`
  - values: `developer-experience`, `infrastructure`, `operational-resilience`, `security-and-compliance`, `workloads-foundation`
- `/sub <value>`
  - values: `contributor-strategy-and-advocacy`, `mentoring`, `project-reviews`
- `/toc`

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
- Keeps label state clean by removing/replacing mutually exclusive labels in command-driven flows.

In short: use `/` commands when you want explicit control, and rely on automatic labeling for baseline triage and path-based routing.
