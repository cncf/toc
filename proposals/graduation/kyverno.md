# Kyverno Graduation Proposal

[Kyverno](https://kyverno.io/) is a Kubernetes policy engine that was created to validate, mutate, generate, and cleanup  Kubernetes resources, including custom resources. Kyverno also offers built-in image verification rules to help secure the software supply chain by verifying signatures and attestations in either Sigstore Cosign and CNCF Notation formats. 

Kyverno policies are managed as Kubernetes resources and no new language is required to write policies. Kyverno also uses Kubernetes API objects for policy reporting and managing policy exceptions, making it a popular choice for platform teams using Kubernetes.

Kyverno was accepted as a CNCF Sandbox project in November 2020, and graduated to Incubating status in July 2022. The GitHub stars for Kyverno itself has experienced an impressive [growth](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=Kyverno&var-metric=Stargazers&var-table=swatchers&var-pref=&var-met1=watch&var-met2=watch&from=1656604800000&to=now) from 2537 to [4700](https://github.com/kyverno/kyverno), and there are 2,500+ members registered for the [Kyverno community Slack](https://main.kyverno.io/community/#slack-channel). The Kyverno organization has [1452](https://kyverno.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=All) contributors from [280+](https://kyverno.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=contributions) companies. Kyverno is used in production and at scale by several end user organizations like Deutsche Telekom, Spotify, The US Department of Defense, LinkedIn, Vodafone, and Yahoo.

The project maintainers believe that Kyverno meets the requirements for Graduation status as detailed below:

## Graduation State Criteria

### * Have committers from at least two organizations.

Kyverno has maintainers from 4 different organizations, see [maintainers.md](https://github.com/kyverno/kyverno/blob/main/MAINTAINERS.md).

1. Nirmata
2. Stackwatch (Kubecost)
3. Giant Swarm GmbH
4. Ohio Supercomputer Center

### * Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

Kyverno has achieved an [OpenSSF Best Practices Badge](https://www.bestpractices.dev/en/projects/5327).

### * Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

Kyverno completed a [fuzzing security audit](https://kyverno.io/blog/2023/09/06/kyverno-completes-fuzzing-security-audit/) and a [third-party security review](https://kyverno.io/blog/2023/11/28/kyverno-completes-third-party-security-audit/). 

The security review was conducted in collaboration with the [CNCF](https://www.cncf.io/), [Ada Logics](https://adalogics.com/) and [OSTIF](https://ostif.org/). The Kyverno project has addressed all issues from the audits.

### * Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

The Kyverno project governance policies are documented in [GOVERNANCE.md](https://github.com/kyverno/kyverno/blob/main/GOVERNANCE.md).

The committer process in Kyverno is defined in project roles [Contributors](https://main.kyverno.io/community/#contributors) and [Code Owners](https://main.kyverno.io/community/#code-owners) documents, which includes the full committer lifecycle, onboarding, offboarding, and emeritus criteria. The committers are listed in [OWNERS.md](https://github.com/kyverno/kyverno/blob/main/OWNERS.md).

### * Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

The project role [maintainers](https://main.kyverno.io/community/#maintainers) covers the criteria, process, offboarding and emeritus conditions. Maintainers may be subject to removal, based on a public vote, if they have made less than 30 contributions over a span of 6 months. Maintainers who are removed will be moved to an emeritus status.

The current maintainers list can be found in [MAINTAINERS.md](https://github.com/kyverno/kyverno/blob/main/MAINTAINERS.md).

### * Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to [FAQs](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for guidelines on identifying adopters.

The list of organizations and users that have publicly shared how they are using Kyverno is managed at [ADOPTERS.md](https://github.com/kyverno/kyverno/blob/main/ADOPTERS.md).

## Incubation Details

### * Link to Incubation Due Diligence(DD) Document

* [Incubation proposal](https://github.com/cncf/toc/pull/784)
* [Incubation Due Diligence](https://docs.google.com/document/d/18dWgOd2MUQz3RXI1R9vKntL3ULyZhOD1HEtijGOeaWg/edit#heading=h.amgfsmvtn6jy)

### * Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

There are no outstanding issues from the Incubation level due diligence. The only issue identified, and its resolution, are detailed below:

#### Use the official CNCF Kyverno Zoom account 

An official Kyverno Zoom account was requested from the CNCF and all project meetings now use this account.

