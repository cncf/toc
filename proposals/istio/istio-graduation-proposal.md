[Istio](https://istio.io) is an open source service mesh that transparently provides a uniform and efficient way to secure, connect, and monitor services in cloud native applications. It provides zero-trust networking, policy enforcement, traffic management, load balancing, and monitoring, without requiring applications to be rewritten.

Istio [applied for CNCF transfer](https://github.com/cncf/toc/pull/827) in April 2022 and [was accepted in September 2022](https://www.cncf.io/blog/2022/09/28/istio-sails-into-the-cloud-native-computing-foundation/).  The [proposal](https://github.com/cncf/toc/blob/main/proposals/incubation/istio.md) and [due diligence](https://docs.google.com/document/d/1cQiigR5WHQHvo_krUXO6uEaGSB2dWNRkR0cHCAoF5QA/edit) from that application are linked for reference. As the due diligence was completed within the last few months, we believe the information contained within is still up-to-date.

## Graduation criteria requirements

We believe Istio meets the graduation stage requirements in the following way:

### Have committers from at least two organizations.

Istio uses "[maintainer](https://github.com/istio/community/blob/master/ROLES.md#maintainer)" as the noun for a member allowed to accept code contributions. At the time of writing, the project has [61 maintainers](https://github.com/istio/community/blob/master/org/teams.yaml#L12), representing [16 organizations](https://docs.google.com/spreadsheets/d/1KBu4C6HehvMO8fWCgPJbPdvdDcthXjIktrJCHiq4JWg/edit#gid=0).

(The total number of maintainers has decreased since the incubation application as [maintainers move to emeritus status if they have not been active within 1 year](https://github.com/istio/community/blob/master/ROLES.md?plain=1#L234-L236).)

### Have achieved and maintained a Open Source Security Foundation Best Practices Badge.

Istio’s OpenSSF badge can be seen [here](https://bestpractices.coreinfrastructure.org/en/projects/1395).

### Have completed an independent and third party security audit with results published of similar scope and quality as the [Envoy example](https://github.com/envoyproxy/envoy#security-audit), and all critical vulnerabilities need to be addressed before graduation.

A new audit, conducted after joining the CNCF, is [complete and published](https://istio.io/latest/blog/2023/ada-logics-security-assessment/). No critical vulnerabilities were detected.

### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

Istio is led by two groups, the [Steering Committee](https://github.com/istio/community/tree/master/steering) (Steering) and the [Technical Oversight Committee](https://github.com/istio/community/blob/master/TECH-OVERSIGHT-COMMITTEE.md) (TOC).

Steering oversees administrative aspects of the project, including governance, branding, and marketing. The TOC is responsible for cross-cutting project and design decisions. Please see the incubation DD for more context.

Istio uses "[maintainer](https://github.com/istio/community/blob/master/ROLES.md#maintainer)" as the noun for a member allowed to accept code contributions. Contributor roles, including requirements of each role and the privileges that each role grants, are defined in [ROLES.md](https://github.com/istio/community/blob/master/ROLES.md).  Current maintainers are listed in [teams.yaml](https://github.com/istio/community/blob/master/org/teams.yaml) and emeritus maintainers in [emeritus.yaml](https://github.com/istio/community/blob/master/org/emeritus.yaml).

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

Members of Istio’s [Steering Committee](https://github.com/istio/community/blob/master/steering/README.md) are authorized to interact with the CNCF on behalf of the project. The members are selected on an annual basis.

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

Just under 100 users are documented on our [users & case studies page](https://istio.io/latest/about/case-studies), including written case studies and videos of conference talks discussing production use.
