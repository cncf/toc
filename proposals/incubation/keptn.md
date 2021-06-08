# Keptn Incubating Proposal

## Background

Keptn is a cloud-native application life-cycle orchestration tool. Its goal is not to replace tools that are already present in an organization, but to connect and orchestrate them to support the life-cycle management of applications. Keptn enables DevOps engineers as well as SREs to scale delivery and operations automation in their organizations. 

Keptn builds upon declarative definitions for multi-stage environments, SLO-based quality gates, and auto-remediation and integrates with tools via an event-based approach. 



## Sandbox proposal artifacts

- Sandbox proposal [PR #237](https://github.com/cncf/toc/pull/237)
- Sandbox proposal [Doc](https://github.com/cncf/toc/blob/main/proposals/keptn.adoc)
- Sandbox proposal [presentation](https://docs.google.com/presentation/d/1k1QcjeTmb1ZeYtwMntLw1IVTbJXGerGnGExow9oa7i4/edit#slide=id.g60b676c4ff_2_0)
- Sandbox proposal [review and recommendation document](https://docs.google.com/document/d/12Uhkv5aMnvM_HFvpYDuDTN2Zzc83qKOLYmp-k_RGWB8/edit#heading=h.kd4eg2uz3lt0)

### Sandbox Acceptance

Keptn was accepted into the CNCF Sandbox on [June 24th, 2020](https://lists.cncf.io/g/cncf-toc/message/4853).

### Progress since joining the CNCF Sandbox

Since joining the CNCF Sandbox, Keptn has made substantial progress in various dimensions: user adoption, feature set, ecosystem growth, and community growth.

**User Adoption**: Keptn has been adopted by 25+ organizations (highlights being listed in section [requirements](#incubation-stage-requirements)) and seen an 6x increase of weekly downloads of the Keptn CLI. The successful usage of Keptn has been presented in Keptn User Groups, blogs, as well as presentations of Keptn users in meetups.

**Feature Set**: Significant improvements have been made to Keptn since the Sandbox stage. Some of them are highlights here:

- *Delivery assistant*: promotion after an SLO-based quality gate evaluation can be done manually, automated or even via 3rd party integrations using e.g., a ChatOps approach with Slack.
- *Support for closed-loop remediation*: Remediation sequences orchestrated by Keptn include evaluation of the executed remediation actions to verify if a remediation sequence can be aborted as a result of a fixed issue or might need escalation if automated remediation can not solve the issue.
- *Clear separation of control-plane and execution-plane*: Both components can be installed separately and allow for different execution-planes that are connected to Keptn. 
- *Multi-cluster support*: Keptn control-plane and execution plane can run on separate clusters in the same or even in different cloud providers or on-prem. This allows, e.g., Keptn to orchestrate delivery of application even if each stage/environment is managed on a separate Kubernetes cluster.
- *Support for individual tasks in sequences*: User defined tasks can be added to a task sequence allowing for customized sequences.
- *Support for multiple parallel stages to be executed*: TBD


**Ecosystem Growth**: Keptn has grown its ecosystem by adding support for more than 15 tools and addd and strengthened integrations with CNCF projects as well as other tools. Besides, the Keptn team is providing [templates](https://github.com/keptn-sandbox?q=template&type=&language=&sort=) to foster new tools integrations. 
- Added integrations: ArgoRollouts, LitmusChaos, Jenkins, Slack, Splunk, Gitlab, Grafana, Jira, AnsibleTower, XMatters, Locust, Artillery, ZenDesk, Azure DevOps, OneChart (from Gimlet), and counting
- Strengthened integrations in CNCF Ecosystem: CloudEvents, Prometheus, Helm

**Community**: We have significantly grown our community, on average by the factor 2 spanning across our multiple channels. In Slack, which is our preferred way to interact with the Keptn community, we have more than double our weekly active users (from around 50 to 110+), and all other community channels have shown significant increase as well as shown in the following. Besides the community growth, we have established weekly Keptn developer and community meetings (see public CNCF calendar) as well as monthly Keptn user groups to foster exchange between Keptn adopters that share best practices on their Keptn usage. 


### Project and Community Statistics

| Statistic | Sandbox | Current | Multiplier |
|-|-|-|-|
| [GitHub stars](https://github.com/keptn/keptn/stargazers) | 410 | 893+ | >2x |
| Commits? |  | 5,500+ |
| [Releases](https://github.com/keptn/keptn/releases) | 18 | 30 | >1.5x |
| [Contributors (commits)](https://github.com/keptn/keptn/graphs/contributors) | 23 | 56+ | >2x |
| [Twitter followers](https://twitter.com/keptnProject) | 495 | 1,133+ | >2x |
| [Slack members](https://slack.keptn.sh) | WHERE TO FIND THIS NUMBER | 645+ (93,000+ messages sent) | TBA |
| [YouTube subscribers](https://www.youtube.com/c/keptn) | 78 | 235+ | >3x | 


## Collaboration within the CNCF Ecosystem

  * [Crossplane](https://www.youtube.com/watch?v=V8L-JTpkEpc&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=3) 
  * [LitmusChaos](https://www.youtube.com/watch?v=_DgCc4-BLW8) presented at KubeCon 2021 europe
  * [Tekton](https://cdeliveryfdn.slack.com/archives/C0151BTKEJX/p1622202061013700) 
  * [ArgoRollouts](https://www.youtube.com/watch?v=w-E8FzTbN3g)
  * [Helm]
  * [Cloudevents](https://cloudevents.io/) 

## Incubation Stage Requirements

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope

1. [Schlumberger](https://slb.com/), an oilfield services company working in more than 120 countries, has currently 4 applications evaluated with Keptn quality gates. Evaluations make use of 10-20 SLIs, one of them even 90 SLIs per quality gate evaluation. Their integration triggers Keptn quality gates via Azure DevOps automation on a daily basis.
    - Reference: Keptn user group TBA June 15

1. [Vitality Group](https://www.vitalitygroup.com/) is a leading behavior change platform to make people healthier and enhance and protect their lives. They are triggering Keptn quality gates by utilizing an integration with Jenkins. Currently, Vitality has 22 services onboarded in DEV, TEST & QA environments. They are running quality evaluations with Keptn multiple times a day in DEV, once per day in TEST, and Adhoc via the Keptn API in QA.
    - Reference: TBA

1. [Kitopi](https://www.kitopi.com/), the world's leading cloud kitchen platform, is using Keptn to evaluate the resilience of their applications. Using Keptn quality gates and Prometheus data, their evaluations are running in several nightly runs over 3 chaos stages, summing up in total of over 700 runs as of March 2021. Find a reference of Adrian Gonciarz, Lead QA Engineer, [who presented their use case](https://www.youtube.com/watch?t=222&v=m_RVxVQQrHo).

1. [Intuit](https://www.intuit.com/), a leading provider of financial software in the United States of America, is using Keptin conjunction with other tools in a system called DISTRO to allow for distributed load testing and evaluations thereof. The tooling includes Galting, Argo, and Keptn. Sumit Nagal, Principal Engineer, has summarized their application a [blog](https://sumitnagal.medium.com/distributed-load-testing-using-argo-in-kubernetes-distro-132c350f8733) and [presented it in a Keptn user group](https://www.youtube.com/watch?v=Omprl5OFtEw&t=402s)

1. Raiffeisen 
    - Use Case TBA
    - Reference TBA


There are many more adopters of Keptn and organizations evaluating the project as well. Organizations that agreed to have them listed publicly can be found in the [ADOPTERS.md](https://github.com/keptn/keptn/blob/master/ADOPTERS.md) file which currently lists 10 adopters, while some of these have only consented to sharing their details privately with the TOC, if required.

### Have a healthy number of committers. 

Keptn currently has contributions from [more than 15 different organizations](https://keptn.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=commits) and a [total of 50+ individual contributors](https://github.com/keptn/keptn/graphs/contributors) to the core project. 

There are 10 maintainers and a clear [membership definition](https://github.com/keptn/community/blob/master/COMMUNITY_MEMBERSHIP.md) on how to become a member, approver, and maintainer of the project. An example of a recent member that have been promoted via this process can be found [here](https://github.com/keptn/community/issues/45) and [here](https://github.com/keptn/community/issues/42). 

Due to the architecture of the project, it is worth to mention that 


### Demonstrate a substantial ongoing flow of commits and merged contributions

The project averages at ~620 contributions from around ~16 contributors per month - DEVSTATS https://keptn.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all 





### A clear versioning scheme


### Incubation Proposal Resources

TBA



