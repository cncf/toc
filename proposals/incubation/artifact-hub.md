# Artifact Hub Incubation Proposal

## Introduction

[Artifact Hub](https://artifacthub.io) is a web-based application that enables finding, installing, and publishing packages and configurations for cloud native projects that are part of an open source non-profit.

Discovering artifacts to use with cloud native projects can be difficult. If every cloud native project that needs to share artifacts creates its own Hub this creates a fair amount of repeat work for each project and a fractured experience for those trying to find the artifacts to consume. The Artifact Hub attempts to solve that by providing a single experience for consumers that any CNCF project can leverage.

At the moment, the following artifacts kinds are supported *(with plans to support more projects to follow)*:

- [CoreDNS plugins](https://coredns.io/)
- [Falco configurations](https://falco.org/)
- [Helm charts](https://helm.sh/)
- [Helm plugins](https://helm.sh/docs/topics/plugins/)
- [KEDA scalers](https://keda.sh/)
- [Keptn integrations](https://keptn.sh)
- [Kubectl plugins (Krew)](https://krew.sigs.k8s.io/)
- [OLM operators](https://github.com/operator-framework)
- [Open Policy Agent (OPA) policies](https://www.openpolicyagent.org/)
- [Tekton tasks and pipelines](https://tekton.dev/)
- [Tinkerbell actions](https://tinkerbell.org/)

Artifact Hub allows users to:

- [Discover](https://artifacthub.io/packages/search), [install](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=install) and [publish](https://artifacthub.io/docs/topics/repositories/) packages and configurations
- Explore content like Helm charts [schemas](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=values-schema) and [templates](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=template) in an interactive way
- Subscribe to packages' new releases and security alerts notifications, via email or webhooks
- Visualize packages' [security reports](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub/1.2.0?modal=security-report)
- Inspect packages' [changelog](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=changelog)

## Incubation stage requirements

#### Document that it is being used successfully in production by at least three independent end users

As of today, more than **800 registered organizations** are publishing their content on `artifacthub.io`. **1.2K repositories** containing artifacts of *12 different kinds* have been listed so far, which accounts for a total of [6K packages and 95K+ releases](https://artifacthub.io/packages/search). The site has more than **8K registered users** and 300K unique users have used it during the last year. These numbers keep growing by the day. For more details and some graphs please see the [Artifact Hub stats page](https://artifacthub.io/stats).

Some examples of organizations listing their content on Artifact Hub:

- [Atlassian](https://artifacthub.io/packages/search?org=atlassian)
- [Bitnami](https://artifacthub.io/packages/search?org=bitnami)
- [Datadog](https://artifacthub.io/packages/search?org=datadog)
- [Hashicorp](https://artifacthub.io/packages/search?org=hashicorp)
- [Prometheus](https://artifacthub.io/packages/search?org=prometheus)

It's also worth mentioning that the Helm Hub was moved to Artifact Hub on [October 2020](https://helm.sh/blog/helm-hub-moving-to-artifact-hub/).

In addition to this, we know of some organizations that are also considering to use Artifact Hub (the software) internally, but we don't have yet a proper list of adopters.

#### Have a healthy number of committers

The project currently has 3 [maintainers](https://github.com/artifacthub/hub/blob/master/OWNERS) from 2 different organizations:

- Sergio Castaño Arteaga (Linux Foundation)
- Cintia Sánchez García (Linux Foundation)
- Matt Farina (SUSE)

#### Demonstrate a substantial ongoing flow of commits and merged contributions

The [project's changelog](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=changelog) lists all versions released so far, including when they were released and the changes introduced on each of them.

Some interesting DevStats links:

- [Activity](https://artifacthub.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&from=now-1y&to=now&var-period=w&var-repogroups=All)
- [Contributions](https://artifacthub.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-1y&to=now&var-period=w&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all)
- [Stars and forks](https://artifacthub.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now)
- [Issues](https://artifacthub.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=now-1y&to=now&var-period=w&var-repogroup_name=All)
- [Pull requests](https://artifacthub.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)
- [Commits](https://artifacthub.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=w&var-repogroups=All)

Please see [Artifact Hub DevStats](https://artifacthub.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m) for more details.

#### A clear versioning scheme

Artifact Hub follows [semantic versioning](https://semver.org).

The project [maintains a detailed changelog](https://artifacthub.io/packages/helm/artifact-hub/artifact-hub?modal=changelog) that lists the changes introduced by each version.

#### Clearly documented security processes explaining how to report security issues to the project

The process to report a security problem in Artifact Hub is described in the [SECURITY.md document](https://github.com/artifacthub/hub/blob/master/SECURITY.md).
