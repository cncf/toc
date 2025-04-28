# TAG Project Reviews

v1.0

CNCF Technical Advisory Groups (TAGs) help with Project Reviews for new projects interested in joining the CNCF, as well as projects moving between maturity levels.

## Project Reviews

To assist in evaluating projects applying to enter or moving levels within the CNCF, projects are encouraged to do one of the following:
- Present an overview and demo of the project to the CNCF (TAGs) or TAG Working Group(s) that their project aligns with.

**or**

- Complete the [General Technical Review][gtrq] (GTR) questions and open a PR to the project's directory in the TOC repo for the TAG and/or TOC to review instead of scheduling a project presentation.

_**Suggestion:** The GTR questions may be kept in the project's GitHub organization as a living document and updated as needed. When applying to move levels, a PR with a copy of the GTR can then be opened to record the project's state at that point-in-time._


#### What to do if multiple TAGs are suggested

At times, project maintainers will be encouraged to interact with multiple TAGs, Working Groups and/or Kubernetes SIGs. In these instances, the project may complete the GTR questions and request reviews from the different groups or work with TAG leadership to coordinate a shared project presentation to present to all requested parties.

### Scheduling

TAG leadership should proactively reach out to the project maintainers to schedule a project presentation, if warranted.

If a project review issue is created in the TAG repo the expectation is that the TAG will assign someone to the issue as well as find a suitable time for the project to present.
The TAG may suggest the project complete the [General Technical Review questions][gtrq] if the TAG does not have availability within 60 days.

## Presentation Guidance

Project maintainers interested in applying to join the CNCF - or move levels within the CNCF - should be prepared with slides or the General Technical Review questions and a short demo, while also reserving time for community engagement. 

Project maintainers may ask the TAG leadership for additional time. Projects may also record the presentation and demo separately for async discussion with the TAG if there are timezone and/or TAG resource conflicts. The presentation and demo should then be shared on the Project's application, if applicable.

#### Suggested Presentation Timing Guidance

Please see the table below for timing guidance. 

| Application Level | Presentation + Short Demo Time | TAG Questions and Discussion |
| -------- | -------- | -------- |
| Sandbox     | 20 Minutes     | 10 Minutes     |
| Incubation | 40 Minutes | 15 Minutes |
| Graduation | 40 Minutes | 15 Minutes |

### Presentation Agenda

The presentation being delivered to the TAG should be focused on technical specifics using the [General Technical Review questions][gtrq] as the basis.
The General Technical Review questions follow the cloud native software lifecycle day schemas and the TAG may ask for additional questions to be answered for the project presentation. *Please see the [GTR questions][gtrq] for more information.*

Please see the table below for questions that should be addressed per level to satisfy the Sandbox, Incubation and Graduation TAG presentation requirements as well as the Incubation and Graduation engineering principles requirements.

**REMINDER:** Not all questions are relevant to all projects.
It's fully expected that projects will mark some as not-applicable.

| General Technical Review Questions | Sandbox | Incubation | Graduation |
| -------- | -------- | -------- | -------- |
| Day 0 - Scope     | required     | required     | required     |
| Day 0 - Usability     | required     | required     | required     |
| Day 0 - Design     | required     | required     | required     |
| Day 0 - Installation     | required     | required     | required     |
| Day 0 - Security     | required     | required     | required     |
| Day 1 - Project Configuration and Activation     | optional     | required     | required     |
| Day 1 - Project Enablement and Rollback     | optional     | required     | required     |
| Day 1 - Rollout, Upgrade and Rollback Planning     | optional     | required     | required     |
| Day 2 - Scalability/Reliability     | optional     | required     | required     |
| Day 2 - Observability     | optional     | required     | required     |
| Day 2 - Dependencies     | optional     | required     | required     |
| Day 2 - Troubleshooting     | optional     | required     | required     |
| Day 2 - Security     | optional     | required     | required     |

## Domain Technical Review

After the project has completed a presentation to the TAG, or completed the General Technical Review questions, the assigned TAG lead should complete a domain technical review.

Sandbox Project Review Assignments: https://github.com/orgs/cncf/projects/14/views/5

Incubation and Graduation Project Review Assignments: tbd

### Who should conduct project reviews?

A TAG or WG lead can either self-assign or be assigned by the TOC to conduct a project review.
They **MAY** delegate the review to a community member with domain expertise.

### How long should a review take?

It should take no longer than an hour to complete a review.
The review should be a brief synopsis of the project that captures key highlights as well as any **glaring** gaps that the project may have.

### When are domain project reviews requested?

A TAG domain project review should be completed when:
* Sandbox projects in the upcoming queue for TOC Sandbox Review
* when the TOC requests a domain technical review of a project applying for Incubation or Graduation

### What is needed to complete a review

The TAG member assigned to complete the domain technical review references: 
* the [original General Technical Review questions][gtrq]
* the project's presentation or the completed GTR questions
* the [TAG Domain Technical Review template](tag-domain-technical-review-template.md)
* TAG Security [self or joint assessment](https://tag-security.cncf.io/community/assessments/) (if applicable)

### Where is a review submitted?

The TAG project reviewer should submit the review as a comment on the project's application.
If the project is in the upcoming queue when applying for Sandbox, the TAG is requested to submit a review no later than two (2) weeks to allow for TOC review/clarification and to allow projects time to respond. 



[gtrq]: general-technical-questions.md

