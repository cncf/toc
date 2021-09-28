# Strimzi Sandbox Annual Review

Strimzi is a project focused on running Apache Kafka on Kubernetes. Apache Kafka is a leading platform for building real-time data pipelines. It provides support for high-throughput / low-latency messaging, as well as sophisticated development options that cover all the stages of a distributed data streaming pipeline, from ingestion, through storage to processing. But running it on Kubernetes can be complex and tedious. Strimzi makes it easy by using the operator pattern and declarative deployments. It provides operators not just for deploying Kafka clusters, but also for managing topics, users or connectors to make sure running Apache Kafka feels Kube-native.

Strimzi also integrates with many other CNCF projects such as:
* Kubernetes
* Prometheus for metrics and monitoring
* Open Policy Agent for authorization
* KEDA for autoscaling of Kafka clients
* Jaeger and Open Tracing for tracing

Strimzi can be also installed using Helm charts or using OperatorHub.io which is part of the Operator Framework project.

Strimzi joined CNCF Sandbox in September 2019 ([slides from TOC presentation proposing Strimzi to CNCF Sandbox](https://docs.google.com/presentation/d/1bijEpuwaaa6jR1D5PAjyW731-j6Xc1TFHJuUh_FwwK8/edit#slide=id.g5628e016b1_46_346)). This is our second annual review. The annual review from 2020 can be found [here](https://github.com/cncf/toc/blob/main/reviews/2020-strimzi-annual.md).

## Contribution activity

During the last year, we continued with our release pace of one release every one or two months with 6 minor and 3 patch releases. We also managed to keep the development pace of the project.

Some of the statistics confirming it are:
* [Get 56 new contributors (PRs)](https://strimzi.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)
* [Get contributions (review, comment, commit, issue or PR) from 71 different companies](https://strimzi.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions)
* [On average 24 PRs per week](https://strimzi.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)
* [Continued increase of stargazers and forks of our GitHub repositories](https://strimzi.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now)
* The [#strimzi Slack channel](https://cloud-native.slack.com/archives/CMH3Q3SNP) now has over 1200 members who create a vibrant community with many questions and discussions
* Our [YouTube channel](https://www.youtube.com/c/Strimzi) has now over 250 subscribers and our [Twitter account](https://twitter.com/strimziio) over 1600 followers
* Published over 15 blog posts and several videos

Despite the growing number of new contributors, most of the contributions are still carried by the core contributors from Red Hat. We would like to have more people contribute regularly to the project and not just occasionally with one off PRs.

All Strimzi GitHub metrics can be found in the [CNCF DevStats dashboards](https://strimzi.devstats.cncf.io/d/8/dashboards).

## Maintainers

Strimzi has currently 7 maintainers from two companies (Red Hat and IBM). Two new maintainers were added since the last annual review. The full maintainer list can be found in our [GitHub repository](https://github.com/strimzi/governance/blob/master/MAINTAINERS).

Additionally to the maintainers, 8 other contributors have _Write_ permissions to individual repositories and act as owners for the sub-projects within these repositories. The full list can be found in our [GitHub repository](https://github.com/strimzi/governance/blob/main/COMPONENT-OWNERS) as well.

## Adoption

Strimzi project sees growth in adoption both in terms of production users as well as in terms of companies engaging with the project via issues, PRs or Slack channels. The list of our adopters can be found in our [GitHub repository](https://github.com/strimzi/strimzi-kafka-operator/blob/master/ADOPTERS.md) and includes end-users such as:

* [Axual](https://axual.com/)
* [Grupo MASMOVIL](https://www.grupomasmovil.com/)
* [Helvetia](https://helvetia.com/)
* [Marlow Navigation](https://marlow-navigation.com/)
* [SBB CFF FFS](https://www.sbb.ch/en/home.html)
* [Swisscom](https://www.swisscom.ch/)
* [Smart.iX](https://smart-ix.ai/)

Strimzi is also adopted by several vendors and used in their products:
* [IBM](https://www.ibm.com/cloud/event-streams)
* [Lightbend](https://cloudflow.io/)
* [Red Hat](https://www.redhat.com/en/resources/amq-streams-datasheet)

_(While Red Hat is owned by IBM, Strimzi is currently used in two separate product offerings - one from IBM and one from Red Hat)_

## Last years achievements

There are several things we achieved during the last year which we think are worth mentioning:
* We successfully migrated the Strimzi project to the new versions of the Kubernetes CRD or Ingress APIs. Despite the related effort we managed to do it in time to be ready for the Kubernetes 1.22 release.
* While our main focus is on running the Apache Kafka server components on Kubernetes, we also created and released several new components / sub-projects which make it easier to run Kafka clients on Kubernetes. This includes the Apache Kafka configuration providers which make it easier to use data from Kubernetes Config Maps and Secrets or from environment variables in configuration of Kafka clients.
* We maintain several components which are used by Apache Kafka users even outside of Strimzi. This includes for example the [library for OAuth authentication](https://github.com/strimzi/strimzi-kafka-oauth) in Apache Kafka brokers and clients, the [extensions](https://github.com/strimzi/mirror-maker-2-extensions) for Apache Kafka Mirror Maker 2 and the [quotas plugin](https://github.com/strimzi/kafka-quotas-plugin) for Apache Kafka. We also started work on an _encryption proxy_ for topic / message level encryption to improve the security of Apache Kafka based infrastructure or a Canary tool to measure availability of the Kafka clusters.
* We established the [Strimzi proposals](https://github.com/strimzi/proposals) process for proposing and discussing important changes or new features.
* We improved the flexibility of our operators to make them more configurable. We also improved significantly the experience for users of Kafka Connect by adding support for automatically adding new connectors to integrate Apache Kafka with other systems.

## Past Goals

We managed to achieve some of the goals we set in the past annual review:
* We managed to continue development and growth of the Strimzi project and do regular releases
* We attracted more users
* We attracted new contributors

There were also some goals we didn't achieve:
* While we attracted more contributors, most of them are not regular long-term contributors
* We planned to release Strimzi 1.0 after Apache ZooKeeper dependency is removed from Apache Kafka. This change in Apache Kafka itself will bring significant changes to our APIs and it is something we want to incorporate before the 1.0 release. However, the Apache ZooKeeper dependency removal is delayed in the Apache Kafka project. So while this remains our plan, we didn't manage to complete in the previous year and it remains to be planned for next year.

In the last year's goals, we also planned to work on a Strimzi UI. After careful considerations, we decided that at this time we do not have enough contributors with the right skill-set to work on the UI. So we decided to postpone this goal.

## Current Goals

For the next year, we would like to achieve the following goals:
* Keep attracting more users to Strimzi and Apache Kafka on Kubernetes
* Keep broadening the contributor base with a focus on trying to get also more regular and long-term committers. This should also lead to adding more maintainers from other companies than just IBM and Red Hat.
* Release Strimzi 1.0. While we see Strimzi as production-ready and have production users, we are still on `0.x` versions. We should move to 1.0 during the next year. We plan to work on the following features before releasing 1.0:
    * Improve the use of Cruise Control for automatic performance tuning of the Kafka clusters
    * Move our main Custom Resources to `v1`
    * Remove Zookeeper from Apache Kafka clusters managed by Strimzi (ongoing effort in Apache Kafka community)
* Releasing Strimzi Canary tool for monitoring availability of Kafka clusters
* Better integrate tools such as Vault or Cert-manager to improve the support for custom TLS certificates / certification authorities
* Incubation

_(High level roadmap is available as a [project](https://github.com/orgs/strimzi/projects/1) in our GitHub organization.)_

## How can the CNCF help?

According to the feedback from many of our users and contributors, the speaking and marketing opportunities are very helpful for growing our user and contributor base. For example the Office Hours sessions held during the KubeCon conferences are very popular among existing and new users and help us to stay in contact with users and contributors.

## Incubation

We think that we match the incubation criteria regarding adoption and contributions and healthy number of committers for a project of our size. Assuming the TOC members agree, we plan to start working on incubation materials / proposal in the upcoming weeks.
