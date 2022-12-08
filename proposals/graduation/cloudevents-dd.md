# CloudEvents Due Diligence

## Statement of CNCF Alignment to TOC Principles
1. Project is self-governing

Yes

2. Is there a documented Code of Conduct that adheres to the CNCF guidelines?
Yes. The governance docs point to:  https://github.com/cncf/foundation/blob/master/code-of-conduct.md

3. Does the project have production deployments that are high quality and high-velocity? (for incubation and graduated projects).

Yes. While CloudEvents itself, at its core, is a specification without an official implementation we have a list of known adopters of the specification on our https://cloudevents.io website and have gathered a list of companies who would be willing to discuss their usage of CloudEvents with the TOC members as part of the application for graduated status.

While not an exhaustive list, vendors such as [Google](https://cloud.google.com/functions/docs/writing/write-event-driven-functions) and [Microsoft](https://learn.microsoft.com/en-us/azure/event-grid/cloudevents-schema) have public documentation on using CloudEvents. And numerous open source projects have support for CloudEvents beyond what is listed in the [open source](https://github.com/cloudevents/spec/blob/main/docs/open-source.md) list. As examples, [dapr](https://docs.dapr.io/developing-applications/building-blocks/pubsub/pubsub-overview/),
[knative](https://knative.dev/docs/eventing/), [opentelemetry](https://opentelemetry.io/docs/reference/specification/trace/semantic_conventions/cloudevents/),
[spring](https://spring.io/blog/2020/12/10/cloud-events-and-spring-part-1).

Additionally, the CloudEvents github repo hosts community driven SDKs that support a wide range of languages (CSharp, Go, Java, Javascript, PHP, PowerShell, Python, Ruby and Rust) as well as many of the CloudEvent protocol bindings.

4. Is the project committed to achieving the CNCF principles and do they have a committed roadmap to address any areas of concern raised by the community?

Yes. Concerns raised by the community will be managed through github
issues: https://github.com/cloudevents/spec/issues. See the
[Contributing doc](https://github.com/cloudevents/spec/blob/main/docs/CONTRIBUTING.md)
for more information.

5. Document that the project has a fundamentally sound design without obvious critical compromises that will inhibit potential widespread adoption

List of adopters can be found on the "adopters" section of the
[website](https://cloudevents.io) as well as the list of OSS projects
that are using it
[here](https://github.com/cloudevents/spec/blob/main/docs/open-source.md).

6. Document that the project is useful for cloud native deployments & degree that it's architected in a cloud native style

See previous bullet.

7. Document that the project has an affinity for how CNCF operates and understand the expectation of being a CNCF project.

CloudEvents has been a CNCF project since 2018 and as such has a history of
being a good, and successful, member of the CNCF community.

### Incubating Stage Graduation (Exit Requirements)
1. Document that it is being used successfully in production by at least three independent direct adopters which with focus on adequate quality and scope defined.

Some adopters include: Adobe, Alibaba, European Commission, Google, IBM, Microsoft, Tencent, TriggerMesh and VMware. See the full list of adopters on the "adopters" section of the
[website](https://cloudevents.io) as well as the list of OSS projects

2. Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

CloudEvents, being a specification project, doesn't have formal committers.
Rather each change to the specifications requires approval from the entire
community. All comments on a PR must be resolved to the community's
satisfaction before it is merged. This, in essence, gives everyone veto power.

In the rare cases the community can not agree then a vote is taken by
the members of the community who have been to 3 out of the last 4 weekly
calls. This has only happened less than 5 times during the entire life of the
project. However, to gauge interest from the community, please see the
historical list of weekly call attendees
[here](https://docs.google.com/spreadsheets/d/1bw5s9sC2ggYyAiGJHEk7xm-q2KG6jyrfBy69ifkdmt0/edit?pli=1#gid=0)
and the "voting rights" column shows about 10 different companies regularly
attend the calls and have voting rights.

3. Demonstrate a substantial ongoing flow of commits and merged contributions.

See activity over the last 2 years
[here](https://cloudevents.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&var-period=w&var-repogroups=%22cloudevents%2Fspec%22&from=1606798800000&to=1672549199000)

4. Have committers from at least two organizations.

See item 2.

5. Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

Yes, see: https://bestpractices.coreinfrastructure.org/projects/6770

6. Adopted the CNCF Code of Conduct.

Yes, see: https://github.com/cloudevents/spec#community-and-docs

7. Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

See section 2 above, and
[here](https://github.com/cloudevents/spec/blob/main/docs/CONTRIBUTING.md).

8. Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

See section 2 above, and
[here](https://github.com/cloudevents/spec/blob/main/docs/CONTRIBUTING.md).

9. Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website), guidelines for identifying adopters can be found in the [FAQs](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter).

See list of adopters on the "adopters" section of the
[website](https://cloudevents.io) as well as the list of OSS projects

### Documentation of CNCF Alignment (if not addressed above):
* name of project (must be unique within CNCF)

CloudEvents

* project description (what it does, why it is valuable, origin and history)

CloudEvents is a specification for describing common event data in popular
formats and transports to provide interoperability across services, platforms
and systems.

* statement on alignment with CNCF charter mission

See intro here: https://github.com/cloudevents/spec#readme

* sponsor from TOC (sponsor helps mentor projects)

Originally it was:
- Ken Owens <ken.owens @ mastercard.com>
- Brian Grant <briangrant @ google.com>

* license (charter dictates Apache 2 by default)

Yes, see https://github.com/cloudevents/spec/blob/main/LICENSE

* source control (GitHub by default)

Yes, see https://github.com/cloudevents/spec

* external dependencies (including licenses)

None

* release methodology and mechanics

See https://github.com/cloudevents/spec/blob/main/docs/GOVERNANCE.md#release-process-and-versioning

* community size and any existing sponsorship

See weekly call attendee list: https://docs.google.com/spreadsheets/d/1bw5s9sC2ggYyAiGJHEk7xm-q2KG6jyrfBy69ifkdmt0/edit?pli=1#gid=0
They regularly have about 10 different companies represented on each weekly
call.

## Technical
* An architectural, design and feature overview should be available. (add link)

See the
[Primer](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/primer.md).

* What are the primary target cloud-native use cases? Which of those:
  * Can be accomplished now.
  * Can be accomplished with reasonable additional effort (and are ideally already on the project roadmap).
  * Are in-scope but beyond the current roadmap.
  * Are out of scope.

See https://github.com/cloudevents/spec/blob/main/docs/ROADMAP.md for the
historical roadmap. Right now CloudEvents is at version 1.0 and they are
addressing any questions from the community as they arise. The project is
now focusing on other eventing lifecycle pain-points which are being
addressed by the creation of a Discovery specification, a Subscription
Manager specification and a Schema Registry specification. These new specs
are still under active development.

* What are the current performance, scalability and resource consumption bounds of the software? Have these been explicitly tested? Are they appropriate given the intended usage (e.g. agent-per-node or agent-per-container need to be lightweight, etc)?

Not applicable.

* What exactly are the failure modes? Are they well understood? Have they been tested? Do they form part of continuous integration testing? Are they appropriate given the intended usage (e.g. cluster-wide shared services need to fail gracefully etc)?

Not applicable.

* What trade-offs have been made regarding performance, scalability, complexity, reliability, security etc? Are these trade-offs explicit or implicit? Why? Are they appropriate given the intended usage? Are they user-tunable?

Not applicable.

* What are the most important holes? No HA? No flow control? Inadequate integration points?

Not applicable.

* Code quality. Does it look good, bad or mediocre to you (based on a spot review). How thorough are the code reviews? Substance over form. Are there explicit coding guidelines for the project?

Not applicable for the spec itself, but we do host some community driven
SDKs that are at varying degrees of maturity. The most popular ones
(e.g. golang) are being used in other OSS projects (such as Knative).

* Dependencies. What external dependencies exist, do they seem justified?

None

* What is the release model? Versioning scheme? Evidence of stability or otherwise of past stable released versions?

See https://github.com/cloudevents/spec/blob/main/docs/GOVERNANCE.md#release-process-and-versioning
and https://github.com/cloudevents/spec/blob/main/docs/RELEASES.md

* What is the CI/CD status? Do explicit code coverage metrics exist? If not, what is the subjective adequacy of automated testing? Do different levels of tests exist (e.g. unit, integration, interface, end-to-end), or is there only partial coverage in this regard? Why?

Not applicable.

* What licensing restrictions apply? Again, CNCF staff will handle the full legal due diligence.

Not applicable.

* What are the recommended operational models? Specifically, how is it operated in a cloud-native environment, such as on Kubernetes?

Not applicable.

## Project
* Do we believe this is a growing, thriving project with committed contributors?
* Is it aligned with CNCF's values and mission?
* Do we believe it could eventually meet the graduation criteria?
* Should it start at the sandbox level or incubation level?
* Does the project have a sound, documented process for source control, issue tracking, release management etc.
* Does it have a documented process for adding committers?
* Does it have a documented governance model of any kind?
* Does it have committers from multiple organizations?
* Does it have a code of conduct?
* Does it have a license? Which one? Does it have a CLA or DCO? Are the licenses of its dependencies compatible with their usage and CNCF policies? CNCF staff will handle the full legal due diligence.
* What is the general quality of informal communication around the project (slack, github issues, PR reviews, technical blog posts, etc)?
* How much time does the core team commit to the project?
* How big is the team? Who funds them? Why? How much? For how long?
* Who are the clear leaders? Are there any areas lacking clear leadership? Testing? Release? Documentation? These roles sometimes go unfilled.
* Besides the core team, how active is the surrounding community? Bug reports? Assistance to newcomers? Blog posts etc.
* Do they make it easy to contribute to the project? If not, what are the main obstacles?
* Are there any especially difficult personalities to deal with? How is this done? Is it a problem?
* What is the rate of ongoing contributions to the project (typically in the form of merged commits).

## Users
* Who uses the project? Get a few in-depth references from 2-4 of them who actually know and understand it.
* What do real users consider to be its strengths and weaknesses? Any concrete examples of these?
* Perception vs Reality: Is there lots of buzz, but the software is flaky/untested/unused? Does it have a bad reputation for some flaw that has already been addressed?

## Context
* What is the origin and history of the project?

See https://github.com/cncf/toc/blob/main/proposals/sandbox/cloudevents.md

* Where does it fit in the market and technical ecosystem?
* Is it growing or shrinking in that space? Is that space growing or shrinking?
* How necessary is it? What do people who don't use this project do? Why exactly is that not adequate, and in what situations?
* Clearly compare and contrast with peers in this space. A summary matrix often helps. Beware of comparisons that are too superficial to be useful, or might have been manipulated so as to favor some projects over others. Most balanced comparisons will include both strengths and weaknesses, require significant detailed research, and usually there is no hands-down winner. Be suspicious if there appears to be one.

