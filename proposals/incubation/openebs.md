# OpenEBS Incubation Stage Proposal
 
## Table of Contents

- [Background](#background)
- [Incubation Stage Requirements](#incubation-stage-requirements)
- [Recap: What is OpenEBS](#recap-what-is-openebs)

## Background

[OpenEBS](https://openebs.io/) is a leading example of Open Source Container Attached Storage, built using Cloud Native Architecture, and it simplifies running Stateful Applications on Kubernetes. Some of the early adopters of OpenEBS include Arista, Comcast, Optoro and and CNCF itself. There are several others listed on the community [Adopters list](https://github.com/openebs/openebs/blob/master/ADOPTERS.md).  

OpenEBS is currently a Sandox stage project and was approved as a [CNCF Sandbox Project on May 14th, 2019](https://github.com/cncf/toc/pull/224). OpenEBS recently completed its [2020 Annual Review](https://github.com/cncf/toc/pull/433).

OpenEBS is a thriving project with an active community and committed contributors. We believe we are ready to be considered joining Incubation Stage. This proposal focuses on OpenEBS project compliance with the [the Incubation Stage criteria](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage).

## Incubation Stage Requirements

*Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.*

OpenEBS has a large number of successful adopters, some of those have chooses to speak publicly about their usage. We can see them listed in the GitHub [Adopters list](https://github.com/openebs/openebs/blob/master/ADOPTERS.md). Based on the usage trends and discussion in the slack, we are seeing more and more users adopting OpenEBS (not yet in ADOPTERS.md because they can’t make it public yet).

Some of the public references include:
- [Arista Networks](https://github.com/openebs/openebs/blob/master/adopters/arista/README.md) and [Comcast](https://github.com/openebs/openebs/blob/master/adopters/comcast/README.md) using OpenEBS for providing Developer infrastructure by hosting Gerrit, Maven, Chartmueseum and internal tools. 
- [CNCF](https://github.com/openebs/openebs/blob/master/adopters/cncf/README.md) using OpenEBS for hosting DevStats and other tooling.
- [Optoro](https://github.com/openebs/openebs/tree/master/adopters/optoro) runs all of their Stateful workloads on OpenEBS and have been able to migrate to OpenEBS smoothly and completely transparent to their end users.
- [Orange](https://github.com/openebs/openebs/blob/master/adopters/orange/README.md), Cort, [CLEW Medical](https://github.com/openebs/openebs/blob/master/adopters/clewmedical/README.md) and other companies using OpenEBS as the default storage solution for their on-premise Kubernetes clusters


*Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.*

- We have 7 maintainers, 20 reviewers, overall from 5 different companies. The details are here: https://github.com/openebs/openebs/blob/master/MAINTAINERS

- Maintainers are added and removed from the project as per the policies outlined in the project governance: https://github.com/openebs/openebs/blob/master/GOVERNANCE.md

*Demonstrate a substantial ongoing flow of commits and merged contributions.*

As a CNCF project, development activity for OpenEBS can be inspected using [DevStats](https://openebs.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m). Some interesting trends include:
- A sustained monthly [release cadence](https://openebs.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-1y&to=now-1h) since August 2019.
- An increasing monthly [commit rate](https://openebs.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-2y&to=now&var-period=m&var-repogroups=%22openebs%2Fapi%22&var-repogroups=%22openebs%2Fcharts%22&var-repogroups=%22openebs%2Fcstor%22&var-repogroups=%22openebs%2Fcstor-csi%22&var-repogroups=%22openebs%2Fcstor-operators%22&var-repogroups=%22openebs%2Fistgt%22&var-repogroups=%22openebs%2Fjiva%22&var-repogroups=%22openebs%2Fjiva-csi%22&var-repogroups=%22openebs%2Fjiva-operator%22&var-repogroups=%22openebs%2Flibcstor%22&var-repogroups=%22openebs%2Flinux-utils%22&var-repogroups=%22openebs%2Fmaya%22&var-repogroups=%22openebs%2FMayastor%22&var-repogroups=%22openebs%2Fnode-disk-manager%22&var-repogroups=%22openebs%2Fvelero-plugin%22&var-repogroups=%22openebs%2Fwebsite%22&var-repogroups=%22openebs%2Fzfs-localpv%22) (excluding documentation).
- Trending at over 340 [PRs closed](https://openebs.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All) per month.
- Averaging over 5 [new contributors](https://openebs.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now) a month in 2020.
- More than 10 [companies](https://openebs.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions) made over 5 contributions in the last year.
- A total of 34 [companies](https://openebs.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions) contributed in the last year.

*A clear versioning scheme.*

We use the [semantic versioning](http://semver.org/) numbering system. The releases are announced using [github’s releases](https://github.com/openebs/openebs/releases) and followed by slack, mailing list, and social media announcements. 

We release a new minor version every month and a patch version as the demand arises on the last monthly minor release. Typically we seem to reach a cadence of a year for a major release. We are currently moving towards v2.0. 

## Recap: What is OpenEBS
 
OpenEBS is cloud native storage for stateful applications on Kubernetes where "cloud native" means following a loosely coupled architecture for providing data services to Kubernetes stateful workloads. OpenEBS follows a container attached storage (CAS) model.  Every volume that is surfaced to an application is backed by a dedicated controller pod and a set of replica pods. From an implementation perspective: 
- OpenEBS provides a variety of pluggable Storage Engines that are tailored to provide the right storage services to cloud-native workloads. The engines range from orchestrating Local Volumes for Distributed workloads to providing high-available replicated storage for Databases.
- OpenEBS is developed completely in Userspace enabling it to run on any OS / Platform without modification.
- OpenEBS uses the Kubernetes platform itself to provide highly available, horizontally scalable Persistent Volumes using the storage attached to the Kubernetes nodes.

The benefits of using OpenEBS can be summarized based on Kubernetes user persona as follows.

For the Kubernetes SRE / Administrator OpenEBS offers:
- A simple to set up and manage storage solution for Kubernetes clusters that is built on Kubernetes and uses all the standard Kubernetes tooling like Declarative API, kubectl, Prometheus, Grafana and so forth for provisioning, monitoring and Day 2 operations. A familiar approach to maintenance such as install and upgrades, as OpenEBS itself is comprised of containerized microservices and standard extensions to Kubernetes.
- Declarative API using standard Kubernetes primitives (StorageClasses, Custom Resources, ResourceQuotas, Taints/Tolerations and so forth) and Kubernetes Operators enabling administrators to:
   * Set up different classes of Storage services for application workloads.
   * Secure using RBAC, PodSecurityPolicies, OPA, Istio and so forth.
   * Automate storage operations and allow OpenEBS to be installed via platform tools like Helm, BOSH, Terraform, Ansible and Kubernetes bootstrapping projects like [Gravitational](https://gravitational.com/gravity/docs/storage/), [Lokomotive](https://github.com/kinvolk/lokomotive-kubernetes/wiki/The-big-picture#openebs-operator).
- Standard API for Provisioning, Snapshots, Backup and Restore using standard API enabling administrators to implement data protection and migration capabilities by plugging into other projects like Velero.
- Simplified administration and management of the storage attached to the nodes by making block devices first-class Kubernetes resources.
- OpenEBS offers an Open Source alternative to proprietary and non Kubernetes based Storage Products and Services without vendor lock-in for both on-premise and cloud environments.

For Developers using Kubernetes OpenEBS offers:

- Greater control over the storage services delivered to their Stateful Workloads without any shared dependencies due to the use of a shared storage system or service.
- Simple to set up of production-like clusters even when using Stateful workloads on their development machine. This consistency across environments helps to harden applications and improve developer productivity by having access to the entire application deployment stack.
- Light-weight storage engines with ARM support enabling homegrown Kubernetes clusters for running Stateful workloads for developers.

_Note: The OpenEBS project's alignment with Cloud Native and comparison with other Cloud Native projects are covered in details in the [2020 OpenEBS Annual Review]()._

