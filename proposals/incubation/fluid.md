Fluid was accepted to CNCF on April 28, 2021 at the Sandbox maturity level.. This doc proposes promoting Fluid  to a CNCF Incubating project.
## About Fluid
Fluid is a Kubernetes-native distributed data cache and workload orchestrator that enables AI/Big Data Applications to use data from any storage more efficiently with a high-level abstraction manner  and without changes to the applications themselves.

Fluid aims to turn different distributed cache systems (Alluxio, JuiceFS, Vineyard,  CubeFS and so on) into self-managing, self-scaling, self-healing and observable cache services inside of Kubernetes by providing the common framework of Fluid. Fluid enables Kubernetes schedulers to make intelligent, topology-aware scheduling plans regarding where the distributed data cache system is located. 

Fluid contains the following key features:

- **Dataset Abstraction**: Implements the unified abstraction for datasets from multiple storage sources, with observability features to help users evaluate the need for scaling the cache system.

- **Scalable Cache Runtime**: Offers a unified access interface for data operations with different runtimes, enabling access to third-party storage systems.

- **Automated Data Operations**: Provides various automated data operation modes to facilitate integration with automated operations systems.

- **Elasticity and Scheduling**: Enhances data access performance by combining data caching technology with elastic scaling, portability, observability, and data affinity-scheduling capabilities.

- **Runtime Platform Agnostic**: Supports a variety of environments and can run different storage clients based on the environment, including native, edge, Serverless Kubernetes clusters, and Kubernetes multi-cluster environments.

Fluid is built in the Kubernetes native fashion. It lies between existing underlying cloud native storage systems and the upper layer data-intensive applications.  Specifically, Fluid's architecture is logically split into **control plane** and **data plane**. The following diagram shows the different components.

- **Control Plane**：The control plane is composed of **Dataset/Runtime Operator** and **Application Manager**
   **Dataset/Runtime Operator**: Responsible for the scheduling and orchestration of datasets and their supporting runtimes in Kubernetes. This includes scheduling, migration, and elastic scaling of the runtime for datasets, as well as automated operations for dataset support, such as fine-grained data preheating, such as specifying preheating for a specific folder; controlling metadata backup and recovery to improve data access performance for scenarios with massive small files; and setting pinning policies for cached data to avoid performance fluctuations caused by data eviction.
   **Application Manager**:  Responsible for the scheduling and operation of application Pods that use datasets, which is divided into two core components: the Scheduler and the Webhook.
   (1) Scheduler: schedule application Pods that use datasets in the Kubernetes cluster. By incorporating cached information obtained from the Runtime, Pods that use datasets are preferentially scheduled to nodes that have data caching, without the need for users to specify caching nodes. 
  (2) Sidecar Webhook: For Kubernetes environments where the csi-plugin cannot be run, the Sidecar webhook automatically replaces the PVC with a FUSE sidecar and controls the startup order of containers in the Pod to ensure that the FUSE container starts first.

- **Data Plane**：The data plane is composed of **Runtime Plugin** and **CSI Plugin**.
   **Runtime Plugin**: As a highly extensible plugin, it can support various data access engines. Fluid achieves this by abstracting some common features, such as the use of cache media, quotas, directories, etc., making it extensible with different distributed cache engine implementation technologies. For example, the AlluxioRuntime uses a Master-Slave architecture, while the JuiceFSRuntime uses a Worker P2P architecture, both of which can be configured in the CRD of the Runtime. This plugin not only supports specific Runtimes like Alluxio and JuiceFS, but also supports a generic ThinRuntime, enabling users to access generic storage without the need for development.
   **CSI Plugin**: The storage client is started in a containerized manner, completely decoupled from the business container. Upgrading the CSI plugin will not affect the business container, and it also supports deploying multiple versions of the storage client in the same Kubernetes cluster. Running the client independently in a Pod also provides observability within the Kubernetes system. Additionally, resource quotas can be set for the client's computing resources.

**In Kubernetes environment**, unlike traditional PVC-based storage abstraction, Fluid takes an Application-oriented perspective to abstract the “process of manipulating data on Kubernetes”. It introduces the concept of elastic Dataset and implements it as a first-class citizen in Kubernetes to enable Dataset CRUD operation, permission control, and data access acceleration. 

-  Zero Intrusion: Seamlessly integrates with the Kubernetes ecosystem
Scalability - Supports various caching engines, containerized operation of each component
-  High Elasticity: Supports horizontal scaling of computation, implements CRD, Operator, etc., based on Kubernetes
-  Multi-Platform: Supports Serverless and edge Kubernetes

