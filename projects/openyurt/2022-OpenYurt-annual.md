# OpenYurt Annual Review

## Background

[OpenYurt](https://openyurt.io/) is an extension of the upstream Kubernetes for edge computing use cases.
It has been designed to meet various DevOps requirements against typical edge infrastructures. With OpenYurt,
users are able to manage their edge applications as if they were running in the cloud infrastructure.
It addresses specific challenges for cloud-edge orchestration in Kubernetes such as unreliable or disconnected
cloud-edge networking, edge node autonomy, edge device management, region-aware deployment and so on.
Currently, OpenYurt primarily provides the following capabilities:
- A node daemon, YurtHub, which acts as an agent between api-server and local daemons such as kubelet, could fulfill
  requirements of edge autonomy, service topology optimization and traffic consolidation, etc.
- NodePool management for region-aware application deployment and higher level autonomy raised to a bunch of assimilated nodes as a whole.
- A set of controllers to deal with cloud edge collaboration such as node and NodePool autonomy, OTA upgrade, networking setup, etc.
- An unified networking solution to enable both north-south and east-west communication.
- An integration with [EdgeX Foundry](https://github.com/edgexfoundry) to allow managing edge devices through Kubernetes.

### Alignment with CNCF

- OpenYurt was accepted as a CNCF Sandbox project on Sept 8, 2020.

## Development metrics

The OpenYurt devstats page and dashboards can be found [here](https://openyurt.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, OpenYurt currently has [134](https://openyurt.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [36](https://openyurt.devstats.cncf.io/d/5/companies-table?orgId=1)
companies. On average, there are [82 commits per month](https://openyurt.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) contained within [47 merged PRs per month](https://openyurt.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now).
- [New PRs in last year](https://openyurt.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
- The community has grown since the project entered the CNCF sandbox.
  - We held weekly community meetings constantly (total 98 as of May 2023). The meeting records can be found in [here](https://search.bilibili.com/video?keyword=openyurt). The average number of  meeting attendees is ~20.
  - Number of contributors: 40+ -> **134**
  - Github stars: 1000+ -> **1500+**
  - Github forks: 200+ -> **300+**
  - Contributing organizations: 22 -> **36**

## Maintainers

We now have [eleven maintainers](https://github.com/openyurtio/openyurt/blob/master/MAINTAINERS.md) from eight organizations.
- Chao Zheng (ByteDance) @charleszheng44
- Chenglong Wang (Inspur) @luckymrwang
- Fei Guo (Microsoft) @Fei-Guo	
- Lifang Zhang (China Telecom) @DrmagicE
- Linbo He (Alibaba) @rambohe-ch
- Shaoqiang Chen (Intel) @gnunu
- Shida Qiu (Alibaba) @SataQiu
- Tao Chen (Zhejiang University) @qclc
- Wuming Liu (VMware) @lwmqwer
- Yixing Jia (VMware)	@yixingjia
- Zhengguang Zhang (Independent) @zzguang

## Project adoption

OpenYurt has been adopted as the foundation of public cloud Kubernetes edge solutions. Many public services in Alibaba cloud,
CDN and IoT for example, have fully leveraged OpenYurt to manage their edge infrastructures spread across multiple regions.

Some publicly announced adopters include:
- Alibaba Cloud: Uses OpenYurt as base framework and integrated with other cloud services(like SLB etc.) to provider hosted edge kubernetes service.
- Sangfor Technologies Inc: A company that focuses on providing security services, and uses OpenYurt for edge autonomy and operation and maintenance communication.
- China Telecom: Use OpenYurt for managing edge nodes across different idc region and network.
- Sony Group Corporation: Uses Openyurt on internal PoC for cloud and edge container orchestration.
- Lixiang Auto Inc: A company that designs and produces new energy vehicles, and uses OpenYurt for managing edge nodes, deploying edge applications.
- Shanghai Cue Co.,Ltd: Uses OpenYurt to manage retail edge AIBoxes.

As for now, there are registered 20+ public adopters. We can predict the number will keep growing.

[Here](https://github.com/openyurtio/openyurt/issues/300) is an adopter list.

## Project goals

Our primary goals include:
- Extending Kubernetes to solve specific challenges for cloud-edge orchestration.
- Be platform agnostic and non-intrusive, meaning users can enable all OpenYurt capabilities for a cloud cluster or on-premise cluster easily.
- Highly collaborative for integrating other edge computing platforms or technologies, such as IOT solution EdgeX Foundry.
- Build an open and organized community for resolving cloud-edge orchestration problems and leveraging the Kubernetes ecosystem.

We set up three SIGs according to our goals: SIG-control-plane for edge management, SIG-data-plane for networking, and SIG-iot for device management.

in the past year, we have acquired these achievements:
- SIG-control-plane:
  - Add governance capability for NodePool
  - Multiplexing cloud-edge traffic
  - Daemonset Update Model
  - Yurthub Cache Model Refactoring
  - Add edge autonomy e2e tests
  - Yurtadm join Refactoring
  - Enhance the service topology ability of YurtHub
  - Static Pod Upgrade Model
  - Install openyurt components using helm
  - Optimize the pods recovery efficiency when edge nodes restart
- SIG-data-plane:
  - Unify cloud edge communication solution for OpenYurt
- SIG-iot:
  - EdgeX AutoCollector for infusing EdgeX upstream update and verification
  - EdgeX LTS Support

Major releases last year:
- 1.0.0
  - nodepool API upgrade to v1beta1
  - code quality improvement
  - installation upgrade, helm support added
- 1.1.0
  - Support OTA/Auto upgrade model for DaemonSet workload
  - Support autonomy feature validation in e2e tests
  - Improve Yurthub configuration for enabling data filter function

The latest release of OpenYurt is 1.3.1.
  
In the near future, we aim to achieve the following technical goals:
- SIG-control-plane:
  - Support independent configuration in each NodePool.
  - Support multiple tenant isolation management.
  - Improve cluster scalability to contain much more worker nodes.
- SIG-data-plane:
  - Support SLB as public network exporter for gateway Raven.
  - Support http(s) request forwarding based on node name to deal with possible problems caused by edge node IP duplication.
  - Improve route configuration(policy route + ipset + iptables mark) based on Pod IPs instead of subnets of nodes.
- SIG-iot:
  - Abstract a common interface, Yurt-Device-Interface, to support different IoT solutions to connect with OpenYurt, thereby supporting a richer IoT ecosystem.
  - Through the Yurt-Device-Interface common interface, provide a set of benchmarks to test the performance of the connected IoT solution.
  - Add a Yurt-IoT-Manager to integrate more IOT platforms.
- Improving user experiences. A new SIG, SIG-ui has been set up to improve overall OpenYurt installation and ease of use.
- A complete project roadmap can be found in [here](https://github.com/openyurtio/community/blob/main/roadmap.md).

## Incubation readiness
- Yes, we have submitted [due diligence document](https://docs.google.com/document/d/1St_JIS6Ntt1q39eOyUYs7U5enXIiYhVY1-6qb6b-6Rs) for the incubation review.
