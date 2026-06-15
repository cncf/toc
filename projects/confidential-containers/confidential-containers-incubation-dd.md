# Confidential Containers Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1504)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Confidential Containers

### Criteria Evaluation

Kevin Wang, Faseela K, and Lin Sun conducted the due diligence of Confidential Containers, which applied for Incubation. This effort was also supported by a [General Technical Review from Matt Young](https://github.com/cncf/toc/pull/2051) and a [Governance Review from Josh Gavant](https://github.com/cncf/toc/pull/2081). The project has completed the criteria that show its maturity at Incubation.

#### Noteworthy Implementations

- **Vendor neutrality enforced on the Steering Committee**: The Steering Committee caps each organization at two seats and currently spans seven companies (Alibaba, IBM, Intel, AMD, Red Hat, NVIDIA, Microsoft), with a process to keep membership tracking the project's major contributors.
- **Governance iterated over time**: The governance document has changed as the project has run, with examples including inactive-maintainer removal rules and provisions for members who change employer or leave ([#235](https://github.com/confidential-containers/confidential-containers/pull/235), [#329](https://github.com/confidential-containers/confidential-containers/pull/329), [#348](https://github.com/confidential-containers/confidential-containers/pull/348), [#339](https://github.com/confidential-containers/confidential-containers/pull/339), [#338](https://github.com/confidential-containers/confidential-containers/pull/338)).
- **Alignment with its main dependency, Kata Containers**: Kata Containers contributors sit on the CoCo Steering Committee, and the 6-week release cadence follows the Kata lifecycle.
- **Subproject-to-repository mapping**: Each subproject lives in its own repository, and the mapping of components to subprojects is documented on the project [website](https://confidentialcontainers.org/docs/architecture/design-overview/#components).
- **Secure-by-default architecture**: The design denies the host access to container images (pulled inside the guest enclave via `image-rs`), releases secrets only after attestation through Trustee, and applies deny-by-default OPA policies on the Kata Agent API to block host interference.
- **Signed build provenance at SLSA Build Level 2**: The project generates signed `in-toto` provenance via GitHub Actions for components such as `kata-containers`, `guest-components`, and `cloud-api-adaptor`.
- **Consistent, well-documented release process**: The project releases on a 6-week cadence and runs each release against a standardized checklist.

#### Blockers (Previously Raised, Now Resolved)

- **Access Control and 2FA Enforcement**: The project was asked to document and enforce 2FA for privileged roles and make the access-control model publicly auditable. This was resolved in [confidential-containers/confidential-containers#351](https://github.com/confidential-containers/confidential-containers/pull/351), which updated `governance.md` to explicitly require 2FA for org members, maintainers, Steering Committee members, and Security Managers, and to describe how maintainer access is granted and revoked via GitHub teams.
- **Security Response Role Documentation**: The project was asked to formally define security response roles and document how vulnerability reports are handled. This was resolved through [confidential-containers/confidential-containers#351](https://github.com/confidential-containers/confidential-containers/pull/351) and [confidential-containers/.github#25](https://github.com/confidential-containers/.github/pull/25), which introduced the Security Manager role in `governance.md`, aligned `SECURITY.md` with the governance language, and added a link from `CONTRIBUTING.md` to the security reporting guidance.

#### Recommended Enhancements (Non-Blocking)

- **Publish a public maintainer and team-membership list**: Maintainer and functional-role assignments are managed through private GitHub teams referenced in `CODEOWNERS`, so they are not publicly readable. The TOC reviewers recommend publishing this membership through declarative configuration using a tool such as [cilium/team-manager](https://github.com/cilium/team-manager) or [CLOWarden](https://github.com/cncf/clowarden).
- **State vendor neutrality explicitly in governance**: Vendor neutrality is enforced through the two-seat-per-organization Steering Committee limit, but the governance documents do not state it as a principle. The reviewers recommend adding a vendor-neutrality clause before graduation.
- **Move governance and community docs into a community repository**: With 10+ active non-fork repositories, the reviewers suggest a dedicated community repository to hold governance and community documentation.
- **Document subproject governance details**: The subproject removal process, per-subproject maturity status, and a public per-subproject maintainer list are not yet documented and should be completed before graduation.
- **Integrate with the official CNCF calendar**: Weekly community meetings are documented in a public Google Doc; adding them to the official CNCF calendar would make them easier to find.
- **Keep the contributing guide current**: The contributing guide has not been updated since 2024 and should be reviewed periodically to match the current state of the project.
- **Fix the short-term roadmap board link**: The short-term roadmap link to the Confidential Containers GitHub board in `roadmap.md` is broken (outdated "view") and should be corrected.
- **Recruit more adopters and case studies**: Adopter verification was met through interviews; the project should keep recruiting adopters and encouraging public case studies.

### Adoption Evaluation

The TOC interviewed four adopters of Confidential Containers — IBM, NVIDIA, AccuKnox, and TDC — spanning hardware vendors, cloud and security vendors, telecom, and enterprise IT across multiple geographies. All four reported pre-production or dev/test usage on recent releases (versions 0.13 through 0.18), which matches the level of adoption expected for Incubation. Several adopters also contribute upstream or hold maintainer roles, and most track the project closely as the foundation for downstream products.

Adopters consistently described Confidential Containers as the primary cloud-native project offering a full Trusted Execution Environment pathway for containerized workloads. The most commonly cited strengths were its multi-vendor maintainer base and hardware flexibility — covering Intel TDX, IBM Secure Execution, and NVIDIA GPUs — which drew adopters away from single-vendor alternatives such as Intel SGX-based Gramine. Adopters also valued the security-first design and the active maintainer engagement through Slack, GitHub, and the weekly community meeting.

Reported value centered on securing sensitive workloads in untrusted environments, a unified hybrid-cloud approach across diverse hardware, and faster integration compared with building a confidential stack in-house. Because the interviewees are still pre-production, few could quantify business value yet, though several noted that archiving the project would have a significant negative impact on their products.

Common areas for improvement were documentation (stale or broken links, a missing threat model, and the lack of a compatibility matrix), ease of configuration, clearer use-case messaging over deep technical detail, and concern that the project's scope is expanding into areas already served by other CNCF projects. Overall, Confidential Containers is in active use at a level appropriate for Incubation, with a diverse and engaged adopter base and a clear path to deeper production adoption as hardware availability and documentation mature.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

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

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

  <!-- (TOC Evaluation goes here) -->
  The project provides appropriate documents for installation and configuration, e.g.: <https://confidentialcontainers.org/docs/getting-started/>

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

  Since CoCo has 10+ active non-fork repos, TOC reviewers suggest to consider creating a community repository and maintain governance and community relavent docs there.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

  <!-- (TOC Evaluation goes here) -->
  The governance documentation is kept up to date with the current project state, including Steering Committee leadership transitions, organizational representation changes, and changes to maintainer lifecycle processes.

  Some examples are:
  - removed AMD rep Ryan Savino from SC and added to emeritus <https://github.com/confidential-containers/confidential-containers/pull/348>
  - Intel maintainers update <https://github.com/confidential-containers/confidential-containers/pull/339>
  - Update Microsoft representative maintainer <https://github.com/confidential-containers/confidential-containers/pull/338>

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  <!-- (TOC Evaluation goes here) -->
  As outlined in [GOVERNANCE.md](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md), the CoCo project effectively operationalizes vendor neutrality through structural mechanisms, specifically the two-seat limit per organization on the Steering Committee.

  The TOC reviewers note that the governance documents do not state 'vendor neutrality' as an explicit principle, and recommend adding a vendor-neutrality clause before the project advances to graduation.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  <!-- (TOC Evaluation goes here) -->
  [GOVERNANCE.md#decision-making](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md#decision-making) explicitly documents a consensus-driven framework. It establishes a clear voting protocol for critical decisions, specifically leadership changes and governance modifications, requiring a defined supermajority threshold (2/3rds of current SC members) when consensus is not achieved.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

  <!-- (TOC Evaluation goes here) -->
  The project's [GOVERNANCE.md](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md#becoming-a-project-maintainer) defines how maintainer roles are assigned via GitHub teams. Functional roles like "security managers" are also documented. However, the TOC reviewers noted that since these assignments are managed via private GitHub teams, a public list of these roles is not currently available.

  The TOC reviewers suggest the project either maintain a public list manually or use a tool such as [team-manager](https://github.com/cilium/team-manager) to improve openness and transparency.

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

  <!-- (TOC Evaluation goes here) -->
  The [GOVERNANCE.md](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md) documents the complete lifecycle, including onboarding (building trust/contributions) and removal processes for both Maintainers and the Steering Committee.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

  <!-- (TOC Evaluation goes here) -->
  The project demonstrates maintainer lifecycle outcomes through recorded updates in various sub-projects, such as [Trustee](https://github.com/confidential-containers/trustee/issues?q=is%3Aissue++in%3Atitle+maintainer) and [guest-components](https://github.com/confidential-containers/guest-components/issues?q=is%3Aissue++in%3Atitle+maintainer).

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

  <!-- (TOC Evaluation goes here) -->
  Subproject leadership and contribution follow the organization-level [governance](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md).

  The TOC reviewers note that the following subproject governance details are not yet documented and recommend completing them before graduation:
  - The subproject removal process.
  - The maturity status of individual subprojects.
  - A public per-subproject maintainer list.

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

  <!-- (TOC Evaluation goes here) -->
  The project documents maintainers list with all the relevant information at <https://github.com/confidential-containers/confidential-containers/blob/main/MAINTAINERS>.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

  <!-- (TOC Evaluation goes here) -->
  Based on [LFX Insights](https://insights.linuxfoundation.org/project/confcont/contributors), the project has a broad group of active contributors from multiple organizations, which is appropriate for its scale.

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

  <!-- (TOC Evaluation goes here) -->
  The project uses `CODEOWNERS` files across its repositories to enforce ownership in alignment with the documented GitHub team-based governance.

  While `CODEOWNERS` files reference GitHub teams, the membership of those teams is only visible to members of the organization and not publicly accessible. The TOC reviewers suggest considering a tool that manages org and team membership declaratively via public configuration, such as [cilium/team-manager](https://github.com/cilium/team-manager) ([config example](https://github.com/cilium/community/tree/main/ladder)), [CLOWarden](https://github.com/cncf/clowarden) ([config example](https://github.com/cncf/people)), or a similar tool, to improve transparency and auditability of project membership.

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

  <!-- (TOC Evaluation goes here) -->
  The project has adopted the CNCF Code of Conduct, as documented in its [CODE_OF_CONDUCT.md](https://github.com/confidential-containers/confidential-containers/blob/main/CODE_OF_CONDUCT.md).

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

  <!-- (TOC Evaluation goes here) -->
  The Code of Conduct is discoverable in the `.github` repository and linked from the project's metadata.

- [x] **All subprojects, if any, are listed.**

  <!-- (TOC Evaluation goes here) -->
  The project lists its subprojects and components on its [website](https://confidentialcontainers.org/docs/architecture/design-overview/#components), including Trustee, guest-components, cloud-api-adaptor, operator, trustee-operator, and td-shim.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

  <!-- (TOC Evaluation goes here) -->
  The contributor ladder and roles are defined in the [governance document](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md#community-members-and-roles).

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

  <!-- (TOC Evaluation goes here) -->
  The process for submitting changes is clearly defined in the [CONTRIBUTING.md](https://github.com/confidential-containers/.github/blob/main/CONTRIBUTING.md) and on the project's [website](https://confidentialcontainers.org/docs/contributing/#making-contributions).

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

  <!-- (TOC Evaluation goes here) -->
  The project uses the `#confidential-containers` channel on CNCF Slack as its primary public communication channel.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

  <!-- (TOC Evaluation goes here) -->
  Communication channels, including Slack and the weekly community meeting, are documented in the [contributing guide](https://confidentialcontainers.org/docs/contributing/#connecting-with-the-community).

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

  <!-- (TOC Evaluation goes here) -->
  Weekly community meetings are held and documented in a [public Google Doc](https://docs.google.com/document/d/1E3GLCzNgrcigUlgWAZYlgqNTdVwiMwCRTJ0QnJhLZGA/).

  The TOC reviewers recommend integrating the meetings with the official CNCF calendar to improve discoverability.

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

  <!-- (TOC Evaluation goes here) -->
  A [contributing guide](https://confidentialcontainers.org/docs/contributing/) is maintained. The TOC reviewers noted it should be reviewed periodically to match the current state of the project.

- [x] **Demonstrate contributor activity and recruitment.**

  <!-- (TOC Evaluation goes here) -->
  Contributor activity is actively tracked and demonstrated via [LFX Insights](https://insights.linuxfoundation.org/project/confcont) and [CNCF DevStats](https://confidentialcontainers.devstats.cncf.io/). 

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
  - The General Technical Review was updated on 24-Feb-2026 (currently under review), and can be discovered at <https://github.com/cncf/toc/pull/2051>.

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
  According to the [roadmap.md](https://github.com/confidential-containers/confidential-containers/blob/main/roadmap.md), the project maintains its:
  - short-term roadmap at the GitHub boards (including the [Confidential containers github board](https://github.com/orgs/confidential-containers/projects/6) and [Trustee github board](https://github.com/orgs/confidential-containers/projects/10)), and
  - its mid/long-term roadmap at the project's website based on use-case driven development. Ref: <https://confidentialcontainers.org/docs/use-cases/>

  The short-term roadmap Confidential containers github board link is unavailable with an outdated "view", needs to be fixed. Ref: <https://github.com/confidential-containers/confidential-containers/blob/15cbeeff8282e6cea15f66900f66b5936ab0971e/roadmap.md#L10>

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**
  - The General Technical Review was updated on 24-Feb-2026 (currently under review), and can be discovered at <https://github.com/cncf/toc/pull/2051>.

    <!-- (TOC Evaluation goes here) -->
    The project maintains a detailed architectural overview on its [website](https://confidentialcontainers.org/docs/architecture/design-overview/), covering key mechanisms such as Pod-Centric Virtualization, host deprivileging, and the remote attestation (Trustee) architecture.

- [x] **Document the project's release process.**

  <!-- (TOC Evaluation goes here) -->
  The project maintains a well-documented and consistent release process.
  
  It follows a [6-week release cadence](https://github.com/confidential-containers/confidential-containers/blob/main/README.md) aligned with the Kata Containers lifecycle.
  
  The community utilizes a standardized release checklist ([.github/ISSUE_TEMPLATE/release-check-list.md](https://github.com/confidential-containers/confidential-containers/blob/main/.github/ISSUE_TEMPLATE/release-check-list.md)) to ensure quality and consistency across releases.

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

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

  <!-- (TOC Evaluation goes here) -->
  This was initially found blocking but fixed by [confidential-containers/confidential-containers#351](https://github.com/confidential-containers/confidential-containers/pull/351).

  According to the updated [governance.md](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md), the project now explicitly requires 2FA for GitHub org members, maintainers, Steering Committee members, and Security Managers. It also documents that maintainer access is managed through GitHub teams referenced in each repository's `CODEOWNERS` file, while clarifying that GitHub org membership is not a formal governance tier.

  Issue [#349](https://github.com/confidential-containers/confidential-containers/issues/349) was then closed as completed on Mar 3, 2026.

- [x] **Document assignment of security response roles and how reports are handled.**

  <!-- (TOC Evaluation goes here) -->
  This was initially found blocking but fixed by [confidential-containers/confidential-containers#351](https://github.com/confidential-containers/confidential-containers/pull/351) and [confidential-containers/.github#25](https://github.com/confidential-containers/.github/pull/25).

  According to the updated [governance.md](https://github.com/confidential-containers/confidential-containers/blob/main/governance.md), the project now defines the Security Manager role directly, including who holds that role, how additional Security Managers are added or removed, and how they may coordinate with affected external parties on a need-to-know basis during pending advisories.

  In parallel, [confidential-containers/.github#25](https://github.com/confidential-containers/.github/pull/25) aligned `SECURITY.md` with the governance language, replaced the older "security champions" terminology, and added a link from `CONTRIBUTING.md` to the security reporting guidance so the reporting path is easier to find.

  Taken together, these changes address the documentation gap raised in issue [#350](https://github.com/confidential-containers/confidential-containers/issues/350).

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
