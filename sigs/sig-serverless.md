## CNCF Serverless SIG Charter

## Introduction

This is the charter referred to in “[CNCF
SIGs](https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#sig-charter)”
by the CNCF TOC, and consistent with the [proposed SIG
definition](https://github.com/cncf/toc/blob/master/sigs/proposed.md).

Borrowing from the [Serverless
Whitepaper](https://github.com/cncf/wg-serverless/tree/master/whitepapers/serverless-overview),
Serverless computing refers to the concept of building and running
applications that do not require server management. Allowing developers
to focus on writing their applications’ business logic rather than the
complexities of the platform hosting/managing their applications. As such,
this SIG will focus on technologies related to enabling this simplified
developer experience.

## Areas Considered In Scope

Technologies related to hosting, managing, developing and integrating of
Serverless applications from an application developer's perspective.
The following are example areas that are considered in scope:

* container hosting platforms that attempt to abstract their complexities away
  from the application developer
* orchestration and management of workloads running on those platforms from the
  application developer's perspective
* tooling to aid in the development and deployment of those workloads
* integration technologies that aim to promote interoperability between the
  various systems that an application developer may interact with during the
  management of their applications on those platforms
* workload interoperability - e.g. portability of functions,
  interaction with hosting environment/runtime
* interoperability and connectivity with backend services used by workloads

## Areas Considered Out Of Scope

Anything not considered in scope above is out of scope.  See also “Interfaces
with Related Groups” below.

An example would include technology related to the inner-workings of a
particular platform. For example, a new autoscaler for Kubernetes.

## SIG Mission Statement

To enable application developers to more easily create, deploy and manage their
applications without the need to become experts in the hosting technologies
being used.

Developers should focus their attention on the development of their
applications. When they are required to spend a significant portion of their
time learning and interacting with the hosting platform of those applications
then that is time that should have been spent working on their business
objective.

To that end, the SIG will:

1. Provide valuable and unbiased information to the TOC,
   End Users and Projects of the CNCF regarding areas considered in scope
   (see above).
2. Collaborate effectively with other related groups (see below).
3. Help to maintain the continued health of the CNCF Projects deemed
   to be within the scope of this SIG (see below)
4. Identify and fill gaps in the landscape of CNCF Projects within scope
   either through encouragement of existing projects to join the CNCF
   or through development of new artifacts (e.g. specs, code, papers).
5. Maintain any "living" documentation, such as landscape and state of the
   community documents.

Specific SIG deliverables are as per the above, and the [general SIG
responsibilities set out by the CNCF
 TOC](https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#responsibilities--empowerment-of-sigs).

## Current CNCF Projects Considered To Be Within The Scope Of This SIG

* CloudEvents
* Serverless Workflow
* CloudEvents Discovery API
* CloudEvents Subscription API

## Interfaces With Other Related Groups

It is not always clear where the line is between
Serverless and other cloud native hosting technologies, therefore the exact
choice of a project being within scope of SIG-Serverless versus one of the
other CNCF SIGs will need to be consider on a case by case basis. Often the
decision might be based on the project's relationship to existing projects
within the SIG.

*   **[Several Kubernetes SIGs](https://github.com/kubernetes/community)**
    cover Kubernetes-specific workload, scheduling, autoscaling, execution
    and other related abstractions, interfaces, and implementations of
    these interfaces.  We will maintain communication with these Kubernetes
    SIGs where needed.  Our aim is to avoid unnecessary duplication of
    effort by the two groups and maintain clear and consistent messaging
    to our end user community and projects.
*   **[CNCF Security SIG](https://github.com/cncf/sig-security)**
    works on the more general area of cloud-native security including
    authentication, authorization, encryption, accounting, auditing, and
    related topics.  We defer as much as possible to this group to deal
    with general security-related issues and liaise closely with them on
    how to deal with security areas where these arise.
*   **[CNCF App Delivery SIG](https://github.com/cncf/sig-app-delivery)**
    is focussed on the development, deployment, operation and testing of
    cloud-native applications. We collaborate with this SIG where it
    pertains to helping to ensure that the required underlying workload
    execution abstractions and mechanisms are suitably provided to support
    these application-level delivery needs.
*   **[CNCF Runtime SIG](https://github.com/cncf/sig-runtime)**
    includes projects that are focused on the runtime used to host
    cloud-native applications. While we expect a lot of collaboration with this
    SIG, the difference can be thought of as SIG-Runtime will be more
    concerned with the specifics of how runtimes host the application, while
    SIG-Serverless will be more about how developers interact with the
    runtime. And in particular, in a way that attempts to abstract the
    complexities of that runtime.

## Operating Model

This SIG follows the
[standard operating
 guidelines](https://github.com/cncf/toc/blob/master/sigs/cncf-sigs.md#operating-model)
provided by the TOC unless otherwise stated here.

**Current TOC Liaison:**  ????

**Co-Chairs:** [Doug Davis](https://github.com/duglin/),
  [Mark Peek](https://github.com/markpeek/),
  [Ken Owens](https://github.com/kenowens12)

**Tech Leads:**
  - Workflow: See project
  [owners](https://github.com/cncf/wg-serverless/blob/master/workflow/spec/governance/owners.md)
  - CloudEvents: See project
  [governance](https://github.com/cloudevents/spec/blob/master/GOVERNANCE.md)
  - Discovery & Subscription: See project
  [governance](https://github.com/cloudevents/spec/blob/master/GOVERNANCE.md)

**Other named roles:** None at present; will be identified as needed

## Meeting Schedule

The CNCF SIG-Serverless group meets every Thursday at 12pm Eastern.

Zoom:
[https://zoom.us/my/cncfsigserverless](https://zoom.us/my/cncfsigserverless)

Mailing list: Join SIG-Serverless mailing list at
[lists.cncf.io](https://lists.cncf.io)

Slack channel: [TBD](TBD) 
