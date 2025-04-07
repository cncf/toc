# CubeFS Graduation Due Diligence

- Link to [CubeFS Graduation application](https://github.com/cncf/toc/pull/1140)

  <!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Graduation Evaluation Summary for CubeFS

### Criteria Evaluation

Kevin Wang conducted the due diligence of CubeFS who applied for graduation. The project has completed the criteria that show its maturity at graduation.

The following criteria implementations are noteworthy to call out:

- A stable and easily maintainable distributed file system, with excellent concurrent write performance, scalability and extensibility.
- Has a strong and growing community of maintainers and adopters, ensuring its long-term sustainability.
- Provided diverse channels for community users and contributors to interact, and public meeting links, recordings and notes can be easily found.
- Provided high-quality, well-organized documentation and practical examples to help users quickly learn and master the project.
- The project completed the  [third party security audit by Ada Logics](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/security/CubeFS-security-audit-2023-report.pdf) with no high-level or critical issues found. The project team actively resolved the 12 detected issues, clearly demonstrated commitment to security, which is praiseworthy.

The following actions were provided to the project that were considered blocking but have since been resolved:

- Removed the Project Lead role, previously held by one individual and considered conflicting with community neutrality. And instead, established a Technical Steering Committee (TSC) with a defined number of seats and neutrality requirements.
- Updated the governance documents to clarify the management rules for subprojects.
- Added a [RELEASE.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/RELEASE.md) file, including updating the release process to reflect the latest engineering principles criteria.
- Updated the governance documents to include roadmap changing process.

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

- TOC Reviewer recommends finalizing the ongoing Governance Review by the TAG Contributor Strategy to achieve a more comprehensive community governance.
- TOC Reviewer recommends to extend its current management rules to cover all repos under the GitHub organization, including non-subproject repos, and archive any repos that are no longer being maintained.
- Make full use of `cubefs-community` repo, as the canonical location for community governance-related documents.
- TOC Reviewer recommends improving the project's extensibility and adding more pluggable framework designs. This would facilitate easier customization and easier decoupling of downstream enhancements by adopters.
- TOC Reviewer recommends improving feature lifecycle management, including feature maturity, configuration, commit history, and release notes. This would make it easier for adopters to track community updates and backport patches as needed.
- TOC Reviewer recommends to establish Adopters WG or other similar group, as a dedicated space for adopters to collaborate on shared features and functionality to continue improving and enhancing the project.
- To foster a more inclusive global community, TOC Reviewer recommends making a plan for global community development. This plan may include initiatives like English-language community meetings and cultivating contributors from various regions to better support adopters worldwide.
- TOC Reviewer recommends organizing dedicated TSC meeting, in order to keep TSC members engaged.
- To enhance community decision-making transparency, the TOC Reviewer recommends the project provide explicit records of voting processes, e.g. manual vote counts or using [gitvote](https://github.com/cncf/gitvote).
- TOC Reviewer recommends to add explicit description of platforms supported in the [RELEASE.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/RELEASE.md)
- TOC Reviewer recommends to cross reference the [roadmap governance(https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#roadmap)] and [change process](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#changes-in-project-roadmap) on the [ROADMAP.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/ROADMAP.md) to make it easier to find for potential contributors.
- And for the [roadmap change process](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#changes-in-project-roadmap), it's recommended to include collecting roadmap proposals through public channels, and use more community fashion phrasing, which would encourage contributors to join the discussion and better understand where the project is heading to.
- TOC Reviewer recommends to update security policy to include an embargo and private disclosure period before doing public disclosure for security vulnerabilities. And tagging a release clearly as "security-fixes-only" will help users to prioritize an upgrade.

### Adoption Evaluation

The adopter interviews reflect the production grade usage of CubeFS for the applied Graduation level. According to the feedback, CubeFS simplifies the interaction between users and storage infrastructure by providing a high-performance, scalable, reliable and easy-to-maintain distributed file system. Developers leverage CubeFS to address the challenges of complex distributed file systems, including large-scale data storage, high concurrent access, data reliability and the complexity of system maintenance. Adopters have been using CubeFS in the production environment for many years, managing data ranging from hundreds of terabytes to petabytes and supporting the access of millions of clients. Overall, the stability, reliable performance and active community of CubeFS have built the trust of adopters.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation.

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

  The CubeFS team has given presentation on the TAG Storage meeting on April 24, 2024. [link to recording](https://www.youtube.com/watch?v=UgRBZhzfr4w)

  The project has requested Governance Review by TAG Contributor Strategy at: <https://github.com/cncf/tag-contributor-strategy/issues/649>

- [x] **TAG provides insight/recommendation of the project in the context of the landscape**

  TAG Storage recommendation can be found in [this doc](https://docs.google.com/document/d/1D-Y1XVmPNu_g4Vtl6ZhmueR2ATmuGjoe8Bnc-uTFsqI/edit).
  > TAG Storage has reviewed CubeFS for its graduation request, provided suggestions for updating CubeFS-csi and CubeFS-helm projects to resolve CVEs, and the team responded quickly to address the issues. We recommend that the CubeFS team continue to keep these projects up to date. We believe that the project’s health is sound in general, the customer adoption has been increasing, and the project is at a mature state and it is ready to move to the graduation level.

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

  - **Neutral resources** - CubeFS has its own channels (community branded and managed), including:
    - Homepage: <https://cubefs.io>
    - Mailing list: <users@cubefs.groups.io>
    - Slack: <https://cubefs.slack.com>
    - WeChat: <https://github.com/cubefs/cubefs/issues/604>
    - Twitter: <https://x.com/cubefs_storage>
    - Community Meeting: <https://meeting.tencent.com/dm/IpNLDdT7uyEH>

  - **Governance** - [GOVERNANCE.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md) defined vendor-neutrality requirements on the TSC, Maintainers, Committers, including:
    - [The structure of the Technical Steering Committee Section](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#the-structure-of-the-technical-steering-committee) says: `No single vendor can exceed 50% of the total number of personnel.`
    - [The structure of the Maintainers Section](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#the-structure-of-the-maintainers) says: `No single vendor can exceed 50% of the total number of personnel.`
    - [The structure of the Committers Section](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#the-structure-of-the-committers) says: `No single vendor can exceed 50% of the total number of personnel.`

- [x] **Review and acknowledgement of expectations for graduated projects and requirements for moving forward through the CNCF Maturity levels.**
  - [x] Met during Project's application on 10-Apr-2024 as a kick-off meeting.

  The [initial PR](https://github.com/cncf/toc/pull/1140) was submitted on 5-Aug-2023. The TOC Reviewer was assigned on 9-Apr-2024.

  On 10-Apr-2024, the TOC Reviewer and project maintainers met during the kick-off meeting of the Graduation process, discussed the changes and update in the TOC repo regarding the criteria and expectations for moving levels, release, and freeze time period for KubeCons.

  The project provided [updates](https://github.com/cubefs/cubefs/issues/3298) using the recommended template on 18-Jun-2024.

  The TOC Reviewer and project maintainers met multiple time during the due diligence review. Some [suggested action items](https://github.com/cncf/toc/pull/1140#issuecomment-2195009379) were provided during the review, and the project maintainers have been highly responsive throughout the process.

  Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisifies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

  - [CubeFS introduction documentation](https://cubefs.io/docs/master/overview/introduction.html) introduces the CubeFS architecture and main features.
  - [CubeFS installation documentation](https://cubefs.io/docs/master/deploy/env.html) covers several ways of deployment.
  - [CubeFS end user documentation](https://cubefs.io/docs/master/user-guide/volume.html) includes basics operations as creating a volume, using volume and using cli tool.

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

  CubeFS has been continuously updating governance doc to reflect project growth, some examples are:
  - CubeFS initial governance: <https://github.com/cubefs/cubefs/blob/v1.1.0/GOVERNANCE.md>
  - Added Committer role in May. 2023: <https://github.com/cubefs/cubefs/pull/1976>
  - Added Steering Committee in Apr. 2024: <https://github.com/cubefs/cubefs/pull/3312>
  - Update maintainer list according to activity and add steering committee member: <https://github.com/cubefs/cubefs/pull/3311>
  - Update the Governance Document to eliminate the role of the leader: <https://github.com/cubefs/cubefs/pull/3382>
    The description of 'project lead' implies a somewhat authoritarian role, but with the establishment of a Steering Committee, the Steering Committee should be considered the highest decision-making body.Thus CubeFS delete the role of 'project lead'.
  - Adding governance rules related to SIGs.: <https://github.com/cubefs/cubefs/pull/3430>
  - Adding governance rules related to RoadMap.: <https://github.com/cubefs/cubefs/pull/3430>
  - Renamed Steering Committee to Technical Steering Committee, and clarified responsibilities and lifecycle of TSC, maintainers, and committers, SIGs: <https://github.com/cubefs/cubefs/pull/3478>

### Required

- [x] **Clear and discoverable project governance documentation.**

  CubeFS governance documentation: [GOVERNANCE.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md)

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**


  Examples showing that governance is up to date:
  - Update the document of the roadmap: <https://github.com/cubefs/cubefs/pull/3358>
  - Add new Committer: <https://github.com/cubefs/cubefs/pull/3384>
  - Nomination of Chairs and Tech Leads for SIG-CSI and SIG-Docs: <https://github.com/cubefs/cubefs-community/pull/13>

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  CubeFS has clear vendor-neutrality description the [governance doc](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md), including matters related to information transparency, channel transparency, decision-making, and other aspects among vendors.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  - Decision making process on leadership roles: [GOVERNANCE.md#the-tsc-decision-making-process](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#the-tsc-decision-making-process)
  - Contribution acceptance: [CONTRIBUTING.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/CONTRIBUTING.md)
  - Requests to the CNCF: [GOVERNANCE.md#cubefs-and-cncf](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#cubefs-and-cncf)
  - Changes to governance or project goals
    - Changes to governance or project goals: [GOVERNANCE.md#changes-in-project-governance](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#changes-in-project-governance)
    - Technical Steering Committee Member is responsible for roadmap formulation: [GOVERNANCE.md#expectations-from-the-technical-steering-committeetsc](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#expectations-from-the-technical-steering-committeetsc)

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

  Main CubeFS project role update according to governance doc:
  - [Becoming a TSC member](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#becoming-a-tsc-member)
  - [Changes in TSC](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#changes-in-tsc)
  - [Becoming a Maintainer](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#becoming-a-maintainer)
  - [Changes in Maintainer membership](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#changes-in-maintainership)
  - [Becoming a committer](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#becoming-a-committer)
  - [Changes in committer membership](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#changes-in-commitership)

  Governance of to SIGs. Rules for assignment, onboarding, and removal: [GOVERNANCE.md#sig](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#sig)
  Product Security Committee Membership: Rules for assignment, onboarding, and removal: [security-release-process.md#product-security-committee-membership](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/security/security-release-process.md#product-security-committee-membership)

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

  CubeFS documents maintainers list at: <https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/MAINTAINERS.md>

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

  CubeFS currenlty has 14 top level maintainers from: JD.com, BEIKE, OPPO, Bytedance, LinkedIn, XFusion. Ref: <https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/MAINTAINERS.md>

  Activities of maintainers can be found at: [chubaofs.devstats.cncf.io](https://chubaofs.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-companies=All)

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**
  
  CubeFS has a clear maintainer lifecycle process documented in their governance doc:
  - Becoming a maintainer: [GOVERNANCE.md#becoming-a-maintainer](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#becoming-a-maintainer)
  - Changes in maintainership: [GOVERNANCE.md#changes-in-maintainership](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#changes-in-maintainership)

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

  Demonstrations regarding the maintainer lifecycle of the CubeFS project:
  - <https://github.com/cubefs/cubefs/pull/3386> Chuanqing Zhang was added as a new committer
  - <https://github.com/cubefs/cubefs/pull/3384> Shuqiang Zheng was added as a new committer
  - <https://github.com/cubefs/cubefs/pull/3311> updated maintainer list including removal of inactive maintainers, and added steering committee members

- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

  According to the [Maintainers list](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/MAINTAINERS.md), CubeFS currently has top level maintainers from OPPO, JD.com, BEIKE, Bytedance, LinkedIn, and additional committers from BIGO, VIVO.

  Definition of Maintainers and Committers can be found in the [GOVERNANCE.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md).
  Both Maintainers and Committers require diversified membership: `No single vendor can exceed 50% of the total number of personnel.`

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

  CubeFS uses a github CODEOWNERS mechanism to manage the code access between different community roles. Ref: <https://github.com/cubefs/cubefs/tree/master/.github/CODEOWNERS>

- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

  Documented at [GOVERNANCE.md#code-of-conduct](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#code-of-conduct)
  
- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

  CNCF Code of conduct is cross-linked in the [Code of Conduct of CubeFS](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/CODE_OF_CONDUCT.md)

- [x] **All subprojects, if any, are listed.**

  The following projects are associated with CubeFS and maintained as sub-projects.
  - [cubefs-helm](https://github.com/cubefs/cubefs-helm): CubeFS installation to helm in the Kubernetes ecosystem
  - [cubefs-csi](https://github.com/cubefs/cubefs-csi): CSI (Container Storage Interface) plugin for CubeFS
  - [cubefs-dashboard](https://github.com/cubefs/cubefs-dashboard): Dashboard for CubeFS

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

  According to [Governance.md#sub-projects](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#sub-projects), sub-projects can have their own repositories but follow the same governance mechanism as the main project

  Subprojects Governance descriptions can be found at:
  - cubefs-helm Governance: <https://github.com/cubefs/cubefs-csi#governance>
  - cubefs-csi Governance: <https://github.com/cubefs/cubefs-csi#governance>
  - cubefs-dashboard Governance: <https://github.com/cubefs/cubefs-dashboard#governance>

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

  Cubefs has the following roles for contributors that are related to code and non-code contributions:
  - Technical Steering committee member: [GOVERNANCE.md#expectations-from-the-technical-steering-committeetsc](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#expectations-from-the-technical-steering-committeetsc)
  - Maintainer: [GOVERNANCE.md#expectations-from-maintainers](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#expectations-from-maintainers)
  - Committer: [GOVERNANCE.md#expectations-from-committers](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#expectations-from-committers)
  - SIG member: [GOVERNANCE.md#expectations-from-sigs-member](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#expectations-from-sigs-member)
  - Product Security Committee (PSC): [security-release-process.md#product-security-committee-psc](https://github.com/cubefs/cubefs/blob/master/security/security-release-process.md#product-security-committee-psc)

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

  Defined in [CONTRIBUTING.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/CONTRIBUTING.md), in the root path of CubeFS main repo.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

  CubeFS has the following public communications channel for users and contributors documented in the [Project README](https://github.com/cubefs/cubefs#community)
  - Website: <https://cubefs.io/>
  - Mailing list: <users@cubefs.groups.io>
  - Slack: <https://cubefs.slack.com>
  - WeChat: <https://github.com/cubefs/cubefs/issues/604>
  - X/Twitter: <https://x.com/cubefs_storage>

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

  - The communication channels for CubeFS documented at <https://github.com/cubefs/cubefs#community>
  - Besides public channels, CubeFS has a private mailing list <security@cubefs.groups.io> for users reporting security vulnerabilities. Ref: [SECURITY.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/SECURITY.md)

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

  CubeFS currently holds monthly community meeting, integrated with [CNCF calendar](https://www.cncf.io/calendar/)
  Meeting minutes and recordings listed at: [CubeFS meeting schedule](https://github.com/cubefs/cubefs-community/wiki/Meeting-Schedule)

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

  CubeFS contribution workflow documented at: [CONTRIBUTING.md#workflow](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/CONTRIBUTING.md#workflow)

- [x] **Demonstrate contributor activity and recruitment.**

  - Contributor activities measured by devstats: [chubaofs.devstats.cncf.io](https://chubaofs.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Since%20joining%20CNCF&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-companies=All)  
  - Contributor activity measured by GitHub contributor dashboard: [The contributions of contributors](https://github.com/cubefs/cubefs/graphs/contributors)
  - Examples of recruiting new committers according to contributor's contributions:
    - Add [shuqiang-zheng](https://github.com/shuqiang-zheng) to Committers list: <https://github.com/cubefs/cubefs/pull/3384>
    - Add [zhangchuanqing](https://github.com/zhangchuanqing5658) to Committers list: <https://github.com/cubefs/cubefs/pull/3386>

  - Recruiting new contributors by participating in developer events
    - [Developer activity 2024](https://github.com/cubefs/cubefs/issues/3105)
    - [Developer activity 2023](https://github.com/cubefs/cubefs/issues/1920)

## Engineering Principles

### Suggested

N/A

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

  Project goal from project [README.md#what-can-you-build-with-cubefs](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/README.md#what-can-you-build-with-cubefs): 

  > As an open-source distributed storage, CubeFS can serve as your datacenter filesystem, data lake storage infra, and private or hybrid cloud storage. 
  > In particular, CubeFS enables the separation of storage/compute architecture for databases and AI/ML applications.
  > 
  > Some key features of CubeFS include:
  > 
  > - Multiple access protocols such as POSIX, HDFS, S3, and its own REST API
  > - Highly scalable metadata service with strong consistency  
  > - Performance optimization of large/small files and sequential/random writes
  > - Multi-tenancy support with better resource utilization and tenant isolation
  > - Hybrid cloud I/O acceleration through multi-level caching
  > - Flexible storage policies, high-performance replication or low-cost erasure coding

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

  CubeFS Introduction: [link](https://cubefs.io/docs/master/overview/introduction.html)

  According to [CubeFS doc](<https://cubefs.io/docs/master/overview/introduction.html#application-scenarios>), typical use cases are:
  - Big Data Analytics
  - Deep Learning/Machine Learning
  - Container Shared Storage
  - Database & Middleware
  - Online Services
  - Traditional NAS to Cloud

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

  CubeFS has a public roadmap doc at [ROADMAP.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/ROADMAP.md)

- [x] **Roadmap change process is documented.**

  CubeFS documents its roadmap rules and changing process in [GOVERNANCE.md#roadmap](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/GOVERNANCE.md#roadmap)

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

  - CubeFS architecture overview: <https://cubefs.io/docs/master/overview/architecture.html#architecture>
  - Design details are documented under the "Design" section, an example is: <https://cubefs.io/docs/master/design/master.html>

- [x] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:**

  - [x] **Release expectations (scheduled or based on feature implementation)**

    CubeFS documents their release frequency as needed (beta and official releases), which can be regarded as based on feature implementation.
  - [x] **Tagging as stable, unstable, and security related releases**

    CubeFS uses beta to mark their unstable releases. Ref: [RELEASE.md#types-of-releases](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/RELEASE.md#types-of-releases).

    Security release process is documented at: [security-release-process.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/security/security-release-process.md). CubeFS doesn't have explicit tagging rule for security releases. Though this is not required, tagging a release with "security-fixes-only" alike markers would be helpful for users to prioritize upgrades.

  - [x] **Information on branch and tag strategies**

    > | Type      | Versioning     | Branch               | Frequency                    |
    > | ----      | ----------     | ---------            | ---------                    |
    > | beta      | vX.Y.Z-beta    | release-X.Y.Z-Beta   | as needed (at branch time)   |
    > | official  | vX.Y.Z         | release-X.Y.Z        | as needed (post beta)        |

  - [x] **Branch and platform support and length of support**

    No specific description of platform supported, according to the [artifacts-included-in-the-release](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/RELEASE.md#artifacts-included-in-the-release), currently only amd64 binaries are maintained by the community.
    Length of support clearly documented, support latest 3 minor releases.

  - [x] **Artifacts included in the release.**
  - **Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.**

    Each release note specifies the artifacts included in the release. Example: [CubeFS v3.3.2 release notes](https://github.com/cubefs/cubefs/releases/tag/v3.3.2)

- [x] **History of regular, quality releases.**

  History of CubeFS releases and changelogs: <https://github.com/cubefs/cubefs/releases>

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

- [x] **Achieving OpenSSF Best Practices silver or gold badge.**

  CubeFS has achieved the OpenSSF Best Practices silver badge: <https://www.bestpractices.dev/projects/6232>

  [![OpenSSF Best Practices](https://www.bestpractices.dev/projects/6232/badge)](https://www.bestpractices.dev/projects/6232)

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

  CubeFS has a clear security vulnerability report guide at: [SECURITY.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/SECURITY.md#report-a-vulnerability)

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  - CubeFS commnity is currently WIP in the enforcement of two-factor authentication for all members: <https://github.com/cubefs/cubefs/issues/3487>.
  - DCO sign-off and review&approval by maintainers are required for all the incoming pull-request Ref: <https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/CONTRIBUTING.md#workflow>
    > Every pull request that merges code to the master branch needs to be approved by at least one core maintainer for code review and pass all checks (including the DCO check) before it can be merged.
  - CubeFS also enabled the following static and dynamic scanning, security scanning to help ensure the code quality:
    - [gofumpt](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/docker/script/run_format.sh)
    - golint:In file [docker-compose.yml:469](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/docker/docker-compose.yml#L469-L477)
    - gosec:In file [docker-compose.yuml:479](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/docker/docker-compose.yml#L479-L487)
    - Fuzz testing [cubefs: add base fuzzers cncf/cncf-fuzzing#387](https://github.com/cncf/cncf-fuzzing/pull/387)
  - CI integration includes ci-test-unit, ci-test-s3 and ci-sast: [ci.yml](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/.github/workflows/ci.yml)

- [x] **Document assignment of security response roles and how reports are handled.**

  The CubeFS [Security Release Process](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/security/security-release-process.md) documents response roles and process of handling reports.

- [x] **Document Security Self-Assessment.**
  
  <https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/security/CubeFS-self-assessment.md>

- [x] **Third Party Security Review.**

  - [x] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

  CubeFS has passed the Third Party Security Review, Ref: [CubeFS-Security-Audit-2023-report](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/security/CubeFS-security-audit-2023-report.pdf).
  All found issues have been fixed, ref: page4 in the report "Executive summary".
  Security advisories of the fixes: [link](https://github.com/cubefs/cubefs/security/advisories?state=Triage)

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

  CubeFS has achieved OpenSSF Best Practices passing badge: <https://www.bestpractices.dev/en/projects/6232>

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

  The [ADOPTERS.md](https://github.com/cubefs/cubefs/blob/206d5ddadf1f99abde6401b7aa18b57fc46e6bed/ADOPTERS.md) documents adopters with adoption level and success stories.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

  The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation.
  CubeFS has been adopted by a large base of end users, most of which prefer to remain anonymous.

- [x] **TOC verification of adopters.**

  CubeFS Maintainers provided the TOC sponsor with a list of 7 users who agreed to be interviewed for the Graduation Due Diligence process.
  Refer to the [Adoption portion of this document](#adoption) for details.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

  - Integration with Hadoop: <https://cubefs.io/docs/master/ecology/hadoop.html>
  - Integration with Kubernetes: <https://cubefs.io/docs/master/ecology/k8s.html>
  - Integration with Prometheus: <https://cubefs.io/docs/master/ecology/prometheus.html>
  - Integration with Grafana: <https://cubefs.io/docs/master/ecology/grafana.html>



#### Adoption

##### Adopter 1 - NetEase / Internet

NetEase has integrated CubeFS into its cloud-native infrastructure since 2020, using it as a distributed file system solution. After a thorough evaluation of various options, CubeFS was selected for its superior scalability and reliability. Currently in production, NetEase's deployment spans over 200 servers and manages more than 15 petabytes of user data.

The core features of CubeFS that NetEase finds most valuable are high availability, fault tolerance, and the ability to handle large data volumes efficiently. These functionalities have made CubeFS an integral part of their cloud operations. The project's compatibility with Kubernetes has also been a significant advantage.

During the implementation phase, NetEase found the documentation quite helpful, although there was an initial challenge in understanding the architecture. They received community support that helped them overcome these initial hurdles.

NetEase upgrades their deployment of CubeFS with one minor release per year and patch releases as needed. CubeFS releases are usually stable and upgrades can be done with automated scripts in most cases.

Looking ahead, NetEase intends to increase its engagement with the CubeFS community and explore additional features that could enhance their cloud services. They view the project as a valuable asset that has not only met their storage needs but also contributed to their learning and development in distributed systems.

The adoption of CubeFS has been a significant benefit for NetEase, enhancing its cloud services with a robust and scalable storage solution. The project's stability and clear scope and roadmap have built trust with NetEase, making it a reliable component for its cloud-native infrastructure needs.

August, 2024

Ref: [Adopter Interview - NetEase](./cubefs-adopter-interview-netease.md)

##### Adopter 2 - BEIKE / Housing Transactions and Services

BEIKE has been using CubeFS since early 2020, starting with version 2.3. Currently, they are using version 3.3.1 in a large-scale production environment with a total capacity of 30PB (triple-replica redundancy), and some legacy environments are still running on version 2.4.

The main motivation for adoption was CubeFS's overall read/write and concurrency performance, scalability, as well as the ability to seamlessly integrate with existing infrastructure.

The adoption resulted in saving storage costs and easier debugging and problem resolution. The disk throttling feature has been useful, and the fault tolerance and data consistency features have been crucial in maintaining service continuity, even during peak business periods.

During the migration from NFS to CubeFS, they faced challenges with decreased write performance for small files. They also met issues with the HDFS data integration tool due to lack of maintenance. And BEIKE found the CubeFS documentation to be comprehensive and the community support to be responsive, which has been instrumental in the successful integration and ongoing maintenance of the system.

The community is open and friendly, BEIKE mainly participates in discussions in the WeChat group and attends monthly meetings, and contributes back when possible. Looking ahead, BEIKE hopes to be more involved in feature planning, enhancement discussions, and foster deeper collaboration with other community members.

Overall, CubeFS's strengths include concurrent read and write operations and efficient bandwidth utilization. Areas for improvement include cleaner git commit history management, making features pluggable, and improving compatibility for upgrading historical versions. The community also needs to improve its promotion and outreach efforts.

September, 2024

Ref: [Adopter Interview - BEIKE](./cubefs-adopter-interview-beike.md)

##### Adopter3 - Live Streaming and Short Video

Adopter3 started to use CubeFS in the second half of 2022. It began with testing and was deployed for production in January 2023. The key reason for adopting CubeFS was its ability to handle concurrent writes and the capability to horizontally scale metadata nodes. These were major advantages compared to HDFS and CephFS, which had constraints in performance, maintenance, and handling small files.

CubeFS has been convenient to use with its FUSE mounting. It has been integrated with the Kubernetes clusters and is also used directly on physical machines in some scenarios. The community documentation has been helpful to the adopter for learning about APIs, RESTful commands, and the architecture. Some internal business units have transitioned from CephFS to CubeFS, benefiting from reduced maintenance time. The replacement of HDFS with CubeFS is in progress.

In the future, Adopter3 plans to use the EC (Erasure Coding) feature when the community's architecture becomes more stable. Currently, they are focusing on replacing HDFS. Adopter3 has 3 committers in the community, interacting with maintainers primarily through WeChat group discussions.

Overall, the CubeFS project is considered stable. However, there is room for improvement in areas such as covering more corner cases, promoting the project more effectively, and enhancing the format and content of issue and PR descriptions to facilitate better maintenance.

September, 2024

Note: Adopter3 preferred to remain anonymous. The interview details are kept in a private file for CNCF TOC review.

##### Adopter4 - JD.com / Online Retailing

JD.com has been using CubeFS since 2018 as the foundation for its entire storage infrastructure. This adoption supports a diverse range of unstructured storage needs across the company's retail operations, including middleware, online and offline business, big data, and AI training, etc. In production for more than 6 years, they have multiple clusters, with the largest one consisting of over 4,000 servers and a total storage capacity exceeding 300TB, serving over 1 million clients concurrently. They use the 2020 version from the community and manually backport bug fixes.

JD.com choses CubeFS because it is customizable for their specific scenarios and supports large-scale clusters. CubeFS supports operating on the same dataset with different protocols simultaneously, like POSIX and S3, while alternatives like CephFS support only one. The adopter has also evaluated MooseFS, however it didn't met their requirements in scalability and stability.

JD.com has found the project documentation invaluable, particularly the design documents, which facilitated a deeper understanding and customization of CubeFS. The adoption has led to significant value, including reduced maintenance and resource costs, and improved resource utilization, reaching up to 98% in some environments.

From the adopter's perspective, the CubeFS community has rational governance and responsive maintainers. JD.com participates by sharing internal solutions, having maintainers review PRs, and interacting via GitHub and WeChat groups.

JD.com has a positive view of the CubeFS community, with rational governance and responsive maintainers. JD.com participates by sharing internal solutions, having maintainers review PRs, and interacting via GitHub and WeChat groups. They will continue to actively engage, contribute solutions, propose features, seeing the relationship as mutually beneficial.

Overall, CubeFS has strengths such as handling extreme scenarios and meeting enterprise requirements. JD.com suggests that the project could benefit from more community events and wider adoption promotion to reach its full potential.

October, 2024

Ref: [Adopter Interview - JD.com](./cubefs-adopter-interview-jd.com.md)
