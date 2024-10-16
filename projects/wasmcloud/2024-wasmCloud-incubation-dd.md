# WasmCloud Incubation Due Diligence

- Link to [Incubation application issue]()

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for $PROJECT

### Criteria Evaluation

_$TOCMEMBER conducted the due diligence of $PROJECT who applied for $LEVEL. The project [has/has not] completed the criteria that show its maturity at $LEVEL. The following criteria implementations are noteworthy to call out... $NOTABLES. The following actions were provided to the project that were considered blocking but since resolved... $BLOCKERS. The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project... $RECOMMENDATIONS._

### Adoption Evaluation

_The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY._

### Final Assessment

_[The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS]._

## Application Process Principles

### Suggested

N/A

### Required

- [ ] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**
  - This was completed and occurred on DD-MMM-YYYY, and can be discovered at $LINK.

<!-- (TOC Evaluation goes here) --> 

- [ ] **TAG provides insight/recommendation of the project in the context of the landscape**

<!-- (TOC Evaluation goes here) --> 

- [ ] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
- Met during Project's application on DD-MMM-YYYY.

<!-- (TOC Evaluation goes here) --> 
https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/ADOPTERS.md
- [ ] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisifies the Due Diligence Review criteria.

- [ ] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

<!-- (TOC Evaluation goes here) --> 

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Clear and discoverable project governance documentation.**

  WasmCloud's governace is listed at the root of the WasmCloud repo in [GOVERNANCE.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/GOVERNANCE.md)

- [ ] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

