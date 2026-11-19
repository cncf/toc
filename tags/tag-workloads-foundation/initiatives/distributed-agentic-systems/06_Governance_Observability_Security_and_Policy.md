# 6. Governance, observability, security and policy

Status: outline. Owner: Elankumaran Srinivasan.

- OpenTelemetry GenAI semantic conventions and spans; metrics, events, logs and traces; baggage
- LLM cost, token and latency tracking. Gap: these numbers are not comparable across implementations today, so an operator cannot bound spend from telemetry (see 6b)
- Policy and compliance as code, with OSCAL Compass as one implementation (Anca Sailer)
- 6a. Verifiable per-call evidence (subsection, proposed by Sankalp Gilda): independently verifiable records of gateway behavior that third parties can audit; vendor neutral, anchored on in-toto and the OWASP agent control work
- 6b. Telemetry schema and cost (subsection, proposed by Elankumaran Srinivasan): a cost figure carries its unit and its source, and a declared budget is checked after the run from its record; vendor neutral, anchored on the OpenTelemetry GenAI conventions and the FinOps FOCUS specification

## 6b. Telemetry schema and cost

Status: outline. Author: Elankumaran Srinivasan.

An amount without its unit and its source is not comparable. Budget is a separate question from cost: a limit that was configured is not a limit that was enforced.

Each requirement is a checklist item plus one line on how an implementer checks it.

- [ ] R1. A cost figure carries its unit and its source. Record a monetary amount only alongside an ISO 4217 currency code, and never record a token or credit count in a currency-denominated field. Declare whether the provider that served the call returned the amount or it was computed locally. No attribute carries the pricing data behind a computed amount, so that is declared absent rather than defaulted. Check: aggregating two providers' costs fails closed when a currency is absent, a token count occupies a currency field, or the source of a computed amount is undeclared; and where two layers record one operation, each record identifies the instrumentation that emitted it, so a consumer can de-duplicate before summing.
- [ ] R2. A declared budget is checked after the run from its record. Record the declared limit with its unit, an observed count whose vantage is declared (6a R2) and which comes from a component the enforcer does not control (6a R4), and the declared mechanism of the enforcer that produced this record. If the enforcer reports its own counter, record it with its unit; a counter without a unit is a missing fact. Treat the mechanism as a claim the record lets a reader falsify. Check: a limit that held is distinguished from one merely requested by a retained record of the requests actually made, and not by the enforcer's own counter; a conformance check exercises a server that ignores the stop request, since one that honours it makes the mechanisms indistinguishable; and an undeterminable mechanism is unclassifiable, per 6a R3.

Patterns: the amount is returned by the provider in the response, computed at a gateway against a price table whose revision is recorded, or reconciled after the fact against a billing export; each record declares which.

Requirement numbering is scoped to this subsection. A cross-framework count of consumed iterations stays out of scope, because the implementations measured bind their limits to different native units, so the resulting counts are not comparable.

How an implementer checks it today: OpenTelemetry's GenAI [conformance report](https://github.com/open-telemetry/semantic-conventions-conformance/blob/6807b1d5e676/docs/data/conformance.json) records which conventions each instrumented project emits. At revision `6807b1d5e676` it covers 57 targets, of which 3 emit a cost figure at all, under 12 attribute and 2 metric names spanning four namespaces. The registry defines no cost attribute at all, so every one of those names is an implementation's own choice; one of them sits inside `gen_ai.usage.*`, a namespace the registry does own.

Neither requirement binds to an attribute name; any attribute satisfying them would be defined in the OpenTelemetry GenAI SIG. None of the three emitters carries a currency or a source beside the amount. Both cost metrics do declare a currency, as the metric instrument's unit, and they spell it differently: one `USD`, the other `{USD}`. Neither spelling is conventional, because the metric-unit guidance says nothing about currency and uses curly braces to annotate dimensionless counts. That is read from the pinned library versions rather than from the report, which records no unit. An instrument unit is readable, but it is fixed per instrument, so it cannot vary by provider or by request.

Builds on: the OpenTelemetry GenAI semantic conventions and their conformance report, and the FinOps FOCUS specification, whose v1.0 `BillingCurrency` column already shows the shape R1 needs: it is required, non-null and constrained to a currency-code format. FOCUS governs billing exports and not runtime telemetry, so the constraint is borrowed and not inherited. The record a cost points back to is 6a R1; 6b defines no second binding.

Open: which fields are shared with 6a so one value is not specified twice, whether a consumed-iteration count can be made comparable at all without a declared unit, and where a locally computed amount's price basis is retained.
