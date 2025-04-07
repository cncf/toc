# CubeFS Graduation Proposal
CubeFS, an open-source project that joined CNCF as a sandbox project in December 2019, has made remarkable progress since then. With its growing community of contributors and users including big names like Oppo, JD, Xiaomi, NetEase, and BIGO, CubeFS has already scaled up to hundreds of petabytes (PB+). As of July 2022, CubeFS has become an incubating project and we are confident that it is now ready to graduate.
## Introduction

[CubeFS](https://cubefs.io)  is a cloud-native storage product that was introduced to the CNCF storage SIG in June 2019. It offers shared storage solutions for various applications and is widely used in scenarios such as AI, big data, and database management. CubeFS boasts several key features that make it a reliable option for users

**Multi-Protocol Interoperability** By unifying access to multiple protocols such as HDFS, POSIX, and S3, it enables one set of data to be accessed in multiple protocols , avoiding the need to copy data between different access approaches.

**Multi-Engine Support** CubeFS supports two storage engines, multiple copies and erasure coding, and has made a lot of multidimensional optimizations for these engines. This enables the flexibly to meet the demands of different business scenarios for IO performance and storage cost.

**Strong Scalability** The metadata of CubeFS has horizontal scalability, and a single cluster can easily support billions of files and exabytes of storage capacity.
**High Performance** It supports multi-level caching, dedicated optimizations for small files, and various high-performance replication protocols.

**Multi-tenancy** CubeFS provides multi-tenant management and achieves access isolation and performance non-interference between tenants through namespace management and QoS.

**Cloud-Native** CubeFS can be easily deployed, run, and upgraded on Kubernetes . With the CSI plugin, PODs can easily access storage volumes.

## Graduation State Criteria
#### Have committers from at least two organizations
* [OPPO](https://www.oppo.com/cn/) Consumer electronics and mobile communications company. It uses CubeFS to support its cloud platform and diverse customers running on it.
* [JD.com](https://www.jd.com) Top e-commerce company in China. It has the largest CubeFS cluster of all adopters, and uses CubeFS to provide storage solutions to its containerized applications in a large scale Kubernetes cluster, and to withstand an overwhelming data flow during sales promotions.
* [BIGO](https://www.bigo.tv/) It uses CubeFS to store logs for AI platform applications that running in the container environment, because CubeFS have excellent concurrent processing capability.
* [Xiaomi Corporation](https://www.mi.com/c/about) It uses CubeFS to serve various applications such as AI, Big Data, shared file system and so on.
* [Netease game company](http://www.netease.com/) It uses CubeFS to support Elasticsearch and AI machine learning business.
#### Have achieved and maintained a Core Infrastructure Initiative [Best Practices Badge](https://bestpractices.coreinfrastructure.org/projects/6232)
* CubeFS‘s OpenSSF badge can be seen in Best Practices Badge.

#### Have completed an independent and third party security audit with results published of similar scope and quality as this example which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation
* The security audit has been initiated by OSTIF and is in progress.

#### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers
* The CubeFS project governance is defined in the [GOVERNANCE](https://github.com/cubefs/cubefs/blob/master/GOVERNANCE.md).

#### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence

* How to participate in the CubeFS project is defined in [CONTRIBUTING](https://github.com/cubefs/cubefs/blob/master/CONTRIBUTING.md) ,The list of maintainers has been stored in [MAINTAINERS](https://github.com/cubefs/cubefs/blob/master/MAINTAINERS.md).

#### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to FAQs for guidelines on identifying adopters
* The CubeFS project has a public list of Project adopter which is maintained in [ADOPTERS](https://github.com/cubefs/cubefs/blob/master/ADOPTERS.md).

## Incubation Details
#### Link to Incubation Due Diligence(DD) Document
* [Incubation proposal](https://github.com/cncf/toc/pull/586)
* [CubeFS Due Diligence](https://docs.google.com/document/d/1WDJEeRDO8pHAetyCbU7TftZfvxSATc3bsLj7VJeNoJ0/edit#heading=h.560y4815aj5r)
#### Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document
* There are currently no known outstanding concerns or recommendations regarding
the CubeFS project.
