# KServe CNCF Incubator Project Proposal

## Name of Project:

KServe

## Description

KServe is a community driven open source project, aiming to deliver a cloud-native, scalable, extensible serverless ML inference platform. It provides an open standard control and data plane for serving production ML models including Scikit-learn, Tensorflow, PyTorch and Hugging face Transformer/large language models(LLMs).
It encapsulates the complexity of autoscaling, networking, health checking, and server configuration to bring cutting edge serving features like GPU Autoscaling, Scale to Zero, and Canary Rollouts to your ML deployments. It enables a simple, pluggable, and complete story for Production ML Serving including prediction, pre-processing, post-processing and explainability. 
Features:

KServe is a standard, cloud agnostic Model Inference Platform for serving predictive and generative AI models on Kubernetes, built for highly scalable use cases.
Provides performant, standardized inference protocol across ML frameworks including OpenAI specification for generative models.
Support modern serverless inference workload with request based auto-scaling including scale-to-zero on CPU and GPU.
Provides high scalability, density packing and intelligent routing using ModelMesh.
Simple and pluggable production serving for inference, pre/post processing, monitoring and explainability.
Advanced deployments for canary rollout, pipeline, ensembles with InferenceGraph.

Statement on alignment with CNCF mission
Since its inception in 2019, previously known as KFServing under Kubeflow, the KServe project has gained significant traction with support from leading industry players such as Bloomberg, IBM, RedHat, Intuit, SAP, AliCloud. It has fostered a diverse and inclusive community including both end users and cloud vendors. Embracing an open governance model, KServe has established a clear process for promoting reviewers and approvers, ensuring transparency and fairness in decision-making.

As of Mar 2024, the community has 233 contributors from over 30 organizations, and hosts community meetings, and meetups across the world. KServe's powerful development experience helps data scientists, ML engineers deploy their ML models in the Kubernetes environments on production. 

The application to become a CNCF incubating project is the next big milestone for the KServe community, and we believe that CNCF is the right foundation for the project given KServe’s roots in the Kubernetes community and the Cloud native ecosystem (see: Kubernetes, Istio, Knative, OpenTelemetry, gRPC, Buildpack). Being part of CNCF will bring KServe closer to this cloud native ecosystem, with the potential to increase the velocity of feedback and innovation over time. In turn, KServe can bring years of community expertise in enabling deploy AI/ML workloads on production.


## Roadmap

[KServe Roadmap](https://github.com/kserve/kserve/blob/master/ROADMAP.md)

## Sponsors from TOC

Ricardo Arvaena

## Preferred maturity level

Incubating

## License

KServe is licensed under the Apache 2.0 license.

## Source control

[https://github.com/kserve](https://github.com/kserve) (everything in the kserve org)

## External Dependencies

KServe has external dependencies licensed under the following CNCF-approved licenses:

* Apache 2.0
* BSD License
* ISC License
* MIT License


## Maintainers and Committers

KServe welcomes commits from anyone in the community. The project uses "maintainer" as the noun for a member with sufficient permissions to merge PR in specific repo. That role can be defined at any level of each repo using the OWNERS file which in turn grants that contributor permission to “/LGTM” and “/APPROVE” PRs. These comments on PR subsequently trigger automation that actually merges that PR. 

The current snapshot of the maintainers as defined above (company affiliation based on user public GitHub profile). 

### Maintainers:

- [njhill](https://github.com/njhill) - Nick Hill (IBM)
- [rachitchauhan43](https://github.com/rachitchauhan43) - Rachit Chauhan(Intuit)
- [yuzisun](https://github.com/yuzisun) - Dan Sun(Bloomberg)

### Committers:

- [andyi2it] - Andrews Arokiam (Idea2IT)
- [alexagriffith] - Alexa Griffith (Bloomberg)
- [ckadner] - Christian Kadner (IBM)
- [cmaddalozzo] - Curtis Maddalozzo (Bloomberg)
- [israel-hdez] - Edgar Hernandez (Redhat)
- [lizzzcai] - Lize Cai (SAP)
- [sivanantha321] - Sivanantham (Idea2IT)
- [terrytangyuan] - Yuan Tang (Redhat)

Infrastructure requests (CI/CNCF Cluster)

KServe uses Github Action for its CI needs. If KServe does get accepted for incubation in CNCF, the community will work towards migrating the ownership of this infrastructure to the CNCF.

https://github.com/kserve/kserve/tree/master/.github

## Communication Channels

The KServe community includes a variety of Slack forums for general discussion, feature requests, job postings, and sub-project, platform, or geo-specific channels. Popular KServe Slack channels are listed here 
[Kubeflow kserve channel](https://kubeflow.slack.com/archives/CH6E58LNP)
[CNCF kserve channel](https://cloud-native.slack.com/archives/C06AH2C3K8B)

## Issue tracker

- KServe - [https://github.com/kserve/kserve/issues](https://github.com/kserve/kserve/issues)
- ModelMesh - [https://github.com/kserve/modelmesh-serving](https://github.com/kserve/modelmesh-serving)

## Community

- Community Github Repo - [https://github.com/kserve/community](https://github.com/kserve/community)
- KServe Blog - [https://github.com/kserve/website/tree/main/docs/blog/articles](https://github.com/kserve/website/tree/main/docs/blog/articles)

## Code of Conduct

## Adopters

- KServe Python SDK: [https://github.com/kserve/kserve/network/dependents](https://github.com/kserve/kserve/network/dependents)
- KServe: [https://github.com/kserve/website/blob/main/docs/community/adopters.md](https://github.com/kserve/website/blob/main/docs/community/adopters.md)

## Security Process:

Vulnerability issues are scanned at regular intervals
Reporting security issues: https://github.com/kserve/kserve/blob/master/SECURITY.md

## Releases

The release handbook - [https://github.com/kserve/kserve/blob/master/release/RELEASE_PROCESS_v2.md](https://github.com/kserve/kserve/blob/master/release/RELEASE_PROCESS_v2.md)

## Community size

The KServe project now has 233 contributors from over 30 organizations. Vendors like Redhat, IBM, Alicloud and end users like Bloomberg, Intuit have been actively engaged with the project and the community. 

The KServe community has been hosting the biweekly US/EU working group with 20-30 attendees. For more information see [https://github.com/kserve/community](https://github.com/kserve/community)

The KServe github organization has more than 3k stars and 1k+ forks across all the repositories

* [KServe](https://github.com/kserve/kserve)
* [ModelMesh Serving](https://github.com/kserve/modelmesh-serving)
* [ModelMesh](https://github.com/kserve/modelmesh)
* [Open Inference Protocol](https://github.com/kserve/open-inference-protocol)
* [Website](https://github.com/kserve/website)

