# Metal3.io 2021 Annual Review

## Background

There are a number of great open source tools for bare metal host
provisioning, including [Ironic](https://opendev.org/openstack/ironic/). [Metal3.io](http://metal3.io/)
aims to build on these technologies to provide a Kubernetes native API
for managing bare metal hosts via a provisioning stack that is also
running on Kubernetes. We believe that Kubernetes Native Infrastructure,
or managing your infrastructure just like your applications, is a
powerful next step in the evolution of infrastructure management.

The Metal³ project is also building integration with the Kubernetes
Cluster-API project, allowing Metal³ to be used as an infrastructure
backend for Machine objects from the [Cluster API](https://github.com/kubernetes-sigs/cluster-api).

## Include a link to your project’s devstats page.

Please refer to Metal³ [DevStats page](https://metal3.devstats.cncf.io/d/8/dashboards?orgId=1&from=now%2Fy&to=now-1h).

## How many maintainers do you have, and which organisations are they from?

Currently the project has 19 unique maintainers from Red Hat and Ericsson
Software Technologies across the main repositories. Some of the people
maintain more than one GitHub repository.

Complete list of unique maintainers across different repositories can be
found [here](https://github.com/metal3-io/metal3-docs/blob/master/maintainers/ALL-OWNERS).

## What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so.

Metal³ is incorporated into a number of open source projects, including [Airship](https://www.airshipit.org/),
[OpenShift](https://github.com/openshift/), and [Das Schiff](https://github.com/telekom/das-schiff). Each of these is associated with fully supported
vendor products and/or production deployments.

We are soliciting public references to add to the [ADOPTERS file](https://github.com/metal3-io/metal3-docs/blob/master/ADOPTERS.md).

## How has the project performed against its goals since the last review?

This is our first annual review report after joining CNCF. As such, we have
included some of the major achievements that took place after entertaining CNCF
sandbox.

- Cluster-api-provider-metal3 GitHub project
    - node reuse feature ([PR #169](https://github.com/metal3-io/cluster-api-provider-metal3/pull/169))
    - remediation controller ([PR #157](https://github.com/metal3-io/cluster-api-provider-metal3/pull/157))

- Baremetal operator GitHub project
    - authenticating to Ironic ([#601](https://github.com/metal3-io/baremetal-operator/pull/601))
    - added RAID configuration support ([#134](https://github.com/metal3-io/metal3-docs/pull/134))
    - added select vendor-independent BIOS configuration options ([#63](https://github.com/metal3-io/metal3-docs/pull/63))
    - support automatic secure boot ([#161](https://github.com/metal3-io/metal3-docs/pull/161))
    - add boot-iso API to BareMetalHost ([#150](https://github.com/metal3-io/metal3-docs/pull/150))
    - optimised concurrent provisioning performance and tested at scale up to 1000 Hosts ([PR #725,](https://github.com/metal3-io/baremetal-operator/pull/725) [Issue #905](https://github.com/metal3-io/baremetal-operator/issues/905))
    - added custom agent image controller ([#183](https://github.com/metal3-io/metal3-docs/pull/183))

## What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

Here is the list of the major changes that we are anticipating to happen until next year.

**New features**
- Split the current BareMetalHost CRD inspection data into a new CRD
- Add support for vendor-specific firmware configuration
	
**General improvements**
- Freeze the BareMetalHost CRD
- Start releasing process for BMO
- Establish a formal security reporting process
- CAPM3 v1beta1 migration
- Improve end-to-end testing

**Documentation improvements**
- Improve end user documentation

## How can the CNCF help you achieve your upcoming goals?

1. It would be great if CNCF provided some of the infrastructure for reporting
security issues - not necessarily becoming a CNA, but e.g. hosting security
reporting mailing lists at an official address.

2. We have several container images that we are pulling constantly in our CI system
from Docker Hub for every single integration test that we run on almost every pull
request. Due to the [rate limitation of API calls to the Docker Hub](https://www.docker.com/increase-rate-limits),
we had to find a workaround (not the best) to overcome the rate limit. It would be very helpful if
CNCF could help us to get a paid subscription with Docker Hub which would have less
limits than an unauthenticated/personal user account (200 API calls to the Docker
Hub within 6 hours). 

## Do you think that your project meets the criteria for incubation?

We are aware that we are lacking a documented mechanism for reporting security issues. 
The BareMetalHost CRD is still at version v1alpha1 and additional fields have been
added to it over time. We are working on freezing the API and adopting a versioning
scheme for the baremetal-operator. The cluster-api-provider-metal3 already has stable
CRDs and a versioning scheme. Other than that we believe the project meets the criteria
for incubation.
