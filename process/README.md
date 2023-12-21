# I. Overview

This policy describes the Cloud Native Computing Foundation (CNCF) project life cycle process, from sandbox to archival and more. It describes the stages a project may be admitted under and what the criteria and expectations are for a given stage, as well as the acceptance criteria for a project to move from one stage to another. It also describes the Annual Review Process through which those changes will be evaluated and made.

Project progression, movement from one stage to another, allows projects to participate at the level that is most appropriate for them given where they are in their lifecycle.


# II. Stages - Definitions & Expectations

CNCF projects have a maturity level of sandbox, incubating, or graduated. Archived is for projects no longer in active development. The maturity level is a signal by CNCF as to what sorts of enterprises should be adopting different projects. Projects increase their maturity by demonstrating their sustainability to CNCF’s Technical Oversight Committee: that they have adoption, a healthy rate of changes, committers from multiple organizations, have adopted the CNCF Code of Conduct, and have achieved and maintained the OpenSSF Best Practices Badge.

![Project Stages](https://github.com/cncf/toc/blob/main/process/project-stages.png)

## Sandbox:

The CNCF Sandbox is the entry point for early stage projects and has four goals:

* Encourage public visibility of experiments or other early work that can add value to the CNCF mission and build the ingredients of a successful Incubation level project
* Facilitate alignment with existing projects if (and only if) this is desired
* Nurture projects (e.g. via CNCF Service Desk requests)
* Remove possible legal and governance obstacles to adoption and contribution by ensuring all projects adhere to CNCF legal, code of conduct and IP Policy requirements


## Incubating:
Incubating projects have access to all resources listed at https://cncf.io/services-for-projects, and have a presence at CloudNativeCon.


## Graduated:
Graduated projects signal the highest level of maturity for a CNCF project.


## Archived:
Archived projects are no longer in active development and are only archived after a TOC vote.

# III. Project Proposal Process

Introduction:
This governance policy sets forth the proposal process for projects to be accepted into the CNCF.
The process is the same for both existing projects which seek to move into the Foundation, and new projects to be formed within the CNCF.

## Project Proposal Requirements:

## Sandbox Projects:

Projects being submitted to the CNCF at the sandbox level are intended to be the entry point for early stage projects and are not required to undergo due diligence.

Sandbox projects should be early-stage projects that the CNCF TOC believes warrant experimentation.

* New projects that are designed to extend one or more CNCF projects with functionality or interoperability libraries.
* Independent projects that fit the CNCF mission and provide potential for a novel approach to existing functional areas (or are an attempt to meet an unfulfilled need)
* Projects commissioned or sanctioned by the CNCF, including initial code for CNCF WG collaborations, and "experimental" projects
* Any project that realistically intends to join CNCF Incubation in future and wishes to lay the foundations for that

To apply for inclusion into the Sandbox, projects should use apply at [sandbox.cncf.io](https://sandbox.cncf.io).
The TOC will review on a rotating basis, currently every two months as of June 2020.


## Project Graduation Process: Sandbox to Incubating

Incubating projects are required to undergo due diligence as a part of the process to move from Sandbox to Incubation. Due Diligence is driven by a TOC sponsor, with two weeks for public comment before a vote is called.

Criteria:
To be accepted to incubating stage, a project must meet the sandbox stage requirements plus:

* Document that it is being used successfully in production by at least three independent adopters which, in the TOC’s judgement, are of adequate quality and scope.
* Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.
* Demonstrate a substantial ongoing flow of commits and merged contributions.
* Since these metrics can vary significantly depending on the type, scope and size of a project, the TOC has final judgement over the level of activity that is adequate to meet these criteria
* A clear versioning scheme.
* Specifications must have at least one public reference implementation.

![Incubating](https://github.com/cncf/toc/blob/main/process/incubation-process.png)

Projects currently in progress for consideration at the Incubating stage are tracked: https://github.com/orgs/cncf/projects/27/views/9

## (3) Project Graduation Process: Incubating to Graduation
Projects that wish to move from Incubating to Graduation should open a PR confirming the following criteria:
* Have committers from at least two organizations.
* Have achieved and maintained a [OpenSSF Best Practices Badge](https://bestpractices.coreinfrastructure.org/).
* Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.
* Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.
* Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.
* Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.
* Please include a short one-page narrative based off the Graduation template, no more than 500 words.

Projects moving from incubation to graduation are tracked here: https://github.com/orgs/cncf/projects/27/views/9

## (4) Archiving Projects

Open source projects have a lifecycle and there are times that projects become inactive due to a variety of reasons. There are also cases where a project may no longer want to be supported by the TOC, or the TOC may no longer wish to recommend the use of a project.
Archiving Criteria
When voting on a proposal to archive a project, TOC members may wish to consider whether the project continues to meet the criteria for CNCF acceptance. The TOC may also look at activity levels in the project (https://all.devstats.cncf.io/d/53/projects-health-table?orgId=1), although it is important to note that there is a difference between a mature project that doesn't get much attention any more but is stable, versus a project that is inactive.

### Voting Process

To archive a project:
* A proposal must be put forth to the TOC repo
* The TOC will inform the project maintainers, CNCF End-User community and wider community of all archiving proposals
* The proposal must remain open for at least 2 weeks of discussion after the maintainers are informed.
* A vote must be finalized with 2/3 approval from the TOC

### Archiving Process

See the [Archiving process](archiving.md).

# IV. Annual Review Process

The TOC is deprecating the current annual review process as of September 2023 in favor of a proposed automated review system. 