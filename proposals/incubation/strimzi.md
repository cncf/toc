# Strimzi Incubation Stage Proposal

## Introduction

Strimzi is a project focused on running [Apache Kafka][ak] on [Kubernetes][kube].
Apache Kafka is the _de facto_ standard platform for building real-time data pipelines.
It provides support for high-throughput / low-latency messaging, as well as sophisticated development options that cover all the stages of a distributed data streaming pipeline, from ingestion, through storage, and on to processing.
But running Kafka on Kubernetes can be complex and tedious.
Strimzi makes it easy by using the Kubernetes operator pattern and declarative deployments.
It provides operators not just for deploying Kafka clusters, but also for managing topics, users and connectors to make sure running Apache Kafka feels Kube-native.

Strimzi also integrates with many other CNCF projects such as:

* Kubernetes
* [Prometheus][prom] for metrics and monitoring
* [Open Policy Agent][opa] for authorization
* [KEDA][keda] for autoscaling of Kafka clients
* [Jaeger][jaeger] and [OpenTelemetry][otel] for tracing

In addition, Strimzi is used by [Knative][knative] to provide its Kafka source.

Strimzi can be installed using [Helm][helm] charts or using Operator Lifecycle Manager, which is part of the [Operator Framework][operatorframework] project.
Both the Helm charts and OLM bundles are available on [Artifact Hub][artifacthub].

Strimzi joined CNCF Sandbox in September 2019

## License

The Strimzi project is licensed under [Apache License 2.0](https://github.com/strimzi/strimzi-kafka-operator/blob/main/LICENSE).

## Statement on alignment with the CNCF mission

> The Foundation's mission is to make cloud native computing ubiquitous

While Kafka is _capable_ of horizontal scalability, achieving such things as zero-downtime upgrades, and scaling clusters within Kubernetes without an operator is too high a barrier for many organizations.
This prevents ubiquitous use of Kafka anywhere with cloud native platforms such as Kubernetes.

Strimzi provides declarative APIs for Apache Kafka resources, including broker clusters, Kafka Connect clusters, connectors, and Kafka MirrorMaker instances (for inter-cluster replication).

Thanks to the declarative APIs, Strimzi simplifies the way to deploy a Kafka cluster on Kubernetes as well as helps to automate the day-2 operations like upgrading brokers configuration, Kafka version, and security related issues.

By automating the operational burdens of running Kafka, Strimzi brings this important technology into the hands of anyone using Kubernetes.

## Community

In the Strimzi community, users and developers can interact in several places:

* A public #strimzi channel, on the CNCF Slack workspace, with almost 1700 users.
* Mailing lists.
* Github issues and discussions.
* A public "community call" every other week. The full list with all recordings is [here](https://www.youtube.com/watch?v=wK4k3cAstbQ&list=PLpI4X8PMthYfONZopcRd4X_stq1C14Rtn).
* Forums and events. Strimzi has participated in numerous KubeCon "Office Hours" events.

Strimzi encourages contributions from the community, and has a good record of merging such contributions.


## Incubation Stage Requirements

### Production usage

Strimzi is used in production by a number of organizations (see [ADOPTERS.md](https://github.com/strimzi/strimzi-kafka-operator/blob/main/ADOPTERS.md) for some public references).
Some of the public adopters are presented on the homepage of the official Strimzi [website](https://strimzi.io/).
In addition, Strimzi maintainers are aware of further non-public users of both Strimzi and commercially supported distributions of Strimzi. 

### Committers

Strimzi has currently 7 maintainers from three companies (Red Hat, IBM, and CN Group).
The full maintainer list can be found in our [GitHub repository](https://github.com/strimzi/governance/blob/master/MAINTAINERS).

In addition to the maintainers, 8 other contributors have _Write_ permissions to individual repositories and act as owners for the sub-projects within these repositories.
The full list can be found in our [GitHub repository](https://github.com/strimzi/governance/blob/main/COMPONENT-OWNERS) as well.

### Commits and merged contributions

Current status of the Strimzi project:

* 2133 commits in the last 12 months
* 411 contributors in the last 12 months
* more than 4k stars and ~1.2 forks on github.

See our [devstats](https://all.devstats.cncf.io/d/54/project-health-table?orgId=1&var-repogroup_name=Strimzi) for more details.

### Versioning scheme

Strimzi uses a 1–2 month release schedule for the "core" operators component.
All Strimzi component releases use [semantic versioning](https://semver.org/).
The project's release version in described in [RELEASES.md](https://github.com/strimzi/governance/blob/main/RELEASES.md).

### Security processes

Security is taken seriously by the Strimzi team. Any vulnerabilities should be reported to Strimzi maintainers privately, to minimize potential attacks against current users of Strimzi before they are fixed.
Reported vulnerabilities will be investigated and patched on the next patch (or minor) release as soon as possible.
The project's security processes are described in [SECURITY.md](https://github.com/strimzi/.github/blob/main/SECURITY.md).

## Future Plans

### Functionality

There are several major features in development or under consideration that will further increase Strimzi capabilities.

Kafka itself is migrating from ZooKeeper for metadata management to its own Raft-based quorum, known as [KRaft][kip-500].
For Strimzi the challenge is ensure that users can access this new functionality and provide a migration path from existing ZooKeeper-based clusters.

To support more sophisticated use cases, we want to improve the manageability of Kafka clusters by integrating [Cruise Control][cc] more deeply within Strimzi.
We also intend to investigate enabling Kafka clusters to be spread over multiple Kubernetes clusters.

Finally, we see that there is a lot of value in leveraging the community's operational experience by providing common tooling, standard operating procedures, and documentation.

The Strimzi roadmap of new future features is available [here](https://github.com/orgs/strimzi/projects/1).

### Usage and community

In the short to medium term we see two main drivers of adoption. 

Strimzi is already seeing significant production usage across many companies, but this is at odds with its current 0.x version number.
Announcing Strimzi 1.0 will help signal the software's maturity, especially to those who have been held back by an apparent lack of production-readiness. 

A second driver is the increased visibility of the project if Strimzi is accepted into CNCF incubation.
Taking Knative as an example, we would like to more actively promote Strimzi as a building block for other cloud-native technologies with data streaming requirements.

We hope that expanding the number and diversity of organizations using Strimzi will have a corresponding increase on the number of committers and contributions being made to the project.

[ak]: https://kafka.apache.org/
[kube]: https://kubernetes.io/
[prom]: https://prometheus.io/
[opa]: https://www.openpolicyagent.org/
[keda]: https://keda.sh/
[jaeger]: https://www.jaegertracing.io/
[otel]: https://opentelemetry.io/
[knative]: https://knative.dev/
[helm]: https://helm.sh/
[operatorhub]: https://operatorhub.io/
[operatorframework]: https://operatorframework.io/
[kip-500]: https://cwiki.apache.org/confluence/display/KAFKA/KIP-500%3A+Replace+ZooKeeper+with+a+Self-Managed+Metadata+Quorum
[cc]: https://github.com/linkedin/cruise-control
[artifacthub]: https://artifacthub.io/