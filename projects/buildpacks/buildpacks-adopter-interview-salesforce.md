# Buildpacks Adopter Interview - Salesforce

**Meeting Time:** Jan 28, 2026 5PM CET  
**Attendees:**  
- **TOC:** Faseela K, TOC Sponsor of Buildpacks  
- **Adopter:** Andrey Falko, Software Engineering Architect, Salesforce (internal CI systems)

---

## About the Adopter Interview

The purpose of the CNCF TOC Adopter Interview is to evaluate the maturity, adoption patterns, and operational impact of the Buildpacks project within production environments.

Raw notes are recorded during the interview and refined into this formal version for adopter review, correction, and approval before being included in Buildpacks due diligence documentation.

---

## Organization Introduction

### Can you give us an overview of your organization and what it does?

Salesforce is a large-scale enterprise software company with a significant internal development organization. I work within the CI systems team, which primarily supports internal code repositories across Salesforce. These CI systems serve more than 10,000 developers and handle the majority of our continuous integration and deployment workflows. Our CI systems allow us to safely move code from development into production, running automated tests, packaging code, signing artifacts, performing security scans, and orchestrating deployments. While we also support a small subset of Salesforce open source projects, the predominant focus is on internal production services.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

Before adopting Buildpacks, Salesforce relied largely on Dockerfiles for building images and running tests. Our team saw Buildpacks as a way to simplify and standardize the build process, particularly for our production configurations and integration tests. The “easy button” philosophy of Buildpacks, zero-configuration builds, strong defaults, and reproducibility, was particularly attractive. Additionally, Salesforce’s historical relationship with Heroku, which co-founded Buildpacks, made collaboration on adoption smoother. The project’s documentation and strong community support further reinforced our confidence in adopting it.

## Usage Scenario

### How long has your organization used the project?

Salesforce began experimenting with Buildpacks in 2022 and entered full production use shortly thereafter. The experimental phase lasted approximately four months, after which Buildpacks were rolled out across all major CI workflows.

### What were the main motivations to adopt the project and which key features do you use today?

The main motivations were to reduce manual configuration, achieve reproducible builds, and streamline integration testing and deployment workflows. Buildpacks allow us to package production configurations, integration tests, and, in some cases, service code into consistent, reusable units. Integration test Buildpacks are particularly valuable, as they support automated rebasing for dependency updates and security patches, which frees service owners from maintaining these updates manually.

### What is the current level of usage (pre-production, production) and scale?

Buildpacks are fully in production. On average, we run over 10,000 builds per day, with peaks reaching up to 50,000 builds depending on activities such as security patching. The usage is accelerating, partly due to the adoption of AI-assisted pull request generation.

### What version is used and what is your update cadence with the project?

We aim to update Buildpacks at least monthly, though in practice, updates occur roughly twice a year, mainly due to regression testing.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Adopting Buildpacks was relatively smooth, particularly because Salesforce has internal expertise from Heroku contributors. The initial challenges were mostly around deep infrastructure updates, such as upgrading an OS version that required changes across nested Buildpacks. These updates were labor-intensive but manageable with proper coordination. Since 2022, the experience has stabilized, and adoption for integration tests and configurations has become straightforward. Experimental features, such as optional flags for advanced functionality, have also been successfully implemented by several hundred services without major issues.

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

The project documentation has been critical to adoption. We rely heavily on the end-to-end guides for reproducing Buildpacks builds locally, which allow new users to onboard without issues. The clarity and completeness of the documentation significantly reduce the need for internal handholding.

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

The primary benefits have been ease of use and reduced manual effort. Buildpacks facilitate zero-configuration builds, standardized workflows, and reproducible builds for both development and production. While cost savings compared with Dockerfiles are marginal, the consistency, predictability, and speed improvements have been substantial, particularly for integration tests.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

We plan to extend the zero-configuration model to service builds, increasing usability across more use cases. Salesforce continues to contribute to Buildpacks upstream, particularly around multi-architecture support and internal stack integrations, while leveraging Heroku’s expertise for ongoing development.

## Perception

### What is your perception in terms of the project’s:

* Community openness : The community is open and responsive, with maintainers providing timely support for contributions and experimental features.
* Governance : Governance is structured, with clear channels for contributing and reporting issues. Salesforce works closely with Heroku maintainers to ensure alignment.
* Community growth potential : There is strong potential for growth, as other large companies such as Google and Bloomberg have also adopted Buildpacks. Increasing participation from additional vendors would further strengthen the ecosystem.
* Maintainer diversity and ladder : While Salesforce and Heroku contributors are predominant, maintainers are responsive and collaborative, which has facilitated adoption and contribution.
* Maintainer response:

### How are you participating in the project community?

Our team interacts with the upstream community mainly through Slack and internal coordination with Heroku maintainers. For most contributions, direct public engagement is minimal, as internal expertise suffices.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

Salesforce’s internal team generally did not need extensive engagement with the public Buildpacks community, because we have a strong internal presence of Heroku contributors who are deeply familiar with the Buildpacks stack. When our team did need to interact with upstream maintainers, it was primarily for specific contributions such as enabling multi-architecture support (e.g., ARM64/Graviton) or integrating internal sources for network-isolated builds. These engagements were conducted via Slack within the CNCF community and through direct collaboration with Heroku maintainers. The outcomes of these interactions were consistently positive and timely, allowing our contributions and customizations to be incorporated smoothly. For most day-to-day use cases, Salesforce relies on internal expertise rather than public community engagement.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

The key strengths include excellent documentation, reproducible builds, ease of use, and responsiveness of maintainers. Buildpacks support zero-configuration workflows, improve CI efficiency, and integrate well with Salesforce’s internal systems. Cross-pollination through the CNCF enhances visibility and strengthens the ecosystem.

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

One factor that could hold Buildpacks back from reaching its full potential is the difficulty for teams to migrate legacy applications from Dockerfiles to Buildpacks. While new projects adopt Buildpacks readily, older applications often remain on Dockerfiles due to the lack of a straightforward migration path. Facilitating this transition would likely increase adoption and overall impact.

### In your opinion, what can the project do better?

Potential improvements include smoother migration paths from legacy Dockerfiles to Buildpacks and expanded zero-configuration support for service builds. Increasing participation from additional companies would enhance vendor neutrality and community diversity.
