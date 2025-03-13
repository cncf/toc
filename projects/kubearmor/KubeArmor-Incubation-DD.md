# KubeArmor Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1326)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for KubeArmor

### Criteria Evaluation

_$TOCMEMBER conducted the due diligence of KubeArmor who applied for $LEVEL. The project [has/has not] completed the criteria that show its maturity at $LEVEL. The following criteria implementations are noteworthy to call out... $NOTABLES. The following actions were provided to the project that were considered blocking but since resolved... $BLOCKERS. The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project... $RECOMMENDATIONS._

### Adoption Evaluation

_The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY._

### Final Assessment

_[The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS]._

## Application Process Principles

### Suggested

N/A

### Required

- [X] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

This was completed within TAG Security's Policy Working Group and occurred on 09 JUN 2021, the recording can be discovered at https://youtu.be/W8PlIBXT1hA?si=nJZ5k9rlRcHFDjHp.

A second presentation was delivered on 09 OCT 2024, the recording can be discovered at 


- [X] **TAG provides insight/recommendation of the project in the context of the landscape**

The TAG provided the following recommendation on the [presentation issue](https://github.com/cncf/tag-security/issues/1372#issuecomment-2403861372):
> The project has seen significant growth and adoption since the original acceptance to Sandbox. Having conducted threat modeling and implemented a number of static analysis and security processes in the development lifecycle provides a healthy project stance for security.

- [ ] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The project's site does not explicitly convey vendor neutrality. The donating company's logo and statement of creation occurs at least twice on the home page. There is no mention of vendor neutrality or open source on the main site. The blogs heading on the site states "open source solutions for kubernetes security" but can easily be interpreted that KubeArmor is not open source. The only reference to Cloud Native is the community group link in the site header. Further, the inclusion of modelarmor on the kubearmor site directs users to accuknox but is listed on kubearmor's site as "open source". ModelArmor is in the KubeArmor org repo.  The project lacks web separation to ensure neutrality.

Additionally, there are several references to AccuKnox, owning, initiating, or driving aspects of KubeArmor that need reconciled for the project to be considered vendor neutral. The [Policy-templates repo](https://github.com/kubearmor/policy-templates?tab=readme-ov-file#-contributions) is one such example where, in addition to the reference of AccuKnox driven initiative, also shows a dependency with [AccuKnox's travis-ci for KubeArmor](https://github.com/kubearmor/policy-templates/blob/c955043c5d16b6935e45a7b14e123fc14e80e9ac/README.md?plain=1#L4).

- [ ] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
- While the project was onboarded on 16-NOV-2021, the KubeArmor security group is email address for `support@accuknox.com` persists and does not provide separation between the business/product/service, and the project (requirement for sandbox)

- [ ] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisifies the Due Diligence Review criteria.

- [X] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The KubeArmor project has a variety of documentation to assist adopters and developers, as well as contributors, working with the project. The Documentation is clearly divided in to discoverable headers, with distinction on use cases from the overall project docs and guides.

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [ ] **Clear and discoverable project governance documentation.**

The Project's governance is kept in their primary repo for the project but is not linked from the Readme or on the project docs. While discoverable and clear, the governance does not cover considerations like inactivity of a maintainer. It also indicates an "as-needed" basis for inviting new maintainers to the project but does not specifiy what those needs are. This may place the project in an unhealthy posture in the future if not proactive defining its maintainership needs and building a succession of future maintainers from its active contributor pool.

Recommend the project include a link to governance in the Readme as well as in the project docs so adopters and contributors better understand the governance model of the project as community-driven. Also recommend the project define as-needed and consult with contributor strategy to build diveristy of its maintainers through its governance.

- [ ] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The Project's goverance has only received [one update](https://github.com/kubearmor/KubeArmor/commit/3fe550b02af2b2eba9de0efecafc8b86d678501e) since creation, and does not reflect iterative improvements by application of governance. 

- [ ] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

The Project's governance does not contain sufficient information to validate its governance activities and the repository does not have any activities to reflect exercising their [described process to become a maintainer](https://github.com/kubearmor/KubeArmor/blob/30345613e48533ec00732e2e7a0382058e04b83b/GOVERNANCE.md?plain=1#L22).

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

The Project's governance does not include any references to neutrality and does not provide any guidance regarding neutrality for the Project's blog and website. It has been confirmed that the project's website DOES include a link for ModelArmor which takes users to Accuknox's website and is considered in conflict with CNCF's requirements for vendor neutrality. It is noted that the KubeArmor security group is email address for `support@accuknox.com` which does not meet expectations for neutrality and does not provide separation between the business/product/service, and the project (requirement for sandbox).

- [ ] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

The Project's governance references the process by which maintainers are nominated and confirmed. The project's contribution guide does not define how criteria for contributions to be accepted, but does provide enough information (such as DCO requirement) that it could be inferred. There is no information regarding how the project may make requests of the CNCF or how it makes changes to its governance or project goals (project-level decision making).

- [ ] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

The Project does not have any documentation regarding assignment, onboarding, and removal of specific role of function-based members and teams.

- [ ] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

The Project does not have any documentation regarding the maintainer lifecycle process.

- [ ] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Given there is no documentation regarding the maintainer lifecycle process, nothing can be demonstrated.

- [X] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

The project does not appear to have any subprojects at the time of application for incubation.

### Required

- [X] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

The Project's maintainer list includes eight maintainers, 4 from AccuKnox, 2 from CERN, 1 from Dankook University, and another identified as independent. Maintainers listed include links to their GitHub profiles. No domain of responsibility is defined as the project is small enough the project's other repos are considered are joint responsibilities. As the project continues to grow, this will change.

- [X] **A number of active maintainers which is appropriate to the size and scope of the project.**

Yes, for now but without changes to its operations the project may become unhealthy (victim of its own success). The project has eight maintainers, of which at least two are regularly active. The project's current backlog of PRs (91 including fix(deps)) and open issues (155) as of March 2025 suggest that the current activity is slowing down on the project and there will need to be a focused effort by the project to onboard new maintainers (moving contributors to maintainers, formalizing that governance more), re-activate current maintainers, and generally continue to make progress on the project in driving down the open issues and PRs against their defined roadmap.

- [ ] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

The project does not have documented roles beyond maintainers listed however the CODEOWNERS file calls out three default reviewers (indicating a new role which needs documented) and does not define process to add new individuals to this role. One individual defined in the CODEOWNERS is not listed as a maintainer and they appear to no longer be active in the project as of March 2024.

- [ ] **Document agreement that project will adopt CNCF Code of Conduct.**

The project has adopted a separate code of conduct that is modified beyond the CNCF CoC, and the current link to the CNCF CoC is broken (indicating an older version was used). The KubeArmor CoC does include reporting information specific to the project. 

The project needs to adopt the CNCF CoC and may add provisions above and beyond what the CNCF CoC details, and is encouraged to maintain reporting information with a mechanism for reporting in the event the complaint is against one of the maintainers. The CoC needs to also link to those maintainers so reporters know who to contact.

- [X] **CNCF Code of Conduct is cross-linked from other governance documents.**

The Project's Governance links to their current CoC file, last updated 4 years ago.

- [ ] **All subprojects, if any, are listed.**

The project does not define sub-projects and does not have a component/repo construct to assist adopters in understanding what is in scope of KubeArmor. There are some project repos that indicate a specialty sub-project (such as the policy-templates) which detail being "commuity-owned" but do not consider any additional governance or documentation associated with how they are maintained beyond annotation that it is an initiative from AccuKnox. This needs decoupled to provide clear separation between the Project, its Community, and the donating company. 

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [ ] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) --> 

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

Note: this section may be augemented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [ ] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) --> 

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

