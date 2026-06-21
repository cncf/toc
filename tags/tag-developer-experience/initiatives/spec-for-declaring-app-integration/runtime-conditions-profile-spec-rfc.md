# Runtime Conditions Profile Specification (Draft)

## Status

**Draft - Request for Comments**

This document defines the core Runtime Conditions Profile format.

First-party extension drafts define common integration vocabulary separately:

- `https://runtimeconditions.io/extensions/common-integrations:v1alpha1`
- `https://runtimeconditions.io/extensions/env-configuration:v1alpha1`

---

# 1. Purpose

A Runtime Conditions Profile declares the external runtime integrations required by one application workload.

The profile describes requirements. It does not describe implementations, provisioning actions, deployment topology, credentials, secret values, or concrete target-environment values.

---

# 2. Scope

This specification defines:

- The Runtime Conditions Profile document shape
- Workload identity fields
- Profile metadata labels
- The core Condition object shape
- Extension declaration and resolution rules
- Extension definition structure
- Validation layers
- Conformance requirements

This specification does not define concrete Condition vocabulary. Condition kinds, interface types, field values, and type-specific fields are defined by extensions.

Concerns specifically beyond the scope of this specification include:

- Infrastructure provisioning behavior
- Platform-specific resource models
- Deployment topology
- Runtime configuration values
- Secret material
- Internal workload behavior
- Observability or code-analysis mechanisms used to generate a profile

---

# 3. Profile Document

Profiles MAY be serialized as YAML or JSON.

Serialized profiles are invalid if:

- A mapping contains duplicate keys
- A required field is missing
- A required field is present with `null`
- A field has a type other than the type defined by this specification

Optional fields SHOULD be omitted when unused.

## 3.1 Complete Profile Example

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: checkout-service
  labels:
    owner.example.com/team: payments
    lifecycle.example.com/stage: production

workload:
  uri: https://github.com/example-org/checkout-service
  version: v1.2.3

extensions:
  - https://runtimeconditions.io/extensions/common-integrations:v1alpha1
  - https://runtimeconditions.io/extensions/env-configuration:v1alpha1

conditions:
  - name: primary-db
    kind: datastore
    interface:
      type: relational
      engine: postgres
    configuration:
      env:
        - property: hostname
          name: POSTGRES_HOST
        - property: port
          name: POSTGRES_PORT
        - property: database
          name: POSTGRES_DATABASE
        - property: username
          name: POSTGRES_USERNAME
        - property: password
          name: POSTGRES_PASSWORD
          sensitive: true

  - name: payments-api
    kind: api
    interface:
      type: http
      operations:
        - method: POST
          path: /charge
    configuration:
      env:
        - property: baseUrl
          name: PAYMENTS_API_URL
        - property: token
          name: PAYMENTS_API_TOKEN
          sensitive: true
```

## 3.2 Top-Level Fields

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `apiVersion` | string | YES | Runtime Conditions Profile API version |
| `kind` | string | YES | Document kind |
| `metadata` | object | YES | Profile metadata |
| `workload` | object | YES | Workload identity |
| `extensions` | array | YES | Extension identifiers required by the profile |
| `conditions` | array | YES | Runtime Conditions declared by the workload |

`apiVersion` MUST be:

```text
runtimeconditions.io/v1alpha1
```

`kind` MUST be:

```text
RuntimeConditionsProfile
```

`extensions` MAY be empty if and only if the profile uses no extension-defined vocabulary.

`conditions` MAY be empty if and only if the profile is intended to explicitly declare that the workload has no external runtime integration requirements.

## 3.3 Metadata

```yaml
metadata:
  name: checkout-service
  labels:
    owner.example.com/team: payments
    lifecycle.example.com/stage: production
