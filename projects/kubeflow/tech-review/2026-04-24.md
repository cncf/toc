# General Technical Review - Kubeflow / Graduation

- **Project:** Kubeflow

- **Project Version:** Every Kubeflow sub-project has its own version.

- **Website:** https://www.kubeflow.org/

- **Date Updated**: 2025-09-05

- **Template Version**: v1.0

- **Description**:

Kubeflow is the foundation of tools for AI Platforms on Kubernetes.

AI platform teams can build on top of Kubeflow by using each project independently or deploying the
entire AI reference platform to meet their specific needs. The Kubeflow AI reference platform is
composable, modular, portable, and scalable, backed by an ecosystem of Kubernetes-native projects
that cover every stage of the
[AI lifecycle.](https://www.kubeflow.org/docs/started/architecture/#kubeflow-projects-in-the-ai-lifecycle)

Whether you’re an AI practitioner, a platform administrator, or a team of developers,
Kubeflow offers modular, scalable, and extensible tools to support your AI use cases.

## What are Kubeflow Projects

Kubeflow is composed of multiple open source projects that address different aspects of the AI
lifecycle. These projects are designed to be usable both independently and as part of the Kubeflow
AI reference platform. This provides flexibility for users who may not need the full end-to-end
AI platform capabilities but want to leverage specific functionalities, such as model training or
model serving.

### Kubeflow Projects in scope for CNCF Graduation

- Kubeflow Spark Operator
- Kubeflow Notebooks
- Kubeflow Trainer
- Kubeflow Katib
- Kubeflow Model Registry
- Kubeflow Pipelines

## What is the Kubeflow AI Reference Platform

The Kubeflow AI reference platform refers to the full suite of Kubeflow projects bundled together
with additional integration and management tools. Kubeflow AI reference platform deploys the
comprehensive toolkit for the entire AI lifecycle. The Kubeflow AI reference platform can be
installed via [Packaged Distributions](https://www.kubeflow.org/docs/started/installing-kubeflow/#packaged-distributions)
or [Kubeflow Manifests](https://www.kubeflow.org/docs/started/installing-kubeflow/#kubeflow-manifests).

## Day 0 - Planning Phase

### Scope

#### Describe the roadmap process

Kubeflow projects are managed by community members that are part of working groups.
Each working group defines and agrees on the features for each release. The release cadence for
each working group varies according to community agreement among the working groups.

Then, all Kubeflow projects have individual roadmap files in the Git repository defining each
release and available to the public. This ensures we have a standard structure for each proposed
feature, auditing, versioning, and transparency, since it is recorded along with the history in the Git repo.

For more information, check ROADMAP for each Kubeflow Project:

- [Kubeflow Spark Operator](https://github.com/kubeflow/spark-operator/blob/master/ROADMAP.md)
- [Kubeflow Trainer](https://github.com/kubeflow/trainer/blob/master/ROADMAP.md)
- [Kubeflow Katib](https://github.com/kubeflow/katib/blob/master/ROADMAP.md)
- [Kubeflow Hub](https://github.com/kubeflow/hub/blob/main/ROADMAP.md)
- [Kubeflow Pipelines](https://github.com/kubeflow/pipelines/blob/master/ROADMAP.md)

Community-wide changes are proposed as [Kubeflow Enhancement proposals (KEPs)](https://github.com/kubeflow/community/tree/master/proposals)
in the `kubeflow/community` repository or in the [Kubeflow sub-projects KEPs](https://github.com/kubeflow/trainer/tree/master/docs/proposals).

#### Describe the target persona or user(s) for the project

1. Users: Data Scientists, ML Engineer, AI Practitioners, Data Engineers, AI Practitioners.
2. Operators: MLOps Engineers, AIOps engineers, Platform Engineers, AI Platform Engineers.
3. Vendors: Vendors and projects building Kubernetes based AI Platform products.

#### Explain the primary use case for the project. What additional use cases are supported by the project

The goal of Kubeflow is to run Cloud Native AI workloads for every stage in AI lifecycle. By
using Kubeflow projects users can develop and deploy AI applications.

The primary use-cases include:

- Large-scale data processing and feature engineering.
- Distributed pre-training of foundation models.
- Post-training and fine-tuning of LLMs.
- Hyperparameter optimization and model tuning.
- LLM inference and multi-host serving.

Additional Use Cases:

- End-to-End GenAI pipeline building.
- Interactive AI development.
- Multiple users / projects with hard multi-tenancy on the same cluster.

#### Explain which use cases have been identified as unsupported by the project

As Kubeflow is composed of multiple projects, each working group makes its own determinations as t
what will be excluded from them. However we have an overarching theme and governance structure
(Steering Committee) that has identified the following areas as not being a priority for all projects:

- The projects are deployed in any Kubernetes (each release will specify tested versions),
  regardless of the underlying infrastructure, independently through Kubernetes manifests leveraging
  Kustomize and/or Helm Charts. However, the project doesn’t provide an implementation to be deployed
  on infrastructure besides Kubernetes. - We do not officially enforce a deployment method or distribution.

- Kubeflow doesn’t provide a GitOps implementation, however Kubeflow manifests can be integrated
  into a GitOps solution. For example, Platform Engineers can create an ArgoCD Application (CRD)
  to install and configure Kubeflow projects. by providing Kubeflow individual project manifests,
  for example Pipelines. The GitOps application will read from the Kubeflow Pipelines manifest and
  Argo CD will deploy the configurations in the target cluster.

#### Describe the intended types of organizations who would benefit from adopting this project

Kubeflow is intended to be used by any organization which needs to run AI workloads on Kubernetes,
in any of the AI lifecycle stages an organization might choose to use just one or more projects,
such as Pipelines or Training. Organizations can also use all the projects from Kubeflow to increase
the user experience to build AI workloads. Additionally, organizations can develop their own
customizations on top of the Kubeflow platform or choose to build distributions to help other
organizations adopt a customized platform based on Kubeflow projects.

As Kubeflow maintains flexibility, organizations can choose their own path according to their needs.

We encourage adopters to be part of the adopters list in [GitHub](https://github.com/kubeflow/community/blob/master/ADOPTERS.md).

Examples of organizations that use Kubeflow:

- AWS
- Red Hat
- Capital One
- CERN
- Google
- Alibaba Cloud
- Bloomberg
- IBM
- Cisco
- Huawei
- Microsoft
- Tencent
- DHL Data & Analytics
- Telia
- Roblox
- Toyota
- PepsiCo
- Volvo
- and others…

#### Describe any completed end user research and link to any reports.

We regularly undertake user research about Kubeflow and its users. Research could be done by
working groups, during events, or by conducting at least one study annually. We ensure these surveys
are visible to the public and shared with the community.

Here are some results from previous years.

- [Kubeflow 2025 Survey](https://docs.google.com/forms/d/11cSe5vmGLrGekJISHBMfjVh_97WFGuhcvGnd0l5aNLg/edit#responses)
- [2025:UX designers supporting Model Registry conducted a series of user sessions to understand preferred interaction patterns (link)](https://docs.google.com/forms/d/11cSe5vmGLrGekJISHBMfjVh_97WFGuhcvGnd0l5aNLg/edit#responses)
- [Kubeflow Survey 2024](https://github.com/kubeflow/community/issues/708)
- [Kubeflow Survey 2023](https://blog.kubeflow.org/kubeflow-user-survey-2023/)
- [Kubeflow Survey 2022](https://blog.kubeflow.org/kubeflow-user-survey-2022/)

### Usability

#### How should your target personas interact with your project

AI Practitioner - Kubeflow SDK, Kubeflow UIs

Platform Admins - Operator guides, installing and configuring any Kubeflow projects via Helm charts
or Kustomize manifests predefined and available in the Kubeflow documentation from the command line.

- [Getting Started guide](https://www.kubeflow.org/docs/started/introduction/).
- [How to install Kubeflow projects](https://www.kubeflow.org/docs/started/installing-kubeflow/).
- [How to contribute to Kubeflow](https://www.kubeflow.org/docs/about/contributing/).

#### Describe the user experience (UX) and user interface (UI) of the project

Kubeflow user experience in each project is a collection of projects, the user experience for the
projects are each with their own [interfaces, APIs and SDKs](https://www.kubeflow.org/docs/started/architecture/#kubeflow-interfaces).

##### Describing User Experience through SDK

We are working on a unified [Kubeflow SDK](https://github.com/kubeflow/sdk) that gives
AI practitioners Python-native experience to interact with Kubeflow APIs.

Through Kubeflow SDK, users will be able to interact with the different projects, increasing user experience and reducing complexity.

<div style="text-align: center;">
  <img
    src="https://raw.githubusercontent.com/kubeflow/sdk/main/docs/images/persona_diagram.svg"
    width="600"
    title="Kubeflow SDK Personas"
    alt="Kubeflow SDK Personas"
  />
</div>
 
Example to interact with Kubeflow Trainer using Kubeflow SDK:

```py
from kubeflow.trainer import TrainerClient, BuiltinTrainer, TorchTuneConfig

# Fine-Tune Llama 3.2
job_id = TrainerClient().train(
    runtime=TrainerClient().get_runtime(name="torchtune-llama3.2-1b"),
    trainer=BuiltinTrainer(
        config=TorchTuneConfig(
            resources_per_node={
                "memory": "200G",
                "gpu": 1,
            },
        )
    ),
)

# Wait for TrainJob to complete
TrainerClient().wait_for_job_status(job_id)

# Print TrainJob logs
print("\n".join(TrainerClient().get_job_logs(name=job_id)))
```

Example to interact with Kubeflow Katib using Kubeflow SDK

```py
from kubeflow.optimizer import OptimizerClient, search
from kubeflow.trainer import TrainerClient, BuiltinTrainer, TorchTuneConfig

OptimizerClient().optimize(
    objective="loss",
    mode="min",
    num_trials=5,
    trainer=BuiltinTrainer(
        config=TorchTuneConfig(
            resources_per_node={
                "memory": "200G",
                "gpu": 1,
            },
            lr=search("0.1", "0.2", "lognormal"),
            num_epochs=search("4", "8", "uniform"),
        )
    ),
    runtime=TrainerClient().get_runtime(name="torchtune-llama3.2-1b"),
)
```

Find more details and examples about [Kubeflow SDK](https://github.com/kubeflow/sdk).

##### Describing User Experience through APIs

- Kubeflow Pipelines

REST API is available under the `/pipeline/` HTTP path. For example, if you host Kubeflow at
https://kubeflow.example.com, the API will be available at https://kubeflow.example.com/pipeline/.

The API is documented using swagger and examples about its usage
[can be found here](https://www.kubeflow.org/docs/components/pipelines/reference/api/kubeflow-pipeline-api-spec/)

- Kubeflow Model Registry

Model Registry REST API is available under the `/api/model_registry/` HTTP path. More information
[can be found here](https://www.kubeflow.org/docs/components/hub/reference/rest-api/)

The Model Catalog Service provides a read-only discovery service for ML models across multiple
catalog sources. It acts as a federated metadata aggregation layer, allowing users to search and
discover models from various external catalogs through a unified REST API.

##### Describing User Interfaces through web UI

[Kubeflow Central Dashboard](https://www.kubeflow.org/docs/components/central-dash/overview/)
acts as a hub for the AI platform and tools by exposing the UIs of components running in the cluster.

<div style="text-align: center;">
  <img
    src="https://github.com/kubeflow/website/blob/498691a0eae4bef5ae5ae7935dd367a547274f85/content/en/docs/images/dashboard/homepage.png?raw=true"
    width="600"
    title="Kubeflow Dashboard"
    alt="Kubeflow Dashboard"
  />
</div>

End users can access the Central Dashboard installed version on their organization according to
the user and access control setup previously by Platform Admin. Users can access Kubeflow projects
such as Kubeflow Pipelines (KFP) to manage experiments, pipeline definitions, runs, and recurrent
runs. KFP Artifacts to track artifacts produced by pipelines stored in MLMD. KFP Executions to
track executions of pipeline components stored in MLMD. Kubeflow Katib experiments to manage
AutoML experiments. Kubeflow Notebooks to manage Kubeflow Notebooks. Kubeflow TensorBoards to
manage TensorBoard instances. Kubeflow Volumes to manage Kubernetes PVC Volumes. Contributors page
to manage contributors of profiles (namespaces) that you own.

More information can be found [in the Kubeflow Dashboard docs](https://www.kubeflow.org/docs/components/central-dash/overview/).

#### Describe how this project integrates with other projects in a production environment

All Kubeflow projects are Kubernetes native and so fit into the wider ecosystem of Kubernetes
based tools. Kubeflow projects extensively use tools from the cloud native ecosystem,
including, Argo Workflow, Istio, Helm, Knative, KServe, JobSet, Kueue, and other projects.

Specific components integrate with popular AI frameworks where applicable (e.g. Kubeflow Trainer
integrates with PyTorch and other model frameworks)

The following diagram gives an overview of the
[Kubeflow Ecosystem](https://www.kubeflow.org/docs/started/architecture/#kubeflow-ecosystem)
and how it relates to the wider Kubernetes AI landscapes.

<div style="text-align: center;">
  <img
    src="https://github.com/kubeflow/website/blob/498691a0eae4bef5ae5ae7935dd367a547274f85/content/en/docs/started/images/kubeflow-architecture.drawio.svg?raw=true"
    width="600"
    title="Kubeflow Ecosystem"
    alt="Kubeflow Ecosystem"
  />
</div>

Specifically for a production environment, Kubeflow can integrate and leverage the following
resources and projects:

- Cloud Providers and underlying infrastructure
  - Kubeflow can leverage the underlying infrastructure provided through Kubernetes, such as access
    hardware accelerators from Intel, Nvidia, and AMD through Kubernetes Operators.
  - Organizations can run Kubeflow projects on any Kubernetes platform and use cloud provider
    services, such as scalability across different nodes, taking advantage of projects such as Open Cluster Management.
- Configuration as code
  - GitOps approach can be used to set up Kubeflow setup and configuration, for example, using
    Argo CD to manage configuration as code in multiple clusters.
- Security
  - Security is becoming a key aspect, Kubeflow projects implement security best practices in
    Kubernetes including Pod Security Standards, Network policies, RBAC. Additionally, they can
    leverage projects such as the External Secrets Operator.
  - Data Scientists can also use Sigstore projects to sign and validate their models to avoid type
    of attacks such as tampering with Models to ensure the integrity and authenticity of models
    persists during the Model Development Lifecycle.
- Simplify User experience
  - Organizations can simplify Data scientists' experience by scaling Platform Engineers by
    providing software templates through Backstage to access environments quickly, leveraging
    Kubeflow projects such as Notebooks.

### Design

#### Explain the design principles and best practices the project is following

- Loosely Coupled and Distributable Services
  - Built as modular, independent microservices that can scale and evolve.
  - Services communicate through well defined APIs, enabling flexible orchestration and integration.
- Kubernetes native with Declarative APIs and Automation
  - Uses Kubernetes CRD for each project.
  - Integrate natively with Kubernetes core APIs like Jobs, Pods, Deployment.
  - Supports automated deployment, orchestration, and CI/CD workflows.
  - Enables GitOps management for reproducible and auditable operations.
- Portability and Platform-Agnosticism
  - Designed to run on any Kubernetes cluster, including onprem, edge, public cloud or hybrid environments.
  - Avoids vendor lock-in while ensuring consistent behavior across different infrastructures.
- Observability
  - Provides metrics, logs, and traces to monitor system performance and workflow execution.
  - Enables debugging, performance analysis, and monitoring of workloads and processes.
- Resilience and Availability
  - Leverages platform features such as health checks, automatic restarts, and scaling to maintain high availability.
  - Designed to tolerate failures and minimizing disruption to workloads.
- Security and Multi-Tenancy
  - Implements access control, isolation, and resource quotas to separate workloads and users.
  - Supports secure communication and policy enforcement to maintain a safe multi-tenant environment.
- Extensibility and Interoperability
  - Provides APIs and runtime contracts to allow integration of new tools, frameworks, or workflows.
  - Supports extension and customization without breaking existing functionality.

Kubeflow overall has strong [contributing guidelines](https://www.kubeflow.org/docs/about/contributing/)
that inform our development and design. Each project also provides new contributors specific
guidance on their GitHub repositories:

- [Kubeflow Spark Operator](https://github.com/kubeflow/spark-operator/blob/master/CONTRIBUTING.md)
- [Kubeflow Notebooks](https://github.com/kubeflow/notebooks/blob/main/CONTRIBUTING.md)
- [Kubeflow Trainer](https://github.com/kubeflow/trainer/blob/master/CONTRIBUTING.md)
- [Kubeflow Katib](https://github.com/kubeflow/katib/blob/master/CONTRIBUTING.md)
- [Kubeflow Model Registry](https://github.com/kubeflow/model-registry/blob/main/CONTRIBUTING.md)
- [Kubeflow Pipelines](https://github.com/kubeflow/pipelines/blob/master/CONTRIBUTING.md)

For the new features every Kubeflow project follows
[the KEP guidelines](https://github.com/kubeflow/community/tree/master/proposals#kep-format-to-propose-and-document-enhancements-to-kubeflow)

#### Outline or link to the project’s architecture requirements

We do follow the same practices as Kubernetes for APIs, with alpha, beta, and stable status of APIs.
Details [are described here](https://www.kubeflow.org/docs/started/support/#component-status).

#### Define any specific service dependencies the project relies on in the cluster

The following projects are required to install Kubeflow projects:

- Istio, Knative, Cert-manager
- Detailed information [can be found in the Kubeflow manifests](https://github.com/kubeflow/manifests#kubeflow-components-versions).

Specific projects have other dependencies:

- Kubeflow Pipelines: Argo Workflows>=v3.1
- Kubeflow Trainer: JobSet>=v0.8.0
- Kubeflow Katib: MySQL>=v8.0
- Kubeflow Model Registry: MySQL>=v8.0

#### Describe how the project implements Identity and Access Management

Kubeflow projects use a pluggable system of IAM that connects back to Kubernetes IAM in most cases.
As a reference implementation, the optional Kubeflow Manifests/Dashboard components use the
following IAM systems:

- [OAuth2 Proxy](https://github.com/kubeflow/manifests#oauth2-proxy)
- [Kubeflow manifest with Dex](https://github.com/kubeflow/manifests#dex)

General security Talks about Kubeflow including an architectural introduction:

- [KubeCon / Kubeflow Summit 2024 Lightning Talk](https://youtu.be/f7C3v0gJRVI).
- [KubeCon 2023 Hardening Kubeflow Security for Enterprise Environments](https://youtu.be/wQaOa4Kjxs0).
- [Kubeflow Summit 2023 Security Working Group Update](https://youtu.be/XKHVt2yxQFo).
- [Blog Hardening Kubeflow Security for Enterprise Environments](https://blogs.vmware.com/opensource/2023/06/20/hardening-kubeflow-security-for-enterprise-environments-2/).

#### Describe how the project has addressed sovereignty

Kubeflow’s projects can be self-hosted on any Kubernetes cluster, including air-gapped environments,
which is critical for those organizations using disconnected environments. Additionally,
Kubeflow projects support multi-tenancy, which allows organizations to isolate workloads in a
Kubernetes cluster, which can be used with Network Policies to restrict isolation from different
components by ports or application name. Kubeflow projects support custom service accounts,
allowing organizations to control how Kubeflow projects interact with the cluster by leveraging RBAC.
Further, Platform Administrators can decide which users in Kubernetes will have access to the
Kubeflow projects by creating different users and RBAC policies. Kubeflow Notebooks also supports
profiles, which allows simplifying the permission management.

In terms of data management, data sovereignty is managed by those who deploy or package the projects.

#### Describe any compliance requirements addressed by the project

Kubeflow projects are extensible which allows users to fit their internal compliance requirements.
As a result, specific common compliance frameworks (SOC-2, GDPR, etc.) are the responsibility of
end users and vendors. However, we aim to provide a strong foundation through reference architectures
similar things from which to build on.

#### Describe the project’s High Availability requirements

The end users can adjust the replicas. Kubeflow project controllers support leader election,
[for example Kubeflow Trainer](https://github.com/kubeflow/trainer/blob/master/cmd/trainer-controller-manager/main.go#L80).

#### Describe the project’s resource requirements, including CPU, Network and Memory

The following table shows the resource requirements for each Kubeflow project, calculated as the
maximum of actual usage and configured requests for CPU/memory, plus storage requirements from PVCs:

The maximum looks hefty so rather consider in general the maximum (manifest requests, average usage).

| Component               | CPU (cores) | Memory (Mi) | Storage (GB) |
| ----------------------- | ----------- | ----------- | ------------ |
| Cert Manager            | 3m          | 130Mi       | 0GB          |
| Dex + OAuth2-Proxy      | 3m          | 28Mi        | 0GB          |
| Istio                   | 2300m       | 3502Mi      | 0GB          |
| Kubeflow KServe         | 600m        | 1200Mi      | 0GB          |
| Kubeflow Katib          | 9m          | 471Mi       | 13GB         |
| Kubeflow Core           | 35m         | 841Mi       | 0GB          |
| Metadata                | 78m         | 687Mi       | 30GB         |
| Kubeflow Model Registry | 510m        | 2112Mi      | 0GB          |
| Other                   | 20m         | 354Mi       | 6GB          |
| Kubeflow Pipelines      | 970m        | 3552Mi      | 90GB         |
| Kubeflow Spark Operator | 4m          | 41Mi        | 0GB          |
| Kubeflow Trainer        | 3m          | 25Mi        | 0GB          |
| **Total**               | **4535m**   | **12943Mi** | **139GB**    |

#### Describe the project’s storage requirements, including its use of ephemeral and/or persistent storage

Each project can configure storage in different ways. This is also true of the manifests which
configure storage for a collection of projects.
This [can be seen here](https://github.com/kubeflow/manifests/pull/3091).

#### Please outline the project’s API Design

Various Kubeflow projects offer APIs and Python SDKs. See the following sets of reference documentation:

- [Pipelines reference docs](https://www.kubeflow.org/docs/components/pipelines/reference/)
  for the Kubeflow Pipelines API and SDK, including the Kubeflow Pipelines domain-specific language (DSL).
- [Kubeflow SDK](https://github.com/kubeflow/sdk)
  to interact with Kubeflow APIs with Python-native interface.

  See also [Kubeflow APIs and SDKs](https://www.kubeflow.org/docs/started/architecture/#kubeflow-apis-and-sdks).

Kubeflow CRDs are following Kubernetes best practices for
[the API changes](https://kubernetes.io/docs/concepts/overview/kubernetes-api/#api-changes).

#### Describe the project’s API topology and conventions

[Kubeflow APIs and SDKs documentation](https://www.kubeflow.org/docs/started/architecture/#kubeflow-apis-and-sdks).

#### Describe the project defaults

Some defaults can be seen [in the Kubeflow manifests](https://github.com/kubeflow/manifests#installation).

Additionally, API defaults are shown in the API docs, such as:
[Kubeflow Trainer](https://github.com/kubeflow/trainer/blob/master/pkg/apis/trainer/v1alpha1/trainjob_types.go#L128)
or [Kubeflow Katib](https://github.com/kubeflow/katib/blob/master/pkg/apis/controller/experiments/v1beta1/experiment_types.go#L46).

#### Outline any additional configurations from default to make reasonable use of the project

[Kubeflow manifests guides](https://github.com/kubeflow/manifests/blob/master/README.md).

Some public and listed distributions have their own ways
[to install Kubeflow AI platform](https://www.kubeflow.org/docs/started/installing-kubeflow/#packaged-distributions)

#### Describe any new or changed API types and calls—including to cloud providers—that will result from this project being enabled and used

Deploying Kubeflow manifests and individual projects results in exposing new APIs and the possibility
to call configured 3rd party APIs if integrations exist. These all have to be explicitly set by the users.

#### Describe compatibility of any new or changed APIs with API servers, including the Kubernetes API server

Kubeflow follows the same practice for [API compatibility as Kubernetes](https://kubernetes.io/docs/concepts/overview/kubernetes-api/#api-changes).

#### Describe versioning of any new or changed APIs, including how breaking changes are handled

Many Kubeflow projects use Kubernetes CRDs, and for these resources follows
[the same deprecation policy as Kubernetes](https://kubernetes.io/docs/reference/using-api/deprecation-policy/).

#### Describe the project’s release processes, including major, minor and patch releases

Every Kubeflow project follow its own release lifecycle, for example
[Kubeflow Trainer](https://github.com/kubeflow/trainer/tree/master/docs/release)
or [Kubeflow Katib](https://github.com/kubeflow/katib/tree/master/docs/release).

However, community also maintain [the Kubeflow AI reference platform](https://www.kubeflow.org/docs/kubeflow-platform/releases/)
releases which install all Kubeflow projects together for end-to-end AI platform:

### Installation

#### Describe how the project is installed and initialized

[Kubeflow Installation guide](https://www.kubeflow.org/docs/started/installing-kubeflow/).

Kubeflow projects can be installed as a standalone applications or together using
the Kubeflow Manifests or Kubeflow Distributions (public or private).

#### How does an adopter test and validate the installation

Distributions can verify the installation [by following this guide](https://github.com/kubeflow/manifests?tab=readme-ov-file#installation)
or [executing this test suites](https://github.com/kubeflow/manifests/blob/master/.github/workflows/full_kubeflow_integration_test.yaml).

Kubeflow projects docs also explain how platform admins should validate the installation of individual
applications, [for example Kubeflow Trainer](https://www.kubeflow.org/docs/components/trainer/operator-guides/installation/#installing-the-kubeflow-trainer-controller-manager)

### Security

#### Provide a link to the project’s cloud native [security self assessment](https://tag-security.cncf.io/community/assessments/).

- [Kubeflow Security Self Assessment](https://github.com/kubeflow/community/blob/master/security/self-assessment.md)

#### How are you satisfying the tenets of cloud native security projects

1. Make security a design requirement.

Kubeflow projects are built with security as a fundamental concern. Kubeflow projects follow
Kubernetes and cloud native security best practices. By leveraging Kubernetes Custom Resources
Definition (CRDs), Role-Based Access Control (RBAC), network policies, and pod security standards,
Kubeflow projects seamlessly integrate into cloud native secure environments. Additionally,
Kubeflow projects re-use functionality from other cloud native tools like Istio, Cert-Manager,
Kubernetes and inherit its security best practices.

Default configuration for Kubeflow projects follow best practices such as rootless containers,
limited privileges, etc.

2. Applying secure configuration has the best user experience

Default configuration for Kubeflow projects follow best practices such as rootless containers,
limited privileges, etc. Users can migrate workloads to more secure configurations without breaking
changes, leveraging Kubernetes’ declarative model.

3. Selecting insecure configuration is a conscious decision

Insecure options require explicit user setup and configurations.

#### Describe how each of the [cloud native principles](https://github.com/cncf/toc/blob/main/DEFINITION.md) apply to your project

Kubeflow uses cloud native principles by building on Kubernetes and other cloud native technologies
while extending them in our composable projects.

Kubeflow projects use containers as a fundamental unit of task/deployment. For example,
every TrainJob uses separate containers to train models, allowing users to package their training
code along with its dependencies in a containerized environment

Kubeflow projects are designed to run natively on Kubernetes and define various CRDs for each
AI workload. For example, TrainJob CR for model training or Notebook CR for interactive development.
It uses native Kubernetes primitives like Deployments, Services, Job, JobSet, and CRDs to
orchestrate AI workloads.

Kubeflow follows a loosely-coupled microservices architecture model, where each project can
be deployed as a standalone service and be integrated into desired AI platform on Kubernetes.

Kubeflow APIs are declarative and follow Kubernetes best practices. For instance,
SparkApplication to manage Spark jobs or TrainJob to manage distributed training jobs.

#### How do you recommend users alter security defaults in order to "loosen" the security of the project

Users can install Kubeflow Pipelines without multi-user isolation if they want to loosen the
security of deployment. Users suggested to follow the Kubeflow AI reference platform deployment
mode to install Kubeflow Pipelines control plane
[in multi-tenant model](https://www.kubeflow.org/docs/components/pipelines/operator-guides/multi-user/).
It provides isolation between Pipelines and Runs for users.

#### Security Hygiene

##### Describe the frameworks, practices and procedures the project uses to maintain the basic health and security of the project

- **Robust CI/CD infrastructure:** Kubeflow projects have automated unit, integration, and end-to-end
  tests that are integrated into the CI pipelines to ensure code stability and correctness of PRs.

- **Dependency management:** Kubeflow projects leverage lock files like `go.mod` and `pyproject.toml`,
  and progressing towards Software Bill of Materials (SBOM) generation.

- **Release stability:** Kubeflow projects maintain release branches and release tags which allow
  maintainers create patch releases to address bug fixes and security vulnerabilities. Also,
  projects maintain changelogs and roadmaps to ensure traceability and transparency for all changes.

- **Static and Dynamic Code Analysis:** Kubeflow projects CI integrates code linting, static,
  and dynamic code analysis tools to enforce code quality and detect potential issues before they
  reach production.

- **Secure Defaults:** Defaults manifests and Helm charts are configured with security best practices,
  like pod security standards, rootless containers, and enforces least privileges.

- **Code Review Process:** Kubeflow projects’ maintainers follow a strong review process with peer
  review and explicit approval before PRs being merged into the main branch.

- **Open Governance:** Kubeflow follows open governance with regular public meetings and
  communication channels to keep users aware of development and roadmap.

- **Secure Image Build:** Image builds are validated through CI/CD checks to ensure they are
  reproducible, security, and consistent.

Some examples can be found here:

- [Kubeflow Spark Operator](https://github.com/kubeflow/spark-operator/tree/master/.github/workflows)
- [Kubeflow Katib](https://github.com/kubeflow/katib/tree/master/.github/workflows)
- [Kubeflow Trainer](https://github.com/kubeflow/trainer/tree/master/.github/workflows)
- [Kubeflow Model Registry](https://github.com/kubeflow/model-registry/tree/main/.github/workflows)
- [Kubeflow Pipelines](https://github.com/kubeflow/pipelines/tree/master/.github/workflows)
- [Kubeflow Manifests](https://github.com/kubeflow/manifests/blob/master/.github/workflows/full_kubeflow_integration_test.yaml)

##### Describe how the project has evaluated which features will be a security risk to users if they are not maintained by the project

Features are discussed and reviewed in the community, with security implications considered during
design and code review processes. The project tracks the security health of dependencies and evaluates
the impact of vulnerabilities or unmaintained projects. The project identifies potential risks by
analyzing the attack surface, especially for features that interact with user-supplied code,
custom containers, user supplied models or external data sources.

#### Cloud Native Threat Modeling

##### Explain the least minimal privileges required by the project and reasons for additional privileges

Each Kubeflow project's control plane needs to watch for its CRDs. For example, Kubeflow Trainer:
TrainJob and ClusterTrainingRuntime, Spark Operator for SparkApplications, Katib for Experiments, etc.
Controllers have minimal privileges to orchestrate its CRDs and corresponding resources that need to be created

Kubeflow control plane is installed cluster-wide, and it requires RBAC to manage namespace-scoped resources.

The default controllers run with PodSecurityStandards restricted and rootless containers,
minimizing the risk of privilege escalation.

Users who interact with Kubeflow SDK get namespace-scoped permission to interact with required
resources. For example, users’ role for
[Kubeflow Trainer can be found here](https://github.com/kubeflow/trainer/blob/master/manifests/overlays/kubeflow-platform/kubeflow-trainer-roles.yaml#L17).
All roles are aggregated in the `kubeflow-edit` cluster role after installing Kubeflow Manifests.

##### Describe how the project is handling certificate rotation and mitigates any issues with certificates

Kubeflow projects use [cert-manager to generate and rotate certificates](https://github.com/kubeflow/manifests/tree/master/common/cert-manager).
Those certificates are used for validation and mutation webhooks across Kubeflow projects. Cert-manager
handles the issuance, renewal, and rotation of these certificates automatically without downtime.

##### Describe how the project is following and implementing [secure software supply chain best practices](https://project.linuxfoundation.org/hubfs/CNCF_SSCP_v1.pdf)

- **Automated CI/CD infrastructure:** Kubeflow projects have automated unit, integration, and
  end-to-end tests that are integrated into the CI pipelines to ensure code stability and reduce
  risk of introducing vulnerabilities.

- **Vulnerability Scanning:** Kubeflow uses tools for vulnerability scanning like Dependabot,
  Trivy to identify CVEs and address known CVEs in dependencies.

- **Dependency management:** Kubeflow projects leverage lock files like `go.mod` and `pyproject.toml`,
  and progressing towards Software Bill of Materials (SBOM) generation.

- **Code Review Process:** Kubeflow projects’ maintainers follow a strong review process with peer
  review and explicit approval before PRs being merged into the main branch.

- **DCO Check:** Committers are required to sign and comply with the Developer Certificate of Origin
  (DCO) to affirm the legitimacy and authorship of their contributions.

- **Branch Protection:** The project enforces branch protection rules to prevent unauthorized changes,
  enforce status checks, require pull request reviews, and control who can push to protected branches.

- **Prevent Secrets in Source Code:** Proactively prevents committing secrets to the source code
  by using tools and GitHub workflows that detect sensitive data in PRs.

- **License compliance:** Kubeflow uses
  [the FOSSA product](https://app.fossa.com/reports/bb8e2d41-254a-4af3-9044-e7f484c34dd1)
  provided by the CNCF to scan for license compliance on a regular and ongoing basis.

## Day 1 - Installation and Deployment Phase

### Project Installation and Configuration

#### Describe what project installation and configuration look like

Each project has its [own standalone installation guide](https://www.kubeflow.org/docs/started/installing-kubeflow/#standalone-kubeflow-components).

Kubeflow projects are packaged by [distributions and vendors](https://www.kubeflow.org/docs/started/installing-kubeflow/#kubeflow-platform)
into platform products which can be used to install our tools.

We provide optional manifests to deploy all Kubeflow projects as
[Kubeflow AI reference platform](https://github.com/kubeflow/manifests/blob/master/README.md).

### Project Enablement and Rollback

#### How can this project be enabled or disabled in a live cluster? Please describe any downtime required of the control plane or nodes

Users can set the replica count to 0 in the Kubeflow projects deployment. Existing AI workloads
should not be impacted since it won’t be reconciled by controllers.

Updating [the Kubeflow AI Reference Platform](https://github.com/kubeflow/manifests/blob/master/README.md#upgrading-and-extending).

The installation guide described [when control plane is ready](https://www.kubeflow.org/docs/components/trainer/operator-guides/installation/#installing-the-kubeflow-trainer-controller-manager).

#### Describe how enabling the project changes any default behavior of the cluster or running workloads

Enable projects by scaling replica count back to 1. All running AI workloads will be reconciled again
by controllers, and they perform the appropriate updates to the CRDs.

Istio, Knative, cert-manager might interfere with existing installations.

#### Describe how the project tests enablement and disablement

Conformance program [is work in progress](https://github.com/kubeflow/kubeflow/tree/master/conformance/1.7)
to ensure tests across all Kubeflow projects.

#### How does the project clean up any resources created, including CRDs

The Kubeflow projects control plane can be deleted by cleanup the appropriate resources,
for example:

```sh
kubectl delete -k https://github.com/kubeflow/trainer.git/manifests/overlays/manager?ref=v2.0.0
```

The command will cleanup all CRDs and control plane deployment. Since Kubeflow CRDs maintain
`ownerReference` for associated resources, all resources will be removed after deleting the
actual CRDs. For example:

```sh
kubectl delete trainjob --all --all-namespaces
```

Check [this guide to cleanup](https://github.com/kubeflow/manifests/blob/master/README.md#upgrading-and-extending)
the Kubeflow AI reference platform.

### Rollout, Upgrade and Rollback Planning

#### How does the project intend to provide and maintain compatibility with infrastructure and orchestration management tools like Kubernetes and with what frequency

Kubeflow projects publish its supported Kubernetes version for every release. The supported versions
are evaluated and upgraded on every release. We support
[Kubernetes 1.31+ and test on 1.32+](https://github.com/kubeflow/manifests/blob/master/tests/install_KinD_create_KinD_cluster_install_kustomize.sh)

#### How the project handles rollback procedures

Some projects support leader election and HA to make sure that long-running workloads are complete.
Users can scale up or down replicas of controllers during rollback.

#### How can a rollout or rollback fail? Describe any impact to already running workloads

Newer revision will not be active for many reasons including insufficient cluster resources for the
newer spec, invalid configuration spec. Traffic will not be switched to the new unless the newer
revision is ready to accept traffic. Hence, the already running workloads will not be affected
in any case if rollout fails.

#### Describe any specific metrics that should inform a rollback

Kubeflow projects CRDs expose status that inform about activity of AI workloads. Additionally,
controllers are exposed various Prometheus metrics to indicate workload status.

#### Explain how upgrades and rollbacks were tested and how the upgrade->downgrade->upgrade path was tested

Currently, it’s being manually tested by users, but automated tests are work in progress.

#### Explain how the project informs users of deprecations and removals of features and APIs

All API changes are backward compatible and changes are announced in the release notes and
changelogs. If some APIs are deprecated, newer versions of APIs are introduced with user awareness.
Kubeflow CRDs follow Kubernetes best practices for API compatibility.

For example [Kubeflow Trainer breaking changes](https://github.com/kubeflow/trainer/blob/master/CHANGELOG.md#breaking-changes).

#### Explain how the project permits utilization of alpha and beta capabilities as part of a rollout

At the moment, functionality is logged explicitly if a feature is alpha or beta. We do not have
feature gates. Additionally, we are working on conversion webhook that helps users with updated
API versions of CRDs.

## Day 2 - Day-to-Day Operations Phase

### Scalability/Reliability

#### Describe how the project increases the size or count of existing API objects

In general, the API object count complexity is linear in the number of users and workloads. Users
can list deployed Kubeflow CRDs across all namespaces.

For example, to list all TrainJob from Kubeflow Trainer run:

```sh
kubectl get trainjob --all-namespaces
```

The core controllers and resources are constant (beyond replicating specific controllers),
see [the full kustomize build here](https://github.com/kubeflow/manifests?tab=readme-ov-file#install-with-a-single-command).

#### Describe how the project defines Service Level Objectives (SLOs) and Service Level Indicators (SLIs)

Kubeflow currently doesn’t provide the SLOs and SLIs, however we follow
[the Kubernetes SLOs](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md)

#### Describe any operations that will increase in time covered by existing SLIs/SLOs

As above.

#### Describe the increase in resource usage in any components as a result of enabling this project, to include CPU, Memory, Storage, Throughput

Resources requirements for Kubeflow projects [are set here](https://github.com/kubeflow/manifests/pull/3091#issuecomment-3016609243).

#### Describe which conditions enabling / using this project would result in resource exhaustion of some node resources

There are some specific issues for each project, we should list some of them (e.g. pipelines
controller can have open file handler issues). Since many workloads are GPU-intensive,
Kubernetes platform admins need to ensure that nodes have enough capacity to run AI workloads with
Kubeflow projects.

#### Describe the load testing that has been performed on the project and the results

Some intensive load testing has been performed by Kubeflow users, for example some LLM foundation
models have been trained using Kubeflow Trainer across a few hundreds GPUs. At KubeCon some users
share their experience of running more than 10,000 GPUs with Kubeflow Training Operator.

#### Describe the recommended limits of users, requests, system resources, etc. and how they were obtained

Users on the order of hundreds are known to work. Scaling beyond hundreds of users may require an
increase in the requests/limits replicas of some deployments, or scaling dependencies like databases.

#### Describe which resilience pattern the project uses and how, including the circuit breaker pattern

Kubeflow uses the Kubernetes resilience pattern to manage controllers with HA. If one replica fails, another part of the control plane takes responsibility to orchestrate workloads.

### Observability Requirements

#### Describe the signals the project is using or producing, including logs, metrics, profiles and traces. Please include supported formats, recommended configurations and data storage

Kubeflow controllers expose Prometheus metrics to report workload status. Controllers also expose
logs and status for platform admins to ensure stability.

#### Describe how the project captures audit logging

Platform admins can leverage [Kubernetes audit](https://kubernetes.io/docs/tasks/debug/debug-cluster/audit/)
for Kubeflow projects.

#### Describe any dashboards the project uses or implements as well as any dashboard requirements

[Kubeflow Dashboard requirements](https://www.kubeflow.org/docs/components/central-dash/overview/).

#### Describe how the project surfaces project resource requirements for adopters to monitor cloud and infrastructure costs, e.g. FinOps That must happen on the Kubernetes namespace level

Users are recommended to use third-party tools like Kubecost to measure cloud and infrastructure
cost of running Kubeflow projects.

#### Which parameters is the project covering to ensure the health of the application/service and its workloads

Most project use Kubernetes [Liveness and Readiness probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).

For example [Kubeflow Trainer controller manager](https://github.com/kubeflow/trainer/blob/master/manifests/base/manager/manager.yaml#L30-L43).

#### How can an operator determine if the project is in use by workloads

- Check the Pods in `kubeflow-profil`e labeled namespaces.
- Check the CRDs in user’s namespaces
- Check the Kubeflow Dashboard resources.

#### How can someone using this project know that it is working for his instance

- [Kubeflow Manifests guide](https://github.com/kubeflow/manifests#installation).
- Run the [Kubeflow manifest test suite](https://github.com/kubeflow/manifests/blob/master/.github/workflows/full_kubeflow_integration_test.yaml)
- Run the upstream examples, for example [Kubeflow Trainer PyTorch training](https://github.com/kubeflow/trainer/blob/master/examples/pytorch/question-answering/fine-tune-distilbert.ipynb).

### Dependencies

#### Describe the specific running services the project depends on in the cluster

Most projects depend on Cert-Manager>=v1.16, and Istio>=v1.26

Specific projects have other dependencies:

- Kubeflow Pipelines: Argo Workflows>=v3.1
- Kubeflow Trainer: JobSet>=v0.8.0
- Kubeflow Katib: MySQL>=v8.0
- Kubeflow Model Registry: MySQL>=v8.0

For details, please take a look at General security Talks about Kubeflow including
[an architectural introduction](https://github.com/kubeflow/manifests#kubeflow-components-versions).

#### Describe the project’s dependency lifecycle policy

We follow Kubernetes deprecation policy for supported versions. For example, we support the
3-4 latest versions of Kubernetes to deploy Kubeflow projects.

#### How does the project incorporate and consider source composition analysis as part of its development and security hygiene? Describe how this source composition analysis (SCA) is tracked

Various static and dynamic code analysis tools are enforced as described above.

#### Describe how the project implements changes based on source composition analysis (SCA) and the timescale

Various static and dynamic code analysis tools are enforced as described above.

### Troubleshooting

#### How does this project recover if a key component or feature becomes unavailable

Specifics depend on the Kubeflow Project (KFP, Katib, etc.). Kubeflow projects are cloud and
Kubernetes native apps, so fault tolerance is strongly tied to the health of the underlying cluster.
In practice user workloads sometimes have a retry policy and the Kubeflow core services should
just recover automatically.

More details are described in the previous sections.

#### Describe the known failure modes

Each Kubeflow project handles failure modes differently beyond native Kubernetes fault tolerance.
Many of them are configured at the application level in user code.

### Security

#### How is the project executing access control

As described above each Kubeflow projects require limited RBAC to manage its CRDs

#### Cloud Native Threat Modeling

Described in the section above.

#### How does the project ensure its security reporting and response team is representative of its community diversity (organizational and individual)

Every Kubeflow project follows security guidelines.

- Kubeflow Spark Operator [security policy](https://github.com/kubeflow/spark-operator/blob/master/SECURITY.md).
- Kubeflow Notebooks [security policy](https://github.com/kubeflow/notebooks/blob/master/SECURITY.md).
- Kubeflow Trainer [security policy](https://github.com/kubeflow/trainer/blob/master/SECURITY.md).
- Kubeflow Katib [security policy](https://github.com/kubeflow/katib/blob/master/SECURITY.md).
- Kubeflow Model Registry [security policy](https://github.com/kubeflow/model-registry/blob/main/SECURITY.md).
- Kubeflow Pipelines [security policy](https://github.com/kubeflow/pipelines/blob/master/SECURITY.md).

#### How does the project invite and rotate security reporting team members

Active project maintainers are responsible to ensure that security reports are addressed. Each
project has its own security reporting and disclosure policy. For projects which use GitHub’s security
disclosure system, access control is managed by having write access to the relevant github repository.
