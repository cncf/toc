# SuperEdge 2022Annual Review

This is the annual review for the [SuperEdge](https://superedge.io) project for 2022.

## Background

SuperEdge is an open source **container management system for edge computing** to manage compute resources and container applications in multiple edge regions. These resources and applications, in the current approach, are managed as one single **Kubernetes** cluster. A native Kubernetes cluster can be easily converted to a SuperEdge cluster.

SuperEdge has the following characteristics:

* **Kubernetes-native**: SuperEdge extends the powerful container orchestration and scheduling capabilities of Kubernetes to the edge. It makes nonintrusive enhancements to Kubernetes and is fully compatible with all Kubernetes APIs and resources. Kubernetes users can leverage SuperEdge easily for edge environments with minimal learning.

* **Edge autonomy**: Now SuperEdge provides L4/L5 edge autonomy. 
  
  L3: When the network connection between the edge and the cloud is unstable, or the edge node is offline, the node can still work independently. But at this time, edge node can't do anything associated with writing operation, such as create/delete/update.
  
  L4/L5: Now SuperEdge has released v0.9.0 which support Kins [[English](docs/components/kins.md), [中文](docs/components/kins_CN.md)] Feature. Kins will provision a lightweight K3s cluster based on SuperEdge, which can be operated totally offline. L4 level will leverage the single master edge K3s cluster, and L5 level will provision 3 master edge K3s cluster.

* **Distributed node health monitoring**: SuperEdge provides edge-side health monitoring capabilities. SuperEdge can continue to monitor the processes on the edge side and collect health information for faster and more accurate problem discovery and reporting. In addition, its distributed design can provide multi-region monitoring and management.

* **Built-in edge orchestration capability**: SuperEdge supports automatic deployment of multi-regional microservices.
  Edge-side services are closed-looped, and it effectively reduces the operational overhead and improves the fault
  tolerance and availability of the system.

* **Network tunneling**: SuperEdge ensures that Kubernetes nodes can operate under different network situations. It
  supports network tunnelling using TCP, HTTP, HTTPS and SSH.

SuperEdge was initiated by the following companies: Tencent, Intel, VMware, Huya, Cambricon, Captialonline and Meituan.

## Development Metrics

The SuperEdge devstats page and dashboards can be found [here](https://superedge.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, SuperEdge currently has [58](https://superedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [16](https://superedge.devstats.cncf.io/d/5/companies-table?orgId=1)
  companies. 
- There are the [New PRs in last year](https://superedge.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
- Most of the contributions come from maintainers who work at Tencent. Since joining the CNCF:
  * Number of stars: 591 -> 935
  * Number of forks: 137 -> 196

## Maintainers

SuperEdge currently has 3 core maintainers. Since joining the CNCF,
3 contributors become maintainers.

The original maintainers are as follow:

| Maintainer Name | Github Name                                   | Company           |
|:---------------:|:---------------------------------------------:|:-----------------:|
| Gerald Wang     | [neweyes](https://github.com/neweyes)         | Tencent           |
| Fighting Du     | [duyanghao](https://github.com/duyanghao)     | Linklogis         |
| Attlee Wang     | [attlee-wang](https://github.com/attlee-wang) | Tencent           |
| Random Cheng    | [Beckham007](https://github.com/Beckham007)   | Tsinghua Unigroup |
| Kaiyue Chen     | [chenkaiyue](https://github.com/chenkaiyue)   | Tencent           |
| Shubiao Yao     | [shubiao-yao](https://github.com/shubiao-yao) | Sina              |
| Fee Li          | [00pf00](https://github.com/00pf00)           | Tencent           |
| Yiwei Chen      | [yiwei-C](https://github.com/yiwei-C)         | Tencent           |
| Wenhu Wang      | [wenhuwang](https://github.com/wenhuwang)     | ZTO               |
| Roy Liang       | [lianghao208](https://github.com/lianghao208) | Tencent           |

The new members are as follow:

| Maintainer | GitHub Username                              | Company |
| ---------- | -------------------------------------------- | ------- |
| malc0lm    | [@oilbeater](https://github.com/malc0lm)     | Tencent |
| dodiadodia | [@dodiadodia](https://github.com/dodiadodia) | Tencent |
| zizhuoy    | [@zizhuoy](https://github.com/zizhuoy)       | Tencent |

We hope to see more diversity in the maintainer team.

## Adoption

We didn't strictly distinguish the adoption before and after joining the CNCF.
Belows are just the list of companies that we know have deployed SuperEdge in their own infrastructure or integrated SuperEdge as part of their products.

- NIO
- XDF.CN
- SANY
- SHUKUN
- ZOOMLION
- ZUOYEBANG.com
- Tencent Meeting

## Project perform

Since joining the CNCF, SuperEdge has published key features:

- NodeUnit and NodeGroup enhancement to use CRD to manipulate the Node Pools
- Lite-apiserver enhancement to cache critical resources with the purpose of reducing network traffic
- Application-grid-wrapper support to watch `Ingress` resource to enable `nginx-ingress-controller` at NodeUnit
- Reconstruct `Tunnel` architecture to support Cloud and Edge Level 7 communication
- Originally release `Kins` feature to support whole lifecycle management of edge K*s cluster

## Project Goals

The goal of SuperEdge this year is to enhancement Cloud-edge collaboration, including:

- Enhance Kins feature to improve user experience and usability
- Combine various edge scenarios to develop corresponding plugins, such as support edge AI device(Orin, etc)
- Reconstruct SuperEdge installation tools, using helm to install SuperEdge modules on standard K8s cluster
- Provide more document about the user guide and internals 

## Incubation readiness

We think SuperEdge still needs further exploration to get border adoption in Edge Industry. We are looking forward to meeting the incubation criteria in near future.
