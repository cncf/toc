# 3. Gateway and data plane

Status: outline. Owner: Akanksha Trehun.

- Agent-to-agent and agent-to-tool routing
- Gateway API Inference Extensions (InferencePool) and session-aware fan-out
- Bidirectional SSE, JSON-RPC streaming and multiplexing, proxy extensions
- Framing under review: the waypoint-style sidecarless L7 architecture transfers to agent traffic; the protocol awareness does not. Session pinning must key on the MCP session ID, server-initiated push breaks the client-initiates assumption, and per-tool authorization needs JSON-RPC payload introspection beyond HTTPRoute matching
- Open question: conformance behaviors as an extension contract on Gateway API primitives, or a new session-aware primitive

