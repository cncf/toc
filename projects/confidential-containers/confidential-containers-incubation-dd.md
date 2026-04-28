# Confidential Containers Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1504)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Confidential Containers

### Criteria Evaluation

_$TOCMEMBER conducted the due diligence of Confidential Containers who applied for $LEVEL. The project [has/has not] completed the criteria that show its maturity at $LEVEL. The following criteria implementations are noteworthy to call out... $NOTABLES. The following actions were provided to the project that were considered blocking but since resolved... $BLOCKERS. The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project... $RECOMMENDATIONS._

### Adoption Evaluation

_The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY._

### Final Assessment

_[The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS]._

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**

  <!-- (TOC Evaluation goes here) -->
  - This was completed and occurred on 28-Aug-2024, and can be discovered at <https://zoom.us/rec/share/nCSdjZxN4DjxA7RopMqocddvDBPzxuO-Xpa2u-9xzB8MbumatD1Vtw0ePpNrobHy.Cby9t12K1IJfq_tk>

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

  <!-- (TOC Evaluation goes here) -->  
  The CoCo project uses vendor neutral resources for project host, communication etc.
  Some examples are:
  - CoCo Github: <https://github.com/confidential-containers>
  - CoCo Project Website: <https://confidentialcontainers.org/>
  - CoCo Slack: <https://cloud-native.slack.com/archives/C039JSH0807>

- [x] **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		

  <!-- (TOC Evaluation goes here) -->
  - The project contacts and TOC Reviewers had a kick-off meeting on Jan. 16th, set expectations and discussed general steps & timelines.

- [ ] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [ ] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

  <!-- (TOC Evaluation goes here) -->
  The project provides approriate documents for installation and configuration, e.g.: <https://confidentialcontainers.org/docs/getting-started/>

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

  <!-- (TOC Evaluation goes here) -->
  The CoCo project has been continuously updating governance doc, some examples are:
  - Update governance doc to include rules of removing inactive maintainers <https://github.com/confidential-containers/confidential-containers/pull/235>
  - Add provisions to the governance document for members who move from one company to another or who become inactive or leave the project. <https://github.com/confidential-containers/confidential-containers/pull/329>


- [x] **Clear and discoverable project governance documentation.**

  <!-- (TOC Evaluation goes here) -->
  The [project governance doc](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md) is maintained in the main repository.

  **Suggestion by Kevin:** Since CoCo has 10+ active non-fork repos, TOC reviewers suggest to consider creating a community repository and maintain governance and community relavent docs there.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

  <!-- (TOC Evaluation goes here) -->
  The project maintains an active and up-to-date governance framework that accurately reflects the current project state. The documentation is regularly updated to capture Steering Committee leadership transitions, organizational representation changes, and refinements to maintainer lifecycle processes.

  Some examples are:
  - removed AMD rep Ryan Savino from SC and added to emeritus <https://github.com/confidential-containers/confidential-containers/pull/348>
  - Intel maintainers update <https://github.com/confidential-containers/confidential-containers/pull/339>
  - Update Microsoft representative maintainer <https://github.com/confidential-containers/confidential-containers/pull/338>

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  <!-- (TOC Evaluation goes here) -->
  As outlined in [GOVERNANCE.md](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md), the CoCo project effectively operationalizes vendor neutrality through structural mechanisms, specifically the two-seat limit per organization on the Steering Committee.

  **Suggestion by Kevin:** It is noted that the documentation currently lacks an explicit definition of 'vendor neutrality' as a core principle. The TOC reviewers recommend explicitly codifying a Vendor Neutrality clause to align with CNCF best practices before the project advances to graduation.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  <!-- (TOC Evaluation goes here) -->
  [GOVERNANCE.md#decision-making](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md#decision-making) explicitly documents a consensus-driven framework. It establishes a clear voting protocol for critical decisions, specifically leadership changes and governance modifications, requiring a defined supermajority threshold (2/3rds of current SC members) when consensus is not achieved.

- [ ] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [ ] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **CNCF Code of Conduct is cross-linked from other governance documents.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **All subprojects, if any, are listed.**

<!-- (TOC Evaluation goes here) --> 

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [ ] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) --> 

## Engineering Principles

### Suggested

