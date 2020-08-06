# Open Service Mesh CNCF Sandbox Proposal

## Name of project

Open Service Mesh

## Project description

Open Service Mesh (OSM) is a lightweight, extensible, cloud native service mesh that allows users to uniformly manage, secure, and get out-of-the-box observability features for highly dynamic microservice environments.

## Overview

Open Service Mesh (OSM) is a lightweight and extensible cloud native service mesh.

OSM takes a simple approach for users to uniformly manage, secure, and get out-of-the box observability features for highly dynamic microservice environments.

Using the CNCF Envoy project, OSM implements Service Mesh Interface (SMI) for securing and managing your microservice applications.

### About Open Service Mesh

OSM runs on Kubernetes. The control plane implements Envoy's xDS and is configured with SMI APIs. OSM injects an Envoy proxy as a sidecar container next to each instance of an application.

The data plane (the set of Envoy proxies running as part OSM) executes rules around access control policies, implements routing configuration, and captures metrics. The control plane continually programs the data plane to ensure policies and routing rules are up to date and ensures the data plane is healthy.

### Features

- Easily and transparently configure traffic shifting for deployments
- Secure end-to-end service to service communication by enabling mTLS
- Define and execute fine grained access control policies for services
- Observability and insights into application metrics for debugging and monitoring services
- Integrate with external certificate management services/solutions with a pluggable interface
- Onboard applications onto the mesh by enabling automatic sidecar injection of Envoy proxy
- Flexible enough to handle both simple and complex scenarios through SMI and Envoy XDS APIs

## Statement on alignment with CNCF charter mission

Open Service Mesh is well-aligned with the CNCF mission to make cloud native computing ubiquitous. OSM enables users to uniformly manage, secure, and observe service-to-service communication in highly dynamic microservice environments. We hope for OSM to be a community-led project that it will spur collaboration on new and existing APIs for Service Mesh Interface (SMI), the service mesh spec which is a CNCF project. We intend for OSM to have open governance and be in a place where we can easily collaborate with the community, which is why we are donating it to the CNCF.

## Comparison with similar projects (inside or outside the CNCF), including what differentiates this project

- **[Linkerd](https://github.com/linkerd/linkerd)** is a service mesh which is an incubating CNCF project and implements SMI APIs but does not use the Envoy data plane.
- **[Kuma](https://github.com/kumahq/kuma)** is a service mesh which is a sandbox CNCF project and uses the Envoy data plane but does not implement SMI APIs.
- **[Istio](https://github.com/istio/istio)** is a non-CNCF project which is a service mesh using the Envoy data plane.

Open Service Mesh implements SMI APIs, uses the Envoy data plane, and plans to be a sandbox CNCF project.

## Sponsor from TOC

Sponsors from the TOC will be added as they are confirmed (as applicable, based on the fact we’re submitting this under the new lightweight sandbox guidelines).

## Preferred maturity level
Sandbox

## License

Currently MIT (https://github.com/openservicemesh/osm/blob/main/LICENSE) - happy to relicense to Apache2 at donation.

## Source control

We use GitHub: https://github.com/openservicemesh/osm

## External dependencies (including licenses)

Currently OSM uses go 1.14 with the following dependencies:

        github.com/AlekSi/gocov-xml v0.0.0-20190121064608-3a14fb1c4737
        github.com/Azure/azure-sdk-for-go v34.0.0+incompatible
        github.com/Azure/go-autorest/autorest v0.10.0
        github.com/Azure/go-autorest/autorest/azure/auth v0.1.0
        github.com/Azure/go-autorest/autorest/to v0.2.0
        github.com/Azure/go-autorest/autorest/validation v0.1.0 // indirect
        github.com/Masterminds/semver/v3 v3.1.0
        github.com/axw/gocov v1.0.0
        github.com/cncf/udpa/go v0.0.0-20200629203442-efcf912fb354 // indirect
        github.com/deckarep/golang-set v1.7.1
        github.com/envoyproxy/go-control-plane v0.9.6
        github.com/golang/protobuf v1.4.2
        github.com/golangci/golangci-lint v1.30.0
        github.com/google/go-cmp v0.5.0
        github.com/google/uuid v1.1.1
        github.com/googleapis/gnostic v0.3.1 // indirect
        github.com/gorilla/mux v1.7.3
        github.com/hashicorp/vault/api v1.0.4
        github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
        github.com/jstemmer/go-junit-report v0.0.0-20190106144839-af01ea7f8024
        github.com/matm/gocov-html v0.0.0-20200509184451-71874e2e203b
        github.com/mitchellh/gox v1.0.1
        github.com/onsi/ginkgo v1.13.0
        github.com/onsi/gomega v1.10.1
        github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
        github.com/pkg/errors v0.9.1
        github.com/prometheus/client_golang v1.0.0
        github.com/rs/zerolog v1.18.0
        github.com/servicemeshinterface/smi-sdk-go v0.4.1
        github.com/spf13/cobra v1.0.0
        github.com/spf13/pflag v1.0.5
        golang.org/x/lint v0.0.0-20200302205851-738671d3881b
        google.golang.org/grpc v1.27.0
        gopkg.in/yaml.v2 v2.3.0
        helm.sh/helm/v3 v3.2.0
        k8s.io/api v0.18.0
        k8s.io/apimachinery v0.18.0
        k8s.io/cli-runtime v0.18.0
        k8s.io/client-go v0.18.0
        k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6 // indirect
        rsc.io/letsencrypt v0.0.3 // indirect

## Initial committers (how long working on project, companies they represent)

Currently 25 committers over the last 8 months. While the committers thus far are from Microsoft, our intent is community ownership, governance, and contribution.

## Infrastructure requests (CI / CNCF Cluster)

We have no infrastructure requests at this time.

## Communication channels (slack, irc, mailing lists)

So far we have a google group mailing list (https://groups.google.com/g/openservicemesh)  and GitHub. We have requested a channel on Kubernetes Slack: https://github.com/kubernetes/community/issues/5009

## Issue tracker

https://github.com/openservicemesh/osm/issues

## Website

https://openservicemesh.io/

## Release methodology and mechanics

Checklist in progress https://github.com/openservicemesh/osm/issues/1406

Here’s what we did for the initial release:

- Enter desired release number on the main branch, like v.0.1.0
- Create a release branch called release-v0.1.0
- Send PRs/commits to that branch which updates the chart AppVersion and default image tag version
- Cut a release by checking out the release branch and pushing a tag v0.1.0 to origin which will kick off the GitHub Action to cut all the artifacts

## Social media accounts

Twitter: @openservicemesh

## Community size and any existing sponsorship

We have about 25 committers from Microsoft thus far, but now that we’re open source we plan to build a broad-based open source community. As for sponsorship, Microsoft currently pays for the project’s GitHub Actions.

## Who is currently known to be using the project? Are they using it in production and at what scale? (It may be hard to obtain accurate data for this, but any supporting evidence of usage is helpful)

As this is a brand new open source project, all usage is yet to come.

## Project logo in svg format

All logos are available in https://github.com/openservicemesh/osm/tree/main/community/logos

The svg main logo is here: https://github.com/openservicemesh/osm/blob/main/community/logos/osm-logo.svg
