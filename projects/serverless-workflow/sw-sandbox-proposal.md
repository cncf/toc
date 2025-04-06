# Serverless Workflow

**Name of project**: Serverless Workflow

**Description**:

TOC created the Serverless Working Group to investigate
the Serverless landscape. The outputs of the WG included:

- a [whitepaper](https://github.com/cncf/wg-serverless/blob/master/whitepapers/serverless-overview/README.md) that:
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

The Working Group's first focus area was around making it easier for middleware to process Events,
which led to the development of the [CloudEvents](https://cloudevents.io) specification.

Additionally, the Serverless WG has agreed to form a sub working group to work on
the Serverless Workflow specification.

Serverless Workflow is a vendor-neutral specification for defining the model of workflows responsible for 
orchestrating event-driven serverless applications.

Goals of the Serverless Workflow sub-group are to :

- Standardize Serverless Workflow model and definition
- Facilitate Serverless Workflow portability
- Be completely vendor neutral
- Support both stateless and stateful Serverless Workflow implementations
- To define a light-weight and powerful Serverless Workflow model

Functional scope of the Serverless Workflow specification includes:

- Define and orchestrate steps/states involved in a serverless application.
- Define which functions are executed in each step.
- Define which event or combination of events trigger function execution.
- Define function execution behavior (sequential, parallel, etc).
- Specify information filtering throughout the execution of the serverless workflow.
- Define error conditions with retries.
- Define correlation between events and workflow instances.
- Specify workflow integration with third-party services.

Here is diagram that illustrates the Serverless Workflow functional flow:

<p align="center">
<img src="https://raw.githubusercontent.com/cncf/wg-serverless/master/workflow/spec/media/spec/sample-serverless-workflow2.png" height="450px" alt="Serverless Workflow Diagram"/>
</p>

Below are some of the core Serverless Workflow Specification documents:

* [Examples](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/examples.md)
* [Use Cases](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/usecases.md)
* [Community Contributions](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/community/open-source.md)
* [Logo](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/media/logo-small-text.png)

The work on this specification is currently being executed within the
CNCF's Serverless WG Workflow sub-group, but with the large growth of the specification,
and the short-term plan for a release of our first milestone (v0.1), it would make
sense for this work to be more formalized as a new sandbox project under the CNCF.

The goals/roadmap of the project include:

- finalize release v0.1.
- define requirements for implementation compliance.
- define requirements for Serverless Workflow API/SDK.
- continue improving and verifying specification quality by
promoting and testing community implementations.

You can find the project roadmap document [here](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/roadmap.md).

You can find the Serverless Workflow blog [here](https://serverlessworkflow.blogspot.com/).

**Statement on alignment with CNCF mission**:

Being born out of the CNCF's Serverless Working Group, the Serverless Workflow
project (and its members) share the CNCF's goals of promoting Cloud Native
technologies, and offering choice to our consumers through an open
interoperability specification, as shown by the significant participation
from key industry companies.

We believe that the CNCF provides the proper home for this due to its
commitment to the promotion and development of open, vendor-neutral projects.
Additionally, the wide breadth of the CNCF members will provide the feedback
necessary to ensure the Serverless Workflow specification isn't too limited in its
scope and appeals to as many constituents of the cloud native community
as possible.

**Sponsor / Advisor from TOC**:

**Preferred maturity level**:

Sandbox

**License**:

Apache License v2.0

**Source control repositories**:

CNCF Serverless WG repository (workflow directory): https://github.com/cncf/wg-serverless/tree/master/workflow

**External dependencies**:

None

**Initial Maintainers**:

The [GOVERNANCE](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/governance/readme.md)
documentation explains the process used within the Serverless Workflow group.

The [OWNERS](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/governance/owners.md)
file lists community members currently responsible for oversight, review, and approval of changes to specification 
documents.

**Infrastructure Requests**:

The Serverless Workflow Group requests a dedicated github repository within the https://github.com/CNCF namespace for the Serverless Workflow specification.
The Serverless Workflow Group requests a dedicated website.
The Serverless Workflow Group requests a dedicated mailing list.

**Communication Channels**:

Mailing list: https://lists.cncf.io/g/cncf-wg-serverless

Slack: #serverless-workflow Slack channel within CNCF's Slack workspace https://slack.cncf.io/

We have monthly zoom calls, first Monday every month @ 10-11am PT:
https://zoom.us/my/cncfserverlesswg

We track our monthly meeting information here: https://docs.google.com/document/d/1xwcsWQmMiRN24a7o7oy9MstzMroAup31oOkM5Dru1jQ

**Issue tracker**:

Issues are tracked with GitHub Issues: https://github.com/cncf/wg-serverless/issues

Changes are tracked with GitHub PRs: https://github.com/cncf/wg-serverless/pulls

**Website**:

Serverless Workflow does not have its own website.

**Release Methodology and Mechanics**

With the large growth of the specification in the last months, we are formalizing the
specification roadmap at the moment and working on the v.0.1 release.

Beyond what will be defined there, the sub-group will decide when
significate progress has been made to warrant a new release.

**Social Media Accounts**:

**Contributor statistics**:

Meeting attendance is tracked [here](https://docs.google.com/document/d/1xwcsWQmMiRN24a7o7oy9MstzMroAup31oOkM5Dru1jQ/edit#heading=h.g2rizfze8av2).

As can been seen in that document, Serverless Workflow monthly calls have regular
attendance.

Without implying endorsement, the following companies have attended at least
one meeting:

Nokia, Red Hat, Camunda, Futurewei, Pivotal / VMware, IBM, Twilio, Codid

And the following have voting rights (today):

Red Hat, Camunda, Futurewei

We continue to keep working on increasing our community interested in this specification.
