# LitmusChaos Sandbox Annual Review

## Table of Contents

- [Background](#background)
- [Alignment with Cloud Native Computing Foundation](#alignment-with-cloud-native-computing-foundation)
- [LitmusChaos's Year in Review](#litmuschaos-year-in-review)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

Litmus is a cloud-native chaos engineering framework. It provides a complete set of tools required by developers 
and SREs to carry out chaos experiments easily and in a Kubernetes-native way. The project has “Declarative Chaos” as the 
fundamental design goal and keeps the community at the center for growing the chaos experiments. 

Litmus has the following salient components:

- Chaos CRDs: Primarily, there are three chaos CRDs – ChaosExperiment, ChaosEngine and ChaosResult. Chaos is defined, run, and managed using the above CRDs. 
  ChaosEngine binds a target application (or infrastructure component) with a ChaosExperiment CR. When run by the operator, the result will be stored in ChaosResult CR.
  
- Chaos Operator: This operator is built using the Operator SDK framework and manages the lifecycle of a chaos experiment.

- Chaos Portal: Litmus Portal is a centralized web portal for creating, scheduling, and monitoring chaos workflows (that consist of one or more experiments).
  Chaos workflows can be scheduled on local or remote Kubernetes clusters registered with the portal. Teams can collaborate to manage chaos through the portal.
  
  ### Accepted into CNCF Sandbox
  
  - [Sandbox presentation slides](https://docs.google.com/presentation/d/1w94WgGmsgRXu4xFbkkVIPHtEl5V4eOg6FxgVikGimYs/edit#slide=id.g7e5d7aa8b9_2_75)
  - [Sandbox presentation video](https://youtu.be/O40GK24Ea4k?t=1106)
  - [Recommendation from SIG-App-Delivery chairs](https://docs.google.com/document/d/10G1fOySiYeCYIkWhz_ueyXTBi04_pQVRkDBdt2ryCFo/edit#heading=h.6pidoniplc5n)
  
  ## Alignment with Cloud Native Computing Foundation
  
  LitmusChaos falls in the scope of [CNCF App Delivery TAG](https://github.com/cncf/tag-app-delivery).
  
  Reliable application delivery and operations are very important and key to the success of technology adoption of various CNCF and 
  ecosystem projects. The practice of Chaos Engineering is a necessity for validating service reliability. LitmusChaos aims to become 
  the toolset to measure and achieve that reliability. The Chaos Hub today hosts application and platform specific experiments and more 
  CNCF project specific experiments will find their way into the hub. It is already being used by CNCF projects such as OpenEBS, Pravega, 
  Okteto and working groups like CNF-WG (as part of the CNF-TestSuite) for the mentioned purposes. 
  
  On the other hand, LitmusChaos also integrates with several projects in the CNCF landscape in the areas of app delivery, CI/CD and 
  Observability to enrich the usecases and help developers, SREs to have a seamless experience with regard to performing chaos tests. 

  In summary, LitmusChaos project clearly aligns with CNCF mission of making cloud native ubiquitous by being a readily usable and 
  easily extensible toolset for measuring and enhancing the reliability of Cloud Native applications and platforms.
  
  The following list captures a list of ecosystem projects that Litmus is integrated with: 
  
  Ansible 
  Argo Workflows
  Keptn
  GitHub Actions
  Gitlab
  Grafana
  Helm
  Okteto
  Prometheus
  Spinnaker
  
  ## LitmusChaos's Year in Review
  
  The following notable features were added over the past year: 
  
  - Litmus Portal to simplify management of chaos 
  - Support for chaos on containerd, crio runtime 
  - Golang SDK & migration to golang-based experiments 
  - Complex chaos workflows with argo integration 
  - Declarative Hypothesis via Litmus Probes
  - Different modes (namespaced, admin/cluster)  of operation
  - GitOps integration - Both for keeping chaos configuration at Git and to achieve event triggered chaos
  - Integration with other tools such as Argo Workflow, Keptn, GitLab, GitHub and Spinnaker
  - Interleaved chaos charts for open observability
  - Support for chaos against non kubernetes resources (infra components like VMs, Disks)

  Key Project Statistics: 
  
  ### Community Stats (graphs / dashboards) 

- Stars
  - 1.9k from 650 (~3x)

- Slack Membership
  - 780+ from 237 (~3x)

- Usage metrics 
  - Operator Installations / Experiment Run Count
    - Operator installations: 127268 
    - Experiments Run: 329008

  - New Experiments Added
    - 50+ (~2x)
  
