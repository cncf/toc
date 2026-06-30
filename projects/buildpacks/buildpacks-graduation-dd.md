# Buildpacks Graduation Due Diligence

- Link to [cncf/toc/issues#1538](https://github.com/cncf/toc/issues/1538)

## Graduation Evaluation Summary for Buildpacks

### Criteria Evaluation

[Faseela K](https://github.com/kfaseela) conducted the due diligence of Cloud Native Buildpacks (CNB) who applied for Graduation. The project has completed the criteria that show its maturity at the level of Graduation.

The following criteria implementations are noteworthy to call out:
- CNB is a specification-first project (Platform API, Buildpack API) with reference implementations (lifecycle, pack), driving adoption across cloud providers, CI/CD platforms, and enterprise developer platforms.
- [General Technical Review](https://github.com/cncf/toc/pull/2047) completed on 20-Feb-2026, covering architecture, cloud-native use cases, differentiation, and ecosystem integrations.
- Third-party security audit (OSTIF/Quarkslab) completed, findings published, and all moderate/low findings tracked to resolution in specific pack and lifecycle milestones.
- Four production adopter interviews (Google, Salesforce, Adopter 3, 7SIGNAL) confirm graduation-level usage across diverse industries at meaningful scale since 2020–2022.
- The spec has 9+ independent platform implementations (Google Cloud, Heroku, kpack, GitLab, DigitalOcean, Spring Boot, Dokku, Fly.io, and more), demonstrating broad adoption and real-world interoperability. The RFC process (341+ RFCs, mature FCP process, supermajority voting) reflects a well-established and community-driven approach to spec evolution.

**Items raised and resolved during review:**
- **Contributor ladder** ([#282](https://github.com/buildpacks/community/issues/282)): Added explicit contributor ladder and role-summary table to [contributors/guide.md](https://github.com/buildpacks/community/blob/main/contributors/guide.md); updated [GOVERNANCE.md](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md) with Team Lead appointment, Security Response Team composition, and 2FA requirement. Resolved via [PR#291](https://github.com/buildpacks/community/pull/291)
- **Security process** ([#284](https://github.com/buildpacks/community/issues/284)): Enhanced [SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md) with response timelines, required report contents, and disclosure guidance. Resolved via [PR#6](https://github.com/buildpacks/.github/pull/6)
- **Code of Conduct** ([#285](https://github.com/buildpacks/community/issues/285)): Org-level CoC updated to explicitly reference the CNCF CoC. Resolved via [PR#5](https://github.com/buildpacks/.github/pull/5)
- **Audit/self-assessment discoverability** ([#286](https://github.com/buildpacks/community/issues/286)): Audit report and self-assessment links added to [SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md). Resolved via [PR#8](https://github.com/buildpacks/.github/pull/8)
- **Meeting links** ([#283](https://github.com/buildpacks/community/issues/283)): Stale recording links updated to current YouTube channel. Resolved via [PR#882](https://github.com/buildpacks/docs/pull/882)
- **Maintainer affiliations** ([#294](https://github.com/buildpacks/community/issues/294)): Affiliation was documented for TOC members but not consistently for Team Leads, Maintainers, and Component Maintainers. Resolved via [PR#301](https://github.com/buildpacks/community/pull/301): all team listings converted to Markdown tables with an explicit Organization column.

**Non-blocking recommendations:**
- **Vendor-neutrality of project direction**: 2 of 5 TOC seats are vacant, and 2 of 3 active members are from Salesforce/Heroku (related companies per CNCF charter). The project has opened [RFC#339](https://github.com/buildpacks/rfcs/pull/339) proposing a Steering Committee model with a 3+2 seat structure and a hard 40% corporate representation cap. Given the spec-project context (broad adoption across 9+ independent platform implementations demonstrates the spec is not being steered in a closed direction), this is treated as a non-blocking recommendation. Unlike pure spec projects (e.g., in-toto, TUF), Buildpacks also controls the reference implementation (`lifecycle`), which is the de facto conformance mechanism for the spec, making the concentration of TOC voting power more significant in practice. RFC#339 adoption is therefore particularly important and the project is expected to present a 6-month status update to the TOC on its adoption and implementation.
- Add a cross-reference from [GOVERNANCE.md#roadmap](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#roadmap) to the RFC process to make the roadmap change path more discoverable ([#292](https://github.com/buildpacks/community/issues/292)).
- **Conformance mechanism documentation** ([#299](https://github.com/buildpacks/community/issues/299)): The project has no standalone conformance test suite; interoperability is enforced through the `lifecycle` reference implementation. It is recommended to explicitly document that `lifecycle` serves as the de facto conformance mechanism for the spec, giving implementers a clear and authoritative reference.
- **Spec-level security guidance for implementers** ([#300](https://github.com/buildpacks/community/issues/300)): The current security work (third-party audit, SECURITY.md, response process) is implementation-focused. Since the spec involves executing untrusted code (buildpacks run during container builds), it is recommended to add spec-level security considerations for platform implementers: trust boundaries between platform, lifecycle, and buildpack execution; filesystem isolation requirements; and network access controls during builds.

### Adoption Evaluation

The adopter interviews confirm Cloud Native Buildpacks is in active production use at graduation maturity. Four adopters were interviewed (Google, Salesforce, Adopter 3, and 7SIGNAL), spanning cloud platforms, enterprise software, cybersecurity, and SaaS. All four report production usage at meaningful scale since 2020–2022, with Buildpacks serving as a core part of their container build and delivery workflows. Common themes across interviews include reproducible builds, reduced operational burden compared to Dockerfile-centric approaches, security benefits from rebasing and standardized base images, and positive engagement with the maintainer community. Detailed summaries are in the [Adoption](#adoption) section below.

Beyond the four interviewed adopters, the [public ADOPTERS.md](https://github.com/buildpacks/community/blob/main/ADOPTERS.md) lists 25+ organizations in production use — a mix of adopters using the CNB reference implementation (`lifecycle`, `pack`) and organizations with their own internal spec-compliant implementations. Among the latter, Google's independent implementation ([GoogleCloudPlatform/buildpacks](https://github.com/GoogleCloudPlatform/buildpacks)), which tracks the spec directly without using the reference lifecycle, is the strongest indicator that CNB functions as a genuine interoperability standard.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation. The TOC requests the project meet with the TOC in 6 months to present a status update on the non-blocking recommendations raised during this review. At that time the TOC would like to see:

- Progress on [RFC#339](https://github.com/buildpacks/rfcs/pull/339) (Steering Committee model with strengthened vendor-neutrality and end-user representation), including adoption timeline and any governance changes implemented.
- A plan or progress toward explicitly documenting the conformance mechanism for the spec ([#299](https://github.com/buildpacks/community/issues/299)).
- A plan or progress toward adding spec-level security guidance for platform implementers ([#300](https://github.com/buildpacks/community/issues/300)).

### Criteria

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**

<!-- (TOC Evaluation goes here) -->

  - Tech Review request: https://github.com/cncf/toc/issues/1983
  - General Technical Review (GTR): https://github.com/cncf/toc/pull/2047

- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

<!-- (TOC Evaluation goes here) -->

Project infrastructure and resources are vendor-neutral: the GitHub org, website, Slack, and communication channels are community-owned and not controlled by any single vendor:
- GitHub org(s): https://github.com/buildpacks, https://github.com/buildpacks-community
- Website: https://buildpacks.io/
- Slack: https://slack.buildpacks.io/
- TOC structure: https://github.com/buildpacks/community/blob/main/TEAMS.md#technical-oversight-committee

TOC seat composition is addressed as a non-blocking recommendation; the project has proactively opened [RFC#339](https://github.com/buildpacks/rfcs/pull/339) to restructure governance with a stronger representation cap and end-user seats. See the non-blocking recommendations in the summary and the "Governance clearly documents vendor-neutrality" criterion for full detail.

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		

<!-- (TOC Evaluation goes here) -->

   - Met during initial application on 31-08-2018.
   - Additional review/acknowledgement: TOC sponsor kickoff meeting with Buildpacks maintainers on 2025-12-16 to set Graduation expectations.

Acknowledgement evidence against prior CNCF TOC artifacts:
- Sandbox proposal PR: https://github.com/cncf/toc/pull/150
- Incubation review PR: https://github.com/cncf/toc/pull/338
- Process / criteria references:
  - https://github.com/cncf/toc/blob/main/process/README.md
  - https://github.com/cncf/toc/blob/main/process/graduation_criteria.md

- [ ] **Due Diligence Review.**

<!-- (TOC Evaluation goes here) -->

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

<!-- (TOC Evaluation goes here) -->

Documentation links for this criterion:
- Installation documentation: https://buildpacks.io/docs/app-journey/
- Contributing/community documentation: https://buildpacks.io/community/#contributing
- “Cloud Native Buildpacks Incubation Due Diligence” (Google Doc): https://docs.google.com/document/d/1tb3mK5cJmaQLO8xR__9NaH2GMrdn3WPjAZFBJYsXrxY/edit?usp=sharing

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

<!-- (TOC Evaluation goes here) -->

  - [GOVERNANCE.md in buildpacks/community](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md) documents formal governance structure including teams, roles, voting, RFC process, and vendor-neutrality commitments.
  - [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) and emeritus status demonstrate active evolution; substantial emeritus membership indicates role transitions and project maturation over time.

### Required

- [x] **Clear and discoverable project governance documentation.**

<!-- (TOC Evaluation goes here) -->

  - [GOVERNANCE.md in buildpacks/community](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md) provides comprehensive governance documentation (repositories, TOC, teams, roles, voting, RFC process, emergencies, vendor-neutrality, security, updating governance).
  - [README.md](https://github.com/buildpacks/community/blob/main/README.md) links to [GOVERNANCE.md](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md), making it highly discoverable.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

<!-- (TOC Evaluation goes here) -->

  - Technical Oversight Committee details documented in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md#technical-oversight-committee) with current member list and GitHub handles.
  - Committee activities tracked in [buildpacks/community Issues](https://github.com/buildpacks/community/issues), [Discussions](https://github.com/buildpacks/community/discussions), and [RFCs](https://github.com/buildpacks/rfcs).
  - TOC meetings are public with notes documented; [community calendar](https://buildpacks.io/community/) lists meeting times.
  - Team Leads, maintainers, and component maintainers documented in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) with current membership and contact information.

- [x] **Governance clearly documents [vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

<!-- (TOC Evaluation goes here) -->

  - [GOVERNANCE.md#vendor-neutrality](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#vendor-neutrality) documents a vendor-neutrality commitment with a cap on seats per organization.
  - Current TOC has 3 active members, 2 of whom are from Heroku/Salesforce (related companies per CNCF charter), with 2 seats vacant. The project has opened [RFC#339](https://github.com/buildpacks/rfcs/pull/339) to restructure governance with a stronger representation cap and end-user seats.
  - Given the spec's broad adoption across 9+ independent platform implementations, the spec is demonstrably not being steered in a closed direction. This is treated as a non-blocking recommendation, with a 6-month status update to the TOC expected on RFC#339 adoption.

Vendor-neutrality is documented in governance. The ongoing RFC#339 addresses current TOC composition and is tracked as a non-blocking recommendation.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

<!-- (TOC Evaluation goes here) -->

  - Changes to leadership roles, contribution acceptance, and governance follow the [RFC process](https://github.com/buildpacks/rfcs#rfc-process) per [GOVERNANCE.md#rfc-process](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#rfc-process).
  - Day-to-day decisions use lazy consensus (no objection + at least one approval), while higher-stakes decisions such as leadership elections and governance changes require a supermajority. Both thresholds are documented in [GOVERNANCE.md](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md).

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

<!-- (TOC Evaluation goes here) -->

  - Roles documented: Team Lead, Maintainer, Component Maintainer, Contributor, TOC member; nomination and election processes defined for each.
  - All team roles described in [GOVERNANCE.md#teams](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#teams) and [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md).
  - Security Response Team: [GOVERNANCE.md#security](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#security) documents composition ("all TOC members and all Team Leads") and responsibilities. Assignment and removal follow the underlying TOC/Team Lead role lifecycle.
  - Progress delivered by [buildpacks/community/pull#291](https://github.com/buildpacks/community/pull/291): added an explicit contributor ladder and role-summary matrix to [contributors/guide.md](https://github.com/buildpacks/community/blob/main/contributors/guide.md), and updated [GOVERNANCE.md](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md) with Team Lead appointment, Security Response Team composition, and 2FA requirement.
  - Inactivity-based offboarding is covered by the emeritus process in [GOVERNANCE.md#emeritus](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#emeritus) for all roles.

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

<!-- (TOC Evaluation goes here) -->

  - [GOVERNANCE.md#maintainers](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#maintainers): "New maintainers must already be contributors, must be nominated by an existing maintainer, and must be elected by a supermajority of CNB Team Leads and the TOC. Likewise, maintainers may resign or be removed by a super-majority of Team Leads and the TOC."
  - Emeritus process: [GOVERNANCE.md#emeritus](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#emeritus) documents self-nomination path, 3-month inactivity threshold, and reactivation process. "Out of courtesy, a notice for nomination should be given to members that fall under this category prior to being nominated."
  - Onboarding guidance: [contributors/guide.md](https://github.com/buildpacks/community/blob/main/contributors/guide.md) provides general contributor onboarding.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

<!-- (TOC Evaluation goes here) -->

  - [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) documents current active members and extensive emeritus lists across all teams (Implementation, Platform, Learning, Buildpack Authors' Tooling).
  - Git history examples: [buildpacks/community commit c837cd1](https://github.com/buildpacks/community/commit/c837cd15ce1e70f2a968be114b44e8f64d39a601#diff-7a8831df97c49707228dec1240256385588189e5e1bce211d5cdb23e7f94dfceR19) shows Natalieparellano progression from contributor ([commit 0fb5283](https://github.com/buildpacks/community/commit/0fb5283eb0458e655e8a65d1d6362c3d9be021c6)) to maintainer ([commit ee0965f](https://github.com/buildpacks/community/commit/ee0965f6802f15e534060ba3909849d0a96b24dc)) to TOC member.
  - [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) shows substantial emeritus membership, indicating the project actively uses role transitions and has experienced team evolution.

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

<!-- (TOC Evaluation goes here) -->

  - [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) provides current TOC members, Team Leads, Maintainers, and Component Maintainers with GitHub handles and team assignments.
  - [GOVERNANCE.md#technical-oversight-committee](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md) describes TOC responsibilities and links to TEAMS.
  - Project contact is documented at [cncf-buildpacks-maintainers@lists.cncf.io](mailto:cncf-buildpacks-maintainers@lists.cncf.io).
  - Affiliation gap tracked in [buildpacks/community/issues#294](https://github.com/buildpacks/community/issues/294) and resolved via [buildpacks/community/pull#301](https://github.com/buildpacks/community/pull/301): all team listings in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) converted to Markdown tables with an explicit Organization column covering active and emeritus members across all teams.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) -->

  - 3 active Technical Oversight Committee members (2 vacant seats out of max 5).
  - 5 Team Leads across 5 functional teams.
  - Approximately 10+ active Maintainers across teams.
  - ~20+ Project Contributors and Component Maintainers listed in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md).
  - Project scope: multi-component specification, reference implementation (pack, lifecycle, imgutil), registry ecosystem, documentation, and buildpack tooling. Leadership breadth (TOC, Team Leads, Maintainers across 5 teams) is appropriate to project scale and maturity.

- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

<!-- (TOC Evaluation goes here) -->

  - Active maintainers are distributed across multiple independent organizations across all five teams, as documented in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md).

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

<!-- (TOC Evaluation goes here) -->

  - [CODEOWNERS in buildpacks/community](https://github.com/buildpacks/community/blob/main/CODEOWNERS) assigns governance files ([GOVERNANCE.md](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md), [ROADMAP.md](https://github.com/buildpacks/community/blob/main/ROADMAP.md), OWNERS, [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md), [README.md](https://github.com/buildpacks/community/blob/main/README.md)) to @buildpacks/toc and @buildpacks/learning-maintainers.
  - All major component repositories include CODEOWNERS files (e.g., [pack](https://github.com/buildpacks/pack/blob/main/CODEOWNERS), [lifecycle](https://github.com/buildpacks/lifecycle/blob/main/CODEOWNERS)) defining ownership per documented teams.
  - GitHub Teams (e.g., @buildpacks/platform-maintainers, @buildpacks/toc) mirror documented roles in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md).

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

<!-- (TOC Evaluation goes here) -->

  - [buildpacks/.github/CODE_OF_CONDUCT.md](https://github.com/buildpacks/.github/blob/main/CODE_OF_CONDUCT.md) explicitly states: "This project follows the CNCF Code of Conduct."
  - [GOVERNANCE.md](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md) references CoC: "fostering a healthy and welcoming community, including by defining and enforcing our [Code of Conduct](https://github.com/buildpacks/.github/blob/main/CODE_OF_CONDUCT.md)."
  - The org-level CoC was updated to explicitly follow the CNCF CoC in response to [buildpacks/community/issues#285](https://github.com/buildpacks/community/issues/285), merged via [buildpacks/.github/pull#5](https://github.com/buildpacks/.github/pull/5).

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

<!-- (TOC Evaluation goes here) -->

  - [GOVERNANCE.md](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md) links to [buildpacks/.github/CODE_OF_CONDUCT.md](https://github.com/buildpacks/.github/blob/main/CODE_OF_CONDUCT.md).
  - [README.md](https://github.com/buildpacks/community/blob/main/README.md) also links to CoC.
  - The org-level CoC now explicitly links to the CNCF CoC, establishing a clear cross-link chain from governance documents to the CNCF CoC.

- [x] **All subprojects, if any, are listed.**

<!-- (TOC Evaluation goes here) -->

  - Each team and the subprojects/repositories they own are listed under [GOVERNANCE.md#teams](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#teams).
  - The [buildpacks-community GitHub organization](https://github.com/buildpacks-community) hosts independent community subprojects (e.g., kpack) with their own governance.

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

<!-- (TOC Evaluation goes here) -->

  - [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) documents team leadership and contributor roles; each team owns specific repositories and components, making subproject ownership and leadership discoverable.
  - For a spec project, subprojects are the reference implementations orbiting the spec. Team-based ownership documented in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md) and [GOVERNANCE.md#teams](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#teams) covers leadership, contribution paths, and the add/remove process (via RFC and TOC decision) for each subproject.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) -->

  - [contributors/guide.md](https://github.com/buildpacks/community/blob/main/contributors/guide.md) includes an explicit contributor ladder diagram and progression path through Team Lead and TOC.
  - The same file includes a role-summary table that documents eligibility, selection, and lifecycle expectations for Project Contributor, Component Maintainer, Maintainer, Team Lead, and TOC Member.
  - The contributor ladder and role-summary table were added in response to [buildpacks/community/issues#282](https://github.com/buildpacks/community/issues/282), merged via [buildpacks/community/pull#291](https://github.com/buildpacks/community/pull/291).

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) -->

  - [buildpacks/.github/CONTRIBUTING.md](https://github.com/buildpacks/.github/blob/main/CONTRIBUTING.md) documents the pull-request workflow and DCO sign-off requirements across all Buildpacks repositories.
  - [buildpacks.io/community](https://buildpacks.io/community/) and [contributors/guide.md](https://github.com/buildpacks/community/blob/main/contributors/guide.md) cover multiple contribution types: code, docs, RFCs, and triage, with explicit pointers to each path.
  - [README.md](https://github.com/buildpacks/community/blob/main/README.md) serves as the main entry point linking to all contribution and governance resources.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) -->

  - [buildpacks.io/community](https://buildpacks.io/community/) documents public channels including Slack (`#buildpacks`), GitHub Discussions, Mailing List, and StackOverflow.
  - [README.md](https://github.com/buildpacks/community/blob/main/README.md#meetings) documents public working-group meetings and links to attendance details.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) -->

  - All public channels (Slack, GitHub Discussions, Mailing List, StackOverflow) are named and linked on [buildpacks.io/community](https://buildpacks.io/community/), including team-scoped Slack channels for each sub-team (covering implementation, platform, learning, buildpack authors, maintainers, and more).
  - Non-public channel: `security@buildpacks.io` and GitHub Security Advisories are used for private vulnerability reporting and are documented with their purpose in [buildpacks/.github/SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md).

[buildpacks.io/community](https://buildpacks.io/community/) provides a comprehensive inventory of all public channels. The only non-public channel (security reporting) is documented in [SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md) with its purpose.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) -->

  - [buildpacks.io/community](https://buildpacks.io/community/) publishes recurring working-group meeting cadence, public Zoom links, agenda link, and a visible meeting schedule/calendar block.
  - [README.md#meetings](https://github.com/buildpacks/community/blob/main/README.md#meetings) points contributors to the public community meeting details.
  - Meeting-recording links were updated to the current YouTube channel in response to [buildpacks/community/issues#283](https://github.com/buildpacks/community/issues/283), merged via [buildpacks/docs/pull#882](https://github.com/buildpacks/docs/pull/882).

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) -->

  - [buildpacks/.github/CONTRIBUTING.md](https://github.com/buildpacks/.github/blob/main/CONTRIBUTING.md) documents baseline contribution workflow and DCO sign-off expectations used across Buildpacks repositories.
  - [contributors/guide.md](https://github.com/buildpacks/community/blob/main/contributors/guide.md) provides structured contribution guidance across contribution types (code, docs, RFCs, triage), contributor ladder, role expectations, and FAQs.
  - Role progression and lifecycle detail was strengthened in response to [buildpacks/community/issues#282](https://github.com/buildpacks/community/issues/282), resolved via [buildpacks/community/pull#291](https://github.com/buildpacks/community/pull/291).
  - [buildpacks.io/community](https://buildpacks.io/community/) provides newcomer-oriented paths (`good first issue` links, repository pointers, and contribution onboarding guidance), complementing the deeper governance/contributor docs.

- [x] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) -->

  - Recruitment pathways are actively documented: `good first issue` entry points and onboarding guidance on [buildpacks.io/community](https://buildpacks.io/community/), plus role progression in [contributors/guide.md](https://github.com/buildpacks/community/blob/main/contributors/guide.md).
  - Ongoing community activity is visible through recurring public working-group meetings and open discussion channels (Slack, mailing list, GitHub Discussions).
  - [DevStats](https://buildpacks.devstats.cncf.io/) shows sustained activity across commits, PRs, issues, and contributors over the project lifetime, with continued throughput in recent windows. See [Overall Project Statistics](https://buildpacks.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1), [Developer Activity by Company](https://buildpacks.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-metric=contributions), [New Contributors](https://buildpacks.devstats.cncf.io/d/52/new-contributors-table?orgId=1&var-period_name=Last%20year&from=now-1y&to=now), and [Issues/PRs flow](https://buildpacks.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1) for current numbers.
  - [LFX Insights](https://insights.linuxfoundation.org/project/buildpacks?timeRange=past365days&start=2025-06-27&end=2026-06-27) (past 365 days) reports an overall health score of **Excellent**, with active quarterly contributors and healthy retention. Contribution activity is concentrated in a small set of core maintainers and organizations , consistent with the spec project model, and reinforcing the vendor-neutrality recommendations already noted in this review.

Contributor activity is positive and sustained. Exact metrics are best read from DevStats and LFX Insights directly, as they reflect live data.

## Engineering Principles

### Suggested

N/A

### Required

**Note:** A General Technical Review was completed on 20-Feb-2026, and can be discovered at [cncf/toc/pull#2047](https://github.com/cncf/toc/pull/2047). The following criteria are satisfied through this review and supporting documentation.

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**

<!-- (TOC Evaluation goes here) -->

  - Satisfied via General Technical Review: [cncf/toc/pull#2047](https://github.com/cncf/toc/pull/2047).
  - The GTR captures CNB differentiation, project objectives, and why CNB addresses cloud-native build and supply-chain needs differently from Dockerfile-centric workflows.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases. _This requirement may also be satisfied by completing a General Technical Review._**

<!-- (TOC Evaluation goes here) -->

  - Satisfied via General Technical Review: [cncf/toc/pull#2047](https://github.com/cncf/toc/pull/2047).
  - The GTR documents what CNB does, why it exists, and representative cloud-native use cases.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

<!-- (TOC Evaluation goes here) -->

  - Public Roadmap maintained at [buildpacks/community ROADMAP.md](https://github.com/buildpacks/community/blob/main/ROADMAP.md)
  - Recent planning documented in [buildpacks/rfcs/pull#325](https://github.com/buildpacks/rfcs/pull/325)
  - The Technical Oversight Committee (TOC) is explicitly responsible for roadmap direction, team leadership, and cross-cutting concerns as documented in [TEAMS.md](https://github.com/buildpacks/community/blob/main/TEAMS.md#technical-oversight-committee)
  - Discoverability gap noted in [buildpacks/community/issues#292](https://github.com/buildpacks/community/issues/292): roadmap history is not clearly presented for all expected periods (for example, 2024/2025 are not clearly discoverable from the current roadmap index/path).

Public roadmap exists with clear ownership; remaining discoverability gaps (including missing/unclear 2024/2025 roadmap visibility) are tracked in [buildpacks/community/issues#292](https://github.com/buildpacks/community/issues/292).

- [x] **Roadmap change process is documented.**

<!-- (TOC Evaluation goes here) -->

  - The RFC process serves as the roadmap change process for this spec project; [GOVERNANCE.md#rfc-process](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#rfc-process) documents how changes are proposed, discussed, and approved.
  - A minor documentation enhancement to add a cross-reference from [GOVERNANCE.md#roadmap](https://github.com/buildpacks/community/blob/main/GOVERNANCE.md#roadmap) to the RFC process is tracked in [buildpacks/community/issues#292](https://github.com/buildpacks/community/issues/292).

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.  *This requirement may also be satisfied by completing a General Technical Review and capturing the output in the project's documentation.***

<!-- (TOC Evaluation goes here) -->

  - Satisfied via General Technical Review: [cncf/toc/pull#2047](https://github.com/cncf/toc/pull/2047)
  - Architecture-related documentation referenced in the GTR:
    - [Platform API (spec interface and architecture contract)](https://buildpacks.io/docs/reference/spec/platform-api/)
    - [Buildpack API (lifecycle/buildpack interface contract)](https://buildpacks.io/docs/reference/spec/buildpack-api/)
    - [What is the lifecycle? (phased build architecture and security boundaries)](https://buildpacks.io/docs/for-platform-operators/concepts/lifecycle/)
    - [What is a builder? (builder composition and component model)](https://buildpacks.io/docs/for-platform-operators/concepts/builder/)
    - [What is a buildpack? (detect/build behavior and transformation model)](https://buildpacks.io/docs/for-platform-operators/concepts/buildpack/)

- [x] **Document the project's release process and guidelines publicly in a [RELEASES.md](https://github.com/buildpacks/community/blob/main/RELEASES.md) or equivalent file that defines:** 

<!-- (TOC Evaluation goes here) -->

  - [community/RELEASES.md](https://github.com/buildpacks/community/blob/main/RELEASES.md) serves as a project-level index to component-specific release docs, with each component maintaining its own release process.

  - [x] Release expectations (scheduled or based on feature implementation)
    - [pack/RELEASE.md](https://github.com/buildpacks/pack/blob/main/RELEASE.md) documents a 6-week release cadence; [lifecycle/RELEASE.md](https://github.com/buildpacks/lifecycle/blob/main/RELEASE.md) documents minor/patch/backport release types.
  - [x] Tagging as stable, unstable, and security related releases
    - Lifecycle categorises CVE fixes as "New patch" releases in [lifecycle/RELEASE.md](https://github.com/buildpacks/lifecycle/blob/main/RELEASE.md); pack documents RC (pre-release) and final release tagging. There is no project-wide release; tagging is handled per component.
  - [x] Information on branch and tag strategies
    - Branch and tag strategies are documented in [pack/RELEASE.md](https://github.com/buildpacks/pack/blob/main/RELEASE.md), [lifecycle/RELEASE.md](https://github.com/buildpacks/lifecycle/blob/main/RELEASE.md), and [spec/RELEASE.md](https://github.com/buildpacks/spec/blob/main/RELEASE.md).
  - [x] Branch and platform support and length of support
    - API deprecation and support policy is governed by [RFC #0110](https://github.com/buildpacks/rfcs/blob/main/text/0110-deprecate-apis.md).
  - [x] Artifacts included in the release.
    - Release artifacts and changelogs are published in [Pack releases](https://github.com/buildpacks/pack/releases) and [Spec releases](https://github.com/buildpacks/spec/releases).

Component-level release docs are at [pack/RELEASE.md](https://github.com/buildpacks/pack/blob/main/RELEASE.md), [lifecycle/RELEASE.md](https://github.com/buildpacks/lifecycle/blob/main/RELEASE.md), and [spec/RELEASE.md](https://github.com/buildpacks/spec/blob/main/RELEASE.md).

- [x] **History of regular, quality releases.**

<!-- (TOC Evaluation goes here) -->

  - All three core components have maintained a regular release cadence since their inception: [pack releases](https://github.com/buildpacks/pack/releases) (since 2018), [lifecycle releases](https://github.com/buildpacks/lifecycle/releases) (since 2019), and [spec releases](https://github.com/buildpacks/spec/releases) (versioned API releases ongoing since 2020).

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

<!-- (TOC Evaluation goes here) -->

Buildpacks has an OpenSSF Best Practices **passing** badge (https://www.bestpractices.dev/en/projects/4748). Silver/gold remains a non-blocking improvement target.

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) -->

Buildpacks documents a clear reporting process in [SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md) with:
- private reporting channels (GitHub Security Advisories and security@buildpacks.io),
- required report contents,
- and explicit response workflow.
- https://github.com/buildpacks/.github/blob/main/SECURITY.md
- Reporting process was enhanced with clearer report content requirements, response workflow/timelines, supported versions policy, and disclosure guidance in response to [buildpacks/community/issues#284](https://github.com/buildpacks/community/issues/284), merged via [buildpacks/.github/pull#6](https://github.com/buildpacks/.github/pull/6).

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) -->

The project enforces access-control and repository ownership controls across Buildpacks repositories, including CODEOWNERS-based review ownership and organization/repository policy controls.
- Example CODEOWNERS usage: https://github.com/buildpacks/lifecycle/blob/main/CODEOWNERS
- Security process and disclosure governance: https://github.com/buildpacks/.github/blob/main/SECURITY.md

- [x] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) -->

[SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md) documents that the project security team (maintainers) handles reports, acknowledges reports within 72 hours, triages with assigned maintainers, coordinates fixes privately, and discloses through advisories and release notes.
- https://github.com/buildpacks/.github/blob/main/SECURITY.md
- Security response roles and handling process were clarified via [buildpacks/.github/pull#6](https://github.com/buildpacks/.github/pull/6) in response to [buildpacks/community/issues#284](https://github.com/buildpacks/community/issues/284).

- [x] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) -->

Buildpacks has a published markdown self-assessment in this CNCF-discoverable repo:
- https://github.com/cncf/toc/blob/main/projects/buildpacks/security-assessment/self-assessment.md

Related context:
- Self-assessment discoverability was improved in response to [buildpacks/community/issues#286](https://github.com/buildpacks/community/issues/286), with audit and self-assessment links added to [SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md) via [buildpacks/.github/pull#8](https://github.com/buildpacks/.github/pull/8).

- [x] **Third Party Security Review.**

<!-- (TOC Evaluation goes here) -->

Buildpacks completed a third-party security audit with OSTIF/Quarkslab and published findings publicly.
- https://ostif.org/buildpacks-audit-complete/
- https://medium.com/buildpacks/announcing-findings-from-security-audit-b4701f4e8b4b
- Audit references are linked from [SECURITY.md](https://github.com/buildpacks/.github/blob/main/SECURITY.md).

  - [x] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings.

    All moderate and low findings from the audit have been tracked and addressed across pack and lifecycle milestones. See [pack 0.35.0 milestone](https://github.com/buildpacks/pack/milestone/50) and [lifecycle 0.20.0 milestone](https://github.com/buildpacks/lifecycle/milestone/41) for remediation rollout.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) -->

Buildpacks has achieved the OpenSSF Best Practices passing badge:
- https://www.bestpractices.dev/en/projects/4748

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

<!-- (TOC Evaluation goes here) -->

Buildpacks maintains a public adopters list in the community repository.
- https://github.com/buildpacks/community/blob/main/ADOPTERS.md

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

<!-- (TOC Evaluation goes here) -->

Four adopters were interviewed (Google, Salesforce, Adopter 3, 7SIGNAL). All four reported production usage at meaningful scale and described Buildpacks as core to container build and delivery workflows.

An adopter list was reviewed to verify level-appropriate usage (production use for graduation, dev/test use for incubation).

- [x] **TOC verification of adopters.**

<!-- (TOC Evaluation goes here) -->

Verification was completed through four adopter interviews and approved summaries captured in this folder.

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

<!-- (TOC Evaluation goes here) -->

Buildpacks documentation includes integration guidance for major CNCF and non-CNCF ecosystems (for example Kubernetes, Tekton, Podman, Docker, and OCI-compatible registries/runtimes).
- https://buildpacks.io/docs/
- https://buildpacks.io/docs/for-platform-operators/how-to/integrate-ci/tekton/
- https://buildpacks.io/docs/for-app-developers/how-to/special-cases/build-on-podman/
- https://buildpacks.io/docs/for-platform-operators/how-to/integrate-ci/docker/


#### Adoption

The adopter interviews show Buildpacks is in production use at graduation maturity. Adopters report production-scale usage, measurable operational benefit, and positive maintainer/community engagement. The interviews also demonstrate the spec functioning as a genuine standard: Google maintains an independent implementation ([GoogleCloudPlatform/buildpacks](https://github.com/GoogleCloudPlatform/buildpacks)) that tracks the spec directly rather than using the reference implementation, while others use the reference implementation ecosystem (`pack` CLI, Paketo buildpacks), confirming adoption across both paths.

Common strengths across interviews:
- Consistent and reproducible source-to-image workflows at scale
- Reduced operational burden versus hand-maintained Dockerfiles
- Security and maintenance gains from rebasing, standardized base images, and SBOM-related workflows
- Good documentation and responsive maintainer/community interactions

Common improvement themes:
- Upgrade/version-alignment complexity across multiple Buildpacks components
- Migration friction for some legacy Dockerfile-heavy applications
- Continued improvement on multi-architecture support and rebase capabilities

##### Adopter 1 - Google / Cloud Platform (Serverless)

Interviewed: Jan 2026

Production usage across Cloud Run, Cloud Functions, and App Engine since 2020. Google reported Buildpacks as the standardized mechanism for source deployment, with very high build volume and strong value from open specification alignment and rebase-driven security updates.

Reference: [Buildpacks Adopter Interview - Google](./buildpacks-adopter-interview-google.md)

##### Adopter 2 - Salesforce / Enterprise Software (Internal CI)

Interviewed: Jan 2026

Production usage since 2022 in internal CI systems serving a large developer base, with sustained high daily build volume. Salesforce highlighted reproducibility, reduced manual effort, and strong utility for integration-test and configuration packaging workflows.

Reference: [Buildpacks Adopter Interview - Salesforce](./buildpacks-adopter-interview-salesforce.md)

##### Adopter 3 - Cybersecurity

Interviewed: Jan 2026

Production usage since 2022 across multiple teams. Adopter 3 described Buildpacks as a standard container-build path with custom buildpacks/builders where needed, citing security, rebasing, and standardization benefits, plus active community collaboration.

Reference: [Buildpacks Adopter Interview - Adopter 3](./buildpacks-adopter-interview-adopter-3.md)

##### Adopter 4 - 7SIGNAL / SaaS Network Monitoring

Interviewed: Jan 2026

Production usage for over three years across multiple applications. 7SIGNAL emphasized low setup overhead, low maintenance, consistent outcomes across projects, and improved security posture compared to Dockerfile-centered workflows.

Reference: [Buildpacks Adopter Interview - 7SIGNAL](./buildpacks-adopter-interview-7signalinc.md)
