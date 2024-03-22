ChaosBlade was accepted into the CNCF Sandbox on April  27th, 2021. This doc proposes promoting ChaosBlade to a CNCF Incubating project.
## About ChaosBlade
ChaosBlade is a versatile Chaos Engineering platform that orchestrates chaos experiments on Kubernetes and Host environments. 

By covering comprehensive fault injection methods in container, Pod, Node, network, file system, middleware, cloud vendor infrastructure, running application process, and even the kernel, ChaosBlade aims at providing a neutral, universal Chaos Engineering platform that enables cloud-native applications to be as resilient as they should be.

ChaosBlade contains the following main components:

- **ChaosBlade-Box Console**：The ChaosBlade visualization component mainly provides a set of user-friendly Web UI, through which users can arrange and manage chaos engineering experiments.
- **ChaosBlade-Box Server**：The core logic component is mainly responsible for the management and arrangement of chaos engineering experiments, probe and application management. Including components, Chaos Engine: exercise engine, including process orchestration, security control, exercise report and other functions; Chaos Runner: exercise executor, compatible with a variety of execution tools; Chaos Experinece: exercise experience library, etc.
- **ChaosBlade-Agent**：The core logic component is deployed on the host of the user terminal or in the Kubernetes cluster. It is mainly responsible for establishing a connection with ChaosBlade-Box Server, reporting the heartbeat and serving as a command delivery channel.
- **ChaosBlade**：The main execution tool can perform fault injection on different environments such as the Host and Kubernetes, and can perform fault interference on system network devices, file systems, kernels, and applications running on the system.

**In Kubernetes environment**, ChaosBlade is a cloud-native chaos engineering framework. The project has “Chaos Experiment as a Custom Resource” as the fundamental design goal and keeps the community at the center for growing the chaos experiments.

- ChaosBlade CRDs: By defining Kubernetes CRD (chaosblades.chaosblade.io) to manage chaos experiments, each experiment has a very clear execution status. 
- ChaosBlade Operator: The operator is built using Operator SDK framework and manages the lifecycle of  a chaos experiment.
- ChaosBlade Tool:  The tool is deployed through the daemonset, used to receive commands issued by the operator, and perform drills on the target pod, container, and node.

**In Host environment**, Directly download the ChaosBlade binary executable file package to the target host, and through the CLI command for fault injection.
## ChaosBlade in the CNCF Sandbox
### Community & Growth
The ChaosBlade devstats page and dashboards can be found [here](https://chaosblade.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open). Under CNCF we have seen significant growth,The community has grown since the project entered the CNCF sandbox:

- We held bi-weekly community meetings constantly (total 28 as of Jun 2024). The meeting records can be found in [here](https://space.bilibili.com/22390414/video). The average number of  meeting attendees is ~30.
- Number of contributors: 60+ -> **149**
- Github stars: 3800+ -> **5400+**
- Github forks: 600+ -> **900+**
- Contributing organizations: 10+ -> **31**
- maintainer: 4 ->** 5**
-  Working with the community to generate content: we have published 20 blog posts on the [ChaosBlade blog](https://chaosblade.io/blog/). 
### Governance

The project matured remarkably. We have established [the Community Contributor Ladder](https://github.com/chaosblade-io/community/blob/main/Contributor_Ladder.md), and quite a few active and qualified contributors
have been promoted to maintainers. Currently, we have 5 maintainers from four different organizations.
### Technology
ChaosBlade was accepted into the CNCF Sandbox on April  27th, 2021. Since that date, the maintainers have added substantial new features:

- Chaos Dashboard (box platform) has been fully upgraded to provide a [more friendly interactive interface](https://chaosblade.io/blog/2022/06/24/ChaosBlade-Box-a-New-Version-of-the-Chaos-Engineering-Platform-Has-Released)
- Chaos Dashboard(box platform) support internationalized, [Chinese and English switching ](https://www.alibabacloud.com/blog/chaosblade-box-a-new-version-of-the-chaos-engineering-platform-has-released_599069)
- The open source agnet is used as a command channel, and it can also automatically report pod data in a cloud-native environment, which is convenient for k8s cluster to execute chaos experiments.
- Ansible can help users install agent.
- Promentheus can help users monitor chaos experiments.
- New types of fault injection:
   - Time travel chaos experiment
   - JVM Add plugin zookeeper and clickhouse
   - Golang chaos experiment
   - Alicloud chaos experiment
   - AWS chaos experiment
   - Nginx chaos experiment
   - Redis chaos experiment
- Support for chaos on containerd runtime (along with existing docker support)
- Support for ARM64 nodes 
- Jvm performance optimization, reducing the impact on app processes when executing jvm chaos experiments.
### Collaboration with other CNCF projects

- **Ansible:** ChaosBlade provides a way to install Agents and ChaosBlade tool [through ansible](https://github.com/chaosblade-io/chaosblade-box#pre-run-application-for-host). Facilitate automatic installation of agnet.
- **GitHub Action**: to enable users to run chaos experiments even during the development phase, the [ChaosBlade-action](https://github.com/chaosblade-io/chaosblade-action) project, allowing ChaosBlade to run in the workflow of GitHub Actions. This way, ChaosBlade can easily be integrated into daily system development and testing.
- **Prometheus**: The Box platform supports access to prometheus, so that it is convenient to observe the situation of the system being exercised, and it is used to verify the hypothesis of the chaos engineering experiment.
- **Kubevela**: [kubevela](https://github.com/kubevela/catalog/pull/471) make ChaosBlade as an addon for chaos engineering. This will allow chaosblade to pass through kubevela conveniently Installation and deployment in a multi-cluster/hybrid environment.
## Comparison with other CNCF Chaos Engineering projects

ChaosBlade, Chaos Mesh, and Litmus Chaos are chaos engineering tools designed for Kubernetes environments and are used to conduct fault injection experiments in distributed systems to test the stability and reliability of the system. Although they have some overlap in goals and application scenarios, ChaosBlade has some unique advantages. For example:

- **Lightweight**: ChaosBlade is designed to be lightweight and easy to use, and provides a simple command line interface to perform fault injection operations.
- **Modularity**: Supports fault simulation in multiple scenarios, such as network delay, increased CPU load, disk I/O failure, etc., and can be expanded according to **different platforms** (such as Kubernetes, applications, host level). Aboult 300+ fault experiment.
- **Cross-platform**: Not only supports Kubernetes, but also supports chaos experiments directly on physical or virtual machines. Supports multiple operating systems, such as **linux, windows**.
- Support fault injection for application processes in **different languages**, such as **java, c++, goland, nodejs **etc.
- Supports fault injection for **multiple cloud vendors**, such as **AWS, Alibaba Cloud**. Support fault injection for **middleware**, such as** nginx, redis.**
- ChaosBlade enables users to design and orchestrate complex chaos experiments, rather than a mere fault injection chaos tool. The platform directly controls physical and k8s nodes, thereby directly injecting faults, making it convenient for users and reducing the cost of chaos engineering experiments.

Although having conceptual overlaps with Kubeedge and Superedge, ChaosBlade is driven by distinct technical paths for emerging new challenges.















