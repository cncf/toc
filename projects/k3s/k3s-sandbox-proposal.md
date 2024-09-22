# K3s CNCF Sandbox Project Proposal

Authors: 
- Darren Shepherd, Rancher Labs, CTO
- Craig Jellick, Rancher Labs, Director of Engineering


## Background

Link to Presentation - Coming soon. We'd like to have some guidance on what you want to see in the presentation.

[**Link to GitHub project**](https://github.com/rancher/k3s)

### Project Goal

k3s's goal is to provide a CNCF certified, production-grade Kubernetes distribution that is lightweight and easy to configure and operate. In doing so, k3s seeks to promote and improve the adoption of Kubernetes for the following use cases:
- Edge
- IoT
- CI Pipelines
- Development-to-production consistent environments
- ARM devices
- Embedding k8s in software appliances

### Current Status
k3s reached GA quality in November of 2019. It maintains a release cadence consistent with upstream Kubernetes, seeking to release within one month of upstream for the three most recent Kubernetes minor releases. As of this writing, the most recent releases are as follows:
| Upstream Kubernetes | k3s |
| ------------- | ------------- |
| [v1.18.2](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.2)  | [v1.18.2+k3s1](https://github.com/rancher/k3s/releases/tag/v1.18.2%2Bk3s1)  |
| [v1.17.5](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.5)  | [v1.17.5+k3s1](https://github.com/rancher/k3s/releases/tag/v1.17.5%2Bk3s1)  |
|  [v1.16.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.16.9) | [v1.16.9+k3s1](https://github.com/rancher/k3s/releases/tag/v1.16.9%2Bk3s1)  |

The project has a vibrant community. It has over 12,000 stars on GitHub and 89 code contributors. It's most recent stable release has over [60,000 downloads across three different architectures](https://api.github.com/repos/rancher/k3s/releases/24836622).

### Future Plans
k3s's roadmap includes:
- Improving stability and support across operating systems and architectures
- Enhancing security through SELinux support and CIS compliance
- Developing HA solutions appropriate for small edge clusters where even three nodes is cost prohibitive
- Improving support for Kubernetes cloud providers
- Improved support for database backends, including embedded etcd

while continuing to maintain release cadence and conformance with upstream Kubernetes.

## Project Scope
### Clear project definition
As previously stated, k3s is a CNCF certified, production-grade, lightweight Kubernetes distribution. It is packaged with sane defaults to achieve a fully functional cluster with minimal dependencies.

To achieve this, k3s has the following scope:
1. Packaged as a single binary
1. Shipped with a lightweight storage backend based on sqlite3 as the default storage mechanism. etcd3, MySQL, Postgres are also supported
1. Wrapped in simple launcher that handles a lot of the complexity of TLS certificates and kubernetes configuration options
1. Brings its own dependencies such that just a sane kernel and cgroup mounts are needed. k3s packages the following required dependencies:
    - containerd
    - Flannel
    - CoreDNS
    - CNI
    - Host utilities (iptables, socat, etc)
    - Ingress controller (traefik)
    - Embedded service loadbalancer
    - Embedded network policy controller
  
### Value-add to the CNCF ecosystem
k3s is bringing Kubernetes to environments and use cases that are not easily accessible to upstream Kubernetes, such as edge and IoT. In doing so, it is promoting the overall growth and adoption of Kubernetes. When coupled with Kubernetes's explosive growth in the cloud and the datacenter, k3s is helping Kubernetes to become the universal compute platform.

### Alignment with other CNCF projects
***Does the project align and actively collaborate with other CNCF projects?***

k3s is incredibly well aligned with other CNCF projects. It leverages existing CNCF projects such as Kubernetes, CNI, CRI, CSI, containerd, CoreDNS, and Helm. In doing so, its maintainers contribute to these projects by opening issues, submitting PRs, promoting them on social media, and engaging in technical discussions.

***Does the project require any specific versions of projects (or APIs) to interoperate? (e.g. K8s API, CSI, CNI, CRI)?***

k3s generally seeks to support the latest stable versions of the CNCF projects that it leverages. For projects that are themselves closely tied to Kubernetes, such as CNI, CRI, and CSI, k3s supports the same versions that upstream Kubernetes supports.

***Does the project augment or benefit other CNCF projects?***

As stated, k3s benefits Kubernetes by bringing it to environments that are not easily accessible to upstream Kubernetes. In focusing on secondary use cases and environments, k3s also illuminates problems or bugs in the other CNCF projects it leverages. Here are a few examples: [kubernetes#89554](https://github.com/kubernetes/kubernetes/issues/89554), [cri-tools#585](https://github.com/kubernetes-sigs/cri-tools/issues/585), [kubernetes#88376](https://github.com/kubernetes/kubernetes/issues/88376), and [containerd#2991](https://github.com/containerd/containerd/issues/2991).

### Anticipated use cases
With hundreds of thousands of downloads and a large and growing community of users, k3s's use cases are well proven. These use cases have already been stated: edge, IoT, ARM devices, and software appliances. One major use case that has not yet been addressed is that of fleet management. This is tied directly to the "edge" and IoT use cases, but is worth calling out explicitly. k3s's ease of deployment, operations, and upgrades makes the managing thousands or even millions of small Kubernetes clusters a possibility.

Alignment with SIG Reference Model
Does the project align with the SIG CNCF reference model and which capabilities does it require/provide at each level of the reference model.

### High level architecture
Being a distribution of Kubernetes, k3s's architecture looks largely similar to that of Kubernetes itself. However, there are a few key changes:
1. Kubernetes control plane processes such as the api server, scheduler, and controller-manager on server nodes and kubelet and kube-proxy on agent nodes are embedded within a single process
1. Flannel for CNI is managed within the k3s process
1. Containerd for CRI is shipped with and launched by k3s
1. By default, the datastore backend is an embedded SQLite database rather than etcd
1. A websocket-based tunnel proxy is used to manage all control plane communication between agent and server nodes.

This architecture is illustrated in the following diagram: 
![k3s](https://k3s.io/images/how-it-works-k3s.svg "K3s Architecture")


## Formal Requirements
Here is k3s's status on [the formal requirements](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#sandbox-stage)
- Require 3 of the 11 TOC members to step forward as sponsors to enter the sandbox - Pending
- Adopt the CNCF Code of Conduct - [k3s is adopting the CNCF Community Code of Conduct](https://github.com/rancher/k3s/pull/1783)
- Adhere to CNCF IP Policy (including trademark transferred) - We agree to this policy and will work through the requirements with the CNCF as the process moves forward
- List their sandbox status prominently on website/readme - We agree to do this

_Are there any anticipated issues with any of the criteria?_

No, we anticipate no issues.

_Has the TOC been approached for sponsorship? If yes, which TOC members have agreed to sponsor the project?_

Not officially.

## CNCF IP Policy
_Becoming a sandbox project requires adoption of the CNCF IP Policy: https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy_

_Note: there is a grace period after becoming a sandbox period to enable projects to adopt the policy, however, some prep is required to ensure there are no major blockers._

_Has the IP policy been reviewed?_

Yes.

_List the repos for the project and their current license_

- github.com/rancher/k3s - Apache 2.0
- github.com/rancher/k3s-upgrade - Apache 2.0
- github.com/rancher/helm-controller - Apache 2.0
- github.com/rancher/klipper-helm - Apache 2.0
- github.com/rancher/klipper-lb - Apache 2.0
- github.com/rancher/k3s-root - Apache 2.0

_List any dependent repos (upstream/downstream) that are required to build the project (including but not limited to libraries, commercial tools, plugins)_

k3s is committed to using open source throughout its build process. GitHub is used to maintain the code. Drone is used as the CI system.

The following is a table of dependent libraries and their libraries, as reported by [askalono](https://github.com/amzn/askalono/releases)
| Library's license file | License |
| ------------- | ------------- |
| gopkg.in/natefinch/lumberjack.v2/LICENSE | MIT |
| gopkg.in/yaml.v2/LICENSE | Apache-2.0 |
| gopkg.in/yaml.v2/NOTICE 1.0 | Apache-2.0 |
| gopkg.in/yaml.v2/LICENSE.libyaml  | MIT |
| gopkg.in/robfig/cron.v2/LICENSE 0.966967 | MIT |
| gopkg.in/square/go-jose.v2/LICENSE | Apache-2.0 |
| gopkg.in/square/go-jose.v2/json/LICENSE  | BSD-3-Clause |
| gopkg.in/warnings.v0/LICENSE  | BSD-2-Clause |
| gopkg.in/gcfg.v1/LICENSE | BSD-3-Clause |
| gopkg.in/inf.v0/LICENSE | BSD-3-Clause |
| cloud.google.com/go/LICENSE | Apache-2.0 |
| go.opencensus.io/LICENSE | Apache-2.0 |
| vbom.ml/util/LICENSE | MIT |
| go.etcd.io/bbolt/LICENSE | MIT |
| go.etcd.io/etcd/LICENSE | Apache-2.0 |
| go.etcd.io/etcd/NOTICE | Confidence threshold not high enough for any known license |
| gonum.org/v1/gonum/LICENSE  | BSD-3-Clause |
| google.golang.org/grpc/LICENSE | Apache-2.0 |
| google.golang.org/genproto/LICENSE | Apache-2.0 |
| google.golang.org/appengine/LICENSE | Apache-2.0 |
| google.golang.org/api/LICENSE  | BSD-3-Clause |
| google.golang.org/api/googleapi/internal/uritemplates/LICENSE  | MIT |
| k8s.io/metrics/LICENSE | Apache-2.0 |
| k8s.io/heapster/LICENSE | Apache-2.0 |
| k8s.io/client-go/LICENSE | Apache-2.0 |
| k8s.io/kube-aggregator/LICENSE | Apache-2.0 |
| k8s.io/cloud-provider/LICENSE | Apache-2.0 |
| k8s.io/cluster-bootstrap/LICENSE | Apache-2.0 |
| k8s.io/cri-api/LICENSE | Apache-2.0 |
| k8s.io/apiextensions-apiserver/LICENSE | Apache-2.0 |
| k8s.io/gengo/LICENSE | Apache-2.0 |
| k8s.io/kube-openapi/LICENSE | Apache-2.0 |
| k8s.io/kube-controller-manager/LICENSE | Apache-2.0 |
| k8s.io/utils/LICENSE | Apache-2.0 |
| k8s.io/utils/inotify/LICENSE  | BSD-3-Clause |
| k8s.io/utils/inotify/PATENTS | Confidence threshold not high enough for any known license |
| k8s.io/legacy-cloud-providers/LICENSE | Apache-2.0 |
| k8s.io/kube-scheduler/LICENSE | Apache-2.0 |
| k8s.io/cli-runtime/LICENSE | Apache-2.0 |
| k8s.io/csi-translation-lib/LICENSE | Apache-2.0 |
| k8s.io/kubelet/LICENSE | Apache-2.0 |
| k8s.io/kube-proxy/LICENSE | Apache-2.0 |
| k8s.io/api/LICENSE | Apache-2.0 |
| k8s.io/apiserver/LICENSE | Apache-2.0 |
| k8s.io/kubernetes/LICENSE | Apache-2.0 |
| k8s.io/kubernetes/third_party/forked/golang/LICENSE  | BSD-3-Clause |
| k8s.io/kubernetes/third_party/forked/golang/PATENTS | Confidence threshold not high enough for any known license |
| k8s.io/kubernetes/third_party/forked/gonum/graph/LICENSE  | BSD-3-Clause |
| k8s.io/klog/LICENSE | Apache-2.0 |
| k8s.io/kubectl/LICENSE | Apache-2.0 |
| k8s.io/code-generator/LICENSE | Apache-2.0 |
| k8s.io/apimachinery/LICENSE | Apache-2.0 |
| k8s.io/component-base/LICENSE | Apache-2.0 |
| go.mongodb.org/mongo-driver/LICENSE | Apache-2.0 |
| go.uber.org/zap/LICENSE.txt  | MIT |
| go.uber.org/atomic/LICENSE.txt  | MIT |
| go.uber.org/multierr/LICENSE.txt  | MIT |
| golang.org/x/crypto/LICENSE  | BSD-3-Clause |
| golang.org/x/crypto/PATENTS | Confidence threshold not high enough for any known license |
| golang.org/x/tools/LICENSE  | BSD-3-Clause |
| golang.org/x/tools/PATENTS | Confidence threshold not high enough for any known license |
| golang.org/x/net/LICENSE  | BSD-3-Clause |
| golang.org/x/net/PATENTS | Confidence threshold not high enough for any known license |
| golang.org/x/oauth2/LICENSE  | BSD-3-Clause |
| golang.org/x/time/LICENSE  | BSD-3-Clause |
| golang.org/x/time/PATENTS | Confidence threshold not high enough for any known license |
| golang.org/x/sys/LICENSE  | BSD-3-Clause |
| golang.org/x/sys/PATENTS | Confidence threshold not high enough for any known license |
| golang.org/x/text/LICENSE  | BSD-3-Clause |
| golang.org/x/text/PATENTS | Confidence threshold not high enough for any known license |
| golang.org/x/sync/LICENSE  | BSD-3-Clause |
| golang.org/x/sync/PATENTS | Confidence threshold not high enough for any known license |
| sigs.k8s.io/structured-merge-diff/v3/LICENSE | Apache-2.0 |
| sigs.k8s.io/yaml/LICENSE | Confidence threshold not high enough for any known license |
| sigs.k8s.io/apiserver-network-proxy/konnectivity-client/LICENSE | Apache-2.0 |
| sigs.k8s.io/kustomize/LICENSE | Apache-2.0 |
| github.com/Rican7/retry/LICENSE  | MIT |
| github.com/gorilla/websocket/LICENSE | BSD-2-Clause |
| github.com/gorilla/mux/LICENSE  | BSD-3-Clause |
| github.com/mitchellh/go-wordwrap/LICENSE.md| MIT |
| github.com/mitchellh/mapstructure/LICENSE | MIT |
| github.com/cyphar/filepath-securejoin/LICENSE  | BSD-3-Clause |
| github.com/gophercloud/gophercloud/LICENSE | Apache-2.0 |
| github.com/natefinch/lumberjack/LICENSE | MIT |
| github.com/davecgh/go-spew/LICENSE  | ISC |
| github.com/PuerkitoBio/urlesc/LICENSE  | BSD-3-Clause |
| github.com/PuerkitoBio/purell/LICENSE  | BSD-3-Clause |
| github.com/blang/semver/LICENSE  | MIT |
| github.com/karrick/godirwalk/LICENSE  | BSD-2-Clause |
| github.com/kubernetes-sigs/cri-tools/LICENSE | Apache-2.0 |
| github.com/evanphx/json-patch/LICENSE | BSD-3-Clause |
| github.com/urfave/cli/LICENSE  | MIT |
| github.com/urfave/cli/v2/LICENSE  | MIT |
| github.com/docker/libnetwork/LICENSE | Apache-2.0 |
| github.com/docker/go-metrics/LICENSE | Apache-2.0 |
| github.com/docker/go-metrics/LICENSE.docs| CC-BY-SA-4.0 |
| github.com/docker/go-metrics/NOTICE | Confidence threshold not high enough for any known license |
| github.com/docker/go-events/LICENSE | Apache-2.0 |
| github.com/docker/docker/LICENSE | Apache-2.0 |
| github.com/docker/docker/NOTICE | Confidence threshold not high enough for any known license |
| github.com/docker/docker/pkg/symlink/LICENSE.APACHE | Apache-2.0 |
| github.com/docker/docker/pkg/symlink/LICENSE.BSD  | BSD-3-Clause |
| github.com/docker/go-connections/LICENSE | Apache-2.0 |
| github.com/docker/distribution/LICENSE | Apache-2.0 |
| github.com/docker/go-units/LICENSE | Apache-2.0 |
| github.com/docker/spdystream/LICENSE | Apache-2.0 |
| github.com/docker/spdystream/LICENSE.docs| CC-BY-SA-4.0 |
| github.com/JeffAshton/win_pdh/LICENSE | BSD-3-Clause |
| github.com/Nvveen/Gotty/LICENSE  | BSD-2-Clause-FreeBSD |
| github.com/grpc-ecosystem/go-grpc-prometheus/LICENSE | Apache-2.0 |
| github.com/daviddengcn/go-colortext/LICENSE | Confidence threshold not high enough for any known license |
| github.com/robfig/cron/LICENSE 0.966967 | MIT |
| github.com/tchap/go-patricia/LICENSE | MIT |
| github.com/fsnotify/fsnotify/LICENSE  | BSD-3-Clause |
| github.com/beorn7/perks/LICENSE  | MIT |
| github.com/go-bindata/go-bindata/LICENSE | Confidence threshold not high enough for any known license |
| github.com/go-sql-driver/mysql/LICENSE 1.0 | MPL-2.0 |
| github.com/go-stack/stack/LICENSE.md| MIT |
| github.com/canonical/go-dqlite/LICENSE | Apache-2.0 |
| github.com/vishvananda/netns/LICENSE | Apache-2.0 |
| github.com/vishvananda/netlink/LICENSE | Apache-2.0 |
| github.com/hashicorp/golang-lru/LICENSE 1.0 | MPL-2.0 |
| github.com/Azure/go-ansiterm/LICENSE | MIT |
| github.com/Azure/go-autorest/tracing/LICENSE | Apache-2.0 |
| github.com/Azure/go-autorest/logger/LICENSE | Apache-2.0 |
| github.com/Azure/go-autorest/autorest/mocks/LICENSE | Apache-2.0 |
| github.com/Azure/go-autorest/autorest/LICENSE | Apache-2.0 |
| github.com/Azure/go-autorest/autorest/date/LICENSE | Apache-2.0 |
| github.com/Azure/go-autorest/autorest/to/LICENSE | Apache-2.0 |
| github.com/Azure/go-autorest/autorest/validation/LICENSE | Apache-2.0 |
| github.com/Azure/go-autorest/autorest/adal/LICENSE | Apache-2.0 |
| github.com/Azure/azure-sdk-for-go/LICENSE | Apache-2.0 |
| github.com/Azure/azure-sdk-for-go/NOTICE | Confidence threshold not high enough for any known license |
| github.com/golang/groupcache/LICENSE | Apache-2.0 |
| github.com/golang/mock/LICENSE | Apache-2.0 |
| github.com/golang/protobuf/LICENSE  | BSD-3-Clause |
| github.com/json-iterator/go/LICENSE  | MIT |
| github.com/miekg/dns/LICENSE | BSD-3-Clause |
| github.com/GoogleCloudPlatform/k8s-cloud-provider/LICENSE | Apache-2.0 |
| github.com/google/gofuzz/LICENSE | Apache-2.0 |
| github.com/google/go-cmp/LICENSE  | BSD-3-Clause |
| github.com/google/btree/LICENSE | Apache-2.0 |
| github.com/google/tcpproxy/LICENSE | Apache-2.0 |
| github.com/google/uuid/LICENSE  | BSD-3-Clause |
| github.com/google/cadvisor/LICENSE | Apache-2.0 |
| github.com/buger/jsonparser/LICENSE  | MIT |
| github.com/Microsoft/go-winio/LICENSE | MIT |
| github.com/Microsoft/hcsshim/LICENSE | MIT |
| github.com/mrunalp/fileutils/LICENSE | Apache-2.0 |
| github.com/shurcooL/sanitized_anchor_name/LICENSE  | MIT |
| github.com/containerd/continuity/LICENSE | Apache-2.0 |
| github.com/containerd/containerd/LICENSE | Apache-2.0 |
| github.com/containerd/containerd/NOTICE | Confidence threshold not high enough for any known license |
| github.com/containerd/typeurl/LICENSE | Apache-2.0 |
| github.com/containerd/go-cni/LICENSE | Apache-2.0 |
| github.com/containerd/go-runc/LICENSE | Apache-2.0 |
| github.com/containerd/fifo/LICENSE | Apache-2.0 |
| github.com/containerd/cgroups/LICENSE | Apache-2.0 |
| github.com/containerd/ttrpc/LICENSE | Apache-2.0 |
| github.com/containerd/cri/LICENSE | Apache-2.0 |
| github.com/containerd/console/LICENSE | Apache-2.0 |
| github.com/mxk/go-flowrate/LICENSE  | BSD-3-Clause |
| github.com/gregjones/httpcache/LICENSE.txt  | MIT |
| github.com/konsorten/go-windows-terminal-sequences/LICENSE  | MIT |
| github.com/liggitt/tabwriter/LICENSE  | BSD-3-Clause |
| github.com/pquerna/cachecontrol/LICENSE | Apache-2.0 |
| github.com/go-openapi/loads/LICENSE | Apache-2.0 |
| github.com/go-openapi/strfmt/LICENSE | Apache-2.0 |
| github.com/go-openapi/validate/LICENSE | Apache-2.0 |
| github.com/go-openapi/analysis/LICENSE | Apache-2.0 |
| github.com/go-openapi/runtime/LICENSE | Apache-2.0 |
| github.com/go-openapi/spec/LICENSE | Apache-2.0 |
| github.com/go-openapi/spec/license.go| Apache-2.0 |
| github.com/go-openapi/swag/LICENSE | Apache-2.0 |
| github.com/go-openapi/errors/LICENSE | Apache-2.0 |
| github.com/go-openapi/jsonpointer/LICENSE | Apache-2.0 |
| github.com/go-openapi/jsonreference/LICENSE | Apache-2.0 |
| github.com/peterbourgon/diskv/LICENSE  | MIT |
| github.com/checkpoint-restore/go-criu/LICENSE | Apache-2.0 |
| github.com/juju/errors/LICENSE 0.912832 | LGPL-3.0-only |
| github.com/lxc/lxd/COPYING| Apache-2.0 |
| github.com/imdario/mergo/LICENSE  | BSD-3-Clause |
| github.com/jonboulle/clockwork/LICENSE | Apache-2.0 |
| github.com/fatih/camelcase/LICENSE.md| MIT |
| github.com/NYTimes/gziphandler/LICENSE | Apache-2.0 |
| github.com/euank/go-kmsg-parser/LICENSE | Apache-2.0 |
| github.com/dgrijalva/jwt-go/LICENSE  | MIT |
| github.com/sirupsen/logrus/LICENSE | MIT |
| github.com/mistifyio/go-zfs/LICENSE | Apache-2.0 |
| github.com/godbus/dbus/LICENSE | BSD-2-Clause |
| github.com/BurntSushi/toml/COPYING| MIT |
| github.com/googleapis/gnostic/LICENSE | Apache-2.0 |
| github.com/exponent-io/jsonpath/LICENSE | MIT |
| github.com/containernetworking/plugins/LICENSE | Apache-2.0 |
| github.com/containernetworking/cni/LICENSE | Apache-2.0 |
| github.com/bhendo/go-powershell/LICENSE | MIT |
| github.com/spf13/afero/LICENSE.txt 0.94437 | Apache-2.0 |
| github.com/spf13/pflag/LICENSE  | BSD-3-Clause |
| github.com/spf13/cobra/LICENSE.txt 0.94437 | Apache-2.0 |
| github.com/syndtr/gocapability/LICENSE  | BSD-2-Clause |
| github.com/modern-go/reflect2/LICENSE | Apache-2.0 |
| github.com/modern-go/concurrent/LICENSE | Apache-2.0 |
| github.com/armon/circbuf/LICENSE | MIT |
| github.com/seccomp/libseccomp-golang/LICENSE  | BSD-2-Clause |
| github.com/cpuguy83/go-md2man/v2/LICENSE.md| MIT |
| github.com/matttproud/golang_protobuf_extensions/LICENSE | Apache-2.0 |
| github.com/matttproud/golang_protobuf_extensions/NOTICE | Confidence threshold not high enough for any known license |
| github.com/lib/pq/LICENSE.md  | MIT |
| github.com/opencontainers/image-spec/LICENSE | Apache-2.0 |
| github.com/opencontainers/runc/LICENSE | Apache-2.0 |
| github.com/opencontainers/runc/NOTICE | Confidence threshold not high enough for any known license |
| github.com/opencontainers/go-digest/LICENSE.code| Apache-2.0 |
| github.com/opencontainers/go-digest/LICENSE.docs| CC-BY-SA-4.0 |
| github.com/opencontainers/selinux/LICENSE | Apache-2.0 |
| github.com/opencontainers/runtime-spec/LICENSE | Apache-2.0 |
| github.com/rakelkar/gonetsh/LICENSE | Apache-2.0 |
| github.com/asaskevich/govalidator/LICENSE | MIT |
| github.com/rancher/dynamiclistener/LICENSE | Apache-2.0 |
| github.com/rancher/wrangler/LICENSE | Apache-2.0 |
| github.com/rancher/wrangler-api/LICENSE | Apache-2.0 |
| github.com/rancher/kine/LICENSE | Apache-2.0 |
| github.com/rancher/helm-controller/LICENSE | Apache-2.0 |
| github.com/rancher/remotedialer/LICENSE | Apache-2.0 |
| github.com/jmespath/go-jmespath/LICENSE 1.0 | Apache-2.0 |
| github.com/aws/aws-sdk-go/NOTICE.txt,Confidence threshold not high enough for any known license |
| github.com/aws/aws-sdk-go/LICENSE.txt| Apache-2.0 |
| github.com/bronze1man/goStrongswanVici/LICENSE | MIT |
| github.com/container-storage-interface/spec/LICENSE | Apache-2.0 |
| github.com/prometheus/client_golang/LICENSE | Apache-2.0 |
| github.com/prometheus/client_golang/NOTICE | Confidence threshold not high enough for any known license |
| github.com/prometheus/client_model/LICENSE | Apache-2.0 |
| github.com/prometheus/client_model/NOTICE | Confidence threshold not high enough for any known license |
| github.com/prometheus/common/LICENSE | Apache-2.0 |
| github.com/prometheus/common/NOTICE | Confidence threshold not high enough for any known license |
| github.com/prometheus/procfs/LICENSE | Apache-2.0 |
| github.com/prometheus/procfs/NOTICE | Confidence threshold not high enough for any known license |
| github.com/satori/go.uuid/LICENSE  | MIT |
| github.com/mailru/easyjson/LICENSE  | MIT |
| github.com/mindprince/gonvml/LICENSE | Apache-2.0 |
| github.com/MakeNowJust/heredoc/LICENSE | MIT |
| github.com/emicklei/go-restful/LICENSE | MIT |
| github.com/lithammer/dedent/LICENSE | MIT |
| github.com/russross/blackfriday/v2/LICENSE.txt| BSD-2-Clause |
| github.com/russross/blackfriday/LICENSE.txt| BSD-2-Clause |
| github.com/flosch/pongo2/LICENSE | MIT |
| github.com/cilium/ebpf/LICENSE  | MIT |
| github.com/theckman/go-flock/LICENSE  | BSD-3-Clause |
| github.com/coreos/go-systemd/LICENSE | Apache-2.0 |
| github.com/coreos/go-systemd/NOTICE | Confidence threshold not high enough for any known license |
| github.com/coreos/flannel/LICENSE | Apache-2.0 |
| github.com/coreos/flannel/NOTICE | Confidence threshold not high enough for any known license |
| github.com/coreos/go-oidc/LICENSE | Apache-2.0 |
| github.com/coreos/go-oidc/NOTICE | Confidence threshold not high enough for any known license |
| github.com/coreos/go-iptables/LICENSE | Apache-2.0 |
| github.com/coreos/go-iptables/NOTICE | Confidence threshold not high enough for any known license |
| github.com/coreos/pkg/LICENSE | Apache-2.0 |
| github.com/coreos/pkg/NOTICE | Confidence threshold not high enough for any known license |
| github.com/inconshreveable/mousetrap/LICENSE 1.0 | Apache-2.0 |
| github.com/rubiojr/go-vhd/LICENSE | MIT |
| github.com/gogo/googleapis/LICENSE | Apache-2.0 |
| github.com/gogo/protobuf/LICENSE | BSD-3-Clause |
| github.com/mattn/go-sqlite3/LICENSE | MIT |
| github.com/mattn/go-shellwords/LICENSE | MIT |
| github.com/chai2010/gettext-go/LICENSE  | BSD-3-Clause |
| github.com/rootless-containers/rootlesskit/LICENSE | Apache-2.0 |
| github.com/pkg/errors/LICENSE  | BSD-2-Clause |
| github.com/munnerz/goautoneg/LICENSE 0.953917 | BSD-3-Clause |
| github.com/ghodss/yaml/LICENSE | Confidence threshold not high enough for any known license |
| github.com/vmware/govmomi/vim25/xml/LICENSE  | BSD-3-Clause |
| github.com/vmware/govmomi/LICENSE.txt| Apache-2.0 |

Additionally, k3s ships with or leverages the following technologies:
| Technology | License |
| ------------- | ------------- |
| https://github.com/containous/traefik | MIT |
| https://github.com/helm/helm | Apache 2.0 |
| https://github.com/coredns/coredns | Apache 2.0 |
| https://github.com/buildroot/buildroot | GPL v2 | 
| https://github.com/mirror/busybox | GPL v2

_What actions are required to be compliant with the IP policy?_

This is under review.

## Other Considerations
### _Please note, these are not gating criteria but rather to:_
- Collect a standard set of information for each project
- Provides a point in time capture of the state of the project which makes it easier to track progress at future reviews and / or promotion
- Help projects to prepare for SIG and TOC presentation
- Allow the SIG to review the project and perform due diligence for incubation
- Provide the TOC with the information required to accept sponsorship of a project and/or votes
- Identify and rectify any significant issues / blockers prior to presenting to the TOC and acceptance as a CNCF project

### Cloud Native
_Does the project meet the definition of Cloud Native? The CNCF charter states:_
```
“Cloud native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds. Containers, service meshes, microservices, immutable infrastructure, and declarative APIs exemplify this approach.


“These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.”
```

Yes.

### Project and Code Quality
_Are there any metrics around code quality? Are there good examples of code reviews? Are there enforced coding standards?_

_What are the performance goals and results? What performance tradeoffs have been made? What is the resource cost?_

We've used the Kubernetes scale tests to judge performance and help users size their k3s instances. This is documented here: https://rancher.com/docs/k3s/latest/en/installation/installation-requirements/#cpu-and-memory

_What is the CI/CD system? Are there code coverage metrics? What types of tests exist?_

- Drone is the CI system. PRs are exercised here: https://drone-pr.rancher.io/rancher/k3s
- We largely leverage the Sonobuoy test suites from Kubernetes. Additionally, we are currently developing automated tests for exercising the different installation configurations possible with k3s.

_Is there documentation?_

Yes, it is currently hosted at https://rancher.com/docs/k3s/latest/en/

_How is it deployed?_

k3s is deployed as a single binary on one or more machines. This is covered in detail in the documentation.

_How is it orchestrated?_

k3s is orchestrated by sharing a registration token and URL between nodes in the cluster. A node joins the cluster by hitting the registration URL with the shared token.

_How will the project benefit from acceptance into the CNCF?_

k3s will benefit from acceptance into the CNCF by seeing wider adoption and contributions. Organizations that put a high value on vendor-neutral platforms will see that k3s is a successful open source technology that goes beyond a single company.


_Has a security assessment by the security SIG been done? If not, what is the status/progress of the assessment?_

No.
