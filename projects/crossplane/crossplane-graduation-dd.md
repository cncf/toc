# Crossplane Graduation Due Diligence

- Link to [Graduation application issue](https://github.com/cncf/toc/issues/1788)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Graduation Evaluation Summary for Crossplane

### Criteria Evaluation

TOC Members, Karena Angell and Dims, conducted the due diligence of Crossplane who applied for CNCF Graduation. TAG Leads Josh Gavant and Mario Fahlandt conducted the governance review of Crossplane. The project has completed the criteria that show its maturity at the level of Graduation. The following criteria implementations are noteworthy to call out:

$NOTABLES. 

The following actions were provided to the project that were considered blocking but since resolved:

* Vendor Neutrality: The project needs to demonstrate a stronger separation between upstream/downstream concerns and commercial vendors should only appear in an upstream context in accordance with CNCF policies.
    * The project has done considerable work to ensure a clean separation between upstream and downstream and clarified within its documentation.
    * Add Community Extension Projects list page #999
    * The TOC recommends the project meet with the end user TAB and TOC in a joint session after the next Steering Committee vote as a touch point in commitment to additional vendor neutrality initiatives put in place this year.
* Survivability: The project needs to demonstrate long term survivability and resilience to sponsoring vendors shutting down.
    * 
* Project Lifecycle Policies: There is confusion around the status and health of Crossplane ecosystem projects which negatively affects the community's ability to adopt the project.
* Project status list
  * Tracked in Provide a list of active extensions and their maturity crossplane/crossplane#6834
  * Fixed in Add Community Extension Projects list page crossplane/docs#999 and governance: community extension projects list crossplane/crossplane#6836
  * Live on https://docs.crossplane.io/latest/learn/community-extension-projects/
* CNCF project maintainers list out of date
  * Mentioned in Crossplane Governance Review for Graduation Application #1887 (comment)
  * PR opened to fix in Update Crossplane maintainers in project-maintainers.csv foundation#1143

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:

* Link release docs to release processes #1004
* Remove stale note about default registry #1002
* Update header navigation community link to commercial #995
* "community" link on docs.crossplane.io
  * Tracked in Top banner "Community" link goes to "Commercial" crossplane/docs#994
  * Fixed in Update header navigation community link to commercial crossplane/docs#995
  * Live on https://docs.crossplane.io/
* Clarify purpose of SIGs
  * Tracked in Define "SIG" in Crossplane crossplane/crossplane#6833
  * Fixed in README: clarify purpose of SIGs crossplane/crossplane#6835
  * Live on https://github.com/crossplane/crossplane?tab=readme-ov-file#special-interest-groups-sig


### Adoption Evaluation

The adopter interviews reflect a project in use for the level which the project applied.

Adopters of the Crossplane project that have chosen to share their adoption story publicly can be found in the [ADOPTERS.md](https://github.com/crossplane/crossplane/blob/main/ADOPTERS.md) file in the main Crossplane repository. Currently, there are over 60 public adopters of the project, and there are more that are willing to share their story with the TOC privately. Some notable Crossplane public adopters include Nike, Autodesk, Grafana, NASA Science Cloud, Elastic, Akamai, SAP, IBM, VMWare Tanzu, and Nokia.

The interviews show that Crossplane is a complex project a steep learning curve. However, this initial adoption investment allows for a unified control plane that transforms a Kubernetes implementation from a simple container orchestrator into a control plane for most cloud infrastructure. This approach consolidates infrastructure orchestration within the familiar Kubernetes API and eliminates the need for separate state management files by treating cloud resources as native Kubernetes objects.

The abstractions provided by the Crossplane project allow these adopter platform teams to create simple and secure infrastructure templates. These abstractions enforce organizational best practices by default, dramatically simplifying the resource provisioning process for developers within the adopter organizations.
This shift toward a centralized, Kubernetes-driven model has yielded specific, measurable outcomes in both development speed and operational security.

One of the most significant, concrete improvements experienced by adopters is the reduction of time required to provision cloud resources. The move to a GitOps-centric workflow with Crossplane led to considerable accelerations in deployment velocity: processes that once took months could now be completed in minutes.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation.

### Criteria

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**
  - This was completed and occurred on 07-02-2024, and can be discovered in [this comment](https://github.com/cncf/toc/pull/1254#issuecomment-1978304903).

- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

Crossplane operates according to well defined vendor-neutral governance in https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md, and project communication, messaging, and collaboration is demonstrated as vendor-neutral.

The official project charter states that the project is vendor-neutral as well: https://github.com/crossplane/crossplane/blob/main/CHARTER.md#what-crossplane-is

The Crossplane project made numerous changes to improve the project's vendor neutrality. Previously blocking this application, the project has now met the criteria for this requirement:

* Established a vendor neutral registry for Crossplane extension packages (xpkg.crossplane.io)
* Updated all project documentation (e.g., https://docs.crossplane.io/) to refer only to vendor neutral resources and artifacts
* Defined clear governance, policies, and lifecycle for community extension projects, which includes policies around vendor neutrality so that it is enforced consistently throughout the project ecosystem
* Migrated critical project release artifacts to vendor neutral CNCF owned accounts, improving project sustainability in the face of any vendor shutting down
* All vendor specific code has been removed from the upstream code base

The TOC recommends the project meet with the end user TAB and TOC in a joint session after the next Steering Committee vote as a touch point in commitment to additional vendor neutrality initiatives put in place this year.

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
   - [x] Met during the project's original proposal PR [Crossplane Graduation Proposal #1254](https://github.com/cncf/toc/pull/1254) on 05-Feb-2024, the second application after the format change on 13-Aug-2024, and with the TOC sponsors during due diligence at Kubecon London 31-Mar-2025.

The Crossplane project has reviewed and understands the expectations as it has continued to move forward through the maturity levels as described in the process [README](https://github.com/cncf/toc/blob/main/process/README.md) and [graduation criteria.](https://github.com/cncf/toc/blob/main/.github/ISSUE_TEMPLATE/template-graduation-application.md)

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The Crossplane project is thoroughly documented, most notably the [user documentation](https://docs.crossplane.io/v2.0/), [contributing guide](https://docs.crossplane.io/contribute/), [CLI](https://docs.crossplane.io/v2.0/cli/) and [API](https://docs.crossplane.io/v2.0/api/) reference documentation, and [build guide](https://docs.crossplane.io/v2.0/guides/install-from-source/). The project itself is complex and has a steep learning curve, as evidenced during the adopter interviews. However, the documentation itself is appropriate for a complex project such as Crossplane that meets the criteria for Graduation.

## Governance and Maintainers

Note: this section is augmented by the Governance Review completed by Mario Fahlandt and Josh Gavant from the TOC Project Reviews subproject.

### Suggested

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The project governance has undergone a few revisions in its history since the project's creation. These commits/updates can be found in the git history at [https://github.com/crossplane/crossplane/commits/main/GOVERNANCE.md](https://github.com/crossplane/crossplane/commits/main/GOVERNANCE.md). 

The most recent important governance update was ratified on 21-Feb-2025 which established the policies for Community Extension Projects: [crossplane/crossplane#6290](https://github.com/crossplane/crossplane/pull/6290). This was in direct response to feedback on vendor neutrality provided by the TOC sponsors during graduation due diligence for [#1397](https://github.com/cncf/toc/issues/1397).

### Required

- [x] **Clear and discoverable project governance documentation.**

The Crossplane project has had governance in place since entry into the CNCF Sandbox, which can be found in the main repo’s GOVERNANCE.md file. All aspects of the life cycle for Crossplane leadership positions, including the steering committee and repository maintainers (committers) are described in detail within this governance document. The steering committee members, currently from Upbound, Apple, and Nokia, can also be found in the project governance. Repository maintainers can be found in the OWNERS.md file of each separate Crossplane repository that make up the project.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

The governance is up to date with the latest iteration of the steering committee membership, which was last updated after the most recent election occurred in February 2025: [crossplane/crossplane#6304](https://github.com/crossplane/crossplane/pull/6304). All processes for approvals, maintainers, conflict resolution, community extension projects, etc. are defined and up to date in this governance document.

All meetings within the Crossplane community and ecosystem are tracked in the [community calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/crossplane). This calendar as well as other ways to get involved are highlighted prominently in the project's main [README](https://github.com/crossplane/crossplane?tab=readme-ov-file#get-involved).

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

The governance has a "maximum representation" section that outlines how vendor neutrality is enforced over the lifetime of the project and leadership elections: https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#maximum-representation

Community Extension Projects have clearly defined policies in the project governance that promote vendor neutrality as well, for example the strict requirement that all projects must publish artifacts to the vendor neutral community registry xpkg.crossplane.io: https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#policies-for-community-extension-projects

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

Changes to governance has a clearly defined process in https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#updating-the-governance.

Project leadership (steering committee) election process is defined in https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#election-process. This process was followed for the most recent 2025 election.

Process for how each individual repository under the crossplane organization(s) are maintained can be found in https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#repository-governance. Additionally, policies and guidelines for Community Extension Projects is clearly defined in https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#policies-for-community-extension-projects.

Contribution acceptance is augmented by the contributing guide with https://github.com/crossplane/crossplane/tree/main/contributing#contributing-code and https://github.com/crossplane/crossplane/tree/main/contributing#code-review-process.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

* Repository maintainers are added by: https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#becoming-a-maintainer
* Repository maintainers are removed by: https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#removing-a-maintainer
* Steering committee members are elected by: https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#election-process

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

* Repository maintainers are added by: https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#becoming-a-maintainer
* Repository maintainers are removed by: https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#removing-a-maintainer

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Using the an example repository maintainers (OWNERS.md), the history of these files demonstrates the maintainer membership changes over time, with both additions and removals (or movement to emeritus status):

https://github.com/crossplane/crossplane/commits/main/OWNERS.md
https://github.com/crossplane/crossplane-runtime/commits/main/OWNERS.md
https://github.com/crossplane-contrib/provider-sql/commits/master/OWNERS.md
https://github.com/crossplane-contrib/provider-ansible/commits/main/OWNERS.md

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

The steering committee membership and details can be found in https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#initial-steering-committee, and contact info for the committee as a whole is provided.

The maintainers of each repository in the crossplane and crossplane-contrib organizations are listed in the OWNERS.md file in each individual repository. For example:

core Crossplane: https://github.com/crossplane/crossplane/blob/main/OWNERS.md
crossplane-runtime: https://github.com/crossplane/crossplane-runtime/blob/main/OWNERS.md
provider-sql: https://github.com/crossplane-contrib/provider-sql/blob/master/OWNERS.md
provider-palette: https://github.com/crossplane-contrib/provider-palette/blob/main/OWNERS.md
Community Extension Projects must keep these lists up to date, as part of the defined policies and lifecycle expectations.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

Across the entire Crossplane project, there are [183 different companies](https://crossplane.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Since%20moving%20to%20incubating%20state&var-metric=committers) that have committers (individuals with write permissions) on at least one repository. 

The documented maintainers for the core and extension repositories also match the activity of active maintainers for the project at the level applied.

- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

Across the entire Crossplane project, there are [183 different companies](https://crossplane.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Since%20moving%20to%20incubating%20state&var-metric=committers) that have committers (individuals with write permissions) on at least one repository, which demonstrates organizational diversity.

Also, the steering committee for the Crossplane project is composed of individuals from 3 separate organizations: Apple, Nokia, and Upbound.

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

Yes, OWNERS.md files in each Crossplane project repository should reflect the documented maintainer roles defined in the governance. For example, https://github.com/crossplane/crossplane/blob/main/OWNERS.md lists current and emeritus maintainers for that repo. The ownership being up to date is an enforced policy for all Community Extension Projects.

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

Crossplane project and community adhere to the CNCF Code of Conduct, e.g., https://github.com/crossplane/crossplane/blob/main/CODE_OF_CONDUCT.md.

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

The CNCF Code of Conduct is linked from the root of the core Crossplane repository https://github.com/crossplane/crossplane/blob/main/CODE_OF_CONDUCT.md, where GitHub automatically displays it from its "Code of conduct" tab.

- [x] **All subprojects, if any, are listed.**

Crossplane does not have formally defined "subprojects", but all repositories under the crossplane and crossplane-contrib repository adhere to the well defined governance.

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

As the project does not have subprojects defined, this requirement does not apply.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

Contributor roles fall into 3 tiers: member, maintainer, and steering committee. The roles and expectations are described in:

* [new member process](https://github.com/crossplane/org/blob/main/processes/new-member.md)
* [becoming a maintainer](https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#becoming-a-maintainer)
* [steering committee elections](https://github.com/crossplane/crossplane/blob/main/GOVERNANCE.md#election-process)

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

All repositories in the Crossplane project accept issues and changes from the community through the standard Github workflows:

Issues: https://github.com/crossplane/crossplane/issues
Changes (PRs): https://github.com/crossplane/crossplane/pulls

Both [issues](https://github.com/crossplane/crossplane/tree/main/.github/ISSUE_TEMPLATE) and [PRs](https://github.com/crossplane/crossplane/blob/main/.github/PULL_REQUEST_TEMPLATE.md) have templates to standardize and guide the contributor experience.

The [Contributing guide](https://github.com/crossplane/crossplane/tree/main/contributing) also describes how changes are accepted, what the contributor can expect to experience, and tips for making a successful contribution.

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

N/A

### Required

- [ ] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  - _If applicable_ A General Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document what the project does, and why it does it - including viable cloud native use cases. _This requirement may also be satisfied by completing a General Technical Review._**
  - _If applicable_ a General Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

- [ ] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Roadmap change process is documented.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.  *This requirement may also be satisfied by completing a General Technical Review and capturing the output in the project's documentation.***
  -  _If applicable_ a general Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:** 

  - [ ] Release expectations (scheduled or based on feature implementation)
  - [ ] Tagging as stable, unstable, and security related releases
  - [ ] Information on branch and tag strategies
  - [ ] Branch and platform support and length of support
  - [ ] Artifacts included in the release.
  - Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.

<!-- (TOC Evaluation goes here) --> 

- [ ] **History of regular, quality releases.**

<!-- (TOC Evaluation goes here) --> 

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [ ] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **Third Party Security Review.**

  - [ ] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

<!-- (TOC Evaluation goes here) --> 

- [ ] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) --> 

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

<!-- (TOC Evaluation goes here) --> 

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

<!-- (TOC Evaluation goes here) --> 

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [x] **TOC verification of adopters.**

<!-- (TOC Evaluation goes here) --> 

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The Crossplane project has integrations with many other CNCF and non-CNCF projects that are not only clearly documented, they are also widely publicized to increase adoptability:

* [Kubernetes](https://docs.crossplane.io/v2.0/get-started/install/) is extended by Crossplane to connect it to external, non-Kubernetes resources, and allows platform teams to build custom Kubernetes APIs to consume those resources.
* [Helm](https://docs.crossplane.io/v2.0/get-started/install/) is the main way to install Crossplane into a control plane.
* [ArgoCD](https://docs.crossplane.io/latest/guides/crossplane-with-argo-cd/#configuring-argo-cd-with-crossplane) is used frequently to sync Crossplane resources and definitions from a Git repository to the control plane to enable GitOps workflows.
* [Flux](https://www.cncf.io/blog/2022/07/26/how-to-apply-gitops-to-everything-with-crossplane-and-flux/) also enables GitOps workflows for Crossplane resources.
* [gRPC](https://docs.crossplane.io/latest/guides/write-a-composition-function-in-python/) powers the communication between Crossplane's core composition engine (client) and the Functions (server) within a user defined composition pipeline.
* [Prometheus](https://docs.crossplane.io/latest/guides/metrics/) metrics provide observability on Crossplane's internal behavior/health as well as statistics about the resources that Crossplane is managing.
* [Harbor](https://docs.crossplane.io/latest/guides/install-from-source/) can serve as a container registry for Crossplane packages.
* [Open Policy Agent](https://blog.crossplane.io/crossplane-v0-10-compose-and-publish-your-own-infrastructure-crds-velero-backup-restore-compatibility-and-more/) is commonly used with Crossplane to enforce organizational policy on Crossplane resources.
* [Kyverno](https://www.cncf.io/online-programs/cncf-on-demand-webinar-securely-provisioning-cloud-resources-using-kyverno-and-crossplane/) also enforces policy to ensure secure provisioning of resources with Crossplane.
* [ArtifactHub](https://artifacthub.io/packages/helm/crossplane/crossplane) indexes all versions of Crossplane's main Helm chart for installation into control planes.
* [Backstage](https://2022.platformcon.com/talk/how-to-build-an-idp-with-backstage-crossplane-and-argo-cd) is often used as a developer portal on top of Crossplane to offer a comprehensive Internal Developer Platform.
* [Dapr](https://blog.crossplane.io/crossplane-and-dapr/) and Crossplane work well together to expose resources provisioned by Crossplane for consumption by developers with Dapr.
* [KubeVela](https://kubevela.io/docs/end-user/components/cloud-services/provision-cloud-resources-by-crossplane/) supports Crossplane as an add-on to provision resources.
* [KCL](https://blog.crossplane.io/function-kcl/) has quickly become one of the favored languages to write Crossplane composition logic via function-kcl.
* [Velero](https://docs.crossplane.io/latest/guides/disaster-recovery/) can backup and restore the resources of Crossplane to perform disaster recovery.
* [K8GB](https://www.k8gb.io/examples/crossplane/globalapp/) documents how to use Crossplane to build resilient, scalable multicluster environments integrated with k8gb for DNS-based failover.

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