##### Adopter 1 - [Open Horizon](https://github.com/open-horizon) / Edge

_Joe Pearson (IBM), TAC chair at LF Edge_
JANUARY 2025

Open Horizon is a LF Edge Foundation project that provides a platform for managing the service software lifecycle of containerized workloads and related machine learning assets. Open Horizon has been working with the KubeArmor project for approximately 18 months to ensure platform adopters can gain confidence that workloads deployed on the platform are secure and operating within the defined policy. As Open Horizon is designed for bare metal hosts and clusters, they reached out to KubeArmor to provide security coverage across multiple environments. Open Horizon’s integration with KubeArmor is compatible with any container runtime running on bare metal.

The Open Horizon team was originally unfamiliar with the runtime security space but leveraged their network of colleagues and community members to initiate discovery of open source projects in the space that provide protection and multiple layers (such as the file system and operating system) through a simple, easy to use security policy. KubeArmor’s capabilities align well with the premise of how Open Horizon operates (through defined policies) and is compatible with their deployment policies.

Open Horizon’s KubeArmor integration is not officially finalized, as the two projects plan to provide two different deployment models. The first is delivery of the KubeArmor client as a containerized workload delivered by Horizon (near complete) and the second is by embedding KubeArmor into Open Horizon (development happening later this year). The intent is to make KubeArmor available to developers through these deployment models so regardless of which model is chosen, organizations adopting Open Horizon have the benefit of security policy enforcement out of the box when properly configured.

Open Horizon’s development and integration with KubeArmor is against the “latest” tag of the project and has plans to pin the version of KubeAmor for their embedded implementation.

Open Horizon’s collaboration with KubeArmor allows the teams to work independently on their respective code bases with a regularly occurring, dedicated subgroup meeting to ensure they’re aligned and working towards the same objectives, features, and functionality. Where KubeArmor requires changes to Open Horizon, primarily through updates to Open Horizon documentation, the KubeArmor team has submitted the changes following Open Horizon’s existing process, demonstrating collaborative participation in open source. Open Horizon shared that KubeArmor’s docs are very impressive and thorough, providing the Open Horizon team with some tips and tricks that are used to support generation of Open Horizon’s documentation. They also shared that KubeArmor’s current maturity and operation is similar to a Stage 2 project in LF Edge (akin to CNCF Incubating) albeit slightly early.

