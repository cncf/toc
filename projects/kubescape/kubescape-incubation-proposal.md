# Kubescape Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1291)

## Incubation Evaluation Summary for Kubescape

### Criteria Evaluation

Emily Fox (@TheFoxAtWork) conducted the due diligence of Kubescape who applied for Incubation Level. The project has completed the criteria that show its early maturity at Incubation level. While the project meets the expectations of an early-stage Incubating level project, it could greatly benefit from a dedicated focus on its community, documentation, and overall governance.

The following was notable for the project and worth calling out here:
* Project engaged the TOC sponsor for DD to ensure they were aware of the expectations for Incubation level projects and could begin taking steps to positions itself for successful graduation in the future.
* During this DD the TOC sponsor found several areas of the project that while meeting the minimum criteria, did not fully meet the intent of the criteria. The project worked to improved their implementation in several areas and focused to improve its governance, community, and project documentation. The project is in a much better place as a result of leaning-into these changes and if continued, should foster a thriving, vibrant and open community.

The following actions were provided to the project that were considered blocking but since resolved:
* Complete the project's Security Self-Assessment
 * Completed with [PR #1424](https://github.com/cncf/tag-security/pull/1424)
* Ensure the maintainers file is up to date and reflects the current affiliation of the maintainers
 * Fully completed with [Commit 5a0b97a](https://github.com/kubescape/project-governance/commit/5a0b97a374b630483e123dc8a37d787f8e556b63)
* Document the project's release process
 * Completed with [Commit d938389](https://github.com/kubescape/project-governance/commit/d938389e639db99544e0cc43ecd30d1a5dc47e28)

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:
* Reorganization/improved structure and interlinking of the documentation, repos, and files
  * Captured for work and completed through the [project's pre-incubation board](https://github.com/orgs/kubescape/projects/7/views/1)
* Improve the contributing guidelines to better explain how contributors can get involved and understand expectations
 * Completed with [PR #23](https://github.com/kubescape/project-governance/pull/23)
* Improve the project's Roadmap and cadence of planning to support project activities, like planning for moving levels, such that these boards are referenceable and understandable in the [Roadmap document](https://github.com/kubescape/project-governance/blob/main/ROADMAP.md) such that contributors can see what work is currently underway versus planned later.
* Ensure the Kubescape site is updated to reflect the current capabilities of the project and provides adopters with clear information on getting started and engaging with the project.
 * Completed with [PR #79](https://github.com/kubescape/kubescape.io/pull/79/files)
* Clarify information on how contributors can become maintainers
 * The [updated governance file](https://github.com/kubescape/project-governance/blob/main/GOVERNANCE.md) for the project provides information on decision making and how maintainers are added to the project.
* Provide a high-level architecture diagram of each of the components, and how they work together, with links to their repos be provided on the Kubescape website and in the project's primary or project-governance repos to improve discovery. There are numerous architecture diagrams identified during the DD that cover various aspects of the project.
* The [Security-Insights file has also been updated](https://github.com/kubescape/kubescape/commit/e87bf7b723fdad67037f0db741777cb1b721d50f) to reflect the listing of maintainers
* Ensure the [project's OpenSSF Scorecard accurately reflects the security posture of the project's repositories](https://scorecard.dev/viewer/?uri=github.com%2Fkubescape%2Fkubescape) and for items that are incomplete, work to address them. Security projects are expected to be the exemplar of good security practices and security hygiene - to include ensuring existing reporting tools correctly reflect their posture
 * Make additional improvements to its repository security posture by enforcing branch protection, ensuring code-reviews are performed on pull requests and changes are not committed directly to the project without review from a second individual who is not the author of the commit.
 * It is noted that the best practices badging in the scorecard is inaccurately recorded as there appears to be an additional record being pulled.
* Provide a contributor ladder and improve the clarity around how contributors become committers and eventually maintainers. This should also include details about how and when maintainers and committers are removed from their roles
* Engage with TAG Contributor Strategy for a comprehensive Governance Review and discussion on community engagement and expansion to increase company diversity and variety of contributors to the project.

### Adoption Evaluation

The adopter interviews reflect a project in appropriate use for the level which the project applied. Interviews with Adopters commend the project's ease of use, breadth of capability, and friendliness of its maintainers as well as continually cite the comprehensiveness of its capabilities in supporting organizations looking to secure their clusters. 

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation Level.

The Kubescape project is considered a very early Incubation level project which meets the minimum expectations and criteria, with a few areas of increasing maturity. The project could greatly benefit from more comprehensive documentation and activity captures/messages of what specifically is transpiring, completed, and why. The current operations of the project indicate a closely knit group of maintainers which often operates with implicit knowledge and understanding that would provide substantial value to potential community members not in the group if properly captured in the repo as well-linked documents, as PRs linked and referenced in open issues which demonstrate clear, transparent, and reasoned paths to resolution: problem --> implementation --> closure, and include additional comments and decisions that make assumptions and project knowledge explicit to the uninformed community member.

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

- [X] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [X] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The project provides adopters with appropriate documentation for installation and configuration of Kubescape given its maturity level, to include necessary [prerequisites](https://kubescape.io/docs/install-operator/). 

The project recently made changes to their getting started page to provide greater clarity to individuals unfamiliar with the project and the documentation also includes three tutorials to assist individuals in becoming familiar with Kubescape, how to harden a cluster, and using kubescape to enforce security policies on github repositories.

## Governance and Maintainers

Note: This section was completed leveraging TAG Contributor Strategy's Governance Review template and was conducted on the project's governance document from the following commit: [accc8a38347297ef41d72deef617af5a8995d22b](https://github.com/kubescape/kubescape/commit/accc8a38347297ef41d72deef617af5a8995d22b).

The following items are considered non-blocking for Incubation level but if unresolved will see the community stagnate and the project fail to meet criteria for Graduation.
* Define a contributor ladder which leverages existing merit of contributions (such as committer role) for maintainership eligibility, 
* Create a clear focus for community engagement
  * Define the use of “ideas” versus proposals and issues for the project, make better use of the [proposals folder structure](https://github.com/kubescape/kubescape/tree/d82e8daa25d616a8ea2ac75e6e723a29805db727/docs/proposals) and define this as part of the contributing file.

### Suggested

- [X] **Clear and discoverable project governance documentation.**

The project's governance file is linked from the main project repository to a dedicated project-governance repo.  It contains sufficient information on how the project is governed for an early stage incubating project. The project could improve the current governance by linking their official resolutions and Maintainer Meeting notes in the governance so they are more readily discoverable (note they only exist in the [community.md file](https://github.com/kubescape/project-governance/blob/fda8a8122d1a0c0133d5bf216cae6ec092adb290/COMMUNITY.md?plain=1#L27)). It is also recommended the project evaluate existing governance templates (such as the [maintainer council template](https://contribute.cncf.io/maintainers/templates/governance-maintainer/)) to improve how the project is governed and formally develop a [contributor ladder](https://github.com/cncf/project-template/blob/main/CONTRIBUTOR_LADDER.md).

- [ ] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The governance appears largely unchanged since being added.

- [X] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

The governance file was updated to reflect quarterly Maintainer Meetings (of which only 1 meeting is recorded in the notes), and the current Maintainer Meeting notes references discoverable activities that show the application of its governance (i.e. [proposals for new maintainers](https://github.com/kubescape/project-governance/issues?q=label%3Aproject-candidacy+is%3Aclosed))

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

While the governance is not explicit in calling out neutrality, the overall structure does support vendor-neutrality.

- [ ] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

The [governance file](https://github.com/kubescape/project-governance/blob/main/GOVERNANCE.md) partially covers these items, specifically calling out how new maintainers are approved and how changes are proposed to the project and its governance.

- [ ] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

The project does not currently have this.

- [ ] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

The project does not define a maintainer lifecycle beyond activity within a 12-month window for the maintainer to remain in role.

- [ ] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Not demonstrable with the current governance.

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

The project does have several components which fall under the overall project governance but the project does not include any add/remove process for these components.

### Required

- [X] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

Kubescape maintains its list of maintainers for the core repository at https://github.com/kubescape/kubescape/blob/master/MAINTAINERS.md. This file has been changed 8 times since its creation on 06-JAN-2022.

- [X] **A number of active maintainers which is appropriate to the size and scope of the project.**

The project's maintainer listing was refreshed with review dates. The project's level of activity show a reasonable distribution of 2 active maintainers and 1 active committer (per the project's governance documentation).

- [X] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

In November 2024, the project [reviewed their access controls and made changes](https://github.com/kubescape/project-governance/issues/14) to ensure ownership and access is carried by two primary groups. The project's code and doc ownership appear to match the documented governance roles for the project.

- [X] **Document agreement that project will adopt CNCF Code of Conduct.**

[The project has adopted the CNCF Code of Conduct.](https://github.com/kubescape/kubescape/commit/0eb27389daeff7dc545258cfe0f488c6b8c7a2f4) The project also has included an [additional modifier that is NOT in conflict with the CNCF CoC](https://github.com/kubescape/project-governance/blob/main/CODE_OF_CONDUCT.md)

- [X] **CNCF Code of Conduct is cross-linked from other governance documents.**

The project's code of conduct, which links to CNCF's, is cross-linked through several project documents, to include the [governance file](https://github.com/kubescape/project-governance/blob/fda8a8122d1a0c0133d5bf216cae6ec092adb290/GOVERNANCE.md?plain=1#L9), the [README](https://github.com/kubescape/project-governance/blob/fda8a8122d1a0c0133d5bf216cae6ec092adb290/README.md?plain=1#L24) for project-governance, and the [contributing file](https://github.com/kubescape/project-governance/blob/fda8a8122d1a0c0133d5bf216cae6ec092adb290/CONTRIBUTING.md?plain=1#L20)

- [X] **All subprojects, if any, are listed.**

The Kubescape project does not have any known subprojects. The project does [maintain a listing of repositories in its org that are subject to its governance](https://github.com/kubescape/project-governance/blob/main/repositories-under-governance.json). These repositories contain components necessary for interoperability with other cloud native projects, installation tools, actions, and other elements essential to the project.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

While the project defines several roles for contributors, additional enhancements would greatly improve understanding and actionability for how an individual may progressively become more involved in the project through the defined roles.

### Required

- [X] **Clearly defined and discoverable process to submit issues or changes.**

The project's contributing file contains a variety of information on how to get started, however the [Community file](https://github.com/kubescape/project-governance/blob/main/COMMUNITY.md#github-issues) and [Governance file](https://github.com/kubescape/project-governance/blob/main/GOVERNANCE.md#processes) contain specific information on opening issues and proposing changes to the project.

- [X] **Project must have, and document, at least one public communications channel for users and/or contributors.**

The project's [README](https://github.com/kubescape/kubescape?tab=readme-ov-file#community) and [project-governance/community file](https://github.com/kubescape/project-governance/blob/main/COMMUNITY.md) contain links to both the users and developers channel for the project. The project could benefit by linking the [kubescape-dev slack channel on the contributing file](https://github.com/kubescape/project-governance/blob/fda8a8122d1a0c0133d5bf216cae6ec092adb290/CONTRIBUTING.md?plain=1#L68).

- [X] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

The project has listed and documented all communications channels in their relevant files of the repo. See the previous criteria evaluation for more information. There are no known non-public project communication channels.

- [X] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

The project regularly holds public meetings, with [public notes](https://github.com/kubescape/project-governance/blob/main/COMMUNITY.md#community-meetings) and [recordings](https://www.youtube.com/channel/UCgf472-CiLXkfeoPWBmh5hg). Information on meeting schedule is [available in the community file](https://www.youtube.com/channel/UCgf472-CiLXkfeoPWBmh5hg).

- [X] **Documentation of how to contribute, with increasing detail as the project matures.**

The project provides a simple and straightforward contributing guide that is sufficient for a project just entering incubation, however it is recommended this be developed further to be more robust and cover all aspects of contributing to the project, to include triage, PR reviews, etc. The project would benefit from linking the [third party dependency policy](https://github.com/kubescape/kubescape/blob/d82e8daa25d616a8ea2ac75e6e723a29805db727/docs/environment-dependencies-policy.md) with the Contributing file and the Governance file.

The project's original contributing file had been modified 14 times since it was initially committed. After migrating the content to the project-governance/contributing file, it has been modified an additional 3 times. Through each iteration, the robustness and clarity in detail has been improved.

- [X] **Demonstrate contributor activity and recruitment.**

In the past year, the project's contributor activity has decreased by approximately 31% (per LFX Insights query dtd 19 DEC 2024). Further analysis into the project shows fewer new contributors since June 2024, which show drift indicative of slowed activity and recruitment - likely the result of the project's current focus on stability, performance, and scaling improvements as these activities require more indepth understanding of the project's current capabilities. Looking into 8knot, the Kubescape project  receives the vast majority of contributions through repeat contributors, with first time contributors and contributor activity increasing annually in October before steadily decreasing again towards July on an annual basis.


It is worth noting that the project calls out new contributors in addition to all contributors that contributed to a given release, with the [3.0.19 release (12 NOV 2024)](https://github.com/kubescape/kubescape/releases/tag/v3.0.19) being the most recent with a new contributor to the project.


## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

The project does not currently have a change process for its Roadmap documented and  does not appear to use milestones or appear to associate issues or PRs with planned releases. This will be essential to the project moving forward in order to not only facillitate more clear direction for the project, but encourage more participation by potential contributors.

- [ ] **History of regular, quality releases.**

Since 2022, the project has been increasing the rigor of their release process, providing better detail post-release and no adopters have reported any issues. The current release cadence does not appear to be fixed yet, with documentation citing 6-8 weeks between releases and the release history showing more frequent releases than this.

### Required 

- [X] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

The project's goals and objectives are documented through it's website: _... offers hardening, posture management, and runtime security to ensure robust protection for Kubernetes environments. The project's website also describes additional open source projects in use by Kubescape and how they are used within the project to for specific use cases (i.e. image scanning with Grype, OPA for security control posture verification, etc.).

- [X] **Document what the project does, and why it does it - including viable cloud native use cases.**

The project's website provides adopters and contributors with information about aspects of the project, what they do, and why. An overview for each of project's components is located on each readme and provides information on how it is used to achieve a targeted cloud native security outcome (image scanning, triggering actions from events, etc.)

- [X] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The project has [a primary documented public roadmap](https://github.com/kubescape/project-governance/blob/main/ROADMAP.md) and an [older roadmap](https://github.com/kubescape/kubescape/blob/e1db7f3704199441724a85c16786bbf85ec76948/docs/roadmap.md) that needs replaced or provide reference to the current one ([captured in issue #1777](https://github.com/kubescape/kubescape/issues/1777)). The Project has linked the current Roadmap in their README to improve discovery of it, and has created several project boards that show forward looking planning. It is recommended the project establish a practice to review the roadmap at least annually and leverage project boards to track defined goals and objectives to completion.

- [X] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

The project has two locations of architecture diagrams that demonstrate how components of the Kubescape work together and lower level detail how each work. The first architecture file is in the [kubescape/kubescape repo](https://github.com/kubescape/kubescape/blob/master/docs/architecture.md) and is two years old. It covers the general component architecture, the Kubescape CLI with the OPA engine, and the Operator. The second architecture diagram is a series of diagrams located in the [helm-charts for the kubescape-operator's readme's section on in-cluster components](https://github.com/kubescape/helm-charts/blob/main/charts/kubescape-operator/README.md#in-cluster-components-overview) and is far more expansive on content specific to the operator (the primary method to run Kubescape as a set of microservices inside a kubernetes cluster for cluster compliance and vulnerability management). An overview of the Kubescape architecture which references the operator, storage, and other components in the context of a cluster provides additional information. It is worth noting that [additional diagrams on the CLI and Operator are stored in the security self-assessment for the project](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/kubescape/self-assessment.md#actors).

It is recommended the project ensure that a high-level architecture diagram of each of the components, and how they work together, with links to their repos be provided on the Kubescape website and in the project's primary or project-governance repos to improve discovery.

- [X] **Document the project's release process.**

The project's release process is documented in the [release file](https://github.com/kubescape/project-governance/blob/main/RELEASE.md) of the project-governance repo. It includes the release schedule, steps with links to actions workflows, and items to be performed post-release. While the release process explains a schedule of every 6-8 weeks, the project issues releases more frequently.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [X] **Clearly defined and discoverable process to report security issues.**

The project maintains a [centralized security.md](https://github.com/kubescape/project-governance/blob/main/SECURITY.md) file which includes a straightforward process to report security issues and is linked to by other project repo's security.md files.

- [X] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

The project performed a [review of its access controls](https://github.com/kubescape/project-governance/issues/14) as well as evaluated[ enforcement of 2FA on the project's maintainer accounts](https://github.com/kubescape/project-governance/issues/13).

The project should make additional improvements to its repository security posture by enforcing branch protection, ensuring code-reviews are performed on pull requests and changes are not committed directly to the project without review from a second individual who is not the author of the commit.

- [X] **Document assignment of security response roles and how reports are handled.**

The [central SECURITY.md](https://github.com/kubescape/project-governance/blob/main/SECURITY.md) file explains that maintainers are responsible for responding to reported vulnerabilities and will test to confirm the existence of the reported vulnerability. They further detail the issuance of Security advisory after confirmed and the expected timelines for disclosure.

- [X] **Document Security Self-Assessment.**

The project [completed its security self-assessment](Completed with [PR #1424](https://github.com/cncf/tag-security/pull/1424)) and it is [available on the CNCF TAG Security's repo](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/kubescape/self-assessment.md).

- [X] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The project has [achieved OpenSSF's best practices badge at the passing level](https://www.bestpractices.dev/en/projects/6944), however it appears an additional record is referenced by OpenSSF scorecards that does not accurately reflect its badge.

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
