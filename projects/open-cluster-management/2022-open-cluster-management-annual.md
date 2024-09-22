# Open Cluster Management Sandbox Annual Review


## Background

[Open Cluster Management](https://open-cluster-management.io/) (OCM) is a project focused on simplifying management of multiple Kubernetes clusters regardless of where they are deployed. Open APIs are evolving within this project for cluster registration, work distribution, dynamic placement of policies and workloads, and much more. In doing so, OCM helps promote the overall growth and adoption of Kubernetes and its adjacent cloud-native technologies.

Open Cluster Management (OCM) provides a framework to enable any capability within the Kubernetes ecosystem to orchestrate across multiple clusters and cloud providers. The built-in primitives allow an understanding of cluster inventory, cluster workload placement, and cluster workload distribution. The hub-agent model of OCM allows for new kinds of agent behavior to be injected into managed clusters easily. New behaviors can further configure clusters or collect information about clusters (such as observability data or search indexing). The project also provides two out of the box capabilities: multicluster application delivery and orchestration, and multicluster policy-based governance. Additional integrations and add-ons can be easily developed through the use of the robust OCM Add-on Framework.


### Alignment with CNCF

- OCM was accepted as a CNCF Sandbox project on November 09, 2021.

## Development metrics

The OCM devstats page and dashboards can be found [here](https://openclustermanagement.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, OCM currently has [64](https://openclustermanagement.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [19](https://openclustermanagement.devstats.cncf.io/d/5/companies-table?orgId=1)
companies. On average, there are [~1085 contributions per month](https://openclustermanagement.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) contained within [134 merged PRs per month](https://openclustermanagement.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=prs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).
- [New PRs in last year](https://openclustermanagement.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
- The community has grown since the project entered the CNCF sandbox.
  - We consistently hold weekly community meetings. The meeting records can be found [here](https://www.youtube.com/@OpenClusterManagement/videos) for different timezone/language community members. The average number of meeting attendees is ~10.
  - Number of contributors: 120+ -> **150+**
  - Github stars: 600+ -> **1000+**
  - Github forks: 330+ -> **560+**
  - Contributing organizations: 5 -> **19**

Regarding these metrics, we would perceive the community as healthy and rising.

## Maintainers

We have established [the contributor ladder](https://github.com/open-cluster-management-io/community/blob/main/CONTRIBUTOR_LADDER.md), and quite a few active and qualified contributors joined to become members. We currently have [7 maintainers from 2 organizations](https://github.com/cncf/foundation/blob/main/project-maintainers.csv) as maintainers.

- Qiu Jian (Red Hat)	@qiujian16
- Joshua Packer	(Red Hat)	@jnpacker	
- David Eads	(Red Hat)	@deads2k	
- Mike Ng	(Red Hat)	@mikeshng	
- Min Kim	(Alibaba)	@yue9944882	
- Yong Feng	(Alibaba)	@luckyfengyong

## Project adoption

OCM has several production environment adopters. They are:

- [Alibaba Cloud](https://www.aliyun.com/)
- [Ant Group](https://www.antgroup.com/)
- [Red Hat](https://www.redhat.com/)

There are several more projects and companies that are experimenting with OCM:

- [Argo Workflow](https://argoproj.github.io/argo-workflows/)
- [ESET](https://www.eset.com/)
- [Expedia](https://www.expediagroup.com/)
- [TietoEVRY](https://www.tietoevry.com/)

## Project goals

The primary goal of OCM is to enable capabilities within the Kubernetes ecosystem to orchestrate across multiple clusters and cloud providers. In the past year, we have achieved the following development outcomes:
- Release cadence: 4 minor releases, roughly once every three months. The latest release is v0.9.0.
- Key features added to the project (complete changelog can be found [here](https://open-cluster-management.io/community/releases/)):
  - 0.6
    - New add-ons
      - Cluster Proxy, allows a proxied connection from the hub to a managed cluster's API ingress.
      - Managed Service account, allows the hub to define service account and receive its token from a managed cluster.
    - Sync status of applied resources in ManifestWork API.
    - Placement extensible scheduling.
    - Helm chart interface for addon framework.
  - 0.7
    - New major features to Clusteradm CLI tool.
      - Create, bind and view ManagedClusterSet.
      - Create and view ManifestWork.
      - Check the multicluster controlplane status.
      - Upgrade hub and klusterlet.
    - A default ManagedClusterSet
    - New PolicySet API.
  - 0.8
    - ManifestWork update strategy.
    - Global ClusterSet.
    - Configuring feature gates for Klusterlet and Cluster Manager.
    - Support host alaises for Klusterlet.
    - Running policy addon using Clusteradm CLI tool.
  - 0.9
    - De-escalate ManifestWork agent privilege on Managed Clusters.
    - Support referencing the add-on configuration with add-on related APIs.
    - Allow targeting specific services within Managed Clusters.
    - Application add-on is now able to expose custom Prometheus metrics via the Git subscription.

In the near future, we aim to achieve the following technical goals:
- Running multiple Hubs on a single cluster.
- Introduce ArgoCD pull model integration.
- Support to communicate between the hub and managed clusters in a proxy environment.
- Add delete options to ManagedCluster to cleanup ManifestWorks and Addons. 
- Enhance Add-on Framework capabilities: rolling upgrade, install strategy, and install template.

A complete project roadmap can be found [here](https://github.com/orgs/open-cluster-management-io/projects/2/views/9).

We have built close connections with members in the CNCF organization for project integration. 
We have attended multiple conferences such as KubeCon, DevConf and others.
We are actively looking for more user adoptions and community collaborations.

## How the CNCF can help to achieve the upcoming goals

- More chances to advocate the project.
- More chances to collaborate with other projects in CNCF or even out of CNCF.
- Technical writing support for project documents.

## Incubation readiness

- We are preparing for the incubation proposal.
