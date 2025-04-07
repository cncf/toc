# Distributed Application Runtime CNCF Proposal

## Name of project

Distributed Application Runtime (Dapr)

## Project description

The Distributed Application Runtime (Dapr) is a portable, event-driven runtime that makes it easy for developers to build resilient, microservice stateless and stateful applications that run on the cloud and edge. 

## Overview

Using Dapr you can build microservice applications using any language, any framework, and run them anywhere.

Today developers are comfortable with web + database application architectures but not with microservice application architectures which are inherently distributed. It’s hard to become a distributed systems expert, nor should they have to. Developers want to focus on business logic, while leaning on the platforms to imbue their applications with scale, resiliency, maintainability, elasticity and the other attributes of cloud-native architectures.

Dapr codifies the best practices for building microservice applications into open, independent, building blocks exposed via a consistent API that enables you to build portable applications with the language and framework of your choice. Each building block is completely independent and you can use one, some, or all of them in your application.

Dapr is platform agnostic meaning you can run your applications locally, on any Kubernetes cluster, and other hosting environments. This enables you to build microservice applications that can run on the cloud and edge.

### About Dapr

Dapr exposes its APIs as a sidecar architecture, either as a container or as a process, not requiring the application code to include any Dapr runtime code. This makes integration with Dapr easy from other runtimes, as well as providing separation of the application logic.  

Dapr can be hosted in multiple environments, including self hosted for local development, a group of VMs, and Kubernetes. Dapr has a control plane that is deployed and integrated with the hosting environment, with Kubernetes currently being the preferred choice for production workloads.  

Dapr released as a v1.0 project in Feb 2021 and has achieved the stability and enterprise readiness to be designated production ready. Dapr is used in production today by multiple companies. 

