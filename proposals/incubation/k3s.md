# k3s Incubation Stage Proposal

k3s (https://k3s.io/) joined the CNCF as a Sandbox project in 2020. As the k3s team, we believe we have met the criteria to graduate to the Incubation stage of a CNCF project.

## Incubation State Requirements

1. *Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.*

    * Chick-Fil-A, a US-based fast food restaurant, uses k3s extensively in their stores' POS systems (outlined here)[https://medium.com/chick-fil-atech/what-we-learned-from-launching-edge-compute-from-enterprise-architecture-1dc34e49482f]

    * (Gitpod.io)[gitpod.io], a cloud development environment, is [run on k3s](https://github.com/gitpod-io/gitpod/pull/18521)

    * (Home Depot)[homedepot.com] a US-based hardware store, runs k3s extensively across their (~2300 stores)[https://www.datacenterknowledge.com/open-source/home-depot-upgrades-2300-retail-edge-locations-using-suse-rancher-k3s]
    
    * (Hypergiant)[https://www.hypergiant.com/] chose K3s to run AI/ML workloads (in space aboard a satellite)[https://sanjimoh.medium.com/putting-kubernetes-in-space-e0918a3dce26]

2. *Have a healthy number of committers.*

    According to (GitHub)[https://github.com/k3s-io/k3s/pulse/monthly] k3s has had 130 commits (across all k3s release branches) from 12 authors in just the last month alone.  According to (LFX Insights)[https://insights.lfx.linuxfoundation.org/foundation/cncf/overview?project=k3s&bestPractice=false there have been 1146 contributors in the last year.  Additionally, there have been commits from 353 companies according to (devstats)[https://k3s.devstats.cncf.io/d/5/companies-table?orgId=1]

3. *Demonstrate a substantial ongoing flow of commits and merged contributions.*

    Commits to master over the last quarter: **240** (Ref LFX Insights)[https://insights.lfx.linuxfoundation.org/foundation/cncf/overview/github?project=k3s&bestPractice=false&repository=all&dateFilters=Last%20Quarter&dateRange=2024-01-01%20to%202024-03-31&compare=PP&granularity=week&hideBots=true&routedFrom=Github]

    Issues opened over the last quarter **442** (Ref LFX Insights)[https://insights.lfx.linuxfoundation.org/foundation/cncf/reports/activities?project=k3s&bestPractice=false&repository=all&dateFilters=Last%20Quarter&dateRange=2024-01-01%20to%202024-03-31&compare=PP&granularity=week&hideBots=true&routedFrom=Github]

    Issues closed over the last quarter: **419** (Ref LFX Insights)[https://insights.lfx.linuxfoundation.org/foundation/cncf/reports/activities?project=k3s&bestPractice=false&repository=all&dateFilters=Last%20Quarter&dateRange=2024-01-01%20to%202024-03-31&compare=PP&granularity=week&hideBots=true&routedFrom=Github]

    New PRs over the last quarter: **310** (Ref LFX Insights)[https://insights.lfx.linuxfoundation.org/foundation/cncf/reports/activities?project=k3s&bestPractice=false&repository=all&dateFilters=Last%20Quarter&dateRange=2024-01-01%20to%202024-03-31&compare=PP&granularity=week&hideBots=true&routedFrom=Github]

    All data was gathered thanks to [LFX Insights](https://insights.lfx.linuxfoundation.org/foundation/cncf/overview/github?project=k3s&bestPractice=false&repository=all&dateFilters=Last%20Quarter&dateRange=2024-01-01%20to%202024-03-31&compare=PP&granularity=week&hideBots=true&routedFrom=Github)

4. *A clear versioning scheme.*

    The version scheme follows the upstream versioning scheme of Kubernetes, with an additional +k3sX build (i.e. our latest version was v1.29.4+k3s1) 

## What is k3s: Recap

K3s is a fully compliant Kuberenetes distribution with the following enhancements:

* Distributed as a single binary or minimal container image.
* Lightweight datastore based on sqlite3 as the default storage backend. etcd3, MySQL, and Postgres are also available.
* Wrapped in simple launcher that handles a lot of the complexity of TLS and options.
* Secure by default with reasonable defaults for lightweight environments.
* Operation of all Kubernetes control plane components is encapsulated in a single binary and process, allowing K3s to automate and manage complex cluster operations like distributing certificates.
* External dependencies have been minimized; the only requirements are a modern kernel and cgroup mounts.

### Architecture

K3s bundles the following major components:

* containerd / cri-dockerd container runtime (CRI)
* Flannel Container Network Interface (CNI)
* CoreDNS Cluster DNS
* Traefik Ingress controller
* ServiceLB Load-Balancer controller
* Kube-router Network Policy controller
* Local-path-provisioner Persistent Volume controller
* Spegel distributed container image registry mirror
* Host utilities (iptables, socat, etc)

For more information, please see (K3s Architecture)[https://docs.k3s.io/architecture]

## Statement on alignment with the CNCF mission

The K3s project aligns with CNCF's mission to make cloud-native computing ubiquitous. With its small footprint and ease of use, K3s brings more developers and tinkerers into the cloud-native ecosystem by making Kubernetes easy.  With its ability to run on smaller devices, K3s opens up more use cases tha ever for Kubernetes, allowing folks to learn and use Kubernetes on something as small as a Raspberry Pi, or something as large as a datacenter.  The ease of use and small size allows for the use of Kubernetes in CI systems and allows for a more rapid development cycle for apps running on Kubernetes, helping to grow the community around Kubernetes.


## Project Progress since Sandbox

K3s has joined the CNCF in 2020. Since then, we've seen the continuous growth of the K3s community.

K3s has continued to grow its user base, it's base of contributors, as well as its feature set since first being accepted into the CNCF Sandbox.  Since being accepted into the sandbox, K3s has added:

* Embedded Registry Support via (Spegel)[https://github.com/spegel-org/spegel]
* IPV6-only and dual-stack support
* CA Certificate Rotation
* Secrets Encryption
* Embedded ETCD support
* Multus Support

Among countless other bug fixes, new OS support, and support for new Kubernetes versions.


### Community growth

We've seen tremendous growth since we joined CNCF four years ago. Compare 2020 to 2024:

* GitHub stars: 11k+ vs 26k+ (2+x)
* Github Downloads: ~300k vs ~9Million (30x)
