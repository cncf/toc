# Cloud Native Buildpacks Incubating Stage Review

Cloud Native Buildpacks (CNB) is currently a Sandbox-stage member project. Please refer to the [CNB Sandbox stage proposal](https://github.com/cncf/toc/blob/master/proposals/buildpacks.adoc) for alignment with the CNCF and details on sandbox requirements.

Deliverables include:
* [spec](https://github.com/buildpacks/spec), the specification for CNB
* [lifecycle](https://github.com/buildpacks/lifecycle), a reference implementation
* [pack](https://github.com/buildpacks/pack), a local CLI for building apps using CNB
* [libbuildpack](https://github.com/buildpacks/libbuildpack), Go language binding for the CNB API

Deliverables do not include buildpacks themselves. Instead this project is about fostering an ecosystem where individuals or companies can use the tools provided by this project to create and use them.

## Community
Most people interested in buildpacks will not make contributions to the spec itself, but instead use the tools provided by the project. For instance, users may build their own buildpacks or run existing buildpacks to package their applications. Builder images on Docker Hub and other registries provide one mechanism by which users consume buildpacks.  Docker Hub lists 10m+ pulls from  the most popular public [builder](https://hub.docker.com/r/cloudfoundry/cnb) [images](https://hub.docker.com/r/heroku/buildpacks).

We've introduced a [RFC Process](https://github.com/buildpacks/rfcs/blob/master/text/0004-rfc-process.md): a well-documented procedure for proposing substantive changes to the project used by everyone (including maintainers). The project has accepted RFCs from VMware/Pivotal, Salesforce/Heroku, and Google.

There's an active community engaged in our twice weekly working groups, with attendance over the last two months of around 13-20 people, representing roughly five companies/projects (including Salesforce/Heroku, VMware/Pivotal, Deft, Dokku, Microsoft, and Google).

Since joining the sandbox project, our [slack instance](https://slack.buildpacks.io) has grown to 400+ users. We consider users who engage with us an equally important part of the community as maintainers.

To make it clear how the project operates, we've published a [governance policy](https://github.com/buildpacks/community/blob/master/GOVERNANCE.md). We also expect everyone to abide by our [Code of Conduct](https://github.com/buildpacks/.github/blob/master/CODE_OF_CONDUCT.md). To kickoff the new year and help with alignment we've published a [roadmap for 2020](https://github.com/buildpacks/community/blob/master/ROADMAP.md).

## Incubating Stage Criteria
To be accepted into Incubation, a project must meet the Sandbox requirements plus:

* Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.
  * [Project riff](https://projectriff.io/blog/2018/11/19/announcing-riff-0-2-0)
  * [Microsoft Azure Container Registry](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-tasks-pack-build)
  * [Google Cloud Run](https://cloud.google.com/blog/products/serverless/introducing-cloud-run-button-click-to-deploy-your-git-repos-to-google-cloud)
  * [Pivotal kpack](https://content.pivotal.io/blog/introducing-kpack-a-kubernetes-native-container-build-service)
  * [Pivotal Build Service](https://content.pivotal.io/blog/pivotal-build-service-now-alpha-assembles-and-updates-containers-in-kubernetes)
  * [Salesforce Evergreen](https://developer.salesforce.com/blogs/2019/11/introducing-salesforce-evergreen.html)
  * [Weave Firekube](https://www.weave.works/blog/firekube-fast-and-secure-kubernetes-clusters-using-weave-ignite)
  * [Azure Spring Cloud](https://content.pivotal.io/blog/azure-spring-cloud-a-new-way-to-run-spring-boot-apps-atop-kubernetes)

* Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.
  * Owners - 4
  * Mantainers - 6
  * Committers -  20
  * Maintainers are added and removed as per the policy outlined in the [project governance](https://github.com/buildpacks/community/blob/master/GOVERNANCE.md).


* Demonstrate a substantial ongoing flow of commits and merged contributions.
  * Releases
    * pack has 6 minor releases, with [v0.6.0](https://github.com/buildpacks/pack/releases) being the latest.
    * lifecycle has 5 releases, with [v0.5.0](https://github.com/buildpacks/lifecycle/releases/tag/v0.5.0) being the latest.
    * libbuildpack has [42 releases](https://github.com/buildpacks/libbuildpack/releases).
  * [pack Commit Activity](https://github.com/buildpacks/pack/graphs/commit-activity)
  * [CNCF DevStats](https://buildpacks.devstats.cncf.io)
    * [GitHub Activity from the last year](https://buildpacks.devstats.cncf.io/d/8/dashboards?orgId=1&from=now-1y&to=now-1h&refresh=15m)
    * [Commits per week over the last 6 months](https://buildpacks.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=now-6M&to=now)
    * [Issues Opened/Closed per week over the last 6 months](https://buildpacks.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=now-6M&to=now)
    * [New PRs per week over the last 1 year](https://buildpacks.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)
* A clear versioning scheme.
  * Versioning for our components and their compatibility is publicly documented in our [community repo](https://github.com/buildpacks/community/blob/master/VERSIONING.md).
* Specifications must have at least one public reference implementation.
  * [lifecycle](https://github.com/buildpacks/lifecycle)
  * [libbuildpack](https://github.com/buildpacks/libbuildpack)
