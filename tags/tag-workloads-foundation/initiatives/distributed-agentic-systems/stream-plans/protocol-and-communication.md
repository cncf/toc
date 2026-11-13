# Stream plan: protocol and communication interoperability

Technical chapter: [Chapter 02](../02_Protocol_and_Communication_Interoperability.md)
Stream owner: Moeez Khan (@MuhammadMoeezKhan), Independent; volunteered in `#initiative-distributed-agentic-systems` on November 5, 2026, lazy consensus to November 9
Team: owner confirmed; contributors invited, none confirmed yet (see section 2)
Status: scoping
Last updated: 2026-11-13
Next walkthrough: November 19, 2026 session

## 1. Purpose and intended outcome

Platform teams increasingly run agents and tools that speak MCP and A2A inside Kubernetes, and some also carry payment authorization (AP2) or call models through a common interface (Open Responses). Kubernetes already tells a caller where an endpoint is. It does not tell the caller what that endpoint can do, whether its contract is still the one the caller validated against, or what happened when a call failed partway. This chapter helps platform engineers, gateway and runtime implementers, and protocol maintainers decide:

- where agent and tool capabilities are registered and discovered in a cluster, and how they stay current;
- where payload schemas and protocol versions are validated (caller, gateway, or server), and at what cost;
- which guarantees each protocol actually provides, and which the platform must add, especially for failures and outcomes that are uncertain after a lost response.

What differs from distributed systems generally: capabilities are described for a model and discovered at runtime; contracts change while clients are connected; interactions can be long-running tasks with interrupted states; capability metadata is mostly self-declared by the server; and a lost response to a side-effecting call leaves an outcome the caller cannot infer.

Minimum useful outcome for v0.1, within about 1.5 pages:

1. A responsibility comparison of MCP, A2A and AP2, plus a note on Open Responses, with pinned versions: what each specifies for discovery, contract and schema, version negotiation, authentication, long-running work, and failure representation, and where each contract stops.
2. Discovery and validation patterns for a single cluster, with their failure behavior, walked through the shared reference scenario.
3. A short list of evidence-backed interoperability gaps for chapter 08.

The chapter does not design protocols. It describes what the specifications define and what a Kubernetes platform needs to support them natively.

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| Moeez Khan (@MuhammadMoeezKhan) | MCP specification and client/server behavior under failure: lost responses, retries, partial results | Owner; scope, MCP column, failure representation, drafting, coordination | Confirmed |
| Akanksha Trehun (@magic-peach) | Gateway and data plane stream owner; listed protocol interest | Transport, streaming, and session boundary with chapter 03 | Invited |
| Aman Singh (@amansinghoriginal) | Chapter 07 owner; MCP Working Group for Triggers; event-driven execution | Component roles and event delivery boundary with chapter 07 | Invited |
| Mohammad Hassan | Listed protocol interoperability and runtime interest | A2A investigation and review | Invited |

- Missing expertise or perspectives: A2A implementation experience; AP2 or agent-payments practice; registry and naming (MCP Registry, ANS, IETF DAWN); and practitioners running MCP or A2A in production on Kubernetes. Cross-stream reviewers from gateway, runtime, state, and governance.
- Recruitment action: the owner posts a call in `#initiative-distributed-agentic-systems` once this plan is open for review, with a first task per area, and asks practitioners who share use cases through the [use-case template](../use-cases/TEMPLATE.md) to review the comparison. Checkpoint: contributors named by the December 3 session.
- How to join: reply in `#initiative-distributed-agentic-systems` or comment on this plan's pull request with the area you would like to take.
- First task for a new contributor: take one row of the responsibility comparison (for example, how A2A 1.0 and MCP 2026-07-28 represent a failed or interrupted call) and fill it from the primary specification, with section references and anything the specification leaves unsaid. About an hour; the draft table lives in the working document.

People listed as invited have not agreed to any work. Rows change to confirmed when they say so.

## 3. Scope and priority questions

