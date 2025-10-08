# OpenFGA Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1287)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for OpenFGA

### Criteria Evaluation

# Criteria Evaluation

$TOCMember conducted the due diligence of **OpenFGA**, which applied for incubation. The project has <TBD> the criteria that demonstrate its maturity at incubation.

## Noteworthy Implementations

The OpenFGA team has shown substantial progress across CNCF incubation criteria:

- **Open Governance Structure**:  
  - Governance is documented in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md), with clear roles for maintainers, TSC responsibilities, and decision-making processes.  
  - While the governance model follows CNCF guidelines, the current maintainer pool is primarily Okta-driven, which raises neutrality concerns noted by adopters.

- **Strong Community and Contributor Growth**:  
  - Over 100 active contributors in the past year ([DevStats](https://all.devstats.cncf.io/d/53/projects-health-table?orgId=1)), with consistent high contribution velocity (1,400+ contributions/month).  
  - Public community meetings occur monthly with agendas and notes available ([community page](https://openfga.dev/docs/community#monthly-community-meetings)).  
  - RFC-driven roadmap development ([RFC repo](https://github.com/openfga/rfcs)) enables transparent technical decision-making.

- **Robust Engineering and Security Practices**:  
  - Documented [release process](https://github.com/openfga/openfga/blob/main/RELEASES.md) with regular releases and QA checks.  
  - Completed [security self-assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/self-assessment.md) and [joint assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/joint-assessment.md).  
  - Holds the [OpenSSF Best Practices badge](https://www.bestpractices.dev/en/projects/6374).

- **Industry Adoption**:  
  - Production deployments at **Canonical, Docker, Agicap, Read AI, Stacklok, and Configu**, demonstrating scalability to billions of tuples and thousands of requests/sec.  
  - Adopters report high-quality documentation, responsiveness of maintainers, and strong security model.

- **Ecosystem Integrations**:  
  - Integrates with CNCF projects including Kubernetes/Helm, OpenTelemetry, Prometheus, and Grafana.  
  - SDKs in multiple languages lower adoption barriers.

## Blockers (Previously Raised, Now Resolved)

- **Security Review**: Completion of [security self-assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/self-assessment.md) and [joint assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/joint-assessment.md), addressing design and threat modeling gaps.  
- **Governance Documentation**: Expanded governance and maintainer details in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) and [MAINTAINERS.md](https://github.com/openfga/.github/blob/main/MAINTAINERS.md), ensuring clarity in contributor pathways and decision-making.  
- **Adopter Verification**: Multiple independent production adopters documented in [ADOPTERS.md](https://github.com/openfga/community/blob/main/ADOPTERS.md), confirming real-world use cases.

## Recommended Enhancements (Non-Blocking)

The following items remain open and may not block incubation, but should be completed before graduation:

- **Maintainer Diversity and Neutrality**: Expand the maintainer base beyond Okta to reduce single-vendor dependence and strengthen long-term neutrality. *(Tracking issue: TBD)*

- **Transparency in Support**: Adopter feedback revealed reliance on **private support channels with Okta maintainers**, which creates inequity and goes against CNCF’s recommendation for open, transparent community support. The project should discourage private support channels and ensure all discussions happen in public forums.

- **Governance Evolution**: Document clearer processes for roadmap approvals, leadership decisions, and potential subproject governance. *(Tracking issue : TBD)*

- **Maintainer Lifecycle**: Define and demonstrate onboarding, offboarding, and emeritus processes. *(Tracking issue : TBD)*

- **Developer Experience**: Address adopter feedback on DSL tooling improvements, batch operations, and idempotent writes. *(Tracking issue : TBD)*

- **Community Outreach**: Increase visibility through case studies, blogs, and talks to grow adoption beyond early contributors. *(Tracking issue: TBD)*

### Adoption Evaluation

_The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY._

### Final Assessment

_[The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS]._

### Due Diligence

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**
  - OpenFGA presented to the TAG Security on [2022-06-01](https://github.com/cncf/tag-security/issues/902), with a [recording of the session](https://www.youtube.com/watch?v=un2KKTE25qg) to increase awareness of the project in the security context.

- [x] **TAG provides insight/recommendation of the project in the context of the landscape.**
  - TAG Security discussed OpenFGA and provided feedback in [issue #1339](https://github.com/cncf/tag-security/issues/1339), confirming the relevance of the project in the context of cloud-native authorization.

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**
  - OpenFGA manages its own [communication channels](https://contribute.cncf.io/maintainers/community/vendor-neutrality/#communication), holds community events and meetings on [vendor-neutral hosting](https://contribute.cncf.io/maintainers/community/vendor-neutrality/#hosting), makes architectural decisions in a collaborative manner [Architectural decisions](https://contribute.cncf.io/maintainers/community/vendor-neutrality/#architectural-decisions), and follows self-governance per [Governance](https://contribute.cncf.io/maintainers/community/vendor-neutrality/#governance).

- [x] **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
  - OpenFGA completed this step during its Sandbox application on 2024-03-15.

- [x] **Due Diligence Review.**
  - Completion of this due diligence document, resolution of concerns, and presentation for public comment satisfies the criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**
  - Documentation and examples include [Running OpenFGA in Production](https://openfga.dev/docs/getting-started/running-in-production), [OpenFGA CLI](https://github.com/openfga/cli), and [sample stores](https://github.com/openfga/sample-stores#sample-stores) to help users integrate OpenFGA in their applications.

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [ ] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

OpenFGA has documented iterative governance updates in the community repository. Relevant discussions can be found in [OpenFGA community issues](https://github.com/openfga/community/issues) tracking governance changes. 

- [x] **Clear and discoverable project governance documentation.**

Governance is publicly documented, defining TSC roles, maintainers, contributor responsibilities, and decision-making processes. See [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) 

- [ ] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Some governance updates are partially documented; discussions tracked in [TOC issue #1287](https://github.com/cncf/toc/issues/1287)

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

Vendor-neutral practices are implemented for hosting, communication, and architecture decisions. See [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) 

- [ ] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

Decision-making is partially documented in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) 

- [ ] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

Some documentation exists in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) regarding assignment and onboarding. 

- [ ] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

Maintainer lifecycle partially documented in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) 

- [ ] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Maintainer changes are tracked in [MAINTAINERS.md](https://github.com/openfga/community/blob/main/MAINTAINERS.md) 

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

No subprojects currently documented.

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

OpenFGA has 22 active maintainers with clearly defined roles. See [MAINTAINERS.md](https://github.com/openfga/community/blob/main/MAINTAINERS.md) 

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

Maintainer count covers repositories, CLI, SDKs, and documentation. 

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

Ownership in repositories aligns with roles defined in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) 

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

Maintainers and contributors agree to the CNCF CoC. 

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

Link: [CODE_OF_CONDUCT.md](https://github.com/openfga/.github/blob/main/CODE_OF_CONDUCT.md) 

- [ ] **All subprojects, if any, are listed.**

No subprojects currently exist.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

OpenFGA defines multiple contributor roles in its governance documentation, including maintainers, committers, and reviewers, with responsibilities outlined in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md#roles-and-responsibilities).

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

The project has a publicly documented contribution process in [CONTRIBUTING.md](https://github.com/openfga/openfga/blob/main/CONTRIBUTING.md) that explains how to submit issues, PRs, and participate in discussions.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

OpenFGA maintains public communication channels for announcements, discussion, and support, including a blog at [openfga.dev/blog](https://openfga.dev/blog) and Slack workspace [#openfga](https://cloud-native.slack.com/archives/C06G1NNH47N).

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.). List any non-public communications channels and what their special purpose is.**

The community documentation lists all public communication channels at [openfga.dev/docs/community](https://openfga.dev/docs/community), including Slack channels, GitHub discussions, and blog updates.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

OpenFGA hosts regular monthly community meetings with schedules documented at [openfga.dev/docs/community#monthly-community-meetings](https://openfga.dev/docs/community#monthly-community-meetings).

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

Contribution guidelines are described in detail in [CONTRIBUTING.md](https://github.com/openfga/.github/blob/main/CONTRIBUTING.md), including steps for new contributors, reviewers, and maintainers.

- [x] **Demonstrate contributor activity and recruitment.**

OpenFGA had 109 committers in the last 12 months as reported on [CNCF DevStats](https://all.devstats.cncf.io/d/53/projects-health-table?orgId=1), indicating active recruitment and participation from the community.

## Engineering Principles

### Suggested

- [x] **Roadmap change process is documented.**
  - OpenFGA follows a structured [RFC (Request for Comments) process](https://github.com/openfga/rfcs) for substantial changes to the project or roadmap. This ensures that contributors can review, comment, and approve significant changes, providing transparency and community input.

- [x] **History of regular, quality releases.**
  - OpenFGA maintains a track record of high-quality releases, with detailed metrics available through DevStats:
    - [Contribution activity](https://openfga.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now-2d): average 1450 contributions per month
    - [Issues opened/closed](https://openfga.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All): 46.5 issues opened and 28 closed per month
    - [PRs opened/closed](https://openfga.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All): 159 opened and 144 closed per month
    - [Open PR age](https://openfga.devstats.cncf.io/d/19/open-pr-age-by-repository-group?orgId=1&var-period=d7&var-repogroup_name=All): median 5 hours, 1 min over 7-day moving average

### Required 

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outline how this project fulfills an outstanding need and/or solves a problem differently.**
  - OpenFGA addresses a critical gap in application authorization by providing high-performance, relationship-based access control inspired by Google Zanzibar. It helps developers implement fine-grained access control, centralizing authorization logic to reduce errors, increase security, and scale reliably. Security risks like [Broken Object Level Authorization](https://owasp.org/API-Security/editions/2023/en/0xa1-broken-object-level-authorization/) are mitigated using OpenFGA’s approach.
  - Related CNCF projects can be found at [related-projects.md](https://github.com/openfga/community/blob/main/related-projects.md).

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**
  - OpenFGA is a flexible authorization solution that evaluates object-user relationships for access decisions. It supports modeling Role-Based, Attribute-Based, and Relationship-Based Access Control systems.
    - [Google Zanzibar paper](https://research.google/pubs/zanzibar-googles-consistent-global-authorization-system/) inspired modeling
    - SDKs for multiple languages allow easy integration into cloud-native applications.
    - Scales efficiently to billions of relationships, suitable for startups and large enterprises alike.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**
  - The [OpenFGA Roadmap](https://github.com/orgs/openfga/projects/1) tracks planned features and strategic priorities.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**
  - Detailed architecture and component design are documented in [Product Architecture](https://github.com/openfga/openfga/blob/main/docs/architecture/architecture.md), demonstrating scalability and cloud-native deployment patterns.
  - **Additional information from Project Reviews Subproject and General Technical Review:**  
    - The project team was engaged in a technical review on **August 8**, using OpenFGA as an example for the Project Reviews Subproject. 
    - The project team was noted to be flexible and easy to work with during this process, facilitating smooth evaluation.
    - The General Technical Review covered the full **cloud native software lifecycle**:  
      - **Day 0 – Planning Phase:** Roadmap process, target personas (developers, architects, platform teams), primary use cases (authorization engine), and differentiation from other solutions like OPA.  
      - **Day 1 – Installation and Deployment Phase:** Helm chart deployment, enabling/disabling in live clusters, configuration management, rollout, upgrade and rollback strategies including immutability and shadow checks.  
      - **Day 2 – Day-to-Day Operations Phase:** Scalability, reliability, observability (logs, metrics, dashboards), SLOs/SLIs, dependency management, and troubleshooting known failure modes.  
    - Security principles followed: secure by design and secure by default, support for multiple AuthN strategies, incremental TLS enablement, immutable authorization models, and structured access control.  
    - Links to supporting assessments and documentation:  
      - [Self Assessment](https://github.com/cncf/toc/blob/main/projects/openfga/security-assessment/self-assessment.md)  
      - [Joint Assessment](https://github.com/cncf/toc/blob/main/projects/openfga/security-assessment/joint-assessment.md)  
      - [Cloud Native Threat Model](https://github.com/cncf/toc/blob/main/projects/openfga/security-assessment/joint-assessment.md#threat-model)  
    - Meeting notes and additional review documentation are available in the [Google Doc](https://docs.google.com/document/d/1BxXhJbfhHc0WxF-EXAMPLELINK/edit) for reference.

- [x] **Document the project's release process.**
  - OpenFGA’s [release process](https://github.com/openfga/openfga/blob/main/RELEASES.md) defines versioning, QA procedures, and release steps to ensure reliable delivery of updates.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**
  - OpenFGA maintains a clear vulnerability reporting process through its [security policy](https://github.com/openfga/community/security/policy), specifying how to report vulnerabilities, triage them, and the expected response timelines.  
  - Security reports are handled with a structured workflow that includes acknowledgment, triage, mitigation, and disclosure stages.

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**
  - All members of the OpenFGA organization are required to enable [two-factor authentication](https://docs.github.com/en/authentication/securing-your-account-with-two-factor-authentication-2fa) for GitHub, ensuring secure access control to the repositories.  
  - Repository permissions are strictly managed, with role-based access control and separation of duties for code and documentation contributions.

- [x] **Document assignment of security response roles and how reports are handled.**
  - Security response roles are defined in [GOVERNANCE.md](https://github.com/openfga/.github/blob/main/GOVERNANCE.md) and the [security policy](https://github.com/openfga/community/security/policy), detailing how reports are handled, triaged, and escalated.  
  - OpenFGA rotates security team responsibilities to ensure diverse coverage and minimize single points of failure.

- [x] **Document Security Self-Assessment.**
  - OpenFGA completed a self-assessment documented at [Self-Assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/self-assessment.md) and a joint assessment at [Joint Assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/openfga/joint-assessment.md), reviewing security practices, threat modeling, and compliance with CNCF recommendations.  
  - The assessments validate adherence to cloud-native security principles, including secure defaults, configurable authentication/authorization strategies, and robust dependency management.  
  - OpenFGA’s immutable authorization models, “deny by default” approach, and incremental TLS enablement reduce risk and support secure deployments at scale.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**
  - OpenFGA holds the [OpenSSF Best Practices Badge](https://www.bestpractices.dev/en/projects/6374), demonstrating adherence to key security practices including vulnerability management, secure development, supply chain security, and community awareness.  
  - Continuous scanning of dependencies, CVE tracking, pinned GitHub Actions workflows, and secure project management settings reinforce adherence to OpenSSF best practices.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**
  - OpenFGA maintains a public list of adopters at [ADOPTERS.md](https://github.com/openfga/community/blob/main/ADOPTERS.md), which includes organizations using OpenFGA in development, trial, and production environments. Many adopters prefer not to disclose publicly.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters (these are not required to be in the publicly documented list of adopters)**
  - Highlighted production use cases:
    1. **Canonical**: OpenFGA embedded in [LXD](https://discourse.ubuntu.com/t/lxd-5-20-has-been-released/40865#authorization-restructure-7), [MicroCloud](https://www.gamingdeputy.com/canonical-unveils-microcloud-a-toolkit-for-rapid-cluster-deployment/), and [JAAS](https://github.com/canonical/openfga-operator) for multi-layer authorization.
    2. **Stacklok**: Migrated Minder platform from Open Policy Agent to OpenFGA for relationship-based access control ([Stacklok Blog](https://stacklok.com/blog/using-openfga-to-build-a-relationship-based-authorization-model-in-minder)).
    3. **Configu**: Chose OpenFGA for configuration authorization due to CNCF backing and battle-tested standards ([Configu Blog](https://configu.com/blog/authorization-over-configurations-using-openfga/)).
    4. **Docker**: Utilizes OpenFGA for authorization within Docker Hub.

- [x] **TOC verification of adopters.**
  - Adopters were provided to TOC for verification of production, trial, and development usage as part of the incubation review. Refer to the Adoption section and [ADOPTERS.md](https://github.com/openfga/community/blob/main/ADOPTERS.md).

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**
  - OpenFGA integrates with multiple CNCF and open-source projects:
    - [OpenTelemetry](https://github.com/openfga/js-sdk/blob/main/docs/opentelemetry.md) for tracing and telemetry
    - [Kubernetes / Helm](https://github.com/openfga/helm-charts) for deployment
    - [Grafana](https://github.com/openfga/openfga/blob/main/telemetry/grafana/dashboards/openfga.json) dashboards for monitoring
    - [Prometheus](https://github.com/openfga/openfga/blob/main/telemetry/prometheus/prometheus.yml) metrics collection
    - [ArtifactHUB](https://artifacthub.io/packages/helm/openfga/openfga) for Helm chart distribution
      
#### Adoption

##### Adopter 1 - Read AI / Meeting Intelligence
_August 2025_

Read AI, a leading meeting intelligence platform serving over 100,000 organizations, adopted OpenFGA in production to achieve scalable, high-performance, and secure authorization. They were drawn to OpenFGA for its Zanzibar-inspired model, detailed documentation, responsive maintainers, and self-hosting capabilities. Since April 2023, Read AI’s deployment handles over 5,200 requests/sec with low latency, managing billions of tuples reliably. Adoption of OpenFGA has improved internal design practices, reduced infrastructure costs, and provided confidence in secure data authorization. The project’s strong technical foundation and active community have enabled smooth integration, though the adopter notes that broader maintainer diversity could further enhance governance and sustainability.

Reference: [OpenFGA Adopter Interview - Read AI](/projects/openfga/openfga-adopter-interview-readai.md)

##### Adopter 2 - Agicap / Fintech  
_September 2025_ 

Agicap, a European fintech serving over 8,000 customers, adopted OpenFGA in April 2023 to replace an internally built RBAC solution with a scalable ReBAC-based authorization system. Today, OpenFGA underpins all authorization across Agicap’s SaaS platform, running in both production and pre-production environments at ~250 RPS. The project’s flexibility, strong documentation, and responsive maintainers were key factors in adoption. Agicap actively contributes bug fixes and performance improvements upstream, participates in community meetings, and has co-presented with maintainers at KubeCon. Key strengths cited include OpenFGA’s robust ReBAC model, reliability at production scale, and an engaged community, while areas for improvement include maintainer diversity and support for idempotent batch writes.

Reference: [OpenFGA Adopter Interview - Agicap](/projects/openfga/openfga-adopter-interview-agicap.md)

##### Adopter 3 - Docker / Cloud Infrastructure & Developer Tools
_September 2025_

Docker adopted OpenFGA in early 2024 to replace its internal access control system with a flexible, ReBAC-based solution. The project is used in production across multiple services, handling 100–150 RPS, and has demonstrated measurable value by centralizing permission management, reducing manual overhead, and accelerating integration of new products. Key drivers for adoption included self-hosting, open-source availability, mature SDKs, and an active, responsive community. Docker highlighted the project’s robust performance, clear documentation, and strong maintainers as core strengths, while noting potential improvements in DSL tooling, batch evaluation optimizations, and broader marketing to increase awareness and adoption.

Reference: [OpenFGA Adopter Interview - Docker](/projects/openfga/openfga-adopter-interview-docker.md)

##### Adopter 4 - Grafana / Cloud Observability and Monitoring
_September 2025_

Grafana Labs has adopted OpenFGA as the foundation of its multi-tenant authorization system, moving away from a legacy single-tenant engine. They emphasize OpenFGA’s strengths in combining storage and evaluation, supporting per-tenant schemas, and enabling scalable ReBAC-based architectures. The project stands out for its production readiness, comprehensive documentation, and ease of schema iteration, which reduced engineering burden and accelerated product development. From a community perspective, Grafana values OpenFGA’s CNCF governance alignment, responsive maintainers, and open, proactive engagement. They also note its strong potential to become a standard in authorization due to its design for scale and backing from adopters. These qualities make OpenFGA well-positioned for CNCF Incubation.

Reference: [OpenFGA Adopter Interview - Grafana](/projects/openfga/openfga-adopter-interview-grafana.md)

##### Adopter 4 - Zuplo / API Management
_September 2025_

Zuplo uses OpenFGA to implement scalable relationship-based access control (ReBAC) across its API management platform. They highlight production readiness, PostgreSQL support for edge authorization, a single authorization model, and responsive maintainers as the project’s key strengths. OpenFGA has streamlined their authorization processes, reduced manual effort, and is viewed as well-positioned for CNCF Incubation.

Reference: [OpenFGA Adopter Interview - Zuplo](/projects/openfga/openfga-adopter-interview-zuplo.md)
