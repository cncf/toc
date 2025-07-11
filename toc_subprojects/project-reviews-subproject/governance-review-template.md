# Governance Review Template

What follows is a governance review and assessment for the ______ project. This
review is carried out by members of the Project Reviews subproject. The review
may have been done because of a change in maturity level for the project, at the
request of the TOC, or as a request by the project itself. If requested by the
project, the review will be provided to the project maintainers. Otherwise the
review will be PR'ed into the project's directory within the cncf/toc repo.

## Summary and Assessment

_Status summary:_

_Exemplary: project has an extraordinary level of governance development and
implementation, and can be used as an example for other projects_

_Satisfactory: project has appropriate governance for its maturity level and is
following that governance_

_Mostly Satisfactory: project has mostly appropriate governance, but needs to
fix one or two things_

_Needs Work: project's governance is lacking and inadequate for its current
level of maturity, and needs substantial work to overcome that_

_In Crisis: project appears to have several outstanding, critical issues that
require attention, clarity, and/or direct guidance from CNCF in order to return
to a state of good health and standing._


**Status:** Exemplary | Satisfactory | Mostly Satisfactory | Needs Work | In Crisis

_Short paragraph summarizing the general state of project governance.
In the event the project governance requires attention, notify the TOC liaison
for their awareness._


### Executing the Assessment

_A brief description that details the timebox the assessment occurred and the
individuals involved in the assessment._


### Must-Fix Items

**The following issues have been identified that need to be resolved before
[project milestone or other requirement]:**

_The items in the list should be summarized, have a prioritized ordering
and are expected to be considered Critical or High,  these are explicit and
implicit blockers to project advancement within the CNCF Matriculation process.
For each item in this list, a corresponding detailed description should be
supplied. Note that which items are required depends on the project's maturity
level._

_File an issue with the project for each item on this list, and link to it._


### Points of Excellence

**The following aspects of governance are exemplary, and can be referenced as
examples for other projects to copy:**

_List of governance aspects where the project is exceeding expectations, or any
novelty in their approach to governance._


### Areas for Improvement

**Over the next year, the project should work on the following issues to improve
its governance, these are considered non-blocking:**

_This is a summarized listing of longer term improvement areas for the project.
These items are strongly encouraged but not required for the project's
maturity level. Fully detailed descriptions are found in the Finding Table.
Items listed here should be in priority ordering and would be considered
Medium and Low._


## Review

**The following review primarily consists of an audit on the project's
self-assessment in their Graduation application.**

_Link to Project Application with Governance and Community Sections_


### Governance Summary

_Short narrative describing the governance type of the project, some general
information about its leadership, and the project's general status and
maturity._

_If the project has any unusual aspects to its governance, describe
them here._

_Link to the project's existing documents where applicable._


### Governance Evolution

**Governance has continuously been iterated upon by the project as a result of
their experience applying it, with the governance history demonstrating
evolution of maturity alongside the project's maturity evolution.**

_How has the project's governance evolved over time?  Is the project steadily
refining/advancing its governance as the project grows and resolves issues?_


### Discoverability

**Clear and discoverable project governance documentation.**

_Where are governance documents located? Primary repo, Community Repo, somewhere
else?_

_How easy is it for potential contributors to find and read the governance
documentation? Is it findable from the project web page? Are governance files
named clearly, and interlinked across the projects repos to the primary?_


### Accuracy and Clarity

**Governance is up to date with actual project activities, including any
meetings, elections, leadership, or approval processes.**

_Are the governance documents up to date?  Do they accurately reflect
current project participants, code and subproject status, etc? Check for:_
- _meetings actually happening_
- _elections actually happening_
- _leadership being updated_
- _PR approval process following documentation_


**Governance clearly documents [vendor-neutrality] of project direction.**

_Check that governance is documented to be vendor-neutral, or at least not
vendor-specific, including:_
- _leadership positions are vendor-neutral_
- _meetings etc. are community-owned and open_
- _blogs and communications channels are community-owned and have guidelines_ 
- _Check for general openness of the governance_
- _Are meetings open or closed?_  
- _Are public records of decisions kept?_


### Decisions and Role Assignments

**Document how the project makes decisions on leadership roles, contribution
acceptance, requests to the CNCF, and changes to governance or project goals.**


**Document how role, function-based members, or sub-teams are assigned,
onboarded, and removed for specific teams (example: Security Response
Committee).**

_Check that Governance.md and other documents cover all of the above._

_This should include:_
- _Definitions of leadership roles_
- _Qualifications, duties, and selection process_
- _Removal process, including removal for inactivity and CoC violations_
- _Who has approver rights_
- _Process of updating the roadmap_


### Maintainers and Maintainer Lifecycle

**Document complete list of current maintainers, including names, contact
information, domain of responsibility, and affiliation.**


**A number of active maintainers which is appropriate to the size and scope of
the project.**

_Check maintainers.md, governance.md, contributor ladder and other files.
Maintainers file should have contact info (not just GH handles), other useful
information, and employer._

_Do a quick sanity check on the number of maintainers.A very large critical
project applying for graduation which has three maintainers is probably at
bus-factor risk._

_In contrast, a narrowly-scoped project with 50 maintainers is probably listing
people who are not actually maintainers. Also check that maintainers on the
list are active using Devstats, GitHub and LFX Insights._

_Query inactive maintainers._

