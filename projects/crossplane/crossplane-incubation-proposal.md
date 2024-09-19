# Crossplane CNCF Incubation Proposal

## Background

Crossplane is an open source Kubernetes add-on that enables platform teams to assemble
infrastructure from multiple vendors, and expose higher level self-service APIs for application
teams to consume. Crossplane effectively enables platform teams to quickly put together their own
opinionated platform declaratively without having to write any code, and offer it to their
application teams as a self-service Kubernetes-style declarative API.

Both the higher level abstractions as well as the granular resources they are composed of are
represented simply as objects in the Kubernetes API, meaning they can all be provisioned and managed
by kubectl, GitOps, or any tools that can talk with the Kubernetes API. To facilitate reuse and
sharing of these APIs, Crossplane supports packaging them in a standard OCI image and distributing
via any compliant registry.

Platform engineers are able to define organizational policies and guardrails behind these
self-service API abstractions. The developer is presented with the limited set of configuration that
they need to tune for their use-case and is not exposed to any of the complexities of the low-level
infrastructure below the API.  Access to these APIs is managed with Kubernetes-native RBAC, thus
enabling the level of permissioning to be at the level of abstraction.

While extending the Kubernetes control plane with a diverse set of vendors, resources, and
abstractions, Crossplane recognized the need for a single consistent API across all of them. To this
end, we have created the Crossplane Resource Model (XRM).  XRM extends the Kubernetes Resource Model
(KRM) in an opinionated way, resulting in a universal experience for managing resources, regardless
of where they reside. When interacting with the XRM, things like credentials, workload identity,
connection secrets, status conditions, deletion policies, and references to other resources work the
same no matter what provider or level of abstraction they are a part of.

The functionality and value of the Crossplane project can be summarized at a very high level by
these two main areas:

1. Enabling infrastructure owners to build custom platform abstractions (APIs) composed of granular
   resources that allow developer self-service and service catalog use cases
1. Providing a universal experience for managing infrastructure, resources, and abstractions
   consistently across all vendors and environments in a uniform way, called the Crossplane Resource
   Model (XRM)

### Sandbox Proposal Artifacts

