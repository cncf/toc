# Cilium Project Proposal

## Name of Project
Cilium

## Description

Cilium is an open source project for providing and transparently securing network
connectivity, and loadbalancing between application workloads such as
application containers or processes. Cilium operates at Layer 3/4 to provide
traditional networking and security services as well as Layer 7 to protect and
secure use of modern application protocols such as HTTP, gRPC and Kafka. Cilium
implements the Container Networking Interface (CNI) and integrates directly with
Kubernetes.

A new Linux kernel technology called eBPF is at the foundation of Cilium. It
supports dynamic insertion of eBPF bytecode into the Linux kernel at various
integration points such as: network IO, application sockets, and tracepoints to
implement security, networking and visibility logic. eBPF is highly efficient
and flexible. (To learn more about eBPF, visit [eBPF.io](eBPF.io).)

## Features

Simple Networking
-----------------

A simple flat Layer 3 network with the ability to span multiple clusters
connects all application containers. IP allocation is kept simple by using host
scope allocators. This means that each host can allocate IPs without any
coordination between hosts.

The following multi node networking models are supported:

* **Overlay:** Encapsulation-based virtual network spanning all hosts.
  Currently VXLAN and Geneve are baked in but all encapsulation formats
  supported by Linux can be enabled.

  When to use this mode: This mode has minimal infrastructure and integration
  requirements. It works on almost any network infrastructure as the only
  requirement is IP connectivity between hosts which is typically already
  given.

* **Native Routing:** Use of the regular routing table of the Linux host.
  The network is required to be capable to route the IP addresses of the
  application containers.

  When to use this mode: This mode is for advanced users and requires some
  awareness of the underlying networking infrastructure. This mode works well
  with:

  - Native IPv6 networks
  - In conjunction with cloud network routers
  - If you are already running routing daemons


Protect and secure APIs transparently
-------------------------------------

Ability to secure modern application protocols such as REST/HTTP, gRPC and
Kafka. Traditional firewalls operates at Layer 3 and 4. A protocol running on a
particular port is either completely trusted or blocked entirely. Cilium
provides the ability to filter on individual application protocol requests such
as:

- Allow all HTTP requests with method ``GET`` and path ``/public/.*``. Deny all
  other requests.
- Allow ``service1`` to produce on Kafka topic ``topic1`` and ``service2`` to
  consume on ``topic1``. Reject all other Kafka messages.
- Require the HTTP header ``X-Token: [0-9]+`` to be present in all REST calls.

See the section `Layer 7 Policy`_ in our documentation for the latest list of
supported protocols and examples on how to use it.

Secure service to service communication based on identities
-----------------------------------------------------------

Modern distributed applications rely on technologies such as application
containers to facilitate agility in deployment and scale out on demand. This
results in a large number of application containers to be started in a short
period of time. Typical container firewalls secure workloads by filtering on
source IP addresses and destination ports. This concept requires the firewalls
on all servers to be manipulated whenever a container is started anywhere in
the cluster.

In order to avoid this situation which limits scale, Cilium assigns a security
identity to groups of application containers which share identical security
policies. The identity is then associated with all network packets emitted by
the application containers, allowing to validate the identity at the receiving
node. Security identity management is performed using a key-value store.

Secure access to and from external services
-------------------------------------------

Label based security is the tool of choice for cluster internal access control.
In order to secure access to and from external services, traditional CIDR based
security policies for both ingress and egress are supported. This allows to
limit access to and from application containers to particular IP ranges.

Load balancing
--------------

Distributed load balancing for traffic between application containers and to
external services. The loadbalancing is implemented using BPF using efficient
hashtables allowing for almost unlimited scale and supports direct server
return (DSR) if the loadbalancing operation is not performed on the source
host.
*Note: load balancing requires connection tracking to be enabled. This is the
default.*

Monitoring and Troubleshooting
------------------------------

The ability to gain visibility and to troubleshoot issues is fundamental to the
operation of any distributed system. While we learned to love tools like
``tcpdump`` and ``ping`` and while they will always find a special place in our
hearts, we strive to provide better tooling for troubleshooting. This includes
tooling to provide:

- Event monitoring with metadata: When a packet is dropped, the tool doesn't
  just report the source and destination IP of the packet, the tool provides
  the full label information of both the sender and receiver among a lot of
  other information.

