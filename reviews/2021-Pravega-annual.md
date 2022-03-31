# Pravega.io Sandbox 2021 Annual Review

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
    - [Hackathons](#hackathons)
    - [Conference presentations](#conference-presentations)
  - [Current Goals](#current-goals)
  - [Help with Goals](#help-with-goals)
  - [Incubation](#incubation)

## Project Devstats
> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

The Pravega GitHub organization was active across 30 repositories in 2021, including:
- **Core Pravega**: The code base for the streaming storage engine, including the Java client, a CLI and long-term storage adapters;
- **Connectors**: Apache Flink, Apache Spark, Presto, Trino, GStreamer;
- **Helm charts and operators**: For Pravega, ZooKeeper, BookKeeper;
- **Additional client bindings**: For Rust and Python;
- **Schema Registry**: To enable schema management;
- **Samples**: Code samples for various components in the GitHub organization.

In 2021, the Pravega organization saw around 1,000 PRs, among which about 500 are in Pravega alone.

GitHub Forks and Stars have increased by 22% for Pravega repo alone. Stars have increased by 25% across all repositories.

| Metric (all repos) | Jan. 3rd, 2021 | Jan. 9th, 2022 | YoY change |
|--------------------|----------------|----------------|------------|
| Stargazers         | 1356           | 1693           |25% increase|
| Forks              | 280            | 350            |25% increase|
| Open Issues & PRs  | 424            | 231            |46% decrease|

This year saw 34 new contributors across all repositories.

https://pravega.devstats.cncf.io/

## Maintainers
> How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

We have organized contributors with write access into two levels. The Steering Committee comprises the contributors that manage the organization and map to the CNCF maintainers. Committers have access to groups of repositories that we organized into: Core, Ecosystem, Operators, Integration and Tools.

Maintainers file: https://github.com/pravega/.github/blob/main/MAINTAINERS

Number of maintainers (Steering Committee): 5

Number of organizations for maintainers: 3 (Dell, Wheels Up, Amazon)

Number of committers: 28

## Adoption
> What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

One end-user company has joined the project and is increasing their contributions, having added support for Prometheus metrics as a recent major contribution.

This year has seen significant adoption of and commitment to Pravega internal to Dell Technologies, especially in the realms of storage products, edge sensor collection, video analytics and AI.

Pravega remains vital to Dell’s [Streaming Data Platform](https://www.delltechnologies.com/en-us/storage/streaming-data-platform.htm) product, where Pravega is seeing significant growth in the edge, with use cases having thousands of edge deployments.

**Data Flow from Sensors to the Edge and the Cloud using Pravega**<br>
https://cncf.pravega.io/blog/2021/03/23/data-flow-from-sensors-to-the-edge-and-the-cloud-using-pravega/

Pravega’s ZooKeeper Operator continues to see broad adoption, used by [Banzai Cloud Kafka Operator](https://banzaicloud.com/docs/supertubes/kafka-operator/install-kafka-operator/#install-zookeeper), [Apache Solr Operator](https://github.com/apache/solr-operator/blob/master/example/dependencies/zk_operator.yaml), [Mesosphere](https://docs.d2iq.com/dkp/kommander/2.1/workspaces/applications/catalog-applications/dkp-applications/zookeeper-operator/), [Tencent BlueKing Container Service](https://github.com/Tencent/bk-bcs/blob/master/install/kubernetes/zookeeper-operator.yaml), [KOARCH.de](https://github.com/janstrohschein/KOARCH/tree/master/Big_Data_Platform/Kubernetes/Kafka_Broker#zookeeper), and has seen contributions from developers at Fortune 500 & Top 100 Fastest Growing companies.

## Project Performance
> How has the project performed against its goals since the last review? (We won’t penalize you if your goals changed for good reasons.)

**2021 Milestones:**
- Open governance instituted
- Frequent major releases containing new features and new integrations
- External open source contributions on the rise
- Community outreach programs established

### Open Governance

Committers have been organized into 4 teams: Core, Operators, Ecosystem, and Integration & Tools. Pravega committer teams manage a total of 43 repositories. Open governance has been instituted following a model inspired by Apache projects: https://github.com/pravega/.github/blob/main/governance.md

### Releases & Features

Pravega has made steady progress closely following its published roadmaps through regular major and minor releases: 0.8.1, 0.9.0, 0.9.1, 0.10.0, 0.10.1 and 0.10.2.

Features enabled by these releases include performance improvements, Simplified Long Term Storage plugin system, Consumption Based Retention policies, Java 11 support, improved authorization, Rust and Python clients, improved test reliability, Stream Tags, Key Value Tables 2.0, Health Check, CLI improvements, TLS 1.3.

### Open Source

Open source drivers for Pravega have been added to OpenMessaging Benchmark, Debezium, and Project Alvarium. The Akka Alpakka library has expanded to support Key Value Tables. A new ZIO library was prototyped. Connectors to Presto and Trino were released.

**OpenMessaging Benchmark Pravega Driver**<br>
PR: https://github.com/openmessaging/benchmark/pull/199<br>
README: https://github.com/openmessaging/benchmark/tree/master/driver-pravega#pravega-benchmarks

**Debezium Server Pravega Sink**<br>
PRs: https://github.com/debezium/debezium/pulls?q=is%3Apr+author%3Aderekm+is%3Aclosed<br>
Docs: https://debezium.io/documentation/reference/operations/debezium-server.html#_pravega

**Project Alvarium Stream Provider (Alvarium is a Stage 1 LF Edge project)**<br>
Website: https://www.lfedge.org/projects/alvarium/<br>
Code: https://github.com/project-alvarium/alvarium-sdk-java/blob/main/src/main/java/com/alvarium/streams/PravegaStreamProvider.java

**Alpakka Pravega Connector Key Value Tables support**<br>
PR: https://github.com/akka/alpakka/pull/2651<br>
Docs: https://doc.akka.io/docs/alpakka/current/pravega.html (docs will reflect KVT support in Alpakka’s next release)

**ZIO Pravega Connector**<br>
GitHub: https://github.com/cheleb/zio-pravega<br>
README: https://github.com/cheleb/zio-pravega/blob/master/README.md#pravega-zio

**Presto and Trino Connectors**<br>
GitHub: https://github.com/pravega/presto-connector<br>
Getting Started: https://github.com/pravega/presto-connector/tree/main/getting-started

### Community Outreach

Pravega Community meetings were founded under Cloud Native Community Groups:<br>
https://community.cncf.io/pravega-community/

Four community meetings overviewing new features and connectors were held between March and September:<br>
https://www.youtube.com/playlist?list=PLgODz_jrU0amcesAG6NngJxWVWT1znYYX

One community meeting was held in China:<br>
https://www.youtube.com/watch?v=MUyuEcKPBzg

Meetup presentations are providing content for Pravega’s YouTube channel:<br>
https://www.youtube.com/channel/UCiEkLCWdnjSpub76ZEZcdqw

Dell QE had an engagement with LitmusChaos community to exhibit how fault injection tests are used to validate Pravega:<br>
https://youtu.be/EU_g8cKa1G8

Dell China and PingCAP collaborated on an evaluation of Change Data Capture using Pravega, TiDB & Flink SQL:<br>
https://asktug.com/t/topic/92873<br>
https://pingcap.medium.com/building-a-real-time-data-warehouse-with-tidb-and-pravega-a44fba92b3fa

We joined Presto Technical Steering Committee and Trino Community Broadcast meetings to review the new Pravega Presto/Trino connector:<br>
Presto TSC meeting: https://youtu.be/4EZrZ4vaAys<br>
Trino Community Broadcast: https://trino.io/episodes/28.html

### Blog posts

In 2021, the following blog posts were published:

**When Speed meets Parallelism – Pravega performance under parallel streaming workloads**<br>
https://cncf.pravega.io/blog/2021/03/10/when-speed-meets-parallelism-pravega-performance-under-parallel-streaming-workloads/

**Data Flow from Sensors to the Edge and the Cloud using Pravega**<br>
https://cncf.pravega.io/blog/2021/03/23/data-flow-from-sensors-to-the-edge-and-the-cloud-using-pravega/

**Pravega Flink Connector 101**<br>
https://cncf.pravega.io/blog/2021/11/01/pravega-flink-connector-101/

### Hackathons

Unbounded Hackathon was held in English for US/APAC:<br>
https://unboundedhackathon.com/

ITMO University hackathon held in Russia:<br>
https://news.itmo.ru/en/announce/65679/

Pravega China Hackathon was held in China:<br>
https://mp.weixin.qq.com/s/JvLDz4sTq32hOQcnyb0IDg

Flink Forward Asia Hackathon (Flink & Pravega Hackathon), Pravega community was the major sponsor:<br>
https://tianchi.aliyun.com/competition/entrance/531936/introduction

### Conference presentations

Asia Innovation Summit presentation:<br>
https://www.youtube.com/watch?v=l4uXgHOn-xQ

NASSCOM presentation:<br>
https://www.youtube.com/watch?v=gspgqAFJmGQ

StrangeLoop presentation:<br>
https://www.youtube.com/watch?v=AI1M7Wr4_6w

Flink Forward Asia (FFA) 2021, one keynote session and one session on Flink ecosystem from Pravega was presented:<br>
https://flink-forward.org.cn/

“Flink & Pravega” technical whitepaper, a certified joint solution between Flink and Pravega, was published in FFA 2021:<br>
https://flink-learning.org.cn/article/detail/4c6f8c1bf87f3a83bac266690f07a0bb
https://flink-learning.org.cn/article/detail/5a7e33f36930c5f6da328aea88102fb9

PingCAP DevCon 2021 presentation:<br>
https://www.bilibili.com/video/BV1D64y1W71M/

## Current Goals
> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

Working on new features and tools in the 0.11 release and beyond. https://github.com/pravega/pravega/blob/master/documentation/src/docs/roadmap.md

Focusing on popularizing Pravega and increasing its adoption. Generating mixed-media content at all skill levels, especially beginner and advanced material. Disseminating content more broadly across social media by leveraging individual networks.

New integrations to broaden Pravega’s applicability as a component in off-the-self solutions.

## Help with Goals
> How can the CNCF help you achieve your upcoming goals?

There are a few options:
- **Streaming events**: Having a Day 0 event in the 2022 KubeCon series, “Cloud Native Streaming Day,” with other streaming-oriented projects: Pravega, Strimzi, NATS, gRPC, CloudEvents, etc;
- **Documentation**: Support for content generation of technical documentation. We are converting documentation from MkDocs to Docusaurus, combining docs across several repos into a single docs site, allowing for advanced formatting such as doc variables (e.g., version numbers). As we continue to improve our docs site, additional assistance is desirable;
- **Test infrastructure**: We need always-on Kubernetes clusters to run our system tests, longevity tests and Litmus Chaos tests in the open. Pravega codebase includes a suite of system tests, but these are presently used to confirm releases internally. Whether via Prow infra or the cncf.ci initiative, assistance in gaining access to test execution resources at the CNCF would help us move these validations into the public realm.

## Incubation
> Do you think that your project meets the criteria for incubation?

No, only one qualifying end-user presently.
