## CNCF Graduation Criteria v1.3

Every CNCF project has an associated maturity level. Proposed CNCF projects should state their preferred maturity level. A two-thirds supermajority is required for a project to be accepted as incubating or graduated. If there is not a supermajority of votes to enter as a graduated project, then any graduated votes are recounted as votes to enter as an incubating project. If there is not a supermajority of votes to enter as an incubating project, then any graduated or incubating votes are recounted as sponsorship to enter as a sandbox project. If there is not enough sponsorship to enter as a sandbox stage project, the project is rejected. This voting process is called fallback voting.

The criteria for each stage of maturity are described below, and there is a separate [Project Proposal process](https://github.com/cncf/toc/blob/main/process/project_proposals.md).

Incubating and graduated projects have access to all resources listed at [cncf.io/projects](https://cncf.io/projects) but if there is contention, more mature projects will generally have priority.

### Sandbox Stage

To be accepted in the sandbox a project must

* Apply to join the sandbox using the [form](https://github.com/cncf/sandbox/issues/new/choose)
* Adopt the CNCF [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
* Adhere to CNCF [IP Policy](https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy) (including trademark transferred)
* List their sandbox status prominently on website/readme

See the [CNCF Sandbox Guidelines v1.0](https://github.com/cncf/toc/blob/main/process/sandbox.md) for the detailed process.

### Incubating Stage

*Note: The incubation level is the point at which we expect to perform full [due diligence](https://github.com/cncf/toc/blob/main/process/due-diligence-guidelines.md) on projects.*

To be accepted to incubating stage, a project must meet the sandbox stage requirements plus:

 * Document that it is being used successfully in production by at least three independent direct adopters which, in the TOC’s judgement, are of adequate quality and scope. For the definition of an adopter, see https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter.

 * Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.
 * Demonstrate a substantial ongoing flow of commits and merged contributions.
 * Since these metrics can vary significantly depending on the type, scope and size of a project, the TOC has final judgement over the level of activity that is adequate to meet these criteria
 * A clear versioning scheme.
 * Clearly documented security processes explaining how to report security issues to the project, and describing how the project provides updated releases or patches to resolve security vulnerabilities 
 * Specifications must have at least one public reference implementation.

### Graduation Stage

To graduate from sandbox or incubating status, or for a new project to join as a graduated project, a project must meet the incubating stage criteria plus:

 * Have committers from at least two organizations.
 * Have achieved and maintained a Core Infrastructure Initiative [Best Practices Badge](https://bestpractices.coreinfrastructure.org/).
 * Have completed an independent and third party security audit with results published of similar scope and quality as the following example (including critical vulnerabilities addressed): https://github.com/envoyproxy/envoy#security-audit and all critical vulnerabilities need to be addressed before graduation.
 * Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.
 * Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should be preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.
 * Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.
 * Receive a supermajority vote from the TOC to move to graduation stage. Projects can attempt to move directly from sandbox to graduation, if they can demonstrate sufficient maturity. Projects can remain in an incubating state indefinitely, but they are normally expected to graduate within two years.
