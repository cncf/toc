# Cloud Native and OCI Compliant Inner-Loop Tooling & Packaging for AI Engineers

https://github.com/cncf/toc/issues/1740

Integrating the AI developer inner loop into an end-to-end CI/CD process leveraging cloud native technologies and tooling

## Initiative description

Focus on inner loop development which incorporates everything an AI engineer does on a local environment before code or models ever reach CI/CD in a cloud native environment:

* Local container workspaces: Reference inner loop workflow using desktop tooling such as Podman Desktop / Podman AI Lab for root-less, GPU-aware experimentation, including template images for PyTorch/LLM stacks and volume-mounted datasets. ​
* Unified Inner-Loop CLI: Hardening local inference and agentic frameworks via container-based tooling, allowing engineers to spin up inference, RAG, and multi-agent services locally with a single command.
* Standard packaging of artefacts: Drive convergence between various implementations such as ModelKit, ModelCar towards the emerging ModelPack spec to create a single OCI-manifest that can hold model weights, metadata and SBOM.
* Inner-loop supply-chain security: Integrate Notary v2 / model authenticity and transparency via Sigstore, LF AI & Data Model Openness Framework-generated model & data cards, plus SBOM annotations directly into the OCI artefact so that security & openness are “baked in” before CI. ​
* Fast hand-off to outer loop: Provide reference GitOps flows (Flux/Argo) that pull the signed artefact into KServe with ModelPack image-mount optimisation, and register versions in Kubeflow Model Registry.

## Deliverable(s) or exit criteria

* An technical POC showing <10 min “idea-to-inference” path for cloud native agent development on a developer laptop.
* Clearly documented standards for OCI artefact standardization across runtimes and registries.
* Specification / procedure to achieve MOF Class III compliant model distributions via any OCI registry.
* Standardised process for leveraging model signing with artefacts-level provenance to support a verified end-to-end CI/CD reference pipeline including outer loop for AI engineering.

## Project Scope & Intent - Cloud Native AI Developer Workflow Interoperability

### Overview and Intent
AI developers today frequently work in fragmented local environments that are disconnected from cloud native operational workflows. While emerging standards like ModelPack and OCI-aligned AI artifact initiatives provide the “packaging” foundations, there is no unified interoperability specification that defines how these artifacts must be structured, secured, and described to move seamlessly from a developers environments into a Kubernetes-based production system.

The goal of this initiative is to define a minimal Interoperability Specification (a “Compliance Profile”) for AI Artifacts. Rather than rebuilding the OCI layer structure, this initiative defines the **Standardized Metadata Contract** that must exist on top of packaging formats like **ModelPack**. This ensures that any <a href="https://github.com/cncf/foundation/blob/main/style-guide.md#1-cloud-native-and-open-source" target="_blank">“Cloud Native Ready”</a> AI artifact contains the mandatory identity, security, and runtime information required that enables a cohesive developer inner loop and GitOps-driven delivery.

This initiative intentionally builds on existing OCI-aligned packaging efforts rather than redefining artifact storage or layer mechanics.

### Scope Overview
This initiative defines the **Interoperability Layer** for AI artifacts, bridging the gap between raw packaging and operational deployment. 

Within this scope, the initiative will explore and document:
* **An Interoperability Profile Spec:** A set of mandatory annotation conventions and metadata requirements (the “Manifest Contract”).
* **Compliance & Trust Requirements:** Standards for signing, SBOMs, and openness classification.
* **Workflow Reference Patterns:** Validating the spec through “Local Environment-to-Cluster” GitOps and runtime integration.

The initiative is intended to encourage ecosystem alignment and workflow interoperability rather than define new standalone packaging specifications or runtime standards.

### In-Scope Areas
#### 1. The Interoperability Specification
Define a structured, minimal specification for AI artifacts to be considered “Cloud Native Interoperable”. This does not define OCI layering but specifies the mandatory metadata:
* **Annotation Conventions:** Standardize keys for runtime frameworks (e.g., vllm), hardware accelerators (e.g., nvidia-gpu), and lifecycle status.
* **Agentic Assets:** Standardizing the packaging of “skills”, prompt templates and workflow definitions. 
  * To ensure interoperability, the internal format for skills will align with the <a href="https://agentskills.io/home" target="_blank">agentskills.io</a> community standard. 
  * The spec defines how these standardized skills are encapsulated into the OCI layers for consistent distribution and discovery.
  * The initiative may leverage Skill DLC as the primary reference for demonstrating how these assets are dynamically loaded and managed.
* Interface Definitions: Define the "Ingredient List” for the different classes of artifacts (Models, RAG contexts, and Agentic Assets).

This includes defining how artifacts relate to and compose with one another.

