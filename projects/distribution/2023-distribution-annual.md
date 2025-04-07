# Distribution 2023 Annual Review

This is the annual review for the [distribution](https://github.com/distribution/distribution) project for 2023.

## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Project goals](#project-goals)
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)

## Background

Distribution provides an implementation of the Open Source Registry for storing and distributing container images [and other application content] via the [OCI Distribution Specification](https://github.com/opencontainers/distribution-spec). The goal of this project is to provide a simple, secure, and scalable base for building a large scale registry solution or running a simple private registry. It is a core building block for many registry operators including [Docker Hub](https://hub.docker.com/), GitHub Container Registry, [GitLab Container Registry](https://gitlab.com/gitlab-org/container-registry), [DigitalOcean Container Registry](https://www.digitalocean.com/products/container-registry), as well as the [CNCF Harbor Project](https://goharbor.io/), and VMware Harbor Registry.

Additionally, the project provides various libraries (a.k.a. `Go` programming language packages) that are widely used across large number of actively maintained OSS projects such as [moby](https://github.com/moby/moby/), [argocd](https://github.com/argoproj-labs), etc.

## DevStats

Distribution DevStats dashboard can be found [here](https://distribution.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=now-2y&to=now).

* Stargazers and forks have been steadily growing as can be seen [here](https://distribution.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1).
* Commits in the past month show that the repository is active (see [here](https://distribution.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&var-period=m&var-repogroups=All&from=now-6M&to=now))
* Contributors and companies stats show steady activity as can be seen [here](https://distribution.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&var-period=d7&var-repogroup_name=All&from=now-1y&to=now), but so is the growing list of first time contributors as seen [here](https://distribution.devstats.cncf.io/d/52/new-contributors-table?orgId=1)

## Maintainers

The full list of the current maintainers can be found [here](https://github.com/distribution/distribution/blob/main/MAINTAINERS).

## Adoption

There is very little doubt the project's adoption continue to be quite high and is steadily growing due to more companies and developers hosting their own private registries. The project is at the core of the services provided by varying sizes of organisations from the small ones all the way to the large enterprises. Equally, the libraries provided by the project have carried on being used by many OSS projects including various container runtime implementations.

## Project goals

The Distribution project has the long term goal of providing a secure tool chain for distributing content hat allow users to:
* Take advantage of an efficient, secure and reliable way to store, manage, package and exchange content
* Hack/roll their own on top of healthy open-source components
* Implement their own home made solution through good specs, and solid extensions mechanism.

The most immediate goal of the project is to release `v3` version, that's been in making for a while.

## CNCF membership

Distribution mainly benefited from CNCF by gaining the insights into the project development and adding of additional security scanning integrations such as GitHub's CodeQL and the OSS Fuzzing system.

## Incubation

At this time Distribution is not ready to apply for incubation due to the pending `v3` release work that must be done before we start considering this as an option.
