# Runtime Conditions Profile Specification (Draft)

## Status

**Draft — Request for Comments**

This document is an early working draft of the Runtime Conditions Profile specification.

It is intentionally not marked with a version and is being published solely to solicit early feedback from the community.

This draft is expected to evolve significantly based on review and discussion before a stable version is tagged.

## Request for Feedback

The authors are particularly interested in feedback on:

- Core Condition model structure
- Extension model design
- Validation behavior and layering
- Namespacing approach
- Overall scope boundaries

Early architectural feedback is strongly encouraged.

---

# 1. Purpose

The Runtime Conditions Profile provides a **portable declaration of required external runtime capabilities** needed for an application workload to function successfully.

These capabilities may include:

- HTTP services
- Relational databases
- Caches
- Vendor-defined integration services

The Runtime Conditions Profile:

- **SHOULD be generated automatically when possible**
- **MAY be authored manually when automated generation is not feasible**
- **MUST remain valid regardless of generation method**
- **MUST remain implementation-neutral**
- **MUST remain infrastructure-agnostic**

The profile defines **requirements**, not implementations.

---

# 2. Scope

This specification defines a portable format for describing the external capabilities that an application workload depends on in order to function properly. These dependencies represent integrations with services that exist outside the workload itself, such as HTTP APIs, databases, caches, and message systems.

The Runtime Conditions Profile models each dependency as an independent  requirement that describes the expected interface characteristics needed to interact with an external system. The specification focuses on describing what capabilities must be present, without describing how those capabilities are implemented or fulfilled.

This specification is limited to externally satisfied integrations and does not attempt to describe internal execution behavior, infrastructure configuration, deployment topology, or platform-specific provisioning. It also does not require or depend on any upstream observation system, although such systems may be used to generate Runtime Conditions Profiles.

---

# 3. Core Design Principles

## 3.1 Declarative

Profiles MUST be declarative documents describing what is required, not how to fulfill it.

A Runtime Conditions Profile MUST be associated with a uniquely identifiable
workload and SHOULD correspond to a specific version of that workload. The profile version SHOULD align with the workload version.

A Runtime Conditions Profile MUST describe exactly one workload identity
and MUST NOT represent multiple unrelated workloads within a single profile.

## 3.2 Portable

Profiles SHOULD be portable across environments and platforms when expressed using only core specification vocabulary.

Profiles that use extension-defined vocabulary MAY introduce platform-specific or vendor-specific semantics. Such profiles remain portable to the extent that the required extensions are available.

## 3.3 Implementation-Neutral

Profiles MUST describe required capabilities without prescribing how those capabilities are implemented or provisioned.

Core specification vocabulary MUST remain vendor-neutral and MUST NOT encode assumptions about specific infrastructure implementations.

Vendor-specific or platform-specific identifiers MAY be used only when introduced through declared extensions.

Profiles MUST NOT encode:

- Infrastructure configuration details
- Deployment topology
- Resource sizing
- Geographic placement
- Provider-specific provisioning instructions

## 3.4 Extensible

Profiles MAY include extension-defined vocabulary to describe capabilities beyond those defined in the core specification.

Profiles that use extension-defined vocabulary MUST identify the extensions on which that vocabulary depends.

Use of extensions MUST NOT alter or redefine the meaning of core specification vocabulary.

## 3.5 Deterministically Validatable

Profiles MUST adhere to the structural and semantic validation rules defined by the core specification.

Profiles that reference extension-defined vocabulary MUST also adhere to the validation rules defined by those extensions.

---

# 4. Runtime Conditions Profile Structure

A Runtime Conditions Profile defines a collection of independent runtime Conditions.

Examples in this specification are expressed using YAML for readability. The data model defined by this specification is serialization-neutral and MAY be represented using YAML, JSON, or other compatible formats.

## Top-Level Structure

```yaml
## Top-Level Structure

apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: example-profile

workload:
  uri: https://github.com/example-org/example-service
  version: v1.2.3

extensions:
  - core
  - aws.runtime/v1alpha1

conditions:
  - name: primary-db
    kind: datastore
    interface:
      type: relational

  - name: payments-api
    kind: api
    interface:
      type: http
      operations:
        - method: POST
          path: /charge
```

---

# 5. Condition Model

Each Condition represents an **independent required runtime dependency**.

## Condition Fields


| Field       | Required | Description                                       |
| ----------- | -------- | ------------------------------------------------- |
| `kind`      | YES      | Required capability classification                |
| `interface` | YES      | Interface definition required for matching        |
| `name`      | NO       | Unique identifier within profile                  |
| `optional`  | NO       | Whether the Condition is optional. Defaults to `false` |


## Optional Conditions

