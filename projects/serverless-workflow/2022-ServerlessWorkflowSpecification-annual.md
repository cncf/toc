# Serverless Workflow Specification 2022 Annual Review

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

Serverless Workflow includes [Synapse](https://github.com/serverlessworkflow/synapse), a Kubernetes-native runtime engine for executing workflows that follows the specification DSL definition.

In addition, Serverless Workflow provides support for several other open-source projects and specifications in the cloud-native
space:

* CloudEvents
* OpenAPI
* AsyncAPI
* GraphQL
* OData
* OAuth2

## Year in Review

This year was for stabilization and discussions around the upcoming 1.0 release. The ecosystem kept growing:

* Two new releases for the [sdk-java](https://github.com/serverlessworkflow/sdk-java/releases)
* One new release for the [sdk-typescript](https://github.com/serverlessworkflow/sdk-typescript/releases)
* Four new releases for the [sdk-go](https://github.com/serverlessworkflow/sdk-typescript/releases)
* Seven minor releases for the [sdk-net](https://github.com/serverlessworkflow/sdk-net/tags)
* Four minor releases for the [Synapse](https://github.com/serverlessworkflow/synapse/releases) runtime
* Added [sdk-python](https://github.com/serverlessworkflow/sdk-python), a Python SDK into our ecosystem

From the community perspective we also had a good year:

<!-- I couldn't get this number for Twitter 
* Over 100 new followers on our [twitter channel](https://twitter.com/CNCFWorkflow).
-->
* Over 210 new stars on our [specification github repo](https://github.com/serverlessworkflow/specification).
* Over 200 people attending our project kiosk at KubeCon NA 2022
* Presented in-person at KubeCon EU 2022 for more than 100 people.
* Over 200 people attending our project office hours at 2022 KubeCon EU.
* Over 50 people attending our project office hours at 2022 KubeCon China.

## Annual Review Contents

- **Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity.**

Project [DevStats page](https://serverlessworkflow.devstats.cncf.io).
The info for the span of one year shows:

* Over 50 new forks
* Over 80% increase to Synapse (runtime implementation) github stars compared to last year
* Over 101% increase to github stars compared to last year

- **How many maintainers do you have, and which organizations are they from?**

Serverless Workflow currently has [3 project maintainers](https://github.com/serverlessworkflow/specification/blob/main/MAINTAINERS.md)

<!-- Alphabetical order -->
- Charles d'Avernas, Neuroglia
- Ricardo Zanini, Red Hat
- Tihomir Surdilovic, Temporal Technologies

- **What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?**  

Both the adoption and community interest has been steadily increasing over the course of the year.
Most notable adoptions have been by:

<!-- Alphabetical order -->
- [Apache EventMesh](https://eventmesh.apache.org/), a new generation serverless event middleware for building distributed event-driven applications
- [Automatiko](https://automatiko.io/), a workflow automation runtime
- [FaasNet](https://github.com/simpleidserver/FaasNet), FaasNet makes it easy to deploy functions and API to Kubernetes without repetitive, boiler-plate coding.
- [OpenShift Serverless Logic](https://developers.redhat.com/articles/2022/08/15/how-openshift-serverless-logic-evolved-improve-workflows), a Red Hat product under Tech Preview integrated with their flagship product, OpenShift
- [Synapse](https://github.com/serverlessworkflow/synapse), a Kubernetes-based workflow runtime which has joined the Serverless Workflow ecosystem

These are the companies that have adopted the Serverless Workflow Specification:

<!-- Alphabetical order -->
- [CAF](https://caf.io), Serverless Workflow is the core technology behind every KYC/KYB solution allowing them to customize it for their clients seamlessly.
- [Huawei](https://www.huaweicloud.com/intl/en-us/product/functiongraph.html), Huawei FunctionGraph hosts event-driven functions in a serverless context while ensuring high availability, high scalability, and zero maintenance.
- [IBM](https://www.ibm.com/) As active members of the open-source KIE community, the BAMOE team from IBM's Digital Business Automation division is highly committed to standards within the business automation domain, and CNCF Serverless Workflow is no different. Eventually, IBM plans to incorporate the CNCF Serverless Workflow format into its product offering, providing more choices to customers to take advantage of BAMOE workflow capabilities.
- [Neuroglia](https://neuroglia.io/), Neuroglia is a consultancy and solution design company for the digital transformation of companies and their services.
- [OpenEnterprise](https://automatiko.io/), OpenEnterprise Automatiko helps you build better services and functions based on workflows expressed with well known standards.
- [Red Hat](https://redhat.com/), Red Hat sponsors the development of Kogito Serverless Workflow, which is a tool for building cloud-native workflow applications.
- [Tantl](https://www.tantl.com/), Tantl is making it easy for developers to build internal workflows, such as allowing customer support reps to quickly process refunds.
- [Temporal](https://temporal.io/), Temporal is the open source microservice orchestration platform for writing durable workflows as code.

There are a few other companies that are in touch with us and using the specification, but can't disclose at the moment.

- **How has the project performed against its goals since the last review?**

This is our second annual review since becoming a Sandbox project.
This year we looked for stabilization of the specification by having more
discussions with current implementation project leaders to achieve a good
balance between the standards and production-level use cases.

We decided to grow slowly and now we are reaching to a point to fix most of
the open issues and discussions before releasing 0.9 version, and then the final
1.0 by the end of the year.

One of the goals we achieved, was to set a new governance model to balance the
responsibilities amongst all the maintainers.

We had to do a little detour duo to the progress of projects implementing and using
the specification in many production-level use cases.

- **What are the current goals of the project?**

Specification [roadmap](https://github.com/serverlessworkflow/specification/tree/main/roadmap) and [progress tracker](https://github.com/orgs/serverlessworkflow/projects/1/views/2).

Our main goals for the project include:

- Release specification version 1.0 by late 2023
- Create a specification TCK
- Add integrations with different workflow DSLs
- Improve our community tooling support

- **How can the CNCF help you achieve your upcoming goals?**

- Help us in promoting the project (Blogs, Twitter, KubeCon, etc)
- If feasible help our project via crowdfunding [here](https://crowdfunding.lfx.linuxfoundation.org/projects/serverless-workflow)
- Keep providing us with opportunities to have project office hours and talks at KubeCons

- **Do you think that your project meets the criteria for incubation?**

We believe we have made significant progress toward this goal and that the project is ready for incubation.
We would like however to release Serverless Workflow specification v1.0 before starting this process
(scheduled for late 2023).

## Project Links

* [Website](https://serverlessworkflow.io/)
* [GitHub](https://github.com/serverlessworkflow)
* Slack:[CNCF](http://slack.cncf.io) / #serverless-workflow
* [Twitter](https://twitter.com/CNCFWorkflow)
