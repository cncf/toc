# KServe Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1367)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for KServe

### Criteria Evaluation

Kevin Wang and Faseela K conducted the due diligence of KServe who applied for incubation. The project has completed the criteria that show its maturity at incubation.

#### Noteworthy Implementations

The KServe team has demonstrated substantial progress across CNCF incubation criteria:

- **Open and Transparent Governance**:
  - KServe has established a clear and well-documented governance model in its [Technical Charter](https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md), which includes a strong commitment to vendor neutrality.
  - The project's maintainer list, contributor ladder, and community roles are all publicly documented ([MAINTAINERS.md](https://github.com/kserve/community/blob/main/MAINTAINERS.md), [ROLES.md](https://github.com/kserve/community/blob/main/ROLES.md), [membership.md](https://github.com/kserve/community/blob/main/membership.md)), ensuring transparency and accessibility.

- **Strong Community Engagement**:
  - The project has a long-standing history of community engagement, with [bi-weekly meetings](https://github.com/kserve/community/tree/main?tab=readme-ov-file#community-meeting) running for over five years.
  - Contributor promotion and activity are openly tracked and demonstrated through public pull requests (e.g., [#3484](https://github.com/kserve/kserve/pull/3484), [#3966](https://github.com/kserve/kserve/pull/3966)), fostering an active and welcoming community.

- **Robust Engineering and Security Practices**:
  - KServe maintains a public [roadmap](https://github.com/kserve/kserve/blob/master/ROADMAP.md) that includes strategic initiatives like GenAI inference support and enhanced security.
  - The project follows a quarterly release schedule with a documented [release process](https://github.com/kserve/kserve/blob/master/release/RELEASE_PROCESS_v2.md).
  - The project has strengthened its security posture by achieving an OpenSSF Best Practices badge ([badge](https://www.bestpractices.dev/en/projects/6643)), implementing regular image scans, and completing a [security self-assessment](https://github.com/kserve/community/blob/main/security/self-assessment.md).

- **Broad Industry Adoption**:
  - KServe has achieved widespread adoption across various industries, with production deployments at major companies such as CyberAgent, Nutanix, Cloudera, and SAP.
  - The project has also improved its adopter documentation in response to feedback from the TOC Reviewers ([community#47](https://github.com/kserve/community/issues/47)).

#### Blockers (Previously Raised, Now Resolved)

- **Standardized Code of Conduct Reference and Management**: The project initially had inconsistent Code of Conduct files. This was resolved by creating a standard `CODE_OF_CONDUCT.md` file across all repositories, ensuring uniformity, ref: <https://github.com/kserve/community/issues/29>.
- **Improved Adopter Documentation**: In response to TOC reviewers' feedback, the project expanded its adopter documentation to provide clearer references and greater visibility into production use cases,
ref: <https://github.com/kserve/community/issues/47>.
- **Improved Security Self-Assessment**: The security self-assessment was updated to clarify details regarding SBOMs and vulnerability scanning, providing a more comprehensive overview of the project's sec
urity practices, ref: <https://github.com/kserve/community/issues/44>.

#### Recommended Enhancements (Non-Blocking)

The following items remain open may not block incubation, but should be completed and demonstrated before graduation.

- **Governance Decision Process**: Explicit documentation of leadership, contribution acceptance, and governance changes ([community#53](https://github.com/kserve/community/issues/53)).  
- **Maintainer Lifecycle**: Full lifecycle (onboarding, offboarding, emeritus status) not yet fully documented ([community#52](https://github.com/kserve/community/issues/52)).  
- **Subproject Governance**: Leadership and maturity status for subprojects not documented ([community#53](https://github.com/kserve/community/issues/53)).  
- **Security Reporting Links**: `SECURITY.md` not yet linked from community repo readme ([community#42](https://github.com/kserve/community/issues/42)).  
- **Roadmap Change Process**: Governance should specify how roadmap changes are proposed and approved (gap in [charter](https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md)).  
- **Release Documentation**: Clarify release frequency, support policy, and rename release doc for consistency ([community#41](https://github.com/kserve/community/issues/41)).  
- **Security Contact Migration**: Move vulnerability reporting address from LF AI & Data to CNCF domain in future ([community#42](https://github.com/kserve/community/issues/42)).  
- **CI & Testing Resources**: Address limited CI infrastructure, potentially via CNCF resources.  
- **Documentation Consolidation**: Reduce fragmentation between website and GitHub sources (raised in adopter interview feedback).  

### Adoption Evaluation

KServe has demonstrated consistent and growing adoption across diverse industries and geographies, with production deployments ranging from large-scale multi-cloud enterprise platforms to specialized internal AI infrastructure. The project is used by major organizations including Cloudera, CyberAgent, Nutanix, SAP, and others, spanning sectors such as enterprise software, cloud infrastructure, online media, gaming, and financial services, with deployments supporting both generative and predictive AI workloads at scale. Use cases span multi-cloud enterprise AI platforms, large GPU clusters, internal developer platforms, and hyperscaler-agnostic services, serving thousands of tenants in some cases.

Adopters highlighted KServe’s Kubernetes-native design, flexible serving runtime mechanism, broad ML framework support, and advanced scaling features (autoscaling, zero-scale) as key benefits. These capabilities have enabled significant operational efficiency and cost savings—CyberAgent reported ~40% savings on GPU costs—while maintaining production stability and smooth upgrade paths.

The project is backed by a diverse set of active contributors from multiple companies, with open governance and strong community support. Areas for improvement noted by adopters include consolidating documentation, simplifying dependencies on Knative/Istio, increasing maintainer coverage across time zones, and enhancing outreach to showcase KServe’s strengths and adoption stories.

Overall, KServe is a mature, production-proven model serving platform with broad applicability, strong community backing, and clear potential for continued growth under CNCF incubation.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

  <!-- (TOC Evaluation goes here) -->  
  The KServe maintainers has provided GTR file at: <https://github.com/kserve/community/blob/main/docs/general-technical-review-cncf-incubating.md>

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

  <!-- (TOC Evaluation goes here) -->  
  Yes. KServe is utilizing vendor neutral resources for communication, testing, hosting and governance.
  Some examples are:
  - KServe Github: <https://github.com/kserve/kserve>
  - KServe Slack: <https://cloud-native.slack.com/archives/C06AH2C3K8B>
  - Testing - KServe CI Configuration: <https://github.com/kserve/kserve/tree/master/.github/workflows>
  - KServe Communigy Governance: <https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md>
  Ref: <https://github.com/kserve/community/issues/26#issuecomment-2866903229>

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**

  <!-- (TOC Evaluation goes here) -->  
  - The project contacts and TOC Reviewers had a kick-off meeting on Feb. 18th, set expectations and discussed general steps & timelines.

- [x] **Due Diligence Review.**

  The initial [application](https://github.com/cncf/toc/issues/1365) was submitted on July 6th, 2024.

  The TOC reviewers started working the due diligence on Feb. 5th, 2025, and met with the project maintainers multiple times during the DD evaluation.

  The project maintainers have been highly responsive throughout the process.

  Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisifies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

  <!-- (TOC Evaluation goes here) -->
  Refer to the following links:
  - Installation: <https://kserve.github.io/website/latest/admin>
  - End user documentation: <https://kserve.github.io/website/master/modelserving/control_plane/>
  - Example Blog post for a new release: <https://kserve.github.io/website/0.15/blog/articles/2025-05-27-KServe-0.15-release/x>

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Clear and discoverable project governance documentation.**

  <!-- (TOC Evaluation goes here) -->
  KServe project documents its open govermance at [KSERVE-TECHICAL-CHARTER.md](https://github.com/kserve/community/blob/main/KSERVE-TECHICAL-CHARTER.md)

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

  <!-- (TOC Evaluation goes here) -->
  KServe project Governance has iterated organically as it has gained experienced over the years.  
  Relevant details worth mentioning are:
  - The roof governance document is tracked at: <https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md>.
  - Contributor Roles and contributor ladder processes are streamlined over the past years as well as team member onboarding/offboarding process is well defined at: <https://github.com/kserve/community/blob/main/membership.md>.
  - [KServe Technical Steering Committee](https://github.com/kserve/community/blob/main/TECHNICAL-STEERING-COMMITTEE.md) was established in Apr. 2025. <https://github.com/kserve/community/pull/18>
  - The Project currently has 19 [maintainers](https://github.com/kserve/community/blob/main/MAINTAINERS.md) and according to [lfx insights](https://insights.linuxfoundation.org/project/kserve/contributors?timeRange=past365days&start=2024-08-14&end=2025-08-14&widget=active-contributors) it has received contributions from 499 contributors who come from different companies.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

  <!-- (TOC Evaluation goes here) -->
  Governance is up to date and KServe project has been running [bi-weekly community meeting](https://github.com/kserve/community/tree/main?tab=readme-ov-file#community-meeting) for 5 years now.  
  KServe also regularly promote contributors with the established voting and approval process. Some examples are:
  - <https://github.com/kserve/community/issues/21>
  - <https://github.com/kserve/kserve/pull/3484>
  - <https://github.com/kserve/kserve/pull/3213>
  - <https://github.com/kserve/kserve/pull/3966>

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  <!-- (TOC Evaluation goes here) -->
  [KSERVE-TECHNICAL-CHARTER.md](https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md#5-vendor-neutrality) explicilty documents vendor neutrality of project direction.

  The TOC reviewers suggest adding rules such as "a limitation on the number/percentage of seats that a single vendor/company could take," "the total number of seats for TSC members," and "the term of office for the TSC" into the governance. These rules should be demonstrated before applying for graduation. This is being tracked under [Provide more detailed TSC charter](https://github.com/kserve/community/issues/52)

- [ ] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  <!-- (TOC Evaluation goes here) -->
  The KServe project currently uses GitHub issues for discussion and follows a git-vote process to make final decisions. For example: [requests to the CNCF](https://github.com/kserve/community/issues/10).

  While these practices are established and functional, they have not yet been formally documented as written governance rules. The TOC reviewers recommend documenting explicit decision-making processes rather than relying solely on examples. This requirement is being tracked under [Document how the KServe project makes decisions](https://github.com/kserve/community/issues/53).

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

  <!-- (TOC Evaluation goes here) -->
  KServe documents it's community roles, associated responsibilities, and role management at <https://github.com/kserve/community/blob/main/ROLES.md>

- [ ] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

  <!-- (TOC Evaluation goes here) -->
  KServe has documented Maintainer lifecycle relevant processes in two documents: <https://github.com/kserve/community/blob/main/membership.md> and <https://github.com/kserve/community/blob/main/ROLES.md>.
  
  As decribed in <https://github.com/kserve/community/issues/39>, the TOC reviewers recommend consolidating all community role definitions into "membership.md" for better organization and clarity.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

  <!-- (TOC Evaluation goes here) -->
  - Updates to the KServe OWNER file can be found at: <https://github.com/kserve/kserve/commits/master/OWNERS>
  
  - Some of the good examples are:
    - Add new PR reviewers and approvers: <https://github.com/kserve/kserve/pull/3213>
    - Update list of reviewers according to activities: <https://github.com/kserve/kserve/pull/3484>
    - PR that adds the promoted maintainers: <https://github.com/kserve/kserve/pull/3966>

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

  <!-- (TOC Evaluation goes here) -->
  The maintainers listed the following 2 subprojects, however no descriptions or rules about subprojects were found in the [KSERVE-TECHNICAL-CHARTER.md]:
  - KServe ModelMesh:  Yuan Tang, [terrytangyuan@gmail.com](mailto:terrytangyuan@gmail.com)
  - KServe Open Inference Protocol: Dan Sun [dsun20@bloomberg.net](mailto:dsun20@bloomberg.net)

  The TOC reviewers suggest documenting subproject leadership, contribution, maturity status, including add/remove process. This is being tracked under [Document how KServe project makes decisions](https://github.com/kserve/community/issues/53)

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

  <!-- (TOC Evaluation goes here) -->
  KServe has a clear [KServe maintainers.md](https://github.com/kserve/community/blob/main/MAINTAINERS.md) file for all the maintainer relevant information.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

  <!-- (TOC Evaluation goes here) -->
  The project's activity level shows a reasonable distribution of maintainers. Of the top 10 contributors over the past 180 days, 7 are active maintainers.  
  Refer to the "Contributor Leaderboard" in [LFX insights - KServe](https://insights.linuxfoundation.org/project/kserve/contributors?widget=contributors-leaderboard&timeRange=past180days&start=2025-02-27&end=2025-08-26) for more details.


- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

  <!-- (TOC Evaluation goes here) -->
  KServe uses the OWNERS file to manage code and doc approval access among approvers and reviewers, and clearly matches documented governance roles.
  - Code owners:
    - KServe: <https://github.com/kserve/kserve/blob/master/OWNERS>
    - modelmesh serving: <https://github.com/kserve/modelmesh-serving/blob/main/OWNERS>
    - models web app: <https://github.com/kserve/models-web-app/blob/master/OWNERS>
  - Doc owners:
    - website: <https://github.com/kserve/website/blob/main/OWNERS>

- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

  <!-- (TOC Evaluation goes here) -->
  KServe has documented their adoption of CNCF Code of Conduct, in [CONTRIBUTING.md#code-of-conduct](https://github.com/kserve/community/blob/main/CONTRIBUTING.md#code-of-conduct) and [KSERVE-TECHNICAL-CHARTER.md#4-compliance-with-policies](https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md#4-compliance-with-policies)

  The Kserve community has also created a `CODE_OF_CONDUCT.md` file at the root level of each repository.
  Ref: <https://github.com/kserve/community/issues/29>

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

  <!-- (TOC Evaluation goes here) -->
  Yes, it's cross-linked from:
  - <https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md>
  - <https://github.com/kserve/community/blob/main/ROLES.md>
  - <https://github.com/kserve/community/blob/main/docs/SUBPROJECTS.md>

- [x] **All subprojects, if any, are listed.**

  <!-- (TOC Evaluation goes here) -->
  According to [SUBPROJECTS.md](https://github.com/kserve/community/blob/main/docs/SUBPROJECTS.md), KServe currently has the following subprojects:
  - ModelMesh: <https://github.com/kserve/modelmesh-serving>
  - Open Inference Protocol: <https://github.com/kserve/open-inference-protocol>

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

  <!-- (TOC Evaluation goes here) -->  
  The project describes its contributor ladder in [membership.md](https://github.com/kserve/community/blob/main/membership.md). Contributor roles include "member", "reviewer", "approver" and "leader".
  Requirements and recruiting/retiring process are clearly documented.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

  <!-- (TOC Evaluation goes here) -->  
  The project describes issue and PR process in its [Contribution Guide](https://github.com/kserve/community/blob/main/CONTRIBUTING.md).

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

  <!-- (TOC Evaluation goes here) -->  
  Communication channels are documented at <https://github.com/kserve/community?tab=readme-ov-file#questions-and-issues>

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

  <!-- (TOC Evaluation goes here) -->
  KServe community channels can be found at <https://github.com/kserve/community#questions-and-issues>.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

 <!-- (TOC Evaluation goes here) -->  
  KServe uses [LFX calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/kserve?view=month) for running community meeting.
  The [readme of kserve/community repo](https://github.com/kserve/community#community-meeting) clearly documents meeting time, meeting link, community calendar, and meeting notes:

  > You can also find these meetings on [the community calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/kserve?view=month), along with other major community events.
  >
  > Meeting agendas and notes can be accessed in the [working group document](https://docs.google.com/document/d/1KZUURwr9MnHXqHA08TFbfVbM8EAJSJjmaMhnvstvi-k).
  >
  > You can access the meeting recordings on [the community calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/kserve?view=month) by clicking on the respective date's event details.

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

  <!-- (TOC Evaluation goes here) -->  
  KServe has a cleary contribute guide at <https://github.com/kserve/community/blob/main/CONTRIBUTING.md>

- [x] **Demonstrate contributor activity and recruitment.**

  <!-- (TOC Evaluation goes here) -->
  The project [contributor guide](https://github.com/kserve/community/blob/main/CONTRIBUTING.md#project-membership) describes an activity based project membership recruitment.  
  Two examples of using the contributor activity data are:  
  - Update list of reviewers according to activities: <https://github.com/kserve/kserve/pull/3484>
  - PR that adds the promoted maintainers: <https://github.com/kserve/kserve/pull/3966>

## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

  <!-- (TOC Evaluation goes here) -->
  The KServe 2024-2025 roadmap is captured at <https://github.com/kserve/kserve/blob/master/ROADMAP.md>. However, the project's [Governance document](https://github.com/kserve/community/blob/main/KSERVE-TECHNICAL-CHARTER.md) does not cover how the roadmap is created or how changes to it are handled.

  While this is not a requirement for incubation, the TOC reviewers recommend that the project document and demonstrate this process before applying for graduation.

- [x] **History of regular, quality releases.**

  <!-- (TOC Evaluation goes here) -->
  KServe follows quarterly releases as per <https://github.com/kserve/kserve/releases> and <https://github.com/kserve/kserve/tags>.
  
  Even though the release process from a development perspective is captured under: <https://github.com/kserve/kserve/blob/master/release/RELEASE_PROCESS_v2.md>, it is recommended to document at a high level the frequency/type of releases and support windows also.

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

  <!-- (TOC Evaluation goes here) -->
  KServe project goals, objectives and how it fits in the Cloud Native landscape are well captured under <https://github.com/kserve/kserve?tab=readme-ov-file#why-kserve>

  > **Why KServe?**
  > - KServe is a standard, cloud agnostic **Model Inference Platform** for serving predictive and generative AI models on Kubernetes, built for highly scalable use cases.
  > - Provides performant, **standardized inference protocol** across ML frameworks including OpenAI specification for generative models.
  > - Support modern **serverless inference workload** with **request based autoscaling including scale-to-zero** on **CPU and GPU**.
  > - Provides **high scalability, density packing and intelligent routing** using **ModelMesh**.
  > - **Simple and pluggable production serving** for **inference**, **pre/post processing**, **monitoring** and **explainability**.
  > - Advanced deployments for **canary rollout**, **pipeline**, **ensembles** with **InferenceGraph**.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

  <!-- (TOC Evaluation goes here) -->
  The high level usage of the project is captured at <https://github.com/kserve/kserve?tab=readme-ov-file#kserve>.  
  The project website also contains good documentation: <https://kserve.github.io/website/latest/> on usage, and administration.  
  The project also has done a good job in capturing various KServe usages, demos and presentations at <https://kserve.github.io/website/master/community/presentations>.  
  The project also lists adopters at <https://kserve.github.io/website/latest/community/adopters>.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

  <!-- (TOC Evaluation goes here) -->
  KServe has documented their 2024-2025 roadmap in <https://github.com/kserve/kserve/blob/master/ROADMAP.md> and the major highlights include:  
  1. Support GenAI inference
  2. Graduate core inference capability to stable/GA
  3. Graduate KServe Python SDK to 1.0
  4. Graduate InferenceGraph
  5. Secure InferenceService
  6. Improve KServe 1.0 documentation

  Feature Request process is also documented clearly at: [community/CONTRIBUTING.md at main · kserve/community](https://github.com/kserve/community/blob/main/CONTRIBUTING.md#contributing-a-feature).
  
  However roadmap document is present only for 2024-2025 and not previous years. Hoping that 2024-2025 trend continues, and roadmaps are similarly published for upcoming years too to indicate a healthy project. Similarly, it is recommended to mention the roadmap change process in the Governance document.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

  <!-- (TOC Evaluation goes here) -->
  KServe website has detailed information on the project architecture and design at the User Guide section. For example: <https://kserve.github.io/website/latest/modelserving/control_plane> and <https://kserve.github.io/website/latest/modelserving/data_plane/data_plane>.

- [x] **Document the project's release process.**

  <!-- (TOC Evaluation goes here) -->
  The release process for KServe is documented in [`RELEASE_PROCESS_v2.md`](https://github.com/kserve/kserve/blob/master/release/RELEASE_PROCESS_v2.md).

  The TOC reviewers recommend several improvements to this documentation:
  - **Rename the file** to `RELEASE.md` to avoid confusion, as the `_v2` suffix is misleading.
  - **Add a user-focused section** at the beginning of the document that outlines:
    - Release cadence
    - Supported versions and types (e.g., Beta)
    - Support policy and End-of-Life (EOL) announcements
  - **Clarify the release roles**, specifying that the process is performed by authorized release managers who are chosen for each release (ideally defined in the governance document).

  These recommendations are being tracked under [Document history of regular, quality releases](https://github.com/kserve/community/issues/41).

## Security

Note: this section may be augemented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

  <!-- (TOC Evaluation goes here) -->
  KServe’s security vulnerability reporting process is outlined at [SECURITY.md](https://github.com/kserve/kserve/blob/master/SECURITY.md).  
  The mail id is currently under lfaidata foundation, it is recommended to move to a CNCF domain at an appropriate time.  
  As part of the TOC review process the SECURITY.md has been updated with detailed information on how to report security vulnerabilities and the file has been linked to all major repositories of KServe.  
  See <https://github.com/kserve/community/issues/42> and <https://github.com/kserve/community/issues/43> for more details.

  It's also recommended to link SECURITY.md to <https://github.com/kserve/community?tab=readme-ov-file#questions-and-issues>

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  <!-- (TOC Evaluation goes here) -->
  Access control is enforced via OWNERS and CODEOWNERS files:
  - <https://github.com/kserve/kserve/blob/master/OWNERS> for Kserve repo
  - <https://github.com/kserve/community/blob/main/.github/CODEOWNERS> for community repo

  And the respective roles are specified in <https://github.com/kserve/community/blob/main/membership.md#community-membership>.  

  Members of KServe organization are required to enable two factor authentication <https://github.com/kserve/community/blob/main/membership.md#requirements> in their Github account.

- [x] **Document assignment of security response roles and how reports are handled.**

  <!-- (TOC Evaluation goes here) -->
  <https://github.com/kserve/kserve/blob/master/SECURITY.md> specifies that only active maintainers of the project has access to reported security vulnerabilities. SECURITY.md has detailed information on KServer supported versions, Prevention Mechanisms, Private/Public Disclosure Process and time taken for mitigation, and Communication channels used for the disclosure.

  <https://github.com/kserve/community/blob/main/ROLES.md> defines responcibilities of [KServe Security Team](https://github.com/kserve/community/blob/main/ROLES.md#kserve-security-team-kserve-security) and [ModelMesh Security Team](https://github.com/kserve/community/blob/main/ROLES.md#modelmesh-security-team-modelmesh-security). The assignment of security response roles is documented as the follwoing:

  Ref: <https://github.com/kserve/community/blob/main/ROLES.md#role-progression>:
  > Contributor → Security Team: Based on sustained contributions to the CVE fixes and security related improvements

- [x] **Document Security Self-Assessment.**

  <!-- (TOC Evaluation goes here) -->
  KServe Security Self Assessment is available at <https://github.com/kserve/community/blob/main/security/self-assessment.md>. In addition:
  - Improvements were done as part of <https://github.com/kserve/community/issues/44> to improve the self-assessment documentation to clarify about SBOMs.
  - KServe snyk scan running twice a week, ref: <https://github.com/kserve/kserve/blob/master/.github/workflows/scheduled-image-scan.yml>.
  - There is also a scan for go source code as well. Ref: <https://github.com/kserve/kserve/blob/master/.github/workflows/scheduled-go-security-scan.yml>.
  - <https://github.com/kserve/community/blob/main/security/self-assessment.md#development-pipeline> has detailed information regarding the same.

  It's recommended that the links given in <https://github.com/kserve/kserve/blob/master/CONTRIBUTING.md> just point to <https://github.com/kserve/community/blob/main/CONTRIBUTING.md> to make is easier to access.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

  <!-- (TOC Evaluation goes here) -->
  KServe has obtained OpenSSF Best Practicess passing badge: <https://www.bestpractices.dev/en/projects/6643>

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

  <!-- (TOC Evaluation goes here) -->
  KServe adopters are tracked at: <https://kserve.github.io/website/docs/community/adopters>.  
  The project has done a great job collecting various KServe demos and presentations at: <https://kserve.github.io/website/docs/community/presentations>.  
  KServe did additional work as per the TOC review under <https://github.com/kserve/community/issues/47> to enhance <https://kserve.github.io/website/docs/community/adopters> with the right references from various demos and presentations.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

  <!-- (TOC Evaluation goes here) -->
  The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

  The adopters document about recorded 9 production grade adopters, and the TOC reviewers have interviewed 4 qulified adopters that are using KServe in production.

- [x] **TOC verification of adopters.**

  <!-- (TOC Evaluation goes here) -->
  Refer to the Adoption portion of this document for more details.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

  <!-- (TOC Evaluation goes here) -->
  KServe officially supports integration with CNCF projects Istio, Knative, gRPC, OpenTelemetry. KServe installation guide using Istio/Knative has been described at: <https://kserve.github.io/website/latest/admin/serverless/serverless>.  
  KServe ServiceMesh guide using Istio is documented at: <https://kserve.github.io/website/latest/admin/serverless/servicemesh>.  
  Based on TOC reviewers' recommendation, KServe team is working on additional enhancements as part of <https://github.com/kserve/community/issues/46> to incorporate other integrations from projects like Kubeflow, minio etc.

#### Adoption

##### Adopter 1 - CyberAgent, Inc. (Internet Services & Entertainment)

  CyberAgent, Inc., a leading internet and entertainment company in Japan, operates across media, advertising, and gaming. To support its diverse AI-driven business operations, CyberAgent built Cycloud, a private cloud platform with a dedicated bare-metal GPU cluster in its own Japanese data center. This cluster, hosting over 80 nodes and ~500 GPUs, powers between 1,500 and 2,500 pods, delivering GPU-as-a-Service, distributed training, and a global inference endpoint. KServe serves as the core model serving solution within Cycloud, enabling the company to offer flexible, scalable inference services across business units.

  CyberAgent adopted KServe four years ago after evaluating multiple architectures, including a DIY approach and a Knative + Istio stack. KServe was chosen for its built-in autoscaling and zero-scale features, which optimize costly GPU usage, as well as for its deep integration with Knative and Istio that reduced operational complexity. Today, KServe supports ~80 inference services, each consuming between 1 and 8 GPUs during peak usage. Scaling is demand-driven, and zero-scale ensures idle services release resources to maximize efficiency.

  Operationally, CyberAgent upgrades KServe annually, currently running v0.14.0 after eight successful upgrades since v0.6. Adoption was smooth thanks to strong documentation and examples, though production troubleshooting proved challenging due to limited deep-dive technical guides. Despite this, KServe’s templating mechanism (serving runtime) allowed for seamless support of multiple ML frameworks and custom containers. The platform achieved ~40% cost savings versus alternatives, and was built by just four developers in six months—highlighting KServe’s efficiency.

  CyberAgent’s future plans include active community participation and feature requests such as Gateway API Inference Extension, zero-scale for raw deployment mode, vertical pod autoscaling, and Dynamic Resource Allocation support. While praising KServe’s governance, maintainer responsiveness, and strong technical foundations, CyberAgent noted opportunities to improve technical decision transparency and global collaboration through more asynchronous communication. KServe remains a critical component in enabling CyberAgent’s large-scale, cost-efficient AI infrastructure.

  August 2025
  Reference: [KServer Adopter Interview - CyberAgent](https://github.com/kevin-wangzefeng/toc/blob/kserve-incubation-wip/projects/kserve/kserve-adopter-interview-cyberagent.md)

##### Adopter 2 - Nutanix/Cloud Service Provider

  A global leader in cloud software, Nutanix has integrated KServe into its production cloud platform since 2021, following extensive experimentation from the project’s early stages. Their choice of KServe over alternative solutions highlights the platform’s maturity, driven by its Kubernetes-native design, flexible and rich feature set, and robust support for both predictive and generative AI model inference.

  KServe’s standard inference APIs abstract away the complexity of underlying ML frameworks, making it accessible for production use. Nutanix has specifically called out features such as accelerator support and autoscaling as key differentiators that align with their enterprise-grade requirements.

  In addition to technical capabilities, Nutanix reports that the KServe community is active, supportive, and well-governed with a strong commitment to vendor neutrality. They commend the project’s high-quality documentation, which has facilitated seamless adoption and upgrades. Each new KServe release is thoroughly tested by Nutanix internally, and backward compatibility is consistently upheld—demonstrating the project's engineering maturity.

  Nutanix is also an active contributor to the KServe project, with five engineers engaged, including three official reviewers. They have expressed a strong commitment to continue contributing and to deepen their collaboration with other CNCF ecosystem projects such as Envoy and components within the model inference domain. This level of production usage, community involvement, and sustained commitment from a major enterprise vendor like Nutanix reflects positively on KServe’s readiness for CNCF incubation.

  July 2025
  Reference: [KServer Adopter Interview - Nutanix](https://github.com/kevin-wangzefeng/toc/blob/kserve-incubation-wip/projects/kserve/kserve-adopter-interview-nutanix.md)

##### Adopter 3 -  Cloudera (Enterprise Data Platform Provider)

  Cloudera has integrated KServe as the model serving backbone of its AI Inference product within the Cloudera Data Platform (CDP). The decision followed a comprehensive evaluation of build-versus-buy options in 2022–2023, during which KServe was selected for its open governance under LF AI & Data and CNCF, enterprise-friendly licensing, active and approachable community, and vendor-neutral roadmap. KServe’s Kubernetes-native architecture, built on Knative Serving, aligned with Cloudera’s cloud-native design principles and avoided duplicating existing infrastructure capabilities.

  KServe has been in use for nearly two years, with approximately one year of production deployment since the product’s GA release in 2024. Current production usage spans multiple industries, including banking, government, healthcare, and oil & gas, with customers running both generative AI (GenAI) and traditional machine learning models. Cloudera primarily leverages KServe’s customizable cluster serving runtimes, traffic management, and scaling features, enabling rapid onboarding of diverse model types across hybrid and multi-cloud environments.

  The integration experience was smooth due to KServe’s compatibility with existing Kubernetes-based infrastructure. Documentation quality was cited as a strength, with the project’s requirement for synchronized code and documentation updates ensuring accuracy and completeness. KServe has delivered measurable value by accelerating time-to-market (estimated savings of at least two years in development effort), enabling multi-cloud operations, and supporting rapid adoption of emerging AI capabilities. Cloudera actively contributes upstream, including code, documentation, and new features such as enhanced inference logging with cloud storage backends.

  Cloudera views KServe’s community governance, technical quality, and growth potential positively, though it notes challenges such as a small maintainer pool, slow PR review cycles, and constrained CI infrastructure. Future plans include deepened community involvement, governance participation, and continued upstreaming of customer-driven features to avoid maintaining internal forks. Cloudera believes CNCF incubation could help address resource constraints, improve build/test infrastructure, and expand maintainer diversity across time zones.

  July 2025
  Reference: [KServer Adopter Interview - Cloudera](https://github.com/kevin-wangzefeng/toc/blob/kserve-incubation-wip/projects/kserve/kserve-adopter-interview-cloudera.md)

##### Adopter 4 – SAP (Enterprise AI Platform Provider)

  SAP AI Core, a multi-cloud AI execution platform within SAP’s Business Technology Platform, selected KServe in mid-2020 for its Kubernetes-native, cloud-agnostic serving capabilities. The platform placed a premium on seamless integration with their Kubernetes infrastructure, multi-tenancy, and scalability—criteria that KServe met more effectively compared to alternative solutions, particularly given its Kubeflow heritage and alignment with CNCF best practices.

  In production since 2021, SAP AI Core currently employs KServe across thousands of internal and external tenants. The service leverages KServe’s InferenceService API to abstract model serving logistics, autoscale workloads, and interact with multi-cloud object storage. Critical features include serverless scale-to-zero for cost-efficient edge cases, as well as storage initializers supporting cross-cloud object stores, accelerating deployment and operational agility.

  While the overall integration experience has been positive, SAP identified operational complexities arising from KServe’s dependencies—most notably Knative and Istio—during upgrades. Documentation is robust but fragmented between the website and GitHub repository, prompting a recommendation for consolidation. Nonetheless, KServe’s multi-cloud support, alignment with enterprise governance, and accelerated time-to-market have brought tangible value and reduced infrastructure burdens for SAP’s AI services.

  SAP continues to contribute to KServe with active development and feature requests, with multiple engineers participating in the project. The community is praised for its responsiveness and diversity under CNCF governance, though meeting times favor US contributors. SAP looks forward to deeper involvement, improved adopter visibility, and reduced operational complexity in future project iterations.

  August 2025
  Reference: [KServer Adopter Interview - SAP](https://github.com/kevin-wangzefeng/toc/blob/kserve-incubation-wip/projects/kserve/kserve-adopter-interview-sap.md)
