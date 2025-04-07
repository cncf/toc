# metallb 2022 Annual Review

This is a Kubernetes ToC Annual Review for the [metallb](https://metallb.universe.tf/) project for 2022.

<!-- omit in toc -->
## Table of Contents

- [Background](#background)
- [DevStats](#development-metrics)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Latest Achievements](#latest-achievements)
- [Project Goals](#project-goals)
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)

## Background

[MetalLB](https://metallb.universe.tf/) is a load-balancer implementation for bare metal Kubernetes clusters, using standard routing protocols.

Kubernetes does not offer an implementation of network load balancers (Services of type LoadBalancer) for bare-metal clusters. The implementations of network load balancers that Kubernetes does ship with are all glue code that calls out to various IaaS platforms (GCP, AWS, Azure…). If you’re not running on a supported IaaS platform (GCP, AWS, Azure…), LoadBalancers will remain in the “pending” state indefinitely when created.

Bare-metal cluster operators are left with two lesser tools to bring user traffic into their clusters, “NodePort” and “externalIPs” services. Both of these options have significant downsides for production use, which makes bare-metal clusters second-class citizens in the Kubernetes ecosystem.

MetalLB aims to redress this imbalance by offering a network load balancer implementation that integrates with standard network equipment, so that external services on bare-metal clusters also “just work” as much as possible.

MetalLB works with both relatively simple networking topology using the L2 mode that does not require any additional network component, or using the BGP protocol when BGP enabled routers are available.

More information about the MetalLB architecture can be found [here](https://metallb.universe.tf/concepts/)

## Development metrics

The MetalLB devstats page and dashboards can be found [here](https://metallb.devstats.cncf.io/).

Here are some links worth mentioning:

- [Stars and forks](https://metallb.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now)
- [Issues](https://metallb.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1)
- [PRs](https://metallb.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&var-period=d&var-repogroup_name=All)
- [Commits](https://metallb.devstats.cncf.io/d/2/commits-repository-groups?orgId=1)

## Maintainers

There are three maintainers actively working on the project (alphabetically):

| Maintainer          | GitHub ID                                         | Affiliation |
| ------------------- | ------------------------------------------------- | ----------- |
| Federico Paolinelli | [fedepaol](https://github.com/fedepaol)           | Red Hat     |
| Graeme Lawes        | [gclawes](https://github.com/gclawes)             | Nasdaq      |
| Ori Braunshtein     | [oribon](https://github.com/oribon)               | Red Hat     |

## Adoption

MetalLB is probably the most popular solution for exposing services of type LoadBalancer in bare metal deployments. It is widely used by users running kubernetes in homelabs
(an [issue opened](https://github.com/metallb/metallb/issues/1481#issuecomment-1176716132) against MetalLB revealed a wide community of users in the k8s@home discord channel).
MetalLB is part of [multiple tutorials and videos](https://www.google.com/search?q=metallb+kubernetes+tutorial) written by the community and suggested by multiple projects
such as [kubespray](https://github.com/kubernetes-sigs/kubespray/blob/master/docs/metallb.md), [kind](https://kind.sigs.k8s.io/docs/user/loadbalancer/) and [microk8s](https://microk8s.io/docs/addon-metallb) (to mention some).

Despite not having a clear estimate of the number users / companies consuming MetalLB, the images related to the [controller](https://quay.io/repository/metallb/controller) / [speaker](https://quay.io/repository/metallb/speaker) components show an average of
40k "actions" per day for the controller image (which is deployed in single instance) and 150k "actions" per day for the speaker image.

We have a [who is using MetalLB](https://github.com/metallb/metallb/issues/5) issue to provide some form of tracking, but that is not really indicative.

On the enterprise side, MetalLB is distributed as part of Red Hat's Openshift
since [version 4.9](https://docs.openshift.com/container-platform/4.9/networking/metallb/about-metallb.html), [Microsoft's Azure HCI](https://learn.microsoft.com/en-us/azure-stack/aks-hci/deploy-metallb), [Platform9](https://platform9.com/docs/kubernetes/metallb-addon2), Google's [Anthos on premise clusters](https://cloud.google.com/anthos/clusters/docs/on-prem/1.11/how-to/bundled-load-balance-metallb), Amazon's [EKS anywhere](https://anywhere.eks.amazonaws.com/docs/tasks/packages/metallb/) and suggested in [Canonical's kubernetes distribution](https://ubuntu.com/kubernetes/docs/metallb).

In the next year we hope to have a clearer identity / sizing of the userbase via surveys and other statistics.

## Latest Achievements

The biggest achievement of the past year was adding a very comprehensive CI based on github actions and kind. Going from no end to end tests to
covering all the existing and new features lowers the bar for new contributions and avoids the risk of regressions, not to mention a quicker development
cycle.

Feature wise, the following improvements were done:

- Replaced the native BGP implementation with FRR, which made the implementation of new features faster
- Moved from a configmap based configuration to a custom resource based configuration
- Added support for IPV6 and dual stack clusters
- Added support for BFD for quicker failover
- Added support for interface selection in case of L2 mode
- Added more flexible options for announcement, such as node selector and BGP peer selector
- Added a MetalLB Operator to deploy MetalLB via OLM
- Images are co-signed using cosign
- A proper kustomize friendly layout to deploy the various configuration flavours

On top of that, we kept the project aligned with the evolution of the platform, moving from pod security policy to pod security admission, removing deprecated apis and
adding support for loadbalancer class.

## Project Goals

<!-- What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation? -->

MetalLB must continue to evolve with two goals in mind. On one hand it must keep being the (nearly) zero configuration load balancer that enables loadbalancers on Kubernetes clusters running on RaspberryPIs, on the other it must offer optional configuration knobs to accommodate the requirements of enterprise customers that are using MetalLB on more complex scenarios and topologies.

On top of that, the project must keep being contributor friendly and open to discuss new feature requests and responsive to contributors willing to help.

## CNCF Membership

<!-- How can the CNCF help you achieve your upcoming goals? -->

We don't have particular needs that must be addressed. The interaction with the CNCF has always been productive and we are already leveraging a few of the benefits provided by the CNCF.

## Incubation

This was the first year after the sandboxing process finished. Despite MetalLB probably already meets the majority of the criteria for Incubation, we will probably invest more time in filling the gaps and starting the process next year.
