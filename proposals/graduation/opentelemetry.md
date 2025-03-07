# OpenTelemetry Graduation Proposal

Since moving to incubation in August 2021, OpenTelemetry has made great strides in its mission to provide ubiquitous, high-quality telemetry of cloud-native systems for observability. OpenTelemetry has maintained its incredible [project velocity](https://www.cncf.io/blog/2023/01/11/a-look-at-the-2022-velocity-of-cncf-linux-foundation-and-top-30-open-source-projects/), the second highest in the CNCF ecosystem behind Kubernetes. Over [150 companies, and over 1000 developers](https://opentelemetry.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1) contribute to the project each month (as of March 2024), and OpenTelemetry is used in production at member organizations such eBay, Adobe, Apple, and many more. In addition, we’ve shipped a [demo](https://github.com/open-telemetry/opentelemetry-demo) to show our work, and continued our mission to align standards across the observability space by absorbing the Elastic Common Schema into our [Semantic Conventions](https://github.com/open-telemetry/semantic-conventions), which are expanding to include conventions for the observability of emerging technology like Generative AI (Large Language Models). We continue to push the state of the art through our work in adopting [continuous profiling](https://github.com/open-telemetry/oteps/blob/main/text/profiles/0239-profiles-data-model.md) and [client side instrumentation](https://github.com/open-telemetry/oteps/blob/main/text/0202-events-and-logs-api.md) as crucial components of cloud-native observability. Our current and future work is documented on our [project’s Kanban board](https://github.com/orgs/open-telemetry/projects/29).

In 2023, we achieved our primary goals – stabilizing tracing, metrics, and logging signals, as well as sunsetting OpenTracing and OpenCensus, our progenitor projects. Our journey towards OpenTelemetry becoming a truly built-in part of cloud-native systems begins here, with graduation.

## Graduation State Criteria

OpenTelemetry, by design, is a very broad project. We support 13+ languages,
each maintained by the community. This breadth makes it challenging to perform
certain graduation requirements in each repository. Thus, the GC proposes a more
narrowly scoped set of components that we define as “Core Components”. These
repositories are as follows:

- [OpenTelemetry Specification](https://github.com/open-telemetry/opentelemetry-specification)
- [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-collector)
- [OpenTelemetry Go](https://github.com/open-telemetry/opentelemetry-go)
- [OpenTelemetry Java](https://github.com/open-telemetry/opentelemetry-java)
- [OpenTelemetry .NET](https://github.com/open-telemetry/opentelemetry-net)
- [OpenTelemetry Python](https://github.com/open-telemetry/opentelemetry-python)

> Core components have been identified based on several factors, including GitHub Stars, opentelemetry.io docs traffic, and ecosystem importance.

### Have committers (maintainers) from at least two organizations

OpenTelemetry's governance process ensures that all governance functions include
members from multiple organizations. A maximum of [two members of
governance](https://github.com/open-telemetry/community/blob/main/governance-charter.md#maximal-representation)
can be from the same organization, and [no more than 25% of the technical
committee can be from the same
organization](https://github.com/open-telemetry/community/blob/main/tech-committee-charter.md#no-over-representation).
In addition, individual SIGs are encouraged to have broad representation from
multiple organizations.

### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/)

Yes, we have achieved this for some repositories (such as the [Collector](https://www.bestpractices.dev/en/projects/8404)), and we anticipate rolling it out to other core components this year.

### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation

An independent, third-party security audit was completed of our core components, and all high-priority issues were dealt with during its duration.
You can [read our announcement](https://opentelemetry.io/blog/2024/security-audit-results/) that contains a link to the audit report.

### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers

Project governance, committer lifecycle, and other documents can be found in the [community repository](https://github.com/open-telemetry/community)
[Governance Charter](https://github.com/open-telemetry/community/blob/main/governance-charter.md). A representative list of maintainers, approvers, and their emeritus can be seen in the [Collector README](https://github.com/open-telemetry/opentelemetry-collector/blob/main/README.md)

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence

Maintainer lists and governance lists are located in the [community repository](https://github.com/open-telemetry/community). The [Governance Charter](https://github.com/open-telemetry/community/blob/main/governance-charter.md) describes how elections are performed; the Governance Committee is the primary interface between the project and the CNCF.

### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to [FAQs](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for guidelines on identifying adopters

Adopters can be found on [opentelemetry.io](https://opentelemetry.io/ecosystem/adopters/)
