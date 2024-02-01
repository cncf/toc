# Kuma proposal for incubation

## **About Kuma**

Kuma is a platform agnostic open-source control plane for service mesh and microservices.
It can run and be operated natively across both Kubernetes and VM environments, making it easy to adopt by every team in the organization.

Bundling Envoy as a data-plane, Kuma can instrument any L4/L7 traffic to secure, observe, route and enhance connectivity between any service or database.
It can be used natively in Kubernetes via CRDs or via a RESTful API across other environments like VMs and Bare Metal.

Kuma provides the best multi-zone support in the industry to run a distributed service mesh across clouds, clusters,
including hybrid Kubernetes + VM meshes, with native cluster synchronization and automatic cross-zone connectivity.

**Key Features**:

- **Hybrid Universal Mode** - By automating and abstracting connectivity across enterprise environments, hybrid universal mode enables service meshes to run and communicate simultaneously on heterogeneous environments out-of-the box, 
including VMs, data centers and multiple Kubernetes clusters, on both standalone and multi-zone deployments.
- **Multi-Zone Deployment** - To deploy a service mesh across multiple Kubernetes clusters, VMs,
containerized environments like AWS ECS and Fargate, and enable cross-zone connectivity across the entire organization.
- **Advanced Control Plane Replication** - With global and remote control planes, Kuma automatically synchronizes service mesh policies
across a distributed environment out of the box across clouds, on both Kubernetes and VMs.
- **Ingress Data Plane Mode** - Automates cross-platform and cross-cluster service mesh communication out-of-the-box,
which previously would take months of set-up effort for enterprises.
- **Built-in Service Discovery** - A native universal DNS service discovery API enables operational visibility 
by abstracting away underlying services running across multiple platforms and zones, as if these were all running in one zone.
- **Zero Trust Security** - Provide identity and encryption to all service traffic, from L4 to L7,
with various CA backends (including built-in), and automatic certificate rotation. Traffic permissions can also be easily set to ensure appropriate access control.
- **Sophisticated Routing & Control** - Provide fine-grained traffic control capabilities, such as circuit breakers and health checks,
to enhance L4 routing. Easily implement blue/green deployments and canary releases to alleviate the stress of deploying new versions and/or features into your service mesh,
including zone traffic shifting.
- **Observability with Prometheus and Grafana** - Natively integrate with Prometheus for auto-discovery and metrics collection,
and then use Grafana dashboards to monitor performance and ensure service mesh health.
- **L7 tracing** - Identify system failures and performance Kuma’s Mesh’s Traffic-Trace policy.
- **Attribute-based segmentation** - Group services by attributes to efficiently manage operations at scale.
- **Kubernetes-native support** - Natively manage services running in Kubernetes through CRDs to ensure a single source of truth for configurations.

## **Sandbox project acceptance**

