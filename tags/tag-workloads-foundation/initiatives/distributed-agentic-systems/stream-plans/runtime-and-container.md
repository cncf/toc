# Stream plan: runtime and container abstraction

Technical chapter: [Chapter 04](../04_Runtime_and_Container_Abstraction.md)
Stream owner: Srikanth Rao, as recorded in the [initiative README](../README.md#workstreams)
Team: to identify; no additional contributors confirmed here
Status: forming
Last updated: 2026-11-05
Next walkthrough: to agree

**Starter draft for team review.** Questions and outputs below are suggestions based on the chapter outline, not agreed assignments or conclusions. Use the [planning guide](./README.md) and [full template](./TEMPLATE.md) to develop this plan in 45–60 minutes.

## 1. Purpose and intended outcome

Help platform teams choose workload, isolation, scheduling, and lifecycle boundaries for agents on Kubernetes, distinguishing process recovery from task recovery.

Proposed minimum useful outcome for v0.1: deployment and isolation decision guidance plus a restart/scaling walkthrough showing where existing Kubernetes primitives suffice and where application cooperation is needed.

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| To identify | Kubernetes workloads, sandboxing, scheduling, or agent operations | Investigate, draft, review, or upstream liaison | Not yet identified |

- Missing perspectives and recruitment action: agree outreach for sandbox/runtime maintainers, scheduling expertise, and practitioners handling agent task recovery.
- How to join: contact the stream owner in `#initiative-distributed-agentic-systems` with your experience or proposed contribution.
- Suggested first task: compare an ordinary Pod deployment with a sandboxed workload for one explicit threat model; record evidence and trade-offs rather than selecting a preferred project.

## 3. Scope and priority questions

- In scope: Pod, sidecar, and CRD modeling; container practices; microVM/confidential isolation; HPA, Kueue, DRA; lifecycle and task recovery.
- Out of scope: building a runtime, deep-diving one framework, or assuming a new Agent CRD is required.
- Assumptions and terminology: define agent instance, execution boundary, task lifetime, and recovery ownership.
- Single-cluster baseline: Kubernetes workload lifecycle and scheduling within one cluster.
- Distributed considerations: placement and recovery assumptions across clusters, edge constraints, and unavailable dependencies.
- Boundaries: agree durable checkpoint semantics with state, registration with protocol, draining/reconnect with gateway, and isolation/identity requirements with governance.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| Proposed v0.1 | When do Pods, sidecars, or higher-level APIs provide suitable lifecycle boundaries? | Workload decision guidance | To agree |
| Proposed v0.1 | Which threat and tenancy requirements justify stronger isolation? | Isolation alternatives and trade-offs | To agree |
| Proposed v0.1 | Which signals and resource constraints should drive scaling and scheduling? | Capability mapping and operational considerations | To agree |
| Proposed v0.1 | What must cooperate with Kubernetes to resume or safely terminate a task? | Restart and recovery pattern | To agree |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) | General container practices; security runtime isolation | Baseline to reference rather than repeat | Published March 23, 2026; agent-specific lifecycle extensions need evidence |
| [Chapter 04 outline](../04_Runtime_and_Container_Abstraction.md) | Workload, sandbox, scaling, and scheduling candidates | Starting inventory including Agent Sandbox, kagent, HPA, Kueue, and DRA | Select primary sources and versions; inclusion is not endorsement or validation |

## 5. Milestones and substantive work

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Confirm team, scope, and source inventory | To agree | To agree | Proposed | Not yet created |
| First useful chapter draft and lifecycle walkthrough | To agree | To agree | Proposed | Not yet created |
| Architecture and gap-analysis inputs | To agree | To agree | Proposed | Not yet created |
| Cross-stream review and integration | To agree | To agree | Proposed | Not yet created |

Create issues for substantive work once scoped; keep detailed tasks in those issues.

## 6. Current biweekly walkthrough — to schedule

- Changes or findings: none recorded in this starter.
- Suggested material: replace a specialist agent Pod during the [shared reference scenario](./README.md#proposed-shared-reference-scenario); identify lifecycle, identity, and state transitions.
- Feedback needed: what can restart automatically, and what requires task-level recovery logic?
- Blocker / help needed: confirm contributor and state/governance review coverage.
- Next output, contributor, and checkpoint: agree at kickoff; link meeting notes, issues, and PRs here.

## 7. Findings and integration — develop progressively

No findings established here. Use the [template's findings table and handoff prompts](./TEMPLATE.md#7-findings-and-integration--develop-progressively). Link workload boundaries, lifecycle flows, and isolation decisions to [chapter 07](../07_Reference_Architecture.md). Document evidenced limitations before proposing APIs or other upstream work in [chapter 08](../08_Gap_Analysis_and_Recommendations.md).

## 8. Review and readiness — use later

Use the [common readiness checklist](./TEMPLATE.md#8-review-and-readiness--use-later). Record review feedback and deferred questions here or in linked issues. This starter is not ready for integration as technical guidance.
