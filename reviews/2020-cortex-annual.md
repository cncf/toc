# Cortex Annual Review 2020

## Background

[Cortex Sandbox proposal](https://github.com/cncf/toc/pull/125)
[Original CNCF TOC Meeting slides](https://docs.google.com/presentation/d/190oIFgujktVYxWZLhLYN4q8p9dtQYoe4sxHgn4deBSI/edit#slide=id.g380c8a0114_0_178)

[Cortex](https://github.com/cortexproject/cortex) is a horizontally scalable, highly available, multi-tenant, Prometheus API-compatible service that offers a long-term storage solution.

For teams looking for a Prometheus-like solution that offers the following over Prometheus:

* Long-term metrics storage in a variety of cloud based and on-prem NoSQL and ObjectStorage data stores

* Tenancy model supporting commercial SaaS offerings or large/multiple Kubernetes installations requiring data separation

* On-demand Prometheus instance provisioning

* A highly-available system that benefits from cloud-native architectures run with Kubernetes

* A highly scalable Prometheus experience that horizontally scales

* The ability to handle large metric topologies in a single instance with no federation

Cortex was presented at the [CNCF TOC meeting on 6/5/2018](https://docs.google.com/presentation/d/190oIFgujktVYxWZLhLYN4q8p9dtQYoe4sxHgn4deBSI/edit#slide=id.g25ca91f87f_0_0). The project and the community has grown a lot since then, and the project is more active and mature now.

Notable improvements in the last year include:

* We started doing versioned releases at a regular cadence.
* An easy to use single-process version for people to get started quicker and scale.
* Up to 10x query performance increase due to techniques such as query parallelisation and sharding.
* Cortex now uses up to 50% less disk space.
* Cortex’s alerting and recording rule layer is now horizontally scalable.
* Cortex now uses a Write Ahead Log (WAL) for higher crash resilience.
* Because of the WAL and the single process/binary mode, Cortex is now much more stable and easier to run, with more improvements on their way.

Further, a lot of the work in Cortex also involved improvements in upstream Prometheus. Here are a few improvements the Cortex team submitted to the Prometheus project:

* https://github.com/prometheus/prometheus/pull/4588
* https://github.com/prometheus/prometheus/pull/5707
* https://github.com/prometheus/tsdb/pull/642
* https://github.com/prometheus/prometheus/pull/5316
* https://github.com/prometheus/prometheus/pull/5131

## Alignment with Cloud Native

Cortex fully supports the CNCF’s goal for scalability: "Ability to support all scales of deployment, from small developer-centric environments to the scale of enterprises and service providers."

There are many ways to provide a scalable and available metric system for Kubernetes. Cortex, with its tenancy model combined with both the high-availability and horizontally scalable architecture, serves this goal directly. Further, while having no dependency on Kubernetes, Cortex is built with Kubernetes in mind and most users deploy it in Kubernetes.

Cortex also provides a robust way for users to scale their Prometheus servers, and has resulted in a lot of improvements in Prometheus itself.

### Comparison with Thanos

Thanos is another CNCF project that provides high-availability and long-term storage to Prometheus. Both Thanos and Cortex make different trade-offs that will appeal to different use-cases.

* Cortex is a centralized store while Thanos holds the recent data at the edge in Prometheus servers themselves. This presents a different tradeoff — pushing writes to a central location with Cortex vs. pulling data at query time with Thanos. The tradeoff results in query latency and availability differences. Having said that, Thanos now supports an experimental receive service which can be used to centralise the data.

* Multitenancy is built into Cortex, which makes it a good option for larger organisations that need to keep the data from separate teams separate. Note: Thanos has multitenancy on the roadmap.

* Cortex leverages the built-in Prometheus remote-write API whereas the Thanos architecture allows for incremental adoption and reuse of existing Prometheus deployments.

* Cortex uses a lot of caching to speed up queries and we are actively working with the Thanos community to bring much of the same improvements to Thanos.

With all this, the Cortex and Thanos communities are [constantly collaborating](https://twitter.com/ThanosMetrics/status/1230065830391664640) with each other, the recent ones being using the Cortex query-frontend to provide query caching for Thanos and exploration of writing blocks and using the Thanos query path in Cortex. There are plans to collaborate further and the community and usage of both the projects is only growing!

This is further aided by the large overlap of Prometheus, Cortex, and Thanos maintainership and close coordination between all three projects.

## Questions for annual review

1. _Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats._

* [Commits per week over the last 12 months](https://cortex.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=w&var-repogroups=All)

* [New PRs per week over the last 1 year](https://cortex.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)

You can see how we’re showing healthy activity (44 commits per week average over the year) and that the project is constantly evolving and improving.

2. _How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)_

We have eight maintainers from four different companies. We’ve added three maintainers and had three maintainers step down due to changes of priorities/companies. The details are [here](https://github.com/cortexproject/cortex/blob/master/MAINTAINERS). We also modified our governance to make sure that each company gets only two votes, to make sure one company cannot take over the project.

3. _What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)_

We have [the list of public adopters in the repo](https://github.com/cortexproject/cortex/blob/master/ADOPTERS.md). Some notable users include:

* [Electronic Arts](https://www.ea.com/), is using Cortex for scaling their Prometheus servers.
* [DigitalOcean](https://www.digitalocean.com/), is using Cortex for scaling their Prometheus servers.
* [GoJek](https://www.gojek.com/), is using Cortex to build Lens, the unified internal monitoring platform for all the services in its fleet.
* [Grafana Labs](https://www.grafana.com/), uses Cortex to run a commercial hosted Prometheus service.
* [Mayadata](https://www.mayadata.io/), is using Cortex to monitor the storage nodes for the users of its platform.
* [Weaveworks](https://www.weave.works/), uses Cortex to run a commercial hosted Prometheus service.

In the past six months, we’ve focused on improving the documentation and created a [website](https://cortexmetrics.io/) which made for a better on-boarding experience. We’re seeing more and more users adopting Cortex (not yet in ADOPTERS.md because they can’t make it public yet) and being active in our Slack channel. Based on the trends we’re seeing, we’re confident that the list of end users will increase dramatically in 2020. Further, we’re working on adding case-studies with end users to our website.

4. _How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)_

We think Cortex has done really well against the goal of providing a horizontally scalable multi-tenant Prometheus service. We’ve made significant improvements in our query and ingest path that made it orders of magnitude more performant since we joined the CNCF. Cortex can confidently handle a significant scale (over 100 mil+ active series) with ease. We’ve also made it easier to set up and operate; with the team focused on improving those aspects further to increase adoption.

5. _What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?_

The current goals of the project are:

1. **Increase adoption**: Cortex has been running in production for over three years now and can handle significant scale with ease. We’ve made the on-boarding and operations process easier, and we are now focusing on increasing adoption. This will be through case studies of existing users, improved documentation, guides, and dedicated blog posts.


2. **Make Cortex even easier to run**: Though operating Cortex has gotten simpler, it's still quite complex with several moving parts and a NoSQL store dependency. We’re working with the Thanos community to see if we can leverage the TSDB blocks to provide an ObjectStore-only backed Cortex cluster. The NoSQL store gives us major wins in performance, but we’re now load-testing the ObjectStore backed system to see if we can have sufficient performance and scalability. We’re also working on a single process mode where one can just horizontally scale the “Cortex” processes instead of having to manage several microservices. If both these experiments are fruitful, then we’ll soon have a very simple way to operate the system that only has a dependency on an ObjectStore. 

See also some of the roadmap for the project: https://grafana.com/blog/2020/01/21/the-future-of-cortex-into-the-next-decade/

6. _How can the CNCF help you achieve your upcoming goals?_

We need someone to do user interviews and write up the case studies, so it would be great if we could get someone from the CNCF to help with different case studies we want to do. 
We would also benefit a lot from a technical writer to improve the documentation and version the documentation published on the website. 
We’re ready to launch the 1.0 version of Cortex and we request being funneled interview, article, and blogging opportunities from media for the launch.

7. Do you think that your project meets the criteria for incubation?

Yes, please see our incubation application here: https://github.com/cncf/toc/pull/315.
