# OpenCost Annual Review 2022

June 2022-September 2023

## Background

[OpenCost](https://opencost.io) is a vendor-neutral open source project for measuring and allocating Kubernetes infrastructure and container costs in real time. This project combines a [specification](https://github.com/opencost/opencost/blob/develop/spec) as well as a Golang implementation of these detailed requirements. OpenCost currently fully supports Amazon Web Services, Microsoft Azure, Google Cloud Platform, and on-premises with partial support for Alibaba Cloud and Scaleway.

OpenCost was announced as a CNCF Sandbox Project in June of 2022. It was originally developed and open-sourced by [Kubecost](https://kubecost.com/).

## Development

There are three [OpenCost GitHub](https://github.com/opencost) repositories:
* [opencost/opencost](https://github.com/opencost/opencost) - the primary repository for the project, it contains the [OpenCost Specification](https://github.com/opencost/opencost/blob/develop/spec), the Golang implementation of the specification, and the Node/React [web interface](https://github.com/opencost/opencost/blob/develop/ui/).
* [opencost/opencost-helm-chart](https://github.com/opencost/opencost-helm-chart) - the Helm chart for configuring and deploying OpenCost with a wide variety of installation options.
* [opencost/opencost-website](https://github.com/opencost/opencost-website) - the source for the [OpenCost](https://opencost.io) website and documentation.

### DevStats

The OpenCost devstats page and dashboard are found [here](https://opencost.devstats.cncf.io/).

The project currently has PRs from [113 contributors](https://opencost.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=OpenCost) across [41 companies](https://opencost.devstats.cncf.io/d/21/prs-authors-companies-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=OpenCost). According to devstats, over the last year the project averages [437 contributions](https://opencost.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=OpenCost&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) from 43 contributors per month.

### Releases

OpenCost repositories follow the [Semantic Versioning 2.0.0](https://semver.org/) convention. So far there have been 12 OpenCost releases, with the initial OpenCost 1.94 released in June 2022 and as of September 2023 the current release is 1.106. The 1.x series are considered stable and have been in production for years.

The release cadence is:
* There have been no additional major releases since OpenCost was added to the CNCF.
* Minor versions are released every month.
* Patch versions are released as necessary.
* The releases are announced using GitHub’s releases, followed by Slack (CNCF/#opencost), Mastodon, Twitter, and discussion in the fortnightly Community Meetings.

### Maintainers

The OpenCost project has documented the [governance policy](https://github.com/opencost/opencost/blob/develop/GOVERNANCE.md) for designating [Committers](https://github.com/orgs/opencost/teams/opencost-committers) and [Maintainers](https://github.com/opencost/opencost/blob/develop/MAINTAINERS.md) per repository and we are transitioning towards a more community-driven development model. Current development decisions are primarily discussed in GitHub Issues and in the respective working group and community meetings. To ensure transparent decision-making, OpenCost will adopt a Request for Comments (RFC) process similar to other projects for any substantial changes to the project.

* OpenCost currently has 7 Maintainers and 12 Committers from 4 organizations with expanded permissions in GitHub (there are 75  contributors to the repository).
* The OpenCost Helm chart has 2 Maintainers and 31 contributors so far.
* The OpenCost website has 33 contributors to date.

## Adoption

Given the single deployment of OpenCost per Kubernetes cluster, there are easily thousands of deployments of OpenCost in production. There are potentially orders of magnitude greater when considering inclusion in Grafana Cloud's [Kubernetes cost monitoring](https://www.opencost.io/blog/grafana-cloud), Kubecost’s commercial offering, and being embedded in several other commercially supported products. There is an [ADOPTERS.md](https://github.com/opencost/opencost/blob/develop/ADOPTERS.MD) with several organizations publicizing their usage.

The OpenCost containers average over [10k daily downloads](https://quay.io/repository/kubecost1/kubecost-cost-model).

### Community

OpenCost strives to become a supportive and welcoming community. Most conversations occur in the [CNCF Slack](http://slack.cncf.io) #opencost channel and through GitHub Issues and Pull Requests. There is an [OpenCost Community Calendar](https://bit.ly/opencost-calendar) with a fortnightly recurring [Community Meeting](https://bit.ly/opencost-meeting) ([meeting notes](https://docs.google.com/document/d/1JFB_-sjF8F9UWet1c-dWixdMZY4hri23UlQG5FX5xfY/edit)). OpenCost also has a presence on [LinkedIn](https://www.linkedin.com/showcase/opencost/), [Mastodon](https://hachyderm.io/@opencost), and [Twitter](https://twitter.com/open_cost). Additionally OpenCost has been presented at several community events and had a presence at the KubeCon EU and will have a kiosk and working group session at KubeCon NA.

Since becoming a CNCF project OpenCost has seen substantial growth in the community.

| [Stats](https://opencost.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=OpenCost)              | Current (September 11, 2023) | Sandbox |
| ------------------ | --------------------------- | ------- |
| Stargazers         | 4025 | 2000 |
| Commits            | 4734 | 2739 |
| Forks              | 451 | 192 |
| Contributors       | 446 | 40 |
| CNCF Slack members | 702 | - |
| LinkedIn followers | 214 | - |

## Roadmap

The current roadmap may be found [here](https://github.com/opencost/opencost/blob/develop/ROADMAP.md). We will revisit it at the upcoming KubeCon NA during our working group session.

### Feature Timeline

* The initial 1.94 OpenCost release was made in June of 2022.
* Support for GCP spot instances was added in the 1.95 release.
* Additional AWS regions were added in 1.96.
* The 1.97 release included initial support for European cloud provider ScaleWay.
* Kubernetes 1.25 support was included with 1.98.
* The OpenCost UI was included in the default installation starting with version 1.99.
* Additional documentation for installation, upgrade, uninstall and Helm Chart instructions have all been greatly expanded.
* Version 1.99 also included initial support for China-based cloud provider Alibaba Cloud.
* 1.100 release including performance enhancements with code contributions from [Grafana Labs](https://www.opencost.io/blog/grafana-labs).
* The initial 1.0.0 release of the OpenCost Helm chart provides expanded installation options.
* A [swagger.json](https://raw.githubusercontent.com/opencost/opencost/develop/docs/swagger.json) OpenAPI document is now provided and [API documentation](https://www.opencost.io/docs/integrations/api) was added to the OpenCost website.
* OpenCost 1.103 brought enhanced Azure billing integration and a daily CSV export of allocation data with code contributions from [Microsoft](http://aka.ms/aks/OpenCost-AKS).
* OpenCost 1.104 included terminology based on the [FinOps Open Cost and Usage Specification](https://www.opencost.io/blog/focus).
* Documentation for [custom on-premises pricing](https://www.opencost.io/docs/configuration/on-prem) is now available.
* 1.104 Kubecost contributed non-k8s billing integration and cost monitoring to the project for future enhancements.
* Configuring the [AWS Spot Instance Data Feed](https://www.opencost.io/docs/configuration/AWS#aws-spot-instance-data-feed-configuration) is now documented.
* The [OpenCost Backstage plugin](https://github.com/backstage/backstage/tree/master/plugins/opencost) will be included in the upcoming 1.18.0 Backstage release.

### Goals

The goal of OpenCost is to be the default real-time cloud cost monitoring solution for Kubernetes.

* Increase adoption and awareness. We're working on this through community involvement, support, and outreach through events and social platforms.
* Cultivate community developer base with more contributors and maintainers from additional organizations. This is part of our community outreach as well as additional developer exposure through integrations with other projects.
* Expanded public cloud coverage through soliciting contributions for supporting additional cloud providers and developer documentation.
* External Asset Costs will bring support for services not directly managed by Kubernetes infrastructure. Examples of currently untracked costs that could be linked to Kubernetes infrastructure include network costs, monitoring, and external database services. The initial External Asset Costs prototype will add support for object storage across AWS, Azure, and GCP with documentation and a new specification being developed in parallel. These efforts should encourage external contributions from both vendors and users.
* Carbon footprint costs. Initial conversations with the TAG Environmental Sustainability were had at KubeCon EU with the intention of adding the carbon footprint for Kubernetes workloads once a suitable billing API source is available. Microsoft Azure has made metrics available through “early access” in their billing API, we hope to start development soon. Similar conversations have been had with ThoughtWorks Cloud Carbon Footprint project and will likely start development in parallel.
* The FinOps Foundation has participated in development conversations as OpenCost looks to standardize on the [FinOps Open Cost and Usage Specification](http://focus.finops.org/) (FOCUS). Initial FOCUS patches have already been merged and development continues.

### How the CNCF can help to achieve the upcoming goals

We appreciate the support provided by the CNCF and look forward to participating in events and online presentations. End-user awareness and contributor recruitment are our biggest needs at this point. Any assistance with our [OpenCost Incubation proposal](https://github.com/cncf/toc/pull/1046/) would be greatly appreciated and expand our opportunities.

## Incubation

We have submitted a [proposal for OpenCost Incubation](https://github.com/cncf/toc/pull/1046/) and are working with our TOC Sponsor Ricardo Rocha. We are presenting to the TAG Observability September 19, 2023.