- [x] **Roadmap change process is documented.**

  <!-- (TOC Evaluation goes here) -->
  The project adopts a [use-case driven development approach](https://confidentialcontainers.org/blog/2024/02/16/introduction-to-confidential-containers-coco/#foundational-principles), focusing community efforts on key scenarios rather than feature-centric models. The roadmap and scope are determined via consensus in the [Steering Committee and weekly community meetings](https://docs.google.com/document/d/1E3GLCzNgrcigUlgWAZYlgqNTdVwiMwCRTJ0QnJhLZGA/edit), with changes documented in the public [roadmap.md](https://github.com/confidential-containers/confidential-containers/blob/main/roadmap.md).

- [x] **History of regular, quality releases.**

  <!-- (TOC Evaluation goes here) -->
  The project documents historical releases with change notes at: <https://github.com/confidential-containers/confidential-containers/tree/main/releases>

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  -  The general Technical Review was updated on 24-Feb-2026 (currently in Draft status), and can be discovered at <https://github.com/cncf/toc/pull/2051>.

  <!-- (TOC Evaluation goes here) -->
  The project's goals and differentiation are documented through its [website](https://confidentialcontainers.org): _"...enables cloud native data in use protection by leveraging hardware Trusted Execution Environments (TEEs)."_ The project's documentation and draft GTR also describe how it solves this problem differently by encapsulating unmodified Kubernetes pods inside confidential VMs, completely shielding workloads from host operating systems, cluster administrators, and cloud providers.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

  <!-- (TOC Evaluation goes here) -->
  The project clearly documents its core function, motivation, and practical applications:

  - What the project does: <https://confidentialcontainers.org/docs/overview/#what-is-the-confidential-containers-project>
  - Why it does it: <https://confidentialcontainers.org/blog/2024/02/16/introduction-to-confidential-containers-coco/#motivation>
  - Viable use cases include the following: (Ref: <https://confidentialcontainers.org/docs/use-cases/>)
    - Confidential AI
    - Switchboard Oracles: Securing Web3 Data with Confidential Containers
    - Secure Supply Chain (Trusted Pipeline)

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

  <!-- (TOC Evaluation goes here) -->
  According to the [roadmap.md](https://github.com/confidential-containers/confidential-containers/blob/main/roadmap.md) the project maintains its：
  - short-term roadmap at the GitHub boards (including the [Confidential containers github board](https://github.com/orgs/confidential-containers/projects/6) and [Trustee github board](https://github.com/orgs/confidential-containers/projects/10)), and
  - its mid/long-term roadmap at the project's website based on use-case driven development. Ref: <https://confidentialcontainers.org/docs/use-cases/>

  **TODO for maintainers**: the short-term roadmap Confidential containers github board link is unavailable with outdated "view", needs to be fixed.

- [ ] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**
  -  _If applicable_ a general Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

    <!-- (TOC Evaluation goes here) -->
    The project maintains a detailed architectural overview on its [website](https://confidentialcontainers.org/docs/architecture/design-overview/), demonstrating a robust and viable cloud-native design. It effectively outlines key mechanisms such as Pod-Centric Virtualization, host deprivileging, and its sophisticated remote attestation (Trustee) architecture.

- [ ] **Document the project's release process.**

<!-- (TOC Evaluation goes here) --> 

## Security

### Suggested

N/A

### Required

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance.

- [x] **Clearly defined and discoverable process to report security issues.**

  <!-- (TOC Evaluation goes here) -->
  The project has a published security policy via the GitHub org-wide security policy:
  - Security policy (`SECURITY.md`): <https://github.com/confidential-containers/.github/blob/main/SECURITY.md>

  The policy instructs reporters to use GitHub’s “Report a vulnerability” (private reporting) mechanism rather than filing public issues.

- [ ] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  <!-- (TOC Evaluation goes here) -->
  **Maintainers’ input (from application):** The incubation application references a maintainers list as evidence that maintainers use 2FA:
  - Maintainers list: <https://github.com/confidential-containers/confidential-containers/blob/main/MAINTAINERS>

  **TOC reviewer assessment / gap (from confidential-containers/confidential-containers#349):**
  - A maintainer list does **not** by itself provide auditable evidence that 2FA is **required and enforced** for GitHub org members/maintainers.
  - Reviewers need a **clear, stable, public** reference describing:
    - how access is granted/revoked,
    - what controls prevent unauthorized changes (e.g., CODEOWNERS / required reviews / branch protection expectations),
    - whether 2FA is required/enforced for privileged roles.

  **Action / follow-up requested:** Project to document the access control model and explicitly document 2FA requirements and/or org enforcement:
  - <https://github.com/confidential-containers/confidential-containers/issues/349>

- [ ] **Document assignment of security response roles and how reports are handled.**

  <!-- (TOC Evaluation goes here) -->
  **Maintainers’ input (from application):**
  - The project points to the org-level `SECURITY.md` as the place describing handling of reports:
    - <https://github.com/confidential-containers/.github/blob/main/SECURITY.md>

  **TOC reviewer assessment / gap (from confidential-containers/confidential-containers#350):**
  - While `SECURITY.md` describes the reporting mechanism, it historically referenced “maintainers and security champions” without a clearly defined, discoverable role model elsewhere (e.g., in governance/community docs).
  - Reviewers need more explicit documentation of:
    - who the security responders are (org-wide vs per subproject),
    - how membership/ownership is determined and maintained (onboarding/offboarding/continuity),
    - high-level responsibilities (triage/coordination/communication), without duplicating the mechanics already in `SECURITY.md`.
  - Also recommended: cross-link security reporting guidance from contributing documentation (and other relevant guides), so users can easily find the security reporting process.

  **Action / follow-up requested:**
  - <https://github.com/confidential-containers/confidential-containers/issues/350>

- [x] **Document Security Self-Assessment.**

  <!-- (TOC Evaluation goes here) -->
  The project has published a Security Self-Assessment:
  - <https://tag-security.cncf.io/community/assessments/projects/confidential-containers/self-assessment/>

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

  <!-- (TOC Evaluation goes here) -->
  The project holds an OpenSSF Best Practices passing badge:
  - <https://www.bestpractices.dev/projects/5719>

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

The list of adopters can be found in the project repo: https://github.com/confidential-containers/confidential-containers/blob/main/ADOPTERS.md which includes the usage level. Most adopters are at the beta level.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, dev/test for incubation. The TOC was able to verify confirm at least dev/test use with all interviewed adopters.

- [x] **TOC verification of adopters.**

The CoC maintainers provided the TOC with a list of 5-6 adopters who agreed to be interviewed for the Incubation Due Diligence process. 4 of these adopters were interviewed. The adoption portion of this document contains interview summaries from adopters who approved public attribution. All adopters are verified using CoC at the level appropriate for incubation.

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

This is documented in the [alignment documentation](https://github.com/confidential-containers/confidential-containers/blob/main/alignment.md) in the project repo. This document contains the list of CNCF projects that are directly related to CoCo, list of non CNCF projects that are directly related to CoCo along with other potential related projects.

#### Adoption

##### Adopter 1 - IBM/Technology

In a March 2, 2026 interview, Nicolas Mäding explained that IBM has been contributing to and promoting Confidential Containers since 2021, including running internal trials. IBM considers it the most mature option within the Cloud Native Computing Foundation ecosystem, largely due to strong industry participation. Their work emphasizes s390x, IBM Secure Execution, and Trustee.

Production adoption depends on two key factors: support for bare metal and a standardized Trustee model. Both are still evolving following a Q4 2025 tech preview that IBM plans to harden. The team currently runs version 0.16.0, contributes upstream, and intends to stay aligned with future releases since the project underpins their Kubernetes product.

A key tension lies in roadmap priorities—particularly the emphasis on x86 and public cloud versus on-premises environments and IBM Z/LinuxONE systems. IBM advocates for a more balanced approach across architectures and deployment models.

Documentation is considered solid, especially architecture diagrams, though improved threat modeling would be beneficial as the project matures. Maintainers are highly engaged via Slack, GitHub, and Zoom. IBM sees the project’s primary value as enabling a unified hybrid-cloud approach across diverse hardware. Archiving the project would have a significant negative impact.

Improvement area: greater openness to non-x86 architectures and alternative deployment models.

Refer to the full [interview report](IBM-interview.md) for more details.

##### Adopter 2 - AccuKnox/Technology

In a February 10, 2026 interview, AccuKnox (represented by Rahul Jadhav, interviewed by Lin Sun) described using Confidential Containers (CoCo) in 5G security—specifically for intent-based security in 5G ORAN and Core environments.

A key use case is protecting the 5G UDR (Unified Data Repository), which stores highly sensitive data such as subscription keys. Because this data often resides in public cloud environments—where cloud service provider root access must be assumed as a threat—CoCo provides a mechanism to secure these workloads using hardware-backed isolation.

AccuKnox views CoCo as uniquely capable of delivering a full Trusted Execution Environment (TEE) pathway for containerized workloads. Features such as cloud API adapters have been particularly valuable, especially after early integration challenges with AWS.

Their current usage remains pre-production which matches the expected level for incubation. They are running version 0.13.0 (approximately one year behind) and have not kept related components like 5g-blueprint-controls fully updated. While documentation was helpful, they relied heavily on Slack support from community members such as Pradeepta.

They presented their 5GSEC work to the community and were encouraged to list themselves as adopters. AccuKnox frames CoCo’s value in terms of improving security posture in untrusted cloud environments, where TEE-based solutions are often the only viable option.

Gaps identified: Limited hardware availability, Interface instability and High cost. They recommend clearer messaging and better articulation of real-world use cases, rather than focusing solely on deep technical detail. Their enterprise future plan includes integrating CoCo into AccuKnox Enterprise, particularly for government and federal deployments.

Refer to the full [interview report](AccuKnox-interview.md) for more details.

##### Adopter 3 - NVIDIA/Technology

In a January 23, 2026 interview, Dan Middleton from NVIDIA described roughly two years of involvement with Confidential Containers as maintainers, ecosystem contributors, and adopters. He explained confidential computing as leveraging CPU and GPU features to create protected memory regions—similar to VM-level isolation—shielded from other tenants. CoCo enables this capability in a Kubernetes-native way.

NVIDIA’s motivations fall into three categories:
- Enabling GPU support so AI workloads can leverage confidential computing
- Building products on top of CoCo (e.g., integration with Red Hat OpenShift)
- Meeting internal and external datacenter security requirements with a cloud-native confidential stack

NVIDIA views CoCo as the leading cloud-native confidential computing project due to its momentum, feature set, and broad ecosystem support across cloud providers, hardware vendors, and ISVs. Current status is pre-production which matches the expected level for incubation. They expect early access external offerings, with GA targeted for late Q1 2026. They are using version: 0.18, and tracking releases closely (downstream users follow OpenShift cadence).

Adoption was relatively smooth due to internal expertise. Initial integration challenges—mainly around compatibility with other CNCF components like container runtimes—have largely been resolved. Documentation has improved significantly, especially with the introduction of a unified docs site.

NVIDIA actively engages through community meetings, GitHub, and Slack. Key strengths of the project is strong security-first design, broad and diverse maintainer base and faster and less complex integration compared to alternatives (e.g., kubelet-in-confidential-VM).

Key improvement areas are more end-user feedback and visibility, easier usability and simplified per-tenant security configuration.

Refer to the full [interview report](NVIDIA-interview.md) for more details.

##### Adopter 4 - TDC/Technology

In a January 29, 2026 interview, TDC (Denmark)—Nino Wael, Lasse, and Martin Rasmussen, interviewed by Lin Sun—describes a small IT-security-focused group that knew CoCo earlier but seriously evaluated it from early 2025. The main pull is running AI workloads securely, then broader confidential workloads. They previously looked at Gramine (Intel SGX-centric) and moved toward CoCo for broader hardware options (including Intel TDX) and multi-vendor contributors—seen as a major advantage. The current stage is pre-production / POC only, so no measurable business value yet; subjectively, CoCo is much easier than Gramine. They try to stay on the latest release (CoCo 0.18) and follow upgrade paths.

Pain points is that basic tutorials exist, but combining features often breaks without extra config, pushing them to read source. Docs sometimes 404, go stale, or describe removed code; some endpoints lack documentation for their small integration layer. Slack is helpful but insufficient alone. They treat community engagement as mandatory given complexity and a steep learning curve.
Scope critique: they feel the project sometimes sprawls into areas better handled elsewhere (e.g. Keycloak for authz), including non-core items like a Trustee admin web UI—they’d prefer deeper CNCF integration over building everything in-tree.

Strengths of the project is its hardware flexibility, strong vendor participation, approachable contribution, and a single pane to confidential containers versus stitching per-vendor docs. Improvements areas are a compatibility matrix, clearer marking of unsupported paths (e.g. Azure CSI wrapper), and cleaner, truthful documentation.

Refer to the full [interview report](TDC-DK-interview.md) for more details.
