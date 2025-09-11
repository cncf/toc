#  Governance Review Open Telemetry (OTel) Example

# WIP - Not complete

What follows is a governance review and assessment for the Open Telemetry
project.

* **Date Completed:** August 2025
* [Open Telemetry Graduation Application]


## Summary and Assessment


**Status:** Exemplary

Open Telemetry is a large project and has appropriate governance for a project
of its size. They delegate ownership to Special Interest Groups (SIGs), with
each SIG responsible for subprojects in their respective domain.

They have exceptional contributor documentation, end-user engagement, and a
steady pipeline of new contributors.


### Executing the Assessment

[Bob Killen](https://github.com/mrbobbytables/) executed the review in August
2025 as part of the [Open Telemetry Graduation Application].


### Must-Fix Items

**The following issues have been identified that need to be resolved before
[project milestone or other requirement]:**

No issues have been identified that need to be resolved before Graduation.


### Points of Excellence

**The following aspects of governance are exemplary, and can be referenced as
examples for other projects to copy:**

- OTel publishes [marketing guidelines] that provide specific guidance around
  vendor neutrality as well as documenting [how to respond] should a vendor
  run afoul of their marketing rules.
  - Very few projects do this, but they should. As a follow-up, a template
    should be created based off of it that other projects should be able to use.

- OTel has a high contributor diversity, with the "[top contributors]" for the
  last year only account for 2-3% of all contributions.
  - This is a strong indicator of a healthy project with the effort and
    contributions being very well distributed across the organization.


[top contributors]: https://insights.linuxfoundation.org/project/opentelemetry/contributors?timeRange=past365days&start=2024-07-31&end=2025-07-31&widget=contributor-dependency


### Areas for Improvement

**Over the next year, the project should work on the following issues to improve
its governance, these are considered non-blocking:**

There are two very minor things that don't require fixing, but might help with
general discoverability:
- In the [contributor or maintainer guides], provide a brief description of
  [OpenTelemetry Enhancement Proposals] (OTEPs), and [link to the repo].
  - While its not expected that most newer contributors would dive right into 
    drafting a large scale change proposal, its still useful to both new 
    contributors and end users to see how the larger scale technical design
    processes are thought through.
- Publish the OTel meeting recordings to YouTube.
  - The OTel meetings are tracked in a [spreadsheet] that link directly to the
    recordings. The spreadsheet is linked from many locations within the OTel
    GitHub Org and website, however publishing to a service like YouTube
    would make them much more discoverable to the broader community.


[contributor or maintainer guides]: https://github.com/open-telemetry/community/tree/b6a1f396c50b280c74b2d0d887ee0d1c37b59948/guides
[OpenTelemetry Enhancement Proposals]: https://github.com/open-telemetry/opentelemetry-specification/blob/8aa1d27d9cf8c224299dcf0b58bfd5fcaacf88eb/oteps/README.md
[link to the repo]: https://github.com/open-telemetry/opentelemetry-specification

---

## Review

<!--
Bob Tip: When linking to files in GitHub, reference them with a specific commit.
a Governance review is a point-in-time review and things can change. Referencing
the commits will ensure the links will still work, and show what the project
was like at the time of the review.
-->

**The following review primarily consists of an audit on the project's
self-assessment in their Graduation application.**


### Governance Summary

Groups:
- [Governance Committee] (GC) - Top level governing body and final escalation
  point that is responsible creating and continuing to iterate OTel's project
  governance. GC members are elected by the Otel's contributors.
- [Technical Committee][tc] (TC) - Top level technical authority that provides
  the overall technical vision for the project and oversees releases, coding
  standards, spec changes etc. TC members are appointed by the GC.
- [Special Interest Group] (SIG) -  Domain or topic focused subgroups that are
  responsible for their overall area.


[Governance Committee]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/governance-charter.md#scope
[tc]: https://github.com/open-telemetry/community/blob/ffa1e82df098d2bf7140e67e8ef18d2de3c14fdc/tech-committee-charter.md#responsibilities-of-the-technical-committee
[Special Interest Group]: https://github.com/open-telemetry/community/blob/142fbd7bc301247933df044db3833338e538aba6/guides/contributor/processes.md#special-interest-groups-sigs


### Governance Evolution

**Governance has continuously been iterated upon by the project as a result of
their experience applying it, with the governance history demonstrating
evolution of maturity alongside the project's maturity evolution.**
<br>
**Incubating:** Suggested | **Graduated:** Suggested

OTel's project governance has been updated several times since it was initially
created. Many of the changes reflect the project's evolution and growth; a few
examples are:

- [May of 2025] the TC added sponsorship requirements
- [September of 2024] the GC updated the requirements for a 'member in standing'
  who in the OTel community may vote in elections.
- [February of 2024] Changed the scope of the GC to include Project Management
- [June of 2023] the TC codified a section on escalations for Technical
  Decisions

<!--
Bob comments: I looked though their documentation and the git history of the 
major governance docs to look for changes along examples of it being put in to
use. Project behavior -should- reflect their governance.

Also, growth and evolution does not necessarily mean MORE governance. Often for
more mature projects it can be beneficial to remove/refactor/update some of the
governance requirements to reduce the toil of the current maintainers.

-->

[May of 2025]: https://github.com/open-telemetry/community/commit/ffa1e82df098d2bf7140e67e8ef18d2de3c14fdc
[September of 2024]: https://github.com/open-telemetry/community/commit/22a02ad2823ee0445d8e685e86c686e67f809537
[February of 2024]: https://github.com/open-telemetry/community/commit/a0f3b55f0305855b19474af4e686cdf02bba19a3
[June of 2023]: https://github.com/open-telemetry/community/commit/f7271e4db4772c2c44fa4130f12777a24a78d955


### Discoverability

**Clear and discoverable project governance documentation.**
<br>
**Incubating:** Suggested | **Graduated:** Required

Otel's governance documentation is easily found in the project's
[community repo]. It's a pinned repo and linked from the root of the project's
[GitHub Org].

Additionally the relevant sections are linked from their [contributor guide] on
the project's webiste (https://opentelemetry.io).

<!--
Bob comment: No real issue here, their governance is clearly discoverable.

Something to watch for - some projects don't centralize the gov docs in one
place and its easy for them to get out of sync. In these instances I suggest
consolidating on one location and updating all the other locations with pointers
to the new canonical location.

-->

[community repo]: https://github.com/open-telemetry/community
[GitHub Org]: https://github.com/open-telemetry
[contributor guide]: https://opentelemetry.io/docs/contributing/

### Accuracy and Clarity

**Governance is up to date with actual project activities, including any
meetings, elections, leadership, or approval processes.**
<br>
**Incubating:** Suggested | **Graduated:** Required

It is up to date and functioning as written. Some examples of where the
information can be found to verify:

- Their meetings are listed [through multiple calendar options]
- Recordings of their meetings can be found in a [google sheet], and the zoom
  recording links are public.
- Their most recent GC election (2024) is easily accessible through the
  [community repo's elections directory].

<!--
Bob comment: Similar to previous. Their info is easily discoverable.
Not necessarily related to this, but the project doesn't use their official
youtube account that much, and a minor recommendation could be to upload
videos there.

https://www.youtube.com/@otel-official/videos

-->

[multiple calendar options]: https://github.com/dyladan/community/blob/50767bd503704ca2c95f1a31df9d8ac5f7312c43/README.md#calendar
[google sheet]: https://docs.google.com/spreadsheets/d/1SYKfjYhZdm2Wh2Cl6KVQalKg_m4NhTPZqq-8SzEVO6s/edit?gid=0#gid=0
[community repo's elections directory]: https://github.com/open-telemetry/community/tree/fe12ecb89802550b4d3daa9a30646157c540b391/elections/2024


**Governance clearly documents [vendor-neutrality] of project direction.**
<br>
**Incubating:** Suggested | **Graduated:** Required

Otel has both documented and demonstrated good vendor-neutral practices.

- OTel's [Mission, Vision, and Values] explicitly call our vendor-neutrality as
  a core tenant of their.
- Their leadership roles have [maximal representation] rules that apply a limit
  of no more than 2 elected members may be from the same company.
- Their collectors support multiple vendors such as Amazon, Google, and
  Microsoft, as well as other platforms directly such as Kubernetes
  demonstrating they are not tied to a single vendor.
- OTel's [marketing guidelines] provide specific guidance around vendor
  neutrality as well as documenting [how to respond] should a vendor run afoul
  of their marketing rules.

<!-- 
Bob Comment: Otel has done a good job of integrating vendor-neutrality into
their governance. The marketing guidelines and documenting how to respond are
EXCELLENT and something that would be great for more projects to adopt.
-->


[Mission, Vision, and Values]: https://github.com/open-telemetry/community/blob/e9411eec532047b73e68acb492ff7dd3e2b91256/mission-vision-values.md
[maximal representation]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/governance-charter.md#maximal-representation
[marketing guidelines]: https://opentelemetry.io/community/marketing-guidelines/
[how to respond]: https://opentelemetry.io/community/marketing-guidelines/#concerns-and-consequences


### Decisions and Role Assignments

**Document how the project makes decisions on leadership roles, contribution
acceptance, requests to the CNCF, and changes to governance or project goals.**
<br>
**Incubating:** Suggested | **Graduated:** Required

[Decision making and respectful disagreement] is documented as a core value
within OTel's governance. They have several methods for decision making
depending on scope.

- For large-scale technical decisions, Otel has an [RFC process called OTEPs] or
  OpenTelemetry Enhancement Proposals.
- For leadership roles, both the GC and TC have their own election methods.
  - [GC] - GC is a 9 member body that serve 2 year terms, and is responsible
      for the overall governance of the project.
    - No more than 2 members may be from the same company.
    - GC members are elected by active contributors ([members of standing])
      that meet a minimal threshold of contributions. Elections for the roles
      are held using [time-limited approval voting] provided by [Helios], with
      the top candidates being elected to their respective positions.
  - [TC] - The TC is responsible for the [overall technical development] of the
      project and has no max member limit or terms.
    - No more than 25% of the TC may be from the same employer.
    - New TC members may be nominated by any [members of standing], and must
      be sponsored by one existing TC member.
    - Candidates are elected by the current TC from a super-majority vote of the 
      TC members via [Condorcet] or [Single Transferable Vote] methods.
- For [donations of pre-existing code], both the TC and GC are involved in the
  process. The TC does an initial investigation and creates a donation proposal.
  Then it moves to the GC to do a review, trademark audit etc. and final
  approval.
- For general voting on project issues and escalations, they have a
  [documented policy] within their TC charter.
- For all other role related decisions, the specific requirements are outlined
  in their [membership, roles, and responsibilities docs].

  For checks and balances between the GC and TC, the GC is
  **[responsible for the TC charter]** providing an "escape hatch" mechanism
  should there be an unaccounted for issue within the scope of the TC.

[GC]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/governance-charter.md#elections
[members of standing]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/governance-charter.md#members-of-standing
[time-lmited approval voting]: https://en.wikipedia.org/wiki/Approval_voting
[Helios]: https://vote.heliosvoting.org/
[TC]: https://github.com/open-telemetry/community/blob/91525d6a34b1ebb98317bd7b0595fb3b00030e3b/tech-committee-charter.md#establishment-of-the-technical-committee
[overall technical development]: https://github.com/open-telemetry/community/blob/91525d6a34b1ebb98317bd7b0595fb3b00030e3b/tech-committee-charter.md#responsibilities-of-the-technical-committee
[Condorcet]: https://en.wikipedia.org/wiki/Condorcet_method
[Single Transferrable Vote]: https://en.wikipedia.org/wiki/Single_transferable_vote
[Decision making and respectful disagreement]: https://github.com/open-telemetry/community/blob/91525d6a34b1ebb98317bd7b0595fb3b00030e3b/mission-vision-values.md#respectfully-disagree
[documented policy]: https://github.com/open-telemetry/community/blob/91525d6a34b1ebb98317bd7b0595fb3b00030e3b/tech-committee-charter.md#voting-on-project-issues
[donations of pre-existing code]: https://github.com/open-telemetry/community/blob/91525d6a34b1ebb98317bd7b0595fb3b00030e3b/guides/contributor/donations.md
[membership, roles, and responsibilities docs]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/guides/contributor/membership.md
[responsible for the TC charter]: https://github.com/open-telemetry/community/blob/91525d6a34b1ebb98317bd7b0595fb3b00030e3b/tech-committee-charter.md#responsibilities-of-the-technical-committee


**Document how role, function-based members, or sub-teams are assigned,
onboarded, and removed for specific teams (example: Security Response
Committee).**
<br>
**Incubating:** Suggested | **Graduated:** Required

Most of the roles, qualifications and the onboarding/offboarding guidelines are
in OTel's [membership, roles, and responsibilities docs].

Additionally, SIGs can create their own roles or expand on requirements within
their purview to meet the needs of the SIG (e.g. [OTel's Collector Contributor Guide]).

There are a few other roles, such as the [community-manager] that are fully
documented in the [GC's charter]. In all instances they have been documented
sufficiently.

[OTel's Collector Contributor Guide]: https://github.com/open-telemetry/opentelemetry-collector/blob/8951ae8caefcc33fd1767f304701237ce0bcf4b6/CONTRIBUTING.md#membership-roles-and-responsibilities
[community-manager]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/governance-charter.md#community-manager-role
[GC's charter]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/governance-charter.md


### Maintainers and Maintainer Lifecycle

**Document a complete maintainer lifecycle process (including roles, onboarding,
offboarding, and emeritus status).**
<br>
**Incubating:** Suggested | **Graduated:** Required

The full maintainer lifecycle is documented in their [Membership, Roles, and Responsibilities doc].

A  copy of their membership table is copied here for convenience.

| **Role**              | **Responsibilities**                                              | **Requirements**                                                                                                       | **Defined by**                                                                           |
|-----------------------|-------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------|
| member                | active contributor in the community.  reviewer of PRs             | sponsored by 2 approvers or maintainers. multiple contributions to the project.                                        | OpenTelemetry GitHub org member.                                                         |
| triager               | assist with project management and backlog organization.          | nominated by a maintainer. help with issues for one month.                                                             | CONTRIBUTING, CODEOWNERS, or the README.                                                 |
| approver              | approve incoming contributions                                    | highly experienced and active reviewer + contributor to a subproject                                                   | [CODEOWNERS] in GitHub                                                                   |
| maintainer            | set direction and priorities for a subproject                     | demonstrated responsibility and excellent technical judgement for the subproject                                       | [CODEOWNERS], GitHub Team and repo ownership in GitHub                                   |
| emeritus              | position of honor for former maintainers, approvers, and triagers | must have previously held a community role and not have been removed from that role for a [Code of Conduct] violation. | Listed as an emeritus maintainer/approver/triager in CONTRIBUTING, CODEOWNERS, or README |
| specification sponsor | trusted collaborators for the specification                       | nominated by [Technical Committee]                                                                                     | Listed in [Community Members -> Specification and Protos]                                |

[Membership, Roles, and Responsibilities doc]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/guides/contributor/membership.md
[CODEOWNERS]: https://help.github.com/en/articles/about-code-owners
[Community Members -> Specification and Protos]: https://github.com/open-telemetry/community/blob/397865e630debdeb4a492d070c387e0173b42dfb/community-members.md#specifications-and-proto
[Code of Conduct]: https://github.com/open-telemetry/community/blob/ffa1e82df098d2bf7140e67e8ef18d2de3c14fdc/tech-committee-charter.md
[Technical Committee]: https://github.com/open-telemetry/community/blob/ffa1e82df098d2bf7140e67e8ef18d2de3c14fdc/tech-committee-charter.md


**Demonstrate usage of the maintainer lifecycle with outcomes, either through
the addition or replacement of maintainers as project events have required.**
<br>
**Incubating:** Suggested | **Graduated:** Required

Multiple examples of the maintainer lifecycle can be found in the PRs &
[commit history] of OTel's [Community Member List].

[commit history]: https://github.com/open-telemetry/community/commits/main/community-members.md


**Document complete list of current maintainers, including names, contact
information, domain of responsibility, and affiliation.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel lists their members, affiliation, and responsibilities in their
[Community Member List].


**A number of active maintainers which is appropriate to the size and scope of
the project.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel has a significant number of active maintainers, with an
[individual contributor dependency] of 60 contributors authoring 51% of the
commits in the past year.

[individual contributor dependency]: https://insights.linuxfoundation.org/project/opentelemetry/contributors?timeRange=past365days&start=2024-07-31&end=2025-07-31&widget=contributor-dependency


**Project maintainers from at least 2 organizations that demonstrates
survivability.**
<br>
**Incubating:** N/A | **Graduated:** Required

OTel has strong maintainer diversity with no one organization having more than
15% of the top level maintainers. This is mirrored in the current top
[contributing organization list] within lfx insights.

[contributing organization list]: https://insights.linuxfoundation.org/project/opentelemetry/contributors?timeRange=past365days&start=2024-07-31&end=2025-07-31&widget=organizations-leaderboard


### Ownership

**Code and Doc ownership in Github and elsewhere matches documented governance
roles.**
<br>
**Incubating:** Required | **Graduated:** Required

From a random sampling of 10 repos, they all align with what has been
documented.

### Code of Conduct

**Document adoption and adherence to the CNCF Code of Conduct or the project's
CoC which is based off the CNCF CoC and not in conflict with it.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel has [adopted the CNCF Code of Conduct], and documents how to report CoC
violations within the code-of-conduct doc at the root of the community repo.

[adopted the CNCF Code of Conduct]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/code-of-conduct.md


**CNCF Code of Conduct is cross-linked from other governance documents.**
<br>
**Incubating:** Required | **Graduated:** Required

The [CoC is cross-linked] from their governance docs.

[CoC is cross-linked]: https://github.com/search?q=org%3Aopen-telemetry%20code-of-conduct.md&type=code


### Subprojects

**All subprojects, if any, are listed.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel lists their [SIGs and Working Groups] in the README within the community
repo which is generated from [sigs.yaml]. Their SIGs essentially function as
subprojects.

OTel originally tracked their subprojects/repos and ownership in a
[private admin repo], however they have opened a [PR to update sigs.yaml] in
their community repo and have created a workflow that will keep it updated in
the future.


**If the project has subprojects: subproject leadership, contribution, maturity
status documented, including add/remove process.**
<br>
**Incubating:** Suggested | **Graduated:** Required

OTel lists their [SIGs and Working Groups] in the README within the community
repo which is generated from [sigs.yaml].

OTel originally tracked their subprojects/repos and ownership in a
[private admin repo], however they have opened a [PR to update sigs.yaml] in
their community repo and have created a workflow that will keep it updated in
the future.

SIGs and subprojects follow the project level guidelines for membership,
contributing etc, but MAY define additional criteria or roles to meet the
needs of the subproject e.g. [opentelemetry-collector].


[opentelementry-collector]: https://github.com/open-telemetry/opentelemetry-collector/blob/8951ae8caefcc33fd1767f304701237ce0bcf4b6/CONTRIBUTING.md#membership-roles-and-responsibilities

### Contributors and Community

**Contributor ladder with multiple roles for contributors.**
<br>
**Incubating:** Suggested | **Graduated:** Suggested

OTel's [contributor ladder] is well documented and easily discoverable from the
community repo.

[contributor ladder]: https://github.com/open-telemetry/community/blob/fe12ecb89802550b4d3daa9a30646157c540b391/guides/contributor/membership.md


**Clearly defined and discoverable process to submit issues or changes.**
<br>
**Incubating:** Required | **Graduated:** Required

From a random sampling of 5 repos, they all had issue templates for the general
set of things like features, bugs, etc.

For large-scale technical decisions, Otel has an [RFC process called OTEPs] or
OpenTelemetry Enhancement Proposals. The process and steps itself are very well
documented, but not as easily discoverable. There are no specific links or
references in the community repo and contributor/maintainer guide except for
a comment in the [guidelines around donating a repo].


[guidelines around donating a repo]: https://github.com/open-telemetry/community/blob/706839937eb5886bb07da6bd9d63245839e3cc2f/guides/contributor/donations.md?plain=1#L46

**Project must have, and document, at least one public communications channel
for users and/or contributors.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel has all their public [communication channels] available in the
community/repo.


**List and document all project communication channels, including subprojects
(mail list/slack/etc.). List any non-public communications channels and what
their special purpose is.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel has all their public [communication channels] available in the
community/repo, other lists and private channels are listed in the
[list of assets] in the community repo.

[list of assets]: https://github.com/open-telemetry/community/blob/073235c1e018bc658235e384ff2429ade74898eb/assets.md


**Up-to-date public meeting schedulers and/or integration with CNCF calendar.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel lists their public meetings in the [communication channels] section of
their docs, and recordings are made available via a [spreadsheet] that is linked
in various locations including the footer of their website (camera icon).


**Documentation of how to contribute, with increasing detail as the project
matures.**
<br>
**Incubating:** Required | **Graduated:** Required

Otel has documentation for contributors both within their community repo and
published on their website. They have documentation targeted towards
maintainers in their community repo, and these are supported by additional
documentation (e.g. contributing.md) in each repo.

From the random sampling of contributing docs across the project that were
looked at, all were sufficient to get someone up to speed, and many went into
further detail with architecture diagrams etc.

Examples:
- [opentelemetry-rust - otlp architecture](https://github.com/open-telemetry/opentelemetry-rust/blob/f00456f94f9f331f2d63b1d34f9d34bfe8a48eed/docs/design/otlp.md)
- [opentelemetry-go - contributing.md](https://github.com/open-telemetry/opentelemetry-go/blob/d9e9038f076299eeb377d7fa706addd4c37b6271/CONTRIBUTING.md)
- [semantic-conventions - database spans](https://github.com/open-telemetry/semantic-conventions/blob/7d1dbf2958dcbf5ed40e0c233dff6675716cb76b/docs/database/database-spans.md)


**Demonstrate contributor activity and recruitment.**
<br>
**Incubating:** Required | **Graduated:** Required

OTel has a [very active contributor community], and has shown regular growth
with [contributors being promoted to org members].


[very active contributor community]: https://insights.linuxfoundation.org/project/opentelemetry/contributors?timeRange=past365days&start=2024-08-01&end=2025-08-01&widget=active-contributors
[contributors being promoted to org members]: https://github.com/open-telemetry/community/issues?q=is%3Aissue%20state%3Aclosed%20label%3Aarea%2Fgithub-membership



[Open Telemetry Graduation Application]: https://github.com/cncf/toc/issues/1739
[RFC process called OTEPs]: https://github.com/open-telemetry/opentelemetry-specification/blob/773093474e5c2538b66006f4a5ada8d3f44a1e7f/oteps/README.md
[project milestone or other requirement]: https://github.com/cncf/toc/tree/main/process#how-to-apply-to-move-levels
[vendor-neutrality]: https://contribute.cncf.io/maintainers/community/vendor-neutrality/
[Community Member List]: https://github.com/open-telemetry/community/blob/397865e630debdeb4a492d070c387e0173b42dfb/community-members.md
[communication channels]: https://github.com/open-telemetry/community/blob/073235c1e018bc658235e384ff2429ade74898eb/README.md#special-interest-groups
[SIGs and Working Groups]: https://github.com/open-telemetry/community/blob/073235c1e018bc658235e384ff2429ade74898eb/README.md#special-interest-groups
[sigs.yaml]: https://github.com/open-telemetry/community/blob/4bce418b38347b8911fced682be2563bdec54794/sigs.yml
[PR to update sigs.yaml]: https://github.com/open-telemetry/community/pull/2901
[private admin repo]: https://github.com/open-telemetry/community/blob/073235c1e018bc658235e384ff2429ade74898eb/guides/maintainer/github-admin-processes.md#guidelines
[spreadsheet]: https://docs.google.com/spreadsheets/d/1SYKfjYhZdm2Wh2Cl6KVQalKg_m4NhTPZqq-8SzEVO6s/edit?gid=0#gid=0