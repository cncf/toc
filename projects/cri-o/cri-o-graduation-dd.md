# CNCF TOC


# CRI-O Due Diligence 


## Promotion to CNCF Graduation

**Authors:** Peter Hunt (Red Hat), Mrunal Patel (Red Hat), Sasha Grunert (Red Hat), Karena Angell (Red Hat)

**TOC Sponsor(s):** Cathy Zhang (TOC), Nikhita Raghunath (TOC)


# CRI-O Graduation Proposal


## This Document

The purpose of this document is to provide the CRI-O project technical review Due Diligence [described here](https://github.com/cncf/toc/blob/master/process/dd-review-template.md).
Its goal is to provide data to help the CNCF TOC to decide on promoting CRI-O to Graduation.


## Background

* [Link to TOC PR](https://github.com/cncf/toc/pull/917)
* [Link to Presentation](https://docs.google.com/presentation/d/1DwMfxbwJxeJyp-d-kpVcYvnQthwAjw8cnLMs9KQVK7o/edit?usp=sharing)
* [Link to GitHub project](https://github.com/cri-o)


## Project Introduction

CRI-O is a project that has one aim: being a secure, performant, and stable Container Runtime Interface ([CRI](https://kubernetes.io/docs/concepts/architecture/cri/))
implementation for the Kubelet to orchestrate Open Container Initiative (OCI) containers in production Kubernetes environments.
Specifically, CRI-O is primarily responsible for managing three types of resources:

* Images
    * Authenticating with an OCI registry.
    * Verifying OCI image integrity.
    * Pulling OCI images.
    * Provisioning disk resources to store images.
* Pods
    * Creating and managing the lifecycle of pod level resources like namespaces and network resources
        * For networking resources, CRI-O delegates the configuration to [CNI](https://cni.dev)
    * Providing resource usage statistics for a collection of containers.
* Containers
    * Translating a CRI request to create a container into an OCI runtime specification.
    * Provisioning disk resources for storing OCI containers.
    * Creating and managing the lifecycle of OCI containers through an OCI runtime.
    * Monitor and redirect output of containers to a specified log location.
    * Catching and notifying the Kubelet of when OCI containers stop.
    * Providing resource usage statistics about single containers.
    * Allowing to run commands within containers, attach to their output, and forward ports to them.

CRI-O also provides extended metrics and profiling data beside fulfilling the Container Runtime Interface and managing the above resources.
For example, CRI-O ships a set of [default Prometheus metrics](https://github.com/cri-o/cri-o/blob/main/tutorials/metrics.md) as well as
[support for OpenTelemetry traces](https://github.com/cri-o/cri-o/blob/main/tutorials/tracing.md) for extended debugging of containers.


### Architecture

Below can be seen an architecture diagram of CRI-O. CRI-O’s responsibilities begin with requests from the Kubelet, which happen over
[gRPC](https://grpc.io) and are defined by the Kubernetes [CRI](https://kubernetes.io/docs/concepts/architecture/cri/) protocol.
There are two pieces of the CRI: the image service and runtime service.

The image service is responsible for image related operations, like pulling, listing and removing images.
CRI-O uses the [OCI registry authentication definition](https://github.com/opencontainers/distribution-spec/blob/main/spec.md), and the [OCI image definition](https://github.com/opencontainers/image-spec/blob/main/spec.md)
to conform to industry-standard image moving and storing methods.
In addition to docker registry authentication, CRI-O also has the option to query [sigstore](https://sigstore.dev) to verify image signatures.
CRI-O uses the [containers image](https://github.com/containers/image) library to do this image pulling, which is a library shared by projects like
[Podman](https://podman.io), [Buildah](https://buildah.io) and [Skopeo](https://github.com/containers/skopeo). 

The runtime service is responsible for both container and pod related operations, for instance running and removing pods,
and creating, starting and removing containers. It uses a utility called [pinns](https://github.com/cri-o/cri-o/tree/release-1.26/pinns)
to create namespaces for the pods, and [CNI](https://cni.dev) to create the networking resources for pods.
It also uses an OCI compliant runtime like [runC](https://github.com/opencontainers/runc), [crun](https://github.com/containers/crun)
or [Kata Containers](https://katacontainers.io/) to do many container operations.
It also uses a utility called [conmon](https://github.com/containers/conmon), or its successor [conmon-rs](https://github.com/containers/conmon-rs)
to monitor containers’ lifecycle: including catching when they exit and saving the output in a specified format for the CRI.

![CRI-O Architecture](https://github.com/haircommander/toc/blob/pics/proposals/graduation/cri-o-arch.png)

Architecturally, it could be considered very similar to other container managers like Podman, [docker](https://www.docker.com/),
or [containerd](https://containerd.io). However one key difference separates CRI-O from the rest:
All of these functions are customized for the requirements of Kubernetes.

This customization applies to its security posture:

* Since CRI-O only supports Kubernetes, the only avenue an unprivileged malicious user has to attack it is to trick the Kubelet into causing CRI-O to perform the exploit. This severely limits CRI-O’s attack surface.
    * For instance, of the results the [audit](https://docs.google.com/document/d/1eQkSNdXmk4qIA1ghu3NGQ6xCF7-xqxIyYFRYVyit7Gc/edit?usp=sharing) found, the only vulnerability that was high risk was exploitable by a user with access to the Kubernetes API.
    * Note: direct access to the CRI-O socket by a container is considered a privileged operation and should be limited by Kubernetes constructs like [Pod Security Admission](https://kubernetes.io/docs/concepts/security/pod-security-admission/).
* It tightens cluster security by implementing features in a way that is easily configurable by a Kubernetes admin/user.
    * User namespaces are able to be configured via Kubernetes annotation
    * Includes a configuration knob: `seccomp_use_default_when_empty` that turns on seccomp even if the user didn’t specify it.
    * Ships containers with a reduced Linux capability set by default.
* It has a reduced set of operations compared with other container runtimes, as these capabilities are not needed by the Kubelet, and thus has an even smaller attack surface. Note: the smaller attack surface doesn’t directly mean CRI-O is more secure than other container managers. However, having a smaller scope means there is less code in general to be insecure, as well as it reduces the maintenance complexity of keeping that code secure. Some of the capabilities absent from CRI-O that are present in other container managers are:
    * Building images, and any image creation/manipulation operation
    * Pushing images, and any image moving operation
    * Different clients having different pods and containers (containerd namespaces)
    * Pod and container events outside of those needed for EventedPLEG
    * Windows container support
    * Limited HTTP API surface, with operations present to support [cAdvisor](https://github.com/google/cadvisor) or [Kubevirt](https://kubevirt.io/) 
        * HTTP API can be seen [here](https://github.com/cri-o/cri-o/blob/a67e7776c0d82d380e5677040ee114cb85d1c178/server/inspect.go#L123). The inspect endpoints are for cAdvisor to query for containers that exist to relate them to newly created cgroups. The pause/unpause endpoints are present to allow for Kubevirt to scale up/down pods quickly.
    * A dedicated CLI. Users who wish to debug against CRI-O can use [crictl](https://github.com/kubernetes-sigs/cri-tools/tree/master/cmd/crictl)

This customization applies to its performance:

* It is optimized for operations that happen frequently in Kubernetes.
    * A very common operation the Kubelet performs (before the EventedPLEG feature) is asking the CRI implementation to list pods, containers and images. Instead of constantly querying the OCI runtime and disk for the state of those resources, CRI-O uses a combination of caching and inotify to make returning the state as fast as possible.
* It takes advantage of specific Kubelet behavior to optimize in heavy node load situations.
    * The Kubelet requests to create a pod or container with a timeout to make sure the CRI implementation receives the request. However, in situations of load on the node (for instance, disk throttling or when the networking plugin is overwhelmed), the creation of a pod or container can take longer than that timeout. With no special behavior, this causes the Kubelet and CRI implementation to go back and forth requesting and failing until the pod or container is created. Custom CRI-O behavior throttles the Kubelet’s requests and reduces the amount of work happening to make sure the pods and containers start as quickly as possible.
* It optimizes for common CNI plugin patterns.
    * Often, CNI plugins are a daemonset running on nodes that register themselves with a runtime by dynamically placing a CNI config once the pod is ready. However, the Kubelet doesn’t have any special knowledge about which pods are related to the CNI plugins. CRI-O has customized behavior on node startup that waits to attempt to create pods that require the CNI plugin to be up, which allows nodes to bootstrap faster.
* It can add features configurable through the Kubernetes API with pod annotations that give additional flexibility to pod authors.
    * Kubernetes added support for user namespaces in 1.25.0. However, CRI-O added an experimental annotation `io.kubernetes.cri-o.userns-mode` in 1.20.0, which helped inform the eventual pod API, and gave users access to user namespaces in Kubernetes early.

This customization applies to its stability:

* It is released in lock-step with Kubernetes (following the same major/minor release cycle).
* It is battle tested thousands of test cases:
    * Kubernetes e2e suite, verifying Kubernetes conformance
    * SIG-Node conformance test suite, testing particular behavior with the kubelet
    * critest suite, verifying conformance with the CRI spec
    * Integration tests that test CRI-O specific e2e behavior
    * Unit tests that verify CRI-O specific functional behavior 

To achieve this, it utilizes industry-standard specifications and tools:

* Utilizes Open Containers Initiative (OCI) runtime and image specs to be interoperable with other container tools.
* Built in tandem with the Kubernetes Container Runtime Interface (CRI) to ensure compatibility with its one purpose: Kubernetes
* Delegates to CNI plugins for configuring pod networking resources.

Ultimately, despite all of these interesting properties, CRI-O aims to serve its purpose in a stable_boring_ way.
Users should be able to install and configure CRI-O to their Kubernetes clusters, and promptly be able to forget about it.

CRI-O was born in the Kubernetes incubator in 2016. It grew to production stable there,
and was inducted into the CNCF as an incubating project in 2019. Since then, CRI-O has only further stabilized and delivered value.
It has been adopted by 7 new organizations, runs on tens of thousands of clusters, has released 11 new minor versions,
around 100 patch releases, and has had more than 4000 commits to the main branch.


## Future Plans

The initial roadmap of CRI-O was lightweight and followed the main Kubernetes Container Runtime Interface (CRI) development lifecycle.
This is partially because additional features on top of that are either integrated into a CRI-O release as soon as they’re ready,
or are tracked through the Milestone mechanism in GitHub. Another reason is that feature availability is mostly tied to Kubernetes releases,
and thus most of its long-term goals are already tracked in SIG-Node through the Kubernetes Enhancement Proposal (KEP) process.
Finally, CRI-O’s long-term roadmap outside of features being added by SIG-Node is in part described by its mission: to be a secure,
performant and stable implementation of the Container Runtime Interface. However, all of these together do construct a roadmap,
and [the referencing GitHub document](https://github.com/cri-o/cri-o/blob/main/roadmap.md) is describing how.


## Project Scope


### Value-add to the CNCF ecosystem

CRI-O is the only remaining (the now archived [rkt](https://github.com/rkt/rkt) had a similar goal) container manager whose purpose
is strictly to be a CRI implementation for Kubernetes. It is able to provide users with a simple, clear, and performant container manager
for their cloud native workloads (all of which are built upon some entity like CRI-O).

While there are multiple CRI implementations, and another already exists in the CNCF landscape: [containerd](https://containerd.io),
having multiple implementations furthers the idea of the Kubernetes CRI being an interface for a user to plug in a container runtime into the Kubelet.
Without multiple implementations, the sole implementation of the interface could cause unintended implementation dependencies,
where the sole implementation has hidden behavior that the Kubelet relies on, raising the barrier of entry for new implementations.
Having multiple implementations also motivates collaboration in the open in Kubernetes SIG-Node,
where representatives from both runtimes come together to achieve the shared goal of making the CRI expressive and complete for all user workload needs.
These two communities also collaborate on multiple pieces of code together, like
[Container Device Interface](https://github.com/container-orchestrated-devices/container-device-interface),
[Node Resource Interface](https://github.com/containerd/nri/), and also use shared components like [runc](https://github.com/opencontainers/runc)
and other projects in the Open Containers Initiative. Finally, there have been multiple instances of the two communities collaborating on security patches,
like in [CVE-2022-1708](https://github.com/cri-o/cri-o/security/advisories/GHSA-fcm2-6c3h-pg6j)
and [CVE-2022-27652](https://github.com/cri-o/cri-o/security/advisories/GHSA-4hj2-r2pm-3hc6).
All of this collaboration paints a picture of two communities that can work together through friendly competition to ensure the CRI specification
and OCI ecosystem are healthy and have sound implementations.


### Alignment with other CNCF projects

CRI-O occupies the same piece of the Kubernetes stack as other implementations of the CRI – most notably [containerd](https://containerd.io).
Both communities work together under Kubernetes SIG-Node to define the CRI spec (which is a protocol that uses [gRPC](https://grpc.io/))
and ensure compatibility between implementations. CRI-O also utilizes [CNI](https://cni.dev) for provisioning of networking resources of the pods.
CRI-O integrates with both [Prometheus](https://prometheus.io) and [OpenTelemetry](https://opentelemetry.io) for reporting of metrics and tracing.


## Formal Requirements

Below is evidence proving CRI-O meets the formal
[CNCF graduation criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#graduation-stage).


#### Have committers from at least two organizations.

CRI-O has had major (more than 10 commits all-time) contributions from the following organizations:

- Red Hat: 27
- Intel: 3
- SUSE: 3
- Independent: 3
- IBM: 1

The company breakdown of the OWNERS file is:

Approvers:
- Red Hat: 7
- Intel: 1

Reviewers:
- Red Hat: 2
- Intel: 1

While the formal requirement of graduation criteria is two organizations being a committer, which is clearly the case,
the spirit of the requirement seeks to prove the project is company agnostic.
The above committer breakdown paints a picture of a Red Hat dominated project, but it’s not necessarily so.
While Red Hat has invested heavily in CRI-O strategically, looking at the
[dev stats](https://crio.devstats.cncf.io/d/5/companies-table?orgId=1), CRI-O has also been heavily invested in by SUSE and some by Intel and IBM.
While the recent [dev stats](https://crio.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions)
indicate very heavy Red Hat investment, this can also be attributed to the general stabilizing of CRI-O, which needs fewer features as time goes on.

Further, something that is not reflected in the github stats are the attendance at the bi-weekly meetings,
where community members and maintainers are regularly in attendance representing both Red Hat, IBM, and Intel.

In general, it is not the desire of the CRI-O team to have such lop-sided company involvement.
The community has made strides to expand its contributor base, including:

* Refining [GOVERNANCE](https://github.com/cri-o/cri-o/blob/main/GOVERNANCE.md) and [CONTRIBUTING.md](https://github.com/cri-o/cri-o/blob/main/CONTRIBUTING.md) documentation.
    * In particular, making the distinction between community member, organization member, reviewer and approver, to provide a ladder for new contributors to work towards approver.
    * Since incubation, CRI-O has added reviewer and organization member tiers, and further defined the rights and responsibilities of them, which opens the path towards approver status in the project to newcomers.
* Tagging issues as `good first issue` and `help wanted` when applicable.
* Joining a bug bash during Kubecon US 2021
* Regularly requesting submitters of issues whether they would like to be mentored in the process of fixing their issue. Below are some examples:
    * [https://github.com/cri-o/cri-o/issues/5862](https://github.com/cri-o/cri-o/issues/5862) 
    * [https://github.com/cri-o/cri-o/issues/5736](https://github.com/cri-o/cri-o/issues/5736)
    * [https://github.com/cri-o/cri-o/issues/5578](https://github.com/cri-o/cri-o/issues/5578)
    * [https://github.com/cri-o/cri-o/issues/5030](https://github.com/cri-o/cri-o/issues/5030) 
    * [https://github.com/cri-o/cri-o/issues/4046](https://github.com/cri-o/cri-o/issues/4046)
    * [https://github.com/cri-o/cri-o/issues/3916](https://github.com/cri-o/cri-o/issues/3916)
    * [https://github.com/cri-o/cri-o/issues/1400](https://github.com/cri-o/cri-o/issues/1400) 
* Regularly advertising CRI-O during Kubecon, through both talks and office hours, to help move interested parties towards maintainership.
* Participating in Google Summer of Code in 2023.
* Encouraging new contributors to join the CRI-O organization with an encouraging note on github issues or through mentoring via slack.
* In the last year, only [23%](https://all.devstats.cncf.io/d/53/projects-health-table?orgId=1000) of code contributions came from Red Hat, which shows the non-Red Hat community is growing.


#### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

Here is CRI-O’s best practices badge:

[https://bestpractices.coreinfrastructure.org/en/projects/2298](https://bestpractices.coreinfrastructure.org/en/projects/2298)


#### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

In the spring of 2022, the CNCF sponsored a [security audit](https://github.com/cri-o/cri-o/blob/main/security/2022_security_audit_adalogics.pdf) of CRI-O,
coordinated by OSTIF and completed by Ada Logics and Chainguard. It included fuzzing, manual investigation, and a review of CRI-O’s software supply chain.
In it, Ada Logics found [one](https://github.com/cri-o/cri-o/security/advisories/GHSA-fcm2-6c3h-pg6j) high level vulnerability
(which also helped find a similar vulnerability in containerd), and a couple of low/medium ones.
All of the high/medium ones have been addressed (the linked advisory has reference to the fixed versions of the high risk one),
and the low priority are being tracked.
The number of high priority vulnerabilities were so limited in part due to CRI-O’s threat model of being scoped to the needs of Kubernetes.
The only high risk vulnerabilities are ones where a malicious user is able to trick the Kubelet into tricking CRI-O with a CRI request.
All the low/medium ones are not exploitable in production.

Further, CRI-O is in the process of evaluating its supply chain,
recently integrating sigstore validation and fulcio signing for its packaged and static artifacts as part of the response to the audit.

In summarizing CRI-O’s security posture, Ada Logics said: “CRI-O is a well-written project that has a high level of security assurance”.

The standardized CRI-O security process can be found in the [SECURITY.md](https://github.com/cri-o/cri-o/blob/main/SECURITY.md) in the CRI-O GitHub repository.


#### Explicitly define a Project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

- [GOVERNANCE.md](https://github.com/cri-o/cri-o/blob/main/GOVERNANCE.md)

- [OWNERS](https://github.com/cri-o/cri-o/blob/main/OWNERS)

- [CONTRIBUTING.md](https://github.com/cri-o/cri-o/blob/main/CONTRIBUTING.md)

Generally, the CRI-O project is run by the approvers. They have jurisdiction over the code changes, governance processes, and direction of the project.
Decisions are largely made as described in the roadmap (described in brief in the future work section above),
where primarily CRI-O aims to conform to the Kubernetes CRI, and then secondarily aims to do so in an efficient and secure way.
As such, code direction is largely dictated by those aims, and there have not been many instances of a direction decision that needed to be made. 

An example of a decision not directly related to upstream SIG-Node was to implement [conmon-rs](https://github.com/containers/conmon-rs)
as opposed to the legacy [conmon](https://github.com/containers/conmon). The direction was decided upon by the approvers,
then discussed in the community meetings, and then development was taken upstream to the repository.


#### Have a public list of Project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the Project website). For a specification, have a list of adopters for the implementation(s) of the spec.

CRI-O has a number of end-users. Those that have wished to publicly acknowledge their adoption of CRI-O are included in the
[ADOPTERS](https://github.com/cri-o/cri-o/blob/main/ADOPTERS.md) file. Users are encouraged to add themselves to the list by
submitting a pull-request in the [README](https://github.com/cri-o/cri-o/blob/main/README.md#adopters)


## Incubation Details

_**Project should address each area listed below**_


#### Link to Incubation Due Diligence(DD) Document

CRI-O was inducted in incubation stage in [April of 2019](https://github.com/cncf/toc/blob/main/proposals/incubation/cri-o.adoc).


#### Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

CRI-O being inducted to incubation predates the DD process. The initial PR and discussion can be found [here](https://github.com/cncf/toc/pull/177),
and there do not seem to have been follow-up items that needed addressing.
