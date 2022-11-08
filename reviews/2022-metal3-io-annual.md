# Metal3-io 2022 Annual Review

## Metrics

This [DevStats](https://metal3.devstats.cncf.io/d/8/dashboards?orgId=1&from=now-1y&to=now-1h&editPanel=2)
graph shows that we have a fairly consistent number of contributions
throughout the year.

### Downloads per repository
We have fairly high number of downloads of our release manifests and
artifacts. We showcase two examples below.

[Cluster API Provider Metal3](https://hanadigital.github.io/grev/?user=metal3-io&repo=cluster-api-provider-metal3)

[IP Address Manager](https://hanadigital.github.io/grev/?user=metal3-io&repo=ip-address-manager)

In addition, we have a very high number of downloads of our container images
from [quay.io](https://quay.io/). Since the download statistics is not
publicly visible, we are not adding any link.

As an example of our traction the aggregated number of container image
downloads from the Metal3's container repository from `10-26-2022 to
11-2-2022` was on `average 12193` and the peak was `15755 during a single day`.

![](https://raw.githubusercontent.com/Nordix/metal3-clusterapi-docs/master/2022-metal3-io-annual-review/2022-metal3-io-container-downloads.png)


## Current maintainers

The following table contains both the project's approvers and reviewers combined.

| Reviewer/Approver  | GitHub ID      | Affiliation                    |
| ------------------ | -------------- | ------------------------------ |
| Andrea Fasano      | andfasano      | Red Hat                        |
| Bob Fournier       | bfournie       | Red Hat                        |
| Derek Higgins      | derekhiggins   | Red Hat                        |
| Dmitry Tantsur     | dtantsur       | Red Hat                        |
| Riccardo Pittau    | elfosardo      | Red Hat                        |
| Honza Pokorny      | honza          | Red Hat                        |
| Himanshu Roy       | hroyrh         | Red Hat                        |
| Zane Bitter        | zaneb          | Red Hat                        |
| Iury Gregory       | iurygregory    | Red Hat                        |
| Kashif Khan        | kashifest      | Ericsson Software Technology   |
| Furkat Gofurov     | furkatgofurov7 | Ericsson Software Technology   |
| Mohammed Boukhalfa | mboukhalfa     | Ericsson Software Technology   |
| Lennart Jern       | lentzi90       | Ericsson Software Technology   |
| Adil Ghaffar       | adilGhaffarDev | Ericsson Software Technology   |
| Moshiur Rahman     | smoshiur1237   | Ericsson Software Technology   |
| Sunnatillo Samadov | Sunnatillo     | Ericsson Software Technology   |
| Adam Rozman        | Rozzii         | Ericsson Software Technology   |

Complete list of approvers across different repositories can be
found [here](https://github.com/metal3-io/metal3-docs/blob/master/maintainers/ALL-OWNERS).
We have introduced the option of moving people to emeritus approvers and
reviewers list for people who became less active in the review/maintenance
process.

## Adoption

The project's adopters list can be found [here](https://github.com/metal3-io/metal3-docs/blob/main/ADOPTERS.md).

**Companies that are using Metal3 in production:**

`IKEA IT AB`: "IKEA IT AB uses Metal3 to handle Bare Metal provisioning and
lifecycle management in its CAPI-Based bare metal cloud infrastructure."

`Ericsson`: "We have chosen Metal3 as a bare metal provisioner for Ericsson’s
Cloud Container Distribution since Metal3 is a forerunner when it comes to
Kubernetes on top of bare metal servers. Besides the robust features of
Metal3, a very involved community and clearly defined roadmap have been
the key factors for us to choose Metal3 as a core component."

`Red Hat`: Red Hat's OpenShift distribution includes Metal3 as part of its
solution for automating the deployment of bare metal clusters.

`Fujitsu`: "As a server vendor, we are developing and using Metal3 to
provide Fujitsu servers as Kubernetes bare metal nodes."

`Deutsche Telekom`: "We have various Telco applications that require bare
metal infrastructure. We rely on metal3 for provisioning since it supports
layer 3 only deployments which makes it easier to use in our complex
networking setup than other options."

**Other open source projects that are utilizing Metal3:**

[Airship](https://www.airshipit.org)

[DT Technik "Das SCHIFF"](https://github.com/telekom/das-schiff)

[OpenShift](https://github.com/openshift/)

In addition to the official adopters there are projects like
[Medik8s](https://www.medik8s.io/) who are providing options
for users to use Metal3 specific features like Metal3 based
[remediation](https://www.medik8s.io/remediation/metal3/metal3/).

Furthermore maintainers have noticed that there are many more companies
that are actively using Metal3, as representatives of said companies are
opening issues, writing proposals and in general contribute to the project
but have not registered themselves as official adopters.

## Project performance against its goals since the last review

This is our second annual review report after joining CNCF. We have included
here some of the major achievements that took place during the past year.
One of the major achievement was to release beta releases of Cluster API
Provider Metal3 and IP Address Manager. With this, Cluster API Provider
Metal3 API has become stable and production ready.

In last years report we listed some goals for the year that has now passed.
Here is a summary of those goals and their status:

**Goals achieved in the past year**
- Split the current BareMetalHost CRD inspection data into a new CRD
  - :heavy_check_mark: The inspection data now [has its own CRD (HardwareData)](https://github.com/metal3-io/baremetal-operator/pull/1099).
- Add support for vendor-specific firmware configuration
  - :heavy_check_mark: Implemented in [#901](https://github.com/metal3-io/baremetal-operator/pull/901/) and [#938](https://github.com/metal3-io/baremetal-operator/pull/938)
- Freeze the BareMetalHost CRD
  - :x: We are not quite ready to do this yet.
- Start releasing process for BMO
  - :heavy_check_mark: A release process has been [defined and documented](https://github.com/metal3-io/baremetal-operator/blob/84a7ea97d75d560b35cb205bb63f4163c02c6f59/docs/releasing.md).
- Establish a formal security reporting process
  - :heavy_check_mark: Security Process has been [introduced.](https://github.com/metal3-io/metal3-docs/pull/293)
- CAPM3 v1beta1 migration
  - :heavy_check_mark: The `v1beta1` migration was done for the [v1.1.0 release](https://github.com/metal3-io/cluster-api-provider-metal3/releases/tag/v1.1.0).
- Improve end-to-end testing
  - :heavy_check_mark: Adopted the CAPI e2e framework. ([Initial PR](https://github.com/metal3-io/cluster-api-provider-metal3/pull/194))

**Documentation improvements**
- Improve end user documentation
  - :heavy_check_mark: Added a user-guide in form of the [Metal3 Book](https://metal3io.netlify.app/)

### More detailed accomplishments per repository

- **metal3-docs**
    - Introduce [Metal3 Book](https://metal3io.netlify.app/) for comprehensive user-guide across the
      organization.

- **cluster-api-provider-metal3**
    - Add v1beta1 types and related changes [#342](https://github.com/metal3-io/cluster-api-provider-metal3/pull/342)
    - Support IP Reuse for BMHs using Preallocations [#656](https://github.com/metal3-io/cluster-api-provider-metal3/pull/656)
    - Introduce additional providerID format [#563](https://github.com/metal3-io/cluster-api-provider-metal3/pull/563)
    - Adoption of CAPI e2e framework [Initial PR](https://github.com/metal3-io/cluster-api-provider-metal3/pull/194)
    - Add live-iso support to CAPI Metal3 provider [#189](https://github.com/metal3-io/cluster-api-provider-metal3/pull/189)

- **baremetal-operator**
	- Release process for BMO [#1150](https://github.com/metal3-io/baremetal-operator/pull/1150)
	- HardwareData custom resource for host inspection data [#1099](https://github.com/metal3-io/baremetal-operator/pull/1099)
	- PreprovisioningImage API and integration [#936](https://github.com/metal3-io/baremetal-operator/pull/936)
	- New API and integration of hostFirmwareSettings and firmwareSchema resources [#901](https://github.com/metal3-io/baremetal-operator/pull/901/), [#938](https://github.com/metal3-io/baremetal-operator/pull/938)
	- Enable RAID for Redfish-based iDRAC driver flavor [#1095](https://github.com/metal3-io/baremetal-operator/pull/1095)
	- Reboot API Implementation [#424](https://github.com/metal3-io/baremetal-operator/pull/424)
	- Add physicalDisks and controller parameters to HardwareRAID [#1062](https://github.com/metal3-io/baremetal-operator/pull/1062)
	- Implement explicit reboot mode options [#795](https://github.com/metal3-io/baremetal-operator/pull/795)
	- Add live-iso support [#759](https://github.com/metal3-io/baremetal-operator/pull/759)
	- Add support for detached annotation [#827](https://github.com/metal3-io/baremetal-operator/pull/827)
	- Add support for ironic custom deploy interface [#884](https://github.com/metal3-io/baremetal-operator/pull/884)

- **ip-address-manager**
  -  It's a comprehensive tool to manage static IP address allocations in
     Cluster API Provider Metal3. It has its own controller and quite a few new
     features are added in this repo.

 **NOTE**
`A few of the above mentioned features is a result of work done not only in the previous year but before that time period also. We have put it here to highlight these since we didn't put them in the previous review.`

### General Improvements

  - Added kubesec for security analysis of Kubernetes manifest files.
  - Added gosec for security analysis of go files.
  - Enhanced e2e and full stack CI building process.
  - Bi-weekly triaging of issues.
  - Restrict interface usage of ironic for better security.
  - Rename all the ‘master’ branches to ‘main’.

## Future Goals

  - One of the main goals for 2022 was to help adoption of Metal3 thus the
    Metal3 community has started writing a [User-guide](https://metal3io.netlify.app).
    The community plans to extend the guide continuously and also provide
    separate tutorials that will showcase the deployment process on bare metal
    environment that is as close to a production environment as possible.

  - In 2022 certain aspects of Metal3 security has been improved, but there
    is more planned specially related to the topics of container security
    and related to the testing of Metal3 security features during e2e tests.

  - There are improvements done on the CI infrastructure of the Metal3
    project, but there are more to come in terms of securing and speeding the
    test processes to help contributors receive feedback faster.

  - There is long-running cooperation between the `Openstack-Ironc` and the
    `Metal3-io` communities, thus in the spirit of cooperation there is an
    ongoing effort to run Metal3 CI jobs against [Ironic upstream](https://review.opendev.org/c/openstack/ironic/+/863873).

  - There is an ongoing process that will continue in the future to align
    the Metal3 e2e test better with the CAPI e2e tests as well as to rework
    the developer environment that is currently used to better fit the CI
    use cases.

  - There are ongoing discussion related to expanding the features of BMO,
    CAPM3 and IPAM e.g.:
    - Adding better support for Ironic root device hint operators.
    - Supporting use-cases where the user would like to manage the lifecycle
      of bare metal machines from a K8S cluster but don't want to include
      the machines as nodes in the cluster.
    - Adding support for the CAPI IPAM contract.
    - Adding support for selective disk cleaning in baremetal-operator.
    - Adding support for multitenancy.

## How can the CNCF help you achieve your future goals?

1. We would like to ask CNCF to guide us through the fuzzing process using [CNCF fuzzing](https://github.com/cncf/cncf-fuzzing).

2. We would like to ask CNCF to help Metal3-io project to improve its visibility
   in any way possible e.g.:
     - finding advocate(s) for the project
     - facilitating cooperation between Metal3 and other FOSS projects
     - or providing with Metal3-io with any other opportunity to improve its
       traction/visibility


## Security processes

The Metal3-io project's general security policy is located [here](https://metal3io.netlify.app/security_policy.html).

In addition to the project's security policy, all the active repositories
should have a `SECURITY_CONTACTS` file.

The `SECURITY_CONTACTS` file informs contributors about:
  - the project's security policy
  - basic rules of security vulnerability disclosure

Optionally the `SECURITY_CONTACTS` may contain:
  - repository specific security fix back porting information
  - repository specific security fix release information

## Incubation

We think the Metal3-io project meets the criteria of incubation. We have shown healthy
number of contributions and sustainable process of feature additions in Metal3-io. Our
adopters list is also growing steadily. We will be opening proposal for incubation
very soon.


## Metal3-io Links

- [User-Guide](https://metal3io.netlify.app/introduction.html)
- [Quick start guide](https://metal3.io/try-it.html)

### Socials

- [Twitter](https://twitter.com/metal3_io)
- [Youtube](https://www.youtube.com/channel/UC_xneeYbo-Dl4g-U78xW15g)
- [Website](https://metal3.io/)
- [GitHub](https://github.com/metal3-io)
- [Slack](https://kubernetes.slack.com/messages/CHD49TLE7)
- [Mailing list/Google group](https://groups.google.com/g/metal3-dev)

### Talks

- [Kubernetes meetup, Helsinki, May 2022](https://youtu.be/LTzIudvLs9A?t=8332)
- [Open Infrastructure Foundation talk on Metal3, Dec 2021](https://youtu.be/rjSC6cJ9YY8)
- [CNCF Webinar on Metal3, Dec 2020](https://youtu.be/-BvQWz2cfCg)
- [Openshift Commons talk on Metal3, Oct 2020](https://youtu.be/HVKwWAE1nUE)

