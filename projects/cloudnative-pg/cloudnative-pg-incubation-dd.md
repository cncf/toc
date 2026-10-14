# CloudNativePG Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1961)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for CloudNativePG

### Criteria Evaluation

Karina Angell ([@angellk](https://github.com/angellk)) and Chad Beaudin ([@chadbeaudin](https://github.com/chadbeaudin)) conducted the due diligence of CloudNativePG, which applied for Incubation. The project has substantially completed the criteria that show its maturity at the Incubation level, pending the adopter interviews that are the current live stage of this DD.

The following criteria implementations are noteworthy to call out: a mature engineering posture for a Sandbox project (public roadmap, documented architecture, a multi-version release process, signed artifacts, SBOMs, and provenance attestations); strong security fundamentals (a documented vulnerability-reporting path, a named Security Response Team, a completed Security Self-Assessment, and an OpenSSF Best Practices passing badge at 100%); and governance that is explicitly documented as vendor-neutral in direction with a complete and demonstrated maintainer-lifecycle process.

No actions were identified that are blocking at the Incubation level. (An earlier draft review had flagged a missing Security Self-Assessment and single-organization governance concentration as blocking; on review against the actual Incubation criteria, the SSA is complete and the "maintainers from ≥2 organizations" / organization-balance-mechanism bars are graduation-stage requirements, not incubation ones — see the Governance note below.)

The following recommendations are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability: provide the General Technical Review (GTR) and Governance Review links (or have the sponsors produce the GTR as part of this DD), replace the remaining placeholder dates/links in the application issue, and continue the tracked work toward an organization-balance mechanism ahead of any future graduation application.

### Adoption Evaluation

_Pending._ The adopter interviews are the current live stage of this due diligence. The project publishes a [public adopters list](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/ADOPTERS.md) spanning cloud, finance, retail, telecom, public sector, and technology verticals. Per the DD process the sponsors will interview 5–7 adopters and record adopter-approved summaries in the `#### Adoption` section below before this DD advances to public comment.

### Final Assessment

_In progress._ The TOC's evaluation found no blocking items against the Incubation criteria. The DD can proceed through adopter interviews → internal TOC review → public comment once the GTR and Governance Review inputs are in place and the application-issue placeholders are resolved. The final assessment will be recorded after adopter interviews are complete.

## Application Process Principles

### Suggested

- [ ] **Engage with domain-specific TAG(s) to present the technical architecture of the project.**
  - To be confirmed with the project during the DD kickoff; the project has an active public community presence (meetings, Slack, mailing list).

<!-- (Project assertion goes here) -->

### Required

- [ ] **Complete a [General Technical Review (GTR)](../project-reviews/general-technical-questions.md).**
  - Not yet linked. The application issue #1961 leaves the GTR link as a `$LINK` placeholder. Recommend the project file a Technical Review request (TOC `tech-review.yml`), or the sponsors complete the GTR as part of this DD and commit the snapshot at `projects/cloudnative-pg/tech-review/YYYY-MM-DD.md`.

<!-- (Project assertion goes here) -->

- [ ] **Complete a [Governance Review](../project-reviews/governance-review-template.md).**
  - Not yet linked. No governance-review artifact is referenced in the application issue. Recommend the project file a Governance Review request (TOC `gov-review.yml`).

<!-- (Project assertion goes here) -->

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

Project metadata and resources are hosted under the vendor-neutral `cloudnative-pg` GitHub organization and in the CNCF landscape. No vendor-neutrality issues were found in code or documentation during review.

<!-- (TOC Evaluation goes here) -->

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
- Acknowledged as part of the project's Sandbox onboarding and this Incubation application. The exact Sandbox-acceptance date should be confirmed against the official vote record (the application body still carries a `DD-MMM-YYYY` placeholder).

<!-- (TOC Evaluation goes here) -->

- [ ] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presentation for public comment satisfies the Due Diligence Review criteria. This document is that review, in progress.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The project maintains extensive [documentation](https://cloudnative-pg.io/documentation/) covering installation, configuration, operations, and end-user guides.

<!-- (TOC Evaluation goes here) -->

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review if completed as a suggested item prior to application.

### Suggested

- [ ] **Complete a Governance Review**

Not completed/linked as of this review — see the Required Governance Review item above.

<!-- (TOC Evaluation goes here) -->

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The project moved governance into a dedicated [`cloudnative-pg/governance`](https://github.com/cloudnative-pg/governance) repository and restructured into a Steering Committee plus subproject committees; the git history of that repo demonstrates iteration.

<!-- (TOC Evaluation goes here) -->

- [x] **Clear and discoverable project governance documentation.**

Governance is documented in [GOVERNANCE.md](https://github.com/cloudnative-pg/governance/blob/main/GOVERNANCE.md) in the dedicated governance repository.

<!-- (TOC Evaluation goes here) -->

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

The governance repo documents the current Steering Committee, Security Response Team, Infrastructure Team, and subproject committees; the roster matches the organization's team structure. The affiliation-update and election cadence should be confirmed with the project during DD.

<!-- (TOC Evaluation goes here) -->

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

The governance documentation states the project's commitment to vendor-neutral direction.

<!-- (TOC Evaluation goes here) -->

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

[GOVERNANCE.md](https://github.com/cloudnative-pg/governance/blob/main/GOVERNANCE.md) documents the Steering Committee decision process, voting, and amendment procedures.

<!-- (TOC Evaluation goes here) -->

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

Role assignment and lifecycle for the Steering Committee, Security Response Team, Infrastructure Team, and the four subproject maintainer committees are documented in the governance repo and the [contributor ladder](https://github.com/cloudnative-pg/governance/blob/main/CONTRIBUTOR_LADDER.md).

<!-- (TOC Evaluation goes here) -->

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

The [contributor ladder](https://github.com/cloudnative-pg/governance/blob/main/CONTRIBUTOR_LADDER.md) and [GOVERNANCE.md](https://github.com/cloudnative-pg/governance/blob/main/GOVERNANCE.md) define roles, onboarding, offboarding, and emeritus status.

<!-- (TOC Evaluation goes here) -->

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

The project restructured into a Steering Committee and subproject committees and documents the resulting roster. [MAINTAINERS.md](https://github.com/cloudnative-pg/governance/blob/main/MAINTAINERS.md) describes the current five-person seeding as "the starting point, not the intended end state," demonstrating active lifecycle management.

<!-- (TOC Evaluation goes here) -->

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

The governance repo documents four subproject maintainer committees with leadership and lifecycle.

<!-- (TOC Evaluation goes here) -->

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

The complete maintainer list, with names, domains of responsibility, and affiliation, is in [MAINTAINERS.md](https://github.com/cloudnative-pg/governance/blob/main/MAINTAINERS.md).

> **Governance concentration — TOC note (not an Incubation blocker).**
> As of this review, every maintainer across all governing bodies (Steering Committee, Security Response Team, Infrastructure Team, and all four subproject committees) is affiliated with EDB — currently 100% single-organization. Under the **Incubation** criteria this is **not a blocking item**: incubation requires vendor-neutral *direction*, affiliation transparency, and a working maintainer lifecycle, all of which are satisfied. The "maintainers from at least 2 organizations" and "organization-balance mechanism for governance decisions" criteria are **graduation-stage** Required items, not incubation ones.
> The project is already tracking this work ahead of graduation:
> - [cloudnative-pg/governance#68](https://github.com/cloudnative-pg/governance/issues/68) — implement organization-balanced voting (open)
> - [cloudnative-pg/governance#69](https://github.com/cloudnative-pg/governance/issues/69) — formalize a non-EDB maintainer pathway (open)
> - [cloudnative-pg/governance#81](https://github.com/cloudnative-pg/governance/pull/81) — draft org-balanced voting mechanism (open PR: Steering Committee capped at 2 seats per organization, per-subproject representation plus elected Community Representative seats, staggered terms, ranked-choice elections)
>
> Recorded as a positive graduation-readiness signal.

<!-- (TOC Evaluation goes here) -->

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

The project has an appropriate number of active maintainers for its scope: a five-member Steering Committee plus four subproject maintainer committees and a documented Component Owner ladder.

<!-- (TOC Evaluation goes here) -->

- [x] **Code and Doc ownership in GitHub and elsewhere matches documented governance roles.**

GitHub team/merge-approval structure aligns with the documented governance roles; to be spot-checked against OWNERS/CODEOWNERS during DD.

<!-- (TOC Evaluation goes here) -->

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

The project adopts a Code of Conduct based on the CNCF CoC; path to be confirmed in the primary repo during DD.

<!-- (TOC Evaluation goes here) -->

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

CoC references appear in the governance/contributing documentation; cross-link completeness to be confirmed during DD.

<!-- (TOC Evaluation goes here) -->

- [x] **All subprojects, if any, are listed.**

The four subproject committees are listed in the governance repo.

<!-- (TOC Evaluation goes here) -->

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

The project documents a [contributor ladder](https://github.com/cloudnative-pg/governance/blob/main/CONTRIBUTOR_LADDER.md) with multiple roles.

<!-- (TOC Evaluation goes here) -->

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

[CONTRIBUTING.md](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/CONTRIBUTING.md) documents how to submit and manage issues and changes.

<!-- (TOC Evaluation goes here) -->

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

Public communication channels (Slack, mailing list, community meetings) are documented on the [project website](https://cloudnative-pg.io/) and in the repository.

<!-- (TOC Evaluation goes here) -->

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.). List any non-public communications channels and what their special purpose is.**

Public channels are documented; the Security Response path (`security@cloudnative-pg.io`) is the designated non-public channel for vulnerability reports.

<!-- (TOC Evaluation goes here) -->

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

The project runs regular public community meetings; CNCF calendar integration to be confirmed during DD.

<!-- (TOC Evaluation goes here) -->

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

Contribution documentation is maintained in [CONTRIBUTING.md](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/CONTRIBUTING.md) and on the project website.

<!-- (TOC Evaluation goes here) -->

- [x] **Demonstrate contributor activity and recruitment.**

The project shows ongoing contribution and release activity. Recommend corroborating organizational contribution patterns via [LFX Insights](https://insights.linuxfoundation.org/) during DD.

<!-- (TOC Evaluation goes here) -->

## Engineering Principles

### Suggested

- [x] **Roadmap change process is documented.**

The project maintains a public [ROADMAP.md](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/ROADMAP.md).

<!-- (TOC Evaluation goes here) -->

- [x] **History of regular, quality releases.**

The project has a documented multi-version release process and a history of regular [releases](https://github.com/cloudnative-pg/cloudnative-pg/releases).

<!-- (TOC Evaluation goes here) -->

### Required

- [x] **Document project goals and objectives that illustrate the project's differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  - _If applicable_ a General Technical Review was not yet linked (`$LINK` placeholder in #1961). The [project documentation](https://cloudnative-pg.io/documentation/) describes CloudNativePG as a Kubernetes operator for managing PostgreSQL workloads, with its goals and differentiation documented.

<!-- (TOC Evaluation goes here) -->

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

The project documentation covers its purpose and cloud native use cases for running PostgreSQL on Kubernetes.

<!-- (TOC Evaluation goes here) -->

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

A public [ROADMAP.md](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/ROADMAP.md) is maintained.

<!-- (TOC Evaluation goes here) -->

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**
  - _If applicable_ a General Technical Review was not yet linked. Architecture and software design are documented in the [project documentation](https://cloudnative-pg.io/documentation/).

<!-- (TOC Evaluation goes here) -->

- [x] **Document the project's release process.**

The release process is documented, including a multi-version release procedure with signed artifacts, SBOMs, and provenance.

<!-- (TOC Evaluation goes here) -->

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance if completed as a suggested item prior to application.

### Suggested

- [ ] **Complete a joint security assessment with TAG Security and Compliance**

Not completed. SSA-only is acceptable at the Incubation level; a joint assessment with TAG Security and Compliance is a candidate for the graduation track.

<!-- (TOC Evaluation goes here) -->

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

The [security policy](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/SECURITY.md) documents how to report security issues via `security@cloudnative-pg.io`.

<!-- (TOC Evaluation goes here) -->

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

Access control follows the GitHub organization model with documented maintainer/merge-approval roles. 2FA enforcement to be confirmed during DD (TOC recommendation if undocumented).

<!-- (TOC Evaluation goes here) -->

- [x] **Document assignment of security response roles and how reports are handled.**

The governance documentation defines a named Security Response Team and the handling process for reports.

<!-- (TOC Evaluation goes here) -->

- [x] **Document Security Self-Assessment.**

The project has completed a Security Self-Assessment, committed at [`.github/security-self-assessment.md`](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/.github/security-self-assessment.md) and referenced from the project's `SECURITY-INSIGHTS.yml`. Verified present during this review.

<!-- (TOC Evaluation goes here) -->

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The project holds a [passing OpenSSF Best Practices badge](https://www.bestpractices.dev/en/projects/9933) at 100% (187% tiered). Verified during this review.

<!-- (TOC Evaluation goes here) -->

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

A public [adopters list](https://github.com/cloudnative-pg/cloudnative-pg/blob/main/ADOPTERS.md) is maintained.

<!-- (TOC Evaluation goes here) -->

- [ ] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

To be validated during the pending adopter interviews.

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [ ] **TOC verification of adopters.**

Pending — refer to the Adoption portion of this document (adopter interviews are the live stage).

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The project documents its integration with Kubernetes and related ecosystem components in the [project documentation](https://cloudnative-pg.io/documentation/).

<!-- (TOC Evaluation goes here) -->

#### Adoption

_Pending adopter interviews. Adopter-approved summaries will be recorded here (anonymized to industry vertical where required) once the sponsors complete 5–7 interviews._

##### Adopter 1 - $COMPANY/$INDUSTRY

_To be completed after interview._

##### Adopter 2 - $COMPANY/$INDUSTRY

_To be completed after interview._

##### Adopter 3 - $COMPANY/$INDUSTRY

_To be completed after interview._

<!-- Specification Project Considerations section intentionally omitted: CloudNativePG is a Kubernetes operator, not a specification project defining normative behavior for independent implementation. -->