<!-- (TOC Evaluation goes here) --> 

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  WasmCloud has two processes for decision making scoped at the org (top) level or at the (sub)project level.
  - [Org Level](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/GOVERNANCE.md#decision-making-at-the-wasmcloud-org-level) - Adding, archiving, or removing subprojects, CoC type issuese etc.
  - [Project Level](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/GOVERNANCE.md#decision-making-at-the-wasmcloud-project-level) - Day to day operations or [Architectural Decision Records](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/CONTRIBUTING.md#proposing-an-idea)

  Both use a lazy consensus model, or where required a super/simple-majority. They do explicitly mention that when proposing a design change that there should be sufficient input from different invested parties (e.g. vendors) before a decision is made.
  This level of decision-making is sufficient for a project of this size.

- [ ] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**



- [ ] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) --> 

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

  They do and align with the information in [MAINTAINERS.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/MAINTAINERS.md#maintainers). These are mapped to the CODEOWNERS files used in each repo.


- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

It is documented as a requirement for [each role defined in the contributor ladder](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/CONTRIBUTION_LADDER.md).


- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

  It is cross-linked in [GOVERNANCE.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/GOVERNANCE.md#code-of-conduct)


- [x] **All subprojects, if any, are listed.**

  Their subprojects are enumerated along with their owners in [MAINTAINERS.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/MAINTAINERS.md#maintainers)


## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

  The project details a contributor ladder with three levels along with an additional role for org management and an explicit emeritus status.

  - [Community Member](https://github.com/wasmCloud/wasmCloud/blob/c0f57e4b7fa810d8c4cd6cc42de2a8efa6b9f0f7/CONTRIBUTION_LADDER.md#community-member)
  - [Contributor](https://github.com/wasmCloud/wasmCloud/blob/c0f57e4b7fa810d8c4cd6cc42de2a8efa6b9f0f7/CONTRIBUTION_LADDER.md#contributor)
  - [Project Maintainer](https://github.com/wasmCloud/wasmCloud/blob/c0f57e4b7fa810d8c4cd6cc42de2a8efa6b9f0f7/CONTRIBUTION_LADDER.md#project-maintainer)
  - [Org Manintainers](https://github.com/wasmCloud/wasmCloud/blob/c0f57e4b7fa810d8c4cd6cc42de2a8efa6b9f0f7/CONTRIBUTION_LADDER.md#org-maintainers) (Additional role with Org Management permissions)
  - [Emeritus Maintainers](https://github.com/wasmCloud/wasmCloud/blob/c0f57e4b7fa810d8c4cd6cc42de2a8efa6b9f0f7/CONTRIBUTION_LADDER.md#emeritus-maintainers)

  Each role has a list of responsibilities and details the steps needed to obtain that role.



### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

  The [WasmCloud Contributor Guide](https://wasmcloud.com/docs/contributing/contributing-guide) is mirrored from GitHub with the direct links from from the front page of the [WasmCloud website](https://wasmcloud.com/) and fully details the [Issue Submission](https://wasmcloud.com/docs/contributing/contributing-guide#issues) and [PR process](https://wasmcloud.com/docs/contributing/contributing-guide#pull-requests).

  They additionally outline and link to their process for feature requests which requires am [Architectural Decision Record](https://wasmcloud.com/docs/contributing/contributing-guide#proposing-an-idea), a form of design proposal.


- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

  WasmCloud's primary communication platform is Slack. It is used for both users and contributors.
The [wasmcloud slack](https://slack.wasmcloud.com/) is linked from from their site, and various places in GitHub.


- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

  - [WasmCloud slack](https://slack.wasmcloud.com/) - Used for all major communications for both users & contributors
  - [Maintainer Mailing List]() - Private use for maintainer notifications from CNCF
  - WasmCloud Mailing List | TBD - For general purpose async comms


- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

  Their schedule and meeting agenda/notes are available on the project webiste under the [community section](https://wasmcloud.com/community). They also regularly [stream & post their meetings to YouTube](https://www.youtube.com/@wasmCloud/streams).



- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

  The [WasmCloud Contributor Guide](https://wasmcloud.com/docs/contributing/contributing-guide) is mirrored from GitHub with the direct links from from the front page of the [WasmCloud website](https://wasmcloud.com/) and fully details the [Issue Submission](https://wasmcloud.com/docs/contributing/contributing-guide#issues) and [PR process](https://wasmcloud.com/docs/contributing/contributing-guide#pull-requests).

  They additionally outline and link to their process for feature requests which requires am [Architectural Decision Record](https://wasmcloud.com/docs/contributing/contributing-guide#proposing-an-idea), a form of design proposal.


- [x] **Demonstrate contributor activity and recruitment.**



## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

The `[wasmCloud roadmap](https://wasmcloud.com/docs/roadmap/)` is discussed openly and transparently during planned quarterly open community meetings. Example: `[Q3 planning meeting](https://wasmcloud.com/community/2024/07/02/community-meeting)`. Outside of the quarterly meetings, there is no documented roadmap process. TOC recommends that after meeting the requirements for Incubation, a process is created to document more thoroughly how features are added to the roadmap, as well as document a roadmap change process. This is being tracked by `[wasmCloud issue 3343](https://github.com/wasmCloud/wasmCloud/issues/3343)`.

- [ ] **History of regular, quality releases.**

The wasmCloud/wasmCloud monorepo `[had frequent releases](https://github.com/wasmCloud/wasmCloud/releases)` between February 2021 and May 2021. The project was admitted to CNCF Sandbox in July 2021. Between May 2021 and August 2023, there were no releases of the project. The project applied for CNCF Incubation in October 2023. Recommend that the project continues regular, quality releases post moving levels to CNCF Incubation.

### Required 

- [ ] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document what the project does, and why it does it - including viable cloud native use cases.**

<!-- (TOC Evaluation goes here) --> 

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The wasmCloud project documents and maintains a `[public roadmap](https://github.com/orgs/wasmCloud/projects/7)` as a GitHub project and is `[cross-linked to their website](https://wasmcloud.com/docs/roadmap/)`.

- [ ] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

<!-- (TOC Evaluation goes here) --> 

- [x] **Document the project's release process.**

The project's release process is documented by a `[RELEASE.md](https://github.com/wasmCloud/wasmCloud/blob/main/RELEASE.md)` and `[RELEASE_RUNBOOK.md](https://github.com/wasmCloud/wasmCloud/blob/main/RELEASE_RUNBOOK.md)` documents. Recommend the project update the Release documentation with a more specific schedule.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

  Information on reporting a security issue is present in [SECUIRTY.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/SECURITY.md) and it is cross-linked from their website.


- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  **NOTE:** 2FA is a requirement that is enforced via the CNCF GitHub Enterprise account.

  WasmCloud limits broader GitHub access to a designated "[Org Maintainer](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/CONTRIBUTION_LADDER.md#org-maintainers)" role. This role is a smaller subset of the maintainers, limiting general access.

  While not a blocker, this could be improved by adopting some tooling such as [clowarden](https://github.com/cncf/clowarden), [prow & peribolos](https://docs.prow.k8s.io/docs/components/cli-tools/peribolos/), [sheriff](https://github.com/electron/sheriff) or other tool that can manage org membership & permissions in a more transparent manner.

- [x] **Document assignment of security response roles and how reports are handled.**

The security response roles and security process are thoroughly documented in the `[SECURITY.md](https://github.com/wasmCloud/wasmCloud/blob/main/SECURITY.md)` file in the wasmCloud/wasmCloud monorepo.

- [ ] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The wasmCloud project has achieved the `[OpenSSF best practices passing badge](https://www.bestpractices.dev/en/projects/6363)`.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

  Adopters are listed in [ADOPTERS.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/ADOPTERS.md) at the root of the WasmCloud repo.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [ ] **TOC verification of adopters.**

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The wasmCloud project maintains `[a list of integrations](https://wasmcloud.com/docs/integrations/)` on their website of both CNCF projects as well as non-CNCF projects.

#### Adoption

##### Adopter 1 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR

##### Adopter 2 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR

##### Adopter 3 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR
