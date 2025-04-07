# Thanos Incubation Stage Proposal
 
During Thanos presentation at the [CNCF TOC meeting on 7/9/2019](https://docs.google.com/presentation/d/1jhzJlSAAJNNil1nIYp60eSMH3LPd6AwqHLt3vEAzMSg/edit#slide=id.g5cdf155fc3_0_4)
we were suggested by TOC members to apply straight for [the Incubation Stage](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage).
We all agreed back then that we will join Sandbox first and after that perform Due Diligence required for the Incubation Stage. 
 
This proposal aims to address the Diligence requirements.
 
As the Thanos team, we believe we are ready to be considered joining Incubation Stage.

## Incubation State Requirements

1. *Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.*

Thanos has a large number of successful adopters that are running Thanos on production. 
Some of those who choose to speak publicly about their usage, we can see in the list of public adopters on our website [https://thanos.io/](https://thanos.io/).
Some of them decided to go beyond that and shared publicly about their successes with Thanos:
 
* [Improbable](https://improbable.io) [started Thanos](https://improbable.io/blog/thanos-prometheus-at-scale)
and were using the project on the production from the earliest days. [This blog post](https://improbable.io/blog/thanos-architecture-at-improbable)
describes the usage in detail here.
* [Monzo](https://monzo.com/) is UK’s banking startup that was one of the earliest adopters and contributors to Thanos. 
They described their success in detail [in this blog post](https://monzo.com/blog/2018/07/27/how-we-monitor-monzo).
* [uSwitch](https://www.uswitch.com/) shared their experience [on Medium](https://medium.com/uswitch-labs/making-prometheus-more-awesome-with-thanos-fbec8c6c28ad).
* [Alibaba Cloud](https://us.alibabacloud.com/) presented their usage of Thanos during CNCF KubeCon China 2019. Video is available [here](https://www.youtube.com/watch?v=ZS6zMksfipc).
* [Red Hat](https://www.redhat.com/) is using Thanos in production. One of the usages was described [here](https://blog.openshift.com/federated-prometheus-with-thanos-receive/). 
* [HelloFresh](https://engineering.hellofresh.com/) and their blog post about the Thanos usage is available [here](https://engineering.hellofresh.com/monitoring-at-hellofresh-part-1-architecture-677b4bd6b728). 
* [Taboola](https://engineering.taboola.com/) shared their story too [here](https://engineering.taboola.com/monitoring-and-metering-scale/).
* [Banzai Cloud](https://banzaicloud.com/) recently shared their production usage of Thanos in a multi-cluster environment in [this blog post](https://banzaicloud.com/blog/multi-cluster-monitoring/).
They also maintain a [Thanos helm chart](https://github.com/banzaicloud/banzai-charts/tree/master/thanos).
 
2. *Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.*
 
We have 7 maintainers, 4 triage members, overall from 8 different companies. The details are here: https://thanos.io/maintainers/
 
3. *Demonstrate a substantial ongoing flow of commits and merged contributions.*
 
Commits per week over the last 6 months: **28**

![](https://github.com/bwplotka/toc/blob/pics/proposals/thanos-commits.png?raw=true)

Issue Opened/Closed per week over the last 6 months: **22.29**
 
![](https://github.com/bwplotka/toc/blob/pics/proposals/thanos-openedclosedissues.png?raw=true)
 
New PRs per week over the last 6 months: **18**

![](https://github.com/bwplotka/toc/blob/pics/proposals/thanos-prs.png?raw=true)

All data was gathered thanks to [CNCF devstats dashboards](https://thanos.devstats.cncf.io/d/8/dashboards?orgId=1&from=now-6M&to=now-1h)
 
4. *A clear versioning scheme.*
 
We have regular releases every 6 weeks with dedicated Release Shepherd for each. All is documented [here](https://thanos.io/release-process.md/).
 
Release cadence and the process is very similar to the [Graduated CNCF Prometheus project](https://github.com/prometheus/prometheus/blob/master/RELEASE.md) in order to reuse the best practices.
Additionally, it is convenient and familiar for Prometheus Ecosystem Community.

## What is Thanos: Recap
 
Thanos is a metric system that aims for being a simple and cheap option to centralize and scale-out the Prometheus based systems.
Those goals did not change since the [Sandbox proposal](https://github.com/cncf/toc/pull/256) and Thanos still achieves all of this by:

* Gradual and incremental installation model. Users can seamlessly transform vanilla Prometheus system into partial Thanos, optionally upgrade into fully-fledged Thanos deployment with near-unlimited retention capabilities in separate isolated steps.
* Single, widely available, cheap and optional dependency: object storage.
* Built-in Prometheus HA support and downsampling capabilities.
* Flexible deployment model due to a cloud-native, microservice approach using gRPC.

Concretely the aims of the project are:

* Global query view of metrics.
* Unlimited retention of metrics, including downsampling.
* High availability of components, including support for Prometheus HA.

## Statement on alignment with the CNCF mission:

The Thanos project aligns well with the mission of the CNCF by empowering organizations to build, run and monitor applications at scale.

It can be used to monitor and alert on microservices and infrastructure that spans over not only multiple clusters but over multiple clouds with a simple approach to adoption and scaling.

Thanos aims to enable people to take their current Prometheus ecosystem and scale it out beyond their current capabilities, gradually and without downtime.

Building on top of the work that is being done by the Prometheus community we can scale out to give Thanos users a global view of all their metrics without the need for aggregations.

## Comparison to Cortex

Cortex is another CNCF project that provides high-availability and long-term storage to Prometheus. Both Thanos and Cortex make different trade-offs that will appeal to different use-cases. Thanos approached the problem with the focus on simplicity, cost efficiency, and gradual installation. It takes advantage of having only a single dependency (object storage), single stable API (StoreAPI) and is integrating with Prometheus via both pull and push-based methods. 

On the other hand, Cortex is more mature and has more performance optimizations focused to reduce the latency for complex metric queries. It has also better multitenancy support.

It’s worth to note that Cortex and Thanos’s communities are **constantly collaborating with each other**, the recent ones
being using the cortex-frontend to provide query caching for Thanos and exploration of writing blocks and using the Thanos query path in Cortex.
Example: [Thanos PR](https://github.com/thanos-io/thanos/pull/1881), [Cortex PR](https://github.com/cortexproject/cortex/pull/1942). There are plans to collaborate further and the community and usage of both the projects are only growing!

## Project Progress since Sandbox

* Improved documentation accordingly to the [CII Best Practices](https://github.com/thanos-io/thanos/issues/1389).
* We agreed on our [Thanos Governance](https://thanos.io/governance.md/). We are also one of the first projects which
applied *the rule of maximum 2 votes from a single company* to ensure the project independence.
* All key metrics around contributions, downloads and committers have seen significantly grown.
* Thanos being officially incorporated by [several new customers](https://github.com/thanos-io/thanos/commits/master/website/static/logos). Unofficially by even more based on spoken feedback.
* Three more [maintainers](https://github.com/thanos-io/thanos/commits/master/MAINTAINERS.md) joined.
* The project has seen an increase in integrations e.g: Memcached support, new object storage implementations, tracing backends.
* Major improvements in stability and resource consumption were done. Some of them were done with collaboration with Prometheus and Cortex projects,
**which proves why open source is amazing!**
* Thanos participates in both the Google Summer of Code and Community Bridge 

## Community Bridge

We applied to the Community Bridge program with various tasks for Thanos project. Details can be found [here](https://people.communitybridge.org/project/f51284ab-f652-47b1-9819-cd4135e75c00).

We are currently mentoring one developer from India within this program and she has already contributed significantly to the Thanos project.

## Google Summer of Code

We plan to apply for GSoC as well, we will submit the task there as well before deadline.

## Features & Roadmap

Following is a list of features we plan to add in next iterations with the help of the community.
The list is also available in updated form via [GitHub milestones](https://github.com/thanos-io/thanos/projects).

* [Move tracing to OpenTelemetry.](https://github.com/thanos-io/thanos/issues/1972)
  * Purpose: Reuse and stability.
* [Improve loading blocks with binary format and mmap.](https://thanos.io/proposals/201912_thanos_binary_index_header.md/)
  * Purpose: Lower memory footprint of Store Gateway, faster startup
* [Contribute to Cortex Query-frontend to make response cache work better against Thanos.](https://github.com/thanos-io/thanos/issues/1651)
  * Purpose: Consistency and resilience of the caching layer, downsampling support.
* [LabelNames and LabelValues time range limit.](https://github.com/thanos-io/thanos/issues/1811)
  * Purpose: Lower memory usage of Store Gateway, lower latency of LabelValues/LabelNames requests
* [Request logging](https://github.com/thanos-io/thanos/issues/1706) 
  * Purpose: Query audit, debuggability for slow queries, matching trace to request
* [Native gRPC certificate rotation & mTLS everything.](https://github.com/thanos-io/thanos/pull/1672)
  * Purpose: Security, Simpler on-prem installation
* [Adhoc e2e Benchmark for Thanos](https://github.com/thanos-io/thanos/issues/1707)
  * Purpose: Confidence in optimizations we do in the upstream
* [Race-free compactions](https://thanos.io/proposals/201901-read-write-operations-bucket.md/)
  * Purpose: Consistency and data correctness during various operations, support for eventual consistency buckets.
  * **Proposed for community bridge and the first task**
* [Offline deduplication](https://github.com/thanos-io/thanos/issues/1014)
  * Purpose: Saves resources per request on Store and Querier, 
* [Query limits against OOMs e.g series limits](https://github.com/thanos-io/thanos/issues/448)
  * Purpose: Reliability
* [Compaction bucket web UI](https://github.com/thanos-io/thanos/issues/1657)
  * Purpose: Debuggability
* [Deletions](https://github.com/thanos-io/thanos/issues/1598)
  * Purpose: GDPR, removing broken series, retention per series
* [Get Receive component out of experimental stage](https://github.com/thanos-io/thanos/issues/1093)
  * Purpose: Reliability, Multitenancy, Usability for community