The `optional` field is a boolean that indicates whether a Condition is optional.

When `optional` is omitted, it MUST be understood as `false`, meaning the Condition is required. It MAY be defined explicitly, but only setting `optional: true` changes downstream behavior.

The logic for determining what makes a Condition optional is beyond the scope of this specification. It is the responsibility of both the developer and the platform team to decide how optional integration Conditions are handled.

---

# 6. Core Condition Kinds

The core specification defines a set of **capability classes**, referred to as **Kinds**, that represent common categories of externally satisfied runtime dependencies.

Each Condition MUST declare exactly one `kind`.

Kinds represent **broad capability families**, not specific technologies.

The following core kinds are defined:

- `api` — External service integrations such as APIs
- `datastore` — Persistent data storage systems
- `cache` — Volatile data storage optimized for fast access

These kinds are intentionally broad so that multiple interaction models or implementation families can be expressed within the same capability class through the `interface` block.

---

# 7. Interface Model

Each Condition MUST define an `interface` block describing how the workload interacts with the declared capability.

The `interface` block defines:

- The **interaction model**
- The **matching characteristics**
- Any **additional details required for fulfillment matching**

## Interface Structure

```yaml
interface:
  type: <interface-type>
```

The `type` field identifies the interaction model associated with the declared `kind`.

Additional fields MAY be defined within `interface` depending on the declared `kind` and `interface.type`.

Interface definitions are validated based on:

- The declared `kind`
- The declared `interface.type`
- Core validation rules
- Extension validation rules, when applicable

---

# 8. Core Interface Types

This section describes the **structure and purpose** of core interface types. The set of currently supported interface types is defined in the validation section.

---

## 8.1 API Interface

API interfaces describe callable external services such as APIs.

API interfaces typically define:

- Request methods
- Request paths
- Optional request schemas
- Optional response schemas
- An optional reference to an external API specification document

Example:

```yaml
kind: api
interface:
  type: http
  spec:
    format: openapi
    uri: https://github.com/example-org/example-service/openapi.yaml
    version: ^1.2.0
  operations:
    - method: POST
      path: /charge
      requestBodySchema:
        amount: integer
        currency: string
      responseSchema:
        id: string
        status: string
```

### API Interface Fields


| Field        | Required | Description                                         |
| ------------ | -------- | --------------------------------------------------- |
| `type`       | YES      | Identifies the API interaction model                |
| `spec`       | NO       | Reference to an external API specification document |
| `operations` | NO       | Explicit list of operations the workload depends on |


Allowed values for `interface.type` are defined in Section 9.2.

At least one of `spec` or `operations` MUST be present. Both MAY be declared together. When both are present and disagree, the `operations` declaration takes precedence over `spec`.

### API Specification Fields

The `spec` block references an external API specification document describing the API the workload integrates with.

```yaml
spec:
  format: openapi
  uri: https://github.com/example-org/example-service/openapi.yaml
  version: ^1.2.0
```


| Field     | Required | Description                                            |
| --------- | -------- | ------------------------------------------------------ |
| `format`  | YES      | API specification format. Only `openapi` is currently supported  |
| `uri`     | YES      | Location of the external specification document        |
| `version` | NO       | Required version of the referenced document, as an exact version or version constraint |


Only the OpenAPI specification is currently supported. Other API specification formats exist and MAY be supported in future versions.

### Version Constraint Syntax

`interface.spec.version` declares which version of the referenced specification document the workload requires. It accepts either an exact version or a version constraint expression, using the operators commonly found in dependency manifests such as `package.json`.

Versions MUST follow Semantic Versioning, expressed as `MAJOR.MINOR.PATCH`.

The following constraint operators are supported:


| Syntax     | Name                | Meaning                                          |
| ---------- | ------------------- | ------------------------------------------------ |
| `1.2.3`    | Exact               | Matches only `1.2.3`                             |
| `=1.2.3`   | Exact               | Equivalent to `1.2.3`                            |
| `>1.2.3`   | Greater than        | Matches any version higher than `1.2.3`          |
| `>=1.2.3`  | Minimum             | Matches `1.2.3` and any higher version           |
| `<1.2.3`   | Less than           | Matches any version lower than `1.2.3`           |
| `<=1.2.3`  | Maximum             | Matches `1.2.3` and any lower version            |
| `^1.2.3`   | Compatible (caret)  | Matches `>=1.2.3` and `<2.0.0` (no major change) |
| `~1.2.3`   | Approximate (tilde) | Matches `>=1.2.3` and `<1.3.0` (no minor change) |


When `version` is omitted, no version constraint is applied and any version of the referenced document is acceptable.

Validation rules:

