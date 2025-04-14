# General Technical Review questions
v1.0
## Introduction

The General Technical Review questions can be completed by a project in lieu of a presentation to a Technical Advisory Group (TAG) as well as to satisfy the Engineering Principle requirements of a Sandbox application or Due Diligence for moving levels. 

The questions are designed to prompt **_design thinking_** for projects that would like to one day be a graduated project.

The intent is to understand how the project has adopted and aligned with the CNCF maturation levels as well as encourage good design & security best practices. 


<!-- 
The General Technical Review questions can be completed by a project team in lieu of a presentation to a Technical Advisory Group (TAG) as well as to satisfy several of the Engineering Principle requirements for applying to CNCF Sandbox as well as applying to move to Incubation and Graduation.

For the purposes of the general technical review and further domain reviews, the questions are designed to prompt design thinking for ready-for-production projects and architecture. The intent is to understand and validate the cloud native software lifecycle of the project and whether it is in line with the CNCF Maturation process and levels. 

Project maintainers are expected to have designed the project for cloud native use cases and workloads, as well as taking a ‘secure by design and secure by default’ approach that enables adopters and consumers of the project the ability to ‘loosen’ the defaults in a manner that suits their environment, requirements and risk tolerance.

These questions are to gather knowledge about the project. Project maintainers are expected to answer to the best of their ability. **_Not every question will be addressable by every project._**

**Suggestion:** A recorded demo or diagram(s) may be easier to convey some of the concepts in the questions below. The project maintainers may provide a link to a recorded demo or add architectural diagrams along with your GTR questionnaire.

-->

### General Technical Review questions 

The questions follow the cloud native software lifecycle day schemas:

**Day 0 - Planning Phase. (Sandbox)** - This phase covers design and architecture of the cloud native project.

**Day 1 - Installation and Deployment Phase (Incubation)** - This phase covers initial installation and deployment of the design developed during Day 0 - Planning Phase. 

**Day 2 - Day-to-Day Operations Phase (Graduated)** - This phase covers post-deployment operations in production-ready environments to include monitoring, maintenance, auditing and troubleshooting.


### How to use this template

Make a copy of the template below and answer questions related to your project level to the best of your ability.
_**Not every question will be addressable or relevant to every project.**_
If this is the case for your project, please mark it as not-applicable (N/A) and provide a brief explanation. 

**NOTE:** The questions are cumulative e.g. if you are applying for incubation or graduation, you should answer both day 0 and day 1 questions etc.

#### Tips

* Treat the GTR questionnaire as a living document and keep a copy of it in your project's own repo. The GTR questions are helpful to both contributors and users and will make updating it in the future less work when you want to apply to move levels.
* Answer more questions than the requirement for your level if it _makes sense for your project_. e.g. if you have documentation covering the different forms of observability in the Day-2 requirements.
* You **CAN** link out to your own project's documentation, but be sure to link to it in a _versioned_ form. e.g. link to it at a specific commit instead of the `main` branch, or versioned website.
* A recorded demo or diagram(s) may be easier to convey some of the concepts in the questions below. You may provide a link to a recorded demo or add architectural diagrams along with your GTR questionnaire.
* If you are unsure or have a question about any section below, **please ask**. Chances are you're not the only one with a question and the template should be updated with additional guidance.

---

# General Technical Review - [Project Name] / [Level]

- **Project:**
- **Project Version:** 
- **Website:**
- **Date Updated:** YYYY-MM-DD
- **Template Version:** v1.0
- **Description:** <!-- Short project description --> 



## Day 0 - Planning Phase

### Scope

  * Describe the roadmap process, how scope is determined for mid to long term features, as well as how the roadmap maps back to current contributions and maintainer ladder?
  * Describe the target persona or user(s) for the project?
  * Explain the primary use case for the project. What additional use cases are supported by the project?
  * Explain which use cases have been identified as unsupported by the project.  
  * Describe the intended types of organizations who would benefit from adopting this project. (i.e. financial services, any software manufacturer, organizations providing platform engineering services)?  
  * Please describe any completed end user research and link to any reports.

### Usability

  * How should the target personas interact with your project?  
  * Describe the user experience (UX) and user interface (UI) of the project.  
  * Describe how this project integrates with other projects in a production environment.

### Design

  * Explain the design principles and best practices the project is following.   
  * Outline or link to the project’s architecture requirements? Describe how they differ for Proof of Concept, Development, Test and Production environments, as applicable.  
  * Define any specific service dependencies the project relies on in the cluster.  
  * Describe how the project implements Identity and Access Management.  
  * Describe how the project has addressed sovereignty.  
  * Describe any compliance requirements addressed by the project.  
  * Describe the project’s High Availability requirements.  
  * Describe the project’s resource requirements, including CPU, Network and Memory.  
  * Describe the project’s storage requirements, including its use of ephemeral and/or persistent storage.  
  * Please outline the project’s API Design:  
    * Describe the project’s API topology and conventions  
    * Describe the project defaults  
    * Outline any additional configurations from default to make reasonable use of the project  
    * Describe any new or changed API types and calls \- including to cloud providers \- that will result from this project being enabled and used  
    * Describe compatibility of any new or changed APIs with API servers, including the Kubernetes API server   
    * Describe versioning of any new or changed APIs, including how breaking changes are handled  
  * Describe the project’s release processes, including major, minor and patch releases.

