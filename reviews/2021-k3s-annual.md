# K3s 2021 Annual Review

## Table of Contents

- [Background](#background)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

K3s is a Lightweight Kubernetes distribution that ships all needed components and utilities,
sans images, embedded in a single binary that weighs in at under 60mb (with compressed air-gap images adding ~150mb).

K3s is a [fully conformant](https://github.com/cncf/k8s-conformance/pulls?q=is%3Apr+k3s) production-ready distribution with some additional flavor:

1. Packaged as a single binary.
2. Support for sqlite3 as the default storage backend. Etcd3, MySQL, and Postgres are also supported.
3. Wraps Kubernetes and other components in a single, simple launcher.
4. Secure by default with reasonable defaults for lightweight environments.
5. Minimal to no OS dependencies (just a sane kernel and cgroup mounts needed).
6. Eliminates the need to expose a port on Kubernetes worker nodes for the kubelet API by exposing this API to the Kubernetes control plane nodes over a websocket tunnel.
7. Internet installation via [https://get.k3s.io](https://get.k3s.io "Y.O.L.O.")

K3s bundles the following technologies together into a single cohesive distribution:

* [runc](https://github.com/opencontainers/runc) as the default OCI runtime
* [Containerd](https://containerd.io/) as the default CRI
* [Flannel](https://github.com/coreos/flannel) as the default CNI
* [CoreDNS](https://coredns.io/) as kube-dns
* [Metrics Server](https://github.com/kubernetes-sigs/metrics-server) exposing metrics
* [Traefik](https://containo.us/traefik/) for ingress
* [Klipper-lb](https://github.com/k3s-io/klipper-lb) as an embedded service loadbalancer provider
* [Kube-router](https://www.kube-router.io/) for network policy
* [Helm-controller](https://github.com/k3s-io/helm-controller) to allow for CRD-driven deployment of helm manifests
* [Kine](https://github.com/k3s-io/kine) as a datastore shim that allows etcd to be replaced with other databases
* [Local-path-provisioner](https://github.com/rancher/local-path-provisioner) for provisioning volumes using local storage
* [Host utilities](https://github.com/k3s-io/k3s-root) such as iptables/nftables, ebtables, ethtool, & socat

Additionally, K3s simplifies Kubernetes operations by maintaining functionality for:

* Managing the TLS certificates of Kubernetes components
* Managing the connection between worker and server nodes
* Drop-in manifest apply for all built-in and custom resources
* Managing an embedded etcd cluster

### Accepted into CNCF Sandbox

**K3s was accepted as a CNCF Sandbox project on August 19, 2020**

- Sandbox proposal [pr #447](https://github.com/cncf/toc/pull/447)
- Sandbox proposal [doc](https://github.com/cncf/toc/blob/main/proposals/sandbox/k3s.md)

### Alignment with Cloud Native Computing Foundation

K3s falls in the scope of [CNCF Runtime TAG](https://github.com/cncf/tag-runtime). We recently presented an overview of K3s to TAG Runtime, which can be viewed [here](https://youtu.be/ptoiC8fAPms).


**Key features added**:

- etcd snapshots
- etcd member management
- etcd backup and restore
- automatic discovery of nvidia container runtime
- containerd v1.5.x
- traefik v2.x
- selinux support

## Annual Review Contents

### Include a link to your project’s devstats page.
Here is a link to our devstats page: https://k3s.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m

In reviewing this data, we noticed a few issues. We've opened the following issues to resolve them:
 - [[bug] missing data for k3s](https://github.com/cncf/devstats/issues/320)
 - [Change to the repos that are in k3s devstats](https://github.com/cncf/devstats/issues/319)

Unfortunately, until these issues are resolved, it is difficult to draw accurate conclusions from the devstats page. In the meantime, here is some data that can be gleaned directly from GitHub:
| Statistic | Sandbox | Current |
|-|-|-|
| [GitHub Stars](https://github.com/k3s-io/k3s/stargazers) | ~14k | ~18k |
| [Commits](https://github.com/k3s-io/k3s/commits/master) | ~75/month | ~65/month |
| [Merged PRs](https://github.com/k3s-io/k3s/pulls) | ~35/month | ~45/month |
| [Releases](https://github.com/k3s-io/k3s/releases) | 100 | 171 |

### How many maintainers do you have, and which organizations are they from?

Our maintainers list: https://github.com/k3s-io/k3s/blob/master/MAINTAINERS

Currently, the project has 11 maintainers, all with SUSE:
- [Brad Davidson](https://github.com/brandond "@brandond")
- [Brian Downs](https://github.com/briandowns "@briandowns")
- [Craig Jellick](https://github.com/cjellick "@cjellick")
- [Chris Wayne](https://github.com/cwayne18 "@cwayne18")
- [Derek Nola](https://github.com/dereknola "@dereknola")
- [Jacob Blain Christen](https://github.com/dweomer "@dweomer")
- [Hussein Galal](https://github.com/galal-hussein "@galal-hussein")
- [Darren Shepherd](https://github.com/ibuildthecloud "@ibuildthecloud")
- [Manuel Buil](https://github.com/manuelbuil "@manuelbuil")
- [Chris Kim](https://github.com/Oats87 "@Oats87")
- [Michal Rostecki](https://github.com/vadorovsky "@vadorovsky")

### What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?

Based on issues filed on GitHub and interactions on Slack, K3s continues to see a healthy adoption from many different types of users from hobbyists to large organizations. However, we haven't begun to formally track adopters or ask their permission to list them as such. We will start to do so now.

One adoption highlight is Civo's use of K3s to power its Kubernetes cloud offering, see: https://www.civo.com/kubernetes.

K3s is also used as a building block for other projects and products. Here are a few that are independent of K3s, but leverage or build on top of it:
* [k3d](https://github.com/rancher/k3d)
* [k3sup](https://github.com/alexellis/k3sup)
* [Rancher Desktop](https://rancherdesktop.io/)
* [RKE2](https://github.com/rancher/rke2)

We also track downloads of our binaries from GitHub (the primary way to install K3s). Over the past year, our all-time number of downloads grew from around 1.2 million to 2.7 million. Note that this number does not represent unique users. It is simply the number of times the k3s binaries have been downloaded.


### How has the project performed against its goals since the last review?

This is our first annual review, so here is a list of our accomplishments since joining the CNCF. When we submitted for inclusion in the CNCF, we outlined the following roadmap items:

- Improving stability and support across operating systems and architectures
  - _We did achieve this by improving our ARM builds and supporting more operating systems_
- Enhancing security through SELinux support and CIS compliance
  - _We did introduce SELinux support and published a guide evaluating K3s against the CIS Kubernetes benchmark_
- Developing HA solutions appropriate for small edge clusters where even three nodes are cost prohibitive
  - _We did not complete this. We chose instead to focus on embedded etcd support, which meets the use cases we were concerned about._
- Improving support for Kubernetes cloud providers
  - _We chose not to do this. Cloud providers are moving out of tree and K3s supports that mode properly. We decided that was enough._
- Improved support for database backends, including embedded etcd
  - _We achieved this. K3s now supports an "embedded etcd" mode where it will launch and manage etcd, including snapshot and restore functionality._


Beyond these above items, here is a list of our accomplishment highlights over the past year. We've organized these into "Keeping our house in order" to represent improvements around process, community, and contributor experience and "Features and fixes" for major enhancements or fixes to our project:

- Keeping our house in order
  - Improved framework and documentation for adding unit, integration, and end-to-end tests
  - Publishing weekly status and accomplishment reports
  - Enabled GitHub Discussions for better distinction between questions and issues
  - Contributor process quality of life improvements
    - Issue and PR templates
    - Stalebot integration
  - Maintained release cadence with upstream Kubernetes
    - 3-4 scheduled patch releases a month
    - Minor (aka feature) release once every 3-4 months
- Features and fixes
  - Embedded etcd support -  HA, snapshotting, and restoring
  - GAed SELinux support
  - Improved airgap support
  - Traefik v2 integration
  - cgroupv2 support
  - Kine (DB shim) performance and reliability improvements including transaction-locked pure-sql compaction, metrics support, and stubbed out GRPC services for compatibility with other distros such as kubeadm
  - Kubernetes CIS hardening guide evaluation and conformance
  - stargz snapshotter support
  - Nvidia device plugin support
  - Many more - hundreds of issues closed across 40+ release since last October

Additionally, the maintainer team has contributed to projects that K3s relies on or integrates with, including:
* [container-selinux](https://github.com/containers/container-selinux)
* [containerd](https://github.com/containerd/containerd)
* [etcd](https://github.com/etcd-io/etcd)
* [flannel](https://github.com/flannel-io/flannel)


One area where we've fallen short is in getting the Kubernetes patches we carry submitted to upstream. We didn't have the bandwidth to do so. Of these patches, our biggest/most interesting change is the use of our project [KINE](https://github.com/k3s-io/kine) to allow for different SQL backends for Kubernetes. We'd like to submit this to upstream Kubernetes in the upcoming year.

### What are the current goals of the project?

Similar to our accomplishments, our goals can be organized into "Keeping our house in order" and "Features":
- Keeping our house in order
  - Improve and expand automated testing
  - Overhaul user documentation
  - Implement build and release improvements
  - Improve contributor quality of life improvements: docs, bots, and process improvements
  - Better engagement with upstream(s) via contributions (PR and KEP)
- Features
  - Implement expanded networking support
    - IPv4/IPv6 dualstack support
    - More first class CNI choices
    - SR-IOV support
  - Implement security and compliance improvements 
    - Continue to improve SELinux support
    - FIPS 140-2/3 crypto validation (a U.S. government standard for cryptography)
    - Automated base image scanning and revving for CVE mitigation
    - Ensure proper support and integration for PSP v2
    - Custom CIS-like hardening guide for K3s
  - Expand our support matrix
    - An ever-expanding list of supported OSes
  - Speculative (highly subject to change)
    - Windows
    - Embedded registry or image airgapping solution
    - Deeper “tiny edge” integration (with projects like Akri and smarter-device-manager)

We want to achieve the above roadmap while also maintaining cadence with upstream Kubernetes. This means:
- Delivering patch releases that lag mere days behind upstream (with CVEs ideally same day or day after)
- Delivering minor releases that lag at most 2-4 weeks behind upstream

### How can the CNCF help you achieve your upcoming goals?

The CNCF has been helpful thus far by contributing CI resources. As we worked towards incubation, we'll likely reach out with specific, tactical asks as we encounter them.

### Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

At this time, we aren't ready to submit for incubation status. Our goal is to submit next year. We'll use this upcoming year to prep for that and improve the areas of the project that need it.


## Project Links
 - [web://k3s.io](https://k3s.io)
 - [github://k3s-io/k3s](https://github.com/k3s-io/k3s)
 - [slack://rancher-users#k3s](https://rancher-users.slack.com/archives/CGGQEHPPW)
