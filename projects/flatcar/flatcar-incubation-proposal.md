# **Flatcar Project Proposal**


## Name of Project

Flatcar


## Description

Flatcar Container Linux is a container optimized operating system (OS) that ships a minimal OS image, which includes only the tools needed to run containers. The OS is shipped through an immutable filesystem and includes automatic atomic updates that provides a minimal attack surface. 


### Overview

Flatcar is a minimal operating system for containerized workloads with the following key features:



* **Container Optimized**: Flatcar’s OS image includes only those packages needed to run containers. This minimalist approach reduces the amount of software to manage, as well as the potential attack surface.
* **Secure and Immutable File System**: The OS is deployed to an immutable filesystem (read-only /usr partition), which eliminates a whole category of security vulnerabilities that modify installed OS files (see, e.g. CVE-2019-5736).
* **Declaratively Provisioned**: Node configuration is defined in a YAML file which is applied on first boot. Thereafter, node configuration is generally treated as immutable, avoiding “configuration drift” and enabling management at scale.
* **Auto Updating (and Rollback)**: Updates are shipped as validated images and applied in an atomic operation. If the update fails, then the system automatically reverts to the previous image.


### Background and Project History

In 2013, CoreOS, Inc., launched the first container-optimized Linux, based on the insight that containerization of applications enabled a new approach to the host OS. Leveraging Google’s innovations with ChromeOS on the desktop, CoreOS promised security, reliability, and manageability at scale for servers running containers. It rapidly gained adoption and an enthusiastic user base. 

Following the acquisition of CoreOS by Red Hat in 2018, Kinvolk GmbH announced Flatcar Container Linux as a friendly fork of the original CoreOS Container Linux. When Red Hat declared the end-of-life of the original CoreOS, many users seeking continued support with a seamless upgrade path migrated to Flatcar. However, Flatcar’s benefits extended beyond simply being a drop-in replacement for CoreOS Container Linux, and it continued to attract new adopters. 

In 2021, Kinvolk was acquired by Microsoft. In the acquisition announcement, Microsoft’s Brendan Burns stated: “We know the CoreOS community has been on a winding journey over the years—we want to assure the Flatcar community that Microsoft and the Kinvolk team will continue to collaborate with the larger Flatcar community on the evolution of Flatcar Container Linux. Microsoft is committed to Flatcar Container Linux community development and will invest in working with the Flatcar community to create a growth path forward together.” 

Since then, the Flatcar team has been guided by the principle of Flatcar being a community-driven project, establishing regular community office hour meetings, opening up roadmap planning, simplifying the build process, opening up builds and CI to community contributors, shepherding community members through the contribution process, improving documentation for contributors, and more. 

The culmination of this process is to transfer Flatcar into a foundation home that will guarantee long-term independent governance, encouraging more in the community to get involved, and enabling even closer collaboration with related cloud native projects.


## Sponsor / Advisor from TOC

To be assigned by the TOC.


## Unique Identifier

flatcar


## License



