## Background
OpenTelemetry is an integrated set of APIs and libraries that generate, collect, and describe telemetry about distributed systems. This data includes basic context propagation, distributed traces, metrics, and other signals in the future. OpenTelemetry makes it easy to get critical telemetry data out of your services: for each language it offers a single set of APIs, libraries, and data specifications, and developers can make use of which components they see fit.

OpenTelemetry is well-suited to diverse software systems, from client applications, large monoliths, or highly-distributed microservices. It also provides a collector component that can proxy, aggregate, and enrich telemetry before it is exported to a backend.

Supported languages :
- .NET (SDK  and Auto-Instrumentation)
- C/C++ SDK
- GoLang SDK
- Java (SDK and Auto-Instrumentation)
- JavaScript SDK
- Python SDK
- Ruby SDK
- Erland/Elixir SDK
- PHP SDK
- Rust SDK
- Swift SDK

## Sandbox Presentation 
- [Slides](https://docs.google.com/presentation/d/1roUFtkiVw25N7nRuUnMa-ziNhb6vzhZAe43jpBfiDiw/edit#slide=id.g5724df7a45_0_258)
- [Proposal](https://github.com/cncf/toc/blob/f4ecf1bf391867d0840434b638cd9ef6cb2f57fe/proposals/sandbox/opentelemetry.adoc)
- [PR](https://github.com/cncf/toc/pull/233)


## Alignment with CloudNative
OpenTelemetry delivers all of the raw materials required for robust observability of modern software applications. Portability, vendor neutrality, and common data formats for collected signals are crucial for cloud-native computing as highlighted in the [CNCF charter](https://www.cncf.io/about/charter/). OpenTracing and OpenCensus are both robust projects, but their overlapping charters have created confusion and wasted work in the larger open-source community. As such, OpenTelemetry also adds considerable value to the CNCF ecosystem by creating the shared destination for the OpenTracing and OpenCensus projects and communities.


## Highlights from the past year

OpenTelemetry declared Beta for the collector and five languages(Java, JavaScript, Python, Go, and .Net APIs and SDKs). The Beta milestone had several requirements such as
Implementation of metrics, traces, context propagation and resource metadata for version <= 0.3 in the [spec](add link to repository)
Documentation for getting started
Ability to export in Jaeger, Prometheus and Zipkin
Testing coverage and automate release process
For more information on the Beta requirements and the process, please refer to the [Beta Launch Plan](https://docs.google.com/document/d/1eviggoIguOS89dgKL-8ntUOCfPP6FLWAoQkGoBDsld0/edit?usp=sharing)


Adoption! There are 10 public adopters including vendors, end users and a cloud provider. It is a wide set of use cases that highlights the flexibility of the project.

Jaeger embraces OpenTelemetry with contributions and integrations. [Blog posts with details](https://medium.com/jaegertracing/jaeger-embraces-opentelemetry-collector-90a545cbc24)


## Dev Stats
*Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.* 

[OpenTelemetry devstats](https://opentelemetry.devstats.cncf.io/) 

Notable Highlights:
Activity per repository over 6 months shows a consistent growth https://opentelemetry.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&from=now-6M&to=now

There are 75 companies committing to OpenTelemetry https://opentelemetry.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=commits

New PRs per week have gone from  ~80 / week  in July of 2019 to ~225-250/week in July of 2020 https://opentelemetry.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-2y&to=now

## Maintainers
*How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)*

There are 9 project-wide maintainers on the governance committee
- Ben Sigelman (Lightstep)
- Bogdan Drutu (Splunk)
- Constance Caramanolis (Splunk)
- Liz Fong-Jones (Honeycomb)
- Morgan James McLean (Google)
- Sarah Novotny (Microsoft)
- Sergey Kanzhelev (Google)
- Ted Young(Lightstep)
- Yuri Shkuro (Uber)

The governance committee has a rule of [maximal representation](https://github.com/open-telemetry/community/blob/master/governance-charter.md#maximal-representation) to ensure that no single company is over represented. 

Each SIG has additional maintainers/approvers. Names and links to those additional names can be found [here](https://github.com/open-telemetry/community/blob/master/community-members.md)


## Adoption
*What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)*

A list of our adopters can be found [here](https://github.com/open-telemetry/community/blob/master/ADOPTERS.md)


## Comparison against last year
*How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)*

OpenTelemetry joined the CNCF as a sandbox project in May of last year. There is no review to compare against. Please refer to highlights section for notable accomplishments since joining CNCF. 

## Goals
*What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?*

The following are “must-do” goals:
- Release GA(1.0). Similar to Beta, GA has requirements around implementation, documentation and testing. In addition, the guarantee is made that the APIs and SDKs must not take any breaking changes between release candidates, the GA releases, and any subsequent releases that implement spec version < 2.0. [GA Launch Plan Documentation](https://docs.google.com/document/d/1dyXxsmpcHLvlpyvZc3HXBfKF-vrfJlMLAN6CQVzCEos/edit?usp=sharing)
- Archive OpenTracing and OpenCensus. 
- Meet incubation criteria. 

Project roadmap can be found [here](https://opentelemetry.io/project-status/)



The following are goals that are beneficial to the project and community, but are just “nice to have” / stretch goals for the next twelve months:
- Reference implementation of log support. OpenTelemetry focused on providing support for metrics and traces. The desire for log support increased as the project gained momentum. [Ex. Log working group https://github.com/open-telemetry/community#logs-working-group ] Adding support for logs makes opentelemetry a more compelling option for instrumentation. 
- Integration with a CNCF project.

## Collaboration with CNCF
*How can the CNCF help you achieve your upcoming goals?*

There are two primary ways that CNCF can help with the growth of the project:
- Continued support at CNCF events. This 
- Promoting integration with other CNCF Projects. The CNCF TOC can facilitate introductions with other projects that could be interested in adopting OpenTelemetry. 

## Incubation
*Do you think that your project meets the criteria for incubation?*

We believe that the project is close to meeting incubation criteria.
There are two areas that we want to expand upon more before applying: 
- Documenting end user Adoption - On the list provided above, there are several strong end users, but there is no public documentation(blogs, tech talks) of what it looks like. 
- Versioning - API spec versioning is specified [here](https://github.com/open-telemetry/opentelemetry-specification#versioning) and labeling [here](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/library-guidelines.md#version-labeling ). We need to document the policy around API breaking changes.
