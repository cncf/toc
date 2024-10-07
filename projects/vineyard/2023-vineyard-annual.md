# Vineyard 2023 Annual Review

## Background

[Vineyard](https://v6d.io) is an in-memory immutable data manager that provides
**out-of-the-box high-level** abstraction and **zero-copy in-memory** sharing for
distributed data in big data tasks, such as graph analytics, numerical computing, and
machine learning. Vineyard is design to address the inefficiency of data sharing in
big data analytical workflows on Kubernetes.

**Vineyard provides**:

1. Efficient in-memory data management and zero-copy sharing across different systems.
2. Out-of-the-box high-level data abstraction for distributed objects (e.g., tensors, tables,
   graphs, and distributed datasets) and efficient polyglot support (currently including C++,
   Python, Go, Rust and Java).
3. Seamless integration with Kubernetes for cluster deployment and management, workloads
   orchestration, and observability.
4. Out-of-the-box integration with workflow orchestration engines (including
   [Apache Airflow](https://github.com/v6d-io/v6d/tree/main/python/vineyard/contrib/airflow),
   [Flyte](https://github.com/v6d-io/v6d/tree/main/python/vineyard/contrib/kedro)
   and [Kubeflow Pipelines](https://github.com/v6d-io/v6d/tree/main/k8s/examples)),
   providing end-users with a unified and intrusive experience to leverage Vineyard in
   their data-intensive workflows to improving performance.

**Alignment with CNCF**:

1. Vineyard builds on Kubernetes for deploying and scaling, and the objects are observable
   in Kubernetes as CRDs.
2. Vineyard makes efficient zero-copy sharing possible for data-intensive workflows on
   cloud-native infrastructure by a data-aware Kubernetes scheduler plugin.
3. Vineyard adopts an immutable object design, which aligns with the immutable infrastructure
   of the cloud-native environment.
4. Vineyard aligns with the CNCF effort on helping migrate batching system workflows to cloud
   native environments.

## Development

### DevStats

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

- Number of contributors: 26 -> 40
- Github stars: 600+ -> ~750
- Github forks: 80+ -> 110+

### Highlights of new features

Vineyard published 8 release (one release about per 1.5 month) since the last annual review.
The major new features and improvements include:

- Language SDKs in Rust and Go, where the Rust SDK was a collaboration with our external
  end-user, and enabled users seamlessly and efficiently interoperating their data between
  Python and Rust.
- Integration with the workflow engine Kedro, and gained attention in the Kedro community.
- Vineyard supports the [Apache Hive](https://github.com/apache/hive) data processing engine,
  letting users can easily connect traditional data processing pipelines built with the
  Hadoop ecosystem with emerging big-data and AI applications (e.g., applications in the
  [PyData](https://pandas.pydata.org/community/ecosystem.html) community).
- A initial version of CSI driver, which helped Vineyard aligned with the Kubernetes platform
  and enables users to leverage Vineyard in their Kubeflow pipelines to optimize the data
  sharing between steps with only minor changes to their existing source code.

### Academic Research

We have conducted a series of research work around Vineyard and published the paper
_Vineyard: Optimizing Data Sharing in Data-Intensive Analytics_ in SIGMOD 2023, a top-tier
conference in the data management community.

- Wenyuan Yu, Tao He, Lei Wang, Ke Meng, Ye Cao, Diwen Zhu, Sanhong Li, Jingren Zhou.
  Vineyard: Optimizing Data Sharing in Data-Intensive Analytics. ACM SIG Conference on
  Management of Data (SIGMOD), industry, 2023. [https://dl.acm.org/doi/10.1145/3589780](https://dl.acm.org/doi/10.1145/3589780).

## Maintainers

> How many maintainers do you have, and which organisations are they from? (Feel free
> to link to an existing MAINTAINERS file if appropriate.)

We currently have 10 maintainers and 2 committers and have [a maintainer list on Github](https://github.com/v6d-io/v6d/blob/main/MAINTAINERS.md).

- **Initial maintainers**

  | Name | GitHub ID | Affiliation | Email |
  | --- | --- | --- | --- |
  | Tao He | [sighingnow](https://github.com/sighingnow) | Alibaba | [linzhu.ht@alibaba-inc.com](mailto:linzhu.ht@alibaba-inc.com) |
  | Xiaojian Luo | [luoxiaojian](https://github.com/luoxiaojian) | Alibaba | [lxj193371@alibaba-inc.com](mailto:lxj193371@alibaba-inc.com) |
  | Ke Meng | [septicmk](https://github.com/septicmk) | Alibaba | [mengke.mk@alibaba-inc.com](mailto:mengke.mk@alibaba-inc.com) |
  | Wenyuan Yu | [wenyuanyu](https://github.com/wenyuanyu) | Alibaba | [wenyuan.ywy@alibaba-inc.com](mailto:wenyuan.ywy@alibaba-inc.com) |
  | Weibin Zeng | [acezen](https://github.com/acezen) | Alibaba | [qiaozi.zwb@alibaba-inc.com](mailto:qiaozi.zwb@alibaba-inc.com) |
  | Siyuan Zhang | [siyuan0322](https://github.com/siyuan0322) | Alibaba | [siyuanzhang.zsy@alibaba-inc.com](mailto:siyuanzhang.zsy@alibaba-inc.com) |
  | Diwen Zhu | [andydiwenzhu](https://github.com/andydiwenzhu) | Alibaba | [diwen.zdw@alibaba-inc.com](mailto:diwen.zdw@alibaba-inc.com) |

- **New maintainers in this year**

  | Name | GitHub ID | Affiliation | Email |
  | --- | --- | --- | --- |
  | Ye Cao | [dashanji](https://github.com/dashanji) | Alibaba | [caoye.cao@alibaba-inc.com](mailto:caoye.cao@alibaba-inc.com) |
  | Shumin Yuan | [vegetableysm](https://github.com/vegetableysm) | Alibaba | [yuanshumin.ysm@alibaba-inc.com](mailto:yuanshumin.ysm@alibaba-inc.com) |
  | Denghao Li | [lidh15](https://github.com/lidh15) | PingAn Tech | [lidhrandom@gmail.com](mailto:lidhrandom@gmail.com) |

## Adoption

> What do you know about adoption, and how has this changed since your last review / since
> you joined Sandbox? If you can list companies that are end-users of your project, please
> do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

We have tracked the following two major adoption since [our last annual review]()

- _StartDT (Qidianyun)_: transiting towards production stage
  - _StartDT_ is a startup company in China, providing a cloud-native data platform for
    big-data analytics and machine learning applications.
  - Vineyard is currently used in their Python-centric data processing pipelines to
    share distributed dataframe artifacts between steps, and help build a composable
    and efficient data processing platform to end-users.

    Vineyard has passed their eager-evaluation, and they are working on building their
    distributed data processing platform on top of Vineyard.

- _PingAn Tech_: production stage
  - _PingAn_ is a large-scale fin-tech company in China.
  - Vineyard is used in their data science platform to support efficient dataset sharing and
    management among data science researchers.

    The status of Vineyard in their platform has been transited from testing to production
    stage, and one of their engineers has become a maintainer of the Vineyard project.

Besides these two major companies, since our last annual review, we have also noticed some other
questions about using Vineyard in machine learning inference scenarios, but we haven't tracked
the actual adoption yet.

## Project Goals

> How has the project performed against its goals since the last review? (We won't penalize you
> if your goals changed for good reasons.)

Vineyard has successfully archived the goals about easing the getting started process for
end-users from three aspects:

1. Out-of-the-box integration with data processing systems, especially Spark and Hive,
   the most popular data processing engines in the big data community;
2. Data processing pipeline orchestration: providing non-intrusive interfaces to help users
   migrate their existing data processing pipelines to Vineyard on Kubernetes and finally
   benefit from the efficient data sharing;
3. Seamless inter-operability with other systems in the cloud-native environments: we invest
   a lot of effort in the Vineyard operator to help use deploy vineyard along with their
   workloads in a non-intrusive, declarative way and has tested the functionality with
   _GraphScope_ in end-users production environments.

Besides, Vineyard has successfully attracted new end-users from the big data community to
adopt Vineyard in their own data processing platform, and the feedback from the Kedro
community is also positive.

> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

Our current goals are mainly focused on the attracting more end-user to adopt Vineyard in
their scenarios from different domains. Specifically, we are keeping moving towards the
following goals in the next year:

1. Optimizing our current Kubeflow integration and find more opportunities to evaluate
   and deploy Vineyard in production machine learning applications;
2. Publish our integration with the big data processing systems to their end-user
   community and gather feedback for further improvements;
3. Seeking more opportunities to evaluate Vineyard in the emerging LLM applications,
   for both data preprocessing, training, and inference serving to see if Vineyard
   can bring added value to these applications as where the data cost is usually high;
4. Getting engaged with the Batch System WG in CNCF to seek opportunities about
   further collaboration with other projects in CNCF.

## CNCF membership

> How can the CNCF help you achieve your upcoming goals?

Vineyard has incredibly benefited from CNCF since accepted as a sandbox project. We believe
the end-users in the CNCF community are critical for Vineyard to become successful. With
the help of CNCF service desk, we have successfully built a new website for Vineyard, which
is more friendly to end-users. We are also working on components like CSI driver and hope
that could make the inter-operation with other projects in the CNCF community easier.

We will host a _Project Kiosk_ in this KubeCon China and hope to get more feedback from the
community, and hope to get more feedback from the community. Furthermore, we hope we could have
more opportunities to introduce our project to border end-users in the CNCF community to
increase adoption.

## Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)?

We think our project vineyard still needs further exploration to get border adoption in the
end-user's production deployment and gather more feedback, and we are looking forward to
meeting the incubation criteria in the near future.
