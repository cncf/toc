# Illustrative stream plan: state, memory and persistence

**Example only.** This is a fictional starter exercise informed by the [October 22 discussion](../meeting-notes/2026-10-22.md), not Victor Lu's plan, a record of his draft, or an assignment to any contributor. Roles, activities, milestones, and proposed technical directions below require agreement in a real plan.

Technical chapter: [Chapter 05](../05_State_Memory_and_Persistence.md)
Stream owner: Example facilitator (unassigned)
Team: Example drafting and reviewer roles below; no participants assigned
Status: scoping (illustrative)
Last updated: 2026-11-05
Next walkthrough: to agree

## 1. Purpose and intended outcome

Help platform teams choose ownership, persistence, and recovery boundaries for agent task state, conversation context, and longer-term memory. A restarted Pod does not by itself recover the meaning of an agent task, and replaying a tool call may repeat an external side effect.

Minimum useful outcome for v0.1: a state classification and a recovery decision table, illustrated by one interrupted task. Do not attempt to prescribe a universal memory architecture.

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| Unassigned: state author | Kubernetes persistence and agent state | Draft state classification | Needed; not invited |
| Unassigned: practitioner reviewer | Task recovery in an operating system | Challenge recovery assumptions | Needed; not invited |
| Unassigned: security reviewer | Data tenancy and retention | Review access and deletion boundaries | Needed; not invited |

- Missing expertise: production agent operations, database recovery, checkpoint integrity, and disconnected deployment experience.
- Recruitment action: facilitator to agree an outreach owner and checkpoint at kickoff.
- How to join: introduce your relevant experience in `#initiative-distributed-agentic-systems` and ask the actual stream owner where help is needed.
- First task for a new contributor: compare what a volume snapshot preserves with what an interrupted agent task needs to resume; bring primary documentation and list assumptions.

## 3. Scope and priority questions

- In scope: state categories, ownership, checkpoint boundaries, recovery, retention, and garbage collection.
- Out of scope: building a memory framework, choosing a preferred database, or designing a model's reasoning strategy.
- Assumptions and terminology: distinguish ephemeral context, durable task state, conversation history, and retrievable long-term memory; validate these definitions with the team.
- Single-cluster baseline: the coordinator and specialist run as Kubernetes workloads, with durable storage accessible after Pod replacement.
- Distributed considerations: replication lag, conflicting updates, and disconnected recovery are follow-on questions, not assumed solved by storage replication.
- Boundaries: runtime owns process lifecycle; state owns durable recovery information; protocol and gateway teams coordinate reconnect and replay behavior; governance reviews access, retention, and audit implications. These leads are proposed, not agreed.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| v0.1 | What state must survive a Pod restart, and who owns it? | State classification and ownership table | Unassigned |
| v0.1 | What makes a checkpoint sufficient and safe to resume? | Recovery pattern and limitations | Unassigned |
| v0.1 | How are duplicate tool actions and stale context handled? | Cross-stream decision guidance | Unassigned |
| v0.1 | When can task data be deleted without breaking recovery or audit needs? | Retention and garbage-collection considerations | Unassigned |
| Later | What changes with replication across disconnected clusters? | Extension questions and evidence needs | Unassigned |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) | Availability and fault tolerance; governance data privacy; security agent data access | Baseline for recovery, retention, and isolation considerations | Published March 23, 2026; not evidence of task-level checkpoint semantics |
| [Kubernetes persistent volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) | Volume lifecycle independent of a particular Pod | Starting point for persistence boundaries | Select documentation version during investigation; does not establish application recovery correctness |
| [Kubernetes volume snapshots](https://kubernetes.io/docs/concepts/storage/volume-snapshots/) | Storage snapshot capability | Compare storage recovery with task checkpoint requirements | Driver capabilities and application consistency need verification |
| [Existing use-case template](../use-cases/TEMPLATE.md) | Runtime status, decisions, and gaps | Collect practitioner evidence without treating a design as production experience | No system evidence collected for this fictional example |

These are starting sources, not a complete ecosystem comparison. Investigate database/operator recovery patterns and agent checkpoint implementations before recommending an approach.

## 5. Milestones and substantive work

All targets and contributors below are illustrative and unassigned. No issues have been created for this example.

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Review definitions and priority questions | Agree at kickoff | Initial walkthrough | Proposed | Not created |
| Draft state ownership table and interrupted-task flow | Agree at kickoff | Following walkthrough | Proposed | Not created |
| Review replay and retention boundaries with adjacent streams | To agree | After first draft | Proposed | Not created |
| Integrate chapter, architecture inputs, and any evidenced gaps | To agree | Before release review | Proposed | Not created |

## 6. Current biweekly walkthrough — proposed agenda

- What changed or did we learn? No findings yet; scope and candidate sources are prepared.
- Concrete material: draw the [shared incident-investigation scenario](./README.md#proposed-shared-reference-scenario), marking where state exists before and after a Pod restart.
- Feedback needed: does the checkpoint describe only agent progress, or also the outcome of external actions? Who resolves an uncertain tool-call outcome?
- Help needed: runtime, protocol, gateway, and governance reviewers for the relevant boundaries.
- Next output: a state ownership table and a recovery flow, with contributors and checkpoint to agree at the walkthrough.
- Related records: link the actual meeting notes or issues once discussion occurs.

This is a design walkthrough, not a tested recovery guarantee.

## 7. Findings and integration — develop progressively

No findings are established by this example. The following is a proposed investigation, not a confirmed ecosystem gap.

| Finding / pattern / open gap | Evidence and limitations | Related streams / coordination needed | Destination / next action |
|---|---|---|---|
| Investigate whether storage recovery alone is sufficient for safe task resumption | Open question; compare documented application semantics and practitioner experience | Runtime lifecycle, gateway reconnect, protocol retries, governance evidence | Investigate before proposing chapter guidance or a gap |

### Reference architecture handoff

Proposed input: component responsibilities for coordinator, specialist, checkpoint store, and tool service; a restart sequence with ownership and decision points. Link reviewed content to [chapter 07](../07_Reference_Architecture.md) once drafted.

### Gap-analysis handoff

No standardization proposal yet. If investigation reveals a missing recovery contract, document what existing capabilities cannot do and distinguish guidance, implementation, and interoperability needs before proposing an upstream destination in [chapter 08](../08_Gap_Analysis_and_Recommendations.md).

## 8. Review and readiness — use later

- [ ] Definitions and scope reviewed by the actual team.
- [ ] Persistence and task-recovery claims supported by evidence.
- [ ] Replay, tenancy, retention, and failure implications reviewed across streams.
- [ ] Architecture inputs and any evidenced gaps linked.
- [ ] Deferred questions, review feedback, and follow-up work recorded.

No items are checked because this example has not undergone technical review. Use the full checklist in [TEMPLATE.md](./TEMPLATE.md) for a real plan.
