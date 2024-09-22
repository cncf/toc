# How TAGs Help with Project Reviews and Engagement

CNCF Technical Advisory Groups help with Project Reviews for new projects interested in joining the CNCF at the Sandbox or Incubation level, as well as projects moving between maturity levels: to Incubation or Graduation.

Projects themselves may engage with the TAGs:
* when exploring engagement within the cloud native ecosystem
* prior to submitting their application for CNCF consideration
* after submitting their application for CNCF consideration
* for guidance in moving levels within the CNCF
* upon direction by the CNCF Technical Oversight Committee (TOC) or Linux Foundation staff

## Project Engagement

Project engagement within the Cloud Native ecosystem is a vital way to raise project visibility as well as gain insights from technical leaders in the CNCF community for project maintainers. 

As the Technical Advisory Groups (TAGs) are an entry point for many projects to engage with the Cloud Native ecosystem, TAGs are encouraged to guide projects to join the TAGs, Working Groups, collaborate with other projects as well as explore various avenues to further engage with the greater community.

## Project Reviews

To assist in evaluating projects applying to enter or moving levels within the CNCF, projects are encouraged to do one of the following:
- Present an overview and demo of the project to the CNCF Technical Advisory Group(s) (TAGs) or TAG Working Group(s) that their project aligns with. If the project is unsure how or where they align, they may ask in the [CNCF #toc slack channel](https://cloud-native.slack.com/archives/C0MP69YF4).

**or**

- Complete the [General Technical Review] (GTR) questions and open a PR to the project's directory in the TOC repo for the TAG and/or TOC to review instead of scheduling a project presentation.

_**Suggestion:** The GTR questionnaire may be kept in the project's GitHub organization as a living document and updated as needed. When applying to move levels, a PR with a copy of the GTR can then be opened to record the project's state at that point-in-time._


#### What to do if multiple TAGs are suggested

At times, project maintainers will be encouraged to interact with multiple TAGs, Working Groups and/or Kubernetes SIGs. In these instances, the project may complete the GTR questionnaire and request reviews from the different groups or work with TAG leadership to coordinate a shared project presentation to present to all requested parties.

### Scheduling

TAG leadership should proactively reach out to the project maintainers to schedule a project presentation, if warranted.

If an issue is created in the TAG repo - either automatically by the automation being implemented by LF staff - or manually by TOC, Project Maintainers or LF Staff, the expectation is that the TAG will assign someone to the issue as well as find a suitable time for the project to present. The TAG may suggest the project complete the General Technical Review questionnaire if the TAG does not have availability within 60 days.

## Presentation Guidance

Project maintainers interested in applying to join the CNCF - or move levels within the CNCF - should be prepared with slides or the General Technical Review questions and a short demo, while also reserving time for community engagement. 

Project maintainers may ask the TAG leadership for additional time. Projects may also record the presentation and demo separately for async discussion with the TAG if there are timezone and/or TAG resource conflicts. The presentation and demo should then be shared on the Project's application, if applicable.

#### Presentation Timing Guidance

Please see the table below for timing guidance. 

| Application Level | Presentation + Short Demo Time | TAG Questions and Discussion |
| -------- | -------- | -------- |
| Sandbox     | 20 Minutes     | 10 Minutes     |
| Incubation | 40 Minutes | 15 Minutes |
| Graduation | 40 Minutes | 15 Minutes |

### Presentation Agenda

The presentation being delivered to the TAG should be focused on technical specifics using the [General Technical Review questions] as the basis. The General Technical Review questions follow the cloud native software lifecycle day schemas and the TAG may ask for additional questions to be answered for the project presentation. *Please see the [GTR questions] for more information.*

Please see the table below for questions that should be addressed per level to satisfy the Sandbox, Incubation and Graduation TAG presentation requirements as well as the Incubation and Graduation engineering principles requirements:
* Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.
* Document what the project does, and why it does it - including viable cloud native use cases.

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

After the project has completed a presentation to the TAG, or completed the General Technical Review questionnaire, the assigned TAG leader should complete a domain technical review.

### Who May Complete a Review

A TAG Chair, Technical Lead or Working Group lead may choose to self-assign. Alternatively, a technical member of the community with domain knowledge may be encouraged to complete a review with TAG leadership oversight. 

LF staff are implementing automation to open an issue in the TAG repos as well as capture assignees on the project boards. TAGs are expected to regularly groom their backlog and assign members to project reviews.

Sandbox Project Review Assignments: https://github.com/orgs/cncf/projects/14/views/5

Incubation and Graduation Project Review Assignments: TBD

### What is needed to complete a review

The TAG member assigned to complete the domain technical review references: 
* the [original General Technical Review questionnaire]
* the project presentation or the completed GTR questionnaire
* the [TAG Domain Technical Review template]
* TAG Security [self] or [joint assessment]

### How long should a review take?

The TAG member should take no longer than an hour to complete a review. The review should be a brief synopsis of the project that captures key highlights as well as any glaring gaps that the project may have.

### Where is a review submitted?

The TAG project reviewer should submit the review as a comment on the project's application.

