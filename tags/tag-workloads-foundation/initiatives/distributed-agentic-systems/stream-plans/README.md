# Stream planning guide

This guide supports the common section-owner template requested at the [October 22, 2026 working session](../meeting-notes/2026-10-22.md). It is for the whole stream team: owners, authors, reviewers, practitioners, and new contributors.

A stream plan starts as a lightweight onboarding document and becomes an ongoing delivery record. It is not a second technical chapter or an additional approval process. The [charter](../charter.md) continues to govern scope, vendor neutrality, decisions, and contributions.

## Start together in 45–60 minutes

Use your stream's starter plan below. For a new stream, copy [TEMPLATE.md](./TEMPLATE.md). The [state and memory example](./EXAMPLE-state-memory.md) illustrates the expected level of detail; it does not assign work or commit the actual stream team.

| Time | Activity | Result |
|---|---|---|
| 5 minutes | Read the charter, chapter outline, and relevant baseline standards | Shared context |
| 10 minutes | Define the problem, intended readers, and minimum useful v0.1 outcome | Section 1 |
| 10 minutes | Identify contributors, missing expertise, and a first task for newcomers | Section 2 |
| 10–15 minutes | Agree scope boundaries and three to five priority questions | Section 3 |
| 5–10 minutes | Link initial sources and existing capabilities | Section 4 |
| 5–10 minutes | Identify milestones and the next walkthrough contribution | Sections 5–6 |

Do not try to finish the research in this exercise. Mark unknowns explicitly and name a next action where possible. Drafting, investigation, and review can happen together. Develop sections 7–8 as findings emerge; they are not entry requirements.

## Stream plans

These are starter drafts based on the existing chapter outlines. Recorded owners are carried over from the initiative README; suggested questions, outputs, and milestones need team review. No new participation or delivery commitments are implied.

| Stream | Plan | Technical chapter |
|---|---|---|
| Protocol and communication interoperability | [Starter plan](./protocol-and-communication.md) | [Chapter 02](../02_Protocol_and_Communication_Interoperability.md) |
| Gateway and data plane | [Starter plan](./gateway-and-data-plane.md) | [Chapter 03](../03_Gateway_and_Data_Plane.md) |
| Runtime and container abstraction | [Starter plan](./runtime-and-container.md) | [Chapter 04](../04_Runtime_and_Container_Abstraction.md) |
| State, memory and persistence | [Starter plan](./state-memory-and-persistence.md) | [Chapter 05](../05_State_Memory_and_Persistence.md) |
| Governance, observability, security and policy | [Starter plan](./governance-observability-security-policy.md) | [Chapter 06](../06_Governance_Observability_Security_and_Policy.md) |

## Develop the plan through biweekly walkthroughs

Proposed meeting format, subject to community agreement:

- Brief update from every stream: progress, next output, and help needed.
- Rotate deeper walkthroughs across streams, with time for affected teams to discuss boundaries and decisions. Choose the next streams in advance according to available meeting time and dependencies.
- Walk through concrete material: a problem statement, diagram, source comparison, draft pattern, failure scenario, or chapter passage. A software demo is not required.
- Ask what evidence supports a recommendation, what is still assumed, and what changes for other streams.
- Record decisions and actions in meeting notes or linked issues. Update the plan's current walkthrough and milestones; do not maintain a parallel meeting archive.

Documented capabilities and expert review can support early drafts. Distinguish those from demonstrated or production experience. A shared-scenario walkthrough helps expose assumptions; it is not proof that an implementation works or a requirement to build new software.

Use a small milestone table in each plan. Link substantive investigations, drafting tasks, and unresolved decisions to issues; link contributions to PRs. Detailed work belongs in those issues and PRs, not in duplicate task lists. Use target checkpoints agreed by the team rather than inventing deadlines or commitments for others.

## Proposed shared reference scenario

Use one simple scenario to connect the streams and inform [chapter 07](../07_Reference_Architecture.md). It is a design exercise, not a claim about a running system or a required implementation stack.

A user asks a coordinator agent to investigate an operational incident and propose a remediation. The coordinator delegates analysis to a specialist agent and retrieves information through an MCP tool service. It stores a durable task checkpoint and produces a recommendation. Any state-changing remediation requires explicit human approval and scoped authorization. The baseline is a single Kubernetes cluster.

| Stream | Questions to bring to the scenario |
|---|---|
| Protocol | How are agent and tool capabilities discovered, contracts validated, and failures represented? |
| Gateway | How are long-lived interactions routed, sessions maintained, and disconnects handled? |
| Runtime | What is the workload and isolation boundary, and what happens if an agent Pod restarts? |
| State | What survives restart, who owns it, and how are duplicate or stale actions avoided? |
| Governance | How are identity, delegation, approval, tenancy, telemetry, and audit evidence enforced? |

Start with a Pod restart and a broken connection. Later consider a second cluster or disconnected edge deployment where relevant, identifying assumptions that no longer hold. Streams can use additional examples when the common scenario does not exercise their important concerns, such as agent payments.

Keep the scenario framework- and vendor-neutral. If a walkthrough describes a real system, document it through the existing [use-case template](../use-cases/TEMPLATE.md), including whether it is running in production, in test, or is design only.

## Keep technical outputs consistent, not identical

Teams can organize chapters differently. Each should cover the problem and terminology, baseline standards, existing capabilities, patterns and decision guidance, operational and failure considerations, evidence, and unresolved gaps.

For each substantive pattern, use:

> Problem → Context → Approach → Alternatives and trade-offs → Failure behavior → Evidence → Related streams

Provide component roles, interfaces, flows, and decision points to [chapter 07](../07_Reference_Architecture.md). Provide evidence-backed unmet requirements and candidate upstream destinations to [chapter 08](../08_Gap_Analysis_and_Recommendations.md). Link to the technical content rather than copying it into the plan.

Reference the [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) and explain the additional Kubernetes guidance, rather than restating the baseline. The streams reorganize and extend that material; cite the actual relevant sections rather than assuming identical headings. Investigate capabilities before proposing standards. Projects illustrate patterns, never constitute endorsements, and no chapter should be based on a single project.
