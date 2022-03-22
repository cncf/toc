# Due Diligence Guidelines

This page provides guidelines to those leading or contributing to due
diligence exercises performed by or on behalf of the Technical
Oversight Committee of the CNCF.

## Introduction

Part of the evaluation process in deciding upon initial or continued
inclusion of projects into the CNCF is a Technical Due Diligence
('Tech DD') exercise.  Ultimately the voting members of the TOC will,
on the basis of this and other information, vote for or against the
inclusion of each project at the relevant time.

## Leading a Technical Due Diligence

### Primary Goals

To enable the voting TOC members to cast an informed vote about a
project, it is crucial that each member is able to form their own
opinion as to whether and to what extent the project meets the agreed
upon [criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.adoc) for
sandbox, incubation, or graduation. As the leader of a DD, your job
is to make sure that they have whatever information they need,
succinctly and readily available, to form that opinion.

As a secondary goal, it is in the interests of the broader CNCF
ecosystem that there exists some reasonable degree of consensus across
the community regarding the inclusion or otherwise of projects at the
various maturity levels.  Making sure that the relevant information is
available, and any disagreement or misunderstanding as to its
validity are ideally resolved, helps to foster this consensus.

### Where to start

* Make sure you are clear on the [TOC Principles](https://github.com/cncf/toc/blob/main/PRINCIPLES.md),
  the [project proposal process](https://github.com/cncf/toc/blob/main/process/project_proposals.adoc),
  the [graduation criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.adoc),
  and the [desired cloud native properties](https://www.cncf.io/about/charter/).  The project sponsor (a member
  of the TOC) should have assisted in crafting the proposal to explain why it is a good fit for the CNCF. If anything is
  unclear to you, reach out to the project sponsor or, failing that, the TOC mailing list for advice.
* Make sure you read, in detail, the relevant [project proposal](https://github.com/cncf/toc/tree/main/proposals).
  This will usually be in the form of an [open pull request](https://github.com/cncf/toc/pulls).
  Consider holding off on commenting on the PR until you have completed the next three steps.
* Take a look at some [previous submissions](https://github.com/cncf/toc/pulls?utf8=%E2%9C%93&q=is%3Apr)
  (both successful and unsuccessful) to help calibrate your expectations.
* Verify that all of the basic [project proposal requirements](https://github.com/cncf/toc/blob/main/process/project_proposals.adoc) have been provided.
* Do as much reading up as you need to (and consult with experts in the specific field) in order to familiarize yourself with the technology
  landscape in the immediate vicinity of the project (and do not only use the proposal and that project's documentation as a guide in this regard).
* At this point, you should have a very clear technical idea of what exactly the project actually does and does not do, roughly how it compares with and differs from
  similar projects in its technology area, and/or a set of unanswered questions in those regards.
* Go through the [graduation criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.adoc) and for each item,
  decide for yourself whether or not you have enough information to make a strong, informed call on that item.
  * If so, write it down, with motivation.
  * If not, jot down what information you feel you are missing.
  * Also take note of what unanswered questions the community might have posted in the PR review that you consider
    to be critically important.

### Some example questions that will ideally need clear answers

Most of these should be covered in the project proposal document.  The
due diligence exercise involves validating any claims made there,
verifying adequate coverage of the topics, and possibly summarizing
the detail where necessary.

#### Technical

* An architectural, design, and feature overview should be available.
  ([example](https://github.com/docker/notary/blob/master/docs/service_architecture.md),
  [example](https://github.com/docker/notary/blob/master/docs/command_reference.md))
* What are the primary target cloud native use cases?  Which of those:
  * Can be accomplished now.
  * Can be accomplished with reasonable additional effort (and are ideally already on the project road map).
  * Are in-scope but beyond the current road map.
  * Are out of scope.
* What are the current performance, scalability, and resource consumption bounds of the software?  Have these been explicitly tested?
  Are they appropriate given the intended usage (e.g. agent-per-node or agent-per-container need to be lightweight, etc)?
* What exactly are the failure modes?  Are they well understood?  Have they been tested?  Do they form part of continuous integration testing?
  Are they appropriate given the intended usage (e.g. cluster-wide shared services need to fail gracefully etc)?
* What trade-offs have been made regarding performance, scalability, complexity, reliability, security etc?  Are these trade-offs explicit or implicit?
  Why?  Are they appropriate given the intended usage?  Are they user-tunable?
* What are the most important holes? No high availability? No flow control? Inadequate integration points?
* Code quality.  Does it look good, bad, or mediocre to you (based on a spot review).  How thorough are the code reviews? Substance over form.
  Are there explicit coding guidelines for the project?
* Dependencies.  What external dependencies exist; do they seem justified?  Note: all core dependencies should be listed in the document along with the details of relevant repositories.
* What is the release model?  Versioning scheme?  Evidence of stability or otherwise of past stable released versions?
* What is the CI/CD status?  Do explicit code coverage metrics exist? If not, what is the subjective adequacy of automated testing?
  Do different levels of tests exist (e.g. unit, integration, interface, end-to-end), or is there only partial coverage in this regard?  Why?
* What licensing restrictions apply?  Again, CNCF staff will handle the full legal due diligence.
* What are the recommended operational models?  Specifically, how is it operated in a cloud native environment, such as on Kubernetes?

#### Project

The key high-level questions that the voting TOC members will be looking to have answered are (from the [graduation criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.adoc)):

* Do we believe this is a growing, thriving project with committed contributors?
* Is it aligned with CNCF's values and mission?
* Do we believe it could eventually meet the graduation criteria?
* Should it start at the sandbox level or incubation level?

Some details that might inform the above include:

* Does the project have a sound, documented process for source control, issue tracking, release management etc.
* Does it have a documented process for adding committers?
* Does it have a documented governance model of any kind?
* Does it have committers from multiple organizations?
* Does it have a code of conduct?
* Does it have a license? Which one? Does it have a CLA or DCO? Are the licenses of its dependencies compatible with their usage and CNCF policies?
  CNCF staff will handle the full legal due diligence.
* What is the general quality of informal communication around the project (Slack, GitHub issues, PR reviews, technical blog posts, etc)?
* How much time does the core team commit to the project?
* How big is the team? Who funds them? Why? How much? For how long?
* Who are the clear leaders? Are there any areas lacking clear leadership? Testing? Release? Documentation? These roles sometimes go unfilled.
* What is the rate of ongoing contributions to the project (typically in the form of merged commits).

#### Users

* Who uses the project?  Get a few in-depth references from 2-4 of them who actually know and understand it.
* What do real users consider to be its strengths and weaknesses?  Any concrete examples of these?
* Perception vs Reality:  Is there lots of buzz, but the software is flaky/untested/unused?  Does it have a bad reputation for some flaw that has already been addressed?

#### Contributor experience

* Besides the core team, how active is the surrounding community? Bug reports? Assistance to newcomers? Blog posts etc.
* Is it easy to contribute to the project as an external contributor? If not, what are the main obstacles?
* Are there any especially difficult personalities to deal with? How is this done? Is it a problem?
* Getting interviews with 2-3 external contributors is advisable for DD process, both from the community and technical perspective. It can help to identify technical depth in areas like extensibility, API design, and general code architecture.
* For more in-depth review of the contributor experience, consulting with [tag-contributor-strategy](https://github.com/cncf/tag-contributor-strategy) is always a good idea.

#### Context

* What is the origin and history of the project?
* Where does it fit in the market and technical ecosystem?
* Is it growing or shrinking in that space?  Is that space growing or shrinking?
* How necessary is it? What do people who do not use this project do?  Why exactly is that not adequate, and in what situations?
* Clearly compare and contrast with peers in this space.  A summary matrix often helps.
  Beware of comparisons that are too superficial to be useful, or might have been manipulated so as to favor some projects over others.
  Most balanced comparisons will include both strengths and weaknesses, require significant detailed research, and usually there is no hands-down winner.
  Be suspicious if there appears to be one.

#### Other advice

* Bring in other people (e.g. from your company) who might be more familiar with a
  particular area than you are, to assist where needed. Even if you know the area,
  additional perspectives from experts are usually valuable.
* Conduct as much of the investigation in public as is practical. For example, favor explicit comments on the
  submission PR over private emails, phone calls etc. By all means, conduct whatever communication might be
  necessary to do a thorough job, but always try to summarize these discussions in the PR so that others can follow along.
* Explicitly disclose any vested interest or potential conflict of interest that you, the project sponsor,
  the project champion, or any of the reviewers have in the project. If this creates any significant concerns regarding
  impartiality, it is usually best for those parties to excuse themselves from the submission and its evaluation.
* Fact check where necessary. If an answer you get to a question does not smell right, check the underlying data, or get a second/third opinion.
