# Dapr Graduation Due Diligence

- [Dapr Graduation application issue](https://github.com/cncf/toc/issues/1354)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Graduation Evaluation Summary for Dapr

### Criteria Evaluation

[Karena Angell](https://github.com/angellk) and [Davanum Srinivas](https://github.com/dims) conducted the due diligence of Dapr who applied for Graduation. The project has completed the criteria that show its maturity at the applied level. 

The following criteria implementations are noteworthy to call out:

- An important [language agnostic](https://docs.dapr.io/concepts/overview/#developer-language-sdks-and-frameworks) project designed to simplify the development of [microservices applications](https://docs.dapr.io/concepts/overview/#microservice-building-blocks-for-cloud-and-edge) by providing building blocks for common patterns with a wide adopter base.
- Provides [abstractions](https://docs.dapr.io/reference/components-reference/) for Input/Output Bindings, Pub Sub, State Stores, Secret Stores, Name resolvers, Configuration stores, Middlewares, Workflow services etc with over 110 concrete swappable implementations in total.
- Security is a [fundamental property](https://docs.dapr.io/concepts/security-concept/) central to Dapr's design.
- Well defined and exercised build and [release process](https://github.com/dapr/community/blob/master/release-process.md).
- Implementation of the [Steering and Technical committee](Steering and Technical committee) to guide both the community governance and technical leadership to the project.
- The completion of the [external security audit](https://www.cncf.io/blog/2023/09/06/dapr-completes-2023-security-audit-increasing-enterprise-confidence/) by Ada Logics and the lack of any high or critical issues demonstrate a strong commitment to security. All seven identified issues have been addressed, which is commendable.
- Dapr has an complete [contributing guide](https://docs.dapr.io/contributing/), including information general guidance, contributing to docs, SDKs, info on the roadmap etc. This is complemented with a dedicated github repository [dapr/community](https://github.com/dapr/community) makes discovery of information quick and easy for new folks.
- The project makes sure that its [public meetings](https://github.com/dapr/community?tab=readme-ov-file#communication-and-discord) are accessible, with notes, videos, and easy to find meeting links. 

The following actions were provided to the project that were considered blocking but have since been resolved:
- TOC reviewers are [recommending](https://github.com/dapr/community/issues/530) some of the STC meetings to be public to give a chance for the STC to interact with the project maintainers/users.
- [Update the Code of Conduct](https://github.com/dapr/community/issues/538) to reference the CNCF Code of Conduct. Currently the Code of Conduct is adapted from the Contributor Covenant, however one of the incubation criteria is to explicitly reference and use the [CNCF CoC](https://github.com/cncf/foundation/blob/main/code-of-conduct.md).

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

- TAG Security recommends documenting [Security Self-Assessment](https://tag-security.cncf.io/community/assessments/guide/self-assessment/)
- TAG Security recommends a [joint assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/guide/joint-assessment.md) as well building on the self-assessment above. The existing work already done in Dapr including the several audits and documentation of the [threat model](https://docs.dapr.io/concepts/security-concept/#threat-model) will come in handy.
- TAG Contributor Experience recommends adopting some tooling such as [clowarden](https://github.com/cncf/clowarden), [prow & peribolos](https://docs.prow.k8s.io/docs/components/cli-tools/peribolos/), [sheriff](https://github.com/electron/sheriff) or other tool that can manage org membership & permissions in a more transparent manner.
- TOC reviewers [recommend](https://github.com/dapr/community/pull/540/files) a better definition of a "sub-project" as counting each repo as a sub-project ends up with too many of them. Having a better aggregation will help create named roles like "technical leads" which sets up folks better for going up the contributor ladder. TOC reviewers especially worry about "bus factor", having the next rung of leadership mentored and ready will alleviate the situation.
- TOC reviewers recommend [separation](https://github.com/dapr/community/issues/532) of duties between STC, Code of Conduct and Security related folks. Some work was already done in terms of the mailing list. Having distinct groups of folks responsible for growth of community going forward.
- TOC reviewers recommend introducing the [concept of Emeritus](https://github.com/dapr/community/issues/438) in various roles to acknowledge the work of folks who spent time in the community, but no longer are active as well as give a easy path for them to get active again when they choose to.
- The [Dapr component lifecycle certification](https://docs.dapr.io/operations/components/certification-lifecycle/) could use a single page for a quick lookup of component states in the lifecycle as well as a documented process to deprecate and eventually remove items that are no longer supported.
- The [Security Reporting Process](https://docs.dapr.io/operations/support/support-security-issues/) should be enhanced to include an [embargo](https://contribute.cncf.io/maintainers/security/security-guidelines/#34-embargo-policy) and private security notification process etc. In addition, we feel that tagging a release clearly as "security-fixes-only" will help users to prioritize an upgrade.
- Documentation updates to better cover Day 2 operations, more complex examples especially for people who are developers but not experts in cloud native were cited in some of the adopter interviews. Also a much clearer distinction/description of relationship with service meshes would be welcome as well (when do you use one vs another or one with another!).
- Dapr has grown by leaps and bounds beyond its .NET roots across Javascript, Python and Java based workloads and should continue to work towards making all the platforms be on par in terms of adoption. Other investments would be better upgrades, and simplified certificate rotation especially when users have lots of clusters to deal with (without needing an additional product/service)
- In terms of expanding beyond converting end users into contributors, it would be good to look at mentoring programs like GSoC, LFX etc as well as end users typically focus on what exactly they need and we need folks who look at the overall system and focus on the core as well.

### Adoption Evaluation

The adopter interviews reflect the project use for the Graduation level to which the project has applied. [Dapr](https://dapr.io/) simplifies cloud infrastructure interactions by providing developers a consistent API that targets multiple backend systems like databases, message brokers, configuration stores, and secret stores. Developers leverage Dapr to address complex distributed systems challenges, including durable workflows, business orchestration, and state conflict resolution. The project also aims to enhance application-level security rather than relying solely on network or infrastructure-level protections. Dapr enables platform teams to limit infrastructure access to specific apps and automatically creates data isolation per service when integrating with databases. Additionally, Dapr offers native authentication and authorization middleware for developers to use. In these terms, Dapr has brought substantial value to adopters, as it provided an opportunity to make distributed system development accessible to everyone. Overall, the stability of the components, sdks and their API builds a strong sense of trust with its adopters as well.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation.

### Criteria


## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**
 
- Presentation to [TAG App Delivery on Sept 18, 2024](https://github.com/cncf/tag-app-delivery/issues/748)
- Presentation to [TAG Security on Sept 25, 2024](https://github.com/cncf/tag-security/issues/1364)
- Offline review with [TAG Contributor Strategy filed on Oct 4, 2024](https://github.com/cncf/tag-contributor-strategy/pull/724)

- [x]  **TAG provides insight/recommendation of the project in the context of the landscape**

  In addition to the presentations mentioned above, Dapr project has presented to TAG App Delivery in the past as well (for example during [incubation process](https://github.com/cncf/tag-app-delivery/issues/130)) and is a subject of [discussions/whitepapers](https://tag-app-delivery.cncf.io/search/?q=dapr) in various working groups. TAG App Delivery and TAG Contributor Strategy are supportive of Dapr's progression to graduated status within CNCF. The project demonstrates maturity in its technology/deliverables and a strong commitment to security and community practices.

- TAG Security evaluation can be found in this [comment](https://github.com/cncf/tag-security/issues/1364#issuecomment-2376994865).

- TAG Contrib Strategy evaluation can be found in this [PR #724](https://github.com/cncf/tag-contributor-strategy/pull/724/files).

- TAG Application Delivery is represented in this due diligence by both TOC reviewers' domain expertise.

- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

  The project fulfills the vendor-neutral criteria. The main communication channel is [Discord](https://github.com/dapr/community?tab=readme-ov-file#communication-and-discord). Community meetings are held on [Zoom/Discord](https://github.com/dapr/community?tab=readme-ov-file#community-calls) with a bi-weekly cadence, showcased on the CNCF calendar and recordings are posted to youtube. Community call agenda is tracked as [issues in their community repository](https://github.com/dapr/community/issues?q=is%3Aissue+%22Community+Call%22) and notes are posted there as well.

  Regarding the architectural decisions, Dapr uses a [Proposals process](https://github.com/dapr/proposals) for any significant changes to design or updates to the SDKs etc. The community is welcome to submit proposals for consideration and discussion with the community and decisions are made there in public.

  Currently, the project has 21 maintainers affiliated with 8 companies: Microsoft 10; Diagrid 5; 1 each Nirmata, Stripe, Intel, Automattic, Independent, Proximus

  The project is self-governing, adhering to the principles of exercising minimal viable governance to enable community driven success. To further the sustainable development of the project, the  [Steering and Technical Committee](https://github.com/dapr/community/blob/master/steering-and-technical-committee-charter.md) manages to represent the voice of the users, direct the technical work and help guide development efforts towards solving concrete and immediate problems for Dapr users.

- [x] **Review and acknowledgement of expectations for [graduated](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
   - [x] Met during Project's application on 12-07-2024.

[Initial Issue](https://github.com/cncf/toc/pull/1168) was submitted on 18-09-2023. The project opened a [new issue](https://github.com/cncf/toc/issues/1354) using the recommended template on 18-06-2024.

  The TOC and project maintainers met during the kick-off meeting of the Graduation process to lay out the new process, what to expect, timelines etc. We created a shared, temporary channel with the TOC reviewers and the project leadership for communication. TOC reviewers opened public issues in various repositories asking questions and nudging the team to work on things all of which were cross linked to the main [issue](https://github.com/cncf/toc/issues/1354).

  The project maintainers have been highly responsive throughout the process.

  Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**
 
  - [Installation documentation](https://docs.dapr.io/getting-started/install-dapr-cli/)
  - [Getting started](https://docs.dapr.io/getting-started/)
  - [Tutorials](https://docs.dapr.io/getting-started/tutorials/)
  - [Developing Applications](https://docs.dapr.io/developing-applications/)
  - [Operations](https://docs.dapr.io/operations/)

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The project fulfills this criteria. Recently two maintainers without activity have entered PRs to be removed. Earlier in the year there was a [cleanup of STC/Maintainers](https://github.com/dapr/community/issues/456) as well. This activity indicates the project governance is being adhered to. The STC was [bootstrapped in 2021](https://github.com/dapr/community/pull/83) and there was an [election in 2023](https://github.com/dapr/community/tree/master/elections/2023-STC) as the members have 2 year terms.

### Required

- [x] **Clear and discoverable project governance documentation.**

Project Governance can be found [here](https://github.com/dapr/community/blob/master/steering-and-technical-committee-charter.md).

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

  The project is up to date with actual project activities, this includes iterations on the STC process and meetings, as well as its influence on the roadmap.

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  Decisions are made by maintainers who are actively involved in the project, and and there is a clear path outlined for people to become [maintainers](https://github.com/dapr/community/blob/master/community-membership.md#maintainer). Current maintainers are spread across several companies and the project is also encouraging new maintainers to join. Changes are reflected back into the [CNCF project-maintainers.csv](https://github.com/cncf/foundation/pulls?q=is%3Apr+dapr) periodically. 

  The project [roadmap](https://github.com/dapr/community/blob/master/roadmap.md) is based on [Proposals](https://github.com/dapr/proposals) process and public github [milestones](https://github.com/dapr/dapr/milestones) reflecting vendor neutrality at its core.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  This is included in project governance and contributors docs, including:
- [Community membership](https://github.com/dapr/community/blob/master/community-membership.md) outlines the responsibilities and privileges of contributor roles.
- [Release Process](https://github.com/dapr/community/blob/master/release-process.md) doc outlines what goes into a release and how the release team works on features etc.
- [Contributing flow](https://docs.dapr.io/contributing/) for new features, bugs and overall project management.
- [Feature Lifecycle](https://github.com/dapr/proposals?tab=readme-ov-file#feature-lifecycle-outline) delineates lifecycle of features/components and a defined set of milestones / requirements for progression between the lifecycle phases of those.
- [Steering and Technical Committee](https://github.com/cert-manager/community/blob/main/STEERING.md) charter has the formal process for governance and technical oversight for the project.
- A separate doc focused on [Community Manager role](https://github.com/dapr/community/blob/master/COMMUNITY-MANAGER.md) is available.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

- Onboarding documentation is available for [maintainers](https://github.com/dapr/community/blob/master/community-manager-onboarding-1.md) and for [community managers](https://github.com/dapr/community/blob/master/community-manager-onboarding-2.md).
- Five different levels of community roles are defined in the [project documentation](https://github.com/dapr/community/blob/master/community-membership.md). 

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

- List of all maintainers are in [MAINTAINERS.md](https://github.com/dapr/community/blob/master/MAINTAINERS.md).
- List of community managers are in [COMMUNITY-MANAGER.md](https://github.com/dapr/community/blob/master/COMMUNITY-MANAGER.md#current-community-managers)

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

  The 20+ maintainers take care of a lot of repositories between them that host different components, SDKs etc.

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

- The community repository has the documentation in [community-membership.md](https://github.com/dapr/community/blob/master/community-membership.md)

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

- You can see the additions and deletions in the history of the [MAINTAINERS.md](https://github.com/dapr/community/commits/master/MAINTAINERS.md) and the pull requests to the CNCF [project-maintainers.csv](https://github.com/search?q=repo%3Acncf%2Ffoundation+dapr&type=pullrequests&s=updated&o=desc)

- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

Current count there are 10 maintainers from Microsoft and 5 from Diagrid

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

- Yes, [github teams and CODEOWNERS](https://github.com/dapr/dapr/blob/master/CODEOWNERS) are used to provide access to code and documentation

- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

- [CODE-OF-CONDUCT.md](https://github.com/dapr/community/blob/master/CODE-OF-CONDUCT.md) has been updated to reflect that CNCF CoC is in effect.

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

- CODE-OF-CONDUCT.md currently has [25+ files linking to it](https://github.com/search?q=org%3Adapr%20CODE-OF-CONDUCT.md&type=code). 

- [x] **All subprojects, if any, are listed.**

Dapr treats each of its repositories [as a subproject](https://github.com/dapr/community/blob/master/community-membership.md).

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

The project currently lacks clarity on what constitutes a sub-project. Currently, each repo is considered a sub-project and therefore lists a volume of them. Having a better aggregation will help create named roles like "technical leads" which sets up folks better for going up the contributor ladder. It should further be noted in the event of substantial leadership departure ("bus factor"), having the next rung of leadership mentored and ready will alleviate any potential issues.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

- [5 levels](https://github.com/dapr/community/blob/master/community-membership.md) of code related contributors and a [separate role for a community manager](https://github.com/dapr/community/blob/master/COMMUNITY-MANAGER.md) is defined in governance documentation.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

- [contributing-overview](https://docs.dapr.io/contributing/contributing-overview/) is defined in the docs website.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

- [Discord](https://github.com/dapr/dapr?tab=readme-ov-file#community) is the platform of choice for online communications for all the participants.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

  In addition to the various channels on slack, `codeofconduct@dapr.io`, `stc@dapr.io` and `security@dapr.io` are some private mailing aliases to reach the respective folks.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

- [Bi-Monthly calls](https://tockify.com/cncf.public.events/monthly?search=dapr) are available in the CNCF calendar

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

- There are separate [Maintainer Guide](https://docs.dapr.io/contributing/docs-contrib/maintainer-guide/) and [Contributor Guide](https://docs.dapr.io/contributing/docs-contrib/contributing-docs/) in addition to a lot of documentation in the [dapr/community](https://github.com/dapr/community) repository.

- [x] **Demonstrate contributor activity and recruitment.**

- Devstats has some good charts to show activity like the [users-statistics-by-repository-group](https://dapr.devstats.cncf.io/d/48/users-statistics-by-repository-group?orgId=1) and [companies-contributing-in-repository-groups](https://dapr.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1)
- Github contributor graph shows [consistent activity](https://github.com/dapr/dapr/graphs/contributors) as well over time even as some people leave and others join.

## Engineering Principles

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

  The goal of the Dapr project is to codify the traditional 3-tier software architecture pattern - client/server/database - into a cloud native microservice architecture via software building blocks. The objective is to negate the traditional drawbacks of the 3-tier software pattern by enabling agility, redundancy, scale-out/scale-in, resiliency, platform independence and high availability through a cloud native software design.
  
  An additional objective is to simplify building microservice applications for traditional and cloud native application developers by providing standardized service building blocks and familiar SDKs for multiple programming languages - allowing organizations with teams that utilize disparate programming languages a lower barrier to entry for writing and maintaining cloud native microservice applications. 

  The Dapr project is currently targeted for medium to large enterprises. The most common adoption pattern being seen mostly in financial service, retail and healthcare. The target persona is application developers, who interact with Dapr via SDKs, Rest APIs or native gRPC clients.
  
  Dapr is an API driven system and the UX code is driven around SDKs, direct HTTP calls and gRPC clients.
  
  Dapr integrates with multiple CNCF projects: it uses gRPC as its communication layer, with Kubernetes it allows operators to configure CRDs to manage connections to underlying infrastructure. Dapr also integrates with NATS for pub/sub, and uses OpenTelemetry for generating metrics and traces.
  
  As options for installation, Dapr has: a Kubernetes Operator, a Helm chart, and a dedicated CLI that performs extra validation to make sure Dapr is deployed correctly. To test and validate the installation, an adopter may use Helm or the dedicated CLI to use the Kubernetes API to make sure the installation is valid.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

  Dapr reduces the complexity of interacting with cloud infrastructure such as databases, message brokers, configuration stores, secret stores etc. by providing developers with a consistent API that can target multiple infrastructure backends (database, message broker etc.). Developers use Dapr to tackle hard distributed systems challenges like durable workflows, business orchestration, state conflict resolution etc. 
  
  An additional goal of the project is to secure applications at the application level instead of relying on network or infrastructure level security. Dapr allows platform teams to scope access to infrastructure to specific applications. And when integrating with a database, Dapr automatically creates data isolation per service among others. 
  
  Dapr also includes native authentication and authorization middleware for developers to use.
  
  Although Dapr utilizes sidecars, the project has specified that it is not a [service mesh replacement](https://v1-13-1.docs.dapr.io/concepts/service-mesh/). According to Adopter interviews, there is still confusion on where Dapr fits in with Service Mesh technologies.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

  The Dapr project has a documented release process that includes a 2 week “feature definition stage” in which maintainers triage items based on community feedback into the milestone. This determines the short term roadmap for a given release. For long term items they have a dedicated proposal feature where big ticket items are opened by the community and maintainers, and are then discussed and voted on. Once an issue is approved, it will roll into a specific release during the feature definition stage.

  The [Dapr high-level roadmap](https://github.com/dapr/community/blob/master/roadmap.md) is maintained in the Dapr community repository and is updated quarterly. The current update cadence is not frequent enough, allowing for out of sync information - [tracking issue #4377](https://github.com/dapr/docs/issues/4377) in the Dapr docs repo has been created to track updates to this documentation. The current roadmap documentation also does not sufficiently reflect the current process and is being tracked by the same issue.

- [x] **Roadmap change process is documented.**

  [Tracking issue #4377](https://github.com/dapr/docs/issues/4377) in the Dapr docs repo has been created to track Roadmap process improvements including creating and documenting a roadmap change process. The current changes by the Dapr maintainers are sufficient to meet the Graduation requirement, however, all recommendations should be completed soon after Graduation.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

  The current project documentation thoroughly documents information for the application developer persona including the software design of the project. 
  
  The documentation outlines a basic [overview](https://docs.dapr.io/concepts/overview/) of the project including the software building blocks and how to use them.
  
  The project documentation also includes a section on [how to develop applications](https://docs.dapr.io/developing-applications/) and general [Operations guidance](https://docs.dapr.io/operations/).
  
  The documentation also refers to [how to contribute](https://docs.dapr.io/contributing/) to the project in multiple ways, not just through code contribution, but giving presentations, helping drive priorities, contributing to documentation and the various SDKs.
  
  Examples of cloud native use cases and integrations are included in the Ecosystem section below.
  
  A gap in documentation is for the SRE / Operations administrator persona in an easily accessible location within the documentation. Tracking issues [4383](https://github.com/dapr/docs/issues/4383) and [4384](https://github.com/dapr/docs/issues/4384) have been created in the Dapr docs repo to track recommendations to remediate the gaps.

- [x] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:** 

  - [x] Release expectations (scheduled or based on feature implementation)
  
  The project documents their release process in a [Release Process document](https://github.com/dapr/community/blob/master/release-process.md) that is linked to the [RELEASE.md](https://github.com/dapr/dapr/blob/master/RELEASE.md) file. Releases are aimed to be scheduled on a quarterly basis, however, the release cadence is more frequent with a healthy number of features, contributions and quick security fixes.
  
  Each release has a tracking issue that is tracking milestones per component to be added to the current release. Minor release example: [Dapr v1.10.09 Security release](https://github.com/dapr/dapr/issues/6685)
  
  Major releases also have a tracking board to provide further information and management for the release. Example: [Dapr v1.5.0 Tracking Issue](https://github.com/dapr/dapr/issues/8017) and [Dapr v1.5.0 Tracking Board]
  
  There is a discrepancy between the Support Release Policy and the Release Process - [Issue #8188 in the Dapr repo](https://github.com/dapr/dapr/issues/8188) has been created to track clarifications to each. 
  
  - [x] Tagging as stable, unstable, and security related releases

  Releases aren't specifically tagged as stable, unstable or security. They are [tagged](https://github.com/dapr/dapr/tags) as either a stable release or a "Release Candidate", which is inherently unstable as it is still undergoing testing prior to stable release. 
  
  However, a tracking issue may be specified as a security release in the tracking issue for the Minor release, but not tagged as such.
  
  Bug fixes are included in both Release Candidates and Stable releases. They can be included in a major or minor release alongside other features, or solely in a minor release and are documented thoroughly. Example: [Dapr 1.14.4 release](https://github.com/dapr/dapr/releases/tag/v1.14.4).
  
  [Issue #8188 in the Dapr repo](https://github.com/dapr/dapr/issues/8188) has been created to track tagging policy clarification.

  The current tagging strategy is sufficient to meet the requirements for Graduation with the recommendation to complete the issues identified soon after Graduation.
  
  - [x] Information on branch and tag strategies

  The Dapr project provides information on branch strategies in the [Maintainer guide](https://docs.dapr.io/contributing/docs-contrib/maintainer-guide/).

  The project does not currently provide information on tag strategies. [Issue #8188 in the Dapr repo](https://github.com/dapr/dapr/issues/8188) has been created to track adding tag strategies to the Dapr process documentation.

  The current information on branch strategies is sufficient to meet the requirements for Graduation. The recommendation is to add tag strategies shortly after Graduation to improve the end user experience.

  - [x] Branch and platform support and length of support

  The project specifies platform support and length of support in their [Support Release Policy](https://docs.dapr.io/operations/support/support-release-policy/) document. There is a discrepancy between the Support Release Policy and the Release Process - [Issue #8188 in the Dapr repo](https://github.com/dapr/dapr/issues/8188) has been created to track. 

  - [x] Artifacts included in the release.
  - Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.

  Each release note specifies the artifacts included in the release. Example: [Dapr 1.14.4 release notes](https://github.com/dapr/dapr/releases/tag/v1.14.4)

  - [x] **History of regular, quality releases.**

  The [Dapr releases page](https://github.com/dapr/dapr/releases/) shows a history of frequent, quality releases for Major, Minor and patch/security releases. The Dapr project responds quickly to identified bugs and CVEs and has a robust release process. At the time of publishing this due diligence, the Dapr project has 263 releases.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

  The project has not applied for the silver or gold OpenSSF Best Practices badge at this time.

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

  The Security reporting process is clearing outlined on the [Reporting security issues page](https://docs.dapr.io/operations/support/support-security-issues/) located in the Support section of the Dapr.io website. 

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  Members of the Dapr GitHub organization are [required to enable two-factor authentication](https://github.com/dapr/community/blob/master/community-membership.md) on their GitHub account. 
  
  Github teams are used to control access to the repositories, some of the important ones are:
- @dapr/approvers-components-contrib 
- @dapr/approvers-dapr
- @dapr/maintainers-components-contrib 
- @dapr/maintainers-dapr 

  The folks on the STC and folks with a special role called [ADMINISTRATORS](https://github.com/dapr/community/blob/master/ADMINISTRATORS.md#current-administrators) have github admin access to all repositories in the Dapr org. There is selection, onboarding and offboarding process for ADMINISTRATORS documented there as well.

- [x] **Document assignment of security response roles and how reports are handled.**

  For ease of use, adopters are requested to send an email to security@dapr.io and anyone on the [MAINTAINERS.md](https://github.com/dapr/community/blob/master/MAINTAINERS.md) file are authorized to access the security report.
  
  The Security reporting process also defines response SLAs as best effort within 3 days and once a security issue is identified as a vulnerability, the reviewer responds to the reporter with a timeline. From report to fix, the project maintainers attempt to resolve within 30 days.

- [x] **Document Security Self-Assessment.**

  The Dapr project has documented the project's [security posture](https://docs.dapr.io/concepts/security-concept/) and [threat model](https://docs.dapr.io/concepts/security-concept/#threat-model). The project has also presented to TAG Security. TAG Security has recommended completing a self-assessment and joint assessment although has not identified glaring security issues.
  
  The CNCF Graduation application was not clear on what "security self-assessment" meant so the project has put forth the security posture document and threat model (linked above), the third party fuzzing audit and security audits (linked in the next section), the [OpenSSF Scorecard](https://scorecard.dev/viewer/?uri=github.com/dapr/dapr) and [CLOMonitor](https://clomonitor.io/projects/cncf/dapr). 
  
  TOC reviewers do not view completing the TAG Security self-assessment as a blocker to Graduation given review of all presented materials.
  
  TOC issue [1451](https://github.com/cncf/toc/pull/1451) was created and merged to link to the TAG Security self-assessment for future applications.

- [x] **Third Party Security Review.**

  - [x] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

  The Dapr project went through a [Security Audit](https://docs.dapr.io/docs/Dapr-september-2023-security-audit-report.pdf) in September of 2023 with third party Ada Logics. The report identifies 7 issues which have been fixed and one that was fixed post-report.
  
  In June of 2023, the Dapr project went through a [Fuzzing Audit](https://docs.dapr.io/docs/Dapr-june-2023-fuzzing-audit-report.pdf), also with third party Ada Logics. The report identified 3 issues. 

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The project has a passing [OpenSSF Best Practices](https://www.bestpractices.dev/en/projects/5044) badge.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

A list of public adopters can be found in the [ADOPTERS.md](https://github.com/dapr/community/blob/master/ADOPTERS.md) file.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC reviewers with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

The project is used by a large number of organizations around the globe, many who publicly speak about their Dapr use and adoption. For the purpose of this due diligence, most adopters requested anonymity.

- [x] **TOC verification of adopters.**

Refer to the Adoption portion of this document.

Maintainers provided the TOC reviewers with a list of 7 adopters from different geographic regions who agreed to be interviewed for the Graduation Due Diligence process. 5 of these adopters were officially interviewed. The adoption portion of this document contains interview summaries from adopters who approved public use.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The Dapr project has extensive documentation outlining integrations and/or compatibility with Kubernetes as well as with other CNCF projects, including:
* [Autoscaling with KEDA, a CNCF Graduated project](https://docs.dapr.io/developing-applications/integrations/autoscale-keda/)
* [Using the gRPC interface, a CNCF Incubating project](https://docs.dapr.io/developing-applications/integrations/grpc-integration/)
* [Deploy Dapr with Helm, a CNCF Graduated project](https://docs.dapr.io/operations/hosting/kubernetes/kubernetes-deploy/)
* [Using the Open Telemetry Collector, a CNCF Incubating project](https://docs.dapr.io/operations/observability/tracing/otel-collector/)
* [Support declaration for Prometheus, a CNCF Graduated project](https://docs.dapr.io/operations/support/support-release-policy/)
* [Deploy to Kubernetes with Helm, a CNCF Graduated project](https://docs.dapr.io/operations/hosting/kubernetes/kubernetes-deploy/)
* [Hosting Dapr on Kubernetes, a CNCF Graduated project](https://docs.dapr.io/operations/hosting/kubernetes/)
* [Setup a Minikube cluster, a Kubernetes sub-project](https://docs.dapr.io/operations/hosting/kubernetes/cluster/setup-minikube/)
* [How to setup a KiND cluster, a Kubernetes sub-project](https://github.com/kubernetes-sigs/kind)
* [Dapr and OpenFunction, a CNCF Sandbox project](https://blog.dapr.io/posts/2022/05/14/how-dapr-helps-to-build-a-cloud-agnostic-faas-platform/)
* [Using Dapr with Linkerd or Istio, CNCF Graduated projects](https://v1-13-1.docs.dapr.io/concepts/service-mesh/)

Also compatibility with non-CNCF projects, including:
* [Distributed Tracing with Zipkin](https://github.com/dapr/quickstarts/tree/master/tutorials/observability)
* [Using a Kafka binding to Trigger Events](https://github.com/dapr/quickstarts/tree/master/tutorials/observability)

#### Adoption

##### Adopter 1 - Financial Services

August 2024

Adopter 1 has requested the adopter interview be kept private. The notes are kept on file for CNCF TOC review and inform the findings of this Graduation Due Diligence and recommendation to proceed to CNCF Graduation level.

##### Adopter 2 - Healthcare

August 2024

Adopter 2 has requested the adopter interview be kept private. The notes are kept on file for CNCF TOC review and inform the findings of this Graduation Due Diligence and recommendation to proceed to CNCF Graduation level.

##### Adopter 3 - Financial Services

August 2024

Adopter 3 has requested the adopter interview be kept private. The notes are kept on file for CNCF TOC review and inform the findings of this Graduation Due Diligence and recommendation to proceed to CNCF Graduation level.

##### Adopter 4 - Sharper Image / Retail

August 2024

Aaron Olds, Senior Software Architect at Sharper Image

Sharper Image is a retail company that also runs a commerce cloud. The Sharper Image team initially looked at the Critter stack which is an Asp.Net stack that required a major rewrite. The developers did not want a major rewrite and started looking at service meshes on the CNCF website and subsequently went to KubeCon NA Detroit in 2022 and made friends with the Dapr project team.

Aaron started at Sharper Image in August 2021. At that time, the microservices framework was in-house and the original Architect/Software Engineer left the company.  Aaron was tasked with replacing the in-house solution with a COTS alternative. Aaron found the “Learning Dapr” book and reached out to Yarun, core Dapr maintainer and architect. In March 2022, Aaron went to a hackathon with the Dapr project team and converted one microservice in a day, which was using gRPC.

Aaron had to do some work to convert to use Pub-Sub with Dapr. They were using services communication as well as a secrets store and some state store. They run all the communication for their backend services, which is a commerce cloud, in Dapr. So, when a customer hits “buy” - the Sharper Image team gets the cart and starts processing the order. They send notifications, process fullfillments and returns. It all runs in AWS using EKS.

The Sharper Image team currently runs Dapr 1.13 in production. They also have Dapr 1.14rc7 installed locally where they are testing features.

To orchestrate updates, Sharper Image uses Diagrid Conductor (paid product). They also use Datadog to monitor (had looked at Jaeger). The Sharper Image team had issues with some upgrades and posted to the community Discord instance. They verified the subscriptions were correct and when applying manifest files to the cluster they verified the subscriptions existed - it turns out they didn’t have auth set up correctly. In order to troubleshoot, they used Diagrid Conductor (paid product), however, they think they would be able to troubleshoot with an alternate tool that collects and monitor logs.

Aaron wrote a whitepaper for Diagrid on deploying their first application which was still focused on Diagrid Conductor. Aaron’s predecessor had written a CLI to deploy everything for the developers with more automation built into the homegrown app. However, now developers had to understand manifest files.

All secrets for the Sharper Image deployment are in an AWS secret store and they use the AWS client locally. Everything is installed in Dapr, so if they need to test communication, they port forward to a Dapr container.

They have mixed dev local environments with 7 developers in total. 5 are backend developers and 2 are frontend developers. Of the 7 developers there are 3 macs and 4 windows machines. They no longer deploy locally but deploy through a bitbucket pipeline

Aaron has tried to continue to abstract from the developers and has changed the implementation of the backend to use Dapr.

The processors and memory for the old framework used to take between 3-6 cores to run and up to 3-8 GB ram. It was inefficient, had memory leaks and used 20 nodes in Kubernetes with lots of taints. The current footprint with 1 Dev, 1 Test, and 2 Prod clusters is probably only 8 nodes.

The future with Dapr is currently uncertain however they will be doing more with Lambda and looking at using Diagrid’s other tool (Catalyst - API Manager). They want to be able to reach out to other services.

Aaron enjoys the open source aspect of the Dapr project submits discussions to Discord. He feels comfortable asking for features, however, it is unclear if his team submits issues/PRs or whether they are mostly submitted through the support contract they have with Diagrid.

Aaron believes the overall Governance of the project is good and someone(s) will continue to maintain the project if Diagrid is not available.

##### Adopter 5 - Healthcare

August 2024

ADOPTER 5 is a Healthcare company that sells SaaS-based solutions.

The decision to use Dapr was made mid-2021. The Dapr integration began at the experimentation phase and was chosen as a component in the technology stack modernization effort. The technology stack had been a .NET framework-based stack from approximately 2005. The Dapr API turned out to be sufficient when looking at modernizing the technology stack as the key drivers were:
* a need to transition to the cloud
* provide Service Discovery capabilities
* vendor agnostic 
* reducing the cognitive load on the application developers. 

There are currently 4 versions of Dapr in use. One team is using the latest release candidate. Not all teams update at the same time and the majority are on 1.11.x. If features are missing the teams update to a newer version.

Diagrid Conductor (a paid product) is used for upgrades. The Infrastructure-as-Code team can do updates manually if required.

When the IaC team sets up a cluster, they install Dapr by default. New projects have templates for Helm charts, Pub-Sub, Dapr configuration, etc and they focus on making it easy to pull functionality in via .NET plugins/manifests and dependency injection. ADOPTER 5 is primarily a .NET shop while some developers use Node as well.

ADOPTER 5 uses the Dapr GitHub repo to look at issues and check for error messages from the Dapr logs for troubleshooting. Also, they look for autocreate topics for Pub-Sub brokers. ADOPTER 5 has posted messages in the Dapr Discord channel 4-5 times as well as commented on a couple of GitHub issues and find that while GitHub usually has good discussions, the Discord community experience is mixed.

ADOPTER 5 has a contract with Diagrid for Production support, which also provides help with upgrades.
Having a contract helps with timely responses. Other community members are doing good work, so the project will continue if Diagrid developers aren’t available. 

It’s a bit early in ADOPTER 5's production deployment to measure value. Only 3-4 applications are using Dapr right now and it’s early in their modernization effort. However, the Pub-Sub functionality is already showing benefits and they were able to standardize on Apache Kafka which was as simple as changing a yaml file.

In the future, ADOPTER 5 would like to run all of the newer services with Dapr. The team is looking to rewrite the current software (macro-lith) to micro services. So, approximately 90-95% of services will start using Dapr including Pub-Sub, secrets and workflow. 

ADOPTER 5 does not have an open source culture so the team is not looking at contributing upstream. 

The community itself seems to be growing as new folks are presenting all the time in the bi-weekly calls and there are new attendees coming in from other companies. The community appears to be a relatively diverse, responsive group. Project maintainers are pretty responsive in GitHub, mixed in Discord.

ADOPTER 5 sees the main project strengths as having a broad feature set, cloud portability, built-in Service Discovery, and robust resiliency policies. 

In order to achieve its full potential, ADOPTER 5 would like to see Workflows move out of alpha and to reduce the time it takes features to reach production quality. Also, ADOPTER 5 would like to see something for out-of-cluster Service Discovery.
