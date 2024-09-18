## Project name and description

"Weave Flux" (or just "Flux") is the GitOps Kubernetes Operator: it
automatically ensures that the state of a cluster matches the config
in git and images are upgraded automatically. Flux uses an operator in
the cluster to trigger deployments inside Kubernetes. It supports Helm
as well, with a subproject called the "Helm operator".

Weaveworks started the project in July 2016 and is receiving more and
more contributions from a growing community. The team of maintainers
is very active and made up from people from different companies.

## Alignment with CNCF charter mission

TODO

## Sponsor from TOC

Michelle Noorali, Xiang Li.

## Preferred maturity level

Sandbox.

## License

[Apache 2.0 license](https://github.com/weaveworks/flux/blob/master/LICENSE).

## Source control

https://github.com/weaveworks/flux

## External Dependencies

Generated license analysis:

| Dependency                                                                                    | License analysis |
|-----------------------------------------------------------------------------------------------|------------------|
| github.com/weaveworks/flux                                                                    | Apache License 2.0
| github.com/weaveworks/flux/vendor/cloud.google.com/go/compute/metadata                        | Apache License 2.0
| github.com/weaveworks/flux/vendor/contrib.go.opencensus.io/exporter/ocagent                   | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/2opremio/go-k8s-portforward                      | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/Azure/go-autorest                                | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/BurntSushi/toml                                  | MIT License
| github.com/weaveworks/flux/vendor/github.com/Masterminds/goutils                              | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/Masterminds/semver                               | MIT License (98%)
| github.com/weaveworks/flux/vendor/github.com/Masterminds/sprig                                | MIT License (98%)
| github.com/weaveworks/flux/vendor/github.com/aws/aws-sdk-go                                   | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/beorn7/perks/quantile                            | MIT License (98%)
| github.com/weaveworks/flux/vendor/github.com/bradfitz/gomemcache/memcache                     | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/census-instrumentation/opencensus-proto/gen-go   | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/cyphar/filepath-securejoin                       | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/davecgh/go-spew/spew                             | ISC License (98%)
| github.com/weaveworks/flux/vendor/github.com/dgrijalva/jwt-go                                 | MIT License (98%)
| github.com/weaveworks/flux/vendor/github.com/docker/distribution                              | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/docker/go-metrics                                | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/docker/libtrust                                  | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/docker/spdystream/spdy                           | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/ghodss/yaml                                      | ? (BSD 3-clause "New" or "Revised" License, 83%) - actually MIT
| github.com/weaveworks/flux/vendor/github.com/go-kit/kit                                       | MIT License
| github.com/weaveworks/flux/vendor/github.com/go-logfmt/logfmt                                 | MIT License
| github.com/weaveworks/flux/vendor/github.com/gobwas/glob                                      | MIT License
| github.com/weaveworks/flux/vendor/github.com/gogo/googleapis/google/rpc                       | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/gogo/protobuf                                    | BSD 3-clause "New" or "Revised" License (91%)
| github.com/weaveworks/flux/vendor/github.com/gogo/status                                      | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/golang/gddo/httputil/header                      | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/golang/groupcache/lru                            | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/golang/protobuf                                  | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/google/go-cmp/cmp/internal                       | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/google/gofuzz                                    | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/google/uuid                                      | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/googleapis/gnostic                               | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/gophercloud/gophercloud                          | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/gorilla/mux                                      | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/gorilla/websocket                                | BSD 2-clause "Simplified" License (98%)
| github.com/weaveworks/flux/vendor/github.com/grpc-ecosystem/grpc-gateway                      | BSD 3-clause "New" or "Revised" License (97%)
| github.com/weaveworks/flux/vendor/github.com/hashicorp/go-cleanhttp                           | Mozilla Public License 2.0
| github.com/weaveworks/flux/vendor/github.com/hashicorp/golang-lru/simplelru                   | Mozilla Public License 2.0
| github.com/weaveworks/flux/vendor/github.com/huandu/xstrings                                  | MIT License
| github.com/weaveworks/flux/vendor/github.com/imdario/mergo                                    | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/jmespath/go-jmespath                             | ? (The Unlicense, 35%) - actually Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/json-iterator/go                                 | MIT License
| github.com/weaveworks/flux/vendor/github.com/matttproud/golang_protobuf_extensions/pbutil     | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/modern-go/concurrent                             | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/modern-go/reflect2                               | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/ncabatoff/go-seq/seq                             | MIT License
| github.com/weaveworks/flux/vendor/github.com/opencontainers/go-digest                         | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/opencontainers/image-spec/specs-go/v1            | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/opentracing-contrib/go-stdlib/nethttp            | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/opentracing/opentracing-go                       | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/pkg/errors                                       | BSD 2-clause "Simplified" License
| github.com/weaveworks/flux/vendor/github.com/pkg/term/termios                                 | BSD 2-clause "Simplified" License
| github.com/weaveworks/flux/vendor/github.com/prometheus/client_golang/prometheus              | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/prometheus/client_model/go                       | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/prometheus/common                                | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/prometheus/procfs                                | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/ryanuber/go-glob                                 | MIT License
| github.com/weaveworks/flux/vendor/github.com/sirupsen/logrus                                  | MIT License
| github.com/weaveworks/flux/vendor/github.com/spf13/cobra                                      | Apache License 2.0 (95%)
| github.com/weaveworks/flux/vendor/github.com/spf13/pflag                                      | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/github.com/uber/jaeger-client-go                            | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/uber/jaeger-lib/metrics                          | Apache License 2.0
| github.com/weaveworks/flux/vendor/github.com/weaveworks/common                                | Apache License 2.0 (96%)
| github.com/weaveworks/flux/vendor/github.com/weaveworks/go-checkpoint                         | Mozilla Public License 2.0
| github.com/weaveworks/flux/vendor/github.com/weaveworks/promrus                               | Apache License 2.0
| github.com/weaveworks/flux/vendor/go.opencensus.io                                            | Apache License 2.0
| github.com/weaveworks/flux/vendor/golang.org/x/crypto                                         | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/golang.org/x/net                                            | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/golang.org/x/oauth2                                         | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/golang.org/x/sync/semaphore                                 | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/golang.org/x/sys/unix                                       | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/golang.org/x/text                                           | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/golang.org/x/time/rate                                      | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/google.golang.org/api/support/bundler                       | BSD 3-clause "New" or "Revised" License (96%)
| github.com/weaveworks/flux/vendor/google.golang.org/genproto                                  | Apache License 2.0
| github.com/weaveworks/flux/vendor/google.golang.org/grpc                                      | Apache License 2.0
| github.com/weaveworks/flux/vendor/gopkg.in/inf.v0                                             | BSD 3-clause "New" or "Revised" License (97%)
| github.com/weaveworks/flux/vendor/gopkg.in/yaml.v2                                            | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/api                                                  | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/apiextensions-apiserver/pkg                          | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/apimachinery                                         | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/client-go                                            | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/helm/pkg                                             | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/klog                                                 | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/kube-openapi/pkg/util/proto                          | Apache License 2.0
| github.com/weaveworks/flux/vendor/k8s.io/utils                                                | Apache License 2.0
| github.com/weaveworks/flux/vendor/sigs.k8s.io/yaml                                            | ? (BSD 3-clause "New" or "Revised" License, 83%) - actually MIT

## Initial committers

Peter Bourgon (ex-Weaveworks) and Michael Bridgen (Weaveworks) were
the initial committers. Michael is still at the helm as the lead
maintainer today.

## Infrastructure requests

There are no infrastructure requests; Flux drives CI and releases
using CircleCI, which is adequate at present.

## Communication channels

Development discussion is principally via public mailing list and
Slack channel (and GitHub issues/PRs). The [project
README](https://github.com/weaveworks/flux#help) enumerates the user
support channels.

## Issue tracker

In the [GitHub project](https://github.com/weaveworks/flux).

## Website

None at present; documentation is in markdown format in the [GitHub
repo](https://github.com/weaveworks/flux/tree/master/docs).

## Release methodology and mechanics

Container images and executables for releases are built in CI. The
process of preparing a release is [explained
here](https://github.com/weaveworks/flux/blob/master/internal/docs/releasing.md).

## Social media accounts

None.

## Community size and existing sponsorship

Flux has 1967 stars on Github, 530 people on Slack. The project counts
90 individual committers up until now with a lot of drive-by
contributions, and a few quite experienced project contributors who
are working on Flux on an ongoing basis.

The Helm side of things gets a lot of attention.

This [recent blog
post](https://www.weave.works/blog/weave-increased-our-flux-capacity)
about Flux activity in the past year might give a good sense of
activity.

### Project logo

Flux does not have an official logo.
