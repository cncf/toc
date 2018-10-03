# Harbor Incubating Stage Review

Harbor is currently a CNCF sandbox project. Please refer to Harbor's initial
[sandbox proposal](../proposals/harbor.adoc) for discussion on Harbor's
alignment with the CNCF and details on sandbox requirements.

In the time since being accepted as a sandbox project, Harbor has demonstrated
healthy growth and progress.

* [v1.6.0 is the latest
releases](https://goharbor.io/blogs/harbor-1.6.0-release/), shipped on
September 7th, marking our 7th major feature release. New features include:

  * [Support for hosting Helm charts](https://github.com/goharbor/harbor/issues/4922)
  * [Support for RBAC via LDAP groups](https://github.com/goharbor/harbor/issues/3506)
  * [Replication filtering via labels](https://github.com/goharbor/harbor/issues/4861)
  * [Major refactoring to coalesce to a single PostgreSQL database](https://github.com/goharbor/harbor/issues/4855)

* A [formalized governance
policy](https://github.com/goharbor/community/blob/master/GOVERNANCE.md) has
been approved and instituted for the project, and two new maintainers from
different companies have joined the project to help Harbor continue to grow.

## Incubating Stage Criteria

In addition to sandbox requirements, a project must meet the following
criteria to become an incubation-stage project:

* Document that it is being used successfully in production by at least three
independent end users which, in the TOC’s judgement, are of adequate quality
and scope.

  * Adopters: [https://github.com/goharbor/harbor/blob/master/ADOPTERS.md](https://github.com/goharbor/harbor/blob/master/ADOPTERS.md)

* Have a healthy number of committers. A committer is defined as someone with
the commit bit; i.e., someone who can accept contributions to some or all of
the project.

  * Maintainers of the project are listed in
[https://github.com/goharbor/harbor/blob/master/OWNERS.md](https://github.com/goharbor/harbor/blob/master/OWNERS.md). There are 11 maintainers working on Harbor from 3 different
companies (VMware, Caicloud and Hyland Software)

  * Maintainers are added and removed from the project as per the policies
outlined in the project governance:
[https://github.com/goharbor/community/blob/master/GOVERNANCE.md](https://github.com/goharbor/community/blob/master/GOVERNANCE.md).

* Demonstrate a substantial ongoing flow of commits and merged contributions.

  * Releases: 7 major releases ([https://github.com/goharbor/harbor/releases](https://github.com/goharbor/harbor/releases))

  * Roadmap: [https://github.com/goharbor/harbor/wiki/Harbor-Roadmap](https://github.com/goharbor/harbor/wiki/Harbor-Roadmap)

  * Contributors: [https://github.com/goharbor/harbor/graphs/contributors](https://github.com/goharbor/harbor/graphs/contributors)

  * Commit activity: [https://github.com/goharbor/harbor/graphs/commit-activity](https://github.com/goharbor/harbor/graphs/commit-activity)

  * CNCF DevStats: [https://harbor.devstats.cncf.io/](https://harbor.devstats.cncf.io/)
    * [Last 30 days activity on GitHub](https://harbor.devstats.cncf.io/d/8/dashboards?refresh=15m&orgId=1&from=now-30d&to=now-1h)
    * [Community Stats](https://harbor.devstats.cncf.io/d/3/community-stats?orgId=1&var-period=d7&var-repo_name=goharbor%2Fharbor)

Further details of Harbor's growth and progress since entering the sandbox
stage as well as use case details from the Harbor community can be found in this
[slide
deck](https://docs.google.com/presentation/d/1aBQnE96kKatc1_t3E97lJBwiWvL-3GTitojuv-nWMuo/).

## Security

Harbor's codebase has been analyzed and reviewed by VMware's internal product
security team.

* Static analysis has been performed on Harbor via
[gosec](https://github.com/securego/gosec)
* Software decomposition via AppCheck, Snyk and retire.js with goal of
discovering outdated or vulnerable packages
* Manual code analysis / review
* Vulnerability assessment via multiple scanners
* Completed threat model

In addition to this security work the Harbor maintainers are partnering with
the CNCF to schedule a third-party security audit of Harbor.
