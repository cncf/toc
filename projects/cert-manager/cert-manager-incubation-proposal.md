# cert-manager Incubation Proposal

cert-manager joined the CNCF as a sandbox project just over a year ago, and we firmly believe that the project is ready to take the next step and graduate to the Incubation stage.

cert-manager provides reliable, automated Kubernetes-native management of X.509 certificates. It reduces the risk of certificate-related outages and enables PKI at scale, and is treated by many teams as de-facto part of the core components of a Kubernetes or OpenShift deployment.

The wider project comprises cert-manager itself - which adds Kubernetes Custom Resource Definitions (CRDs) for certificate-related concepts, and provides controllers to manage them - along with a selection of subprojects which focus on solving other issues relating to certificate management in cloud native environments. That said, and for the avoidance of doubt, we're only proposing cert-manager itself - the main repo - for incubation in this document.

The main cert-manager repo has 8.4k stars on GitHub and hundreds of projects downstream which depend on cert-manager code in some way. Thousands of people have contributed to the project in some way and cert-manager's use and prominence remains strong.

## Incubation Criteria Checklist

Based on the [graduation criteria][GraduationCriteria], cert-manager seems to clearly tick the boxes required for incubation:

### Production Use

> Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

cert-manager is used in production by _many_ more than three independent end users. Several major cloud native projects including [Istio][IstioCertManager], the [Kubernetes Cluster API][ClusterAPICertManager] and [Knative][KnativeCertManager] either suggest or require that cert-manager be installed.

Some other prominent examples include:

- The CNCF project ArgoCD uses cert-manager in its [demo instance][ArgoDemoInstance],
- Walmart, who mentioned it in a [Keynote speech they gave][WalmartKeynote],
- Render, who gave a great talk on cert-manager at [cert-manager Community Day 2020][RenderTalk]

According to the [companies-table DevStats dashboard][CompaniesTable] over a thousand different companies (plus ~1330 "Independents") have made some kind of contribution (defined as somebody who made a review, comment, commit, created PR or issue) to cert-manager. That implies that the floor of cert-manager's reach can be measured in the thousands, but there could be many more users who happily use cert-manager without contributing back or raising GitHub issues.

One prominent example of a cert-manager user might be the CNCF themselves, using cert-manager [in the DevStats project][DevStatsCertManager]!

In addition we might measure some production use through the many large companies who have seen fit to create addons for cert-manager to help it better integrate into their environments. That includes (at the time of writing):

- at least [49 public repositories][ExternalDNSSolvers] implementing an external ACME DNS-01 solver
- at least [8 public repositories][ExternalIssuers] implementing complete external issuers for certificates

Finally, we offer a place for users to identify themselves in a [USERS.md file][USERSmd] on our repository.

### Committer Numbers

> Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

cert-manager currently has 9 maintainers across 3 organisations, any of whom can accept changes to the codebase. Community contributions from many different organizations are very common and gladly received. The majority of the maintainers are part of the Jetstack cert-manager team, which has the open source cert-manager project as one of its primary focus points, although there are 2 other maintainers from other organizations - one at Apple, one at Thomas More University in Belgium.

The [companies-table DevStats dashboard][CompaniesTable] is illustrative of the number of committers and contributors we have. We've seen:

- 4 different companies which have made over a thousand contributions each,
- 8 different companies which have made over a hundred contributions each,
- and over a thousand different companies which have contributed something to the project.

In 2021, cert-manager saw [102 new first-time-contributors][NewContributors2021] to the project. Our latest release as of creating this proposal, cert-manager v1.7, saw at least 7 different individual community contributors, and one of the major features in that release was [contributed by a community member][ReleaseNotes17].

2022 has started strongly too; at the time of writing there have been [21 new contributors][NewContributors2022] already in this calendar year.

In terms of code changes only and ignoring other types of contribution, we see an [average of 10 different individual committers][CommittersCertManager] per month in the 6 months before February 2022 for cert-manager alone, rising to an [average of 20 per month][CommittersAllRepos] across all repositories in the same time period.

### Ongoing Commits

> Demonstrate a substantial ongoing flow of commits and merged contributions

All recent releases of cert-manager have included major features contributed by both the core maintainer team and community members, driven by a continuous flow of commits every month.

Looking at CNCF DevStats, cert-manager saw an average of [39 pull requests opened][AvgOpenedPRsCertManager] per month in the 6 months before February 2022 (excluding bots), with an average of [30 PRs merged per month][AvgMergedPRsCertManager] in the same time period.

The above numbers only count the main cert-manager repo; as mentioned previously, there are several cert-manager subprojects being regularly worked on too, including the cert-manager website. Across all repos we see an average of [78 merged PRs][AvgMergedPRsAllRepos] per month in the same time period as above.

