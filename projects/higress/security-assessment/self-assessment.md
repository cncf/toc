# Higress Security Self-Assessment

## Metadata

| | |
| --- | --- |
| Assessment stage | Complete project self-assessment, approved in [higress-group/higress#4177](https://github.com/higress-group/higress/pull/4177); updated 2026-08-06 |
| Software | <https://github.com/higress-group/higress> |
| Review state | Prepared from public project evidence; no independent security review or audit |
| Evidence branches | [`higress-group/community@main`](https://github.com/higress-group/community/tree/main) for this assessment and project-wide policies; [`higress-group/higress@main`](https://github.com/higress-group/higress/tree/main) for software evidence |
| Assessment outline | [CNCF TAG Security Self-Assessment](https://tag-security.cncf.io/community/assessments/guide/self-assessment/) |
| Intended TOC snapshot | `projects/higress/security-assessment/self-assessment.md` |
| Security provider | No. Higress provides security features, but its primary function is API gateway traffic management. |
| Languages | Go, C++, Rust, AssemblyScript, shell, and Helm/YAML |
| SBOM | Not generated for release artifacts. Go module files, Cargo lock data, and container build files provide dependency inputs. |

### Security Links

| Document | Location |
| --- | --- |
| Vulnerability reporting and response | [`SECURITY.md`](https://github.com/higress-group/higress/blob/main/SECURITY.md) |
| Architecture | [`docs/architecture.md`](https://github.com/higress-group/higress/blob/main/docs/architecture.md) |
| Helm defaults | [`helm/core/values.yaml`](https://github.com/higress-group/higress/blob/main/helm/core/values.yaml) |
| OpenSSF Best Practices | [Passing](https://www.bestpractices.dev/projects/12667), verified 2026-08-04 |


## Overview

Higress translates declarative ingress, Gateway API, service discovery, and
plugin configuration into xDS consumed by an Envoy-based gateway. It accepts
untrusted downstream traffic, selects upstream services, applies traffic and
security policy, and proxies requests and responses.

### Actors and Actions

- **Cluster administrator:** installs/upgrades Higress, grants RBAC, configures
  exposure, certificates, registries, and security contexts.
- **Gateway operator/platform engineer:** creates routes, services, policies,
  plugins, credentials, and observability configuration.
- **Application owner:** requests routes/policy and operates upstream services.
- **End user/client:** sends potentially hostile network requests.
- **Plugin author/provider:** supplies code that executes in the gateway's Wasm
  sandbox or native filter boundary.
- **External provider/registry:** supplies service discovery data, plugins,
  identity metadata, certificates, or AI/model APIs when configured.
- **Project maintainer/release manager:** reviews changes, responds to reports,
  and publishes releases.

### Background

Higress runs a gateway that accepts requests from clients and forwards them to
configured backend services or AI model providers. Operators define routes,
access policies, certificates, services, and plugins through Kubernetes or
supported service registries. The Higress controller reads those settings and
sends the resulting gateway configuration to Envoy, which applies the required
authentication, authorization, traffic changes, and routing. Because Higress
sits between external clients and backend systems, it processes untrusted
requests and may handle sensitive headers, request bodies, AI prompts,
responses, and TLS credentials. Installed plugins and external integrations
extend this request path and therefore become part of its security boundary.

### Architecture and Data Flow

The security-relevant flow is:

1. An authenticated Kubernetes user or automation writes Ingress, Gateway API,
   Higress/Istio, Secret, and policy resources to the Kubernetes API.
2. Higress controller and discovery components watch authorized resources,
   optional service registries, and certificate sources, then translate the
   desired state into xDS and SDS configuration.
3. Gateway pilot agents and Envoy processes receive configuration and secret
   material. Existing gateways may retain their last accepted configuration if
   the control plane becomes unavailable.
4. Untrusted clients connect to the gateway. Envoy and configured filters parse,
   authenticate, authorize, transform, and route requests to upstream services
   or explicitly configured AI/model providers.
5. The gateway emits access logs, metrics, and traces to destinations selected
   by the operator. Depending on configuration, those signals may include
   sensitive request metadata.

### Data Assets and Trust Boundaries

| Asset or boundary | Security relevance |
| --- | --- |
| Kubernetes configuration and status | Unauthorized mutation can redirect traffic, weaken policy, load code, or disrupt availability. |
| TLS keys, certificates, tokens, and upstream credentials | Disclosure or replacement can enable impersonation, traffic decryption, or unauthorized upstream access. |
| Request/response bodies, AI prompts, and model outputs | May contain application secrets, personal data, proprietary data, or attacker-controlled content. |
| xDS/SDS control-plane boundary | Integrity and availability determine what code, routes, clusters, policy, and secrets the data plane uses. |
| Gateway process boundary | Native Go/C++ filters execute inside the proxy trust boundary; a defect can affect the full gateway process. |
| Wasm plugin boundary | Wasm isolates a plugin from the native gateway process to the extent provided by the runtime, but an installed plugin can still observe or modify request and response data exposed through the filter API. Higress `matchRules` scope intended plugin execution; they are not a capability boundary against a malicious plugin. |
| External integration boundary | Registries, identity providers, certificate issuers, Redis, observability systems, OCI registries, and model providers have separate operators and trust. |
| Source and release boundary | Compromise of source control, CI identities, dependencies, images, or plugin artifacts can affect every adopter. |

### Goals

- Preserve the integrity and availability of configuration delivery between
  control and data planes, reject invalid xDS updates, and make transport trust
  assumptions explicit for operators.
- Terminate and originate TLS as configured and distribute private key material
  through Kubernetes Secrets and SDS.
- Isolate Wasm plugin execution from the gateway process to the extent provided
  by the Envoy/Wasm runtime.
- Enforce configured routing, authentication, authorization, traffic, and data
  policies consistently.
- Contact only the external registries, identity systems, observability
  backends, AI providers, and upstreams explicitly configured by the operator.
- Provide a private vulnerability reporting and coordinated disclosure path.

### Non-goals

Higress does not secure a compromised Kubernetes control plane, node, cluster
administrator, upstream service, identity provider, model provider, registry,
or native plugin. It does not guarantee that user-authored policy is correct,
provide regulatory certification, or replace network segmentation, secrets
management, PKI governance, application security, or incident response.

## Self-Assessment Use

This document is an internal analysis by the Higress project. It is not an
independent audit, certification, or attestation. It gives adopters and CNCF
reviewers an initial view of security boundaries, practices, and known gaps.

## Security Functions and Features

### Critical

- xDS configuration generation, transport, validation, and last-known-good
  behavior between controller/discovery and gateways.
- TLS termination/origination, SDS secret delivery, certificate issuance and
  rotation paths, and private-key access.
- Kubernetes RBAC, ServiceAccounts, token reviews, subject-access reviews, and
  admission/configuration validation.
- HTTP/TCP parsing, routing, upstream selection, and request/response mutation.
- Plugin loading and Wasm sandbox boundary; native Go/C++ filters share the
  gateway process trust boundary.
- Release workflows, container/plugin registries, dependency inputs, and
  published artifacts.

### Security Relevant

- Authentication and authorization plugins (JWT, OIDC, key, HMAC, basic auth),
  WAF, rate limiting, request blocking, and data masking.
- Pod/container security contexts, host networking, privileged mode, RBAC
  toggles, network exposure, and admin/debug endpoints.
- Access, audit-style, metrics, and trace output, which may contain sensitive
  request metadata depending on operator configuration.
- External service registries, Redis, certificate issuers, identity providers,
  OCI registries, and AI/model providers.

### Who configures and enforces security controls

The [operator guide](https://github.com/higress-group/community/blob/db86b75/docs/security-operator-responsibilities.md)
lists the actions required from the people running Higress. This guide is
proposed in the community repository and is awaiting merge.

Kubernetes enforces the RBAC permissions and container resource limits chosen
by operators. Higress and Envoy apply the routes and request policies they
configure. Operators must review custom resources, plugin settings, and image
sources before deployment. Configuration validation does not establish that a
policy is safe or that an image is trustworthy. For example, a valid route can
still point to the wrong backend, and a loaded plugin can still be malicious.
Operators must select trusted images and use their deployment tooling to check
signatures or attestations where available; they must not assume Higress checks
these automatically.

### Threat Model

This project-authored model uses the actors, assets, flows, and boundaries above.
Priority reflects generic impact for a shared production gateway; adopters must
adjust it for their tenancy model and data classification.

| ID | Priority | Threat scenario | Current mitigation and residual risk |
| --- | --- | --- | --- |
| HG-TM-01 | High | A principal with excessive Kubernetes permissions creates or changes routes, policies, Secrets, or plugins to hijack traffic or bypass controls. | The Kubernetes API server authenticates principals and applies RBAC and admission policy before Higress sees the resources; Higress then validates and translates admitted resources into data-plane configuration. The boundary depends on operator-managed Kubernetes policy. Cluster administrators remain trusted, and the controller currently has broad cluster-wide permissions. |
| HG-TM-02 | High | A compromised controller, gateway ServiceAccount, or ClusterRoleBinding exposes TLS keys or upstream credentials. | Kubernetes RBAC controls API access to Secrets; Higress/Istio SDS delivers secret material to Envoy rather than embedding it in images. Gateway and controller roles are separate, but secret access and several controller rules remain broad and require minimization. |
| HG-TM-03 | High | A malicious or compromised plugin/image executes code, exfiltrates traffic, or alters policy. | Explicit installation is the control point. Higress `defaultConfig` applies globally to the selected gateway workload, while `matchRules` scope intended execution by ingress, domain, or service. These rules are not a security sandbox against a malicious installed plugin. Wasm isolates the plugin from the native gateway process to the extent provided by the runtime, while native filters share that process; release/plugin artifacts lack project-wide signature, SBOM, and provenance guarantees. |
| HG-TM-04 | High | Control-plane or xDS/SDS channel compromise injects malicious configuration or secret material, or prevents updates. | Envoy validates delivered resources and can retain the last accepted configuration. During a control-plane partition, policy removal or tightening cannot propagate, so stale permissive configuration may remain active until connectivity is restored. Operators must protect control-plane identities, endpoints, network paths, and Kubernetes access; transport and deployment assumptions require a dedicated hardening guide. |
| HG-TM-05 | High | Hostile downstream traffic exploits an Envoy/filter parser defect or exhausts connections, memory, CPU, sockets, or upstream capacity. | Envoy validation, resource limits, timeouts, rate-limit/circuit-breaker features, probes, and multiple replicas can limit impact. Safe values are deployment-specific and timely Envoy/Higress patching remains essential. |
| HG-TM-06 | High | A tenant accidentally or deliberately attaches a route or policy to another tenant's gateway and exposes or disrupts traffic. | Kubernetes RBAC and admission policy control who may create or change resources. Higress's Gateway API translation enforces a listener's `allowedRoutes.namespaces` for Route-to-Gateway attachment and separately checks `ReferenceGrant` for cross-namespace backend and Secret references. A Route from a namespace not selected by `allowedRoutes` does not attach, and a cross-namespace reference without a matching `ReferenceGrant` is rejected. Higress does not make a hostile multi-tenant cluster safe when administrators grant overlapping write privileges. |
| HG-TM-07 | High | Request bodies, credentials, or responses are disclosed to a configured plugin, observability backend, identity service, or external provider. When Higress is used as an AI gateway, request bodies can include AI prompts and responses can contain model output. | Integrations require operator configuration and can be restricted through credentials and network policy. Operators remain responsible for egress allowlists, provider contracts, log redaction, residency, and data-retention controls. |
| HG-TM-08 | Medium | Logs, metrics, traces, admin/debug endpoints, or configuration dumps expose credentials or sensitive request metadata. | Admin interfaces are not intended as public endpoints and telemetry is configurable. Access control, redaction, retention, and exposure are deployment responsibilities; unsafe logging or endpoint exposure remains possible. |
| HG-TM-09 | High | Certificate expiration, issuer compromise, or unauthorized Secret replacement causes outage or impersonation. | SDS and automatic HTTPS support dynamic certificate updates and renewal. Operators must secure issuers, monitor expiry/renewal, test emergency rotation, and control Secret writers. |
| HG-TM-10 | High | A source-control, dependency, GitHub Actions, registry, or maintainer-account compromise produces malicious release artifacts. | Public review, CI tests, license checks, CodeQL, pinned dependency/submodule inputs, and GitHub release workflows provide layers. Missing universal immutable action pinning, project-wide release SBOM, artifact-signature, and machine-readable provenance-attestation guarantees, and documented repository access controls leave material residual risk. |

The model should be reviewed after material architecture, privilege, plugin,
release-pipeline, or trust-boundary changes. The project does not yet enforce a
review cadence or require a named security reviewer, which is itself a process
gap.

## Project Compliance

The open-source project does not claim PCI-DSS, SOC 2, ISO 27001, GDPR, or
other regulatory certification. Deployers are responsible for assessing their
configuration and operational environment. Source is Apache-2.0 and pull
requests run license header and dependency-license checks.

## Secure Development Practices

Pull requests are publicly reviewed and run build/unit tests with Go race
detection, Gateway API and Higress conformance tests, plugin tests, and license
checks. CodeQL runs on pull requests targeting `main`, pushes to `main`, and a
weekly schedule; the project also enforces a scoped Go vet warning gate. A
`golangci-lint` target remains in the repository, but current CI does not invoke
it, so it provides no active CI control.
Release tags trigger image and CLI/CRD artifact builds. Dependency inputs are
versioned, but the project does not yet provide SBOMs, artifact signatures, or
signed provenance attestations consistently across all release artifacts. Not
all workflow actions are pinned to immutable commits. The public repository
does not prove a required reviewer count, signed-commit requirement,
organization-wide 2FA, or branch-protection configuration; those controls
require separate repository-settings evidence.

Ordinary project-team communication uses GitHub issues, pull requests,
discussions, localized community channels, and Discord. Inbound users use the
same public channels. GitHub Private Security Advisories are the required,
vendor-neutral, and authoritative private vulnerability intake documented in
`SECURITY.md`. Reporters may also use the Alibaba Security Response Center as
an optional additional channel. An SRT member with ASRC access may correlate an
ASRC case with the project record, but ASRC access is not required for SRT
membership or project handling.
Releases, the project website, and the WeChat Official Account are outbound
channels. [`COMMUNITY.md`](https://github.com/higress-group/community/blob/main/COMMUNITY.md)
is the authoritative inventory of public, subproject, and narrowly scoped
non-public channels.

Higress operates in the Kubernetes networking and cloud-native gateway
ecosystem. It implements Ingress and Gateway API and builds on Envoy, Istio,
OCI, Prometheus/OpenTelemetry conventions, and optional service registries.

## Security Issue Resolution

[`SECURITY.md`](https://github.com/higress-group/higress/blob/main/SECURITY.md)
prohibits public vulnerability reports and
requires reporters to use GitHub Private Security Advisories as the
authoritative project intake. Alibaba Security Response Center submission is
optional. The named Security Response Team is the current maintainer list.
For each case it assigns a triage coordinator, fix lead, independent reviewer
and release lead, and disclosure lead, with at least two unconflicted members.
The policy documents conflicts and escalation. Its targets are acknowledgement
within three business days and triage within 14 days, followed by private fix
development, coordinated disclosure (typically within 90 days), a GitHub
Security Advisory, and a CVE request where appropriate.

An operational security incident in an adopter environment remains the
adopter's responsibility. The project handles defects in project code and
artifacts. For confirmed project vulnerabilities, maintainers triage impact,
develop and release a fix, notify affected users through a GitHub Security
Advisory and release information, and coordinate timing with the reporter.

## Appendix

### Known Gaps

- Project-wide release SBOM, artifact-signature, and machine-readable
  provenance-attestation coverage is incomplete.
- The controller's ClusterRole is broad and its default container security
  context is empty. Gateway non-root defaults depend on Kubernetes/platform
  capability; legacy fallback adds `NET_BIND_SERVICE` and allows escalation.
- No dedicated fuzzing/DAST program or automated upgrade/downgrade matrix is
  documented.
- This threat model is project-authored, has not been independently validated,
  and is not yet backed by a published data-flow diagram with explicit trust
  boundaries or an independent security audit.

### Known Issues Over Time

Published project advisories are available from the repository's
[Security Advisories](https://github.com/higress-group/higress/security/advisories)
page. This assessment does not claim that the absence of a public advisory for
a period means no vulnerability existed. The project has not published an
aggregate vulnerability history or mean-time-to-remediation report.

### OpenSSF Best Practices

The [Higress OpenSSF entry](https://www.bestpractices.dev/projects/12667)
reached **Passing** on 2026-08-04. Supporting evidence includes CodeQL coverage
for pull requests, pushes to `main`, and weekly scans
([#4186](https://github.com/higress-group/higress/pull/4186)); an enforced Go
vet warning gate
([#4193](https://github.com/higress-group/higress/pull/4193)); and the completed
Critical/High CodeQL remediation tracker
([#4207](https://github.com/higress-group/higress/issues/4207)), which records
zero open Critical and High findings and the disposition of every finding in
scope. Dedicated project fuzzing/DAST remains a defense-in-depth improvement,
but it is not an unmet requirement for the achieved Passing badge.

### Example Use Cases

1. A platform team exposes Kubernetes services through Gateway API with TLS,
   JWT authentication, rate limiting, and Prometheus metrics.
2. An AI platform routes requests across model providers while applying token
   quotas, content policy, and request/response observability.
3. A microservice platform discovers Nacos/Consul services and exposes them
   through stable API routes without reloading the data plane.

### Related Projects and Vendors

Envoy supplies the proxy foundation; Istio supplies xDS/control-plane building
blocks; Kubernetes Gateway API and Ingress supply standard configuration APIs.
Kong, Apache APISIX, Envoy Gateway, Traefik, and ingress controllers address
overlapping gateway use cases with different APIs, extension models, and
operational tradeoffs. Commercial products may distribute or manage Higress,
but they are outside this open-source security assessment.
