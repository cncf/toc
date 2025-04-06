# Confidential Containers 2022 Annual Review

## Background

Confidential Containers is a CNCF Sandbox project that enables cloud native
confidential computing by taking advantage of a variety of hardware platforms
and technologies. The project was accepted as a CNCF Sandbox in March 2022.

Confidential Containers builds on existing and emerging hardware security
technologies such as Intel SGX, Intel TDX, AMD SEV, and IBM Z Secure Execution,
in combination with new software frameworks to help better secure data-in-use 
for users of cloud native workloads.

The main goal of the project is to standardize confidential computing at the
pod level and simplify its consumption in Kubernetes. This is in order to enable
Kubernetes users to deploy confidential container workloads using familiar
workflows and tools, without extensive knowledge of underlying confidential
computing technologies.

## DevStats

The Confidential Containers project is a collection of logically and
programmatically bound sub-projects, hosted across multiple git repositories
under the Confidential Containers GitHub organization.

Below is a chart of all GitHub activities across all Confidential Containers
repositories, since it's acceptance as a CNCF project:
* [Confidential Containers GithHub Contributions per repository (2022-03-08 to Now)](https://confidentialcontainers.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&from=1646694000000&to=1686261599000)

The following chart is an aggregating view of all contributions to the
Confidential Containers project as whole, for the same time period:
* [Confidential Containers GithHub Contributions (2022-03-08 to Now)](https://confidentialcontainers.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=1646694000000&to=now-1h)

The project hourly activity for the same period is shown in the below chart:
* [Confidential Containers GitHub Hourly Activity (2022-03-08 to Now)](https://confidentialcontainers.devstats.cncf.io/d/8/dashboards?orgId=1&from=1646694000000&to=now-1h&viewPanel=2&refresh=15m)

Confidential Containers new Contributors:
* [Confidential Containers New Contributors (2022-03-08 to Now)](https://confidentialcontainers.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1646694000000&to=now-1h)

These charts and statistics show a few interesting points:

* Since its acceptance as a CNCF project, the project has seen a steady stream
  of contributions.
* A very important metric for Confidential Containers: More than 40 new
  contributors have joined the project over the past year.
* Contributions are spread over different repositories, showing the width of the
  project:
  * On the client side, the `enclave-cc`, `attestation-agent`, `td-shim` and
    `image-rs` projects take the bulk of the contributions.
  * On the attestation part of the project, most of the contributions went to
    the `kbs` and `attestation-service` repositories.
  * The `cloud-api-adaptor` and `operator` sub-projects merged most of the
    infrastructure related contributions to Confidential Containers.

## Maintainers

The Confidential Containers project maintainership structure is defined as
follows:

* Each sub-project has its own list of maintainers, tracked through both a
  per-repo MAINTAINERS file and a corresponding GitHub team. We have multiple
  instances of individuals maintaining several sub-projects.
* A Technical Steering Committee (TSC) provides decision-making and strategic
  oversight for the project.

The [TSC](https://github.com/confidential-containers/community/blob/main/MAINTAINERS)
is composed of individuals from multiple companies (Alibaba, AMD, IBM, Intel,
Red Hat and Rivos):

- [@ariel-adam](https://github.com/ariel-adam), Ariel Adam, Red Hat
- [@bpradipt](https://github.com/bpradipt), Pradipta Banerjee, Red Hat
- [@dcmiddle](https://github.com/dcmiddle), Dan Middleton, Intel
- [@fitzthum](https://github.com/fitzthum), Tobin Feldman-Fitzthum, IBM
- [@jiazhang0](https://github.com/jiazhang0), Zhang Jia, Alibaba
- [@jiangliu](https://github.com/jiangliu), Jiang Liu, Alibaba
- [@larrydewey](https://github.com/larrydewey), Larry Dewey, AMD
- [@magowan](https://github.com/magowan), James Magowan, IBM
- [@peterzcst](https://github.com/peterzcst), Peter Zhu, Intel
- [@sameo](https://github.com/sameo), Samuel Ortiz, Rivos

## Adoption

Cloud Providers (Microsoft, Alibaba, IBM) and Open Source Vendors (e.g. Red Hat)
are pursuing offerings which will include Confidential Containers. They are
actively engaged in the community.

For example, Red Hat now supports [Confidential Containers with Openshift on Azure](https://www.redhat.com/en/blog/confidential-containers-azure-openshift-setup-guide),
and Azure supports [Confidential Containers on AKS](https://superuser.openinfra.dev/articles/aligning-with-kata-confidential-containers-to-achieve-zero-trust-operator-deployments-with-aks/)
as well.

## Project Goals

### Performance

This is the first annual Review for Confidential Containers, so the comparison
baseline is the the moment when Confidential Containers was accepted as a CNCF
Sandbox project, i.e. March 2022.

Since then, the Confidential Containers project published 6
[releases](https://github.com/confidential-containers/documentation/tree/main/releases)
by following a time-based, 6 weeks period release cadence.

Starting from the initial release, Confidential Containers delivered on its
initial and major goal of enabling seamless use and deployment of unmodified
containers through Kubernetes, across multiple TEEs and hardware platforms.

The project first 6 releases brought the following major key features and
enhancements:

- A stabilized, enhanced, and wider CI that is key to the project's fast release
  cadence.
- A hardware backed CI, i.e. a CI that runs tests on actual confidential
  computing hardware instances like e.g. Intel TDX, AMD SEV, and Intel SGX.
- A Kubernetes Operator for easily deploying the Confidential Containers
  runtime, stack, and payload on Kubernetes clusters.
- Support for process-based TEEs like e.g. SGX.
- Technical debt reduction, in particular the dependencies on `umoci` and
  `skopeo` are gone.
- A modular, [RATS-compliant](https://datatracker.ietf.org/doc/rfc9334/) Key
  Broker Service (KBS) that implements a full remote attestation service. The
  Confidential Container KBS provides its own attestation service but also
  supports external ones like Intel Amber.

### Current Goals

At a high level, the current project focus is on stability and reliability
increase, performance improvements, and customer adoption expansion.

This translates into the following goals:

1. CI/CD improvements
1. Host-Based container images pulling
1. Extended attestation model
1. Policy and secure container configuration
1. Kata Containers upstreaming
1. Make the project currently features-driven implementation more focused on
   customer use cases.

## How Can CNCF Help the Project?

## Incubation Criteria

We are in the process of getting direct adopters to consume the Confidential
Containers solution; both members of our community and external. Our target is
to have at least 3 direct adopters by the end of 2023, helping us to grow to
the next level of incubation.

Aside from customer adoption, we also need to improve our current vulnerability
management process. Based on the [CNCF Security Guidelines](https://contribute.cncf.io/maintainers/security/security-guidelines/),
we have to define a more formal way of tracking security issues, from the
initial reporting steps up to the project release management for rapidly
addressing them.