- Policy decision tracing: Why is a packet being dropped or a request rejected.
  The policy tracing framework allows to trace the policy decision process for
  both, running workloads and based on arbitrary label definitions.

- Metrics export via Prometheus: Key metrics are exported via Prometheus for
  integration with your existing dashboards.

- [Hubble](https://github.com/cilium/hubble/): An observability platform specifically written for Cilium. It
  provides service dependency maps, operational monitoring and alerting,
  and application and security visibility based on flow logs.

## Sponsor / Advisor from TOC
Justin Cormack

## Unique Identifier
cilium

## License

* The main agent is licensed under [Apache 2.0](https://github.com/cilium/cilium/blob/master/LICENSE)
* Portions of the eBPF datapath code are licensed under the [GPL](https://github.com/cilium/cilium/blob/master/bpf/COPYING)

## Maturity Level

Incubating

## Source control repositories

https://github.com/cilium

## External dependencies

From [go.mod](https://github.com/cilium/cilium/blob/master/go.mod):

- github.com/Azure/azure-sdk-for-go MIT License
- github.com/Azure/go-autorest Apaache 2.0
- github.com/asaskevich/govalidator MIT License
- github.com/aws/aws-sdk-go-v2 Apache 2.0
- github.com/aws/smithy-go Apache 2.0
- github.com/blang/semver MIT License
- github.com/cncf/udpa Apache 2.0
- github.com/containernetworking/cni Apache 2.0
- github.com/containernetworking/plugins Apache 2.0
- github.com/davecgh/go-spew ISC License
- github.com/docker/distribution Apache 2.0
- github.com/docker/docker Apache 2.0
- github.com/docker/go-connections Apache 2.0
- github.com/docker/libnetwork Apache 2.0
- github.com/fsnotify/fsnotify BSD-3-Clause
- github.com/go-bindata/go-bindata Public Domain
- github.com/go-openapi/errors Apache 2.0
- github.com/go-openapi/loads APache 2.0
- github.com/go-openapi/runtime Apache 2.0
- github.com/go-openapi/spec Apache 2.0
- github.com/go-openapi/strfmt Apache 2.0
- github.com/go-openapi/swag Apache 2.0
- github.com/go-openapi/validate Apache 2.0
- github.com/gogo/protobuf BSD-3-Clause
- github.com/golang/protobuf BSD-3-Clause
- github.com/google/go-cmp BSD-3-Clause
- github.com/google/gopacket BSD-3-Clause
- github.com/google/gops BSD-3-Clause
- github.com/google/renameio Apache 2.0
- github.com/google/uuid BSD-3-Clause
- github.com/gorilla/mux BSD-3-Clause
- github.com/hashicorp/consul/api MPL-2.0
- github.com/hashicorp/go-immutable-radix MPL-2.0
- github.com/hashicorp/golang-lru MPL-2.0
- github.com/ishidawataru/sctp Apache 2.0
- github.com/jeremywohl/flatten MIT
- github.com/kevinburke/ssh_config MIT
- github.com/kr/pretty MIT
- github.com/mattn/go-shellwords MIT
- github.com/miekg/dns BSD-3-Clause
- github.com/morikuni/aec MIT
- github.com/onsi/ginkgo MIT
- github.com/onsi/gomega MIT
- github.com/opencontainers/go-digest Apache 2.0
- github.com/opencontainers/image-spec Apache 2.0
- github.com/optiopay/kafka MIT
- github.com/pmezard/go-difflib https://github.com/pmezard/go-difflib/blob/master/LICENSE
- github.com/prometheus/client_golang Apache 2.0
- github.com/prometheus/client_model Apache 2.0
- github.com/russross/blackfriday/v2 Simplified BSD
- github.com/sasha-s/go-deadlock Apache 2.0
- github.com/servak/go-fastping MIT
- github.com/shirou/gopsutil/v3 BSD
- github.com/sirupsen/logrus MIT
- github.com/spf13/cobra Apache 2.0
- github.com/spf13/pflag BSD-3-Clause
- github.com/spf13/viper MIT
- github.com/stretchr/testify MIT
- github.com/vishvananda/netlink Apache 2.0
- github.com/vishvananda/netns Apache 2.0
- go.etcd.io/etcd Apache 2.0
- go.uber.org/goleak MIT
- golang.org/x/crypto BSD-3-Clause
- golang.org/x/net BSD-3-Clause
- golang.org/x/sync BSD-3-Clause
- golang.org/x/sys BSD-3-Clause
- golang.org/x/time BSD-3-Clause
- golang.org/x/tools BSD-3-Clause
- google.golang.org/genproto Apache 2.0
- google.golang.org/grpc Apache 2.0
- google.golang.org/protobuf BSD-3-Clause
- gopkg.in/check.v1 BSD-3-Clause
- gopkg.in/ini.v1 Apache 2.0
- gopkg.in/natefinch/lumberjack.v2 MIT
- gopkg.in/yaml.v2 Apache 2.0
- k8s.io/api Apache 2.0
- k8s.io/apiextensions-apiserver Apache 2.0
- k8s.io/apimachinery Apache 2.0
- k8s.io/client-go Apache 2.0
- k8s.io/code-generator Apache 2.0
- k8s.io/klog/v2 Apache 2.0
- k8s.io/kube-openapi Apache 2.0
- sigs.k8s.io/controller-tools Apache 2.0
- sigs.k8s.io/structured-merge-diff/v4  Apache 2.0
- sigs.k8s.io/yaml MIT

### Runtime dependencies

- LLVM Apache 2.0
- iproute2 GPL

## Issue tracker

See [https://github.com/cilium/cilium/issues](https://github.com/cilium/cilium/issues)

## Initial committers

See [MAINTAINERS](https://github.com/cilium/cilium/blob/master/MAINTAINERS.md)

There are currently committers from eight companies, with the majority from Isovalent. The company block vote limit described in the [project governance](https://docs.cilium.io/en/latest/contributing/governance/commit_access/#company-block-vote-limit) already ensures that Isovalent do not have a controlling majority vote. 

## Infrastructure requirements (CI / CNCF Cluster)

N/A

## Statement on alignment with CNCF mission

Cilium's mission is to provide open source networking and network security for
the cloud native ecosystem. Cilium is deeply integrated with Kubernetes, etcd,
Prometheus, Envoy, and other CNCF projects. More and more users are adopting Cilium
and making it a key part of their infrastructure. As such, there is a mutual
interest in building a strong community involving Cilium core developers and
users to provide a neutral home for Cilium, and guarantee open governance for
the future.

## Communication channels

- Slack: http://cilium.io/slack (5877 members)
- GitHub: https://github.com/cilium/cilium

## Website

https://cilium.io/

## Release methodology and mechanics

Cilium employs semantic versioning to name each release ([release process details](https://docs.cilium.io/en/v1.9/contributing/release/)) and compiled container images hosted on [Quay](https://quay.io/repository/cilium/cilium) and [Docker Hub](https://hub.docker.com/u/cilium). SHA256 checksums are provided for all distributed binaries. The release process is fully automated. A new minor release is released every 4 months with extensive release notes. In between minor releases, micro releases for the last 3 minor releases are provided regularly with backports of crucial bugfixes and security fixes.

## Security processes

See [SECURITY.md](https://github.com/cilium/cilium/blob/master/SECURITY.md)

## Roadmap

The roadmap is driven by community requests and feedback:
 * GitHub [issues](https://github.com/cilium/cilium/issues?q=is%3Aopen+is%3Aissue+label%3Akind%2Fenhancement) with the `enhancement` label
 * Annual public user surveys ([2021 survey](https://cilium.io/blog/2021/02/25/user-survey-2021))

The features committed to the next release are tracked in GitHub projects ([1.10 release example](https://github.com/cilium/cilium/projects/131)).

## Social media account
Twitter: https://twitter.com/ciliumproject

## Community size

Adopters: Adobe, Alibaba, Canonical, Datadog, DigitalOcean, GitLab, Google, SAP Concur, etc.
See [USERS.md](https://github.com/cilium/cilium/blob/master/USERS.md) for a list of over 20 publicly-announced adopters.

GitHub stars: +7.8K

Container image pulls: +100M across Docker Hub + Quay

Slack: 5800+ members

As of April 6, 2021, we have 263 people who have directly contributed code to
the main repository cilium/cilium. Numerous people have contributed by
reporting issues, testing, helping out on Slack, and so on.

We hold a weekly community meeting with 20-40 attendees each week.
