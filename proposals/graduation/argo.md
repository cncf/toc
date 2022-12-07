# Argo Graduation Proposal

Since joining the CNCF in April 2020 as an incubation-level project, the Argo Project has made great progress in growing the community of contributors and users including the addition of Red Hat and Codefresh along-side BlackRock as partners in governing and running the project. On behalf of the Argo Project team, we believe that Argo is ready for graduation.

## Description

The Argo Project is a set of Kubernetes-native tools for deploying and running jobs and applications on Kubernetes. All the Argo tools are implemented as controllers and custom resources. These tools can be used independently but are even more powerful when used together. As a CNCF graduated project, all top level projects must meet graduation maturity requirements.

**Argo Workflows** enables creation of complex parallel workflows as Kubernetes resources and is used in many different use cases from CI/CD pipelines to DAG-based ML training workflows. It is the workflow engine behind the open source Kubeflow Pipelines.

**Argo Events** provides declarative management of event-based dependencies and triggers for Kubernetes resources based on various event sources. A common use of Argo Events is to trigger Argo Workflows and to generate events for long-lived services deployed using Argo CD.

**Argo CD** supports declarative GitOps-based deployment of any Kubernetes resource, including Argo Events, services and deployments across multiple k8s clusters.

**Argo Rollouts** provides declarative progressive delivery of application resources using deployment strategies such as blue/green and canary for Kubernetes.

Each of these tools is built to be Kubernetes native and are implemented as controllers and Custom Resources.  These tools can be used independently, but there is great benefit in using them together to create complex applications. Argo Events, for example, can kick off Argo Workflows which can generate and queue additional Argo Events which are processed by long-lived services deployed using Argo CD that use Argo Rollouts to control the deployment process.

Argo has been presented at Kubernetes community meetings, KubeCon conferences, and many other venues by both project members and the community at large.

* https://indico.cern.ch/event/950886/attachments/2122812/3573279/20201014_CERN_ZeroToKubernetesPhysicsAnalysis.pdf
* https://medium.com/arthur-engineering/picking-a-kubernetes-orchestrator-airflow-argo-and-prefect-83539ecc69b
* https://www.cncf.io/wp-content/uploads/2020/12/Argo_-Real-Enterprise-scale-with-Kubernetes.pdf
* KubeCon: https://www.youtube.com/watch?v=yeVkTTO9nOA
* KubeCon: https://www.youtube.com/watch?v=OdzH82VpMwI&feature=youtu.be
* KubeCon: https://www.youtube.com/watch?v=ZK510prml8o
* KubeCon: https://www.youtube.com/watch?v=VXrGp5er1ZE&t=0s&index=135&list=PLj6h78yzYM2PZf9eA7bhWnIh_mK1vyOfU
* Sig Apps: https://www.youtube.com/watch?v=aWDIQMbp1cc
* K8s Community Meeting: https://www.youtube.com/watch?v=LRDoKOLOlf8
* K8s Community Meeting: https://www.youtube.com/watch?v=GeB50xG-gmc&feature=youtu.be&t=81
* TGIK by Joe Beda: https://www.youtube.com/watch?v=M_rxPPLG8pU&start=859

## Statement on alignment with CNCF mission

The Argo project is well-aligned with the CNCF’s mission to make cloud native computing ubiquitous. We are completely aligned to "empower organizations to build and run scalable applications" including the adoption of Kubernetes and declarative APIs. All the Argo tools are implemented as controllers and custom resources. They use/integrate with other CNCF projects like gRPC, Prometheus, NATS, Helm..

#### Sponsors / Advisors from TOC
* Lei Zhang
* Cathy Zhang

#### Project name
* Argo

#### Unique identifier
* argo, argoproj

#### Preferred maturity level
* graduated

#### The Argo Community is looking for the following by becoming a CNCF graduated project:

* Promote public visibility to add value to the CNCF mission.
* Foster collaborative development to further the overall CNCF mission.
* Closer integration and collaboration with other CNCF projects.
* Reassurance for the Argo Community that the project will be a permanent part of the CNCF ecosystem well into the future.

