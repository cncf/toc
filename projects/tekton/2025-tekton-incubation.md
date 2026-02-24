# Tekton Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1310)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Tekton

### Criteria Evaluation

Jeremy Rickard (@jeremyrickard) and Chad Beaudin (@chadbeaudin) conducted the due diligence of Tekton who applied for incubation. Tekton is currently a graduated project in the Continuous Delivery Foundation (CDF). Tekton is applying to join the Cloud Native Computing Foundation (CNCF) at the incubating level, so the following actions will be required upon onboarding to the CNCF:

- [ ] Adoption of the CNCF Code of Conduct 

This would be a normal onboarding activity, so we have not created a tracking issue for this. The project should complete the normal sandbox onboarding steps as part of their acceptance to the CNCF. 

Additionally, it would be good for the project to update some documentation around security processes, including how members of the security response team are selected and include a SECURITY.md file. This is also part of the sandbox onboarding process.

The project has completed the criteria that show its maturity at the incubating level.
 

### Adoption Evaluation

The adopter interviews reflect a project that exceeds the expectations for a project at the incubating level. The project is used by end users across a number of industries and is also a key dependency of Shipwright project, a CNCF sandbox project.

Adopters highlighted the Kubernetes native design of the project and the ability to adopt GitOps models for management of pipelines and to shorten onboarding times for new projects through high levels of reuse and the ability to integrate with other cloud native tools via its Kubernetes native design. 

