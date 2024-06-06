# Pravega.io Sandbox 2022 Annual Review

## Table of Contents

  - [Project Devstats](#project-devstats)
  - [Maintainers](#maintainers)
  - [Adoption](#adoption)
  - [Project Performance](#project-performance)
    - [Open Governance](#open-governance)
    - [Releases & Features](#releases--features)
    - [Open Source](#open-source)
    - [Community Outreach](#community-outreach)
    - [Blog posts](#blog-posts)
    - [Conference presentations](#conference-presentations)
  - [Current Goals](#current-goals)
  - [Help with Goals](#help-with-goals)
  - [Incubation](#incubation)

## Project Devstats
> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. 
> Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

The Pravega GitHub organization was active across 30 public repositories in 2022, including:
- **Core Pravega**: The code base for the streaming storage engine, including the Java client, a CLI and long-term storage adapters;
- **Connectors**: Apache Flink, Apache Spark, Presto, Trino, GStreamer, Debeizum;
- **Helm charts and operators**: For Pravega, ZooKeeper, BookKeeper;
- **Additional client bindings**: For C#, NodeJS, and Go;
- **Schema Registry**: To enable schema management;
- **Samples**: Code samples for various components in the GitHub organization.

In 2022, the Pravega organization saw around 483 new PRs, among which about 246 are in Pravega alone.

For Pravega repo alone, GitHub Forks and Stars have increased by 10.3% and 8.2%, respectively.
Across all the repositories, the activity trend is shown in the table below:

| Metric (all repos) | Jan. 1st, 2022 | Dec. 31th, 2022 | YoY change |
|--------------------|----------------|----------------|------------|
| Stargazers         | 2138           | 2360           |10.4% increase|
| Forks              | 716            | 771            |7.7% increase|
| Open Issues & PRs  | 474            | 438            |7.6% decrease|

This year saw 25 new contributors across all repositories.

Source: https://pravega.devstats.cncf.io

## Maintainers
> How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing 
> MAINTAINERS file if appropriate.)

We continue the organization of contributors with write access into two levels. 
The Steering Committee comprises the contributors that manage the organization 
and map to the CNCF maintainers. Committers have access to groups of repositories that 
we organized into: Core, Ecosystem, Operators, Integration and Tools.

Maintainers file: https://github.com/pravega/.github/blob/main/MAINTAINERS

Number of maintainers (Steering Committee): 7 (+2 compared to 2021)

Number of organizations for maintainers: 3 (Dell, Wheels Up, Amazon)

Number of committers (across all groups): 32 (+4 compared to 2021)

## Adoption
> What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? 
> If you can list companies that are end users of your project, please do so. (Feel free to link to an existing 
> ADOPTERS file if appropriate.)

Pravega is the core storage engine of Dell’s [Streaming Data Platform](https://www.delltechnologies.com/en-us/storage/streaming-data-platform.htm) 
product. Dell is focusing on exploiting Pravega and Dell's Streaming Data Platform in the areas of storage products, 
edge deployments, video analytics, and data management. What is more,
[Atos](https://atos.net/en/2022/press-release_2022_05_04/atos-launches-innovative-edge-to-cloud-5g-and-ai-enabled-solution) 
is one of the main partners working with Dell to exploit Dell's Streaming Data Platform even further.

Pravega has been adopted as core architectural part of two new European Research projects: 
[NEARDATA (101092644)](https://neardata.eu/) and [CLOUDSKIN (101092646)](https://cloudskin.eu/).
Pravega will be used to satisfy several use-cases, such as computer-assisted surgery, near real-time
storage of genomic data, and smart agriculture, among others. We expect academic and industrial
project partners to know and adopt Pravega even beyond the scope of the projects.

## Project Performance
> How has the project performed against its goals since the last review? (We won’t penalize you if your 
> goals changed for good reasons.)

**2022 Milestones:**
- Open governance is well-established and disseminated across the community
- Frequent major releases containing new features and new integrations
- External open source contributions and increasing industrial exploitation of Pravega
- Community outreach activities executed (conferences, blogs, etc.)

### Open Governance

Committers are organized into 4 teams: Core, Operators, Ecosystem, and Integration & Tools. 
Pravega committer teams manage a total of 43 repositories. Open governance has been instituted following a model 
inspired by Apache projects: https://github.com/pravega/.github/blob/main/governance.md

### Releases & Features

Pravega has made steady progress closely following its published roadmaps through regular major and minor 
releases: 0.10.2, 0.11.0, 0.12.0.

Features enabled by these releases include test stabilization and increased coverage, improvements
in hash-based KV Tables, new Admin CLI repair commands, AWS S3 long-term storage binding,
Prometheus support for metrics, new metrics, Large Events support, and security fixes, among others.

### Open Source

This year, the Pravega project includes a more mature RUST client and three new client bindings (C#, NodeJS, and Go)
as well as the Debezium sink:

**New Pravega client binding for NodeJS**<br>
https://github.com/pravega/pravega-client-rust/tree/master/nodejs

**New Pravega client binding for Go**<br>
https://github.com/pravega/pravega-client-rust/tree/master/golang

**New Pravega client binding for C# (in collaboration with Washington State University)**<br>
https://github.com/pravega/pravega-client-csharp

**New Debezium sink for Pravega**<br>
https://debezium.io/documentation/reference/operations/debezium-server.html#_pravega

**Added Pravega as a use case for Bookkeeper**<br>
PR: https://github.com/openmessaging/benchmark/pull/199 <br> README: https://bookkeeper.apache.org/docs/overview/

### Community Outreach

Pravega Community meetings were founded under Cloud Native Community Groups:<br>
https://community.cncf.io/pravega-community/

Community meeting in which Olivier Nouguier presents an introduction to ZIO alongside his ZIO Pravega connector:<br>
https://www.youtube.com/watch?v=uNOmAw5nL7o

We have organized a summer internship with Washington State University to work on the C# Pravega binding:
https://github.com/WSUCptSCapstone-Fall2022Spring2023/dell-pravegaapi

Organization of a course for Pravega at Duke Kunshan University (China):
https://mp.weixin.qq.com/s/ULZZOfhPmts9xqdn1Ds-kA

Joint project with MatrixOne and Pravega for university students last summer (China): 
https://summer-ospp.ac.cn/2022/#/org/prodetail/22fa00455

The Zookeeper Operator repository is increasing its activity and concentrating external contributions:
https://github.com/pravega/zookeeper-operator

### Blog posts

In 2022, the following blog posts were published:

**Change Data Capture with Pravega + Debezium**<br>
https://cncf.pravega.io/blog/2022/08/16/change-data-capture-with-pravega-debezium/

### Conference presentations

Presentation of Pravega & ZIO at Scala.io (Oct. 2022):<br>
https://www.youtube.com/watch?v=lXdkDg0RvNI

Presentation of Pravega at SDC India 2022 (Aug. 2022):<br>
https://www.youtube.com/watch?v=1YrhO8yWmus

Presentation at WOSCx (Jun. 2022):<br>
https://www.youtube.com/watch?v=0x_4rLiKbMo

Pravega Office Hour at KubeconEU (Apr. 2022):<br>
https://www.youtube.com/watch?v=huDZ0boA-Do

Flink Forward Asia (Jan. 2022):<br>
https://developer.aliyun.com/special/ffa2021/live

Flink Forward Asia (Jan. 2022):<br>
https://www.bilibili.com/video/BV1RY41187fJ/?p=2


## Current Goals
> What are the current goals of the project? For example, are you working on major new features? 
> Or are you concentrating on adoption or documentation?

Working on new features and tools in the 0.13 release and beyond: 
https://github.com/pravega/pravega/blob/master/documentation/src/docs/roadmap.md

Focusing on popularizing Pravega and increasing its adoption. 
Generating mixed-media content at all skill levels, especially beginner and advanced material. 
Disseminating content more broadly across social media by leveraging individual networks.

Working on adapting Pravega to Edge environments to align it with the current main
exploitation field.

## Help with Goals
> How can the CNCF help you achieve your upcoming goals?

There are a couple of options:
- **Documentation**: Support for content generation of technical documentation and blog posts. 
- **Test infrastructure**: We need always-on Kubernetes clusters to run our system tests, including 
longevity tests and Litmus Chaos tests in the open.

## Incubation
> Do you think that your project meets the criteria for incubation?

Not yet, only one qualifying end-user presently.
