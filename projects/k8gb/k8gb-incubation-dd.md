# k8gb Incubation Due Diligence

**Status:** Adoption evaluation complete (3 interviews, 3 orgs, 3 geographies). Ready public comment → TOC vote.

- **Incubation application:** [cncf/toc#1472](https://github.com/cncf/toc/issues/1472)
- **Project:** Sandbox since 2021-03-30 → applying **Incubation** · [k8gb-io/k8gb](https://github.com/k8gb-io/k8gb) · [k8gb.io](https://www.k8gb.io/)
- **Primary DD authoring:** @TheFoxAtWork and @ricardorocha
- **GTR:** [cncf/toc#2021](https://github.com/cncf/toc/issues/2021) · [PR #2029](https://github.com/cncf/toc/pull/2029) (in flight)
- **Governance review:** [cncf/toc#1963](https://github.com/cncf/toc/issues/1963) · [PR #2020](https://github.com/cncf/toc/pull/2020)
- **Security SSA on `main`:** [self-assessment.md](https://github.com/cncf/toc/blob/main/projects/k8gb/security-assessment/self-assessment.md)
- **OpenSSF / CLOMonitor:** [clomonitor.io/projects/cncf/k8gb](https://clomonitor.io/projects/cncf/k8gb) — verify passing badge before public PR
- **Adopter verification (interviews):** @angellk and @kgamanji, 3 completed
- **CNCF policies:** [https://www.cncf.io/policies/](https://www.cncf.io/policies/) · [https://github.com/cncf/foundation/tree/main/policies-guidance](https://github.com/cncf/foundation/tree/main/policies-guidance)

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1472)

## Incubation Evaluation Summary for K8GB

### Criteria Evaluation

Emily Fox and Ricardo Rocha conducted the due diligence of K8GB who applied for incubation. Karena Angell and Katie Gamanji have completed the adopter interviews. The project has completed the criteria that show its maturity at Incubation.

The following criteria implementations are noteworthy to call out:

The project was very responsive to requests for changes.

The following actions were provided to the project that were considered blocking but since resolved:
* Completion of vendor-neutrality for API group
    * API group change:  [k8gb-io/k8gb#2180](https://github.com/k8gb-io/k8gb/issues/2180)
    * Implementation: [k8gb-io/k8gb#2203](https://github.com/k8gb-io/k8gb/pull/2203).

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:
* It is recommended the project define activity requirements for maintainers to ensure access to the codebase is tightly controlled and reduce the possibility of a compromise occurring through an inactive maintainer role.
* It is recommended k8gb's README, governance, and contributing files be updated to provide information on when and where community meetings are held.  Be sure to include the date and time with the link to attend.
* It is recommended that the project document their overall planning and roadmap update process to better guide and inform adopters and contributors of how to introduce changes (features, fixes, enhancements) to the project. This can be incorporated into the Governance document for easy discoverability.
* It is recommended that the project, as part of their planning and roadmap update documentation, include information on decision making (in GOVERNANCE.md), contribution acceptance (in CONTRIBUTING.md), and general process for introducing changes to the governance and project scope/goals (in GOVERNANCE.md).

### Adoption Evaluation

Three adopter interviews were conducted across three organizations in three geographies spanning financial services and managed services. All three run k8gb in production, have contributed upstream, and report stable, low-maintenance operation. The project fills a genuine gap — DNS-based Kubernetes-native global server load balancing — with no equivalent open-source alternative identified by any adopter.

**Combined production scale:** 220+ clusters, 1,700+ nodes, 23,000+ pods across 3 continents in regulated industries.

**Consistent findings across all adopters:**
- **No alternative exists.** All three independently evaluated proprietary alternatives (F5, AviNetworks, Azure Traffic Manager, HAProxy, Nginx) and found none offering native Kubernetes integration with DNS-based GSLB at their required scale and cloud-agnosticism.
- **Architecture is praised.** Described as "elegant," "does one thing and does it very good," with a simple, small codebase that is easy to understand and maintain.
- **All adopters are contributors.** Upstream contributions include Windows DNS integration, Istio Virtual Services support, core features, and documentation. One adopter's contributor became a maintainer through this process.
- **Operationally stable.** Zero rollbacks reported. Strong trust in non-breaking-changes principle. CVE response within 2 weeks. Quarterly releases with ad hoc as needed.
- **Self-sustaining.** All three know the codebase well enough to self-maintain. All stated they would fork and maintain the project if it were archived, with a preference for formal upstream maintainership.

**Areas for improvement (non-blocking):**
- Onboarding documentation for users coming from L7/HTTP load balancing backgrounds (DNS-based GSLB is unfamiliar)
- Signed container images not yet available
- DNS service limits at scale required optimization (zone delegation) — resolved

**Must-fix (resolved):** API group renamed from `k8gb.absa.oss/v1beta1` to `k8gb.io/v1beta1` ([#2180](https://github.com/k8gb-io/k8gb/issues/2180), closed Mar 2026).

The adopter interviews reflect a project in active, production use at significant scale, appropriate for the Incubation level. The project demonstrates healthy community engagement, contributor-driven adoption, and a clear, well-defined scope with no viable open-source competitor.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

## Application Process Principles

### Suggested

N/A

### Required

- [X] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**
  - The Project completed a General Technical Review on 30-Jan-2026 (Kashif Khan, TAG Infrastructure), rated Satisfactory. It can be discovered at [projects/k8gb/tech-review/2026-01-30.md](https://github.com/cncf/toc/blob/main/projects/k8gb/tech-review/2026-01-30.md). Issue: [cncf/toc#2021](https://github.com/cncf/toc/issues/2021).

- [X] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The Project is in the process of correcting the only remaining item that may be perceived as non-vendor-neutral (API group) and has made significant progress such that this is considered non-blocking for Incubation.
API group change:  [k8gb-io/k8gb#2180](https://github.com/k8gb-io/k8gb/issues/2180)
Initial Implementation: [k8gb-io/k8gb#2203](https://github.com/k8gb-io/k8gb/pull/2203)

- [X] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**

Met during Project's application on March 30, 2021 at the Sandbox maturity level. Refer to the project's Onboarding PR [PROJECT ONBOARDING] k8gb [sandbox#251](https://github.com/cncf/sandbox/issues/251) for completion of this.

- [ ] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [X] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The project has a plethora of documentation that spans multiple use types, from quick start to tutorials and presentations.

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

* The Project underwent a Governance Review in January of 2026. The results may be found in projects/k8gb/governance-review/2026-01-21.md
* The details and notes from the Governance review are reflected below.

### Suggested

- [X] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The project has embarked on a recent effort to formalize maintainer onboarding and offboarding while preparing for Incubation. Previously, the project adhered to a lighter-weight process and therefore the CODEOWNERS file history does not reflect the new documented election and governance since those changes were made. [An upcoming election will exercise these new changes](https://github.com/k8gb-io/k8gb/issues/2210).

- [X] **Clear and discoverable project governance documentation.**

The Project's governance is readily discoverable and clearly laid out in the [GOVERNANCE.md](https://github.com/k8gb-io/k8gb/blob/master/GOVERNANCE.md) file.

- [X] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

The Project has recently made changes to their governance to reflect their current project activities, elections, and approvals.

- [X] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

The project's governance details [Vendor-neutrality as a key principle](https://github.com/k8gb-io/k8gb/blob/5a38dd692dd4f7a9adcdc5715355750c0393b7cf/GOVERNANCE.md?plain=1#L13).

- [ ] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

The project does not have this documented in their Governance or contributor file.

It is recommended that the project, as part of their planning and roadmap update documentation, include information on decision making (in GOVERNANCE.md), contribution acceptance (in CONTRIBUTING.md), and general process for introducing changes to the governance and project scope/goals (in GOVERNANCE.md).

- [X] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

The project only defines the Maintainer and Project Lead roles within their Governance. Each of these describe how they are filled and removed.

- [X] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

The project does not formally have a complete maintainer lifecycle process documented. A very lightweight lifecycle could be inferred from its current Governance.

- [X] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Examples of maintainers being onboarded and offboarded in the following commits under the previous process:

- https://github.com/k8gb-io/k8gb/commit/1ecee28519f3a70ad4fe24ef5f47a6c93f2734d2
- https://github.com/k8gb-io/k8gb/commit/9ab4528bfc103bbd4fbf580510491fd3386aed90
- https://github.com/k8gb-io/k8gb/commit/c5ea88b4be8ee20a9698717f2f6204af86a45c2e
- https://github.com/k8gb-io/k8gb/commit/5b35c35f5bfc643ffd180b82bfde6067b2fcaa79

- [X] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

k8gb does not contain sub-projects at this time.

### Required

- [X] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

The Project's CODEOWNERS file contains the names, GitHub handles, and affiliation of each. The project does not have additional roles beyond Project Lead, which is assigned in the CODEOWNERS file.

- [X] **A number of active maintainers which is appropriate to the size and scope of the project.**

The project lists a total of 7 maintainers across 5 organizations. Some maintainers are substantially more active than others, with some contributions only occurring once a year by a given maintainer. It should be noted that the majority of contributions come from two organizations only.

It is recommended the project define activity requirements for maintainers to ensure access to the codebase is tightly controlled and reduce the possibility of a compromise occurring through an inactive maintainer role.

- [X] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

The Code and Doc ownership is governed by the [CODEOWNERS](https://github.com/k8gb-io/k8gb/blob/master/CODEOWNERS) file.

- [X] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

The Project's [CoC file](https://github.com/k8gb-io/k8gb/blob/master/CODE_OF_CONDUCT.md) states it adheres to the CNCF CoC and provides a link directly to that file.

- [X] **CNCF Code of Conduct is cross-linked from other governance documents.**

The CNCF CoC is cross-linked from the governance file via the project's CoC.

- [X] **All subprojects, if any, are listed.**

k8gb has no sub-projects at this time.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

The project does not have a contributor ladder defined yet.

### Required

- [X] **Clearly defined and discoverable process to submit issues or changes.**

The project adheres to traditional GitHub workflows and is documented in their [CONTRIBUTING.md](https://github.com/k8gb-io/k8gb/blob/master/CONTRIBUTING.md) file.

- [X] **Project must have, and document, at least one public communications channel for users and/or contributors.**

The Project has a CNCF Slack Channel, [#k8gb](https://cloud-native.slack.com/archives/C021P656HGB) that is public for users and contributors.

- [X] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

The project has several public communications channels:
- CNCF Slack: [#k8gb](https://cloud-native.slack.com/archives/C021P656HGB)
- LinkedIn: <https://www.linkedin.com/company/k8gb/>
- X: <https://x.com/k8gb_io>
- Community meetings: <https://zoom-lfx.platform.linuxfoundation.org/meetings/k8gb>
- Web site & blog: <http://k8gb.io/>

The project also has a security reporting alias that is made up of maintainers and not open to the public due to the sensitivity of communications.

- [X] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

The project holds meetings regularly and they may be found at
<https://zoom-lfx.platform.linuxfoundation.org/meetings/k8gb>.

While non-blocking, k8gb's README, governance, and contributing files do not provide information on when and where community meetings are held. The README mentions they happen and are uploaded to YouTube, but no date or time and it doesn't have the link for people to find out how to attend. This is a simple fix, and is especially important for contributors who may have more to discuss than a simple issue will permit.

- [X] **Documentation of how to contribute, with increasing detail as the project matures.**

The project's
[CONTRIBUTING.md](https://github.com/k8gb-io/k8gb/commits/master/CONTRIBUTING.md)
has been continuously iterated on as the project grows in maturity.

- [X] **Demonstrate contributor activity and recruitment.**

The project's [LFX page](https://insights.linuxfoundation.org/project/k8gb/contributors) provides some information regarding contributor activity. Most contributors are identified as maintainers. The project's slack channel and LinkedIn activity shows efforts for recruitment and contribution to potential new users.

## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

The project does not currently document their roadmap planning process to capture changes and updates.

It is recommended that the project document their overall planning and roadmap update process to better guide and inform adopters and contributors of how to introduce changes (features, fixes, enhancements) to the project. This can be incorporated into the Governance document for easy discoverability.

- [X] **History of regular, quality releases.**

The Project has a [well documented release process](https://github.com/k8gb-io/k8gb/blob/master/CONTRIBUTING.md#release-process) that contributes to regular, quality releases. While the current releases do not have a fixed cadence (roughly quarterly), they do occur with regularity (once a month or thereabout). Reviewing closed issues, there are few bugs filed that correspond with a given release.

### Required

- [X] **Document project goals and objectives that illustrate the project's differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  -  _If applicable_ a General Technical Review was completed on 30-JAN-2026, and can be discovered in the [projects/k8gb](https://github.com/cncf/toc/tree/main/projects/k8gb/tech-review/2026-01-30.md) folder.

A domain technical review was satisfactorily completed and reviewed as part of the project's due diligence.

The DTR found that k8gb provides a cloud native, vendor-neutral Global Server Load Balancer (GSLB). Unlike traditional GSLB solutions that require expensive hardware or a central management cluster, k8gb uses a decentralized, DNS-based control plane. This removes single points of failure and allows for a lightweight, Kubernetes-native operational model.

- [X] **Document what the project does, and why it does it - including viable cloud native use cases.**

The DTR noted k8gb automates multi-cluster failover and geo-aware traffic routing. Primary use cases include disaster recovery, high availability, and multi-cloud resilience by dynamically updating DNS records based on the real-time health of Kubernetes Ingress and Service resources.

- [X] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The [project's roadmap](https://github.com/orgs/k8gb-io/projects/2/views/2) can be found as a project board located in their repository.

- [X] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**
  -  _If applicable_ a Technical Review was completed/updated on 30-Jan-2026, and can be discovered in the [projects/k8gb](https://github.com/cncf/toc/tree/main/projects/k8gb/tech-review/2026-01-30.md) folder.

The DTR found the k8gb architecture is decentralized; each participating cluster runs the k8gb operator and an embedded CoreDNS instance. It leverages external-dns to handle zone delegation. This architecture ensures that even if one cluster is completely lost, the remaining clusters continue to serve valid DNS and route traffic correctly.

- [X] **Document the project's release process.**

k8gb follows a [mature release](https://github.com/k8gb-io/k8gb/blob/master/CONTRIBUTING.md#release-process) cadence (roughly quarterly) using Semantic Versioning (SemVer). The pipeline is fully automated via GitHub Actions, encompassing container image signing, Helm chart publishing, and SLSA provenance generation.

## Security

### Suggested

N/A

### Required

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance.

- [X] **Clearly defined and discoverable process to report security issues.**

The project has a discoverable [SECURITY.md](https://github.com/k8gb-io/k8gb/commit/eb15fe8dd6a2a7734b8ec1d11d231b987dc0e458) file that outlines a clear process for reporting vulnerabilities to the maintainers and specifies the expected response timeline.

- [X] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

Maintainers enforce strict branch protection on GitHub. All code changes require at least one (often two) maintainer approvals. 2FA is required for all organization members, and the project uses fine-grained GitHub Action permissions. Additional information regarding the project's security practices may be found in their [self-assessment.md](https://github.com/cncf/toc/blob/main/projects/k8gb/security-assessment/self-assessment.md#development-pipeline).

- [X] **Document assignment of security response roles and how reports are handled.**

The project does not have a dedicated security response team. Responsibility for security triage is assigned to the core maintainers. Procedures for handling reports and coordinating fixes are documented within their governance and security policies.

- [X] **Document Security Self-Assessment.**

The project has completed and had their [Security Self-Assessment](https://github.com/cncf/toc/blob/main/projects/k8gb/security-assessment/self-assessment.md) merged on December 4th 2025.

- [X] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The project has achieved [passing badge](https://www.bestpractices.dev/en/projects/4866/passing) for the OpenSSF's best practices badging program.

## Ecosystem

### Suggested

N/A

### Required

- [X] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

The project maintains a list of Adopters in their [ADOPTERS.md](https://github.com/k8gb-io/k8gb/blob/master/ADOPTERS.md) file of the repo with a pinned issue to encourage adopters to add themselves to the file.

- [X] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The TOC conducted 3 adopter interviews across 3 independent organizations spanning financial services and managed services in 3 geographies. All 3 run k8gb in production at significant scale (220+ clusters combined). All 3 are active upstream contributors. Details in the Adoption Evaluation section above.

- [X] **TOC verification of adopters.**

3 interviews completed and verified. Refer to the Adoption Evaluation section of this document.

- [X] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

k8gb is deeply integrated into the CNCF landscape, specifically:

* CoreDNS: Uses it as the primary DNS engine.
* external-dns: For automating cloud DNS provider updates.
* Prometheus/Grafana: For observability and automated health-check metrics.
* Gateway API: Native support for modern Kubernetes networking standards.
* Cloud Providers: Automated integration with AWS Route53, Azure DNS, and GCP Cloud DNS.

#### Adoption

##### Adopter 1 — Financial Services

Production use for ~7 years (project originator). 200 clusters (~1,500 nodes, 1,200 on-prem), 200 services, ~23,000 pods. After a 6-7 month evaluation of proprietary alternatives, selected k8gb for native Kubernetes integration, vendor independence, and self-service at scale. Active upstream contributor; runs upstream binaries. Now positioning k8gb as non-optional infrastructure in next-generation platform. Reports lowest payment-system incident rate among peers over past 5 years, attributed partly to k8gb's failover capability.

##### Adopter 2 — Financial Services

Production use across ~20 clusters in 8-10 regions (multi-cloud + on-prem). Active-active pattern with automatic failover. Active contributor — contributed DNS integration support upstream. Self-supporting; knows codebase well. Architecture described as "elegant." CVE response within 2 weeks. 70% reduction in DNS-related incident response times. Only adopts CNCF-governed open source.

##### Adopter 3 — Managed Services

Production use for ~2 years. 100-200 services across multi-cluster Kubernetes with service mesh. Adopted k8gb because cloud provider DNS solutions were unsuitable for private-space multi-cluster deployments. Contributor's colleague became a maintainer through upstream contribution. Zero rollbacks. Described project as "does one thing and does it very good." Willing and able to self-maintain if needed.
