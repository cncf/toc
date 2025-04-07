# Kyverno 2021 Annual Review

This is a Kubernetes ToC Annual Review for the [Kyverno](https://kyverno.io) project for 2021.

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

Kyverno ("govern" in Greek) is a policy management solution designed for Kubernetes. 

The main goal of the Kyverno project is to simplify Kubernetes configuration management using policies for security and automation. Kyverno makes it easy for Kubernetes administrators to write and manage Kubernetes policies and for Kubernetes workload owners (for example, developers) to consume policy results and address issues.

In addition to validating resources (in either audit or enforce modes), Kyverno has full support for mutating and generating Kubernetes resource and is often used to automate costly and error-prone handoffs across operations, security, and development roles.

Kyverno uses Kubernetes custom resource definitions for cluster-wide and namespaced policies, and the [Kubernetes Policy WG policy report](https://github.com/kubernetes-sigs/wg-policy-prototypes/tree/master/policy-report) for reporting policy results. Kyverno runs as an admission controller and as a standalone command-line tool.

Since Kyverno is focused on Kubernetes, it leverages Kubernetes patterns, idioms, and tools such as:
* owner references
* pod and pod controller relations
* structural schemas in OpenAPIv3 
* strategic merge patches

Kyverno's focus on Kubernetes and its subsequent ability to leverage Kubernetes internals makes it intuitive and easy to use for Kubernetes administrators and operators.

Kyverno was accepted as a CNCF Sandbox project on [November 9th, 2020](https://docs.google.com/spreadsheets/d/1Nnh_usr0tSZxaUpxTusqeIqKxMmvuEViRkyO9e_Do40/edit#gid=1136111842).

## DevStats

> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

The Kyverno DevStats dashboards can be found [here](https://kyverno.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&viewPanel=4).

We are using the following metrics as key indicators of community health:
* [Stargazers](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Stargazers&var-table=swatchers&var-pref=all&var-met1=watch&var-met2=watch&from=now-1y&to=now): 331% annual growth from 524 -> 1735.
* [Issue Creators](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Issue%20creators&var-table=scommunity_health&var-pref=&var-met1=chealthissue&var-met2=&from=now-1y&to=now): 500% annual growth from 45 -> 225.
* [Code Committers](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Code%20committers&var-table=scommunity_health&var-pref=&var-met1=chealthcommit&var-met2=&from=now-1y&to=now): 500% annual growth from 18 -> 90.
* [Code Commenters](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Code%20commenters&var-table=scommunity_health&var-pref=&var-met1=chealthcomment&var-met2=&from=now-1y&to=now): 509% annual growth from 22 -> 112.

## Maintainers

Kyverno currently has 5 maintainers from 4 different companies:

| Maintainer           | GitHub ID                                     | Affiliation |
| -------------------- | --------------------------------------------- | ----------- |
| Chip Zoller           | [@chipzoller](https://github.com/chipzoller)   | Dell Technologies  |
| Jim Bugwadia          | [@JimBugwadia](https://github.com/JimBugwadia)   | Nirmata Inc.  |
| Marcel Muller        | [@MarcelMue](https://github.com/MarcelMue)    | Giant Swarm GmbH |
| Shuting Zhao          | [@realshuting](https://github.com/realshuting)   | Nirmata Inc.   |
| Trey Dockendorf       | [@treydock](https://github.com/treydock)      | Ohio Supercomputer Center |
| Vyankatesh Kudtarkar	| [@vyankyGH](https://github.com/vyankyGH)	| Nirmata Inc.  |
| Prateek Pandey        | [@prateekpandey14](https://github.com/prateekpandey14) | Nirmata                   |
| Sambhav Kothari       | [@samj1912](https://github.com/samj1912)               | Bloomberg                 |


The maintainer list is managed in our [GitHub repository](https://github.com/kyverno/kyverno/blob/main/MAINTAINERS.md) along with our [governance policy](https://kyverno.io/community/).

## Adoption

Adopters who are publicly referenceable are listed in the [Kyverno GitHub repository ADOPTERS.md file](https://github.com/kyverno/kyverno/blob/main/ADOPTERS.md).

Kyverno adopters include:
 
**Projects**
* [Flux](https://github.com/fluxcd/flux2-multi-tenancy/tree/main/infrastructure/kyverno-policies) 
* [OpenEBS](https://github.com/openebs/charts/tree/main/charts/openebs/templates/kyverno)
* [cert-manager](https://github.com/jetstack/cert-manager/tree/master/devel/addon/kyverno)
* [CNF WG tests](https://github.com/cncf/cnf-testsuite)

**Organizations**
* [Ohio Supercomputer Center](https://www.osc.edu/)
* [VSHN](https://www.vshn.ch/en/)
* [Coinbase](https://www.coinbase.com/)
* [Mandiant](https://www.mandiant.com/)
* [Giant Swarm](https://www.giantswarm.io/)
* [Vodafone Group Plc](https://www.vodafone.com/)
* [Deutsche Telekom](https://www.telekom.com/en)
* [Bloomberg](https://www.techatbloomberg.com/) 

A list of companies providing commercial services and solutions for Kyverno is available [here](https://kyverno.io/support/).

## Project Goals

> How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

The main goals of the project have been in the following categories:

* **Complex policy support**: Several new features such as [JMESPath support](https://kyverno.io/docs/writing-policies/variables/), [API Server Lookups](https://kyverno.io/docs/writing-policies/external-data-sources/#variables-from-kubernetes-api-server-calls), and [foreach](https://kyverno.io/docs/writing-policies/validate/#foreach) have been introduced to allow declarative processing of complex policy logic.

* **Availability**: Kyverno has added the ability to support multiple replicas for fault-tolerance and increased availability in Release 1.4.0 (June 2021).

* **Scalability**: Several fixes and improvements have been made to reduce memory usage and handle large clusters. This remains an on-going area of test and improvement.

* **Security**: Kyverno participated in the CNCF TAG Security [Security Pals](https://github.com/cncf/tag-security/issues/554) program. As a result, Kyverno has adopted recommended security best practices. The [Security](https://main.kyverno.io/docs/security/) section in the Kyverno documentation lists these and acts as a security guide for adopters.

* **Community and Awareness**: Kyverno has grown from initially having maintainers from a single company, to maintainers from multiple organizations. The community has also grown significantly as reflected by [DevStats](https://kyverno.devstats.cncf.io). While Kyverno awareness has grown, more work is needed to make the community fully aware of all of its features and capabilities.

> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

1. **Scalability**: Kyverno is increasingly being used in large production clusters. We aim to dramatically improve scalability in the subsequent releases.

2. **Security**: In addition to current security processes, a goal for Kyverno is to improve software supply chain security and achieve [SLSA Level 4](https://slsa.dev/levels).

3. **Adoption**: We aim to continue growing Kyverno adoption. A survey conducted by Nirmata at KubeCon US 2021 showed Kyverno adoption is currently around 8%. Our goal is to exceed 30% adoption by Q4 2022.

## CNCF membership

> How can the CNCF help you achieve your upcoming goals?

Kyverno has benefited greatly from the CNCF and its community, especially in the areas of awareness and adoption. The Kyverno project has also participated in the [LFX Mentorship program](https://mentorship.lfx.linuxfoundation.org/#projects_all), KubeCon project office hours, and the Community Infrastructure Lab.

Some areas where we can use help are:
 
* Best practices for automated scale testing.
* A security assessment and guidance on fuzzing and other security best practices to improve the Kyverno security posture. 
* Continued support for webinar, blogs, and other community activities to increase visibility and adoption.
* Add Kyverno to the Certified Kubernetes Security Specialist (CKS) examination and any other CNCF materials that reference the Open Policy Agent (OPA) project, to help drive awareness and promote fairness across projects.

## Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

Yes. The Kyverno project team has started planning a proposal for incubation and plans to submit it in the next month. This work is tracked in [issue #2838](https://github.com/kyverno/kyverno/issues/2838).

## Project Links
 - [Website](https://kyverno.io)
 - [Github](https://github.com/kyverno)
 - [Slack](https://slack.k8s.io/#kyverno)
 - [Twitter](https://twitter.com/kyverno)
 
