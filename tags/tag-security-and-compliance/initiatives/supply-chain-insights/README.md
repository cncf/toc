# CNCF Software Supply Chain Insights

Creation issue: [#1709](https://github.com/cncf/toc/issues/1709)

## Objective

Pravide a template for particpants wishing to run a scalable system to automate the collection, correlation, and understanding of CNCF Projects' supply chain metadata.  This initiative is related to, but separate from the CNCF staff efforts to establish a GUAC-based supply chain metadata instance to answer CNCF priorities, but may be informed by that effort to enable other organizations to replicate and expand upon the work.

### Desired outcomes

1. Establish whether a community-operated supply chain analysis platform template is worth establishing as a long-term subproject.

1. Provide feedback to utilized CNCF and OpenSSF supply chain insights projects (such as GUAC) about the ability to adopt and use their platforms.

1. Assess the utility of and establish patterns for CNCF projects to publish supply chain metadata, including both dependency information and other supply chain security data.

## Logistics

This initiative is being led by Matt Young (@halcyondude), with TAG-SC support from Evan Anderson (@evankanderson).  The initiative meets weekly on Mondays at 1600GMT on the [TAG SC calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/tag-security-and-compliance?view=list).

[Meeting notes](https://notes.cncf.io/cmBr4VUwS3qSHo3ABM6Tmw)

## Milestones

Further milestones will be established once milestone 1 is complete, and the ability to answer selected supply chain questions has been assessed.

### Milestone 1

Configuration, queries, documentation, and other necessary content for a CNCF community member to be able to set up a GUAC instance running on a single-node cluster (e.g. kind) which can ingest SBOMs for the software running on the cluster and answer the following questions:

* What projects would be impacted if a given library changed to an incompatible non-OSI license ("rug pull")?
* What projects would be impacted by the discovery of a CVE in a given library?

Further questions might include:

* Which libraries would cause the _largest impact_ due to a license change.