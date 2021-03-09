## Table of Contents

- [Background](#background)
- [Recap: What is LitmusChaos](#recap-what-is-litmuschaos)
- [Progress Since Joining CNCF](#progress-since-joining-cncf)
- [Incubation Stage Requirements](#incubation-stage-requirements)
- [Statement on alignment with the CNCF mission](#statement-on-alignment-with-the-cncf-mission)
- [Comparison with other CNCF chaos engineering projects](#comparison-with-other-cncf-chaos-engineering-projects)
- [Roadmap](#roadmap)

## Background


- [Sandbox presentation slides](https://docs.google.com/presentation/d/1w94WgGmsgRXu4xFbkkVIPHtEl5V4eOg6FxgVikGimYs/edit#slide=id.g7e5d7aa8b9_2_75)
- [Recommendation from SIG-App-Delivery chairs](https://github.com/cncf/toc/pull/391#issuecomment-624115660) (refer to end of doc) 
  - [Sandbox presentation video](https://youtu.be/O40GK24Ea4k?t=1106)
- CNCF incubation presentation (TBA)
- [CNCF blog on Introduction to LitmusChaos - 28 Aug 2020](https://www.cncf.io/blog/2020/08/28/introduction-to-litmuschaos/)
- [CNCF blog on defining cloud native chaos engineering - 6 Nov 2019](https://www.cncf.io/blog/2019/11/06/cloud-native-chaos-engineering-enhancing-kubernetes-application-resiliency/)
- [Project update to Sig-App-Delivery](https://docs.google.com/presentation/d/1zNzSRc0N4I1_LsQsUnpcispQAJMU9ni6k-vfTQXefLo/edit?usp=sharing)
- [Incubation presentation to Sig-App-Delivery](https://docs.google.com/presentation/d/1jir2aHrfkXUsOxhObyjxgNcWy0eZ2XPlaFhDD8i9FWc/edit?usp=sharing)

## Recap: What is LitmusChaos

### What is LitmusChaos ?
Litmus is a chaos engineering framework for Kubernetes. It provides a complete set of tools required by Kubernetes developers and SREs to carry out chaos experiments easily and in Kubernetes-native way. The project has “Declarative Chaos” as the fundamental design goal and keeps the community at the center for growing the chaos experiments. 
Litmus has the following components:
- Chaos Portal: Litmus Portal is a centralized web portal for creating, scheduling, and monitoring chaos workflows. A chaos workflow is a set of chaos experiments. Chaos workflows can be scheduled on remote Kubernetes clusters from the portal. SRE teams can share the portal while managing chaos through the portal.
- Chaos Operator: This operator is built using the Operator SDK framework and manages the lifecycle of a chaos experiment.
- Chaos CRDs: Primarily, there are three chaos CRDs – ChaosEngine, ChaosExperiment, and ChaosResult. Chaos is built, run, and managed using the above CRDs. ChaosEngine binds a target application (or infrastructure component) with a ChaosExperiment CR. When run by the operator, the result will be stored in ChaosResult CR.
- Chaos experiments or the ChaosHub: Chaos experiments are custom resources that define low level chaos intent. The YAML specifications for these custom resources are hosted at the public ChaosHub (https://hub.litmuschaos.io).

### What are the use cases?

Typical use cases of Litmus include: chaos testing CI/CD pipelines, in staging and production environments with a goal of either benchmarking the resilience or improving the resilience

*Link to adopters*

https://github.com/litmuschaos/litmus/blob/master/ADOPTERS.md


## Progress Since Joining CNCF

### Project improvements in this period (features & governance) 

#### Features:

- Litmus Portal to simplify management of chaos 
- Support for chaos on containerd, crio runtime 
- Golang SDK & migration to golang-based experiments 
- Complex chaos workflows with argo integration 
- Declarative Hypothesis via Litmus Probes
- Different modes (namespaced, admin/cluster)  of operation
- GitOps integration - Both for keeping chaos configuration at Git and to achieve event triggered chaos
- Integration with other tools such as Argo Workflow, Keptn, GitLab, GitHub and Spinnaker
- Interleaved chaos charts for open observability

#### Governance:

- Moved to a SIG based model with chairs maintaining the associated sub-projects (group of repositories) under the purview of the respective SIG, while the project leadership group comprising a subsection of the chairs holds commit bit across repos in the project. 

- Chairs from Wrike (sig-deployment), Infracloud (sig-orchestration), Amazon & HSBC (sig-docs), Intuit (sig-integration), NetApp & Deutche Bank(sig-observability)


### Community Stats (graphs / dashboards) 

- Stars
  - 1.5k from 650 (2x)

- Slack Membership
  - 555 from 237 (~134%+)

- Usage metrics 
  - Operator Installations / Experiment Run Count
    - Operator installations: 41667 
    - Experiments Run: 215294

  - New Experiments Added
    - 48 from 29 (17) 

### External contributions

- Intuit 
- RedHat
- Microsoft
- Dino Systems
- Navi Technologies
- IAG
- Autopilot
- Container Solutions
- Okteto 
- HSBC
- Deutsche Telecom
- Orange

#### References / Highlights: 
- [Special Interest Groups](https://github.com/litmuschaos/litmus/wiki/Special-Interest-Groups)
- Meetups by Community, Joint Meetups 
  - Our sync up, [Argo meetup (Sumit)](https://youtu.be/eqm8bhrFuiY?t=1453), [okteto](https://youtu.be/MVWmRD9B470), [WeScale(French)](https://youtu.be/tr5C8bZ0elc), [Kubernetes/Docker meetup (malaysia)](https://youtu.be/bXwhXpw74No), [Cloud Native, Madison](https://www.youtube.com/watch?v=wRLKiMXbBbo&t=2620s)
- [Thoughtworks Radar](https://www.thoughtworks.com/radar/tools?blipid=202010081#.X7H-CwT19VA.link) (published. Oct 20, 2020) 
- [User Feedback](https://www.youtube.com/watch?v=PiGqQc_UYPc&amp;feature=youtu.be) (Andreas Krivas, Engineering Manager, ContainerSolutions AG) 

#### Integrations with CNCF Ecosystem
- [Argo Workflows](https://dev.to/ksatchit/chaos-workflows-with-argo-and-litmuschaos-2po5) (constructing complex chaos workflows) 
- Argo Rollouts (select pods managed by argo rollouts annotated for chaos)
- [Keptn](https://github.com/keptn-sandbox/litmus-service) (executing chaos tests as part of the keptn CD pipeline before quality gate evaluation)
- [Spinnaker](https://www.armory.io/blog/litmuschaos-in-your-spinnaker-pipeline/) (build chaos stages as part of the spinnaker CD pipeline before app promotion)
- [Gitlab](https://dev.to/uditgaurav/litmuschaos-gitlab-remote-templates-6l2) (chaos tests as a gitlab job )
- [Okteto](https://dev.to/ksatchit/litmus-sdk-devtest-your-chaos-experiments-with-okteto-4dkj) (chaos operator is available as a ready template usage of Okteto CLI for development of litmus experiments) 
- [Prometheus & Grafana](https://grafana.com/grafana/dashboards/12096) (chaos event [interleaved](https://dev.to/ksatchit/monitoring-litmus-chaos-experiments-198a) app dashboards) 

#### Catalog listing / Open Marketplace / SaaS Integrations 
- [Okteto Cloud](https://github.com/okteto/litmus-on-okteto)
- [Kublr](https://kublr.com/on-demand-videos/kubernetes-chaos-engineering-with-mayadata-and-kublr/) 

#### Artifacts availability 
- Operatorhub
- Helm Hub / Artifact Hub
- jFrog ChartCenter

### Litmus user presentations at ChaosCarnival
- [Network Chaos (Andreas Krivas, ContainerSolutions)](https://www.youtube.com/watch?v=whdDP41Omd8&list=PLBuYBMjBLBzHPuPsvdbJvKu1KxSowWDYl&index=3)
- [Chaos Testing on OpenShift (Jordi Gill, Red Hat)](https://www.youtube.com/watch?v=VITGHJ47gx8&list=PLBuYBMjBLBzHPuPsvdbJvKu1KxSowWDYl&index=6)
- [Putting Chaos into Continuous Delivery (Juergen Etzlstorfer, Keptn](https://www.youtube.com/watch?v=j7THMrIGKO4&list=PLBuYBMjBLBzGGxnZn7DVMQ2tbh7Gwkw7G&index=8&t=1277s)
- [Building Reliability Pipelines with Chaos Workflows (Michael Knyazev, IAG)](https://www.youtube.com/watch?v=7yBFgqUo01E&list=PLBuYBMjBLBzHPuPsvdbJvKu1KxSowWDYl&index=14)
- [Event-Driven Chaos Injection  (Raj and Shoumo, ChaosNative](https://www.youtube.com/watch?v=7cF3rwcZMcA&list=PLBuYBMjBLBzHPuPsvdbJvKu1KxSowWDYl&index=8)
- [Bring Chaos into your Development Environment (Ramiro Berrelleza, Oketeto)](https://www.youtube.com/watch?v=3No4-1H6slc&list=PLBuYBMjBLBzGGxnZn7DVMQ2tbh7Gwkw7G&index=13&t=1830s)
- [Chaos Engineering in Telco Cloudnative Infra (Samar and Vaibhav, Orange)](https://www.youtube.com/watch?v=UOhjFbCrncw&list=PLBuYBMjBLBzHPuPsvdbJvKu1KxSowWDYl&index=4&t=201s )

### Litmus 2.0
When we started the Litmus project, the goal of this project was to create a complete platform to practice chaos engineering at scale in a Kubernetes way. Of course this had to be done incrementally, first create a toolset for chaos injection and then add additional features to make it a platform. Litmus 1.x achieved the goal of keeping it completely open source, creating a ChaosHub and creating the required CRDs, Operators and Schedulers. With Litmus 1.x, users have a working chaos engineering toolset aligned with the original goals.
Over time, the monthly cadence releases added the following features. 
- Chaos experiments become building blocks of a ChaosWorkflow, to allow users to create a larger chaos scenarios. 
- A portal to centrally visualize the chaos workflows, get chaos analytics, get the teaming in place for collaboration of chaos workflows.
- Chaos GitOps for highly scalable automation of chaos workflows. Chaos can now be trigged as a result of a change to an application. This integrates with other CD tools like ArgoCD and FluxCD
- Chaos Interleaved dashboards. A step toward open observability that is interleaved with chaos incident details

With all these features, Litmus is a comprehensive platform for chaos engineering.

**Note:** Litmus itself is composed of microservices. For 2.0, more microservices are added and the the existing ones will continue to work. Litmus 2.0 is completely backward compatible. No features are deprecated. Migration path is about constructing new artifacts such as ChaosWorkflows that include the current chaos experiments in use by the users. 

## Incubation Stage Requirements

### Production usage

  - [Intuit](https://www.intuit.com/?utm_source=github&utm_campaign=litmuschaos_repo): Uses the BYOC (Bring-Your-Own-Chaos) capability of Litmus, where existing experiments can be orchestrated with Litmus (in the process offering blast radius & app filter capabilities, on-demand/run-once execution, chaos artifacts that lend themselves to GitOps, Kubernetes chaos events  etc.,) 
			
		References: 
    - https://github.com/litmuschaos/litmus/blob/master/adopters/organizations/intuit.md
    - https://sumitnagal.medium.com/chaos-journey-279924051d57
    - https://youtu.be/cvH6R4wRVrg?t=1055


  - [Wipro](https://www.wipro.com/en-IN/infrastructure/wipros-appanywhere/?utm_source=github&utm_campaign=litmuschaos_repo): Uses litmus as part of its “AppAnywhere Gecko” managed service. to validate and verify the reliability of applications moved into Kubernetes as part of end-to-end automation. Also makes chaos experiments available as self-service for customers of the platform. 

    References:

    - https://github.com/litmuschaos/litmus/blob/master/adopters/organizations/wipro.md

  - [WeScale](https://www.wescale.fr/): Uses Litmus experiments to enhance resiliency of clients’ Kubernetes clusters. The experiments are integrated with the respective clients CI pipelines, with pod-level chaos experiments attempted first before moving to cluster level ones.
		
		References: 
    - https://github.com/litmuschaos/litmus/blob/master/adopters/organizations/wescale.md
    - https://blog.wescale.fr/2020/09/29/wespeakcloud-quest-ce-que-le-chaos-engineering-et-comment-peut-il-sappliquer-a-votre-infrastructure/


  - [NetApp](https://www.netapp.com) [(chaos as part of QA cycles)](https://github.com/litmuschaos/litmus/blob/master/adopters/organizations/netapp.md), which also contributes to the Litmus project, with one of the SIG-Chairs for Observability. 


#### Known list of LitmusChaos users
Litmus is known to be used by the following organizations which are at various stages of testing.

- [ContainerSolutions](https://www.container-solutions.com/): Uses Litmus experiments and infrastructure to manage chaos engineering in a large Kubernetes cluster at one of their clients. They have been using Litmus for more than 6 months and contributed to various chaos experiments based on their field usage experience. 
  - https://youtu.be/PiGqQc_UYPc
- [Daimler AG](https://www.daimler.com/)
- [Auto Pilot](https://www.autopilothq.com/)
- [IAG Limited](https://www.iag.com.au/)
- [F5 Networks](https://f5.com)
- [Oracle](https://www.oracle.com)
- [Microsoft](https://www.microsoft.com)
- [Navi Technologies](https://navi.com/)
- [RedHat](https://www.redhat.com)
- [Akridata](https://www.akridata.com/)
- [Workday](https://www.workday.com/)
- [Anuta Networks](https://www.anutanetworks.com/)
- [Argo AI](https://www.argo.ai/)
- [Nokia](https://nokia.com)
- [AWS](https://aws.amazon.com)
- [Aicumen Technologies](https://www.aicumen.com/)
- [ANZ](https://www.anz.com.au/)
- [QVC](https://www.qvc.com)

### Committers
- Have a healthy number of committers

  - Maintainers list
    - Overall project maintainers: 
      - [Sumit Nagal (Intuit)](https://github.com/sumitnagal)
      - [Jayesh Kumar (AWS)](https://github.com/k8s-dev) 
      - [Karthik Satchitanand (MayaData)](https://github.com/ksatchit)
      - [Uma Mukkara (MayaData)](https://github.com/umamukkara)
      - [Maria Kotlyarevskaya (Wrike)](https://github.com/Jasstkn)

  - Committers across Organizations (per SIG / sub-projects)

    - SIG Integrations: [litmus-python, chaos-charts, chaos-workflows, litmus-go, litmus-ansible]
      - [Sumit Nagal (Intuit)](https://github.com/sumitnagal)
      - [Ondra Machacek (RedHat)](https://github.com/machacekondra)
      - [Ramiro Berrelleza (Okteto)](https://github.com/rberrelleza)

    - SIG Documentation: [litmus-docs] 
      - [Jayesh Kumar (AWS)](https://github.com/k8s-dev)
      - [Divya Mohan (HSBC)](https://github.com/divya-mohan0209)

    - SIG-Orchestration [chaos-operator, elves, chaos-runner, admission-controller, chaos-scheduler]
      - [Rahul M Chheda (InfraCloud)](https://github.com/rahulchheda)

    - SIG-Deployment [chaos-helm, litmus]
      - [Maria Kotlyarevskaya (Wrike)](https://github.com/Jasstkn)

    - SIG Observability [chaos-observability] 
      - [Kristin Barkardottir (NetApp)](https://github.com/xkbarkar), 
      - [Shantanu Deshpande (Deutsche Bank)](https://github.com/ishantanu)

### Demonstrate substantial flow of commits / healthy merge rate (backed-up by devstats)

- Release Cadence
  - [Sustained monthly release](https://litmuschaos.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-1y&to=now-1h) cadence (15th of every month) (this is measured against releases on litmuschaos/litmus)
  - Patch releases on demand (1-2) between monthly releases for individual components / sub-projects

- New Contributors added in last 6 months: 50  
  - (81 since 12 months) [starting 25th Nov 2019](https://litmuschaos.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)

- [PRs opened/merged per month](https://litmuschaos.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All&from=now-1y&to=now): ~290/~280 

- No. of contributors month-on-month over the last year
  - 6.75 per month (including hacktoberfest)
  - 4.16 per month (excluding hacktoberfest)

- No. of contributing orgs over the last 1 year
  - More than [10 companies made over 5 contributions](https://litmuschaos.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions)

### A clear versioning Scheme

- Semver for Releases: 
  - We use the [semantic versioning](https://semver.org/#semantic-versioning-specification-semver) numbering system. 
  - The releases are announced using [github’s releases](https://github.com/litmuschaos/litmus/releases) and followed by slack and social media announcements.
  - Helm Versioning Practices also as per semver

- Minor & Patch Release 
  - Based on user issues/requests


## Statement on alignment with the CNCF mission

The practices of Chaos Engineering is a necessity for achieving service reliability. Reliable application delivery and operations are very important and key to the success of technology adoption of various CNCF projects and ecosystem projects. LitmusChaos project goal is to become the toolset to measure and achieve that reliability.  Litmus extends the existing experiences of both the cloud-native developers and SREs in adding chaos tests to either CI pipelines or to their operational infrastructure. Litmus integrates with other CNCF projects such as Argo Workflows, Argo rollouts, Prometheus and Keptn to enrich the experience of cloud native communities during their efforts to run their application services reliably. LitmusChaos Hub hosts the application specific experiments and CNCF project specific experiments will find their way into the hub. CoreDNS is one such example. 

In summary, LitmusChaos project clearly aligns with CNCF mission of making the cloud native ubiquitous by being the readily usable and easily extensible toolset for measuring and enhancing the reliability of Cloud Native applications and platforms.



## Comparison with other CNCF chaos engineering projects

- Litmus is an end to end chaos framework which allows the users to practice chaos engineering as per the [principles of chaos](https://principlesofchaos.org/) by allowing them to:
  - Minimize blast radius: 
    - Through annotation checks at higher level controllers like deployments/statefulsets/daemonsets/argo rollouts) 
    - Add affect percentage of pods belonging for these controllers

  - Define Hypothesis around experiments 
    - Via Litmus probes, to define diverse steady-state conditions & SLOs (ranging from downstream service availability, prometheus metrics, kubernetes resource status etc.,) 
    - Detailed chaosresults which indicate state of experiments & verdicts based on predefined SLO

  - Simulate real world failures & Observe events
    - By providing a wide range of pod-level and node-level chaos experiments
    - Through creation of complex chaos workflows that combine multiple faults via Litmus Portal 
    - Generating chaos events for each stage of experiment execution that can be viewed as prometheus metrics

  - Aiding continuous & automated execution 
    - Through the scheduled workflows with a rich scheduling policy 

- It achieves this while providing a kubernetes native user-experience to the developers & SREs by ensuring that both the chaos intent (experiment details like fault etc.,) as well as run properties (pod scheduling policies, resource requests & limits,  experiment image properties, monitoring, cleanup) etc., are defined via Custom Resources. 

- It also satisfies the need of multi-tenant architectures & addresses security needs by providing different modes (namespaced & admin)  of operation and also by specifying a per-experiment RBAC that interested users can leverage. 

- Importantly, Litmus orchestrates both native experiments as well as any other, i.e., supports BYOC and helps to homogenize the operational flow for these.

- In other words, the focus of the project is not only around providing means to inject different chaos types/failures (for example, projects like chaos-mesh & pumba are more focused on providing chaos injection methodologies), but also practice chaos engineering in a wholesome way. 

### Comparison blog by a community user
Here is a blog comparing different chaos engineering tools for kubernetes : https://blog.container-solutions.com/comparing-chaos-engineering-tools


## Roadmap

[Link to the Roadmap](https://github.com/litmuschaos/litmus/blob/master/ROADMAP.md)

Roadmap classification based on categories: 

- Features
  - Increased support for different container runtimes wrt I/O chaos
  - Integration with service-mesh ecosystem to provide experiments with greater control/granularity wrt service-level chaos (http, grpc)
  - Additional application specific chaos experiments
  - Support for platform chaos scenarios as applicable

- Integrations
  - Increased off-the-shelf chaos-interleaved monitoring dashboards for different popular applications
  - Continue to work with Keptn, Spinnaker, Okteto, Argo & other communities to drive the left-shift of chaos into CI/CD pipelines
- Governance
  - Operationalize the SIG-Testing & SIG-CI groups to improve the resiliency of the LitmusChaos framework & aid with reference implementations of chaos-enabled pipelines for community
  - Collaborate with contributing organizations to come up with an improved graduation process for committers & maintainers 
