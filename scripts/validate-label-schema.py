#!/usr/bin/env python3
"""
Validate that all labels referenced by issue templates are defined in
.github/labels.yaml (labels section only).
"""

from __future__ import annotations

from dataclasses import dataclass
from pathlib import Path
import re
import sys


ROOT = Path(__file__).resolve().parents[1]
LABELS_FILE = ROOT / ".github" / "labels.yaml"
TEMPLATES_DIR = ROOT / ".github" / "ISSUE_TEMPLATE"


@dataclass
class LabelRef:
    file: Path
    line: int
    label: str


def read_defined_labels(path: Path) -> set[str]:
    text = path.read_text(encoding="utf-8")
    labels_section = text.split("\nruleset:\n", 1)[0]
    return set(re.findall(r"^- name: (.+)$", labels_section, flags=re.M))


def parse_labels_block(lines: list[str], start_idx: int, file: Path) -> list[LabelRef]:
    refs: list[LabelRef] = []
    labels_indent = len(lines[start_idx]) - len(lines[start_idx].lstrip(" "))
    idx = start_idx + 1
    while idx < len(lines):
        raw = lines[idx]
        stripped = raw.strip()
        if not stripped:
            idx += 1
            continue
        current_indent = len(raw) - len(raw.lstrip(" "))
        if current_indent <= labels_indent:
            break
        match = re.match(r"^\s*-\s*(.+?)\s*$", raw)
        if match:
            refs.append(LabelRef(file=file, line=idx + 1, label=match.group(1).strip()))
        idx += 1
    return refs


def parse_yaml_template(path: Path) -> list[LabelRef]:
    lines = path.read_text(encoding="utf-8").splitlines()
    refs: list[LabelRef] = []
    for idx, line in enumerate(lines):
        if re.match(r"^\s*labels\s*:\s*$", line):
            refs.extend(parse_labels_block(lines, idx, path))
    return refs


def parse_md_frontmatter_template(path: Path) -> list[LabelRef]:
    lines = path.read_text(encoding="utf-8").splitlines()
    if not lines or lines[0].strip() != "---":
        return []
    end_idx = -1
    for idx in range(1, len(lines)):
        if lines[idx].strip() == "---":
            end_idx = idx
            break
    if end_idx == -1:
        return []
    frontmatter = lines[: end_idx + 1]
    refs: list[LabelRef] = []
    for idx, line in enumerate(frontmatter):
        if re.match(r"^\s*labels\s*:\s*$", line):
            refs.extend(parse_labels_block(frontmatter, idx, path))
    return refs


def read_template_label_refs(template_dir: Path) -> list[LabelRef]:
    refs: list[LabelRef] = []
    for path in sorted(template_dir.iterdir()):
        if not path.is_file():
            continue
        suffix = path.suffix.lower()
        if suffix in {".yml", ".yaml"}:
            refs.extend(parse_yaml_template(path))
        elif suffix == ".md":
            refs.extend(parse_md_frontmatter_template(path))
    return refs


def main() -> int:
    defined = read_defined_labels(LABELS_FILE)
    template_refs = read_template_label_refs(TEMPLATES_DIR)

    undefined = [ref for ref in template_refs if ref.label not in defined]

    if undefined:
        print("Undefined label references found in issue templates:")
        for ref in undefined:
            rel = ref.file.relative_to(ROOT)
            print(f"  - {rel}:{ref.line} -> {ref.label}")
        print("\nAdd these labels to .github/labels.yaml or fix the template references.")
        return 1

    print("OK: all issue template labels are defined in .github/labels.yaml")
    return 0


if __name__ == "__main__":
    sys.exit(main())