Kuma was accepted as a Sandbox project in June 2020: [Sanbox proposal](https://github.com/cncf/toc/blob/main/proposals/sandbox/kuma.md).

## **Progress since sandbox**

Feature Summary:

**2023**:
- Completely re-written suite of policies, with greater flexibility for matching
- Flexible ability to filter API output by metadata
- Many L7 traffic manipulation improvements and features added
- OpenTelemetry support for logs, traces, and metrics
- New proxy patch policy with full JSONPatch support
- New load balancing policy with greater flexibility and power
- Ability to do offline token signing with pre-generated tokens and public key

**2022**:
- Second major release (2.0)
- Zone egress introduced for better control of outbound traffic
- New builtin gateway (based on Envoy) capability
- eBPF support for transparent proxying (init container & CNI)
- Added troubleshooting tooling in the CLI / GUI
- Kubernetes gateway API support
- Completely re-written transparent proxy for traffic redirection
- Cross-mesh, cross-zone communication via the builtin gateway
- Support for ARM and MacOS environments
- Ability to customize sidecar and init containers with JSON patches
- Single-command observability installation
- Completely re-written CNI plugin
- Customizable TLS versions and ciphers across the project
- Metrics filtering to avoid overwhelming Prometheus collector
- Projected service account support in Kubernetes for more securely scoped dataplane identity

**2021**:
- Multiple protocol fault injection support
- Multiple new load balancing algorithms added
- Support for Kubernetes jobs & headless services
- New L7 traffic routing policy
- Datadog support
- Significant performance improvements
- Dynamic / live service map plugin for Grafana
- New rate limiting policy
- Multiple new graphs and charts in Grafana dashboards
- Timeout policy added
- IPV6 support
- Full demo app added
- mTLS permissive mode
- New virtual outbound policy for better custom DNS control

**2020 (since incubation, 6/20)**:
- First major release (1.0)
- StatefulSet support
- Official Helm charts
- gRPC support
- Zone-aware load balancing
- Kafka support
- UI overhaul
- AWS ECS & Fargate support
- Transparent proxying (traffic redirection) in universal (VM) mode
- Retry policy added

## **Governance**
Kuma maintains an open governance policy in accordance with CNCF guidelines.
The [governance policy](https://github.com/kumahq/kuma/blob/master/GOVERNANCE.md), [contributing guidelines](https://github.com/kumahq/kuma/blob/master/CONTRIBUTING.md),
and [security processes](https://github.com/kumahq/kuma/blob/master/SECURITY.md) are all documented.

Kuma hosts [public community meetings](https://docs.google.com/document/d/1HgnK3wJIEg8uFlivdrhrPZYWTpElWWu3mhFDXj-bMWQ/edit) on a monthly basis.
Over 200 individuals have attended these calls with ongoing participation from 28 people from ~14 organizations. 

## **Licenses and Dependencies**

A full list of licenses and dependencies are automatically generated by [this github action](https://github.com/kumahq/kuma/actions/workflows/bom.yaml) and an export is in the annexes. 

## **Project Health & Community stats**

Kuma has a [CLOMonitor](https://clomonitor.io/projects/cncf/kuma) score of 89, passes [OpenSSF Best Practices](https://bestpractices.coreinfrastructure.org/en/projects/5576), and has an A+ on [Go Report Card](https://goreportcard.com/report/github.com/kumahq/kuma).

Kuma maintains a [Slack community](https://join.slack.com/t/kuma-mesh/shared_invite/zt-1wi2o3uok-x0KmKJiSzjQy7NgNkC8IAA) of over 1,380 members
(300% growth since Sandbox acceptance)

The Kuma repo has over 3,400 Stars (+130%) and 300+ Forks on [GitHub](https://github.com/kumahq/kuma).

## **Integrations with the CNCF Ecosystem**

- [Envoy](https://envoyproxy.io)
  - Kuma is built on top of – and is a control-plane for – a fleet of Envoy proxies.
- [Flux (Flagger)](https://docs.flagger.app/tutorials/kuma-progressive-delivery)
  - Built-in integration with Flagger for automated canary releases using Kuma Traffic Routing.
- [Jaeger](https://kuma.io/docs/2.6.x/policies/meshtrace/#zipkin-1)
  Built-in integration with Jaeger / Zipkin for traces emitted from Kuma workloads.
- [Prometheus](https://prometheus.io/docs/prometheus/2.6.x/configuration/configuration/#kuma_sd_config)
  - Built-in integration with Prometheus to auto-configure and scrape Kuma workload and control-plane metrics.
- [Kubernetes](https://kuma.io/docs/2.6.x/production/dp-config/dpp-on-kubernetes/)
  - Many direct and indirect integration points with Kubernetes. First-class support as a deployment platform for Kuma.
- [SPIFFE](https://kuma.io/docs/2.6.x/policies/mutual-tls/#mutual-tls)
  - Kuma utilizes the SPIFFE libraries and ID format for generated workload certificates.
- [CNI](https://kuma.io/docs/2.6.x/production/dp-config/cni/)
  - Kuma provides a CNI plugin as an option for redirecting workload traffic.
- OpenTelemetry
  - Kuma has built-in support for [traces](https://kuma.io/docs/2.6.x/policies/meshtrace/#opentelemetry), [logs](https://kuma.io/docs/2.6.x/policies/meshaccesslog/#opentelemetry), and [metrics](https://kuma.io/docs/2.6.x/policies/meshmetric/#opentelemetry-experimental) being emitted in OTEL format / to an OTEL collector.
- [Meshery](https://docs.meshery.io/service-meshes/adapters/kuma)
  - Meshery includes a built-in adaptor for Kuma.
- [gRPC](https://kuma.io/docs/2.6.x/policies/protocol-support-in-kuma/)
  - Kuma both supports, and is built upon gRPC communication and APIs.

## **Incubation Stage Requirements**

### Adopters in production

- **American Airlines** is leveraging Kuma for their Developer Experience Product, [Runway](https://tech.aa.com/2021-12-21-runway-pt1/).
- **Clearsense** is using Kuma to modernize [healthcare data architecture](https://konghq.com/blog/enterprise/clearsense-healthcare-kong-gateway-kuma).
- **Comcast**
- **Kong** is repackaging Kuma into its enterprise offering "Kong Mesh" to provide service mesh technology to Top Fortune 500 and Top Global 2000 organizations like the top US airlines, top worldwide telcos and [financial institutions](https://www.youtube.com/watch?v=SaPuSXe0JLs&t=736s), and more where Kuma - by proxy - is the underlying technology that powers billions of requests across every industry vertical.
- **Koyeb** is using Kuma for their [Serverless Platform](https://www.koyeb.com/blog/building-a-multi-region-service-mesh-with-kuma-envoy-anycast-bgp-and-mtls).
- **Maersk**
- **MS3** built their enterprise integration platform, [Tavros](https://konghq.com/blog/enterprise/ms3-kong-gateway-kuma), with Kuma which supports 10k+ transactions/second.
- **[Tinymile](https://tinymile.ai/)** is using Kuma to revolutionize the third-party delivery industry.
- **Zomato**

There are 300+ organizations using Kuma including 23 Fortune 500 and 53 Global 2000 companies.

Additional end users are available for private references. 

### Healthy number of committers to the project
Kuma has [90](https://github.com/kumahq/kuma/graphs/contributors) project members from 13 different organizations contributing to the project. There are [8](https://github.com/kumahq/kuma/blob/master/OWNERS.md) maintainers.

### Substantial growth in # of commits and merges

Project growth since sandboxing: 

- [Contribution Activity](https://kuma.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=1590994800000&to=now) - Average of 1,641 commits/month from 28 different contributors
- [Commits by Repository](https://kuma.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&var-period=d7&var-repogroups=All&from=1590994800000&to=now) - 17,000+ commits to date
- [New Contributors](https://kuma.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1496300400000&to=now) - 125 new contributors, 346% growth since Sandbox
- [Issues & PRs Open/Closed](https://kuma.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1590994800000&to=now&var-period=d7&var-repogroup_name=All) - 2,420 (55/month) Issues opened, 1,960 (44/month) Issues closed. 9,980 (232/month) PRs opened, 9,420 (219/month) PRs closed.
- [Contributing Companies](https://kuma.devstats.cncf.io/d/5/companies-table?orgId=1) - 81 different companies (+113%) have made over 73,000 (965%) contributions to Kuma
- [Opened to Merged](https://kuma.devstats.cncf.io/d/16/opened-to-merged?orgId=1&from=now-3y&to=now-1w&var-period=y&var-repogroup_name=All) - Average time to merge on PRs is 7.66 hours.

### Clear versioning scheme

Kuma uses the [Semantic Versioning scheme](https://semver.org/#semantic-versioning-specification-semver) with the format of MAJOR.MINOR.PATCH.
Releases are documented at: https://github.com/kumahq/kuma/releases.

### Clearly documented security processes

[Kuma Security Overview](https://github.com/kumahq/kuma/security) - Security vulnerabilities will be disclosed via release notes, issues and Github advisories with severity score higher than [4.0](https://www.first.org/cvss/calculator/3.1) will have an advisory published.

[Kuma Security Policy](https://github.com/kumahq/kuma/security/policy) - We use Github's Security advisories for reporting security vulnerabilities.

## Statement on Alignment with the CNCF Mission

Kuma believes in simplifying service mesh to be accessible from Day 0, to help enable any organization to be cloud native from the start.
Being built on Envoy from the beginning, we are committed to building upon and furthering adoption of existing cloud native projects.
We foster growth of the ecosystem with support and integrations for 10 other CNCF projects and continue to expand that list.
We believe in open, community-driven development and adoption with a strong commitment to continue diversifying our contributor base,
creating a safe and fair contribution environment,
and expanding the overall footprint of Kuma within the cloud native ecosystem in partnership with the foundation and other CNCF projects. 

## Roadmap

Kuma releases a minor version on a 10 week release cycle.

The roadmap is tracked using milestones: https://github.com/kumahq/kuma/milestones

## Annexes

### Licenses

```yaml
# File extracted from the github action: https://github.com/kumahq/kuma/actions/workflows/bom.yaml
# with the following command: 
# gh run download -n licenses -R kumahq/kuma $(gh run list -b master -R kumahq/kuma -w "Build bom and licenses"  -L1 -s completed --json databaseId --jq '.[0].databaseId'); yq -oyaml '.components[] | select(.type == "library") | [{"name": .name, "version": .version, "license": .evidence.licenses[0].license.id}]' licenses.json > licenses.yaml; rm licenses.json
- name: cirello.io/pglock
  version: v1.14.1
  license: Apache-2.0
- name: dario.cat/mergo
  version: v1.0.0
  license: BSD-3-Clause
- name: github.com/Azure/go-ansiterm
  version: v0.0.0-20230124172434-306776ec8161
  license: MIT
- name: github.com/BurntSushi/toml
  version: v1.3.2
  license: MIT
- name: github.com/Masterminds/goutils
  version: v1.1.1
  license: Apache-2.0
- name: github.com/Masterminds/semver/v3
  version: v3.2.1
  license: MIT
- name: github.com/Masterminds/sprig/v3
  version: v3.2.3
  license: MIT
- name: github.com/Microsoft/go-winio
  version: v0.6.1
  license: MIT
- name: github.com/Microsoft/hcsshim
  version: v0.11.4
  license: MIT
- name: github.com/Nordix/simple-ipam
  version: v1.0.0
  license: Apache-2.0
- name: github.com/asaskevich/govalidator
  version: v0.0.0-20210307081110-f21760c49a8d
  license: MIT
- name: github.com/aws/aws-sdk-go
  version: v1.49.6
  license: Apache-2.0
- name: github.com/bakito/go-log-logr-adapter
  version: v0.0.3-0.20231211113354-bfa42fa7e121
  license: Apache-2.0
- name: github.com/beorn7/perks
  version: v1.0.1
  license: MIT
- name: github.com/boombuler/barcode
  version: v1.0.1
  license: MIT
- name: github.com/cenkalti/backoff/v4
  version: v4.2.1
  license: MIT
- name: github.com/census-instrumentation/opencensus-proto
  version: v0.4.1
  license: Apache-2.0
- name: github.com/cespare/xxhash/v2
  version: v2.2.0
  license: MIT
- name: github.com/cilium/ebpf
  version: v0.12.3
  license: MIT
- name: github.com/cncf/xds/go
  version: v0.0.0-20231109132714-523115ebc101
  license: Apache-2.0
- name: github.com/containerd/containerd
  version: v1.7.11
  license: Apache-2.0
- name: github.com/containerd/log
  version: v0.1.0
  license: Apache-2.0
- name: github.com/containernetworking/cni
  version: v1.1.2
  license: Apache-2.0
- name: github.com/containernetworking/plugins
  version: v1.4.0
  license: Apache-2.0
- name: github.com/cpuguy83/dockercfg
  version: v0.3.1
  license: MIT
- name: github.com/cpuguy83/go-md2man/v2
  version: v2.0.3
  license: MIT
- name: github.com/cyphar/filepath-securejoin
  version: v0.2.4
  license: BSD-3-Clause
- name: github.com/davecgh/go-spew
  version: v1.1.2-0.20180830191138-d8f796af33cc
  license: 0BSD
- name: github.com/docker/distribution
  version: v2.8.2+incompatible
  license: Apache-2.0
- name: github.com/docker/docker
  version: v24.0.7+incompatible
  license: Apache-2.0
- name: github.com/docker/go-connections
  version: v0.4.0
  license: Apache-2.0
- name: github.com/docker/go-units
  version: v0.5.0
  license: Apache-2.0
- name: github.com/emicklei/go-restful/v3
  version: v3.11.2
  license: MIT
- name: github.com/envoyproxy/go-control-plane
  version: v0.12.0
  license: Apache-2.0
- name: github.com/envoyproxy/protoc-gen-validate
  version: v1.0.4
  license: Apache-2.0
- name: github.com/evanphx/json-patch/v5
  version: v5.9.0
  license: BSD-3-Clause
- name: github.com/exaring/otelpgx
  version: v0.5.3
  license: Apache-2.0
- name: github.com/fatih/color
  version: v1.16.0
  license: MIT
- name: github.com/fsnotify/fsnotify
  version: v1.7.0
  license: BSD-3-Clause
- name: github.com/ghodss/yaml
  version: v1.0.0
  license: null
- name: github.com/go-errors/errors
  version: v1.4.2
  license: MIT
- name: github.com/go-logr/logr
  version: v1.4.1
  license: Apache-2.0
- name: github.com/go-logr/stdr
  version: v1.2.2
  license: Apache-2.0
- name: github.com/go-logr/zapr
  version: v1.3.0
  license: Apache-2.0
- name: github.com/go-ole/go-ole
  version: v1.2.6
  license: MIT
- name: github.com/go-openapi/jsonpointer
  version: v0.20.0
  license: Apache-2.0
- name: github.com/go-openapi/jsonreference
  version: v0.20.2
  license: Apache-2.0
- name: github.com/go-openapi/swag
  version: v0.22.4
  license: Apache-2.0
- name: github.com/go-sql-driver/mysql
  version: v1.6.0
  license: MPL-2.0
- name: github.com/go-task/slim-sprig
  version: v0.0.0-20230315185526-52ccab3ef572
  license: MIT
- name: github.com/gobuffalo/flect
  version: v1.0.2
  license: MIT
- name: github.com/goburrow/cache
  version: v0.1.4
  license: BSD-3-Clause
- name: github.com/gobwas/glob
  version: v0.2.3
  license: MIT
- name: github.com/gogo/protobuf
  version: v1.3.2
  license: 0BSD
- name: github.com/golang-jwt/jwt/v4
  version: v4.5.0
  license: MIT
- name: github.com/golang-migrate/migrate/v4
  version: v4.17.0
  license: MIT
- name: github.com/golang/groupcache
  version: v0.0.0-20210331224755-41bb18bfe9da
  license: Apache-2.0
- name: github.com/golang/protobuf
  version: v1.5.3
  license: BSD-3-Clause
- name: github.com/gonvenience/bunt
  version: v1.3.5
  license: MIT
- name: github.com/gonvenience/neat
  version: v1.3.12
  license: MIT
- name: github.com/gonvenience/term
  version: v1.0.2
  license: MIT
- name: github.com/gonvenience/text
  version: v1.0.7
  license: MIT
- name: github.com/gonvenience/wrap
  version: v1.1.2
  license: MIT
- name: github.com/gonvenience/ytbx
  version: v1.4.4
  license: MIT
- name: github.com/google/gnostic-models
  version: v0.6.8
  license: Apache-2.0
- name: github.com/google/go-cmp
  version: v0.6.0
  license: BSD-3-Clause
- name: github.com/google/gofuzz
  version: v1.2.0
  license: Apache-2.0
- name: github.com/google/pprof
  version: v0.0.0-20230602150820-91b7bce49751
  license: Apache-2.0
- name: github.com/google/uuid
  version: v1.6.0
  license: BSD-3-Clause
- name: github.com/gorilla/websocket
  version: v1.5.0
  license: BSD-2-Clause
- name: github.com/grpc-ecosystem/go-grpc-prometheus
  version: v1.2.0
  license: Apache-2.0
- name: github.com/grpc-ecosystem/grpc-gateway/v2
  version: v2.19.0
  license: BSD-3-Clause
- name: github.com/gruntwork-io/go-commons
  version: v0.8.0
  license: MIT
- name: github.com/gruntwork-io/terratest
  version: v0.46.11
  license: Apache-2.0
- name: github.com/hashicorp/errwrap
  version: v1.1.0
  license: MPL-2.0
- name: github.com/hashicorp/go-multierror
  version: v1.1.1
  license: MPL-2.0
- name: github.com/hoisie/mustache
  version: v0.0.0-20160804235033-6375acf62c69
  license: MIT
- name: github.com/homeport/dyff
  version: v1.6.0
  license: MIT
- name: github.com/huandu/xstrings
  version: v1.4.0
  license: MIT
- name: github.com/imdario/mergo
  version: v0.3.16
  license: BSD-3-Clause
- name: github.com/inconshreveable/mousetrap
  version: v1.1.0
  license: Apache-2.0
- name: github.com/jackc/pgpassfile
  version: v1.0.0
  license: MIT
- name: github.com/jackc/pgservicefile
  version: v0.0.0-20221227161230-091c0ba34f0a
  license: MIT
- name: github.com/jackc/pgx/v5
  version: v5.5.2
  license: MIT
- name: github.com/jackc/puddle/v2
  version: v2.2.1
  license: MIT
- name: github.com/jmespath/go-jmespath
  version: v0.4.0
  license: Apache-2.0
- name: github.com/josharian/intern
  version: v1.0.0
  license: MIT
- name: github.com/json-iterator/go
  version: v1.1.12
  license: MIT
- name: github.com/kelseyhightower/envconfig
  version: v1.4.0
  license: MIT
- name: github.com/kennygrant/sanitize
  version: v1.2.4
  license: BSD-3-Clause
- name: github.com/klauspost/compress
  version: v1.17.4
  license: Apache-2.0
- name: github.com/kumahq/protoc-gen-kumadoc
  version: v0.3.1
  license: Apache-2.0
- name: github.com/lib/pq
  version: v1.10.9
  license: MIT
- name: github.com/lucasb-eyer/go-colorful
  version: v1.2.0
  license: MIT
- name: github.com/lufia/plan9stats
  version: v0.0.0-20211012122336-39d0f177ccd0
  license: BSD-3-Clause
- name: github.com/magiconair/properties
  version: v1.8.7
  license: BSD-2-Clause
- name: github.com/mailru/easyjson
  version: v0.7.7
  license: MIT
- name: github.com/mattn/go-ciede2000
  version: v0.0.0-20170301095244-782e8c62fec3
  license: MIT
- name: github.com/mattn/go-colorable
  version: v0.1.13
  license: MIT
- name: github.com/mattn/go-isatty
  version: v0.0.20
  license: MIT
- name: github.com/mattn/go-zglob
  version: v0.0.2-0.20190814121620-e3c945676326
  license: MIT
- name: github.com/miekg/dns
  version: v1.1.58
  license: BSD-3-Clause
- name: github.com/mitchellh/copystructure
  version: v1.2.0
  license: MIT
- name: github.com/mitchellh/go-homedir
  version: v1.1.0
  license: MIT
- name: github.com/mitchellh/go-ps
  version: v1.0.0
  license: MIT
- name: github.com/mitchellh/hashstructure
  version: v1.1.0
  license: MIT
- name: github.com/mitchellh/reflectwalk
  version: v1.0.2
  license: MIT
- name: github.com/moby/patternmatcher
  version: v0.6.0
  license: Apache-2.0
- name: github.com/moby/spdystream
  version: v0.2.0
  license: Apache-2.0
- name: github.com/moby/sys/mountinfo
  version: v0.7.1
  license: Apache-2.0
- name: github.com/moby/sys/sequential
  version: v0.5.0
  license: Apache-2.0
- name: github.com/moby/term
  version: v0.5.0
  license: Apache-2.0
- name: github.com/modern-go/concurrent
  version: v0.0.0-20180306012644-bacd9c7ef1dd
  license: Apache-2.0
- name: github.com/modern-go/reflect2
  version: v1.0.2
  license: Apache-2.0
- name: github.com/morikuni/aec
  version: v1.0.0
  license: MIT
- name: github.com/munnerz/goautoneg
  version: v0.0.0-20191010083416-a7dc8b61c822
  license: BSD-3-Clause
- name: github.com/mxk/go-flowrate
  version: v0.0.0-20140419014527-cca7078d478f
  license: BSD-3-Clause
- name: github.com/natefinch/atomic
  version: v1.0.1
  license: MIT
- name: github.com/onsi/ginkgo/v2
  version: v2.15.0
  license: MIT
- name: github.com/onsi/gomega
  version: v1.31.1
  license: MIT
- name: github.com/opencontainers/go-digest
  version: v1.0.0
  license: CC-BY-SA-4.0
- name: github.com/opencontainers/image-spec
  version: v1.1.0-rc5
  license: Apache-2.0
- name: github.com/opencontainers/runc
  version: v1.1.10
  license: Apache-2.0
- name: github.com/patrickmn/go-cache
  version: v2.1.0+incompatible
  license: MIT
- name: github.com/pkg/errors
  version: v0.9.1
  license: BSD-2-Clause
- name: github.com/pmezard/go-difflib
  version: v1.0.1-0.20181226105442-5d4384ee4fb2
  license: BSD-3-Clause
- name: github.com/power-devops/perfstat
  version: v0.0.0-20210106213030-5aafc221ea8c
  license: MIT
- name: github.com/pquerna/otp
  version: v1.2.0
  license: Apache-2.0
- name: github.com/prometheus/client_golang
  version: v1.18.0
  license: Apache-2.0
- name: github.com/prometheus/client_model
  version: v0.5.0
  license: Apache-2.0
- name: github.com/prometheus/common
  version: v0.46.0
  license: Apache-2.0
- name: github.com/prometheus/procfs
  version: v0.12.0
  license: Apache-2.0
- name: github.com/russross/blackfriday/v2
  version: v2.1.0
  license: BSD-2-Clause
- name: github.com/sergi/go-diff
  version: v1.3.1
  license: MIT
- name: github.com/sethvargo/go-retry
  version: v0.2.4
  license: Apache-2.0
- name: github.com/shirou/gopsutil/v3
  version: v3.23.11
  license: BSD-3-Clause
- name: github.com/shoenig/go-m1cpu
  version: v0.1.6
  license: MPL-2.0
- name: github.com/shopspring/decimal
  version: v1.3.1
  license: MIT
- name: github.com/sirupsen/logrus
  version: v1.9.3
  license: MIT
- name: github.com/slok/go-http-metrics
  version: v0.11.0
  license: Apache-2.0
- name: github.com/soheilhy/cmux
  version: v0.1.5
  license: Apache-2.0
- name: github.com/spf13/cast
  version: v1.5.1
  license: MIT
- name: github.com/spf13/cobra
  version: v1.8.0
  license: Apache-2.0
- name: github.com/spf13/pflag
  version: v1.0.5
  license: BSD-3-Clause
- name: github.com/spiffe/go-spiffe/v2
  version: v2.1.7
  license: Apache-2.0
- name: github.com/stretchr/testify
  version: v1.8.4
  license: MIT
- name: github.com/testcontainers/testcontainers-go
  version: v0.27.0
  license: MIT
- name: github.com/texttheater/golang-levenshtein
  version: v1.0.1
  license: MIT
- name: github.com/tklauser/go-sysconf
  version: v0.3.12
  license: BSD-3-Clause
- name: github.com/tklauser/numcpus
  version: v0.6.1
  license: Apache-2.0
- name: github.com/tonglil/opentelemetry-go-datadog-propagator
  version: v0.1.1
  license: Apache-2.0
- name: github.com/urfave/cli
  version: v1.22.12
  license: MIT
- name: github.com/virtuald/go-ordered-json
  version: v0.0.0-20170621173500-b18e6e673d74
  license: BSD-3-Clause
- name: github.com/vishvananda/netlink
  version: v1.2.1-beta.2
  license: Apache-2.0
- name: github.com/vishvananda/netns
  version: v0.0.4
  license: Apache-2.0
- name: github.com/xeipuuv/gojsonpointer
  version: v0.0.0-20190905194746-02993c407bfb
  license: Apache-2.0
- name: github.com/xeipuuv/gojsonreference
  version: v0.0.0-20180127040603-bd5ef7bd5415
  license: Apache-2.0
- name: github.com/xeipuuv/gojsonschema
  version: v1.2.0
  license: Apache-2.0
- name: github.com/yusufpapurcu/wmi
  version: v1.2.3
  license: MIT
- name: go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful
  version: v0.47.0
  license: Apache-2.0
- name: go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
  version: v0.47.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel
  version: v1.22.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc
  version: v0.45.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel/exporters/otlp/otlptrace
  version: v1.22.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
  version: v1.22.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel/metric
  version: v1.22.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel/sdk
  version: v1.22.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel/sdk/metric
  version: v1.22.0
  license: Apache-2.0
- name: go.opentelemetry.io/otel/trace
  version: v1.22.0
  license: Apache-2.0
- name: go.opentelemetry.io/proto/otlp
  version: v1.1.0
  license: Apache-2.0
- name: go.uber.org/atomic
  version: v1.10.0
  license: MIT
- name: go.uber.org/multierr
  version: v1.11.0
  license: MIT
- name: go.uber.org/zap
  version: v1.26.0
  license: MIT
- name: golang.org/x/crypto
  version: v0.18.0
  license: BSD-3-Clause
- name: golang.org/x/exp
  version: v0.0.0-20231006140011-7918f672742d
  license: BSD-3-Clause
- name: golang.org/x/mod
  version: v0.14.0
  license: BSD-3-Clause
- name: golang.org/x/net
  version: v0.20.0
  license: BSD-3-Clause
- name: golang.org/x/oauth2
  version: v0.16.0
  license: BSD-3-Clause
- name: golang.org/x/sync
  version: v0.6.0
  license: BSD-3-Clause
- name: golang.org/x/sys
  version: v0.16.0
  license: BSD-3-Clause
- name: golang.org/x/term
  version: v0.16.0
  license: BSD-3-Clause
- name: golang.org/x/text
  version: v0.14.0
  license: BSD-3-Clause
- name: golang.org/x/time
  version: v0.3.0
  license: BSD-3-Clause
- name: golang.org/x/tools
  version: v0.17.0
  license: BSD-3-Clause
- name: gomodules.xyz/jsonpatch/v2
  version: v2.4.0
  license: Apache-2.0
- name: gonum.org/v1/gonum
  version: v0.14.0
  license: BSD-3-Clause
- name: google.golang.org/appengine
  version: v1.6.8
  license: Apache-2.0
- name: google.golang.org/genproto
  version: v0.0.0-20231212172506-995d672761c0
  license: Apache-2.0
- name: google.golang.org/genproto/googleapis/api
  version: v0.0.0-20240102182953-50ed04b92917
  license: Apache-2.0
- name: google.golang.org/genproto/googleapis/rpc
  version: v0.0.0-20240102182953-50ed04b92917
  license: Apache-2.0
- name: google.golang.org/grpc
  version: v1.61.0
  license: Apache-2.0
- name: google.golang.org/protobuf
  version: v1.32.0
  license: BSD-3-Clause
- name: gopkg.in/inf.v0
  version: v0.9.1
  license: BSD-3-Clause
- name: gopkg.in/natefinch/lumberjack.v2
  version: v2.2.1
  license: MIT
- name: gopkg.in/yaml.v3
  version: v3.0.1
  license: Apache-2.0
- name: gotest.tools/v3
  version: v3.5.0
  license: Apache-2.0
- name: helm.sh/helm/v3
  version: v3.14.0
  license: Apache-2.0
- name: k8s.io/api
  version: v0.29.1
  license: Apache-2.0
- name: k8s.io/apiextensions-apiserver
  version: v0.29.1
  license: Apache-2.0
- name: k8s.io/apimachinery
  version: v0.29.1
  license: Apache-2.0
- name: k8s.io/client-go
  version: v0.29.1
  license: Apache-2.0
- name: k8s.io/component-base
  version: v0.29.1
  license: Apache-2.0
- name: k8s.io/klog/v2
  version: v2.120.1
  license: Apache-2.0
- name: k8s.io/kube-openapi
  version: v0.0.0-20231010175941-2dd684a91f00
  license: Apache-2.0
- name: k8s.io/kubectl
  version: v0.29.1
  license: Apache-2.0
- name: k8s.io/utils
  version: v0.0.0-20230726121419-3b25d923346b
  license: Apache-2.0
- name: sigs.k8s.io/controller-runtime
  version: v0.17.0
  license: Apache-2.0
- name: sigs.k8s.io/controller-tools
  version: v0.14.0
  license: Apache-2.0
- name: sigs.k8s.io/gateway-api
  version: v1.0.1-0.20231106150010-c39f0c1568c1
  license: Apache-2.0
- name: sigs.k8s.io/json
  version: v0.0.0-20221116044647-bc3834ca7abd
  license: Apache-2.0
- name: sigs.k8s.io/structured-merge-diff/v4
  version: v4.4.1
  license: Apache-2.0
- name: sigs.k8s.io/yaml
  version: v1.4.0
  license: MIT
```
