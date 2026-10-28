# Cozystack Incubation Due Diligence

- Link to [Incubation application issue](https://github.com/cncf/toc/issues/1916)

## Incubation Evaluation Summary for Cozystack

### Criteria Evaluation

Faseela K conducted the due diligence of Cozystack, which applied for Incubation. Cozystack is a CNCF Sandbox project with evidence of technical maturity, production use by a commercial cloud provider, a complete technical review, a security self-assessment, a public roadmap, regular releases with maintained patch lines, and a contributor ladder exercised in both directions.

The project has a clear technical scope for building private and public clouds on Kubernetes; a deployment model that works from a single machine to multi-datacenter production; operational and contributor documentation; a published vendor-neutrality commitment; and an adopter base spanning consultancy, hosting, sovereign and public cloud, GPU and educational environments documented in `ADOPTERS.md`, broadened further by interviewed adopters in scientific computing and higher-education research.

_The following criteria implementations are noteworthy to call out:_

- The project's differentiation criterion is clear and substantiated, not just asserted: Cozystack is a Kubernetes-native platform delivering virtual machines, managed Kubernetes clusters, and managed databases from a single control plane, built from well-known CNCF/open-source components rather than as a black box, a scope and approach distinct from both proprietary virtualization stacks and narrower single-purpose projects.
- The project responds to upstream-dependency risk directly: the `etcd-operator` subproject was moved out of the founding vendor's GitHub organization into `cozystack` and its API group renamed from `etcd.aenix.io` to `etcd-operator.cozystack.io` (v1.6.0).
- The release process is documented in-repo and the tag history matches it, including concurrent patch activity on older release lines, evidence of genuinely maintained backport branches.
- The project publishes its own OSS health metrics and OpenSSF Best Practices badge verification state, rather than only pointing at third-party dashboards.
- The security documentation is candid, including explicit disclosure of where the project is not meeting its own remediation targets.
- The sub-project model is documented with a packaged-versus-independent-tooling split, per-repository owners, published maturity status, and an accounting of forks and private repositories.
- Governance was iterated during this due diligence rather than defended: an unfilled Directors role was removed, organization-wide default Code of Conduct/SECURITY.md files were added, and the subproject listings were reconciled between the application and `GOVERNANCE.md`.
- The project proactively monitors its own dependency licensing rather than waiting to be asked: three of the eight outstanding license exceptions were self-identified and filed without prompting, two of them early in this review (20-Jul-2026); that experience led the project to run a more thorough internal audit, which surfaced the remaining exceptions as the review progressed, including three covering independent tooling (`talos-meta-tool`, `terraform-provider-cozystack`, dashboard console assets) that isn't even part of the packaged distribution.
- Adopter feedback is strongly positive on the integrated, easy-to-use stack: data sovereignty and freedom from vendor lock-in are recurring adoption drivers across interviews, and adopters describe the multi-tenancy model as more mature than anything they could have built themselves. The publicly documented adopter base in `ADOPTERS.md` independently spans the same themes at broader scale, including sovereign and public clouds, hosting providers, and GPU/educational environments.

_The following actions were provided to the project that were considered blocking but have since been resolved:_

- **[cozystack/community#76](https://github.com/cozystack/community/issues/76) - Vendor-neutrality: website footer/GitHub org LinkedIn pointed to Ænix's corporate accounts**: fixed in [cozystack/website#690](https://github.com/cozystack/website/pull/690) (merged).
- **[cozystack/community#78](https://github.com/cozystack/community/issues/78) - SECURITY.md/CODE_OF_CONDUCT.md not present across all cozystack org repos**: fixed in [cozystack/.github#8](https://github.com/cozystack/.github/pull/8) (merged).
- **[cozystack/community#80](https://github.com/cozystack/community/issues/80) - Subproject listings don't align between application and GOVERNANCE.md**: fixed by reconciling the subproject list between the application and `GOVERNANCE.md`.

_One condition remains outstanding. This does not block this review or the TOC vote, but it does block the project's actual move to Incubation level and the associated benefits until resolved:_

- **Eight CNCF license exception requests filed, all still pending review**:
  - Percona MongoDB operator (Apache-2.0) with SSPL-1.0 deploy-time payload ([cncf/foundation#1465](https://github.com/cncf/foundation/issues/1465))
  - RabbitMQ operator manifests and OpenBao Helm chart, MPL-2.0 ([cncf/foundation#1466](https://github.com/cncf/foundation/issues/1466))
  - Talos bootstrap provider (CABPT) manifests and Talos Linux images, MPL-2.0 ([cncf/foundation#1539](https://github.com/cncf/foundation/issues/1539))
  - Sidero Labs Go libraries in `talm`, MPL-2.0 ([cncf/foundation#1541](https://github.com/cncf/foundation/issues/1541))
  - LINSTOR server, GPL-3.0-only ([cncf/foundation#1542](https://github.com/cncf/foundation/issues/1542))
  - Sidero Labs Go libraries in `talos-meta-tool`, MPL-2.0 ([cncf/foundation#1570](https://github.com/cncf/foundation/issues/1570))
  - HashiCorp Terraform plugin libraries in `terraform-provider-cozystack`, MPL-2.0 ([cncf/foundation#1571](https://github.com/cncf/foundation/issues/1571))
  - noVNC, Geist font and Simple Icons in the dashboard console, MPL-2.0/OFL-1.1/CC0-1.0 ([cncf/foundation#1572](https://github.com/cncf/foundation/issues/1572))

  Not blocking for this review or the TOC vote; blocked for the project's move to Incubation level and incubation-level benefits until all eight are approved.

_The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:_

- Carrying the vendor-neutrality commitment into `GOVERNANCE.md` ([cozystack/community#97](https://github.com/cozystack/community/issues/97)); tracked under [cozystack/cozystack#4934](https://github.com/cozystack/cozystack/pull/4934).
- Broadening organizational representation among active maintainers: per [MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md), Ænix holds 6 of 9 active maintainer seats (about 67%), below the 75% threshold that would trigger a mandatory practice period, but still worth active attention and recruitment effort from other organizations.
- Publishing an API stability and deprecation policy; tracked under [cozystack/community#112](https://github.com/cozystack/community/issues/112).
- Documenting the subproject add/remove process; tracked under [cozystack/cozystack#4933](https://github.com/cozystack/cozystack/pull/4933).
- Minor DCO and CodeQL branch-protection refinements: neither check was a required status check on `main`; now tracked under [cozystack/cozystack#4932](https://github.com/cozystack/cozystack/pull/4932).
- Kamaji's upstream dependency concentration, noted for awareness based on adopter interviews (not a tracked DD item).

_Non-blocking recommendations already completed during this review:_

- **[cozystack/community#75](https://github.com/cozystack/community/issues/75) - GTR snapshot submission**: snapshot PR [cncf/toc#2305](https://github.com/cncf/toc/pull/2305) merged 2026-10-01.
- **[cozystack/community#79](https://github.com/cozystack/community/issues/79) - Directors role undocumented**: fixed in [cozystack/cozystack#4296](https://github.com/cozystack/cozystack/pull/4296) (merged).
- **[cozystack/community#81](https://github.com/cozystack/community/issues/81) - Security self-assessment snapshot submission**: snapshot merged in [cncf/toc#2306](https://github.com/cncf/toc/pull/2306).
- **Improvements for new-contributor onboarding**: contributor-welcoming/CI-approval guidance added to `GOVERNANCE.md` ([cozystack/cozystack#4845](https://github.com/cozystack/cozystack/pull/4845)), alongside further communication-channel updates ([cozystack/website#744](https://github.com/cozystack/website/pull/744)).


### Adoption Evaluation

Three adopters were interviewed, together meeting and exceeding the dev/test bar expected for Incubation: a scientific-computing centre in pre-production, a university research institute running a live multi-site deployment, and a regional cloud and managed-services provider in commercial production since early 2025. All three interviews are presented anonymized. No adopter reported an upstream blocker. Full findings are summarized under [TOC verification of adopters](#toc-verification-of-adopters) in the Ecosystem section below.

Common strengths reported by adopters:

- An integrated, open platform that's simple to operate, with no black box.
- Data sovereignty and freedom from vendor lock-in.
- A multi-tenancy model more mature than anything they could have built themselves.
- Self-service provisioning that cuts out manual operational work.
- Cross-industry applicability, including public-sector use and regulated-industry interest.
- Responsive, hands-on maintainers.
- Upgrade safety across varied release cadences.

Common areas for improvement reported by adopters:

- Day-2 operations and troubleshooting documentation.
- An API stability and deprecation policy.
- Clearer upgrade notes between minor versions.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Incubation. The blocking remediation raised during this review has been completed and independently verified. The project’s move to Incubation level, and the benefits that come with it, wait on approval of the eight license exception requests the project filed with the CNCF Governing Board, all of which remain open.

## Application Process Principles

### Suggested

- [x] **Engage with domain-specific TAG(s) to present the technical architecture of the project.**

The project presented to TAG Infrastructure on 05-NOV-2025; [recording](https://youtu.be/gJ4GYeLB-V0?t=2066). The project also engaged TAG App Delivery during Sandbox review ([cncf/tag-app-delivery#719](https://github.com/cncf/tag-app-delivery/issues/719)).

### Required


- [x] **Complete a [General Technical Review (GTR)](../../operations/project-reviews/general-technical-questions.md).**
  - GTR questionnaire completed in [cozystack/cozystack#3572](https://github.com/cozystack/cozystack/pull/3572); snapshot merged in [cncf/toc#2305](https://github.com/cncf/toc/pull/2305).

<!-- (Project assertion goes here) -->

- [x] **Complete a [Governance Review](../../operations/project-reviews/governance-review-template.md).**
  - Requested in [cncf/toc#2270](https://github.com/cncf/toc/issues/2270); the TOC sponsor completed the review directly in [## Governance and Maintainers](#governance-and-maintainers) below.

<!-- (Project assertion goes here) -->

- [x] **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

Repositories are hosted under the `cozystack` GitHub organization, images published to `ghcr.io/cozystack`, and the website hosted on Netlify. Community channels are CNCF Slack, Kubernetes Slack, Telegram, GitHub Discussions, and the cncf-cozystack-dev mailing list. The `etcd-operator` subproject was moved from the `aenix-io` organization into the `cozystack` organization, with its API group renamed from `etcd.aenix.io` to `etcd-operator.cozystack.io` (v1.6.0). Commercial support is listed neutrally at [cozystack.io/support](https://cozystack.io/support/), with two vendors presented equally, and the README links to that page rather than to any individual company.

Vendor-neutrality concerns raised in [cozystack/community#76](https://github.com/cozystack/community/issues/76) were fixed via [cozystack/website#690](https://github.com/cozystack/website/pull/690) (merged) and verified independently.

<!-- (TOC Evaluation goes here) -->

- [x] **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
  - Met and verified during DD. Cozystack is currently Sandbox, was accepted on 2025-03-02 via [cncf/sandbox#322](https://github.com/cncf/sandbox/issues/322), and the incubation application issue ([cncf/toc#1916](https://github.com/cncf/toc/issues/1916)) records acknowledgement of CNCF maturity expectations.

<!-- (TOC Evaluation goes here) -->

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

Documentation is published at [cozystack.io/docs](https://cozystack.io/docs/):

- [Getting-started tutorial](https://cozystack.io/docs/getting-started/)
- [Deployment guide](https://cozystack.io/docs/install/)
- [Use-case guides](https://cozystack.io/docs/guides/use-cases/)
- [Live demo](https://cozystack.io/demo/)

## Governance and Maintainers

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject if completed as a suggested item prior to application.

### Suggested

- [x] **Complete a Governance Review with the Project Reviews subproject**

Requested in [cncf/toc#2270](https://github.com/cncf/toc/issues/2270); the governance criteria in this section were evaluated directly by the TOC member.

<!-- (TOC Evaluation goes here) -->

- [x] **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

Governance documentation has evolved alongside the project: `GOVERNANCE.md` and `CONTRIBUTOR_LADDER.md` were introduced during Sandbox and have been revised since, with the Emeritus offboarding and reactivation process documented after it was first applied to two maintainers ([PR #2717](https://github.com/cozystack/cozystack/pull/2717)), and a new section on welcoming first-time contributors and CI-approval practices added to `GOVERNANCE.md` in [cozystack/cozystack#4845](https://github.com/cozystack/cozystack/pull/4845). Project assets also moved from the founding vendor's GitHub organization into the `cozystack` organization over the same period, alongside further communication-channel updates ([cozystack/website#744](https://github.com/cozystack/website/pull/744)).

Observation: revisions to `GOVERNANCE.md` since its creation have added new content: dependency and sub-project sections, maintainer lifecycle detail, and now contributor-welcoming/CI-approval guidance, rather than changing the core decision-making mechanics themselves.

- [x] **Clear and discoverable project governance documentation.**

[GOVERNANCE.md](https://github.com/cozystack/cozystack/blob/main/GOVERNANCE.md), [MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md), [CONTRIBUTOR_LADDER.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md) and [CONTRIBUTING.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTING.md) are present at the repository root and cross-link to each other and to the Code of Conduct. `GOVERNANCE.md` documents community roles, the maintainer lifecycle, and the project decision-making process.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Governance documents the maintainer lifecycle and decision-making process, and the maintainer roster in [MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md) is kept current, now spanning Active, Reviewer, and Emeritus sections as the roster has grown since application. Community meetings are held on a published schedule with recordings and notes.

[GOVERNANCE.md](https://github.com/cozystack/cozystack/blob/main/GOVERNANCE.md) defined a Directors role that was never filled or elected; removed via [cozystack/community#79](https://github.com/cozystack/community/issues/79).

<!-- (TOC Evaluation goes here) -->

- [ ] **Governance clearly documents [vendor-neutrality](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

[CODE_OF_CONDUCT.md](https://github.com/cozystack/cozystack/blob/main/CODE_OF_CONDUCT.md) contains a Vendor Neutrality Manifesto, but it's a statement of principles rather than a structural protection: `GOVERNANCE.md` has no vendor-neutrality statement or org-balance mechanism, with no cap on the proportion of the decision-making body affiliated to a single organization. Tracked in [cozystack/community#97](https://github.com/cozystack/community/issues/97); in progress under [cozystack/cozystack#4934](https://github.com/cozystack/cozystack/pull/4934).

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

[GOVERNANCE.md](https://github.com/cozystack/cozystack/blob/main/GOVERNANCE.md) documents leadership decisions (a defined path to becoming a maintainer; self-service or decided move to Emeritus) and contribution acceptance (contributions follow [CONTRIBUTING.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTING.md), with merge decisions the prerogative of the maintainers). A general decision-making process covers all project decisions: consensus of the maintainers, failing which a simple-majority vote with one vote per maintainer.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

Reviewers hold function-based responsibility for a specific project area under [CONTRIBUTOR_LADDER.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md) (roles detailed above), onboarded by pull request adding them to the owning file for one or more directories with approval by two existing approvers, and removed via inactivity thresholds, involuntary removal by majority vote, or the Stepping Down / Emeritus process. [SECURITY.md](https://github.com/cozystack/cozystack/blob/main/SECURITY.md) similarly defines a function-based security response team (see Security section below).

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

[CONTRIBUTOR_LADDER.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md) documents the full maintainer lifecycle. Onboarding requires nomination by a current Maintainer via pull request against `MAINTAINERS.md`, agreement by the nominee, and approval by a majority of current Maintainers. Offboarding is covered by inactivity thresholds, involuntary removal or demotion by majority vote, and a [Stepping Down / Emeritus process](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md#stepping-downemeritus-process) that specifies technical offboarding (access revocation, `CODEOWNERS` updates, credential rotation) within five business days.

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

The lifecycle has been exercised in both directions since Sandbox. New maintainers were onboarded during the v1.3 cycle with documented areas of responsibility. The Emeritus process was applied for the first time in [PR #2717](https://github.com/cozystack/cozystack/pull/2717), with the Active / Emeritus separation added to [MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md) at the same time, and the technical offboarding and reactivation process documented in `CONTRIBUTOR_LADDER.md` as a result of that experience. The Emeritus section has since grown further, and a Reviewers section has been added alongside it, as the ladder continues to be exercised.

Observation: Ænix holds 6 of 9 active maintainer seats (see [MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md)), below the 75% threshold that would trigger a TOC practice-period recommendation, but it reinforces the vendor-neutrality recommendation already tracked in [cozystack/community#97](https://github.com/cozystack/community/issues/97). Contribution volume is even more concentrated than seats, so broadening representation should mean more than adding seats, worth watching at Graduation.

- [ ] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

Subprojects inherit the parent project's governance, contribution, and maintainer-lifecycle rules: `GOVERNANCE.md` states Cozystack maintainers are maintainers of every sub-project by default, with per-repo `CODEOWNERS` identifying day-to-day owners (e.g. [cozystack/local-ccm#4](https://github.com/cozystack/local-ccm/pull/4)), and maturity status (Stable) is published for each packaged sub-project. The subproject listings also differed between the application and `GOVERNANCE.md`; fixed via [cozystack/community#80](https://github.com/cozystack/community/issues/80) and verified independently.

Missing: no documented process for adding or removing a subproject; in progress under [cozystack/cozystack#4933](https://github.com/cozystack/cozystack/pull/4933).

### Required

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

[MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md) lists 9 active, 4 Reviewer-tier, and 4 emeritus maintainers with name, GitHub username, company, and area of responsibility. GitHub usernames serve as the contact route, supported by the reporting paths documented in [SECURITY.md](https://github.com/cozystack/cozystack/blob/main/SECURITY.md).

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

[MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md) lists 9 active maintainers, appropriate to the project's size and scope.

- [x] **Code and Doc ownership in GitHub and elsewhere matches documented governance roles.**

[.github/CODEOWNERS](https://github.com/cozystack/cozystack/blob/main/.github/CODEOWNERS) matches the Reviewer/Maintainer model in [CONTRIBUTOR_LADDER.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md): a Reviewer nomination adds a GitHub username to `CODEOWNERS` for one or more paths and to the Reviewers section of `MAINTAINERS.md`. Ownership is defined for docs as well as code: `/docs/` and the root-level governance/community files (e.g. `GOVERNANCE.md`, `MAINTAINERS.md`, `SECURITY.md`) each have their own `CODEOWNERS` entry, distinct from the core code paths. [SECURITY.md](https://github.com/cozystack/cozystack/blob/main/SECURITY.md)'s named security response team are all listed as Active Maintainers in [MAINTAINERS.md](https://github.com/cozystack/cozystack/blob/main/MAINTAINERS.md).

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

[CODE_OF_CONDUCT.md](https://github.com/cozystack/cozystack/blob/main/CODE_OF_CONDUCT.md) adopts the CNCF Code of Conduct by reference, without modification. Raised in [cozystack/community#78](https://github.com/cozystack/community/issues/78), fixed via org-level default files in [cozystack/.github#8](https://github.com/cozystack/.github/pull/8) (merged) and verified independently.

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

Linked from [GOVERNANCE.md](https://github.com/cozystack/cozystack/blob/main/GOVERNANCE.md) and [CONTRIBUTOR_LADDER.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md).

- [x] **All subprojects, if any, are listed.**

All subprojects are listed in [GOVERNANCE.md](https://github.com/cozystack/cozystack/blob/main/GOVERNANCE.md); this was fixed via [cozystack/community#80](https://github.com/cozystack/community/issues/80) after the application's subproject list and `GOVERNANCE.md`'s governed repositories were found not to align, and verified independently.

## Contributors and Community

Note: this section may be augmented by the completion of a Governance Review from the Project Reviews subproject.

### Suggested

- [x] **Contributor ladder with multiple roles for contributors.**

[CONTRIBUTOR_LADDER.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md) defines four roles (Community Participant, Contributor, Reviewer, Maintainer), each with documented responsibilities, requirements, and privileges. Requirements for advancement are quantified and measurable against project contribution data.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

[CONTRIBUTING.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTING.md) documents the contribution process, and the repository README routes contributors by change type: bug reports to GitHub Issues after checking for duplicates, feature requests to an issue, cross-cutting or architectural changes to a design proposal in the [community repository](https://github.com/cozystack/community/tree/main/design-proposals), and questions or early ideas to GitHub Discussions. A commit convention covering type, scope, and sign-off is enforced by a commit-trailers CI check, and the pull request template requires a release note, a downstream-repository assessment, and screenshots for user-interface changes. The project website's [community page](https://cozystack.io/community/) links contributor onboarding, good first issues, the developer guide, design proposals, and the contributor ladder.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

The [community page](https://cozystack.io/community/) documents five public channels with their intended purpose: [Telegram](https://t.me/cozystack) for day-to-day questions, [#cozystack on CNCF Slack](https://cloud-native.slack.com/archives/C08BQJD95J7) and [on Kubernetes Slack](https://kubernetes.slack.com/archives/C06L3CPRVN1), [GitHub Discussions](https://github.com/cozystack/cozystack/discussions) for usage questions and early ideas, and the [cncf-cozystack-dev mailing list](https://lists.cncf.io/g/cncf-cozystack-dev) for announcements and meeting-notes access. The same channels are linked from the repository README.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.). List any non-public communications channels and what their special purpose is.**

The [community page](https://cozystack.io/community/) provides a single inventory of all project communication channels with the purpose of each: Telegram, #cozystack on CNCF Slack and Kubernetes Slack, GitHub Discussions, and the [cncf-cozystack-dev mailing list](https://lists.cncf.io/g/cncf-cozystack-dev). Subprojects do not operate separate channels; discussion routes through the parent project's channels. The project states that no non-public channels are used for project work, with the exception of private vulnerability reporting via GitHub Private Vulnerability Reporting, documented in [SECURITY.md](https://github.com/cozystack/cozystack/blob/main/SECURITY.md).

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

The project holds a public community meeting on a published schedule: see the [community page](https://cozystack.io/community/) for details, [community_meeting.md](https://github.com/cozystack/community/blob/main/community_meeting.md) for how it runs, and the [meeting archive](https://github.com/cozystack/community/blob/main/meetings/README.md) for past recordings.

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

Contribution documentation is layered by audience: the [community page](https://cozystack.io/community/) for orientation and non-code contribution paths, [contributor onboarding](https://github.com/cozystack/community/blob/main/contributors/README.md) covering how to pick a task, what environment is needed, and how to get a first pull request merged, [CONTRIBUTING.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTING.md) for process and routing, and the [developer guide](https://cozystack.io/docs/development/) for local builds and end-to-end testing. [Design proposals](https://github.com/cozystack/community/tree/main/design-proposals) provide a documented path for architectural and cross-cutting changes, and [CONTRIBUTOR_LADDER.md](https://github.com/cozystack/cozystack/blob/main/CONTRIBUTOR_LADDER.md) documents progression through project roles. This documentation has expanded during the project's time in Sandbox, with the contributor ladder, design-proposal process, and contributor onboarding guide all added over that period.

- [x] **Demonstrate contributor activity and recruitment.**

Unique contributors by first commit have grown year over year, matched again in 2026 with months left in the year, per the project's [OSS health snapshot](https://cozystack.io/oss-health/devstats/), [DevStats dashboard](https://cozystack.devstats.cncf.io/), and [LFX Insights](https://insights.linuxfoundation.org/project/cozystack/contributors). A [contributor onboarding guide](https://github.com/cozystack/community/blob/main/contributors/README.md) and a `good first issue` label support this.

<!-- (TOC Evaluation goes here) -->

## Engineering Principles

### Suggested

- [x] **Roadmap change process is documented.**

Feature requests are opened as GitHub issues; cross-cutting or architectural changes go through a design proposal in the [community repository](https://github.com/cozystack/community/tree/main/design-proposals) (the design-proposals `README.md` frames this as encouraged rather than mandatory). A `triage/needs-proposal` label on `cozystack/cozystack` routes issues that need one. The community repository holds 12 design proposals to date. Design proposals are a standing agenda item at the [community meeting](https://github.com/cozystack/community/blob/main/community_meeting.md), presented as a linked PR ahead of the call.

<!-- (TOC Evaluation goes here) -->

- [x] **History of regular, quality releases.**

The [release history](https://github.com/cozystack/cozystack/releases) shows a predictable cadence: 7 minor lines (v1.0–1.6) at roughly monthly intervals since February 2026, backed by ~38 patch releases, with older lines (e.g. v1.0.x) still receiving patches concurrently with newer ones, evidence of maintained backport branches rather than a single moving line. [OpenSSF Best Practices](https://www.bestpractices.dev/en/projects/10177) is passing at 100%.

Observation: the project's own GTR self-disclosed a period of unreliable scheduled nightly end-to-end test runs. Verified independently, this has since trended positive, with the most recent scheduled runs passing consistently.

<!-- (TOC Evaluation goes here) -->

### Required

- [x] **Document project goals and objectives that illustrate the project's differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

The project documents its goals and positioning in the repository README and [guides section](https://cozystack.io/docs/guides/), covering three distinct use cases: building a public cloud, a private cloud via Infrastructure-as-Code, and use as a Kubernetes distribution for bare metal ([use-case guides](https://cozystack.io/docs/guides/use-cases/)). A General Technical Review was also completed, merged in-tree at [GENERAL_TECHNICAL_REVIEW.md](https://github.com/cozystack/cozystack/blob/main/GENERAL_TECHNICAL_REVIEW.md) ([cozystack/cozystack#3572](https://github.com/cozystack/cozystack/pull/3572)); the snapshot was merged into `cncf/toc` via [cncf/toc#2305](https://github.com/cncf/toc/pull/2305).

<!-- (TOC Evaluation goes here) -->

- [x] **Document what the project does, and why it does it - including viable cloud native use cases.**

The [Introduction](https://cozystack.io/docs/introduction/) documents what Cozystack is (a Kubernetes-based framework for building a private cloud) and why: a tenant system for multi-team/multi-company isolation, built-in replicated storage, VMs running natively in Kubernetes alongside containers, and managed databases without a separate virtualization layer. The three [use-case guides](https://cozystack.io/docs/guides/use-cases/) (public cloud (PaaS for multiple customers), private cloud via Infrastructure-as-Code, and Kubernetes distribution for bare metal) ground these features in concrete cloud native scenarios.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The project publishes a [roadmap page](https://cozystack.io/docs/roadmap/) in its documentation, pointing to a public [GitHub project board](https://github.com/orgs/cozystack/projects/1) and routing roadmap questions to GitHub Discussions or issues; the board is genuinely active, with completed items tracing to shipped features alongside open draft items for forward planning. Forward-looking planning is additionally tracked through [design proposals](https://github.com/cozystack/community/tree/main/design-proposals) in the community repository, which are discussed at the community meeting alongside releases.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation. _This requirement may also be satisfied by completing a General Technical Review._**
  - A General Technical Review was completed and merged in-tree at [GENERAL_TECHNICAL_REVIEW.md](https://github.com/cozystack/cozystack/blob/main/GENERAL_TECHNICAL_REVIEW.md) (2026-08-12).

- [x] **Document the project's release process.**

The release process is documented in-repo at [docs/release.md](https://github.com/cozystack/cozystack/blob/main/docs/release.md) and linked from the public docs' [upgrade guide](https://cozystack.io/docs/operations/cluster/upgrade/). It covers the three release types (release candidate, regular, patch), the branch-freeze model (`release-X.Y` is cut at the first RC, closing the line to everything but cherry-picks/backports), and the promotion flow (a regular release is the exact RC image retagged by digest, never rebuilt), matching the concurrent multi-line patch activity already observed in the tag history.

<!-- (TOC Evaluation goes here) -->

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security and Compliance if completed as a suggested item prior to application.

### Suggested

- [ ] **Complete a joint security assessment with TAG Security and Compliance**

A self-assessment exists in-tree at [docs/security/self-assessment.md](https://github.com/cozystack/cozystack/blob/main/docs/security/self-assessment.md), but no issue has been filed against [cncf/toc](https://github.com/cncf/toc) for a joint assessment. Required security criteria below were evaluated directly by the TOC member.

<!-- (TOC Evaluation goes here) -->

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

[SECURITY.md](https://github.com/cozystack/cozystack/blob/main/SECURITY.md) documents private reporting via GitHub Private Vulnerability Reporting (with named maintainer fallback), response/triage targets, and a coordinated-disclosure window. Demonstrated in practice by [GHSA-g883-q79m-8225](https://github.com/cozystack/cozystack/security/advisories/GHSA-g883-q79m-8225), patched in v1.3.3. A finding from a third-party penetration test reported by an adopter was also patched in about a day, a positive responsiveness signal.

<!-- (TOC Evaluation goes here) -->

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

A branch ruleset requires reviewed PRs on the default branch, [.github/CODEOWNERS](https://github.com/cozystack/cozystack/blob/main/.github/CODEOWNERS) assigns per-path ownership, and CI was migrated from long-lived PATs to short-lived `cozystack-ci` GitHub App tokens ([#2351](https://github.com/cozystack/cozystack/pull/2351), [#2383](https://github.com/cozystack/cozystack/pull/2383)). [docs/release.md](https://github.com/cozystack/cozystack/blob/main/docs/release.md) also describes tag-protection rulesets restricting `v*` tag creation/deletion to the CI app. Org-wide 2FA enforcement was confirmed directly with the Cozystack maintainers.

Observation: per [SECURITY.md](https://github.com/cozystack/cozystack/blob/main/SECURITY.md) and [docs/security/self-assessment.md](https://github.com/cozystack/cozystack/blob/main/docs/security/self-assessment.md)'s own disclosure, CodeQL and DCO are not among the required status checks on `main`; in progress under [cozystack/cozystack#4932](https://github.com/cozystack/cozystack/pull/4932).

- [x] **Document assignment of security response roles and how reports are handled.**

[SECURITY.md](https://github.com/cozystack/cozystack/blob/main/SECURITY.md) names five maintainers as the security response team for reported vulnerabilities, with the handling flow (private report → acknowledge → triage → coordinated disclosure) already demonstrated in practice. A rotating Security Champion role owns day-to-day security process ownership (CVE-triage clock, scan review, escalation), rotating every six months with a one-week handoff overlap; the current holder is named in `SECURITY.md`.

- [x] **Document Security Self-Assessment.**

A self-assessment exists in-tree at [docs/security/self-assessment.md](https://github.com/cozystack/cozystack/blob/main/docs/security/self-assessment.md). Snapshot requested in [cozystack/community#81](https://github.com/cozystack/community/issues/81), merged via [cncf/toc#2306](https://github.com/cncf/toc/pull/2306).


- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The project holds an OpenSSF Best Practices passing badge ([project 10177](https://www.bestpractices.dev/projects/10177)). The badge is displayed in the site footer across cozystack.io, and the project publishes the current badge state and verification status at [cozystack.io/oss-health/openssf](https://cozystack.io/oss-health/openssf/). Security practice metrics are also tracked externally at [LFX Insights](https://insights.linuxfoundation.org/project/cozystack/security).

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

[ADOPTERS.md](https://github.com/cozystack/cozystack/blob/main/ADOPTERS.md) lists 8 organizations with contact, date added, and a description of use for each, stating that entries are running Cozystack in production. Entries span February 2024 to December 2025. The documented use cases cover consultancy, hosting providers, sovereign and public clouds, GPU and educational environments, and use as a Kubernetes distribution. Deployment scale is not published; scale and production status are addressed through the adopter interviews below.


- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The publicly documented adopter list in [ADOPTERS.md](https://github.com/cozystack/cozystack/blob/main/ADOPTERS.md) above lists 8 organizations at production use, satisfying this criterion on its own. Three adopters were additionally interviewed for deeper verification, summarized below under **TOC verification of adopters**; all three are presented anonymized.

- [x] **TOC verification of adopters.**

**Complete:** Three anonymized adopter interviews were completed, satisfying TOC verification of adopters for this review; the three-adopter minimum itself is independently met by the public adopter list above. Interview notes are kept privately for internal TOC use and verification; [#### Adoption](#adoption) below lists interview status only.

**Findings (three interviews):** usage ranged from pre-production (a scientific-computing centre) to a live multi-site deployment (a university research institute) to commercial production since early 2025 (a regional cloud and managed-services provider). No adopter reported an upstream blocker.

**Strengths adopters cited:**

1. **One integrated, open platform** — adopters describe it as a "cloud in a box": virtual machines, Kubernetes-as-a-service and managed databases behind a single control plane, built from well-known CNCF components so operators can understand and debug every layer themselves rather than working with a black box. One adopter found it accessible enough that a colleague without dedicated DevOps experience completed a full deployment within a couple of days; another runs a large multi-tenant footprint across multiple datacenters with a very small operations team.
2. **Data sovereignty and freedom from vendor lock-in** — the main adoption driver for all three, and decisive in two cases following licensing changes elsewhere in the market.
3. **Multi-tenancy** — adopters describe the tenant model as more mature than anything they could have built themselves, and two had abandoned in-house attempts before adopting. Isolation held up in every deployment, including where tenants are unrelated paying customers.
4. **Self-service provisioning** — cuts out manual operational work, with users creating clusters, databases and virtual machines without raising a ticket.
5. **Cross-industry applicability** — one adopter runs a commercial multi-tenant cloud serving customers across industries, while also deploying on-premise for a public-sector IT department and fielding interest from regulated-industry organizations, evidence the platform fits beyond a single vertical.
6. **Responsive, hands-on maintainers** — a requested feature implemented within days, a reported security finding patched in about a day, proactive notification of security-relevant issues, and help getting an unsupported hardware configuration working.
7. **Upgrade safety** — no adopter reported an upgrade that broke their deployment, across cadences from days after release to a validated two-to-three-month cycle.

**Improvements adopters would like to see:**

1. **Day-2 operations documentation** — troubleshooting guides and known failure modes, particularly for storage, backups and stuck tenant clusters.
2. **API stability and deprecation policy** — the most consistent ask, independently raised by the project's own General Technical Review and tracked as a non-blocking recommendation above.
3. **Clearer upgrade notes** between minor versions.

Raised by a single adopter: per-tenant backup-health observability, tenant-definable firewall rules, and clearer entry-level guidance that a single-machine deployment is a legitimate way to start, which the project has since addressed.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The [platform stack reference](https://cozystack.io/docs/guides/platform-stack/) documents every integrated component with its role, upstream source, and license. CNCF projects integrated include Kubernetes, KubeVirt, Cilium, Kube-OVN, Flux, cert-manager, Velero, Harbor, External Secrets Operator, MetalLB, CoreDNS, ExternalDNS, Cluster API and HAMi; non-CNCF integrations include Talos Linux, Kamaji, LINSTOR, Keycloak, ClickHouse, OpenSearch, Kafka, NATS, VictoriaMetrics and Grafana. The project additionally documents a path for [adding applications from external sources](https://cozystack.io/docs/applications/external/) to the catalog, with a reference example repository, and publishes a [Kubernetes Conformance](https://cozystack.io/compliance/kubernetes-conformance/) page.

<!-- (TOC Evaluation goes here) -->

#### Adoption

Interview findings are summarized under [TOC verification of adopters](#toc-verification-of-adopters) in [## Ecosystem](#ecosystem). All interview notes are kept privately for internal TOC use and verification.

##### Adopter 1 - Scientific research and higher-education computing centre

**Status:** Interview complete (September 2026). Anonymized.

##### Adopter 2 - Higher education and research

**Status:** Interview complete (October 2026). Anonymized.

##### Adopter 3 - Sovereign regional cloud and managed-services provider

**Status:** Interview complete (October 2026). Anonymized.
