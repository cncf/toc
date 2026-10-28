# KubeVirt Graduation Due Diligence

KubeVirt enables users to run virtual machines alongside containers on
Kubernetes. It brings Kubernetes APIs, scheduling, networking, and storage to
workloads that cannot readily be containerized.

- [Graduation application](https://github.com/cncf/toc/issues/1822)
- [Graduation DD template](../../operations/toc-templates/template-dd-pr-graduation.md)
- [Project repository](https://github.com/kubevirt/kubevirt)
- [Project website](https://kubevirt.io)
- [CNCF project page](https://www.cncf.io/projects/kubevirt/)
- [LFX Insights dashboard](https://insights.linuxfoundation.org/project/kubevirt)
- [Public adopter list](https://github.com/kubevirt/kubevirt/blob/main/ADOPTERS.md)

## Graduation Evaluation Summary for KubeVirt

### Criteria Evaluation

Jeremy Rickard and Mario Fahlandt conducted the due diligence for KubeVirt's
graduation application.

The following areas are noteworthy:

- KubeVirt brings the Kubernetes operational model to virtual machines, allowing
  adopters to manage VM and container workloads together.
- Red Hat IT has run production clusters since 2023. NVIDIA reports 50-60
  production clusters supporting GeForce NOW. Some of its largest clusters have
  about 800 nodes and 3,500 VMs each. Swisscom uses KubeVirt for customer-facing
  Kubernetes infrastructure.
- The project has regular releases, a public roadmap, and a documented process
  for proposing and reviewing enhancements.
- The documentation covers installation, development, and production operations.
  Both interviewed adopters found it useful.
- Governance documents vendor-neutral decision-making and a contributor ladder.
  Recent maintainer promotions and retirements show these processes in use.
- The project explicitly follows the CNCF Code of Conduct and documents how
  reports are handled and escalated.
- The external security audit's High and Medium findings have been fixed.
  The remaining Low and Informational findings have public trackers.

The technical and governance reviews are satisfactory.

We recommend the following non-blocking improvements. The criteria below
include the details and tracking issues:

- Continue fixing the remaining Low and Informational audit findings.
- Publish SBOMs with releases and sign container images.
- Complete the security-release runbook and set timeframes for addressing scan
  findings.
- Improve rollback documentation and tests, sizing guidance, and access to
  benchmarks. Define component SLIs/SLOs.
- Create a plan to help SIG and working-group leaders from underrepresented
  organizations become project maintainers.
- Clarify SIG and subproject ownership and leadership, how groups are added
  or retired, and which shared project policies apply.

### Adoption Evaluation

The adopter interviews and Swisscom reference architecture describe production
use at a level consistent with graduation. Adopters use KubeVirt for enterprise
IT, cloud gaming, and customer-facing Kubernetes services. Managing virtual
machines through Kubernetes and the ability to leverage the Kubernetes ecosystem were common reasons for adoption.

Red Hat and NVIDIA described KubeVirt as stable and useful at scale. Both found
the documentation helpful and reported positive interactions with maintainers.
NVIDIA also described an open community that accepted changes for its workloads
even when other maintainers did not share those needs.

Adopters identified opportunities to improve the project. NVIDIA highlighted
reviewer and maintainer bandwidth, Red Hat expressed interest in Boot from SAN,
and Swisscom described challenges with networking and production examples. 

### Final Assessment

We recommend KubeVirt for graduation based on its production adoption,
engineering practices, governance, and security work.

### Criteria

## Application Process Principles

### Suggested

- [ ] **Engage with domain-specific TAG(s) to present the technical architecture of the project.**

The [graduation application](https://github.com/cncf/toc/issues/1822) cites the
project-authored GTR. A record of TAG engagement or a presentation is still needed to mark this suggested item complete.

### Required


- [x] **Complete a [General Technical Review (GTR)](../../operations/project-reviews/general-technical-questions.md).**

The project completed its
[technical review](https://github.com/kubevirt/community/blob/main/CNCF/general_technical_review.md)
in January 2025. The
[independent graduation review](https://github.com/cncf/toc/pull/2287) is
satisfactory with open follow-up items and is pending merge.

- [x] **Complete a [Governance Review](../../operations/project-reviews/governance-review-template.md).**

The [Governance Review](https://github.com/cncf/toc/pull/2258) is complete and
satisfactory, with non-blocking recommendations. The PR is pending merge.

- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The [project website](https://kubevirt.io/) and
[repository](https://github.com/kubevirt/kubevirt) present KubeVirt as a CNCF
project without requiring a particular vendor's distribution. The
[installation guide](https://kubevirt.io/user-guide/cluster_admin/installation/)
uses upstream Kubernetes manifests, with OKD as a separate option.
[Quickstarts](https://kubevirt.io/user-guide/quickstarts/) cover kind and
minikube, and the [cloud quickstart](https://kubevirt.io/quickstart_cloud/)
covers multiple providers. The public documentation reviewed is consistent
with the vendor-neutrality expectations.

- [x] **Review and acknowledgement of expectations for graduated projects and requirements for moving forward through the CNCF Maturity levels.**
   - [x] Acknowledgement verified in the application.

The [application](https://github.com/cncf/toc/issues/1822) acknowledges the
expectations for graduated projects and maturity progression.

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and
presentation for public comment satisfy this criterion.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

The [user guide](https://kubevirt.io/user-guide/),
[installation instructions](https://kubevirt.io/user-guide/cluster_admin/installation/),
[quickstarts and examples](https://kubevirt.io/user-guide/quickstarts/), and
[developer documentation](https://github.com/kubevirt/kubevirt/tree/main/docs)
provide a good range of end-user and developer documentation. Red Hat reported
that it could meet its implementation needs without inspecting source code.
Improvements are tracked for
[capacity planning](https://github.com/kubevirt/user-guide/issues/1034),
[benchmark discoverability](https://github.com/kubevirt/kubevirt/issues/19266),
and [security self-assessment discoverability](https://github.com/kubevirt/kubevirt/issues/19264).

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review if completed as a suggested item prior to application.

### Suggested

- [x] **Complete a Governance Review**

The [Governance Review](https://github.com/cncf/toc/pull/2258) found the project
satisfactory and recommended improvements to group leadership, discoverability,
and maintainer diversity. Those recommendations are non-blocking. The completed
review is pending merge.

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

The project has updated governance as its community has evolved. For example:

[kubevirt/community#366](https://github.com/kubevirt/community/pull/366)
added maintainer responsibilities, offboarding, mentoring, and emeritus
provisions in January 2025.
[kubevirt/community#365](https://github.com/kubevirt/community/pull/365)
added reviewer/approver removal guidance, and
[kubevirt/community#414](https://github.com/kubevirt/community/pull/414)
clarified maintainer responsibilities for project direction in August 2025.

### Required

- [x] **Clear and discoverable project governance documentation.**

Governance and membership policy are easy to find from the
[community README](https://github.com/kubevirt/community/blob/main/README.md)
and the [contributor guide](https://kubevirt.io/user-guide/contributing/).

A remaining tracking item was opened to improve discoverability in the project website: 
[kubevirt/kubevirt.github.io#1049](https://github.com/kubevirt/kubevirt.github.io/issues/1049).

The Governance Review treats the additional navigation links as non-blocking.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Governance documentation is generally up-to-date. There are some tracked items to clarify documentation for some SIGs and Working Groups:

The SIG Documentation and SIG Testing documentation still needs chairs,
charters, and directory entries, tracked in
[kubevirt/community#457](https://github.com/kubevirt/community/issues/457) and
[kubevirt/community#458](https://github.com/kubevirt/community/issues/458).

The Governance Review considers the chair and charter improvements non-blocking.

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

[Maintainer responsibilities](https://github.com/kubevirt/community/blob/main/GOVERNANCE.md#maintainer-responsibilities)
explicitly include ensuring vendor neutrality. The
[membership policy](https://github.com/kubevirt/community/blob/main/membership_policy.md)
commits to inclusive, vendor-neutral decisions and open meetings, while the
[VEP process](https://github.com/kubevirt/enhancements#process)
provides public proposal review and cross-SIG participation.

Six of the eight maintainers work for Red Hat, which is enough to meet both
the simple-majority voting rule and the two-thirds threshold for governance
changes. 

This concentration is worth addressing even though the policies
support vendor-neutral participation. The
[maintainer-diversity plan](https://github.com/kubevirt/community/issues/468)
is a non-blocking recommendation from the Governance Review.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

[Governance](https://github.com/kubevirt/community/blob/main/GOVERNANCE.md)
documents maintainer nominations, voting, CNCF resource requests, and
governance amendments. The
[membership policy](https://github.com/kubevirt/community/blob/main/membership_policy.md)
defines contribution and leadership roles, while
[contribution guidance](https://github.com/kubevirt/kubevirt/blob/main/CONTRIBUTING.md)
and the VEP process document technical acceptance and project direction.
For SIG/WG structure,
[kubevirt/community#467](https://github.com/kubevirt/community/issues/467)
specifically requests a documented proposal, approval, recognition, and
retirement process.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

The membership policy documents sponsorship, reviewer and approver progression,
leadership responsibilities, and removal procedures. Documentation follow-ups cover
the SIG/WG lifecycle process in
[kubevirt/community#467](https://github.com/kubevirt/community/issues/467),
and leadership and charter documentation in
[kubevirt/community#457](https://github.com/kubevirt/community/issues/457),
[kubevirt/community#458](https://github.com/kubevirt/community/issues/458)
and
[kubevirt/community#459](https://github.com/kubevirt/community/issues/459).

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

The [maintainer lifecycle](https://github.com/kubevirt/community/blob/main/GOVERNANCE.md#selecting-maintainers)
is well documented. It covers qualifications, nomination, voting, onboarding,
retirement and handover, inactivity and removal, emeritus status, and
reinstatement.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Recent changes show the maintainer lifecycle in use:

[kubevirt/community#398](https://github.com/kubevirt/community/pull/398)
added Alice Frosi and Luboslav Pivarc in April 2025;
[kubevirt/community#406](https://github.com/kubevirt/community/pull/406)
records Alice's voluntary retirement and move to emeritus in June 2025.
[kubevirt/community#439](https://github.com/kubevirt/community/pull/439)
added Itamar Holder in April 2026, with approvals from maintainers affiliated
with Red Hat, Google, and NVIDIA. All three changes are merged.

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

The [maintainer list](https://github.com/kubevirt/community/blob/main/MAINTAINERS.md)
includes eight maintainers with names, GitHub contact handles, employers, and
responsibilities. It also includes an emeritus roster.

For October 1, 2025 through October 1, 2026, the
[LFX contributor dashboard](https://insights.linuxfoundation.org/project/kubevirt/contributors?timeRange=custom&start=2025-10-01&end=2026-10-01)
reports 697 active contributors, 121 reviewers, and 123 active organizations.
It lists eight maintainers, matching the count in the project roster.
Seventeen contributors account for about 52% of contribution activity.

These figures use "All activities" with "Include collaborations" turned off.
They count activities such as commits, pull requests, and issue participation,
not just merged code changes. [LFX excludes bot actions](https://insights.linuxfoundation.org/docs/metrics/development/)
from contributor metrics.

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

The project has a broad set of reviewers and approvers documented in
[`OWNERS_ALIASES`](https://github.com/kubevirt/kubevirt/blob/main/OWNERS_ALIASES).

This includes core approvers and SIG-specific roles for compute, storage,
networking, testing, and scale. Alongside the project maintainers, this provides
coverage appropriate to the project's size and scope.

- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

KubeVirt has maintainers from Red Hat, Google, and NVIDIA. Six work for Red Hat,
so there is room to improve the balance. 

For the same reporting period, Red Hat accounts for 71,874 contribution
activities, about 91% on the
[LFX organization leaderboard](https://insights.linuxfoundation.org/project/kubevirt/contributors?timeRange=custom&start=2025-10-01&end=2026-10-01#organizations-leaderboard).

The diversity plan in
[kubevirt/community#468](https://github.com/kubevirt/community/issues/468)
is a non-blocking, one-year goal and the project should focus on improving this balance. We recommend a follow up at one year to review with the TOC. 

- [x] **Code and Doc ownership in GitHub and elsewhere matches documented governance roles.**

In general, code and documentation ownership matches the documented governance
roles. The [core ownership aliases](https://github.com/kubevirt/kubevirt/blob/main/OWNERS_ALIASES)
and [documentation ownership aliases](https://github.com/kubevirt/user-guide/blob/main/OWNERS_ALIASES)
assign reviewers and approvers by area, including SIG-specific roles.

Remaining repository-to-SIG mismatches are tracked for
[SIG Control Plane](https://github.com/kubevirt/community/issues/465) and the
[release and enhancement repositories](https://github.com/kubevirt/community/issues/466).
[kubevirt/community#470](https://github.com/kubevirt/community/pull/470) proposes
the Control Plane subproject declaration and clarifies its boundary with SIG Compute.

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

The [community Code of Conduct](https://github.com/kubevirt/community/blob/main/code-of-conduct.md)
explicitly states that KubeVirt and its community abide by the CNCF Code of
Conduct. The
[organization-wide policy](https://github.com/kubevirt/.github/blob/main/CODE_OF_CONDUCT.md)
does the same.  Membership policy requires contributors to adhere to the conduct standards.


Governance assigns confidential report handling to maintainers,
provides CNCF escalation when a maintainer is involved, and excludes an accused
maintainer from the relevant closed meeting. A dedicated code of conduct committee would be a better option for the project, to remove the project maintainers from any potential conflict. 

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

[Governance](https://github.com/kubevirt/community/blob/main/GOVERNANCE.md#code-of-conduct)
links to the community policy, which links directly to the CNCF Code of
Conduct. The
[membership policy](https://github.com/kubevirt/community/blob/main/membership_policy.md#new-contributors)
links to both policies, and the
[contributor guide](https://kubevirt.io/user-guide/contributing/)
also directs contributors to the conduct requirements.

- [x] **All subprojects, if any, are listed.**

The project lists its subprojects in
[`sigs.yaml`](https://github.com/kubevirt/community/blob/main/sigs.yaml),
grouped by SIG with links to the relevant OWNERS files. Non-blocking follow-ups
clarify SIG Control Plane's subproject declaration in
[kubevirt/community#465](https://github.com/kubevirt/community/issues/465), with
a proposed update in
[kubevirt/community#470](https://github.com/kubevirt/community/pull/470), and
the release and enhancement repository mappings in
[kubevirt/community#466](https://github.com/kubevirt/community/issues/466).

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

The [SIG and subproject inventory](https://github.com/kubevirt/community/blob/main/sigs.yaml)
records ownership and leadership, while the
[membership policy](https://github.com/kubevirt/community/blob/main/membership_policy.md)
describes subproject leadership roles and contribution responsibilities.
We recommend making clear which leadership, contribution, maturity, and
add/remove rules each subproject follows, including where it uses shared
project policies. This and the
[SIG/WG lifecycle follow-up](https://github.com/kubevirt/community/issues/467)
are non-blocking.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

The [contributor ladder](https://github.com/kubevirt/community/blob/main/membership_policy.md#contributor-ladder)
defines contributor, organization member, reviewer, approver, SIG chair,
subproject lead, WG chair, and project maintainer roles. It provides
role-specific requirements and responsibilities, sponsorship and nomination
paths, and inactivity/removal guidance.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

[`CONTRIBUTING.md`](https://github.com/kubevirt/kubevirt/blob/main/CONTRIBUTING.md)
documents issue and PR submission, VEPs for substantial changes, testing,
DCO signoff, review assignment, and escalation when a contribution needs
attention. The
[contributor guide](https://kubevirt.io/user-guide/contributing/)
provides entry points for code and non-code contributions.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

The [community README](https://github.com/kubevirt/community#socializing)
and [website community page](https://kubevirt.io/community/)
document the public `#virtualization` and `#kubevirt-dev` channels in the
Kubernetes Slack workspace, the `kubevirt-dev` Google Group, and open community
meetings.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

Project-wide channels and the
[SIG/WG directory](https://github.com/kubevirt/community/blob/main/sig-list.md)
provide a public inventory of meetings.
[Governance](https://github.com/kubevirt/community/blob/main/GOVERNANCE.md#meetings)
documents private maintainer communications and closed meetings for security
reports and conduct matters; the
[security policy](https://github.com/kubevirt/kubevirt/blob/main/SECURITY.md)
documents private vulnerability reporting.

Contact and meeting details are tracked for
[SIG Network](https://github.com/kubevirt/community/issues/460),
[SIG Buildsystem](https://github.com/kubevirt/community/issues/461),
[s390x WG](https://github.com/kubevirt/community/issues/462),
[ARM WG](https://github.com/kubevirt/community/issues/463), and
[SIG Storage and WG-AIE](https://github.com/kubevirt/community/issues/464).

[kubevirt/community#471](https://github.com/kubevirt/community/pull/471) is an
open PR documenting ARM WG contacts and current meeting status. 

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

The [community meeting guide](https://github.com/kubevirt/community/blob/main/community_meeting.md)
documents weekly open meetings, agendas, notes, recordings, and a
[public calendar](https://calendar.google.com/calendar/u/0/embed?src=kubevirt@cncf.io).
The SIG/WG directory also publishes schedules for several groups.

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

The [contributor guide](https://kubevirt.io/user-guide/contributing/)
progresses from Git and open-source prerequisites through first contributions,
documentation/community work, code contributions, and leadership resources.
The [core contribution guide](https://github.com/kubevirt/kubevirt/blob/main/CONTRIBUTING.md)
adds testing, CI permission handling for new contributors, DCO, review, and
VEP expectations, with links to the developer guide. The merged governance
cross-link in
[kubevirt/kubevirt#19279](https://github.com/kubevirt/kubevirt/pull/19279)
is a recent improvement.

- [x] **Demonstrate contributor activity and recruitment.**

The core repository had
[203 merged PRs](https://github.com/kubevirt/kubevirt/pulls?q=is%3Apr+is%3Amerged+merged%3A2026-09-15..2026-10-15)
between September 15 and October 15, 2026, including automated changes.

The project welcomes newcomers and provides paths into review and leadership
roles. For example,
[kubevirt/kubevirt#19268](https://github.com/kubevirt/kubevirt/issues/19268)
shows a new contributor being welcomed and completing the linked PR.
The
[September reviewer promotion](https://github.com/kubevirt/kubevirt/pull/19222)
and [maintainer promotion](https://github.com/kubevirt/community/pull/439)
show contributors becoming reviewers and maintainers.

## Engineering Principles

### Suggested

N/A

### Required

- [x] **Document project goals and objectives that illustrate the project's differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**

  - The [project-authored GTR](https://github.com/kubevirt/community/blob/main/CNCF/general_technical_review.md)
    was merged on 2025-01-28 and documents the project's goals, scope, and
    differentiation.

KubeVirt lets users manage virtual machines through the same APIs, RBAC,
networking, storage, and scheduling model as container workloads. This gives
organizations a way to bring existing VM workloads onto Kubernetes without
first having to containerize them.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases. _This requirement may also be satisfied by completing a General Technical Review._**

The [project website](https://kubevirt.io), repository documentation, and GTR
explain how KubeVirt manages VM workloads through Kubernetes custom resources
and controllers. The
[public adopter list](https://github.com/kubevirt/kubevirt/blob/main/ADOPTERS.md)
reports use cases including legacy applications, CI runners, virtual desktops,
GPU-backed AI/ML VMs, confidential-computing environments, training platforms,
and telecommunications workloads.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The project maintains a public
[roadmap](https://github.com/kubevirt/sig-release/blob/main/ROADMAP.md), a
[VEP repository](https://github.com/kubevirt/enhancements), a public
[enhancements tracking board](https://github.com/orgs/kubevirt/projects/21),
and rolling
[upcoming changes](https://github.com/kubevirt/sig-release/blob/main/upcoming-changes.md).

- [x] **Roadmap change process is documented.**

Roadmap changes follow the public VEP process. SIGs review proposals and
prioritize them for releases. The roadmap makes clear that a planned feature
may move to a later release if it misses code freeze.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review and capturing the output in the project's documentation._**

The [user guide](https://kubevirt.io/user-guide/) and GTR document the
major components: `virt-operator` deploys components and CRDs; `virt-api`
serves VM, VMI, and migration APIs; `virt-controller` reconciles KubeVirt
objects; `virt-handler` manages instances on each node; and `virt-launcher`
hosts each VM workload in a pod.

KubeVirt reuses Kubernetes RBAC, CNI, CSI, scheduling, and Pod Security Standards.
It supports live migration, GPU passthrough, SR-IOV, disk-image handling through
CDI, and multiple architectures. The documentation would benefit from more
[sizing guidance](https://github.com/kubevirt/user-guide/issues/1034), particularly
control-plane requirements, per-VM overhead, and production recommendations.

- [x] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:**

  - [x] Release expectations (scheduled or based on feature implementation)
  - [x] Tagging as stable, unstable, and security related releases
  - [x] Information on branch and tag strategies
  - [x] Branch and platform support and length of support
  - [x] Artifacts included in the release.

The release process is documented in
[`docs/release.md`](https://github.com/kubevirt/kubevirt/blob/main/docs/release.md),
with a companion automated release procedure and
[backport policy](https://github.com/kubevirt/kubevirt/blob/main/docs/release-branch-backporting.md).
KubeVirt uses a time-based cadence aligned behind Kubernetes releases. Signed
semantic-version tags distinguish alpha, beta, release-candidate, GA, and patch
releases. Stable branches use the `release-<major>.<minor>` convention and
accept backports according to the project's documented policy.

A release reaches [end of life](https://github.com/kubevirt/kubevirt/blob/main/docs/release.md#end-of-life-eol)
when the support period for its corresponding Kubernetes release ends.
KubeVirt does not currently enforce this cutoff; maintainers may continue
approving backports after EOL.

The
[Kubernetes compatibility policy](https://github.com/kubevirt/kubevirt/blob/main/docs/kubernetes-compatibility.md)
and [support matrix](https://github.com/kubevirt/sig-release/blob/main/releases/k8s-support-matrix.md)
document supported platform versions. Releases include source tags, container
images, `virtctl` binaries, and release notes. The following improvements are
recommended:

1. Publish SBOMs with release artifacts:
   [kubevirt/kubevirt#19261](https://github.com/kubevirt/kubevirt/issues/19261).
2. Sign container images and document verification:
   [kubevirt/kubevirt#19262](https://github.com/kubevirt/kubevirt/issues/19262).
   Source tags are signed, but container-image signing is still being discussed.

The proposed security-release runbook in
[kubevirt/kubevirt#19203](https://github.com/kubevirt/kubevirt/pull/19203)
would add embargo handling, release roles, and disclosure-day procedures. It is
still open.

- [x] **History of regular, quality releases.**

The project released monthly before v1.0 and moved to three minor
releases per year after v1.0 in 2023. The release history includes stable
patches such as v1.9.1 and prereleases such as v1.10.0-beta.0. Releases include
human-readable notes and are gated through Prow-based unit and end-to-end
testing.

The GTR raised a question about rollback after a completed upgrade, tracked in
[kubevirt/kubevirt#19267](https://github.com/kubevirt/kubevirt/issues/19267).
The [maintainer's response](https://github.com/cncf/toc/pull/2287#discussion_r4084863374)
clarifies that rollback of an incomplete or stuck upgrade is supported. New API
fields, CRDs, and features are blocked until rollout completes. Once an upgrade
has completed and new features may have been used, downgrade behavior is
undefined and discouraged. The open
[documentation PR](https://github.com/kubevirt/user-guide/pull/1033) explains
this boundary, but does not provide the CI coverage requested in the issue.

Other recommendations are to make existing
[performance benchmarks easier to find](https://github.com/kubevirt/kubevirt/issues/19266)
and define [component SLIs/SLOs](https://github.com/kubevirt/kubevirt/issues/19265).


## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance if completed as a suggested item prior to application.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

KubeVirt has the passing badge but has not yet achieved Silver or Gold.

- [ ] **Complete a joint security assessment with TAG Security and Compliance**

We did not find a completed joint assessment with TAG Security and Compliance.
The published
[self-assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/kubevirt/self-assessment.md)
and external audit are separate reviews.

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

[`SECURITY.md`](https://github.com/kubevirt/kubevirt/blob/main/SECURITY.md)
documents private reporting to the project's CNCF security mailing list,
requested report contents, confidential handling, private vulnerability
reporting as an alternate path, and disclosure through GitHub Security
Advisories and project communications.

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

KubeVirt requires members to enable two-factor authentication on their GitHub
accounts. The
[membership checklist](https://github.com/kubevirt/community/blob/80200f687f3bd8fb746647fce5780b3474126a16/membership_checklist.md?plain=1#L7)
requires applicants to confirm, "I have enabled 2FA on my GitHub account."
This is also a requirement in the
[membership policy](https://github.com/kubevirt/community/blob/main/membership_policy.md#requirements).

The [main branch](https://api.github.com/repos/kubevirt/kubevirt/branches/main)
was protected when checked. Required checks included DCO, build, unit,
end-to-end, and lint checks. The
[Prow configuration](https://github.com/kubevirt/project-infra/blob/main/github/ci/prow-deploy/kustom/base/configs/current/config/config.yaml)
requires `lgtm` and `approved` labels for normal merges and uses the repository's
OWNERS roles.

- [x] **Document assignment of security response roles and how reports are handled.**

`SECURITY.md` assigns response to KubeVirt maintainers with escalation
support from vendor security teams and documents private triage, embargoed
fixes, advisories, patch releases, and public notices.

[kubevirt/kubevirt#19203](https://github.com/kubevirt/kubevirt/pull/19203)
proposes a runbook with specific security-release roles and an incident
workflow. This would be a useful improvement to the existing policy; the PR
is still open.

[kubevirt/kubevirt#19263](https://github.com/kubevirt/kubevirt/issues/19263)
requests a remediation-time commitment for FOSSA/SonarCloud software composition
analysis findings. The proposed runbook describes urgency rather than numeric
timelines, so this remains a separate non-blocking recommendation.

- [x] **Document Security Self-Assessment.** The completed self-assessment must be available in `cncf/toc` under `projects/<project>/security-assessment/self-assessment.md`.

KubeVirt previously published a self-assessment in the archived TAG Security repository.
[cncf/toc#2310](https://github.com/cncf/toc/pull/2310) was opened to migrate the
assessment back into this repository at
`projects/kubevirt/security-assessment/self-assessment.md`. 

[kubevirt/community#469](https://github.com/kubevirt/community/issues/469)
tracks the discrepancy. Separately,
[kubevirt/kubevirt#19264](https://github.com/kubevirt/kubevirt/issues/19264)
tracks linking the assessment from `SECURITY.md`. The pull request to add that link,
[kubevirt/kubevirt#19300](https://github.com/kubevirt/kubevirt/pull/19300),
is intended to merge after the TOC migration.

The review accepts the completed self-assessment. Its migration into this repository and the related documentation updates are non-blocking follow-ups.

- [x] **Third Party Security Review.**

  - [x] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

Quarkslab performed a 37-day audit in early 2025, facilitated by OSTIF and funded
by the Sovereign Tech Fund in coordination with CNCF. It covered
`virt-operator`, `virt-api`, `virt-controller`, `virt-handler`, and
`virt-launcher`. The
[full report](https://ostif.org/wp-content/uploads/2025/10/KubeVirt_OSTIF_Report_25-06-2150-REP_v1.2.pdf)
and summaries from
[OSTIF](https://ostif.org/kubevirt-audit-complete/),
[CNCF](https://www.cncf.io/blog/2025/12/17/kubevirt-undergoes-ostif-security-audit/),
and [KubeVirt](https://kubevirt.io/2025/Announcing-KubeVirt-Security-Audit-Results.html)
are public.

The audit reported 15 findings: one High, seven Medium, four Low, and three
Informational.

The High finding,
[CVE-2025-64324](https://github.com/kubevirt/kubevirt/security/advisories/GHSA-46xp-26xh-hpqh),
was patched in 1.6.1 and 1.7.0. Published
[security advisories](https://github.com/kubevirt/kubevirt/security/advisories)
also document fixes for six Medium findings. The remaining Medium finding,
MEDIUM-12, was addressed by
[requiring control-plane placement](https://github.com/kubevirt/kubevirt/pull/15157);
the fix is present in v1.7.0.

The remaining Low findings are tracked in
[LOW-4](https://github.com/kubevirt/kubevirt/issues/19286),
[LOW-13](https://github.com/kubevirt/kubevirt/issues/19287),
[LOW-14](https://github.com/kubevirt/kubevirt/issues/19288), and
[LOW-15](https://github.com/kubevirt/kubevirt/issues/19289).
The Informational findings are tracked in
[INFO-1](https://github.com/kubevirt/kubevirt/issues/19290),
[INFO-10](https://github.com/kubevirt/kubevirt/issues/19291), and
[INFO-11](https://github.com/kubevirt/kubevirt/issues/19292).

These cover certificate identity checks, component and device access, sidecar
integrity, and error handling. Fixes remain outstanding; the proposed INFO-10
mitigation is partial. The project meets the audit and remediation-tracking
criteria without having completed every Low and Informational fix.
The tracking issues should record fixes, affected releases, and any risks
the project accepts.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

KubeVirt holds the
[OpenSSF Best Practices passing badge](https://www.bestpractices.dev/projects/3203) at 100%.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

[`ADOPTERS.md`](https://github.com/kubevirt/kubevirt/blob/main/ADOPTERS.md)
groups organizations into end-user, integration, and vendor categories and
usually includes a year, link, and use-case description. Adding missing dates
and distinguishing production from evaluation would make the list more useful.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The Red Hat and NVIDIA interviews and Swisscom reference architecture describe
production use by three independent organizations. They cover enterprise IT,
cloud gaming, and customer-facing Kubernetes services. The public adopter list
also reports use across AI, telecommunications, training, and cloud
infrastructure.

The project provided the TOC with a list of adopters for verification of use of the project at the level expected, i.e. production use for graduation, dev/test for incubation.

- [x] **TOC verification of adopters.**

Red Hat IT was interviewed on June 26, 2026, and NVIDIA GeForce NOW on July 28,
2026. Both described production use, operational value, and interactions with
the project.

Both organizations employ KubeVirt maintainers. The Red Hat interview covers
IT operations, while the NVIDIA interview also includes an active contributor's
perspective. The NVIDIA interviewee was working toward a reviewer role.
Swisscom, which is not represented among current maintainers, provides an
additional operational perspective through its published reference architecture.
That record is accepted in place of a third interview under the
[DD guide](../../operations/dd-toc-guide.md#conducting-adopter-interviews).

Refer to the Adoption portion of this document.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

KubeVirt documents integrations with Kubernetes and other CNCF and non-CNCF
projects.

| Integration | CNCF relationship | Documentation |
| :--- | :--- | :--- |
| Kubernetes: CRDs, controllers, API aggregation, RBAC, and Pod Security Standards | CNCF graduated | [Architecture and installation](https://kubevirt.io/user-guide/cluster_admin/installation/) |
| Containerized Data Importer | KubeVirt sub-project | [Repository](https://github.com/kubevirt/containerized-data-importer) |
| Cilium | CNCF graduated | [Cilium and KubeVirt lab](https://isovalent.com/labs/cilium-kubevirt/) |
| Prometheus | CNCF graduated | [Monitoring guide](https://kubevirt.io/user-guide/operations/component_monitoring/) |
| containerd and CRI-O | CNCF graduated | [Installation requirements](https://kubevirt.io/user-guide/cluster_admin/installation/) |
| Multus and SR-IOV | CNCF ecosystem | [Interfaces and networks](https://kubevirt.io/user-guide/network/interfaces_and_networks/) |
| Backup and disaster recovery | Commercial ecosystem | [Kasten VM protection](https://docs.kasten.io/latest/usage/vm_protection/), [Trilio VM backups](https://docs.trilio.io/kubernetes/user-guide/getting-started-with-management-console/navigating-intro/backup-and-recovery/virtual-machines), [CloudCasa KubeVirt backup/restore](https://docs.cloudcasa.io/help/reference-vm.html), and [Portworx VM backups](https://docs.portworx.com/portworx-backup-on-prem/use-px-backup/backup-restore/create-backup/backup-vm) |
| oVirt and OKD | Non-CNCF/community | [oVirt external provider guide](https://www.ovirt.org/documentation/administration_guide/index.html#proc-adding-kubevirt-openshift-as-an-external-provider_external_providers) and [OKD Virtualization](https://docs.okd.io/latest/virt/about_virt/about-virt.html) |

Commercial and community products built on KubeVirt provide indirect-adoption
evidence, including OpenShift Virtualization, SUSE Virtualization, KubeSphere
Virtualization, Kubermatic KKP Virtualization, Azure Operator Nexus, Cozystack,
Deckhouse, Oracle Cloud Native Environment, Spectro Cloud Palette VMO, and
Platform9 Managed KubeVirt. 

#### Adoption

##### Adopter 1 - Red Hat

Red Hat IT operates on-premises and cloud infrastructure for internal and
customer workloads.

Brian Atkisson was interviewed by Mario Fahlandt on June 26, 2026. Red Hat IT
began proof-of-concept work in 2021 and put production clusters into service
in 2023. At the time of the interview, it reported approximately 40 clusters
and 500-1000 nodes, using large nodes to support high virtual-machine density.

Red Hat IT adopted KubeVirt to operate virtual machines alongside containerized
and serverless workloads through Kubernetes and OpenShift. Its earlier
environment relied heavily on Red Hat Virtualization and OpenStack. Key
benefits included bringing existing VM and bare-metal workloads into the
platform, reusing Kubernetes storage and backup integrations, applying
infrastructure as code to fleet management, and using live migration.

For its most important workloads, Red Hat IT described an annual major-version
update cadence, with more frequent maintenance releases for other clusters.
It had upgraded more frequently when it needed new features, but now considers
the project stable and sufficient for most of its requirements.

Red Hat IT found what it needed in the documentation without reading source
code. KubeVirt helped it consolidate virtualization teams and reduce the compute,
storage, and load balancers it needed. Customers also had a single way to manage
VMs instead of having to learn several different APIs.

Red Hat IT described maintainers as responsive and said earlier bug reports
and feature requests had been addressed. It has needed less direct engagement
recently and did not offer an overall assessment of the community or its
governance. It plans to continue using KubeVirt and is interested in Boot from
SAN support. Red Hat highlighted managing VMs alongside other workload types
and bringing the Kubernetes ecosystem to virtual machines as the project's
main strengths.

The entire adopter interview can be found here:
[Red Hat Adopter Interview](kubevirt-adopter-interview-red-hat.md).

##### Adopter 2 - NVIDIA

NVIDIA's GeForce NOW platform runs online gaming workloads in virtual machines.

Alay Patel was interviewed by Jeremy Rickard on July 28, 2026. NVIDIA reported
using KubeVirt for several years, since at least version 0.35. In a follow-up,
NVIDIA clarified that it has about 50-60 production clusters, with some of the
largest having about 800 nodes and 3,500 VMs each.

NVIDIA chose KubeVirt as an open-source way to manage VM workloads through
Kubernetes. The team found the APIs and Kubernetes integrations especially useful
and said replacing KubeVirt would be difficult.

NVIDIA was using KubeVirt 1.7 at the time of the interview and generally runs
about two versions behind, because rolling out upgrades across its deployment
takes considerable time. Early adoption involved upgrade problems and API
breaks when moving between versions. NVIDIA now considers the API surface
stable and described contributing API round-trip tests around the transition
to v1 to help validate upgrades.

NVIDIA found the GPU-integration examples, user guide, and contribution
documentation useful. It uses performance and scale data to qualify upgrades
and highlighted openness, scale, and the API surface as project strengths.

NVIDIA described the community as open and receptive, including support for
changes that NVIDIA needed even when those changes did not address other
maintainers' needs. The team discussed feature requests through Slack, GitHub
issues and pull requests, and mailing lists, and said it was generally satisfied
with the results.
The interviewee participates in Dynamic Resource Allocation integration,
device-management work, and SIG Scale, and was working toward a reviewer role.
The team reported driving five or six enhancements and plans to develop more
community members, reviewers, and maintainers.

NVIDIA would like to see more reviewer and maintainer capacity and further
improvements to API evolution and conventions. It found the contributor ladder
well documented, but attracting additional maintainers remains difficult.

The entire adopter interview can be found here:
[NVIDIA Adopter Interview](kubevirt-adopter-interview-nvidia.md).

##### Adopter 3 - Swisscom / Telecommunications and Cloud

Swisscom provides a Kubernetes Service for Swiss enterprise customers. Its
[published CNCF reference architecture](https://architecture.cncf.io/architectures/swisscom-kubernetes-service/)
is used in place of a third adopter interview.

Swisscom reports using KubeVirt since 2024, with v1.5.0 recorded when the
architecture was published. Its Pathfinders operations team uses KubeVirt to
create and manage ephemeral control-plane and worker VMs for Swisscom's
customer-facing Kubernetes Service. KubeVirt runs on bare-metal infrastructure
and provides the VM abstraction for these cluster nodes; `kubevirt-csi`
provides persistent storage to user clusters.

The architecture describes a sovereign, highly available private-cloud service
for Swiss enterprise customers. Swisscom chose open source components to reduce
vendor lock-in and licensing dependency, retain control of platform
development and operations, and keep customer data under Swiss jurisdiction.
Swisscom also described several operational challenges:
limited production-scale KubeVirt examples and reference designs, networking
complexity, limited enterprise support options, and the need to build
specialized in-house skills. Swisscom operates the platform with professional
support from Kubermatic.
