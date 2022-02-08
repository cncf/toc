# Kyverno for incubating 

## Table of Contents

- [Background](#background)
- [About Kyverno](#about-kyverno)
- [Progress Since Joining CNCF](#progress-since-sandbox)
- [Incubation Stage Requirements](#incubation-stage-requirements)
- [Statement on alignment with the CNCF mission](#statement-on-alignment-with-the-cncf-mission)
- [Comparison with other CNCF Policy Management projects](#comparison-with-other-cncf-chaos-engineering-projects)
- [Roadmap](#roadmap)

## Background

Kubernetes has a declarative configuration management system that allows users to specify the desired state of resources which Kubernetes controllers continuously reconcile with the current system state. Kyverno uses Kubernetes custom resource definitions for cluster-wide and namespaced policies, and the [Kubernetes Policy WG policy report](https://github.com/kubernetes-sigs/wg-policy-prototypes/tree/master/policy-report) for reporting policy results.

Kyverno was accepted as a CNCF Sandbox project on [November 9th, 2020](https://docs.google.com/spreadsheets/d/1Nnh_usr0tSZxaUpxTusqeIqKxMmvuEViRkyO9e_Do40/edit#gid=1136111842).

## About Kyverno

Kyverno is a policy engine designed for Kubernetes. With Kyverno, policies are managed as Kubernetes resources and no new language is required to write policies. This allows using familiar tools such as `kubectl`, `git`, and `kustomize` to manage policies. Kyverno policies can validate, mutate, and generate Kubernetes resources. The Kyverno CLI can be used to test policies and validate resources as part of a CI/CD pipeline.

### High-level Architecture

Kyverno comprises of four major components, each component is described below.

**Webhook**: The `Webhook` component registers as a validating and mutating admission webhook and receives `AdmissionReview` requests from the API server to validate and mutate configuration changes, based on policies. Users can configure which namespaces and resources the webhooks will receive via command line options or the ConfigMap.

The `Webhook` also creates and updates `GenerateRequest` and `PolicyChangeRequest` resources to trigger updates via other Kyverno controllers.

**Webhook Monitor**: On startup, Kyverno's `Webhook Monitor` component generates a self-signed certificate (or uses a user-provided certificate) and auto-creates the webhook configurations required to register Kyverno as an admission webhook. The component also periodically monitors if Kyverno is receiving webhook events and recreates the certificate and webhook configurations if needed.

**Generate Controller**: The `Generate Controller` watches `GenerateRequest` resources and creates, updates, and deletes Kubernetes resources based on Kyverno [generate rules](https://kyverno.io/docs/writing-policies/generate/). The `Generate Controller` also watches for changes in policy definitions to update generated resources.

**Policy Controller**: The `Policy Controller` performs periodic background scans on existing configurations and creates or updates policy reports based on changes and background scans. The `Policy Controller` watches `ReportChangeRequest` resources and creates, updates, and delete Kyverno [Policy Report](https://kyverno.io/docs/policy-reports/) resources. The `Policy Controller` also watches for changes in policies definitions to update policy reports.

## Progress Since Sandbox

### Feature Summary

- Support multiple replicas for high availability and scalability
- Improved observability via Prometheus Metrics over both the state and execution of cluster-applied Kyverno policies
- Implement Kubernetes [Pod Security Standards definitions](https://kubernetes.io/docs/concepts/security/pod-security-standards/) with compatbility testing across versions
- Standalone Kyverno CLI to check for policy compliance in CI/CD pipelines
- Support for unit tests via CLI for policy as code best practices
- Introduce `JMESPath` support for complex processing
- `API Server Lookups` and `ConfigMap Lookups` for data-driven policies
- Add `foreach` block for processing sub-elements in resource declarations
- Dynamic webhook management with a fail-closed default based on the configured policy
- Automated scale testing to measure the performance across a wide-range of clusters settings
- Optimize memory usage for large clusters with hundreds of namespaces and thousands of resources
- Ability to fetch and use image configuration data from OCI registries
- Integrations with sigstore for software supply chain security verification of signatures and attestations

### Project Governance

The Kyverno community has matured significantly. We have on-boarded maintainers across different companies, and we host weekly contributors meetings and monthly community meetings. We codified our [governance processes](https://main.kyverno.io/community/#project-governance), [contributing guidelines](https://github.com/kyverno/kyverno/blob/main/CONTRIBUTING.md), and [security processes](https://main.kyverno.io/docs/security/).


### Community Stats (graphs / dashboards) 
The Kyverno DevStats dashboards can be found [here](https://kyverno.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&viewPanel=4).

We are using the following metrics as key indicators of community health:

- [Stargazers](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Stargazers&var-table=swatchers&var-pref=all&var-met1=watch&var-met2=watch&from=now-1y&to=now): 372% growth from 524 -> 1949.
- [Issue Creators](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Issue%20creators&var-table=scommunity_health&var-pref=&var-met1=chealthissue&var-met2=&from=now-1y&to=now): 602% growth from 45 -> 271.
- [Code Committers](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Code%20committers&var-table=scommunity_health&var-pref=&var-met1=chealthcommit&var-met2=&from=now-1y&to=now): 594% growth from 18 -> 107.
- [Code Commenters](https://kyverno.devstats.cncf.io/d/81/community-health?orgId=1&var-repo_name=All&var-metric=Code%20commenters&var-table=scommunity_health&var-pref=&var-met1=chealthcomment&var-met2=&from=now-1y&to=now): 573% growth from 22 -> 126.

Slack has grown to 722 members since July 2019, and Twitter account has added 738 followers since Sep 2020.

### References / Highlights

**Presentations by community members**:
- [Kubernetes policy management with Kyverno](https://www.containerconf.de/veranstaltung-13729-0-kubernetes-policy-management-mit-kyverno.html), Nov 2021
- [Utilizing Kubernetes and Open OnDemand to Support Virtual Classroom Labs](https://sc21.supercomputing.org/presentation/?id=exses102&sess=sess479), Nov 2021
- [Kubernetes with Open OnDemand Using Kyverno](https://sc21.supercomputing.org/presentation/?id=ws_hpcsysp102&sess=sess326), Nov 2021
- [Replacing PSPs? Keep Bad Pods out of your cluster using Kyverno!](https://www.youtube.com/watch?v=AmJUFH7n33c&t=290s), Oct 2021
- [Cloud Native Live: SLSA with Cosign and Kyverno to secure software delivery](https://www.cncf.io/online-programs/cloud-native-live-slsa-with-cosign-and-kyverno-to-secure-software-delivery/), Oct 2021
- [Kubernetes Philly, October 2021 - Kubernetes Policy Management with Kyverno](https://www.youtube.com/watch?v=Am7mvIQWx4E&ab_channel=KubernetesPhilly), Oct 2021


**Reviews and blogs**:
- [Let's Deploy : kyverno](https://www.youtube.com/watch?v=Y3m6cYS3ytM&ab_channel=GuilhemLettron), by Guilhem Lettron, Jan 2022
- [Securing Kubernetes with Kyverno](https://cloudyuga.guru/hands_on_lab/kyverno-introduction/), by Oshi Gupta, Jan 2022
- [No! You shall not pass! Kyverno’s here!](https://medium.com/@ul_Timate/no-you-shall-not-pass-kyvernos-here-def1d376a4f8), by Kamonnop Arunrat, Dec 2021
- [Admission Controller for Secure Supply Chain Verification - Kyverno](https://boxboat.com/2021/12/06/secure-supply-chains-kyverno/), by Parth Patel, Dec 2021
- [K8s policy with Kyverno](https://www.blakyaks.com/resources/k8s-policy-with-kyverno), Oct 2021
- [Enforcing Policy as Code using Kyverno in Kubernetes](https://dev.to/aws-builders/enforcing-policy-as-code-using-kyverno-in-kubernetes-3epk), by Sivamuthu Kumar, Oct 2021
- [A Zero Trust Approach for Securing the Supply Chain of Microservices Packaged as Container Images](https://gkovan.medium.com/a-zero-trust-approach-for-securing-the-supply-chain-of-microservices-packaged-as-container-images-89d2f5b7293b), by Gerry Kovan, Sep 2021
- [Automate Your Security Practices and Policies on OpenShift With Kyverno](https://cloud.redhat.com/blog/automate-your-security-practices-and-policies-on-openshift-with-kyverno), by Gareth Healy, Sep 2021
- [Kyverno - A Kubernetes native policy manager (Policy as Code)](https://blog.mimacom.com/kyverno/), by Arun Singh Sisodiya, Sep 2021
- [Simplify Kubernetes Cluster Management with Kyverno](https://movi.hashnode.dev/simplify-kubernetes-cluster-management-with-kyverno-ckt6yxjqy0duy95s14groe7h4), by Anita Ihuman, Sep 2021


### Integrations with CNCF Ecosystem
- [FluxCD](https://github.com/fluxcd/flux2-multi-tenancy#enforce-tenant-isolation): Kyverno policies to enforce tenant isolation
- [LitmusChaos](https://litmuschaos.github.io/litmus/experiments/concepts/security/kyverno-policies/): manage permissions for execution of litmus experiments
- [OpenEBS](https://github.com/openebs/charts/tree/main/charts/openebs/templates/kyverno): build Kyverno policies to secure configurations
- [Cert-Manager](https://github.com/jetstack/cert-manager/tree/master/devel/addon/kyverno): adopt Kyverno Pod Security Standards policies


## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._**

Full list of our adopters can be found [here](https://github.com/kyverno/kyverno/blob/main/ADOPTERS.md).

- [Ohio Supercomputer Center](https://www.osc.edu/): Support Kubernetes multi-user workflows through Open OnDemand.
- [Coinbase](https://www.coinbase.com/): Use Kyverno for mutation, to replace hand-written Webhooks, and generation to project common Kubernetes objects into many similar namespaces.
- [Vodafone Group Plc](https://www.vodafone.com/): Policy enforcement and automation on an internal k8s service offering.
- [Deutsche Telekom](https://www.telekom.com/en): Use Kyverno to enforce Policies on managed clusters to prevent right escalation of internal customers and to enforce security rules.

Additional end users are available for private references.

**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

- There are 8 maintainers for Kyverno from five different organizations. The maintainers list can be found [here](https://github.com/kyverno/kyverno/blob/main/MAINTAINERS.md)

**_Demonstrate a substantial ongoing flow of commits and merged contributions._**
- [Commit Activity](https://kyverno.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=1604851200000&to=now)
- [New Contributors added since Sandbox](https://kyverno.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1604851200000&to=now): 93
- [Issues open/closed since Sandbox](https://kyverno.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1604851200000&to=now&var-period=m&var-repogroup_name=All): 885/1.09k
- [Issues open/closed since Sandbox](https://kyverno.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1604851200000&to=now&var-period=m&var-repogroup_name=All): 1.16k/1.21k
- [No. of contributing organizations in the past two years](https://kyverno.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%202%20years&var-metric=contributions)
  - More than 20 companies made over 20 contributions
  - A wide range of contributions over 90 companies

**_A clear versioning scheme._**

Kyverno uses the [Semantic Versioning scheme](https://semver.org/#semantic-versioning-specification-semver). Kyverno v1.0.0 was released in Dec 2019. This project follows a given version number MAJOR.MINOR.PATCH. Details can be found [here](https://main.kyverno.io/docs/release-management/#versioning).

## Statement on Alignment with the CNCF Mission

A recent [CNCF security survey](https://www.cncf.io/blog/2021/10/12/cloud-native-security-microsurvey-more-than-80-of-organizations-want-to-build-modern-security-systems-with-open-source-software/) revealed that more than 80% of organizations are looking for modern security systems with open source software. Kyverno's goal is to help address and close this gap. Kyverno secures and automates Kubernetes configurations by providing a powerful but easy-to-use policy engine that allows developers, operators, and security engineers to leverage policies as contracts for collaboration. By making Kubernetes easier to use and secure, Kyverno's mission directly aligns with the Foundations mission of making cloud native ubiquitious.

## Comparison with Other CNCF Policy Management Projects

### Comparison to OPA/Gatekeeper

Open Policy Agent (OPA) is a general-purpose policy engine where policies are defined using Rego. [Gatekeeper]([OPA/Gatekeeper](https://open-policy-agent.github.io/gatekeeper/website/docs/) enables OPA/Rego policies to be used in Kubernetes and allows customization of Rego policies using constraints.

Kyverno was designed for Kubernetes and it's policies leverage the OpenAPIv3 structural schema used for Kubernetes resources. In addition, Kyverno uses Kubernetes patterns, idioms, and best-practices, so it feels familiar to Kubernetes DevSecOps teams. Kyverno policies, and policy results, are custom resources that can be managed using the same declarative configuration management principles used in Kubernetes and by widely adopted tools like kubectl and Kustomize. 

Here is the video and the blog published by community compares policy management tools for Kubernetes:
- [Kubernetes Policy Management Tools Compared - OPA with Gatekeeper vs. Kyverno](https://www.youtube.com/watch?v=9gSrRNmmKBc&feature=youtu.be&ab_channel=DevOpsToolkit), by Viktor Farcic
- [Kubernetes Policy Comparison: OPA/Gatekeeper vs Kyverno](https://neonmirrors.net/post/2021-02/kubernetes-policy-comparison-opa-gatekeeper-vs-kyverno/), by Chip Zoller

## Roadmap

Kyverno maintains a steady pace of releases driven by a growing set of requirements and use cases. The planned releases can be viewed using the issue [milestones](https://github.com/kyverno/kyverno/milestones?direction=asc&sort=due_date&state=open). 

Some longer term features may include:
- Flexible deployment models for Kyverno beyond admission controls
- Adopting Common Expressions Language (CEL), or similar processors
- Expansion of use cases for applications security and automation
- SLSA Level 3 compliance for build security
