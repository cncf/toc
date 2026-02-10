# Cloud Native Security Controls Refresh

Creation Issue: https://github.com/cncf/toc/issues/1910

## Objective

The objective is to refresh the existing [Cloud Native Security Controls Catalog](https://docs.google.com/spreadsheets/d/1GUohOTlLw9FKUQ3O23X7ypvJLXN-B3veJGe6YE6JYfU/edit) artifact to enhance its consistency and maintenance by expressing the content in a standardized, machine-readable format.

## Logistics

The initiative meets weekly on Fridays at 17:30 UTC on the [TAG Security and Compliance calendar](https://zoom-lfx.platform.linuxfoundation.org/meetings/tag-security-and-compliance?view=list).

* [Meeting notes](https://notes.cncf.io/64WPMKmDTOO2WnDI0av_Rw)
* [Slack channel](https://cloud-native.slack.com/archives/C09TLL22PK9)

## Planned Milestones

### Milestone 1: CNSC Gemara-compatible Release Process

**Goal:** Convert existing controls into a standardized, machine-readable format that can automatically generate artifacts in a continuous release pipeline, such as markdown, PDF, and/or web page.

| **Checkpoint** | **Activity Description**                                                                              |
|:---------------|:------------------------------------------------------------------------------------------------------|
| 1.1            | Parse and format existing controls to a [Gemara](https://gemara.openssf.org/)-based Guidance Document |
| 1.2            | Set up basic tooling and pipeline for release artifacts                                               |

### Milestone 2: Project-specific Gemara Compatibility Assessment

**Goal:** Determine the value of a longstanding sub-project aimed at creating technology-specific artifacts, pausing to assess practicality after creating detailed threat catalogs.

| **Checkpoint** | **Activity Description**                                                       |
|:---------------|:-------------------------------------------------------------------------------|
| 2.1            | Identify a project to create "Layer 2" Gemara artifacts for (e.g. Kyverno)     |
| 2.2            | Create a capabilities document to identify the features of the technology      |
| 2.3            | Create a threats document to map common threats to the technology capabilities |
| 2.4            | Assess gaps and progress to determine whether to continue to Milestone 3       |

### Milestone 3: Project-specific CNSC Artifacts

**Goal:** Determine the value of a longstanding sub-project aimed at creating technology-specific control catalogs.

| **Checkpoint** | **Activity Description**                                                                         |
|:---------------|:-------------------------------------------------------------------------------------------------|
| 3.1            | Identify integration points between Milestone 1 and Milestone 2 artifacts                        |
| 3.2            | Create control objectives which mitigate the threats identified in Milestone 2                   |
| 3.3            | Create assessment requirements to guide implementation and evaluation of the control objectives  |
| 2.4            | Assess gaps and progress to determine whether to propose a sub-project for this type of activity |