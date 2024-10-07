# Propose Backstage to Incubation Stage 

## Table of Contents

- [Introduction](#introduction)
- [Background](#background)
  - [Recap: What is Backstage](#recap-what-is-backstage)
- [Progress Since Sandbox](#progress-since-sandbox)
  - [Community](#community)
  - [Technical](#technical)
- [Incubation Stage Requirements](#incubation-stage-requirements)
  - [Adhere to CNCF IP Policy](#adhere-to-cncf-ip-policy)
  - [Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope](#document-that-it-is-being-used-successfully-in-production-by-at-least-three-independent-end-users-which-in-the-toc’s-judgement-are-of-adequate-quality-and-scope)
  - [Clear versioning scheme & release methodology](#clear-versioning-scheme-&-release-methodology)
  - [Demonstrate a substantial ongoing flow of commits and merged contributions.](#demonstrate-a-substantial-ongoing-flow-of-commits-and-merged-contributions)
  - [Security](#security)
- [Statement on Alignment with the CNCF Mission](#statement-on-alignment-with-the-cncf-mission)
  - [Alignment with Other CNCF Projects](#alignment-with-other-cncf-projects)
- [Future Plans](#future-plans)

## Introduction

* Name: Backstage 
* Website: [backstage.io](https://backstage.io/)
* Source Control: [Backstage on Github](https://github.com/backstage/backstage)
* Code of Conduct: [Backstage Code of Conduct](https://github.com/backstage/backstage/blob/master/CODE_OF_CONDUCT.md)
* License: [Apache 2](https://github.com/backstage/backstage/blob/master/LICENSE)

## Background

Backstage joined CNCF as a sandbox project in September 2020. 

* [Link to Backstage Sandbox Proposal](https://github.com/cncf/toc/pull/476) 
* [Link to Backstage presentation to App delivery SIG](https://docs.google.com/presentation/d/1raY4eVh4zbkt1KAlSWXGkEfSwo2XN7UaDli1GWbh5p8/edit#slide=id.g7f5557e96c_0_18)
* [App delivery SIG meeting notes](https://docs.google.com/document/d/1OykvqvhSG4AxEdmDMXilrupsX2n1qCSJUWwTc3I7AOs/edit#heading=h.dhgrk589hocy)
* [Link to Github organisation](https://github.com/backstage/backstage)

### Recap: What is Backstage

[Backstage​​](https://backstage.io/) is an open platform for building developer portals. Powered by a centralized software catalog, Backstage restores order to your microservices and infrastructure and enables product teams to ship high-quality code quickly — without compromising autonomy.

Backstage unifies all your infrastructure tooling, services, and documentation to create a streamlined development environment from end to end and is based upon three core feature sets: 

* **[Create](https://backstage.io/blog/2020/08/05/announcing-backstage-software-templates)** - Backstage uses software templates to simplify setup, standardize tooling, and deploy with the click of a button. Using automated templates, engineers can spin up a new microservice, website, or other software component with your organization’s best practices built-in, right from the start.
* **[Manage](https://backstage.io/blog/2020/06/22/backstage-service-catalog-alpha)** - The software catalog is a centralized system that keeps track of ownership and metadata for all the software in your ecosystem (services, websites, libraries, data pipelines, etc). The catalog is built around the concept of metadata yaml files stored together with the code, which are then harvested and visualized in Backstage.
* **[Explore](https://backstage.io/blog/2020/09/08/announcing-tech-docs)**- How to find stuff? Backstage surfaces all the relevant information to users in relation to their components and services. From service metadata through the catalog, tooling to explore and manage through plugins, and documentation available right alongside the code. A single pane of glass to reduce context switching that can be built on with search or other discoverability tools, improving overall productivity and for engineers and users.

## Progress Since Sandbox

### Community

**Outreach:** As a project we hold regular community meetings each month, highlighting what is happening in the community with presentations from multiple adopters and discussions on current topics and interests in the community. Community sessions have on average **80 attendees**, are uploaded to a [community playlist hosted on Youtube](https://www.youtube.com/watch?v=3dV3aZo6JN8&list=PLf1KFlSkDLIBmA5TLXn2BzEHmwWzckP8y) and gained **4792 views**.

* [Backstage Community Hub on Github](https://github.com/backstage/community)
* [Backstage Community on Discord](https://discord.gg/gNPPdTHc)

The community statistics are extracted from [CNCF devstats](https://backstage.devstats.cncf.io/).


<table>
  <tr>
   <td><strong>Type</strong>
   </td>
   <td><strong>Start of Sandbox</strong>
   </td>
   <td><strong>In 2021</strong>
   </td>
   <td><strong>Current Total</strong>
   </td>
  </tr>
  <tr>
   <td>Adopters
   </td>
   <td>32
   </td>
   <td>28
   </td>
   <td>45
   </td>
  </tr>
  <tr>
   <td>Commits
   </td>
   <td>13,250
   </td>
   <td>8,540
   </td>
   <td>18,000
   </td>
  </tr>
  <tr>
   <td>Pull Requests
   </td>
   <td>3084
   </td>
   <td>1936
   </td>
   <td>4855
   </td>
  </tr>
  <tr>
   <td>Contributions
   </td>
   <td>34,107
   </td>
   <td>21,357
   </td>
   <td>284,000
   </td>
  </tr>
  <tr>
   <td>Forks
   </td>
   <td>903
   </td>
   <td>640
   </td>
   <td>1438
   </td>
  </tr>
</table>

### Technical

Since joining the sandbox, Backstage has technically evolved in multiple ways, some of the more substantial developments include;

* The plugin system has evolved in several ways
    * More flexibility in how plugins are added into the application
    * Cross version compatibility has been introduced to simplify maintenance
    * It is now easier to mix and match different open source and internal plugins
* Improved Catalog scalability
    * Horizontal scalability of the Catalog Backend with processing split into smaller chunks rather than one big loop.
    * Load tested and instrumented the catalog with a large number of entities.
    * Added processing cache to reduce computation in large environments.
* The introduction of Scaffolder Actions
    * Ability to create and compose custom actions - with a myriad of custom actions being built by adopters and the community. [#5153](https://github.com/backstage/backstage/pull/5153), [#5849](https://github.com/backstage/backstage/pull/5849), [#6340](https://github.com/backstage/backstage/pull/6340)
    * Horizontal scalability of the scaffolder backend
    * [https://backstage.io/blog/2021/07/26/software-templates-are-now-in-beta](https://backstage.io/blog/2021/07/26/software-templates-are-now-in-beta) 
* Introduced Backstage search alpha. With its composable frontend and extensible backend, engineers can design and build the search tool that suits their organization’s needs.
    * Integrates with sources from both within and outside the Backstage ecosystem
    * Pluggable search backends with implementations including Elasticsearch, Lunr and PostgreSQL.
    * [https://backstage.io/blog/2021/06/24/announcing-backstage-search-platform](https://backstage.io/blog/2021/06/24/announcing-backstage-search-platform)

## Incubation Stage Requirements

#### **_Adhere to CNCF [IP Policy](https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy)_**

As an existing Sandbox project this is already in place

#### **_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope_**

A list of ~40 public adopters can be found [here](https://github.com/backstage/backstage/blob/master/ADOPTERS.md). We’ve worked directly with ~75 companies in various stages of adopting, many of whom are non-public in their adoption.

The following are three highlighted examples of adopters: 

* [American Airlines](https://www.aa.com/) - A major American airline and the world's largest airline when measured by fleet size. In their [own words](https://tech.aa.com/2021-02-03-american-loves-open-source/); _Backstage, or “Runway” as we call it internally, has quickly become a central hub for creating new projects and has helped us radically streamline project creation._
* [Netflix](https://www.netflix.com/) - An over-the-top content platform and production company. Netflix looked at a lot of different options to fill their developer portal needs. They chose Backstage because it’s built for extensibility and is an open source solution supported by a strong community. [Watch the panel discussion](https://youtu.be/ajN9-dWSVYs?t=1012s).
* [Zalando](https://www.zalando.de/) - Online fashion platform Zalando decided to move away from their existing developer console and migrated to Backstage instead for its extensibility. With help from the open source community, Zalando’s small internal team went from proof-of-concept to building a platform serving thousands of developers. [Watch their presentation from our community event](https://youtu.be/6sg5uMCLxTA?t=175).
* [Telus](https://www.telus.com/en/) - Canada’s largest telecom. Realizing that trying to find the owner of services by yelling in Slack channels was no longer tenable, TELUS turned to Backstage for its focus on developer experience. [Watch their presentation from our community event](https://youtu.be/UZTVjv-AvZA?t=188).

#### **_Clear versioning scheme & release methodology_**

* Backstage packages follow [semantic versioning](https://semver.org/)
* We maintain a changelog for each of the Backstage packages within the repo, as example our [CLI changelog](https://github.com/backstage/backstage/blob/9d0c74b00d22e9467674ae9a5fed4c80501e59cd/packages/cli/CHANGELOG.md)
* As an indication of health and maturity we also maintain a stability index for the Backstage API

We also follow a regular release cycle of Thursday each week and have fully [documented this flow](https://backstage.io/docs/plugins/publishing#creating-a-new-release) to set proper expectations across the community for our releases.

#### **_Demonstrate a substantial ongoing flow of commits and merged contributions._**

* [Commit Activity](https://backstage.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-2y&to=now&var-period=w&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) 
    * In August 2021 there were a total of 676 contributions from 92 contributors. Up from a total of 35 contributors the year before.
    * There are a total of 489 total contributors to the project as of September 2021
* [New Contributors over past year](https://backstage.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)
    * On average we have 6 new contributors joining each week
* There are on average 49 new PRs from non-maintainers on the project per week
* In the last 30 days the average time to close PRs was 8 hours 11 minutes

Full details can be found at [CNCF devstats](https://backstage.devstats.cncf.io/) dashboards for Backstage.

#### **_Security_**

Security processes are clearly documented [here](https://github.com/backstage/backstage/blob/master/SECURITY.md).

## Statement on Alignment with the CNCF Mission

Backstage restores order to an organization's infrastructure and enables product teams to ship high-quality code quickly without compromising autonomy. Its goal is to enable teams and organisations to take control of the growing number of services and complexity they maintain.   As highlighted in the [CNCF charter](https://www.cncf.io/about/charter/),  Backstage aims to empower organisations to build and run modern scalable applications with ease, bringing together technologies and products that makes sense for them through a single pane of glass. 

In alignment with the [CNCF values](https://github.com/cncf/foundation/blob/master/charter.md#3-values), Backstage hosts a growing community of open plugins unlocking new functionality for the community and enables Backstage to run cloud agnostic to suit users needs. As a project we believe in an open Backstage that encourages the growth of our community helping to guide and develop Backstage both technically and in product direction. 

### Alignment with Other CNCF Projects

**_Does the project align and actively collaborate with other CNCF projects._**

The main mechanism for extending Backstage and configuring initial deployments for organizations is through our [plugin model](https://backstage.io/docs/plugins/#docsNav). Contributors, end-users, 3rd party vendors and organisations can extend Backstage functionality and surface information or tooling to support Backstage users. There is a growing [plugin marketplace](https://backstage.io/plugins) with currently 45 plugins for platform owners to integrate. Including some plugins for existing CNCF projects;

* [K8s plugin](https://backstage.io/docs/features/kubernetes/overview)
* [Argo](https://roadie.io/backstage/plugins/argo-cd/?utm_source=backstage.io&utm_medium=marketplace&utm_campaign=argo-cd)
* [Harbor](https://github.com/BESTSELLER/backstage-plugin-harbor/blob/master/README.md)
* [GitOps](https://github.com/backstage/backstage/tree/master/plugins/gitops-profiles)
* [GKE](https://github.com/BESTSELLER/backstage-plugin-gkeusage/blob/master/README.md) (alongside k8s)

We also make use of [Helm Charts](https://backstage.io/docs/deployment/#docsNav) as another CNCF integration to deploy Backstage.

## Future Plans

The direction of travel for Backstage is highly influenced by the community contributions and the core maintainers will continue to support the requests, as it has been until today. Saying that, the clear commitment is to continue to be “community driven”, scaling the community support accordingly with the increasing number of [adopters](https://github.com/backstage/backstage/blob/master/ADOPTERS.md), requests of support (in forums, chats, github issues, etc.) and contributions.

In addition to that, the core maintainers want to define a coherent roadmap for the whole product, focusing the improvements on strategic theming that we can summarize as follows.

* **Maturity,** to land to a stable version of the platform and a shared release plan, supporting the adopters with a clear lifecycle and cadence of releases.
* **Enablement** of new use cases and features, to extend the audience of the users who can benefit from the platform and empower them in doing better and faster (abstraction of the jobs to be done).
* **Foster innovation**, through the experimentation of the new features that the community can try and evaluate for future values and enhancements.

As a tangible plan for the platform enhancements developed by the core maintainers, but not only by them, you can check the [backstage.io roadmap page](https://backstage.io/docs/overview/roadmap) announcing the upcoming features for the ongoing quarter.

To get closer to the community, trying to take advantage of the momentum, we want to organize an in-person event entirely dedicated to Backstage, once the conditions allow it. The opportunity of a Backstage Developer Conference is loudly highlighted by the community and it would bring the involvement to the next level, where core-maintainers and contributors/enthusiasts will discuss together about the evolution of the platform.

The future ambitions for Backstage include the “sustainability” topic. Backstage is living a fast growth in adoption, maturity, and features, and part of its growth is about going beyond the current investments and investors. While we see the community growing with the current maintainers and investors, we would like to see more diversity in the core team. In addition, we believe that a truly successful open source project needs commercial ventures to continue to grow and thrive, and we are thinking of promoting initiatives for additional services specifically focusing enterprises and business use cases.
