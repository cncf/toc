# Knative project proposal

## Name of the project

Knative

## Description

Knative is a set of serverless API that provides higher level abstractions that enables developers to be more productive by simplifying and extending the Kubernetes interfaces for common stateless application patterns.

Knative consists of two sets of APIs – Serving and Eventing – and includes a reference implementation and conformance tests for both the Serving and Eventing components. These two components provide a simplified (easier to learn) and enriched (more features, less effort) experience for developers writing HTTP and event-driven applications when compared with baseline Kubernetes objects.

Knative has also been a proving ground for several interesting Kubernetes-related innovations: higher-level multi-provider HTTP abstractions beyond the Ingress API, automatic sharding and leader election of controllers using inversion of control, and polymorphic usage of Kubernetes API resources (aka “duck typing”) as an extension mechanism.

## Features

Knative Serving builds on Kubernetes to manage all aspects of the serverless application and function lifecycle, from deployment to scaling, routing, and provisioning SSL certificates. The primary resources of Serving are the Service, Route, Configuration, and Revision objects, which provide a clean abstraction for defining lifecycle and routing operations on an application. A pluggable networking layer backs these applications and functions with support from projects such as Istio, Ambassador, Contour, Kourier, Gloo, and many others.

Knative Eventing enables developers to configure event-driven architectures based on the CloudEvents HTTP binding. Event-driven architectures enable teams to decouple internal services in both time and interface, and seamlessly integrate with external event providers. Knative has enabled an ecosystem of event sources with common routing infrastructure using Broker and Trigger abstractions over multiple transports, enabling delivery to multiple destinations, including Knative Services, Kubernetes resources, and even VMs.

In addition to Serving and Eventing, all other parts of the Knative project in Github will be transferred to the CNCF – documents, a client implementation, a Kubernetes operator, and any other materials in the child Knative repos.

## Sponsor / Advisor from TOC

Davanum Srinivas (dims)

## License

Knative is licensed under the Apache 2.0 license

## Source control repositories

The specifications, reference implementation, clients, and tests are stored in public Github repositories as children under <https://github.com/knative> and all of the Knative Sandbox projects under <https://github.com/knative-sandbox/>.

## External dependencies

