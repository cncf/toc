# Kubescape Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1291)

## Incubation Evaluation Summary for Kubescape

### Criteria Evaluation

Emily Fox (@TheFoxAtWork) conducted the due diligence of Kubescape who applied for Incubation Level. The project has completed the criteria that show its early maturity at Incubation level. While the project meets the expectations of an early-stage Incubating level project, it could greatly benefit from a dedicated focus on its community, documentation, and overall governance.

The following actions were provided to the project that were considered blocking but since resolved:
* Completion of the security self-assessment
* Document the project's release process

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:
* Reorganization/improved structure and interlinking of the documentation, repos, and files
* Increased detail and robustness of the existing community, contributing, and governance documentation that provides better focus for potential contributors
* Provide a contributor ladder and improve the clarity around how contributors become committers and eventually maintainers. This should also include details about how and when maintainers and committers are removed from their roles
* Engage with TAG Contributor Strategy for a comprehensive Governance Review and discussion on community engagement and expansion to increase company diversity and variety of contributors to the project.
  * Complete the recommendations detailed in the governance and maintainer section (bullets and as full text) prior to initiating that engagement with TAG Contributor Strategy
* Provide tutorials that walk through and explain how the project is used and removes implicit assumptions in the Getting Started guide, recommend reaching out to new adopters and new community members/contributors for feedback and improvements that could be made
* Detail the sub-projects and repos of Kubescape to include a brief descriptor of what they are, if they are core to the project, their status, and who is responsbile for maintaining them, linking appropriately to the overall project governance

### Adoption Evaluation

The adopter interviews reflect a project in appropriate use for the level which the project applied. Interviews with Adopters commend the project's ease of use, breadth of capability, and friendliness of its maintainers as well as continually cite the comprehensiveness of its capabilities in supporting organizations looking to secure their clusters. 

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation Level.

The project could greatly benefit from several recommendations detailed throughout this Due Diligence, much of which would be addressed by reorganizing the project's website documentation, removing duplicative or out-of-date content, and establishing a dedicated area for governance, contributions, and community items for easier discoverability. It appears the structure for this has been started but not completed.

## Application Process Principles

### Suggested

N/A

### Required

