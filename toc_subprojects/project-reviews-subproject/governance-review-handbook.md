# Governance Reviewer Role Handbook

# Goals

The goal of the Governance Review process is to help projects improve their own governance.  As a reviewer, it is your goal to:

1. Perform one or more Governance Reviews in order to help projects complete their requirements for matriculation  
2. Assist projects in identifying areas where their governance could use improvement and carrying out that improvement  

A successful Governance Review results in projects discovering, and fixing, issues around their governance processes or documentation, and to make a plan to gradually improve their governance continuously over the life of the project.

# Requirements

Governance reviewers should fulfill the following requirements:

* More than one year experience in any leadership position in a public open source project, or some other equivalent community management experience  
* Familiarity with the CNCF governance and matriculation requirements  
* Commitment to a principled impartial review  
  * This may prevent reviewing projects where the reviewer has a conflict of interest (see below)
* Availability to spend 2-4 hours performing a review, and complete it in a timely fashion

# Process

The following is the process for initiating, executing and completing a governance review:

### Getting Started

* Projects request a governance review by filling in [this form](https://github.com/cncf/toc/issues/new?template=gov-review.yml)
    * Open requests are labeled ['governance/review'](https://github.com/cncf/toc/issues?q=is%3Aissue%20state%3Aopen%20label%3Areview%2Fgovernance).
    * If the review is triggered by a matriculation application then the reviewer may open the issue themselves or may request the project do so.
* A reviewer should comment on the governance issue to note that they are starting the review. A simple `/assign @githubname` is sufficient - the purpose is to notify readers.
* The review must be based on a copy of the template at [github.com/cncf/toc/toc_subprojects/project-reviews-subproject/governance-review-template.md](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/governance-review-template.md). There are markdown comments in that template not shown in the rendered version in GitHub so be sure to copy the [raw](https://raw.githubusercontent.com/cncf/toc/refs/heads/main/toc_subprojects/project-reviews-subproject/governance-review-template.md) version.
   * Reviewers can collaborate in any efficient way, such as by opening a page in [notes.cncf.io](https://notes.cncf.io) or sharing a Google Doc.

### Executing the Review

* Questions and answers required for the governance review may have been answered in a matriculation application for a project. Be sure to review that first. The reviewer's job is to a) review and confirm what has been submitted and b) research the project's repos and structure to be sure answers are accurate and complete.
* Questions, issues and clarifications required during the process of the review should be put in the Governance Review issue so that they are public for future viewers and TOC members to consider.
* At the reviewers' discretion they may contact project submitters and maintainers in CNCF Slack for deeper conversations, but remember to record significant items in the public issue.

### Completing the Review

* When the assessment has been completed it must be submitted as a PR to the TOC repository under [projects](https://github.com/cncf/toc/tree/main/projects). Each project has a folder within which a "governance-review" folder should exist or be created, as in `projects/crossplane/governance-review`. The review should be saved as a dated markdown document in that folder, e.g., `projects/crossplane/governance-review/2025-10-10.md`.
* The assessment is a current point-in-time snapshot of the project's governance. No need to wait for things to be fixed to get the PR merged.
* The PR can only be merged after the TOC Liaison approves the PR.

# Conflicts of Interest

The following kinds of projects may impose a conflict of interest on reviewers, so in general each reviewer should avoid taking on reviews that would cause them.

* The reviewer's employer is the primary sponsor of the project
* The project competes with a project in which the reviewer has a leadership role
* The reviewer has had personal conflicts with one or more leaders of the project
* The reviewer has some other close personal relationship with one or more leaders of the project

If you are assigned a project that triggers a conflict of interest, please alert one of the leads for the Governance WG so that you can be re-assigned.
