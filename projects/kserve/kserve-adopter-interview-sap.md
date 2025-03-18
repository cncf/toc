# KServe Adopter Interview - SAP

**Meeting Time:** 10th July 2025  

**Attendees:**  
- **TOC:** Kevin Wang, TOC Sponsor of KServe  
- **TOC:** Faseela K, TOC Sponsor of KServe  
- **Adopter:** Lize Cai, Architect @ SAP AI Core (Machine Learning Platform for SAP)  

---

## About the Adopter Interview

The intent of this interview is to assess the maturity, adoption, and real-world usage of the KServe project within production environments, from the perspective of an active adopter. These insights contribute to the CNCF TOC’s due diligence process and KServe’s maturity evaluation.  

Raw notes are recorded during the interview and refined into this formal version for adopter review, correction, and approval before being included in KServe’s due diligence documentation.

---

## Organization Introduction

### Can you give us an overview of your organization and what it does?

SAP is a multinational software corporation specializing in enterprise software solutions that enable organizations to manage business operations and customer relationships effectively. Within SAP’s Business Technology Platform (BTP), **SAP AI Core** serves as a dedicated service for managing the execution and operational lifecycle of AI assets in a standardized, scalable, and hyperscaler-agnostic manner. The platform supports both **internal SAP teams** and **external enterprise customers**, enabling them to develop, deploy, and operate AI workloads with robust governance, security, and scalability.

---

## Motivation

### Compared with other products in this space (proprietary and open source), what drew you to the project?

The decision to adopt KServe within SAP AI Core was made prior to my joining the team. During the evaluation phase, alternative solutions such as BentoML and Seldon Core were also carefully assessed for their production-grade model serving capabilities. Each of these offerings had strengths, particularly in simplifying deployment and providing Docker-focused workflows.

KServe, however, stood out for its Kubernetes-native architecture, which aligned closely with SAP AI Core’s platform requirements. Its cloud-native design, strong alignment with CNCF principles, and support from multiple established organizations gave additional confidence in its long-term sustainability. Originating from the Kubeflow project also added to its credibility and maturity. These factors made KServe a strong match for SAP AI Core’s needs around multi-tenancy, scalability, and standardized AI serving.

---

## Usage Scenario

### How long has your organization used the project?

SAP AI Core began adopting KServe in **July 2020**, with the platform entering **full production use in 2021**.

### What were the main motivations to adopt the project and which key features do you use today?

Our primary motivation was KServe’s **seamless integration with the Kubernetes ecosystem**, enabling a consistent operational model for AI workloads. Its native scalability and abstraction of complex serving infrastructure aligned perfectly with our multi-cloud deployment strategy.

Today, we primarily utilize **KServe’s InferenceService API** to deploy, manage, and scale AI models in production. This API abstracts infrastructure details, allowing data scientists and developers to focus on model logic while the platform manages lifecycle, autoscaling, and networking.

### What is the current level of usage (pre-production, production) and scale?

We operate KServe in **full production** across multiple environments. SAP AI Core serves **thousands of tenants**—both internal and external—leveraging KServe for diverse AI workloads. Due to confidentiality constraints, exact numbers are not disclosed.

### What version is used and what is your update cadence with the project?

We currently operate **KServe v0.11** in production. We do not strictly follow the upstream release cadence; instead, upgrades are triggered by:
- The need to address **security vulnerabilities**,
- The requirement for **specific new features**,
- Or when a given version reaches **end-of-support**.

### Can you walk us through your adoption and integration experience? What challenges did you face?

Overall, the adoption process was smooth and aligned well with our Kubernetes-centric infrastructure. However, KServe’s **external dependencies**—notably Istio and Knative—introduce complexity during upgrades. Coordinating version compatibility and upgrade sequencing for these components can be operationally challenging, particularly in regulated enterprise environments.

### Did you find the information in the repo or the project docs valuable to your implementation?

KServe’s documentation is generally comprehensive, but content is **distributed between the GitHub repository and the official website**, and these sources are occasionally out of sync. The website’s "Getting Started" guides are excellent for onboarding, but deeper technical details often require searching through the repository. We recommend consolidating documentation into a **single source of truth** on the project website to reduce ambiguity.

### Has your implementation of the project provided measurable value?

Yes. SAP AI Core is inherently **multi-cloud**, and KServe’s Kubernetes-native, multi-cloud-ready design has significantly accelerated integration across hyperscalers.  
Key benefits include:
- **Storage Initializers** enabling cross-cloud object storage access,
- **Serverless scaling** capabilities (scale-to-zero) for cost efficiency in narrow AI use cases,
- Reduction in infrastructure management overhead, enabling faster delivery cycles.

### Do you have any future plans regarding the project?

Yes. Our team actively contributes to KServe, with **3–4 engineers** regularly submitting code, documentation, and feature requests. Recent contributions include:
- [Feature Request #4574](https://github.com/kserve/kserve/issues/4574)
- [Feature Request #4575](https://github.com/kserve/kserve/issues/4575)

We are also exploring ways to officially list SAP as an adopter on the KServe website and are closely monitoring **KServe + GenAI** initiatives for future adoption.

---

## Perception

### What is your perception of the project’s community, governance, and diversity?

The KServe community is **active, responsive, and diverse**. Questions raised in Slack channels typically receive prompt and helpful responses, surpassing the engagement levels of many other open source communities.

Governance is well-structured, with regular project health checks, transparent decision-making, and continuous onboarding of new contributors and committers. The **company diversity** within the maintainer group mitigates the risk of project capture by a single vendor.

However, meeting schedules are currently optimized for US time zones, making it difficult for contributors in Asia to participate consistently. Although there is an Asian-friendly meeting slot, attendance is limited due to fewer core contributors joining.

### How do you participate in the project community?

We engage via Slack, GitHub issues, and discussions. While timezone constraints limit live meeting attendance, we participate in the Asian-friendly maintainer meetings when possible.

### Have you engaged with maintainers directly, and what was the outcome?

Yes. We have engaged the community and maintainers for:
- Feature requests,
- Bug fixes,
- Documentation improvements,
- Upgrading outdated dependencies (often driven by internal compliance scans).

Outcomes have been positive, with community members responding constructively and collaborating effectively on required changes.

---

## Project Strengths

### In your opinion, what are the overall strengths of the project?

- Long-term **stability** (active for over 6 years) and sustained **community engagement**,
- Backing from multiple large organizations with strong AI focus,
- CNCF governance providing licensing clarity and long-term sustainability,
- Balanced corporate influence, avoiding single-vendor dominance,
- Open, responsive communication channels with maintainers.

---
