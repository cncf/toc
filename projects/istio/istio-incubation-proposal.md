### Name of the project

Istio

### Description

Istio is an open source service mesh that transparently provides a uniform and efficient way to secure, connect, and monitor services in cloud native applications. It provides zero-trust networking, policy enforcement, traffic management, load balancing, and monitoring, without requiring applications to be rewritten.

Features include:

* Secure service-to-service communication based on zero-trust principles, using TLS encryption and strong identity-based authentication and authorization
* Automatic load balancing for HTTP, gRPC, WebSocket, and TCP traffic
* Fine-grained control of traffic behavior with rich routing rules, retries, failovers, and fault injection
* A policy API supporting access controls, rate limits and quotas, that can be connected to external authorization systems
* Automatic generation of metrics, logs, and traces for all traffic within a cluster, including cluster ingress and egress
* A data plane platform allowing for extension and customization of proxy behavior

Istio is designed for extensibility, and can handle a diverse range of deployment needs. Istio’s control plane runs on Kubernetes, and is capable of managing workloads running in the same Kubernetes cluster, other Kubernetes clusters, and even workloads running outside Kubernetes on VMs.

### Sponsor / Advisor from TOC

Dave Zolotusky, with help from Erin Boyd (TOC) and Lee Calcote (TAG Network)

### License

Istio is licensed under the Apache 2.0 license.

### Source control repositories

* https://github.com/istio (everything in the org)
* https://github.com/istio-ecosystem (everything in the org)

### External dependencies

Istio has external dependencies licensed under the following CNCF-approved licenses:

* Apache 2.0
* BSD License
* ISC License
* MIT License

