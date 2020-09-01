# Crossplane CNCF Sandbox Proposal

## Name of project

Crossplane

## Project description

### Introduction

Crossplane is a Kubernetes add-on used to extend and enhance clusters so users can provision and
manage infrastructure, managed services, and applications using kubectl, GitOps, or any tools that
can talk with the Kubernetes API. With Crossplane, cloud or on-prem infrastructure is represented as
a CRD that is configured using YAML and deployed into a cluster alongside existing application
configuration. Whether you're using a managed Kubernetes service or self hosting Kubernetes,
Crossplane can be installed into your cluster.

Crossplane empowers users to define, compose, and publish their own infrastructure abstractions with
declarative YAML on top of the base infrastructure CRDs resulting in new abstracted infrastructure
CRDs. With these abstractions, Crossplane lets you put in place policy guardrails to hide
infrastructure complexity from applications.

Crossplane supports application definition models allowing users to standardize on a single workflow
for both application and infrastructure deployment. The first specification Crossplane supports is
the Open Application Model (OAM) originally created by Alibaba and Microsoft. The OAM and Crossplane
communities have collaborated on a new version of the OAM specification that unifies the Crossplane
and OAM approaches to managing infrastructure and applications by fully embracing Kubernetes-based
runtimes. Crossplane also has great support for the native Kubernetes application model with
deployments and pods.