#### License
* Apache License 2.0

#### Source control repositories

* https://github.com/argoproj/argo-workflows
* https://github.com/argoproj/argo-events
* https://github.com/argoproj/argo-cd
* https://github.com/argoproj/argo-rollouts

#### External Dependencies

Argo depends on the following external software components:

* Kubernetes (Apache Software License 2.0)
* Dex (Apache Software License 2.0)
* Redis (BSD)
* React (MIT)
* GRPC (Apache Software License 2.0)
* Golang (Apache Software License 2.0)

#### Initial Committers

* Alex Collins (Intuit)
* Alexander Matyushentsev (Akuity)
* Jann Fischer (Red Hat)
* Jesse Suen (Akuity)
* Jonathan West (Red Hat)
* Oleg Sucharevich (Elastic)
* Vaibhav Page (BlackRock)

#### Infrastructure requests

* None

#### Communication Channels

* Slack: https://cncf.slack.com/ #argo-cd #argo-workflow #argo-rollouts #argo-events
* Mailing List: https://groups.google.com/forum/#!forum/argoproj 
* Community Docs: https://github.com/argoproj/argoproj/blob/master/community/README.md

#### Issue tracker

* https://github.com/argoproj/argo-workflows/issues 
* https://github.com/argoproj/argo-events/issues
* https://github.com/argoproj/argo-cd/issues
* https://github.com/argoproj/argo-rollouts/issues

#### Website
* https://argoproj.github.io/ 

#### Release methodology and mechanics

* Major feature release about four times per year with additional minor/patch releases
* Some users also run HEAD of master particularly for trying and testing release candidates.

#### Social media accounts

* Twitter: @argoproj

#### Existing sponsorship
* CNCF, Intuit, Akuity, BlackRock, Codefresh and Red Hat are the sponsors of Argo

#### Community size

Argoproj

* 25,000 stars
* 14,000 slack members
* 6,100 forks
* 815 contributors
* 240 end user companies and organizations

#### Production usage

Argo is known to be actively used in production by the following organizations:

* https://github.com/argoproj/argo-workflows/blob/master/USERS.md
* https://github.com/argoproj/argo-events/blob/master/USERS.md
* https://github.com/argoproj/argo-cd/blob/master/USERS.md
* https://github.com/argoproj/argo-rollouts/blob/master/USERS.md

## Graduation State Criteria

#### Have committers from at least two organizations

The Argo governance committee consists of members from Intuit, Red Hat, Blackrock, and Codefresh. Committers are represented from a much broader set of organizations.

#### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/)

Core Infrastructure Initiative Best Practices Badges have been completed and are being maintained.
* https://bestpractices.coreinfrastructure.org/projects/4486
* https://bestpractices.coreinfrastructure.org/projects/3834
* https://bestpractices.coreinfrastructure.org/projects/3830
* https://bestpractices.coreinfrastructure.org/projects/3832

#### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation

Third party security audit has been completed with Trail of Bits and ADA Logics as the 3rd party auditors.
https://github.com/argoproj/argoproj/blob/master/docs/argo_threat_model.pdf
https://github.com/argoproj/argoproj/blob/master/docs/argo_security_final_report.pdf
https://github.com/argoproj/argoproj/blob/master/docs/audit_fuzzer_adalogics_2022.pdf
https://github.com/argoproj/argoproj/blob/master/docs/argo_security_audit_2022.pdf

#### Explicitly define a project governance and committer process

* https://github.com/argoproj/argoproj/blob/master/community/README.md
* https://github.com/argoproj/argoproj/blob/master/community/GOVERNANCE.md
* https://github.com/argoproj/argoproj/blob/master/community/membership.md

#### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website)

* https://github.com/argoproj/argo-workflows/blob/master/USERS.md
* https://github.com/argoproj/argo-events/blob/master/USERS.md
* https://github.com/argoproj/argo-cd/blob/master/USERS.md
* https://github.com/argoproj/argo-rollouts/blob/master/README.md
