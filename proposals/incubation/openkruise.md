# OpenKruise proposal for CNCF Incubation

## About OpenKruise

[OpenKruise](https://openkruise.io) is an extended component suite for Kubernetes,
which mainly focuses on application automations, such as deployment, upgrade, ops and availability protection.
Most features provided by OpenKruise are built primarily based on CRD extensions.
They can work in pure Kubernetes clusters without any other dependencies.

Now OpenKruise provides the following capabilities:

- **Advanced workloads**, which support not only the basic features that are similar to the upstream Workloads in Kubernetes, but also more advanced abilities such as in-place update, configurable scale/upgrade strategies, parallel operations.
- **Sidecar container management**, which defines, injects and even upgrades sidecar containers with no effect on application containers.
- **Multiple domain management**, which empowers workload to support multi-domain and elastic deployment, so that users can define the rules about how their applications should be deployed over different kinds of nodes, e.g., nodes across availability zones.
- **Enhanced operations**, such as restarting containers in-place, pre-downloading images on specific nodes, controlling containers starting priority in a Pod and distributing resources over multiple namespaces.
- **Application availability protection**, which can prevent unexpected Kubernetes resources deletion during cascading deletion and prevent application from disruption or SLA degradation in voluntary disruption scenarios.

OpenKruise was accepted as a CNCF Sandbox project on Nov 10, 2020.

* [Project Annual Review](https://github.com/cncf/toc/blob/main/reviews/2021-OpenKruise-annual.md)
* [Presentation for TAG-runtime](https://docs.google.com/presentation/d/1bzCNwjRtVJ7YMAYcsVWu_1YLCdvMfYk2uzAlQChdS5w/edit#slide=id.p1)
* [Presentation for TAG-app-delivery](https://docs.google.com/presentation/d/16lH6yXEpHcYZiErnXtJG8QR-C_4g07Jz9h_qIByRpGM/edit#slide=id.g127c7cee73c_2_75)

## Continuous Momentum

The OpenKruise devstats page and dashboards can be found [here](https://openkruise.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

* Least Release: v1.2.0
* Num of Contributors: 50+ => *110+*
* Github Stars: 1700+ => *3200+*
* Github Forks: 220+ => *520+*
* Contributing member organizations: 30+ => *50+*
* Maintainers (including Approvers & Reviewers): 3 (1 company) => *7* (5 companies)
* Docker image Pulls: *10M+*
* Member of DingTalk group + Slack channel: 600+ -> *1300+*

## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._**

Now there are at least 40+ adopters that have used OpenKruise in production, full list of registered adopters can be found [here](https://github.com/openkruise/kruise/issues/289).
Highlighted examples includes:
- [**Alibaba**](https://www.alibaba.com/): Using Kruise for managing all their applications and sidecars in large-scale clusters.
- [**LinkedIn**](https://www.linkedin.com/): Using CloneSet for managing large-scale workloads (in-place updates, enhanced PVC support).
- [**MeiTuan**](https://about.meituan.com/en): Using CloneSet and AdvanceStatefulSet for managing stateless and stateful apps. 
- [**Ctrip**](https://www.ctrip.com/): Using CloneSet and AdvanceStatefulSet for managing stateless and stateful apps. They need inplace and partial updates.
- [**Lyft**](https://www.lyft.com/): Using Kruise CloneSets and bin packing for selective downscaling of stateful nodes.

**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

OpenKruise currently has [110+ different PR authors](https://openkruise.devstats.cncf.io/d/22/prs-authors-table?orgId=1) and contributions from [50+ different organizations](https://openkruise.devstats.cncf.io/d/5/companies-table?orgId=1), which includes commits from [more than 15 organizations](https://openkruise.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=commits).

There are [17](https://github.com/orgs/openkruise/people) project members for OpenKruise, out of which seven are project maintainers from five different organizations,
and a clear [membership definition](https://github.com/openkruise/community/blob/master/community-membership.md) on how to become a member, approver, and maintainer of the project.
The maintainers list can be found [here](https://github.com/openkruise/community/blob/master/MAINTAINERS.md).

**_Demonstrate a substantial ongoing flow of commits and merged contributions._**

A short summary has been provided in the aforementioned community section. Detailed statistics can be found in the following devstats links:
- [Commit Activity](https://openkruise.devstats.cncf.io/d/74/contributions-chart?orgId=1)
- [New contributors since sandboxing](https://openkruise.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1573369200000&to=now)
- [Issues open/closed since sandboxing](https://openkruise.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1599721200000&to=now)
- [PRs open/closed since sandboxing](https://openkruise.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1599721200000&to=now)
- [Contributing organizations](https://openkruise.devstats.cncf.io/d/5/companies-table?orgId=1)

**_A clear versioning scheme._**

OpenKruise uses the [semantic versioning scheme](https://semver.org/#semantic-versioning-specification-semver).
OpenKruise follows a quarterly release cadence with version numbers using format of MAJOR.MINOR.PATCH. The latest release is v1.2.0.

We have releases documented at: https://github.com/openkruise/kruise/releases

The latest 4 releases are:

* [release 1.0.0](https://github.com/openkruise/kruise/releases/tag/v1.0.0)
* [release 1.0.1](https://github.com/openkruise/kruise/releases/tag/v1.0.1)
* [release 1.1.0](https://github.com/openkruise/kruise/releases/tag/v1.1.0)
* [release 1.2.0](https://github.com/openkruise/kruise/releases/tag/v1.2.0)
* release 1.3 (WIP)

## Future plan

We are running our project in a fairly open model so that we can react to the needs of our community promptly. In addition, we have recently introduced a [roadmap](https://github.com/openkruise/kruise/projects/2) page in github so that the community members can easily track where we are going. There are still quite a few critical and high demanding features that are under development, including but not limited to the following:
- A new CRD named PodMarker to mark pods by number, labels, nodes, probes, etc.
- Decoupled liveness probe which supports restart containers on failure with limited rate.
- Binding nodes for statefulset pods after being recreated.
- Enhanced application rollout for traffic shifting and batch upgrades.

## Summary

We believe that OpenKruise is ready for Incubation. The project closely aligns with CNCF values on openness and fairness. It has been used in production by at least 40+ organizations and it seems to meet all of the required criteria to move to the next stage.

We're really looking forward to getting started with the process and we're keen to move as quickly as possible. We - the maintainers - will gladly make ourselves available to assist with any due diligence research, to make any presentations required and generally to assist in any way we can be helpful.
