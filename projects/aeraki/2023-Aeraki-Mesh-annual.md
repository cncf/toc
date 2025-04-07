# Aeraki Mesh 2023 Sandbox Annual Review

This is the annual review for the [Aeraki Mesh](https://aeraki.net) project for 2023.

## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Project Perform](#project-perform)
- [Project Goals](#project-goals)
- [Incubation Readiness](#incubation-readiness)
- [Project Links](#project-links)

## Background

[Aeraki Mesh](https://aeraki.net) is an open-source project to manage any layer-7 protocols in
a service mesh. It supports Dubbo, Thrift, Redis, Kafka, ZooKeeper, and any proprietary protocols
can be supported as well with the help of the MetaProtocol framework.

### What does Aeraki stand for ?
Aeraki [Air-rah-ki] is the Greek word for 'breeze'. We hope this breeze can help Kubernetes and Istio
sail further in the cloud native adventure.

### Why Aeraki Mesh ?
While service mesh becomes an important infrastructure for microservices, many(if not all) service mesh
implementations mainly focus on the HTTP protocol and treat other protocols as plain TCP traffic. Aeraki Mesh
has been created to provide a non-intrusive, highly extendable way to manage any layer-7 traffic in an Istio
service mesh, including Dubbo, Thrift, Redis, etc., and private protocols as well.

### How Aeraki Mesh Solves the Problem ?
Aeraki Mesh consists of the following components:

- Aeraki: Aeraki is the control plane component. It provides high-level, user-friendly traffic management rules to operations, translates the rules to envoy filter configurations, and leverages Istio’s EnvoyFilter API to push the configurations to the sidecar proxies. Aeraki also serves as the RDS server for MetaProtocol proxies in the data plane. Contrary to Envoy RDS, which focuses on HTTP, Aeraki RDS is aimed to provide a general dynamic route capability for all layer-7 protocols.
- MetaProtocol Proxy: MetaProtocol Proxy is the data plane component. It provides common capabilities for Layer-7 protocols, such as load balancing, circuit breaker, routing, rate limiting, fault injection, and auth. Layer-7 protocols can be built on top of MetaProtocol. To add a new protocol into the service mesh, the only thing you need to do is to implement the codec interface and a couple of lines of configuration. If you have special requirements which can’t be accommodated by the built-in capabilities, MetaProtocol Proxy also has an application-level filter chain mechanism, allowing users to write their own layer-7 filters to add custom logic into MetaProtocol Proxy.
Dubbo , Thrift , bRPC and a number of other protocols have been implemented based on MetaProtocol. More protocols are on the way. If you're using a close-source, proprietary protocol, you can also manage it in your service mesh simply by writing a MetaProtocol codec for it.

Aeraki Mesh was launched in November 2020, and was accepted into the CNCF as a Sandbox project in June
2022.

## DevStats

 [Aeraki Mesh Devstats](https://aerakimesh.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=now-1y&to=now) shows that contributions have remained strong in the last year. Since joining the CNCF, we have attracted 13 contributors from 11 organizations including Tecent, Huawei, Microsoft, AlaudaInc, etc.

  - [13 contributors](https://aerakimesh.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All)
  - [11 Organizations](https://aerakimesh.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=commits)


## Maintainers

Aeraki Mesh currently has 7 core maintainers from 6 companies.

|                                                                           |                        Name                         |    Company    |            Repos            |
|:-------------------------------------------------------------------------:|:---------------------------------------------------:|:-------------:|:---------------------------:|
| <img width="30px" src="https://avatars.githubusercontent.com/u/1321762">  |    [zhaohuabing](https://github.com/zhaohuabing)    |  Tetrate.io   | aeraki, meta-protocol-proxy |
| <img width="30px" src="https://avatars.githubusercontent.com/u/15240576"> |  [huanghuangzym](https://github.com/huanghuangzym)  |    Alauda     |     meta-protocol-proxy     | 
| <img width="30px" src="https://avatars.githubusercontent.com/u/3876735">  |         [smwyzi](https://github.com/smwyzi)         |     Baidu     |     meta-protocol-proxy     |
| <img width="30px" src="https://avatars.githubusercontent.com/u/12084604"> |     [whitefirer](https://github.com/whitefirer)     | Tencent Music |           aeraki            |
| <img width="30px" src="https://avatars.githubusercontent.com/u/28644213"> | [Sad-polar-bear](https://github.com/Sad-polar-bear) |  BOSS Zhipin  |           aeraki            |
| <img width="30px" src="https://avatars.githubusercontent.com/u/48784001"> |        [Xunzhuo](https://github.com/Xunzhuo)        |    Tencent    |           aeraki            |
| <img width="30px" src="https://avatars.githubusercontent.com/u/19339970"> |     [tanjunchen](https://github.com/tanjunchen)     |     Baidu     |           aeraki            |

We hope to see more diversity in the maintainer team.

## Adoption

We encourage individuals to submit a comment on [Who is using Aerkai Mesh?](https://github.com/aeraki-mesh/aeraki/issues/105)
to document their use cases. As of now, there're a dozen companies that we know have deployed Aeraki Mesh in their
own infrastructure or integrated Aeraki Mesh as part of their products. Some of the most well-known names include:

- Tencent
- Huawei
- Baidu
- BOSS Zhipin
- Tencent Music
- Alauda
- Shopline


## Project Perform

- *Releases* - Since joining the CNCF, Aeraki Mesh has published 3 major releases. Now, Aeraki Mesh can provide the same traffic management and observability capabilities to non-HTTP protocols as the HTTP protocol.
- *Expand and broaden the Aeraki Mesh community.* - Aeraki Mesh has increased the breath of maintainers and contributors.
- *Grow adoption* - Aeraki Mesh has been adopted by big companies such as Tencent, Baidu, and Huawei.

## Project Goals

The goal of Aeraki Mesh this year is to increase the diversity in both contributors and adoption.
Some key works include:
- Improve the user experience in installation, upgrade and troubleshooting.
- Provide more document about the user guide, tutorial and internals.
- Host or join more events to attract more contribution and adoption.

## Incubation Readiness

   Aeraki Mesh does not currently meet all criteria for incubation, but is on track to do so. Some work needs to be done to achieve the [Best Practices Badge](https://bestpractices.coreinfrastructure.org/), and we also have to define security processes for the project on how to report and address security issues.

## Project Links

* [Website](https://aeraki.net)
* [Github](https://github.com/aeraki-mesh)
