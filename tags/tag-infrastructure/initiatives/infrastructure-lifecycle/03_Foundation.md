# Foundation

Infrastructure lifecycle management relies on foundational concepts, defining how
resources are provisioned, updated, and decommissioned. While some of the concepts
and tools focus on a declarative approach, others follow an imperative or mixed
approach. Similarly, some of them follow an approach to continuously reconcile
their state, while others are executed on demand.

In the following sections, some of these patterns, as well as their benefits and
potential drawbacks, are covered.

## Cloud Provider and Infrastructure APIs

An infrastructure API is the programmable interface through which resources are
provisioned, changed, and removed. It is the prerequisite for every pattern in
this brief, because it replaces manual consoles and ticket queues with an
interface that tools can call directly and repeatably.

These APIs exist at several layers. At the bottom is physical hardware. Above it,
virtualization exposes core compute, storage, and network. Infrastructure as a
service adds those same primitives together with identity, and managed services
expose higher level products such as databases and queues. On top of all of this,
cluster level APIs model whole Kubernetes clusters as resources. A distinction
that runs through the rest of this brief is whether a resource lives inside a
cluster or outside it, and the layer an API sits at usually decides which.

The physical layer is the exception, because hardware has no native management API
in the way a cloud service does. Provisioning tools reach it through out of band
interfaces such as Redfish and IPMI exposed by the server's management controller,
together with network boot, and build an API-like surface on top of those.

Tools do not call these APIs directly. Each one integrates through a provider, a
plugin that maps the tool's model onto a specific API. The providers and
controllers described later all follow this shape, which is why support for a new
platform usually means writing a provider rather than changing the tool.

Several properties of these APIs shape the trade-offs discussed later. Operations
are often asynchronous and only eventually consistent, so a resource reported as
created may not yet be usable. Calls are rate limited and subject to quotas, and
credentials carry a defined scope. These are the reasons the tools built on top
need retries, backoff, and a way to wait for a resource to settle.

## Declarative vs. Imperative Configuration

The two configuration models differ in what the user describes. Declarative
configuration states the desired end result and leaves the tool to work out the
steps that reach it. Imperative configuration states the steps themselves, as an
ordered sequence of commands. The difference is the model, not the file format: a
declarative desired state is often written in JSON, YAML, or HCL, but the format
is incidental to the idea.

The models exist because calling infrastructure APIs by hand is awkward. Resources
depend on each other, so a load balancer backend has to exist before the frontend
that uses it, and an identity role before the compute that assumes it. A tool that
understands these relationships removes that burden from the user.

Declarative configuration is the default in the cloud native world. It relies on
an engine that compares the desired state against what currently exists and applies
only the difference, usually as a plan step followed by an apply step. Two
properties matter here and recur throughout the brief. An operation is idempotent
when applying it repeatedly leaves the same result, so a run that changes nothing
is safe. Convergence is the process of repeatedly moving the actual state toward
the declared one. Because the user describes the destination rather than the
route, the same configuration can be applied again without tracking what has
already been done.

Imperative configuration executes commands directly and in the order given. It is
more transparent and gives finer control, and it expresses conditionals and loops
naturally, because it runs in a normal flow of execution. The cost is that the
user owns ordering, failure handling, and retries, and has to account for the
current state of the system before acting.

Declarative configuration has its own costs. Logic that is simple to write
imperatively, such as a conditional or a loop, is often awkward to express. Because
the engine acts implicitly, a failure can be harder to trace back to a cause, and
the abstraction that hides API detail also hides where something went wrong.

The two are not exclusive. The authoring language is independent of the model, so a
desired state can be produced by a general purpose program rather than written by
hand, a point the next section develops. Many declarative tools also provide escape
hatches that run imperative steps for cases the declarative model cannot express,
and physical provisioning commonly pairs a declarative description of the target
host with an imperative workflow that powers it on, boots it, and writes the image.

## On-Demand vs. Continuously Reconciled

The two execution models differ in what causes a change to happen. In the on demand
model, a change runs when a person or an external system triggers it, through a
portal, a CLI, or a pipeline. In the continuously reconciled model, a running
component watches a declared source of truth and acts on its own to keep the live
system matching it.

Reconciliation works through a control loop. The component observes the current
state, compares it against the desired state, acts to close any difference, and
repeats. This loop is what makes a reconciled system self healing: when the live
state drifts from the declared one, whether through a manual change or a failure,
the next pass corrects it without anyone intervening.

The models trade immediacy against durability. On demand execution gives immediate
feedback, since the result of a run is visible when it finishes, and with the right
tooling a run is idempotent and flexible. Its weakness is that nothing watches the
system afterward, so it drifts as reality diverges from the last applied intent.
Reconciliation offers eventual consistency against a persistent desired state and
holds the system there over time, at the cost of immediate feedback.

Reconciliation also has a standing cost. The component that runs the loop is itself
infrastructure with its own lifecycle to operate, upgrade, and secure, and it has
to be bootstrapped by some other means first. There is no built in equivalent of a
plan step that previews what a loop will do before it does it. And because the loop
asserts ownership of the fields it manages, a change made directly against the live
system is reverted on the next pass, and two components managing the same field will
overwrite each other, which is why reconciled systems provide a way to pause a
resource during an incident.

The reconciliation model also assumes that resources can be created and destroyed
through an API cheaply and quickly. That assumption holds for cloud resources and is
weaker for physical hosts, where provisioning takes minutes to boot, image, and
inspect, deletion means wiping and powering down rather than removing a record, and
the pool of hardware is finite and discovered rather than requested on demand.

The choice can also follow compliance rather than preference. Where a change must
pass review and approval, an on demand model with an explicit trigger fits the
control. Where a rule must hold continuously across every resource, such as a
required access control, a reconciling component that keeps asserting it fits
better. The distinction is sometimes framed as push versus pull, or pipelines
versus controllers, and the two can be combined, with some parts of a system
reconciled continuously and others changed on demand.

_The remaining sections below (DSL vs. Programming Language, Stateful vs.
Stateless, Trade-offs, Pattern Maturity) are still in progress, drafted
collaboratively outside this repo. Tracked in
[#1631](https://github.com/cncf/toc/issues/1631)._

## DSL vs. Programming Language

## Stateful vs. Stateless

## Trade-offs

## Pattern Maturity

