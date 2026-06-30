# Security and privacy considerations

Runtime Conditions Profiles are designed to describe requirements without embedding fulfillment data.

Profiles must not contain secret values, credentials, customer data, protected data, or concrete target-environment values. Environment configuration entries name workload-facing inputs such as `POSTGRES_PASSWORD` or `PAYMENTS_API_TOKEN`, but they do not carry the corresponding value. Downstream adapters are responsible for mapping those requested inputs to Secrets, ConfigMaps, or other target-environment artifacts.

Some things to consider:

- Profiles can still reveal useful information about a workload's expected behavior. For example, a profile may disclose that a workload calls a payments API, depends on Redis, or needs an object store.
- Generators should avoid emitting values discovered from local developer environments.
- Validators should reject fields that violate the profile shape, unresolved vocabulary, and extension JSON Schema rules.
- Adapters should preserve the profile's distinction between workload requirements and target-environment fulfillment choices.
