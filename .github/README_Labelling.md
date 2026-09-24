# TOC Labeling Guide

Labeling and review chat-ops in this repository are handled by
[cncf/prow-github-actions](https://github.com/cncf/prow-github-actions), configured in
[`.github/prow.yaml`](prow.yaml) and driven by [`.github/workflows/prow.yml`](workflows/prow.yml).

There are three labeling paths:

1. **Slash commands in comments** (manual, explicit)
2. **Automatic labeling** (`needs-*` rules and OWNERS-based path labels)
3. **Review and merge commands** (`/lgtm`, `/approve`, `/hold`)

## 1) Slash Commands in Comments

Write a command at the start of a line in a normal issue or PR comment.

- A command must start a line; one mentioned mid-sentence is ignored. Several commands may be
  placed on separate lines of the same comment.
- Format is `/command value [value ...]`. Values are case-insensitive.
- Commands inside code blocks or blockquotes are ignored.
- Every command has a `/remove-<command> value` form.
- A command only applies labels that already exist in the repository. If one is missing, run
  the label-sync job (Actions -> Prow -> Run workflow).

### Core triage/grouping

| Command | Values | Notes |
|---|---|---|
| `/kind <value>` | `dd`, `docs`, `election`, `enhancement`, `initiative`, `meeting`, `moving-levels`, `publication`, `review`, `subproject` | stacks; clears `needs-kind` |
| `/triage <value>` | `valid`, `needs-information`, `duplicate`, `not-planned` | exclusive; clears `needs-triage` |
| `/tag <value>` | `developer-experience`, `infrastructure`, `operational-resilience`, `security-and-compliance`, `workloads-foundation` | stacks; clears `needs-group` |
| `/sub <value>` | `contributor-strategy-and-advocacy`, `mentoring`, `project-reviews` | stacks; clears `needs-group` |
| `/label toc` | | plain `toc` label; clears `needs-group` |
| `/pub <value>` | `blog`, `tech-paper` | exclusive |
| `/level <value>` | `archived`, `graduation`, `incubation`, `sandbox` | exclusive |
| `/review <value>` | `governance`, `health`, `security`, `tech` | stacks |

### DD lifecycle

Labels in these families are `dd-<stage>/<value>` (for example `dd-status/in-progress`).

| Command | Values | Notes |
|---|---|---|
| `/dd-triage <value>` | `needs-adopters`, `needs-more-information`, `needs-security-assessment` | stacks |
| `/dd-status <value>` | `ready-for-assignment`, `in-progress`, `in-comment-period`, `in-voting`, `complete`, `waiting` | exclusive |
| `/dd-adopters <value>` | `not-started`, `in-progress`, `complete` | exclusive |
| `/dd-gov-review <value>` | `not-started`, `in-progress`, `complete` | exclusive |
| `/dd-tech-review <value>` | `not-started`, `in-progress`, `complete` | exclusive |
| `/dd-sec-review <value>` | `not-started`, `in-progress`, `complete` | exclusive |

When a DD moves to `ready-for-assignment`, the assigning TOC member should also clear any
`dd-triage/*` labels (`/remove-dd-triage <value>`) and set the four review trackers to
`not-started` (`/dd-adopters not-started`, `/dd-gov-review not-started`, and so on) in the
same comment. `dd/needs-triage` is applied with `/label dd/needs-triage` when a DD application
arrives and removed with `/remove-label dd/needs-triage` once triage is done.

### Initiatives and votes

| Command | Values | Notes |
|---|---|---|
| `/toc-initiative <value>` | `AI` | label is `toc-initiative/<value>` |
| `/init <value>` | `not-started`, `in-progress`, `complete`, `stale` | exclusive |
| `/vote <value>` | `nomination`, `open`, `closed` | exclusive |
| `/help` | | adds `help wanted` |
| `/good-first-issue` | | adds `good first issue` and `help wanted` |

### Other built-in commands

| Command | Who | Description |
|---|---|---|
| `/assign [@user ...]`, `/unassign [@user ...]` | anyone | assign or unassign (yourself if no user given) |
| `/cc [@user ...]`, `/uncc [@user ...]` | anyone | request or dismiss a review |
| `/auto-cc` | anyone, PRs only | request reviewers from the OWNERS files of the changed files |
| `/close [not-planned]`, `/reopen` | collaborators or the author | close or reopen |
| `/retitle <title>` | collaborators | rename the issue or PR |
| `/milestone <name>`, `/milestone clear` | collaborators | set or clear the milestone |
| `/lock [reason]` | collaborators | lock the conversation |
| `/remove <label> ...` | collaborators | remove any label by name |
| `/check-required-labels` | anyone | re-evaluate the `needs-*` rules |
| `/retest`, `/test all`, `/test <workflow>` | reviewers, PRs only | re-run GitHub Actions runs on the head commit |
| `/ok-to-test` | reviewers (not the author) | approve held workflow runs on a first-time contributor's fork PR |

The full command reference is in the
[prow-github-actions docs](https://github.com/cncf/prow-github-actions/blob/main/docs/commands.md).

## 2) Automatic Labeling

### `needs-*` helper labels

These are added when a required category is absent and removed as soon as a matching label
arrives, whether via a `/` command, the GitHub UI, or an OWNERS path label. Removing a
`needs-*` label by hand while the category is still missing re-adds it.

| Helper label | Removed when |
|---|---|
| `needs-kind` | Any `kind/*` label is present |
| `needs-triage` | Any `triage/*` label is present |
| `needs-group` | Any `toc`, `tag/*`, or `sub/*` label is present |

### Mutually exclusive label groups

Families marked exclusive above (`triage`, `pub`, `level`, `init`, `vote`, `dd-status`,
`dd-adopters`, `dd-gov-review`, `dd-tech-review`, `dd-sec-review`) replace any existing
label of the same family when applied via their command. Labels applied through the GitHub
UI are not de-duplicated.

### Path labels (PR-based)

`OWNERS` files in the directories below carry a `labels:` list. A PR touching files under
one of them gets that label when opened or updated, which in turn clears `needs-group`.

| Directory | Label applied |
|---|---|
| `tags/tag-developer-experience/` | `tag/developer-experience` |
| `tags/tag-infrastructure/` | `tag/infrastructure` |
| `tags/tag-operational-resilience/` | `tag/operational-resilience` |
| `tags/tag-security-and-compliance/` | `tag/security-and-compliance` |
| `tags/tag-workloads-foundation/` | `tag/workloads-foundation` |
| `toc_subprojects/contributor-strategy-and-advocacy-subproject/` | `sub/contributor-strategy-and-advocacy` |
| `toc_subprojects/mentoring-subproject/` | `sub/mentoring` |
| `toc_subprojects/project-reviews-subproject/` | `sub/project-reviews` |

## 3) Review and Merge

Review authority comes from the `OWNERS` files (root and per directory), which mirror
[`.github/CODEOWNERS`](CODEOWNERS). `.github/CODEOWNERS` remains the source for GitHub's own
required-review protections; the `OWNERS` files decide who may use the commands below.

| Command | Who | Effect |
|---|---|---|
| `/lgtm` | an OWNERS reviewer or approver of a changed file, not the author | adds `lgtm`, bound to the current head commit; a new push removes it |
| `/lgtm cancel` | same, or the author | removes `lgtm` |
| `/approve` | an OWNERS approver of a changed file | records approval; `approved` is added once every changed file is covered |
| `/approve cancel` | same | withdraws approval |
| `/hold` | anyone | adds `do-not-merge/hold` |
| `/hold cancel`, `/unhold` | anyone | removes the hold |

A PR merges automatically (merge commit) once it carries `lgtm` and `approved`, has no
`do-not-merge/*`, `needs-rebase` or `hold` label, and GitHub reports it mergeable
(required reviews and checks satisfied). A bot comment starting with `[APPROVALNOTIFIER]`
tracks who still needs to approve.

## Maintaining labels

Label names, colors and descriptions live in [`.github/prow.yaml`](prow.yaml). A push that
changes that file runs the label-sync job, which creates missing labels and fixes drifted
colors or descriptions. It never deletes or renames labels; do that in the repository's
label settings.
