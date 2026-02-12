# Contributions 

## Overview
Given the rapid pace of change in technology in this domain there is a **high-change revision cycle**. This means the further data in this space could evolve rapidly, resulting in items and recommendations listed, becoming stale or out of date. 

For this reason, please ensure to review the CNCF web site for newly released documents and whitepapers.

It should be noted that this checklist is aimed to be concise and limited in size, not as a replacement of broader more in depth whitepapers and documents which are dedicated to a single given application function or technology domain (ie: security).

For future release contribution please find details on the process below: 

---

## How to Contribute 

### 1) Propose a Change (Intake & Scoping)
1. Create a ticket/issue for every change (even small ones).
2. Include:
   - **Summary** of the change
   - **Rationale** (why this is needed)
   - **Evidence/Sources** (links, papers, data)
   - **Affected sections** (headings/page refs)
   - **Impact/Risk** (does it change claims, guidance, diagrams?)
   - **Suggested reviewers** (SME/editor/owner)

**Acceptance criteria:** the TOC confirms the change is in-scope and works with the corresponding TAG to schedule it into the next planned revision cycle.

---

### 2) Draft the Update (Working Branch / Working Copy)
1. Create a working branch tied to the ticket:
   - `feature/<ticket-id>-short-title` (recommended)
2. Keep changes **atomic**:
   - One topic per change set (commit/PR) where possible
3. Preserve traceability:
   - Prefer diffs (PR/MR) or track-changes-style editing
4. Update “neighboring” artifacts as needed:
   - Citations/references
   - Figures/diagrams and captions
   - Glossary/definitions (if terminology changes)

---

### 3) Request Review (Peer + SME Validation)
Open a PR/MR and request reviewers:
- **SME reviewer(s):** correctness of technical content
- **Editor:** structure, clarity, style consistency
- **Maintainer/Owner:** scope alignment and final approval path

**Review checklist**
- Claims are accurate and **verifiable**
- New/changed content includes **sources**
- Terminology is consistent across the document
- No contradictions introduced across sections
- References, links, and numbering are correct
- Diagrams/figures are updated and sourced

---

## Versioning

Choose one of the following schemes (the maintainer will confirm which is authoritative for this whitepaper).

### Option A: Semantic Versioning (Reader-Friendly)
Format: `MAJOR.MINOR.PATCH`
- **MAJOR**: significant restructuring or changed guidance/claims (high revision)
- **MINOR**: meaningful additions (new sections) without changing core conclusions
- **PATCH**: typos, formatting, broken links, small clarifications

Examples:
- `2.0.0` major rewrite / high-change update
- `2.1.0` new sections and updated examples
- `2.1.2` citation fixes, wording tweaks, link corrections


## Roles and Responsibilities
- **Maintainer/Owner:** scope control, merge approvals, release tagging/publishing
- **SME Reviewer(s):** validates technical correctness
- **Editor:** ensures readability, structure, and style consistency
- **Contributor:** drafts changes, responds to review feedback, updates sources

---

## Decision Log (When Required)
For high-impact or contentious changes (e.g., altered recommendations, changed claims):
- Record: **decision**, **rationale**, **date**, **approver(s)**, and **links** to evidence and PR/issue.

---

## Getting Help
If you’re unsure about scope, versioning, or review requirements:
- Open an issue labeled **question** (or equivalent)
- Tag/mention the maintainer and proposed SME reviewer(s)

--- 

## Checklist for PR/MR Authors
- [ ] Ticket/issue exists and is linked
- [ ] Changes are scoped and atomic
- [ ] Sources added/updated for new claims
- [ ] Figures/diagrams updated (and sourced if applicable)
- [ ] Terminology is consistent across sections
- [ ] Changelog entry prepared (if required by the maintainer)
- [ ] Requested SME + editorial review