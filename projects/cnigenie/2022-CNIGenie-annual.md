# CNI Genie 2022 Annual Review

This is the annual review for the [CNI Genie](https://github.com/cni-genie/CNI-Genie) project for 2022.

## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Project goals](#project-goals)
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)

## Background

CNI-Genie is a multi-networking plugin for Kubernetes that is very easy to configure. It makes container orchestrator to seamlessly connect to the user's choice(s) of CNI plugins installed on the host.

Here are some key features already supported by CNI-Genie:

* Dynamic Plugin Selection: Using multiple CNI plugins for different pods in same cluster.
* Multi-IP Assignment: Multiple interfaces per containers with IPs coming from different plugins.
* Logical Network CRD Objects: Realization of different provisions w.r.t IP management. (Eg. customized subnet usage)
* Network Level Policy: Application of network policy at network level.
* Support for Kubernetes NPWG multi-network spec v1.
* Default Networks: Configuring default network(s) once for each of the pods coming up.


CNI Genie consists of 3 major components:

* Genie Plugin: Manages all multi-IP related provisions.
* Network Admission Controller: Validation of logical network CRD objects.
* Network Policy Controller: Application of policies at network level.

## DevStats

CNI Genie DevStats dashboard [here](https://cnigenie.devstats.cncf.io/d/4/company-statistics-by-repository-group?orgId=1). For the past couple of years, progress and activity on this project had stalled after the company had to reprioritize resources following the events that unfolded in 2019 and its ensuing consequences.

Now, with some of those uncertainties behind us, we are in a better position to commit resources and are now looking to redirect & reinvigorate this project to address the needs we have identified for our customers and for the community. We have a new maintainer for the project with fresh ideas based on the latest technologies. We are also looking to staff up some more to help implement these ideas. We continue to remain a sandbox project at this time.


## Maintainers

CNI Genie currently has a [new maintainer](https://github.com/cni-genie/CNI-Genie/blob/master/MAINTAINERS) from Futurewei Technologies:

Current Maintainer(s):

| Maintainer       | GitHub Username                               | Company        |
| ---------------- | --------------------------------------------- | -------------- |
| Vinay Kulkarni   | [@vinaykul](https://github.com/vinaykul)      | Futurewei, USA |

Initial Maintainers:

| Maintainer          | GitHub Username                                          | Company       |
| ------------------- | -------------------------------------------------------- |---------------|
| Jun Du              | [@m1093782566](https://github.com/m1093782566)           | Huawei, China |
| Shushant Kumar      | [@sushanthakumar](https://github.com/sushanthakumar).    | Huawei, China |
| Zefeng (Kevin) Wang | [@kevin-wangzefeng](https://github.com/kevin-wangzefeng) | Huawei, China |


## Adoption

Unfortunately, no adoption data is available at this time.


## Project Goals

### Previous Goals

CNI-Genie was originally designed to enable multihoming for Kubernetes pods by enabling users to specify desired number of interfaces and the respective CNI drivers for those interfaces.

As mentioned above, this project had stalled for over two years due to resource reprioritization following 2019 events. So we have not tracked how we did against [goals](https://github.com/cni-genie/CNI-Genie/blob/master/ROADMAP-old.md) that were earlier set.

While the initial goals to offer CNI choice & flexibility were great, users really only care that they get reliable multihoming capability with performant data-plane (traffic throughput) and control-plane (low network-ready latencies for pods) out of box that works well. This became clear from a 2022 KubeCon EU stories. One talk in particular discussed the complexities they encountered using alternative solutions for multihomed networking. In the end, the other solution did not work because they designed their network for scale out rather than scale up. The ability to pick and choose CNI drivers becomes a much more appealing feature with a default that intuitively works well when networks start scaling.

### New Approach

For the past couple of years, we architected and built a pod networking solution
based on eBPF/XDP called the [Mizar project](https://github.com/centaurus-cloud/mizar).
Mizar was designed for fast data-plane performance by relying on eBPF/XDP to provide
the overlay networking that completely by-passes the host network stack to ferry
traffic between containers.

It was also built with a control-plane design to provide low-latency network-readiness for pods in order to handle the cloud native networking needs where pods rapidly come and go. Mizar also provides native multi-tenancy network isolation and was designed for scale out networking. The goal at that time was to provide a CNI networking solution for our scale out pod orchestration solution called project [Arktos](https://github.com/centaurus-cloud/arktos). We recently successfully integrated Mizar and Arktos and also demonstrated its multi-tenant networking capabilities in Arktos scaleout architecture at the Linux Foundation Open Source Summit in Austin, TX in June 2022.

We now realize that Mizar's eBPF/XDP technology can also address the critical cloud networking problems that we and others in the community face with multi-homed networking at scale.

#### New Goals

We have identified the following goals to integrate select Mizar's features into CNI-Genie:

- Add out-of-box fast & scalable eBPF/XDP based pod networking capability.
- Add ability for users to select the isolated networks to connect their pods into.
- Allow users to operate multiple groups of pods in their own isolated networks.
- Eliminate the (per-packet) overhead of network policies to achieve isolation.
- Add ability to CNI-Genie for users to select native network isolation using
  VPC isolation concept from Mizar's design.
- Complete the control plane design to provide reliability and failover through
  distributed hash tables to store pod network groupings & connectivity information.
- Natively offer Network Quality of Service (QoS) to allow users to assign relative
  network traffic priorities to their pods.

We may still bring in some of the items tracked in the [old roadmap](https://github.com/cni-genie/CNI-Genie/blob/master/ROADMAP-old.md) as their relevance and our resource allocation picture becomes clearer in the coming months.

#### 2022 - 2023 Goals

For the next one year, we plan to take a few small steps and accomplish following:

- Identify and on-ramp new additional maintainer(s) for the project.
- Implement basic XDP multihomed pod networking features:
  - Implement pod-to-pod eBPF/XDP based multihomed networking with built-in isolation.
  - Implement service-to-pod eBPF/XDP based multihomed networking with built-in isolation.
- Implement simple and very basic XDP based egress gateway. (stretch goal)
- Ensure ability to configure other CNI providers is retained.
- Restart community engagement for the project.
- Prototype and present the new CNI-Genie roadmap features at conferences and collect feedback.


## CNCF Membership

CNI Genie has mainly benefited from the visibility that comes from inclusion in the CNCF as a sandbox project. This is an area that we hope to get continued support. Once we are ready with initial prototypes of the new approach, we hope to get opportunities to socialize our work and raise awareness of the new developments in this project. This is an area we believe CNCF can really help with.


## Incubation

Currently, CNI Genie remains a sandbox project. It is not ready to apply for incubation due to the limited visibility on how well the new goals may be received. We hope this becomes clearer as we roll out and socialize our first prototypes and we can better measure performance and community interest.