```

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `name` | string | YES | Human-readable profile name |
| `labels` | object | NO | Machine-readable profile and workload classification labels |

`metadata.name` MUST be a non-empty string.

`metadata.labels`, when present, MUST be a string-to-string mapping.

Label keys MUST be non-empty strings.

Label values MUST be strings.

Label keys SHOULD be namespaced:

```text
<namespace>/<name>
```

Examples:

- `compliance.example.com/hipaa`
- `owner.example.com/team`
- `lifecycle.example.com/stage`
- `risk.example.com/criticality`

The `runtimeconditions.io/` label namespace is reserved for the Runtime Conditions core specification and first-party Runtime Conditions extensions.

Extensions MAY document label key conventions. Label keys are metadata. They are not Condition vocabulary and do not participate in extension resolution unless a future version of this specification defines otherwise.

Labels MAY be used for selection, filtering, policy, ownership, reporting, and lifecycle workflows.

Labels MUST NOT define runtime dependency requirements or change the meaning of Conditions.

Labels MUST NOT contain secrets, protected data, personal data, customer data, or concrete target-environment values.

## 3.4 Workload

```yaml
workload:
  uri: https://github.com/example-org/checkout-service
  version: v1.2.3
```

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `uri` | string | YES | Stable workload identifier |
| `version` | string | NO | Workload version described by the profile |

`workload.uri` MUST be a non-empty string.

`workload.version`, when present, MUST be a non-empty string.

A Runtime Conditions Profile MUST describe exactly one workload identity.

---

# 4. Condition Model

Each Condition represents one external runtime dependency requirement.

## 4.1 Condition Fields

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `name` | string | NO | Unique Condition name within the profile |
| `optional` | boolean | NO | Whether the Condition is optional. Defaults to `false` |
| `kind` | string | YES | Extension-defined integration classification |
| `interface` | object | YES | Workload-facing interface requirement |

Condition shape:

```yaml
conditions:
  - name: primary-db
    optional: false
    kind: datastore
    interface:
      type: relational
```

`conditions[].name`, when present, MUST be a non-empty string and MUST be unique within the profile.

`conditions[].optional`, when omitted, MUST be interpreted as `false`.

`conditions[].kind` MUST be a non-empty string and MUST be defined by exactly one resolved extension.

`conditions[].interface` MUST be an object.

Extensions MAY define additional Condition fields.

## 4.2 Interface Fields

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `type` | string | YES | Extension-defined interface type for the declared `kind` |

`conditions[].interface.type` MUST be a non-empty string and MUST be defined by exactly one resolved extension for the declared `kind`.

Extensions MAY define additional `interface` fields.

---

# 5. Extensions

Profiles that use extension-defined vocabulary MUST declare the required extensions.

First-party extensions are extensions. They are not core vocabulary and are not implicit.

Implementations MAY bundle support for first-party extensions, but generated profiles MUST still declare the vocabulary they use.

## 5.1 Extension Identifiers

Extension identifiers MUST have this form:

```text
<uri>:<version>
```

The version delimiter is the final colon in the identifier. The extension URI is everything before that delimiter. The version is everything after that delimiter.

The extension URI MUST be an absolute HTTP or HTTPS URI.

The version MUST be a non-empty string.

This specification does not define the syntax or semantics of extension versions. Version interpretation is defined by the extension author and the Adapter.

Examples:

- `https://runtimeconditions.io/extensions/common-integrations:v1alpha1`
- `https://runtimeconditions.io/extensions/env-configuration:v1alpha1`
- `https://extensions.example.com/runtimeconditions/aws-object-store:2026.06.0`

Extension identifiers are case-sensitive.

## 5.2 Extension Declarations

```yaml
extensions:
  - https://runtimeconditions.io/extensions/common-integrations:v1alpha1
  - https://runtimeconditions.io/extensions/env-configuration:v1alpha1
```

The `extensions` array MUST NOT contain duplicate extension identifiers.

Each `extensions` item MUST be a valid extension identifier.

Declared extensions MAY depend on other extensions.

Declared extensions and transitive dependencies form the resolved extension set.

Dependency resolution MUST be deterministic and MUST NOT depend on declaration order.

A profile is invalid if:

- A declared extension cannot be resolved
- A transitive dependency cannot be resolved
- Extension dependency resolution contains a cycle
- The resolved extension set contains a vocabulary definition conflict

---

# 6. Extension Definitions

Extensions are defined as independent artifacts.

## 6.1 Extension Definition Fields

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsExtensionDefinition

metadata:
  uri: https://runtimeconditions.io/extensions/common-integrations
  version: v1alpha1

spec:
  kinds: []
