# Cilium Graduation Proposal

[Cilium](https://cilium.io)  is a networking, observability, and security solution with an eBPF based dataplane. As a CNI, it provides a simple flat Layer 3 network with the ability to span multiple clusters in either a native routing or overlay mode. Cilium is L7-protocol aware and can enforce network policies on L3-L7 using an identity based security model that is decoupled from network addressing.

Cilium implements distributed load balancing for traffic between pods and to external services, and is able to fully replace kube-proxy, using efficient hash tables in eBPF allowing for almost unlimited scale. It also supports advanced functionality like integrated ingress and egress gateway, bandwidth management and service mesh, and provides deep network and security visibility and monitoring through Hubble and Tetragon.


Cilium itself has 16,200+ GitHub stars and 624 contributors and there are 22,000+ GitHub stars and 800+ contributors across the various repos in the [Cilium GitHub org](https://github.com/cilium) and 16,500+ users registered for the [Cilium community Slack](https://cilium.herokuapp.com/). Cilium was [open sourced in 2015](https://github.com/cilium/cilium/commit/7fa3c60eb7dbe7a5a4caea3aab0396f75a8b10c7), [hit the 1.0 milestone in 2018](https://cilium.io/blog/2018/04/24/cilium-10/), and [joined the CNCF at Incubation level in October 2021](https://www.cncf.io/blog/2021/10/13/cilium-joins-cncf-as-an-incubating-project/). It currently powers production systems at scale all around the world, including at
[Adobe](https://www.youtube.com/watch?v=7UQ2CU6UEGY), [Bell Canada](https://www.youtube.com/watch?v=vJaOKGWiyvU), [CapitolOne](https://www.youtube.com/watch?v=hwOpCKBaJ-w&ab_channel=eBPFSummit), 
[Datadog](https://www.cncf.io/case-studies/datadog/)
[IKEA](https://www.youtube.com/watch?v=sg-F_R-ZVNc), [Nexxiot](https://www.cncf.io/case-studies/nexxiot/) [Sky](https://www.youtube.com/watch?v=u-4naOMfs_w), [Utmost](https://www.cncf.io/case-studies/utmost/) 
and [many more](https://github.com/cilium/cilium/blob/main/USERS.md). It is also the default CNI for managed Kubernetes on GKE, Anthos, EKS-A, AKS, and other cloud providers.

The project maintainers believe that Cilium meets the requirements for Graduation status. 

## Graduation State Criteria

### Have committers from at least two organizations.

Cilium currently has [committers from](https://github.com/cilium/cilium/blob/main/MAINTAINERS.md) AMD, Datadog, Deepfence, Google, Independent, Isovalent, Palantir, and Red Hat.

### Have achieved and maintained an [OpenSSF Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

Cilium has achieved an [OpenSSF Best Practices Badge](https://bestpractices.coreinfrastructure.org/de/projects/1269).

### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

Cilium received independent security and fuzzing audits through CNCF from OSTIF. You can find the blog [here](https://www.cncf.io/blog/2023/02/13/a-well-secured-project-cilium-security-audits-2022-published/) and [each](https://github.com/cilium/cilium.io/blob/main/Security-Reports/CiliumSecurityAudit2022.pdf) [report](https://github.com/cilium/cilium.io/blob/main/Security-Reports/CiliumFuzzingAudit2022.pdf) respectively. All findings have been remediated. Finally, Cilium had an [LFX mentee](https://github.com/cncf/mentoring/tree/main/lfx-mentorship/2022/03-Sept-Nov#cilium) increase the security of the release process including signed SBOMs and signed release artifacts.

### Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

Project governance including granting and revoking rights, expectations, and voting are outlined in [governance documents](https://docs.cilium.io/en/latest/community/governance/). All committers to the project can be found in [MAINTAINERS.md](https://github.com/cilium/cilium/blob/main/MAINTAINERS.md).

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

Committers may be removed after 6 months of inactivity and are removed based on the [process outlined in the docs](https://docs.cilium.io/en/v1.12/community/governance/commit_access/#revoking-commit-access). Those removed will be changed to [emeritus in the MAINTAINERS.md](https://github.com/cilium/cilium/blob/main/MAINTAINERS.md#cilium--hubble-emeritus-committers)

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.

See [USERS.md](https://github.com/cilium/cilium/blob/main/USERS.md) and the [Cilium website](https://cilium.io/adopters/).

## Incubation Details

Cilium was accepted as an [Incubation level project in October 2021](https://github.com/cncf/toc/pull/637). There are no outstanding issues raised during the incubation due diligence that remain to be addressed.
