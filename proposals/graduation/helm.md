# Helm Graduation Proposal

Helm originally joined the CNCF when Kubernetes became the first CNCF project. Helm was a sub-project of Kubernetes at that time. Since joining the CNCF as an incubating sister project to Kubernetes in June 2018, Helm has released v3, created a governance appropriate for Helm, and has continued to watch the use of Helm grow and evolve.

On behalf of the maintainers, we believe Helm is ready for [graduation stage](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#graduation-stage) and meets the v1.3 criteria. The Helm [due diligence questions](https://docs.google.com/document/d/1-5ncwLwikeQrAmeVAPuCYlWLYMCw2OmQuouTYCECYFk/edit), following the [guidelines](https://github.com/cncf/toc/blob/master/process/due-diligence-guidelines.md), can be found in a separate Google Docs document.

Helm has seen dramatic growth since joining incubation. For example,

* [At KubeCon/CloudNativeCon US in San Diego Helm was 3rd most popular project](https://www.cncf.io/wp-content/uploads/2020/01/KubeCon_NA_19_Report.pdf).
* Helm now has more than 1 million downloads in a months time.
* Installations via Homebrew for MacOS have grown more than 12 times since entering incubation.
* Helm v3, a major change, was released

## Graduation State Criteria

_**Project should address each graduation criteria listed below**_

### * Have committers from at least two organizations.

The [Helm project governance](https://github.com/helm/community/blob/master/governance/governance.md) defines two types of maintainers in the Org Maintainers who oversee the overall project and those who are responsible for different sub-projects (e.g., the Helm CLI, the charts repository, etc). The Org Maintainers currently have people from 7 companies and the governance does not allow a single company to have a majority of representatives. The Helm CLI, the primary project, currently has maintainers from 7 companies.

### * Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

The Helm CLI has achieved the [CII best practices badge](https://bestpractices.coreinfrastructure.org/en/projects/3131) and is currently close to obtaining the silver badge (more than 95% of the way there). This includes having a [security assurance case](https://github.com/helm/community/tree/master/security-assurance-case).

### * Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

The Helm CLI has completed a 3rd party security review. [You can find the details and a link to the full report on the Helm blog](https://helm.sh/blog/2019-11-04-helm-security-audit-results/). In the conclusion of the report the authors noted:

> To conclude, in light of the findings stemming from this CNCF-funded project, Cure53 can only state that the Helm project projects the impression of being highly mature. This verdict is driven by a number of different factors described above and essentially means that Helm can be recommended for public deployment, particularly when properly configured and secured in accordance to recommendations specified by the development team.

### * Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

Helm has a [documented governance](https://github.com/helm/community/blob/master/governance/governance.md) that outlines numerous things including roles, responsibilities, how decisions are made, how maintainers are elected and removed, and more. The [Helm teams are documented](https://github.com/helm/community/blob/master/Teams.md) and that includes links to the OWNERS files for each team.

### * Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.

Many companies and organizations use Helm. Helm has been downloaded more than a million time in a month, across all the Helm sub-projects there are maintainers from more than 25 companies, the [Helm Hub](https://hub.helm.sh/charts) lists repositories from more than 150 organizations, and many more organizations than that have their people in the Helm issues asking support questions. The [Helm website](https://helm.sh) has long had a list of adopters. More recently the Helm code has added an [adopters file](https://github.com/helm/helm/blob/master/ADOPTERS.md) with a select number of organizations. Organizations are only added when the organization provides permission.

## Incubation Details

_**Project should address each area listed below**_

### * Link to Incubation Due Diligence(DD) Document

Helm entered the CNCF as an incubating project after being part of Kubernetes when it came into the CNCF. Helm was a part of Kubernetes when it graduated. The Helm due diligence at the time is captured in the [incubating proposal](https://github.com/cncf/toc/blob/master/proposals/incubation/helm.adoc).

### * Address any concerns or recommendations from the SIG and/or TOC sponsor(s) from the DD Document

There are currently no known outstanding concerns or recommendations with Helm.
