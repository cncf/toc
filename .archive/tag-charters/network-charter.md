# CNCF TAG Network

Primary Authors: Lee Calcote

Reviewed and Contributed to by: Matt Klein, Ken Owens

# Introduction

With the increased prevalence of microservice-based distributed systems, networking as a discipline has never been more critical in the efficient operation of modern software deployments. Networking primitives, including load balancing, observability, authentication, authorization, policy, rate limiting, QoS, mesh networks, legacy infrastructure bridging, and so on are now receiving substantial development and investment throughout the industry.

# Mission Statement

TAG Network&#39;s charter is inspired from the [CNCF Networking WG](https://github.com/cncf/wg-networking). With an ever steady eye to the needs of workloads and developers who create them and operators who run them, TAG Network&#39;s mission is to enable widespread and successful development, deployment and operation of resilient and intelligent network systems in cloud native environments through these activities to:

1. **Clarify and inform.** Provide valuable and objective information to the TOC, End Users and Projects of the CNCF regarding areas considered in-scope. Strengthen the project ecosystem to meet the needs of end users and project contributors. Educate and inform users with unbiased, effective, and practically useful information.
2. **Collaborate and interrelate.** Effectively interface with other related groups internal and external to the CNCF and connect the dots to facilitate meaningful collaborative progression of relevant topics. Engage more communities and create an on-ramp to effective TOC contribution &amp; recognition.
3. **Assist and attract projects.** Helping to maintain the continued health of CNCF networking-related projects. Focus attention and resources on helping foster project maturity, systematically across CNCF projects. Clarify relationship between projects, CNCF project staff, and community volunteers. Identify gaps in the landscape of CNCF networking-related projects. Find and attract projects to fill these gaps through invitation to present, diligence on proposed projects, and in essence, act as a funnel for TOC project reviews in this area.
4. **Afford impartial stewardship.** Provide and maintain a vendor neutral venue for relevant thought validation, discussion, and project feedback. Establish procedures to assist where necessary.

## In Scope

We strive to understand the fundamental characteristics of different networking approaches with respect to availability, scalability, performance, consistency, observability, security, ease-of-use, cost, performance and operational complexity; and relate these to their suitability to various cloud native use cases. Networks, network services and methodologies suitable for and commonly used in [modern cloud-native environments](https://github.com/cncf/toc/blob/main/DEFINITION.md) are considered in scope:

- Especially where these differ significantly from network systems and approaches previously commonly used in traditional enterprise data center environments.
- Consideration for public, private, and hybrid cloud environments.
- Network functions and services, including but not limited to, the discovery, monitoring, management, routing, load balancing, firewalling, quality of service, provided by service proxy, API gateway, service mesh, remote procedure call, coordination &amp; service discovery.
- Provide an umbrella TAG for Networking WG and Universal Data Plane API Working Group (UDPA-WG).
  - [Networking WG](https://github.com/cncf/wg-networking) will fold directly into TAG Network.
  - UDPA-WG will continue with existing calls and mailing list.
- Facilitating an understanding of network abstraction APIs for meshes, proxies and other network services.
- Influence and provide guidance on the definition of network abstraction specifications. Network TAG may review and recommend adoption.
- Networking extends to services and workloads not running on Kubernetes. Network TAG focuses on the integration and interoperability of networking running those workloads, not on those (physical) networking systems themselves unto their own.

### Current CNCF Network-centric Projects

Generally, projects listed in the [CNCF Landscape](https://landscape.cncf.io/?group=projects-and-products&view-mode=grid&tag=network) under the categories of Service Mesh, API Gateway, Coordination &amp; Service Discovery, Service Proxy, Remote Procedure Call, and Cloud Native Network are considered in-scope of this TAG.

## Out of Scope

Generally anything not considered in scope. The TAG may touch on physical networking hardware as it relates to the running of or interoperability of cloud native networking, but is not a stated focus of the TAG unto its own. See also Overlap and Relations with other Related Groups.

# Overlap and Relations with other Groups and Projects

The area of networking interfaces with essentially all areas of the cloud native landscape.

- Current CNCF Network-centric Projects - we maintain close communication with these projects and their communities to avoid unnecessary duplication of effort and inconsistent messaging wherever possible.
- Kubernetes Network SIG - is focused towards Kubernetes-specific network abstractions, interfaces, and implementations of these interfaces. We maintain close communication with this Kubernetes SIG, with several individuals actively involved in both. Our aim is to avoid unnecessary duplication of effort by the two groups, and maintain clear and consistent messaging by the two groups to our end user community and projects.
- CNCF Security TAG - works on the more general area of cloud-native security including authentication, authorization, encryption, accounting, auditing and related topics. Each of these topics are directly applicable and often implemented as network services. We defer as much as possible to this group to deal with general security-related issues, and liaise closely with them on how to deal with network-specific security as these are a continual area of concern.
- CNCF App Delivery TAG - will be focused on the development, deployment, operation and testing of cloud-native applications. We collaborate with this TAG where this pertains to Networking and Traffic, particularly around application deployment, debugging, and monitoring.
- CNCF Storage TAG - primarily with respect to service discovery as cloud native name services as related etcd.
- Workload Profile Definition and Measurement Project - this project aims to define an industry standard definition for workload profiles and a standard set of measures and metrics that represent those profiles with the goal of developing an open source standard benchmarking to validate these. The CNCF Network TAG will engage as and when a network-related specification or set of benchmarks is presented.
- [Service Mesh Performance](https://smp-spec.io) - this project defines an industry standard definition for capturing service mesh performance profiles and a standard set of measures and metrics that represent those profiles with the goal of developing an open source standard index to understand these.
- Multi-Vendor Service Mesh Interoperation - specifies a set of API standards for enabling service mesh federation.

# Responsibilities &amp; Deliverables

Clarify and further concepts of cloud native networking, traffic management and common design patterns and best practices of these network services. Explore definition and practicality of new cloud native networking technologies focused on serving distributed applications.

## Responsibilities

- Identification of areas in need of focus in cloud native networking.
- Ensure perspectives of users, customers, and vendors are represented.
- Facilitate an understanding of the boundaries between related **network technologies**, including but not limited to API gateways, client libraries, service meshes, proxies, service and name directories, management, control and data planes, policy, security, and so on.
- Facilitate an understanding of **network services** and an expansion of the definition of what these are as existing types are reincarnated or new types emerge in a cloud native way.
- Evaluation of project proposals and facilitation of their diligence. Establish a framework for evaluation.
- Work on integrating common tooling into different projects, particularly where that tooling is a CNCF project (but the targets may not be).
- Growing CNCF external relationships with interested parties, e.g. NIST, ONAP, LFEdge and other security standards bodies.
- Develop and maintain a relevant landscape and whitepapers. Provide common design patterns and best practices.

## Deliverables

- **Proposed project review** - a pre-review of upcoming project proposed for adoption, allowing projects to have a soft review, gauge and garner CNCF interest.
  - Provide recommendations for new CNCF projects.
- **New projects** - recommendation of sandbox projects where exploration is needed.
- **TAG reports to TOC** - scheduled regular reporting to ToC on ongoing and completed work.
- **Cloud Native networking whitepaper(s)** - framework document for the cloud native traffic and networking space.
- Cloud Native Networking Patterns
- **Network Landscape** - a deeper and category-specific project landscape (e.g. [service mesh landscape](https://layer5.io/landscape)).
- **Existing project review** - review of previously adopted projects as a health check of projects in the category.

# Governance &amp; Operations

This TAG follows the [standard operating model](https://github.com/cncf/toc/blob/main/tags#operating-model) provided by the TOC unless otherwise stated here.

