# Fluid 2022 Annual Review

This is the annual review for the [Fluid](https://github.com/fluid-cloudnative/fluid) project in the past year.

# Content Outline

- [Background](#Background)
- [Development metrics](#Development_metrics)
- [Maintainer](#Maintainer)
- [Project adoption](#Project_adoption)
- [Project goals](#Project_goals)
- [CNCF Membership](#CNCF_Membership)
- [Incubation readiness](#Incubation_readiness)

# Background

Fluid is an open source Kubernetes-native Distributed Dataset Orchestrator and Accelerator for data-intensive applications, such as big data and AI applications. It primarily provides the following features:

* **Native Support for DataSet Abstraction**: Fluid provides data warm-up and acceleration for cloud applications by using a distributed cache engine (Alluxio) in Kubernetes with Observability, Portability and Horizontal Scalability.

* **Cloud Data Warming up and Accessing Acceleration**: Fluid provides data warm-up and acceleration for cloud applications by using a distributed cache engine (Alluxio) in Kubernetes with Observability, Portability and Horizontal Scalability.

* **Co-Orchestration for Data and Application**: During application scheduling and data placement on the cloud, taking both the app's characteristics and data location into consideration, to improve the performance.

* **Support Multiple Namespaces Management**: User can create and manage datasets in multiple namespaces.

* **Support Heterogeneous Data Source Management**: Unify the Data access for OSS, HDFS, CEPH and Other underlayer storages.

Here are several key concepts:

* **Dataset**: A DataSet is a set of data logically related that can be used by computing engines, such
  as Spark for big data analytics and TensorFlow for AI applications. Intelligently leveraging data
  often creates core industry values. Managing DataSets may require features in different
  dimensions, such as security, version management and data acceleration. We hope to start with data
  acceleration to support the management of datasets.
* **Runtime**: The execution engine that enforces dataset security, provides version management and data
  acceleration capabilities. The Runtime defines a set of interfaces to manage DataSets in their
  life cycle, so the management and acceleration of datasets can be implemented behind these
  interfaces.


### Alignment with CNCF

- Fluid was accepted as a CNCF Sandbox project on April 28th, 2021.


# Development metrics

The Fluid devstats page and dashboards can be found [here](https://fluid.devstats.cncf.io/).

Here are some links worth mentioning:

* [Stars and forks](https://fluid.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now)
* [Issues](https://fluid.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1)
* [PRs](https://fluid.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&var-period=d&var-repogroup_name=All)
* [Commits](https://fluid.devstats.cncf.io/d/2/commits-repository-groups?orgId=1)

According to devstats, Fluid currently has [281](https://fluid.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [22](https://fluid.devstats.cncf.io/d/5/companies-table?orgId=1) companies. On average, there are [~300 contributions per month](https://fluid.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) contained within [50 merged PRs per month](https://fluid.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=prs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).

The community has grown since the project entered the CNCF sandbox.

* We held bi-weekly community meetings constantly (total 28 times as of June 2021). The meeting records can be found in [here](https://github.com/fluid-cloudnative/community/wiki/Meeting-Agenda-and-Notes). The average number of meeting attendees is around 20.
* Github stars increased from **470+ to 1080+**
* Github forks increased from **130+ to 390+**
* Github commits increased from **670+ to 1350+**
* Contributing organizations including:
  - Alibaba, Unisound AI, China Telecom, CCB Fintech Company, SAP, IBM, Baidu, Bytedance, Tencent, PureSoftware, DMetaSoul, VIPKid, Weibo, Qihoo 360.

Starting from Fluid joining CNCF Sandbox project, the contributors opened 413 new issues and 330 of the issues(containing the previous issues) were closed. Meanwhile, they opened 794 new PRs and 806 of the total PRs(containing the previous PRs) are merged into the master branch. Also, we have held bi-weekly community meetings and the meeting notes and agenda can be found [here](https://github.com/fluid-cloudnative/community/wiki).

# Maintainer

Currently, the project has 5 maintainers. The following lists the detailed information of the maintainers.

| Name | GitHub ID | Affiliation |
| ---- | --------- | ----------- |
| [Rong Gu](mailto:gurong@nju.edu.cn) | [RongGu](https://github.com/RongGu) | Nanjing University |
| [Yang Che](mailto:cheyagn52@gmail.com) | [cheyang](https://github.com/cheyang) | Alibaba |
| [Bin Fan](mailto:binfan@alluxio.com) | [apc999](https://github.com/apc999) | Alluxio |
| [Zhihao Xu](mailto:trafalgarz@outlook.com) | [TrafalgarZZZ](https://github.com/TrafalgarZZZ) | Nanjing University |
| [Kai Zhang](mailto:wsxiaozhang@gmail.com) | [wsxiaozhang](https://github.com/wsxiaozhang) | Alibaba |

Also in this year, we have added 4 project committers due to their outstanding contribution.
The following lists the detailed information of the project committers.

| Name | GitHub ID | Affiliation |
| ---- | --------- | ----------- |
| [Yuandong Xie](mailto:xieydd@gmail.com) | [xieydd](https://github.com/xieydd) | Tencent Cloud |
| [Lingwei Qiu](mailto:qlw705706@gmail.com) | [yangyuliufeng](https://github.com/yangyuliufeng) | China Telecom |
| [Tao Wang](mailto:wangtaod13@gmail.com) | [frankleaf](https://github.com/frankleaf) | Alibaba |
| [Weiwei Zhu](mailto:zww@hdls.me) | [zwwhdls](https://github.com/zwwhdls)|JuiceData|

More descriptions of maintainers and committers are listed in
the [MAINTAINERS_COMMITTERS.md](https://github.com/fluid-cloudnative/fluid/blob/master/MAINTAINERS_COMMITTERS.md) file.

# Project adoption

Fluid has been adopted as the foundation of public cloud Kubernetes AI/Big Data solutions. Many companies have fully leveraged Fluid build their cloud native AI/Big Data platform across hybrid, multi-cloud, serverless environments.

Beyond this, Fluid have **15 adopters** totally, and the whole adopters list can be found in the existing [ADOPTERS file](https://github.com/fluid-cloudnative/fluid/blob/master/ADOPTERS.md).

Since Fluid joining CNCF Sandbox project, there are **9 new adopters**(containing phase change) which are listed as follows.

| Organization | Description of Use |
  | ------------ | ------------------ |
| [AnalyticDB On Alibaba Cloud](https://www.aliyun.com/product/ApsaraDB/ads)  | LakeCache Containers |
| [Bilibili](http://www.bilibili.com/)  | Bilibili AI platform |
| [Gaussian Robotics](https://www.gaussianrobotics.com/)  | AI platform on k8s |
| [Metaapp](https://www.metaapp.cn/)  | Recommendation |
| [Platform of Artificial Intelligence On Alibaba Cloud](https://www.aliyun.com/product/bigdata/product/learn)   | PAI Deep Learning Containers |
| [Qihoo 360](http://www.360.cn/)  |  BigData platform |
| [Weibo](http://www.weibo.com/)  | Weibo deep learning platform |
| [鹏行智能](https://www.pxing.com/)  | Pxing AI training platform |
| [数元灵](https://www.dmetasoul.com/)  | MetaSpore Distributed Multimodal AI platform |


# Project goals

Fluid targets on the Kubernetes-native Distributed Dataset Orchestrator and Accelerator for data-intensive applications. Following the above goals, in the last year we released a total of 2 versions and each version added the following features:

- v0.6.0
	- New features
    	- Support dataset cache autoscaling and cronscaling
    	- Add dataset mounting point dynamically update feature
    	- Enhance dataset cache aware Pod scheduling
    	- Enhance HA support for cache Runtime
    	- Support new cache Runtime：GooseFS
	- Bugs fixed
    	- Fix [if alluxioruntime is nonroot, databackup will fail](https://github.com/fluid-cloudnative/fluid/issues/745)
    	- Fix [Node labels exceeds maximum length limit for long namespace and name](https://github.com/fluid-cloudnative/fluid/issues/704)
- v0.7.0
	- New features
    	- Support fuse sidecar auto injection for all the runtimes，it’s helpful for no CSI
      	environment
    	- Support fuse auto recovery and upgrade
    	- Support lazy fuse mount mode
    	- Support New cache runtime: JuiceFS
	- Bugs fixed
    	- Fixes [Failed to update status of dataload](https://github.com/fluid-cloudnative/fluid/issues/1436)
    	- Fixes [Failed to delete dataload when target dataset is removed](https://github.com/fluid-cloudnative/fluid/issues/1419)
    	- Fixes [node-driver-registrar will not receive any volume umount in subdirectories of kubelet-dir](https://github.com/fluid-cloudnative/fluid/issues/1048)

In the future, we plan to add the following new features in Fluid:

- Integration with BigData Analysis Applications
- More native support for AI applications running on Serverless
- Cross-namespace data accessing API design
- Enhance cache warmup feature
- Disk resource management for better cache scheduling

At the same time, the community has been committed to the maintenance of the documentation such as adding more examples or video for new features.
Besides, the [Fluid website](https://fluid-cloudnative.github.io/) is under maintenance all the time.

# CNCF Membership

It can be seen from the performance of the last year that Fluid has received a lot of benefits since joining the CNCF Sandbox project.
We hope that CNCF can increase the promotion of Fluid so that Fluid will have more opportunities to show its work and gain more contributors and users.

## How the CNCF can help to achieve the upcoming goals
Similar to many other Sandbox projects, we wish to get the following help from CNCF:
- More channels to advocate the project.
- More chances to collaborate with other projects in CNCF or even out of CNCF.
- Technical writing support for project documents.

# Incubation readiness

Fluid intends to remain a CNCF Sandbox project currently. We are not ready to apply for the CNCF Incubation project yet since we have joined the CNCF Sandbox project just for a relatively short time.
In the future, we are intended to develop more committers, contributors and users, and then we will consider applying for the CNCF Incubation project.