# CNCF Distributed Agentic Systems Initiative

Creation issue: [#1746](https://github.com/cncf/toc/issues/1746)

## Description

Enterprise AI is moving from single-model inference to distributed systems of collaborating agents: long-running, probabilistic loops with dynamic tool calls, persistent memory, and streaming transport. Cloud native patterns designed for stateless, deterministic microservices do not cover these workloads.

This initiative operationalizes and extends the [CNCF Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) for Kubernetes. The standards define the what; this initiative delivers the how: reference architectures and Kubernetes patterns that platform teams can run. The output is a versioned, living document maintained in this directory through pull requests, with tagged releases recorded in [CHANGELOG.md](./CHANGELOG.md) so each milestone is citable while the work continues.

## Objective

- Publish a reference architecture and pattern catalogue for running distributed agent workloads on Kubernetes.
- Extend the Cloud Native Agentic Standards where cluster environments need more than the standards currently specify, for example agent discovery, session-aware data planes, and comparable cost telemetry.
- Produce a gap analysis and standards proposals (for example MCP for clusters, an Agent CRD schema, OpenTelemetry conformance for agents) and route them to the right upstream groups.
- Align with adjacent efforts rather than duplicate them: the Kubernetes AI Gateway and Agentic Network work, AAIF, LF Edge, ETSI, and the LF AI & Data AI safety work.

## Workstreams

| Stream | Scope | Owner |
|---|---|---|
| Protocol and communication interoperability | MCP, A2A, AP2, agent discovery and naming, schema validation | open |
| Gateway and data plane | Agent-to-agent and agent-to-tool routing, Gateway API Inference Extensions, session-aware fan-out, bidirectional SSE, JSON-RPC streaming | Akanksha Trehun |
| Runtime and container abstraction | Pod, sidecar and CRD modeling, container best practices, microVM and confidential containers, HPA, Kueue and DRA integration, task recovery | Srikanth Rao |
| State, memory and persistence | Conversational state, long-term memory, vector storage, distributed state replication, garbage collection | Victor Lu |
| Governance, observability, security and policy | OpenTelemetry GenAI conventions, MELT, LLM cost and token telemetry, policy and compliance as code, verifiable per-call evidence | Elankumaran Srinivasan |

Owners are confirmed at working sessions and recorded in the meeting notes. Interest in an open stream can be raised in the Slack channel.

## Logistics

The initiative is coordinated by Pavan Madduri (@pmady), with Kante Yin (@kerthcet) and Vincent Caldeira (@caldeirav), who authored the initiative, as co-leads. It meets every other Thursday at 9:00 AM US Central (14:00 UTC) on the [TAG Workloads Foundation calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/tag-workloads-foundation?view=list). Recordings are published on the [TAG Workloads Foundation YouTube channel](https://www.youtube.com/@CNCFTAGWorkloadsFoundation).

- Slack: `#initiative-distributed-agentic-systems` on the [CNCF Slack](https://slack.cncf.io)
- Meeting notes: [meeting-notes/](./meeting-notes/)
- Working draft and contributor register: [Google Doc](https://docs.google.com/document/d/1bRWy8vhke16mMhpPlEE_Dv71fTmkRWG-4Cy1hdTHUb0/edit)
- Charter and terms of reference: [charter.md](./charter.md)

## How we work

Decisions are made by lazy consensus. Owners volunteer for streams; work that has not moved by the following session can be reassigned. Contributions go in through pull requests to this directory and are reviewed by the approvers listed in [OWNERS](./OWNERS). Content is vendor neutral; the rules are in the terms of reference in [charter.md](./charter.md).

## Document structure

| File | Content |
|---|---|
| [01_Introduction.md](./01_Introduction.md) | Problem statement, audience, scope and non-goals |
| [02_Protocol_and_Communication_Interoperability.md](./02_Protocol_and_Communication_Interoperability.md) | Stream 1 |
| [03_Gateway_and_Data_Plane.md](./03_Gateway_and_Data_Plane.md) | Stream 2 |
| [04_Runtime_and_Container_Abstraction.md](./04_Runtime_and_Container_Abstraction.md) | Stream 3 |
| [05_State_Memory_and_Persistence.md](./05_State_Memory_and_Persistence.md) | Stream 4 |
| [06_Governance_Observability_Security_and_Policy.md](./06_Governance_Observability_Security_and_Policy.md) | Stream 5 |
| [07_Reference_Architecture.md](./07_Reference_Architecture.md) | End-to-end architecture and pattern catalogue |
| [08_Gap_Analysis_and_Recommendations.md](./08_Gap_Analysis_and_Recommendations.md) | Gaps, standards proposals, upstream routing |
| [CHANGELOG.md](./CHANGELOG.md) | Versioned releases of the document |

## Milestones

1. Working sessions 1 and 2 (September 24 and October 8, 2026): direction, terms of reference, workstreams and owners confirmed. Done.
2. Stream outlines in this directory: October 22, 2026.
3. First drafts of all streams: November 2026.
4. Reference architecture first draft and document release v0.1: before KubeCon + CloudNativeCon Europe, March 2027.
5. v1.0 release after community review, date to be set at the February 2027 session.

## Related efforts

- [CNCF Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) (the base this initiative extends)
- Kubernetes Gateway API Inference Extensions and the Agentic Network working group
- Agentic AI Foundation (AAIF)
- LF Edge SuperBlueprint
- ETSI cloud native agentic blueprint for telco
- LF AI & Data AI Safety Best Practices
