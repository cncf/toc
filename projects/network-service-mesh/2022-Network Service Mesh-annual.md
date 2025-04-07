# Network Service Mesh Sandbox Annual Review 2022

## Table of Contents
* [Background](#background)
* [Alignment with Cloud Native](#alignment-with-cloud-native)
* [Highlights of Last Year](#highlights-of-last-year)
* [Annual Review Questions](#annual-review-questions)
* [Project Links](#project-links)

## Background
Network Service Mesh provides a Hybrid/Multi-cloud IP Service mesh. Inspired by L7 service meshes like Linkerd, Istio, 
and others, Network Service Mesh maps the concept of a service mesh from L7 payloads to L2/L3 payloads.

Network Service Mesh recognizes and respects that the existing Kubernetes Networking provides excellent service for 
common developer use cases.

Network Service Mesh requires no changes to the CNI plugin being used, or to Kubernetes itself to function.  
It runs as an additional layer of infrastructure on top of stock out of the box Kubernetes.

When Network Service Mesh is installed on a K8s cluster, sophisticated network connectivity is simplified for the developer. 
They need only specify, in an annotation on their Pod, the name of the Network Service they wish to consume, 
such as "db-replication".  Network Service Mesh will ensure 
that Network Service is available to their Pods (not to the cluster as a whole) without using CNI.  Standard 
Kubernetes Networking is still provided by CNI.  Network Service Mesh takes care to insure that it is completely 
non-conflicting with Kubernetes Networking.

### Example: **db-replication**

Database workloads across multiple K8s clusters in different public clouds and on prem can 
connect to a vL3 Network Service that exists solely for the purpose of database replication, and only has
database workloads that need to participate in such replication connected to it.  This is done declaratively.  The developer
simply adds a simple one line annotation to the Pod spec for their DB Pods.

## Alignment with Cloud Native

### Loose Coupling
IP Networking has traditionally been strongly coupled to the runtime domain (K8s cluster, VIM, datacenter, etc).
Where a workload ran determined what networking was available to it and the semantics of what networking features
it could receive.

Network Service Mesh allows workloads to be 'loosely coupled' to the Network Service (or Network Services) that are
relevant to their particular work independent of the environment they are running in.  Workloads running in different
clusters hybrid/multi-cloud scenarios can connect to a common Network Service that provides the networking features
needed for their particular collaboration.

Network Service Mesh does this without disturbing (and taking care not to break) the existing networking provided
by the runtime.  K8s Networking via CNI is untouched, and continues to work as expected.

Network Service Mesh makes IP networking itself Cloud Native by loosely coupling with the existing 
immutable infrastructure in K8s to deliver whatever non-standard Networking needs a workload has with minimal toil.

NSM is a complement to higher level L7 Service Meshes (Linkerd), providing an additional 
connectivity,security and observability at lower layers.

### Minimal Toil
The Network Service Mesh enabled developers access Network Services with minimal toil:

* Developers ask for what Network Service they want by name
* Developers can add metadata to their request for a Network Service with labels.
* Developers simply get the Network Service they asked for

Developers never have to know or think about 35 year old legacy Networking concepts like:

* Subnets
* Interfaces
* Routes
* IPAM

### Immutable Infrastructure

Network Service Mesh shifts IP networking from being 'Infrastructure' to being a selection of 'Network Services'.
This allows infrastructure to remain immutable while meeting a much wider variety of requirements.  
By scoping the selection of 'Network Services' to the granularity of workloads (rather than clusters), Network Service Mesh
allows different workloads to consume different Network Services that could be conflicting if forced into cluster level
infrastructure.

## Highlights of Last Year
###  Community
   In the last year Network Service Mesh has had contributions from:
   *  [30 contributors](https://networkservicemesh.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All)
   *  [9 Organizations](https://networkservicemesh.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=commits)
      Including:
       *  Cisco
       *  Ericsson Software Technology
       *  Red Hat

###  Releases
   
Network Service Mesh had three release in the last year:
* [v1.0.0](https://networkservicemesh.io/docs/releases/v1.0.0/) - Jun 19, 2021
* [v1.1.0](https://networkservicemesh.io/docs/releases/v1.1.0/) - Nov 28, 2021
* [v1.2.0](https://networkservicemesh.io/docs/releases/v1.2.0/) - Feb 15, 2021

## Annual Review Questions

1.  _Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats._

    [Network Service Mesh Devstats](https://networkservicemesh.devstats.cncf.io/)
    
    Highlights:
    *  [30 contributors](https://networkservicemesh.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All)
    *  [9 Organizations](https://networkservicemesh.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=commits)
       
       Including:
       *  Cisco
       *  Ericsson Software Technology
       *  Red Hat
    * [Commits](https://networkservicemesh.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=d7&var-repogroups=All)
      have remained strong in the last year.
    
2.  _How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)_

    Network Service Mesh has five top level MAINTAINERS:
    * Andre Sobolov (Xored)
    * Ed Warnicke (Cisco)
    * Frederick Kautz (Anthem)
    * Nikolay Nikolaev (Juniper)
    * Denis Tingaikin (Xored) - **New since last review**
    
    In addition, as the Network Service Mesh has grown, it has recently begun adding sub-MAINTAINERS to specific repos in the organization:
    * Alexandre Menezes (Red Hat) - [nsm-operator](https://github.com/networkservicemesh/nsm-operator)
    * Denis Tingajkin (Xored) - [fanout](https://github.com/networkservicemesh/fanout)
    * Dimitrios Markou (Ericsson Software Technology) - [sdk-ovs](https://github.com/networkservicemesh/sdk-ovs), [cmd-forwarder-ovs](https://github.com/networkservicemesh/cmd-forwarder-ovs) - **New since last review**
    * Periyasamy Palanisamy (Ericsson Software Technology) - [sdk-ovs](https://github.com/networkservicemesh/sdk-ovs), [cmd-forwarder-ovs](https://github.com/networkservicemesh/cmd-forwarder-ovs) - **New since last review**
    * Laszlo Kiraly (Ericsson Software Technology) - [cmd-nse-remote-vlan](https://github.com/networkservicemesh/cmd-nse-remote-vlan) - **New since last review**
    
3.  _What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)_

    At least four companies are beginning to build this into next generation architecture:
    * [Anthem](https://anthem.ai/)
    * [Ericsson](https://www.ericsson.com/en)
    * [Intel](https://www.intel.com/)
    * [Cisco](https://www.cisco.com/)

4. _How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)_
   Goals from last time:
      - *v1.0.0 release.* - Done. The refactor of NSM took longer than expected, but NSM has now settled into a regular 60 day release cadence.
      - *Continue to expand and broaden the NSM community.* - NSM has increased the breath of maintainers overall, particularly in sub-maintainers for particular repos
      - *Grow adoption across: Service Provider* - “Ericsson is actively contributing to NSM to enable 5G specific use cases for cloud native network functions. We have multiple NSM based solutions in our roadmap, targeting live deployments by the end of 2022.”
      - *Grow adoption across: Edge/IoT* - "NSM is being considered in the Intel Smart Edge Open roadmap for Service Function Chaining (SFC) of FD.io VPP based Container Network Functions (CNFs)"
      - *Grow adoption across: Enterprise* - The NSM vL3 and Application Service Mesh over vL3 (see below) features need to be delivered to enable traction in Enterprise.

5. _What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?_
    1. Maintaining a regular (currently 60 day)  release cadence.  Next scheduled release is [NSM v1.3.0](https://networkservicemesh.io/docs/roadmap/v1.3.0/)
    2. Build out Enterprise Network Services:
      1. vL3 - Virtual L3 - a flat IP networking domain to which workloads from different clusters/clouds/etc can connect and over which they can communicate
      2. Application Service Mesh (L7) over vL3 (Note: This goal involves *running* an existing service mesh like Linkerd/Istio/OpenServiceMesh/Kuma over a Network Service Mesh provided vL3.  NSM will *not* be writing its own L7 functionality)
     
6. _How can the CNCF help you achieve your upcoming goals?_
   
   CNCF has been incredibly helpful to Network Service Mesh.  All of our requests have been either been met
   enthusiastically, or we have been informed cannot be provided until we reach incubation or graduation as a matter of
   policy set by the Board and/or TOC.

7. _Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?_

   Network Service Mesh does not currently meet all criteria for incubation, but is on track to do so.
   
   Network Service Mesh believes that it meets or can trivially meet all criteria for incubation *except* adoption.  
   There is interest from adopters in 2022, we hope to make progress on adoption this year.

## Project Links

* [Website](https://networkservicemesh.io/)
* [Github](https://github.com/networkservicemesh/)
* [Weekly Meeting Minutes](https://docs.google.com/document/d/1C9NKjo0PWNWypROEO9-Y6haw5h9Xmurvl14SXpciz2Y/edit#heading=h.rc9df0a6n3ng)
* [Twitter](https://twitter.com/nservicemesh)
* [Linkedin](https://www.linkedin.com/company/networkservicemesh)
* [Youtube Channel](https://www.youtube.com/channel/UCXB3ccAbw9qLXSwDNzrkWhA)
* [Network Service Mesh Sandbox Proposal](https://github.com/cncf/toc/pull/212)
* [Network Service Mesh TOC Presentation Slides](https://docs.google.com/presentation/d/1bijEpuwaaa6jR1D5PAjyW731-j6Xc1TFHJuUh_FwwK8/edit#slide=id.g54dcbc5536_7_215)
