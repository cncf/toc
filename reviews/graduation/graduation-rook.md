# Rook Graduation Proposal

Rook is currently an incubation stage project.  Please refer to Rook's [incubation stage
proposal](incubation-rook.md) and [sandbox stage proposal](../proposals/rook.adoc) ("inception" at
time of acceptance) for details on the acceptance criteria at the time of those previous proposals.

In the time since being accepted to the incubation stage, Rook has demonstrated healthy growth of
the community, continually released new features, and expanded storage provider support.  This
proposal focuses on Rook's maturity and compliance with the [CNCF Graduation Criteria
v1.3](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md).

Every major project and community statistic has [**increased
2-3x**](https://docs.google.com/presentation/d/1mMPYMDC4JMGWhoL3FzFgeasSLJepNwYMfwQD-T_gET4/edit#slide=id.g3e2163a66f_0_0)
(contributors, commits, organizations, GitHub stars, etc.).  Some have even **increased 10x**, such
as container downloads going from 13.3M to 124.9M.

**Four major releases** were completed (0.9, 1.0, 1.1, and 1.2), as well as **28 patch releases**,
for a total of **32 releases overall**. Across these releases, operators for **4 new storage
providers** were integrated, with a total of **6 storage providers** currently supported:

* EdgeFS, Cassandra, and NFS orchestration released in v0.9
* YugabyteDB orchestration added in v1.1

Support for both Ceph and EdgeFS has been declared **Stable** in this time frame, reflecting the
increased maturity that has only been possible from impressive engagement from the community.  The
community has provided strong and consistent feedback that they have been able to rely on the
stability of these operators, even through upgrading across all the releases of new versions.

The [formalized governance policy](https://github.com/rook/rook/blob/master/GOVERNANCE.md) has been
updated to institute a project steering committee and add multiple new maintainers, doubling the
number of [maintainers](https://github.com/rook/rook/blob/master/OWNERS.md) since incubation,
increasing the diversity of maintainers, and helping the project continue to grow.

Many other project maturity milestones have been achieved, such as the completion of a [full
security audit](https://github.com/rook/rook/blob/master/SECURITY.md#security-audit), a [security
disclosure process](https://github.com/rook/rook/blob/master/SECURITY.md) has been defined, and a
[100% passing level criteria](https://bestpractices.coreinfrastructure.org/en/projects/1599) has
been achieved for the Core Infrastructure Initiative Best Practices Badge.

## Graduation State Criteria

### * Have committers from at least two organizations.

Rook currently has [235 contributors](https://github.com/rook/rook/graphs/contributors) from [107
organizations](https://rook.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=commits)

### * Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge/(https://bestpractices.coreinfrastructure.org/).

We have achieved 100% passing level criteria for our [CII best practices
badge](https://bestpractices.coreinfrastructure.org/en/projects/1599).  The badge is displayed in
the header of the main repository's
[`README.md`](https://github.com/rook/rook/blob/master/README.md)

### * Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

A third party security audit was performed in December 2019 by [Trail of
Bits](https://www.trailofbits.com/). The full security report has been published and is [available
for download](https://drive.google.com/file/d/1rOwrwYmBUpLUm6W5J5rhXvdVit818hWJ/view?usp=sharing).

All critical vulnerabilities have been fixed and released.

### * Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

The project governance and committer process was defined as part of the Incubation criteria, but has
been recently updated to incorporate a steering committee for the project in addition to adding more
maintainers from a greater diversity of organizations.  Full details can be found in the formally
described [`GOVERNANCE.md`](https://github.com/rook/rook/blob/master/GOVERNANCE.md) file.

### * Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.

A partial list of production adopters can be found in the
[ADOPTERS.md](https://github.com/rook/rook/blob/master/ADOPTERS.md) file in the `rook/rook` repo.

## Proposal Presentation

Further details of Rook's growth and progress since entering the incubation stage as well as
production adopter testimonials and use case details from the Rook community can be found in this
[slide
deck](https://docs.google.com/presentation/d/1mMPYMDC4JMGWhoL3FzFgeasSLJepNwYMfwQD-T_gET4/edit?usp=sharing).

## Incubation Details

### * Link to Incubation Due Diligence(DD) Document

Due diligence has been recently completed with the help of the CNCF Storage SIG and can be found in
this [due diligence
document](https://docs.google.com/document/d/1acp9gJ1D_qflHKJBg4gB-nZwZQs87_Dh9uiH4pITc_U/edit?usp=sharing).

More details from when Rook was accepted into incubation stage can be found in the [original
incubating proposal](incubation-rook.md).

### * Address any concerns or recommendations from the SIG and/or TOC sponsor(s) from the DD Document

The CNCF Storage SIG and TOC sponsor Saad Ali has recommended Rook to graduate. This recommendation
was given to the TOC in the TOC Meeting on March 3, 2020. Slide with recommendation is directly
linked in the [meeting slide
deck](https://docs.google.com/presentation/d/1ChMwj1LJuPVp64PXMG6RGe3_sX0wAerQp0ycObbFq1M/edit#slide=id.g7c144c2455_0_7)
here.

The recommendation for graduation can also be found in this [SIG recommendation
doc](https://docs.google.com/document/d/185vGMt4yQSJvDx8r8mjpq6-Ez9To75vlSYxLBGD-HpE/edit).

The Rook maintainers presented this graduation proposal to the SIG on February 26, 2020.  The
summary of the discussion can be found in the [agenda
document](https://docs.google.com/document/d/1DigEag4UUpf53qYBEr50YIdVJJvhXhxHN5ATj-js-IA/edit) for
this date.  Below is a high level summary of the discussion points:

* Project has met and exceeded criteria for graduation and continues to grow (doubling and tripling
  contributions statistics)
* Suggestion: Improve perception of Rook as a storage orchestrator (control plane) as opposed to
  being a storage provider (data plane) to reduce confusion.
  * The https://rook.io/ website has since been updated with messaging to make this distinction more
    clear to the public
* Question: How much work is it to create a new operator in the Rook project?
  * User experience could be improved for creating new operators, the focus so far has been on
    getting a single cohesive framework.
  * Each storage provider has individual needs, so not all things will be common and reusable across
    the providers. Each provider does have to write some logic and management code that is specific
    to their project.
* Question: How does Rook provide HA/Backup?
  * Rook relies on kubernetes best practices for running pods. There are health checks in place to
    ensure storage provisioned by Rook is available and ready.
  * The project has been maturing for 3 years, so we now feel confident about the hardening of the
    robustness of the control plane
  * Since Rook is not part of the data path, these concerns are lower as it’s not handling real time
    read and write requests
    * Operators are stateless, which lowers their complexity
