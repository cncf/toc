# General Technical Review - Project Copacetic / Sandbox

- **Project:** Copacetic
- **Project Version:** v0.12.0-rc.2
- **Website:** https://project-copacetic.github.io/copacetic/
- **Date Updated:** 2025-10-24
- **Template Version:** v1.0
- **Description:** Copacetic (CLI “copa”) is an open-source tool (written in Go, using BuildKit) that allows direct patching of container image vulnerabilities (at the OS/package level) without full rebuilds. It supports integrating with vulnerability scanners (like Trivy), multiple package managers, distros, and is intended for use in CI/CD pipelines and operational workflows.

---

## Day 0 - Planning Phase

### Scope

#### Roadmap / Feature Planning & Maintainer Ladder
Copacetic has a contributor ladder and governance docs (`CONTRIBUTOR_LADDER.md`, `MAINTAINERS.md`).  
The roadmap is visible via GitHub (release notes, issues) and features are prioritized via user feedback in community calls.  
Mid to long term features are agreed upon by the community and maintainers. As mentioned in the contributor ladder, there are certain contribution metrics that need to be met before becoming a maintainer.

#### Target Persona(s)
- DevSecOps engineers  
- Platform engineers  
- Security teams  

#### Explain the primary use case for the project. What additional use cases are supported by the project?
**Primary Use Case:** Copacetic is primarily used to patch container images directly, currently targeting operating system-level vulnerabilities.  
This provides teams with the ability to patch CVEs without rebuilding from source, which is especially useful for teams that don't rebuild often.  
**Additional Use Cases:** The project is exploring application-level patching and registry-integrated workflows for continuous patching.

#### Unsupported Use Cases
Application-level patching is being explored but is currently unsupported.  
Additionally, End-of-Service-Life (EOSL) images and Windows images are not supported by Copacetic.

#### Organizations who would benefit from adopting this project
Any organization that works with containerized applications could benefit from Copacetic's security focus.

#### End User Research
No formal studies have been published; most community feedback has come primarily through GitHub issues, Slack, and community events like KubeCon.  
Many users have reported successfully running Copacetic in their CI/CD pipelines and registries (thanks to our collaboration with Azure Container Registry).

---

### Usability

#### How do users interact with the project?
Users typically interact through the CLI, GitHub Action, or Docker Desktop Extension.

#### Describe the user experience (UX) and user interface (UI) of the project
The user experience typically revolves around the CLI with flags and JSON/text output.  
There are guides, FAQs, and video tutorials highlighted on the website.

#### Describe how this project integrates with other projects in a production environment
This project integrates with Azure Container Registry for patching container images.  
Azure Container Registry’s continuous patching feature relies on Copacetic as the patching mechanism.

---

### Design

#### Explain the design principles and best practices the project is following
**Best Practices**
- Extensible architecture with different scanner reports or omitting scan reports.  
- Immutability — does not modify the original image; produces a new patched image.  
- Discards patch layer to prevent buildup of patched layers on subsequent patches.  
- Thorough integration and unit testing.

**Copa Design Tenets**
- Accelerate container patching by eliminating dependency on upstream base image updates.  
- Work with the existing ecosystem of container images.  
- Enable parties other than image authors to address container vulnerabilities.  
- Focused, composable design — do one thing well and integrate cleanly with other tools.

#### Outline of the project’s architecture requirements
**Core Components**
- Copa CLI  
- BuildKit (build engine)  
- Container Runtime (Docker/Podman)  
- Package Repositories  
- Optional Vulnerability Scanner  

Documentation: [Copa Design](https://project-copacetic.github.io/copacetic/website/design
):
Copa’s architecture requirements do not differ across environments as long as it can access a local or remote BuildKit engine.

#### Define any specific service dependencies the project relies on in the cluster
- **BuildKit:** Container build engine that executes patching operations  
  - Compatible with BuildKit v0.11+  
  - Can run as: Docker’s built-in BuildKit, standalone daemon, container, buildx builder, or remote instance  
- **OS Package Repositories:** Source of security updates  
- **Container Runtime:** Manages images and registry interactions  
  - Docker Engine v24.0+ (with containerd image store) OR Podman v4.0+  
- **Container Registry:** Stores and retrieves images  
  - Supports Docker Hub, GHCR, Harbor, ECR, ACR, GCR, Quay, Artifactory  
  - Uses Docker credential helpers for authentication  

Note: Copa does not run as a Kubernetes service by default.

#### Describe how the project implements Identity and Access Management
**Authentication Mechanisms**
- **Container Registry Authentication:** Delegates to runtime credentials (`docker login`, `podman login`, etc.)
- **Kubernetes:** Uses image pull secrets when running in-cluster.
- **Cloud Providers:** Supports IAM-based authentication (ECR, ACR, GCR automatic auth).
- **BuildKit Authentication:**  
  - SSH-based: SSH key authentication.  
  - TLS-based: Mutual TLS with certificate-based identity verification (`--tlscert`, `--tlskey`, `--tlscacert`).  
- **File System Permissions:** Standard Unix permissions, Docker group/root for BuildKit socket access.

