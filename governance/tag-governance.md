# CNCF Technical Group Governance

This document outlines the scope and responsibilities of the CNCF Technical
Groups

The CNCF Technical Groups are comprised of four types of groups that serve
different purposes or have different scopes:

- **Technical Advisory Group (TAG)**: Primary organizational unit that is
  aligned with an industry problem domain or a group that serves a common need
  across TAGs & CNCF Projects.   
- **Subproject**: An ongoing or “permanent” project as part of a TAG or the TOC
  directly.  
- **Initiative:** A lightweight organizational unit used for TAG or TOC work
  that is tightly scoped and time-bound.  
- **Community Group:** A topic or domain focused group used to serve as a
  rallying point for community members to discuss, share knowledge, and
  coordinate potential future initiatives. 

- [CNCF Technical Group Governance](#cncf-technical-group-governance)
- [Keywords](#keywords)
- [Technical Advisory Group (TAG)](#technical-advisory-groups-\(tags\))
  - [Goals](#goals)
  - [Formation](#formation)
  - [Requirements](#requirements)
  - [Roles, Terms, and Elections](#roles,-terms,-and-elections)
    - [Chair](#chair)
      - [Requirements](#requirements-1)
      - [Duties](#duties)
      - [Activity Expectations](#activity-expectations)
      - [Escalations](#escalations)
    - [Tech Lead](#tech-lead)
      - [Requirements](#requirements-2)
      - [Duties](#duties-1)
      - [Activity Expectations](#activity-expectations-1)
      - [Escalations](#escalations-1)
    - [Other Roles](#other-roles)
- [Subprojects](#subprojects)
  - [Goals](#goals-1)
  - [Formation](#formation-1)
  - [Requirements](#requirements-3)
  - [Roles, Terms, and Elections](#roles,-terms,-and-elections-1)
    - [Subproject Leads](#subproject-leads)
      - [Requirements](#requirements-4)
      - [Duties](#duties-2)
      - [Activity Expectations](#activity-expectations-2)
      - [Escalations](#escalations-2)
    - [Other Roles](#other-roles-1)
- [Initiatives](#initiatives)
  - [Goals](#goals-2)
  - [Formation](#formation-2)
    - [TAG Initiatives](#tag-initiatives)
    - [Subproject Initiatives](#subproject-initiatives)
    - [TOC Initiatives](#toc-initiatives)
  - [Requirements](#requirements-5)
  - [Roles](#roles)
- [Community Groups](#technical-community-groups)
  - [Rules & Guidelines](#rules-&-guidelines)    - [Formation](#formation-3)
    - [Formation](#formation-3)

## Keywords

This document uses [RFC2119 keywords] to indicate requirement levels as clearly
and concisely as possible. RFC2119 has long been used by many governance and
standards bodies, such as the [Internet Engineering Task Force (IETF)] and the
[Kubernetes project].

- **MUST** - This word, or the terms "**REQUIRED**" or "**SHALL**," means that
  the item is required.  
- **MUST NOT** - This phrase, or the phrase "**SHALL NOT**", means that the item
  is not permitted or must be avoided.  
- **SHOULD** -  This word, or the adjective "**RECOMMENDED**," means the item is
  suggested but may be ignored for valid reasons.  
- **SHOULD NOT** - This phrase, or the phrase "**NOT RECOMMENDED**," means the
  item is acceptable under certain circumstances.  
- **MAY** - This word, or the adjective "**OPTIONAL**," means the item is truly
  optional.


[RFC2119 keywords]: https://www.ietf.org/rfc/rfc2119.txt
[Internet Engineering Task Force (IETF)]: https://en.wikipedia.org/wiki/Internet_Engineering_Task_Force
[Kubernetes project]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance-requirements.md


## Technical Advisory Groups (TAGs)

A CNCF TAG will oversee and coordinate the interests pertaining to a logical
area of needs across projects, other CNCF groups, and users. They are long-lived
groups that report directly to the TOC. Areas covered by TAGs may be vertically
focused on industry problem domains that align with the projects within the CNCF
and/or horizontal that support the other TAGs or the projects themselves.

**Examples:**
- **Vertical:** Infrastructure, Workloads  
- **Horizontal:** Security & Compliance, Operational Resilience


### Goals

- Strengthen the project ecosystem to meet the needs of end users and project
  contributors.
- Serve as a “bridge” between projects to surface common problems and drive
  alignment.  
- Identify gaps in the CNCF project portfolio. Find and attract projects to fill
  these gaps.  
- Reduce some project workload on TOC while retaining executive control & tonal
  integrity with this elected body.  
- Provide technical expertise during the evaluation of projects to join the CNCF
  and as they move levels.  
- Educate and inform users with unbiased, effective, and practically useful
  information scoped to cloud native.  
- Focus attention & resources on helping foster project maturity, systematically
  across CNCF projects.


### Formation

TAGs are formed when there is a documented, outstanding gap in the technical
domain coverage, needs, and direction within cloud native. TAGs are expected to
be very long lived groups that perform standing functions for projects and their
adopters.

Due to the longevity of a TAG, existing groups interested in becoming a TAG need
to engage the TOC, the TOC will review the request against the current needs and
movement of the ecosystem and make a determination.

Example:  
A domain-focused community group has regularly submitted and completed
initiatives that cover gaps in the current TAG structure. After a discussion
with the TOC, a draft charter with scopes is drafted. The TOC then votes on the
TAG charter PR, and the TAG is formed after the vote is passed.


### Requirements

- MUST have a charter that outlines scope, potential overlaps with other TAGs,
  and any operational deviations from TAGs.  
- MUST enumerate any roles within the TAG and the responsibilities of each along
  with any role membership requirements and lifecycle.  
- MUST maintain TAG “metadata” - a list of leads, subprojects, and initiatives.  
- Where applicable, MUST participate in Project Reviews & Contributor Strategy
  TOC subproject.
- If applicable, MUST document services provided to other TAGs or projects.
- MUST hold at least one regularly scheduled *public* monthly meeting, recorded
  and uploaded to YouTube.
- MUST provide periodic updates to the TOC and TOC liaisons on their health,
  subprojects, and initiatives.
- SHOULD provide timely review and completion of project and TOC requests related
  to the moving levels process.  
- SHOULD provide regular ecosystem updates on the domain, such as through the
  regular CNCF community meeting, KubeCon presentations, or to the TOC directly.  
- SHOULD organize KubeCon + CloudNativeCon presence; such as CFPs, Project
  Pavilion booths etc.  
- SHOULD form and document subprojects to organize specific ongoing work or
  services.  
- SHOULD produce well researched and informed technical white papers, best
  practices, and other ecosystem supportive assets.  
- SHOULD surface gaps or ecosystem trends to related projects.  
- SHOULD check in with graduated and incubating projects to understand their
  health, features, changes, etc. and notify the TOC of anything significant,
  any gaps, novelty, interesting, concerning.


### Roles, Terms, and Elections

#### Chair

A Chair is an individual who is responsible for overseeing and driving the
planning, execution, and accountability of efforts within the TAG. TAG chairs
are expected to equally share responsibilities, and should rotate/take turns as
“primary” chair to allow responsibilities to be executed while ensuring
consistency, availability, and continuity over the course of Chair terms and
availability.

Chairs build unity in the purpose of the group for which they govern and
oversee. This includes sufficient consideration of ideas, topics, and issues
presented, facilitation of healthy discussion, delegation of tasks (as
appropriate), and ensuring their group's adherence to both the TOC’s technical
leadership principles and the TOC’s operating principles.

Chairs ensure agreed tasks are carried out and that group decisions have
consensus. They are responsible for reporting progress to their parent body
(TOC).

Chairs also primarily perform administrative functions including collecting and
compiling topics for the agenda, chairing the meeting, ensuring that quality
meeting minutes are published, and follow-up actions tracked and resolved.

The expected minimum time commitment is at least 10% of their working hours, but
expected to increase (e.g up to 20%) during times with a high volume of requests
including, but not limited to sandbox inclusion, due diligence for
incubation/graduation, and/or or/and KubeCon preparations.


##### Requirements

- MUST be approved by the TOC.  
- MUST adhere to the [Technical Leadership Principles].  
- Membership MUST be tracked alongside the TAG metadata.  
- MUST have prior experience within the CNCF (e.g. as an active TAG community
  member), but MAY be approved by their parent governing body for an explicitly
  documented reason such as extensive domain expertise, formation of a new
  governing body (e.g. new TAG or Working Group), or demonstrating significant
  aptitude for people & community management.


##### Duties

- SHOULD drive charter changes (including creation) to get community buy-in but
  MAY delegate content creation to other leads or community members.  
- SHOULD define how priorities and commitments are managed and set a cap on
  concurrent efforts of the group per quarter. MAY delegate to other leads as
  needed.  
- MUST ensure leadership roles are documented, accurate, and updated in any
  applicable locations (GitHub teams, mailing lists, etc).  
- MUST ensure all subprojects are documented and operate within the scope
  defined in their charter.  
- MUST coordinate communication and be a connector with other community groups
  like TAGs and the TOC, but MAY delegate the actual communication and creation
  of content to other contributors where appropriate.  
- SHOULD facilitate meetings but MAY delegate to other Leads or active community
  members.  
- SHOULD mentor future Leads as part of active and continuous succession
  planning.  
- MUST, in conjunction with the Tech Leads, regularly triage their issues & PRs
  but MAY delegate to other contributors to fulfill these responsibilities.  
- MUST organize KubeCon + CloudNativeCon presence (CFPs, Project Pavilion booth,
  etc) and approve content presented by or about the TAG, but MAY delegate to
  other contributors.  
- MUST, in conjunction with the Tech Leads, regularly provide status updates on
  their groups, subprojects, and any other initiatives.   
- MUST coordinate service desk requests on behalf of the TAG.  
- MUST keep the TOC liaisons informed of the group’s activities and progress
  against the charter.


##### Activity Expectations

- Chairs serve 2-year terms and MAY submit for re-election.  
- Chairs SHOULD remain active and responsive.  
- Chairs MAY decide to step down at any time and propose a replacement to
  complete their term.   
  - The candidate SHOULD be supported by a majority of the current active leads,
    with support backed up by documentation, such as links to GitHub activity
    or documented deliverables.
  - The candidate is REQUIRED to be approved by the TOC.  
- Chairs taking an extended leave of 1 or more months MUST coordinate with
  other leads to ensure the role is adequately covered during their leave.  
- Chairs going on leave for 1-3 months MUST work with other Chairs and Tech
  Leads to identify a temporary replacement.
- Active Leads are REQUIRED to notify their parent governing body (e.g., TAG
  notifying their TOC liaison) and receive confirmation of notice before
  pursuing removal of an inactive or Lead found to be in dereliction of duty.
- Inactive Leads MUST be removed through a super-majority vote of the parent
  governing body if there is no other explicitly defined removal policy.
- If there is no defined removal policy and not enough active parent governing
  body leads to hold a vote, the active members MUST escalate to the next-level
  parent governing body, where those members may count as additional voters.


##### Escalations

- Lead membership disagreements MAY be escalated to their parent governing body,
  the CoCC as applicable, or in the event of further dispute, up to the top
  governing body itself (TOC).


#### Tech Lead

A Tech Lead is an individual responsible for providing subject matter expertise
of a defined domain that supports the objectives of the TAG or initiative. They
may share this responsibility with other Tech Leads as appropriate or directed
by the TOC or their TAG Chair. Tech Leads are accountable to the Chair or TOC
for the group they are assigned.

Tech Leads ensure timely delivery and completion of defined milestones and
deliverables for TAGs or initiatives for which they are assigned. 

Tech Leads are also responsible for guiding the technical aspects of efforts by
a TAG or initiative. This includes ensuring neutrality (no kingmaking) in
recommendations and enforcement of TOC operating principles in deliverables.
They ensure quality and integrity in the deliverables of each group, providing
technical mentorship and guidance to contributors of their group, and
collaborating within and across groups to ensure consistency and concurrence.  

The expected minimum time commitment is at least 10% of their working hours, but
expected to increase (e.g up to 20%) during times with a high volume of requests
including, but not limited to sandbox inclusion, due diligence for
incubation/graduation, and/or or/and KubeCon preparations.


##### Requirements

- SHOULD be nominated by Chairs.  
- MUST be approved by the TOC.  
- MUST adhere to the [Technical Leadership Principles].  
- Membership MUST be tracked alongside the TAG metadata.  
- MUST have prior experience within the CNCF (e.g. as an active TAG community
  member), but MAY be approved by their parent governing body for an explicitly
  documented reason such as extensive domain expertise, formation of a new
  governing body (e.g. new TAG or Working Group), or demonstrating significant
  aptitude for people & community management. 

##### Duties

- MUST approve & facilitate the creation of new subprojects.  
- MUST approve & facilitate decommissioning of existing subprojects.  
- If applicable MUST work with the TOC to evaluate the technical fit and state
  of projects applying to join the CNCF or move levels. MAY delegate to other
  leads or domain experts.  
- MUST review technical outputs or artifacts from the group for accuracy. This
  includes white papers, code, other forms of technical documentation, etc.  
- SHOULD mentor future Leads as part of active and continuous succession
  planning.  
- MUST, in conjunction with the Chairs, regularly triage their issues & PRs but
  MAY delegate to other contributors to fulfill these responsibilities.  
- MUST, in conjunction with the Chairs, regularly provide updates on the status
  of their groups, subprojects, and any other initiatives.   
- SHOULD attend regular group meetings.

##### Activity Expectations

- Tech Leads serve 2 year terms, and MAY submit for re-election.  
- Tech Leads MAY decide to step down at any time and propose a replacement to
  complete their term.   
  - The candidate SHOULD be supported by a majority of the current active leads,
    with support backed up by documentation, such as links to GitHub activity
    or documented deliverables.
  - The candidate is REQUIRED to be approved by the TOC.  
- Tech Leads taking an extended leave of 1 or more months MUST coordinate with
  other leads to ensure the role is adequately covered during their leave.  
- Tech Leads going on leave for 1-3 months MUST work with other Leads to
  identify a temporary replacement.  
- Active Leads are REQUIRED to notify their parent governing body (e.g., TAG
  notifying their TOC liaison) and receive confirmation of notice before
  pursuing removal of an inactive or Lead found to be in dereliction of duty.
- Inactive Leads MUST be removed through a super-majority vote of the parent
  governing body if there is no other explicitly defined removal policy.
- If there is no defined removal policy and not enough active parent governing
  body leads to hold a vote, the active members MUST escalate to the next-level
  parent governing body, where those members may count as additional voters.


##### Escalations

- Lead membership disagreements MAY be escalated to their parent governing body,
  the CoCC as applicable, or in the event of further dispute, up to the top
  governing body itself (TOC).


#### Other Roles

TAGs MAY have other roles, but are REQUIRED to be documented in the TAG charter
and adhere to the requirements outlined in the Other Roles section of the CNCF
Community Group Leadership Requirements.

**NOTE:** Roles that assume some of the responsibilities of the Chair or Tech
Lead role MUST be described in the TAG charter and REQUIRE TOC approval.

Roles that do not assume any Chair or Tech Lead responsibilities may be created
at the TAG Leads discretion as long as they adhere to the following
requirements:

- MUST be documented with a description, scope, and responsibilities.  
- MUST have a documented lifecycle policy including how they are appointed, any
  potential terms, and an inactivity policy.  
- MUST be tracked alongside the other TAG metadata.  
- MUST adhere to the [Technical Leadership Principles].  
- SHOULD remain active and responsive in their Roles.


## Subprojects

Specific on-going work efforts within TAGs or under the TOC are divided into
subprojects. Subprojects may be code, documentation, services, or any other
persistent work effort. Subprojects have leads that function as the primary
organizer, domain experts, or points of contact. 

### Goals

- Assist the projects, TOC, and other TAGs  
- Remove ambiguity and provide guidance, increasing knowledge and understanding
  on the subproject domain  
- Support projects and end users in addressing items in scope of the subproject  
- Coordinate awareness, understanding, concurrence of the subproject
  deliverables with projects, end users, and other defined stakeholders  
- Advance and accelerate the cloud native mission within the scope of the
  subproject  
- Enact the technical vision of the TOC within scope of the subproject


### Formation

Subprojects require a charter describing their purpose, scope, roles, and any
regular deliverables or services they provide.  This charter MUST be approved by
the TOC. 

It is strongly recommended to share a draft charter and work with the TOC early
to ensure a streamlined approval process.


### Requirements

- MUST be documented including purpose, scope, roles and any deliverables or
  services they provide.  
- MUST remain actively in use. A subproject that has gone 6+ months without
  delivering on its goal or providing a service should be considered for
  retirement or put on “pause”.  
- MUST regularly report on status to leads. Updates SHOULD occur at least
  quarterly.  
- SHOULD have at least 2 subproject leads. In the absence of defined subproject
  leads, responsibility for the subproject automatically falls to the Chairs and
  Tech Leads of the TAG. For TOC Subprojects, subproject leads fall to the TOC
  Shadows first and TOC Members next, as needed.


### Roles, Terms, and Elections

#### Subproject Leads

Subproject leads are community members and subject matter experts tasked with
organizing and shepherding the execution of the subproject. They adhere to the
requirements, duties and expectations outlined in the Subproject Lead role
description in CNCF Community Group Leadership Requirements.

Recommended minimum number of leads: 2  
Term: 1 year  
Election requirements: Super-majority vote of the parent governing body.

Subproject Leads are the technical authority for a subproject. They MUST have
demonstrated both good judgement and responsibility towards the health of that
subproject.


##### Requirements

- MUST be approved by a super-majority vote of the parent governing body.  
- MUST adhere to the [Technical Leadership Principles].  
- MUST be nominated by the TAG’s Tech Leads (minimum 1) for TAG led subprojects,
  or by the TOC (minimum 1) for TOC led subprojects.
- Membership MUST be tracked alongside the parent governing body metadata.  
- SHOULD have prior experience within the CNCF (e.g. as an active TAG community
  member), but MAY be approved by their parent governing body for an explicitly
  documented reason such as extensive domain expertise, formation of a new
  governing body (e.g. new TAG or Working Group), or demonstrating significant
  aptitude for people & community management. 


##### Duties

- SHOULD set direction and make or approve design decisions for their
  subproject - either directly or through delegation.   
  - Design decisions that make changes outside of the subproject’s defined scope
    MUST be approved by the Tech Lead(s) of the TAG.  
- MUST actively participate in triaging issues and reviewing PRs related to
  their subproject.  
- MUST regularly report subproject status to TAG leadership.  
- SHOULD be an escalation point for technical discussions and decisions in the
  subproject.  
- SHOULD set milestone priorities.   
- SHOULD ensure a healthy process for discussion and decision making is in
  place. MAY make decisions to resolve conflicts.


##### Activity Expectations

- Subproject Leads serve 1 year terms, and  MAY submit for re-election.  
- Subproject Leads MAY decide to step down at any time and propose a replacement
  to complete their term.   
  - The candidate SHOULD be supported by a majority of the current active leads,
    with support backed up by documentation, such as links to GitHub activity
    or documented deliverables. 
  - The candidate is REQUIRED to be approved by the parent TAG’s Tech Leads, or
    in the case of TOC lead subprojects, the TOC itself.
- Subproject Leads taking an extended leave of 1 or more months MUST coordinate
  with other leads to ensure the role is adequately covered during their leave.  
- Subproject Leads going on leave for 1-3 months MAY work with other Leads to
  identify a temporary replacement.  
- Active Leads are REQUIRED to notify their parent governing body (e.g., TAG
  notifying their TOC liaison) and receive confirmation of notice before
  pursuing removal of an inactive or Lead found to be in dereliction of duty.
- Inactive Leads MUST be removed through a super-majority vote of the parent
  governing body if there is no other explicitly defined removal policy.
- If there is no defined removal policy and not enough active parent governing
  body leads to hold a vote, the active members MUST escalate to the next-level
  parent governing body, where those members may count as additional voters.


##### Escalations

- Lead membership disagreements MAY be escalated to their parent governing body,
  the CoCC as applicable, or in the event of further dispute, up to the top
  governing body itself (TOC).


#### Other Roles

Subprojects MAY have other roles, but are REQUIRED to be documented in the
subproject charter and adhere to the requirements outlined in the Other Roles
section of the CNCF Community Group Leadership Requirements.

**NOTE:** Roles that assume some of the responsibilities of the Subproject Lead
role MUST be described in the Subproject charter and REQUIRE approval from the
parent governing body.

Roles that do not assume any Subproject Lead responsibilities may be created at
the Subproject Leads discretion as long as they adhere to the following
requirements:

- MUST be documented with a description, scope, and responsibilities.  
- MUST have a documented lifecycle policy including how they are appointed, any
  potential terms, and an inactivity policy.  
- MUST be tracked alongside the other TAG metadata.  
- MUST adhere to the [Technical Leadership Principles].  
- SHOULD remain active and responsive in their Roles.


## Initiatives

Initiatives are short-term, time bound work that have a pre-defined objective
and exit criteria. Some examples include white papers, one-off reports, related
to a subproject, or working with the TOC on a specific initiative (e.g. DTR).
They do not have a charter, or defined leads, but can have separate meetings if
it is beneficial to the initiative completing its goal. Each initiative MUST
have a TAG Tech Lead, Subproject lead, or TOC member assigned to oversee and
ensure timely completion and closure of the work within the scope of the parent
governing group's charter.


### Goals

- Serve as a lightweight organizational unit to coordinate work.  
- Recognize community members who have helped execute and deliver on the
  initiative.  
- Ensure high-velocity completion of scoped effort with clear outcomes aligned
  with the CNCF mission and TOC vision.

### Formation

There are three types of initiatives with different formation requirements, TAG
initiatives, subproject initiatives, and TOC initiatives.


#### TAG Initiatives

General initiatives can be any "piece of work," such as white papers, reports,
etc., that align with the scope of a TAG. The Tech Leads manage TAG Initiatives
and MUST have at least one assigned as the sponsor to oversee the initiative's
execution. They MAY delegate some of the responsibility of organizing and
shepherding the initiative, but are still ultimately responsible for it.

TAG initiatives can span multiple TAGs when something spans multiple tech
domains, however one is ultimately assigned as the responsible group for
execution and reporting. Initiatives that span many TAGs, or have no clear owner
should be proposed as TOC initiatives. 


#### Subproject Initiatives

Subproject-based initiatives are defined in the subproject’s charter and are
tailored to meet the subproject’s needs. They are essentially a customized
GitHub issue template used to track the work of the subproject.


#### TOC Initiatives

The TOC may create and own initiatives when the scope of work spans the majority
of TAGs or coordinate work around specific goals. In these instances, at least
one TOC member MUST be assigned to organize, shepherd, and report on the
initiative's status.


### Requirements

- MUST be documented, including goals, milestones, deliverables, and any exit
  criteria (e.g., a tracking GitHub issue).  
- Initiatives MUST have more than two participants not including the assigned
  lead for oversight.  
- SHOULD actively be worked on. An initiative that has made little progress in
  3 months should be canceled or paused until other contributors can be
  identified.   
- Initiatives SHOULD not last longer than 2 quarters but MAY request an
  extension from their parent governing body with reasonable justification.  
- MUST check in at least monthly to assess the initiative's continued need,
  progress, and ability to meet.  
- SHOULD acknowledge everyone who actively helped complete the initiative.  
- Subproject based initiatives MAY have different requirements that are outlined
  in the subproject charter.


### Roles

Assigned Tech Leads, Subproject Leads or TOC members serve as the general
oversight of initiatives depending on the parent governing body.


## Technical Community Groups

Technical Community Groups are a part of the broader
[CNCF Community Groups Program]. They are topic and domain-focused or
locality-based groups used as a rallying point for community members to discuss,
share knowledge, and coordinate potential future initiatives. They are free to
propose and shepherd initiatives when there is an explicit deliverable or
objective they wish to work on.

Community Groups have [much broader exposure], and have a suite of tools
tailored for community management, making them the best place for
discussion-oriented groups.

Additionally, it is likely that future TAGs and subprojects will likely start as
community groups. If they have proposed and completed numerous initiatives and
cover a gap in the currently defined groups, then they may engage the TOC to
determine if it makes sense to propose the formation of a TAG.


[much broader exposure]: https://community.cncf.io/


### Rules & Guidelines

Adhere to the rules and guidelines outlined in the
[Cloud Native Community Groups Program].


### Formation

Community Groups follow the formation process outlined in
[Community Groups README].

**NOTE:** Currently Technology focused groups do not have an issue template.
These groups will require TOC approval, and will be provisioned with a GitHub
discussion category for async communication instead of the built-in forum within
Bevy (Community Groups backing platform).

[Community Groups README]: https://github.com/cncf/communitygroups?tab=readme-ov-file#how-to-apply


<!-- Common Links -->

[Technical Leadership Principles]: https://github.com/cncf/toc/blob/main/.md#technical-leadership-principles
[CNCF Community Groups Program]: https://github.com/cncf/communitygroups