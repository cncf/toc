# Thanos Proposal
 
**Name of project:** Thanos

**Description:** 

[Thanos](http://thanos.io) is a metric system that aims for being simple and cheap option to centralize and scale out the Prometheus based systems. 
It achieves that by: 

* Gradual and incremental installation model. Users can seamlessly transform vanilla Prometheus system into partial Thanos, 
optionally upgrade into fully fledged Thanos deployment with near unlimited retention capabilities in separate isolated steps.
* Single, widely available, cheap and optional dependency: object storage
* Built-in Prometheus HA support and downsampling capabilities
* Flexible deployment model due to a cloud-native, microservice approach using gRPC.

Concretely the aims of the project are:
* Global query view of metrics.
* Unlimited retention of metrics, including downsampling.
* High availability of components, including support for Prometheus HA.

**Statement on alignment with CNCF mission:**

The Thanos project aligns well with the mission of the CNCF by empowering organizations to build, run and monitor applications at scale.

It can be used to monitor and alert on microservices and infrastructure that spans over not only multiple clusters but over 
multiple clouds with a simple approach to adoption and scaling.

Thanos aims to enable people to take their current Prometheus ecosystem and scale it out beyond their current capabilities, 
gradually and without downtime. 

Building on top of the work that is being done by the Prometheus community we can scale out to give Thanos users a 
global view of all their metrics without the need of aggregations.

Thanos has similar goals to Cortex project that is already a CNCF Sandbox project. However, Thanos approached the 
problem with the focus on simplicity and gradual installation. It takes advantage of having only single dependency 
(object storage) and is integrating with Prometheus in slightly different way. On the other hand Cortex is more mature 
and have more performance optimizations. As per [No Kingmakers & One Size Does Not Fit All](https://github.com/cncf/toc/blob/master/PRINCIPLES.md#no-kingmakers--one-size-does-not-fit-all) CNCF principle we think it makes
sense to have both projects part of CNCF. Worth to note that we are already collaborating with Cortex team on reusable improvements
that would benefit both projects. In fact, we believe that being together in CNCF will improve the collaboration even further.

**Sponsors / Advisors from TOC:** Xiang Li

**Unique identifier:** Thanos

**Preferred maturity level:** Sandbox

**License:** [Apache License 2.0](https://github.com/improbable-eng/thanos/blob/master/LICENSE)

**Source control repositories:** https://github.com/improbable-eng/thanos

We also own neutral github organization which we plan to use: https://github.com/thanos-io
 
**External Dependencies:**
 
* Golang ([BSD](https://github.com/golang/go/blob/master/LICENSE))
* gogo/protobuf ([BSD](https://github.com/gogo/protobuf/blob/master/LICENSE))
* GRPC-go ([Apache Software License 2.0](https://github.com/grpc/grpc-go/blob/master/LICENSE))
* Jaeger ([Apache Software License 2.0](https://github.com/jaegertracing/jaeger/blob/master/LICENSE))
* OpenTracing-go ([Apache Software License 2.0](https://github.com/opentracing/opentracing-go/blob/master/LICENSE))
* Prometheus ([Apache Software License 2.0](https://github.com/prometheus/prometheus/blob/master/LICENSE))
* Prometheus/tsdb ([Apache Software License 2.0](https://github.com/prometheus/tsdb/blob/master/LICENSE))

**Initial Committers:**

* Bartek Płotka - Improbable - from Nov 2017
* Fabian Reinartz - Google - Nov 2017 - Jan 2018  

**Current maintainers:**

[Public list](https://github.com/improbable-eng/thanos/blob/master/MAINTAINERS.md)

* Bartek Plotka - Improbable
* Dominic Green - Improbable
* Frederic Branczyk - RedHat
* Giedrius Statkevičius - Adform
* Povilas Versockas - Utility Warehouse

**Infrastructure requests (CI / CNCF Cluster):**

* Prow CI 

**Communication Channels:**

* Improbable Public Slack: https://improbable-eng.slack.com/
* Issue tracker: https://github.com/improbable-eng/thanos/issues

**Issue Tracker:** Github

**Website:** https://thanos.io

**Release methodology and mechanics:** 

* Semantic versioning
* Release every 6 weeks
* Mostly automated release process: [details](https://thanos.io/release-process.md/)

**Social media accounts:** https://twitter.com/ThanosMetrics

**Existing sponsorship:** [Improbable](https://improbable.io)

**Community size:**

* 3550+ stars
* 110+ contributors
* 400+ forks

**Production usage:**

Companies that are actively using Thanos in production and publicly share that:

* [Alibaba](https://www.alibaba.com/)
* [Adform](https://site.adform.com/)
* [Adobe](https://www.adobe.com/)
* [Cdiscount](https://www.cdiscount.com/)
* [GitLab](https://gitlab.com)
* [eBay](https://www.ebay.com) 
* [HelloFresh](https://www.hellofresh.com/)
* [Improbable](https://improbable.io)
* [Jetstack](https://www.jetstack.io/)
* [Meltwater](https://www.meltwater.com/) 
* [Monzo](https://monzo.com/)
* [Qonto](https://qonto.eu/)
* [Seznam.cz](https://www.seznam.cz/)
* [Utility Warehouse](https://www.utilitywarehouse.co.uk/)
* [Red Hat](https://www.redhat.com/en)
* [Tiket.com](https://www.tiket.com/)
* [uSwitch](https://www.uswitch.com/)



