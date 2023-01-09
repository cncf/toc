# Vineyard 2022 Annual Review

## Background

[Vineyard](https://v6d.io) is an in-memory immutable data manager that provides
**out-of-the-box high-level** abstraction and **zero-copy in-memory** sharing for
distributed data in big data tasks, such as graph analytics, numerical computing, and
machine learning.

**Vineyard provides**:

1. Efficient in-memory data management and zero-copy sharing across different systems.
2. Out-of-the-box high-level data abstraction for distributed objects (e.g., tensors, tables,
   graphs) and efficient polyglot support (currently including C++, Python, and Java).
3. Built-in streaming support for data accessing and across system pipelining.
4. An extensible driver framework and a set of efficient built-in drivers for eliminating
   the boilerplate part in computation engines, e.g., I/O, serialization, and checkpointing.

**Alignment with CNCF**:

1. Vineyard builds on Kubernetes for deploying and scaling and the objects are observable
   in Kubernetes as CRDs.
2. Vineyard makes efficient zero-copy sharing possible for data-intensive workflows on
   cloud-native infrastructure by a data-aware Kubernetes scheduler plugin.
3. Vineyard adopts a immutable object design, which aligns with the immutable infrastructure
   of the cloud-native environment.

**Vineyard was accepted as a CNCF sandbox project on Apr 28th, 2021.**

## DevStats

> Include a link to your project’s devstats page. We will be looking for signs of consistent
> or increasing contribution activity. Please feel free to add commentary to add colour to
> the numbers and graphs we will see on devstats.

- Stargazers and Forks
  - [https://vineyard.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-2y&to=now](https://vineyard.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-2y&to=now)
- Commits per week
  - [https://vineyard.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=w&var-repogroups=All&from=now-6M&to=now](https://vineyard.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=w&var-repogroups=All&from=now-6M&to=now)
- Contributors and Companies
  - [https://vineyard.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&var-period=d7&var-repogroup_name=All&from=now-1y&to=now](https://vineyard.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&var-period=d7&var-repogroup_name=All&from=now-1y&to=now)

The vineyard community has grown since the project entered the CNCF sandbox.

- Number of contributors: 11 -> 26
- Github stars: 300+ -> 600+
- Github forks: 20+ -> 80+
- Contributing organizations: 1 -> 12

## Maintainers

> How many maintainers do you have, and which organisations are they from? (Feel free
> to link to an existing MAINTAINERS file if appropriate.)

We currently have 7 maintainers and 2 committers and have [a maintainer list on Github](https://github.com/v6d-io/v6d/blob/main/MAINTAINERS.md).

- **Initial maintainers**

  | Name | GitHub ID | Affiliation | Email |
  | --- | --- | --- | --- |
  | Tao He | [sighingnow](https://github.com/sighingnow) | Alibaba | [linzhu.ht@alibaba-inc.com](mailto:linzhu.ht@alibaba-inc.com) |
  | Xiaojian Luo | [luoxiaojian](https://github.com/luoxiaojian) | Alibaba | [lxj193371@alibaba-inc.com](mailto:lxj193371@alibaba-inc.com) |
  | Wenyuan Yu | [wenyuanyu](https://github.com/wenyuanyu) | Alibaba | [wenyuan.ywy@alibaba-inc.com](mailto:wenyuan.ywy@alibaba-inc.com) |
  | Weibin Zeng | [acezen](https://github.com/acezen) | Alibaba | [qiaozi.zwb@alibaba-inc.com](mailto:qiaozi.zwb@alibaba-inc.com) |
  | Siyuan Zhang | [siyuan0322](https://github.com/siyuan0322) | Alibaba | [siyuanzhang.zsy@alibaba-inc.com](mailto:siyuanzhang.zsy@alibaba-inc.com) |
  | Diwen Zhu | [andydiwenzhu](https://github.com/andydiwenzhu) | Alibaba | [diwen.zdw@alibaba-inc.com](mailto:diwen.zdw@alibaba-inc.com) |

- **New maintainers in this year**

  | Name | GitHub ID | Affiliation | Email |
  | --- | --- | --- | --- |
  | Ke Meng | [septicmk](https://github.com/septicmk) | Alibaba | [mengke.mk@alibaba-inc.com](mailto:mengke.mk@alibaba-inc.com) |

- **New Committers in this year**

  | Name | GitHub ID | Affiliation | Email |
  | --- | --- | --- | --- |
  | Lihong Lin | [linlih](https://github.com/linlih) | PKU | [linlh@stu.pku.edu.cn](mailto:linlh@stu.pku.edu.cn) |
  | Pei Li | [peilii](https://github.com/peilii) | CMU | [peili.dev@gmail.com](mailto:peili.dev@gmail.com) |

## Adoption

> What do you know about adoption, and how has this changed since your last review / since
> you joined Sandbox? If you can list companies that are end-users of your project, please
> do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

We know several cases where vineyard has been adopted in both testing and production environments.

- _GraphScope_: production stage
  - _GraphScope_ is an open-source graph processing platform.
  - Vineyard is used in graphscope to provide distributed shared in-memory storage for different
    graph processing engines.
- _weilaisudu_: transiting from testing to production stage
  - _weilaisudu_ is the company behind the project [Mars](https://github.com/mars-project/mars),
    a distributed scientific computing engine that provides numpy and pandas like API.
  - Vineyard is used as the shared-memory storage for actors that do computation on chunks.
- _ESRF_: testing stage
  - _ESRF_ is a joint research facility situated in France, one of the biggest x-ray science
    facilities in Europe.
  - VIneyard is used in the BLISS software to serve as the shared storage between sensors and
    data processing jobs.
- _PingAn_: testing stage
  - _PingAn_ is a large-scale fin-tech company in China.
  - Vineyard is used in a research platform to support efficient dataset sharing and management
    among data science researchers.

We have also integrated with the apache-airflow project, which is a workflow orchestration engine
and has been widely adopted. We have published airflow-vineyard-provider on Astronomer Registry
and received much feedback from end-users, but we haven't tracked the actual adoption yet.

## Project Goals

> How has the project performed against its goals since the last review? (We won't penalize you
> if your goals changed for good reasons.)

Vineyard has successfully archived the goal of bringing value to big data analytical workflows
on Kubernetes. We have shown the gain in an internal project which involves both ETL, graph
computation, and machine learning jobs.

Our goal hasn't changed since becoming CNCF sandbox project and we are still aiming at supporting
a more efficient big data analytical workflow on the cloud-native infrastructure. Specifically,
we'll keep moving towards  following goals in the next year:

1. Providing efficient cross-engine data sharing for data-intensive workflows in Kubernetes
2. Integrating with projects in the cloud-native community for orchestration and scheduling
   and integrating with more big data computing engines to improve the end-to-end efficiency.
3. Building a new cloud-native paradigm for big data applications working  together. By
   integrating Vineyard, Kubernetes can help orchestrating data and workloads together for
   better alignment and efficiency.

We still need to do more to engage end-users to show the value-added of the vineyard project.

## CNCF membership

> How can the CNCF help you achieve your upcoming goals?

Vineyard has incredibly benefited from CNCF since accepted as a sandbox project. We believe
the end-users in the CNCF community are critical for Vineyard to become successful. We have
submitted proposals for the KubeCon and CNCF Conferences in the past year but got rejected.
We hope we could have more opportunities to introduce our project to border end-users in the
CNCF community to increase adoption.

## Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)?

We think our project vineyard still needs further exploration to get border adoption in the
production environment and we are looking forward to meeting the incubation criteria in near
future.
