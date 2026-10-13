# NATS Graduation Due Diligence

- Link to [Graduation application issue](https://github.com/cncf/toc/issues/2042)
- Link to [project-maintained General Technical Review](https://github.com/nats-io/nats-general/blob/main/GENERAL-TECHNICAL-REVIEW.md)
- Link to [project-side remediation tracker](https://github.com/nats-io/nats-general/issues/106)
- Link to [NATS governance](https://github.com/nats-io/nats-general/blob/main/GOVERNANCE.md)
- Link to [NATS maintainer and project inventory](https://github.com/nats-io/nats-general/blob/main/MAINTAINERS.md)
- Link to [NATS security policy](https://github.com/nats-io/nats-general/blob/main/SECURITY.md)


## Graduation Evaluation Summary for NATS

### Criteria Evaluation

Karena Angell conducted the due diligence of NATS, which applied for Graduation. NATS is a long-running CNCF incubating project with evidence of technical maturity, broad production adoption, a complete technical review, a security self-assessment, a public roadmap, regular releases, and a substantial maintainer base.

The project has a clear technical scope for high-performance cloud native messaging; a lightweight deployment model that works from edge to multi-cloud environments; operational documentation; an organization-balanced voting model for vendor neutrality; and a public adopter base across AI, automotive, energy, financial services, retail, technology, and telecom.

#### Project notables

- **Broad technical scope:** NATS covers subject-based messaging, pub/sub, request/reply, queue groups, service communication, and JetStream persistence/streaming in one project scope, with deployment documentation from local development to clustered, supercluster, leaf-node, edge, and Kubernetes environments.
- **Project-maintained operations ecosystem:** The project maintains Kubernetes resources and Helm charts in [`nats-io/k8s`](https://github.com/nats-io/k8s), controllers in [`nats-io/nack`](https://github.com/nats-io/nack), and tooling including [`natscli`](https://github.com/nats-io/natscli), [`nats-surveyor`](https://github.com/nats-io/nats-surveyor), and [`prometheus-nats-exporter`](https://github.com/nats-io/prometheus-nats-exporter), making deployment, operations, and monitoring assets part of the project-maintained ecosystem.
- **Client and adoption breadth:** NATS maintains official clients for Go, JavaScript / TypeScript, Python, Rust, C, Java, .NET, Ruby, and Swift. Public materials list 100+ adopters across AI, automotive, energy, financial services, retail, technology, and telecom, indicating use across multiple implementation languages and industry contexts.
- **Maintainer and contribution data:** `nats-general/MAINTAINERS.md` is consistent with the maintainer activity shown in [LFX Insights](https://insights.linuxfoundation.org/project/nats/contributors?timeRange=past365days&start=2025-09-29&end=2026-09-29) for the past 365 days. LFX reports 2 organizations accounting for 67% of contributions and 6 contributors accounting for 55%, with Synadia at 50.39% on the organization leaderboard. This provides both a maintainer-count cross-check and contribution-concentration data to evaluate against the project's organization-balanced voting model.
- **Documentation, governance, and tooling context:** Project-controlled resources include docs, examples, release guidance, architecture/design records, organization-balanced voting rules, and repository-specific `CLAUDE.md` files for AI-assisted tooling context. These are notable because operational guidance, decision context, governance rules, and AI/tooling guardrails are published in project repositories rather than only in external or informal channels.

Project-side blocking remediation has been completed and all remediation tracking issues are closed in [`nats-io/nats-general#106`](https://github.com/nats-io/nats-general/issues/106). All project-side blocking remediation issues `#106` through `#116` are closed as of 2026-10-13:

- **Vendor-neutrality cleanup:** docs copyright was updated in [`nats.docs.v2#98`](https://github.com/nats-io/nats.docs.v2/pull/98); roadmap input and nightly-image language were updated in [`nats-site#1112`](https://github.com/nats-io/nats-site/pull/1112); support page vendor listing was converted to a neutral vendor directory with inclusion criteria in [`nats-site#1114`](https://github.com/nats-io/nats-site/pull/1114); and the download page now explains why nightly images are currently published under the `synadia/nats-server` Docker Hub namespace in [`nats-site#1117`](https://github.com/nats-io/nats-site/pull/1117). Tracked and closed in [`nats-io/nats-general#107`](https://github.com/nats-io/nats-general/issues/107).
- **Maintainer metadata:** [`nats-general#117`](https://github.com/nats-io/nats-general/pull/117) converted `MAINTAINERS.md` to a table with maintainer names, contacts, GitHub handles, and affiliations. CNCF `.project/maintainers.yaml` remains handle-only by design; CNCF maintainer affiliation data is sourced through OpenProfile / LFX PCC automation rather than custom fields in `.project`. Tracked and closed in [`nats-io/nats-general#108`](https://github.com/nats-io/nats-general/issues/108).
- **Maintainer lifecycle:** [`nats-general#118`](https://github.com/nats-io/nats-general/pull/118) added active maintainer role expectations, emeritus status, return from emeritus, affiliation changes, voluntary and involuntary offboarding, and access / ownership transition language. Tracked and closed in [`nats-io/nats-general#109`](https://github.com/nats-io/nats-general/issues/109).
- **Asynchronous governance records:** [`nats-general#119`](https://github.com/nats-io/nats-general/pull/119) documents that NATS governance is asynchronous, that proposals/discussions/decisions happen publicly on GitHub, and that meeting/Slack/mailing-list outcomes only take effect when recorded in GitHub issues or PRs. Tracked and closed in [`nats-io/nats-general#111`](https://github.com/nats-io/nats-general/issues/111).
- **Subproject / repository scope:** [`nats-general#119`](https://github.com/nats-io/nats-general/pull/119) added subproject governance, addition/removal, acceptance, transition, and archival process; [`nats-general#120`](https://github.com/nats-io/nats-general/pull/120) added a project/subproject inventory with repository, language, category/domain, maturity/support status, and domain owner/maintainers. Tracked and closed in [`nats-io/nats-general#110`](https://github.com/nats-io/nats-general/issues/110).
- **Security response role documentation:** [`nats-general#121`](https://github.com/nats-io/nats-general/pull/121) added a canonical NATS security policy in `nats-general/SECURITY.md`, including reporting channels, supported versions, response timeline, disclosure policy, security response team membership/access, appointment/removal model, and conflicts-of-interest handling. Tracked and closed in [`nats-io/nats-general#112`](https://github.com/nats-io/nats-general/issues/112).
- **Security policy discoverability:** [`nats-general#121`](https://github.com/nats-io/nats-general/pull/121) added the canonical security policy; [`nats-io/.project#6`](https://github.com/nats-io/.project/pull/6) updated `.project/project.yaml` to the live `nats-general/SECURITY.md` path; and [`nats-server#8744`](https://github.com/nats-io/nats-server/pull/8744) added a `nats-server/SECURITY.md` shim that points to the canonical policy and the `nats-server` private advisory reporting flow. Tracked and closed in [`nats-io/nats-general#113`](https://github.com/nats-io/nats-general/issues/113) and [`nats-io/nats-general#116`](https://github.com/nats-io/nats-general/issues/116).
- **Repository hygiene baseline:** [`nats-io/nats-general#114`](https://github.com/nats-io/nats-general/issues/114) was filed and closed as complete.
- **`.project` repository-scope metadata alignment:** [`nats-io/.project#5`](https://github.com/nats-io/.project/pull/5) expanded the `.project/project.yaml` repository list to match the documented NATS project/subproject inventory, and [`nats-io/.project#6`](https://github.com/nats-io/.project/pull/6) updated the security policy link to the canonical NATS security policy. Tracked and closed in [`nats-io/nats-general#115`](https://github.com/nats-io/nats-general/issues/115).

The following are non-blocking recommendations for the project:

- Add a short governance history section or link to notable governance changes.
- Make decision records easier to discover through governance issues, vote PRs, meeting notes, or an equivalent public decision log.
- Continue improving secure-by-default deployment guidance, especially around TLS/mTLS and production Kubernetes deployments.
- Confirm and document org-level 2FA / branch-protection posture where possible, or enable stronger 2FA controls if not already enforced. This is a non-blocking recommendation because the public contribution, governance, and security docs already require maintainer review and access-control practices sufficient for this DD.
- Generate or publish SBOMs automatically for release artifacts if not already covered by current release automation.

### Adoption Evaluation

Adopter verification is complete for this Due Diligence through three anonymized production-adopter interviews. The interviewees represented independent organizations using NATS in large-scale production environments, including cloud/edge services, multi-region and multi-cloud messaging infrastructure, and on-premises Kubernetes-based network/service platforms.

Across the three interviews, adopters consistently described NATS as production-critical infrastructure. Reported uses included pub/sub messaging, control-plane/event signaling, work queues, telemetry, key-value or persistence capabilities, MQTT-related workflows, and JetStream-backed operational patterns. The interviews support the conclusion that NATS is used in appropriate production capacity by at least three independent adopters.

Common strengths reported by adopters:

- **Simplicity and operational leverage:** adopters valued getting messaging, persistence, key-value patterns, and related communication primitives from one lightweight system.
- **Reliability and resilience:** adopters reported running NATS as highly available or zero-downtime infrastructure across regions, cloud environments, Kubernetes clusters, edge platforms, or on-premises deployments.
- **Cloud-native deployment fit:** adopters reported successful use with Kubernetes, Helm, GitOps workflows, canary/rolling upgrades, and automated deployment pipelines.
- **Open-source confidence:** adopters explicitly valued that core functionality is available in the open project and is not dependent on a proprietary edition.
- **Responsive community and maintainers:** adopters cited positive experiences with the community and rapid upstream issue resolution.

Common areas for improvement reported by adopters:

- Improve documentation on high availability boundaries, especially what is and is not included in the HA set for specific configurations.
- Improve credential rotation and related tooling, including integrations with common secret-management systems.
- Make secure deployment patterns easier to adopt, including decentralized credential/root-certificate management and compliance-oriented guidance.
- Consider additional ecosystem integrations requested by adopters, such as stronger database/global-consistency patterns, industrial protocol support, and options attractive for embedded or supply-chain-sensitive environments.

### Final Assessment

The TOC has found the project to have satisfied the criteria for Graduation after completion and closure of project-side remediation tracked in [`nats-io/nats-general#106`](https://github.com/nats-io/nats-general/issues/106).

### Criteria

## Application Process Principles

### Suggested

- [ ] **Engage with domain-specific TAG(s) to present the technical architecture of the project.**
  - A domain-specific TAG presentation was not identified during this draft review. The project completed a General Technical Review, which provides the technical architecture and operational review material needed for DD.

### Required

- [x] **Complete a [General Technical Review (GTR)](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/general-technical-questions.md).**
  - The [project-maintained GTR](https://github.com/nats-io/nats-general/blob/main/GENERAL-TECHNICAL-REVIEW.md) was most recently revised on 2026-01-19 in `nats-io/nats-general` and was reviewed as part of DD.

The GTR covers project scope, personas, use cases, roadmap process, architecture, IAM, sovereignty considerations, HA, resource requirements, API design, installation, security, rollout, observability, dependencies, troubleshooting, compliance, and Day-2 operations.

- [x] **Complete a [Governance Review](https://github.com/cncf/toc/blob/main/toc_subprojects/project-reviews-subproject/governance-review-template.md).**

The governance review was completed during TOC DD and verified against the project's public governance, maintainer, security, and remediation materials. The initial status was **Mostly Satisfactory**. Subsequent remediation PRs addressed maintainer lifecycle, public/asynchronous governance records, subproject lifecycle, maintainer affiliation/domain information, and security response role assignment. No local-only review artifact is required as evidence in this DD.

- [x]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

Met after remediation. NATS governance documents organization-balanced voting: maintainers from the same company or organization collectively receive one organization vote, and independent maintainers receive their own vote. Maintainer affiliations are now listed in `MAINTAINERS.md`; affiliation changes are documented in governance; and public website/support/docs surfaces have been updated to reduce Synadia-specific presentation.

Completed during due diligence:

- [`nats-io/nats-general#107`](https://github.com/nats-io/nats-general/issues/107) — vendor-neutrality cleanup for project website/support/events/docs surfaces
- [`nats-io/nats-general#108`](https://github.com/nats-io/nats-general/issues/108) — maintainer metadata with affiliation, org-vote group, and domain ownership
- [`nats-io/nats-general#115`](https://github.com/nats-io/nats-general/issues/115) — `.project` metadata alignment with graduation scope

- [x] **Review and acknowledgement of expectations for [Sandbox](https://sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**
   - [x] Met and verified during DD. NATS is currently Incubating, was accepted on 2018-03-15 via https://github.com/cncf/toc/pull/81, and the graduation issue records acknowledgement of CNCF maturity expectations.

- [x] **Due Diligence Review.**

Completion of this due diligence document, resolution of concerns raised, and presentation for public comment satisfies the Due Diligence Review criteria.

- [x] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

NATS has extensive public documentation at https://docs.nats.io, including installation, configuration, security, monitoring, JetStream, Kubernetes, developer, and reference protocol documentation. The project also publishes examples at https://natsbyexample.com/ and maintains architecture / decision records at https://github.com/nats-io/nats-architecture-and-design.

## Governance and Maintainers

Note: this section is augmented by the governance review performed during TOC DD and verified against public project governance materials.

### Suggested

- [x] **Complete a Governance Review with the Project Reviews subproject**

A governance review has been completed by the TOC member conducting the due diligence. Status: **Mostly Satisfactory** before remediation. The project has completed the blocking remediation issues identified through that review.

- [x]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

Due diligence verified that NATS governance has evolved from its initial governance model to the current organization-based voting model through public governance documentation and remediation PRs. A short governance history or links to notable governance changes remains an open documentation (non-blocking) recommendation.

### Required

- [x] **Clear and discoverable project governance documentation.**

Governance is available at https://github.com/nats-io/nats-general/blob/main/GOVERNANCE.md. The maintainer list is available at https://github.com/nats-io/nats-general/blob/main/MAINTAINERS.md.

- [x] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

Met after remediation. Governance documents maintainer elections, voting thresholds, contribution acceptance, and CNCF request paths. [`nats-general#119`](https://github.com/nats-io/nats-general/pull/119) documents NATS' actual asynchronous governance model: governance proposals, discussion, votes, and decision records are recorded in public GitHub issues and pull requests; technical design decisions are recorded in `nats-architecture-and-design`; and meeting, Slack, or mailing-list outcomes are not final until recorded on GitHub. Tracked in [`nats-io/nats-general#111`](https://github.com/nats-io/nats-general/issues/111).

- [x] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

Met after remediation. NATS uses organization-balanced voting so no single organization receives more than one organization vote regardless of maintainer count. Independent maintainers receive their own vote. Governance now states that a maintainer's organization is the affiliation listed in `MAINTAINERS.md` when a vote opens, documents how affiliation changes are handled, and requires subproject additions not to introduce dependencies, ownership arrangements, or other restrictions that would compromise vendor-neutral governance.

- [x] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

Met. Governance documents consensus-first decision making, formal votes when needed, maintainer election/removal by 2/3 organization vote, contribution acceptance by maintainers for the affected repo/component, and CNCF requests via `info@nats.io`. Maintainers are responsible for project direction, PR review, releases, and component maintenance responsibilities. Governance changes require a formal vote.

- [x] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

Met after remediation. Maintainer assignment/removal is documented in governance. `nats-general/SECURITY.md` now documents the security response team, reporting paths, who receives GitHub private vulnerability reports and `security@nats.io`, appointment/removal model, disclosure expectations, and advisory publication path. Remediated by [`nats-general#121`](https://github.com/nats-io/nats-general/pull/121) and tracked in [`nats-io/nats-general#112`](https://github.com/nats-io/nats-general/issues/112).

- [x] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

Met after remediation. Governance documents active maintainer role expectations, maintainer nomination/election by formal vote, resignation and removal, emeritus status, return from emeritus to active maintainer status, affiliation changes, voluntary offboarding, removal/offboarding, and access / ownership transition. Remediated by [`nats-general#118`](https://github.com/nats-io/nats-general/pull/118) and tracked in [`nats-io/nats-general#109`](https://github.com/nats-io/nats-general/issues/109).

- [x] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

Met. DD verified closed `nats-general` maintainer-vote issues/PRs exercising maintainer lifecycle updates, and the maintainer list has been actively maintained.

- [x] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

Met after remediation. [`MAINTAINERS.md`](https://github.com/nats-io/nats-general/blob/main/MAINTAINERS.md) lists current maintainers with names, contact emails, GitHub handles, and affiliations. The same file now includes a project/subproject inventory with repository, language, category/domain, maturity/support status, and domain owner/maintainers. [LFX Insights](https://insights.linuxfoundation.org/project/nats/contributors?timeRange=past365days&start=2025-09-29&end=2026-09-29) matches the public maintainer count. CNCF `.project/maintainers.yaml` remains handle-only by design; CNCF affiliation metadata is sourced through OpenProfile / LFX PCC automation. Tracked in [`nats-io/nats-general#108`](https://github.com/nats-io/nats-general/issues/108).

- [x] **A number of active maintainers which is appropriate to the size and scope of the project.**

Met. The maintainer list is substantial and appears appropriate for the size and scope of the project and activity verified during due diligence.

- [x] **Project maintainers from at least 2 organizations that demonstrates survivability.**

Met. `MAINTAINERS.md` identifies maintainers from Synadia, Striveworks, NanoNative, Pinetec Limited, Kajabi, NN Group, Salesforce, Workiva, Luxant Solutions, Essential Personnel, and independent maintainers; this activity has been confirmed during due diligence.

- [x] **Code and Doc ownership in GitHub and elsewhere matches documented governance roles.**

Met for public due diligence purposes. Governance describes maintainer authority, PR approval, subproject ownership, and offboarding/access transition. `MAINTAINERS.md` now includes repository/domain owner-maintainer assignments across project and documentation repositories. Public CODEOWNERS / permission evidence is limited by GitHub visibility, but the public governance and maintainer inventory now align.

- [x] **Document adoption and adherence to the CNCF Code of Conduct or the project's CoC which is based off the CNCF CoC and not in conflict with it.**

Met. Governance states that NATS follows the CNCF Code of Conduct.

- [x] **CNCF Code of Conduct is cross-linked from other governance documents.**

Met. `GOVERNANCE.md` references the CNCF Code of Conduct.

- [x] **All subprojects, if any, are listed.**

Met after remediation. DD verified the server, official client libraries, Kubernetes-related repositories, monitoring repositories, and additional tooling in scope through public project inventory and repository review. `MAINTAINERS.md` now provides a durable project/subproject inventory with repository, language, category/domain, maturity/support status, and domain owner/maintainers. Tracked in [`nats-io/nats-general#110`](https://github.com/nats-io/nats-general/issues/110).

Licensing review: the graduation-scope repositories each publish an Apache License 2.0 `LICENSE` file. This is consistent with CNCF expectations for project licensing.

- [x] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**

Met after remediation. The DD includes a provisional maturity/support assessment for each graduation-scope repository. `MAINTAINERS.md` now connects repositories to domain owner/maintainers and maturity/support status. Governance now documents subproject addition/removal, community-contributed project acceptance criteria, review and acceptance, transition into the project, and repository archival. Remediated by [`nats-general#119`](https://github.com/nats-io/nats-general/pull/119) and [`nats-general#120`](https://github.com/nats-io/nats-general/pull/120); tracked in [`nats-io/nats-general#110`](https://github.com/nats-io/nats-general/issues/110).

## Contributors and Community

Note: this section is augmented by the governance review performed during TOC DD and verified against public project governance materials.

### Suggested

- [ ] **Contributor ladder with multiple roles for contributors.**

Partially met. NATS documents contributors and maintainers, but a richer contributor ladder with intermediate roles is not clearly described. This is suggested rather than required for graduation.

### Required

- [x] **Clearly defined and discoverable process to submit issues or changes.**

Met. The project publishes contribution guidance at https://nats.io/contributing/ and repository-specific guidance such as https://github.com/nats-io/nats-server/blob/main/CONTRIBUTING.md. The `nats-server` guide describes opening issues for enhancements/bugs, PR expectations, sign-off requirements, test expectations, and review flow.

- [x] **Project must have, and document, at least one public communications channel for users and/or contributors.**

Met. Public channels include Slack (https://slack.nats.io/), GitHub Discussions on respective repositories, Google Groups (https://groups.google.com/g/natsio), YouTube (https://www.youtube.com/@NATS_io), and public issue trackers.

- [x] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

Met. DD verified public communication channels including Slack, GitHub Discussions, Google Groups, YouTube, and email. Subproject communication occurs through the same public project channels and through repository-specific GitHub issues, discussions, and pull requests. `security@nats.io` is documented as the non-public security reporting channel.

- [x] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

Met through documented asynchronous governance rather than a standing public meeting cadence. Governance now states that NATS has no standing governance meeting; proposals, discussion, and decisions happen in public on GitHub; ad hoc meeting, Slack, or mailing-list outcomes are only effective once recorded in the relevant GitHub issue or pull request. Tracked in [`nats-io/nats-general#111`](https://github.com/nats-io/nats-general/issues/111).

- [x] **Documentation of how to contribute, with increasing detail as the project matures.**

Met. NATS has project-wide contribution guidance and repository-specific contribution guidance for `nats-server`, with detailed expectations for issues, PRs, sign-off, tests, review, and release branches.

- [x] **Demonstrate contributor activity and recruitment.**

Met. DD verified active community activity through public repository activity, public communication channels, `nats-server` contributor data, and LFX Insights. DD re-checks recorded 214 GitHub contributors for `nats-server`; LFX Insights reported 24 maintainers, 57 reviewers, organization-dependency of 2 organizations for 67% of contributions, and contributor-dependency of 6 contributors for 55% of contributions over the past 365 days.

## Engineering Principles

### Suggested

N/A

### Required

- [x] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently. _This requirement may also be satisfied by completing a General Technical Review._**
  - The [project-maintained General Technical Review](https://github.com/nats-io/nats-general/blob/main/GENERAL-TECHNICAL-REVIEW.md) was completed/updated on 2026-01-19 and reviewed during DD.

Met. The GTR describes NATS as a lightweight, high-performance messaging system for cloud-native applications, IoT, edge, microservices, streaming/event-driven architectures, control planes, and multi-cloud / hybrid deployments. The core differentiation is a small, simple, deploy-anywhere messaging system with subject-based addressing, very low latency, optional persistence through JetStream, and support for topologies from single server to clusters and superclusters.

- [x] **Document what the project does, and why it does it - including viable cloud native use cases. _This requirement may also be satisfied by completing a General Technical Review._**
  - The [project-maintained General Technical Review](https://github.com/nats-io/nats-general/blob/main/GENERAL-TECHNICAL-REVIEW.md) was completed/updated on 2026-01-19 and reviewed during DD.

Met. The GTR and docs cover publish/subscribe, request/reply, queue groups, JetStream persistence/streaming, service discovery, IoT/edge messaging, streaming/event-driven systems, control planes, and multi-cloud/hybrid connectivity. NATS integrates with Kubernetes, Prometheus, Terraform, Jaeger, Zipkin, and OpenTelemetry-related tracing workflows.

- [x] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

Met. The project points to the public roadmap at https://nats.io/about/#roadmap and describes the roadmap as maintainer-led but informed by GitHub, Slack, adopters, vendor feedback, commercial users and support providers, support experience, community needs, and Architecture Decision Records in https://github.com/nats-io/nats-architecture-and-design. The Synadia-specific roadmap-input wording identified during DD was remediated in [`nats-site#1112`](https://github.com/nats-io/nats-site/pull/1112).

- [x] **Roadmap change process is documented.**

Met. The roadmap input sources and features/changes are usually created as GitHub issues and concrete proposals are made as architecture and decision records. The site roadmap now lists community, adopter, vendor, commercial-user/support-provider, support, and maintainer input sources without singling out Synadia as a privileged channel.

- [x] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.  *This requirement may also be satisfied by completing a General Technical Review and capturing the output in the project's documentation.***

Met. Architecture is covered in the public docs and https://github.com/nats-io/nats-architecture-and-design. The project describes a lightweight, deploy-anywhere messaging system built around NATS servers, clusters, superclusters, leaf nodes, subject-based APIs, authentication and authorization models, and optional persistence through JetStream streams and consumers. The documented design supports cloud native use cases including service-to-service communication, publish/subscribe messaging, request/reply, queue groups, event-driven systems, control planes, IoT and edge messaging, multi-region and multi-cloud connectivity, Kubernetes-based deployment, monitoring, and operational patterns for high availability and rollout/rollback. This documentation demonstrates that NATS has a mature architectural model suitable for production cloud native environments.

- [x] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:**

  - [x] Release expectations (scheduled or based on feature implementation)
  - [x] Tagging as stable, unstable, and security related releases
  - [x] Information on branch and tag strategies
  - [x] Branch and platform support and length of support
  - [x] Artifacts included in the release.
  - Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.

Met. `nats-server` publishes release guidance at https://github.com/nats-io/nats-server/blob/main/RELEASES.md. As of NATS 2.11, the server follows a six-month release cycle; version tagging follows semantic versioning; release branches support patch releases; the current and previous minor series receive bug fixes and security patches; binaries are released via GitHub Releases; and Docker images are available as the official image.

- [x] **History of regular, quality releases.**

Met. DD verified the documented release policy and public GitHub Releases/tags. Prior to 2.12 the release cadence was scope-based; as of 2.12, the documented minor release cadence is six months. GitHub Releases and tags provide public release history. Recent public `nats-server` releases include `v2.15.1` and `v2.14.8` on 2026-10-09, demonstrating active release maintenance.

## Security

Note: this section is augmented by the security self-assessment reviewed during DD, the public NATS security policy, public advisory materials, and third-party audit references.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

Not currently met. DD verified that NATS has achieved the OpenSSF Best Practices passing badge; silver/gold remains aspirational. Silver/gold is suggested, not required, for graduation.

- [ ] **Complete a joint security assessment with TAG Security and Compliance**

Partially met. NATS provided a security self-assessment for DD and a threat-model document, and public security materials are available through the NATS security policy and advisories site. The self-assessment indicates it was created through the TAG Security Pals process and is pending maintainer changes. Remaining verification items: formal joint-assessment record status and self-assessment status updates. This is a suggested criterion and is not blocking graduation.

### Required

- [x] **Clearly defined and discoverable process to report security issues.**

Met after remediation. Security issues can be reported through GitHub private vulnerability reporting or, as a fallback, by email to `security@nats.io`. NATS maintains a security advisory policy at https://advisories.nats.io/advisory-policy and advisories at https://advisories.nats.io/. [`nats-general#121`](https://github.com/nats-io/nats-general/pull/121) added canonical `nats-general/SECURITY.md`; [`nats-io/.project#6`](https://github.com/nats-io/.project/pull/6) updated `.project/project.yaml` to that live policy path; and [`nats-server#8744`](https://github.com/nats-io/nats-server/pull/8744) added `nats-server/SECURITY.md` as a discoverable shim to the canonical policy and `nats-server` advisory reporting flow. Tracked in [`nats-io/nats-general#113`](https://github.com/nats-io/nats-general/issues/113) and [`nats-io/nats-general#116`](https://github.com/nats-io/nats-general/issues/116).

- [x] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

Met for public DD purposes. Public contribution and governance docs require maintainer review before merge, the self-assessment describes review expectations and maintainer approval requirements, and `nats-server` publishes CODEOWNERS at https://github.com/nats-io/nats-server/blob/main/.github/CODEOWNERS. Org-level 2FA and branch-protection details are not fully public; confirming or enabling stronger 2FA controls is captured as a non-blocking recommendation rather than a graduation blocker because the public process and security self-assessment support the criterion.

- [x] **Document assignment of security response roles and how reports are handled.**

Met after remediation. The self-assessment documents the responsible disclosure process, vulnerability response process, advisory publication process, security reporting channels, and incident-response considerations. `nats-general/SECURITY.md` now documents security response team membership/access, appointment/removal model, reporting paths, response timeline, disclosure policy, and advisory publication path. Remediated by [`nats-general#121`](https://github.com/nats-io/nats-general/pull/121) and tracked in [`nats-io/nats-general#112`](https://github.com/nats-io/nats-general/issues/112).

- [x] **Document Security Self-Assessment.** The completed self-assessment must be available in `cncf/toc` under `projects/<project>/security-assessment/self-assessment.md`.

Met. The security self-assessment was reviewed during DD. It covers metadata, security links, actors, actions, goals, non-goals, security functions/features, secure development practices, security issue resolution, known issues, CII/OpenSSF status, and related projects.

- [x] **Third Party Security Review.**

  - [x] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

Met. A Trail of Bits audit was completed in May 2025 and 100% of findings were resolved or accepted: https://ostif.org/nats-audit-complete/. There is also a prior Cure53 audit at https://github.com/nats-io/nats-general/blob/main/reports/Cure53_NATS_Audit.pdf.

- [x] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

Met. NATS has achieved the OpenSSF Best Practices passing badge: https://www.bestpractices.dev/projects/1895.

## Ecosystem

### Suggested

N/A

### Required

- [x] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

Met. DD verified that NATS maintains a public adopter list at https://github.com/nats-io/nats-general/blob/main/ADOPTERS.md, public adopter logos, and CNCF case studies.

- [x] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

Met. DD verified the public adopter list includes more than 100 organizations across multiple industries, including technology, financial services, automotive, telecommunications, software, and the broader cloud native ecosystem. TOC verification also confirmed production use through three independent adopter interviews.

The project provided the TOC with adopter contacts for verification of use of the project at the level expected, i.e. production use for graduation.

- [x] **TOC verification of adopters.**

**Complete:** Three independent production-adopter interviews were completed, satisfying the graduation minimum. Interview notes and recordings are kept privately for internal TOC use and verification. [#### Adoption](#adoption) lists interview status by industry vertical only (anonymous).

**Findings (three interviews):**

Across the interviews, adopters described NATS as production-critical infrastructure used for cloud, edge, multi-region, multi-cloud, and on-premises deployments. Reported use cases included pub/sub messaging, control-plane and event signaling, work queues, telemetry, key-value and persistence-oriented workflows, MQTT-related patterns, and JetStream-backed operations. The deployments described were not experimental or limited pilots; adopters reported relying on NATS for large-scale distributed services, resilient messaging infrastructure, and long-running service/network operations.

**Strengths adopters cited:**

1. **Simplicity and operational leverage** — Adopters valued getting messaging, persistence, key-value patterns, and related communication primitives from one lightweight system. Several contrasted NATS favorably with heavier messaging stacks that require more dedicated infrastructure, staffing, or chargeback models.
2. **Reliability and resilience** — Adopters reported running NATS as highly available or zero-downtime infrastructure across regions, cloud environments, Kubernetes clusters, edge platforms, and on-premises deployments. They described NATS as suitable for production-critical service communication and operational workflows.
3. **Cloud-native deployment fit** — Adopters reported successful use with Kubernetes, Helm, GitOps, canary or rolling upgrades, and automated deployment pipelines. NATS fit existing platform-engineering workflows without requiring a heavyweight operational model.
4. **Open-source confidence and community responsiveness** — Adopters valued that core capabilities are available in the open project rather than reserved for a proprietary enterprise edition. They also cited positive experiences with the community and maintainers, including rapid upstream issue response.

**Improvements adopters would like to see:**

1. **High-availability documentation** — Clearer documentation on what is and is not included in the HA set for specific configurations, including operational boundaries and failure-mode expectations.
2. **Credential and secret-management tooling** — Easier credential rotation and stronger tooling or examples for integrations with common secret-management systems.
3. **Secure-deployment guidance** — Clearer guidance for decentralized credentials, root certificates, and compliance-oriented operating models.
4. **Ecosystem and deployment options** — Additional guidance or support for database/global-consistency patterns, industrial protocol use cases, and implementation/runtime options attractive for embedded or supply-chain-sensitive environments.

- [x] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

Met. NATS project resources document integrations and compatibility across both CNCF and non-CNCF projects:

| Integration | CNCF? | Project resource |
| :--- | :--- | :--- |
| **Kubernetes** deployment | yes | NATS deployment docs: https://docs.nats.io/learn/deployment/ and Kubernetes assets in [`nats-io/k8s`](https://github.com/nats-io/k8s) |
| **Helm** charts | yes | NATS Helm charts in [`nats-io/k8s`](https://github.com/nats-io/k8s/tree/main/helm/charts/nats) |
| **NATS Kubernetes controller/operator work** | yes | [`nats-io/nack`](https://github.com/nats-io/nack) |
| **Prometheus** monitoring | yes | NATS monitoring docs: https://docs.nats.io/learn/monitoring/, [`prometheus-nats-exporter`](https://github.com/nats-io/prometheus-nats-exporter), and [`nats-surveyor`](https://github.com/nats-io/nats-surveyor) |
| **OpenTelemetry / Jaeger / Zipkin-style tracing workflows** | yes / mixed | NATS distributed tracing docs: https://docs.nats.io/using-nats/developer/receiving/tracing |
| **Terraform** | no | [`terraform-provider-jetstream`](https://github.com/nats-io/terraform-provider-jetstream) |
| **MQTT** | no | NATS MQTT compatibility docs: https://docs.nats.io/learn/mqtt/ |

These project-maintained docs and repositories demonstrate that NATS is compatible with common cloud-native deployment, monitoring, tracing, infrastructure-as-code, and messaging ecosystems.

#### Adoption

Interview findings are summarized under [TOC verification of adopters](#toc-verification-of-adopters) in [## Ecosystem](#ecosystem). All adopter interview notes and recordings are privately kept for TOC use and verification.

##### Adopter 1 - Cloud and edge services

**Status:** Interview complete (September 2026). Anonymous.

##### Adopter 2 - Enterprise messaging platform

**Status:** Interview complete (September 2026). Anonymous.

##### Adopter 3 - Network services

**Status:** Interview complete (September 2026). Anonymous.
