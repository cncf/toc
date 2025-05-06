# Governance Reviewer Role Handbook

WIP: This guide is not complete, but the instructions it contains are a starting point.

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

For anybody interested in doing governance reviews, I wanted to write a summary of how to get started with that work:

* Projects send a request by filling in a [form](https://github.com/cncf/tag-contributor-strategy/issues/new?assignees=jberkus%2Cgeekygirldawn&labels=wg%2Fgovernance&projects=&template=governance-review-request.yaml&title=%5BGovernance+Review%5D%3A+PROJECT+NAME) (some open requests are [here](https://github.com/cncf/tag-contributor-strategy/issues?q=is%3Aissue+is%3Aopen+label%3Awg%2Fgovernance+%22governance+review%22))  
* That form should have pointers to governance documents, activities and other things necessary for your assessment.  
* The review should follow this [template](https://github.com/cncf/tag-contributor-strategy/blob/main/governance/reviews/template.md). There are markdown comments in that template, not shown in GitHub UI. See the [raw](https://raw.githubusercontent.com/cncf/tag-contributor-strategy/main/governance/reviews/template.md) file for that.  
* When the template is filled with the assessment, the output should be in a PR, to be put in the TAG repository under [governance/assessments/projects](https://github.com/cncf/tag-contributor-strategy/tree/main/governance/assessments/projects).  Each project gets a folder, and reviews are saved under the date the first draft of the review was completed.
* The assessment is to be done for a snapshot of a project governance. No need to wait for things to be fixed to get the PR merged.  
* The PR can only be merged after the TOC Liaison approves the PR.

You can collaborate with other TAG members using Google Docs.  
The discussions (disagreements about a certain governance aspect that you find unsatisfactory for example) with projects should be done on the PR when it is opened.  
That doesn't mean you can't interact with the projects of course and ask them questions and get their help to find out where something is (e.g. meeting recordings, governance docs, etc)

# Conflicts of Interest

The following kinds of projects may impose a conflict of interest on reviewers, so in general each reviewer should avoid taking on reviews that would cause them.

* The reviewer's employer is the primary sponsor of the project
* The project competes with a project in which the reviewer has a leadership role
* The reviewer has had personal conflicts with one or more leaders of the project
* The reviewer has some other close personal relationship with one or more leaders of the project

If you are assigned a project that triggers a conflict of interest, please alert one of the leads for the Governance WG so that you can be re-assigned.