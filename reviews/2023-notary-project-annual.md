# Notary Project CNCF Annual Review 2023

## Background

The Notary Project is a set of specifications and tools intended to provide a cross-industry standard for securing software supply chains by using authentic container images and other OCI artifacts. Notary Project specification and tooling provides signing and verification workflows for OCI artifacts, signature portability across OCI compliant registries, and integration with 3rd party key management solutions through a plugin model. See the [Project overview](https://github.com/notaryproject/.github#notary-project-overview) for more details about the project history and background.

Notary Project has been accepted as an incubating project by CNCF in 2019. After several years of iteration, the Notary Project maintainers [announced the major release on Aug 22, 2023](https://notaryproject.dev/blog/2023/announcing-major-release/), including [Notary Project specifications v1.0.0](https://github.com/notaryproject/specifications/releases/tag/v1.0.0), [notation v1.0.0](https://github.com/notaryproject/notation/releases/tag/v1.0.0), [notation-go v1.0.0](https://github.com/notaryproject/notation-go/releases/tag/v1.0.0), and [notation-core-go v1.0.0](https://github.com/notaryproject/notation-core-go/releases/tag/v1.0.0) which are ready for production use. 

As part of our commitment to security, the Notary Project maintainers engaged with CNCF to set up continuous fuzzing of the source code and completed a security audit in 2023. All vulnerabilities found during the testing and the audit were fixed before the release of the libraries and the CLI. Below are links to the security reports:

- [Notation Security Audit Report 2023](https://github.com/notaryproject/notaryproject/blob/main/security/reports/audit/ADA-notation-security-audit-23.pdf)
- [Notary Project fuzzing audit report 22-23](https://github.com/notaryproject/notaryproject/tree/main/security/reports/fuzzing/ADA-fuzzing-audit-22-23.pdf)

You can learn more about the Notary Project on the [notaryproject.dev](https://notaryproject.dev) website. 

## Project’s devstats page

The Notary Project devstats page and dashboards can be found [here](https://notary.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

- According to devstats, contributors in all Notary Project repositories increased to [209](https://notary.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20decade&var-repogroup_name=All) totally, up from [63](https://notary.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Before%20joining%20CNCF&var-repogroup_name=All) before joining CNCF. 
- [26](https://notary.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now) new contributors engaged in all Notary Project repositories over the past year (from Oct 2022 to Oct 2023).
- PR contributors in all Notary Project repositories come from [23 organizations](https://notary.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=prcreators) over the past year (from Oct 2022 to Oct 2023). Now PR contributors in all Notary Project repositories come from [70](https://notary.devstats.cncf.io/d/21/prs-authors-companies-table?orgId=1) organizations in total.
- All [Stars](https://notary.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=now-1y&to=now) increased from 3146 to 3513, all [Forks](https://notary.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=1665331200000&to=1696953599000) increased from 658 to 808 over the past year (from Oct 2022 to Oct 2023).
- Notation CLI, SDK, and Spec had 8, 10, and 2 releases respectively in 2023.
- Organized 84 community meetings in 2023, see [archived meeting notes](https://github.com/notaryproject/meeting-notes).

Regarding these metrics listed above, we would perceive the Notary Project community as healthy and growing in 2023.

## Maintainers

The Notary Project community has established [the community governance](https://github.com/notaryproject/.github/blob/main/GOVERNANCE.md) to engage active and qualified contributors. Now Notary Project has 14 maintainers from different organizations.

- Feynman Zhou, Microsoft
- Junjie Gao, Microsoft
- Justin Cormack, Docker
- Niaz Khan, Amazon
- Milind Gokarn, Amazon
- Pritesh Bandi, Amazon
- Patrick Zheng, Microsoft
- Rakesh Gariganti, Amazon
- Samir Kakkar, Amazon
- Steve Lasker, RKVST
- Shiwei Zhang, Microsoft
- Toddy Mladenov, Microsoft
- Vani Rao, Individual
- Yi Zha, Microsoft

## Project Adoption

| Organization | Website | Scenario/Use case | 
| --- | --- | --- | 
| Amazon | https://aws.amazon.com/ | Using Notation in AWS Signer for signing artifact, see the [reference](https://aws.amazon.com/blogs/containers/announcing-container-image-signing-with-aws-signer-and-amazon-eks/) |
| containerd | http://containerd.io/ | Integrate Notation signing and verification in containerd nerdctl, see the [reference](https://github.com/containerd/nerdctl/blob/main/docs/notation.md) |
| Harbor | https://goharbor.io/ | Support Notation signature and adds Notation policy checker, see the [reference](https://goharbor.io/blog/harbor-2.9/)
| Kyverno | https://kyverno.io/ | Image signing and verification, see the [reference](https://kyverno.io/docs/writing-policies/verify-images/notary/) |
| Microsoft | https://www.microsoft.com/ | Sign and verify artifact, see the [reference](https://techcommunity.microsoft.com/t5/apps-on-azure-blog/announcing-notation-azure-key-vault-plugin-v1-0-for-signing/ba-p/3920895) |
| Ratify |  https://ratify.dev/ | Validating OCI Artifact signatures using Notation library, see the [reference](https://ratify.dev/docs/1.0/reference/crds/verifiers#notation) |
| Venafi | https://www.venafi.com/ | Build Venafi CodeSign Protect Signature and Verification Plugin for Notation CLI, see the [reference](https://github.com/Venafi/notation-venafi-csp)  |
| Zot | https://zotregistry.io | Support Notary Project signature storage, see the [reference](https://github.com/project-zot/zot/tree/main/demos#image-signatures)

## Project Goals

Notary Project provides a set of specifications and tools intended to build a cross-industry standard for securing software supply chains. Notary Project aims to help protect enterprises, open-source communities, and users from supply chain attacks, ensures produced software artifacts have not been tampered with and provides security policies to determine which validated artifacts are allowed to be used in users' environments. It provides signing and verification approach to help software publishers and users establish integrity and authenticity of content distribution.

The long-term project [goals](https://github.com/notaryproject/specifications/blob/main/requirements/requirements.md#goals) and [scenarios](https://github.com/notaryproject/specifications/blob/main/requirements/scenarios.md) are available on the [Notary Project Specifications](https://github.com/notaryproject/specifications) repository.

In the near future, the upcoming milestone scheduled for the final quarter of 2023 is [v1.1.0](https://github.com/notaryproject/notation/milestone/17) slated for release in late November 2023. The following enhancements are considered in this release: 

- Sign and verify arbitrary data
- Sign local container images before they are pushed to registries
- Improve user experience for plugin management
- Improve user experience of managing trust store and trust policy

The Notary Project community uses the [project planning board](https://github.com/orgs/notaryproject/projects/10) for roadmap and status tracking. Users can also use the GitHub milestones to track the detailed plan and progress of each repository.

## How can the CNCF help achieve upcoming goals

- Help maintainers measure the project adoption metrics
- Marketing support and get more chances to promote the project to broader users
- More chances to collaborate with other security ecosystem projects in CNCF
- Technical writing support