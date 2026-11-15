# Stream plan: protocol and communication interoperability

Technical chapter: [Chapter 02](../02_Protocol_and_Communication_Interoperability.md)
Stream owner: open, as recorded in the [initiative README](../README.md#workstreams)
Team: to identify; no additional contributors confirmed here
Status: forming
Last updated: 2026-11-05
Next walkthrough: to agree

**Starter draft for team review.** Questions and outputs below are suggestions based on the chapter outline, not agreed assignments or conclusions. Use the [45–60 minute starter exercise](./README.md#start-together-in-4560-minutes) and [full template](./TEMPLATE.md) to develop this plan.

## 1. Purpose and intended outcome

Help platform teams understand agent-to-agent and agent-to-tool contracts, discovery, and validation in Kubernetes deployments.

Proposed minimum useful outcome for v0.1: capability and discovery guidance that distinguishes protocol responsibilities from cluster integration and identifies evidenced interoperability limitations.

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| To identify | MCP, A2A, AP2, discovery, or operator experience | Scope, investigate, draft, review, or upstream liaison | Not yet identified |

- Missing perspectives and recruitment action: agree an owner and outreach for protocol maintainers, registry/identity expertise, and practitioners.
- How to join: volunteer in `#initiative-distributed-agentic-systems`; the stream owner is currently open.
- Suggested first task: compare discovery metadata and validation responsibilities in two primary specifications; record versions and unresolved questions.

## 3. Scope and priority questions

- In scope: MCP, A2A, AP2, discovery and naming, dynamic registration, filtering, and schema validation.
- Out of scope: implementing a new protocol or selecting a preferred framework.
- Assumptions and terminology: define agent, tool, capability, registry, and discovery before comparing mechanisms.
- Single-cluster baseline: distinguish Kubernetes endpoint discovery from agent capability discovery.
- Distributed considerations: identify cross-cluster naming, trust, and unavailable-registry behavior where applicable.
- Boundaries: coordinate transport/session behavior with gateway, lifecycle registration with runtime, and trust/delegation with governance. Agree leads together.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| Proposed v0.1 | Which problems do MCP, A2A, and AP2 address, and where do their contracts stop? | Capability and responsibility comparison | To agree |
| Proposed v0.1 | How are capabilities registered, discovered, and kept current? | Discovery patterns and failure considerations | To agree |
| Proposed v0.1 | Where are payload schemas and protocol compatibility validated? | Validation decision guidance | To agree |
| Proposed v0.1 | Which interoperability needs remain unmet by existing specifications and integrations? | Evidence-backed findings or gaps | To agree |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) | Control and communication; security agent identity | Baseline for protocol, discovery, and trust guidance | Published March 23, 2026; verify current protocol claims against primary sources |
| [Chapter 02 outline](../02_Protocol_and_Communication_Interoperability.md) | Existing topic inventory | Starting scope, not a technical finding | Outline; select primary specification versions and project evidence during kickoff |

## 5. Milestones and substantive work

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Confirm owner, team, scope, and sources | To agree | To agree | Proposed | Not yet created |
| First useful chapter draft and scenario walkthrough | To agree | To agree | Proposed | Not yet created |
| Architecture and gap-analysis inputs | To agree | To agree | Proposed | Not yet created |
| Cross-stream review and integration | To agree | To agree | Proposed | Not yet created |

Create issues for substantive work once scoped; keep detailed tasks in those issues.

## 6. Current biweekly walkthrough — to schedule

- Changes or findings: none recorded in this starter.
- Suggested material: map discovery and validation onto the [shared reference scenario](./README.md#proposed-shared-reference-scenario). Use a separate example for payments if needed.
- Feedback needed: where do protocol guarantees end and deployment responsibilities begin?
- Blocker / help needed: stream owner and contributor recruitment.
- Next output, contributor, and checkpoint: agree at kickoff; link the meeting notes, issues, and PRs here.

## 7. Findings and integration — develop progressively

No findings established here. Use the [template's findings table and handoff prompts](./TEMPLATE.md#7-findings-and-integration--develop-progressively). Link component contracts and discovery flows to [chapter 07](../07_Reference_Architecture.md); route only evidence-backed unmet requirements to [chapter 08](../08_Gap_Analysis_and_Recommendations.md).

## 8. Review and readiness — use later

Use the [common readiness checklist](./TEMPLATE.md#8-review-and-readiness--use-later). Record review feedback and deferred questions here or in linked issues. This starter is not ready for integration as technical guidance.
