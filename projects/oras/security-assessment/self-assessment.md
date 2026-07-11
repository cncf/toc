---
title: ORAS Self-Assessment

---

# ORAS Self-Assessment

Security reviewers: Andrew Block

This document is intended to aid in roadmapping, and the onboarding of new maintainers and provide a good introduction to the security model of ORAS.

## Table of Contents

* [Metadata](#metadata)
  * [Security links](#security-links)
* [Overview](#overview)
  * [Actors](#actors)
  * [Actions](#actions)
  * [Background](#background)
  * [Goals](#goals)
  * [Non-goals](#non-goals)
* [Self-assessment use](#self-assessment-use)
* [Security functions and features](#security-functions-and-features)
* [Project compliance](#project-compliance)
* [Secure development practices](#secure-development-practices)
* [Security issue resolution](#security-issue-resolution)
* [Appendix](#appendix)

## Metadata

| | |
|-----------|------|
| Assessment Stage | Incomplete |
| Software | <https://github.com/oras-project> |
| Security Provider? | No |
| Languages | Go, Python, Java, .NET, Rust |
| Software Bill of Materials | ORAS does not currently publish SBOM's as part of the release process of any of its subprojects. |
| | |

### Security links

| Doc | url |
| -- | -- |
| Security file | <https://github.com/oras-project/community/blob/main/SECURITY.md> |
| Security Policy | <https://oras.land/community/reporting_security_concerns> |

## Overview

ORAS (OCI Registry As Storage) is the de facto tool for working with OCI Artifacts and provides CLI and client libraries to distribute artifacts across OCI-compliant registries.

### Background

OCI registries are a core component within cloud computing as it provides storage for container images. Over time, there became an interest to store additional types of content aside from container images which led to the standardization of OCI artifacts as a solution. The ORAS project was born out of the need to enable the management and lifecycle of OCI artifacts along with the interaction with OCI registries.

### Actors

1. [oras](https://github.com/oras-project/oras) is a golang based CLI which serves as the primary tool for end user use.
2. Client libraries are available in a variety of programming language for use for managing OCI artifacts using ORAS:

    1. [oras-go](https://github.com/oras-project/oras-go)
    2. [oras-py](https://github.com/oras-project/oras-py)
    3. [oras-dotnet](https://github.com/oras-project/oras-dotnet)
    4. [oras-java](https://github.com/oras-project/oras-java)
    5. [rust-oci-client](https://github.com/oras-project/rust-oci-client)
3. A [Custom GitHub Action](https://github.com/oras-project/setup-oras) to enable the installation of the ORAS CLI within [GitHub Actions](https://docs.github.com/en/actions)
4. A [MCP Server](https://github.com/oras-project/oras-mcp) using the [Model Context Protocol](https://modelcontextprotocol.io) to manage OCI artifacts in remote registries

### Actions

#### Interacting with OCI Registries by Client Libraries

The interface between end users/applications and OCI registries is facilities by a series of client side libraries. These libraries are written in a variety of popular programming languages and provide the functionality to interact with remote OCI registries.

#### Integration with Client Libraries

ORAS client libraries are the primary interface between tooling contained within the project (such as the `oras` CLI) and third party applications. Capabilities are provides to manage the lifecycle of OCI artifacts locally (including interacting with the local filesystem/Operating System) and remote OCI registries.

#### Using the oras Command Line Interface (CLI)

The `oras` Command Line Interface (CLI) represents how most end users and tools interact with the ORAS project. Methods are available for managing OCI artifacts locally as well as in remote OCI registries in an easy to use, human-friendly utility.

#### AI Workflow Interactions

AI Applications can integrate with functionality provided by the ORAS project using the [Model Context Protocol](https://modelcontextprotocol.io/) and the [oras-mcp](https://github.com/oras-project/oras-mcp) project.

### Goals

The goals associated with the ORAS project are to simply the lifecycle working with OCI artifacts. Specifically:

**Awareness of OCI Artifacts and their benefits**

Demonstrate the use of OCI artifacts as a method of packaging and storing content using OCI structures.

**Creation and management of OCI Artifacts**

Tooling to support the creation and management of OCI artifacts both within a local environment/machine along with remote OCI registries.

**Tooling for the mangement of OCI assets**

In addition to OCI artifacts themselves, the project provides utilities for interacting with additional OCI related resources, such as Manifests.

**Third party integration**

While the `oras` Command Line Interface (CLI) is the primary interface for most end users, the project seeks to enable integration by third party utilities by providing a set of client libraries in a variety of programming languages. 

### Non-Goals

While not the primary areas of focus, the ORAS project seeks to provide the following:

**Enhance the Software Supply Chain**

OCI artifacts that are created and managed by ORAS commonly include assets related to the Software Supply Chain. The ability to leverage these resources as part of the software development and delivery lifecycle increases the overall security posture of projects utilizing such approaches.

**Support the evolution of OCI standards**

The use of OCI Artifacts and the structures governed by the Open Container Initiative through the ORAS project spreads awareness of the OCI community, their standards, and offers opportunities to further conversations to support the use of OCI related resources today and tomorrow.

## Self-assessment Use

This self-assessment is created by the ORAS team to perform an internal analysis of the project's security. It is not intended to provide a security audit of ORAS, or function as an independent assessment or attestation of ORAS's security health.

This document serves to provide ORAS users with an initial understanding of ORAS's security, where to find existing security documentation, ORAS plans for security, and general overview of ORAS security practices, both for development of ORAS as well as security of ORAS.

This document provides ORAS maintainers and stakeholders with additional context to help inform the roadmap creation process, so that security and feature improvements can be prioritized accordingly.

## Security functions and features

### Critical

**Credential Management**

The ORAS project enables users to communicate with protected OCI registries. The project adheres to standards as specified by OCI and Docker which includes both the format as well as the storage mechanisms.

### Security Relevant

**Transport Protocol with OCI Registries**

ORAS and the underlying libraries provide protections to ensure that communications with remote resources are protected (communicating over secure channels along with the verification of TLS certificates). Functionality is available within the project to enable end users and systems to both specify how the connection should be achieved as wel as bypass these protections (disabling TLS verification as well as communicating over unencrypted channels) if desired.

**Methods of Specifying Authentication Details**

Credentials associated with remote OCI registries may be stored locally for later reuse. End users may choose alternately provide credentials at runtime when communicating with remote OCI registries.

**Material Stored in OCI Artifacts**

ORAS facilitates the creation and management of OCI artifacts. End users may choose to store sensitive information either within the metadata (Manifests/Descriptors) or in the content of the artifact. Appropriate security practices should be utilized to ensure any sensitive information.

## Project Compliance

ORAS does not document meeting any particular compliance standards.

## Secure Development Practices

The ORAS project follows established CNCF and OSS best practices for code development and delivery.

### Development Pipeline

All source code is maintained in Git within the [oras-project GitHub Organization](https://github.com/oras-project). The project has a detailed set of guide for individuals interested in contributing to the project. These resources can be found below:

* [Contribution Guide](https://oras.land/community/contributing_guide)
* [Developer Guide](https://oras.land/community/developer_guide)

#### Contributing to the Project

Contributors to ORAS are required to sign their commits, adhering to the [Developer Certificate of Origin (DCO)](https://probot.github.io/apps/dco). This practice ensures the integrity of the code by verifying that the changes are made by the person who claims to have made them. Contributors use the Signed-off-by line in commit messages to signify their adherence to these requirements. Git has a `-s` command-line option to append this automatically to commit messages.

Code contribution should follow the [Fork and Pull workflow](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork) which provides a streamlined mannner of accepting contributions by the community.

Several roles have been established to goven the activities within the ORAS project. Descriptions of each role, their responsibilities, and the process for which individuals within the community can fulfil such obligations are detailed within the [Contributor Ladder](https://oras.land/community/contributor_ladder) documentation.

#### Release Process

The steps involved including the detailed procedures for proposing and releasing ORAS project related content is documented wihtin the [Release Process](https://oras.land/community/release-process) section of the ORAS project website and includes the following:

* Prerequisites activities that must be performed prior to performing a release
* Proposing a release
* Detailed steps for updating source code
* Updating documentation
* Announcing a release
* Published released artifacts to external sources

### Community Channels

The ORAS project provides several different ways methods, either synchroouslty or asynchronously, for the Open Source community to interact with the project. A list of key resources cn be foud below:

* [GitHub](https://github.com/oras-project/) - Issues/Pull Requests/Discussios
* [#oras](https://cloud-native.slack.com/archives/CJ1KHJM5Z) channel within the [CNCF Workspace](https://cloud-native.slack.com)
* Bi-weekly community meetings
* Social Media ([X/Twitter](https://x.com/orasproject) and [YouTube](https://www.youtube.com/@cncforasproject))

The full set of assets can be found on the [Community Resources](https://oras.land/community/community_resources) page of the ORAS website.

## Security Issue Resolution

The ORAS security policy is maintained in the SECURITY.md file within the [community repository](https://github.com/oras-project/community) and can be quickly found through the [GitHub Security Overview](https://github.com/oras-project/oras/security/policy).

### Responsible Disclosure Practice

The ORAS project accepts vulnerability reports through the [GitHub Vulnerability Reporting](htt‌ps://docs.github.com/en/code-security/security-advisories/guidance-on-reporting-and-writing/privately-reporting-a-security-vulnerability) tool.

Anyone can make a report by going to the reporting form in any of the ORAS project GitHub repositories. In the event that a report is received, a maintainer will collaborate directly with the reporter through the Security Advisory until it is resolved.

### Incident Response

A dedicated security team consisting of a subgroup of the project maintainers has been established within the ORAS project. Their job is to manage the lifecycle of security related concerns within the project.

In the event that a vulnerability is reported, the security team will collaborate to determine the validity and criticality of the report. Based on these findings, the fix will be triaged and the security team will work to issue a patch in a timely manner.

Patches will be made to all versions that are currently supported under the project's security policy. Information will be disseminated to the community through all appropriate outbound channels as soon as possible based on the circumstance.

## Appendix

**Known Issues Over Time**
  
The ORAS project does not contain any known vulnerability in any released of the released versions. If a vulnerability is discovered in the future, an issue would be raised in within the GitHub repository associated with the vulnerability and disseminated within the documented project communication channels.

**OpenSSF Best Practices**

Several of the repositories within the ORAS project have begun their journey toward attaining [OpenSSF Best Practices](https://www.bestpractices.dev/en) accreditation:

  * [oras](https://www.bestpractices.dev/en/projects/5562)
  * [oras-java](https://www.bestpractices.dev/en/projects/10047)

Issues have been created within the project to account for the areas of improvements needed to meet OpenSSF Best Practices guidelines. It is on the ORAS project roadmap that all code repositories attain OpenSSF Best Practices accreditation.

**Case Studies**

The ORAS project has gained adoption by a variety of projects within the CNCF and Open Source Community. A list of of entities that are either leveraging or have integrated the ORAS project can be found on the [Adopters](https://oras.land/adopters) page of the ORAS website.

The following are specific examples where ORAS has become a fundamental aspect of their solution:

- [Helm](https://helm.sh) is a package manager for Kubernetes. ORAS has become a key integration for Helm as it has enabled the ability to store Helm charts as OCI artifacts for storage within OCI registries. By simplifying how Helm Charts are distributed and consumed using common tooling within Cloud Native environments, it has become the recommended option for managing Helm content.

**Related Projects / Vendors**

While the primary goal of the ORAS project focuses solely on the management of OCI artifacts, other solutions are available that provide overlapping capabilities with the ORAS within their solution:

- [Podman](https://podman.io) - Includes capabilities for creating, managing and publishing OCI artifacts to remote OCI registries.
