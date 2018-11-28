# Jaeger Graduation Application

Jaeger was open sourced in April, 2017 and joined the CNCF in September, 2017. The project has grown significantly over time. As of Oct-30-2018 ([devstats][devstats]):

  * 6185 stars on GitHub
  * 878 forks
  * 224 authors of commits and pull requests
  * 350 issue creators  

Jaeger fulfills all CNCF incubating and graduation criteria. The following application links to the required information to become a graduated project.

## Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

Known end users include: Alibaba Cloud, Base CRM (acquired by ZenDesk), Candide, Circonus, ContaAzul, FarmersEdge, GrafanaLabs, Massachusetts Open Cloud, Northwestern Mutual, Nets, Stagemonitor, RiksTV, SeatGeek, SpotHero, Vistar Media, Uber, Weave HQ, Weaveworks, Zenly.

Some specific examples:

  * [Uber Technologies](https://eng.uber.com/distributed-tracing/) has been running Jaeger in production since 2015. Over 2000 microservices are integrated with Jaeger, which collects about 5 billion spans per day. Data mining techniques on traces are used to significantly reduce root cause analysis time during outages.
  * [Weaveworks](https://www.weave.works/) empowers developers and DevOps teams to build better software faster. Using Jaeger in production since Spring 2018. ~70 microservices, ~600 containers, most of services are instrumented with Jaeger. Using AWS Elasticsearch as the span storage backend, with two weeks retention. 
    * Bryan Boreham at KubeCon EU 2018. [How we used Jaeger and Prometheus to deliver lightning-fast user queries](https://www.youtube.com/watch?v=qg0ENOdP1Lo).
  * [ContaAzul](http://contaazul.com), ERP for small and medium business in Brazil, one legacy monolith with ~60 microservices around it, 20 of them already instrumented and the rest being prioritized. Microservices are deployed to a self-maintained Kubernetes cluster on AWS. Each Kubernetes node holds a Jaeger agent responsible for collecting spans from its node pods. Span storage is on AWS Elasticsearch Service; it currently has 2B searchable documents related to Jaeger spans.
    * Quote: "Jaeger is helping us to increase observability since Feb 2018 and is praised by our developers."
  * [Weave HQ](https://www.getweave.com/) threads together data, software and communication platforms to build stronger relationships at the point of contact. Deployed at over 8000 medical offices. Using Jaeger since early 2018; ~150 active microservices, ~500 containers, ~90% of microservices are instrumented for tracing. Using Cassandra as span storage backend, 1 week retention, ~2M spans/day.
    * Fun fact 1: started collecting fewer metrics in Prometheus because Jaeger works much better for root cause analysis than having extra metrics.
    * Fun fact 2: most application logs are written to Jaeger spans via the OpenTracing API and viewed in Jaeger UI in the context of the trace, instead of log aggregation tools like ELK.
  * [Candide](https://candide.eu/), a mobile app for gardeners, uses Jaeger in production running on GKE with an Elasticsearch backend hosted on Elastic Cloud. 30+ microservices.

Jaeger is also being integrated with other open source projects:

  * OpenCensus libraries and agent ship with [exporters for Jaeger](https://opencensus.io/guides/exporters/supported-exporters/java/jaeger/)
  * [Istio comes with Jaeger included](https://istio.io/docs/tasks/telemetry/distributed-tracing/)
  * [Envoy works with Jaeger native client](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/jaeger_native_tracing)
  * Eclipse Trace Compass incubator supports [importing traces from Jaeger](https://github.com/tuxology/tracevizlab/tree/master/labs/303-jaeger-opentracing-traces)

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

# Demonstrate a substantial ongoing flow of commits and merged contributions.

https://jaeger.devstats.cncf.io/d/4/companies-stats?orgId=1&var-period=m&var-metric=activity&var-repogroup_name=All&var-companies=All 

# Have committers from at least two organizations.

We currently have 7 committers from 2 organizations, Uber and Red Hat.

# Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

https://bestpractices.coreinfrastructure.org/en/projects/1273 

# Adopt the CNCF Code of Conduct.

https://github.com/jaegertracing/jaeger/blob/master/CODE_OF_CONDUCT.md

# Explicitly define a project governance and committer process. 

  * Governance: https://github.com/jaegertracing/jaeger/blob/master/GOVERNANCE.md
  * Committers: https://github.com/jaegertracing/jaeger/blob/master/CODEOWNERS 

# Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

https://github.com/jaegertracing/jaeger/blob/master/ADOPTERS.md

[devstats]: https://jaeger.devstats.cncf.io/d/18/project-statistics?orgId=1&var-period_name=Last%20decade&var-repogroup_name=All
[committers]: https://github.com/jaegertracing/jaeger/blob/master/CODEOWNERS
