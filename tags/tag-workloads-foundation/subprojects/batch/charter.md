# TAG Workloads Foundation Batch Subproject Charter

## Mission

The cloud-native batch scheduling ecosystem is fragmented — different projects tackle job scheduling, queueing, and resource management in incompatible ways. The Batch subproject brings together maintainers and users across the ecosystem to reduce that fragmentation: aligning on common Kubernetes APIs and primitives, developing best practices, and improving outcomes for batch workloads — whether HPC, AI/ML, data analytics, or CI — in cloud-native environments.

## Scope

### In Scope

To reduce fragmentation in the Kubernetes batch ecosystem: congregate leads and users from different external and internal projects and user groups (CNCF TAGs, Kubernetes sub-projects focused on batch-related features such as topology-aware scheduling) in the batch ecosystem to gather requirements, validate designs and encourage reutilization of core Kubernetes APIs.

The following recommendations for enhancements:

* Additions to the batch API group, currently including Job and CronJob resources that benefit batch use cases such as HPC, AI/ML, data analytics and CI.
* Primitives for job-level queueing, not limited to the Kubernetes Job resource. Long-term, this could include multi-cluster support.
* Primitives to control and maximize utilization of resources in fixed-size clusters (on-prem) and elastic clusters (cloud).
* Benchmarking models for Batch systems
* Data Locality
* User Stories
* Scheduling support for specialized hardware (Accelerators, NUMA, Networking, etc.)

### Out of Scope

* Addition of new API kinds that serve a specialized type of workload. The focus should be on general APIs that specialized controllers can build on top of.
* Uses of the batch APIs as support for serving workloads (eg. backups, upgrades, migrations). These can be served by existing SIGs.
* Proposals that duplicate the functionality of core Kubernetes components (job-controller, kube-scheduler, cluster-autoscaler).
* Job workflows or pipelines. Mature third party frameworks serve these use cases with the current Kubernetes primitives. But additional primitives to support these frameworks could be in scope.

## Deliverables

* **Project Landscape** — a living catalogue of batch scheduling projects in the cloud-native ecosystem, maintained at [bsi-landscape.netlify.app](https://bsi-landscape.netlify.app/).
* **Whitepapers and Technical Research** — the subproject produces papers and research on topics relevant to cloud-native batch scheduling, such as benchmarking of batch systems, data locality, scheduling best practices, and user stories. An initial series of five whitepapers is complete, with more planned as the space evolves.
