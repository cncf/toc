# Falco Graduation Proposal

Falco is a threat detection engine providing runtime security alerting for containerized applications and the cloud. It joined the CNCF as a sandbox project in 2018-10-10 and evolved into incubation on 2020-01-08.

Falco is in widespread use within Kubernetes deployments, including at **AWS, Qonto, and Shopify**; and is integrated into a variety of derivative products such as **R6/Phoenix, Giant Swarm/Falco App, gVisor, StackRox, Tarian, Sysdig Secure,** among others.

**Website:** falco.org
**GitHub:** [falcosecurity org](https://github.com/falcosecurity)
**Core repositories**:

|                                         Repository                                          |                               Description               |
|---------------------------------------------------------------------------------------|--------------------------------------------------------------------------|
| [falcosecurity/charts](https://github.com/falcosecurity/charts)                       | Helm charts for running Falco with Kubernetes                            |
| [falcosecurity/deploy-kubernetes](https://github.com/falcosecurity/deploy-kubernetes) | Kubernetes deployment resources for Falco                                |
| [falcosecurity/falco](https://github.com/falcosecurity/falco)                         | Main Falco repository                                            |
| [falcosecurity/falco-website](https://github.com/falcosecurity/falco-website)         | Source code of [falco.org](https://falco.org) website      |
| [falcosecurity/libs](https://github.com/falcosecurity/libs)                           | libsinsp, libscap, the kernel module driver, and the eBPF driver sources |
| [falcosecurity/plugin-sdk-go](https://github.com/falcosecurity/plugin-sdk-go)         |                                                                         Plugins SDK |
| [falcosecurity/plugins](https://github.com/falcosecurity/plugins)                     | Plugins registry                                                   |



**Key highlights as of 2021-11-04 & growth since April 2021**
* Github stars: **5,300+ (26% growth)**
* Core repo contributors: **165 (65% increase)**
* Forks: **708 (7% increase)**
* Slack channel members: **1,695 (126% growth)**
* Downloads:
    * **41.7M** Docker hub pulls (**25% growth**)
    * **1.2M** AWS ECR 
* Contributor community growth: the number of contributors measured by devstats has grown from **55** at Incubation acceptance, to **165** as of Q3 2022.
* Number of organizations with maintainers: since incubation the project has added maintainers from **4** new organizations.
* Features: the most significant additions since acceptance for Incubation include:
    * Stable release schedule
    * Contribution of the whole *libs* source code into the Falcosecurity organization
    * Plugin framework, allowing event sources other than syscalls
    * new eBPF probe
    * gVisor integration


## Graduation state criteria

### Have committers from at least two organizations.

Since incubation, the amount of contributions coming from organizations other than Falco original creator (Sysdig) has grown steadily. During the last year, around **38%** of the contributions came from a diverse group of committers that includes: **Innoteam, Amazon, Samsung, IBM, Mercari, RedHat and many individual contributors**. More details are available [here](https://falco.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions). 

The core repositories have maintainers (committers) from these organizations:

- Clastix
- Chainguard
- Polytechnic of Turin
- RedHat
- Sysdig

In addition, the following organizations are represented by maintainers (committers) among non-core repositories:

- Amazon
- IBM
- LOVOO
- Secureworks


The [MAINTAINERS.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS.md) lists all maintainers and their affiliations.

### Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

Falco meets CII Best Practices passing level criteria (100%). More details [here](https://bestpractices.coreinfrastructure.org/en/projects/2317).

### Have completed an independent and third party security audit with results published of similar scope and quality as this example which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.


Falco underwent a CNCF sponsored, independent security audit from Cure53 in July 2019, which is available in the Falco repository [here](https://github.com/falcosecurity/falco/blob/master/audits/SECURITY_AUDIT_2019_07.pdf). Issues found during the audit were then [tracked and addressed](https://github.com/falcosecurity/falco/security/advisories).

A further CNCF sponsored, independent security audit from Quarkslab was completed in January 2023, and no critical vulnerabilities were found. Other issues found were [tracked and addressed](https://github.com/falcosecurity/libs/security/advisories).

### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

Falco has a comprehensive governance policy, recently clarified and reviewed with the Governance WG of the CNCF TAG Contributor Strategy. Highlights of the governance framework include:

- Principles: openness, respect, transparency and vibrant evolution are Falco’s principles.
- Leadership: maintainers of the core repositories serve as the oversight and direction-setting group for the whole project, with limits set on per-company voting power.
- Maintainer lifecycle: clear definition of how project members move from contributor to maintainer, to retiring to emeritus status.
- Decision-making and voting: clear rules on lazy consensus by default, and how votes should be conducted.


Docs laid out here:

- The [GOVERNANCE.md]( https://github.com/falcosecurity/evolution/blob/main/GOVERNANCE.md) document defines the project governance and processes. 
- The [evolution](https://github.com/falcosecurity/evolution) repository, which aims to document the evolution process of The Falco Project, provides complementary documentation.
- [OWNERS](https://www.kubernetes.dev/docs/guide/owners/) files (one per repository) are used to define approvers (ie. maintainers), reviewers and emeritus. Examples: [falco repository OWNERS](https://github.com/falcosecurity/falco/blob/master/OWNERS), [libs repository OWNERS](https://github.com/falcosecurity/libs/blob/master/OWNERS), [charts respository OWNERS](https://github.com/falcosecurity/charts/blob/master/OWNERS).
- The [MAINTAINERS.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS.md) file is automatically generated and summarizes the content of all OWNERS files across all repositories.


### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.


Maintainership-related decisions are taken with respect to the rules established in our [governance](https://github.com/falcosecurity/evolution/blob/main/GOVERNANCE.md#maintainership). The [MAINTAINERS_GUIDELINES.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS_GUIDELINES.md) document provides guidelines for the implementation of these decisions.

Maintainers are defined by [OWNERS](https://www.kubernetes.dev/docs/guide/owners/) files on a per-repository basis. The [MAINTAINERS.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS.md) file summarizes the content of all OWNERS files across all repositories. The document is automatically updated by our infra and frequently audited using the PR review process.

As per our governance, core maintainers (listed by [MAINTAINERS.md#core-maintainers](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS.md#core-maintainers)) are those maintaining the [core repositories](https://github.com/falcosecurity/evolution#official) of the project. Core maintainers are responsible for overseeing the overall project health and growth, and speaking on behalf of the project, and also entitled to interact with the CNCF.




### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to FAQs for guidelines on identifying adopters.

The public list is in the [ADOPTERS.md](https://github.com/falcosecurity/falco/blob/master/ADOPTERS.md) file.

## Incubation Details

### Link to Incubation Due Diligence(DD) Document

[Incubation DD Document](https://docs.google.com/document/d/1TJCzW8dQ6858lw2UNY-H5LMnvEd4GzwjuOcDInimeyA/edit#)

### Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

No outstanding issues raised during the incubation due diligence remain to be addressed.