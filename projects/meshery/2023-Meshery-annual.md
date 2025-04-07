# Meshery Sandbox Project Annual Report

As a self-service engineering platform, Meshery enables collaborative design and operation of cloud native infrastructure. Meshery is an extensible, cloud native management platform, offering multi-cluster management, lifecycle, configuration, and performance management of Kubernetes clusters and their workloads.

## Devstats ([link](https://meshery.devstats.cncf.io/))

- [130 companies represented by contributors (all-time)](https://meshery.devstats.cncf.io/d/5/companies-table?orgId=1)
- [41 countries represented by contributors (all-time)](https://meshery.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=All)
- [3350 stars (all-time)](https://meshery.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=All)

| Type | Sep 2020 - Sep 2021 | Sep 2021 - Sep 2022 | Sep 2022 - Sep 2023 | Total |
| --- | --- | --- | --- | --- |
| Commits | 11.03K | 11.02K (-0.09%) | 13.5K (+22.45%) | 34k |
| Pull Requests | 2K | 1.9K (-5%) | 2.3K (+21.05%) | 6.4k |
| Contributions | 28.72K | 25.29K (-11.94%) | 31.3K (+23.73%) | 101k |
| Contributors | 177 | 127 (-28.25%) | 226 (+77.95%) | 678 |
| Forks | 306 | 305 (-0.33%) | 731 (+139.67%) | 1,615 |


## Maintainers ([link](https://github.com/meshery/meshery/blob/master/MAINTAINERS.md))

### CLI Maintainers

| Name                    | GitHub            | Affiliation |
| ----------------------- | ----------------- | ----------- |
| Hussaina Begum          | hexxdump          | VMware      |
| Aadhitya Amarendiran    | alphaX86          | Citi        |
| Jerod Culpepper         | cpepper96         | SAIC        |
| Antonette Caldwell      | acald-creator     | Acquia      |

### UI Maintainers

| Name                | GitHub                 | Affiliation |
| ------------------- | ---------------------- | ----------- |
| Abhishek Kumar      | Abhishek-kumar09       | Layer5      |
| Uzair Shaikh        | MUzairS15              | Layer5      |
| Nikhil Ladha        | Nikhil-Ladha           | Red Hat     |
| Antonette Caldwell  | acald-creator          | Acquia      |

### Adapter Maintainers

| Name                | GitHub        | Affiliation |
| ------------------- | ------------- | ----------- |
| Aisuko Li           | aisuko        | RMIT        |
| Dhiraj Gedam        | dheerajng     | Citrix      |
| Haim Helman         | thehh1974     | VMware      |
| Hussaina Begum      | hexxdump      | VMware      |
| Ashish Tiwari       | revolyssup    | API7        |
| Michael Gfeller     | mgfeller      | Computas AS |
| Antonette Caldwell  | acald-creator | Acquia      |
| Xin Huang           | gyohuangxin   | Intel       |

### CI / Build & Release Maintainers

| Name                  | GitHub             | Affiliation |
| --------------------- | ------------------ | ----------- |
| Ashish Tiwari         | revolyssup         | API7        |
| Pranav Singh          | theBeginner86      | Layer5      |
| Mario Arriaga         | MarioArriaga92     | F5          |

### Docs Maintainers

| Name              | GitHub          | Affiliation |
| ----------------- | --------------- | ----------- |
| Adithya Krishna   | adithyaakrishna | Red Hat     |
| Lee Calcote       | leecalcote      | Layer5      |

### Site Maintainers

| Name                    | GitHub       | Affiliation    |
| ----------------------  | -----------  | -------------- |
| Nikhil Ladha            | Nikhil-Ladha | Red Hat        |
| Aaditya Narayan Subedy  | asubedy      | Fast Retailing |

## Adopters ([link](https://github.com/meshery/meshery/blob/master/ADOPTERS.md))

- ARM
- Asteria Aerospace
- BootLabs
- Ericsson
- HPE Labs- Cloudless
- HPE Security Engineering
- Intel
- Intraserve Systems
- Metabit Trading
- Metabob
- Red Hat
- SolarWinds

## Project Goals

Meshery has greatly expanded beyond it's initial service mesh-centric focus to support integrations (~220 currently) across the entire cloud native ecosystem. Meshery's configuration and orchestration abilities include a growing list of supported infrastructure.

Meshery's [roadmap](https://github.com/meshery/meshery/blob/master/ROADMAP.md) has remained consistent over this past year. The community of contributors has been in execution mode, dialing now on it's final, planned architectural components to a v1.0. Signicant strides have been made with [Meshery's architecture](https://docs.meshery.io/concepts), including incorporation of WASM for OPA execution in it's UI client, a client-side user permissioning framework, delivery a [Docker Extension](https://docs.meshery.io/installation/platforms/docker-extension), the addition of support for Postgres, 6 new [extension points](https://docs.meshery.io/extensibility), self-documenting end-to-end test framework that includes a published [platform compatibility matrix](https://docs.meshery.io/installation/platforms), and a number of other features. Additionally, Meshery's community is blooming with Meshery LFX internship being the #1 most popular out of all of the LF.

## How can the CNCF help?

### Help promoting use of the [Cloud Native Playground](https://play.meshery.io)

Under Meshery's mission to enable the ease by which and the collaborative nature by which cloud native infrastructure is managed, the Meshery project has created a cloud native playground - a CNCF playground -  in which all CNCF projects can be configured and deployed. This playground is being hosted in the CNCF labs. We could use the TOC's assistance in encouraging other CNCF projects to represent their deployment best practices in the [Meshery Catalog](https://meshery.io/catalog). Infrastructure designs available in the Meshery Catalog represent patterns and templates that end users can quickly apply in their environments. Ideally, each CNCF project incorporates their deployment models and suggested best practices into the catalog.

The Meshery Catalog is available in the cloud native playground and is an excellent resources for participating CNCF End User Organizations. It is of benefit to all projects and all end users.

### Uplifting of all Helm charts in Artifact Hub

Meshery has the ability to visualize Helm charts and [screenshot the visualization](https://github.com/meshery/meshery/pull/8498#issuecomment-1681369650). We think this would be a valuable enhancement to users of Artifact Hub. Connecting the two projects for an exploratory discussion would be helpful.

### Incubation Readiness Assessment and Project Sponsors

The TOC can assist by highlighting any concerns in terms of Meshery's readiness for Incubation, identity Incubation sponsors, and to queue for due diligence, if appropriate.

## Ready for Incubation?

Yes. Based on the graduation criteria, Meshery seems to already tick the relevant boxes required for incubation. We believe Meshery to be a strong candidate for incubation based on the following:

**Document that it is being used successfully in production by at least three independent end users**
Meshery is used in production by many more than three independent end users (see above). [SMI adopted](https://smi-spec.io/blog/validating-smi-conformance-with-meshery/) Meshery as it's conformance tool.

**Have a healthy number of committers**
Meshery has 10 regular maintainers across 8 organisations. Community contributions from many different organizations are very common and warmly received.

**Demonstrate a substantial ongoing flow of commits and merged contributions**
Evidence of this can be seen under DevStats above.

**A clear versioning scheme**
Meshery has had a clear versioning scheme (see [Build and Release](https://docs.meshery.io/project/contributing/build-and-release)) and a [regular release cadence](https://docs.meshery.io/project/releases).

**Clearly documented security processes explaining how to report security issues**
The [Security page](https://docs.meshery.io/project/security-vulnerabilities) in Meshery Docs page outlines the vulnerability reporting process, has a mailing list for reporting (security@meshery.dev), has a [published and resolved CVE](https://docs.meshery.io/project/security-vulnerabilities#list-of-announced-vulnerabilities).

Meshery is at "[passing](https://bestpractices.coreinfrastructure.org/projects/3564)" level for Core Infrastructure Best Practices.

Preemptively, if there are any concerns about Meshery's eligibility for incubation, we'd be more than happy to address them.
