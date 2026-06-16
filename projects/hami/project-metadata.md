# HAMi — CNCF `.project` resources (chair index)

**Canonical repo:** https://github.com/Project-HAMi/.project

Use this repo for centralized project metadata (landscape sync, maintainer roster, resource paths). Refresh before citing paths in the [incubation DD](./hami-incubation-dd.md).

| File | Role |
| :--- | :--- |
| [`project.yaml`](https://github.com/Project-HAMi/.project/blob/main/project.yaml) | Identity, maturity log, primary repo, website, security/governance/legal/doc paths, landscape category |
| [`maintainers.yaml`](https://github.com/Project-HAMi/.project/blob/main/maintainers.yaml) | `project-maintainers` and `reviewers` teams |
| [`ADOPTERS.md`](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md) | Adopters list (`adopters.path` in `project.yaml`) |

**LFX Insights (past year):** https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01

**Schema / automation:** https://github.com/cncf/automation/tree/main/utilities/dot-project

**Chair mirror (kcontrol, not on this repo):** `kcontrol/cncf_toc/context/projects/hami-canonical-project-info.md` — snapshot for DD cross-checks; update when `.project` changes on `main`.

## TOC review artifacts (`cncf/toc`)

| Artifact | Status | Link |
| :--- | :--- | :--- |
| GTR | Merged 2026-06-01 | [tech-review/2026-04-14.md](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md) · [#2118](https://github.com/cncf/toc/issues/2118) |
| Security SSA | Complete on `main` | [security-assesment/self-assessment.md](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md) |
| TAG joint assessment | Not filed | Schedule per TAG-Security when ready (non-blocking for incubation) |
| OpenSSF Best Practices | Passing (verified 2026-06-02) | https://www.bestpractices.dev/en/projects/9416 |

**Security follow-ups (non-blocking)** — tracked in [hami-incubation-dd.md § Security](./hami-incubation-dd.md#security): SECURITY.md index in community README; org 2FA in community-membership; PST roster in SECURITY.md; SSA date refresh; joint assessment before graduation; SLSA/attestations; optional clowarden for org ACLs.

**DD verification notes**

- [x] Security — five Required criteria evaluated in chair DD (2026-06-02)
- [ ] Website: `project.yaml` vs application [#1775](https://github.com/cncf/toc/issues/1775) (`project-hami.io` vs `project-hami.github.io`)
- [ ] Sub-repos (HAMi-core, volcano-vgpu-device-plugin, HAMi-WebUI) — in application; not all in `project.yaml` `repositories`
- [ ] Sandbox maturity issue link in `project.yaml` (AUTO-DETECTED comment — confirm `cncf/sandbox#132`)