* Sandbox Proposal PR: [PR #454](https://github.com/cncf/toc/pull/454)
* Sandbox Proposal doc: [Proposal
  doc](https://github.com/cncf/toc/blob/master/proposals/sandbox/crossplane.md)
* Sandbox SIG app-delivery presentation: [Slide
  deck](https://docs.google.com/presentation/d/1HEQtVMwQxrpkQg2UWUNi0RPdjmJ6TaZLr50Cr_anNeQ/edit?usp=sharing)

### Sandbox Acceptance

Crossplane was accepted into the CNCF Sandbox on June 24th, 2020.

## Incubation Proposal Resources

* Due Diligence Document: [DD doc](https://docs.google.com/document/d/1OBGR7h7mqQ1zE1yheIH8oPkTHiVaEH_DdT6QCTdF370/edit?usp=sharing)
* SIG App-Delivery presentation: [Slide deck](https://docs.google.com/presentation/d/12YLgn_J5WyvW7s4dab8eF0ebquWzB4Zv3bzi-klUW1w/edit?usp=sharing)

## Progress since joining the CNCF Sandbox

The biggest accomplishment from the Crossplane community is the release of [v1.0, the first major
milestone](https://blog.crossplane.io/announcing-crossplane-v1-0/) in the project’s history, which
was followed recently by even more improvements in
[v1.1](https://blog.crossplane.io/crossplane-v1-1-vault-integration-enhanced-composition-and-aws-provider-code-generation/),
[v1.2](https://blog.crossplane.io/crossplane-v1-2-0/), and
[v1.3](https://github.com/crossplane/crossplane/releases/tag/v1.3.0).
With these versions of Crossplane, the project is now stable and ready to be used in production
scenarios. This is further validated by project adopters successfully running Crossplane in
production environments. The core APIs have been moved to v1 (stable) and future enhancements and
improvements are intended to be done in a backwards compatible way with a clear upgrade path.  We do
not anticipate any further breaking changes in subsequent minor and patch releases.

With Crossplane reaching maturity, we are seeing end-users adopt Crossplane and find success running
it in their production environments. We consistently hear from users that have built their own
custom platforms APIs to enable self-service provisioning for their developers and accelerate the
delivery of their software to production. The community growth has been quite encouraging and we are
pleased to be able to include specific details about some of their adoption success in the sections
below.

The Crossplane community has expanded its collaboration with a rich ecosystem of partners, building
relationships and combining development efforts with a broader set of cloud providers and projects
in meaningful ways:

* IBM developed and released the [first version of their Crossplane provider for IBM
  Cloud](https://developer.ibm.com/blogs/ibm-joins-the-crossplane-community/), enabling the
  provisioning and management of 85+ hosted services from the IBM Cloud Catalog.
* Engineering teams from both [AWS and
  Azure](https://blog.crossplane.io/accelerating-crossplane-provider-coverage-with-ack-and-azure-code-generation-towards-100-percent-coverage-of-all-cloud-services/)
  have joined forces with the Crossplane community to share their code generation pipelines,
  allowing the Crossplane providers for AWS and Azure to increase their coverage and more easily
  stay updated with changes in their cloud offerings.
* RedHat partners with Crossplane for infrastructure provisioning, with a focus on hybrid scenarios.
  Crossplane has been available on [operatorhub.io](https://operatorhub.io/operator/universal-crossplane)
  since it’s v0.9 release.
* Alibaba Cloud's internal platforms have adopted Crossplane to model and manage their
  infrastructure resources.  Their [KubeVela project also uses
  Crossplane](https://kubevela.io/docs/platform-engineers/cloud-services/#install-and-configure-crossplane)
  as a cloud resource provider to provision infrastructure resources and bind to applications.
* GitLab has integrated Crossplane into [GitLab
  ADO](https://docs.gitlab.com/ee/user/clusters/crossplane.html) so users can easily provision and
  manage infrastructure from their Kubernetes clusters. Crossplane has been shipped with every
  GitLab release since version 12.5.
* Packet/Equinix developed the [Crossplane Provider for Equinix
  Metal](https://feedback.equinixmetal.com/changelog/crossplane-provider-for-equinix-metal), as well
  as
  [provider-ssh](https://www.youtube.com/watch?v=CxwmrKi1jtM&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=6)
  as a way to provision, manage, and run scripts against bare metal instances and VMs
* The following are examples of ecosystem projects that have collaborated and integrated with
  Crossplane:
  * [OPA](https://www.youtube.com/watch?v=TaF0_syejXc&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=13),
    a CNCF graduated project, makes for a very powerful combination with Crossplane to set global
    policy on the provisioning of cloud resources, further extending the self-service infrastructure
    capabilities.
  * [Falco](https://www.youtube.com/watch?v=uYmnZwsfblI&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=8),
    the CNCF runtime security project created by Sysdig, partnered with Crossplane to show how Falco
    can be used to secure infrastructure environments provisioned by Crossplane.
  * [NATS](https://www.youtube.com/watch?v=Rh6twPUPcGc&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=10),
    the CNCF open source messaging system, works alongside Crossplane to deliver secure solutions
    for end users.
  * [Linkerd](https://www.youtube.com/watch?v=o0-xkk6pCGM&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=15),
    the CNCF service mesh, can utilize Crossplane to configure the service mesh across a set of
    remote clusters.
  * Development efforts are underway with
    [Cluster-API](https://github.com/kubernetes-sigs/cluster-api/issues/4095) to use Crossplane as
    the [infrastructure engine for
    Cluster-API](https://www.youtube.com/watch?v=YBfSlh2duwo&list=PL510POnNVaaazaBA7jNh1f6Qz1LxuB5Ku&index=5),
    which would deduplicate community effort and accelerate support for more providers.
  * [Keptn](https://www.youtube.com/watch?v=V8L-JTpkEpc&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=3),
    which provides cloud native application life-cycle orchestration, can utilize Crossplane for
    infrastructure management in DevOps workflows.
  * [Sourcegraph](https://www.youtube.com/watch?v=WtsrrNLmmfM&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=4)
    can utilize Crossplane to enhance the experience of deploying Sourcegraph on Kubernetes by
    providing access to cloud provider managed services from within the cluster.
  * [SUSE](https://www.youtube.com/watch?v=exg_zrg16SI&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=6)
    worked with Crossplane to demonstrate the new Seccomp operator and build a Seccomp profile for
    Crossplane.
  * [OpenFaas](https://www.youtube.com/watch?v=XphQgB87U-s&list=PL510POnNVaaYFuK-B_SIUrpIonCtLVOzT&index=14)
    partnered with Crossplane to deploy and consume managed services alongside serverless functions.

Composition, the flagship feature of Crossplane, has also been promoted to a stable API since being
accepted into the Sandbox.  Composition is arguably the most compelling feature of Crossplane and is
where the project is seeing significant traction, since it enables platform engineers to essentially
build their own PaaS in a declarative way without having to write any code.  Platform teams use the
Composition feature to build their own custom declarative APIs, including policy guardrails, that
allow their application teams to provision the infrastructure resources they need in a self-service
manner.

The composition feature area has started getting significant contributions from the broader
community to enable new features and scenarios that are important to them, further increasing the
adoption, growth, and maturity of the project.  It has also enabled an ecosystem of solution focused
“configurations” to be built and shared with the community, removing some of the obstacles with
getting started and productive building your own platform with Crossplane.

The installation and management of providers (e.g. provider-aws) and configurations in Crossplane
has greatly improved with the second iteration of Crossplane’s package manager.  The package author
can express the dependencies it has on other packages, as well as version constraints, so that all
required packages (and their versions) are installed automatically.  Packages can also be
automatically upgraded and rolled back to ease adoption and roll out of new versions.

The Crossplane community held a very successful Community Day event to coincide with the release of
v1.0.  The speaker line up was quite impressive, including the [keynote
panel](https://www.youtube.com/watch?v=WGfYlssfIIk&list=PL510POnNVaaZx1QDIn_-77igM1COuHaHV&index=2)
of Kubernetes creators Joe Beda, Brian Grant, and Brendan Burns, which was moderated by Kelsey
Hightower.  Many adopters of the project also gave talks on their specific use cases and success
stories with Crossplane, while project maintainers led deep dives through the various feature areas
of the project.

We’ve also seen impressive growth in the community statistics, indicating greater adoption and a
growing ecosystem of people getting involved with the project.  Many stats have increased by at
least 50% in the last 9 months since being accepted into Sandbox, such as the number of GitHub stars
(~2,000 to 3,680+), commits to the project (~2,000 to 3,430+). Other metrics have seen greater
relative increases, such as container downloads (13x), contributors (3x), and Slack members (4x).
One particular area of growth that we’re most excited about is that we have started to receive
non-trivial feature contributions from the greater community, such as [bi-directional
patching](https://github.com/crossplane/crossplane/pull/2077) of composite resources, and
[composition patchSets](https://github.com/crossplane/crossplane/pull/1972), which indicates a more
serious adoption and commitment on their part.

### Project and Community Statistics

| Statistic | Sandbox | Current |
|-|-|-|
| GitHub stars | 2,120 | 3,680+ |
| Container downloads | 1,440,000 | 18,880,000+ |
| Releases | 24 | 32 |
| Commits | 2,215 | 3,430+ |
| Contributors (commits) | 48+ contributors | 184+ contributors from 105+ companies |
| Maintainers | 10 maintainers from 3 orgs | 18 maintainers from 5 orgs |
| Twitter followers | 1,484 | 3,132+ |
| Slack members | 610 | 2,744+ |
| Youtube subscribers | 215 | 743+ |

## Incubation Requirements

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope

1. [Fortune 100 Company interviewed by the TOC] - they are running Crossplane in 15 dev clusters and
   5 production clusters. They use it to enable teams to self-service on AWS infrastructure
   including RDS and other resources alongside their container workloads. They have been in
   production with Crossplane for at least a year.
1. [Zego](https://www.zego.com) provides insurance to both individuals and fleets who use vehicles
   to earn money. They are using Crossplane compositions to make it easier for developers to combine
   their applications with AWS infrastructure including databases, IAM roles, and other resources.
   They are in production with Crossplane.
1. [RipCord](https://www.ripcord.com) - they are running Crossplane in production clusters serving
   10s of tenants. They are using parameterized compositions to create instances of cloud resources
   for each of the tenants. This includes AWS buckets, queues, sns topics, IAM, and more.
1. [BBD](https://bbdsoftware.com/) - is a software development company that helps businesses from
   many sectors with digital solutions that drive seamless processes and smart systems. They are
   possibly the oldest production adopters as they have been running Crossplane in their production
   environments for over 18 months now. They are making effective use of Compositions within their
   production clusters and are currently managing around 7 teams and around 70 of production
   workloads. They have plans to continue scaling their Crossplane usage throughout various clients
   and projects.
1. [CloudCheckr](https://cloudcheckr.com/) - provides total visibility and insight into cloud
   management to lower costs, maintain security and compliance, and optimize resources.  They have
   successfully deployed Crossplane into production, in fact replacing many instances of their
   previous usage of Terraform.  CloudCheckr uses Crossplane to provision and manage cloud services
   for their production applications running in Kubernetes, while easily integrating with their
   existing CI/CD pipelines.
1. [Deutsche Bahn](https://www.bahn.com/en/view/index.shtml) - the largest railway operator in
   Europe, is now using Crossplane in production environments to empower their developers with
   self-service infrastructure. They are using it to launch EKS clusters and connect them to AWS
   cloud services including RDS and others. They are heavy users of Crossplane Compositions and use
   them to define their own opinionated platform on-top of AWS.
1. [Mothership](https://www.mothership.com/) has reimagined freight for modern commerce, bringing
   technology to the freight industry to provide on demand same day delivery services, real time
   tracking, and fully digital paperless processes. Their Crossplane usage began with provisioning
   databases like Amazon RDS from within Kubernetes. Building on the active reconciliation and
   status updates provided by the control plane, they have also been able to implement
   post-provisioning steps and day-2 operations. “It's been a huge win for our platform team. We're
   not worried about configuration drift and there's no requirement on our platform team to help
   developers provision their databases. We view Crossplane as a great building block for day-2
   infrastructure management.”
1. [Squiz](https://www.squiz.net/) enables their customers to create and choreograph their digital
   services with an open Digital Experience Platform (DXP). They are currently using Crossplane in
   production to provide services to their product development teams. They have been a fantastic
   contributor to the project with great feature additions to the composition engine, and have also
   started a new provider effort to manage CloudFlare resources:
   [provider-cloudflare](https://github.com/benagricola/provider-cloudflare). Squiz has found the
   composition functionality to be particularly valuable for their use cases, "the Composition
   mechanism enables a way of describing infrastructure and applications that is uniquely powerful
   when compared to other IAC / Platform tools".
1. [VSHN](https://www.vshn.ch/en/) is a leading Swiss company specializing in DevOps, Docker,
   Kubernetes, OpenShift and 24/7 Cloud Operations. Using Crossplane as a foundation, they have
   built a “Self-Service Managed Services Marketplace for developers to choose from various
   different service plans”. To enhance compatibility with existing provisioning workflows, they
   have developed an “[Open Service Broker capable
   client](https://github.com/vshn/crossplane-service-broker)” based on Crossplane that is fully
   integrated with their [platform and
   deployments](https://kb.vshn.ch/app-catalog/explanations/crossplane_service_broker.html).
1. [Accenture](https://www.accenture.com/) - they have adopted Crossplane and are building a
    repeatable solution that can be used across their customer base. They are using Crossplane
    compositions to define entire cloud-native deployments including Kubernetes clusters, and
    managed services around them.  They were an early adopter of Crossplane and have continuously
    provided valuable feedback on the features and capabilities of the project. They have been
    successful in building self-service infrastructure platforms for their clients, or providing an
    “infra-sdk” as they refer to it. Crossplane has been an effective solution to help them ensure
    that “teams create cloud resources in compliance with the company”.

There are many more adopters of Crossplane and companies evaluating the project as well. Some of
these have only consented to sharing their details privately with the TOC, if required.

### Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project

Across the various projects in the Crossplane organization, we currently have **18** maintainers
from **5** different organizations, including Upbound, Alibaba, Red Hat, Microsoft, and independent.

The requirements and process for [becoming a
maintainer](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md#becoming-a-maintainer)
is clearly defined in the project governance.

The [steering
committee](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md#steering-committee)
has been established and is currently in a bootstrapping period.  We intend to fill all seats in the
steering committee as a prerequisite for our full Graduation proposal.

### Demonstrate a substantial ongoing flow of commits and merged contributions

The Crossplane devstats page and dashboards can be found
[here](https://crossplane.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

The project currently has over
[184+](https://crossplane.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [105+
companies](https://crossplane.devstats.cncf.io/d/5/companies-table?orgId=1). According to devstats,
the project averages [317 commits per
month](https://crossplane.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now)
contained within [105 merged PRs per
month](https://crossplane.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now).

### A clear versioning scheme

All repositories in the Crossplane organization follow strict [semantic
versioning](https://semver.org/).  We also follow a regular release cycle of every 8 weeks and have
fully [documented this flow](https://crossplane.io/docs/v1.1/reference/release-cycle.html) to set
proper expectations across the community for our releases.
