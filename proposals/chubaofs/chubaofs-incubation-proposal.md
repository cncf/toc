# Propose ChubaoFS to Incubation Stage 

## Table of Contents

- [Background](#background)
- [Recap: What is ChubaoFS](#recap-what-is-chubaofs)
- [Progress Since Sandbox](#progress-since-sandbox)
- [Incubation Stage Requirements](#incubation-stage-requirements)
- [Statement on Alignment with the CNCF Mission](#statement-on-alignment-with-the-cncf-mission)
- [Comparison with Other CNCF Storage Projects](#comparison-with-other-cncf-storage-projects)
- [Future Plans](#future-plans)

## Background

ChubaoFS (https://chubao.io) was first [presented to the CNCF storage SIG](https://docs.google.com/document/d/15Cnz3AHooqtml0xWhF6Hgei4gpWiRJbwoVCO_pN2z4c/edit?usp=sharing) in June 2019, and was accepted as a sandbox project in December 2019.

## Recap: What is ChubaoFS

ChubaoFS is a distributed storage system aims to provide storage solutions for containerized applications deployed in a large scale container platform. 

### High-level Architecture

ChubaoFS consists of a metadata subsystem, a data subsystem, and a resource manager, and can be accessed by different clients (as a set of application processes) hosted on the containers through different file system instances called volumes.

The metadata subsystem stores the file metadata, and consists of a set of meta nodes. Each meta node consists of a set of meta partitions.

The data subsystem stores the file contents, and consists of a set of data nodes. Each data node consists of a set of data partitions.

The volume is a logical concept in ChubaoFS and consists of one or multiple meta partitions and one or multiple data partitions. Each partition can only be assigned to a single volume. From a client’s perspective, the volume can be viewed as a file system instance that contains data accessible by the containers. A volume can be mounted to multiple containers so that files can be shared among different clients simultaneously, and needs to be created at the very beginning before the any file operation. A ChubaoFS cluster deployed at each data center can have hundreds of thousands of volumes, whose data sizes vary from a few gigabytes to several terabytes.

Generally speaking, the resource manager periodically communicates with the metadata subsystem and data subsystem to manage the meta nodes and data nodes, respectively. Each client periodically communicates with the resource manager to obtain the up-to-date view of the mounted volume. A file operation usually initiates the communications from the client to the corresponding meta node and data node directly, without the involvement of the resource manager. The updated view of the mounted volume, as well as the file metadata are usually cached at the client side to reduce the communication overhead.

### Features

- Scalability for metadata
- Multi-tenancy
- Easy to expand capacity
- Serve large amount of clients simultaneously
- Optimized for both large and small files
- Converged filesystem and S3-compatible interfaces

## Progress Since Sandbox

### Community

The community statistics are extracted from [CNCF devstats](https://chubaofs.devstats.cncf.io/).

- Commits: 516 -> 1995
- Code committers: 17 (5 companies) -> 33 (8 companies)
- Pull requests: 56 -> 666
- Contributors: 27 (5 companies) -> 68 (10 companies)
- Contributions: 1112 -> 3202
- Forkers: 110 -> 175
- Watchers: 851 -> 1327

### Technical

- ChubaoFS was initially developed as a distributed filesystem. However, after joining CNCF, ChubaoFS introduced S3-compatible interface and released v2.0.0 version which gives customers more options to access one copy of data.
- [ChubaoFS CSI driver](https://github.com/chubaofs/chubaofs-csi) released supporting both CSI spec v0.3 and v1.0 versions.
- [ChubaoFS Helm](https://github.com/chubaofs/chubaofs-helm) released.

## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._**

There is a list of [adopters](https://github.com/chubaofs/chubaofs/blob/master/ADOPTERS.md) documented, in which there are a few worth mentioning:

- [JD.com](https://www.jd.com) Top e-commerce company in China. It has the largest ChubaoFS cluster of all adopters, and uses ChubaoFS to provide storage solutions to its containerized applications in a large scale Kubernetes cluster, and to withstand an overwhelming data flow during sales promotions.
- [OPPO](https://www.oppo.com) Consumer electronics and mobile communications company. It uses ChubaoFS to support its cloud platform and diverse customers running on it.
- [Meizu](https://www.meizu.com) Consumer electronics company. It uses ChubaoFS to serve various applications such as database, risk control, cloud backup, push service and so on. It also uses ChubaoFS to support DevOps management.

**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

We have several [maintainers](https://github.com/chubaofs/chubaofs/blob/master/MAINTAINERS.md) covering the main repository and sub-projects.

**_Demonstrate a substantial ongoing flow of commits and merged contributions._**

- [Commit Activity](https://github.com/chubaofs/chubaofs/graphs/commit-activity)
- [Git clones and Visitors](https://github.com/chubaofs/chubaofs/graphs/traffic)
- [Contributors](https://github.com/chubaofs/chubaofs/graphs/contributors)
- [CNCF devstats](https://chubaofs.devstats.cncf.io/)

**_A clear versioning scheme._**

ChubaoFS uses [Semantics Versioning](https://semver.org/). There are 17 [releases](https://github.com/chubaofs/chubaofs/releases) right now, and the latest release version is v2.2.2. The main reason why we changed the major version is that ChubaoFS introduced S3-compatible interface in v2.0.0 which is a huge leap that makes it beyond the scope of distributed filesystem. We plans to release a minor version every 3 months approximately.


## Statement on Alignment with the CNCF Mission

ChubaoFS project aligns with the CNCF mission of making the cloud native ubiquitous by providing a solid cloud storage solution which helps users build stateless cloud native applications.

ChubaoFS serves as an on-premise backend storage. It was designed to support cloud native applications, and some of the key features make it feasible. To name a few, large number of client instances served simultaneously in a single cluster, data isolation through volumes, easy to expand capacity, and so on.

## Comparison with Other CNCF Storage Projects

Although there are a few storage projects in CNCF, ChubaoFS is the only distributed filesystem among all, and is optimized for both large and small files.

### Comparison to Rook

[Rook](https://rook.io/) is storage operators for Kubernetes, which _turns distributed storage systems into self-managing, self-scaling, self-healing storage services_. ChubaoFS can act as an on-premise backend storage system for Rook.

### Comparison to Longhorn

[Longhorn](https://longhorn.io) is a _distributed block storage system for Kubernetes_. Longhorn and ChubaoFS apply to different using scenarios since ChubaoFS is a distributed file and object storage system.

### Comparison to OpenEBS

[OpenEBS](https://openebs.io) is a storage solution for Kubernetes which supports a range of storage engines. ChubaoFS focuses more on the storage engine itself, to provide reliable, scalable and resilient distributed storage system.

## Future Plans

- Integrate with Rook.
- Optimize cross zone feature to improve robustness.
- Spark remote shuffle service plugin which helps spark jobs to run in container platforms.
