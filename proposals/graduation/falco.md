# Falco Graduation Proposal

Falco, the Runtime Security Project, joined the CNCF in October 2018 as a Sandbox Project.

On January 8th, 2020 Falco was [promoted to the Incubating stage](https://www.cncf.io/blog/2020/01/08/toc-votes-to-move-falco-into-cncf-incubator/). Since then, the Falco community, users and maintainers continued to actively invest in Falco, fueling sustained growth in terms of adopters, governance, contributions, features and integrations with other CNCF projects.

At this stage, we believe that Falco is ready for the next natural step: Graduation.

Key highlights as of April 19th, 2021:

- Github stars: 4.200+
- Core repository contributors: 100+
- Forks: 660+
- Releases: 110
- Slack channel members: 750+
- Downloads:
  - 24M+ Docker hub pulls (+570% since Incubation)
  - 30M+ driver downloads from our CDN (+680% since Incubation)
	- Since Falco compiles a driver for the host when it starts (and fallbacks to download a prebuilt one) this is a very conservative proxy of how many times Falco is installed
- Top 3 media mentioned Incubated CNCF project for last 3 months (January, February, March 2021)
- Keynote mention at Kubecon+CloudNativeCon 2020 -- Constance Caramolis ([youtube](https://www.youtube.com/watch?v=uNpZA5fz5o8))
- End user talk at Kubecon 2020 -- Shane Lawrence, Shopify ([youtube](https://www.youtube.com/watch?v=rBqBrYESryY))
- End user talk at Kubecon 2020 -- Natch Ruengsakulrach & Eric Hollis, MathWorks ([slides](https://static.sched.com/hosted_files/kccncna20/aa/KubeCon_NA_Virtual_2020-Cyber_Kill_Chain_Falco.pdf))
- Inclusion in CNCF CKS certification
- Several vendors have built commercial offerings on top of Falco (e.g. [GitLab](https://about.gitlab.com/blog/2020/08/18/how-gitlab-can-help-you-secure-your-cloud-native-applications/), [SumoLogic](https://www.sumologic.com/solutions/kubernetes/))
- The Falco organization is evolving from a single tool into an ecosystem of projects, some of which are gaining substantial traction (e.g. [falcosidekick](https://github.com/falcosecurity/falcosidekick): 390K+ Docker hub pulls, [falco-exporter](https://github.com/falcosecurity/falco-exporter): 260K+ Docker hub pulls)

## Graduation State Criteria

### ✅ Have committers from at least two organizations.

Since incubation, the amount of contributions coming from organizations other than Falco original creator (Sysdig) has grown steadily. During the last year, around 45% of the contributions came from a diverse group of committers that includes Innoteam, Amazon, Samsung, IBM, Mercari, Red Hat and many individual contributors. More details are available [here](https://falco.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions).

[Maintainers](https://github.com/falcosecurity/.github/blob/master/maintainers.yaml) of [Falco projects](https://github.com/falcosecurity) also come from a wide set of companies (Amazon, IBM, VMWare, RedHat, Sysdig, Qonto, move:elevator GmbH, Mercari, Timber).

### ✅ Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

Falco meets CII Best Practices passing level criteria (100%). More details [here](https://bestpractices.coreinfrastructure.org/en/projects/2317).
Also, the current score for silver criteria is 49% while gold level is at 26%.

### ✅ Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

Falco underwent a CNCF sponsored, independent security audit from Cure53 in July 2019, which is available in the Falco repository [here](https://github.com/falcosecurity/falco/blob/master/audits/SECURITY_AUDIT_2019_07.pdf). Issues found during the audit were then [tracked and addressed](https://github.com/falcosecurity/falco/security/advisories).

The Falco project is looking to undergo a second security audit later this year.

### ✅ Explicitly define a project governance and the contributing process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

The governance process for the Falco project is well described in our organization-wide [GOVERNANCE.md](https://github.com/falcosecurity/.github/blob/master/GOVERNANCE.md), in which we have clear rules and processes for:

- Maintainers appointment
- Maintainers duties and responsibilities
- Conflict resolution
- Addition/removal of repositories

New repositories and organization resources have a dedicated workflow which is regulated inside the [evolution](https://github.com/falcosecurity/evolution) repository.

Maintainers for all projects are listed in a dedicated [yaml file](https://github.com/falcosecurity/.github/blob/master/maintainers.yaml) inside [falcosecurity/.github](https://github.com/falcosecurity/.github) repository. This file is automatically updated using OWNERS files stored inside repositories and repositories subfolders.

Last but not least, the contributing process for the whole organization is described [here](https://github.com/falcosecurity/.github/blob/master/CONTRIBUTING.md).

### ✅ Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.

Falco adopters are listed on both our [official website](https://falco.org/) and into the [ADOPTERS.md](https://github.com/falcosecurity/falco/blob/master/ADOPTERS.md) inside our core repository.

## Incubation Detail

### Link to Incubation Due Diligence(DD) Document

Incubation DD document can be found [here](https://docs.google.com/document/d/1TJCzW8dQ6858lw2UNY-H5LMnvEd4GzwjuOcDInimeyA/edit?ts=5dacfd96#heading=h.378jkvcve1nq).

### Address any concerns or recommendations from the SIG and/or TOC sponsor(s) from the DD Document

One important concern emerged during the incubation due diligence process was related to the codebase’s vendor independence: some of the Falco key data collection dependencies (libsinsp, libscap, kernel module and eBPF probe) were coming from an OSS repository under the Sysdig Inc. GitHub organization.

To address this issue, Sysdig has since contributed those components to the CNCF via the Falcosecurity organization (now available at [falcosecurity/libs](https://github.com/falcosecurity/libs)). Read more on the CNCF blog [here](https://www.cncf.io/blog/2021/02/24/sysdig-contributes-falcos-kernel-module-ebpf-probe-and-libraries-to-the-cncf/).

Another concern was about the ownership of the release artifacts since they were under Sysdig's controlled S3 buckets. This has also been addressed, and now all the artifacts (rpm, deb, kernel modules, eBPF probes) are hosted under a CNCF-owned AWS account and published to https://download.falco.org/. This process, along with all other automations, is maintained and orchestrated in the [test-infra](https://github.com/falcosecurity/test-infra) repository.
