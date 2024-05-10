# CNCF TOC FAQ

A FAQ geared towards CNCF TOC and project issues.

## How do I join the TOC?

The TOC is a body elected by a variety of constituents. There is a public election schedule:
https://github.com/cncf/toc/blob/main/operations/election-schedule.md

The best way to get involved is to start attending TOC meetings and especially by participating in TAG(s)/WG(s).

If you have specific focus areas, CNCF TAG meetings are good opportunity to dive in:
https://github.com/cncf/toc/tree/main/tags

## What are CNCF TAGs?

The CNCF Technical Advisory Groups (TAGs) scale contributions by the CNCF technical and user community, while retaining integrity and increasing quality in support of our mission.

https://github.com/cncf/toc/tree/main/tags

## What type of governance is my CNCF project expected to follow?

CNCF does not require its hosted projects to follow any specific governance model by default.

Instead, CNCF [specifies](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md) that graduated projects need to "[e]xplicitly define a project governance and committer process." 

This varied and [open governance](https://github.com/opengovernance/opengovernance.dev) approach has led to different projects defining what is best and optimized for their community: 

* https://github.com/containerd/project/blob/master/GOVERNANCE.md
* https://github.com/coredns/coredns/blob/master/GOVERNANCE.md
* https://github.com/envoyproxy/envoy/blob/main/GOVERNANCE.md
* https://github.com/helm/community/blob/master/governance/governance.md
* https://github.com/kubernetes/community/blob/master/governance.md
* https://prometheus.io/governance/
* https://thanos.io/governance.md/

All project assets like trademarks, domains, builds, registries, github are neutrally owned by the foundation members versus a single vendor.

## Does the CNCF TOC control CNCF projects in any way?

Neither the CNCF Governing Board (GB) nor the Technical Oversight Committee (TOC) is responsible for managing CNCF-hosted projects. 

Instead, the maintainers of those projects manage them; this includes defining the governance process. The GB is responsible for the budget.

TOC members are available to provide guidance and conflict resolution if desired to the projects but do not control them. The TOC also helps mature projects through the various CNCF project maturity levels to ensure projects meet the expected [graduation criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md). To date, we have had no meaningful disagreements between the TOC and project maintainers.

Please see the [TOC Principles](https://github.com/cncf/toc/blob/main/PRINCIPLES.md) for more details.

## What type of support do CNCF projects receive?

The CNCF provides a variety of services that are accessible by maintainers via the ServiceDesk: https://github.com/cncf/servicedesk

## What is the definition of an adopter?
Adopters of a CNCF project are any organization that successfully leverages that project in the manner it was intended or repackages it as a core component of a service offering. To increase understanding of how the project is adopted, we distinguish a direct adopter from a transitive adopter; i.e. a Service Provider using Kubernetes as the underlying container orchestrator for their Knative offering, in this scenario the Service Provider is a direct adopter of Kubernetes and their Knative users are direct adopters of Knative but transitive adopters of Kubernetes.

Direct adopters of a project are responsible for the project's development, packaging, configuration, or deployment in their use of it. Transitive adopters of a project are not responsible for the project's development, packaging, configuration, or deployment, however they may receive the benefits that the project provides. In the case of the example above, the Service Provider's use of Kubernetes enables their customers to use Knative, which relies on Kubernetes as part of its architecture. However, if the Service Provider's implementation enables their customers to develop, package, configure, or deploy Kubernetes as part of the Knative offering, those customers are then also direct adopters.

The TOC’s intent in identifying adopters is to better understand the operational or production-level use of a given project (which can also be a specification) by its operators or users. We do this to ascertain the level of maturity the project has reached, its interactions with adopters, and its likelihood to continue growing within and supporting the ecosystem. Original sponsoring organizations, if any, may be included in the ADOPTERS.md file.  It is the discretion of the TOC to determine whether a given adopter, for the purposes of interviews, can provide quality feedback on where the project is and where it is going next in their use and interactions with the project.

[Cloud native](https://github.com/cncf/toc/blob/main/DEFINITION.md) project adopters may be any one of the following. Please note that any single company can fall under several categories at once. If that's the case, it should enumerate all that apply and only be listed once. For the purposes of identifying the adopters of a project, we encourage the use of `adopter` to include all of the below, if a project needs to identify the category of adopter we encourage them to use the convention provided below.
* A CNCF End-User member - Companies and organizations who are End-User members of the CNCF (https://www.cncf.io/enduser/ & https://landscape.cncf.io/card-mode?enduser=yes). This group is identified in the written form by the convention `End-User`, capitalizing the `E` in `End`, hyphenating the two words, and capitalizing the `U` in `User`.
* Another project - an open source project that leverages a CNCF project as part of their solution, integrates with for compatibility and interoperability, or is used in the supply chain of another project
* End users - companies and organizations that are not CNCF End-User members that use the project and cloud native technologies internally,  or build upon a cloud native open source project but do not sell the cloud native project externally as a service offering (those are Service Providers). This group is identified in the written form by the convention `end user`, uncapitalized and unhyphenated.
* A Service Provider - a Service Provider is an organization that repackages an open source project as a core component of a service offering, sells cloud native services externally
  * APIs
  * SaaS
  * A Service Provider’s customers are considered transitive adopters and should be excluded from identification within the ADOPTERS.md file.
  * Examples of Service Providers (and not end users) include cloud providers (e.g., Alibaba Cloud, AWS, Google Cloud, Microsoft Azure), some infrastructure software vendors, and telecom operators (e.g., AT&T, China Mobile). Refer to the [vendor category on the landscape](https://landscape.cncf.io/category=cncf-members&enduser=no&format=card-mode&grouping=category) for more examples.
* Consultancy - an entity whose purpose is to assist other organizations in developing a solution leveraging cloud native technology. They may be embedded in the end user team and is responsible for the execution of the service. Service Providers may also provide consultancy services, they may also package cloud native technologies for reuse as part of their offerings. These function as proxies for an end user.

Projects may leverage the above guidelines to list organizations in their ADOPTERS.md file within their repo, projects are not required to identify the type of adopter in their ADOPTERS.md file. 

If you’re not sure if your organization falls into any of these categories you can ask in the #toc slack channel, on the TOC mailing list, or email info@cncf.io and we’ll help you.

## What does CNCF do with its events revenues, especially for projects?

CNCF uses revenues from events first and foremost to cover the cost of events. The remaining money, which can be thought of as the events’ profits, are used to directly support CNCF’s hosted projects. Two recent examples of this include funding $250,000 for the Kubernetes security audit and funding nearly $100,000 for Jepsen testing of etcd. The services CNCF provides to its hosted projects are detailed in the [CNCF Annual Report](https://www.cncf.io/cncf-annual-report-2018/) and the Project Journey Reports for [Kubernetes](https://www.cncf.io/cncf-kubernetes-project-journey/) and Envoy.

## So CNCF invests all events revenues back into the events or to support its hosted projects?

**Yes**! In fact, CNCF invests both all events profits *and* a significant portion of its membership revenue in support of its hosted projects.

## Does CNCF make decisions based on “pay-to-play” considerations?

**No**. CNCF’s charter [Section 3(c)](https://github.com/cncf/foundation/blob/main/charter.md) explicitly calls out fairness as a core value and requires us to ‘avoid undue influence, bad behavior or “pay-to-play” decision-making.’ Also projects are [openly governed](https://www.cncf.io/blog/2019/08/30/cncf-technical-principles-and-open-governance-success/) and technical participation never requires any form of membership. The correct way to view this is that CNCF members **PAY TO SUSTAIN** and govern the shared budget which goes to sustain project activities such as events, scholarships and security audits.

## Does membership or sponsorship level ever affect project-related decisions? 

**No**. *All project-related decisions are made by the project maintainers*. Maintainership and governance processes are decided by the projects without regard to CNCF membership. Here’s more detail on [open governance](https://www.cncf.io/blog/2019/08/30/cncf-technical-principles-and-open-governance-success/).

## Why do members join if they don’t get to control CNCF and its projects?

The main member benefit is to support the cloud native ecosystem and to be publicly seen supporting it. You can see other benefits on the [Join](https://www.cncf.io/about/join/) page, but none of them involve control over projects. The governing board has ultimate control over CNCF, but CNCF is managed through a written charter and it consists of many disparate and competing interests. As a result, decision-making is consensus-based.

## Do I have to use my real name when contributing via DCO/CLA?

The DCO requires the use of a real name that can be used to identify someone in case there is an issue about a contribution they made.

A real name does not require a legal name, nor a birth name, nor any name that appears on an official ID (e.g. a passport). Your real name is the name you convey to people in the community for them to use to identify you as you.

See the full guidelines here: https://github.com/cncf/foundation/blob/main/dco-guidelines.md
