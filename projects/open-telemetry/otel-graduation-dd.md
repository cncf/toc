# OpenTelemetry Graduation Due Diligence

- Link to [[Graduation] OpenTelemetry Graduation Application](https://github.com/cncf/toc/issues/1739)

## Graduation Evaluation Summary for OpenTelemetry

### Criteria Evaluation

[Emily Fox](https://github.com/TheFoxAtWork) and [Davanum Srinivas](https://github.com/dims) conducted the due diligence of OpenTelemetry who applied for Graduation. The project has completed the criteria that show its maturity at the applied level. 

The following criteria implementations are noteworthy to call out:
- [OpenTelemetry](https://opentelemetry.io/) is an open-source observability framework that standardizes the way applications generate, collect, and export telemetry data (traces, metrics, and logs). It emerged from the merger of two earlier projects namely [OpenTracing](https://opentracing.io/) and [OpenCensus](https://opencensus.io/).
- OpenTelemetry is designed for modern distributed systems and microservices architectures, which are central to cloud-native paradigms ([What is OpenTelemetry?](https://opentelemetry.io/docs/what-is-opentelemetry)).
- The project's design principles emphasize vendor-neutrality, data ownership by users, and a single set of APIs and conventions, all of which promote broader adoption and integration of cloud native technologies. ([Specification Principles](https://opentelemetry.io/docs/specs/otel/specification-principles/))
- [Key Features](https://opentelemetry.io/docs/concepts/) include [Multi-language Support](https://opentelemetry.io/docs/languages/), Context Propagation, Flexible Backend Support, Auto-Instrumentation etc.
- Healthy [Specification process](https://opentelemetry.io/docs/specs/) that drives innovation and implementations. [Decoupled components](https://opentelemetry.io/docs/concepts/components/) that makes it easy for end users to consume what they really want to with upfront [maturity levels](https://opentelemetry.io/status/) for each component.
- OpenTelemetry has a complete [contributing guide](https://opentelemetry.io/docs/contributing/) provides clear instructions for anyone interested in improving or expanding the OpenTelemetry code, docs and website.
- The OpenTelemetry community repository makes it very easy to get started, find meetings, slack channels, sponsors, liaisons etc right in the main [README](https://github.com/open-telemetry/community?tab=readme-ov-file#specification-sigs). Specific shoutout to the [Membership, Roles, and Responsibilities](https://github.com/open-telemetry/community/blob/main/guides/contributor/membership.md) page that has the contributor ladder for folks to aspire to larger roles in the project.
- There are two active bodies, the [Governance Committee (GC)](https://github.com/open-telemetry//blob/main/governance-charter.md) and the [Technical Committee (TC)](https://github.com/open-telemetry/community/blob/main/tech-committee-charter.md) that provide governance and ensure that the community is healthy.
- The most [recent Audit](https://opentelemetry.io/blog/2024/security-audit-results/) in July 2024 in collaboration with OSTIF and conducted by 7ASecurity, resulted in one CVE finding and five hardening recommendations, for which OpenTelemetry resolved the CVE finding and is in progress on implementing each hardening recommendation from the audit.

The following actions were provided to the project that were considered blocking but have since been resolved to a degree that are deemed satisfactory for Graduation:
- Clarification regarding stable components of distributions - Initial adopter interviews highlighted confusion regarding stability expectations of OTEL distributions. The project has taken this initial feedback and provided a [plan for "Stable by Default"](https://github.com/open-telemetry/opentelemetry-specification/pull/4813). This initiative has evolved from a simple policy change into an open sdiscussion on balancing the project's stability reputation with the practical reality of experiemental community contributions. Refer to the non-blocking section for resolution recommendations.
- Missing consistent adherence to the project's governance - Though the central body responsible for Security seems to be SIG-Security, the TC seems to be still carrying the burden both in terms of documentation and process as the definition for a "Security Response Committee (SRC)" is still in draft status and probably understaffed ([Security Response Committee - Current Status? Future Plan?](https://github.com/open-telemetry/community/issues/2835)).
- [Missing information regarding Roadmap changes](https://github.com/open-telemetry/community/issues/2894) - the project leverages the Discussions feature of GitHub to document Roadmap planning activities that occur in SIG meetings. The Project's Governance Committee charter details the GC as being responsible for ensuring the Roadmap is updated, but does not convey the process by which that occurs. While the project does maintain a project proposal process, this is not integrated into Roadmap development and maintenance efforts. It was recommended the Governance Committee document the process by which they and the sub-components make changes or consider efforts to be prioritized and included within the Roadmap; connecting project proposal to roadmap inclusion. This was completed with the merge of [PR 2911](https://github.com/open-telemetry/community/pull/2911).
- [Completion of best practice badging for core components of the project](https://github.com/open-telemetry/community/issues/2919) - The project initially completed best practices passing badging for the collector, Java, go, Android, .NET. After engaging the project to bring additional repositories in scope, python, the specification, and protocol were completed.

The following recommendations were provided to the project that are non-blocking in the TOC's assessment but should be completed by the project to ensure continued viability of the project:
- As each repository has its own release cadence, a recommendation is to have a central jump off point and document some things that are common across these would be useful ([issue #2814](https://github.com/open-telemetry/community/issues/2814))
- Complete Stable-by-Default. What follows are recommendations by the TOC to progress and close ongoing discussion. The project should move to merge the conceptual framework of PR #4813 while deferring the specific implementation details for individual SDKs to their respective SIGs, provided they adhere to the Explicit Opt-In and Stability Metadata principles recommended below.
  - Implement feature capability metadata for standardization in the specification rather than using a binary Stable/Experimental flag for the entire distribution. Desired Outcome: permit automated tooling like the OTEL collector builder to generate stability reports during build to alert maintainers if they are inadvertently shipping experimental code in a stable release.
  - Formalize Core versus contrib tiering and expectations which contribut to user confusion on lifecycle support for components. Desired Outcome: Mirror successful models used by others where different repositories have different stability "SLAs".
  - Mandate a programmatic barrier for experiental features to prevent innovation stifling. Desired Outcome: ensure no user accidently relies on an unstable API, protecting the stability gurantees while allowing community contributed content to co-exist. This may be through a requirement for any experiemental feature in a stable distribution to be disabled by default and require an explicit config flag to activate (e.g., enable_unstable_feature:true or a specific feature gate).

### Adoption Evaluation

The adopter interviews reflect a project in use for the level which the project applied. Uniquely this project demonstrated a strategic, high-maturity shift over the course of the adopter interviews conducted. Initially the project did not adhere to or enforce conformance for its community contributed content – where a substantial portion of adoption occurs from –  resulting in buggy releases. After engagement by the TOC with the project to identify the root cause of these issues, the project took meaningful steps to address the feedback provided. The final adopter interview notes improvements in documentation and structural guidance which serve as an accelerator to project adoption.

### Final Assessment

The TOC finds that OpenTelemetry has successfully transitioned from a high-potential framework to a mature, stable, and foundational pillar of the cloud native ecosystem. While initial adopter feedback highlighted concerns regarding the consistency and stability of community-contributed content, the project has demonstrated a marked strategic shift by implementing rigorous structural improvements and governance oversight to address these gaps. The project maintains an exemplary governance status, with a robust contributor ladder and a clear separation of concerns between its Governance and Technical Committees. Its commitment to vendor neutrality is enforced through strict representation limits and mission-driven guidelines.

The TOC concludes that OpenTelemetry has met all criteria for Graduation. The project’s proactive response to feedback regarding component stability—combined with its ubiquitous industry adoption and healthy, multi-organization contributor base—confirms its readiness for Graduation.

### Criteria

## Application Process Principles

### Suggested

N/A

### Required

- [X] **Give a presentation and engage with the domain specific TAG(s) to increase awareness**

The project [presented to TAG Observability on Dec 18th 2024](https://www.youtube.com/watch?v=1sT5wfHUx5M) to discuss the future of OpenTelemetry and observability in general. The presentation also covered current and future challenges in the observability space.

- [X]  **TAG provides insight/recommendation of the project in the context of the landscape**

The Technical Advisory Group for Operational Resilience provided the TOC with the following positive recommendation regarding OpenTelementry:

> After reviewing the December 2024 meeting, OpenTelemetry DevStats, and discussing among ourselves within the TAG, we are unanimous in stating that OpenTelemetry has reached a high level of maturity and is widely adopted across the industry. It plays a critical and foundational role in the cloud-native observability ecosystem, offering standardized telemetry for traces, metrics, and logs. With active community contributions and growing production usage, OTEL is becoming the de facto standard for telemetry collection and propagation. The project is also leading significant efforts in areas such as profiling, centralized management, and control of observability agents via OpAMP, as well as multi-signal correlation, making it an essential component of resilient system design.

> We recommend OTEL favorably and recognize it as a strategic project within the cloud native landscape. 

- [X]  **All project metadata and resources are [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/).**

The OpenTelemetry project has exemplary documentation regarding vendor neutrality and subsequently incorporated governance processes to enforce this neutrality. In addition to setting [maximum representation on governance and technical committees](https://github.com/open-telemetry/community/blob/main/governance-charter.md#maximal-representation), the project also has incorporate neutrality into their [mission and values](https://github.com/open-telemetry/community/blob/main/mission-vision-values.md) as well as provided [marketing guidance](https://opentelemetry.io/community/marketing-guidelines/) to ensure vendor-neutrality. A random sampling of blog posts and other project resources and data reinforce the project's committment to vendor neutrality as none of the content sample demonstrated any bias or favoritism of one vendor over another.

- [X] **Review and acknowledgement of expectations for [sandbox](sandbox.cncf.io) projects and requirements for moving forward through the CNCF Maturity levels.**		
   - [X] Met during Project's [re-application with the new Issue template on 30-May-2025](https://github.com/cncf/toc/issues/1739).

Completion of this due diligence document, resolution of concerns raised, and presented for public comment satisfies the Due Diligence Review criteria.

- [X] **Additional documentation as appropriate for project type, e.g.: installation documentation, end user documentation, reference implementation and/or code samples.**

As a graduating project, OpenTelemetry has sufficient documentation for its project type, with robust user documentation that covers the user persona, demos of the project, details regarding its specifications, spec-supported project and vendor listing, migration information, and architecture diagrams among many other material topics. 

## Governance and Maintainers

Note: this section has been augmented by the completion of a Governance Review from Project Reviews TOC Sub-project.

OTEL's [Governance review was performed in August 2025](https:github.com/cncf/toc/projects/open-telemetry/governance-review/2025-08-14). The TOC has found no conflict with the governance review group's findings.

> **Status:** Exemplary

> Open Telemetry is a large project and has appropriate governance for a project of it's size. They delegates ownership to Special Interest Groups (SIGs), with each SIG responsible for subprojects in their respective domain.

> They have exceptional contributor documentation, end-user engagement, and a
steady pipeline of new contributors.

### Suggested

- [X]  **Governance has continuously been iterated upon by the project as a result of their experience applying it, with the governance history demonstrating evolution of maturity alongside the project's maturity evolution.**

### Required

- [X] **Clear and discoverable project governance documentation.**

- [X] **Governance is up to date with actual project activities, including any meetings, elections, leadership, or approval processes.**

- [X] **Governance clearly documents [vendor-neutral](https://contribute.cncf.io/maintainers/community/vendor-neutrality/) of project direction.**

- [X] **Document how the project makes decisions on leadership, contribution acceptance, requests to the CNCF, and changes to governance or project goals.**

- [X] **Document how role, function-based members, or sub-teams are assigned, onboarded, and removed for specific teams (example: Security Response Committee).**

- [X] **Document complete list of current maintainers, including names, contact information, domain of responsibility, and affiliation.**

- [X] **A number of active maintainers which is appropriate to the size and scope of the project.**

- [X] **Document a complete maintainer lifecycle process (including roles, onboarding, offboarding, and emeritus status).**

- [X] **Demonstrate usage of the maintainer lifecycle with outcomes, either through the addition or replacement of maintainers as project events have required.**

- [X] **Project maintainers from at least 2 organizations that demonstrates survivability.**

- [X] **Code and Doc ownership in Github and elsewhere matches documented governance roles.**

- [X] **Document agreement that project will adopt CNCF Code of Conduct.**

- [X] **CNCF Code of Conduct is cross-linked from other governance documents.**

- [X] **All subprojects, if any, are listed.**


- [X] **If the project has subprojects: subproject leadership, contribution, maturity status documented, including add/remove process.**


## Contributors and Community

Note: this section has been augmented by the completion of a Governance Review from Project Reviews TOC Sub-project.

OTEL's [Governance review was performed in August 2025](https:github.com/cncf/toc/projects/open-telemetry/governance-review/2025-08-14). The TOC has found no conflict with the governance review group's findings.

### Suggested

- [X] **Contributor ladder with multiple roles for contributors.**

### Required

- [X] **Clearly defined and discoverable process to submit issues or changes.**

- [X] **Project must have, and document, at least one public communications channel for users and/or contributors.**

- [X] **List and document all project communication channels, including subprojects (mail list/slack/etc.).  List any non-public communications channels and what their special purpose is.**

- [X] **Up-to-date public meeting schedulers and/or integration with CNCF calendar.**

- [X] **Documentation of how to contribute, with increasing detail as the project matures.**

- [X] **Demonstrate contributor activity and recruitment.**

## Engineering Principles

### Suggested

N/A

### Required

- [X] **Document project goals and objectives that illustrate the project’s differentiation in the Cloud Native landscape as well as outlines how this project fulfills an outstanding need and/or solves a problem differently.**

The project provides its [mission, vision, and values publicly](https://opentelemetry.io/community/mission/). The Vision covers 5 guiding points which complement and underscore characteristics of cloud native (e.g. loosely coupled), while their Values outline how the project fulfills the telemetry needs of the ecosystem; priorizing compatibility, stability, resilience, and performance to ensure effective observability through ubiquitous high-quality, portable telemetry.

- [X] **Document what the project does, and why it does it - including viable cloud native use cases.**

The Project's documentation features a page "[What is OpenTelemetry](https://opentelemetry.io/docs/what-is-opentelemetry/)" which provides an overview of what the project does and doesn't do (notably it does not provide the storage or visualization for telemtry data, offloading this to other cloud native projects or commerical tools to meet adopters where they are).

While not explicitly called out in the project documentation, the information available provides adopters with an understanding of the use cases solves for. It is recommended the project update its "What is OpenTelemtry?" to include the following desired use:
> The project is designed specifically to address the observability needs inherent in cloud native using a 'hands-off' approach through instrumentation, where comparatively traditional methods of monitoring, observation, and eventual troubleshooting are too resource intensive, manual, or inconsistent to utilize in distributed cloud systems.

- [X] **Document and maintain a public roadmap or other forward looking planning document or tracking mechanism.**

The Project's Community repository contains a generalized [Roadmap](https://github.com/open-telemetry/community/blob/main/roadmap.md) updated [November of 2024](https://github.com/open-telemetry/community/commit/58f1679f9997d899eed89ccfe8e4720faf808386) as of this writing.

Several of the [Project's sub-components](https://github.com/search?q=org%3Aopen-telemetry%20roadmap&type=code) also detail their own Roadmaps that convey what work is planned.

- [X] **Roadmap change process is documented.**

The Project recently modified the [project management](https://github.com/open-telemetry/community/blob/main/project-management.md) file to conveys how the project manages its boards, roadmap, and proposals. A dedicated file exists for [Roadmap management](https://github.com/open-telemetry/community/blob/main/roadmap-management.md). 

- [X] **Document overview of project architecture and software design that demonstrates viable cloud native use cases, as part of the project's documentation.**

The Project maintains an [architecture diagram](https://opentelemetry.io/docs/) on the Documentation landing page, which describes the various sub-components in relation to one-another. An overview of the 1.47 spec release does [convey the terminology and concepts of the OpenTelemetry client at a high level](https://opentelemetry.io/docs/specs/otel/overview/#opentelemetry-client-architecture). 

It is may be beneficial for potential users if the project also refer to this diagram from the spec version Overview to couple the concepts in the spec overview with the architecture of the project itself, however this is not required.

- [X] **Document the project's release process and guidelines publicly in a RELEASES.md or equivalent file that defines:** 

  - [X] Release expectations (scheduled or based on feature implementation)
  - [X] Tagging as stable, unstable, and security related releases
  - [X] Information on branch and tag strategies
  - [X] Branch and platform support and length of support
  - [X] Artifacts included in the release.
  - Additional information on topics such as LTS and edge releases are optional. Release expectations are a social contract between the project and its end users and hence changes to these should be well thought out, discussed, socialized and as necessary agreed upon by project leadership before getting rolled out.

The Project maintains a [Release process](https://github.com/open-telemetry/community/blob/main/RELEASE.md) which lays out the expectations, practices, and versioning for releases. It also includes specification and component specific information regarding releases.

The project leverages tagging to indicate experimental, stable, deprecated for releases for which [detailed explanations convey expectations](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/versioning-and-stability.md). While the project does not leverage "unstable" or "security" tagging, the TOC does not find this to be an issue.

The project primarily leverages the main branch for development. The OpenTelemetry Collector project leverages release branches to manage their release process. Git Tags are used for official release processes and artifact publishing.

The projects official documentation specifies their component lifecycle status, with Stable covering backward compatibility and long-term support.

OTEL has several release arifact types, such as container images and language-specific SDKs published to their respective repositories. The specification itself is the primary artifact and has its own dedicated GitHub repository for protcol bugger definitions.

- [X] **History of regular, quality releases.**

The overall project has a very structured approach to development with a long-standing commitment to stability. Specifically, Collector and the language SDKS have a consisent release cadence, some occuring every few weeks to ensure relevancy and partiy with the broader ecosystem. The project's use and distinction of experimental and stable tagging for components indicates attention to the quality of components across their lifecycle and releases. The [release notes](https://github.com/open-telemetry/opentelemetry-collector/releases) further highlight bug fixes within each release, with consideration for fix backporting on older, minor releases.

Several Adopters reported a majority of issues regarding the project's releases. This was raised with the project who has taken steps to improve the clarity regarding official OTEL distribution stability.

## Security

Note: this section may be augmented by a joint-assessment performed by TAG Security.

### Suggested

- [ ] **Achieving OpenSSF Best Practices silver or gold badge.**

Due to the structure of the OTEL project, several components of OTEL have either completed passing level or are in progress.

### Required

- [X] **Clearly defined and discoverable process to report security issues.**

Users who identify potential security issues have discoverable areas (org repo level readme and security.md) that direct them to the overarching [Security Policy](https://github.com/open-telemetry/community/security/policy) of the project and lists two methods for reporting.

- [X] **Enforcing Access Control Rules to secure the code base against attacks (Example: two factor authentication enforcement, and/or use of ACL tools.)**

The project uses branch protection rules that prevent direct pushes to main. Additionally, [otelbot](https://github.com/open-telemetry/community/blob/main/assets.md#otelbot) that is configured with specific branch protection rules that supplement the limitations of GitHub Actions. The project's entire [asset.md](https://github.com/open-telemetry/community/blob/main/assets.md) file is remarkably well documented and covers a breadth of information regarding the project's tools, services, and accounts.

- [X] **Document assignment of security response roles and how reports are handled.**

The Security SIG for the project is responsbile for not only security initiatives, but also for responding to reported vulnerabilities. The SIG has its own [charter](https://github.com/open-telemetry/sig-security/blob/main/sig-security-charter.md) that outlines its scope and responsibilities. The project has a well documented [vulnerability response](https://github.com/open-telemetry/sig-security/blob/main/security-response.md) process.

- [X] **Document Security Self-Assessment.**

The project's [self-assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/open-telemetry/self-assessment.md) is document in TAG Security's repository. 

- [X] **Third Party Security Review.**

  - [X] Moderate and low findings from the Third Party Security Review are planned/tracked for resolution as well as overall thematic findings, such as: improving project contribution guide providing a PR review guide to look for memory leaks and other vulnerabilities the project may be susceptible to by design or language choice ensuring adequate test coverage on all PRs.

The project completed a third-party security audit from May to June of 2024. All high findings were resolved. [Moderate and Low findings are tracked in the following issues](https://github.com/search?q=org%3Aopen-telemetry+%2FOTE-01%2F&type=issues&p=1):
- [OTE-01-001: Usage of Multiple Vulnerable Dependencies (Low)](https://github.com/open-telemetry/opentelemetry-collector/issues/10766) - OPEN
- [OTE-01-002: Possible DYLIB Injection on MacOS Client (Medium)](https://github.com/open-telemetry/opentelemetry-collector-releases/issues/586) - OPEN/stale
- [OTE-01-003: Enhanced Security Against MitM via TLS MinVersion (Info)](https://github.com/open-telemetry/opentelemetry-collector/issues/10470) - OPEN
- [OTE-01-004: Possible DoS Attacks on HTTP Services (Medium)](https://github.com/open-telemetry/opentelemetry-collector/issues/10469) - OPEN
- [OTE-01-005 WP1: Linux Binaryu Hardening Recommendations](https://github.com/open-telemetry/opentelemetry-collector-releases/issues/618) Reopened, as PR was [reverted](https://github.com/open-telemetry/opentelemetry-collector-releases/commit/62dfc10402322ae4e2cdbdd92a0c0cc797f1b1f4)

- [X] **Achieve the Open Source Security Foundation (OpenSSF) Best Practices passing badge.**

The project initially had a few repositories with passing badge. After engaging with the project and identifying additional repositories in scope, the project [completed best practice badging in September 2025](https://github.com/open-telemetry/community/issues/2919).

## Ecosystem

### Suggested

N/A

### Required

- [X] **Publicly documented list of adopters, which may indicate their adoption level (dev/trialing, prod, etc.)**

The project maintains their list of adopter's on the project's website: https://opentelemetry.io/ecosystem/adopters/ 

- [X] **Used in appropriate capacity by at least 3 independent + indirect/direct adopters, (these are not required to be in the publicly documented list of adopters)**

The project provided the TOC with a list of adopters for verification of use of the project at the level expected. The TOC conducted 5 Adopter Interviews that demonstrate the projects' use in a capacity suitable for Graduated projects.

- [X] **TOC verification of adopters.**

The TOC received the adopters list for verification of the project at the graduation level.

Refer to the Adoption portion of this document.

- [X] **Clearly documented integrations and/or compatibility with other CNCF projects as well as non-CNCF projects.**

The project maintains a [comprehensive list of integrations](https://opentelemetry.io/ecosystem/integrations/) with other projects and products that provide first-class support of OTEL.

#### Adoption

##### Adopter 1 - Parseable/Observability Vendor

###### Debabrata Panigrahi, August 2025

Parseable has been using OpenTelemetry for two years, approximately since it was open sourced, it is currently in use in production. Paresable’s product is an observability tool that includes OTEL and additional tooling as a binary for customers. In their own words, they provide an enterprise-scale observability platform, with Parquet's compression, S3's durability, and Apache Arrow's performance to turn Petabytes of telemetry into instant insights.

Parseable targets an n-1 version of the project for inclusion in their binary. As new versions are released that provide security updates or breaking changes, they pull in those updates but their overall cadence is more relaxed as of late to continue to provide enterprise-grade performance and stability of their product.

When first looking at the space of observability, fluentbit had far more accessible documentation. OpenTelemetry’s documentation was initially subpar, making instrumentation and integration challenging. They used ChatGPT to supplement the documentation but encountered issues with hallucinations due to lack of available reference material (most of observability use cases and tutorials focused on fluentbit nearly exclusively). They noted a lag between when project documentation is updated and when that information becomes available through GenAI tools.

The information in the OTEL repository was somewhat valuable, but the format of OTEL outputs was a bit unclear, especially when compared to Fluentbit. Parsable identified a need for more tutorials and best practices to clarify which instrumentation is most effective with the given output formats. Parseable primarily use metrics and traces, with CPU utilization being a key metric for their customers.

Parseable is not very active in the community but have engaged via Slack and have subscribed to the mailing list and commented on GitHub issues. They have found that most of their issues and support questions are already covered by the community. They plan to contribute back some work related to the OTEL demo app with FluentBit once they’ve made the effort more publicly testable. Parseable finds the community to be a substantial value of the project itself.

Parseable is particularly interested in the OTEL Arrow protocol, which they are following closely and intend to use once it’s GA. They believe that while Fluentbit is more mature, the OTEL community is more focused on next-generation observability. They also find that OTEL is better for newer columnar formats like Parquet.

Parseable’s implementation of OTEL within their product has been "immensely useful." They have found that OTEL is faster than Fluentbit for their use case, though there is always a trade-off between memory and performance. They have also seen benefits in data compression which saves bandwidth. While OTEL doesn't directly help with storage costs, their use of Parquet with their own compression allows them to save costs.

The project could be improved by focusing on reducing costs for users. They also mentioned concerns about data loss on the agent side as a general concern. Overall, the project could greatly benefit from more tutorials across different use cases, more detailed instruction on instrumentation, and a focused effort on cost and performance efficiencies.

They do not have any major plans beyond their interest in OTEL Arrow. The project currently works well for their use case.

Many of their customers are legacy companies that are new to cloud native observability. These customers have found that Fluentbit is often easier to get started with due to the larger amount of available documentation and articles. However, OTEL is a natural next step beyond Fluentbit as it was designed to provide comprehensive observability out of the gate which provides a more holistic view of system behavior. 

##### Adopter 2 - Delivery Hero/quick commerce

###### Elena Kovalenko, August 2025

Delivery Hero is a global online food and quick commerce delivery company based in Berlin, Germany. It operates a network of platforms that connect customers with a wide range of restaurants and local shops across more than 70 countries and global yearly revenue of $50 billion. The company facilitates the delivery of various goods, including prepared meals, groceries, and household essentials, often through its network of riders.

The company started using OpenTelemetry collectors in non-production environments in February 2024, followed by full adoption in production at scale in August 2025. Their primary objective was to strategically eliminate vendor lock-in within their observability stack, breaking free from vendor-specific agents and SDKs, as well as to create an architecture where future transitions between vendors would require minimal effort and cost, giving them greater flexibility and negotiating power. While the adoption has provided significant strategic value, the journey has highlighted challenges with the stability of contributed components and feature gaps in client-side SDKs.

Delivery Hero uses OpenTelemetry collectors as building blocks in their complex telemetry ingestion pipeline. Their federated collectors alone are running on 155 Kubernetes clusters. In July 2025, the company ingested over 1 PB of logs and 1 PB of traces, along with 200 million metric series, to their observability vendor backend using OpenTelemetry collectors. It's important to note that they sample in only 4% of traces, which means an estimated 25 petabytes of traces actually passed through their collectors that month.

While the core components (including OTLP Arrow receiver and exporter) have been effective, Delivery Hero faced stability issues with components from the contrib repository, a community contributed and maintained collection of components that do not provide the same stability and quality assurances of the components within the main project. Version updates have occasionally introduced regressions. In one instance, an update to the OTTL (OpenTelemetry Transformation Language) library caused unnecessary memory allocation, leading to increased costs. Another update enabled debug-level logging by default, which significantly increased log storage expenses. While bugs were fixed quickly by the community, these incidents made Delivery Hero more cautious about their update cadence.

The core documentation for the data model, tail sampling, and semantic conventions was found to be very helpful for getting started. However, Delivery Hero noted inconsistencies between the project website and the GitHub repository, finding GitHub to be the more reliable source of truth. They identified a need for more documentation covering advanced use cases, which they plan to contribute back to the community.

Delivery Hero has actively engaged with the community by submitting pull requests, reviewing issues, and participating in local meetups. While the community was welcoming, response times on critical issues for Delivery Hero were sometimes slow but understandable. This led Delivery Hero to fork components temporarily to implement fixes at their own velocity while waiting for changes to be merged upstream.

The primary value of adopting OpenTelemetry has not been immediate cost reduction but rather long-term strategic advantage. The biggest benefit is the significant reduction in the estimated cost and effort of any future vendor migrations. This freedom gives Delivery Hero immense leverage in contract negotiations. While instrumentation is currently fragmented—with some teams still using vendor-specific or Prometheus SDKs—the long-term goal is to standardize all server-side applications on OTEL.

Delivery Hero is committed to deepening its investment in OpenTelemetry. They are looking to target 100% OTEL instrumentation for all server-side applications and using OTEL as a common observability language to standardize practices across its federated company structure. As Delivery Hero continues to expand and refine their overall architecture, they plan to contribute documentation on advanced use cases back to the community.

Delivery Hero  identified a key area holding the project back from reaching its full potential. The client-side SDKs are perceived as lacking crucial features available in mature, vendor-specific SDKs. This feature gap was significant enough that it prevented Delivery Hero from adopting OTEL for client-side tracing, forcing them to select a separate vendor solution for that domain.

##### Adopter 3 - Anonymous / Software Technology

This adopter has not approved the release or final content of the interview. It is available to TOC members upon request.
August 2025

##### Adopter 4 - Intuit/FinTech
###### Venkatesh Rangarajan & Ashwini Dulam, September 2025

Intuit, the global technology platform behind products like QuickBooks and TurboTax, adopted OpenTelemetry about two and a half years ago. Intuit has a slow, deliberate update cadence, only upgrading their OTEL components quarterly or semi-annually, common large enterprises like Intuit who prioritize stability over cutting-edge features. 

Intuit has been using OpenTelemetry in production to solve a critical business challenge: reducing incident recovery time (MTTR). While already proficient with metrics and logs for detecting issues (MTTD), their teams struggled to understand complex service dependencies during an incident. The core goal was to gain a complete view of an entire workflow, identify the precise failing service, and understand the full customer impact, thereby slashing their MTTR. 

This need for actionable, end-to-end visibility prompted their strategic move from their previous vendor, Lightstep, to the OpenTelemetry project. This would empower their incident response teams with actionable insights to pinpoint the source of a problem and dramatically accelerate MTTR. The initial challenge was overcoming a fragmented landscape of tools like Lightstep, where inconsistent, "spotty" adoption in their organization prevented a true end-to-end view which would enable identifying the specific failing service–previously a slow, manual process. By building a comprehensive observability platform centered on OTEL, they have gained an end-to-end view of service dependencies, enabling faster root cause analysis. While the project has delivered immense value for their server-side architecture and allowed for cost-effective scale, the journey has also highlighted significant maturity gaps in front-end instrumentation, which remains a major adoption hurdle.

Intuit's adoption was a strategic, phased process that revealed challenges at scale and highlighted key differences in maturity across the OTEL ecosystem. The team began by creating a drop-in replacement for Lightstep to ensure a smooth transition.  Intuit experienced no major performance issues with the collector at scale, only needing to perform minor memory tweaks. 

They found that implementing tracing for backend services and APIs was a relatively seamless process. As adoption grew, Intuit's massive data volume pushed the storage capabilities of off-the-shelf solutions like Grafana Tempo to their limit. To solve this, the team engineered a sophisticated, custom data lake backend. They now store trace data in the Parquet format on Amazon S3, using their own custom indexing service for discoverability. This innovative approach demonstrates the flexibility of the OTEL specification but also points to a maturity gap in turn-key storage solutions for hyper-scale users.

The most significant ongoing challenge is front-end instrumentation. Intuit described this process as "very hard" and "high touch," requiring extensive manual effort and coordination across many developer teams to propagate trace context. Intuit expressed a desire for more standardization in the agent space to ensure they can reuse the tooling they've built. They briefly experimented with other solutions like Langfuse for their agent infrastructure but ran into issues reconciling separate trace IDs. This experience reinforced the value of a unified observability view and highlights their need for a stable, standardized agent ecosystem where tools align with OTEL formats.

Despite the challenges, OpenTelemetry has delivered tangible, high-impact results for Intuit. The primary goal was achieved. Incident response teams now start with the OTEL-powered call graph to identify the failing asset before diving into metrics and logs. This workflow has "substantially improved" recovery times. Current accuracy with OpenTelemetry, still in the early stages, is at 40% in identifying the root cause of an incident. Intuit’s concrete goal is to increase this to 80% accuracy, with their previous vendor, achieving 100% service coverage was cost-prohibitive. OpenTelemetry unlocked their ability to do this. Reaching this threshold is the critical prerequisite for unlocking the next level of operational maturity: enabling automated remediation actions like automatic traffic switching and service rollbacks. This is crucial as it pushes OpenTelemetry beyond just a diagnostic tool into a foundational component for a future, self-healing infrastructure.

Intuit shared the following potential improvement areas for the project:
This is the biggest gap holding the project back is out-of-the-box automatic instrumentation for front-end API calls to remove the manual adoption burden.
The project's success also relies on broader support, such as cloud vendors like AWS natively propagating trace headers.
More focus on building out performance tooling and defining a clear roadmap for collecting performance data from both web clients and backend services.
In order to understand specific configurations, Intuit’s team often has to resort to reading the project's source code. This highlights a gap between functional documentation and the comprehensive, expert-level examples needed for complex, at-scale deployments typical of large adopters.

Intuit's experience demonstrates OpenTelemetry’s stable and strategically valuable relevance for server-side observability at scale. Its core components provide the foundation needed for a sophisticated organization to build a powerful, cost-effective incident response platform. 

The project's primary weakness lies in its immature front-end instrumentation capabilities, which act as a significant barrier to achieving effortless, universal, end-to-end tracing. For prospective adopters, this means that while backend observability can be achieved with high confidence, a successful full-stack implementation will require a substantial and dedicated engineering investment for the client side.

##### Adopter 5 - Large Media Company
###### Person A & Person B, April 2026

The company adopted OpenTelemetry approximately five years ago, making them one of the longer-tenured adopters among those interviewed. The company uses OpenTelemetry across four standardized languages — JVM (currently on hold), TypeScript, Python, and Golang — for logs, metrics, and traces. Their approach centers on building opinionated wrapper libraries around the official OTEL SDKs and APIs, with Person A's team most deeply experienced with the Node/TypeScript offering. They also self-host an OpenTelemetry collector and generally stay up to date via Dependabot, though the collector version has lagged at times.

The primary motivation for adopting OpenTelemetry was vendor agnosticism. The company had previously migrated from New Relic to Datadog as their primary observability vendor and recognized that vendor lock-in was a recurring risk. By standardizing on OpenTelemetry, they ensured that the core instrumentation and collection layer would remain consistent regardless of which vendor sits on the receiving end. Person B emphasized that adoption was framed not as a mandate but as a strategic vision: teams consume the platform libraries without needing to know that OpenTelemetry is running behind the scenes. This abstraction has already proven its value — when the team moved log shipping from Sumo Logic to Datadog, the transition was significantly easier because the underlying OTEL-based collection architecture stayed the same.

Integration was largely straightforward for backend services. Person A noted that the SDK itself was a good experience and the documentation was helpful. The most time-consuming challenge was not with OpenTelemetry itself but with ensuring compatibility with Datadog's proprietary tracing header standards and other vendor-specific conventions. Person A also found Honeycomb's open-source SDK useful as a design reference when building the company's own wrapper libraries. The team contributed back to the project early on — a team member submitted a PR to the Go SDK when features were still missing in the project's earlier days. Cross-language maturity gaps (e.g., metrics vs. logs vs. traces at different stages across SDKs) required the team to carefully sequence which capabilities they rolled out to engineers, though this is no longer a significant issue.

The reference architecture provided by the OpenTelemetry project was singled out as the most valuable resource. The company adopted a similar model internally, creating their own reference architecture to demonstrate how to use their OTEL-based platform libraries. This pattern of "reference architecture as adoption accelerator" was highlighted as the single most important artifact they leveraged from the project.

OpenTelemetry has delivered measurable value at the company. Over one hundred production services have adopted the platform libraries. A major distributed tracing initiative tracked adoption across the organization, and the team found a correlation between observability adoption rates, telemetry volume flowing into Datadog, and improvements in time-to-detect for incidents. Beyond detection, the strategic value lies in reduced engineering effort when switching between vendors or trialing new tools — the standard libraries and collection infrastructure remain constant, giving the platform team flexibility to experiment without reimplementing instrumentation.

The project is fully in production at scale. The team is now in an "optimized usage" phase, shifting focus from adoption breadth to adoption depth — ensuring services are using tracing to its fullest extent rather than just having it enabled.

The most significant gap holding the project back, in the company's view, is front-end observability — particularly for client-rendered services and event-based communication patterns. The team struggles to get a complete, connected view from the end user's browser through to backend services due to context propagation challenges. Notably, they pointed out that their vendor (Datadog) has not solved this problem either, framing it as a genuine industry-wide innovation opportunity rather than a project-specific shortcoming.

Person A and Person B expressed satisfaction with their overall OpenTelemetry experience and acknowledged a desire to become more engaged with the community. Person A identified front-end observability as the strategic area where the company could contribute most meaningfully and asked about the best path in — whether through the SDK SIG, the End User SIG, or a dedicated working group. Person B noted that they would facilitate introductions to the OpenTelemetry Governance Committee and community managers, and that Person A would be speaking at an upcoming OTel Night event in New York.

If the project were to be archived, Person B was candid: the company would use it for a while but would eventually move on to a vendor-provided solution. They noted that the organization does not yet have a strong open-source maintenance culture, though they are working to develop one. Forking and long-term maintenance would be a significant burden given the breadth of the project.

Regarding CNCF maturity levels, both interviewees were familiar with the Sandbox, Incubating, and Graduated tiers. Person B noted that maturity level is a useful signal for understanding a project's strategic direction but is not a blocking factor in adoption decisions. The company has used Sandbox-level projects in Tier 0 production systems — for example, adopting Karpenter before it progressed beyond Sandbox and deprecating the cluster autoscaler in its favor. The key evaluation criteria are community health, extensibility, and whether the project is dominated by a single vendor or contributor, rather than the maturity label itself.
