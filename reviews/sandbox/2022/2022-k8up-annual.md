# K8up Annual Review 2022

This is the very first annual review of the [k8up](https://github.com/k8up-io/) project.

## Background

K8up is a Kubernetes backup operator based on Restic that will handle PVC and application backups on a Kubernetes or OpenShift cluster.

## DevStats

The [K8up DevStats Dashboards](https://k8up.devstats.cncf.io/d/8/dashboards?orgId=1&from=1622498400000&to=1685656799000) shows (from the past two years):

* 6 companies contributed to the project.
* 18 different contributing developers.

Around 11 releases where made, including one breaking change from 1.x to 2.x.

## Maintainers

The K8up maintainers are tracked in the [GitHub Team "Maintainer"](https://github.com/orgs/k8up-io/teams/maintainer/members), as described in our [OWNERS.md](https://github.com/k8up-io/k8up/blob/master/OWNERS.md) file.

Current Maintainer(s):

| Maintainer       | GitHub Username                                   | Company |
| ---------------- | ------------------------------------------------- | ------- |
| Tobias Brunner   | [@tobru](https://github.com/tobru)                | VSHN AG |
| Simon Beck       | [@Kidswiss](https://github.com/Kidswiss)          | VSHN AG |
| Nicolas Bigler   | [@TheBigLee](https://github.com/TheBigLee)        | VSHN AG |
| Fabian Fischer   | [@glrf](https://github.com/glrf)                  | VSHN AG |
| Lukasz Widera    | [@wejdross](https://github.com/wejdross)          | VSHN AG |
| Gabriel Saratura | [@zugao](https://github.com/zugao)                | VSHN AG |
| Liene Luksika    | [@lieneluksika ](https://github.com/lieneluksika) | VSHN AG |

Our wish is to onboard maintainers from outside VSHN AG and hope to find some in the upcoming months.

## Adoption

Adoption of K8up is tracked in our [ADOPTERS.md](https://github.com/k8up-io/k8up/blob/master/ADOPTERS.md) file.
As of 2023-06-01 we have three adopters documented:

* [VSHN](https://www.vshn.ch)
* [amazee.io](https://www.amazee.io)
* [Lagoon](https://github.com/uselagoon)

## Project Goals

### Goals Since Last Review

This is K8up's first review, so there are no past goals to report on.

### Current Project Goals

Project health:
* Onboard more maintainers from other companies than VSHN AG
* Increase adoption to gain traction for the project
* Progress K8up's stability and reliability to be rock-solid and a fully reliable data protection tool
* Fix annoying bugs which are being discovered while the adoption rises

Features and technology:
* Renew operator patterns to be on-par with current best practices (e.g. replace controller configuration via environment variables with a configuration CRD)
* Use native Kubernetes CronJobs for scheduling backup jobs, rather than using our home-grown solution
* Refresh Restic repository configuration
* Further enhance the newly built CLI tool
* Support multiple backups per namespace

## Help from CNCF for the project

The CNCF membership has helped the project already tremendously since it's adoption into it.
We've seen increased visibility and contribution from various people and organizations.
Being present on CNCF events like KubeCon EU gave the project a real boost.

## Incubation

Currently, K8up remains a sandbox project. It is not ready to apply for incubation yet, as it hasn't seen broad adoption as of today.
We also would like to have more maintainers outside VSHN AG and are happy to welcome them.