- If `version` is present, it MUST be a valid Semantic Versioning version or a supported constraint expression
- Constraint operators MUST be drawn from the supported set above

### Operation Fields


| Field               | Required | Description         |
| ------------------- | -------- | ------------------- |
| `method`            | YES      | HTTP method         |
| `path`              | YES      | Request path        |
| `requestBodySchema` | NO       | Minimum required request body fields and their types |
| `responseSchema`    | NO       | Minimum required response fields and their types     |


### Schema Fields

`requestBodySchema` and `responseSchema` describe the data structures an operation depends on. Each is expressed as a map whose keys are field names and whose values declare the JSON Schema type of each field.

The declared fields represent the **minimum set of fields that MUST be present** in the external API. The external service MAY expose additional fields beyond those declared; only the declared fields participate in matching.

Because the structure is expressed using field names and JSON Schema types rather than a specific API specification format, a Condition can be matched programmatically against any specification the external service publishes — including, but not limited to, the OpenAPI document referenced through `interface.spec`.

A type declaration is one of:

- A JSON Schema type keyword: `string`, `number`, `integer`, `boolean`, or `null`
- A nested object, declared by mapping field names to further type declarations
- An array, declared as a single-element list containing the element type declaration

Example:

```yaml
operations:
  - method: POST
    path: /charge
    requestBodySchema:
      amount: integer
      currency: string
      customer:
        id: string
        email: string
      lineItems:
        - sku: string
          quantity: integer
    responseSchema:
      id: string
      status: string
      paid: boolean
```

### Validation Expectations

- At least one of `spec` or `operations` MUST be present
- If `operations` is present, it MUST be non-empty
- `method` MUST be a valid HTTP method
- `path` MUST be a non-empty string
- If present, `requestBodySchema` and `responseSchema` MUST be maps whose fields declare supported JSON Schema types
- Declared schema fields represent the minimum required set of fields and MAY be a subset of the fields exposed by the external API
- If `spec` is present, `spec.format` and `spec.uri` MUST be present
- `spec.format` MUST be `openapi`
- If `spec.version` is present, it MUST be a valid semantic version or supported version constraint expression

---

## 8.2 Datastore Interface

Datastore interfaces describe persistent storage systems.

Datastore interfaces identify the storage model used by the workload and MAY include additional matching details about the datastore engine.

Example:

```yaml
kind: datastore
interface:
  type: relational
  engine: postgres
```

### Datastore Interface Fields


| Field    | Required | Description               |
| -------- | -------- | ------------------------- |
| `type`   | YES      | Datastore interface type  |
| `engine` | NO       | Specific datastore engine |


If `engine` is provided, it MUST be valid for the declared datastore type.

Allowed values for `interface.type` and `engine` are defined in Section 9.2.

---

## 8.3 Cache Interface

Cache interfaces describe volatile key/value storage systems.

Example:

```yaml
kind: cache
interface:
  type: key_value
  engine: redis
```

### Cache Interface Fields


| Field    | Required | Description             |
| -------- | -------- | ----------------------- |
| `type`   | YES      | Cache interface type    |
| `engine` | NO       | Specific caching engine |


If `engine` is provided, it MUST be valid for the declared cache type.

Allowed values for `interface.type` and `engine` are defined in Section 9.2.

---

# 9. Core Validation Rules

Validation ensures that Conditions are structurally correct and semantically consistent.

Validation occurs in multiple phases.

---

## 9.1 Structural Validation

A Condition is invalid if:

- `kind` is missing
- `interface` is missing
- `interface.type` is missing

If a `name` field is provided, it MUST be unique within the profile.

---

## 9.2 Kind-to-Interface Type Validation

Each `kind` supports a defined set of valid `interface.type` values.

### API

Allowed interface types:

- `http`

Additional validation rules:

- At least one of `spec` or `operations` MUST be present
- If `operations` is present, it MUST NOT be empty
- If `spec` is present, `spec.format` and `spec.uri` MUST be present
- `spec.format` MUST be `openapi`
- If `spec.version` is present, it MUST be a valid semantic version or supported version constraint expression

Allowed HTTP Methods:

- GET
- HEAD
- POST
- PUT
- PATCH
- DELETE
- OPTIONS
- TRACE

---

### Datastore

Allowed interface types:

- `relational`
- `document`

Allowed `engine` values for `type: relational`:

- `postgres`
- `mysql`
- `mariadb`
- `sqlserver`
- `oracle`
- `sqlite`

Allowed `engine` values for `type: document`:

- `mongodb`
- `couchbase`

Additional validation rules:

- `engine` is OPTIONAL
- If `engine` is present, it MUST be valid for the declared datastore type

---

### Cache

Allowed interface types:

- `key_value`

