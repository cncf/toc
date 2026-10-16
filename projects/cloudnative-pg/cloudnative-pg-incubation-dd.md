# CloudNativePG Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1961)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

**Date:** 2026-10-16 (initial draft 2026-07-05)
**Lead TOC member:** Chad Beaudin
**Due diligence:** Chad Beaudin and Karena Angell
**Adopter interviews:** Karena Angell
**Application:** [cncf/toc#1961](https://github.com/cncf/toc/issues/1961)
**Project:** [cloudnative-pg/cloudnative-pg](https://github.com/cloudnative-pg/cloudnative-pg)
**Website:** https://cloudnative-pg.io
**Level:** Incubation (from Sandbox, accepted 2025-01-21)
**Application filed:** 2025-11-11 by Gabriele Bartolini (EDB)
**Project contacts:** Armando Ruocco, Floor Drees, Francesco Canovai, Gabriele Bartolini, Jonathan Gonzalez (emeritus), Leonardo Cecchi, Marco Nenciarini, Philippe Scorsolini (emeritus, Upbound)

**Related documents:**
- [OpenSSF Best Practices badge #9933](https://www.bestpractices.dev/projects/9933)
- [TAG Data presentation (Sep 2024)](https://www.youtube.com/watch?v=Nsu5Em8Ig88)
- [Security Self-Assessment (Gemara model)](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/.github/security-self-assessment.md)
- [SECURITY-INSIGHTS.yml](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/SECURITY-INSIGHTS.yml)
- [Governance restructuring PR #71](https://github.com/cloudnative-pg/governance/pull/71) (merged)
- [GOVERNANCE.md](https://github.com/cloudnative-pg/governance/blob/main/GOVERNANCE.md): governance, decision-making, subproject, security, and conduct model
- [CONTRIBUTOR_LADDER.md](https://github.com/cloudnative-pg/governance/blob/main/CONTRIBUTOR_LADDER.md): contributor progression and maintainer lifecycle
- [Incubation readiness epic #65](https://github.com/cloudnative-pg/governance/issues/65): tracks DD follow-up items

## Incubation Evaluation Summary for CloudNativePG

### Criteria Evaluation

TOC members Chad Beaudin (lead) and Karena Angell evaluated CloudNativePG for Incubation. Karena Angell conducted the adopter interviews. The project meets the incubation criteria for technical maturity, documentation, ecosystem integration, and community engagement, and has a strong security posture. [PR #71](https://github.com/cloudnative-pg/governance/pull/71) has merged and substantially improves the governance model. Ratification, org-balanced voting, the non-EDB maintainer pathway, and Steering Committee formation remain open in project tracking as recommendations rather than required blockers.

**Notables:**
- 4,818 active contributors, 1,790 contributing organizations (LFX)
- 51 documented adopters spanning 7+ verticals and 10+ countries; 3 hyperscalers (Google Cloud, Microsoft Azure, IBM) build commercial products on CloudNativePG
- OpenSSF Best Practices passing 100%, silver 87% (7 criteria remaining); security health score 91
- Governance restructuring ([PR #71](https://github.com/cloudnative-pg/governance/pull/71)) delivered a federated subproject model, explicit Security Response Team language, updated contributor ladder, subproject maintainer committees, gitvote profiles, and clearer ownership mapping; org-balanced voting and Steering Committee formation remain tracked recommended follow-ups
- No StatefulSets or sidecars; the operator manages Pods and PVCs directly, with LSN-based failover promotion and isolation-aware self-fencing
- CNPG-I plugin interface separates extensibility from core, enabling external backup providers and authentication validators without core modifications
- 96 releases across 16 minor versions; 3 concurrent release branches maintained with coordinated same-day patch releases and RC cycles
- Supply chain security: cosign signatures, Sigstore bundles, SBOM, SLSA provenance attestation
- SECURITY-INSIGHTS.yml v2.2.0 with Gemara-based self-assessment mapped to FINOS Common Cloud Controls
- ROADMAP.md explicitly documents employer influence on prioritization
- LFX Mentorship program documented with active participation; +27% contributors YoY, +33% orgs YoY

**Non-blocking recommendations:**
1. Complete the OpenSSF silver badge (7 criteria remaining: accessibility, i18n, regression tests, secure design documentation, input validation documentation, hardening documentation, assurance case)
2. Submit the self-assessment to TAG Security in the standard template format (the current Gemara format is thorough but may not meet TAG Security expectations for graduation)
3. Use project-domain emails in the SECURITY-INSIGHTS.yml administrator contacts and goreleaser maintainer field (currently `@enterprisedb.com`; the vendor-neutral security@cloudnative-pg.io is already the primary security contact)
4. Expand integration blueprints and reference examples for common CNCF ecosystem combinations, as requested in adopter interviews
5. Improve day-two operational visibility for DBAs (maintenance and audit transparency, data migration support, troubleshooting tooling), as requested in adopter interviews
6. Grow review and issue-triage capacity while keeping a high review bar, particularly as AI-assisted pull request volume increases

### Adoption Evaluation

ADOPTERS.md lists 51 organizations across multiple verticals and geographies, updated as recently as June 2026. Adopters include hyperscalers building products on CNPG (Google Cloud with EDB Community 360, Microsoft Azure with DocumentDB Operator, IBM Cloud Pak), enterprises in production (Tesla, Ericsson, Novo Nordisk), government agencies (Netherlands Tax Administration, Turksat), and research institutions (Vera Rubin Observatory). The adopter list includes both end users and managed service providers building on the platform.

The TOC received three adopter questionnaire responses and completed three adopter interviews, which satisfies the incubation adopter verification requirement. The interviewed organizations are independent of one another: a European municipal IT provider running an internal database-as-a-service platform, a large global technology vendor that embeds CloudNativePG in self-hosted products, and a large U.S. company operating a large fleet of installations. Each interview is summarized in the Adoption section below. Organizations are not named.

All three use CloudNativePG in production, either directly or embedded in products their customers run in production. Adopters cited documentation quality, Kubernetes-native operations, high availability, and a dependable upgrade path as strengths. Individual adopters also highlighted declarative database provisioning, fleet-wide security patching through image catalogs, backup and point-in-time recovery, the project's security posture, and the maintainers' PostgreSQL expertise. All three engage with the project beyond use, through testing, issue reports, pull requests, design feedback, or community channels, and one has been invited to become a subproject maintainer.

Improvement requests were non-blocking for incubation and are reflected in the recommendations above. Two adopters raised maintainer concentration: one sees it as a sustainability concern, and the other has assessed the risk and keeps contingency options. The third sees CNCF ownership of the project's trademarks and intellectual property as the main protection against single-vendor risk. Other requests included more issue-triage capacity, more integration blueprints, cross-cluster switchover orchestration, better DBA maintenance and audit visibility, data migration support, a troubleshooting UI, and keeping review standards high as AI-assisted pull request volume grows.

### Final Assessment

CloudNativePG is technically strong and has credible production adoption, and adopter verification is complete. [PR #71](https://github.com/cloudnative-pg/governance/pull/71) resolves part of the governance concern by restructuring technical authority into subproject committees and documenting security response roles. The DD is ready to proceed to public comment and the TOC vote, with the open governance items tracked as recommended follow-ups rather than required blockers.

**Required before vote:**

1. Public comment period
2. TOC vote

**Complete:**

1. TOC adopter verification: three adopter interviews completed.

**Resolved verification item:**

1. Website vendor-neutrality concern: verified against [CNCF vendor-neutrality guidance](https://contribute.cncf.io/maintainers/community/vendor-neutrality/). The homepage attribution link points to the EDB homepage, and the project [Commercial Support](https://cloudnative-pg.io/support/) page lists multiple vendors.

**Recommended follow-ups:**

1. Governance restructuring ratification ([governance#66](https://github.com/cloudnative-pg/governance/issues/66))
2. Org-balanced voting implementation ([governance#68](https://github.com/cloudnative-pg/governance/issues/68))
3. Non-EDB maintainer pathway formalization ([governance#69](https://github.com/cloudnative-pg/governance/issues/69))
4. Steering Committee formation and first meeting with published minutes ([governance#70](https://github.com/cloudnative-pg/governance/issues/70))

**Post-incubation:** 6-month progress update on org-balanced voting in practice, Steering Committee operation, external maintainer pathway progress, and LFX org dependency trend.

## Application Process Principles

### Suggested

- [x] **Engage with domain-specific TAG(s) to present the technical architecture of the project.**
  - Presented to TAG Data on 2024-09-25. Recording: https://www.youtube.com/watch?v=Nsu5Em8Ig88
  - **Status: Met** (via TAG Data presentation).

### Required

- [ ] **Complete a [General Technical Review (GTR)](../../operations/project-reviews/general-technical-questions.md).**
  - No GTR on `cncf/toc` main branch. The TAG Data presentation covers technical architecture but is not a formal GTR.
  - The application's architecture documentation (cloudnative-pg.io/documentation/current/architecture/, contribute/technical-architecture.md) provides GTR-equivalent coverage.
  - CloudNativePG filed its application on 2025-11-11, before the GTR became a required incubation item on 2026-04-16 ([cncf/toc@dffcced](https://github.com/cncf/toc/commit/dffcced)). The template in effect at the time of application required either a TAG presentation or a GTR, and the project satisfied that requirement with the TAG Data presentation on 2024-09-25.
  - **Status: Not formally completed. Not required at the time of application; application documentation covers the same ground and the DD lead assessment substitutes. Non-blocking.**

- [ ] **Complete a [Governance Review](../../operations/project-reviews/governance-review-template.md).**
  - No formal Governance Review has been completed.
  - DD lead conducted a governance assessment as part of this evaluation (see Governance Assessment section below).
  - CloudNativePG filed its application on 2025-11-11, before the Governance Review became a required incubation item on 2026-04-16 ([cncf/toc@dffcced](https://github.com/cncf/toc/commit/dffcced)).
  - **Status: Not formally completed. Not required at the time of application; the DD lead assessment substitutes. Non-blocking.**

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**
  - Website: cloudnative-pg.io (not EDB-branded); GitHub org: cloudnative-pg; Releases: GitHub Releases; Documentation: cloudnative-pg.io via GitHub Pages; Slack: #cloudnativepg-users on CNCF Slack; Meetings: LFX Zoom + CNCF calendar; Governance: separate cloudnative-pg/governance repo.
  - Website vendor-neutrality concern verified resolved 2026-10-13 against CNCF guidance: the homepage EDB logo is historical attribution linking to the EDB homepage, and the [Commercial Support](https://cloudnative-pg.io/support/) page lists multiple vendors.
  - **Status: Met.**

- [x] **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
  - Application references moving-level criteria. Accepted to Sandbox 2025-01-21.
  - **Status: Met.**

- [x] **Due Diligence Review.**
  - This document. Adopter verification complete; ready for public comment / TOC vote.
  - **Status: Complete for DD draft.**

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**
  - Installation/upgrade, Quickstart, Architecture, Use cases, Operator capability levels (all under cloudnative-pg.io/documentation/current/), plus the CNPG Playground (github.com/cloudnative-pg/cnpg-playground).
  - **Status: Met.**

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review if completed as a suggested item prior to application.

### Governance Assessment (DD lead)

CloudNativePG's governance is documented in a dedicated [governance repository](https://github.com/cloudnative-pg/governance), with the current model captured in [GOVERNANCE.md](https://github.com/cloudnative-pg/governance/blob/main/GOVERNANCE.md) and [CONTRIBUTOR_LADDER.md](https://github.com/cloudnative-pg/governance/blob/main/CONTRIBUTOR_LADDER.md). Following the TOC's triage recommendations, the project undertook a governance restructuring ([PR #71](https://github.com/cloudnative-pg/governance/pull/71)), transitioning from a flat maintainer-council model to a federated subproject model. This addresses much of the TOC's feedback; the org-balance and steering components are tracked separately as recommended follow-ups.

**Merged in PR #71:** 4 Subproject Maintainer Committees (Core Projects, Supply Chain, Community/Docs/Ecosystem, Extensibility), each with technical authority over its own repositories; a restructured contributor ladder (Community Participant, Contributor, Reviewer, Component Owner, Subproject Maintainer) with Steering Committee as a representative role rather than a rung; a named and defined Security Response Team; a per-subproject ownership model with generated `COMPONENT_OWNERS.md` / `CONTRIBUTORS.md`; and governance vote templates / gitvote profiles for amendments and maintainer-lifecycle actions.

**Organizational concentration:** LFX organizational dependency is 71%. All 5 active top-level maintainers are EDB employees, but the project has 9 non-EDB authors contributing in the last 30 days and 10 formal contributors spanning 7 countries. PR #71 improves the governance structure and technical authority model; external maintainer elevation and org-balanced governance remain tracked follow-ups ([governance#68](https://github.com/cloudnative-pg/governance/issues/68), [#69](https://github.com/cloudnative-pg/governance/issues/69), [#70](https://github.com/cloudnative-pg/governance/issues/70)). Under the **incubation** criteria, single-organization concentration is **not a blocking item** (the "≥2 organizations" and org-balance-mechanism requirements are graduation-stage); it is recorded here as a graduation-readiness item that is already being addressed.

### Suggested

- [ ] **Complete a Governance Review.** No formal Governance Review has been completed; the DD lead governance assessment above is used in its place. **Status: Not formally completed.**

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**
  - Adopted gitvote for public transparent votes; emeritus process exercised (Jonathan Gonzalez, Philippe Scorsolini); Component Owners tier added; [PR #71](https://github.com/cloudnative-pg/governance/pull/71) replaced the flat model with a federated subproject model. Org-balanced voting and Steering Committee formation remain tracked ([#68](https://github.com/cloudnative-pg/governance/issues/68), [#70](https://github.com/cloudnative-pg/governance/issues/70)).
  - **Status: Met for incubation.**

- [x] **Clear and discoverable project governance documentation.** Governance repo with [GOVERNANCE.md](https://github.com/cloudnative-pg/governance/blob/main/GOVERNANCE.md) at root; every sub-project points to the canonical governance repo. **Status: Met.**

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.** Public developer meetings documented; gitvote active; maintainer/component-owner lists current; community meetings on LFX/CNCF calendar. **Status: Met.**

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.** Values include "Community over product/company," openness, fairness, inclusivity. Steering org caps / org-balanced voting remain open ([#70](https://github.com/cloudnative-pg/governance/issues/70), [#68](https://github.com/cloudnative-pg/governance/issues/68)) as recommended follow-ups. **Status: Met for incubation.**

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.** [PR #71](https://github.com/cloudnative-pg/governance/pull/71) documents subproject decisions, public vote workflows, amendments, and maintainer-lifecycle actions; private mailing-list votes documented for security/conduct. Org-balanced voting remains open ([#68](https://github.com/cloudnative-pg/governance/issues/68)). **Status: Met for incubation.**

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).** Subproject Maintainers voted by committee; Component Owners assigned by subproject maintainers; Contributors by simple majority; Steering Committee: one representative per subproject plus Community Representatives elected by ranked-choice vote; removal documented for all tiers. **Status: Met.**

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).** All phases documented in [GOVERNANCE.md](https://github.com/cloudnative-pg/governance/blob/main/GOVERNANCE.md). **Status: Met.**

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.** 2 emeritus maintainers (Jonathan Gonzalez/EDB, Philippe Scorsolini/Upbound); lifecycle exercised for departures; ownership changes via public governance issues/PRs. Non-EDB maintainer pathway formalization tracked ([#69](https://github.com/cloudnative-pg/governance/issues/69)). **Status: Met for lifecycle usage; external pathway is a recommended follow-up.**

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.** 4 formally defined subprojects with per-subproject maintainer committees and dedicated files listing repositories/maintainers/components; MAINTAINERS.md restructured into per-subproject tables. Steering representation tracked ([#70](https://github.com/cloudnative-pg/governance/issues/70)). **Status: Met.**

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.** MAINTAINERS.md lists 5 active + 2 emeritus with names, GitHub handles, emails, affiliations. **Status: Met.**

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.** 5 active maintainers + 18 component owners for a project with 8,929 stars, 96 releases, 18 sub-components. **Status: Met.**

- [x] **Code and Doc ownership in GitHub and elsewhere matches documented governance roles.** Every project has GOVERNANCE.md pointing to the canonical repo; component owners have Write access to their repos; maintainers have Maintain/Admin. **Status: Met.**

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.** CODE_OF_CONDUCT.md references the CNCF CoC; enforcement documented (private maintainer mailing list; CNCF staff for maintainer-involved violations). **Status: Met.**

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.** Cross-linked from main repo CODE_OF_CONDUCT.md and governance docs. **Status: Met.**

- [x] **All subprojects, if any, are listed.** All projects listed under the cloudnative-pg GitHub org; COMPONENT_OWNERS.md provides ownership mapping. **Status: Met.**

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.** Tiers: Community Participant, Contributor, Reviewer, Component Owner, Subproject Maintainer; Component Owner provides folder-scoped CODEOWNERS recognition; Steering Committee kept off the ladder as a representative role; separate [CONTRIBUTOR_LADDER.md](https://github.com/cloudnative-pg/governance/blob/main/CONTRIBUTOR_LADDER.md). **Status: Met.**

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.** CONTRIBUTING.md redirects to the governance repo's canonical CONTRIBUTING.md; separate contribute/ directory; DCO documented. **Status: Met.**

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.** CNCF Slack: #cloudnativepg-users. **Status: Met.**

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.). List any non-public communications channels and what their special purpose is.** Public: CNCF Slack, GitHub Issues, GitHub Discussions. Non-public: security@cloudnative-pg.io (vulnerability reports), conduct@cloudnative-pg.io (CoC), maintainer mailing list (governance). **Status: Met.**

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.** LFX Zoom integration; CNCF calendar integration confirmed. **Status: Met.**

- [x] **Documentation of how to contribute, with increasing detail as the project matures.** CONTRIBUTING.md, contribute/ (development environment, E2E testing, release procedure), DCO, LFX Mentorship program docs. **Status: Met.**

- [x] **Demonstrate contributor activity and recruitment.** DevStats + LFX Insights: 4,818 active contributors, 1,790 contributing orgs; YoY +27% contributors, +33% orgs, +21% stars, +79% forks; LFX Mentorship active; 9 non-EDB authors in last 30 days. **Status: Met.**

## Engineering Principles

### Suggested

- [x] **Roadmap change process is documented.** ROADMAP.md documents the planning process (GitHub Projects + Milestones), explicit prioritization drivers, and acknowledges that employer priorities influence the roadmap. **Status: Met.**

- [x] **History of regular, quality releases.** 96 releases across 16 minor versions; minor releases every 3-4 months; 3 concurrent release branches (N, N-1, N-2); coordinated same-day patch releases; RCs since v1.24.0; most recent v1.30.0 (2026-06-29) with coordinated v1.29.2/v1.28.4. **Status: Met.**

### Required

- [x] **Document project goals and objectives that illustrate the project's differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  - TAG Data presentation (2024-09-25) covers differentiation; website documents scope, use cases, operator capability levels; FAQ documents differentiation from other PostgreSQL operators. No formal GTR; application docs provide equivalent coverage.
  - **Status: Met.**

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.** Comprehensive documentation at cloudnative-pg.io; use cases and operator capability levels documented. **Status: Met.**

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.** ROADMAP.md with a GitHub Projects board; Milestones for per-release planning. **Status: Met.**

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._** Technical architecture (contribute/technical-architecture.md) and architecture/controller/container-image docs on the site. No formal GTR; application docs provide equivalent coverage. **Status: Met.**

- [x] **Document the project's release process.** contribute/release_procedure.md: trunk-based development, cherry-picked backports, feature freeze, ordered release execution, manual smoke testing, supply chain security (cosign, SBOM, SLSA), multi-artifact coordination (Helm, OLM, docs, Krew, API tagging). **Status: Met.**

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance if completed as a suggested item prior to application.

### Suggested

- [ ] **Complete a joint security assessment with TAG Security and Compliance.**
  - Not completed. The current self-assessment uses the Gemara model, which is thorough but not the standard TAG Security format; producing a standard-format assessment is recommended before graduation. A self-assessment is sufficient at incubation.
  - **Status: Not completed (non-blocking at incubation).**

### Required

- [x] **Clearly defined and discoverable process to report security issues.** SECURITY.md documents reporting to security@cloudnative-pg.io (private maintainer Google Group), with GitHub Security Advisories for disclosure. **Status: Met.**

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)** 2FA enforced on the GitHub org (per OpenSSF badge); DCO required; branch protection on main; CodeQL + Snyk static analysis. **Status: Met.**

- [x] **Document assignment of security response roles and how reports are handled.** PR #71 added explicit Security Response Team documentation; MAINTAINERS.md lists its membership and triage role; GOVERNANCE.md documents closed meetings / private list for security matters; SECURITY-INSIGHTS.yml lists administrators and reporting metadata. The team is currently all-EDB, a consequence of the broader governance concentration to be addressed via the non-EDB maintainer pathway rather than a separate security blocker. **Status: Met.**

- [x] **Document Security Self-Assessment.** Present at [`.github/security-self-assessment.md`](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/.github/security-self-assessment.md), Gemara model v1.2.0 mapped to FINOS Common Cloud Controls Core v2025.10, with threat and capability catalogs; referenced from SECURITY-INSIGHTS.yml v2.2.0. **Status: Met.**

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.** [Passing badge 100%](https://www.bestpractices.dev/en/projects/9933) (67/67, achieved 2025-01-29, badge #9933); silver 87% (7 criteria remaining). **Status: Met.**

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)** [ADOPTERS.md](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/ADOPTERS.md) lists 51 organizations across verticals and geographies, updated 2026-06-17. **Status: Met.**

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**
  - 3 questionnaire responses received; 3 TOC interviews completed.
  - **Status: Met.**

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [x] **TOC verification of adopters.**
  - Complete: three adopter interviews conducted. See the Adoption section.
  - **Status: Met.**

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.** Dedicated docs for Cilium (CNCF Graduated) and External Secrets Operator (CNCF Sandbox); additional integrations (Prometheus, cert-manager, VolumeSnapshots, Helm, OLM) documented. **Status: Met.** (Recommendation: expand the "CNCF Projects Integrations" section beyond Cilium/ESO.)

#### Adoption

Three TOC adopter interviews completed; adopter verification is complete for the incubation DD.

##### Adopter 1 - European municipal IT provider / Public sector

Running CloudNativePG in production since mid-2023, this adopter provides database-as-a-service for internal application teams and affiliated organizations. It selected CNPG after comparing PostgreSQL operators, citing documentation quality as the primary differentiator. A staged migration from a legacy database platform required data-type remapping but completed without rollback.

CNPG standardizes previously manual database provisioning through reusable declarative configurations, reducing operational work and accelerating integration. Releases are evaluated through phased non-production and production environments during maintenance windows. The adopter contributes through testing, issue reports, and smaller pull requests, and described maintainer encouragement as a positive path to community participation.

The adopter views maintainer concentration and issue-triage capacity as sustainability concerns, although it remains confident in the project's direction. It could migrate to another operator if required but lacks capacity to supply maintainers itself. Requested improvements include broader maintainer diversity, more cross-project integration blueprints, and a management layer for cross-cluster or multi-cloud switchovers.

##### Adopter 2 - Large global technology vendor / Technology

This adopter embeds CloudNativePG as an operator in its products (resilient metadata store, some relational workloads) across OpenShift and other Kubernetes platforms; multi-year use includes pre-open-source downstream collaboration. Customer-operated product upgrades orchestrate the database lifecycle alongside the larger product, so the adopter depends on resilient stable/LTS releases with forward- and backward-compatible upgrade and rollback paths.

The adopter highlighted CNPG's custom workload controller, high availability, security posture, and deep PostgreSQL expertise. Early collaboration contributed lessons around CRD/schema evolution, Operator Lifecycle Manager and Helm deployment complexity, and isolation of multiple operators on shared clusters. The adopter has built internal competence and participates through design and contribution feedback, while retaining several contingency options if the primary vendor stopped maintaining the project.

Desired improvements include day-to-day DBA maintenance and auditing transparency, data repatriation and migration support, and a troubleshooting UI. The adopter reported positive collaboration with the community, while noting that its long-standing relationship with the project may not represent every user's experience.

##### Adopter 3 - Large U.S. company

This adopter operates a large fleet of CloudNativePG installations, adopted as part of the company's strategy of investing in open source and Kubernetes. The team finds the documentation high quality and works from it directly, and also participates in the community Slack. It has built standard processes so that installations are consistent across the fleet.

The adopter described the main benefit as bringing Kubernetes operating practices to the database tier. Its leading example was image catalogs: when a security patch is released, PostgreSQL container and base images are rebuilt automatically and can be rolled out across the fleet without logging in to individual servers, in contrast to the largely manual patching still common for databases. Backups use object storage and volume snapshots configured through CRDs, with scheduled backups and point-in-time recovery to a specific timestamp. The adopter is migrating its fleet from the original built-in backup system to the newer plugin-based backup architecture, which it described as a normal part of the software lifecycle.

Asked how it would respond if the primary vendor stopped funding development, the adopter said the most important protection is CNCF ownership of the project's trademarks and intellectual property, which makes a single-company relicensing far more difficult. It pointed to other open source projects where a license change forced users to fork and rebuild a community under a new name. It noted that other large users it has spoken with could work together to keep the project maintained, likely at a slower feature pace. It considers the project production-ready today.

The adopter is aware of the published governance restructuring proposal and has been invited to become a maintainer of the extensions and playground subprojects. It recommended making more use of the playground, which uses two kind clusters to simulate a two-region deployment with high availability, monitoring, and backups, for training and demonstrations. Its main concern for the project was the growing volume of pull requests generated with AI tools by contributors without a deep understanding of PostgreSQL internals. It stressed that a database project must keep a high bar for review to avoid corruption and data loss, while still mentoring contributors who are prepared to invest the time.

<!-- Specification Project Considerations section intentionally omitted: CloudNativePG is a Kubernetes operator, not a specification project defining normative behavior for independent implementation. -->

---

## Appendix A: Technical Assessment

### Architecture

CloudNativePG is a Kubernetes operator built in Go that manages the full lifecycle of PostgreSQL clusters.

**Core design decisions:** No StatefulSets (operator manages Pods and PVCs directly, enabling LSN-based promotion rather than ordinal index); no sidecars (an Instance Manager Go binary runs as PID 1, eliminating Patroni/repmgr/Stolon dependencies); immutable, non-root, read-only-root-filesystem application containers.

**High availability:** Single-primary / multiple hot-standby via native streaming replication; three auto-managed Services (`-rw`, `-ro`, `-r`); LSN-based failover promotion with per-cluster Lease mutex to prevent split-brain; isolation-aware liveness probes (primary self-fences on full network partition); synchronous replication (priority `first` and quorum `any`); PodDisruptionBudgets; `pg_rewind` reintegration.

**Cross-cluster / DR:** Replica clusters across Kubernetes clusters via WAL shipping or streaming; controlled switchover or emergency failover; the operator's authority is scoped to a single Kubernetes cluster, so cross-cluster failover requires external orchestration.

**Backup and PITR:** WAL archiving to object stores (S3, Azure Blob, GCS, MinIO) via the Barman Cloud CNPG-I plugin; full backups via CNPG-I and VolumeSnapshots; PITR with parallel WAL fetching; bootstrap modes `initdb`/`recovery`/`pg_basebackup`.

**Extensibility:** The CNPG-I plugin interface adds functionality (backup providers, auth validators) without modifying the core operator.

**Additional CRDs:** `Backup`, `ScheduledBackup`, `Database`, `DatabaseRole`, `Pooler` (PgBouncer), `Publication`, `Subscription`, `ImageCatalog`, `ClusterImageCatalog`. **Monitoring:** Prometheus exporter (port 9187), Grafana dashboards, JSON-structured logging.

### Technical Strengths

1. No-StatefulSet / no-sidecar architecture with Mermaid diagrams and source-reference tables
2. LSN-based failover with Lease-based split-brain prevention
3. Isolation-aware self-fencing on network partition
4. CNPG-I plugin interface for ecosystem extensibility
5. CSI volume snapshot integration for near-instant cloning
6. Supply chain security (cosign, Sigstore, SBOM, SLSA)
7. E2E test suite including failure/chaos scenarios
8. 80%+ code coverage (per OpenSSF badge)
9. Level 5 "Auto Pilot" Operator Capability
10. Regular KubeCon presence 2022-2026, including talks with major adopters

### Sub-projects and Ecosystem

The cloudnative-pg GitHub org contains 34 repositories, including the core operator, cnpg-i / cnpg-i-machinery (plugin interface + framework), plugin-barman-cloud, plugin-pvc-backup, barman-cloud, machinery, api, governance, charts, docs, website, artifacts, cnpg-playground, container-image repos (postgres/postgis/extensions/trunk/pgbouncer/pgcrash), grafana-dashboards, ciclops, chaos-testing, daggerverse, community-operators, and postgres-keycloak-oauth-validator.

### CNCF Ecosystem Integrations

Dedicated docs: Cilium (CNCF Graduated), covering CiliumNetworkPolicy for pod-to-pod security; External Secrets Operator (CNCF Sandbox), covering automated password rotation and pushing secrets to an external KMS with reload labels. Additional (documented, no dedicated page): Prometheus (metrics exporter on 9187), cert-manager (TLS), Kubernetes VolumeSnapshots (CSI backup), Helm, OLM. Recommendation: expand the "CNCF Projects Integrations" section beyond Cilium/ESO.

## Appendix B: Security Assessment

**Security documentation:** SECURITY.md (reporting to security@cloudnative-pg.io, GitHub Security Advisories); SECURITY-INSIGHTS.yml v2.2.0 (last reviewed 2026-02-25); Security Self-Assessment at `.github/security-self-assessment.md` (Gemara v1.2.0 mapped to FINOS CCC Core v2025.10, with threat and capability catalogs; non-standard format); Security Response Team named in governance and listed in MAINTAINERS.md.

**OpenSSF Best Practices:** passing 100% (67/67, 2025-01-29, badge #9933); silver 87% (7 remaining: accessibility, i18n, regression tests, secure design docs, input validation docs, hardening docs, assurance case).

**Access control:** GitHub-org 2FA enforcement; DCO; branch protection; CodeQL + Snyk. **Supply chain:** cosign, Sigstore bundles, SBOM, SLSA provenance, Renovate dependency updates.

## Appendix C: Adopter Assessment

**ADOPTERS.md:** 51 organizations (updated 2026-06-17) across cloud/infra (EDB, Google Cloud, Microsoft Azure, IBM, Akamai, Nutanix), automotive (Tesla), telecom (Ericsson, Turksat), healthcare/pharma (Novo Nordisk), financial services (PZU, GEICO Tech, Alpcot, Rootline), government (Netherlands Tax Administration, GARR), science (Vera Rubin Observatory), and MSPs (Clastix, Tembo, OptimaData, Enix, Aenix, Syself, Giant Swarm, Bitnami/VMware). Hyperscalers Google Cloud, Microsoft, and IBM build commercial products on CNPG.

**Questionnaire status:** 5-7 adopters requested (triage 2026-07-03); 3 responses received; 3 TOC interviews completed; adopter verification complete for the incubation DD.

**Community endorsements:** 4 GitHub endorsements: Matteo Bianchi (@mbianchidev), Ian Stanton (@ianstanton), @zoe-tek, Jerome Petazzoni (@jpetazzo).

## Appendix D: Gaps Table

| Criteria | Status | Severity | Notes |
|----------|--------|----------|-------|
| TOC adopter verification | Complete | Resolved | 3 questionnaire responses; 3 interviews completed |
| Website vendor neutrality | Complete | Resolved | Homepage EDB logo links to EDB homepage (historical attribution); Commercial Support page lists multiple vendors |
| Governance restructuring ratified | Not complete | Recommended / non-blocking | governance#66 open |
| Org-balance mechanism | Not complete | Recommended / non-blocking | governance#68 open |
| Non-EDB maintainer pathway | Not complete | Recommended / non-blocking | governance#69 open |
| Steering Committee formed | Not complete | Recommended / non-blocking | governance#70 open; first meeting w/ minutes recommended |
| Security response roles documented | Complete | Resolved | PR #71 added explicit Security Response Team docs |
| Metadata email addresses | Not complete | Recommended / non-blocking | SECURITY-INSIGHTS.yml / goreleaser use `@enterprisedb.com`; primary security contact is vendor-neutral |
| GTR | Not formal | Non-blocking | Not required when the application was filed (2025-11-11); TAG presentation and application docs provide equivalent coverage |
| Governance Review | Not formal | Non-blocking | Not required when the application was filed (2025-11-11); DD lead assessment substitutes |
| TAG Security standard self-assessment | Non-standard format | Recommended / non-blocking | Gemara format rigorous; standard format may be needed for graduation |
| OpenSSF silver badge | 87% complete | Recommended / non-blocking | 7 criteria remaining; primarily documentation |
| CNCF integration blueprints | Partial | Recommended / non-blocking | Adopter requested more practical examples |
| Day-two operational transparency / UX | Partial | Recommended / non-blocking | Adopter requested DBA maintenance/auditing visibility, data repatriation/migration, troubleshooting UI |
| Review and issue-triage capacity | Ongoing | Recommended / non-blocking | Adopters raised triage capacity and review rigor as AI-assisted PR volume grows |
