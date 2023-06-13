# Kubeflow CNCF Incubator Project Proposal

## Name of Project:

Kubeflow

## Description

Kubeflow project is an extendable Machine Learning (ML) platform made of multiple distinct components to address specific stages of the ML lifecycle, including:

* **Mode Development** - includes services to create and manage interactive Jupyter notebooks. Users can customize their notebook deployment and their compute resources to suit their data science needs. Supports local workflows with cloud deployment when ready.
* **Model Training**
    * **Training Operator** - custom TensorFlow training job operator that can be used to train ML models. Supports distributed jobs. Configures the training controller to use CPUs or GPUs and to suit various cluster sizes.
    * **Pipelines** - comprehensive solution for deploying and managing end-to-end ML workflows. Supports rapid and reliable experimentation. Allows users to schedule and compare runs, and examine detailed reports on each run.
* **Multi-framework** - supports models beyond TensorFlow, including PyTorch, Apache MXNet, MPI, XGBoost, Chainer, and more. Integrates with Istio and Ambassador for ingress, Nuclio as a fast multi-purpose serverless framework, and Pachyderm for managing your data science pipelines.

## Statement on alignment with CNCF mission

Since its inception in 2017, when Kubeflow was first [open sourced by Google](https://kubernetes.io/blog/2017/12/introducing-kubeflow-composable/), the project strove to make ML on Kubernetes simple, portable and scalable. In May 2020, **with the [v1.0](https://blog.kubeflow.org/releases/2020/03/02/kubeflow-1-0-cloud-native-ml-for-everyone.html)** release, Kubeflow reached maturity across a core set of its stable applications, and graduated Kubeflow Serving as an independent project ([KServe](https://github.com/kserve/kserve)), which is now incubating in [Linux Foundation AI & Data](https://lfaidata.foundation/) foundation. 

Today, the community has close to 200 contributors from over 30 organizations, and hosts regular contributor summits, community meetings, and meetups across the world. The broader Kubeflow ecosystem also includes today a number [distributions](https://www.kubeflow.org/docs/distributions/) across multiple cloud service providers and on-prem environments. Kubeflow’s powerful development experience helps data scientists build, train and deploy their ML models while also enabling ML operation teams in the enterprise to deploy and scale advanced workflows in a variety of infrastructures. 

To meet the evolving needs of its community, Kubeflow needs a vendor-neutral foundation. This application to become a CNCF incubating project is the next big milestone for the Kubeflow community, and we believe that CNCF is the right foundation for the project given Kubeflow's roots in the Kubernetes community and the Cloud native ecosystem (see: Kubernetes, Argo, Istio, Knative, Kustomize). Being part of CNCF will bring Kubeflow closer to this project ecosystem, with the potential to increase the velocity of feedback and innovation over time. In turn, Kubeflow can bring years of community expertise in enabling the use of containers for running AI/ML workloads.


## Roadmap

* Kubeflow - [https://github.com/kubeflow/kubeflow/blob/master/ROADMAP.md](https://github.com/kubeflow/kubeflow/blob/master/ROADMAP.md)
* Pipelines - [https://github.com/kubeflow/pipelines/blob/master/ROADMAP.md](https://github.com/kubeflow/pipelines/blob/master/ROADMAP.md)
* Katib - [https://github.com/kubeflow/katib/blob/master/ROADMAP.md](https://github.com/kubeflow/katib/blob/master/ROADMAP.md)
* Training Operator - [https://github.com/kubeflow/training-operator/blob/master/docs/roadmap.md](https://github.com/kubeflow/training-operator/blob/master/docs/roadmap.md)

## Sponsors from TOC

[Ricardo Rocha](https://github.com/rochaporto)

## Preferred maturity level

Incubating

## License

Kubeflow is licensed under the Apache 2.0 license.

## Source control

[https://github.com/kubeflow](https://github.com/kubeflow) (everything in the Kubeflow org)


## External Dependencies

Kubeflow has external dependencies licensed under the following CNCF-approved licenses:

* Apache 2.0
* BSD License
* ISC License
* MIT License

The dependency for each sub-project and their license information are included in:

* Kubeflow - [https://github.com/kubeflow/kubeflow/search?q=in%3Apath+third_party+license](https://github.com/kubeflow/kubeflow/search?q=in%3Apath+third_party+license)
* Pipelines - [https://github.com/kubeflow/pipelines/tree/master/third_party](https://github.com/kubeflow/pipelines/tree/master/third_party)
* Training-operator - [https://github.com/kubeflow/training-operator/tree/master/third_party_licenses](https://github.com/kubeflow/training-operator/tree/master/third_party_licenses)
* Katib - [https://github.com/kubeflow/katib/blob/master/go.mod](https://github.com/kubeflow/katib/blob/master/go.mod)
* MPI Operator - [https://github.com/kubeflow/mpi-operator/blob/master/go.mod](https://github.com/kubeflow/mpi-operator/blob/master/go.mod)
When possible, sub-projects in Kubeflow will be brought in line with CNCF guidelines before submission or, in extreme circumstances, the Kubeflow project will seek CNCF governing board exceptions.

## Maintainers

Kubeflow welcome commits from anyone in the community. Due to its multi-repo structure with distinct contributor roles, the project uses "maintainer" as the noun for a member with sufficient permissions to merge PR in specific repo. That role can be defined at any level of each repo using the OWNERS file which in turn grants that contributor permission to “/LGTM” and “/APPROVE” PRs. These comments on PR subsequently trigger automation that actually merges that PR. 

The current snapshot of the maintainers as defined above (company affiliation based on user public GitHub profile). 

* Kubeflow
    * [pdmack](https://github.com/pdmack) - Pete MacKinnon (NVIDIA)
    * [james-jwu](https://github.com/james-jwu) - James Wu (Google)
    * [zijianjoy](https://github.com/zijianjoy) - James Liu (Google)
* Kubeflow Notebooks & Common Kubeflow Components
    * [elikatsis](https://github.com/elikatsis) - Ilias Katsakioris (Arrikto)
    * [kimwnasptd](https://github.com/elikatsis) - Kimonas Sotirchos (Arrikto)
    * [StefanoFioravanzo](https://github.com/StefanoFioravanzo) - Stefano Fioravanzo (Arrikto)
    * [thesuperzapper](https://github.com/thesuperzapper) - Mathew Wicks
    * [yanniszark](https://github.com/yanniszark) - Yannis Zarkadas (Arrikto)
* Kubeflow Manifests
    * [elikatsis](https://github.com/elikatsis) - Ilias Katsakioris (Arrikto)
    * [kimwnasptd](https://github.com/elikatsis) - Kimonas Sotirchos (Arrikto)
    * [PatrickXYS](https://github.com/PatrickXYS) - Yao Xiao (Google)
    * [StefanoFioravanzo](https://github.com/StefanoFioravanzo) - Stefano Fioravanzo (Arrikto)
    * [yanniszark](https://github.com/yanniszark) - Yannis Zarkadas (Arrikto)
* Pipelines
    * [Animeshsingh](https://github.com/animeshsingh) - Animesh Singh (IBM)
    * [chensun](https://github.com/chensun) - Chen Sun (Google)
    * [connor-mccarthy](https://github.com/connor-mccarthy) - Connor McCarthy (Google)
    * [gkcalat](https://github.com/gkcalat) - Ablai Akhazhanov (Google)
    * [IronPan](https://github.com/IronPan) - Yang Pan (Google)
    * [james-jwu](https://github.com/james-jwu) - James Wu (Google)
    * [jlyaoyuli](https://github.com/jlyaoyuli) - Joe Li (Google)
    * [Linchin](https://github.com/Linchin) - Lingqing Gan (Google)
    * [zijianjoy](https://github.com/zijianjoy) - James Liu (Google)
* Training-operator
    * [Gaocegege](https://github.com/Gaocegege) - Ce Gao (TensorChord)
    * [Jeffwan](https://github.com/Jeffwan) - Jiaxin Shan (Bytedance)
    * [Johnugeorge](https://github.com/Johnugeorge) - Johnu George (Nutanix)
    * [Terrytangyuan](https://github.com/Terrytangyuan) - Yuan Tang (Akuity)
    * [zw0610](https://github.com/zw0610) - Wang Zhang (Tencent)
* Katib
    * [Andreyvelich](https://github.com/Andreyvelich) - Andrey Velichkevich (Apple)
    * [Gaocegege](https://github.com/Gaocegege) - Ce Gao (TensorChord)
    * [Johnugeorge](https://github.com/Johnugeorge) - Johnu George (Nutanix)
    * [tenzen-y](https://github.com/tenzen-y) - Yuki Iwai (CyberAgent)
* MPI Operator
    * [Terrytangyuan](https://github.com/Terrytangyuan) - Yuan Tang (Akuity)
    * [Alculquicondor](https://github.com/terrytangyuan) - Aldo Culquicondor (Google)
    * [tenzen-y](https://github.com/tenzen-y) - Yuki Iwai (CyberAgent)
## Infrastructure requests (CI/CNCF Cluster)

Kubeflow uses AWS for its CI needs. The community is already working on streamlining the infrastructure provisioning and operations ([https://github.com/kubeflow/testing/issues/1006](https://github.com/kubeflow/testing/issues/1006)). If Kubeflow does get accepted for incubation in CNCF, the community will work towards migrating the ownership of this infrastructure to the CNCF.

Each distribution offering Kubeflow as a service operates their own test/build/release infrastructure. These will continue being the responsibility of each provider and will not transition to CNCF ownership. The maintainers of these infrastructures are available in the platform-specific Slack channels:

* AWS - [https://kubeflow.slack.com/archives/CKBA5D0MU](https://kubeflow.slack.com/archives/CKBA5D0MU)
* Azure - [https://kubeflow.slack.com/archives/CUW6SLCPR](https://kubeflow.slack.com/archives/CUW6SLCPR)
* Platform - [https://kubeflow.slack.com/archives/CKH7V1M7F](https://kubeflow.slack.com/archives/CKH7V1M7F)

## Communication Channels

The Kubeflow community includes a variety of Slack forums for general discussion, feature requests, job postings, and sub-project, platform, or geo-specific channels. Popular Kubeflow Slack channels are listed here - [https://www.kubeflow.org/docs/about/community/#kubeflow-slack](https://www.kubeflow.org/docs/about/community/#kubeflow-slack)


## Issue tracker

* Kubeflow - [https://github.com/kubeflow/kubeflow/issues](https://github.com/kubeflow/kubeflow/issues)
* Pipelines - [https://github.com/kubeflow/pipelines/issues](https://github.com/kubeflow/pipelines/issues)
* Training-operator -  [https://github.com/kubeflow/training-operator/issues](https://github.com/kubeflow/training-operator/issues)
* Katib - [https://github.com/kubeflow/katib/issues](https://github.com/kubeflow/katib/issues)
* MPI Operator - [https://github.com/kubeflow/mpi-operator/issues](https://github.com/kubeflow/mpi-operator/issues)

## Community

* Repo - [https://github.com/kubeflow/community](https://github.com/kubeflow/community)
* Website - [https://www.kubeflow.org/docs/about/community/](https://www.kubeflow.org/docs/about/community/)
* Workgroups & SIGs - [https://github.com/kubeflow/community/blob/master/wg-list.md](https://github.com/kubeflow/community/blob/master/wg-list.md)
* Kubeflow Blog - [https://blog.kubeflow.org](https://blog.kubeflow.org)

## Code of Conduct

[https://github.com/kubeflow/community/blob/master/CODE_OF_CONDUCT.md](https://github.com/kubeflow/community/blob/master/CODE_OF_CONDUCT.md)

## Social media accounts

* Youtube - [https://www.youtube.com/kubeflow](https://www.youtube.com/kubeflow)
* Twitter - [https://twitter.com/kubeflow](https://twitter.com/kubeflow)
* LinkedIn - [https://www.linkedin.com/company/kubeflow](https://www.linkedin.com/company/kubeflow)
* Google Groups - [https://groups.google.com/g/kubeflow-discuss](https://groups.google.com/g/kubeflow-discuss)
* Google Workspace (used to manage community calendar, Drive, etc.) 

## Existing sponsorship

None

## Adopters

* Kubeflow - [https://github.com/kubeflow/kubeflow/network/dependent](https://github.com/kubeflow/kubeflow/network/dependents)
* Pipelines - [https://github.com/kubeflow/pipelines/network/dependents](https://github.com/kubeflow/kubeflow/network/dependents)
* Training-operator -  [https://github.com/kubeflow/training-operator/network/dependents](https://github.com/kubeflow/training-operator/network/dependents)
* Katib - [https://github.com/kubeflow/katib/network/dependents](https://github.com/kubeflow/katib/network/dependents)
* MPI Operator - [https://github.com/kubeflow/mpi-operator/network/dependents](https://github.com/kubeflow/mpi-operator/network/dependents)

## Releases

* The release handbook - [https://github.com/kubeflow/community/blob/master/releases/handbook.md](https://github.com/kubeflow/community/blob/master/releases/handbook.md)
* Kubeflow releases - [https://github.com/kubeflow/community/tree/master/releases](https://github.com/kubeflow/community/tree/master/releases)

## Community size

The Kubeflow project has now close to 200 contributors from over 30 organizations, and the Kubeflow community has hosted a number of summits and contributor meetups across the world. The broader Kubeflow ecosystem includes a number distributions across multiple cloud service providers and on-prem environments. For more information see [https://www.kubeflow.org/docs/about/community/](https://www.kubeflow.org/docs/about/community/)

The Kubeflow project repo-level metrics: 

* [Kubeflow](https://github.com/kubeflow/kubeflow) - 11.9k stars, 2k forks 
* [Pipelines](https://github.com/kubeflow/pipelines) - 3k stars, 1.3k forks 
* [Training-operator](https://github.com/kubeflow/training-operator) - 1.2k stars, 500 forks 
* [Katib](https://github.com/kubeflow/katib) - 1.2k stars, 362 forks 
* [MPI Operator](https://github.com/kubeflow/mpi-operator) - 318 stars, 190 forks