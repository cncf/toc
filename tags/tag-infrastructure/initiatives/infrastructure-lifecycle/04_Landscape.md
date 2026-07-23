# Use Cases & Landscape

The following section exemplifies the previously described patterns. Each use case describes the main pattern, relevant projects, and the trade-offs that should be considered before adoption. The section also provides an overview of the CNCF landscape for infrastructure lifecycle management.

## Case 1: Declarative Tooling

Declarative tooling allows operators to define the desired state of the infrastructure without describing every step required to create it. The underlying tool is responsible for ordering, dependency handling, retries, and applying the required changes.

Crossplane is the clearest example in the Cloud Native Computing Foundation (CNCF) landscape of this pattern. It extends the Kubernetes API with Custom Resource Definitions that represent infrastructure resources. This allows teams to provision and compose infrastructure using the same APIs, RBAC model, tooling, and GitOps workflows which are used for existing Kubernetes workloads. Crossplane controllers continuously reconcile the desired state against the actual state of the infrastructure, ensuring that any drift is corrected automatically.

Outside of the CNCF landscape, popular tools such as Terraform and OpenTofu use a similar declarative model but typically operate outside Kubernetes. They run as standalone tools and store state in a local or remote storage solution.

These tools are commonly used together with Kubernetes and Crossplane, but their operating models are different. Terraform and OpenTofu apply changes when they are explicitly executed, while Crossplane uses running controllers to continuously reconcile the defined state.

### Positives:
* Declarative code is easy to work with, and usually predictable.
* Infrastructure is defined as code 
* Changes can be reviewed before they are applied
* Drift can be detected when reviewing the current state against the desired state

### Negatives:
* Troubleshooting can be complex, for instance because the tool performs many actions implicitly.
* Complex infrastructure implementations can be hard to express in a declarative way, especially when they require conditional logic or loops.
* Reconciliation is not continuous in all tools, which can lead to drift if the desired state is not applied frequently enough.

## Case 2: Microservices - Continuously reconciled

Continuously reconciled systems use control loops that compare the current state with the desired state and automatically correct any differences.

Argo and Flux apply this model to Kubernetes workloads and infrastructure through GitOps. The desired state is stored in Git, while controllers running in the cluster continuously reconcile the environment against that state.

This model works particularly well for platforms with many microservices and frequent changes. Teams do not need to manually trigger every deployment or correction, and resources can automatically be restored if they are changed or deleted outside the defined workflow.

This flow can also be further extended by using GitOps tools outside of the CNCF landscape, such as the Flux Operator, which is built by the CNCF Flux core maintainers, licensed through controlplane.

Positives:
* Changes are applied automatically, which reduces the risk of drift and misconfiguration.
* The system is self-healing, which can reduce operational overhead.
* Troubleshooting is often easier, as the system can be observed and controlled through a single source of truth (e.g., Git).
### Negatives:
* Changes are not necessarily applied immediately.
* Dependent on the reconciliation interval and the state of the controller.
* Troubleshooting also requires operators to understand reconciliation status, events, and controller logs rather than following a single pipeline execution.

## Case 3: Programming Language driven

Programming-language-driven tooling allows infrastructure to be defined using general-purpose programming languages instead of structured configuration languages, like YAML and JSON.

cdk8s is an open-source software development framework which allows Kubernetes resources to be defined using languages such as TypeScript, Python, Java, and Go. The code is then compiled into standard Kubernetes manifests that can be applied to the cluster.

This approach makes it possible to use familiar software development concepts such as functions, loops, conditionals, packages, reusable components, and unit tests. It can therefore be easier to build and share more complex abstractions across teams. Furthermore, this is an imperative approach, which allows for more complex logic to be expressed than what is typically possible with declarative configuration.
    
Similar tools outside CNCF include Pulumi and the AWS Cloud Development Kit. These apply the same model to broader cloud provider APIs rather than only Kubernetes resources.

### Positives:
* Flexibility in defining complex infrastructure.
* Ability to use familiar programming constructs and testing frameworks.

### Negatives:
* Resulting state can be harder to inspect and compare.
* Requires developers to be familiar with the programming language and framework.

## Case 4: Policy-driven

Policy-driven infrastructure separates governance and enforcement from the tooling used to provision resources.

Open Policy Agent and Kyverno allow policies to be defined as code and applied to infrastructure and application configuration. Policies can validate, reject, or mutate resources during Kubernetes admission or as part of a CI/CD pipeline or in runtime during creation of resources.

This pattern can be used with both on-demand and continuously reconciled workflows. Policies are applied independently of how the resource was created, allowing platform teams to enforce security, compliance, operational, and cost requirements across multiple provisioning tools.

### Positives:
* Policies enforce security, compliance, operational, and cost requirements consistently across environments.
* Decouples governance from provisioning tools, allowing teams to choose the best tool for the job.

### Negatives:
* Policies must evolve together with the resources they validate.
* Poorly designed or overly restrictive policies can block valid changes.
* Policies should be tested, versioned, and introduced gradually.   

## Landscape

As the vendor-neutral home for cloud native software, CNCF hosts and fosters the
projects that give operators an end-to-end path through the infrastructure
lifecycle patterns described above. Across the Graduated, Incubating, and
Sandbox maturity levels, these projects share open governance, a Developer
Certificate of Origin (DCO) or Contributor License Agreement, and stewardship
by the community rather than by any single vendor:

* Provisioning and infrastructure APIs: Crossplane (Incubating) gives
  operators a Kubernetes-native API and control plane for declaratively
  provisioning and composing infrastructure.
* Continuous delivery and reconciliation: Argo (Graduated) and Flux
  (Graduated) provide GitOps-based continuous reconciliation for applications
  and infrastructure.
* Programming-language-driven definitions: cdk8s (Sandbox) lets teams
  author Kubernetes configuration using general-purpose programming languages.
* Policy and compliance: Open Policy Agent (Graduated) and Kyverno
  (Incubating) provide policy-as-code enforcement across the infrastructure
  lifecycle.

Together, these projects let operators build a complete cloud native
infrastructure lifecycle practice on software that is community-owned, vendor-neutral, and overseen by the Technical Oversight Committee (TOC) through a
transparent maturity process.

The rest of the landscape, outside of our governance, is made up of tools that
implement many of the same patterns under different stewardship models. Some,
such as Terraform, are commercially licensed and controlled by a single
vendor. Others, such as OpenTofu, are open source but governed directly by the
Linux Foundation rather than by CNCF. Others still, such as Pulumi and the AWS
Cloud Development Kit, are backed by a single company or cloud provider. These
tools are widely adopted and often used alongside CNCF projects in practice.

For operators, the distinction between what CNCF projects and the rest of the
landscape is useful when weighing licensing terms, vendor lock-in risk,
governance neutrality, and long-term roadmap control alongside the technical
merits of any given tool.

## Summary
