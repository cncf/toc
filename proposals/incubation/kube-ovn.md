# Kube-OVN proposal for CNCF Incubation

## About Kube-OVN

Kube-OVN is a bridge between tradition SDN and modern Cloud Native that integrates the OVN-based Network Virtualization with Kubernetes. 
It offers an advanced Container Network Fabric for Enterprises with the most functions, extreme performance and the easiest operation.

Kube-OVN provides lots of features, here are some unique ones across the ecosystem:

**Namespaced Subnet**: Unlike most other CNIs that implement subnet per Node network topology, 
Kube-OVN implement subnet per Namespace network topology,
which gives more possibility for cluster administrators to control the network address range, network type, 
ACLs, routes and NAT rules for each Namespace. And the subnet not bind to a specific Node also gives more possibility 
for IPAM and high level network abstract.

**Overlay and Underlay Coexist**: Most other CNIs can only provide one type of network (either overlay or underlay) in a cluster.
Due to the Namespace Subnet design in Kube-OVN, we can provide different type of networks in one cluster by using different subnets.
Users in one cluster can create different type of subnet due to the application requirements.

**VPC Level Network**: Most other CNIs implement a simple flat L2 or L3 network in Kubernetes cluster, 
which is just like the classical network in tradition IaaS. Kube-OVN provides more modern VPC type multi-tenant network just like most 
IaaS provide today. This gives Kubernetes an opportunity to rebuild the datacenter infrastructure and 
actually some community users are building the new public/private cloud with Kubernetes, KubeVirt and Kube-OVN.

## Progress Since Sandbox

Kube-OVN was accepted as a CNCF Sandbox project on Jan 27, 2021.

### Community

The Kube-OVN devstats page and dashboards can be found [here](https://kubeovn.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, Kube-OVN currently has [115](https://kubeovn.devstats.cncf.io/d/22/prs-authors-table?orgId=1) PRs authors from [42](https://kubeovn.devstats.cncf.io/d/5/companies-table?orgId=1)
  organizations. On average, there are [~700 contributions per month](https://kubeovn.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) contained within [72 merged PRs per month](https://kubeovn.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=prs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).
- The community has grown significantly since the project entered the CNCF sandbox.
    - We held biweekly community meetings constantly. The meeting records can be found [here](https://space.bilibili.com/1932512297?spm_id_from=333.337.search-card.all.click). The average number of meeting attendees is ~20.
    - Number of PRs authors: 31 -> **115+**
    - Number of PRs authors organizations: 13 -> **42**
    - Number of Contributors: 62 -> **330**
    - Last Release: **v1.12.11**
    - Member of WeChat Group: **1426+**

### Governance

The Kube-OVN community has established [the community governance](https://github.com/kubeovn/kube-ovn/blob/master/GOVERNANCE.md), and promoted 4 maintainers since join CNCF Sandbox. 
Currently, we have 6 maintainers from 4 organizations.

We codified our [contributing guidelines](https://github.com/kubeovn/kube-ovn/blob/master/CODE_STYLE.md), and [security processes](https://github.com/kubeovn/kube-ovn/blob/master/SECURITY.md).


## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._**

Now there are at 50+ adopters that have used Kube-OVN in production, some adopters can be found [here](https://github.com/kubeovn/kube-ovn/blob/master/USERS.md) 
most of them don't have right to put the company names here, but we have frequently communication and can make sure they are using Kube-OVN.

**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

There are [115](https://kubeovn.devstats.cncf.io/d/22/prs-authors-table?orgId=1) PR authors in Kube-OVN, out of which there are 6 maintainers from 4 different organizations.
The maintainers can be found [here](https://github.com/kubeovn/kube-ovn/blob/master/MAINTAINERS).


**_Demonstrate a substantial ongoing flow of commits and merged contributions._**

A short summary has been provided in the aforementioned community section. Detailed statistics can be found in the following devstats links:
- [Commit Activity](https://kubeovn.devstats.cncf.io/d/74/contributions-chart?orgId=1)
- [New contributors](https://kubeovn.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1573369200000&to=now)
- [Issues open/closed](https://kubeovn.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1573369200000&to=now)
- [PRs open/closed](https://kubeovn.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1573369200000&to=now)
- [Contributing organizations](https://kubeovn.devstats.cncf.io/d/5/companies-table?orgId=1)

**_A clear versioning scheme._**

Kube-OVN uses the [semantic versioning scheme](https://semver.org/#semantic-versioning-specification-semver). 
The latest release is v1.12.11.

## Statement on Alignment with the CNCF Mission

Kube-OVN is trying to bring tradition SDN functions to cloud native ecosystem. 
Kube-OVN is a CNI implementation and has deeply integrated with Kubernetes, Prometheus, Submariner, Cilium and other CNCF projects.
We also see more and more users are trying to integrate Kube-OVN with KubeVirt together to rebuild the datacenter infrastructure. 
We believe Kube-OVN has extended cloud native computing to more scenarios.

## Future plan

The goal of Kube-OVN this year is to improve the stability and increase the diversity in both contributors and adoption . Some key works include:

- Build the test infrastructure to make writing and running test at ease.
- Improve the test coverage and inject errors in test to make Kube-OVN more robust.
- Improve the user experience in installation, upgrade and troubleshooting.
- Provide more document about the user guide, tutorial and internals.

