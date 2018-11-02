# Jaeger Graduation Application

Jaeger was open sourced in April, 2017 and joined the CNCF in September, 2017. The project has grown significantly over time. As of Oct-30-2018 ([devstats][devstats]):

  * 6185 stars on GitHub
  * 878 forks
  * 224 authors of commits and pull requests
  * 350 issue creators  

Jaeger fulfills all CNCF incubating and graduation criteria. The following application links to the required information to become a graduated project.

## Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

Known end users include: Alibaba Cloud, Base CRM (acquired by ZenDesk), Circonus, ContaAzul, FarmersEdge, GrafanaLabs, Massachusetts Open Cloud, Northwestern Mutual, Nets, Stagemonitor, RiksTV, SeatGeek, SpotHero, Vistar Media, Uber, Weave, Weaveworks, Zenly.

Jaeger is also being integrated with other open source projects:

  * OpenCensus libraries and agent ship with [exporters for Jaeger](https://opencensus.io/guides/exporters/supported-exporters/java/jaeger/)
  * [Istio comes with Jaeger included](https://istio.io/docs/tasks/telemetry/distributed-tracing/)
  * [Envoy works with Jaeger native client](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/jaeger_native_tracing)

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
