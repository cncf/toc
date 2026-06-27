# Extension definitions

Runtime Conditions Profiles have a small, extensible core. Concrete integration vocabulary is defined by extensions.

This directory includes:

- `first-party/common-integrations-v1alpha1.yaml` - vocabulary for common integration categories such as APIs, datastores, and caches. Intended to be supported upon release.
- `first-party/env-configuration-v1alpha1.yaml` - vocabulary for declaring workload-facing environment variable inputs. Intended to be supported upon release.
- `third-party-example/aws-object-store-v1alpha1.yaml` - illustrative third-party extension for an AWS S3-style object store condition. Intended as an example for third-party authorship.

The first two files are the extension artifacts referenced by the complete examples. They are included so reviewers can see that condition kinds, interface types, fields, field values, and JSON Schema validation can be declared outside the core specification.

The AWS file is intentionally provider-specific. It exists to show the expected third-party extension path: an extension can declare its own URI, depend on first-party extensions, define provider-specific vocabulary, and constrain that vocabulary using JSON Schema.
