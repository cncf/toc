# KubeEdge Graduation Proposal

Since becoming a CNCF incubation-level project in September 2020, the KubeEdge project has made significant progress in developing contributors and users. In the past few years, the project also **helps the cloud native ecosystem to expand to more industries outside the data center**, and has been widely used in various new fields such as Smart CDN, Smart Campus, Intelligent Transpiration, Vehicle, Aerospace, Smart logistics, Financial, Blockchain and other fields. On behalf of KubeEdge team, we believe that KubeEdge is ready for graduation.

## Background

KubeEdge is an open source system for extending native containerized application orchestration capabilities to hosts at the edge. It's built upon Kubernetes and provides fundamental infrastructure support for networking, application deployment and metadata synchronization between cloud and edge. In addition, KubeEdge supports MQTT, allowing developers to design logic and enable resource constraints for device communications at the edge. KubeEdge decouples and simplifies Kubernetes modules.

![KubeEdge Architecture](https://raw.githubusercontent.com/kubeedge/kubeedge/26f47ccdb47b4fcd3c6eef9a7d97dffa49d619be/docs/images/kubeedge_arch.png)

It aims to resolve three major challenges for edge computing: 1) network reliability and bandwidth limit between cloud and edge, 2) resource constraint at edge, 3) highly distributed and heterogeneous device management.

KubeEdge provides:

1. **Seamless Cloud-Edge Communication** for both metadata and data
2. **Edge Autonomy**: Autonomous operation of Edge even when disconnected from cloud.
3. **Low Resource Requirement**: KubeEdge can work in constrained resource situations (low memory, low bandwidth, low compute)
4. **Simplified Device Communication**: Easy communication between applications and devices for IOT and Industrial Internet of Things (_IIoT_)
5. **Service discovery and traffic proxy** for edge application

In addition to these basic capabilities, the community is constantly exploring new directions for edge scenarios, including:

1. **Edge-cloud synergy AI**, Sedna subproject implement across edge-cloud collaborative training and collaborative inference capabilities, such as joint inference, incremental learning, federated learning, and lifelong learning. Sedna supports popular AI frameworks, such as TensorFlow, Pytorch, PaddlePaddle, MindSpore.
2. **Edge-Cloud synergy robots**, based on KubeEdge, to implement cloud robot platform, to integrate cloud computing technologies(cloud native, AI, storage etc) with robots, to support heterogeneous robots, including mobile robots, robotics arm, and many other robot platforms.

Link to GitHub repos:

- The KubeEdge Main project: <https://github.com/kubeedge/kubeedge>
- Subsidiary projects:
  - EdgeMesh: <https://github.com/kubeedge/edgemesh>
  - Mappers-go <https://github.com/kubeedge/mappers-go>
  - Sedna: <https://github.com/kubeedge/sedna>

## Alignment with Cloud Native

