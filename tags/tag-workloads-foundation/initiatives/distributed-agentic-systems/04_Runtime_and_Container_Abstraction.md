# 4. Runtime and container abstraction

Status: outline. Owner: Srikanth Rao.

Help platform teams choose workload, isolation, scaling, and lifecycle boundaries for agents on Kubernetes, and distinguish process recovery from task recovery. Projects named below are reference points for pattern comparison, not recommendations; the chapter is organized around capabilities, not projects.

## Modeling agents as Kubernetes workloads

- Why core workload models strain: an agent instance is a long-running, stateful, singleton execution with a stable identity and probabilistic, dynamic work, which fits neither the stateless Deployment model nor the numbered StatefulSet model
- Two-layer pattern: separate the agent orchestration layer from the secure execution layer
  - Orchestration layer: model routing, prompt and tool configuration (MCP servers, CLIs, APIs), conversation state, workflows, and approval gates
  - Execution layer: isolated environments for the risky part of the work — running model-chosen code, scripts, and browsers; file generation; persistent workspaces
  - Request flow: the orchestration layer receives the request and drives the conversation; when a task requires executing code or performing risky actions, it launches or attaches an execution environment, the work runs inside it, and results return to the orchestration layer to summarize back to the user
  - Why not run the agent directly in a plain Pod: an agent co-located with its tools shares the Pod's filesystem, process space, and identity, so one session's blast radius covers the whole workload; multi-tenancy is harder, and executing generated code inside the agent's own container is riskier
  - What a separated execution layer adds per session: a dedicated runtime, stable identity, persistent storage, pause/resume lifecycle, and a kernel-level or microVM isolation boundary
  - Where the layering matters most: coding agents, browser-use agents, computer-use agents, autonomous remediation, and anything executing arbitrary user-generated code
  - Illustration: a coding agent asked to fix a failing test reads repository context, plans the change, and selects the affected test suite in the orchestration layer; it then spins up an isolated sandbox with a persistent workspace (cloned repository), applies the patch and runs the suite inside it, returns the diff and results to the orchestration layer, and the sandbox is suspended with its workspace state resumable for a follow-up session
- Baseline patterns with core Kubernetes objects:
  - Pod / Deployment: replicated, stateless agent runtimes; the platform owns restarts and scale
  - StatefulSet: ordered, stable sets of agent instances
  - Job / CronJob: one-shot and scheduled runs (evaluations, batch work)
  - Native sidecar containers: tool execution, credential handling, and telemetry attached to an agent Pod
  - CRD + controller: the agent as a first-class cluster object with its own status, lifecycle, and conditions
- Decision axes: who owns the lifecycle (platform, framework controller, or application), what identity and state the instance needs, how quickly new instances must become available, and whether the agent executes model-chosen code outside its own Pod
- Reference points for the two layers, to compare rather than adopt:
  - Agent Sandbox project (Kubernetes SIG Apps) — execution layer: a Sandbox CRD for a single stateful, singleton workload with stable identity, persistent storage, and pause/resume; extensions add SandboxTemplate, SandboxClaim, and SandboxWarmPool for pre-warmed, low-latency allocation; isolation is delegated to sandbox runtimes through RuntimeClass
  - kagent (LF project) — orchestration layer: declarative AgentTemplate (model, prompt, tools), Harness (execution policy and placement), and Agent (pairs the two) compiled by a controller, with Sessions for running conversations; the 0.x model materialized agents as Deployments, while 1.0 runs Sessions on Agent Substrate, the execution backend kagent promotes, which pre-warms worker capacity and suspends idle sandboxes to snapshots
  - Overlap: Agent Substrate and Agent Sandbox pursue the same goal — secure, isolated agent execution on Kubernetes — so composing a orchestration layer with an execution layer requires comparing the options rather than assuming they interoperate
- Open questions: where does a standard agent workload CRD belong (Kubernetes, a SIG, or the framework layer), what is the minimal contract (identity, state, lifecycle, pause/resume) that keeps agent workloads portable across implementations, and what is the interface between a orchestration layer and an execution layer that keeps the composition portable

## Container best practices

- Apply the Cloud Native Agentic Standards baseline: distroless or minimal images, non-root, read-only root filesystems where feasible
- Image signing and admission validation for agent and tool images
- Keep credentials out of agent processes where possible; egress-gateway credential injection as one pattern to investigate

## Isolation

- Trust problem: agents execute model-chosen commands and process untrusted content, so the default Pod trust model is insufficient
- Options and trade-offs: kernel-level isolation (gVisor), microVMs (Kata, Firecracker), confidential containers and TEEs
- RuntimeClass as the mechanism that separates the isolation choice from the workload definition
- Per-instance identity and egress control: network policies, egress allowlists, per-agent credentials

## Scaling and scheduling

- Candidate signals: task queue depth, active sessions, model latency, token throughput, GPU utilization
- HPA (including scale-to-zero behavior) and event-driven scaling for bursty agent traffic
- Kueue for queued and batch agent work; DRA for accelerator allocation
- Pre-warming and warm pools for fast sandbox allocation, and their cost trade-offs

## Task recovery and lifecycle

- Distinguish process recovery (the platform restarts a container) from task recovery (the agent resumes meaningful progress)
- Restart, pause/resume (hibernation), and snapshot semantics; what survives a Pod replacement
- Draining and safe termination: in-flight tool calls, external side effects, and approval gates
- Boundaries: durable checkpoint semantics with the state stream, reconnect and replay with the gateway and protocol streams, identity and audit evidence with the governance stream

