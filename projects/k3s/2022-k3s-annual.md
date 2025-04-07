# K3s 2022 Annual Review

## Table of Contents

- [Background](#background)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

K3s is a Lightweight Kubernetes distribution that ships all needed components and utilities,
sans images, embedded in a single binary that weighs in at under 70mb (with compressed air-gap images adding ~150mb).

K3s is a [fully conformant](https://github.com/cncf/k8s-conformance/pulls?q=is%3Apr+k3s) production-ready distribution with some additional flavor:

1. Packaged as a single binary.
2. Support for sqlite3 as the default storage backend. Etcd3, MySQL,MariaDB, NATS JetStream, and Postgres are also supported.
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

K3s falls in the scope of [CNCF Runtime TAG](https://github.com/cncf/tag-runtime). We presented an overview of K3s to TAG Runtime, which can be viewed [here](https://youtu.be/ptoiC8fAPms).


**Key features added**:

- containerd v1.6.x
- etcd 3.5
- NATS JetStream added as supported storage backend
- MicroOS support
- Increased focus on automated testing across multiple OS/configs
- encryption key secret rotation
- bootstrap reconciliation refactor
- dual-stack support
- IPv6-only support
- Automatic detection of nvidia container runtime
- Automatic detection of systemd cgroup manager
- Control-plane egress proxy (similar to Konnectivity service) using existing agent websocket tunnel

## Annual Review Contents

### Include a link to your project’s devstats page.
Here is a link to our devstats page: https://k3s.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m


Github specific stats:

| Statistic | Sandbox | 2021 | 2022 | 
| -- | -- | -- | -- |
| [GitHub Stars](https://github.com/k3s-io/k3s/stargazers) | ~14k | ~18k | ~21k | 
| [Commits](https://github.com/k3s-io/k3s/commits/master) | ~75/month | ~65/month | [~100/month](https://github.com/k3s-io/k3s/pulse/monthly)  |
| [Merged PRs](https://github.com/k3s-io/k3s/pulls) | ~35/month | ~45/month | [~66/month](https://github.com/k3s-io/k3s/pulse/monthly)  | 
| [Releases](https://github.com/k3s-io/k3s/releases) | 100 | 171 | 189 |

### How many maintainers do you have, and which organizations are they from?

Our maintainers list: https://github.com/k3s-io/k3s/blob/master/MAINTAINERS

Currently, the project has 17 maintainers, all with SUSE:
- [Brad Davidson](https://github.com/brandond "@brandond")
- [Brian Downs](https://github.com/briandowns "@briandowns")
- [Brooks Newberry](https://github.com/brooksn "@brooksn")
- [Caroline Davis](https://github.com/carolid "@carolid")
- [Chris Kim](https://github.com/Oats87 "@Oats87")
- [Chris Wayne](https://github.com/cwayne18 "@cwayne18")
- [Derek Nola](https://github.com/dereknola "@dereknola")
- [Hussein Galal](https://github.com/galal-hussein "@galal-hussein")
- [Manuel Buil](https://github.com/manuelbuil "@manuelbuil")
- [Matt Trachier](https://github.com/matttrach "@matttrach")
- [Max Ross](https://github.com/rancher-max "@rancher-max")
- [MD Rahman](https://github.com/mdrahman-suse "@mdrahman-suse")
- [Nikolai Shields](https://github.com/nikolaishields "@nikolaishields")
- [Roberto Bonafiglia](https://github.com/rbrtbnfgl "@rbrtbnfgl")
- [Shylaja Devadiga](https://github.com/shylajadevadiga "@shylajadevadiga")
- [Thomas Ferrandiz](https://github.com/thomasferrandiz "@thomasferrandiz")
- [Justin Janes](https://github.com/VestigeJ "@VestigeJ")

### What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?

Based on issues filed on GitHub and interactions on Slack, K3s continues to see a healthy adoption from many different types of users from hobbyists to large organizations. While we have setup an ADOPTERS.md, we've not yet start to reach out to known adopters in earnest to get them listed.  There are however several known adopters, specifically Civo Cloud, Rocketchat, and Gitpod.  We will continue to reach out to these and other adopters to ensure they are properly tracked in ADOPTERS.md.

Additionally, through our devstats, you can see a healthy number of contributors from outside collaborators across several companies, in fact you can see in the last year, we've accepted contributions from over [180 different companies](https://k3s.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions).


K3s is also used as a building block for other projects and products. Here are a few that are independent of K3s, but leverage or build on top of it:
* [k3d](https://github.com/k3d-io/k3d)
* [k3sup](https://github.com/alexellis/k3sup)
* [Rancher Desktop](https://rancherdesktop.io/)
* [RKE2](https://github.com/rancher/rke2)
* [k3ai](https://k3ai.github.io/)
* [Kairos](https://kairos.io)

We also track downloads of our binaries from GitHub (the primary way to install K3s). Over the past year, our all-time number of downloads grew from around 2.7 million to over 4 million. Note that this number does not represent unique users or installations. It is simply the number of times the k3s binaries have been downloaded directly from GitHub.


### How has the project performed against its goals since the last review?

Over the last year, we have worked largely around process, automation, testing, quality, and documentation as well as continuing to keep up with upstream Kubernetes releases, and adding new features requested by the community.  Much of what we have accomplished over the last year speaks to these goals, as we have continued to push on processes and testing specifically, as well as getting our releases out much more quickly.  Previously, we often did not release a k3s release for a new upstream Kubernetes minor release until .2 or .3, whereas over this last year we have begun to get our first releases out along with .0 upstream.  We have also pushed to add more transparency and accountability to the community, by enabling GitHub discussions, and posting weekly reports detailing PR's opened+closed, issues opened+closed, while highlighting PR's from the community, which have continued to increase in frequency.


Beyond these above items, here is a list of our accomplishment highlights over the past year. We've organized these into "Keeping our house in order" to represent improvements around process, community, and contributor experience and "Features and fixes" for major enhancements or fixes to our project:

- Keeping our house in order
  - Improved framework and documentation for adding unit, integration, smoke, and end-to-end tests
  - Publishing weekly status and accomplishment reports
  - Enabled and curated GitHub Discussions for better distinction between questions and issues
  - Move k3s docs to [docs.k3s.io](http://docs.k3s.io), built on Docusaurus and served via GitHub Pages (no longer on SUSE Infrastructure)
  - Move [k3s.io website](http://k3s.io) onto Docusaurus/GitHub Pages (no longer on SUSE Infrastructure)
  - Continued focus on adding CI tests and improvements
  - Contributor process quality of life improvements
    - Issue and PR templates
    - Stalebot integration
  - Implemented [ADR](https://github.com/k3s-io/k3s/tree/master/docs/adrs) (Architecture Design Record) processes to add transparency and accountability to product design decisions 
  - Maintained and improved release cadence with upstream Kubernetes
    - 3-4 scheduled patch releases a month
    - Minor (aka feature) release once every 3-4 months
- Features and fixes
  - Adding support for several OS's (RHEL9, Ubuntu 22.04, OpenSUSE MicroOS)
  - Dual-stack support
  - IPv6-only support
  - Improved certificate rotation
  - Encryption key secrets rotation
  - Performance testing
  - Cleanup deprecated flags
  - Tooling to measure code coverage in e2e tests
  - Automated tooling to facilitate high quality releases
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
  - Continue to improve and expand automated testing
  - Continue to implement build and release improvements
  - Improve contributor quality of life improvements: docs, bots, and process improvements
  - Better engagement with upstream(s) via contributions (PR and KEP)
- Features
  - Improve VPN support
  - Implement security and compliance improvements 
    - Continue to improve SELinux support
    - FIPS 140-2/3 crypto validation (a U.S. government standard for cryptography)
    - Automated base image scanning and revving for CVE mitigation
  - Expand our support matrix
    - An ever-expanding list of supported OSes
  - Speculative (highly subject to change)
    - Embedded registry or image airgapping solution
    - Deeper “tiny edge” integration (with projects like Akri and smarter-device-manager)
    - Continue to bring down binary size whenever possible

We want to achieve the above roadmap while also maintaining cadence with upstream Kubernetes. This means:
- Delivering patch releases that lag mere days behind upstream (with CVEs ideally same day or day after)
- Delivering minor releases that lag at most 2-4 weeks behind upstream

### How can the CNCF help you achieve your upcoming goals?

The CNCF has been helpful thus far by contributing CI resources. As we work towards incubation, we'll likely reach out with specific, tactical asks as we encounter them.

### Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

We believe that we will meet the criteria for incubation in the coming months, if not sooner.  We have a very healthy contributor base outside of the core group as evidenced by devstats, and will continue to work to list adopters.  We hope to begin the process within the coming months, and look forward to working with the CNCF on it.


## Project Links
 - [web://k3s.io](https://k3s.io)
 - [github://k3s-io/k3s](https://github.com/k3s-io/k3s)
 - [slack://rancher-users#k3s](https://rancher-users.slack.com/archives/CGGQEHPPW)