* All Flatcar project-specific code is licensed under MIT, BSD (3-clause and 2-clause), or [Apache 2.0](https://github.com/cilium/cilium/blob/master/LICENSE).
* Unmodified redistributed OS and SDK components are licensed under their respective licenses, including e.g. the Linux kernel which is under GPL. This is standard for projects which ship Linux OS images as part of their build artifacts, and there is precedent in the CNCF for shipping Linux OS images - for example, several providers in the Kubernetes Cluster API project full Linux images (including Kubernetes bits) as part of its CI validation process, using the “image-builder” CNCF project (see e.g. [https://cluster-api-aws.sigs.k8s.io/topics/images/built-amis.html](https://cluster-api-aws.sigs.k8s.io/topics/images/built-amis.html)).
* The full list of packages and their respective licenses is automatically generated and shipped with each release.


## Maturity Level

Incubating


## Alignment with CNCF mission

Flatcar is aligned with the CNCF’s mission of making cloud native computing ubiquitous by providing a secure, immutable foundation for running scalable container-based applications in a wide variety of environments, such as public, private, and hybrid cloud. 

As one of the first cloud native technologies, it is natural for Flatcar to be part of the cloud native ecosystem alongside such projects as Kubernetes, etcd, and containerd. Flatcar is the “missing foundational layer” in the cloud native stack: if Flatcar is accepted into the CNCF, it will be possible to deploy a complete CNCF software stack, from OS through container runtime, orchestration, and all the higher level projects.

It is important to note that Flatcar is fundamentally a cloud-native project. It is not a general purpose operating system; its raison d’être is to host containers. It embeds etcd as a core part of the operating system (in fact, etcd was originally developed by CoreOS for the container Linux use case, and only subsequently adopted by Kubernetes). It ships with containerd as its container execution environment. And following a substantial investment by the Flatcar and ClusterAPI teams, it is now fully integrated with Cluster API for automated provisioning of Kubernetes nodes running Flatcar.  


## Source control repositories

[https://github.com/flatcar](https://github.com/flatcar)

[https://github.com/kinvolk/nebraska](https://github.com/kinvolk/nebraska) (note: this is planned to move under flatcar org)


## External dependencies

At the time of writing Flatcar packages and ships 288 upstream open source projects in the OS image, such as e.g. the Linux kernel, systemd, and containerd. An additional 268 upstream projects are packaged and shipped with the SDK container, resulting in 556 open source project dependencies total.

The list of OS image packages is published with every release. For the current stable it can be viewed here: [https://stable.release.flatcar-linux.net/amd64-usr/current/flatcar_production_image_packages.tx](https://stable.release.flatcar-linux.net/amd64-usr/current/flatcar_production_image_packages.txt)


## Roadmap and Issue Tracking

Flatcar uses GitHub extensively for tracking of issues (both raised by the development team and end users), and roadmap items.

Issues are raised and tracked in the [project’s issue tracker](http://issues.flatcar.org/). 

Items on the [project roadmap](https://github.com/orgs/flatcar/projects/7) are derived from requests coming from the community and discussion about which items will be part of the roadmap are done in the public, during office hours, release planning meetings, through Slack and Matrix chats and also in mailing lists forums.

Community contributions are welcomed, and new contributors in particular are encouraged to tackle issues with the tag “good first issue”.


## Maintainers and Contributors

The Flatcar codebase has contributions from over 110 people, across both its origins as CoreOS Container Linux and subsequent fork as Flatcar. In the last year, it has had 20 different contributors.

In the last 12 months, there have been 10,446 commits to the project; overall there were more than 49K commits since the beginning of the project. 

The set of maintainers with commit permissions is maintained in the [Maintainers list](https://github.com/flatcar/Flatcar/blob/main/MAINTAINERS.md)


## Infrastructure requirements (CI / CNCF Cluster)

Flatcar requires infrastructure for 

* building the distribution, for running system / acceptance tests
* publishing / serving OS images, with caches for global distribution
* providing OS updates.

The Flatcar build infrastructure is on Equinix Metal, as are the image servers and caching network (in total ~20 servers of various sizes). System / acceptance tests are run in a variety of vendor clouds (as required to test images specific to those environments) as well as on private cloud infrastructure (like qemu and VMWare) hosted on Equinix Metal’s bare metal servers. Public cloud testing infrastructure is allocated dynamically and freed after testing. Lastly, the update server runs on AWS.

The Equinix Metal infrastructure is currently fully sponsored by Equinix Metal. Upon acceptance of the project into CNCF, Equinix Metal has confirmed its intent to continue that sponsorship under the existing mechanism for sponsored CNCF infrastructure. Where additional infrastructure vendors are willing and able to provide resources to Flatcar as a CNCF project, that would be welcomed.

The full list of Infrastructure is listed in the Appendix. No additional infrastructure is expected to be required at this time. 


## Release methodology and mechanics

Each release goes through stabilization channels from Alpha -> Beta -> Stable. Furthermore the Flatcar project maintains an LTS channel derived from Stable on a regular (approximately annual) basis.


### Alpha channel

The Alpha channel is where new features, packages, and other major updates are introduced. New major versions are released to Alpha first; the major version number is set to the days since Sunday, June 30, 2013. Alpha follows a rapid release cadence - new major releases every 2-4 weeks - and is meant for development use. While each Alpha release undergoes full testing, new feature sets may not be production ready. Only when full feature sets are implemented, Alpha is promoted to Beta. On average, every second Alpha major release is promoted.


### Beta channel

The Beta channel is where a release series gestates and stabilizes in preparation for promotion to the stable channel. New Beta releases ship major new feature sets and are considered appropriate for users to deploy in production as a limited proportion of the total fleet: Beta “canaries” can ensure major new releases work well with user-specific environments and applications, enabling users to identify and report potential issues early, so they can be fixed for Stable. On average, every second major Beta release is promoted to Stable.


### Stable channel

The Stable channel is intended for production environments. A Stable major release is published roughly every 3 months and may receive a number of patch level releases during its lifetime. Support for a Stable release ends when a new major Stable release is published.

LTS

Long term support releases are derived from a “golden” Stable release that feature an LTS kernel as well as other LTS versions of core components. LTS follows an approximately annual cadence and receives patch level updates throughout its lifetime. Legacy LTS releases are supported until 6 months after a new LTS version is released, allowing sufficient migration time for users to switch to the new version.


## Ensuring Quality

Releases to each channel undergo quality assurance in multiple stages. While feedback on stability and robustness is key for maintaining quality, each stage uses different quality feedback mechanisms.



* Before a release:
    * Nightlies to catch build errors
    * Limited automated testing / smoke testing to catch major issues
* During release process (reviewed in GO / No GO meeting):
    * Full automated release tests to ensure release quality
    * Number of tests per platform to ensure sufficient coverage
        * with new tests being added based on bugs reported after a release
* After release is published:
    * Changes / features spend sufficient soak time in Alpha and Beta
    * Fixes based on bug reports, w/ tests added to test suite where applicable
    * Coverage is ensured based on Flatcar Instance Data report’s Alpha, Beta usage
        * Alpha / Beta canary coverage may be increased by reaching out to and work with users, customers, and partners when necessary


## Release Cadence

The cadence discussed below is based on experience with previous releases, particularly Alpha. That said, following the cadence is not a strict requirement, but serves as guidelines and for expectation management. There might be multiple “feature” or “bugfix” releases in a row.

**New releases are typically published at the beginning of the week** (Monday or Tuesday). Flatcar updates automatically; releasing at the end of the week would potentially create issues in users’ fleets right before, or during, the weekend.

Package updates and new features transition through a stabilization process:



* Introduced in Alpha (“feature” Alpha). Alpha releases must pass all tests.
* Feature-complete Alpha releases are promoted to Beta.
* Beta is considered fit for production and enable users to canary test new features.
* Stable is production-ready and has been tested by Beta users.
* LTS is based on stable and enables users to switch major versions only once a year.

**Alpha** releases have a high release cadence (target every two weeks) to swiftly iterate on new features.



1. Alpha introduces new major release versions and is usually not patched.
2. Alpha releases must pass the full set of tests on all vendors,

**Beta** releases stabilize features coming from Alpha. Beta is considered production-ready and enables users to detect issues with their workloads (e.g. by using canaries).



1. While Alpha might contain work-in-progress features, Beta releases are considered feature complete.
2. A new major Beta release is published roughly every month (i.e. every second Alpha goes Beta). The Beta channel may receive patch releases to address issues detected by users.
3. Feedback from users who run Beta canaries are addressed before a major release transitions to Stable, therefore ensuring their nodes on the Stable channel only receive releases that do not cause issues.

**Stable** releases are production-tested by users running Beta canaries.



1. A new major Stable is released roughly every 2 - 3 months (about every second Beta goes Stable).
2. Major Stable releases receive a number of patch releases during its lifetime.
3. A Stable release is considered supported until the next Stable is released.

**LTS** releases are based on Stable releases. LTS are for users who seek to minimize maintenance overhead.



1. A new major LTS is released every 12 months. LTS only receives patch level updates during its lifetime.
2. LTS has a lifetime of 18 months, leaving LTS users with 6 months to upgrade after a new major LTS has been released.


## Security


### OS Architecture security

Flatcar is a minimal Linux distribution designed with security in mind. Our OS image, which does not include a package manager thus cannot be extended at runtime, only ships the minimal set of services and applications that are required for running containers. OS images and update image are cryptographically signed, thus tamper-proof. The Flatcar OS binaries are stored on an immutable file system on a separate, read-only partition, so core operating system components cannot be overridden (either maliciously or accidentally). The OS binaries partition’s integrity is cryptographically verified by dm-verity on each boot.


### Supply chain security 

Ensuring the integrity of the OS is an essential element of the total supply chain security story for cloud native deployments. Flatcar takes this seriously, and has had a strong focus on ensuring integrity of images and updates since its inception. With the recent introduction of the SLSA framework, we had the opportunity to review compliance with those requirements. Today, Flatcar exceeds the SLSA Level 3 standard, and we are working to address the few identified gaps to meet Level 4.

In summary, the key points are as follows:

* Flatcar provides a reproducible build shipping an immutable OS image. 
* The whole distribution is built from source, ensuring integrity and provenance of all source files. Flatcar attests source inputs and binary outputs of all packages using cryptographic checksums in [in-toto / SLSA 0.1 JSON format](https://slsa.dev/provenance/v0.1).
* The build process produces signed artifacts, and has an offline signing procedure before artifacts are made available through the automated update infrastructure. In addition, an SBOM is delivered with every build.
* Flatcar is read-only and its integrity is verified at runtime on every machine.

The full details of Flatcar’s compliance with SLSA is included in the Flatcar documentation: https://www.flatcar.org/docs/latest/reference/supply-chain/.


### Security Process

The Flatcar maintainers team is relentlessly focused on keeping Flatcar secure for our users.

Security issues are identified through a variety of channels.



* We track security issues from a variety of sources, such as MITRE CVE entries, Gentoo GLSAs, public announcements, and public and private security mailing lists.
* Members of the maintainers team are subscribed to public and private (non-disclosure) security lists. 
* Users can report security issues via a private email channel ([security@flatcar-linux.org](mailto:security@flatcar-linux.org)) rather than having to use public github issues. This is documented on the project web page ([https://www.flatcar.org/security](https://www.flatcar.org/security)).

Once known, issues enter a private tracker with limited access. Our security team (a subset of the maintainers team) with on-call rotation assesses incoming reports and tracks issues, disclosure dates (if applicable), and mitigations and upstream patches. Embargoed patches are ingested in private builds which are published at the time the embargo lifts, supplying immediate remediations to our users.

On a weekly basis the security team also look into upstream security trackers and makes a severity assessment (independent from existing CVSS score) and then decide whether an urgent patch is required, or if the fix can wait until next release.

If accepted into CNCF, we would envisage moving over time towards an even more comprehensive security release process aligned with that of Kubernetes (https://github.com/kubernetes/committee-security-response/blob/main/security-release-process.md).


## Community


### Website

[https://flatcar.org](https://flatcar.org)


### Communication Channels



* Matrix (via element.io): [https://app.element.io/#/room/#flatcar:matrix.org](https://app.element.io/#/room/#flatcar:matrix.org)
* Slack: [https://kubernetes.slack.com/archives/C03GQ8B5XNJ](https://kubernetes.slack.com/archives/C03GQ8B5XNJ)
* Mailing list: [https://groups.google.com/g/flatcar-linux-user](https://groups.google.com/g/flatcar-linux-user)  ([mailto:flatcar-linux-user@googlegroups.com](mailto:flatcar-linux-user@googlegroups.com)) 
* GitHub: [https://github.com/flatcar/flatcar](https://github.com/flatcar/flatcar) 


### Community Meetings

More information on the meetings, dial-in instructions and links to recordings can be found in the [main repository’s readme](https://github.com/flatcar/flatcar#monthly-office-hours-and-release-planning).



* Office Hours (Monthly): Discusses news and happenings in the project and features live demos (often by guests) on Flatcar or related technologies.
* Release Planning (fortnightly): Follows the Alpha channel’s release cadence and discusses features of the current release as well as plans / items to include in the next release.


### Social media

Twitter: [twitter.com/flatcar](https://twitter.com/flatcar)

Mastodon: [hachyderm.io/@flatcar](https://hachyderm.io/@flatcar)


### Adopters

Flatcar has built up a broad community of adopters from large enterprises to small startups across a variety of industries, and is deployed at scale in diverse cloud and on-premise environments. For example:



* **Adobe** runs Flatcar on over 18,000 nodes in their fleet of Kubernetes clusters across multiple cloud providers and private data centers in 22 different regions worldwide.
* **Equinix Metal** [runs](https://kinvolk.io/blog/2021/02/case-study-equinix-metal-builds-on-flatcar/) its bare metal cloud control plane in Kubernetes atop Flatcar. They say: “The management overhead is massively reduced. Flatcar allows us to start thinking about our bare metal nodes like they’re cloud instances, we can think of them as more ephemeral. Also important to us is the atomic nature of Flatcar updates, which comes with the ability to rollback easily.”
* **Giant Swarm** [deploys](https://www.giantswarm.io/blog/time-to-catch-a-new-train-flatcar-linux) its managed Kubernetes service with Flatcar, having migrated from CoreOS Container Linux, enabling the deployments of customers such as Adidas and Vodafone.
* **Wipro Business Solutions** uses Flatcar Linux to power their hybrid/multi-cloud PostgreSQL containerized DBaaS platform. Each provisioned database is running on a dedicated lightweight stack with Flatcar as the foundational OS running on each database VM. In addition the DBaaS API itself and all supporting machines use Flatcar as well. They say: “Flatcar has proven to be a well-supported rocket solid OS with minimal attack surface, built in update mechanism and integrated docker daemon. Ignition brings in an early boot provisioning utility that perfectly adds to the full automation approach of the PostgreSQL DBaaS platform. We use it at scale on-prem with OpenStack cloud but also with public clouds like Google and Tencent.”

Users can share their adoption of Flatcar, with details of their use cases, in the [ADOPTERS.md](https://github.com/flatcar/Flatcar/blob/main/adopters.md) file in the Flatcar repo. 

GitHub stars: 435 (on the [https://github.com/flatcar/Flatcar](https://github.com/flatcar/Flatcar) repo)

Approximately 43,000 nodes access the public update server, an increase of 25% over the past 12 months. This only counts endpoints that contact the public server for updates. It excludes those behind a firewall, with updates disabled, using a private update server, or using full OS release images to update (node re-deployment / blue-green updates).


# Appendix


## Current Infrastructure



* Fixed server infrastructure on Equinix Metal
    * 5x s3.xlarge.x86 - Nightly build / release build servers
    * 1x c2.large.arm - Nightly build / release build server for ARM64 images and tests
    * 1x m3.large.x86 - build artifacts caching server for nightly / release builds
    * 1x t1.small.x86 \
1x c3.large.arm64 \
1x m3.large.x86 build / test servers for GitHub PRs
    * 1x c3.small.x86 w/ VSphere - VMWare vendor tests server
    * 2x s3.xlarge.x86 - OS image (“origin”) servers
    * 4x c3.small.x86 - OS image caching servers
* Fixed server infrastructure on AWS (for the Nebraska Update Server)
    * 2 x  t3.micro bastion EC2 instances (staging and live)
    * 1 x t3.micro staging Nebraska EC2 instance
    * 1 x t3.medium main Nebraska EC2 instance
    * 1 x ELB 
    * 1 x KMS (currently w/ with 22 keys)
    * 1 x db.t3.small staging RDS instance
    * 1 x db.t3.2xlarge prod RDS instance
    * 1x RDS backups
    * 1x Route53
    * 1x S3 bucket for data export
    * Kinesis firehose + WAF
    * CloudWatch alarms
* Dynamically allocated servers for vendor-tests of nightly builds and of release builds:
    * Equinix Metal
        * 12x c3.small.x86, \
12x c3.large.arm for max. 6h (usually ~1h) per build for full test suite
        * 1x c3.medium.x86, \
M3.large.x86, \
S3.xlarge.x86, \
n2.xlarge.x86 per build for ~10min for smoke-test
    * AWS
        * 24x t3.small, \
24x a1.large for max. 6h  (usually &lt; 30min) per build for full test suite
        * 1x m4.2xlarge per build for smoke-test (&lt; 5min).
    * Azure
        * 60 Standard_D2s_v4, \
60 Standard_D2pls_v5 for max 6h (usually &lt;1h) per build for full test suite
    * Digital Ocean
        * 24x s-2vcpu-2gb for max. 4h (usually &lt; 30min) per build for full test suite
    * Google Compute Engine
        * 12x n1-standard-2 for mak. 6h (usually &lt; 1h) per build for full test suite
