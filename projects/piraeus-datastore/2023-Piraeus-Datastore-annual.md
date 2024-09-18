# Piraeus Datastore Sandbox Annual Review 2023

## Background
Piraeus Datastore is a cloud-native storage system that aims to make node-local storage available to Kubernetes in
a fast, reliable and highly-available way. Piraeus is built on tried and tested Linux storage systems in order to
achieve the goals of reliability as well as making it universally usable.

At the heart of the Piraeus Datastore project sits the [Piraeus Operator](https://github.com/piraeusdatastore/piraeus-operator)
which sets up the whole storage system for the user, as well as manage day 2 operations. This includes:

* Deployment of the open-source replication component ([DRBD](https://github.com/linbit/drbd)).
* Deployment of the open-source storage management component ([LINSTOR](https://github.com/linbit/linstor-server)).
* Deployment of the Piraeus CSI driver components.
* Deployment of the optional high-availability controller for faster fail-over of stateful workloads and storage
  topology aware scheduler. 
* Configuration of node local storage to be made available in the cluster.
* Configuration of Prometheus metrics scraping.

### Comparison to similar projects

Piraeus Datastore can be compared to other CNCF projects that also offer cloud-native storage.

#### Comparison with rook

Rook is an open source cloud-native storage orchestrator, currently focused on provisioning CephFS clusters, as well as
NFS volumes and Cassandra databases.

In comparison, Piraeus focuses not just on not just operating storage itself, but also enhancing Kubernetes operations
where possible. For example, Piraeus offers enhanced storage locality with a dedicated scheduler plugin, and faster
fail-over of stateful workloads using a dedicated Kubernetes controller.

#### Comparison with Longhorn

Longhorn is a distributed block storage system for Kubernetes. Longhorn implements distributed block storage using
containers and microservices. Longhorn creates a dedicated storage controller for each block device volume and
synchronously replicates the volume across multiple replicas stored on multiple nodes.

In comparison, Piraeus only uses a fixed number of containers, leveraging Linux kernel components to achieve replication
across the network. As a consequence, replication will continue to work even in the face of disruptions to the container
runtime.

#### Comparison with OpenEBS

OpenEBS is a CNCF Sandbox project, providing a fast and highly reliable container storage solution. OpenEBS runs
entirely in userspace, easing deployment and maintenance.

In comparison, Piraeus leans heavily on Linux kernel components, benefiting from the lower resource overhead as well
as faster execution in comparison to userspace solutions, at the cost of some portability.

## Highlights of Last Year
###  Community

Last year we saw some major contributions from community members:

* Improved handling of systems with SELinux enforcement enabled. [PR](https://github.com/piraeusdatastore/piraeus-operator/pull/477)
* Documentation on how to Piraeus Datastore on different Kubernetes distributions. [Documentation for Talos Linux](https://www.talos.dev/v1.4/kubernetes-guides/configuration/storage/#piraeus--linstor)

###  Releases

Piraeus Datastore saw a number of releases in the last year, notably the release of the long worked on v2.0.0, along
with minor releases to fix some bugs. The notable releases were:

* [v1.9.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v1.9.0) - Jul 27, 2022
* [v1.10.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v1.10.0) - Oct 19, 2022
* [v2.0.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v2.0.0) - Feb 23, 2023
* [v2.1.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v2.1.0) - Apr 24, 2023

## Annual Review Questions

1. _Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats._

    [Piraeus Datastore Devstats](https://piraeus.devstats.cncf.io/)

    Highlights:
    * [15 new contributors](https://piraeus.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)
    * [Commits](https://piraeus.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=d7&var-repogroups=All)
     have increased in the last year.

2. _How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)_

    Piraeus Datastore has four [MAINTAINERS](https://github.com/piraeusdatastore/piraeus/blob/master/MAINTAINERS.md):

   | Name | GithubID | Affiliation |
   |------|----------|-------------|
   | Alex Zheng | [alexzhc](https://github.com/alexzhc) | DaoCloud |
   | Joel Colledge | [JoelColledge](https://github.com/JoelColledge) | LINBIT |
   | Moritz Wanzenböck | [WanzenBug](https://github.com/WanzenBug) | LINBIT |
   | Roland Kammerer | [rck](https://github.com/rck) | LINBIT |

3. _What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)_

    Piraeus Datastore does not directly track adoption. Some known adopters are listed here:
   * [flant.com](https://flant.com/) in their [deckhouse platform](https://deckhouse.io/en/)
   * [phoenix-media.eu](https://www.phoenix-media.eu)
   * [Fateflow Federated AI](https://github.com/rafflescity/rafflescity/blob/main/adoptions/federated_ai.md)
   * [Talos Linux lists Piraeus as supported storage option](https://www.talos.dev/v1.4/kubernetes-guides/configuration/storage/#piraeus--linstor)

5. _How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)_

   > *Upgrade our operator* - current operator version lags behind the Operator SDK, making it harder to maintain.

      This goal was achieved with release of Operator v2.0.0, using the latest Operator SDK. It has been updated
      regularly since.

   > *Upgrade our documentation* - While most things are documented, the docs are not well organized.

      With the release of Operator v2.0.0, we updated our documentation to have a clear structure. The new documentation
      has tutorials and how-to guides which were missing in the old version.
      [Docs](https://github.com/piraeusdatastore/piraeus-operator/tree/v2/docs)

   > *Increase community activity* - we would like to see more adoption and contributions. We believe this is mainly
         an issue of outreach and visibility.

      With the release of Operator v2.0.0, we saw renewed activity in the community.

6. _What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?_

   1. While Operator v2.0.0 improved our documentation, there is still places that have not been updated to the new
      release. We would fill these gaps.
   2. We currently have no documented upgrade path from Operator v1 to v2. We would like to tackle this problem in the
      coming year.
   3. We would like to improve observability of Piraeus Datastore, specifically monitoring and alerting using
      Prometheus. We feel this is an important part of any reliable storage system, where we can certainly improve.

7. _How can the CNCF help you achieve your upcoming goals?_

   CNCF has been helpful to Piraeus Datastore. Providing opportunities to increase visibility of Piraeus Datastore.
   Currently, the burden is on the Piraeus maintainers to make use of these opportunities. 

8. _Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)?_

   Piraeus does not currently meet all criteria for incubation. We would like to see increased adoption and less
   reliance on the original maintainers before moving towards incubation.

## Project Links

* [Website](https://piraeus.io/)
* [Github](https://github.com/piraeusdatastore/)
