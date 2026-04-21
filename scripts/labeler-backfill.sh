#!/usr/bin/env bash
#
# labeler-backfill.sh
#
# Re-runs the auto-labeler (ghcr.io/cncf/gha-labeler) against every open PR
# and every open issue in cncf/toc. Useful when:
#   - The labeler workflow failed for old PRs (e.g. fork PRs pinned to an old
#     workflow file via GitHub's "Re-run jobs" behavior).
#   - Labels have been added/changed in .github/labels.yaml and existing open
#     items need to be reprocessed.
#
# This script calls the exact same container image that the GitHub Actions
# workflow uses (.github/workflows/labeler.yaml), so behavior matches 1:1.
#
# Prereqs:
#   - gh CLI authenticated with write access to the target repo (gh auth status)
#   - docker daemon running locally
#   - jq installed
#
# Usage:
#   chmod +x scripts/labeler-backfill.sh
#   DRY_RUN=1 ./scripts/labeler-backfill.sh   # preview only (no API writes)
#   ./scripts/labeler-backfill.sh             # apply labels
#
# Overridable env vars:
#   OWNER       (default: cncf)
#   REPO        (default: toc)
#   LABELS_URL  (default: main branch labels.yaml)
#   IMAGE       (default: ghcr.io/cncf/gha-labeler:main)
#   DRY_RUN     (default: 0; set to 1 to preview)
#   SCOPE       (default: all; one of: all, prs, issues)

set -euo pipefail

OWNER=${OWNER:-cncf}
REPO=${REPO:-toc}
LABELS_URL=${LABELS_URL:-https://raw.githubusercontent.com/cncf/toc/refs/heads/main/.github/labels.yaml}
IMAGE=${IMAGE:-ghcr.io/cncf/gha-labeler:main}
DRY_RUN=${DRY_RUN:-0}
SCOPE=${SCOPE:-all}

for bin in gh docker jq; do
  if ! command -v "$bin" >/dev/null 2>&1; then
    echo "error: '$bin' is required but not installed" >&2
    exit 1
  fi
done

if [[ -z "${GITHUB_TOKEN:-}" ]]; then
  GITHUB_TOKEN=$(gh auth token)
fi
export GITHUB_TOKEN

if [[ "$DRY_RUN" != "1" ]]; then
  echo "Pulling labeler image: $IMAGE"
  docker pull --quiet "$IMAGE" >/dev/null
fi

run_labeler() {
  local number=$1 body=$2 files=${3:-}
  if [[ "$DRY_RUN" == "1" ]]; then
    local preview=${body//$'\n'/ }
    printf '  [dry-run] #%s files=%q body=%q\n' \
      "$number" "$files" "${preview:0:120}"
    return
  fi
  docker run --rm -e GITHUB_TOKEN "$IMAGE" \
    "$LABELS_URL" "$OWNER" "$REPO" "$number" "$body" "$files"
}

label_prs() {
  echo "== Open PRs in $OWNER/$REPO =="
  local numbers
  numbers=$(gh pr list --repo "$OWNER/$REPO" --state open --limit 500 \
    --json number --jq '.[].number')
  if [[ -z "$numbers" ]]; then
    echo "  (none)"
    return
  fi
  while read -r n; do
    [[ -z "$n" ]] && continue
    echo "-> PR #$n"
    local body files
    body=$(gh pr view "$n" --repo "$OWNER/$REPO" --json body --jq '.body // ""')
    files=$(gh pr view "$n" --repo "$OWNER/$REPO" --json files \
      --jq '[.files[].path] | join(",")')
    run_labeler "$n" "$body" "$files"
  done <<< "$numbers"
}

label_issues() {
  echo "== Open Issues in $OWNER/$REPO =="
  local numbers
  numbers=$(gh issue list --repo "$OWNER/$REPO" --state open --limit 500 \
    --json number --jq '.[].number')
  if [[ -z "$numbers" ]]; then
    echo "  (none)"
    return
  fi
  while read -r n; do
    [[ -z "$n" ]] && continue
    echo "-> Issue #$n"
    local body
    body=$(gh issue view "$n" --repo "$OWNER/$REPO" --json body --jq '.body // ""')
    run_labeler "$n" "$body" ""
  done <<< "$numbers"
}

case "$SCOPE" in
  all)    label_prs; label_issues ;;
  prs)    label_prs ;;
  issues) label_issues ;;
  *) echo "error: SCOPE must be one of: all, prs, issues (got: $SCOPE)" >&2; exit 2 ;;
esac

echo "Done."
