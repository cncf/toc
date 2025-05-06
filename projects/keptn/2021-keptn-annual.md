# Keptn Sandbox Annual Review

## Table of Contents

- [Background](#background)
- [Alignment with Cloud Native Computing Foundation](#alignment-with-cloud-native-computing-foundation)
- [Keptn's Year in Review](#keptns-year-in-review)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

[Keptn](https://keptn.sh/) is a cloud-native application life-cycle orchestration tool. Its goal is not to replace tools that are already present in an organization but to connect and orchestrate them to support the life-cycle management of applications. Keptn enables DevOps engineers as well as Site Reliability Engineers (SRE) to scale delivery and operations automation in their organizations. 

Keptn builds upon declarative definitions for multi-stage environments, SLO-based quality gates, and auto-remediation and integrates with tools via an event-based approach. 

**Keptn CloudEvents:** The event-based approach is built upon a well-defined defined set of *Keptn events*; currently in [v0.2.2](https://github.com/keptn/spec/blob/0.2.2/cloudevents.md). All *Keptn events* conform to the CloudEvents specification in version [v1.0](https://github.com/cloudevents/spec/blob/v1.0/spec.md). The CloudEvents specification is a vendor-neutral specification for defining the format of event data. In the course of the Keptn project, the event data is defined for the use-cases of application *delivery* and *remediation* as well as life-cycle orchestration tasks such as *deployment*, *test*, *evaluation*, *release*, *problem*, etc. The specification of Keptn CloudEvents is not limited to the mentioned tasks and can be easily extended by following the proposed [format](https://github.com/keptn/spec/blob/0.2.2/cloudevents.md#type). The Keptn project is currently in the progress of aligning the *Keptn events*  with the [event specification](https://github.com/cdfoundation/sig-events) from the Continuous Delivery Foundation (CDF) with the goal of establishing an industry-wide eventing standard for application life-cycle orchestration. 

**Keptn Control-Plane:** Keptn is built for Kubernetes and consists of a couple of Keptn core services that altogether form the Keptn control-plane. 
The control-plane is responsible to orchestrate the life-cycle of an application managed by Keptn.
Execution-plane service can connect to the control-plane to interact with Keptn via CloudEvents sent through NATS. The CloudEvents are currently stored in a MongoDB that serves as the datastore for all events that are sent via Keptn and allows for full traceability of life-cycle events. The architecture of the Keptn project can be found in the [Keptn documentation](https://keptn.sh/docs/concepts/architecture/).

**Keptn Quality Gates:** A central component of Keptn are [quality gate evaluations](https://keptn.sh/docs/concepts/quality_gates/) based on service-level objectives (SLOs). Therefore, Keptn builds upon SRE best practices such as service-level indicators (SLIs) and allows to declaratively define SLOs for them. These SLOs define quality criteria for the applications and act as a gatekeeper during software delivery before promoting any application or microservice from one environment (e.g,. hardening) to the next environment (e.g., production). 

**Keptn Life-Cycle Orchestration:** Keptn's architecture allows any tool to be integrated into the application life-cycle orchestration managed by Keptn. These *execution plane* services can run within the same cluster as the Keptn control plane or on different clusters, allowing to orchestrate multi-cluster deployments, tests, evaluations, and operational tasks such as remediation orchestration or ChatOps. 

### Accepted into CNCF Sandbox

**Keptn was accepted as a CNCF Sandbox project on June 24, 2020**

- Sandbox proposal [PR #237](https://github.com/cncf/toc/pull/237)
- Sandbox proposal [Doc](https://github.com/cncf/toc/blob/main/proposals/keptn.adoc)
- Sandbox proposal [presentation](https://docs.google.com/presentation/d/1k1QcjeTmb1ZeYtwMntLw1IVTbJXGerGnGExow9oa7i4/edit#slide=id.g60b676c4ff_2_0)
- Sandbox proposal [review and recommendation document](https://docs.google.com/document/d/12Uhkv5aMnvM_HFvpYDuDTN2Zzc83qKOLYmp-k_RGWB8/edit#heading=h.kd4eg2uz3lt0)


## Alignment with Cloud Native Computing Foundation

Keptn falls in the scope of [CNCF App Delivery TAG](https://github.com/cncf/tag-app-delivery).


**Keptn targets on:**
  - Providing an orchestration for different projects of the CNCF, e.g., for deployment, testing, observability, etc. 
  - Integrating with the CNCF ecosystem via CloudEvents and working on a common format to allow interoperability in CD systems.
  - Automating configuration and management of tooling by using declarative SRE best-practices such as SLOs to define quality gates, alerting rules, etc based on a single source of truth.

In the following, we provide a list of **ecosystem projects** that have collaborated or integrated with Keptn.

* [Prometheus](https://github.com/keptn-contrib/prometheus-service): configuring Prometheus by creating scrape jobs and alerting rules based on SLOs managed by Keptn and utilizing data from Prometheus for Keptn quality gates.
* [LitmusChaos](https://www.youtube.com/watch?v=_DgCc4-BLW8): orchestrating Chaos tests and load tests as part of a Keptn CD sequence as evaluating its impact with SLO-based quality-gates (presented at KubeCon 2021 Europe)
* [Tekton](https://www.youtube.com/watch?v=Qq3fTt6g2FM&list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC&index=6): bridging the gap between CI and CD by defining interoperable events that can be exchanged between Tekton and Keptn 
* [ArgoRollouts](https://www.youtube.com/watch?v=w-E8FzTbN3g): orchestrating a delivery sequence via ArgoRollouts for canary releases with Keptn quality gates to proceed or stop the rollout
* [Crossplane](https://www.youtube.com/watch?v=V8L-JTpkEpc&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=3): utilization of Crossplane for infrastructure management in DevOps workflows and Keptn CD sequences.
* [Helm](https://github.com/keptn/keptn/tree/master/helm-service): deployment of applications via Helm charts orchestrated by Keptn
* [Gitlab](https://www.youtube.com/watch?v=fyS8m8VoayM): automating deployment validation using Keptn’s SLI/SLO-based quality gates orchestrated in Gitlab CI/CD pipelines.
* [Istio](https://keptn.sh/docs/0.8.x/continuous_delivery/deployment_helm/): usage of Istio for traffic shifting between blue/green deployments. Keptn rewrites Istio virtual services and therefore manages the traffic shifting.
* [CloudEvents](https://github.com/keptn/spec/blob/master/cloudevents.md): all events that are sent to and from the Keptn control-plane make use of the CloudEvents specification. 
* [NATS](https://nats.io/): Keptn is using NATS as its message system internally and [contributed back to the NATS project](https://github.com/nats-io/k8s/pull/222).
* [Grafana](https://github.com/keptn-sandbox/grafana-service): Integration to automatically create dashboards for services managed by Keptn. 
* [Ansible](https://github.com/keptn-sandbox/ansibletower-service): Integration to trigger Ansible Tower playbooks as part of Keptn's orchestrated remediation sequences. 
* [Jenkins](https://github.com/keptn-sandbox/keptn-jenkins-library): Jenkins shared library for integrating Keptn use cases with Jenkins pipelines.
* [Auto-remediation working group](https://docs.google.com/document/d/1_WlLP6oLcHe0yyC7kXH2hB3i9bOPvIArp83NohE78FU/edit#): Keptn maintainers initiated a working group to discuss the current state and requirements of auto-remediation and to define the future of automated operations where auto-remediation is a big part of.


## Keptn's Year in Review

**Key features added**

- *Delivery assistant*: Promotion after an SLO-based quality gate evaluation can be done manually, automated, or even via 3rd party integrations using e.g., a ChatOps approach with Slack.
- *Support for closed-loop remediation*: Remediation sequences orchestrated by Keptn include evaluation of the executed remediation actions to verify if a remediation sequence can be aborted as a result of a fixed issue or might need escalation if automated remediation can not solve the issue.
- *Clear separation of control-plane and execution-plane*: Both components can be installed separately and allow for different execution-planes that are orchestrated by a Keptn control-plane. The separation of the planes enables: 
    - *Multi-cluster support*: Keptn control-plane and execution-plane can run on separate clusters in the same or even in different cloud providers, or on-prem. This allows, e.g., Keptn to orchestrate the delivery of the application even if each stage/environment is managed on a separate Kubernetes cluster.
    - *Parallel sequence executions*: The execution-planes are independent from each other meaning that they can execute sequences (e.g., for delivery or remediation) in parallel. This enables the use-case of deploying an app into multiple environments simultaneously. 
- *Support for individual tasks in sequences*: User-defined tasks can be added to a task sequence allowing for customized delivery or remediation use-cases.
- *Hardening of Keptn*: Introduced RBAC for all Keptn core services to restrict permissions on the Kubernetes cluster, added security contexts for core services, and core services are not running root users anymore.
- *Concept for subscribing execution-plane services (aka. integrations) to the control-plane*: Developed an approach to manage the subscription of execution-plane services to the Keptn control-plane.
- *Reduced resource consumption*: From Keptn 0.6.2 (the latest stable release as the project achieved Sandbox) to Keptn 0.8.3 (latest stable release as of today), the resource footprint has been reduced by various improvements: (1) removing Istio from the default Keptn installation, (2) switching from Helm 2 to Helm 3 what removed Tiller, (3) defining explicit resource quota for the Keptn core services, and (4) extracting services that are needed for delivery/remediation use-cases but not for the core functionality of Keptn (see: *Clear separation of control-plane and execution-plane* mentioned above).

**Keptn community and adoption growth** in the last year

| Statistic | Sandbox | Current | Multiplier |
|-|-|-|-|
| [GitHub stars](https://github.com/keptn/keptn/stargazers) | 410 | 959+ | >2x |
| [Commits](https://github.com/keptn/keptn) | 3,751 | 5,500+ | ~1.5x |
| [Releases](https://github.com/keptn/keptn/releases) | 18 | 30+ | >1.5x |
| [Contributors (commits)](https://github.com/keptn/keptn/graphs/contributors) | 23 | 56+ | >2x |
| [Twitter followers](https://twitter.com/keptnProject) | 495 | 1,133+ | >2x |
| [Slack weekly active users](https://slack.keptn.sh) | 50 | 110+ (645+ users in total) | >2x |
| [YouTube subscribers](https://www.youtube.com/c/keptn) | 78 | 235+ | >3x | 


## Annual Review Contents

### Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

- The project averages at [~620 contributions from around ~16 contributors](https://keptn.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) per month according to keptn.devstats.cncf.io contained within [~126 merged PRs](https://keptn.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now-9d) on average per month for the last year.


### How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

Keptn has established a formal process to deal with ownership/maintainership of the project. The process definition can be found in our [COMMUNITY_MEMBERSHIP.md](https://github.com/keptn/community/blob/master/COMMUNITY_MEMBERSHIP.md) and has been applied to approval or promotion of membership for our members, as can be seen e.g., [here](https://github.com/keptn/community/issues/45) and [here](https://github.com/keptn/community/issues/42).

Currently, the project has 11 maintainers, currently all affiliated with Dynatrace.
- Alois Reitbauer (@AloisReitbauer)
- Andreas Grabner (@grabnerandi)
- Johannes Bräuer (@johannes-b)
- Florian Bacher (@bacherfl)
- Jürgen Etzlstorfer (@jetzlstorfer)
- Elisabeth Lang (@laneli)
- Daniel Khan (@danielkhan)
- Ermin Muratovic (@ermin-muratovic)
- Klaus Strießnig (Kirdock)
- Christian Kreuzberger (@christian-kreuzberger-dtx)
- Bernd Warmuth (@warber)

Due to the architecture of the project, each ecosystem integration has their own maintainers, they can be found in the [Keptn-Contrib](https://github.com/keptn-contrib) and [Keptn-Sandbox](https://github.com/keptn-sandbox) organizations. The ones listed here are only for the Keptn core project. 


### What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

Some adoption highlights from the last year are:

1. [Schlumberger](https://slb.com/), an oilfield services company working in more than 120 countries, has currently 4 applications evaluated with Keptn quality gates. Evaluations make use of 10-20 SLIs, one of them even 90 SLIs per quality gate evaluation. Their integration triggers Keptn quality gates via Azure DevOps automation daily. It has been [presented it in a Keptn user group](https://youtu.be/9YYjk1e9ZGE).

1. [Vitality Group](https://www.vitalitygroup.com/) is a leading behavior change platform to make people healthier and enhance and protect their lives. They are triggering Keptn quality gates by utilizing an integration with Jenkins. Currently, Vitality has 22 services managed by Keptn in DEV, TEST & QA environments. They are running quality evaluations with Keptn multiple times a day in DEV, once per day in TEST, and ad-hoc via the Keptn API in QA. Reference will be added once publicly available.

1. [Kitopi](https://www.kitopi.com/), the world's leading cloud kitchen platform, is using Keptn to evaluate the resilience of their applications. Using Keptn quality gates and Prometheus data, their evaluations are running in several nightly runs over 3 chaos stages, summing up in a total of over 700 runs as of March 2021. Find a reference of Adrian Gonciarz, Lead QA Engineer, [who presented their use case](https://www.youtube.com/watch?t=222&v=m_RVxVQQrHo).

1. [Intuit](https://www.intuit.com/), a leading provider of financial software in the United States of America, is using Keptn in conjunction with other tools in a system called DISTRO to allow for distributed load testing and evaluations thereof. The tooling includes Gatling, Argo, and Keptn. Sumit Nagal, Principal Engineer, has summarized their application a [blog](https://sumitnagal.medium.com/distributed-load-testing-using-argo-in-kubernetes-distro-132c350f8733) and [presented it in a Keptn user group](https://www.youtube.com/watch?v=Omprl5OFtEw&t=402s)

1. [Raiffeisen Software](https://www.r-software.at/) is a leading software manufacturer for financial applications in Europe. They are running Keptn quality gates for internet banking, both for their mobile app and desktop applications evaluating 30+ SLIs and SLOs for each release. The software quality is evaluated weekly via load tests, and integrated into Jenkins. [Find their story here](https://medium.com/keptn/keptn-automates-release-readiness-validation-for-austrian-online-banking-software-eaaab7ad7856?source=friends_link&sk=a271e10817106d5ecefaeb3828cf35e8).


There are many more adopters of Keptn and organizations evaluating the project as well. Organizations that agreed to have them listed publicly can be found in the [ADOPTERS.md](https://github.com/keptn/keptn/blob/master/ADOPTERS.md) file which currently lists 10 adopters, while some of these have only consented to share their details privately with the TOC, if required.


### How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

- The Keptn project has established a 4 weeks release cadence since version Keptn 0.8.0. Before that releases were based on value increments. 
For our releases, Keptn maintainers follow the [release checklist](https://github.com/keptn/keptn/wiki/Release-Checklist) as documented for the latest releases [0.8.1](https://github.com/keptn/keptn/issues/3483), [0.8.2](https://github.com/keptn/keptn/issues/3746 ), and [0.8.3](https://github.com/keptn/keptn/issues/4089) were a dedicated [release manager](https://github.com/keptn/keptn/wiki/Release-Checklist#github-issues-milestone-etc) is responsible for driving the release process. If need be, there is a [hotfix checklist](https://github.com/keptn/keptn/wiki/Hotfix-Checklist) that has been executed for a [hotfix release](https://github.com/keptn/keptn/issues/3909) in the past.

- The project has also streamlined process for [GOVERNANCE](https://github.com/keptn/keptn/blob/master/GOVERNANCE.md), [MEMBERSHIP](https://github.com/keptn/community/blob/master/COMMUNITY_MEMBERSHIP.md), a clear [CONTRIBUTION](https://github.com/keptn/keptn/blob/master/CONTRIBUTING.md) guide, and established monthly [Keptn user group meetings](https://keptn.sh/community/meetings/#keptn-user-group) and continued weekly [Keptn community and developer meetings](https://keptn.sh/community/meetings/#keptn-developer-meeting).

- From a feature set perspective, the project successfully implemented [multi-cluster support](https://github.com/keptn/keptn/releases/tag/0.8.0), [self-healing procedures](https://github.com/keptn/keptn/releases/tag/0.7.0) via feature flag, rollbacks, and custom action providers, and is [working actively in the CDF](https://github.com/cdfoundation/sig-events/tree/main/poc) to support interoperability of CD tools by defining a common set of CD events. 




### What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

- The complete project [roadmap is publicly available](https://github.com/keptn/enhancement-proposals/blob/master/roadmap.md) and upcoming building blocks are listed below:

  - **[Zero-downtime Upgrades & High Availability](https://github.com/keptn/enhancement-proposals/pull/48)**: *Running Keptn at scale without downtime for end-users.*
      * Critical components in Keptn are running with multiple replicas, utilizing a rolling upgrade mechanism, and graceful shutdowns for various upgrading scenarios.
  - **[Keptn Uniform support](https://github.com/keptn/enhancement-proposals/issues/42)**: *Seamless integration of DevOps tooling into the sequence orchestration by Keptn.* 
      * UI-support for managing connected integrations and means to configure their subscriptions.
  - **Security: Access Control**: *Allow fine-grained access control for interacting with Keptn.* 
      * Basic functionality to control user access and user/API permissions on different entities (e.g., Keptn projects)
  - **Alignment with [Continous Delivery Foundation (CDF) eventing standard](https://github.com/cdfoundation/sig-events)**: *Establish an industry-wide eventing standard jointly with the CDF.* 
      * A translation layer will translate *CDF events* into *Keptn events*, followed by a step-wise integration of the established eventing standard into Keptn. [Recording of SIG meeting](https://www.youtube.com/watch?v=Qq3fTt6g2FM&list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC&index=6)
  - **Execution-Plane support**: *Central component (e.g., Operator) for installing and managing execution-planes.* 
      * For handling multiple execution-planes that are connected to one Keptn control-plane, a central component (*in the sense of an Operator*) has to be in place that allows installing an execution-plane, handles the communication back to the control-plane, and is responsible for operating the integrations (which are running on the execution-plane).
  - **Multi-tenancy**: *Reduction of resource footprint by multi-tenant capabilities.* 
      * This concludes architectural changes in Keptn core for maintaining multiple clients with multiple projects by one Keptn deployment.


- Integrations that are currently planned for Keptn, in no particular order:

  * [ArtifactHub](https://github.com/keptn/keptn/issues/3406) to provide a central hub for Keptn itself and Keptn integrations. Currently, integrations are curated via GitHub, and the ArtifactHub can provide a holistic view for the user on all available integrations.
  * [OpenTelemetry](https://github.com/keptn/enhancement-proposals/pull/30) to align the Keptn's internal "Keptn context" with the trace-context supported by OpenTelemetry.
  * [ArgoCD](https://github.com/keptn-sandbox/argocd-service/issues/1) for deployments as an alternative via Helm charts and ArgoRollouts that are already supported by Keptn.
  * [Flux](https://github.com/keptn/keptn/issues/2669), similar to ArgoCD, for deployments as an alternative via Helm charts and ArgoRollouts that are already supported by Keptn
  * [Falco](https://github.com/keptn-sandbox/falco-service/issues/1) to identify any security issues and have its ruleset managed by Keptn and use Keptn to orchestrate counter-action in response to identified security threats.
  * [Vault](https://github.com/keptn/keptn/issues/4342) to support encrypted secret management in Keptn.
  * [Snyk](https://github.com/keptn/keptn/issues/417) to be integrated into the CD process to check for security vulnerabilities of container images as part of a quality gate and as a dedicated security scanning task during multi-stage delivery.

- We want to further drive adoption of Keptn by engaging actively with the CNCF community.

### How can the CNCF help you achieve your upcoming goals?

- The CNCF can best help us in providing visibility for our project to drive adoption. This could be done by featuring Keptn in webinars, blogs, at community and industry events or via other publicly available resources. 

- The CNCF ecosystem is another driver for Keptn. We are already engaging with multiple CNCF projects and will further drive collaboration with the help of the CNCF.


### Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

- Yes, we've also filed our [incubation proposal](https://github.com/cncf/toc/pull/670).

## Project Links
 - [Website](https://keptn.sh)
 - [Github](https://github.com/kept/keptn)
 - [Slack](https://slack.keptn.sh)
 - [Twitter](https://twitter.com/keptnProject)
 - [YouTube](https://youtube.com/c/keptn)