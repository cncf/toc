# OpenFGA Adopter Interview - Docker  

**Meeting Time:** 9/11/2025  
**Meeting Record:** TBD  
**Attendees:**  
- TOC: Faseela K, TOC Sponsor of KServe  
- TOC Shadow: Ricardo Aravena  
- Adopter: Gurleen Sethi, Senior Software Engineer at Docker, Inc.

## Organization Intro  

### Can you give us an overview of your organization and what it does?  
Docker provides tools that help developers build, share, run, and verify applications across environments. By abstracting away environment configuration and infrastructure management, Docker accelerates software delivery and improves developer productivity.  

## Motivation  

### Compared with other products in this space (proprietary and open), what drew you to the project?  
Docker evaluated multiple access control systems, including Authzed and Ory Keto, as well as non–ReBAC systems. ReBAC was identified as more flexible and better suited for Docker’s needs.  

Key drivers for adopting OpenFGA included:  
- Self-hosted and open source, with straightforward deployment (e.g., Docker Compose, runnable in under five minutes).  
- Backed by CNCF and supported by companies with strong security backgrounds (e.g., Okta/Auth0).  
- Mature SDKs, APIs, and testing tools.  
- Active community and responsive maintainers, with prompt resolution of feature requests and issues.  

## Usage Scenario  

### How long has your organization used the project?  
OpenFGA was adopted in February 2024 and has been running in production since March 2024.  

### What were the main motivations to adopt the project and which key features do you use today?  
- Relationship-Based Access Control (ReBAC).  
- Domain-Specific Language (DSL) for modeling.  
- Rich, consistent CLIs and SDKs.  

### What is the current level of usage (pre-production, production) and scale?  
- Production usage.  
- 100–150 RPS.  
- [Internal note, not for publication]: Running on 10 containers (1 CPU, 1Gi memory each), although 4–5 would be sufficient.  

### What version is used and what is your update cadence with the project?  
- Always deploy the latest release.  
- New releases are applied immediately to staging, tested for two weeks, then promoted to production.  

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?  
The adoption process was straightforward. OpenFGA was easy to run locally, the documentation was clear, and the integration with existing systems was simple. To ensure a safe migration, Docker initially ran OpenFGA in parallel with its internal authorization system. Every permission check was sent to both systems, and results were compared to verify consistency. Once both systems returned identical outputs, Docker began incrementally routing production traffic exclusively through OpenFGA.  

The primary challenge during adoption was related to scaling issues, particularly with batch access checks across multiple records. However, these were quickly resolved by the upstream team. Another area identified for improvement is the DSL parsing, which, although already superior to comparable systems, could still be enhanced. Docker also noted the absence of a paid support model, with community Slack channels serving as the only support mechanism.  

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?  
Yes, the documentation was extremely valuable. The modeling guides in particular—both the Getting Started guide and the Advanced modeling guide—were instrumental in helping Docker understand the concepts behind ReBAC and in applying them to real-world scenarios.  

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.  
The adoption of OpenFGA has delivered measurable improvements. Docker’s previous system was rigid and difficult to extend, often requiring significant manual work to add new permissions. With OpenFGA, modifications are centralized in the modeling file, dramatically reducing manual intervention. This has also accelerated the integration of new Docker products into the access control system, enabling faster delivery and reduced operational overhead.  

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.  
Docker intends to expand its internal adoption of OpenFGA across more services. The team actively raises feature requests and issues in the community and plans to continue doing so. Currently, one team member is engaged with the upstream project, but as adoption grows, broader participation is anticipated.  

## Perception  

### What is your perception in terms of the project’s:  
- **Community openness:** The community is open and welcoming, and adopters find it easy to engage and get help through channels like Slack and GitHub.  
- **Governance:** While Docker is not deeply familiar with the formal governance model, the way the project is run suggests a healthy level of structure and collaboration.  
- **Community growth potential:** There is strong potential for growth, with more organizations adopting OpenFGA and an active pace of contributions.  
- **Maintainer diversity and ladder:** Most of the core maintainers today are from Okta, but contributions are starting to come in from other companies as well. Community members from different backgrounds are active on Slack and GitHub, which suggests healthy potential for broader maintainer diversity over time.  
- **Maintainer response:** Maintainers are highly responsive, often addressing feature requests, performance issues, and bugs quickly and constructively.  

### How are you participating in the project community?  
Docker participates primarily through the community Slack channel, GitHub issues, and monthly community meetings.  

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?  
Yes. Docker has engaged with the community and maintainers on new feature requests, performance issues, and bug reports. Most interactions have taken place through Slack and GitHub, and in each case, the outcomes were positive and timely.  

## Project Strengths  

### In your opinion, what are the overall strengths of the project?  
Docker sees OpenFGA’s main strengths as its highly active and responsive community, strong and diverse maintainers, and the robustness of its design and performance. The project is also supported by excellent documentation and modeling guides, which lower the barrier to adoption and make the concepts of ReBAC accessible to new users.  

## Project Improvements  

### Is there something you feel that holds the project back from reaching its ultimate potential?  
One area that could be improved is the DSL and the tooling that surrounds it. While OpenFGA’s DSL is already better than comparable systems, additional enhancements would make it even more powerful. Another limitation is visibility: greater marketing and awareness-building could help OpenFGA reach a broader audience, as it provides capabilities that are highly relevant to many organizations.  

### In your opinion, what can the project do better?  
From Docker’s perspective, performance optimizations in certain complex scenarios would provide significant benefits. For example, batch checks for the same user and object across multiple relations currently require each check to be evaluated separately. Introducing a pre-evaluation step could minimize database lookups and improve efficiency. Similarly, in some cases where the `and` condition is used in the DSL, performance could be further optimized.  
