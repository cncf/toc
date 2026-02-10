# CNCF TAG Infrastructure

## Mission

To define and advance practices, standards, and assessments related to core cloud native infrastructure areas that support and enable scalable, resilient, secure, and performant cloud native systems, applications, and architectures. This supports the CNCF's technical vision by addressing critical problems faced by adopters and contributing to a robust cloud native ecosystem.

## Scope

This TAG covers the following sub-domains and topics within its scope:
- Data: Data refers to all digital information created, handled, and utilized by applications. Both structured and unstructured forms of data are included. This includes, but is not limited to the following.
  - User interactions
  - System logs
  - Performance metrics
  - Configuration settings
  - Permanent records like database entries
- Storage: Storage used in cloud native environments include block stores, file systems, object stores, databases, key-value stores, streaming and messaging, and related caching mechanisms. Our focus is on understanding the following fundamental characteristics and relating them to cloud native use cases.
  - Availability
  - Scalability
  - Performance
  - Durability
  - Consistency
  - Ease-of-use
  - Cost
  - Operational complexity
- Networking: Core connectivity and traffic management in cloud native systems. This can include:
  - DNS, gateways, load balancing, and service discovery
  - Service mesh for inter-service communication
  - Connectivity and traffic management across environments
  - Policy enforcement for network security and compliance
  - Observability and monitoring of network behavior
- Compute: Encompasses kernel and hypervisor-level compute primitives that underpin cloud native infrastructure (workloads running in Public/Private Cloud, Edge, IoT, Batch, Big Data, AI/ML/LLMs, etc). This can include:
  - Internal components of container runtimes (at the kernel/user boundary)
  - System-level resource isolation (namespaces, cgroups)
  - Hypervisor interactions, and node-level bootstrap agents
  - Primitives supporting specialized computing elements beyond CPUs, including DPUs, GPUs, TPUs, FPGAs, ASICs, etc.
- Infrastructure Management: End-to-end lifecycle management across public/private clouds, on-premises, edge, and hybrid environments. This can include:
  - Declarative and imperative IaC
  - Control planes and orchestration
  - Drift detection and correction
  - Policy enforcement and compliance
  - Testing and validation
  - Automation of operations
  - Observability and monitoring
  - Languages, schemas, and protocols enabling infrastructure-as-code tooling
- Edge: Edge computing extends cloud native principles to resource-constrained, distributed environments, often characterized by intermittent connectivity, diverse hardware, and location-aware workloads. This can include:
  - Edge-ready infrastructure
  - Cross-architecture workloads
  - Resilient multi-cluster deployments
  - Low latency, power efficiency (carbon footprint), remote ops
  - Cloud–edge interoperability
- Sovereignty: Ensures cloud native systems operate within jurisdictional, regulatory, and organizational constraints. This can include:
  - Infrastructure and data portability across regions and providers
  - Runtime policy enforcement for compliance and governance
  - Operational continuity in disconnected or restricted environments

## Out of Scope
- Scheduling, batch, serving, etc. are part of the TAG Workloads Foundation charter.
- Hardware design and procurement are out of scope for this charter
- Non cloud native technologies (according to the CNCF Cloud Native definition)
- Anything not explicitly mentioned in the scope above is out of scope.

## Anticipated Deliverables
Expected outputs of the TAG Infrastructure may include:
- Subprojects that represent ongoing services or programs requiring stewardship.
- Initiatives submitted to the TOC that are lightweight, time-bound, and objective-focused units of work. Examples include Cloud Native Disaster Recovery White Paper and multicluster performance optimization analysis.
- Other outputs consistent with advancing practices in the TAG's scope, such as frameworks, guidelines, or whitepapers.

These deliverables are intended to provide services to projects, other TAGs, or the TOC.

## Success Criteria
Success is based on the effective establishment and operation of Subprojects and Initiatives for the TAG, the creation and dissemination of valuable resources such as best practices and assessments, and the TAG's ability to maintain focus on its defined scope and align with TOC requirements. Success will also involve recruiting new leadership and community members into the TAG to support, drive, and establish deliverables within the TAG’s scope.

## Coordination
The TAG Infrastructure will coordinate with various stakeholders within the CNCF ecosystem:
- CNCF Projects: The TAG provides services to projects and its work helps align projects within the CNCF ecosystem.This coordination ensures alignment across the foundation and provides pathways for community focuses to be supported
- Other TAGs: Coordination is essential as TAGs serve needs across projects and other TAGs.
- TOC Subprojects: TAG leadership participates in TOC Subprojects like Project Reviews and Contributor Strategy.
- Community Groups: Community Groups are encouraged to discuss initiative ideas and may submit applications for initiatives within a TAG.

## Alignment with the CNCF TOC Charter
The TAG Infrastructure charter is directly aligned with the CNCF TOC charter. The TOC is the technical governing body responsible for maintaining the technical vision and driving common practices across projects. The TOC's vision is problem-centric, encouraging projects to solve challenges faced by adopters. By focusing on fundamental infrastructure problems such as Data, Storage, Network, Compute, and Service Mesh, the TAG Infrastructure directly addresses significant problems faced by cloud native adopters and projects in building and managing their underlying infrastructure. The TAG's work in defining practices, conducting assessments (as shown in example Subprojects and Initiatives), and potentially creating guidelines contributes to driving common practices across the ecosystem, as mandated by the TOC. The TAG's work in defining best practices, frameworks, and performing assessments contributes to driving common practices and aligning projects within the ecosystem, as mandated by the TOC.
