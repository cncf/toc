# In-toto Adopter Interview - GitHub

Interviewee: Zach Steindler, Principal Eng at GitHub
Interview date: Oct 7, 2024

## Organization Intro

### Can you give us an overview of your organization and what it does?

GitHub is a website where people work on code together.  Very popular in OSS for people to share code and build artifacts. Also used widely by enterprise.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

There are primiarly 2 things that drew us to the project:

-  We started using in-toto when we added build provenance. In-toto collects source projects to write specifications. 
- In-toto use cases were attractive to us.  There aren’t really other projects out there as an alternative that has lots of other projects using it.

## Usage Scenario

### How does your organization use the project and how long have you used it?

GitHub owns npm, we released npm provenance in 2022 which uses in-toto.  We use the in-toto framework to create publish attestation. Last year we released github artifact attestation, so anything you build with github can have build provenance.  We also use SBOM and use in-toto to represent it.

### What version is used and what is your update cadence with the project?

We maintain our own version of custom predicate.  We are currently up-to-date and we update as needed.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

It has been pretty smooth. There are docs around how to produce custom predicates. There are docs on how to produce build provenance. Libraries are relatively straightforward to use. Can’t think of any challenges we had!

### Did you find the information in the repo valuable to your implementation? What specifically?

Yes! Pretty good docs for in-toto attestation repo, SLSA(Supply Chain level for software artifacts) repo, very good repo.

### Has your implementation of the project provided measurable value? 

Tens of thousands of people make use of npm provenance and github artifact attestation.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

Yes!  For GitHub releases, we plan to make it immutable by leveraging in-toto attestation. Besides that, nothing concrete. We always keeping track of new attestation released from in-toto.

## Perception

### What is your perception in terms of the project’s:

#### Community openness

Very open, I participated in the slack channel in CNCF, and have created issues/PRs that have been resolved.

#### Governance

Don’t think I attended any meeting. Some of the PRs have been reviewed by the Governance/steering committee, they were prompt and thorough in review.

#### Community growth potential

Could be biased, we are definitely invested in the ecosystem and believe in the growth of it.

#### Maintainer diversity and ladder

Multiple Xs of diversity. There is some diversity in terms of gender and people’s background (industry & academic & non profit OSS foundation).

#### Maintainer response

Couple of PRs made by me were handled well.  Things are resolved in a reasonable amount of time.

### How are you participating in the project community?

Yes but not recently.  About 6 months ago, I attended some community meetings and submitted PRs.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

So far, I have good experience with PRs.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

Community discussions are great and how they bring them (industry & academic & non profit OSS foundation). Really thinking ahead and anticipating needs before people need them.  Continue to be an active community.

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

Not really.  Struggle to come up with an answer.  Only worry is if there are lots of layoffs, would people have time to contribute in-toto?

### In your opinion, what can the project do better?

Continue to think about where the industry is headed and anticipate the needs. They have demonstrated the ability to do so so far.
