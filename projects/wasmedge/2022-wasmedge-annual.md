# WasmEdge Runtime-2022-Annual Review

## Background

WasmEdge is a lightweight, high-performance, and extensible WebAssembly runtime. It is the fastest Wasm VM today. The WasmEdge Runtime provides a well-defined execution sandbox for its contained WebAssembly bytecode program. The runtime offers isolation and protection for operating system resources (e.g., file system, sockets, environment variables, processes) and memory space. The most important use case for WasmEdge is to safely execute user-defined or community-contributed code as plug-ins in a software product (e.g., SaaS, software-defined vehicles, edge nodes, or even blockchain nodes). It enables third-party developers, vendors, suppliers, and community members to extend and customize the software product.

The ultimate goal for WasmEdge is to bring cloud-native and serverless application paradigms to edge, streaming, web3, and SaaS computing environments.

The unique features and advantages of WasmEdge include the following.

* One of the [fastest](https://github.com/WasmEdge/WasmEdge#performance) WebAssembly VMs on the market (based on **LLVM AOT**)
* WasmEdge feature extensions
  * **Network sockets** ([Rust](https://github.com/second-state/wasmedge_wasi_socket) and [JavaScript](https://github.com/second-state/wasmedge-quickjs#http-request) SDKs)
  * **Async polling** (for Rust Future and JS async)
  * Tensorflow inference ([Tutorial](https://www.secondstate.io/articles/wasi-tensorflow/))
  * Key value storage
  * Database connector
  * **Gas meters** for resource constraints
* JavaScript support
  * [ES6 module](https://github.com/second-state/wasmedge-quickjs#es6-module-support) and std API support
  * **Implement JS APIs in Rust** ([Tutorial](https://www.secondstate.io/articles/embed-rust-in-javascript/))
  * Import C native shared library functions as JS functions ([Tutorial](https://www.secondstate.io/articles/call-native-functions-from-javascript/))
* Cloud native management and orchestration
  * [CRI-O and Kubernetes compatibility](https://www.secondstate.io/articles/manage-webassembly-apps-in-wasmedge-using-docker-tools/)
  * Sidecar apps in Kubernetes-based service meshes
  * Dapr microservices ([Tutorial](https://www.secondstate.io/articles/dapr-wasmedge-webassembly/))
* Cross-platform support
  * Linux OSes dated back to 2010 for both x86 and ARM CPUs
  * Mac OS X for both x86 and m1 CPUs
  * Windows
  * Microkernel and RTOS (e.g., the highly secure [seL4 microkernel](https://github.com/second-state/wasmedge-sel4))
* Easy extensibility
  * Build customized runtimes with native functions in [C](../embed/c/ref.md#host-functions) or [GO](../embed/go/ref.md#host-functions)
* Easy to embed into a host application
  * Embed WasmEdge functions in [C](../embed/c.md), [Go](https://www.secondstate.io/articles/extend-golang-app-with-webassembly-rust/), [Rust](https://github.com/WasmEdge/WasmEdge/tree/master/bindings/rust), [Node.js](https://www.secondstate.io/articles/getting-started-with-rust-function/) and Python host applications
  * Embedded function runtime for service mesh proxies (e.g., [proxy-wasm](https://github.com/proxy-wasm/proxy-wasm-cpp-host/pull/193) for Envoy and MOSN proxies)


## Sandbox Acceptance

WasmEdge is accepted into the CNCF Sandbox on 27th April 2021.

## Development metrics

The WasmEdge devstats page and dashboard could be found [here](https://wasmedge.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

* The project currently has [107](https://wasmedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1) contributors from 28 companies. According to devstats, the project averages [798 contributions per month](https://wasmedge.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) within [35+ merged PRs per month](https://wasmedge.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now).

* [New PRs in last year](https://wasmedge.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1).
* The community has grown since WasmEdge entered the CNCF Sandbox.
    * We have held monthly meetings since October 2021. The meeting note could be found [here](https://docs.google.com/document/d/1iFlVl7R97Lze4RDykzElJGDjjWYDlkI8Rhf8g4dQ5Rk/edit), and the recording meeting video could be found [here](https://www.youtube.com/playlist?list=PL4H9HXvnb_jYbsHBOMk-jayEMGrl09VSO).
    * Number of contributors: [6](https://wasmedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Before%20joining%20CNCF&var-repogroup_name=All) -> 98
    * Number of stars: 819 -> 3372
    * Number of forks: 43 -> 356

## Maintainers

We now have four maintainers from Second State, the same as the wasmedge initially. But we have established the [community governance guideline](https://github.com/WasmEdge/WasmEdge/blob/master/docs/GOVERNANCE.md) and will promote more contributors to be maintainers.

* Michael Yuan (Second State) @juntao
* Hung-Ying Tai (Second State) @hydai
* Yi-Ying He (Second State) @q82419
* Shen-Ta Hsieh (Second State) @ibmibmibm

## Project progress and adoption

The WasmEdge community has released five versions since joining CNCF. 

With the [version 0.8.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.0), the [version 0.8.1](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.1), the [version 0.8.2](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.2), the [version 0.9.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.9.0), and the [version 0.10.0](https://github.com/WasmEdge/WasmEdge/releases/tag/0.10.0), WasmEdge delivers important features that are important for cloud-native users.

* Networking socket. WasmEdge provides a Rust API and a JavaScript API for non-blocking network I/O, which is a crucial feature for running WasmEdge applications as cloud services.
* Container and K8s tooling. WasmEdge is an OCI-compliant runtime, so it could be managed and orchestrated by K8s and K8s tooling, like crun, containerd, CRI-O, K8s, Kind, OpenYurt, KubeEdge, SuperEdge.
* JavaScript support. WasmEdge now supports the full JavaScript language spec as well as common JavaScript module systems, including ES6, CommonJS (CJS), and NPM modules.
* Support for WebAssembly feature extensions such as WASI-NN and WASI-Crypto.
* Support for optional WebAssembly specs and proposals from the community, including the SIMD proposal, the Tail Call proposal, the extended-const proposal, the multi-memories proposal, the Import/Export mutable globals proposal, the Non-trapping float-to-int conversions proposal, the Sign-extension operators proposal, and the Multi-value proposal.
* Improved cross-platform compatibility. WasmEdge added support for macOS, Windows, Andriod, seL4, and OpenHarmony operating systems, as well as ARM32, ARM64, Apple M1 and M2 hardware.
* Enhanced host SDKs in C and Go languages.
* New host SDKs in Rust, Python, and Java languages.
* A new host / Wasm data binding toolkit called the `wasmedge-bindgen`.
* A new plugin system that makes WasmEdge native host functions and extensions easier to develop and install.
* A lot of performance improvements.

The WasmEdge community has expanded its collaboration with many open-source projects in the cloud-native space. Some of these projects now have production customers using WasmEdge.

* Suborbital, an application framework for Wasm, supports WasmEdge as an underlying WebAssembly runtime engine.
* The Proxy-wasm project added support for WasmEdge, enabling Wasmedge apps to run as Envoy proxy extensions.
* crun, the OCI Container Runtime, supports WasmEdge as a container runtime.
* YoMo, a streaming data processing framework, integrates WasmEdge as the runtime for embedded functions.
* The Nebula Graph and TiDB communities used WasmEdge to power UDFs (User Defined Functions) in database queries.
* KubeEdge (CNCF project), SuperEdge (CNCF project), and OpenYurt (CNCF project), all support WasmEdge apps on edge networks.
* Dapr (CNCF project), a runtime framework for microservices, demonstrated its capability to run WasmEdge programs as sidecar-based microservices.
* nginx-wasm-module, created by the APISIX team, enables WasmEdge applications to handle and respond to web requests within Nginx servers.
* Leading blockchain projects such as Polkadot / Substrate and FileCoin, are working on using WasmEdge as the underlying runtime to power their node software and smart contract execution engines.
* The developer community has also created demonstrations on running WasmEdge-based serverless functions in cloud providers including AWS Lambda, Vercel, Netlify, and Tencent Cloud.

Furthermore, the WasmEdge project is used by numerous commercial companies including VMWare, FutureWei, Huawei, ByteDance, Xunlei, and Ant Financial.

## Project goals

This is the WasmEdge project's first Annual review after joining the CNCF Sandbox.

For the immediate future, WasmEdge will focus on the following areas.

* Technical goals
    * Support more WebAssembly proposals and specs. In particular, we will provide full support for the WebAssembly component model and `wit-bindgen`.
    * Support TLS and HTTPS connections from WasmEdge apps.
    * Adopt popular application libraries and frameworks for the WasmEdge networking socket API. Examples including Rust's Tokio framework, database drivers, and client SDKs for common web services.
    * Support the entire Node.js API and NPM module ecosystem in WasmEdge's JavaScript runtime.
    * Expand our support for edge AI frameworks and models from Tensorflow, Pytorch, Paddlepaddle, from WasmEdge's AI inference extensions.
    * Improve WasmEdge support in popular application frameworks such as Dapr, Suborbital, Fermyon, essa-rs, etc.
    * Better support for host SDKs in Rust, Go, Java and Python for embedding WasmEdge in other applications.
    * Support edge computing platforms such as RISC-V and OpenWRT.
    * Debugging, monitoring, and observability tools for WasmEdge.

Our 2022 roadmap could be found [here](https://github.com/WasmEdge/WasmEdge/blob/master/docs/roadmap.md).

* User adoption goals
    * Drive end-user adoption through integrating more open-source projects, including popular databases and data analytics tools. 
    * Work closely with end-users to bring WasmEdge into production systems in at least 3-5 commercial companies.

* Community and community collaborations
    * Recruit and support proposal maintainers from companies other than Second State.
    * More community collaborations with open source projects in Linux Foundation and Apache Foundation.
    * Incorporate WasmEdge into leading Linux distributions as officially supported packages.

## How the CNCF can help to achieve the upcoming goals

In the past year, the WasmEdge project has received significant support from the CNCF. That includes LFX Mentorship opportunities, publication of tech articles on the CNCF blog, Project Pavilion and Office Hour opportunities at Kubecon, and legal support for the WasmEdge trademark in the US and China. Going forward, we plan to work closely with the CNCF on the following additional areas.

* Strategy to recruit and retain contributors.
* Guidance and support for the project's technical documentation.
* Opportunities to collaborate with other CNCF projects.

## Incubation readiness

WasmEdge aims to be incubation ready by the end of 2022.