#### 2. Metadata, Relationships, & MOF Mapping
Define how artifacts describe themselves and their dependencies to enable cross-tool discovery:
* **MOF-to-OCI Mapping:** Standardize how the LF AI & Data Model Openness Framework (MOF) classifications (e.g., Class I, II, III) are represented as machine-readable OCI metadata.
* **Lineage & Authorship:** Standardizing metadata for provenance, versioning, and authorship to ensure clear ownership as artifacts move from local environments to registries.
* **Relationship Mapping:** Defining minimal relationships to metadata conventions between related artifacts (e.g., model → skill → pipeline) within the OCI manifest.
* **Dependencies & Compatibility:** A schema for declaring software dependencies and infrastructure requirements (e.g., specific CUDA versions or vRAM minimums) to ensure cross-runtime interoperability.
* **Large Binary Asset Optimization:** Establishing best practices and metadata conventions for the efficient handling of large-scale binary artifacts (multi-GB model weights) within OCI registries, focusing on registry compatibility and layer deduplication. 
* **Alignment with CNCF:** Build on existing efforts (e.g., <a href="https://modelpack.org/" target="_blank">ModelPack</a>, <a href="https://kitops.org/docs/modelkit/intro/" target="_blank">ModelKit</a>, <a href="https://github.com/redhat-ai-services/modelcar-catalog" target="_blank">ModelCar</a>)

#### 3. Supply Chain Security and Transparency
Define the mandatory “Trust Profile” for AI artifacts to ensure they are verifiable before entering production:
* **Cryptographic Identity:** Standardize artifact signing and verification using Sigstore and Notary v2 at the point of creation on a developer's machine.
* **Transparency Manifests:** Mandatory requirements for SBOM (Software Bill of Materials) generation and attachment for all artifact layers.
* **Provenance Metadata:** Defining the "Hardened Provenance" requirements to ensure the journey from local experimentation to a secure registry is immutable and documented.

The goal is to ensure artifacts are trusted and verifiable before entering CI/CD pipelines.

#### 4. Developer Inner-Loop & Workflow Interoperability
Define the operational patterns that allow the specification to be utilized in a portable "local environment-to-cluster" journey.
* **Workflow Consistency:** Documenting how existing OCI-aligned tools (ModelPack, ModelKit, ModelCar) can produce artifacts that adhere to this initiative's compliance spec.
* **Local Execution Patterns:** Reference patterns for running specified artifacts in local, container-based environments to ensure high-fidelity parity with remote clusters.
* **Rapid Iteration Flow:** Validation of the spec through a reference implementation targeting a sub-10-minute "idea-to-inference" experience.

#### 5. GitOps and Kubernetes Integration Patterns
Define the "Handoff" patterns for how artifacts transition into production cloud native systems.
* **GitOps Delivery Patterns:** Reference architectures for pulling compliant artifacts into Flux or Argo CD workflows.
* **Runtime Integration:** Standardized patterns for the seamless deployment of artifacts into serving platforms (e.g., <a href="https://kserve.github.io/website/" target="_blank">KServe</a>, <a href="https://github.com/vllm-project/vllm" target="_blank">vLLM</a>) and registration into model registries (e.g., Kubeflow Model Registry).
* **Enterprise Requirements:** Ensuring the promotion spec accounts for air-gapped, regulated, and hybrid-cloud infrastructure constraints.

#### 6. Real-World Deployment Considerations
Ensure the approach accounts for:
* Air-gapped and regulated environments
* Enterprise security and compliance requirements
* Regulated environments
* Hybrid and multi-platform infrastructure
* Resource-constrained developer environments

This ensures the solution is practical and broadly applicable.

#### 7. Ecosystem Collaboration
This initiative will be developed in collaboration with:
* ModelPack and related OCI-aligned initiatives
* CNCF projects
* LF AI & Data communities
* OpenSSF and supply chain security initiatives
* Kubernetes AI and platform engineering communities

The intent is to align efforts across communities rather than define a solution in isolation.

### Out of Scope Areas
* Define new low-level binary compression or OCI layering techniques (deferring to OCI/ModelPack)
* Define model architectures or ML training frameworks
* Mandate specific vendor-locked developer tools
* Standardize outer-loop pipelines beyond reference integration patterns

### Definition of Success
* **A Published Interoperability Spec:** A validated specification that existing tools can adopt to ensure cloud native readiness.
* **Cross-Tool Portability:** Demonstrated ability for an artifact built by one tool to be verified and deployed by a different runtime.
* **The "10-Minute Flow":** A successful reference implementation demonstrating the journey from a local idea to a running inference service on Kubernetes.
* **Ecosystem Alignment:** Broad adoption of the "Compliance Profile" metadata across CNCF and LF AI & Data communities.

