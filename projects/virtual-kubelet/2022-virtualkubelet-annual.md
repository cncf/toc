# VirtualKubelet 2021 Annual Review

## Background

Virtual Kubelet enables clients to register a node with Kubernetes but have
the actual node implementation and how to run a pod deferred to the client.
The goal of Virtual Kubelet is to provide tooling to build custom pod
runtimes instead of providing the runtime itself.

## Devstats

[Dashboard](https://virtualkubelet.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m)

The maintainers work to get more users and contributors to join the project, that has paid of with the amount 
of changes to the core have started to increase again in the last year with more efforts close more regression 
issues that have introduced with Kubernetes v1.24. In addition, the project started to get more contributions 
from new users, besides working on the stabilization and bug fixes now that Virtual Kubelet is running in more 
environments and with different workloads.

## Maintainers

Virtual Kubelet has 4 core maintainers from Microsoft, VMWare, Netflix, and
an independent. Additionally there are several projects within the VK org
which are specific to a VK implementation (such as azure ACI) which have
their own maintainers.

[VK Owners](https://github.com/virtual-kubelet/community/blob/master/OWNERS.md)

## Adoption

Virtual Kubelet has been adopted by a number of projects which use VK to implement their own pod runtime.
Some of these have been adopted into the VK org for the purposes of governance and gathering community.

Projects using VK:

- [Azure ACI](https://github.com/virtual-kubelet/azure-aci) Use's Azure's ACI container service to run pods
- [Systemk](https://github.com/virtual-kubelet/systemk) Uses systemd to run pods
- [Nomad](https://github.com/virtual-kubelet/nomad) Run pods using nomad
- [Tensile Kube](https://github.com/virtual-kubelet/tensile-kube) Run pods on another Kubernetes cluster

There is also wider adoption outside of these projects. As an example,
Netflix uses their own Virtual Kubelet implementation finely tuned to their
environment to run pods on Docker. If you have watched Netflix, you have hit VK.

New providers have been introduced, `StackPath` and `Liqo` providers.

## Goals

In the past year we have two new releases and see how things perform in newly adopted environments. We 
are now looking at what does "2.0" look like, i.e. where does VK fall short such that it requires breaking 
changes to the core functionality and fix regression issues. Overall we want to work towards an interface 
that is less runtime and more composability, more helpers, downward APIs etc. We want to provide a good set 
of functionality that implementers can take and use directly, build their own, or even compose together.
