# Fluid Incubation Due Diligence

- [Fluid Incubation application issue](https://github.com/cncf/toc/issues/1317)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Fluid

### Criteria Evaluation

Katie Gamanji (@kgamanji) and Alex Chircop (@chira001) conducted the due diligence of Fluid who applied for Incubation. The project has completed the criteria that show its maturity at Incubation. The following criteria implementations are noteworthy to call out:
- TODO ... $NOTABLES. 

The following actions were provided to the project that were considered blocking but since resolved:
- the project presented to TAG Storage over 1y ago. Considering the ongoing TAG Reboot initiative, maintainers were advised to complete the General Technical Review for the project
- the security contact page was returning a 404 error in [Reporting a Vulnerability](https://github.com/fluid-cloudnative/fluid/blob/master/SECURITY.md#private-disclosure-processes) and should be fixed
- provide a  community meeting calendar to simplify the path for contributors and adopters to join future calls. It is specified that the calls are on a bi-weekly cadence, but the date of the next meeting date is unknown
- fix the inconsistency in community meeting times. The agenda to propose topics specifies a Zoom link as opposed to DingTalk Group where the calls are actually taking place. Consolidation and removal of references to unused resources is required. 
- the link to CoC in Join Our [Community as a Member](https://github.com/fluid-cloudnative/fluid/blob/a6c6343a0ee448f1498aafcdfecada55d5fa283b/CONTRIBUTING.md#join-our-community-as-a-member) returns a 404 error and should be fixed. 

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

- Consider renaming the development branch from master -> main e.g. [Kubernetes Default Branch Migration](https://www.kubernetes.dev/resources/rename/). Also, note that renaming the branch would require an update for the contributor and release process guidelines
- Community meetings are taking place regularly and are held in Mandarin. No blockers here, however, would suggest working with CNCF staff to use an automated transcript tool if possible.
- take action on the outstanding items that lower the [OpenSSF scorecard score](https://scorecard.dev/viewer/?uri=github.com/fluid-cloudnative/fluid)
- for the listed [Fluid Adopters](https://github.com/fluid-cloudnative/fluid/blob/master/ADOPTERS.md) all GitHub handles for "contact" are 404ing and pointing to email addresses of adopters. Email addresses could represent a privacy risk, and we would recommend pointing to GitHub handlers instead.

### Adoption Evaluation

The adopter interviews reflect the project use for the Incubation level to which the project has applied.

TODO

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.


## Application Process Principles

### Suggested

N/A

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**
  - The project presented to TAG Storage on 10-07-2024, and can be discovered at [here](https://www.youtube.com/watch?v=kBZNRbP4nOU).

<!-- (TOC Evaluation goes here) --> 

- [x] **TAG provides insight/recommendation of the project in the context of the landscape**

<!-- (TOC Evaluation goes here) --> 

The feedback from the TAG was positive, however it was conducted more than 1 year ago at the time of the project assessment for Incubation.  The project maintainers were asked to complete the General Technical Review for Fluid, which can be found [here](https://github.com/fluid-cloudnative/community/blob/1215ac22dc7b757cc9b63eee8660c3e15210ba8f/docs/general-technical-review-cncf-incubating.md).

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

<!-- (TOC Evaluation goes here) --> 

The project fulfills the vendor-neutral criteria. The main communication channel is on DingTalk and WeChat groups. Also the project uses [Slack](https://cloud-native.slack.com/archives/C02ADG209SP), on the CNCF workspace via `#fluid` channel. 
[Community meetings](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) are held on DingTalk with a bi-weekly cadence. Users are also capable of finding the meeting notes on [Fluid GitHub Community Repository](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) and to follow up on the latest topics of discussion.


- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
- Met during Project's application on 27-Apr-2021.

<!-- (TOC Evaluation goes here) --> 

Fluid Sandbox application can be found [here](https://github.com/cncf/sandbox/issues/245).

The project maintainers have understood the expectations for Sandbox projects and requirements for moving forward through the CNCF incubation level.

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

<!-- (TOC Evaluation goes here) --> 

The project maintainers have added the following additional documentation as follows.
- [Installation documentation](https://fluid-cloudnative.github.io/docs/next/get-started/installation)
- [User documentation](https://fluid-cloudnative.github.io/docs/next)
- [Architecture documentation](https://fluid-cloudnative.github.io/docs/next/core-concepts/architecture-and-concepts)

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

<!-- (TOC Evaluation goes here) --> 

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


### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) --> 
The project provides [3 levels of engagement routes](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/GOVERNANCE.md) for community members, each one building on the previous responsibilities:
- Contributors - base level interaction on the project through comments on issues or pull request 
- Committers - community members who have shown that they are committed to the continued development of the project through ongoing engagement 
- Maintainers - lead the development of the project, through contributions to increasingly complicated PRs/designs and review PRs/designs, under the guidance of the existing maintainers

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) --> 
Fluid makes use of GitHub issues and PRs. The following guidelines are available:
- [Filling an issue](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#filing-issues), including templates for:
  - [Report a bug](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/.github/ISSUE_TEMPLATE/bug-report.md)
  - [Document improvements](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/.github/ISSUE_TEMPLATE/docs-defect.md)
  - [Feature Request](https://github.com/fluid-cloudnative/fluid/blob/master/.github/ISSUE_TEMPLATE/feature-request.md)
- with a list of examples for [code contributions](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#code-contributions)

The project also provides a guide on how to set up a [development workspace
](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#setting-up-development-workspace) and submit a PR through available [template](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/.github/PULL_REQUEST_TEMPLATE.md).


- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) --> 

Public communication channels for Fluid are listed [here](https://github.com/fluid-cloudnative/fluid?tab=readme-ov-file#community), including a DingTalk and WeChat groups, and Slack. 

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) --> 
The main channel of communication are listed below: 
- CNCF Slack channel: [#fluid](https://cloud-native.slack.com/archives/C02ADG209SP)
- WeChat Group: Fluid Open Source Group
- DingTalk Group: Fluid Open Source Group

[Community meetings](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) are held on [DingTalk Group](https://qr.dingtalk.com/action/joingroup) with a bi-weekly cadence. 

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) --> 
Community members are able to join community meetings on a bi-weekly basis. The meeting notes can be found on [Fluid GitHub Community Repository](https://github.com/fluid-cloudnative/community/wiki/Meeting-Schedule) and to follow up on the latest topics of discussion.

Community meetings are held in Mandarin, and are a great way to engage with the maintainers and get involved in project development. 

TODO: calendar for meetings


- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) --> 

Details on how to contribute and engage with the project are listed on the [Contributing Guidelines](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md), including a path to [join the community as a member](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CONTRIBUTING.md#join-our-community-as-a-member). The contributing members are also encouraged to follow the [Code of Conduct](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/CODE_OF_CONDUCT.md) as part of getting involved. 


- [x] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) --> 
The project has a good history of adding new contributors e.g. [listing 10 new contributors](https://github.com/fluid-cloudnative/fluid/commit/37ddcf63e88303578fee99e3970465fbddc404e3) in addition to existing maintainers. Also, there is a good distribution of involved organizations, ensuring the continuity of the vendor-neutral development of Fluid.  

Also, over the last 2 years the project has a consistent ratio of contribution vs involved contributors, that can observed in the [DevStats contribution distribution](https://fluid.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now). 


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

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) --> 
The [SECURITY.md](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/SECURITY.md#reporting-a-vulnerability) file outlines the vulnerability reporting process, including a mailing list (fluid.opensource.project@gmail.com) monitored by the maintainers that can be used for private disclosures.

Community members have the option to reach out to [security contacts](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/SECURITY_CONTACTS) to kickstart any private or public disclosure processes. 

Successful examples of fixed reported security issues, can be found below:
- [On a compromised node, the fluid-csi service account can be used to modify node specs](https://github.com/fluid-cloudnative/fluid/security/advisories/GHSA-93xx-cvmc-9w3v)
- [OS Command Injection for Fluid Users with JuicefsRuntime](https://github.com/fluid-cloudnative/fluid/security/advisories/GHSA-wx8q-4gm9-rj2g)


- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) --> 

2FA required for org members; branch protections enabled.

To be added as a GitHub member within the Fluid organization, each member should enable two-factor authentication (2FA). Also, the repository has branch protection rules enabled, which enforce certain workflows or requirements before a collaborator can push changes to a branch in the repository.

- [x] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) --> 
One of the core responsibility for maintainers is to respond to time-sensitive security release processes. Although this should be a rare occurrence, if a serious vulnerability is found the maintainers are expected to dedicate time to the fix, which is a process that might take up to several full days of work to implement.

[GitHub Security Advisor](https://github.com/fluid-cloudnative/fluid/security/advisories) is used as the communication channel during the process of identifying, fixing & shipping the mitigation of the reported vulnerability.

The advisory will only be made public when the patched version is released to inform the community of the breach and its potential security impact.

Additionally, security scanning is enabled for the  Fluid project and the[security contacts](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/SECURITY_CONTACTS) are responsible for assessing and providing a fix for the reported vulnerability.


- [x] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

The completed security self-assessment can be found [here](https://github.com/fluid-cloudnative/fluid/blob/3c08916896125990963b34d0b7fb6caae2c78bda/security/self-assessment.md).


- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) --> 
The project has a passing [OpenSSF Best Practices](https://www.bestpractices.dev/en/projects/4886) badge, with a 100% completion level, which is linked in [Fluid](https://github.com/fluid-cloudnative/fluid/tree/master) repository.

In addition, Fluid has [OpenSSF Scorecard report](https://scorecard.dev/viewer/?uri=github.com/fluid-cloudnative/fluid) via a score of 9.1 at the time of due diligence assessment. 



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


##### Adopter 3 - $COMPANY/$INDUSTRY

_If the Adopting organization needs to remain anonymous, stating the industry vertical is sufficient._
MONTH YEAR
