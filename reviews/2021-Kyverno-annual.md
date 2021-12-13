# Kyverno 2021 Annual Review

This is a Kubernetes ToC Annual Review for the [Kyverno](https://kyverno.io) project.

## Table of Contents
- [Background](#background)
- [DevStats](#development-metrics)
- [Maintainers](#maintainers)
- [Adoption](#adoption) 
- [Project goals](#project-goals) 
- [CNCF Membership](#annual-review-contents)
- [Incubation](#incubation)
- [Project Links](#project-links)

## Background

Kyverno (which means "govern" in Greek) is a policy management solution designed for Kubernetes. 

The main goal for the Kyverno project is to simplify Kubernetes configuration management using policies for security and automation. Kyverno making it easy for Kubernetes administrators to write and manage Kubernetes policies, and for Kubernetes users (i.e. developers) to consume policy results and address issues. 

Kyverno has full support for mutating and generating Kubernetes resources, and hence can be used to automate costly and error-prone handoffs across operations, security, and development roles.

Kyverno uses Kubernetes custom resources for policies and policy reporting. Kyverno runs as an admission controller and a command-line tool.

Since Kyverno is focused on Kubernetes, it leverages Kubernetes patterns, idioms, and tools such as:
* owner references
* pod and pod controller relations
* structural schemas in OpenAPIv3 
* strategic merge patches

This focus on Kubetnetes and the subsequent ability to leverage Kubernetes internals makes Kyverno intuitive and easier to use for Kubernetes administrators and users.

Kyverno was accepted as a CNCF Sandbox project on [November 9th, 2020](https://docs.google.com/spreadsheets/d/1Nnh_usr0tSZxaUpxTusqeIqKxMmvuEViRkyO9e_Do40/edit#gid=1136111842).

## DevStats

> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

The Kyverno DevStats dashboards can be found [here](https://kyverno.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&viewPanel=4).

We are using the following metrics as key indicators of community health:
* [Stargazers](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Stargazers&var-table=swatchers&var-pref=all&var-met1=watch&var-met2=watch&from=now-1y&to=now): 331% annual growth from 524 -> 1735.
* [Issue Creators](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Issue%20creators&var-table=scommunity_health&var-pref=&var-met1=chealthissue&var-met2=&from=now-1y&to=now): 500% annual growth from 45 -> 225.
* [Code Committers](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Code%20committers&var-table=scommunity_health&var-pref=&var-met1=chealthcommit&var-met2=&from=now-1y&to=now): 500% annual growth from 18 -> 90.
* [Code Commenters](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Code%20commenters&var-table=scommunity_health&var-pref=&var-met1=chealthcomment&var-met2=&from=now-1y&to=now): 509% annual growth from 22 -> 112.

## Maintainers (Anita)

Kyverno currently has 5 maintainers from 4 different companies:

| Maintainer           | GitHub ID                                     | Affiliation |
| -------------------- | --------------------------------------------- | ----------- |
| Jim Bugwadia          | [@JimBugwadia](https://github.com/JimBugwadia)   | Nirmata Inc.  |
| Shuting Zhao          | [@realshuting](https://github.com/realshuting)   | Nirmata Inc.   |
| Chip Zoller           | [@ChipZoller](https://github.com/chipzoller)   | Dell  |
| Marcel Mueller        | [@MarcelMue](https://github.com/MarcelMue)    | Giant Swarm GmbH |
| Trey Dockendorf       | [@treydock](https://github.com/treydock)      | Ohio Supercomputer Center |

The full maintainer list can be found in our [GitHub repository](https://github.com/kyverno/blob/main/MAINTAINERS.md) along with our [governance policy](https://kyverno.io/community/).

## Adoption

We list public adopters on the [Kyverno GitHub repository](https://github.com/kyverno/kyverno/blob/main/ADOPTERS.md). Some of these are:
 
**Projects**
* [Flux](add link to Kyverno use) 
* [OpenEBS](add link to Kyverno use)
* [cert-manager](add link to Kyverno use)
* [CNF WG tests](add link to Kyverno use)

**Organizations**
* [Ohio Supercomputer Center](https://www.osc.edu/)
* [VSHN](https://www.vshn.ch/en/)
- Coinbase
- Mandiant
- Giant Swarm
- Vodafone Group Plc
- Deutsche Telekom

There are a number of other adopters not listed here, who are privately referencable.

A list of companies providing commercial services and solutions for Kyverno is available [here](https://kyverno.io/support/).

## Project goals (Jim)

> How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

The main goals of the project have been in the following categories:

* **Complex policy support**: Several new features such as [JMESPath support](https://kyverno.io/docs/writing-policies/variables/), [API Server Lookups](https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-calls), and [foreach](https://kyverno.io/docs/writing-policies/validate/#foreach) have been introduced to allow declarative processing of complex policy logic.

* **Availibility**: Kyverno has added the ability to support multiple replicas for fault-tolerance and increased availability in Release 1.4.0 (June 2021).

* **Scalability**: Several fixes and improvenents have been made to reduce memory usage and handle large clusters. This remains an on-going area of test and improvement.

* **Security**: Kyverno participated in the CNCF TAG Security [Security Pals](https://github.com/cncf/tag-security/issues/554) program. As a result, Kyverno has adopted recommended security best practices. The [Security](https://main.kyverno.io/docs/security/) section in the Kyverno documentation lists these and acts as a security guide for adopters.

* **Community and Awareness**: Kyverno has grown from a initially having maintainers from a single community, to maintainers from multiple organizations. The community has also grown significantly as reflected by [DevStats](https://kyverno.devstats.cncf.io). While Kyverno awareness has grown, more work is needed to make the community fully aware of all of its features and capabilities.

> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

1. **Scalability**: Kyverno is increasingly being used in large production clusters. We aim to dramatically improve scalability in the subsequent releases.

2. **Security**: In addition to current security processes, a goal for Kyverno is to achieve [SLSA Level 4](https://slsa.dev/levels) in 2022.

3. **Adoption**: We aim to continue growing Kyverno adoption. A survey conducted by Nirmata at KubeCon US 2021 showed Kyverno adoption is currently around 8%. Our goal is to exceed 30% adoption by Q4 2022.

## CNCF membership

> How can the CNCF help you achieve your upcoming goals?

Kyverno has benefited greatly from the CNCF and its community, especially in the areas of awareness and adoption. The Kyverno project has also participated in the [LFX Mentorship program](https://mentorship.lfx.linuxfoundation.org/#projects_all), KubeCon project office hours, and the Community Infrastructure Lab.

Some areas where we can use help are:
 
* Best practices for automated scale testing.
* A security assessment and guidance on fuzzing and other security best practices to improve the Kyverno security posture. 
* Continued support for webinar, blogs, and other community activities to increase visibility and adoption.

## Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)?

Yes. The Kyverno project has started planning a proposal for incubation and plan to submit this in the next month. This work is tracked at [github issue](...).

## Project Links
 - [Website](https://kyverno.io)
 - [Github](https://github.com/kyverno)
 - [Slack](https://slack.k8s.io/#kyverno)
 - [Twitter](https://twitter.com/kyverno)



