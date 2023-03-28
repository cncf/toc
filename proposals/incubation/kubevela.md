# KubeVela proposal for CNCF Incubation

## About KubeVela

[KubeVela](https://kubevela.io/) is a modern software delivery control plane. The goal is to make deploying and operating applications across today's hybrid, multi-cloud environments easier, faster and more reliable. Applications created with KubeVela use the best practices of modern applications by default: they are able to scale with clouds, they use infrastructure as code, they are observable and secure.

Currently, KubeVela primarily provides the following key features:

* **Deployment as Code**: KubeVela declares your deployment plan as [workflow](https://kubevela.io/docs/getting-started/core-concept), run it automatically with any CI/CD or GitOps system, extend or re-program the workflow steps with CUE. It can glue and orchestrate all of your IaC based infrastructure capabilities as reusable modules, share with the community as [addons](https://kubevela.net/docs/reference/addons/overview). The deployment workflow in KubeVela is powered by [Open Application Model](https://oam.dev/).

* **Built-in security, compliance and observability building blocks**: Users can choose from the wide range of LDAP integrations provided out-of-box, enjoy enhanced [multi-cluster authorization and authentication](https://kubevela.io/docs/platform-engineers/auth/advance), pick and apply fine-grained RBAC modules and customize them per their own supply chain requirements. All delivery process has fully automated observability.

* **Multi-cloud/hybrid-environments app delivery as first-class citizen**: KubeVela natively supports multi-cluster/hybrid-cloud scenarios such as progressive rollout across test/staging/production environments, automatic canary, blue-green and continuous verification, rich placement strategy across clusters and clouds, along with automated cloud environments provision.


KubeVela was accepted as a CNCF Sandbox project on June 22, 2021.

* [Project Annual Review](https://github.com/cncf/toc/blob/main/reviews/2022-kubevela-annual.md)
* [Sandbox presentation](https://docs.google.com/presentation/d/1CWCLcsKpDQB3bBDTfdv2BZ8ilGGJv2E8L-iOA5HMrV0/edit#slide=id.p)
* [KubeVela Due Diligence for Incubation](https://docs.google.com/document/d/17owtXTes1TSkx0fwpV4_xRRrRL7D_jVN1c_iNsaRZ6M/edit)

## Progress Since Sandbox

### Community

The KubeVela devstats page and dashboards can be found [here](https://kubevela.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- According to devstats, KubeVela currently has [220](https://kubevela.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from [56](https://kubevela.devstats.cncf.io/d/5/companies-table?orgId=1)
companies. On average, there are [~2000 contributions per month](https://kubevela.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) contained within [150 merged PRs per month](https://kubevela.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=prs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).
- The community has grown significantly since the project entered the CNCF sandbox.
  - We held weekly community meetings constantly (total 84 as of July 2022). The meeting records can be found in [here](https://www.youtube.com/channel/UCSCTHhGI5XJ0SEhDHVakPAA/videos) and [here](https://space.bilibili.com/180074935/channel/seriesdetail?sid=1842207) for different timezone/language community members. The average number of meeting attendees is ~30.
  - Number of contributors: 90+ -> **220+**
  - Github stars: 1900+ -> **4000+**
  - Github forks: 250+ -> **560+**
  - Contributing organizations: 20+ -> **50+**
  - Last Release: **v1.5.0**
  - Member of DingTalk/Wechat group, and Slack channel: **4400+**

### Governance

The KubeVela community has matured significantly. We have established [the community membership roadmap](https://github.com/kubevela/community/blob/main/community-membership.md), and quite a few active and qualified contributors
have been promoted to approvers and maintainers. Currently, we have 6 maintainers and 8 approvers from 7 organizations.

We host weekly contributors and community meetings. We codified our [contributing guidelines](https://kubevela.io/docs/contributor/overview#contribution-guide), and [security processes](https://github.com/kubevela/kubevela/blob/master/SECURITY.md).

### Technology

We have significantly expanded the covered technical areas for the project development since KubeVela joined the Sandbox.
Currently, KubeVela has made feature developments in at least the following technical areas:
- **Multi-Cluster Delivery**: we have developed a project called [cluster-gateway](https://github.com/oam-dev/cluster-gateway) for routing kubernetes api traffic to multiple kubernetes clusters. It can support direct connect clusters using kubeconfig as secret, or integrate with [Open-Cluster-Management(OCM)](https://open-cluster-management.io/). As a result, the application layer of KubeVela shares the same interface no matter the underlying implementation.
- **Cloud Resource Delivery**: KubeVela natively support Crossplane for cloud resource provisioning, we have also developed a project called [terraform-controller](https://github.com/kubevela/terraform-controller) to make terraform-based cloud resource delivery available with Kubernetes CRD API. As a result, KubeVela can provide a wide range of cloud resources that supported by both Crossplane and Terraform ecosystem.
- **Lightweight App Delivery**: we have developed [velad](https://github.com/kubevela/velad) to provide lightweight KubeVela integrated with [K3s](https://k3s.io/) that can run in a single node with high availability. It can help local development or work as a lightweight control plane for multi-cluster/hybrid cloud app distribution.
- **UI Console**: we have built an end-to-end [UI console](https://github.com/kubevela/velaux) to help new users to quickly learn all features provided by KubeVela. It also help end users to adopt KubeVela easily as an out-of-box application platform.
- **Addon Ecosystem**: we have developed an addon ecosystem that can be used to extend the functionality of KubeVela. There's a growing catalog of addons available in [this repo](https://github.com/kubevela/catalog) including lots of CNCF projects.


## Incubation Stage Requirements

**_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope._**

Now there are at least 30+ adopters that have used KubeVela in production, full list of registered adopters can be found [here](https://github.com/kubevela/kubevela/issues/1662).

Highlighted examples include:
- [**Alibaba**](https://www.alibaba.com/): Using KubeVela as the core to deliver and manage applications across hybrid, multi-cluster environments.
- [**China Merchants Bank**](http://www.cmbchina.com/): Using KubeVela to build hybrid cloud-native application platform to unify the whole process from build, ship and run.
- [**4Paradigm**](https://www.4paradigm.com/): Using KubeVela to build a machine learning platform for AIOps.
- [**Bytedance**](https://www.bytedance.com/): Using OAM/KubeVela and Crossplane to provide advance Gaming PaaS abilities.
- [**Guidewire**](https://www.guidewire.com/): Using KubeVela to have the ability to extend kubernetes to be an ideal PaaS abstracting cloud provider and also the application centricity of deployment through workflow and an open spec.
- [**Li Auto**](https://www.lixiang.com/): Using KubeVela to build multi-cloud, multi-env and multi-cluster unified management application platform.
- [**Springer Nature**](https://www.springernature.com/): Using KubeVela to build a CloudFoundry like PaaS for internal usage.
- [**NetEase Games**](http://neteasegames.com/): Using KubeVela for game cluster deployment and istio app deployment.
- [**YOUZAN**](https://www.youzan.com/): Using KubeVela to build Cloud Native application management platform with rollout and multi-cluster capability.
- [**SHEIN**](https://www.shein.com/): Using KubeVela to help reduce custom resource workload docking time on building application delivery and management platform.


**_Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project._**

There are [51](https://github.com/orgs/kubevela/people) project members for KubeVela, out of which there are 6 maintainers, 8 approvers and 15 reviewers from **16** different organizations.
The OWNERS file can be found [here](https://github.com/kubevela/community/blob/main/OWNERS.md).


**_Demonstrate a substantial ongoing flow of commits and merged contributions._**

A short summary has been provided in the aforementioned community section. Detailed statistics can be found in the following devstats links:
- [Commit Activity](https://kubevela.devstats.cncf.io/d/74/contributions-chart?orgId=1)
- [New contributors since sandboxing](https://kubevela.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1573369200000&to=now)
- [Issues open/closed since sandboxing](https://kubevela.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1599721200000&to=now)
- [PRs open/closed since sandboxing](https://kubevela.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1599721200000&to=now)
- [Contributing organizations](https://kubevela.devstats.cncf.io/d/5/companies-table?orgId=1)

**_A clear versioning scheme._**

KubeVela uses the [semantic versioning scheme](https://semver.org/#semantic-versioning-specification-semver).

kubevela follows roughly once every two months release cadence with version numbers using format of MAJOR.MINOR.PATCH. The latest release is v1.5.0.

We have releases documented at: https://github.com/kubevela/kubevela/releases.

## Statement on Alignment with the CNCF Mission

The cloud-native ecosystem has been growing very fast, while modeling the deployment of a microservice application on top of cloud-native technology is a highly fragmented and challenging process. KubeVela aims to make deploying and operating applications across today's hybrid, multi-cloud environments easier, faster and more reliable. It will not only simplify the application delivery experience, while also be flexible enough to satisfy the fast growth of businesses requirements.

We believe it strictly aligns with the mission of CNCF in empowering organizations to build and run applications in modern cloud-native environments and benefit from the large growing cloud native technologies.


## Future plan

We are running our project in a fairly open model so that we can react to the needs of our community promptly. We are used to using [github millstone](https://github.com/kubevela/kubevela/milestone/9) to dynamically manage the issues on the plan.

There are still quite a few critical and high demanding features that are under development, including but not limited to the following:

- Improve the user experience for cloud resource provisioning and consuming with delivery workflow.
- Enhance the security for the whole CI/CD delivery process in hybrid/multi-cluster scenario.
- Provide user defined observability with more automation.

 In addition, we have [roadmap](https://kubevela.io/docs/roadmap/) page on the doc site so that the community members can easily track where we are going in a wider perspective.
