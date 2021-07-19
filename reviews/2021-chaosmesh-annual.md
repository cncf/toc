# Chaos Mesh® 2021 Annual Review

- [Chaos Mesh® 2021 Annual Review](#chaos-mesh-2021-annual-review)
  - [Background](#background)
  - [Alignment with Cloud Native](#alignment-with-cloud-native)
  - [Year in Review](#year-in-review)
  - [Annual Review Contents](#annual-review-contents)
  - [Project Links](#project-links)

## Background

Chaos Mesh is a versatile Chaos Engineering platform that orchestrates chaos experiments on Kubernetes environments. 

By covering comprehensive fault injection methods in container, Pod, network, file system, and even the kernel, Chaos Mesh aims at providing a neutral, universal Chaos Engineering platform that enables cloud-native applications to be as resilient as they should be. 

Chaos Mesh contains the following components:

![Chaos Mesh Architecture](https://chaos-mesh.org/assets/images/chaos-mesh-79ecf122938b5863a1baf9d64b1ccd8d.svg) 

* **Chaos Dashboard:** a Web UI for managing, designing, monitoring Chaos Experiments, through which users can operate and observe chaos experiments. It also provides RBAC authorization.

* **Chaos Controller Manager:** the core logic component of Chaos Mesh. It’s responsible for the scheduling and managing chaos experiments. The Controller Manager contains multiple CRD Controllers, such as Workflow Controller, Scheduler Controller, and various types of fault controllers.

* **Chaos Daemon:** the main executive component of Chaos Mesh. Chaos Daemon runs in [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) mode and has privileged permissions by default (can be turned off). Chaos Daemon mainly interferes with specific network devices, file systems, kernels, etc. by intruding into the target Pod Namespace.

**Chaos Mesh was accepted into CNCF Sandbox on July 15, 2020**

* [Sandbox proposal PR](https://github.com/cncf/toc/pull/367)

* [Sandbox proposal doc](https://github.com/cncf/toc/blob/1809ba96b31a562c8968265f06b5395c462c8b35/proposals/chaos-mesh.md)


* [Sandbox presentation slides](https://docs.google.com/presentation/d/115QvSCnT6ROwwV1lK-R10d4MW54fQhtVM2uZDGscpEk/edit#slide=id.g80f082b201_0_0)

## Alignment with Cloud Native

Chaos Mesh falls into the scope of[ CNCF TAG Network](https://github.com/cncf/tag-network).

Chaos Mesh adopts a Kubernetes architecture that uses Kubernetes CustomResourceDefinitions (CRDs) to define chaos objects, making it naturally integrated with the Cloud Native ecosystem.

By orchestrating chaos experiments via comprehensive fault injection methods, Chaos Mesh provides a vendor-neutral, scalable chaos engineering platform that enables cloud-native applications to be **_resilient, manageable, and observable_**.  

In addition, it integrates several other projects in the cloud-native ecosystem, such as Helm, Prometheus, and Grafana.

## Year in Review

* Chaos Mesh 1.0 released, and moving closely to 2.0.
* Key features added:
  * Chaos Dashboard as a visual aid to help users manage and monitor chaos experiments through a Web UI.
  * Workflow orchestration support for users to arrange chaos experiments on their own.
  * Interrupt mechanism and automatic detection of application status steps.
  * Authorization mechanism based on Kubernetes RBAC permission policies.
  * New types of fault injection:
      * StressChaos
      * DNSChaos
      * JVMChaos
      * AWSChaos
      * HTTPChaos
  * Provides complete native event information.
  * Grafana plug-in to help users monitor chaos experiments.
  * A toolkit, Chaosd, so that users can simulate failures directly through physical nodes.

* Community
  * Stars: 3.7k (from 1.5k)
  * Contributors: 102 (from 23)
  * Commits: 1,100+ (from 240)
  * Forks: 430+ (from 88)
  * Adopters: 40+

## Annual Review Contents

* **Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity.**

  According to [chaosmesh.devstats.cncf.io](http://chaosmesh.devstats.cncf.io), the Chaos Mesh project average monthly stats in the past year are as follows:
    * [980 contributions](https://chaosmesh.devstats.cncf.io/d/74/contributions-chart?orgId=1);
    * [40 issus opened and 31 closed](https://chaosmesh.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All);
    * [41 active contributors](https://chaosmesh.devstats.cncf.io/d/74/contributions-chart?orgId=1);
    * [82 merged PRs](https://chaosmesh.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all).

* **How many maintainers do you have, and which organizations are they from?** 

    Chaos Mesh currently has[ 8 Maintainers](https://github.com/chaos-mesh/chaos-mesh/blob/master/MAINTAINERS.md#project-maintainers) from 3 organizations and 9 Committers from 4 organizations. The roles and the corresponding responsibilities are defined in the Project [Governance](https://github.com/chaos-mesh/chaos-mesh/blob/master/GOVERNANCE.md). 

* **What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?**  

    This past year in CNCF has seen huge growth in adoption, from 6 to over 40.  Some adoption highlights are:
    
  * [Apache APISIX](https://github.com/apache/apisix) \
  Apache APISIX is a dynamic, real-time, high-performance open source API gateway, which provides rich traffic management features, such as load balancing, dynamic upstream and more. \
  APISIX integrates Chaos Mesh into open-source CI pipelines, to further enhance APISIX's resiliency and robustness.
  * [ByteDance](https://bytedance.com/en/) \
ByteDance is a technology company operating a range of content platforms including TikTok, Lark, Helo, Vigo Video, Douyin, and Huoshan in over 150 markets and 75 languages. \
ByteDance's self-developed chaos engineering platform is mainly used by the company's own technology system. As there are some cloud-native deployment services involved, they integrated Chaos Mesh as the underlying fault injection engine, which is a key supplement to ByteDance’s chaos engineering platform.
  * [NetEase Fuxi Lab](https://fuxi.163.com/en/about.html) \
    NetEase Fuxi AI Lab is China’s first professional game AI research institution. Researchers use their Kubernetes-based Danlu platform for algorithm development, training and tuning, and online publishing.  \
    They use Chaos Mesh to improve the stability of their internal hybrid cloud. In addition, their users with cloud platforms also access Chaos Mesh to test the stability of user services. See [Use case](https://chaos-mesh.org/blog/how-a-top-game-company-uses-chaos-engineering-to-improve-testing).
  * [KingNet](https://www.kingnet.com/) \
KingNet’s main business includes the development, operation and distribution of premium entertainment content. \
KingNet mainly uses Chaos Mesh for testing the availability of multiple data centers and microservice links. Chaos Mesh also helps them with mocking service unavailability or abnormal network conditions.
  * [Qiniu Cloud](https://qiniu.com/en) \
Qiniu Cloud is a distributed cloud system that carries massive amounts of data, and is one that requires high data consistency and high availability, with the data quantity level at 1 trillion+.  \
To ensure the reliability of cloud storage products, they use Chaos Mesh to perform chaos tests on metadata and the underlying storage system under conditions such as: single point of failure of services, network abnormality, abnormal resource consumption (CPU, memory, I/O), etc.
  * [Tencent](https://www.tencent.com/en-us) \
After Tencent Interactive Entertainment migrated their online operations to the Tencent Cloud Kubernetes engine, they wished to provide users with a more stable and reliable experience, which is why they introduced Chaos Mesh. Tencent mainly use Chaos Mesh to simulate the following types of failures:
        * Fault isolation, such as simulating pod abnormality, and checking whether the system can automatically isolate fault instances;
        * Service degradation, such as simulating a downstream recommended service failure through network failure, and verifying whether the local cache is effective;
        * Verifying if the alarm works, for example, purposefully burning the CPU to 90%, and checking whether the alarm is timely issued in time.
  * [Xpeng](https://en.xiaopeng.com/) \
Xpeng Motors is China's leading smart electric vehicle designer and manufacturer, as well as a technology company integrating cutting-edge Internet and AI innovation. They use Chaos Mesh in the following scenarios: 
        * Rolling updates of microservices and lossless verification of traffic;
        * Microservices, multi-registries, multi-party synchronization, and traffic lossless verification;
        * mqtt cluster two-way subscription verification;
        * Exactly-once consumer business verification for message queues;
        * Simulation of weak 4G network for in-vehicle systems, saving drive test costs;
        * AIOPS anomaly detection dataset generation.

* **How has the project performed against its goals since the last review?**

  * From the feature’s perspective, 5 new fault injection types are now supported, namely: StressChaos, DNSChaos, JVMChaos, AWSChaos, and HTTPChaos; we have also updated the Chaos Dashboard for better user experience; Workflow orchestration support for users to arrange chaos experiments on their own; authorization mechanism based on Kubernetes RBAC permission policies; Grafana plug-in to help users monitor chaos experiments; and last but not least, Chaosd, a toolkit provided by Chaos Mesh, to enables users to simulate failures directly through physical nodes.
  * From the community’s perspective, Chaos Mesh now has documentation for different groups of community members, including [a user guide](https://chaos-mesh.org/docs/user_guides/installation/) and [developer guide](https://chaos-mesh.org/docs/development_guides/development_overview), [quick start guides](https://chaos-mesh.org/docs/get_started/get_started_on_kind), [use cases](https://chaos-mesh.org/docs/use_cases/multi_data_centers), and [a contributing guide](https://github.com/chaos-mesh/chaos-mesh/blob/master/CONTRIBUTING.md). The Chaos Mesh [Governance](https://github.com/chaos-mesh/chaos-mesh/blob/master/GOVERNANCE.md) was published in early 2021, which makes how to contribute to the project, participate in decision making, and how to earn merits within the community clearer. We have also established monthly [Chaos Mesh community meetings](https://github.com/chaos-mesh/chaos-mesh#community-meeting), where community members are invited to share the latest tech and community updates, as well as to exchange their experience of using and contributing to Chaos Mesh. We also maintained the community Slack channel under CNCF, which currently has 317 community members, where community support is timely provided, and other Chaos Mesh related issues and events discussed and shared. 
* **What are the current goals of the project**?  
  * Feature-wise, our short-term goal is to iteratively test and improve Chaos Mesh’s functions, user experience, and reduce costs for chaos engineering. The long-term goal is to build a chaos engineering ecology, including building a shared chaos experiment case platform, providing plug-ins to help users freely expand fault injection types and procedures, as well as to achieve more with observability.
  * We also would like to continuously improve the onboarding journey of users and contributors, which includes but is not limited to updating English documents and adding more Chinese documents, generating more use cases and chaos engineering practice guides to help users better conduct and implement chaos experiments.
* **How can the CNCF help you achieve your upcoming goals?**
    * Cloud Native chaos engineering is a relatively new, but an up and coming discipline. CNCF can best help us in promoting chaos engineering, so that more people become aware of the technology, which can drive the implementation of chaos engineering in organizations.
    * CNCF can also help Chaos Mesh by providing visibility to drive adoption and attract contributions, for example, by featuring Chaos Mesh in webinars, blogs, at community and industry events, or via other social media channels and resources.
    * The CNCF ecosystem is another driver for Chaos Mesh. Building collaborations with other CNCF projects to make Chaos Mesh a more wholesome project is another goal of ours CNCF can help drive collaborations with other projects under the CNCF landscape.
* **Do you think that your project meets the criteria for incubation?**
    * Yes, we do. We have already created the PR for it and presented our proposal in front of TAG Network, see [https://github.com/cncf/toc/pull/683](https://github.com/cncf/toc/pull/683). 


## Project Links
* [Website](https://chaos-mesh.org/)
* [GitHub](https://github.com/chaos-mesh/chaos-mesh)
* Slack:[ CNCF](http://slack.cncf.io) / #project-chaos-mesh
* [Twitter](https://twitter.com/chaos_mesh)
* [YouTube Channel](https://www.youtube.com/channel/UC4OwT4QTd0ML3YNnV1ybT6g)
* [Monthly community meeting](https://community.cncf.io/chaos-mesh-community/)
    * [ Meeting Minutes](https://docs.google.com/document/d/1H8IfmhIJiJ1ltg-XLjqR_P_RaMHUGrl1CzvHnKM_9Sc/edit#heading=h.t56mbwjhouvx)