KubeEdge falls in the scope of [CNCF Runtime TAG](https://github.com/cncf/sig-runtime), and is also well-aligned with the CNCF's mission to make cloud native computing ubiquitous.

### Value-add to the CNCF ecosystem

As the CNCF's first cloud-native edge computing framework and incubation project, KubeEdge has helped **expanding cloud native ecosystem to more industries outside the data center**, including Smart CDN, Smart Campus, Intelligent Transpiration, Vehicle, Aerospace, Smart logistics, Financial, Blockchain and other fields.

- Building an open edge computing platform with cloud native technologies to **extend the boundary of cloud native**
- Helping users extending their business architecture, applications, services, etc. from cloud to edge in same experience
- Implementing extensible cloud-edge architecture based on Kubernetes
- Seamless development, deployment and execution of complex workloads at edge with optimized resources

### Alignment with other CNCF projects

As a CNCF project, KubeEdge integrates with the CNCF landscape and is built on top of the following CNCF projects and brings them to edge computing.

- Kubernetes: KubeEdge built on Kubernetes, compatible with native Kubernetes APIs and allows calls to these native APIs to manage edge nodes and devices.
- containerd, cri-o: KubeEdge supports run containers in edge node with containerd, cri-o though CRI.
- Prometheus: KubeEdge provides Prometheus metrics to monitor edge nodes and edge applications.
- WasmEdge: KubeEdge supports run lightweight wasm applications on edgenode with WasmEdge.
- KubeFlow: KubeEdge integrates with Kubeflow to run AI Jobs on edge nodes.
- ArgoCD: KubeEdge integrates with ArgoCD to provide GitOps on the edge.

In the future, we are planning to expand these integrations including edge service mesh with Istio and Envoy, edge storage with CSI related projects, edge serverless with Knative, OpenFunction, etc.

## Updates since becoming a CNCF Incubation project

### 1. Continuous Community Momentum

- Num of Contributors: 300 => 1400+
- Github Starts: 2500 => 6800+
- Github Forks: 600 => 1950+
- Contributing member organizations: 25 => 100+

### 2. Development & Stability & Security

Since 2018, the community has witnessed more than 50 releases. With excellent capabilities, KubeEdge has made technological breakthroughs along the way:

- Support for **100,000 edge nodes** in a Single Cluster
- **DMI**: cloud native edge device management framework
- **EdgeMesh**: HA edge computing architecture
- **Sedna**: edge-cloud synergy AI framework
- **Ianvs**: distributed collaborative AI benchmarking platform
- Support KubeEdge on Android
- White Paper on Cloud Native Edge Computing Threat Model and Security Protection Technical
- CNCF's **first** batch of integrated fuzzing projects
- CNCF's **first** SLSA L3 project

The figure below shows the key nodes in the evolution of the KubeEdge project：

![KubeEdge Innovative Journey](https://raw.githubusercontent.com/kevin-wangzefeng/kubeedge-community/kubeedge-graduation/graduation/kubeedge_innovative_journey.png)

### 3. Community

- Open Governance:
  - Governance doc
    - <https://github.com/kubeedge/community/blob/master/GOVERNANCE.md>
  - Established Technical Steering Committee:
    - **7 members** from **6 companies**, 1 Arm, 1 Google, 2 Huawei, 1 DaoCloud, 1 HarmonyCloud, 1 KubeSphere
    - <https://github.com/kubeedge/community/tree/master/committee-technical-steering>
  - Multiple SIGs established:
    - SIG AI
    - SIG Robotics
    - SIG Device IoT
    - SIG Node
    - SIG Networking
    - SIG Security
    - SIG Testing
    - SIG MEC
    - WG Wireless
  - Multiple Special Teams established:
    - Security Team
      - <https://github.com/kubeedge/community/tree/master/team-security>
    - WIP: PM Team

- Community Events:
  - KubeEdge Summit 2022
    - 2-days event, **1,000+ registrations**, **51,800** all-platform viewers
    - **50 speakers** from **35 orgs** (24 enterprises, 5 non-profits, 6 collages)
    - **28** community supporters
  - KubeEdge Acadamic Workshop 2022 (KEAW '22)
    - 2-days event, **800+ registrations**,**41,000+** all-platform viewers
    - **35 speakers** from **20 colleges**
  - KubeEdge Online Courses
    - 4 stages with 22 lessons, 30+ instructors
    - **1600+** learners have joined
    - **1.2k+** viewers per class, **140+** viewers per broadcast

- Remarkable Presentations at KubeCon
  - KubeCon NA 2022
    - KubeEdge: From Fixed Location To Movable Edge, Latest Updates And Future
  - KubeCon EU 2022 **Keynote**
    - Incremental Deep Learning For Satellite with KubeEdge and MindSpore
    - <https://sched.co/yuEH>
  - KubeCon China 2021 **Keynote**
    - Break Through Gravity! Cloud Native Towards Space With KubeEdge
    - <https://sched.co/qx5s>
  - KubeCon China 2020 **Keynote**
    - Managing One Hundred Thousand Edge Nodes With Kubernetes in China's Highway
    - <https://sched.co/cp8l>

## Wider adoption across industries

### End Users

KubeEdge has been widely used in industries such as intelligent transportation, smart campus, manufacturing, finance, aerospace, logistics, energy, and smart CDN, providing an integrated edge-device-cloud collaboration solution for users. The community has reached ongoing completions like the industry's largest cloud native highway project powered by edge-cloud synergy (100,000 edge nodes and 500,000 edge applications are centrally managed. Traffic at provincial borders is 13% smoother, and tolling efficiency is 10 times higher.) Other examples include the industry's first cloud native satellites, vehicles, and oilfields.

1. **[Cloud Native Vehicles (SAIC MAXUS)](https://www.bilibili.com/video/av771524351/)**: KubeEdge is used to connect vehicles to the cloud as nodes, implementing unified management of more than 100,000 vehicles. A single cluster manages millions of devices, facilitating unified device upgrade.
2. **[Smart Campus (XingHai IoT)](https://github.com/kubeedge/kubeedge/blob/master/ADOPTERS.md#xinghai-iot)**: Xinghai IoT is an IoT company that provides comprehensive smart building solutions by leveraging a construction IoT platform, intelligent hardware, and AI. Xinghai IoT now offers services in a wide range of areas. Therefore, its solutions should be portable and replicable and need to ensure real-time data processing and secure data storage. KubeEdge, with services designed for cloud native development and edge-cloud synergy, has become an indispensable part of Xinghai IoT for building smart campuses.
3. **[China's highways ETC system](https://vmblog.com/archive/2021/04/16/managing-100-000-edge-nodes-on-china-s-highways-using-kubeedge.aspx)**: Connected ETC systems and the removal of toll gates at provincial borders give rise to the deployment of a large number of toll gantries and controllers to collect vehicle information. Currently, the KubeEdge-based edge management system manages nearly 100,000 edge nodes and more than 500,000 edge applications in 29 provinces, municipalities, and autonomous regions. It supports the continuous adjustment and update of the highway gantry services, and collects more than 300 million data records every day.
4. **[Smart Customer Service System (China Mobile)]([https://www.cncf.io/blog/2021/08/16/china-mobile-kubeedge-based-customer-service-platform-featuring-edge-cloud-synergy/)**: Branches of China Mobile had headaches such as slow service response, low O&M efficiency, and low resource utilization. KubeEdge helps improve O&M efficiency by over 35%, and the service continuity of their 10086 customer service system reaches 99.999%.
5. **[Cloud Native Satellites](https://vmblog.com/archive/2022/05/06/incremental-deep-learning-for-satellite-with-kubeedge-and-mindspore.aspx)** ("Sky Computing Constellation" jointly developed with Beijing University of Posts and Telecommunications): Integrates KubeEdge with satellite computing to enable satellites to use resources of edge and central clouds simultaneously in space and have AI and multitasking capabilities. By means of collaborative inference between satellites and ground stations, the calculation precision is improved by more than 50%, and the amount of data sent back by satellites to the ground is reduced by 90%. The distributed AI provided by Sedna improves the on-orbit image inference and image recognition accuracy.
6. **[Smart factory (Raisecom)](https://github.com/kubeedge/kubeedge/blob/master/ADOPTERS.md#raisecom-technology-coltd)**: Ensuring industrial production safety is an important requirement of Raisecom. With KubeEdge as the edge computing platform, factory safety is under perfect control using AI, reducing the occurrence of accidents and improving the production efficiency.
7. **[Mine Brain (Jingying Shuzhi)](https://github.com/kubeedge/kubeedge/blob/master/ADOPTERS.md#jingying-shuzhi-technology-co-ltd)**: "Mine Brain" has helped many mining enterprises in Shanxi to realize smart mining in more than 1,000 mines: unified management, batch delivery from the center, and remote upgrade; AI analysis algorithms to deliver instructions from the cloud to the mines and provide real-time risk warning. The identification rate has reached 98%. Centralized monitoring of the IT infrastructure of sub-mining sites reduces O&M costs by 65%. Integrated full-stack IT devices in a cabinet reduce deployment costs by 75%.
8. **[IoT System (SF Express)](https://www.bilibili.com/video/BV1Gg41117d8/?spm_id_from=333.999.0.0)**: Industrial IoT is seeing explosive device connections and data volume. Based on KubeEdge, SF Express built ubiquitous IoT applications for their logistics and supply chain systems, and connected heterogeneous devices from multiple sources by direct connection or cross-level connection through edge nodes. The company effectively streamlined onsite OT production systems. Digital monitoring and intelligent management help improve business operations.
9. **[FinTech Smart Bank (China Construction Bank)](https://www.bilibili.com/video/BV17V4y1E7bZ/?spm\_id\_from=333.999.0.0)**: Banks have high demands on service consistency, system stability, and security. Their IT systems must be scalable and agile on the basis of being highly available and reliable. KubeEdge helps CCB Fintech realize platform-based and cloud-based edge computing to make everything intelligent.
10. **[Large-scale CDN (China Telecom eCloud)](https://mp.weixin.qq.com/s/bPq9NpR6dO8wBvHCoj5RA)**: As a typical product with cloud-network synergy, CDN needs to manage a large number of widely distributed edge nodes. Its core components are complex to be reconstructed, associated components need to be deployed in collaboration, and applications need to be managed in a unified manner to reduce costs and improve efficiency. By June 22, all new and legacy nodes of eCloud have been moved to cloud based on KubeEdge, and they now have a full CDN product line that supports multiple acceleration services for live streaming, VOD, downloads, security, services outside Chinese mainland, as well as whole site and static acceleration, backed by industry-leading precise scheduling.


### Integrators

- **[HuaweiCloud Intelligent EdgeFabric (IEF)](https://www.huaweicloud.com/product/ief.html):** provides you a complete edge computing solution based on KubeEdge, in which cloud applications are extended to the edge. By leveraging edge-cloud synergy, you can manage edge nodes and applications remotely and process data nearby, to meet your requirements for remote control, data processing, analysis, decision-making, and intelligence of edge computing resources. In addition, you can perform O&M in the cloud, including edge node monitoring, application monitoring, and log collection.
- **[DaoCloud Edge Computing (DEC)](https://www.daocloud.io/en/dec):** based on KubeEdge, extends the boundary of cloud native capability to the edge. It integrates the end, edge, and cloud to help enterprises quickly build a safe, flexible, efficient and reliable edge cloud native platform. DEC is widely applicable to industrial Internet, IoT, smart factories, smart parks, smart retailing and other fields. The cloud side registers and manages edge nodes and deploys relevant applications and configurations. The edge side runs edge applications and implements edge autonomy. The end side integrates terminals with multiple protocols and provides standard interface docking equipment capabilities.
- **[HarmonyCloud - EdgeStack](https://harmonycloud.cn/product/2):** builds a cloud-edge collaboration system based on KubeEdge, support access to multiple edge nodes, provide capabilities such as computing power sinking, access to multiple device protocols, and nearby data processing, and accelerate the implementation of edge intelligence for users.
- **[BoCloud - BeyondEdge](https://www.bocloud.com.cn/product/BeyondEdge/):** realizes the bilateral collaboration among cloud, edge and device end with open architecture. The cloud is responsible for edge system registration and management, application, and configuration distribution; the edge for edge application running and access device management; the device end for edge devices running. They work independently or jointly for different scenarios.
- **[KubeSphere](https://www.kubesphere.io/docs/v3.4/installing-on-linux/cluster-operation/add-edge-nodes/):** KubeSphere leverages KubeEdge, to extend native containerized application orchestration capabilities to hosts at edge. With separate cloud and edge core modules, KubeEdge provides complete edge computing solutions.
- **[EMQ - Analytic Engine](https://ekuiper.org/docs/en/latest/integrations/kubeedge/overview.html)** based on KubeEdge, eKuiper can ingest device data from MQTT and provides versatile analytic capabilities for the KubeEdge components to achieve low latency calculation on edge.
- **[Orange Cloud Intelligent EdgeFabric (IEF)](https://cloud.orange-business.com/en/offers/infrastructure-iaas/public-cloud/features/intelligent-edge-fabric/):** provides users with a complete edge computing solution where cloud applications are extended to the edge. By leveraging edge-cloud synergy, users can manage edge nodes and applications remotely while still processing data nearby.
- **[Envision Digital - ENOS Edge](https://www.envisioniot.com/docs/enos-edge/en/2.2.0/howto/console\_configuration/containerized\_service.html):** Based on Kubernetes and KubeEdge, EnOS Edge Container System can manage massive Edge nodes in batches.
- **[Ci4Rail GmbH - EdgeFarm IoT Platform](https://www.ci4rail.com/edgefarm-iot-platform/):** EdgeFarm focuses on the lossless and high-quality acquisition, transmission and processing of all types of field data and thus provides the basis for intelligent maintenance or fleet optimization, regardless of the vehicle type and manufacturer. In addition, EdgeFarm enables the convenient management of user's devices as well as user's applications and thus creates transparency for user's fleet.

## Graduation State Criteria

### 1. Have committers from at least two organizations.

KubeEdge community has 7 Technical Steering Committees from 6 companies, who are the governing body for the KubeEdge project, providing decision-making and oversight pertaining to the KubeEdge project bylaws. The Technical Steering Committee also defines the project values and structure.

KubeEdge has 10 core maintainers from 7 companies, who are responsible for maintaining code development for overall projects. And KubeEdge also has 15+ SIG maintainers who Focus on maintaining corresponding modules.

**Maintainers list**: <https://github.com/kubeedge/kubeedge/blob/master/MAINTAINERS.md>

### 2. Have achieved and maintained a Core Infrastructure Initiative [**Best Practices Badge**](https://bestpractices.coreinfrastructure.org/)

Yes. See links below for Core Infrastructure Initiative Best Practices Badge.

Ref: <https://bestpractices.coreinfrastructure.org/en/projects/3018>

### 3. Have completed an independent and third party security audit

KubeEdge community in 2022 cooperates with Ada Logics Ltd. and The Open Source Technology Improvement Fund performed a holistic security audit of KubeEdge and output a **security auditing report**, including the security threat model and security issues related to the KubeEdge project, which has important guiding significance for the security protection of the KubeEdge project.

The discovered security issues have been fixed and patched to the releases(v1.11.1, v1.10.2, v1.9.4) by KubeEdge maintainers according to the [kubeedge security policy](https://github.com/kubeedge/kubeedge/security/policy).

**KubeEdge Security Audit report**: <https://github.com/kubeedge/community/blob/master/sig-security/sig-security-audit/KubeEdge-security-audit-2022.pdf>

**KubeEdge Threat Model and Security Protection Analysis**: <https://github.com/kubeedge/community/blob/master/sig-security/sig-security-audit/KubeEdge-threat-model-and-security-protection-analysis.md>

**KubeEdge Security Policy**: <https://github.com/kubeedge/community/blob/master/team-security/SECURITY.md>

### 4. Explicitly define a project governance and committer process

This is documented in our governance documentation available here as well as our contribution guide for making contributions.

**Community Governance doc**: <https://github.com/kubeedge/community/blob/master/GOVERNANCE.md>

**Contribution Guide doc**: <https://github.com/kubeedge/community/blob/master/contribute.md>

### 5. Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project

This is documented in our community membership documentation available here for roles with the requirements and responsibilities associated with them.

**Community Membership doc**: <https://github.com/kubeedge/community/blob/master/community-membership.md>

### 6. Have a public list of project adopters for at least the primary repo

KubeEdge provides a full overview of publicly listed end-users on the website, but there are others that are not yet published or in the application process to list but can be interviewed.

**KubeEdge Supporters**: <https://kubeedge.io/#supporters>

## Appendix

### Adopter Interviews

The following adopter interviews have been conducted:
- [x] [China Telecom Cloud](https://www.chinatelecomglobal.com/): The TOC sponsors have interviewed Qidong Zhang from China Telecom Cloud in June 2024, refer here for the [interview report](https://docs.google.com/document/d/1UXuPK9bXqCZzYXLf1WdKXSgiqR8oiXrSkzxT0g6bY6c/edit?usp=sharing).
- [x] [Nio](https://www.nio.com/): The TOC sponsors have interviewed Xuhui Jiang from Nio in July 2024, refer here for the [interview report](https://docs.google.com/document/d/1TgIyzYG8BL5K_osHuJfKuj2YCYdrI766OCSAL_-UUvw/edit?usp=sharing).
- [x] [Zeblok Computational Inc](https://www.zeblok.com/): The TOC sponsors have interviewed Vikarna Tathe from Zeblok Computational Inc in July 2024, refer here for the [interview report](https://docs.google.com/document/d/186uAKT6suU7UR8bbpkiNxEA2nrMZFabDuKynErnjWPM/edit?usp=sharing).
