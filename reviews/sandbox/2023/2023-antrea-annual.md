# Antrea 2023 Sandbox Annual Review

[Antrea](https://antrea.io) is a Kubernetes networking solution intended to be Kubernetes native. It
operates at Layer3/4 to provide networking and security services for a Kubernetes cluster,
leveraging Open vSwitch as the networking data plane.

Antrea was launched in November 2019, and was accepted into the CNCF as a Sandbox project in April 2021.

## Project updates

Here are some of the most significant updates for the project:

* The community has added many significant features over the last year. Notable examples include:
    - Implementation of the new K8s upstream [`AdminNetworkPolicy` API](https://network-policy-api.sigs.k8s.io/api-overview/).
    - Introduction of a new `L7NetworkPolicy` API to enforce network policies at the application layer (L7).
    - ExternalNode feature which enables Antrea to manage and enforce network policies on
      non-Kubernetes Nodes (Virtual Machines and bare-metal servers running the Antrea Agent).
* We started the [Nephe](https://github.com/antrea-io/nephe) sub-project, which extends the Antrea
  NetworkPolicy model to Virtual Machines in the public cloud (AWS, Azure).
* We added a dedicated [UI](https://github.com/antrea-io/antrea-ui) for Antrea, following user
  feedback and the deprecation of the [Octant](https://github.com/vmware-archive/octant) project
  (prior to this, Antrea shipped with an Octant plugin).

## Metrics

The project's devstats page can be found [here](https://antrea.devstats.cncf.io/).
The numbers of contributions and contributors have remained consistent over the last year, with
upwards of 40 monthly contributors on average.
The "PR Time to Engagement" has also been consistent, with an average median time of 8 hours over
the last year (there was a slight improvement compared to our last annual review).

## Maintainers

The [MAINTAINERS](https://github.com/antrea-io/antrea/blob/main/MAINTAINERS.md) file is up-to-date,
and shows 4 current maintainers, all from the same company (VMware):

```
| Maintainer | GitHub ID | Affiliation |
| ---------- | --------- | ----------- |
| Antonin Bas | antoninbas | VMware |
| Jianjun Shen | jianjuns | VMware |
| Quan Tian | tnqn | VMware |
| Salvatore Orlando | salv-orlando | VMware |
```

## Adoption

We have an [ADOPTERS](https://github.com/antrea-io/antrea/blob/main/ADOPTERS.md) file in our
repository, which shows 3 users: Glasnostic, Transwarp, TerraSky. These are users who have been
active in the Antrea community and who have agreed to add their company name and logo to the
file. We have not had any addition to this file since our last review. VMware is also using Antrea
as the networking solution for their K8s-aaS developer platform, which runs production workloads on
large clusters (100+ Nodes).

The Antrea Docker image (Linux + Windows) has been pulled 1.7M times from docker.io in the last
year, which is a clear increase compared to the previous year (0.6M).

While these are not end users, we are also still aware of the following usage of Antrea:
* Antrea is used in GKE and Google Anthos to enforce K8s NetworkPolicies on Windows Nodes.
* Antrea is used as the default CNI in VMware commercial K8s distributions (Tanzu portfolio).

## Performance against previous goals

While this is hard to quantify, we believe that the number of end users of Antrea is increasing. We
see an healthy amount of user issues and user feature requests, with significant interest in our
Egress feature.

We are still struggling to increase the number of *regular* contributors outside of VMware, and
hence have not met our earlier goal of adding a new maintainer from outside of VMware.

We believe we have made good progress on our other stated goals from last year's review: improve
usability of Theia (observability solution), improve feature and API stability. There is still
ongoing work on the latter, as we march towards an Antrea v2.0 release.

## Current goals

While we are always working on adding new features, this year we would like to concentrate on the following goals:

* increase contributions from outside of VMware, and if possible add one new maintainer from outside
  of VMware. This was already our goal last year, and is still a challenge for us.
* Antrea v2.0 release: with this milestone, we want to underline the stability and maturity of the
  project. This new major release will be the opportunity for us to promote our major APIs and
  Feature Gates to GA, and remove some deprecated APIs.
* work on these specific features from our roadmap: L7 NetworkPolicy (improvements to the API and
  implementation), L7 observability in Theia, focus on Telco workloads with native support for
  secondary network interfaces.
* increase e2e test quality, and reduce the frequency of flakes, which are a distraction / annoyance
  for contributors.

## How can the CNCF help you achieve your upcoming goals?

CNCF staff has been very responsive when we had enquiries in the past. We are thankful for the help
in transitioning our website to a CNCF-managed Netlify account. We do not require any specific help
at this stage.

## Do you think that your project meets the criteria for incubation?

Yes, we would like to be considered for incubation.
We think we meet the necessary [criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage):

* Adoption: we believe that we show a sufficient level of [adoption](#Adoption).
* Have a healthy number of committers: at the moment we have 4 committers for the main Antrea
  repository (who are the 4 maintainers), as well as 2 additional committers for the
  [Theia](https://github.com/antrea-io/theia) sub-project, and 1 additional committer for the
  [Nephe](https://github.com/antrea-io/nephe) sub-project.
* Substantial ongoing flow of contributions: can be seen from [devstats](https://antrea.devstats.cncf.io/d/74/contributions-chart?orgId=1).
* Clear versioning scheme: Antrea has followed a clear semantic versioning scheme since before
  release v1.0.0, which is  [documented](https://github.com/antrea-io/antrea/blob/main/docs/versioning.md).
* Clearly documented security processes explaining how to report security issues to the project: see
  our [documentation](https://github.com/antrea-io/antrea/blob/main/SECURITY.md).
* Specifications must have at least one public reference implementation: N/A

In addition to the above, Antrea maintains a Core Infrastructure Initiative [Best Practices Badge](https://bestpractices.coreinfrastructure.org/en/projects/4173).

If the TOC feels that Antrea is not ready for incubation, we would appreciate feedback about the areas in which we need to improve. Thanks!
