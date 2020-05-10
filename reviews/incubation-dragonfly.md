# Dragonfly Incubating Stage Review

Dragonfly is currently a CNCF sandbox project. Please refer to Dragonfly's initial
[sandbox proposal](../proposals/dragonfly.adoc) for discussion on Dragonfly's
alignment with the CNCF and details on sandbox requirements.

In the time since being accepted as a sandbox project, Dragonfly has demonstrated
healthy growth and progress.

* [v0.4.3 is the latest
releases](https://github.com/dragonflyoss/Dragonfly/releases), shipped on
August 5th, 2019. New features include:
  * [Support Prometheus for monitoring](https://github.com/dragonflyoss/Dragonfly/pull/669)
  * [Support image preheat with Harbor](https://github.com/dragonflyoss/Dragonfly/issues/108)
  * [Support HTTPS private image distribution](https://github.com/dragonflyoss/Dragonfly/pull/523)
  * [Optimize efficiency on scheduling and disk IO](https://github.com/dragonflyoss/Dragonfly/pull/819)
  * [refactoring the whole Dragonfly with Golang](https://github.com/dragonflyoss/Dragonfly/issues/346)

* A [formalized governance
policy](https://github.com/dragonflyoss/community/blob/master/GOVERNANCE.md) has been approved and instituted for the project

## Incubating Stage Criteria

In addition to sandbox requirements, a project must meet the following
criteria to become an incubation-stage project.

* Document that it is being used successfully in production by at least three
independent end users which, in the TOC’s judgement, are of adequate quality
and scope.
  * Dragonfly has a list for [https://github.com/dragonflyoss/Dragonfly/blob/master/adopters.md](https://github.com/dragonflyoss/Dragonfly/blob/master/adopters.md) in varies of industries. Some are listed in the following:
    * telecom industry: ChinaMobile, ChinaUnicom, Huawei, ZTE;
    * e-commerce industry: Taobao, JD.com, Shopee, Lazada;
    * cloud providers: Alibaba Cloud, wise2c, Caicloud, DaoCloud;
    * live streaming industry: Huya.com, bilibili;
    * public life service industry: Meituan, ele.me;
    * artificial intelligence industry: iFLYTEK.
    
* Have a healthy number of committers. A committer is defined as someone with
the commit bit; i.e., someone who can accept contributions to some or all of
the project.
  * Maintainers of the project are listed in
[https://github.com/dragonflyoss/Dragonfly/blob/master/MAINTAINERS.md](https://github.com/dragonflyoss/Dragonfly/blob/master/MAINTAINERS.md). 
  * There are 7 maintainers from 4 different companies (Alibaba Group, ByteDance, eBay and Meitu).
  * Maintainers are added and removed from the project as per the policies
outlined in the project governance:
[https://github.com/dragonflyoss/community/blob/master/GOVERNANCE.md](https://github.com/dragonflyoss/community/blob/master/GOVERNANCE.md).

* Demonstrate a substantial ongoing flow of commits and merged contributions.
  * Releases: 11 major releases ([https://github.com/dragonflyoss/Dragonfly/releases](https://github.com/dragonflyoss/Dragonfly/releases))
  * Roadmap: [https://github.com/dragonflyoss/Dragonfly/blob/master/ROADMAP.md](https://github.com/dragonflyoss/Dragonfly/blob/master/ROADMAP.md)
  * Contributors: 51 contributors [https://github.com/dragonflyoss/Dragonfly/graphs/contributors](https://github.com/dragonflyoss/Dragonfly/graphs/contributors)
  * Commit activity: [https://github.com/dragonflyoss/Dragonfly/graphs/commit-activity](https://github.com/dragonflyoss/Dragonfly/graphs/commit-activity)
  * CNCF DevStats: [https://dragonfly.devstats.cncf.io/](https://dragonfly.devstats.cncf.io/)
    * [Last 30 days activity on GitHub](https://dragonfly.devstats.cncf.io/d/8/dashboards?refresh=15m&orgId=1&from=now-30d&to=now-1h)
    * [Community Stats](https://dragonfly.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1)
