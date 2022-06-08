**Containerd Graduation Proposal**

Since the [March 2017 announcement](https://www.cncf.io/announcement/2017/03/29/containerd-joins-cloud-native-computing-foundation/) at KubeCon Berlin that containerd would be joining the CNCF, the containerd community has been busy continuing the development and testing of a stable, core container runtime for cloud native use cases.

In late 2017, containerd announced its [1.0 release](https://blog.docker.com/2017/12/cncf-containerd-1-0-ga-announcement/), followed in early 2018 by a 1.1 release that has been updated with several minor fix releases through 2018. The containerd team has just completed the next major release, [1.2.0](https://github.com/containerd/containerd/releases/tag/v1.2.0), with more complete Windows runtime support and a new runtime shim v2 API that is [proving](https://www.cncf.io/blog/2018/09/17/gsoc-18-kata-containers-support-for-containerd/) [valuable](https://twitter.com/resouer/status/1035066629887905792) for [runc](https://github.com/opencontainers/runc) alternatives like [Kata containers](https://katacontainers.io/).

At this time, we believe containerd is ready for the [graduation stage](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#graduation-stage) within the CNCF, and per the guidelines listed there, we detail our readiness in the section below.

### CNCF Graduation Criteria

**Committers from at least two organizations.**

Containerd has had a variety of maintainers and reviewers since its inception, and currently have 12 committers representing Docker, NTT, Google, IBM, Microsoft, Facebook, Tesla, and Cruise Automation. We also recognize **LGTM** rights for a group we call *reviewers*, of which there are currently eight reviewers representing Alibaba, ZTE, Huawei, Docker, Microsoft and an independent developer.

**Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.**

We have recently validated our CII best practices, and have the badge embedded within our main containerd/containerd repository on GitHub. Our specific status (passing) is shown on [our CII project page here](https://bestpractices.coreinfrastructure.org/en/projects/1271).

**Adopt the CNCF Code of Conduct.**

We have adopted the CNCF code of conduct and prominently make this clear in our [GOVERNANCE.md](https://github.com/containerd/project/blob/master/GOVERNANCE.md) document, which applies to all containerd projects and sub-projects.

**Explicitly define a project governance and committer process.**

Our project governance is clearly laid out in our [GOVERNANCE.md](https://github.com/containerd/project/blob/master/GOVERNANCE.md) document, including details on how to become a maintainer and how maintainer and contribution processes are handled. The maintainers for containerd, the CRI project, and all sub-projects are common, and maintained in our core project repo in the [MAINTAINERS](https://github.com/containerd/project/blob/master/MAINTAINERS) file.

**Have a public list of project adopters for at least the primary repository.**

Containerd began life prior to its CNCF adoption as a lower-layer runtime manager for the Docker engine. Continuing today, containerd has widest usage and adoption as the layer between the Docker engine and the OCI runc executor. However, as containerd and its CRI plugin project have merged in January 2018, the combined use of containerd and the CRI plugin has grown to include several public cloud providers, as well as several projects who are attracted to the simplicity of the Go client API library for embedding container runtime capabilities.

**_IBM Cloud Kubernetes Service (IKS)_** - offers containerd as the CRI runtime for v1.11 and, soon, v1.12.

**_IBM Cloud Private (ICP)_** - IBM's on-premises cloud offering has containerd as a "tech preview" CRI runtime for the Kubernetes offered within this product for the past two releases, and plans to fully migrate to containerd in a future release.

**_Google Cloud Kubernetes Engine (GKE)_** - offers containerd in "alpha clusters" on recent versions of Kubernetes.

**_Cloud Foundry_** - The [Guardian container manager](https://github.com/cloudfoundry/guardian) for CF has been using OCI runC directly with additional code from CF managing the container image and filesystem interactions, but have recently migrated to use containerd as a replacement for the extra code they had written around runC.

**_Alibaba's PouchContainer_** - The Alibaba [PouchContainer](https://github.com/alibaba/pouch) project uses containerd as its runtime for a cloud native offering that has unique isolation and image distribution capabilities.

**_Rancher's Rio project_** - Rancher Labs [Rio](https://github.com/rancher/rio) project uses containerd as the runtime for a combined Kubernetes, Istio, and container "Cloud Native Container Distribution" platform.

**_Eliot_** - The [Eliot](https://github.com/ernoaapa/eliot) container project for IoT device container management uses containerd as the runtime.

**_Balena_** - Resin's [Balena](https://github.com/resin-os/balena) container engine, based on moby/moby but for edge, embedded, and IoT use cases, uses the containerd and runc stack in the same way that the Docker engine uses containerd.

**_LinuxKit_** - the Moby project's [LinuxKit](https://github.com/linuxkit/linuxkit) for building secure, minimal Linux OS images in a container-native model uses containerd as the core runtime for system and service containers.

**_BuildKit_** - The Moby project's [BuildKit](https://github.com/moby/buildkit) can use either runC or containerd as build execution backends for building container images. BuildKit support has also been built into the Docker engine in recent releases, making BuildKit provide the backend to the `docker build` command.

**_AWS Firecracker_** - The AWS [Firecracker VMM project](http://firecracker-microvm.io/) uses containerd through integration via a v2 shim and custom external snapshotter. More details on their integration is available in the [Firecracker containerd project](https://github.com/firecracker-microvm/firecracker-containerd).

**_Kata containers_** - The [Kata containers](https://katacontainers.io/) lightweight-virtualized container runtime project integrates with containerd via a custom v2 shim implementation that drives the Kata container runtime.

**_Docker engine_** - As noted in the opening paragraph, Docker continues to consume containerd as a key component within the Docker engine stack, and is actively working to remove Docker engine implementations where containerd implementations can be used as a replacement.

With contributions from Microsoft and AWS, we believe that Azure and AWS are also looking at containerd for potential use within their public cloud offerings as well.

**_Other Projects_** - While the above list provides a cross-section of well known uses of containerd, the simplicity and clear API layer for containerd has inspired many smaller projects around providing simple container management platforms. Two that have come from containerd community participants directly are Michael Crosby's [boss](https://github.com/crosbymichael/boss) project and Evan Hazlett's [stellar](https://github.com/ehazlett/stellar) project, as examples of higher layer functionality that can easily be built on the containerd base.

**NOTE:** *The containerd community will be happy to provide adoption and production usage
statistics where available and/or possible to enumerate.*

**Receive a supermajority vote from the TOC to move to graduation stage.**

We believe this document prepares us for a TOC vote on the graduation readiness of containerd.
