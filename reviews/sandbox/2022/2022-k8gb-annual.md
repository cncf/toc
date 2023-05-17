# k8gb 2022 Annual Review

This is a Kubernetes ToC Annual Review for the [k8gb](https://k8gb.io) project for 2022.

<!-- omit in toc -->
## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Latest Achievements](#latest-achievements)
- [Project Goals](#project-goals)
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)
- [Project Links](#project-links)

## Background

[k8gb](https://www.k8gb.io/) is an open source DNS-based global load balancing solution for Kubernetes.

k8gb focuses on traffic load balancing across geographically dispersed Kubernetes clusters using different load balancing strategies to achieve high availability.
k8gb works on Layer 7. It registers as a resolver in the Edge DNS server(e.g. Infoblox, Route53, NS1) and provides DNS resolution for the services based on the given strategy and the Kubernetes workloads' health status.
k8gb implements Kubernetes operator pattern with custom controller configurable either via standard Ingress resources or Gslb custom resources.

k8gb also relies on well-known projects from CNCF and Kubernetes ecosystems in its work:

- [CoreDNS](https://coredns.io/)
- [ExternalDNS](https://github.com/kubernetes-sigs/external-dns)

More information about k8gb architecture can be found [here](https://www.k8gb.io/docs/components.html)

k8gb was accepted as a CNCF Sandbox project on [March 30th, 2021](https://docs.google.com/spreadsheets/d/1Nnh_usr0tSZxaUpxTusqeIqKxMmvuEViRkyO9e_Do40/edit#gid=1136111842).

## DevStats

[k8gb's DevStats page](https://k8gb.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1609520400000&to=1641056399000&var-period=w&var-repogroup_name=All) shows the weekly number of
newly opened pull requests on the main git repo.

| Metric | Pre-Sandbox | Today |
| ------ | ----------- | ----- |
| Stars  | ~100        | ~450  |
| Forks  | 10-20       | ~50   |

## Maintainers

There are six maintainers working on the project (alphabetically):

| Maintainer      | GitHub ID                                         | Affiliation |
| --------------- | ------------------------------------------------- | ----------- |
| Dinar Valev     | [k0da](https://github.com/k0da)                   | Absa        |
| Donovan Muller  | [donovanmuller](https://github.com/donovanmuller) | Absa        |
| Jiri Kremser    | [jkremser](https://github.com/jkremser)           | Absa        |
| Michal Kuritka  | [kuritka](https://github.com/kuritka)             | Absa        |
| Timofey Ilinykh | [somaritane](https://github.com/somaritane)       | Absa        |
| Yury Tsarev     | [ytsarev](https://github.com/ytsarev)             | Upbound     |

## Adoption

k8gb is widely used in production by the [Absa Group](https://www.absa.africa/absaafrica/), where it was originally developed.
There are several other companies trying to adopt k8gb, but we don't have a public confirmation available yet.

## Latest Achievements

- Support for bigger multi-cluster deployments, not limited to two clusters
- Better observability (more metrics, convenience tools for metrics showcasing)
- Integration with [Operator Lifecycle Manager](https://olm.operatorframework.io/) and [operatorhub.io](https://operatorhub.io/)
- Integration with [Artifact Hub](https://artifacthub.io/packages/search?ts_query_web=k8gb&sort=relevance&page=1)
- Various Quality improvements: FakeDNS, more terratests, additional SAST checks, real DNS server (`bind`) in tests
- Ingress API version upgrade (`v1beta1 -> v1`), support for latest k8s versions
- Spreading the word about k8gb in the community ([KubeCon, FOSDEM, DataOnKubernetes](https://www.youtube.com/playlist?list=PL9-JmDqpp9x_hAYS4zlhnDa5PTJPqmBnV))
- Moving to [open transparent project meetings](https://www.youtube.com/playlist?list=PL9-JmDqpp9x8mhAyK8DpCixGLOHmZrqbZ)
- Last but not least, k8gb is honored to be mentioned in the "[Kubernetes – An Enterprise Guide - Second Edition](https://www.packtpub.com/product/kubernetes-an-enterprise-guide-second-edition/9781803230030)" book by By Marc Boorshtein and Scott Surovich

## Project Goals

<!-- What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation? -->

Right now the project is focused on wider adoption and better visibility in the community.
We also plan to expand the integration use cases to other managed cluster environments, such as Azure, GCP.
Historically the project has proven to work on vanilla on-prem Kubernetes clusters and AWS/EKS and should work in other environments as well, but we need external help to make sure it works in these environments.

## CNCF Membership

<!-- How can the CNCF help you achieve your upcoming goals? -->

We hope that CNCF membership will help us to reach a wider audience and get more contributors involved in the project.
One of the possibilities is CNCF-sponsored events.
We also would like to explore the opportunities to expand the integration cases in cloud managed k8s environments, such as AWS, Azure, GCP, DigitalOcean, etc.

## Incubation

Currently, we only have one publicly confirmed adoption case, so we do not meet all the adoption criteria for the [incubating](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage) stage yet.
The goal is to continue gathering the community around the project and get publicly available adopters.

## Project Links

- [Website](https://k8gb.io/)
- [Github](https://github.com/k8gb-io/k8gb)
- [Weekly Meeting Minutes](https://docs.google.com/document/d/1YdpEVhtyCKvwFtXR7cn1Kn2Xc_tdNskoFnhHmFPbtA4/edit#heading=h.w7l85jeih9j6)
- [Project Calendar](https://calendar.google.com/calendar/u/0/embed?src=ena2n7rlpb6sj14llroc274uvs@group.calendar.google.com)
- [Youtube Channel](https://www.youtube.com/channel/UCwvtktvdZu_pg-t-INvuW5g)
