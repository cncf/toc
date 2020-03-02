# CNCF SIG Observability Charter

- [CNCF SIG Observability Charter](#cncf-sig-observability-charter)
  - [Introduction](#introduction)
  - [Mission](#mission)
  - [Areas considered in Scope](#areas-considered-in-scope)
  - [Areas considered out of Scope](#areas-considered-out-of-scope)
  - [Roadmap & Initial Efforts](#roadmap--initial-efforts)
  - [Governance](#governance)
  - [Operations](#operations)

*Initially authored by [Matthew Young][matthew young] with grateful review and
contributions from:
[Alex Nauda][Alex Nauda],
[Alois Reitbauer][Alois Reitbauer],
[Bartłomiej (Bartek) Płotka][Bartłomiej (Bartek) Płotka],
[Daniel Khan][Daniel Khan],
[Daniel Prata][Daniel Prata],
[Lincoln Sward][Lincoln Sward],
[Matthias Loibl][Matthias Loibl],
[Michael Hausenblas][Michael Hausenblas],
[Ricardo Aravena][Ricardo Aravena],
[Richard Hartmann][Richard Hartmann],
[Sergey Kanzhelev][Sergey Kanzhelev],
[Steve Flanders][Steve Flanders],
[Ted Young][Ted Young],
[Tigran Najaryan][Tigran Najaryan],
[Tommy Chong][Tommy Chong],
and [Umair Ishaq][Umair Ishaq].*

<!-- TODO: please put github names here -->
[Matthew Young]:              https://github.com/halcyondude
[Alex Nauda]:                 @
[Alois Reitbauer]:            @
[Bartłomiej (Bartek) Płotka]: @
[Daniel Khan]:                @
[Daniel Prata]:               @
[Lincoln Sward]:              @
[Matthias Loibl]:             @
[Michael Hausenblas]:         @
[Ricardo Aravena]:            @
[Richard Hartmann]:           @
[Sergey Kanzhelev]:           @
[Steve Flanders]:             @
[Ted Young]:                  @
[Tigran Najaryan]:            @
[Tommy Chong]:                @
[Umair Ishaq]:                @

## Introduction

This document describes the purpose and operations of the Cloud Native
Computing Foundation ([CNCF]) Special Interest Group ([SIG]) on Observability.

This [SIG] focuses on topics pertaining to the observation
of [cloud native][cn-def] workloads. Additionally, it produces supporting
material and best practices for end-users and provides guidance and
coordination for CNCF projects working within the SIG’s scope.

A full list of [CNCF projects][projs] can be found at [landscape.cncf.io].

[cncf]: https://www.cncf.io
[projs]: https://www.cncf.io/projects
[landscape.cncf.io]: https://landscape.cncf.io
[sig]:  https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md
[cn-def]: https://github.com/cncf/toc/blob/master/DEFINITION.md

## Mission

Consistent with the CNCF [SIG] definition, the mission of SIG Observability
is to:

- Foster and grow the ecosystem of observability related projects, users, and
  maintainers.
- Identify and report gaps in the CNCF's project portfolio on topics of
  observability to the TOC and the wider CNCF community.
- Collect, curate, champion, and disseminate patterns and current best practices
  related to the observation of cloud-native systems that are effective and
  actionable.
- Educate and inform users with unbiased, accurate, and pertinent information.
- Educate and help other CNCF projects in regarding observability techniques and
  practices available within the CNCF.
- Provide and maintain a vendor-neutral venue for relevant thought validation,
  discussion, and project feedback.
- Provide a ladder for community members to become involved with the technical
  oversight of projects within the SIG's scope in an open, transparent, and
  inclusive way.

## Areas considered in Scope

Observability focuses on patterns, projects, tools, and techniques related to
topics such as:

- Methodologies for instrumenting, collecting, processing, storing, querying,
  curating, and correlating observational data such as metrics, logging/events,
  trace spans, and profiling of cloud native workloads.
- Using distributed trace tooling to observe a series of calls between
  microservices to understand where time is being spent.
- Managing the complexity, operational cost, and resource consumption of
  observability tools and systems at the enterprise scale.
- Best practices for meaningful alerting, queries, and operational dashboards
  including how to manage things including rules, definitions, thresholds and
  policies.
- How developers, operators, SRE, IT, and other actors comprehend, process, and
  reason on distributed cloud-native systems.
- Projects that incorporate novel & insightful approaches to utilizing
  observability data such as:
  - ML, model training, Bayesian networks, and other data science techniques
    that enable anomaly & intrusion detection.
  - correlating resource consumption with costing data to reduce the total cost
    of cloud native infrastructure
  - Using observability data exposed by service meshes, orchestrators, and other
    metric sources to inform continuous deployment tooling (e.g. Canary
    Predicates/Judges).
- Objective curation and generation of case studies pertaining to delivering
  observability tools/systems to end users.
- Best practices around observability and its continuous improvement, e.g. post
  mortems, runbooks
- Provide guidance around and foster interoperability between observability
  solutions without trying to enforce one specific standard
- Foster understanding of the prerequisites and corner-stones of observability
  like SLI/KPI, service objectives, and internal/external commitments.

The following is a non-exhaustive sample list of activities and deliverables
that are in-scope for this SIG

- Summary and overview of projects available in the community.
- Catalog of reference architectures that draw from CNCF projects, combining
  them in useful and novel ways.
- Definitions of implementations and patterns for best practices for
  delivering observability tooling at enterprise scale.
- Tooling composition and tool chain creation based on existing projects.
- Best practices for operations and monitoring workflows using CNCF Projects.
- Organizing and helping to provide visibility to Meetups, Blogs, and Podcasts
  related to the scope of the SIG.
- Guidance for application development and architecture that is observable.
- Replicatable reference architectures.
- Patterns for observing application delivery pipelines.
- Education regarding instrumentation cloud native workloads.
- Processing and Accessing relevant observability data at scale.
- Policy and security controls for observabilty data.
- Creating artifacts as part of CI/CD pipelines that facilitate observation of
  services.  Concrete examples might be:
  - service profiles for Linkerd
  - debug binaries or other diagnostic metadata.
  - representative trace spans from failing CI tests.

## Areas considered out of Scope

Anything not explicitly considered in the scope above.

Examples include:

- Datastores that are not primarily used for observability. Those datastores
  might be in the scope of SIG Storage.
- Security aspects that need to be present when setting up cloud native
  infrastructure, these might be more relevant for SIG Security.
- How cloud native applications that need observability are deployed; this would
  fall in the scope of SIG App Delivery
- Tools and projects that are used to run cloud native workloads that in some
  cases need observability would fall under the scope of SIG-Runtime.

## Roadmap & Initial Efforts

- Contribute to [due diligence reports][ddr] to assist the CNCF TOC for projects
  in the scope of the SIG.
- Facilitate webinars and presentations from CNCF projects and domain experts in
  the scope of the SIG.
- Formation of [SIG working group(s)][sigwg] as resource capacity and member
  contribution allows.

  > _SIGs may choose to spawn focussed and time-limited working groups to achieve some of their responsibilities (for example, to produce a specific educational white paper, or portfolio gap analysis report). Working groups should have a clearly documented charter, timeline (typically a few quarters at most), and set of deliverables. Once the timeline has elapsed, or the deliverables delivered, the working group dissolves, or is explicitly re-chartered._

[ddr]: https://github.com/cncf/toc/blob/master/process/due-diligence-guidelines.md
[sigwg]: https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#responsibilities--empowerment-of-sigs

## Governance

- This SIG follows the [standard operating model][som] provided by the TOC
  unless otherwise stated here.

[som]: https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#operating-model

## Operations

- Formation of the SIG follows the [documented process][sigform].
- [Roles][sigroles] for SIG Observability
  - TOC Liaison: *Jeff Brewer*\*
  - SIG Chairs: Matt Young, *Ricardo Aravena*\*
  - Tech Leads: Michael Hausenblas, Bartłomiej Płotka, *Richard Hartmann*\*

\*_**(TODO: need confirmation)**_

[sigform]: https://github.com/cncf/toc/tree/master/sigs#sig-formation-process
[sigroles]: https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#sig-member-roles

- Contact
  - Slack channel: #sig-observability @ [https://cloud-native.slack.com](https://cloud-native.slack.com)
  - Email List: [cncf-sig-observability@lists.cncf.io](mailto:cncf-sig-observability@lists.cncf.io)
- Meeting Schedule:
  - TBD - pending feedback from SIG members
  - [https://www.cncf.io/community/calendar](https://www.cncf.io/community/calendar/)
