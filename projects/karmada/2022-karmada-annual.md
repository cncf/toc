# Karmada 2022 Annual Review

## Table of Contents

* [Background](#background)
  * [Sandbox Acceptance](#sandbox-acceptance)
* [Continous Community Momentum](#continous-community-momentum)
* [Robust Developer Ecosystem](#robust-developer-ecosystem)
* [Extensive Adoption](#extensive-adoption)
* [Rich Integration](#rich-integration)
* [Project Goals](#project-goals)
* [How CNCF Can Help](#how-cncf-can-help)
* [Incubation](#incubation)

## Background

Karmada (Kubernetes Armada) is a Kubernetes management system that enables you to run your cloud-native applications across multiple Kubernetes clusters and clouds, with no changes to your applications. By speaking Kubernetes-native APIs and providing advanced scheduling capabilities, Karmada enables truly open, multi-cloud Kubernetes.

Karmada aims to provide turnkey automation for multi-cluster application management in multi-cloud and hybrid cloud scenarios, with key features such as centralized multi-cloud management, high availability, failure recovery, and traffic scheduling.

Karmada provides default policy sets to ensure high application availability from multiple dimensions in different scenarios. Region-independent cluster management allows more flexible service deployment.

Karmada presents its value in the following two aspects:

1. Provides a smooth path for organizations to evolve from a single-cluster architecture to a multi-cloud one, while channeling the cloud native technology ecosystem to their native Kubernetes capabilities in use.
2. Provides functional modules covering HA scheduling, failover, multi-cluster service discovery, multi-cloud cluster lifecycle management, etc, as well as policy sets for a variety of scenarios, which allow users to customize their own multi-cloud platforms.

### Sandbox Acceptance

Karmada was proposed and accepted as a sandbox project in September 2021 ([Sandbox project onboarding issue](https://github.com/cncf/toc/issues/721)).

## Continuous Community Momentum

In the past year, we joined hand with end users across industries and witnessed successful implementations of Karmada in many real-world projects.
We have built channels to reach out to our users to know better what they want and need, such as meetups and online activities.
Also, we continue to refine our documentation and use cases so that our users can get started with our project faster.

Before our journey of developing Karmada, we were deeply engaged in the KubeFed project led by the Kubernetes multi-cluster SIG.
Such engagement brought us insights and expertise to develop Karmada. Karmada is built based on Kubernetes Federation v1 and v2 and inherits some basic concepts from them.
We also marked our presence in upstream communities.

- In the Kubernetes community, we led the design of the [work API](https://github.com/kubernetes-sigs/work-api) for multi-cluster deployments. Our contributions make us the [maintainers of this project](https://github.com/kubernetes-sigs/work-api/blob/master/OWNERS).
  Meanwhile, we are actively involved in the [mcs-api](https://github.com/kubernetes-sigs/mcs-api) and [about-api](https://github.com/kubernetes-sigs/about-api) projects hosted by Kubernetes multi-cluster SIG.
  We are also one of the earliest projects to adopt the above APIs of the upstream communities. Our users can easily use them in Karmada. For example, we provide [multi-cluster service discovery](https://karmada.io/docs/next/userguide/service/multi-cluster-service) based on multi-cluster service APIs.
- Besides Kubernetes, we are open to integration of more CNCF projects and already have some amazing cases (details available [here](#rich-integration)).

The above efforts take the community to a new height since the project entered the CNCF sandbox.
Numbers can tell the story.

Karmada currently has [330+](https://karmada.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&viewPanel=1) contributors from [50+](https://karmada.devstats.cncf.io/d/5/companies-table?orgId=1) organizations and [14](https://karmada.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&viewPanel=1) countries and regions.
More details about our developers are available in [#Robust Developer Ecosystem](#robust-developer-ecosystem).

* We hold bi-weekly community meetings regularly (38 times as of September 2022), with about 25 attendees each time. [Meeting minutes](https://docs.google.com/document/d/1y6YLVC-v7cmVAdbjedoyR5WL0-q45DBRXTvz5_I7bkA/edit).
* [GitHub stars](https://karmada.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1) increased from **1,600+ to 2,600+**
* [GitHub forks](https://karmada.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1) increased from **200+ to 500+**
* [Number of contributions](https://karmada.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now&editPanel=5) increased to **1,500+ per month**
* [Number of new PRs](https://karmada.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&editPanel=1) increased to **30+ per week**
* Number of commits increased from **300+ to 1,500+**
* Number of PR reviewers increased to **100+**

More devStats information can be found at: https://karmada.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m.

## Robust Developer Ecosystem

We attach great importance to developers.
We have established [the community membership](https://github.com/karmada-io/community/blob/main/community-membership.md) and wrote a solid contributor guide. Based on this open and transparent mechanism, developers can continue to grow while equally participating in the project.
We have established a mature mentor mechanism in the community. Some active and qualified contributors, mentored attentively, have been promoted to members or maintainers.
Also, we are actively involved in various open source events, such as [LFX mentorship](https://github.com/karmada-io/karmada/issues/2371), helping college students better grow in the open source community.

In the past year, our project expanded to 330+ contributors from 50+ different organizations, including end users,cloud vendors, academic institutions and individual developers, showing the good diversity of the community.

![Contributors](https://github.com/karmada-io/community/raw/28aae23875fbf4a2ae5d0df4b478190118c16262/cncf_process/2022_karmada_annual/images/contributors.png)

Among the contributors, there are **23** people recognized by the community as reviewers and approvers.
Most of contributors are newcomers to our project last year.
We are very excited to see them grow in the community and look forward to their greater contributions.


| Reviewer/Approver | GitHub ID           | Affiliation             |
|-------------------|---------------------|-------------------------|
| Chauncey          | @chaunceyjiang      | DaoCloud                |
| Fan Bao           | @carlory            | DaoCloud                |
| Future            | @pidb               | AIML INSTITUTE          |
| Hanbo Li          | @mrlihanbo          | Huawei                  |
| Hongcai Ren       | @RainbowMango       | Huawei                  |
| Ihor Sychevskyi   | @Arhell             | Phalcon                 |
| Jinhui Ying       | @ikaven1024         | DiDi                    |
| Jun Qian          | @qianjun1993        | Tencent                 |
| Kevin Wang        | @kevin-wangzefeng   | Huawei                  |
| Prodan            | @prodanlabs         | Independent             |
| Rui Fang          | @Garrybest          | Tencent                 |
| Rui Jiang         | @jrkeen             | BMW-CIH                 |
| Rupesh Gelal      | @rgrupesh           | Independent             |
| Tao Ling          | @Tingtal            | Huawei                  |
| Tiecheng Shen     | @Poor12             | Huawei                  |
| Wang Bing         | @pigletfly          | Youzan                  |
| Xinzhao Xu        | @iawia002           | Kubesphere              |
| Yan Li            | @huntsman-li        | Hualala                 |
| Yasong Li         | @liys87x            | Hualala                 |
| Yiheng Ci         | @lfbear             | VIPKID                  |
| Yingjun Wu        | @wuyingjun-lucky    | China Mobile            |
| Zhe Cheng         | @lonelyCZ           | SEL-Zhejiang University |
| Zhen Chang        | @XiShanYongYe-Chang | Huawei                  |

Compared with maintainers initially, Karmada currently has **7** maintainers from **4** different companies:


| Maintainer  | GitHub ID         | Affiliation |
|-------------|-------------------|-------------|
| Hanbo Li    | @mrlihanbo        | Huawei      |
| Hongcai Ren | @RainbowMango     | Huawei      |
| Kevin Wang  | @kevin-wangzefeng | Huawei      |
| Lei Xue     | @carmark          | Tencent     |
| Shiyi Xie   | @GitHubxsy        | Huawei      |
| Yifan Shen  | @zoroyouxi        | ICBC        |
| Yiheng Ci   | @lfbear           | VIPKID      |

The full maintainer list can be found in our [GitHub repository](https://github.com/karmada-io/karmada/blob/master/MAINTAINERS.md).

## Extensive Adoption

We once released new features every month. Starting from version 1.0.0, we changed our pace to every 3 months, with the aim to deliver more stable features and smooth upgrade.
The biggest accomplishment of the community for the past year is the release of v1.0.0, the first milestone for the project, which means Karmada is now stable and ready to be used in production environments.
We are glad that this has been validated by project adopters successfully running Karmada in their own prod.
Future enhancements and improvements are intended to be done in a backwards compatible way with a clear upgrade path. We do not anticipate any further breaking changes in subsequent minor and patch releases.

With Karmada reaching maturity, more and more users adopt Karmada as the core to build their multi-cloud and multi-cluster container orchestration engine for higher service efficiency.
The community is witnessing a promising growth. Some success cases provided below set a good example for future adopters.

Karmada now has 11 adopters that using Karmada in `production` environments as of Oct 25, 2022.

* **VIPKID** is an online English teaching platform which delivers 150 million training sessions across countries and regions. To provide better services, VIPKID deploys applications by region closer to teachers and trainees.
Therefore, VIPKID purchased dozens of clusters from multiple cloud vendors around the world to build its global infrastructure.
How to manage multi-region and multi-vendor clusters has become a new challenge for VIPKID.
To solve the this challenge, VIPKID finally chooses Karmada. __Yiheng Ci, Senior Backend R&D Expert at VIPKID__, said:

  > "We surveyed and evaluated many solutions in the communities and decided on Karmada. **After Karmada is introduced, the multi-cluster aggregation layer is truly unified.**
  > We can manage resources by application on the Karmada control plane. We only need to interact with Karmada, not the clusters, which simplifies containerized application management and enables our PaaS platform to fully focus on service requirements. With Karmada integrated into GitOps, system components can be easily released and upgraded in each cluster, exponentially more efficient than before."
  >

  Refer to [VIPKID case study](https://karmada.io/docs/next/casestudies/vipkid) for more details.
* **ICBC** services run with specific requirements and on a large-scale cloud native infrastructure. There are four typical requirements, namely, high availability (HA) deployment, cross-cluster auto scaling, cross-cluster scheduling, and a specific dependency on Kubernetes versions.
ICBC is running over **100** Kubernetes clusters and more than **280,000** containers. They are managed by the container cloud platform in a unified manner. However, ICBC also faces the following four problems: 1) Limited availability. A Kubernetes cluster is also a fault domain, requiring automatic recovery across fault domains. 2) Limited resources. Application scheduling and auto scaling are available only to single clusters. 3) Non-transparent clusters. Upper-layer applications
are not aware of cluster differences. 4) Duplicate configuration. Although services are configured on the cloud management platform in a unified manner, their configurations need to be delivered to each cluster separately.
To solve these problems, ICBC finally chooses Karmada. __Yifan Shen, PaaS Platform Architect at ICBC__ said:

  > "For cross-cluster deployment models and scheduling, Karmada provides independent propagation policies, defined using a policy API, that can be reused across applications. For example, ICBC has its independent platform team and service team. This model can well adapt to their organizational structures."
  >

  The use case details can be found [here](https://karmada.io/zh/blog/2022/03/22/multi-cluster-management-with-an-ocean-of-nodes/).

* **Vivo** is a tech company that designs and develops smartphones, smartphone accessories, software and online services. To better manage their large-scale cloud native infrastructure, Vivo requires the following capabilities: 1) Unified management of multiple Kubernetes clusters and resource management by application; 2) Cross-cluster auto scaling and scheduling for
more efficient use of cross-cluster resources; 3) Dynamic migration of applications for quick failover; 4) Better DR
To obtain the above capabilities, Vivo chooses Karmada. According to __Rong Zhang from Vivo__:

  > "As we are migrating services to Kubernetes, we have more and more clusters. Both their scale and O&M difficulty increase sharply. To support multi-cluster deployments, we invested much in self development, but the returns were not satisfying, so we chose Karmada." 
  >

* **Trip.com** is a Chinese multinational online travel company that provides services including accommodation reservation, transportation ticketing, packaged tours and corporate travel management. Before using Karmada, Trip.com used KubeFed to build their multi-cloud container platform. However, KubeFed uses non-native Kubernetes APIs, which makes it difficult to migrate existing clusters to their multi-cloud platform. The KubeFed community is also becoming inactive.
So Trip.com seeks smooth migration from KubeFed, as well as cross-cluster auto scaling and scheduling for their multi-cloud platform. __HongHui Yue, Lead of Infrastructure at Trip.com__ said:

  > "We have used KubeFed for a long time to manage our clusters. However, we encountered the problems of high reconstruction costs and no auto scaling across clusters. We are glad that Karmada helps us solve these problems."
  > 

* **AIML INSTITUTE** is a tech company that helps enterprises build integrated cloud native solutions for digital transformation. Built on Kubernetes, their platform is cloud-vendor-independent. Customers can host service applications without regard for vendor differences.
Their customers are demanding more on multicloud, and the O&M complexity increases sharply with the growing scale and number of clusters. The company needs to find a way to satisfy their customers. With Karmada, existing resource definitions can be migrated into multi-cluster environments without modification.
In the words of AIML INSTITUTE:

  > "Karmada brings in the following benefits for AIML INSTITUTE in multi-cluster management: 1) Existing resource definitions can be migrated into multi-cluster environment without modification, owing to Karmada's compatibility with Kubernetes native APIs. 2) Based on the Karmada cluster APIs, Cluster Controllers, and pull/push modes,
  > a multi-cluster control standard is established for unified output of the multi-cluster management capabilities of any upper-layer system. 3) Scenario-specific multi-cluster resource scheduling and orchestration are realized based on Karmada's built-in controllers, scheduler plugins, and extended scheduler equivalent to that of Kubernetes. 4) The flexible architecture design of Karmada enables the existing single-cluster systems to be quickly switched to multi-cluster ones."
  >

  The use case details can be found [here](https://karmada.io/docs/next/casestudies/ci123).
* **Huawei Cloud**'s UCS(Ubiquitous Cloud Native Service) adopts Karmada to connect and manage cross-cloud and cross-region clusters and the cloud native applications running in them.
UCS provides consistent experience in cloud native deployment, management, and application ecosystem.
From the UCS team:

  > "Running on Karmada, UCS schedules multi-cloud resources in a unified manner and bursts on-premises applications to the public cloud. Supporting multiple types of distributed deployment policies, UCS can find the best-fit location to deploy your application based on global resource distribution, service characteristics, geographical locations, network QoS, and resource balancing."
  >

  The use case details can be found [here](https://support.huaweicloud.com/intl/en-us/productdesc-ucs/ucs_productdesc_0001.html).

The full adopter list can be found in our [website](https://karmada.io/docs/next/casestudies/adopters).

## Rich Integration

In the past year, Karmada actively cooperates with other CNCF projects and helps them better transition to multi-cluster based on Kubernetes native APIs.

The following lists ecosystem projects integrated with Karmada:

* [Argo CD](https://karmada.io/docs/userguide/cicd/working-with-argocd), a CNCF incubating project, makes for a very powerful combination with Karmada to manage workloads across clusters and simplify the delivery of multi-cloud applications, further extending the self-service infrastructure capabilities.
* [Flux](https://fluxcd.io/flux/use-cases/karmada/), the CNCF CI/CD project, partnered with Karmada to support Helm chart propagation and customization across clusters among multiple regions and multiple vendors.
* [Kyverno](https://karmada.io/docs/userguide/security-governance/working-with-kyverno), a CNCF incubating project, works alongside Karmada to unify management of security policies for multiple clusters. With Karmada, users do not have to tediously configure security policies in member clusters.
  [The discussions](https://github.com/kyverno/kyverno/discussions/4282) we initiated in the kyverno community also provided a new way of thinking for kyverno to manage multi-cluster security policies.
* [Velero](https://karmada.io/docs/administrator/backup/working-with-velero), which provides the ability to backup and migrate Kubernetes resource and persistent volumes, can utilize Karmada to implement the unified backup and recovery of resources across multiple clusters.
  We are glad that AIML INSTITUTE used this solution in production.
* [Istio](https://karmada.io/docs/userguide/service/working-with-istio-on-flat-network), the CNCF service mesh project, can utilize Karmada to implement multi-cluster traffic governance in both flat and non-flat network planes.
* In terms of multi-cluster observability, Karmada works with [Prometheus](https://karmada.io/docs/administrator/monitoring/working-with-prometheus) to build a multi-cloud unified monitoring view. And [Filebeat](https://karmada.io/docs/administrator/monitoring/working-with-filebeat), which provides the ability to collect logs, can utilize Karmada to
  implement unified log collection in multi-cloud scenarios.
* [Submariner](https://karmada.io/docs/next/userguide/network/working-with-submariner), which provides the ability to flatten the network between clusters, works alongside with Karmada to provide the capabilities of application management and service discovery in the upper layer of the multi-cluster network.
* Besides above, we also integrate with [OPA(gatekeeper)](https://karmada.io/docs/next/userguide/security-governance/working-with-gatekeeper), Thanos and so on.

## Project Goals

> How has the project performed against its goals since the last review?

Karmada aims to provide turnkey automation for multi-cluster application management in multi-cloud and hybrid cloud scenarios,
with key features such as centralized multi-cloud management, high availability, failure recovery, and traffic scheduling.

Following the above goals, we concentrated efforts on:

* Building an extensible, open architecture and a stable core API and feature set that is reliable and mature
* Helping users to deploy Karmada in their production environments

In the last year we released 1 major version and 6 minor versions. Version 1.0.0 marks a stage of steady iteration for the Karmada project.
Also, as CNCF's first multi-cloud container orchestration project, we solved many pain points of users using containers in multi-cloud scenarios and brought an excellent user experience to users.
We are glad to see more and more users are adopting Karmada and running it in production successfully.

We summarize our work in the last year as follows:

* **More powerful cross-cluster scheduling capabilities**. In the last year, we have done a lot of work on cross-cluster scheduling of applications. Facing a variety of user scenarios, we introduced richer scheduling strategies and a new component named `karmada-scheduler-estimator` to improve scheduling accuracy.
What's more, we introduced `karmada-descheduler` to reschedule in case of cluster resource imbalance. Based on this, users can distribute their applications more easily among clusters across regions and vendors.
* **Easier to manage the cluster and the resources on the cluster**. In the past year, we introduced a global uniform resource view across clusters.
  - We enhanced the framework of the resource interpreter to enable users to customize the resource status collected to the control plane.
  - We introduced the aggregated API to allow users to connect and operate member clusters from the control plane.
  - We introduced a new component `karmada-search` to implement global search and global proxy for multi-cluster resources.
    Based on this, users no longer have to deal with tedious and repetitive configuration management, nor to endure the management costs caused by fragmented API access portals.
* **More convenient cross-cluster service governance**. We support multi-cluster service discovery and a one-stop multi-cluster network solution. Now users can enable service governance for cross-cluster with Karmada.
* **High availability for multi-cloud applications**. We support cross-cluster failover of applications.
When a cluster fails, applications can smoothly migrate from the failed cluster to another available cluster. During the replicas migration process, Karmada ensures that the service replicas do not drop to zero, thereby ensuring that the service will not be interrupted.
Based on this, users can use Karmada to manage distributed applications with confidence.
* **Smoother multi-cluster O&M experience**. In the past year, We continue to optimize our command line tools. Now users can complete the creation of the Karmada control plane and bootstrap token-based cluster registration in just a few clicks.
Meanwhile, we provide not only a smooth version upgrade path but also a clear approach to migrate from Kubefed or promote legacy workloads from Kubernetes.

> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

The Karmada project continues to grow and has its sight on short-term and long-term goals:

* **Multi-cluster security**. Multi-cluster security presents many challenges.
  - Karmada manages clusters centrally, which eases cluster management but brings some security risks.
    The aim is to minimize the permissions of the control plane and reduce blast radius to member clusters after a control plane failure.
  - Another target is on multi-cluster access control with finer-grained security isolation on the control plane.
* **Larger scale landing**. We have started early on improving user experience in large-scale scenarios in the past releases.
Multicloud means larger resource pools, which brings greater performance pressure to the multi-cloud management platform.
In large-scale scenarios, we will continue to reduce the access latency, resources consumed, and stability impact on the system under high concurrency.
* **Multi-cluster observability**. Multi-cluster observability is an important part in a multi-cluster orchestration.
  - We will continue to explore integrated solutions for multi-cloud monitoring and logging.
  - We will continue to optimize multi-cluster events so that users can better maintain the multi-cloud management platform.
* **Multi-cluster application distribution**. There are still many challenges in application distribution across multiple clusters.
  - For plug-ins coupled with the cluster, we will dive into simplifying the plug-in installation and lifecycle management.
  - When applications have sequential relationships between clusters, we aim to provide a multi-cluster workflow to simplify the deployment of these applications.
* **Integration with ecosystem**.
  The integration with the CNCF ecosystem has always been a major focus of our efforts.
  We believe that cooperation in different fields can better serve our users.
  We will continue to seek cooperation with other mature projects to form an open and prosperous ecosystem.

The planed new features are listed in [roadmap](https://github.com/karmada-io/karmada/blob/master/ROADMAP.md).

## How CNCF Can Help

Similar to many other Sandbox projects, we wish to get the following help from CNCF:

- More channels to advocate the project.
- More chances to collaborate with other projects, both CNCF and ecosystem ones.
- Technical writing support for project documents.

## Incubation

We are preparing for the incubation proposal.