#### Describe how the project has addressed sovereignty
- **Local Data Processing:** Images remain local where Copa runs.  
- **On-Premises Deployment:** Runs entirely within private infrastructure.  
- **Private Mirrors:** Supports internal or regional package mirrors.  
- **No Vendor Lock-in:** Works with any OCI-compliant regis
#### Describe any compliance requirements addressed by the project
N/A

#### Describe the project’s High Availability requirements
Copa is a CLI tool, not a service.  
Each invocation is stateless and independent.  
Failed operations do not corrupt source images; patch operations are atomic.

#### Describe the project’s resource requirements
Dependent on image size and complexity; CPU, memory, and network usage scale linearly with image layer count.

#### Describe the project’s storage requirements
Uses ephemeral BuildKit cache and temporary build context.  
Storage requirements scale with the image and optional vulnerability reports.

#### Please outline the project’s API Design

**API Topology and Conventions:**  
Copa does not provide a REST or Kubernetes API — it is a CLI tool.

**Defaults:**
- Output tag: `<original-tag>-patched`  
- BuildKit address: Auto-detected (Docker → buildx → local socket).  
- Timeout: 10 minutes.  
- Registry: Same as source image.  
- Push: False by default (`--push` to enable).  
- Patching mode: Comprehensive unless a scan report is provided.

**Additional Configurations:**
- Custom BuildKit Address: `--addr tcp://buildkit.example.com:8443`  
- TLS Configuration: `--tlscert`, `--tlskey`, `--tlscacert`  
- Custom Output Tag: `--tag` or `-t`  
- Push to Registry: `--push`  
- Ignore Errors: `--ignore-errors`  
- Timeout Adjustment: `--timeout 30m`  
- Experimental Features: `EXPERIMENTAL_BUILDKIT_SOURCE_POLICY`  
- Vulnerability Reports: `--r` for targeted patching.  
- Platform-Specific Patching: `--platform`.

**New or Changed APIs:**  
None — Copa introduces no new APIs.

**Versioning and Breaking Changes:**  
Uses semantic versioning (SemVer). Breaking changes only introduced in major versions.

#### Describe the project’s release processes
Releases follow `vX.Y.Z` (semantic versioning).  
- **Major Releases:** Breaking changes.  
- **Minor Releases:** Backward-compatible feature additions.  
- **Patch Releases:** Backward-compatible fixes.

Automated via **GoReleaser** upon tag push.

---

### Installation

#### Describe how the project is installed and initialized
Copa has a minimal install procedure:
- macOS/Linux: `brew install copa`  
- Pre-built binaries via GitHub Releases  
- Build from source: `git clone` and `make`  

**Prerequisites:** Docker or Podman runtime, BuildKit instance, optional scanner (e.g., Trivy).  
Copa automatically detects BuildKit — no config files or setup needed for basic use.

#### How does an adopter test and validate the installation?
**Version Check:**  By running the CLI on any image.


### Security

