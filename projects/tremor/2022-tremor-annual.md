# Tremor Sandbox Annual Review

[Tremor](https://www.tremor.rs/) is a cloud-native event-processing engine designed for operator and systems reliability engineer convenience in the design, development, and support of continuous streaming systems deployed and operated at scale.

Tremor was initially designed for the needs of Wayfair’s observability platform engineering teams within our infrastructure organization and was first deployed to production in October 2018 for load-shedding and traffic-shaping of massive logging and metrics workloads. It has since been adopted for transaction orchestration, and our development platforms teams for dynamic code analysis on production PHP deployments to minimize the observable overhead.

Tremor was open-sourced in February 2020, and adopted into the CNCF sandbox in September 2020.

The Tremor devstats page can be found [here](https://tremor.devstats.cncf.io/d/8/dashboardsm).


## Architecture

Tremor is a runtime environment that supports pluggable connectors to consume and contribute events to the outside world. Tremor supports a growing number of connectors to systems such as Kafka, InfluxDB via Telegraf, ElasticSearch through its REST API, the Open Telemetry suite and a number of Google Cloud services.

Tremor's language is a holistic approach to configuring the system starting from defining deployments, connecting components using a query-like syntax and an ETL focus that allows for the creation of pipelines that can be used to transform and enrich events.

## Maintainers

Tremor has seven primary maintainers from three organisations:

*	Darach Ennis and Heinz Gies founded the project and continue to contribute.
*	Matthias Wahl has been with the project since September 2020.
*	Sharon Koech is a graduated LFX Mentee who redesigned the Tremor www and is now our lead documentation maintainer and technical writer.
*	Ramona Luczkiewicz has been with the project since April 2022.
*	Natali Vlatko of Wayfair’s OSPO and Ana Hobden of Hoverbear Consulting are also maintainers of Tremor and assist with mentee selection and with community moderation to ensure that the code of conduct is adhered to by all members of the Tremor community's Discord Server.

## Notable Releases Since the 2021 Annual Review

### v0.12 - New Connectivity via Connectors and Added Deployment Language

This is a community-focused release. Feedback from our LFX mentorships was that writing connectors was hard - this promoted the connectors refactoring in the `0.12` release. The release is a major rewrite of the connector internals and related runtime facilities to deliver better consistency to users and a better programming model for connector authors.

The deployment language work was feedback from production users that favored dropping YAML-based deployment configuration for unifying deployment configuration with the Tremor language. The connectors refactoring allowed for the introduction of the deployment language, associated runtime and API changes to be undertaken holistically.

### v0.13 - Quality of Life Improvements

The work in progress `0.13` release concentrates on Quality of Life improvements over `0.12` that introduce minor backwards compatibility breaks for application writers based on feedback from early adopters of the new deployment model that have migrated applications to the new model.

### v0.14 - Basic Clustering for Cloud-based Autoscaling

RAFT-based consensus mechanism for maintaining cluster state such as registered applications and pipelines.

This provides the basic capabilities for nodes to join and leave the cluster enabling auto-scaling and a shared cluster-wide configuration.

## Notable Community Contributions

### [Plugin Development Kit](https://nullderef.com/blog/plugin-abi-stable/)

The plugin development kit is a huge contribution to the Tremor ecosystem. It allows for the creation of Tremor connectors that are not part of the core binary. This makes contributing to Tremor a significantly easier task and allows for the creation of connectors that are separately maintained. The PDK will be productized into Tremor in a future release once performance overhead is brought into line with the current release.

## Sandbox Activity

### Sandbox Proposal and Inclusion Artefacts

Sandbox proposal [PR #492](https://github.com/cncf/toc/pull/492), [content](https://github.com/cncf/toc/blob/main/proposals/sandbox/tremor.md) and [inclusion](https://lists.cncf.io/g/cncf-toc/message/5273?p=%2C%2C%2C20%2C0%2C0%2C0%3A%3ACreated%2C%2Ctremor%2C20%2C2%2C0%2C76717642) references cited for completeness.

### Tremor Community

Launch of the Tremor community on [Discord](https://chat.tremor.rs/) - this has grown to **235** members, **30** of which are online on a daily basis.

Most of the day to day coding, design work, refactoring, bug hunting sessions and planning are now conducted in the open on Discord in the `General` channel or the `Summer` or `Winter` voice/live presentation channels - and we are continuously evolving our governance so that our processes, practices and decision-making are by the community, for the community.

In addition, we hold bi-weekly [office hours](https://community.cncf.io/events/details/cncf-tremor-community-presents-tremor-office-hours-2022-01-31/) for the community.

### TremorCon 2022

For the second year in a row, we held [TremorCon](https://community.cncf.io/events/details/cncf-tremor-community-presents-tremor-con-2022/)- a one-day annual Tremor conference showcasing the work of our 2nd and 3rd LXF mentorship cohorts and community members. This was hosted on Twitch, and the videos are archived on [YouTube](https://youtube.com/playlist?list=PLNTN4J6tdf22rEnqZlVRWHjiDyRBNy3Iw).

### Tremor Adoption Case Studies Launch

The initial launch of [Tremor use-case studies](https://www.tremor.rs/blog/tags/case-study) for Wayfair’s ongoing adoption and sponsoring of Tremor.

Tremor hosts open office hours for end-users and contributors bi-weekly as part of the [Tremor CNCF Community](https://community.cncf.io/tremor-community/). We have begun canvassing for adoption case studies in this and other forums with our starting set of adoption case studies for community reference and guidance. A list can be found in the [Incubation Stage Requirements](#document-that-is-being-used-successfully-in-production-by-at-least-three-independent-end-users-which-in-the-tocs-judgment-are-of-adequate-quality-and-scope) section.

In addition, we had some adopters talk at TremorCon 2022.

### Tremor Mentorships

#### LFX Spring 2022

- Sasha Pourcelot ( `@scrabsha` ) - [Database Connectors](https://github.com/tremor-rs/tremor-runtime/pulls?q=is%3Apr+author%3Ascrabsha).
- Prashant Mishra ( `@PrimalPimmy` ) - [CI and Release Process Improvements](https://github.com/tremor-rs/tremor-runtime/pulls?q=is%3Apr+author%3APrimalPimmy).

#### LFX Summer 2022

- Rebecca Abli ( `@znwar` ) - [Pluggable Logging](https://github.com/tremor-rs/tremor-runtime/pulls?q=is%3Apr+author%3Aznwar+).
- Carol Geng ( `@carolgeng ` ) - [Hygenic Error Handling and Validation for Pipelines](https://github.com/tremor-rs/tremor-runtime/pulls?q=is%3Apr+author%3Acarolgeng).

## <a name="document-that-is-being-used-successfully-in-production-by-at-least-three-independent-end-users-which-in-the-tocs-judgment-are-of-adequate-quality-and-scope"></a> Incubation Stage Requirements

### Document That is Being Used Successfully in Production by at Least Three Independent End-users Which, in the TOC’s Judgment, are of Adequate Quality and Scope

We have documented a number of use-cases, beginning with our originating sponsor and founding adopter - Wayfair. 

In addition to that, we have a number of organic adoptions from the community:

* [Instadeq](https://www.youtube.com/watch?v=og3kIey0X5E)
* [TietoEvry for the Norwegian Postal Service](https://www.youtube.com/watch?v=ZK6fGYTS-HM)
* [A collaboration with TDEngine](https://tdengine.com/tdengine-and-tremor-solution-for-system-monitoring-and-alerting/)
* [Aiven](https://www.youtube.com/watch?v=6PdTvpydnJM)


### Have a Healthy Number of Committers

Tremor's core codebase has contributions from at least [**17** different organizations](https://tremor.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=commits) and more than [**42**](https://github.com/tremor-rs/tremor-runtime/graphs/contributors) different individuals. This does not count contributions to the web presence, documentation or supporting libraries.

### Demonstrate a Substantial Ongoing Flow of Commits and Merged Contributions

Out of the last **100** merged pull requests to the main repository, **22** were made from non-core contributors (this excludes Dependabot or other automated PRs). The community contributions are significant and valuable in their own right and add a lot to the core value of Tremor as a solution for prospective adopters.

This puts Tremor close to over 20% ratio of team to community PRs.

### A Clear Versioning Scheme

The project's versioning scheme conforms to the [semantic versioning](https://semver.org/spec/v2.0.0.html) scheme and is implemented as follows:

- Major: currently 0 as the core implementation is active and in progress.
- Minor: is incremented for new major features or breaking changes.
- Patch: is incremented for bug fixes and minor new features.

The project has a 3-6 month release cadence for minor increments since version v0.7.0. Patch releases are produced continuously as production needs demand.

For our releases, maintainers follow the [release process](https://github.com/tremor-rs/tremor-runtime/blob/main/RELEASE_PROCESS.md) now fully run by an [automated release process](https://www.tremor.rs/blog/2022/06/03/LFX-Blog-Prashant) contributed by a CNCF mentee.

For security patch releases, we follow a separate [disclosure process](https://github.com/tremor-rs/tremor-runtime/blob/main/SECURITY.md) and announce CVEs accordingly. For example, we have published one [CVE](https://nvd.nist.gov/vuln/detail/CVE-2021-39228) against Tremor for a language behavior that could have been exploited and was fixed in a [security patch release](https://github.com/tremor-rs/tremor-runtime/pull/1217).

## Roadmap

We do not currently have a roadmap available as a publicly accessible GitHub project but are prototyping the changes to governance, documentation, and processes that would support a roadmap for the current release, a tentative roadmap for the next release, and a horizon looking 2 releases into the future.

These changes are forward-looking.

In broad terms, and subject to community feedback and changes, there are:

### V0.14 - Basic Clustering for Cloud-based Autoscaling

RAFT-based consensus mechanism for maintaining cluster state such as registered applications and pipelines.

This provides the basic capabilities for nodes to join and leave the cluster enabling auto-scaling and a shared cluster-wide configuration.

### V0.15 - Clustering Enhancements

Introduction of a micro / macro architecture using a ring-based topology for the macro ring in the spirit of Riak's V-Nodes for ease of scalability

### Release Strategy

As a production-focused project, Tremor maintainers favor essential production-impacting features, enhancements and fixes over roadmap completeness.

This set of practical planning and delivery constraints of roadmap definition and release strategy against a planned and visible roadmap is not thoroughly documented at this time. This is acknowledged as a flaw in our governance model and is being reviewed.

The V0.12 release will be the final release in the current governance model that was introduced with the open-sourcing of the project. Subsequent releases will use an evolution of the governance model driven by community feedback and supported by community decisions in regularly held Tremor Open Office Hours where changes are discussed, feedback gathered and decisions voted upon.

We have referenced the known governance model issues in this review document.
