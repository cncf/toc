# Open Cluster Management (OCM) Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1884)

## Incubation Evaluation Summary for Open Cluster Management (OCM)

### Criteria Evaluation

Kevin Wang conducted the due diligence of Open Cluster Management (OCM), which applied for Incubation, with adopter interviews conducted together with Joseph Sandoval and Ricardo Aravena. The project has completed the Required criteria that demonstrate its maturity at Incubation.

The following criteria implementations are noteworthy to call out:

- **Governance and Maintainers**: project metadata and infrastructure are hosted on vendor-neutral infrastructure, with maintainers from multiple organizations.
- **Engineering Principles**: the project maintains a [public roadmap board](https://github.com/orgs/open-cluster-management-io/projects/2/views/9) and an active [enhancements repository](https://github.com/open-cluster-management-io/enhancements) for feature proposals; release cadence follows a documented minor/patch schedule in `community/RELEASE.md`.
- **Security**: both `ocm/SECURITY.md` and the project website direct users to `community/SECURITY.md` for vulnerability reporting guidance.
- **Ecosystem**: [ocm/ADOPTERS.md](https://github.com/open-cluster-management-io/ocm/blob/main/ADOPTERS.md) documents 10 adopter organizations and 18 CNCF/non-CNCF integrations; the TOC completed adopter interviews with seven organizations, three of which have no maintainer affiliation with the project.

No blocking issues were identified during this due diligence review. The project's governance, security, and engineering practices demonstrate a level of maturity appropriate for Incubation.

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

- **Governance and Maintainers**: The maintainer roster already spans multiple organizations. Documenting how governance safeguards vendor neutrality as the project grows would build on this foundation. Updating the subproject inventory to reflect accepted additions and repository moves, with leadership and maturity status, would make project responsibilities clearer, as would alignment between `MAINTAINERS.md` and subproject OWNERS lists (starting with `pmorie`, `suigh`, and `tesshuflower`).
- **Contributors and Community**: The contributor ladder still contains a `[TODO: appropriate election]` placeholder for Organization Member voting rights. Clarifying those rights and how the two `CONTRIBUTING.md` files reference or supersede each other would give contributors more consistent guidance. Linking the existing LFX meeting calendar from `community/COMMUNICATION.md` would make meetings easier to find.
- **Engineering Principles**: The roadmap change process is described in the GTR. Maintaining that guidance in a governance or community document, with an explanation of how the enhancements repository relates to the roadmap board, would make the process easier to follow.
- **Security**: An explicit two-factor authentication requirement in `GOVERNANCE.md` or `CONTRIBUTOR_LADDER.md` would make access expectations clear. Naming Core Maintainers as the security response role in `SECURITY.md`, consistent with the contributor ladder, and incorporating the report-handling procedure from `SELF_ASSESSMENT.md` into maintained security guidance would clarify responsibilities.


### Adoption Evaluation

The adopter interviews show OCM in use at a level appropriate for Incubation. The TOC interviewed seven organizations, including three independent adopters. Two of these, a professional services organization and a fintech company, are kept anonymous in this report and use OCM in production. Salesforce uses it in pre-production for two tier-zero services. Four additional interviews were conducted with maintainer-affiliated adopters.

Adopters use OCM to distribute workloads and shared services across clusters, with placement flexibility and its extensible add-on framework among the reasons for choosing it. The professional services organization reported reducing vulnerability remediation time from weeks to days through its combined GitOps and OCM rollout workflow. Feedback also identified areas for improvement, including cross-cluster visibility, rollback support for ManifestWorkReplicaSet, and add-on documentation. These reflect operational needs and opportunities for further development, rather than blockers to Incubation.

### Final Assessment

The TOC has found Open Cluster Management to have satisfied the criteria for Incubation and recommends moving the project to Incubation. No blocking issues were identified during this due diligence. The recommendations in this report are non-blocking.

## Application Process Principles

### Suggested

- [x] **Engage with domain-specific TAG(s) to present the technical architecture of the project.**

  The project presented at the CNCF TAG Runtime meeting on 06-Mar-2025: <https://www.youtube.com/watch?v=ex3vdlCDPXw>.

### Required

- [x] **Complete a [General Technical Review (GTR)](../toc_subprojects/project-reviews-subproject/general-technical-questions.md).**

<!-- (TOC Evaluation goes here) -->
The project maintainers have published a [self-authored General Technical Review](https://github.com/open-cluster-management-io/community/blob/main/cncf/GTR.md), which is under review by the Project Reviews subproject ([cncf/toc#2173](https://github.com/cncf/toc/issues/2173)). The TOC reviewers suggest completing the review after this due diligence. This follow-up is non-blocking for Incubation.

- [x] **Complete a [Governance Review](../toc_subprojects/project-reviews-subproject/governance-review-template.md).**

<!-- (TOC Evaluation goes here) -->
The project maintainers have published a governance self-assessment, which is under review by the Project Reviews subproject ([cncf/toc#2263](https://github.com/cncf/toc/issues/2263)). The TOC reviewers suggest completing the review after this due diligence. This follow-up is non-blocking for Incubation.

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

<!-- (TOC Evaluation goes here) -->
Yes. OCM's project metadata and resources are hosted on vendor-neutral infrastructure:
- GitHub organization: <https://github.com/open-cluster-management-io>
- Project domain: <https://open-cluster-management.io/>
- Governance: <https://github.com/open-cluster-management-io/community/blob/main/GOVERNANCE.md>

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
- Met during Project's application on 28-SEP-2023.

<!-- (TOC Evaluation goes here) -->
Yes. Acknowledged in the project's application, with continued activity on the onboarding issue <https://github.com/cncf/sandbox/issues/227>. Kevin Wang (TOC Reviewer) also met in person with OCM project contacts at KubeCon + CloudNativeCon Japan (July 2026) to discuss this due diligence.

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

<!-- (TOC Evaluation goes here) -->
Refer to the following links:
- Installation: <https://open-cluster-management.io/docs/getting-started/installation/>
- Quick start: <https://open-cluster-management.io/docs/getting-started/quick-start/>
- Architecture / concepts: <https://open-cluster-management.io/docs/concepts/architecture/>
- User scenarios: <https://open-cluster-management.io/docs/scenarios/>
- Blog: <https://open-cluster-management.io/blog/>

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

<!-- (TOC Evaluation goes here) -->
The OCM project has been continuously updating its governance doc, some examples are:
- Update the contributor ladder: <https://github.com/open-cluster-management-io/community/pull/187>
- Update the community meeting time with office hours and update the communication page: <https://github.com/open-cluster-management-io/community/pull/188>
- Add `GOVERNANCE.md` for OCM incubation checks: <https://github.com/open-cluster-management-io/community/pull/205>
- Refactor governance to define a federated core/subproject structure, a maintainer council, voting thresholds, subproject lifecycle, communications, and an amendment process: <https://github.com/open-cluster-management-io/community/pull/228>

- [x] **Clear and discoverable project governance documentation.**

<!-- (TOC Evaluation goes here) -->
The [project governance doc](https://github.com/open-cluster-management-io/community/blob/main/GOVERNANCE.md) is public and directly discoverable, with a table of contents covering project structure, maintainer selection/removal, decision making, subprojects, security, release process, Code of Conduct, communications, and amendments.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

<!-- (TOC Evaluation goes here) -->
The governance documentation is kept up to date with the current project state, including maintainer roster changes and structural refactors.

Some examples are:

- Refactored governance to define a federated core/subproject structure, a maintainer council, voting thresholds, subproject lifecycle, communications, and an amendment process: <https://github.com/open-cluster-management-io/community/pull/228>
- Added Meng Yan as addon-contrib maintainer: <https://github.com/open-cluster-management-io/community/pull/232>
- Added Mike Ng as api maintainer: <https://github.com/open-cluster-management-io/community/pull/234>

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

<!-- (TOC Evaluation goes here) -->
Governance states that ideas and contributions are accepted on technical merit, with maintainer nominations evaluated without prejudice to employer or demographics <https://github.com/open-cluster-management-io/community/blob/main/GOVERNANCE.md>, and the README separately describes OCM as vendor-neutral and standards-based <https://github.com/open-cluster-management-io/ocm/blob/main/README.md>. However, governance does not state an explicit vendor-neutrality rule, company-seat limit, or conflict-of-interest procedure. The [maintainer roster](https://github.com/open-cluster-management-io/community/blob/main/MAINTAINERS.md) lists 25 active maintainers from eight organizations, including 13 from Red Hat.

TOC reviewer suggests explicitly documenting the basis on which vendor neutrality is enforced in governance (e.g., a maintainer-seat limit per organization or a conflict-of-interest procedure), so that neutrality is safeguarded as the project and its organizational makeup continue to evolve; the maintainer roster already spans multiple organizations, which is a positive sign for neutrality in practice.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

<!-- (TOC Evaluation goes here) -->
Governance explicitly defines consensus and formal voting, simple-majority and two-thirds thresholds, maintainer appointment/removal, core-component and subproject addition/removal, community comment periods, CNCF representation, and a governance amendment process <https://github.com/open-cluster-management-io/community/blob/main/GOVERNANCE.md>.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

<!-- (TOC Evaluation goes here) -->
[CONTRIBUTOR_LADDER.md](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTOR_LADDER.md) defines onboarding, eligibility, nomination, approval, and removal for contributor, organization-member, reviewer, maintainer, and core-maintainer roles.

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

<!-- (TOC Evaluation goes here) -->
[CONTRIBUTOR_LADDER.md](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTOR_LADDER.md) clearly documents progression through reviewer and maintainer roles, nomination and majority approval, inactivity thresholds, involuntary removal/demotion, stepping down, and emeritus status, and governance separately defines becoming, resigning, removing, and converting core maintainers to emeritus status.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

<!-- (TOC Evaluation goes here) -->
Updates to the MAINTAINERS.md file can be found at: <https://github.com/open-cluster-management-io/community/commits/main/MAINTAINERS.md>

Some examples are:
- Add new Guidewire maintainers: <https://github.com/open-cluster-management-io/community/pull/210>
- Add Ben Perry and Tyler Gillson as maintainers: <https://github.com/open-cluster-management-io/community/pull/214>
- Move maintainers to emeritus status: <https://github.com/open-cluster-management-io/community/pull/209>
- Add Meng Yan as addon-contrib maintainer: <https://github.com/open-cluster-management-io/community/pull/232>
- Add Mike Ng as api maintainer: <https://github.com/open-cluster-management-io/community/pull/234>

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

<!-- (TOC Evaluation goes here) -->
The org-level [GOVERNANCE.md](https://github.com/open-cluster-management-io/community/blob/main/GOVERNANCE.md) lists subprojects (Addon-Contrib, Argo Integration, Cluster Proxy Addon, Policy Addon, Managed Service Account Addon, Lab, and an open-ended "and others as they are added to the ecosystem"), and documents subproject decision making, contributor management, escalation, and add/remove processes. Still missing: per-subproject maturity status, per-subproject leadership/maintainer lists, and a working link for the Addon-Contrib subproject — the listed repository URL contains a typo (`gopen-cluster-management-io`) and 404s.

TOC reviewer suggests publishing a canonical, complete subproject inventory with leadership and maturity/status per subproject, and fixing the broken Addon-Contrib link.

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

<!-- (TOC Evaluation goes here) -->
[community/MAINTAINERS.md](https://github.com/open-cluster-management-io/community/blob/main/MAINTAINERS.md) lists 25 active and 2 emeritus maintainers with names, GitHub handles, public affiliations, and separate responsibility mappings for core areas, SDK/tools, community/docs, and add-ons.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) -->
According to [MAINTAINERS.md](https://github.com/open-cluster-management-io/community/blob/main/MAINTAINERS.md), the project currently has 25 named maintainers from 8 organizations, with responsibility mappings for core areas, SDK/tools, community/docs, and add-ons. [LFX Insights](https://insights.linuxfoundation.org/project/openclustermanagement/contributors?timeRange=past365days) separately reports **214** active contributors across **32** organizations, activity on **322** of the last 365 days, and **467** pull requests opened over the same period (past year, all repos).

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

<!-- (TOC Evaluation goes here) -->
OCM grants approver rights through per-subproject OWNERS files, matching the Maintainer role defined in [CONTRIBUTOR_LADDER.md](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTOR_LADDER.md): "Added to the OWNERs file in one or more Subprojects." Spot checks across nine repositories (`ocm`, `community`, `registration`, `work`, `placement`, `addon-framework`, `api`, `managed-serviceaccount`, `cluster-proxy`) show the mechanism is kept current: the most recently added maintainer, Kahiro Okina, already appears as an approver.

TOC reviewer suggests reconciling MAINTAINERS.md with the OWNERS approver lists across subproject repositories, starting with `pmorie`, `suigh`, and `tesshuflower`, and publishing an all-repository ownership audit; three individuals (`pmorie`, `suigh`, `tesshuflower`) hold OWNERS approver rights in several of these repositories but do not appear in the central [MAINTAINERS.md](https://github.com/open-cluster-management-io/community/blob/main/MAINTAINERS.md) roster.

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

<!-- (TOC Evaluation goes here) -->
The public community Code of Conduct <https://github.com/open-cluster-management-io/community/blob/main/CODE_OF_CONDUCT.md> states that the community abides by the CNCF Code of Conduct and documents confidential reporting and CNCF escalation when a maintainer is involved.

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

<!-- (TOC Evaluation goes here) -->
`GOVERNANCE.md` has a dedicated Code of Conduct section linking directly to the CNCF Code of Conduct, and `CODE_OF_CONDUCT.md` independently repeats that CNCF link <https://github.com/open-cluster-management-io/community/blob/main/GOVERNANCE.md#code-of-conduct>.

- [x] **All subprojects, if any, are listed.**

<!-- (TOC Evaluation goes here) -->
The [governance subproject list](https://github.com/open-cluster-management-io/community/blob/main/GOVERNANCE.md#subprojects) is not up to date. Its six named entries are unchanged since the list was [introduced on 17 July 2025](https://github.com/open-cluster-management-io/community/commit/251aa14f85ed15740ee91eedc8bb7e1373c1df41). The community subsequently approved moving Dynamic Scoring Framework out of Addon-Contrib into a standalone repository ([open-cluster-management-io/community#236](https://github.com/open-cluster-management-io/community/issues/236)); that repository was created on 6 July 2026, and the [migration PR](https://github.com/open-cluster-management-io/addon-contrib/pull/125) merged on 15 August 2026, but the governance list does not name it.

The TOC reviewers suggest updating the list to reflect accepted subprojects and repository moves, recording when it was last updated, and using issue templates for subproject additions and removals that include a corresponding list update. This is a non-blocking documentation follow-up for Incubation.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) -->
[community/CONTRIBUTOR_LADDER.md](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTOR_LADDER.md) documents a five-tier ladder (Contributor → Organization Member → Reviewer → Maintainer → Core Maintainer), each with explicit responsibilities, requirements, and privileges, plus a concrete PR-based nomination/approval mechanism; it has been substantively revised twice since its 2022-12-02 introduction (PR [#187](https://github.com/open-cluster-management-io/community/pull/187), 2024-08-01; PR [#222](https://github.com/open-cluster-management-io/community/pull/222), 2025-08-11).

TOC reviewer suggests resolving the Organization Member tier's outstanding `[TODO: appropriate election]` voting-rights placeholder in the live document.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) -->
The process for submitting issues and changes is clearly defined in [ocm/CONTRIBUTING.md](https://github.com/open-cluster-management-io/ocm/blob/main/CONTRIBUTING.md), linked directly from the primary repository's README. It covers the full patch workflow, DCO sign-off, local dev-environment setup, and repo-specific detail like the API/CRD change walkthrough.

TOC reviewer suggests clarifying whether `ocm/CONTRIBUTING.md` or [community/CONTRIBUTING.md](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTING.md) should reference or supersede the other for the ground they cover in common (Terms, DCO, Code of Conduct, patch workflow, issue management), so the content has one source of truth instead of two copies that can drift apart.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) -->
Public communication channels are documented in [ocm/README.md#community](https://github.com/open-cluster-management-io/ocm/blob/main/README.md#community): a Slack channel (`#open-cluster-mgmt` on the Kubernetes Slack, joinable via a public, unauthenticated self-service invite per [community/COMMUNICATION.md](https://github.com/open-cluster-management-io/community/blob/main/COMMUNICATION.md)), a Google Group mailing list, a community-meeting calendar, and a YouTube channel.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) -->
[community/COMMUNICATION.md](https://github.com/open-cluster-management-io/community/blob/main/COMMUNICATION.md) itemizes each channel with its purpose (Slack, mailing list, weekly community meetings/office hours, a WeChat group joined by request), consistent with `ocm/README.md`'s independent listing, and `GOVERNANCE.md` separately lists the project's non-public channels and their special purposes: security reports to `OCM-security@googlegroups.com` and Code of Conduct reports to `cncf-ocm-maintainers@lists.cncf.io`.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) -->
OCM documents weekly community meetings and office hours in [community/COMMUNICATION.md](https://github.com/open-cluster-management-io/community/blob/main/COMMUNICATION.md), with a public Google Calendar. The community also completed CNCF calendar onboarding on 15 September 2025, with a maintainer confirming the [LFX meeting calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/openclustermanagement?view=month) in [open-cluster-management-io/community#224](https://github.com/open-cluster-management-io/community/issues/224).

TOC reviewer suggests linking the existing LFX meeting calendar from `community/COMMUNICATION.md` to improve discoverability.

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) -->
[ocm/CONTRIBUTING.md](https://github.com/open-cluster-management-io/ocm/blob/main/CONTRIBUTING.md) covers how to contribute, and it has grown since the repository's 2021 creation, each revision adding detail rather than replacing it: recent additions include a [dev-environment guide](https://github.com/open-cluster-management-io/ocm/blob/main/CONTRIBUTING.md#setting-up-your-dev-environment-coding-and-debugging) and an AI-contribution policy ([#1547](https://github.com/open-cluster-management-io/ocm/pull/1547), 2026-08-04). The org-wide sibling, `community/CONTRIBUTING.md`, stays a short, generic notice, but day-to-day contribution activity follows the document that actually grew.

- [x] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) -->
[LFX Insights](https://insights.linuxfoundation.org/project/openclustermanagement/contributors?timeRange=past365days) (past year, all repos): **214** active contributors across **32** organizations; activity on **322** of the last 365 days; **467** pull requests opened. OCM recruits contributors through sustained KubeCon Pavilion presence in North America, China, Japan, and Europe, plus recent talks such as [ArgoCon EU 2025 - Scale GitOps With the Argo CD Agent and Open Cluster Management](https://www.youtube.com/watch?v=trR8OcmizXQ) and [KubeCon EU 2025 - How the SIG-Multicluster API Specifications Are Used for Real World](https://www.youtube.com/watch?v=I9GV4N23dvE).

## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

<!-- (TOC Evaluation goes here) -->
According to the [General Technical Review](https://github.com/open-cluster-management-io/community/blob/main/cncf/GTR.md), enhancement issues enter a backlog, scope and priority are reviewed in bi-weekly community meetings, and changes affecting APIs or multiple components require a proposal in the `enhancements` repository. No equivalent description was found elsewhere, and the GTR itself is a point-in-time assessment document submitted for TAG review, not a maintained source of truth for community governance rules.

TOC reviewer suggests documenting the roadmap change process in an authoritative, maintained location (e.g. `GOVERNANCE.md` or the `community` repository), rather than only in the GTR.

- [x] **History of regular, quality releases.**

<!-- (TOC Evaluation goes here) -->
OCM follows the cadence set out in [community/RELEASE.md](https://github.com/open-cluster-management-io/community/blob/main/RELEASE.md): minor releases roughly every three to four months, with patch releases as needed in between. Over the past year this produced three minor releases plus several patch releases. [Release notes](https://open-cluster-management.io/docs/release/) record dated changes with links to the relevant issues and pull requests.

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**

<!-- (TOC Evaluation goes here) -->
[ocm/README.md#project-goals-and-objectives](https://github.com/open-cluster-management-io/ocm/blob/main/README.md#project-goals-and-objectives) and [ocm/README.md#differentiation-in-the-cloud-native-landscape](https://github.com/open-cluster-management-io/ocm/blob/main/README.md#differentiation-in-the-cloud-native-landscape) describe OCM's goals and differentiation: vendor-neutral multicluster APIs, an implementation of the SIG-Multicluster specification, a common `PlacementDecision` API, and an extensible add-on framework. The [General Technical Review](https://github.com/open-cluster-management-io/community/blob/main/cncf/GTR.md) also defines the target users and the project's in-scope and out-of-scope responsibilities.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

<!-- (TOC Evaluation goes here) -->
[ocm/README.md#what-ocm-does](https://github.com/open-cluster-management-io/ocm/blob/main/README.md#what-ocm-does) explains OCM's hub-spoke, declarative approach to multicluster management. The [General Technical Review scope](https://github.com/open-cluster-management-io/community/blob/main/cncf/GTR.md#scope) defines its target users, primary use cases, and the boundary that cluster provisioning and lifecycle remain out of scope. The project website documents 11 practical [user scenarios](https://open-cluster-management.io/docs/scenarios/), including workload deployment, Argo CD integration, multiple hubs, and Cluster API registration.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

<!-- (TOC Evaluation goes here) -->
According to community [README](https://github.com/open-cluster-management-io/community/blob/main/README.md), the project uses [GitHub Projects board](https://github.com/orgs/open-cluster-management-io/projects/2/views/9) as the public roadmap.

The project also uses [enhancements repository](https://github.com/open-cluster-management-io/enhancements) for proposing, discussing, and tracking significant new features and improvements to Open Cluster Management.

The TOC reviewers suggest to clarify how the enhancements repository relates to, and feeds into, the roadmap board.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**

  <!-- (TOC Evaluation goes here) -->
  The [architecture overview](https://open-cluster-management.io/docs/concepts/architecture/) describes OCM's hub-spoke model: the hub hosts the multi-cluster control plane, while agents on managed clusters pull and reconcile the desired configuration. It explains cluster registration and its double opt-in security model, work distribution through `ManifestWork`, placement decisions, and the add-on framework for extending these capabilities.

- [x] **Document the project's release process.**

<!-- (TOC Evaluation goes here) -->
[community/RELEASE.md](https://github.com/open-cluster-management-io/community/blob/main/RELEASE.md) documents the versioning scheme, the cadence for major, minor, and patch releases, and a three-stage planning, development, and release process.

## Security

### Suggested

N/A

### Required

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance.

- [x] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) -->
The security issue reporting process is documented in [community/SECURITY.md](https://github.com/open-cluster-management-io/community/blob/main/SECURITY.md). Per that document, reports go to `OCM-security@googlegroups.com`, are acknowledged and analyzed within five working days, and the reporter is kept updated through triage, fix, and release, following the [Kubernetes security reporting process](https://kubernetes.io/docs/reference/issues-security/security/).

Both [ocm/SECURITY.md](https://github.com/open-cluster-management-io/ocm/blob/main/SECURITY.md) and the [website security page](https://open-cluster-management.io/docs/security/) point directly to this same community file rather than duplicating the content, keeping a single source of truth.

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) -->
According to [community/CONTRIBUTOR_LADDER.md](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTOR_LADDER.md), being listed as an Approver or Reviewer in an [OWNERS](https://github.com/open-cluster-management-io/ocm/blob/main/OWNERS) file grants "GitHub or CI/CD rights to approve pull requests in specific directories," and Prow checks enforce those rights on every pull request; the OWNERS files themselves are only the resulting name lists, not the rule. No governance or membership document defines an explicit two-factor authentication requirement for organization members, and GitHub's org-level 2FA enforcement setting is not publicly visible to confirm whether it is enforced.

TOC reviewer suggests stating an explicit two-factor authentication requirement for GitHub organization members in `GOVERNANCE.md` or `CONTRIBUTOR_LADDER.md`.

- [x] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) -->
According to [community/SECURITY.md](https://github.com/open-cluster-management-io/community/blob/main/SECURITY.md), reported vulnerabilities are handled by "OCM project owners," who acknowledge and analyze each report, work with the reporter in confidence, and coordinate a fix through to release.

However, no role named "project owner" appears in [community/CONTRIBUTOR_LADDER.md](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTOR_LADDER.md); the closest documented match is the Core Maintainer responsibility of "Responding to security compromise reports according to OCM security procedures."

[ocm/SELF_ASSESSMENT.md#security-issue-resolution](https://github.com/open-cluster-management-io/ocm/blob/main/SELF_ASSESSMENT.md#security-issue-resolution) adds that a maintainer works directly with the reporter, the maintainer team assesses validity and severity, and patches are issued for the three most recent minor releases; however, this self-assessment is itself a point-in-time security review document, not a maintained source of truth for community role assignments.

TOC reviewer suggests reconciling the role name in `SECURITY.md` with `CONTRIBUTOR_LADDER.md` by naming Core Maintainers as the role responsible for handling reports, and recording the report-handling procedure described in `SELF_ASSESSMENT.md` in an authoritative, maintained location instead.

- [x] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) -->
[ocm/SELF_ASSESSMENT.md](https://github.com/open-cluster-management-io/ocm/blob/main/SELF_ASSESSMENT.md) covers the project's actors, security goals and non-goals, security features, compliance, and incident response, with the [Deployment Pipeline](https://github.com/open-cluster-management-io/ocm/blob/main/SELF_ASSESSMENT.md#deployment-pipeline) section documenting PR sign-off, code-review, and approval requirements. Its "Security functions and features" table also flags a `work executor subject` feature as "disabled by default, should consider enabling it by default in the future." In addition, independent CI workflows corroborate those pipeline controls:
- [.github/workflows/dco.yml](https://github.com/open-cluster-management-io/ocm/blob/main/.github/workflows/dco.yml) enforces DCO sign-off on every pull request.
- [.github/workflows/dependency-review.yml](https://github.com/open-cluster-management-io/ocm/blob/main/.github/workflows/dependency-review.yml) scans dependency manifests on every pull request and can block merges that introduce known-vulnerable packages.
- [.github/workflows/ossf-scorecard.yml](https://github.com/open-cluster-management-io/ocm/blob/main/.github/workflows/ossf-scorecard.yml) runs OpenSSF Scorecard supply-chain checks on pushes to `main` and on a weekly schedule.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) -->
OCM holds an [OpenSSF Best Practices passing badge](https://www.bestpractices.dev/projects/5376).

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

<!-- (TOC Evaluation goes here) -->
[ocm/ADOPTERS.md](https://github.com/open-cluster-management-io/ocm/blob/main/ADOPTERS.md) lists 12 entries in its "End users" table: 10 are typed as adopter organizations (e.g., Alibaba Cloud, Ant Group, Red Hat, eBay, Guidewire Software), and 2 are typed as "Integration" (RamenDR, KubeStellar Console) rather than adopters. The table has no dedicated adoption-level column, but most use-case descriptions convey the level directly, e.g. Red Hat's entry states RHACM "is used in multiple production installations," and eBay's states it "manage[s] application deployment across 190+ clusters."

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

<!-- (TOC Evaluation goes here) -->
[community/MAINTAINERS.md](https://github.com/open-cluster-management-io/community/blob/main/MAINTAINERS.md) shows at least five organizations listed in ADOPTERS.md without a current maintainer seat: Alibaba Cloud, Ant Group, Xiao Hong Shu, Groww, and VESSL AI. That satisfies the independent-adopter count expected for incubation, based on the public adopter statements rather than direct TOC interviews.

- [x] **TOC verification of adopters.**

<!-- (TOC Evaluation goes here) -->
The TOC interviewed seven adopting organizations. Three have no OCM maintainers and no role in project governance: a professional services organization, Salesforce, and a fintech company. The Adoption section below contains summaries for these three adopters, with the professional services organization and fintech company kept anonymous. The other four organizations employ OCM maintainers, and in each of those interviews the interviewee was a maintainer.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

<!-- (TOC Evaluation goes here) -->
[ocm/ADOPTERS.md#ecosystem](https://github.com/open-cluster-management-io/ocm/blob/main/ADOPTERS.md#ecosystem) lists 18 CNCF and non-CNCF project integrations, several with their own integration-specific documentation: [Cluster API](https://github.com/open-cluster-management-io/ocm/tree/main/solutions/cluster-api), [KubeVela](https://kubevela.io/docs/platform-engineers/system-operation/working-with-ocm/), [Kueue](https://kueue.sigs.k8s.io/docs/tasks/manage/setup_multikueue/#optional-setup-multikueue-with-open-cluster-management), [Istio](https://github.com/open-cluster-management-io/multicluster-mesh), and [Submariner](https://github.com/stolostron/submariner-addon), etc.

#### Adoption

##### Adopter 1 - Anonymous / Professional services

September 2026. Interviewer: Joseph Sandoval.

A large professional-services organization has worked with OCM for approximately one year and now runs it in production. Forty clusters are enrolled on v1.1.0, and the team continues to bring clusters operated by smaller internal groups into the platform.

The primary driver was supporting multi-cluster deployments without doubling deployment and infrastructure costs. The team initially considered a conventional Argo CD hub-and-spoke model, but that approach would have required applications to run active-active across multiple clusters. OCM Placement allows the platform to deploy an application to any eligible set of clusters without choosing specific clusters in advance.

The organization also created an add-ons project: one definition for security, logging, monitoring, and other baseline services that OCM distributes across all 40 clusters. This gives the team one source of truth for services that must run everywhere.

The production deployment has remained on OCM v1.1.0. During pre-production, the team updated more frequently while working through bugs and missing features with the maintainers. Adoption introduced little friction because there was no previous multi-cluster management system to migrate. The documentation was reliable, and the team found many implementation details directly in the Helm chart configuration and accompanying Markdown files.

Maintainers were responsive through Slack and GitHub. The team also has access to a weekly maintainer sync for larger users, although it has not needed to attend recently.

The clearest measured benefit is faster vulnerability remediation. Moving from Helm wrapped in Terraform to GitOps simplified the release workflow, while OCM automated rollout across the 40 enrolled clusters. Together, these changes reduced remediation time from weeks to days. Most remaining delay comes from required stakeholder notice rather than hands-on deployment work.

The next objective is zero-downtime maintenance across multiple clusters. That capability has not yet been realized because north-south traffic routing remains an external blocker.

The organization has made OCM a critical part of its platform. If the project were archived, the team would need to fork it immediately, although it does not currently have capacity to assume upstream maintainership. It identified limited market awareness as OCM's main weakness, noting that platform vendors often had little familiarity with the project.

The team considers OCM's independence from any single cloud provider a major strength. Its roadmap makes OCM the standard deployment-targeting mechanism across all clusters, with future support for multi-cloud workloads and failover between providers.

Note: This adopter preferred to remain anonymous. The interview details are kept in a private file for CNCF TOC review.

##### Adopter 2 - Salesforce / Enterprise SaaS

September 2026. Interviewer: Joseph Sandoval.

Salesforce began evaluating OCM through a proof of concept in June 2025. It is currently using OCM v1.3 in pre-production across approximately 20 clusters, supporting two tier-zero services, with production targeted by the end of the year. The team intends to remain one minor release behind the latest version.

The core problem was deploying services across multiple clusters while preserving customers' visibility into their manifests and placement decisions. OCM's ManifestWorkReplicaSet and Placement APIs addressed that need directly.

The team also evaluated Karmada and Azure Fleet Manager. Karmada did not fit Salesforce's preferred extension model because it would introduce additional data-store and API-server components, as well as dependencies on Karmada-specific tooling rather than standard Kubernetes clients. Azure Fleet Manager was still too early in its development at the time. OCM was the best fit because of its CRD-based architecture, alignment with standard Kubernetes patterns, and responsive maintainers. GitHub questions received quick replies, and pull requests were often merged within a week.

Adoption was straightforward. Several engineers could run OCM locally, and existing AWS authentication support eased its integration into deployment pipelines. Support for GCP authentication remains a potential gap as Salesforce considers expanding there.

ManifestWorkReplicaSet was still alpha when Salesforce adopted it, so the team encountered several maturity gaps. These included the absence of rollback support, subject access review limited to service accounts rather than users and groups, and the need for internal Argo Rollouts integration. Salesforce built some capabilities internally and has not yet contributed all of them upstream. Cross-cluster visibility was another challenge because most existing tools operate against one cluster at a time. The team spent several months developing workarounds. Core documentation for work distribution and placement was useful, while add-on documentation was less complete.

Because the deployment remains in pre-production, Salesforce does not yet have measurable production outcomes. The immediate benefit is improved resilience for tier-zero services running across multiple clusters. OCM is also integrated with five or six internal tools. If the project were archived, the team said it would need to fork and continue maintaining it.

The team considers OCM's hub-and-spoke RBAC model and separation of controller responsibilities major strengths. Its main gap is the lack of a unified aggregation and visibility layer across managed clusters. The Work API also remains at v1alpha1. Salesforce is considering upstream contributions for rollback, dashboards, and manifest customization. Longer term, it sees OCM as the scheduling engine for its next-generation multi-cluster platform, replacing static deployment pipelines with dynamic, API-driven provisioning.

Note: The interview details are kept in a private file for CNCF TOC review.

##### Adopter 3 - Anonymous / Fintech

The adopter runs OCM in production. It needed a cloud native way to operate across multiple regions and cloud providers, and uses OCM's hub-and-spoke model and manifest-based workflow to replicate configuration and distribute components across clusters. The adopter updates quarterly, driven by the features it needs.

The adopter had hit limits running Argo CD across many clusters, found OCM's concepts easier to understand than Karmada's, and preferred it to a proprietary option it evaluated. The team understood the core model within days. Replicating configuration was straightforward; configuration overrides were harder, so the team built an adjacent tool to fill that gap.

The documentation was strong enough to get the team running quickly, and the team found the source code and enhancement proposals well structured. The team engaged through Slack and GitHub, and maintainers were responsive. The main benefit is a single control plane and consolidated view across clusters. OCM is a deep dependency it would keep maintaining if the project were archived. It wants a UI and better cross-cluster status reporting, and plans to expand its use and stay involved.

Note: This adopter is kept anonymous by default. Additional details may be published in accordance with the adopter's disclosure preferences. The interview details are kept in a private file for CNCF TOC review.

##### Additional interviews: maintainer-affiliated organizations

The TOC reviewers also conducted four interviews with adopters whose organizations employ OCM maintainers.
