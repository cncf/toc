# OpenELB 2022 Annual Review

This is a Kubernetes ToC Annual Review for the [OpenELB](https://openelb.io/) project for 2022.

## Table of Contents

- [Background](#background)
- [DevStats](#development-metrics)
- [Contributors and Maintainers](#contributors-and-maintainers)
- [Adoption](#adoption)
- [Project Goals](#project-goals)
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)

## Background

OpenELB is an open-source load balancer implementation designed for exposing the LoadBalancer type of Kubernetes services in bare metal, edge, and virtualization environments.

In cloud-based Kubernetes clusters, Services are usually exposed by using load balancers provided by cloud vendors. However, cloud-based load balancers are unavailable in bare-metal or on-premise environments. OpenELB allows users to create LoadBalancer Services in bare-metal, edge, and virtualization environments for external access, and provides the same user experience as cloud-based load balancers.

## Development metrics

The OpenELB devstats page and dashboards can be found [here](https://openelb.devstats.cncf.io/).

Here are some links worth mentioning:

- [Stars and forks](https://openelb.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now)
- [Issues](https://openelb.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1)
- [PRs](https://openelb.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&var-period=d&var-repogroup_name=All)
- [Commits](https://openelb.devstats.cncf.io/d/2/commits-repository-groups?orgId=1)


## Contributors and Maintainers

There are two maintainers and two contributors actively working on the project (alphabetically):

| Name           | GitHub ID                                           | Role        | Affiliation |
| -------------- | --------------------------------------------------- | ----------- | ----------- |
| Amal Thundiyil | [amal-thundiyil](https://github.com/amal-thundiyil) | Contributor |             |
| Anurag Pathak  | [AnuragThePathak](https://github.com/AnuragThePathak) | Contributor |             |
| Chauncey Jiang | [chaunceyjiang](https://github.com/chaunceyjiang)   | Maintainer  | DaoCloud    |
| Ryken Ren      | [renyunkang](https://github.com/renyunkang)         | Maintainer  | QingCloud   |


## Adoption

OpenELB has been adopted by [many companies](https://github.com/openelb/openelb/blob/master/ADOPTERS.md) all over the world. Belows are just the list of companies that we know have deployed OpenELB in their own clusters:
- baiwang-580
- cvte
- qingcloud
- sbs
- jollychic
- benlai
- yicheng-health
- wisdomworld


## Project Goals & Update
<!-- What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation? -->
Our current goal is to improve product functions, expand users, and discover contributors.

Project Features：
- Refactor openelb so that it can be better used in more complex environments.
- Support for compatibility with more cni plugins.
- Support for more routing protocols.
- Improve openelb console to make user operations more convenient.
- Solve the single point of failure of layer2 mode to achieve failover.
- Support eip to bind different namespaces.
- Improve openelb monitoring metrics and support importing to grafana.

Docs：
- Improve the docs more in "How it works", in addition to "How to use".
- Publish some related technical blogs.

Community：
- keep being contributor friendly and open to discuss new feature requests and responsive to contributors willing to help.


In 2022 we've shipped features like:
- Support user to configure custom container image
- Add some simple monitoring metrics
- Support vip mode
- Support to assign ip by configuring default EIP

## CNCF Membership

In 2022 we participated in the [lfx-mentorship](https://github.com/cncf/mentoring/tree/main/programs/lfx-mentorship/2022/02-Summer#openelb) event and achieved certain results. Two active contributors participated in the development, but then had to suspend maintenance and contributions due to their academic pressure.


## Incubation

OpenELB is not yet ready for incubation. At this stage, the documentation of openelb is relatively missing and relatively incomplete. The community is not very active, and there are fewer users. Therefore, we are actively attracting more open source contributors to join, and promoting it to let more users know and use it.