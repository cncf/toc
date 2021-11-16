# Tremor Sandbox Annual Review

Tremor is a cloud-native event processing engine designed for operator and systems reliability engineer convenience in the design, development, and support of continuous streaming systems deployed and operated at scale.

Tremor was initially designed for the needs of Wayfair’s observability platform engineering teams within our infrastructure organization and was first deployed to production in October 2018 for Load-shedding and Traffic shaping of massive Logging and Metrics workloads.. It has since been adopted for transaction orchestration, and our development platforms teams for dynamic code analysis on production PHP deployments to minimize the observable overhead.

Tremor was open-sourced in February 2020, and adopted into the CNCF early-stage sandbox in September 2020.

## Architecture

Tremor is a runtime environment that supports pluggable connectors to consume and contribute events to the outside world. Tremor supports a growing number of connectors to systems such as Kafka, InfluxDB via Telegraf, ElasticSearch through its REST API, and using AMQP to systems such as RabbitMQ or the Ably SaaS.

Tremor hosts business logic implemented in an expression-oriented scripting language designed for efficient ETL and processing of semi-structured JSON-like data which it decodes using SIMD-accelerated JSON processing via the [rust port of SIMD-JSON](https://github.com/simd-lite/simd-json) which is also maintained by the tremor maintainers.

Tremor also supports a query language for describing complex event flows with a rich palette of pluggable operators for quality of services - such as batch event processing, and backpressure operators, and a write-ahead-log for transactional delivery. The tremor runtime supports acknowledgments, guaranteed delivery, and circuit breaker primitives internally and connectors are QoS aware - relieving the burden of recovery and restoring systems to good operating conditions from operators of tremor-based systems.

Tremor is currently adding a deployment language to further support operators with hygienic errors and hygienic deployments to allow larger tremor-based systems to be composed modularly and for complex components ( sub-applications ) to be reused.

## Maintainers

Tremor has 6 primary maintainers.

*	Darach Ennis and Heinz Gies founded the project and continue to contribute.
*	Matthias Wahl has been with the project since September 2020.
*	Sharon Koech is a graduated LFX Mentee who redesigned the tremor www and is now our lead documentation maintainer and technical writer.
*	Natali Vlatko of Wayfair’s OSPO and Ana Hobden of Hoverbear Consulting are also maintainers of tremor and assist with mentee selection for mentorships on tremor and with community moderation to ensure that the code of conduct is adhered to by all members of the tremor communities discord service.

## Notable releases since open source

### V0.7	The first open-source release of tremor

This version introduced the governance and community model and was tested against a panel of Wayfair external OSS stakeholders and industry experts to align the project goals with OSS expectations and formed the basis of our open source strategy. One of our project advisors went on to deploy Tremor into production in his organization where Tremor is still used.


### V0.7.3	Introduction of the pipeline state mechanism for stateful processing

Until this time logic deployed within tremor was purely event-based and stateless. With this release the state mechanism allowed query applications to maintain state over time and events.


### V0.8	- The introduction of script and query language support for modules

Growing size and complexity of production deployments necessitated the introduction of modularity features. This release saw the introduction of a preprocessor, a functional programming language extension to the scripting language and a general purpose syntax for defining and using modules in the script and query language. A library of reusable functions largely built in rust and exposed to the query and scripting languages was developed. Applications based on tremor can now be defined with reusable modular logic developed to support application domains.


### V0.9	- The introduction of high-level QoS primitives and linked transports

This release was a return to the grass roots of tremor and focused on primitives for quality of service. The runtime internals were leveraged to underlay the mechanics to drive circuit breakers, propagate message acknowledgements and enable tremor-based applications to selectively introduce recoverable guaranteed delivery through connectors refactored for this purpose, and runtime mechanisms to support these guarantees.


### V0.10	- Breaking changes for string interpolation

This was a relatively minor release based on features but plugged a design flaw in Tremor’s support for string interpolation where the interpolation demarcation overlapped with regular expression syntax. The interpolation syntax was changed from `{..}` pairs to `#{..}` to resolve the design flaw and the solution verified against production regression cases.


### V0.11	- The introduction of CNCF OpenTelemetry as a source or sink connector

This release saw the introduction of CNCF OpenTelemetry source and sink connectors allowing Tremor to interpose itself between OpenTelemetry based systems. This release also added CNCF NATS connectivity and a `stdin` connector. This is the first suite of connectors built in tremor that have an associated set of modules to assist tremor developers with their OpenTelemetry based processing needs.


### V0.12 - First fully community-focused release

This release is the current primary development target. 

See the roadmap section later in this document.

## Notable community contributions

### [AMQP Connectivity](https://github.com/tremor-rs/tremor-runtime/pull/1001/)

Nokia Corporation contributed the AMQP connectivity in tremor. This allows tremor to connect to AMQP exchanges and to publish or consume events over AMQP such as to RabbitMQ or to Ably’s SaaS using their AMQP connectivity.

### [SNMalloc](https://github.com/microsoft/snmalloc)

Microsoft Corporation - we leverage the snmalloc allocator by Matt Parkinson et al at Microsoft Research and used the mimalloc allocator ( also by MSR ) previously. This collaboration is ongoing and has seen contributions from tremor to that project and vice-versa.

### [New Relic Connectivity](https://github.com/tremor-rs/tremor-runtime/pull/866)

Grover contributed the New Relic connectivity to tremor.

## Notable community adoptions

Wayfair was the originator and remains the primary sponsor of tremor development supporting 3 of the tremor maintainers. A number of production Wayfair use cases have been captured and documented.

* [PHP dead code detection](https://www.tremor.rs/blog/2021/11/09/php-dead-code-detection)
* [Unified Observability Platform](https://www.tremor.rs/blog/2021/11/08/uop)
* [Multi-Participant Transaction Orchestration](https://www.tremor.rs/blog/2021/11/07/search)
* [Modularity](https://www.tremor.rs/blog/2021/11/06/modularity)
* [Kubernetes and Sidecars](https://www.tremor.rs/blog/2021/11/05/kubernetes-sidecars)
* [Data Flow](https://www.tremor.rs/blog/2021/11/04/data-flow)
* [Data Distribution](https://www.tremor.rs/blog/2021/11/03/data-distribution)
* [Traffic Shaping](https://www.tremor.rs/blog/2021/11/02/traffic-shaping)	

Tremor is also used by [Instadeq](https://instadeq.com/) and others and we are collaborating with other organizations such as [Vectorized](https://vectorized.io/) ( [Redpanda](https://vectorized.io/redpanda/), [collaboration](https://github.com/tremor-rs/tremor-redpanda) ), [Ably](https://ably.com/) ( Messaging SaaS ) and others such as the [Vector](https://vector.dev/) team ( was timber.io  ) at DataDog. We are also working with other CNCF Projects through our support of CNCF OpenTelemetry, CNCF NATs, adoption of CNCF TiKV's raft implementation in a future release, and have had discussions with the CNCF Services Workflow specification authors on potential adaptation of Services Workflow to tremor.

We have yet to capture adoptions using use case studies but have begun this process by recording and blogging Wayfair’s primary production use cases and using these as examples to other adopters to reference when blogging about their adoption of Tremor in production.
	
## Sandbox Activity

### Sandbox Proposal and Inclusion Artefacts

Sandbox proposal [PR #492](https://github.com/cncf/toc/pull/492), [content](https://github.com/cncf/toc/blob/main/proposals/sandbox/tremor.md) and [inclusion](https://lists.cncf.io/g/cncf-toc/message/5273?p=%2C%2C%2C20%2C0%2C0%2C0%3A%3ACreated%2C%2Ctremor%2C20%2C2%2C0%2C76717642) references cited for completeness.

### Tremor Community

Launch of the tremor community based on Discord at [chat.tremor.rs](https://chat.tremor.rs/) - this has grown to 187 members, of which ~30 are online on a daily basis.

Most of the day to day coding, design work, refactoring, bug hunting sessions and planning are now conducted in the open on the tremor discord in the `General` voice/live presentation channel - and we are continuously evolving our governance so that our processes, practices and decision making are by the community, for the community.

For example, the latest RFC for aggregate functions in tremor-script was a community contribution that was ratified during our community office hours. 

We are yet to update our governance model to reflect that RFC acceptance is now community office hour ratified and voted upon - this is planned.

### Tremor Web and Documentation Redesign

Organizing a review of the tremor web experience by LF/CNCF Technical writers. Celeste Horgan’s review led to our ongoing Docs Project and an LFX Mentorship which resulted in tremor’s new web presence.

- [Tremor Docs Project](https://github.com/tremor-rs/tremor-www/projects/1)
- Adoption of Netlify for tremor’s www
- Redesign and information architecture for www - led by Sharon Keoch
- Documentation refactoring and enhancement - led by Sharon Keoch
- Documentation Github Project - new process introduced by maintainers
- Sharon adjoined to documentation/www lead as a core maintainer
- Adoption of Equinix and the CNCF Cluster for [benchmarking reports](https://www.tremor.rs/benchmarks/ )

### TremorCon

[TremorCon](https://www.youtube.com/watch?v=T-WukekIUMs&list=PLNTN4J6tdf20vy14FVOazLTdou_8xyvfe) - a one day annual tremor conference showcasing the work of our first two LXF mentorship cohorts and community members. This was hosted on Twitch with a live Q&A with pre-recorded presentations. The videos are archived on YouTube.

### Tremor Adoption Case Studies Launch

The initial launch of [tremor use case studies](https://www.tremor.rs/blog/tags/case-study) for Wayfair’s ongoing adoption and sponsoring of Tremor.


Tremor hosts open office hours for end-users and contributors every other week as part of the [Tremor CNCF Community](https://community.cncf.io/tremor-community/). We have begun canvassing for adoption case studies in this and other forums with our starting set of adoption case studies for community reference and guidance.


### Tremor Mentorships

#### LFX Spring 2021

- Jigyasa ( `@jigyasak05` ) - GCP Cloud Storage and Pub/Sub connectivity
- Nupur ( `@murex971` ) - Syslog support
- Rohit ( `@kurious-diru` ) - Property-based testing with Erlang/OTP and Quickcheck
- Akshat ( `@humancalico` ) - Benchmark CI prototype work

#### LFX Summer 2021

- Yatin ( `@yatinmaan` ) - Subqueries in tremor query language DSL
- Sharon ( `@skoech` ) - WWW and documentation redesign

#### LFX Autumn 2021

- Daksh ( `@dak-x` ) - Amazon AWS S3 Connectivity
- 1 failed candidate - Error handling improvements and redesign

#### Tremor Summer of Code 2021

- Sanskar ( `@aryan9600` ) - gRPC - Google Remote Procedure Call connectivity
- Mario ( `@marioortizmanero` ) - Plugin Development Kit ( PDK )

## Incubation Stage Requirements

Largely meeting, but with gaps as cited below.

Our governance model, processes and capture of adoption case studies is currently being reworked for more effective community participation at this time.

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgment, are of adequate quality and scope

We have 8 use cases documented from our originating sponsor and founding adopter - Wayfair. We have begun the capture process for other adoptions in the OSS community around tremor.

### Have a healthy number of committers

Tremor has contributions from [14 different organizations](https://tremor.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=commits) and more than [31](https://github.com/keptn/keptn/graphs/contributors) different individuals.

### Demonstrate a substantial ongoing flow of commits and merged contributions

Out of the last 60 merged pull requests to the main repository 19 were made from non-core contributors (this excludes Dependabot related PRs). The community contributions are significant and valuable in their own right and add a lot to the core value of tremor as a solution for prospective adopters.

​​### A clear versioning scheme

The projects versioning scheme conforms to the [semantic versioning](https://semver.org/spec/v2.0.0.html) scheme and is implemented as follows

- Major: currently 0 as the core implementation is active and in progress
- Minor: is incremented for new major features or breaking changes
- Patch: is incremented for bug fixes and minor new features

The project has a 3-6 month release cadence for minor increments since version v0.7.0. Patch releases are produced continuously as production needs demand.

For our releases, maintainers follow the [release process](https://github.com/tremor-rs/tremor-runtime/blob/main/RELEASE_PROCESS.md) supported by a [release script](https://github.com/tremor-rs/tremor-runtime/blob/main/release.sh).

For security patch releases, we follow a separate [disclosure process](https://github.com/tremor-rs/tremor-runtime/blob/main/SECURITY.md) and announce CVE’s accordingly. For example, we have published one [CVE](https://nvd.nist.gov/vuln/detail/CVE-2021-39228) against tremor for a language behavior that could have been exploited and was fixed in a [security patch release](https://github.com/tremor-rs/tremor-runtime/pull/1217).

## Roadmap

We do not currently have a roadmap available as a publicly accessible GitHub project but are prototyping the changes to governance, documentation, and processes that would support a roadmap for the current release, a tentative roadmap for the next release, and a horizon looking 3 releases into the future.

These changes are forward looking.

In broad terms, and subject to community feedback and changes, there are:

### V0.12 - new connectivity via connectors and added deployment language

This is a community focused release. Feedback from our LFX 
mentorships was that writing connectors was hard - this promoted the 
connectors refactoring in the v0.12 release.

The deployment language work was feedback from our production users and 
is tightly coupled to runtime changes that spring from the connectors 
rework.

### V0.13 - Basic clustering for cloud-based autoscaling

RAFT-based consensus mechanism (using CNCF TiKV’s raft implementation)

Ring-style topology inspired by riak_core’s variant of the Dynamo 
architecture from Amazon AWS - including adoption of riak-core style 
virtual nodes ( vnode ).

Native tremor Network protocol and cluster gossip protocols based on 
web-sockets.

Introduction of micro ( raft based ) and macro ( ring topology based ) layers to create a unified view on a tremor cluster with the ability of nodes joining and leaving in an automated fashion to enable cloud based auto-scaling of nodes.

The micro ring is a subset of the whole ( macro ) cluster of nodes where raft consensus is hosted as a service on behalf of the wider cluster allowing consensus and cluster works to be independently scaled.

### V0.14 - Clustering enhancements

Advanced clustering features such as generalized rate-limit based load balancing, distributed cluster-wide key-value storage to dynamically change configuration and separation of ingress and egress groups, cluster-wide event notifications service.

These features may also see the introduction of language features that allow tremor deployments to be composed with distribution of events mediated by cluster aware partitioning, grouping, distribution and orchestration logic for applications deployed into a tremor cluster.

### Release Strategy

As a production focused project, Tremor maintainers favor essential production-impacting features, enhancements and fixes over roadmap completeness.

This set of practical planning and delivery constraints of roadmap definition and release strategy against a planned and visible roadmap is not thoroughly documented at this time. This is acknowledged as a flaw in our governance model and being reviewed.

The V0.12 release will be the final release in the current governance model that was introduced with the open sourcing of the project. Subsequent releases will use an evolution of the governance model driven by community feedback and supported by community decisions in regularly held Tremor Open Office Hours where changes are discussed, feedback gathered and decisions voted upon.

We have referenced the known governance model issues in this review document.
