# Volcano Sandbox Annual Review

## Table of Contents

- [Background](#background)
- [Alignment with Cloud Native](#alignment-with-cloud-native)
- [Year in Review](#year-in-review)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

[Volcano](https://volcano.sh/en/) is a system for running high-performance workloads on Kubernetes. It features powerful batch scheduling capability that Kubernetes cannot provide but is commonly required by many classes of high-performance workloads, including Machine Learning, Deep Learning, Big Data, Bioinformatics Computing, etc. These types of workloads typically run on generalized domain frameworks like TensorFlow, Spark, PyTorch, MPI, etc. Volcano is integrated with these frameworks to allow users to run their applications without extra adaptation efforts while enjoying remarkable batch scheduling.


**Volcano provides:**
1. Powerful support to mainstream computing architectures such as Spark/Flink/Tensorflow/PyTorch.
2. Rich advanced scheduling strategies including Gang Scheduling/Backfill/Bin-pack, etc.
3. Support mixed scheduling of heterogeneous resources, e.g. CPU, Memory, GPU, etc.


**Volcano was accepted as a CNCF Sandbox project on Apr 10, 2020.**
- [Volcano Sandbox Proposal](https://github.com/cncf/toc/blob/main/proposals/volcano.md)
- [Original CNCF TOC meeting slides](https://docs.google.com/presentation/d/1RdplRxmUicu0Y03VoMzap3Fb4amwjvMdLH9K7QxOFiM/edit?usp=sharing)

## Alignment with Cloud Native

Volcano falls in the scope of [CNCF Runtime SIG](https://github.com/cncf/sig-runtime).


**Volcano targets on:**
  - leveraging high performance (especially AI/Big Data/HPC) workloads, to run on Kubernetes
  - providing unified API for all computing architectures
  - providing rich and scalable scheduling strategies framework
  - supporting heterogeneous devices such as GPU/FPGA

## Year in Review
- Moved to 1.3, latest version v1.3.0
- 1759+ commits; 1.8k+ stars; 400+ forks; 220+ contributors
- 30+ Contributing organizations
- ~20 Public User Adoptions including 10 production grade.
- Key Features Added: added monitoring components/resource reservation support/HDRF support/added TDM plugin/added SLA plugin/task-topology support.

## Annual Review Contents

* Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

  - [New PRs last 1 year](https://volcano.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1586448000000&to=1622476799000&var-period=w&var-repogroup_name=All)

  - [Commits Repository Groups last 1 year](https://volcano.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=w&var-repogroups=All)

  - Community growth
    - Entered CNCF sandbox in April 2020
    - Continuous momentum
      - Num of Contributors: 60 => **220+**
      - Github Stars: 800+ => **1800+**
      - Github Forks: 190+ => **400+**
      - Contributing member organizations: 14 => **30+**
      - Contributing organizations including:
        - Huawei, Amazon, Tecent, Baidu, ByteDance, NetEase,   
        ZTE, HP, Inspur, IBM, ARM, Oracle, Visa

* How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

  Initial maintainers
  - Klaus Ma (@k82cn, Huawei)
  - Kevin Wang (@kevin-wangzefeng, Huawei)
  - Zhonghu Xu (@hzxuzhonghu, Huawei)
  - Quinton Hoole (@quinton-hoole, Huawei)
  - Animesh Singh (@animeshsingh, IBM)
  - Jun Gong (@hex108, Tencent)

  New in the Year

  - Thor Wu (@Thor-wl, Huawei)
  - William Wang (@william-wang, Huawei)
  - Peng Wang (@wpeng102, Huawei)
  - Liang Tang (@shinytang6, Baidu)
  - Zhengyu Xu (@zen-xu, Ruitian Investment)
  - Wentian Jiang (@alcorj-mizar, CCB)
  - Chi Zhang (@hudson741, BOSSzhipin)
  - Yan Zhou (@merryzhou, Cambricon)


* What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

  Number of public adopters grew from 13 to 20, with **production grade adopters grew from 2 to 10**.

  Part of adopters are:
  - HUAWEI CLOUD
  - Tencent
  - iQIYI
  - VIPS (VIP.com)
  - Ruitian Investment
  - Xiaohongshu
  - DiDi

  Refer to: [Volcano Community Adopters](https://github.com/volcano-sh/volcano/blob/master/docs/community/adopters.md) for details


* How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

  - Development
    - Release Cadence: 1 minor release every three month
    - Moved to 1.3.0
    - Versioning Scheme: following [Semantic Versioning](https://semver.org/)
    - Key features added to the project:
      - v1.0:
        - GPU Sharing
        - Preempt and reclaim enhancement
        - Dynamic scale up and down
        - Support integrate with flink operator
        - Support DAG job with Argo

      - v1.1:
        - Add monitor component
        - Support resource reservation for big job automatically
        - Support HDRF

      - v1.2:
        - Add TDM plugin
        - Add SLA plugin

      - v1.3:
        - Support minAvailable at task level
        - Support minSuccess for job
        - Support task-topology


* What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

  - Technical
    - More integration with mainstream computing architectures
    - Richer advanced scheduling algorithms
    - [Project Development Roadmap](https://github.com/volcano-sh/volcano/blob/master/docs/community/roadmap.md)

  - Community
    - Wider user adoption
    - Wider developer participation
    - More evangelism


* How can the CNCF help you achieve your upcoming goals?

  - We need more speaking and marketing opportunities to help attract more contributors and user adoptions.
  - We also need some technical writers to help improve documentation and website content.

* Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?
  - Yes, incubation proposal is under preparation.

## Project Links
 - [Website](https://volcano.sh/en/)
 - [Github](https://github.com/volcano-sh/volcano)
 - [Slack](https://volcano-sh.slack.com)
 - [Mailing List](https://groups.google.com/forum/#!forum/volcano-sh)
 - [Twitter](https://twitter.com/volcano_sh)
 - [Issue Tracker](https://github.com/volcano-sh/volcano/issues)