```

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `apiVersion` | string | YES | Runtime Conditions API version |
| `kind` | string | YES | MUST be `RuntimeConditionsExtensionDefinition` |
| `metadata` | object | YES | Extension identity |
| `spec` | object | YES | Extension vocabulary, dependencies, and validation schemas |

`metadata.uri` MUST identify the extension URI and MUST be an absolute HTTP or HTTPS URI.

`metadata.version` MUST identify the extension version.

`<metadata.uri>:<metadata.version>` MUST be a valid extension identifier.

## 6.2 Extension Spec Fields

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `dependencies` | array | NO | Exact extension identifiers required by this extension |
| `kinds` | array | NO | Condition kinds defined by this extension |
| `interfaceTypes` | array | NO | Interface types defined by this extension |
| `conditionFields` | array | NO | Condition-level fields defined by this extension |
| `interfaceFields` | array | NO | Interface-level fields defined by this extension |
| `fieldValues` | array | NO | Field values defined by this extension |
| `schemas` | array | NO | JSON Schema validation schemas defined by this extension |

An extension MUST define at least one vocabulary item or validation schema.

Each `dependencies` item MUST be an exact extension identifier.

## 6.3 Vocabulary Definition Fields

Each `kinds` entry MUST include:

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `name` | string | YES | Condition kind defined by the extension |

Each `interfaceTypes` entry MUST include:

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `name` | string | YES | Interface type defined by the extension |
| `targetKind` | string | YES | Condition kind for which the interface type is valid |

Each `conditionFields` entry MUST include:

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `name` | string | YES | Condition-level field defined by the extension |
| `appliesToKinds` | array | NO | Kinds to which the field applies |
| `appliesToInterfaceTypes` | array | NO | Interface types to which the field applies |
| `appliesToAllKinds` | boolean | NO | Whether the field may apply to all Conditions. Defaults to `false` |

Each `conditionFields` entry MUST declare either `appliesToKinds` or `appliesToAllKinds: true`.

Each `interfaceFields` entry MUST include:

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `name` | string | YES | Interface-level field defined by the extension |
| `targetKind` | string | YES | Condition kind for which the field is valid |
| `targetType` | string | YES | Interface type for which the field is valid |

Each `fieldValues` entry MUST include:

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `field` | string | YES | Field path whose values are being defined |
| `targetKind` | string | YES | Condition kind scope |
| `targetType` | string | NO | Interface type scope |
| `values` | array | YES | Non-empty list of defined values |

Each `fieldValues.values` entry MUST be unique within that `fieldValues` entry.

Each `schemas` entry MUST include:

| Field | Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| `id` | string | YES | Schema identifier unique within the extension |
| `description` | string | YES | Human-readable summary |
| `appliesToKind` | string | NO | Condition kind scope |
| `appliesToInterfaceType` | string | NO | Interface type scope |
| `schema` | object | YES | JSON Schema 2020-12 schema |

Each `schemas[].schema` object MUST be a valid JSON Schema using the JSON Schema 2020-12 dialect.

Each extension schema validates a Condition object after the profile has been converted to the JSON data model.

A schema applies to a Condition when all declared scope fields match the Condition. Omitted scope fields do not restrict applicability.

## 6.4 Extension Definition Example

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsExtensionDefinition

metadata:
  uri: https://aws.example.com/runtimeconditions/object-store
  version: v1alpha1

spec:
  kinds:
    - name: aws.object_store

  interfaceTypes:
    - name: aws.s3
      targetKind: aws.object_store

  interfaceFields:
    - name: bucketClass
      targetKind: aws.object_store
      targetType: aws.s3

  fieldValues:
    - field: interface.bucketClass
      targetKind: aws.object_store
      targetType: aws.s3
      values:
        - standard
        - archive

  schemas:
    - id: aws-s3-interface
      appliesToKind: aws.object_store
      appliesToInterfaceType: aws.s3
      description: Validates the aws.s3 interface shape.
      schema:
        $schema: https://json-schema.org/draft/2020-12/schema
        type: object
        required:
          - kind
          - interface
        properties:
          kind:
            const: aws.object_store
          interface:
            type: object
            required:
              - type
            properties:
              type:
                const: aws.s3
              bucketClass:
                enum:
                  - standard
                  - archive
            additionalProperties: true
```

---

# 7. Extension Vocabulary Rules

Extensions MAY define:

- New Condition kinds
- New interface types
- New Condition fields
- New interface fields
- New field values
- JSON Schema validation schemas

Extensions MUST NOT:

