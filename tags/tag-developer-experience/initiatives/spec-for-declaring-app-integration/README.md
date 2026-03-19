# Specification for declaring application integration dependencies

https://github.com/cncf/toc/issues/1797

Spec that lists integration dependencies needed for an application to run successfully, e.g. DBs, message bus, etc.

## Initiative description

One of the common things we hear in software development is that if we package an application a certain way, it can run anywhere. While that may be true in terms of executing the command that runs it, it doesn’t guarantee that the application won’t exit 1 within a few milliseconds when it can’t connect to a database or cache, or start throwing 500s when it tries to call to a peer service that hasn’t been deployed. Yes, it can run, but can it run successfully?

This initiative would set out to create a project-agnostic specification that for declaring a service’s dependencies, so that it could indicate to an application deployment team what other services must be readily available for the application to run successful - e.g. APIs, databases, caches, blob stores, message buses, etc. This is not meant to replace conversations between application engineers and production support. Rather, it’s meant to give them a standardized reference that they can use to enable their discussion.

The terms project agnostic specification is used to build this as part of a solution in the past, e.g. Radius. Since the Radius approach packaged the spec in with the solution, the spec has only evolved to the degree that the solution can support it, most notably with a strong focus on Azure services. Making the spec project-agnostic would create a solution-independent language that can be implemented however necessary by conforming projects, including Radius. In fact, this effort would likely look to Radius, Porter, and other projects like them for inspiration. But the benefit of a spec is that it could also be leveraged for other use cases, such as guaranteeing abstraction coverage in Dapr, or for integration mocking in Microcks. Furthermore, it can support security efforts that are focused on expected application behavior, such as Software Bill of Behavior (SBOB) initiatives.

## Deliverable(s) or exit criteria

* first deliverable is research and report on existing specs and previous attempts that have failed (and maybe could be resurrected) to better understand the problem space
  * the outcome of this first step will dictate the direction of subsequent steps, but the following should still apply regardless
  * this is the first potential exit point if, for whatever reason, the conclusion in the report is that this is too complex and we recommend that a spec not be created
* after the report has been delivered to the TAG, identify (and potentially extend) and existing spec, or as a last resort create a specification that indicates what services must be available to an application to ensure that it can run successfully
  * First release will be scoped to a specific integration set, e.g. external APIs, or required data stores (DB, cache)
    * TBD based on group discussion
  * The specification must exist in a format that can be consumed programmatically by spec consumers, e.g. JSON, YAML, etc.
* A plan and project description for automatic generation of dependency declarations that adhere to the spec during application development
  * to be clear, this is the description or recommendation of a project, not the project itself
* Guidance for consumers of dependency declarations for how to consume the spec and interpret it for downstream use
  * likely in the form of a short white paper
