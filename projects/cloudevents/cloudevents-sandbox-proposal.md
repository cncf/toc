# CloudEvents

**Name of project**: CloudEvents

**Description**:

Last year the CNCF TOC created the Serverless Working Group to investigate
the Serverless landscape. The outputs of the WG included:
- a [whitepaper](https://github.com/cncf/wg-serverless#whitepaper) that:
  - defines Serverless and its terminology
  - describes common use cases for the technology
  - compares it with other Cloud Native technologies and \*aaS environments
  - describes the common architecture of Serverless platforms
- a [landscape document](https://docs.google.com/spreadsheets/d/10rSQ8rMhYDgf_ib3n6kfzwEuoE88qr0amUPRxKbwVCk/edit#gid=0)
  that lists well-known open-source and proprietary Serverless platforms
  and tools
- a set of recommended next steps for the WG, as part of the whitepaper:
  - encourage more Serverless technology vendors and open source developers
    to join the CNCF
  - foster an open ecosystem by establishing interoperable APIs, in particular
    around: Events, Deployments and Workflows
  - provide additional education as needed

One of the recommendations, interoperability around Events, was agreed to
by the TOC and the WG began to develop a new specification for how
Events that are transferred between an event producer and an event consumer
should be formalized. The purpose of this would be to better enable
interoperability between these components such that basic processing of
the events (such as routing) can be achieved without having to require
knowledge of the event's structure in advance, or without understanding
the application specific data of the event.

The work on this specification is currently being done within the
CNCF's Serverless Working Group, but with the release of our first
milestone (v0.1), it would make sense for this work to be more
formalized as a new sandbox project under the CNCF.

The goals/roadmap of the project include:
- moving the specification to v1.0. A baseline format for an Event
  to enable broad adoption within the Cloud community, and in particular
  for Serverless/FaaS implementations
- define protocol mappings for popular transports, such as HTTP
- define serialization mappings for popular formats, such as JSON

**Statement on alignment with CNCF mission**:

Being born out of the CNCF's Serverless Working Group, the CloudEvents
project (and its members) share the CNCF's goals of promoting Cloud Native
technologies, and offering choice to our consumers through an open
interoperability specification, as shown by the significant participation
from key industry companies.

We believe that the CNCF provides the proper home for this due to its
commitment to the promotion and development of open, vendor-neutral projects.
Additionally, the wide breadth of the CNCF members will provide the feedback
necessary to ensure the CloudEvents specification isn't too limited in its
scope and appeals to as many constituents of the cloud native community
as possible.

**Sponsor / Advisor from TOC**:
- Ken Owens <ken.owens @ mastercard.com>
- Brian Grant <briangrant @ google.com>

**Preferred maturity level**: Sandbox

**License**: Apache License v2.0

**Source control repositories**: 

CloudEvents org: https://github.com/cloudevents

CloudEvents repo for the specification: https://github.com/cloudevents/spec

**External dependencies**: None

**Initial Maintainers**:

The CloudEvents group does not have "maintainers" that approve
Pull Requests (PRs) like traditional GitHub projects. Rather, the group
discusses/reviews PRs in the PRs themselves and then when consensus is reached
they are approved during our weekly calls. If consensus can not be reached
then a formal vote is taken. 

Voting rights: each member company designates a "primary" and "alternate"
member whose attendance at the weekly calls is tracked. Any member company
that attends three out of the last four meetings (current meeting not included)
has voting rights.

We also have this
[GOVERNANCE](https://github.com/cloudevents/spec/blob/master/GOVERNANCE.md)
doc which explains the processes we follow.

**Infrastructure Requests**: None

**Communication Channels**:

Mailing list: CloudEvents uses the CNCF Serverless WG mailing list:
https://groups.google.com/forum/#!forum/cncf-wg-serverless but we may
move to our own dedicated mailing list when/if the Serverless WG starts
a second project.

Slack: There is a #cloudevents Slack channel under CNCF's Slack workspace.

We have weekly zoom calls (9am PT on Thursdays):
https://zoom.us/my/cncfserverlesswg

**Issue tracker**:

Issues are tracked with GitHub Issues:https://github.com/cloudevents/spec/issues

Changes are tracked with GitHub PRs: https://github.com/cloudevents/spec/pulls

**Website**:

CloudEvents has its own website at: https://cloudevents.io

**Release Methodology and Mechanics**

CloudEvents has a set of milestones defined in its
[roadmap](https://github.com/cloudevents/spec/blob/master/roadmap.md)
document. Beyond what is defined there, the group will decide when
significate progress has been made to warrant a new release.

**Social Media Accounts**:

Twitter: @CloudEventsDemo 

**Contributor statistics**:

Attendance is tracked [here](https://docs.google.com/spreadsheets/d/1bw5s9sC2ggYyAiGJHEk7xm-q2KG6jyrfBy69ifkdmt0/edit?pli=1#gid=0).
As can been seen in that document, CloudEvents weekly calls have regular
attendance from most major cloud vendors, averaging nearly 30 people
each week.

Without implying endorsement, the following companies have attended at least
one meeting:
Accenture, Alibaba, Amazon, Bitnami/Kubeless, Cisco, Clay, CNCF,
Collinson Group, Cuemby, Google, Hauwei, Honeycomb.io, Huawei, IBM, iguazio,
infraCloud, Intel, JP Morgan, JS Foundation, Mastercard, Microsoft, NAIC,
Nordstrom, OpenFaaS, Oracle, Particular Software, Pivotal, Progress, Red Hat,
RX-M, SAP, Serverless, Singlepoint, Solar Winds, solo.io, Splunk, VMWare

And the following have voting rights (today), which means they regularly
attend the weekly calls:
Alibaba, CNCF, Google, Huawei, IBM, iguazio, Intel, JS Foundation, Microsoft,
NAIC, Nordstrom, Oracle, Red Hat, SAP, Serverless, VMWare

In terms of adoption, the following companies participated in the KubeCon
EU CloudEvents demo:
Alibaba, Google, Hauwei, IBM, iguazio, Microsoft, Oracle, Red Hat, SAP,
Serverless, VMWare

Azure recently announced official support for CloudEvents in their
[Event Grid](https://docs.microsoft.com/en-us/azure/event-grid/cloudevents-schema),
and Serverless announced support for it in their
[Event Gateway](https://serverless.com/learn/event-gateway/).
