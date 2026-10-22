# Use case: <system name>

Project or team: <name, link>
Contact: <GitHub handle or Slack name>
Status: <running in production | running in test | design only>
Last updated: <date>

## What it does

Two or three sentences. Who uses it, what the agents do, roughly what scale (agents, requests, clusters).

## Architecture

A short description or a diagram link. Which components are agents, tools, gateways, memory stores, orchestrators.

## Streams it touches

Tick the streams this system exercises and say in a line how.

- [ ] Protocol and communication interoperability:
- [ ] Gateway and data plane:
- [ ] Runtime and container abstraction:
- [ ] State, memory and persistence:
- [ ] Governance, observability, security and policy:

## Standards and patterns exercised

Which parts of the Cloud Native Agentic Standards and which patterns in this document the system uses (MCP, A2A, Gateway API Inference Extensions, OpenTelemetry GenAI conventions, Agent Sandbox, and so on).

## Decisions and why

The choices the team made at the decision points, and the reasons. For example: graph database vs vector store, reuse of an existing PostgreSQL vs a new vector operator, sidecar vs separate Pod, session pinning at the gateway vs in the agent.

## Gaps

Where the standards or this document did not help, or where the team had to invent something. These feed chapter 08.

## References

Code, docs, talks, issues.
