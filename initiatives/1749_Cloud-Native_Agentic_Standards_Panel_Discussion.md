# Cloud Native Agentic Standards Panel Discussion

> **Working draft.** Prep material for a KubeCon NA 2026 co-located session. This is
> collaboration scratch space in a public fork — it is **not** an initiative artifact and is
> not intended for submission to `cncf/toc`. Anything durable that comes out of the panel
> should be PR'd into
> [`1749_Cloud-Native_Agentic_Standards_Checklist.md`](./1749_Cloud-Native_Agentic_Standards_Checklist.md)
> per its [CONTRIBUTING](./1749_Cloud_Native_Agentic_Standards_Checklist_CONTRIBUTING.md).

---

## 1. Session

| | |
|---|---|
| **Title** | Cloud Native Agentic Standards: A Progress Report from the CNCF AI TCG |
| **Event** | Cloud Native AI Day — CNCF-hosted Co-located Event, KubeCon + CloudNativeCon NA 2026 |
| **Format** | Panel |
| **Date / time** | November 9, 2026 |
| **Length** | 35 min |
| **Moderator** | Vincent Caldeira |

### Abstract (as submitted)

As agentic AI expands across the cloud native ecosystem, a lack of uniform standardization
and interoperability threatens secure enterprise scaling.

In this session, co-authors from the CNCF AI TCG provide an essential structural and
operational update based on our recent foundational standards checklist for
Kubernetes-driven workloads. We will break down the core pillars required for reliable
multi-agent systems: communication protocols like MCP and A2A, identity and authorization
frameworks including SPIFFE/SPIRE, granularity in time-series observability metrics, and
flexible governance models designed to mitigate emergent AI behaviors.

Attendees will walk away with an agnostic view of cloud-native best practices to deploy
securely, scale reliably, and ensure explainability across distributed agent networks.

---

## 2. Panelists

| Name | GitHub |
|---|---|
| Vincent Caldeira | [@caldeirav](https://github.com/caldeirav) |
| Josh Halley | [@joshhalley](https://github.com/joshhalley) |
| Nina Polshakova | [@npolshakova](https://github.com/npolshakova) |
| Nimisha Mehta | [@nimishamehta5](https://github.com/nimishamehta5) |

---

## 3. Source material

The panel is a progress report on
[`1749_Cloud-Native_Agentic_Standards_Checklist.md`](./1749_Cloud-Native_Agentic_Standards_Checklist.md).
Its five pillars, for mapping questions:

| § | Pillar | Covers |
|---|---|---|
| General | Container baseline | Least privilege, image provenance, MELT observability, fault tolerance |
| Control and Communication | Protocols | MCP, A2A, AP2, tool sprawl, schema validation, registries/discovery, REST/gRPC/Kafka |
| Observability | Telemetry | Token/TTFT/TPOT/ITL metrics, inference cost, OTel traces + spans, canonical logging |
| Governance | Lifecycle | Pre-deploy evaluation, trajectory assessment, agent-as-a-judge, explainability, MOF |
| Security | Trust | Agent identity (SPIFFE/SVID), tenancy (JIT, ABAC/PBAC), data access |

---

## 4. Draft questions

