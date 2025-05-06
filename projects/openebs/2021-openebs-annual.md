# OpenEBS Annual Review 2021


## Background

[OpenEBS](https://openebs.io/) is a leading example of Open Source Container Attached Storage, built using Cloud Native Architecture, and it simplifies running Stateful Applications on Kubernetes. OpenEBS manages the local storage available on the Kubernetes worker nodes and dynamically creates distributed persistent block volumes. OpenEBS is implemented using microservices pattern, the storage management and volumes are provided using containers - that are managed by Kubernetes itself. 

Reading through the [OpenEBS adoption stories](https://github.com/openebs/openebs/blob/master/ADOPTERS.md), the primary reasons for using OpenEBS are: 
* OpenEBS provides a collection of data engines that are tailored for the type of Kubernetes environments that users have and the needs of the Stateful workloads. For example, OpenEBS can be used on nodes that don't have any disks to those that have NVMe devices. OpenEBS can be used to provide Local Volumes for Distributed Stateful Workloads and provide Highly Available volumes for workloads that need protection against node failures. 
* OpenEBS increases the productivity of the application teams - by allowing them to work with platform teams to provide a cluster of nodes with storage devices and application teams can run Kubernetes and optimally manage the compute, memory, network and storage available to them. 
* OpenEBS reduces the friction between the application and platform teams. Platform teams can setup a Cloud Volume like services on premises for their development teams. Platform teams that have adopted Cloud for running their Kubernetes services also benefit by using a fixed set of cloud volumes attached to the nodes and let the application teams operate within those boundaries, instead of providing access to developers to create (and forget) cloud volumes.
* OpenEBS improves the availability of the storage using its synchronous replication for distributing the volumes. The volume data is always fully available on the replica nodes - there by reducing the overhead incurred in maintaining the metadata of the volume. Helps with easier rebuild and lower blast radius when compared to scale-out distributed volumes. OpenEBS also has been used to replicate volumes across availability zones on cloud where such features are not yet available. 

Last but not the least, OpenEBS community support has also been spotted as one of the main reasons for using OpenEBS compared to other storage solutions. 

OpenEBS was approved as a [CNCF Sandbox Project on May 14th, 2019](https://github.com/cncf/toc/pull/224). Here are some background links:

- [OpenEBS Sandbox Proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/openebs.adoc)
- [Original CNCF TOC Meeting Slides](https://docs.google.com/presentation/d/10Z5SsA90MyKLgSSvmGAyufsxoAABIfLTaHDwG7oNQB0/edit#slide=id.g59b9c9aad1_0_245)
- [CNCF Storage WG Slides](https://docs.google.com/presentation/d/1MEx9EyvFW1cLZhFxUvWVd4bzoRS7PNnMqgmbIZaoulw/edit#slide=id.g34883d674a_0_601)
- [OpenEBS 2020 Annual Review PR](https://github.com/cncf/toc/pull/433)
- [OpenEBS Incubation Proposal PR](https://github.com/cncf/toc/pull/506) 


## Develoment metrics

Here are a few highlights from the project's [devstats page](https://openebs.devstats.cncf.io/), looking at the 1 year period after
2020 annual review (Jul 2020 - Jun 2021):

 - This year we welcomed [101 new contributors](https://openebs.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1593455400000&to=1625077799000) across our repositories.
 - The average number of [active contributors](https://openebs.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=1593541800000&to=1625077799000) increased to 77, up from 58.
 - These active contributors come from 15 different [companies](https://openebs.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&from=1593541800000&to=1625077799000), up from 10.
 - The [activity](https://openebs.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&from=1561919400000&to=now&var-period=m&var-repogroups=All) across repositories shows a clear steady increase.
   
### Maintainers

- There are a total of [10 maintainers](https://github.com/openebs/openebs/blob/master/MAINTAINERS), up from 7 last year and
- 29 additional reviewers. The [reviewers](https://github.com/openebs/openebs/blob/master/GOVERNANCE.md#reviewers) maintain one or more repositories and are on their way to becoming organization maintainers. This is up grom 17 last year. 

Majority of the maintainers still are from MayaData, while ~10 maintainers are contributing from other companies. 



### Other metrics

Besides contributor metrics, the following numbers also highlight how the
project has gathered more interest since last year:

- [~7k github stars](https://star-history.t9t.io/#openebs/openebs), up from 6k since last year.
- Community Slack was migrated from MayaData hosted slack to Kubernetes Slack:
   - The general [#openebs](https://kubernetes.slack.com/messages/openebs/) channel now has 650+ members.
   - The developer [#openebs-dev](https://kubernetes.slack.com/messages/openebs-dev/) currently has ~200 members.
   
## Adoption

There are many more adopters of OpenEBS and organizations evaluating the project as well. Organizations that agreed to have them listed publicly can be found in the [ADOPTERS.md](https://github.com/openebs/openebs/blob/master/ADOPTERS.md) file which currently lists 25+ adopters, doubled since last year.


## Project goals

**How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)**

Our goals from last year annual review were:

- *Continue to expand our community of user and contributors. We are particularly focused on improving the contributor experience.*
  
  We have made some progress in this regard. However we still need to get a lot better in terms of onboarding new contributors, by making it easier to contribute via running contributor sessions, regular contributor meetings, additional tooling and refactoring of the repositories to let contributors focus on specific areas for their contribution. 

- *Improve the documentation and website.*

  We have made progress in engaging new contributors to help with improving the documentation and website. New content is being worked on to be part of the upcoming 3.0 release. We also kick-started a process to capture the feedback that comes through the community slack into [GitHub Community tracker projects](https://github.com/orgs/openebs/projects) to fix the website and documentation aspects. 
  

- *Work on the [Roadmap items](https://github.com/openebs/openebs/blob/master/ROADMAP.md), which primarily include:*

  Our primary goal from a feature readiness standpoint was to hit GA with the OpenEBS data engines. There has been a lot of progress made in this regard as can be seen from the updates done to the [Roadmap items](https://github.com/openebs/openebs/blob/master/ROADMAP.md). In addition we had to work on a few additional items to address user requests, deal with the changes in the eco-system around Travis/DockerHub and improving the contributor experience. Few noteworthy changes are as follows: 
  - Migrate Jiva and cStor from external-storage provisioner to CSI driver based provisioning 
  - Enhance Jiva to make use of OpenEBS Local Volumes instead of `hostpath` volumes
  - Enhance Jiva and cStor volumes to deal with Read-only mounts and multi-attach errors that happen due a abrupt node reboots 
  - Support for multi-arch images - with native support for AMD64 and ARM64
  - Enhance Local PV volumes to be backed by LVM and Device partitions in addition to hostpath and ZFS from earlier releases
  - Enhance OpenEBS velero plugin to perform ZFS Local PV incremental backups 
  - Enhance Mayastor to handle Kubernetes node failure scenarios. Working closely with the upstream SPDK community to push NVMe and other fixes 
  - Migrating all the CI pipelines to GitHub Actions
  - Refactor and modularize our system (e2e) tests so that they can be run against a data engine. Helps to release each engines independently
  - Refactor and modularize the helm charts to release each data engine independently to allow users to install only the required components
  - Split most of the mono repos into functionality based repos for better manageability and contributor experience
   
**What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?**

Our goals for the next year will be focused on:
  - Continue to progress on hitting the GA for all the data engines by listening to users and working on the [planned items](https://github.com/openebs/openebs/blob/master/ROADMAP.md) 
  - Focus more on community, reference architectures and solution documents that can increase the adoption of Kubernetes for running Stateful workloads. 
  - Follow through with the feedback being received from the CNCF Storage TAG to help with Incubation of the project


## CNCF membership

### How can the CNCF help you achieve your upcoming goals?

First off, thank you for all the help and support provided by CNCF team in the past year in the following areas for the OpenEBS project:
* Guidance on the licensing related questions from CNCF legal and coaching from the CNCF Storage TAG on how to address some of the concerns. 
* Providing venues during KubeCon for presenting Project Office hours 
* CNCF Mentorship and Bugbash programs that have helped with increasing the contributions to the project. 
* Helping on queries related to process and governance. 

We are looking forward to the continued support in the above areas and as well as get ready to taken help in terms of additional tooling that CNCF provides around security auditing, CI infrastructure and creative services. 


### Incubation

**Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?**

Yes, we do, and we are planning to present the progress we have made since our last [CNCF Storage TAG review](https://github.com/cncf/toc/pull/506#issuecomment-756153457) and move towards the next steps.

