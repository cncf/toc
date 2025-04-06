# Propose OpenTelemetry to Incubation Stage 

## Table of Contents

- [Background](#background)
- [Recap: What is OpenTelemetry](#recap-what-is-opentelemetry)
- [Progress Since Sandbox](#progress-since-sandbox)
- [Incubation Stage Requirements](#incubation-stage-requirements)
- [Statement on Alignment with the CNCF Mission](#statement-on-alignment-with-the-cncf-mission)
- [Comparison with Other CNCF Observability Projects Projects](#comparison-with-other-cncf-observability-projects)
- [Future Plans](#future-plans)

## Background
OpenTelemetry became a sandbox project in May 2019. 
- [Slides](https://docs.google.com/presentation/d/1roUFtkiVw25N7nRuUnMa-ziNhb6vzhZAe43jpBfiDiw/edit#slide=id.g5724df7a45_0_258)
- [Proposal](https://github.com/cncf/toc/blob/f4ecf1bf391867d0840434b638cd9ef6cb2f57fe/proposals/sandbox/opentelemetry.adoc)
- [PR](https://github.com/cncf/toc/pull/233)

## Recap: What is OpenTelemetry
OpenTelemetry is a vendor-neutral telemetry management framework for ensuring cloud-native software is observable. It is a collection of tools, APIs, and SDKs. You can use it to instrument, generate, collect, and export telemetry data consisting of metrics, logs, and traces for analysis in order to understand your software's performance and behavior.

### High-level Architecture

OpenTelemetry is an integrated set of APIs and libraries that generate, collect, and describe telemetry about distributed systems. This data includes distributed traces, metrics, and other signals in the future; it embraces standards for correlation and context propagation, but also supports various formats needed by end users. OpenTelemetry makes it easy to get critical telemetry data out of your services: for each language it offers a set of APIs, libraries, and data specifications, and developers can make use of which components they see fit.

The project is made up of three major components:

- [Specification](https://github.com/open-telemetry/opentelemetry-specification) with the [protos/definitions](https://github.com/open-telemetry/opentelemetry-proto) 
- Client libraries (APIs, SDKs, and instrumentation) -- language-specific. For example [Java](https://github.com/open-telemetry/opentelemetry-java) 
- [Collector](https://github.com/open-telemetry/opentelemetry-collector/blob/master/docs/design.md)

The design of OpenTelemetry focuses on backwards compatibility and a clean separation of concerns.

The first component is a stand-alone instrumentation API.  This API is separate from any implementation, and is designed to remain backwards compatible even as new features are added. This stability enables OSS libraries (web frameworks, HTTP clients, databases, etc) to natively instrument themselves without the risk of instrumentation creating a dependency conflict when the libraries are composed together into a larger application.

We believe this is a novel feature, which comes from our OpenTracing roots. Having no easy means to communicate runtime information is a serious pain point in OSS software development today. Native instrumentation allows library authors to focus on the runtime observability of their OSS code, not just testing in development. Library authors will be able to communicate what is important about running their software, providing warnings and tuning advice to their users.

OpenTelemetry is designed to allow experimental signals, such as metrics and logging, to be added without affecting the stability of existing functionality such as tracing, baggage, and resources. OpenTelemetry is also designed to be fully backwards compatible with OpenTracing.

The second component is the OpenTelemetry SDK - our implementation of the API.  The application owner installs and configures the SDK in order to consume API calls and export telemetry. The SDK is also designed with backwards compatibility in mind. Our goal is to ensure there is a clean upgrade path for application owners, allowing them to always stay on the latest version, and consume security patches and new features without concern that their older existing plugins will break, or otherwise need to be pinned to an older version of the SDK.

A detailed description of how our design client creates the above benefits can be found [here](https://github.com/open-telemetry/oteps/blob/main/text/0147-upgrade-procedures.md).

The third component is the OpenTelemetry Collector. The Collector is a stand-alone process that handles data processing, configuration, and protocol translation. Operators run Collectors to manage the type of complex observability pipelines which are becoming common, where telemetry is being sent to multiple, ever changing backends and analysis tools. By managing this workload via the Collector instead of the clients, operators are able to control their observability pipeline without the need to restart or otherwise affect running applications.

OpenTelemetry uses its own protocol, OTLP, to combine all telemetry into a single stream of data. At the edge of the system, a Collector then converts OTLP into the protocol required by a particular backend.

These cleanly separated components compose together into a robust telemetry system, which provides both the stability and flexibility needed to be a long term solution for end users. This removes the need to painfully reinstrument and then redeploy a new set of technology every time a new, useful analysis tool becomes available.

### Features

OpenTelemetry is well-suited to diverse software systems, from client applications, large monoliths, or highly-distributed microservices. It also provides a collector component that can proxy, aggregate, and enrich telemetry before it is exported to a backend.

Supported languages :
- .NET (SDK  and Auto-Instrumentation)
- C/C++ SDK
- GoLang SDK
- Java (SDK and Auto-Instrumentation)
- JavaScript SDK
- Python SDK
- Ruby SDK
- Erlang/Elixir SDK
- PHP SDK
- Rust SDK
- Swift SDK

OpenTelemetry also offers support for popular open-source projects:

- **Traces:** full support for Jaeger and Zipkin. The trace specification is stable (1.0) as is the DotNet and Java instrumentation. Python is in RC and expected to achieve stable status (1.0) in the next month. Other languages to follow.
- **Metrics:** basic support for Prometheus (includes OpenMetrics though translation to OTLP needs to be addressed prior to metrics GA). The Metrics specification is expected to be 1.0 along with client libraries in the 2H2021.
- **Logs:** alpha support for Fluentd (Stanza recently joined project). GA is expected in the first half of next year.

In addition, the platform is extensible so adding more standards is possible. Finally, one of the project's core features is being pluggable and vendor-agnostic – today it supports a variety of receivers and exporters – thus helping ensure widespread adoption. The core OpenTelemetry libraries and SDKs integrate with popular CNCF and open source ecosystem projects, and vendors may choose to publish plugins in separate -contrib repos or produce their own distributions.

## Progress Since Sandbox

### Community

The community statistics are extracted from [CNCF devstats](https://opentelemetry.devstats.cncf.io/).

| Type | Start of Incubation | 2021 |
| ------ | --------------- | --------- |
| Commits| 66 | 20181 |
| Code Committers | 6 | 843 (142 companies) |
| Pull Requests | 17 | 15218 |
| Contributors | 17 | 2,188 (142 companies) |
| Contributions | 201 | 182,741 |
| Forks | 9 | 2281 |
| Watchers| 14 | 6479 |

**Outreach:** A priority of the project is to build a healthy and collaborative community so we can achieve our ambitious technical goals. In addition to frequent SIG meetings (https://github.com/open-telemetry/community#special-interest-groups), the project chose to foster the community in multiple ways such as: 
- [OTel Community Day 2020](https://events.linuxfoundation.org/open-telemetry-community-day) 
- [Community Survey](https://github.com/open-telemetry/community/issues/370#issuecomment-719072644)
- [User Feedback](https://medium.com/opentelemetry/opentelemetry-garnering-user-feedback-for-a-better-product-650b4433e805)
- [OTel Celebrates IWD 2021](https://medium.com/opentelemetry/opentelemetry-observes-international-womens-day-2021-4493a157f119)

### Technical

OpenTelemetry's central technical achievement has been its convergence of two vendor-neutral and numerous proprietary telemetry frameworks into one set of specifications and SDKs. As per the community section, getting dozens of vendors and end-users to agree on a common specification was non-trivial. Given these constraints, we have achieved:
- Full transition support of OpenTracing and OpenCensus in preparation for sunsetting these projects.
- Specification reached Beta in 2020.
- Tracing specification is [GA(v1.0.0)](https://medium.com/opentelemetry/opentelemetry-specification-v1-0-0-tracing-edition-72dd08936978) as of February 2021. 
- Designing the collector to be [extensible](https://medium.com/opentelemetry/extending-the-opentelemetry-collector-with-your-own-components-64c10cf675db)
- Spinning up several [SIGS](https://github.com/open-telemetry/community#special-interest-groups)
- Support of several languages 


## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._**

A list of our adopters can be found [here](https://github.com/open-telemetry/community/blob/master/ADOPTERS.md)

The following are three highlighted examples of adopters: 
- [AWS](https://aws.amazon.com/) AWS Distro for OpenTelemetry - An open source distribution of OpenTelemetry which includes the Collector with key receivers and exporters and Java, JavaScript, .Net and Go API/SDKs. See aws-otel.github.io. AWS also has monitoring services including CloudWatch, Amazon Managed Service for Prometheus and AWS X-Ray that use the OpenTelemetry client/agent for data collection and consumption. See https://aws.amazon.com/otel
- [Shopify](https://www.shopify.com/)  Shopify is an ecommerce platform for hosting online stores. At Kubecon EU 2020, Francis Bogsanyi shared their journey on adopting OpenTelemetry. https://sched.co/ZekF 
- [ecobee](https://www.ecobee.com/)  ecobee is a company that makes smart thermostats, temperature sensors and other home products. It uses OpenTelemetry instrumentation  https://www.honeycomb.io/case-studies/bees-working-together-how-s-engineers-adopted-honeycomb/
- [HelloFresh](https://www.hellofresh.com/) uses OpenTelemetry, including inside of their common healthcheck library. https://github.com/hellofresh/health-go/pull/55

**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

There are 11 project-wide maintainers that form the governance committee. This committee  has a rule of [maximal representation](https://github.com/open-telemetry/community/blob/master/governance-charter.md#maximal-representation) to ensure that no single company is over represented. In addition, there are 7 project-wide technical maintainers that form the [technical committee](https://github.com/open-telemetry/community/blob/main/tech-committee-charter.md).

Each SIG has additional maintainers/approvers. Names and links to those additional names can be found [here](https://github.com/open-telemetry/community/blob/master/community-members.md)

**_Demonstrate a substantial ongoing flow of commits and merged contributions._**

- [Commit Activity](https://opentelemetry.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-2y&to=now&var-period=w&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all)
- [Git clones and Visitors to community page](https://github.com/open-telemetry/community/graphs/traffic)
- [New Contributors over past year](https://opentelemetry.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)
- [CNCF devstats](https://opentelemetry.devstats.cncf.io/)

**_A clear versioning scheme._**

SemVer is followed with maturity levels clearly defined and detailed [here](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/versioning-and-stability.md).

## Statement on Alignment with the CNCF Mission

OpenTelemetry delivers all of the raw materials required for robust observability of modern software applications. Portability, vendor neutrality, and common data formats for collected signals are crucial for cloud-native computing as highlighted in the [CNCF charter](https://www.cncf.io/about/charter/). OpenTracing and OpenCensus are both robust projects, but their overlapping charters have created confusion and wasted work in the larger open-source community. As such, OpenTelemetry also adds considerable value to the CNCF ecosystem by creating the shared destination for the OpenTracing and OpenCensus projects and communities.

## Comparison with Other CNCF Observability Projects

There are two categories of projects that OpenTelemetry can be compared to: (1) projects with their own backend implementations, for which OpenTelemetry aims to provide compatible export, and (2) projects that are API standardization without reference implementations, for which OpenTelemetry aims to support ingesting & exporting.

Jaeger and Prometheus are examples of backend projects; OpenMetrics and W3C Trace Context are examples of a standardization project.

### Comparison to Jaeger
Jaeger is a distributed tracing backend that is fully supported by the OpenTelemetry SDKs which can export data in Jaeger-native format. [Jaeger maintainers stated](https://medium.com/jaegertracing/jaeger-embraces-opentelemetry-collector-90a545cbc24) that once OpenTelemetry SDKs reach maturity and parity with Jaeger’s own client libraries, Jaeger SDKs will be declared end-of-life. Jaeger backend components are also being rebuilt on top of the OpenTelemetry Collector framework, in order to consolidate the efforts on developing new features like tail based sampling and to natively understand and consume OpenTelemetry’s data model (OTLP).

### Comparison to Prometheus 
Prometheus has a couple components in their ecosystem: client libraries and prometheus server.

Prometheus server is a metrics backend that is going to be fully supported by the OpenTelemetry SDKs which will be able to export data in Prometheus/OpenMetrics format.

There is overlap between OpenTemeletry client libraries and Prometheus client libraries. The main difference is that OpenTelemetry client libraries are backend agnostic and support advanced features like request correlation (via baggages), allowing users to customize aggregation and select the backend of their choice. The OpenTelemetry instrumentation API is explicitly designed to be embedded in shared OSS libraries without risk of compatibility issues.

## Future Plans

There are roadmaps for the project as a whole as well as for each SIG, available as GitHub project boards and `roadmap.md` files. The only community issues would be those tracked in GitHub issues.

Examples include:
- https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/roadmap.md
- https://github.com/open-telemetry/opentelemetry-java-instrumentation/blob/main/docs/ga-requirements.md
- https://github.com/orgs/open-telemetry/projects/5 

Per signal, the roadmaps are:
- Tracing: 1.0 specification, 1.0 DotNet, 1.0 Java, RC1 Python today; goal is 1.0 stability for all beta instrumentation libraries in next month. Collector tracing stability by mid-year.
- Metrics: specification is expected to reach 1.0 around mid-year.
  - Prometheus & OpenMetrics data model & wire format compatibility will be reached by middle of 2021
  - Full semantic, etc Prometheus compatibility including staleness handling & up handling will be reached by end of 2021 will include full Prometheus & OpenMetrics compatibility. GA for instrumentation libraries is planned in 2H2021 for the languages with tracing GA support.
  - For more information, please refer to this published [roadmap](https://medium.com/opentelemetry/opentelemetry-metrics-roadmap-f4276fd070cf)
- Logs are currently not a focus area, and not expected to be finished within 2021
  - For more information, please refer to the Logging section within this [document](https://docs.google.com/document/d/1vV3teqg9vmpaxgMS8JiQDghUcnpd3Qunq_tJu9sQaRE/edit?usp=sharing)