From knative/serving go.mod:
* Apache-2.0 cloud.google.com/go
* Apache-2.0 contrib.go.opencensus.io/exporter/ocagent
* Apache-2.0 contrib.go.opencensus.io/exporter/prometheus
* Apache-2.0 contrib.go.opencensus.io/exporter/zipkin
* Apache-2.0 github.com/Azure/azure-sdk-for-go
* Apache-2.0 github.com/Azure/go-autorest/autorest
* Apache-2.0 github.com/Azure/go-autorest/autorest/adal
* Apache-2.0 github.com/Azure/go-autorest/autorest/date
* Apache-2.0 github.com/Azure/go-autorest/autorest/to
* Apache-2.0 github.com/Azure/go-autorest/autorest/validation
* Apache-2.0 github.com/Azure/go-autorest/logger
* Apache-2.0 github.com/Azure/go-autorest/tracing
* Apache-2.0 github.com/aws/aws-sdk-go
* Apache-2.0 github.com/census-instrumentation/opencensus-proto
* Apache-2.0 github.com/docker/cli
* Apache-2.0 github.com/docker/distribution
* Apache-2.0 github.com/docker/docker
* Apache-2.0 github.com/go-logr/logr
* Apache-2.0 github.com/golang/glog
* Apache-2.0 github.com/golang/groupcache
* Apache-2.0 github.com/google/go-containerregistry
* Apache-2.0 github.com/google/go-containerregistry/pkg/authn/k8schain
* Apache-2.0 github.com/google/gofuzz
* Apache-2.0 github.com/google/mako
* Apache-2.0 github.com/googleapis/gnostic
* Apache-2.0 github.com/influxdata/tdigest
* Apache-2.0 github.com/matttproud/golang_protobuf_extensions
* Apache-2.0 github.com/modern-go/concurrent
* Apache-2.0 github.com/modern-go/reflect2
* Apache-2.0 github.com/opencontainers/go-digest
* Apache-2.0 github.com/openzipkin/zipkin-go
* Apache-2.0 github.com/prometheus/client_golang
* Apache-2.0 github.com/prometheus/common
* Apache-2.0 github.com/prometheus/procfs
* Apache-2.0 github.com/prometheus/statsd_exporter
* Apache-2.0 github.com/vdemeester/k8s-pkg-credentialprovider
* Apache-2.0 go.opencensus.io
* Apache-2.0 gomodules.xyz/jsonpatch/v2
* Apache-2.0 google.golang.org/genproto
* Apache-2.0 gopkg.in/yaml.v2
* Apache-2.0 k8s.io/api
* Apache-2.0 k8s.io/apiextensions-apiserver
* Apache-2.0 k8s.io/apimachinery
* Apache-2.0 k8s.io/cloud-provider
* Apache-2.0 k8s.io/component-base
* Apache-2.0 k8s.io/klog
* Apache-2.0 k8s.io/klog/v2
* Apache-2.0 k8s.io/kube-openapi
* Apache-2.0 k8s.io/legacy-cloud-providers
* Apache-2.0 knative.dev/caching
* Apache-2.0 knative.dev/hack
* Apache-2.0 knative.dev/networking
* Apache-2.0 knative.dev/pkg
* Apache-2.0 knative.dev/serving
* BSD-2-Clause github.com/gorilla/websocket
* BSD-3-Clause github.com/evanphx/json-patch
* BSD-3-Clause github.com/evanphx/json-patch/v5
* BSD-3-Clause github.com/gogo/protobuf
* BSD-3-Clause github.com/google/go-cmp
* BSD-3-Clause github.com/google/go-github/v27
* BSD-3-Clause github.com/google/go-querystring
* BSD-3-Clause github.com/google/uuid
* BSD-3-Clause github.com/grpc-ecosystem/grpc-gateway
* BSD-3-Clause golang.org/x/crypto
* BSD-3-Clause golang.org/x/oauth2
* BSD-3-Clause golang.org/x/sync
* BSD-3-Clause golang.org/x/sys
* BSD-3-Clause golang.org/x/text
* BSD-3-Clause golang.org/x/time
* BSD-3-Clause google.golang.org/api
* BSD-3-Clause google.golang.org/protobuf
* BSD-3-Clause gopkg.in/inf.v0
* ISC github.com/davecgh/go-spew
* MIT github.com/beorn7/perks
* MIT github.com/blang/semver/v4
* MIT github.com/cespare/xxhash/v2
* MIT github.com/docker/docker-credential-helpers
* MIT github.com/form3tech-oss/jwt-go
* MIT github.com/go-kit/log
* MIT github.com/go-logfmt/logfmt
* MIT github.com/gobuffalo/flect
* MIT github.com/josharian/intern
* MIT github.com/json-iterator/go
* MIT github.com/kelseyhightower/envconfig
* MIT github.com/mailru/easyjson
* MIT github.com/sirupsen/logrus
* MIT github.com/tsenart/vegeta/v12
* MIT go.uber.org/atomic
* MIT go.uber.org/automaxprocs
* MIT go.uber.org/multierr
* MIT go.uber.org/zap
* MIT gopkg.in/yaml.v3
* MIT sigs.k8s.io/yaml
* MPL-2.0-no-copyleft-exception github.com/hashicorp/golang-lru