#### Please provide a link to the project’s cloud native security self assessment
[Copacetic Security Self-Assessment](https://github.com/cncf/tag-security/blob/main/community/assessments/projects/copacetic/self-assessment.md)

#### Please review the Cloud Native Security Tenets from TAG Security. How are you satisfying the tenets of cloud native security projects?
- **Reduce attack surface:** Patches known OS/package CVEs in container images, removing exploitable entry points.  
- **Complement, not replace:** Works with scanners (e.g., Trivy, Grype) via an extensible scanner interface; does not replace other controls.  
- **Rapid response:** Enables quick mitigation when new CVEs are disclosed.  
- **No unnecessary change:** Does not modify application code/configs.  
- **Immutability:** Produces new images rather than mutating existing ones.  
- **Transparent layers:** Patch layer is distinct and auditable.  
- **CI/CD-first:** Ships a GitHub Action and Docker Desktop Extension for automation.

#### Describe how each of the cloud native principles apply to your project
- **Containers:** Container-native patching workflow.  
- **Immutable infrastructure:** Always creates a new image; never edits in place.  
- **Resilience:** Shrinks exposure window by quickly removing known CVEs.  
- **Manageability:** Simple CLI, clear outputs and flags.  
- **Observability:** Logs and resulting layers are inspectable.  
- **Automation:** Easy integration into CI/CD pipelines.  
- **Dynamic environments:** Cloud/vendor agnostic, works with any OCI registry.

#### How do you recommend users alter security defaults in order to "loosen" the security of the project?
- **Insecure BuildKit connections:** Allow unencrypted TCP when necessary (see docs: [custom BuildKit address](https://project-copacetic.github.io/copacetic/website/custom-address)).  
- **Custom/untrusted repositories:** Replace official package repos with mirrors or private sources as needed (see FAQ: [replace package repositories](https://project-copacetic.github.io/copacetic/website/faq#can-i-replace-the-package-repositories-in-the-image-with-my-own)).

#### Security Hygiene
- CNCF Code of Conduct and contribution guidelines.  
- Built on CNCF BuildKit.  
- Automated dependency/security tooling: Dependabot, Snyk.  
- Unit/integration/lint tests; CodeQL for static analysis; CodeCov for coverage.

#### Describe how the project has evaluated which features will be a security risk to users if they are not maintained by the project
- Security-critical areas: patching engine, BuildKit integration, scanner report ingestion.  
- These are actively maintained and prioritized when issues arise (Slack/GitHub).  
- Risks like outdated dependency handling or incorrect scan parsing are triaged quickly.

#### Cloud Native Threat Modeling

##### Explain the least minimal privileges required by the project and reasons for additional privileges
- **Container runtime:** Read/pull images; write/push patched images (Docker/Podman socket).  
- **BuildKit:** Local/remote socket or TCP connection (rootless supported).  
- **Package repos:** Egress to fetch updates; credentials only for private repos.  
- **Filesystem:** Read scan reports; temporary build context writes.  
- **Additional:** Registry auth when pushing to private registries.

##### Describe how the project is handling certificate rotation and mitigates any issues with certificates
- Certificate lifecycle is delegated to the environment (BuildKit, runtime, registry).  
- Copa stores no certs; honors system and endpoint TLS configuration (`--tlscert`, `--tlskey`, `--tlscacert` when applicable).

##### Describe how the project is following and implementing secure software supply chain best practices
- **Immutable artifacts & digests.**  
- **Verification & rescanning** of patched images in CI/CD.  
- **Automated deps:** Dependabot.  
- **Static analysis:** CodeQL; Snyk for vuln detection.  
- **Releases:** GoReleaser for reproducible, consistent builds.

---

## Day 1 - Installation and Deployment Phase

### Project Installation and Configuration

#### Describe what project installation and configuration look like
- **Install:**  
  - macOS/Linux: `brew install copa`  
  - GitHub Releases: download prebuilt binaries  
  - From source: `make` / `make install`  
  - Docker Desktop Extension: GUI, zero-config for local dev  
- **Prerequisites:** Docker or Podman; a BuildKit instance; optional scanner (e.g., Trivy) for targeted patching.  
- **Configuration:** CLI flags on `copa patch`  
  - `--addr` (BuildKit address), `-i` (input image), `-r` (scan report/file or directory for multi-arch), `-t` (output tag), `--push`, `--platform`, `--timeout`, `--ignore-errors`.

---

### Project Enablement and Rollback

#### How can this project be enabled or disabled in a live cluster? Please describe any downtime required of the control plane or nodes
- **N/A** — Copacetic is a CLI/tooling component, not a control plane/node daemon. No cluster downtime implications.

#### Describe how enabling the project changes any default behavior of the cluster or running workloads
- **N/A** — Operates out-of-band on images; it does not mutate cluster defaults or running pods.

#### Describe how the project tests enablement and disablement
- **N/A** — Not a cluster-resident service. Patched images are validated through standard CI/CD (build → scan → tests) before deployment.

#### How does the project clean up any resources created, including CRDs?
- **N/A** — Creates no CRDs or cluster resources. Artifacts:  
  - New patched image (local cache/registry).  
  - Optional OpenVEX file (local file, user-managed).

---

### Rollout, Upgrade and Rollback Planning

#### How does the project intend to provide and maintain compatibility with infrastructure and orchestration management tools like Kubernetes and with what frequency?
- By adhering to **OCI image** standards. No special integration needed; compatibility is continuous by design.

#### Describe how the project handles rollback procedures
- **Tool rollback:** reinstall prior Copacetic version (Homebrew or GitHub Releases).  
- **Workload rollback:** redeploy previous **image digest** (immutable reference). Copa never mutates existing images.

#### How can a rollout or rollback fail? Describe any impact to already running workloads
- **Rollout can fail if:** package dependency conflicts break patching; registry push fails (auth/network); functional tests fail on the patched image.  
- **Rollback can fail if:** prior tag/digest has been deleted/garbage-collected; orchestrator references a mutable tag instead of a digest.  
- **Impact:** Running workloads are **unaffected**; Copa only creates new images. Changes apply on redeploy.

#### Describe any specific metrics that should inform a rollback
- **Vulnerability scan deltas:** regressions or expected CVE reductions not realized.  
- **CI/CD outcomes:** functional/integration test failures post-patch.  
- **Copa/build signals:** non-zero exit codes; failed package installs; error logs.

#### Explain how upgrades and rollbacks were tested and how the upgrade→downgrade→upgrade path was tested
- Regression suites across multiple distros/images.  
- **Idempotency checks** (re-patching same image yields identical output).  
- Cross-version validation: patch with N-1, verify with N; downgrade and re-upgrade to confirm reproducibility and artifact integrity.  
- Manual verification using released binaries to mirror operator flows.

#### Explain how the project informs users of deprecations and removals of features and APIs
- **GitHub Releases** and **`CHANGELOG.md`**.  
- CLI deprecation warnings.  
- CNCF Slack/mailing list announcements.  

#### Explain how the project permits utilization of alpha and beta capabilities as part of a rollout
- Experimental features are **opt-in** (e.g., `--experimental` flags or env vars), clearly labeled in docs/release notes, allowing early trials without impacting stable users.