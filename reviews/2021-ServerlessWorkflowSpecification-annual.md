# Serverless Workflow Specification 2021 Annual Review

- [Background](#background)
- [Alignment with Cloud Native](#alignment-with-cloud-native)
- [Year in Review](#year-in-review)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

Serverless Workflow is a vendor-neutral, open-source, and fully community-driven ecosystem
for defining and running DSL-based workflows that target the serverless technology domain.

This project is composed of:

* [Specification](https://github.com/serverlessworkflow/specification/blob/main/specification.md) for defining DSL-based workflows.
* [Developer SDKs](https://github.com/serverlessworkflow/specification#sdks) that provide support for many programming languages.
* [Workflow runtimes](https://github.com/serverlessworkflow/specification#runtime) part of the project ecosystem, and support the execution of specification DSL.
* [Developer tooling](https://github.com/serverlessworkflow/specification#Tooling) support for writing DSL-based workflows.

Serverless Workflow was approved as a Cloud Native Sandbox level project on July 14, 2020.
* [TOC review PDF](https://github.com/serverlessworkflow/specification/blob/main/community/presentations/2020-4-15-toc-pres.pdf).
* [TOC sandbox proposal PR](https://github.com/cncf/toc/pull/376)

## Alignment with Cloud Native

Serverless Workflow project falls under the [CNCF "App Definition and Development"](https://landscape.cncf.io/card-mode?category=app-definition-and-development&grouping=category) category.

Serverless Workflow is a member project of the [CNCF Serverless Working Group](https://github.com/cncf/wg-serverless).

Serverless Workflow includes [Synapse](https://github.com/serverlessworkflow/synapse), a Kubernetes-native runtime engine for executing workflows that follows the
specification DSL definition.

In addition, Serverless Workflow provides support for several other open-source projects and specifications in the cloud-native
space:
* CloudEvents
* OpenAPI
* AsyncAPI
* GraphQL
* OData
* OAuth2

## Year in Review

This year was very exciting for the project. Some of the most notable accomplishments include:
* Released specification version [0.6](https://github.com/serverlessworkflow/specification/releases/tag/v0.6)
* Released specification version [0.7](https://github.com/serverlessworkflow/specification/releases/tag/v0.7)
* Added [Synapse](https://github.com/serverlessworkflow/synapse), a Kubernetes-native runtime into our ecosystem
* Added [sdk-net](https://github.com/serverlessworkflow/sdk-net), a .NET SDK into our ecosystem
* Added [sdk-typescript](https://github.com/serverlessworkflow/sdk-typescript), a TypeScript SDK into our ecosystem
* Added [two new project maintainers](https://github.com/serverlessworkflow/specification/blob/main/MAINTAINERS.md)

From the community perspective we also had a good year:
* Over 100 new followers on our [twitter channel](https://twitter.com/CNCFWorkflow).
* Over 200 new stars on our [specification github repo](https://github.com/serverlessworkflow/specification).
* Over 300 people attending our project office hours at 2021 KubeCon EU.
* Presented at KubeCon EU 2021, KubeCon NA 2020, and KubeCon EU 2020
* Mentioned as a key component of open-source microservices architectures at [InfoQ](https://www.infoq.com/articles/microservices-inside-out/)
* Over 40 unique visitors per day on our [website](https://serverlessworkflow.io/)
* Participated at the KubeCon EU 2021 BugBash

## Annual Review Contents

- **Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity.**

Project [DevStats page](https://serverlessworkflow.devstats.cncf.io).
The info for the span of one year shows:
* Over 500 merged PRs
* ~2000 commits by 30+ contributors
* Community contributions from 10+ different companies
* Over 100% increase to github stars compared to last year

- **How many maintainers do you have, and which organizations are they from?**

Serverless Workflow currently has [5 project maintainers](https://github.com/serverlessworkflow/specification/blob/main/MAINTAINERS.md)
- Tihomir Surdilovic, Temporal Technologies
- Manuel Stein, Nokia Bell Labs
- Ricardo Zanini, Red Hat
- Charles d'Avernas , Neuroglia
- Antonio Mendoza Pérez, Independent


- **What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?**

Both the adoption and community interest has been steadily increasing over the course of the year.
Most notable adoptions have been by:

- [Kogito](https://kogito.kie.org/), a Red Hat project automation runtime
- [Automatiko](https://automatiko.io/), a workflow automation runtime
- [Synapse](https://github.com/serverlessworkflow/synapse), a Kubernetes-based workflow runtime which has joined the Serverless Workflow ecosystem

We also have a number of integrations that are currently work-in-progress which include
integrations with [Temporal](https://temporal.io/).

- **How has the project performed against its goals since the last review?**

This is our first annual review since becoming a Sandbox project.
We have surpassed all of our goals that we have set for this year.
We have been able to not only surpass the goals of the main specifications,
but also to go from just hosting a specification to creating
an entire workflow ecosystem around it. This includes SDKs, Tooling (VSCode, IntelliJ),
runtimes (Synapse), etc.

Over this year we were able to create a workflow DSL which is in our opinion
at this time the most feature-rich and most powerful workflow DSL that exists.

- **What are the current goals of the project?**

Specification roadmap: https://github.com/serverlessworkflow/specification/tree/main/roadmap

Our main goals for the project include:
- Release specification version 1.0 by end of 2021
- Add more SDKs in different languages
- Create a specification TCK
- Add integrations with different workflow DSLs
- Improve our community tooling support
- Add at least 2 more integrations with existing workflow runtimes by middle of 2022


- **How can the CNCF help you achieve your upcoming goals?**
- Help us in promoting the project (Blogs, Twitter, KubeCon, etc)
- If feasible help our project via crowdfunding [here](https://crowdfunding.lfx.linuxfoundation.org/projects/serverless-workflow)
- Keep providing us with opportunities to have project office hours and talks at KubeCons


- **Do you think that your project meets the criteria for incubation?**

We believe we have made significant progress toward this goal and that the project is ready for incubation.
We would like however to release Serverless Workflow specification v1.0 before starting this process
(scheduled for end-of-year 2021).

## Project Links
* [Website](https://serverlessworkflow.io/)
* [GitHub](https://github.com/serverlessworkflow)
* Slack:[CNCF](http://slack.cncf.io) / #serverless-workflow
* [Twitter](https://twitter.com/CNCFWorkflow)