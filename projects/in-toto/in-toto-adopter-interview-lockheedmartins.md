# In-toto Adopter Interview - Lockheed Martin

Interviewee: Ian Dunbar-hall, Head of Open Source Program Office, Lockheed Martin
Interview date: Sept 3, 2024

## Organization Intro

### Can you give us an overview of your organization and what it does?

[Lockheed Martin](https://www.lockheedmartin.com/en-us/contact.html) is a leading aerospace and defense company.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

I recall it was at KubeCon either 2020 or 2021, and I went to the contribfest for the in-toto project. The project solves a foundational need in our space, securing software supply chain, which is very critical to our customers for delivering high quality products. Without it, it has resulted in very bad outcomes, Solarwinds, or Crowdstrike incidents, for example.

## Usage Scenario

### How does your organization use the project and how long have you used it?

For the in-toto specification, we don’t directly work with specification but consume it as part of libraries.

For the in-toto subprojects (application or libraries), we started to use the libraries in our OSS projects in Jan 2023. We have also incorporated in-toto attestation for corporate networks for any OSS projects that come to internal & external use. Basically any time when we consume any OSS projects, we check on the following:
- Is the OSS project approved for use in our company?
- How do we know if someone maliciously modified it in the corporate network?
- Can we still adopting it for products we are delivering to our customers?

### What version is used and what is your update cadence with the project?

We update the libraries fairly regularly, whenever any core library gets updated.
Specification changes are also adopted as part of the library update. We don’t implement the spec ourselves.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Overall, a really positive experience! 
- Well organized oss projects, very strong community behind it.
- Very large enterprises and universities are involved. Easy to get support.
- Well beyond community compared with other graduated projects.
- We also contributed some PR and got good feedback/reviews.  Testify donated a subproject under in-toto (command line client to create attestation) and we use it to create attestation for everything.
- Lots of end users and vendors.

### Did you find the information in the repo valuable to your implementation? What specifically?

Yes, without a ton of background, we were able to quickly incorporate the in-toto python libraries.  Slack/community meetings are both helpful. Easy to get technical recommendations. Docs are quite good for such a complex problem.

### Has your implementation of the project provided measurable value? 

Yes. The in-toto capabilities in our products were demo-ed frequently. We are writing a white paper around how to use in-toto across public sectors for supply chain security via the public sector of the CNCF end user groups.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

- Need to drive more adoptions within the company.
- We are working on the SBOMit project within OpenSSF, which is built on top of in-toto:
  * https://github.com/sbomit/specification
  * [Bomctl](https://github.com/bomctl/bomctl) will be announced as a sandbox project as part of OpenSSF this thurs which also uses in-toto.

## Perception

### What is your perception in terms of the project’s:

#### Community openness

Very active community, weekly meeting for some part of the in-toto ecosystem. Variety of the people involved speak very well for the project.

#### Governance

Steering committee was formed and a lot of people involved know the spec well, a good technical community.

#### Community growth potential

Good growth and expansion.

#### Maintainer diversity and ladder

Totally see this growing.

#### Maintainer response

Nothing but quick response and feedback.

### How are you participating in the project community?

We have done in-toto related talks at CloudNativeSecurityCon, OpenSource Summit etc. We are pretty active in the community, being on the community calls.  We are also maintaining a few other projects which are built on top of in-toto. 

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

A variety of reasons. My first one was a first time in-toto user and then I wanted to incorporate in-toto in our OSS tool, and a few other minor contributions.  All were addressed timely with good feedback.  Really good experience which led me to be more deeply involved.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

Beyond the community, and diversity of others using it, the spec is also very flexible. You can add to it, or expand it or create a unique solution with it. 

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

Growing the list of attestation types of in-toto will strengthen the project more.  For example, adding an OSS program approval.

### In your opinion, what can the project do better?

- Has done a really good job of getting people using in-toto and sometimes people don’t really realize they are using it.
 * [SLSA](https://slsa.dev/) is built on top of in-toto, which is very well known.  Yet people don’t realize it.

- Needs more marketing or branding because people don’t realize it as much as they are using it.
