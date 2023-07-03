# Schemahero 2022 Annual Review

This is the Kubernetes TOC Annual Review for the [Schemahero](https://schemahero.io) project for 2022.

## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption) 
- [Project goals](#project-goals) 
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)
- [Project Links](#project-links)

## Background

Schemahero is an implementation of declarative database schema management. The goal of the Schemahero project is to simplify management of database schema changes in various environments. Schemahero allows developers to define the desired state of a database schema and the Schemahero project will generate the migration commands needed to migrate an existing database schema to the desired schema.

Declarative change management is the cloud-native way to manage a database schema. Similar to how Kubernetes brings a common API to define the desired state of a cluster, Schemahero implements a common API to define the desired schema of a database.

In addition to schema management, Schemahero aims to implement approval workflows to simplify change management of schemas in organizations that have strict policies or regulatory requirements to manage the database. Schemahero can work with change management tools that integrate with git to provide a policy-driven approach to database schema management.

Schemahero's initial implementation is a Kubernetes operator, making it easy to manage any supported database schema using existing Kubernetes CI/CD tooling, including gitops workflows. This also allows Schemahero manifests to be compatible with many other CNCF projects for policy validation and enforcement, deployment pipelines, change management, and more.

Schemahero was accepted as a CNCF Sandbox project on [October 10th, 2020](https://docs.google.com/spreadsheets/d/1Nnh_usr0tSZxaUpxTusqeIqKxMmvuEViRkyO9e_Do40).

## DevStats

> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

The Schemahero DevStats dashboards can be found [here](https://schemahero.devstats.cncf.io/d/8/dashboards?orgId=1&viewPanel=2&from=now-2y&to=now-1h&editPanel=2). 

We are using the following metrics from the DevStats metrics as key indicators of community health:

* [Stargazers and Forks](https://schemahero.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1): 265% annual growth from 194 -> 516. Stars are a key way for us to measure awareness of the project.

* [Commits per week](https://schemahero.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=w&var-repogroups=All&from=now-6M&to=now): Always looking for consistent and then increased activity in the projects

* [Contributors and Companies](https://schemahero.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&var-period=d7&var-repogroup_name=All&from=now-1y&to=now): This has recently started to become a key metric that we want to transition from consistent to increasing. More on this goal below.

## Maintainers

Schemahero currently has 7 maintainers, but currently all are from a single company:

| Maintainer           | GitHub Username                                     | Affiliation |
| -------------------- | --------------------------------------------- | ----------- |
| Marc Campbell           | [@marccampbell](https://github.com/marccampbell)   | Replicated  |
| Dmitriy Ivolgin          | [@divolgin](https://github.com/divolgin)   | Replicated  |
| Ethan Mosbaugh        | [@emosbaugh](https://github.com/emosbaugh)    | Replicated |
| Andrew Lavery          | [@laverya](https://github.com/laverya)   | Replicated   |
| Dan Stough | [@danstough](https://github.com/danstough) | Replicated |
| Salah Aldeen Al Saleh | [@sgalsaleh](https://github.com/sgalsaleh) | Replicated |
| Treva Nichole Williams | [@OGtrilliams](https://github.com/OGtrilliams) | Replicated |


The maintainer list is managed in our [GitHub repository](https://github.com/schemahero/schemahero/blob/main/OWNERS.md).  Our [governance policy](https://schemahero.io/community/) is maintained on the project website.

## Adoption

In October 2021, we created an Adopters file to start to track publicly referenceable uses of the projects. Adopters who are publicly referenceable will be listed in the [SchemaHero GitHub repository ADOPTERS file](https://github.com/schemahero/schemahero/blob/main/ADOPTERS).

Schemahero adopters include:
 
**Projects**
* [KOTS](https://github.com/replicatedhq/kots) 

**Organizations**
* [Replicated](https://www.replicated.com/)

**Evaluating Organizations**
* [Tyro](http://www.tyro.com/)
* Several unnamed organizations are evaluating (but not yet selected) Schemahero

## Project Goals

> How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

When we were accepted as a Sandbox project, it was early for Schemahero. Our initial goals were to share the idea and leverage the CNCF platform to get feedback from additional potential users of the project. To do this, we needed to focus on some specific features such as additional database support and functionality such as secrets management (integration with Vault and others) in order to ensure that the Schemahero solution is broad enough to be adopted in various architectures.

To that end, we've released several significant minor versions and have evolved the project from 2 databases with 1 methods of secrets to now supporting 6 databases with 3 secrets management methods. This does allow Schemahero to be compatible with more environments, so we've made significant progress against our previously stated goals.

The following list will show our 4 previous areas of focus with a status update on each:

* **Additional database support**: When joining the sandbox, Schemahero supported Postgres and Mysql. Today we support Postgres, Mysql, Cockroachdb, Yugabyte, Cassandra, SQLite, and are working on TimescaleDB support. We continue to believe that making a single syntax and workflow for all database schema definitions will be a key value of the project and gain more adoption.

* **Additional environment support**: Schemahero has added support for configuration from Hashicorp Vault (via a community contribution) and Amazon Web Services parameter store. In addition, Schemahero has added support for publishing a software bill of materials (SBOM) with each release using sigstore. We are working to "meet the end user where they are" and prioritizing work that decreases the effort needed to adopt a Schemahero POC.

* **Additional use cases**: In addition to schema management, we've identified that Schemahero only solves some of the problems solved by traditional database schema management tools. Until schemahero can replace all of the functionality of traditional solutions, it's difficult to get an organization to convert. We've started working on the primary missing use cases: support for fixtures and seed data, and deploying migrations that need to execute code or other user-defined migration processes (data migrations).

* **Community and Awareness**: This is an opportunity for the project to do better. We've been focusing on building functionality and some maturity to the project, and must start to be more visible in the community to get feedback. In 2021, the project joined the CNCF BugBash, had a desk at the CNCF Project Pavilion (LA), and participated in live office hours during Kubecon. We've also started hosting live and recorded community calls monthly on zoom, and have had community participation in these.

> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

The current goals of the Schemahero have changed from our original. We currently have three goals:

1. **Integration**: There are other, popular, non-CNCF tools that are popular database schema management projects and products. We've been evaluating integrating these projects and products (OSS and proprietary) into Schemahero so that an organization can adopt Schemahero alongside an existing tool, not just as a replacement for existing tools.

2. **Adoption**: Schemahero is currently being evaluated in several organizations, and our primary goal is to work with these (and future) organizations during the proof of value phase. These early adopters have expressed interest, and are continuing to slowly roll Schemahero into various environments. Feature work (database support and other functionality) will be focused on areas that can increase adoption.

3. **Community**: We aim to continue growing the Schemahero community. Schemahero was an early project when it joined the sandbox, and needed to continue to grow in maturity and find a use case. We will continue to rely on the community for this, and are expecting to spend significant time working with additional committers, maintainers, and end users of the project. We want to have more committers and more maintainers vs more commits.

## CNCF membership

> How can the CNCF help you achieve your upcoming goals?

Schemahero has benefited from inclusion in the CNCF, especially in the areas of awareness and confidence as an early project. Many current evaluations have come from awareness as part of the CNCF. This has also helped to increase confidence (lower risk) for some as the project is still early.

Some areas where we'd like to get continued and additional help are:

* Opportunities to present and discuss Schemahero in order to increase adoptions. This can be webinars, office hours, events, blogs, or other formats.
* Add Schemahero to the Certified Kubernetes Security Specialist (CKS) examination to help build a community of expertise in the project.
* Resources to create case studies that demonstrate the value of Schemahero in both increased velocity and lower risk when managing database schemas.

## Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

Not at this time. It's the goal of the Schemahero project to continue to build in the Sandbox for the next year to increase production-grade adoption and add additional maintaining organizations. Once the project has achieved production status and has moved out of evaluation phases at adopting organizations, we will start planning for incubation.

## Project Links
 - [Website](https://schemahero.io)
 - [GitHub](https://github.com/schemahero)
 - [Slack](https://slack.k8s.io/#schemahero)
 - [Twitter](https://twitter.com/schemahero)
