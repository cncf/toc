# Fluid Incubation Due Diligence

- [Fluid Incubation application issue](https://github.com/cncf/toc/issues/1317)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Fluid

### Criteria Evaluation

Katie Gamanji (@kgamanji) and Alex Chircop (@chira001) conducted the due diligence of Fluid who applied for Incubation. The project has completed the criteria that show its maturity at Incubation. The following criteria implementations are noteworthy to call out:
- The project is mature and deployed in large scale production systems for a number of years
- Fluid offers standardised ways for multiple applications to access different datasets and cache them via standardised APIs
- Integration with a good number of backend systems for data access and caching allows applications to scale in a standard way across complex systems
- A healthy ecosystem and community exists, and the future roadmap is based on feedback from adopters

The following actions were provided to the project that were considered blocking but since resolved:
- the project presented to TAG Storage over 1y ago. Considering the ongoing TAG Reboot initiative, maintainers were advised to complete the General Technical Review for the project
- the security contact page was returning a 404 error in [Reporting a Vulnerability](https://github.com/fluid-cloudnative/fluid/blob/master/SECURITY.md#private-disclosure-processes) and should be fixed
- provide a  community meeting calendar to simplify the path for contributors and adopters to join future calls. It is specified that the calls are on a bi-weekly cadence, but the date of the next meeting date is unknown
- fix the inconsistency in community meeting times. The agenda to propose topics specifies a Zoom link as opposed to DingTalk Group where the calls are actually taking place. Consolidation and removal of references to unused resources is required. 
- the link to CoC in Join Our [Community as a Member](https://github.com/fluid-cloudnative/fluid/blob/a6c6343a0ee448f1498aafcdfecada55d5fa283b/CONTRIBUTING.md#join-our-community-as-a-member) returns a 404 error and should be fixed. 

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

- Consider renaming the development branch from master -> main e.g. [Kubernetes Default Branch Migration](https://www.kubernetes.dev/resources/rename/). Also, note that renaming the branch would require an update for the contributor and release process guidelines
- Community meetings are taking place regularly and are held in Mandarin. No blockers here, however, would suggest working with CNCF staff to use an automated transcript tool if possible.
- take action on the outstanding items that lower the [OpenSSF scorecard score](https://scorecard.dev/viewer/?uri=github.com/fluid-cloudnative/fluid)
- for the listed [Fluid Adopters](https://github.com/fluid-cloudnative/fluid/blob/master/ADOPTERS.md) all GitHub handles for "contact" are 404ing and pointing to email addresses of adopters. Email addresses could represent a privacy risk, and we would recommend pointing to GitHub handlers instead.

### Adoption Evaluation

The adopter interviews reflect the project use for the Incubation level to which the project has applied. We have conducted the interviews of 3 organizations that span across different usages, scalability and feature requirements.

Noteworthy highlights, are that Alibaba Group has been running Fluid in their production environment for over 3 years, with over 5,000 dataset creations and deletions per day. Fluid provided measurable value for Alibaba Group, with significant changes in terms of efficiency improvements and cost savings. Alibaba Group also benefited from the zero-downtime upgrades of the underlying storage engine, ensuring uninterrupted model training across more than 10,000 nodes.

On the other side, Adopter 2 has been using the project for just under 1 year, with over 5,000 nodes, 600+ dataset creations and deletions on demand every day. Fluid delivered measurable value by slashing new storage integration time by 70% and accelerating training cycles 3x through distributed caching – turning multi-cloud data bottlenecks into a competitive advantage. 



### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.


## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**
  - The project presented to TAG Storage on 10-07-2024, and can be discovered at [here](https://www.youtube.com/watch?v=kBZNRbP4nOU).

<!-- (TOC Evaluation goes here) --> 

- [x] **TAG provides insight/recommendation of the project in the context of the landscape**

<!-- (TOC Evaluation goes here) --> 

The feedback from the TAG was positive, however it was conducted more than 1 year ago at the time of the project assessment for Incubation.  The project maintainers were asked to complete the General Technical Review for Fluid, which can be found [here](https://github.com/fluid-cloudnative/community/blob/1215ac22dc7b757cc9b63eee8660c3e15210ba8f/docs/general-technical-review-cncf-incubating.md).

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

<!-- (TOC Evaluation goes here) --> 

The project fulfills the vendor-neutral criteria. The main communication channel is on DingTalk and WeChat groups. Also the project uses [Slack](https://cloud-native.slack.com/archives/C02ADG209SP), on the CNCF workspace via `#fluid` channel. 
[Community meetings](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) are held on DingTalk with a bi-weekly cadence. Users are also capable of finding the meeting notes on [Fluid GitHub Community Repository](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) and to follow up on the latest topics of discussion.


- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
- Met during Project's application on 27-Apr-2021.

<!-- (TOC Evaluation goes here) --> 

Fluid Sandbox application can be found [here](https://github.com/cncf/sandbox/issues/245).

The project maintainers have understood the expectations for Sandbox projects and requirements for moving forward through the CNCF incubation level.

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

<!-- (TOC Evaluation goes here) --> 

The project maintainers have added the following additional documentation as follows.
- [Installation documentation](https://fluid-cloudnative.github.io/docs/next/get-started/installation)
- [User documentation](https://fluid-cloudnative.github.io/docs/next)
- [Architecture documentation](https://fluid-cloudnative.github.io/docs/next/core-concepts/architecture-and-concepts)

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Clear and discoverable project governance documentation.**

Governance documentation is located here: https://github.com/fluid-cloudnative/fluid/blob/master/GOVERNANCE.md


- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

Fluid has maintainers from 5 organisations, and committers from many others.

Updated list: https://github.com/fluid-cloudnative/fluid/blob/master/MAINTAINERS_COMMITTERS.md

Regular meetings are held bi-weekly: https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule

Contributors are regularly promoted using the established voting/approval process: https://github.com/fluid-cloudnative/community/issues?q=is%3Aissue%20state%3Aclosed

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Latest elected members are listed [here](https://github.com/fluid-cloudnative/community/issues?q=is%3Aissue+is%3Aclosed), including 9 new committers in the last year.

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

<!-- (TOC Evaluation goes here) --> 
Decisions are made by maintainers who are actively involved in the project, and and there is a clear path outlined for people to become maintainers. Current maintainers are spread across several companies and the project is also encouraging new maintainers to join. Contibuting members and thier respective companies are listed [here](https://github.com/fluid-cloudnative/fluid/blob/master/MAINTAINERS_COMMITTERS.md).

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

<!-- (TOC Evaluation goes here) --> 
The [Technical Oversight Committee（TOC)](https://github.com/fluid-cloudnative/fluid/blob/master/GOVERNANCE.md#groups-and-leadership) is in charge of making decisions related to the future of the project, such as governance policies, management of sub-structures, and security processes. The TOC is might be required to participate in the vote in various situations. A vote is passed with 2/3 of the votes.

In addition, the project decisions are discussed in open issues and we use vote process to make the final decision. For example: [Agree to accept New Committers](https://github.com/fluid-cloudnative/community/issues/60).

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

<!-- (TOC Evaluation goes here) --> 
The project has [5 levels of responsibility](https://github.com/fluid-cloudnative/fluid/blob/ea2351bd72437af66d6f18577209f4878d62e01e/GOVERNANCE.md#groups-and-leadership), each one building on the previous:

- Contributor: comments on an issue or pull request, people who add value to the project (whether it’s triaging issues, writing code, or organizing events), or anybody with a merged pull request
- Committer: community members who have shown that they are committed to the continued development of the project through ongoing engagement with the community
- Maintainer: Maintainers are expected to contribute increasingly complicated PRs/designs and review PRs/designs, under the guidance of the existing maintainers. One who wants to be a maintainer should have been working for the project for at least 3 months.
- Technical Oversight Committee (TOC): The TOC functions as the core management team that oversees the community. The TOC has additional responsibilities over and above those of Maintainers
- Community Chair: Community chair is primarily responsible for performing community development work and administrative functions.

Fluid also has a [fluid-security team](https://github.com/fluid-cloudnative/fluid/blob/master/SECURITY.md) responses to the any security related issues.



- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

<!-- (TOC Evaluation goes here) --> 
All maintainers share all domains of responsibility. Current list of maintainers can be found [here](https://github.com/fluid-cloudnative/fluid/blob/master/MAINTAINERS_COMMITTERS.md), including ways to become a contributing member of the project through the listed [Governance](https://github.com/fluid-cloudnative/fluid/blob/ea2351bd72437af66d6f18577209f4878d62e01e/GOVERNANCE.md#groups-and-leadership) roles and responsabilities.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

<!-- (TOC Evaluation goes here) --> 
An example of adding new maintainers and also moving existing maintainers to emeritus status is listed [here](https://github.com/fluid-cloudnative/fluid/commit/96b140d99a58efb9a260c543f3206dda90b2a5bc).

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

<!-- (TOC Evaluation goes here) --> 
N/A

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

<!-- (TOC Evaluation goes here) --> 
The list of current maintainers is [here](https://github.com/fluid-cloudnative/fluid/blob/ea2351bd72437af66d6f18577209f4878d62e01e/MAINTAINERS_COMMITTERS.md).

- [ ] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) --> 
There are 7 active maintainers involved in the project development, from various organizations, such as Nanjing University, Alibaba, Alluxio, JuiceData and China Telecom. 


- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

<!-- (TOC Evaluation goes here) --> 

This is covered in the [contrubuiting guide](https://github.com/fluid-cloudnative/fluid/blob/ea2351bd72437af66d6f18577209f4878d62e01e/CONTRIBUTING.md) for Fluid.

- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

<!-- (TOC Evaluation goes here) --> 
The Code of Conduct can be found [here](hhttps://github.com/fluid-cloudnative/fluid/blob/ea2351bd72437af66d6f18577209f4878d62e01e/CODE_OF_CONDUCT.md), which referenced the CNCF Code of Conduct.

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

<!-- (TOC Evaluation goes here) --> 
[Contributing guide](https://github.com/fluid-cloudnative/fluid/blob/ea2351bd72437af66d6f18577209f4878d62e01e/CONTRIBUTING.md#code-of-conduct) refers to the agreed Code of Conduct.

- [x] **All subprojects, if any, are listed.**

<!-- (TOC Evaluation goes here) --> 
N/A

## Contributors and Community


### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) --> 
The project provides [3 levels of engagement routes](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/GOVERNANCE.md) for community members, each one building on the previous responsibilities:
- Contributors - base level interaction on the project through comments on issues or pull request 
- Committers - community members who have shown that they are committed to the continued development of the project through ongoing engagement 
- Maintainers - lead the development of the project, through contributions to increasingly complicated PRs/designs and review PRs/designs, under the guidance of the existing maintainers

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) --> 
Fluid makes use of GitHub issues and PRs. The following guidelines are available:
- [Filling an issue](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#filing-issues), including templates for:
  - [Report a bug](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/.github/ISSUE_TEMPLATE/bug-report.md)
  - [Document improvements](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/.github/ISSUE_TEMPLATE/docs-defect.md)
  - [Feature Request](https://github.com/fluid-cloudnative/fluid/blob/master/.github/ISSUE_TEMPLATE/feature-request.md)
- with a list of examples for [code contributions](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#code-contributions)

The project also provides a guide on how to set up a [development workspace
](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#setting-up-development-workspace) and submit a PR through available [template](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/.github/PULL_REQUEST_TEMPLATE.md).


- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) --> 

Public communication channels for Fluid are listed [here](https://github.com/fluid-cloudnative/fluid?tab=readme-ov-file#community), including a DingTalk and WeChat groups, and Slack. 

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) --> 
The main channel of communication are listed below: 
- CNCF Slack channel: [#fluid](https://cloud-native.slack.com/archives/C02ADG209SP)
- WeChat Group: Fluid Open Source Group
- DingTalk Group: Fluid Open Source Group

[Community meetings](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) are held on [DingTalk Group](https://qr.dingtalk.com/action/joingroup) with a bi-weekly cadence. 

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) --> 
Community members are able to join community meetings on a bi-weekly basis. The meeting notes can be found on [Fluid GitHub Community Repository](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) and to follow up on the latest topics of discussion.

Community meetings are held in Mandarin, and are a great way to engage with the maintainers and get involved in project development. 

TODO: calendar for meetings


- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) --> 

Details on how to contribute and engage with the project are listed on the [Contributing Guidelines](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md), including a path to [join the community as a member](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#join-our-community-as-a-member). The contributing members are also encouraged to follow the [Code of Conduct](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CODE_OF_CONDUCT.md) as part of getting involved. 


- [x] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) --> 
The project has a good history of adding new contributors e.g. [listing 10 new contributors](https://github.com/fluid-cloudnative/fluid/commit/37ddcf63e88303578fee99e3970465fbddc404e3) in addition to existing maintainers. Also, there is a good distribution of involved organizations, ensuring the continuity of the vendor-neutral development of Fluid.  

Also, over the last 2 years the project has a consistent ratio of contribution vs involved contributors, that can observed in the [DevStats contribution distribution](https://fluid.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now). 


## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **History of regular, quality releases.**

<!-- (TOC Evaluation goes here) --> 

### Required 

- [ ] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document what the project does, and why it does it - including viable cloud native use cases.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document the project's release process.**

<!-- (TOC Evaluation goes here) --> 

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) --> 
The [SECURITY.md](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/SECURITY.md#reporting-a-vulnerability) file outlines the vulnerability reporting process, including a mailing list (fluid.opensource.project@gmail.com) monitored by the maintainers that can be used for private disclosures.

Community members have the option to reach out to [security contacts](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/SECURITY_CONTACTS) to kickstart any private or public disclosure processes. 

Successful examples of fixed reported security issues, can be found below:
- [On a compromised node, the fluid-csi service account can be used to modify node specs](https://github.com/fluid-cloudnative/fluid/security/advisories/GHSA-93xx-cvmc-9w3v)
- [OS Command Injection for Fluid Users with JuicefsRuntime](https://github.com/fluid-cloudnative/fluid/security/advisories/GHSA-wx8q-4gm9-rj2g)


- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) --> 

2FA required for org members; branch protections enabled.

To be added as a GitHub member within the Fluid organization, each member should enable two-factor authentication (2FA). Also, the repository has branch protection rules enabled, which enforce certain workflows or requirements before a collaborator can push changes to a branch in the repository.

- [x] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) --> 
One of the core responsibility for maintainers is to respond to time-sensitive security release processes. Although this should be a rare occurrence, if a serious vulnerability is found the maintainers are expected to dedicate time to the fix, which is a process that might take up to several full days of work to implement.

[GitHub Security Advisor](https://github.com/fluid-cloudnative/fluid/security/advisories) is used as the communication channel during the process of identifying, fixing & shipping the mitigation of the reported vulnerability.

The advisory will only be made public when the patched version is released to inform the community of the breach and its potential security impact.

Additionally, security scanning is enabled for the  Fluid project and the[security contacts](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/SECURITY_CONTACTS) are responsible for assessing and providing a fix for the reported vulnerability.


- [x] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

The completed security self-assessment can be found [here](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/security/self-assessment.md).


- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) --> 
The project has a passing [OpenSSF Best Practices](https://www.bestpractices.dev/en/projects/4886) badge, with a 100% completion level, which is linked in [Fluid](https://github.com/fluid-cloudnative/fluid/tree/master) repository.

In addition, Fluid has [OpenSSF Scorecard report](https://scorecard.dev/viewer/?uri=github.com/fluid-cloudnative/fluid) via a score of 9.1 at the time of due diligence assessment. 



## Ecosystem

### Suggested

N/A

### Required

- [ ] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

<!-- (TOC Evaluation goes here) --> 

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [ ] **TOC verification of adopters.**

<!-- (TOC Evaluation goes here) --> 

Refer to the Adoption portion of this document.

- [ ] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

<!-- (TOC Evaluation goes here) --> 

#### Adoption

##### Alibaba Group - technology company

Alibaba Group has been running Fluid in their production environment for over 3 years, with over 5,000 dataset creations and deletions per day on OSS/HDFS/NAS. The organization is currently using version 1.0.1, with a small portion still running on 0.9.2.

The project optimized Alibaba Group's large-scale AI training pipelines by speeding up training sample and model loading. The automatic scaling of the distributed cache improved GPU utilization for over 2,000 concurrent tasks. Data-aware scheduling increased throughput by 40% and reduced GPU idle cycles by 30%. The organization also uses Fluid for scheduling and co-locating multiple data cache clusters along with the application Pods that utilize them.

Alibaba Group chose Fluid over alternatives like Datashim and Alluxio’s operator due to its superior storage versatility. It supports diverse data sources (JuiceFS, PVC, GPFS, Lustre) without cloud platform constraints, working seamlessly with both CSI and sidecar modes. Additionally, through Kubernetes CRDs, it simplifies operations like cache scaling, data preloading, migration, and custom operations, integrating perfectly with existing infrastructure. More importantly， Fluid can support large scale deployments. According to their test, Fluid can reliably manage a cluster with over 10 thousand nodes. In 24 hours, there were more than 5,000 Fluid Dataset CRs and over 6,000 AI jobs accessing data produced via mounted Fluid Datasets.

The adopter's experience in using and integrating Fluid with existing services has been mainly positive. In most cases, the documentation provided in the GitHub repository was sufficient to get the examples working. For other scenarios, the team referenced some best-practice blog posts written by the community. The main challenges encountered were stability at a large scale (which Alibaba Group addressed by implementing real-time monitoring of storage mount statuses, allowing Fluid to automatically repair abnormal mount points to ensure data access availability and reduce manual intervention) and the evolving needs of AI users, which demanded a high level of API abstraction.

The adopter was able to engage with the community members through available channels, mainly to report issues they encountered and to propose new feature requests. Alibaba Group used GitHub issues and pull requests to discuss with the wider community and maintainers.

Overall, Fluid has provided measurable value for Alibaba Group, with significant changes in terms of efficiency improvements and cost savings. The adopter also benefited from the zero-downtime upgrades of the underlying storage engine, ensuring uninterrupted model training across more than 10,000 nodes.

August 2025

##### Adopter 2- technology company

Adopter 2 has been using the project for just under 1 year. The company operates a multi-cloud architecture, making Fluid essential for cross-cloud data access compatibility for containers and storage, performance optimization for data-intensive workloads, and cost reduction on cross-cloud dedicated lines. Fluid operates in production with over 5,000 nodes, 600+ dataset creations and deletions on demand every day. The current version used is 1.0.5.

Some of the key features that the organization leverages are unified storage orchestration via ThinRuntime that allows for integrating diverse storage systems (self-developed storage, JuiceFS, GPFS, cloud provider‘s storage) and eliminates storage-specific CSI plugins, which reduces management overhead (ref: https://github.com/fluid-cloudnative/fluid/blob/master/addons/readme.md). PVC-based acceleration for all storage types enables acceleration of existing PersistentVolumeClaims (PVCs) without migration and delivers consistent performance for heterogeneous storage (ref: https://fluid-cloudnative.github.io/docs/next/tutorials/storage-acceleration/accelerate-pvc-with-fluid). While cross-namespace cache sharing enables multiple teams to share cached datasets across Kubernetes namespaces, which eliminates redundant data copies with an overall 40% bandwidth cost reduction (ref: https://fluid-cloudnative.github.io/docs/next/tutorials/dataset-creation/share-data-across-namespace-csi).

Adopter 2 has chosen Fluid over storage-specific operators or proprietary solutions after thorough research. Fluid provides unified management through a single control plane, orchestrating diverse storage systems and eliminating fragmented, operator-based management. It also offers vendor agnosticism, as Fluid avoids cache/cloud provider lock-in, critical for adopters' hybrid environment spanning Alibaba Cloud, Baidu Cloud, Tencent Cloud, and on-premises IDCs. Finally, the project enables large-scale production validation.  With over 5,000 nodes in production, Fluid’s battle-tested scalability and optimizations were essential, and it's an area where the project excels.

Overall, the main technical flexibility that sealed Adopter 2's decision was dual runtime modes. Fluid supports both CSI driver and FUSE sidecar modes, allowing the adopter to use CSI for standard Kubernetes clusters, leveraging sidecar mode for serverless containers where CSI is unsupported and a consistent data plane - a flexibility that enables the organization to maintain uniform data access across all environments without platform-specific reengineering.

The adopter's experience in using and integrating Fluid with existing services has been mainly positive. The team found Fluid’s documentation helpful (https://fluid-cloudnative.github.io/docs). The primary adoption challenge centered on troubleshooting complexity in production environments. Issues could originate from Fluid's control plane, underlying storage systems (e.g. JuiceFS/GPFS), or Kubernetes' PersistentVolume (PV) layer – requiring expertise across all domains to isolate failures. The adopter has actively engaged with Fluid’s community to resolve their production challenges. They communicated with community maintainers via the DingTalk group, reporting issues they encountered and expressing their needs for new functionalities.

The project delivered measurable value by slashing new storage integration time by 70% and accelerating training cycles 3x through distributed caching – turning multi-cloud data bottlenecks into a competitive advantage. If Fluid were archived, the organization would face significant operational disruption due to its deep integration into its hybrid cloud infrastructure, as 100% of cross-cloud data workflows (caching, multi-storage access, serverless integration) is based on Fluid. Features like ThinRuntime abstraction and data operations(scale, data load) lack drop-in alternatives. In this scenario, the organization would fork Fluid’s codebase to maintain critical functionality for a long time, with a dedicated contributor team of 2-3 people.

Based on Adopter 2's hybrid cloud practice, the most critical growth opportunity for Fluid is deep integration with multi-cluster orchestration ecosystems like Karmada, Clusternet, and Open Cluster Management. The project could introduce some improvements by reducing maintainers' workloads. Fluid should enrich docs with a deep technical dive (e.g. how ThinRuntime abstracts storage under the hood) and battle-tested best practices (like the 5,000-node tuning checklist). This would transform tribal knowledge into self-service answers.

Resulting from the adopter's production experience and strategic roadmap, their future engagement with Fluid will focus on Data-Aware Multi-Cluster Scheduling, Cost-Optimized Runtime Switching, and Fluid at Scale Group. The main vectors of work are focused on co-developing cache-aware placement policies with multi-cluster orchestration projects (e.g. Karmada/OpenClusterManagement), enabling intelligent workload scheduling based on dataset availability and performance thresholds.

The team also plans to implement dynamic runtime selection that automatically matches storage engines (JuiceFS/Jindo/Alluxio) to workload patterns – optimizing for performance, cost, or data characteristics without manual intervention.
Through Fluid at Scale Group, Adopter 2 aims to publish enterprise-grade blueprints for 10,000+ node deployments and co-maintain disaster recovery protocols for mission-critical environments.

August 2025

##### Adopter 3 -  financial services

Adopter has been using the project for 3 years in production having previously tested for 6 months in dev environments.   Adoption was at an early stage in the project development, has received lots help from the community and has seen the project mature over time.

The project is in use in two areas of the business covering financial trading as well as analytics processes on up to 1000 nodes in production.   It was adopted as a solution that was able to cache different types of data, deal with very high levels of throughput, whilst providing a common interface for different teams and applications.  Adoption started in a cloud environment, and needed to be able to work with different cloud providers as well as operate in on-prem infrastructure.

Fluid supports multiple backends with a standard API, including JuiceFS, which was an important factor in the adoption process.  Other projects were considered, but they did not have the flexibility, reliability, or the ability to scale to 1000 nodes.

Documentation was limited early in the project, and the active community was important to the adoption.  Over time, the documentation has improved, but there is still work to do.  That said, the documentation is good enough for new teams to be able adopt the project effectively.   The code is also very well documented and the code documention often acts as a guide for project usage.

Integrations with the ecosystem and backend components have improved over time, and the early problems have been addressed, so the system is very reliable in production.   Multiple plugins now provide a similar experience and maturity level.   As Fluid is now tightly integrated with multiple important stateful applications in the adopter's environment, upgrades are performed once a year.

The adopter has committed to the project, and is planning to continue using the project in product into the future, and has commented that the project could benefit from broadening the community outside of China.

August 2025
