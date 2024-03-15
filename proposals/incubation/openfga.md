# Table of Contents

- [Table of Contents](#table-of-contents)
- [About OpenFGA](#about-openfga)
- [Statement on Alignment with CNCF Mission](#statement-on-alignment-with-cncf-mission)
- [Progress since sandbox](#progress-since-sandbox)
- [Developer Community](#developer-community)
- [Incubation Stage Requirements](#incubation-stage-requirements)
  - [Adopters In Production](#adopters-in-production)
    - [1. Canonical](#1-canonical)
    - [2. Stacklok](#2-stacklok)
    - [3. Configu](#3-configu)
  - [Healthy Committers](#healthy-committers)
  - [Flow Of Contributions](#flow-of-contributions)
  - [Security](#security)
    - [License](#license)
- [Future plans](#future-plans)

# About OpenFGA
OpenFGA is a high-performance and flexible authorization solution that allows developers to build fine-grained access control using an easy-to-read modeling language and friendly APIs. 

Inspired by Google Zanzibar, OpenFGA is a centralized authorization engine that evaluates decisions by determining whether a relationship exists between an object and a user. Each check request references the authorization model against the known object relationships and returns an authorization decision (i.e. true or false).

- **Model any authorization system** - OpenFGA is inspired by the [Google Zanzibar paper](https://research.google/pubs/zanzibar-googles-consistent-global-authorization-system/) for Relationship-Based Access Control, and also solves problems for Role-based Access Control and some Attribute-Based Access Control use cases. The modeling language is powerful enough for engineers to create complex relationships but friendly enough for other stakeholders on the team to read and understand. 
- **Blazing fast** - OpenFGA is designed to answer authorization check calls in milliseconds across billions of relationships, which lets it scale with projects of any size. It works just as well for small startups building single applications as it does for enterprise companies building platforms on a global scale.
- **Works with existing code** - SDKs for several of the most popular languages have already been written, making it easy to integrate and grow alongside your applications.

# Statement on Alignment with CNCF Mission

As the world continues to move to a more digital, collaborative ecosystem of applications with ever-increasing objects, developers are scrambling to keep up and evolve their authorization systems to be more relationship-focused. But authorization is difficult to get right. [OWASP's Top 10 security risks](https://owasp.org/API-Security/editions/2023/en/0x11-t10/) include 3 on Authorization, with the top vulnerability being [Broken Object Level Authorization](https://owasp.org/API-Security/editions/2023/en/0xa1-broken-object-level-authorization/). 

Just like [Open Policy Agent](https://www.cncf.io/projects/open-policy-agent-opa/) for cloud infrastructure, application developers want a cloud-native option to add fine grained access control to their application logic without recreating a new solution every time they need to protect a new object type. Centralizing authorization enables application developers to build against a single predictable pattern regardless of their authorization needs. This approach to authorization will continue to serve them regardless of scale or pivoting through a digital transformation journey.

A list of CNCF projects that target solving access control in different ways can be found at [openfga/community/related-projects.md](https://github.com/openfga/community/blob/main/related-projects.md).

# Progress since sandbox
OpenFGA was proposed and accepted as a CNCF sandbox project on September 14, 2022 at the Sandbox maturity level.

Since then, OpenFGA has released:

- Added support for [MySQL as a storage backend](https://openfga.dev/docs/getting-started/setup-openfga/docker#using-mysql) ([#210](https://github.com/openfga/openfga/pull/210)).
- An [OpenTelemetry metrics integration](https://openfga.dev/docs/getting-started/setup-openfga/configure-openfga#telemetry) ([#360](https://github.com/openfga/openfga/pull/360)).
- Support for ABAC Conditional Relationships ([#1220](https://github.com/openfga/openfga/pull/1220)).
- A cross-platform [CLI](https://github.com/openfga/cli) to interact with an OpenFGA server.
- [Visual Studio Code extension](https://github.com/openfga/vscode-ext) to improve developer experience during model creation.
- A feature to list all the objects a user has a relationship ([#145](https://github.com/openfga/openfga/pull/145)).
- Support for [type restrictions](https://openfga.dev/docs/configuration-language#the-direct-relationship-type-restrictions) in the OpenFGA modeling language.
- Released a [Python](https://github.com/openfga/python-sdk) and [Java](https://github.com/openfga/java-sdk) SDKs.
- A [Helm Chart](https://github.com/openfga/helm-charts) to simplify deployment in Kubernetes clusters. 
- GitHub Actions for [testing](https://github.com/marketplace/actions/openfga-model-testing-action) and [deploying](https://github.com/marketplace/actions/openfga-model-deploy-action) models through CI/CD.
- An [ANTLR based language parser](https://github.com/openfga/language) with bindings for JS, Go and Java.

# Developer Community
It takes a lot of work from a lot of people to build a great, open-sourced authorization system. Which is why we value this community and receive genuine pleasure when we witness it rapidly growing. 

Since accepted as a CNCF sandbox project on September 2022:

- [Github stars](https://github.com/openfga/openfga/stargazers) grew from 398 → 2052 (**415%** increase)
- [Github forks](https://github.com/openfga/openfga) grew from 25 → 140 (**460%** increase)
- [OpenFGA repo contributors](https://github.com/openfga/openfga/graphs/contributors) grew from 14 → 44 (**214%** increase)
- [X (Twitter)](https://twitter.com/OpenFGA) followers grew from 206 → 542 (**163%** increase)
- [Developer activity](https://openfga.devstats.cncf.io/d/9/developer-activity-counts-by-repository-group-table?orgId=1) increased from 58 → 280 (**383%** increase)
- 676 [Discord](https://discord.com/channels/759188666072825867/930524706854031421) community members (actively migrating to the CNCF [#openfga](https://cloud-native.slack.com/archives/C06G1NNH47N) channel)
- [CNCF DevStats Dashboard](https://openfga.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1)

We host a monthly community meeting the 2nd Thursday of every month, at 11AM EST (8AM PST/3PM UTC). We have meet 19 times since September 2022, with about ~16 attendees each meeting. 
- [Meeting Minutes](https://docs.google.com/document/d/1Y6rbD0xpGLVl-7CmeMgxi56_a0ibIQ_RojvWBbT9MZk/edit#)
- [Recordings](https://www.youtube.com/playlist?list=PLUR5l-oTFZqUneyHz-h4WzaJssgxBXdxB) of past meetings
- [Youtube playlist](https://www.youtube.com/playlist?list=PLUR5l-oTFZqUAdAibhLw7l5IdqDnQ5gga) of presentations by the community


Highlighted community blogs:
- [Introduction to OpenFGA](https://www.ericksegaar.com/2024/03/04/introduction-into-openfga/) by Erick Segaar
- [Fine-grained authorization with OpenFGA and OAuth](https://www.godaddy.com/resources/news/authorization-oauth-openfga) by Jacob Brooks (GoDaddy)
- [Keycloak integration with OpenFGA](https://embesozzi.medium.com/keycloak-integration-with-openfga-based-on-zanzibar-for-fine-grained-authorization-at-scale-d3376de00f9a) by Martin Besozzi
- [Relation-based access control using Verifiable Credentials](https://medium.com/@excid/relation-based-access-control-using-verifiable-credentials-d8e542a0ce1) by ExcID

# Incubation Stage Requirements
## Adopters In Production

A list of OpenFGA adopters can be found at [openfga/community/ADOPTERS.md](https://github.com/openfga/community/blob/main/ADOPTERS.md), plus many more that haven’t been disclosed.

Three production examples to highlight:

### 1. Canonical
They are embedding OpenFGA into several different layers of their Ubuntu Pro stack.

- [LXD](https://discourse.ubuntu.com/t/lxd-5-20-has-been-released/40865#authorization-restructure-7) - a container and virtual machine manager
- [MicroCloud](https://www.gamingdeputy.com/canonical-unveils-microcloud-a-toolkit-for-rapid-cluster-deployment/) - a deployment center for computing clusters with shared distributed data storage and a secure virtual network
- [JAAS](https://github.com/canonical/openfga-operator) - a managed Juju orchestration engine as a service
 

### 2. Stacklok
Stacklok recently [revamped their authorization model and engine in Minder](https://stacklok.com/blog/using-openfga-to-build-a-relationship-based-authorization-model-in-minder), an open source software supply chain security platform. They switched from a database-backed authorization implementation using [Open Policy Agent](https://www.openpolicyagent.org/) to a multi-tenant, relationship-based authorization model using [OpenFGA](https://openfga.dev/). 

### 3. Configu
Configu is an open source software for streamlining, testing, and automating application configurations across environments. They [specifically picked OpenFGA](https://configu.com/blog/authorization-over-configurations-using-openfga/) because it was a CNCF backed third-party authorization system that allows them to build upon battle-tested authorization standards saving them valuable implementation time not recreating the wheel for a problem that has already been solved for developers. 

## Healthy Committers

OpenFGA welcomes feedback and contributions from anyone in the community who might be interested in helping the project. Due to its multi-repo structure with distinct contributor roles, the project uses "maintainer" as the noun for a member with sufficient permissions to merge PR in specific repo. This is the group of people responsible for organizing the backlog of issues in OpenFGA and other issues across repositories, reviewing pull requests, and all code within this repository.

OpenFGA had [91 committers in the last 12 months](https://all.devstats.cncf.io/d/53/projects-health-table?orgId=1) and [19 active maintainers](https://github.com/openfga/community/blob/main/MAINTAINERS).

## Flow Of Contributions 

All notable changes in OpenFGA projects are documented in a [CHANGELOG](https://github.com/openfga/openfga/blob/main/CHANGELOG.md) file. The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) and adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). 

Detailed statistics can be found in the following openfga.devstats.cncf.io links:
- [Contribution activity](https://openfga.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now-2d): 1450 average contributions per month
- [Issues open/closed](https://openfga.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All): average 46.5 issues opened and 28 issues closed per month
- [PRs open/closed since sandbox](https://openfga.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All): average 159 PRs opened and 144 PRs closed per month
- [Open PR Age](https://openfga.devstats.cncf.io/d/19/open-pr-age-by-repository-group?orgId=1&var-period=d7&var-repogroup_name=All): 5 hours 1 min median 7 day moving average 


## Security

View the [security self assessment](https://github.com/cncf/tag-security/blob/main/assessments/projects/openfga/self-assessment.md).

We have an [open PR](https://github.com/cncf/tag-security/pull/1235) pending to be reviewed by the TAG Security team.

OpenFGA vulnerability management is described in the official project security documentation [SECURITY.md](https://github.com/openfga/.github/blob/main/SECURITY.md).

### License

OpenFGA is licensed under the [Apache 2.0 license](https://github.com/openfga/openfga/blob/main/LICENSE).

# Future plans

OpenFGA follows an [RFC (Request For Comments)](https://github.com/openfga/rfcs) process for substantial changes to the project. An RFC is a way to propose, communicate and coordinate on new efforts for the project.

OpenFGA updates a [visible roadmap](https://github.com/orgs/openfga/projects/1) for anyone to view.

Two major upcoming improvements:

1. [Modular Authorization Models](https://github.com/orgs/openfga/projects/1?pane=issue&itemId=35518190) ([RFC](https://github.com/openfga/rfcs/blob/main/20231212-modular-models.md))- Splits a single authorization model into smaller “modules” allowing Large enterprise organizations with multiple teams to update their part of the authorization model maintained by that team without affecting other teams. 
1. [ListUsers API endpoint](https://github.com/orgs/openfga/projects/1?pane=issue&itemId=24490812) ([RFC](https://github.com/openfga/rfcs/pull/15))- In the same way we provide a [ListObjects endpoint](https://openfga.dev/api/service#/Relationship%20Queries/ListObjects) that list all resources for a specific user & relation, we will provide a way to list which users have a specific relationship with a specific resource.

Plus, we plan to release a slew of updates to the SDKs, like [Java Springboot security improvements](https://github.com/orgs/openfga/projects/1/views/1?pane=issue&itemId=10720070) and [JetBrain IDE integration](https://github.com/orgs/openfga/projects/1/views/1?pane=issue&itemId=49602107). We also plan to release a [DynamoDB storage adapter](https://github.com/orgs/openfga/projects/1/views/1?pane=issue&itemId=49602035) to complement the Postgres and MySQL storage adapters.