# Stream plan: governance, observability, security and policy

Technical chapter: [Chapter 06](../06_Governance_Observability_Security_and_Policy.md)
Stream owner: Elankumaran Srinivasan, as recorded in the [initiative README](../README.md#workstreams)
Team: to identify; no additional participation confirmed here
Status: forming
Last updated: 2026-11-05
Next walkthrough: to agree

**Starter draft for team review.** Questions and outputs below are suggestions based on the chapter outline, not agreed assignments or conclusions. Use the [planning guide](./README.md) and [full template](./TEMPLATE.md) to develop this plan in 45–60 minutes. Coordinate existing subsection proposals with their contributors rather than assuming new commitments.

## 1. Purpose and intended outcome

Help platform teams enforce and observe agent identity, delegation, tenancy, policy, and operational limits, with evidence that supports debugging and audit without unnecessarily exposing sensitive context.

Proposed minimum useful outcome for v0.1: cross-cutting control responsibilities, telemetry and cost guidance with explicit limitations, and an auditable-action walkthrough. Prioritize with the team so the scope remains deliverable.

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| To identify | OpenTelemetry, workload identity, authorization, policy, evaluation, audit, or operations | Investigate, draft, review, or upstream liaison | Not yet identified |

- Missing perspectives and recruitment action: agree outreach for security practitioners, end-user operators, telemetry maintainers, and existing policy/evidence subsection contributors.
- How to join: contact the stream owner in `#initiative-distributed-agentic-systems` with your experience or proposed contribution.
- Suggested first task: map identity, authorization, and trace context across one delegated tool call, noting sensitive data and evidence limitations.

## 3. Scope and priority questions

- In scope: GenAI telemetry and MELT, cost/token/latency measurements, policy and compliance as code, verifiable per-call evidence, and cross-stream security implications.
- Out of scope: vendor-specific controls, legal compliance guarantees, or treating telemetry alone as proof of compliant or correct behavior.
- Assumptions and terminology: distinguish user/workload identity, delegated authority, evaluation, operational telemetry, and independently verifiable evidence.
- Single-cluster baseline: explicit trust and tenancy boundaries across agents, tools, gateways, and stores.
- Distributed considerations: identity federation, disconnected authorization behavior, policy distribution, and evidence/telemetry continuity across clusters.
- Boundaries: agree enforcement and instrumentation responsibilities with all streams; avoid leaving cross-cutting controls unowned.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| Proposed v0.1 | Where are identity, delegation, approval, and least-privilege policies enforced? | Cross-stream control responsibility mapping | To agree |
| Proposed v0.1 | Which telemetry supports task diagnosis and comparable cost measurements, and with what limits? | Semantic-convention and measurement guidance | To agree |
| Proposed v0.1 | What evidence supports auditing a delegated action without exposing unnecessary sensitive data? | Evidence pattern, trust assumptions, and trade-offs | To agree |
| Proposed v0.1 | How are policy/evaluation decisions and failure behavior made explicit and reviewable? | Governance decision guidance | To agree |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) | Observability, governance, and security | Baseline to operationalize, not restate | Published March 23, 2026; recommendations need concrete deployment context |
| [OpenTelemetry GenAI agent spans](https://opentelemetry.io/docs/specs/semconv/gen-ai/gen-ai-agent-spans/) | Agent instrumentation conventions | Evaluate current coverage and measurement consistency | Record selected semantic-convention version and stability; do not assume all implementations conform |
| [Chapter 06 outline](../06_Governance_Observability_Security_and_Policy.md) | Cost comparability, policy, and proposed per-call evidence | Coordinate existing OSCAL Compass, in-toto, and OWASP-related investigations | Outline and proposals are not validated outcomes or endorsements; collect primary sources and independent perspectives |

## 5. Milestones and substantive work

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Confirm team, subsection boundaries, and v0.1 priorities | To agree | To agree | Proposed | Not yet created |
| First useful chapter draft and delegated-action walkthrough | To agree | To agree | Proposed | Not yet created |
| Architecture and gap-analysis inputs | To agree | To agree | Proposed | Not yet created |
| Cross-stream review and integration | To agree | To agree | Proposed | Not yet created |

Create issues for substantive work once scoped; keep detailed tasks in those issues.

## 6. Current biweekly walkthrough — to schedule

- Changes or findings: none recorded in this starter.
- Suggested material: trace identity, policy checks, human approval, telemetry, and evidence through the [shared reference scenario](./README.md#proposed-shared-reference-scenario).
- Feedback needed: what happens if authorization or telemetry infrastructure is unreachable, and which claims can the remaining evidence support?
- Blocker / help needed: confirm subsection participation and reviewer coverage from other streams.
- Next output, contributor, and checkpoint: agree at kickoff; link meeting notes, issues, and PRs here.

## 7. Findings and integration — develop progressively

No findings established here. Use the [template's findings table and handoff prompts](./TEMPLATE.md#7-findings-and-integration--develop-progressively). Link control boundaries and telemetry/evidence flows to [chapter 07](../07_Reference_Architecture.md). Provide supported measurement, policy, or interoperability limitations and candidate upstream destinations to [chapter 08](../08_Gap_Analysis_and_Recommendations.md).

## 8. Review and readiness — use later

Use the [common readiness checklist](./TEMPLATE.md#8-review-and-readiness--use-later). Record review feedback and deferred questions here or in linked issues. This starter is not ready for integration as technical guidance.
