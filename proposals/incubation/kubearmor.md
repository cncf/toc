# KubeArmor for incubating 

- [Background](#background)
- [Progress Since Joining CNCF](#progress-since-joining-cncf)
  - [Development metrics](#development-metrics)
  - [Maintainers](#maintainers)
- [Incubation Stage Requirements](#incubation-stage-requirements)
  - [Adopters](#document-that-it-is-being-used-successfully-in-production-by-at-least-three-independent-end-users-which-in-the-tocs-judgement-are-of-adequate-quality-and-scope)
  - [Have a healthy number of committers](#have-a-healthy-number-of-committers)
  - [Demonstrate a substantial ongoing flow of commits and merged contributions](#demonstrate-a-substantial-ongoing-flow-of-commits-and-merged-contributions)
  - [A clear versioning scheme](#a-clear-versioning-scheme)
  - [Have committers from at least two organizations](#have-committers-from-at-least-two-organizations)
  - [Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge](#have-achieved-and-maintained-a-core-infrastructure-initiative-best-practices-badge)
  - [Adopted the CNCF Code of Conduct](#adopted-the-cncf-code-of-conduct)
  - [Explicitly define a project governance and committer process](#explicitly-define-a-project-governance-and-committer-process)
  - [Have a public list of project adopters for at least the primary repo](#have-a-public-list-of-project-adopters-for-at-least-the-primary-repo)
- [Security processes](#security-processes)
- [License](#license)
- [Source control repositories](#source-control-repositories)
- [Roadmap](#roadmap)
- [External Dependencies](#external-dependencies)


## Background

[KubeArmor](https://kubearmor.io/) is a runtime security engine for kubernetes and other cloud workloads. Users can harden their workloads, create Zero Trust security posture i.e., allowing specific systems actions and denying rest using KubeArmor security policies. KubeArmor supports [inline mitigation](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/differentiation.md) for preventing attacks. Due to its versatility, KubeArmor has found its relevance in Edge and 5G Control Plane security scenarios as well.

KubeArmor primarily provides the following key features:

- [Workload Hardening](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/hardening_guide.md)

    KubeArmor provides a set of hardening policies that are based on industry-leading compliance and attack frameworks such as CIS, MITRE, NIST-800-53, and STIGs. These policies are designed to help you secure your workloads in a way that is compliant with these frameworks and recommended best practices.

- [Zero Trust Policies](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/least_permissive_access.md)

    KubeArmor is a tool that helps organizations enforce a zero trust posture within their Kubernetes clusters. It allows users to define an allow-based policy that specifies the specific system behavior that is allowed, and denies or audits all other behavior.

- [Network Visibility & Microsegmentation](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/network_segmentation.md)

    KubeArmor has visibility into the network connections made into or from the pods. It has visibility across socket(), bind(), connect(), accept() calls. KubeArmor can help auto-generate k8s network policy which are enforceable using most of the k8s CNIs (Container Network Interface). 

- [Application Behavior](https://github.com/kubearmor/KubeArmor/blob/main/getting-started/workload_visibility.md)

    KubeArmor has visibility into systems and application behavior. KubeArmor summarizes/aggregates the information and provides a user-friendly view to figure out the application behavior.

## Progress Since Joining CNCF

### Development metrics

The KubeArmor devstats page and dashboards can be found [here](https://kubearmor.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

- The community has significantly grown since the project entered the CNCF sandbox.
  - We held biweekly community meetings consistently (total 43 since Sep 2021). The community did not skip a single meeting since its inception. The meeting records can be found in [here](https://docs.google.com/document/d/1IqIIG9Vz-PYpbUwrH0u99KYEM1mtnYe6BHrson4NqEs/edit).
  - Number of contributors: 30+ -> **150+**
  - Github stars: 100+ -> **1070+**
  - Github forks: 30+ -> **265**
  - Contributing organizations: 5+ -> **30+**
- [New PRs in last year](https://kubearmor.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1)
- KubeArmor maintainer team has mentored more than 10 candidates as part of LFX and GSoC mentorships.

### Maintainers

Initial maintainers were from AccuKnox primarily. We have independent maintainers and few other contributors who are shaping up to take the ownership of the modules. Following is the list sorted in alphabetical order:

- Achref Ben Saad (AccuKnox) @achrefbensaad
- Ankur Kothiwal (CERN) @Ankurk99
- Anurag Kumar (Independent) @kranurag7
- Barun Acharya (AccuKnox) @daemon1024
- Jaehyun Nam (Dankook Univ) @nam-jaehyun
- Rahul Jadhav (AccuKnox) @nyrahul

## Incubation Stage Requirements

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope

Adoption of KubeArmor is tracked in our [ADOPTERS.md](https://github.com/kubearmor/kubearmor/blob/main/ADOPTERS.md) file.

Beyond this, we have received interests from other organizations such as:
- [LF 5G Super Blueprint to use KubeArmor](https://wiki.lfnetworking.org/pages/viewpage.action?pageId=82905466) for 5G security.
- [KubeArmor Operator](https://catalog.redhat.com/software/containers/kubearmor/kubearmor-operator/6458b1ee3f880286315e7cfb) is also available on Red Hat Ecosystem Catalog.
- Securing 5G Control Plane using KubeArmor. [[ref1](https://nephio.org/5g-security-enhanced-open-radio-access-networks/)] [[ref2](https://5gsec.com/5G_Security/5G-Kubearmor.html)]

### Have a healthy number of committers

KubeArmor currently has[6 Maintainers](https://github.com/KubeArmor/KubeArmor/blob/main/MAINTAINERS.md#project-maintainers) from 4 organizations and 6 Committers from 4 organizations. The roles and the corresponding responsibilities are defined in the [Project Governance](https://github.com/KubeArmor/KubeArmor/blob/main/GOVERNANCE.md). 


### Demonstrate a substantial ongoing flow of commits and merged contributions

According to devstats, KubeArmor currently has [252](https://kubearmor.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1) contributors from [40](https://kubearmor.devstats.cncf.io/d/5/companies-table?orgId=1) companies belonging to 15 countries.

The project averages at [~100 contributions from around ~16 contributors](https://kubearmor.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) per month according to kubearmor.devstats.cncf.io contained within [30 merged PRs](https://kubearmor.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now-9d) on average per month for the last year.

### A clear versioning scheme

KubeArmor uses the [semantic versioning scheme](https://semver.org/#semantic-versioning-specification-semver).

KubeArmor follows roughly once every two months release cadence with version numbers using format of MAJOR.MINOR.PATCH. The latest release is v1.1.1.

We have releases documented at: https://github.com/KubeArmor/KubeArmor/releases.

KubeArmor has a release cadence of once in two month release cycle.

### Have committers from at least two organizations

KubeArmor has committers from 3 organisations and 1 independent committer.

### Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge

KubeArmor has achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](achieved and maintained a Core Infrastructure Initiative Best Practices Badge)

### Adopted the CNCF Code of Conduct

KubeArmor adopts [CNCF Code of Conduct](https://github.com/kubearmor/KubeArmor/blob/main/CODE_OF_CONDUCT.md)

### Explicitly define a project governance and committer process

KubeArmor has a defined Project Governance and Committer Process in [GOVERNANCE.md](https://github.com/kubearmor/KubeArmor/blob/main/GOVERNANCE.md)
Current list of active maintainers is maintained and updated at [MAINTAINERS.md](https://github.com/kubearmor/KubeArmor/blob/main/MAINTAINERS.md)

### Have a public list of project adopters for at least the primary repo 

Adoption of KubeArmor is tracked in our [ADOPTERS.md](https://github.com/kubearmor/kubearmor/blob/main/ADOPTERS.md) file.

## Security processes

See [SECURITY.md](https://github.com/KubeArmor/KubeArmor/blob/main/SECURITY.md)

## License

KubeArmor is licensed under the Apache 2.0 license.

## Source control repositories

https://github.com/KubeArmor

## Roadmap

The backlog/roadmap for KubeArmor can be found [here](https://github.com/orgs/kubearmor/projects/9/views/1).

## External Dependencies

From [go.mod](https://github.com/kubearmor/kubearmor/blob/main/KubeArmor/go.mod):

- github.com/Masterminds/goutils https://github.com/Masterminds/goutils/blob/v1.1.1/LICENSE.txt Apache-2.0
- github.com/Masterminds/semver/v3 https://github.com/Masterminds/semver/blob/v3.2.1/LICENSE.txt MIT
- github.com/Masterminds/sprig/v3 https://github.com/Masterminds/sprig/blob/v3.2.3/LICENSE.txt MIT
- github.com/beorn7/perks/quantile https://github.com/beorn7/perks/blob/v1.0.1/LICENSE MIT
- github.com/cespare/xxhash/v2 https://github.com/cespare/xxhash/blob/v2.2.0/LICENSE.txt MIT
- github.com/cilium/cilium/pkg/mountinfo https://github.com/cilium/cilium/blob/v1.13.2/LICENSE Apache-2.0
- github.com/cilium/ebpf https://github.com/cilium/ebpf/blob/v0.11.0/LICENSE MIT
- github.com/containerd/containerd https://github.com/containerd/containerd/blob/v1.7.1/LICENSE Apache-2.0
- github.com/containerd/ttrpc https://github.com/containerd/ttrpc/blob/v1.2.2/LICENSE Apache-2.0
- github.com/containerd/typeurl/v2 https://github.com/containerd/typeurl/blob/v2.1.1/LICENSE Apache-2.0
- github.com/davecgh/go-spew/spew https://github.com/davecgh/go-spew/blob/v1.1.1/LICENSE ISC
- github.com/docker/distribution https://github.com/docker/distribution/blob/v2.8.1/LICENSE Apache-2.0
- github.com/docker/docker https://github.com/docker/docker/blob/v24.0.7/LICENSE Apache-2.0
- github.com/docker/go-connections https://github.com/docker/go-connections/blob/v0.4.0/LICENSE Apache-2.0
- github.com/docker/go-units https://github.com/docker/go-units/blob/v0.5.0/LICENSE Apache-2.0
- github.com/emicklei/go-restful/v3 https://github.com/emicklei/go-restful/blob/v3.10.2/LICENSE MIT
- github.com/evanphx/json-patch/v5 https://github.com/evanphx/json-patch/blob/v5.6.0/v5/LICENSE BSD-3-Clause
- github.com/fsnotify/fsnotify https://github.com/fsnotify/fsnotify/blob/v1.6.0/LICENSE BSD-3-Clause
- github.com/go-logr/logr https://github.com/go-logr/logr/blob/v1.2.4/LICENSE Apache-2.0
- github.com/go-openapi/jsonpointer https://github.com/go-openapi/jsonpointer/blob/v0.19.6/LICENSE Apache-2.0
- github.com/go-openapi/jsonreference https://github.com/go-openapi/jsonreference/blob/v0.20.2/LICENSE Apache-2.0
- github.com/go-openapi/swag https://github.com/go-openapi/swag/blob/v0.22.3/LICENSE Apache-2.0
- github.com/gogo/protobuf https://github.com/gogo/protobuf/blob/v1.3.2/LICENSE BSD-3-Clause
- github.com/golang/groupcache/lru https://github.com/golang/groupcache/blob/41bb18bfe9da/LICENSE Apache-2.0
- github.com/golang/protobuf https://github.com/golang/protobuf/blob/v1.5.3/LICENSE BSD-3-Clause
- github.com/google/gnostic https://github.com/google/gnostic/blob/v0.6.9/LICENSE Apache-2.0
- github.com/google/go-cmp/cmp https://github.com/google/go-cmp/blob/v0.5.9/LICENSE BSD-3-Clause
- github.com/google/gofuzz https://github.com/google/gofuzz/blob/v1.2.0/LICENSE Apache-2.0
- github.com/google/uuid https://github.com/google/uuid/blob/v1.3.0/LICENSE BSD-3-Clause
- github.com/hashicorp/hcl https://github.com/hashicorp/hcl/blob/v1.0.0/LICENSE MPL-2.0
- github.com/huandu/xstrings https://github.com/huandu/xstrings/blob/v1.4.0/LICENSE MIT
- github.com/imdario/mergo https://github.com/imdario/mergo/blob/v0.3.15/LICENSE BSD-3-Clause
- github.com/josharian/intern https://github.com/josharian/intern/blob/v1.0.0/license.md MIT
- github.com/json-iterator/go https://github.com/json-iterator/go/blob/v1.1.12/LICENSE MIT
- github.com/magiconair/properties https://github.com/magiconair/properties/blob/v1.8.7/LICENSE.md BSD-2-Clause
- github.com/mailru/easyjson https://github.com/mailru/easyjson/blob/v0.7.7/LICENSE MIT
- github.com/matttproud/golang_protobuf_extensions/pbutil https://github.com/matttproud/golang_protobuf_extensions/blob/v1.0.4/LICENSE Apache-2.0
- github.com/mitchellh/copystructure https://github.com/mitchellh/copystructure/blob/v1.2.0/LICENSE MIT
- github.com/mitchellh/mapstructure https://github.com/mitchellh/mapstructure/blob/v1.5.0/LICENSE MIT
- github.com/mitchellh/reflectwalk https://github.com/mitchellh/reflectwalk/blob/v1.0.2/LICENSE MIT
- github.com/modern-go/concurrent https://github.com/modern-go/concurrent/blob/bacd9c7ef1dd/LICENSE Apache-2.0
- github.com/modern-go/reflect2 https://github.com/modern-go/reflect2/blob/v1.0.2/LICENSE Apache-2.0
- github.com/munnerz/goautoneg https://github.com/munnerz/goautoneg/blob/a7dc8b61c822/LICENSE BSD-3-Clause
- github.com/opencontainers/go-digest https://github.com/opencontainers/go-digest/blob/v1.0.0/LICENSE Apache-2.0
- github.com/opencontainers/image-spec/specs-go https://github.com/opencontainers/image-spec/blob/v1.1.0-rc3/LICENSE Apache-2.0
- github.com/opencontainers/runtime-spec/specs-go https://github.com/opencontainers/runtime-spec/blob/v1.1.0-rc.2/LICENSE Apache-2.0
- github.com/pelletier/go-toml/v2 https://github.com/pelletier/go-toml/blob/v2.0.7/LICENSE MIT
- github.com/pkg/errors https://github.com/pkg/errors/blob/v0.9.1/LICENSE BSD-2-Clause
- github.com/prometheus/client_golang/prometheus https://github.com/prometheus/client_golang/blob/v1.15.1/LICENSE Apache-2.0
- github.com/prometheus/client_model/go https://github.com/prometheus/client_model/blob/v0.4.0/LICENSE Apache-2.0
- github.com/prometheus/common https://github.com/prometheus/common/blob/v0.43.0/LICENSE Apache-2.0
- github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg https://github.com/prometheus/common/blob/v0.43.0/internal/bitbucket.org/ww/goautoneg/README.txt BSD-3-Clause
- github.com/prometheus/procfs https://github.com/prometheus/procfs/blob/v0.9.0/LICENSE Apache-2.0
- github.com/shopspring/decimal https://github.com/shopspring/decimal/blob/v1.3.1/LICENSE MIT
- github.com/sirupsen/logrus https://github.com/sirupsen/logrus/blob/v1.9.0/LICENSE MIT
- github.com/spf13/afero https://github.com/spf13/afero/blob/v1.9.5/LICENSE.txt Apache-2.0
- github.com/spf13/cast https://github.com/spf13/cast/blob/v1.5.0/LICENSE MIT
- github.com/spf13/jwalterweatherman https://github.com/spf13/jwalterweatherman/blob/v1.1.0/LICENSE MIT
- github.com/spf13/pflag https://github.com/spf13/pflag/blob/v1.0.5/LICENSE BSD-3-Clause
- github.com/spf13/viper https://github.com/spf13/viper/blob/v1.15.0/LICENSE MIT
- github.com/subosito/gotenv https://github.com/subosito/gotenv/blob/v1.4.2/LICENSE MIT
- go.uber.org/atomic https://github.com/uber-go/atomic/blob/v1.11.0/LICENSE.txt MIT
- go.uber.org/multierr https://github.com/uber-go/multierr/blob/v1.11.0/LICENSE.txt MIT
- go.uber.org/zap https://github.com/uber-go/zap/blob/v1.24.0/LICENSE.txt MIT
- golang.org/x/crypto https://cs.opensource.google/go/x/crypto/+/v0.9.0:LICENSE BSD-3-Clause
- golang.org/x/exp https://cs.opensource.google/go/x/exp/+/c95f2b4c:LICENSE BSD-3-Clause
- golang.org/x/net https://cs.opensource.google/go/x/net/+/v0.10.0:LICENSE BSD-3-Clause
- golang.org/x/oauth2 https://cs.opensource.google/go/x/oauth2/+/v0.8.0:LICENSE BSD-3-Clause
- golang.org/x/sys/unix https://cs.opensource.google/go/x/sys/+/v0.10.0:LICENSE BSD-3-Clause
- golang.org/x/term https://cs.opensource.google/go/x/term/+/v0.8.0:LICENSE BSD-3-Clause
- golang.org/x/text https://cs.opensource.google/go/x/text/+/v0.9.0:LICENSE BSD-3-Clause
- golang.org/x/time/rate https://cs.opensource.google/go/x/time/+/v0.3.0:LICENSE BSD-3-Clause
- gomodules.xyz/jsonpatch/v2 https://github.com/gomodules/jsonpatch/blob/v2.2.0/v2/LICENSE Apache-2.0
- google.golang.org/genproto/googleapis/rpc/status https://github.com/googleapis/go-genproto/blob/daa745c078e1/LICENSE Apache-2.0
- google.golang.org/grpc https://github.com/grpc/grpc-go/blob/v1.56.3/LICENSE Apache-2.0
- google.golang.org/protobuf https://github.com/protocolbuffers/protobuf-go/blob/v1.30.0/LICENSE BSD-3-Clause
- gopkg.in/inf.v0 https://github.com/go-inf/inf/blob/v0.9.1/LICENSE BSD-3-Clause
- gopkg.in/ini.v1 https://github.com/go-ini/ini/blob/v1.67.0/LICENSE Apache-2.0
- gopkg.in/yaml.v2 https://github.com/go-yaml/yaml/blob/v2.4.0/LICENSE Apache-2.0
- gopkg.in/yaml.v3 https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE MIT
- k8s.io/api https://github.com/kubernetes/api/blob/v0.26.4/LICENSE Apache-2.0
- k8s.io/apiextensions-apiserver/pkg/apis/apiextensions https://github.com/kubernetes/apiextensions-apiserver/blob/v0.27.1/LICENSE Apache-2.0
- k8s.io/apimachinery/pkg https://github.com/kubernetes/apimachinery/blob/v0.26.4/LICENSE Apache-2.0
- k8s.io/apimachinery/third_party/forked/golang https://github.com/kubernetes/apimachinery/blob/v0.26.4/third_party/forked/golang/LICENSE BSD-3-Clause
- k8s.io/client-go https://github.com/kubernetes/client-go/blob/v0.26.4/LICENSE Apache-2.0
- k8s.io/component-base/config https://github.com/kubernetes/component-base/blob/v0.27.1/LICENSE Apache-2.0
- k8s.io/cri-api/pkg/apis/runtime/v1 https://github.com/kubernetes/cri-api/blob/v0.27.1/LICENSE Apache-2.0
- k8s.io/klog/v2 https://github.com/kubernetes/klog/blob/v2.100.1/LICENSE Apache-2.0
- k8s.io/kube-openapi/pkg https://github.com/kubernetes/kube-openapi/blob/8b0f38b5fd1f/LICENSE Apache-2.0
- k8s.io/kube-openapi/pkg/internal/third_party/go-json-experiment/json https://github.com/kubernetes/kube-openapi/blob/8b0f38b5fd1f/pkg/internal/third_party/go-json-experiment/json/LICENSE BSD-3-Clause
- k8s.io/kube-openapi/pkg/validation/spec https://github.com/kubernetes/kube-openapi/blob/8b0f38b5fd1f/pkg/validation/spec/LICENSE Apache-2.0
- k8s.io/utils https://github.com/kubernetes/utils/blob/9f6742963106/LICENSE Apache-2.0
- k8s.io/utils/internal/third_party/forked/golang/net https://github.com/kubernetes/utils/blob/9f6742963106/internal/third_party/forked/golang/LICENSE BSD-3-Clause
- sigs.k8s.io/controller-runtime https://github.com/kubernetes-sigs/controller-runtime/blob/v0.14.6/LICENSE Apache-2.0
- sigs.k8s.io/json https://github.com/kubernetes-sigs/json/blob/bc3834ca7abd/LICENSE Apache-2.0
- sigs.k8s.io/structured-merge-diff/v4 https://github.com/kubernetes-sigs/structured-merge-diff/blob/v4.2.3/LICENSE Apache-2.0
- sigs.k8s.io/yaml https://github.com/kubernetes-sigs/yaml/blob/v1.3.0/LICENSE MIT