From knative/eventing go.mod:
* Apache-2.0 cloud.google.com/go
* Apache-2.0 cloud.google.com/go/storage
* Apache-2.0 contrib.go.opencensus.io/exporter/ocagent
* Apache-2.0 contrib.go.opencensus.io/exporter/prometheus
* Apache-2.0 contrib.go.opencensus.io/exporter/zipkin
* Apache-2.0 github.com/census-instrumentation/opencensus-proto
* Apache-2.0 github.com/cloudevents/conformance
* Apache-2.0 github.com/cloudevents/sdk-go/observability/opencensus/v2
* Apache-2.0 github.com/cloudevents/sdk-go/v2
* Apache-2.0 github.com/go-logr/logr
* Apache-2.0 github.com/golang/glog
* Apache-2.0 github.com/golang/groupcache
* Apache-2.0 github.com/google/gofuzz
* Apache-2.0 github.com/google/mako
* Apache-2.0 github.com/googleapis/gnostic
* Apache-2.0 github.com/influxdata/tdigest
* Apache-2.0 github.com/matttproud/golang_protobuf_extensions
* Apache-2.0 github.com/modern-go/concurrent
* Apache-2.0 github.com/modern-go/reflect2
* Apache-2.0 github.com/openzipkin/zipkin-go
* Apache-2.0 github.com/prometheus/client_golang
* Apache-2.0 github.com/prometheus/client_model
* Apache-2.0 github.com/prometheus/common
* Apache-2.0 github.com/prometheus/procfs
* Apache-2.0 github.com/prometheus/statsd_exporter
* Apache-2.0 github.com/spf13/cobra
* Apache-2.0 github.com/wavesoftware/go-ensure
* Apache-2.0 go.opencensus.io
* Apache-2.0 go.opentelemetry.io/otel
* Apache-2.0 gomodules.xyz/jsonpatch/v2
* Apache-2.0 google.golang.org/genproto
* Apache-2.0 google.golang.org/grpc
* Apache-2.0 gopkg.in/yaml.v2
* Apache-2.0 k8s.io/api
* Apache-2.0 k8s.io/apiextensions-apiserver
* Apache-2.0 k8s.io/apimachinery
* Apache-2.0 k8s.io/apiserver
* Apache-2.0 k8s.io/client-go
* Apache-2.0 k8s.io/klog
* Apache-2.0 k8s.io/klog/v2
* Apache-2.0 k8s.io/kube-openapi
* Apache-2.0 k8s.io/utils
* Apache-2.0 knative.dev/eventing
* Apache-2.0 knative.dev/hack
* Apache-2.0 knative.dev/hack/schema
* Apache-2.0 knative.dev/pkg
* Apache-2.0 knative.dev/reconciler-test
* Apache-2.0 sigs.k8s.io/structured-merge-diff/v4
* BSD-2-Clause github.com/gorilla/websocket
* BSD-2-Clause github.com/pkg/errors
* BSD-3-Clause github.com/evanphx/json-patch
* BSD-3-Clause github.com/evanphx/json-patch/v5
* BSD-3-Clause github.com/fsnotify/fsnotify
* BSD-3-Clause github.com/gogo/protobuf
* BSD-3-Clause github.com/golang/protobuf
* BSD-3-Clause github.com/google/go-cmp
* BSD-3-Clause github.com/google/go-github/v27
* BSD-3-Clause github.com/google/go-querystring
* BSD-3-Clause github.com/google/uuid
* BSD-3-Clause github.com/googleapis/gax-go/v2
* BSD-3-Clause github.com/grpc-ecosystem/grpc-gateway
* BSD-3-Clause github.com/imdario/mergo
* BSD-3-Clause github.com/pmezard/go-difflib
* BSD-3-Clause github.com/rickb777/date
* BSD-3-Clause github.com/rickb777/plural
* BSD-3-Clause github.com/rogpeppe/fastuuid
* BSD-3-Clause github.com/spf13/pflag
* BSD-3-Clause golang.org/x/crypto
* BSD-3-Clause golang.org/x/net
* BSD-3-Clause golang.org/x/oauth2
* BSD-3-Clause golang.org/x/sync
* BSD-3-Clause golang.org/x/sys
* BSD-3-Clause golang.org/x/term
* BSD-3-Clause golang.org/x/text
* BSD-3-Clause golang.org/x/time
* BSD-3-Clause golang.org/x/xerrors
* BSD-3-Clause google.golang.org/api
* BSD-3-Clause google.golang.org/protobuf
* BSD-3-Clause gopkg.in/inf.v0
* BSD-3-Clause gopkg.in/tomb.v1
* ISC github.com/blendle/zapdriver
* ISC github.com/davecgh/go-spew
* MIT github.com/beorn7/perks
* MIT github.com/blang/semver/v4
* MIT github.com/go-kit/log
* MIT github.com/go-logfmt/logfmt
* MIT github.com/gobuffalo/flect
* MIT github.com/josharian/intern
* MIT github.com/json-iterator/go
* MIT github.com/kelseyhightower/envconfig
* MIT github.com/mailru/easyjson
* MIT github.com/mitchellh/go-homedir
* MIT github.com/nxadm/tail
* MIT github.com/onsi/ginkgo
* MIT github.com/pelletier/go-toml/v2
* MIT github.com/robfig/cron/v3
* MIT github.com/stretchr/testify
* MIT github.com/tsenart/vegeta/v12
* MIT github.com/valyala/bytebufferpool
* MIT go.uber.org/atomic
* MIT go.uber.org/automaxprocs
* MIT go.uber.org/multierr
* MIT go.uber.org/zap
* MIT gopkg.in/yaml.v3
* MIT sigs.k8s.io/yaml
* MPL-2.0 github.com/hashicorp/golang-lru
* MPL-2.0-no-copyleft-exception github.com/hashicorp/go-cleanhttp
* MPL-2.0-no-copyleft-exception github.com/hashicorp/go-retryablehttp

## Issue tracker

