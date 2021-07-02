# CNCF Storage TAG Charter

Primary Author: Quinton Hoole

2019-04-23

Reviewed and contributed to by Saad Ali, Alex Chircop,
Xing Yang, Xiang Li, and other members of the CNCF Storage Working
Group and community.

# Introduction

This is the charter referred to in [CNCF
TAGs](https://github.com/cncf/toc/blob/main/tags/cncf-tags.md#tag-charter)
by the CNCF TOC.

# Areas Considered In Scope

Storage systems and approaches suitable for and commonly used in
[modern cloud-native
environments](https://github.com/cncf/toc/blob/main/DEFINITION.md) are considered in scope:

*   especially where these differ significantly from storage systems
    and approaches previously commonly used in traditional enterprise
    data center environments,
*   and which are not already adequately covered by other groups within the CNCF (see below).  
*   storage includes block stores, file systems, object stores, databases, key-value stores, and related caching mechanisms. 

We strive to understand the fundamental characteristics of different
storage approaches with respect to availability, scalability,
performance, durability, consistency, ease-of-use, cost and
operational complexity; and relate these to their suitability to
various cloud-native use cases. This is covered in more detail in the
[CNCF Storage Landscape White Paper](
http://bit.ly/cncf-storage-whitepaper).

# Areas Considered Out Of Scope

Anything not considered in scope above (:-)).  See also “Interfaces
with Related Groups” below.

Examples include:

*   Details of general-purpose low-level storage media like magnetic
    hard disks, solid-state storage devices, non-volatile memory, etc
    (which tend not to be very different irrespective of whether or
    not they’re being used in a cloud-native environment).
*   General authentication, authorization, accounting, auditing etc
    (even though these clearly apply to storage systems too) - because
    AAA etc is clearly the domain of the CNCF Security TAG.
*   Standardizing container storage interfaces - this is the domain of CSI.
*   Defining storage abstraction APIs for container orchestrators -
    these are the domains of the Storage TAGs or SIGs specific to each
    orchestrator, e.g. Kubernetes Storage SIG.

# TAG Mission Statement

To enable widespread and successful storage of persistent state in
cloud-native environments through:

1. Providing valuable and objective information to the TOC, End Users
   and Projects of the CNCF regarding areas considered in scope (see
   above).
2. Collaborating effectively with other related groups (see below).
3. Helping to maintain the continued health of CNCF Storage Projects
   (see below)
4. Identifying and filling gaps in the landscape of CNCF Storage
   Projects.

# Current CNCF Storage Projects

1. TiKV
2. etcd
3. Vitess
4. Rook
5. [TBD] OpenEBS

# Interfaces With Other Related Groups

*   **Kubernetes Storage SIG** - is focussed towards
    Kubernetes-specific storage abstractions, interfaces, and
    implementations of these interfaces.  We maintain close
    communication with this Kubernetes SIG, with several individuals
    actively involved in both.  Our aim is to avoid unnecessary
    duplication of effort by the two groups, and maintain clear an
    consistent messaging by the two groups to our end user community
    and projects.
*   **CSI** - is focussed on defining an industry standard “Container
    Storage Interface” (CSI) that will enable storage vendors to
    develop a plugin once and have it work across a number of
    container orchestration systems.  Again, we maintain close
    communication with this group, and avoid unnecessary duplication
    of effort and inconsistent messaging wherever possible.
*   **CNCF Security TAG** - works on the more general area of
    cloud-native security including authentication, authorization,
    encryption, accounting, auditing and related topics.  We defer as
    much as possible to this group to deal with general
    security-related issues, and liaise closely with them on how to
    deal with storage-specific security areas where these arise.
*   **CNCF App-Delivery TAG** - be focussed on the
    development, deployment, operation and testing of cloud-native
    applications.  We collaborate with this TAG where this pertains to
    Storage.
*   **K8s Apps SIG** - has done some work on how Kubernetes apps use
    storage, as well as how storage systems (including databases) may
    be deployed on Kubernetes .  We collaborate with Apps SIG and make
    sure that important topics are well covered.
*   **[Kubernetes Service Catalog SIG](https://github.com/kubernetes/community/tree/master/sig-service-catalog)**-
    works on enabling external managed software offerings such as
    datastore services offered by public cloud providers.


# Operating Model

This TAG follows the [standard operating
guidelines](https://github.com/cncf/toc/blob/main/tags/cncf-tags.md#operating-model)
provided by the TOC unless otherwise stated here.

**TOC Liaison:** [Erin Boyd](https://github.com/erinboyd), [Saad Ali](https://github.com/saad-ali)

**TAG Chairs:** [Alex Chircop](https://github.com/chira001), [Quinton Hoole](https://github.com/quinton-hoole), [Xing Yang](https://github.com/xing-yang) 

**Tech Leads:** Nick Connolly, Sugu Sougoumarane, Luis Pabon

**Other named roles: **None at present; will be identified and staffed as needed.