- In scope: MCP, A2A and AP2 as used between agents and tools in a cluster; agent discovery and naming; dynamic registration and deregistration; filtering; payload schema and protocol version validation; how each protocol represents errors, interrupted work, and uncertain outcomes. Open Responses is covered as a related model-facing interface, to say where it sits relative to the others.
- Out of scope: implementing a new protocol, selecting a preferred framework, product comparisons, and protocol governance (which belongs to the protocol projects and AAIF).
- Assumptions and terminology: define agent, tool, capability, registry, discovery, contract, and outcome before comparing mechanisms, and align them with chapter 07's shared terms. Note that "protocol session" means different things across versions: MCP 2026-07-28 removed protocol-level sessions that 2025-11-25 had.
- Single-cluster baseline: Kubernetes Service DNS and EndpointSlices give endpoint discovery, not capability discovery. The baseline is agents and tool servers as workloads in one cluster, reached directly or through a gateway.
- Multi-cluster, edge, or disconnected considerations: cross-cluster naming and trust (ANS v2, IETF DAWN), registry unavailability, and static registration for air-gapped sites. Recorded as later investigations; the Agentic Standards already flag the air-gap case.
- Boundaries with other streams (proposed leads, to agree with each owner):
  - Gateway (chapter 03): protocol describes what each specification says about streams, sessions, headers, and reconnection; gateway leads on routing, affinity, and proxy behavior. MCP 2026-07-28's removal of `Mcp-Session-Id` and its new `Mcp-Method` and `Mcp-Name` headers affect both chapters.
  - Runtime (chapter 04): runtime leads on who registers and deregisters an agent as Pods come and go; protocol covers what registration metadata contains and how clients learn it changed.
  - State (chapter 05): state leads on durable deduplication and checkpoints; protocol covers what each protocol offers for duplicate detection and retries (for example, A2A `messageId` and MCP request re-issue rules).
  - Governance (chapter 06): governance leads on identity, authorization enforcement, and per-call evidence (6a); protocol covers the security schemes, card signing, and mandates the specifications define.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| v0.1 | Which problems do MCP, A2A and AP2 address, and where do their contracts stop? | Responsibility comparison with pinned versions | Moeez (MCP); help needed for A2A and AP2 |
