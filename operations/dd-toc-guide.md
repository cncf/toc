# TOC Due Diligence guide

This document provides the TOC with guidance on how to execute due diligence of CNCF projects for each level of maturity.  It complements the Moving Levels process detailed in the [Process Directory](../process/README.md).

## Quick Links

Getting Started: **[Triage applications](#initial-triageevaluation-prior-to-assignment)** | **[TOC Assignment](#toc-members-step-forward-to-be-assigned)** | **[Kick-off](#kicking-off-the-due-diligence)**

The Due Diligence (DD):  **[Due Diligence](#completing-due-diligence)** | **[Finalizing DD](#finalizing-the-due-diligence)** | **[Adopter Interviews](#conducting-adopter-interviews)**

Wrapping it up: **[TOC internal review](#toc-internal-review)** | **[Public Comment](#public-comment-period)** | **[Voting](#voting)**

## What is due diligence?

Due diligence is the process by which the TOC performs an independent review of CNCF projects to assess their posture, maturity, and adoption across a variety of technical, governance, and community focuses. The intent of the due diligence is to provide project and adopters with an independent senior technical evaluation of a project's readiness for production. Similar to how organizations have established software development processes and check points prior to software delivery or deployment that ensure the software meets the organization's goals and objectives, the due diligence is a point in time artifact of the project's achievement for meeting the goals and objectives expected for their maturity level. By performing the due diligence on CNCF projects, the TOC supports our adopters in gaining confidence that the project has been reviewed against documented criteria for their maturity level, can understand any deviations from those criteria, and may not need to repeat this type of evaluation, rather they may incorporate or leverage the contents of the due diligence to guide and information their decision making. It also conveys to adopters the potential level of effort in adopting and integrating the project into their architecture and infrastructure.  For projects, the due diligence provides an evaluation of the project against the expectations of adopters across multiple focuses. It can and often will include additional recommendations to the project that may support them in reaching the next level of maturity, improving their documentation or architecture, and in some cases highlight opportunities to distinguish themselves and their features or functionality from other projects within the same area.

Due Diligence is typically conducted when a project has applied to move levels, after the project completes and asserts their adherence to the criteria for the level they are seeking to move to. For example, an incubating project must complete the graduation criteria before applying to graduate. Part of the TOC's responsibility is to not only independently verify those assertions, but to understand those assertions in the context of the individual project, its domain, its interoperability, and subsequent alignment with the Foundation, ecosystem, adopters, and overall industry.

The criteria for moving levels sets the expectations for what projects are to complete prior to applying for the next level. Criteria are outcome oriented and have been developed to provide projects with maximum flexibility in implementation while demonstrating the characteristics and hallmarks of maturing projects. TOC members may exercise their discretion in evaluating conformance to the criteria, however deviations, waivers, and exceptions should be clearly articulated and reasoned such that the desired outcome is still met through compensating mechanisms or demonstrated to not be relevant.

The process of conducting due diligence is not to create unnecessary work for projects, rather to ensure projects are set up to succeed in a way that works for them against the values, principles, and definitions of cloud native and the expectations adopters carry for scalable, secure, production-ready software.

It is important to remember that the CNCF is for the benefit of its members. This means the CNCF, through the TOC, defines the technical and community expectations for projects to receive benefits of the foundation. The criteria, due diligence process, security audits, templates, code of conduct, and community management practices are several expectations for projects to receive benefits. Adherence to these, and confirmation of adherence for each project's maturity level is essential to continuously receive those benefits. 

As a matter of balance, it is important to note that while these benefits provide value to projects, the TOC is responsible for continuously assessing the criteria, due diligence process, and overall levels framework to ensure we are adapting to shifts and changes in the landscape and market, as well as responding to the needs and limitations experienced by projects. _The community and projects are our partners in effecting that change, so feedback is crucial._

## How to conduct due diligence

We will not focus on the events that have transpired prior to a project's application to move levels, choosing instead to focus on the point when the TOC is engaged.

Currently, there exist three levels in the CNCF:
* _Sandbox_ - experimental and early stage projects worth exploring their viability and adoption probability within the ecosystem
  * _Sandbox_ projects undergo due diligence when they apply to _Incubation_.
  * _Sandbox_ projects do not receive due diligence in order to join the Foundation.
* _Incubating_ - projects that have proven concepts, are receiving market and adopter attention, and increasing their robustness, security, scalability, and maturity
  * _Incubation_ projects receive due diligence when they apply for _Graduation_.
* _Graduated_ - projects that are stable, resilient, secure, and highly mature, they are capable of meeting the expectations of production-ready software
  * _Graduated_ projects do not receive additional due diligence once they reach this level.
  * By the time projects have _graduated_, they will have undergone two rounds of due diligence, once during the application to _incubation_, and a second time when they apply to _graduate_ - often a refresh depending on age since the last due diligence.

Projects not already in the Foundation may apply directly to Incubation if they feel they are ready. At this point of application, they undergo due diligence that also considers their fit in the CNCF, much the same as the considerations for inclusion of sandbox projects.

It is critical that TOC members strive to complete due diligence in a reasonable amount of time. The process involves multi-parties that requires coordination and clear communication of expectations. Delays in completing due diligence can create friction with projects and may encounter TOC member term endings, requiring project to start fresh with a new member. It is expected the process will take time, adopter interviews may be difficult to schedule in a timely fashion, so being upfront on these expectations is important. **Each TOC member is expected to manage their project's application timeline to the best of their ability and reduce any delays where possible.**

Every TOC member is expected to conduct due diligence of CNCF projects and triage those applications. In cases where there may be a perceived [conflict of interest](#conflicts-of-interest), the due diligence must have two TOC members participating in order to dissolve any illusion of bias (for or against).

TOC members may not take on anymore than two (2) projects for due diligence at a given time unless one of the following conditions is true:
* the TOC member is functioning as a guide to new TOC members learning this process
* the TOC member is serving as a secondary to the primary TOC member conducting a due diligence in order to offer additional domain support or other relevant subject matter expertise
* the TOC member has two projects in voting
* the TOC member has one project in voting, and another in progress

We expect all TOC members to be mindful of their obligations and timelines, whether they be work, cloud native, or personal and manage their workload accordingly.

### Project has applied to move levels

**Projects apply to move levels by submitting an Issue on the TOC repo that details their conformance to the criteria.** These Issues go into the [TOC backlog](https://github.com/orgs/cncf/projects/27/views/9) and are worked on a first-in, first-out/started basis to the extent that is reasonably practical - each TOC member has subject matter expertise in one domain or another and each TOC member serves as a liaison to a Technical Advisory Group that may or may not align with that domain. The TOC strives to align their skills and background with the project applying they are assigned to ensure they have the context to understand the project within that domain. 

The issue will contain a limited set of information about the project at the time of its application, commonly asserting its conformance to the stated criteria with links to where or descriptions as to how they are implemented.

### Initial triage/evaluation prior to assignment

All TOC members are expected to assist in the triaging of project applications to move levels to ensure that when a TOC member is ready to be assigned, the project is ready to be evaluated with no immediate blockers that would inhibit or delay the TOC's engagement.

This light-weight triage/evaluation must cover the list below (it is not exhaustive and is a minimum triage set from the [incubation template retrieved 12 DEC 2025](https://github.com/cncf/toc/blob/c2943ffc98064dd88e9ef9c4afd5a8856898942f/.github/ISSUE_TEMPLATE/template-incubation-application.md)):
* Adoption Assertion includes the Adopters file link, and the project has an entry in the Adopter's form responses to provide 5-7 adopters to reach out to. [Check here](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for the definition of an Adopter.
* Application Process Principles provides 
  * Either: the link to the Recording, issue, and/or meeting notes from a TAG meeting where the project presented with the domain specific TAG -or- completion of the [General Technical Review (GTR)](../tags/resources/toc-supporting-guides/general-technical-questions.md) or [Domain Technical Review (DTR)](../tags/resources/toc-supporting-guides/tag-domain-technical-review-template.md) in lieu of a TAG meeting
  * assertion of vendor neutrality
  * assertion of review and acknowledge of expectations of CNCF projects and requirements for moving forward through the CNCF maturity levels
  * provided any additional documentation links the project feels is appropriate for its type
* Governance and Maintainers provides
  * link to the project's governance
  * any notes on governance iteration
  * Maintainers file is linked
  * lists number of active maintainers
  * link to or information regarding doc and code ownership
  * link to the project's code of conduct (should link to CNCF CoC)
  * link to CoC references in Governance docs (can be a link to governance only and we recommend linking in contributions as well)
  * link to subproject listing, if applicable.
* Contributors and Community provides
  * link to contributing file or other file that describes issues or change submissions (i.e. enhancement proposal process)
  * link to file containing project communication channels
  * links to information on community meetings, recordings, and/or notes
  * link to the contributing file
  * information on active contributors (i.e. quantity, contribution metrics, etc.) and documented efforts to attract more contributors (i.e. issues, presenting at conferences, slack messages, mailing lists, etc.)
* Engineering Principles
  * link to information on the project goals and objectives, scan for the use cases or problems the project addresses
  * link to information on supported use cases, what the project does, etc.
  * link to roadmap, project board, or milestones
  * link to project architecture diagram and documentation
  * link to project release process
* Security
  * link to joint assessment if available
  * link to security.md, scan for a process to report issues
  * link to (scorecard: scan for branch protection, token permissions, SAST, and CI best practices in results) or link to other evidence of repo hardening
  * link to security report resolution process and roles
  * link to completed or PR filed self-assessment
  * link to best practices badge, confirm it is "passing" and 100% complete 
* Ecosystem
  * link to adopters file
  * link to integrations/ compatibility information of other projects and products

Projects should NOT be blocked if they do not have a Governance Review or a GTR/DTR completed. Both the Governance review and GTR/DTRs depend on the availability of our community members in our TAGs which cannot be guaranteed.

If some of the criteria are not yet met, or missing, the TOC member triaging will add a comment detailing all items that are unmet or missing and close the application; affixing the "not-ready" label and move the card to the "Not-Ready-Will Return" column of the [TOC project board](https://github.com/orgs/cncf/projects/27/views/9)'s Applications to Move levels tab. Projects are expected to re-apply upon completion of outstanding items. When the project is ready to reapply, they should link to the previous application so the TOC may leverage and reuse as much prior work as reasonable.

Once the TOC has triaged the application and found all criteria to have content, the TOC member performing triage comment the application is complete and ready for assignment. The TOC member will affix the "ready" label and move the project from the "new" column on the application's board. The project's application will be updated by the TOC member with a comment that details where work still needs to be completed, next steps associated with completion of those, and an estimated timeframe that the project is likely to complete those items by. 

TOC members are expected to triage projects in the "new" column on the board for projects that are returning from a previous not-ready application, verify completion, and move them to the top of the ready for assignment column. 

TOC members are to prioritize selecting projects from the ready for assignment column over the new column to expedite the queue and make the best use of TOC time.

### TOC member(s) step forward to be assigned

Commonly referred to as the Project's Application Sponsor, TOC members assign themselves to projects to sponsor the application for moving levels. Sponsoring an application ensures a focused point of contact exists for both the project and the TOC in completing the Due Diligence, public comment, and execution of voting. 

The TOC member that assigns themselves a project to sponsor the application for moving levels may request a secondary TOC member to support the due diligence according to eligibility.

TOC members ready to perform due diligence a project's application will socialize this internally with the TOC to provide opportunity for other members to participate. Once a TOC member or members is determined, those TOC members must assign themselves to the Issue and move the issue's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Assigned".

TOC members may occasionally be contacted by projects asking the TOC member to sponsor their application to move levels. When this occurs, politely reminder projects of the TOC's first-in, first-started process and let them know it is subject to TOC availability. Follow up internally to the TOC that you were contacted and by which project so we may all be aware of which projects are reaching out as it may identify a backlog, timing, or communication issue.

#### Conflicts of interest

The Due Diligence process provides several mechanisms that serve as check points for bias or uncomprehensive evaluation. The public comment period is the most open and transparent point at which the community and others may independently review the diligence performed by the TOC.  Additionally, TOC members conduct an internal review of the due diligence to verify all aspects of the project have been thoroughly reviewed.

A TOC member will require a co-sponsor for a project if any of the following conflicts apply:
* Hard Conflicts
  * a project maintainer,
  * direct report to a current maintainer,
  * paid to work on the project
  * has a significant financial interest directly tied to the success of the project,
* Soft Conflicts
  * belongs to the same organization of the project,
  * or uses it in their work

 This does not mean they can't have any involvement with a project at all as contributing to pull requests or adopting the project are signals of a healthy interest and knowledge of the project. To ensure appropriate evaluation without bias, a second, unconflicted TOC member must be assigned to co-sponsor the project with them.

If a conflict of interest is present, the TOC member will state they have a conflict and seek a second _primary_ sponsor on the project's application issue prior to proceeding. 

### Kicking off the due diligence

Once the project is assigned to the TOC member(s), the TOC member(s) engages the project's maintainers or leadership group to kick off the due diligence. This can be done by commenting on the Issue, starting a slack channel (public or private), slack group direct message, email exchange, or thread in the project's primary communication channel. We strive to keep these discussions transparent and visible but should consider any potential sensitive issues that come about as a result of the review (resolution of vulnerabilities for instance). 

Any form of communication must include _two members_ of CNCF staff to ensure consistency and continuity throughout the process. 

TOC members, with support from CNCF staff, should schedule a meeting with the project to the extent possible given availability and timezones. Asynchronous kick-offs can occur, but may result in additional back and forth or delays. Each Kick-off meeting should have a central kick-off document that allows the TOC and the project to capture expectations, decisions, timelines, and other pertinent references needed for successful completion of the due diligence. A [kick-off meeting template](toc-templates/template-kickoff-notes.md) is located in the [toc-templates](toc-templates/) folder.

Once the Kick-off is scheduled, the TOC member(s) should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "In Due Diligence". It is *strongly* recommended that you inform the project to verify compliance with the CNCF's licensing policy (set forth in the [Section 11 of the Charter](https://github.com/cncf/foundation/blob/master/charter.md) with [additional information here](https://github.com/cncf/foundation/blob/main/allowed-third-party-license-policy.md)).

### Completing Due Diligence

NOTE: the Due Diligence can be completed in conjunction with adopter interviews, some TOC members find completion of the DD as informative to conducting interviews, but not in all cases.

The TOC will use the appropriate Due Diligence for the project's applied level as the basis for a PR ([Incubation template](toc-templates/template-dd-pr-incubation.md), [Graduation template](toc-templates/template-dd-pr-graduation.md)) and will evaluate the project's assertions in the application issue against the discoverable and publicly available sites, repos, files, and other artifacts of the project. The TOC's evaluation against each criteria goes in the corresponding area of the PR template.

Previously, the project was responsible for completing the due diligence such that it allowed the TOC member(s) to review and comment. Due to the extensive back and forth in this prior process and with recent changes to the criteria, the TOC has altered the process leverage a Due Diligence PR as the TOC's assessment of the projects completion of the criteria. Therefore TOC members are expected to complete the Due Diligence PR with support from the project and TAG(s).

As the TOC member(s) reviews the criteria, any deviations or implementation notes from the review should be recorded within the due diligence PR as part of their evaluation for that specific criteria. The TOC member will provide an overall evaluation statement that summarizes the content of the due diligence once the Adopter Interview Summaries are recorded. For more information on adopter interviews, refer to [Conducting Adopter Interviews](#conducting-adopter-interviews).

As an example, let's say a project is asserting their sub-projects have leadership, contribution information, maturity statuses, and add/removal processes. The TOC member(s), in confirming this may determine that the sub-projects share the same contributor file on the org's evolution directory and may change the maturity level in accordance with their documented process, but the process is not clear as to what initiates that change or who. For a sandbox project seeking incubation, this may be non-blocking but the TOC may recommend the project improve the documented process. As such, the TOC member(s) will record their finding and corresponding recommendation.

Another example, if the TOC member(s) is looking over the project's stated goals and objectives that have not changed since the project was accepted into the CNCF, and they are now applying to Graduate, the TOC member(s) will ask the project to clarify or provide additional information as to why the project hasn't initiated any changes and if they still feel those goals and objectives are accurate for the future of the project. The TOC member(s) will then summarize the response and record it in the PR under the corresponding criteria evaluation.

It is expected that the TOC's evaluation of a project's completion of the criteria may reveal a mismatch in understanding or an unexpected implementation. Documenting the TOC evaluation in the Due Diligence PR provides the project, TAGs, community, adopters, and TOC with a point of reference to understand if the criteria are meeting the outcomes required of a project for a certain maturity level, or if compensating mechanisms that supplement or augment the criteria are in place that work best for that specific project.

#### Sub-projects

When reviewing a project, you may find multiple sub-projects within their organization or repositories. It is important that the Due Diligence document reflect the scope of  evaluation for a given project and what that includes. Sub-projects are often broken out into two categories: packaged with the project, and independent but related to the project (SDK reference implementations for instance). 

Sub-projects packaged with the main project are within scope of the due diligence evaluation. Any sub-project not packaged is outside the scope of the due diligence evaluation and must be indicated in the due diligence as "un-evalauted" and the reason. TOC members must verify that the main project has sufficient governance of its sub-projects that allows adopters and contributors to understand the stability, maturity, and requirements or criteria to be a sub-project, their relation to the main project, and other information necessary commensurate with the maturity level of the project.

#### Licensing

In the course of conducting a project's due diligence, you may become aware of license issues, outstanding exception requests, or other matters concerning the licensing of the project and its dependencies. If in the course of Due Diligence you find the project appears to be missing a license exception for their code or dependencies (dynamic or statically linked), notify the project to [file a license exception request](https://github.com/cncf/foundation/issues/new/choose). The project's application to move levels may not go to a vote until the exception has been approved, however activities such as adopter interviews and public comment may proceed. It is expected that sandbox projects, when accepted, undergo a licensing review to ensure compliance with [CNCF's licensing policy](https://github.com/cncf/foundation/blob/main/license-exceptions/README.md) and exceptions. Refer to the [allowed third party licensing policy](https://github.com/cncf/foundation/blob/main/allowed-third-party-license-policy.md) for more information.

#### Graduated project security audits

TOC members who sponsor projects seeking graduation are expected to review the results of the audit to confirm the project has resolved all critical and high findings at a minimum. Additional findings are expected to be tracked for resolution by the project. In reviewing the audit of the project, you may find additional recommendations or deltas in the project's operational security (incident response, PR reviewer guidelines, lack of regression tests, etc.) that should be addressed. These are typically non-blocking, however the you should take care to note them within the due diligence after conveying their need and importance of completion to the project. You may reach out to TAG Security for assistance in identifying these areas or reach out to a TOC member with a background in Security.

#### Specification project

If a project is a specification project such as the TUF, SPIFFE and in-toto projects, there really is very little additional development that would need to happen and it is reasonable for the specification project to have only a few maintainers. For a specification project, it is required to have at least one implementation and that reference implementation DOES NOT need to be part of the project undergoing due diligence. The reference implementation should have  sufficient adoption to assess maturity and viability of the specification.  In many successful specifications, different adopters will implement their own copy for a variety of reasons. Each individual implementation may have limited diversity of maintainers, adoption, etc. but as a whole it can be broad and diverse.

### Finalizing the Due Diligence

When the TOC has finished their criteria evaluation, they should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Adopter Interviews & Project Discussion" and re-engage the project to elevate and discuss any items needing clarity, correction, or improvement. This includes notifying the project of any recommendations. Recommendations and discussion points may be copied into the kick-off document to facilitate discussion and to provide for additional context and discussion with the project until they are finalized.

The TOC member(s) may then file the PR and place it in draft until the Adopter Interviews are completed.

#### Engaging TAGs

If the project has previously engage with a TAG, the TOC may request an update or summary regarding the engagement and the TAG's assessment of the project in the context of the domain of the project. For instance, a policy-as-code project may present to TAG Security, TAG Security may summarize their recommendations to the project and link to supplemental material (like the project's self-assessment or joint-assessment) which is then added in the Due Diligence PR. 

Feedback by the TAG is encouraged prior to Due diligence being initiated, as requesting feedback when due diligence is occurring may prolong the due diligence process. It is expected that the project will have completed all criteria prior to applying, if this is not the case, due diligence should not commence.

### Conducting Adopter Interviews

In order to appropriately ascertain the adoption of a project, the TOC interviews a sampling of the project's adopters to understand how it is being used, what problems it is solving, the ease of adoption and integration, the community and contribution experience, and learn how adopters are experiencing the project's maturity level.

The TOC member(s) begin reaching out to the 5-7 potential adopters provided by the project to be interviewed and work with the TAB in conducting the interview. The TOC, with support from CNCF staff and the TAB, is responsible for engaging adopters, gathering publication consent, scheduling, conducting, summarizing, gathering final approval, and including the approved summary of the interview within the Due Diligence.

Projects should not coach or instruct adopters with answers to interview questions and may encourage adopters to be open and transparent during the interviews. All interview notes are kept private unless permission is received from the adopter's organization for the notes to be made public.

The TOC maintains a [core list of questions](/operations/toc-templates/template-adopter-questions.md) intended to initiate discussion with adopters, but may add additional questions, or skip questions depending on the course of the interview and the organization's level of comfort in providing responses.

Interviews typically do not take more than 30-60 minutes to complete, and TOC and TAB members should anticipate about 1 hour of time dedicated to summarizing adopter responses. Some interviews may take more or less time, it is dependent upon the interview, any prior interviews that may introduce conflicting stories among adopters, or other concerns noted by the TOC members and the community.

#### Reaching out to Adopters

Once a TOC member has a listing of potential interviewees, they should leverage the [Adopter Interview Request email template](/operations/toc-templates/template-adopter-interview-request.md) to engage and include any TAB members interested in supporting the interview. The email template provides the essential information needed for interviewers to coordinate with their marketing, PR, or other outreach team for approval and allows adopters the opportunity to remain anonymous.

It is imperative that the TOC honor any anonymity concerns and work to address them with adopters, the final approval of the summarized response is a mechanism that allows us to confirm with the adopter their comfort and approval of the final content intended for publication and make any corrections they feel are warranted. TOC members are free to bring in the Chair or Vice Chair to assist in addressing such concerns.

To ease scheduling with Adopters, TOC members are recommended to either include set aside dates/times for adopters as part of the initial email, or to provide a scheduling link to expedite scheduling and avoid delay.

It is anticipated that a minimum of three adopter interviews are required to appropriately ascertain adoption of a project. However in the course of interviewing, you may find that you need additional adopters or types of adopters to be interviewed.

For projects moving from Incubation to Graduation, if considerable time has passed since Incubation (according to the TOC's judgement), the TOC should refresh the Adopter interviews. This may be done by reaching out to previous interviewees, by engaging a new group of adopters for interviews, or some combination thereof. If the time period between Incubation and Graduation is fairly recent, the TOC member(s) should exercise their judgement in choosing to pursue additional interviews. That decision should be recorded with justification in the adoption section of the template.

#### Recording Responses

**Adopter interviews are expected to be interactive**. The [adopter questions template](toc-templates/template-adopter-questions.md) should serve as a starting point for questions when interviewing, however TOC members are expected to use their best judgement in asking questions, deep diving on responses, and crafting additional questions or skipping others. 

You may need to record the meeting to fully capture the responses or take sufficient notes that you can summarize the discussion and convey, with enough breadth, how the adopter is using the project, what environments (such as dev, test, prod), their engagement with the project, use, experience, and future plans.

TOC and TAB members are NOT to email questions to adopters in order for them to write in their responses. If there is a language or time zone challenge, TOC members are expected to inform the broader TOC and seek assistance.

#### Summary Approval

TOC members will summarize responses to the questions asked in a separate, non-public document until the Adopter approves the content. 

Once you have summarized the interview, typically about 5 paragraphs in length depending on the detail allowed to be conveyed, share the doc with the Interviewee so they have permissions to make any changes. Provide the Interviewee with the doc, ask them to make any changes and provide approval two weeks from the date of sending. This provides you with a timeframe to follow-up with them so that it does not get lost or overlooked.

#### Add Interview Summaries to Due Diligence

Once approval is received, the summarized content is copied into the Due Diligence PR in the appropriate section. Linking to the summary is NOT recommended as it may convey identifying information in the history if not properly access controlled which will circumvent anonymity assurances.

### Summarizing the Due Diligence Evaluation

Once you have completed evaluating the project's implementation of the criteria and included the Adopter Interviews, you can summarize the overall evaluation.

Evaluation summary is composed of two parts: the Criteria and the Adoption. The evaluation summary should read like an executive summary of the due diligence, conveying any notable areas and the assessment of sufficient adoption for the level the project applied. An example structure is provided below:

> $TOCMEMBER conducted the due diligence of $PROJECT who applied for $LEVEL. The project [has/has not] completed the criteria that show its maturity at $LEVEL. The following criteria implementations are noteworthy to call out... $NOTABLES. The following actions were provided to the project that were considered blocking but since resolved... $BLOCKERS. The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project... $RECOMMENDATIONS.
> 
> The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY.
> 
> [The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS].

## TOC Internal Review

Once the TOC member(s) has completed the Due Diligence, they should create a PR in their personal TOC repo and share the link with the TOC for review. The TOC member(s) should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "TOC Review".

The TOC member(s) should craft a slack message thread that provides the project name, level applied, links to the PR and issue, and thread any additional call outs, questions, or potential issues warranting further discussion not otherwise captured in the PR. The internal review is expected to last 1 week, assuming no issues are brought up. 

The TOC member(s) is responsible for updating the project with the change in status for internal review.

### Reapplication

In the event a project was not ready to move levels after the due diligence was completed and the project has reapplied through an issue, the previous or new TOC member(s) assigned will initiate a new Due Diligence based on the previous one. The TOC should refresh the prior evaluations with corresponding dates to show changes and improvements and amend the evaluation statements accordingly. 

## Public Comment Period

Assuming no issues have come up during the TOC internal review, the TOC member may put the due diligence out for public comment. The TOC member(s) should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Public Comment".
 
TOC members are to leverage the [public comment template](toc-templates/template-dd-public-comment.md) and be mindful of the timeline to consider if a freeze is in effect or soon will be. All public comment messages are to be sent on the [TOC public mailing list](https://lists.cncf.io/g/cncf-toc/topics).  Once sent, they should be linked on the PR and the project notified.

## Voting

Once the public comment period is complete, the TOC member needs to adjudicate any responses and comments on the PR. Provided no new or blocking information has been shared, the TOC member should request CNCF Staff to initiate voting on the PR, at which point the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Voting". 

TOC members, especially those who filed the PR, need to record their vote by the thumbs up or thumbs down on the PR comment where gitvote was initiated.

## Press Coordination and Completion

Once voting has passed, the CNCF Staff take over the remainder of the process, coordinating press releases and announcements for the project. CNCF Staff will update the project' card on the  [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) accordingly.
