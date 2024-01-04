# TOC Due Diligence guide

This document provides the TOC with guidance on how to execute due diligence of CNCF projects for each level of maturity.  It complements the Moving Levels process detailed in the [Moving Levels Directory](../moving-levels/).

## What is due diligence?

Due diligence is the process by which the TOC performs an independent review of CNCF projects to assess their posture, maturity, and adoption across a variety of technical, governance, and community focuses. The intent of the due diligence is to provide project and adopters with an independent senior technical evaluation of a project's readiness for production. Similar to how organizations have established software development processes and check points prior to software delivery or deployment that ensure the software meets the organization's goals and objectives, the due diligence is a point in time artifact of the project's acheivement for meeting the goals and objectives expected for their maturity level. By performing the due diligence on CNCF projects, the TOC supports our adopters in gaining confidence that the project has been reviewed against documented criteria for their maturity level, can understand any deviations from those criteria, and may not need to repeat this type of evaluation, rather they may incorporate or leverage the contents of the due diligence to guide and information their decision making. It also conveys to adopters the potential level of effort in adopting and integrating the project into their archicture and infrastructure.  For projects, the due diligence provides an evaluation of the project against the expectations of adopters across multiple focuses. It can and often will include additional recommendations to the project that may support them in reaching the next level of maturity, improving their documentation or architecture, and in some cases highlight opportunities to distinguish themselves and their features or functionality from other projects within the same area.

Due Diligence is typically conducted when a project has applied to move levels, after the project completes and asserts their adherence to the criteria for the level they are seeking to move to. For example, an incubating project must complete the graduation criteria before applying to graduate. Part of the TOC's responsiblity is to not only independently verify those assertions, but to understand those assertions in the context of the individual project, its domain, its interoperability, and subsequent alignment with the Foundation, ecosystem, adopters, and overall industry.

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

It is critical that TOC members strive to complete due diligence in a reasonable amount of time. The process involves multi-parties that requires coordination and clear communication of expectations. Delays in completing due diligence can create friction with projects and may encounter TOC member term endings, requiring project to start fresh with a new member. It is expected the process will take time, adopter interviews may be difficult to schedule in a timely fashion, so being upfront on these expectations is important.

Every TOC member is expected to conduct due diligence of CNCF projects. In cases where there may be a perceived conflict of interest, the due diligence must have two TOC members participating in order to dissolve any illusion of bias (for or against).

TOC members may not take on anymore than two (2) projects for due diligence at a given time unless one of the following conditions is true:
* the TOC member is functioning as a guide to new TOC members learning this process
* the TOC member is has two projects in voting
* the TOC member has one project in voting, and another in progress

We expect all TOC members to be mindful of their obligations and timelines, whether they be work, cloud native, or personal and manage their workload accordingly.

### Project has applied to move levels

