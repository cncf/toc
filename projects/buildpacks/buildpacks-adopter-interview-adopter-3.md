# Buildpacks Adopter Interview - Adopter 3

**Meeting Time:** TBD  
**Attendees:**  
- **TOC:** Faseela K, TOC Sponsor of Buildpacks  
- **Adopter:** Adopter 3 representative

---

## About the Adopter Interview

The purpose of the CNCF TOC Adopter Interview is to evaluate the maturity, adoption patterns, and operational impact of the Buildpacks project within production environments.

Raw notes are recorded during the interview and refined into this formal version for adopter review, correction, and approval before being included in Buildpacks due diligence documentation.

---

## Organization Introduction

### Can you give us an overview of your organization and what it does?

Adopter 3 is a cybersecurity company focused on visibility, analytics, and automation for security operations.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

Adopter 3 was drawn to Buildpacks for several key reasons:

* The ability to control supported base images through builders
* A standardized build process that works across multiple programming languages
* Built-in support for generating software bills of materials (SBOMs)
* Rebasing capabilities that simplify vulnerability patching
* Reduction in the need for application teams to manage container packaging and image construction

Overall, Buildpacks aligned well with Adopter 3's goal of standardizing container builds while offloading complexity from development teams.

## Usage Scenario

### How long has your organization used the project?

Adopter 3 began experimenting with Buildpacks and has been using it for production workloads since 2022.

### What were the main motivations to adopt the project and which key features do you use today?

Buildpacks are used as the standard mechanism for building containers across multiple languages. Key features in use today include:

* A consistent container build process across languages
* Rebasing to address operating system–level vulnerabilities
* Leveraging the open-source ecosystem of builders and buildpacks to accelerate onboarding
* Community engagement through Slack and working group meetings
* Buildpack caching to improve build performance, with identified opportunities for further optimization

### What is the current level of usage (pre-production, production) and scale?

Buildpacks are currently used to build containers for production workloads across multiple teams and a shared CI/CD system. Adopter 3 initially used Paketo buildpacks, but encountered inefficiencies that led to the development of custom buildpacks to improve efficiency, simplicity, and consistency.

### What version is used and what is your update cadence with the project?

Adopter 3 uses the pack CLI to build containers. Paketo buildpacks are used extensively, with varying Buildpacks API versions depending on the buildpack.

Updates to pack are not fully automated and typically occur when addressing security vulnerabilities or when important new features are required. Production workloads are updated more regularly, while some older versions occasionally require upgrades due to deprecated or unsupported base images.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

I inherited the project after a former colleague had already started a proof-of-concept (POC). At the time, I had no prior knowledge of Buildpacks and only basic experience with container building. I prioritized understanding the technology, which involved reading the documentation extensively.

The POC demonstrated that we could build most of our applications with minimal or no configuration changes. We then focused on customizing the process for apps that required it. For instance, some of our internal applications use a web server not supported out of the box by the available buildpacks. We also had a custom build pattern where specific artifacts are produced by teams to streamline the creation of applications on instances or containers. After initially considering moving away from this custom pattern, we realized that creating our own custom buildpack was the perfect fit. This process led to my engagement with the community and a significantly deeper understanding of the technology.

This effort resulted in us creating our own composite/meta buildpacks, which then required building out pipelines and tooling to support the ecosystem's reliance on semantically versioned buildpacks and artifacts. Our need to use custom base images also led us to create our own builders.

One of the initial appeals of the project was the ease of quickly getting started with open-source builders and buildpacks. We chose Paketo buildpacks for their broad support for languages such as Java, Python, and Golang.

We encountered two main challenges. First, some internal applications required custom buildpacks due to web server requirements not covered by available open-source buildpacks, which ultimately proved beneficial by allowing us to optimize our builds. Second, we faced a lack of support for multiple architectures, specifically arm64. This led to small contributions and deeper project involvement to help advance the multi-arch effort.

Supply chain security was a major benefit of adoption. Rebasing significantly reduced the effort required to patch OS-level vulnerabilities, saving developers considerable time. While reproducible builds are not a primary focus, Adopter 3 emphasizes standardization and patching.

Adopter 3 also maintains internal integration tests and automation to track and consume the latest Paketo releases, with plans to further automate buildpack update workflows.

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

Adopter 3's adoption journey began with the Paketo Buildpacks project, which was particularly helpful in demonstrating how buildpacks adapt behavior based on application source code. Additional valuable resources included:

* buildpacks.io documentation
* The Buildpacks Platform and API specifications on GitHub
* The buildpacks samples repository and example applications

