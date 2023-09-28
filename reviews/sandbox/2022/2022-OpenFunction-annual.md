# OpenFunction 2022 Annual Review

## Background

OpenFunction is a cloud-native open source FaaS (Function as a Service) platform aiming to let you focus on your business logic without having to maintain the underlying runtime environment and infrastructure. You can generate event-driven and dynamically scaling Serverless workloads by simply submitting business-related source code in the form of functions.

Its core features includes:
- Cloud agnostic and decoupled with cloud providers’ BaaS
- Pluggable architecture that allows multiple function runtimes
- Support both sync and async functions
- Unique async functions support that can consume events directly from event sources
- Support generating OCI-Compliant container images directly from function source code.
- Flexible autoscaling between 0 and N
- Advanced async function autoscaling based on event sources’ specific metrics
- Simplified BaaS integration for both sync and async functions by introducing Dapr
- Advanced function ingress & traffic management powered by K8s Gateway API (In Progress)
- Flexible and easy-to-use events management framework
- Support using WasmEdge as workload runtime

## Sandbox Acceptance

OpenFunction was accepted as a CNCF sandbox project on April 27th, 2022. 

## DevStats

The OpenFunction devstats page and dashboard could be found [here](https://openfunction.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=1650988800000&to=now).

* The project currently has [35 contributors](https://openfunction.devstats.cncf.io/d/22/prs-authors-table?orgId=1) from [15 companies] (https://openfunction.devstats.cncf.io/d/21/prs-authors-companies-table?orgId=1&var-period_name=Last%202%20years&var-repogroup_name=All). 
  
* Since joining CNCF Sandbox, we've seen [a steady growth](https://openfunction.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=1650988800000&to=now):
  - Totally 35 contributors created 398 PRs (964 commits in total) with 594 PR reviews from 18 PR reviewers, you can find more details [here](https://openfunction.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Since%20joining%20CNCF&var-repogroup_name=All). 
  - Number of starts: 536 -> 1166
  - Number of forks: 70 -> 142

## Maintainers

We've established the [TOC](https://github.com/OpenFunction/community/blob/main/TOC.md) with 7 members from 6 companies. The core maintainers includes:

* Benjamin Huo (QingCloud Technologies) @benjaminhuo
* Wanjun Lei (QingCloud Technologies) @wanjunlei
* Yifei Wang (QingCloud Technologies) @wrongerror
* Tian Fang (AscentStream) @tpiperatgod
* Haili Zhang (UISEE) @webup
* Lize Cai (SAP) @lizzzcai

We have been holding community meeting every two weeks since Feb 2022. The meeting notes could be found [here](https://github.com/OpenFunction/community/discussions?page=1), and the meeting recording can be found [here](https://space.bilibili.com/438908638/search/video?keyword=OpenFunction%20%E4%BE%8B%E4%BC%9A).

## Project goals & updates

We're going to evolve OpenFunction into a modern cloud agnostic Serverless & FaaS platform that can be a open source alternative for Serverless & FaaS platforms from cloud providers.

The OpenFunction community has released 6 versions since joining CNCF:
- [v1.1.1](https://github.com/OpenFunction/OpenFunction/releases/tag/v1.1.1)
- [v1.1.0](https://github.com/OpenFunction/OpenFunction/releases/tag/v1.1.0)
- [v1.0.0](https://github.com/OpenFunction/OpenFunction/releases/tag/v1.0.0)
- [v0.8.1](https://github.com/OpenFunction/OpenFunction/releases/tag/v0.8.1)
- [v0.8.0](https://github.com/OpenFunction/OpenFunction/releases/tag/v0.8.0)
- [v0.7.0](https://github.com/OpenFunction/OpenFunction/releases/tag/v0.7.0)

The feature highlights includes:
- Support OpenFunction Gateway powered by K8s Gateway API in v0.7.0
- Add Dapr proxy to support Dapr standalone mode instead of Dapr sidecar mode in v0.8.0
- Support WasmEdge runtime in v1.0.0
- Support Dapr State Management & Trigger refactoring

## Adoptions

- The autonomous driving company [UISEE](https://www.uisee.com/en/index.html) is using OpenFunction to process their vehicle data in the cloud in a Serverless style.
- China Unicom is using OpenFunction to build FaaS platform for their cloud
- WeBank is evaluating OpenFunction to process their data and also planning to integrate OpenFunction with their [Apache EventMesh project](https://eventmesh.apache.org/)
- [ximalaya](https://www.ximalaya.com/) is using OpenFunction to process their audio data
- [YuanShan AI](http://www.yuanshan-ai.com/) is using OpenFunction to process their data from Edge IoT devices

## Incubation readiness

Not yet, OpenFunction need more time to be more mature.