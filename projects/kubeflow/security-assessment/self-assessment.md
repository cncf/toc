# Kubeflow Security Self-Assessment

This document is a Kubeflow security self-assessment.

Authors:

- Kubeflow Community
- Alex (@droctothorpe)
- Andrey Velichkevich (@andreyvelich)
- Francisco Javier Arceo (@franciscojavierarceo)
- Humair Khan (@HumairAK)
- Julius von Kohout (@juliusvonkohout)
- Manabu McCloskey (@nabuskey)
- Matteo Mortari (@tarilabs)
- Mathew Wicks (@thesuperzapper)
- Vara Bonthu (@vara-bonthu)

# Table of Contents

- [Metadata](#metadata)
- [Overview](#overview)
  - [Background](#background)
  - [Actors](#actors)
  - [Actions](#actions)
  - [Goals](#goals)
  - [Non-Goals](#non-goals)
- [Self-assessment Use](#self-assessment-use)
- [Security functions and features](#security-functions-and-features)
- [Project Compliance](#project-compliance)
- [Secure Development Practices](#secure-development-practices)
  - [Development pipeline](#development-pipeline)
  - [Software Bill of Materials](#software-bill-of-materials)
  - [Communication Channels](#communication-channels)
  - [Ecosystem](#ecosystem)
- [Security issue resolution](#security-issue-resolution)
  - [Responsible disclosure practice](#responsible-disclosure-practice)
  - [Incident response](#incident-response)
- [Appendix](#appendix)

# Metadata

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Information</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Assessment Stage</td>
      <td>Complete</td>
    </tr>
    <tr>
      <td>Software</td>
      <td>
        <a href="https://github.com/kubeflow/spark-operator">Kubeflow Spark Operator</a>, 
        <a href="https://github.com/kubeflow/notebooks">Kubeflow Notebooks</a>, 
        <a href="https://github.com/kubeflow/trainer">Kubeflow Trainer</a>, 
        <a href="https://github.com/kubeflow/katib">Kubeflow Katib</a>, 
        <a href="https://github.com/kubeflow/model-registry">Kubeflow Model Registry</a>, 
        <a href="https://github.com/kubeflow/pipelines">Kubeflow Pipelines</a>
      </td>
    </tr>
    <tr>
      <td>Security Provider?</td>
      <td>No. Kubeflow projects are not security providers</td>
    </tr>
    <tr>
      <td>Languages</td>
      <td>Python, Go, TypeScript</td>
    </tr>
  </tbody>
</table>

# Overview

[Kubeflow](https://www.kubeflow.org/) consists of multiple open source projects that run on
Kubernetes and cover every stage of the [AI lifecycle](https://www.kubeflow.org/docs/started/architecture/#kubeflow-projects-in-the-ai-lifecycle).
Each project can be used independently or composed into an end-to-end AI platform.

Kubeflow subprojects have a control plane that orchestrates user-supplied AI workloads on Kubernetes,
so its security model builds on, and inherits, the security boundary of the underlying cluster.
Kubeflow aims to provide the following security properties:

- **Authenticated and authorized access**: access to Kubeflow APIs is mediated by Kubernetes RBAC,
  and access to Kubeflow web interfaces is authenticated and authorized at the Istio ingress gateway
  together with OAuth2 Proxy. Kubeflow does not introduce an authorization path that bypasses
  Kubernetes RBAC.
- **Tenant isolation**: Kubeflow's control-plane components (controllers, admission webhooks, and
  API servers, such as the Katib experiment and suggestion controllers) are shared, trusted services
  operated by the platform administrator, not by end users. End users — including multiple, mutually
  untrusted tenants — interact with Kubeflow only by creating namespaced custom resources (for
  example, `Experiment`, `TrainJob`, or `SparkApplication`) in the namespaces they are granted
  access to. Kubernetes namespaces, RBAC, and, where configured, network policies isolate each
  tenant, so one tenant cannot read or control another tenant's resources or workloads without an
  explicit grant. See the [Kubeflow Community Distribution architecture](https://github.com/kubeflow/community-distribution#architecture)
  for how this isolation is composed into an end-to-end platform.
- **Least-privilege control plane**: Kubeflow controllers run under scoped service accounts, and
  admission webhooks validate user-submitted resources before they are admitted.

These properties rely on the following trust assumptions and operator responsibilities. The
Kubeflow control-plane components are trusted services that the platform administrator (operator)
runs in dedicated system namespaces; end users neither run them nor have access to them. The
operator is responsible for:

- deploying, hardening, and protecting the control-plane components from compromise;
- configuring the Kubernetes primitives Kubeflow depends on (RBAC, namespaces, network policies,
  the Istio ingress gateway, and the identity provider);
- enforcing pod-level isolation for workloads — the reference [Kubeflow Community Distribution](https://github.com/kubeflow/community-distribution#architecture)
  applies Pod Security Standards `restricted` (falling back to `baseline` where a workload requires
  it) together with network policies by default, and standalone installations of individual Kubeflow
  projects are expected to apply the same `restricted`-by-default baseline;
- trusting the container images and user-supplied code they choose to run.

The control plane further assumes that the underlying Kubernetes cluster and its components (API
server, etcd, kubelet, container runtime, scheduler, and CNI) are correctly operated.

When these assumptions do not hold, the corresponding guarantees no longer apply. If RBAC or
namespace isolation is misconfigured, tenant isolation can break; if the ingress gateway or identity
provider is bypassed or misconfigured, the web-interface access controls no longer hold. Kubeflow
executes user-supplied code and does not sandbox it beyond the isolation Kubernetes provides, so
protecting the cluster and other tenants from a malicious or compromised workload relies on the
Kubernetes mechanisms (RBAC, namespace isolation, Pod Security Admission, network policies, and
resource quotas) that the operator configures. If a control-plane component itself is compromised,
the blast radius is bounded by its service account: an attacker can disrupt orchestration for the
resources that component manages — for example, a denial of service by failing to reconcile or
schedule jobs — and can read or write the namespaced resources that component is authorized to
access, but least-privilege, namespace-scoped RBAC means a single compromised controller does not by
itself grant access to unrelated tenants' data.

Please refer to [the official documentation](https://www.kubeflow.org/docs/) for more information.

## Background

Building AI platforms in cloud native environments poses many challenges given the uniqueness of
AI workloads. While Kubernetes provides a robust foundation for container orchestration, it was not
originally designed to support AI workloads. Some of the gaps including gang-scheduling, accelerator
management, monitoring, autoscaling, managing stateful workloads, often resulting in operational
complexity and inefficiencies.

Kubeflow projects address these challenges by creating Custom Resource Definitions (CRDs)
on top of Kubernetes to orchestrate AI workloads for Data Processing, AI Development, Model Training,
Model Optimization, ML Pipelines and MLOps governance. Additionally, Kubeflow offers a variety of interfaces including a
Python SDK and web UIs which allow users to focus on their models rather than the underlying
complexities.

Kubeflow is designed to be portable, extensible, and framework-agnostic. It offers pluggable
architecture that allows platform engineers integrate their internal systems with Kubeflow tools.

Every Kubeflow project is designed to address specific parts of the AI lifecycle. For example, Kubeflow
Trainer manages model training and Kubeflow Spark Operator manages data processing workloads.

## Actors

Kubeflow consists of **six** individual projects which allow organizations to build end-to-end
AI platforms. The following diagram shows the AI lifecycle and role of Kubeflow projects:

![ai-lifecycle](images/ai-lifecycle-kubeflow.drawio.svg)

The actors for each Kubeflow project are explained in the following sections:

### Kubeflow Spark Operator

![spark-operator](images/spark-operator.png)

- Spark Operator controller: A controller that watches for events of SparkApplication CRDs and acts on the watch events. It includes a submission runner that runs Spark submit for submissions received from the controller, and a Spark pod monitor that watches for Spark pods and sends pod status updates to the controller.

- Spark Mutating Webhook: a Mutating Admission Webhook that handles customizations for Spark driver
  and executor Pods based on the annotations on the Pods added by the controller.

Detailed information can be found here in the official
[Kubeflow Spark Operator docs](https://www.kubeflow.org/docs/components/spark-operator/overview/#architecture).

### Kubeflow Notebooks

- Notebook controller: controller that watches events of Notebook CRDs. Kubeflow Notebooks supports
  three type of IDEs: JupyterLab, RStudio, and VSCode. Supported images are explained in
  [this guide](https://www.kubeflow.org/docs/components/notebooks/container-images/#official-images).

- Culling controller: controller that watches for Notebooks activity and cull resources in case of
  idleness.

- Notebook web UI: web interface that allow users to manage their Notebooks. User can leverage this
  UI to create IDE and start model development. The web UI is integrated with Kubeflow Dashboard
  which gives users ability to see all created Notebooks.

### Kubeflow Katib

![katib](images/katib.png)

- Experiment controller: controller that watches events of Experiment CRDs which manage single
  hyperparameter tuning job. User can specify several parameters in Experiment such as objective
  to define metric that user wants to achieve, search space to define set of all hyperparameter
  values, and search algorithm to use for optimization job (e.g. bayesian optimization )

- Suggestion controller: controller that watches events of Suggestion CRDs which manage set of
  hyperparameter values that the hyperparameter tuning process has proposed. Suggestion is
  responsible to manage algorithm service.

- Trail controller: controller that watches events of Trial CRDs which manage one iteration of
  hyperparameter tuning process. A Trial corresponds to one worker job instance with a list of
  parameter assignments. The list of parameter assignments corresponds to a Suggestion.

- Katib webhooks: Validates and mutates CRD resources to ensure they conform to Katib standards
  and best practices. Katib also manages admission webhook to mutate metrics collector sidecar
  container into Trial workers.

- Katib SDK: Python SDK which allows ML engineers manage their Experiments without learning YAML
  configurations and Kubernetes APIs.

- Katib web UI: web interface that allow users to manage their Experiments. User can leverage this
  UI to create Experiment and compare results.

### Kubeflow Trainer

![trainer](images/trainer.svg)

- TrainingRuntime controller: controller that watches events of TrainingRuntime and
  ClusterTrainingRuntime CRDs. It is responsible to check whether runtime is unused and can be
  deleted or updated by users.

- TrainJob controller: controller that watches events of TrainJob CRDs. It is
  responsible to orchestrate distributed training job based on TrainJob and TrainingRuntime
  configurations. It creates additional resources like ConfigMap for MPI hostfile to ensure
  appropriate orchestration of jobs.

- Trainer webhook: Validates and mutate CRD resources to ensure they conform to Kubeflow Trainer
  standards and best practices.

- External Integrations: JobSet controller to manage orchestration of multiple jobs within single
  workload.

- Kubeflow Trainer SDK: Python SDK which allows ML engineers manage their TrainJobs without
  learning YAML configurations and Kubernetes APIs.

### Kubeflow Model Registry

![model-registry](images/model-registry.png)

- OpenAPI/REST Server: This component exposes a high-level REST API of the Model Registry.
  The REST API offers end-user capabilities focused on the domain model of Model Registry, like
  register a model, version a model, get a catalog of models, manage the deployment statutes of a model

- Python SDK: Python SDK to interact with Model Registry. This tool can be used by a user to execute
  operations such as retrieving the registered models, get model’s deployment status, model’s version etc.

Detailed information can be found here in the official
[Kubeflow Model Registry docs](https://www.kubeflow.org/docs/components/model-registry/reference/architecture/).

### Kubeflow Pipelines (KFP)

![pipelines](images/pipelines.png)

- KFP API Server: A gRPC and HTTP API server that accepts user requests to manage their ML pipelines.
  Users connect to the API server via REST or KFP SDK.

- Pipelines SQL DB: database that contains metadata about runs and pipelines. KFP API Server writes
  data to this DB.

- Persistent Agent: controller that watches for events of Workflow and ScheduledWorkflow CRDs.
  It reports status to the KFP API Server.

- Scheduler Workflow controller: controller that watches for events of ScheduledWorkflow CRDs.
  It creates runs for recurring runs. When a recurring run is created, the controller is in charge of
  creating the underlying Workflow CR at the scheduled times.

- ML Metadata gRPC server: gRPC server that manage data in MLMD SQL database.

- External Integrations:
  - Argo Workflow controller to manage orchestration of Workflow CRDs.
  - S3-complaint object store: object store that contains artifacts that launcher pods read and
    write to.
  - MLMD SQL DB: database that contains metadata around artifacts, executions, and lineage. All
    pipeline pods send and receive data from MLMD. KFP UI also communicates with this database.

- Envoy Gateway: Support KFP UI -> MLMD reads via HTTP/1.0 using gRPC-web protocol.

- Launcher Pod: Runs user's code that was written in KFP SDK. It uses the images specified
  by the user and it downloads Python packages from an external pip repository (e.g. PyPI).

## Actions

### Kubeflow Spark Operator

- Spark Application submission: Users create SparkApplication CRDs, specifying configuration for the
  Spark Job.

- Spark Operator controller automatically performs appropriate actions to orchestrate Spark jobs on
  Kubernetes

- Security and Access Control: Spark Operator leverages Kubernetes RBAC for Spark drivers and
  executors. This allows administrators to define who can create, modify, or delete SparkApplications
  and associated pods within the specific namespaces, enabling proper multi-tenant isolation.

### Kubeflow Notebooks

- Notebook creation: User access Kubeflow Notebook UI to create Notebook. Kubeflow Dashboard manages
  authentication and authorization for user via Istio gateway.

- Security and Access Control: Kubeflow Notebooks leverages Kubernetes RBAC to ensure that user can
  manage Notebooks.

### Kubeflow Katib

- Experiment creation: User creates Experiment CRDs, specifying configuration for the hyperparameter
  tuning job. User can create Experiment using: Python SDK, Kubeflow Dashboard, or `kubectl`. In
  case of Python SDK and `kubectl`, Katib relies on Kubernetes RBAC, in case of Kubeflow Dashboard
  user access it via Istio gateway.

- Security and Access Control: Katib leverages Kubernetes RBAC to ensure that user can
  manage Experiments.

### Kubeflow Trainer

- TrainJob creation: User creates TrainJob CRDs, specifying configuration for the training job.
  User should create TrainJob using Python SDK. Python SDK relies on Kubernetes RBAC to access
  Kubernetes cluster.

- TrainingRuntime creation: User creates TrainingRuntime or ClusterTrainingRuntime CRDs, specifying
  template configuration for the training job. User should create CRDs via `kubectl`.

- Security and Access Control: Kubeflow Trainer leverages Kubernetes RBAC to ensure that user can
  manage TrainJobs and TrainingRuntimes.

### Kubeflow Model Registry

- Save models: User utilizes Python client to interact with Model Registry. Also, user can use
  Kubeflow Dashboard to manage models. In that case, user access it via Istio gateway

### Kubeflow Pipelines

- Create pipelines: User utilized Python client to manage ML pipelines. Also,
  can use Kubeflow Dashboard to access KFP UI. In that case, user access it via Istio gateway

- Security and Access Control: Kubeflow Pipelines leverages Kubernetes RBAC to ensure that user can
  manage Pipelines.

## Goals

- **Foundation for AI Platforms on Kubernetes**: Kubeflow projects provide a standard, cloud native
  solution to build end-to-end AI platforms on Kubernetes. The Kubeflow tools are scalable and
  composable which gives organizations flexibility to adopt them.

- **Simplify Complexity**: Abstract the Kubernetes complexity for ML users to manage and scale AI
  workloads. Define Python-first experience for users, so they can focus on their models rather
  than the underlying infrastructure.

- **Large Scale Model Development**: Provide capability to develop, train, optimize, and deploy
  LLMs at scale.

- **Out of the Box ML Framework Support**: Support a wide range of ML frameworks like PyTorch,
  JAX, HuggingFace, DeepSpeed, XGboost, and others to run AI workloads.

- **Advanced Orchestration Patterns**: Support advanced orchestration patterns such as gang-scheduling,
  High-Performance Computing jobs, and GPU cost optimization to run AI workloads.

- **End-to-End GenAI on Kubernetes**: Provides simple APIs and interfaces for users to perform
  end-to-end GenAI use-cases including LLM fine-tuning, RAG, LLM optimizations, LLM inference,
  AI agent orchestration.

- **Interactive AI development**: Provide interactive development environment for ML users to
  develop and deploy AI models.

## Non-Goals

- Kubeflow doesn't re-invent ML frameworks and packages (e.g. PyTorch, JAX).
- Kubeflow is not a replacement for GitOps systems like ArgoCD.
- Kubeflow doesn't enforce a deployment method or distribution for Kubeflow projects.
- Kubeflow does not manage the security of external dependencies or third-party plugins beyond
  standard Kubernetes best practices.

# Self-assessment Use

This self-assessment is created by the Kubeflow team to perform an internal analysis of the
project’s security. It is not intended to provide a security audit of Kubeflow, or function as an
independent assessment or attestation of Kubeflow’s security health.

This document serves to provide Kubeflow users with an initial understanding of Kubeflow’s security,
where to find existing security documentation, Kubeflow plans for security, and general overview of
Kubeflow security practices, both for development of Kubeflow projects as well as security of
Kubeflow projects.

This document provides the CNCF TAG-Security with an initial understanding of Kubeflow to assist in
a joint-assessment, necessary for projects under incubation. Taken together, this document and the
joint-assessment serve as a cornerstone for if and when Kubeflow seeks graduation and is preparing
for a security audit.

# Security functions and features

| Component                          | Applicability | Description of Importance                                                                               |
| ---------------------------------- | ------------- | ------------------------------------------------------------------------------------------------------- |
| Kubernetes RBAC & Network Policies | Critical      | Leverages Kubernetes RBAC and network policies to restrict access to resources and isolate workloads.   |
| Kubeflow Web UI                    | Critical      | Leverages Istio Gateway and OAuth2 Proxy to authenticate and authorize users for Kubeflow UIs           |
| Secure Defaults                    | Relevant      | Default configurations are designed to minimize exposure and follow Kubernetes security best practices. |
| Monitoring & Logging               | Relevant      | Integrates with Prometheus, Grafana, and logging systems for observability and incident response.       |

# Project Compliance

Kubeflow projects do not currently claim compliance with specific security standards
(e.g., PCI-DSS, ISO, GDPR). The Kubeflow projects follow open source and Kubernetes security best
practices and is following the OpenSSF Best Practices. Kubeflow projects achieve the OpenSSF badges:

- Kubeflow Spark Operator: https://www.bestpractices.dev/en/projects/10524
- Kubeflow Notebooks: https://www.bestpractices.dev/en/projects/9942
- Kubeflow Katib: https://www.bestpractices.dev/projects/9941
- Kubeflow Trainer: https://www.bestpractices.dev/projects/10435
- Kubeflow Model Registry: https://www.bestpractices.dev/en/projects/9937
- Kubeflow Pipelines: https://www.bestpractices.dev/en/projects/9938

# Secure Development Practices

## Development pipeline

- Committers are required to agree to the Developer Certificate of Origin (DCO) for each and
  every commit by simply stating you have a legal right to make the contribution.
- At least one reviewer is required for a pull request to be approved.
- Automated tests for unit, integration, and end-to-end testing.
- Automated code quality checks and linting.
- Publicly documented contribution and code review guidelines
  - https://www.kubeflow.org/docs/about/contributing/
  - https://github.com/kubeflow/spark-operator/blob/master/CONTRIBUTING.md
  - https://github.com/kubeflow/notebooks/blob/main/CONTRIBUTING.md
  - https://github.com/kubeflow/trainer/blob/master/CONTRIBUTING.md
  - https://github.com/kubeflow/katib/blob/master/CONTRIBUTING.md
  - https://github.com/kubeflow/pipelines/blob/master/CONTRIBUTING.md
  - https://github.com/kubeflow/model-registry/blob/main/CONTRIBUTING.md

- Tide is enabled to automatically merge the PRs after approvers add `/lgtm` and `/approve` labels.

- Release process for each Kubeflow project is described in their GitHub repositories:
  - https://github.com/kubeflow/spark-operator/blob/master/docs/release.md
  - https://github.com/kubeflow/trainer/blob/master/RELEASE.md
  - https://github.com/kubeflow/katib/blob/master/docs/release/README.md
  - https://github.com/kubeflow/pipelines/blob/master/RELEASE.md
  - https://github.com/kubeflow/model-registry/blob/main/RELEASE.md

## Software Bill of Materials

The package versions for each Kubeflow project can be found in the appropriate repositories:

- Kubeflow Spark Operator
  - Go modules: https://github.com/kubeflow/spark-operator/blob/master/go.mod
  - FOSSA report: https://app.fossa.com/reports/b1c3b2c0-a1bf-4572-b3b2-8455729c8cf9
- Kubeflow Notebooks
  - Go modules: https://github.com/kubeflow/notebooks/blob/notebooks-v1/components/notebook-controller/go.mod
  - FOSSA report: https://app.fossa.com/reports/8720a333-0c53-4a2d-a3c9-92a2435d4d7c
- Kubeflow Trainer
  - Go modules: https://github.com/kubeflow/trainer/blob/master/go.mod
  - Python packages: https://github.com/kubeflow/sdk/blob/main/pyproject.toml
  - FOSSA report: https://app.fossa.com/reports/bb8e2d41-254a-4af3-9044-e7f484c34dd1
- Kubeflow Katib
  - Go modules: https://github.com/kubeflow/katib/blob/master/go.mod
  - Python packages: https://github.com/kubeflow/katib/blob/master/sdk/python/v1beta1/setup.py
  - FOSSA report: https://app.fossa.com/reports/9762748b-b1fb-43ed-84d1-bee58e61e619
- Kubeflow Model Registry
  - Go modules: https://github.com/kubeflow/model-registry/blob/main/go.mod
  - Python packages: https://github.com/kubeflow/model-registry/blob/main/clients/python/pyproject.toml
  - FOSSA report: https://app.fossa.com/reports/0559a384-a6ce-4a37-b286-fcc47c962790
- Kubeflow Pipelines
  - Go modules: https://github.com/kubeflow/pipelines/blob/master/go.mod
  - Python packages: https://github.com/kubeflow/pipelines/blob/master/sdk/python/requirements.txt
  - FOSSA report: https://app.fossa.com/reports/295e1511-cf5e-405c-abc3-eddac2ef03ae

All Kubeflow container images can be found here: https://github.com/orgs/kubeflow/packages

The SBOMs for any released Kubeflow container image can be accessed using the following command:

```
docker sbom ghcr.io/kubeflow/trainer/trainer-controller-manager:v2.0.0-rc.0
```

## Communication Channels

### Internal

- Kubeflow projects have a dedicated CNCF Slack channel `kubeflow-contributors` used for developer
  communication.
- Kubeflow KSC has a private mailing list for maintainers and user requests: `ksc@kubeflow.org`

### Inbound

- Kubeflow Projects have dedicated CNCF Slack channels to communicate with users and maintainers:
  - `#kubeflow-spark-operator`
  - `#kubeflow-notebooks`
  - `#kubeflow-trainer`
  - `#kubeflow-katib`
  - `#kubeflow-hub`
  - `#kubeflow-pipelines`

- Kubeflow Working Groups host weekly community meetings: https://www.kubeflow.org/docs/about/community/#list-of-available-meetings

### Outbound

- Kubeflow users this Slack channel for announcements: `#kubeflow-announcements`

- Kubeflow has a dedicated mailing: `kubeflow-discuss@googlegroups.com`

- Kubeflow also uses various social media:
  - LinkedIn: https://www.linkedin.com/company/kubeflow/
  - X: https://x.com/kubeflow
  - BlueSky: https://bsky.app/profile/kubefloworg.bsky.social
  - YouTube Channel: https://www.youtube.com/@Kubeflow and https://www.youtube.com/@KubeflowCommunity

## Ecosystem

Kubeflow projects extensively use cloud native ecosystem, including Kubernetes, Argo Workflow,
Istio, Helm, Knative, KServe, JobSet, Kueue, and other projects.

Many CNCF projects have integrations with Kubeflow tools like KubeEdge and Volcano have integrations
with Kubeflow Trainer.

# Security Issue Resolution

Kubeflow projects maintain security policy in their corresponding `SECURITY.md` files:

- https://github.com/kubeflow/spark-operator/blob/master/SECURITY.md
- https://github.com/kubeflow/notebooks/blob/master/SECURITY.md
- https://github.com/kubeflow/trainer/blob/master/SECURITY.md
- https://github.com/kubeflow/katib/blob/master/SECURITY.md
- https://github.com/kubeflow/model-registry/blob/main/SECURITY.md
- https://github.com/kubeflow/pipelines/blob/master/SECURITY.md

## Responsible Disclosure Practice

Vulnerability reports are accepted via GitHub's private vulnerability reporting tool or via a
private mailing list `ksc@kubeflow.org`. Maintainers collaborate with reporters to triage and
resolve issues, and advisories are published as needed.

## Incident Response

Upon receiving a vulnerability report, the Kubeflow maintainers triage the issue, determine impact,
and work to release a patch for supported versions. Users are notified via release notes,
documentation, and community channels. Only a limited number of maintainers have access to
the reports to avoid excessive disclosure of vulnerabilities.

# Appendix

- _Security Policies_: each Kubeflow project documents its security policy in a `SECURITY.md` file:
  - https://github.com/kubeflow/spark-operator/blob/master/SECURITY.md
  - https://github.com/kubeflow/notebooks/blob/master/SECURITY.md
  - https://github.com/kubeflow/trainer/blob/master/SECURITY.md
  - https://github.com/kubeflow/katib/blob/master/SECURITY.md
  - https://github.com/kubeflow/model-registry/blob/main/SECURITY.md
  - https://github.com/kubeflow/pipelines/blob/master/SECURITY.md
- _Known Issues Over Time_: Issues are tracked in GitHub Issues, and vulnerabilities are tracked
  and disclosed through GitHub Security Advisories for each Kubeflow project. Maintainers triage
  reported vulnerabilities, release patches for supported versions, and publish advisories as
  needed (see [Incident Response](#incident-response)). The current set of published and open
  advisories for each project can be reviewed here:
  - https://github.com/kubeflow/spark-operator/security/advisories
  - https://github.com/kubeflow/notebooks/security/advisories
  - https://github.com/kubeflow/trainer/security/advisories
  - https://github.com/kubeflow/katib/security/advisories
  - https://github.com/kubeflow/model-registry/security/advisories
  - https://github.com/kubeflow/pipelines/security/advisories
- _OpenSSF Best Practices_: Kubeflow projects achieved the passing level criteria and some projects
  achieve silver and gold status. Community is currently working to achieve silver status for
  each Kubeflow project.
- _Case Studies_: Kubeflow is used by a lot of organizations including AstraZeneca, AWS, Bloomberg,
  CERN, DHL, Google, IBM, NVIDIA, Nutanix, Red Hat, Toyota, and many others. The extensive list of
  adopters for each project can be found here: https://github.com/kubeflow/community/blob/master/ADOPTERS.md
- _Related Projects / Vendors_: Related projects including KServe, Volcano, MLFlow, Airflow. Vendors including
  SageMaker, VertexAI, AzureML, OpenShift AI.
- _Competitive projects_: MLFlow, Airflow, Ray.
  Kubeflow projects differentiate by providing Kubernetes-native solution, extensible and portable.
  That gives organizations more flexibility to integrate Kubeflow tools within AI platforms.
