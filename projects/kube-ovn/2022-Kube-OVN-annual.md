# Kube-OVN 2022Annual Review

This is the annual review for the [Kube-OVN](https://kube-ovn.io) project for 2022.

## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Project perform](#project-perform)
- [Project goals](#project-goals)
- [Incubation readiness](#incubation-readiness)

## Background

Kube-OVN is an open source cloud-native Kubernetes container network orchestration system,
which greatly enhances the functionality, security, manageability and performance of Kubernetes
container network by leveling the mature network functions from tradition SDN to Kubernetes and
bringing a unique value to the Kubernetes ecosystem.

## DevStats

Kube-OVN DevStats dashboard [here](https://kubeovn.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=now-2y&to=now)
shows an increase in contribution activities. Most of the contributions come from maintainers
who work at Alauda and China Telecom. Since joining the CNCF, we have attracted 37 new committers
from companies like ByteDance, Inspur, EasyStack and 99Cloud.

## Maintainers

Kube-OVN currently has 5 core maintainers from Alauda and China Telecom. Since joining the CNCF,
3 contributors become maintainers, two from Alauda and one from China Telecom.

| Maintainer   | GitHub Username                                | Company       |
|--------------|------------------------------------------------|---------------|
| Mengxin Liu  | [@oilbeater](https://github.com/oilbeater)     | Alauda        |
| Zujian Zhang | [@zhangzujian](https://github.com/zhangzujian) | Alauda        |
| Hongzhen Ma  | [@hzma](https://github.com/hongzhen-ma)        | Alauda        |
| Riming Fan   | [@fanriming](https://github.com/fanriming)     | China Telecom |
| Tao Liu      | [@lut777](https://github.com/lut777)           | Alauda        |

We hope to see more diversity in the maintainer team.

## Adoption

We didn't strictly distinguish the adoption before and after joining the CNCF.
Belows are just the list of companies that we know have deployed Kube-OVN in their own 
infrastructure or integrated Kube-OVN as part of their products.

- Alauda
- China Telecom
- ByteDance
- Inspur
- EasyStack
- 360
- Yealink
- ZTO
- CEC
- CMSChina

## Project perform

Since joining the CNCF, Kube-OVN has published 3 major releases and 10 minor releases.

Key features added to the project:
- VPC function enhancement like eip/nat/policy route/load balancer/security group
- Vlan/underlay network refactor and user experience enhancement
- Performance boost that increase the pps and throughput by a large margin
- Enhancement for KubeVirt like live-migration, multiple interface and sr-iov/ovs-dpdk support

Contributors and adaptors are mainly come from China mainland. We expect to attract more participation
outside of China after joining CNCF. However, this not meet our expect, and we need to focus on it this 
year.  

## Project Goals

The goal of Kube-OVN this year is to increase the diversity in both contributors and adoption.
Some key works include:
- More combination with projects in CNCF, like KubeVirt, Cilium, Submariner and MetaLB to provide 
a whole network solution.
- Improve the user experience in installation, upgrade and troubleshooting.
- Provide more document about the user guide, tutorial and internals 
- Host or join more events to attract more contribution

## Incubation readiness

For adoption, there is already a large number of users and the total image download count has 
exceeded 1.6M. However,the maintainers and majority commits are mainly come from two companies, 
we are not sure if it meets the incubation criteria. We hope to enter the incubation process 
this year, and we are preparing for it.