Five modules are licensed under the Mozilla Public License, and all have been [granted exceptions](https://github.com/cncf/foundation/blob/main/license-exceptions/cncf-exceptions-2019-11-01.spdx) by the CNCF governing board.

Licenses for most Istio dependencies are automatically included at https://github.com/istio/istio/tree/master/licenses. A snapshot of modules and their corresponding licenses [can be found below](#licenses).

Projects in istio-ecosystem will be brought in line with CNCF guidelines before submission.

### Issue tracker

https://github.com/istio/istio/issues

### Committers

Commits to Istio, no matter how small, are welcome from anyone in the community. 

Istio uses "maintainer" as the noun for a member allowed to approve code contributions. See https://github.com/istio/community/blob/master/ROLES.md#maintainer for more information.

* Members are mapped to teams in https://github.com/istio/community/blob/master/org/teams.yaml.  A current list of all maintainers is at the top of the file.
* Teams are mapped to project paths in https://github.com/istio/istio/blob/master/CODEOWNERS.

A dashboard of Istio maintainers can be seen at https://eng.istio.io/maintainers.

### Infrastructure requirements (CI / CNCF Cluster)

The project uses Prow infrastructure run by Google Cloud. Artifacts are stored in Google Cloud Storage, and containers published to registries including Docker Hub and Google Container Registry. 

Google will migrate ownership of this infrastructure to the CNCF and, along with other community members, sponsor it going forward.

### Statement on alignment with CNCF mission

Istio aims to provide security, observability and policy for traffic on top of Kubernetes. If the goal of the CNCF is to make cloud native computing ubiquitous, then a goal of Istio could be expressed as making that ubiquity safe and easy to manage.

Istio is a service mesh, and service meshes are called out in the CNCF charter as exemplifying the Cloud Native approach.

Istio stands on the shoulders of several CNCF projects, such as Kubernetes, Envoy, gRPC, Prometheus and SPIFFE. It can be installed with Helm, and integrations exist for projects including Knative, Flagger, Jaeger, Open Policy Agent, and OpenTelemetry.

### Communication channels

* Slack: https://slack.istio.io
* Discourse: https://discuss.istio.io
* Shared Drive: https://drive.google.com/corp/drive/folders/0ADmbrU7ueGOUUk9PVA

### Website

https://istio.io/

### Release methodology and mechanics

We produce new builds of Istio for each commit. Around once per quarter, we build a minor release and run through several additional tests as well as release qualification. We release patch versions for issues found in minor releases.

Detailed information on Istio's release management can be found in [its documentation](https://istio.io/latest/docs/releases/supported-releases/) and [its developer wiki](https://github.com/istio/istio/wiki/Istio-Release-Management).

### Security processes

Istio has a responsible and mature vulnerability submission and disclosure program, managed by the Product Security Working Group.  The processes are documented at https://istio.io/latest/docs/releases/security-vulnerabilities/.

Vendors meeting certain published criteria receive advance notice of security vulnerabilities through Istio's early disclosure list, which is documented at https://github.com/istio/community/blob/master/EARLY-DISCLOSURE.md.

### Roadmap

[Our 2022 roadmap](https://docs.google.com/document/d/1gfruq22BXKYgKA-BttHv2gE-woq6lu1kreA8VyIdQXc/edit#heading=h.gjdgxs) highlights the following goals:

* Stabilizing Istio's API surface
* Improving upgrades and troubleshooting
* Enhancing extensibility
* Expanding reach
* Further security hardening

### Social media account

* Twitter: https://twitter.com/istiomesh
* YouTube: https://www.youtube.com/c/istio

### Community size

Istio has had over [8000 individual contributors](https://istio.teststats.cncf.io/d/18/overall-project-statistics-table?orgId=1) on the project’s GitHub repositories. 

The project has over [32,000 GitHub stars](https://istio.teststats.cncf.io/d/18/overall-project-statistics-table?orgId=1) across its repositories.

[Over 70 companies](https://istio.teststats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20decade&var-metric=prs) have submitted at least 8 Pull Requests to Istio.

#### Contribution

Istio's proud parents are Google and IBM, alongside the Lyft team who built the Envoy proxy.

Significant contributions come from:

* US technology companies, including Red Hat, Cisco, VMware, and Intel
* Chinese cloud providers and vendors, including Huawei, Tencent, Alibaba, and DaoCloud
* Companies founded to bring Istio solutions to market, including Tetrate, Aspen Mesh, and Solo.io
* Istio's end users, including Auto Trader UK, Salesforce, SAP, and Yahoo!

#### End users

According to the CNCF annual survey, Istio is the most widely adopted service mesh.  A list of [public reference customers can be found on the istio website](https://istio.io/latest/about/case-studies/), including a number of case studies in text or video format.

The first annual IstioCon event, held online in 2021, had over 4,000 attendees. The project also hosts monthly community meetups.

### Licenses

At the time of project submission, the following modules were used with the following licenses.  For an up-to-date list, please refer to the [go.mod](https://github.com/istio/istio/blob/master/go.mod) file and licenses directory in each repository.

| Module                                                                                                          | Apache     | BSD | ISC | MIT | Mozilla |
| --------------------------------------------------------------------------------------------------------------- | ---------- | --- | --- | --- | ------- |
| [cloud.google.com/go](http://cloud.google.com/go)                                                               | Apache 2.0 |     |     |     |         |
| [cloud.google.com/go/compute](http://cloud.google.com/go/compute)                                               | Apache 2.0 |     |     |     |         |
| [cloud.google.com/go/logging](http://cloud.google.com/go/logging)                                               | Apache 2.0 |     |     |     |         |
| [cloud.google.com/go/security](http://cloud.google.com/go/security)                                             | Apache 2.0 |     |     |     |         |
| [contrib.go.opencensus.io/exporter/prometheus](http://contrib.go.opencensus.io/exporter/prometheus)             | Apache 2.0 |     |     |     |         |
| [github.com/AdaLogics/go-fuzz-headers](http://github.com/AdaLogics/go-fuzz-headers)                             | Apache 2.0 |     |     |     |         |
| [github.com/antlr/antlr4/runtime/Go/antlr](http://github.com/antlr/antlr4/runtime/Go/antlr)                     |            | BSD |     |     |         |
| [github.com/asaskevich/govalidator](http://github.com/asaskevich/govalidator)                                   |            |     |     | MIT |         |
| [github.com/Azure/go-ansiterm](http://github.com/Azure/go-ansiterm)                                             |            |     |     | MIT |         |
| [github.com/Azure/go-autorest](http://github.com/Azure/go-autorest)                                             | Apache 2.0 |     |     |     |         |
| [github.com/Azure/go-autorest/autorest](http://github.com/Azure/go-autorest/autorest)                           | Apache 2.0 |     |     |     |         |
| [github.com/Azure/go-autorest/autorest/adal](http://github.com/Azure/go-autorest/autorest/adal)                 | Apache 2.0 |     |     |     |         |
| [github.com/Azure/go-autorest/autorest/date](http://github.com/Azure/go-autorest/autorest/date)                 | Apache 2.0 |     |     |     |         |
| [github.com/Azure/go-autorest/logger](http://github.com/Azure/go-autorest/logger)                               | Apache 2.0 |     |     |     |         |
| [github.com/Azure/go-autorest/tracing](http://github.com/Azure/go-autorest/tracing)                             | Apache 2.0 |     |     |     |         |
| [github.com/beorn7/perks](http://github.com/beorn7/perks)                                                       |            |     |     | MIT |         |
| [github.com/BurntSushi/toml](http://github.com/BurntSushi/toml)                                                 |            |     |     | MIT |         |
| [github.com/cenkalti/backoff/v4](http://github.com/cenkalti/backoff/v4)                                         |            |     |     | MIT |         |
| [github.com/census-instrumentation/opencensus-proto](http://github.com/census-instrumentation/opencensus-proto) | Apache 2.0 |     |     |     |         |
| [github.com/cespare/xxhash](http://github.com/cespare/xxhash)                                                   |            |     |     | MIT |         |
| [github.com/cespare/xxhash/v2](http://github.com/cespare/xxhash/v2)                                             |            |     |     | MIT |         |
| [github.com/chai2010/gettext-go](http://github.com/chai2010/gettext-go)                                         |            | BSD |     |     |         |
| [github.com/cheekybits/genny](http://github.com/cheekybits/genny)                                               |            |     |     | MIT |         |
| [github.com/cheggaaa/pb/v3](http://github.com/cheggaaa/pb/v3)                                                   |            | BSD |     |     |         |
| [github.com/cncf/udpa/go](http://github.com/cncf/udpa/go)                                                       | Apache 2.0 |     |     |     |         |
| [github.com/cncf/xds/go](http://github.com/cncf/xds/go)                                                         | Apache 2.0 |     |     |     |         |
| [github.com/containerd/stargz-snapshotter/estargz](http://github.com/containerd/stargz-snapshotter/estargz)     | Apache 2.0 |     |     |     |         |
| [github.com/containerd/typeurl](http://github.com/containerd/typeurl)                                           | Apache 2.0 |     |     |     |         |
| [github.com/containernetworking/cni](http://github.com/containernetworking/cni)                                 | Apache 2.0 |     |     |     |         |
| [github.com/containernetworking/plugins](http://github.com/containernetworking/plugins)                         | Apache 2.0 |     |     |     |         |
| [github.com/coreos/go-oidc/v3](http://github.com/coreos/go-oidc/v3)                                             | Apache 2.0 |     |     |     |         |
| [github.com/cpuguy83/go-md2man/v2](http://github.com/cpuguy83/go-md2man/v2)                                     |            |     |     | MIT |         |
| [github.com/cyphar/filepath-securejoin](http://github.com/cyphar/filepath-securejoin)                           |            | BSD |     |     |         |
| [github.com/davecgh/go-spew](http://github.com/davecgh/go-spew)                                                 |            |     | ISC |     |         |
| [github.com/decred/dcrd/dcrec/secp256k1/v4](http://github.com/decred/dcrd/dcrec/secp256k1/v4)                   |            |     | ISC |     |         |
| [github.com/docker/cli](http://github.com/docker/cli)                                                           | Apache 2.0 |     |     |     |         |
| [github.com/docker/distribution](http://github.com/docker/distribution)                                         | Apache 2.0 |     |     |     |         |
| [github.com/docker/docker](http://github.com/docker/docker)                                                     | Apache 2.0 |     |     |     |         |
| [github.com/docker/docker-credential-helpers](http://github.com/docker/docker-credential-helpers)               |            |     |     | MIT |         |
| [github.com/envoyproxy/go-control-plane](http://github.com/envoyproxy/go-control-plane)                         | Apache 2.0 |     |     |     |         |
| [github.com/envoyproxy/protoc-gen-validate](http://github.com/envoyproxy/protoc-gen-validate)                   | Apache 2.0 |     |     |     |         |
| [github.com/evanphx/json-patch](http://github.com/evanphx/json-patch)                                           |            | BSD |     |     |         |
| [github.com/evanphx/json-patch/v5](http://github.com/evanphx/json-patch/v5)                                     |            | BSD |     |     |         |
| [github.com/exponent-io/jsonpath](http://github.com/exponent-io/jsonpath)                                       |            |     |     | MIT |         |
| [github.com/fatih/camelcase](http://github.com/fatih/camelcase)                                                 |            |     |     | MIT |         |
| [github.com/fatih/color](http://github.com/fatih/color)                                                         |            |     |     | MIT |         |
| [github.com/felixge/fgprof](http://github.com/felixge/fgprof)                                                   |            |     |     | MIT |         |
| [github.com/florianl/go-nflog/v2](http://github.com/florianl/go-nflog/v2)                                       |            |     |     | MIT |         |
| [github.com/fsnotify/fsnotify](http://github.com/fsnotify/fsnotify)                                             |            | BSD |     |     |         |
| [github.com/fvbommel/sortorder](http://github.com/fvbommel/sortorder)                                           |            |     |     | MIT |         |
| [github.com/go-errors/errors](http://github.com/go-errors/errors)                                               |            |     |     | MIT |         |
| [github.com/go-kit/log](http://github.com/go-kit/log)                                                           |            |     |     | MIT |         |
| [github.com/go-logfmt/logfmt](http://github.com/go-logfmt/logfmt)                                               |            |     |     | MIT |         |
| [github.com/go-logr/logr](http://github.com/go-logr/logr)                                                       | Apache 2.0 |     |     |     |         |
| [github.com/go-openapi/jsonpointer](http://github.com/go-openapi/jsonpointer)                                   | Apache 2.0 |     |     |     |         |
| [github.com/go-openapi/jsonreference](http://github.com/go-openapi/jsonreference)                               | Apache 2.0 |     |     |     |         |
| [github.com/go-openapi/swag](http://github.com/go-openapi/swag)                                                 | Apache 2.0 |     |     |     |         |
| [github.com/go-task/slim-sprig](http://github.com/go-task/slim-sprig)                                           |            |     |     | MIT |         |
| [github.com/gobwas/glob](http://github.com/gobwas/glob)                                                         |            |     |     | MIT |         |
| [github.com/goccy/go-json](http://github.com/goccy/go-json)                                                     |            |     |     | MIT |         |
| [github.com/gogo/protobuf](http://github.com/gogo/protobuf)                                                     |            | BSD |     |     |         |
| [github.com/golang-jwt/jwt/v4](http://github.com/golang-jwt/jwt/v4)                                             |            |     |     | MIT |         |
| [github.com/golang/groupcache](http://github.com/golang/groupcache)                                             | Apache 2.0 |     |     |     |         |
| [github.com/golang/protobuf](http://github.com/golang/protobuf)                                                 |            | BSD |     |     |         |
| [github.com/google/btree](http://github.com/google/btree)                                                       | Apache 2.0 |     |     |     |         |
| [github.com/google/cel-go](http://github.com/google/cel-go)                                                     | Apache 2.0 | BSD |     |     |         |
| [github.com/google/go-cmp](http://github.com/google/go-cmp)                                                     |            | BSD |     |     |         |
| [github.com/google/go-containerregistry](http://github.com/google/go-containerregistry)                         | Apache 2.0 |     |     |     |         |
| [github.com/google/gofuzz](http://github.com/google/gofuzz)                                                     | Apache 2.0 |     |     |     |         |
| [github.com/google/pprof](http://github.com/google/pprof)                                                       | Apache 2.0 |     |     |     |         |
| [github.com/google/shlex](http://github.com/google/shlex)                                                       | Apache 2.0 |     |     |     |         |
| [github.com/google/uuid](http://github.com/google/uuid)                                                         |            | BSD |     |     |         |
| [github.com/googleapis/gax-go/v2](http://github.com/googleapis/gax-go/v2)                                       |            | BSD |     |     |         |
| [github.com/googleapis/gnostic](http://github.com/googleapis/gnostic)                                           | Apache 2.0 |     |     |     |         |
| [github.com/gorilla/mux](http://github.com/gorilla/mux)                                                         |            | BSD |     |     |         |
| [github.com/gorilla/websocket](http://github.com/gorilla/websocket)                                             |            | BSD |     |     |         |
| [github.com/gregjones/httpcache](http://github.com/gregjones/httpcache)                                         |            |     |     | MIT |         |
| [github.com/grpc-ecosystem/go-grpc-middleware](http://github.com/grpc-ecosystem/go-grpc-middleware)             | Apache 2.0 |     |     |     |         |
| [github.com/grpc-ecosystem/go-grpc-prometheus](http://github.com/grpc-ecosystem/go-grpc-prometheus)             | Apache 2.0 |     |     |     |         |
| [github.com/hashicorp/errwrap](http://github.com/hashicorp/errwrap)                                             |            |     |     |     | MPL     |
| [github.com/hashicorp/go-multierror](http://github.com/hashicorp/go-multierror)                                 |            |     |     |     | MPL     |
| [github.com/hashicorp/go-version](http://github.com/hashicorp/go-version)                                       |            |     |     |     | MPL     |
| [github.com/hashicorp/golang-lru](http://github.com/hashicorp/golang-lru)                                       |            |     |     |     | MPL     |
| [github.com/hashicorp/hcl](http://github.com/hashicorp/hcl)                                                     |            |     |     |     | MPL     |
| [github.com/huandu/xstrings](http://github.com/huandu/xstrings)                                                 |            |     |     | MIT |         |
| [github.com/imdario/mergo](http://github.com/imdario/mergo)                                                     |            | BSD |     |     |         |
| [github.com/inconshreveable/mousetrap](http://github.com/inconshreveable/mousetrap)                             | Apache 2.0 |     |     |     |         |
| [github.com/jonboulle/clockwork](http://github.com/jonboulle/clockwork)                                         | Apache 2.0 |     |     |     |         |
| [github.com/josharian/intern](http://github.com/josharian/intern)                                               |            |     |     | MIT |         |
| [github.com/josharian/native](http://github.com/josharian/native)                                               |            |     |     | MIT |         |
| [github.com/json-iterator/go](http://github.com/json-iterator/go)                                               |            |     |     | MIT |         |
| [github.com/klauspost/compress](http://github.com/klauspost/compress)                                           | Apache 2.0 | BSD |     | MIT |         |
| [github.com/kr/pretty](http://github.com/kr/pretty)                                                             |            |     |     | MIT |         |
| [github.com/kr/text](http://github.com/kr/text)                                                                 |            |     |     | MIT |         |
| [github.com/kylelemons/godebug](http://github.com/kylelemons/godebug)                                           | Apache 2.0 |     |     |     |         |
| [github.com/lestrrat-go/backoff/v2](http://github.com/lestrrat-go/backoff/v2)                                   |            |     |     | MIT |         |
| [github.com/lestrrat-go/blackmagic](http://github.com/lestrrat-go/blackmagic)                                   |            |     |     | MIT |         |
| [github.com/lestrrat-go/httpcc](http://github.com/lestrrat-go/httpcc)                                           |            |     |     | MIT |         |
| [github.com/lestrrat-go/iter](http://github.com/lestrrat-go/iter)                                               |            |     |     | MIT |         |
| [github.com/lestrrat-go/jwx](http://github.com/lestrrat-go/jwx)                                                 |            |     |     | MIT |         |
| [github.com/lestrrat-go/option](http://github.com/lestrrat-go/option)                                           |            |     |     | MIT |         |
| [github.com/liggitt/tabwriter](http://github.com/liggitt/tabwriter)                                             |            | BSD |     |     |         |
| [github.com/lucas-clemente/quic-go](http://github.com/lucas-clemente/quic-go)                                   |            |     |     | MIT |         |
| [github.com/magiconair/properties](http://github.com/magiconair/properties)                                     |            | BSD |     |     |         |
| [github.com/mailru/easyjson](http://github.com/mailru/easyjson)                                                 |            |     |     | MIT |         |
| [github.com/MakeNowJust/heredoc](http://github.com/MakeNowJust/heredoc)                                         |            |     |     | MIT |         |
| [github.com/marten-seemann/qpack](http://github.com/marten-seemann/qpack)                                       |            |     |     | MIT |         |
| [github.com/marten-seemann/qtls-go1-16](http://github.com/marten-seemann/qtls-go1-16)                           |            | BSD |     |     |         |
| [github.com/marten-seemann/qtls-go1-17](http://github.com/marten-seemann/qtls-go1-17)                           |            | BSD |     |     |         |
| [github.com/marten-seemann/qtls-go1-18](http://github.com/marten-seemann/qtls-go1-18)                           |            | BSD |     |     |         |
| [github.com/Masterminds/goutils](http://github.com/Masterminds/goutils)                                         | Apache 2.0 |     |     |     |         |
| [github.com/Masterminds/semver/v3](http://github.com/Masterminds/semver/v3)                                     |            |     |     | MIT |         |
| [github.com/Masterminds/sprig/v3](http://github.com/Masterminds/sprig/v3)                                       |            |     |     | MIT |         |
| [github.com/mattn/go-colorable](http://github.com/mattn/go-colorable)                                           |            |     |     | MIT |         |
| [github.com/mattn/go-isatty](http://github.com/mattn/go-isatty)                                                 |            |     |     | MIT |         |
| [github.com/mattn/go-runewidth](http://github.com/mattn/go-runewidth)                                           |            |     |     | MIT |         |
| [github.com/matttproud/golang\_protobuf\_extensions](http://github.com/matttproud/golang_protobuf_extensions)   | Apache 2.0 |     |     |     |         |
| [github.com/mdlayher/netlink](http://github.com/mdlayher/netlink)                                               |            |     |     | MIT |         |
| [github.com/mdlayher/socket](http://github.com/mdlayher/socket)                                                 |            |     |     | MIT |         |
| [github.com/miekg/dns](http://github.com/miekg/dns)                                                             |            | BSD |     |     |         |
| [github.com/mitchellh/copystructure](http://github.com/mitchellh/copystructure)                                 |            |     |     | MIT |         |
| [github.com/mitchellh/go-homedir](http://github.com/mitchellh/go-homedir)                                       |            |     |     | MIT |         |
| [github.com/mitchellh/go-wordwrap](http://github.com/mitchellh/go-wordwrap)                                     |            |     |     | MIT |         |
| [github.com/mitchellh/mapstructure](http://github.com/mitchellh/mapstructure)                                   |            |     |     | MIT |         |
| [github.com/mitchellh/reflectwalk](http://github.com/mitchellh/reflectwalk)                                     |            |     |     | MIT |         |
| [github.com/moby/buildkit](http://github.com/moby/buildkit)                                                     | Apache 2.0 |     |     |     |         |
| [github.com/moby/spdystream](http://github.com/moby/spdystream)                                                 | Apache 2.0 |     |     |     |         |
| [github.com/moby/term](http://github.com/moby/term)                                                             | Apache 2.0 |     |     |     |         |
| [github.com/modern-go/concurrent](http://github.com/modern-go/concurrent)                                       | Apache 2.0 |     |     |     |         |
| [github.com/modern-go/reflect2](http://github.com/modern-go/reflect2)                                           | Apache 2.0 |     |     |     |         |
| [github.com/monochromegane/go-gitignore](http://github.com/monochromegane/go-gitignore)                         |            |     |     | MIT |         |
| [github.com/natefinch/lumberjack](http://github.com/natefinch/lumberjack)                                       |            |     |     | MIT |         |
| [github.com/nxadm/tail](http://github.com/nxadm/tail)                                                           |            |     |     | MIT |         |
| [github.com/onsi/ginkgo](http://github.com/onsi/ginkgo)                                                         |            |     |     | MIT |         |
| [github.com/onsi/gomega](http://github.com/onsi/gomega)                                                         |            |     |     | MIT |         |
| [github.com/opencontainers/go-digest](http://github.com/opencontainers/go-digest)                               | Apache 2.0 |     |     |     |         |
| [github.com/opencontainers/image-spec](http://github.com/opencontainers/image-spec)                             | Apache 2.0 |     |     |     |         |
| [github.com/openshift/api](http://github.com/openshift/api)                                                     | Apache 2.0 |     |     |     |         |
| [github.com/pelletier/go-toml](http://github.com/pelletier/go-toml)                                             | Apache 2.0 |     |     | MIT |         |
| [github.com/peterbourgon/diskv](http://github.com/peterbourgon/diskv)                                           |            |     |     | MIT |         |
| [github.com/pkg/errors](http://github.com/pkg/errors)                                                           |            | BSD |     |     |         |
| [github.com/pmezard/go-difflib](http://github.com/pmezard/go-difflib)                                           |            | BSD |     |     |         |
| [github.com/prometheus/client\_golang](http://github.com/prometheus/client_golang)                              | Apache 2.0 |     |     |     |         |
| [github.com/prometheus/client\_model](http://github.com/prometheus/client_model)                                | Apache 2.0 |     |     |     |         |
| [github.com/prometheus/common](http://github.com/prometheus/common)                                             | Apache 2.0 |     |     |     |         |
| [github.com/prometheus/procfs](http://github.com/prometheus/procfs)                                             | Apache 2.0 |     |     |     |         |
| [github.com/prometheus/prom2json](http://github.com/prometheus/prom2json)                                       | Apache 2.0 |     |     |     |         |
| [github.com/prometheus/prometheus](http://github.com/prometheus/prometheus)                                     | Apache 2.0 |     |     |     |         |
| [github.com/prometheus/statsd\_exporter](http://github.com/prometheus/statsd_exporter)                          | Apache 2.0 |     |     |     |         |
| [github.com/PuerkitoBio/purell](http://github.com/PuerkitoBio/purell)                                           |            | BSD |     |     |         |
| [github.com/PuerkitoBio/urlesc](http://github.com/PuerkitoBio/urlesc)                                           |            | BSD |     |     |         |
| [github.com/rivo/uniseg](http://github.com/rivo/uniseg)                                                         |            |     |     | MIT |         |
| [github.com/rogpeppe/go-internal](http://github.com/rogpeppe/go-internal)                                       |            | BSD |     |     |         |
| [github.com/russross/blackfriday](http://github.com/russross/blackfriday)                                       |            | BSD |     |     |         |
| [github.com/russross/blackfriday/v2](http://github.com/russross/blackfriday/v2)                                 |            | BSD |     |     |         |
| [github.com/ryanuber/go-glob](http://github.com/ryanuber/go-glob)                                               |            |     |     | MIT |         |
| [github.com/shopspring/decimal](http://github.com/shopspring/decimal)                                           |            |     |     | MIT |         |
| [github.com/sirupsen/logrus](http://github.com/sirupsen/logrus)                                                 |            |     |     | MIT |         |
| [github.com/spaolacci/murmur3](http://github.com/spaolacci/murmur3)                                             |            | BSD |     |     |         |
| [github.com/spf13/afero](http://github.com/spf13/afero)                                                         | Apache 2.0 |     |     |     |         |
| [github.com/spf13/cast](http://github.com/spf13/cast)                                                           |            |     |     | MIT |         |
| [github.com/spf13/cobra](http://github.com/spf13/cobra)                                                         | Apache 2.0 |     |     |     |         |
| [github.com/spf13/jwalterweatherman](http://github.com/spf13/jwalterweatherman)                                 |            |     |     | MIT |         |
| [github.com/spf13/pflag](http://github.com/spf13/pflag)                                                         |            | BSD |     |     |         |
| [github.com/spf13/viper](http://github.com/spf13/viper)                                                         |            |     |     | MIT |         |
| [github.com/stoewer/go-strcase](http://github.com/stoewer/go-strcase)                                           |            |     |     | MIT |         |
| [github.com/stretchr/testify](http://github.com/stretchr/testify)                                               |            |     |     | MIT |         |
| [github.com/subosito/gotenv](http://github.com/subosito/gotenv)                                                 |            |     |     | MIT |         |
| [github.com/vbatts/tar-split](http://github.com/vbatts/tar-split)                                               |            | BSD |     |     |         |
| [github.com/vishvananda/netlink](http://github.com/vishvananda/netlink)                                         | Apache 2.0 |     |     |     |         |
| [github.com/vishvananda/netns](http://github.com/vishvananda/netns)                                             | Apache 2.0 |     |     |     |         |
| [github.com/VividCortex/ewma](http://github.com/VividCortex/ewma)                                               |            |     |     | MIT |         |
| [github.com/xeipuuv/gojsonpointer](http://github.com/xeipuuv/gojsonpointer)                                     | Apache 2.0 |     |     |     |         |
| [github.com/xeipuuv/gojsonreference](http://github.com/xeipuuv/gojsonreference)                                 | Apache 2.0 |     |     |     |         |
| [github.com/xeipuuv/gojsonschema](http://github.com/xeipuuv/gojsonschema)                                       | Apache 2.0 |     |     |     |         |
| [github.com/xlab/treeprint](http://github.com/xlab/treeprint)                                                   |            |     |     | MIT |         |
| [github.com/yl2chen/cidranger](http://github.com/yl2chen/cidranger)                                             |            |     |     | MIT |         |
| [go.opencensus.io](http://go.opencensus.io)                                                                     | Apache 2.0 |     |     |     |         |
| [go.opentelemetry.io/proto/otlp](http://go.opentelemetry.io/proto/otlp)                                         | Apache 2.0 |     |     |     |         |
| [go.starlark.net](http://go.starlark.net)                                                                       |            | BSD |     |     |         |
| [go.uber.org/atomic](http://go.uber.org/atomic)                                                                 |            |     |     | MIT |         |
| [go.uber.org/multierr](http://go.uber.org/multierr)                                                             |            |     |     | MIT |         |
| [go.uber.org/zap](http://go.uber.org/zap)                                                                       |            |     |     | MIT |         |
| [golang.org/x/crypto](http://golang.org/x/crypto)                                                               |            | BSD |     |     |         |
| [golang.org/x/mod](http://golang.org/x/mod)                                                                     |            | BSD |     |     |         |
| [golang.org/x/net](http://golang.org/x/net)                                                                     |            | BSD |     |     |         |
| [golang.org/x/oauth2](http://golang.org/x/oauth2)                                                               |            | BSD |     |     |         |
| [golang.org/x/sync](http://golang.org/x/sync)                                                                   |            | BSD |     |     |         |
| [golang.org/x/sys](http://golang.org/x/sys)                                                                     |            | BSD |     |     |         |
| [golang.org/x/term](http://golang.org/x/term)                                                                   |            | BSD |     |     |         |
| [golang.org/x/text](http://golang.org/x/text)                                                                   |            | BSD |     |     |         |
| [golang.org/x/time](http://golang.org/x/time)                                                                   |            | BSD |     |     |         |
| [golang.org/x/tools](http://golang.org/x/tools)                                                                 |            | BSD |     |     |         |
| [golang.org/x/xerrors](http://golang.org/x/xerrors)                                                             |            | BSD |     |     |         |
| [gomodules.xyz/jsonpatch/v2](http://gomodules.xyz/jsonpatch/v2)                                                 | Apache 2.0 |     |     |     |         |
| [gomodules.xyz/jsonpatch/v3](http://gomodules.xyz/jsonpatch/v3)                                                 | Apache 2.0 |     |     |     |         |
| [gomodules.xyz/orderedmap](http://gomodules.xyz/orderedmap)                                                     |            |     |     | MIT |         |
| [google.golang.org/api](http://google.golang.org/api)                                                           |            | BSD |     |     |         |
| [google.golang.org/appengine](http://google.golang.org/appengine)                                               | Apache 2.0 |     |     |     |         |
| [google.golang.org/genproto](http://google.golang.org/genproto)                                                 | Apache 2.0 |     |     |     |         |
| [google.golang.org/grpc](http://google.golang.org/grpc)                                                         | Apache 2.0 |     |     |     |         |
| [google.golang.org/protobuf](http://google.golang.org/protobuf)                                                 |            | BSD |     |     |         |
| [gopkg.in/inf.v0](http://gopkg.in/inf.v0)                                                                       |            | BSD |     |     |         |
| [gopkg.in/ini.v1](http://gopkg.in/ini.v1)                                                                       | Apache 2.0 |     |     |     |         |
| [gopkg.in/square/go-jose.v2](http://gopkg.in/square/go-jose.v2)                                                 | Apache 2.0 |     |     |     |         |
| [gopkg.in/tomb.v1](http://gopkg.in/tomb.v1)                                                                     |            | BSD |     |     |         |
| [gopkg.in/yaml.v2](http://gopkg.in/yaml.v2)                                                                     | Apache 2.0 |     |     |     |         |
| [gopkg.in/yaml.v3](http://gopkg.in/yaml.v3)                                                                     |            |     |     | MIT |         |
| [helm.sh/helm/v3](http://helm.sh/helm/v3)                                                                       | Apache 2.0 |     |     |     |         |
| [istio.io/api](http://istio.io/api)                                                                             | Apache 2.0 |     |     |     |         |
| [istio.io/client-go](http://istio.io/client-go)                                                                 | Apache 2.0 |     |     |     |         |
| [istio.io/gogo-genproto](http://istio.io/gogo-genproto)                                                         | Apache 2.0 |     |     |     |         |
| [istio.io/pkg](http://istio.io/pkg)                                                                             | Apache 2.0 |     |     |     |         |
| [k8s.io/api](http://k8s.io/api)                                                                                 | Apache 2.0 |     |     |     |         |
| [k8s.io/apiextensions-apiserver](http://k8s.io/apiextensions-apiserver)                                         | Apache 2.0 |     |     |     |         |
| [k8s.io/apimachinery](http://k8s.io/apimachinery)                                                               | Apache 2.0 |     |     |     |         |
| [k8s.io/cli-runtime](http://k8s.io/cli-runtime)                                                                 | Apache 2.0 |     |     |     |         |
| [k8s.io/client-go](http://k8s.io/client-go)                                                                     | Apache 2.0 |     |     |     |         |
| [k8s.io/component-base](http://k8s.io/component-base)                                                           | Apache 2.0 |     |     |     |         |
| [k8s.io/klog/v2](http://k8s.io/klog/v2)                                                                         | Apache 2.0 |     |     |     |         |
| [k8s.io/kube-openapi](http://k8s.io/kube-openapi)                                                               | Apache 2.0 |     |     |     |         |
| [k8s.io/kubectl](http://k8s.io/kubectl)                                                                         | Apache 2.0 |     |     |     |         |
| [k8s.io/utils](http://k8s.io/utils)                                                                             | Apache 2.0 |     |     |     |         |
| [sigs.k8s.io/controller-runtime](http://sigs.k8s.io/controller-runtime)                                         | Apache 2.0 |     |     |     |         |
| [sigs.k8s.io/gateway-api](http://sigs.k8s.io/gateway-api)                                                       | Apache 2.0 |     |     |     |         |
| [sigs.k8s.io/json](http://sigs.k8s.io/json)                                                                     | Apache 2.0 | BSD |     |     |         |
| [sigs.k8s.io/kustomize/api](http://sigs.k8s.io/kustomize/api)                                                   | Apache 2.0 |     |     |     |         |
| [sigs.k8s.io/kustomize/kyaml](http://sigs.k8s.io/kustomize/kyaml)                                               | Apache 2.0 |     |     |     |         |
| [sigs.k8s.io/mcs-api](http://sigs.k8s.io/mcs-api)                                                               | Apache 2.0 |     |     |     |         |
| [sigs.k8s.io/structured-merge-diff/v4](http://sigs.k8s.io/structured-merge-diff/v4)                             | Apache 2.0 |     |     |     |         |
| [sigs.k8s.io/yaml](http://sigs.k8s.io/yaml)                                                                     |            | BSD |     | MIT |         |
