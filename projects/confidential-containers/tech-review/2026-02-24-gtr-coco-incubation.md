# General Technical Review - Confidential Containers (CoCo) / Incubation

- **Project:** Confidential Containers (CoCo)
- **Project Version:** v0.18.0
- **Website:** https://confidentialcontainers.org
- **Date Updated:** 2026-02-24
- **Template Version:** v1.0
- **Description:** Confidential Containers is an open-source CNCF Sandbox project that integrates Trusted Execution Environments (TEEs) into Kubernetes. It encapsulates pods inside confidential virtual machines (VMs), allowing cloud-native workloads to leverage hardware-based confidential computing to protect data-in-use with minimal to no application modification.

## Day 0 - Planning Phase

### Scope

  * Describe the roadmap process, how scope is determined for mid to long term features, as well as how the roadmap maps back to current contributions and maintainer ladder?
  
    The project adopts a [use-case driven development approach](https://confidentialcontainers.org/blog/2024/02/16/introduction-to-confidential-containers-coco/#foundational-principles), focusing community efforts on key scenarios rather than feature-centric models. The roadmap and scope are determined via consensus in the [Steering Committee and weekly community meetings](https://docs.google.com/document/d/1E3GLCzNgrcigUlgWAZYlgqNTdVwiMwCRTJ0QnJhLZGA/edit), with changes documented in the public [roadmap.md](https://github.com/confidential-containers/confidential-containers/blob/main/roadmap.md).

  * Describe the target persona or user(s) for the project?
  
    The project defines five distinct trust boundaries and [cloud native personas](https://confidentialcontainers.org/docs/architecture/trust-model/cloud-native-personas/):
    1. Infrastructure Operator (Cloud Provider)
    2. Orchestration Operator (Kubernetes Admin)
    3. Workload Provider (Software vendor providing the container image)
    4. Container Image Provider (note: there may be [expected overlap](https://confidentialcontainers.org/docs/architecture/trust-model/cloud-native-personas/#container-image-provider))
    5. Data Owner (The end-user whose data is processed)
    
    ![CoCo cloud personas](https://confidentialcontainers.org/persona_model.svg)
    _from: https://confidentialcontainers.org/persona_model.svg_
        
    CoCo protects the Workload Provider and Data Owner from both the Infrastructure and Orchestration Operators.

  * Explain the primary use case for the project. What additional use cases are supported by the project?

    [Confidential AI / Model Inferencing](https://confidentialcontainers.org/blog/2024/04/15/memory-protection-for-ai-ml-model-inferencing), allowing organizations to deploy proprietary LLMs or process sensitive data on untrusted infrastructure without risk of IP theft or memory inspection. Additional use cases include Secure Supply Chain / Confidential CI/CD (building and signing OCI images in a TEE), Federated Learning, and [Crypto Oracles for Web3](https://confidentialcontainers.org/blog/2025/03/11/how-switchboard-oracles-leverage-confidential-containers-for-next-generation-web3-security).

  * Explain which use cases have been identified as unsupported by the project.
  
    As outlined in [Trust Model::Out of Scope](https://confidentialcontainers.org/docs/architecture/trust-model/trust-model/#out-of-scope), CoCo inherits the properties of the underlying hardware TEE. It explicitly **does not protect against**:
    * Hardware side-channel attacks.
    * Effective Denial of Service (DoS) by the host and/or orchestrator (the untrusted control plane can simply refuse to schedule or execute the VM).

  * Describe the intended types of organizations who would benefit from adopting this project. (i.e. financial services, any software manufacturer, organizations providing platform engineering services)?
  
    Highly regulated industries (finance, healthcare, government), AI/ML practitioners protecting proprietary models/data, and Web3 infrastructure providers (e.g., Switchboard), CI/CD vendors, Kubernetes service providers wishing to provide TEE as a differentiator.

  * Please describe any completed end user research and link to any reports.

    CoCo maintains a public [`ADOPTERS.md`](https://github.com/confidential-containers/confidential-containers/blob/main/ADOPTERS.md#confidential-containers-adopters). Rather than attempting to summarize and make choices about what to add, we've reproduced what's contained in `ADOPTERS.md`.  Please see that file for additional information regarding Adoper types (https://github.com/confidential-containers/confidential-containers/blob/main/ADOPTERS.md#adopter-types).

| Organization/Company                                              | Project/Product                                               | Usage level              | Adopter type                     | Details                                                                   |
|-------------------------------------------------------------------|---------------------------------------------------------------|--------------------------|----------------------------------|---------------------------------------------------------------------------|
|[Alibaba Cloud (Aliyun)](https://www.alibabacloud.com/)| [Elastic Algorithm Service](https://www.alibabacloud.com/help/en/pai/user-guide/eas-model-serving/?spm=a2c63.p38356.0.0.2b2b6679Pjozxy) and [Elastic GPU Service](https://www.alibabacloud.com/help/en/egs/) | Beta | Service Provider | Both services use sub-projects of confidential containers to protect the user data and AI model from being exposed to CSP (For details mading.ma@alibaba-inc.com) |
| [Edgeless Systems](https://www.edgeless.systems/)                 | [Contrast](https://github.com/edgelesssys/contrast)           | Beta                     | Service Provider / Consultancy   | Contrast runs confidential container deployments on Kubernetes at scale.                                |
| [IBM](https://www.ibm.com/z)                                    | [IBM LinuxONE](https://www.ibm.com/linuxone)                  | Beta                     | Service Provider                 | Confidential Containers with Red Hat OpenShift Container Platform and IBM® Secure Execution for Linux (see [details](https://www.ibm.com/blog/confidential-containers-with-red-hat-openshift-container-platform-and-ibm-secure-execution-for-linux/))  |
|NanhuLab|Trusted Big Data Sharing System |Beta |Service Provider |The system uses confidential containers to ensure that data users can utilize the data without being able to view the raw data.(No official website yet. For details: yzc@nanhulab.ac.cn) |
| [KubeArmor](https://www.kubearmor.io/) | Runtime Security | Beta | Another project | An open source project that leverages CoCo as part of their solution, integrates with for compatibility and interoperability, or is used in the supply chain of another project [(5GSEC)](https://github.com/5GSEC/nimbus/blob/main/examples/clusterscoped/coco-workload-si-sib.yaml). |
| [Red Hat](https://www.redhat.com/en) | [OpenShift confidential containers](https://www.redhat.com/en/blog/learn-about-confidential-containers) | Beta | Service Provider | Confidential Containers are available from [OpenShift sandboxed containers release version 1.7.0](https://docs.redhat.com/en/documentation/openshift_sandboxed_containers/1.7/) as a tech preview on Azure cloud for both Intel TDX and AMD SEV-SNP. The tech preview also includes support for confidential containers on IBM Z and LinuxONE using Secure Execution for Linux (IBM SEL).|
| [ByteDance](https://www.bytedance.com/) | Jeddak Sandbox | Beta | End-User / Service Provider | Jeddak Sandbox leverages CoCo to protect the data privacy in the process of the company's business  (for details chendian.imtyrant@bytedance.com) |
| [Intel](https://www.intel.com/) | [Enterprise-RAG](https://github.com/opea-project/Enterprise-RAG/blob/main/docs/tdx.md)<br>[OPEA](https://github.com/opea-project/GenAIInfra/tree/main/helm-charts/TDX.md) | Beta | End-User / Service Provider | Intel runs confidential container deployments on Kubernetes with Intel TDX |
| [Switchboard](https://www.switchboard.xyz/) | [Decentralized Crypto Oracle](https://docs.switchboard.xyz/) | Production | Service Provider | Running our Decentralized Oracles code in CoCo on AMD SEV SNP bare metal machines |
|[JDCloud](https://www.jdcloud.com)|JoyScale |Beta |End-User / Service Provider  | JoyScale leverages CoCo to protect the AI data privacy in the process of the company's business and end user. (For details: huoqifeng1@jd.com)|
|[Kubermatic](https://www.kubermatic.com/)| Kubeone | Beta | Service Provider / Consultancy | Running confidential containers on baremetal kubeone clusters. |

   **TODO (Maintainers):** Please provide a brief summary or links to any additional adopter interviews, user surveys, or formal UX research (if any) conducted during the Sandbox phase.

### Usability

  * How should the target personas interact with your project?
  
    Users interact via standard Kubernetes primitives. Enabling CoCo for a workload simply requires specifying the appropriate `runtimeClassName` (e.g., `kata-qemu-tdx`, `kata-remote`) and relevant annotations on standard Pod specifications. Advanced users interact with the [KBS Client Tool](https://confidentialcontainers.org/docs/attestation/client-tool/) to provision resource and attestation policies.

  * Describe the user experience (UX) and user interface (UI) of the project.
  
    The project emphasizes transparent deployment of unmodified containers without providing it's own "UI." The UX is CLI and GitOps driven, utilizing standard k8s mechanisms for workload deployment and Helm for infrastructure deployment. There was an operator but it has recently been deprecated on 2026-02-01, ([deprecated operator repo](https://github.com/confidential-containers/operator/blob/main/README.md)) in favor of a [Helm chart](https://github.com/confidential-containers/charts?tab=readme-ov-file#confidential-containers-helm-chart).

  * Describe how this project integrates with other projects in a production environment.
  
    CoCo natively integrates with [Kata Containers](https://katacontainers.io/) for VM isolation, standard CRI runtimes (`containerd` 1.7+ / `CRI-O`), and snapshotters (like `nydus-snapshotter`). It integrates seamlessly with higher-level orchestration tools like KServe for AI workloads.

### Design

  * Explain the design principles and best practices the project is following.   
  
    **Pod-Centric Virtualization:** Balances the Trusted Computing Base (TCB) size by putting the Pod inside the enclave (rather than a single container or a whole Node), enabling secure local container networking. See [Design Overview](https://confidentialcontainers.org/docs/architecture/design-overview/).

    ![pod-centric-virt](https://confidentialcontainers.org/img/coco-design-diagram-1.png)

    **Host Deprivileging:** Images are pulled and unpacked *inside* the guest enclave via `image-rs`, preventing the host from accessing container filesystems. See [Design Overview::Image Pulling](https://confidentialcontainers.org/docs/architecture/design-overview/#image-pulling)

    ```mermaid
    sequenceDiagram
    kubelet->>containerd: create container
    containerd->>nydus snapshotter: load container snapshot
    nydus snapshotter->>image-rs: download container
    kubelet->>containerd: start container
    containerd->>kata shim: start container
    kata shim->>kata agent: start container
    ```

    **Attestation (Trustee) Architecture:** Decouples attestation and secret delivery from the workload orchestration.

    * high level synopsis: https://confidentialcontainers.org/docs/architecture/design-overview/#attestation
    * Trustee Attestation (arch): https://confidentialcontainers.org/docs/attestation/architecture
        * Key Broker Service (KBS) protocol: [kbs_attestation_protocol.md](https://github.com/confidential-containers/trustee/blob/main/kbs/docs/kbs_attestation_protocol.md)

    from the (very well written) docs:


    > Confidential Containers also provides components inside the guest and elsewhere to facilitate attestation. Attestation is a crucial part of confidential computing and a direct requirement of many guest operations. For example, to unpack an encrypted container image, the guest must retrieve a secret key. Inside the guest the confidential-data-hub and attestation-agent handle operations involving secrets and attestation. Again, these components are beyond the scope of traditional Kata deployments and are located in the guest components repository.
    >
    >The CDH and AA use the KBS Protocol to communicate with an external, trusted entity. Confidential Containers provides Trustee as an attestation service and key management engine that validates the guest TCB and releases secret resources.

    ```mermaid
    sequenceDiagram
    workload->>CDH: request secret
    CDH->>AA: get attestation token
    AA->>KBS: attestation request
    KBS->>AA: challenge
    AA->>KBS: attestation
    KBS->>AA: attestation token
    AA->>CDH: attestation token
    CDH->>KBS: secret request
    KBS->>CDH: encrypted secret
    CDH->>workload: secret
    ```

    Finally, the docs for Trustee / Attestation contain a diagram worth reproducing here entitled "putting the peices together" ([design-overview.md#putting-the-pieces-together](https://github.com/confidential-containers/confidentialcontainers.org/blob/main/content/en/docs/architecture/design-overview.md#putting-the-pieces-together))
    
    > If we take Kata Containers and add guest image pulling and attestation, we arrive at the following diagram, which represents Confidential Containers.
    
    ```mermaid
    flowchart TD
        kubelet-->containerd
        containerd-->kata-shim
        kata-shim-->hypervisor
        containerd-->nydus-snapshotter

        subgraph guest
            kata-agent-->pod
            ASR-->CDH
            CDH<-->AA
            pod-->ASR
            image-rs-->pod
            image-rs-->CDH
        end
        subgraph pod
            container-a
            container-b
            container-c
        end
        subgraph Trustee
            KBS-->AS
            AS-->RVPS
        end

    AA-->KBS
    CDH-->KBS
    nydus-snapshotter-->image-rs
    hypervisor-->guest
    kata-shim<-->kata-agent
    image-rs<-->registry

    ```

  * Outline or link to the project’s architecture requirements? Describe how they differ for Proof of Concept, Development, Test and Production environments, as applicable.  
  
    [Hardware Requirements](https://confidentialcontainers.org/docs/getting-started/prerequisites/hardware/): 
    
    Production requires a bare-metal node supporting TEEs (Intel TDX, AMD SEV-SNP, IBM Secure Execution) or Cloud APIs for "Peer Pods". There are runtime classes for all of these defined:

    ```bash
    $ kubectl get runtimeclasses
    NAME             	HANDLER          	AGE
    kata             	kata-qemu        	26m
    kata-clh         	kata-clh         	26m
    kata-qemu        	kata-qemu        	26m
    kata-qemu-coco-dev    kata-qemu-coco-dev    26m
    kata-qemu-sev    	kata-qemu-sev    	26m
    kata-qemu-snp    	kata-qemu-snp    	26m
    kata-qemu-tdx    	kata-qemu-tdx    	26m

    ```
    
    Development/PoC environments can run on standard VMs using the `kata-qemu-coco-dev` runtime class, which emulates the stack without hardware TEE guarantees. The docs contain a great blog (["Confidential Containers without confidential hardware"](https://confidentialcontainers.org/blog/2024/12/03/confidential-containers-without-confidential-hardware/)) that describes this in detail.

  * Define any specific service dependencies the project relies on in the cluster.
  
    Requires `containerd` (1.7+) or `CRI-O`, and standard K8s control plane components. The Nydus snapshotter is required to route image pulls to the guest. Remote attestation relies on the external [Trustee](https://confidentialcontainers.org/docs/attestation/) (KBS, AS, RVPS) services.

  * Describe how the project implements Identity and Access Management.  
  
    Control plane installation uses Kubernetes RBAC. Inside the enclave, identity is verified cryptographically via [Attestation Evidence](https://confidentialcontainers.org/docs/features/get-attestation/) matching reference values. Secrets are bound to the specific verified state of the hardware and software via KBS resource policies.
    
    **Of Note** (from looking at the community meeting notes), Additionally, the community is actively working on "Trustworthy Workload Identity" (TWI) architecture to bridge hardware TEE evidence with cloud-native identity standards like SPIFFE/SPIRE, allowing CoCo workloads to securely authenticate to external services (like databases or vaults) 
    
  * Describe how the project has addressed sovereignty.  
  
    Sovereignty is the core value proposition: Data owners maintain absolute cryptographic control over their data and workloads, even when running on third-party public clouds, by relying on hardware-rooted trust rather than cloud provider trust.

  * Describe any compliance requirements addressed by the project.  
  
    Provides technical enforcement for data-in-use protection, assisting organizations with GDPR, HIPAA, and strict financial compliance requirements by guaranteeing workload confidentiality and integrity.

  * Describe the project’s High Availability requirements.  
  
    The in-cluster CoCo deployment scales automatically with Kubernetes worker nodes. The external `Trustee` service can be configured for High Availability by scaling the deployment replicas (`KbsDeploymentSpec.replicas`) via the [Trustee Operator](https://confidentialcontainers.org/blog/2026/02/11/deploy-trustee-in-kubernetes/).

  * Describe the project’s resource requirements, including CPU, Network and Memory.  
  
    Worker nodes require virtualization support and a recommended minimum of 8GB RAM and 4 CPUs to accommodate the hypervisor/Kata overhead.

  * Describe the project’s storage requirements, including its use of ephemeral and/or persistent storage.  
  
    Memory is protected natively by the TEE. Because memory is an expensive resource, CoCo implements [Trusted Ephemeral Storage](https://github.com/confidential-containers/confidential-containers/blob/main/guides/ephemeral_storage.md) using LUKS2 encryption via `dm-crypt` for container images and Read/Write layers, backed by host block devices (e.g., via `open-local` CSI).

  * Please outline the project’s API Design:  
    * Describe the project’s API topology and conventions  
    
      Extends Kubernetes via standard Custom Resource Definitions (CRDs) (e.g., `TrusteeConfig`, `KbsConfig`). Inside the pod, the Confidential Data Hub (CDH) exposes a local REST API (e.g., `127.0.0.1:8006/cdh/resource`) for workload interaction.
      
    * Describe the project defaults  
    
      The project operates on a "secure by default" posture, enforcing strict isolation at multiple API boundaries:

      * Host-to-Guest API: The Kata Agent API inside the enclave utilizes an OPA policy that is "deny-by-default" for invasive operations (e.g., default ExecProcessRequest := false), preventing untrusted orchestrator admins from executing arbitrary commands in the guest.
      * Workload-facing APIs: The internal REST API used to fetch raw attestation reports is explicitly disabled by default (agent.guest_components_rest_api must be manually enabled).
      * Secret Release: The Trustee Key Broker Service (KBS) defaults to a policy that secrets are only released to cryptographically verified TEEs.
      * Image Management: By default, container image pulling is diverted away from the host Kubelet/Containerd and forced inside the guest enclave via image-rs.
      
    * Outline any additional configurations from default to make reasonable use of the project  
    
      Users provide [Init-Data](https://confidentialcontainers.org/docs/features/initdata/) via pod annotations to inject specific Agent Policies and specify KBS URLs.
      
    * Describe any new or changed API types and calls \- including to cloud providers \- that will result from this project being enabled and used  
    
      The Cloud API Adaptor (Peer Pods) translates Kubelet pod creation into Cloud Provider VM creation APIs (AWS EC2, Azure VMs, GCP Compute).
      
    * Describe compatibility of any new or changed APIs with API servers, including the Kubernetes API server   
    
      APIs follow standard Kubernetes semantic versioning.
      
    * Describe versioning of any new or changed APIs, including how breaking changes are handled  
    
      The project follows semantic versioning and clearly communicates breaking changes in release notes.

  * Describe the project’s release processes, including major, minor and patch releases.
  
    The project follows a strict [6-week release cadence](https://github.com/confidential-containers/confidential-containers/blob/main/README.md) aligned closely with the Kata Containers lifecycle.

### Installation

  * Describe how the project is installed and initialized, e.g. a minimal install with a few lines of code or does it require more complex integration and configuration?  
  
    Installation is highly streamlined using [Helm charts](https://confidentialcontainers.org/docs/getting-started/installation/). (Note: The legacy Operator has been explicitly [deprecated](https://github.com/confidential-containers/operator/blob/main/README.md) in favor of Helm). A single Helm command (`helm install coco oci://ghcr.io/confidential-containers/charts/confidential-containers`) installs the daemonsets, kata-deploy binaries, and registers the necessary `RuntimeClasses` based on the target architecture. The Attestation components are deployed separately via the [Trustee Operator](https://confidentialcontainers.org/blog/2026/02/11/deploy-trustee-in-kubernetes/) or Docker Compose.

  * How does an adopter test and validate the installation?
  
    Adopters test installation by verifying the presence of the `kata-qemu-*` runtime classes, then deploying a sample pod referencing that `runtimeClassName`. Validation of actual hardware encryption can be done by parsing the container's `dmesg` logs for TEE active messages (e.g., `Memory Encryption Features active: AMD SNP`). See the [Troubleshooting Guide](https://github.com/confidential-containers/confidential-containers/blob/main/guides/troubleshooting.md).

### Security

  * Please provide a link to the project’s cloud native [security self assessment](https://tag-security.cncf.io/community/assessments/).  
  
    [Confidential Containers Self Assessment](https://tag-security.cncf.io/community/assessments/projects/confidential-containers/self-assessment/)

  * Please review the [Cloud Native Security Tenets](https://github.com/cncf/tag-security/blob/main/community/resources/security-whitepaper/secure-defaults-cloud-native-8.md) from TAG Security.  
    * How are you satisfying the tenets of cloud native security projects?  
    
      **Secure by Default:** The architecture denies host access to container images (pulled inside the guest) and ensures secrets are only released if attestation succeeds.
      
    * Describe how each of the cloud native principles apply to your project.  
    
      **Zero Trust:** Trust is established via hardware measurements and cryptographic attestation, not network perimeter or node identity.
      
    * How do you recommend users alter security defaults in order to "loosen" the security of the project? Please link to any documentation the project has written concerning these use cases.  
    
      Users configure the [Three Policy Layers](https://confidentialcontainers.org/docs/attestation/policies/) (Agent Policy, Resource Policy, Attestation Policy). For testing, users can "loosen" defaults by using the `kata-qemu-coco-dev` runtime class or setting permissive OPA policies (e.g., allowing `ExecProcessRequest` in the Kata Agent).

  * Security Hygiene
    * Please describe the frameworks, practices and procedures the project uses to maintain the basic health and security of the project.
    
      The project achieved the [OpenSSF Best Practices passing badge](https://www.bestpractices.dev/projects/5719). It strictly enforces DCOs, uses automated linting/static analysis (`rustfmt`, `clippy`), and maintains a private vulnerability reporting pipeline defined in the org-level `SECURITY.md`.
      
    * Describe how the project has evaluated which features will be a security risk to users if they are not maintained by the project?
    
      The project heavily scrutinizes the [Kata Agent API](https://confidentialcontainers.org/docs/architecture/trust-model/trust-model/#crossing-the-trust-boundary), enforcing OPA policies to block endpoints that could allow host interference (e.g., restricting `ExecProcessRequest` and `CreateContainerRequest`).

  * Cloud Native Threat Modeling
    * Explain the least minimal privileges required by the project and reasons for additional privileges.  
    
      The Kata Agent API is heavily restricted using OPA policies (Rego) ensuring the host Kubelet cannot perform arbitrary commands or mounts into the enclave unless explicitly whitelisted by the workload owner.
      
    * Describe how the project is handling certificate rotation and mitigates any issues with certificates.
    
      **TODO (Maintainers):** Please describe the mechanisms for rotating the internal TLS/mTLS certificates used between Trustee, the CDH, and the Attestation Agent.
      
      * The Trustee service uses trustee-operator, which (TODO: validate) relies on cert-manager. 
      
    * Describe how the project is following and implementing [secure software supply chain best practices](https://project.linuxfoundation.org/hubfs/CNCF\_SSCP\_v1.pdf) 
    
      The project has achieved SLSA Build Level 2 (see [blog](https://confidentialcontainers.org/blog/2025/02/17/confidential-containerscoco-and-supply-chain-levels-for-software-artifacts-slsa), automatically generating signed provenance in `in-toto` format via GitHub Actions for components like `kata-containers`, `guest-components`, and `cloud-api-adaptor`.
 

## Day 1 \- Installation and Deployment Phase

### Project Installation and Configuration

  * Describe what project installation and configuration look like.
  
    Installation is managed declaratively via Helm values files (e.g., `-f values/kata-s390x.yaml` or `--set kata-as-coco-runtime.debug=true`). Trustee is configured via `TrusteeConfig` and `KbsConfig` CRDs which mount configurations and reference values to the KBS.

### Project Enablement and Rollback

  * How can this project be enabled or disabled in a live cluster? Please describe any downtime required of the control plane or nodes.
  
    Enablement happens transparently on a per-pod basis. Existing workloads are unaffected. To enable CoCo for a specific pod, the user simply adds `runtimeClassName: kata-qemu-[tee-type]` to the Pod spec. No control plane downtime is required.

  * Describe how enabling the project changes any default behavior of the cluster or running workloads.
  
    For CoCo pods, standard container image pulling behavior is preempted. The `nydus-snapshotter` routes the pull request to `image-rs` running *inside* the guest VM. This prevents the host `containerd` from keeping an unencrypted copy of the image, meaning images are not shared across pods on the same node.

  * Describe how the project tests enablement and disablement.
  
    Tested heavily via automated CI functional and integration tests across multiple TEE hardware environments (TDX, SNP, SGX). 
    
    ** On SGX and `enclave-cc` deprecation (https://github.com/confidential-containers/enclave-cc?tab=readme-ov-file#proposed-deprecation).  Suport for process-based confidential computing via `enclave-cc` (which provided SGX support) was [offically deprecated in v0.18.0](https://github.com/confidential-containers/confidential-containers/issues/315), and the enclave-cc project has been archived.  While the Trustee component can still technically validate SGX evidence for external workloads, CoCo's core runtime itself no longer executes workooads via SGX.

  * How does the project clean up any resources created, including CRDs?
  
    Running `helm uninstall` removes the deployed daemonsets and runtime classes. Trustee CRDs and controllers are removed via standard `kubectl delete` commands.

### Rollout, Upgrade and Rollback Planning

  * How does the project intend to provide and maintain compatibility with infrastructure and orchestration management tools like Kubernetes and with what frequency?
  
    The project maintains compatibility with Kubernetes 1.24+ and Containerd 1.7+, executing a regular 6-week release cadence.

  * Describe how the project handles rollback procedures.
  
    Because the runtime is deployed via Helm, rolling back involves using standard `helm rollback` commands to revert to a previous version of the `kata-deploy` configuration.

  * How can a rollout or rollback fail? Describe any impact to already running workloads.
  
    **TODO (Maintainers):** Describe any specific failure modes during upgrades/downgrades. For instance, do existing VMs keep running if the host-level `kata-shim` or `containerd` drops connection? Are there state-migration issues with Trustee CRDs during a rollback?

  * Describe any specific metrics that should inform a rollback.
  
    A spike in Pods stuck in `StartError` state (e.g., due to failing attestation, or `image-rs` failing to retrieve decryption keys from the KBS) indicates a severe configuration or network routing failure requiring rollback. See the [Troubleshooting Guide](https://github.com/confidential-containers/confidential-containers/blob/main/guides/troubleshooting.md).

  * Explain how upgrades and rollbacks were tested and how the upgrade-\>downgrade-\>upgrade path was tested.
  
    **TODO (Maintainers):** Detail the CI processes validating version-to-version upgrade and downgrade paths.

  * Explain how the project informs users of deprecations and removals of features and APIs.  
  
    Significant architectural shifts are documented clearly in repository READMEs (e.g., the [explicit deprecation notice](https://github.com/confidential-containers/operator/blob/main/README.md) of the `operator` in favor of the `charts` repository) and detailed in formal release notes on the [project blog](https://confidentialcontainers.org/blog/releases/).

  * Explain how the project permits utilization of alpha and beta capabilities as part of a rollout.
  
    Developers can utilize the `kata-qemu-coco-dev` runtime class to test attestation workflows, init-data, and resource policies locally on standard hardware before migrating to production TEEs. See the [coco-dev guide](https://github.com/confidential-containers/confidential-containers/blob/main/guides/coco-dev.md).