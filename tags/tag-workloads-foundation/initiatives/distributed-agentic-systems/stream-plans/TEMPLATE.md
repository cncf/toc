# Stream plan: <name>

Technical chapter: <link>
Stream owner: <name / handle, or open>
Team: <names / handles; distinguish confirmed from invited>
Status: <forming | scoping | drafting | reviewing | ready for integration>
Last updated: <YYYY-MM-DD>
Next walkthrough: <date / to agree>

Complete sections 1–4 and an initial milestone and next activity in sections 5–6 together in 45–60 minutes. Unknowns are welcome; identify how to resolve them. Develop sections 7–8 over subsequent walkthroughs. See the [planning guide](./README.md) and [illustrative example](./EXAMPLE-state-memory.md).

## 1. Purpose and intended outcome

Describe the operational problem, intended readers, and what they should be able to decide after reading the chapter. What is different for agentic workloads rather than distributed systems generally?

Minimum useful outcome for v0.1: <smallest coherent contribution, not an exhaustive topic inventory>

## 2. Team and expertise

| Person / handle | Expertise / perspective | Intended contribution | Participation |
|---|---|---|---|
| <name> | <domain or practitioner experience> | <investigate, draft, review, upstream liaison> | <confirmed / invited> |

- Missing expertise or perspectives: <include practitioner and cross-stream review needs>
- Recruitment action: <who will approach whom, by which checkpoint>
- How to join: <contact / channel / issue; use public handles, not private contact details>
- First task for a new contributor: <bounded, useful task with context or issue link>

Do not treat a suggested expert as a confirmed contributor or assign work without agreement.

## 3. Scope and priority questions

- In scope: <capabilities and problems>
- Out of scope: <explicit non-goals>
- Assumptions and terminology: <definitions or links>
- Single-cluster baseline: <deployment assumptions>
- Multi-cluster, edge, or disconnected considerations: <where relevant; otherwise explain why>
- Boundaries with other streams: <overlap, proposed lead, and coordination needed>

Start with three to five questions. Separate v0.1 priorities from later investigations.

| Priority | Question to answer | Intended output | Contributor |
|---|---|---|---|
| <v0.1 / later> | <operational decision or uncertainty> | <pattern, comparison, guidance, finding, gap> | <agreed name / help needed> |

## 4. Starting knowledge and sources

| Source / project | Relevant section or capability | Why it matters / question informed | Version / date and evidence limitations |
|---|---|---|---|
| <link> | <baseline guidance or existing capability> | <connection to priority question> | <documented, demonstrated, production experience, unverified; details> |

- Which [Cloud Native Agentic Standards](https://www.cncf.io/blog/2026/03/23/cloud-native-agentic-standards/) recommendations apply, and what are we adding?
- What do Kubernetes and existing projects already provide? What remains uncertain or insufficient?
- Where do sources disagree? Which practitioner experience or upstream contact could help?

Prefer primary sources and record versions for changing specifications. Organize around capabilities and patterns, not project showcases. Citation does not imply endorsement; no chapter should be based on one project.

## 5. Milestones and substantive work

| Milestone / output | Responsible contributor(s) | Target checkpoint | Status | Issue / PR |
|---|---|---|---|---|
| Initial scope and team plan | <agree> | <agree> | <status> | <link / not yet created> |
| First useful chapter draft | <agree> | <agree> | <status> | <link / not yet created> |
| Architecture and gap-analysis inputs | <agree> | <agree> | <status> | <link / not yet created> |
| Review and integration | <agree> | <agree> | <status> | <link / not yet created> |

Link substantive work to issues, with investigation approach and expected evidence in the issue. Keep detailed tasks there. Agree targets with contributors; mark deferred work and the reason explicitly.

## 6. Current biweekly walkthrough — <date>

- What changed or did we learn? <summary>
- What concrete material can we walk through? <diagram, comparison, scenario, pattern, draft, or link>
- Which recommendation, assumption, or boundary needs feedback? <question>
- What is blocked, and what help is needed? <request>
- What will we bring to the next meeting? <output, agreed contributor, checkpoint>
- Related issues, PRs, or meeting notes: <links>

Use the [shared reference scenario](./README.md#proposed-shared-reference-scenario) where helpful, adding stream-specific examples as needed. A walkthrough is not necessarily a demo or implementation validation. Record the evidence actually available and its limitations. Keep previous discussion in linked meeting notes or issues, not an expanding log here.

## 7. Findings and integration — develop progressively

| Finding / pattern / open gap | Evidence and limitations | Related streams / coordination needed | Destination / next action |
|---|---|---|---|
| <summary> | <source, practitioner review, walkthrough, experiment> | <streams and proposed lead; agree boundaries> | <chapter / issue / upstream> |

Keep technical findings in the chapter; link rather than duplicate them here. For each substantive pattern, use: **Problem → Context → Approach → Alternatives and trade-offs → Failure behavior → Evidence → Related streams**.

### Reference architecture handoff

For [chapter 07](../07_Reference_Architecture.md), link component roles, interfaces and flows (protocol, identity, state, telemetry), deployment choices, decision points, and failure behavior. Identify dependencies that need another stream's review.

### Gap-analysis handoff

For [chapter 08](../08_Gap_Analysis_and_Recommendations.md), identify the unmet requirement, evidence, existing capability and its limitation, and whether the gap is documentation, implementation, interoperability, or specification. Suggest an upstream destination and next action without implying upstream acceptance.

## 8. Review and readiness — use later

- [ ] Priority questions are answered or explicitly deferred with a reason.
- [ ] Scope, terminology, and boundaries are reviewed with affected streams.
- [ ] Existing capabilities, proposed guidance, and possible standards work are distinguished.
- [ ] Sources, versions, assumptions, and evidence limitations are visible.
- [ ] Security, tenancy, observability, recovery, and distributed deployment implications are addressed or explicitly deferred.
- [ ] Relevant technical and practitioner feedback is recorded, including unresolved objections.
- [ ] Chapter 07 and 08 inputs are linked, or the absence of an identified gap is explained.
- [ ] Follow-up work and review contributions are linked to issues or PRs.

This checklist supports review; the [charter's contribution and decision process](../charter.md#terms-of-reference) still applies.
