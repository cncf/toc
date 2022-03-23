# Chaos Mesh Incubation Stage Proposal 

Chaos Mesh was accepted into the CNCF Sandbox on July 15th, 2020. This doc proposes promoting Chaos Mesh to a CNCF Incubating project.

* **Name of Project:** Chaos Mesh  
* **Unique Identifier**: chaos-mesh
* **License**: Apache License, Version 2.0
* **[GitHub]**: https://github.com/pingcap/chaos-mesh
* **Website**: [https://github.com/pingcap/chaos-mesh](https://github.com/pingcap/chaos-mesh)
* **Sponsor from TOC**: TBD
* **Preferred Maturity Level**: Incubating
* **Sandbox References**
  *  [Sandbox proposal PR](https://github.com/cncf/toc/pull/367)
  * [Sandbox presentation slides](https://docs.google.com/presentation/d/115QvSCnT6ROwwV1lK-R10d4MW54fQhtVM2uZDGscpEk/edit#slide=id.g80f082b201_0_0)

## Project description

Chaos Mesh is a versatile Chaos Engineering platform that orchestrates chaos experiments on Kubernetes environments. By covering comprehensive fault injection methods in container, Pod, network, file system, and even the kernel, Chaos Mesh aims at providing a neutral, universal Chaos Engineering platform that enables cloud-native applications to be as resilient as they should be. 

Chaos Mesh contains the following main components:

*   **Chaos Dashboard**: a Web UI for managing, designing, monitoring chaos experiments, through which users can operate and observe chaos experiments. It also provides RBAC authorization.
*   **Chaos Controller Manager**: the core logic component of Chaos Mesh. It’s responsible for the scheduling and managing chaos experiments. The Controller Manager contains multiple CRD Controllers, such as Workflow Controller, Scheduler Controller, and various types of fault controllers.
*   **Chaos Daemon**: the main executive component of Chaos Mesh. Chaos Daemon runs in [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) mode and has privileged permissions by default (can be manually turned off). Chaos Daemon mainly interferes with specific network devices, file systems, kernels, etc. by intruding into the target Pod Namespace.

# Alignment with Cloud Native

Chaos Mesh falls into the scope of [CNCF TAG Network](https://github.com/cncf/tag-network).

Chaos Mesh adopts a Kubernetes architecture that uses Kubernetes CustomResourceDefinitions (CRDs) to define chaos objects, making it naturally integrated with the Cloud Native ecosystem.

By orchestrating chaos experiments via comprehensive fault injection methods, Chaos Mesh provides a vendor-neutral, scalable chaos engineering platform that enables cloud-native applications to be **_resilient, manageable, and observable_**. 

Chaos Mesh has the following integrations within the CNCF ecosystem:

*   **Grafana**: a separate [Grafana plug-in](https://github.com/chaos-mesh/chaos-mesh-datasource) is included, which allows users to directly display real-time chaos experiment information on the application monitoring panel, which improves the observability of chaos experiments.
*   **GitHub Action**: to enable users to run chaos experiments even during the development phase, the [chaos-mesh-action](https://github.com/chaos-mesh/chaos-mesh-action) project, allowing Chaos Mesh to run in the workflow of GitHub Actions. This way, Chaos Mesh can easily be integrated into daily system development and testing.
*   **Argo: **[TiPocket](https://github.com/pingcap/tipocket) provides an automated test platform that enables a fully automated chaos engineering testing loop for a distributed database, which includes deploying applications, running workloads, injecting exceptions, and business checks.
*   **Civo:** Chaos Mesh has been made [available](https://github.com/civo/kubernetes-marketplace/tree/master/chaos-mesh) on [Civo](https://github.com/civo)’s Kubernetes Marketplace, where Kubernetes users can deploy Chaos Mesh and start conducting chaos engineering conveniently.

## Progress since joining CNCF Sandbox

### Key features added
*   Chaos Dashboard as a visual aid to help users manage and monitor chaos experiments through a Web UI.
*   Workflow orchestration support for users to arrange chaos experiments on their own.
*   Interrupt mechanism and automatic detection of application status steps.
*   Authorization mechanism based on Kubernetes RBAC permission policies.
*   New types of fault injection:
    *   StressChaos
    *   DNSChaos
    *   JVMChaos
    *   AWSChaos
    *   HTTPChaos
*   Provides complete native event information.
*   Grafana plug-in to help users monitor chaos experiments.
*   A toolkit, Chaosd, so that users can simulate failures directly through physical nodes.


### Governance

Chaos Mesh published its [Governance ](https://github.com/chaos-mesh/chaos-mesh/blob/master/GOVERNANCE.md)in the beginning of 2021, which makes how to contribute to the project, participate in decision making, and how to earn merits within the community much clear for community members.

### Community growth

Chaos Mesh is a thriving project with an active community. One year under CNCF we have seen significant growth:

*   Stars: 3.7k (from 1.5k)
*   Contributors: 100+ (from 23)
*   Commits: 1,000+ (from 240)
*   Forks: 420+ (from 88)
*   Community Slack: 315+

What’s more, The quality of an open source community determines whether the community members are willing to stick around and get involved in the long run. In this regard, we’ve been working hard on:

*   Continuously enriching documentation and optimizing its structure. So far, we have developed a complete set of documentation for different groups of audiences, including [a user guide](https://chaos-mesh.org/docs/user_guides/installation/) and [developer guide](https://chaos-mesh.org/docs/development_guides/development_overview), [quick start guides](https://chaos-mesh.org/docs/get_started/get_started_on_kind), [use cases](https://chaos-mesh.org/docs/use_cases/multi_data_centers), and [a contributing guide](https://github.com/chaos-mesh/chaos-mesh/blob/master/CONTRIBUTING.md). All are constantly updated per each release.
*   Working with the community to generate content: in the past year, we have published 11 blog posts on the [Chaos Mesh blog](https://chaos-mesh.org/blog/).
*   There have also been 13 [community generated articles](https://github.com/chaos-mesh/chaos-mesh#community-blogs) on various platforms, such as the CNCF blog, Medium, opensource.com, etc. Among them is [an interactive tutorial](https://chaos-mesh.org/interactive-tutorial), published on O’Reilly’s Katakoda site. 
*   Repurposing and amplifying [videos and tutorials](https://github.com/chaos-mesh/chaos-mesh#community-talks) generated in community meetings, live streams, conferences, webinars, and meetups on different social media channels.


## Comparison with other CNCF Chaos Engineering projects

As compared to other Cloud Native chaos engineering tools, Chaos Mesh has the following advantages: 

*   Our goal is to build a **chaos engineering ecology**, so that users of Chaos Mesh can conduct chaos experiments from start to finish on a single platform. 
*   Chaos Mesh uses Kubernetes CustomResourceDefinitions (CRDs) to manage chaos experiments, is easy to scale, and is closely integrated with the Kubernetes ecosystem, enabling users to use Chaos Mesh in combination with other tools within the Kubernetes ecosystem easily.
*   **Thanks to its native orchestration engine, users do not need to install other components to perform chaos experiments, reducing component dependencies and risks.**
*   Chaos Mesh has an easy-to-use UI (Chaos Dashboard) for chaos experiments, which reduces the cost of chaos experiments.
*   Chaos Mesh currently provides physical (VM) and Kubernetes fault injection methods, respectively.
*   Faults injected through Chaos Mesh are not intrusive to the applications, which means users do not need to make any adjustments, and can use Chaos Mesh directly in the production environment.

In comparison to ChaosBlade, Chaos Mesh embraces cloud-native, is more closely integrated with the Kubernetes ecosystem and aims to build a **chaos engineering ecology**. Chaos Mesh enables users to design and orchestrate complex chaos experiments, rather than a mere fault injection chaos tool.

When compared to LitmusChaos, Chaos Mesh puts more focus on reducing the cost of chaos engineering. Chaos Mesh provides native orchestration, meaning that users do not need to rely on external tools for conducting chaos experiments, whereas Litmus users have to rely on other tools for Workflow. What’s more, Chaos Mesh has a Chaos Engineering toolkit, Chaosd, which can be used to inject failures into physical nodes, providing more possibilities for chaos tests.


## Incubation requirements

**Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.**

This is a [list of Chaos Mesh users](https://github.com/chaos-mesh/chaos-mesh/blob/master/ADOPTERS.md) that have gone into production. Below are some user cases to highlight:

* [Apache APISIX](https://github.com/apache/apisix) \
Apache APISIX is a dynamic, real-time, high-performance open source API gateway, which provides rich traffic management features, such as load balancing, dynamic upstream and more. \
APISIX integrates Chaos Mesh into open-source CI pipelines, to further enhance APISIX's resiliency and robustness.

 * [ByteDance](https://bytedance.com/en/) \
ByteDance is a technology company operating a range of content platforms including TikTok, Lark, Helo, Vigo Video, Douyin, and Huoshan in over 150 markets and 75 languages. \
ByteDance's self-developed chaos engineering platform is mainly used by the company's own technology system. As there are some cloud-native deployment services involved, they integrated Chaos Mesh as the underlying fault injection engine, which is a key supplement to ByteDance’s chaos engineering platform.

*  [NetEase Fuxi Lab](https://fuxi.163.com/en/about.html) 
  
    NetEase Fuxi AI Lab is China’s first professional game AI research institution. Researchers use their Kubernetes-based Danlu platform for algorithm development, training and tuning, and online publishing. 

    They use Chaos Mesh to improve the stability of their internal hybrid cloud. In addition, their users with cloud platforms also access Chaos Mesh to test the stability of user services. See more in their [Use case](https://chaos-mesh.org/blog/how-a-top-game-company-uses-chaos-engineering-to-improve-testing).

* [KingNet](https://www.kingnet.com/) 

    KingNet’s main business includes the development, operation and distribution of premium entertainment content. 

    KingNet mainly uses Chaos Mesh for testing the availability of multiple data centers and microservice links. Chaos Mesh also helps them with mocking service unavailability and abnormal network conditions.

*   [Qiniu Cloud](https://qiniu.com/en)

    Qiniu Cloud is a distributed cloud system that carries massive amounts of data, and is one that requires high data consistency and high availability, with the data quantity level at 1 trillion+.  

    To ensure the reliability of cloud storage products, they use Chaos Mesh to perform chaos tests on metadata and the underlying storage system under conditions such as: single point of failure of services, network abnormality, abnormal resource consumption (CPU, memory, I/O), etc.
*   [Tencent](https://www.tencent.com/en-us) \
After Tencent Interactive Entertainment migrated their online operations to the Tencent Cloud Kubernetes engine, they wished to provide users with a more stable and reliable experience, which is why they introduced Chaos Mesh. Tencent mainly use Chaos Mesh to simulate the following types of failures:
    * Fault isolation, such as simulating pod abnormality, and checking whether the system can automatically isolate fault instances;
    * Service degradation, such as simulating a downstream recommended service failure through network failure, and verifying whether the local cache is effective;
    * Verifying if the alarm works, for example, purposefully burning the CPU to 90%, and checking whether the alarm is timely issued in time.

*   [Xpeng](https://en.xiaopeng.com/) 

    Xpeng Motors is China's leading smart electric vehicle designer and manufacturer, as well as a technology company integrating cutting-edge Internet and AI innovation. They use Chaos Mesh in the following scenarios: 

    * Rolling updates of microservices and lossless verification of traffic;
    * Microservices, multi-registries, multi-party synchronization, and traffic lossless verification;
    * mqtt cluster two-way subscription verification;
    * Exactly-once consumer business verification for message queues;
    * Simulation of weak 4G network for in-vehicle systems, saving drive test costs;
    * AIOPS anomaly detection dataset generation.

**Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.**

Chaos Mesh currently has[8 Maintainers](https://github.com/chaos-mesh/chaos-mesh/blob/master/MAINTAINERS.md#project-maintainers) from 3 organizations and 8 Committers from 3 organizations. The roles and the corresponding responsibilities are defined in the [Project Governance](https://github.com/chaos-mesh/chaos-mesh/blob/master/GOVERNANCE.md). 

**Demonstrate a substantial ongoing flow of commits and merged contributions**

Chaos Mesh currently has over 100 contributors from [45+ companies](https://chaosmesh.devstats.cncf.io/d/5/companies-table?orgId=1). According to devstats, the project has on average [193 commits per month](https://chaosmesh.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-1y&to=now-2d&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all), along with [93 merged PRs per month](https://chaosmesh.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-1y&to=now-2d&var-period=m&var-metric=prs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all).

**A clear versioning scheme**

All Chaos Mesh repositories follow [Semantic Versioning 2.0.0 convention](https://semver.org/). So far we have 33 releases, with Chaos Mesh 1.0 released on Sep 25, 2020, and the current release is v2.0.0-beta.2.

The release cadence are as follows:

*   Major version is released every 6 months
*   Minor version is released every 2 months.
*   Patch version is released as necessary.

The releases are announced using [GitHub’s releases](https://github.com/chaos-mesh/chaos-mesh/releases), followed by Slack (CNCF/#project-chaos-mesh) and social media ([Twitter](https://twitter.com/chaos_mesh)) announcements.