### Versioning

> A clear versioning scheme

cert-manager has had a clear versioning scheme and a regular release cadence for over a year. Details can be seen on the cert-manager website under our [Supported Releases][] page, which documents full and clear support lifecycles for the main cert-manager project. The versioning scheme is similar to Kubernetes in that we try to avoid making any breaking changes wherever possible but if we do have to make a breaking change, we don't do so without signposting that the change is coming for several versions.

Being closely tied to Kubernetes, cert-manager also makes great efforts to be clear which versions of Kubernetes it is compatible with.

### Security Process

> Clearly documented security processes explaining how to report security issues to the project, and describing how the project provides updated releases or patches to resolve security vulnerabilities

A security reporting mechanism was added in 2021 with a clear [security process][] and a dedicated email address: `cert-manager-security@googlegroups.com`. This security policy applies to all subprojects under the cert-manager umbrella too.

Our security process states that we'll aim to reply to any reports as soon as we possibly can and that we'll try to formulate a plan of action within 24 hours of acknowledging the report.

## The cert-manager Project

### Encouraging Open Source

We endeavour to help people to contribute to cert-manager and to open source more widely, including:

- Two of our maintainers mentoring Google Summer of Code projects in 2021, with ambitions to get involved with GSoC 2022 too.
- Marking some of our repos for [Hacktoberfest](https://hacktoberfest.digitalocean.com/) for visibility
- Trying to find "good first issues" and mark them where possible, making it easier for new contributors to get started

We've run regular community meetings for years, too; daily standups are open to the public to join, and we also run a community meeting every other week which gives anyone a chance to participate. We're proud that some people have chosen to make their first ever open-source commits to the cert-manager project.

For support we maintain a [cert-manager-dev](https://groups.google.com/g/cert-manager-dev) Google group, and we're very active on Slack to help people out, too.

### Subprojects

The cert-manager GitHub org also includes a variety of subprojects which help with various certificate-related tasks in Kubernetes clusters. These include:

- [istio-csr][]: Enables the use of cert-manager for issuing [Istio](https://istio.io/) certificates
- [approver-policy][]: Adds the `CertificateRequestPolicy` resource to define policy over whether a given certificate should be signed by cert-manager
- [csi-driver][]: Kubernetes CSI plugin to facilitate seamlessly requesting and mounting certificate key pairs into pods
- [csi-lib][]: A generalised library for building CSI drivers that interact with cert-manager's `CertificateRequest` resource
- [trust][]: cert-manager, but for certificate trust bundles rather than for issuing certificates

Looking forward, we believe there's a lot of potential for the `trust` project. We believe many clusters will lack a coherent overview of their exposure related to
TLS trust, and the project's experience with TLS leads to a great opportunity for us to help with this issue.

NB: These tools are part of the wider cert-manager project but are not explicitly part of the application for incubation. We mention them as potential areas for future development for the project, and all of the tools touch on cert-manager in some way.

## Summary

We believe that cert-manager is ready for Incubation. The project closely aligns with CNCF values on openness and fairness; it's used by thousands upon thousands of people; and it seems to meet all of the required criteria to move to the next stage.

We're really looking forward to getting started with the process and we're keen to move as quickly as possible. We - the maintainers - will gladly make ourselves available to assist with any due diligence research, to make any presentations required and generally to assist in any way we can be helpful.

[GraduationCriteria]: https://github.com/cncf/toc/blob/da2c6f1a14ea9592ca0d1794da376e0c874b5a6f/process/graduation_criteria.adoc#incubating-stage

[IstioCertManager]: https://istio.io/latest/docs/ops/integrations/certmanager/
[DevStatsCertManager]: https://github.com/cncf/devstats#architecture
[ClusterAPICertManager]: https://cluster-api.sigs.k8s.io/developer/guide.html#cert-manager
[KnativeCertManager]: https://knative.dev/development/install/serving/installing-cert-manager/
[ArgoDemoInstance]: https://cd.apps.argoproj.io/applications/cert-manager
[WalmartKeynote]: https://www.youtube.com/watch?v=sfPFrvDvdlk&t=709s
[RenderTalk]: https://youtu.be/XlkGsBGk9W4
[UsersMD]: https://github.com/cert-manager/cert-manager/blob/5f5ba9eff67a5a8b23ef2b88bc4403c20ead979e/USERS.md

[ExternalDNSSolvers]: https://github.com/topics/cert-manager-webhook
[ExternalIssuers]: https://cert-manager.io/docs/configuration/external/

[CompaniesTable]: https://certmanager.devstats.cncf.io/d/5/companies-table?orgId=1

[NewContributors2021]: https://certmanager.devstats.cncf.io/explore?left=%5B%221609459200000%22,%221640995199000%22,%22psql%22,%7B%22alias%22:%22%22,%22dsType%22:%22influxdb%22,%22format%22:%22table%22,%22groupBy%22:%5B%7B%22params%22:%5B%22$__interval%22%5D,%22type%22:%22time%22%7D,%7B%22params%22:%5B%22null%22%5D,%22type%22:%22fill%22%7D%5D,%22orderByTime%22:%22ASC%22,%22policy%22:%22default%22,%22query%22:%22%22,%22rawQuery%22:true,%22rawSql%22:%22select%20count(*)%20from%20%5C%22snew_contributors_data%5C%22%20where%20$__timeFilter(dt)%20and%20series%20%3D%20%27ncdall%27%20and%20period%20%3D%20%27d%27%22,%22refId%22:%22A%22,%22resultFormat%22:%22table%22,%22select%22:%5B%5B%7B%22params%22:%5B%22value%22%5D,%22type%22:%22field%22%7D,%7B%22params%22:%5B%5D,%22type%22:%22mean%22%7D%5D%5D,%22tags%22:%5B%5D,%22datasource%22:%7B%22type%22:%22postgres%22,%22uid%22:%22P172949F98CB31475%22%7D,%22timeColumn%22:%22time%22,%22metricColumn%22:%22none%22,%22group%22:%5B%5D,%22where%22:%5B%7B%22type%22:%22macro%22,%22name%22:%22$__timeFilter%22,%22params%22:%5B%5D%7D%5D%7D%5D&orgId=1

[NewContributors2022]: https://certmanager.devstats.cncf.io/explore?left=%5B%221640995200000%22,%22now%22,%22psql%22,%7B%22alias%22:%22%22,%22dsType%22:%22influxdb%22,%22format%22:%22table%22,%22groupBy%22:%5B%7B%22params%22:%5B%22$__interval%22%5D,%22type%22:%22time%22%7D,%7B%22params%22:%5B%22null%22%5D,%22type%22:%22fill%22%7D%5D,%22orderByTime%22:%22ASC%22,%22policy%22:%22default%22,%22query%22:%22%22,%22rawQuery%22:true,%22rawSql%22:%22select%20count(*)%20from%20%5C%22snew_contributors_data%5C%22%20where%20$__timeFilter(dt)%20and%20series%20%3D%20%27ncdall%27%20and%20period%20%3D%20%27d%27%22,%22refId%22:%22A%22,%22resultFormat%22:%22table%22,%22select%22:%5B%5B%7B%22params%22:%5B%22value%22%5D,%22type%22:%22field%22%7D,%7B%22params%22:%5B%5D,%22type%22:%22mean%22%7D%5D%5D,%22tags%22:%5B%5D,%22datasource%22:%7B%22type%22:%22postgres%22,%22uid%22:%22P172949F98CB31475%22%7D,%22timeColumn%22:%22time%22,%22metricColumn%22:%22none%22,%22group%22:%5B%5D,%22where%22:%5B%7B%22type%22:%22macro%22,%22name%22:%22$__timeFilter%22,%22params%22:%5B%5D%7D%5D%7D%5D&orgId=1

[CommittersCertManager]: https://certmanager.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=cert-manager&var-country_name=All&var-company_name=All&var-company=all&from=1625094000000&to=1643673599000
[CommittersAllRepos]: https://certmanager.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=1625094000000&to=1643673599000

[ReleaseNotes17]: https://cert-manager.io/docs/release-notes/release-notes-1.7/#additional-certificate-output-formats

[AvgOpenedPRsCertManager]: https://certmanager.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&var-period=m&var-repogroup_name=cert-manager&from=1625094000000&to=1643673599000
[AvgMergedPRsCertManager]: https://certmanager.devstats.cncf.io/d/24/prs-merged-repository-groups?orgId=1&var-period=m&var-repogroups=%22cert-manager%22&from=1625094000000&to=1643673599000
[AvgMergedPRsAllRepos]: https://certmanager.devstats.cncf.io/d/24/prs-merged-repository-groups?orgId=1&var-period=m&var-repogroups=%22cert-manager%22&from=1625094000000&to=1643673599000

[Supported Releases]: https://cert-manager.io/docs/installation/supported-releases/
[security process]: https://github.com/cert-manager/cert-manager/blob/5f5ba9eff67a5a8b23ef2b88bc4403c20ead979e/SECURITY.md
[istio-csr]: https://github.com/cert-manager/istio-csr
[approver-policy]: https://github.com/cert-manager/approver-policy
[csi-driver]: https://github.com/cert-manager/csi-driver
[csi-lib]: https://github.com/cert-manager/csi-lib
[trust]: https://github.com/cert-manager/trust
