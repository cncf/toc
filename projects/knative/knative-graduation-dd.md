# Knative Graduation Due Diligence

- Link to [Graduation application issue](https://github.com/cncf/toc/issues/1509)
- Based on the [graduation DD template](https://github.com/cncf/toc/blob/main/operations/toc-templates/template-dd-pr-graduation.md)
- Project Repo(s): github.com/knative, github.com/knative-extensions
- Project Site: https://www.knative.dev
- Sub-Projects: Serving, Eventing, Functions, Client
- Communication: https://cloud-native.slack.com/ #knative and various channels with #knative- prefix
- Project points of contacts: steering@knative.team

## Graduation Evaluation Summary for Knative

### Criteria Evaluation

Lin Sun & Jeremy Rickard conducted the due diligence of Knative who applied for graduation. The project has completed the criteria that show its maturity at graduation. The following criteria implementations are noteworthy to call out:
- Knative is a mature project that are endorsed by multiple adopters.
- The project is not only vendor neutral but also has a very diverse set of maintainers, adopters and integrators. The steering committee is composed of members from 4 different set of organizations.
- Knative documentation is fairly clear and easy to follow.
- The project does an excellent job of making sure that its working group public meetings are well documented.

The following actions were provided to the project that were considered blocking but since resolved:

- Add a [governance.md file](https://github.com/knative/community/pull/1685) to the community repo to make it visible and easy discover.
- Clarify [subproject documentation](https://github.com/knative/community/pull/1703)
- Clarify [maintainer affliation](https://github.com/knative/community/pull/1694).
- [Updated maintainer list](https://github.com/knative/community/pull/1700) to be accurate.
- Achieved [100% passing on OpenSSF best practice badge](https://www.bestpractices.dev/en/projects/5913#analysis).
- [Removed mentioning of community meetings](https://github.com/knative/community/pull/1704) as it is not active any more.

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

- [GTR review](https://github.com/knative/community/issues/1706) is being started by the project.
- Asked the knative team to [add more production readiness documentation](https://github.com/knative/community/issues/1707).

### Adoption Evaluation

The adopter interviews showed project usage at a level compatible with CNCF graduation. Adopters interviewed deploy Knative in production systems as part of managed platform offerings, supporting a variety of workload types. The Serving sub-project tas the most adoption and was cited repeatedly for it's scalability and stability, and have realized value in terms of cost-savings. 

Adopters indicated that Knative is quite mature for the use cases that it solves, has an active community, and has reliable documentation. Several adopters expressed a desire to see additional documentation around advanced administrative concepts and several additionally highlighted an opportunity for Knative to evolve to support emerging workloads within the A.I. space, particularly around supporting new developments in inference serving. 


### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation.

### Criteria

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

TAG Runtime - Presentation Feb 3rd, 2022 ([notes](https://docs.google.com/document/d/1k7VNetgbuDNyIs_87GLQRH2W5SLgjgOhB6pDyv89MYk/edit?tab=t.0#heading=h.wjqgcvq8889g), [video](https://www.youtube.com/watch?v=Qt--cUJOaQY&feature=youtu.be))
TAG App Delivery - Presentation Nov 8th, 2023 ([slides](https://docs.google.com/presentation/d/1ZxsRVTYW2vdlqQqVkLTQupLIDe4Ie8aBzRy4PNDaxmY/edit#slide=id.p)). Note this was in person during KubeCon thus no recording.

- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

Searched the GitHub repos, browsed through the project website, all look vendor neutral.

- [X] **Review and acknowledgement of expectations for [graduated](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
   - [X] We [met](https://github.com/cncf/toc/issues/1509#issuecomment-2844841621) with the project lead and went over the expectation and requirements for graduated project on 01-May-2025.

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [X] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

Knative the installation doc explained multiple ways to install clearly. Saw plenty of code samples. Saw limited production recommendation doc and asked the team to add more.

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

[Governance review for the Knative project](https://docs.google.com/document/d/1kwlizyag6EsMn9GECYSam-beSosQZhLj5Q1R_FLYQyo/edit?tab=t.0) was performed in May 2025 and it looks reasonable and healthy for the project. No major must fix issues were flagged. 

### Required

- [X] **Clear and discoverable project governance documentation.**

The TOC was able to discover the [governance file](https://github.com/knative/community/blob/main/GOVERNANCE.md) easily after suggesting the team to add one.

- [X] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Based on the governance file, the TOC was able to locate the leadership team, recent meetings with notes and recording, past election results.

- [X] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

It is clearly documented in the governance doc: https://github.com/knative/community/blob/main/GOVERNANCE.md

- [X] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

The Knative [steering committee doc](https://github.com/knative/community/blob/main/STEERING-COMMITTEE.md) has detailed explanations on how project elects leaders, updates to the charter doc and make decisions, who is elegible to vote etc.

- [X] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

[Community roles doc](https://github.com/knative/community/blob/main/ROLES.md) explain different roles and what to do when a given role becomes inative. Also the project has clear [security policy](https://github.com/knative/community/tree/main?tab=security-ov-file#readme) defined.

- [X] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

Both [steering committee](https://github.com/knative/community/blob/main/STEERING-COMMITTEE.md#committee-members) and [working group leads](https://github.com/knative/community/blob/main/working-groups/WORKING-GROUPS.md#working-groups) are documented. The entire [list of maintainers](It is documented here: https://github.com/knative/community/blob/main/MAINTAINERS.md) is documented as well.

- [X] **A number of active maintainers which is appropriate to the size and scope of the project.**

For the current cadence of changes in the project and backlog of work, the project has sufficient active maintainers to sustain its current and future momentum.

- [X] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

This is [documented](https://github.com/knative/community/blob/main/ROLES.md#approver) well including onboarding, offboarding and emeritus status.

- [X] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

From a few most recent PRs, this appears to be managed well. For example:
- Remove someone off a working group lead: https://github.com/knative/community/pull/1682 
- Add someone to a working group lead: https://github.com/knative/community/pull/1677

- [X] **Project maintainers from at least 2 organizations that demonstrates survivability.**

Confirmed that knative has maintainers from at least 2 organizations based on the [steering committee](https://github.com/knative/community/blob/main/STEERING-COMMITTEE.md#committee-members) and [working group leads](https://github.com/knative/community/blob/main/working-groups/WORKING-GROUPS.md#working-groups) documents.

- [X] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

Based on my limited browsing of the repository, it does seem to match, nothing is flagged. 

- [X] **Document agreement that project will adopt CNCF Code of Conduct.**

This is documented in the home page of the community repo: https://github.com/knative/community

- [X] **CNCF Code of Conduct is cross-linked from other governance documents.**

Yes, this is documented at https://github.com/knative/community/blob/main/CODE-OF-CONDUCT.md. Also linked from https://github.com/knative/community/blob/main/GOVERNANCE.md and https://github.com/knative/community/blob/main/ROLES.md as well as [Steering charter](https://github.com/knative/community/blob/main/STEERING-COMMITTEE.md#charter)

- [X] **All subprojects, if any, are listed.**

Subprojects are managed by Knative working groups (https://github.com/knative/community/blob/main/working-groups/WORKING-GROUPS.md).  The three main (user-facing) subprojects are:
* Serving
* Eventing
* Functions
Which are documented as top-level tabs on https://knative.dev/ (though we are also doing a tech docs assessment and landing page redesign to help highlight these more clearly).

A PR has been opened to add additional subproject info on this.  https://github.com/knative/community/pull/1703


- [X] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

All repos in knative and knative-extensions are owned by a working group (see above), with the following process for requesting a new repository or promoting an extension repository to core:
https://github.com/knative/community/blob/main/REPOSITORY-GUIDELINES.md 

Working group leadership is managed by the process in https://github.com/knative/community/blob/main/ROLES.md, with current leads in https://github.com/knative/community/blob/main/working-groups/WORKING-GROUPS.md (and Peribolos used to manage actual role assignments)

Each extension repo should have its own reported maturity status, as described in https://github.com/knative/community/blob/main/REPOSITORY-GUIDELINES.md#additional--experimental-repositories

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [X] **Clearly defined and discoverable process to submit issues or changes.**

To request a new feature or propose changes to roadmap, users can use this link: https://github.com/knative/community/blob/main/mechanics/FEATURE-TRACKS.md. To change the charter, it is documented in the steering doc.

- [X] **Project must have, and document, at least one public communications channel for users and/or contributors.**

Yes, knative has a slack channel in CNCF. It was quite active.

- [X] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

Social media related discussion are documented here: https://github.com/knative/community/blob/main/SOCIAL-MEDIA.md
Working group meeting and community meeting are also documented in the community repo: https://github.com/knative/community?tab=readme-ov-file#meetings-and-work-groups
Slack guideline with channel info are listed here: https://github.com/knative/community/blob/main/SLACK-GUIDELINES.md

- [X] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

This is very well documented in https://github.com/knative/community/blob/main/CALENDAR.MD.

- [X] **Documentation of how to contribute, with increasing detail as the project matures.**

Details on how to contribute are documented [here](https://github.com/knative/community/blob/main/CONTRIBUTING.md).

- [X] **Demonstrate contributor activity and recruitment.**

The project has run quite a bunch of [LF mentorship programs](:https://github.com/cncf/mentoring/tree/main/programs) and other programs to recruit more contributors. There's lots of Knative projects if you look at the LF mentorship program list, multiple projects every year:

* CNCF - Knative: Design and Implement Levels for Educational Game (2025 Term 1)
* CNCF - Knative: Improve Knative Eventing Onboarding (2024 Term 2)
* CNCF - Knative: Applying pre-prepared website design (2024 Term 2)
* Knative - Contributor Journey Research (2024 Term 1)
* Knative - Cross Namespace Event Links (2024 Term 1)
* CNCF - Knative: Porting Knative Serving to Microshift (2023 Term 2)

Prior to the CNCF acceptance, the project also participated Google Summer of Code directly:
https://github.com/knative/community/blob/main/google-summer-of-code/gsoc-2022.md
https://github.com/knative/community/blob/main/google-summer-of-code/gsoc-2023.md


## Engineering Principles

### Suggested

N/A

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

The goal of Knative is to provide a platform-agnostic solution for running serverless deployments. It does this with two main subprojects: _serving_ and _eventing_. 

The serving project is intended to provide a set of Kubernetes Custom Resource Definitions (CRDs) that  provide a common toolkit and API framework for serverless workloads. 

The goal of the Knative Eventing subproject is to define common primitives to enable composing event-processing applications through configuration, rather than application code.

Another Knative objective is to provide loose coupling of all components. It provides an robust [duck typing](https://knative.dev/docs/concepts/duck-typing/#example) system for it's components.


- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

There is a good overview of the [serving components](https://github.com/knative/specs/blob/main/specs/serving/overview.md) in the project Github repository. One of the primary goals of Knative is to provide a capability to scale workloads to zero when there is no demand for them. This feature helps reduces costs and helps better utilize resources. This use case was identified by all adopters that were interviewed as part of the due diligence process.

The eventing subproject builds on the [Cloud Events Specification](https://github.com/cloudevents/spec) in order to enable loose coupling and late-binding of event producers and consumers. The [components](https://github.com/knative/specs/blob/main/specs/eventing/overview.md) of and [motivation](https://github.com/knative/specs/blob/main/specs/eventing/motivation.md) for eventing is clearly documented. The ability to integrate legacy systems in an event based architecture with Knative was identified as a key feature by more than one adopter interviewed for this due diligence.

The project documentation also has an extensive [testimonials](https://knative.dev/docs/about/testimonials/) and [case studies](https://knative.dev/docs/about/case-studies/) highlighting viable cloud native use cases.


- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The project maintains public roadmaps in Github using Github project boards.

Each subproject maintains a separate roadmap. For example:

[Serving](https://github.com/orgs/knative/projects/53)
[Client](https://github.com/orgs/knative/projects/37/)
[Eventing](https://github.com/orgs/knative/projects/38)
[Functions](https://github.com/orgs/knative/projects/49)

- [x] **Roadmap change process is documented.**

The project maintains a [consistent set of instructions](https://github.com/knative/community/blob/main/mechanics/ROADMAPS.md) for maintaining roadmaps across subprojects. 

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

The Knative documentation provides a basic [overview](https://knative.dev/docs/concepts/) of Knative concepts. Further detailed information can be found about [eventing](https://github.com/knative/specs/blob/main/specs/eventing/overview.md) resources and [serving](https://knative.dev/docs/serving/) resources. The documentation also provides relevant digrams depicting [serving architecture](https://knative.dev/docs/serving/architecture/) and [request flow](https://knative.dev/docs/serving/request-flow/) that provide additional supporting evidence of the cloud native use cases the project supports. 

The serving and eventing documentation includes topics targeted toward both [developers](https://knative.dev/docs/eventing/event-delivery/#configuring-broker-event-delivery) and [administrators](https://knative.dev/docs/serving/knative-kubernetes-services/). 


- [x] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:** 

  - [x] Release expectations (scheduled or based on feature implementation)
  - [x] Tagging as stable, unstable, and security related releases
  - [x] Information on branch and tag strategies
  - [x] Branch and platform support and length of support
  - [x] Artifacts included in the release.
  - Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.

Knative documents the project release processes and guidelines in several files:

[RELEASE SCHEDULE](https://github.com/knative/community/blob/8921c0ad7fff98bc1a5d357910ae4fe806f09aa0/mechanics/RELEASE-SCHEDULE.md)

[RELEASE VERSIONING PRINCIPLES](https://github.com/knative/community/blob/8921c0ad7fff98bc1a5d357910ae4fe806f09aa0/mechanics/RELEASE-VERSIONING-PRINCIPLES.md)

These documents provide a good overview on the release schedule, how branches are used and supported, and what end users can expect in terms of API version compatability and compatability with Kubernetes. 

While this documentaiton does not detail what artifacts are included in a release, the project uses Github for releases, including for Serving manifests and artifacts are easily observed. GCR is used for container image releases and released manifests use SHA references. 


- [x] **History of regular, quality releases.**

Each of the subprojects within Knative utilize Github to publish releases and have a lengthy history of releases. KNative Serving had a v1.0.0 release in November 2021 and has had regular releases since then. Releases appear to have good release notes and clearly indicate breaking changes and bug fixes. 

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

The project has a [SECURITY.md](https://github.com/knative/community/blob/main/SECURITY.md) file defined in the community repo. 

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

Two factor authentication is required for the [knative](https://github.com/knative) org as well as the [knative-extensions](https://github.com/knative-extensions) org in Github.

- [x] **Document assignment of security response roles and how reports are handled.**

The project has a [vulnerability disclosure response policy](https://github.com/knative/community/blob/main/working-groups/security/responding.md) that outlines the Product Security Taskforce, as well as the process for fixing an issue identified as a vulnerability and general timelines for remediation.

The project also has a documented policy for [early disclosure of security vulnerabilities](https://github.com/knative/community/blob/main/working-groups/security/disclosure.md) for distributors of Knative.

- [x] **Document Security Self-Assessment.**

Knative has completed a [security self-assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/knative/self-assessment.md).

- [x] **Third Party Security Review.**

  - [ ] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

Knative underwent a [Third Party Security Audit](https://github.com/knative/docs/blob/main/reports/ADA-knative-security-audit-2023.pdf) by Ada Logics in 2023. The audit found 16 security issues, 15 of which were fixed with upstream patches. The remaining issue was not addressed, because it was only exploitable if fundamental security assumptions of the cluster were already broken.

- [X] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The passing badge is currently shown as 100%.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

A list of public adopters can be found in the [ADOPTERS.md](https://github.com/knative/community/blob/main/ADOPTERS.MD) file.


- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC with a list of adopters for verification of use of the project in production. The project is used by a many global organizations and the TOC was able to verify confirm production use with all interviewed adopters.  

- [x] **TOC verification of adopters.**

The Knative maintainers provided the TOC with a list of 20 adopters from different geographic regions and business segments who agreed to be interviewed for the Graduation Due Diligence process. 5 of these adopters were interviewed. The adoption portion of this document contains interview summaries from adopters who approved public attribution. All adopters recommended Knative for graduation and commented on project maturity. Scalability and stability were common strengths identified by adopters. The Knative Serving project especially was highlighted as providing a great deal of value to adopters.

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The Knative documentation highlights a number of integrations with Kubernetes as well as with other CNCF projects, including:

[Istio](https://knative.dev/docs/install/installing-istio/) 
[Helm](https://knative.dev/docs/install/operator/knative-with-operators/#install-via-helm-option-2)
[Contour](https://knative.dev/docs/install/yaml-install/serving/install-serving-with-yaml/#__tabbed_1_3)
[NATS](https://knative.dev/docs/install/yaml-install/eventing/install-eventing-with-yaml/#__tabbed_1_3)
[cert-manager](https://knative.dev/docs/install/installing-cert-manager/)
[Backstage](https://knative.dev/docs/install/installing-backstage-plugins)
[Open Telemetry](https://knative.dev/docs/serving/observability/metrics/collecting-metrics/#about-opentelemetry)

Additionally, Knative integrates with non CNCF projects such as:

[Apache Kafka](https://knative.dev/docs/install/yaml-install/eventing/install-eventing-with-yaml/#optional-install-a-default-channel-messaging-layer)
[rabbit mq](https://knative.dev/docs/install/yaml-install/eventing/install-eventing-with-yaml/#__tabbed_2_3)

Several of these integrations were mentioned during adoption interviewes as well.


#### Adoption

##### Adopter 1 - YMeadows

Y Meadows provides AI-Driven Automation for Business Operations, such as order operations, customer inquiries, and account management.

YMeadows was interviewed as a Knative adopter on June 4, 2025. They have been using Knative in production for about five years, supporting all of their customers. 

YMeadows originally adopted Knative for it's "scale to zero" features in order to control costs. The open source nature of the project, along with it's minimal dependencies, and active community presence were also strong factors in adopting the project.

When adopting Knative, YMeadows encountered some scaling issues, which were resolved by replacing the ingress component they were using. They generally found the documentation for the project to be adequate, however they had some issues related to supporting high availability and to address timeouts due to the size of workloads being used. 

YMeadows has a positive opinion of the Knative community and views the project as having open and transparent governance, as well as a helpful community with responsive maintainers. 

YMeadows would like to see internal TLS as a default configuration, more support for Kourier as a default networking plugin, and a smoother operator upgrade experience.

Overall, YMeadows recommended Knative for graduation and highlighted the cost reduction benefits the project helped deliver and highlighted the stability, scalability, documentation, and alignment with the general Kubernetes philosophy as overall strengths of the project.   

The entire adopter interview can be found here: [YMeadows Adopter Interview](knative-adopter-interview-ymeadows.md)

##### Adopter 2 - SVA System Vertrieb Alexander GmbH

Note: This adopter uses Knative via the Red Hat OpenShift Serverless product for compliance reasons. 

SVA System Vertrieb Alexander GmbH is one of the leading system integrators in Germany.

SVA System Vertrieb Alexander GmbH was interviewed as a Knative adopter on June 18, 2025. They adopted Knative in 2022, and are currently running the project in production, handling over a million events daily with 99.9% availability across 11 organizational units.

SVA System Vertrieb Alexander GmbH chose Knative for its open standards, clean abstraction layers, extensible architecture, and strong governance. The plug-and-play nature and integration with messaging systems like Kafka were key factors. 

SVA System Vertrieb Alexander GmbH found Knative easy to adopt and saw dramatically reduced onboarding times for customers and were able to implement improved compliance controls in regulated environments. When adopting Knative, SVA System Vertrieb Alexander GmbH faced some challenges relating to documentation, particularly around administrative topics. 

SVA System Vertrieb Alexander GmbH indicated they have a positive perception of the governance of Knative, although they would like to see increased contributor diversity. 

Overall, they recommended the project for graduation in part due to it's stability, performance, ease of adoption. 

The entire adopter interview can be found here: [SVA System Vertrieb Alexander GmbH Adopter Interview](knative-adopter-interview-sva.md)

##### Adopter 3 - Gojek

Gojek is a technology company in Indonesia, offering ride-hailing, e-commerce, and delivery services to millions of users across Indonesia and Singapore.

Gojek initially started using Knative in 2020 via the KServe project, before migrating directly to Knative. It is currently used in production at a very large scale, supporting millions of users and 100,000+ requests per second at peak times. 

Gojek adopted Knative while building a self-serve ML model-serving and experimentation platform for data scientists. They chose Knative for it's ease of use, flexibility, and because it was an active open-source project. 

While adopting, they had a positive experience overall. They saw immediate value and were able to deliver a platforms that were scalable, flexible, and had stable APIs. After the COVID-19 pandemic, cost optimization became a focus and they wanted to explore vertical autoscaling, but Knative primarily supported horizontal scaling at the time. They proposed a feature request to integrate VPA into Knative, but it wasn't finalized. This was not due to the project's unwillingness; rather, the complexity of integrating it into Knative was very high. 

Gojek views the project as mature and effective, and did not offer any areas for improvement. 

The entire adopter interview can be found here: [Gojek Adopter Interview](knative-adopter-interview-gojek.md)


##### Adopter 4 - Cloud Service Provider

Adopter 4 provides a platform for AI workloads. They elected to provide this adopter interview anonymously.

Adopter 4 has run Knative in production for about four years, as part of their hosted platform. When building their platform, they wanted to provide the following features:

* Autoscaling
* Throughput and latency for metrics
* Scale to zero 

Knative provided these capabilities, which meant that they did not have to build a solution on their own. Adoption for them was straightforward, with very few issues. Problems they did encounter were easily to identify and resolve using project documentation and by reviewing Github issues. 

Adopter 4 does not actively participate in the community, finding they rarely need to engage with the community because the project is mature. Adopter 4 did comment that while the documentation is generally good, and in some areas excellent, some advanced topics were harder to find in documentation adn required some review of code and markdown files in order to find answers to some questions.

Overall, Adopter 4 feels that Knative is very mature for the problems that it solves, but commented that Knative might not be the right solution for emerging workloads and the project should consider how to evolve to support changes in the space. 

The entire adopter interview can be found here: [Adopter 4 Adopter Interview](knative-adopter-interview-adopter-4.md)

##### Adopter 5 - IBM Cloud

IBM Cloud offers a serverless platform called IBM Cloud Code Engine, that supports multiple containerized workload types, and is partially based on Knative.

Within IBM Cloud Code Engine, Knative has been used in production for over five years, and supports thousands of Knative services per cluster and region. The platform primarily uses Knative Serving, but also some features from Knative Eventing.

IBM Cloud chose Knative because of it's Kubernetes-native design, as well as it's support for scale to zero for workloads. At the time of adoption, Knative was the most promising option available. 

IBM Cloud encountered no major obstacles with Knative itself; most challenges were related to Istio integration and scaling. IBM uses custom patches for scale-related issues. Container startup latency, especially for larger workloads, was another issue that was faced while providing support for function based workloads. 

IBM Cloud views the Knative community as very open and feels that the project generally has good governance. Maintainer response is generally good, but they have had some instances where issues took longer to resolve than anticipated.

Generally, they view Knative as very mature and support graduation of the project.

The entire adopter interview can be found here: [IBM Cloud Adopter Interview](knative-adopter-interview-adopter-ibm.md)
