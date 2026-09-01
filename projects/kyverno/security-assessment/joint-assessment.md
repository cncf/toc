# Kyverno TAG-Security Joint Assessment

**Completed:** TBD

**Security reviewers:** Andrew Martin, John Kinsella, Wesley Steehouwer (@dutchshark), Robert Ficcaglia, Tom Cope, Giovanni Baggio, Justin Cappos

**Project security lead:** Jim Bugwadia, Shuting Zhao

**Project team:** Jim Bugwadia, Shuting Zhao

- **Source code:** https://github.com/kyverno/kyverno
- **Website:** https://kyverno.io/
- **Tracking issue:** https://github.com/cncf/toc/issues/1703
- **Changes tracking issue:** https://github.com/kyverno/kyverno/issues/15335

## Background

Kyverno is a policy engine for Kubernetes that enforces security, operational, and best practice policies across Kubernetes resources. It operates as an admission controller with validating and mutating webhooks, and provides CLI tools for offline policy evaluation. Kyverno is designed to prevent insecure configurations, ensure compliance through user-defined policies, and automate resource management via mutation and generation capabilities.

The project has expanded beyond Kubernetes admission control to support policy validation for non-Kubernetes resources (Terraform, Dockerfiles, cloud configurations, and others) through new policy types (`ValidatingPolicy`, `ImageValidatingPolicy`, `MutatingPolicy`) that utilize CommonExpressionLanguage (CEL).

### Maturity

Kyverno is in the CNCF incubation stage. Notable features include:

- **Admission control:** Policy enforcement via validating and mutating webhooks
- **Image verification:** Signature and attestation verification using Sigstore Cosign and Notary
- **Background scanning:** Periodic policy evaluation and resource mutation via cron schedules
- **Policy reports:** Policy violation and compliance reporting via `PolicyReport` and `ClusterPolicyReport` resources
- **Resource generation and cleanup:** Automated resource creation and deletion based on policies
- **CLI tools:** Off-cluster policy evaluation and testing for CI/CD pipelines
- **Multi-tenant support:** Policy Exceptions for narrowly-scoped, auditable policy bypasses

## Summary

### Design

Kyverno comprises four main controllers handling distinct aspects of policy enforcement:

1. **Admission Controller** - Receives admission requests from the Kubernetes API server; processes validation, mutation, and image verification rules; manages webhook configurations dynamically
2. **Reports Controller** - Generates and reconciles final policy reports from ephemeral intermediate resources
3. **Background Controller** - Processes generate and mutate-existing rules on existing cluster resources
4. **Cleanup Controller** - Processes cleanup and deletion policies via cron-triggered jobs

The project demonstrates strong foundational security architecture, including:
- Separation of concerns across controllers with fine-grained RBAC
- Certificate management and renewal for webhook TLS
- Circuit breakers to prevent unbounded resource growth
- Support for both CLI and admission control enforcement paths
- Comprehensive policy reporting and exception tracking

### Analysis

The joint assessment found a well-maintained project with good security practices:

**Strengths:**
- Third-party security audits completed (Ada Logics, Nov 2023)
- SBOM generation and availability for releases
- Container image signing via Sigstore Cosign
- Supply chain security practices aligned with SLSA 3
- Extensive threat model documentation referencing Kubernetes SIG Security admission controller threat model
- Responsive project team and comprehensive documentation

**Areas requiring attention:**
- **Global Context Cache:** Cross-namespace access and lack of explicit cache bounds create potential for data leakage; cache is in-memory and shared across policies.
- **External Data Lookups:** Documentation should more explicitly emphasize risks of network calls and recommend authentication/encryption for all external service integrations 
- **CEL Cost Budget:** `ValidatingPolicy` and `ImageValidatingPolicy` do not enforce RuntimeCELCostBudget, unlike classic admission validation
- **Webhook Authentication:** Default installation does not authenticate API server requests to admission webhooks
- **Network Policy:** Default installation includes no network policy; Kyverno can access external systems but lacks default outbound restrictions
- **Helm Chart Signing:** Helm charts are not signed with Sigstore Cosign
- **CLI Threat Model:** No formal threat model for CLI or offline evaluation scenarios

## Recommendations

### CNCF Recommendations

1. **Global Context Cache Restrictions** - Implement namespace-scoped access controls and configurable cache bounds to limit data leakage between tenants (see [kyverno/kyverno#15359](https://github.com/kyverno/kyverno/issues/15359))
2. **Webhook Request Authentication** - Add built-in support for authenticating API server requests to admission webhooks as per Kubernetes best practices (see [kyverno/kyverno#16559](https://github.com/kyverno/kyverno/issues/16559))
3. **CEL Cost Budget Alignment** - Enforce RuntimeCELCostBudget consistently across all CEL policy types to prevent resource exhaustion attacks (see  [kyverno/kyverno#14495](https://github.com/kyverno/kyverno/issues/14495))
4. **Threat Model Mapping** - Map all published security advisories to the Kubernetes SIG Security admission controller threat model to close coverage gaps (see [kyverno/website#2091](https://github.com/kyverno/website/issues/2091)-Done)
5. **Helm Chart Signing** - Sign Helm charts with Sigstore Cosign to ensure supply chain security alignment across all distribution mechanisms
6. **SAST Integration** - Continue implementation of static analysis security testing (SAST) tools like Semgrep in CI/CD pipelines (see [PR #16464](https://github.com/kyverno/kyverno/pull/16464))

Tracked in [kyverno/kyverno#15335](https://github.com/kyverno/kyverno/issues/15335)

### Project Recommendations

1. **External Data Lookup Documentation** - Expand security documentation to explicitly call out risks of network calls and prominently recommend authentication and encryption for all external service integrations (see [kyverno/website PR #2086](https://github.com/kyverno/website/pull/2086)-Done)
2. **Network Policy Guidance** - Provide default network policy templates and expand documentation on risks of running without network policies; offer CLI command to auto-generate platform-specific network policies (see website PRs #2078, #2085)
3. **Architecture Documentation** - Update architecture diagrams to clearly separate logical components from physical deployment architecture (see website PR #2083)
4. **Security Advisories Visibility** - Link to published security advisories from the main security documentation page (see website PR #2017)
5. **CLI Threat Model** - Develop threat model for CLI use cases and non-admission-control deployment scenarios
6. **Blog Post Links** - Fix outdated blog post links in security documentation (see issue #15336)

### Additional Recommendations

1. Consider implementing additional SAST tools (Nancy, Trivy) alongside Semgrep for comprehensive vulnerability detection
2. Establish automated security advisory review and threat model update process
3. Consider threat model refresh against SLSA 1.0 specification (currently documented against v0.1)
4. Document and provide examples for secure configuration in high-security environments (e.g., air-gapped clusters, restricted networks)
5. Implement periodic security advisory backtesting against threat model to ensure completeness of coverage

## Acknowledgments

The Kyverno team was responsive, engaged, and thoughtful throughout the assessment process. The comprehensive security documentation and third-party audit results demonstrate a commitment to security-first development practices.
