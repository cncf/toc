# Krustlet

## Description

Krustlet is a Kubernetes Kubelet implementation for running WebAssembly (Wasm) modules on nodes.
Krustlet acts as a Kubelet by listening on the event stream for new pods that the scheduler assigns
to it based on specific Kubernetes tolerations. It also can be extended to enable other types of
non-Wasm runtimes using a "Provider Model" similar to [Virtual
Kubelet](https://github.com/virtual-kubelet/virtual-kubelet), but its originating use case is for
Wasm.

## Project Goal

The primary goal of Krustlet is to be the easiest way to integrate Wasm support as seamlessly as
possible into Kubernetes. It will allow for users to run heterogeneous clusters where both
containers and Wasm modules can run (though not necessarily on the same node)

## Current Status

Krustlet is very close to a 1.0 release. It was created in late 2019 and has grown along with the
rest of the Wasm community. Essentially 1.0 is stating that that basic features (but not all
advanced features such as pod resources, full networking stack integration, health checks) for a
kubelet are in place and the software is reliable enough to use in production. Continued growth of
the project will include implementing more of the advanced kubelet features and tracking upstream
changes to the the WebAssembly Systems Interface spec (WASI).

## Value for the CNCF ecosystem

Krustlet is a great fit for the CNCF ecosystem. First and foremost, it is a Kubernetes-related
project. That in and of itself is not enough to justify inclusion, but Krustlet has proven to be an
interesting alternative runtime for Kubernetes and has driven interest in the wider Kubernetes
community around using Wasm and other non-container workloads. As Sandbox projects are meant to
"[e]ncourage public visibility of experiments or other early work that can add value to the CNCF
mission and build the ingredients of a successful Incubation level project," Krustlet is a natural
fit as Wasm is an emerging technology that many in the CNCF have expressed an interest in. As seen
at the latest KubeCon EU, community interest in Wasm grew to the point where it merited its own
adjacent event. Given the rapidly evolving Wasm ecosystem, it puts Krustlet in a common place where
it can facilitate alignment with the various projects and ideas coming from that evolution.

## Community Size and Existing Sponsorships

At the time of writing, the Krustlet project has 1.7k stars on GitHub and 39 different code and
documentation contributors. With the release of 1.0 we expect project adoption and additional
community participation to increase.

Currently, the project is sponsored by Microsoft

## Current Maintainers

The project currently has 6 active core maintainers:

- Ivan Towlson (Microsoft)
- Matt Butcher (Microsoft)
- Matt Fisher (Microsoft)
- Kevin Flansburg (Moose Consulting)
- Radu Matei (Microsoft)
- Taylor Thomas (Microsoft)

## Source code location

Currently, the source code is located at https://github.com/deislabs/krustlet.

This will be moving shortly to https://github.com/krustlet/krustlet before acceptance as a sandbox
project

### Krator

The Krustlet repository currently contains the
['krator`](https://github.com/deislabs/krustlet/tree/main/crates/krator) and
[`krator-derive`](https://github.com/deislabs/krustlet/tree/main/crates/krator-derive) projects.
These are currently in the process of being split out into their own project that will also be
proposed separately for addition into the Sandbox. They will not be coming along as part of the
Krustlet project, although they will remain related.

## Sponsor from TOC

TBD

## License

Apache 2.0

## Website

https://krustlet.dev

## Social Media Accounts

Twitter: @krustlet

## Infrastructure Needs

Krustlet has a custom Windows GitHub Actions runner used for e2e testing that is provided by
Microsoft Azure.
