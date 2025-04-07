# Longhorn Incubation Stage Proposal

Longhorn (https://longhorn.io/) joined CNCF a year ago. As the Longhorn team, we believe we're ready to be considered for joining the Incubation stage of the CNCF project.

## Incubation State Requirements

1. *Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.*

    * Tribunal Regional Eleitoral do Pará (www.tre-pa.jus.br) is the regional electoral court of the state of Pará, Brazil. It is using Longhorn in production as the storage backend for Prometheus, Minio and pgAdmin on-premises.

    * Cerner (https://www.cerner.com) is a health information technology company. It is using Longhorn for persistent storage and highly available data replication for its Rancher-based Cernerworks Enterprise Hosting Observability platform.

    * Tyk (https://tyk.io) is an open-source API and service management platform. With the latest Tyk Cloud product (https://tyk.io/api-gateway/cloud/), users can launch their own deployments of Tyk control planes and API gateway clusters in the locations of their choice. Each Tyk control plane runs an individual persistent 6-node Redis cluster among other things. Longhorn helps Tyk to back hundreds of dynamically provisioned Redis cluster nodes with resilient persistent storage volumes in multiple Kubernetes clusters.

2. *Have a healthy number of committers.*

    We currently have 50 committers across 10 different companies.

3. *Demonstrate a substantial ongoing flow of commits and merged contributions.*

    Commits per week over the last 6 months: **51.26**

    Issues opened per week over the last 6 months: **24.16**

    Issues closed per week over the last 6 months: **18.24**

    New PRs per week over the last 6 months: **29.44**

    All data was gathered thanks to the [CNCF devstats dashboards.](https://longhorn.devstats.cncf.io/)

4. *A clear versioning scheme.*

    The version of Longhorn is based on Sematic Versioning 2.0.0. The version number follows the pattern of MAJOR.MINOR.PATCH. We are planning to do about 3 minor versions per year, and major releases according to necessity. We will do patch releases according to necessity.

## What is Longhorn: Recap

Longhorn is an open-source distributed storage software for Kubernetes. It is a lightweight and portable implementation of persistent storage that works for any Kubernetes cluster. The project implements distributed block device volumes that can be mounted as read-write by a single node (ReadWriteOnce).

Longhorn is designed to leverage existing Linux technologies as much as possible, rather than building a complex storage technology stack from scratch. The software takes advantage of modern high-speed and high capacity SSD and NVMe storage, and proven Linux storage features like sparse files and cgroups.

The most distinct feature of Longhorn is to implement each volume as an independent microservice. By leveraging Kubernetes to orchestrate the volumes, it implements a highly feature-rich enterprise-grade distributed block storage system with only tens of thousands of lines of Go code.

### Architecture

Longhorn consists of the following major components:

***Engine and Replicas***

The Longhorn engine is the microservice used to implement each volume. The engine implements the data plane of reading and writing the block devices. Each Longhorn volume is implemented using one engine and multiple replicas. The engine always runs on the same node as the pod consuming the volume, whereas Longhorn replicas reside on different nodes to ensure redundancy.

Longhorn delivers added resiliency because the data plane for each volume is separate. A software bug that causes one volume to malfunction does not impact other volumes in the system.

***Manager***

The Longhorn manager is built on top of Kubernetes. It implements the CSI driver for Longhorn and leverages Kubernetes to ensure the proper functioning of all engines and replicas. The manager relies on the Kubernetes controller pattern to manage the volume status. It also utilizes Kubernetes to implement pod orchestration, cron jobs, node failure detection, pod failure detection, log collection, etc.

***Other Components***

Besides engine and manager, Longhorn also has an intuitive GUI and a test suite that covers every feature of Longhorn. Integration test cases are required for new features and major bug fixes.

## Statement on alignment with the CNCF mission

The Longhorn project aligns with CNCF's mission to make cloud-native computing ubiquitous. Today, persistent storage support varies from Kubernetes cluster to Kubernetes cluster. As a result, users cannot deploy stateful workloads that require persistent volumes and expect their stateful workloads to work on any Kubernetes cluster. By creating a portable and lightweight implementation of persistent storage volumes, Longhorn can make persistent storage support available on every Kubernetes cluster and thus can help drive the adoption of Kubernetes technology.

## Comparison to existing CNCF storage projects

The goal of Longhorn is to provide reliable and portable persistence storage support for every Kubernetes cluster.

To reach this goal, we not only need to make sure Longhorn has the best user experience, but we also need to make sure the product is easy to understand so it can be easily maintained.

### Comparison to Rook

Rook is a storage orchestrator for Kubernetes. It transformed the existing storage solutions, e.g. Ceph, into a Cloud Native application.

Longhorn is a full-stack storage software, which is built with Cloud Native in mind. Longhorn's position is more similar to Rook + Ceph + a third-party backup solution.

### Comparison to OpenEBS

OpenEBS is a full-stack storage software. It has a list of swappable engines, including Jiva, cStor, local storage, and the new Mayastor. Jiva is the first GA engine available on OpenEBS, and it's a fork from the Longhorn engine.

To simplify the user experience, Longhorn focuses efforts on the Longhorn engine only, which results in better integration between the data layers and the management layer.

Also, Longhorn supports off-site backup and disaster recovery, which are not supported by OpenEBS.

Longhorn also provides a built-in UI to enhance the user experience.

## Project Progress since Sandbox

Longhorn has joined the CNCF in October 2019. Since then, we've seen the tremendous growth of the Longhorn community.

Longhorn has reached v1.0 GA this May, and we're planning to release v1.1 this November.

### Community activity

We're hosting community meeting and office hour every month. The community meeting recordings can be found at https://www.youtube.com/channel/UCGk1f-LCVmccf1pX4OvF1Ig .

For the last year, we've participated in KubeCon NA 2019, Cloud-Native + Open-Source Summit China 2020, as well as KubeCon EU 2020.

We're also going to host a booth and office hours in KubeCon NA 2020.

### Community growth

We've seen tremendous growth since we joined CNCF one year ago. Compare October 2019 to October 2020:

* GitHub stars: 600+ vs 2k+ (3x)
* Slack members: 200+ vs 900+ (4x)
* Known active node count: 2.7k vs 14k+ (5x)

![Longhorn community growth](https://rancher-longhorn.s3-us-west-1.amazonaws.com/longhorn-community-growth-2020.png)

## Features & Roadmap

### Current latest release (v1.0.2)
* Enterprise-grade distributed storage software
* Incremental snapshot of the block storage
* Backup to an off-cluster backup store (NFS or S3-compatible object storage), built on efficient change block detection mechanism
* Volume expansion
* Cross-cluster disaster recovery volume with defined Recovery Time Objective (RTO) and Recovery Point Objective (RPO)
* Cross-availability-zone replica scheduling
* Automated non-disruptive upgrade. You can upgrade the entire Longhorn software stack without disrupting running volumes.
* Intuitive GUI

### Upcoming release (v1.1.0)
* RWX support
* Prometheus support
* CSI Snapshotter support
* Data locality
* ARM64 support

### Future plans
* Longhorn CLI
* SPDK support
* Application backup and restore