With this updated specification in place, Microsoft and Alibaba have joined the Crossplane community
and Crossplane is now [the Kubernetes implementation of the Open Application
Model](https://oam.dev/#implementations).

Crossplane functionality can be summarized by 3 main areas that will be covered in more detail in
the sections below:

1. Provision and manage any cloud infrastructure using the Kubernetes API
1. Publish abstract infrastructure resources as Kubernetes API types (CRDs)
1. Run and deploy applications alongside infrastructure

### Provision and Manage Cloud Infrastructure & Services

Crossplane extends Kubernetes with infrastructure provisioning and management capabilities.
Crossplane Providers include CRDs and controllers that operate on cloud provider APIs (GCP, AWS,
Azure, Alibaba) and on-prem infrastructure alike. Infrastructure configurations can be authored with
idiomatic Kubernetes YAML, versioned and managed with GitOps workflows, and provisioned manually
with kubectl or any tool that talks with the Kubernetes API.

### Publishing Custom Infrastructure Resources as Kubernetes API types

Crossplane also allows you to author YAML to define, compose, and publish your own composite
infrastructure resources on top of the infrastructure primitives included. Complex infrastructure
configurations, including policy guardrails, can be hidden behind your own infrastructure resources,
resulting in a layer of infrastructure resources that can be consumed dynamically and safely by
applications.

A practical example of this is an infrastructure owner wanting to provide a MySQL database to their
application developer teams and ensure that it is configured with secure networking connectivity and
according to their company policies and practices, without the developers needing to know any of
those complicated details.

With Crossplane, custom infrastructure resources can be authored with YAML (no code) that composes
the Crossplane CRDs for a managed service database (e.g., Amazon RDS), security groups, networking
configuration, and firewall rules.  All of this gets published as a new CRD that users can provision
alongside their existing application configurations.  The end result is secure and compliant
infrastructure that is created the same way the application itself is.

This also enables scenarios for multi-cloud abstractions, and creating an internal service catalog
of approved services and resources that applications can consume.

### Run and Deploy Applications Alongside Their Infrastructure

The Crossplane and OAM communities share a team-centric philosophy and believe a strong separation
of concerns combined with the proven Kubernetes declarative model is the best way to manage
infrastructure and cloud native applications.

Teams can publish infrastructure resources for applications to consume, define application
components independent of infrastructure, and compose both into a complete application configuration
-- all using idiomatic Kubernetes YAML that can be deployed with kubectl from app delivery pipelines
or with GitOps workflows.

Crossplane support for the [new version of the OAM
specification](https://github.com/oam-dev/spec/blob/master/design/20200105-spec-v1alpha2-kubernetes-friendly.md)
enables you to use Crossplane and OAM together by [installing the Crossplane helm
chart](https://crossplane.io/docs) into any Kubernetes cluster.

### Origin and History

The Crossplane project was first conceived in November 2017 by the same engineers that created
[Rook](https://rook.io/), another CNCF project that has matured and is currently under consideration
for graduation.  Within the Rook project, we saw firsthand the benefits Kubernetes provided for
storage resources from the Volume abstraction.  This was a first glimpse at application developers
declaring and managing infrastructure (storage volumes) alongside and in the same way they were
configuring their applications.  We quickly saw that the benefits from this approach could be
applied to many more types of infrastructure beyond just volumes, such as databases, buckets,
caches, clusters, etc.

A small team within [Upbound](https://upbound.io/) worked on the initial concepts, architecture,
controllers, CRDs, and documentation and made the [first public v0.1 release on December 4th,
2018](https://blog.upbound.io/introducing-crossplane-open-source-multicloud-control-plane/).  The
project has grown steadily since then and has continued to gain traction with early adopters as well
as some cloud providers.  Of significant note is the recent collaboration with Alibaba and Microsoft
to support the Open Application Model and standardize the way applications and infrastructure are
declared and managed. The project has also been integrated into GitLab to drive auto-devops
workflows on multiple cloud providers.

Crossplane has released multiple new versions of the project’s software and has most recently
released version [v0.11.0](https://github.com/crossplane/crossplane/releases/tag/v0.11.0).

## Statement on alignment with CNCF charter mission

The Crossplane project is well aligned with the CNCF’s charter and mission of making cloud native
computing ubiquitous. Crossplane follows the same API-centric, declarative configuration and
automation approach pioneered by Kubernetes, and extends it past application orchestration to all
the infrastructure and services applications need. The end result is users are now able to unify
application orchestration and infrastructure management around the same tooling and processes.

Crossplane enables best in class infrastructure and services of the cloud to be offered directly to
cloud native applications regardless of where they run. With Crossplane, it doesn’t matter if your
applications run in the cloud or on-prem, the project gives them the resources they need to run and
scale in modern, dynamic environments.  Furthermore, they can be deployed and consumed in a common
way with existing tools and familiar patterns from the ecosystem.

Crossplane is completely open source (Apache 2.0 license) and follows an open governance, in order
to foster and sustain a community of broad collaboration and progress across the entire ecosystem in
a vendor-neutral and agnostic way. The patterns and innovations of Crossplane that are helping to
make cloud native computing more ubiquitous are also accessible for everyone.

## Comparison with similar projects (inside or outside the CNCF), including what differentiates this project

While there are many projects that address similar issues, none of them encapsulate the full use
case that Crossplane addresses. This list is not exhaustive and is not meant to provide a deep
analysis of the following projects, but instead to motivate why Crossplane was created.

### Open Service Broker and Service Catalog

The [Open Service Broker](https://www.openservicebrokerapi.org/) and the [Kubernetes Service
Catalog](https://kubernetes.io/docs/concepts/extend-kubernetes/service-catalog/) are able to
dynamically provision cloud services from Kubernetes. As a result it shares similar goals with
Crossplane. However, service broker does not have the ability to define, compose, and publish your
own infrastructure resources to the Kubernetes API in a no-code way. Crossplane goes further by
enabling infrastructure operators to hide infrastructure complexity and include policy guardrails,
with a team-centric approach and a strong separation of concerns, so applications can easily and
safely consume the infrastructure they need, using any tool that works with the Kubernetes API.
Solutions like the [GCP implementation of Open Service
Broker](https://cloud.google.com/kubernetes-engine/docs/concepts/google-cloud-platform-service-broker)
have been deprecated in favor of a more Kubernetes-native solution, but one that is Google-specific
and closed source. Unlike the Kubernetes Service Catalog, Crossplane is implemented completely as
Kubernetes controllers and does not require an external service broker to operate.

### GCP Config Connector

The [GCP Config Connector](https://cloud.google.com/config-connector/docs/overview) is the GCP
replacement for Open Service Broker, and implements a set of Kubernetes controllers that are able to
provision managed services in GCP. It defines a set of CRDs for managed services like CloudSQL, and
controllers that can provision them via their cloud APIs. It is similar to Crossplane in that it can
provision managed services in GCP. Crossplane goes further by enabling you to provision managed
services from any cloud provider and the ability to define, compose, and publish your own
infrastructure resources in a no-code way. Crossplane supports a team-centric approach with a strong
separation of concerns, that enables applications to easily and safely consume the infrastructure
they need, using any tool that works with the Kubernetes API. GCP Config Connector is closed-source
and only supports GCP resources.

### AWS Service Operator

The [AWS Service Operator](https://github.com/awslabs/aws-service-operator) is a recent project that
implements a set of Kubernetes controllers that are able to provision managed services in AWS. It
defines a set of CRDs for managed services like DynamoDB, and controllers that can provision them
via AWS CloudFormation. It is similar to Crossplane in that it can provision managed services in
AWS. Crossplane goes further by enabling you to provision managed services from any cloud provider
and the ability to define, compose, and publish your own infrastructure API types in Kubernetes in a
no-code way. Crossplane supports a team-centric approach with a strong separation of concerns, that
enables applications to easily and safely consume the infrastructure they need, using any tool that
works with the Kubernetes API.

### AWS CloudFormation, GCP Deployment Manager, Azure Resource Manager and Others

These products offer a declarative model for deploying and provisioning infrastructure in each of
the respective cloud providers. They only work for one cloud provider, are generally closed source,
and offer little or no extensibility points, let alone being able to extend the Kubernetes API to
provide your own infrastructure abstractions in a no-code way. We have considered using some of
these products as a way to implement resource controllers in Crossplane. These projects use an
Infrastructure as Code approach to management, while Crossplane offers an API-driven control plane.

### Terraform and Pulumi

[Terraform](https://www.terraform.io/) and [Pulumi](https://www.pulumi.com/) are tools for
provisioning infrastructure across cloud providers. It offers a declarative configuration language
with support for templating, composability, referential integrity and dependency management.
Terraform can declaratively manage any compatible API and perform changes when the tool is run by a
human or in a deployment pipeline. Terraform is an Infrastructure as Code tool, while Crossplane
offers an API-driven control plane. Terraform can be used on-top of Crossplane using the standard
Kubernetes terraform provider.

## Sponsor from TOC

Sponsors from the TOC will be added as they are confirmed.

## Preferred maturity level

Sandbox

## License

Apache 2:
[https://github.com/crossplane/crossplane/blob/master/LICENSE](https://github.com/crossplane/crossplane/blob/master/LICENSE)

## Source control

The Crossplane project is composed of multiple source code repositories, all of which are hosted on
Github under the `crossplane` organization:
[https://github.com/crossplane](https://github.com/crossplane)

The core repository can be found at
[https://github.com/crossplane/crossplane](https://github.com/crossplane/crossplane).  A library of
reusable logic and interfaces that almost all controllers in the Crossplane logic depend on can be
found in
[https://github.com/crossplane/crossplane-runtime](https://github.com/crossplane/crossplane-runtime).

Some implementations of specific providers are also hosted in the Crossplane organization:

* [https://github.com/crossplane/provider-gcp](https://github.com/crossplane/provider-gcp)
* [https://github.com/crossplane/provider-aws](https://github.com/crossplane/provider-aws)
* [https://github.com/crossplane/provider-azure](https://github.com/crossplane/provider-azure)
* [https://github.com/crossplane/provider-rook](https://github.com/crossplane/provider-rook)
* [https://github.com/crossplane/provider-alibaba](https://github.com/crossplane/provider-alibaba)

## External dependencies (including licenses)

All dependencies (internal and external) of all projects in the Crossplane organization can be found
in the [Bill of Materials
(BOM)](https://docs.google.com/spreadsheets/d/1Z5kPXTWqeahBVnsbI2Q8NboRYQFxYBXdRIYnqKaHDJ0/edit?usp=sharing)
and [Audit/Due Diligence
report](https://drive.google.com/file/d/17zVD1jEE3jMGijx5DdNN_KpWtEFKMJs5/view?usp=sharing).  All
licensing data has been discovered and analyzed by [FOSSA](https://fossa.com/).  All projects in the
Crossplane organization pass the FOSSA license scan.

## Initial committers (how long working on project, companies they represent)

This list of initial committers represents maintainers of repositories within the Crossplane
organization that have write permissions to the respective repositories that they maintain.

| Name | GitHub | First Contribution | Organization |
| ---- | ------ | ------------------ | ------------ |
| Jared Watts | jbw976 | 2018-09-07 | Upbound |
| Bassam Tabbara | bassam | 2018-09-14 | Upbound |
| Nic Cope | negz | 2019-01-24 | Upbound |
| Daniel Mangum | hasheddan | 2019-01-21 | Upbound |
| Muvaffak Onus | muvaf | 2019-07-15 | Upbound |
| Hongchao Deng | hongchaodeng | 2020-03-09 | Alibaba |
| Artur Souza | artursouza | 2020-03-27 | Microsoft |
| Ryan Zhang | ryanzhang-oss | 2020-04-21 | Alibaba |
| Lei Zhang (Harry) | resouer | 2020-04-22 | Alibaba |
| Sun Jianbo | wonderflow | 2020-04-26 | Alibaba |

## Infrastructure requests (CI / CNCF Cluster)

We do not have any specific infrastructure requests at the current time.  All CI/CD is currently
self hosted using Jenkins and GitHub Actions.

## Communication channels (slack, irc, mailing lists)

The official communication channels for the project can be found in the README at
[https://github.com/crossplane/crossplane/blob/master/README.md#get-involved](https://github.com/crossplane/crossplane/blob/master/README.md#get-involved),
but are included here for convenience:

* Slack: [https://slack.crossplane.io/](https://slack.crossplane.io/)
* crossplane-dev mailing list:
  [https://groups.google.com/forum/#!forum/crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* info email: [info@crossplane.io](mailto:info@crossplane.io)
* Twitter: [https://twitter.com/crossplane_io](https://twitter.com/crossplane_io)

## Issue tracker

We prefer to track issues in the same code repository that the issue is relevant to, so each repo
under the Crossplane GitHub organization will have its own set of issues.  For example, the core
Crossplane repo issue tracker can be found at
[https://github.com/crossplane/crossplane/issues](https://github.com/crossplane/crossplane/issues).

## Website

The official public facing website can be found at [https://crossplane.io/](https://crossplane.io/)

## Release methodology and mechanics

The full release process is described in complete detail in this document:
[https://github.com/crossplane/crossplane/blob/master/docs/release/release-process.md](https://github.com/crossplane/crossplane/blob/master/docs/release/release-process.md)

Here is a brief summary of the release philosophy and methodology:

* Each repository under the crossplane organization can be independently versioned and released. In
    practice, we typically release all repos at the same time for convenience.
* All repositories follow strict [semantic versioning](https://semver.org/)
* A monthly release cadence is consistently followed, with a new release around the middle of every
    month
  * Minor version releases until v1.0
  * Latest release: v0.11.0
* Hotfixes (patch releases) are performed on an as needed basis after triaging issues for severity
    and impact
* Most non-library/package only repositories are released as container images to
    [DockerHub](https://hub.docker.com/orgs/crossplane)
* Build steps are automated with Makefiles and [Jenkins
    pipelines](https://jenkinsci.upbound.io/blue/pipelines/)
  * [https://github.com/crossplane/crossplane/blob/master/Makefile](https://github.com/crossplane/crossplane/blob/master/Makefile)
  * [https://github.com/crossplane/crossplane/blob/master/Jenkinsfile](https://github.com/crossplane/crossplane/blob/master/Jenkinsfile)
* Build process consistency is aided by reusable logic:
  * build submodule: [https://github.com/upbound/build/](https://github.com/upbound/build/)
  * scripts and pipelines: [https://github.com/crossplane/cicd](https://github.com/crossplane/cicd)

## Social media accounts

* Twitter: [https://twitter.com/crossplane_io](https://twitter.com/crossplane_io)
* Youtube:
  [https://www.youtube.com/channel/UC19FgzMBMqBro361HbE46Fw](https://www.youtube.com/channel/UC19FgzMBMqBro361HbE46Fw)
* Blog: [https://blog.crossplane.io/](https://blog.crossplane.io/)

## Community size and any existing sponsorship

The following statistics, when related to code repositories, are for the main crossplane repo of
[https://github.com/crossplane/crossplane](https://github.com/crossplane/crossplane).

* GitHub Stars: 2,120+
* Container downloads: 1,440,000+
* Releases: 24
* Commits: 2,215+
* Forks: 188+
* Contributors: 48+ contributors
* Maintainers: 10 maintainers from 3 organizations
* Twitter followers: 1,484+
* Slack members: 610+
* Youtube subscribers: 215+

### Who is currently known to be using the project? Are they using it in production and at what scale?

#### 45PRESS

45PRESS recently started evaluating Crossplane as a control plane powering their internal cloud
platform which is already built on Kubernetes. They’re building this platform to quickly, easily,
and programmatically provision infrastructure for their customers. Before Crossplane, 45PRESS would
have to spin up an environment manually whenever they got a new customer. Thanks to Crossplane, this
can be done automatically as users sign up for their services.

#### Alibaba

Alibaba’s application/serverless platform is fully OAM based. Alibaba collaborated with Crossplane
on [OAM Kubernetes runtime](https://github.com/crossplane/oam-kubernetes-runtime) which is the
candidate library for Alibaba’s web scale application/serverless systems both internally and core of
public cloud products.

#### GitLab

GitLab has [integrated Crossplane into GitLab
ADO](https://docs.gitlab.com/ee/user/clusters/crossplane.html) so users can easily provision and
manage infrastructure from their Kubernetes clusters. Crossplane has been shipped with every GitLab
release since version 12.5.

#### Mastercard

Mastercard is one of the largest credit card processors in the world. Internally, their
infrastructure environments span many different clouds and data centers. To unify them all, they’re
beginning to deploy Crossplane which will connect to services across clouds and provide teams with a
central point of management.

#### Microsoft

Recently, Microsoft collaborated with the Crossplane community to make Crossplane the Kubernetes
implementation of OAM application definition specification. Microsoft uses Crossplane to demo the
OAM specification on Kubernetes, showing how applications can be deployed across environments using
the OAM traits and scopes supported by Crossplane.

#### Red Hat

Crossplane recently integrated with the [Operator Framework](https://github.com/operator-framework)
to natively integrate with OpenShift as an alternative to the Open Service Broker API for
provisioning, managing, and consuming external services by cloud native applications. Red Hat’s
Emerging Technologies group and Upbound did a few Proof of Concept (PoC) validations with the goal
of running Crossplane as a fully supported OpenShift Operator installed from the OpenShift Operator
Catalog.

#### Swisscom

Swisscom is a major telecommunications provider in Switzerland. The company's Cloud Lab is
evaluating Crossplane to help give them a single interface for provisioning and managing services
from public and private cloud infrastructure.

#### Upbound

Upbound is building their marquee product, Upbound Cloud, using Crossplane. Upbound Cloud is an
enterprise grade managed service of Crossplane which layers on top of existing infrastructure and
transforms it to an internal cloud platform.

#### VSHN

VSHN is a leading Swiss company specializing in DevOps, Docker, Kubernetes, OpenShift and 24/7 Cloud
Operations. VSHN helps software developers run their apps and is integrating Crossplane into their
managed service offering of Kubernetes. With Crossplane, VSHN customers will be able to provision
the resources managed by VSHN that their applications need without having to submit a support
ticket.

## Project logo in svg format

All Crossplane artwork and assets can be found in the artwork repository:
[https://github.com/crossplane/artwork](https://github.com/crossplane/artwork)

The current main project logo in svg format can be found in that artwork repo here:
[https://raw.githubusercontent.com/crossplane/artwork/master/logo/logo-horizontal-bluetext.svg](https://raw.githubusercontent.com/crossplane/artwork/master/logo/logo-horizontal-bluetext.svg)
