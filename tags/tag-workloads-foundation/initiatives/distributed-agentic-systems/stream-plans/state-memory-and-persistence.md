# Stream plan: state, memory and persistence

Technical chapter: [Chapter 05](../05_State_Memory_and_Persistence.md)
Stream owner: Victor Lu, as recorded in the [initiative README](../README.md#workstreams)
Team: to identify; no additional contributors confirmed here
Status: forming
Last updated: 2026-11-05
Next walkthrough: to agree

**Starter draft for team review.** This does not reproduce Victor's working draft. Questions and outputs below are suggestions from the chapter outline and meeting notes, not agreed assignments or conclusions. Use the [planning guide](./README.md), [full template](./TEMPLATE.md), and separate [illustrative example](./EXAMPLE-state-memory.md) to develop this plan in 45–60 minutes.

## 1. Purpose and intended outcome

Help platform teams define state ownership, persistence, and recovery for agent conversation context, durable tasks, and long-term memory, building on existing database and Kubernetes patterns.

Proposed minimum useful outcome for v0.1: guiding principles, a state ownership classification, and recovery/retention decision guidance illustrated by an interrupted task.

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| To identify | Kubernetes data services, agent memory, recovery, or distributed operations | Investigate, draft, review, or upstream liaison | Not yet identified |

- Missing perspectives and recruitment action: agree outreach for database recovery, checkpoint integrity, data tenancy, and multi-cluster/edge practitioners.
- How to join: contact the stream owner in `#initiative-distributed-agentic-systems` with your experience or proposed contribution.
- Suggested first task: compare a documented storage recovery mechanism with application checkpoint requirements; identify what each can and cannot establish.

## 3. Scope and priority questions

- In scope: short-term context, long-term memory, storage/recovery patterns, replication, checkpoint integrity, and garbage collection.
- Out of scope: building a memory framework, choosing a preferred database, or defining another general-purpose standard without evidence of need.
- Assumptions and terminology: distinguish ephemeral context, durable task state, conversation history, and retrievable memory; validate terms with the team.
- Single-cluster baseline: define which state outlives agent Pod replacement and who owns it.
- Distributed considerations: replication, recovery, and context management across clusters and disconnected edge/telco environments.
- Boundaries: agree process/task recovery with runtime, reconnect/replay with protocol and gateway, and retention/access/evidence with governance.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| Proposed v0.1 | What state exists, who owns it, and how long must it survive? | State classification and guiding principles | To agree |
| Proposed v0.1 | Which existing storage and recovery capabilities meet agent requirements? | Capability and limitation mapping | To agree |
| Proposed v0.1 | What is required for checkpoint integrity and safe task resumption? | Recovery pattern and trade-offs | To agree |
| Proposed v0.1 | How are shared state, retention, and garbage collection coordinated? | Lifecycle and deployment decision guidance | To agree |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) | Availability; governance privacy; security data access | Baseline for resilience, retention, and isolation | Published March 23, 2026; task-level state semantics require further investigation |
| [October 22 discussion](../meeting-notes/2026-10-22.md) and [chapter outline](../05_State_Memory_and_Persistence.md) | Problem-first approach, database patterns, recovery, checkpointing, distributed context | Existing direction; coordinate with the owner's working draft | Notes and outline are not implementation evidence; select primary sources during kickoff |

## 5. Milestones and substantive work

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Review starter against existing draft; confirm team and scope | To agree | To agree | Proposed | Not yet created |
| First useful chapter draft and interrupted-task walkthrough | To agree | To agree | Proposed | Not yet created |
| Architecture and gap-analysis inputs | To agree | To agree | Proposed | Not yet created |
| Cross-stream review and integration | To agree | To agree | Proposed | Not yet created |

Create issues for substantive work once scoped; keep detailed tasks in those issues.

## 6. Current biweekly walkthrough — to schedule

- Changes or findings: none recorded in this starter; review the existing working draft with the owner.
- Suggested material: mark state and recovery boundaries before and after a Pod restart in the [shared reference scenario](./README.md#proposed-shared-reference-scenario).
- Feedback needed: when does persisted data suffice, and when is application-level coordination needed to resume safely?
- Blocker / help needed: confirm contributor and runtime/governance review coverage.
- Next output, contributor, and checkpoint: agree at kickoff; link meeting notes, issues, and PRs here.

## 7. Findings and integration — develop progressively

No findings established here. Use the [template's findings table and handoff prompts](./TEMPLATE.md#7-findings-and-integration--develop-progressively). Link state-store responsibilities and recovery flows to [chapter 07](../07_Reference_Architecture.md). Route only evidenced limitations to [chapter 08](../08_Gap_Analysis_and_Recommendations.md); distinguish documentation, implementation, interoperability, and specification needs.

## 8. Review and readiness — use later

Use the [common readiness checklist](./TEMPLATE.md#8-review-and-readiness--use-later). Record review feedback and deferred questions here or in linked issues. This starter is not ready for integration as technical guidance.
