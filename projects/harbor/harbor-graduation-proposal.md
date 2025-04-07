== Harbor Graduation Proposal

Since joining the CNCF in [July 2018](https://www.cncf.io/blog/2018/07/31/cncf-to-host-harbor-in-the-sandbox/) as a Sandbox project and in [November 2018](https://www.cncf.io/blog/2018/11/13/harbor-into-incubator/) moving to an Incubating Project, Harbor has built a healthy ecosystem of users, maintainers, and production implementations.

On behalf of the maintainers team, we believe Harbor is ready for the CNCF [Graduation stage](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#graduation-stage) and meets the v1.2 criteria. We address the [CNCF Technical Due Diligence questions](tbd) in a separate Google Docs document (referred to as the `Tech-DD` from now on).

*Name of project:* Harbor

*Project Description:* Harbor is an open source cloud native registry that provides trust, compliance, performance, and interoperability. As a private on-premises registry, Harbor fills a gap for organizations that prefer not to use a public or cloud-based registry or want a consistent experience across clouds. Harbor secures images with role-based access control, scans images for vulnerabilities, and signs images as trusted. A CNCF Incubating project, Harbor helps you consistently and securely manage images across cloud native compute platforms like Kubernetes and Docker.

Harbor in essence solves common problems in organizations building cloud native applications by offering them a choice to deploy their own private on-premises registry and apply a consistent policy across artifact management in any cloud.

=== CNCF Alignment & Why does CNCF need a container registry? 

The CNCF has an impressive portfolio of projects that can be leveraged to build and run complex distributed systems; a gap, however, exists without a secure container registry. In particular, no other open source container registry offers the feature set present in Harbor.

Harbor's features and community are a natural fit for the CNCF. CNCF offers a vendor-neutral home for the project, while increasing community involvement and feature velocity, with a tighter alignment between Harbor and other CNCF projects.

Harbor's core mission aligns well with graduated projects like Kubernetes and the container ecosystem. The CNCF's mission is to “create and drive the adoption of a new computing paradigm that is optimized for modern distributed systems environments capable of scaling to tens of thousands of self-healing multi-tenant nodes.” We believe container registries are essential to achieve this mission. Harbor, as a mature open source registry is a logical complement to the CNCF's existing portfolio of graduated projects.

*Registry Landscape:* A comparison of Harbor with other non-CNCF similar projects is available in the Tech-DD.

==== Features

The mission of Harbor is to provide users in cloud native environments the ability to confidently manage and securely serve container images and Helm charts. To do so, Harbor stores, signs, and scans content. Here are some of the key features of Harbor organized by user needs:

*Security & Compliance*  
Users ask for a comprehensive policy, registry and data ownership, identity federation with built-in multitenancy. Harbor delivers:

 * Security & Vulnerability Scanning
 * CVE Exception Policy
 * Image Signing
 * Quotas Management
 * Retention Policies
 * Audit logging
 * OIDC/LDAP Integration for Identity Federation
 * Role Based Access Control & CLI secrets
 * Project Isolation and Multitenancy

*Infrastructure*
Users ask for the ability to deploy the registry on any infrastructure (private, public, hosted, edge), data locality, and the ability to  execute containers on Kubernetes and Docker. Harbor delivers on every one of these requests including:

 * Internationalization (currently English, Chinese, Spanish, Portuguese, French, and Turkish)

*Scale & Control*
Users ask for controlled access to artifacts, the ability to replicate resources based on business needs, and a scalable registry. Harbor delivers:

 * Replicate Harbor artifacts to another instance of Harbor
 * Replicate Harbor artifacts to Docker Registry
 * Replicate Harbor artifacts to every major cloud provider's registry, including Docker Hub, Huawei Cloud, AWS, Azure, GCP, Alibaba Cloud

*Automation & Extensibility*
Users ask for the registry to have the ability to be plug-n-play with existing investments in infrastructure and services. Harbor delivers:

 * Syslog integration
 * Webhooks
 * Extensible REST API and graphical UI
 * Robot Accounts

To learn more about Harbor's features, read and view the following resources:

 * [Intro to Harbor](https://blogs.vmware.com/cloudnative/2018/06/14/harbor-delivers-a-trusted-cloud-native-registry)
 * [Going Big: Harbor 1.8 Takes Security and Replication to New Heights](https://goharbor.io/blogs/announcing-harbor-1.8/)
 *	[Harbor 1.9 tackles multi-tenancy with key enterprise features](https://goharbor.io/blogs/harbor-1.9/)
 *	[Harbor CNCF webinar on Security and Day 2 operations](https://www.cncf.io/webinars/kubernetes-in-highly-restrictive-environments- meeting-the-needs-of-enterprise-governance-security-2/)
 * [Harbor CNCF webinar on Gain Confidence in Compliance: Advanced Image Scanning with Harbor](https://www.cncf.io/webinars/gain-confidence-in-compliance-advanced-image-scanning-with-harbor/)

=== Project Timeline and Snapshot

A brief history of Harbor is included in the Tech-DD.

== Production Users

A comprehensive adoption guide is included in the Tech-DD.

== In-Flight Features

The Harbor team is currently working on improving Harbor with our 1.10 release. The theme of the release is `Security & Compliance` and some of the key features include:

 * Immutable Images and Repositories
 * OIDC Group Support
 * Limited Guest Role
 * CLI Secret and Robot account enhancements
 * Interrogation Service
 * Pluggable out-of-tree scanning by Aqua Trivy
 * Pluggable out-of-tree scanning by Anchore Engine and Anchore Enterprise

The direction of the project has been generally guided by our open source community and users. There are a plethora of GitHub issues requesting various features that we prioritize based on popularity of user requests and engineering capacity. Our community has been involved in the addition of several new important features, including webhooks and the pluggable scanner work.

A roadmap for future features, including those listed above, can be found in the Tech-DD. The project welcomes contributions of any kind: features, bug fixes, documentation, bug reporting via issues, security testing, and project management to help track and prioritize workstreams.

== Use Cases

The following is a list of common use-cases for Harbor users:

 * *On-prem container registry* – organizations with the desire to host sensitive production images on-premises can do so with Harbor
 * *Edge container registry* – organizations with the desire to host images near edge compute nodes because links ultimately fail
 * *Vulnerability scanning* – organizations can scan images before they are used in production to safeguard against threats. Images with failed vulnerability scans can be blocked from being pulled
 * *Image signing* – images can be signed via Notary to ensure provenance
 * *Role-based Access Control* – integration with LDAP (and AD) or OIDC to provide user- and group-level permissions
 * *Image replication* – production images can be replicated to disparate registries, providing disaster recovery, load balancing and the ability for organizations to replicate images to different geos to provide a more expedient image pull

== CNCF Requirement Details
 * *Preferred Maturity Level:* Graduation Stage
 * *TOC Sponsor:* Joe Beda
 * *License:* Apache License 2.0
 * *Source control repositories / issue tracker:* https://github.com/goharbor
 * *Website:* https://goharbor.io
 * *Release Methodology and Mechanics:* Covered in Tech-DD
 * *Receive a supermajority vote from the TOC to move to graduation stage:* We believe this document and the Tech-DD prepares us for a TOC vote on the graduation readiness of Harbor.

== Social Media Accounts (more details and membership numbers in the Tech-DD):

 * *Twitter:* https://twitter.com/project_harbor
 * *Users Google Groups:* https://lists.cncf.io/g/harbor-users/topics
 * *Developer Google Groups:* https://lists.cncf.io/g/harbor-dev/topics
 * *Slack:* https://cloud-native.slack.com/messages/harbor
 * *WeChat:* Harbor Community

== Contributor Statistics
Covered in depth in the Tech-DD

== Architecture
A detailed Harbor architecture and its various components is covered in the Tech-DD
