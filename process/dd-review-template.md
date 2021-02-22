# Due Diligence Project Review Template
This page provides a template to help those leading or contributing to due diligence exercises performed by or on behalf of the Technical Oversight Committee of the CNCF. Please also read the [Due Diligence Guidelines](https://github.com/cncf/toc/blob/main/process/due-diligence-guidelines.md). 

## Introduction
The decision to graduate or promote a project depends on the TOC sponsors of the project performing and documenting the evaluation process in deciding upon initial or continued inclusion of projects through a Technical Due Diligence ('Tech DD') exercise. Ultimately the voting members of the TOC will, on the basis of this and other information, vote for or against the inclusion of each project at the relevant time.

## Technical Due Diligence
### Primary Goals
To enable the voting TOC members to cast an informed vote about a project, it is crucial that each member is able to form their own opinion as to whether and to what extent the project meets the agreed upon criteria for sandbox, incubation or graduation. As the leader of a DD, your job is to make sure that they have whatever information they need, succinctly and readily available, to form that opinion.

As a secondary goal, it is in the interests of the broader CNCF ecosystem that there exists some reasonable degree of consensus across the community regarding the inclusion or otherwise of projects at the various maturity levels. Making sure that the relevant information is available, and any disagreement or misunderstanding as to it's validity are ideally resolved, helps to foster this consensus.

## Statement of CNCF Alignment to TOC Principles
1. Project is self-governing
2. Is there a documented Code of Conduct that adheres to the CNCF guidelines?
3. Does the project have production deployments that are high quality and high-velocity? (for incubation and graduated projects).
(Sandbox level projects are targeted at earlier-stage projects to cultivate a community/technology)
4. Is the project committed to achieving the CNCF principles and do they have a committed roadmap to address any areas of concern raised by the community?
5. Document that the project has a fundamentally sound design without obvious critical compromises that will inhibit potential widespread adoption
6. Document that the project is useful for cloud native deployments & degree that it's architected in a cloud native style
7. Document that the project has an affinity for how CNCF operates and understand the expectation of being a CNCF project.

## Review of graduation criteria and desired cloud native properties
/* Use appropriate Section */

### Sandbox Graduation (Exit Requirements)
1. Document that it is being used successfully in production by at least three independent end users which with focus on adequate quality and scope defined.
2. Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.
3. Demonstrate a substantial ongoing flow of commits and merged contributions.

### Incubating Stage Graduation (Exit Requirements)
1. Document that it is being used successfully in production by at least three independent end users which with focus on adequate quality and scope defined.
2. Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.
3. Demonstrate a substantial ongoing flow of commits and merged contributions.
4. Have committers from at least two organizations.
5. Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.
6. Adopted the CNCF Code of Conduct.
7. Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.
8. Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

### Documentation of CNCF Alignment (if not addressed above):
* name of project (must be unique within CNCF)
* project description (what it does, why it is valuable, origin and history)
* statement on alignment with CNCF charter mission
* sponsor from TOC (sponsor helps mentor projects)
* license (charter dictates Apache 2 by default)
* source control (GitHub by default)
* external dependencies (including licenses)
* release methodology and mechanics
* community size and any existing sponsorship

## Technical
* An architectural, design and feature overview should be available. (add link)
* What are the primary target cloud-native use cases? Which of those:
  * Can be accomplished now.
  * Can be accomplished with reasonable additional effort (and are ideally already on the project roadmap).
  * Are in-scope but beyond the current roadmap.
  * Are out of scope.
* What are the current performance, scalability and resource consumption bounds of the software? Have these been explicitly tested? Are they appropriate given the intended usage (e.g. agent-per-node or agent-per-container need to be lightweight, etc)?
* What exactly are the failure modes? Are they well understood? Have they been tested? Do they form part of continuous integration testing? Are they appropriate given the intended usage (e.g. cluster-wide shared services need to fail gracefully etc)?
* What trade-offs have been made regarding performance, scalability, complexity, reliability, security etc? Are these trade-offs explicit or implicit? Why? Are they appropriate given the intended usage? Are they user-tunable?
* What are the most important holes? No HA? No flow control? Inadequate integration points?
* Code quality. Does it look good, bad or mediocre to you (based on a spot review). How thorough are the code reviews? Substance over form. Are there explicit coding guidelines for the project?
* Dependencies. What external dependencies exist, do they seem justified?
* What is the release model? Versioning scheme? Evidence of stability or otherwise of past stable released versions?
* What is the CI/CD status? Do explicit code coverage metrics exist? If not, what is the subjective adequacy of automated testing? Do different levels of tests exist (e.g. unit, integration, interface, end-to-end), or is there only partial coverage in this regard? Why?
* What licensing restrictions apply? Again, CNCF staff will handle the full legal due diligence.
* What are the recommended operational models? Specifically, how is it operated in a cloud-native environment, such as on Kubernetes?

## Project
* Do we believe this is a growing, thriving project with committed contributors?
* Is it aligned with CNCF's values and mission?
* Do we believe it could eventually meet the graduation criteria?
* Should it start at the sandbox level or incubation level?
* Does the project have a sound, documented process for source control, issue tracking, release management etc.
* Does it have a documented process for adding committers?
* Does it have a documented governance model of any kind?
* Does it have committers from multiple organizations?
* Does it have a code of conduct?
* Does it have a license? Which one? Does it have a CLA or DCO? Are the licenses of its dependencies compatible with their usage and CNCF policies? CNCF staff will handle the full legal due diligence.
* What is the general quality of informal communication around the project (slack, github issues, PR reviews, technical blog posts, etc)?
* How much time does the core team commit to the project?
* How big is the team? Who funds them? Why? How much? For how long?
* Who are the clear leaders? Are there any areas lacking clear leadership? Testing? Release? Documentation? These roles sometimes go unfilled.
* Besides the core team, how active is the surrounding community? Bug reports? Assistance to newcomers? Blog posts etc.
* Do they make it easy to contribute to the project? If not, what are the main obstacles?
* Are there any especially difficult personalities to deal with? How is this done? Is it a problem?
* What is the rate of ongoing contributions to the project (typically in the form of merged commits).

## Users
* Who uses the project? Get a few in-depth references from 2-4 of them who actually know and understand it.
* What do real users consider to be its strengths and weaknesses? Any concrete examples of these?
* Perception vs Reality: Is there lots of buzz, but the software is flaky/untested/unused? Does it have a bad reputation for some flaw that has already been addressed?

## Context
* What is the origin and history of the project?
* Where does it fit in the market and technical ecosystem?
* Is it growing or shrinking in that space? Is that space growing or shrinking?
* How necessary is it? What do people who don't use this project do? Why exactly is that not adequate, and in what situations?
* Clearly compare and contrast with peers in this space. A summary matrix often helps. Beware of comparisons that are too superficial to be useful, or might have been manipulated so as to favor some projects over others. Most balanced comparisons will include both strengths and weaknesses, require significant detailed research, and usually there is no hands-down winner. Be suspicious if there appears to be one.
