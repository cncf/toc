# Lima Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1348)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Lima

### Criteria Evaluation

[Chad Beaudin](https://github.com/chadbeaudin) and [Ricardo Rocha](https://github.com/rochaporto) conducted both due diligence and adopter interviews for Incubation. [Ahmed Bebars](https://github.com/abebars) from the CNCF TAB also participated in the adopter interviews. The project has completed the criteria that shows its maturity at the applied level.

The TOC would like to highlight some areas from the project review:
* Lima is widely used either directly by end users and adopters or indirectly via projects that rely on it as a core dependency. Rancher Desktop or Colima are examples of the latter mentioned during the adopter interviews. In this respect Lima is a special project and this was taken into account by the TOC when considering adequate adopters to be interviewed
* The founders and maintainers of the project are very active and well known members of the cloud native community. This shows in the level of integration of Lima with other tools in the ecosystem, as well as in the open source best practices being followed
* The project is unique in its goals and fills a gap in the ecosystem for both macOS and Windows environments

The following required incubation items were raised and taken into account by the project during the review period:
* [Tracking affiliation information for maintainers](https://github.com/lima-vm/lima/issues/3794)
* [Tracking of code and doc ownership](https://github.com/lima-vm/lima/issues/3795)
* [Pointer for users to submit issues to the project](https://github.com/lima-vm/lima/issues/3796)

Other items suggested for incubation were taken as issues in the project to be followed up:
* [Tracking of emeritus maintainers](https://github.com/lima-vm/lima/issues/3793)
* [Improvements to the developer's guide](https://github.com/lima-vm/lima/issues/3792)
* [Setting up regular meeting slots](https://github.com/lima-vm/lima/issues/3798)

### Adoption Evaluation

The adopter interviews showed project usage at a level compatible with CNCF incubation. Adopters covered both direct usage and indirectly via projects like Rancher Desktop and Colima, both relying heaving on Lima. Adopters rely on Lima to manage virtual machines but mainly to offer container development environments for macOS and Windows.

Adopters highlight as main strengths of the project the stability of the releases with no backwards incompatible changes found even when staying close to the latest version, no issues seen when upgrading the operating system versions or moving to different architectures, the support for multiple VM networking modes, and the ability to use the same tool in both macOS and Windows. Adopters also reported as one key strength the founders and maintainers being well known and recognized in the community, the very good responsiveness and follow up to proposed fixes and contributions and the health of the community overall.

Adopters see opportunities in better definition and priorization of roadmap items and better communication and dissemination of the project in relevant communities, potentially engaging a community manager to help with these tasks. Adopters would also like to see performance improvements on boot time speeds with alternatives to cloud-init, as well as clarifications where Lima stands for users relying on the project via higher levels tools.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

The last update was given on [2022-10-06](https://docs.google.com/document/d/1k7VNetgbuDNyIs_87GLQRH2W5SLgjgOhB6pDyv89MYk/edit?tab=t.0#heading=h.nyb77lbc7y9z) to TAG-Runtime.

- [x] **TAG provides insight/recommendation of the project in the context of the landscape**

Given the ongoing TAG reboot, this task is taken over by the TOC.

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

No issues were found during due diligence, both code and documentation are vendor neutral.

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		

All done during the [sandbox onboarding process](https://github.com/cncf/sandbox/issues/207), completed on 2024-06-10.

- [x] **Due Diligence Review.**

The initial application was initially submitted June 04, 2024.

The project submitted initial DD information along with [the application](https://github.com/cncf/toc/issues/1348).

Project maintainers were responsive throughout the process. The process satisfies the due diligence review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The project documentation is available at: https://lima-vm.io/docs/.

It provides extensive information on the different platforms supported, the installation options for the different environments (both binary and source), as well as usage and configuration guides.

The [developer guide](https://lima-vm.io/docs/dev/) is not particularly complete and could benefit from an update - [issue opened](https://github.com/lima-vm/lima/issues/3792). This is seen as an improvement and not a blocker for incubation.

## Governance and Maintainers

### Suggested

- [x] **Clear and discoverable project governance documentation.**

Lima's governance is documented in its [main docs](https://lima-vm.io/docs/community/governance/), along with the definition of the different roles (contributor, maintainer, committer, reviewer) and the process for promotion and demotion of Maintainers.

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The governance documentation has been consolidated in [May 2024](https://github.com/lima-vm/lima/commit/4a1c05325a18be8a6de900838198a75923553184). Since then there has been one significant update linking to policies regarding DCO, code of conduct and 3rd party dependencies.

As the governance process has matured, the TOC recommends a check with TAG-CS for validation of the existing practices.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Governance is up to date with actual practices.

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

Vendor neutrality is clearly stated and [linked](https://lima-vm.io/docs/community/governance/#maintainership) from the governance docs.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

Leadership in the project is assigned to Maintainers, which can be one of:
* Committer: full maintainership
* Reviewer: limited maintainership, no access to merge PRs

Acceptance of contributions can be done by Committers only, but Reviewers can participate in the process.

Decisions to governance or other are done by [a majority vote](https://lima-vm.io/docs/community/governance/#other-decisions) by Committers.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

Roles are well documented in the [governance documentation](https://lima-vm.io/docs/community/governance/#roles).

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

Roles are [well documented](https://lima-vm.io/docs/community/governance/#current-maintainers) with information on how promotion / demotion is done. No information on emeritus maintainers is currently kept - [requested here](https://github.com/lima-vm/lima/issues/3793).

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Example of an [addition of a maintainer](https://github.com/lima-vm/lima/pull/2383) in June 8 2024. [Another example](https://github.com/lima-vm/lima/commit/95a1dd971a879acc1f246becd66772221247643a) from Nov 18 2024.

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

There is a [dedicated page](https://lima-vm.io/docs/community/subprojects/) listing all sub-projects and stating maintainers are shared with the main Lima project.

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

The list of current maintainers is [kept here](https://lima-vm.io/docs/community/governance/#current-maintainers) but no information is available on affiliation - [requested here](https://github.com/lima-vm/lima/issues/3794).

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

There are currently 4 active committers and 3 active reviewers, which suits the size and scope of the project.

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

Not currently done or not obvious - [requested here](https://github.com/lima-vm/lima/issues/3795).

- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

Done in the main [governance document](https://lima-vm.io/docs/community/governance/).

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

Done in the main [governance document](https://lima-vm.io/docs/community/governance/) and in the [github home](https://github.com/lima-vm/lima?tab=readme-ov-file#code-of-conduct).

- [x] **All subprojects, if any, are listed.**

There is a [dedicated page](https://lima-vm.io/docs/community/subprojects/) listing all sub-projects.

## Contributors and Community

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

[Well documented](https://lima-vm.io/docs/community/governance/#addition-and-promotion-of-maintainers) in the governance document.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

The [contributing guide](https://lima-vm.io/docs/community/contributing/) documents how to propose changes to the project.

There is currently no reference in the docs on how users can submit issues - [requested here](https://github.com/lima-vm/lima/issues/3796).

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

Communication channels are [documented here](https://github.com/lima-vm/lima?tab=readme-ov-file#communication-channels) and include GitHub discussions and the CNCF slack (#lima). Also documented in the [community page](https://lima-vm.io/docs/community/).

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

Communication channels include:
* [GitHub Discussions](https://github.com/lima-vm/lima/discussions)
* [CNCF Slack](#lima)

These apply to sub-projects as well.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

Project mentions no regular meetings are currently set due to timezone diversity. Given the project still maintains a very active community, the TOC accepts this as a non-blocker for incubation, but recommends the project considers setting up regular meetings in diverse timezones to help engage newcomers in the project - [requested here](https://github.com/lima-vm/lima/issues/3798).

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

A [contributing guide](https://lima-vm.io/docs/community/contributing/) exists with links to the approval process and governance roles. It could be improved by adding information on how people willing to help the project can get more engaged and find open issues to work on - [requested here](https://github.com/lima-vm/lima/issues/3797).

- [x] **Demonstrate contributor activity and recruitment.**

As per [LFX Insights](https://insights.linuxfoundation.org/project/lima/contributors?timeRange=past365days&start=2024-08-05&end=2025-08-05) has 293 contributors in the last 12 months from 101 organizations. 2 contributors are responsible for 51% of the activity with 3 more with over 5% of the contributions. 2 organizations are responsible for 60% of the activity with 3 more with over 5% of the contributions.


## Engineering Principles

### Suggested

- [x] **Roadmap change process is documented.**

The roadmap is kept via labels in github issues, as [documented here](https://lima-vm.io/docs/community/roadmap/).

The TOC recommends that moving forward the project considers extending the roadmap process to include more discussion in the project, going beyond the simple availability of a contributor to work on the items.

- [x] **History of regular, quality releases.**

[Releases](https://github.com/lima-vm/lima/releases) are made often and well documented.

### Required 

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

Purpose and goals of the project are well summarized in the [project homepage](https://lima-vm.io/docs/).

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

The purpose of the project is documented in its [homepage](https://lima-vm.io/docs/) and documents the integration with containerd and nerdctl as other projects in the cloud native ecosystem.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The project keeps a [roadmap](https://lima-vm.io/docs/community/roadmap/) page with roadmap issues kept with appropriate [labels](https://github.com/lima-vm/lima/issues?q=is%3Aissue+is%3Aopen+label%3Aroadmap). As documented above in the DD, the TOC recommends the roadmap definition is extended to include more discussion in the whole project in addition to the availability of a contributor to work on items.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

The project documents both the [VM lifecycle](https://lima-vm.io/docs/dev/internals/#vm-lifecycle) and the different ways to [run guest machines](https://lima-vm.io/docs/config/vmtype/). This is an item that while present could benefit from improvements in the developer's guide - [issue opened](https://github.com/lima-vm/lima/issues/3792).

- [x] **Document the project's release process.**

The [release process](https://lima-vm.io/docs/community/governance/#release-process) is well documented in the governance documentation.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

The process to report security issues is [documented here](https://github.com/lima-vm/.github/blob/main/SECURITY.md) and linked as a [tab](https://github.com/lima-vm/lima?tab=security-ov-file) from the main readme file in the github home of the project. As well as a dedicated section in the [project's homepage](https://lima-vm.io/docs/security/).

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

The requirement for 2FA is stated in the definition of the Committer role in the [governance guide](https://lima-vm.io/docs/community/governance/#roles).

- [x] **Document assignment of security response roles and how reports are handled.**

Process for handling security reports is documented [here](https://github.com/lima-vm/.github/blob/main/SECURITY.md).

- [x] **Document Security Self-Assessment.**

The security self-assessment has been [completed](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/lima/self-assessment.md) with TAG-Security.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

Achieved with report [available here](https://www.bestpractices.dev/en/projects/6505).


## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

A list of adopters is available in the main project [github page](https://github.com/lima-vm/lima/tree/master?tab=readme-ov-file#adopters).

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

Validated during the adopter interviews. See TOC highlights in the Criteria Evaluation for more details on indirect adoption.

- [x] **TOC verification of adopters.**

See the adoption section of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

Both containerd and nerdctl are documented in the [main project page](https://lima-vm.io/docs/) and additional integrations in the [github home](https://github.com/lima-vm/lima/tree/master?tab=readme-ov-file#adopters).

#### Adoption

##### AWS

###### Usage

AWS started using Lima as part of their project Finch in November 2022 (prior to the 1.0 release). Initially considered as a tool to manage Linux virtual machines, Lima can be used beyond this initial goal since it's container centric and integrates well with nerdctl. Another motivation to use the project includes support for different VM networking modes, with new ones still being added without any issue observed for existing deployments.

AWS felt confident with the project after an initial evaluation and thanks to the founders and authors of the project being well known and recognized members of the community. Originally targeting macOS only, added support for Windows stopped efforts to find a different solution for these environments.

AWS uses the project in production, currently by 1000s of users. AWS stays as close as possible to the latest version and has observed no backwards incompatible changes until today. Internally they rely on an automated workflow triggering tests and integration from the latest upstream versions.

AWS is active in the upstream project, contributing major features (Windows support) and bug fixes. Moving forward AWS is particularly interested in performance improvements for boot speed times.

###### Perception and Feeling

Issues during the project adoption were handled quickly by the maintainers and the community. In some cases the proposed fixes came from AWS and were incorporated quickly.

During pre-1.0 the documentation was good but disorganized and spread over multiple markdown files. The situation today has improved significantly with a well structured website and easy to find information and diagrams.

Communication with the project is mostly done via GitHub issues and discussions, with slack used to find any previous discussion on a topic before opening GitHub issues.

If needed AWS could consider stepping up and becoming maintainers of the project but there is no issue with community engagement today.

###### Strengths and Improvements

AWS considers that it would take them weeks to months to implement a tool similar to Lima, highlighting savings in engineering effort. A main strength of the project comes from the set of founders and maintainers being very experienced in the community and well connected and knowledgeable with other projects in the ecosystem. Stability is another strength with no observed backward incompatibilities even with many new features being added. Finally support for macOS and Windows is also seen as a relevant strength since they can rely on the same tool for both environments.

Suggestions for improvements include better tracking and review of roadmap items, including discussions on prioritization and on the feature design and development. AWS looks forward to performance improvements for boot time speed and suggests looking for alternatives to cloud-init.

##### SUSE

###### Usage

SUSE adopted Lima in late 2021 integrating with Rancher Desktop, with the first releases based on Lima happening that same year. Later in 2022 a SUSE employee became a maintainer of the upstream Lima project.

The main motivation for adoption was to bring a container runtime to Rancher Desktop running inside a virtual machine. This gave more control over the environment compared to minikube with a custom VM OS. SUSE also considered other VM management software or even writing their own wrapper around QEMU prior to looking at Lima.

SUSE is currently relying on Lima in production as a desktop application and has 10k+ instances deployed. The Lima dependencies are updated regularly syncing with the latest version and internal project releases happen every 6 to 8 weeks.

Jan Dubois from SUSE is currently one of the core maintainers of Lima.

###### Perception and Feeling

Adoption of the project depended on a few custom requirements being taken by Lima, which was done quickly when working with the upstream project. The documentation was seen as good enough at the start and they realize it got much better now.

Interactions with the project are mostly done async due to timezone differences with some core maintainers in the project, relying on github issues and slack. Direct communication with other adopters (such as AWS) happens in a more direct way.

SUSE is not aware of any viable alternative to their current setup and Lima is now very much engrained in Rancher Desktop. Community engagement is very good, but if this would ever change they would consider forking the project and continue development.

###### Strengths and Improvements

SUSE highlights the responsiveness of the founder and main maintainer (Akihiro Suda) as a core strength of the project. The project also eases handling problems that are by nature hard to handle, giving socket management as an example. This allows them to have fewer people working on these areas as they can leverage the upstream features instead.

SUSE sees this as a good example of giving a little to a project and getting a lot back from it.

Improvements as seen by SUSE include establishing regular meetings, with maintainers in different timezones chairing those sessions. SUSE also suggests the project looks for a community manager to ensure it is better disseminated in the right communities.

##### Intuit

Intuit relies on Lima via Rancher Desktop. Some of the remarks below refer to the latter and apply to Lima only as a core dependency.

###### Usage

Intuit started looking at Lima via the Rancher Desktop project in 2022 with usage starting in 2023. From February 2024 it's available to all users.

Intuit switched to Rancher Desktop from Docker Desktop as a tool to deploy and manage local kubernetes and docker developer environments and in particular to build container images. The user base is mostly on macOS with a minimal set also on Windows. The move occurred after an initial callout to the Docker Desktop community asking for a fully open source alternative. 

Usage of Rancher Desktop and Lima is currently production level with 8500 Mac deployments and 500 in Windows. Intuit stays close to the latest Rancher Desktop releases (about 2 weeks behind the latest), which is also close to the Lima releases.

###### Perception and Feeling

The move from Docker Desktop to Rancher Desktop was done with a fraction of two people and saw a smooth transition for over 8500 individual installations. Intuit saw no need to assign a dedicated team to this process. The migration concluded and went well above expectations.

Documentation was appropriate when needed with most questions being related to automation and enterprise rollouts of Rancher Desktop. Engagement with the community is seen by Intuit as some of the best ever experienced, with community health and responsiveness being one of the early checks done. A few zoom calls and direct slack communication was done during the early adoption period.

Intuit considers that if ever needed it would be hard to find a similarly smooth solution and would consider moving to a cloud based solution for developer environments instead of maintaining a fork.

###### Strengths and Improvements

For Intuit the main strength of both Rancher Desktop and Lima is the ability to achieve a similar solution to that of Docker Desktop based purely on open source tools. This also brought significant savings when accounting to the number of users and licensing costs of their previous solution.

Other strengths include the reliability and stability and the fact that no issues have been seen in production. The project community is small but very active on slack.

Intuit believes the project would benefit from additional awareness and better marketing and that it should focus more on enterprise roll-out solutions (particularly for Rancher Desktop).

##### Abiosoft

Usage of Lima in this context is done via Colima. The interview was done with the author of Colima, Abiola Ibrahim.

###### Usage

Colima has been relying on Lima since 2020, with the main motivation coming from one of the core features of the project - the ability to run headless virtual machines on macOS. Podman and Docker were also checked as alternatives.

Colima does not yet have a 1.0 release. It relies on the latest available version of Lima, and acts as an early tester of new releases. Colima plans to follow on the recently added Windows support, and add similar support directly in Colima.

The authors of Colima contributed the support for nested virtualization in Lima.

###### Perception and Feeling

Conversations with the Lima project have been very productive from the early start, with direct communication with the maintainers of the Lima project. Minor challenges during adoption included support for file watchers (originally with little interest from Lima, but later included); support for network providers and integration with Colima, with good ideas coming from replicating available functionality in Podman.

GitHub issues are the main source for communication, complemented by slack. Very happy with the support and level of response from the Lima maintainers. It appears the project is also gaining popularity, and Abiola gives as an example the multiple contributions to the project coming from GSOC internships.

###### Strengths and Improvements

Abiola sees the clear instructions and ease of contributions as a main strength of the project, together with it being an unique project offering virtual machine support for Mac environments.

Suggested improvements to the project include maturing the Windows support, as not all features are currently available across all platforms. Not a project weakness but expected as the project matures. Recent changes to cleanup the distribution channels has also slightly complicated the setup on Mac environments, with `lima-additional-guestagents` having to be installed independently for running VM with a non-native architecture (e.g., Intel on ARM).

##### Garner

###### Usage

Garner has been using Lima since 2022 as part of Rancher Desktop, as an alternative to Docker Desktop. The main motivations for adoption came from the openness to feedback and contributions from the Lima project.

Current usage covers ~20 developers supporting both development and testing workloads, with only 1 developer not adopting the project. Version in use is 1.2.1 and upgrades follow the Rancher Desktop upgrades.

###### Perception and Feeling

Garner had a good experience adopting the project and saw it being easy to setup and build. The contribution guides were clear and they highlight this is unlike their experiences with other open source projects.

Garner's engagement with the community has been minimal with some direct discussions with the core maintainers about specific pull requests, which were merged within a day. Garner sees this as a sign of a healthy project.

###### Strengths and Improvements

Garner sees as the main strength of the project the cost savings compared to alternatives such as OrbStack or Docker Desktop. Garner also points out Lima fills a niche, is a very active and healthy project and their experience is that it just works, highlighting that they've been through multiple macOS version upgrades including a move from Intel to ARM.

Garner would like to see better documentation regarding the system architecture, especially when used with Rancher Desktop to better understand where Lima fits in the stack. This could help with marketing the project as well.
