---
adopter: CyberAgent, Inc.
organization: CyberAgent, Inc.
project: KServe
date: 2025-07-16
status: Approved
toc_sponsors:
  - Kevin Wang
  - Faseela K
attendees:
  - Yuki Iwai (CyberAgent, Inc.)
  - Kevin Wang (TOC Sponsor, KServe)
  - Faseela K (TOC Sponsor, KServe)
---

# KServe Adopter Interview – CyberAgent, Inc.

## Organization Introduction

### Can you give us an overview of your organization and what it does?

CyberAgent, Inc. is a leading internet and entertainment enterprise in Japan, operating across three primary domains: **Media & IP** (including broadcasting, streaming services, and blog hosting), **Advertising**, and **Gaming** (mobile and consumer titles). Across these sectors, the organization strategically employs advanced AI technologies to optimize content delivery, enhance advertising targeting, and improve end-user experiences.

A cornerstone of its AI infrastructure is **Cycloud**, a proprietary private cloud platform engineered to deliver robust computing resources to internal engineers, creators, and business stakeholders. Cycloud includes a purpose-built **Machine Learning (ML) platform**, comprising a private bare-metal GPU cluster hosted in CyberAgent’s own Japanese data center.

The GPU cluster consists of over **80 compute nodes** and approximately **500 GPUs**, supporting between **1,500 and 2,500 pods** concurrently. The platform offers **GPU-as-a-Service (GPUaaS)**, single-node and distributed model training, and a globally accessible endpoint for inference and model serving. **KServe** is deployed as the primary framework for model inference and serving operations within this environment.

---

## Motivation

### Compared with other products in this space (proprietary and open source), what drew you to the project?

During the ML platform’s initial design phase four years ago, CyberAgent evaluated multiple architectural options: a **custom-built (DIY) solution**, a **pure Knative + Istio deployment**, and **KServe (with Knative and Istio)**.  

The DIY approach presented significant complexity and operational overhead, making it less sustainable long-term. **KServe** was selected due to:

- **Native zero-scale and autoscaling support**, critical for optimizing costly GPU resources and reducing operational expenditures.
- Seamless **integration with Knative and Istio**, which simplified deployment and ongoing management.
- A mature, well-documented serving mechanism that reduced the need for bespoke solutions.

---

## Usage Scenario

### How long has your organization used the project?

KServe has been in **full production use** at CyberAgent for approximately **four years**.

### What were the main motivations to adopt the project and which key features do you use today?

The adoption was driven by the need to **serve a diverse range of ML frameworks** and meet the requirements of multiple internal business units. KServe’s **ServingRuntime** templating mechanism provided a unified, flexible way to deploy models across different frameworks and custom containers.

Key features in production use include:

- **Autoscaling and zero-scale** — enabling inference services to scale down to zero when idle and scale out dynamically upon receiving requests.
- **Demand-driven scaling**, eliminating the need for rigid scheduling and maximizing infrastructure utilization.
- **Multi-framework serving** with consistent operational processes.

### What is the current level of usage (pre-production, production) and scale?

KServe operates in **production** with around **80 active inference services**. Each service may consume between **1 and 8 GPUs** at peak demand. Due to zero-scale, not all services run concurrently; scaling decisions are made dynamically based on live demand.

### What version is used and what is your update cadence with the project?

- Current version: **KServe v0.14.0**
- Update frequency: **Annual**, or sooner for critical security patches and high-value new features.
- Upgrade history: ~8 upgrades since **KServe v0.6**, with coordinated updates alongside Istio to ensure zero downtime in production.

### Can you walk us through your adoption and integration experience? What challenges did you encounter?

Initial adoption was smooth, aided by **clear getting-started guides** and deployment documentation.  

Challenges emerged primarily during **production troubleshooting**, such as diagnosing container crashes or investigating traffic routing issues. The official documentation lacked deep diagnostic workflows for complex scenarios involving sidecars and layered routing (Knative + Istio). This sometimes required significant manual tracing and reverse-engineering of request flows.

### Did you find the information in the repo or the project docs valuable?

Yes. The **GitHub repository** and official documentation provided numerous **framework-specific examples** and deployment templates that were critical in the implementation phase.

### Has your implementation of the project provided measurable value?

Yes — with KServe, the ML platform was delivered in **six months** by a team of **four engineers**.  
Key measurable outcomes include:

- ~**40% cost savings** versus alternative solutions.
- Significant **resource efficiency** gains via autoscaling and zero-scale.
- Reduced operational burden thanks to standardized serving workflows.

### Do you have any future plans regarding the project?

Yes, CyberAgent has identified the following high-priority feature requests:

1. **Gateway API Inference Extension (GIE) support**
2. **Zero-scale support for raw deployment mode**
3. **Vertical Pod Autoscaling (VPA) integration**
4. **Dynamic Resource Allocation (DRA) support** in KServe

---

## Perception

### What is your perception in terms of the project’s:
- **Governance** – Strong and effective.
- **Community growth potential** – High, with consistent activity and new contributors.
- **Maintainer diversity and response** – Excellent responsiveness and healthy maintainer base.
- **Community openness** – Generally positive, though greater transparency in technical decision-making is desired, especially regarding alternative approaches considered during feature design.

### How are you participating in the project community?

Engagement is primarily through **GitHub issues and PRs**. Regular attendance at community meetings is limited due to **timezone constraints** (Japan vs. Europe/North America).

### Have you engaged with community members or maintainers directly?

Yes. Most engagement occurs via **Slack** and **GitHub**, supplemented by in-person discussions at **KubeCon**.  
Example: A recent feature introduced for raw deployment mode lacked public context; clarification was sought and provided via Slack and GitHub by maintainers.

---

## Project Strengths

### In your opinion, what are the overall strengths of the project?

- Comprehensive **multi-framework support**
- Advanced **autoscaling and zero-scale** capabilities
- Flexible **ServingRuntime** design enabling diverse workloads and internal customization

---

## Project Improvements

### Is there something that holds the project back from its ultimate potential?

- **Technical decision transparency** — Some proposals remain in private Google Docs instead of public GitHub discussions.
- **Timezone barriers** — Collaboration is harder for contributors in Asia.

### In your opinion, what can the project do better?

- Move **all proposal discussions** to open, searchable public platforms.
- Enhance **asynchronous communication** to better support global participation.