One area identified for improvement is documentation around [exec.d](https://buildpacks.io/docs/for-buildpack-authors/how-to/write-buildpacks/use-exec.d/), particularly with more concrete examples. Overall, the existing documentation has been sufficient to onboard multiple internal team members, though Adopter 3 has not yet contributed documentation improvements upstream.

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

Yes.

* Consistent base image usage and automated updates

Because we set up our own builders and buildpacks we introduced automation to keep base images up-to-date and pick up updates from upstream (paketo) buildpacks. Initially the updates were done manually in order to understand all of the moving pieces. After some time though we were able to fully automate the process so all new builds use the latest buildpacks, builders, and base images. Using buildpacks means new builds will always use approved base images that are kept up-to-date.

* Custom buildpack for custom build process

The fact that anyone can create a buildpack to meet their specific needs was a huge help to us. This meant that adoption of buildpacks did not require development teams to change their build process, and instead we were able to create a buildpack to make the transition seamless.

* Standardized pipeline and tooling that support semantic versioning

Builders and buildpacks all rely on semantic versioning, which led to the development of new tooling and a ci/cd pipeline with support for building and releasing artifacts with semantic versioned releases. This pipeline also builds containers using buildpacks and has become the standard pipeline for new projects. This has simplified the developer experience and reduced the number of pipelines that need to be maintained.

* Rebasing

Using buildpacks has reduced the burden of patching vulnerabilities in container images. Rebasing has allowed us to remove the burden of patching os-level vulnerabilities, so that developers can focus on their app code and dependencies.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

Adopter 3 has several team members contributing to Buildpacks, with deeper involvement in the Paketo project, including maintainership. The team is actively collaborating with Buildpacks maintainers on an RFC to enhance rebasing capabilities.

## Perception

### What is your perception in terms of the project’s:

* Community openness
  + Excellent. Adopter 3 has participated in mentorship initiatives focused on multi-architecture support.
* Governance
  + Excellent - Excellent. The project is perceived as welcoming and focused on solving real problems rather than vendor influence.
* Community growth potential
  + Strong, particularly with ongoing work to enhance rebasing.
* Maintainer diversity and ladder
  + Positive, though Adopter 3 has not explored formal maintainer roles within the Buildpacks project itself.
* Maintainer response
  + Excellent, very open for collaboration

### How are you participating in the project community?

Adopter 3 joins working group meetings regularly and engages with the community on Slack, including participation in mentorship programs with Buildpacks maintainers focused on multi-arch support. The team follows RFCs and helps out community members as needed.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

Adopter 3 participates through Slack, working group meetings, RFC discussions, and mentorship programs. Engagement has consistently resulted in positive outcomes, particularly during early adoption when guidance and advice were needed.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

Strong standardization across languages

Consistent, repeatable build processes

Rebasing and supply chain security features

SBOM and provenance support

Active, welcoming community with accessible communication channels

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

Multi-architecture support was initially a challenge for Adopter 3, particularly in the early stages of adoption. However, significant progress has been made through collaboration across the Buildpacks ecosystem. While continued work is still needed, the overall momentum is positive and moving in the right direction.

Enhancing the capabilities of rebasing remains another important area of focus. Rebasing already delivers meaningful value, particularly for vulnerability patching, but there is an opportunity to expand its impact further. Adopter 3 is actively contributing to this effort through the RFC process.

Buildpacks are also not always a good fit for every use case. In particular, teams that rely heavily on complex, multi-stage Dockerfile-based builds may find Buildpacks less suitable. At Adopter 3, Dockerfiles are supported where Buildpacks do not make sense, although Buildpacks remain the default and recommended standard.

Finally, upgrades can be challenging due to the need to coordinate across multiple components, including pack, builders, buildpacks, base images, and Docker versions. While upgrading the pack CLI itself is generally straightforward, Docker compatibility issues have occasionally caused friction. Recent releases have improved upgrade guidance, particularly around API version support, which has helped mitigate some of these challenges.

### In your opinion, what can the project do better?

Overall, the Buildpacks project is doing the right things by actively supporting its community through working group meetings and Slack, and by listening to users to drive improvements and new development efforts.

That said, continued focus on improving the upgrade experience would further benefit users. While documentation around API version releases and supported functionality is available, older versions of pack may still reference images or configurations that are no longer supported, which can force unplanned upgrades. Clearer guidance and tooling around version alignment across the ecosystem would help reduce this friction.

Adopter 3's approach reflects this reality: teams are allowed to use Dockerfiles when Buildpacks are not a good fit, but Buildpacks remain the standard path. Continued refinement in documentation, upgrade workflows, and compatibility guidance would make this standard even easier to adopt and maintain.
