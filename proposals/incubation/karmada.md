# Karmada proposal for CNCF Incubation

## Table of Contents

* [About Karmada](#about-karmada)
  * [Sandbox Acceptance](#sandbox-acceptance)
* [Progress Since Sandbox](#progress-since-sandbox)
  * [Continous Community Momentum](#continous-community-momentum)
  * [Robust Developer Ecosystem](#robust-developer-ecosystem)
  * [Technical Advances](#technical-advances)
* [Incubation Stage Requirements](#incubation-stage-requirements)
* [Statement on Alignment with the CNCF Mission](#statement-on-alignment-with-the-cncf-mission)
* [Future Plan](#future-plan)

## About Karmada

Karmada (Kubernetes Armada) is a Kubernetes management system that enables you to run your cloud-native applications across multiple Kubernetes clusters and clouds, with no changes to your applications. By speaking Kubernetes-native APIs and providing advanced scheduling capabilities, Karmada enables truly open, multi-cloud Kubernetes.

Karmada aims to provide turnkey automation for multi-cluster application management in multi-cloud and hybrid cloud scenarios, with key features such as centralized multi-cloud management, high availability, failure recovery, and traffic scheduling.

Karmada provides default policy sets to ensure high application availability from multiple dimensions in different scenarios. Region-independent cluster management allows more flexible service deployment.

Karmada presents its value in the following two aspects:

1. Provides a smooth path for organizations to evolve from a single-cluster architecture to a multi-cloud one, while channeling the cloud native technology ecosystem to their native Kubernetes capabilities in use.
2. Provides functional modules covering HA scheduling, failover, multi-cluster service discovery, multi-cloud cluster lifecycle management, etc, as well as policy sets for a variety of scenarios, which allow users to customize their own multi-cloud platforms.

### Sandbox Acceptance

Karmada was proposed and accepted as a sandbox project in September 2021 ([Sandbox project onboarding issue](https://github.com/cncf/toc/issues/721)).

## Progress Since Sandbox

### Continous Community Momentum

In the past year, we joined hand with end users across industries and witnessed successful implementations of Karmada in many real-world projects.
We have built channels to reach out to our users to know better what they want and need, such as meetups and online activities.
Also, we continue to refine our documentation and use cases so that our users can get started with our project faster.

Before our journey of developing Karmada, we were deeply engaged in the KubeFed project led by the Kubernetes multi-cluster SIG.
Such engagement brought us insights and expertise to develop Karmada. Karmada is built based on Kubernetes Federation v1 and v2 and inherits some basic concepts from them.
We also marked our presence in upstream communities.

- In the Kubernetes community, we led the design of the [work API](https://github.com/kubernetes-sigs/work-api) for multi-cluster deployments. Our contributions make us the [maintainers of this project](https://github.com/kubernetes-sigs/work-api/blob/master/OWNERS).
  Meanwhile, we are actively involved in the [mcs-api](https://github.com/kubernetes-sigs/mcs-api) and [about-api](https://github.com/kubernetes-sigs/about-api) projects hosted by Kubernetes multi-cluster SIG.
  We are also one of the earliest projects to adopt the above APIs of the upstream communities. Our users can easily use them in Karmada. For example, we provide [multi-cluster service discovery](https://karmada.io/docs/next/userguide/service/multi-cluster-service) based on multi-cluster service APIs.
- Besides Kubernetes, we are open to integration of more CNCF projects and already have some amazing cases (details available [here](https://github.com/karmada-io/community/blob/main/cncf_process/2022_karmada_annual/2022-karmada-annual.md#rich-integration)).

The above efforts take the community to a new height since the project entered the CNCF sandbox.
Numbers can tell the story.

Karmada currently has [330+](https://karmada.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&viewPanel=1) contributors from [50+](https://karmada.devstats.cncf.io/d/5/companies-table?orgId=1) organizations and [14](https://karmada.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&viewPanel=1) countries and regions.
More details about our developers are available in [#Robust Developer Ecosystem](#robust-developer-ecosystem).

* We hold bi-weekly community meetings regularly (38 times as of September 2022), with about 25 attendees each time. [Meeting minutes](https://docs.google.com/document/d/1y6YLVC-v7cmVAdbjedoyR5WL0-q45DBRXTvz5_I7bkA/edit).
* [GitHub stars](https://karmada.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1) increased from **1,600+ to 2,900+**
* [GitHub forks](https://karmada.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1) increased from **200+ to 500+**
* [Number of contributions](https://karmada.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now&editPanel=5) increased to **1,700+ per month**
* [Number of new PRs](https://karmada.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&editPanel=1) increased to **30+ per week**
* Number of commits increased from **300+ to 1,700+**
* Number of PR reviewers increased to **110+**

Specific devStats can be found at: https://karmada.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m.

### Robust Developer Ecosystem

We attach great importance to developers and our community has matured significantly.
We have established [the community membership](https://github.com/karmada-io/community/blob/main/community-membership.md) and wrote a solid contributor guide. Based on this open and transparent mechanism, developers can continue to grow while equally participating in the project.
We have established a mature mentor mechanism in the community. Some active and qualified contributors, mentored attentively, have been promoted to members or maintainers.
Also, we are actively involved in various open source events, such as [LFX mentorship](https://github.com/karmada-io/karmada/issues/2371), helping college students better grow in the open source community.

In the past year, our project expanded to **330+** contributors from **50+** different organizations, including end users,cloud vendors, academic institutions and individual developers, showing the good diversity of the community.

![Contributors](https://github.com/karmada-io/community/raw/28aae23875fbf4a2ae5d0df4b478190118c16262/cncf_process/2022_karmada_annual/images/contributors.png)

Among the contributors, there are **70** members and **23** approvers/reviewers.
Most of contributors are newcomers to our project last year.
We are very excited to see them grow in the community and look forward to their greater contributions.
The full approver and reviewer list can be found at https://github.com/karmada-io/community/blob/main/REVIEWERS.md.

We now have **7** maintainers from **4** different companies.
The full maintainer list can be found at https://github.com/karmada-io/karmada/blob/master/MAINTAINERS.md.

We codified our [contributing guidelines](https://github.com/karmada-io/karmada/blob/master/CONTRIBUTING.md) and [security processes](https://github.com/karmada-io/community/blob/main/security-team/SECURITY.md).

### Technical Advances

In the past year we released 1 major version and 6 minor versions. Version 1.0.0 marks a stage of steady iteration for the Karmada project.
Also, as CNCF's first multi-cloud container orchestration project, we solved many pain points of users using containers in multi-cloud scenarios and brought an excellent user experience to users.
We are glad to see more and more users are adopting Karmada and running it in production successfully.

We have significantly expanded the covered technical areas for the project development since Karmada joined the Sandbox.
Karmada has made feature developments in at least the following technical fields:

* **More powerful cross-cluster scheduling capabilities**. In the past year, we have done a lot on cross-cluster scheduling of applications. Facing a variety of user scenarios, we introduced richer scheduling strategies and a new component named `karmada-scheduler-estimator` to improve scheduling accuracy.
  What's more, we introduced `karmada-descheduler` for rescheduling in case of cluster resource imbalance. Based on this, users can distribute their applications more easily among clusters across regions and vendors.
* **Easier to manage the cluster and its resources**. We introduced a uniform resource view for global clusters.
  - We enhanced the framework of the resource interpreter to enable users to customize the resource status collected to the control plane.
  - We introduced the aggregated API to allow users to connect and operate member clusters from the control plane.
  - We introduced a new component `karmada-search` to implement global search and global proxy for multi-cluster resources.
  With these, users are freed from repetitive configurations and the management costs caused by fragmented API access portals.
* **More convenient cross-cluster service governance**. We support multi-cluster service discovery and a one-stop multi-cluster network solution. Now users can enable service governance for cross-cluster with Karmada.
* **High availability for multi-cloud applications**. We support cross-cluster failover of applications.
  When a cluster fails, applications can smoothly migrate to an available cluster. During replicas migration, Karmada ensures that the service replicas do not drop to zero, thereby ensuring that the service will not be interrupted.
  Based on this, users can use Karmada to manage distributed applications with confidence.
* **Smoother multi-cluster O&M experience**. In the past year, We continue to optimize our command line tools. Now users can complete the creation of the Karmada control plane and bootstrap token-based cluster registration in just a few clicks.
  Meanwhile, we provide not only a smooth version upgrade path but also a clear approach to migrate from Kubefed or promote legacy workloads from Kubernetes.

## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent direct adopters which, in the TOC’s judgement, are of adequate quality and scope._**

11 adopters are using Karmada in their `production` environments as of Oct 25, 2022. The full adopter list can be found on our [website](https://karmada.io/docs/next/casestudies/adopters).

Highlighted examples include:

* **[VIPKID](https://www.vipkid.com/)**: Uses Karmada to build a multi-cloud, multi-vendor PaaS platform across regions. The use case details can be found [here](https://karmada.io/docs/next/casestudies/vipkid).
* **[ICBC](http://www.icbc.com.cn/ICBC/EN/)**: Uses Karmada to build a multi-cloud platform to manage **100** Kubernetes clusters and more than **280,000** containers. The use case details can be found [here](https://karmada.io/zh/blog/2022/03/22/multi-cluster-management-with-an-ocean-of-nodes/).
* **[Vivo](https://www.vivo.com/)**: Builds a more resilient and elastic hybrid-multi-cloud platform using Karmada.
* **[Trip.com](https://www.trip.com/)**: Uses Karmada to build a multi-cloud platform with cross-cluster auto scaling and scheduling.
* **[Hurricane Engines](https://orang.cloud/)**: The [ORANG](https://orang.cloud/) low-code platform provides the service for users to build powerful applications quickly with zero lines of code. ORANG uses Karmada to deploy applications across multi-cloud. The use case details can be found [here](https://karmada.io/docs/next/casestudies/ci123).
* **[Huawei Cloud](https://www.huaweicloud.com/intl/en-us/)**: Huawei Cloud's UCS ([Ubiquitous Cloud Native Service](https://support.huaweicloud.com/intl/en-us/productdesc-ucs/ucs_productdesc_0001.html)) adopts Karmada to connect and manage cross-cloud and cross-region clusters and the cloud native applications running in them.
  UCS provides consistent experience in cloud native deployment, management, and application ecosystem. The use case details can be found [here](https://support.huaweicloud.com/intl/en-us/productdesc-ucs/ucs_productdesc_0001.html).

**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

There are 70 project members for Karmada, out of which there are 7 [maintainers](https://github.com/karmada-io/karmada/blob/master/MAINTAINERS.md), 23 [approvers/reviewers](https://github.com/karmada-io/community/blob/main/REVIEWERS.md) from 15 different organizations.

**_Demonstrate a substantial ongoing flow of commits and merged contributions._**

A short summary has been provided in [#Continous Community Momentum](#continous-community-momentum). Detailed statistics can be found in the following devstats links:

* [Contribution Activity](https://karmada.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-1y&to=now): 1700+ per month
* [New contributors since sandbox](https://karmada.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1631635200000&to=now)
* [Issues open/closed since sandbox](https://karmada.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1631635200000&to=now): average 45 issues opened and 35 issues closed per month
* [PRs open/closed since sandbox](https://karmada.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1631635200000&to=now): average 127 PRs opened and 115 PRs closed per month
* [Contributing organizations](https://karmada.devstats.cncf.io/d/5/companies-table?orgId=1): 50+ companies and organizations

**_A clear versioning scheme._**

Karmada uses the [semantic versioning scheme](https://semver.org/#semantic-versioning-specification-semver). We have minor release every three months, for patch releases the cadence is roughly one month, or when critical fixes are required. The latest release is v1.4.0.

We have releases documented at: https://karmada.io/docs/next/releases

**_Clearly documented security processes explaining how to report security issues to the project, and describing how the project provides updated releases or patches to resolve security vulnerabilities_**

Karmada has a responsible and mature vulnerability submission and disclosure program, managed by the Karmada Security Team. The processes are documented at https://github.com/karmada-io/community/blob/main/security-team/SECURITY.md.

## Statement on Alignment with the CNCF Mission

As surveyed, more than 93% of enterprises are using the services of multiple cloud vendors at the same time. Cloud native technologies and the cloud market continue to mature. Multi-cloud and multi-cluster deployments become the choice of future.
Karmada aims to provide turnkey automation for multi-cluster application management in multi-cloud and hybrid cloud scenarios, with key features such as centralized multi-cloud management, high availability, failure recovery, and traffic scheduling.
With Karmada, users can not only evolve from a single-cluster architecture to a multi-cloud one smoothly, but also customize their own multi-cloud platforms with a variety of policy sets.

Karmada strictly aligns with the mission of CNCF in empowering organizations to build and run applications in modern cloud native environments and benefit from the evolving cloud native technologies.

## Future Plan

The Karmada project continues to grow and has its sight on short-term and long-term goals:

* **Multi-cluster security**.
Multi-cluster security presents many challenges.
  - Karmada manages clusters centrally, which eases cluster management but brings some security risks.
    The aim is to minimize the permissions of the control plane and reduce blast radius to member clusters after a control plane failure.
  - Another target is on multi-cluster access control with finer-grained security isolation on the control plane.
* **Larger scalability**.
We have been working on improving user experience in large-scale scenarios with the previous releases.
  We launched a large-scale test on Karmada to obtain baseline performance metrics for Karmada managing multiple Kubernetes clusters.
  The test results show that Karmada can stably support 100 large-scale clusters with 500,000 nodes connected at the same time, running more than 2 million pods.
  Multicloud means larger resource pools, which brings greater performance pressure to the multi-cloud management platform.
  In large-scale scenarios, we will continue to reduce the access latency, resources consumed, and stability impact on the system under high concurrency.
* **Multi-cluster observability**.
Multi-cluster observability is an important part in a multi-cluster orchestration. We will continue to:
  - Explore integrated solutions for multi-cloud monitoring and logging.
  - Optimize multi-cluster events so that users can better maintain the multi-cloud management platform.
* **Multi-cluster application distribution**.
There are still many challenges in application distribution across multiple clusters.
  - For plug-ins coupled with the cluster, we will dive into simplifying the plug-in installation and lifecycle management.
  - When applications have sequential relationships between clusters, we aim to provide a multi-cluster workflow to simplify the deployment of these applications.
* **Integration with ecosystem**.
  The integration with the CNCF ecosystem has always been a major focus of our efforts.
  We actively cooperate with other CNCF projects and helps them better transition to multi-cluster based on Kubernetes native APIs, such as Argo CD and Flux.
  We believe that cooperation in different fields can better serve our users.
  We will continue to seek cooperation with other mature projects to form an open and prosperous ecosystem.

The planed new features are listed in [roadmap](https://github.com/karmada-io/karmada/blob/master/ROADMAP.md).
