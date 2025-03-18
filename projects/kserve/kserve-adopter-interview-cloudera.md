# KServe Adopter Interview – Cloudera

**Meeting Time:** July 1, 2025  
**Attendees:**  
- **TOC:** Kevin Wang, TOC Sponsor of KServe  
- **TOC:** Faseela K, TOC Sponsor of KServe  
- **Adopter:** Zoram Thanga
- **Adopter:** Cory Johannsen 

---

## About the Adopter Interview

The purpose of the CNCF TOC Adopter Interview is to evaluate the maturity, adoption patterns, and operational impact of the KServe project within production environments.  
The responses below reflect Cloudera’s technical experience and strategic assessment of KServe based on its integration into the Cloudera Data Platform (CDP) AI Inference services.

---

## Organization Introduction

### Can you provide an overview of your organization and its core activities?  

**A:**  
Cloudera is a global enterprise data cloud provider specializing in delivering a unified, hybrid data platform for multi-cloud and on-premises environments. Its flagship offering, the Cloudera Data Platform (CDP), enables secure, scalable, and compliant data lifecycle management, covering ingestion, storage, analytics, machine learning, and AI workloads. CDP is engineered to meet stringent enterprise requirements for governance, security, and operational resilience, enabling organizations to transform heterogeneous datasets into actionable business intelligence and AI-driven solutions.

---

## Motivation

### What factors led you to adopt KServe over other solutions in this domain?  

**A:**  
In 2022–2023, Cloudera conducted a technical evaluation of strategies to deliver a next-generation AI inference service. Options included developing a proprietary serving platform, extending the Cloudera Data Science Workbench, or adopting an open-source serving framework.  

After a structured analysis, KServe was selected due to:  
- **Open governance model** under LF AI & Data and CNCF  
- **License suitability** for enterprise distribution  
- **Active and approachable community** with transparent processes  
- **Neutrality in roadmap ownership**—no single vendor dictates priorities  
- **Alignment with cloud-native principles**, leveraging Kubernetes and Knative Serving rather than re-implementing core primitives

This ensured both long-term technical sustainability and alignment with Cloudera’s open-source-first strategy.

---

## Usage Scenario

### How long has your organization been using KServe?  
**A:**  
Adoption began in late 2022, with KServe fully integrated into our AI Inference product architecture by mid-2023. The service reached general availability in 2024 and has been in production for approximately one year.

### What are the primary use cases and features currently leveraged?  
**A:**  
KServe serves as the foundational serving layer for both generative AI (GenAI) and traditional AI models within CDP. Key capabilities in use include:  
- **Custom runtime creation** using KServe’s extensible cluster serving runtime framework  
- **Kubernetes-native deployment patterns**, inheriting operational benefits from Knative Serving  
- **Flexible scaling and traffic management** to optimize inference workloads  
These features enable rapid onboarding of diverse model types with minimal operational overhead.

### What is the current scale and operational maturity?  
**A:**  
The system is in production with approximately five enterprise customers across banking, government, healthcare, and oil & gas sectors. Customers currently operate a small number of large language models (LLMs) in production, integrated with Cloudera’s security and governance features. While usage volume is modest at present, the adoption trajectory is positive as client onboarding progresses.

### What version is deployed and what is your upgrade policy?  
**A:**  
Cloudera is currently shipping KServe v0.14, with planned upgrades to v0.15 and v0.16 within a two-to-three-month timeframe. Our target cadence is to adopt each quarterly community release, incorporating relevant security patches promptly. All updates undergo automated downstream integration testing across supported environments to validate compatibility and performance.

### How was the integration experience? Were there significant challenges?  
**A:**  
Integration into CDP was technically straightforward due to KServe’s alignment with core cloud-native components (Kubernetes, Istio, Knative). No KServe-specific blockers were encountered. Any integration complexity was limited to general Kubernetes-based infrastructure patterns rather than the KServe layer itself.

### How effective is the project documentation?  
**A:**  
KServe’s documentation is comprehensive, technically accurate, and closely synchronized with the codebase due to a mandatory documentation policy for contributions. The architecture, while non-trivial, is well-documented with clear conceptual and operational guidance. Cloudera actively contributes documentation improvements to address gaps, and our engineering teams have found the onboarding materials particularly effective for new contributors.

### What tangible value has KServe delivered?  
**A:**  
- **Accelerated time-to-market**—estimated reduction of at least two years in internal development effort  
- **Multi-cloud compatibility** with minimal integration friction  
- **Rapid adoption of GenAI advancements** via upstream community releases  
- **Operational modularity**, enabling selective feature enablement without excess overhead

### What are your future plans regarding KServe?  
**A:**  
Cloudera intends to expand its contributions, including:  
- Enhancements to inference logging for cost efficiency and scalability  
- Additional log storage backends (e.g., Azure Blob, further Google Cloud integrations)  
- Increased governance participation with the goal of attaining maintainer status  
- Upstreaming customer-driven features to avoid maintaining long-lived forks

---

## Perception

### How would you assess KServe’s community, governance, and growth potential?  

**A:**  
- **Openness:** The community actively encourages participation from new contributors, regardless of experience level.  
- **Governance:** Feature reviews are thorough and technically rigorous, ensuring long-term maintainability.  
- **Growth Potential:** High, driven by the expansion of AI workloads across industries.  
- **Maintainer Diversity:** Currently limited to a small group concentrated in Southeast Asia and the US East Coast, creating some time zone latency.  
- **Responsiveness:** Positive overall, though review timelines can be extended due to resource constraints.

**Challenges Noted:**  
- Limited core maintainer bandwidth affects PR review turnaround  
- Build/test infrastructure capacity constraints contribute to slow CI cycles  
- Expectation that CNCF incubation will enable infrastructure improvements

### How is Cloudera participating in the community?  
**A:**  
- Code and documentation contributions  
- Internal and customer evangelism  
- Technical reviews of community design proposals and contributions

### Have you engaged with maintainers directly?  
**A:**  
Yes, primarily via Slack and GitHub issues/PRs. Outcomes have been positive, though asynchronous communication across time zones can delay feedback cycles.

---

## Project Strengths

- Transparent governance and absence of single-vendor dominance  
- Technically robust architecture with strong documentation  
- Welcoming and collaborative community culture  
- Adoption by multiple reputable enterprises, contributing to project longevity  
- Rapid integration of emerging AI capabilities

---

## Project Improvements

### What factors currently limit the project’s potential?  
**A:**  
- Dependency on Knative for serverless capabilities, which faces its own resourcing challenges  
- Limited number of active maintainers, leading to bottlenecks in PR review and release throughput  
- Build pipeline inefficiencies (long execution times, occasional infrastructure-related failures)

### What specific improvements would you recommend?  
**A:**  
- Expand maintainer pool across additional time zones to improve review responsiveness  
- Increase community investment in CI/CD infrastructure to reduce build times and instability  
- Improve visibility and responsiveness in asynchronous channels such as Slack  
- Evaluate options to reduce operational dependency on Knative if upstream velocity declines

---
