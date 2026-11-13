# Stream plan: gateway and data plane

Technical chapter: [Chapter 03](../03_Gateway_and_Data_Plane.md)
Stream owner: Akanksha Trehun, as recorded in the [initiative README](../README.md#workstreams)
Team: to identify; no additional contributors confirmed here
Status: forming
Last updated: 2026-11-05
Next walkthrough: to agree

**Starter draft for team review.** Questions and outputs below are suggestions based on the chapter outline, not agreed assignments or conclusions. Use the [planning guide](./README.md) and [full template](./TEMPLATE.md) to develop this plan in 45–60 minutes.

## 1. Purpose and intended outcome

Help platform teams choose routing and session-management patterns for agent-to-agent, agent-to-tool, and model interactions without assuming inference routing solves all agent traffic requirements.

Proposed minimum useful outcome for v0.1: a routing responsibility and decision table, with long-lived session and disconnect behavior illustrated by a walkthrough.

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| To identify | Gateway API, agent transports, proxies, or platform operations | Investigate, draft, review, or upstream liaison | Not yet identified |

- Missing perspectives and recruitment action: agree outreach for gateway maintainers, protocol/session expertise, and end-user operators.
- How to join: contact the stream owner in `#initiative-distributed-agentic-systems` with your experience or proposed contribution.
- Suggested first task: trace one documented MCP connection lifecycle through a gateway; identify transport-version assumptions and failure questions.

## 3. Scope and priority questions

- In scope: agent and tool routing, inference routing integration, session-aware fan-out, streaming, multiplexing, and proxy behavior.
- Out of scope: choosing one proxy product or treating a proposed gateway primitive as already agreed.
- Assumptions and terminology: distinguish inference requests, agent/tool sessions, and transport-specific streaming behavior. Verify the chapter outline's transport terminology against selected specification versions.
- Single-cluster baseline: gateway, agent, and tool endpoints within a Kubernetes deployment.
- Distributed considerations: cross-cluster routing, session placement, reconnect, and control-plane reachability.
- Boundaries: agree protocol contracts with protocol, session recovery with state/runtime, and per-tool authorization and evidence with governance.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| Proposed v0.1 | Which routing needs are covered by current Gateway API and inference capabilities? | Existing-capability mapping | To agree |
| Proposed v0.1 | What session affinity and streaming behavior do selected agent protocols require? | Session and transport decision guidance | To agree |
| Proposed v0.1 | What happens when a connection or backend fails mid-task? | Failure and recovery walkthrough | To agree |
| Proposed v0.1 | Which needs fit existing extension contracts, and which remain demonstrably unmet? | Evidence for architecture choices and possible upstream work | To agree |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) | Inference availability; control and communication | Baseline for routing, resilience, and communication | Published March 23, 2026; assess agent traffic separately from inference |
| [Gateway API](https://gateway-api.sigs.k8s.io/) and [Inference Extension](https://gateway-api-inference-extension.sigs.k8s.io/) | Routing resources and inference integration | Assess existing primitives before proposing extensions | Select API and implementation versions; conformance and agent-session behavior need verification |
| [Chapter 03 outline](../03_Gateway_and_Data_Plane.md) | Session-aware architecture questions | Existing hypotheses to investigate, not conclusions | Framing explicitly under review |

## 5. Milestones and substantive work

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Confirm team, scope, and protocol versions | To agree | To agree | Proposed | Not yet created |
| First useful chapter draft and disconnect walkthrough | To agree | To agree | Proposed | Not yet created |
| Architecture and gap-analysis inputs | To agree | To agree | Proposed | Not yet created |
| Cross-stream review and integration | To agree | To agree | Proposed | Not yet created |

Create issues for substantive work once scoped; keep detailed tasks in those issues.

## 6. Current biweekly walkthrough — to schedule

- Changes or findings: none recorded in this starter.
- Suggested material: draw routing and session ownership in the [shared reference scenario](./README.md#proposed-shared-reference-scenario), then interrupt a connection.
- Feedback needed: which layer owns reconnect, replay, and authorization on resumed interactions?
- Blocker / help needed: confirm contributor and protocol/governance review coverage.
- Next output, contributor, and checkpoint: agree at kickoff; link meeting notes, issues, and PRs here.

## 7. Findings and integration — develop progressively

No findings established here. Use the [template's findings table and handoff prompts](./TEMPLATE.md#7-findings-and-integration--develop-progressively). Link gateway responsibilities, session flows, and decision points to [chapter 07](../07_Reference_Architecture.md). Link evidenced limitations and candidate upstream destinations to [chapter 08](../08_Gap_Analysis_and_Recommendations.md), without assuming a new primitive is necessary.

## 8. Review and readiness — use later

Use the [common readiness checklist](./TEMPLATE.md#8-review-and-readiness--use-later). Record review feedback and deferred questions here or in linked issues. This starter is not ready for integration as technical guidance.
