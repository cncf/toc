== CNCF Graduation Criteria v1.3

Every CNCF project has an associated maturity level. Proposed CNCF projects should state their preferred maturity level. A two-thirds supermajority is required for a project to be accepted as incubating or graduated. If there is not a supermajority of votes to enter as a graduated project, then any graduated votes are recounted as votes to enter as an incubating project. If there is not a supermajority of votes to enter as an incubating project, then any graduated or incubating votes are recounted as sponsorship to enter as an sandbox project. If there is not enough sponsorship to enter as an sandbox stage project, the project is rejected. This voting process is called fallback voting.

The criteria for each stage of maturity are described below, and there is a separate https://github.com/cncf/toc/blob/main/process/project_proposals.adoc[Project Proposal process].

Incubating and graduated projects have access to all resources listed at https://cncf.io/projects[https://cncf.io/projects] but if there is contention, more mature projects will generally have priority.

=== Sandbox Stage

To be accepted in the sandbox a project must

* Apply to join the sandbox using the https://docs.google.com/forms/d/1bJhG1MuM981uQXcnBMv4Mj9yfV5_q5Kwk3qhBCLa_5A/edit[form]
* Adopt the CNCF https://github.com/cncf/foundation/blob/master/code-of-conduct.md[Code of Conduct]
* Adhere to CNCF https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy[IP Policy] (including trademark transferred)
* List their sandbox status prominently on website/readme

See the https://github.com/cncf/toc/blob/main/process/sandbox.md[CNCF Sandbox Guidelines v1.0] for the detailed process.

=== Incubating Stage

*Note: The incubation level is the point at which we expect to perform full https://github.com/cncf/toc/blob/main/process/due-diligence-guidelines.md[due diligence] on projects.*

To be accepted to incubating stage, a project must meet the sandbox stage requirements plus:

 * Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope. For the definition of an end user, see https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-end-user

 * Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.
 * Demonstrate a substantial ongoing flow of commits and merged contributions.
 * Since these metrics can vary significantly depending on the type, scope and size of a project, the TOC has final judgement over the level of activity that is adequate to meet these criteria
 * A clear versioning scheme.
 * Clearly documented security processes explaining how to report security issues to the project, and describing how the project provides updated releases or patches to resolve security vulnerabilities 
 * Specifications must have at least one public reference implementation.

=== Graduation Stage

To graduate from sandbox or incubating status, or for a new project to join as a graduated project, a project must meet the incubating stage criteria plus:

 * Have committers from at least two organizations.
 * Have achieved and maintained a Core Infrastructure Initiative https://bestpractices.coreinfrastructure.org/[Best Practices Badge].
 * Have completed an independent and third party security audit with results published of similar scope and quality as the following example (including critical vulnerabilities addressed): https://github.com/envoyproxy/envoy#security-audit and all critical vulnerabilities need to be addressed before graduation.
 * Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.
 * Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.
 * Receive a supermajority vote from the TOC to move to graduation stage. Projects can attempt to move directly from sandbox to graduation, if they can demonstrate sufficient maturity. Projects can remain in an incubating state indefinitely, but they are normally expected to graduate within two years.
