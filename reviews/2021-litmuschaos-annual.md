# LitmusChaos Sandbox Annual Review

## Table of Contents

- [Background](#background)
- [Alignment with Cloud Native Computing Foundation](#alignment-with-cloud-native-computing-foundation)
- [LitmusChaos Year in Review](#litmuschaos-year-in-review)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

Litmus is a cloud-native chaos engineering framework. It provides a complete set of tools required by developers 
and SREs to carry out chaos experiments easily and in a Kubernetes-native way. The project has “Chaos Experiment as a Custom Resource” as the 
fundamental design goal and keeps the community at the center for growing the chaos experiments. 

Litmus has the following salient components:

- Chaos CRDs: Primarily, there are three chaos CRDs – ChaosExperiment, ChaosEngine and ChaosResult. Chaos is defined, run, and managed using the above CRDs. 
  The ChaosExperiment is essentially a definition of the fault. ChaosEngine binds a target application (or infrastructure component) with a ChaosExperiment CR and 
  allows specification of run properties, steady-state hypothesis etc., When run by the operator, the result of the execution will be stored in ChaosResult CR.
  
- Chaos Operator: This operator is built using the Operator SDK framework and manages the lifecycle of a chaos experiment.

- Chaos Portal: Litmus Portal serves as a centralized orchestration platform that provides a web UI for creating, scheduling, and monitoring chaos workflows 
(that consist of one or more experiments). Chaos workflows can be scheduled on local or remote Kubernetes clusters registered with the portal. Teams can 
collaborate to manage chaos through the portal.

Amongst the most common reasons for its adoption as shared by its end-users are: 

- Availability of wide variety of off-the-shelf experiments via the ChaosHub
- Kubernetes-native approach to chaos which makes it easy to extend and integrate with CI/CD flows
- Ability to include steady-state hypothesis validation as part of automated experiment runs
- Adherence to/facilitates experimentation as per the standard principles of chaos 

The community support provided by LitmusChaos is another strong factor in its adoption and continued usage. 

  ### Accepted into CNCF Sandbox
  
  LitmusChaos was accepted as a [Sandbox project on 25th June 2020](https://github.com/cncf/toc/pull/391#issuecomment-649630917)
  
  - [Sandbox presentation slides](https://docs.google.com/presentation/d/1w94WgGmsgRXu4xFbkkVIPHtEl5V4eOg6FxgVikGimYs/edit#slide=id.g7e5d7aa8b9_2_75)
  - [Sandbox presentation video](https://youtu.be/O40GK24Ea4k?t=1106)
  
  ## Alignment with Cloud Native Computing Foundation
  
  LitmusChaos falls in the scope of [CNCF App Delivery TAG](https://github.com/cncf/tag-app-delivery).
  
  Reliable application delivery and operations are very important and key to the success of technology adoption of various CNCF and 
  ecosystem projects. The practice of Chaos Engineering is a necessity for validating service reliability. LitmusChaos aims to become 
  the toolset to measure and help achieve that reliability. The Chaos Hub today hosts application and platform specific experiments and more 
  CNCF project specific experiments will find their way into the hub. It is already being used by CNCF projects such as OpenEBS, Pravega, 
  Okteto and working groups like CNF-WG (as part of the CNF-TestSuite) for the aforementioned purposes. 
  
  On the other hand, LitmusChaos also integrates with several projects in the CNCF landscape in the areas of app delivery, CI/CD and 
  Observability to enrich the usecases and help developers, SREs to have a seamless experience with regard to performing chaos tests. 

  In summary, LitmusChaos project clearly aligns with CNCF mission of making cloud native ubiquitous by being a readily usable and 
  easily extensible toolset for measuring and enhancing the reliability of Cloud Native applications and platforms.
  
  The following list captures a list of ecosystem projects that Litmus is integrated with or utilizes: 
  
  - Ansible: Litmus supports building chaos experiments as ansible playbook and even provides an [SDK/scaffolding utility](https://github.com/litmuschaos/litmus-ansible/tree/master/contribute/developer_guide) to bootstrap the experiments. 
  
  - Argo Workflows: Litmus Portal uses [argo Workflows](https://dev.to/ksatchit/chaos-workflows-with-argo-and-litmuschaos-2po5) as a unit of execution. A Chaos Workflow is an [enhanced Argo Workflow](https://dev.to/litmus-chaos/workflow-creation-with-litmus-2-0-beta-29nc) that executes one or more chaos experiments and has the intelligence to collate the results as per the specified hypothesis 
 
  - Keptn: The Keptn [litmus service](https://dev.to/litmus-chaos/part-2-evaluating-application-resiliency-with-keptn-and-litmuschaos-use-case-and-demo-2fl) provides a mechanism to run chaos experiments on event triggers as part of Keptn CD pipeline.
  
  - GitHub Actions: Litmus provides [chaos actions](https://dev.to/uditgaurav/setup-github-chaos-action-using-kind-cluster-part-2-4dbm) that can be used in GitHub Workflows to effect specific experiments on transient or staging clusters. 
  
  - Gitlab: Litmus experiments are available as GitLab [remote templates](https://dev.to/uditgaurav/litmuschaos-gitlab-remote-templates-6l2) that can be included into .gitlab-ci.yml files as part of a pipeline.
  
  - Helm: The Litmus control plane & experiments are bundled as [helm charts](https://dev.to/ksatchit/deployment-modes-in-litmuschaos-1bpa) with control for nature of installation (admin/namespaced mode)
  
  - Okteto: The Okteto CLI is used as part of the [chaos experiment test procedure](https://dev.to/ksatchit/litmus-sdk-devtest-your-chaos-experiments-with-okteto-4dkj) while the [Okteto Cloud](https://okteto.com/blog/okteto-and-litmus-webinar) (SaaS) has an inbuilt template for Litmus Installation. 
  
  - Prometheus: Litmus both generates (chaos) Prometheus metrics as well as consumes (app/infra metrics) it as part of its [observability functionality](https://dev.to/ksatchit/observability-considerations-in-chaos-the-metrics-story-6cb). The chaos metrics
  are used to achieve chaos "interleaved" dashboards while the app metrics are used by the promProbe to validate hypothesis. 
  
  - Spinnaker: Litmus provides a [preconfigured job plugin](https://www.armory.io/blog/litmuschaos-in-your-spinnaker-pipeline/) to add chaos stages in a spinnaker pipeline
  
  ## LitmusChaos Year in Review
  
  The past year saw overall development of the project in terms of features, community growth and adoption. 
  
  The following notable features were added:  
  
  - Litmus Portal to simplify management of chaos across clusters
  - Support for chaos on containerd, crio runtime (along with existing docker support)
  - Golang SDK & migration to golang-based experiments 
  - Support for creation of 'scenarios' involving multiple faults via Argo-based Chaos Workflows 
  - Declarative Hypothesis definition and validation via Litmus Probes
  - Continuous chaos via flexible schedule, with support to halt/resume or (manual/conditional) abort experiments
  - Different modes of operation (namespaced, admin/cluster) 
  - GitOps support - Git as the single source of truth for chaos definition 
  - Support for ARM64 nodes
  - Integration with other tools such as Keptn, GitLab, GitHub and Spinnaker
  - Enhanced chaos metrics and chaos-interleaved app dashboards for observability
  
  ### Key Project Statistics  

  | Statistic | Sandbox | Current | Multiplier |
  |-|-|-|-|
  | [GitHub stars](https://github.com/litmuschaos/litmus/stargazers) | 650 | 1.9k | ~3x |
  | [Commits](https://github.com/litmuschaos) | 2654| 6623 | >2x |
  | [Releases](https://github.com/litmuschaos/litmus/releases) | 14 | 50+ | >3x |
  | Contributors (PRs,commits,issues,comments) | 90 | 173 | ~2x |
  | Slack Membership (#litmus in K8s slack) | 213 | 783 | >3x |
  | Youtube Subscribers | 32 | 332 | >10x |

## Annual Review Contents

### Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

The project averages around [255 commits from around 27 contributors per month](https://litmuschaos.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) contained within [186 PRs](https://litmuschaos.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) on an average per month over the past year. 

There are around [1236 contributions (across PRs, issues, comments, reviews) from 50 contributors](https://litmuschaos.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) on an average per month

### How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

Currently, the project has [9 maintainers](https://github.com/litmuschaos/litmus/blob/master/MAINTAINERS) along with a growing list of reviewers who contribute and lead specific sub-projects/areas within Litmus. The process of adding maintainers is detailed in the [Governance](https://github.com/litmuschaos/litmus/blob/master/GOVERNANCE.md) documentation. 

### What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

LitmusChaos adoption has grown significantly since becoming a sandbox project. End users of Litmus include organizations that meet the CNCF end user criteria apart from well-known vendors (that use it for the resilience evaluation of their cloud-native projects/products and other internal devops needs). Some of them are listed in the [ADOPTERS](https://github.com/litmuschaos/litmus/blob/master/ADOPTERS.md) file. There are other adopters/users of litmus that haven't added themselves but have consented to speak with the CNCF ToC (some of these are in the [public domain](https://www.youtube.com/watch?v=PiGqQc_UYPc&feature=youtu.be), while others prefer to maintain the details private). 

### How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

When we started the Litmus project, the goal of this project was to create a complete platform to practice chaos engineering at scale in a Kubernetes-native way. 
This had to be done incrementally, first create a toolset for chaos injection and then add additional features to make it a platform. Litmus, when it was accepted 
into sandbox had achieved the first goal of creating a 'toolset' that was  open source, had a ChaosHub to search/pull readymade experiments along with the required 
CRDs and Operators. 

The next set of features we focused on the past year were to tackle the challenges around (a) chaos experimentation process and (b) operations around chaos engineering to create a comprehensive 'platform', that resulted in Litmus 2.0 (in advanced beta, fully backward compatible with existing Litmus 1.x releases): 

**Experimentation Process**

- Visualizing impact of chaos 
- Ability to validate application behaviour during experiment runs to determine verdict 
- Creation of complex scenarios involving more than one fault, in desired order/sequence
- Add conditions for auto-abort of experiments 
- Associate a resiliency metric that connects the application and the chaos scenario  

**Chaos Engineering Operations**

- Enabling teams to collaborate on chaos experimentation
- Management of chaos across a fleet of clusters from a single control plane
- Single source of truth for chaos artifacts & bringing chaos into the ambit of GitOps 
- Operating the chaos framework in airgapped Environments 

Apart from the technical features, other project goals (based on [recommendations](https://docs.google.com/document/d/10G1fOySiYeCYIkWhz_ueyXTBi04_pQVRkDBdt2ryCFo/edit#heading=h.6pidoniplc5n) by the TAG-APP-Delivery) are being pursued: 

- Growing the community (refer: Key Project Stats)
- Increasing awareness around chaos engineering (Cloud Native Chaos was discussed in several CNCF meetups hosted by the respective groups in Bangalore, Linz, Madison, Switzerland etc., apart from conducting bootcamps/workshops to explain chaos engineering in general. Litmus featured significantly in the [ChaosCarnival 2021](https://docs.google.com/presentation/d/1vIO41Pbybf9viQH1_1fNf0Ex66IPXwpKkvAJBQglE4w/edit#slide=id.p))
- Begin contributing to resilience testing efforts of other CNCF projects/Workgroups (Ex: Resiliency suite of the CNF-Conformance TestSuite)

### What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

- LitmusChaos aims to fulfill the chaos testing needs of users that are yet to migrate to Kubernetes fully or operate hybrid environments (examples: some services on 
Kubernetes & others on cloud instances or VMWare VMs/baremetal, services across several cloud providers etc.,) while allowing reuse of a common control plane and 
using the same declarative approach to define the experiment. While this capability is available to a certain extent already via out-of-band/provider-api-driven 
experiments (instance and disk failures on AWS, GCP, Azure, VMWare) more work needs to be done in this area. 

  Other in-progress features and live backlog can be viewed 
in the project [ROADMAP](https://github.com/litmuschaos/litmus/blob/master/ROADMAP.md#in-progress-under-active-development)

- The community is also working on simplifying user onboarding and contributions via improved documentation (terms of structure/information flow, content & infrastructure). 
 This is another significant activity that is being focused on.

- The project Governance is set to improved further with the introduction of Special Interest Groups within the project to drive better decision making and aid all-round
development of the project via prioritization and adequate representation in the monthly release cadence. The concept is being trialed actively in the areas of Documentation 
& Observability, with folks from different organizations leading/chairing the group and driving requirements along with the project maintainers. 

- As part of the efforts to build awareness around chaos engineering in general and the influence/challenges brought about by the cloud native paradigm, the LitmusChaos 
community is working to set-up a [cross-TAG/project working group](https://docs.google.com/document/d/1scr9uuvG1g1xpIHPs3314FqeFufE31ustTVnRMrX3gI/edit) that is expected 
to become operational and produce a whitepaper summarizing findings and practices. 

### How can the CNCF help you achieve your upcoming goals?

The project is receiving great pointers and suggestions from TAG-App-Delivery, under the purview of which Litmus was accepted into Sandbox. 
The office-hours and project booth during KubeCon are very helpful in introducing the project to interested folks. Events like Container Garage are also great 
to discuss/talk about the technology and appraise the community on the capabilities. We would like to acknowledge the support and thank CNCF for the same. 

- The CNCF can help us in providing increased visibility for the project to drive further adoption via blogs, webinars or suggest channels to reach out to the end-user 
groups to present some compelling stories that we have within the litmus community as also to gather more feedback around requirements and pain points. 

- The CNCF can help in bringing together like-minded folks to create/formalize the chaos working group so we can strengthen the agenda and get broader viewpoints 
on the board.

- We look forward to receiving guidance/tooling around security auditing, CI infrastructure and creative services that can make the project better. 

- As part of the project goals for the upcoming year, there is increased collaboration planned with other CNCF projects (in terms of providing a joint solution/integration as
well as help them by enabling project(app)-specific experiments). A forum to ease these interactions or present findings jointly would be much beneficial. 

### Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

Yes, we've also filed our [incubation proposal](https://github.com/cncf/toc/pull/569).

## Project Links
 - [Website](https://litmuschaos.io)
 - [Github](https://github.com/litmuschaos/litmus)
 - Slack #litmus in [Kubernetes](https://slack.k8s.io/) workspace
 - [Twitter](https://twitter.com/litmuschaos)




