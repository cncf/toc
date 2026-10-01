# TOC Due Diligence guide

This document provides the TOC with guidance on how to execute Due Diligence of CNCF projects for each level of maturity.  It complements the Moving Levels process detailed in the [Process Directory](../process/README.md).

## Quick Links

Getting Started: **[Triage applications](#initial-triageevaluation-prior-to-assignment)** | **[TOC Assignment](#toc-members-step-forward-to-be-assigned)** | **[Kick-off](#kicking-off-the-due-diligence)**

The Due Diligence (DD):  **[Due Diligence](#completing-due-diligence)** | **[Finalizing DD](#finalizing-the-due-diligence)** | **[Adopter Interviews](#conducting-adopter-interviews)**

Wrapping it up: **[TOC internal review](#toc-internal-review)** | **[Public Comment](#public-comment-period)** | **[Voting](#voting)**

## What is Due Diligence?

Due diligence is the process by which the TOC performs an independent review of CNCF projects to assess their posture, maturity, and adoption across a variety of technical, governance, and community focuses. The intent of the Due Diligence is to provide project and adopters with an independent senior technical evaluation of a project's readiness for production. Similar to how organizations have established software development processes and check points prior to software delivery or deployment that ensure the software meets the organization's goals and objectives, the Due Diligence is a point in time artifact of the project's achievement for meeting the goals and objectives expected for their maturity level. By performing the Due Diligence on CNCF projects, the TOC supports our adopters in gaining confidence that the project has been reviewed against documented criteria for their maturity level, can understand any deviations from those criteria, and may not need to repeat this type of evaluation, rather they may incorporate or leverage the contents of the Due Diligence to guide and information their decision making. It also conveys to adopters the potential level of effort in adopting and integrating the project into their architecture and infrastructure.  For projects, the Due Diligence provides an evaluation of the project against the expectations of adopters across multiple focuses. It can and often will include additional recommendations to the project that may support them in reaching the next level of maturity, improving their documentation or architecture, and in some cases highlight opportunities to distinguish themselves and their features or functionality from other projects within the same area.

Due Diligence is typically conducted when a project has applied to move levels, after the project completes and asserts their adherence to the criteria for the level they are seeking to move to. For example, an incubating project must complete the graduation criteria before applying to graduate. Part of the TOC's responsibility is to not only independently verify those assertions, but to understand those assertions in the context of the individual project, its domain, its interoperability, and subsequent alignment with the Foundation, ecosystem, adopters, and overall industry.

The criteria for moving levels sets the expectations for what projects are to complete prior to applying for the next level. Criteria are outcome oriented and have been developed to provide projects with maximum flexibility in implementation while demonstrating the characteristics and hallmarks of maturing projects. TOC members may exercise their discretion in evaluating conformance to the criteria, however deviations, waivers, and exceptions should be clearly articulated and reasoned such that the desired outcome is still met through compensating mechanisms or demonstrated to not be relevant.

The process of conducting Due Diligence is not to create unnecessary work for projects, rather to ensure projects are set up to succeed in a way that works for them against the values, principles, and definitions of cloud native and the expectations adopters carry for scalable, secure, production-ready software.

It is important to remember that the CNCF is for the benefit of its members. This means the CNCF, through the TOC, defines the technical and community expectations for projects to receive benefits of the foundation. The criteria, Due Diligence process, security audits, templates, code of conduct, and community management practices are several expectations for projects to receive benefits. Adherence to these, and confirmation of adherence for each project's maturity level is essential to continuously receive those benefits. 

As a matter of balance, it is important to note that while these benefits provide value to projects, the TOC is responsible for continuously assessing the criteria, Due Diligence process, and overall levels framework to ensure we are adapting to shifts and changes in the landscape and market, as well as responding to the needs and limitations experienced by projects. _The community and projects are our partners in effecting that change, so feedback is crucial._

## How to conduct Due Diligence

We will not focus on the events that have transpired prior to a project's application to move levels, choosing instead to focus on the point when the TOC is engaged.

Currently, there exist three levels in the CNCF:
* _Sandbox_ - experimental and early stage projects worth exploring their viability and adoption probability within the ecosystem
  * _Sandbox_ projects undergo Due Diligence when they apply to _Incubation_.
  * _Sandbox_ projects do not receive Due Diligence in order to join the Foundation.
* _Incubating_ - projects that have proven concepts, are receiving market and adopter attention, and increasing their robustness, security, scalability, and maturity
  * _Incubation_ projects receive Due Diligence when they apply for _Graduation_.
* _Graduated_ - projects that are stable, resilient, secure, and highly mature, they are capable of meeting the expectations of production-ready software
  * _Graduated_ projects do not receive additional Due Diligence once they reach this level.
  * By the time projects have _graduated_, they will have undergone two rounds of Due Diligence, once during the application to _incubation_, and a second time when they apply to _graduate_ - often a refresh depending on age since the last Due Diligence.

Projects not already in the Foundation may apply directly to Incubation if they feel they are ready. At this point of application, they undergo Due Diligence that also considers their fit in the CNCF, much the same as the considerations for inclusion of sandbox projects.

It is critical that TOC members strive to complete Due Diligence in a reasonable amount of time. The process involves multi-parties that requires coordination and clear communication of expectations. Delays in completing Due Diligence can create friction with projects and may encounter TOC member term endings, requiring project to start fresh with a new member. It is expected the process will take time, adopter interviews may be difficult to schedule in a timely fashion, so being upfront on these expectations is important. **Each TOC member is expected to manage their project's application timeline to the best of their ability and reduce any delays where possible.**

Every TOC member is expected to conduct Due Diligence of CNCF projects and triage those applications. In cases where there may be a perceived [conflict of interest](#conflicts-of-interest), the Due Diligence must have two TOC members participating in order to dissolve any illusion of bias (for or against).

TOC members may not take on anymore than two (2) projects for Due Diligence at a given time unless one of the following conditions is true:
* the TOC member is functioning as a guide to new TOC members learning this process
* the TOC member is serving as a secondary to the primary TOC member conducting a Due Diligence in order to offer additional domain support or other relevant subject matter expertise
* the TOC member has two projects in voting
* the TOC member has one project in voting, and another in progress

We expect all TOC members to be mindful of their obligations and timelines, whether they be work, cloud native, or personal and manage their workload accordingly.

### Project has applied to move levels

**Projects apply to move levels by submitting an Issue on the TOC repo that details their conformance to the criteria.** These Issues go into the [TOC backlog](https://github.com/orgs/cncf/projects/27/views/9) and are worked on a first-in, first-out/started basis to the extent that is reasonably practical - each TOC member has subject matter expertise in one domain or another and each TOC member serves as a liaison to a Technical Advisory Group that may or may not align with that domain. The TOC strives to align their skills and background with the project applying they are assigned to ensure they have the context to understand the project within that domain. 

The issue will contain a limited set of information about the project at the time of its application, commonly asserting its conformance to the stated criteria with links to where or descriptions as to how they are implemented.

### Initial triage/evaluation prior to assignment

All TOC members are expected to assist in the triaging of project applications to move levels to ensure that when a TOC member is ready to be assigned, the project is ready to be evaluated with no immediate blockers that would inhibit or delay the TOC's engagement.

This light-weight triage/evaluation must cover the list below (it is not exhaustive and is a minimum triage set from the [incubation template retrieved 12 DEC 2025](https://github.com/cncf/toc/blob/c2943ffc98064dd88e9ef9c4afd5a8856898942f/.github/ISSUE_TEMPLATE/template-incubation-application.md)):
* Adoption Assertion includes the Adopters file link, and the project has an entry in the [Adopter Submission Form](https://docs.google.com/forms/d/1n1oLC6IKj5-7S_xeEjIdEjbtS9SWniuAo7IIOyLFuK8/edit) responses to provide 5-7 adopters to reach out to. [Check here](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for the definition of an Adopter.
* Application Process Principles provides 
  * link to a completed [General Technical Review (GTR)](project-reviews/general-technical-questions.md). The GTR snapshot should be submitted as a PR into `projects/<project-name>/tech-review/` in this repo, named `YYYY-MM-DD.md` to reflect the date of the snapshot. Projects may maintain a working copy in their own repository, but the TOC repo entry is the required artifact for the DD.
  * link to a completed [Governance Review](project-reviews/governance-review-template.md)
  * (optional) link to a technical presentation to a CNCF TAG, or links to notes or issues about such a presentation
  * assertion of vendor neutrality
  * assertion of review and acknowledgement of expectations of CNCF projects and requirements for moving forward through the CNCF maturity levels
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
  * link to completed self-assessment in `cncf/toc` under `projects/<project>/security-assessment/self-assessment.md`
  * link to best practices badge, confirm it is "passing" and 100% complete 
* Ecosystem
  * link to adopters file
  * link to integrations/ compatibility information of other projects and products

Projects should complete Technical and Governance reviews themselves and store them in a .project folder or GitHub repository in their own organization. CNCF associates will review and upload vetted snapshots of the project's copies to [github.com/cncf/toc/projects](../projects), but projects should not be blocked by this process, as it depends on availability of community members in CNCF TAGs which cannot be guaranteed.

If some of the criteria are not yet met, or missing, the TOC member triaging will add a comment detailing all items that are unmet or missing and close the application; affixing the "not-ready" label and move the card to the "Not-Ready-Will Return" column of the [TOC project board](https://github.com/orgs/cncf/projects/27/views/9)'s Applications to Move levels tab. Projects are expected to re-apply upon completion of outstanding items. When the project is ready to reapply, they should link to the previous application so the TOC may leverage and reuse as much prior work as reasonable.

Once the TOC has triaged the application and found all criteria to have content, the TOC member performing triage comment the application is complete and ready for assignment. The TOC member will affix the "ready" label and move the project from the "new" column on the application's board. The project's application will be updated by the TOC member with a comment that details where work still needs to be completed, next steps associated with completion of those, and an estimated timeframe that the project is likely to complete those items by. 

TOC members are expected to triage projects in the "new" column on the board for projects that are returning from a previous not-ready application, verify completion, and move them to the top of the ready for assignment column. 

TOC members are to prioritize selecting projects from the ready for assignment column over the new column to expedite the queue and make the best use of TOC time.

### TOC member(s) step forward to be assigned

Commonly referred to as the Project's Application Sponsor, TOC members assign themselves to projects to sponsor the application for moving levels. Sponsoring an application ensures a focused point of contact exists for both the project and the TOC in completing the Due Diligence, public comment, and execution of voting. 

The TOC member that assigns themselves a project to sponsor the application for moving levels may request a secondary TOC member to support the Due Diligence according to eligibility. They may also request that a member of the project-reviews subproject verify and commit the project's Technical and Governance reviews to the TOC repo.

TOC members ready to perform Due Diligence a project's application will socialize this internally with the TOC to provide opportunity for other members to participate. Once a TOC member or members is determined, those TOC members must assign themselves to the Issue and move the issue's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Assigned".

TOC members may occasionally be contacted by projects asking the TOC member to sponsor their application to move levels. When this occurs, politely remind projects of the TOC's first-in, first-started process and let them know it is subject to TOC availability. Follow up internally to the TOC that you were contacted and by which project so we may all be aware of which projects are reaching out as it may identify a backlog, timing, or communication issue.

#### Conflicts of interest

The Due Diligence process provides several mechanisms that serve as check points for bias or uncomprehensive evaluation. The public comment period is the most open and transparent point at which the community and others may independently review the diligence performed by the TOC.  Additionally, TOC members conduct an internal review of the Due Diligence to verify all aspects of the project have been thoroughly reviewed.

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

### Kicking off the Due Diligence

Once the project is assigned to the TOC member(s), the TOC member(s) engages the project's maintainers or leadership group to kick off the Due Diligence. This can be done by commenting on the Issue, starting a slack channel (public or private), slack group direct message, email exchange, or thread in the project's primary communication channel. We strive to keep these discussions transparent and visible but should consider any potential sensitive issues that come about as a result of the review (resolution of vulnerabilities for instance). 

Any form of communication must include _two members_ of CNCF staff to ensure consistency and continuity throughout the process. 

TOC members, with support from CNCF staff, should schedule a meeting with the project to the extent possible given availability and timezones. Asynchronous kick-offs can occur, but may result in additional back and forth or delays. Each Kick-off meeting should have a central kick-off document that allows the TOC and the project to capture expectations, decisions, timelines, and other pertinent references needed for successful completion of the Due Diligence. A [kick-off meeting template](toc-templates/template-kickoff-notes.md) is located in the [toc-templates](toc-templates/) folder.

Once the Kick-off is scheduled, the TOC member(s) should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "In Due Diligence". It is *strongly* recommended that you inform the project to verify compliance with the CNCF's licensing policy (set forth in the [Section 11 of the Charter](https://github.com/cncf/foundation/blob/master/charter.md) with [additional information here](https://github.com/cncf/foundation/blob/main/policies-guidance/allowed-third-party-license-policy.md)).

### Completing Due Diligence

NOTE: the Due Diligence can be completed in conjunction with adopter interviews, some TOC members find completion of the DD as informative to conducting interviews, but not in all cases.

The TOC will use the appropriate Due Diligence for the project's applied level as the basis for a PR ([Incubation template](toc-templates/template-dd-pr-incubation.md), [Graduation template](toc-templates/template-dd-pr-graduation.md)) and will evaluate the project's assertions in the application issue against the discoverable and publicly available sites, repos, files, and other artifacts of the project. The TOC's evaluation against each criteria goes in the corresponding area of the PR template.

Previously, the project was responsible for completing the Due Diligence such that it allowed the TOC member(s) to review and comment. Due to the extensive back and forth in this prior process and with recent changes to the criteria, the TOC has altered the process to leverage a Due Diligence PR as the TOC's assessment of the projects completion of the criteria. Therefore TOC members are expected to complete the Due Diligence PR with support from the project and TAG(s).

As the TOC member(s) reviews the criteria, any deviations or implementation notes from the review should be recorded within the Due Diligence PR as part of their evaluation for that specific criteria. The TOC member will provide an overall evaluation statement that summarizes the content of the Due Diligence once the Adopter Interview Summaries are recorded. For more information on adopter interviews, refer to [Conducting Adopter Interviews](#conducting-adopter-interviews).

As an example, let's say a project is asserting their sub-projects have leadership, contribution information, maturity statuses, and add/removal processes. The TOC member(s), in confirming this may determine that the sub-projects share the same contributor file on the org's evolution directory and may change the maturity level in accordance with their documented process, but the process is not clear as to what initiates that change or who. For a sandbox project seeking incubation, this may be non-blocking but the TOC may recommend the project improve the documented process. As such, the TOC member(s) will record their finding and corresponding recommendation.

Another example, if the TOC member(s) is looking over the project's stated goals and objectives that have not changed since the project was accepted into the CNCF, and they are now applying to Graduate, the TOC member(s) will ask the project to clarify or provide additional information as to why the project hasn't initiated any changes and if they still feel those goals and objectives are accurate for the future of the project. The TOC member(s) will then summarize the response and record it in the PR under the corresponding criteria evaluation.

It is expected that the TOC's evaluation of a project's completion of the criteria may reveal a mismatch in understanding or an unexpected implementation. Documenting the TOC evaluation in the Due Diligence PR provides the project, TAGs, community, adopters, and TOC with a point of reference to understand if the criteria are meeting the outcomes required of a project for a certain maturity level, or if compensating mechanisms that supplement or augment the criteria are in place that work best for that specific project.

#### Sub-projects

When reviewing a project, you may find multiple sub-projects within their organization or repositories. It is important that the Due Diligence document reflects the scope of evaluation for a given project and what that includes. Sub-projects are often broken out into two categories: packaged with the project, and independent but related to the project (SDK reference implementations for instance). 

Sub-projects packaged with the main project are within scope of the Due Diligence evaluation. Any sub-project not packaged is outside the scope of the Due Diligence evaluation and must be indicated in the Due Diligence as "un-evalauted" and the reason. TOC members must verify that the main project has sufficient governance of its sub-projects that allows adopters and contributors to understand the stability, maturity, and requirements or criteria to be a sub-project, their relation to the main project, and other information necessary commensurate with the maturity level of the project.

#### Licensing

In the course of conducting a project's Due Diligence, you may become aware of license issues, outstanding exception requests, or other matters concerning the licensing of the project and its dependencies. If in the course of Due Diligence you find the project appears to be missing a license exception for their code or dependencies (dynamic or statically linked), notify the project to [file a license exception request](https://github.com/cncf/foundation/issues/new/choose). The project's application to move levels may not go to a vote until the exception has been approved, however activities such as adopter interviews and public comment may proceed. It is expected that sandbox projects, when accepted, undergo a licensing review to ensure compliance with [CNCF's licensing policy](https://github.com/cncf/foundation/blob/main/license-exceptions/README.md) and exceptions. Refer to the [allowed third party licensing policy](https://github.com/cncf/foundation/blob/main/policies-guidance/allowed-third-party-license-policy.md) for more information.

#### Graduated project security audits

TOC members who sponsor projects seeking graduation are expected to review the results of the audit to confirm the project has resolved all critical and high findings at a minimum. Additional findings are expected to be tracked for resolution by the project. In reviewing the audit of the project, you may find additional recommendations or deltas in the project's operational security (incident response, PR reviewer guidelines, lack of regression tests, etc.) that should be addressed. These are typically non-blocking, however you should take care to note them within the Due Diligence PR after conveying their need and importance of completion to the project. You may reach out to TAG Security for assistance in identifying these areas or reach out to a TOC member with a background in Security.

#### Specification projects

A specification project is one that defines normative behavior (using keywords such as MUST, SHALL, SHOULD per [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119)) that third parties can independently implement. This includes projects that produce protocol definitions, API contracts, data format standards, or interoperability frameworks intended for multi-vendor implementation.

CNCF projects generally fall into three categories for DD purposes:

- **Specification projects** produce normative documents that define behavior for independent implementers. The spec text is the primary deliverable.
- **Reference implementations** accompany a specification and demonstrate how to implement it. A reference implementation may be part of the same CNCF project as the spec or may exist as a separate project.
- **Standalone tools and libraries** are software projects evaluated under the standard DD criteria without spec-specific considerations.

A project may be **hybrid** - producing both a specification and a reference implementation. For hybrid projects, the reference implementation must independently meet the standard (non-spec) DD criteria for governance, adoption, and engineering maturity. The spec and implementation components should each be evaluated on their own merits.

##### When spec-specific evaluation applies

The TOC member conducting DD should apply the spec-specific criteria in this section when the project self-identifies as a specification project in its application, or when the TOC member determines the project includes normative spec content. If there is ambiguity about whether a project qualifies, the TOC member should discuss with the project during kick-off.

##### Spec-specific evaluation criteria

In addition to the standard DD criteria, spec projects are evaluated on the following:

**Implementations.** Spec projects are expected to have multiple independent implementations that demonstrate the spec's viability and adoptability:

- For incubation: at least 2 independent implementations, with at least 1 from outside the primary contributing organization to the spec authoring.
- For graduation: at least 3 independent implementations, from at least 2 organizations.
- A reference implementation created by the spec authors counts toward the implementation threshold but does not satisfy it alone. At least one implementation must come from an organization other than the primary contributing organization to the spec authoring.
- Qualifying implementations must have documented production usage. An implementation that exists only as a proof of concept or test harness does not satisfy this criterion.
- Individual implementations may have limited diversity of maintainers or adoption on their own, but as a whole the implementation ecosystem should be broad and diverse.
- The reference implementation does not need to be part of the project undergoing Due Diligence.
- Some specifications define multiple distinct components that may each have different implementation ecosystems (for example, a security spec may have many client-side implementations but fewer server-side implementations, or an attestation spec may have broad generation support but fewer policy verification implementations). When a spec has asymmetric component coverage, the TOC should evaluate implementation thresholds at the component level where appropriate and use judgment about whether the overall ecosystem demonstrates sufficient independent adoption. The thresholds above represent the expected baseline for the spec's core functionality; components with fewer implementations should have a clear rationale (such as inherently smaller deployment scope) documented in the Due Diligence.

**Spec version stability.** The spec should have at least one stable or versioned release. For graduation, the spec should demonstrate stability through multiple versioned releases and a track record of managed, backward-compatible evolution.

**Conformance and interoperability.** The project should provide documented conformance testing, an interoperability test suite, or other evidence that implementations can be validated against the spec. For graduation, a conformance test suite should be maintained and exercised across implementations.

**Spec lifecycle maturity.** The project should have a formal proposal process for normative changes to the spec (for example: OTEPs, TAPs, gRPC Proposals, or equivalent). This demonstrates that the spec evolves through structured community input rather than unilateral changes.

**Spec editor or equivalent role.** The project should have a named individual or role responsible for maintaining the spec text, with appropriate access and authority. This role may go by different titles across projects.

##### Maintainer considerations for spec projects

It is reasonable for a specification project to have fewer maintainers than a large software project, as the primary deliverable is a normative document rather than a codebase. The number of maintainers should be appropriate to the size and scope of the project. However, spec projects should still demonstrate multi-organization participation in the spec development process. Spec projects should be multi-vendor driven - vendor neutrality thresholds apply with the same or greater weight as for software projects, because a spec controlled by a single vendor risks becoming a proprietary standard rather than an open one.

##### Adopter definition for spec projects

For specification projects, "adopters" are the implementations of the spec and the organizations behind them, not users of a single reference implementation. When conducting adopter interviews for spec projects, the TOC should interview implementers from multiple independent implementations rather than only users of one implementation. This provides stronger evidence of the spec's viability and adoption than user interviews alone.

##### Standards path (optional)

Spec projects may optionally pursue formal standardization through the Joint Development Foundation (JDF). This is not a CNCF maturity criterion and is not evaluated during Due Diligence. Projects interested in this path should coordinate with CNCF staff separately.

#### Governance concentration and practice period

During DD, TOC members should review the project's organizational contribution concentration using [LFX Insights](https://insights.lfx.linuxfoundation.org/) to corroborate the maintainer list. For projects with **>75% organizational dependency** and no org-balance mechanism (such as [org-balanced voting](https://github.com/cncf/project-template/blob/main/GOVERNANCE-org-balanced-voting.md), steering committee with org caps, or equivalent structural protection), the TOC recommends the project implement governance changes and demonstrate them in practice for **3 months** before the DD continues.

Exit criteria for the practice period:

- Org-balanced voting or steering committee adopted and documented
- At least 3 governance meetings held with published minutes
- At least one governance decision made through the new process
- No regression in LFX org dependency

This is not punitive -- it verifies that governance works under real conditions before the project moves forward. Projects with lower concentration or existing structural protections may proceed without a practice period. The TOC member should record the decision (whether a practice period is needed and why) in the DD PR.

For governance best practices and anti-patterns to look for during DD, see the [governance guidance blog post](https://www.cncf.io/blog/), the [governance templates](https://contribute.cncf.io/projects/best-practices/governance/templates/), and the [governance remediation process](governance-remediation-process.md).

### Finalizing the Due Diligence

When the TOC has finished their criteria evaluation, they should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Adopter Interviews & Project Discussion" and re-engage the project to elevate and discuss any items needing clarity, correction, or improvement. This includes notifying the project of any recommendations. Recommendations and discussion points may be copied into the kick-off document to facilitate discussion and to provide for additional context and discussion with the project until they are finalized.

The TOC member(s) may then file the PR and place it in draft until the Adopter Interviews are completed.

#### Engaging TAGs

If the project has previously engage with a TAG, the TOC may request an update or summary regarding the engagement and the TAG's assessment of the project in the context of the domain of the project. For instance, a policy-as-code project may present to TAG Security, TAG Security may summarize their recommendations to the project and link to supplemental material (like the project's self-assessment or joint-assessment) which is then added in the Due Diligence PR. 

Feedback by the TAG is encouraged prior to Due Diligence being initiated, as requesting feedback when Due Diligence is occurring may prolong the Due Diligence process. It is expected that the project will have completed all criteria prior to applying, if this is not the case, Due Diligence should not commence.

### Conducting Adopter Interviews

In order to appropriately ascertain the adoption of a project, the TOC interviews a sampling of the project's adopters to understand how it is being used, what problems it is solving, the ease of adoption and integration, the community and contribution experience, and learn how adopters are experiencing the project's maturity level.

As part of the application process, projects are required to submit 5-7 potential adopters via the [Adopter Submission Form](https://docs.google.com/forms/d/1n1oLC6IKj5-7S_xeEjIdEjbtS9SWniuAo7IIOyLFuK8/edit). The TOC member(s) begin reaching out to these adopters to be interviewed and work with the TAB in conducting the interview. The TOC, with support from CNCF staff and the TAB, is responsible for engaging adopters, gathering publication consent, scheduling, conducting, summarizing, gathering final approval, and including the approved summary of the interview within the Due Diligence.

Projects should not coach or instruct adopters with answers to interview questions and may encourage adopters to be open and transparent during the interviews. All interview notes are kept private unless permission is received from the adopter's organization for the notes to be made public.

The TOC maintains a [core list of questions](/operations/toc-templates/template-adopter-questions.md) intended to initiate discussion with adopters, but may add additional questions, or skip questions depending on the course of the interview and the organization's level of comfort in providing responses.

Interviews typically do not take more than 30-60 minutes to complete, and TOC and TAB members should anticipate about 1 hour of time dedicated to summarizing adopter responses. Some interviews may take more or less time, dependent upon the interview, any prior interviews that may introduce conflicting stories among adopters, or other concerns noted by the TOC members and the community.

#### Reaching out to Adopters

Once a TOC member has a listing of potential interviewees, they should leverage the [Adopter Interview Request email template](/operations/toc-templates/template-adopter-interview-request.md) to engage and include any TAB members interested in supporting the interview. The email template provides the essential information needed for interviewers to coordinate with their marketing, PR, or other outreach team for approval and allows adopters the opportunity to remain anonymous.

It is imperative that the TOC honor any anonymity concerns and work to address them with adopters, the final approval of the summarized response is a mechanism that allows us to confirm with the adopter their comfort and approval of the final content intended for publication and make any corrections they feel are warranted. TOC members are free to bring in the Chair or Vice Chair to assist in addressing such concerns.

To ease scheduling with adopters, TOC members are recommended to either include set aside dates/times for adopters as part of the initial email, or to provide a scheduling link to expedite scheduling and avoid delay.

It is anticipated that a minimum of three adopter interviews are required to appropriately ascertain adoption of a project. However in the course of interviewing, you may find that you need additional adopters or types of adopters to be interviewed.

> **Note on Interview Participant Composition:** TOC members should prioritize interviewing adopter contacts who work with the project in an operational or user capacity. If the primary contact is a project maintainer or holds a project governance role, ask them to include a colleague from the same organization who uses the project operationally. Contributors from adopting organizations who have become maintainers can reflect strong community engagement, which is a positive signal. However, if such interviewees make up a meaningful proportion of the pool, overall results may skew overly positive. If you do not feel you have collected enough diverse perspectives, request interviews with additional adopters.

> **Reference Architecture or Case Study Substitution:** A project's inclusion in a published [TAB Reference Architecture](https://github.com/cncf/tab/blob/main/process/reference-architectures.md) or a published [CNCF Case Study](https://github.com/cncf/foundation/blob/main/policies-guidance/case-study-guidelines.md) may serve as a replacement for an individual adopter interview. Both represent validated, production-level adoption by the submitting organization and can provide equivalent signal to an adopter interview. When evaluating a Reference Architecture or Case Study, the TOC member should confirm that it demonstrates an independent organization's operational use of the project, rather than relying solely on the project team's or a vendor's account. It is at the TOC member's discretion who is leading the Due Diligence for Moving Levels to determine whether the Reference Architecture or Case Study meets the necessary standards and includes the requisite information the Due Diligence requires to determine whether the project meets the requirements of the level in process. Where it does, TOC members should note the Reference Architecture or Case Study in the adoption section of the Due Diligence in place of the interview summary it replaces.

#### Spec project adopter interviews

For specification projects, the adopter interview process differs from standard software projects. Because the primary deliverable is a normative document rather than a single codebase, "adopters" are the independent implementations of the spec and the organizations behind them.

When interviewing spec project adopters, the TOC should prioritize implementers from multiple independent implementations rather than users of a single reference implementation. In addition to the standard adopter interview questions, the following questions are relevant for spec project implementers:

- Which parts of the spec does your implementation cover?
- How do you track conformance with spec updates?
- Have you participated in the spec proposal process?
- What is your experience with interoperability testing against other implementations?
- How does your implementation handle areas where the spec is ambiguous or silent?

The goal is to assess whether the spec is implementable, adopted across organizations, and producing interoperable outcomes - not just whether one implementation has users.

For projects moving from Incubation to Graduation, if considerable time has passed since Incubation (according to the TOC's judgement), the TOC should refresh the Adopter interviews. This may be done by reaching out to previous interviewees, by engaging a new group of adopters for interviews, or some combination thereof. If the time period between Incubation and Graduation is fairly recent, the TOC member(s) should exercise their judgement in choosing to pursue additional interviews. That decision should be recorded with justification in the adoption section of the template.

#### Recording Responses

**Adopter interviews are expected to be interactive**. The [adopter questions template](toc-templates/template-adopter-questions.md) should serve as a starting point for questions when interviewing, however TOC members are expected to use their best judgement in asking questions, deep diving on responses, and crafting additional questions or skipping others. 

You may need to record the meeting to fully capture the responses or take sufficient notes that you can summarize the discussion and convey, with enough breadth, how the adopter is using the project, what environments (such as dev, test, prod), their engagement with the project, use, experience, and future plans.

TOC and TAB members are NOT to email questions to adopters in order for them to write in their responses. If there is a language or time zone challenge, TOC members are expected to inform the broader TOC and seek assistance.

#### Summary Approval

TOC members will summarize responses to the questions asked in a separate, non-public document until the Adopter approves the content. 

Once you have summarized the interview, typically about 5 paragraphs in length depending on the detail allowed to be conveyed, share the doc with the interviewee so they have permissions to make any changes. Provide the Interviewee with the doc, ask them to make any changes and provide approval two weeks from the date of sending. This provides you with a timeframe to follow-up with them so that it does not get lost or overlooked.

#### Add Interview Summaries to Due Diligence

Once approval is received, the summarized content is copied into the Due Diligence PR in the appropriate section. Linking to the summary is NOT recommended as it may convey identifying information in the history if not properly access controlled which will circumvent anonymity assurances.

### Summarizing the Due Diligence Evaluation

Once you have completed evaluating the project's implementation of the criteria and included the Adopter Interviews, you can summarize the overall evaluation.

Evaluation summary is composed of two parts: the Criteria and the Adoption. The evaluation summary should read like an executive summary of the Due Diligence, conveying any notable areas and the assessment of sufficient adoption for the level the project applied. An example structure is provided below:

> $TOCMEMBER conducted the Due Diligence of $PROJECT who applied for $LEVEL. The project [has/has not] completed the criteria that show its maturity at $LEVEL. The following criteria implementations are noteworthy to call out... $NOTABLES. The following actions were provided to the project that were considered blocking but since resolved... $BLOCKERS. The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project... $RECOMMENDATIONS.
> 
> The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY.
> 
> [The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS].

## TOC Internal Review

Once the TOC member(s) has completed the Due Diligence, they should create a PR in their personal TOC repo and share the link with the TOC for review. The TOC member(s) should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "TOC Review".

The TOC member(s) should craft a slack message thread that provides the project name, level applied, links to the PR and issue, and thread any additional call outs, questions, or potential issues warranting further discussion not otherwise captured in the PR. The internal review is expected to last 1 week, assuming no issues are brought up. 

The TOC member(s) is responsible for updating the project with the change in status for internal review.

### Reapplication

In the event a project was not ready to move levels after the Due Diligence was completed and the project has reapplied through an issue, the previous or new TOC member(s) assigned will initiate a new Due Diligence based on the previous one. The TOC should refresh the prior evaluations with corresponding dates to show changes and improvements and amend the evaluation statements accordingly. 

## Public Comment Period

Assuming no issues have come up during the TOC internal review, the TOC member may put the Due Diligence out for public comment. The TOC member(s) should move the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Public Comment".
 
TOC members are to leverage the [public comment template](toc-templates/template-dd-public-comment.md) and be mindful of the timeline to consider if a freeze is in effect or soon will be. All public comment messages are to be sent on the [TOC public mailing list](https://lists.cncf.io/g/cncf-toc/topics).  Once sent, they should be linked on the PR and the project notified.

## Voting

Once the public comment period is complete, the TOC member needs to adjudicate any responses and comments on the PR. Provided no new or blocking information has been shared, the TOC member should request CNCF Staff to initiate voting on the PR, at which point the project's card on the [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) to "Voting". 

TOC members, especially those who filed the PR, need to record their vote by the thumbs up or thumbs down on the PR comment where gitvote was initiated.

## Press Coordination and Completion

Once voting has passed, the CNCF Staff take over the remainder of the process, coordinating press releases and announcements for the project. CNCF Staff will update the project' card on the  [Application to Move Levels board](https://github.com/orgs/cncf/projects/27/views/9) accordingly.
