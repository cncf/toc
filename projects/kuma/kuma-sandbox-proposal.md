Hello,

We would like to propose to donate Kuma to CNCF as a SandBox project. We are opening this GitHub issue according to the process outlined [here](https://github.com/cncf/toc/blob/master/process/project_proposals.md) with the format outlined [here](https://github.com/cncf/toc/issues/344).

Authors:

* Jakub Dyszkiewicz (@jakubdyszkiewicz)
* Marco Palladino (@subnetmarco)

# Background

* [Link to TOC PR](#)
* [Link to Presentation](https://docs.google.com/presentation/d/1zbeFCR2JBuQWH0XIfYmGpoFerv2fg8Wr7ncHSso860A/edit?usp=sharing)
* [Link to GitHub project](https://github.com/Kong/kuma)

## Project Goal

Kuma is a modern control plane for Service Mesh and Microservices. It can run and be operated natively across both Kubernetes and VM environments, making it easy to adopt by every team in the organization with the goal of accelerating both Kubernetes and Service Mesh adoption.

By bundling the Envoy proxy as the underlying data-plane technology, Kuma can instrument any L4/L7 traffic to secure, observe, route and enhance connectivity between any service or database. It can be used natively in Kubernetes via CRDs, while at the same time providing a RESTful API, a native CLI tool and a built-in GUI that can be used to better integrate Kuma with the rest of the organization.

While Kuma provides easy to use Policy abstractions for most use-cases, Kuma also allows for the configuration of the underlying Envoy data-planes in a more fine-grained manner via the `ProxyTemplate` policy. By doing so, Kuma can be used by both first-time users of Service Mesh, as well as the most experienced ones who want greater control of the underlying networking stack.

## Current Status

* **Project releases**: Kuma has made 11 releases so far, with the latest one being 0.5 RC1 on April 17th 2020. The first release of Kuma was 0.1 on September 10th 2019.

* **Community status:** Kuma has introduced open governance guidelines in April 2020, and has bi-weekly community calls to discuss the status of the project and its roadmap. It has about 1,400 stars on GitHub and 18 total contributors so far.

Kuma is currently being adopted by a few enterprise organizations to build zero-trust security model across different lines of business and platforms, as well as to accelerate the adoption of Envoy across the board.

## Future Plans

**Near-Term**:

* Support for L7 HTTP routing policies (in addition to the existing L4 routing policy)
* Support for L7 gRPC routing policies (in addition to the existing L4 routing policy)
* Discussing roadmap for SMI integration (most requested feature on GitHub)
* Introducing pluggable backends for mTLS, logging and tracing policies
* Adding support for multi-cluster Kubernetes deployments
* Support for traffic shadowing and auto-retries
* Support for L7 HTTP traffic permissions (in addition to the existing L4 traffic permission policy)
* Support for L7 gRPC traffic permissions (in addition to the existing L4 traffic permission policy)
* Support for L7 gRPC logging (in addition to the existing L4 traffic log policy)
* GUI wizards for every entity (in addition to Mesh and Dataplane entities)

**Long-Term**:

* OPA integration
* Introducing dynamic loading of Kuma policies at runtime to enhance extensibility
* Service Stickiness support
* Transparent proxying support in Universal mode
* OpenCensus support
* WebAssembly support
* Extending Health-Checking
* Hooks for RBAC controls in RESTful API (+ CLI)
* Internationalization
* Replicable benchmarks to demonstrate performance of Envoy as the sidecar proxy
* ADA support for enterprise adoption
* And more

# Project Scope

## Clear project definition

**_Does the project have a clear and well defined scope_**

The project aims to provide a universal control plane for Service Mesh with native Envoy support to secure, observe, and orchestrate service connectivity in any application - and across every platform - with ease of use, portability and extensibility in mind.

## Value-add to the CNCF ecosystem

**_Does the project have a clear value add to the current project ecosystem. How does it relate to other projects with overlapping capabilities._**

Kuma aligns with the [CNCF mission](https://github.com/cncf/foundation/blob/master/charter.md) as explained in chapter 1, and should advance the adoption of Envoy as the de-facto data plane for any modern architecture. As such Kuma is the only control plane for Service Mesh with native Envoy support that is in the process of being donated to the CNCF.

Envoy is quickly becoming the data-plane of the future for any modern application, and today the CNCF ecosystem doesn't have a donated and vendor-neutral control plane to help organizations increase the adoption of Envoy. As more digital products are built with cloud-native technologies, Kuma wants to advance the adoption of Envoy and other cloud-native projects across every organization and every industry with an open, neutral, portable and easy to use solution, and building an inclusive and diverse community that can support the growth of the project in the CNCF home.

## Alignment with other CNCF projects

**_Does the project align and actively collaborate with other CNCF projects?_**

Kuma is fully aligned in supporting CNCF projects and integrations in order to help every user build a modern cloud-native infrastructure with a modern Service Mesh solution. In specific:

* Kuma can be deployed natively on Kubernetes.
* Kuma supports Envoy as its data plane proxy technology.
* Kuma integrates with Prometheus natively.
* Kuma integrates with Jaeger natively.
* Users of Kuma asked for SMI integration (top requested feature).
* Users of Kuma asked for OPA integration.
* Kuma is planning support for gRPC as a native protocol in its policies.

And more.

**_Does the project require any specific versions of projects (or APIs) to interoperate? (e.g. K8s API, CSI, CNI, CRI)?_**

We leverage the APIs of some projects (like Kubernetes) but we don't require _specific_ versions.

**_Does the project augment or benefit other CNCF projects?_**

Kuma benefits the adoption of Envoy, Prometheus, Jaeger and of course Kubernetes.

## Anticipated use cases

**_What problems does the project solve?_**

* **Security and Compliance**:  Automatically provides encryption and identity to any L4/L7 workload including (but not limited to) APIs, Microservices and Databases. This accelerates the security compliance journey of the organization, reduces risks and enables a Zero-Trust model across the entire organization.

* **Traffic Reliability**: Improves the reliability of all the traffic running across any workload, any cloud and any platform (running on VMs or Kubernetes, both legacy and modern) by providing ready to use Envoy-based network policies.

* **Infrastructure Control**: Provides out-of-the-box policies to manage the service traffic lifecycle and helps enable architectural control over the entire infrastructure. By reducing the overhead, teams can now focus more on delivering business logic and meeting their goals. Eliminating management of the network abstractions also reduces shadow IT practices.

* **Productivity and Speed**: Reduces the network management overhead on the applications and provides a Service Mesh “overlay” across the entire organization.  This allows the team to focus on delivering better and more reliable products faster - while leveraging ready-to-use network connectivity policies powered by Kuma.

* **Distributed Observability**: Provides distributed metrics, logs and network traces across every workload running in the organization. This helps the organization identify connectivity and application errors in a shorter time frame and in more areas, and accelerates the delivery of better uptime.

* **Kubernetes Adoption**: Decouples the networking management from existing VM-based applications to simplify the adoption of Kubernetes and containers, while providing security and observability to the transition reducing risks and accelerating the K8s journey.

## Alignment with SIG Reference Model

**_Does the project align with the SIG CNCF reference model and which capabilities does it require/provide at each level of the reference model._**

## High level architecture

**_Describe the overall architecture of the project. Feel free to add diagrams._**

There are two main components of Kuma that are very important to understand:

* **Control-Plane**: Kuma is first and foremost a control-plane that will accept user input to create and configure Policies like Service Meshes and Traffic Routes, and in order to add services and configure their behavior within the Meshes you have created.
* **Data-Plane**: Kuma also bundles a data-plane implementation based on top of Envoy for convenience, in order to get up and running quickly. An instance of the data-plane will run alongside every instance of our services, and it will process both incoming and outgoing requests for the service.

A minimal Kuma deployment involves one or more instances of the control-plane (kuma-cp), and one or more instances of the data-planes (kuma-dp) which will connect to the control-plane as soon as they startup. Kuma supports two modes:

### Kubernetes mode

When running on Kubernetes, Kuma will store all of its state and configuration on the underlying Kubernetes API Server, therefore requiring no dependency to store the data. Kuma will also automatically inject the sidecar Envoy proxy to the workloads.

![Kubernetes](https://kuma.io/images/docs/0.2.0/diagram-08.jpg "Kuma Kubernetes Architecture")

### Universal mode

When running in Universal mode, Kuma will require a PostgreSQL database to store its state. The PostgreSQL database and schema will have to be initialized accordingly to the installation instructions.

Unlike kubernetes mode, Kuma won't require the kuma-injector executable to run:

![Universal](https://kuma.io/images/docs/0.2.0/diagram-09.jpg "Kuma Universal Architecture")

More information can be accessed from the official documentation: https://kuma.io/docs/latest/documentation/overview/

# Formal Requirements

**_Document that the project fulfills the requirements as documented in the [CNCF graduation criteria for sandbox](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#sandbox-stage)_**

* We [have adopted](https://github.com/Kong/kuma/blob/master/CODE_OF_CONDUCT.md) the CNCF Code of Conduct.
* Will adhere to the CNCF IP policy.

**_Are there any anticipated issues with any of the criteria ?_**

No.

**_Has the TOC been approached for sponsorship? If yes, which TOC members have agreed to sponsor the project?_**

* Matt Klein

# CNCF IP Policy

**_Becoming a sandbox project requires adoption of the CNCF IP Policy: https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy_**

Kuma's source code is licensed under Apache License Version 2.0.

**_Has the IP policy been reviewed?_**

**_List the repos for the project and their current license_**

| github.com/Masterminds/sprig v2.20.0+incompatible                                   | MIT                                                                            |
|-------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|
| github.com/dgrijalva/jwt-go v3.2.0+incompatible                                     | MIT                                                                            |
| github.com/emicklei/go-restful v2.9.6+incompatible                                  | MIT                                                                            |
| github.com/envoyproxy/go-control-plane v0.9.1-0.20191108215040-b0f2cec0e187         | Apache-2.0                                                                     |
| github.com/envoyproxy/protoc-gen-validate v0.3.0-java.0.20200311152155-ab56c3dd1cf9 | Apache-2.0, MIT, BSD-2-Clause                                                  |
| github.com/ghodss/yaml v1.0.0                                                       | BSD-3-Clause, BSD-3-Clause OR MIT                                              |
| github.com/go-logr/logr v0.1.0                                                      | Apache-2.0                                                                     |
| github.com/go-logr/zapr v0.1.0                                                      | Apache-2.0                                                                     |
| github.com/golang-migrate/migrate/v4 v4.8.0                                         | MIT                                                                            |
| github.com/golang/protobuf v1.3.2                                                   | BSD-3-Clause                                                                   |
| github.com/hoisie/mustache v0.0.0-20160804235033-6375acf62c69                       | MIT                                                                            |
| github.com/kelseyhightower/envconfig v1.4.0                                         | MIT                                                                            |
| github.com/lib/pq v1.1.1                                                            | BSD-2-Clause, MIT                                                              |
| github.com/onsi/ginkgo v1.12.0                                                      | MIT                                                                            |
| github.com/onsi/gomega v1.9.0                                                       | MIT-style                                                                      |
| github.com/patrickmn/go-cache v2.1.0+incompatible                                   | MIT                                                                            |
| github.com/pkg/errors v0.8.1                                                        | BSD-2-Clause, BSD-3-Clause                                                     |
| github.com/prometheus/common v0.4.1                                                 | Apache-2.0                                                                     |
| github.com/prometheus/prometheus v0.0.0-00010101000000-000000000000                 | Apache-2.0                                                                     |
| github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749                       | MIT                                                                            |
| github.com/spf13/cobra v0.0.5                                                       | Apache-2.0, Apache-2.0 OR BSD-3-Clause OR GPL-3.0-only OR LGPL-3.0-only OR MIT |
| github.com/spiffe/go-spiffe v0.0.0-20190820222348-6adcf1eecbcc                      | Apache-2.0                                                                     |
| github.com/spiffe/spire v0.0.0-20190905203639-e85640baca1d                          | BSD-3-Clause, Apache-2.0                                                       |
| go.uber.org/multierr v1.1.0                                                         | MIT                                                                            |
| go.uber.org/zap v1.9.1                                                              | MIT                                                                            |
| google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55                       | Apache-2.0                                                                     |
| google.golang.org/grpc v1.23.0                                                      | BSD-3-Clause                                                                   |
| gopkg.in/yaml.v2 v2.2.8                                                             | BSD-3-Clause, BSD-3-Clause OR MIT                                              |
| k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b                                       | Apache-2.0                                                                     |
| k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d                              | Apache-2.0, BSD-3-Clause                                                       |
| k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible                 | Apache-2.0                                                                     |
| k8s.io/helm v2.14.3+incompatible                                                    | Apache-2.0, Apache-2.0 OR BSD-3-Clause                                         |
| sigs.k8s.io/controller-runtime v0.2.2                                               | Apache-2.0                                                                     |
| sigs.k8s.io/testing_frameworks v0.1.1                                               | Apache-2.0                                                                     |
| github.com/onsi/ginkgo v1.10.2                                                      | MIT                                                                            |
| github.com/onsi/gomega v1.7.0                                                       | MIT-style                                                                      |

**_List any dependent repos (upstream/downstream) that are required to build the project (including but not limited to libraries, commercial tools, plugins)_**

The Kuma GUI, which is by default embedded into Kuma, is been developed at https://github.com/Kong/kuma-gui

**_What actions are required to be compliant with the IP policy?_**

Transfer ownership of Kuma's trademark and logo assets to the Linux Foundation.

# Other Considerations

## _Please note, these are not gating criteria but rather to:_

* **_Collect a standard set of information for each project_**
* **_Provides a point in time capture of the state of the project which makes it easier to track progress at future reviews and / or promotion_**
* **_Help projects to prepare for SIG and TOC presentation_**
* **_Allow the SIG to review the project and perform due diligence for incubation_**
* **_Provide the TOC with the information required to accept sponsorship of a project and/or votes_**
* **_Identify and rectify any significant issues / blockers prior to presenting to the TOC and acceptance as a CNCF project_**

## Cloud Native

**_Does the project meet the definition of Cloud Native? The CNCF charter states:_**

```
“Cloud native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds. Containers, service meshes, microservices, immutable infrastructure, and declarative APIs exemplify this approach.


“These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.”
```

Yes.

## Project and Code Quality

**_Are there any metrics around code quality? Are there good examples of code reviews? Are there enforced coding standards?_**

We run static analysis on every PR, [this](https://github.com/Kong/kuma/pull/633) is an example of code review and we try to address codes styling in those reviews.

**_What are the performance goals and results? What performance tradeoffs have been made? What is the resource cost?_**

The main goal for any Kuma user is to make sure that the service-to-service communication is as fast as it can get: for this goal we are relying fully on Envoy, since Kuma itself is not on the execution path of the service traffic. No trade-offs have been made when it comes to Kuma itself, but as you can imagine there is always room for improvements (like pushing only selective changes to the underlying Envoy proxies as opposed to dynamically generating the entire configuration at once).

**_What is the CI/CD system? Are there code coverage metrics? What types of tests exist?_**

We are using CircleCI, and we provide both unit tests and integration tests when it comes to the features that have been built. The tests are executed on every PR, and they are a criteria of acceptance in order to be merged.

**_Is there documentation?_**

Yes, primarily at [kuma-website](https://github.com/Kong/kuma-website) in the `docs` folder. Also accessible at https://kuma.io/docs/.

**_How is it deployed?_**

By using the official release tooling in the `tools` folder. Kuma is being compiled on a variety of different operating systems by leveraging Docker, and then each resulting output is being compressed and made available to the end-user.

**_How is it orchestrated?_**

Kuma can be fully orchestrated in Kubernetes when running in Kubernetes mode, and can be integrated via its RESTful API or native CLI on other environments when running in Universal mode. The `kumactl` CLI tool also provides helpers in Kubernetes to install Kuma, to install Prometheus + Grafana as well as Jaeger for tracing.

**_How will the project benefit from acceptance into the CNCF?_**

We will be providing a vendor-neutral home for the project, and since we will be donating the first control plane with native Envoy support to the community, we would like Kuma to become a neutral control plane primitive for higher level use-cases that the industry may want to build on top. Kuma will also benefit from being considered by other CNCF projects when it comes to providing integration points to their technology.

**_Has a security assessment by the security SIG been done? If not, what is the status/progress of the assessment?_**

No.

## Promotion to Incubation

### Open Governance

**_How are committers chosen?_**

Using the following criteria:

* Commitment to the project
* Ability to write good code
* Ability to collaborate with the team
* Understanding of how the team works
* Understanding of the project's code base and coding style

**_How are architectural and roadmap decisions made?_**

By opening GitHub issues to discuss proposals, as well as bringing up discussion topics in the bi-weekly community calls.

**_How many decision makers are outside the sponsoring organization._**

None at the moment, although we have one committer that has demonstrated continued contributions to the project that is potentially on track to be added as an additional maintainer.

## Adoption

**_Who are the current maintainers?_**

* Jakub Dyszkiewicz, @jakubdyszkiewicz
* Ilya Lobkov, @lobkovilya
* Marco Palladino, @subnetmarco

**_How long has the project been developed for?_**

Since September 10th, 2019.

**_Is there a commercial version of the project or a primary commercial sponsor?_**

There is no commercial version of the project. Kong Inc is the primary commercial sponsor.

**_Is the project used in production? If so, please list some of the accounts._**

It's being used by a series of companies, including:

* Kong
* Sabre

Other enterprise companies are using Kuma, but we cannot disclose their names because of legal contractual obligations. They will have to come forward and disclose it themselves.

**_Does the project participate in a CNCF User Group?_**

No.

## Vendor Independence

**_Is the project reasonably independent from the sponsoring vendor?_**

Kong Inc - the original author of the project - has sponsored most of the development and in preparation to the CNCF donation we have been promoting project independence via public open governance guidelines and community calls. We are open to accept new maintainers to the project and we welcome every contributions that meets the project guidelines.

**_Are all communication channels and project resources hosted just for this project or with other CNCF projects/resources?_**

We use a Kuma [slack channel](https://kuma.io/community).

**_Is all code that is part of the project hosted and part of the CNCF managed orgs and repos?_**

Not at moment since we are in the process of donating, but it can obviously be after the donation.

**_Are all defaults for upstream reporting either unset or community hosted infrastructure (i.e. doesn’t point to vendor hosted SaaS control plane or analytics server for usage data)? Is all project naming independent of vendors?_**

We are using Google Analytics and Splunk on our accounts, but those can certainly be transferred.

**_Relevant Assets regarding vendor independence_**

* [https://kuma.io](https://kuma.io)
* [https://github.com/Kong/kuma](https://github.com/Kong/kuma)
* [https://github.com/Kong/kuma-website](https://github.com/Kong/kuma-website)
* [https://github.com/Kong/kuma-gui](https://github.com/Kong/kuma-gui)
* [https://twitter.com/kumamesh](https://twitter.com/kumamesh)
* [https://kuma-mesh.slack.com/](https://kuma-mesh.slack.com/)