# Artifact Hub 2021 Annual Review

## Background

[Artifact Hub](https://artifacthub.io) is a web-based application that enables finding, installing, and publishing packages and configurations for cloud native projects that are part of an open source non-profit.

Discovering artifacts to use with cloud native projects can be difficult. If every cloud native project that needs to share artifacts creates its own Hub this creates a fair amount of repeat work for each project and a fractured experience for those trying to find the artifacts to consume. The Artifact Hub attempts to solve that by providing a single experience for consumers that any CNCF project can leverage.

At the moment, the following artifacts kinds are supported *(with plans to support more projects to follow)*:

- [CoreDNS plugins](https://coredns.io/)
- [Falco configurations](https://falco.org/)
- [Helm charts](https://helm.sh/)
- [Helm plugins](https://helm.sh/docs/topics/plugins/)
- [KEDA scalers](https://keda.sh/)
- [Kubectl plugins (Krew)](https://krew.sigs.k8s.io/)
- [OLM operators](https://github.com/operator-framework)
- [Open Policy Agent (OPA) policies](https://www.openpolicyagent.org/)
- [Tekton tasks](https://tekton.dev/)
- [Tinkerbell actions](https://tinkerbell.org/)

Artifact Hub allows users to:

- [Discover](https://artifacthub.io/packages/search), [install](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=install) and [publish](https://artifacthub.io/docs/topics/repositories/) packages and configurations
- Explore content like Helm charts [schemas](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=values-schema) and [templates](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=template) in an interactive way
- Subscribe to packages' new releases and security alerts notifications, via email or webhooks
- Visualize packages' [security reports](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub/0.19.0?modal=security-report)
- Inspect packages' [changelog](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=changelog)

## DevStats

[Artifact Hub DevStats](https://artifacthub.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m)

Some interesting DevStats links:

- [Activity](https://artifacthub.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&from=1581165447000&to=1625047861000&var-period=w&var-repogroups=All)
- [Contributions](https://artifacthub.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=1581165447000&to=1625047861000&var-period=w&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all)
- [Stars and forks](https://artifacthub.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=1581165447000&to=1625047861000)
- [Issues](https://artifacthub.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1581165447000&to=1625047861000&var-period=w&var-repogroup_name=All)
- [Pull requests](https://artifacthub.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1581165447000&to=1625047861000)
- [Commits](https://artifacthub.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=1581165447000&to=1625047861000&var-period=w&var-repogroups=All)

## Maintainers

The project currently has 3 maintainers from 2 different organizations:

- Sergio Castaño Arteaga (Linux Foundation)
- Cintia Sánchez García (Linux Foundation)
- Matt Farina (SUSE)

## Adoption

The project has experienced a lot of growth since it joined the CNCF as a sandbox project last year:

- The Helm Hub was moved to Artifact Hub on [October 2020](https://helm.sh/blog/helm-hub-moving-to-artifact-hub/).
- More than 600 registered organizations are now publishing content on `artifacthub.io`.
- We have more than 6K registered users, 200K+ unique users have visited the site during the last year and 4K+ users rely on `artifacthub.io` on a daily basis (just 3 months ago it was only 1.5K daily users).
- Almost 900 repositories with artifacts of 9 different kinds have been added so far, which accounts for a total of [3.6K packages and 50K releases](https://artifacthub.io/packages/search). Those numbers keep growing by the day.

For more details and some graphs please see the [Artifact Hub stats page](https://artifacthub.io/stats).

In addition to this, we know of some organizations that are also considering to use Artifact Hub (the software) internally, but we don't have yet a proper list of adopters.

## Current goals

For the next year we'd like to achieve the following goals:

- **Grow user base**. The project has reached a certain level of maturity and stability (version `1.0.0` has just been released -June 22nd-), but it's still a relatively new project and many users don't know about it yet. Our goal is to onboard as many users as possible so that they can help shaping the future of Artifact Hub with their contributions and ideas.
- **New features**. We want to keep improving the users’ experience by adding more [useful features](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=changelog).
- **More contributors**. We want to grow the number of contributors to the Artifact Hub project.

## How can the CNCF help

We think there are two important ways that CNCF can help us with the growth of the project:

- **Growing user base**. The project can use all the visibility it can get to increase its user base, and we are hoping that CNCF can help us with that.
- **Integration with other CNCF projects**. During the last few months, @caniszczyk has helped us adding support for Tekton tasks, KEDA scalers and CoreDNS plugins, and we are looking forward to adding support for more artifacts kinds.

## Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

Yes, we believe that Artifact Hub is ready and we now have [an issue](https://github.com/artifacthub/hub/issues/1418) to track progress at proposing a move to incubation.