- Redefine core fields
- Change core field meaning
- Define additional top-level profile fields
- Define additional fields within `metadata` or `workload`
- Encode secrets or concrete target-environment values

Core-reserved top-level fields:

- `apiVersion`
- `kind`
- `metadata`
- `workload`
- `extensions`
- `conditions`

Core-reserved Condition fields:

- `name`
- `optional`
- `kind`
- `interface`

Core-reserved interface field:

- `type`

## 7.1 Vocabulary Definition Scope

Vocabulary definition scope is determined by:

- Vocabulary category
- Field path, for field values
- Target Condition `kind`, when applicable
- Target `interface.type`, when applicable

A profile is invalid if the resolved extension set contains more than one definition for the same vocabulary item in the same scope.

Two definitions conflict even when they are identical.

An extension that uses vocabulary defined by another extension MUST declare a dependency on that extension. It MUST NOT redefine that vocabulary in the same definition scope.

An extension MAY define vocabulary with overlapping purpose when it uses a distinct name or definition scope.

## 7.2 Namespacing

Extension-defined vocabulary SHOULD use namespaced identifiers when the vocabulary is vendor-specific, platform-specific, experimental, or likely to conflict.

Namespacing applies to:

- `kind` values
- `interface.type` values
- Condition field names
- Interface field names
- Field values that are not intended to be shared vocabulary

Examples:

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

Unqualified vocabulary MAY be used when it resolves to exactly one definition in the resolved extension set.

---

# 8. Validation

Validation occurs in this order:

1. Core structural validation
2. Extension declaration resolution
3. Extension dependency resolution
4. Vocabulary definition and conflict validation
5. Extension JSON Schema validation

## 8.1 Structural Validation

A profile is structurally invalid if:

- `apiVersion` is missing or is not `runtimeconditions.io/v1alpha1`
- `kind` is missing or is not `RuntimeConditionsProfile`
- `metadata` is missing or is not an object
- `metadata.name` is missing or is not a non-empty string
- `metadata.labels` is present and is not a string-to-string mapping
- `workload` is missing or is not an object
- `workload.uri` is missing or is not a non-empty string
- `workload.version` is present and is not a non-empty string
- `extensions` is missing or is not an array
- `conditions` is missing or is not an array
- Any required field is present with `null`
- Any mapping contains duplicate keys

A Condition is structurally invalid if:

- It is not an object
- `kind` is missing or is not a non-empty string
- `interface` is missing or is not an object
- `interface.type` is missing or is not a non-empty string
- `name` is present and is not a non-empty string
- `optional` is present and is not a boolean

Condition names, when present, MUST be unique within the profile.

## 8.2 Vocabulary Validation

A profile is invalid if:

- A Condition `kind` is not defined by exactly one resolved extension
- An `interface.type` is not defined by exactly one resolved extension for the declared `kind`
- An extension-defined Condition field is not defined by exactly one resolved extension for its scope
- An extension-defined interface field is not defined by exactly one resolved extension for its scope
- An extension-defined field value is not defined by exactly one resolved extension for its scope
- Any resolved extension dependency is missing
- Resolved extensions contain a dependency cycle
- Resolved extensions contain a vocabulary definition conflict
- Any applicable extension JSON Schema validation fails

## 8.3 Validity Levels

Validators SHOULD distinguish:

| Level | Description |
| ----- | ----------- |
| **Structural validity** | The document satisfies the core shape and type rules |
| **Extension-resolved validity** | All extensions resolve and all vocabulary has exactly one definition in scope |
| **Semantic validity** | The profile satisfies all applicable extension JSON Schema validations |

A core-only profile with an empty `conditions` array can be structurally valid.

A profile with non-empty `conditions` is not extension-resolved valid unless every Condition kind, interface type, extension-defined field, and extension-defined value resolves to exactly one definition.

## 8.4 Validator Diagnostics

Validation errors SHOULD identify:

- Error category
- Location within the profile
- Offending field or value
- Relevant extension, when applicable
- Expected valid type or vocabulary, when practical

Suggested error categories:

- `structural`
- `unknown-extension`
- `extension-dependency`
- `unresolved-vocabulary`
- `vocabulary-conflict`
- `semantic`

---

# 9. Conformance

## 9.1 Profile Conformance

A conforming profile MUST:

