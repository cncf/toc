# Strimzi Sandbox Annual Review

Strimzi is a project focused on running Apache Kafka on Kubernetes.
Apache Kafka is a leading platform for building real-time data pipelines.
It provides support for high-throughput / low-latency messaging, as well as sophisticated development options that cover all the stages of a distributed data streaming pipeline, from ingestion, through storage to processing.
Running Kafka on Kubernetes can be complex and laborious.
Strimzi makes it easy by using the operator pattern and declarative deployments.
It provides operators not just for deploying Kafka clusters, but also for managing topics, users or connectors to make sure running Apache Kafka feels Kube-native.

Strimzi also integrates with many other CNCF projects such as:
* Kubernetes
* Prometheus for metrics and monitoring
* Open Policy Agent for authorization
* KEDA for autoscaling of Kafka clients
* Jaeger, OpenTracing and OpenTelemetry for tracing

Strimzi can be also installed using Helm charts or using OperatorHub.io which is part of the Operator Framework project.

Strimzi joined CNCF Sandbox in September 2019 ([slides from TOC presentation proposing Strimzi to CNCF Sandbox](https://docs.google.com/presentation/d/1bijEpuwaaa6jR1D5PAjyW731-j6Xc1TFHJuUh_FwwK8/edit#slide=id.g5628e016b1_46_346)). 
This is our third annual review. 
The annual review from 2021 can be found [here](https://github.com/cncf/toc/blob/main/reviews/2021-Strimzi-annual.md).

## Contribution activity

During the last year, we continued with our pace of one release every one or two months.
We also managed to keep the pace of development steady in terms of new contributors and PRs.

The following figures illustrate the continued momentum of the Strimzi project:
* [Over 50 new contributors (PRs)](https://strimzi.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)
* [Get contributions (review, comment, commit, issue or PR) from over 60 different companies](https://strimzi.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions)
* [On average 23 PRs per week](https://strimzi.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)
* [Continued increase of stargazers and forks of our GitHub repositories](https://strimzi.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now)
* A vibrant community of almost 2000 members discussing Strimzi on the [#strimzi Slack channel](https://cloud-native.slack.com/archives/CMH3Q3SNP)
  Over 250 members on the developer and contributor-focused [#strimzi-dev Slack channel](https://cloud-native.slack.com/archives/C018247K8T0)
* Over 500 subscribers on the Strimzi [YouTube channel](https://www.youtube.com/c/Strimzi)
* Almost 2000 followers of the Strimzi [Twitter account](https://twitter.com/strimziio)
* Over 20 [blog posts and instructional videos](https://strimzi.io/blog/) on the Strimzi website describing new features or providing guidance

As in previous years, we were able to attract new contributors.
But only a few of them contribute regularly and most of the contributions are still carried by the core contributors.
We would like to have more people contribute regularly to the project rather than creating the occasional PR.
To improve this, we created a [Join the Strimzi Community](https://strimzi.io/join-us/) section on our website.  
The section describes how the community can contribute and grow.

All Strimzi GitHub metrics can be found in the [CNCF DevStats dashboards](https://strimzi.devstats.cncf.io/d/8/dashboards).

## Maintainers

Strimzi has 9 maintainers from three different companies (Red Hat, CN Group, and IBM). 
Two new maintainers were added since the last annual review. 
The full maintainer list can be found in our [GitHub repository](https://github.com/strimzi/governance/blob/master/MAINTAINERS).

In addition to the maintainers, 7 other contributors have _Write_ permissions to individual repositories and act as owners for the sub-projects within these repositories. 
The full contributor list can be found in our [GitHub repository](https://github.com/strimzi/governance/blob/main/COMPONENT-OWNERS) as well.

## Adoption

The Strimzi project sees growth in adoption in terms of the number of production users and also companies engaging with the project by raising issues or PRs, or using the Slack channels. 
The list of our adopters can be found in our [GitHub repository](https://github.com/strimzi/strimzi-kafka-operator/blob/master/ADOPTERS.md) and includes the following users:

* [Atruvia](https://atruvia.de/)
* [CN Group](https://www.cngroup.dk/)
* [Grupo MASMOVIL](https://www.grupomasmovil.com/)
* [Helvetia](https://helvetia.com/)
* [Marlow Navigation](https://marlow-navigation.com/)
* [SBB CFF FFS](https://www.sbb.ch/en/home.html)
* [Swisscom](https://www.swisscom.ch/)
* [Smart.iX](https://smart-ix.ai/)

Strimzi is also adopted by several vendors and used in their products:
* [Axual](https://axual.com/)
* [IBM](https://www.ibm.com/cloud/event-streams)
* [Lightbend](https://cloudflow.io/)
* [Red Hat](https://www.redhat.com/en/resources/amq-streams-datasheet)

_(While Red Hat is owned by IBM, Strimzi is currently used in two separate product offerings - one from IBM and one from Red Hat)_

## Last years achievements

There are several things we achieved during the last year which we think are worth mentioning:
* Keeping Strimzi up-to-date with the latest Kubernetes releases and API changes.
* Adopting new Kafka versions in a timely fashion.
* Resolving security challenges quickly, such as Log4Shell.
* Adding platform support for AArch64 and s390x.
* Maintaining focus on running the Apache Kafka server components on Kubernetes, whilst creating and maintaining several smaller components / sub-projects to make it easier to run Kafka clients on Kubernetes.
  This year, we're working on supporting Service Bindings with the [Strimzi Access Operator](https://github.com/strimzi/kafka-access-operator) sub-project, and also a Strimzi-based [Apache Kafka test container](https://github.com/strimzi/test-container).

## Past Goals

We managed to achieve some of the goals we set in the past annual review:
* Continuing the development and growth of the Strimzi project with regular releases
* Gaining more users
* Gaining new contributors

There were also some goals we didn't achieve:
* While we attracted more contributors, most of them are not regular long-term contributors
* The removal of the ZooKeeper dependency from Apache Kafka is progressing more slowly than expected.
  This applies both for Apache Kafka itself, but also for readiness of this feature in Strimzi.
  We managed to add some building blocks for this over the last year such as with the introduction of `StrimziPodSets`.
  With `StrimziPodSets`, Strimzi handles the creation and management of pods.  
  But more work remains to be done on this.
  In the past, we decided to have a Strimzi 1.0 version after ZooKeeper was removed.
  This might have been a mistake because the removal of ZooKeeper is taking much longer than initially expected.
  But after careful consideration, we have decided to stick with this plan.

## Current Goals

For the next year, we would like to achieve the following goals:
* Attracting more users to Strimzi and Apache Kafka on Kubernetes.
* Broadening the contributor base with a focus on trying to get also more regular and long-term committers.
  This should also lead to adding more maintainers from companies beyond Red Hat, CN Group, and IBM.
* Continuing with ZooKeeper removal and towards the release of Strimzi 1.0.
* Continuing with other improvements and new features such as:
    * Improved support for automatic performance tuning of the Kafka clusters (based on the [Cruise Control project](https://github.com/linkedin/cruise-control))
    * Supporting Kafka clusters across multiple Kubernetes clusters
    * Supporting different TLS certificate management tools such as cert-manager
* Reaching an _incubation_ level of project maturity.

_(High level roadmap is available as a [project](https://github.com/orgs/strimzi/projects/1) in our GitHub organization.)_

## How can the CNCF help?

According to the feedback from many of our users and contributors, speaking and marketing opportunities are very helpful for growing our user and contributor base. 
For example the Office Hours sessions held during the KubeCon conferences are very popular among existing and new users, and help us to stay in contact with users and contributors.

## Incubation

We think that we match the incubation criteria regarding adoption and contributions.
We also have a healthy number of committers for a project of our size. 
In May, we proposed [Strimzi for Incubation](https://github.com/cncf/toc/pull/842).
The incubation process is currently ongoing with Erin Boyd as the TOC sponsor.
