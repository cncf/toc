# Kubescape Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1291)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Kubescape

### Criteria Evaluation

_Emily Fox (@TheFoxAtWork) conducted the due diligence of Kubescape who applied for Incubation Level. The project [has/has not] completed the criteria that show its maturity at Incubation lelve. The following criteria implementations are noteworthy to call out... $NOTABLES. The following actions were provided to the project that were considered blocking but since resolved... $BLOCKERS. The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project... $RECOMMENDATIONS._

### Adoption Evaluation

_The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY._

### Final Assessment

_[The TOC has found the project to have satisfied the criteria for $LEVEL/ The TOC's evaluation of the project shows a needed focus to complete the outstanding blockers and reapply when the following conditions are met ... $CONDITIONS]._

## Application Process Principles

### Suggested

N/A

### Required

- [X] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**
  - This was completed and occurred on 10-JULY-2024, and can be discovered at the [TAG's meeting notes for that session](https://docs.google.com/document/d/1XDmNG1P1YRnidQJEuZFKufaXaIenNHiY-TdYMlt0Un4/edit#heading=h.cbsth4x8m975) and the [issue link for that presentation](https://github.com/cncf/tag-security/issues/1262).

The Project presented to TAG Security on 10-JULY-2024 and provided the community with an update on new capabilities of the project, feedback by the community and the TAG was engaging and did not appear to raise any concerns with the project.

- [X] **TAG provides insight/recommendation of the project in the context of the landscape**

At the July 10th, 2024 presentation to TAG Security, Kubescape was provided the following recommendations: 

> Recommendation for the project to engage with the End User TAB for seeking feedback on these features.

> Recommendation to look at existing open source projects (SIG-infra, etc.) with substantial infrastructure (build systems for instance) to see if they are willing to test out those features.

"Features" referring to several new capabilities that allows the project to auto-generate policies/configuration that support its goals to provide a runtime-powered k8s security platform.  To provide all k8s security needs under a single project.

- [X] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The project's communications, planning, and docs align with CNCF stated guidelines for project neutrality. It is noted that the project does maintain a providers file, for which ARMO is the only one listed. However, any contributor is welcome to submit a PR to add additional providers in accordance with the information provided on the [project's docs page for providers](https://kubescape.io/docs/providers/).

- [X] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
- Met during Project's [initial application on 13-DEC-2022 to CNCF](https://github.com/cncf/toc/issues/978) and reflected through initial application to [Incubation on 13-NOV-2023](https://github.com/cncf/toc/pull/1209) and subsequent [application in the updated format on 4-APR-2024](https://github.com/cncf/toc/issues/1291)

- [ ] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisifies the Due Diligence Review criteria.

- [ ] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The project provides adopters with appropriate documentation for installation and configuration of Kubescape, to include necessary prerequisites, 

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [ ] **Clear and discoverable project governance documentation.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

<!-- (TOC Evaluation goes here) --> 

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

Kubescape maintains its list of maintainers for the core repository at https://github.com/kubescape/kubescape/blob/master/MAINTAINERS.md. This file has been changed 6 times since its creation on 06-JAN-2022.

The project has the following additional respositories with Maintainer files:
* Helm-charts: https://github.com/kubescape/helm-charts/blob/main/MAINTAINERS.md
* regolibrary: https://github.com/kubescape/regolibrary/blob/master/MAINTAINERS.md - added 05-NOV-2023, no changes since. No roles defined, clear association with regolibrary

- [ ] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document agreement that project will adopt CNCF Code of Conduct.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **CNCF Code of Conduct is cross-linked from other governance documents.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **All subprojects, if any, are listed.**

<!-- (TOC Evaluation goes here) --> 

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

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

- [ ] **Roadmap change process is documented.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **History of regular, quality releases.**

<!-- (TOC Evaluation goes here) --> 

### Required 

- [ ] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document what the project does, and why it does it - including viable cloud native use cases.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document the project's release process.**

<!-- (TOC Evaluation goes here) --> 

## Security

Note: this section may be augemented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [ ] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) --> 

## Ecosystem

### Suggested

N/A

### Required

- [ ] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

<!-- (TOC Evaluation goes here) --> 

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [ ] **TOC verification of adopters.**

<!-- (TOC Evaluation goes here) --> 

Refer to the Adoption portion of this document.

- [ ] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

<!-- (TOC Evaluation goes here) --> 

#### Adoption

##### Adopter 1 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR

##### Adopter 2 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR

##### Adopter 3 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR
