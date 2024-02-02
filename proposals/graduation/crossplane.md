# Crossplane Graduation Proposal

Crossplane extends the Kubernetes API to enable platform teams to manage a wide
variety of infrastructure resources from multiple vendors. These resources can
be composed into higher level self-service APIs for application teams to
consume. Crossplane effectively enables platform teams to quickly put together
their own opinionated platform declaratively without having to write any code,
and offer it to their application teams as a self-service Kubernetes-style
declarative API.

Crossplane was initially accepted into the CNCF Sandbox in June 2020 then
progressed to Incubation in September 2021. Since the promotion to Incubation,
the Crossplane project has grown considerably and made significant strides along
multiple dimensions of maturity. In light of the achievements outlined in this
proposal, we recommend that the project is now ready for full Graduation.

A complete [due diligence
document](https://docs.google.com/document/d/1OBGR7h7mqQ1zE1yheIH8oPkTHiVaEH_DdT6QCTdF370/edit?usp=sharing)
was prepared by the project team when applying for Incubation and reviewed by
TAG App Delivery to provide their feedback and recommendations.  This document
has now been updated in preparation for Graduation to include notable project
progress and accomplishments since Incubation and how the specific concerns
raised by the TAG have been addressed.

Project health metrics tracked by the CNCF consistently demonstrate that the
community has continued to thrive with both adoption of the technology as well
as a strong base of contributors to the project.  We [grew our contributor base
20%](https://www.cncf.io/blog/2024/07/11/as-we-reach-mid-year-2024-a-look-at-cncf-linux-foundation-and-top-30-open-source-project-velocity/)
in the last year and are currently in the top 10% of all CNCF projects for
contributor authors, at position [12 out of
190](https://docs.google.com/spreadsheets/d/1FyVjdO9kMnz4hmEkyYLpl101cBDg_ZGHMkR5j_ce5JU/edit?gid=0#gid=0).
This number of PR authors has grown 4x from 184 at the time of Incubation to over
[799](https://crossplane.devstats.cncf.io/d/22/prs-authors-table?orgId=1)
currently and the diversity of companies contributing also quadrupled from 105 to
[395](https://crossplane.devstats.cncf.io/d/5/companies-table?orgId=1). The
overall number of contributors to the project is now almost
[2,400](https://crossplane.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=All).

Leadership at the highest level of the Crossplane project now includes
individuals from Upbound (the founding company) as well as Apple and Nokia, two
major end user companies that have adopted Crossplane as a critical component
within their environments.

Adoption of Crossplane has clearly increased since Incubation as adopters have
started publicly sharing their usage of Crossplane in the main repo’s
[ADOPTERS.md](https://github.com/crossplane/crossplane/blob/master/ADOPTERS.md)
file. There are now more than 60 companies who have publicly endorsed the
project, which is a fraction of the overall adoption base that we interact daily
with throughout the Crossplane community on Github (8.3K stars) and Slack (10.7K
members). Some notable public adopters include Nike, Autodesk, Grafana, NASA
Science Cloud, Elastic, Akamai, SAP, IBM, VMWare Tanzu, and Nokia. There are
additional adopters that are also willing to share their usage with the TOC
privately.

The ecosystem of extensions to Crossplane has grown massively since Incubation.
Many adopters of the project have written their own providers and extensions to
Crossplane and often share those with the community. Additionally, software and
technology vendors have taken their own initiative to create Crossplane
providers for their services to facilitate their adoption by the Crossplane
ecosystem.  Currently, there are over 70 providers that adopters have chosen to
donate to the community in the [crossplane-contrib
organization](https://github.com/crossplane-contrib/?q=provider-&type=all&language=&sort=),
each with their own autonomous team of maintainers, which is another sign of a
thriving contributor base for the project.  Across all of the Crossplane
projects there are [148 companies with
“committers”](https://crossplane.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Since%20moving%20to%20incubating%20state&var-metric=committers)
(maintainers with merge permissions) in the period since Incubation.

The CNCF has sponsored the creation of a Crossplane certification and the
content is currently being built by contributors to the project. Once completed,
it will appear alongside the other Cloud Native Certifications on
https://www.cncf.io/training/certification/. This official certification from
the CNCF will further assist the broader cloud native ecosystem to become
proficient with Crossplane and adopt it within their organizations to
efficiently build their own cloud native control planes.

The trustworthiness of the Crossplane project has further progressed through the
completion of key maturity milestones in the time since Incubation. Two separate
security audits have been completed by ADA logics, the first focusing on [fuzz
testing](https://blog.crossplane.io/fuzzing-security-audit/) in March of 2023
and the second being a [complete general security
audit](https://blog.crossplane.io/security-audit-2023/) that was published in
July 2023. The project has also put into place a [security disclosure
policy](https://github.com/crossplane/crossplane/security/policy) and overall
has a strong security posture with the defined policies and audits being
completed. Additionally, Crossplane has finalized its OpenSSF Best Practices
[badge](https://www.bestpractices.dev/en/projects/3260), demonstrating
compliance with industry maturity standards across multiple dimensions.

## Graduation State Criteria

### * Have committers from at least two organizations.

Across the entire Crossplane project, there are [148 different
companies](https://crossplane.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Since%20moving%20to%20incubating%20state&var-metric=committers)
that have committers (individuals with write permissions) on at least one
repository, which is a great demonstration of organizational diversity.

Also, the steering committee for the Crossplane project is composed of
individuals from 3 separate organizations: Apple, Nokia, and Upbound.

### * Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

Crossplane has invested in an OpenSSF Best Practices badge (formerly known as
Core Infrastructure Initiative (CII) Best Practices badge) and completed all
sections to earn a passing badge.  The OpenSSF Best Practices badge is displayed
prominently on the main Crossplane repository
[README.md](https://github.com/crossplane/crossplane?tab=readme-ov-file), and
all the details can be explored on the project’s page at
[https://www.bestpractices.dev/en/projects/3260](https://www.bestpractices.dev/en/projects/3260).

### * Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

Crossplane completed two separate security audits within 2023, both of which
were performed by ADA Logics. The first audit focused on
[fuzzing](https://blog.crossplane.io/fuzzing-security-audit/) and was completed
in March 2023, followed by a more intense [general security
audit](https://blog.crossplane.io/security-audit-2023/) that was broader in
scope and completed in July 2023. The full report details can be found in the
security folder of the main Crossplane repo:

* [Fuzzing
  audit](https://github.com/crossplane/crossplane/blob/master/security/ADA-fuzzing-audit-22.pdf)
* [General security
  audit](https://github.com/crossplane/crossplane/blob/master/security/ADA-security-audit-23.pdf)

In the general security audit, the ADA Logics team identified a total 16 issues,
with 7 being deemed Low severity, 8 Medium, and 1 of High severity. All issues
were reported in accordance with Crossplane’s responsible disclosure [security
policy](https://github.com/crossplane/crossplane/security/policy#reporting-a-vulnerability).
CVEs were published for 2 of these 16 issues:

* [https://nvd.nist.gov/vuln/detail/CVE-2023-37900](https://nvd.nist.gov/vuln/detail/CVE-2023-37900)
* [https://nvd.nist.gov/vuln/detail/CVE-2023-38495](https://nvd.nist.gov/vuln/detail/CVE-2023-38495) 

At the time of publishing the audit report, 15 of the 16 issues had been fixed
in the codebase and patch releases were published for all currently supported
versions of Crossplane.  The final 16th issue was in alpha code that was
subsequently removed, thus resolving 100% of the issues found during the
security audit.

### * Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

The Crossplane project has had well defined governance in place since entry into
the CNCF Sandbox, which can be found in the main repo’s
[GOVERNANCE.md](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
file. All aspects of the life cycle for Crossplane leadership positions,
including the steering committee and repository maintainers (committers) are
described in detail within this governance document.  The steering committee
members, currently from Upbound, Apple, and Nokia, can be found in the project
governance also. Repository maintainers can be found in the OWNERS.md file of
each separate Crossplane repository that make up the project.

### * Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

The complete lifecycle, including requirements and process for addition and
removal of the maintainer teams and steering committee, is described in the
project
[GOVERNANCE.md](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
document.


### * Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to [FAQs](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for guidelines on identifying adopters.

Adopters of the Crossplane project that have chosen to share their adoption
story publicly can be found in the
[ADOPTERS.md](https://github.com/crossplane/crossplane/blob/master/ADOPTERS.md)
file in the main Crossplane repository. Currently, there are over 60 public
adopters of the project, and there are more that are willing to share their
story with the TOC privately. Some notable Crossplane public adopters include
Nike, Autodesk, Grafana, NASA Science Cloud, Elastic, Akamai, SAP, IBM, VMWare
Tanzu, and Nokia.

## Incubation Details
_**Project should address each area listed below**_

### * Link to Incubation Due Diligence(DD) Document

A complete [due diligence
document](https://docs.google.com/document/d/1OBGR7h7mqQ1zE1yheIH8oPkTHiVaEH_DdT6QCTdF370/edit?usp=sharing)
was prepared by the project team when applying for Incubation and reviewed by
TAG App Delivery to provide their feedback and recommendations.

### * Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

A brief update to the due diligence document has been provided for Graduation
that focuses on notable project progress and accomplishments since the
Incubation due diligence and how the project has addressed concerns provided by
TAG App Delivery at the time of Incubation.

This document can be found in [Crossplane Due Diligence for
Graduation](https://docs.google.com/document/d/1gylzTKDdAJOir28fkApKcR_f6HUxDcQWq9CMF1ZSAa8/edit?usp=sharing).
