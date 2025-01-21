# CNCF TAG App Delivery Charter

Primary Authors: Michelle Noorali, Alexis Richardson

Approved September 3rd, 2019

Reviewed and contributed to by:

* Quinton Hoole
* Gareth Rushgrove
* Alois Reitbauer
* Doug Davis
* Lei Zhang
* Roger Klorese
* Taylor Dolezal
* Jianbo Sun
* Nicolas Trangez
* Carolyn Van Slyck
* Jeremy Rickard
* Marc Campbell
* Shantanu Deshpande
* Xander Grzywinski
* Kunal Parmar
* Ricardo Aravena
* Dave Forgac
* Max Körbächer

## **Introduction**

The charter describes the operations of the CNCF TAG Application Delivery. The Application Delivery TAG focuses on delivering cloud native applications which involves multiple phases including building, deploying, managing, and operating. Additionally, the TAG produces supporting material and best practices for end-users, and provide guidance and coordination for CNCF projects working within the TAG’s scope.

## **Mission**

Consistent with the [CNCF TAG definition](https://github.com/cncf/toc/blob/main/tags/cncf-tags.md), the mission of CNCF TAG App Delivery is:

* To collaborate on areas related to developing, distributing, deploying, managing and operating secure cloud-native applications with the target of delivering application in manner of cloud native.
* To develop informational resources including guides, tutorials and white papers to give the community an understanding of best practices, trade-offs, and value-adds regarding to application delivery.
* To identify suitable projects and gaps in the landscape, periodically updating the TOC with suggested actions in a structured manner. This includes helping the TOC with assessments and due diligence of prospective new projects.

## **Areas considered in Scope**

TAG Application Delivery focuses on the following topics of the lifecycle of cloud-native applications:

* Application definition, including description, parameter and configuration
* Guidance and practice for application design and development
* Application bundling and deployment
* Package management
* Application delivery workflow and strategy
* Configuration source driven workflow
* Release management

The TAG will work on developing best practices, fostering collaboration between related projects, working on improving tool interoperability as well as proposing new initiatives and projects when blank spots in the current landscape are identified. 

For CNCF projects, the scope of application delivery TAG engages, amongst others, with the application management focused ones, for example:

* Brigade
* Buildpacks
* CloudEvents
* Flux
* Helm
* Kubernetes

The following, non exhaustive, sample list of activities and deliverables are in-scope for the TAG:

* Education material to help provide guidance for the community
    * Summary and overview of projects available in the community
    * Definitions of implementations and patterns for best practices for delivering cloud native applications at enterprise scale
    * Tooling composition and tool chain creation based on existing projects. 
    * Best practices for development, operations and monitoring workflows
* Application definition, development and distribution
    * Application descriptor
    * Artifacts (e.g. bundles, packaging, images), package management and repositories
    * Configuration customization (e.g. templating and parameter)
* Guidance for application development and architecture
    * Reference design of microservices, decomposing monoliths
    * Reference design of event-based architectures
* Standard and generic application delivery pattern
    * Typical rollout strategy implementations with reference architectures
    * Best practices of configuration source driven workflow for application delivery
    * Best practices of standard application delivery pipeline
* Other tools or practices which may be involved in the workflow of application delivery, including but not limited to:
    * Building and configuration
    * Deploying, Orchestration and Patching
    * Policy and security
    * Debugging and monitoring
    * Hosting environments and interoperability (e.g. PaaS, FaaS, CaaS,...)
    * CI/CD
* Serverless - Serverless applications are a core part of cloud-native application development and the Serverless WG will migrate to live under this TAG.

## **Areas considered out of Scope**

Anything not explicitly considered in the scope above. Example include:

* Testing.
* Non cloud-native applications. 
* Developer tools which are not closely related to delivering application to cloud.
* Specific programming model or developing framework.
* Definition of standards for components like container images and other infrastructure-level building blocks of cloud-native applications. 

## **Roadmap**

* White paper explaining "Cloud Native Applications" and the current state of the art as seen in the community:
    * **Clarify the terminology** currently in use in the cloud native application space, and the relationships between the various terms.
    * Figure out patterns and practices of application definition, distribution, and delivery in the manner of "cloud native" in current community, it may include areas like application development, but it mainly focuses around how to deliver application to cloud.
    * Provide some general examples of **how these patterns and practices are currently being used in production** in public or private or hybrid cloud environments.
* Creating an application delivery landscape based on the outputs from the white paper. The TAG will help to reduce confusion and educate users by identifying finer-grained problems and implementation choices. Eg "project X focuses on application level rollout strategy", “project Y focuses on GitOps”.
* Define and standardize generic rollout models covering typical app delivery patterns with concrete use cases and practices. 
    * The models include but are not limited to Blue-Green Deployment, A/B Testing, Canary Deployment and Analysis, Progressive Traffic Shifting and GitOps.
    * The models are not expected to be bound to any specific runtime or execution engine. If one has to, it may not be a good candidate for this part.

## **Governance**

### Cross-group relationships

Lifecycle management of applications is a broad and mainstream topic of Cloud Native computing; therefore this TAG may collaborate with most of the other CNCF TAGs and projects. However, the following groups might have the largest potential interactions:

* **TAG Security** - The publication of guidance or tutorials by the TAG could see the adoption of insecure practices if security isn’t considered as a prerequisite for publication. Collaborating with TAG Security on reviews should help to ensure guidance doesn’t lead to propagating insecure patterns of usage.
* **Kubernetes SIG Apps** - Many projects currently under Kubernetes SIG Apps may overlap with CNCF TAG App Delivery. The application delivery TAG will focus on end-to-end aspects of these projects, including non-Kubernetes platforms and projects where applicable; while Kubernetes SIG Apps will focus on Kubernetes-specific runtime-level concerns. Close collaboration is expected to happen within these two groups around different phases for application delivery.

## **Operations**

* TOC Liaisons: Katie Gamanji, Lin Sun
* TAG chairs: [Lian Li](https://github.com/lianmakesthings), [Roberth Strand](https://github.com/roberthstrand).
* Tech Leads: [Sarah Christoff](https://github.com/schristoff), [Dylan Page](https://github.com/genpage)
* See [roles](https://github.com/cncf/tag-security/blob/main/governance/roles.md#role-of-chairs) for more information
* Slack channel: #tag-app-delivery in CNCF workspace - [https://cloud-native.slack.com/messages/CL3SL0CP5](https://cloud-native.slack.com/messages/CL3SL0CP5) 

## **Contact**

* [Slack Channel (#tag-app-delivery)](https://cloud-native.slack.com/messages/CL3SL0CP5 )
* Join TAG-App-Delivery at [lists.cncf.io](http://lists.cncf.io)
