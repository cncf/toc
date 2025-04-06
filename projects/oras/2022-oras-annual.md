
# ORAS CNCF Sandbox Annual Review 2022

## Background

Registries are evolving as generic artifact stores. To enable this goal, the [ORAS](https://oras.land/) project provides a way to push, discover and pull [OCI Artifacts](https://github.com/opencontainers/artifacts) to and from OCI Registries, including support for the new [OCI Artifact reference types](https://github.com/opencontainers/image-spec/blob/main/artifact.md).

Users seeking a generic registry client can benefit from the [ORAS CLI](https://oras.land/CLI/), while developers can build their own OCI clients using the [oras-go](https://github.com/oras-project/oras-go) or [oras-py](https://github.com/oras-project/oras-py) SDKs.

ORAS was accepted as a CNCF Sandbox project in June 2021.

## Development Metrics

The ORAS devstats page and dashboards can be found [here](https://oras.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, ORAS contributors increased to [76](https://oras.devstats.cncf.io/d/22/prs-authors-table?orgId=1), up from [30](https://oras.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Before%20joining%20CNCF&var-repogroup_name=All) before joining CNCF. 
- On average, there are [~574 contributions per month](https://oras.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) contained within [25 merged PRs per month](https://oras.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=prs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).
- These contributors come from [46 companies](https://oras.devstats.cncf.io/d/5/companies-table?orgId=1), up from [31](https://oras.devstats.cncf.io/d/5/companies-table?orgId=1).
- All [Stars](https://oras.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now) increased from 504 to 835, all [Forks](https://oras.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now) increased from 71 to 182 in the past one year.
- ORAS CLI, SDK, and spec had 6, 16, and 4 releases respectively since joining CNCF.
- The total downloads of ORAS CLI is **418,573**.
- [The number of new PRs has tripled in the last year](https://oras.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
- Organized 22 community meetings in the past one year, see [meeting notes](https://hackmd.io/P-O6n222TcSMoJgHmTTduw?view).

Regarding these metrics listed above, we would perceive the community as healthy and growing.

## Maintainers

The ORAS community has established [the community governance](https://github.com/oras-project/community/blob/main/governance/GOVERNANCE.md) to engage active and qualified contributors, who have been promoted to maintainers, release managers, and subproject owners. 

 - Avi Deitcher (@deitch) from Aptimia 
 - Josh Dolitsky (@jdolitsky) from Chainguard
 - Sajay Antony (@sajayantony) from Microsoft
 - Shiwei Zhang (@shizhMSFT) from Microsoft
 - Steve Lasker (@stevelasker) from Microsoft
 - Sylvia Lei (@Wwwsylvia) from Microsoft
 - Vanessasaurus (@vsoch) from Lawrence Livermore National Lab 
 - Justin Cormack (@justincormack) from Docker
 - Michael Brown (@michaelb990) from Amazon
 - Yi Zha (@yizha1) from Microsoft
 - Feynman Zhou (@feynmanzhou) from Microsoft
 
## Project Adoption

ORAS CLI, ORAS Go and Python SDK are designed to help users and developers manage [OCI Distribution](https://github.com/opencontainers/distribution-spec) based artifacts. ORAS empowers the secure supply chain by enabling users to leveraging the existing services they already have across their development to production environments.

Currently, the biggest cloud providers like Microsoft Azure and AWS are using ORAS to interact with container registries. ORAS Go SDK has been integrated and adopted by some industry-leading ISVs and OSS projects. Here are part of known adopters: 

- [Amazon ECR](https://aws.amazon.com/ecr/)
- [Alibaba Cloud Service Mesh](https://www.alibabacloud.com/help/en/alibaba-cloud-service-mesh/latest/use-oras-to-simplify-wasm-based-asm-instance-extension)
- [GitHub](https://github.com/)
- [Helm](https://v3.helm.sh/docs/topics/registries/)
- [Singularity](https://sylabs.io/guides/3.1/user-guide/cli/singularity_push.html)
- [Microsoft Azure - ACR](https://azure.microsoft.com/en-us/products/container-registry/)
- [Notary v2](https://github.com/notaryproject/notation)
- [KubeApps](https://github.com/vmware-tanzu/kubeapps) by VMware Tanzu
- [VMware Application Catalog](https://tanzu.vmware.com/application-catalog)
- [UOR Framework](https://universalreference.io/) by Red Hat
- [soci-snapshotter](https://github.com/awslabs/soci-snapshotter) by AWS
- [Zot](https://github.com/project-zot/zot)

## Project Goals

As cloud native development continues to grow, we have seen increased community interest in evolving registries to natively push, discover, copy and pull a graph of supply chain artifacts. Artifact references are important for many use cases such as adding Software Bill of Materials (SBoM), security scan results, and container image (artifact) signatures. 

ORAS 0.15 and future milestones will provide more capabilities to easily manage OCI content and interact with registries. It will empower the container secure supply chain and focus on the following areas:

- Be able to manage repository, tag, manifest, and blob
- Support and migrate to OCI reference types
- Support push, discover, pull, copy artifacts from OCI Image Layout
- Support all OCI-compatible registries
- Add E2E testing to enhance the project quality and stability

See the ORAS [Roadmap](https://github.com/oras-project/community/blob/main/Roadmap.md) for more details.

## How the CNCF Can Help To Achieve the Upcoming Goals

- More chances to advocate for the project to broader users
- More chances to collaborate with other projects in CNCF
- Certification service for maintainers and active contributors
- Technical writing and website design support

## Incubation Readiness

With the current status and the existing plans we believe that ORAS is on a good track towards incubating stage. So we are planning to meet the incubating criteria and submit our application in the first half of 2023.

We believe that the ORAS project is in a healthy growing status in both user adoption and contributors. We are working to properly document existing production usage and we expect that the contributor growth will also imply a diverse group of maintainers.


