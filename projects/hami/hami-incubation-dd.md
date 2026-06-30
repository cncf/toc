# HAMi Incubation Due Diligence

**Status:** Complete — ready for public comment PR to [cncf/toc](https://github.com/cncf/toc).

HAMi (Heterogeneous AI Computing Virtualization Middleware) enables Kubernetes clusters to share and isolate GPU and accelerator resources across workloads without application code changes. It is used in production across education, cloud platform, logistics, real estate, automotive, and tech platform verticals, with 70+ organizations on the [public adopters list](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md). The TOC has found HAMi to have satisfied the criteria for CNCF Incubation.

- **Incubation application:** [cncf/toc#1775](https://github.com/cncf/toc/issues/1775)
- **CNCF project metadata (`.project`):** [Project-HAMi/.project](https://github.com/Project-HAMi/.project) — `project.yaml`, `maintainers.yaml`, [ADOPTERS.md](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md)
- **Project:** Sandbox accepted 2024-08-21 → applying **Incubation** · [Project-HAMi/HAMi](https://github.com/Project-HAMi/HAMi) · [project-hami.io](https://project-hami.io/) (`.project` lists [project-hami.github.io/HAMi](https://project-hami.github.io/HAMi/); both resolve to the same documentation site)
- **GTR on `main`:** [tech-review/2026-04-14.md](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md) ([#2113](https://github.com/cncf/toc/pull/2113)) · tracking [#2118](https://github.com/cncf/toc/issues/2118)
- **Security SSA:** [self-assessment.md](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md) (complete on `main`)
- **OpenSSF Best Practices:** [passing — project 9416](https://www.bestpractices.dev/en/projects/9416) (TOC member verified 2026-06-02)
- **LFX Insights (past year):** [HAMi project dashboard](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01) — 2025-06-01 through 2026-06-01, all org repos
- **Adopter verification (interviews):** TOC member interviews with Project-HAMi adopters (minimum three per incubation process). Findings are summarized under [TOC verification of adopters](#toc-verification-of-adopters); [#### Adoption](#adoption) is interview status only.
- **CNCF policies:** [https://www.cncf.io/policies/](https://www.cncf.io/policies/) · [https://github.com/cncf/foundation/tree/main/policies-guidance](https://github.com/cncf/foundation/tree/main/policies-guidance)

## Incubation Evaluation Summary for HAMi

### Criteria Evaluation

Karena Angell, TOC Member and Chair, conducted the due diligence of HAMi for **Incubation**. Adopter interviews were conducted by Karena Angell and Kevin Wang, TOC Member and Vice Chair.

The project has completed the criteria that show its maturity at Incubation. The following are considered note-worthy:

- **Security:** Product Security Team process documented with a **~48h response target** — a concrete SLA that most projects at this stage leave undefined.
- **Governance and Maintainers:** **Five maintainers across four organizations** for a GPU middleware project; multi-vendor GPU scope TOC-verified in direct maintainer conversations at KubeCon EU Amsterdam 2026.
- **Contributors and Community:** [LFX Insights](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01) overall health **Excellent** — **100** quarterly active contributors, **360** of 365 active days, ~**945** PRs/month, ~**5**-day mean PR merge time; bilingual (English and Chinese) community meetings; CONTRIBUTING policy explicitly addresses AI-assistance disclosure.
- **Ecosystem:** **Five** independent [CNCF case studies](https://www.cncf.io/case-studies/?_sft_lf-project=hami) with production narratives across distinct verticals — unusually strong evidence at incubation; [70+ organizations](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md) on the public adopters list. TOC adopter verification complete across three interviews.

The following actions were provided to the project that were considered blocking but since resolved:

- [Update the link in the Governance.md to directly reference the CNCF Code of Conduct](https://github.com/Project-HAMi/community/issues/42)

**Security (TOC member evaluation, 2026-06-02):** From a security standpoint, HAMi meets the criteria for incubation. The project has completed a thorough [Security Self-Assessment on `cncf/toc` `main](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md)`, maintains a discoverable reporting path through SECURITY.md and GitHub Security Advisories, and follows a PR-only merge model with DCO sign-off and multi-maintainer review. The Product Security Team process is documented in the SSA, and the [OpenSSF Best Practices passing badge](https://www.bestpractices.dev/en/projects/9416) was verified at passing level on 2026-06-02. A TAG joint assessment has not yet been filed; that is acceptable at incubation and is something the project should plan for before graduation. Criterion-level evidence appears in [## Security](#security).

**Recommendations (non-blocking):**

- improve multi-repo discoverability of SECURITY.md (e.g. [OpenFGA community security policy](https://github.com/openfga/community/security/policy), [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md), [KServe incubation DD](https://github.com/cncf/toc/blob/main/projects/kserve/kserve-incubation-dd.md))
- document org-wide **2FA** expectations (e.g. [OpenFGA incubation DD](https://github.com/cncf/toc/blob/main/projects/openfga/openfga-incubation-dd.md), [KServe membership requirements](https://github.com/kserve/community/blob/main/membership.md))
- name PST / security responders in SECURITY.md or governance docs (e.g. [OpenFGA GOVERNANCE#core](https://github.com/openfga/.github/blob/main/GOVERNANCE.md#core), [KServe ROLES.md](https://github.com/kserve/community/blob/main/ROLES.md))
- refresh SSA date when processes change
- pursue TAG joint assessment and supply-chain maturity before graduation (e.g. [OpenFGA joint assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/joint-assessment.md))
- add org ACL tooling as sub-repos grow (e.g. [clowarden](https://github.com/cncf/clowarden), [peribolos](https://docs.prow.k8s.io/docs/components/cli-tools/peribolos/), [sheriff](https://github.com/electron/sheriff) — [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md))

**Governance (TOC member evaluation, 2026-06-02):** HAMi meets incubation governance expectations. Public governance documentation, a five-maintainer roster across four organizations, a contributor ladder enforced through OWNERS, and a CNCF-aligned Code of Conduct are in place. **TOC verification (2026-06-02):** Multi-vendor GPU support reflects ecosystem scope (maintainer discussions at multiple events, including KubeCon EU Amsterdam 2026). Vendor-neutrality values are stated in [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) but not yet spelled out as procedural rules; [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) does not document a roadmap change process. Container images identify employer mailboxes in Dockerfile `LABEL maintainer`. A governance review was conducted by the TOC member as part of this DD (see [## Governance and Maintainers](#governance-and-maintainers)), not as a stand-alone [Project Reviews Subproject](https://github.com/cncf/toc/tree/main/toc_subprojects/project-reviews-subproject) review.

**Recommendations (non-blocking):**

- fix `governance.md` link to `CODE-OF-CONDUCT.md` vs actual filename
- add written vendor-neutrality rules to governance (e.g. [OpenFGA incubation DD](https://github.com/cncf/toc/blob/main/projects/openfga/openfga-incubation-dd.md), [KServe GOVERNANCE.md](https://github.com/kserve/community/blob/main/GOVERNANCE.md))
- replace employer-specific Dockerfile `LABEL maintainer` values (`info@dynamia.ai`, `projecthami@dynamia.ai`, `opensource@4paradigm.com` in [docker/Dockerfile](https://github.com/Project-HAMi/HAMi/blob/master/docker/Dockerfile) and sibling Dockerfiles) with project-neutral contact (`The HAMi Authors`, a neutral list email) and OCI annotations (`org.opencontainers.image.authors`, `org.opencontainers.image.source`) per [CNCF vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) — peer pattern: Prometheus, Thanos, Falco, Cilium
- add stable public links or in-repo meeting notes for leadership decisions
- cite public GitHub PRs for reviewer/maintainer promotions — see criterion-level notes in [## Governance and Maintainers](#governance-and-maintainers) and [## Contributors and Community](#contributors-and-community)
- merge/document subproject governance on `main` before graduation
- reduce **organization dependency** (two orgs >51% of contributions per [LFX Insights](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01)) before graduation

**Community (TOC member evaluation, 2026-06-02):** HAMi presents a healthy contributor community for incubation. The project documents a clear contributor ladder, straightforward issue and PR paths, and multiple public channels—including Discord, CNCF Slack, a mailing list, and bilingual community meetings with LFX Zoom for English sessions. **TOC member verified (2026-06-02):** Vendor-neutral [CNCF Slack #hami-dev](https://cloud-native.slack.com/archives/C07T10BU4R2) is present (quiet relative to Discord and other channels). HAMi public community meetings are listed on the [CNCF calendar](https://www.cncf.io/calendar/) (LFX Project Control Center), **bi-weekly on alternating Wednesdays**. CONTRIBUTING guidance is in place across the org. [LFX Insights](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01) over the past year shows strong sustained activity (100 quarterly active contributors, 360 active days of 365), which supports the Required bar. Details and per-criterion notes are in [## Contributors and Community](#contributors-and-community).

**Recommendations (non-blocking):**

- centralize channel index on [Project-HAMi/community](https://github.com/Project-HAMi/community) README (e.g. [OpenFGA community docs](https://openfga.dev/docs/community), [KServe incubation DD](https://github.com/cncf/toc/blob/main/projects/kserve/kserve-incubation-dd.md))
- keep Discord and other active channels as primary entry points; post release notes and meeting reminders to [CNCF Slack #hami-dev](https://cloud-native.slack.com/archives/C07T10BU4R2) regularly (channel is quiet today)
- document WeChat groups with purpose on community repo (e.g. [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md))
- cite public PRs for reviewer/maintainer promotions (e.g. [KServe #3484](https://github.com/kserve/kserve/pull/3484), [#3966](https://github.com/kserve/kserve/pull/3966); [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md) for LFX-based recruitment evidence)

**Ecosystem (TOC member evaluation, 2026-06-15):** HAMi has built strong adoption evidence for incubation. The project maintains a public adopters list (`[.project` ADOPTERS.md]([https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md)), 70+ organizations), **five** independent [CNCF case studies](https://www.cncf.io/case-studies/?_sft_lf-project=hami) with production narratives, and documented CNCF integrations (Kubernetes, Volcano, Koordinator). **TOC verification of adopters** is **complete** (three of three minimum interviews) with **three** strengths and **five** improvement themes from completed interviews and are under [TOC verification of adopters](#toc-verification-of-adopters) in [## Ecosystem](#ecosystem); [#### Adoption](#adoption) is interview status only.

**Recommendations (non-blocking):**

- indicate adoption level (prod/dev/trial) on public list (e.g. [OpenFGA ADOPTERS.md](https://github.com/openfga/community/blob/main/ADOPTERS.md))
- label CNCF vs non-CNCF on integrations docs (e.g. [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md))
- reconcile site adopters page with `.project` list

**Engineering Principles (TOC member evaluation, 2026-06-02):** Engineering Principles evidence meets the criteria for Incubation and is supplemented by the merged [GTR on `cncf/toc` `main](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md)` (2026-04-14), architecture docs, and a documented [release process](https://github.com/Project-HAMi/HAMi/blob/master/docs/release-process.md). **[HAMi Roadmap #923](https://github.com/Project-HAMi/HAMi/issues/923) is closed** (2025-09-28); per-release tracking issue [#1615](https://github.com/Project-HAMi/HAMi/issues/1615) is also closed—TOC recommends a **persistent** public roadmap (project board or in-repo document), not only cross-linking the closed issue. [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) does not document roadmap change process. Regular semver releases from v2.4.0 (2024-09) through v2.9.0 (2026-05) support the Suggested release-history criterion. All seven Engineering Principles checkboxes are `[x]`; GTR satisfies goals/architecture Required items the application left for TOC completion. See [## Engineering Principles](#engineering-principles).

**Recommendations (non-blocking):**

- document roadmap change process in [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) (no roadmap process documented there today; practice lives in issues and meetings per [## Engineering Principles](#engineering-principles))
- establish and maintain a **persistent** public roadmap — see [## Engineering Principles](#engineering-principles) criterion-level recommendations for details and peer examples
- add user-facing release cadence and support-window summary at top of [release-process.md](https://github.com/Project-HAMi/HAMi/blob/master/docs/release-process.md) (e.g. [OpenFGA RELEASES.md](https://github.com/openfga/openfga/blob/main/RELEASES.md), [wasmCloud RELEASE.md](https://github.com/wasmCloud/wasmCloud/blob/main/RELEASE.md), [KServe RELEASE_PROCESS_v2.md](https://github.com/kserve/kserve/blob/master/release/RELEASE_PROCESS_v2.md))
- expand long-chain upgrade/rollback test matrices ([GTR Day 1](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md) note)
- complete release automation steps documented
- align container image maintainer/OCI labels with project-neutral contact — see [## Governance and Maintainers](#governance-and-maintainers) vendor-neutrality recommendations
- publish **performance benchmark data** covering expected overhead and isolation quality across common concurrent workload scenarios (multi-GPU training, single-GPU training, mixed training and inference, varying framework configurations)

### Adoption Evaluation

**Adoption Evaluation** is based on TOC member live adopter interviews. 

**Three** adopter interviews are complete (minimum met); **three** strengths and **five** improvements areas are summarized under [TOC verification of adopters](#toc-verification-of-adopters) in [## Ecosystem](#ecosystem). [#### Adoption](#adoption) below records interview status by industry vertical only.

**What adopters reported (summary):** Adopters consistently cited multi-year production stability, effective vGPU memory isolation for bursty development jobs and underutilized inference workloads, and responsive maintainers. The most-requested improvements are richer vGPU-slice-level observability metrics, published performance benchmark data, and operational guidance for large-fleet `hami-device-plugin` rollouts. Full synthesis is in [TOC verification of adopters](#toc-verification-of-adopters).

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

## Application Process Principles

**Meets all Required criteria; non-blocking recommendations below.**

### Suggested

- **Engage with domain-specific TAG(s) to present the technical architecture of the project.**
  - HAMi presented to [TAG Runtime](https://github.com/cncf/tag-runtime) on 14-May-2024; [session recording](https://youtu.be/SwdEoQYkMsE).

Per [incubation application #1775](https://github.com/cncf/toc/issues/1775): the project engaged TAG Runtime through this presentation (application criterion: engage domain-specific TAG(s) via presentation or GTR). HAMi is adopted in the [CNCF landscape](https://landscape.cncf.io/?item=orchestration-management--scheduling-orchestration--hami); the application also records TAG insight/recommendation in landscape context at that entry.

### Required

- **Complete a [General Technical Review (GTR)](../toc_subprojects/project-reviews-subproject/general-technical-questions.md).**
  - Completed 2026-04-14; artifact on [cncf/toc `main](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md)` (merged via [#2113](https://github.com/cncf/toc/pull/2113), 2026-06-01). Tracking: [#2118](https://github.com/cncf/toc/issues/2118).
- **Complete a [Governance Review](../toc_subprojects/project-reviews-subproject/governance-review-template.md).**
  - **TOC member evaluation (2026-06-02):** Incubation governance criteria are evaluated in [## Governance and Maintainers](#governance-and-maintainers) by the TOC assignee per [#1775](https://github.com/cncf/toc/issues/1775)—not via a stand-alone Project Reviews Subproject review merged to `projects/hami/` on `cncf/toc` `main`. A separate subproject governance review remains optional for graduation planning.
- **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

Project asserts vendor-neutral metadata and resources in [#1775](https://github.com/cncf/toc/issues/1775). **TOC verification (2026-06-02):** Vendor-neutral [CNCF Slack #hami-dev](https://cloud-native.slack.com/archives/C07T10BU4R2) confirmed (low activity relative to Discord and other channels—see [## Contributors and Community](#contributors-and-community)). Multi-vendor GPU/accelerator support reflects **ecosystem scope**, not single-vendor control of project direction, based on TOC member discussions with maintainers at multiple events including **KubeCon EU Amsterdam 2026**. [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) does **not** document a roadmap change process (see [## Engineering Principles](#engineering-principles)); historical planning used [HAMi Roadmap #923](https://github.com/Project-HAMi/HAMi/issues/923) (**closed** 2025-09-28)—TOC recommends a persistent project board or roadmap document. Container images use employer-domain Dockerfile maintainer labels (see [## Governance and Maintainers](#governance-and-maintainers)); **non-blocking** recommendation to adopt project-neutral image metadata before graduation.

- **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
- Sandbox accepted 2024-08-21; moving-levels expectations acknowledged in application 2025-07-11 per [#1775](https://github.com/cncf/toc/issues/1775).

Confirmed against landscape sandbox acceptance date and application checklist.

- **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

This document.

- **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

Installation, user, and reference material on [https://project-hami.io](https://project-hami.io) and linked repositories (per [#1775](https://github.com/cncf/toc/issues/1775)).

Documentation is discoverable from the project site and GitHub org. Engineering Principles criteria are evaluated in [## Engineering Principles](#engineering-principles); the merged GTR satisfies goals and architecture Required items that the application left for TOC completion ([#1775](https://github.com/cncf/toc/issues/1775)). Security documentation includes the merged [Security Self-Assessment](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md) on `cncf/toc` (see [## Security](#security)).

## Governance and Maintainers

**Meets all Required criteria; non-blocking recommendations below.**

**Governance evaluation (TOC member, 2026-06-02):** This section satisfies the incubation governance review requirement through TOC member due diligence per [#1775](https://github.com/cncf/toc/issues/1775), [Project-HAMi/.project](https://github.com/Project-HAMi/.project), and public governance artifacts—not a stand-alone [Project Reviews Subproject](https://github.com/cncf/toc/tree/main/toc_subprojects/project-reviews-subproject) governance review merged to `cncf/toc` `main`.

### Suggested

- **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

Per [#1775](https://github.com/cncf/toc/issues/1775): governance has been iterated over the past year; history is visible in [Project-HAMi/community commits](https://github.com/Project-HAMi/community/commits/main/).

The project has been iterating on governance since sandbox acceptance (2024-08-21), and commit history on the community repository shows ongoing updates to governance documents—a healthy sign that practice is evolving with experience.

- **Clear and discoverable project governance documentation.**

[governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) (values, membership pointer, meetings, CoC, charter changes); contributor ladder in [community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md); roster in [MAINTAINERS.md](https://github.com/Project-HAMi/HAMi/blob/master/MAINTAINERS.md).

Governance documentation is public and reachable from the org README and `.project` metadata paths.

**Recommendations:**

1. resolve `governance.md` link to `CODE_OF_CONDUCT.md` vs actual file `[CODE-OF-CONDUCT.md](https://github.com/Project-HAMi/community/blob/main/CODE-OF-CONDUCT.md)`

- **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

[#1775](https://github.com/cncf/toc/issues/1775): [Meetings](https://github.com/Project-HAMi/community/blob/main/governance.md#meetings) (public developer meeting agenda); closed meetings for security/CoC; decisions tracked in [closed community issues](https://github.com/Project-HAMi/community/issues?q=is%3Aissue+state%3Aclosed).

The weekly public meeting agenda is documented, and the application describes how maintainer decisions are recorded.

**Recommendations (non-blocking):**

1. add stable public links (or meeting notes in-repo) for leadership decisions, not only private docs

- **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

[governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) values include “Community over Product or Company”; application [#1775](https://github.com/cncf/toc/issues/1775) left this checkbox unchecked.

**TOC verification (2026-06-02):** Multi-vendor GPU and accelerator support reflects **ecosystem scope** rather than single-vendor control of project direction, based on TOC member discussions with maintainers at multiple events including **KubeCon EU Amsterdam 2026**.

**Partial (non-blocking for incubation):** The project's stated values align with vendor-neutrality principles, but there is not yet a dedicated section in [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) explaining how neutral direction is maintained in practice (for example, peer incubation DDs document written vendor-neutrality rules in [OpenFGA incubation DD](https://github.com/cncf/toc/blob/main/projects/openfga/openfga-incubation-dd.md) and [KServe charter](https://github.com/kserve/community/blob/main/GOVERNANCE.md)). [LFX Insights](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01) flags **organization dependency**—two organizations account for more than 51% of contributions—which is worth monitoring alongside maintainer employers listed in [MAINTAINERS.md](https://github.com/Project-HAMi/HAMi/blob/master/MAINTAINERS.md). Published container images use employer-specific `LABEL maintainer` values (`info@dynamia.ai`, `projecthami@dynamia.ai`, `opensource@4paradigm.com`) in [docker/Dockerfile](https://github.com/Project-HAMi/HAMi/blob/master/docker/Dockerfile) and sibling Dockerfiles; graduated CNCF peers favor project-collective or neutral contacts (Prometheus `The Prometheus Authors`, Thanos `The Thanos Authors`, Falco `cncf-falco-dev@lists.cncf.io`, Cilium `maintainer@cilium.io`).

**Recommendations (non-blocking):**

1. add written vendor-neutrality rules to [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md)
2. replace Dockerfile employer maintainer emails with project-neutral labels (`The HAMi Authors`, neutral list email) and OCI annotations (`org.opencontainers.image.authors`, `org.opencontainers.image.source`) per [CNCF vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/)
3. document neutral comms, hosting, and architectural decision process before graduation

- **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

[#1775](https://github.com/cncf/toc/issues/1775): maintainer membership via maintainer meeting ([notes](https://docs.google.com/document/d/1mRUhNm3spCdRLcxFBuCYNfwQJqSO9RwXi82IHF6M4Xk/edit)); weekly meeting for reviewers/approvers/governance/goals ([agenda](https://docs.google.com/document/d/1YC6hco03_oXbF9IOUPJ29VWEddmITIKIfSmBX8JtGBw/edit)); charter changes require 2/3 maintainer vote per [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md#modifying-this-charter).

Decision processes are described and practiced; meeting cadence and membership paths are discoverable. Not blocking for incubation.

**Recommendations (non-blocking):**

1. strengthen formal written rules for CNCF-facing requests (e.g. [OpenFGA community #504](https://github.com/openfga/community/issues/504), [KServe #53](https://github.com/kserve/community/issues/53))

- **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

[#1775](https://github.com/cncf/toc/issues/1775): no separate sub-teams beyond maintainers; roles (member → reviewer → approver → maintainer) in [community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) with [OWNERS](https://github.com/Project-HAMi/HAMi/blob/master/OWNERS) enforcement.

For the project's current structure, documenting roles through the contributor ladder and OWNERS files is adequate. A dedicated security response committee is not required at incubation; maintainers and the Product Security Team handle security per [SSA § Security issue resolution](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md#security-issue-resolution) and [## Security](#security) below.

- **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

[community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) defines requirements per role; **Inactivity** and **Involuntary Removal or Demotion** sections cover emeritus/demotion paths.

The maintainer lifecycle is documented, including inactivity and involuntary removal paths. The project's `community-membership.md` is marked work in progress.

**Recommendations (non-blocking):**

1. tighten emeritus naming and add public examples in a follow-up

- **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

[#1775](https://github.com/cncf/toc/issues/1775): maintainer group notes ([Google Doc](https://docs.google.com/document/d/1mRUhNm3spCdRLcxFBuCYNfwQJqSO9RwXi82IHF6M4Xk/edit)); roster of five maintainers across four affiliations in [MAINTAINERS.md](https://github.com/Project-HAMi/HAMi/blob/master/MAINTAINERS.md) (aligned with [maintainers.yaml](https://github.com/Project-HAMi/.project/blob/main/maintainers.yaml)).

Maintainer group notes and a five-maintainer roster across four affiliations demonstrate lifecycle use in practice.

**Recommendations:**

1. cite public GitHub issues/PRs for promotions (e.g. community membership requests) in addition to private notes for transparency

- **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

Sub-repos per [#1775](https://github.com/cncf/toc/issues/1775): [HAMi-core](https://github.com/Project-HAMi/HAMi-core), [volcano-vgpu-device-plugin](https://github.com/Project-HAMi/volcano-vgpu-device-plugin), [HAMi-WebUI](https://github.com/Project-HAMi/HAMi-WebUI). Application referenced extended [community-membership](https://github.com/Project-HAMi/community/blob/add_sub_projects/community-membership.md#requirements-1) on a branch.

**Partial:** subprojects are listed and named in application/GTR; `community-membership.md` on `main` references subprojects in role tables but lacks a full subproject charter/add-remove process on `main`.

**Recommendations:**

1. merge/document subproject governance on `main` before graduation

### Required

- **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

[MAINTAINERS.md](https://github.com/Project-HAMi/HAMi/blob/master/MAINTAINERS.md) — five maintainers with email and employer; roster mirrored in [Project-HAMi/.project `maintainers.yaml](https://github.com/Project-HAMi/.project/blob/main/maintainers.yaml)`.

The maintainer roster is complete for incubation.

**Recommendations (non-blocking):**

1. add clearer domain-of-responsibility per technical area as the project grows

- **A number of active maintainers which is appropriate to the size and scope of the project.**

Five named maintainers from **four organizations** (dynamia.ai ×2, NVIDIA, 4Paradigm, independent) for a multi-vendor GPU scheduling project with several sub-repos.

Five maintainers from four organizations is appropriate for a multi-vendor GPU scheduling project at incubation. [LFX Insights](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01) (past year, all repos) reports **100** quarterly active contributors, **27%** quarterly contributor retention, and contributor dependency of **7** people for 51%+ of commits.

**Recommendations:**

1. grow reviewer/approver bench and reduce **organization dependency** (two organizations account for 51%+ of contributions per LFX) to align with vendor-neutrality goals before graduation

- **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

[OWNERS](https://github.com/Project-HAMi/HAMi/blob/master/OWNERS) and org [OWNERS](https://github.com/Project-HAMi/community/blob/main/OWNERS) align with reviewer/approver/maintainer ladder in [community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md).

The OWNERS model on the primary repository is consistent with documented reviewer, approver, and maintainer roles.

- **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

[CODE-OF-CONDUCT.md](https://github.com/Project-HAMi/community/blob/main/CODE-OF-CONDUCT.md) (based on [CNCF Code of Conduct](https://www.cncf.io/conduct/)); violations escalated per [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md#code-of-conduct) to CNCF CoC Committee.

The project has adopted a CNCF-aligned Code of Conduct and describes an escalation path to the CNCF CoC Committee.

- **CNCF Code of Conduct is cross-linked from other governance documents.**

[community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) links [CODE-OF-CONDUCT.md](https://github.com/Project-HAMi/community/blob/main/CODE-OF-CONDUCT.md). [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md#code-of-conduct) references `./CODE_OF_CONDUCT.md` (filename mismatch).

CoC is linked from community-membership.md. The `governance.md` reference to `./CODE_OF_CONDUCT.md` contains a filename mismatch (actual file: `CODE-OF-CONDUCT.md`).

**Recommendations (non-blocking):**

1. fix `governance.md` CoC link to `CODE-OF-CONDUCT.md` and cross-link from `governance.md` opening section

- **All subprojects, if any, are listed.**

[#1775](https://github.com/cncf/toc/issues/1775) and GTR: [HAMi-core](https://github.com/Project-HAMi/HAMi-core), [volcano-vgpu-device-plugin](https://github.com/Project-HAMi/volcano-vgpu-device-plugin), [HAMi-WebUI](https://github.com/Project-HAMi/HAMi-WebUI) (primary scheduling repo: [HAMi](https://github.com/Project-HAMi/HAMi)).

Subprojects are listed in the application and GTR.

**Recommendations (non-blocking):**

1. add sub-repos to `.project` when metadata model allows

## Contributors and Community

**Meets all Required criteria; non-blocking recommendations below.**

**Community metrics ([LFX Insights](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01), past year 2025-06-01 → 2026-06-01, all repositories):** Overall health is **Excellent**, with **100** quarterly active contributors, **27%** quarterly contributor retention, and activity on **360** of 365 days. The project averages roughly **945** new PRs per month, resolves issues in about **18** days on mean, and merges PRs in about **5** days. GitHub shows **4,284** stars and **981** forks. These figures provide context for the criterion-by-criterion evaluation below.

### Suggested

- **Contributor ladder with multiple roles for contributors.**

[community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) — member → reviewer → approver → maintainer (OWNERS / MAINTAINERS enforcement).

The contributor ladder is documented and enforced through OWNERS and MAINTAINERS files. LFX shows a large active contributor pool (100 quarterly active contributors) relative to five named maintainers—a healthy sign of breadth relative to the maintainer roster.

**Recommendations (non-blocking):**

1. cross-link [community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) and [contributing.md](https://github.com/Project-HAMi/community/blob/main/contributing.md) from [Project-HAMi/community](https://github.com/Project-HAMi/community) README (e.g. [OpenFGA community hub](https://openfga.dev/docs/community))

### Required

- **Clearly defined and discoverable process to submit issues or changes.**

[GitHub issue templates](https://github.com/Project-HAMi/HAMi/issues/new/choose); contribution workflow in [CONTRIBUTING.md](https://github.com/Project-HAMi/HAMi/blob/master/CONTRIBUTING.md) (primary repo, including AI-assistance disclosure and `make verify`) and [community/contributing.md](https://github.com/Project-HAMi/community/blob/main/contributing.md). Good-first and help-wanted labels documented in CONTRIBUTING.md.

Issue templates and CONTRIBUTING guidance are easy to find from the primary repository README and the community org. The workflow—including AI-assistance disclosure and local verification steps—is described at a depth appropriate for incubation.

**Recommendations (non-blocking):**

1. add a single “Questions and issues” index on [Project-HAMi/community](https://github.com/Project-HAMi/community) README pointing to CONTRIBUTING and issue templates (e.g. [OpenFGA community hub](https://openfga.dev/docs/community), [KServe community README](https://github.com/kserve/community?tab=readme-ov-file#questions-and-issues))

- **Project must have, and document, at least one public communications channel for users and/or contributors.**

Public channels documented in [HAMi README — Community](https://github.com/Project-HAMi/HAMi/blob/master/README.md#community): [Discord](https://discord.gg/Amhy7XmbNq) (recommended), [CNCF Slack #hami-dev](https://cloud-native.slack.com/archives/C07T10BU4R2), [hami-project Google Group](https://groups.google.com/forum/#!forum/hami-project), [project-hami.io](https://project-hami.io), and GitHub issues.

HAMi offers several well-documented public entry points. **TOC verification (2026-06-02):** Vendor-neutral [CNCF Slack #hami-dev](https://cloud-native.slack.com/archives/C07T10BU4R2) is present and appropriate for CNCF users; the channel is **quiet** relative to Discord and other primary channels. That meets incubation expectations for a documented public channel.

**Recommendations (non-blocking):**

1. continue treating Discord and other active channels as primary community entry points (as README recommends)
2. post release notes, meeting reminders, and periodic updates to **#hami-dev** so the CNCF Slack channel stays current for CNCF users

- **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

Per [#1775](https://github.com/cncf/toc/issues/1775) and public docs:


| Channel                                                                           | Scope                                                                                                                                     | Public               |
| --------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | -------------------- |
| [Discord](https://discord.gg/Amhy7XmbNq)                                          | Users and contributors (recommended)                                                                                                      | yes                  |
| [CNCF Slack #hami-dev](https://cloud-native.slack.com/archives/C07T10BU4R2)       | Users and contributors (vendor-neutral; low activity relative to Discord)                                                                 | yes                  |
| [hami-project mailing list](https://groups.google.com/forum/#!forum/hami-project) | Announcements and discussion                                                                                                              | yes                  |
| [GitHub Issues](https://github.com/Project-HAMi/HAMi/issues)                      | Bug reports, features (org sub-repos: HAMi-core, volcano-vgpu-device-plugin, HAMi-WebUI)                                                  | yes                  |
| [project-hami.io](https://project-hami.io) / [blog](https://project-hami.io/blog) | Docs, releases, case studies                                                                                                              | yes                  |
| WeChat discussion groups (2)                                                      | Chinese-speaking community (#1775: ~500 + ~151 members)                                                                                   | no — regional chat   |
| Maintainer closed meetings                                                        | Security reports, CoC violations per [governance.md#meetings](https://github.com/Project-HAMi/community/blob/main/governance.md#meetings) | no — maintainer-only |


Public channels are documented; non-public channels have stated purpose in governance or application.

**Recommendations (non-blocking):**

1. publish WeChat and maintainer-meeting notes on [Project-HAMi/community](https://github.com/Project-HAMi/community) README with purpose text (e.g. [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md)); keep Discord and other active channels as primary entry points; regularly update **#hami-dev** with announcements

- **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

[README — Community](https://github.com/Project-HAMi/HAMi/blob/master/README.md#community): [meeting agenda and notes](https://docs.google.com/document/d/1YC6hco03_oXbF9IOUPJ29VWEddmITIKIfSmBX8JtGBw/edit#heading=h.g61sgp7w0d0c); Chinese community meeting Friday 16:00 UTC+8 weekly ([Tencent meeting link](https://meeting.tencent.com/dm/Ntiwq1BICD1P)); English community meeting Wednesday 16:00 UTC+8 biweekly ([LFX Zoom](https://zoom-lfx.platform.linuxfoundation.org/meeting/95994137931?password=55b961b5-3e8e-4040-8657-0f2d26511f1d)). [governance.md#meetings](https://github.com/Project-HAMi/community/blob/main/governance.md#meetings) references the public developer meeting.

**TOC verification (2026-06-02):** HAMi is listed on the [CNCF public community calendar](https://www.cncf.io/calendar/) (LFX Project Control Center). Public project community meetings appear **bi-weekly on alternating Wednesdays** on that calendar.

Schedules and links are public; English sessions use LFX (e.g. [OpenFGA community meetings](https://openfga.dev/docs/community), [KServe incubation DD](https://github.com/cncf/toc/blob/main/projects/kserve/kserve-incubation-dd.md)). CNCF/LFX calendar integration satisfies this Required criterion.

**Recommendations (non-blocking):**

1. cross-link [CNCF calendar](https://www.cncf.io/calendar/) from [Project-HAMi/community](https://github.com/Project-HAMi/community) README
2. add a meeting archive on the project site (e.g. [wasmCloud community archive](https://wasmcloud.com/community))

- **Documentation of how to contribute, with increasing detail as the project matures.**

[CONTRIBUTING.md](https://github.com/Project-HAMi/HAMi/blob/master/CONTRIBUTING.md): CoC, community expectations, AI-assistance disclosure policy, fork/PR workflow, `make verify`, code review expectations, good-first-issue pointers. [community/contributing.md](https://github.com/Project-HAMi/community/blob/main/contributing.md) mirrors org-level guidance; [community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) defines role progression.

CONTRIBUTING documentation covers the topics incubation expects—CoC, workflow, review expectations, and pointers for newcomers—and the project has gone further than many peers by documenting AI-assistance disclosure policy. That level of detail signals maturity and sets clear expectations for contributors.

- **Demonstrate contributor activity and recruitment.**

[LFX Insights](https://insights.linuxfoundation.org/project/hami?timeRange=past365days&start=2025-06-01&end=2026-06-01) (past year, all repos): overall health **Excellent**; **100** quarterly active contributors; **360** active days; high PR volume; mean PR merge lead time **~5** days; **27%** contributor retention; contributor dependency spread across **7** people for 51%+ of commits.

Meets incubation bar for contributor activity (peer examples: [OpenFGA incubation DD](https://github.com/cncf/toc/blob/main/projects/openfga/openfga-incubation-dd.md) cites DevStats; [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md) cites LFX dashboard).

**Recommendations:**

1. document recruitment/onboarding outcomes in public GitHub issues or PRs (not only private meeting notes)—e.g. peer incubation DDs cite concrete promotions ([KServe #3484](https://github.com/kserve/kserve/pull/3484), [#3966](https://github.com/kserve/kserve/pull/3966); [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md) cites LFX recruitment metrics)
2. broaden organizational participation beyond the two dominant employers flagged in LFX organization dependency

## Engineering Principles

**Meets all Required criteria; non-blocking recommendations below.**

**General Technical Review:** Completed 2026-04-14; artifact on [cncf/toc `main](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md)` (merged [#2113](https://github.com/cncf/toc/pull/2113)). GTR version baseline: v2.8.0; current release at DD review: [v2.9.0](https://github.com/Project-HAMi/HAMi/releases/tag/v2.9.0) (2026-05-19).

The TOC member relies on the merged GTR for differentiation, architecture, and release-process depth; criterion-level notes below confirm incubation satisfaction and record non-blocking recommendations.

### Suggested

- **Roadmap change process is documented.**

Per [#1775](https://github.com/cncf/toc/issues/1775) and [GTR § Scope](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md): scope and roadmap are maintained through GitHub issues and maintainer review; per-version tracking (e.g. [#1615](https://github.com/Project-HAMi/HAMi/issues/1615), closed); after each release the weekly community meeting discusses upcoming goals and assigns owners on tracking issues; users reply on issues to request features. Historical public roadmap hub: [HAMi Roadmap #923](https://github.com/Project-HAMi/HAMi/issues/923) (**closed** 2025-09-28). **[governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) does not document a roadmap change process**—only meetings and charter mechanics.

Roadmap planning is practiced through community meetings and release-scoped issues per GTR, but there is **no open persistent roadmap hub** at TOC review (2026-06-02). The Suggested criterion is satisfied in practice; governance doc and durable roadmap artifact gaps are non-blocking for incubation.

**Recommendations (non-blocking):**

1. document how roadmap changes are proposed and approved in [governance.md](https://github.com/Project-HAMi/community/blob/main/governance.md) (e.g. [OpenFGA RFC process](https://github.com/openfga/rfcs), [wasmCloud #3343](https://github.com/wasmCloud/wasmCloud/issues/3343), [KServe incubation DD](https://github.com/cncf/toc/blob/main/projects/kserve/kserve-incubation-dd.md))
2. establish and maintain a **persistent** public roadmap—GitHub Project board, in-repo `ROADMAP.md`, or equivalent—and link it from [Project-HAMi/community](https://github.com/Project-HAMi/community) README, [HAMi README](https://github.com/Project-HAMi/HAMi/blob/master/README.md), and [project-hami.io](https://project-hami.io/); do not rely on closed [#923](https://github.com/Project-HAMi/HAMi/issues/923) as the active planning surface; peer examples: [OpenFGA org project board](https://github.com/orgs/openfga/projects/1), [wasmCloud ROADMAP.md](https://github.com/wasmCloud/wasmCloud/blob/main/ROADMAP.md), [Cilium community roadmap](https://github.com/cilium/cilium/blob/main/Documentation/community/roadmap.rst)

- **History of regular, quality releases.**

[GitHub Releases](https://github.com/Project-HAMi/HAMi/releases): semver tags from [v2.4.0](https://github.com/Project-HAMi/HAMi/releases/tag/v2.4.0) (2024-09-29, post-sandbox) through [v2.9.0](https://github.com/Project-HAMi/HAMi/releases/tag/v2.9.0) (2026-05-19). Recent cadence includes v2.8.x patch line (Apr–May 2026) and v2.7.0 (2025-09). [Release process](https://github.com/Project-HAMi/HAMi/blob/master/docs/release-process.md) documents MINOR releases roughly every 1–3 months and PATCH releases for critical fixes or roughly monthly. Tag-driven workflows publish container images and Helm charts with each release (`[.project` Docker Hub `projecthami/hami](https://github.com/Project-HAMi/.project)`).

Release history is regular and appropriate for a fast-moving incubation-stage project; chart and container images published with tags per release workflow.

**Recommendations (non-blocking):**

1. publish high-level release cadence and support expectations for operators (not only maintainer runbook) at top of [release-process.md](https://github.com/Project-HAMi/HAMi/blob/master/docs/release-process.md) (e.g. [OpenFGA RELEASES.md](https://github.com/openfga/openfga/blob/main/RELEASES.md), [wasmCloud RELEASE.md](https://github.com/wasmCloud/wasmCloud/blob/main/RELEASE.md), [KServe RELEASE_PROCESS_v2](https://github.com/kserve/kserve/blob/master/release/RELEASE_PROCESS_v2.md))

### Required

- **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. *This requirement may also be satisfied by completing a General Technical Review.***
  - *If applicable* a general Technical Review was completed/updated on **2026-04-14**, and can be discovered at [projects/hami/tech-review/2026-04-14.md](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md).

[GTR § Scope](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md): HAMi is Kubernetes middleware for **heterogeneous AI device sharing, isolation, and scheduling** across GPUs/NPUs and other accelerators—without application code changes. Primary use case: share and isolate accelerator memory/compute for multiple workloads on Kubernetes. **Non-goals:** not a model serving framework, training framework, or cluster autoscaler ([GTR](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md)). Landscape positioning: complements CNCF batch/scheduling stacks ([Volcano](https://volcano.sh/), [Koordinator](https://koordinator.sh/)) documented in GTR and [## Ecosystem](#ecosystem). [CNCF landscape — HAMi](https://landscape.cncf.io/?item=orchestration-management--scheduling-orchestration--hami).

HAMi addresses heterogeneous GPU/accelerator utilization and isolation in Kubernetes—a distinct niche from generic schedulers or ML frameworks. Differentiation and goals are documented for incubation via the merged GTR and project site. The GTR satisfies this Required criterion.

- **Document what the project does, and why it does it - including viable cloud native use cases.**

[GTR § Scope and Usability](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md): target personas include platform engineers, MLOps teams, individual workload owners, and IT operations; interaction via Helm, Pod specs/annotations, and Kubernetes-native APIs. Use cases: GPU memory/core sharing, topology-aware and policy-driven scheduling, heterogeneous vendor devices, dynamic MIG (NVIDIA). [project-hami.io](https://project-hami.io/) · [Installation](https://project-hami.io/docs/installation/online-installation) · [README quick start](https://github.com/Project-HAMi/HAMi/blob/master/README.md).

Cloud native use cases and operator interaction model are documented in GTR and user-facing docs; production integration paths (mutating webhook, scheduler extender, device plugins) are described for incubation.

- **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

Per [#1775](https://github.com/cncf/toc/issues/1775): historical roadmap hub [HAMi Roadmap #923](https://github.com/Project-HAMi/HAMi/issues/923) (**closed** 2025-09-28; formerly rolling version sections); per-release tracking issues such as [#1615](https://github.com/Project-HAMi/HAMi/issues/1615) (**closed**; v2.9 line) ([GTR § Scope](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md)). Community users historically added requests via issue comments on roadmap threads.

**TOC verification (2026-06-02):** No **open** persistent public roadmap or project board at review time; forward planning continues through releases, community meetings, and ad-hoc issues per GTR. Historical GitHub-issue roadmap practice meets the incubation bar in substance; maintainers should re-establish a durable planning artifact before graduation.

**Recommendations (non-blocking):**

1. establish a **persistent** public roadmap — see "Roadmap change process is documented" Suggested criterion above for details and peer examples

- **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. *This requirement may also be satisfied by completing a General Technical Review.***
  - *If applicable* a general Technical Review was completed/updated on **2026-04-14**, and can be discovered at [projects/hami/tech-review/2026-04-14.md](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md).

[Design overview](https://github.com/Project-HAMi/HAMi/blob/master/docs/develop/design.md) · [GTR § Design](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md): core components—**MutatingWebhook**, **Scheduler Extender**, **Device Plugins**, in-container control libraries ([HAMi-core](https://github.com/Project-HAMi/HAMi-core) for NVIDIA); design principles include Kubernetes-native control paths, no application code changes, pluggable heterogeneous devices, policy-driven scheduling, hard isolation on shared devices, and separation of responsibilities. Uses standard K8s APIs (Pods, Nodes, annotations, device plugin gRPC)—no required CRD for core workflows. [GTR § Best practices](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md): CI gates (lint, unit tests, chart packaging, scheduling/device paths), realistic hardware/driver matrices, and documentation alongside features.

Architecture documentation satisfies incubation via GTR and in-repo design doc; viable cloud native deployment (Helm, scheduler extender, device plugin) is explicit.

- **Document the project's release process.**

[docs/release-process.md](https://github.com/Project-HAMi/HAMi/blob/master/docs/release-process.md): semver (MAJOR/MINOR/PATCH timelines), release branches, cherry-pick policy, tag-driven GitHub Actions (images, Helm chart to `gh-pages`, release notes), and maintainer post-release checklist. [GTR § Release process](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md) · [GTR Day 1 — Rollout, upgrade, rollback](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md): Helm install/upgrade/uninstall; `helm rollback` for in-place rollback; running workloads generally stable since v2.5; scheduler HA via leader election since v2.8; deprecations communicated in docs and community meetings with typical one-release transition window; e2e required before merge; long-chain upgrade→downgrade→upgrade matrices are not yet comprehensive.

Release process is documented for incubation; automation covers images and charts with manual verification steps.

**Recommendations (non-blocking):**

1. add operator-facing cadence/support summary at document top (e.g. [OpenFGA RELEASES.md](https://github.com/openfga/openfga/blob/main/RELEASES.md), [wasmCloud RELEASE.md](https://github.com/wasmCloud/wasmCloud/blob/main/RELEASE.md), [Flatcar RELEASE.md](https://github.com/flatcar/Flatcar/blob/main/RELEASE.md))
2. complete release automation steps marked “planing” [sic] in [release-process.md](https://github.com/Project-HAMi/HAMi/blob/master/docs/release-process.md) (automated e2e, website build, version compatibility tests, changelog upload)
3. expand upgrade/rollback test documentation per [GTR Day 1](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md) (long-chain matrices not yet comprehensive)
4. align container image `LABEL maintainer` / OCI metadata with project-neutral contact — see [## Governance and Maintainers](#governance-and-maintainers) vendor-neutrality recommendations for details
5. document **large-fleet rollout/rollback operational guidance** including batch-by-batch strategies, validation gates between batches, and rollback procedures

## Security

**Meets all Required criteria; non-blocking recommendations below.**

Note: this section may be augmented by a joint assessment performed by TAG Security and Compliance. **Security Self-Assessment** on [cncf/toc `main](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md)` (complete); TAG joint assessment not yet filed — non-blocking for incubation. GTR references the same SSA ([tech-review § Security](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md)).

### Suggested

N/A

### Required

- **Clearly defined and discoverable process to report security issues.**

[SECURITY.md](https://github.com/Project-HAMi/HAMi/blob/master/SECURITY.md) and [GitHub security policy](https://github.com/Project-HAMi/HAMi/security/policy) define reporting; private reports via [GitHub Security Advisories](https://github.com/Project-HAMi/HAMi/security/advisories/new). SSA [Security links](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md#security-links) table indexes policy, SECURITY.md, and maintainer contacts.

Reporting paths are clear and discoverable from the primary repository and the SSA security-links table.

**Recommendations (non-blocking):**

1. link [SECURITY.md](https://github.com/Project-HAMi/HAMi/blob/master/SECURITY.md) from [Project-HAMi/community](https://github.com/Project-HAMi/community) README and contributing docs—and from each major org repo if applicable (e.g. [OpenFGA community security policy](https://github.com/openfga/community/security/policy), [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md), [KServe incubation DD](https://github.com/cncf/toc/blob/main/projects/kserve/kserve-incubation-dd.md))

- **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

SSA [Secure development practices](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md#secure-development-practices): all code changes via PR (no direct commits), DCO sign-off, multi-maintainer review, golangci-lint and CI gates. [OWNERS](https://github.com/Project-HAMi/HAMi/blob/master/OWNERS) and [community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) define reviewer/approver/maintainer paths.

Code-base protection meets the incubation bar via PR-only workflow and role-based ownership.

**Recommendations (non-blocking):**

1. document **two-factor authentication** for `Project-HAMi` org members in [community-membership.md](https://github.com/Project-HAMi/community/blob/main/community-membership.md) (e.g. [OpenFGA incubation DD](https://github.com/cncf/toc/blob/main/projects/openfga/openfga-incubation-dd.md), [KServe membership requirements](https://github.com/kserve/community/blob/main/membership.md), [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md))
2. add org ACL tooling as sub-repos grow (e.g. [clowarden](https://github.com/cncf/clowarden), [peribolos](https://docs.prow.k8s.io/docs/components/cli-tools/peribolos/), [sheriff](https://github.com/electron/sheriff) — [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md))

- **Document assignment of security response roles and how reports are handled.**

SSA § [Security issue resolution](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md#security-issue-resolution): Product Security Team (PST), responsible disclosure, ~48h response target, triage and advisory publication via GitHub Security Advisories. [MAINTAINERS.md](https://github.com/Project-HAMi/HAMi/blob/master/MAINTAINERS.md) lists contacts.

Roles and handling are documented in SSA and SECURITY.md; no separate security sub-team (aligned with Governance eval — acceptable at incubation).

**Recommendations:**

1. name PST members or link a short roster in SECURITY.md for external reporters (e.g. [OpenFGA GOVERNANCE#core](https://github.com/openfga/.github/blob/main/GOVERNANCE.md#core), [KServe ROLES.md](https://github.com/kserve/community/blob/main/ROLES.md))

- **Document Security Self-Assessment.**

[HAMi Security Self-Assessment](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md) on `cncf/toc` `main` — metadata **Assessment Stage: Complete**; covers threat-relevant components (webhook, scheduler extender, device plugin, in-container enforcement), secure development pipeline, and issue resolution.

SSA satisfies incubation documentation.

**Recommendations (non-blocking):**

1. refresh “Last updated” when process changes
2. pursue TAG **joint assessment** before graduation (e.g. [OpenFGA joint assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/joint-assessment.md), [Flatcar joint assessment on `cncf/toc](https://github.com/cncf/toc/blob/main/projects/flatcar/security-assessment/joint-assessment.md)`)
3. continue supply-chain attestations/SLSA adoption noted in SSA (SBOM at [community/sbom.md](https://github.com/Project-HAMi/community/blob/main/sbom.md); e.g. [OpenFGA incubation DD](https://github.com/cncf/toc/blob/main/projects/openfga/openfga-incubation-dd.md))

- **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

[HAMi — OpenSSF Best Practices (passing)](https://www.bestpractices.dev/en/projects/9416) — cited in SSA appendix; TOC member verified passing level 2026-06-02.

Badge at passing level; verified 2026-06-02 (peer incubation DDs used the same verification approach, e.g. [OpenFGA](https://www.bestpractices.dev/en/projects/6374), [wasmCloud](https://www.bestpractices.dev/en/projects/6363), [KServe](https://www.bestpractices.dev/en/projects/6643)).

**Recommendations (non-blocking):**

1. maintain badge through incubation review; re-verify at [bestpractices.dev](https://www.bestpractices.dev/en/projects/9416) before graduation public comment

## Ecosystem

**Meets all Required criteria; non-blocking recommendations below.**

For **Incubation**, adopters may be in dev/test or production use; **graduation** expects production use. HAMi is listed in the [CNCF landscape](https://landscape.cncf.io/?item=orchestration-management--scheduling-orchestration--hami).

**Public adopters:** [project-hami.io/adopters](https://project-hami.io/adopters), [HAMi#4](https://github.com/Project-HAMi/HAMi/issues/4). **Canonical list:** [Project-HAMi/.project `ADOPTERS.md](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md)` (70+ organizations; sourced from [website `adopters.json](https://github.com/Project-HAMi/website/blob/master/src/data/adopters.json)`).

### Suggested

N/A

### Required

- **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

[Project-HAMi/.project `ADOPTERS.md](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md)` — 70+ organizations across cloud, telecom, finance, automotive, education, and logistics (e.g. Dynamia, DaoCloud, 4Paradigm, Baidu AI Cloud, Huawei, LinkedIn, SAP, SF Technology, NIO, Ke Holdings, Prep EDU, China Mobile, Ping An Bank, ZStack). Per-adopter **adoption level** (prod/dev/trial) is not indicated on the public list today.

Public list exceeds incubation expectations for breadth; canonical `.project` path aligns with CNCF metadata conventions.

**Recommendations (non-blocking):**

1. add per-adopter **adoption level** (prod/dev/trial) on the public list (e.g. [OpenFGA ADOPTERS.md](https://github.com/openfga/community/blob/main/ADOPTERS.md))
2. enrich with case-study links (e.g. [OpenFGA ADOPTERS.md](https://github.com/openfga/community/blob/main/ADOPTERS.md), [KServe adopters + presentations](https://github.com/cncf/toc/blob/main/projects/kserve/kserve-incubation-dd.md), [community#47](https://github.com/kserve/community/issues/47))
3. reconcile [project-hami.io/adopters](https://project-hami.io/adopters) with `.project` list (site page unavailable as of TOC review 2026-06-02)

- **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

The project meets incubation bar for ≥3 independent adopters based on five public CNCF case studies and a diverse public adopter list: [CNCF end-user case studies](https://www.cncf.io/case-studies/?_sft_lf-project=hami) and [ADOPTERS.md](https://github.com/Project-HAMi/.project/blob/main/ADOPTERS.md).

**Adoption Evaluation** and TOC verification below depend on TOC member interviews with adopters—not on application assertions or case-study text alone.

- **TOC verification of adopters.**

**Complete (2026-06-15):** The TOC members completed **three** adopter interviews (minimum **three** for incubation). Interview notes and recordings are kept privately for internal TOC use and verification. [#### Adoption](#adoption) lists interview status by industry vertical only (anonymous).

**Findings (three interviews):**

**Strengths adopters cited:**

1. **Production maturity** — Multi-year **production** use (approximately two to four years) for on-premises AI/LLM platforms, commercial heterogeneous GPU services, and large-scale ML infrastructure at a major US technology company. Adopters span education, cloud platform, and tech platform verticals.
2. **vGPU isolation and cost** — **Virtual GPU memory isolation** and heterogeneous scheduling reduce wasted whole-GPU allocation; fractional use, simpler GPU node bootstrap, and (for GPU service providers) delivery without maintaining a proprietary vGPU stack. Specific high-value use cases include bursty interactive development jobs where GPUs are otherwise idle, legacy training workloads on newer GPU generations with underutilized capacity, and offline inference jobs consuming less than 50% of a GPU's resources.
3. **Documentation and community** — Initial installation from public documentation and GitHub was described as straightforward; public docs are sufficient for minimum viable deployment. The project's responsiveness to external contributions reinforces community health signals from [## Contributors and Community](#contributors-and-community)

**Improvements adopters would like to see:**

1. **GPU observability** — Deeper metrics, profiling, and OpenTelemetry-style tracing; continued investment in operational UI (**HAMi WebUI**). At the vGPU slice level, HAMi's metrics endpoint exposes only basic utilization and memory usage.
2. **Performance matrix** — No published benchmark data for expected performance impact across concurrent workload scenarios (multi-GPU training, single-GPU training, mixed training and inference, varying framework configurations). Adopters spent evaluation time running their own benchmarks to understand performance trade-offs before deploying to production which is data that the project could publish to reduce the pre-adoption evaluation burden.
3. **Rollback difficulty at scale** — `hami-device-plugin` replaces the NVIDIA device plugin as a DaemonSet; on large clusters this requires careful batch-by-batch rollout because the device plugin cannot be shut down and replaced atomically. If issues arise mid-rollout, rollback is difficult. This operational complexity is distinct from NVIDIA GPU Operator integration challenges.
4. **KubeVirt** — Stronger **vGPU** support when workloads run under KubeVirt.
5. **Vendor cloud-native maturity** — adopters would like to see continued maturation of  per-vendor drivers, SDKs, and virtualization paths to reduce custom integration work.

Three adopter interviews are complete (minimum met). [#### Adoption](#adoption) remains a status roster only (e.g. [OpenFGA incubation DD](https://github.com/cncf/toc/blob/main/projects/openfga/openfga-incubation-dd.md), [KServe incubation DD](https://github.com/cncf/toc/blob/main/projects/kserve/kserve-incubation-dd.md)).

- **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

Documented integrations ([GTR](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md), [SSA](https://github.com/cncf/toc/blob/main/projects/hami/security-assesment/self-assessment.md), application [#1775](https://github.com/cncf/toc/issues/1775)):


| Integration                                                                                    | CNCF?             | Documentation                                                                                                                                                                                                                   |
| ---------------------------------------------------------------------------------------------- | ----------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Kubernetes** (scheduler extender, device plugin, mutating webhook)                           | yes               | [project-hami.io docs](https://project-hami.io/docs/); [GTR](https://github.com/cncf/toc/blob/main/projects/hami/tech-review/2026-04-14.md)                                                                                     |
| **[Volcano](https://volcano.sh/)**                                                             | yes               | [HAMi Volcano vGPU guide](https://github.com/Project-HAMi/HAMi/blob/master/docs/how-to-use-volcano-vgpu.md); [Volcano GPU sharing](https://github.com/volcano-sh/volcano/blob/master/docs/user-guide/how_to_use_gpu_sharing.md) |
| **[Koordinator](https://koordinator.sh/)**                                                     | yes               | [Device scheduling with HAMi](https://koordinator.sh/docs/user-manuals/device-scheduling-gpu-share-with-hami)                                                                                                                   |
| **[cert-manager](https://cert-manager.io/)** (webhook TLS)                                     | yes               | Optional production hardening per GTR                                                                                                                                                                                           |
| **[KAI Scheduler](https://github.com/kai-scheduler/KAI-Scheduler)**                            | yes               | GPU sharing integration in progress ([KAI-Scheduler#60](https://github.com/kai-scheduler/KAI-Scheduler/pull/60))                                                                                                                |
| **Grafana** / observability                                                                    | yes               | [Dashboard examples](https://github.com/Project-HAMi/HAMi/blob/master/docs/dashboard.md); [HAMi WebUI guide](https://project-hami.io/docs/userguide/hami-webui-user-guide)                                                      |
| Multi-vendor accelerators (NVIDIA, Ascend, MLU, Hygon, Iluvatar, MetaX, Enflame, Cambricon, …) | non-CNCF hardware | [Supported devices](https://project-hami.io/docs/userguide/device-supported); [HAMi-core](https://github.com/Project-HAMi/HAMi-core)                                                                                            |


CNCF scheduling and Kubernetes integration paths are documented for incubation; Volcano and Koordinator guides match application assertions in [#1775](https://github.com/cncf/toc/issues/1775).

**Recommendations (non-blocking):**

1. add a single integrations page that labels **CNCF** vs **non-CNCF** partners (e.g. [wasmCloud integrations](https://wasmcloud.com/docs/integrations/), [wasmCloud incubation DD](https://github.com/cncf/toc/blob/main/projects/wasmcloud/2024-wasmCloud-incubation-dd.md))
2. document KAI Scheduler integration status ( [KAI-Scheduler#60](https://github.com/kai-scheduler/KAI-Scheduler/pull/60))
3. expand HAMi scheduler monitor metrics beyond basic utilization and memory, or document integration paths for adopters to bridge `dcgm-exporter` physical-GPU metrics with HAMi's virtual GPU slice abstraction

#### Adoption

Interview findings are summarized under [TOC verification of adopters](#toc-verification-of-adopters) in [## Ecosystem](#ecosystem). All adopter interview notes are privately kept for TOC use and verification.

##### Adopter 1 - Education

**Status:** Interview complete (April 2026). Anonymous.

##### Adopter 2 - Cloud platform

**Status:** Interview complete (April 2026). Anonymous.

##### Adopter 3 - ML infrastructure

**Status:** Interview complete (June 2026). Anonymous.