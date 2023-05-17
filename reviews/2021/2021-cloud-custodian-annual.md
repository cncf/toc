# Cloud Custodian Sandbox Annual Report 2021

# Background

Cloud Custodian is a YAML DSL based stateless rules engine for cloud audit, management, and governance.

Cloud Custodian exists as a policy/governance as code tool for organizations to enable realtime detective controls to ensure that organizational policies are being adhered to by their cloud infrastructure. 
Policy/governance as code is about bringing some of the best practices of software engineering to policy/governance, ie. using version control, code review, testing, and continuous integration and deployment tools.

Custodian can be used to manage AWS, Azure, and GCP environments by ensuring real time compliance to security policies (like encryption and access requirements), tag policies, operations best practices, and cost management via garbage collection of unused resources and off-hours resource management. 
Kubernetes and OpenStack support is present in alpha stages at the moment.

Some of its features:

* Ability to implement detective controls to ensure adherence to organizational policies
* Supports real-time detection, reporting/notification, and remediation
* Consistent outputs and telemetry (blob, logs, trace, metrics) with provider native sinks.
* Minimal installation, usable as a cli query/investigative tool, or as an operations tool in a compliance as code environment.


### Intended Use

Custodian enables reusable vocabularies of filters and actions that can be used to implement many semantic policies across hundreds of resource types. 
For example an action like `stop` a vm instance could be used for offhours savings, tag compliance, or incident response. 
Some additional examples use cases might be

* Identify and Remediate load balancers or storage buckets not
  configured for logging
* Turning off development servers/clusters and databases off at night
  to realize cost savings.
* Finding underutilized resources and sending an email to their
  creator to reduce the size.
* Enforcing tag compliance policies on resources.
* Finding resources with embedded access control policies that are
  setup to give access across the org boundary.
* In response to a security event, used as a remediation tool to
  snapshot, disk, network isolate a server, change cloud credentials,
  and install forensics tools.

## Preferred maturity level

Applying for Incubation Soon(tm)

## Sponsor / Advisor from Toc

Ricardo Rocha

## License

