# WasmEdge Runtime-2023-Annual Review

## Background

WasmEdge is a lightweight, high-performance, and extensible WebAssembly runtime. The ultimate goal for WasmEdge is to bring cloud-native and serverless application paradigms to edge, streaming, web3, and SaaS computing environments.

The unique features and advantages of WasmEdge include the following.

* One of the [fastest](https://github.com/WasmEdge/WasmEdge#performance) WebAssembly VMs on the market (based on **LLVM AOT**) See another bechmark [here](https://00f.net/2023/01/04/webassembly-benchmark-2023/).
* WasmEdge feature extensions
  * **Network sockets** ([Rust](https://github.com/second-state/wasmedge_wasi_socket), C and [JavaScript](https://github.com/second-state/wasmedge-quickjs#http-request) SDKs)
  * **Async polling** (for Rust Future and JS async)
  * Tensorflow, Pytorch, OpenVINO inference ([Tutorial](https://wasmedge.org/docs/category/neural-networks-for-wasi))
    * [Mediapipe-rs support for Mediapipe models](https://github.com/WasmEdge/mediapipe-rs)
  * Key value storage
  * Database connector including MySQL-based databases, Redis and PostgreSQL. ([Tutorial](https://wasmedge.org/docs/category/database-drivers))
  * **Gas meters** for resource constraints
* JavaScript support
  * [ES6 module](https://github.com/second-state/wasmedge-quickjs#es6-module-support) and std API support
  * **Implement JS APIs in Rust** ([Tutorial](https://www.secondstate.io/articles/embed-rust-in-javascript/))
  * Import C native shared library functions as JS functions ([Tutorial](https://www.secondstate.io/articles/call-native-functions-from-javascript/))
* Cloud native management and orchestration
  * [CRI-O and Kubernetes compatibility](https://wasmedge.org/docs/category/deploy-wasmedge-apps-in-kubernetes)
  * Docker Desktop integration
  * Sidecar apps in Kubernetes-based service meshes
  * Dapr microservices with Dapr WasmEdge SDK ([Tutorial](https://github.com/second-state/dapr-sdk-wasmedge))
* Cross-platform support
  * Linux OSes dated back to 2010 for both x86 and ARM CPUs
  * Mac OS X for both x86 and M1/M2 CPUs
  * Windows
  * Microkernel and RTOS (e.g., the highly secure [seL4 microkernel](https://github.com/second-state/wasmedge-sel4))
  * OpenWRT
  * RISC-V
  * Android
* Easy extensibility
  * Extend and enhance the runtime with custom functions in [C](https://wasmedge.org/docs/embed/c/reference/latest#host-functions) or [GO](https://wasmedge.org/docs/embed/go/reference/latest#host-functions))
* Easy to embed into a host application
  * Embed WasmEdge functions in [C](https://wasmedge.org/docs/category/c-sdk-for-embedding-wasm-functions), [Go](https://wasmedge.org/docs/category/go-sdk-for-embedding-wasm-functions), and [Rust](https://wasmedge.org/docs/category/rust-sdk-for-embedding-wasm-functionst)-based host applications.
  * Embed WasmEdge functions in service mesh proxies (e.g., [proxy-wasm](https://github.com/proxy-wasm/proxy-wasm-cpp-host/pull/193) for Envoy and MOSN proxies)


## Sandbox Acceptance

WasmEdge is accepted into the CNCF Sandbox on 27th April 2021.

## Development metrics

The WasmEdge devstats page and dashboard could be found [here](https://wasmedge.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

* The project currently has [177](https://wasmedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from 28 companies. According to devstats, the project averages [789 contributions per month](https://wasmedge.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) within [47+ merged PRs per month](https://wasmedge.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).

* [New PRs in last year](https://wasmedge.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
* The community has grown since WasmEdge entered the CNCF Sandbox.
    * We have held monthly meetings since October 2021. The meeting notes could be found [here](https://docs.google.com/document/d/1iFlVl7R97Lze4RDykzElJGDjjWYDlkI8Rhf8g4dQ5Rk/edit), and the recording meeting video could be found [here](https://www.youtube.com/playlist?list=PL4H9HXvnb_jYbsHBOMk-jayEMGrl09VSO).
    * Number of contributors: [6](https://wasmedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Before%20joining%20CNCF&var-repogroup_name=All) -> 140
    * Number of stars: 819 -> 6200
    * Number of forks: 43 -> 558

## Maintainers

We now have sixteen reviewers, committers, and maintainers from Second State, Bytedance, Giant Network Group Co., Ltd., SRA VJTI, University of Science and Technology of China, Southeast University, Nanjing University and Purdue University. We have established the [community governance guideline](https://github.com/WasmEdge/WasmEdge/blob/master/docs/GOVERNANCE.md) and will promote more contributors to maintainers.

The core maintainers are

* Michael Yuan (Second State) @juntao
* Hung-Ying Tai (Second State) @hydai
* Yi-Ying He (Second State) @q82419
* Shen-Ta Hsieh (Second State) @ibmibmibm

## Project progress and adoption

The WasmEdge community has released ten versions since joining CNCF. 

* [version 0.8.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.0)
* [version 0.8.1](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.1)
* [version 0.8.2](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.2)
* [version 0.9.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.9.0)
* [version 0.10.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.10.0)
* [version 0.11.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.11.0)
* [version 0.11.1](https://github.com/WasmEdge/WasmEdge/releases/tag/0.11.1)
* [version 0.11.2](https://github.com/WasmEdge/WasmEdge/releases/tag/0.11.2)
* [version 0.12.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.12.0)
* [version 0.12.1](https://github.com/WasmEdge/WasmEdge/releases/tag/0.12.1)
* [version 0.13.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.13.0)


With each release, WasmEdge delivers important features that are important to cloud-native users.

* Advanced networking with the WasmEdge sockets library. It supports non-blocking sockets, DNS, TLS. That allows WadmEdge to run modern HTTPS client and server applications. 
* Support for the Rust Tokio library. It allows many Tokio-based async Rust applications, especially async networking apps, to run seamlessly on WasmEdge.
* Container and K8s tooling. WasmEdge is an OCI-compliant runtime, so it could be managed and orchestrated by K8s and K8s tooling, like crun, youki, runwasi, containerd, CRI-O, K8s, Kind, OpenYurt, KubeEdge, SuperEdge.
* An official package in Fedora and Red Hat Enterprise Linux (RHEL) distributions. 
* JavaScript support. WasmEdge now supports the full JavaScript language spec as well as common JavaScript module systems, including ES6, CommonJS (CJS), and NPM modules.
* Support for WebAssembly feature extensions such as WASI-NN and WASI-Crypto.
* Support for optional WebAssembly specs and proposals from the community, including the SIMD proposal, the Tail Call proposal, the extended-const proposal, the multi-memories proposal, the Import/Export mutable globals proposal, the Non-trapping float-to-int conversions proposal, the Sign-extension operators proposal, the Multi-value proposal, the Reference Types proposal, the Bulk memory operations proposal and the threads proposal.
* Improved cross-platform compatibility. WasmEdge added support for macOS, Windows, Andriod, seL4, OpenWRT, RISC-V, and OpenHarmony operating systems, as well as ARM32, ARM64, Apple M1 and M2 hardware.
* Enhanced host SDKs in C, Rust, and Go languages.
* A host / Wasm data binding toolkit called the `wasmedge-bindgen`.
* A new plugin system that makes WasmEdge native host functions and extensions easier to develop and install.
* A lot of performance improvements.

The WasmEdge community has expanded its collaboration with many open-source projects in the cloud-native space in the last year. Some of these projects now have production customers using WasmEdge.

* OpenFunction (CNCF project) supports WasmEdge as a container runtime.
* WebAssembly Languages Runtime, maintained by VMWare, supports WasmEdge as a runtime
* LF Edge eKuiper and shifu uses WasmEdge to process data streamed from IoT devices.
* Kuasar supports WasmEdge as a container runtime.
* containerd's runwasi supports WasmEdge as a container runtime.
* youki, the OCI Container Runtime, supports WasmEdge as a container runtime.
* Essa-rs ueses WasmEdge to execute serverless functions on the essa-rs platform.
* The OpenGauss and libsql communities used WasmEdge to power UDFs (User Defined Functions) in database queries.
* Apache Traffic Server uses WasmEdge as a plugin to process the traffic flow.

Another highlighted collaboration is that Docker Desktop integrates WasmEdge to support Wasm containers.

Furthermore, the WasmEdge project is used by numerous commercial companies including Docker, VMWare, FutureWei, Huawei, 5miles, and ByteDance. You can check out the full list [here](https://wasmedge.org/docs/contribute/users).

## Project goals

This is the WasmEdge project's second Annual review after joining the CNCF Sandbox.

For the immediate future, WasmEdge will focus on the following areas.

* Technical goals
    * Support more WebAssembly proposals and specs, such as threading, stack switching, logging, and debugging. In particular, we aim to be one of the first Wasm runtimes to fully support the WebAssembly component model spec.
    * Improve support for microservices. We will enable WasmEdge applications to access commonly used databases, messaging services, proxies, sidecars, and other services.
    * Improve support for data streaming functions. We will work with data streaming partners to support embedded functions in WasmEdge in those frameworks.
    * Improve support for AI inference workloads in Wasm. We will continue adding support for popular models in computer vision, voice recognition, and document AI. We will also enable inference on LLMs such as Llama.
    * Support AI training on Wasm. We are collaborating with ByteDance and the Ray community to support WasmEdge on Ray nodes.
    * Support the complete Node.js API and NPM module ecosystem in WasmEdge's JavaScript runtime.
    * Improve WasmEdge support in popular application frameworks such as Dapr, Spin, etc.
    * Better support for host SDKs in Rust, Go, Java, and Python for embedding WasmEdge in other applications.
    * Debugging, monitoring, and observability tools for WasmEdge.
    * Further optimize the developer experience and user experience for WasmEdge plugins. That would make it easy for the community to add new features to WasmEdge.


Our 2023 roadmap could be found [here](https://github.com/WasmEdge/WasmEdge/blob/master/docs/ROADMAP.md).

* User adoption goals
    * Drive adoption through integrating more open-source projects. 
    * Continue working closely with end-users to bring WasmEdge into production systems in at least 3-5 commercial companies.

* Community and community collaborations
    * Recruit and support proposal maintainers from more companies other than Second State.
    * More community collaborations with open source projects in Linux Foundation and Apache Foundation.
    * Attending continuously the mentorship programs like [Google Summer of Code](https://summerofcode.withgoogle.com/), [Google Season of Docs](https://developers.google.com/season-of-docs), [LFX mentorship](https://lfx.linuxfoundation.org/tools/mentorship/) and [Open Source Program Plan](https://summer-ospp.ac.cn/) to drive community contributions by students.

## How the CNCF can help to achieve the upcoming goals

In the past two years, the WasmEdge project has received significant support from the CNCF. That includes LFX Mentorship opportunities, publication of tech articles on the CNCF blog, Project Pavilion and Office Hour opportunities at KubeCon, and legal support for the WasmEdge trademark in the US and China. Going forward, we plan to work closely with the CNCF on the following additional areas.

* Strategy to recruit and retain contributors.
* Guidance and support for the project's technical documentation.
* Opportunities to collaborate with other CNCF projects.

## Incubation readiness

The WasmEdge team believes that the project is incubation ready, and had [submitted a incubation proposal](https://github.com/cncf/toc/pull/1102).

