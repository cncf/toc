# in-toto Graduation Proposal

in-toto, an open-source project that joined CNCF as a sandbox project in August 2019, and incubation in March 2022. 
Since then, in-toto has experienced a remarkable degree of adoption within various ecosystems and usecases. These include 
cases such as GitHub's, Gitlab's and Tekton among others. Due to this, we are confident that in-toto is ready to graduate.

## Introduction

[in-toto](https://in-toto.io) is a framework for software supply chain security. It offers a mechanism to capture and evaluate security properties as software artifacts are created and distributed. 
Through in-toto's layered architecture, information such as SBOMs, SLSA attestations and OpenVEX statements can be seamlessly collected, cryptographically signed, collected and evaluated.

in-toto protects the software supply chain as a whole. In order to achieve this, in-toto provides a series of mechanisms to define:

- Who can carry out each step  (step authentication)
- How the artifacts between each step interconnect with each other (artifact flow integrity)
- What steps are to be carried in the supply chain (supply chain integrity)

This way, in-toto can allow project owners to define (and update) the topology of the supply chain in a file called a software supply chain layout (or just layout for short). 
In addition to this, in-toto provides a way for actors (or functionaries) in the supply chain to provide cryptographically-signed attestations that can be used to verify that all steps within the supply chain were carried out according to the specification.

## Graduation State Criteria
#### Have committers from at least two organizations
1. Google
2. Kusari
3. NYU
4. Verizon
5. Intel
6. Purdue University
7. TestifySec

#### Have achieved and maintained a Core Infrastructure Initiative [Best Practices Badge](https://bestpractices.coreinfrastructure.org/projects/1523?criteria_level=2)
* in-toto‘s holds a gold OpenSSF best practices badge

#### Have completed an independent and third party security audit with results published of similar scope and quality as this example which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation
* The security audit has been completed [earlier this year](https://in-toto.io/security-audit-23/)

#### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers

The project’s [GOVERNANCE.md](https://github.com/in-toto/community/blob/main/GOVERNANCE.md) and [contributor instructions](https://github.com/in-toto/community/blob/main/CONTRIBUTING.md) cover the committer lifecycle, onboarding, offboarding, and emeritus criteria.  Any participant may commit, so long as their code is approved by a project maintainer for the implementation for that codebase.  The current [maintainers](https://github.com/in-toto/in-toto/blob/develop/MAINTAINERS.txt) are also listed in the repository.

The in-toto specification has a separate process by which changes and additions are proposed and vetted.  This is through the ITE (in-toto enhancement) [process](https://github.com/in-toto/ITE), which involves a public proposal of a specification change which is discussed by the [community](https://github.com/in-toto/ITE/blob/master/ITE/1/README.adoc).


#### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence

As above, the project’s [GOVERNANCE.md](https://github.com/in-toto/community/blob/main/GOVERNANCE.md) covers the criteria for maintainers including the onboarding, offboarding, and emeritus criteria.   These are audited at an annual cadence by the in-toto steering committee, as is described in the project’s GOVERNANCE document.

While the maintainers of both the in-toto specification and its implementations work together for the health of the project, for the purposes of CNCF interactions, the in-toto [maintainers](https://github.com/in-toto/in-toto/blob/develop/MAINTAINERS.txt) are the ones who will interact with the CNCF on behalf of the project.  The current [maintainers](https://github.com/in-toto/in-toto/blob/develop/MAINTAINERS.txt) are listed in the repository. 

#### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to FAQs for guidelines on identifying adopters

There is a public “friends list” of project adopters [here](https://github.com/in-toto/friends).


## Incubation Details
#### Link to Incubation Due Diligence(DD) Document
* [Incubation proposal](https://github.com/cncf/toc/pull/393)
* [in-toto Due Diligence](https://docs.google.com/document/d/1ZBiErp4sPRhNV5lFjb5pByHpTpuozuFTJA9Zv-SDHpE/edit#heading=h.xty7ca1yvruj)

#### Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

- Verify in-toto's supply chain with in-toto in-toto/issue#278

See [#444](https://github.com/in-toto/in-toto/pull/444)

- Improve introductory documentation to clearly communicate security scope docs/issue#15

We have worked on this process in our [community
repository](https://github.com/in-toto/community).  We also have put in a
[request](https://cncfservicedesk.atlassian.net/servicedesk/customer/portal/1/CNCFSD-1517)
to the CNCF documentation experts to give us an outsider’s perspective.

- Additional integrations, examples and/or documented case studies (such as: in-toto/issue#284, roadmap#3)

This documentation was added to our project adopters page.

- Consider encoding best practices in default implementation (such as issue#287)

We’ve worked to address issues (such as issue#287) to reduce user confusion.  We have also worked with communities that are building policies on in-toto (such as [SLSA](https://slsa.dev/)) to ensure that they have integrated it in a secure manner.

- Proceed with CII silver badge & roadmap. 

We have since obtained [silver and gold OpenSSF Best Practices Badges](https://bestpractices.coreinfrastructure.org/en/projects/1523?criteria_level=2)

- Formal security audit: no blocking issues for a formal code audit

See above 

- Additional organizations contributing to as core members of the development team, recommend addressing documentation issues above in advance of CNCF promotion

See list of contributors above

