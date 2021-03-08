# Service Mesh Interface Annual Review 2021

## Background

[Service Mesh Interface](https://smi-spec.io) is a specification for service meshes that run on Kubernetes. It defines a common standard that can be implemented by a variety of providers. This allows for both standardization for end-users and innovation by providers of Service Mesh technology. SMI enables flexibility and interoperability, and covers the most common service mesh capabilities.

The [Service Mesh Interface](https://github.com/servicemeshinterface) GitHub organization not only hosts the main [Service Mesh Interface Specification](https://github.com/servicemeshinterface/smi-spec) repo, but also related repos for the website and adapters.

- [smi-spec.io](https://github.com/servicemeshinterface/smi-spec.io) - project website
- [SMI Metrics API](https://github.com/servicemeshinterface/smi-metrics) - an implementation of the Traffic Metrics Spec
- [SMI SDK for go](https://github.com/servicemeshinterface/smi-sdk-go) - Kubernetes client implementation for SMI
- [SMI adapter for Istio](https://github.com/servicemeshinterface/smi-adapter-istio) - Kubernetes operator which implements the Service Mesh Interface(SMI) Traffic Split, Traffic Access Control, and Traffic Specs APIs to work with Istio

**SMI was accepted as a CNCF Sandbox project in April 2020**
- [SMI Sandbox Proposal](https://github.com/cncf/toc/blob/main/proposals/sandbox/smi.adoc)
- [SMI Sandbox entry blog post](https://smi-spec.io/blog/smi-joins-cncf/)

## Year in Review

- CNCF Sandbox project acceptance and onboarding (including GitHub updates and CNCF tooling)
- Spec evolution - [with current status viewable on the repo](https://github.com/servicemeshinterface/smi-spec#service-mesh-interface-documents)
  - With a release in early 2020 right before entering the CNCF and three releases since then, SMI has continued to progress: https://github.com/servicemeshinterface/smi-spec/releases
    - In [v0.4.0](https://github.com/servicemeshinterface/smi-spec/releases/tag/v0.4.0) we updated Traffic Specs, Traffic Split, Traffic Target, and Traffic Metrics.
    - In [v0.5.0](https://github.com/servicemeshinterface/smi-spec/releases/tag/v0.5.0), we added top level spec field to Traffic Specs and Traffic Target
    - In [v0.6.0](https://github.com/servicemeshinterface/smi-spec/releases/tag/v0.6.0), we introduced updates to Traffic Specs, Traffic Access, Traffic Metrics, and Traffic Split
- Ongoing involvement with the community: CNCF slack and an open community call every other week.
  - SMI blog with [community contribution](https://smi-spec.io/blog/building-an-alternative-service-mesh/)
  - New community members joining calls regularly: [communications links](https://github.com/servicemeshinterface/smi-spec#communications)
  - [SMI conformance testing project](https://docs.google.com/document/d/1HL8Sk7NSLLj-9PRqoHYVIGyU6fZxUQFotrxbmfFtjwc/edit#heading=h.rkzc80m1a5sr)
- Presentations including:
  - [Introduction to Service Mesh Interface](https://www.cncf.io/webinars/introduction-to-service-mesh-interface/)
  - [Navigating the service mesh ecosystem](https://www.cncf.io/webinars/navigating-the-service-mesh-ecosystem/)
  - [Multi Cluster and Multi Mesh Patterns](https://sched.co/abc9)
  - [Service Mesh on Arrested DevOps](https://www.arresteddevops.com/service-mesh/)
  - [Fireside Chat: SMI and the Future of Service Mesh](https://www.youtube.com/watch?v=g9C4-i5e8kw&list=PLRQaLsJQ82CJKzI148URGHWvXlkCq83qA&index=11&t=0s)
  - [Edward ServiceMeshHands](https://www.softwarecircus.io/software-circus-nightmares-on-cloud-street-schedule-videos)

## Annual Review Items

- *Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.*
  - Year in review for the main [SMI devstats dashboard](https://smi.devstats.cncf.io/d/8/dashboards?orgId=1&from=now-365d&to=now-1h&refresh=1d)
    - This shows the greater activity in spring 2020 around CNCF entry time, then ongoing activity overall. While devstats can show releases, repo activity is not the only metric for a spec. It’s useful to consider nuance in the content of the updates (specifically, how it’s been clarified due to direct input from implementers.)
  - Participation from [various organizations](https://smi.devstats.cncf.io/d/5/companies-table?orgId=1)


- *How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)*
    - SMI currently has [11 core maintainers](https://github.com/servicemeshinterface/smi-spec/blob/main/CODEOWNERS)
        - Thomas Rampelberg, Buoyant
        - Lachlan Evenson, Microsoft
        - Lee Calcote, Layer5
        - Nic Jackson, HashiCorp
        - Jason Hansen, Microsoft
        - Stefan Prodan, Weaveworks
        - Michelle Noorali, Microsoft
        - Bridget Kromhout, Microsoft
        - Idit Levine, Solo.io
        - Michael Hausenblas, Amazon Web Services
        - Tarun Pothulapati, Buoyant

- *What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)*

    - Existing implementations of the spec when SMI joined Sandbox (some have evolved/renamed):
        - [Consul Connect](https://www.consul.io/docs/connect): service segmentation (via adapter)
        - [Flagger](https://flagger.app/): progressive delivery operator
        - [Gloo Mesh](https://www.solo.io/products/gloo-mesh): multi-cluster service mesh management plane
        - [Istio](https://istio.io/): connect, secure, control, observe (via adapter)
        - [Linkerd](https://linkerd.io/): ultralight service mesh
        - [Traefik Mesh](https://traefik.io/traefik-mesh/): simpler service mesh
        - [Rio](https://rio.io/): application deployment engine

    - New implementations of the spec this year:
        - [Meshery](https://layer5.io/meshery): the service mesh management plane
        - [Open Service Mesh](https://openservicemesh.io/): lightweight and extensible cloud native service mesh
        - [Argo Rollouts](https://argoproj.io/): advanced deployment & progressive delivery controller
        - [NGINX Service Mesh](https://www.nginx.com/products/nginx-service-mesh): lightweight and seamless
  - Ecosystem overview: https://smi-spec.io/#ecosystem


- *How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)*
    - This is SMI's first annual review. Project goals thus far have included:
        - Maintain open pathways for community interaction
        - Maintain progress on API updates by moving initial APIs to stable and identifying other APIs for inclusion in SMI.
   - SMI has progressed towards those goals this year.

- *What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?*
  - Evolve the APIs through alpha to beta and finally stable
  - Update adapters for ongoing compatibility
  - Look into introducing a new SMI controller SDK to make it easier for implementations to adopt SMI
  - Increase participation in community by supporting ecosystem tooling integrations (Argo, Flagger, etc.)

- *How can the CNCF help you achieve your upcoming goals?*
    - The SMI project has benefited from CNCF webinars and events to increase awareness.
    - SMI has also used the opportunity of project office hours during KubeCon to reach out to the community.

- *Do you think that your project meets the criteria for incubation?*
    - We plan to file for incubation soon, yes. While some details appear to be focused on suitable metrics for source code, not necessarily specifications, we think we can meet the criteria for incubating projects.
