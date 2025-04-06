# OpenYurt Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/pull/797)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for OpenYurt

### Criteria Evaluation

[Ricardo Rocha](https://github.com/rochaporto) conducted both due diligence and adopter interviews for Incubation. The project has completed the criteria that shows its maturity at the applied level.

The TOC would like to highlight some areas from the project review:
* OpenYurt has an active and diverse community, and adopters have highlighted this as one of its main strengths
* Its architecture extends instead of replacing existing core components in Kubernetes. This is an approach that differentiates it from other solutions in the ecosystem and that was also highlighted as a key strength by adopters
* The governance and contributors guides are well defined but were last updated over 3 years ago. The TOC suggests a review together with TAG Contributor Strategy to consider if improvements can be made to the current member roles and contributor ladder

The following required incubation items were raised and implemented by the project during the review period:
* A security self assessment following the template from TAG-Security
* Setting up of the best practices badge from the OpenSSF

Other items suggested for incubation were taken as issues in the project and have also been followed up:
* Documenting integrations with other CNCF and non-CNCF projects to facilitate adoption by end users
* Documenting the process for definition of the project roadmap
* Clarification of public contact for security vulnerability reports

### Adoption Evaluation

The adopter interviews showed project usage at a level compatible with CNCF incubation. Adopters integrate OpenYurt for management of edge devices on a variety of use cases, and highlight its extensibility of upstream vanilla Kubernetes as one of its key strengths - some adopters highlighted this is a differentiating point when compared to similar projects in the ecosystem which do the same by replacing some of the core components in Kubernetes. The project has a well established and responsive community in a variety of communication channels.

Adopters also suggested improvements in resource consumption, multi-tenancy support and recommend the project focuses on edge AI use cases looking forward.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

  The project [has presented](https://www.youtube.com/watch?v=L6X3Bx0e7B0) at TAG-Runtime, but given the significant evolution of both project and the ecosystem a new presentation is to be scheduled with an update.

- [x] **TAG provides insight/recommendation of the project in the context of the landscape**

OpenYurt is a Kubernetes-native project designed to seamlessly extend Kubernetes capabilities to cloud-edge environments, addressing critical needs in IoT, industrial automation, and other edge-intensive domains. Its architecture employs CRDs, NodePools for regional edge node management, and local caching mechanisms to ensure reliability and offline functionality in low-connectivity scenarios. With tools like Helm for installation and yurtadm for managing edge components, OpenYurt provides a user-friendly, modular approach to deploying Kubernetes across distributed edge regions. Its alignment with CNCF principles is evident in its extensibility, cloud-native design, and focus on enabling Kubernetes to manage diverse workloads across geographically dispersed environments.

During CNCF incubation, OpenYurt must provide evidence of API stability, scalability benchmarks, and fault tolerance in production-grade deployments. Expanding community engagement, demonstrating governance maturity, and showcasing broader adoption across industries will further validate its readiness. Additionally, improved documentation, including comparisons with similar projects like KubeEdge and an articulated roadmap for the next 12–24 months, will solidify its position as a leading solution for cloud-edge orchestration. With these enhancements, OpenYurt is well-positioned to drive innovation in distributed Kubernetes architectures and make a meaningful impact in the CNCF ecosystem.

The CNCF TAG-Runtime recommends OpenYurt for incubation, recognizing its strong alignment with cloud native principles and its potential to address critical challenges in cloud-edge orchestration. With its Kubernetes-native design, extensibility, and growing community, OpenYurt is well-positioned to impact the CNCF ecosystem as it matures significantly.

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

  No issues were found during due diligence, both code and documentation are vendor neutral.

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		

  Done during the project's application, onboarding completed [Mar 2, 2021](https://github.com/cncf/sandbox/issues/270).

- [x] **Due Diligence Review.**

  The initial application was initially submitted Mar 07, 2022. After some work required on the project's side the TOC due diligence kick-off happened May 06, 2024.

  The project submitted an initial DD doc with information, [available here (google docs)](https://docs.google.com/document/d/1St_JIS6Ntt1q39eOyUYs7U5enXIiYhVY1-6qb6b-6Rs/edit?tab=t.0).

  Project maintainers were responsive throughout the process.

  The process satisfies the due diligence review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

  OpenYurt's [documentation](https://openyurt.io/docs/) is extensive and covers architecture, concepts, installation and configuration. It also includes the required information for [development](https://openyurt.io/docs/developer-manuals/how-to-build-and-test) and [contribution](https://openyurt.io/docs/developer-manuals/how-to-contribute).

## Governance and Maintainers

### Suggested

- [x] **Clear and discoverable project governance documentation.**

  OpenYurt's governance is documented under the project's github repo, under [GOVERNANCE.md](https://github.com/openyurtio/openyurt/blob/master/GOVERNANCE.md). Some aspects, such as the [different roles](https://github.com/openyurtio/community/blob/main/community-membership.md) of community membership are particularly well documented.

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

  The last update has happened 3 years ago and it is recommended that a review is performed together with TAG-CS.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

  Governance is up to date with actual practices.

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  Although there is no mention of vendor in the project governance documentation, it states its commitment to public decision making.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  The decision making process is simple and [documented](https://github.com/openyurtio/openyurt/blob/master/GOVERNANCE.md#decision-making-process), focusing on consensus and highlighting 3rd party involvement for conflict resolution.

  Given the size of the project this is reasonable.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

  Community membership and roles are [well documented](https://github.com/openyurtio/community/blob/main/community-membership.md) with details on how to grow to different roles in the project.

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

  Roles are [well documented](https://github.com/openyurtio/community/blob/main/community-membership.md) including dealing with inactive members.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

  The current list of maintainers is [well maintained](https://github.com/openyurtio/community/blob/main/OWNERS.md#maintainers) and includes the list of [emeritus members](https://github.com/openyurtio/community/blob/main/OWNERS.md#emeritus-members).

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

  N/A

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

  The full list of maintainers is well maintained in a [MAINTAINERS.md](https://github.com/openyurtio/openyurt/blob/master/MAINTAINERS.md) file along with their affiliation. The content is a subset of a more extensive list of reviewers, approvers and maintainers in a [OWNERS.md](https://github.com/openyurtio/community/blob/main/OWNERS.md) which also includes emeritus members.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

  The list of maintainers is extensive with a diverse affiliation. As of November 2024, it lists 9 maintainers from 7 organizations.

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

  A OWNERS file in the main repo matches the community file listing approvers and reviewers.

- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

  The code of conduct is [linked](https://github.com/openyurtio/openyurt/blob/master/CODE_OF_CONDUCT.md) and is listed as a requirement from the [contributing guide](https://github.com/openyurtio/openyurt/blob/master/CONTRIBUTING.md).

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

  A specific [CODE_OF_CONDUCT.md](https://github.com/openyurtio/openyurt/blob/master/CODE_OF_CONDUCT.md) links to the CNCF code of conduct.

- [x] **All subprojects, if any, are listed.**

  N/A

## Contributors and Community

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

  The [community membership](https://github.com/openyurtio/community/blob/main/community-membership.md) documents the different roles in the project.

  It documents well the responsibilities of each role, along with the requirements to move along those roles, both in terms of sponsor support and contributions. The process for emeritus/inactive community members is also well documented.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

  The [contributor guide](https://github.com/openyurtio/openyurt/blob/master/CONTRIBUTING.md) details how to report generic and security issues, and how to contribute code and documentation.

  It also documents expected best practices for commits, references to issues and the required workspace preparation.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

  The primary communication channel is DingTalk, with a mailing list and a slack channel also linked in the [contact information](https://github.com/openyurtio/openyurt/tree/master?tab=readme-ov-file#contact).

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

  - [Mailing list](https://groups.google.com/g/openyurt/)
  - [Slack](https://join.slack.com/t/openyurt/shared_invite/zt-2ajsy47br-jl~zjumRsCAE~BlPRRsIvg): for english speakers
  - [DingTalk](12640034121): for chinese speakers

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

  Meetings are APAC friendly and linked in the CNCF calendar, with [recordings available](https://space.bilibili.com/484245424/video) in bilibili.com and up to date as of November 2024. 

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

  The [contributors guide](https://github.com/openyurtio/openyurt/blob/master/CONTRIBUTING.md) is extensive and linked in the project's [home page](https://openyurt.io/docs/developer-manuals/how-to-contribute).

- [x] **Demonstrate contributor activity and recruitment.**

  The project has seen a reduction in the total number of contributors in the last 12 months, but still sees over 20 monthly contributors, according to [LFX Insights](https://insights.lfx.linuxfoundation.org/foundation/cncf/overview/github?project=openyurt&repository=&routedFrom=Github). Of those, 4 are responsible for ~50% of all contributions with over 120 contributors total in the last 12 months.

  Regarding organizations, 1 organization is responsible for 90% of the contributions with the remaining work spread over multiple other organizations.

## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

  A [roadmap file](https://github.com/openyurtio/community/blob/main/roadmap/roadmap-2024.md) exists with yearly updates. No procedure is available describing how the roadmap is created or discussed, [an item](https://github.com/openyurtio/openyurt/issues/2188) to be addressed by the project.

- [x] **History of regular, quality releases.**

  OpenYurt has a [history of releases](https://github.com/openyurtio/openyurt/releases) with extensive and complete release notes for the community. While claiming 2 to 3 months, the project seems to currently follow a 6 month release cycle in practice, with multiple patch releases in between.

### Required 

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

  The [OpenYurt website](https://openyurt.io/) has the required information and details.

  The [Getting Started/Introduction](https://openyurt.io/docs/) explicitly says what differentiates the project from others in the ecosystem. A section on core concepts expands on its [architecture](https://openyurt.io/docs/core-concepts/architecture) and components.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

  OpenYurt has an extensive set of user manuals - example for [storage](https://openyurt.io/docs/user-manuals/storage/edge-local-storage) - covering the use for specific areas and use cases. The project also documents its [resource requirements](https://openyurt.io/docs/usage-conditions/resource-and-system-requirements).

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

  A [roadmap](https://github.com/openyurtio/community/blob/main/roadmap/roadmap-2024.md) exists in the community's github repository.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

  The project's [architecture](https://openyurt.io/docs/core-concepts/architecture) page covers this topic in detail. The user manuals in the same website further expand on the use for each of the target end users.

- [x] **Document the project's release process.**

  The release cadence and [process](https://github.com/openyurtio/openyurt/blob/master/RELEASE-PROCESS.md) are well documented in the github repository.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

  The information is documented in a [SECURITY.md](https://github.com/openyurtio/openyurt/blob/master/SECURITY.md) file and linked from the contributing guide in github. Following [this issue](https://github.com/openyurtio/openyurt/issues/2187) consolidation was done to show the same info also the website.

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  Access control is available via CODEOWNERS after [this issue](https://github.com/openyurtio/openyurt/issues/2223), with the initial groups:
  * OpenYurt-Maintainers

  Improvements to this setup, with finer grained groups and ACLs is being worked on [here](https://github.com/openyurtio/openyurt/issues/2225).

- [x] **Document assignment of security response roles and how reports are handled.**

  Handling of security reports is documented in the [SECURITY.md](https://github.com/openyurtio/openyurt/blob/master/SECURITY.md) file, relying on reports to maintainers. 

  The target email is now using the openyurt domain following [this issue](https://github.com/openyurtio/openyurt/issues/2186).

- [x] **Document Security Self-Assessment.**

  Issue requesting self assessment [opened here](https://github.com/openyurtio/openyurt/issues/2184). Completed.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

  Issue requesting the OpenSSF badge [opened here](https://github.com/openyurtio/openyurt/issues/2185). Implemented.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

  A list of [ADOPTERS.md](https://github.com/openyurtio/community/blob/main/ADOPTERS.md) is available in the community repo.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

  The list of adopters (see above) documents over 20 production deployments as of November 2024.

  Also see adopter interviews.

- [x] **TOC verification of adopters.**

  See the *adoption section* in this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

  The [architecture](https://github.com/openyurtio/openyurt?tab=readme-ov-file#architecture) section documents the dependencies and integrations with Kubernetes.

  [An issue](https://github.com/openyurtio/openyurt/issues/2189) has been opened to complement this with the information on the integrations with other CNCF projects. This was done and is [available here](https://openyurt.io/docs/integrations/).


#### Adoption

##### Adopter 1 - Information Technology

Adopter 1 is a senior software engineer working in the area of edge computing.

###### Usage

Adopter 1 reports interest in the project started in 2020 with adoption in 2021. OpenYurt is integrated with multiple solutions at the company which are then packaged and made available to end users.

Adopter 1 reported integration with the company's existing edge computing platform, filling the role of managing edge devices. Examples of usage include switching off edge devices during the night and bringing them back in the morning; AI on the edge for object recognition, including car plate recognition at parking lots and similar usage in retail and restaurants.

###### Perception and Feeling

Adopter 1 participates actively in the development of the project and has dedicated developers for this task. Those developers are active in regular meetings.

Interaction with the project happens mostly via GitHub issues, with TikTok and WeChat also used to interact with the community.

###### Strengths and Improvements

Adopter 1 seeds the added networking stability to standard kubernetes as one of the main strengths, particularly important for edge workloads. The additional primitive for managing `nodepools` is also a valuable addition to the standard `daemonset` and `deployment`.

Adopter 1 sees installation as the main pain point and sees the same happening with other end users. No single installation method is ideal, leading to distinct methods depending on the use case.

Adopter 1 sees improvements can come with better support for AI workloads, especially for inference but also including the usage of remote devices for distributed training. Adopter 1 sees this as an opportunity for OpenYurt to pioneer this type of workload.

##### Adopter 2 - Telecommunications

Adopter 2 is a project manager in the cloud computing area, working at the innovation center of a large telecommunications company.

###### Usage

Adopter 2 has been using OpenYurt since 2023, performing regular updates twice a month. The project is fully integrated with their kubernetes facilities, allowing a single installation method for third party applications focusing on validation and certification of edge devices. Adopter 2 relies on OpenYurt to analyse videos from cameras

Adopter 2 considered other solutions, having decided to use OpenYurt as it adds functionality on top instead of changing Kubernetes.

###### Perception and Feeling

Adopter 2 is an active participant in community meetings for the project every 2 weeks. They also contribute code back to the community, fixing issues with static analysis tooling.

Adopter 2 mainly communicates with the community via WeChat, but also uses GitHub and the regular community meetings.

###### Strengths and Improvements

Adopter 2 sees the automated development capabilities as a way of greatly simplifying development of edge node applications. They also see the simplified operations as a main strength, with simplified installation methods for those remote devices in particular for third party applications.

Adopter 2 would like to see better multi-tenancy support, including tenant separation.

##### Adopter 3 - Financial and Industrial Services

Adopter 3 is an engineer working on Kubernetes for over 4 years and responsible for cloud edge products for over 2 years, in close contact with the OpenYurt project. Adopter 3 is responsible for financial and industrial products.

###### Usage

Adopter 3 chose OpenYurt as it extends existing vanilla kubernetes deployments, allowing for an edge solution that is compatible with existing clusters. This gave a solution that is less intrusive than other projects such as KubeEdge.

Adopter 3 runs OpenYurt for production workloads in multiple deployments, including a major electricity provider in China as well as smaller providers around the country. This includes the monitoring of 1000s of electricity sub-stations, with a setup where OpenYurt is deployed along KubeVirt. This solutions has been in place for almost 2 years.

Adopter 3 follows the upstream releases closely.

Adopter 3 is looking to extending its usage of the project to improve connectivity between edge devices, as well as improved isolation and better network stability. They expect this will open usage to other use cases, including automotive and drones.

Adopter 3's main contributions to the project have been in the area of bandwidth usage optimization and reduction of packet loss, as well as security enhancements in the area of encryption.

###### Perception and Feeling

Adopter 3 mainly communicates with the project via DingTalk. They can see maintainers review open issues weekly, with discussions on both issues and roadmap.

Adopter 3 has been invited multiple times to present their use cases, and follows presentations from other users as well. They see it as a very active community.

###### Strengths and Improvements

Adopter 3 recently made a comparison with KubeEdge. This comparison showed OpenYurt has a higher overall resource consumption likely to the existing of explicit nodes on the edge as it follows the native kubernetes deployments. Reducing this numbers would open up new use cases where resource usage is particularly important.

Adopter 3 recommends the project follows this particular topic looking forward.

##### Adopter 4 - Research

Adopter 4 is a software engineer at a research laboratory and currently a OpenYurt maintainer.

###### Usage

Adopter 4 relies on OpenYurt for research work in the area of scheduling optimization on edge environments. Adopter 4 focuses in particular on the integration with other projects such as Istio, and adds features to the project where appropriate and needed.

Adopter 4 also works together with other end users helping them integrate the tools in their environments.

###### Perception and Feeling

Adopter 4 mainly relies on DingTalk but also Slack to communicate with the project's community. Adopter 4 also actively participates helping other end users. They see very good interactions with other maintainers in the project, where GitHub is the main communication channel.

Adopter 4 joins the weekly meetings when required or when the topic is interesting.

###### Strengths and Improvements

Adopter 4 also has members working on KubeEdge. OpenYurt's main attraction is being less intrusive to the upstream kubernetes components, giving the example that the upstream kubelet is used without any modification.

Adopter 4 thinks the project could become more popular and should focus on attracting new users. It also recommends focusing on further integration with AI tools in particular EdgeAI modules for inference.

Adopter 4 also suggest network usage should be a key focus area, reducing the traffic in the underlay network to minimize the impact on slow network deployments.
