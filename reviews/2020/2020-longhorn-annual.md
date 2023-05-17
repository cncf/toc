# 2020 Longhorn Annual Review

## Background

[Longhorn](https://longhorn.io) is an open-source distributed storage software for Kubernetes. It is a lightweight and portable implementation of persistent storage that works for any Kubernetes cluster. The project implements distributed block device volumes that can be mounted as read-write by a single node (ReadWriteOnce).

Longhorn is designed to leverage existing Linux technologies as much as possible, rather than building a complex storage technology stack from scratch. The software takes advantage of modern high-speed and high capacity SSD and NVMe storage, and proven Linux storage features like sparse files and cgroups.

The most distinct feature of Longhorn is to implement each volume as an independent microservice. By leveraging Kubernetes to orchestrate the volumes, it implements a highly feature-rich enterprise-grade distributed block storage system with only tens of thousands of lines of Go code.

**Longhorn was accepted into CNCF Sandbox on Oct 11, 2019**
- [Longhorn Sandbox Proposal](https://github.com/cncf/toc/pull/259)
- [TOC proposal slide](https://docs.google.com/presentation/d/1_Fr8DPaKNb-wyM2GQqWC_PtcuatScjnbGFGvySHyAIc/edit?usp=sharing)

## Annual Review Questions

* Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

  * Commits per week over the last 6 months: **51.26**

  * Issues opened per week over the last 6 months: **24.16**

  * Issues closed per week over the last 6 months: **18.24**

  * New PRs per week over the last 6 months: **29.44**

  * [CNCF DevStats](https://longhorn.devstats.cncf.io)

* How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

  * There are 2 maintainers in total that work for Rancher Labs.

* What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

  The end users include:

  * Tribunal Regional Eleitoral do Pará (www.tre-pa.jus.br) is the regional electoral court of the state of Pará, Brazil. It is using Longhorn in production as the storage backend for Prometheus, Minio and pgAdmin on-premises.

  * Cerner (https://www.cerner.com) is a health information technology company. It is using Longhorn for persistent storage and highly available data replication for its Rancher-based Cernerworks Enterprise Hosting Observability platform.

  * Tyk (https://tyk.io) is an open-source API and service management platform. With the latest Tyk Cloud product (https://tyk.io/api-gateway/cloud/), users can launch their own deployments of Tyk control planes and API gateway clusters in the locations of their choice. Each Tyk control plane runs an individual persistent 6-node Redis cluster among other things. Longhorn helps Tyk to back hundreds of dynamically provisioned Redis cluster nodes with resilient persistent storage volumes in multiple Kubernetes clusters.

  Also, we've seen tremendous growth of the community in general since we joined CNCF one year ago. Compare October 2019 to November 2020:

  * GitHub stars: 600+ vs 2k+ (3x)

  * Slack members: 200+ vs 1k+ (4x)

  * Known active node count: 2.7k vs 14k+ (5x)

* How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

  Longhorn has perform very well compared to one years ago. Besides the above metrics we saw, we've also achieved the follow milestones:

  * First beta release v0.7.0 on Nov 2019.

  * First GA release v1.0.0 on May 2020.

  * Participated in KubeCon NA 2019, Cloud-Native + Open-Source Summit China 2020, as well as KubeCon EU 2020.

* What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

  At the moment, the main focus of the project is feature implementation. In the upcoming v1.1.0 release, we've implemented a series of new features to enhance user experience for the node/cluster maintenance. We've also implemented RWX as well as ARM64 support.

  We're looking into ways to further improve the performance and stability of the product.

* How can the CNCF help you achieve your upcoming goals?

  CNCF has helped a lot in the past year. From setting up the framework for the Longhorn website, to providing the infrastructure support to the Longhorn test environment.

  We do wish to get more help from marketing perspective, but it's understandable that the graduated and incubation projects have the priority.

* Do you think that your project meets the criteria for incubation?

  Yes, we've already [submitted a proposal](https://github.com/cncf/toc/pull/553) for incubation.
