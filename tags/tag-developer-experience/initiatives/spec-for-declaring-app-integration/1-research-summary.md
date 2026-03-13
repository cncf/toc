# Research Summary: Existing Specifications and Problem Space for Application Integration Dependencies

**Initiative:** [Specification for declaring application integration dependencies](https://github.com/cncf/toc/issues/1797)  
**Responsible group:** CNCF TAG Developer Experience  
**Document:** First deliverable - research and report on existing specs and previous attempts.

---

## Executive summary

Research on five existing specifications (CNAB, OCM, ORD, Radius, Score) confirms a well-understood problem space: all are declarative and differ in granularity, versioning, and resolution. No widely cited failed or abandoned specification was identified; solution coupling is the main constraint observed. The TAG should proceed with a specification. The initiative can be distinguished by pursuing an emergent, auto-generated dependency model derived from observed application behavior (e.g., via eBPF tooling) rather than manual declarative authoring.

---

## 1. Purpose and scope

This document summarizes research on existing specifications and prior attempts relevant to declaring a service’s integration dependencies (e.g., APIs, databases, caches, blob stores, message buses) so that deployment teams know what must be available for an application to run successfully.

The primary source is the [Application Dependencies Research](https://github.com/api-evangelist/application-research) repository, a contribution to this initiative. That repository evaluates five specifications in depth and produces a dependency comparison, shared vocabulary, and recommendations. This report synthesizes that work and adds a brief review of failed or abandoned attempts. Detailed artifacts (JSON Schema, examples, OpenAPI, apis-json indices) remain in the linked repository.

---

## 2. Problem statement

Packaging an application so that it “runs anywhere” typically addresses execution of the process, not the availability of dependencies. Applications can exit quickly or return errors when required backing services are missing (databases, caches, peer APIs, message buses, etc.). A project-agnostic specification for declaring these integration dependencies would give application and deployment teams a common, standardized reference. It would not replace direct collaboration but would support it and enable downstream use cases such as abstraction coverage (e.g., Dapr), integration mocking (e.g., Microcks), and behavior-focused security (e.g., Software Bill of Behavior).

---

## 3. Research source and methodology

The [Application Dependencies Research](https://github.com/api-evangelist/application-research) repository was produced as a contribution to this initiative. [Kin Lane](https://github.com/kinlane) (Naftiko) is a key participant in the initiative and led this research. The methodology included:

- Extracting and analyzing JSON Schema across all evaluated specifications to capture core properties
- Summarizing specification priorities (properties and schema focus) and operational priorities (governance, tooling, community)
- Producing examples, mock APIs (OpenAPI for use with Microcks), and apis-json indices
- Evaluating each specification’s approach to dependencies and producing a structured [dependencies comparison](https://github.com/api-evangelist/application-research/blob/main/dependencies-comparison.md)

The repository also provides per-specification profiles (operations, schema, dependencies) for each of the five specs. This report focuses on the markdown documentation and the dependencies comparison; the full repository should be consulted for schemas, examples, and tooling artifacts.

---

## 4. Existing specifications evaluated

The research evaluates five specifications that relate to application definition, packaging, or dependency declaration. All are declarative: they describe what an application or workload needs in a human- or tool-authored format.

### 4.1 Cloud Native Application Bundle (CNAB)

CNAB is a package format for bundling, installing, and managing distributed applications in a cloud-agnostic way. Dependencies are expressed as bundle-to-bundle references with semantic version ranges. The specification includes image inventories, credential and parameter declarations, and invocation images for installation. It is well suited to reliable, repeatable installation and air-gapped or enterprise deployment. CNAB is maintained by the CNAB community; tooling includes Porter (CNCF Sandbox).

### 4.2 Open Component Model (OCM)

OCM describes software components and their relationships for lifecycle and compliance. Dependencies are component references with version constraints. It supports multiple resource types (e.g., container images, Helm charts), source traceability, and signing. It targets software supply chain management, compliance, and distribution across environments. The specification and tooling are maintained by the Open Component Model project.

### 4.3 Open Resource Discovery (ORD)

ORD is a protocol for publishing and discovering application and service metadata. It describes integration requirements at the API and event contract level for discovery and catalog use. Dependencies are structured as integration dependencies with aspects, API resources, and event resources; they support minimum versions, mandatory/optional semantics, and explicit direction (inbound/outbound). It is oriented toward enterprise API catalogs, service discovery, and integration documentation. The specification is maintained by the Open Resource Discovery community.

### 4.4 Radius

Radius is an open-source, cloud-native application platform. It defines runtime connections between application components and backing infrastructure. Dependencies are expressed as named connections (e.g., database, cache) with resource identifiers (UCP-style), optional IAM configuration, and environment variable injection. It supports recipe-based provisioning and application graph visualization. Radius is a CNCF project. The initiative description notes that Radius packaged the spec with the solution, so the spec has evolved mainly within the constraints of that solution, with a strong focus on Azure services; a project-agnostic spec would allow a solution-independent language that conforming projects (including Radius) could implement.

### 4.5 Score

Score is a platform-agnostic, developer-centric workload specification. Dependencies are abstract resource types (e.g., postgres, redis, volume) with optional class and shared-resource identity; they are referenced via placeholders resolved at deployment. It does not mandate version constraints; the implementation chooses versions. Score is a CNCF project and is designed for portability across platforms (e.g., score-compose, score-k8s, score-humanitec).

---

## 5. Dependencies comparison

The repository’s [dependencies comparison](https://github.com/api-evangelist/application-research/blob/main/dependencies-comparison.md) structures how each specification models dependencies.

| Aspect | ORD | Radius | Score | CNAB | OCM |
| ------ | --- | ------ | ----- | ---- | --- |
| Primary use case | API/service discovery & catalog | Cloud-native app deployment | Platform-agnostic workload spec | Application packaging & installation | Component lifecycle |
| Dependency granularity | API/event level | Resource connections | Abstract resource types | Image/invocation level | Component references |
| Version constraints | minVersion on aspects | None explicit | None | SemVer ranges | Exact or ranges |
| Direction | Explicit (inbound/outbound) | Implicit via connections | Implicit via placeholders | Implicit via credentials | Explicit source→target |

Findings from the comparison:

- **Declaration style:** ORD uses structured objects with explicit metadata; Radius uses named connection maps and resource URIs; Score uses type-based abstract resources; CNAB uses bundle references and version ranges; OCM uses component references and versions.
- **Resolution:** ORD resolves at discovery/design time; Radius and Score at deployment time (resource lookup or implementation provisioning); CNAB at installation time (bundle resolution); OCM at build/deploy time (descriptor resolution).
- **Sharing:** Each spec has a different mechanism (e.g., correlation IDs/groups in ORD, shared resource ID in Radius, type+class+id in Score, parameters/credentials in CNAB, component references in OCM).
- **Optional vs required:** ORD has explicit mandatory flags; Radius and Score treat declared connections/resources as required; CNAB distinguishes requires vs provides; OCM references are required unless expressed via labels/annotations.

The comparison concludes that these schemas can be complementary (e.g., Score workloads deployed via Radius, OCM components containing CNAB bundles, ORD describing APIs of applications deployed with any of the others).

---

## 6. Shared vocabulary and operational priorities

The research aggregated specification properties (schema and property-level focus) and operational properties (governance, tooling, community) across the five specifications. A common vocabulary does exist: the projects overlap on many concepts—dependencies, resources, configurations, metadata, versions, and the like—but with varying degrees of uniformity. Naming, structure, and emphasis differ by spec (e.g., ORD’s integration dependencies and aspects vs. Score’s type-based resources vs. Radius’s named connections). That overlap supports mapping and alignment between specs and can inform the design of any new dependency-focused specification; the variation indicates that a single, uniform vocabulary has not yet coalesced across the ecosystem.

---

## 7. Recommendations from the research

The Application Dependencies Research repository recommends:

- **Specification:** Create a new specification focused on application dependencies, expressible in JSON or YAML.
- **ABNF grammar:** Define an Augmented Backus-Naur Form (ABNF) grammar to map the new spec’s properties to existing spec properties.
- **Converter:** Implement a converter driven by that grammar for use by other tooling.
- **Reference:** Publish documentation and examples covering a wide range of dependency types.
- **Sandbox:** Provide OpenAPI, JSON Schema, and Microcks examples for a reference application dependency sandbox.
- **Services and tooling:** Maintain catalogs of reference services and tooling used in the reference architecture and sandbox.
- **Experience:** Align developer experience with the operational priorities observed across the evaluated specifications.
- **Governance:** Apply governance rules (e.g., Spectral or Vacuum) for completeness and consistency across specifications.

---

## 8. Failed or abandoned attempts

A search for failed, abandoned, or deprecated specifications specifically for declaring application integration dependencies in the cloud-native space did not identify a well-documented, widely cited abandoned spec. Findings:

- CNAB has not achieved very broad adoption relative to early expectations; it remains in use (e.g., via Porter) and is not deprecated. No formal “failure” or end-of-life was found.
- General cloud-native deprecations (e.g., Ingress NGINX controller end-of-life) concern specific implementations, not dependency-declaration specifications.
- In other ecosystems (e.g., Python PEP 508, Gradle dependency notation), dependency declaration formats have evolved or been refined rather than abandoned; those are not direct analogues to application integration dependency specs.

**Conclusion:** There is no clearly documented “failed” specification that the initiative could resurrect. The main lesson from the research is that solution-coupled specifications tend to evolve within the solution’s constraints. A project-agnostic specification would avoid that coupling.

---

## 9. Conclusion and recommendation

**Problem space:** The research confirms that multiple specifications already model application or workload structure and dependencies (CNAB, OCM, ORD, Radius, Score). All of them are declarative: they assume human- or tool-authored declarations of what an application needs. They differ in granularity (API/event vs resource vs bundle/component), versioning, resolution time, and sharing semantics. The comparison and shared vocabulary provide a basis for aligning or extending existing specs or for defining a new dependency-focused spec. No widely cited failed or abandoned specification was identified; the main risk noted is solution coupling, which constrains evolution and portability.

**Recommendation:** The TAG should **proceed with a specification** for declaring application integration dependencies.

The recommended direction is distinguished from the five evaluated specifications in an important way: those specifications are **declarative** (authors write or generate a static description of dependencies). The initiative can pursue an **emergent** specification, where dependency declarations are **derived from observed application behavior** (e.g., via eBPF-based system call hooks) and **automatically generated**. That places the primary emphasis on **auto-generation** of dependency declarations during development or observation, rather than on manual declarative authoring. Such a spec would complement existing declarative specs (e.g., by feeding into Score, Radius, or ORD) while addressing a different source of truth and workflow. Subsequent deliverables should therefore include a plan and project description for automatic generation of dependency declarations that adhere to the spec, in addition to identifying or defining the specification itself and providing consumer guidance.

---

## 10. References

- [CNCF TOC Issue #1797 — Specification for declaring application integration dependencies](https://github.com/cncf/toc/issues/1797)
- [Application Dependencies Research](https://github.com/api-evangelist/application-research) (Kin Lane / Naftiko contribution to this initiative)
- [Application Dependencies Research — Dependencies comparison](https://github.com/api-evangelist/application-research/blob/main/dependencies-comparison.md)
- [Application Dependencies Research — Cloud Native Application Bundle](https://github.com/api-evangelist/application-research/blob/main/cloud-native-application-bundle.md)
- [Application Dependencies Research — Open Component Model](https://github.com/api-evangelist/application-research/blob/main/open-component-model.md)
- [Application Dependencies Research — Open Resource Discovery](https://github.com/api-evangelist/application-research/blob/main/open-resource-discovery.md)
- [Application Dependencies Research — Radius](https://github.com/api-evangelist/application-research/blob/main/radius.md)
- [Application Dependencies Research — Score](https://github.com/api-evangelist/application-research/blob/main/score.md)
