# Service Mesh Interface Annual Review 2022

## Background

[Service Mesh Interface](https://smi-spec.io) is a specification for service meshes that run on Kubernetes. It defines a common standard that can be implemented by a variety of providers. This allows for both standardization for end-users and innovation by providers of Service Mesh technology. SMI enables flexibility and interoperability, and covers the most common service mesh capabilities.

The [Service Mesh Interface](https://github.com/servicemeshinterface) GitHub organization not only hosts the main [Service Mesh Interface Specification](https://github.com/servicemeshinterface/smi-spec) repo, but also related repos for the website and adapters:

- [smi-spec.io](https://github.com/servicemeshinterface/smi-spec.io) - project website
- [SMI Metrics API](https://github.com/servicemeshinterface/smi-metrics) - an implementation of the Traffic Metrics Spec
- [SMI SDK for go](https://github.com/servicemeshinterface/smi-sdk-go) - Kubernetes client implementation for SMI
- [SMI adapter for Istio](https://github.com/servicemeshinterface/smi-adapter-istio) - Kubernetes operator which implements the Service Mesh Interface(SMI) Traffic Split, Traffic Access Control, and Traffic Specs APIs to work with Istio
- [SMI Controller SDK](https://github.com/servicemeshinterface/smi-controller-sdk) - a successor to smi-sdk-go that provides a simple abstraction layer for making your mesh SMI-compliant
- [Istio SMI Controller](https://github.com/servicemeshinterface/istio-smi-controller) - a successor to the SMI istio adapter that uses the smi-controller-sdk

## Year in Review

- Spec stabilization and augmentation - [with current status viewable on the repo](https://github.com/servicemeshinterface/smi-spec#service-mesh-interface-documents)
  - 2022 was a year of maturation and stability for SMI. After releasing v0.6 of the specification, we turned our attention to ensuring that the spec had time to breathe (so to speak) and that the community had time to adopt the changes. This resulted in fruitful launches of collaborations such as [SMI Conformance Testing GitHub Action](https://github.com/layer5io/meshery-smi-conformance-action), and [Service Mesh Performance](https://smp-spec.io/). These efforts didn’t stop real progress in the spec itself, as issues and features continued to be discussed, triaged, and implemented via the bi-weekly community calls and various ad-hoc design documents. 
- Ongoing involvement with the community: CNCF slack and an open community call every other week.
  - SMI blog with [community contribution](https://smi-spec.io/blog/pipelining-service-mesh-specifications/)
  - New community members joining calls regularly: [communications links](https://github.com/servicemeshinterface/smi-spec#communications)

- [SMI Conformance Testing blog post](https://smi-spec.io/blog/validating-smi-conformance-with-meshery/)

## Annual Review Items

- *Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.*
  - Year in review for the main [SMI devstats dashboard](https://smi.devstats.cncf.io/d/8/dashboards?orgId=1&from=now-365d&to=now-1h&refresh=1d)
    - This shows a sustained, consistent level of contributions to the spec with changes based on discussion between maintainers and the community (with an expected dropoff around KubeCon events and the holidays). While devstats can show releases, repo activity is not the only metric for a specification. It’s useful to consider nuance in the content of the updates (specifically, how it’s been clarified due to direct input from implementers.)
  - Participation from [various organizations](https://smi.devstats.cncf.io/d/5/companies-table?orgId=1)

- *How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)*
  - SMI currently has [12 core maintainers](https://github.com/servicemeshinterface/smi-spec/blob/main/CODEOWNERS)
    - Thomas Rampelberg, Airbnb
    - Lachlan Evenson, Microsoft
    - Lee Calcote, Layer5
    - Nic Jackson, HashiCorp
    - Jason Hansen, Microsoft
    - Stefan Prodan, Weaveworks
    - Bridget Kromhout, Microsoft
    - Idit Levine, Solo.io
    - Michael Hausenblas, Amazon Web Services
    - Tarun Pothulapati, Buoyant
    - Keith Mattix, Microsoft
    - Jared Byers, NGINX Inc.

- *What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)*

  - Existing implementations of the spec at the beginning of the year (some have evolved/renamed):
    - [Consul Connect](https://www.consul.io/docs/connect): service segmentation (via adapter)
    - [Flagger](https://flagger.app/): progressive delivery operator
    - [Gloo Mesh](https://www.solo.io/products/gloo-mesh): multi-cluster service mesh management plane
    - [Istio](https://istio.io/): connect, secure, control, observe (via adapter)
    - [Linkerd](https://linkerd.io/): ultralight service mesh
    - [Traefik Mesh](https://traefik.io/traefik-mesh/): simpler service mesh
    - [Rio](https://rio.io/): application deployment engine
    - [Meshery](https://layer5.io/meshery): the service mesh management plane
    - [Open Service Mesh](https://openservicemesh.io/): lightweight and extensible cloud native service mesh
    - [Argo Rollouts](https://argoproj.io/): advanced deployment & progressive delivery controller
    - [NGINX Service Mesh](https://www.nginx.com/products/nginx-service-mesh): lightweight and seamless

  - Partnerships discussed around the spec this year
    - [Cilium mesh](https://cilium.io/blog/2021/12/01/cilium-service-mesh-beta): An eBPF powered service mesh
    - [Gateway API]https://gateway-api.sigs.k8s.io/): successor to Kubernetes Ingress
    - [Kubeflow](https://www.kubeflow.org/) - ML platform currently using Istio
    - [Kuma](https://kuma.io/) - Service mesh from Kong
  - Ecosystem overview: <https://smi-spec.io/#ecosystem>

- *How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)*

Last years goals were:

```markdown
- Evolve the APIs through alpha to beta and finally stable
  - Update adapters for ongoing compatibility
  - Look into introducing a new SMI controller SDK to make it easier for implementations to adopt SMI
  - Increase participation in community by supporting ecosystem tooling integrations (Argo, Flagger, etc.)
```
  
SMI has progressed in many of these goals this year. We’ve elected 2 new maintainers via organic community participation, and we’ve identified key areas where we can continue to foster adoption. Through our efforts in community outreach, we’ve been able to safely design updates to the specification through input from partners representing 6 different service meshes (Istio, Consul, Linkerd, Kuma, NGINX, and OSM). The new smi-controller-sdk was launched, allowing us to effectively deprecate the adapters in favor of a more maintainable solution with a lower barrier to entry.

- *What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?*
  - Invest heavily in partnerships with other meshes around Gateway API
  - Refocus efforts on creating APIs for policy enforcement rather than traffic routing
  - Continue to support our existing implementations and ecosystem tooling integrations
  - Release our first beta version

- *How can the CNCF help you achieve your upcoming goals?*
  - The SMI project has benefited from CNCF webinars and events to increase awareness.
  - SMI has also used the opportunity of project office hours during KubeCon to reach out to the community.

- *Do you think that your project meets the criteria for incubation?*
  - We plan to file for incubation soon, yes. While some details appear to be focused on suitable metrics for source code, not necessarily specifications, we think we can meet the criteria for incubating projects.
