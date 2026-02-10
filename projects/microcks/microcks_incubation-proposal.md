# Microcks Incubation Due Diligence

- [Microcks Incubation application issue](https://github.com/cncf/toc/issues/1552)

<!-- This template provides the TOC with the outline for completing due diligence of a project to move levels. This universal template is designed to capture all criteria so the TOC may ensure prior level criteria do not regress. As part of completing the due diligence, the TOC member should update the template to convey the level the project applied for the criteria by bolding the level indicated where the criteria is relevant. -->

## Incubation Evaluation Summary for Microcks

### Criteria Evaluation

Katie Gamanji conducted the due diligence of Microcks who applied for Incubation. The project [has/has not] completed the criteria that show its maturity at Incubation. The following criteria implementations are noteworthy to call out:
- Microcks has an outstanding  list of adopters [ADOPTERS.md](https://github.com/microcks/.github/blob/main/ADOPTERS.md).
- ... $NOTABLES. 

The following actions were provided to the project that were considered blocking but since resolved:
- For the main project, update the list of maintainers to include active contributors and develop a plan for onboarding new members.  
- For subprojects, identify reasons for stalled contributions and encourage activity or implement archiving initiatives as needed. 
- Onboard at least 1 maintainer to ensure the voting clause on project maintainers and deliverables is fulfilled. The other option is to update the voting guidelines. 
- Clearly outline the project goals and objectives and/or out-of-scope items. This can be solved by completing the GTR or separately in the project documentation. Similar example already available as part of the security [self-assessment](https://github.com/yada/tag-security/blob/main/community/assessments/projects/microcks/self-assessment.md#goals).
- Provide timelines for release milestone.
- Provide clear guidelines on how to cut/deploy a release, including cadence, versioning, etc., under a Release.md or similar file.

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:
- The project presented to TAG App Delivery last year, hence their feedback and review are relevant as it's pretty recent. The extra recommendation is to liaise with the new TAGs and present to TAG Workloads Foundation and TAG Developer Experience. The other option is to complete the self-assessment [GTR](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/general-technical-questions.md).
- List all available subprojects with their respective code owners (if applicable). This will also help the discoverability of ongoing initiatives related to Microcks.

### Adoption Evaluation

_The adopter interviews reflect a project [in use/too early] for the level which the project applied. They show ... $INTERVIEWSUMMARY._

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation.

## Application Process Principles

### Suggested

N/A

### Required

- [x] **Engage with the domain specific TAG(s) to increase awareness through a presentation or completing a General Technical Review.**

<!-- (TOC Evaluation goes here) --> 
The project presented to the former TAG App Delivery on February 19, 2025, during the TAG General Meeting. The recording is available on the [TAG YouTube channel](https://www.youtube.com/watch?v=UmGNJQUMlfo&t=856s). The presented deck can be found [here](https://docs.google.com/presentation/d/1dQpmqwOKFroAHPzhAzSXip1AnM96WUH-9UrVNiBgtP0/edit?slide=id.g244de45f28a_0_13#slide=id.g244de45f28a_0_13). The feedback from the TAG members was positive, with no objections or specific recommendations.

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

<!-- (TOC Evaluation goes here) --> 
The project fulfils the vendor-neutral criteria. The main communication channel are [Slack](https://cloud-native.slack.com/archives/C05BYHW1TNJ) and [Discord](https://discord.com/invite/JA4rbcGzk7). Community meetings are held on [Zoom](https://github.com/microcks/community/blob/main/JOIN-OUR-MEETINGS.md), bi-weekly with US and APAC friendly timezones, which are showcased on the [CNCF calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/microcks?view=month). Users are also capable of join [GitHub Discussions](https://github.com/orgs/microcks/discussions) to follow up on the latest topics of interest for the project. 

The Microcks [website](https://microcks.io/) also follows the [CNCF guidelines](https://github.com/cncf/foundation/blob/main/website-guidelines.md) to ensure neutrality, and offers a variety of social media platforms: [LinkedIn](https://www.linkedin.com/company/microcks/), [X](https://twitter.com/microcksio), [BlueSky](https://bsky.app/profile/microcks.io), [Mastodon](https://mastodon.social/@microcksio), [YouTube](https://www.youtube.com/c/Microcks), and other channels do not endorse any particular vendor.

- [x] **Review and acknowledgement of expectations for [Sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		

<!-- (TOC Evaluation goes here) --> 
Met during Project's application on June 13, 2023, with its review available [on YouTube](https://www.youtube.com/watch?v=C4bU6A3RToc&t=2777s). 

Microcks Sandbox application can be found [[Sandbox] Microcks #37](https://github.com/cncf/sandbox/issues/37). The project maintainers have understood the expectations for Sandbox projects and requirements for moving forward through the CNCF incubation level.


- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

<!-- (TOC Evaluation goes here) --> 
The project documentation follows the [Diátaxis methodology](https://diataxis.fr/), and in 2024, maintainers undertook a six-month complete refactoring to enhance overall clarity and structure. The project maintainers have added the following additional documentation as follows.

- [Installation documentation](https://microcks.io/documentation/guides/installation/)
- [User documentation](https://microcks.io/documentation/guides/usage/)
- [Architecture documentation](https://microcks.io/documentation/overview/main-concepts/)
- [Blog posts from adopters and contributors](https://microcks.io/blog/)
- [Demos](https://github.com/orgs/microcks/repositories?q=demo)

Additionally, there are a number of guides and blogs from community members on technical and high-level topics, such as [How Microcks fit to unify inner and outer loops in cloud native](https://www.linkedin.com/pulse/how-microcks-fit-unify-inner-outer-loops-cloud-native-kheddache/), [Traefik API Mocking On-Premises Setup](https://doc.traefik.io/traefik-hub/api-mocking/on-premises-setup), and [Bump.sh Mocking with Microcks](https://docs.bump.sh/guides/bump-sh-tutorials/mocking-with-microcks/).

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [ ] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

<!-- (TOC Evaluation goes here) --> 


- [ ] **Clear and discoverable project governance documentation.**

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

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

<!-- (TOC Evaluation goes here) --> 
The current list of maintainers and subproject owners are listed in [MAINTAINERS.md](https://github.com/microcks/.github/blob/main/MAINTAINERS.md). 


- [ ] **A number of active maintainers which is appropriate to the size and scope of the project.**

<!-- (TOC Evaluation goes here) --> 
Currently there are 2 maintainers listed. Most contributions are made these maintainers. For instance, [69% of contributions](https://insights.linuxfoundation.org/project/microcks/contributors) were made by 2 people. There are also community members that have contributed significantly to the project and currently hold the second place in terms of overall contributions, however not listed as a contributor. 

For subprojects, some of the listed code owners for subprojects, seem to not have contributed or have a low-activity in the last year.

**Recommendation**: 
- For the main project, update the list of maintainers to include active contributors and develop a plan for onboarding new members.  
- For subprojects, identify reasons for stalled contributions and encourage activity or archive initiatives as needed. 

- [x] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

<!-- (TOC Evaluation goes here) --> 
This is implemented with the use of `CODEOWNERS` files, such as for the [main repository](https://github.com/microcks/microcks/blob/23aac2b52c5349947583672aefdc0d2b1fe2d8ce/CODEOWNERS).

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

<!-- (TOC Evaluation goes here) --> 
Microcks project adheres to the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md), aligning with the values of collaboration, transparency, and inclusivity that define open source.

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

<!-- (TOC Evaluation goes here) --> 
Yes, the CNCF Code of Conduct is cross-linked from the main governance definition.

- [ ] **All subprojects, if any, are listed.**

<!-- (TOC Evaluation goes here) --> 
The following subprojects are listed:
  | Name                          | URL                                                                 |
  |-------------------------------|---------------------------------------------------------------------|
  | microcks                      | [GitHub](https://github.com/microcks/microcks)                      |
  | microcks-cli                  | [GitHub](https://github.com/microcks/microcks-cli)                  |
  | microcks-operator             | [GitHub](https://github.com/microcks/microcks-operator)             |
  | microcks-postman-runtime      | [GitHub](https://github.com/microcks/microcks-postman-runtime)      |
  | microcks-testcontainers-java  | [GitHub](https://github.com/microcks/microcks-testcontainers-java)  |
  | microcks-testcontainers-go    | [GitHub](https://github.com/microcks/microcks-testcontainers-go)    |
  | microcks-testcontainers-node  | [GitHub](https://github.com/microcks/microcks-testcontainers-node)  |
  | microcks-testcontainers-dotnet| [GitHub](https://github.com/microcks/microcks-testcontainers-dotnet) |
  | microcks-quarkus              | [GitHub](https://github.com/microcks/microcks-quarkus)              |
  | microcks-java-client          | [GitHub](https://github.com/microcks/microcks-java-client)          |
  | microcks-go-client            | [GitHub](https://github.com/microcks/microcks-go-client)            |
  | microcks-docker-desktop-extension | [GitHub](https://github.com/microcks/microcks-docker-desktop-extension) |
  | microcks-backstage-provider   | [GitHub](https://github.com/microcks/microcks-backstage-provider)   |
  | microcks-jenkins-plugin       | [GitHub](https://github.com/microcks/microcks-jenkins-plugin)       |
  | microcks-spectral-ruleset     | [GitHub](https://github.com/microcks/microcks-spectral-ruleset)     |
  | import-github-action          | [GitHub](https://github.com/microcks/import-github-action)          |
  | test-github-action            | [GitHub](https://github.com/microcks/test-github-action)            |
  | hub.microcks.io               | [GitHub](https://github.com/microcks/hub.microcks.io)               |
  | microcks.io                   | [GitHub](https://github.com/microcks/microcks.io)  

In the [maintainers file](https://github.com/microcks/microcks/blob/23aac2b52c5349947583672aefdc0d2b1fe2d8ce/MAINTAINERS.md#code-owners), however only a handful of above projects are listed. 

**Recommendation**: 
- list all available subprojects with their respective code owners (if applicable). This will also help the discoverability of ongoing initiatives related to Microcks

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

<!-- (TOC Evaluation goes here) --> 

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

<!-- (TOC Evaluation goes here) --> 
The project provides a [summary of the contribution flow
](https://github.com/microcks/.github/blob/d4f58662a652fe2af179332fa63f3babb18bf0c0/CONTRIBUTING.md#summary-of-the-contribution-flow), including:
- [Open an issue](https://github.com/microcks/microcks/issues/new) to report a bug or request a feature
- Pull Requests - to propose changes to existing specification, or fixing typos or obvious errors. PR should include [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/#summary) specification in order to be merged. Specifications are listed [here](https://github.com/microcks/microcks-docker-desktop-extension/blob/7e198a49957443eb3c6ac2697d60ff5365ab5181/CONTRIBUTING.md#conventional-commits).

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

<!-- (TOC Evaluation goes here) --> 
Links to public communication channels are provided on the organization repo [README](https://github.com/microcks/#-join-microcks-community) and community repo [README](https://github.com/microcks/community/blob/main/README.md), which is linked to all existing sub-projects.

These are also listed on the Microcks website: https://microcks.io/community/.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

<!-- (TOC Evaluation goes here) --> 
There are multiple public channels of communication available, such as:
| Platforms       | Link                                                                 |
|-----------------|----------------------------------------------------------------------|
| 💬 Discord (preferred) | https://microcks.io/discord-invite                                   |
| 💬 Slack (alternate)   | https://cloud-native.slack.com/archives/C05BYHW1TNJ                  |
| 💬 Discussions         | https://github.com/orgs/microcks/discussions                         |

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

<!-- (TOC Evaluation goes here) --> 
Community meetings are held on [Zoom](https://github.com/microcks/community/blob/main/JOIN-OUR-MEETINGS.md), bi-weekly with US and APAC friendly timezones, which are showcased on the [CNCF calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/microcks?view=month).

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

<!-- (TOC Evaluation goes here) --> 
[Microcks contributing guidelines](https://github.com/microcks/.github/blob/main/CONTRIBUTING.md) contains details on how to engage with the project using issues and PRs. These guidelines also allows members of the community to increase their level of participation and navigation of the [contributing ladder](https://github.com/microcks/microcks/blob/master/GOVERNANCE.md#contributor-ladder).

- [x] **Demonstrate contributor activity and recruitment.**

<!-- (TOC Evaluation goes here) --> 
The Governance is exercised on a need basis, such as updating the [maintainer list](https://github.com/microcks/.github/issues/18) and adding [emeritus](https://github.com/microcks/.github/commit/87644bce98589a17bf3d26947464a1c790fb1f40) members.

Additionally, the project  participated in the LFX Mentorship Program in 2025 with [7 projects and 7 mentees](https://www.linkedin.com/posts/microcks_opensource-activity-7302340683829268480-t1Ye?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAAwPXkBEN6NhnFGO6FP3z2IT09A1SW07Pg), to encourage new community members to join the project.

## Engineering Principles

### Suggested

- [ ] **Roadmap change process is documented.**

<!-- (TOC Evaluation goes here) --> 

- [ ] **History of regular, quality releases.**

<!-- (TOC Evaluation goes here) --> 

### Required 

- [ ] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  -  _If applicable_ a general Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

<!-- (TOC Evaluation goes here) --> 
There are plenty of resources on project features and capabilities, such as [Why Microcks?](https://microcks.io/blog/why-microcks/). Also, there is an outline on how [different actors can benefit from using the project](https://microcks.io/documentation/overview/what-is-microcks/#who-can-use-microcks), including API Owners, developers, quality assurance and platform engineers. 

**Recommendations**: 
- Clearly outline the project goals and objectives and/or out-of-scope items. This can be solved by completing the GTR or separately in the project documentation. Similar example already available as part of the security [self-assessment](https://github.com/yada/tag-security/blob/main/community/assessments/projects/microcks/self-assessment.md#goals).

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

<!-- (TOC Evaluation goes here) --> 

Plenty of available resources on project features and capabilities, such as [Why Microcks?](https://microcks.io/blog/why-microcks/), [main concepts](https://microcks.io/documentation/overview/main-concepts/), and [What is Microcks?](https://microcks.io/documentation/overview/what-is-microcks/).

Also, there are exiting real-world use cases shared by adopters:
  - [Revolutionizing API Strategy: Lombard Odier's Success Story with Microcks](https://microcks.io/blog/lombard-odier-revolutionizing-api-strategy/)
  - [CNAM Partners with Microcks for Automated SOAP Service Mocking](https://microcks.io/blog/cnam-soap-service-mocking/)
  - [J.B. Hunt: Mock It till You Make It with Microcks](https://microcks.io/blog/jb-hunt-mock-it-till-you-make-it/)

- [ ] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

<!-- (TOC Evaluation goes here) -->

Microcks roadmap is available [here](https://github.com/microcks/.github/blob/main/ROADMAP.md), managed through [GitHub Milestones](https://github.com/microcks/microcks/milestone/57). The roadmap can be filtered by [area](https://github.com/orgs/microcks/projects/1/views/1) or by [status](https://github.com/orgs/microcks/projects/1/views/2), providing different perspectives on the ongoing work.

**Recommendation**:
- provide timelines for milestone completion

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**
  -  _If applicable_ a general Technical Review was completed/updated on DD-MMM-YYYY, and can be discovered at $LINK.

<!-- (TOC Evaluation goes here) --> 

 Available Microcks architecture and software design concepts:
  - [Architecture & deployment options](https://microcks.io/documentation/explanations/deployment-options/)
  - [Deployment topologies](https://microcks.io/documentation/explanations/deployment-topologies/)
  - [Monitoring & Observability](https://microcks.io/documentation/explanations/monitoring/)
  - [Application Configuration](https://microcks.io/documentation/references/configuration/application-config/)


- [ ] **Document the project's release process.**

<!-- (TOC Evaluation goes here) --> 

From the available [Microcks releases](https://github.com/microcks/microcks/releases), it can be observed that container images adhere to a versioning scheme where the **x.y.z** or **x.y.z-fix-N** (for critical fixes) tag denotes a stable release from a GitHub repo tag and is immutable. Additionally, there are mutable tags like `latest` and `nightly` that point to the most recent stable or potentially unstable build, respectively. 

Microcks components are distributed as OCI container images for container runtimes such as Docker or Podman. There is also a guide on [Microcks container images](https://microcks.io/documentation/references/container-images/), software supply chain security, including SBOM and provenance attestations. 

The project has fully [automated the build and release process](https://github.com/microcks/microcks/issues/1468) so all delivered components and their provenance attestations are signed using the GitHub Action provided identities (following the in-toto framework).


**Recommendation**:
- provide clear guidelines on how to cut/deploy a release, including cadence, versioning, etc., under a Release.md or similar file

## Security

### Suggested

N/A

### Required

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance.

- [x] **Clearly defined and discoverable process to report security issues.**

<!-- (TOC Evaluation goes here) --> 
Documented in the [Microcks Security.md](https://github.com/microcks/.github/blob/main/SECURITY.md) file, including [how to report a vulnerability](https://github.com/microcks/.github/blob/d4f58662a652fe2af179332fa63f3babb18bf0c0/SECURITY.md#reporting-a-vulnerability) and [public disclosure timing](https://github.com/microcks/.github/blob/d4f58662a652fe2af179332fa63f3babb18bf0c0/SECURITY.md#public-disclosure-timing).

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

<!-- (TOC Evaluation goes here) --> 
As a GitHub hosted project, Microcks relies on the GitHub authentication mechanisms. All the maintainers and bots (GitHub Actions and Workflow) use two factor authentication and sign commits. The maintainers are responsible for regularly reviewing and updating the organization's membership enforcing 2FA and commit signature checks.

- [x] **Document assignment of security response roles and how reports are handled.**

<!-- (TOC Evaluation goes here) --> 
The security team is made up of a subset of the project [maintainers and code owners](https://github.com/microcks/.github/blob/main/GOVERNANCE.md#maintainers-code-owners-contributors-and-adopters) who are willing and able to respond to vulnerability reports.

The Fix Lead drives the schedule using their best judgment based on severity, development time, and release manager feedback. Security fixes are typically merged into the current development branch and are due for release in the next minor version. The maintainers may create a fix release upon request or, if deemed necessary, as part of a critical security fix.


- [x] **Document Security Self-Assessment.**

<!-- (TOC Evaluation goes here) --> 

Microcks Security Self-Assessment is available [here](https://github.com/yada/tag-security/blob/ce401790fff476ed73f2788574cf7cb52ab64bdc/community/assessments/projects/microcks/self-assessment.md). 

To reinforce the overall commitment to this task and enhance their understanding, Microcks maintainers have completed the Linux Foundation Training & Certification: Security Self-Assessments for Open Source Projects ([LFEL1005](https://training.linuxfoundation.org/express-learning/security-self-assessments-for-open-source-projects-lfel1005/)).


- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

<!-- (TOC Evaluation goes here) --> 
The  Open Source Security Foundation (OpenSSF) score for the main repository is 99% ([View Score](https://www.bestpractices.dev/fr/projects/7513)). However, the Microcks team is committed to improving this score to 100% and achieving the passing stage.

Maintainers currently have a temporary outstanding issue with [Microcks UI-related dependencies](https://github.com/orgs/microcks/discussions/1458) which cannot be upgraded. This issue is specific to the Microcks UI and does not impact the core services of Microcks, which are typically used directly by applications relying on Microcks. This work is in progress, with the aim to resolve it in the coming months.

The project has also made significant efforts to enhance the overall security and compliance across all 19 repositories using [CLOMonitor checks](https://clomonitor.io/projects/cncf/microcks). This was an improvement [process initiated in June 2024](https://github.com/microcks/microcks/issues/1201), showcasing continued commitment to improving project security and best practices.

Also, here is a list of tools the project uses for security of the supply chain:
- Sonar Cloud,
- FOSSA,
- Cosign / Sigstore,
- Clair / Docker Scout,
- Syft

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

<!-- (TOC Evaluation goes here) --> 
Microcks public adopters list is available in [ADOPTERS.md](https://github.com/microcks/.github/blob/main/ADOPTERS.md) file.

Maintainers have reviewed and understood the [CNCF definition of an adopter](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter). During TAG review, it was suggested adding the CNCF adopters categories, which have since has been implemented and tracked in this [issue](https://github.com/microcks/.github/issues/39).


- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

<!-- (TOC Evaluation goes here) --> 
The project provided the TOC with a list of adopters for verification of use of the project. Microcks has been adopted in production at scale at each of the adopters interviewed by the TOC.

- [x] **TOC verification of adopters.**

<!-- (TOC Evaluation goes here) --> 

Refer to the Adoption portion of this document.

- [ ] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

<!-- (TOC Evaluation goes here) --> 
A diagram of available integrations are available in the [Microcks update deck](https://docs.google.com/presentation/d/1dQpmqwOKFroAHPzhAzSXip1AnM96WUH-9UrVNiBgtP0/edit#slide=id.g32b6d8da232_0_279), which was also presented to TAG AppDelivery.

Available integrations:
  - Helm: https://microcks.io/documentation/references/configuration/helm-chart-config/
  - Kubernetes Operator: https://microcks.io/documentation/guides/installation/kubernetes-operator/
  - OIDC: https://microcks.io/blog/mocking-oidc-redirect/
  - Prometheus: https://microcks.io/documentation/explanations/monitoring/#technical-metrics
  - Grafana: https://microcks.io/documentation/explanations/monitoring/#grafana-dashboard
  - OpenTelemetry: https://microcks.io/blog/observability-for-microcks-at-scale/ & https://microcks.io/documentation/explanations/monitoring/#opentelemetry-support
  - Keycloak: https://microcks.io/documentation/guides/administration/users/ & https://microcks.io/documentation/references/configuration/security-config/#identity-management
  - gRPC: https://microcks.io/documentation/tutorials/first-grpc-mock/
  - CloudEvents: https://microcks.io/blog/simulating-cloudevents-with-asyncapi/
  - Cosign: https://microcks.io/documentation/references/container-images/#signatures
  - SLSA: https://microcks.io/documentation/references/container-images/#provenance
  - Podman: https://microcks.io/documentation/guides/installation/podman-compose/
  - Backstage: https://microcks.io/blog/backstage-integration-launch/
  - Traefik: https://doc.traefik.io/traefik-hub/api-mocking/on-premises-setup
  - GitLab: https://about.gitlab.com/blog/2023/09/27/microcks-and-gitlab-part-one/
  - Tekton: https://microcks.io/documentation/guides/automation/tekton/
  - Jenkins: https://microcks.io/documentation/guides/automation/jenkins/
  - GitHub Actions: https://microcks.io/documentation/guides/automation/github-actions/
  - Dagger: https://daggerverse.dev/mod/github.com/fluent-ci-templates/microcks-pipeline@645fe89a0d2a46afbfb778a938cddc06d26b4c4c
  - Testcontainers: https://microcks.io/documentation/guides/usage/developing-testcontainers/ & https://testcontainers.com/modules/microcks/
  - Docker Extension: https://www.docker.com/blog/get-started-with-the-microcks-docker-extension-for-api-mocking-and-testing/

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