[Apache 2](https://github.com/cloud-custodian/cloud-custodian/blob/master/LICENSE.txt)

# Sandbox Proposal Artifacts

[Slides] -- ? 
[Proposal Document](https://github.com/cncf/toc/blob/main/proposals/sandbox/cloudcustodian.md)
[Pull Request](https://github.com/cncf/toc/pull/480)

# Sandbox Acceptance 

c7n was accepted into the CNCF Sandbox on June 25, 2020

# Highlights since Sandbox

Governance as Code Day as Part of KubeCon + CloudNativeCon
    - 2,500+ registrants, 310 live attendees, top 50 hashtags for KubeCon
    - [Recordings from the day](https://www.youtube.com/watch?v=4J7ZaO8oDgg&list=PLtIlR7WdaxTEj45N63lUgrd2IhS_gD3pe&index=1)
    
- [State of Cloud Custodian 2021](https://youtu.be/DB52-yaIrqY)
  - A summary video for our community, covers many aspects of the project

## Goals for 2022

- New enhancement process for feature development
- GCP moving from beta to GA
- ARM64 docker images
- Extracting core from the AWS provider (have each cloud provider in its own package)
- Kubernetes Admission Controller
- Policy Enforcement against IaC definitions like Terraform (aka shift left)
- End User Policy Testing framework

# Incubation Requirements

## Dev Statistics
"Have a healthy number of committers"
"Demonstrate a substantial ongoing flow of commits and merged contributions"

- [c7n devstats](https://cloudcustodian.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&viewPanel=4)

| Statistic | Sandbox | Current | Growth | 
|-|-|-|-|
| GitHub Stars | 2,940+ | 3,888 | 32% |
| Releases     | 77+    | 116+  | 50% | 
| Commits      | 3,292+ | 3,677 | 12% |
| Forks        | 856+   | 1,100+| 29% |
| Contributors | 258+   | 320+  | 24% |
| Maintainers  | 14     | 10    | -29% |
| Affiliations between Maintainers | 4 | 6 | 50% |

## Maintainers

- [Current Maintainers](https://github.com/cloud-custodian/cloud-custodian/blob/master/OWNERS.md)

#### Leads

- Kapil Thangavelu @kapilt (@kapilt) - Project Lead - Stacklet

#### Aws Provider Maintainers

- Todd Stansell (@tjstansell) - 23andme
- John Hillegass (@JohnHillegass) - Fidelity
- Sonny Shi (@thisisshi) - Stacklet
- David Filiatrault (@FireballDWF) - Amazon
- Jamison Roberts (@jtroberts83) - Stacklet

#### Gcp Provider Maintainers

- Karol Lassak (@ingwarsw) - Ingwar & co.

#### Azure Provider Maintainers

- Stefan Gordon (@stefangordon) - Stacklet
- Andy Luong (@aluong) - Microsoft
- Kiril Logachev (@logachev) - Stacklet

#### Emeritus

Since the opensource inception of Cloud Custodian in April 2016, we have had numerous maintainers who over the course of years are no longer current contributors, but we still celebrate their contributions and they are listed here.

- Kit Ewbank (@ewbankkit)
- Mandeep Bal (@mandeepbal)
- Erin Welch (@erwelch)
- Darcy Laylock (@Sutto)

# Security
"Clearly documented security processes explaining how to report security issues to the project"

Dedicated [security mailing list](https://github.com/cloud-custodian/cloud-custodian#security) for potential vulnerability reports. 

More security details in: 
[Joint Review](https://github.com/cncf/tag-security/blob/cf1f38ff803fb3f830209e16618221844a9dea68/assessments/projects/custodian/joint-review.md)
[Self-assessment](https://github.com/cncf/tag-security/blob/ef2d51d54697f4037a4a697b405d270df2579620/assessments/projects/custodian/self-assessment.md)

Note: Both documents pending merge, will update accordingly

# Adoption

[ADOPTERS.md](https://github.com/cloud-custodian/cloud-custodian/blob/master/ADOPTERS.md)

- Downloads from end users have increased 4x since this time last year
- Over a million downloads a month
- Over 60 million total downloads from Docker Hub

## Source control

The Cloud Custodian project is composed of multiple source code repositories, all of which are hosted on Github under the `cloud-custodian` organization:

- [https://github.com/cloud-custodian](https://github.com/cloud-custodian)

The core repository can be found at

- [https://github.com/cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian)

## Issue tracker

All issues are found in the appropriate repo in our GitHub organization:
[https://github.com/cloud-custodian](https://github.com/cloud-custodian).

## Website

The official public facing website can be found at [https://cloudcustodian.io/](https://cloudcustodian.io/)

## Communication channels (slack, irc, mailing lists)

The official communication channel for the project can be found at: 

- https://gitter.im/cloud-custodian/cloud-custodian

Less active channels include a [mailing list](https://groups.google.com/g/cloud-custodian) and [Stack Overflow](https://stackoverflow.com/questions/tagged/cloudcustodian).

### Appendix

* CII Best Practices. https://bestpractices.coreinfrastructure.org/en/projects/3402

* Case Studies / Articles

  - https://stelligent.com/2017/05/15/cloud-custodian-cleans-up-your-cloud-clutter/
  - https://www.cloudsecops.com/aws-security-audit-using-cloud-custodian-for-aws/
  - https://aws.amazon.com/blogs/opensource/announcing-cloud-custodian-integration-aws-security-hub/
  - https://aws.amazon.com/blogs/opensource/compliance-as-code-and-auto-remediation-with-cloud-custodian/
  - https://medium.com/slalom-technology/managing-cloud-compliance-at-scale-with-aws-security-hub-and-cloud-custodian-73c630863a59

* We have a separate document with more exhaustive threat modeling
  exercises contributed by Capital One
  https://docs.google.com/document/d/1S9zQZaT6G1TA3IAx6YNL0f7G938xaFZ-bziszhuxMZg/edit?usp=sharing

## Project logo in svg format

All graphics and artwork can be found in the primary Cloud Custodian repo:
[https://github.com/cloud-custodian/cloud-custodian/tree/master/docs/logos](https://github.com/cloud-custodian/cloud-custodian/tree/master/docs/logos)
