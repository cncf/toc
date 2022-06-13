# Trickster Sandbox Annual Review

Trickster is an HTTP Proxy Cache that specializes in accelerating requests for Time Series data, with support for Prometheus, InfluxDB, ClickHouse and others. Trickster became a CNCF Sandbox Project in March of 2021.

## Annual Review Questions

### Metrics

| Metric       |  2018-2021 |     Today   |
| ------------ | ---------- | ----------- |
| Docker Pulls | 500k+      | 1.9m+       |
| Stars        | 1.3k+      | 1.7k+       |
| Forks        | ~90        | 157         |

The CNCF provides Trickster GitHub metrics through its [devstats site](https://trickster.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

Reviewing the devstats site above, one can observe that Trickster's contributor activity remained steady following the transfer into CNCF. In the fall of 2021, project activity dropped off due to the Project Lead taking a medical leave through spring 2022.

### Maintainers

Trickster currently has [3 maintainers](https://github.com/trickstercache/trickster/blob/main/MAINTAINERS.md):

| Maintainer     | Company                | Role         |
| -------------- | ---------------------- | ------------ |
| [James Ranson](@jranson)   | Virga                  | Project Lead |
| [Chris Randles](@crandles)  | Warner Bros. Discovery | Maintainer   |
| [Adam Ross](@limitlessearth)      | Comcast                | Maintainer   |

All Trickster maintainers are LGBTQ+.

### Adoption

The Trickster project continues to see modest-but steady growth and adoption since entering the CNCF Sandbox. Pull Requests have been submitted by 9 new contributors, while feature request and bug reports have been filed from numerous first time submitters - including Chronosphere, who requested and worked with us on a backend request sharding feature. We do not have an ADOPTERS file, but will plan to add one to the project.

From a user standpoint, we engage the community largely through #trickster channels on the CNCF and Gophers slack orgs. In 2018 Martin Chodur, a now-long-time contributor to the Trickster project, submitted a Trickster 0.x status dashboard to Grafana's online catalog, so users could monitor their Trickster deployments in Grafana. Earlier this year, Martin posted to one of our slack channels that he'd updated the status dashboard to take advantage of all of the new metrics Trickster provides in recent releases. So we see new contributors, we also see long-time users still contributing back to the project.

In June 2022, we welcomed a new regular contributor - a summer intern working at @VirgaIO who has dedicated time each week to work on Trickster.

### Past Goals

Over the last year, our primary goal has been to release version 2.0 of Trickster (features below) at the end of 2021. Trickster 2.0 is currently in beta, with a diverse beta user base filing issues and PRs.

Features completed for 2.0:

- [Time series sharding](https://github.com/trickstercache/trickster/blob/main/docs/timeseries_sharding.md) to 1 or more backends (breaking up a large time range from a user into many smaller time range requests aligned with backend partition schemes)
- Brotli and ZStd Compression - no other proxy cache currently supports all of Brotli, ZStd, Gzip and deflate out of the box
- Support for accelerating additional Grafana data source types, like MySQL
- Support for the Flux query language (InfluxDB)
- [Application Load Balancer](https://github.com/trickstercache/trickster/blob/main/docs/alb.md) with unique capabilities geared towards modern workloads
- Fully exportable, reusable packages; so any golang project can import and use any of Trickster's capabilities in-process

### Current Goals

As the Project Lead recently returned from medical leave, we are ramping up our work on Trickster 2.0, and our primary goal is to deliver a GA release coinciding with CloudNativeCon NA this fall. Our progress on 2.0 is [documented here](https://github.com/trickstercache/trickster/blob/main/docs/new-changed-2.0.md), and we are well on our way to a GA release.

Following the 2.0 GA release this fall, the maintainers will regroup and publish a new roadmap for 2023.

### How can the CNCF help?

We are so appreciative of the work Kapunahele from the CNCF contributed to build out a docs sites at <https://trickstercache.org> for the project. As we approach the GA release later this year, we definitely intend to do a docs refresh, and we were hoping we could work with her on that aspect. And of course, we'd some help spreading the word when our big launch day arrives!

### Incubation

At this time, we're comfortable in the Sandbox, but perhaps next year we'll meet the requirements for incubation!
