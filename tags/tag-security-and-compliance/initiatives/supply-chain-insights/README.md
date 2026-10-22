# CNCF Software Supply Chain Insights

Creation issue: [#1709](https://github.com/cncf/toc/issues/1709)

## Objective

Provide a template for particpants wishing to run a scalable system to automate the collection, correlation, and understanding of CNCF Projects' supply chain metadata.  This initiative is related to, but separate from the CNCF staff efforts to establish a GUAC-based supply chain metadata instance to answer CNCF priorities, but may be informed by that effort to enable other organizations to replicate and expand upon the work.

## Results (Summary)

Exploration was performed in https://github.com/halcyondude/supply-chain-security-collector/, and the results were [presented at the TAG Security and Compliance meeting on 8 April 2026](https://github.com/halcyondude/supply-chain-security-collector/tree/main/docs/presentations/2026-04-08-tag-sc).

The exploration ran into a number of **blocking problems** which would need to be addressed before this initiative was re-attempted.  A parallel CNCF-staff effort used a substantially different approach (rebuild SBOMs from project sources without using project-published metadata) with substantially more success, indicating that project-published SBOMs may not currently be sufficient for supply chain analysis purposes.

### Problems Discovered

1. **Project-published SBOMs are not easily discoverable nor of uniform quality.**

   We manually collected 15 known-SBOM-publishing graduated CNCF projects (where "known" means they had some documentation about SBOMs), along with a survey of 225 graduated through sandbox projects.  We were unable to reliably locate published SBOMs through the GitHub GraphQL interfaces, and ended up resorting to scanning for known SBOM tools in GitHub Actions.  Once located, SBOMs did not necessarily contain e.g. dependency tree information, instead presenting dependencies as a flat list, which frustrated deeper analysis on dependency cut points.

1. **GUAC and GraphQL is not suitable for large-scale analytics**

   The GUAC usage of Postgres and GraphQL is not very performant and has difficulty expressing more complex queries.  We ended up extracting supply chain data from GUAC to DuckDB and then LadybugDB in order to be able to run Cypher queries.  This effort ended up being a substantial diversion from the main initiative.

1. **CI workflows are highly divergent**

   Given the difficulties in locating SBOMs from the first section, we analyzed usage of tools such as cosign and SBOM generators (trivy, syft, spdx-sbom-generator, tern).  Given the wide range of CI and release workflows, it is difficult to recommend a single or small set of tooling improvements across the CNCF ecosystem.  In turn, this reduces the leverage to improve the quality of published SBOMs and signatures of the same.

## Process

Also in https://notes.cncf.io/cmBr4VUwS3qSHo3ABM6Tmw

1. Establish whether a community-operated supply chain analysis platform template is worth establishing as a long-term subproject.

1. Provide feedback to utilized CNCF and OpenSSF supply chain insights projects (such as GUAC) about the ability to adopt and use their platforms.

1. Assess the utility of and establish patterns for CNCF projects to publish supply chain metadata, including both dependency information and other supply chain security data.

## Logistics

This initiative is being led by Matt Young (@halcyondude), with TAG-SC support from Evan Anderson (@evankanderson).  The initiative meets weekly on Mondays at 1600GMT on the [TAG SC calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/tag-security-and-compliance?view=list).

* [Meeting notes](https://notes.cncf.io/cmBr4VUwS3qSHo3ABM6Tmw)
* [GitHub project](https://github.com/orgs/cncf/projects/80/views/4)

## Milestones

Further milestones will be established once milestone 1 is complete, and the ability to answer selected supply chain questions has been assessed.

### Milestone 1

Configuration, queries, documentation, and other necessary content for a CNCF community member to be able to set up a GUAC instance running on a single-node cluster (e.g. kind) which can ingest SBOMs for the software running on the cluster and answer the following questions:

* What projects would be impacted if a given library changed to an incompatible non-OSI license ("rug pull")?
* What projects would be impacted by the discovery of a CVE in a given library?

Further questions might include:

* Which libraries would cause the _largest impact_ due to a license change.