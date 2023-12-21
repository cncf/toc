# Tinkerbell 2022 Annual Review

This is the annual review for the [Tinkerbell](https://tinkerbell.org) project for 2023.

## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Project goals](#project-goals)
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)

## Background

Tinkerbell is a cloud native bare metal provisioning engine. It enables provisioning machines through the same concepts that are used in cloud providers. It provides APIs for integration with other systems.

The Tinkerbell stack consists of the following services. Most services are a separate repository in the Tinkerbell GitHub organization.

- [Tink Server](https://github.com/tinkerbell/tink) is responsible for serving tasks to be run by Tink Worker and updating the state of tasks as reported by Tink worker.
Tink Worker is responsible for retrieving and executing workflow tasks. It reports the state of tasks back to Tink Server. It is pre-packaged into our default in-memory provisioning OS, Hook.

- [Tink Worker](https://github.com/tinkerbell/tink) is responsible for retrieving and executing workflow tasks. It reports the state of tasks back to Tink Server.

- [Tink Controller](https://github.com/tinkerbell/tink) is responsible for rendering workflow templates and managing workflow state as Tink Worker's report on their task status'. It is an internal component that users generally do not need to interact with.

- [Boots](https://github.com/tinkerbell/boots) is Tinkerbell's DHCP server. It handles DHCP requests, hands out IPs, and serves up iPXE. It uses the Tinkerbell client to pull and push hardware data. It only responds to a predefined set of MAC addresses so it can be deployed in an existing network without interfering with existing DHCP infrastructure.

- [Hook](https://github.com/tinkerbell/hook) is Tinkerbell's default in-memory installation environment for bare metal. Hook executes workflow tasks that result in a provisioned machine.
Optional services

- [Hegel](https://github.com/tinkerbell/hegel) is a metadata service that can be used during the configuration of a permanent OS. This enables native integration with cloud-init and other configuration management tools.

- [PBnJ](https://github.com/tinkerbell/pbnj) is a microservice that can communicate with baseboard management controllers (BMCs) to control power and boot settings.

- [Rufio](https://github.com/tinkerbell/rufio) is a Kubernetes controller that facilitates baseboard management controller interactions. It functions similarly to PBnJ but with a Kubernetes focused API.

- [Cluster API Provider Tinkerbell](https://github.com/tinkerbell/cluster-api-provider-tinkerbell) is the CAPI provider for Tinkerbell. It enables provisioning Kubernetes clusters on bare metal using Tinkerbell.

## DevStats

Please find our devstats page [here](https://tinkerbell.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

2022 was an exciting year for the Tinkerbell project. Tinkerbell had multiple releases across the numerous underlying services that make up the Tinkerbell stack. We migrated our backend datastore from Postgres to Kubernetes and published Helm charts. We created and integrated the Rufio service into our CAPI Provider. We staffed our first project pavilion kiosk and had a [talk about Tinkerbell](https://kccncna2022.sched.com/event/182In/bare-metal-chronicles-intertwinement-of-tinkerbell-cluster-api-and-gitops-katie-gamanji-apple) presented at KubeCon NA in Detroit. We created our [official governance doc](https://github.com/tinkerbell/org/blob/main/GOVERNANCE.md) and onboarded individuals to all the associated roles for each of our services. We also created an official [roadmap](https://github.com/tinkerbell/roadmap) repository and began tracking our high level goals. We held community calls almost every week and helped lots of folks getting started in our [slack channel](https://cloud-native.slack.com/archives/C01SRB41GMT).

## Maintainers

The following individuals were Tinkerbell [maintainers](https://github.com/tinkerbell/org/blob/main/GOVERNANCE.md#maintainer) as of 2022-12-31. We have [CODEOWNER](https://github.com/tinkerbell/pbnj/blob/main/.github/CODEOWNERS) files in each of our repositories that list the maintainers for that repository. Linking to each of them would be a burden for all but this [issue](https://github.com/tinkerbell/org/issues/10) was used to track role invitations and can be used to see the full list of maintainers for 2022.

| Maintainer        | GitHub Username                                      | Company       |
| ----------------- | ---------------------------------------------------- | ------------- |
| Manuel Mendez     | [@mmlb](https://github.com/mmlb)                     | Equinix Metal |
| Micah Hausler     | [@micahhausler](https://github.com/micahhausler)     | AWS           |
| Jacob Weinstock   | [@jacobweinstock](https://github.com/jacobweinstock) | AWS           |
| Nahum Shalman     | [@nshalman](https://github.com/nshalman)             | Equinix Metal |
| Chris Doherty     | [@chrisdoherty4](https://github.com/chrisdoherty4)   | AWS           |
| Thomas Stromberg  | [@tstromberg](https://github.com/tstromberg)         | Chainguard    |
| Daniel Finneran   | [@thebsdbox](https://github.com/thebsdbox)           | Isovalent     |
| Jason DeTiberus   | [@detiber](https://github.com/detiber)               | Cisco         |
| Marques Johansson | [@displague](https://github.com/displague)           | Equinix Metal |
| Jeremy Tanner     | [@jeremytanner](https://github.com/jeremytanner)     | Tailscale     |
| Mara              | [@DailyAlice](https://github.com/DailyAlice)         | Equinix Metal |

## Adoption

We do not have an official ADOPTERS file. We have been informed by individuals from the following companies that they have met the CNCF [definition of an adopter](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter).

- AWS - [EKS Anywhere](https://anywhere.eks.amazonaws.com/docs/getting-started/baremetal/overview/)
- Equinix Metal - [Bare metal cloud](https://deploy.equinix.com/)
- Kubermatic GmbH

There is also one known open source project that has integrated Tinkerbell components into its product; [Seeder](https://github.com/harvester/seeder), which is a part of the Harvester project from Rancher.

## Project goals

As this is the first review for Tinkerbell, we don't have a previous set of goals to review. The projects current goals are to (1) improve the user experience around getting started with Tinkerbell, (2) update the [docs](https://docs.tinkerbell.org/) to catch up with the current state of the project, and (3) roll out the v1alpha2 API across all services. Details on these goals and the work in progress can be found in our [roadmap repository](https://github.com/orgs/tinkerbell/projects/16/views/1).

## CNCF Membership

We don't have any requests from the CNCF at this time.  

## Incubation

The Tinkerbell project is not ready for to move to incubation. With that said, we are working towards that goal. Increased sustained committers and documenting adopters are the two areas we are currently focusing on.
