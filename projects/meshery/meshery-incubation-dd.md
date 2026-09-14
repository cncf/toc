# Meshery Incubation Due Diligence

Meshery is a cloud native infrastructure design and operations platform that enables teams to design, deploy, and manage Kubernetes-based infrastructure and service meshes through a visual interface and extensible model-based architecture. Originally focused on service mesh management, the project has evolved to support 380+ integrations across the cloud native ecosystem. It is used in production and development across multiple organizations, see the [public adopters list](https://github.com/meshery/meshery/blob/master/ADOPTERS.md).

- **Incubation application:** [cncf/toc#1386](https://github.com/cncf/toc/issues/1386)
- **Project:** Sandbox since 2021-06-22 → applying **Incubation** · [meshery/meshery](https://github.com/meshery/meshery) · [meshery.io](https://meshery.io/) · [CNCF landscape](https://landscape.cncf.io/?item=app-definition-and-development--application-definition-image-build--meshery)
- **LFX Insights (past year):** [Meshery project dashboard](https://insights.linuxfoundation.org/project/meshery) - contributor and activity signals cited in [## Contributors and Community](#contributors-and-community)
- **Sub-projects in scope:** Core org repos listed in [GOVERNANCE.md#subprojects-and-domains](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#subprojects-and-domains) and [MAINTAINERS.md](https://github.com/meshery/meshery/blob/master/MAINTAINERS.md) - including [schemas](https://github.com/meshery/schemas), [meshkit](https://github.com/meshery/meshkit), [meshsync](https://github.com/meshery/meshsync), [meshery-operator](https://github.com/meshery/meshery-operator), and mesh adapters (Istio, Linkerd, …). **Extensions** in [meshery-extensions](https://github.com/meshery-extensions) are governed separately per [GOVERNANCE.md#extensions](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#extensions-githubcommeshery-extensions). **Kanva** is no longer a core Meshery component (see [governance review](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md)).
- **Governance review on `main`:** [governance-review/2026-06-05.md](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md) ([#2187](https://github.com/cncf/toc/pull/2187) merged 2026-06-08, closes [#2140](https://github.com/cncf/toc/issues/2140)) - **Status: Satisfactory**
- **Security SSA:** Submitted by the project and reviewed by the TOC member - [cncf/toc#2264](https://github.com/cncf/toc/pull/2264) (authored by Meshery maintainers, 2026-08-04). The criterion (project-authored SSA, reviewed by the TOC) is met; maintainers addressed several non-blocking clarifications in-thread (see [## Security](#security)).
- **OpenSSF Best Practices:** [passing - project 3564](https://www.bestpractices.dev/en/projects/3564) (TOC verified badge level 2026-06-08)
- **Adopter verification (interviews):** Minimum three independent adopter interviews per the incubation process; see [ADOPTERS.md](https://github.com/meshery/meshery/blob/master/ADOPTERS.md)

## Sub-project scope (incubation DD)

| Scope | Repos / components | DD evaluation |
| :--- | :--- | :--- |
| **In scope - core platform** | [meshery/meshery](https://github.com/meshery/meshery), [meshery-operator](https://github.com/meshery/meshery-operator), [meshsync](https://github.com/meshery/meshsync), [meshkit](https://github.com/meshery/meshkit), [schemas](https://github.com/meshery/schemas), mesh adapters (Istio, Linkerd, …) per [GOVERNANCE.md#subprojects-and-domains](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#subprojects-and-domains) | Full criteria in this DD |
| **In scope - extensions (governed separately)** | [meshery-extensions](https://github.com/meshery-extensions) org - lifecycle per [GOVERNANCE.md#extensions](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#extensions-githubcommeshery-extensions) | Evaluated for extensibility model; not every extension repo audited line-by-line |
| **Out of scope / not core** | **Kanva** UI extension (separated per [governance review](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md)) | Not required for incubation bar |
| **License** | [Apache-2.0](https://github.com/meshery/meshery/blob/master/LICENSE) on primary repo - **TOC:** no blocking third-party license exceptions identified in public review; escalate to the foundation if any are discovered during review |

## Incubation Evaluation Summary for Meshery

### Criteria Evaluation

Karena Angell, TOC Member and Chair, conducted the due diligence of Meshery for **Incubation** as TOC sponsor ([#1386](https://github.com/cncf/toc/issues/1386)). This document consolidates the section evaluations and the adopter interviews and is prepared for public comment.

The project demonstrates substantial incubation maturity in governance, community, documentation, and ecosystem adoption signals. The following implementations are noteworthy:

- **Application Process:** TAG Infrastructure architecture presentation (2021-11-04); sandbox acceptance [#1264](https://github.com/cncf/toc/pull/1264) (2024-03-01); merged [Governance Review](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md) (**Satisfactory**, 2026-06-08).
- **Governance and Maintainers:** Exemplary [Maintainer Council](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md) model across core subprojects; [Certified Meshery Contributor (CMC)](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#certified-meshery-contributor-cmc) program; maintainers across many organizations per [MAINTAINERS.md](https://github.com/meshery/meshery/blob/master/MAINTAINERS.md); governance iterated through [#19772](https://github.com/meshery/meshery/pull/19772) during matriculation review.
- **Contributors and Community:** Public Slack, mailing lists, weekly [meshery.io/calendar](https://meshery.io/calendar) meetings, layered [contributor ladder](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#roles-and-the-contributor-ladder), and [LFX Insights](https://insights.linuxfoundation.org/project/meshery) activity.
- **Engineering Principles:** Discoverable [overview](https://docs.meshery.io/project/overview), [architecture](https://docs.meshery.io/concepts/architecture), [ROADMAP.md](https://github.com/meshery/meshery/blob/master/ROADMAP.md), and [release process](https://docs.meshery.io/project/contributing/build-and-release); active semver releases (e.g. [v1.0.39](https://github.com/meshery/meshery/releases/tag/v1.0.39) 2026-06-08).
- **Security:** Discoverable reporting path; [OpenSSF passing](https://www.bestpractices.dev/en/projects/3564). SSA submitted as [cncf/toc#2264](https://github.com/cncf/toc/pull/2264) and under public review, with maintainers addressing non-blocking clarifications in-thread.
- **Ecosystem:** Public [ADOPTERS.md](https://github.com/meshery/meshery/blob/master/ADOPTERS.md) with production narratives (HPE, Intel, DigitalOcean, Ericsson, F5, Red Hat, …); [integrations catalog](https://docs.meshery.io/extensibility/integrations).

The following actions were blocking governance matriculation but **resolved during the governance review** (see [#2187](https://github.com/cncf/toc/pull/2187)):

- Kanva separation from core Meshery documentation; [GOVERNANCE.md](https://github.com/meshery/meshery/pull/19772) alignment with CNCF project templates ([#19461](https://github.com/meshery/meshery/issues/19461), [#19447](https://github.com/meshery/meshery/issues/19447), [#19460](https://github.com/meshery/meshery/issues/19460)).

**Blocking items resolved during due diligence:**

1. **SSA submitted to `cncf/toc`** - **Resolved.** Submitted as [cncf/toc#2264](https://github.com/cncf/toc/pull/2264) (2026-08-04); under public review with maintainers addressing non-blocking clarifications in-thread. See [## Security](#security).
2. **Vendor neutrality labeling** ([meshery/meshery#21137](https://github.com/meshery/meshery/issues/21137)) - **Resolved.** The project implemented the "Official" vs "Community" extension tiering it described in its [CNCF blog post](https://www.cncf.io/blog/2026/03/04/scaling-organizational-structure-with-mesherys-expanding-ecosystem/) (2026-03-04). Support tiers are now expressed three ways across `meshery-extensions` repos - a GitHub custom property, a repo topic, and a README badge - with drift detection in CI ([meshery-extensions/.github#34](https://github.com/meshery-extensions/.github/pull/34), 2026-08-06). The `meshery` org profile vendor reference was removed ([meshery/.github#68](https://github.com/meshery/.github/pull/68), 2026-08-06), the recognition/badge system now resolves to `meshery.io/community/recognition`, and the contributing entry point resolves to the project-native `docs.meshery.io/project/contributing`. Extension lifecycle governance for `meshery-extensions` (formal policies, archival process, vendor-neutral registry) remains a graduation-track item.

Completion of this due diligence document and the public comment period satisfies the Due Diligence Review criterion. A General Technical Review was not a requirement when Meshery submitted its incubation application (2024-03-01); see [Application Process Principles](#application-process-principles).

**DD Recommendations (non-blocking unless noted).**

| Priority | Recommendation | Effort |
|----------|---------------|--------|
| ~~**Blocking**~~ | ~~SSA submitted to `cncf/toc` by project~~ - **Resolved.** [cncf/toc#2264](https://github.com/cncf/toc/pull/2264) submitted 2026-08-04, under public review. | ~~Medium~~ |
| ~~**Blocking**~~ | ~~Vendor neutrality labeling~~ - **Resolved.** "Official" vs "Community" support tiers implemented across `meshery-extensions` (GitHub custom property, repo topic, README badge) with CI drift detection ([meshery-extensions/.github#34](https://github.com/meshery-extensions/.github/pull/34)); `meshery` org profile vendor reference removed ([meshery/.github#68](https://github.com/meshery/.github/pull/68)); badge/recognition system and contributing link resolve to project domains ([meshery/meshery#21137](https://github.com/meshery/meshery/issues/21137)). | ~~Low~~ |
| High | Public roadmap - **partially addressed.** ROADMAP.md now links to GitHub milestones, and v1.1.0 through v1.5.0 carry due dates from 2026-09 through 2028-09. Residual: keep ROADMAP.md in sync with the milestone set and close completed milestones. | Low |
| High | Scope reframing - **partially addressed.** The repository description now reads "the cloud native manager," moved off "service mesh manager." Align the application narrative and site positioning with the broadened "cloud native infrastructure design and operations platform" scope. | Low |
| Medium | Structure simplification - archive dead adapter repos, consolidate active adapters, label vendor academy repos. Reduce from 41 active repos to ~15-20. | Low-Medium |
| Medium | CODEOWNERS - no CODEOWNERS file exists; monorepo structure needs review routing per domain | Low |
| Low | Backlog triage - 1,103 open issues, 723 open PRs; document a triage policy. Milestones are now current (v1.1.0 through v1.5.0, all dated). | Medium |
| Low | CI workflow consolidation - 58 GitHub Actions workflows; review for overlap and consolidation | Medium |
| n/a | Suggest considering project rename - "Meshery" derives from "mesh" and reinforces legacy service mesh identity. Worth raising with maintainers as part of scope reframing. | Low |

**Track for graduation (non-blocking at incubation).** Verified against the closed vendor-neutrality tracking issue [meshery/meshery#21137](https://github.com/meshery/meshery/issues/21137); its "not blocking at incubation" items carry forward here. At incubation the expectation is a documented plan; at graduation, implementation.

- Extension lifecycle governance for `meshery-extensions` - health criteria, maturity labels, and an archival process. The `meshery-extensions` GOVERNANCE.md currently points to the main MAINTAINERS file; the general maintainer lifecycle (stepping down, inactivity, emeritus) applies, and a formal extension-repo lifecycle policy is the remaining piece.
- Full subproject maturity list with status - an official list of core versus extension subprojects (governance review, 2026-06-05). Extension support tiers (Official, Community) are already labeled through the support-labels machinery; a full maturity-status list is the remaining piece.
- Vendor-neutral container registry - **in progress.** Container images now publish to the project-named `meshery/` Docker Hub namespace (`meshery/meshery` and the meshsync, operator, and adapter images actively built), and the `layer5io/` image namespace is retired. Remaining: confirm the `meshery/` account is under project-controlled or CNCF-neutral ownership, and complete the residual `github.com/layer5io/...` module-path and docs cleanup.
- Contributor concentration and succession - the founder still contributes about five times the second contributor (13,013 versus 2,427 commits). GOVERNANCE.md documents the maintainer lifecycle (stepping down, inactivity, emeritus) and makes organizational concentration visible through MAINTAINERS.md; an explicit succession plan is the remaining piece.
- TAG Security joint assessment - none on file in `cncf/tag-security`; pursue before graduation.

**Non-blocking recommendations - three-month written update:** None of the recommendations above is a condition of incubation. As a lightweight checkpoint, the TOC requests a written update from the maintainers by **2026-11-22** (approximately three months after due diligence) on the roadmap and progress for the non-blocking recommendations, including the open documentation-guidance thread on components that materially affect the security posture. 

**Security (TOC member evaluation, 2026-08-04):** From a security standpoint, Meshery satisfies incubation criteria for reporting path, access control, and OpenSSF badge. The project maintains a discoverable reporting path through [docs.meshery.io/project/security-vulnerabilities](https://docs.meshery.io/project/security-vulnerabilities) and `security@meshery.dev`, and the [OpenSSF Best Practices badge](https://www.bestpractices.dev/en/projects/3564) is at passing level (verified 2026-06-08). **The SSA has been submitted and reviewed by the TOC member** as [cncf/toc#2264](https://github.com/cncf/toc/pull/2264) (2026-08-04), authored by Meshery maintainers following the TAG Security template. The criterion is a project-authored self-assessment reviewed by the TOC, and it is met. The review generated substantive engagement and the maintainers addressed several non-blocking clarifications in-thread (trust model, actor model, credential-at-rest posture, TLS/mTLS posture, and reclassification of development-process controls). The review surfaced no blocking incubation findings; the extension provenance items are consistent with the existing track-for-graduation framing. A TAG joint security assessment has not been filed; that is acceptable at incubation and should be planned for before graduation. Criterion-level evidence and the review outcomes appear in [## Security](#security).

**Governance (TOC member evaluation, 2026-06-08):** Meshery meets incubation governance expectations. A stand-alone [governance review](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md) was completed by the [Project Reviews Subproject](https://github.com/cncf/toc/tree/main/toc_subprojects/project-reviews-subproject) and merged to `cncf/toc` `main` on 2026-06-08 ([#2187](https://github.com/cncf/toc/pull/2187)) with status **Satisfactory**. Exemplary Maintainer Council model, six maintainers across six organizations, CMC contributor certification program, and governance iterated through [#19772](https://github.com/meshery/meshery/pull/19772) during matriculation review. **Resolved during due diligence:** the incubation vendor-neutrality items - contributing link, badge/recognition system on `badges.layer5.io`, `meshery-extensions` support-tier labeling, and the `meshery` org profile vendor reference - are complete; project metadata now uses project domains (`docs.meshery.io`, `meshery.io`) per [CNCF vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/). **Track for graduation:** Extension lifecycle governance for `meshery-extensions` - formal lifecycle policies, an archival process, a vendor-neutral registry for community extensions, and a subproject listing with unified governance scope. At incubation, document a plan; at graduation, implement it. Criterion-level evidence appears in [## Governance and Maintainers](#governance-and-maintainers).

**Community (TOC member evaluation, 2026-06-08):** Meshery presents a healthy contributor community for incubation. The project documents a contributor ladder, clear contribution paths, and multiple public channels including Slack ([slack.meshery.io](https://slack.meshery.io/)), community forum, mailing lists, and weekly meetings. **TOC member verified (2026-06-08):** [LFX Insights](https://insights.linuxfoundation.org/project/meshery) shows strong multi-org contributor activity. The [CMC certification program](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#certified-meshery-contributor-cmc) is a distinctive recruitment mechanism. Contributor concentration should be monitored - the founder has approximately 5x more contributions than the #2 contributor, which is common at incubation but should be tracked for graduation. Criterion-level evidence appears in [## Contributors and Community](#contributors-and-community).

**Ecosystem (TOC member evaluation, 2026-08-04):** Meshery has a strong public adopter list ([ADOPTERS.md](https://github.com/meshery/meshery/blob/master/ADOPTERS.md)) with 15+ organizations across multiple verticals including HPE Labs, HPE Security Engineering, Intel, DigitalOcean, Ericsson, and others. **TOC verification of adopters is complete** - three independent adopter interviews conducted, meeting the incubation minimum. The 380+ integrations catalog and extensibility model demonstrate ecosystem breadth. Criterion-level evidence appears in [## Ecosystem](#ecosystem).

**Engineering Principles (TOC member evaluation, 2026-08-04):** Engineering Principles evidence meets incubation criteria. Architecture, installation, concepts, and release process documentation is discoverable from [meshery.io](https://meshery.io/) and GitHub. **Roadmap:** `ROADMAP.md` now links to GitHub milestones with target dates (v1.1.0 through v1.5.0). TOC recommends keeping ROADMAP.md in sync with the milestone set and closing completed milestones. **Scope evolution:** The project has evolved from "service mesh manager" to "cloud native infrastructure design and operations platform" - this is legitimate scope evolution supported by the 380+ model-based integrations, but the repo structure (13 service mesh adapter repos) still tells the old story. Structure simplification would reinforce the new narrative. Criterion-level evidence appears in [## Engineering Principles](#engineering-principles).

### Adoption Evaluation

Three independent adopter interviews are complete, meeting the incubation minimum. Interview findings are summarized under [TOC verification of adopters](#toc-verification-of-adopters). TOC verification rests on these interviews, not on application assertions or ADOPTERS.md text alone.

### Final Assessment

Meshery meets the incubation criteria. Governance is **Satisfactory** per the merged [governance review](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md); community, engineering principles, security, and ecosystem criteria are documented with criterion-level evidence below. The security self-assessment is submitted ([cncf/toc#2264](https://github.com/cncf/toc/pull/2264)) and under public review, with maintainers addressing non-blocking clarifications in-thread. The incubation vendor-neutrality items are resolved. Three independent adopter interviews are complete. This due diligence records non-blocking recommendations - structure simplification, keeping the public roadmap current, and scope reframing from "service mesh manager" to "cloud native infrastructure design and operations platform" - and a set of graduation-track items (extension lifecycle governance, vendor-neutral registry, TAG Security joint assessment). On this basis the due diligence is prepared for public comment.

## Application Process Principles

**Meets all Required criteria.**

### Suggested

- [x] **Engage with domain-specific TAG(s) to present the technical architecture of the project.**
  - TAG Infrastructure presentation on **2021-11-04** - [recording](https://www.youtube.com/watch?v=FPMde6EHcJU&t=3936s) (per [#1386](https://github.com/cncf/toc/issues/1386)).

### Required

- **N/A** - **[General Technical Review (GTR)](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/general-technical-questions.md).**
  - A GTR was not a requirement when Meshery submitted its incubation application (2024-03-01), so it does not apply to this application.

- [x] **Complete a [Governance Review](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/governance-review-template.md).**
  - Revised **2026-06-05** - [governance-review/2026-06-05.md](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md) on `main` ([#2187](https://github.com/cncf/toc/pull/2187), merged 2026-06-08). **Status: Satisfactory.**

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

**Met.** Vendor-neutrality values are documented at [GOVERNANCE.md#vendor-neutrality](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#vendor-neutrality), and the project took concrete steps during the governance review (Kanvas separation from core documentation, Slack rebranded to project identity). A review of both GitHub orgs (`meshery`, `meshery-extensions`) identified incubation and graduation findings; the incubation findings were resolved during due diligence and are recorded below.

**Blocking for incubation - resolved during due diligence** (project metadata on vendor domains):

| Finding | Location | Resolution |
|---------|----------|------------|
| Recognition system on vendor domain | [GOVERNANCE.md](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md) referenced `badges.layer5.io` | **Resolved** - recognition/badge system now resolves to `meshery.io/community/recognition` |
| Contributing link to vendor site | Application [#1386](https://github.com/cncf/toc/issues/1386) linked to `layer5.io/community/newcomers` | **Resolved** - contributing entry point resolves to `docs.meshery.io/project/contributing` |
| Org profile promotes commercial product | `meshery`/`meshery-extensions` org profiles referenced Layer5 Cloud | **Resolved** - vendor reference removed ([meshery/.github#68](https://github.com/meshery/.github/pull/68)) |
| Extension support tiers unlabeled | `meshery-extensions` repos did not distinguish "Official" vs "Community" | **Resolved** - support tier applied as a GitHub custom property, repo topic, and README badge across repos, with CI drift detection ([meshery-extensions/.github#34](https://github.com/meshery-extensions/.github/pull/34)) |

**Non-blocking for incubation, track for graduation:**

| Finding | Graduation expectation |
|---------|---------------------|
| Container images on vendor Docker namespace (`layer5io/`) | Publish images to a project-controlled or CNCF-neutral registry |
| No extension lifecycle governance for `meshery-extensions` | Define lifecycle policies with health criteria, maturity labels, and an archival process |
| No official subproject maturity list | Publish a maturity-status list for core and extension subprojects |
| Vendor-branded archived adapter repos ([meshery-octarine](https://github.com/meshery-extensions/meshery-octarine), [meshery-osm](https://github.com/meshery-extensions/meshery-osm), [meshery-nginx-sm](https://github.com/meshery-extensions/meshery-nginx-sm)) | Archive unmaintained extensions and publish extensions to a vendor-neutral registry |
| Academy repos as vendor-specific training ([play-digitalocean](https://github.com/meshery-extensions/play-digitalocean), [play-tcs](https://github.com/meshery-extensions/play-tcs)) | Archive or label as vendor-maintained |
| Kanvas repos in extensions org ([kanvas-snapshot](https://github.com/meshery-extensions/kanvas-snapshot), [kanvas-action](https://github.com/meshery-extensions/kanvas-action)) | Label as vendor-maintained extensions with clear ownership |
| Release infrastructure dependencies | Verify any maintainer can publish release artifacts without relying on vendor-controlled accounts |

**Extension model public framing:** The project published [Scaling Organizational Structure with Meshery's Expanding Ecosystem](https://www.cncf.io/blog/2026/03/04/scaling-organizational-structure-with-mesherys-expanding-ecosystem/) on the CNCF blog (2026-03-04), framing `meshery-extensions` with "Official" vs "Community" extension labels. This is encouraging - the project has publicly articulated the two-org governance model and the tiered support expectations. The graduation expectation is that vendor-maintained extensions, including any commercial products hosted in the extensions org, are clearly labeled with their vendor ownership, and that the commercial relationship between core maintainers and the vendor is disclosed. During due diligence the project implemented the "Official" vs "Community" labeling described in the blog post, applying a support tier to `meshery-extensions` repos as a GitHub custom property, a repo topic, and a README badge, with CI drift detection.

The project acted on the vendor-neutrality feedback: the incubation blocking findings are resolved, and the graduation-track items are documented below. At graduation the expectation is full resolution: vendor-neutral artifact registry, extension lifecycle governance with archival policies, all vendor-specific branding removed from upstream, critical release artifacts on CNCF-owned or project-controlled accounts, and commercial extensions clearly labeled with vendor ownership.

- [x] **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
  - Sandbox accepted **2021-06-22**; moving-levels expectations acknowledged in application **2024-03-01** per [#1264](https://github.com/cncf/toc/pull/1264).

- [x] **Due Diligence Review.**
  - Complete: all incubation blocking items resolved, three independent adopter interviews done, and the document prepared for the two-week public comment period preceding the TOC vote.

Completion of this due diligence document, resolution of concerns raised, and presentation for public comment satisfies the Due Diligence Review criteria.

*This document, finalized for public comment. Incubation blocking items (SSA submission and TOC review, adopter verification, vendor-neutrality labeling) are resolved.*

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

[Installation](https://docs.meshery.io/installation), [Concepts](https://docs.meshery.io/concepts), [Guides](https://docs.meshery.io/guides), [Contributing](https://docs.meshery.io/project), [Reference](https://docs.meshery.io/reference) - discoverable from [meshery.io](https://meshery.io/) and GitHub org.

## Governance and Maintainers

**Meets all Required criteria; non-blocking recommendations below.**

**Governance evaluation (TOC member, 2026-06-08):** Satisfactory per merged [Governance Review](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md) ([#2187](https://github.com/cncf/toc/pull/2187)), conducted as a stand-alone [Project Reviews Subproject](https://github.com/cncf/toc/tree/main/toc_subprojects/project-reviews-subproject) review merged to `cncf/toc` `main`. Checkboxes below incorporate that review; TOC sponsor verified alignment **2026-06-08**.

**Blocking - resolved during due diligence:**

1. **vendor neutrality labeling** - **Resolved.** The project implemented the "Official" vs "Community" extension tiering it described in its [CNCF blog post](https://www.cncf.io/blog/2026/03/04/scaling-organizational-structure-with-mesherys-expanding-ecosystem/) (2026-03-04): a support tier is applied to `meshery-extensions` repos as a GitHub custom property, a repo topic, and a README badge, with CI drift detection ([meshery-extensions/.github#34](https://github.com/meshery-extensions/.github/pull/34)). The `meshery` org profile vendor reference was removed ([meshery/.github#68](https://github.com/meshery/.github/pull/68)).
2. **contributing link** - **Resolved.** The contributing entry point resolves to `docs.meshery.io/project/contributing` rather than the vendor domain.
3. **badge/recognition system** - **Resolved.** The recognition/badge system resolves to `meshery.io/community/recognition` per [CNCF vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).

**Non-blocking recommendations (track for graduation):**

4. **extension lifecycle governance** - define policies for `meshery-extensions` repos including health criteria, maturity labels, and an archival process, and list all subprojects with unified governance scope. At incubation, document a plan; at graduation, implement it.
5. add explicit **maturity labels** for core vs extension subprojects (governance review, 2026-06-05) - publish a maturity-status list. Track for graduation.
6. add a CODEOWNERS file for monorepo review routing per domain
7. document **contributor concentration** mitigation and succession plan in GOVERNANCE.md (founder has ~5x more contributions than #2 contributor)

### Suggested

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

[GOVERNANCE.md commit history](https://github.com/meshery/meshery/commits/master/GOVERNANCE.md); significant matriculation update in [#19772](https://github.com/meshery/meshery/pull/19772). Prior examples: [#12162](https://github.com/meshery/meshery/pull/12162), [#7968](https://github.com/meshery/meshery/pull/7968), [#4992](https://github.com/meshery/meshery/pull/4992).

The project has been iterating on governance since before sandbox acceptance (2021-06-22), and commit history shows ongoing updates - a healthy sign that practice is evolving with experience. The matriculation update during the governance review demonstrates responsiveness to TOC feedback.

- [x] **Clear and discoverable project governance documentation.**

[GOVERNANCE.md](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md) covers council structure, subprojects, extensions, meetings, and decision processes.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Maintainer votes in [meshery.io Google Group](https://groups.google.com/a/meshery.io/g/maintainers/search?q=vote); maintainer PRs to [MAINTAINERS.md](https://github.com/meshery/meshery/commits/master/MAINTAINERS.md) (e.g. [#14886](https://github.com/meshery/meshery/pull/14886), [#16359](https://github.com/meshery/meshery/pull/16359)); weekly meeting notes linked from [GOVERNANCE.md#meetings](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#meetings).

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

[GOVERNANCE.md#vendor-neutrality](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#vendor-neutrality).

**TOC verification (updated 2026-08):** Vendor-neutrality values are documented in governance. The project took concrete steps during the governance review (Kanvas separated from core documentation, Slack rebranded to project identity) and completed the incubation vendor-neutrality items during due diligence: the contributing link and the badge/recognition system resolve to project domains (`docs.meshery.io`, `meshery.io`), and `meshery-extensions` support tiers are labeled ([meshery-extensions/.github#34](https://github.com/meshery-extensions/.github/pull/34), [meshery/.github#68](https://github.com/meshery/.github/pull/68)). Community calendar and forum already use meshery.io domains. Graduation-track vendor-neutrality items (vendor-neutral registry, full extension lifecycle governance) are recorded below.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

[GOVERNANCE.md#how-specific-kinds-of-decisions-are-made](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#how-specific-kinds-of-decisions-are-made): maintainer lifecycle for leadership; maintainer approval for contributions; documented CNCF request path; **2/3 Maintainer Council vote** for governance changes.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

Security response via `security@meshery.dev` - [GOVERNANCE.md#security-response](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#security-response). Extension maintainers - [GOVERNANCE.md#extensions-maintainers](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#extensions-maintainers). Domain contribution guides at [docs.meshery.io/project/contributing](https://docs.meshery.io/project/contributing/).

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

[GOVERNANCE.md#roles-and-the-contributor-ladder](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#roles-and-the-contributor-ladder), [becoming a maintainer](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#becoming-a-maintainer), [emeritus maintainers](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#emeritus-maintainers).

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

[MAINTAINERS.md history](https://github.com/meshery/meshery/commits/master/MAINTAINERS.md).

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

[GOVERNANCE.md#subproject-lifecycle](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#subproject-lifecycle); extensions lifecycle at [GOVERNANCE.md#extensions](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#extensions-githubcommeshery-extensions).

**Recommendation (non-blocking):** add explicit maturity labels for core vs extension subprojects (governance review).

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

[MAINTAINERS.md](https://github.com/meshery/meshery/blob/master/MAINTAINERS.md) - domains and affiliations per subproject.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

At least **three maintainers per core domain** documented in MAINTAINERS.md (governance review).

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

Per application [#1386](https://github.com/cncf/toc/issues/1386) and governance review.

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

[CODE_OF_CONDUCT.md](https://github.com/meshery/meshery/blob/master/CODE_OF_CONDUCT.md); referenced from [GOVERNANCE.md#code-of-conduct](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#code-of-conduct).

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

Cross-linked from GOVERNANCE.md (governance review).

- [x] **All subprojects, if any, are listed.**

Core subprojects in [GOVERNANCE.md#subprojects-and-domains](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#subprojects-and-domains) and [MAINTAINERS.md](https://github.com/meshery/meshery/blob/master/MAINTAINERS.md); extensions listed under [extensions maintainers](https://github.com/meshery/meshery/blob/master/MAINTAINERS.md#extensions-maintainers-meshery-extensions).

## Contributors and Community

**Meets all Required criteria; non-blocking recommendations below.**

Note: community criteria also assessed in [Governance Review](https://github.com/cncf/toc/blob/main/projects/meshery/governance-review/2026-06-05.md#contributors-and-community).

**Community metrics ([LFX Insights](https://insights.linuxfoundation.org/project/meshery), [GitHub](https://github.com/meshery/meshery)):** The primary repository has **11,430** stars and **307** contributors. The project maintains a weekly meeting cadence, public Slack, community forum, and mailing lists. The [CMC certification program](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#certified-meshery-contributor-cmc) is a distinctive contributor recruitment mechanism. Six maintainers across six organizations are documented in [MAINTAINERS.md](https://github.com/meshery/meshery/blob/master/MAINTAINERS.md). These figures provide context for the criterion-by-criterion evaluation below.

**Recommendations (non-blocking):**

1. ~~migrate contributing entry point from `layer5.io/community/newcomers` to `docs.meshery.io/project/contributing`~~ - **Resolved during due diligence;** the contributing entry point resolves to `docs.meshery.io/project/contributing`
2. document **contributor concentration** trajectory and broaden organizational participation before graduation
3. cite public GitHub PRs for reviewer/maintainer promotions for transparency

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

[GOVERNANCE.md#roles-and-the-contributor-ladder](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#roles-and-the-contributor-ladder) - Organization Member, Maintainer, subproject leads, Maintainer Council.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

[GOVERNANCE.md#how-to-contribute](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#how-to-contribute) · [docs.meshery.io/project/contributing](https://docs.meshery.io/project/contributing)

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

[GOVERNANCE.md#communication-channels](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#communication-channels) - Slack ([slack.meshery.io](https://slack.meshery.io/)), forum, mailing lists.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.). List any non-public communications channels and what their special purpose is.**

Documented in GOVERNANCE.md; application [#1386](https://github.com/cncf/toc/issues/1386) references website community page.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

[meshery.io/calendar](https://meshery.io/calendar) · [GOVERNANCE.md#meetings](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#meetings)

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

[docs.meshery.io/project/contributing](https://docs.meshery.io/project/contributing).

**TOC verification (updated 2026-08):** The contributing entry point resolves to the project-native `docs.meshery.io/project/contributing` per [CNCF vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/). The prior application link to a commercial-entity domain was migrated during due diligence.

- [x] **Demonstrate contributor activity and recruitment.**

[Contributor graphs](https://github.com/meshery/meshery/graphs/contributors) · [CMC program](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#certified-meshery-contributor-cmc) · [LFX Insights - Meshery](https://insights.linuxfoundation.org/project/meshery)

**Contributor activity:** Meshery shows strong multi-org contributor activity on [LFX Insights](https://insights.linuxfoundation.org/project/meshery) and the [CNCF project page](https://www.cncf.io/projects/meshery/) (large contributor and organization counts, high health score). These signals corroborate recruitment activity and do not substitute for the adopter interviews.

## Engineering Principles

**Meets all Required criteria; non-blocking recommendations below.**

Current release at DD review: [v1.0.64](https://github.com/meshery/meshery/releases) (weekly cadence). Criterion-level notes below confirm incubation satisfaction and record non-blocking recommendations.

**Recommendations (non-blocking):**

1. keep the public roadmap current - ROADMAP.md now links to GitHub milestones with target dates (v1.1.0 through v1.5.0); keep it in sync and close completed milestones
2. add CODEOWNERS for monorepo review routing per domain (Go server, React UI, CLI, docs)
3. **structure simplification** - archive dead adapter repos, consolidate active adapters into a single monorepo, and archive or label vendor academy repos as vendor-maintained. Reduce from 41 active repos to ~15-20. See DD recommendations table.
4. **CI workflow consolidation** - 58 GitHub Actions workflows; review for overlap and consolidation
5. **backlog triage** - 1,103 open issues, 723 open PRs; document a triage policy

### Suggested

- [x] **Roadmap change process is documented.**

[ROADMAP.md](https://github.com/meshery/meshery/blob/master/ROADMAP.md)

**Roadmap note:** ROADMAP.md now links to GitHub milestones with target dates (v1.1.0 through v1.5.0, dated 2026-09 through 2028-09). Keep ROADMAP.md in sync with the milestone set and close completed milestones. See DD recommendations table.

- [x] **History of regular, quality releases.**

[GitHub releases](https://github.com/meshery/meshery/releases) - e.g. [v1.0.39](https://github.com/meshery/meshery/releases/tag/v1.0.39) (2026-06-08), [v1.0.38](https://github.com/meshery/meshery/releases/tag/v1.0.38) (2026-06-05).

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**
  - [project overview](https://docs.meshery.io/project/overview)

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**
  - [logical concepts](https://docs.meshery.io/concepts/logical), [architecture](https://docs.meshery.io/concepts/architecture)

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

[ROADMAP.md](https://github.com/meshery/meshery/blob/master/ROADMAP.md)

**TOC verification:** ROADMAP.md now links to GitHub milestones with target dates (v1.1.0 through v1.5.0). The criterion is satisfied; keeping ROADMAP.md in sync with the milestone set is a non-blocking recommendation. See recommendations above.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**
  - [architecture docs](https://docs.meshery.io/concepts/architecture)

- [x] **Document the project's release process.**

[Build and release](https://docs.meshery.io/project/contributing/build-and-release)

## Security

**Meets all Required criteria; SSA submitted by the project and reviewed by the TOC member; non-blocking recommendations below.**

Note: SSA submitted as [cncf/toc#2264](https://github.com/cncf/toc/pull/2264) (authored and submitted by the project maintainers, following the [TAG Security template](https://github.com/cncf/tag-security/blob/main/community/assessments/guide/self-assessment.md)) and reviewed by the TOC member, which meets the criterion. A TAG Security [joint assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/guide/joint-assessment.md) has not been filed (acceptable at incubation; recommend before graduation).

**SSA public review outcomes (2026-08):** The self-assessment PR drew substantive review from TAG Security contributors. The maintainers addressed the following non-blocking clarifications in-thread; these are documentation and framing improvements rather than blocking incubation findings. As of 2026-08-21 the maintainers reported all review feedback addressed, with the self-assessment text updates captured in follow-up changes; reviewers confirmed the load-bearing authorization-enforcement thread as resolved.

The TOC notes this as an **exemplary community response**. The review drew deep, multi-reviewer technical engagement, and the maintainers responded promptly and thoroughly - reading and citing source code to answer reviewer questions, adding user-facing documentation, broadening the actor model, and updating the self-assessment in-flight. This is the standard of open, evidence-based security engagement the TOC hopes to see in the moving-levels process.

| Item raised | Maintainer response | Status |
|-------------|--------------------|--------|
| Trust model - are core and community plugins both "trusted" within an installation? | Confirmed the framing and added a docs section, [Trusting an extension](https://docs.meshery.io/installation/production/security-hardening/#trusting-an-extension), to make the trust boundary explicit for users | Addressed |
| Actor model - Kubernetes cluster, cluster administrator, and other users as actors beyond shipped components | Additional actor tables developed to broaden the actor inventory | Addressed |
| Development-process controls (DCO sign-off, branch protection, CI scanning gate) listed under security functions | Agreed these are secure-development-practices controls; input sanitization for query ordering remains a product control | Reclassification agreed |
| Credential-at-rest for kubeconfigs/contexts | Clarified: local provider stores contexts and tokens under `~/.meshery` protected by filesystem permissions only (stated in Non-goals); remote provider forwards credentials subject to the provider's at-rest controls; maintainers to add a note at the credential-collection point. Server-side control noted: credential ownership is bound to the authenticated user so a client-supplied `userId` cannot redirect a credential | Addressed with docs note |
| TLS/mTLS posture for server-adapter and broker channels | Clarified: Broker (NATS) token authentication is on by default via a generated Kubernetes Secret; TLS termination is left to the operator to avoid install friction; guidance in the [security hardening](https://docs.meshery.io/installation/production/security-hardening/) production guide. Maintainers filed [meshery/meshery#21444](https://github.com/meshery/meshery/issues/21444) to add transport encryption to the operator-managed broker | Clarified; follow-up filed |
| Self-assessment does not link to existing security documentation | Maintainers supplied the hardening, operational-readiness, authentication, networking, and authorization-extensibility references for inclusion | Addressed with links |
| GraphQL WebSocket origin policy (cross-site WebSocket hijacking) | Resolved by architecture change: the WebSocket transport was retired ([meshery/meshery#20587](https://github.com/meshery/meshery/pull/20587), 2026-07-11) and GraphQL subscriptions moved to REST and Server-Sent Events over the authenticated middleware chain, so the permissive-origin upgrader no longer exists. The Docker Desktop extension's in-VM socket proxy keeps a permissive `CheckOrigin` but is reachable only from within the Docker Desktop VM | Addressed |
| Permission-key enforcement location (UI vs server) | Clarified: the provider authenticates and resolves the capability set per organization; Meshery UI gates rendering for UX; server-side enforcement is per-operation in `server/models/remote_provider.go`, where each gated operation tests the provider-supplied capability set before proceeding. Actor tables and the named enforcement point were added to the self-assessment | Addressed; reviewer-confirmed |

**Deferred to roadmap / track for graduation:** Extension provenance guarantees for out-of-tree and community extensions - digest validation, signature verification, and mandated vulnerability scanning in the [meshery-extensions](https://github.com/meshery-extensions) CI - are not implemented today. Maintainers indicated these can be placed on the roadmap. Sandboxing exists for dynamically loaded UI and server plugins, and curation-level designations ("official", "verified", "community") provide a first user-facing signal. This is consistent with the extension-lifecycle-governance graduation track recorded elsewhere in this DD.

**Non-blocking items noted by reviewers for the record:** Two thin spots were raised as non-blocking and were explicitly not held against the assessment. First, multi-tenant isolation - what refuses a principal scoped to one organization requesting a resource owned by another, and whether that check is per-record or per-collection; the capability model answers what a role may do, not across which tenant boundary. Second, durability of centralized query-order sanitization - a lint rule, typed builder, or failing test to keep a new call site from concatenating an `ORDER BY` without the shared sanitizer, which is what makes it a control rather than a convention. Both are recorded here as track-for-graduation hardening.

**Recommendations:**

1. complete the in-flight self-assessment updates that incorporate the review clarifications (trust-model docs link, actor and adversary tables, credential-at-rest note, TLS/mTLS note, documentation links, development-practices reclassification, and the named server-side enforcement point); maintainers stated on 2026-08-21 that all review feedback is addressed with these updates captured
2. pursue a TAG Security **joint security assessment** before graduation (see the [TAG Security joint assessment guide](https://github.com/cncf/tag-security/blob/main/community/assessments/guide/joint-assessment.md))
3. document SBOM delivery path to consumers (BOM workflow `bom.yaml` exists but delivery is not documented)
4. document server-side permission-key enforcement to complement the UI gating described in the authorization reference
5. add extension provenance controls (digest/signature verification, mandated CI vulnerability scanning for `meshery-extensions`) to the roadmap - track for graduation
6. document operator guidance on which components and configuration choices materially affect the security posture, and how users assess them (open review thread on [cncf/toc#2264](https://github.com/cncf/toc/pull/2264)) - non-blocking; complete as a documentation follow-up

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

[Security vulnerabilities](https://docs.meshery.io/project/security-vulnerabilities) · `security@meshery.dev` per GOVERNANCE

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

[GOVERNANCE.md#github-project-administration](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#github-project-administration) (per application)

- [x] **Document assignment of security response roles and how reports are handled.**

[Security vulnerabilities](https://docs.meshery.io/project/security-vulnerabilities) · [GOVERNANCE.md#security-response](https://github.com/meshery/meshery/blob/master/GOVERNANCE.md#security-response)

- [x] **Document Security Self-Assessment.**

**Met.** Submitted by the project and reviewed by the TOC member - [cncf/toc#2264](https://github.com/cncf/toc/pull/2264) (2026-08-04), authored by Meshery maintainers following the [TAG Security template](https://github.com/cncf/tag-security/blob/main/community/assessments/guide/self-assessment.md).

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

[OpenSSF project 3564](https://www.bestpractices.dev/en/projects/3564) - **passing** (TOC verified 2026-06-08).

## Ecosystem

**Meets all Required criteria; TOC adopter verification complete (three interviews); non-blocking recommendations below.**

For **Incubation**, adopters may be in dev/test or production use; **graduation** expects production use. Meshery is listed in the [CNCF landscape](https://landscape.cncf.io/?item=app-definition-and-development--application-definition-image-build--meshery).

**Public adopters:** [ADOPTERS.md](https://github.com/meshery/meshery/blob/master/ADOPTERS.md) (15+ organizations with production narratives).

**Recommendations (non-blocking, track for graduation):**

1. add per-adopter **adoption level** (prod/dev/trial) on the public list
2. label CNCF vs non-CNCF integrations on the docs site
3. **structure simplification** - archive dead service mesh adapter repos; archive or label vendor academy repos ([play-digitalocean](https://github.com/meshery-extensions/play-digitalocean), [play-tcs](https://github.com/meshery-extensions/play-tcs)) as vendor-maintained with clear ownership.
4. **vendor-neutral container registry** - evaluate migration from the Docker Hub `layer5io/` namespace to a project-controlled namespace.

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

[ADOPTERS.md](https://github.com/meshery/meshery/blob/master/ADOPTERS.md) - includes production narratives (HPE Labs, HPE Security Engineering, Intel, DigitalOcean, Ericsson, Asteria Aerospace, BootLabs, Metabob, Intraserve, and others).

Public list exceeds incubation expectations for breadth. Per-adopter adoption level (prod/dev/trial) is not indicated on the public list today.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

Public ADOPTERS.md alone lists **many** independent organizations; a roster was provided with the application [#1386](https://github.com/cncf/toc/issues/1386). **TOC interviews complete** - three independent adopters interviewed, satisfying formal verification for incubation.

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

**Adoption Evaluation** and TOC verification below depend on TOC member interviews with adopters - not on application assertions or ADOPTERS.md text alone.

- [x] **TOC verification of adopters.**

**Complete:** Three independent adopter interviews conducted, satisfying the incubation minimum. Interview notes and recordings are kept privately for internal TOC use and verification. [#### Adoption](#adoption) lists interview status by industry vertical only (anonymous).

**Strengths adopters cited:**

Adopters value the project's abstraction layer over native Kubernetes configurations, which reduces the learning curve for teams without deep Kubernetes expertise and enables faster onboarding. One adopter built a full internal developer platform on Meshery serving multiple internal teams, selecting it specifically for its multi-tenancy and collaboration capabilities. The ability to manage organizations, teams, and contractors with granular access permissions was a core requirement. The extensibility model allows organizations to create custom models for tools not natively supported by the project, which proved important for adopters integrating with internal or niche tools outside the CNCF ecosystem. Standardized multi-tool load generation and benchmarking provided significant value for teams evaluating service mesh performance trade-offs in development environments - the project supports multiple load-generating tools that can be exercised independently, with results in a standardized format that enables comparison across configurations. The visualization capabilities (designing deployments, configuring filters, showcasing infrastructure to other technology teams) were used for operational work and for communicating architecture decisions to stakeholders. One adopter successfully white-labeled the UI for their internal platform using existing documentation, with no friction from the project's customization policies. 

The community was consistently described as responsive. One adopter contributed a public extension back to the project's extensions org and has a team member serving as maintainer of that repo. Drawing on hands-on use at a prior organization, one adopter relied on the project's performance tooling to benchmark an optimized distribution of a service mesh against its upstream baseline, generating standardized visual comparison reports that demonstrated the value of their acceleration work to customers and replaced manual, script-driven testing. The same adopter used the drag-and-drop dashboard to install that distribution for customers without requiring them to run manual command-line steps, and to enable or disable individual features with a single action so customers could verify the effect directly from the dashboard and performance reports, taking a new environment from zero to a working deployment without deep infrastructure knowledge. This adopter also described the community as notably open: after the maintainers learned of the usage they engaged directly, and the adopter went on to contribute code and documentation and served for a period as a maintainer within the project's community.

**Improvements adopters would like to see:**

Adopters identified documentation gaps in the project's internal workings, describing certain subsystems as a "black box." While installation and getting-started documentation is sufficient for minimum viable deployment, deeper operational documentation for load testing internals and advanced features is lacking. Feature discoverability was a recurring theme across interviews: adopters reported discovering supported capabilities only after asking questions in Slack, suggesting the project's actual functionality significantly exceeds what is communicated through its documentation and website. One adopter explicitly distinguished this as a marketing problem rather than a documentation problem - the documentation itself is good, but the project does not adequately surface what features exist. 

Also, one adopter experienced installer instability during early adoption, with crashes and stuck installations that were resolved through Slack engagement and an upstream pull request. The project has since stabilized with the v1 GA release line, but adopters maintain their own staging environments subscribed to the edge release channel, running custom test suites against each release before deploying to production. 

Adopters with teams using Terraform or GitLab workflows identified gaps in first-class support for these tools - Terraform integration is limited to Kubernetes operators published by cloud providers, and Git integration currently supports only GitHub, though the project has communicated plans to implement a generic Git backend. Reflecting on performance-testing work, one adopter asked the project to broaden the set of supported load-generation tools, noting the default tool did not meet every requirement at the time, including TLS handshake testing. The same adopter recommended more comprehensive guides for the dashboard and its graphical interface, observing that it was not always immediately clear what each component provides.

_Cross-check vs [ADOPTERS.md](https://github.com/meshery/meshery/blob/master/ADOPTERS.md). Public list alone is **not** TOC verification._

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

Documented integrations (application [#1386](https://github.com/cncf/toc/issues/1386), [integrations catalog](https://docs.meshery.io/extensibility/integrations)):

| Integration | CNCF? | Documentation |
|-------------|-------|---------------|
| **Kubernetes** (deployment, management, visualization) | yes | [docs.meshery.io](https://docs.meshery.io/)
| **Istio** (service mesh adapter) | yes | [meshery-istio adapter](https://github.com/meshery-extensions/meshery-istio) |
| **Linkerd** (service mesh adapter) | yes | [meshery-linkerd adapter](https://github.com/meshery-extensions/meshery-linkerd) |
| **Consul** (service mesh adapter) | yes | [meshery-consul adapter](https://github.com/meshery-extensions/meshery-consul) |
| **Kuma** (service mesh adapter) | yes | [meshery-kuma adapter](https://github.com/meshery-extensions/meshery-kuma) |
| **Cilium** (service mesh adapter) | yes | [meshery-cilium adapter](https://github.com/meshery-extensions/meshery-cilium) |
| **NSM** (Network Service Mesh adapter) | yes | [meshery-nsm adapter](https://github.com/meshery-extensions/meshery-nsm) |
| **Prometheus** (metrics) | yes | Via integrations catalog |
| **380+ model-based integrations** | mixed | [integrations catalog](https://docs.meshery.io/extensibility/integrations) - managed through `models/` and `schemas` repo |

CNCF service mesh and Kubernetes integration paths are documented for incubation; the 380+ model-based integrations represent the modern integration architecture that has superseded the per-adapter approach.

**Recommendations (non-blocking):**

1. add a single integrations page that labels **CNCF** vs **non-CNCF** partners

#### Adoption

Interview findings are summarized under [TOC verification of adopters](#toc-verification-of-adopters) in [## Ecosystem](#ecosystem). All adopter interview notes are privately kept for TOC use and verification.

##### Adopter 1 - Technology

**Status:** Interview complete (June 2026)

##### Adopter 2 - IT services

**Status:** Interview complete (July 2026)

##### Adopter 3 - Technology

**Status:** Interview complete (August 2026)

