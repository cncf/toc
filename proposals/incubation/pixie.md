# Pixie for incubation

## Table of contents 

- [About Pixie](#about-pixie)
- [Sponsor / Advisor from TOC](#sponsor--advisor-from-toc)
- [Adoption](#adoption)
- [Development Process](#development-process)
- [Relationships with other projects](#relationships-with-other-projects)
- [Roadmap](#roadmap)


## About Pixie
Pixie is an open source observability tool for Kubernetes applications. Leveraging technology such as eBPF, Pixie automatically collects telemetry data, including full-body requests, resource and network metrics, application profiles, and more. This allows developers to view the high-level state of their cluster (service maps, cluster resources, application traffic) and also drill down into more detailed views (pod state, flame graphs) without having to modify or redeploy their code.

All telemetry data is collected and stored locally in the cluster, utilizing less than 5% of the cluster's CPU. This edge architecture enables users to deploy Pixie in sensitive environments and air gapped clusters. 

Data can be processed using Pixie's flexible Pythonic query language, Pxl, across Pixie's UI, CLI, and APIs. Pixie additionally supports a plugin system which allows users to easily export its telemetry data and integrate with other projects using OpenTelemetry.  

Pixie was accepted as a CNCF sandbox project in June 2021. Since our inclusion in the CNCF, Pixie has seen promising community contribution, growth, and adoption. Companies such as New Relic and VMWare are building products and integrations on top of the Pixie platform, while others, such as Verizon, use Pixie to monitor their Kubernetes clusters. Pixie also has an active and engaged community, with 1300+ Slack members, 50+ contributors, and an average of 245 commits per month. 


## Sponsor / Advisor from TOC

To be assigned by TOC


## Adoption

Pixie has broad adoption across thousands of production systems (see [ADOPTERS.md](https://github.com/pixie-io/pixie/blob/main/ADOPTERS.md) for some public references).

Some other prominent examples include:


* AWS
    - [Amazon EKS Blueprints Quick Start](https://aws-quickstart.github.io/cdk-eks-blueprints/addons/pixie/)
    - [Pixie - Amazon Managed Grafana](https://docs.aws.amazon.com/grafana/latest/userguide/ds-pixie.html) 
    - [EKS Workshop: Monitoring using Pixie](https://www.eksworkshop.com/intermediate/241_pixie/)
* New Relic
    - [Pixie is Now Available as Part of New Relic One](https://newrelic.com/blog/how-to-relic/kubernetes-monitoring-pixie-ga)
* Verizon
    - [Verizon takes innovation to the edge with Pixie and New Relic](https://newrelic.com/customers/verizon?utm_campaign=PostBeyond&utm_source=Twitter&utm_medium=%23362600&utm_term=Verizon)

We can also measure production-use through tools and addons which the community has built on top of Pixie:

* [Cloudwatch AutoInstrumentation with Pixie](https://github.com/anshrma/pixie-autoinstrumentation-with-cloudwatch) 
* [ChaosToolkit Pixie Driver](https://chaostoolkit.org/drivers/pixie/)

We also noticed an uptick on our [Slack](https://slackin.px.dev/) from folks trying out Pixie, asking some questions and using Pixie across their clusters. Our Slack workspace now has 1300+ members.

Our conference presence has also been well-received. We have had significant attendance at our booth and talks at Kubecon EU 2022, and Kubecon NA 2022.



* [Autoscaling Kubernetes Deployments: A (Mostly) Practical Guide - Natalie Serrino](https://www.youtube.com/watch?v=n8t_hbchQcc)
* [Reproducing Production Issues in your CI Pipeline Using eBPF - Matthew LeRay & Omid Azizi](https://www.youtube.com/watch?v=_RQLY4KXXG8)
* [Tracing SSL/TLS Encrypted Microservices with eBPF - Dom Del Nano](https://sched.co/1Auyt)
* [When the Logs Just Don’t Cut It: Root-Causing Incidents Without Re-Deploying Prod - Phillip Kuznetsov](https://sched.co/182IS)
* [Flux + Observability: Featuring Prometheus Operator and Pixie - Somtochi Onyekwere](https://sched.co/1AR9E)


## Development Process


### Versioning

Pixie has several releasable components. The main ones are: control plane, vizier (monitoring backend), K8s operator, Go/Python APIs, and the CLI. 

The control plane utilizes a date-based release scheme, with releases generally happening once a week. We utilize a date-based scheme here since this component is typically managed by automated deployment systems. 

Everything else in Pixie is released using semantic versioning. These components also follow a regular weekly release cadence.


### Security

Given the scope of Pixie’s data collection capabilities, security and vulnerabilities are taken very seriously by the Pixie team. Pixie has a [document](https://github.com/pixie-io/pixie/blob/main/SECURITY.md) describing the security reporting process. In addition, we have performed pen tests and vulnerability scans.


### Project Health

Pixie has a [CLOMonitor score](https://clomonitor.io/projects/cncf/pixie) of 100 and passes [OpenSSF best practices](https://bestpractices.coreinfrastructure.org/en/projects/5027). Pixie also has an A+ on [Go Report Card](https://goreportcard.com/report/px.dev/pixie).


### Repositories

[pixie-io/pixie](https://github.com/pixie-io/pixie): The main repository containing all the code for the control plane, Vizier (monitoring backend), operator, and CLI.

[pixie-io/pxapi.go](https://github.com/pixie-io/pxapi.go): Code for Pixie’s Go API, along with examples.

[pixie-io/pixie-plugin](https://github.com/pixie-io/pixie-plugin): Repository containing all available plugins for Pixie’s Plugin system, which allows users to send Pixie data to other tools using OTLP.

[pixie-io/px.dev](https://github.com/pixie-io/px.dev): The implementation for Pixie’s website: px.dev.

[pixie-io/blog.px.dev](https://github.com/pixie-io/blog.px.dev): Content and code for Pixie’s blog: blog.px.dev.

[pixie-io/docs.px.dev](https://github.com/pixie-io/docs.px.dev): Documentation for how to install and use Pixie: docs.px.dev.


### Issue Tracking

Pixie uses Github Issues for tracking bugs, features, feature requests, and more. Labels are used to denote the type of issue (for example, `kind/bug` vs `kind/feature`, for bugs and feature requests respectively). Labels are also used to indicate where the issue is in the triage process (`triage/*`), and the priority of the task (`priority/*`).

Issues are triaged daily.


### Release Process

Pixie has a weekly release process for each of its components. When a component is ready to be released, a release candidate is made by creating a branch off of `main`. The release candidate branch is pushed to Github, triggering a Jenkins job to build any necessary images and artifacts. Once the build is complete, testing begins (see “Testing Process” below). If testing succeeds, a new branch is made for the release to trigger a Jenkins build for the official release. All releases, changelogs, and release notes are published to the [Version History](https://docs.px.dev/reference/admin/product-updates/) page on px.dev.


### Testing Process

All changes that have been pulled into `main` are carefully unit tested and reviewed. All PRs and updates to PRs trigger a test build in Jenkins. This test build runs through all related unit tests and linters. When a PR is merged in, another build is also triggered to run tests on the entire repository.

Before a release, each release candidate also goes through a round of manual testing. Each component has its own checklist for test cases which cannot be easily automated. Maintainers are responsible for running through the items in the checklist using the release candidate. Once the checklist is complete, the release is triggered.



* [Vizier Checklist](https://docs.google.com/spreadsheets/d/1GOmMu-1T2lAW-LXkDwwz4N0s_l07sQDDWTdXRgFX6wA/edit#gid=0)
* [Operator Checklist](https://docs.google.com/spreadsheets/d/1KVM5fvipz-td7Dm6HHd7T2rD8di9m3vnaxIJ6QGrDXE/edit#gid=0)
* [Control-plane Checklist](https://docs.google.com/spreadsheets/d/1nuDaDN1TqnVL9GKWGe-Wd2N-12c8ZpktfUHZSma2c3I/edit)
* [CLI Checklist](https://docs.google.com/spreadsheets/d/1DPDdjf7GK4q30V7e5hBng7ljz6Ps7_vIqwoU6FINn54/edit#gid=0)


### External Deps

Pixie relies on other great open source libraries and tools, including and not limited to:



* [NATS/Jetstream](https://nats.io/): Pixie uses NATS and Jetstream as its message bus to send messages and data throughout the system. 
* [OpenTelemetry](https://opentelemetry.io/): Pixie uses OTLP as the format to export Pixie data to other tools in Pixie’s Plugin System.
* [etcd](https://etcd.io/): Although Pixie does not persist any collected data, Pixie does need a datastore for metadata. Users running clusters who may not have support for Persistent Volumes can opt to deploy Pixie with etcd as a reliable in-memory datastore.
* [Operator Lifecycle Manager](https://olm.operatorframework.io/): OLM helps deploy Pixie to a user’s cluster, and ensures that the operator remains running and up-to-date.

A more comprehensive list of dependencies can be found [here](https://storage.googleapis.com/pixie-dev-public/oss-licenses/latest.json).


### Contributions

Pixie has a core team of 10+ full-time committers focused on improving the project and adding more capabilities to Pixie.

Pixie additionally welcomes and encourages contributions from the wider community. These contributions range from small enhancements and bug fixes, to larger contributions, such as protocol tracers. A list of Pixie authors can be found in the [AUTHORS](https://github.com/pixie-io/pixie/blob/main/AUTHORS) file. Many of these contributions also come from vendor offerings (Seekret, GroundCover, etc.) which use Pixie as a foundational building block.

At the time of writing, Pixie has:



* [50+ total contributors](https://github.com/pixie-io/pixie/graphs/contributors) 
* [2933 commits during the last 12 months, with an average of 245 commits per month.](https://pixie.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all)
* [More than 4.2k stars and 400+ forks on Github](https://pixie.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1)

**Maintainers List**


<table>
  <tr>
   <td><strong>Name</strong>
   </td>
   <td><strong>Organization</strong>
   </td>
  </tr>
  <tr>
   <td>Zain Asgar
   </td>
   <td>New Relic / Stanford University
   </td>
  </tr>
  <tr>
   <td>Michelle Nguyen
   </td>
   <td>New Relic
   </td>
  </tr>
  <tr>
   <td>Vihang Mehta
   </td>
   <td>New Relic
   </td>
  </tr>
</table>


**Community Meetings**

Pixie holds monthly community meetings at 9am on the first Tuesday of every month. These meetings are intended to share updates about Pixie (new features, announcements), and serve as a forum for the community to interact and ask questions.

Pixie also holds a monthly governance board meeting at 9am on the first Thursday of every month. This meeting is limited to Pixie’s board members to discuss and make decisions about the project’s direction and roadmap.

**Social Media**

Pixie is active on Twitter as [@pixie_run](https://twitter.com/pixie_run) with 1,300+ followers. 

We also have a [Youtube channel ](https://www.youtube.com/channel/UCOMCDRvBVNIS0lCyOmst7eg)with 50+ videos and 300+ subscribers.

**Websites**

[px.dev](https://px.dev/): The main website explaining what Pixie is and links out to other helpful resources for getting started with Pixie.

[docs.px.dev](http://docs.px.dev): Documentation for how to install and use Pixie, for both beginners and advanced users.

[blog.px.dev](http://blog.px.dev): Pixie’s blog containing announcements and technical posts written by both Pixie contributors and guests. Our blog averages 14k+ page views a month.


## Relationships with other projects

Pixie believes in CNCF’s mission to make cloud-native technologies accessible for all. It should be easy and free for developers to gain visibility into their systems. In order to accomplish these goals, Pixie asserts that the [future of observability lies in open source](https://blog.px.dev/plugin-system/). Open source projects give all developers access to powerful new tools. Platforms should also leverage open standards to make it simple for users to flexibly incorporate these tools into their everyday workflows and techstacks.

Pixie’s plugin system utilizes open standards so developers can effortlessly integrate Pixie into their workflows. The plugin system allows users to export their Pixie data to any other tool which accepts OTLP data. This enables users to send data to Jaeger, Prometheus, and more. The use of open standards allows developers to take advantage of each tool’s strengths in conjunction. For example, similar to Prometheus, Pixie can scrape, store, and query time-series data. However, Pixie focuses on the real-time debugging use case and does not provide long-term data retention. Pixie’s plugin system offers access to Pixie’s rich dataset alongside Prometheus’s powerful long-term data store.


## Roadmap

Pixie maintains a steady pace of releases driven by a growing set of requirements and use cases. As Pixie is actively used in thousands of production clusters, we are focusing on the following to ensure Pixie operates smoothly across all environments and workflows:



* Expand integration with other tools
    * Leverage the Pixie Plugin system to enable users to import/export data into other tools. This helps users easily incorporate Pixie into their everyday workflows and standard tooling.
    * Add support for more protocols to integrate with other tools.
* Improve compatibility, protocol support, and robustness
    * Expand matrix of supported K8s distributions, kernel versions, CPU Architectures (ARM), and related deprecation policies.
    * Expand our testing matrix to automatically cover supported environments. 

Pixie’s detailed roadmap is located [here](https://docs.google.com/spreadsheets/d/1tS9BxGaiu8P2Iq-ucxhicyteOH2hjVDRR80jCrsBANw/edit#gid=0).
