# Interoperability Profile — OCI Annotation Conventions (Draft v0.1)

Companion to the [initiative README](./README.md). This lays out the OCI annotation keys for
the v1 Interoperability Profile (the "Manifest Contract" mentioned in the README's
[Scope Overview](./README.md#scope-overview)). Still a draft — open items from review are
listed at the bottom under Open Questions.

## Unit of conformance

A manifest under this profile describes one self-contained model artifact: either a base model
or a single fine-tuned model/adapter packaged as its own OCI artifact. Composite bundles (a base
model plus one or more adapters referenced together) are out of scope for v1 — each artifact in
the bundle carries its own manifest, and how they relate to each other is handled separately
under the README's "Relationship Metadata" work. So `mof.class` and the other fields below
always describe the single artifact the manifest is attached to, never a bundle.

## Requirement levels and evidence

- **MUST** — required for v1 conformance.
- **SHOULD** — recommended; skipping it doesn't break conformance but weakens the trust/portability signal.
- **MAY** — optional/informational.

One thing worth calling out: annotations that assert something about security or openness
(signing, SBOM, provenance, MOF class) are pointers, not proof. Just having the key/value pair
set doesn't mean the underlying evidence exists. Conformance means the referenced artifact is
actually there, tied to the correct digest, and resolvable — e.g. via an OCI referrer/subject
relationship or an attached attestation. That expectation applies to every row in the "Evidence
requirement" column below, not just the ones that spell it out.

## Annotation table

| Key | Requirement | Values | Description | Evidence requirement |
|---|---|---|---|---|
| `org.cncf.ai.interop.profile.version` | MUST | Semantic version (e.g. `1.0.0`) | Which version of this profile the manifest conforms to. Not the same as the artifact's own version — see `org.opencontainers.image.version` below. | Informational; tells a validator which spec version to check against. |
| `org.cncf.ai.artifact.type` | MUST | `model` (only valid value in v1); future versions may add `skill`, `rag-context`, `workflow` | Type of AI artifact the manifest represents. | Informational. |
| `org.cncf.ai.lifecycle.status` | MUST | `experimental`, `validated`, `deprecated`, `product-ready` | Maturity/promotion status, separate from structural conformance — an artifact can be fully compliant and still `experimental`. | Informational; a GitOps policy could gate on this. |
| `org.cncf.ai.model.mof.class` | MUST | `I`, `II`, `III` | LF AI & Data Model Openness Framework class claimed for this artifact (see Unit of conformance above for what "this artifact" means). | Needs `mof.components` present and consistent with the claimed class; may also need a resolvable MOF-generated model/data card. |
| `org.cncf.ai.model.mof.version` | MUST when `mof.class` is set | MOF spec version, e.g. `1.0` | Which MOF spec version the class/components were derived from. | Informational. |
| `org.cncf.ai.model.mof.components` | MUST when `mof.class` is set | Comma-separated list, values drawn from the MOF spec's own component vocabulary — no free-text synonyms | Which MOF components are present. | Values need to be checked against the fixed MOF vocabulary, not accepted as arbitrary strings. |
| `org.cncf.ai.security.signing.framework` | MUST | `sigstore`, `notation` (Notary v2/TUF-based) | Framework used to sign the artifact. | The signature itself must be resolvable and verify against the artifact digest — the annotation by itself isn't enough. |
| `org.cncf.ai.security.sbom.format` | MUST | `spdx`, `cyclonedx` | Format of the attached SBOM. Bumped from SHOULD to MUST to line up with the README calling this "Required Trust Metadata." | SBOM must be resolvable (e.g. via OCI referrers) and tied to this artifact's digest. |
| `org.cncf.ai.security.provenance.type` | MUST | `slsa-v1.0`, `in-toto` | Type of provenance attestation. Bumped to MUST for the same reason as SBOM format. | Attestation must be resolvable with a subject digest matching this artifact. |
| `org.cncf.ai.packaging.format` | SHOULD | `modelpack` | Packaging format of the assembled content. If it doesn't apply, leave the key out rather than setting it to `none`. | Informational. |

## Reused OCI annotations (not redefined)

Identity fields that the base OCI image spec already covers are reused here instead of adding
new `org.cncf.ai.*` keys for the same thing:

| Key | Purpose |
|---|---|
| `org.opencontainers.image.version` | Version of the model artifact itself (not the profile version above). |
| `org.opencontainers.image.title` | Model artifact name. |
| `org.opencontainers.image.authors` | Authorship. |
| `org.opencontainers.image.source` | Source repository/location. |
| `org.opencontainers.image.created` | Creation timestamp. |
| `org.opencontainers.image.licenses` | License identifier(s). |

## Open questions

1. **MOF component vocabulary** — needs to be pinned to a specific published MOF version
   instead of paraphrased here; someone should confirm the exact list against the LF AI & Data
   source document.
2. **Signing framework list** — the README's supply chain security section also mentions
   OpenPubkey and other emerging zero-trust identity protocols. Do those get added as enum
   values, or are they explicitly out of scope for v1?
3. **Composite bundles** — how `mof.class` and trust metadata roll up when multiple compliant
   artifacts ship together isn't solved here; it's deferred to the README's relationship
   metadata work.
4. **Worked example** — this table should get run through a full example manifest end to end
   (build, sign, push, GitOps pull, KServe deploy) before it's locked in.