Serving: <https://github.com/knative/serving/issues>

Eventing: <https://github.com/knative/eventing/issues>

See the github issue lists under each Knative Sandbox project: <https://github.com/knative-sandbox/>

## Initial committers

Reference <https://github.com/knative/community/tree/main/OWNERS_ALIASES> and 
<https://github.com/knative/community/tree/main/peribolos/knative.yaml>

## Infrastructure requirements (CI / CNCF Cluster)

Currently using Prow infrastructure donated by Google and GitHub Actions; plan to migrate to community-owned Prow instances and (where possible) prefer GitHub Actions for automation.

In addition to the above infrastructure, nightly and release images (and test logs) are stored in gcr.io and Google Cloud Storage

## Statement on alignment with CNCF mission

Knative's mission is to build a cloud-native layer for serverless workloads on top of Kubernetes. This mission is closely aligned with the CNCF's mission to make cloud-native computing ubiquitous. We also believe that fostering a vendor-neutral project with open governance will lead to the most innovation and that we can make serverless workloads accessible to all. CNCF’s definition of cloud native technologies (“Cloud native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds”) is exactly the area that Knative targets by making it easy to build scalable applications in virtually any environment.

## Communication channels

Mailing lists:
* <knative-dev@googlegroups.com>
* <knative-users@googlegroups.com> 

Knative.team GSuite domain and shared folder

<https://slack.knative.dev/>

## Website

<https://knative.dev/>

## Release methodology and mechanics

Release is automated via Prow jobs which trigger on the presence of release branches. Configuration is here:
* <https://github.com/knative/test-infra/blob/main/config/prod/prow/jobs/config.yaml>
* <https://github.com/knative/hack/tree/main/release.sh> 

Documentation is here: <https://github.com/knative/release> (this is a project-wide release train)

The current release versioning mechanism is described in [this document](https://docs.google.com/document/d/1_khEw5_l8Zhh4FuY3NM56o_OM-cNWmmq1vXfLCxZqMA/edit). The TL;DR:

Releases branches are created as release-1.x, and tagged as knative-1.x.y for release artifacts and v0.27+x.y for go module imports to avoid semantic versioning problems.

## Specifications

* Specifications: <https://github.com/knative/specs>
* Implementation: <https://github.com/knative/serving> and <https://github.com/knative/eventing> 

## Security processes

See [Vulnerability Disclosure Response Policy](https://github.com/knative/community/blob/main/working-groups/security/responding.md)

## Roadmap

Each Working Group maintains its own roadmap; the Knative TOC reviews working group roadmaps and status on a quarterly basis (rotating, week-by-week)
* [API Core/Serving](https://github.com/orgs/knative/projects/35)
* [Client](https://github.com/orgs/knative/projects/37)
* [Documentation](https://github.com/orgs/knative/projects/27) + [User Experience](https://github.com/orgs/knative/projects/20)
* [Eventing](https://github.com/orgs/knative/projects/21)
* [Eventing Kafka](https://github.com/orgs/knative-sandbox/projects/8)
* [Eventing Sources](https://github.com/orgs/knative/projects/34)
* [Networking](https://github.com/orgs/knative/projects/32)
* [Productivity](https://github.com/knative/test-infra/projects/3)
* [Scaling](https://github.com/orgs/knative/projects/36)
* [Security](https://github.com/orgs/knative/projects/22)

## Social media account

Twitter: <https://twitter.com/KnativeProject>

YouTube: <https://www.youtube.com/c/KnativeProject>

## Community size

As of Nov 16, 2021 Knative has over [1893 contributors](https://knative.teststats.cncf.io/d/18/overall-project-statistics-table?orgId=1) to the project’s repository. 

Vendors like Google, IBM, Red Hat, SAP, TriggerMesh, and VMware, plus staff from more than [40 additional companies](https://knative.teststats.cncf.io/d/4/company-statistics-by-repository-group?orgId=1) are actively engaged with the project and its community.

Adopters include Blue Origin, DeepC, Liquidx, Optum, Puppet and many other companies. See [ADOPTERS.MD](https://github.com/knative/community/blob/main/ADOPTERS.MD) for a list of over 20 publicly-announced adopters and read Knative [Case Studies](https://knative.dev/docs/about/case-studies/deepc/) on its website.

Knative project has over [7,500 GitHub stars](https://knative.teststats.cncf.io/d/18/overall-project-statistics-table?orgId=1) across its repositories.

The project also hosts monthly community meetups with 20-30 of attendees.

