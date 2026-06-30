# Buildpacks Adopter Interview - Google

**Meeting Time:** 1/27/2026 7:30 PM CET  
**Attendees:**  
- **TOC:** Faseela K, TOC Sponsor of Buildpacks  
- **Adopter:** James Ma, Product Manager, Serverless Runtimes, Google Cloud Platform  
- **Adopter:** Matthew Robertson, Tech Lead, Serverless Runtimes, Google Cloud Platform

---

## About the Adopter Interview

The purpose of the CNCF TOC Adopter Interview is to evaluate the maturity, adoption patterns, and operational impact of the Buildpacks project within production environments.

Raw notes are recorded during the interview and refined into this formal version for adopter review, correction, and approval before being included in Buildpacks due diligence documentation.

---

## Organization Introduction

### Can you give us an overview of your organization and what it does?

Google Cloud Platform (GCP) is a public cloud offering a suite of cloud computing services that allow businesses and developers to build, test, and deploy applications on Google’s global infrastructure. The Serverless Runtimes team focuses on products such as Google App Engine Standard, Google App Engine Flexible, Cloud Run, and Cloud Run Functions, all of which provide source-based deployment experiences for customers

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

Prior to 2020, Google used a proprietary source-to-container solution, particularly within App Engine. As part of a broader effort to modernize runtimes and align more closely with open standards, consistent with Google’s involvement in Kubernetes and other CNCF initiatives, the team decided to invest in an open specification rather than continue evolving a closed solution.

Buildpacks were a strong fit because they support one of the core value propositions of Google’s serverless platforms: source deployment, or the ability to securely take arbitrary source code and run it in the cloud. The project offered a mature, open standard with a strong track record of operational excellence from early contributors such as Heroku and Pivotal. Additionally, the specification was flexible enough to allow Google to implement its own buildpacks, base images, and stacks tailored to GCP’s platforms, while relying on the broader community to evolve the specification over time. Buildpacks also aligned well with adjacent CNCF standards around software supply chain concerns.

## Usage Scenario

### How long has your organization used the project?

Google has used Buildpacks in production since October 2020, beginning with Cloud Run and Cloud Functions, and later expanding to App Engine.

Reference: https://cloud.google.com/blog/products/containers-kubernetes/google-cloud-now-supports-buildpacks

### What were the main motivations to adopt the project and which key features do you use today?

At the time of adoption, App Engine already offered a well-loved source deployment experience, but the process was relatively opaque. Google wanted to provide a similar experience across Cloud Run while continuing its commitment to open technologies and collaboration on open standards.

Buildpacks met several key needs:

* An open specification that produces OCI-compliant images
* A build and creation model that aligned well with GCP’s platforms
* Flexibility to implement platform-specific buildpacks, base images, and stacks
* Rebase functionality, which is heavily used to automatically update and secure running images without requiring redeployment

The ability to inspect application source code and determine how to containerize it dynamically is viewed as a core strength compared to Dockerfile-based approaches.

### What is the current level of usage (pre-production, production) and scale?

Buildpacks are used in production across App Engine (Standard and Flexible), Cloud Run, and Cloud Run Functions. Google is fully standardized on Buildpacks for source-based deployment; the primary alternative offered to customers is to supply their own Dockerfile.

At scale, Google processes tens of thousands of image builds per day using its Buildpacks-based systems, with usage growing alongside the popularity of products like Cloud Run.

### What version is used and what is your update cadence with the project?

Google closely tracks the Buildpacks specification but does not always remain on the very latest versions due to the complexity of managing multiple independently versioned components (such as the platform API, buildpack API, lifecycle, and tooling). Updates are approached cautiously to avoid breaking changes for customers, and upgrades may take significant time when compatibility concerns arise. Google also performs regular updates and maintenance in its own implementation at<https://github.com/GoogleCloudPlatform/buildpacks>.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Overall, the experience of adopting and integrating Buildpacks has been smooth. Documentation and community responsiveness were cited as helpful during adoption.

However, operating Buildpacks at Google’s scale has surfaced challenges, particularly around build performance, caching, and resource usage. As customer expectations shift toward faster feedback loops and high-iteration workflows, including AI-generated code, reducing build times while maintaining security and consistency remains an ongoing effort.

Buildpacks are also less suited to highly iterative local development scenarios, where Dockerfiles benefit from more mature and standardized local tooling. Additionally, managing upgrades across multiple Buildpacks-related components has proven complex, with some upgrades taking multiple years due to breaking changes and compatibility concerns.

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

* The documentation is very useful in giving an overview of the project and a snapshot of current features/capabilities
* RFCs are well maintained, and give a perspective of what future looking changes are being proposed to the spec
* The weekly meetings are welcoming and a good source of information
* Slack channel is responsive to ad hoc information and queries

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

The Buildpacks specification and Google’s Buildpacks implementation are mission-critical to GCP’s serverless services. Key measurable benefits include:

* Standardizing source deployment across multiple GCP products
* Significantly reducing operational toil and engineering effort
* Enabling frequent and automated security updates through rebase
* Allowing teams to focus on differentiated customer features rather than maintaining bespoke build systems

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

There are no specific future plans to share at this time. Google expects to continue relying on Buildpacks as the standard mechanism for source deployment across its serverless platforms.

## Perception

### What is your perception in terms of the project’s:

* Community openness
  + Very positive. RFCs are up to date, and maintainers are active across Slack, GitHub, and weekly meetings.
* Governance
  + Fair and effective. Project leaders are open to ideas while maintaining forward momentum.
* Community growth potential
  + N/A
* Maintainer diversity and ladder
  + Google noted that broader organizational representation in long-term maintainership is beneficial for project sustainability. They expressed increased confidence when additional companies such as Bloomberg were represented at the TOC level. This was not raised as a concern about governance or as a blocker to graduation.
* Maintainer response
  + Strong response times and constructive interactions.

### How are you participating in the project community?

Google engages with the community primarily through RFC feedback and GitHub issues, particularly when changes may impact its own implementation of the specification. Contributions are made on an as-needed basis, though the team is not able to commit to sustained upstream maintainership due to internal priorities.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

There are times when we need to reach out to the community of maintainers when there are CVEs, or when we’re trying to collaborate on shared problems as platform maintainers. We usually go through SLack or the weekly meeting. Tactical discussions sometimes occur on the RFCs or GitHub issues

## Project Strengths

### In your opinion, what are the overall strengths of the project?

The project benefits from a strong maintainer team with representation from organizations such as Heroku, VMware, and Bloomberg. Buildpacks are closely tied to real-world, large-scale platform usage, and the specification has proven effective at standardizing source-to-container workflows across diverse environments.

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

N/A

### In your opinion, what can the project do better?

Google noted that while Buildpacks has been stable in production, upgrading across multiple versioned components can be complex at large scale. As a result, upgrades are approached cautiously to avoid breaking customer workflows. Continued stabilization of APIs and clearer upgrade paths would further improve the experience, particularly for large platform providers.
