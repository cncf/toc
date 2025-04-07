**Vitess Graduation Proposal**

Since joining the CNCF in [February 2018](https://www.cncf.io/blog/2018/02/05/cncf-host-vitess/) as an incubation project, Vitess has built a healthy ecosystem of maintainers and production adoptions.

On behalf of the maintainers team, we believe Vitess is ready for [graduation stage](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#graduation-stage) and meets the v1.2 criteria. As the criteria has evolved since February 2018, we include also answers for incubation criteria and address [due diligence questions](https://docs.google.com/document/d/1TDlRdgfTiEWunpav-G8gkaQF7Zk84-9tNAXyv1I0Kws/edit) in a separate Google Docs document.

### CNCF Incubation Criteria

**Production usage by three independent users**

Vitess can be classified as both a middleware and an orchestration system for managing sharded MySQL systems. It depends on a MySQL-like database (most deployments currently use MySQL 5.7) and a topology server (etcd is now the default, but ZooKeeper and Consul are also supported).

Adoption depends on a high degree of trust from users, since Vitess is responsible for the integrity and durability of canonical data, and connectivity from applications to the database depend on Vitess being available. Vitess unavailability generally causes a production outage for most users.

Several Vitess users have been featured in CNCF case studies:

* JD.com was [recently featured](https://www.cncf.io/jdcom-case-study/).
* Square uses Vitess for their [Cash App](https://www.cncf.io/square-case-study/)
* Case studies on Nozzle and Slack will be published shortly.

In addition to these case studies, Slack has estimated that Vitess powers 40%+ of their database servers, with the goal of migrating remaining databases by mid 2020.

**Have a healthy number of committers**

The Vitess [maintainers team](https://github.com/orgs/vitessio/teams/maintainers) team currently consists of 14 members, with a healthy distribution of corporations represented (Hubspot, Pinterest, Nozzle, PlanetScale, Slack, Square).

Over the last 12 months, [130 distinct users](https://vitess.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All) have authored pull requests.

**Demonstrate a substantial ongoing flow of commits and merged contributions**

Vitess maintains a healthy [ongoing flow of commits](https://vitess.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-2y&to=now&var-period=w&var-repogroups=All). In late 2018, documentation moved from the primary repository to vitessio/vitess.

**A clear versioning scheme**

Vitess 3.0 currently uses Semantic Versioning. Similar to [other](http://www.databasesoup.com/2016/05/changing-postgresql-version-numbering.html) [database](https://www.cockroachlabs.com/blog/calendar-versioning/) products semver has been difficult to follow. Following the release of Vitess 4.0 in November 2019, Vitess 5.0 will be released in February 2020 under the [Vitess Release Cycle](https://github.com/vitessio/enhancements/blob/master/veps/vep-1.md).

### CNCF Graduation Criteria

**Committers from at least two organizations.**

The [devstats](https://vitess.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&var-period=m&var-repogroup_name=All&from=now-2y&to=now) from the preceding 2 years show contributions from an average of 14 organizations.

Committers by company [for the preceding 12 months](https://vitess.devstats.cncf.io/d/4/company-statistics-by-repository-group?orgId=1&var-period=m&var-metric=committers&var-repogroup_name=All&var-companies=All&from=now-1y&to=now) lists the top ten organizations as:

1. PlanetScale
2. Square
3. Slack
4. HubSpot
5. Google
6. Nozzle
7. Independent
8. Pinterest
9. SharpSpring
10. BoCloud

**Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.**

We have achieved passing level criteria for [CII Best Practices](https://bestpractices.coreinfrastructure.org/en/projects/1724). A badge is [displayed](https://github.com/vitessio/vitess) in the primary repository's `README.md`.

**Have completed an independent and third party security audit**

A third-party security audit funded by the CNCF was performed in February 2019. The [full report](https://vitess.io/blog/2019-03-12-vitess-security-audit/) is available online. Only minor issues were discovered by the Cure53 team. To emphasize the first few sentences from the report's conclusion:

> The results from this Cure53 assessment funded by CNCF / The Linux Foundation certify that the Vitess database scaler is secure and robust. This very good outcome is achieved by limiting the attack surface, taking appropriate care of user-supplied input with security-driven best practices, as well as - to a certain extent - the usage of the Go language ecosystem. A team of five Cure53 testers investigated the software system during a budgeted period of 18 days in February 2019.

**Adopt the CNCF Code of Conduct.**

We have adopted the CNCF Code of Conduct. A [link](https://github.com/vitessio/vitess/blob/master/CODE_OF_CONDUCT.md) is provided in the primary repository.

**Explicitly define a project governance and committer process.**

The Vitess project governance is defined in the [primary repository](https://github.com/vitessio/vitess/blob/master/GOVERNANCE.md). To draw attention to key details:

* Contributors can propose pull requests, which are reviewed and approved by maintainers.
* New maintainers are nominated by any existing maintainer. There are no special requirements, other than to have shown a willingness and ability to participate in the project as a team player. Maintainers cannot approve their own PRs.
* A Project Management Committee (PMC) of 5 members oversees the process, and is responsible for ensuring the [Guiding Principles](https://github.com/vitessio/vitess/blob/master/GUIDING_PRINCIPLES.md) are followed. Sugu Sougoumarane is the founding and current PMC chair.

**Have a public list of project adopters for at least the primary repository.**

The primary repository lists [key project adopters](https://github.com/vitessio/vitess/blob/master/ADOPTERS.md).

A list of logos is also prominently displayed on the [Vitess.io homepage](https://vitess.io) under "Who uses Vitess". A link to instructions on how to [add a user logo](https://github.com/vitessio/website/#adding-a-user-logo) are included.

**Receive a supermajority vote from the TOC to move to graduation stage.**

We believe this document prepares us for a TOC vote on the graduation readiness of Vitess.
