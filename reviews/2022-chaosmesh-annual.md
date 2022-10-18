
# Chaos Mesh 2022 Annual Review

  - [Background](#background)
  - [Alignment with Cloud Native](#alignment-with-cloud-native)
  - [Year-in Review](#year-in-review)
  - [Annual Review Contents](#annual-review-contents)
  - [Project Links](#project-links)


## Background

Chaos Mesh is a versatile Chaos Engineering platform that orchestrates chaos experiments on Kubernetes environments.


By covering comprehensive fault injection methods in container, Pod, network, file system, and even the kernel, Chaos Mesh aims at providing a neutral, universal Chaos Engineering platform that enables cloud-native applications to be as resilient as they should be.

Chaos Mesh contains the following components:

![Chaos Mesh arch](https://github.com/chaos-mesh/chaos-mesh/blob/master/static/chaos-mesh.png)

* **Chaos Dashboard:** a Web UI for managing, designing, monitoring Chaos Experiments, through which users can operate and observe chaos experiments. It also provides RBAC authorization.

* **Chaos Controller Manager:** the core logic component of Chaos Mesh. It’s responsible for the scheduling and managing chaos experiments. The Controller Manager contains multiple CRD Controllers, such as Workflow Controller, Scheduler Controller, and various types of fault controllers.

* **Chaos Daemon:** the main executive component of Chaos Mesh. Chaos Daemon runs in [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) mode and has privileged permissions by default (can be turned off). Chaos Daemon mainly interferes with specific network devices, file systems, kernels, etc. by intruding into the target Pod Namespace.

**Chaos Mesh was accepted into CNCF Sandbox on July 15, 2020, and was promoted to a CNCF Incubating project on Feb 16, 2022**

* [Sandbox proposal PR](https://github.com/cncf/toc/pull/367)


* [Sandbox presentation slides](https://docs.google.com/presentation/d/115QvSCnT6ROwwV1lK-R10d4MW54fQhtVM2uZDGscpEk/edit#slide=id.g80f082b201_0_0)


* [Chaos Mesh proposal for incubation](https://github.com/cncf/toc/pull/683) 


* [Chaos Mesh Due Diligence for Incubation ](https://docs.google.com/document/d/1QQ3K9fOp-84UyX7OZbIndrIZs6HqF60k5DamKv9y8tM/edit#) 

# Alignment with Cloud Native

Chaos Mesh falls into the scope of[ CNCF TAG Network](https://github.com/cncf/tag-network).

Chaos Mesh adopts a Kubernetes architecture that uses Kubernetes CustomResourceDefinitions (CRDs) to define chaos objects, making it naturally integrated with the Cloud Native ecosystem.


By orchestrating chaos experiments via comprehensive fault injection methods, Chaos Mesh provides a vendor-neutral, scalable chaos engineering platform that enables cloud-native applications to be **_resilient, manageable, and observable_**.  

In addition, it integrates several other projects in the cloud-native ecosystem, such as Helm, Prometheus, and Grafana.

## Year-in Review

* Chaos Mesh was promoted to a **CNCF Incubating** project in Feb 16 2022! 


* We regulated our release/versioning scheme, and had the process [well documented](https://chaos-mesh.org/docs/release-cycle/). Chaos Mesh has 4 major releases and 21 patch releases since 2.0.

* Chaos Mesh core repo contributors has grown from 102 to 157 since last review.

* Chaos Mesh adopters has grown from 40+ to 60+ since last review.

## Annual Review Contents

* **Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.**


    According to [chaosmesh.devstats.cncf.io](http://chaosmesh.devstats.cncf.io), the Chaos Mesh project demonstrates a steady growth since last review, especially after being promoted to Incubating project, with (monthly average):
    * [53 issus opened and 43 closed](https://chaosmesh.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&var-period=m&var-repogroup_name=All);
    * [46 active contributors](https://chaosmesh.devstats.cncf.io/d/74/contributions-chart?orgId=1);
    * [41 merged PRs](https://chaosmesh.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all)

* **How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)**

    Chaos Mesh currently has gone through[ 6 Maintainers](https://github.com/chaos-mesh/chaos-mesh/blob/master/MAINTAINERS.md#project-maintainers) from 3 organizations (PingCAP, ByteDance Xpeng Motor, and one individual),  and 11 Committers from 6 organizations.  

* **What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)**

    This past year in CNCF has seen steady growth in adoption, from 40 to over 60+ adopters. To name a few:

    * Prudential - Use Chaos Mesh to simulate managed outages scenarios for comprehensive testing of distributed microservice product platforms, especially Java-powered runtimes.

    * Digital China - Uses Chaos Mesh to test the application's ability to maintain business functions and to discover errors and vulnerabilities that may occur under extreme conditions.

    * Adobe - Uses pod, stress and network actions in their chaos testing
    
    * AirBnB - Uses Chaos Mesh to test their distributed database
    * Databricks - Integrate Chaos Mesh in their DevOps to improve resiliency of new releases

    On the othe hand, we are glad to see more cloud service vendors are integrating Chaos Mesh in their offering. 

  * [Azure Chaos Studio](https://learn.microsoft.com/en-us/azure/chaos-studio/chaos-studio-tutorial-aks-portal)
  * [AWS Fault Injection Simulator ](https://medium.com/the-cloud-architect/what-is-aws-fault-injection-simulator-and-why-you-should-care-3fbe457ca227)

   See [ Adopters.md](https://github.com/chaos-mesh/chaos-mesh/blob/master/ADOPTERS.md) for more information of Chaos Mesh adoptions.There are many more adopters of Chaos Mesh and organizations evaluating the project as well.

* **How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)**

  * **Project growth**
      * 4 major releases since last review ( 2.1, 2.2, 2.3, 2.4) and 21 patch releases with following new features and enhancements:
          * New Chaos: integration with chaosd to introduce Chaos with Physical Machines
          * New Chaos: integration with more cloud providers including AzureChaos and GCPChaos
          * ARM64 Support: ported PodChaos, StressChaos, NetworkChaos, TimeChaos, HTTPChaos on ARM64
          * Enriching Observability: introduce more logs and metrics
          * Chaos Dashboard: New UI for creating Chaos Experiments and Workflow
          * New StatueCheck CRD, for checking the health state of application
  * **Community growth**
      * Ever since entering the CNCF Sandbox, we have seen substantial growth in the community. Project stats have grown notably in the past year in CNCF: 

        | **Stats**         | **Current** | **Growth Since Last Review** |
        |-------------------|-------------|------------------------------|
        | Stargazers        | 5.2 k       | 1.5 k                        |
        | Commits           | 2.16k       | 1 k+                         |
        | Forks             | 657         | 157                          |
        | Contributors      | 159         | 57                           |
        | Slack members     | 657         | 420                          |
        | Twitter followers | 1.6k        | 1k (approximately)           |

      * Actively Took part in CNCF and self-hosted events and activities to let more people know, learn, and use Chaos Mesh and Chaos Engineering
  
          * Newly added [Chaos Mesh Development Biweekly](https://docs.google.com/document/d/1s9X6tTOy3OGZaLDZQesGw1BNOrxQfWExjBFIn5irpPE/edit?pli=1#heading=h.68utt89hrdrl) meeting for Chaos Mesh committers to discuss project and community progress (focus on the development)
          * 3 Chaos Mesh KubeCon OfficeHours in KubeCon 21 Euro, NA and KubeCon 22 Spain.
          * One KubeCon talk - [Make Cloud Native Chaos Engineering Easier - Deep Dive into Chaos Mesh](https://kccnceu2022.sched.com/#) (maitainer tracks)
          * 2 GSoC Projects
              * [Single binary deployment outside kubernetes environment.](https://summerofcode.withgoogle.com/programs/2022/projects/yKOpuIZh)
              * [RPC cross different namespaces through unix socket](https://summerofcode.withgoogle.com/programs/2022/projects/dFZyvm9F)
          * 2 LFX Mentorship Projects
              * [Monitoring Metrics about Chaos Mesh](https://mentorship.lfx.linuxfoundation.org/project/8db683b0-0273-4a83-9ed9-4c33ee2cfcf0)
              * [Interactive Katacoda Payground for Chaos Mesh](https://mentorship.lfx.linuxfoundation.org/project/09847d84-5d14-4c05-8644-57cdde5b6466)

* **What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?** 

    * Working on supporting multiple Kuberenetes clusters at present, and plan to release in version 2.5 
    * Support more status check types, such as prometheus, Datadog 
    * Provide a plugin approach to extend complex chaos types, such as RabbitMQChaos, RedisChaos, …
    * Refactoring our documentation to follow a more intuitive and accessible structure

* **How can the CNCF help you achieve your upcoming goals?**
    * CNCF can also help Chaos Mesh by providing visibility to drive adoption and attract contributions
    * Drive collaborations and integrations with other projects under the CNCF landscape.
* **Do you think that your project meets the criteria for Graduation?**
    * Not yet. We need to work to drive more adoptions, on the security aspects 

## Project Links

* [Website](https://chaos-mesh.org/)
* [GitHub](https://github.com/chaos-mesh/chaos-mesh)
* Slack:[ CNCF](http://slack.cncf.io) / #project-chaos-mesh
* [Twitter](https://twitter.com/chaos_mesh)
* [YouTube Channel](https://www.youtube.com/channel/UC4OwT4QTd0ML3YNnV1ybT6g)
* [Chaos Mesh Development Biweekly](https://docs.google.com/document/d/1s9X6tTOy3OGZaLDZQesGw1BNOrxQfWExjBFIn5irpPE/edit?pli=1#heading=h.68utt89hrdrl)
* [Monthly community meeting](https://community.cncf.io/chaos-mesh-community/)