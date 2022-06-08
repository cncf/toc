# OpenEBS Annual Review 2020

## Table of Contents

- [Background](#background)
  * [Notable improvements in the last year](#notable-improvements-in-the-last-year-include)
- [Alignment with Cloud Native](#alignment-with-cloud-native)  
  * [Comparison with Longhorn](#comparison-with-longhorn)
  * [Comparison with Rook](#comparison-with-rook)
- [Annual review Questions](#annual-review-questions)

## Background

[OpenEBS](https://openebs.io/) is a leading example of Open Source Container Attached Storage, built using Cloud Native Architecture, and it simplifies running Stateful Applications on Kubernetes. OpenEBS was approved as a [CNCF Sandbox Project on May 14th, 2019](https://github.com/cncf/toc/pull/224).  Some of the early adopters of OpenEBS include Arista, Comcast, and the CNCF and others listed on the community Adopters list.

- [OpenEBS Sandbox Proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/openebs.adoc)
- [Original CNCF TOC Meeting Slides](https://docs.google.com/presentation/d/10Z5SsA90MyKLgSSvmGAyufsxoAABIfLTaHDwG7oNQB0/edit#slide=id.g59b9c9aad1_0_245)
- [CNCF Storage WG Slides](https://docs.google.com/presentation/d/1MEx9EyvFW1cLZhFxUvWVd4bzoRS7PNnMqgmbIZaoulw/edit#slide=id.g34883d674a_0_601)


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

From an implementation perspective:

- OpenEBS provides a variety of pluggable Storage Engines that are tailored to provide the right storage services to cloud-native workloads. The engines range from orchestrating Local Volumes for Distributed workloads to providing high-available replicated storage for Databases.
- OpenEBS is developed completely in Userspace enabling it to run on any OS / Platform without modification.
- OpenEBS uses the Kubernetes platform itself to provide highly available, horizontally scalable Persistent Volumes using the storage attached to the Kubernetes nodes.

### Notable improvements in the last year include:

- Growth in the user and contributor community since OpenEBS became a Sandbox project. 
  * OpenEBS was the second most likely storage solution to be evaluated, out of a broad mix of projects and traditional storage systems and services. Based on [CNCF September 2019 survey](https://www.cncf.io/wp-content/uploads/2020/03/CNCF_Survey_Report.pdf)
  * Registered slack users increased from 1400 to 2000+.
- Added new storage engine - ZFS Local PV (currently in beta) to address the Day 2 operational requirements on Local PV like volume expansion, capacity quota enforcement, snapshot and clone for running.
- Added a new storage engine - MayaStor (expected to achieve beta in June) to address the performance requirements of Stateful workloads running within Kubernetes. The goal of the MayaStor engine is to provide performance as close as Local PV and also provide replication capabilities. MayaStor leverages SPDK to enable the use of NVMe by Kubernetes environments.
- Support for ARM64 builds.
- Reduce the capacity overhead incurred with running Jiva volumes for a longer duration. 
- Enhance cStor API and simply Day 2 operational tasks like replacing Kubernetes nodes, Volume, and storage expansion, handling the corner cases around storage pools hitting maximum usage when overprovisioning is enabled.
- Enhance Dynamic Local PV provisioners for working with complex Kubernetes cluster designs that include: the same hostname being used in different zones, hostname and node name being different, working with nodes having taints.
- Simplification of upgrades by automating the steps via Kubernetes jobs. 
- Maintaining monthly release cadence with each release focusing on stability, enhancements and new alpha features.
- Running security vulnerability scanner (trivy) on all the delivered OpenEBS container images and fixing the critical vulnerabilities prior to release.

In addition, the OpenEBS contributor community has been involved in improvements to some of the other projects that OpenEBS depends on, such as:
- Converting the E2e test framework used by OpenEBS called [Litmus](https://github.com/litmuschaos/litmus) as another project that can be used by any other Kubernetes based products. A proposal to submit [Litmus as a CNCF Sandbox project](https://github.com/cncf/toc/pull/391) is up for review.
- Fixing issues and helping with migration tasks for projects maintained in other Github org such as:
  * https://github.com/kubernetes-csi/
  * https://github.com/kubernetes-sigs
  * https://github.com/spdk
  * https://github.com/grpc
  * https://github.com/vmware-tanzu/velero
  * https://github.com/longhorn/
  * https://github.com/gostor/gotgt

## Alignment with Cloud Native

OpenEBS falls within the scope of they [CNCF Storage SIG](https://github.com/cncf/sig-storage/blob/master/storage-charter.md#areas-considered-in-scope), offering a significantly different approach to implementing a block storage system by making use of cloud-native principles and Kubernetes itself as a key component to delivering storage for application workloads.

OpenEBS and its community are aligned with the CNCF mission.

- Extensible and open architecture based upon Kubernetes. Built with Kubernetes, for Kubernetes. 
- Horizontally scalable with Kubernetes nodes allowing deployments from small developer-centric to large scale enterprise environments. 
- Works well and users have deployed OpenEBS alongside other CNCF projects like Kubernetes, Prometheus, etcd, and others
- Additionally, the OpenEBS community attempts to be extremely inclusive and is one of the more active CNCF related projects in India, actively working with technical universities to introduce interested students to Open Source practices and projects.

### Comparison with Longhorn
Longhorn is another CNCF project. One of the OpenEBS Storage engines, Jiva, was forked out of Longhorn-engine and the OpenEBS community has kept the engine fixes as compatible as possible. 

OpenEBS Jiva engine differs from Longhorn:

- Deployment and control plane operations. OpenEBS Jiva engine continues to provide a complete isolated set of containers for each volume and is designed to self-heal from cluster failures or other chaos. Longhorn currently depends on the control plane pods for managing and recovering from failures.  
- Jiva remains focused on providing volumes for workloads that are short-lived or have low capacity requirements. This allows Jiva to not require inbuilt support for scheduled snapshots and restore. However, Jiva does support the API for administrators to take remote backups for handling complete cluster migration. 

In addition to Jiva, other storage engines included as a part of OpenEBS include:

- LocalPV - including hostpath, device, and ZFS
- cStor - including an efficient COW engine for snapshots and clones
- MayaStor - a soon to be beta engine refactored in Rust and intended to be a high performance engine, with native use of NVMe
  
### Comparison with Rook

Rook is another CNCF project that focuses on:

- Simplifying the management of Storage engines that are developed outside of the CNCF like Ceph and EdgeFS. 
- Provide a framework to develop higher-level abstractions and operators for running stateful workloads like CockroachDB, Cassandra. 

OpenEBS is focused on Storage engines and the standard Kubernetes abstractions to use them.  

One of the ways in which users have benefited from the Rook and OpenEBS projects is to deploy the Stateful Workload using the Rook abstractions with OpenEBS storage classes.  

Areas of future collaboration with the Rook community may include:
- NFS provisioner
- Application definitions
- Further improvements to Kubernetes capabilities such as Block Device Management, Application Backup, Restore and Migration

## Annual review Questions

- *Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.*

  [OpenEBS DevStats](https://openebs.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-1y&to=now-1h).
  
  Some  interesting trends:
  - [Monthly release cadence since Aug 2019](https://openebs.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-1y&to=now-1h).
  - [Increase in average monthly commits (excluding docs)](https://openebs.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-2y&to=now&var-period=m&var-repogroups=%22openebs%2Fapi%22&var-repogroups=%22openebs%2Fcharts%22&var-repogroups=%22openebs%2Fcstor%22&var-repogroups=%22openebs%2Fcstor-csi%22&var-repogroups=%22openebs%2Fcstor-operators%22&var-repogroups=%22openebs%2Fistgt%22&var-repogroups=%22openebs%2Fjiva%22&var-repogroups=%22openebs%2Fjiva-csi%22&var-repogroups=%22openebs%2Fjiva-operator%22&var-repogroups=%22openebs%2Flibcstor%22&var-repogroups=%22openebs%2Flinux-utils%22&var-repogroups=%22openebs%2Fmaya%22&var-repogroups=%22openebs%2FMayastor%22&var-repogroups=%22openebs%2Fnode-disk-manager%22&var-repogroups=%22openebs%2Fvelero-plugin%22&var-repogroups=%22openebs%2Fwebsite%22&var-repogroups=%22openebs%2Fzfs-localpv%22).
  - [Tending towards 8/10 PRs closed per week](https://openebs.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=d7&var-repogroup_name=All). 
  - [2 to 3 new contributors per month. Oct/Nov seeing the highest due to Hacktoberfest](https://openebs.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now).
  - [10+ companies with 5+ contributions, and a total of 29 companies contributing in the last year](https://openebs.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions).  
  
 - *How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)*
   
   OpenEBS source code, documentations and tests are distributed among several repositories. 
   - There are a total of [7 maintainers](https://github.com/openebs/openebs/blob/master/MAINTAINERS) and
   - 17 additional reviewers. The [reviewers](https://github.com/openebs/openebs/blob/master/GOVERNANCE.md#reviewers) maintain one or more repositories and are on their way to becoming maintainers.
   - Looking at the [devstats on companies contributing since joining CNCF](https://openebs.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Since%20joining%20CNCF&var-metric=contributions), we have seen several companies contributing to the project in terms of issues, reviews on PRs and code fixes. MayaData and Sanmina/Viking are already part of the OpenEBS org. We are in the process of onboarding regular contributors from companies like Kinvolk, Orange, IBM, DO, Huawei and several independent contributors.
  
- *What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end-users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)* 

  In the past year since being accepted into Sandbox, we have seen companies coming forward to be listed in the [public adopters](https://github.com/openebs/openebs/blob/master/ADOPTERS.md) in the repo. Based on the usage trends and discussion in the slack, we are seeing more and more users adopting OpenEBS (not yet in ADOPTERS.md because they can’t make it public yet).
  
  Some of the public references include:
  - Arista Networks and Comcast using OpenEBS for providing Developer infrastructure by hosting Gerrit, Maven, Chartmueseum and internal tools. 
  - CNCF using OpenEBS for hosting DevStats and other tooling.
  - Orange, Cort, CLEW Medical and other companies using OpenEBS as the default storage solution for their on-premise Kubernetes clusters. 

- *How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)*

  Though not formally stated, for the first year some goals we had were:
  - Grow the user community by addressing the adoption blockers and increasing awareness about OpenEBS. From the weekly trends in docker pulls, slack interactions, GitHub issues/comments, external publications, we think we have done well. 
  - Improve the awareness that Kubernetes can provide a substrate for functions such as storage and data management.  We hope to better cooperate with fellow CNCF projects such as Vitess and Longhorn in the year to come.
  - Address performance requirements. We looked at the applications that need high performance. Many of these applications are distributed in nature and are resilient against node failures and expect to access disk or cloud volumes directly.  In response, the OpenEBS community has made three flavors of Local PV that can be used to meet the application requirements and the type of storage attached to the nodes. For the non-distributed applications or those that would like built-in migration and encryption support MayaStor development is close to reaching beta and will be able to perform at close to LocalPV speeds.
  - GA of cStor storage engine. We set ourselves the goal of calling the storage engines as GA after the custom resource specs have been marked as v1 and all the positive and negative tests have been automated. We are not quite done yet, though we have made [significant progress](https://github.com/orgs/openebs/projects/9) in this regard by increasing the coverage in automated E2e tests and publishing them [online](https://openebs.ci).  One of the design changes we made along the way was to remove the dependency on a centralized openebs API server for cStor volume management and create its own Operators.
  
- *What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?*
  - Continue to expand our community of user and contributors. We are particularly focused on improving the contributor experience.  We believe that as the project matures and usage continues to grow, we will see an increase in interested contributors.  We plan to learn from the Contributor SIG how to better embrace them and help them to become successful contributors.  While we do have some major external contributors, we know there is more we can do.
  - Improve the documentation and website.
  - Work on the [Roadmap items](https://github.com/openebs/openebs/blob/master/ROADMAP.md), which primarily include:
    * GA of cStor and migration towards CSI 
    * GA of Local PV Operators
    * Beta and GA of Mayastor 

- *How can the CNCF help you achieve your upcoming goals?*

  First off, thank you for all the help and support provided by CNCF team in the past year. Some of the areas where we could use additional help are:
  - We need a lot of help from technical writers to improve the documentation and content published on the website.
  - The work of the Contributor SIG is immediately relevant.  
  - Embracing and communicating the use of Kubernetes in ways that may not have been originally anticipated or intended by the originators of Kubernetes, including as a substrate for a fundamentally new way of managing data and delivering storage capabilities.  For example, it may be useful to increasingly use a phrase such as Container Attached Storage to reflect this pattern.
  - Additional help with reviewing licenses used by sub components.  There may need to be a check if the licenses of these subcomponents - most notably CDDL - can be included in the CNCF whitelist policy or if these sub components should be held in another open source repository such as an OpenSource repository of MayaData or otherwise.
  
- *Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?*
  
  Yes. We are in the process of completing the Due Diligence document and submitting the incubation proposal. 