### CNCF Due Diligence (DD) document
The [Dapr CNCF Due Diligence document](https://docs.google.com/document/d/1GknuJk8Z26KhvFJiblbMSopWbX27nihpxkEG0_TpkCw/edit) provides the necessary details to support submission of Dapr as an Incubations project to CNCF via the App-Delivery SIG. Refer to this document to answer required questions on accepting Dapr into CNCF Incubation status.

### Features

Dapr consists of a high-performance runtime, CLI, language specific SDKs (Java, Python, Go, .NET, PHP, C++, Javascript) and pluggable components to integrate with a range of different technologies including storage, queues, event brokers etc.

Dapr has a set of [building block APIs](https://docs.dapr.io/concepts/building-blocks-concept/) that use pluggable [components](https://docs.dapr.io/concepts/components-concept/) to provide a developer abstraction, code portability and hosting independence. These are the building blocks that Dapr provides.

| Building Block API | Description |
|----------------|-------------|
| Service-to-service invocation | Resilient service-to-service invocation enables method calls, including retries, on remote services wherever they are located in the hosting environment.
| State management| State management is used for storing key/value pairs, and enables long running, highly available, stateful services.
| Publish and subscribe | Publishes events and subscribes to topics using Cloud Events, and enables scalable, event-driven services.
| Resource bindings | Resource bindings with code triggers enables event-driven architectures by receiving and sending events to and from any external source such as databases, queues, file systems, and services like Twitter, Twilio etc.
| Actors | A pattern for stateful and stateless objects that makes concurrency simple with method and state encapsulation. Can be used to build IoT and workflow applications.
| Observability | Dapr emit metrics, logs, and traces to debug and monitor both Dapr and user applications. Distributed tracing uses W3C Trace Context and Open Telemetry supports all monitoring tools including Prometheus and Grafana. 
| Secrets | Secrets management integrates with public cloud and local secret stores to retrieve the secrets for use in application code.

The Dapr CLI provides a simple local development experience to develop, build and test cloud native applications. 

Although Dapr can be used with localhost HTTP or gRPC calls, language SDKs provide a more familiar developer experience and integration with language specific tooling.

## Statement on alignment with CNCF charter mission

Dapr is aligned with the CNCF mission to make cloud native computing ubiquitous for developers. Dapr enables developers to build scalable, cloud native applications faster and  makes them portable across different clouds, bringing true multi-cloud experience to application development. Dapr is a community-led project that has an open governance charter with a large and rapidly growing community. Dapr utilizes and integrates with several existing CNCF technolgies and is an ideal fit with the cloud native serverless ecosystem which is why we are seeking to have Dapr join the CNCF.

## Comparison with similar projects (inside or outside the CNCF), including what differentiates this project

Similar projects outside of CNCF:

- **[Micro](https://github.com/micro/micro)** is a microservices framework that has some similarities to Dapr. A significant difference is that Dapr is a set of APIs that can be used from any language or framework and is incremental in its adoption. Most microservice frameworks today like Micro are restricted to certain languages, require developers to rewrite their code and do not expose a set of common APIs that can be extended by the community.

- Service meshes such as CNCF **[Linkerd](https://github.com/linkerd/linkerd)** and non-CNCF **[Istio](https://github.com/istio/istio)** both have network routing, security and observability that Dapr also provides. The primary difference between Dapr and service meshes is that Dapr operates at the application level by exposing APIs for developers, whereas service meshes operate at the infrastructure networking layer.

Dapr integrates with many CNCF projects, including;

- **[CloudEvents](https://github.com/cloudevents/spec)** for all publish/subscribe events.
- **[gRPC](https://github.com/grpc)** for high-performance remote procedure calls (RPC).
- **[SPIFFE](https://github.com/spiffe/spiffe)** for identifying and securing communications between application services.
- **[Open Telemetry](https://github.com/open-telemetry)** to generate, collect, and export telemetry data.
- **[Helm](https://github.com/helm/helm)** used to deploy Dapr's control plane to Kubernetes.
- **[Prometheus](https://github.com/prometheus)** to collect and analyse metrics. 
- **[Jaeger](https://github.com/jaegertracing/jaeger)** for distributed tracing.
- **[NATS](https://github.com/nats-io)** as a high performance open source messaging system for cloud native applications.
- **[KEDA](https://github.com/kedacore/keda)** for fine-grained autoscaling on Kubernetes.

## Sponsor from TOC

Lei Zhang (Harry)

## Preferred maturity level

Incubation

Having built Dapr in conjunction with customers, adopters and the community who are running production workloads, Dapr is ready for incubation level maturity.  

## License

Currently MIT (https://github.com/dapr/dapr/blob/master/LICENSE) - happy to relicense to Apache2.0 at donation.

## Source control
Dapr is a GitHub organization with several repos 

GitHub: https://github.com/dapr

## External dependencies (including licenses)

Dapr uses Go 1.16 with the following dependencies:

- contrib.go.opencensus.io/exporter/prometheus Apache 2.0
- contrib.go.opencensus.io/exporter/zipkin Apache 2.0
- github.com/AdhityaRamadhanus/fasthttpcors MIT
- github.com/PuerkitoBio/purell BSD-3 Clause
- github.com/agrea/ptr MIT
- github.com/cenkalti/backoff/v4 MIT
- github.com/dapr/components-contrib MIT
- github.com/fasthttp/router BSD-3 Clause
- github.com/fsnotify/fsnotify BSD-3 Clause
- github.com/ghodss/yaml MIT
- github.com/golang/protobuf BSD-3 Clause
- github.com/google/go-cmp BSD-3 Clause
- github.com/google/uuid BSD-3 Clause
- github.com/gorilla/mux BSD-3 Clause
- github.com/grpc-ecosystem/go-grpc-middleware Apache 2.0
- github.com/hashicorp/go-hclog MIT
- github.com/hashicorp/go-msgpack BSD-3 Clause
- github.com/hashicorp/raft MPL-2.0
- github.com/hashicorp/raft-boltdb MPL-2.0
- github.com/json-iterator/go MIT
- github.com/kelseyhightower/envconfig MIT
- github.com/minio/blake2b-simd Apache 2.0
- github.com/mitchellh/mapstructure MIT
- github.com/nats-io/nats-streaming-server Apache 2.0
- github.com/openzipkin/zipkin-go Apache 2.0
- github.com/phayes/freeport BSD-3 Clause
- github.com/pkg/errors BSD-2 Clause
- github.com/prometheus/client_golang Apache 2.0
- github.com/prometheus/client_model Apache 2.0
- github.com/prometheus/common Apache 2.0
- github.com/sendgrid/rest MIT
- github.com/sirupsen/logrus MIT
- github.com/stretchr/testify MIT
- github.com/valyala/fasthttp MIT
- go.opencensus.io Apache 2.0
- go.opentelemetry.io/otel Apache 2.0
- go.uber.org/atomic MIT
- google.golang.org/genproto Apache 2.0
- google.golang.org/grpc Apache 2.0
- google.golang.org/protobuf BSD-3 Clause
- gopkg.in/yaml.v2 Apache 2.0
- k8s.io/api Apache 2.0
- k8s.io/apiextensions-apiserver Apache 2.0
- k8s.io/apimachinery Apache 2.0
- k8s.io/cli-runtime Apache 2.0
- k8s.io/client-go Apache 2.0
- k8s.io/code-generator Apache 2.0
- k8s.io/klog Apache 2.0
- k8s.io/metrics Apache 2.0
- sigs.k8s.io/controller-runtime Apache 2.0
        
## Initial committers (how long working on project, companies they represent)

Current maintainers:
* Yaron Schneider (Microsoft)
* Mark Fussell (Microsoft)
* Rob Landers (Automattic)
* Sky Xiaojian (Alibaba Cloud)
* Artur Souza (Microsoft)
* Ryan Nowak (Microsoft)
* Mukundan Sundar (Microsoft)
* Charlie Stanley (Microsoft)
* Aaron Crawfis (Microsoft)
* Ori Zohar (Microsoft)
* Mark Chmarny (Apple) 

Other initial committers who are in Alumni status:
* Haishi Bai (Microsoft)
* Aman Bhardwaj (Microsoft)
* Shalabh Mohan Shrivastava (Microsoft)
* Leon Mai (Microsoft)
* Young Bu Park (Microsoft)

This link https://github.com/orgs/dapr/teams captures the list of Dapr's current maintainers and approvers.

Dapr is completely developed in the open with transparent community ownership, governance, and contributions. Microsoft has the majority of maintainers currently, however there are maintainers and approvers from the community for several Dapr repos and we are actively adding more. Each repo has a code owners file, for example https://github.com/dapr/dapr/blob/master/CODEOWNERS

## Infrastructure requests (CI / CNCF Cluster)

No infrastructure requests at this time.

## Communication channels (slack, irc, mailing lists)
See the [Dapr community repo](https://github.com/dapr/community) specifically [communication ](https://github.com/dapr/community#communication)

[Dapr Discord](https://discord.com/invite/ptHhX6jc34) is the primary community engagement and Q&A. There is a twice-monthly community call for open Q&A, updates, demos and roadmap.

## Issue tracker
Issues are tracked in each of the Dapr repos https://github.com/dapr

See https://github.com/dapr/dapr/issues

## Website

- Website https://dapr.io/
- Documentation https://docs.dapr.io/ 
- Blog https://blog.dapr.io/posts 

## Release methodology and mechanics

Dapr reached [v1.0 maturity in Feb 2021](https://blog.dapr.io/posts/2021/02/17/announcing-dapr-v1.0/) having worked closely with customers such as Zeiss, Ignition Group and Morgan Stanley as well as adopters such as Alibaba. 

### Milestone release plan
Each milestone release plan is tracked with an issue, for example these are the plans from the v1.0 and v1.1 releases.
- [Milestone release plan v1.0](https://github.com/dapr/dapr/issues/2346)
- [Milestone release plan v1.1](https://github.com/dapr/dapr/issues/2843) 

Each release sets dates and brings together milestones from the maintainers of each of the Dapr repos. Dapr releases on a regular cadence of around every 6 weeks. Current and next milestones for each repo are maintained. See https://github.com/dapr/dapr/milestones

### Milestone release endgame plan
Each release endgame plan is tracked with an issue, for example this is the release endgame plan for v1.0.

 - [Milestone release endgame v1.0](https://github.com/dapr/dapr/issues/2724)

The endgame tracks the steps to cut a supported release, which can go through a number of release candidates (RCs) during the endgame phase. The RCs are opportunities for users in production and community members to raise any issues that they encounter.

Dapr has a versioning policy for releases which is described in the [Versioning policy](https://docs.dapr.io/operations/support/support-versioning/) documentation.

Dapr has a rolling window of supported releases. Dapr's supported release policy is described in the [Supported releases](https://docs.dapr.io/operations/support/support-release-policy/) documentation. 

This [engineering release decision record](https://github.com/dapr/dapr/blob/e6477d9fb05cbc647d17bd4bf2b593a61290853b/docs/decision_records/engineering/ENG-002-Dapr-Release.MD) outlines the methodology and mechanics

## Social media accounts

Twitter: [@daprdev](https://twitter.com/daprdev) <br>
YouTube channel: https://www.youtube.com/channel/UCtpSQ9BLB_3EXdWAUQYwnRA

## Community size and any existing sponsorship

Currently 823 [contributors](https://github.com/cncf/devstats/blob/master/README_K8s.md) over the last 16 months, which has grown rapidly over the last year. This includes committers from Microsoft, Alibaba, Hashicorp, Disney, Zeiss, Morgan Stanley, Citigroup, New Relic, and Dynatrace.

The community has developed and contributed over 70 components for state management (such as Redis, MySQL,and many more), Pub/Sub messaging (such as RabbitMQ, AWS SNS/SQS, Kafka), secret stores (such as Kubernetes secret store, Azure Key Vault, HashiCorp Vault, and more) and both input and output bindings to integrate with services such as SendGrid, Twilio and Twitter.

The Dapr Github repo currently has 11.8k stars.

As for sponsorship, Microsoft currently pays for the project’s GitHub, infrastructure for testing/releases, website hosting and Docker repos.

## Who is currently known to be using the project? Are they using it in production and at what scale? (It may be hard to obtain accurate data for this, but any supporting evidence of usage is helpful)

Dapr has customers running in production with the v1.0 release. Publicly announced ones include; 

- [ZEISS](https://customers.microsoft.com/en-us/story/1336089737047375040-zeiss-accelerates-cloud-first-development-on-azure-and-streamlines-order-processing), an international technology leader in optics and opto-electronics. Their solution is deployed to multiple regions in Microsoft Azure.

- [Ignition Group](https://customers.microsoft.com/en-us/story/1335733425802443016-ignition-group-speeds-development-and-payment-processing-using-dapr-and-azure) a South Africa–based technology business focusing on customer engagement that builds sales-support tools and value-added services for telcos. Their solution is deployed to multiple regions in Microsoft Azure.

- [Roadwork](https://blog.dapr.io/posts/2021/02/09/running-dapr-in-production-at-roadwork/) a company the enables customers to gain actionable insights from their data. 

Others adopter references can be provided who can talk about their projects under a non-disclosure agreement that we have worked with over the last year. 

Dapr is used by Alibaba Cloud in their serverless infrastructure project. 

As part of the Dapr v1.0 release, Dapr focused on [production readiness](https://blog.dapr.io/posts/2021/02/17/announcing-dapr-v1.0/#release-highlights), in areas such as performance, security, High Availability (HA), conformance testing and end-to-end scenario testing.

Dapr has had multiple, independent [security audits](https://docs.dapr.io/concepts/security-concept/#security-audit) from Cure53, CNCF approved security auditor.

## Project logo in svg format

All logos are available in https://github.com/dapr/docs/tree/v1.0/daprdocs/assets/icons

The svg main logo is here: https://github.com/dapr/docs/blob/v1.0/daprdocs/assets/icons/logo.svg

### CNCF Due Diligence (DD) document
The [Dapr CNCF Due Diligence document](https://docs.google.com/document/d/1GknuJk8Z26KhvFJiblbMSopWbX27nihpxkEG0_TpkCw/edit) provides the necessary details to support submission of Dapr as an Incubations project to CNCF.