# CNCF Batch Subproject

<!-- THIS FILE IS AUTO-GENERATED FROM /tags.yaml -->

## Mission Statement
To enhance collaboration among projects, improve interoperability, and empower users to efficiently leverage batch systems in cloud-native environments.

In scope:

To reduce fragmentation in the k8s batch ecosystem: congregate leads and users from different external and internal projects and user groups (CNCF TAGs, k8s sub-projects focused on batch-related features such as topology-aware scheduling) in the batch ecosystem to gather requirements, validate designs and encourage reutilization of core K8s APIs.

The following recommendations for enhancements:

* Additions to the batch API group, currently including Job and CronJob resources that benefit batch use cases such as HPC, AI/ML, data analytics and CI.
* Primitives for job-level queueing, not limited to the k8s Job resource. Long-term, this could include multi-cluster support.
* Primitives to control and maximize utilization of resources in fixed-size clusters (on-prem) and elastic clusters (cloud).
* Benchmarking models for Batch systems
* Data Locality
* User Stories
* Scheduling support for specialized hardware (Accelerators, NUMA, Networking, etc.)

Out of scope:

* Addition of new API kinds that serve a specialized type of workload. The focus should be on general APIs that specialized controllers can build on top of.
* Uses of the batch APIs as support for serving workloads (eg. backups, upgrades, migrations). These can be served by existing SIGs.
* Proposals that duplicate the functionality of core kubernetes components (job-controller, kube-scheduler, cluster-autoscaler).
* Job workflows or pipelines. Mature third party frameworks serve these use cases with the current kubernetes primitives. But additional primitives to support these frameworks could be in scope.

Deliverable(s) or exit criteria:

* Maintaining a landscape document for currently available projects (already published-relocated and maintained)
* Data Locality project-deliverables TBD, but something that helps in this space (already in process)
* Benchmarking suite for Batch systems (already in process)
* User stories published doc for Batch systems (already in process)

[Charter](./charter.md)

## Leadership
### Subproject Leads
- Alex Scammon (**[@](https://github.com/)**)
- Marlow Warnicke (**[@catblade](https://github.com/catblade)**)
- Abishek Malvankar (**[@](https://github.com/)**)

## 💬 Contact
- [Slack Channel](https://cloud-native.slack.com/archives/C08K71W9HAS) on cloud-native.slack.com
- 📧 [Mailing List](https://lists.cncf.io/g/cncf-tag-workloads-foundation)