**Projects apply to move levels by submitting an Issue on the TOC repo that details their conformance to the criteria.** These Issues go into the [TOC backlog](https://github.com/orgs/cncf/projects/27/views/9) and are worked on a first-in, first-out/started basis to the extent that is reasonably practical - each TOC member has subject matter expertise in one domain or another and each TOC member serves as a liasion to a Technical Advisory Group that may or may not align with that domain. The TOC strives to align their skills and background with the project applying they are assigned to ensure they have the context to understand the project within that domain. 

The issue will contain a limited set of information about the project at the time of its application, commonly asserting its conformance to the stated criteria with links to where or descriptions as to how they are implemented.

### TOC member(s) step forward to be assigned

TOC members ready to perform due diligence a project's application will socialize this internally with the TOC to provide opportunity for other members to participate. Once a TOC member or members is determined, those TOC members must assign themselves to the Issue and move the issue's card on the [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Assigned".

TOC members may occassionally be contacted by projects asking the TOC member to sponsor their application to move levels. When this occurs, politely reminder projects of the TOC's first-in, first-started process and let them know it is subject to TOC availability. Follow up internally to the TOC that you were contacted and by which project so we may all be aware of which projects are reaching out as it may identify a backlog, timing, or communication issue.

### Initial evaluation

Once the TOC member is assigned the project, they should perform a cursory, light-weight evaluation of the project's completion of the criteria. If some of the criteria are not yet met, or missing, the TOC should notify the project of the issues requiring resolution before re-applying, and once confirmed by the project, comment publicly on the Issue with those recommendations for resolution and close it. TOC members should use their best judgement in determining if the unmet criteria are simple fixes or if they require substantial effort or time to properly complete. For example, a project applying to graduation should have clear and discoverable governance documentation. If the TOC member cannot find any governance documentation at all, they should engage the project to confirm that none exists. If it does exist, but is not readily discoverable, the TOC member may continue to move forward with due diligence as improving discoverable may be resolved through appropriate linking. If it doesn't exist, the TOC should finish the lightweight review, capture all unmet criteria, and engage the project prior to commenting and closing the Issue.

### Kicking off the due diligence

Once the project is assigned a TOC member and the initial evaluation looks good to proceed, the TOC member engages the project's maintainers or leadership group to kick off the due diligence. This can be done by commenting on the Issue, starting a slack channel (public or private), slack group direct message, email exchange, or thread in the project's primary communication channel. We strive to keep these discussions transparent and visible but should consider any potential sensitive issues that come about as a result of the review (resolution of vulnerabilities for instance).

TOC members, with support from CNCF staff, should schedule a meeting with the project to the extent possible given availability and timezones. Asynchrounous kick-offs can occur, but may result in additional back and forth or delays. Each Kick-off meeting should have a central kick-off document that allows the TOC and the project to capture expectations, decisions, timelines, and other pertinent references needed for successful completion of the due diligence. A [kick-off meeting template](toc-templates/template-kickoff-notes.md) is located in the [toc-templates](toc-templates/) folder.

Once the Kick-off is scheduled, the TOC member should move the project's card on the [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "In Due Diligence".

### Completing Due Diligence

The TOC will use the appropriate Due Diligence for the project's applied level as the basis for a PR ([Incubation template](toc-templates/template-dd-pr-incubation.md), [Graduation template](toc-templates/template-dd-pr-graduation.md)) and will evaluate the project's assertions in the application issue against the discoverable and publically available sites, repos, files, and other artifacts of the project. The TOC's evaluation against each criteria goes in the corresponding area of the PR template.

Previously, the project was responsible for completing the due diligence such that it allowed the TOC member to review and comment. Due to the extensive back and forth in this prior process and with recent changes to the criteria, the TOC has altered the process leverage a Due Diligence PR as the TOC's assessment of the projects completion of the criteria. Therefore TOC members are expected to complete the Due Diligence PR with support from the project and TAG(s).

As the TOC member reviews the criteria, any deviations or implementation notes from the review should be recorded within the due diligence PR as part of their evaluation for that specific criteria. The TOC member will provide an overall evaluation statement that summarizes the content of the due diligence once the Adopter Interview Summaries are recorded. For more information on adopter interviews, refer to [Conducting Adopter Interviews](#conducting-adopter-interviews).

As an example, let's say a project is asserting their sub-projects have leadership, contribution information, maturity statuses, and add/removal processes. The TOC member, in confirming this may determine that the sub-projects share the same contributor file on the org's evolution directory and may change the maturity level in accordance with their documented process, but the process is not clear as to what initiates that change or who. For a sandbox project seeking incubation, this may be non-blocking but the TOC may recommend the project improve the documented process. As such, the TOC member will record their finding and corresponding recommendation.

Another example, if the TOC member is looking over the project's stated goals and objectives that have not changed since the project was accepted into the CNCF, and they are now applying to Graduate, the TOC member will ask the project to clarify or provide additional information as to why the project hasn't iniatiated any changes and if they still feel those goals and objectives are accurate for the future of the project. The TOC member will then summarize the response and record it in the PR under the corresponding criteria evaluation.

It is expected that the TOC's evaluation of a project's completion of the criteria may reveal a mismatch in understanding or an unexpected implementation. Documenting the TOC evaluation in the Due Diligence PR provides the project, TAGs, community, adopters, and TOC with a point of reference to understand if the criteria are meeting the outcomes required of a project for a certain maturity level, or if compensating mechanisms that supplement or augment the criteria are in place that work best for that specific project. 

When the TOC has finished their criteria evaluation, they should move the project's card on the [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Active Review & Discussion" and re-engage the project to elevate and discuss any items neededing clarity, correction, or improvement. This includes notifying the project of any recommendations. Recommendations and discussion points may copied into the kick-off document to faciliate discussion and to provide for additional context and discussion with the project until they are finalized.

The TOC member may then file the PR and place it in draft until the Adopter Interviews are completed.

#### Engaging TAGs

If the project has previously engage with a TAG, the TOC may request an update or summary regarding the engagement and the TAG's assessment of the project in the context of the domain of the project. For instance, a policy-as-code project may present to TAG Security, TAG Security may summarize their recommendations to the project and link to supplemental material (like the project's self-assessment or joint-assessment) which is then added in the Due Diligence PR. 

Feedback by the TAG is encouraged prior to Due diligence being initiated, as requesting feedback when due diligence is occurring may prolong the due diligence process. It is expected that the project will have completed all criteria prior to applying, if this is not the case, due diligence should not commence.

### Conducting Adopter Interviews

After the evaluation has incorporated project feedback and discussion, the TOC member may move the project's card on the [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Adopter Interviews" to begin outreach and scheduling with adopters.

In order to appropriately ascertain the adoption of a project, the TOC interviews a sampling of the project's adopters to understand how it is being used, what problems it is solving, the ease of adoption and integration, the community and contribution experience, and learn how adopters are experiencing the project's maturity level.

The TOC member should request 5-7 potential adopters to be interviewed and work with the TOC on gathering contact information for them. The TOC, with support from CNCF staff, is responsible for engaging adopters, gathering publication consent, scheduling, conducting, summarizing, gathering final approval, and including the approved summary of the interview within the Due Diligence.

The TOC maintains a core list of questions intended to initiate discussion with adopters, but may add additional questions, or skip questions depending on the course of the interview and the organization's level of comfort in providing responses.

Interviews typically do not take more than 30 minutes to complete, and TOC members should anticipate about 1 hour of time dedicated to summarizing adopter responses.

#### Reaching out to Adopters

Once a TOC member has a listing of potential interviewees, they should leverage the [Adopter Interview Request email template](toc-templates/adopter-interview-request-template.md) to engage. The email template provides the essential information needed for interviewers to coordinate with their marketing, PR, or other outreach team for approval and allows adopters the opportunity to remain anonymous. It is imperative that the TOC honor any anonymity concerns and work to address them with adopters, the final approval of the summarized response is a mechanism that allows us to confirm with the adopter their comfort and approval of the final content intended for publication and make any corrections they feel are warranted.

To ease scheduling with Adopters, TOC members are recommended to either include set aside dates/times for adopters as part of the initial email, or to provide a scheduling link to expedite scheduling and avoid delay.

It is anticipated that a minimum of three adopter interviews are required to appropriately ascertain adoption of a project. However in the course of interviewing, you may find that you need additional adopters to be interviewed.

For projects moving from Incubation to Graduation, if considerable time has passed since Incubation (according to the TOC's judgement), the TOC should refresh the Adopter interviews. This may be done by reaching out to previous interviewees, by engaging a new group of adopters for interviews, or some combination thereof. If the time period between Incubation and Graduation is fairly recent, the TOC member should exercise their judgement in choosing to pursue additional interviews. That decision should be recorded with justification in the adoption section of the template.

#### Recording Responses

Adopter interviews are expected to be interactive. The [adopter questions template](toc-templates/template-adopter-questions.md) should serve as a starting point for questions when interviewing, however TOC members are expected to use their best judgement in asking questions, deep diving on responses, and crafting additional questions or skipping others. 

You may need to record the meeting to fully capture the responses or take sufficient notes that you can summarize the discussion and convey, with enough breadth, how the adopter is using the project, what environments (such as dev, test, prod), their engagement with the project, use, experience, and future plans.

#### Summary Approval

TOC members will summarize responses to the questions asked in a separate, non-public document until the Adopter approves the content. 

Once you have summarized the interview, typically about 5 paragraphs in length depending on the detail allowed to be conveyed, share the doc with the Interviewee so they have permissions to make any changes. Provide the Interviewee with the doc, ask them to make any changes and provide approval two weeks from the date of sending. This provides you with a timeframe to follow-up with them so that it does not get lost or overlooked.

#### Add Interview Summaries to Due Diligence

Once approval is received, the summarized content is copied into the Due Diligence PR in the appropriate section. Linking to the summary is NOT recommended as it may convey identifying information in the history if not properly access controlled which will circumvent anonymity assurances.

### Summarizing the Due Diligence Evalution

Once you have completed evaluating the project's implementation of the criteria and included the Adopter Interviews, you can summarize the overall evalution.

Evaluation summary is composed of two parts: the Criteria and the Adoption. The evaluation summary should read like an executive summary of the due diligence, conveying any notable areas and the assessment of sufficient adoption for the level the project applied. An example structure is provided below:

> $TOCMEMBER conducted the due diligence of $PROJECT who applied for $LEVEL. The project [has/has not] completed the criteria that show its maturity at $LEVEL. The following criteria implementations are noteworthy to call out... $NOTABLES. The following actions were provided to the project that were considered blocking but since resolved... $BLOCKERS. The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project... $RECOMMENDATIONS.
> 
> The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY.
> 
> [The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS].

## TOC Internal Review

Once the TOC member has completed the Due Diligence, the TOC member tags the TOC on the PR for an TOC internal review. The TOC member should move the project's card on the [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "TOC Review".

The TOC member should craft a slack message thread that provides the project name, level applied, links to the PR and issue, and thread any additional call outs, questions, or potential issues warranting further discussion not otherwise captured in the PR. The internal review is expected to last 1 week, assuming no issues are brought up.

The TOC member is responsible for updating the project with the change in status for internal review.

### Reapplication

In the event a project was not ready to move levels after the due diligence was completed and the project has reapplied through an issue, the previous or new TOC member assigned will initiate a new Due Diligence based on the previous one. The TOC should refresh the prior evaluations with corresponding dates to show changes and improvements and ammend the evaluation statements accordingly. 

## Public Comment Period

Assuming no issues have come up during the TOC internal review, the TOC member may put the due diligence out for public comment. The TOC member should move the project's card on the [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Public Comment".
 
TOC members are to leverage the [public comment template](toc-templates/template-dd-public-comment.md) and be mindful of the timeline to consider if a freeze is in effect or soon will be. All public comment messages are to be sent on the [TOC public mailing list](https://lists.cncf.io/g/cncf-toc/topics).  Once sent, they should be linked on the PR and the project notified.

## Voting

Once the public comment period is complete, the TOC member needs to adjudicate any responses and comments on the PR. Provided no new or blocking information has been shared, the TOC member should request CNCF Staff to initiate voting on the PR, at which point the project's card on the [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Voting". 

TOC members, especially those who filed the PR, need to record their vote by the thumbs up or thumbs down on the PR comment where gitvote was initiated.

## Press Coordination and Completion

Once voting has passed, the CNCF Staff take over the remainder of the process, coordinating press releases and announcements for the prokject. CNCF Staff will update the project' card on the  [Appliction to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) accordingly.
