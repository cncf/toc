# Kyverno Graduation Due Diligence

- Link to [Graduation application issue](https://github.com/cncf/toc/issues/1967)
- Link to [previous Graduation application issue](https://github.com/cncf/toc/pull/1217)


<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Graduation Evaluation Summary for Kyverno

### Criteria Evaluation

Karena Angell, TOC Member and Chair, conducted the due diligence of Kyverno who applied for Graduation. The due diligence is supplemented by a Governance Review conducted by Josh Gavant, TOC Subproject Project Reviews member, and a General Technical Review conducted by Dylan Page, CNCF TAG Infrastructure Co-Chair. 

The project has completed the criteria that show its maturity at Graduation. The following criteria implementations are noteworthy to call out:
* The ability to write policies in YAML or CEL lowers the learning curve for adopters and allows platforms teams to use existing GitOps and Helm workflows
* The Kyverno CLI allows teams to 'shift-left' by being able to validate policy changes in the CI/CD pipelines prior to reaching production. Recent optimizations have adopters reporting highly improved execution time in large clusters
* [Documented AI Usage Policy](https://github.com/kyverno/community/blob/main/AI_USAGE_POLICY.md)
* Highly responsive and collaborative maintainers that continue to incorporate community feedback into new releases and are quick to respond to critical bug fixes
* Addressed the previous maintainer ladder concerns and continue to evaluate organizational and maintainer diversity
* Granular policy enforcement to help manage specific security deviations in complex environments
* Extensive documentation allowing easy and fast implementation
* Operational stability at scale with reliability required by large adopting organizations
* Useful subprojects such as [Kyverno Chainsaw](https://github.com/kyverno/chainsaw) that enables declarative end-to-end testing in a Kubernetes cluster

The following actions were provided to the project that were considered blocking but since resolved:
* Build a robust maintainer ladder and update governance - and show in action for at least 6 months - to address organizational diversity concerns. See Governance Section.
* [Update CODEOWNERS file to reflect accurate maintainer list](https://github.com/kyverno/kyverno/issues/15110)
* [Vendor Neutrality: Remove references to the contributing organization](https://github.com/kyverno/kyverno/issues/15113)
* [Link Release Procedure to main Release page](https://github.com/kyverno/website/issues/1867)
* [Update the ROADMAP file to include the Roadmap change process](https://github.com/kyverno/kyverno/issues/15478)

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

* [TAG Security & Compliance - Kyverno Security Assessment](https://github.com/kyverno/kyverno/issues/15335)
  * fix blog post links in https://kyverno.io/docs/guides/security/#security-audits #15336
  * Explicitly state risks associated with external data lookups and recommend using authentication and encryption
  * Add checks to restrict cache lookups within namespaced policies ([[Feature] Restrict Global Context access and implement cache bounds #15359](https://github.com/kyverno/kyverno/issues/15359))
  * Improve architecture diagram to separate logical and physical components
  * Review all [published advisories](https://github.com/kyverno/kyverno/security/advisories?state=published) and update threat model
  * Add [security advisories](https://github.com/kyverno/kyverno/security/advisories?state=published) link to the Kyverno security page.
  * Add SAST like semgrep or Nancy
  * Add cache bounds to the GlobalContext ([[Feature] Restrict Global Context access and implement cache bounds #15359](https://github.com/kyverno/kyverno/issues/15359))
  * Authenticate requests from the API server [[docs](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#authenticate-apiservers)]
  * Add threat model for CLI and other deployment options
  * Improve docs on the risk of not having a network policy to restrict Kyverno communications
  * CLI command to generate a network policy for kyverno
  * Sample network policies for common platforms

* [TOC Project Reviews Subproject - Kyverno General Technical Review Assessment](https://github.com/kyverno/kyverno/issues/15473)
  * Confirm the project’s explicit “non-goals” list (if any) beyond subproject statements.
  * Request links to formal UX/user research reports, surveys, or adopter interviews (if maintained publicly or shareable).
  * Capture a concise list of “core CRDs/APIs” and stability policy per API group (if there is a single authoritative page).
  * Request a consolidated threat model (or pointer to one) if maintained outside the self-assessment.
  * Link the definitive doc section describing webhook cert issuance/rotation and operational guidance (if separate from general docs).
  * Clarify recommended “temporary disable” patterns for production incidents (for example, scale to zero vs. change failure behavior vs. selective policy disable).
  * Capture formal SLOs/SLIs (if the project has published targets) and map them to existing metrics.
  * Document recommended alerting thresholds (latency/error/drops) and provide reference dashboards/runbooks.
  * Document dependency lifecycle/SCA workflow end-to-end (tooling, tracking, SLAs for vulnerability remediation) beyond SBOM publication.
  * Summarize “safe-mode” recommendations for production incidents (preferred steps, expected blast radius) from maintainer guidance.
  * Provide explicit process for third-party notices in source and in distributed artifacts (and how it aligns with CNCF attribution guidance).
  * Link the current security response process (reporting, triage, rotation) and evidence of diverse representation in the response team.

### Adoption Evaluation

Adopter feedback for Project Kyverno satisfies the TOC's requirements for moving levels within the CNCF to Graduated status. Adopters highlight Kubernetes-native simplicity as the project's primary strength, specifically the use of YAML and CEL for policy authoring. Users across the board value its ability to go beyond simple validation to provide automation through mutation and generation, as well as its high-performance admission control that maintains site stability under heavy loads. The community and maintainers are consistently recognized for being highly responsive to bug fixes and feature requests.

However, several areas for improvement were identified to support long-term enterprise stability:

* **Operational Transparency & Testing:** Adopters are requesting more robust internal regression and scale testing to prevent production-impacting updates, along with detailed architectural diagrams to clarify internal logic.
* **Documentation & User Journey:** Adopters would like documentation tailored to different user personas and more guidance on complex, real-world use cases, such as AI/ML specialized environments.
* **Tooling & Metrics:** Adopters would also like to see the project unify fragmented sub-projects into a cohesive user experience and providing more granular, real-time observability metrics to better monitor performance at scale.
* **Architectural Evolution:** Future interest includes shifting toward a daemonset architecture, expanding into vulnerability scanning, and further diversifying project governance to ensure long-term stability.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation. The TOC requests the project meet with the TOC in 3 months time to review the non-blocking issues and findings from the TAG Security and Compliance Joint Security Assessment and General Technical Review Assessment. At that time the TOC would like to see the roadmap for any remaining non-blocking concerns from both assessments.

## Criteria

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**
  - This was completed and occurred on 05-NOV-2025, and can be discovered at the [TAG's meeting recording](https://zoom.us/rec/play/Vs086CY6n9shLtswCRCZ8um01rRbSGRRBDBwrDXNt8mGjfVbW-ZExVNgknug7rrlFYtsVkiIvnmhdX-4.83cpFEK7O_eTSNw3?eagerLoadZvaPages=sidemenu.billing.plan_management&accessLevel=meeting&canPlayFromShare=true&from=share_recording_detail&continueMode=true&componentName=rec-play&originRequestUrl=https%3A%2F%2Fzoom.us%2Frec%2Fshare%2FR38A1sMm5PpzxJGhy3oY27JMWmmOJ1CSe3mG4AYjZurs3GQHRWcItW8PNaSELzxN.A19PcH1eMKZTYOCl&autoplay=true&startTime=1762365863000) beginning at 34 minutes.


- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The project has addressed all outstanding issues regarding vendor neutrality.


- [x] **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io/) projects and requirements for moving forward through the CNCF Maturity levels.**		
   - [x] N/A - Met during Project's application on DD-MMM-YYYY.


- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

Additional documentation provided by the project increases the likelihood of end user adoption:
* Kyverno provides comprehensive user documentation (installation guides, configuration, API reference), policy samples, and tutorials on https://kyverno.io.
* The [Releases page](https://kyverno.io/docs/releases/) explains versioning, branch strategies, and release cadence, while the Roadmap file links to release tracker boards and design proposals.
* Sample policies are hosted in a [dedicated repository](https://github.com/kyverno/policies/) and rendered on the [website](https://kyverno.io/policies/).

## Governance and Maintainers

Note: This report is augmented by a [Governance Review](https://github.com/cncf/toc/blob/main/projects/kyverno/governance-review/2025-12-16.md) that was performed by Josh Gavant in late 2025.

### Suggested

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

  The Kyverno project has made multiple governance revisions; notably in 2024 they added a [Community Repo](https://github.com/kyverno/community) to serve as a focal point for project governance.
  
  For example, they [added roles in July 2024](https://github.com/kyverno/community/pull/7), and further expanded it to a [contributor ladder in May 2025](https://github.com/kyverno/community/pull/48).


### Required

- [x] **Clear and discoverable project governance documentation.**

  Kyverno's primary [GOVERNANCE.md document](https://github.com/kyverno/community/blob/main/GOVERNANCE.md) is published in the [kyverno/community](https://github.com/kyverno/community) repo, linked in the [kyverno/kyverno](https://github.com/kyverno/kyverno) repo and referenced on the website's [community page](https://kyverno.io/community/#project-governance).

  It describes all applicable governance processes: principles, contributor roles, a progression process for contributors to become maintainers, meetings, conflict resolution, and how changes are made.


- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

  Kyverno's governance is up to date, and has been updated following suggestions from the [previous governance review](https://github.com/cncf/toc/blob/main/projects/kyverno/governance-review/2025-12-16.md).
  Their meetings are held and notes taken as described on the [community page](https://kyverno.io/community/#meetings) and other documents.
  Elections of new contributors and maintainers are executed via GitVote as configured [here](https://github.com/kyverno/community/blob/main/.gitvote.yml).


- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

  Kyverno's [documentation](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#vendor-neutrality) directly references the [CNCF vendor neutrality guidelines](https://contribute.cncf.io/projects/best-practices/community/vendor-neutrality/).
  
  Additional restrictions on member composition for the steering committee were established as part of the previous [Governance Review](https://github.com/cncf/toc/blob/main/projects/kyverno/governance-review/2025-12-16.md).


- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

  [Decisions about the project, its leadership and governance](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#changes) are made by Kyverno _Maintainers_ who are elected to their positions by existing maintainers.

  The Steering Committee provides non-binding advice and recommendations to the Maintainers. The Maintainers retain operational authority.
  

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

  [Project Roles](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#project-roles) are outlined in Kyverno's primary [GOVERNANCE.md](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md).
  
  Maintainers, along with any subproject maintainers are listed in [MAINTAINERS.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md) in the [kyverno/community](https://github.com/kyverno/community/) repo.


- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

  Both onboarding and offboarding processes are outlined in the [Maintainer section of GOVERNANCE.md](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#maintainers), along with [general offboarding guidance](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#off-boarding-guidance) provided in the doc as well.
  
  Examples of both may be found in the [commit history of MAINTAINERS.md](https://github.com/kyverno/community/commits/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md)


- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

  Examples may be found in the [commit history of MAINTAINERS.md](https://github.com/kyverno/community/commits/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md)


- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

  [MAINTAINERS.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md) lists each maintainer, their GitHub ID and company affiliation. Sub‑project maintainers are similarly documented.


- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

  Kyverno has 15 maintainers total across it's [primary repo and subprojects](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md#kyverno-maintainers-1).


- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

  Kyverno has maintainers from 6 different organizations, along with 2 independent maintainers across their [primary repo and subprojects](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md#kyverno-maintainers-1).


- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

  Membership in private GitHub teams governs control of repositories corresponding to projects and sub‑projects, e.g., kyverno‑maintainers grants control of kyverno/kyverno, kyverno‑json‑maintainers of kyverno/kyverno-json, etc.

  Each project has its maintainers listed separately in the main [MAINTAINERS.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md). Documented roles correspond to actual code and documentation ownership through GitHub teams.

  Each repository also has a CODEOWNERS file; for example the [CODEOWNERS](https://github.com/kyverno/kyverno/blob/30b67627949a9cc6c2d4fdbf5e52783846a3689f/CODEOWNERS) for [kyverno/kyverno](https://github.com/kyverno/kyverno).


- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

  Kyverno's own [CODE_OF_CONDUCT.md](https://github.com/kyverno/community/blob/main/CODE_OF_CONDUCT.md) references the CNCF's code of conduct. Other repositories reference this primary code of conduct.


- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

  Kyverno's own [CODE_OF_CONDUCT.md](https://github.com/kyverno/community/blob/main/CODE_OF_CONDUCT.md) references the CNCF's code of conduct, and is cross-linked from other docs e.g. [GOVERNANCE.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/GOVERNANCE.md#code-of-conduct).


- [x] **All subprojects, if any, are listed.**

  Subprojects are listed in the Kyverno [documentation](https://kyverno.io/docs/subprojects/) as well as on the main [GitHub organization page](https://github.com/kyverno).


- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

  Each subproject has its maintainers listed separately in the main [MAINTAINERS.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md).
  
  The add/remove process is outlined in [GOVERNANCE.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/GOVERNANCE.md)
  

## Contributors and Community

Note: This report is augmented by a [Governance Review](https://github.com/cncf/toc/blob/main/projects/kyverno/governance-review/2025-12-16.md) that was performed by Josh Gavant in late 2025.


### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

  Kyverno's ladder is outlined in their [GOVERNANCE.md](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#project-roles).
  
  They have 4 roles with a defined progression:
  - [New Contributor](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#new-contributors)
  - [Contributor](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#contributors)
  - [Reviewer](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#reviewers)
  - [Maintainer](https://github.com/kyverno/community/blob/4994f36ca91d19e848ea51f66c832e5ba56b1fee/GOVERNANCE.md#maintainers)


### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

  Kyverno's general contributing docs are outlined in [CONTRIBUTING.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/CONTRIBUTING.md) within the [kyverno/community](https://github.com/kyverno/community) repo.
  
  This is augmented by subproject specific contributor guides in each other repo.
  

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

  Kyverno maintains public [Slack channels](https://kyverno.io/community/#slack-channel) and a [mailing list](https://groups.google.com/g/kyverno). 

  A weekly community meeting is [advertised](https://kyverno.io/community/#meetings), [documented](https://docs.google.com/document/d/1kFd4fpAoHS56mRHr73AZp9wknk1Ehy_hTB_KA7gJuy0/), and recorded (links to recordings are in meeting notes).

  A weekly maintainers meeting is also [advertised](https://kyverno.io/community/#maintainers-meeting), [documented](https://docs.google.com/document/d/1I_GWsz32gLw8sQyuu_Wv0-WQrtRLjn9FuX2KGNkvUY4/edit?usp=sharing), and recorded.

  These meetings are also published on [the CNCF calendar](https://www.cncf.io/calendar/).


- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

  The [community page](https://kyverno.io/community/) details Slack channels, mailing lists, GitHub Discussions, weekly meetings, and the adopters program. Sub‑projects use the same channels, and specific projects (e.g., Chainsaw) have dedicated Slack threads.

  A disclosure process and private email address for security issues is documented on the project's [security page](https://kyverno.io/docs/security/#disclosure-process).


- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

  The [community page](https://kyverno.io/community/) lists meeting times and links to agendas stored on Google Docs.

  Meetings are also published on [the CNCF calendar](https://calendar.cncf.io).
  

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

  Kyverno's general contributing docs are outlined in [CONTRIBUTING.md](https://github.com/kyverno/community/blob/cfa604580b15f3f971d6c7b5b4bc395a758940e2/CONTRIBUTING.md) within the [kyverno/community](https://github.com/kyverno/community) repo.
  
  This is augmented by subproject specific contributor guides in each other repo.
   

- [x] **Demonstrate contributor activity and recruitment.**

  New contributors are listed with each [release](https://github.com/kyverno/kyverno/releases). For example the [v1.17 release](https://github.com/kyverno/kyverno/releases/tag/v1.17.0) had 14 new contributors.

  Additional maintainer recruitment can be seen in the commit history of [MAINTAINERS.md](https://github.com/kyverno/community/commits/cfa604580b15f3f971d6c7b5b4bc395a758940e2/MAINTAINERS.md).


## Engineering Principles

Note: This report is augmented by a [General Technical Review](https://github.com/cncf/toc/blob/main/projects/kyverno/tech-review/2026-02-02.md%E2%80%8E) that was performed by Dylan Page in early 2026.

### Suggested

N/A

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  * A General Technical Review was completed on 02-FEB-2026, and can be discovered at [https://github.com/cncf/toc/blob/main/projects/kyverno/tech-review/2026-02-02.md%E2%80%8E](https://github.com/cncf/toc/blob/main/projects/kyverno/tech-review/2026-02-02.md%E2%80%8E).
 

- [x] **Document what the project does, and why it does it - including viable cloud native use cases. _This requirement may also be satisfied by completing a General Technical Review._**
  * A General Technical Review was completed on 02-FEB-2026, and can be discovered at [https://github.com/cncf/toc/blob/main/projects/kyverno/tech-review/2026-02-02.md%E2%80%8E](https://github.com/cncf/toc/blob/main/projects/kyverno/tech-review/2026-02-02.md%E2%80%8E).

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The project maintains a [public roadmap](https://github.com/kyverno/kyverno/blob/main/ROADMAP.md) that is kept up-to-date and includes:
* [Release Tracker board](https://github.com/orgs/kyverno/projects/9/views/1) 
* [Kyverno Design Proposals](https://github.com/kyverno/KDP) 
* [GitHub milestones](https://github.com/kyverno/kyverno/milestones?direction=asc&sort=due_date&state=open)

- [x] **Roadmap change process is documented.**

The project has documented the Roadmap change process in the ROADMAP.md.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.  *This requirement may also be satisfied by completing a General Technical Review and capturing the output in the project's documentation.***
  * A General Technical Review was completed on 02-FEB-2026, and can be discovered at [https://github.com/cncf/toc/blob/main/projects/kyverno/tech-review/2026-02-02.md%E2%80%8E](https://github.com/cncf/toc/blob/main/projects/kyverno/tech-review/2026-02-02.md%E2%80%8E).

The TOC has confirmed project Kyverno’s [documentation](https://kyverno.io/docs/introduction/how-kyverno-works/) describes the architecture (policy controller, admission webhooks, background scan controller) and provides diagrams explaining how it interacts with Kubernetes. The new policy types (ValidatingPolicy, ImageValidatingPolicy, MutatingPolicy, GeneratingPolicy, DeletingPolicy) are explained in [release blogs](https://kyverno.io/blog/releases/) and [documentation](https://kyverno.io/docs/policy-types/).

- [x] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:** 

  - [x] Release expectations (scheduled or based on feature implementation)
  - [x] Tagging as stable, unstable, and security related releases
  - [x] Information on branch and tag strategies
  - [x] Branch and platform support and length of support
  - [x] Artifacts included in the release.
  - Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.

[Blocker/Completed: Make the documented release process more visible](https://github.com/kyverno/website/issues/1867) 

The project has thoroughly documented the release process and procedures and has met the TOC requirements. 

- [x] **History of regular, quality releases.**

The Kyverno project has a [documented history](https://github.com/kyverno/kyverno/releases) of regular, quality releases and is quick to address bugs. The quality was further assessed through the adopter interviews. The TOC recommends the project provide further information regarding the impact each release makes so adopters can assess the urgency of release adoption.

## Security

Note: This report is augmented by a [Kyverno Security Joint Assessment](https://docs.google.com/document/d/1JJVlYu6LNpWs2ukirOmoHMEhCsuOEoseBq6dq-Kcgjc/edit?tab=t.0#heading=h.9i73ao98ff5y) that was driven by Andrew Martin with members of TAG Security and Compliance.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

The Kyverno project maintains a [passing badge](https://www.bestpractices.dev/en/projects/5327/passing) at this point in time.

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

The TOC confirms the Kyverno project maintains a [SECURITY.md](https://github.com/kyverno/community/blob/main/SECURITY.md) file and [security page](https://kyverno.io/docs/security/) describing how to report vulnerabilities via email to kyverno-security@googlegroups.com, what information to include, and expected response timelines.

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

The Kyverno project enforces branch protection and requires two‑factor authentication on all repositories. GitHub pull requests also require approval from reviewers and maintainers before merging.

- [x] **Document assignment of security response roles and how reports are handled.**

The [SECURITY.md](https://github.com/kyverno/community/blob/main/SECURITY.md) file documents security response and assignment: the Kyverno security team will respond within 3–5 days and plan a fix within 7–28 days based on severity. Maintainers manage security releases and coordinate disclosure. 

- [x] **Document Security Self-Assessment.**

The Kyverno project has conducted a [security self assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/kyverno/self-assessment.md) and participated in a thorough [Joint Security Assessment](https://docs.google.com/document/d/1JJVlYu6LNpWs2ukirOmoHMEhCsuOEoseBq6dq-Kcgjc/edit?tab=t.0#heading=h.9i73ao98ff5y) with TAG Security and Compliance. No blockers have been identified at this time, however, the project is highly encouraged to complete the items identified in the [tracking issue](https://github.com/kyverno/kyverno/issues/15335) by the next Kyverno release.

- [x] **Third Party Security Review.**

  - [x] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

A comprehensive third‑party security audit was conducted by Ada Logics. The audit defined a threat model, manually reviewed code, assessed fuzzing coverage, and evaluated supply‑chain security. Ten issues were found (including six CVEs) and promptly fixed in releases v1.10.6 and v1.11.1. The audit concluded that Kyverno complies with SLSA level 3 and found no policy bypasses. 

* [Kyverno Third-Party Security Audit](https://kyverno.io/docs/security/#kyverno-third-party-security-audit-2023)
* [Kyverno Fuzzing Security Audit](https://kyverno.io/docs/security/#kyverno-fuzzing-security-audit-2023)

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The Kyverno project maintains a [passing badge](https://www.bestpractices.dev/en/projects/5327/passing) at this point in time.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

The [ADOPTERS.md](https://github.com/kyverno/kyverno/blob/main/ADOPTERS.md) file lists dozens of organizations using Kyverno, with descriptions of their use cases (production, dev/trial). Examples include Vodafone enforcing policies and automation, Deutsche Telekom enforcing security rules, and VELUX using Kyverno for security and best practices. Additional adopters from finance, cloud providers, media, telecom, and government demonstrate broad production usage.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [x] **TOC verification of adopters.**

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

* Integrates with Kubernetes as an [admission controller](https://kyverno.io/docs/guides/admission-controllers/) for policy enforcement and validation
* [Integrates](https://kyverno.io/docs/subprojects/authz/#_top) with Istio, Envoy Gateway, and agentgateway using Envoy External Authorization
* [Documented instructions to use Kyverno with ArgoCD](https://kyverno.io/docs/installation/platform-notes/#notes-for-argocd-users)  
* [Documented instructions to install Kyverno via Helm](https://kyverno.io/docs/installation/installation/#helm-chart)
* [Backstage plugin to integrate Policy Reporter](https://kyverno.io/docs/subprojects/backstage-plugin/)
 
#### Adoption

##### Adopter 1 - LinkedIn/Technology

JANUARY 2026

Shan Valleru, Senior Software Engineer and Nicholas Wedlake, Staff Security Engineer

LinkedIn’s platform team (Shan Valleru and Nicholas Wedlake) provided feedback on their experience scaling Kyverno. While they praise its ease of use and community responsiveness, they highlighted a need for more rigorous internal testing and advanced documentation for complex enterprise scenarios.

**Project Strengths**

- **YAML-First Authoring:** The ability to write policies in YAML is a major differentiator. It allows security teams to present controls in a format that application developers can easily understand without specialized training in domain-specific languages.
- **Granular Control:** The project offers high-level precision, allowing for specific scoping against kinds, namespaces, and pod elements to effectively block or audit behaviors.
- **Strong CI/CD Integration:** The Kyverno CLI is a key asset. It is used effectively to validate policy changes within the CI/CD pipelines (e.g., against Kind clusters) before they reach production.
- **Modular Deployment:** The evolution toward a simplified Helm chart is highly valued. It allows adopters to toggle off unnecessary components (like background controllers or scan reports) to optimize resource usage.
- **Responsive Community:** The maintainers were commended for their quick response times in fixing critical bugs, particularly those related to admission review latency and memory utilization.

**Areas for Improvement**

- **Regression & Performance Testing:** 
  - **Internal Regressions:** There is a need for more robust internal testing to catch regressions before release (e.g., the instance where a metrics refresh flag stopped functioning after a migration).
  - **Proactive Scale Testing:** Rather than relying on large-scale users to find limits, the project should perform more proactive, high-scale performance testing internally.
- **Documentation for Advanced Use Cases:**
  - **Complex Logic:** There is a lack of documentation for complex, real-world scenarios, such as enforcing pod-level policies based specifically on namespace labels.
  - **Block Nuance:** The documentation needs to better clarify the functional differences and best practices for using validate vs. deny blocks.
- Administrative User Experience: While policy authoring is "low-code," the administrative experience of managing Kyverno at scale is still perceived as requiring "power user" expertise. Simplifying the operational management of the controllers would be beneficial.
- **Project Governance:** Continuing to diversify the maintainer base beyond a single primary company and strengthening the technical leadership ladder remains a point of interest for long-term stability.
Adopter interview notes held privately for TOC review and verification.

##### Adopter 2 - AI Development

Adopter interview notes held privately for TOC review and verification.

JANUARY 2026

##### Adopter 3 - Securities

Adopter interview notes held privately for TOC review and verification.

DECEMBER 2025

##### Adopter 4 - Retail

Adopter interview notes held privately for TOC review and verification.

DECEMBER 2025
