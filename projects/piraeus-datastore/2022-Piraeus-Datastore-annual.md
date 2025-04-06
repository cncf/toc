# Piraeus Datastore Sandbox Annual Review 2022

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

#### Comparison with OpenEBS Longhorn

OpenEBS is a CNCF Sandbox project, providing a fast and highly reliable container storage solution. OpenEBS runs
entirely in userspace, easing deployment and maintenance.

In comparison, Piraeus leans heavily on Linux kernel components, benefiting from the lower resource overhead as well
as faster execution in comparison to userspace solutions, at the cost of some portability.

## Highlights of Last Year
###  Community

Last year we saw the first major contributions from a community member:

* Automatically setting up secured communications between components with generated TLS certificates. [piraeus-operator#263](https://github.com/piraeusdatastore/piraeus-operator/pull/263)
* Making our scheduler plugin independent of the STORK project. [github project](https://github.com/piraeusdatastore/linstor-scheduler-extender/)

###  Releases

Piraeus Datastore saw 4 main releases in the last year, with the occasional necessary patch release:

* [v1.5.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v1.5.0) - May 12, 2021
* [v1.6.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v1.6.0) - Sep 02, 2021
* [v1.7.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v1.7.0) - Dec 14, 2021
* [v1.8.0](https://github.com/piraeusdatastore/piraeus-operator/releases/tag/v1.8.0) - Mar 15, 2022

## Annual Review Questions

1. _Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats._

    [Piraeus Datastore Devstats](https://piraeus.devstats.cncf.io/)

    Highlights:
    * [6 new contributors](https://piraeus.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)
    * [Commits](https://piraeus.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=d7&var-repogroups=All)
     have remained steady in the last year.

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

5. _How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)_

    Since this is the first annual review, there were no formally stated goals. Informally the initial goals since entering the Sandbox were:
    - *Tighten integration with native resources* - In the last year we've added:
      - Integration with Prometheus monitoring.
      - CSI storage capacity tracking.
      - Zone aware volume scheduling out of the box, using existing node labels.
      - Uploading snapshots to object storage based on VolumeSnapshotClass.
    - *Ease initial deployment hurdle* - While we improved the initial set up experience, especially around securing
      endpoints with TLS, we are still not quite were we want to be in terms of running a one command set-up.
    - *Continue to expand and broaden the Piraeus Datastore community.* - While we saw increased participation,
      especially in reporting and discussing issues, we still would like to see more external contributors.

6. _What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?_

    1. *Upgrade our operator* - current operator version lags behind the Operator SDK, making it harder to maintain.
    2. *Upgrade our documentation* - While most things are documented, the docs are not well organized.
    3. *Increase community activity* - we would like to see more adoption and contributions. We believe this is mainly
       an issue of outreach and visibility. 

7. _How can the CNCF help you achieve your upcoming goals?_

   CNCF has been helpful to Piraeus Datastore. Providing opportunities to increase visibility of Piraeus Datastore.
   Currently, the burden is on the Piraeus maintainers to make use of these opportunities. 

8. _Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)?_

   Piraeus does not currently meet all criteria for incubation. We would like to see increased adoption and less
   reliance on the original maintainers before moving towards incubation.

## Project Links

* [Website](https://piraeus.io/)
* [Github](https://github.com/piraeusdatastore/)
