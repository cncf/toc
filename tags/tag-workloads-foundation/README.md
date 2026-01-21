# TAG Workloads Foundation

<!-- THIS FILE IS AUTO-GENERATED FROM /tags.yaml -->

## Mission Statement
To define and advance practices and standards for fundamental cloud native workload execution environments and their related lifecycle management within cloud native systems, applications, and architectures. This supports the CNCF's technical vision by addressing critical problems faced by adopters and contributing to a robust cloud native ecosystem.


[Charter](./charter.md)

## Leadership
### Chairs
- Danielle Tal (**[@miao0miao](https://github.com/miao0miao)**) (Term: 2025-07-01 - 2026-06-30)
- Paco Xu (**[@pacoxu](https://github.com/pacoxu)**) (Term: 2025-07-01 - 2027-06-30)
- Yuan Tang (**[@terrytangyuan](https://github.com/terrytangyuan)**) (Term: 2025-07-01 - 2027-06-30)
### Tech Leads
- Alexander D. Kanevskiy (**[@kad](https://github.com/kad)**) (Term: 2025-07-02 - 2026-06-30)
- Hongcai Ren (**[@RainbowMango](https://github.com/RainbowMango)**) (Term: 2025-07-02 - 2027-06-30)
- Kante Yin (**[@kerthcet](https://github.com/kerthcet)**) (Term: 2025-07-02 - 2027-06-30)
- Rajas Kakodkar (**[@rajaskakodkar](https://github.com/rajaskakodkar)**) (Term: 2025-07-02 - 2027-06-30)
- Stephen Rust (**[@srust](https://github.com/srust)**) (Term: 2025-07-02 - 2026-06-30)

## Meetings
- **TAG Workloads Foundation Meetings**: [Calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/tag-workloads-foundation?view=list) | [Recordings](https://www.youtube.com/@CNCFTAGWorkloadsFoundation) | [Meeting Notes and Agenda](https://notes.cncf.io/s/1aNdplhtl)

## Contact
- Slack: [Tag TAG Workloads Foundation Slack](https://cloud-native.slack.com/archives/C08K71W9HAS)
- [Mailing List](https://lists.cncf.io/g/cncf-tag-workloads-foundation)
- TOC Liaison: Alex Chircop (**[@chira001](https://github.com/chira001)**)
- TOC Liaison: Kevin Wang (**[@kevin-wangzefeng](https://github.com/kevin-wangzefeng)**)

## Subprojects
### Batch
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

- [Mailing List](https://lists.cncf.io/g/cncf-tag-workloads-foundation)
## Initiatives
[TAG Workloads Foundation Initiatives](https://github.com/cncf/toc/issues?q=state%3Aopen%20label%3Atag%2Fworkloads-foundation%20label%3Akind%2Finitiative)