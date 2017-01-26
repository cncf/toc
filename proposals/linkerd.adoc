== linkerd Proposal

*Name of project:* linkerd

*Description:* Linkerd is an open source service mesh for cloud-native applications. It is a layer 5/7 proxy that provides fully decentralized request-level reliability and optimization to service communication, including latency-aware load balancing, retries, failure handling, circuit breaking, and deadlines. It provides granular metrics (both of itself and of downstream services), distributed tracing, TLS, a powerful logical routing layer, and integration with most common service discovery mechanisms.

Applications use linkerd by proxying HTTP, gRPC, or Thrift calls through a local linkerd instance. In many cases, integration with linkerd is a config change (e.g. setting an http_proxy environment variable) rather than a code change. Linkerd acts as a transparent proxy, and so a call to (e.g. "http://foo/bar") will be routed to the appropriate cluster, resolved in service discovery, load balanced, retried, and returned without application code being aware.

Linkerd is built on top of Finagle (https://twitter.github.io/finagle/), the high-concurrency RPC library that powers companies like Twitter, Soundcloud, Pinterest, and ING Bank. By staying close to the Finagle codebase, linkerd takes advantage of Finagle's extensive production usage (https://github.com/twitter/finagle/blob/develop/ADOPTERS.md). Linkerd’s primary author is a core contributor to Finagle.

*Sponsor / Advisor from TOC*: Jonathan Boulle <jonathan.boulle@coreos.com>

*Unique Identifier*: linkerd

*License*: Apache License v2.0

*Maturity Level*: Inception

*Source control repositories*:

* https://github.com/linkerd

*Initial Committers (leads)*:

* Oliver Gould (Buoyant) [@olix0r]
* Alex Leong (Buoyant) [@adleong]

*Infrastructure requirements*: CI and potentially CNCF Community Cluster access. We run integration tests per commit (e.g. testing Kubernetes API integration) and periodic high-volume stress tests.

*Issue tracker*: https://github.com/linkerd/issues

*Mailing lists*: https://groups.google.com/forum/#!forum/linkerd-users

*Website*: https://linkerd.io/

*Community*:  Active & growing!

* Primarily focused on https://slack.linkerd.io (~400 members)
* ~1200 GitHub stars
* ~25 contributors

*Release methodology and mechanics*: Code review on all changes. Releases triggered manually by developer team. Using pre-1.0 semver semantics.

*Social media accounts*: Twitter: @linkerd, Slack: http://slack.linkerd.io/

*Existing sponsorship*: https://buoyant.io/

*External Dependencies*: Runtime: JVM. Build-time: Finagle, netty

*Statement on alignment with CNCF mission*:

Linkerd directly addresses the communications challenges of running microservices at scale, and integrates directly with orchestrated environments such as Kubernetes and DC/OS. Linkerd integrates with two existing CNCF projects: it can use the Kubernetes API directly as a service discovery mechanism, and it exposes metrics in Prometheus format. See this blog post (https://blog.buoyant.io/2016/10/04/a-service-mesh-for-kubernetes-part-i-top-line-service-metrics/) for an example of linkerd, Kubernetes and Prometheus working together.

*Production Usage*:

* Monzo: used in production; k8s-based infra (https://monzo.com/blog/2016/09/19/building-a-modern-bank-backend/)
* NCBI: used in production w/Consul
* Quid: used in production
* Douban: used in production
* Houghton Mifflin Harcourt: used in production; DC/OS-based infra
* Olark: used in production; k8s-based infra (https://www.youtube.com/watch?v=r0nhQwbe8OY&index=107&list=PLj6h78yzYM2PqgIGU1Qmi8nY7dqn9PCr4)

*Other Contributors:*

Currently ~25; >50% non-Buoyant by headcount: https://github.com/BuoyantIO/linkerd/graphs/contributors?type=a
