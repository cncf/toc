# Cloud Custodian

## Name of project

Cloud Custodian

## Project description

### Introduction

Cloud Custodian is a YAML DSL based stateless rules engine for cloud audit, management, and governance.

Some of its features:
* Ability to implement detective controls to ensure adherence to organizational policies
* Supports real-time detection, reporting/notification, and remediation
* Consistent outputs and telemetry (blob, logs, trace, metrics) with provider native sinks.
* Minimal installation, usable as a cli query/investigative tool, or as an operations tool in a compliance as code environment.

As organizations move to cloud infrastructure the need to automate
governance becomes paramount to realize the benefits of self-service
and on demand infrastructure while ensuring compliance with
organizational policies. Different teams within an organization will
often use distinct tools for their infrastructure
provisioning. However, ensuring that the organization as a whole is
well-managed requires additional tooling to ensure that governance
objectives are met regardless of the choice of provisioning tool.

Cloud Custodian exists as a policy/governance as code tool for
organizations to enable detective controls to ensure that
organizational policies are being adhered to. Policy/governance as
code is about bringing some of the best practices of software
engineering to policy/governance, ie. using version control, code
review, testing, and continuous integration and deployment tools.

Custodian can be used to manage AWS, Azure, and GCP environments by
ensuring real time compliance to security policies (like encryption
and access requirements), tag policies, and cost management via
garbage collection of unused resources and off-hours resource
management. Kubernetes and OpenStack support is present in alpha
stages at the moment.


### Intended Use

Custodian enables reusable vocabularies of filters and actions that
can be used to implement many semantic policies. For example an action
like `stop` a vm instance could be used for offhours savings, tag
compliance, or incident response. Some additional examples might be

* Identify and Remediate load balancers or storage buckets not configured for logging
* Turning off development servers/clusters and databases off at night to realize cost savings.
* Finding underutilized resources and sending an email to their creator to reduce the size. 
* Enforcing tag compliance policies on resources.
* Finding resources with embedded access control policies that are setup to give access across the org boundary.
* In response to a security event, used as a remediation tool to
  snapshot, disk, network isolate a server, change cloud credentials,
  and install forensics tools.

## Preferred maturity level

Incubation

## Sponsor / Advisor from Toc

Ricardo Rocha

## License

Apache 2:
https://github.com/cloud-custodian/cloud-custodian/blob/master/LICENSE.txt

## Source control

The Cloud Custodian project is composed of multiple source code
repositories, all of which are hosted on Github under the
`cloud-custodian` organization:
[https://github.com/cloud-custodian](https://github.com/cloud-custodian)

The core repository can be found at
[https://github.com/cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian)

## Issue tracker

All issues are found in the appropriate repo in our GitHub organization:
[https://github.com/cloud-custodian](https://github.com/cloud-custodian).

## Website

The official public facing website can be found at [https://cloudcustodian.io/](https://cloudcustodian.io/)

## Community size and any existing sponsorship


The following statistics, when related to code repositories, are for the main Cloud Custodian repo of
[https://github.com/cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian).

* GitHub Stars: 3,569
* Releases: 86+
* Commits: 3554+
* Forks: 996+
* Contributors: 299+
* Maintainers: 20 maintainers from 5 organizations


### Who is currently known to be using the project? Are they using it in production and at what scale?

# Adopters

Cloud Custodian’s community includes public contributors from a large number of
production users across a variety of industries, companies, and global
users operating Cloud Custodian at enterprise scale.

End User contributors:

 - Capital One
 - Comcast
 - Cox Automotive
 - 23 and Me
 - Transamerica
 - TCF
 - Georgia State University
 - Toyota Connected
 - Stacklet
 - CapitalOne
 - State Farm
 - Infosys
 - Disney
 - Hulu
 - Bloomberg
 - Slalom

Cloud Provider Contributors:

 - Amazon
 - Microsoft

Sourced from cncf devstats.

## Communication channels (slack, irc, mailing lists)

The official communication channel for the project can be found at https://gitter.im/cloud-custodian/cloud-custodian

Less active channels include a mailing list, reddit, and stack overflow.

### Appendix

* CII Best Practices. https://bestpractices.coreinfrastructure.org/en/projects/3402

* Case Studies / Articles

** https://stelligent.com/2017/05/15/cloud-custodian-cleans-up-your-cloud-clutter/
** https://www.cloudsecops.com/aws-security-audit-using-cloud-custodian-for-aws/
** https://aws.amazon.com/blogs/opensource/announcing-cloud-custodian-integration-aws-security-hub/
** https://aws.amazon.com/blogs/opensource/compliance-as-code-and-auto-remediation-with-cloud-custodian/
** https://medium.com/slalom-technology/managing-cloud-compliance-at-scale-with-aws-security-hub-and-cloud-custodian-73c630863a59

* We have a separate document with more exhaustive threat modeling
  exercises contributed by Capital One
  https://docs.google.com/document/d/1S9zQZaT6G1TA3IAx6YNL0f7G938xaFZ-bziszhuxMZg/edit?usp=sharing

## Project logo in svg format

All graphics and artwork can be found in the primary Cloud Custodian repo:
[https://github.com/cloud-custodian/cloud-custodian/tree/master/docs/logos](https://github.com/cloud-custodian/cloud-custodian/tree/master/docs/logos)
