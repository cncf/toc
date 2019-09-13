# Jaeger Graduation Application

Jaeger was open sourced in April, 2017 and joined the CNCF in September, 2017. The project has grown significantly over time. As of Sep-04-2019 ([devstats][devstats]):

  * 8883 stars on the main repository https://github.com/jaegertracing/jaeger
  * 1573 forks
  * 375 authors of commits and pull requests
  * 654 issue creators
  * 1238 contributors

Jaeger fulfills all CNCF incubating and graduation criteria. The following application links to the required information to become a graduated project.

## Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

Known end users include: Alibaba Cloud, Base CRM (acquired by ZenDesk), Candide, Circonus, ContaAzul, FarmersEdge, GrafanaLabs, Massachusetts Open Cloud, Northwestern Mutual, Nets, Stagemonitor, RiksTV, SeatGeek, SpotHero, Ticketmaster, Vistar Media, Uber, Weave HQ, Weaveworks, Zenly.

Some specific examples:

  * [Uber Technologies](https://eng.uber.com/distributed-tracing/)
    * Running Jaeger in production since 2015.
    * Over 2700 microservices are integrated with Jaeger, which collects about 10 billion spans per day, across multiple data centers.
    * Data mining techniques on traces are used to significantly reduce root cause analysis time during outages.
    * Tracing data is utilized by a number of other internal tools, such as using it to reduce duplicate alerts, and to infer data lineage.
  * [Weaveworks](https://www.weave.works/) empowers developers and DevOps teams to build better software faster. Using Jaeger in production since Spring 2018. ~70 microservices, ~600 containers, most of services are instrumented with Jaeger. Using AWS Elasticsearch as the span storage backend, with two weeks retention.
    * Bryan Boreham at KubeCon EU 2018. [How we used Jaeger and Prometheus to deliver lightning-fast user queries](https://www.youtube.com/watch?v=qg0ENOdP1Lo).
  * [ContaAzul](http://contaazul.com), ERP for small and medium business in Brazil.
    * One legacy monolith with ~60 microservices around it, 20 of them already instrumented and the rest being prioritized. Microservices are deployed to a self-maintained Kubernetes cluster on AWS. Each Kubernetes node holds a Jaeger agent responsible for collecting spans from its node pods. Span storage is on AWS Elasticsearch Service; it currently has 2B searchable documents related to Jaeger spans.
    * Quote: "Jaeger is helping us to increase observability since Feb 2018 and is praised by our developers."
  * [GrafanaLabs](https://grafana.com/) runs a managed metrics platform GrafanaCloud.
    * Uses Jaeger in production for observability and query optimization
    * Case study: [Grafana Labs Teams Use Jaeger to Improve Query Performance Up to 10x][grafanalabs].
  * [Weave HQ](https://www.getweave.com/) threads together data, software and communication platforms to build stronger relationships at the point of contact. Deployed at over 8000 medical offices. Using Jaeger since early 2018; ~150 active microservices, ~500 containers, ~90% of microservices are instrumented for tracing. Using Cassandra as span storage backend, 1 week retention, ~2M spans/day.
    * Fun fact 1: started collecting fewer metrics in Prometheus because Jaeger works much better for root cause analysis than having extra metrics.
    * Fun fact 2: most application logs are written to Jaeger spans via the OpenTracing API and viewed in Jaeger UI in the context of the trace, instead of log aggregation tools like ELK.
  * [Candide](https://candide.eu/), a mobile app for gardeners, uses Jaeger in production running on GKE with an Elasticsearch backend hosted on Elastic Cloud. 30+ microservices.
  * [Ticketmaster](https://www.ticketmaster.com/), a ticket sales and distribution company based in Beverly Hills, California, with operations around the world.
    * Traces 100 million transactions per day with Jaeger .
    * Has 50-plus services instrumented and nine groups of engineering teams actively using Jaeger.
    * Case study: [Ticketmaster Traces 100 Million Transactions per Day with Jaeger][ticketmaster].
    * Quote: "Jaeger Tracing is helping us achieve our vision of observability capabilities across different versions of our platforms, how they integrate with each other and how they have grown over time." -- Kraig Amador, Senior Director.
  * [Red Hat](https://www.redhat.com/en), world’s leading provider of enterprise open source solutions, using a community-powered approach to deliver high-performing Linux, cloud, container, and Kubernetes technologies.
    * Jaeger is fully supported by Red Hat as part of its OpenShift Service Mesh product ([press release][openshift]).
  * Many other companies self-reporting on https://stackshare.io/jaeger/in-stacks/ajax

Jaeger is also being integrated with other open source projects:

  * OpenCensus libraries and agent ship with [exporters for Jaeger](https://opencensus.io/guides/exporters/supported-exporters/java/jaeger/)
  * OpenTelemetry project is building support for Jaeger as direct exporters (e.g. in [Java SDK](https://github.com/open-telemetry/opentelemetry-java/pull/481)) as well as from the [OpenTelemetry Service](https://github.com/open-telemetry/opentelemetry-service/blob/master/exporter/README.md#jaeger). See below for more discussion about [Jaeger and OpenTelemetry](#opentelemetry).
  * [Istio comes with Jaeger included](https://istio.io/docs/tasks/telemetry/distributed-tracing/)
  * [Envoy works with Jaeger native client](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/jaeger_native_tracing)
  * Eclipse Trace Compass incubator supports [importing traces from Jaeger](https://github.com/tuxology/tracevizlab/tree/master/labs/303-jaeger-opentracing-traces)

## Project long term sustainability

  * Uber has a team of 7 full time engineers working on Jaeger, both internally and in the open source.
  * Red Hat supports Jaeger as part of its [OpenShift Service Mesh product][openshift].

## Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

Jaeger project contains a number of sub-projects, such as client libraries, integrations with Kubernetes, etc. We do not grant committer rights equally across projects, only as needed. For example:

  * Jaeger backend: Currently, we have 7 [committers][committers] and a few people contributing regularly. Full committers have commit rights across all other repositories.
    * @black-adder (Uber)
    * @jpkrohling (Red Hat)
    * @objectiser (Red Hat)
    * @pavolloffay (Red Hat)
    * @tiffon (Uber)
    * @vprithvi (Uber)
    * @yurishkuro (Uber)
  * Jaeger UI
    * @everett980 (Uber)
    * @tiffon (Uber)
    * @saminzadeh (Uber)
  * Jaeger client for C++
    * @isaachier (Uber)
  * Jaeger client for C#
    * @Falco20019
    * @grounded042 (Chatham Financial)
    * @kevindaviscfc (Chatham Financial)
    * @PhillipChaffee (Chatham Financial)
    * @cwe1ss (c3-ls.com)

We are on the constant lookout for new maintainers to join our ranks. In the past year over 200 people created commits and pull requests to Jaeger code base.

## Demonstrate a substantial ongoing flow of commits and merged contributions.

https://jaeger.devstats.cncf.io/d/4/companies-stats?orgId=1&var-period=m&var-metric=activity&var-repogroup_name=All&var-companies=All

## Have committers from at least two organizations.

We currently have 7 committers from 2 organizations, Uber and Red Hat, in the main backend repository, and more diversity in the client libraries.

## Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

https://bestpractices.coreinfrastructure.org/en/projects/1273

## Have completed an independent and third party security audit

Audit by Cure-53 has been [completed on 2019-05-04](https://github.com/jaegertracing/security-audits). No actual security threats have been identified by the penetration testing. The auditors made a recommendation to add more in-depth security mechanisms to Jaeger instead of relying on the correct deployment and perimeter security. Many of the channels between Jaeger components already support secure communications (https://github.com/jaegertracing/jaeger/issues/1718).

## Adopt the CNCF Code of Conduct.

https://github.com/jaegertracing/jaeger/blob/master/CODE_OF_CONDUCT.md

## Explicitly define a project governance and committer process.

  * Governance: https://github.com/jaegertracing/jaeger/blob/master/GOVERNANCE.md
  * Committers: https://github.com/jaegertracing/jaeger/blob/master/CODEOWNERS

## Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

https://github.com/jaegertracing/jaeger/blob/master/ADOPTERS.md

## Open-source alternatives to Jaeger

### Zipkin

  * https://zipkin.io
  * (Was incubating at Apache Software Foundation, but withdrew in the summer 2019).
  * Offers similar functionality to Jaeger, but lacks certain features, such as adaptive sampling or trace comparisons in the UI.
  * Implemented in Java. Comes with its own instrumentation SDK for Java called Brave.

### Apache Sky-Walking

  * https://skywalking.apache.org/
  * Top level project at Apache Software Foundation.
  * An APM solution that combines tracing, metrics, and alerting.

### Haystack

  * https://expediadotcom.github.io/haystack/
  * Tracing and analysis system from Expedia.
  * Beyond simple tracing, allows to observe trends and setup anomaly detection.

### Pinpoint

  * https://naver.github.io/pinpoint/
  * APM system developed by NaverCorp (S. Korea).
  * Focused mostly on Java applications performance, for which it provides zero-touch instrumentation integrated with many OSS frameworks.

## Jaeger relationship to OpenTracing and OpenCensus

### OpenTracing

OpenTracing is a vendor-neutral set of instrumentation APIs for distributed tracing. Jaeger client libraries are OpenTracing-compliant and can be used with any OpenTracing-based instrumentation in multiple languages.

### OpenCensus

OpenCensus is a different set of APIs for metrics and tracing. It comes with a default implementation included, which in some cases is not separable from the APIs. It is not compatible with OpenTracing instrumentation. OpenCensus libraries include data exporters that can send trace data directly to Jaeger.

### OpenCensus Agent and Collector

OpenCensus Agent and Collector are backend components that can receive data from OpenCensus instrumentation libraries (as well as sources, including Jaeger libraries) and forward the data to various tracing backends, including Jaeger, Zipkin, and Stackdriver (other backends are also supported through extensions). Functionally these two components are very similar to Jaeger Agent and Collector. However, the OpenCensus components are explicitly scoped to data collection and forwarding, they do not include the actual trace backends with storage, UI, and analytical capabilities.

### OpenTelemetry

The OpenTracing and OpenCensus projects have joined forces in a new project [OpenTelemetry](https://opentelemetry.io) that will provide APIs, default implementation, and reusable instrumentation for metrics, tracing, and other observability signals. The OpenTelemetry project also subsumed the OpenCensus Agent and Collector components (https://github.com/open-telemetry/opentelemetry-service).

This is all good news for Jaeger. When the OpenTelemetry libraries reach parity with the existing Jaeger client libraries, the Jaeger project is planing to sunset its own client libraries and recommend that users use the OpenTelemetry instead. Similarly, it is possible that Jaeger project may stop developing Jaeger Agent and Collector and use the respective components from the OpenTelemetry Service, or extend them with Jaeger-specific functionality, such as adaptive sampling (since projects are implemented in Go). This will allow the Jaeger project to focus its efforts on the development of features in the Jaeger backend, including novel visualizations and data mining techniques, and leave the data gathering problem to OpenTelemetry.

It is also likely that OpenTelemetry and Jaeger will converge on a common data model for traces.

The full transition and sunsetting of Jaeger components replicated by OpenTelemetry can take 1-2 years. During the transition the users may continue using the existing Jaeger libraries or start incrementally switching to OpenTelemetry components, as backwards compatibility will be maintained.

For more information, see [Jaeger and OpenTelemetry blog post](https://medium.com/jaegertracing/jaeger-and-opentelemetry-1846f701d9f2).



[devstats]: https://jaeger.devstats.cncf.io/d/18/project-statistics?orgId=1&var-period_name=Last%20decade&var-repogroup_name=All
[committers]: https://github.com/jaegertracing/jaeger/blob/master/CODEOWNERS
[openshift]: https://www.redhat.com/en/about/press-releases/red-hat-launches-openshift-service-mesh-accelerate-adoption-microservices-and-cloud-native-applications
[ticketmaster]: https://medium.com/jaegertracing/ticketmaster-traces-100-million-transactions-per-day-with-jaeger-38ec6cf599f0
[grafanalabs]: https://medium.com/jaegertracing/grafana-labs-teams-observed-query-performance-improvements-up-to-10x-with-jaeger-cec84b0e3609
