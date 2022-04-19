# KubeVirt Incubation Stage Review Proposal

## What is KubeVirt: Refresh

KubeVirt technology addresses the needs of development teams that have adopted or want to adopt Kubernetes but possess existing Virtual Machine-based workloads that cannot be easily containerized. More specifically, the technology provides a unified development platform where developers can build, modify, and deploy applications residing in both Application Containers as well as Virtual Machines in a common, shared environment.

Benefits are broad and significant. Teams with a reliance on existing virtual machine-based workloads are empowered to rapidly containerize applications. With virtualized workloads placed directly in development workflows, teams can decompose them over time while still leveraging remaining virtualized components as is comfortably desired.

## Statement on alignment with the CNCF mission

* Project is self-governing

* Is there a documented Code of Conduct that adheres to the CNCF guidelines?

[Kubevirt CoC](https://github.com/kubevirt/kubevirt/blob/main/CODE_OF_CONDUCT.md)

* Does the project have production deployments that are high quality and high-velocity?

Yes. See End Users below.

* Is the project committed to achieving the CNCF principles and do they have a committed roadmap to address any areas of concern raised by the community?

The KubeVirt project exists to aid users in moving from older frameworks to Cloud Native ones.

* Document that the project has a fundamentally sound design without obvious critical compromises that will inhibit potential widespread adoption

KubeVirt is built on top of standardized architectures and interfaces, particularly the Container Runtime Interface and the Libvirt API. Building on top of agreed cloud native and virtualization standards means that KubeVirt remains compatible with many other technologies, not only now but for years to come.

* Document that the project is useful for cloud native deployments & degree that it’s architected in a cloud native style

* Document that the project has an affinity for how CNCF operates and understand the expectation of being a CNCF project.

KubeVirt has been a CNCF project for two years.

## Sandbox Graduation Requirements

* Document that it is being used successfully in production by at least three independent end users which with focus on adequate quality and scope defined.

See End Users Below. Currently 10 organizations use KubeVirt, some of them for internal use (such as CloudFlare and Nvidia) and some incorporate it into products (such as Red Hat and SUSE).

* Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

* Demonstrate a substantial ongoing flow of commits and merged contributions.

For the last year, KubeVirt has averaged [more than 350 merged PRs per month](https://kubevirt.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) from more than 60 contributors. Of those, [10-20 per month](https://kubevirt.devstats.cncf.io/d/14/new-and-episodic-pr-contributors?orgId=1&from=now-1y&to=now-1w&var-period=m&var-repogroup_name=All) come from new contributors, showing ongoing interest in joining the project.

## Community

### Governance

* [Governance policy](https://github.com/kubevirt/community/blob/master/GOVERNANCE.md)
* [Membership policy](https://github.com/kubevirt/community/blob/master/membership_policy.md)

As the project is in the process of growing beyond its initial group of contributors, we are working on adopting additional written governance rules.

### SIGs

* [SIG list](https://github.com/kubevirt/community/blob/master/sigs.yaml)

Most of those SIGs are not yet independent committees. The Performance and Scale SIG has its own leadership and contributors, and both Documentation and Storage SIG should soon have the same. The community plan is to, as the community grows, foster dedicated contributors in each SIG area. As well as expanding our contribution pipeline, this will provide more opportunities for new contributors to rise to leadership of the project.

### Maintainers

* https://github.com/kubevirt/community/blob/master/MAINTAINERS.md

### GitHub source code repository

* About 50 contributors measured by multiple L or greater sized contributions
* 171 official members in the organization
* 2562 stars
* 615 forks
* 96 watched tags

### Communications
The project utilizes Google Groups as a [mailing list](https://groups.google.com/g/kubevirt-dev) where where users have a chance to interact with core developers and discuss general topics. There are currently (2021/06/01) 436 subscribers.  The [mailing list](https://groups.google.com/g/kubevirt-dev) typically receives 0-5 new threads per day.

The project holds a video conference every week where users have a chance to interact with core developers, discuss general topics and participate in bug triage. Recently the performance and scale working group started conducting a weekly meeting based on the topic. Both meetings are recorded and posted to the project YouTube channel. Meeting notes are also emailed to the general [mailing list](https://groups.google.com/g/kubevirt-dev) to keep the community informed.

The project uses two channels on the CNCF Kubernetes Slack server. #virtualization is used to handle general use topics. The channel hosts over 1000 users. #kubevirt-dev is used for developer oriented topics. The channel hosts 350 users. The channels typically receive 0-5 new threads per day.

Important announcements are relayed via [mailing list](https://groups.google.com/g/kubevirt-dev), website blog and Twitter.

KubeVirt advertises communications channels via https://kubevirt.io/community as well as project [README](https://github.com/kubevirt/kubevirt/blob/master/README.md).

### Adopters

The Kubevirt project tracks adopters via the [Adopters file](https://github.com/kubevirt/kubevirt/blob/main/ADOPTERS.md) files

### CNCF Sponsored Security Audit
This section is a place holder to be filled in when it occurs during incubation.

## Product information

### Feature design proposals
Design proposals to allow community members to gain feedback on their designs from the repo approvers before the community member commits to executing on the design. By going through the design process, developers gain a have a high level of confidence that their designs are viable and will be accepted.

* [design-proposals](https://github.com/kubevirt/community/tree/master/design-proposals)

### Release cadence

Kubevirt has an established and documented release process and cadence

* [Release process](https://github.com/kubevirt/kubevirt/blob/main/docs/release.md)
* [Release cadence](https://github.com/kubevirt/kubevirt/blob/main/docs/release.md#cadence-and-timeline)

### Delivered features

* [DONE] GA v1 API for core KubeVirt APIs
 * API v1 features need to rely on GA’ed Kubernetes entities, fully fledged (incl e.g. explain, validation)
 * An OpenAPI definition as the only source of truth for KubeVirt’s API
 * https://github.com/kubevirt/kubevirt/pull/3349
* [DONE] Zero downtime live updates
* [DONE] Stabilize bridge network binding
* [DONE] Disk hotplug
* [DONE] IPv6 support
* [DONE] Device passthrough support
* [DONE] Numa topology support
* [DONE] Memory metrics gathering support
* [DONE] Affinity / Anti-Affinity rules
* [DONE] Live-Migration support
* [DONE] Offline disk snapshots
* [DONE] SRIOV support
* [DONE] Dynamic SSH Key Injection
* [DONE] Multus support for multiple network interfaces attached to Virtual Machines
* [DONE] Dedicated prow deployment for CI functional tests and automation

### Future Roadmap

* [WIP] Non-root VMI Pods
* [WIP] Establish predictable community release and support patterns
* [WIP] Define a deprecation policy
* [WIP] Review and Revise User Guide
* [WIP] Templating mechanism for VMs
* [WIP] Monitoring and metrics standardization
* [WIP] Online Snapshots
* [WIP] CPU NUMA topology support
* [WIP] SSH proxy ingress support

Work is in progress to convert items into publicly advertised [GitHub Milestones](future roadmap is being mapped out at https://github.com/kubevirt/kubevirt/milestones)

## Incubation Stage Requirements

The KubeVirt project maintainers propose that KubeVirt move to Incubation based on:

* Use in production by 3 significant end users
* A healthy number of committers and a growing committer base in addition, to a healthy online community.
* Demonstrating a substantial ongoing flow of commits and merged contributions that focused on delivering a defined project roadmap and integrations.
* A clear versioning scheme with dev and stable releases.
