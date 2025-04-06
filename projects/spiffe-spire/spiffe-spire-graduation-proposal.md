
# SPIFFE/SPIRE Graduation Proposal

  

The SPIFFE and SPIRE projects were started with an ambitious vision: to automatically provide every workload with an identity that is cryptographically verifiable, automatically rotated, interoperable and portable across a wide variety of platforms and providers. The goal is an interoperable workload identity format (SPIFFE), and supporting tooling to issue those identities (SPIRE), that allows components in a distributed system to ensure the authenticity, integrity and confidentiality of messages passed between them. To accomplish it effectively, the supporting toolchain must be able to deliver identities to workloads across a range of hosting platforms (eg. private estate, public cloud) and runtime orchestration tools for generalized classes of identities (canonical, attributed, or delegated) for varying workload form factors (eg. bare metal machines, virtual machines, containers).

  

At a high level, SPIFFE defines the APIs and documents necessary to accomplish this vision, and SPIRE provides a vendor and platform agnostic implementation of the SPIFFE APIs. The platform agnostic nature of SPIFFE coupled with the flexibility of SPIRE has proven to be a potent combination. The SPIFFE specification has been adopted by several popular open-source projects beyond SPIRE, including Istio, Consul and Kuma. SPIRE has been successfully deployed in widespread production in several organizations (including Square, Uber and Bytedance). The production scalability and robustness of SPIRE has also been proven, with the largest known production deployment exceeding 1,000,000 nodes.

  

The growth of the SPIFFE and SPIRE projects, with support from the CNCF as it moved through the Sandbox and Incubation phases, is evidenced in many dimensions ranging from maturity and robustness to improved user experience and broader reach. Some examples of project growth and improvements over the last one year include:

  

-   12 SPIRE releases, including version 1.0
    
-   'SPIFFE Federation' released and supported natively in SPIRE
    
-   Project Tornjak launched, a web-based UI for managing SPIRE deployments
    
-   SPIRE released support for serverless platforms via [SVID Store plugins](https://blog.spiffe.io/enabling-authenticated-communication-for-serverless-workloads-with-spire-d636bf2f7a91)
    
-   Developed a [maturity framework](https://github.com/spiffe/spiffe/blob/main/MATURITY.md#) for sub-projects
    
-   SPIRE grew its number of platform integrations through the introduction of a half dozen new plugins
    
-   Across Twitter, Github and Slack the community had 915 distinct active members from the period Feb 1, 2021 to Jan 31, 2022. The number of SPIFFE Slack members has grown by over 37%, now exceeding 1,300 people
    
-   Number of SPIRE contributors grown by ~47%
    
-   Number of SPIRE commits grown by nearly ~30%
    

  

SPIFFE and SPIRE governance has also significantly strengthened and diversified to become true vendor-neutral projects. SPIFFE and SPIRE’s collective seven core maintainers represent six distinct organizations. Similarly the project’s steering committee’s 5 members represent four distinct organizations. Over the 12 months from Feb 1, 2021 to Jan 31, 2022 alone, the projects received pull requests from employees of 41 organizations over 4 continents.

  

The SPIFFE community holds regular meetups and community days, and runs an annual co-located zero day event at KubeCon North America (Production Identity Day). The events draw hundreds of practitioners and provide a space for focused conversation and the sharing of new ideas and best practices.

  
  

## Graduation State Criteria

  

### Have committers from at least two organizations.

  

SPIRE has 5 committers from 3 different organizations. SPIFFE has 4 committers from 3 different organizations. The steering committee has 5 members from 4 organizations.

  

### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

  

The project has achieved and maintained CII Best Practices badge since 2018.

  

### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

  

Our [audit](https://github.com/spiffe/spire/blob/main/doc/cure53-report.pdf) was conducted by Cure53 in the first few months of 2021. The auditors found no critical errors, determined that all of the first lines of defense that we had in place were sound and affirmed that the two projects “had been created with security in mind” and that their maintainers have been “on the right track regarding security.”

  

They did find two vulnerabilities: one in the legacy API that was being deprecated at the time, which they labeled high severity; and then a medium severity vulnerability around validation logic and administrative APIs. We promptly issued a release to address both vulnerabilities. We worked very hard to offer changes that were highly targeted, super-safe and as backwards-compatible as possible. For the medium severity issue we also made it so that people could get the high severity fix but have a method to turn off the patch for the medium security fix while they adjusted their deployment, if they needed to. We ported back five versions of the software.

  

### Explicitly define a Project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

  

[GOVERNANCE.md](https://github.com/spiffe/spiffe/blob/main/GOVERNANCE.md)

[MAINTAINERS.md](https://github.com/spiffe/spire/blob/main/MAINTAINERS.md)

  

### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec.


For the SPIFFE specification, a list of “Who uses SPIFFE” SPIFFE is displayed in the project website spiffe.io . The list showcases the variety of projects that implement the spec to both issue and consume SPIFFE IDs.


As for the SPIRE software implementation, the [ADOPTERS.md](https://github.com/spiffe/spire/blob/main/ADOPTERS.md) file in the GitHub repository for SPIRE, catalogs a number of implementors and end user organizations who have spoken publicly of their success stories with SPIRE.

## Incubation Details

### Due diligence document

  

[Incubation Due Diligence Document](https://docs.google.com/document/d/1tkN9YgBSLEUszOflWPHO72qedOaUb3iHfAye45dKJT8/edit#heading=h.378jkvcve1nq)

  

### Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

  

The CNCF Security TAG had made the following recommendations which we have followed up on and completed:

  

* Made public the threat modeling materials from the security assessment.

* Expanded the security response teams to include all maintainers.

  

The CNCF made the following recommendations:

  

* Conduct a formal security review/audit for the project as it is critical to security of services that use it. We completed the audit in the first months of 2021.

 * Provide an avenue for education of SPIFFE/SPIRE and advanced SPIFFE/SPIRE topology use cases for end users. In addition to producing Production Identity Day as a colocated event, we published a [book on SPIFFE](https://spiffe.io/book/) and also created [interactive learning labs](https://www.katacoda.com/spiffe) in katacoda.
