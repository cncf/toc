# OpenKruise Sandbox Annual Review

## Background

[OpenKruise](https://openkruise.io) is an extended component suite for Kubernetes,
which mainly focuses on application automations, such as deployment, upgrade, ops and availability protection.
Most features provided by OpenKruise are built primarily based on CRD extensions.
They can work in pure Kubernetes clusters without any other dependencies.

Currently, OpenKruise primarily provides the following capabilities:
- **Advanced workloads**, which support not only the basic features that are similar to the upstream Workloads in Kubernetes, but also more advanced abilities such as in-place update, configurable scale/upgrade strategies, parallel operations.
- **Sidecar container management**, which defines, injects and even upgrades sidecar containers with no effect on application containers.
- **Multiple domain management**, which empowers workload to support multi-domain and elastic deployment, so that users can define the rules about how their applications should be deployed over different kinds of nodes, e.g., nodes across availability zones.
- **Enhanced operations**, such as restarting containers in-place, pre-downloading images on specific nodes, controlling containers starting priority in a Pod and distributing resources over multiple namespaces.
- **Application availability protection**, which can prevent unexpected Kubernetes resources deletion during cascading deletion and prevent application from disruption or SLA degradation in voluntary disruption scenarios.

### Alignment with CNCF

- OpenKruise was accepted as a CNCF Sandbox project on Nov 10, 2020.

## Development metrics

The OpenKruise devstats page and dashboards can be found [here](https://openkruise.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, OpenKruise currently has [92](https://openkruise.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [39](https://openkruise.devstats.cncf.io/d/5/companies-table?orgId=1)
companies. On average, there are [28 commits per month](https://openkruise.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) contained within [18 merged PRs per month](https://openkruise.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now).
- [New PRs in last year](https://openkruise.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
- A few highlights since the project joined CNCF:
  - Number of contributors: 50+ -> **90+**
  - Github stars: 1700+ -> **2800+**
  - Github forks: 220+ -> **460+**
  - Docker image Pulls: **5M+**
  - Member of DingTalk group + Slack channel: 600+ -> **1000+**
We have held bi-weekly community meetings and the meeting notes and agenda can be found in [here](https://shimo.im/docs/gXqmeQOYBehZ4vqo). 

## Maintainers

[The community membership guidance](https://github.com/openkruise/community/blob/master/community-membership.md) has been established for a long time. Based on this, we have significantly expanded our maintainers. Currently, we have [seven maintainers](https://github.com/openkruise/community/blob/master/MAINTAINERS.md) from five organizations, which demonstrates the growth of the community considering the fact that we only have three maintainers (all from Alibaba) initially.

| Maintainer       | GitHub ID                                     | Affiliation |
| ---------------- | --------------------------------------------- | ----------- |
| Fei Guo          | [Fei-Guo](https://github.com/Fei-Guo)         | Alibaba     |
| Siyu Wang        | [FillZpp](https://github.com/FillZpp)         | Alibaba     |
| Zhen Zhang       | [furykerry](https://github.com/furykerry)     | Alibaba     |
| Robert Everson   | [reverson](https://github.com/reverson)       | Lyft        |
| Henry Wang       | [henrywangx](https://github.com/henrywangx)   | Tencent     |
| Shanjie Han      | [hantmac](https://github.com/hantmac)         | MeiTuan     |
| Yan Shi          | [shiyan2016](https://github.com/shiyan2016)   | Trip        |

## Project adoption

The public adopters that used OpenKruise in production can be found [here](https://github.com/openkruise/kruise/issues/289). Examples include:
- **DouyuTV (China)**: Using Kruise for managing statefulset applications. They need parallel updates.
- **Sto (China)**: Using Kruise for managing statefulset applications, They need inplace upgrades.
- **Boss直聘 (China)**: sidecar,cloneset,statefulset for Corporate Paas.
- **hangyinxiaofei (China)**: sidecar,cloneset,statefulset for Corporate Paas.
- **vanyitech (China)**: Using Kruise for managing cloneset applications, they need inplace upgrades. Using BroadcastJob to clean images.
- **Lyft (US)**: Using Kruise CloneSets and bin packing for selective downscaling of stateful nodes.
- **Ctrip (China)**: Using CloneSet and AdvanceStatefulSet for managing stateless and stateful apps. They need inplace and partial updates.
- **Dmall (China)**: Advance HPA, In-Place Update, Sidecar management, etc.
- **Bringg (Israel)**: Advanced DaemonSet, to solve load balancer connection to api gateway running with daemonset.
- **xiaohongshu (China)**: Using CloneSet and BroadcastJob for managing stateless apps and async job.
- **bixin (China)**: Using CloneSet and BroadcastJob for managing stateless apps and async job.
- **VIPKID (China)**: Using BroadcastJob for managing clusters.
- **zhangmen (China)**: Using CloneSet and AdvanceStatefulSet to deploy apps on production env. They need inplace and partial updates.
- **ihomefnt (China)**: Using sidecarset to inject container.
- **Arkane Systems (US)**: Using BroadcastJob/AdvancedCronTab for periodic per-node tasks and cleanup; using SidecarSet to inject local maintenance containers.
- **永辉科技中心 (China)**: Using CloneSet and SidecarSet to in-place upgrade and manage sidecar pods.
- **OPPO (China)**: Using Kruise for managing large-scale statefulsets to in-place update.
- **Spectro Cloud (US)**: Using Kruise for managing periodic node actions.
- **Suning (China)**: Using CloneSet and SidecarSet to manage stateless apps and sidecar.
- **Deepexi (China)**: Using Kruise for managing large-scale applications and sidecars.
- **哈啰出行 (China)**: In-Place Update, Sidecar management, etc.
- **joyy (China)**: Using CloneSet and AdvanceStatefulSet for managing stateless and stateful apps. 
- **Mobvista (China)**: Using Kruise for log collect sidecarsets and manage stateless cloneset apps.
- **深圳凤凰木网络有限公司 (China)**: Using CloneSet for managing stateless apps.
- **xiaomi (China)**: Using CloneSet for stateless apps.
- **Netease (China)**: Using SidecarSet to in-place upgrade and manage sidecar pods.
- **MeiTuan (China)**: Using CloneSet and AdvanceStatefulSet for managing stateless and stateful apps. 
- **Shopee (China)**: Using CloneSet for stateless apps.
- **Esign (China)**: Using sidecarSet for sidecar container.
- **Wholee (China)**: Using CloneSet for managing stateless apps.

## Project goals

When developing OpenKruise, we aim to relieve the pain of managing large scale cloud native applications in Kubernetes for developers and cluster operators in production environments. The provided features and solutions are expected to be easily integratable for existing Kubernetes use cases.  In the past year, we roughly followed three month release cycle and have accomplished the following development goals:

- Release cadence: 4 minor and 1 major releases. The latest release is v1.0.0.
- Key features added to the project (complete changelog can be found in [here](https://github.com/openkruise/kruise/blob/master/CHANGELOG.md)):
  - 0.7.0
    - New workload: AdvancedCronJob, which supports multiple kinds of job template.
    - Promote Advanced StatefulSet API to v1beta1.
  - 0.8.0
    - A new component named kruise-daemon, which will be deployed on every node.
    - New CRDs: NodeImage and ImagePullJob to support pre-download images on specific nodes.
    - Refactor the controller and webhook for SidecarSet.
  - 0.9.0
    - New feature: ContainerRecreateRequest, which supports restarting containers in existing pods.
    - New feature: Deletion Protection, which protects Kubernetes resources from the cascading deletion mechanism.
  - 0.10.0
    - New feature: PodUnavailableBudget, which can achieve the effect of preventing ALL application disruption or SLA degradation, including pod eviction, deletion, inplace-update and so on.
    - New feature: WorkloadSpread, which supports to constrain the spread of stateless workload, which empowers single workload the abilities for multi-domain and elastic deployment.
  - 1.0.0
    - InPlace Update for container environments.
    - Distribute resources over multiple namespaces.
    - Container launch priority.
    - Kubectl-kruise command line tool.

We are running our project in a fairly open model so that we can react to the needs of our community promptly. In addition, we have recently introduced a [roadmap](https://github.com/openkruise/kruise/projects/2) page in github so that the community members can easily track where we are going. There are still quite a few critical and high demanding features that are under development, including but not limited to the following:
- A new CRD named PodMarker to mark pods by number, labels, nodes, probes, etc.
- Decoupled liveness probe which supports restart containers on failure with limited rate.
- Binding nodes for statefulset pods after being recreated.
- Enhanced application rollout for traffic shifting and batch upgrades.

Besides, we have refactored [OpenKruise website](https://openkruise.io/) using Docusaurus v2 and tidy the documents in it this year. We plan to add more examples and best practices to help newcomers learn and use OpenKruise more easily.

OpenKruse have been integrated by several other open source projects.
- [KubeVela](https://kubevela.io/), a modern application delivery platform and CNCF Sandbox project, has provides [templates to work with OpenKruise workloads](https://github.com/oam-dev/kubevela/tree/master/vela-templates/addons/kruise).
- [MOSN]](https://mosn.io/en/), a popular cloud-native network proxy in China, is working on [using SidecarSet of OpenKruise to hot upgrade the proxy container](https://github.com/mosn/mosn.io/pull/166).

We are looking for more cross community collaborations within and outside of CNCF.

## How the CNCF can help to achieve the upcoming goals

- More channels to advocate the project within CNCF and even to other foundations.
- Writing support for project documents.

## Incubation readiness
- Yes, we are preparing for the incubation proposal.