- Satisfy core structural requirements
- Declare all extensions required to interpret its vocabulary
- Avoid unresolved vocabulary
- Avoid vocabulary definition conflicts
- Satisfy all semantic validation schemas for resolved extensions
- Avoid secret values and concrete target-environment values

## 9.2 Extension Conformance

A conforming extension MUST:

- Use a valid extension identifier
- Provide a valid extension definition artifact
- Identify all vocabulary it defines
- Declare exact-version dependencies on vocabulary defined by other extensions
- Avoid redefining vocabulary defined by another resolved extension
- Respect core field placement and reserved-name rules
- Express semantic validation with JSON Schema 2020-12 schemas

## 9.3 Generator Conformance

A conforming generator MUST emit structurally valid profiles.

A conforming generator MUST NOT emit secret values or concrete target-environment values.

A conforming generator SHOULD fail before emitting a profile with unresolved vocabulary, missing extensions, dependency errors, or known vocabulary conflicts.

## 9.4 Validator Conformance

A conforming validator MUST implement the validation layers in Section 8.

A conforming validator MUST reject structurally invalid profiles, unresolved extensions, dependency cycles, vocabulary conflicts, unresolved vocabulary, and JSON Schema validation failures.

## 9.5 Adapter Conformance

An Adapter consumes a Runtime Conditions Profile for a target platform, catalog, policy system, or deployment workflow.

A conforming Adapter MUST locate declared extension definitions and transitive extension dependencies before interpreting extension-defined vocabulary.

A conforming Adapter MUST NOT treat a structurally valid but extension-unresolved profile as semantically valid.

A conforming Adapter MUST preserve the distinction between profile requirements and target-environment fulfillment choices.

A conforming Adapter MAY map valid Conditions to platform resources, services, policies, configuration inputs, or deployment workflow steps. Such mappings are outside the profile data model.

---

# 10. Examples

## 10.1 Core-Only Structural Profile

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: structural-profile
  labels:
    owner.example.com/team: platform

workload:
  uri: https://github.com/example-org/example-service
  version: v1.2.3

extensions: []

conditions: []
```

## 10.2 Unresolved Conditions

This profile is structurally valid but not extension-resolved valid.

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: unresolved-profile

workload:
  uri: https://github.com/example-org/example-service

extensions: []

conditions:
  - name: primary-db
    kind: relational-database
    interface:
      type: connection
```

## 10.3 Extension-Backed Profile

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: checkout-service
  labels:
    owner.example.com/team: payments
    lifecycle.example.com/stage: production

workload:
  uri: https://github.com/example-org/checkout-service
  version: v1.2.3

extensions:
  - https://runtimeconditions.io/extensions/common-integrations:v1alpha1

conditions:
  - name: primary-db
    kind: datastore
    interface:
      type: relational
      engine: postgres

  - name: payments-api
    kind: api
    interface:
      type: http
      operations:
        - method: POST
          path: /charge
```

## 10.4 Extension-Backed Profile With Configuration

```yaml
apiVersion: runtimeconditions.io/v1alpha1
kind: RuntimeConditionsProfile

metadata:
  name: checkout-service
  labels:
    owner.example.com/team: payments

workload:
  uri: https://github.com/example-org/checkout-service
  version: v1.2.3

extensions:
  - https://runtimeconditions.io/extensions/common-integrations:v1alpha1
  - https://runtimeconditions.io/extensions/env-configuration:v1alpha1

conditions:
  - name: primary-db
    kind: datastore
    interface:
      type: relational
      engine: postgres
    configuration:
      env:
        - property: hostname
          name: POSTGRES_HOST
        - property: port
          name: POSTGRES_PORT
        - property: database
          name: POSTGRES_DATABASE
        - property: username
          name: POSTGRES_USERNAME
        - property: password
          name: POSTGRES_PASSWORD
          sensitive: true

  - name: request-cache
    kind: cache
    interface:
      type: key_value
      engine: redis
    configuration:
      alternatives:
        - env:
            - property: url
              name: REDIS_URL
        - env:
            - property: hostname
              name: REDIS_HOST
            - property: port
              name: REDIS_PORT

  - name: payments-api
    kind: api
    interface:
      type: http
      operations:
        - method: POST
          path: /charge
    configuration:
      env:
        - property: baseUrl
          name: PAYMENTS_API_URL
        - property: token
          name: PAYMENTS_API_TOKEN
          sensitive: true
```