There are also two research papers about fluid projects published in top-tier academic conferences or journals as follows:
1. **Rong Gu, Kai Zhang, Zhihao Xu, et al. [Fluid: Dataset Abstraction and Elastic Acceleration for Cloud-native Deep Learning Training Jobs](https://ieeexplore.ieee.org/abstract/document/9835158). IEEE ICDE, pp. 2183-2196, May, 2022. (Conference Version)**
2. **Rong Gu, Zhihao Xu, Yang Che, et al. [High-level Data Abstraction and Elastic Data Caching for Data-intensive AI Applications on Cloud-native Platforms](https://ieeexplore.ieee.org/document/10249214). IEEE TPDS, pp. 2946-2964, Vol 34(11), 2023. (Journal Version)**

## Fluid in the CNCF Sandbox
### Community & Growth
The Fluid devstats page and dashboards can be found [here](https://fluid.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m). Under CNCF we have seen significant growth,The community has grown since the project entered the CNCF sandbox. As of May 2024 Fluid has increased 748 contributors since joining CNCF, and has grown rapidly over the last year. These include committers from many organizations such as Nanjing University,  Alibaba,  Alluxio, JuiceData, SAP,  Tencent，Baidu, Unisound, China Telecom, NetEase, 4paradigm, ByteDance, IBM, Inception and so on. 

In the past six months, there were 60 commits including 62 merged PRs per month on average. The community has grown rapidly since the project entered the CNCF sandbox:
- We held bi-weekly community meetings constantly. The average number of meeting attendees is ~20(In addition, more than 20 participants watched the online live stream in the community group). There are 50+  community meetings so far. The records of meetings and meeting videos can be found: 
https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule
- Number of contributors: 20 ->** 748**
- Github stars: 200 -> **1500+**
- Github forks: 20 -> **900+**
- Adopters: **37** , **16** in production
  
### Governance

Fluid implements the open and vendor neutral governance model where the key decisions are made via open conversation and consensus in the Fluid community channels. Fluid has established [community membership](https://github.com/fluid-cloudnative/fluid/blob/master/GOVERNANCE.md)  including contributors, committers, maintainers and Technical Oversight Committee (TOC). The TOC functions as the core management team that oversees the community, but members of the TOC do not have significant authority over other members of the community. Now Fluid has four TOC members from three different organizations and our membership definition allows members/approvers to become maintainers or TOC by following the guidelines and satisfying the requirements.
Committers are community members who have shown that they are committed to the continued development of the project through ongoing engagement with the community. One can be a committer by following:

1. Nominated by one of the maintainers or TOC members by open a community issue;
2. Have more than two +1 vote from maintainer or TOC member;
3. Community chair send invitation letters, and add the GitHub user to the “Fluid” team 
Roadmap decisions for the project and the community are made by a consensus in the community. The project’s public roadmap is listed here, and any community member can suggest the roadmap update, like here.


### Technology
Fluid  was accepted into the CNCF Sandbox on April  28th, 2021. Since that date, Fluid has released serveral version substantial new features, you can checkout our [release notes](https://github.com/fluid-cloudnative/fluid/releases) for details.

### Collaboration with other CNCF projects

- **Kubernetes:** Fluid leverages Kubernetes and its API as the control plane..
- **Helm**: is used to deploy Fluid’s control plane to Kubernetes.
- **Vineyard**: is one kind of Fluid runtime, which is used to share distributed data in various big data tasks. Here is the [reference](https://github.com/fluid-cloudnative/fluid-client-python/blob/master/examples/03_dataflow_with_vineyard/vineyard-en.ipynb).
- **Kubeflow**:Kserve can use Fluid to accelerate large language model loading. Here is the [reference](https://github.com/kserve/kserve/blob/master/docs/samples/fluid/README.md).
- **KubeDL:** uses Fluid as the cache engine to  improve the data access speed of datasets which can reduce the time spent in training and increase the utilization rate of GPU and the efficiency of model training. Here is the [reference](https://github.com/kubedl-io/kubedl/blob/master/docs/cache_backend.md).
- **CubeFS**: makes for a very powerful combination with Fluid to provide dataset management, data access acceleration and enhance observability along with resource management. Here is the [reference](https://github.com/fluid-cloudnative/fluid/blob/master/addons/cubefs/v3.2/dev-guide/cubefs-v3.2.md).
- **Knative**: can use Fluid to accelerate data access with sidecar mode which can avoid CSI plugin installation. Here is the [reference](https://github.com/fluid-cloudnative/fluid/blob/master/docs/en/samples/knative.md).
- **Open Cluster Management **:can use Fluid to accelerate data access to multiple managed K8s clusters. Here is the [reference](https://github.com/open-cluster-management-io/addon-contrib/tree/main/fluid-addon) and [demo](https://www.youtube.com/watch?v=mIeKY54nhaU).
- **Prometheus**:to collect and analyze both system and user level metrics. Here is the [reference](https://github.com/fluid-cloudnative/fluid/blob/master/docs/en/operation/monitoring.md).

## Comparison with other CNCF Storage projects

Many existing cloud native  storage projects are  designed for Kubernetes environments. Although they are about data stroage, Fluid has some unique goals and characteriscts. For example:
- **Dataset Abstraction and Cache Management**: Allow users to create distributed data caches from datasets that are stored in different types of remote distributed storage, speeding up data access for your applications running in Kubernetes.
- **High Extensibility**: The distributed cache provider can easily add new types of data caching engines. The storage provider can add new types of data access client by using ThinRuntime Strategy.
- **Data Cache Aware Scheduling**: Make workloads using dataset  become aware of the data cache locality in Kubernetes and schedule them onto the nodes  based on  the cost of data transmission, such as within the same node, rack, availability zone, or region.
- **Elastic Cache**: Provide the ability to elastically scale the distributed cache up and down. This includes manual, timed, and automatic elastic scaling based on metrics.  It’s according to the needs of the workload, which leads to high I/O throughput or cost saving.
- **Data and Task Flow**: Define automated operation flows centered on datasets, control the sequential execution order of data migration, preheating, and tasks based on dependencies.
