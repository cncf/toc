# Athenz-2022-Annual Review

This is the annual review for the [Athenz](https://athenz.io) project for 2022.

## Table of Contents

- [Background](#background)
- [DevStats](#devstats)
- [Maintainers](#maintainers)
- [Adoption](#adoption)
- [Project goals](#project-goals)
- [CNCF Membership](#cncf-membership)
- [Incubation](#incubation)

## Background

Athenz is an open source platform for X.509 certificate based service authentication and fine-grained
access control in dynamic infrastructures. It enables zero trust core principles like traffic encryption,
authentication, authorization, dynamic trust and least privilege access. Athenz was accepted as a
CNCF Sandbox project on January 26th, 2021. It provides the following two major features:

* X.509 Certificate Based Authentication
  - Service identity in the form of short-lived X.509 certificates to all workloads deployed in
    private or public clouds.
* Fine-grained Authorization
  - Authorize all authenticated clients using fine-grained role-based (RBAC) access control with
    industry standard OAuth2 JWT access tokens.

## DevStats

Athenz DevStats dashboard [here](https://athenz.devstats.cncf.io/d/8/dashboards?orgId=1&viewPanel=2&from=now-1y&to=now-1h&refresh=15m)
shows a consistent rate of commits throughout the year. The significant percentage of code commits is
from the core maintainers, most of whom work at Yahoo as most of the products at Yahoo such as Mail,
Sports, Finance and Ads are heavily integrated with Athenz and the team receives a large number of
feature requests and enhancements from those product teams.

The maintainers prefer a periodic release of Athenz servers and libraries as opposed to few releases
with large number of changes. When Athenz joined CNCF as a sandbox project, we had just released the
1.10.4 version of the product. In the last year we’ve released 45 minor versions and the latest Athenz
release now is [1.10.49](https://github.com/AthenZ/athenz/blob/master/CHANGELOG).

## Maintainers

Athenz currently has [3 core maintainers](https://github.com/AthenZ/athenz/blob/master/MAINTAINERS), all from Yahoo:

| Maintainer       | GitHub Username                               | Company     |
| ---------------- | --------------------------------------------- | ----------- |
| Henry Avetisyan  | [@havetisyan](https://github.com/havetisyan)  | Yahoo       |
| Abhijeet Vaidya  | [@abvaidya](https://github.com/abvaidya)      | Yahoo       |
| Ofer Levi        | [@OferLevi85](https://github.com/OferLevi85)  | Yahoo       |

The Athenz Auth0 and Prometheus plugin subprojects have the [following 5 maintainers](https://github.com/AthenZ/athenz-metric-prometheus/blob/master/MAINTAINERS), all from Yahoo! Japan Corporation.

| Maintainer    | GitHub Username                            | Company      |
| ------------- | ------------------------------------------ |--------------|
| Windz Fan     | [@WindzCUHK](https://github.com/WindzCUHK) | Yahoo! Japan |
| Tatsuya Yano  | [@tatyano](https://github.com/tatyano)     | Yahoo! Japan |
| Seitaro Suno  | [@ssunorz](https://github.com/ssunorz)     | Yahoo! Japan |
| Takayuki Ino  | [@falz-tino](https://github.com/falz-tino) | Yahoo! Japan |
| Kyo Fujisaki  | [@kyfujisa](https://github.com/kyfujisa)   | Yahoo! Japan |

Yahoo! Japan is also planning to contribute its Athenz Go-based authorization sidecar to the Athenz project as well.

The Athenz HashiCorp Terraform Plugin subproject has the following [core maintainer](https://github.com/AthenZ//terraform-provider-athenz/blob/master/MAINTAINERS) from Yahoo.

| Maintainer       | GitHub Username                                | Company     |
| ---------------- | ---------------------------------------------- | ----------- |
| Dvir Guttman     | [@dvirguttman](https://github.com/dvirguttman) | Yahoo       |

Athenz Governance Policy is available in our [Github repository](https://github.com/AthenZ/athenz/blob/master/GOVERNANCE.md).

## Adoption

Unfortunately we were not keeping track of our adopters and we just created an [Adopters file](https://github.com/AthenZ/athenz/blob/master/ADOPTERS.md)
in our Github Repository to start to track publicly referenceable uses of the projects.

At this time we’re aware of Athenz being used in a production environment in two companies:
Yahoo and Yahoo Japan. Athenz is heavily supported and used at both companies. At Yahoo, it
is used to provide an x.509 identity to all of our workloads deployed within our Kubernetes
clusters, On-prem data centers and in AWS enabling those components to enforce mutual TLS
authentication for all service to service communications. Athenz Authorization Token Service
enables Yahoo service to utilize TLS bound OAuth2 tokens to implement role based authorization support.

## Project Goals

When Athenz joined CNCF as a sandbox project, its goals were to add additional functionality
and increase adoption. We have added a significant number of new features in the last year. Some notable one are:

* Support Athenz as an OIDC Provider in Kubernetes
* Improve operability of Athenz with Kubernetes Service Mesh by introducing Envoy SDS support for service x.509 certificates.
* An official [Athenz HashiCorp Terraform Provider](https://registry.terraform.io/providers/AthenZ/athenz/latest)
* Support publishing domain changes as events through an interface. An implementation of
  [Apache Pulsar](https://pulsar.apache.org) was delivered as part of the feature
* Supporting of Tags (key/value pairs) for domains, roles and groups.
* Support for multi-version authorization policies providing the capability to quick and easy
  rollback to a previous version if necessary.
* Deliver OAuth 2.0 Rich Authorization Requests feature based on [Internet Draft](https://www.ietf.org/id/draft-ietf-oauth-rar-10.html)

The full list of features and changes are available in our [roadmap](https://github.com/AthenZ/athenz/blob/master/roadmap.md)
and in our [release change logs](https://github.com/AthenZ/athenz/releases/tag/v1.10.49)

One major area we have been focusing on last year is micro-segmentation support based on service identities.
We’re experimenting with eBPF technology to provide such capabilities within Athenz.

To increase adoption we were hoping that we would get a contribution to provide support for
Google Cloud Platform since Athenz already provides support for providing workloads running in AWS
and Azure with service identity x.509 certificates. We recently were notified that
[Vespa](https://vespa.ai) team will likely be contributing support for the Google Cloud Platform.

One final area that we’ll be focusing on is initial setup. Since a typical setup of Athenz
requires installations and coordination of several components: Management Service, Token Service,
UI, Authentication Authority and Certificate Signer, we’ve heard from several users through
our mailing list and/or slack channel that they’re looking for a simplified installation
setup that would allow users bring up a server quickly to determine if it satisfies their
requirements thus can be considered to be deployed at their companies.

We believe once we have support for all three major cloud providers along with micro-segmentation
support, and a simplified installation setup, we’ll be in a better position to focus on
increasing adoption of Athenz.

## CNCF membership

Athenz has mainly benefited in the area of awareness from inclusion in the CNCF as a sandbox
project. This is an area that we’d like to get continued support. During one of the End-User
Community Calls we were advised to present and request a security review from the Security
TAG which we are planning to do once we finalize our micro-segmentation based on service
identities feature.

## Incubation

At this time Athenz is not ready to apply for incubation due to limited adoption. It is our
goal to improve our adoption and get more consistent contributions from additional
organizations before applying for incubation.
