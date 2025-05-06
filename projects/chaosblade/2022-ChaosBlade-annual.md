# ChaosBlade 2022 Annual Review

## Background

ChaosBlade is a versatile Chaos Engineering platform that orchestrates chaos experiments on Kubernetes and Host environments. 

By covering comprehensive fault injection methods in container, Pod, Node, network, file system, middleware, cloud vendor infrastructure, running application process, and even the kernel, ChaosBlade aims at providing a neutral, universal Chaos Engineering platform that enables cloud-native applications to be as resilient as they should be.

ChaosBlade contains the following main components:
![ChaosBlade Architecture](https://chaosblade.io/img/chaosblade-main-components.png)

- **ChaosBlade-Box Console**：The ChaosBlade visualization component mainly provides a set of user-friendly Web UI, through which users can arrange and manage chaos engineering experiments.
- **ChaosBlade-Box Server**：The core logic component is mainly responsible for the management and arrangement of chaos engineering experiments, probe and application management. Including components, Chaos Engine: exercise engine, including process orchestration, security control, exercise report and other functions; Chaos Runner: exercise executor, compatible with a variety of execution tools; Chaos Experience: exercise experience library, etc.
- **ChaosBlade-Agent**：The core logic component is deployed on the host of the user terminal or in the Kubernetes cluster. It is mainly responsible for establishing a connection with ChaosBlade-Box Server, reporting the heartbeat and serving as a command delivery channel.
- **ChaosBlade**：The main execution tool can perform fault injection on different environments such as the Host and Kubernetes, and can perform fault interference on system network devices, file systems, kernels, and applications running on the system.

**In Kubernetes environment**, ChaosBlade is a cloud-native chaos engineering framework. The project has “Chaos Experiment as a Custom Resource” as the fundamental design goal and keeps the community at the center for growing the chaos experiments.

- ChaosBlade CRDs: By defining Kubernetes CRD (chaosblades.chaosblade.io) to manage chaos experiments, each experiment has a very clear execution status. 
- ChaosBlade Operator: The operator is built using Operator SDK framework and manages the lifecycle of  a chaos experiment.
- ChaosBlade Tool:  The tool is deployed through the daemonset, used to receive commands issued by the operator, and perform drills on the target pod, container, and node.

**In Host environment**, Directly download the ChaosBlade binary executable file package to the target host, and through the CLI command for fault injection.
### Alignment with CNCF
**ChaosBlade was accepted into CNCF Sandbox on  April  27th, 2021**

- [Sandbox project onboarding](https://github.com/cncf/toc/issues/647)
- [Sandbox presentation slides](https://docs.google.com/presentation/d/18ch0AQsFp-he-rnr0sH8KhKH_gknd5Qi/edit#slide=id.p1)

## Development metrics
The ChaosBlade devstats page and dashboards can be found [here](https://chaosblade.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, ChaosBlade currently has [131](https://chaosblade.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [30](https://chaosblade.devstats.cncf.io/d/5/companies-table?orgId=1)
companies. On average, there are [54 commits per month](https://chaosblade.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) contained within [14 merged PRs per month](https://chaosblade.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now).
- [New PRs in last year](https://chaosblade.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&var-period=w&var-repogroup_name=All).
- The community has grown since the project entered the CNCF sandbox. 
   - We held bi-weekly community meetings constantly (total 15 as of Jun 2022). The meeting records can be found in [here](https://space.bilibili.com/22390414/video). The average number of  meeting attendees is ~30.
   - Number of contributors: 60+ -> **131**
   - Github stars: 3800+ -> **5100+**
   - Github forks: 600+ -> **850+**
   - Contributing organizations: 10+ -> **30**
## Alignment with Cloud Native Computing Foundation
ChaosBlade adopts a Kubernetes architecture that uses Kubernetes CustomResourceDefinitions (CRDs) to define chaos objects, making it naturally integrated with the Cloud Native ecosystem. 

By orchestrating chaos experiments with a comprehensive fault injection approach, ChaosBlade provides a vendor-neutral, scalable chaos engineering platform that makes cloud-native applications resilient, manageable, and observable.

ChaosBlade has the following integrations within the CNCF ecosystem:

- **Ansible:** ChaosBlade provides a way to install Agents and ChaosBlade tool [through ansible](https://github.com/chaosblade-io/chaosblade-box#pre-run-application-for-host). Facilitate automatic installation of agnet.
- **GitHub Action**: to enable users to run chaos experiments even during the development phase, the [ChaosBlade-action](https://github.com/chaosblade-io/chaosblade-action) project, allowing ChaosBlade to run in the workflow of GitHub Actions. This way, ChaosBlade can easily be integrated into daily system development and testing.
- **Prometheus**: The Box platform supports access to prometheus, so that it is convenient to observe the situation of the system being exercised, and it is used to verify the hypothesis of the chaos engineering experiment.
- **Kubevela**: [kubevela](https://github.com/kubevela/catalog/pull/471) make ChaosBlade as an addon for chaos engineering. This will allow chaosblade to pass through kubevela conveniently Installation and deployment in a multi-cluster/hybrid environment.
## Maintainers

We have established [the community membership roadmap](https://github.com/chaosblade-io/community/blob/main/Contributor_Ladder.md), and quite a few active and qualified contributors
have been promoted to maintainers. Compared to three maintainers (all from Alibaba) initially, we now have [four maintainers](https://github.com/chaosblade-io/chaosblade/blob/master/MAINTAINERS.md) from three organizations.

- [Changjun Xiao](mailto:changjun.xcj@alibaba-inc.com)(Alibaba)
- [Mingxia Cai](https://github.com/MandssS)(Alibaba)
- [Fei Ye](https://github.com/tiny-x)(Perfma)
- [Xudong Guo](https://github.com/sunny0826)(JiHu GitLab)
## Project adoption

The public adopters that used Chaosblade in production can be found [here](https://github.com/chaosblade-io/chaosblade/issues/32), about **50+**. Examples include:

- China Telecom
- IFLYTEK
- VIVO
- Shopee
- TP-Link
- Bigo
- ICBC
- Huobi
- Qunar
- Trip
- HUYA
- Bilibili
- DiDi
- China Mobile
- Shopee
- Huawei Cloud
- Tuhu
- H3C
- ......
## Project goals
Our primary goal  is to create a chaos engineering platform that supports multi-environment and multi-language, to help more enterprises or users realize the implementation of chaos engineering, so as to improve the high availability of their systems. 

In the past year, we have achieved the following key features:

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
   - Nginx chaos experiment
- Support for chaos on containerd runtime (along with existing docker support)
- Support for ARM64 nodes 
- Jvm performance optimization, reducing the impact on app processes when executing jvm chaos experiments.

In the near future, we aim to achieve the following goals:

- Combining ebpf technology to refine chaos experiment
- Enhance platform capabilities and increase observability
- Add support for cri-o container runtime
- Complete the chaos experiment on GCP/AWS/Azure
- Provide more document about the user guide and tutorial. Host or join more events to attract more contribution
-  A complete project roadmap can be found in [here](https://github.com/chaosblade-io/community/blob/main/roadmap.md)

## How the CNCF can help to achieve the upcoming goals

- More channels to advocate the project within CNCF and even to other foundations.
- Writing support for project documents.
## Incubation readiness

- Yes, we are preparing for the incubation proposal.
