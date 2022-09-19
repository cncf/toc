# Volcano CNCF Sandbox Project Proposal

## Name of project:

Volcano

## Description (what it does, why it is valuable, origin and history)

More and more users would like to run "batch" workload, e.g. Tensorflow, Spark, on Kubernetes. Several domain communities, e.g. kubeflow, spark-operator, focus on frameworks enhancement, and requires enhancement to the orchestration layer, such as: 

- Scheduling: gang-scheduling, fair-sharing, reservation/backfill ……
- Job/Queue management: hierarchical queue, indexed job, multiple pod templates, job dependency ……
- Data management: data locality, cache, data aware scheduling …
- Accelerator, e.g. GPU, FPGA ……
- Others: Singularity, Round-trip, Throughput …...

Volcano is a batch system built on Kubernetes for the above requirements. It provides a suite of mechanisms that
are commonly required by many classes of batch & elastic workload including: machine learning/deep learning,
bioinformatics/genomics and other "big data" applications. These types of applications typically run on generalized
domain frameworks like TensorFlow, Spark, PyTorch, MPI, etc, which Volcano integrates with.

## Statement on alignment with CNCF charter mission

**Volcano** aligns closely with the Cloud Native Computing Foundation (CNCF) mission as described in section 1 of the CNCF Charter; and we agree with the CNCF philosophy/vision, and would like to contribute.

More and more user would like to run "batch" workload (e.g. Spark, Tensorflow) in cloud by Kubernetes; and those batch workload requires several enhancements. **Volcano** provides a suite of mechanisms, e.g. scheduling, job/queue management, data management, to help those batch workload onboard to the kubernetes smoothly. **Volcano** provides both different layers of API which is used as a foundation for other cloud native projects, e.g. kubeflow, spark-operator.

## Sponsor from TOC:


## Preferred maturity level:

Sandbox

## License

Apache License 2.0

## Source control

http://github.com/volcano-sh/volcano

## External dependencies (including licenses)

- github.com/golang/glog (Apache License 2.0)
- github.com/onsi/ginkgo (MIT License)
- github.com/onsi/gomega (MIT License)
- github.com/spf13/cobra (Apache License 2.0)
- github.com/spf13/pflag (BSD 3-Clause "New" or "Revised" License)
- k8s.io/client-go (Apache License 2.0)
- k8s.io/kubernetes (Apache License 2.0)
- k8s.io/apimachinery (Apache License 2.0)
- k8s.io/api (Apache License 2.0)
- k8s.io/apiserver (Apache License 2.0)
- k8s.io/apiextensions-apiserver (Apache License 2.0)
- k8s.io/code-generator (Apache License 2.0)
- golang.org/x/crypto (Go's License)
- k8s.io/code-generator (Apache License 2.0)
- github.com/hashicorp/go-multierror (Mozilla Public License 2.0)

## Initial committers

* Klaus Ma (@k82cn, Huawei)
* Zhonghu Xu (@hzxuzhonghu, Huawei)
* Quinton Hoole (@quinton-hoole, Huawei)
* Animesh Singh (@animeshsingh, IBM)
* Jun Gong (@hex108, Tencent)

## Infrastructure requests (CI / CNCF Cluster)

CI (currently use TravisCI)

## Communication channels (slack, irc, mailing lists)

* Mailing lists: https://groups.google.com/forum/#!forum/volcano-sh
* Slack: https://volcano-sh.slack.com

## Issue tracker (GitHub by default)

http://github.com/volcano-sh/volcano

## Website (current version will move to project.cncf.io, see here for guidelines)

http://volcano.sh

## Release methodology and mechanics

Releases are done through Travis CI; container images build there and pushed to Docker Hub.

## Social media accounts

* https://twitter.com/volcano_sh

## Community size and any existing sponsorship

_(Dec 16, 2019)_

- 683 Stars
- 154 Forks
- 74 Contributors
- 2 Releases
- Integrations: 
  - [Spark-Operator](https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/blob/master/docs/volcano-integration.md)
  - [Kubeflow/tf-operator](https://www.kubeflow.org/docs/use-cases/job-scheduling/)
  - [Kubeflow/arena](https://github.com/kubeflow/arena/blob/master/docs/userguide/12-volcanojob.md)
  - [Cromwell](https://github.com/broadinstitute/cromwell/blob/develop/docs/backends/Volcano.md)

## Adopters:

https://github.com/volcano-sh/volcano/blob/master/docs/community/adopters.md

| Organization | Contact | Environment | Description of Use |
| ------------ | ------- | ----------- | ------------------ |
| [Unisound](https://www.unisound.com/) |[@xieydd](https://github.com/xieydd)| Evaluation | Evaluation in ATLAS AI Platform |
| [BIBDR](http://www.bibdr.org/en/) |[@felix5572](https://github.com/felix5572)| Evaluation | Scientific calculations in physics, materials , biology and chemistry. molecular dynamics simulation. |
| [caicloud](https://caicloud.io/) |[@gaocegege](https://github.com/gaocegege)| Evaluation | Scheduler for Distributed DL training Jobs |
| [Baidu](https://baidu.com/) |[@tizhou86](https://github.com/tizhou86)| Testing | Scheduler for Deep Learning Platform to Optimize Performance |
| [GrandOmics](https://www.grandomics.com/) |[@alartin](https://github.com/alartin)| Evaluation | Infrastructure of Hanwell (Huawei Cloud backend of Cromwell which is a Broad Institute implementation of WDL) |
| [Huawei Cloud](https://huaweicloud.com/) |[@tsjsdbd](https://github.com/tsjsdbd)| **Production** | Scheduler & Job Management of AI Container Service and CCI  |
| [JD Retail Infrastructure Department](https://jd.com/) |[@yuanchen8911](https://github.com/yuanchen8911)| Evaluation | Spark on K8S  |
| [kt NexR](https://www.ktnexr.com) |[@minyk](https://github.com/minyk), [@dieselnexr](https://github.com/dieselnexr)| Evaluation | spark scheduler of our next cloud native product. |

## Project logo in svg format

[Volcano SVG logo](https://raw.githubusercontent.com/volcano-sh/volcano/master/docs/images/volcano-logo-svg.svg)
