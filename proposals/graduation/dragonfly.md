# Dragonfly Graduation Proposal

[Dragonfly](https://d7y.io/) joined the CNCF as a sandbox project in October 2018 and became an
incubating project in April 2020. In January 2020, [Nydus](https://nydus.dev/) became
a sub-project of Dragonfly and was widely used for image acceleration.
In April 2021, the Dragonfly v2.0 was released after architectural optimization and code refactoring.

After an extensive period of development and production experience, Dragonfly is not only widely used in image acceleration,
but also has many use cases in file distribution and AI model distribution. The community has made great progress in growing,
it is governed by partners including Alibaba Group, Ant Group, Baidu Group, Dalian University of Technology, ByteDance,
Intel and JiHu. The project maintainers believe that Dragonfly is ready for graduation.

## Introduction

Dragonfly provides efficient, stable and secure file distribution and image acceleration based on P2P technology
to be the best practice and standard solution in cloud native architectures.
It is designed to improve the efficiency and speed of large-scale file distribution and used in the fields of file distribution,
AI model distribution, cache distribution, log distribution and image distribution.

With the production practice, Dragonfly based on P2P technology to accelerate the image is insufficient to support
faster container launching, such as the Function as a Service (FaaS).
Therefore, we created Nydus as a sub-project of Dragonfly to address this need
and [Nydus Snapshotter](https://github.com/containerd/nydus-snapshotter) become a sub-project of containerd
which is an external plugin of containerd for Nydus.
Nydus implements a content-addressable file system that enhances the current OCI image format with
faster container launch speed, better image space and network bandwidth efficiency, and end-to-end data integrity.

### Image Acceleration

Dragonfly supports various container clients such as containerd, Docker, cri-o, ORAS, etc.
It provides three solutions for image acceleration. The first solution is to use Dragonfly to distribute
images based on P2P technology, which is suitable for large-scale cluster. The second solution is to use Dragonfly and
Nydus to distribute accelerated images, which is suitable for large-scale cluster and faster container launching.
The third solution is to use Nydus to distribute accelerated images, which is suitable for faster container launching.

Production practice and statistical data can refer to:

- [Dragonfly integrates nydus for image acceleration practice](https://www.cncf.io/blog/2022/11/21/dragonfly-integrates-nydus-for-image-acceleration-practice/)
- [The evolution of the Nydus Image Acceleration](https://www.cncf.io/blog/2022/11/15/the-evolution-of-the-nydus-image-acceleration/)
- [Volcano Engine: distributed image acceleration practice based on Dragonfly](https://www.cncf.io/blog/2023/04/13/volcano-engine-distributed-image-acceleration-practice-based-on-dragonfly/)
- [Ant Group security technology’s Nydus and Dragonfly image acceleration practices](https://www.cncf.io/blog/2023/05/01/ant-group-security-technologys-nydus-and-dragonfly-image-acceleration-practices/)
- [Stable and efficient image distribution for 1 billion monthly users with Dragonfly](https://www.cncf.io/case-studies/kuaishou-technology/)

### File Distribution

Dragonfly supports large-scale file distribution and uses P2P technology to eliminate the impact of
origin bandwidth limitations. It supports file distribution of protocols including HTTP, HDFS, etc.
Additionally, it also supports different object storage protocols includes S3, OSS, OBS, etc.

Add [Dfstore](https://d7y.io/docs/concepts/terminology/dfstore) to expand the file distribution capability,
it can depend on different types of object storage, such as S3, OSS, OBS, etc. to provide stable object storage capabilities.
Dfstore uses the entire P2P network as a cache when storing objects. Depend on object storage as
the backend to ensure storage reliability. In the process of object storage, P2P cache is effectively
used for fast read and write storage.

### AI Infrastructure

Dragonfly supports distributing data during AI training and AI inference.
In the AI inference, Dragonfly supports [Triton Server](https://github.com/triton-inference-server/server) and [TorchServe](https://github.com/pytorch/serve)
to use Dragonfly distribution model. In the AI training, [Fluid](https://github.com/fluid-cloudnative/fluid) downloads
dataset through Dragonfly when running based on [JuiceFS](https://github.com/juicedata/juicefs).

There are many use cases in the community, using Dragonfly to distribute data based on P2P technology.
In the inference, the concurrent download model of the inference service can effectively relieve the bandwidth pressure of
the model registry through Dragonfly, and improving the download speed. Users can embed the model into the
accelerated image and use Dragonfly and Nydus for faster container launching.

Production practice and statistical data can refer to:

- [Dragonfly: Intro, Updates and AI Model Distribution in the Practice of Kuaishou - Wenbo Qi, Ant Group & Zekun Liu, Kuaishou Technology](https://sched.co/1PTJb)
- [Dragonfly helps Volcano Engine AIGC inference to accelerate image through p2p technology](https://mp.weixin.qq.com/s/kY6DxRFspAgOO23Na4dvTQ)
- [Get faster pull times with Nydus on CoreWeave](https://docs.coreweave.com/cloud-tools/nydus)

### GitHub Repositories

| Name                                                        | Repository                                                                  |
| :---------------------------------------------------------- | :-------------------------------------------------------------------------- |
| Main Dragonfly v2.x repository                              | [dragonflyoss/Dragonfly2](https://github.com/dragonflyoss/Dragonfly2)       |
| Main Dragonfly v1.x repository                              | [dragonflyoss/Dragonfly](https://github.com/dragonflyoss/Dragonfly)         |
| Main Nydus repository                                       | [dragonflyoss/image-service](https://github.com/dragonflyoss/image-service) |
| Canonical location of the Dragonfly API definition          | [dragonflyoss/api](https://github.com/dragonflyoss/api)                     |
| Helm charts for running Dragonfly and Nydus with Kubernetes | [dragonflyoss/helm-charts](https://github.com/dragonflyoss/helm-charts)     |
| Source for the [d7y.io](https://d7y.io/)                    | [dragonflyoss/d7y.io](https://github.com/dragonflyoss/d7y.io)               |
| Main Dragonfly v2.x console repository                      | [dragonflyoss/console](https://github.com/dragonflyoss/console)             |
| Monitoring support for Dragonfly                            | [dragonflyoss/monitoring](https://github.com/dragonflyoss/monitoring)       |

## Graduation State Criteria

### Have committers from at least two organizations.

Dragonfly has [maintainers (committers) from](https://github.com/dragonflyoss/Dragonfly2/blob/main/MAINTAINERS.md)
Alibaba Group, Ant Group, Baidu Group, Dalian University of Technology, ByteDance, Intel and JiHu.

### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

Dragonfly’s OpenSSF best practices badge can be seen [here](https://bestpractices.coreinfrastructure.org/projects/7103).

### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

A third party security audit was performed by Trail of Bits, you can see the full report [here](https://github.com/dragonflyoss/Dragonfly2/blob/main/docs/security/dragonfly-comprehensive-report-2023.pdf).

### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

- See [GOVERNANCE.md](https://github.com/dragonflyoss/Dragonfly2/blob/main/GOVERNANCE.md) that describes governance
  guidelines and maintainer responsibilities.
- See [OWNERS.md](https://github.com/dragonflyoss/Dragonfly2/blob/main/OWNERS.md) about current and emeritus maintainers.
- See [CONTRIBUTING.md](https://github.com/dragonflyoss/Dragonfly2/blob/main/CONTRIBUTING.md) for general contributing guidelines.

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

- See [GOVERNANCE.md](https://github.com/dragonflyoss/Dragonfly2/blob/main/GOVERNANCE.md) that describes governance
  guidelines and maintainer responsibilities.
- See [MAINTAINERS.md](https://github.com/dragonflyoss/Dragonfly2/blob/main/MAINTAINERS.md) about current maintainers.

### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to [FAQs](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for guidelines on identifying adopters.

The public list of Dragonfly adopters is in the [ADOPTERS.md](https://github.com/dragonflyoss/Dragonfly2/blob/main/ADOPTERS.md).

## Incubation Details

### Link to Incubation Due Diligence(DD) Document

[Incubation Due Diligence(DD) Document](https://docs.google.com/document/d/1FauIs9BXRCioUCYPdO9gfy1uiPw9ESKRu0s4bgd5Dgw/edit#).

### Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

There are no outstanding issues raised during the incubation due diligence remain to be addressed.
