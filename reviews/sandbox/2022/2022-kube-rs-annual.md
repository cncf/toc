# kube-rs 2022 Annual Review

This is a Kubernetes ToC Annual Review for the [kube-rs](https://kube.rs) project for 2022.

<!-- omit in toc -->
## Table of Contents

- [kube-rs 2022 Annual Review](#kube-rs-2022-annual-review)
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

[kube-rs](https://kube.rs) is a suite of Rust libraries for Kubernetes for building controllers/operators, or just in general interacting with the apiserver.

kube focuses on fulfilling a similar role to client-go, kubebuilder, and controller-runtime in the go ecosystem.

kube was accepted as a CNCF Sandbox project on [November, 2021](https://clomonitor.io/projects/cncf/kube-rs).

## DevStats

[kube-rs DevStats page](https://kubers.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1609520400000&to=1641056399000&var-period=w) shows the weekly number of newly opened pull requests on the main git repo.

[Contribution chart for 2022 shows a slight upward trend with steady progress](https://kubers.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=1640995200000&to=1672531200000).

The amount of maintainers have stayed the same, but the amount of changes from external contributors adopting the libraries have increased.
86 contributors in the last year which is the highest so far: https://kubers.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-companies=All

~30 different companies contributing last year: https://kubers.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions

Stars / stats generally going up:
https://kubers.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=1641056400000&to=1672592399000&var-period=w
Main repo went from ~1100 stars to ~1600 stars (now at ~2200).

## Maintainers

There are three maintainers working on the project.

| Maintainer                | GitHub ID                              | Affiliation  |
| ------------------------- | -------------------------------------- | ------------ |
| Eirik Albrigtsen          | [@clux](https://github.com/clux)       | Unaffiliated |
| Natalie Klestrup Röijezon | [@nightkr](https://github.com/nightkr) | Stackable    |
| Kaz Yoshihara             | [@kazk](https://github.com/kazk)       | Andela       |

Also on https://kube.rs/maintainers/

## Adoption

kube-rs is widely used in production by the a considerable number of companies that we know about.
Some are listed in https://kube.rs/adopters/
And some recent changes to adopters can be seen in https://github.com/kube-rs/website/commits/main/docs/adopters.md

Note that since we are series of libraries there are big dark numbers here.
We see [4000 repo dependencies on github](https://github.com/kube-rs/kube/network/dependents), [75 package dependencies](https://github.com/kube-rs/kube/network/dependents?dependent_type=PACKAGE), as well as [68 reverse cargo dependencies (published)](https://crates.io/crates/kube/reverse_dependencies).

## Latest Achievements

- Support for Kubernetes metadata api
- Kubernetes Stable Client Badge
- Lots of new documentation on kube.rs
- Creation of go->rust crd conversion tool (kopium)
- Adoption from major projects such as Linkerd
- Lots of positive PR from KubeCon [1](https://www.youtube.com/watch?v=Kp6GQjZixPE), [2](https://www.youtube.com/watch?v=ONkWwoJoF3I) and more [3](https://www.youtube.com/watch?v=65pyIeLtd5Y)

## Project Goals

<!-- What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation? -->

While we have managed to stabilise some interfaces (like the client), we do want to release a 1.0 of the crate suite, and provide full documentation for controller writing, and achieve parity for a few more common cases.

The few big outstanding projects that we want to resolve pre-1.0 are;

- Stream sharing in kube::runtime
- Protobufs support via k8s-pb
- Client usage without Api (could wait)

We are hoping that the crate suite starts to stand on its own without too much feature work required beyond the main ones listed. It is already fulfilling its job as a client api and operator framework for many companies, and hope that smaller gaps and issues will close over time. As maintainers inevitably shift jobs, documentation efforts are also being focused so that future maintainers can be onboarded easily.

## CNCF Membership

<!-- How can the CNCF help you achieve your upcoming goals? -->

CNCF staff has been very responsive when we had enquiries in the past. We do not require any specific help at this stage.

## Incubation

We have [numerous adopters](https://kube.rs/adopters) building Kubernetes controllers with kube-rs in production (including linkerd), and most of users proposals are generally considered/refined and make it into kube to form a steady stream of (monthly or bi-monthly) [releases](https://github.com/kube-rs/kube/releases).

We have a well defined [release process and versioning](https://kube.rs/release-process/) (with compatibilities defined against both [Kubernetes](https://kube.rs/kubernetes-version/) and [Rust](https://kube.rs/rust-version/)), a [security policy](https://kube.rs/security/), and a [reference implementation of a rust controller](https://github.com/kube-rs/controller-rs).

We feel we probably satisfy the requirements for the [incubating stage](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage), depending on the interpretations of the sizes required for our scope.

## Project Links

- [Website](https://kube.rs/)
- [Github](https://github.com/kube-rs/kube)
- [Discord](https://discord.gg/tokio)
