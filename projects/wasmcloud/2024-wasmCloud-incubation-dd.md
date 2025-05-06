# WasmCloud Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1338)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for wasmCloud

### Criteria Evaluation

[Karena Angell](https://github.com/angellk) conducted the due diligence of wasmCloud who applied for Incubation, while [Davanum Srinivas](https://github.com/dims) and [Duffie Cooley](https://github.com/mauilion) assisted with Adopter interviews. The project has completed the criteria that show its maturity at the applied level.

The following criteria implementations are noteworthy to call out:

- The TOC reviewers recommend continuing to review the project Governance and Contributor ladder to ensure named maintainers are active and responsive. In terms of expanding beyond converting end users into contributors, it would be good to look at mentoring programs like GSoC, LFX etc as well. With WebAssembly being a newer technology, it has been slower to attract larger enterprise vendors as maintainers. To ensure wasmCloud sustainability, the TOC recommends a focus on increasing the diversity of vendor and end user representation for the core Org maintainers.
- Currently there is no specific section or policy on vendor-neutrality. A suggestion before applying for graduation would be to integrate a policy such as maximal representation for the decision-making bodies.
- The completion of the [external security audit](https://ostif.org/ostif-has-completed-a-security-audit-of-wasmcloud/) by Trail of Bits and the lack of any high or critical issues demonstrate a strong commitment to security. All five identified issues have been addressed, which is commendable.
- wasmCloud has a complete [contributing guide](https://wasmcloud.com/docs/contributing/), including information general guidance, contributing to docs, SDKs, info on the roadmap etc. This is complemented with a dedicated github repository [wasmCloud/wasmCloud-contrib](https://github.com/wasmCloud/wasmCloud-contrib) for community contributions to providers, components and demos.
- The project does a commendable job ensuring its [public meetings](https://wasmcloud.com/community) - along with notes, and videos - are easy to find. Their ease of discoverability is something other project should look to replicate.
- The wasmCloud project uses composable [components](https://wasmcloud.com/docs/concepts/components) that can be compiled from various languages allowing for interoperability across various public and private environments.
- The wasmCloud project follows [Open Standards](https://wasmcloud.com/docs/concepts/components#open-standards) for its componentized ecosystem and [supports open tooling](https://wasmcloud.com/docs/ecosystem/useful-webassembly-tools/) from the WebAssembly community and ByteCode Alliance. 

The following actions were provided to the project that were considered blocking but have since been resolved:

- The project had not yet completed the TAG Security Self Assessment and subsequently [completed and submitted a thorough self assessment](https://github.com/wasmCloud/wasmCloud/issues/3302).
- The project did not have a Maintainers file that was viewable publicly due to the GitHub Org permissions. The project [created a public Maintainers file](https://github.com/wasmCloud/wasmCloud/issues/3296)]. 
- The project's [Release documentation was not sufficient](https://github.com/wasmCloud/wasmCloud/issues/3295)] to meet Incubation requirements and not easy to find. A new RELEASE.md file was created in the main repo and RELEASE_RUNBOOK.md was moved to the main repo as well.
- To meet vendor neutrality requirements, the TOC reviewers requested the project to [clean up references](https://github.com/wasmCloud/wasmCloud/issues/3297) to the Cosmonic Kubernetes Applier.
- TAG Contributor Strategy recommended [reconciling](https://github.com/cncf/foundation/pull/872) the MAINTAINERS.md file with the CNCF Project Maintainers file which has been completed.

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

- TAG Contributor Strategy recommends adopting some tooling such as [clowarden](https://github.com/cncf/clowarden), [prow & peribolos](https://docs.prow.k8s.io/docs/components/cli-tools/peribolos/), [sheriff](https://github.com/electron/sheriff) or other tool that can manage org membership & permissions in a more transparent manner.
- TOC reviewers [recommend](https://github.com/wasmCloud/wasmCloud/issues/3343) that a process is created to document more thoroughly how features are added to the roadmap, as well as document a roadmap change process.
- TOC reviewers recommend that the project continues regular, quality releases that are easy to find. Any changes should be thoroughly documented in the RELEASE.md and communicated clearly to the community.
- TOC reviewers recommend the project implement a more consistent, predictable release schedule to allow for greater enterprise end user adoption.
- TAG Contributor Experience [recommends](https://github.com/wasmCloud/wasmCloud/issues/3305) adding an on-boarding and off-boarding process for maintainers as well as create a canonical communication mechanism.
- TOC reviewers [recommend](https://github.com/wasmCloud/wasmCloud/issues/3343) additional roadmap process improvements such as how a feature is added to a release or removed from a release, as well as tagging releases for bug and security fixes. This will allow for greater transparency outside the core org maintainers and ease of adoption for end users in enterprise production environments.
- The [Security Reporting Process](https://github.com/wasmCloud/wasmCloud/blob/main/SECURITY.md) should be enhanced to include an [embargo](https://contribute.cncf.io/maintainers/security/security-guidelines/#34-embargo-policy). In addition, we feel that tagging a release clearly as "security-fixes-only" will help users to prioritize an upgrade. 
- TOC reviewers recommended clarifying the concept of Emeritus in various roles to acknowledge the work of folks who spent time in the community, but no longer are active as well as give a easy path for them to get active again when they choose to.
- Documentation updates to better explain the core wasmCloud architecture.
- TOC Reviewers recommend identifying clearly on the [Integrations](https://wasmcloud.com/docs/integrations/) page the supportability of each integration. For example, if a cloud native integration is supported by the identified cloud native project itself, supported by a vendor, or is held in an open repository to be self supported by the adopting end user.

### Adoption Evaluation

The adopter interviews reflect the project use for the level to which the project has applied. [wasmCloud](https://wasmcloud.com/) allows platform teams to easily build and scale WebAssembly applications in a standard, modular fashion. wasmCloud integrates with existing cloud native stacks and is working to close functionality gaps with Kubernetes, in particular. Adopters are excited to see the technology continue to mature to unlock greater use for enterprise production environments. The current application platform architecture has been designed with composability, security, high availability and resiliency in mind and is desirable for adopters who are creating agile, interoperable application platforms.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**
  - This was completed and occurred on 15-JUN-2023, and can be discovered at:
  - [Notes](https://docs.google.com/document/d/1k7VNetgbuDNyIs_87GLQRH2W5SLgjgOhB6pDyv89MYk/edit#heading=h.6kiq9agcez2t)
  - [Recording](https://youtu.be/R4yQrOH1JGA)

- [x] **TAG provides insight/recommendation of the project in the context of the landscape**

  This was brought up at the TAG Runtime meeting where wasmCloud presented.

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

  To meet vendor neutrality requirements, TOC reviewers requested the project to [=[clean up references](https://github.com/wasmCloud/wasmCloud/issues/3297)] to the Cosmonic Kubernetes Applier.
  
  The cloud native integrations have a mix of vendor and non-vendor resources. 

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
  - Met during Project's application on 13-JUL-2021.
  - [Onboarding Issue](https://github.com/cncf/sandbox/issues/230) 

- [x] **Due Diligence Review.**

[Initial Issue](https://github.com/cncf/toc/issues/1338) was submitted on 25-OCT-2023. The project opened a [new issue](https://github.com/cncf/toc/issues/1338) using the recommended template on 24-MAY-2024.

  The TOC and project maintainers met to discuss architecture, findings and discuss next steps. We created a shared, temporary channel with the TOC reviewers and the project leadership for communication. TOC reviewers opened public issues in various repositories asking questions and nudging the team to work on things all of which were cross linked to the main [issue](https://github.com/cncf/toc/issues/1338).

  The project maintainers have been highly responsive throughout the process.

  Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**
  
  WasmCloud's [documentation](https://wasmcloud.com/docs/intro) covers the majority of needed documentation, and includes other useful information such as their roadmap and links to get engaged with the community.

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Clear and discoverable project governance documentation.**

  WasmCloud's governance is listed at the root of the wasmCloud repo in [GOVERNANCE.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/GOVERNANCE.md)

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

  Their governance & contributing docs had a significant overhaul in [July 2022](https://github.com/wasmCloud/wasmCloud/commit/10a8da994f89434dbf556dc46553fc1275819d2c) with minimal changes since then and is appropriate for their size and maturity.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

  Governance is up to date with actual practices.

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  Currently there is no specific section or policy on vendor-neutrality.
  A suggestion before applying for graduation would be to integrate a policy such as maximal representation for the decision-making bodies.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  WasmCloud has two processes for decision making scoped at the org (top) level or at the (sub)project level.
  - [Org Level](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/GOVERNANCE.md#decision-making-at-the-wasmcloud-org-level) - Adding, archiving, or removing subprojects, CoC type issues etc.
  - [Project Level](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/GOVERNANCE.md#decision-making-at-the-wasmcloud-project-level) - Day to day operations or [Architectural Decision Records](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/CONTRIBUTING.md#proposing-an-idea)

  Both use a lazy consensus model, or where required a super/simple-majority. They do explicitly mention that when proposing a design change that there should be sufficient input from different invested parties (e.g. vendors) before a decision is made.
  This level of decision-making is sufficient for a project of this size.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

  All roles are detailed in [CONTRIBUTOR_LADDER.md](https://github.com/wasmCloud/wasmCloud/blob/18dccc362462cc70e82720bb4cb818bda9ae5b12/CONTRIBUTION_LADDER.md) doc with additional onboarding & offboarding material in [GOVERNANCE.md](https://github.com/wasmCloud/wasmCloud/blob/18dccc362462cc70e82720bb4cb818bda9ae5b12/GOVERNANCE.md#new-maintainer-onboarding)

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

  Onboarding, Offboarding, and emeritus status are detailed in [GOVERNANCE.md](https://github.com/wasmCloud/wasmCloud/blob/18dccc362462cc70e82720bb4cb818bda9ae5b12/GOVERNANCE.md#new-maintainer-onboarding).

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

  Several maintainers have been moved to [emeritus status](https://github.com/wasmCloud/wasmCloud/blob/18dccc362462cc70e82720bb4cb818bda9ae5b12/MAINTAINERS.md#emeritus-maintainers).

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

  N/A - WasmCloud does have repo specific maintainers, but they all use the same general processes.

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

  Each repo has an appropriate MAINTAINERS file with the current list of maintainers.
  
  Links to current pinned repo maintainer files:
  
  - [wasmcloud/wasmcloud](https://github.com/wasmCloud/wasmCloud/blob/main/MAINTAINERS.md)
  - [wasmcloud/wadm](https://github.com/wasmCloud/wadm/blob/main/MAINTAINERS.md)
  - [wasmcloud/wasmcloud.com](https://github.com/wasmCloud/wasmcloud.com/blob/main/MAINTAINERS.md)
  - [wasmcloud/wasmcloud-operator](https://github.com/wasmCloud/wasmcloud-operator/blob/main/MAINTAINERS.md)
  - [wasmcloud/component-sdk-go](https://github.com/wasmCloud/component-sdk-go/blob/main/MAINTAINERS.md)
  - [wasmcloud/wadge](https://github.com/wasmCloud/wadge/blob/main/MAINTAINERS.md)

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

  WasmCloud has a sufficient number of active maintainers; however many fall under a single organization. Before applying for graduation they should work towards improving their organizational diversity with conversion from contributors -> maintainers.

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
  - [Org Maintainers](https://github.com/wasmCloud/wasmCloud/blob/c0f57e4b7fa810d8c4cd6cc42de2a8efa6b9f0f7/CONTRIBUTION_LADDER.md#org-maintainers) (Additional role with Org Management permissions)
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

  - [wasmCloud slack](https://slack.wasmcloud.com/) - Used for all major communications for both users & contributors
  - [Maintainer Mailing List](https://lists.cncf.io/g/cncf-wasmCloud-maintainers/) - Private use for maintainer notifications from CNCF

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

  Their schedule and meeting agenda/notes are available on the project website under the [community section](https://wasmcloud.com/community). They also regularly [stream & post their meetings to YouTube](https://www.youtube.com/@wasmCloud/streams).

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

  The [WasmCloud Contributor Guide](https://wasmcloud.com/docs/contributing/contributing-guide) is mirrored from GitHub with the direct links from from the front page of the [WasmCloud website](https://wasmcloud.com/) and fully details the [Issue Submission](https://wasmcloud.com/docs/contributing/contributing-guide#issues) and [PR process](https://wasmcloud.com/docs/contributing/contributing-guide#pull-requests).

  They additionally outline and link to their process for feature requests which requires an [Architectural Decision Record](https://wasmcloud.com/docs/contributing/contributing-guide#proposing-an-idea), a form of design proposal.

- [x] **Demonstrate contributor activity and recruitment.**

  WasmCloud has seen significant active growth over the past 12 months with a near 60% growth in active unique contributors.

  ![image](https://hackmd.io/_uploads/ByHBzrexkl.png)
  Source:   [LFX Insights Dashboard](https://insights.lfx.linuxfoundation.org/foundation/cncf/overview/github?project=wasmcloud&routedFrom=Github&bestPractice=false&repository=)

## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

  The [wasmCloud roadmap](https://wasmcloud.com/docs/roadmap/) is discussed openly and transparently during planned quarterly open community meetings. Example: [Q3 planning meeting](https://wasmcloud.com/community/2024/07/02/community-meeting). Outside of the quarterly meetings, there is no documented roadmap process. TOC [recommends](https://github.com/wasmCloud/wasmCloud/issues/3343) that after meeting the requirements for Incubation, a process is created to document more thoroughly how features are added to the roadmap, as well as document a roadmap change process.

- [x] **History of regular, quality releases.**

  The wasmCloud/wasmCloud monorepo [had frequent releases](https://github.com/wasmCloud/wasmCloud/releases) between February 2021 and May 2021. The project was admitted to CNCF Sandbox in July 2021. Between May 2021 and August 2023, there were no releases of the project in the main repo. In August 2023, the repos were consolidated into the monorepo and releases resumed subsequently in the wasmCloud/wasmCloud monorepo. The project applied for CNCF Incubation in October 2023. Recommend that the project continues regular, quality releases that are easy to find. Any changes should be clearly documented in the RELEASE.md as well as communicated openly to the community.

### Required 

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

  The [WasmCloud website](https://wasmcloud.com/) details this well. It has various levels of detail across several sections, and also answers questions such as what it is, the problem it's trying to solve as well as several use-cases.
  - [Intro](https://wasmcloud.com/docs/intro)
  - [FAQ](https://wasmcloud.com/docs/reference/faq)
  - [Platform Overview](https://wasmcloud.com/docs/concepts/) 

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

  The [WasmCloud website](https://wasmcloud.com/) details this well.  It has various levels of detail across several sections, and also covers how it differentiates & integrates with other projects in the ecosystem.
  - [Intro](https://wasmcloud.com/docs/intro)
  - [FAQ](https://wasmcloud.com/docs/reference/faq)
  - [Platform Overview](https://wasmcloud.com/docs/concepts/)

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

  The wasmCloud project documents and maintains a [public roadmap](https://github.com/orgs/wasmCloud/projects/7) as a GitHub project and is [cross-linked to their website](https://wasmcloud.com/docs/roadmap/).

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

  The [WasmCloud website](https://wasmcloud.com/) details this well in the [Platform Overview](https://wasmcloud.com/docs/concepts/) section of the docs.


- [x] **Document the project's release process.**

  The project's release process is documented by a [RELEASE.md](https://github.com/wasmCloud/wasmCloud/blob/main/RELEASE.md) and [RELEASE_RUNBOOK.md](https://github.com/wasmCloud/wasmCloud/blob/main/RELEASE_RUNBOOK.md) documents. Recommend the project update the Release documentation with a more specific schedule.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

  Information on reporting a security issue is present in [SECURITY.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/SECURITY.md) and it is cross-linked from their website.

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  **NOTE:** 2FA is a requirement that is enforced via the CNCF GitHub Enterprise account.

  WasmCloud limits broader GitHub access to a designated "[Org Maintainer](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/CONTRIBUTION_LADDER.md#org-maintainers)" role. This role is a smaller subset of the maintainers, limiting general access.

  While not a blocker, this could be improved by adopting some tooling such as [clowarden](https://github.com/cncf/clowarden), [prow & peribolos](https://docs.prow.k8s.io/docs/components/cli-tools/peribolos/), [sheriff](https://github.com/electron/sheriff) or other tool that can manage org membership & permissions in a more transparent manner.

- [x] **Document assignment of security response roles and how reports are handled.**

  The security response roles and security process are thoroughly documented in the [SECURITY.md](https://github.com/wasmCloud/wasmCloud/blob/main/SECURITY.md) file in the wasmCloud/wasmCloud monorepo.

- [x] **Document Security Self-Assessment.**

  [WasmCloud's Security Self-Assessment](https://github.com/cncf/tag-security/blob/67039eb2f63e60d5f82911d2a6197c86a3fa9545/community/assessments/projects/wasmcloud/self-assessment.md)

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

  The wasmCloud project has achieved the [OpenSSF best practices passing badge](https://www.bestpractices.dev/en/projects/6363).

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

  Adopters are listed in [ADOPTERS.md](https://github.com/wasmCloud/wasmCloud/blob/e56ef1246475372803269be008babe7725c03733/ADOPTERS.md) at the root of the WasmCloud repo.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

  The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [x] **TOC verification of adopters.**

  Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

  The wasmCloud project maintains `[a list of integrations](https://wasmcloud.com/docs/integrations/)` on their website of both CNCF projects as well as non-CNCF projects. This meets the requirements for Incubation, however, the project should better clarify which are CNCF versus non-CNCF projects in the documentation.

#### Adoption

##### Adopter 1 - Information Technology

OCTOBER 2024

ADOPTER 1 works on a container platform team for an Information Technology services company. Anything related to cloud native projects fall under this platform team. ADOPTER 1 has been tracking WebAssembly for awhile as a technology area to keep an eye on to drive innovation within their platform and company. They had been looking at the multiple options and evaluating Fermyon, Microsoft and downstream wasmCloud from Cosmonic. For ADOPTER 1's use cases, wasmCloud was the best fit as it worked more with the existing serverless use cases that ADOPTER 1 wanted to move to WebAssembly (Wasm) - primarily due to support for Java and Javascript. The platform team also conducted a technical deep dive into the wasmCloud infrastructure and were happy with how the vision aligned with how cloud native is evolving. The team was also pleased with how the project is broken up into components and providers as well as has baked in security. Networking is also the lifeblood of ADOPTER 1's container platform and they liked how wasmCloud's lattice, dynamic, self-healing-type networking allowed for a design that was built for high availability, resiliency and sale. In addition, the project offers multi-tenancy and cross-public and -private cloud support without a degraded experience, which is incredibly important for ADOPTER 1's use cases.

ADOPTER 1 is currently a customer of Cosmonic who employs the majority of wasmCloud maintainers. The platform team had been been passively working with the wasmCloud maintainers for a year and more actively for the last 6 months. Their production applications are getting ready to deploy using Akamai and they plan to reach their edge devices with an existing CDN provider. After that, ADOPTER 1 plans to look at more deployments. The delay is not necessarily due to wasmCloud feature gaps - but Javascript compatibility in Wasm itself.

ADOPTER 1 has an existing platform based on Apache OpenWhisk - an open source serverless platform - which supports third party developers. The platform also uses AWS Lambda as well as Apache Functions. Cloud providers generally limit long running workloads, however, ADOPTER 1 does not want to artificially create a limit to various long running workloads for their end users. ADOPTER 1 could use other providers but it would likely result in costly rewrites to their applications. The current plan to address the limitations is to deprecate the Apache OpenWhisk based version of the existing platform and move it to wasmCloud.

The platform team is currently running wasmCloud version 1.14 and upgrade in place as wasmCloud releases -- which, at this time, is in an ad hoc manner in order to take advantage of testing new features and functionality. In the future, ADOPTER 1 will move to a more consistent release cadence. New releases are testing in ADOPTER 1's test/stage environment and they haven't had to roll back any releases so far. ADOPTER 1's team does have to rewrite application code in order to take advantage of new functionality and not due to any issues with compatibility.

ADOPTER 1's initial install was relatively straightforward due to the team's previous familiarity with the NATS project. The Getting Started Guide was also easy to understand and use. Typically when the team pushes software to production, component functionality breaks due to the nature of the production environment -- however, they have not seen this issue with wasmCloud. Additionally, most of the documentation has been easy to understand with the exception of understanding what comprises the wasmCloud platform itself.

At this time ADOPTER 1 has a support contract with the company sponsoring the majority of org maintainers in order to support the development of features that are useful for ADOPTER 1's use cases and to augment ADOPTER 1's upstream presence.

WasmCloud is working for ADOPTER 1 and is meeting their KPIs which includes first deploying the front end production environment. Next, they will be advocating for more internal adoption. The limiting factors to driving more adoption internally to ADOPTER 1 have been language support. As an Enterprise end user, they are primarily concerned with supporting Java and Javascript and WebAssembly itself hasn't closed all the gaps yet when it comes to supporting these languages that are typically found in larger enterprises.

Even though the wasmCloud maintainer team is primarily from a single company, ADOPTER 1 isn't as concerned about vendor lock-in though hopes moving to Incubation will prompt more vendors to support the project. They would like to see wasmCloud continue to make sure compatibility is maintained across a lot of vendors including Akamai and CloudFlare. At present, wasmCloud maintain industry standards that help lessen the risk to vendor lock-in and ADOPTER 1 would like to see upstream code continue to be ahead of - or at least in sync - with downstream code for Cosmonic's supported version of wasmCloud. ADOPTER 1 would also like to see CDN providers align with this direction, based on open standards. ADOPTER 1 maintains a pluggable backend strategy if needed to pivot to another Wasm provider if the current primary maintainer organization can no longer support the wasmCloud project and the contributor ladder does not sufficiently sustain the project. Thankfully, by the project maintaining Wasm standards, if ADOPTER 1 had to change to another cloud native project, it would be inconvenient but not a blocker to their current platform architecture and direction.

With regards to community, ADOPTER 1 has not felt any barriers to entry though there have been availability and priority limitations that prevent ADOPTER 1 from contributing more upstream. They would like to see more organizational diversity within the community, especially attracting larger vendors that can help sustain the project.

ADOPTER 1 would like to see more effort put into upstream Wasm language support - specifically Java and Javascript - in order to unlock more internal applications to move to wasmCloud. Also, the WasmCloud project has been documenting the technology and could do more to document specific use cases and benefits that adopters can see from adopting the project. Currently, the project is harder for enterprises to adopt with more legacy languages and ADOPTER 1 would like to see continued support for closing the functionality gaps between Kubernetes and their container platform.

##### Adopter 2 - Financial Services

OCTOBER 2024

ADOPTER 2 is a staff engineer of a Financial Services Institution platform team that supports a Function as a Service platform. Prior to adopting wasmCloud, the team had been building a Wasm (WebAssembly) runtime of their own and as part of that saw some issues with the Wasm runtime itself. The platform team initially explored several other open source tools even outside the wasm ecosystem– including Fission.io and Knative – and found wasmCloud to be the better fit for their use case. The team felt that by adopting wasmCloud, they would reduce the work of writing their own Wasm runtime which would allow them to work on improvement in the platform itself instead.

The platform team uses the wasmCloud project as a core component in their Function As A Service platform. They began researching the project in April 2024 and by June 2024 rolled out the project in development and it has been going well. ADOPTER 2 plans to rollout wasmCloud in production in the December 2024 timeframe. WasmCloud 1.0 is currently deployed as of the date of this interview. The team started with version 0.8 and the release cadence feels constant enough to release requested changes within a reasonable time.

ADOPTER 2 is pleased with the documentation, seeing major improvements with version 1.0 and liked that the maintainer team was open to feedback and quite welcoming. The biggest challenge seen by the ADOPTER 2 platform team is the learning curve. The team is still getting familiar with the Rust programming language in order to understand it better and allow for future contributions to the project.

ADOPTER 2 noticed that a majority of the maintainers are in the ByteCloud Alliance and have experience and context around WebAssembly in general rather just specifically the project itself. So, when the ADOPTER 2 team asks questions in the project slack channel, the community involved not only has a strong understanding of the project, but the core technologies as well. This has helped enable the platform team to adopt the project more quickly while they are coming up to speed on the newer technologies.

WasmCloud has provided value to the ADOPTER 2 platform team by simplifying deployment of their Wasm components. Also, the wasmCloud ecosystem comes with capability providers, configuration and secret management that help fill in gaps that aren’t core to Wasm runtimes in general. For example, connectors to databases as well as the HTTP server proxy. This has helped with rapid prototyping and really enables complex application designs that would be more difficult without these providers.

In the future, ADOPTER 2 sees the team getting even more involved as the team gains familiarity with the codebase and underlying technology. Overall, the community has been very open and interactive.

WebAssembly is still an early technology so there is a lot of opportunity for growth – the project seems to be doing a good job of lowering the barrier to entry with features like providers. The project also provides a seamless way of integrating Wasm components with capability providers using wrpc – Wasm component-native RPC – so the implementations are loosely coupled. ADOPTER 2 sees this as an easy way to replace a capability provider with another component in the future. The location agnostic design of wasmCloud will allow for simpler deployment and migration of Function as a Service platforms similar to the platform designed by ADOPTER 2.

##### Adopter 3 - MachineMetrics / Internet of Things

OCTOBER 2024

Jochen, Engineering Manager
Tyler, Staff Engineer

MachineMetrics operates in the Internet of Things (IoT) space, focusing on the collection and processing of data both at the edge and in the cloud. The data collected is presented via an API, enabling tool monitoring and failure prediction. Jochen serves as the engineering manager, while Tyler is a staff engineer on the platform team. The organization is comfortable sharing insights from this interview publicly.

MachineMetrics chose wasmCloud primarily due to its ability to operate under resource constraints at edge devices. This capability allows for effective integration between edge and cloud environments, utilizing existing tools like NATS. While Kubernetes (K8s) is employed in the cloud, it is not used at the edge due to resource limitations. TimescaleDB has been a core component of their stack since 2019, operating on Kubernetes. The wasmCloud operator has proven useful for deploying into existing Kubernetes setups, and MachineMetrics appreciates community contributions to this integration.

MachineMetrics has utilized WebAssembly (WASM) since version 0.8, with a proof of concept established in November of last year. The technology is now production-grade, primarily used for downsampling high-frequency data from edge devices. The portability of WASM is crucial for its adoption, enabling the development of smaller services and event-handling capabilities when issues arise. They began with version 0.82 and are currently using version 1.1, being among the early adopters of operator tooling.

Adopting wasmCloud presented challenges due to its rapidly evolving landscape and a steep learning curve associated with Rust and the WASM ecosystem. However, MachineMetrics appreciates the project's efforts to consolidate and simplify this landscape through open standards like secrets management. The community's leadership in establishing valuable standards for the WASM ecosystem has been beneficial. Despite facing documentation challenges, MachineMetrics found value in example-driven documentation and weekly community updates that helped them implement new features such as secret handling.

The implementation of wasmCloud has allowed MachineMetrics to rethink possibilities within resource-constrained edge environments. It has facilitated architectural decisions based on WebAssembly's portability, enabling developers to concentrate on business logic rather than deployment specifics. This shift helps flatten traditional silos within platform teams that have historically been divided between cloud and edge computing. Looking ahead, MachineMetrics plans to deepen its engagement with wasmCloud by developing providers tailored to their use cases that can also be generalized for broader WebAssembly applications. They envision a future where WebAssembly runs on all edge devices and is integrated across all backends.

MachineMetrics views the wasmCloud community as open and engaging, appreciating regular communication during weekly calls. They feel confident in the project's governance structure, noting that while Cosmonic backs it, there is independence that ensures project continuity even if Cosmonic faces challenges. The community shows strong growth potential as increasing adoption could influence how platform teams approach compute and data platforms. While most maintainers are from Cosmonic, there is a growing diversity among contributors without perceived barriers to participation. MachineMetrics rates their interactions with maintainers highly, describing them as proactive and supportive. They actively participate in the community through weekly calls, contributing provider capabilities, presenting at events such as WasmDay in Paris, and engaging in CNCF Live sessions.

MachineMetrics identifies several strengths within the wasmCloud project: its integration capabilities for resource-constrained edges into compute domains supported by NATS messaging technology; a strong community; and its unique position in blurring lines between cloud and edge computing. However, they acknowledge that adoption remains a challenge due to the learning curve associated with wasmCloud's component model and providers paradigm. Increasing contributions from both users and developers could enhance project growth. Encouraging strong contributors to transition into maintainer roles could also be beneficial.

While new community members are welcomed during weekly calls, MachineMetrics notes a need for improved secret handling features but appreciates the project's responsiveness to security issues as they arise. Overall, they find engagement with the community supportive and conducive to success in their implementations.
