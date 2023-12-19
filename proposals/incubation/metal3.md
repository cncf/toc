# Metal3 proposal for CNCF incubation

## Table of Contents

- [Metal3 proposal for CNCF incubation](#metal3-proposal-for-cncf-incubation)
   - [Table of Contents](#table-of-contents)
   - [About Metal3](#about-metal3)
   - [Sandbox Acceptance](#sandbox-acceptance)
   - [Progress Since Sandbox](#progress-since-sandbox)
   - [Technical Advances](#technical-advances)
      - [Cluster-api-provider-metal3](#cluster-api-provider-metal3)
      - [Baremetal operator](#baremetal-operator)
      - [ip-address-manager](#ip-address-manager)
      - [Ironic-image](#ironic-image)
      - [Other Major Contributions](#other-major-contributions)
   - [Adopters](#adopters)
   - [Maintainers](#maintainers)
   - [Community size](#community-size)
   - [Security processes](#security-processes)
   - [Statement on Alignment with the CNCF Mission](#statement-on-alignment-with-the-cncf-mission)
   - [Alignment with Other CNCF Projects](#alignment-with-other-cncf-projects)

## About Metal3

The Metal3 project (pronounced: “Metal Kubed”) provides components for baremetal
host management with Kubernetes. User can enroll user's baremetal machines,
provision operating system images, and then, if the user likes, deploy Kubernetes
clusters to them. From there, operating and upgrading user's Kubernetes clusters
can be handled by Metal3. Moreover, Metal3 is itself a Kubernetes application,
so it runs on Kubernetes, and uses Kubernetes resources and APIs as its
interface.

Metal3 is one of the providers for the Kubernetes sub-project Cluster API.
Cluster API provides infrastructure agnostic Kubernetes lifecycle management,
and Metal3 brings the baremetal implementation.

Furthermore Metal3 utilizes a component from the OpenStack ecosystem called
Ironic. Ironic is used for booting and installing machines. Metal3 handles the
installation of Ironic as a standalone component (there’s no need to bring along
the rest of OpenStack). Ironic is supported by a mature community of hardware
vendors and supports a wide range of baremetal management protocols which are
continuously tested on a variety of hardware. Ironic and Metal3 community work
closely with a mutual CI which ensures Metal3 can provision machines, no matter
the brand of hardware.

In summary, user can write Kubernetes manifests representing user's hardware and
desired Kubernetes cluster layout. Then Metal3 can:

- Discover hardware inventory
- Configure BIOS and RAID settings on hosts
- Optionally clean a host’s disks as part of provisioning
- Install and boot an operating system image of own choice
- Deploy Kubernetes
- Upgrade Kubernetes or the operating system in clusters with a
 non-disruptive rolling strategy
- Automatically remediate failed nodes by rebooting them and removing them from
 the cluster if necessary

## Sandbox Acceptance

The Metal3 project was accepted into CNCF sandbox on 8th of September 2020. [PR link](https://github.com/cncf/toc/pull/408).

## Progress Since Sandbox

The Metal3 project has advanced a lot since its acceptance to sandbox.
Repositories that have lost their importance have been archived after extensive
community discussions.

As it stands today the project revolves around four core component plus a
development environment, a documentation repository , a community repository,
a website repository and an infrastructure as code (IAAC) repository that
handles the CI/CD processes for the project.

The four core components are:

- [baremetal-operator](https://github.com/metal3-io/baremetal-operator)
- [cluster-api-provider-metal3](https://github.com/metal3-io/cluster-api-provider-metal3)
- [ip-address-manager](https://github.com/metal3-io/ip-address-manager)
- [ironic-image](https://github.com/metal3-io/ironic-image)

In addition Metal3 also has the following repositories:

- Metal3 Community governance documents are located at [community](https://github.com/metal3-io/community)
 repository.
- The development environment can be found at: [metal3-dev-env](https://github.com/metal3-io/metal3-dev-env)
 repository.
- The design documents and the user-guide is located at [metal3-docs](https://github.com/metal3-io/metal3-docs)
 repository.
- The website's source is located at [metal3-io.github.io](https://github.com/metal3-io/metal3-io.github.io)
 repository.
- The CI/CD IAAC repository is at [project-infra](https://github.com/metal3-io/project-infra)
 repository.
- Metal3 documentations is located at [metal3-docs](https://github.com/metal3-io/metal3-docs)
 repository. Metal3 user guide is located at this [link.](https://book.metal3.io/)

## Technical Advances

Aggregated list of technical advances since sandbox acceptance:

### Cluster Api Provider Metal3

- Node reuse feature [#169](https://github.com/metal3-io/cluster-api-provider-metal3/pull/169)
- Remediation controller [#157](https://github.com/metal3-io/cluster-api-provider-metal3/pull/157)
- Add v1beta1 types and related changes [#342](https://github.com/metal3-io/cluster-api-provider-metal3/pull/342)
- Support IP Reuse for BMHs using Preallocations [#656](https://github.com/metal3-io/cluster-api-provider-metal3/pull/656)
- Introduce additional providerID format [#563](https://github.com/metal3-io/cluster-api-provider-metal3/pull/563)
- Adoption of CAPI e2e framework Initial [#194](https://github.com/metal3-io/cluster-api-provider-metal3/pull/194)
- Add live-iso support to CAPI Metal3 provider [#189](https://github.com/metal3-io/cluster-api-provider-metal3/pull/189)
- Open up CAPM3 for any control plane provider [#1311](https://github.com/metal3-io/cluster-api-provider-metal3/pull/1311)
- Add TLS configuration flags for CAPM3 [#1041](https://github.com/metal3-io/cluster-api-provider-metal3/pull/1041)
- Introduce ReconcileError with Transient and Terminal Error type [#910](https://github.com/metal3-io/cluster-api-provider-metal3/pull/910)
- Add flags for configuring rate limits [#959](https://github.com/metal3-io/cluster-api-provider-metal3/pull/959)
- Implement CAPI IPAM contract support [#769](https://github.com/metal3-io/cluster-api-provider-metal3/pull/769)
- Enable concurrency in metal3 controllers [#884](https://github.com/metal3-io/cluster-api-provider-metal3/pull/884)
- Add explicit securitycontext to controllers [#822](https://github.com/metal3-io/cluster-api-provider-metal3/pull/822)
- Introduce separate go modules for e2e tests to allow them to be imported
independently [#800](https://github.com/metal3-io/cluster-api-provider-metal3/pull/800)
- Add webhook readiness and health check [#279](https://github.com/metal3-io/cluster-api-provider-metal3/pull/279)
- Add AGE column to metal3 CRDs [#289](https://github.com/metal3-io/cluster-api-provider-metal3/pull/289)
- Automate releasing process with GitHub actions [#308](https://github.com/metal3-io/cluster-api-provider-metal3/pull/308)
- Implement externally managed annotation [#224](https://github.com/metal3-io/cluster-api-provider-metal3/pull/224)
- Add pivoting test based on the e2e test framework [#232](https://github.com/metal3-io/cluster-api-provider-metal3/pull/232)
- Add internal logging mechanism for e2e test framework [#259](https://github.com/metal3-io/cluster-api-provider-metal3/pull/259)
- Add BMO inject-ca-from custom patching [#223](https://github.com/metal3-io/cluster-api-provider-metal3/pull/223)

### Baremetal Operator

- Authenticating to Ironic [#601](https://github.com/metal3-io/baremetal-operator/pull/601)
- Added RAID configuration support [#134](https://github.com/metal3-io/baremetal-operator/pull/134)
- Added select vendor-independent BIOS configuration options [#63](https://github.com/metal3-io/baremetal-operator/pull/63)
- Support automatic secure boot [#161](https://github.com/metal3-io/baremetal-operator/pull/161)
- Add boot-iso API to BareMetalHost [#150](https://github.com/metal3-io/baremetal-operator/pull/150)
- Optimized concurrent provisioning performance and tested at scale up to 1000 Hosts PR [#725](https://github.com/metal3-io/baremetal-operator/pull/725)
- Added custom agent image controller [#183](https://github.com/metal3-io/baremetal-operator/pull/183)
- Release process for BMO [#1150](https://github.com/metal3-io/baremetal-operator/pull/1150)
- HardwareData custom resource for host inspection data [#1099](https://github.com/metal3-io/baremetal-operator/pull/1099)
- PreprovisioningImage API and integration [#936](https://github.com/metal3-io/baremetal-operator/pull/936)
- New API and integration of hostFirmwareSettings and firmwareSchema resources [#901](https://github.com/metal3-io/baremetal-operator/pull/901)
 , [#938](https://github.com/metal3-io/baremetal-operator/pull/938)
- Enable RAID for Redfish-based iDRAC driver flavor [#1095](https://github.com/metal3-io/baremetal-operator/pull/1095)
- Reboot API Implementation [#424](https://github.com/metal3-io/baremetal-operator/pull/424)
- Add physicalDisks and controller parameters to HardwareRAID [#1062](https://github.com/metal3-io/baremetal-operator/pull/1062)
- Implement explicit reboot mode options [#795](https://github.com/metal3-io/baremetal-operator/pull/795)
- Add live-iso support [#759](https://github.com/metal3-io/baremetal-operator/pull/759)
- Add support for detached annotation [#827](https://github.com/metal3-io/baremetal-operator/pull/827)
- Add support for ironic custom deploy interface [#884](https://github.com/metal3-io/baremetal-operator/pull/884)
- Implement BMH provisioning E2E test [#1403](https://github.com/metal3-io/baremetal-operator/pull/1403)
- Implemented support adding pre-provisining Network data for Ironic node [#1380](https://github.com/metal3-io/baremetal-operator/pull/1380)
- Explicit way to auto-detect disk image ChecksumType [#1480](https://github.com/metal3-io/baremetal-operator/pull/1480)
- Support both Ironic and Inspector inventory formats [#1450](https://github.com/metal3-io/baremetal-operator/pull/1450)
- Expose FirmwareInterface in Node/BMC [#1441](https://github.com/metal3-io/baremetal-operator/pull/1441)
- API version negotiation [#1435](https://github.com/metal3-io/baremetal-operator/pull/1435)

### IP Address Manager

It's a comprehensive tool to manage static IP address allocations in
Cluster API Provider Metal3. It has its own controller and quite a few new
features are added in this repo.

- Add TLS configuration flags for IPAM [#248](https://github.com/metal3-io/ip-address-manager/pull/248)
- Add flags for configuring rate limits [#238](https://github.com/metal3-io/ip-address-manager/pull/238)
- Enable concurrency in ipam controllers [#213](https://github.com/metal3-io/ip-address-manager/pull/213)
- Add explicit securitycontexts to controller [#195](https://github.com/metal3-io/ip-address-manager/pull/195)
- Default LeaderElectionResourceLock to leases [#85](https://github.com/metal3-io/ip-address-manager/pull/85)
- Automate releasing process with GitHub actions [#66](https://github.com/metal3-io/ip-address-manager/pull/66)
- Add AGE column to IPAM CRDs [#51](https://github.com/metal3-io/ip-address-manager/pull/51)

### Ironic Image

- Custom iPXE firmware building and deployment support [#429](https://github.com/metal3-io/ironic-image/pull/429)

### Other Major Contributions

- New repository with iPXE firmware builder container [utility-images](https://github.com/metal3-io/utility-images)
- Improve end user/developer documentation
- Added a [user-guide](https://book.metal3.io/) in form of the Metal3 Book
- Added Metal3 community repository community  

## Adopters

The project's adopters list can be found here.

Companies that are using Metal3 in production:

- **IKEA IT AB:** "IKEA IT AB uses Metal3 to handle Bare Metal provisioning and
 lifecycle management in its CAPI-Based baremetal cloud infrastructure."

- **Ericsson:** "We have chosen Metal3 as a baremetal provisioner for Ericsson’s
 Cloud Container Distribution since Metal3 is a forerunner when it comes to
 Kubernetes on top of bare metal servers. Besides the robust features of
 Metal3, a very involved community and clearly defined roadmap have been
 the key factors for us to choose Metal3 as a core component."

- **Red Hat:** Red Hat's OpenShift distribution includes Metal3 as part of its
 solution for automating the deployment of baremetal clusters.

- **Fujitsu:** "As a server vendor, we are developing and using Metal3 to
 provide Fujitsu servers as Kubernetes baremetal nodes."

- **Deutsche Telekom:** "We have various Telco applications that require
 baremetal infrastructure. We rely on metal3 for provisioning since it supports
 layer 3 only deployments which makes it easier to use in our complex
 networking setup than other options."
- **PITS Global Data Recovery Services:** “The Metal3 is used to manage
 highly-loaded internal infrastructure providing reliable and flexible k8s
 solutions.”
- **SUSE:** “Metal3 is used for automated baremetal deployment as part of the
 SUSE Edge solution.”

Other open source projects that are utilizing Metal3:

- [Airship](https://github.com/airshipit)
- [DT Technik "Das SCHIFF"](https://github.com/telekom/das-schiff)
- [OpenShift](https://github.com/openshift)
- [Sylvia-Projects](https://gitlab.com/sylva-projects/sylva-elements/helm-charts/sylva-capi-cluster)
- [Kanod](https://gitlab.com/Orange-OpenSource/kanod)
- [Medik8s](https://github.com/medik8s?q=metal3&type=all&language=&sort=)

Furthermore maintainers have noticed that there are many more companies
that are actively using Metal3, as representatives of said companies are
opening issues, writing proposals and in general contribute to the project
but have not registered themselves as official adopters.

## Maintainers

A Metal3 maintainer is a participant in the project with the ability to approve
changes proposed to a given repository. Approval access grants a broad ability
to affect the progress of the project as presented by its most important artifact,
the code and related resources. As such it represents a significant level of
trust in an individual's commitment to working with other maintainers and
the community at large for the benefit of the project. Metal3 project also
documents the process to become Maintainer and/or the process to move up in the
Contribution Ladder as well as clearly points out Contributor Roles in the
[community repository](https://github.com/metal3-io/community).

The following table contains the [Metal3's maintainers](https://github.com/metal3-io/community/tree/main/maintainers).

| Approver           | GitHub ID      | Affiliation                    |
| ------------------ | -------------- | ------------------------------ |
| Andrea Fasano      | andfasano      | Red Hat                        |
| Bob Fournier       | bfournie       | Red Hat                        |
| Derek Higgins      | derekhiggins   | Red Hat                        |
| Dmitry Tantsur     | dtantsur       | Red Hat                        |
| Riccardo Pittau    | elfosardo      | Red Hat                        |
| Furkat Gofurov     | furkatgofurov7 | SUSE                           |
| Honza Pokorny      | honza          | Red Hat                        |
| Himanshu Roy       | hroyrh         | Red Hat                        |
| Iury Gregory       | iurygregory    | Red Hat                        |
| Kashif Khan        | kashifest      | Ericsson Software Technology   |
| Lennart Jern       | lentzi90       | Ericsson Software Technology   |
| Mohammed Boukhalfa | mboukhalfa     | Ericsson Software Technology   |
| Adam Rozman        | Rozzii         | Ericsson Software Technology   |
| Moshiur Rahman     | smoshiur1237   | Ericsson Software Technology   |
| Tuomo Tanskanen    | tuminoid       | Ericsson Software Technology   |
| Zane Bitter        | zaneb          | Red Hat                        |

We have also introduced the option of moving people to emeritus approvers and
reviewers list for people who became less active in the review/maintenance
process.

## Community

Metal3 has had [377 individual contributors](https://metal3.devstats.cncf.io/d/5/companies-table?orgId=11)
on the project’s GitHub repositories from [57 different organizations.](https://metal3.devstats.cncf.io/d/5/companies-table?orgId=1)

The project has [778 Stargazers/Watchers](https://metal3.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1)
across its repositories.

[DevStats graph](https://metal3.devstats.cncf.io/d/8/dashboards?orgId=1&editPanel=2&from=now-1y&to=now-1h)
shows that we have a fairly consistent number of contributions throughout the
past year. [CloMonitor stats](https://clomonitor.io/projects/cncf/metal3-io)
also shows Metal3 is doing fairly well in CNCF best practice metrics.

In addition, we have a very high number of downloads of our container images
from quay.io. Since the download statistics is not publicly visible,
we are not adding any link.

As an example of our traction the aggregated number of container image
downloads from the Metal3's container repository from 20-11-2023 to
19-12-2023 was on average 14830 and the peak was 40,584 during a single day.

## Security processes

Metal3 has clearly documented security process. Generic process is documented
[here](https://book.metal3.io/security_policy.html) and each of the repositories
have
[SECURITY.md](https://github.com/metal3-io/.github/blob/main/SECURITY.md) and
[SECURITY_CONTACTS](https://github.com/metal3-io/.github/blob/main/SECURITY_CONTACTS)
file at the root of the repository, clearly detailing the responsible disclosure
process. The security policy itself details the process of getting a vulnerability
fixed, the roles of security team and the security lead, and timelines for
any releases.

Metal3 has a security team, lead by a Security Lead, who is a full-time
security-first engineer working at EST. Security team can be emailed at
[mailto:metal3-security@googlegroups.com](metal3-security@googlegroups.com).
Security team has processed three responsible disclosures in 2023. Two of the
disclosures have led into published security advisory:
[CVE-2023-30841](https://github.com/metal3-io/baremetal-operator/security/advisories/GHSA-9wh7-397j-722m)
and
[CVE-2023-40585](https://github.com/metal3-io/ironic-image/security/advisories/GHSA-jwpr-9fwh-m4g7)

Metal3 also has robust
[releasing process](https://github.com/metal3-io/metal3-docs/blob/main/processes/releasing.md)
and documentation to ensure the security fixes can be delivered for all the
supported software versions. The core components all have their respective
release documentation detailing the release process for that repository:
[BMO](https://github.com/metal3-io/baremetal-operator/blob/main/docs/releasing.md),
[CAPM3](https://github.com/metal3-io/cluster-api-provider-metal3/blob/main/docs/releasing.md),
[IPAM](https://github.com/metal3-io/ip-address-manager/blob/main/docs/releasing.md)
and
[ironic-image](https://github.com/metal3-io/ironic-image/blob/main/docs/releasing.md).

GitHub Actions based release automation is in place, e.g.
[CAPM3](https://github.com/metal3-io/cluster-api-provider-metal3/blob/main/.github/workflows/release.yaml)
and
[release verification script](https://github.com/metal3-io/cluster-api-provider-metal3/blob/main/hack/verify-release.sh)
ensures software has up to date dependencies and no publicly known
vulnerabilities exist in the repository.

During the sandbox phase, Metal3 has put a lot of effort into securing the supply
chain by pinning down Dockerfiles, GitHub workflows, Go modules and other
dependencies as well as ensuring GitHub configuration for the whole organization
is consistent and secure, with branch protection in place and not allowing code
merging without proper reviews and tests. Dependabot has been adopted to
bump dependencies and keep an eye on the published vulnerabilities in the
dependencies, in addition to pre-existing Renovate bot configurations.

Metal3 has been closely tracking its OpenSSF Scorecard evaluations for all of
the repositories, and while there is still plenty of work to do to achieve
very high scores, the average score for a Metal3 repository has risen 1-2 points
during 2023. We are also working towards
[OpenSSF Best Practices badge](https://www.bestpractices.dev/en), formerly known
as CII badge, as well as implementing the other missing checks, like Fuzzing
and Signed Releases.

Metal3 has also improved its security posture by moving into rootless containers
with any elevated permissions provided by capabilities, implementing very
limiting default security context for all containers and defaulting to baseline
PSA where applicable. We have also added configuration options to enforce use
of TLS 1.3 only.

Metal3 also has implemented a (private) security roadmap detailing future
improvements into security posture, process and documentation.

## Statement on Alignment with the CNCF Mission

As cloud native technologies become the norm for organizations to deploy and
manage their applications, Metal3 aims to apply similar patterns such as
declarative APIs to allow administrators to manage their underlying
infrastructure.

The Metal3 project believes in CNCF’s mission to make cloud-native technologies
accessible for all, that is why the Metal3 project is trying to make physical
infrastructure management more accessible to the cloud native ecosystem and to
the worldwide open source software community.

## Alignment with Other CNCF Projects

The Metal3 stack integrates into Kubernetes Cluster API which is a CNCF project.
Metal3's IPAM is also used by cluster-api-provider-vsphere's CI.
Cluster-api-provider-vsphere is a sub-project of sig-cluster-lifecycle and is a
CNCF supported project.
