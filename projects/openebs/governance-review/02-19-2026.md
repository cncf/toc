## using this to draft the review for OpenEBS incubation app

#  Governance Review OpenEBS

# WIP - Not complete

What follows is a governance review and assessment for the OpenEBS
project.

* **Date Completed:** February 2026
* [OpenEBS Incubation Application - https://github.com/cncf/toc/issues/1601]



## Summary and Assessment


**Status:** Satisfactory

OpenEBS is an umbrella project with five sub-projects within it. The core maintainers for OpenEBS 
have done well in establishing the overall governance of the project and working in conjunction
with the sub-project maintainers to establish governance for the sub-projects. All projects
under the OpenEBS umbrella essentially follow the core OpenEBS governance model. 

They have good contributor documentation, end-user engagement, and are in the process of 
growing a pipeline of new contributors.


### Executing the Assessment

[Khallai Taylor](https://github.com/khallai/) completed the review in February
2026 as part of the [OpenEBS Incubation Application - https://github.com/cncf/toc/issues/1601].


### Must-Fix Items

**The following issues have been identified that need to be resolved before
[project milestone or other requirement]:**

No pressing issues 

### Points of Excellence

**The following aspects of governance are exemplary, and can be referenced as
examples for other projects to copy:**

- OpenEBS has a well put together website https://openebs.io/docs with easy to find and search documentation. 
- The quickstart guides for OpenEBS installation and corresponding sub-project prerequisites guides 
are easy to navigate and helps users to start using OpenEBS quickly https://openebs.io/docs/quickstart-guide/installation
- OpenEBS governance is well documented on the Github Governance.md page and there is a correspoding link 
on the website to access the governance model. https://github.com/openebs/community/blob/develop/GOVERNANCE.md
- Meetings are documented/distributed well via youtube/ meeting minutes and slack  https://github.com/openebs/community/discussions/categories/weekly-maintainers-meetings 
https://www.youtube.com/@openebscommunity6021




### Areas for Improvement

**Over the next year, the project should work on the following issues to improve
its governance, these are considered non-blocking:**

There are few areas that may improve governance and general discoverability:

- the creation of a community member list that clearly outlines on one page members/contributors/committees for OpenEBS and the sub-projects (expalined in detail below).
- All CONTRIBUTING.md files should be cleaned up a bit to include the same level of detail as the primary project.
- The contributor pipeline needs to be expanded to grow the diversity of contributors and reach of OpenEBS globally. Cultivating a more diverse contributor pool could help with the recruitment and retainment of future maintainers from various companies globally.
- Link to decisions related to policies and governance on the Governance.md page does not work and needs to be fixed https://github.com/openebs/community/blob/HEAD/GOVERNANCE.md#changes-to-governance-and-policies
- Fix the governance link on the website - under contributing https://openebs.io/community it currently provides a no server error. Fix this or link it directly to the github governance page (check to ensure it is accessible on various browsers not just chrome)

---

## Review

<!--
Khallai Comment: When linking to files in GitHub, reference them with a specific commit.
a Governance review is a point-in-time review and things can change. Referencing
the commits will ensure the links will still work, and show what the project
was like at the time of the review.
-->

**The following review primarily consists of an audit on the project's
self-assessment in their Incubation application.**


### Governance Summary

- OpenEBS goverance is modeled after the Maitainer Council. With maintainers for the core OpenEBS
project and special maintainers for the sub-projects.
- OpenEBS umbrella project acts as a top level governing body
  and final escalation for the five sub-projects within Open EBS. 
- For a decision to be approved there must be 66% (super majority) vote by maintainers.
- There is a process to promote and demote maintainers and special maintainers. Example of the most recent https://github.com/openebs/community/pull/174
- OpenEBS has a weekly maintainer meeting with well documented minutes https://github.com/openebs/community/discussions/categories/weekly-maintainers-meetings


Governance Maintainer model: https://github.com/openebs/community/blob/HEAD/GOVERNANCE.md



### Governance Evolution

**Governance has continuously been iterated upon by the project as a result of
their experience applying it, with the governance history demonstrating
evolution of maturity alongside the project's maturity evolution.**
<br>
**Incubating:** Suggested | **Graduated:** Suggested

Over the past years OpenEBS has made a concerted effort of documenting their governance evolution through:


- April 19, 2024, governance document updated for sandbox application [(https://github.com/openebs/community/commit/056895c0f7f965d5ac07dd2f47566ae76fbfa9c6]
- May 1, 2024, steering committee role removed [https://github.com/openebs/community/pull/39]
- July 9, 2024 rewrites the processes and definitions of the contributors, especially maintainers. [https://github.com/openebs/community/pull/73]
- August 5, 2024 maintainer updates and process [https://github.com/openebs/community/pull/78]
- October 5, 2024 clarification of maintainer role and admim rights [https://github.com/openebs/community/pull/93]
- March 6, 2025, governance document updated for incubation application [(https://github.com/cncf/toc/issues/1601)]


### Discoverability

**Clear and discoverable project governance documentation.**
<br>
**Incubating:** Suggested | **Graduated:** Required

OpenEBS governance documentation can be found in the project's community repo https://github.com/openebs/community/blob/HEAD/GOVERNANCE.md#changes-to-governance-and-policies

The website for OpenEBS allows for the governance information to be accessed from the community page under 
contributing https://openebs.io/community


<!--
Khallai comment: governance is clearly discoverable.
It would be helpful if the sub-projects has well documented and clearly linked governance
pages for the main project and each sub-project.

-->

[community repo]: https://github.com/openebs/community
[GitHub Org]: https://github.com/openebs
[contributor guide]: https://github.com/openebs/community/blob/HEAD/CONTRIBUTING.md

### Accuracy and Clarity

**Governance is up to date with actual project activities, including any
meetings, elections, leadership, or approval processes.**
<br>
**Incubating:** Suggested | **Graduated:** Required

It is up to date and functioning as written. Some examples of where the
information can be found to verify:

- Their meetings are listed [through multiple calendar options]
- Recordings of their meetings can be found on the youtube channel https://www.youtube.com/@openebscommunity6021, and the zoom
  recording links are public.

[Maintainers Meetings]: https://github.com/openebs/community/discussions/categories/weekly-maintainers-meetings

[Monthly Meetings Youtube]:https://www.youtube.com/@openebscommunity6021


**Governance clearly documents [vendor-neutrality] of project direction.**
<br>
**Incubating:** Suggested | **Graduated:** Required

- OpenEBS does not have good vendor-neutral practices. Out of the five core maintainers four are from 
the same organization. https://github.com/openebs/community/blob/develop/MAINTAINERS.md Additionally, the website nor files clearly state how OpenEBS is vendor neutral for the purposes of adoption and marketibility. 
- Meetings are open and listed on the public calendar https://zoom-lfx.platform.linuxfoundation.org/meetings/openebs?view=list and the zoom link is available https://zoom-lfx.platform.linuxfoundation.org/meeting/99475226462?password=f57455ef-fed1-49b1-a38b-110dab2b42cc
- The vision.md page references vendor adoption which include Microsoft Azure, DataCore, and Civo https://github.com/openebs/community/blob/develop/VISION.md


<!-- 
Khallai Comment: Cultivating a more diverse contributor pool could help with the recruitment and retainment of future maintainers from various companies globally. This area needs to be worked on to graduate but is not required for this incbation application. For marketibility and vendor neutrality having a statement on the website and within one of the core docs would work well for the graduation application phase. There is a blub in the incubation application regarding this but its generic and links back to the CNCF site and does not demonstrate how the project puts vendor neutral into practice. 
-->


### Decisions and Role Assignments

**Document how the project makes decisions on leadership roles, contribution
acceptance, requests to the CNCF, and changes to governance or project goals.**
<br>
**Incubating:** Suggested | **Graduated:** Required

- OpenEBS follows the CNCF Code of Conduct
- The Roadmap and milestones are available https://github.com/openebs/community/blob/develop/ROADMAP.md and repo https://github.com/orgs/openebs/projects/78 
- The maintainer election and/or demotion is outlined in the governance.md page https://github.com/openebs/community/blob/develop/GOVERNANCE.md with clear details on vote percentage required and estimated timeline. 


<!-- 
Khallai Comment: Thia area of decisions and role will likely expand as the project approaches graduation. For the incubation application the current leve of detail is sufficient.
-->


**Document how role, function-based members, or sub-teams are assigned,
onboarded, and removed for specific teams (example: Security Response
Committee).**
<br>
**Incubating:** Suggested | **Graduated:** Required

Most of the roles, qualifications and the onboarding/offboarding guidelines are
in OpenEBS umbrella project governance page of the governance.md doc.

### Maintainers and Maintainer Lifecycle

**Document a complete maintainer lifecycle process (including roles, onboarding,
offboarding, and emeritus status).**
<br>
**Incubating:** Suggested | **Graduated:** Required

The full maintainer lifecycle is documented in their [text](OpenEBS/GOVERNANCE.md) and [text](OpenEBS/MAINTAINERS.md) files.

A  copy of their membership table is copied here for convenience.
<BR>

| Level | Role | Responsibilities |
| :---  | :--- | :--- |
| 1     | **Maintainer** | Vote, Develop roadmap, contribution guidelines; Review, Approve/Reject, Merge, Manage repos. Maintainers are elected or removed by the current maintainers. A Maintainer has authority over the OpenEBS umbrella project: the organization and every project, sub-project and repo within the organization.|
| 2     | **Special**<BR>**Maintainer**| Have special expertise in a particular domain within the OpenEBS umbrella project. The domain may be a sub-project, repo or other responsibility as defined by the Maintainers. The maintainers grant a special maintainer a set of authorities and responsibilities for the domain. A special maintainer is expected to join maintainer and community meetings when required. A special maintainer has no responsibilities for the umbrella project, or projects outside their domain.|
| 3     | **Contributor** | Contribute code, test, document the project. A contributor’s authority applies to one or more sub projects. |
| 4     | **Adopter** | Use the OpenEBS product, with or without contributing to the project. An adopter has authority to raise issues, participate in discussions on sub projects within a public forum. |

<BR>


**Demonstrate usage of the maintainer lifecycle with outcomes, either through
the addition or replacement of maintainers as project events have required.**
<br>
**Incubating:** Suggested | **Graduated:** Required

The community repo, [text](OpenEBS/GOVERNANCE.md) and [text](OpenEBS/MAINTAINERS.md) files contain the necessary information and examples of 
the maintainer lifecycle for OpenEBS (umbrella projects).

[commit history]: OpenEBS and its subprojects have a solid and well structured commit process. Each subproject has a dedicated maintainer which is helpful to ensure that PRs are not missed and handled in a timely manner. All projects provide a CONTRIBUTING.md file that provide details on how to contribute to the project. 

Release notes for OpenEBS and its subprojects are well structured and documented https://github.com/openebs/openebs/releases sections like bug fixes and improvements clearly show milestoes and acheivements that the project has accomplished.

<!-- 
Khallai Comment: I would recommend that the rawfile-localpv CONTRIBUTING.md file be cleaned up a bit to include the same level of detail as the other projects (it has not been updated in 6 months). For example include sections on signing your work, pull request checklist etc. The same recommendation for Mayastor https://github.com/openebs/mayastor/blob/develop/CONTRIBUTING.md it is important to direct people to the right resources and repos so that contributing to the subprojects is the same effort as contributing to the main OpenEBS project https://github.com/openebs/rawfile-localpv/blob/develop/CONTRIBUTING.md
-->


**Document complete list of current maintainers, including names, contact
information, domain of responsibility, and affiliation.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS lists their members, affiliation, and responsibilities in their
[text](OpenEBS/MAINTAINERS.md).


**A number of active maintainers which is appropriate to the size and scope of
the project.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS has five active maintainers, with an
[individual contributor dependency] of 3 active contributors authoring 57% of the
commits in the past year.

[individual contributor dependency]:https://insights.linuxfoundation.org/project/openebs/contributors?timeRange=past365days&start=2025-02-09&end=2026-02-09

**Project maintainers from at least 2 organizations that demonstrates
survivability.**
<br>
**Incubating:** N/A | **Graduated:** Required

Not required for incubating application

### Ownership

**Code and Doc ownership in Github and elsewhere matches documented governance
roles.**
<br>
**Incubating:** Required | **Graduated:** Required

From a random sampling of 5 repos, they all align with what has been
documented.

### Code of Conduct

**Document adoption and adherence to the CNCF Code of Conduct or the project's
CoC which is based off the CNCF CoC and not in conflict with it.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS has [adopted the CNCF Code of Conduct], and documents how to report CoC
violations within the [text](OpenEBS/CONTRIBUTING.md) doc.

[adopted the CNCF Code of Conduct]: https://github.com/openebs/community/blob/1d4a43f54b5ed1f43537197987e7a00b353107e8/CODE_OF_CONDUCT.md


**CNCF Code of Conduct is cross-linked from other governance documents.**
<br>
**Incubating:** Required | **Graduated:** Required

The [CoC is cross-linked] from their governance docs.


### Subprojects

**All subprojects, if any, are listed.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS is an umbrella project with five federated subprojects. This is reference in the [text](OpenEBS/README-2.md) and [text](OpenEBS/CONTRIBUTING.md) files. The centralized location for files and project information can be found in the community repository. 

OpenEBS has a well structured roadmap https://github.com/orgs/openebs/projects/78 that tracks the progress of the project and subprojects. Within it milestones can be seen along with release overview and ant the status of the pull request.

 
**If the project has subprojects: subproject leadership, contribution, maturity
status documented, including add/remove process.**
<br>
**Incubating:** Suggested | **Graduated:** Required

OpenEBS lists their subprojects in the [text](OpenEBS/README-2.md) 

The subprojects of OpenEBS are:
Dynamic LocalPV Provisioner - https://github.com/openebs/dynamic-localpv-provisioner/
Local PV ZFS  - https://github.com/openebs/zfs-localpv
Local PV LVM -  https://github.com/openebs/lvm-localpv
Local PV RawFile - https://github.com/openebs/rawfile-localpv
Mayastor  - https://github.com/openebs/mayastor


OpenEBS subprojects follow the project level guidelines for membership,
contributing etc, but MAY define additional criteria or roles to meet the
needs of the subproject as the project matures.



### Contributors and Community

**Contributor ladder with multiple roles for contributors.**
<br>
**Incubating:** Suggested | **Graduated:** Suggested

[contributor ladder]: https://insights.linuxfoundation.org/project/openebs/contributors?timeRange=past365days&start=2025-02-09&end=2026-02-09

<!-- 
Khallai Comment: the creation of a community member list that clearly outlines on one page members/contributors/committees for OpenEBS and the umbrella projects would  be helpful to determine who is actively involved in the project(s). A community-members.md file would be sufficiant and could be orhganized as the team sees fit. The current README.md has a brief summary of community but does not list any of the people involved along with their function.
-->



**Clearly defined and discoverable process to submit issues or changes.**
<br>
**Incubating:** Required | **Graduated:** Required

From a random sampling of 5 repos, they all had issue templates for the general
set of things like features, bugs, etc.

- the community repo in the contributing.md file is where issue process lives https://github.com/openebs/community/blob/develop/CONTRIBUTING.md


**Project must have, and document, at least one public communications channel
for users and/or contributors.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS has all their public communication channels available in the
community/repo.


**List and document all project communication channels, including subprojects
(mail list/slack/etc.). List any non-public communications channels and what
their special purpose is.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS has all their public communication channels available in the
community/repo.

**Up-to-date public meeting schedulers and/or integration with CNCF calendar.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS lists public meetings on the community github https://github.com/openebs/#community page and the meeting calendar is integrated with cncf/lf page https://zoom-lfx.platform.linuxfoundation.org/meeting/99475226462?password=f57455ef-fed1-49b1-a38b-110dab2b42cc 
https://zoom-lfx.platform.linuxfoundation.org/meetings/openebs?view=list
[communication channels] meeting recordings are made available via youtube - https://www.youtube.com/@openebscommunity6021


**Documentation of how to contribute, with increasing detail as the project
matures.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS has documentation for contributors both within their community repo and
published on their website. They have documentation targeted towards
maintainers in their community repo, and these are supported by additional
documentation (e.g. contributing.md) in each repo https://github.com/openebs/community/blob/develop/CONTRIBUTING.md

<!-- 
Khallai Comment: the contributor.md does not contain the same detail and content as the primary OpbenEBS contributor.md file. Suggestion would be to add appropriate detail to all sub-projects so that future contributors know the process for a PR and other core requirements for contributing to OpenEBS and/or the subprojects.
-->

**Demonstrate contributor activity and recruitment.**
<br>
**Incubating:** Required | **Graduated:** Required

OpenEBS has a [growing contributor community]
[growing contributor community]: https://insights.linuxfoundation.org/project/openebs/repository/openebs_community/contributors?timeRange=past365days&start=2025-02-08&end=2026-02-08
[OpeEBS incubation application]: https://github.com/cncf/toc/issues/1537
[project milestone or other requirement]: https://github.com/cncf/toc/tree/main/process#how-to-apply-to-move-levels
[vendor-neutrality]: https://contribute.cncf.io/maintainers/community/vendor-neutrality/
[communication channels]:https://github.com/openebs/community/blob/1d4a43f54b5ed1f43537197987e7a00b353107e8/README.md
[Sub-projects]: https://github.com/openebs



