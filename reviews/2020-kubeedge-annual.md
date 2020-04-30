# KubeEdge Sandbox Annual Review

## Table of Contents

- [Background](#background)
- [Alignment with Cloud Native](#alignment-with-cloud-native)  
- [Highlights of Last Year](#highlights-of-last-year)  
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

[KubeEdge](https://kubeedge.io/en/) is an open source system for extending native containerized application orchestration capabilities to hosts at Edge. It's built upon kubernetes and provides fundamental infrastructure support for network, application deployment and metadata synchronization between cloud and edge.

It aims to resolve three major challenges at edge computing: network reliability and bandwidth limit between cloud and edge, resource constraint at edge, highly distributed and large scalability requirement.

**KubeEdge provides:**
- Seamless Compute at Edge : low resource requirements;low latency; decision at edge; improve quality of edge service
- Simplified Development : regular http or mqtt based applications; containerize and run anywhere at Edge / Cloud
- Kubernetes Native Edge Platform : Only Kubernetes native Edge Platform in CNCF; Seamless workload orchestration
- Abundant Applications at Edge : Easy to get and deploy complex applications to edge (like image recognition, ML ..)

**KubeEdge was accepted as a CNCF Sandbox project on Mar 18, 2019.**
- [KubeEdge Sandbox Proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/kubeedge.adoc)
- [Original CNCF TOC meeting slides](https://docs.google.com/presentation/d/1e1ahun91Abn2xvX7Z8PVgGBz6c7Q7iO027XVzVuffDg/edit#slide=id.g25ca91f87f_0_0)

## Alignment with Cloud Native

KubeEdge falls in the scope of [CNCF Runtime SIG](https://github.com/cncf/sig-runtime).

**KubeEdge targets on:**
  - building an open edge computing platform with cloud native technologies
  - helping users extending their business architecture, applications, services, etc. from cloud to edge in same experience
  - implementing extensible architecture based on Kubernetes
  - integration with CNCF projects, including (but not limited to) containerd, cri-o, Prometheus, etc.
  - seamless development, deployment and run complex workloads at edge with optimized resources

## Highlights of Last Year
- Major, Stable releases (1.x.x ; currently 1.3.x) started
- Total 11 releases including major and minor
- 1400+ commits; 2400+ stars; 600+ forks; 300+ contributors
- 25+ Contributing organizations
- 5 User Adoptions and more under discussions
- Key Features Added : Device Management at Edge, Bluetooth Device Mapper, containerd support, QUIC as alternative to websocket, EdgeMesh, EdgeSite, CSI at Edge, AutoRegistration of Edge Nodes, Device CRD Validation, Reliable Message delivery to Cloud

## Annual Review Contents

* Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

  - [New PRs last 1 year](https://kubeedge.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1553961600000&to=now)
  <!--
    ![New PRs last 1 year](https://kubeedge.devstats.cncf.io/render/d-solo/15/new-prs-in-repository-groups?orgId=1&from=1553961600000&to=1588071511757&var-period=w&var-repogroup_name=All&panelId=1&width=1000&height=500&tz=Asia%2FShanghai)
  -->
  <!--
  - [Commits to main repo last 1 year](https://docs.google.com/spreadsheets/d/1sEHuUg2yyYOSf887mlFoZWNb6Y1dhk0xRyODxpiKi1M/edit?usp=sharing) (using data from git history, since commits data on devstats look not accurate)(using data from git history, since commits data on devstats look not accurate)
  -->

  - [Commits Repository Groups last 1 year](https://kubeedge.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=1551369600000&to=now&var-period=w&var-repogroups=All)
  <!--
    ![Commits Repository Groups last 1 year](https://kubeedge.devstats.cncf.io/render/d-solo/2/commits-repository-groups?orgId=1&from=1551369600000&to=1588072093433&var-period=w&var-repogroups=All&panelId=2&width=1000&height=500&tz=Asia%2FShanghai)
  -->

  - Community growth
    - Entered CNCF sandbox in March 2019
    - Continuous momentum
      - Num of Contributors: 30 => **300+**
      - Github Stars: 400+ => **2400+**
      - Github Forks: 100+ => **600+**
      - Contributing member organizations: 1 => **25+**
      - Contributing member organizations
        - __Cloud__: Huawei Cloud, DaoCloud, TenxCloud, EasyStack
        - __Carriers/Telco__: China Unicom, China Telecom, China Mobile
        - __IT__: inovex, Infoblox, SpanIdea
        - __IOT & Hardware__: ARM, Samsung, Midokura, EMQ
        - __Acadamic__: Zhijiang Lab, USTC, UESTC

* How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

  - [Initial maintainers](https://github.com/kubeedge/kubeedge/blob/master/MAINTAINERS)
  - New Approvers
    - Fei Xu - Huawei
    - Jie Zhang - China Unicom WoCloud
    - Dave Chen - ARM
    - Jan Unterbrink - inovex


* What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

  - Huawei Cloud
  - China Unicom WoCloud
  - Xinghai IoT (Case Study to be uploaded)
  - [Unimatrix.ai](https://gitlab.com/unimatrix)
  - [Akraino](https://wiki.akraino.org/display/AK/KubeEdge+Edge+Service+Blueprint)
  - More adopters under confirmation for making their information public


* How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

  - Development
    - Release Cadence: 1 minor release every three month
    - Moved to 1.0
    - Versioning Scheme: following [Semantic Versioning](https://semver.org/)
    - Key features added to the project:
      - Added v0.3:
        - Device Management on the Edge
        - Bluetooth DeviceMapper
        - KubeEdge Installer
        - Performance Test Framework

      - Added v1.0:
        - containerd
        - QUIC as alternative to websocket
        - Modbus DeviceMapper
        - EdgeMesh
        - EdgeSite

      - Added v1.1:
        - CSI suuport on the Edge
        - Device CRD validaton
        - L4 Proxy suuport in EdgeMesh

      - Added v1.2:
        - Application layer reliable message delivery from cloud to edge
        - KubeEdge Component Config Application
        - Auto registration of edge node

      - v1.3 (under development):
        - integrate CRI-O as Runtime
        - integrate Virtlet as Runtime
        - HA of CloudCore
        - Auto authentication for edge nodes
        - Collecting metrics of containers on the edge from Cloud
        - Collecting logs of containers on the edge from Cloud


* What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

  - Technical
    - More integration with CNCF and LF EDGE projects
    - More reliability
    - More Security
    - [Project Development Roadmap](https://github.com/kubeedge/kubeedge/blob/master/docs/getting-started/roadmap.md)

  - Community
    - Wider user adoption
    - Wider developer participation
    - More evangelism


* How can the CNCF help you achieve your upcoming goals?

  - More speaking opportunities

* Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)?
  - Yes

## Project Links
 - [Website](https://kubeedge.io)
 - [Github](https://github.com/kubeedge/kubeedge)
 - [Slack](https://kubeedge.slack.com)
 - [Mailing List]( https://groups.google.com/forum/#!forum/kubeedge )  
 - [Twitter](https://twitter.com/KubeEdge)
 - [Issue Tracker](https://github.com/kubeedge/kubeedge/issues)
