# CNCF TAG Observability Charter

- [CNCF TAG Observability Charter](#cncf-tag-observability-charter)
  - [Introduction](#introduction)
  - [Mission](#mission)
  - [Areas considered in Scope](#areas-considered-in-scope)
  - [Areas considered out of Scope](#areas-considered-out-of-scope)
  - [Roadmap & Initial Efforts](#roadmap--initial-efforts)
  - [Governance](#governance)
  - [Operations](#operations)

Review and contributions from:
[Alex Nauda][Alex Nauda],
[Alois Reitbauer][Alois Reitbauer],
[Bartłomiej (Bartek) Płotka][Bartłomiej (Bartek) Płotka],
[Daniel Khan][Daniel Khan],
[Daniel Prata][Daniel Prata],
[Lincoln Sward][Lincoln Sward],
[Matthew Young][Matthew Young],
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
[Alex Nauda]:                 @
[Alois Reitbauer]:            https://github.com/aloisreitbauer
[Bartłomiej (Bartek) Płotka]: https://github.com/bwplotka
[Brendan Burns]:              https://github.com/brendandburns
[Daniel Khan]:                @
[Daniel Prata]:               @
[Lincoln Sward]:              @
[Matthew Young]:              https://github.com/halcyondude
[Matthias Loibl]:             https://github.com/metalmatze
[Michael Hausenblas]:         https://github.com/mhausenblas
[Ricardo Aravena]:            https://github.com/raravena80
[Richard Hartmann]:           https://github.com/RichiH
[Sergey Kanzhelev]:           @
[Steve Flanders]:             https://github.com/flands
[Ted Young]:                  @
[Tigran Najaryan]:            @
[Tommy Chong]:                https://github.com/techietommy
[Umair Ishaq]:                https://github.com/umairishaq

## Introduction

This document describes the purpose and operations of the Cloud Native
Computing Foundation ([CNCF]) Technical Advisory Group([TAG]) on Observability.

This [TAG] focuses on topics pertaining to the observation
of [cloud native][cn-def] workloads. Additionally, it produces supporting
material and best practices for end-users and provides guidance and
coordination for CNCF projects working within the TAG’s scope.

A full list of [CNCF projects][projs] can be found at [landscape.cncf.io].

[cncf]: https://www.cncf.io
[projs]: https://www.cncf.io/projects
[landscape.cncf.io]: https://landscape.cncf.io
[tag]:  https://github.com/cncf/toc/blob/main/tags/cncf-tags.md
[cn-def]: https://github.com/cncf/toc/blob/main/DEFINITION.md

## Mission

Consistent with the CNCF [TAG] definition, the mission of TAG Observability
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
  oversight of projects within the TAG's scope in an open, transparent, and
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
- How developers, operators, Site Reliability Engineers (SRE), IT Engineers, and
  other actors comprehend, process, and reason on distributed cloud-native
  systems.
- Projects that incorporate novel & insightful approaches to utilizing
  observability data such as:
  - ML, model training, Bayesian networks, and other data science techniques
    that enable anomaly & intrusion detection.
  - Correlating resource consumption with costing data to reduce the total cost
    of cloud native infrastructure
  - Using observability data exposed by service meshes, orchestrators, and other
    metric sources to inform continuous deployment tooling (e.g. Canary
    Predicates/Judges).
- Objective curation and generation of case studies pertaining to delivering
  observability tools/systems to end users.
- Best practices around observability and its continuous improvement, e.g. post
  mortems, runbooks
- Provide guidance around and foster interoperability between observability
  solutions without trying to enforce one specific standard.
- Foster understanding of the prerequisites and corner-stones of observability
  like SLI/KPI, service objectives, and internal/external commitments.

The following is a non-exhaustive sample list of activities and deliverables
that are in-scope for this TAG

- Summary and overview of projects available in the community.
- Catalog of reference architectures that draw from CNCF projects, combining
  them in useful and novel ways.
- Definitions of implementations and patterns for best practices for
  delivering observability tooling at enterprise scale.
- Tooling composition and toolchain creation based on existing projects.
- Best practices for operations and monitoring workflows using CNCF Projects.
- Organizing and helping to provide visibility to Meetups, Blogs, and Podcasts
  related to the scope of the TAG.
- Guidance for application development and architecture that is observable.
- Replicable reference architectures.
- Patterns for observing application delivery pipelines.
- Education regarding instrumentation cloud native workloads.
- Processing and Accessing relevant observability data at scale.
- Policy and security controls for observability data.
- Creating artifacts as part of CI/CD pipelines that facilitate observation of
  services. Concrete examples might be:
  - Service profiles for Linkerd.
  - Debug binaries or other diagnostic metadata.
  - representative trace spans from failing CI tests.

## Areas considered out of Scope

Anything not explicitly considered in the scope above.

Examples include:

- Datastores that are not primarily used for observability. Those datastores
  might be in the scope of TAG Storage.
- Security aspects that need to be present when setting up cloud native
  infrastructure, these might be more relevant for TAG Security.
- How cloud native applications that need observability are deployed; this would
  fall in the scope of TAG App Delivery
- Tools and projects that are used to run cloud native workloads that in some
  cases need observability would fall under the scope of TAG-Runtime.

## Roadmap & Initial Efforts

- Contribute to Due Diligence (DD) Reviews in accordance with the TOC's
  [Due Diligence Guidelines][ddg] and [Due Diligence Review Template][ddrt] to
  assist the CNCF TOC for projects in the scope of the SIG.
- Facilitate webinars and presentations from CNCF projects and domain experts in
  the scope of the TAG.
- Formation of [TAG working group(s)][tagwg] as resource capacity and member
  contribution allows.

  > _TAGs may choose to spawn focussed and time-limited working groups to achieve some of their responsibilities (for example, to produce a specific educational white paper, or portfolio gap analysis report). Working groups should have a clearly documented charter, timeline (typically a few quarters at most), and set of deliverables. Once the timeline has elapsed, or the deliverables delivered, the working group dissolves, or is explicitly re-chartered._

[ddg]: https://github.com/cncf/toc/blob/master/process/due-diligence-guidelines.md
[ddrt]: https://github.com/cncf/toc/blob/master/process/dd-review-template.md
[tagwg]: https://github.com/cncf/toc/blob/master/tags/cncf-tags.md#responsibilities--empowerment-of-tags

## Governance

- This TAG follows the [standard operating model][som] provided by the TOC
  unless otherwise stated here.

[som]: https://github.com/cncf/toc/blob/main/tags/cncf-tags.md#operating-model

## Operations

- Formation of the TAG follows the [documented process][tagform].
- [Roles][tagroles] for TAG Observability
  - TOC Liaison: Cathy Zhang, Ricardo Rocha
  - TAG Chairs:
    - [Alolita Sharma] [alolita]
    - [Matthew Young][matthew young]
    - [Richard Hartmann][Richard Hartmann]
  - Tech Leads ([definition][tagtldefinition], [TL election][tagtlprocess])
    - Bartłomiej Płotka
    - TBD more

[tagform]: https://github.com/cncf/toc/tree/main/tags#tag-formation-process
[tagroles]: https://github.com/cncf/toc/blob/main/tags/cncf-tags.md#tag-member-roles
[tagtlprocess]: https://github.com/cncf/toc/blob/master/tags/cncf-tags.md#elections
[tagtldefinition]: https://github.com/cncf/toc/blob/master/tags/cncf-tags.md#tech-lead

- Contact
  - Slack channel: #tag-observability @ [https://cloud-native.slack.com](https://cloud-native.slack.com)
  - Email List: [cncf-tag-observability@lists.cncf.io](mailto:cncf-tag-observability@lists.cncf.io)
- Meeting Schedule:
  - Every 2nd Tuesday, 4pm - 5pm, UTC ([CNCF Calendar][cncfcal] for details)

[cncfcal]: https://www.cncf.io/calendar
