# Network Service Mesh Sandbox Annual Review 2020

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

When Network Service Mesh is installed on a K8s cluster, network connectivity is simplified for the developer. 
They need only specify, in an annotation on their Pod, the name of the Network Service they wish to consume, 
such as "db-replication".Network Service Mesh will ensure 
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
   *  [31 contributors](https://networkservicemesh.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All) - up from 18 contributors when NSM was accepted as a CNCF Sandbox project
   *  [18 Organizations](https://networkservicemesh.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=commits)
    
      Including:
      *  Cisco
      *  Ericsson
      *  Intel
      *  Red Hat
      *  SUSE
      *  VMWare

###  Releases
   
Network Service Mesh had [two releases](https://networkservicemesh.io/docs/releases/b/) in the last year:
* [Andromeda](https://networkservicemesh.io/docs/releases/a/) - May 14, 2019
* [Borealis](https://networkservicemesh.io/docs/releases/b/) - Nov 27, 2019

###  Events
* [NSMCon 2019 NA](https://networkservicemesh.io/events/nsmcon2019/) - A full day colocated event at Kubecon NA 2019.
  Funded by sponsorship from:
  * Cisco
  * doc.ai
  * VMware
  * Juniper
* [NSMCon 2020 EU](https://networkservicemesh.io/events/nsmconeu2020/) - Upcoming full day 'colocated' event at Kubecon EU 2020.
  Funded by sponsorship from:
  * Cisco
  * Anthem

## Annual Review Questions

1.  _Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats._

    [Network Service Mesh Devstats](https://networkservicemesh.devstats.cncf.io/)
    
    Highlights:
    *  [31 contributors](https://networkservicemesh.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All) - up from 18 at acceptance into CNCF
    *  [18 Organizations](https://networkservicemesh.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=commits)
       
       Including:
       *  Cisco
       *  Ericsson
       *  Intel
       *  Red Hat
       *  SUSE
       *  VMWare
    * [Commits](https://networkservicemesh.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=d7&var-repogroups=All)
      have accelerated since the beginning of our migration to a more modular multi-repo model.
    
2.  _How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)_

    Network Service Mesh has four top level MAINTAINERS:
    * Andre Sobolov (Xored)
    * Ed Warnicke (Cisco)
    * Frederick Kautz (Doc.ai)
    * Nikolay Nikolaev (Kong)
    
    Of note: Two of those maintainers have remained engaged through employer changes
    * Frederick Kautz: Red Hat -> doc.ai
    * Nikolay Nikolaev -> VMWare -> Kong
    
    In addition, as the Network Service Mesh has grown, it has recently begun adding sub-MAINTAINERS to specific repos in the organization:
    * Alexandre Menezes (Red Hat) - [nsm-operator](https://github.com/networkservicemesh/nsm-operator)
    * Denis Tingajkin (Xored) - [fanout](https://github.com/networkservicemesh/fanout)
    
3.  _What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)_

    At least three companies are beginning to build this into next generation architecture:
    * [Anthem](https://anthem.ai/)
    * [Bell Canada](https://www.bell.ca/)
    * [Doc.ai](https://doc.ai/)
 
    The project is also used in the [CNF Testbed](https://github.com/cncf/cnf-testbed).

4. _How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)_

   This is Network Service Mesh's first annual review, so we have no previously recorded goals from past annual reviews.

5. _What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?_
    1. v1.0.0 release. 
       
       Several key beach-head customers are waiting for v1.0.0 to begin deployment.
    1. Continue to expand and broaden the NSM community.
    1. Grow adoption across:
       * Enterprise
       * Service Provider
       * IoT
     
6. _How can the CNCF help you achieve your upcoming goals?_
   
   CNCF has been incredibly helpful to Network Service Mesh.  All of our requests have been either been met 
   enthusiastically, or we have been informed cannot be provided until we reach incubation or graduation as a matter of
   policy set by the Board and/or TOC.

7. _Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?_

   Network Service Mesh does not currently meet all criteria for incubation, but is on track to do so.
   
   Network Service Mesh believes that it meets all criteria for incubation *except* adoption.  
   There is interest from potential adopters sufficient to meet the adoption criteria as the 
   technology matures.  

## Project Links

* [Website](https://networkservicemesh.io/)
* [Github](https://github.com/networkservicemesh/)
* [Weekly Meeting Minutes](https://docs.google.com/document/d/1C9NKjo0PWNWypROEO9-Y6haw5h9Xmurvl14SXpciz2Y/edit#heading=h.rc9df0a6n3ng)
* [Twitter](https://twitter.com/nservicemesh)
* [Linkedin](https://www.linkedin.com/company/networkservicemesh)
* [Youtube Channel](https://www.youtube.com/channel/UCXB3ccAbw9qLXSwDNzrkWhA)
* [Network Service Mesh Sandbox Proposal](https://github.com/cncf/toc/pull/212)
* [Network Service Mesh TOC Presentation Slides](https://docs.google.com/presentation/d/1bijEpuwaaa6jR1D5PAjyW731-j6Xc1TFHJuUh_FwwK8/edit#slide=id.g54dcbc5536_7_215)

