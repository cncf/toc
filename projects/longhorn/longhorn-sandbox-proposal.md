=== Longhorn CNCF Sandbox Project Proposal

*Name of Project:* Longhorn

*Project Description:*

Longhorn is an open source distributed block storage software for Kubernetes. It is a lightweight and portable implementation of persistent storage that works for any Kubernetes cluster. The project implements distributed block device volumes that can be mounted as read-write by a single node (ReadWriteOnce).

Longhorn is designed to leverage the existing Linux technologies as much as possible, rather than building a complex storage technology stack from scratch. The software takes advantage of modern high-speed and high capacity SSD and NVMe storage, and proven Linux storage features like sparse files and cgroups.

The most distinct feature of Longhorn is to implement each volume as an independent microservice. By leveraging Kubernetes to orchestrate the volumes, it implements a highly feature-rich enterprise-grade distributed block storage system with about 30K line of Golang code.

Here are the major features of Longhorn:

* Enterprise-grade distributed storage software
* Incremental snapshot of the block storage
* Backup to an off-cluster backup store (NFS or S3-compatible object storage), built on efficient change block detection mechanism
* Recurring snapshot and backup
* Cross-cluster disaster recovery volume with defined Recovery Time Objective (RTO) and Recovery Point Objective (RPO)
* Automated non-disruptive upgrade. You can upgrade the entire Longhorn software stack without disrupting running volumes.
* Intuitive GUI

**Architecture**

Longhorn consists of the following major components:

***Engine and Replicas***

The Longhorn engine is the microservice used to implement each volume. The engine implements the data plane of reading and writing the block devices. Each Longhorn volume is implemented using one engine and multiple replicas. The engine always runs on the same node as the pod consuming the volume, whereas Longhorn replicas reside on different nodes to ensure redundancy.

Longhorn delivers added resiliency because the data plane for each volume is separate. A software bug that causes one volume to malfunction does not impact other volumes in the system.

***Manager***

Longhorn manager is built on top of Kubernetes. It implements CSI driver for Longhorn and leverages Kubernetes to ensure the proper functioning of all engines and replicas. The manager relies on the Kubernetes controller pattern to manage the volume status. It also utilizes Kubernetes to implement pod orchestration, cron job, node failure detection, pod failure detection, log collection, etc.

***Other Components***

Besides engine and manager, Longhorn also has an intuitive GUI and a test suite which covers every major feature of Longhorn. Integration test cases are required for new features and major bug fixes.

*Statement on alignment with CNCF mission:*

Longhorn project aligns with CNCF's mission to make cloud-native computing ubiquitous. Today, persistent storage support varies from Kubernetes cluster to Kubernetes cluster. As a result, users cannot deploy stateful workloads that require persistent volumes and expect their stateful workloads to work on any Kubernetes cluster. By creating a portable and lightweight implementation of persistent storage volumes, Longhorn can make persistent storage support available on every Kubernetes cluster and thus can help drive the adoption of Kubernetes technology.

*Sponsors from TOC:* Alexis Richardson and Liz Rice

*Preferred maturity level:* Sandbox

*License:* Apache License v2.0

*Source control:* GitHub

Longhorn repositories include:

* https://github.com/longhorn/longhorn
* https://github.com/longhorn/longhorn-engine
* https://github.com/longhorn/longhorn-instance-manager
* https://github.com/longhorn/longhorn-manager
* https://github.com/longhorn/longhorn-ui
* https://github.com/longhorn/longhorn-tests
* https://github.com/longhorn/liblonghorn
* https://github.com/longhorn/go-iscsi-helper
* https://github.com/longhorn/sparse-tools
* https://github.com/longhorn/backupstore

*External Dependencies:*

** Golang package dependencies:

* https://bitbucket.org/ww/goautoneg              BSD-3-Clause
* https://github.com/Jeffail/gabs                 MIT License
* https://github.com/sirupsen/logrus              MIT License
* https://github.com/aws/aws-sdk-go               Apache License 2.0
* https://github.com/c9s/goprocinfo               MIT License
* https://github.com/container-storage-interface/spec         Apache License 2.0
* https://github.com/docker/docker                Apache License 2.0
* https://github.com/docker/go-connections        Apache License 2.0
* https://github.com/docker/go-units              Apache License 2.0
* https://github.com/frostschutz/go-fibmap        MIT License
* https://github.com/go-ini/ini                   Apache License 2.0
* https://github.com/golang/protobuf	        BSD-3-Clause
* https://github.com/gorilla/context              BSD-3-Clause
* https://github.com/gorilla/handlers             BSD-3-Clause
* https://github.com/gorilla/mux                  BSD-3-Clause
* https://github.com/gorilla/websocket            BSD-3-Clause
* https://github.com/jmespath/go-jmespath         Apache License 2.0
* https://github.com/kubernetes-csi/drivers       Apache License 2.0
* https://github.com/kubernetes-incubator/external-storage    Apache License 2.0
* https://github.com/mitchellh/mapstructure       MIT License
* https://github.com/pkg/errors                   BSD-2-Clause
* https://github.com/rancher/go-rancher           Apache License 2.0
* https://github.com/robfig/cron                  MIT License
* https://github.com/satori/go.uuid               MIT License
* https://github.com/urfave/cli                   MIT License
* https://github.com/yasker/go-websocket-toolbox  Apache 2.0
* https://github.com/yasker/nsfilelock            Apache 2.0
* https://golang.org/x/crypto/			BSD-3-Clause
* https://golang.org/x/net/                       BSD-3-Clause
* https://golang.org/x/sys/                       BSD-3-Clause
* https://golang.org/x/text/                      BSD-3-Clause
* https://google.golang.org/genproto	        Apache License 2.0
* https://google.golang.org/grpc                  Apache License 2.0
* https://gopkg.in/check.v1                       BSD-3-Clause
* https://gopkg.in/yaml.v2			Apache License 2.0
* https://k8s.io/apiextensions-apiserver          Apache License 2.0
* https://k8s.io/kubernetes                       Apache License 2.0

** Binary dependency:

The current default Longhorn frontend TGT has a binary dependency on a modified version of the upstream TGT project. The modified version is hosted at https://github.com/rancher/tgt branch `longhorn`, with GPLv2 license.

*Initial Committers*

 * Sheng Yang @yasker (Rancher Labs), since 10/2015
 * Darren Shepherd @ibuildthecloud (Rancher Labs), since 03/2016

**Current Maintainer**

 * Sheng Yang @yasker (Rancher Labs)

*Infrastructure requests (CI / CNCF Cluster):*

_Development needs:_

None at the moment.

_Production needs:_

None at the moment.

*Communication Channels:*

 * Issue tracker: https://github.com/longhorn/longhorn/issues
 * Slack: https://rancher-users.slack.com `#longhorn-storage` channel
 * Forum: https://forums.rancher.com/c/longhorn

*Website:* https://github.com/longhorn/longhorn/ (To be moved to https://longhorn.io)

*Release methodology and mechanics:*
The version of Longhorn is based on Sematic Versioning 2.0.0. The version number follows the pattern of MAJOR.MINOR.PATCH. We are planning to do about 3 minor versions and 1 major per year. We will do patch releases according to the necessity.

*Social media accounts:*

 * WIP

*Existing sponsorship:* Rancher Labs

*Community size:*

GitHub https://github.com/longhorn/longhorn/stargazers[600+ stars]

Slack channel #longhorn-storage https://rancher-users.slack.com[200+ users]