Engagements by Open Horizon with KubeArmor have been pleasant. Open Horizon shared the KubeArmor team is great to work with, primarily engaged in email and slack but they also make use of Open Horizon’s chat server (matrix) with leadership, senior developer/architect and developer level engagement in the subgroup meeting.

Open Horizon’s integration of KubeArmor has provided closure of several security gaps Open Horizon’s adopters were looking for, such as explicit access control enforcement through policies, but has not had sufficient time to ascertain measurable value. If the project were archived today, Open Horizon would need to find a replacement project to provide equivalent functionality for their adopters. 

Open Horizon shared that a common challenge with many open source projects developed by startups is achieving maintainer diversity and therefore is a key area of focus for the project to ensure survivability as well as meet expectations of a CNCF graduated project.

Open Horizon shared that KubeArmor generally is a well rounded project, encompassing features and capabilities typically not found in early maturity software. KubeArmor has explicit host and access protection as well as containerized workload protection. The project is currently exploring protection for workload connectivity, outside the host or cluster itself. In cases working with other cloud native projects to achieve this. Open Horizon currently has no additional plans with KubeArmor beyond embedding it within Open Horizon as shared earlier.


##### Adopter 2 - R6 Security / Cloud Security

_Zsolt Nemeth, R6 Security_
JANUARY 2025

R6 Security began adoption of KubeArmor in October 2023 before KubeCon+CloudNativeCon Chicago. R6’s architecture utilizes a series of tools that function as triggers with corresponding platform integrations to form security events that spawn automated defense mechanisms (reference Phoenix AMTD) for applications deployed to Kubernetes.

![R6 Phoenix Architecture](https://github.com/r6security/phoenix/blob/main/docs/img/architecture.jpg 'R6 Phoenix Architecture') Source: https://github.com/r6security/phoenix/blob/main/docs/img/architecture.jpg

Initially leveraging Prometheus, and adopting Falco for additional alerting, R6 Security expanded its triggers in November 2023 to take advantage of rulesets and policies for detecting unauthorized activity; this was prior to Falco’s recent response engine, Talon. 

R6 Security’s decision to include KubeArmor, Falco, and RAD Security (formerly KSOC labs) was largely driven by customer requests and adds event-based triggers to R6 Security’s own time-based trigger. Contributing to R6’s decision to expand beyond one project, in 2023 Falco was experiencing upstream support issues that made pure reliance unfeasible for the Company; these issues have since been resolved.  Combined, these technologies provided their customers with comprehensive detection through layered complementary and overlapping capabilities, ensuring very little, if any, adversarial or unauthorized activity is unnoticed. 

KubeArmor covers a small portion of their overall architecture, making up less than 8% of the overall detection. R6 Security has approximately 200-250 users on their freemium edition which includes KubeArmor and Falco, primarily located in south-eastern Europe. More than half of these users leverage KubeArmor in production. R6 Security also shared that approximately 2-3 customers have upgraded to the enterprise equivalent of KubeArmor, while many are still using the project version.

R6 Security defers to its customers for pulling updated versions of projects into their solution. It took R6 Security less than a day to integrate KubeArmor into their architecture, at a time when the project documentation was less robust and helm charts were not available. Since then, significant improvements have been made to the project’s documentation, particularly it’s API docs. R6 Security shared that the project is very responsive to requests, citing an issue a few months ago where one of the triggers was not functioning for their users. After working with the KubeArmor team to diagnose the source of the problem, it was confirmed the issue was on R6 Security’s implementation.

R6 Security shared much of the value their integration with KubeArmor has provided was in providing more robust coverage for their customers as requested. R6 Security provides its paying customers with additional security datasets from these triggers through their grafana dashboard.

If the KubeArmor project were archived now or in the future, R6 Security would likely fork the project to ensure it remained compatible with Kubernetes releases, but the overall impact would be minimal as R6 Security’s custom rules with Prometheus and the existing ruleset with Falco cover a lot.

R6 Security shared that KubeArmor overall is a robust project, but many differentiating features are available with the Enterprise edition from vendor Accuknox, who employs nearly all the maintainers. If and when these features and capabilities are contributed to KubeArmor, the project could experience more adoption, distinguishing itself from other runtime security projects in the ecosystem.  As a security project, the layered PR review process can sometimes slow down the incorporation of new features, but overall was the only remaining challenge R6 Security saw with the project once the documentation had been improved.

The biggest strengths of KubeArmor, as shared by R6 Security, are the complex policies that it allows to be set and configured. Providing the findings as a json file is great for users who wish to manipulate or combine this dataset with other sources like Prometheus.



##### Adopter 3 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR
