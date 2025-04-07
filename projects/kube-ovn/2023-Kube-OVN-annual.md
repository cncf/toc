# Kube-OVN 2023 Annual Review

This is the annual review for the [Kube-OVN](https://kube-ovn.io) project for 2023.

## Table of Contents

- [Kube-OVN 2023 Annual Review](#kube-ovn-2023-annual-review)
  - [Table of Contents](#table-of-contents)
  - [Background](#background)
  - [DevStats](#devstats)
  - [Maintainers](#maintainers)
  - [Adoption](#adoption)
  - [Project perform](#project-perform)
  - [Project Goals](#project-goals)
  - [Incubation readiness](#incubation-readiness)

## Background

Kube-OVN is an open source cloud-native Kubernetes container network orchestration system,
which greatly enhances the functionality, security, manageability and performance of Kubernetes
container network by leveling the mature network functions from tradition SDN to Kubernetes and
bringing a unique value to the Kubernetes ecosystem.

## DevStats

Kube-OVN DevStats dashboard [here](https://kubeovn.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=now-5y&to=now)
shows an increase in contribution activities. Last year [93 developers](https://kubeovn.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-companies=All) contribute to Kube-OVN.

## Maintainers

2 contributors became new maintainers these year and one retired due to work change.

| Maintainer     | GitHub Username                                | Company       |
| -------------- | ---------------------------------------------- | ------------- |
| Mengxin Liu    | [@oilbeater](https://github.com/oilbeater)     | Alauda Inc    |
| Zujian Zhang   | [@zhangzujian](https://github.com/zhangzujian) | Alauda Inc    |
| Hongzhen Ma    | [@hzma](https://github.com/hongzhen-ma)        | Alauda Inc    |
| Riming Fan     | [@fanriming](https://github.com/fanriming)     | China Telecom |
| Bingbing Zhang | [@bobz965](https://github.com/bobz965)         | Yealink Inc   |
| YuDong Wang    | [@wangyd1988](https://github.com/wangyd1988)   | Inspur        |

We hope to see more diversity in the maintainer team.

## Adoption

Belows are the NEW companies this year that we know have deployed Kube-OVN in their own 
infrastructure or integrated Kube-OVN as part of their products.

- 99cloud
- Canonical with Charmed Kubernetes
- China Mobile
- QingCloud Inc with KSV

## Project perform

This year, we are continuing to add new features to Kube-OVN, as well as making some fundamental refactors to facilitate future development.

Some key features and refactors added to the project:
- A brand new Document website with all documents rewrite.
- More underlay network features and optimization like overlay/underlay interconnection, better performance and more resilient.
- Refactor the Virtual Private Cloud (VPC) Network Address Translation (NAT) Gateway implementation to gain greater flexibility.
- More VPC network functions like DHCP, DNS, and LoadBalance.
- Use libovsdb to replace all the connection to OVN to achieve better performance.  

## Project Goals

The goal of Kube-OVN this year is to increase the diversity in both contributors and adoption.
Some key works include:
- More timely update about Kube-OVN new features.
- Accelerate the release cycle.
- Better test coverage.

## Incubation readiness

Yes, we have submitted the [proposal](https://github.com/cncf/toc/pull/910) for incubation.