Allowed `engine` values for `type: key_value`:

- `redis`
- `memcached`

Additional validation rules:

- `engine` is OPTIONAL
- If `engine` is present, it MUST be valid for the declared cache type

---

## 9.3 Invalid Condition Examples

Invalid datastore engine for relational type:

```yaml
kind: datastore
interface:
  type: relational
  engine: mongodb
```

Invalid interface type for cache:

```yaml
kind: cache
interface:
  type: relational
```

Invalid api definition:

```yaml
kind: api
interface:
  type: http
  operations: []
```

---

# 10. Extension Model

The Runtime Conditions Profile supports extension-defined vocabulary.

Extensions allow:

- New kinds
- New interface types
- New interface fields
- Additional validation rules
- Additional allowed values for existing fields where semantically compatible

Extensions MUST NOT redefine core semantics incompatibly.

---

# 11. Extension Declaration

Profiles that reference extension-defined vocabulary MUST identify those extensions.

```yaml
extensions:
  - core
  - aws.runtime/v1alpha1
  - redis.compat/v1
```

---

# 12. Extension Definition Structure

Extensions are defined as independent artifacts.

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: ValidationExtensionDefinition

metadata:
  name: aws.runtime
  version: v1alpha1

spec:

  kinds:
    - name: aws.object_store

  interfaceTypes:
    - name: object_store

  typeExtensions:
    - targetKind: cache
      targetType: key_value
      addEngines:
        - valkey

  validationRules:
    - id: cache-valkey
      appliesToKind: cache
      rule: engine in ["redis","memcached","valkey"]
```

---

# 13. Extension Capabilities

Extensions MAY:


| Action             | Description                               |
| ------------------ | ----------------------------------------- |
| Add Kind           | Introduce new namespaced kind             |
| Add Interface Type | Define new interface types                |
| Add Fields         | Extend interface schema                   |
| Add Rules          | Add semantic validation                   |
| Add Allowed Values | Extend allowed values for existing fields |


---

# 14. Extension Compatibility Rules

Extensions MUST:

- Use namespaced identifiers where they introduce new vocabulary
- Preserve core semantics
- Not redefine core kinds incompatibly
- Not invalidate core-valid documents

Extensions MAY add new allowed values for existing fields only where those values are semantically compatible with the core meaning of the declared `kind` and `interface.type`.

---

# 15. Namespacing Requirements

Extension-defined vocabulary MUST use namespaced identifiers to avoid collisions with core or other extension-defined elements.

Namespacing applies to:

- Extension-defined `kind` values
- Extension-defined `interface.type` values
- Extension-defined field names (when applicable)

Namespacing typically uses a prefix that identifies the originating organization or vendor.

Examples of valid namespaced identifiers:

```yaml
kind: aws.object_store
interface:
  type: aws.s3
```

```yaml
kind: api
interface:
  type: acme.soap
```

Values defined within existing fields, such as `engine`, do not require namespacing unless necessary to prevent
collisions.

---

# 16. Unknown Extension Handling

Profiles MAY reference extension-defined vocabulary through the `extensions` declaration.

If a profile references extension-defined vocabulary that cannot be resolved through its declared extensions, the profile MUST be considered invalid.

This includes cases where:

- A `kind` value is not defined in the core specification or in any declared extension
- An `interface.type` value is not defined in the core specification or in any declared extension
- A field defined by an extension is used without the corresponding extension being declared
- A declared extension cannot be located or resolved during validation

Validation systems MUST reject Conditions that rely on unknown or unresolved extension-defined vocabulary.

---

# 17. Validation Layers

Validation occurs in the following order:

1. Core structural validation
2. Core semantic validation
3. Extension resolution
4. Extension validation

---

# 18. Example: Core-Only Profile

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: checkout-service

workload:
  uri: https://github.com/example-org/checkout-service
  version: v1.2.3

extensions:
  - core

conditions:

  - kind: datastore
    interface:
      type: relational
      engine: postgres

  - kind: cache
    interface:
      type: key_value
      engine: redis

  - kind: api
    interface:
      type: http
      operations:
        - method: POST
          path: /charge
```

---

# 19. Example: AWS Extension Profile

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: storage-enabled

workload:
  uri: https://github.com/example-org/storage-service
  version: v1.4.0

extensions:
  - core
  - aws.runtime/v1alpha1

conditions:

  - kind: aws.object_store
    interface:
      type: aws.s3
```

---

# 20. Summary

The Runtime Conditions Profile defines:

- A portable dependency declaration format
- A structured interface typing model
- Deterministic validation behavior
- A declarative extension mechanism
- Vendor-neutral core semantics

This provides a foundation for reliable downstream capability matching while preserving ecosystem flexibility.