### Installation

  * Describe how the project is installed and initialized, e.g. a minimal install with a few lines of code or does it require more complex integration and configuration?  
  * How does an adopter test and validate the installation?

### Security

  * Please provide a link to the project’s cloud native [security self assessment](https://tag-security.cncf.io/community/assessments/).  
  * Please review the [Cloud Native Security Tenets](https://github.com/cncf/tag-security/blob/main/security-whitepaper/secure-defaults-cloud-native-8.md) from TAG Security.  
    * How are you satisfying the tenets of cloud native security projects?  
    * Describe how each of the cloud native principles apply to your project.  
    * How do you recommend users alter security defaults in order to "loosen" the security of the project? Please link to any documentation the project has written concerning these use cases.  
  * Security Hygiene  
    * Please describe the frameworks, practices and procedures the project uses to maintain the basic health and security of the project.   
    * Describe how the project has evaluated which features will be a security risk to users if they are not maintained by the project?  
  * Cloud Native Threat Modeling  
    * Explain the least minimal privileges required by the project and reasons for additional privileges.  
    * Describe how the project is handling certificate rotation and mitigates any issues with certificates.  
    * Describe how the project is following and implementing [secure software supply chain best practices](https://project.linuxfoundation.org/hubfs/CNCF\_SSCP\_v1.pdf) 

## Day 1 \- Installation and Deployment Phase

### Project Installation and Configuration

  * Describe what project installation and configuration look like.

### Project Enablement and Rollback

  * How can this project be enabled or disabled in a live cluster? Please describe any downtime required of the control plane or nodes.  
  * Describe how enabling the project changes any default behavior of the cluster or running workloads.  
  * Describe how the project tests enablement and disablement.  
  * How does the project clean up any resources created, including CRDs?

### Rollout, Upgrade and Rollback Planning

  * How does the project intend to provide and maintain compatibility with infrastructure and orchestration management tools like Kubernetes and with what frequency?  
  * Describe how the project handles rollback procedures.  
  * How can a rollout or rollback fail? Describe any impact to already running workloads.  
  * Describe any specific metrics that should inform a rollback.  
  * Explain how upgrades and rollbacks were tested and how the upgrade-\>downgrade-\>upgrade path was tested.  
  * Explain how the project informs users of deprecations and removals of features and APIs.  
  * Explain how the project permits utilization of alpha and beta capabilities as part of a rollout.

## Day 2 \- Day-to-Day Operations Phase

### Scalability/Reliability

  * Describe how the project increases the size or count of existing API objects.
  * Describe how the project defines Service Level Objectives (SLOs) and Service Level Indicators (SLIs).  
  * Describe any operations that will increase in time covered by existing SLIs/SLOs.  
  * Describe the increase in resource usage in any components as a result of enabling this project, to include CPU, Memory, Storage, Throughput.  
  * Describe which conditions enabling / using this project would result in resource exhaustion of some node resources (PIDs, sockets, inodes, etc.)  
  * Describe the load testing that has been performed on the project and the results.  
  * Describe the recommended limits of users, requests, system resources, etc. and how they were obtained.  
  * Describe which resilience pattern the project uses and how, including the circuit breaker pattern.

### Observability Requirements

  * Describe the signals the project is using or producing, including logs, metrics, profiles and traces. Please include supported formats, recommended configurations and data storage.  
  * Describe how the project captures audit logging.  
  * Describe any dashboards the project uses or implements as well as any dashboard requirements.  
  * Describe how the project surfaces project resource requirements for adopters to monitor cloud and infrastructure costs, e.g. FinOps  
  * Which parameters is the project covering to ensure the health of the application/service and its workloads?  
  * How can an operator determine if the project is in use by workloads?  
  * How can someone using this project know that it is working for their instance?  
  * Describe the SLOs (Service Level Objectives) for this project.
  * What are the SLIs (Service Level Indicators) an operator can use to determine the health of the service?

### Dependencies

  * Describe the specific running services the project depends on in the cluster.  
  * Describe the project’s dependency lifecycle policy.  
  * How does the project incorporate and consider source composition analysis as part of its development and security hygiene? Describe how this source composition analysis (SCA) is tracked.
  * Describe how the project implements changes based on source composition analysis (SCA) and the timescale.

### Troubleshooting

  * How does this project recover if a key component or feature becomes unavailable? e.g Kubernetes API server, etcd, database, leader node, etc.  
  * Describe the known failure modes.

### Security

  * Security Hygiene  
    * How is the project executing access control?  
  * Cloud Native Threat Modeling  
    * How does the project ensure its security reporting and response team is representative of its community diversity (organizational and individual)?  
    * How does the project invite and rotate security reporting team members?
