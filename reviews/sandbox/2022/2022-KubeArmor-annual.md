# KubeArmor Sandbox Annual Review

## Background

[KubeArmor](https://kubearmor.io/) is a runtime security engine for kubernetes and other cloud workloads. Users can harden their workloads, create Zero Trust security posture i.e., allowing specific systems actions and denying rest using KubeArmor security policies. KubeArmor supports [inline mitigation](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/differentiation.md) for preventing attacks. Due to its versatility, KubeArmor has found its relevance in Edge and 5G Control Plane security scenarios as well.

KubeArmor primarily provides the following key features:

* [Workload Hardening](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/hardening_guide.md)

    KubeArmor provides a set of hardening policies that are based on industry-leading compliance and attack frameworks such as CIS, MITRE, NIST-800-53, and STIGs. These policies are designed to help you secure your workloads in a way that is compliant with these frameworks and recommended best practices.

* [Zero Trust Policies](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/least_permissive_access.md)

    KubeArmor is a tool that helps organizations enforce a zero trust posture within their Kubernetes clusters. It allows users to define an allow-based policy that specifies the specific system behavior that is allowed, and denies or audits all other behavior.
    
* [Network Visibility & Microsegmentation](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/network_segmentation.md)

    KubeArmor has visibility into the network connections made into or from the pods. It has visibility across socket(), bind(), connect(), accept() calls. KubeArmor can help auto-generate k8s network policy which are enforceable using most of the k8s CNIs (Container Network Interface). 

* [Application Behavior](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/workload_visibility.md)

    KubeArmor has visibility into systems and application behavior. KubeArmor summarizes/aggregates the information and provides a user-friendly view to figure out the application behavior.

### Alignment with CNCF

- KubeArmor was accepted as a CNCF Sandbox project on November 16, 2021.
- KubeArmor provides a way to enforce security in k8s native way by leveraging k8s resource model.

## Development metrics

The KubeArmor devstats page and dashboards can be found [here](https://kubearmor.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, KubeArmor currently has [169](https://kubearmor.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1) contributors from [31](https://kubearmor.devstats.cncf.io/d/5/companies-table?orgId=1) companies belonging to 12 countries.
- The community has significantly grown since the project entered the CNCF sandbox.
  - We held biweekly community meetings consistently (total 43 since Sep 2021). The community did not skip a single meeting since its inception. The meeting records can be found in [here](https://docs.google.com/document/d/1IqIIG9Vz-PYpbUwrH0u99KYEM1mtnYe6BHrson4NqEs/edit).
  - Number of contributors: 30+ -> **150+**
  - Github stars: 100+ -> **1070+**
  - Github forks: 30+ -> **265**
  - Contributing organizations: 5+ -> **30+**
- [New PRs in last year](https://kubearmor.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1)
- KubeArmor maintainer team has mentored 10 candidates as part of LFX and GSoC mentorships.

## Maintainers

Initial maintainers were from AccuKnox primarily. We have an independent maintainer and few other contributors who are shaping up to take the ownership of the modules. Following is the list sorted in alphabetical order:

- Achref Ben Saad (AccuKnox) @achrefbensaad
- Ankur Kothiwal (CERN) @Ankurk99
- Anurag Kumar (Independent) @kranurag7
- Barun Acharya (AccuKnox) @daemon1024
- Jaehyun Nam (Dankook Univ) @nam-jaehyun
- Rahul Jadhav (AccuKnox) @nyrahul

## Project adoption

Adoption of KubeArmor is tracked in our [ADOPTERS.md](https://github.com/kubearmor/kubearmor/blob/master/ADOPTERS.md) file. We have recently started maintaining the ADOPTERS doc.

Beyond this, we have received interests from other organizations such as:
- [LF 5G Super Blueprint to use KubeArmor](https://wiki.lfnetworking.org/pages/viewpage.action?pageId=82905466) for 5G security.
- [KubeArmor Operator](https://catalog.redhat.com/software/containers/kubearmor/kubearmor-operator/6458b1ee3f880286315e7cfb) is also available on Red Hat Ecosystem Catalog.
- Securing 5G Control Plane using KubeArmor. [[ref1](https://nephio.org/5g-security-enhanced-open-radio-access-networks/)] [[ref2](https://5gsec.com/5G_Security/5G-Kubearmor.html)]

## Project goals

### Initial Goal

Some notable goals include:
* Simplifying use of Linux Security Modules (such as AppArmor, SELinux, BPF-LSM) for hardening workloads. Aim is to [provide a way](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/differentiation.md) for users to prevent attacks using inline mitigation techniques, rather than post-attack mitigations which is prevalent currently.
* Seamless deployments on restricted hardened distributions such as [Bottlerocket](https://aws.amazon.com/blogs/containers/secure-bottlerocket-deployments-on-amazon-eks-with-kubearmor/), COS, UEK, etc.
* [Support](https://github.com/kubearmor/KubeArmor/wiki/v0.9-Release-Blog) for managed, unmanaged, onprem kubernetes. Support for x86, ARM based k8s platforms.
* Performance improvements and scalability for 1000s of nodes in a cluster.
* ... and many more. The release blogs can be found [here](https://github.com/kubearmor/KubeArmor/wiki#version-release-blogs).

KubeArmor has a release cadence of once in two month release cycle.

The backlog/roadmap for KubeArmor can be found [here](https://github.com/orgs/kubearmor/projects/9/views/1).

## How the CNCF can help to achieve the upcoming goals?

- More chances to advocate the project and reach out to great audience. Many thanks for the Open Source KubeArmor booth in Kubecon 2023, Amsterdam.
- More chances to collaborate with other projects in CNCF or even out of CNCF (such as Linux Foundation, or Apache Foundation). We believe there is a huge scope for KubeArmor to secure Edge and 5G control plane deployments by utilizing security controls enumerated by [ENISA](https://www.enisa.europa.eu/publications/5g-security-controls-matrix), [MITRE FIGHT](https://fight.mitre.org/) etc.
- Technical writing support for project documents.
- KubeArmor has involved in all LFX and GSoC mentorship calls from CNCF since KubeArmor's inception and would continue on that path.
- None of the maintainers of KubeArmor are in western or european timezones and it limits our ability to improve project's visibility. Any advice/guidance would be helpful.

## Incubation readiness

- We think KubeArmor is ready for Incubation and would love some advice from the CNCF team as to how do we go about next steps. The maintainers team have looked at the Incubation guidelines and KubeArmor fulfills all the aspects.
