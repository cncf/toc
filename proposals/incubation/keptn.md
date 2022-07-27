# Keptn Incubating Proposal

## Background

Keptn is a cloud-native application life-cycle orchestration tool. Its goal is not to replace tools that are already present in an organization but to connect and orchestrate them to support the life-cycle management of applications. Keptn enables DevOps engineers as well as SREs to scale delivery and operations automation in their organizations. 

Keptn builds upon declarative definitions for multi-stage environments, SLO-based quality gates, and auto-remediation and integrates with tools via an event-based approach. 

TBA Jürgen


## Sandbox proposal artifacts

- Sandbox proposal [PR #237](https://github.com/cncf/toc/pull/237)
- Sandbox proposal [Doc](https://github.com/cncf/toc/blob/main/proposals/keptn.adoc)
- Sandbox proposal [presentation](https://docs.google.com/presentation/d/1k1QcjeTmb1ZeYtwMntLw1IVTbJXGerGnGExow9oa7i4/edit#slide=id.g60b676c4ff_2_0)
- Sandbox proposal [review and recommendation document](https://docs.google.com/document/d/12Uhkv5aMnvM_HFvpYDuDTN2Zzc83qKOLYmp-k_RGWB8/edit#heading=h.kd4eg2uz3lt0)

### Sandbox Acceptance

Keptn was accepted into the CNCF Sandbox on [June 24th, 2020](https://lists.cncf.io/g/cncf-toc/message/4853).

### Progress since joining the CNCF Sandbox

Since joining the CNCF Sandbox, Keptn has made substantial progress in various dimensions, including user adoption, feature set, ecosystem growth, and community growth.

**User Adoption**: Keptn has been adopted by 25+ organizations (highlights being listed in section [requirements](#incubation-stage-requirements)) and seen a 6x increase of weekly downloads of the Keptn CLI. The successful usage of Keptn has been presented in Keptn User Groups, blogs, as well as presentations of Keptn users in meetups.

**Feature Set**: Significant improvements have been made to Keptn since the Sandbox stage. Some of them are highlights here:

- *Delivery assistant*: Promotion after an SLO-based quality gate evaluation can be done manually, automated, or even via 3rd party integrations using e.g., a ChatOps approach with Slack.
- *Support for closed-loop remediation*: Remediation sequences orchestrated by Keptn include evaluation of the executed remediation actions to verify if a remediation sequence can be aborted as a result of a fixed issue or might need escalation if automated remediation can not solve the issue.
- *Clear separation of control-plane and execution-plane*: Both components can be installed separately and allow for different execution-planes that are orchestrated by a Keptn control-plane. The separation of the planes enables: 
    - *Multi-cluster support*: Keptn control-plane and execution-plane can run on separate clusters in the same or even in different cloud providers, or on-prem. This allows, e.g., Keptn to orchestrate the delivery of the application even if each stage/environment is managed on a separate Kubernetes cluster.
    - *Parallel sequence executions*: The execution-planes are independent from each other meaning that they can execute sequences (e.g., for delivery or remediation) in parallel. This enables the use-case of deploying an app into multiple environments simultaneously. 
- *Support for individual tasks in sequences*: User-defined tasks can be added to a task sequence allowing for customized delivery or remediation use-cases.
- *Hardening of Keptn*: Introduced RBAC for all Keptn core services to restrict permissions on the Kubernetes cluster, added security contexts for core services, and core services are not running root users anymore.
- *Concept for subscribing execution-plane services (aka. integrations) to the control-plane*: Developed an approach to manage the subscription of execution-plane services to the Keptn control-plane.
- *Reduced resource consumption*: From Keptn 0.6.2 (the latest stable release as the project achieved Sandbox) to Keptn 0.8.3 (latest stable release as of today), the resource footprint has been reduced by various improvements: (1) removing Istio from the default Keptn installation, (2) switching from Helm 2 to Helm 3 what removed Tiller, (3) defining explicit resource quota for the Keptn core services, and (4) extracting services that are needed for delivery/remediation use-cases but not for the core functionality of Keptn (see: *Clear separation of control-plane and execution-plane* mentioned above).

**Ecosystem Growth**: Keptn has grown its ecosystem by adding support for more than 15 tools and added and strengthened integrations with CNCF projects as well as other tools. Besides, the Keptn team is providing [templates](https://github.com/keptn-sandbox?q=template&type=&language=&sort=) to foster new tools integrations. 
- Added integrations: ArgoRollouts, LitmusChaos, Slack, Splunk, Gitlab, Grafana, Jira, AnsibleTower, XMatters, Locust, Artillery, ZenDesk, Azure DevOps, OneChart (from Gimlet), and counting
- Strengthened integrations with CNCF Ecosystem: CloudEvents, Prometheus, Helm, Jenkins

TBD double check this list

**Community**: We have significantly grown our community, on average by the factor 2 spanning across our multiple channels. In Slack, which is our preferred way to interact with the Keptn community, we have more than double our weekly active users (from around 50 to 110+), and all other community channels have shown a significant increase as well as shown in the following. Besides the community growth, we have established weekly Keptn developer and community meetings (see public CNCF calendar) as well as monthly Keptn user groups to foster exchange between Keptn adopters that share best practices on their Keptn usage. 


### Project and Community Statistics

| Statistic | Sandbox | Current | Multiplier |
|-|-|-|-|
| [GitHub stars](https://github.com/keptn/keptn/stargazers) | 410 | 893+ | >2x |
| [Commits](https://github.com/keptn/keptn) | 3,751 | 5,500+ | ~1.5x |
| [Releases](https://github.com/keptn/keptn/releases) | 18 | 30+ | >1.5x |
| [Contributors (commits)](https://github.com/keptn/keptn/graphs/contributors) | 23 | 56+ | >2x |
| [Twitter followers](https://twitter.com/keptnProject) | 495 | 1,133+ | >2x |
| [Slack weekly active users](https://slack.keptn.sh) | 50 | 110+ (645+ users in total) | >2x |
| [YouTube subscribers](https://www.youtube.com/c/keptn) | 78 | 235+ | >3x | 


### Collaboration within the CNCF Ecosystem

In the following, we provide a list of ecosystem projects that have collaborated or integrated with Keptn.

* [LitmusChaos](https://www.youtube.com/watch?v=_DgCc4-BLW8): orchestrating Chaos tests and load tests as part of a Keptn CD sequence as evaluating its impact with SLO-based quality-gates (presented at KubeCon 2021 Europe)
* [Tekton](https://cdeliveryfdn.slack.com/archives/C0151BTKEJX/p1622202061013700): bridging the gap between CI and CD by defining interoperable events that can be exchanged between Tekton and Keptn 
* [ArgoRollouts](https://www.youtube.com/watch?v=w-E8FzTbN3g): orchestrating a delivery sequence via ArgoRollouts for canary releases with Keptn quality gates to proceed or stop the rollout
* [Crossplane](https://www.youtube.com/watch?v=V8L-JTpkEpc&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=3): utilization of Crossplane for infrastructure management in DevOps workflows and Keptn CD sequences.
* [Helm](https://github.com/keptn/keptn/tree/master/helm-service): deployment of applications via Helm charts orchestrated by Keptn
* [Gitlab](https://www.youtube.com/watch?v=fyS8m8VoayM): automating deployment validation using Keptn’s SLI/SLO-based quality gates orchestrated in Gitlab CI/CD pipelines.
* [Istio](https://keptn.sh/docs/0.8.x/continuous_delivery/deployment_helm/): usage of Istio for traffic shifting between blue/green deployments. Keptn rewrites Istio virtual services and therefore manages the traffic shifting.
* [CloudEvents](https://github.com/keptn/spec/blob/master/cloudevents.md): all events that are sent to and from the Keptn control-plane make use of the CloudEvents specification. 
* [NATS](https://nats.io/): Keptn is using NATS as its message system internally and [contributed back to the NATS project](https://github.com/nats-io/k8s/pull/222).
* [Ansible](https://github.com/keptn-sandbox/ansibletower-service): Integration to trigger Ansible Tower playbooks as part of Keptn's orchestrated remediation sequences. 
* [Jenkins](https://github.com/keptn-sandbox/keptn-jenkins-library): Jenkins shared library for integrating Keptn use cases with Jenkins pipelines.
* [Auto-remediation working group](https://docs.google.com/document/d/1_WlLP6oLcHe0yyC7kXH2hB3i9bOPvIArp83NohE78FU/edit#): Keptn maintainers initiated a working group to discuss the current state and requirements of auto-remediation and to define the future of automated operations where auto-remediation is a big part of.

Integrations that are currently planned for Keptn, in no particular order:

* [ArtifactHub](https://github.com/keptn/keptn/issues/3406) to provide a central hub for Keptn itself and Keptn integrations. Currently, integrations are curated via GitHub, and the ArtifactHub can provide a holistic view for the user on all available integrations.
* [OpenTelemetry](https://github.com/keptn/enhancement-proposals/pull/30) to align the Keptn's internal "Keptn context" with the trace-context supported by OpenTelemetry.
* ArgoCD for deployments as an alternative via Helm charts and ArgoRollouts that are already supported by Keptn.
* Flux, similar to ArgoCD, for deployments as an alternative via Helm charts and ArgoRollouts that are already supported by Keptn
* Falco to identify any security issues and have its ruleset managed by Keptn and use Keptn to orchestrate counter-action in response to identified security threats.
* Vault to support encrypted secret management in Keptn.
* Snyk to be integrated into the CD process to check for security vulnerabilities of container images as part of a quality gate and as a dedicated security scanning task during multi-stage delivery.


## Incubation Stage Requirements

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope

1. [Schlumberger](https://slb.com/), an oilfield services company working in more than 120 countries, has currently 4 applications evaluated with Keptn quality gates. Evaluations make use of 10-20 SLIs, one of them even 90 SLIs per quality gate evaluation. Their integration triggers Keptn quality gates via Azure DevOps automation daily.
    - Reference: Keptn user group TBA June 15

1. [Vitality Group](https://www.vitalitygroup.com/) is a leading behavior change platform to make people healthier and enhance and protect their lives. They are triggering Keptn quality gates by utilizing an integration with Jenkins. Currently, Vitality has 22 services onboarded in DEV, TEST & QA environments. They are running quality evaluations with Keptn multiple times a day in DEV, once per day in TEST, and Adhoc via the Keptn API in QA.
    - Reference: TBA

1. [Kitopi](https://www.kitopi.com/), the world's leading cloud kitchen platform, is using Keptn to evaluate the resilience of their applications. Using Keptn quality gates and Prometheus data, their evaluations are running in several nightly runs over 3 chaos stages, summing up in a total of over 700 runs as of March 2021. Find a reference of Adrian Gonciarz, Lead QA Engineer, [who presented their use case](https://www.youtube.com/watch?t=222&v=m_RVxVQQrHo).

1. [Intuit](https://www.intuit.com/), a leading provider of financial software in the United States of America, is using Keptn in conjunction with other tools in a system called DISTRO to allow for distributed load testing and evaluations thereof. The tooling includes Gatling, Argo, and Keptn. Sumit Nagal, Principal Engineer, has summarized their application a [blog](https://sumitnagal.medium.com/distributed-load-testing-using-argo-in-kubernetes-distro-132c350f8733) and [presented it in a Keptn user group](https://www.youtube.com/watch?v=Omprl5OFtEw&t=402s)

1. [Raiffeisen Software](https://www.r-software.at/) is a leading software manufacturer for financial applications in Europe. They are running Keptn quality gates for internet banking, both for their mobile app and desktop applications evaluating 30+ SLIs and SLOs for each release. The software quality is evaluated weekly via load tests, and integrated into Jenkins. 
    - Reference TBA (blog)


There are many more adopters of Keptn and organizations evaluating the project as well. Organizations that agreed to have them listed publicly can be found in the [ADOPTERS.md](https://github.com/keptn/keptn/blob/master/ADOPTERS.md) file which currently lists 10 adopters, while some of these have only consented to share their details privately with the TOC, if required.

### Have a healthy number of committers. 

Keptn currently has contributions from [more than 15 different organizations](https://keptn.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=commits) and a [total of 50+ individual contributors](https://github.com/keptn/keptn/graphs/contributors) to the core project. 

There are 10 maintainers and a clear [membership definition](https://github.com/keptn/community/blob/master/COMMUNITY_MEMBERSHIP.md) on how to become a member, approver, and maintainer of the project. An example of a recent member that have been promoted via this process can be found [here](https://github.com/keptn/community/issues/45) and [here](https://github.com/keptn/community/issues/42). 

TBA: steering committee (create an issue for this)

Due to the architecture of the project, it is worth mentioning that contributions and enhancement to Keptn do not necessarily be contributed to [Keptn core](https://github.com/keptn/keptn), but often are contributed to the Keptn ecosystem in terms of [Keptn-Contrib](https://github.com/keptn-contrib) and [Keptn-Sandbox](https://github.com/keptn-sandbox). Both ecosystem repositories currently hold 30+ integrations in total, maintained by 65+ contributors ([contrib](https://github.com/orgs/keptn-contrib/people) / [sandbox](https://github.com/orgs/keptn-sandbox/people)).


### Demonstrate a substantial ongoing flow of commits and merged contributions

The project averages at [~620 contributions from around ~16 contributors](https://keptn.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) per month according to keptn.devstats.cncf.io contained within [126 merged PRs](https://keptn.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now-9d) on average per month for the last year.

### A clear versioning scheme

The projects versioning scheme is related to the [semantic versioning](https://semver.org/spec/v2.0.0.html) structure and implemented for Keptn as follows:

- Major: currently 0 as the [spec](https://github.com/keptn/spec) is actively developed and not yet finalized
- Minor: is incremented if breaking change in the Keptn spec occurs
- Patch: can contain bug fixes and features (including [Keptn Enhancement Proposals](https://github.com/keptn/enhancement-proposals))

The Keptn project has a 4 weeks release cadence since version Keptn 0.8.0. Before that releases were based on value increments. 
For our releases, Keptn maintainers follow the [release checklist](https://github.com/keptn/keptn/wiki/Release-Checklist) as documented for the latest releases [0.8.1](https://github.com/keptn/keptn/issues/3483), [0.8.2](https://github.com/keptn/keptn/issues/3746 ), and [0.8.3](https://github.com/keptn/keptn/issues/4089) were a dedicated [release manager](https://github.com/keptn/keptn/wiki/Release-Checklist#github-issues-milestone-etc) is responsible for driving the release process. If need be, there is a [hotfix checklist](https://github.com/keptn/keptn/wiki/Hotfix-Checklist) that has been executed for a [hotfix release](https://github.com/keptn/keptn/issues/3909) in the past.

## Roadmap

The complete project [roadmap is publicly available](https://github.com/keptn/enhancement-proposals/blob/master/roadmap.md) and upcoming building blocks are listed below:

- **[Zero-downtime Upgrades & High Availability](https://github.com/keptn/enhancement-proposals/pull/48)**: *Utilization of Keptn at scale without downtime for the end-user.* 
    * Critical components in Keptn are running with multiple replicas, utilizing rolling upgrade mechanism, and graceful shutdowns for various upgrading scenarios.
- **[Keptn Uniform support](https://github.com/keptn/enhancement-proposals/issues/42)**: *Seamless integration of DevOps tooling into the sequence orchestration by Keptn.* 
    * UI-support for managing connected integrations and means to configure their subscriptions.
- **Security: Access Control**: *Allow fine-grained access control for interacting with Keptn.* 
    * Basic functionality to control user access and user/API permissions on different entities (e.g., Keptn projects)
- **Execution-Plane support**: *Central component (e.g., Operator) for installing and managing execution-planes.* 
    * For handling multiple execution-planes that are connected to one Keptn control-plane, a central component (*in the sense of an Operator*) has to be in place that allows installing an execution-plane, handles the communication back to the control-plane, and is responsible for operating the integrations (which are running on the execution-plane).
- **Multi-tenancy**: *Reduction of resource footprint by multi-tenant capabilities.* 
    * This concludes architectural changes in Keptn core for maintaining multiple clients with multiple projects by one Keptn deployment.

## Incubation Proposal Resources

TBA
