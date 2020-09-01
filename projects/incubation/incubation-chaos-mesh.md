# Chaos Mesh Project Proposal

# Background

**[TOC PR](https://github.com/cncf/toc/pull/367)** 

**[Presentation Slide](https://docs.google.com/presentation/d/115QvSCnT6ROwwV1lK-R10d4MW54fQhtVM2uZDGscpEk/edit#slide=id.g80f082b201_0_0)**

**[GitHub](https://github.com/pingcap/chaos-mesh)**

**Name of Project:** Chaos Mesh  

**Description**:

Chaos Mesh is a versatile Chaos Engineering platform that orchestrates chaos experiments on Kubernetes environments. It features all-around fault injection methods for complex systems on Kubernetes, covering faults in Pod, network, file system, and even the kernel.

Chaos Mesh was originated from the internal chaos engineering platform at PingCAP to ensure the resilience of TiDB, its distributed database system. As the system evolves and testing requirements multiply, the team realized that they need an easy to use, scalable and universal chaos testing platform. The combination of chaos and Kubernetes became the natural choice.  Chaos Mesh uses CRD to define chaos objects, which makes it naturally integrate with the Kubernetes ecosystem. 

At the current stage, it supports the following fault injection types:

- pod-kill: Simulates Kubernetes Pods being killed
- pod-failure: Simulates Kubernetes Pods being continuously unavailable
- container-kill: Simulates Kubernetes Pod’s container being killed
- network-delay: Simulates network delay
- network-loss: Simulates network packet loss
- network-duplication: Simulates network packet duplication
- network-corrupt: Simulates network packet corruption
- network-partition: Simulates network partition
- I/O delay: Simulates file system I/O delay
- I/O errno: Simulates file system I/O errors
- Time skew: Simulates time jumping forward or backward

Besides the versatile chaos types, it also offers Chaos Dashboard (under development), a visualized panel that shows the impacts of chaos experiments on the online services of the system, which makes chaos tests easily observable and manageable. 

**Sponsor from TOC**: TBD

**Unique Identifier**: chaos-mesh

**Preferred Maturity Level**: Sandbox

**License**: Apache License, Version 2.0

**Source control repositories:** [https://github.com/pingcap/chaos-mesh](https://github.com/pingcap/chaos-mesh) (to be moved to [https://github.com/chaos-mesh](https://github.com/chaos-mesh))

**Issue tracker**: GitHub

**Infrastructure Required**:

Chaos Mesh uses in-house Jenkins CI cluster for integration tests. We plan to use CNCF test cluster to automatically run stability tests and performance tests in the future.

**Website**: [https://github.com/pingcap/chaos-mesh](https://github.com/pingcap/chaos-mesh)

**Documentation:** <https://github.com/pingcap/chaos-mesh/wiki>

**Release methodology and mechanics:**

This is currently being defined. Releases every few months with RC process.

**External dependencies (including licenses):** 

* Apache License 2.0:
    * github.com/containerd/containerd
    * github.com/docker/docker
    * github.com/ethercflow/hookfs
    * github.com/go-logr/logr
    * github.com/grpc-ecosystem/go-grpc-middleware
    * github.com/grpc-ecosystem/go-grpc-prometheus
    * github.com/oklog/run
    * github.com/prometheus/client_golang
    * github.com/vishvananda/netlink
    * github.com/vishvananda/netns
    * github.com/kubernetes-sigs/controller-runtime
* MIT:
    * github.com/ghodss/yaml
    * github.com/jmoiron/sqlx
    * github.com/onsi/ginkgo
   * github.com/onsi/gomega
   * github.com/robfig/cron
   * github.com/unrolled/render

* Mozilla Public License 2.0:
    * github.com/go-sql-driver/mysql

* BSD 3-Clause:
    * github.com/golang/protobuf
    * github.com/gorilla/mux

* BSD:
  * github.com/hanwen/go-fuse

**Initial committers:**

<table>
  <tr>
   <td>Name 
   </td>
   <td> Email 
   </td>
   <td> Focus
   </td>
  </tr>
  <tr>
   <td><a href="https://github.com/siddontang">Siddon Tang</a> 
   </td>
   <td> tl@pingcap.com 
   </td>
   <td> Project Lead
   </td>
  </tr>
  <tr>
   <td><a href="https://github.com/zhouqiang-cl">Qiang Zhou</a> 
   </td>
   <td> zhouqiang@pingcap.com 
   </td>
   <td> Project Lead
   </td>
  </tr>
  <tr>
   <td><a href="https://github.com/cwen0">CWen</a> 
   </td>
   <td> cwen@pingcap.com 
   </td>
   <td> Operator, Dashboard
   </td>
  </tr>
  <tr>
   <td><a href="https://github.com/YangKeao">YangKeao</a> 
   </td>
   <td>yangkeao@pingcap.com 
   </td>
   <td>Operator, Dashboard
   </td>
  </tr>
</table>

**_Community Stats_:**

Although it’s just been open-sourced since Dec 31, 2019, Chaos Mesh has been gaining recognition and popularity quickly in the community, with 1400+ stars received on GitHub in a month. It has also been listed under  [CNCF Cloud Native Interactive Landscape](https://landscape.cncf.io/category=chaos-engineering&format=card-mode&grouping=category).

* Stars: 1500+
* Contributors: 23
* Commits: 240+
* Forks: 88+

**Comparison**

This comparison is intended simply to compare fault injection features supported by Chaos Mesh with other well-known chaos engineering platforms. It is not intended to favor or position one project over another. Any corrections are welcome.  

<table>
  <tr>
   <td>
   </td>
   <td>chaos-mesh
   </td>
   <td>chaosmonkey
   </td>
   <td>chaosblade
   </td>
   <td>chaoskube
   </td>
   <td>Litmus
   </td>
  </tr>
  <tr>
   <td>Platform supported
   </td>
   <td>K8s
   </td>
   <td>VMs/ Container
   </td>
   <td>JVM/Container/K8s
   </td>
   <td>K8s
   </td>
   <td>K8s
   </td>
  </tr>
  <tr>
   <td>CPU burn
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>Mem burn
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>container kill
   </td>
   <td>Y
   </td>
   <td>Y
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>pod failure
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>pod kill
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
   <td>Y
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>network partition
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
  </tr>
  <tr>
   <td>network duplication
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
  </tr>
  <tr>
   <td>network corrupt
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>network loss
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>network delay
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>Y
   </td>
  </tr>
  <tr>
   <td>I/O delay
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
  </tr>
  <tr>
   <td>I/O errno
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
  </tr>
  <tr>
   <td>Time skew
   </td>
   <td>Y
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
   <td>N
   </td>
  </tr>
</table>

**Roadmap**:

See [https://github.com/pingcap/chaos-mesh/blob/master/ROADMAP.md](https://github.com/pingcap/chaos-mesh/blob/master/ROADMAP.md)

**Social media accounts**

Twitter:[ @chaos_mesh](https://twitter.com/chaos_mesh)

**Communication channels**

GitHub

Slack channel: [#sig-chaos-mesh](https://tidbcommunity.slack.com/archives/CS7U0QV8E)

Community meetings (Planing)

**Adopters or potential users?**

- [Celo](https://celo.org/)
- [Dailymotion](https://www.dailymotion.com/)
- [NetEase Fuxi Lab](https://www.crunchbase.com/organization/netease-fuxi-lab)
- [JuiceFS](http://juicefs.com/?hl=en)
- [Meituan-Dianping](https://about.meituan.com/en)
- [PingCAP](www.pingcap.com)
- [Xpeng Motors](https://en.xiaopeng.com/)

**Existing sponsorship**

PingCAP

**Statement on alignment with CNCF charter mission**

Our team believes Chaos Mesh will be a great fit for CNCF. As manifested in the CNCF mission:

_“These techniques enable loosely coupled systems that are **resilient, manageable, and observable**. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil. “_

We believe Chaos Mesh is one of the essential enablements to this mission, and it’s also a great addition to the sandbox project scope. By covering comprehensive fault injection methods in Pod, network, file system, and even the kernel, Chaos Mesh aims at providing a neutral, universal Chaos Engineering platform that enables cloud-native applications to be as resilient as they should be. Chaos Mesh uses CRD to define chaos objects, making it naturally integrated with the Kubernetes ecosystem. In addition, it integrates several other projects in the cloud-native ecosystem, such as Helm, Prometheus, and Grafana.

**What are we looking for from CNCF**:

* Project governance - Chaos Mesh is still in the early phase of development, we would love to receive advice and guidance from CNCF on the community governance process and other community-related aspects.
* Neutral home for collaboration - As a universal Chaos Engineering platform,  Chaos Mesh grows by incorporating feedback and collaborations from the community. We would love to see more collaborations between Chaos Mesh and other cloud-native projects.    
* Community marketing support -  Guidance or sponsorship in terms of events  hosting
* Legal support - Any sort of legal assistance in aspects of potential IP, trademark issues
