Hi, 

We would like to propose to donate LitmusChaos to CNCF as a SandBox project. We have been advised to follow the new process outlined [here](https://github.com/cncf/toc/blob/master/process/project_proposals.md). The template we are following is outlined [here](https://github.com/cncf/toc/issues/344). 

Please consider this proposal and guide us through the process.

Authors:
- Uma Mukkara, @umamukkara, MayaData, Litmus Maintainer
- Karthik S, @ksatchit, MayaData, Litmus Maintainer

---
# Background

**Link to TOC PR**  - This one

**[Link to Presentation](https://docs.google.com/presentation/d/1w94WgGmsgRXu4xFbkkVIPHtEl5V4eOg6FxgVikGimYs/edit#slide=id.g7e62827add_18_0)**

**[Link to GitHub project](https://github.com/litmuschaos/litmus)**


## Project Goal

The goal of LitmusChoas project is to provide infrastructure toolset to do end to end chaos engineering in Kubernetes environments, in a cloud-native way. By making the practice of chaos engineering easy in a highly scaled environment involving many applications, Litmus helps developers and SREs take control of reliability of the systems. Litmus is bringing application developers and users (SREs) together through chaos hub. Developers submit their chaos experiments in the CI environments to the hub and SREs bring the chaos experiments to the deployments or operations.

## Current Status

- **Project releases:** Litmus achieved GA status or 1.0 release in January 2020. Last latest release is 1.3. The project has made [21 releases](https://github.com/litmuschaos/litmus/releases) so far.

- **Community status:** Litmus has monthly contributor calls. The community notes is [here](https://hackmd.io/a4Zu_sH4TZGeih-xCimi3Q). Other stats on  the community: 
  - 650+ stars on GitHub
  - 50+ contributors, incl.
    - MayaData
    - Intuit
    - Wipro
    - ArgoAI
    - Zebrium
    - RingCentral/Dino Systems
    - Container Solutions
    
  References from community adoption: 
  
  - https://www.zebrium.com/blog/using-autonomous-monitoring-with-litmus-chaos-engine-on-kubernetes
  - https://youtu.be/Uwqop-s99LA?t=720 (Intuit has demonstrated Litmus Argo integration and usage in the recent community syncup meeting)
  - https://youtu.be/M7nL1sqQPOU (Andreas Krivas from container solutions has discussed how they have used Litmus in large retail customer environment for Chaos Engineering)

Community is very active both in using the project and reporting issues/enhancements. A list of issues tagged with (project/community) is provided here https://github.com/litmuschaos/litmus/issues?q=is%3Aissue+label%3Aproject%2Fcommunity+is%3Aopen


Chaos experiments are hosted on https://hub.litmuschaos.io. It is a central hub where the application developers or vendors share their chaos experiments so that their users can use them to increase the resilience of the applications in production.

## Completed Roadmap items include: 
Note: We make monthly releases and features are continuously being churned. For an updated roadmap, use this URL ( https://github.com/litmuschaos/litmus/blob/master/ROADMAP.md )

- Declarative Chaos Intent via custom resources
- Chaos Operator to orchestrate chaos experiments
- Off the shelf / ready chaos experiments for general Kubernetes chaos
- Per-experiment minimal RBAC permissions definition
- Helm3 charts for Litmus Chaos (operator, kubernetes/generic chaos charts)
- Support for Kubernetes events for chaos experiments
- Support for admin mode (centralized chaos management from single namespace)
- Centralized Hub for chaos experiments
- Documentation (user & developer guides, integration with other chaos tools)
- Gitlab e2e pipeline for chaos experiments
- Define community sync up schedule

## Future Plans

In summary, LitmusChaos project has achieved the milestone of having a basic infrastructure toolset to do chaos engineering 
on Kubernetes. In the short term, the plan is to add chaos monitoring support and many other application chaos charts to chaos 
hub. In the long term, LitmusChaos team wants to work with many projects or applications in the CNCF landscape to integrate 
Litmus into their chaos engineering or reliability engineering needs.

#### In-Progress (Near-term)

- A UI portal for LitmusChaos to trigger and schedule chaos experiments & workflows. 
- Off the shelf chaos-integrated grafana dashboards for OpenEBS, Kafka, Cassandra #1280
- Support for scheduled (continuous/background) chaos with halt/resume #1223
- Support for OpenShift platform/resources 1406
- Support for complete chaos abort via pre-stop hooks #1284
- Go, Python SDK for Litmus Experiments 1466 #1259
- Support for user defined chaos experiment result definition (ex:json blob as chaos result) #1254
- Pod level resource chaos libraries (memory, disk stress) #877
- HTTP proxy Chaos libraries #1179
- Support for chaos on containerd runtime #1245
- Self-sufficient ChaosHub with downloadable sample chaosengine, experiment-level RBAC manifests & versioned chaos charts #1228
- Support for custom override of chaos-operator, chaos-runner and chaos-experiment attributes #1253 #1252 #1227
- Integration with Argo project for Chaos workflows, support for scenario creation with experiments1465
- Detailed design #1282

#### Backlog

- Support for Kubernetes pod scheduling policies (affinity rules for chaos resources)
- Support for cloudevents compliant chaos events
- Kubectl plugin for CLI based execution of chaos experiments
- Increased chaos metrics via prometheus chaos exporter
- CI (Gitlab) chaos templates
- Migration to native Kubernetes ansible modules for ansible-based experiments
- Improved application Chaos Suites (OpenEBS, Kafka, Cassandra)
- Support for platform (AWS, GKE, vSphere) Chaos

# Project Scope

## Clear project definition

The project aims to provide a chaos engineering framework - comprising of a chaos orchestrator, off-the-shelf chaos experiments for standard cloud native applications, and a central hub for forge collaboration between practitioners of chaos (SRE, DevOps engineers, Kubernetes developers). 

## Value-add to the CNCF ecosystem

*Does the project have a clear value add to the current project ecosystem? How does it relate to other projects with overlapping capabilities?*

Litmus adheres to the “Cloud Native” principles (as explained in this blog [here](https://www.cncf.io/blog/2019/11/06/cloud-native-chaos-engineering-enhancing-kubernetes-application-resiliency/)). In summary: 

- The life cycle management of Litmus happens using a chaos operator 
- Developer or SRE interacts with Litmus through declarative YAMLs. The chaos workflow can be completely automated in a highly scaled environment using GitOps
- Chaos-runner that manages each chaos job is in a container and can survive node reboots
- Any chaos logic which can run inside a container can be consumed through Litmus YAML files. So, anyone with a docker image that contains their chaos can make use of Litmus project to practice chaos engineering in their Kubernetes environment.


## Alignment with other CNCF projects

**Does the project align and actively collaborate with other CNCF projects?**

Litmus provides full-featured chaos experiments for most of the Kubernetes resources. Currently, there are eleven different generic ways to introduce and manage chaos on Kubernetes cluster. Apart from this through chaos hub, we bring application level chaos experiments for other eco system projects such as CoreDNS, Kafka, OpenEBS. 
As part of its near-term roadmap, Litmus also is in the process of creating chaos charts for other sandbox/incubating/graduated CNCF projects. We are working with cncf-ci workgroup to include chaos stage in the CNCF projects. A PR is already in review for CoreDNS project.

Litmus chaos experiments are being extensively used as part of CI pipelines of the OpenEBS, a CNCF sandbox project that provides containerized storage solution for Kubernetes. Reference: https://openebs.ci/


**Does the project require any specific versions of projects (or APIs) to interoperate? (e.g. K8s API, CSI, CNI, CRI)?**

LitmusChaos mainly uses the Kubernetes API and also makes use of the Docker, Containerd APIs in case of some experiments.   

**Does the project augment or benefit other CNCF projects?**

The project benefits other CNCF projects by helping harden their resiliency under various deployment conditions. Chaos engineering for CNCF projects becomes easy.  

## Anticipated use cases

### Staging & Pre-Prod: Random Chaos:

- Gamedays:  The chaos experiments are typically executed in deployment environments hosting several microservices (either staging clusters that mimic production environments or production itself) as part of “Gamedays”, where a hypothesis is tested out, with the results either confirming the hypothesis and leading to fixes either in the infrastructure or the application software. 
- Continuous Chaos: Chaos experiments (as a background service) are continuously and randomly executed against production-like environments with the appropriate visualization and monitoring/observability setup to gauge system resilience against different chained failures, occurring at random instances.  

### CI/CD:

- Organizations can add "Kubernetes generic" and “application specific” chaos experiments as part of a “chaos stage” in their CI pipelines thereby enabling a left shift in improving fault tolerance and failure response. 

### Kubernetes Upgrade Testing:

- Kubernetes upgrades can cause changes in behaviour / response to component failures which can be exposed by performing chaos experiments against them.

### Chaos Engineering in production:
- Litmus plays important role to orchestrate chaos in production using GitOps. 

## Alignment with SIG Reference Model

**Does the project align with the SIG CNCF reference model and which capabilities does it require/provide at each level of the reference model.**

## High level architecture

**Describe the overall architecture of the project. Feel free to add diagrams.**

**Litmus Project Architecture**
LitmusChaos project comes with the following components.
- Chaos Operator
- Chaos CRDs
- Chaos Metrics exporter and
- Chaos Hub (Which is a list of Chaos Customer Resources)

![Litmus Project Architecture diagrams](https://docs.litmuschaos.io/docs/assets/litmus-schematic.png)

Detailed architecture of the project is [here](https://docs.litmuschaos.io/docs/architecture/).

# Formal Requirements

**Document that the project fulfills the requirements as documented in the CNCF graduation criteria for sandbox**

- We are yet to approach the TOC sponsors at the time of submitting this issue. 
- This project has adopted the CNCF [Code of Conduct](https://github.com/litmuschaos/litmus/blob/master/CODE_OF_CONDUCT.md)
- Will adhere to the CNCF IP policy

**Are there any anticipated issues with any of the criteria ?**
No.

**Has the TOC been approached for sponsorship? If yes, which TOC members have agreed to sponsor the project?**
We have not approached any TOC members yet.

# CNCF IP Policy

**Becoming a sandbox project requires adoption of the CNCF IP Policy: https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy**

The source code developed for the LitmusChaos Project is licensed under Apache 2.0.

FOSSA report for the project is maintained [here](https://github.com/litmuschaos/litmus/blob/master/fossa.md).

**Note: there is a grace period after becoming a sandbox period to enable projects to adopt the policy, however, some prep is required to ensure there are no major blockers.**

**Has the IP policy been reviewed?**

**List the repos for the project and their current license**


| Repo Name | License |
| -------- | -------- | 
| [Litmus](https://github.com/litmuschaos/litmus)     |  Apache License Version 2.0     |     
| [Chaos Operator](https://github.com/litmuschaos/chaos-operator)     |  Apache License Version 2.0     |     
| [Chaos Exporter](https://github.com/litmuschaos/chaos-exporter)     |  Apache License Version 2.0     |     
| [Chaos Charts](https://github.com/litmuschaos/chaos-charts)     |  Apache License Version 2.0     |   
| [Documentation](https://github.com/litmuschaos/litmus-docs)     |  Apache License Version 2.0     |
| [Chaos Runner](https://github.com/litmuschaos/chaos-runner)     |  Apache License Version 2.0     |
| [Elves](https://github.com/litmuschaos/elves)     |  Apache License Version 2.0     |
| [Chaos Hub](https://github.com/litmuschaos/charthub.litmuschaos.io)     |  Apache License Version 2.0     |    
| [Test Tools](https://github.com/litmuschaos/test-tools)     |  Apache License Version 2.0     |    
| [E2E](https://github.com/litmuschaos/litmus-e2e)     |  Apache License Version 2.0     | 




**List any dependent repos (upstream/downstream) that are required to build the project (including but not limited to libraries, commercial tools, plugins)**
None

**What actions are required to be compliant with the IP policy?**
None.

# Other Considerations

## **Please note, these are not gating criteria but rather to:**

- **Collect a standard set of information for each project**
- **Provides a point in time capture of the state of the project which makes it easier to track progress at future reviews and / or promotion**
- **Help projects to prepare for SIG and TOC presentation**
- **Allow the SIG to review the project and perform due diligence for incubation**
- **Provide the TOC with the information required to accept sponsorship of a project and/or votes**
- **Identify and rectify any significant issues / blockers prior to presenting to the TOC and acceptance as a CNCF project**

## Cloud Native

**Does the project meet the definition of Cloud Native? The CNCF charter states:**

“Cloud native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds. Containers, service meshes, microservices, immutable infrastructure, and declarative APIs exemplify this approach.


“These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.”

Yes. This project meets the definition of Cloud Native.

## Project and Code Quality

**Are there any metrics around code quality? Are there good examples of code reviews? Are there enforced coding standards?**

- Some standards for the ansible-based experiments: https://github.com/litmuschaos/litmus/blob/master/.github/PULL_REQUEST_TEMPLATE.md
- Standards for go/python-based components (chaos-operator, chaos-runner) include: BCH, GoReportCard 
- CII best practice score of 82% is achieved. https://bestpractices.coreinfrastructure.org/en/projects/3202
- BCH score of 8/10 for litmuschaos/litmus and 10/10 for litmuschaos/loperator and litmuschaos/lexporter repositories.

Examples of review: 
  - https://github.com/litmuschaos/chaos-operator/pull/193
  - https://github.com/litmuschaos/litmus/pull/1015
  - https://github.com/litmuschaos/chaos-runner/pull/33
  - https://github.com/litmuschaos/litmus/pull/1345
  - https://github.com/litmuschaos/litmus-e2e/pull/50

**What are the performance goals and results? What performance tradeoffs have been made? What is the resource cost?**

There are no performance tradeoffs. Litmus adopts a Kubernetes native approach to chaos with custom resources and controllers, 
thereby reusing the Kubernetes infrastructure itself, hence causing no additional resource costs. 

**What is the CI/CD system? Are there code coverage metrics? What types of tests exist?**

Apart from unit-tests, Litmus makes use of GinkGo based BDD framework to create e2e suites for testing the 
chaos framework (chaos experiments, operator, runner, exporter). These are executed via Gitlab pipelines at 
scheduled intervals, with Travis/CircleCI being used for “build” purposes. The build pipelines themselves include 
static code checks (lint) as well as security scans against the docker image artefacts before being pushed into the 
respective public image repositories on dockerhub. 

Code-Coverage has been added to the Chaos-Operator and the community is working on increasing the same via more 
unit-tests. 

**Is there documentation?**

Litmus has documentation explaining both user as well as developer workflows at the [litmus-docs](https://docs.litmuschaos.io/docs/getstarted/) along with individual experiment guides.

**How is it deployed?**

The Litmus chaos operator is installed as a K8s deployment, along with the custom resource definitions (CRDs) for the ChaosEngine, ChaosExperiment and ChaosResult resources. The out-of-the-box chaos experiments are installed as CustomResource YAML manifests. 

**How is it orchestrated?**

Orchestrated by Kubernetes natively. Litmus has a chaos-runner that runs in a container, an operator that is built using the operator-sdk and chaos experiments which are defined as Kubernetes custom resources.


**How will the project benefit from acceptance into the CNCF?**

- This project will have a vendor neutral home. The project will generate interest in many CNCF project members to contribute application level chaos experiments to the chaos hub when it is accepted into CNCF.
- CNCF projects themselves may adopt Litmus more actively when accepted into CNCF.

**Has a security assessment by the security SIG been done? If not, what is the status/progress of the assessment?**

No. We have to approach the security SIG.

# Promotion to Incubation

## Open Governance

**How are committers chosen?**

- Committers are chosen once they make significant contributions (such as additions of new chaos charts/categories, 
litmus backend/chaos orchestration component improvements, or litmus project infra improvements such as docusaurus-based
website, chaoshub) and are committed to the long-term success of the project. 

**How are architectural and roadmap decisions made?**

- Roadmap & architectural decisions are made via consultation and feedback from the Litmus community (users and contributors) 
comprising developers and SREs. The feedback is received via github issues or during the monthly community sync up call.
- ROADMAP items are reviewed by maintainers before acceptance. 

**How many decision makers are outside the sponsoring organization.**

- *Multiple (Independent contributors, Intuit, Wipro) apart from the primary sponsor MayaData.* 

## Adoption

**Who are the current maintainers?**

- "Chandan Kumar",@chandankumar4,MayaData
- "Jayesh Kumar",@k8s-dev,Self
- "Karthik Satchitanand",@ksatchit,MayaData
- "Sumit Nagal",@sumitnagal,Intuit
- "Uma Mukkara",@umamukkara,MayaData

The adopters file is managed [here](https://github.com/litmuschaos/litmus/blob/master/ADOPTERS.md).

**How long has the project been developed for?**

Litmus started in May 2018, Litmuschaos org created in Apr 2019

**Is there a commercial version of the project or a primary commercial sponsor ?**

- MayaData provides support to Litmus components through it's [OpenEBS Enterprise Platform product](https://mayadata.io/product).
- Wipro includes support to Litmus in it's [AppAnywhere managed service](https://www.wipro.com/en-IN/infrastructure/wipros-appanywhere/).

**Is the project used in production? If so, please list some of the accounts.**

*MayaData* 

**Does the project participate in a CNCF User Group?**

No.

## Vendor Independence

**Is the project reasonably independent from the sponsoring vendor?**

Yes

**Are all communication channels and project resources hosted just for this project or with other CNCF projects/resources?**

We use #litmus channel in Kubernetes Slack. 

**Is all code that is part of the project hosted and part of the CNCF managed orgs and repos?**

Yes. All the code will be under CNCF managed orgs and repos. 

**Are all defaults for upstream reporting either unset or community hosted infrastructure (i.e. doesn’t point to vendor hosted SaaS control plane or analytics server for usage data)? Is all project naming independent of vendors?**

**Relevant Assets regarding vendor independence**

https://litmuschaos.io

https://docs.litmuschaos.io

https://hub.litmuschaos.io

https://github.com/litmuschaos
https://twitter.com/litmuschaos
