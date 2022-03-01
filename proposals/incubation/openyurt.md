# OpenYurt for incubation

## Background

[OpenYurt](https://openyurt.io/) was designed to resolve various challenges for users to manage edge infrastructures using
cloud Kubernetes control planes. With OpenYurt, users are able to leverage the entire Kubernetes ecosystems
to manage their edge applications efficiently. OpenYurt was accepted as a CNCF Sandbox project on Sept 8, 2020.

## About OpenYurt

OpenYurt is a platform built on top of Kubernetes. It provides rich capabilities in supporting both the near edge use cases and the
far edge use cases. In general, OpenYurt addresses most of the management problems for supporting those use cases from the following
aspects:
- **Extending Kubernetes with addons**: It introduces various CRDs, corresponding controllers, and several node demons to support
features such as node autonomy, region-aware deployment, region-aware container service, cloud-to-edge tunneling,
inter-region networking and so on.
- **Integrations with other OSS projects**: For instances, [EdgeX Foundry](https://github.com/edgexfoundry) has been integrated in OpenYurt and
users now can manage IoT devices registered in EdgeX Foundry through Kubernetes CRDs; OpenYurt supports running webassembly applications
using [WasmEdge](https://github.com/WasmEdge/WasmEdge).
- **Separate designs for different use cases**: Instead of using a one-for-all approach, OpenYurt proposes different architectures to
handle near edge use cases and far edge use cases differently given that Kubernetes has separate concerning limitations in supporting them.

## Progress Since Sandbox

### Community
The OpenYurt devstats page and dashboards can be found [here](https://openyurt.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, OpenYurt currently has [82](https://openyurt.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [25](https://openyurt.devstats.cncf.io/d/5/companies-table?orgId=1)
companies. On average, there are [39 commits per month](https://openyurt.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) contained within [25 merged PRs per month](https://openyurt.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now).
- The community has grown since the project entered the CNCF sandbox.
  - We held bi-weekly community meetings constantly (total 41 as of Feb 2022). The meeting records can be found in [here](https://search.bilibili.com/video?keyword=openyurt). The average number of meeting attendees is ~40.
  - Number of contributors: 10+ -> **82**
  - Github stars: 200+ -> **1100+**
  - Github forks: 30+ -> **250+**
  - Contributing organizations: 2 -> **25**
  - With the help from community contributors, the [project website](https://openyurt.io/) has been redesigned and an [OpenYurt user experience center](http://47.242.50.237/login) has been built for new learners!

### Governance

The project matured remarkably. We have established [the community membership roadmap](https://github.com/openyurtio/community/blob/main/community-membership.md), and quite a few active and qualified contributors
have been promoted to maintainers. Currently, we have 9 maintainers from six different organizations.

### Technology

We have significantly expanded the covered technical areas for the project development since OpenYurt joined the Sandbox.
Currently, OpenYurt has made feature developments in at least the following technical areas:
- **Edge networking**: we have formed a SIG for edge networking in the community, i.e., project [**raven**](https://github.com/openyurtio/raven).
The SIG works on developing a secure L3 network solution to achieve cross-region communication, which enables the service discovery and Pod accesses
across per-region LANs. The SIG also works on optimizing the existing OpenYurt tunnel solution with much less resource overhead. The related design
document can be found [here](https://github.com/openyurtio/openyurt/blob/master/docs/proposals/20211123-enhancement-of-cluster-networking.md).
- **Edge device management**: we have worked closely with core developers from the EdgeX Foundry community, designed several CRDs and implemented
a set of controllers to interact with EdgeX Foundry instances for IoT device management.
- **Edge scalability**: we have formed a SIG in the community for investigating managing numerous far edge nodes and the running applications, i.e., project [**kole**](https://github.com/openyurtio/kole).
The project aims to break through the limit of the Kubernetes cluster size so that we can support managing
millions of small far edge nodes in a single OpenYurt cluster.
- **Lightweight edge**: we have developed a lightweight version of `kubelet` for tiny edge nodes. It can additionally support a lightweight webassembly runtime (i.e., [WasmEdge](https://github.com/WasmEdge/WasmEdge)).
- **Advanced edge autonomy**: Each edge region is equipped with an apiserver which helps for selecting a health leader agent for reporting
precise edge node states given that not all edge nodes can reliably connect to the cloud control plane, and reducing duplicated metadata transferred
between the edge and the cloud.
- **UI**: we have built a front-end [dashboard](https://github.com/openyurtio/yurt-dashboard) to help new users to quickly learn all features provided by OpenYurt.

### Highlighted presentations/blogs from community members

- [Kubecon NA 2021: Take the Edge to the Cloud with OpenYurt and EdgeX Foundry](https://www.youtube.com/watch?v=rJHXSpQIjaY)
- [Kubecon EU 2022: Extending OpenYurt for End-to-End Deployment of Video Analytics Workloads ](https://www.youtube.com/watch?v=GGtxtX4LF-E)
- [WasmEdge and OpenYurt bring cloud computing to the edge](https://www.cncf.io/blog/2022/02/07/wasmedge-and-openyurt-bring-cloud-computing-to-the-edge/)
- [Integrating OpenYurt with Sangfor edge platform](https://mp.weixin.qq.com/s/sfd7dMKW15U_vIfNtRamlQ)

### Collaboration with other CNCF projects
- [WasmEdge](github.com/WasmEdge/WasmEdge): OpenYurt allows for running webassembly applications by integrating the WasmEdge runtime.


## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._**

Full list of the adopters of OpenYurt can be found [here](https://github.com/openyurtio/openyurt/issues/300).
Highlighted examples includes:
- [**China Merchants Group**](https://www.cmhk.com/en/): Using OpenYurt to manage edge nodes and applications installed in parking lots.
- [**China Telecom**](https://www.chinatelecom-h.com/en/global/home.php): Using OpenYurt to manage nodes across branch sites.
- [**Sangfor Technologies**](https://www.sangfor.com/en): Using OpenYurt for devOps leveraging its edge autonomy capability.
- [**MWPark**](http://www.mwpark.cn/): Using OpenYurt to manage the edge nodes in parking lots spread over the cities.

**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

There are [42](https://github.com/orgs/openyurtio/people) project members for OpenYurt, out of which nine are project maintainers from six different organizations.
The maintainers list can be found [here](https://github.com/openyurtio/openyurt/blob/master/MAINTAINERS.md).

**_Demonstrate a substantial ongoing flow of commits and merged contributions._**

A short summary has been provided in the aforementioned community section. Detailed statistics can be found in the following devstats links:
- [Commit Activity](https://openyurt.devstats.cncf.io/d/74/contributions-chart?orgId=1)
- [New contributors since sandboxing](https://openyurt.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1568098800000&to=now)
- [Issues open/closed since sandboxing](https://openyurt.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1599721200000&to=now)
- [PRs open/closed since sandboxing](https://openyurt.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1599721200000&to=now)
- [Contributing organizations](https://openyurt.devstats.cncf.io/d/5/companies-table?orgId=1)


**_A clear versioning scheme._**

OpenYurt uses the [semantic versioning scheme](https://semver.org/#semantic-versioning-specification-semver).
OpenYurt follows a quarterly release cadence with version numbers using format of MAJOR.MINOR.PATCH. The latest release is v0.6.0.

## Statement on Alignment with the CNCF Mission

It is [predicted](https://www.youtube.com/watch?v=W1v2Gb6URsk) that more than 75% of enterprise data will be generated in the edge side
by 2025. Given the growth of the computation need and the complexities of managing diverse edge environments,
improving the efficiency for operating and managing edge applications
becomes a high-demanding request in the edge computing community. OpenYurt aims to resolve various challenges in managing
edge applications through cloud control planes without compromising the user experiences provided by Kubernetes. We believe it strictly
aligns with the mission of CNCF in enabling end users from the edge computing community to benefit from the existing cloud native technologies.

## Comparison with other CNCF projects

[Kubeedge](https://github.com/kubeedge/kubeedge) and [Superedge](https://github.com/superedge/superedge) are
other two CNCF projects for supporting edge use cases in Kubernetes. Like them, OpenYurt starts with
resolving the edge autonomy problem so that the cluster still functions even if the cloud-edge network connectivity
is unreliable. As the community grows and more and more requirements are brought out, OpenYurt moves quickly
in the feature development to tackle various problems in different areas. For example,
- We has designed and implemented a prototype in OpenYurt to manage up to one million
  far edge nodes in a single Kubernetes cluster (the first attempt in the OSS community to the best of our knowledge).
- We augment the basic node autonomy in OpenYurt with unique capabilities for reducing the
  amount of metadata transferred between the cloud and the edge, hence reducing the cost of using the public network.
- OpenYurt supports secure cross-region(LAN) east-west networking by introducing per-region gateways and using IPSec.

Although having conceptual overlaps with Kubeedge and Superedge, OpenYurt is driven by distinct technical
paths for emerging new challenges.

In the area of IoT device management, OpenYurt choose to provide API abstractions and integrate with an existing
open source IoT device management platform, i.e., EdgeX Foundry, instead of supporting various IoT protocols for
different IoT devices internally, like that has been done in Kubeedge.

## Future plan

OpenYurt maintains a steady quarterly release pace with clear roadmaps discussed in community meetings.
The current release milestones are tracked through the [repo project dashboard](https://github.com/openyurtio/openyurt/projects/2).

In the future, we will keep the SIG organizations for the community collaborations for specific technology domains
which we believe would be efficient given the broad problems we face in the edge computing area.
