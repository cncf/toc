# Rook Incubating Stage Review

Rook is currently a sandbox stage project.  Please refer to Rook's [sandbox stage proposal](../proposals/rook.adoc) ("inception" at time of acceptance) for details on the sandbox requirements.

In the time since being accepted to the sandbox stage, Rook has demonstrated healthy growth and progress.
Two releases were completed, starting with v0.7 on February 21st and then v0.8 on July 18th.
With those releases, Rook extended beyond just orchestration of Ceph and has built a framework of reusable specs, logic and policies for [cloud-native storage orchestration of other providers](https://blog.rook.io/rooks-framework-for-cloud-native-storage-orchestration-c66278014df7).
Operators and CRD types were added for both CockroachDB and Minio in the v0.8 release, initial support for NFS is nearly complete, and other storage providers are also in the works.

The CRD types and support for Ceph has graduated to Beta in the v0.8 release, reflecting the increased maturity that has only been possible from impressive engagement from the community.
Other big features for the Ceph operator include automatic horizontal scaling of storage resources, an improved security model, and support for new environments such as OpenShift.

A [formalized governance policy](https://github.com/rook/rook/blob/master/GOVERNANCE.md) has been approved and instituted for the project, and a [new maintainer](https://github.com/rook/rook/blob/master/OWNERS.md) has also been added to help the project continue to grow.

## Incubating Stage Criteria

To be accepted to incubating stage, a project must meet the sandbox stage requirements plus:

* Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

  * Adopters: [https://github.com/rook/rook/blob/master/ADOPTERS.md](https://github.com/rook/rook/blob/master/ADOPTERS.md)

* Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

  * Maintainers of the project are listed in [https://github.com/rook/rook/blob/master/OWNERS.md](https://github.com/rook/rook/blob/master/OWNERS.md).

  * Maintainers are added and removed from the project as per the policies outlined in the project governance: [https://github.com/rook/rook/blob/master/GOVERNANCE.md](https://github.com/rook/rook/blob/master/GOVERNANCE.md).

* Demonstrate a substantial ongoing flow of commits and merged contributions.

  * Releases: [https://github.com/rook/rook/releases](https://github.com/rook/rook/releases)

  * Roadmap: [https://github.com/rook/rook/blob/master/ROADMAP.md](https://github.com/rook/rook/blob/master/ROADMAP.md)

  * Contributors: [https://github.com/rook/rook/graphs/contributors](https://github.com/rook/rook/graphs/contributors)

  * Commit activity: [https://github.com/rook/rook/graphs/commit-activity](https://github.com/rook/rook/graphs/commit-activity)

  * CNCF DevStats: [https://rook.devstats.cncf.io/](https://rook.devstats.cncf.io/)
    * [Last 30 days activity on Github](https://rook.devstats.cncf.io/d/8/dashboards?refresh=15m&orgId=1&from=now-30d&to=now-1h)
    * [Community Stats](https://rook.devstats.cncf.io/d/3/community-stats?orgId=1)

Further details of Rook's growth and progress since entering the sandbox stage as well as use case details from the Rook community can be found in this [slide deck](https://docs.google.com/presentation/d/1DOgAlX0RyB8hzD7KbmXK4pKu9hFFPY9WiLv-LEy38jo/edit?usp=sharing).
