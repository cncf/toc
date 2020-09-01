# Cortex proposal for CNCF Incubation

## Background
[Request for Incubation](https://github.com/cncf/toc/pull/125)

[Project Review Proposal](https://github.com/cncf/toc/pull/125)

[Original TOC presentation](https://docs.google.com/presentation/d/190oIFgujktVYxWZLhLYN4q8p9dtQYoe4sxHgn4deBSI/edit#slide=id.g380c8a0114_0_178)

Cortex is a horizontally scalable, highly available, multi-tenant, Prometheus API compatible service that offers a long-term storage solution.

For teams looking for a Prometheus solution that offers the following over vanilla Prometheus:

* Long-term metrics storage in a variety of cloud based and on-prem NoSQL data stores


* Tenancy model supporting commercial SaaS offerings or large/multiple Kubernetes installations requiring data separation


* On-demand Prometheus instance provisioning


* A highly-available architecture that benefits from cloud-native architectures run with Kubernetes


* A highly scalable Prometheus experience that horizontally scales


* The ability to handle large metric topologies in a single instance without the need for federation


Cortex was presented at the [CNCF TOC meeting on 6/5/2018](https://docs.google.com/presentation/d/190oIFgujktVYxWZLhLYN4q8p9dtQYoe4sxHgn4deBSI/edit#slide=id.g25ca91f87f_0_0). We’ve grown a lot since then and the project is a lot more active and mature now.

Notable improvements in the last year include:
* An easy to use single process version for people to test things out.
* Queries are now much faster (up to 10x).
* We now use a lot less disk space.
* Cortex is now much more stable and easier to run, with more improvements on their way.
* Our alerting and recording rule layer is now horizontally scalable.

Further, a lot of the work in Cortex also involved improvements in upstream Prometheus and small subset include:
* https://github.com/prometheus/prometheus/pull/4588
* https://github.com/prometheus/prometheus/pull/5707
* https://github.com/prometheus/tsdb/pull/642
* https://github.com/prometheus/prometheus/pull/5316
* https://github.com/prometheus/prometheus/pull/5131

## Alignment with Cloud Native

Cortex fully supports the CNCF’s goal for scalability, "Ability to support all scales of deployment, from small developer centric environments to the scale of enterprises and service providers." 

There are many different ways to provide a scalable and available metric system for Kubernetes. Cortex with it’s tenancy model combined with both the high-availability and horizontally scalability architecture serves this goal directly. Further, while having no dependency on Kubernetes, Cortex is built with Kubernetes in mind and most users deploy it in Kubernetes.
We also provide a robust way for users to scale their Prometheus servers and Cortex has resulted in a lot of improvements in Prometheus itself.

### Comparison with Thanos

Thanos is another CNCF project that provides high-availability and long-term storage to Prometheus. Both Thanos and Cortex make different trade-offs that will appeal to different use-cases:


* Cortex is a centralised store while Thanos holds the recent data at the edge in Prometheus servers themselves. This presents a different tradeoff - pushing writes to a central location with Cortex vs pulling data at query time with Thanos. This results in query latency and availability differences.


* Because of the centralised nature and push based architecture of Cortex, you can enforce quality and quantity of the data being stored, and drop data you don’t want to store long-term.


* Multitenancy is built into Cortex which makes a good option for larger organisations that need to keep the data from separate teams separate. Having said that, even Thanos has multitenancy on the roadmap.


* The Thanos architecture allows for incremental adoption and reuse of existing Prometheus deployments, whereas Cortex leverages the built-in Prometheus remote-write API.

With all this, the Cortex and Thanos communities are constantly collaborating with each other, the recent ones being using the cortex-frontend to provide query caching for Thanos and exploration of writing blocks and using the Thanos query path in Cortex. There are plans to collaborate further and the community and usage of both the projects is only growing!

## Incubation State Requirements

1. _Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._

We have the list of public adopters in the repo, some notable users include:

* [Electronic Arts](https://www.ea.com/), is using Cortex for scaling their Prometheus servers.
* [DigitalOcean](https://www.digitalocean.com/), is using Cortex for scaling their Prometheus servers.
* [GoJek](https://www.gojek.com/), is using Cortex to build lens, the unified internal monitoring platform for all the services in its fleet.
* [Grafana Labs](https://www.grafana.com/), uses Cortex to run a commercial hosted Prometheus service.
* [Mayadata](https://www.mayadata.io/), is using Cortex to monitor the storage nodes for the users of its platform.
* [Weaveworks](https://www.weave.works/), uses Cortex to run a commercial hosted Prometheus service.


2. _Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._

We have 8 maintainers 3 of who are independent and 4 of them from 3 different companies. The details are [here](https://github.com/cortexproject/cortex/blob/master/MAINTAINERS).

3. _Demonstrate a substantial ongoing flow of commits and merged contributions._

We are seeing a constant stream of performance improvements and features from the maintainers and community. See the stats here:

* [Commits per week over the last 6 months](https://cortex.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=now-6M&to=now)

* [Issue Opened/Closed per week over the last 6 months](https://cortex.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=now-6M&to=now)

* [New PRs per week over the last 1 year](https://cortex.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)

4. _A clear versioning scheme._

We now have regular releases documented at: https://github.com/cortexproject/cortex/blob/master/RELEASE.md
We’ve only recently started our release process, but have 3 releases out already.

5. _Roadmap_

We're nowhere near done with Cortex and have a lot of plans for future development, and our roadmap currently includes:

* **Write-Ahead Log (In Progress)** - This would enable crash resiliency in Cortex. Right now we replicate each sample n-ways in the ingesters for that but we do have data-loss if enough ingesters crash.
* **Query Parallelisation (In Progress)** - Parallelise the processing of queries further where possible, enabling us to execute massive queries.
* **Simpler runtime configuration** - This would make Cortex easier to operate.
* **Downsampling** - Lets us store less data and over longer periods of time.
* **Recording Rule Substitution** - Right now whenever users add recording rules, they have to wait a long time for the recording rule to populate data before they can switch their dashboards to it. We’re exploring automatically detecting and replacing the original query with its recording rule when appropriate. Once this is in place, we can also detect regular large queries and automatically replace them with recording rules.
* **No dependencies cortex (no NoSQL/Object Store needed)** - This would make Cortex easier to operate on-prem and on bare-metal.

