# LitmusChaos Graduation Proposal

### About LitmusChaos and its timeline in CNCF

LitmusChaos is an end-to-end chaos engineering platform which helps cloud-native developers and SREs carry out chaos experiments easily, and as defined by the principles of chaos. It extends the Kubernetes API to enable definition of custom fault-injection and resilience validation intent, while providing a wide variety of real-world failure scenarios out of the box via the ChaosHub. Several organizations leverage the ChaosCenter to construct, automate & manage chaos across their cluster fleet, while maintaining isolation at team and infrastructure levels. 

Litmus was initially accepted into the CNCF sandbox in June 2020 and progressed to incubation by January 2022. Since the promotion to incubation, the project has grown considerably and made significant strides across areas (adoption, contributions, major releases, integrations and security). In the light of these achievements, on behalf of the community and maintainer group, we recommend the project is now ready for graduation.    

### Progress since incubation and other project updates

Since moving to incubation in Jan 2022, LitmusChaos has made great strides in its mission to simplify cloud-native chaos engineering and empower more organizations to fortify their cloud native resilience practices. 

The project continues to maintain its monthly release cadence to address issues, make enhancements, address security vulnerabilities and deliver on roadmap items. 

A snapshot of the project stats **since incubation** is provided below: 


No. of New Releases|No. of New Formal Adopters|No. of New Contributors|No. of New Maintainers|No. of new Slack Members 
---------------|--------------------------|---------------|---------------------|---------------------
26|14|50+|12|1200+


### Adoption and contributions

Over 50+ organizations have publicly stated their Litmus adoption stories, with several more actively using it and in discussions with the project team via slack and other channels. LitmusChaos has attained good adoption levels and has contributions coming in from organizations across various domains, including cloud-native software vendors, telcos, airlines, e-commerce operators, fintech etc. Litmus is well represented in several CNCF meetups and conferences (KubeCon, Conf42, Chaos Carnival) and is one of the most active and sought-after participants in the LFX programme. 


## Features and project evolution 

LitmusChaos, the project, has evolved significantly with increased adoption and community feedback, even as the discipline of cloud native chaos engineering becomes mainstream. 

From initial capabilities that focused on providing custom resources to define and orchestrate fault injection (1.x), through additions of resilience probes, chaos metrics and centralized, multi-tenant management console (2.x) and subsequently, support for varied target environments, container runtimes and a powerful CLI (3.x), Litmus has evolved to support enterprise-level usecases.  

A brief snapshot of the project’s evolution is captured below. 

<img width="1121" alt="Screenshot 2024-04-10 at 6 03 38 PM" src="https://github.com/ksatchit/cncf-docs/assets/21166217/e4322c41-9425-4d61-8312-efbfddae2d85">


# Graduation State Criteria

### Note on components and repos

LitmusChaos components are largely composed of the control plane and execution plane components. 

Control plane consists of microservices that power the management, i.e., the web-frontend, chaos-manager/server (exposes chaos apis), authentication server (supports integration with Dex) and chaos database (mongodb). The control plane is usually setup once and in a centrally accessible location within the users’ environment.  

The execution plane comprises a set of custom-controllers to reconcile the chaos custom resources, a subscriber to connect/act as relay with the chaos manager/server, and a prometheus exporter. The execution plane typically maps to a user Kubernetes cluster (or in some cases a namespace), with multiple execution planes/targets attached to a centralized control plane. The execution plane is where the chaos action actually takes place, via some transient jobs orchestrated by the custom chaos controllers. 

The [litmuschaos/litmus](https://github.com/litmuschaos/litmus/) repo is the meta repo associated with the LitmusChaos project and it also hosts the sources for the control plane microservices. It also serves as the location for all project governance (maintainers, security process, committer process, support horizon, code of conduct etc.,) and roadmap related information.


### Have committers (maintainers) from at least two organizations

Yes. LitmusChaos governance process ensures that all governance functions include members from multiple organizations. Considering Litmuschaos has a broad scope and involves multiple components, the maintainers in the project have been divided on the basis of the project components/areas, with ownership of the respective repositories. 

In addition to the maintainers, the project also recognizes a role called Reviewers, which is comprised of regular contributors that have made or influenced significant changes and improvements to the project, and are potentially in line for moving to maintainer roles over time. 

 https://github.com/litmuschaos/litmus/blob/master/MAINTAINERS.md


### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/)

Yes, we have achieved this for all the project components. The badge has been added on the main/meta repository litmuschaos/litmus to reflect the same. The OpenSSF Best Practices badge is displayed prominently on the main LitmusChaos project repository
[README.md](https://github.com/litmuschaos/litmus/blob/master/README.md), and
all the details can be explored on the project’s page at (https://www.bestpractices.dev/projects/3202).

### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation

A formal security review is in progress with the CNCF. In addition, we are adopting SBOM best practices across the project via signed releases, etc..,

### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers

Project governance, committer lifecycle, and other documents can be found in the [community repository](https://github.com/litmuschaos/litmus/blob/master/GOVERNANCE.md)

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence

The [Governance Charter](https://github.com/litmuschaos/litmus/blob/master/GOVERNANCE.md) describes how maintainers are added; the maintainers are the primary interface between the project and the CNCF. It also explains the criteria for moving maintainers to emeritus status. The active and emeritus maintainers can be found in the maintainers file: [maintainers.md](https://github.com/litmuschaos/litmus/blob/master/MAINTAINERS.md)

### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec. Refer to [FAQs](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for guidelines on identifying adopters

Adopters along with their usecases can be found here: https://github.com/litmuschaos/litmus/blob/master/ADOPTERS.md


