# [Project] Graduation Proposal

CloudEvents is a specification for describing common event data in a
consistent (interoperable) way. CloudEvents seeks to dramatically simplify
event declaration and delivery across services and platforms. The project intentionally focuses on events and does not try to generically cover various message types and patterns.

Previous efforts around "event interoperability" focused on projects that
tried to create something new (such as a common event format) where its
success was dependent on everyone changing their existing events to conform
to the new standard. A nearly impossible task.

CloudEvents took a different approach. Instead, CloudEvents leverages the
existing events and messaging transports in order to achieve the
interoperability it's aiming for. By allowing existing events to remain
"as-is" and leveraging the native extensibility model of the transport,
CloudEvents allows for common eventing metadata to be added without major
disruptions to existing message flows and infrastructure.

With that core model, CloudEvents then defined the key metadata that most
events already have (e.g. source, id, event-type) and by pulling them out of
the event's "business data", it allows for middleware to easily route the
events to their final destinations without the need to have a priori
knowledge of each event's schema or to even examine the business data at all.

Since reaching its v1.0 status we've seen a constant increase in adoption
across a wide range of industries. Including both infrastructure providers
as well as end-users. While the initial motivation was the role of events in serverless technologies, the adoption of CloudEvents today exceeds this context. 

It's been 3 years since our v1.0 release and given the
positive feedback we've received and the lack of major issues found with the
specification, we feel that it is time to move from "incubator" to
"graduated" status.

We have a list of companies who are using CloudEvents today and are willing
to discuss their usage, some of which are listed on the "adopters" section
on main [website](https://cloudevents.io/).

## Graduation State Criteria

### * Have committers from at least two organizations.

CloudEvents, being primarily a specification project, operates a little
differently from a code-based OSS project. Our working process is very
collaborative and pull-requests for spec changes are often developed by a
number of people through many long brainstorming sessions, or our weekly
meetings. As such, committers/maintainers of the project are not determined
by who authored a certain number of merged PRs. Instead, they are determined by level of participation - meaning, the company attends 3 out of 4 of the last weekly meetings. This shows the appropriate level of commitment to the project that is rewarded with the right to "vote", if we ever do have a formal vote, which is rare.

Voting is rare because of the process we've define for approving changes to
the specifications. Our process for reviewing and approving PRs is open to
everyone regardless of their status in the project. We have standing rule
that no PR will be merged until all outstanding comments are resolved (which
almost gives everyone 'veto power') and then PR is approved during one of our
weekly meetings. This ensures a collaborative, and community-wide, review and
approval process. As a result nearly all of our PRs have been developed,
reviewed and approved relatively easily. There have only been a small handful
of PRs that ended up needing a formal "vote" to decide which path the group
wanted to go. As a result, having "voting rights", has limited meaning in our
project.

For more information on our processes please see our
[Governance
doc](https://github.com/cloudevents/spec/blob/main/docs/GOVERNANCE.md).

With that explanation, the current list of participants can be found in
this attendance tracking
[spreadsheet](https://docs.google.com/spreadsheets/d/1bw5s9sC2ggYyAiGJHEk7xm-q2KG6jyrfBy69ifkdmt0/edit?pli=1#gid=0).
A green box in the "Voting Rights" column shows which companies are "voting
members". At the time of this document being created we have 10 such members.
In addition, a list of contributors can be found in [contributors.md](https://github.com/cloudevents/spec/blob/main/docs/contributors.md).

### * Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

Yes, see our status
[here](https://bestpractices.coreinfrastructure.org/projects/6770).

### * Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

Yes, see our Trail of Bits audit
[here](https://github.com/cloudevents/spec/blob/main/docs/CE-SecurityAudit-2022-10.pdf).
No critical issues were found and all non-critical issues found have been
resolved.

### * Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

Yes, see our
[Governance](https://github.com/cloudevents/spec/blob/main/docs/GOVERNANCE.md)
and
[Contributing](https://github.com/cloudevents/spec/blob/main/docs/CONTRIBUTING.md)
docs.

### * Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

Yes, see first section of this doc and our
[Governance](https://github.com/cloudevents/spec/blob/main/docs/GOVERNANCE.md)
and
[Contributing](https://github.com/cloudevents/spec/blob/main/docs/CONTRIBUTING.md)
docs.

### * Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to [FAQs](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for guidelines on identifying adopters.

Yes, see the "adopters" section on our main
[website](https://cloudevents.io).

Also, see the list of known OSS projects using CloudEvents
[here](https://github.com/cloudevents/spec/blob/main/docs/open-source.md).

## Incubation Details
_**Project should address each area listed below**_

### * Link to Incubation Due Diligence(DD) Document

See [CloudEvents Due Diligence](cloudevents-dd.md).

### * Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

