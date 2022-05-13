# KubeVela Sandbox Annual Review


## Background

[KubeVela](https://kubevela.io/) is a modern software delivery platform that makes it easier and faster to deploy and operate applications across hybrid, multi-cloud environments. Applications created with KubeVela use the best practices of modern applications by default: they are able to scale with clouds, they use infrastructure as code, they are observable, and they are secure.

Currently, KubeVela primarily provides the following key features:

* **Unified Application Delivery Experience**

    KubeVela introduces a [unified and extensible model (OAM)](https://oam.dev/) that can glue and orchestrate all of your IaC based infrastructure configuration. As a result, KubeVela provides simple user experience for modern application delivery.

* **Automated Deployment across Clusters**

    KubeVela natively supports multi-cluster/hybrid-cloud scenarios such as promotion across clusters, high availability between clusters, automated cloud infrastructure provision.
    
* **Enterprise-Grade Security**

    KubeVela provides enterprise-grade security with fine-grained and custom RBAC, multi-cluster authorization, and first-class LDAP integrations.

* **Centralized Management and Observability**

    KubeVela provides a unified control plane for modeling, provisioning, and deploying applications. The centralized management reduces the burden of looking over each clusters and gives unified experience across platforms. It greatly improve efficiency around troubleshooting and debugging when things go wrong.

* **Declarative and Highly Extensible Workflow**

    KubeVela drives the delivery process with a declarative workflow, it provides automated canary and blue-green deployments with verification and rollback. After the workflow finished, it keeps the reconciliation loops to prevent any unexpected configuration drifts.


### Alignment with CNCF

- KubeVela was accepted as a CNCF Sandbox project on June 22, 2021.

## Development metrics

The KubeVela devstats page and dashboards can be found [here](https://kubevela.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, KubeVela currently has [194](https://kubevela.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [50](https://kubevela.devstats.cncf.io/d/5/companies-table?orgId=1)
companies. On average, there are [~1500 contributions per month](https://kubevela.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) contained within [150 merged PRs per month](https://kubevela.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=prs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).
- [New PRs in last year](https://kubevela.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
- The community has grown since the project entered the CNCF sandbox.
  - We held weekly community meetings constantly (total 74 as of May 2022). The meeting records can be found in [here](https://www.youtube.com/channel/UCSCTHhGI5XJ0SEhDHVakPAA/videos) and [here](https://space.bilibili.com/180074935/channel/seriesdetail?sid=1842207) for different timezone/language community members. The average number of meeting attendees is ~30.
  - Number of contributors: 90+ -> **150+**
  - Github stars: 1900+ -> **3600+**
  - Github forks: 250+ -> **530+**
  - Contributing organizations: 20+ -> **50**

Regarding these metrics, we would perceive the community as healthy and rising.

## Maintainers

We have established [the community membership roadmap](https://github.com/kubevela/community/blob/main/community-membership.md), and quite a few active and qualified contributors
have been promoted to maintainers. Compared to five maintainers (all from Alibaba) initially, we now have [seven maintainers](https://github.com/kubevela/community/blob/main/OWNERS.md#maintainers) from four organizations.

- Jianbo Sun (Alibaba) @wonderflow
- Hongchao Deng (Independent) @hongchaodeng
- Zhengxi Zhou (Alibaba) @zzxwill
- Yue Wang (Tencent) @captainroy-hy
- Qingguo Zeng (Alibaba) @barnettZQG
- Daniel Higuero (NAPPTIVE) @dhiguero
- Da Yin (Alibaba) @Somefive

We also have joined with [13 Reviewers](https://github.com/kubevela/community/blob/main/OWNERS.md#reviewers) and [7 Approvers](https://github.com/kubevela/community/blob/main/OWNERS.md#approvers) in the contribution team from 13 organizations, this would sustain along.

## Project adoption

KubeVela has been adopted as the core of cloud native application delivery solutions. Many public services in Alibaba cloud have fully leveraged KubeVela to build their platform to deliver and manage applications across hybrid, multi-cloud environments.

Beyond this, we have more than 30 end users to adopt KubeVela for developing or production environments. [Here is the list](https://github.com/kubevela/community/blob/main/ADOPTERS.md) which gathers who they are, examples include:

- [Alibaba Cloud](https://www.aliyun.com/)
- [LINE](https://line.me/en/)
- [4Paradigm](https://www.4paradigm.com/)
- [Merchants Bank](http://www.cmbchina.com/)
- [SHEIN](https://www.shein.com/) 
- [Bytedance](https://www.bytedance.com/) 
- [NetEase Games](http://neteasegames.com/) 
- [NAPPTIVE](https://napptive.com)
- [iQiyi](https://www.iqiyi.com/)  
- [Guidewire](https://www.guidewire.com/)  
- [digital china](https://www.digitalchina.com/)  
- [HSBC](https://www.hsbc.com/)  
- [Asiainfo.com](https://www.asiainfo.com/en_us/index.html)  
- [Springer Nature](https://www.springernature.com/)   
- [lilithgame](https://www.lilith.com/)
- [HarmonyCloud](http://www.harmonycloud.cn/) 
- [YOUZAN](https://www.youzan.com/) 
- [Launcher](https://www.lstack.com/) 
- [52tt.com](https://52tt.com/) 
- [deepexi](https://www.deepexi.com/) 
- [Li Auto](https://www.lixiang.com/)  
- [Forchange](https://www.forchange.cn/) 
- [XPeng](https://www.xiaopeng.com/)
- [HUOLALA](https://www.huolala.cn/) 
- [alauda](https://www.alauda.cn/)  
- [linktimecloud](https://www.linktimecloud.com/)  
- [rootcloud](https://www.rootcloud.com/)  
- [jiankangzhilu](https://www.yihu.com/)  
- [cmbyc](https://www.cmbyc.com/)  
- [xshoppy](https://www.xshoppy.com/)  
- [wanwei](http://www.wanwei.com.cn/)  
- [fiberhome](https://www.fiberhome.com/default.aspx)  
- [zhihuiya](https://www.zhihuiya.com/)  

There are more than 2 big corps as Merchants Bank, Guidewire committing their staffs to spend enormous time in collaboration with the KubeVela community. This is truly a sign which shows KubeVela at its core embodies customers' value.


## Project goals

Our primary goal is to make deploying and operating applications across today's hybrid, multi-cloud environments easier, faster and more reliable. In the past year, we have achieved the following development outcomes.
- Release cadence: 4 minor releases, roughly once every three months. The latest release is v1.4.0.
- Key features added to the project (complete changelog can be found in [here](https://github.com/kubevela/kubevela/releases)):
  - 1.1
    - Upgrade the architecture to natively support multi-cluster and hybrid-cloud delivery.
    - Support declarative workflow for multi-cluster promotion and other user defined steps.
    - Support canary and blue-green rollout.
    - Using CUE as the underlying engine for abstraction and workflow.
  - 1.2
    - Support UI console.
    - Support Addon Management.
    - Support CI Integration.
    - Integrate with FluxCD for GitOps.
  - 1.3
    - Support multi-tenancy and isolation.
    - Support configuration management.
    - Support user/project management with API level RBAC.
    - Support SSO login by integrating with dex.
  - 1.4
    - Support multi-cluster authorization globally.
    - Integrate with K3s to run as lightweight Daemon in single node.
    - Integrate with ArgoCD for GitOps.
    - Integrate with Traefik for progressive rollout.
    - Integrate with OCM for pull based multi-cluster capability.

In the near future, we aim to achieve the following technical goals:
- Improve cloud resource provisioning and consuming workflow.
- Enhance the security and provide secure CI/CD capability.
- Integrate with prometheus/opentelemetry/grafana/loki for more observability and automation.
- Improving user experiences.
- A complete project roadmap can be found in [here](https://kubevela.io/docs/roadmap/README).

We have built close connections with members in the CNCF organization for project integration. 
We have organized a few offline meetups for KubeVela community members and sponsored a few academic programs (internship, programming contest etc) with Universities. 
We are looking for more user adoptions and community collaborations constantly.

## How the CNCF can help to achieve the upcoming goals

- More chances to advocate the project.
- More chances to collaborate with other projects in CNCF or even out of CNCF.
- Technical writing support for project documents.

## Incubation readiness

- Yes, we are preparing for the incubation proposal.
