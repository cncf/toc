# Strimzi Sandbox Annual Review

Strimzi is a project focused on running Apache Kafka on Kubernetes. Apache Kafka is a leading platform for building real-time data pipelines. It provides support for high-throughput / low-latency messaging, as well as sophisticated development options that cover all the stages of a distributed data streaming pipeline, from ingestion, through storage to processing. But running it on Kubernetes can be complex and tedious. Strimzi is making it easy by using the operator pattern and declarative deployments. It provides operators not just for deploying Kafka clusters, but also for managing topics, users or connectors to make sure running Apache Kafka feels Kube-native.

Strimzi also integrates with many other CNCF projects such as:
* Kubernetes
* Prometheus for metrics and monitoring
* Open Policy Agent for authorization
* KEDA for autoscaling of Kafka clients
* Jaeger and Open Tracing for tracing

Strimzi can be also installed using Helm charts or using OperatorHub.io which is part of the Operator Framework project.

Strimzi joined CNCF Sandbox in September 2019 ([slides from TOC presentation proposing Strimzi to CNCF Sandbox](https://docs.google.com/presentation/d/1bijEpuwaaa6jR1D5PAjyW731-j6Xc1TFHJuUh_FwwK8/edit#slide=id.g5628e016b1_46_346)).

## Contribution activity

During the last year, we managed to keep or even increase the pace of the project. We also managed to regularly release new versions. Some of the statistics highlights are:
* [Get 55 new contributors (PRs)](https://strimzi.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)
* [Get contributions (review, comment, commit, issue or PR) from 35 different companies](https://strimzi.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions)
* [On average 27 PRs per week (compared to 19 the year before)](https://strimzi.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)
* [Significantly increase the number of stargazers and forks of our GitHub repositories](https://strimzi.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now)
* The [#strimzi Slack channel](https://cloud-native.slack.com/archives/CMH3Q3SNP) now has over 550 members who create a vibrant community with many questions and discussions

Despite the the number of new contributors, most of the contributions are still carried by the core contributors from Red Hat. We would like to have more people contribute regularly to the project and not just occasionally with one-of PRs.

All Strimzi GitHub metrics can be found in the [CNCF DevStats dashboards](https://strimzi.devstats.cncf.io/d/8/dashboards).

## Maintainers

Strimzi has currently 5 maintainers from two companies (Red Hat and IBM). Two new maintainers were added since joining CNCF. The full maintainer list can be found in our [GitHub repository](https://github.com/strimzi/governance/blob/master/MAINTAINERS).

Additionally to the maintainers, 3 other contributors have _Write_ permissions to individual repositories and act as owners for the sub-projects within these repositories.

## Adoption

Strimzi project sees growth in adoption both in terms of production users as well as in terms of companies engaging with the project via issues, PRs or Slack channels. The list of our adopters can be found in our [GitHub repository](https://github.com/strimzi/strimzi-kafka-operator/blob/master/ADOPTERS.md) and includes end-users such as:

* [Axual](https://axual.com/)
* [Grupo MASMOVIL](https://www.grupomasmovil.com/)
* [Helvetia](https://helvetia.com/)
* [Marlow Navigation](https://marlow-navigation.com/)
* [SBB CFF FFS](https://www.sbb.ch/en/home.html)
* [Swisscom](https://www.swisscom.ch/)

Strimzi is also adopted by several vendors and used in their products:
* [IBM](https://www.ibm.com/cloud/event-streams)
* [Lightbend](https://cloudflow.io/)
* [Red Hat](https://www.redhat.com/en/resources/amq-streams-datasheet)

_(While Red Hat is owned by IBM, Strimzi is currently used in two separate product offerings - one from IBM and one from Red Hat)_

## Past Goals

This is our first annual review, so there are no recorded goals from the previous one. But our general goals and expectations when joining CNCF Sandbox were:
* Attract more contributors and grow the community thanks to the independent governance of CNCF
* Attract more users through the increased visibility

These goals are being fulfilled as seen from the new users and new contributors statistics.

## Current Goals

For the next year, we would like to achieve the following goals:
* Keep attracting more users to Strimzi and Apache Kafka on Kubernetes
* Keep broadening the contributor base with focus on trying to get also more regular and long-term committers. This should also lead to adding more maintainers from other companies than just IBM and Red Hat.
* First release of the Strimzi UI
* Release Strimzi 1.0. While we see Strimzi as production-ready and have production users, we are still on `0.x` versions. We should move to 1.0 during the next year. We plan to work on the following features before releasing 1.0:
    * Improve the use of Cruise Control for automatic performance tuning of the Kafka clusters
    * Move our main Custom Resources to `v1`
    * Remove Zookeeper from Apache Kafka clusters managed by Strimzi (ongoing effort in Apache Kafka community)

_(High level roadmap is available as a [project](https://github.com/orgs/strimzi/projects/1) in our GitHub organization.)_

## How can the CNCF help?

According to the feedback from many of our users and contributors, the speaking and marketing opportunities are very helpful for growing our user and contributor base.

## Incubation

We think that we match the incubation criteria regarding adoption and contributions and healthy number of committers for a project of our size. Assuming the TOC members agree, we would like to start working on incubation during the next year.
