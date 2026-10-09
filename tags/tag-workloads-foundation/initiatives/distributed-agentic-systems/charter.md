# Distributed Agentic Systems Initiative Charter and Terms of Reference

Creation issue: [#1746](https://github.com/cncf/toc/issues/1746)
Parent: [TAG Workloads Foundation](../../README.md)
Adopted at the working session of September 24, 2026; amended October 8, 2026.

## Mission

Establish cloud native principles and reference patterns for running distributed AI agent systems on Kubernetes, by operationalizing and extending the CNCF Cloud Native Agentic Standards. The initiative provides architectural guidance and identifies where standards are needed. It does not build new agent frameworks or runtimes.

## Scope

### In scope

Five workstreams, aligned one to one with the domains of the Cloud Native Agentic Standards so that equivalence between the standards and this document is easy to maintain:

1. Protocol and communication interoperability: Model Context Protocol (MCP), Agent-to-Agent (A2A), AP2, agent discovery and naming, dynamic registration, schema validation.
2. Gateway and data plane: agent-to-agent and agent-to-tool routing, Gateway API Inference Extensions, session-aware fan-out, bidirectional server-sent events, JSON-RPC streaming and multiplexing, proxy extensions.
3. Runtime and container abstraction: Pod, sidecar and CRD modeling for agents, container best practices, microVM and confidential container isolation, horizontal autoscaling, Kueue and dynamic resource allocation, task recovery.
4. State, memory and persistence: short-term conversational state, long-term memory, vector storage operators, distributed state replication, garbage collection across dynamic agents.
5. Governance, observability, security and policy: OpenTelemetry GenAI semantic conventions and spans, metrics, logs and traces, LLM cost, token and latency tracking, policy and compliance as code, verifiable per-call evidence.

Single-cluster deployment is the baseline. Multi-cluster and geographically distributed deployment (cross-cluster discovery, state and routing) is treated as a first-class concern from the start.

### Out of scope

- Building new agent frameworks or runtimes.
- Deep dives into any single framework's implementation.
- Vendor-specific product guidance.
- Restating content the Cloud Native Agentic Standards already cover; this document references the standards and extends them.

## Deliverables

| # | Deliverable | Form |
|---|---|---|
| D1 | The document in this directory | Versioned, living; tagged releases in CHANGELOG.md |
| D2 | Reference architecture and pattern catalogue | Chapter 07 |
| D3 | Standards proposals (for example MCP for clusters, Agent CRD schema) | Chapter 08, routed upstream |
| D4 | Gap analysis | Chapter 08 |
| D5 | Alignment with adjacent efforts | Tracked in README.md |
| D6 | Observability and conformance guidance for agents | Chapter 06 |

### Deliverable direction

At the September 24, 2026 session the group chose to operationalize and extend the existing standards rather than write a new standalone whitepaper, and to maintain the result as a GitHub-based living document updated through pull requests. Tagged release versions give each milestone a citable form while the work continues.

## Roles

| Role | Who |
|---|---|
| Coordination and delivery lead | Pavan Madduri (@pmady) |
| Technical direction | Kante Yin (@kerthcet) |
| Initiative author and TOC liaison | Vincent Caldeira (@caldeirav) |
| CNCF program support | Riaan Kleinhans (@riaankleinhans) |
| Stream owners | Listed in README.md and confirmed in meeting notes |

## Terms of reference

- Decisions are made by lazy consensus. Proposals are posted in the Slack channel or as pull requests; silence for 72 hours after a call for objections is consent.
- Stream owners volunteer. Work that has not moved by the following session can be reassigned so the document keeps moving.
- Sessions are held every other Thursday at 9:00 AM US Central and are recorded and published.
- Drafting happens in this directory through pull requests; discussion happens in `#initiative-distributed-agentic-systems`.
- Content is vendor neutral. Projects and products may be cited as implementations of a pattern, never as the basis for one, and no section may be built on a single project.
- Changes to this charter follow the same lazy consensus process and are recorded in CHANGELOG.md.

## Adjacent efforts

| Effort | Organization | Relationship |
|---|---|---|
| Cloud Native Agentic Standards | CNCF AI TCG | Base document this initiative extends |
| AI Gateway and Agentic Network work | Kubernetes, CNCF | Gateway and data plane alignment |
| AAIF | Agentic AI Foundation | Agentic standards alignment |
| SuperBlueprint | LF Edge | Overlapping distributed agent work at the edge |
| Cloud native agentic blueprint for telco | ETSI | Industry blueprint informed by this work |
| AI Safety Best Practices | LF AI & Data | Guardrails that map to the governance stream |
