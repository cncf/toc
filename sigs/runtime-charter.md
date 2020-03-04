## CNCF Runtime SIG Charter


### Primary Author: Quinton Hoole

2019-09-02

Also reviewed and contributed to by:

*   Liz Rice
*   Brian Grant

## Introduction

This is the charter referred to in “[CNCF
SIGs](https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#sig-charter)”
by the CNCF TOC, and consistent with the [proposed SIG
definition](https://github.com/cncf/toc/blob/master/sigs/proposed.md).


## Areas Considered In Scope

Workload execution and management systems, components and interfaces
used in [modern cloud-native
environments](https://github.com/cncf/toc/blob/master/DEFINITION.md)
are in scope, including:



*   generalized orchestration, autoscaling, scheduling, execution, container
    runtimes, sandboxing, virtualization, image packaging/distribution and specialized architectures thereof, e.g. 
    *   those aimed specifically at Edge, IoT, Batch, Big Data, AI/ML, etc
    *   those incorporating specialized computing elements beyond CPUs, including GPUs, TPUs, FPGAs, ASICs, etc.

We strive to understand the fundamental characteristics of different approaches with respect to availability,
scalability, performance, consistency, ease-of-use, cost and operational complexity; and relate these to their
suitability to various cloud-native use cases.


## Areas Considered Out Of Scope

Anything not considered in scope above is out of scope.  See also “Interfaces with Related Groups” below.

Examples include:

*   General authentication, authorization, accounting, auditing, etc (even though these clearly apply to several of the areas in the scope of this SIG)
    - because AAA, etc is clearly the domain of the [CNCF Security SIG](https://github.com/cncf/sig-security).
*   Getting network traffic into and out of workloads, or any of the interfaces related to that (e.g. CNI)
    - these are the domain of the ([proposed](https://github.com/cncf/toc/blob/master/sigs/proposed.md)) CNCF Traffic SIG, and other related network-centric groups.  
*   Local or remote storage associated with workloads
     - this is the domain of the [CNCF Storage SIG](https://github.com/cncf/sig-storage) and related groups.


## SIG Mission Statement

To enable widespread and successful execution of the full spectrum of workload types,
including both general latency-sensitive and batch, as well as more specialized
categories listed as in scope, in cloud-native environments through:



1. Providing valuable and unbiased information to the TOC,
   End Users and Projects of the CNCF regarding areas considered in scope (see above).
2. Collaborating effectively with other related groups (see below).
3. Helping to maintain the continued health of the CNCF Projects deemed
   to be within the scope of this SIG (see below)
4. Identifying and filling gaps in the landscape of CNCF Projects within scope.

Specific SIG deliverables are as per the above, and the [general SIG responsibilities
set out by the CNCF TOC](https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#responsibilities--empowerment-of-sigs).


## Current CNCF Projects considered to be within the Scope of this SIG



1. Kubernetes
2. Containerd
3. Harbor
4. Dragonfly
5. Virtual Kubelet
6. CRI-O
7. KubeEdge
8. KubeVirt


## Interfaces With Other Related Groups



*   **[Several Kubernetes SIGs](https://github.com/kubernetes/community)**
    cover Kubernetes-specific workload, scheduling, autoscaling, execution
    and other related abstractions, interfaces, and implementations of
    these interfaces.  We will maintain communication with these Kubernetes
    SIGs where needed.  Our aim is to avoid unnecessary duplication of
    effort by the two groups and maintain clear and consistent messaging
    to our end user community and projects.
*   **[CNCF Security SIG](https://github.com/cncf/sig-security)**
    - works on the more general area of cloud-native security including
    authentication, authorization, encryption, accounting, auditing, and
    related topics.  We defer as much as possible to this group to deal
    with general security-related issues and liaise closely with them on
    how to deal with security areas where these arise.
*   **[CNCF App Delivery SIG](https://github.com/cncf/sig-app-delivery)**
    is focussed on the development, deployment, operation and testing of
    cloud-native applications.  We collaborate with this SIG where it
    pertains to helping to ensure that the required underlying workload
    execution abstractions and mechanisms are suitably provided to support
    these application-level delivery needs.
*   **[OCI Open Container Initiative](https://www.opencontainers.org/)**
    is an open governance structure focussed on creating industry standards
    around container formats and runtime. CNCF projects in the scope of
    this SIG-Runtime will typically engage directly with OCI, with SIG-Runtime
    being involved as needed.
*   **Note regarding Container Registries/Repositories** and the like:
    While image packaging and distribution (and hence container
    registries/repositories in general) fall within the scope of this
    Runtime SIG, many of their common features and use cases are better
    dealt with by other CNCF SIGS.  For example:
    *   image storage, caching, etc - Storage SIG
    *   Image encryption, signing etc - Security SIG
    *   use of image registries to store and distribute many other types
        of artifacts, and in particular the format of these artifacts,
	including helm charts, OPA policies, public security certificates,
	data sets, machine learning models, etc, etc - the SIG relevant to
	those artifact types, e.g. Apps SIG, Security SIG, etc.


## Operating Model

This SIG follows the [standard operating guidelines](https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#operating-model)
provided by the TOC unless otherwise stated here.

**Current TOC Liaison:**  Brendan Burns and Brian Grant

**Co-Chairs:** [Quinton Hoole](https://www.linkedin.com/in/quintonhoole/),
  [Ricardo Aravena](https://www.linkedin.com/in/raravena),
  [Diane Feddema](www.linkedin.com/in/dianefeddema)

**Tech Leads:** [Klaus Ma](http://www.klaus1982.cn/about/), 2 TBD

**Other named roles: **None at present; will be identified and staffed as needed.


## Meeting Schedule

The CNCF SIG-Runtime group meets twice a month on the 1st and 3rd Thu of
each month at 8am Pacific, 11am Eastern. 

Zoom: [https://zoom.us/my/cncfsigruntime](https://zoom.us/my/cncfsigruntime)

Mailing list: Join SIG-Runtime mailing list at [lists.cncf.io](https://lists.cncf.io)

Slack channel: [https://cloud-native.slack.com/messages/CPBE97SMU](https://cloud-native.slack.com/messages/CPBE97SMU)  
