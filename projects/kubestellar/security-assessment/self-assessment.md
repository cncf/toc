# Security Self-Assessment: KubeStellar Console

This document follows the [CNCF TAG-Security self-assessment template](https://github.com/cncf/tag-security/blob/main/assessments/guide/self-assessment.md).

## Table of Contents

- [Metadata](#metadata)
- [Overview](#overview)
- [Self-Assessment Use](#self-assessment-use)
- [Security Functions and Features](#security-functions-and-features)
- [Project Compliance](#project-compliance)
- [Secure Development Practices](#secure-development-practices)
- [Security Issue Resolution](#security-issue-resolution)
- [Appendix](#appendix)
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

- Provide secure, authenticated access to Kubernetes cluster information
- Never store or transmit Kubernetes credentials beyond the user's local machine
- Enforce principle of least privilege — Console inherits the user's existing kubeconfig RBAC
- Protect user sessions with short-lived JWT tokens
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
| **GitHub OAuth 2.0** | Primary authentication mechanism; no passwords stored |
| **JWT Session Tokens** | Short-lived tokens for session management; validated on every API request |
| **kubeconfig isolation** | Kubernetes credentials never leave the user's machine; kc-agent proxies requests locally |
| **HTTPS transport** | All production traffic encrypted via TLS (Netlify/Ingress termination) |
| **WebSocket authentication** | WebSocket connections require valid JWT before data is transmitted |

### Security Relevant

| Component | Description |
|-----------|-------------|
| **Content Security Policy** | HTTP headers restrict script sources and frame embedding |
| **CORS configuration** | API restricts cross-origin requests to known Console origins |
| **Input sanitization** | User inputs (search queries, mission parameters) sanitized before rendering |
| **Dependency scanning** | Dependabot, npm audit, and CodeQL identify vulnerable dependencies |
| **Secret scanning** | GitHub repository-level secret scanning enabled (not yet integrated into CI/CD workflows) |

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

### Communication Channels

| Channel | Address |
|---------|---------|
| **Internal (maintainers)** | [kubestellar-dev-private@googlegroups.com](mailto:kubestellar-dev-private@googlegroups.com) |
| **Incoming (users)** | GitHub Issues, [#kubestellar-dev Slack](https://cloud-native.slack.com/channels/kubestellar-dev) |
| **Security Announcements** | [kubestellar-security-announce@googlegroups.com](mailto:kubestellar-security-announce@googlegroups.com) |

## Security Issue Resolution

### Responsible Disclosure Process

Vulnerability reports should be sent to [kubestellar-security-announce@googlegroups.com](mailto:kubestellar-security-announce@googlegroups.com). See [SECURITY.md](https://github.com/kubestellar/console/blob/main/SECURITY.md) for full details.

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

1. **AI provider data flow** — Queries sent to external AI providers (Claude, OpenAI, Gemini) may include cluster resource names and metadata. Users should be aware that AI chat queries are processed by third-party services.
2. **Demo mode** — Unauthenticated demo mode uses synthetic data only; no cluster access is possible without authentication.
3. **kc-agent trust model** — The kc-agent runs with the user's local permissions. A compromised kc-agent could access any cluster the user has kubeconfig access to. This is mitigated by the agent running locally (not as a network service).

### Related Security Documentation

- [SECURITY.md](https://github.com/kubestellar/console/blob/main/SECURITY.md) — Vulnerability reporting process
- [SECURITY_CONTACTS](https://github.com/kubestellar/console/blob/main/SECURITY_CONTACTS) — Security response team
- [HARDCODED_URLS.md](https://github.com/kubestellar/console/blob/main/docs/security/HARDCODED_URLS.md) — Explanation of intentional hardcoded URLs
