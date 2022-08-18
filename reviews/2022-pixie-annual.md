# Pixie 2022 Annual Review

## Background

Pixie is an open source observability tool for Kubernetes applications. Leveraging technology such as eBPF, Pixie automatically collects telemetry data, including full-body requests, resource and network metrics, application profiles, and more. This allows developers to view the high-level state of their cluster (service maps, cluster resources, application traffic) and also drill down into more detailed views (pod state, flame graphs). All telemetry data is collected and stored locally in the cluster, utilizing less than 5% of cluster CPU. This data can be processed using Pixie's flexible Pythonic query language, Pxl, across Pixie's UI, CLI, and APIs. 

Pixie was accepted as a CNCF sandbox project in June 2021. Since our inclusion in the CNCF, Pixie has seen promising community contribution and adoption.

## DevStats

Pixie has seen a steady growth of Github activity. The [steady increase in the number of stars](https://pixie.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now) indicates a growing adoption of Pixie from the community.

Pixie has a core team of 10+ full-time contributors that are focused on improving the product and adding new capabilities. Although the majority of our contributors are from New Relic, we are currently working on growing active contributors across other companies. We are [seeing more contributions](https://pixie.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=d7&var-metric=contributions&from=now-1y&to=now) of various sizes from the community. This includes significant contributions (such as new protocols implemented in eBPF) from various organizations. In addition, Pixie is [adopted](https://github.com/pixie-io/pixie/blob/main/ADOPTERS.md) and used by various cloud providers.
 
Overall, Pixie has over 50 contributors to date. 

## Maintainers

| Maintainer       | GitHub Username                               | Company     |
| ---------------- | --------------------------------------------- | ----------- |
| Zain Asgar       | [@zasgar](https://github.com/zasgar)          | New Relic   |
| Michelle Nguyen  | [@aimichelle](https://github.com/aimichelle)  | New Relic   |
| Vihang Mehta     | [@vihangm](https://github.com/vihangm)        | New Relic   |

Our [governance policy](https://github.com/pixie-io/pixie/blob/main/GOVERNANCE.md) is managed in our repository.

## Adoption
 
We are seeing a large uptick in community participation and adoption. Our conference presence has been well-received and we had an overflowing room for our talks at Kubecon EU 2022. 
We also noticed an uptick on our [Slack](https://slackin.px.dev/) from folks trying out Pixie, asking some questions and using Pixie across their clusters.
 
Some users have implemented tools built on top of Pixie. For example:

- [Cloudwatch AutoInstrumentation with Pixie](https://github.com/anshrma/pixie-autoinstrumentation-with-cloudwatch) 
- [ChaosToolkit Pixie Driver](https://chaostoolkit.org/drivers/pixie/)
 
There are also a handful of large companies that have started adopting Pixie and using it widely. See our [Adopters file](https://github.com/pixie-io/pixie/blob/main/ADOPTERS.md).

## Current Goals

This is Pixie's first annual review, so have no previously documented goals to reference.

Pixie is actively used in production by thousands of clusters, this means that we will be spending significant time and resources on expanding compatibility, security and enterprise readiness, and growing our community. These are our high level goals:
 
- Expand integration with other tools:
  - Leverage the Pixie Plugin system to enable users to import/export data into other tools.
  - Add support for a variety of protocols to integrate with various tools..
- Grow our community:
  - Increase the number of Pixie maintainers.
  - Increase active contributors to a healthy self sustaining number.
  - Enable adoption of Pixie for a variety of use cases.
- Expand compatibility, protocol support and robustness:
  - Expand matrix of supported K8s distributions, kernel versions, CPU Architectures (ARM), and related deprecation policies.
  - Expand our testing matrix to automatically cover supported environments. 
 
Pixie’s detailed roadmap is located [here](https://docs.google.com/spreadsheets/d/1tS9BxGaiu8P2Iq-ucxhicyteOH2hjVDRR80jCrsBANw/edit#gid=0).

## How the CNCF Can Help

Getting access to the community through CNCF (for example, booths and talks at KubeCon) has been instrumental to helping us understand our user base and refining our roadmap. We would like more opportunities to interact with the community, such as events, webinars, and more.

## Incubation 

We feel that Pixie is ready to apply for incubation. We are a cross-vendor project with broad adoption across thousands of production systems. Pixie also has a healthy user and contributor community. 

### Adoption

Pixie has several adopters, some of which are listed on our [adopters page](https://github.com/pixie-io/pixie/blob/main/ADOPTERS.md). Many other adopters of Pixie aren’t listed, but here are some press releases talking about their use of Pixie.
  
AWS:

- [Amazon EKS Blueprints Quick Start](https://aws-quickstart.github.io/cdk-eks-blueprints/addons/pixie/)
- [Pixie - Amazon Managed Grafana](https://docs.aws.amazon.com/grafana/latest/userguide/ds-pixie.html) 
- [EKS Workshop: Monitoring using Pixie](https://www.eksworkshop.com/intermediate/241_pixie/)

- [Pixie is Now Available as Part of New Relic One](https://newrelic.com/blog/how-to-relic/kubernetes-monitoring-pixie-ga)
- [Verizon takes innovation to the edge with Pixie and New Relic.](https://newrelic.com/customers/verizon?utm_campaign=PostBeyond&utm_source=Twitter&utm_medium=%23362600&utm_term=Verizon)

### Contributions

We have 10+ active committers to Pixie.

Pixie currently has around 50 contributors. There is a core team of 10+ full-time contributors that are focused on improving the project and adding more capabilities to Pixie. In addition, we also have contributions from the community. Community contributions range from small enhancements to large contributions. Pixie is also used as a foundational building block for various other vendor offerings (Seekret, GroundCover, etc.), and they are actively engaged in adding significant new capabilities to Pixie.

### Versioning

Pixie has several releasable components. The main ones are: control plane, vizier (monitoring backend), K8s operator, Go/Python APIs, and the CLI. 
 
The control plane utilizes a date-based release scheme, with releases typically happening once a week. We utilize a date-based scheme here since this component is typically managed by automated deployment systems. 
 
Everything else in Pixie is released using semver.

### Security

Pixie has a [document](https://github.com/pixie-io/pixie/blob/main/SECURITY.md) describing the security reporting process. In addition, we have performed pen tests and vulnerability scans.


