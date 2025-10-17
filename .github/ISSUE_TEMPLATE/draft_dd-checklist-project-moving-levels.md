---
name: Project Level Change Checklist
about: This template outlines the steps to be completed when a project moves between maturity levels (e.g., Sandbox to Incubating).
title: "[LEVEL CHANGE] $PROJECT"
labels:
- kind/dd
- toc
---

# Project Moving Levels Checklist

This template outlines the administrative steps required when a project moves from one maturity level to another (e.g., Sandbox to Incubating, Incubating to Graduated).

- **Project Name:** $PROJECT
- **New Maturity Level:** $NEW_LEVEL
- **Project moving levels PR:** #

---

## Phase I: Project Application and Initial Triage

#### Project Submission
- [ ] Project has submitted an Issue on the [TOC repo](https://github.com/cncf/toc/issues) detailing conformance to criteria
- [ ] Application added to [TOC Project Board](https://github.com/orgs/cncf/projects/27/views/9)

#### Initial Triage/Evaluation (TOC Member)
All TOC members assist in triaging new applications in the "new" column.

**Adopters**
- [ ] Verify Adoption Assertion includes [Adopters file](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) link
- [ ] 5-7 adopters are in the [Adopter's form responses](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter)

**Process Principles**
- [ ] TAG meeting/[GTR](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/general-technical-questions.md)/[DTR](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/tag-domain-technical-review-template.md) is complete
- [ ] Assertion of vendor neutrality present
- [ ] Assertion of expectations review present

**Governance/Maintainers**
- [ ] Link to governance provided
- [ ] Link to maintainers file provided
- [ ] Active maintainer count listed
- [ ] Doc/code ownership documented
- [ ] Code of Conduct linked (should reference [CNCF CoC](https://github.com/cncf/foundation/blob/main/code-of-conduct.md))
- [ ] Subproject listing provided (if applicable)

**Community**
- [ ] Link to contributing file provided
- [ ] Communication channels documented
- [ ] Community meetings, recordings, and/or notes linked
- [ ] Active contributor information provided
- [ ] Documented efforts to attract more contributors

**Engineering**
- [ ] Project goals and objectives documented
- [ ] Supported use cases documented
- [ ] Link to roadmap/project board provided
- [ ] Architecture diagram and documentation linked
- [ ] Release process documented

**Security**
- [ ] Link to [joint assessment](https://tag-security.cncf.io/community/assessments/) (if available)
- [ ] Link to security.md with process to report issues
- [ ] Link to [OpenSSF Scorecard](https://securityscorecards.dev/) or other repo hardening evidence
- [ ] Security report resolution process and roles documented
- [ ] Link to completed [self-assessment](https://tag-security.cncf.io/community/assessments/)
- [ ] [OpenSSF Best Practices Badge](https://bestpractices.coreinfrastructure.org/) is "passing" and 100%

**Ecosystem**
- [ ] Link to adopters file provided
- [ ] Integrations/compatibility information documented

#### Handle Blockers
- [ ] If criteria missing: add comment, close application, affix "not-ready" label
- [ ] Move issue to "Not-Ready-Will Return" column

#### Triage Completion
- [ ] If ready: add comment, affix "ready" label
- [ ] Move issue from "new" to "ready for assignment" column
- [ ] Update application with comment detailing outstanding work, next steps, and estimated timeframes

## Phase II: TOC Assignment and Kick-off

#### Assignment
- [ ] TOC member stepped forward to sponsor the application
- [ ] [Conflict of interest](https://github.com/cncf/toc/blob/main/operations/dd-toc-guide.md#conflicts-of-interest) check completed
- [ ] Co-sponsor assigned if conflict exists
- [ ] TOC member(s) assigned self to the Issue
- [ ] Issue moved to "Assigned" on the [board](https://github.com/orgs/cncf/projects/27/views/9)

#### Kick-off
- [ ] Engaged project maintainers/leadership (via comment, Slack, or email)
- [ ] Confirmed two CNCF staff members included in communications
- [ ] Kick-off meeting scheduled (or asynchronous kick-off initiated)
- [ ] [Kick-off document](https://github.com/cncf/toc/blob/main/operations/toc-templates/template-kickoff-notes.md) created to capture expectations, decisions, and timelines
- [ ] Informed project to verify compliance with [CNCF licensing policy](https://github.com/cncf/foundation/blob/master/charter.md) (Section 11) and [allowed third-party licenses](https://github.com/cncf/foundation/blob/main/allowed-third-party-license-policy.md)
- [ ] Issue moved to "In Due Diligence"

**Communication Links:**
- Kick-off Document: [Add link]
- Slack Channel/Thread: [Add link]
- Email Thread: [Add link]

## Phase III: Completing Due Diligence

#### Create Due Diligence PR
- [ ] Created DD PR using appropriate template: [Incubation](https://github.com/cncf/toc/blob/main/operations/toc-templates/template-dd-pr-incubation.md) or [Graduation](https://github.com/cncf/toc/blob/main/operations/toc-templates/template-dd-pr-graduation.md)
- [ ] DD PR Link: [Add link to PR]

#### Evaluate Criteria
- [ ] Independently reviewed project's assertions against publicly available artifacts
- [ ] Recorded all deviations in the Due Diligence PR
- [ ] Recorded implementation notes in the Due Diligence PR
- [ ] Recorded recommendations in the Due Diligence PR

#### Sub-project Review
- [ ] Evaluated sub-projects packaged with the main project
- [ ] Indicated any un-evaluated sub-projects and the reason
- [ ] Verified main project has sufficient governance of sub-projects

#### Licensing Issues
- [ ] Reviewed project licensing compliance
- [ ] If issues found: notified project to file [license exception request](https://github.com/cncf/foundation/issues/new/choose)
- [ ] License Exception Issue: [Add link if applicable]

#### Security Audit Review (Graduation Only)
- [ ] Reviewed security audit results
- [ ] Confirmed all critical findings resolved
- [ ] Confirmed all high findings resolved
- [ ] Noted any additional recommendations
- [ ] Security Audit Link: [Add link if applicable]

#### Specification Project Check (If Applicable)
- [ ] Confirmed at least one implementation exists
- [ ] Confirmed implementation has sufficient adoption


## Phase IV: Finalizing DD and Adopter Interviews

#### Engage TAGs (Optional)
- [ ] Requested update/summary from previously engaged TAG (if beneficial)
- [ ] TAG feedback received and incorporated
- [ ] TAG Feedback Link: [Add link if applicable]

#### Select and Reach Out to Adopters
- [ ] Identified 5-7 potential adopters from project's list
- [ ] Used [Adopter Interview Request template](https://github.com/cncf/toc/blob/main/operations/toc-templates/template-adopter-interview-request.md)
- [ ] Included interested TAB members in outreach
- [ ] Offered anonymity option to adopters

#### Schedule and Conduct Interviews
- [ ] Scheduled minimum of three adopter interviews (30-60 mins each)
- [ ] Interview 1 completed - Date: [Add date]
- [ ] Interview 2 completed - Date: [Add date]
- [ ] Interview 3 completed - Date: [Add date]
- [ ] Additional interviews (if needed): [Add details]
- [ ] Used [core questions](https://github.com/cncf/toc/blob/main/operations/toc-templates/template-adopter-questions.md) as starting point
- [ ] Took sufficient notes or recorded meetings

#### Record and Approve Responses
- [ ] Summarized each interview (~5 paragraphs) in separate, non-public documents
- [ ] Interview 1 summary shared with interviewee for approval
- [ ] Interview 2 summary shared with interviewee for approval
- [ ] Interview 3 summary shared with interviewee for approval
- [ ] All summaries approved within two weeks

#### Add to Due Diligence
- [ ] Copied approved summaries into Due Diligence PR (not linked)
- [ ] Wrote Criteria summary statement (executive summary style)
- [ ] Wrote Adoption summary statement (executive summary style)

**Adopter Interview Tracking:**
| Adopter | Contacted | Scheduled | Completed | Summary Approved |
|---------|-----------|-----------|-----------|------------------|
| 1.      | [ ]       | [ ]       | [ ]       | [ ]              |
| 2.      | [ ]       | [ ]       | [ ]       | [ ]              |
| 3.      | [ ]       | [ ]       | [ ]       | [ ]              |
| 4.      | [ ]       | [ ]       | [ ]       | [ ]              |
| 5.      | [ ]       | [ ]       | [ ]       | [ ]              |
| 6.      | [ ]       | [ ]       | [ ]       | [ ]              |
| 7.      | [ ]       | [ ]       | [ ]       | [ ]              |


## Phase V: Review and Voting

#### Internal Review
- [ ] Created PR in Private TOC repo for internal review
- [ ] Shared PR link with TOC
- [ ] Issue moved to "TOC Review"
- [ ] Crafted Slack message thread with project name, level, links, and call-outs
- [ ] Internal review lasted ~1 week
- [ ] Notified project of status change to internal review
- [ ] Addressed any TOC feedback
- [ ] Private TOC Repo PR: [Add link]

#### Public Comment
- [ ] Put Due Diligence out for public comment using [template](https://github.com/cncf/toc/blob/main/operations/toc-templates/template-dd-public-comment.md)
- [ ] Sent to [TOC public mailing list](https://lists.cncf.io/g/cncf-toc/topics)
- [ ] Issue moved to "Public Comment"
- [ ] Public comment period: 2 weeks until [Add end date]
- [ ] Linked public comment message on PR
- [ ] Notified project of public comment period
- [ ] Public Comment Link: [Add link]

#### Adjudicate Comments
- [ ] Reviewed and adjudicated all responses and comments on PR
- [ ] No new blocking information identified
- [ ] Made necessary updates based on feedback

#### Voting
- [ ] Requested CNCF Staff to initiate voting
- [ ] Issue moved to "Voting"
- [ ] TOC member(s) recorded vote on PR
- [ ] Voting completed successfully
- [ ] GitVote Comment: [Add link]

#### Completion - CNCF Staff tasks
- [ ]  TOC vote passed and moving levels PR merged
- [ ] Email Formation team to move the project's level in PCC.
- [ ] Move the project in Google Group mailing lists from its current group to the new one (e.g., from `cncf-sandbox-projects` to `cncf-incubating-projects`to `cncf-graduated-projects`).
- [ ] Update the project's level and maintainer list in the aggregated [CNCF maintainer list](https://maintainers.cncf.io/). https://github.com/cncf/foundation/pull/1153
- [ ] Move the project to its new level on the [CNCF Landscape](https://github.com/cncf/landscape/blob/master/landscape.yml). https://github.com/cncf/landscape/pull/4553
- [ ] Move artwork in [cncf/artwork](https://github.com/cncf/artwork) https://github.com/cncf/artwork/pull/600
- [ ] Sent email to CNCF PR team for marketing the change


### Key Documents and Links

- **Project Application Issue:** [Add link]
- **Kick-off Document:** [Add link]
- **Due Diligence PR:** [Add link]
- **Private TOC Repo PR:** [Add link]
- **Security Audit:** [Add link if applicable]
- **TAG Feedback:** [Add link if applicable]
- **License Exception:** [Add link if applicable]

### Communication Channels

- **Slack Channel/Thread:** [Add link]
- **Email Thread:** [Add summary or link]
- **Other:** [Add any other communication channels]

### Reference Documentation

- [TOC Due Diligence Guide](https://github.com/cncf/toc/blob/main/operations/dd-toc-guide.md)
- [Moving Levels Process](https://github.com/cncf/toc/blob/main/process/README.md)
- [Graduation Criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md)
- [TOC Project Board](https://github.com/orgs/cncf/projects/27/views/9)
- [TOC Principles](https://github.com/cncf/toc/blob/main/PRINCIPLES.md)