In general, adopters felt that the project is quite mature. Adopters identified areas for improvement in usability and documentation. 

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation. 

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**
  -  The Tekton maintainers presented an overview of the project at the [TAG App Delivery meeting on May 15th, 2024](https://youtu.be/8E316JEe3HY?si=V1osmlvTr-WLFtSG&t=836).

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The Tekton website and project repositories are vendor neutral.

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**	

Tekton is not currently a CNCF project. The project is applying directly to incubation status. The project is currently hosted by the [Continuous Delivery Foundation](https://cd.foundation) and has reached the graduated maturity within that foundation. However, the project currently exceeds the expectations for a sandbox project within the CNCF.   

- [x] **Due Diligence Review.**

The initial application was initially submitted May 01, 2024.

An initial set of resources were provided to support the DD with the project's [incubation application issue](https://github.com/cncf/toc/issues/1310).

A due diligence review was performed at the end of 2025. The maintainers of the project were responsive and provided additional information as needed. 

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The Tekton project hosts documentation at https://tekton.dev/docs/. The documentation includes [installation instructions](https://tekton.dev/docs/installation/pipelines/) for generic Kubernetes clusters, along with Google Cloud and OpenShift specific instructions.

The project documentation also includes basic [how to guides](https://tekton.dev/docs/how-to-guides/) that cover core concepts like [cloning a Git repository](https://tekton.dev/docs/how-to-guides/clone-repository/) and [building images](https://tekton.dev/docs/how-to-guides/kaniko-build-push/). More extensive documentation is provided for other concepts. The project also provides [examples](https://github.com/tektoncd/pipeline/tree/main/examples). 

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The initial governance documentation was added in [March 2019](https://github.com/tektoncd/community/commit/f78c3f7d7ef9233a16980803d70650619b92abc5), when Tekton was part of the Knative project. It has been updated and refined several times as the project matured and separated from the Knative project, including creation of a bootstrap governing board and ongoing updates to the Tekton Governance Committee. The project has further demonstrated adherence to its governance through recurring elections. 

- [x] **Clear and discoverable project governance documentation.**

The project governance documentation can be found at https://github.com/tektoncd/community/blob/main/governance.md. A clear contributor ladder is defined in [contributor ladder](https://github.com/tektoncd/community/blob/main/process/contributor-ladder.md)

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

The governance documentation has been regularly updated to reflect leadership positions and reflects current practices. 

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

The project governance documentation specifies [maximum representation](https://github.com/tektoncd/community/blob/main/governance.md#maximum-representation) for membership of the governing committee and explicitly states that governing committee members must represent a cross-section of interests, not just one company. 

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

The Tekton [contributor ladder](https://github.com/tektoncd/community/blob/main/process/contributor-ladder.md) specifies how the project accepts contributions. The [process](https://github.com/tektoncd/community/tree/main/process) folder of the community repository provides additional information on the process surrounding contribution acceptance. 

The project uses a [mailing list](tekton-governance@googlegroups.com) for making governance decisions and taking votes, while leadership changes are handled using [CIVS](https://civs.cs.cornell.edu/). 


- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

Roles are well defined for contributor ladder and for membership on the governing committee, along with onboarding and offboarding checklists. 

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

The governance document specifies [maintainer lifecycle](https://github.com/tektoncd/community/blob/main/governance.md#changes-to-governing-board) and contains a list of emeritus members. The project governance documentation also contains a process for [involuntary removal](https://github.com/tektoncd/community/blob/main/process/contributor-ladder.md#involuntary-removal-or-demotion).

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Examples of the lifecycle can be seen in pull requests to the community repo, such as [this PR](https://github.com/tektoncd/community/pull/1195) from April 10th, 2025 reflecting the 2025 election results. This is also demonstrated in ongoing pull requests to project specific repositories, like this PR moving [@waveywaves to maintainer](https://github.com/tektoncd/pipeline/pull/8743) from May 6th, 2025. 

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

The governance documentation includes a document on [projects](https://github.com/tektoncd/community/blob/main/process/projects.md). Tekton has experimental and non-experimental projects. Non-experimental projects have a common set of requirements. 

Project leadership is defined in an OWNERS file for each project. For example, [pipelines OWNERS](https://github.com/tektoncd/pipeline/blob/main/OWNERS_ALIASES).

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

The current governing committee members, including contact information and affiliation, is defined in [community/governance.md](https://github.com/tektoncd/community/blob/main/governance.md#current-members).

Maintainers for subprojects and area of responsibility can be found in [community/org/org.yml](https://github.com/tektoncd/community/blob/main/org/org.yaml). 

Affiliation for maintainers outside the governing committee is not recorded, but is available from [devstats](https://tekton.devstats.cd.foundation/d/8/dashboards?orgId=1&refresh=15m).  

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

The project has a large number of active maintainers. This was confirmed by comparing the [community/org/org.yml](https://github.com/tektoncd/community/blob/main/org/org.yaml) with the contributions in [devstats](https://tekton.devstats.cd.foundation/d/9/developer-activity-counts-by-repository-group-table?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-country_name=All).

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

Tekton uses [Prow](https://github.com/kubernetes-sigs/prow) and OWNER files for each repository. Reviewer and Approver status match what is in the contributor ladder. 

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

The Tekton project has currently adopted a code of conduct based on the [Contributor Covenant](https://www.contributor-covenant.org/version/1/4/code-of-conduct/). Upon acceptance to the CNCF, the project will need to adopt the CNCF Code of Conduct.   


- [ ] **CNCF Code of Conduct is cross-linked from other governance documents.**

The code of conduct is referenced from other governing documents and project contribution README files, but is the Code of Conduct adopted prior to joining the CNCF.

- [x] **All subprojects, if any, are listed.**

Tekton documents subprojects as components on the project [documentation website](https://tekton.dev/docs/concepts/overview/#:~:text=Tekton%20Pipelines%20is%20the%20foundation,artifacts%20built%20with%20Tekton%20Pipelines).


## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

A [contribution ladder](https://github.com/tektoncd/community/blob/e87982194801c43d9136bd53b67ac3b0fea7b65d/process/contributor-ladder.md) is clearly defined with the following roles:

* Community Participant
* Contributor
* Organization Member
* Reviewer
* Maintainer
* Governance Committee Member

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

The project has a well defined [process](https://github.com/tektoncd/community/tree/e87982194801c43d9136bd53b67ac3b0fea7b65d/process#finding-something-to-work-on) for contributing changes, including proposing new [enhancement proposals](https://github.com/tektoncd/community/blob/e87982194801c43d9136bd53b67ac3b0fea7b65d/process/tep-process.md). Adopter interviews confirm that the enhancement proposal process is followed.  

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

The project lists multiple [communication channels](https://github.com/tektoncd/community/blob/e87982194801c43d9136bd53b67ac3b0fea7b65d/contact.md) including email and Slack.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

* [tekton slack](https://tektoncd.slack.com/)
* [tekton-dev@ maintainer mailing list](https://groups.google.com/forum/#!forum/tekton-dev)
* [tekton-users@ end user mailing list](https://groups.google.com/forum/#!forum/tekton-users)

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

The project is currently not part of the CNCF at the time of evaluation. This will need to be updated after acceptance onto the CNCF.  

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

Each subproject in the repository has a CONTRIBUTING.md file. The Tekton Pipelines [CONTRIBUTING.md](https://github.com/tektoncd/pipeline/blob/main/CONTRIBUTING.md) is a good example of how the project documents how to contribute. 

- [x] **Demonstrate contributor activity and recruitment.**

The Tekton Devstats dashboard shows more than 20 contributors having 20 or more contributions for [the last 6 months](https://tekton.devstats.cd.foundation/d/9/developer-activity-counts-by-repository-group-table?orgId=1&var-period_name=Last%206%20months&var-metric=contributions&var-repogroup_name=All&var-country_name=All).

LFX Insights also has some data for the [Tekton project](https://insights.linuxfoundation.org/project/tekton?timeRange=past365days&start=2025-02-18&end=2026-02-18). Most contributions come from IBM/Redhat affiliated contributors, however unaffiliated contributors represent 14% of contributions for the previous year. There have been numerous small contributions from a large number of companies, indicating that there is ongoing active contributions outside of IBM/Redhat. LFX Insights also indicates a geographic distribution of commits, with at least 25% occurring outside of the United States.

The [community README](https://github.com/tektoncd/community?tab=readme-ov-file#want-to-get-involved) details way to get involved in the project and the contributor ladder has a low barrier to entry, which indicate efforts to grow the contributor base.


## Engineering Principles

### Suggested

- [x] **Roadmap change process is documented.**

The Tekton subprojects utilize Github projects to track roadmaps. The pipeline roadmap is the most up-to-date. The project's [roadmap.md](https://github.com/tektoncd/community/blob/main/roadmap.md#roadmap) file indicates that the project boards are automatically updated through the use of a label. 

- [x] **History of regular, quality releases.**

The Tekton project has demonstrated a history of [regular releases](https://github.com/tektoncd/pipeline/blob/main/releases.md), dating back to 2019, for the Pipelines subproject. Releases include detailed [release notes](https://github.com/tektoncd/pipeline/releases/tag/v1.2.0) and include information on how to verify release attestations.

### Required 

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  -  _If applicable_ a general Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

The Tekton project has a well defined [mission and vision](https://github.com/tektoncd/community/blob/main/roadmap.md#mission-and-vision) document.  

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

The Tekton project maintains an [overview](https://tekton.dev/docs/concepts/overview/) of what Tekton does, what users can do with Tekton, and why. Additionally, they document [user profiles](https://github.com/tektoncd/community/blob/main/user-profiles.md) that are served by the project. 

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The Tekton project maintains a [public roadmap](https://github.com/orgs/tektoncd/projects/26) in a Github project that is automatically populated with issues labeled with `area/roadmap`. 

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**
  -  _If applicable_ a general Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

The core use case of Tekton is providing CI/CD in a Kubernetes native way. The Tekton documentation has an [overview](https://tekton.dev/docs/concepts/overview/) section that explains key concepts. Additionally, there is a [concept model](https://tekton.dev/docs/concepts/concept-model/) page that provides more context on how Tekton implements CI/CD workflows in a Kubernetes native way. 

- [x] **Document the project's release process.**

The Tekton project utilizes a mix of Github Actions and Tekton pipelines to release. The repository contains comprehensive [release process documentation](https://github.com/tektoncd/pipeline/tree/main/tekton#create-an-official-release). 

## Security

### Suggested

N/A

### Required

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance.

- [x] **Clearly defined and discoverable process to report security issues.**

The Tekton project documents a process for [reporting security issues](https://github.com/tektoncd/community?tab=security-ov-file#).

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

The [contributor ladder](https://github.com/tektoncd/community/blob/83f4a856fe227ade7c885825f5f850f41e9dfab7/process/contributor-ladder.md?plain=1#L92) documentation specifies that all contributors must have two-factor authentication enabled on their Github account. The project uses Prow for enforcement of OWNER files for a separation of reviewer and approver roles. 

- [x] **Document assignment of security response roles and how reports are handled.**

The project does not currently specify assignment of security response roles. The project has historically taken volunteers that have appropriate background and skills to handle security response. The project documents how to report a security incident [on Github](https://github.com/tektoncd/community/security). They utilize a Google group for reporting of security vulnerabilities and specify response timeframes.  

- [x] **Document Security Self-Assessment.**

The Tekton project completed a self-assessment while in the CDF, which can be located at https://hackmd.io/FlVw3hJxSyWuVmUy66Qs8g.

Additionally, the Tekton project also engaged in a security audit while in the CDF. A summary of this audit can be found at https://cd.foundation/blog/2022/08/26/tekton-security-review-completed/. 

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The Tekton project has achieved the OpenSSF Best Practices [passing badge](https://www.bestpractices.dev/en/projects/4020/passing).

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

The Tekton project maintains a list of [adopters](https://github.com/tektoncd/community/blob/main/adopters.md) in their community repository. 

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The adopter list provided in the Tekton community repository has eight end-user adopters. 

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [x] **TOC verification of adopters.**

The Tekton project provided a list of adopters to the TOC in support of the incubation due diligence review. Three adopters were interviewed.  

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The Tekton adopter file includes several open source projects, including projects in the CNCF that provide integration with Tekton. The project additionally maintains a list of [Tekton Friends](https://github.com/tektoncd/friends/) that highlight integrations with other projects. Additionally, the project documentation includes information about compatibility with supply chain security projects, like cosign and metrics projects like Prometheus. 

#### Adoption

##### Adopter 1 - Shipwright (CNCF Sandbox Project)

October 2025

[Shipwright](https://github.com/shipwright-io/build) is a CNCF project that enables building container images directly on Kubernetes. That goal is to build the container images where you are using them, so you don't need local Docker builds everywhere on developer machines. 

Shipwright, from its inception, built on core Tekton features like `TaskRuns` and `Parameters`. The Shipwright project made the decision to build on Tekton in early 2020 based on the need for a Kubernetes native way to sequentially execute tasks or build steps. Tekton was really the only project at this time that satisfied the requirements. The maintainers estimated that adopting Tekton saved several developer months of time and accelerated the project. They have found Tekton to be very mature and stable over the last five years, particularly as the various APIs in the project matured. 

The Shipwright maintainers did see room for improvement in the project documentation, including enhancement of the Tekton custom resources to better support `kubectl explain` for those adopting Tekton. 

Reference: [Tekton Adopter Interview - Shipwright](/projects/tekton/tekton-adopter-interview-shipwright.md)


##### Adopter 2 - Travel and Hospitality

October 2025

Adopter 2 operates a devops platform for all of their developers to use. They were formerly a JenkinsX user, but encountered issues scaling as their IT organization continued to grow. They also experienced pain and friction around onboarding new projects and had a centrally managed support system that was seeing an increasing number of requests for help. To address these concerns, they wanted to adopt an open source technology that would allow them to provide declarative pipelines as code in order to shift to a self-service onboarding model. Tekton fit this model well. They have now been using Tekton in production for several years and perform over 70,000 deployments per year using Tekton.

This adopter identified the native integration with Kubernetes and the ability to run the pipelines within Kubernetes on vetted infrastructure as key strengths.  

The adopter feels that there could be more awareness of the project and that it potentially lacks marketing. The adopter thinks that there could be more positioning within the CI/CD space and is optimistic that the project could grow its community after joining the CNCF. They also feel that the project should focus on providing a better UI experience to increase adoption. 

Reference: [Tekton Adopter Interview - Adopter 2](/projects/tekton/tekton-adopter-interview-adopter2.md)


##### Adopter 3 

December 2025

Adopter 3 is a video streaming and ad sales platform. They were also previously a JenkinsX user that migrated to using Tekton natively. Tekton is used in production by the adopter, as well as in staging environments. Tekton handles all pull requests within Adopter 3 and creates 40,000 pods a month in CI/CD pipelines.

Overall, the adopter cited ease of integration, ease of customization and native Kubernetes integration as key strengths. They also cited the openness of project to accept contributions as another strength, and plan to contribute a new feature to the project in 2026. 

The adopter identified the ability to persist log information outside the cluster as an area that could be improved. At scale, they find it difficult to manage long term storage of pipeline results as they are tied to pod lifecycle and at scale tend to be lost as pods are deleted.

Reference: [Tekton Adopter Interview - Adopter 3](/projects/tekton/tekton-adopter-interview-adopter3.md)