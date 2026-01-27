# Cloud Native and OCI Compliant Inner-Loop Tooling & Packaging for AI Engineers

https://github.com/cncf/toc/issues/1740

Integrating the AI developer inner loop into an end-to-end CI/CD process leveraging cloud-native technologies and tooling

## Initiative description

Focus on the developer inner loop, everything an AI engineer does on a laptop/desktop before code or models ever reach CI/CD in a cloud-native environment:

* Local container workspaces: Reference inner loop workflow using desktop tooling such as Podman Desktop / Podman AI Lab for root-less, GPU-aware experimentation, including template images for PyTorch/LLM stacks and volume-mounted datasets. ​
* Unified model build & run CLI: Hardening inference on developer machine and agentic frameworks to leverage container-based tooling so engineers can easily spin-up inference, RAG and multi-agent services locally with one command.
* Standard packaging of artefacts: Drive convergence between various implementations such as ModelKit, ModelCar towards the emerging ModelPack spec to create a single OCI-manifest that can hold model weights, metadata and SBOM.
* Inner-loop supply-chain security: Integrate Notary v2 / model authenticity and transparency via Sigstore, LF AI & Data Model Openness Framework-generated model & data cards, plus SBOM annotations directly into the OCI artefact so that security & openness are “baked in” before CI. ​
* Fast hand-off to outer loop: Provide reference GitOps flows (Flux/Argo) that pull the signed artefact into KServe with ModelPack image-mount optimisation, and register versions in Kubeflow Model Registry.

## Deliverable(s) or exit criteria

* An technical POC showing <10 min “idea-to-inference” path for cloud-native agent development on a developer laptop.
* Clearly documented standards for OCI artefact standardization across runtimes and registries.
* Specification / procedure to achieve MOF Class III compliant model distributions via any OCI registry.
* Standardised process for leveraging model signing with artefacts-level provenance to support a verified end-to-end CI/CD reference pipeline including outer loop for AI engineering.