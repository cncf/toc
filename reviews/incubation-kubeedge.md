# KubeEdge proposal for CNCF Incubation
## Table of Contents

   * [KubeEdge proposal for CNCF Incubation](#kubeedge-proposal-for-cncf-incubation)
      * [Table of Contents](#table-of-contents)
      * [Background](#background)
      * [Alignment with Cloud Native](#alignment-with-cloud-native)
      * [Incubation State Requirements](#incubation-state-requirements)

## Background
- [KubeEdge Sandbox Proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/kubeedge.adoc)
- [Project Review Proposal](https://github.com/cncf/toc/pull/440)
- [Original CNCF TOC meeting slides](https://docs.google.com/presentation/d/1e1ahun91Abn2xvX7Z8PVgGBz6c7Q7iO027XVzVuffDg/edit#slide=id.g25ca91f87f_0_0)
- [Incubation Application Presentation for SIG-Runtime](https://docs.google.com/presentation/d/1k272w_hTCsjwGHY6h1XGJeBMpuV5Cxry2ozyoMDvsPU/edit?usp=sharing)

[KubeEdge](https://kubeedge.io/en/) is an open source system for extending native containerized application orchestration capabilities to hosts at Edge. It's built upon Kubernetes and provides fundamental infrastructure support for network, application deployment and metadata synchronization between cloud and edge.

It aims to resolve three major challenges for edge computing: network reliability and bandwidth limit between cloud and edge, resource constraint at edge, highly distributed and heterogeneous device managements.

**KubeEdge provides:**
1. Seamless Cloud-Edge Communication for both metadata and data
2. Edge Autonomy: Autonomous operation of Edge even when disconnected from cloud.
3. Low Resource Requirement: KubeEdge can work in constrained resource situations (low memory, low bandwidth, low compute)
4. Simplified Device Communication: Easy communication between applications and devices for IOT and Industrial Internet of Things (*IIoT*)

**KubeEdge was accepted as a CNCF Sandbox project on Mar 18, 2019.**
- [KubeEdge Sandbox Proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/kubeedge.adoc)
- [Original CNCF TOC meeting slides](https://docs.google.com/presentation/d/1e1ahun91Abn2xvX7Z8PVgGBz6c7Q7iO027XVzVuffDg/edit#slide=id.g25ca91f87f_0_0)

**Continuous Momentum**

- Least Release: v1.3
- Num of Contributors: 30 => **300+**
- Github Stars: 400+ => **2500+**
- Github Forks: 100+ => **600+**
- Contributing member organizations: 1 => **25+**
- Maintainers (Approvers): 7 (1 company) => **14** (5 companies)
- Notable improvements since sandbox include:
  + Added Device Management framework to simplify edge application development.
  + Provided Bluetooth and Modbus protocol support for Industrial Internet of Things (*IIoT*)
  + Introduced EdgeMesh to simplify service communication between edge nodes
  + More container runtime integration supports: containerd, cri-o, kata containers
  + Added high availability of cloud component CloudCore
  + Added CSI integration support on the Edge
  + Support fetching edge pods' logs in the cloud


## Alignment with Cloud Native

KubeEdge falls in the scope of [CNCF Runtime SIG](https://github.com/cncf/sig-runtime).

**KubeEdge targets on:**
  - Building an open edge computing platform with cloud native technologies
  - Helping users extending their business architecture, applications, services, etc. from cloud to edge in same experience
  - Implementing extensible cloud-edge architecture based on Kubernetes
  - Integration with CNCF projects, including (but not limited to) containerd, cri-o, Prometheus, Envoy, etc.
  - Seamless development, deployment and execution of complex workloads at edge with optimized resources
## Incubation State Requirements
1. _Document that it is being used successfully in production by at least three independent end users which, in the TOC's judgement, are of adequate quality and scope._

We have the list of public adopters in the repo, some notable users include [here](https://github.com/kubeedge/kubeedge/blob/master/ADOPTERS.md):
- [Raisecom](https://github.com/kubeedge/kubeedge/blob/master/ADOPTERS.md#raisecom-technology-coltd), is using KubeEdge to manage AI inferencing applications on the edge to help improve industrial manufacturing safety and efficiency.
- [WoCloud](https://cucc.wocloud.cn/), is using to offer a commercial hosted edge computing service.
- [Xinghai IoT](https://github.com/kubeedge/kubeedge/blob/master/ADOPTERS.md#xinghai-iot), is using KubeEdge in its smart campus platform.
- [KubeSphere](https://kubesphere.io/), is using to offer an open source on-prem edge computing service in its container platform.
- [HUAWEI CLOUD](https://huaweicloud.com), is using to offer a commercial hosted edge computing service.
- User A (awaiting confirmation to publish name), is using KubeEdge in its highway electronic tolling system.
- User B (awaiting confirmation to publish name), is using KubeEdge in its Mine Brain solution and served 1000+ coal mines.

There are a few other adopters we are waiting for confirmation to publish name.

2. _Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._

We have 14 approvers, from 5 different companies, including Arm, China Unicom, Huawei, infoblox.com, inovex.de. See OWNERS files under [KubeEdge repositories](https://github.com/kubeedge) for details.

3. _Demonstrate a substantial ongoing flow of commits and merged contributions._

We are seeing a constant stream of improvements and features from the maintainers and community. See the stats here:

* [Commits per week over the last 6 months](https://kubeedge.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=now-6M&to=now)

* [Issue Opened/Closed per week over the last 6 months](https://kubeedge.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=now-6M&to=now)

* [New PRs per week over the last 6 months](https://kubeedge.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-6M&to=now&var-period=d7&var-repogroup_name=All)

4. _A clear versioning scheme._

KubeEdge follows the [semantic versioning spec](https://semver.org/), and now is under every 3 months release cadence.

We have releases documented at: https://github.com/kubeedge/kubeedge

The latest 4 releases are:

- [release 1.0](https://github.com/kubeedge/kubeedge/blob/master/CHANGELOG-1.0.md)
- [release 1.1](https://github.com/kubeedge/kubeedge/blob/master/CHANGELOG-1.1.md)
- [release 1.2](https://github.com/kubeedge/kubeedge/blob/master/CHANGELOG-1.2.md)
- [release 1.3](https://github.com/kubeedge/kubeedge/blob/master/CHANGELOG-1.3.md)

5. _Roadmap_

We have a lot of plans for future development, and major items in our [roadmap](https://github.com/kubeedge/kubeedge/blob/master/docs/roadmap.md) are:

- Support edge pod metrics collection in the cloud and integrate with Prometheus.
- Gateway based cross subnet communication on the edge.
- Support edge-cloud communication using EdgeMesh.
- Improve edge device management extensibility, support to define/integrate custom device protocol.
- Decentralized Security for applications on the edge.