| v0.1 | How are capabilities registered, discovered, and kept current in a cluster, and what happens when that information is stale or the registry is unavailable? | Discovery patterns and failure considerations | Help needed (registry and naming) |
| v0.1 | Where are payload schemas and protocol versions validated, and what does each placement cost? | Validation decision guidance | To agree |
| v0.1 | How does each protocol represent errors, interrupted work, and outcomes that are unknown after a lost response, and what must the platform add? | Failure representation comparison (the stream's shared-scenario question) | Moeez |
| v0.1 | Which interoperability needs remain unmet by the specifications and current integrations? | Evidence-backed findings for chapter 08 | All |
| Later | How do naming, trust, and discovery work across clusters and organizations? | Extension questions and evidence needs | Help needed |
| Later | How does delegated authorization behave when one agent delegates a payment to another? (AP2 0.2 leaves agent-to-agent delegation out of scope.) | Investigation | Help needed |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) | Control and communication (tools and services, schema validation, discovery and agent registries); Security (agent identity, tenancy, data access) | Baseline to reference, not restate | Published March 23, 2026. Recommends OWASP ANS, which has since moved to IETF drafts |
| [MCP specification](https://modelcontextprotocol.io/specification/2026-07-28) | Stateless requests, `server/discover`, `Mcp-Method` and `Mcp-Name` headers, cacheable list results, error codes, tool annotations as untrusted hints | Contract, discovery freshness, validation points, failure behavior | Revision 2026-07-28; many deployed servers still speak 2025-11-25, so record both where behavior differs |
| [A2A specification](https://a2a-protocol.org/latest/specification/) | Agent Card and `/.well-known/agent-card.json`, card signing (JWS), task states, errors, version negotiation, `messageId` deduplication | Agent discovery, contract, interrupted and failed tasks | 1.0 (v1.0.0 March 12, 2026; v1.0.1 May 28, 2026); documented capability |
| [AP2](https://github.com/google-agentic-commerce/AP2) | Checkout and Payment Mandates (SD-JWT), roles, verification rules | Delegated authorization carried over agent protocols | v0.2.0, April 28, 2026; defines a security layer inside a commerce protocol, not transport or discovery |
| [Open Responses](https://www.openresponses.org/specification) | Model-facing interface: items, streaming events, tool calls, error types | Where a model interface ends and agent and tool protocols begin | Spec 2026-04-24; joined AAIF October 22, 2026 |
| [MCP Registry](https://github.com/modelcontextprotocol/registry) and its generic registry API | Server metadata (`server.json`), versions, status, subregistries | One registry pattern for tool servers | Preview; API v0.1 frozen since October 2025; v1.8.8, November 6, 2026 |
| [IETF DAWN working group](https://datatracker.ietf.org/wg/dawn/about/) and [ANS v2 draft](https://datatracker.ietf.org/doc/draft-narajala-courtney-ansv2/) | Initial discovery of agents with names (DNS, mDNS); DNS-anchored agent identity with a transparency log | Naming, cross-domain discovery, and the Standards' ANS recommendation | DAWN chartered 2026; ANS v2 is an individual draft, -02 of October 26, 2026; neither is a standard yet |
| [AGNTCY Directory](https://github.com/agntcy/dir) and [OASF](https://github.com/agntcy/oasf) | Distributed announce and discovery; agent record schema | A second registry approach to compare | Directory v1.8.0, October 15, 2026; OASF 1.1.0, July 10, 2026; not yet reviewed in depth |
| [Kubernetes Services and DNS](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/) | Endpoint discovery | Baseline for what the cluster already provides | Endpoint, not capability, discovery |
| [kagent](https://github.com/kagent-dev/kagent) | `Agent` and `RemoteMCPServer` custom resources | One Kubernetes-native registration pattern | v1.0.0, November 9, 2026; one implementation among several, not an endorsement |
| [Kubeflow agentic RAG write-up](https://github.com/cncf/endusers/issues/2114) | MCP tools and A2A routing behind a gateway on Kubernetes | Practitioner evidence for discovery, validation, and gateway placement | Documented by its authors as deployed; not verified by this stream |
| [agent-fault-lab](https://github.com/MuhammadMoeezKhan/agent-fault-lab) | Replayable fault scenarios for MCP tool calls (lost response after commit, timeouts) | One test surface for the failure representation question | Maintained by the stream owner; one test surface, not the basis for any recommendation |

- Which Agentic Standards recommendations apply, and what are we adding? The Standards ask teams to standardize tool access protocols, validate payloads with JSON Schema, Protobuf or OpenAPI, use Kubernetes-native DNS, service mesh, or purpose-built registries for discovery, consider static or multicast registration when air-gapped, and use a verifiable naming system such as ANS. This chapter adds which specification versions provide what, where in a cluster validation and discovery run, how discovery stays current and fails, and how failures and unknown outcomes are represented across protocols.
- What do Kubernetes and existing projects already provide? Endpoint discovery (Services, DNS, EndpointSlices), workload identity and mTLS through service meshes, and custom resources for declaring agents and tool servers. Still uncertain: capability metadata freshness, cross-namespace discovery policy, and who checks that declared capabilities match behavior.
- Where do sources disagree? The Standards point to OWASP ANS while naming work has moved to IETF (ANS v2, DAWN). Gateway framing written against MCP sessions predates their removal in 2026-07-28. AP2 0.2 describes itself both as an extension for A2A and MCP and as a layer inside a commerce protocol. Upstream contacts to approach: MCP maintainers and the MCP Registry working group, A2A maintainers, and DAWN participants.

## 5. Milestones and substantive work

Targets below are the owner's own; contributor targets are agreed with them before they are added.

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Initial scope and team plan | Moeez Khan | November 19, 2026 session | In review | This pull request |
| Chapter 02 bullet outline and draft responsibility comparison | Moeez Khan; reviewers to agree | Outline November 19; comparison reviewed December 3 | In progress | Working document tab; issue to create after November 19 |
| Discovery, validation, and failure patterns on the shared scenario | Moeez Khan; contributors to agree | December 17, 2026 | Planned | Issues to create after November 19 |
| First useful chapter draft | Moeez Khan; contributors to agree | January 14, 2027 session | Planned | PR to chapter 02 |
| Architecture and gap-analysis inputs | Moeez Khan; with chapter 07 (Aman Singh) and the chapter 08 owner, to agree | January 28, 2027 | Planned | Not yet created |
| Cross-stream review and integration | Moeez Khan; affected stream owners | February 2027 sessions; chapter final before the v0.1 release review, ahead of KubeCon + CloudNativeCon Europe, March 2027 | Planned | Not yet created |

Drafting through December happens in a tab of the [working document](https://docs.google.com/document/d/1bRWy8vhke16mMhpPlEE_Dv71fTmkRWG-4Cy1hdTHUb0/edit) so contributors can comment freely. The reviewed result moves into the chapter file by pull request from January.

## 6. Current biweekly walkthrough: November 19, 2026

- What changed or did we learn? An owner volunteered on November 5 and this plan opened. The protocol table at the November 8 Maintainer Summit asked: "When an agent in your cluster calls a tool or another agent over MCP or A2A, how does it know what it's allowed to call, and whether a call that timed out actually ran?" Answers will be added when shared.
- What concrete material can we walk through? The chapter 02 bullet outline and a draft responsibility comparison, plus the shared scenario: the coordinator finds the specialist through its Agent Card and calls an MCP tool whose response is lost. What does each specification say happens next?
- Which recommendation, assumption, or boundary needs feedback?
  - Where do protocol guarantees end and gateway responsibilities begin for streams and sessions, given MCP 2026-07-28 is stateless?
  - Should AP2 get its own short section or a note, given payments sit outside the shared scenario?
  - Is Open Responses in chapter 02's scope beyond a positioning note?
- What is blocked, and what help is needed? Contributors for A2A, AP2, and registry and naming; practitioner reviewers.
- What will we bring to the next meeting? The comparison reviewed by at least one contributor per protocol column (owner; December 3).
- Related issues, PRs, or meeting notes: [creation issue #1746](https://github.com/cncf/toc/issues/1746); [chapter 07 skeleton, #2367](https://github.com/cncf/toc/pull/2367); [6a outline, #2369](https://github.com/cncf/toc/pull/2369).

## 7. Findings and integration: develop progressively

No findings are established yet. The rows below are candidate investigations drawn from specification text; none has been demonstrated or reviewed.

| Finding / pattern / open gap | Evidence and limitations | Related streams / coordination needed | Destination / next action |
|---|---|---|---|
| Session-based routing assumptions are version-dependent | MCP 2026-07-28 changelog removes protocol sessions and `Mcp-Session-Id`; documented only | Gateway (lead) | Agree wording with chapter 03 |
| Neither MCP nor A2A requires an idempotency key for side-effecting calls | A2A 1.0 section 3.3.1: agents "may utilize the messageId to detect duplicate messages"; MCP 2026-07-28: a client re-issues a request lost on a broken stream "as a new request with a new request ID"; documented only | State, governance (6a attempt identity) | Investigate with fault scenarios; candidate chapter 08 interoperability gap |
| Capability metadata is self-declared | A2A Agent Cards can be signed (JWS) for authenticity, which does not establish behavior; MCP tool annotations are hints to treat as untrusted unless the server is trusted; documented only | Governance | Investigate; decide whether this is guidance or a gap |
| No Kubernetes-native capability registry; several registry designs coexist | Kubernetes Services, MCP Registry generic API, AGNTCY Directory, custom resources; documented only | Runtime (registration lifecycle) | Pattern comparison for chapter 02; possible gap |

For each substantive pattern the chapter uses: Problem, Context, Approach, Alternatives and trade-offs, Failure behavior, Evidence, Related streams.

### Reference architecture handoff

Proposed input to [chapter 07](../07_Reference_Architecture.md): component roles (capability registry, Agent Card or server metadata endpoint, schema and version validator, protocol-aware intermediary) and the discovery and call flows for the shared scenario, with decision points and failure behavior. Also align chapter 07's "protocol session vs. network connection" term with MCP 2026-07-28. To review with Aman Singh.

### Gap-analysis handoff

None proposed yet. Candidate gaps are listed in the table above and go to [chapter 08](../08_Gap_Analysis_and_Recommendations.md) only with evidence, the existing capability and its limitation, the gap type (documentation, implementation, interoperability, or specification), and a suggested upstream destination (for example, the MCP or A2A projects, or IETF DAWN), without implying acceptance.

## 8. Review and readiness: use later

Uses the [common readiness checklist](./TEMPLATE.md#8-review-and-readiness--use-later). No items are checked; this plan has not been reviewed.
