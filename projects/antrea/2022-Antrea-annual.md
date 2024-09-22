# Antrea 2022 Sandbox Annual Review

[Antrea](https://antrea.io) is a Kubernetes networking solution intended to be Kubernetes native. It
operates at Layer3/4 to provide networking and security services for a Kubernetes cluster,
leveraging Open vSwitch as the networking data plane.

Antrea was launched in November 2019, and was accepted into the CNCF as a Sandbox project in April
2021.

## Project updates

Here are some of the most significant updates for the project:
* Antrea is used on Windows Nodes in GKE (Google Kubernetes Engine) and Google
  [Anthos](https://cloud.google.com/anthos/clusters/docs/on-prem/latest/how-to/user-guide-for-windows-server-os-node-pools)
  to enforce K8s NetworkPolicies.
* Antrea is used in [Tanzu Community Edition](https://github.com/vmware-tanzu/community-edition).
* The Antrea community is running a weekly live show on our Youtube channel to explore different
  topics related to K8s and to container networking: [Antrea
  LIVE](https://www.youtube.com/watch?v=vJ49_BniUwI&list=PLuzde2hYeDBfHDD0zMbmG4QoVaSbkJChZ).
* The community has added many significant features over the last year. One notable example is
  "multi-cluster" support: the ability to connect different K8s clusters running Antrea to provide
  network connectivity, policy enforcement, and service access.
* We started the [Theia](https://github.com/antrea-io/theia) sub-project, an observability (flow
  visibility) solution for Antrea.
* The [website](https://antrea.io) is now automatically updated (by a bot) with each new release, to
  ensure that the user-facing documentation is always up-to-date.

## Metrics

The project's devstats page can be found [here](https://antrea.devstats.cncf.io/).
The numbers of contributions and contributors have remained consistent over the last year, and even
show a slight upward trend since the beginning of 2022.
The project has just reached 100 individual contributors (commits) on Github, even though a
significant number of them only have one or two commits. In the last year, we have had contributions
from different companies, including Intel, Google, and Transwarp Technology.
The "PR Time to Engagement" has also been consistent, with an average of 8 hours over the last year.

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
active in the Antrea community and who have agreed to add their company name and logo to the file.

The 2021 CNCF survey shows that among 518 respondents, 29 indicate that their company/organization
uses Antrea in production, and 53 indicate that they are evaluating Antrea. A significant portion of
these users are based in Asia. The survey data is available [on
Github](https://github.com/cncf/surveys/blob/main/cloudnative/Cloud_Native_Survey_2021-Part_2.xlsx).

While these are not "end users", we are also aware of the following usage of Antrea:
* Antrea is used in GKE and Google Anthos to enforce K8s NetworkPolicies on Windows Nodes.
* Antrea is used as the default CNI in VMware K8s distributions, both commercial and OSS (see [Tanzu
  Community Edition](https://github.com/vmware-tanzu/community-edition)).

## Performance against previous goals

This is our first annual review as a CNCF Sandbox project.
One of our objectives when applying to join the CNCF was to increase visibility for the projects. We
think that this objective has been partially met: some users seem to discover the project only
because it is part of the CNCF, but at the same time we are struggling to increase the number of
regular contributors outside of VMware.

## Current goals

While we are always working on adding new features, this year we would like to concentrate on the
following goals:
* increasing contributions from outside of VMware, and if possible add one new maintainer from
  outside of VMware
* improve usability of Theia, the observability (flow visibility) solution for Antrea
* improve stability of existing Alpha & Beta features
* graduate stable APIs to GA

## How can the CNCF help you achieve your upcoming goals?

CNCF staff has been very responsive when we had enquiries in the past. We do not require any
specific help at this stage.

## Do you think that your project meets the criteria for incubation?

Yes, we would like to be considered for incubation. We think we meet the necessary
[criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage):
* Adoption: the CNCF TOC is currently replacing the "three end users" requirement with a "three
  direct adopters" requirement as part of this [PR](https://github.com/cncf/toc/pull/864). Based on
  this update, we believe that we show a sufficient level of [adoption](#Adoption).
* Have a healthy number of committers: at the moment we have 4 committers for the main Antrea
  repository (who are the 4 maintainers), as well as 2 additional committers for the
  [Theia](https://github.com/antrea-io/theia) sub-project. As the project grows, we plan to add area
  owners who can accept contributions to specific parts of the Antrea project.
* Substantial ongoing flow of contributions: can be seen from
  [devstats](https://antrea.devstats.cncf.io/d/74/contributions-chart?orgId=1).
* Clear versioning scheme: Antrea has followed a clear semantic versioning scheme since before
  release v1.0.0, which is [documented](https://github.com/antrea-io/antrea/blob/main/docs/versioning.md).
* Clearly documented security processes explaining how to report security issues to the project: see
  our [documentation](https://github.com/antrea-io/antrea/blob/main/SECURITY.md).
* Specifications must have at least one public reference implementation: N/A

In addition to the above, Antrea maintains a Core Infrastructure Initiative [Best Practices
Badge](https://bestpractices.coreinfrastructure.org/en/projects/4173).

If the TOC feels that Antrea is not ready for incubation, we would appreciate feedback about the
areas in which we need to improve. Thanks!
