# Runtime Conditions Profile draft and support artifacts

This directory contains the draft Runtime Conditions Profile Specification, and the minimum supporting artifacts needed to review it as a cohesive initiative.

## Contents

- `runtime-conditions-profile-specification.md` - core draft specification.
- `extensions/` - machine-readable extension definition files that define Condition vocabulary, allowed values, and validation schemas used by the examples.
- `examples/` - complete profile documents that show the core shape and extension usage.
- `review-notes/` - short non-normative notes for common review questions.

## Review scope

The core profile specification and extension definition shape are the main items for review. The first-party extension YAML files show how concrete vocabulary can be defined outside the core. The AWS object-store extension is intentionally illustrative: it demonstrates third-party extension mechanics, while clearly falling short of a full third-party implementation.

Generators, platform adapters, and a working demo are intentionally not included here. They can be found at https://github.com/colinjlacy/runtime-conditions-profiles.

For a guided walkthrough explaining the core concepts, see: https://colinjlacy.github.io/runtime-conditions-profiles/
