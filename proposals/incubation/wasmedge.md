# **Propose WasmEdge into Incubation**

## **About WasmEdge**

WasmEdge is a lightweight, high-performance, OCI-compatible, and extensible WebAssembly runtime for cloud-native, edge, and decentralized applications. It powers serverless apps, embedded functions, microservices, smart contracts, and IoT devices.

### **Sandbox project acceptance**

WasmEdge was accepted as a Sandbox project in April 2021. See the [_onboarding issue_](https://github.com/cncf/toc/issues/646).

## **Progress since Sandbox**

Since joining CNCF, WasmEdge had great achievements, including developer community building and technical updates.

### **Community Metrics**

You can find WasmEdge’s devstats page and dashboards [_here_](https://wasmedge.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

Below are some highlights of WasmEdge’s community growth in the sandbox. 

* WasmEdge now has [_168 contributors_](https://wasmedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%202%20years&var-repogroup_name=All) including 14 maintainers from[_53 organizations_](https://wasmedge.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Since%20joining%20CNCF&var-metric=contributions)and [_22 countries and regions_](https://wasmedge.devstats.cncf.io/d/18/overall-project-statistics-table?orgId=1&viewPanel=1). There are [_343.92 commits_](https://wasmedge.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=1619712000000&to=now) per month, with [_44.32 PR_](https://wasmedge.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=1619712000000&to=now) merged from the day when WasmEdge joined the CNCF Sandbox project.
* The community has grown since WasmEdge entered the CNCF Sandbox.
    * We have held monthly meetings since October 2021. The meeting note can be found [_here_](https://docs.google.com/document/d/1iFlVl7R97Lze4RDykzElJGDjjWYDlkI8Rhf8g4dQ5Rk/edit), and the recording meeting video can be found [_here_](https://www.youtube.com/playlist?list=PL4H9HXvnb_jYbsHBOMk-jayEMGrl09VSO).
    * Number of maintainers: 4 → 14
    * Number of contributors: [_6_](https://wasmedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Before%20joining%20CNCF&var-repogroup_name=All) → [_176_](https://wasmedge.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Since%20joining%20CNCF&var-repogroup_name=All)
    * Number of stars: 819 → 6200
    * Number of forks: 43 → 558

### **Technical metrics**

The WasmEdge community has released ten new versions of the software since joining CNCF sandbox.

* [_version 0.8.0_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.0)
* [_version 0.8.1_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.1)
* [_version 0.8.2_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.2)
* [_version 0.9.0_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.9.0)
* [_version 0.10.0_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.10.0)
* [_version 0.11.0_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.11.0)
* [_version 0.11.1_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.11.1)
* [_version 0.11.2_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.11.2)
* [_version 0.12.0_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.12.0)
* [_version 0.12.1_](https://github.com/WasmEdge/WasmEdge/releases/tag/0.12.1)

With each release, WasmEdge delivers important features that enable new cloud native use cases for WasmEdge users and partners.

The list below highlights some of the technical achievements the team has accomplished since joining the CNCF sandbox. 

* Dynamic language support. Besides compiled languages such as Rust and C/C++, WasmEdge has added support for JavaScript, Python and PHP programs. WasmEdge supports the full JavaScript language spec as well as common JavaScript module systems, including ES6, CommonJS (CJS), and NPM modules. Python and PHP support is achieved through our collaboration with VMWare’s WasmLab. 
* Container and K8s tooling. WasmEdge has reached OCI-compliance. That allows it to be managed and orchestrated by standard container tools, such as Docker, crun, containerd, CRI-O, K8s, Kind, OpenYurt, KubeEdge, SuperEdge.
* Advanced networking with the WasmEdge sockets library. It supports non-blocking sockets, DNS, TLS. That allows WadmEdge to run modern HTTPS client and server applications. 
* Support for the Rust Tokio library. It allows many Tokio-based async Rust applications, especially async networking apps, to run seamlessly on WasmEdge.
* Connectors and drivers for commonly used microservice components, such as MySQL, PostgreSQL, Redis, Kafka, and Dapr. 
* Support AI inference applications written in both Rust and JavaScript using Tensorflow, OpenVino, PyTorch trained models.
* Host SDKs in Rust, Go, C, C++, Python, and Java languages.

WasmEdge is the only WebAssembly runtime in the market that could run complete microservices, with support for web servers, HTTPS clients, and database connections. Through partnerships with Docker, Red Hat, and Microsoft, WasmEdge could be seamlessly integrated into the existing cloud native ecosystem.

## **Incubation Stage Requirements**

***Document that it is being used successfully in production by at least three independent direct adopters which, in the TOC’s judgment, are of adequate quality and scope.***

The full list of companies and projects that have incorporated WasmEdge into their product offerings can be found [on our website](https://wasmedge.org/docs/contribute/users).

Below is a selected list of end-user organizations and products that are using WasmEdge for business workloads.

* [**_Docker_**](https://docker.com/): Bundles WasmEdge in Docker Desktop 4.15+ as a container runtime to support Wasm images across the entire suite of Docker developer tools. Details of this use case can be found [_here_](https://docs.docker.com/desktop/wasm/).
* [**_ByteDance_**](https://www.bytedance.com/en/): Uses WasmEdge to run custom logic in service mesh proxies and sidecars.
* [**_ByteTrade_**](https://www.bytetradelab.io/): Uses WasmEdge to run microservices for marketing automation.
* [**_XRPL Labs_**](https://github.com/XRPL-Labs): Use WasmEdge to execute smart contracts for the Ripple blockchain.
* [**_Huawei Cloud_**](https://www.huaweicloud.com/)**:** Uses WasmEdge to run Javascript-based serverless functions for internal business units. 
* [**_Flows Network_**](https://flows.network/)**:** Uses WasmEdge to run user-defined serverless functions for workflow automation.

***Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.***

There are 14 core team members for WasmEdge. Among them, 4 are maintainers, and 10 are committers/reviewers from 5 different organizations (Second State, ByteDance, SRA VJTI, Southeast University, and the University of Science and Technology of China).

Maintainers have admin access to the WasmEdge repo. committers/reviewers have write/triage access to the WasmEdge repo. See more differences [_here_](https://github.com/WasmEdge/WasmEdge/blob/master/docs/GOVERNANCE.md).

***Demonstrate a substantial ongoing flow of commits and merged contributions.***

A short summary has been provided in Community Metrics. Detailed statistics can be found in the following devstats links:

* [_Contribution Activity_](https://wasmedge.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-1y&to=now&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all): 719+ per month
* [_New contributors since sandbox_](https://wasmedge.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1619712000000&to=now)
* [_Issues open/closed since sandbox_](https://wasmedge.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1619712000000&to=now&var-period=m&var-repogroup_name=All): average 39.92 issues opened and 27.96 issues closed per month
* [_PRs open/closed since sandbox_](https://wasmedge.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1619712000000&to=now&var-period=m&var-repogroup_name=All): average 56.68 PRs opened and 53.36 PRs closed per month
* [_Contributing organizations_](https://wasmedge.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Since%20joining%20CNCF&var-metric=contributions): 53 companies and organizations

***A clear versioning scheme.***

WasmEdge has a clear path to release a new version by applying [_Semantic Versioning_](https://semver.org/). Before an official release, the new version will last three phases: Alpha, Beta, and RC. If everything is going well with the three pre-releases, then the official release will come. WasmEdge usually releases every three months. The latest release is v0.12.1.

The WasmEdge release process is documented at: [_https://wasmedge.org/docs/contribute/release_](https://wasmedge.org/book/en/contribute/release_process.html)

***Clearly documented security processes explaining how to report security issues to the project, and describing how the project provides updated releases or patches to resolve security vulnerabilities***

WasmEdge has achieved an Open Source Security Foundation (OpenSSF) best practices badge at the passing level. See details [_here_](https://bestpractices.coreinfrastructure.org/en/projects/5059).

WasmEdge has a responsible and mature vulnerability submission and disclosure program, managed by the WasmEdge Security Team. The process is documented at [_https://github.com/WasmEdge/WasmEdge/blob/master/SECURITY.md_](https://github.com/WasmEdge/WasmEdge/blob/master/SECURITY.md).

WasmEdge is [_in the process_](https://github.com/google/oss-fuzz/pull/8894) of joining Google’s OSS Fuzz program.

## **Statement on Alignment with the CNCF Mission**

WasmEdge aims to provide high-performance, lightweight, and secure container runtime in cloud-native and “edge” scenarios. Compared with other WebAssembly runtimes, WasmEdge is a fully featured and yet lightweight runtime with support for advanced networking, asynchronous functions, AI inference, and container tooling. WasmEdge could be seamlessly integrated with existing cloud-native ecosystems like Kubernetes and Docker.

With WasmEdge, users not only can achieve high performance without sacrificing security but also can use the same cloud-native tools and workflows. For this goal, WasmEdge has integrated with crun, youki, containerd’s runwasi (A CNCF project), and Docker Desktop. Hence, WasmEdge could be managed by Kubernetes (A CNCF project), SuperEdge (A CNCF project), OpenYurt (A CNCF project), and KubeEdge (A CNCF project).

WasmEdge aligns with the mission of CNCF in empowering organizations to bring the cloud-native and serverless application paradigms to “edge” scenarios. 

## **Future Plan**

## WasmEdge has the following technical goals:

* Support more WebAssembly proposals and specs, such as threading, stack switching, logging, and debugging. In particular, we aim to be one of the first Wasm runtimes to fully support the WebAssembly component model spec.
* Improve support for microservices. We will enable WasmEdge applications to access commonly used databases, messaging services, proxies, sidecars, and other services.
* Improve support for data streaming functions. We will work with data streaming partners to support embedded functions in WasmEdge in those frameworks. 
* Improve support for AI inference workloads in Wasm. We will continue adding support for popular models in computer vision, voice recognition, and document AI. We will also enable inference on LLMs such as Llama.
* Support AI training on Wasm. We are collaborating with ByteDance and the Ray community to support WasmEdge on Ray nodes. 
* Support the complete Node.js API and NPM module ecosystem in WasmEdge's JavaScript runtime.
* Improve WasmEdge support in popular application frameworks such as Dapr, Suborbital, Fermyon, essa-rs, etc.
* Better support for host SDKs in Rust, Go, Java, and Python for embedding WasmEdge in other applications.
* Debugging, monitoring, and observability tools for WasmEdge.
* Further optimize the developer experience and user experience for WasmEdge plugins. That would make it easy for the community to add new features to WasmEdge. 

On the community side, we will make a greater effort to recruit and support proposal maintainers from companies other than Second State.
