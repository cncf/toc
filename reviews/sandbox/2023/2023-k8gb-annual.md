# k8gb 2023 Annual Review

This is a Kubernetes ToC Annual Review for the [k8gb](https://k8gb.io) project for 2023.

<!-- omit in toc -->
## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Latest Achievements](#latest-achievements)
- [Project Goals](#project-goals)
- [Performance against the goals since last
  review](#performance-against-the-goals-since-the-last-review)
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

[k8gb's DevStats
page](https://k8gb.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1641056400000&to=1692806399000&var-period=w&var-repogroup_name=All)
shows the weekly number of newly opened pull requests from the beginning of 2022
to the date of this review report(2023-08-23)

We are visibly less heavy on feature PRs this year but good in project maintenance and adoption.

| Metric | Pre-Sandbox | 2022  | 2023 |
| ------ | ----------- | ----- | ---- |
| Stars  | ~100        | ~450  | ~630 |
| Forks  | 10-20       | ~50   | ~70  |

## Maintainers

There are five maintainers working on the project (alphabetically):

| Maintainer      | GitHub ID                                         | Affiliation |
| --------------- | ------------------------------------------------- | ----------- |
| Dinar Valev     | [k0da](https://github.com/k0da)                   | Absa        |
| Donovan Muller  | [donovanmuller](https://github.com/donovanmuller) | Absa        |
| Jiri Kremser    | [jkremser](https://github.com/jkremser)           | Giant Swarm |
| Michal Kuritka  | [kuritka](https://github.com/kuritka)             | Absa        |
| Yury Tsarev     | [ytsarev](https://github.com/ytsarev)             | Upbound     |

## Adoption

k8gb is widely used in production by the [Absa Group](https://www.absa.africa/absaafrica/), where it was originally developed.

[Millennium bcp](https://www.millenniumbcp.pt/) publicly announced the k8gb adoption in 2023.

There are several other companies trying to adopt k8gb, but we don't have a public confirmation available yet.

We started to collect public adopters at https://github.com/k8gb-io/k8gb/blob/master/ADOPTERS.md

## Latest Achievements

- k8gb project was featured in
  https://www.cncf.io/reports/security-slam-north-america-2022/ where we
implemented advanced SLSA-enabled release pipelines and other security
enhancements
- Early stage of Azure support.
- KubeCon EU 2023 [presentation](https://www.youtube.com/watch?v=U46hlF0Z3xs)
  by the project creator and the project adopter/production user.
  We got really positive feedback after the session.
- Production-related fix of [case-insensitive DNS
  queries](https://github.com/k8gb-io/coredns-crd-plugin/pull/54)
- KCD Bengaluru k8gb [presentation](https://www.youtube.com/watch?v=vrDCUIVyc4g)
  by a community member.

## Project Goals

<!-- What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation? -->

- Adoption and visibility
- Stabilize Azure support
- Extend to GCP
- Increase EdgeDNS provider coverage
- Migrate to the new advanced architecture(https://github.com/k8gb-io/k8gb-lite)
- Support more source objects in addition to Ingress

## Performance against the goals since the last review

<!-- How has the project performed against its goals since the last review? -->

We believe that the project performed reasonably well with work on new
hyperscaler support(Azure), getting new public adopter(Millennium bcp), 2 public
presentations at well-attended events and
overall continuous maintenance of the project.

We also followed the TOC
[recommendations](https://github.com/cncf/toc/pull/837#issuecomment-1164576102)

* Maintainer's group is now more diverse from an organization membership perspective(3 Absa,
  1 Giant Swarm, 1 Upbound)
* Created public [roadmap](https://github.com/orgs/k8gb-io/projects/2/views/2)
* Actively worked on new environment (Azure) support
* Created https://github.com/k8gb-io/k8gb/blob/master/ADOPTERS.md

While the project is actively maintained from an adoption and security perspective
we are objectively concerned with the decrease in feature-related PRs.

We plan to ease the entry level for contribution by moving to
https://github.com/k8gb-io/k8gb-lite architecture in the future, where it will be
much easier to extend the project.

## CNCF Membership

<!-- How can the CNCF help you achieve your upcoming goals? -->

We hope that CNCF membership will help us to reach a wider audience and get more contributors involved in the project.
One of the possibilities is CNCF-sponsored events.
We are pretty confident that k8gb has more adopters out there - it would be great
if CNCF would somehow globally advertise the public project usage announcement in general and motivate
companies to add themselves to https://github.com/k8gb-io/k8gb/blob/master/ADOPTERS.md

## Incubation

<!-- Do you think that your project meets the criteria for incubation? -->

Currently, we have 3 [publicly confirmed](https://github.com/k8gb-io/k8gb/blob/master/ADOPTERS.md) adoption cases(+2 in 2023!).

According to the [incubating criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage), we **do** meet all requirements for the Incubation.

## Project Links

- [Website](https://k8gb.io/)
- [Github](https://github.com/k8gb-io/k8gb)
- [Weekly Meeting Minutes](https://docs.google.com/document/d/1YdpEVhtyCKvwFtXR7cn1Kn2Xc_tdNskoFnhHmFPbtA4/edit#heading=h.w7l85jeih9j6)
- [Project Calendar](https://calendar.google.com/calendar/u/0/embed?src=ena2n7rlpb6sj14llroc274uvs@group.calendar.google.com)
- [Youtube Channel](https://www.youtube.com/channel/UCwvtktvdZu_pg-t-INvuW5g)
