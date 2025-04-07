## Table of Contents

- [Background](#background)
- [Recap: What is LitmusChaos](#recap-what-is-litmuschaos)
- [Progress Since Joining CNCF](#progress-since-joining-cncf)
- [Incubation Stage Requirements](#incubation-stage-requirements)
- [Statement on alignment with the CNCF mission](#statement-on-alignment-with-the-cncf-mission)
- [Comparison with other CNCF chaos engineering projects](#comparison-with-other-cncf-chaos-engineering-projects)
- [Roadmap](#roadmap)

## Background

- [CNCF blog on defining cloud native chaos engineering - 6th Nov 2019 ](https://www.cncf.io/blog/2019/11/06/cloud-native-chaos-engineering-enhancing-kubernetes-application-resiliency/)
- [Sandbox presentation slides](https://docs.google.com/presentation/d/1w94WgGmsgRXu4xFbkkVIPHtEl5V4eOg6FxgVikGimYs/edit#slide=id.g7e5d7aa8b9_2_75) & [Sandbox presentation video - 18th March 2020](https://youtu.be/O40GK24Ea4k?t=1106)
- [Recommendations from SIG-App-Delivery upon acceptance as sandbox project - June 25th 2020](https://docs.google.com/document/d/10G1fOySiYeCYIkWhz_ueyXTBi04_pQVRkDBdt2ryCFo/edit#heading=h.6pidoniplc5n) 
- [CNCF blog on Introduction to LitmusChaos - 28 Aug 2020](https://www.cncf.io/blog/2020/08/28/introduction-to-litmuschaos/)
- [Incubation presentation to Sig-App-Delivery - 3rd March 2021 ](https://docs.google.com/presentation/d/1jir2aHrfkXUsOxhObyjxgNcWy0eZ2XPlaFhDD8i9FWc/edit?usp=sharing)

## Recap: What is LitmusChaos

### What is LitmusChaos ?

LitmusChaos is a chaos engineering framework for Kubernetes. The  project aims to provide a complete Chaos Engineering platform to help developers and SREs to carry out chaos experiments easily and as defined by the [principles of chaos](https://principlesofchaos.org/) . It has “cloud-native chaos” as its fundamental design goal and keeps the community at the center for growing the chaos experiments via the ChaosHub. Litmus supports common practices and aids needed in the practice of chaos, including blast radius control, steady-state definition & hypothesis validation, simulation of real-world scenarios, ability to schedule/automate experiments and observability hooks via events and prometheus metrics.

At a high level, Litmus comprises of:

Chaos Control Plane: A centralized chaos management tool called chaos-center, which helps construct, schedule and visualize chaos. The control plane is made up of a graphql server, react-based web-ui/dashboard and mongodb to store chaos data.

Chaos Execution Plane: Made up of a chaos agent or subscriber, multiple operators and a chaos exporter that execute & monitor the experiment within a defined target Kubernetes environment.

ChaosHub: An open-catalog of ready-to-use chaos/fault templates for Kubernetes and cloud platforms. The public chaoshub maintained by the Litmus project is also “embedded” into the chaos-center to aid in picking the right faults as part of a chaos workflow.

### What are the use cases?

LitmusChaos is predominantly used by the following persona:

-   Developers: To run chaos experiments during application development as part of extended devtests.
    
-   DevOps Engineers: To run chaos as a pipeline stage to find bugs when the application is subjected to fail paths as part of e2e validation in CI/CD
    
-   SREs: To plan and schedule chaos experiments into applications and/or supporting infrastructure in staging/pre-prod environments during gamedays or as part of a randomized / continuous background chaos service.

## Progress Since Joining CNCF

### Project improvements in this period 

#### Feature Summary:

- Migration to golang for chaos business logic with "litmus-sdk" created to bootstrap new experiments 
- Introduction of Chaos-Center, a centralized control plane to simplify management of chaos 
- Support for complex chaos workflows via Argo integration 
- Steady-state hypothesis validation via Litmus Probes
- Different modes (namespaced, admin/cluster)  of operation
- Ability to source chaos artifacts from a Git backend & synchronize changes into the chaos-center
- Support for chaos on containerd, cri-o runtime 
- Introduction of newer chaos types for both Kubernetes & cloud infrastructure (AWS,GCP,Azure,VMware)
- Improved suite of prometheus metrics for chaos experiments 

#### Governance Summary: 

- Moving to a SIG based model with leads driving improvements/agenda in the associated sub-projects (group of repositories) under the purview of the respective SIG

- Operational groups include sig-deployment (Leads from Wrike, Klanik), sig-docs (Leads from HSBC, Amazon) & sig-observability (Leads from NetApp, Deutche Bank). 


### Community Stats (graphs / dashboards) 

(Refer to the [DD doc](https://docs.google.com/document/d/1_8svqwzs1xYRHU92Qfd1I6UUXYyb8J3Zlpemg2-n4SU/edit#heading=h.kd4eg2uz3lt0) for detailed stats) 

- Stars
  - 2.3k from 650 

- Slack Membership
  - 1025 from 237 

- Usage metrics 
    - Operator installations: ~240k  
    - Experiments Run: ~570k 

- New Experiments 
    - 51 from 29 (22) 

### External contributions

List of contributing companies: https://litmuschaos.devstats.cncf.io/d/5/companies-table?orgId=1

### Integrations with CNCF Ecosystem

- [Argo Workflows](https://dev.to/ksatchit/chaos-workflows-with-argo-and-litmuschaos-2po5) (constructing complex chaos workflows) & Argo Rollouts (select pods managed by argo rollouts annotated for chaos)
- [Keptn](https://github.com/keptn-sandbox/litmus-service) (executing chaos tests as part of the keptn CD pipeline before quality gate evaluation)
- [Spinnaker](https://www.armory.io/blog/litmuschaos-in-your-spinnaker-pipeline/) (build chaos stages as part of the spinnaker CD pipeline before app promotion)
- [Gitlab](https://dev.to/uditgaurav/litmuschaos-gitlab-remote-templates-6l2) (chaos tests as a gitlab job )
- [Okteto](https://dev.to/ksatchit/litmus-sdk-devtest-your-chaos-experiments-with-okteto-4dkj) (Okteto CLI for development of litmus experiments) 
- [Prometheus & Grafana](https://dev.to/ksatchit/observability-considerations-in-chaos-the-metrics-story-6cb) monitoring chaos impact on applications) 
- [Dex](https://docs.litmuschaos.io/docs/next/user-guides/chaoscenter-oauth-dex-installation) (integrated with litmus auth server to provide improved authentication methods, such as via Google, Github) 
- [Kyverno](https://litmuschaos.github.io/litmus/experiments/concepts/security/kyverno-policies/) (building policies to control permissions for execution of litmus experiments)


### Catalog listing / Open Marketplace / SaaS Integrations 
- [Okteto Cloud](https://github.com/okteto/litmus-on-okteto)
- [Kublr](https://kublr.com/on-demand-videos/kubernetes-chaos-engineering-with-mayadata-and-kublr/) 

#### Artifacts availability 
- Operatorhub
- Helm Hub / Artifact Hub
- jFrog ChartCenter

### References / Highlights: 

- Meetup Presentations by User Community Members
  -  [Argo meetup (Sumit Nagal)](https://youtu.be/eqm8bhrFuiY?t=1453), [Okteto Meetup (Ramiro Berrelleza)](https://youtu.be/MVWmRD9B470), [WeScale Meetup (Akram Riahi) (French)](https://youtu.be/tr5C8bZ0elc), [Be Tech! Meetup (Santander) (Spanish) ](https://www.youtube.com/watch?v=8KyjsRhhRI0) 
- [KubeCon case-study featuring LitmusChaos](https://static.sched.com/hosted_files/kccncna2021/be/Reslience-telco-KubeCon_NA_2021.pptx.pdf)
- [Thoughtworks Radar](https://www.thoughtworks.com/radar/tools?blipid=202010081#.X7H-CwT19VA.link) (published. Oct 20, 2020) 
- [User Testimonial](https://www.youtube.com/watch?v=PiGqQc_UYPc&amp;feature=youtu.be) (Andreas Krivas, Engineering Manager, ContainerSolutions AG) 

#### Litmus user presentations at ChaosCarnival
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
- Chaos GitOps for highly scalable automation of chaos workflows. Chaos can now be triggered as a result of a change to an application. This integrates with other CD tools like ArgoCD and FluxCD
- Chaos Interleaved dashboards. A step toward open observability that is interleaved with chaos incident details

With all these features, Litmus is a comprehensive platform for chaos engineering.

**Note:** Litmus itself is composed of microservices. For 2.0, more microservices are added and the existing ones will continue to work. Litmus 2.0 is completely backward compatible. No features are deprecated. Migration path is about constructing new artifacts such as ChaosWorkflows that include the current chaos experiments in use by the users. 

## Incubation Stage Requirements

### Production usage

(Refer to the detailed adoption info [here](https://github.com/litmuschaos/litmus/blob/master/ADOPTERS.md))

  - [HaloDoc](https://www.halodoc.com/): Uses the chaos-center to control and manage complex chaos workflows across its fleet of pre-prod and production clusters. Also uses the chaos-gitops & event-triggered chaos functionality wherein microservices subscribe to or are mapped to a given workflow that is triggered upon any changes to the said microservice on the cluster. 

- [PoleEmploi](https://www.pole-emploi.fr/accueil/): Uses LitmusChaos along with private chaoshubs to test the resiliency of the Kubernetes infrastructure as well as application microservices in an air-gapped environment. 

- [Intuit](https://www.intuit.com/?utm_source=github&utm_campaign=litmuschaos_repo): Uses the BYOC (Bring-Your-Own-Chaos) capability of Litmus, where existing experiments can be orchestrated with Litmus (in the process offering blast radius & app filter capabilities, on-demand/run-once execution, chaos artifacts that lend themselves to GitOps, Kubernetes chaos events  etc.,) 

- [Orange](https://www.orange.com/): Litmus is being used to test the resilience of OpenStack based Kubernetes environment. Litmus experiments are run against their internal cloud systems and also in some of their CI/CD pipelines.

- [Anuta Networks](https://www.anutanetworks.com/) : Anuta Networks uses Kubernetes for some applications that enable parts of their public cloud. Litmus is being used by their SREs to test the resilience of their Kubernetes platform and these applications.

- [Red Hat](https://www.redhat.com/en) : Litmus is being used to test the maturity and resilience of Red Hat Openshift Virtualization solution. The experiments involve chaos injection on the VMs that host the OpenShift nodes (reboots, shutdowns) as well as pod level chaos while using litmus probes to validate expected behavior during chaos.

- [Container Solutions](https://www.container-solutions.com/) : Container Solutions team has been using Litmus and contributing back to the project at various levels including Network Chaos experiments and Chaos orchestration improvements. They are using Litmus for chaos engineering needs of their client environments which are based on OpenShift for nearly an year.

- [Wipro](https://www.wipro.com/en-IN/infrastructure/wipros-appanywhere/?utm_source=github&utm_campaign=litmuschaos_repo): Uses litmus as part of its “AppAnywhere Gecko” managed service. to validate and verify the reliability of applications moved into Kubernetes as part of end-to-end automation. Also makes chaos experiments available as self-service for customers of the platform. 

- [WeScale](https://www.wescale.fr/): Uses Litmus experiments to enhance resiliency of clients’ Kubernetes clusters. The experiments are integrated with the respective clients CI pipelines, with pod-level chaos experiments attempted first before moving to cluster level ones.

- [NetApp](https://www.netapp.com) [(chaos as part of QA cycles)](https://github.com/litmuschaos/litmus/blob/master/adopters/organizations/netapp.md), which also contributes to the Litmus project, with one of the SIG-Chairs for Observability. 


#### Known list of LitmusChaos users
Litmus is known to be used by the following organizations which are at various stages of testing.

- [LifeCycle](https://www.lifecycle-software.com/)
- [Daimler AG](https://www.daimler.com/)
- [Auto Pilot](https://www.autopilothq.com/)
- [IAG Limited](https://www.iag.com.au/)
- [F5 Networks](https://f5.com)
- [Oracle](https://www.oracle.com)
- [Microsoft](https://www.microsoft.com)
- [Navi Technologies](https://navi.com/)
- [Akridata](https://www.akridata.com/)
- [Workday](https://www.workday.com/)
- [Argo AI](https://www.argo.ai/)
- [Nokia](https://nokia.com)
- [AWS](https://aws.amazon.com)
- [Aicumen Technologies](https://www.aicumen.com/)
- [ANZ](https://www.anz.com.au/)
- [QVC](https://www.qvc.com)

### Committers
- Have a healthy number of committers

  - [Maintainers list](https://github.com/litmuschaos/litmus/blob/master/MAINTAINERS)

### Demonstrate substantial flow of commits / healthy merge rate (backed-up by devstats)

- Release Cadence
  - [Sustained monthly release](https://litmuschaos.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-2y&to=now-1h) cadence (15th of every month) (this is measured against releases on litmuschaos/litmus)
  - Patch releases on demand (1-2) between monthly releases for individual components / sub-projects

- New Contributors added since sandbox: [120](https://litmuschaos.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1593023400000&to=now)

- [PRs opened/merged per month since sandbox](https://litmuschaos.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All&from=now-1y&to=now):  203/1197

- No. of contributors month-on-month over the last year
  - 6.75 per month (including hacktoberfest)
  - 4.16 per month (excluding hacktoberfest)

- No. of contributing orgs over the last couple of years
  - More than [10 companies made over 30 contributions](https://litmuschaos.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%202%20years&var-metric=contributions)

### A clear versioning Scheme

- Semver for Releases: 
  - We use the [semantic versioning](https://semver.org/#semantic-versioning-specification-semver) numbering system. 
  - The releases are announced using [github’s releases](https://github.com/litmuschaos/litmus/releases) and followed by slack and social media announcements.
  - Helm Versioning Practices also as per semver

- Minor & Patch Release 
  - Based on user issues/requests


## Statement on alignment with the CNCF mission

The practice of Chaos Engineering is a necessity for achieving service reliability. Reliable application delivery and operations are very important and key to the success of technology adoption of various CNCF projects and ecosystem projects. LitmusChaos project goal is to become the toolset to measure and achieve that reliability.  Litmus extends the existing experiences of both the cloud-native developers and SREs in adding chaos tests to either CI pipelines or to their operational infrastructure. Litmus integrates with other CNCF projects such as Argo Workflows, Argo rollouts, Prometheus and Keptn to enrich the experience of cloud native communities during their efforts to run their application services reliably. LitmusChaos Hub hosts the application specific experiments and CNCF project specific experiments will find their way into the hub. CoreDNS is one such example. 

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

- In other words, the focus of the project is not only around providing means to inject different chaos types/failures (for example, projects like chaos-mesh & chaos blade are more focused on providing chaos injection methodologies), but also practice chaos engineering in a wholesome way. 

### Comparisons by  community 

Here is some blogs comparing different chaos engineering tools for kubernetes: 

https://blog.container-solutions.com/comparing-chaos-engineering-tools
https://www.gremlin.com/community/tutorials/chaos-engineering-tools-comparison/

**Note**: The blog by Container-Solutions & Gremlin was created between the 1.7.0 - 1.10 releases of litmus. All the roadmap items/needed improvements described in the blog have since been addressed.

## Roadmap

[Link to the Roadmap](https://github.com/litmuschaos/litmus/blob/master/ROADMAP.md)

- Immediate-Term (in-progress)
  -   Additional Chaos on non-kubernetes targets (cloud & baremetal resources)
  -   Improved Observability via audit logs, alerts from the Chaos-Center  
  -   Simplified monitoring (chaos-metric instrumented app dashboards) within the chaos-center  
  -   GRPC Chaos, I/O & JVM chaos librariesc  
  -   Runtime validation of chaos resources via admission controllers  
  -   Simplified ChaosHub setup for air-gapped environments

- Near-Term (Backlog)
  - Improved SLO management via conditional probes, probe chaining
  -   Improved integrations into the CI/CD ecosystem (via chaos-center API)
  -   Enhanced suite of application-specific chaos libraries
  -   Support for cloudevents compliant chaos events

- Governance
  - Operationalize more special interest groups  within litmus around various aspects of chaos engineering & aid with reference implementations in respective areas
  - Collaborate with contributing organizations to come up with an improved graduation process for committers & maintainers
