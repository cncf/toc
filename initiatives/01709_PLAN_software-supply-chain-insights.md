# CNCF Software Supply Chain Insights

This is a plan for Initiative: https://github.com/cncf/toc/issues/1709.

## Table of Contents

- [CNCF Software Supply Chain Insights](#cncf-software-supply-chain-insights)
  - [Table of Contents](#table-of-contents)
  - [Overall Objective](#overall-objective)
  - [Phase 1: Proof of Concept (PoC) \& Foundational Work](#phase-1-proof-of-concept-poc--foundational-work)
    - [M1.1: PoC Scope, Initial GUAC Deployment \& Observability Baseline](#m11-poc-scope-initial-guac-deployment--observability-baseline)
    - [M1.2: PoC Data Ingestion, Validation \& Leveraging Existing Work](#m12-poc-data-ingestion-validation--leveraging-existing-work)
    - [M1.3: PoC Review, Performance/Cost Analysis \& Phase 2 Go/No-Go](#m13-poc-review-performancecost-analysis--phase-2-gono-go)
  - [Phase 2: Production Infrastructure \& Core Integrations](#phase-2-production-infrastructure--core-integrations)
    - [M2.1: Production GUAC Infrastructure, Advanced Deployment \& Scalability Planning](#m21-production-guac-infrastructure-advanced-deployment--scalability-planning)
    - [M2.2: Expanded Data Ingestion \& Automation Framework](#m22-expanded-data-ingestion--automation-framework)
    - [M2.3: Deepen GUAC Observability \& Acknowledge Community Support](#m23-deepen-guac-observability--acknowledge-community-support)
  - [Phase 3: Full Rollout, Ecosystem Integration \& Continuous Improvement](#phase-3-full-rollout-ecosystem-integration--continuous-improvement)
    - [M3.1: Complete Project Onboarding \& Landscape Integration](#m31-complete-project-onboarding--landscape-integration)
    - [M3.2: Launch Health Dashboards, Community Outreach \& Case Study Development](#m32-launch-health-dashboards-community-outreach--case-study-development)
    - [M3.3: Sustained Operations, Contribution \& Iterative Enhancements](#m33-sustained-operations-contribution--iterative-enhancements)
  - [Call for Contribution: Join Us in Shaping CNCF's Future](#call-for-contribution-join-us-in-shaping-cncfs-future)

## Overall Objective

Establish a scalable, GUAC-based system to automate the collection, correlation, and understanding of CNCF Project Software Supply Chain security metadata (SBOM, CycloneDX, attestations). This initiative is aligned with the **upcoming GUAC 1.0 release** ([DRAFT](https://docs.google.com/document/d/1toZFd2ynysmiDIQSZ8UJbiaJfDG7Twt_6P0WPQWyRE0/edit?tab=t.0)), which marks a major milestone for the project, including standardization on PostgreSQL as a primary backing store, enhancing its enterprise readiness. We **attended the GUAC maintainer's meeting meeting today (June 2nd, 2025\)** ([meeting notes](https://docs.google.com/document/d/1bjkxGu1YcMurav5DU-imoKfEN2j-LuNQJFOa__A-m0I/edit?tab=t.0#heading=h.lxhckb5wv0q5)). This plan is informed by their input and feedback.

This initiative will actively engage with the **OpenSSF and GUAC project communities**, positioning CNCF as both a key **Adopter and future Contributor** to the Project, as well as provide a visible collaboration between Foundations (CNCF, OpenSSF). It will also inform and potentially align with broader CNCF initiatives around **conformance testing and resource consumption analysis tools** for projects and their dependencies (including CVE data sources). The **strong, ongoing collaboration with GUAC maintainers** is a key feature of this plan, ensuring a smooth integration and alignment with the GUAC project roadmap.

## Phase 1: Proof of Concept (PoC) & Foundational Work

**Goal:** Validate GUAC's suitability for CNCF data, establish core infrastructure with a limited project set, and leverage early learnings and **direct feedback from GUAC maintainers**.

### M1.1: PoC Scope, Initial GUAC Deployment & Observability Baseline

* **Define:** Key data points (e.g., repo URL, contacts, incubation level, CLOMonitor) for 3-5 diverse CNCF projects.  
* **Deploy:** Stand up a PoC GUAC instance (PostgreSQL backend) with Kubermatic (https://github.com/mfahlandt et al). Secure and establish basic monitoring.  
* **Assess Observability:** Quantify existing GUAC observability levels. **Solicit feedback directly from GUAC maintainers** (building on discussions from the maintainer meeting) on prior art and desired future state for observability, noting their strong interest and willingness to collaborate.  

### M1.2: PoC Data Ingestion, Validation & Leveraging Existing Work

* **Acknowledge Early Work:** Incorporate learnings from **Mario's PoC ingestion scripts (bash) and experiments with GitHub runners on Oracle Cloud (OKE)**.  
* **Develop:** Initial data collectors (scripts/GUAC native) for PoC sources (e.g., cncf/landscape YAML, CLOMonitor API, GitHub API).  
* **Ingest & Verify:** Define minimal GUAC schema, load data, and validate accuracy. Demonstrate basic GraphQL queries for data retrieval.  

### M1.3: PoC Review, Performance/Cost Analysis & Phase 2 Go/No-Go

* **Present:** Showcase PoC findings, capabilities, and learnings (including initial observability assessment) to TOC and key stakeholders, highlighting collaboration with GUAC maintainers.  
* **Measure PoC Performance/Cost:** Analyze the resource consumption (CPU, memory, storage, network for data ingress/egress) and query performance of the PoC deployment. This data is crucial for informing Phase 2 infrastructure planning and will be shared with GUAC maintainers.  
* **Decide:** Greenlight Phase 2 based on PoC success, performance/cost analysis, and refined full-scale requirements.

## Phase 2: Production Infrastructure & Core Integrations

**Goal:** Build out robust production infrastructure, expand data sources, and integrate advanced operational practices, aligning with GUAC ecosystem advancements and **insights from GUAC maintainers**.

### M2.1: Production GUAC Infrastructure, Advanced Deployment & Scalability Planning

* **Architect:** Design HA production GUAC setup (scalable services, resilient PostgreSQL) for all 220+ projects, informed by PoC performance/cost data.  
* **Deploy Strategy:**  
  * Implement GUAC deployment using **Crossplane Resources** or a **custom Kubernetes controller (e.g., via kcp)** for superior lifecycle management over standalone Helm.  
  * Concurrently, enhance the official GUAC Helm chart with production hardening and best practices, in consultation with GUAC maintainers.  
* **Scalability & Ecosystem Alignment:**  
  * Plan for scale-out by investigating and leveraging **additional tools in the GUAC ecosystem**.  
  * Align with GUAC project goals of integrating more deeply with **DataFusion & the Rust ecosystem**.  
  * **Placeholder:** Note OpenTelemetry's parallel efforts in Rust and potential future synergies.  
  * **Caching & Reliability:** Based on PoC findings and **discussions with GUAC maintainers regarding scale and reliability best practices**, plan for and implement appropriate caching mechanisms (e.g., for frequently accessed graph segments, external API call results) and other optional enhancements to ensure performance and resilience.  

### M2.2: Expanded Data Ingestion & Automation Framework

* **Prioritize:** Catalog all target data sources (DevStats, security audit links, CII Badging, TAG membership, etc.) and prioritize integration, considering potential rate limits and access challenges for community data.  
* **Develop Collectors:** Build and test production-grade collectors for prioritized data sources.  
* **Automate:** Implement a robust scheduling system (e.g., Argo Workflows, Kubernetes CronJobs) for regular (target: weekly or better) data ingestion.  
* **Finalize Schema:** Define and document the comprehensive GUAC graph schema.  

### M2.3: Deepen GUAC Observability & Acknowledge Community Support

* **Helm Enhancements:** Augment the GUAC Helm chart (acknowledging the **GUAC team's recent fast turnaround in adding it, a testament to our collaborative engagement**) to include dependencies/configurations for linked observability tools. Engage with CNCF Projects at all levels of the landscape to contribute configs for their projects, such as Perses (https://perses.dev) dashboards.).  
* **OpenTelemetry Initiative:**  
  * **Action:** Create a formal GitHub issue within the GUAC project (or contribute to existing ones, per discussion with maintainers) to advocate for and specify deeper **OpenTelemetry integration**.  
  * **Scope:** Request comprehensive support beyond current metrics, including **standardized logging output, distributed tracing context propagation, and profiling hooks**, aligning with maintainer interest.

## Phase 3: Full Rollout, Ecosystem Integration & Continuous Improvement

**Goal:** Onboard all CNCF projects, integrate with key CNCF platforms, deliver actionable insights, contribute back to the community (leveraging our **strong partnership with GUAC maintainers**), and establish long-term operational excellence.

### M3.1: Complete Project Onboarding & Landscape Integration

* **Full Ingestion:** Systematically ingest data for all remaining CNCF projects into the production GUAC instance.  
* **Automate landscape.cncf.io:** Develop and deploy mechanisms to automatically update relevant project information on the CNCF Landscape using GUAC as the source of truth.  

### M3.2: Launch Health Dashboards, Community Outreach & Case Study Development

* **Develop Dashboards:** Create and launch V1 of TAG & Project Health Dashboards, providing actionable insights derived from GUAC data.  
* **Communicate:** Execute a communication plan (blog posts, documentation, demos) to inform the CNCF community about the new capabilities and data availability.  
* **Develop Case Study:** Document the experience of setting up and scaling GUAC for CNCF's software supply chain security reporting and analysis. This will serve as a valuable **case study for the GUAC project and OpenSSF (co-developed or reviewed with GUAC maintainers)**, providing much-needed real-world deployment, scalability numbers, and cost/performance benchmarks.  

### M3.3: Sustained Operations, Contribution & Iterative Enhancements

* **Maintain & Optimize:** Ensure ongoing system health, performance, and scalability.  
* **Contribute Back:** Actively contribute improvements, learnings, and potentially code (e.g., related to observability, collectors, scalability enhancements) back to the GUAC project and OpenSSF, solidifying our role as a key partner.  
* **Evolve:** Continuously refine data models, add new data sources, and develop new features based on community feedback and CNCF strategic needs, including insights relevant to the **CNCF conformance/resource consumption tools initiative**.

## Call for Contribution: Join Us in Shaping CNCF's Future

This initiative represents a significant opportunity to contribute to a high-impact CNCF effort that will enhance how the entire ecosystem understands and manages project health, onboarding, and software supply chain security. We are building a foundational data platform with GUAC, and its success hinges on diverse expertise.

While we have a strong core of engineering talent available and committed, we are actively seeking passionate individuals from other disciplines to help us realize the full potential of this project. Your skills will be crucial in ensuring this initiative is not only technically sound but also widely adopted, well-understood, and effectively communicated. This is a chance to make a tangible difference, engage with cutting-edge technology, and collaborate closely with the CNCF Technical Oversight Committee (TOC), the CNCF Technical Advisory Board (TAB) for guidance and strategic alignment, and the vibrant GUAC and OpenSSF communities.

We are looking for contributors in the following areas:

* **Project Managers:** To help steer planning, track progress, manage stakeholder communications, and ensure we meet our milestones effectively.  
* **Content Creators / Technical Writers:** To develop clear, engaging documentation, blog posts, case studies, and communication materials for various audiences, from technical contributors to the broader CNCF community.  
* **Technical Product Marketers:** To help define the value proposition, strategize rollout and adoption, gather user feedback, and ensure the outputs of this initiative meet the evolving needs of CNCF projects and the wider ecosystem.  
* **Data Visualization Specialists:** (Implicit in "Health Dashboards" but good to call out) To help design and implement intuitive and actionable dashboards that make complex project data accessible and understandable.  
* **Community Liaisons / Advocates:** To help foster communication between this initiative, TAGs, projects, and the broader community, ensuring feedback is gathered and the project's progress and benefits are widely understood.  
* **Software Supply Chain Domain Experts:** To provide deep insights into supply chain security best practices, threat models, and emerging standards, ensuring our data collection and analysis are comprehensive and relevant.  
* **Cloud Security Professionals & Practitioners:** To offer practical expertise on securing cloud-native applications and infrastructure, advising on data points relevant to security posture, and helping to translate GUAC insights into actionable security improvements for CNCF projects.

If you are passionate about open source, software supply chain security, and enabling the CNCF ecosystem, we encourage you to reach out and explore how you can contribute!
