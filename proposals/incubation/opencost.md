# OpenCost proposal for CNCF Incubation

## Table of Contents

* [About OpenCost](#about-opencost)
  * [Essentials](#essentials)
  * [Sandbox Acceptance](#sandbox-acceptance)
* [Progress Since Sandbox](#progress-since-sandbox)
  * [Project Timeline](#project-timeline)
  * [Growing Community](#growing-community)
  * [Expanding Developer Ecosystem](#expanding-developer-ecosystem)
  * [Technical Advances](#technical-advances)
  * [Documentation Improvements](#documentation-improvements)
* [Incubation Stage Requirements](#incubation-stage-requirements)
* [Statement on Alignment with the CNCF Mission](#statement-on-alignment-with-the-cncf-mission)
* [Future Plans](#future-plans)

## About OpenCost

OpenCost is a vendor-neutral open source project for measuring and allocating Kubernetes infrastructure and container costs in real time. It was originally developed and open-sourced by [Kubecost](https://kubecost.com/). This project combines a [specification](https://github.com/opencost/opencost/blob/develop/spec) as well as a Golang implementation of these detailed requirements.

The goal of OpenCost is to be the default real-time cloud cost monitoring solution for Kubernetes. OpenCost currently fully supports Amazon Web Services, Microsoft Azure, Google Cloud Platform, and on-premises with partial support for Alibaba Cloud and Scaleway. There are easily thousands of deployments of OpenCost in production and potentially orders of magnitude greater when considering inclusion in Kubecost’s commercial offering.

### Essentials

- Website: [OpenCost.io](https://opencost.io)
- Source Control: [OpenCost on Github](https://github.com/opencost)
- Code of Conduct: [OpenCost Code of Conduct](https://github.com/opencost/opencost/blob/develop/CODE_OF_CONDUCT.md)
- License: [Apache 2](https://github.com/opencost/opencost/blob/develop/LICENSE)

### Sandbox Acceptance

OpenCost was proposed and accepted as a Sandbox project in June 2022 ([Sandbox project onboarding issue](https://github.com/cncf/toc/issues/854)).

## Progress Since Sandbox

### Project Timeline

* June 2, 2022: [OpenCost accepted as a CNCF Sandbox Project](https://www.opencost.io/blog/introducing-opencost)
* June 14, 2022: First official OpenCost release (1.94). Previous open source releases were called “kubecost-cost-model”
* June 20, 2022: FinOps X presentation [Introducing OpenCost](https://www.youtube.com/watch?v=L_tvecd8ssY) by Webb Brown (Kubecost) & Chris Aniszczyk (CNCF)
* June 24, 2022: Open Source Summit North America presentation [OpenCost Introduction](https://docs.google.com/presentation/d/1TDMu6V-J3Mu0hljthnfEz672LhoQ6bKGAsJm2hV5Amg/edit#slide=id.g11eb13674e9_0_0) by Webb Brown
* July 15, 2022: 1.95 release supporting GCP spot instances
* August 12, 2022: 1.96 release providing additional AWS regions and memory usage reductions
* September 20, 2022: 1.97 release adding initial ScaleWay support
* November 4, 2022: 1.98 release supporting Kubernetes 1.25 and performance enhancements
* November 23, 2022: [OpenCost website](http://github.com/opencost/opencost-website/) moved to OpenCost GitHub organization
* December 13, 2022: [OpenCost Helm Chart](http://github.com/opencost/opencost-helm-chart/) repository created
* December 14, 2022: 1.99 release which is the first [including OpenCost UI in the default installation](https://www.opencost.io/blog/opencost-ui) and included initial Alibaba Cloud support.
* February 2, 2023: CNCF On-Demand Webinar: [Cloud Cost Monitoring with OpenCost and Kubernetes Optimization Strategies](https://community.cncf.io/events/details/cncf-cncf-online-programs-presents-cncf-on-demand-webinar-cloud-cost-monitoring-with-opencost-and-kubernetes-optimization-strategies/)
* February 8, 2023: 1.100 release including performance enhancements with code contributions from [Grafana Labs](https://www.opencost.io/blog/grafana-labs)
* February 16, 2023: 1.0.0 release of the Helm chart
* February 27, 2023: OpenCost announced as a [FinOps Certified Solution](https://www.opencost.io/blog/finops-certified-solution)
* February 28, 2023: Kick-off for the [External Asset Costs Working Group](https://docs.google.com/document/d/1-d-Vvy1VGHW0sXKiTjTplIUEnrElIlnfMU8sUpEehlA/edit), which will bring support for services not directly managed by the Kubernetes infrastructure
* March 9-12, 2023: OpenCost booth and [presentation](https://www.slideshare.net/mattray/scale-20x-kubernetes-cloud-cost-monitoring-with-opencost-optimization-strategies) at the Southern California Linux Expo
* April 18, 2023: [Microsoft announces they have joined the OpenCost community](http://aka.ms/aks/OpenCost-AKS) with code and feature contributions and a new [Azure Billing Rate Card API supporting OpenCost](https://www.opencost.io/docs/azure-prices)
* April 18-21, 2023: OpenCost has a full-time staffed kiosk at KubeCon EU in the Project Pavilion

### Growing Community

Since joining the CNCF Sandbox, we have seen substantial growth in the OpenCost community. According to [OpenCost DevStats](https://opencost.devstats.cncf.io/) the following stats for the main [OpenCost repository](https://github.com/opencost/opencost/) have grown notably in the past year in CNCF:

* [GitHub Stars](https://opencost.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1) increased from 2000 to over 3500
* [GitHub Forks](https://opencost.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1) increased from 192 to over 400
* [Contributors](https://opencost.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=OpenCost) increased from 40 to 368
* [Committers](https://opencost.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=OpenCost) There are now 65 committers
* There have been over 800 commits to the *develop* branch since launch
* [21 of the top 50 contributors are not Kubecost employees](https://opencost.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20decade&var-metric=contributions&var-repogroup_name=OpenCost&var-country_name=All&var-companies=All)
* [Number of new PRs](https://opencost.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&editPanel=1&from=now-1y&to=now&var-period=w&var-repogroup_name=OpenCost) now average **10 per week**
* [Median PR Time to Engagement](https://opencost.devstats.cncf.io/d/10/pr-time-to-engagement?orgId=1&var-period=d7&var-repogroup_name=OpenCost&from=now-90d&to=now): < 24 hours
* [Median PR Open to Merge](https://opencost.devstats.cncf.io/d/16/opened-to-merged?orgId=1&var-period=d7&var-repogroup_name=OpenCost&from=now-90d&to=now): < 24 hours
* [LinkedIn followers](https://www.linkedin.com/showcase/opencost/): have increased from 70 to 176
* There are now 528 Slack members in #opencost in [CNCF Slack](https://slack.cncf.io/)

We actively host or take part in various community programs and events, tracked on the [OpenCost Community Calendar](https://bit.ly/opencost-calendar). The fortnightly [OpenCost Working Group Community Meeting](https://docs.google.com/document/d/1JFB_-sjF8F9UWet1c-dWixdMZY4hri23UlQG5FX5xfY/edit#), where community and development progress are discussed. Community members and external stakeholders are invited to discuss the project and their experiences with OpenCost to shape future development.

### Expanding Developer Ecosystem

OpenCost has documented our [contributing guidelines](https://github.com/opencost/opencost/blob/develop/CONTRIBUTING.md) and [security policy](https://github.com/opencost/opencost/blob/develop/SECURITY.md).

#### Repositories

There are currently 3 OpenCost repositories:

* [opencost/opencost](https://github.com/opencost/opencost/): The primary OpenCost repository containing the [OpenCost specification](https://github.com/opencost/opencost/blob/develop/spec) and the code for the cost-model and opencost-UI.
* [opencost/opencost-helm-chart](https://github.com/opencost/opencost-helm-chart/): Helm chart for deploying and customizing OpenCost and the OpenCost UI.
* [opencost/opencost-website](https://github.com/opencost/opencost-website/): This repository hosts the [OpenCost.io](https://opencost.io) website, documentation, and blog.

#### Versioning and Releases

OpenCost repositories follow the [Semantic Versioning 2.0.0](https://semver.org/) convention. So far we have 9 OpenCost releases, with the initial OpenCost 1.94 released in June 2022, and the current release is 1.103 from May 2023. The 1.x series are considered stable and have been in production for years.

The release cadence is:
* There have been no additional major releases since OpenCost was added to the CNCF.
* Minor versions are released every month.
* Patch versions are released as necessary.

The releases are announced using [GitHub’s releases](https://github.com/opencost/opencost/releases), followed by Slack (CNCF/#opencost), Mastodon, Twitter, and discussion in the fortnightly Community Meetings.

#### Code Quality and Testing

All incoming Pull Requests are tested with GitHub Actions to execute the unit tests defined in the test and pkg/* directories against amd64 and arm64 platforms with Docker. Code formatting is enforced with go fmt. Kubecost provides integration testing and performance testing as part of their release process.

The Helm chart is tested with the GitHub Action helm/chart-testing-action which provides linting with ct lint and integration testing with ct install against a Kubernetes in Docker cluster.

The OpenCost website creates temporary sites on approved PRs with Vercel to evaluate the changes.

All merged code requires the approval of at least 1 reviewer.

#### Issue Tracking

OpenCost relies on GitHub Issues for tracking bugs, feature requests, and security issues. We actively use labels to distinguish between types of issues. Issues such as bugs or enhancement requests are also correspondingly labeled so that they are prioritized as part of the minor version releases or patch releases.

Issues are triaged daily.

### Technical Advances

Since joining the CNCF as a Sandbox Project in June of 2022, there have been monthly releases of OpenCost on GitHub. Highlights from the features, bug fixes and performance enhancements include:
* Support for GCP spot instances was added in the 1.95 release.
* Additional AWS regions were added in 1.96.
* The 1.97 release included initial support for European cloud provider ScaleWay.
* Kubernetes 1.25 support was included with 1.98.
* The OpenCost UI was included in the default installation starting with version 1.99.
* Version 1.99 also included initial support for China-based cloud provider Alibaba Cloud.
* 1.0.0 release of the OpenCost Helm chart for expanded installation options.
* 1.103 brought enhanced Azure billing integration and a daily CSV export of allocation data.

### Documentation Improvements
Documentation currently exists in both the opencost and opencost-website repositories, but it is presented on http://opencost.io. Since joining the CNCF the website is now in the OpenCost GitHub organization and is aligned with the CNCF’s IP policy to ensure all documentation is under the Creative Commons Attribution 4.0 International License. There have been many updates to the documentation, highlights include:
* Installation, upgrade, uninstall and [Helm Chart instructions](https://www.opencost.io/docs/install) have all been greatly expanded.
* A [swagger.json](https://raw.githubusercontent.com/opencost/opencost/develop/docs/swagger.json) OpenAPI document is now provided and API documentation was added to the OpenCost website. This should make API-based integrations much simpler.
* Documentation for custom [on-premises pricing](https://www.opencost.io/docs/on-prem) is now available.
* Configuring the [AWS Spot Instance Data Feed](https://www.opencost.io/docs/AWS#aws-spot-instance-data-feed-configuration) is now documented.

## Incubation Stage Requirements

### **_Adhere to CNCF [IP Policy](https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy)_**

As an existing Sandbox project this is already in place. All code is licensed Apache v2 and all documentation is Creative Commons Attribution 4.0.

### **_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope_**
A list of public adopters can be found [here](https://github.com/opencost/opencost/blob/develop/ADOPTERS.MD). Additional adopters will be added shortly as this list was recently added.

### **_Clear versioning scheme & release methodology_**

As previously documented, OpenCost follows Semantic Versioning 2.0.0 and releases monthly.

### **_Demonstrate a substantial ongoing flow of commits and merged contributions._**

As previously documented, [https://opencost.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=OpenCost)](OpenCost currently has 368 contributors and 65 committers).

Full details can be found at [CNCF devstats](https://opencost.devstats.cncf.io/) dashboards for OpenCost.

### **_Security_**

The security process for OpenCost is documented in the [SECURITY.md](https://github.com/opencost/opencost/blob/develop/SECURITY.md). Here is an abbreviated summary:
* All code is version-controlled with known committers.
* Dependencies are checked by Dependabot.
* The current and previous releases are supported for security fixes.
* Vulnerability reporting is handled through GitHub-provided private vulnerability reporting.
* Vulnerabilities will be disclosed once confirmed and a public fix released.

### **_Specifications must have at least one public reference implementation._**

The Golang OpenCost project is the implementation of the [OpenCost specification](https://github.com/opencost/opencost/tree/develop/spec).

## Statement on Alignment with the CNCF Mission

OpenCost aligns with the mission of the CNCF by empowering organizations to monitor the cost of their cloud native environments and to make informed decisions on their optimizing their spend and application infrastructure.

### Alignment with Other CNCF Projects

OpenCost is currently aligned with the following CNCF projects:
* Helm: OpenCost has an opencost-helm-chart repository that provides a Helm chart for deploying and upgrading OpenCost installations.
* Kubernetes: OpenCost is installed on Kubernetes and relies on Kubernetes’ APIs for aligning Kubernetes resources with allocated cloud costs.
* Prometheus: In default installations OpenCost stores metrics in Prometheus’ time series database and uses the Prometheus node-exporter to supplement monitoring data.
* Cortex: Storage option used as an alternative to Prometheus, supported by the OpenCost Helm chart.
* Thanos: Storage option used as an alternative to Prometheus, supported by the OpenCost Helm chart and used by Kubecost with their commercial implementation.

## Future Plans

The [External Asset Costs Working Group](https://docs.google.com/document/d/1-d-Vvy1VGHW0sXKiTjTplIUEnrElIlnfMU8sUpEehlA/edit) intends to bring support for services not directly managed by the Kubernetes infrastructure. Examples include network costs, monitoring, and external database services. Supporting AWS’ S3 object storage is the prototype under development with documentation and a new specification being developed in parallel. These efforts should encourage external contributors from both vendors and users. Initial patches from Grafana Labs and Microsoft are refactoring cloud providers to make it easier to introduce new billing and data sources.

The FinOps Foundation is participating in the development conversations as OpenCost looks to standardize on the [FinOps Open Cost and Usage Specification](https://www.finops.org/focus/) (FOCUS). Initial FOCUS patches have already been merged.

There is a [ROADMAP.md](https://github.com/opencost/opencost/blob/develop/ROADMAP.md) and issues are tagged with "enhancement" to signify potential new features for consideration.