- [X] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**
  - This was completed and occurred on 10-JULY-2024, and can be discovered at the [TAG's meeting notes for that session](https://docs.google.com/document/d/1XDmNG1P1YRnidQJEuZFKufaXaIenNHiY-TdYMlt0Un4/edit#heading=h.cbsth4x8m975) and the [issue link for that presentation](https://github.com/cncf/tag-security/issues/1262).

The Project presented to TAG Security on 10-JULY-2024 and provided the community with an update on new capabilities of the project, feedback by the community and the TAG was engaging and did not appear to raise any concerns with the project.

- [X] **TAG provides insight/recommendation of the project in the context of the landscape**

At the July 10th, 2024 presentation to TAG Security, Kubescape was provided the following recommendations: 

> Recommendation for the project to engage with the End User TAB for seeking feedback on these features.

> Recommendation to look at existing open source projects (SIG-infra, etc.) with substantial infrastructure (build systems for instance) to see if they are willing to test out those features.

"Features" referring to several new capabilities that allows the project to auto-generate policies/configuration that support its goals to provide a runtime-powered k8s security platform.  To provide all k8s security needs under a single project.

- [X] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The project's communications, planning, and docs align with CNCF stated guidelines for project neutrality. It is noted that the project does maintain a providers file, for which ARMO is the only one listed. However, any contributor is welcome to submit a PR to add additional providers in accordance with the information provided on the [project's docs page for providers](https://kubescape.io/docs/providers/).

- [X] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
- Met during Project's [initial application on 13-DEC-2022 to CNCF](https://github.com/cncf/toc/issues/978) and reflected through initial application to [Incubation on 13-NOV-2023](https://github.com/cncf/toc/pull/1209) and subsequent [application in the updated format on 4-APR-2024](https://github.com/cncf/toc/issues/1291)

- [ ] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisifies the Due Diligence Review criteria.

- [X] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The project provides adopters with appropriate documentation for installation and configuration of Kubescape given its maturity level, to include necessary [prerequisites](https://kubescape.io/docs/install-operator/), however it is important to note that the docs may introduce some confusion for adopters looking to get started as the [Getting Started guide](https://kubescape.io/docs/getting-started/) has different instructions than [the client](https://kubescape.io/docs/install-cli/) and [cluster install](https://kubescape.io/docs/install-operator/). 

It is recommended the project convert the Getting Started page to be more tutorial-oriented. The getting started page makes several assumptions, such as access to a machine already running a cluster. In order to more efffectively allow potential adopters to try out Kubescape, providing them a tutorial they can run locally that expands on the Overview section, and potentially suggests how the output of Kubescape can be releveraged for achieving security outcomes (i.e. continuous monitoring for compliance with CIS benchmarks) with existing dashboards, tools, and other monitoring and visibility solutions (i.e. linking to the [prometheus operator instructions](https://kubescape.io/docs/operator/prometheus-integration/) or the[ VEX document generation](https://kubescape.io/docs/operator/generating-vex/#what-is-vex)).

## Governance and Maintainers

Note: This section was completed leveraging TAG Contributor Strategy's Governance Review template and was conducted on the project's governance document from the following commit: [accc8a38347297ef41d72deef617af5a8995d22b](https://github.com/kubescape/kubescape/commit/accc8a38347297ef41d72deef617af5a8995d22b).

The following items are considered non-blocking for Incubation level but if unresolved will see the community stagnate and the project fail to meet criteria for Graduation.
* Cross link appropriate files, contributing, maintainers, and governance are all related topics and content and should be appropriately linked to answer community questions and provide a complete understanding of how the project operates and how individuals may grow, progress, and advance the project.
* The maintainer file needs updated to reflect the correct maintainers and their affiliation
* The governance should be updated to:
  * reduce duplication of content, 
  * ensure all names roles are held to the same expectations of conduct and adherence to governance, 
  * define how and when decisions are made by the maintainers, 
  * define a contributor ladder which leverages existing merit of contributions (such as committer role) for maintainership eligibility, 
  * update security-insights to correctly reflect the maintainers of the project
* Create a clear focus for community engagement
  * Ensure the meeting DAY of the week and meeting notes are linked under ##Community on the read me as well as in the contributor guide.
  * Determine appropriate scope for slack channels and actively promote that usage for those channels (within them and within the project docs/repo)
  * Define the use of “ideas” versus proposals and issues for the project, make better use of the [proposals folder structure](https://github.com/kubescape/kubescape/tree/d82e8daa25d616a8ea2ac75e6e723a29805db727/docs/proposals) and define this as part of the contributing file.

The [maintainers file](https://github.com/kubescape/kubescape/commit/74bfb57d3a135645f1c9ea9ab0e78d718dbf76f9) is complete with 5 individuals listed but is out of date as one maintainer currently works at Kaleido. The governance describes the expectations of maintainers and the requirements for consideration as maintainer. It also describes the conditions under which a maintainer is removed, this is limited to their activities on project work within a 12 month window. The governance also details a role known as committer and the method for proposing contributors to become committers. There is no contingency to be a committer prior to application as a maintainer.

The governance details how changes are proposed and how conflicts are resolved with some duplication and additional information on roles and responsibilities of maintainers, committers, and community members. Not all expectations are shared equally among the roles, notably the community members only have a call out to abide by the code of conduct and project governance. This should be updated to ensure equal expectations.

The project maintains a brief [CONTRIBUTING file](https://github.com/kubescape/kubescape/commit/33577139034d2d9d248741dd83496c20a86e204b) that details two types of contributions to appropriately guide contributors on how contributions are made, but does not reference the Ideas category of discussions on the repo as a mechanism for proposing changes. The contributions file includes a link to instructions on building and testing locally which directs users to a Wiki on the project repo and additional transitive links to getting-started docs. The contributing file does not provide any information on communications channels or community meetings and should be updated to do so.

The project does not have a contributor ladder independent of or as part of their contributing guide or governance document. It is recommended the project establish one.

The governance does not detail how decision making by the maintainers occurs nor the manner by which submitted proposals warrant decisions by maintainers. The governance should be updated to incorporate these.

The project does maintain a [security.md file](https://github.com/kubescape/kubescape/commit/accc8a38347297ef41d72deef617af5a8995d22b)  that details how security issues are reported and responded. The project also has a [security-insights file](https://github.com/kubescape/kubescape/commit/915d5d993b389c2243e0d6295c805f6a6fb5c358 ) for adopters and contributors to understand the security practices the project employs.

Meetings are recorded and uploaded to Youtube. The meeting notes link was previously buried in a comment on a Kubescape discussion and has since been added to the README for the projec and linked in the Contributor guide. The meeting notes are titled "Kubescape Developer Sync", with the regular meeting attendees typically not more than two individuals who both employed by ARMO. The meeting notes document only goes back to February of 2024 with gaps in frequency.

The project’s slack channel’s are active, notable the #kubescape primary channel is used by potential adopters and contributors to ask questions of how the project operates or any issues/bugs/challenges experienced. The Kubescape-dev channel does get reminder notifications of the upcoming community meeting but it and the meeting notes document seem more focused on the development of the project which may deter some interested contributors as not being approachable (not devs of the project ergo the meeting may not be appropriate for them to attend). This channel also overlaps in troubleshooting and project functionality discussions which occur in the Kubescape channel generally.

### Suggested

- [X] **Clear and discoverable project governance documentation.**

The Governance document is located in the primary Kubescape repo (kubescape/kubescape/GOVERNANCE.md) It does not contain links or references to other documents. For individuals looking to discovery governance of the project, it is fairly straightforward to find. The project’s website does not provide links to or reference’s the project’s governance. The file is clearly named, and does not appear to be linked from the project’s other repositories. While the current governance is simple and clear, improvements are encouraged.

- [ ] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The governance appears largely unchanged since being added.

- [ ] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

The Governance shows some small drift and deviation but is not considered substantial. Improvements are recommended to bring the current governance execution into alignment with the next maturity level beyond Incubation.

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

While the governance is not explicit in calling out neutrality, the overall structure does support vendor-neutrality.

- [ ] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

The governance does detail how the projects makes _some_ decisions on leadership but improvements could be made to ensure clarity of this process and its expectations. Contribution acceptance, requests, and changes to governance and goals should also be added.

- [ ] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

The project does not currently have this.

- [ ] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

The project does not define a maintainer lifecycle beyond activity within a 12-month window for the maintainer to remain in role.

- [ ] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Not demonstrable with the current governance.

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

The project does have several components but their leadership and governance relationship with the core project is not detailed.

### Required

- [X] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

Kubescape maintains its list of [maintainers for the core repository](https://github.com/kubescape/kubescape/blob/master/MAINTAINERS.md). This file has been changed 6 times since its creation on 06-JAN-2022 as of October 2024.

The project has the following additional respositories with Maintainer files:
* Helm-charts: https://github.com/kubescape/helm-charts/blob/main/MAINTAINERS.md
* regolibrary: https://github.com/kubescape/regolibrary/blob/master/MAINTAINERS.md - added 05-NOV-2023, no changes since. No roles defined, clear association with regolibrary

- [X] **A number of active maintainers which is appropriate to the size and scope of the project.**

The project has 5 maintainers listed on the main repo, 1 appears to be reasonably active within the past year, with additional 'committers' supporting the project. Overall activity of the repo appears to have slowed in the past year, with fewer new PRs at an increased rate of merging, indicative of increasing maturity and focus on stability of the project.

- [X] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

The project's code and doc ownership appear to match the documented governance roles for the project, however several repos under the project are lacking maintainer and contributing files or links to the overarching files for the project.

- [X] **Document agreement that project will adopt CNCF Code of Conduct.**

[The project has adopted the CNCF Code of Conduct.](https://github.com/kubescape/kubescape/commit/0eb27389daeff7dc545258cfe0f488c6b8c7a2f4)

- [X] **CNCF Code of Conduct is cross-linked from other governance documents.**

The Code of Conduct is linked in the [Contributing file](https://github.com/kubescape/kubescape/blob/1cad446fe65f470add1ec68fdc602f1f4c144740/CONTRIBUTING.md?plain=1#L16) and the [Governance file](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/GOVERNANCE.md?plain=1#L7) for the project.

- [X] **All subprojects, if any, are listed.**

The project's organization contains repos for all sub-projects, but does not have a comprehensive singular listing of each and their purpose or status.

The project should inventory its org and provide a navigation guide of its sub-projects, any maintainers of the sub-projects, the intent and focus of the project, as well as its state/status. With 55 repos, it is not immediately clear to contributors and adopters what is and is not supported. Additionally, the .github repo appears to be an initial attempt to move community and governance material to a dedicated area, highly recommend finalizing this.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

It is recommended the project establish a contributor ladder that aligns with the roles described in the governance document.

### Required

- [X] **Clearly defined and discoverable process to submit issues or changes.**

The Contributing file provides information for community members on suggesting changes but could be improved.

- [X] **Project must have, and document, at least one public communications channel for users and/or contributors.**

The project lists its slack channels on its [README under Community](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/README.md?plain=1#L98). They are also included in the Contributing file.

- [X] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

The project has listed all their communication channels. 

- [X] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

The project has regular meetings as reflected in the README and an automated notice of the meeting in their slack channel. Community members have recently requested adding the project meeting to the CNCF calendar, and the project is working towards this.

- [X] **Documentation of how to contribute, with increasing detail as the project matures.**

The project provides a simple and straightforward contributing guide that is sufficient for a project just entering incubation, however it is recommended this be developed further to be more robust and cover all aspects of contributing to the project, to include triage, PR reviews, etc. The project would benefit from linking the [third party dependency policy](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/docs/environment-dependencies-policy.md) with the Contributing file and the Governance file.

- [X] **Demonstrate contributor activity and recruitment.**

The project's slack channels are active with community members asking questions and the project's regular meetings are leveraged to serve as encouragement for more contributions. The project would greatly benefit from more active outreach to community members to take up open issues on the repo.

## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

The project does not currently have a change process for its Roadmap documented and  does not appear to use milestones or appear to associate issues or PRs with planned releases. This will be essential to the project moving forward in order to not only facillitate more clear direction for the project, but encourage more participation by potential contributors.

- [ ] **History of regular, quality releases.**

Since 2022, the project has been increasing the rigor of their release process, providing better detail post-release and no adopters have reported any issues. The current release cadence does not appear to be set yet, but is approaching some normalcy.

### Required 

- [X] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

The project's goals are detailed in their README and in the project's documentation and provides its focus within the cloud native landscape.

- [X] **Document what the project does, and why it does it - including viable cloud native use cases.**

The project's README and docs describe what the project does and why. It also includes information on how the project is to be used in cloud native environments for achieving security outcomes like runtime threat detection and compliance.

- [X] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The project maintains a [ROADMAP](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/docs/roadmap.md) but is difficult to find. It is recommended the project make this file easier to find both in the Project's website docs as well as in the repo.

- [X] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

The project maintains an older [architecture document](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/docs/architecture.md) but should be updated to more appropriately reflects the current capabilities of the project and its sub-projects to match the more recent version of the [architecture located in the project's README](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/README.md#architecture).

- [ ] **Document the project's release process.**

<!-- (TOC Evaluation goes here) --> 

## Security

Note: this section may be augemented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [X] **Clearly defined and discoverable process to report security issues.**

The project maintains a SECURITY.md file which contains the process for reporting security issues.

- [X] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

The project's maintainer accounts are protected with 2FA.

- [X] **Document assignment of security response roles and how reports are handled.**

The SECURITY.md file includes

- [ ] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

- [X] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The project maintains the [OpenSSF's best practices badge at the Passing level.](https://www.bestpractices.dev/en/projects/6944)

## Ecosystem

### Suggested

N/A

### Required

- [X] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

The project's current [ADOPTER.md](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/ADOPTERS.md) file lists the publicly known adopters. As this project provides security functionality for adopters, it is highly probable that more organizations are using the project than indicated publicly as many organizations do not disclose what tools they use to achieve security outcomes.

- [X] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [X] **TOC verification of adopters.**

Refer to the Adoption portion of this document.

- [X] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The project's architecture and [README's "under the hood" section](https://github.com/kubescape/kubescape/blob/master/README.md#under-the-hood) list other open source projects Kubescape leverages as part of its capabilities.

#### Adoption

##### Adopter 1 - Bitnami/Broadcom

_Carlos Rodriguez-Hernandez_
SEPTEMBER 2024

Binami, a catalog which provides packages of open source projects, has packaged and provided the Kubescape project for about a year and a half to its users. Committed to open source projects, Kubescape was identified by Bitnami to add a container security insights tool into their catalog that went beyond singular cloud native security capabilities.

Bitnami’s adoption and integration into their packaging and tooling pipeline for publication into the catalog was seamless and the team did not report any issues with it. It is currently available in their production catalog instance for users. Bitnami uses Kubescape as part of its security assessment and remediation process for packaged open source projects it provides in its catalog. Results of Kubescape’s scanning and reporting are provided to users alongside additional security analysis results to allow users to reason about the security posture of the packaged project.

As soon as new features and capabilities are reported in the Kubescape change log, the Bitnami team evaluates those changes to ascertain potential value for their users. When complete, the Bitnami team employs a delayed update process that allows them to update to the latest version of Kubescape after appropriate stability, testing, and integration is completed; approximately 30 days. Bitnami did not report any challenges or issues with adoption and integration of the project, by themselves or reported by their users, citing it “...was super easy to integrate in the CI/CD pipelines, the service and the tool is very easy to deploy, the output is easy to consume and leverage as an input to other tools…” in their environment.They further mentioned that in completing the proof of concept, Kubescape required no modifications and just worked out of the box.

Bitnami found the project’s documentation and repository to be well organized and easy to navigate, with robust information that did not require them to engage the project for clarification or assistance. The team has collaborated with the Kubescape project to increase awareness of Kubescape’s impact on Bitnami’s Helm charts through blogs that reflect the results of adopting Kubescape (April 2024) as well as Bitnami’s selection of Kubescape for those their catalog (July 2024).

Bitnami points out that the only thing holding the project back from more rampant adoption is the current lack of visibility and awareness to other potential adopters and plans to continue collaborating with the project to increase awareness.

##### Adopter 2 - KYOS Energy Consulting

_Zois Pagoulatos_
OCTOBER 2024

Note: This Adopter uses the Kubescape project through ARMO’s cloud security platform, a vendor offering which provides additional integrations and enrichment on top of Kubescape. This interview focuses on the benefits and experience of the underlying Kubescape project and not on capabilities exclusive to ARMO’s offering. 

KYOS Energy Consulting has been using the Kubescape project for just over a year; March 2023. As part of their cloud-native journey, KYOS began experimenting with Kubernetes to move workloads. As with more organizations early in such digital transformation, KYOS Energy Consulting was looking to scale their security expertise in this new technology area, particularly around configuration security as it continues to be one of the top weaknesses contributing to incidents. They evaluated several projects but needed a solution that was more inclusive of all aspects of Kubernetes Security, instead of point solutions limited to singular capabilities like vulnerability scanning or network policy security alone. After watching an online demo of Kubescape, KYOS Energy Consulting tried out the project in their Kubernetes environment. In the course of their integration, the only challenge cited was in the combination of Kubescape and ArgoCD. An adopter of GitOps, KYOS Energy Consulting engaged the Kubescape team to resolve the drift they were experiencing.  The team was very responsive, and the installation was smooth.

Deployed in production, KYOS Energy Consulting has had positive engagement with the maintainers, deep diving into resource utilization challenges in the cluster, unique and custom applications and technology resulting in interesting notifications, and general debugging.

Through their experience, KYOS Energy Consulting has been quite happy with their choice of Kubescape, citing the project provides nearly all their needs for security in Kubernetes, continuously adding features and capabilities they didn't know they needed, contributing positively to not only the security of their clusters but also allowing their engineers to gain confidence in their adoption of cloud-native technologies and in particular the value of Kubernetes. Kubescape has reduced the volume of false positives, and the network policy suggestions are extremely helpful. The team reviews suggestions from Kubescape prior to applying them to confirm their accuracy and viability in their environment which has relieved pressure on the engineering team in developing and managing such policies. KYOS Energy Consulting further commented on their satisfaction in the integrity of the project, as a small team they plan to continue using the project even if they weren't leveraging the vendor offering. They currently contribute to the Kubescape project unofficially, and are working on transitioning from another cloud provider. They’ve already added new clusters to Kubescape and are excited at the upcoming features under development by the project.


##### Adopter 3 - Anonymous Finance

_Adopter and organization are anonymous_
OCTOBER 2024

Adopter 3 has been using the Kubescape project for approximately 3 years. The Interviewee had familiarity and experience with using Kubescape at their previous employer and had been satisfied with its ability to address the security of their clusters, making introduction of the project to Adopter 3 a clear choice. 

Adopter 3’s motivation in adopting Kubescape was to address the need of shifting operational cluster security left in the development lifecycle, allowing the team to test and verify the security of their clusters before they become operational in production. Their workflow includes a staggered deployment pipeline that leverages Kubescape to scan clusters in a non-prod environment, test and verify the recommendation changes based on scan results from Kubescape, and then deploy those changes in their production environment. This workflow reduces the likelihood that a recommended change from Kubescape regarding their cluster security posture will adversely impact their production workloads, allowing them to catch problematic changes before the impact is realized.

Each time Adopter 3’s pipeline runs, it checks for and pulls the most recent version of Kubescape available, this ensures they are getting updated vulnerability information as part of their workflow. Adopter 3 uses the Kubescape Helm charts for installation, maintaining their Kubescape Manifests in a repo for installation as part of their pipeline. They also execute push based scans in their pipeline, finding the overall experience of installation and integration of Kubescape as simple as Grafana.

Adopter 3 uses the Kubescape core feature set (posture management, configuration, compliance, hardening) in all their environments. They are not currently using Kubescape to provide runtime security monitoring due to their workflow to reduce the risk of production issues. Adopter 3 is leveraging the application profile rule set in a limited capacity at this time, fully testing it before developing plans to roll this out further.

The Kubescape project was early for its time, and when adopted by Adopter 3, was the only tool with clear community support and provided the most complete feature set for managing cluster security which addressed Adopter 3’ s requirements. Adopter 3’s use of Kubescape has provided immeasurable value, as a multi-cloud and multi-cluster entity, Adopter 3 loves the agnostic support of cloud service providers that Kubescape provides, particularly capacity to support their GitOps flow via Flux. Of particular value to Adopter 3, is the speed at which they are able to deploy new, compliant clusters into their production environments. They have shared clusters, production and non-production per region with node pool isolation which leverage canary deploys to roll out compliant clusters. While they have other tools they must actively troubleshoot, their use of Kubescape with Helm and Flux just works.

While Kubescape does not fully encompass needs for multi-layer security at this time, its capacity to facilitate posture management and misconfiguration was noted by Adopter 3 as one of the best. Adopter 3 believes their use of the Kubescape project has contributed meaningfully to their current “zero breach streak”. 

It has been always clear what the project does as the documentation and repo is very complete and well organized. Adopter 3 has a dedicated communication channel with the project to facilitate quick fixes, such as Helm chart backwards compatibility. Adopter 3 has been particularly pleased with the separation between the Vendor which contributed the project to CNCF (Armo), it’s enterprise offering, and the Kubescape community – while separate, there is focus on alignment of capabilities which allows Adopter 3 to continue using the open source project with optionality to pursue an enterprise offering in the future that will work with their existing architecture should they need it.

Over the years of using the Kubescape project, Adopter 3 has developed a friendship with the project, engaging in discussion regarding features and providing feedback to the project on areas of improvement or capabilities and has participated in a presentation regarding functionality of Kubescape at a previous CNCF conference. 

Adopter 3 believes the Kubescape project could do more with an expanded scope but did not cite any particular features they wanted from the project, having an existing positive engagement to field those with the project. They did note the potential for the project to become a fully featured CNAPP solution, a goal they believe Kubescape is pursuing.

Regarding their future plans with Kubescape, Adopter 3 is looking into a way they could create application profiles in their non-production environment; providing coverage of runtime security. Their current path is to build once and deploy many so each deployment matches, this will likely require them to keep the eBPF component in their non-production environment and apply the created profiles in production without having to execute the learning in production environments. Adopter 3 noted the additional overhead of alert management in production will require focus on false-positive noise suppression as one of several reasons to keep this within their non-production environments.
