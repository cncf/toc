# Project Health Reviews

The TOC may choose to review a project when a concern is registered by a community or TAG member or as a result of a health indicator, project event, or other occasion. These reviews are typically specific to the occassion that warranted them but involves researching the problem or indicator, understanding if the metric or occasion is accurate in its depiction and warrants attention, action, or other next step.

When a project health concern is raised, the following must occur:
- A [Project Health issue](https://github.com/cncf/toc/issues/new/choose) is created on the TOC repo
- A TOC member is assigned to coordinate and orchestrate resolution with the project and TOC
- The assigned TOC member engages with the project by filing an issue on their repo regarding the concern, linked to the TOC issue

It is imperative that the TOC maintain documentation regarding engagements with projects to include direct messages, exchanges, and conversations with individuals. Capturing these notes, discussions, and dates is key in the event of TOC member turnover, long-standing or sporadic issues with the same project, or if additional TOC members are needed to weigh in and understand what is occurring. We strive to maintain this documentation in the open, however there are occasions where the TOC may annotate these events in private due to the sensitive nature of items under discussion. Such sensitive matters may include: lay-offs, acquisitions, or other business/proprietary decisions impacting the project.

## Initial outreach

Once an issue is filed with the project regarding concerns, this initiates the first outreach to the project. Occasionally health concerns are miscommunications or lack of clarity issues that just need sorted with the project and don't involve much beyond that. The TOC expects projects to respond to the issue files within two (2) months of filing and the TOC member MUST provide the project with a date. 

For example:
> Related: https://github.com/cncf/toc/pull/1119/files. & #147
> I am the TOC member assigned to conduct the Evaluation of $PROJECT's Health Review. At this time the TOC is requesting an update from the project regarding <list concerns> with expected resolution. We ask the project to address these concerns and provide anticipated dates of resolution. **A response is due to the TOC by January 31st 2024** and may be provided as a comment on this issue (link to TOC issue) within the TOC repo and tagging myself, @TheFoxAtWork.

The nature of the response determines next steps.

### No response

In the event that after two months you do not receive a response, you will need to conduct a health and alignment review unless it is obvious the project is abandoned. If the project is abandoned, the TOC member should propose the project for [archive](../process/archiving.md).

The health and alignment review focuses on the following areas:
- **Project scope and goals remain consistent.**
 - What to expect: The project goals and scope have not deviated greatly from when the project was accepted. Deviations are publicly documented but generally related to the original conditions under which it was accepted or elevated. i.e. a detection engineering project has not suddenly become a data plane.
 - Verify: Review the information about what the project is/does. Check if the project has changed in scope and direction from the original application, and/or due diligence. Also check if the project has had any major architectural shifts that are within reason.
- **Community development and growth is ongoing.**
 - What to expect: The project may hold meetings, discussions, and/or other forms of community-focused communication. The frequency of these is set by the project and should be open to the public. It is expected the project is balancing community growth with project development.
 - Verify: Check out the Project’s docs for information of meetings or community activities such as presence on CNCF calendar for metrics, meeting times/location in the project’s repo or other equivalent communication. Check if there is information on setting up development environments, this may be an indicator of progress towards incubation. Additionally, check if the number of contributors and number of contributing organizations is increasing. Review whether new contributors are core or transient and capture context. Also take note of directed efforts to grow the community. 
- **Appropriate project governance.**
 - What to expect: The project has clearly defined governance commensurate with their maturity level that addresses the core areas needed to successfully operate the project. The project will likely have a minimal history of changes to the governance beyond initial introduction, as they may be too early to fully exercise their governance process (such as maintainer emeritus changes). This governance should capture the basics of decision making, particularly as it pertains to PR reviews and how contributors may participate in the project. More mature projects should have governance that is comprehensive and minimally meets the criteria for the current level of the project.
 - Verify:  Look for changes to governance.md, readme.md, contributing.md or other governance files - such changes may be refining the governance structure, its processes, or even defining project principles, determine if the CNCF CoC is present (either through a link or as a dedicated file - and is up to date). Also verify that the governance is enforced - for example, ensure that the projects don’t merge changes without review, add maintainers without using their process, etc. If any major concerns or issues were presented, check if they were followed up on.
- **Long-term planning exists.**
 - What to expect: The project defines long term goals, objectives, features, or releases. It is not expected the project will achieve every defined goal or objective within the timeframe or release they specify, rather that they are considering the future of the project and tracking completion of those activities. This may be extremely lightweight in the form of identifying primary components in a v1.0 release, or it may be more detailed and use project boards, milestones, and other planning flows.
 - Verify: Look for the presence of an up-to-date roadmap file (determined by the timeframe the roadmap represents), project goals, project boards, etc
- **Development is ongoing and progressive towards versioned releases.**
 - What to expect: Early projects are likely to have significant variability in the frequency by which development happens. They may experience bursts of contributions and commits, slow steady pacing, or a combination thereof. Each project is different, however we should see some form of progress which may be features and bug fixes, stability improvements, project packaging infrastructure, and tests. Projects should not appear to be in maintenance or sustainment mode - this is an indicator of a project health concern.
 - Verify: Review the Overall-project-statistics-table on https://devstats.cncf.io/ for the project, if last year’s overall project statistics table is available, compare. Look to see if a versioned release is present. Additionally, look to see if they are following semantic versioning, are creating new tags on major feature development and there is evidence of past stable released versions (semantic versioning is a more mature requirements so verify the project's level).
- **Adoption is picking up.**
 - What to expect: While the project is early, we may see one or two more adopters, users, or entities exploring the viability of the project for their needs. We’re not likely to see significant changes within their adoption file, however indicators of features requests and indicators of positive attention by potential adopters. More mature projects should have increasing adopters and these may be recorded in the ADOPTERS.md file or equivalent.
 - Verify: Look for changes to the adopter file, organization/company blogs, articles and talks concerning the project, look for issues marked as feature or other labels from non-maintainers. Determine if existing adopters are continuing to engage.
- **Self-awareness.**
 - What to expect: The project has reasonable awareness of their overall maturity and progress to towards the next level (or maintaining their graduation). They understand their path to the next level and any activities still outstanding in order to be ready to apply. If the project is a graduated project, they are aware of the necessary activities in maintaining good standing with their adopters and community.
 - Verify: Look for their readiness to move to levels or indicators of iterative improvements, look to see if any blockers or other activities are identified, if the project has regressed, or if sub-projects are non-conforming.

The assigned TOC member should provide a brief statement summarizing their review of each area in the context of the project. The TOC member's summary should be included on the TOC issue and linked as a comment to the project's issue. The TOC member may then allow the project another 30-60 days to respond. If the project is still unresponsive, the TOC member should propose the project for archive.

### Project responded

#### Undisputed

The project responded and is not disputing the concerns. They've provided anticipated resolution dates. When this occurs, simply follow up with the project to ensure the work is closed out.

#### Disputed

In the event that the concerns are disputed, the TOC member will need to engage directly with the project to bring closure/resolution to the concerns. This may require identifying specific issues, providing recommendations, and requesting commitment from the project (typically via Roadmap) to close out those items. Alternatively, the TOC may engage the corresponding TAG to work with the project on identifying issues and recommendations. In all matters, the path forward must be documented on the TOC issue and followed up to ensure the matter is closed.

#### Additional Notes

There may be occasions where a project is continually being reviewed for health concerns and not making substantive progress in resolving them. It is important to note that the CNCF exists for the benefits of its members. It is those members, through their elected or named individuals who define the conditions and criteria for projects to receive the benefits of the Foundation. If a project chooses to no longer meet the conditions and criteria, or apply the CNCF values, the TOC may choose to archive the project.
