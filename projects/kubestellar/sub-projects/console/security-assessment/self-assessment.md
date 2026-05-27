# Security Self-Assessment: KubeStellar Console

This document follows the [CNCF TAG-Security self-assessment template](https://github.com/cncf/tag-security/blob/main/assessments/guide/self-assessment.md).

> **Scope:** This assessment covers **[kubestellar/console](https://github.com/kubestellar/console)** only — the multi-cluster Kubernetes dashboard and its local agent (`kc-agent`). It does **not** cover [KubeStellar Core](https://github.com/kubestellar/kubestellar), the multi-cluster orchestration control plane. Console is a standalone project that connects to any Kubernetes cluster via standard kubeconfig; it does not depend on KubeStellar Core to function.

## Table of Contents

- [Metadata](#metadata)
- [Overview](#overview)
- [Self-Assessment Use](#self-assessment-use)
- [Security Functions and Features](#security-functions-and-features)
- [Project Compliance](#project-compliance)
- [Secure Development Practices](#secure-development-practices)
- [Security Issue Resolution](#security-issue-resolution)
- [Appendix](#appendix)
  - [Deployment Architecture](#deployment-architecture)
  - [Data Storage Summary](#data-storage-summary)

## Metadata

| | |
|---|---|
| **Software** | [KubeStellar Console](https://github.com/kubestellar/console) |
| **Security Provider** | No — Console is a user-facing dashboard, not a security tool |
| **Languages** | Go (backend), TypeScript/React (frontend) |
| **SBOM** | Not currently generated; planned via GoReleaser in future releases |

## Overview

### Background

KubeStellar Console is an AI-powered multi-cluster Kubernetes dashboard that provides:

- **Cluster management** — Visual overview of resources across multiple Kubernetes clusters
- **Guided install missions** — Step-by-step deployment of CNCF projects with pre-flight checks, validation, and rollback
- **AI-assisted operations** — Natural language queries about cluster state via MCP (Model Context Protocol) bridge
- **Real-time monitoring** — WebSocket-powered live event streaming from connected clusters

### Actors

| Actor | Description |
|-------|-------------|
| **Console User** | Authenticated (GitHub OAuth) or unauthenticated (demo mode) user accessing the web dashboard |
| **kc-agent** | Local agent running on user's machine that bridges browser to local kubeconfig and MCP servers |
| **AI Provider** | External service (Claude, OpenAI, or Gemini) that processes natural language queries via MCP |
| **Kubernetes API** | Target clusters accessed via kubeconfig contexts |
| **Netlify Functions** | Serverless backend for the hosted console.kubestellar.io deployment |

### Actions

| Action | Actor | Description |
|--------|-------|-------------|
| View cluster resources | Console User | Browse pods, deployments, services, events across connected clusters |
| Execute mission | Console User | Run guided install/uninstall missions on target clusters |
| AI chat query | Console User → AI Provider | Ask natural language questions about cluster state |
| Connect clusters | kc-agent | Discover and tunnel local kubeconfig contexts to hosted Console |
| Authenticate | Console User | Sign in via GitHub OAuth 2.0 |

### Goals

- Provide authenticated access to Kubernetes cluster information over HTTPS (TLS 1.2+ enforced by Netlify for the hosted deployment; user-configured TLS for self-hosted). All API endpoints except `/health` and demo-mode routes require a valid JWT session token.
- Never store or transmit Kubernetes credentials to the server. In the hosted deployment (`console.kubestellar.io`), `kc-agent` runs on the user's local machine, reads `~/.kube/config` locally, executes `kubectl` commands locally, and returns only JSON-serialized resource data (pods, deployments, events, etc.) to the Console backend via an authenticated WebSocket tunnel. The kubeconfig file contents — including certificates, tokens, and private keys — never leave the user's machine. See [Deployment Architecture](#deployment-architecture) for a diagram of both hosted and self-hosted data flows.
- Enforce principle of least privilege by inheriting the user's existing Kubernetes RBAC. Console does not create its own ServiceAccount, ClusterRole, or any Kubernetes RBAC objects. The `kc-agent` executes `kubectl` using whatever kubeconfig contexts the user has configured. If a user's kubeconfig grants only `get` access to pods in namespace `default`, Console can only display those pods — it cannot escalate privileges. The Console backend itself has no Kubernetes credentials; it is exactly as privileged as the user's kubeconfig, no more.
- Protect user sessions with JWT tokens validated on every API request
- Prevent cross-site scripting (XSS) and other web application attacks

### Non-Goals

- Console does **not** manage Kubernetes RBAC policies — it inherits existing permissions
- Console does **not** provide network security or service mesh capabilities
- Console does **not** persistently store raw Kubernetes resource data on the server — pods, deployments, and services are fetched on-demand via the Kubernetes API. See [Data Storage Summary](#data-storage-summary) for details on what is persisted.
- Console is **not** a secrets management tool

## Self-Assessment Use

This self-assessment is created by the KubeStellar Console maintainers to perform an internal analysis of the project's security. It is intended to assist TAG-Security in their joint assessment and to identify areas for improvement.

## Security Functions and Features

### Critical

| Component | Description |
|-----------|-------------|
| **GitHub OAuth 2.0** | Primary authentication mechanism; no passwords stored. OAuth scope is limited to `user:email` ([`auth.go:179`](https://github.com/kubestellar/console/blob/main/pkg/api/handlers/auth.go#L179)). Server-side state tokens with 10-minute expiry prevent CSRF during the OAuth redirect flow. |
| **JWT Session Tokens** | Session tokens signed with HMAC-SHA256 (HS256). Algorithm enforcement: the JWT parser accepts **only** HS256 via `jwt.WithValidMethods([]string{"HS256"})` ([`auth.go:36`](https://github.com/kubestellar/console/blob/main/pkg/api/middleware/auth.go#L36)), preventing algorithm confusion attacks. The `none` algorithm and all non-HMAC methods are rejected at parse time. Defense-in-depth: the keyfunc additionally verifies the token's signing method is `*jwt.SigningMethodHMAC` before returning the secret. **Secret management:** In production, the JWT secret must be provided via the `JWT_SECRET` environment variable — the server refuses to start without it ([`server.go:170`](https://github.com/kubestellar/console/blob/main/pkg/api/server.go#L170)). In development mode, a 32-byte (256-bit) cryptographically random secret is generated per server startup via `crypto/rand.Read()`. **Token lifetime:** 7 days (168 hours) with automatic silent refresh when 50% of the lifetime has elapsed. **Storage:** JWT is stored in an HttpOnly cookie (`kc_auth`) with `SameSite=Lax` for CSRF protection. The `Secure` flag is set when the deployment URL uses HTTPS (always true for the hosted deployment; conditional for local development on `http://localhost`). The JWT is not included in URL query parameters or redirect URLs. **Revocation:** Two-tier revocation system using JTI (JWT ID) claims — an in-memory cache for fast lookups backed by persistent SQLite storage that survives server restarts. Tokens are revoked on explicit logout and on refresh (old token revoked before new token issued). |
| **kubeconfig isolation** | Kubernetes credentials never leave the user's machine. See [Deployment Architecture](#deployment-architecture) for details. |
| **HTTPS transport** | All production traffic encrypted via TLS. Hosted deployment uses Netlify's automatic TLS (1.2+). Self-hosted deployments use user-configured TLS termination (ingress controller or reverse proxy). |
| **WebSocket authentication** | Hub WebSocket connections use first-message authentication: the client must send a `{"type":"auth","token":"..."}` JSON message within 5 seconds of connection. The token is validated via the same `ParseJWT` function used by HTTP routes (HS256-only, revocation-checked). Connections idle for more than 90 seconds are terminated. Server-side ping/pong keepalive runs every 30 seconds. |

### Security Relevant

| Component | Description |
|-----------|-------------|
| **Content Security Policy** | HTTP headers restrict script sources and frame embedding |
| **CORS configuration** | API restricts cross-origin requests to known Console origins |
| **Input sanitization** | User inputs (search queries, mission parameters) sanitized before rendering |
| **Dependency scanning** | Dependabot, npm audit, and CodeQL identify vulnerable dependencies |
| **Post-build safety checks** | Automated post-build verification scans production bundles for Vite define corruption, missing critical chunks, MSW leakage, and bundle size regressions |

## Project Compliance

KubeStellar Console does not currently hold formal compliance certifications. The project follows:

- **CNCF Code of Conduct** — Adopted and enforced
- **Apache License 2.0** — CNCF-compatible open-source license
- **DCO (Developer Certificate of Origin)** — Required for all contributions
- **OpenSSF Scorecard** — Automated weekly scoring

## Secure Development Practices

### Development Pipeline

| Practice | Implementation |
|----------|---------------|
| **Static Analysis (Go)** | CodeQL with extended security queries, gosec, nilaway |
| **Static Analysis (TypeScript)** | CodeQL with extended queries, ESLint, TypeScript strict mode null checks |
| **Dependency Management** | Dependabot (Go, npm, GitHub Actions), weekly automated updates |
| **Container Scanning** | Planned — container image vulnerability scanning not yet integrated into CI/CD |
| **Secret Detection** | GitHub repository-level secret scanning enabled; CI/CD integration planned |
| **Code Review** | All changes require PR review; direct commits to main prohibited |
| **Signed Commits** | DCO sign-off required on all commits |
| **Branch Protection** | Main branch protected; CI must pass before merge |
| **Post-Merge Verification** | Automated Playwright E2E tests run against the deployed production site after each merge, targeted by issue labels and changed file paths |

### Communication Channels

| Channel | Address |
|---------|---------|
| **Internal (maintainers)** | [kubestellar-dev-private@googlegroups.com](mailto:kubestellar-dev-private@googlegroups.com) |
| **Incoming (users)** | GitHub Issues, [#kubestellar-dev Slack](https://cloud-native.slack.com/channels/kubestellar-dev) |
| **Security Announcements** | [kubestellar-security-announce@googlegroups.com](mailto:kubestellar-security-announce@googlegroups.com) |

## Security Issue Resolution

### Responsible Disclosure Process

Vulnerability reports should be submitted via [GitHub Security Advisories](https://github.com/kubestellar/console/security/advisories) (preferred) or sent to [kubestellar-security-announce@googlegroups.com](mailto:kubestellar-security-announce@googlegroups.com). See [SECURITY.md](https://github.com/kubestellar/console/blob/main/SECURITY.md) for full details.

### Response Timeline

| Step | Timeline |
|------|----------|
| **Acknowledgment** | Within 3 working days |
| **Initial Assessment** | Within 5 working days |
| **Patch Development** | Critical: 72 hours; High: 7 days |
| **Public Disclosure** | After patch is available, coordinated with reporter |

### Incident Response

The security response team (listed in [SECURITY_CONTACTS](https://github.com/kubestellar/console/blob/main/SECURITY_CONTACTS)) triages all reports. Critical vulnerabilities trigger an immediate patch release; non-critical issues are addressed in the next scheduled release.

## Appendix

### Deployment Architecture

KubeStellar Console supports two deployment modes with different security boundaries:

```
Hosted Mode (console.kubestellar.io)
════════════════════════════════════

  User's Local Machine                          Netlify CDN
 ┌─────────────────────────────┐              ┌───────────────────────┐
 │  kc-agent (127.0.0.1:8585)  │   WebSocket  │  Console Frontend     │
 │  ┌─────────────────────────┐│   (wss://)   │  (React SPA)          │
 │  │ Reads ~/.kube/config    ││◄────────────►│                       │
 │  │ Executes kubectl        ││  JSON only:  │  Netlify Functions     │
 │  │ locally                 ││  pods, svcs, │  (serverless backend)  │
 │  └───────────┬─────────────┘│  events,     └───────────────────────┘
 │              │               │  deployments          ▲
 │              ▼               │                       │ HTTPS
 │  ┌─────────────────────────┐│                       │
 │  │ Kubernetes Clusters     ││              ┌────────┴──────────────┐
 │  │ (via kubeconfig RBAC)   ││              │  User's Browser        │
 │  └─────────────────────────┘│              │  (dashboard UI)       │
 └─────────────────────────────┘              └───────────────────────┘

 ▲ kubeconfig credentials NEVER                 Only JSON resource data
   cross this boundary                          crosses the network


Self-Hosted Mode (localhost:8080)
════════════════════════════════

 ┌───────────────────────────────────────────────────┐
 │  Single Machine / Pod                             │
 │  ┌──────────────────────┐  ┌────────────────────┐ │
 │  │ Go Backend (:8080)   │  │ kc-agent           │ │
 │  │ REST API + WebSocket │◄►│ kubectl + MCP      │ │
 │  │ OAuth + JWT Auth     │  │ reads kubeconfig   │ │
 │  └──────────────────────┘  └─────────┬──────────┘ │
 │                                      │              │
 │                            ┌─────────▼───────────┐  │
 │                            │ Kubernetes API      │  │
 │                            │ (user's RBAC)       │  │
 │                            └─────────────────────┘  │
 └───────────────────────────────────────────────────┘
   kubeconfig stays on this machine in both modes
```

**Key security properties:**

- **kc-agent binds to `127.0.0.1` only** ([`server.go:466`](https://github.com/kubestellar/console/blob/main/pkg/agent/server.go#L466)) — it is not accessible from the network.
- **In hosted mode**, the WebSocket tunnel between kc-agent and the Console backend carries only JSON-serialized Kubernetes resource data (pod specs, deployment status, event lists, etc.). Kubeconfig file contents — including certificates, bearer tokens, and private keys — are never transmitted.
- **In self-hosted mode**, the Go backend and kc-agent run on the same machine. The backend reads kubeconfig via the standard `client-go` library (`clientcmd.LoadFromFile`). No credentials leave the host.
- **kc-agent does expose localhost-only endpoints** (`/kubeconfig/import`, `/kubeconfig/preview`) that accept kubeconfig YAML for importing new cluster contexts. These are bound to `127.0.0.1:8585` and are not reachable from the hosted Console or any external network.

### Data Storage Summary

The following table details all data persisted by KubeStellar Console, both server-side and client-side:

| Storage | Type | What is Stored | Location | Retention |
|---------|------|----------------|----------|-----------|
| **Backend SQLite** | Relational DB | User accounts, dashboard layouts, card configurations, GPU utilization snapshots, revoked JWT tokens, user analytics events, feature requests | Server: `/data/console.db` | Indefinite |
| **Metrics History** | JSON file | Rolling cluster metrics (CPU/memory %, node counts, pod issues, GPU allocation per node) | kc-agent host: `~/.kc/metrics_history.json` | 7 days |
| **Frontend SQLite (OPFS)** | Browser DB | Cached Kubernetes resource data (pods, deployments, services) for performance | Browser OPFS: `/kc-cache.sqlite3` | Until browser cache cleared |
| **Frontend IndexedDB** | Browser DB | Same as OPFS (fallback when OPFS unavailable) | Browser IndexedDB: `kc_cache` | Until browser cache cleared |
| **localStorage** | Browser KV | User preferences, theme, AI settings, UI state | Browser localStorage | Until browser data cleared |
| **In-memory SSE cache** | Process memory | Short-lived Kubernetes API responses | Server memory | 15-second TTL; lost on restart |

**Key points:**

- Raw Kubernetes resource data (pods, deployments, services) is **not** stored server-side — it is fetched on-demand and cached only in browser storage and a 15-second in-memory server cache.
- Kubernetes **credentials** never leave the user's machine — the kc-agent proxies requests locally.
- The backend SQLite database stores user preferences and dashboard configuration, plus GPU utilization snapshots for historical trend analysis.
- The kc-agent's metrics history file (`~/.kc/metrics_history.json`) stores 7 days of aggregated cluster health metrics on the user's local machine for AI-assisted trend analysis.

### Known Issues and Areas for Improvement

1. **AI provider data flow** — Queries sent to external AI providers (Claude, OpenAI, Gemini) may include cluster resource names and metadata (namespace names, pod names, event messages). Users should be aware that AI chat queries are processed by third-party services subject to those providers' data handling policies.
2. **Demo mode** — Unauthenticated demo mode uses synthetic data only; no cluster access is possible without authentication.
3. **kc-agent trust model** — The kc-agent runs with the user's local permissions. A compromised kc-agent could access any cluster the user has kubeconfig access to. This is mitigated by the agent binding to `127.0.0.1` only (not accessible from the network) and running as a user-space process (not a privileged daemon).

### Related Security Documentation

- [SECURITY.md](https://github.com/kubestellar/console/blob/main/SECURITY.md) — Vulnerability reporting process
- [SECURITY_CONTACTS](https://github.com/kubestellar/console/blob/main/SECURITY_CONTACTS) — Security response team
- [HARDCODED_URLS.md](https://github.com/kubestellar/console/blob/main/docs/security/HARDCODED_URLS.md) — Explanation of intentional hardcoded URLs
