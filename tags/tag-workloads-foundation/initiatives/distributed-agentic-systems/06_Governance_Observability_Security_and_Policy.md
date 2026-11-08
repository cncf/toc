# 6. Governance, observability, security and policy

Status: outline. Owner: Elankumaran Srinivasan.

- OpenTelemetry GenAI semantic conventions and spans; metrics, events, logs and traces; baggage
- LLM cost, token and latency tracking. Gap: these numbers are not comparable across implementations today, so an operator cannot bound spend from telemetry
- Policy and compliance as code, with OSCAL Compass as one implementation (Anca Sailer)
- 6a. Verifiable per-call evidence (subsection, proposed by Sankalp Gilda): independently verifiable records of gateway behavior that third parties can audit; vendor neutral, anchored on in-toto and the OWASP agent control work

## 6a. Verifiable per-call evidence

Status: outline. Author: Sankalp Gilda.

Verifiable evidence is a per-call record that a reader who does not trust its producer can check from retained bytes. A policy decision and an observed effect are separate facts: an allowed call does not show that the effect happened.

Each requirement is a checklist item plus one line on how an implementer checks it.

- [ ] R1. One record per call. The record identifies the caller, the tool and the call attempt, binds the arguments and the evaluated policy by digest, and keeps the decision and the effect apart. Check: change an argument, the policy or the call association, and the reader detects it from the retained inputs.
- [ ] R2. Vantage declared per field. Each field says whether the platform observed it, the producer asserted it, or it was reconstructed. A signed assertion is still an assertion. Check: the reader can tell which party supplied each required value, and a reconstructed value never reads as observed.
- [ ] R3. Absence is not a pass. A missing required fact, an unavailable retained input or a conflicting binding gives an explicit incomplete or refused result; the verifier does not fill the gap with a default. Check: remove a required member or its retained source and the claim fails closed.
- [ ] R4. A denominator the producer cannot edit. Any completeness claim is bound to a population supplied by a separately controlled component (an admission controller, a workload registry, a dispatch record), with excluded paths and retention limits declared. Check: drop one call and completeness against that population changes; a producer-chosen subset cannot show full coverage.
- [ ] R5. Signature over canonical bytes. The signature covers one defined byte encoding that a verifier can recompute, as DSSE does with its pre-authentication encoding and COSE (RFC 9052) with its Sig_structure. Check: the verifier recomputes the signed bytes from the retained record and the signature verifies against them.

Patterns, each with its blind spots declared: interception at a gateway (see the gateway note in chapter 3), reconstruction from two retained snapshots, and binding at admission. Retry and fallback records keep distinct attempt identities and their link to the original call, which is where 6a meets resilience.

Conformance is scoped to a declared profile and an exact test-vector revision, with passing and failing controls for each requirement. Catching one bad vector does not show conformance on its own. A signature binds bytes and attribution; whether a call ran as recorded and whether its observer is trusted stay separate checks.

How an implementer checks it today: public test vectors for these requirements exist in [agent-evidence-vectors](https://github.com/probityai/agent-evidence-vectors). Other implementations are welcome beside it; the requirements define the subsection.

Builds on: in-toto and DSSE statement envelopes, IETF SCITT, the OWASP Agent Control Standard, and the OpenTelemetry GenAI conventions. The cost and budget requirements (unit and source beside each amount, a declared budget checked after the run) belong to the telemetry part of this chapter, and 6a's per-call binding is how a reported cost points back to its source.

Open: where Kubernetes keeps the separately controlled population for R4, how records bind across clusters, and which fields are shared with chapter 3 so the same value is not specified twice.
