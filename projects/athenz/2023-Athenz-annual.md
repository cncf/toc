# Athenz-2023-Annual Review

This is the annual review for the [Athenz](https://athenz.io) project for 2023.

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
shows a consistent rate of commits throughout the year. While the significant percentage of code commits is
from the core maintainers who work at Yahoo, we're encouraged to see an increased number of pull requests
from Yahoo Japan developers in the recent months. Athenz GitHub repo has 787 stars and 265 forks.

The maintainers prefer a periodic release of Athenz servers and libraries as opposed to few releases
with a large number of changes. At the last review time, the team had just released version 1.10.49. Since
then, we have released 38 versions, and are now at [1.11.33]
(https://github.com/AthenZ/athenz/blob/master/CHANGELOG).

## Maintainers

Athenz currently has [3 core maintainers](https://github.com/AthenZ/athenz/blob/master/MAINTAINERS), all from Yahoo:

| Maintainer       | GitHub Username                               | Company     |
| ---------------- | --------------------------------------------- | ----------- |
| Henry Avetisyan  | [@havetisyan](https://github.com/havetisyan)  | Yahoo       |
| Abhijeet Vaidya  | [@abvaidya](https://github.com/abvaidya)      | Yahoo       |
| Dvir Guttman     | [@dvirguttman](https://github.com/dvirguttman)| Yahoo       |

One of the previous maintainers, Ofer Levi left Yahoo at the beginning of the year and did not wish
to continue to act as a maintainer and was replaced with Dvir Guttman.

The Athenz K8S Service Identity Agent subproject has the
[following 4 maintainers](https://github.com/AthenZ/k8s-athenz-sia/blob/master/MAINTAINERS),
all from Yahoo! Japan Corporation.

| Maintainer         | GitHub Username                            | Company      |
| ------------------ | ------------------------------------------ |--------------|
| Fan Wingkin        | [@WindzCUHK](https://github.com/WindzCUHK) | Yahoo! Japan |
| Tatsuya Yano       | [@tatyano](https://github.com/tatyano)     | Yahoo! Japan |
| Seitaro Suno       | [@ssunorz](https://github.com/ssunorz)     | Yahoo! Japan |
| Kyo Fujisaki       | [@ mlajkim](https://github.com/mlajkim)    | Yahoo! Japan |

The Athenz HashiCorp Terraform Plugin subproject has the following [core maintainer](https://github.com/AthenZ//terraform-provider-athenz/blob/master/MAINTAINERS) from Yahoo.

| Maintainer       | GitHub Username                                | Company     |
| ---------------- | ---------------------------------------------- | ----------- |
| Dvir Guttman     | [@dvirguttman](https://github.com/dvirguttman) | Yahoo       |


Athenz Governance Policy is available in our [Github repository](https://github.com/AthenZ/athenz/blob/master/GOVERNANCE.md).

## Adoption

The current list of Athenz users are documented in the [Adopters file](https://github.com/AthenZ/athenz/blob/master/ADOPTERS.md)
in our Github Repository.

At this time Athenz is being used in a production environment at two companies:
Yahoo and Yahoo Japan. Athenz is heavily supported and used at both companies. At Yahoo, it
is used to provide an x.509 identity to all of our workloads deployed within our Kubernetes
clusters, on-prem data centers, AWS and GCP public clouds enabling those components to
enforce mutual TLS authentication for all service to service communications. Athenz Authorization
Token Service enables Yahoo service to utilize TLS bound OAuth2 tokens to implement role based
authorization support. It is used similarly at Yahoo Japan. 

## Project Goals

Since our last annual review, the Athenz team has added a significant number of new features. Some notable one are:

* Support for Google Cloud Platform. Athenz Service Identity Agents can now issue x.509 service
  Identity certificates to workloads running in GCP Compute Engine, Kubernetes Service, and Functions.
* Support Athenz as an OIDC Provider for Console Single-Sign-On in GCP with Workforce Identity Federation.
* Publish authentication history and zms domain syncer utilities
* Update UI to use node.Js 18.x and introduce redux support
* The java zms and zts clients were rewritten using Apache Http Client instead of through jersey connectors.
* Updated to use latest Jetty 1.11.x release with Jakarta namespace and jersey 3.x instead of javax namespace and jersey 2.x
* Requires the use of JDK 11.x to compile and build.
  
The full list of features and changes are available in our [roadmap](https://github.com/AthenZ/athenz/blob/master/roadmap.md)
and in our [release change logs](https://github.com/AthenZ/athenz/releases/tag/v1.11.33)

We have started two major features this quarter:
* Integration with [Cert-Manager](https://cert-manager.io) which is a CNCF project
* Implement a K8S operator for Athenz to manage Athenz resources using CRDs

## CNCF membership

Athenz has mainly benefited in the area of awareness from inclusion in the CNCF as a sandbox
project. This is an area that we’d like to get continued support. Athenz team held a Virtual Project
Office hour during Kubecon + CloudNativeCon + North America 2022.

## Incubation

At this time Athenz is not ready to apply for incubation due to limited adoption. We believe it
is already a mature product handling significant production load both at Yahoo and Yahoo Japan,
but it is our goal to improve our adoption and get more consistent contributions from additional
organizations before applying for incubation.
