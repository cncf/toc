# cert-manager Graduation Due Diligence

- [cert-manager graduation application issue](https://github.com/cncf/toc/issues/1306)

## Graduation Evaluation Summary for cert-manager

### Criteria Evaluation

Katie Gamanji (@kgamanji) conducted the due diligence of cert-manager who applied for Graduation. The project has completed the criteria that show its maturity at the applied level. The following criteria implementations are noteworthy to call out:

- A notable stable project with mature capabilities and a wide adopter base.
- The integration capabilities with Istio, SPIFFE, and other systems are noteworthy, enhancing both security and flexibility.
- Thorough documentation of build and deployment processes is excellent.
- Implementation of the steering committee to capture adopters' voice in the project development and roadmap.
- The completion of the external security audit by Ada Logics and the project's efforts to address the findings demonstrate a strong commitment to security. All eight identified issues have been addressed, which is commendable.
- cert-manager has an exceptionally complete [contributing guide](https://cert-manager.io/docs/contributing/), including information about communications, feature approval, expectations, testing, release process, and coding standards.  This is a good example to other projects about what a complete contribution guide should look like.
- The project does an excellent job of making sure that its public meetings are accessible, with notes, videos, and easy to find meeting links.



The following actions were provided to the project that were considered blocking but have since been resolved:

- Share the steering committee meetings and notes as stated per governance: “All recordings and minutes will be posted publicly”, including  progress, input or impact of the committee on the project.
- Update the Roadmap to reflect the ongoing work and planned efforts. Also, as part of the governance, document the roadmap change process.
- Provide an up-to-date list of sub-projects to be included as part of the Graduation evaluation.
- Provide a RELEASE.md file, including updating the release process to reflect the latest shift to the CNCF infrastructure.
- Update the Code of Conduct to reference the CNCF Code of Conduct. Currently the Code of Conduct is adapted from the Contributor Covenant, however one of the incubation criteria is to explicitly reference and use the CNCF CoC https://github.com/cncf/foundation/blob/main/code-of-conduct.md.

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure the continued viability of the project:
- Complete a joint assessment with TAG Security reviewers.
- TAG Security encourages the expansion of the current use of go vet and implement govulncheck as planned. It would be beneficial to update your OpenSSF Best Practices information if these have been addressed.
- TAG Security recommends completing the silver and gold level criteria, as the project likely already meets most of them.
- TAG Security suggests considering security audits for sub-projects like trust-manager and csi-driver. Perhaps the CNCF can batch these in a follow-up audit.
- TAG Contributor Strategy recommends to add in/out of scope information to the readme, roadmap, or contributor documentation.
- TAG Contributor Strategy recommends to add role qualifications for each step on the contributor ladder.  This is particularly important for the maintainers, where other docs refer to Maintainer qualifications that don't exist.
- TAG Contributor Strategy recommends to add a process for removing Maintainers (and SC members) for reasons other than inactivity, such as violating the CoC or disruptive behavior.
- TAG Contributor Strategy recommends to link to the list of official channels and meetings in the contributor docs from the Governance document.
- TAG Contributor Strategy recommends to gradually build up the sub-projects into their own entities, allowing new contributors to take ownership of them.  This will require adding to the main Governance for these roles.
- TAG Contributor Strategy recommends to have another Steering meeting, in order to keep Steering members engaged.
- TAG Contributor Strategy recommends to make sure that the Community repo is linked from appropriate other places, like the main development repos and the contributor docs.
- TAG Contributor Strategy recommends to figure out a low-effort way to record maintainer decisions for posterity, such as a simple text log.

### Adoption Evaluation
<!-- (TOC Evaluation goes here) -->
The adopter interviews reflect the project use for the Graduation level to which the project has applied. They show that the project aims to demystify the concept of TLS and make it approachable for easy consumption. Also, the project is very clear about the functionalities it supports and/or are out of scope, setting pronounced boundaries in its scope and roadmap. In these terms, cert-manager has brought substantial value to adopters, as it provided an opportunity to learn about certificate management and TLS, and automation of manual processes. Overall, the stability of the projects and their API stand out over time, and that builds a strong sense of trust with its adopters.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation.

### Criteria

## Application Process Principles

### Required

- [x] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

<!-- (TOC Evaluation goes here) -->
The project presented to:
 - TAG Security on 2024-05-22: https://github.com/cncf/tag-security/issues/1254
 - TAG Contrib Strategy on 2024-05-02: https://github.com/cert-manager/community/issues/14. As well, the project requested an official review from the TAG in [[Governance Review]: cert-manager]( https://github.com/cncf/tag-contributor-strategy/issues/675) issue.



- [x]  **TAG provides insight/recommendation of the project in the context of the landscape**

<!-- (TOC Evaluation goes here) -->
Both, TAG Security and TAG Contrib Strategy, are supportive of cert-manager's progression to graduated status within CNCF. The project demonstrates maturity and a strong commitment to security and community practices.

- TAG Security evaluation can be found in this [comment](https://github.com/cncf/toc/issues/1306#issuecomment-2211472304).

- TAG Contrib Strategy evaluation can be found in this [PR #694: Final draft of Cert-Manager review](https://github.com/cncf/tag-contributor-strategy/pull/694).


- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

<!-- (TOC Evaluation goes here) -->
The project fulfils the vendor-neutral criteria. The main communication channel is [Slack](https://cert-manager.io/docs/contributing/#slack), on the Kubernetes workspace. Community meetings are held on [Zoom](https://cert-manager.io/docs/contributing/#community-meetings) with a bi-weekly cadence and showcased on the CNCF calendar.  Users are also capable of joining the [mailing list](https://groups.google.com/g/cert-manager-dev) to follow up on the latest topics of discussion.

Regarding the architectural decisions, cert-manager is developed independently and supplies easy integration with multiple providers. Currently, the project has core maintainers from  Venafi, Diagrid and G-Research.

The project is self-governing, adhering to the principles of exercising minimal viable governance to enable community driven success. To further the sustainable development of the project, the  [Steering Committee](https://github.com/cert-manager/community/blob/main/STEERING.md) was introduced to represent the voice of the users and work with the maintainers to help guide development efforts towards solving concrete and immediate problems for cert-manager adopters.

- [x] **Review and acknowledgment of expectations for graduated projects and requirements for moving forward through the CNCF Maturity levels.**		
   - [x] Met during the Project's application on 23-04-2024. [Initial PR](https://github.com/cncf/toc/pull/1212) was submitted on 28-11-2023. The project opened a [new issue](https://github.com/cncf/toc/issues/1306) using the recommended template on 23-04-2024.

<!-- (TOC Evaluation goes here) -->
The TOC and project maintainers met during the kick-off meeting of the Graduation process to discuss the changes in the TOC repository regarding the criteria and expectations for moving levels, as well as the freeze period, since KubeCon & CloudNativeCon was approaching at that time. Additionally, the TOC provided a list of [suggested action items](https://github.com/cncf/toc/issues/1306#issuecomment-2072815310) that maintainers can start working on and address accordingly.  



Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

<!-- (TOC Evaluation goes here) -->
  - [Installation documentation](https://cert-manager.io/docs/installation/)
  - [Getting started](https://cert-manager.io/docs/getting-started/)
  - [Tutorials](https://cert-manager.io/docs/tutorials/)
  - [Code samples and DevOps tips](https://cert-manager.io/docs/installation/compatibility/)


## Governance and Maintainers
Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

<!-- (TOC Evaluation goes here) -->
The project fulfils this criteria. The Governance is exercised regularly, such as moving maintainers to [emeritus status](https://github.com/cert-manager/community/commit/6bd339d01ff06ee49f0bbeda80fb531dccd94c2d) and the creation of the [Steering Committee](https://github.com/cert-manager/community/blob/main/STEERING.md) to bring the end user perspective to the project development.

### Required

- [x] **Clear and discoverable project governance documentation.**

<!-- (TOC Evaluation goes here) -->
Project Governance can be found [here](https://github.com/cert-manager/community).

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

<!-- (TOC Evaluation goes here) -->
Since opening the Graduation issue, the project has updated its Governance and is now up to date with actual project activities, this includes iterations on the steering committee process and meetings, as well as its influence on the roadmap.

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

<!-- (TOC Evaluation goes here) -->
Decisions are made by maintainers who are actively involved in the project, and and there is a clear path outlined for people to become maintainers. Current maintainers are spread across several companies and the project is also encouraging new maintainers to join. In the past, non-maintainers have certainly contributed to decision making process during meetings. [Erik](https://github.com/erikgb) is a good example, he has contributed with a lot of discussion in the daily standups, e.g. suggest an approach to solving a project problem.

The project roadmap is based on the [guide to open-source project roadmaps](https://contribute.cncf.io/maintainers/community/contributor-growth-framework/open-source-roadmaps/#components-of-a-typical-roadmap), which has vendor neutrality at its core.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

<!-- (TOC Evaluation goes here) -->
This is included in project governance and contributors docs, including:
- [Maintainer decision making process](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md#maintainer-decision-making) which outlines that most decisions are made through the principle of lazy consensus. There also exists the option for a maintainer to propose a formal vote. Unless otherwise specified, such a vote would require a simple majority approval from all maintainers to be considered successful.
- [Contributing flow](https://cert-manager.io/docs/contributing/contributing-flow/) for new features, bugs and overall project management.
- [Feature policy](https://cert-manager.io/docs/contributing/policy/) for both, feature requests and PRs, including examples of  previous proposals.
- [Steering Committee](https://github.com/cert-manager/community/blob/main/STEERING.md) charter to bring the end user perspective to the project development.


- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

<!-- (TOC Evaluation goes here) -->
The project has six levels of responsibility, each one building on the previous:

- [Contributor](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#contributor)
- [GitHub Member](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#github-member)
- [Reviewer](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#reviewer)
- [Approver](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#approver)
- [Maintainer](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#maintainer)
- [Admin](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#admin)

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

<!-- (TOC Evaluation goes here) -->
All maintainers share all domains of responsibility. Current list of maintainers can be found [here](https://github.com/cert-manager/community/blob/main/maintainers.csv).

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) -->
The maintainer list currently enumerates 8 maintainers, including [Adam](https://github.com/cert-manager/community/pull/23) who was onboarded recently.


- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

<!-- (TOC Evaluation goes here) -->
Maintainer lifecycle process is documented as part of the GOVERNANCE.md file, including:
- [How to becoming a Maintainer](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md#becoming-a-maintainer)
- [Maintainer privileges](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md#maintainer-privileges)
- [Maintainer responsibilities](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md#maintainer-responsibilities)
- [Stepping down as a Maintainer](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md#stepping-down-as-a-maintainer)

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

<!-- (TOC Evaluation goes here) -->
Recently a couple of maintainers were moved to [emeritus status](https://github.com/cert-manager/community/commit/6bd339d01ff06ee49f0bbeda80fb531dccd94c2d) as they have drifted away from the project activities.

Recently the project onboarded the following new maintainers:
 - Adam ([cert-manager/community#23](https://github.com/cert-manager/community/pull/23)) - which was a good test of the process
 - Erik ([cert-manager/community#32](https://github.com/cert-manager/community/pull/32))

There are a few other community members at differing stages along the maintainership path, including:
- Houssem (Venafi affiliated, has been an approver for istio-csr since January I 2024)
- Peter (Venafi affiliated - member of the cert-manager org and incredibly helpful with a lot of non-code tasks like debugging, testing, admin, taking notes in meetings)

- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

<!-- (TOC Evaluation goes here) -->
As of July 2024, the project has core maintainers from Venafi, Diagrid and G-Research.

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

<!-- (TOC Evaluation goes here) -->
This is documented in the governance process, as [Admin privileges](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md#admin). Maintainers get roles in GCP for managing tests and release infrastructure (part of [prerequisites](https://cert-manager.io/docs/contributing/release-process/#prerequisites) for the release process), and maintainers get elevated privileges in the GitHub org.

- [x] **Document agreement that project will adopt CNCF Code of Conduct.**

<!-- (TOC Evaluation goes here) -->
The project is operating under the CNCF Community Code of Conduct, which is reproduced below from upstream commit [c0d29e5e94be2ac2c1370d663ce3a08ada77921f](https://github.com/cncf/foundation/blob/c0d29e5e94be2ac2c1370d663ce3a08ada77921f/code-of-conduct.md). The Code of Conduct can be found [here](https://github.com/cert-manager/community/blob/528163770f5ba1d75db133dd69f0307b3b555876/CODE_OF_CONDUCT.md).


- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

<!-- (TOC Evaluation goes here) -->
One of the initial requirements to become a contributor for the project is to [follow the cert-manager Code of Conduct](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md#contributor-responsibilities).



- [x] **All subprojects, if any, are listed.**

<!-- (TOC Evaluation goes here) -->
All subprojects are listed on the [project website](https://cert-manager.io/docs/contributing/projects). Currently, these are the listed subprojects:
- [trust-manager](https://github.com/cert-manager/trust-manager)
- [approver-policy](https://github.com/cert-manager/approver-policy)
- [csi-driver](https://github.com/cert-manager/csi-driver)
- [csi-driver-spiffe](https://github.com/cert-manager/csi-driver-spiffe)
- [istio-csr](https://github.com/cert-manager/istio-csr)

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

<!-- (TOC Evaluation goes here) -->
Mentioned at the beginning of the GOVERNANCE.md document:

>This governance charter applies to every project under the cert-manager GitHub organisation. The term "cert-manager project" refers to any work done under the cert-manager GitHub organisation and includes the cert-manager/cert-manager repository itself as well as cert-manager/trust-manager, cert-manager/approver-policy and all the other repositories under the cert-manager GitHub organisation.


## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from TAG Contributor Strategy.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) -->
The project has six levels of responsibility, each one building on the previous:

- [Contributor](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#contributor)
- [GitHub Member](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#github-member)
- [Reviewer](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#reviewer)
- [Approver](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#approver)
- [Maintainer](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#maintainer)
- [Admin](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md#admin)


### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) -->
cert-manager makes use of GitHub issues and PRs. The following templates are available:
- [Report a bug](https://github.com/cert-manager/cert-manager/blob/fc198e91979a5c1e869d2fc45750de2d6aacea7a/.github/ISSUE_TEMPLATE/bug.md)
- [Feature request](https://github.com/cert-manager/cert-manager/blob/fc198e91979a5c1e869d2fc45750de2d6aacea7a/.github/ISSUE_TEMPLATE/feature-request.md)
- [Submit a Pull Request ](https://github.com/cert-manager/cert-manager/blob/fc198e91979a5c1e869d2fc45750de2d6aacea7a/.github/PULL_REQUEST_TEMPLATE.md)

The project also has a [feature policy](https://cert-manager.io/docs/contributing/policy/) guide to set expectations when an lager change is proposed by the community.


- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) -->
Public communication channels are listed on the cert-manager [Contributing guide](https://cert-manager.io/docs/contributing/) page.


- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) -->
The main communication channel is [Slack](https://cert-manager.io/docs/contributing/#slack), on the Kubernetes workspace (#cert-manager-dev).
The projects has two cert-manager channels on [Kubernetes Slack](https://slack.k8s.io/):

- [cert-manager](https://kubernetes.slack.com/messages/cert-manager) - for all users of cert-manager; dedicated for any usage related questions
- [cert-manager-dev](https://kubernetes.slack.com/messages/cert-manager-dev) - for collaboration between cert-manager contributors and maintainers; dedicated for code related questions

Community meetings are held on [Zoom](https://cert-manager.io/docs/contributing/#community-meetings) with a bi-weekly cadence and showcased on the CNCF calendar.

Users are also capable of joining the [mailing list](https://groups.google.com/g/cert-manager-dev) cert-manager-maintainers@googlegroups.com to follow up on the latest topics of discussion.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) -->
cert-manager has bi-weekly (i.e. every two weeks) community meetings happen [on Zoom](https://zoom.us/j/96658445866) at 17:00 London time. Community members can look at the [CNCF calendar](https://tockify.com/cncf.public.events/monthly?search=cert-manager) to check when the next call will take place.

These meetings are an hour-long chat about cert-manager topics. It's a great way to get involved with contributing for the first time; to get answers to any questions you might have; or to propose a new feature which needs some explanation.

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) -->
Details on how to contribute and engage with the project are listed on the [Contributing Flow](https://cert-manager.io/docs/contributing/contributing-flow/) page, including an overview of GitHub bot, bugs, features, creating Pull Requests
and overall Project Management.

- [x] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) -->
The project regularly onboards new contributors, and each minor cert-manager release contains a thank-you list of who was involved. For example, the alpha release of the next minor release already lists 10 contributors and with plans to continue to growing the contributor base: https://github.com/cert-manager/cert-manager/releases/tag/v1.15.0-alpha.0

## Engineering Principles

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

<!-- (TOC Evaluation goes here) -->
The first line of every cert-manager release re-states the project ambition:

>cert-manager is the easiest way to automatically manage certificates in Kubernetes and OpenShift clusters.

Over time, the project has essentially become the standard for managing certificates in a cloud native/Kubernetes-native way.


- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

<!-- (TOC Evaluation goes here) -->

This is extensively documented in the [intro section](https://cert-manager.io/docs/) of the project.

It states the goal of the project to create TLS certificates for workloads in Kubernetes or OpenShift clusters and renews the certificates before they expire. It also highlights a variety of certificate authorities that can be used to obtain certificates, including: Let's Encrypt, HashiCorp Vault, Venafi and private PKI.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

<!-- (TOC Evaluation goes here) -->
The maintainers have iterated on the previous roadmap along with the steering committee and moved it into the community repo. Current roadmap can be found [here](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/ROADMAP.md).

- [x] **Roadmap change process is documented.**

<!-- (TOC Evaluation goes here) -->

Changes to the roadmap are documented [here](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/ROADMAP.md#changes-to-this-roadmap). All proposals should be submitted via a pull requests containing the suggested change.

Significant changes to the roadmap should be discussed in either a steering committee meeting or a biweekly meeting before merging, to raise awareness of the change and to provide an opportunity for discussion.

Insignificant changes, such as updating links to issues, spelling fixes or minor rewordings,  do not need to be discussed in a meeting but should still be shared in Slack.


- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

<!-- (TOC Evaluation goes here) -->
cert-manager architecture is covered as part of the [intro section](https://cert-manager.io/docs/) of the project.


- [ ] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:**

  - [x] Release expectations (scheduled or based on feature implementation)

  The project aims to do regular releases roughly every 4 months but release dates can vary when accounting for holidays, conferences (such as KubeCon & CloudNativeCon), maintainer commitments and other world events.

  - [x] Tagging as stable, unstable, and security related releases

  The project uses the alpha and beta to mark their releases. Security and bug fixes strategy is outlined [here](https://cert-manager.io/docs/releases/#security-and-bug-fixes). The project back-ports important bug fixes — including security fixes — to all currently supported releases.

  - [x] Information on branch and tag strategies

    Tagging strategy is outlined [here](https://cert-manager.io/docs/contributing/release-process/#process-for-releasing-a-version).

  - [x] Branch and platform support and length of support

  Current supported releases can be found [here](https://cert-manager.io/docs/releases/#currently-supported-releases).

  - [x] Artifacts included in the release.

  Based on the [RELEASE.md](https://github.com/cert-manager/cert-manager/blob/master/RELEASE.md) file cert-manager will produce the following artifacts each release:

  - Container Images - Container images for the cert-manager project are published for all cert-manager components.
  - Helm chart - An official Helm chart is maintained within this repo and published to charts.jetstack.io on each cert-manager release.
  - Binaries - Until version 1.15 the cmctl binary was maintained within this repo and published as part of the cert-manager release. For releases after 1.15 the CLI has moved to its own repository. Binary builds are still available for download from this new location.

  - [x] *Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.

  cert-manager 1.12 is a Long Term Support (LTS) release sponsored by Venafi. It will continue to be supported for at least 2 years from release.

<!-- (TOC Evaluation goes here) -->
The release process is documented on the [website](https://cert-manager.io/docs/contributing/release-process/) and the support schedule is documented [here](https://cert-manager.io/docs/releases/).

The release process has been updated to reflect recent changes which mean that any maintainer can do a release, and that it is no longer a requirement to be a Venafi employee.

It is worth pointing out that there is still one thing which requires approval of a Venafi employee - that's to publish the Helm chart to charts.jetstack.io which is difficult to change since that repo is so widely used.

The project intends to provide an alternative install mechanism in the near future (see [cert-manager/cert-manager#7132](https://github.com/cert-manager/cert-manager/pull/7132)) which is entirely vendor neutral. The plan is to evolve away from depending on the chart repo, to OCI registries. Maintainers will keep pushing to the existing repository not to cause any breaking changes, but the long term vision for the primary way of consuming project charts is to pull from OCI registries, which should be accessible by any maintainer.


- [x] **History of regular, quality releases.**

<!-- (TOC Evaluation goes here) -->
The [supported releases](https://cert-manager.io/docs/releases/) page lists project's extensive release history for cert-manager.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

<!-- (TOC Evaluation goes here) -->
N/A

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) -->
The [SECURITY.md](https://github.com/cert-manager/community/blob/main/SECURITY.md) file outlines the vulnerability reporting process, including a mailing list (cert-manager-security@googlegroups.com) monitored by the maintainers.

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) -->
To be added as a GitHub member, each maintainer should enable [two-factor
authentication (2FA)](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/GOVERNANCE.md?plain=1#L67). Also, the project enforces a [DCO Sign Off](https://cert-manager.io/docs/contributing/sign-off/) process, require all contributors to acknowledge that they have the rights to the code they are contributing by signing their commits in git.

- [x] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) -->
One of the core maintainer responsibility is to respond to time-sensitive security release processes. The security response team will endeavour to reply as soon as possible, ideally within 3 working days.

If the team concludes that the reported issue is indeed a security vulnerability in a cert-manager project, at least two members of the security response team will discuss the next steps together as soon as possible, ideally within 24 hours.

As soon as the team decides that the report is of a genuine vulnerability, one of the team will respond to the reporter acknowledging the issue and establishing a disclosure timeline, which should be as soon as possible.

- [x] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) -->
Completed security self-assessment can be found [here](https://github.com/cncf/tag-security/blob/1f24478023d92c76372e92e478238a04cafd1658/community/assessments/projects/cert-manager/self-assessment.md).

- [x] **Third Party Security Review.**

  - [x] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

<!-- (TOC Evaluation goes here) -->

cert-manager completed CNCF-sponsored security audit earlier in 2024. The announcement can be read [here](https://cert-manager.io/announcements/2024/03/18/cert-manager-security-audit/), including the [full report](https://cert-manager.io/docs/announcements/AdaLogics-2023-cert-manager-audit-report.pdf).
A total of 8 issues were raised as part of the audit, of which 5 were low severity, 2 were moderate severity and 1 was informational. All issues have been resolved as of cert-manager v1.12.8, v1.13.4 and v1.14.3.

The completion of the external security audit by Ada Logics and projects' effort to address the findings demonstrate a strong commitment to security. All eight identified issues have been addressed, which was found commendable by TAG Security members.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) -->
The project has a passing [OpenSSF Best Practices](https://www.bestpractices.dev/en/projects/8079) badge, which is linked in cert-manager/cert-manager repository.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

<!-- (TOC Evaluation goes here) -->
A list of public adopters can be found in the [USERS.md](https://github.com/cert-manager/community/blob/fe1aa75b064820fb6c41741dfdb526a60d0b7617/USERS.md) file.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

<!-- (TOC Evaluation goes here) -->
The project provided the TOC with a list of adopters for verification of use of the project at the level expected, e.g. production use for graduation, dev/test for incubation.

The project is used by a large base of organisations, most of which prefer to remain anonymous.

- [x] **TOC verification of adopters.**

<!-- (TOC Evaluation goes here) -->

Refer to the Adoption portion of this document.

Maintainers provided the TOC sponsor with a list of 6 users who agreed to be interviewed for the Graduation Due Diligence process.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

<!-- (TOC Evaluation goes here) -->
Project documentation are clear around Kubernetes and OpenShift integrations. They also document integrations with [Hashicorp Vault](https://cert-manager.io/docs/configuration/vault/), [Let's Encrypt](https://cert-manager.io/docs/configuration/acme/),  [private PKI](https://cert-manager.io/docs/configuration/ca/), [Venafi](https://cert-manager.io/docs/configuration/venafi/), tens of other issuers, SPIFFE (via csi-driver-spiffe) and a huge variety of [cloud DNS providers](https://cert-manager.io/docs/configuration/acme/dns01/).


#### Adoption

##### Adopter 1 - Enterprise software offerings
Adopter 1 has been using cert-manager as a component of a packaged solution distributed to users as part of a SaaS  and enterprise software offerings. The project has been used since Day 1,  for more than 4 years thus far with great success. When evaluating the project for adoption, there were no other mature alternatives, hence, the cert-manager project stood up as a natural contender within the certificate management space.

The core features that the project provides are certificate consumption and issuing within a Kubernetes cluster, in addition to being compatible with external issuers' APIs. Also, from the consumer perspective, it builds end-to-end trust when a user accesses the product via an HTTPS endpoint. This feature set enables users to easily consume TLS within a cluster without too much overhead, fuelling Adopter 1 to integrate cert-manager within their product. Currently, every Kubernetes distribution of their packaged solution will have cert-manager enabled by default. Overall cert-manager has been straightforward to use and adopt, with the challenge of the learning curve necessary to comprehend the architecture and its components. Once grasped cert-manager is easy to use by different personas, such as infrastructure and platform engineers. For developers, the upskilling and educational piece plays a crucial role, as it exposes a lower level of complexity/platform management that is usually out of the scope of their role.

In regards to the release cadences, Adopter 1 keeps up to date with the upstream and will cut a new release for their product within the same week cert-manager distributes a new release. It is worthwhile to mention that over time cert-manager has proven to be a stable project, with a solid core and fundamental functionalities, that has built trust with adopters and the ability to use the latest release with ease.  Documentation is thorough and well-structured, and the latest doc index alterations made it simpler to identify topical guides and references, providing a balanced base to search for information. The code structure is navigable, however, it is necessary to comprehend the concept of Kubernetes controllers to understand the abstraction layer and identify the code for core functionalities

When it comes to community interactions, Adopter 1 expressed interest in Gateway API feature graduating, to which they were invited to a community call to discuss this further. Slack interactions were helpful as well, usually requiring a clear outline of the issue and pinging one of the maintainers to reach a resolution.

Adopter 1 strongly appreciates that cert-manager aims to demystify the concept of TLS and make it approachable for easy consumption. Also, the project is very clear about the functionalities it supports and/or are out of scope, setting pronounced boundaries in its scope and roadmap. In these terms, cert-manager has brought substantial value to Adopter 1, as it provided an opportunity to learn about certificate management and TLS, in addition to exploring a great case study of Kubernetes API implementations (e.g. building CRDs and providing an abstraction layer for managing certificate resources). Overall, the stability of the projects and their API stand out over time, and that builds a strong sense of trust with their adopters.

July, 2024


##### Adopter 2 - Academia/Research

Adopter 2 has been using the cert-manager project for approximately 1 year. One of the main motivations for adopting the project was to provision ACME based TLS certificates for approximately 100 services across multiple clusters. In their previous platform setup, TLS certificates were manually provisioned and updated, hence integrating cert-manager was a great way to automate this process.  In Adopter 2's view, cert-manager is the defacto product in issuing and managing certificates and they tend to choose the project with the largest install base and majority of users, as it makes information finding easier.

Their current usage of cert-manager spans more than 100 certificates in both test and production clusters, supporting internal and public-facing traffic from websites and applications. The update cadence is approximately every 6 months, and since the project has no ingress functionalities, Adopter 2 opted for a relatively conservative update schedule.

In general, Adopter 2's experience of adopting cert-manager was smooth, with a highlight on the straightforwardness of installing the Kubernetes components. One challenge Adopter 2 faced was their usage of non-public CAs (e.g. not Let's Encrypt), hence extra work was necessary for setting up authentication between the CA and cert-manager, but in general, it works fine. Throughout the project exploration, Adopter 2 used primarily the documentation from the project's website and found the annotated example configurations the most helpful. The cert-manager channel on Kubernetes Slack was also a good resource for specific questions, where the user received basic upgrade guidance. Adopter 2 also tried to look for an  IRC channel, but there is not one available at the moment.

The implementation of the project provided measurable value to Adopter 2, especially in the automation of certificate issuing.  In their older container platform, they were manually provisioning and updating TLS certificates.  This was not only time consuming, sometimes taking a couple of hours a week, but also very error-prone. They had no automated alerting for upcoming expiry hence in the past dealt with production certificates lapsed.  They also had to manually check and ensure new certificates matched the old parameters, such as key size and DNS SANs.

Adopter 2 would find valuable extra features around monitoring and alerting on certificate renewals. They are able to monitor certificates externally, but currently, there is no solution for detecting when a renewal fails, or when a certificate gets too close to the expiry date. In regards to their plans around the project, Adopter 2 is happy with current regular updates and is thinking of implementing mTLS as part of a service mesh, which would involve adding client certificates but is not an immediate priority at the time of this interview.

July, 2024


##### Adopter 3 - Software

Adopter 3 has been using cert-manager in production for approximately 4 years.  One of the key factors to use the project was the standard interface to access certificates from applications, enabling the integration with other ecosystem tools. Majority of cloud native software that requires TLS, has integrations with cert-manager, which eases the overall adoption process.

Regarding the integration experience with existing services and applications, Adopter 3 has built a custom issuer to integrate with their internal certificate manager. They also developed and contributed upstream to the cert-manager, focusing on functionalities that ease integration with user workloads and applications. Related to the update cadence, Adopter 3 revises their cert-manager services when upstream has a new release is available, which is also up to individual cluster administrators to apply these changes.

Adopter 3 has found the project documentation helpful, especially for the implementation of the of an internal issuer and KubeBuilder. Also, the interaction with maintainers through  Slack & GitHub was satisfactory.  The project adoption has provided considerable value to Adopter 3,  by standardising access to certificates across all internal  Kubernetes distributions. There is no longer a need to manually create calendar alerts and rotation of certificates, leading to a reduction in outages caused by expired certificates, a better accounting of certificate and private key material.

August, 2024
