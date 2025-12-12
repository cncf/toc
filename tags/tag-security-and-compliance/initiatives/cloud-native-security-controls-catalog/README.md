# Cloud Native Security Controls Refresh

Creation Issue: https://github.com/cncf/toc/issues/1910

## Objective

The objective is to refresh the existing [Cloud Native Security Controls Catalog](https://docs.google.com/spreadsheets/d/1GUohOTlLw9FKUQ3O23X7ypvJLXN-B3veJGe6YE6JYfU/edit) artifact to enhance its consistency and maintenance by expressing the content in a standardized, machine-readable format.

### Desired Outcomes

1. Determine whether a community-maintained controls catalog is worth establishing as a long-term subproject.

2. Assist CNCF project maintainers to build technology-specific best practices guides which could accelerate joint security assessments and streamline project promotions.

3. Establish strategic alignment with other organizations (e.g., FINOS Common Cloud Controls) to reuse control definitions, avoid duplication, and potentially contribute completed sections back to external projects.

### Proposed Deliverables

* A standardized, machine-readable [Gemara](https://gemara.openssf.org/) guidance document that can be rendered to various human-readable formats for consumption.
* A proof-of-concept granular, technology-specific (or project-specific) controls catalog that is measurable against automated assessments and functions as security acceptance criteria.
* Deployment and maintenance tooling to support publishing on `contribute.cncf.io` website under TAG Security and Compliance [publications](https://contribute.cncf.io/community/tags/security-and-compliance/#publications).
* A contributing guide for continued community-driven maintenance (if moving to Milestone 2)

## Logistics

The initiative meets weekly on Fridays at 17:30 UTC on the [TAG Security and Compliance calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/tag-security-and-compliance?view=list).

* [Meeting notes](https://notes.cncf.io/64WPMKmDTOO2WnDI0av_Rw)
* [Slack channel](https://cloud-native.slack.com/archives/C09TLL22PK9)

## Planned Milestones

### Milestone 1

**Goal:** Convert existing controls into a standardized, machine-readable format to test the tooling and community commitment.

| **Milestone** | **Activity Description**                                                                              | **Output Artifact**             |
|:--------------|:------------------------------------------------------------------------------------------------------|:--------------------------------|
| 1.1           | Parse and format existing controls to a [Gemara](https://gemara.openssf.org/)-based Guidance Document | Draft Layer 1 Guidance Document |
| 1.2           | Set up basic tooling and pipeline                                                                     | Working Conversion Pipeline     |

**Checkpoint**: Based on the successful delivery of the guidance document, is there sufficient contributor interest and resource commitment to take on the ongoing 
maintenance required breaking it down into granular, technology-specific controls?

### Milestone 2

**Goal:** Begin the high-effort work of creating granular, technology-specific controls and assessment requirements by aligning with external projects and collaborating with CNCF maintainers.

| **Milestone** | **Activity Description**                                                                                           | **Output Artifacts**                     |
|:--------------|:-------------------------------------------------------------------------------------------------------------------|:-----------------------------------------|
| 2.1           | Categorize Layer 1 controls for future Layer 2 scope (e.g., Cluster, Container Workloads)                          | Draft Layer 2 Controls Catalog Documents |
| 2.2           | Formalize reuse/contribution strategy with FINOS CCC, the OSPS Baseline, and other external catalogs               | Draft Layer 2 Controls Catalog Documents |
| 2.3           | Select a pilot project to begin refining granular Layer 2 controls for their project's technology and capabilities | Layer 2 Controls Catalog Document POC    |
| 2.4           | Expand tooling to all desired formats as requested by the community                                                | Conversion Tooling and Documentation     |

**Checkpoint**: Have we successfully demonstrated the model for collaboration and creation of a sustainable, measurable Controls Catalog artifact?
