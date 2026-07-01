# Foundation

Infrastructure lifecycle management relies on foundational concepts, defining how
resources are provisioned, updated, and decommissioned. While some of the concepts
and tools focus on a declarative approach, others follow an imperative or mixed
approach. Similarly, some of them follow an approach to continuously reconcile
their state, while others are executed on demand.

In the following sections, some of these patterns, as well as their benefits and
potential drawbacks, are covered.

## Cloud Provider / Infrastructure APIs
The main prerequisite for Cloud Native Infrastructure Lifecycle Management is infrastructure
Application Programming Interfaces (APIs) that provide a way to interact with the underlying
infrastructure provider. These APIs can be provided at a certain granularity: while Virtualization
Providers mainly support the core compute layer (compute, storage, network), Cloud Providers often
expose the configuration of their services via their APIs.

For many of the tools mentioned in this brief, SDKs and/or providers are authored either by the vendors
themselves or by their community, which makes it easier for end-users to interact with them.

## Declarative vs. Imperative Configuration

While these APIs, SDKs, and providers can be used directly in, e.g., HTTP requests, there are many things to take care of when delivering infrastructure. As an example,
backend services for load balancers are often needed before a frontend can be deployed. Similarly, service accounts or IAM roles might be needed before compute resources can be created. Therefore, over
time, some tools emerged that cover all of these infrastructure specifics and "hide" them from their users. They can be grouped into two major groups: those with a declarative and those with
an imperative configuration.

In the Cloud Native world and in the Kubernetes space, declarative configuration is the
standard. When writing declarative configuration, a certain desired state gets
defined in a structured format (e.g. JSON, YAML, HCL). This approach requires some software
that is able to translate this definition into instructions that are executed in the right
order, and in a consistent way. The major benefit of this way of dealing with configuration
is that the developer usually doesn't have to care about the implementation of the configuration
itself, only about the desired state, while "someone else" (the author of the underlying software /
provider) takes care of the implementation of certain API specifics.

The alternative - imperative configuration - behaves a bit different. It can be seen as a sequence of 
commands configuring the related infrastructure. While this way of configuration is more transparent
and controllable, the end-user has to implement control mechanisms and failure-handling. In contrast
to declarative configuration, some programming constructs - such as loops and conditionals - can be
implemented in an easier way.

## On-Demand vs. Continuously Reconciled

These are two paradigms in lifecycle management that are fundamental opposites. On the one side, you have changes that are triggered by a person or external system, the other is where the system itself can reconcile any changes by observing a source. The first can be done through a portal, CLI, or scripts, while the other requires that you install and maintain a set of agents whose behavior is defined by configuration or policy.

When you manage infrastructure with on-demand solutions like scripts, either manually run or through an automation pipeline, one of the benefits is that you get immediate feedback. With a reconciliation process, you rely on eventual consistency, which requires all changes to happen at the source. The GitOps principles call this the “desired state” and it should always be what the system tries to adhere to. This means that a continuously reconciled system is self-healing, while any on-demand model has no such guarantee.

On-demand operations can also be more complex than a continuously reconciled one. Managing imperative pipelines and scripts can quickly grow out of proportion, putting cognitive load on the operator. The agents, Kubernetes controllers or operators, are usually singletons used for specific parts of the operation. This can result in a more stable architecture, with fewer variables and less bespoke solutions.

You might think of them in different terms like “push versus pull” or “pipelines” versus controllers. Neither is wrong, and they both have their benefits and disadvantages. You can also opt for a hybrid approach, where certain parts of the system are reconciled continuously, while others are not.

## DSL vs. Programming Language

## Stateful vs. Stateless

## Trade-offs

## Pattern Maturity

