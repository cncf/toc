# Open Policy Agent (OPA) Graduation Proposal

The Open Policy Agent (OPA) project joined the CNCF Sandbox in April 2018 and was promoted to the Incubating stage in April 2019. Since moving to the Incubating stage the project has demonstrated healthy growth in several areas including community, adoption, contributions, governance, features, and integrations. We believe the project is ready to move to the Graduation state.

Key Highlights:

* The community awareness and engagement metrics are up across-the-board year-over-year, e.g., openpolicyagent.org weekly page views increased 85% (to ~18,000 per week), Slack users increased 175% (to 2,680), DockerHub downloads increased over 10x (to ~13M), public GitHub repos containing .rego files increased 120% (to ~400), etc.

* [At KubeCon San Diego](https://blog.openpolicyagent.org/kubecon-us-2019-recap-3e60c70d633a) we held our first OPA user summit featuring talks from Capital One, Pinterest, TripAdvisor, Fugue, Atlassian, and Chef. Year over year, there were approximately twice as many industry presentations from organizations using OPA. More broadly, the community completed the [second annual OPA usage survey](https://blog.openpolicyagent.org/open-policy-agent-survey-summary-spring-2020-adaa46e61f0). The number of responses increased by 175% compared to the previous year. Over 150 organizations completed the survey with 47% reporting OPA running in production, 20% in pre-production or QA, and 24% in experimentation stages.

* The number of companies and number of developers contributing to OPA have increased by 36% and 57% year over year (ref: [devstats](https://opa.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1)).

* Since reaching the Incubating stage, OPA has shipped a number of impactful features that improve performance, usability, availability, security, operations, etc. The OPA Gatekeeper project shipped several features including multi-pod scalability, stable constraint & template interface, semantic logging, fine-grained metrics, dry-run support, standalone audit mode, readiness tracking, and more. The OPA Gatekeeper project has also become a viable alternative to K8s native PSPs. For a complete list of notable changes see the [OPA Releases](https://github.com/open-policy-agent/opa/releases) and the [OPA Gatekeeper Releases](https://github.com/open-policy-agent/gatekeeper/releases) pages.

* Finally, the popular `conftest` tool is now an official OPA sub-project. Conftest helps you write tests against structured configuration using OPA. For more information on Conftest see the [open-policy-agent/conftest](https://github.com/open-policy-agent/conftest) repo.

## Graduation State Criteria

### Have committers from at least two organizations.

We currently have [maintainers](https://github.com/open-policy-agent/opa/blob/master/MAINTAINERS.md) from 3 organizations (Styra, Google, and Microsoft). Approximately 75% of contributions over the past year come from these organizations (ref: [devstats](https://opa.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions)).

### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

We have met the [CII Best Practices passing level criteria](https://bestpractices.coreinfrastructure.org/en/projects/1768).

### Have completed an independent and third party security audit [...]

There have been two independent security audits carried out:

* In March 2020 an independent security audit was performed by Trail of Bits on Gatekeeper and OPA. The results can be found [here](https://github.com/open-policy-agent/gatekeeper/blob/master/SECURITY_AUDIT.pdf). Issues found during the audit were [tracked and resolved](https://github.com/open-policy-agent/gatekeeper/issues?q=+label%3A2020-security-audit+).

* In September 2018 after joining the Sandbox an independent security was performed by Cure 53 on the OPA daemon. The results can be found [here](https://github.com/open-policy-agent/opa/blob/master/SECURITY_AUDIT.pdf). Issues found during the audit were [tracked and resolved](https://github.com/open-policy-agent/opa/issues?q=is%3Aissue+label%3Aaugust-audit).

### Explicitly define a project governance and committer process [...]

OPA adopted a governance model when it joined the Sandbox. Recently we updated the model to clarify process around addition of new repositories and conflict resolution. The updated governance model also accounts for the different areas of expertise that each maintainer holds. The overall governance model is defined [here](https://github.com/open-policy-agent/opa/blob/master/GOVERNANCE.md). The maintainers are listed [here](https://github.com/open-policy-agent/opa/blob/master/MAINTAINERS.md).

### Have a public list of project adopters for at least the primary repo [...]

A list of some of the production adopters can be [here](https://github.com/open-policy-agent/opa/blob/master/ADOPTERS.md).

## Incubation Details

### Link to Incubation Due Diligence(DD) Document

[This document](https://docs.google.com/document/d/19M5fTpe57rQIMNxawRl5wSWvJUapuzY-CkV4O5pvieU/edit#) contains detailed due diligence as of August 2020.

In addition, there have been two previous rounds of due diligence performed on OPA:

* SIG-Security completed an assessment of the project in September 2019. The assessment can be found [here](https://github.com/cncf/sig-security/tree/master/assessments/projects/opa).

* @brendandburns performed a round of due diligence on the project before it was promoted to the Incubating stage. His comments can be found [here](https://github.com/cncf/toc/pull/199#issuecomment-470608626).

### Address any concerns or recommendations from the SIG and/or TOC sponsor(s) from the DD Document

The SIG-Security assessment made recommendations about improvements to documentation that have since been incorporated.