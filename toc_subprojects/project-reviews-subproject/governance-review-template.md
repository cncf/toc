# Governance Review Template

What follows is a governance review and assessment for the ______ project. This review is carried out by members of the Governance Working Group of TAG Contributor Strategy. The review may have been done because of a change in maturity level for the project, at the request of the TOC, or as a request by the project itself. If requested by the project, the review will be provided to the project maintainers. Otherwise, the review will be submitted to the TOC for their follow-up.

Projects may ask TAG Contributor Strategy for assistance in resolving any issues uncovered by the review. The TAG is available via our [slack channel](https://cloud-native.slack.com/archives/CT6CWS1JN), [email](https://lists.cncf.io/g/cncf-tag-contributor-strategy), [GitHub](https://github.com/cncf/tag-contributor-strategy), or by joining our weekly meetings (listed on the [CNCF public calendar](https://www.cncf.io/calendar/)).

## Summary and Assessment

<!--- Status summary:
Exemplary: project has an extraordinary level of governance development and implementation, and can be used as an example for other projects
Satisfactory: project has appropriate governance for its maturity level and is following that governance
Mostly Satisfactory: project has mostly appropriate governance, but needs to fix one or two things
Needs Work: project's governance is lacking and inadequate for its current level of maturity, and needs substantial work to overcome that
In Crisis: project appears to have several outstanding, critical issues that require attention, clarity, and/or direct guidance from CNCF in order to return to a state of good health and standing.
--->

Status: Exemplary|Satisfactory|Mostly Satisfactory|Needs Work|In Crisis

<!--- Short paragraph summarizing the general state of project governance. In the event the project governance requires attention, notify the TOC liaison for their awareness. --->

### Executing the Assessment

<!--- A brief description that details the timebox the assessment occurred and the individuals involved in the assessment. --->

### Must-Fix Items

The following issues have been identified that need to be resolved before [project milestone or other requirement]:

<!--- The items in the list should be summarized, have a prioritized ordering and are expected to be considered Critical or High,  these are explicit and implicit blockers to project advancement within the [CNCF Matriculation process](https://github.com/cncf/toc/tree/main/process#how-to-apply-to-move-levels). For each item in this list, a corresponding detailed description should be supplied. Note that which items are required depends on the project's maturity level. 
File an issue with the project for each item on this list, and link to it. --->

1. 

### Points of Excellence

The following aspects of governance are exemplary, and can be referenced as examples for other projects to copy:

<!--- List of governance aspects where the project is exceeding expectations, or any novelty in their approach to governance. --->

*

### Areas for Improvement

Over the next year, the project should work on the following issues to improve its governance, these are considered non-blocking:

<!--- This is a summarized listing of longer term improvement areas for the project. These items are strongly encouraged but not required for the project's maturity level. Fully detailed descriptions are found in the Finding Table. Items listed here should be in priority ordering and would be considered Medium and Low. --->

*

## Review

The following review primarily consists of an audit on the project's self-assessment in their Graduation application.

[Project Application with Governance and Community Sections](link to project application here)

### Governance Summary

<!--- Short narrative describing the governance type of the project, some general information about its leadership, and the project's general status and maturity. If the project has any unusual aspects to its governance, describe them here.  Link to the project's existing documents where applicable. --->

### Governance Evolution

* Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.

<!--- How has the project's governance evolved over time?  Is the project steadily refining/advancing its governance as the project grows and resolves issues? --->

### Discoverability

* Clear and discoverable project governance documentation.

<!--- Where are governance documents located?  Primary repo, Community Repo, somewhere else? --->

<!--- How easy is it for potential contributors to find and read the governance documentation? Is it findable from the project web page? Are governance files named clearly, and interlinked across the projects repos to the primary? --->

### Accuracy and Clarity

* Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.

<!--- Are the governance documents up to date?  Do they accurately reflect current project participants, code and subproject status, etc?  Check for:
- meetings actually happening
- elections actually happening
- leadership being updated
- PR approval process following documentation --->

* Governance clearly documents [vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.

<!--- Check that governance is documented to be vendor-neutral, or at least not vendor-specific, including:
- leadership positions are vendor-neutral
- meetings etc. are community-owned and open
- blogs and communications channels are community-owned and have guidelines 
Also check for general openness of the governance.  Are meetings open or closed?  
are public records of decisions kept? --->

### Decisions and Role Assignments

* Document how the project makes decisions on leadership roles, contribution acceptance, requests to the CNCF, and changes to governance or project goals.
* Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).

<!--- Check that Governance.md and other documents cover all of the above.  This should include:
- Definitions of leadership roles
- Qualifications, duties, and selection process
- Removal process, including removal for inactivity and CoC violations
- Who has approver rights
- Process of updating the roadmap
--->

### Maintainers and Maintainer Lifecycle

* Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.
* A number of active maintainers which is appropriate to the size and scope of the project.

<!--- Check maintainers.md, governance.md, contributor ladder and other files.
Maintainers file should have contact info (not just GH handles), other useful information, and employer
Do a quick sanity check on the number of maintainers.  A very large critical project applying for graduation which has three maintainers is probably at bus-factor risk.  In contrast, a narrowly-scoped project with 50 maintainers is probably listing people who are not actually maintainers.
Also check that maintainers on the list are active using Devstats, GitHub and LFX Insights.  Query inactive maintainers. --->

* Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).
* Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.

<!--- Look at governance.md/contributor ladder and other documents.  Look at maintainer.md file history.
Look for evidence that maintainers have turned over, particularly new maintainer recruitment
and retiring inactive maintainers.  Also look that qualifications for new maintainers are
clear and that those qualifications are honored. --->

* Project maintainers from at least 2 organizations that demonstrates survivability.

<!--- Look at maintainer employer affiliation. At graduated level, having everyone working for the same
organization demonstrates some probable problems.  Inactive maintainers do not fix the maintainer
diversity problem. Combine this review with lifecycle above; projects with strong recruitment stories
are lower risk even when not very diverse. --->

### Ownership

* Code and Doc ownership in Github and elsewhere matches documented governance roles.

<!--- right now there is very little ability to audit this.  You can record here 
whether the project is using Prow or Sherriff (both of which are good), but if not
you can't verify ownership. Parts of the governance documentation that talk about
assigning permissions is also good, but the absence of those isn't notable. --->

### Code of Conduct

* Document adoption of the CNCF Code of Conduct
* CNCF Code of Conduct is cross-linked from other governance documents.

<!--- The CNCF CoC should be present in all appropriate places in the project, which 
includes their Community or main repository, and public website at a minimum.  Other
community and governance documentation should refer to the CoC; for example, role
removal process should mention that people can be removed for CoC violations.  
If the project is handling their own CoC reports rather than just sending them 
directly to the CNCF CoCC, you need to check if this is adequately documented, including
who receives the reports and how and where they're evaluated. --->

### Subprojects

* All subprojects, if any, are listed.
* If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.

<!--- If the project has subprojects, plugins, or other divisions define them here. For each, is ownership and operation of clearly described? Are any standing committees/teams fully described, including listing their members? Does it conform to, align, and is it within scope of the governance expectations of the project?  Does is list communications channels and standing meetings?
A subproject audit is often very useful to projects because this part of things tends to 
bit-rot quite a bit, and you can help them see where they need to update it. 
If the project has a full listing of suboprojects use the text below.  Completeness hinges on whether subproject listings include full ownership, communications, and alignment info. --->

The project maintains a [listing of subprojects](LINK).  This listing appears to be [Very Complete|Mostly Complete|Incomplete].

<!--- if the project does not maintain a listing, or if the listing is incomplete, then
fill in the audit info below.  Otherwise, delete it.  Ownership and Operation is "who runs this subproject", Communications includes meetings, channels and email, and Project Alignment is "what part of the project is this, and how mature is it".  --->

The project includes the following sub-projects, plugins, and other notable divisions:

| Area | Ownership and Operation | Communications | Project Alignment | Notes |
|:-----|:-----------------------:|:---------------:|:------------------|:---|
|*sub-project*| Complete/Partial/Missing | Complete/Partial/Other | Complete/Partial/Conflict | |

### Contributors and Community

* Contributor ladder with multiple roles for contributors.

<!-- Note that a contributor ladder is NOT required, but it is a good idea.  Check the supplied one;
does it have most of the project roles that actually exist?  Do those roles form a reasonable set
of steps from new contributor to maintainer?  Does each role have qualifications, duties, and powers 
defined?
If the project does not have one, suggest that TAG-CS could help them create one. --->

* Clearly defined and discoverable process to submit issues or changes.

<!-- Files: contributing.md, reviewing.md, release.md, other files which form a
howto for contributors.  Check for the presence of these and their completeness;
as a project grows it should be adding more and more detailed content to its contributor
docs. 
If you have time or reason to, spot check issue and PR activity on the project's main repos,
and the devstats charts related to contributions.  Are contributions by non-maintainers
handled or ignored?  Is the project accumulating a backlog of unreviewed contributions, or
concluding them? --> 

* Project must have, and document, at least one public communications channel for users and/or contributors.
* List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.

<!-- audit the list of communications channels supplied.  Does it include at least one
general channel?  Does it include all channels you know of?  Are those channels public/
owned by the CNCF, or are some of them private channels controlled by someone else? -->

* Up-to-date public meeting schedulers and/or integration with CNCF calendar.

<!-- Does the project have a calendar (ideally integrated with the CNCF calendar)
that lists all documented meetings?  Are there archival videos and minutes kept
of these meetings?  Are the meeting agendas findable and useful? --> 

* Documentation of how to contribute, with increasing detail as the project matures.

<!-- Files: contributing.md, reviewing.md, release.md, other files which form a
howto for contributors.  Check for the presence of these and their completeness;
as a project grows it should be adding more and more detailed content to its contributor
docs. Cite missing parts based on the level of maturity of the project. --->

* Demonstrate contributor activity and recruitment.

<!-- Check provided evidence Spot check issue and PR activity on the project's main repos,
and the devstats charts related to contributions.  Are contributions by non-maintainers
handled or ignored?  Is the project accumulating a backlog of unreviewed contributions, or
concluding them? --> 