**Document a complete maintainer lifecycle process (including roles, onboarding,
offboarding, and emeritus status).**


**Demonstrate usage of the maintainer lifecycle with outcomes, either through
the addition or replacement of maintainers as project events have required.**

_Look at governance.md/contributor ladder and other documents._

_Look at maintainer.md file history._

_Look for evidence that maintainers have turned over, particularly new maintainer
recruitment and retiring inactive maintainers. Also look that qualifications
for new maintainers are clear and that those qualifications are honored._

**Project maintainers from at least 2 organizations that demonstrates
survivability.**

_Look at maintainer employer affiliation. At graduated level, having everyone
working for the same organization demonstrates some probable problems._

_Inactive maintainers do not fix the maintainer diversity problem._

_Combine this review with lifecycle above; projects with strong recruitment
stories are lower risk even when not very diverse._


### Ownership

**Code and Doc ownership in Github and elsewhere matches documented governance
roles.**

_Right now there is very little ability to audit this._

_You can record here whether the project is using Prow or Sherriff (both of which
are good), but if not you can't verify ownership._

_Parts of the governance documentation that talk about assigning permissions is
also good, but the absence of those isn't notable._


### Code of Conduct

**Document adoption of the CNCF Code of Conduct**


**CNCF Code of Conduct is cross-linked from other governance documents.**

_The CNCF CoC should be present in all appropriate places in the project,
which includes their Community or main repository, and public website at a
minimum._

_Other community and governance documentation should refer to the CoC; for
example, role removal process should mention that people can be removed for CoC
violations._

_If the project is handling their own CoC reports rather than just sending them 
directly to the CNCF CoCC, you need to check if this is adequately documented,
including who receives the reports and how and where they're evaluated._

### Subprojects

**All subprojects, if any, are listed.**


**If the project has subprojects: subproject leadership, contribution, maturity
status documented, including add/remove process.**

_If the project has subprojects, plugins, or other divisions define them here._
- _For each, is ownership and operation of clearly described?_
- _Are any standing committees/teams fully described, including listing their
  members?_
- _Does it conform to, align, and is it within scope of the governance
  expectations of the project?  Does is list communications channels and
  standing meetings?_

_A subproject audit is often very useful to projects because this part of things
tends to bit-rot quite a bit, and you can help them see where they need to
update it._

_If the project has a full listing of subprojects use the text below.
Completeness hinges on whether subproject listings include full ownership,
communications, and alignment info._

**The project maintains a [listing of subprojects]**
**This listing appears to be [Very Complete|Mostly Complete|Incomplete].**

_If the project does not maintain a listing, or if the listing is incomplete,
then fill in the audit info below. Otherwise, delete it._

_Ownership and Operation is "who runs this subproject", Communications includes
meetings, channels and email, and Project Alignment is "what part of the project
is this, and how mature is it"._

The project includes the following sub-projects, plugins, and other notable
divisions:

_Subproject Name/Link_
- Ownership & Operation: Complete / Partial / Missing
- Communications: Complete / Partial / Other
- Project Alignment: Complete / Partial / Conflict
- Notes:


### Contributors and Community

**Contributor ladder with multiple roles for contributors.**

_Note: that a contributor ladder is NOT required, but it is a good idea._
- _Check the supplied one; does it have most of the project roles that actually
  exist?_
- _Do those roles form a reasonable set of steps from new contributor to
  maintainer?_
- _Does each role have qualifications, duties, and powers defined?_

_If the project does not have one, suggest that Project Reviews could help them
create one._

**Clearly defined and discoverable process to submit issues or changes.**

_Files: contributing.md, reviewing.md, release.md, other files which form a
howto for contributors._

_Check for the presence of these and their completeness; as a project grows
it should be adding more and more detailed content to its contributor docs._

_If you have time or reason to, spot check issue and PR activity on the
project's main repos, and the devstats charts related to contributions._

_Are contributions by non-maintainers handled or ignored? Is the project
accumulating a backlog of unreviewed contributions, or concluding them?_

**Project must have, and document, at least one public communications channel for
users and/or contributors.**


**List and document all project communication channels, including subprojects
(mail list/slack/etc.). List any non-public communications channels and what
their special purpose is.**

_Audit the list of communications channels supplied._
- _Does it include at least one general channel?_
- _Does it include all channels you know of?_
- _Are those channels public/owned by the CNCF, or are some of them private
  channels controlled by someone else?_

**Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

- _Does the project have a calendar (ideally integrated with the CNCF calendar)
  that lists all documented meetings?_
- _Are there archival videos and minutes kept of these meetings?_
- _Are the meeting agendas findable and useful?_

**Documentation of how to contribute, with increasing detail as the project
matures.**

_Files: contributing.md, reviewing.md, release.md, other files which form a
howto for contributors._

_Check for the presence of these and their completeness; as a project grows it
should be adding more and more detailed content to its contributor docs. Cite
missing parts based on the level of maturity of the project._

**Demonstrate contributor activity and recruitment.**

_Check provided evidence Spot check issue and PR activity on the project's
main repos, and the devstats charts related to contributions._

_Are contributions by non-maintainers handled or ignored?  Is the project
accumulating a backlog of unreviewed contributions, or concluding them?_


[project milestone or other requirement]: https://github.com/cncf/toc/tree/main/process#how-to-apply-to-move-levels
[vendor-neutrality]: https://contribute.cncf.io/maintainers/community/vendor-neutrality/
