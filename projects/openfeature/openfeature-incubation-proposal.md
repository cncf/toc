# OpenFeature Incubation Proposal

## Introduction

- Name: OpenFeature
- Website: [openfeature.dev](https://openfeature.dev)
- Source Control: [OpenFeature on Github](https://github.com/open-feature)
- Code of Conduct: [OpenFeature Code of conduct](https://github.com/open-feature/community#code-of-conduct)
- License: [Apache 2](https://github.com/open-feature/community/blob/main/LICENSE.md)

## Overview

OpenFeature provides a shared, standardized feature flagging client - an SDK - which can be plugged into various 3rd-party feature flagging providers. Whether you're using an open-source system or a commercial product, whether it's self-hosted or cloud-hosted, OpenFeature provides a consistent, unified API for developers to use feature flagging in their applications.

**Evaluation API**
The Evaluation API is the primary component of OpenFeature that application authors interact with. The Evaluation API allows developers to evaluate feature flags to alter control flow and application characteristics.

**Providers**
Providers are responsible for performing flag evaluations. They provide an abstraction between the underlying flag management system and the OpenFeature SDK. Providers can wrap a vendor SDK, call a bespoke flag evaluation REST API, or even parse some locally stored file to resolve flag values. This allows the underlying flag evaluation logic to be changed without requiring a major code refactor.
An application integrator can register one provider at a time. Registering an additional provider will override any previously configured providers. If no provider is set, OpenFeature will no-op and return the default value passed to the evaluation API.
Providers are set through the evaluation API. They're globally registered and a change affects both new and existing OpenFeature clients.

**Evaluation Context**
The evaluation context is a container for arbitrary contextual data that can be used as a basis for dynamic evaluation. Static data such as the host or an identifier for the application can be configured globally. Dynamic evaluation context, such as the IP address of the client in a web application, can be implicitly propagated or explicitly passed to during flag evaluation, and can be merged with static values.

**Hooks**
Hooks are a mechanism that allows for the addition of arbitrary behavior at well-defined points of the flag evaluation life-cycle. Use cases include validation of the resolved flag value, modifying or adding data to the evaluation context, logging, telemetry, and tracking.

**Events**
Events enable the ability to react to state changes in the provider or underlying flag management system. These include changes in provider readiness, error status, or perhaps most interestingly, flag configuration changes.

**FlagD**
Flagd is a feature flag daemon with a Unix philosophy. Think of it as a ready-made, open source, OpenFeature compliant feature flag backend system.


## Sponsors from the TOC

Emily Fox 

## Progress Since Sandbox
### Sandbox Acceptance

OpenFeature was accepted into the CNCF Sandbox on [June 14th, 2022](https://github.com/cncf/toc/issues/857).

### Community

**Outreach:** As a project, we hold regular community meetings each month, highlighting what is happening in the community with presentations from multiple adopters and discussions on current topics and interests in the community. Community sessions have, on average **12 attendees**, are uploaded to a [Project meeting playlist hosted on Youtube](<[https://www.youtube.com/@openfeature834/featured)>), and gained **over 1700 views**.

- [OpenFeature Community Hub on Github](https://github.com/open-feature/community)
- [OpenFeature Community on Slack](https://cloud-native.slack.com/archives/C0344AANLA1)

The community statistics are extracted from [CNCF devstats](https://openfeature.devstats.cncf.io/).

### Technical

#### SDK Progress
The project aims to provide a unified API and SDK for feature flag management in various technology stacks. The OpenFeature SDK provides a mechanism for interfacing with an external evaluation engine in a vendor-agnostic way; it does not handle the flag evaluation logic.

Since achieving sandbox, we have progressed in the development of our SDKs:
- [.NET](https://github.com/open-feature/dotnet-sdk) released v1.2.0
- [Java](https://github.com/open-feature/java-sdk) released v1.3.1
- [PHP](https://github.com/open-feature/php-sdk) released v2.0.1
- [Python](https://github.com/open-feature/python-sdk) released v0.0.9
- [JavaScript](https://github.com/open-feature/js-sdk) released v1.3.1
- [GO](https://github.com/open-feature/go-sdk) released v1.4.0
- [Ruby](https://github.com/open-feature/ruby-sdk) released v0.1.0

#### Introduction of flagD
[Released v0.5.4](https://github.com/open-feature/flagd)
- A simple command line tool for fetching and evaluating feature flags for services, designed to conform with the OpenFeature specification.
- OpenFeature compliant with providers available in many languages
- Multiple flag configuration sources including files, http, and Kubernetes
- Accessible over gRPC and HTTP
- Supports subscriptions to real-time flag change events
- Flexible targeting rules based on JSON Logic
- Lightweight daemon, with an emphasis on performance
- Native support for metrics using Prometheus

#### Provider Growth
Providers are responsible for performing flag evaluations and provide an abstraction between the underlying flag management system and the OpenFeature SDK. Providers can wrap a vendor SDK, call a bespoke flag evaluation REST API, or even parse some locally stored file to resolve flag values. This allows the underlying flag evaluation logic to be changed without requiring a major code refactor.

##### Providers
- .NET: Cloudbees, flagD, Go feature flag, LaunchDarkly, Split
- GO: Cloudbees, ConfigCat, flagD, Flagsmith, Flipt, Go feature flag, LaunchDarkly, Split
- Java: Cloudbees, flagD, Flagsmith, Go feature flag, LaunchDarkly, Split
- JavaScript: Cloudbees, ConfigCat, DevCycle, flagD, Go feature flag, LaunchDarkly, PostHog, Split
- PHP: Cloudbees, flagD, Split

## Incubation Stage Requirements

### **_Adhere to CNCF [IP Policy](https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy)_**

As an existing Sandbox project, this is already in place

### **_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope_**
A list of public adopters can be found [here](https://github.com/open-feature/community/blob/main/ADOPTERS.md). There are ~ 10 companies in various stages of adopting, many of whom are non-public in their adoption.
The following are some examples of adopters:

- [Split](https://www.split.io) - Split is excited to announce participation in OpenFeature, an initiative led by Dynatrace and recently submitted to the Cloud Native Computing Foundation (CNCF) for consideration as a sandbox program - Check out the full article [here](https://www.split.io/blog/split-embraces-openfeature/) 
- [Go Feature Flag](https://gofeatureflag.org) - GO Feature Flag believe in OpenSource and offer providers for the new standard for feature flags OpenFeature - check out a more detailed explanation [here](https://gofeatureflag.org/docs/next/openfeature_sdk/concepts)
- [eBay](https://www.ebay.com) - As part of eBay’s tech-led reimagination, we are investing heavily in improving software delivery so that we can launch features more rapidly and safely to our customers. Over the past few months, we have been developing feature flagging capabilities internally, powered by OpenFeature. - read more details [here](https://tech.ebayinc.com/engineering/openfeature-with-contributions-from-ebay-submitted-to-cncfs-sandbox-program/)
- [Proofpoint](https://www.proofpoint.com/) - Our team recently needed to use one of the largest feature management cloud services. This service has a documented SDK with usage guidelines. Thinking about how to define the exposed interfaces, while keeping it simple and generic, we started doing some research, and encountered OpenFeature. Gladly, the mentioned feature management cloud service has already created an OpenFeature provider. Let me explain why and how it was adopted. - read more details [here](https://openfeature.dev/blog/openfeature-adoption-story)

### **_Clear versioning scheme & release methodology_**

As outlined by our [Repository requirements](https://github.com/open-feature/.github/blob/main/CONTRIBUTING.md#repository-requirements), OpenFeature artifacts adhere to semantic versioning and include meaningful change logs. The OpenFeature specification includes [Document status](https://github.com/open-feature/spec/tree/main/specification#document-statuses) definitions, which are used to indicate the stability level of each specification section.

### **_Demonstrate a substantial ongoing flow of commits and merged contributions._**
* [Commit Activity](https://openfeature.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-1y&to=now&var-period=w&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) 
    * In June 2023 there were a total of 820 contributions from 29 contributors. Up from a total of 64 contributions by 6 contributors the year before.
* [New Contributors over past year](https://openfeature.devstats.cncf.io/d/52/new-contributors-table)
    * On average we have 2 new contributors joining each week
* There are on average 30 new PRs on the project per week
* In the last 30 days the average time to close PRs was 14 hours 4 minutes

Full details can be found at [CNCF devstats](https://openfeature.devstats.cncf.io/) dashboards for OpenFeature.

### **_Security_**

Security processes are documented [here](https://github.com/open-feature/community/security/policy)

## Statement on Alignment with the CNCF Mission

OpenFeature provides a standardized, vendor-neutral specification and tooling for feature flagging.
Feature flags are becoming an essential aspect of the modern software development lifecycle because they allow teams to make decouple deployments from feature releases.
This reduces many risks associated with traditional deployments, allowing engineers to make high-impact changes frequently and predictably.

Standardizing feature flags unlocks increased observability at the feature level, providing additional pathways for increased automation, including intelligent deployment and rollback orchestration.
This aligns with the [CNCF mission](https://github.com/cncf/foundation/blob/main/charter.md#1-mission-of-the-cloud-native-computing-foundation) of democratizing state-of-the-art patterns, making innovations in feature flagging, feature observability, and intelligent deployments accessible to everyone.

### Alignment with Other CNCF Projects

Contributions to other CNCF projects:

- [Feature Flag Semantic Contention in OpenTelemetry](https://opentelemetry.io/docs/reference/specification/trace/semantic_conventions/feature-flags/)

CNCF project used by OpenFeature:

- gRPC
- Helm
- Kubernetes
- Cert Manager
- Operator Framework
- OpenTelemetry
- Artifact Hub

## Future Plans

A world where feature flagging is a core principle of the software development lifecycle, enabling teams to test and release high-quality software with a high degree of safety and confidence.

OpenFeature can achieve this vision by:

**Vendor neutrality**

OpenFeature aims to provide a standardized interface for an application author to use feature flags in a vendor-neutral way. This provides flexibility for the application integrators to choose the tooling that best fits their current requirements while avoiding code-level lock-in.

The vendor-neutral ecosystem enables additional integration options for framework and library authors. Using the standardized interface, they can include feature flagging natively, allowing users to choose the provider that best fits their needs.

**Extensibility**

OpenFeature is extensible while still being easy to work with. It is quick to get started with while providing application authors the flexibility to customize the functionality based on their business requirements and workflows.

**Community**

OpenFeature is a collective effort that benefits from years of experience across multiple verticals, including many of the top feature flag vendors and practitioners. The patterns, idioms, and techniques developed by this collaboration support feature flagging as a core principle of the software development lifecycle.

**Education**

Feature flagging is a simple yet powerful mechanism that improves the entire software development lifecycle by decoupling feature releases from a deployment. It isn't, however, ubiquitously used throughout the industry. OpenFeature is an opportunity to promote good software practices in a vendor-neutral way through developing feature flag awareness.

The direction of travel for OpenFeature is highly influenced by community contributions, and the core maintainers will continue to support the requests, as it has been until today. Saying that, the clear commitment is to continue to be “community-driven”, scaling the community support accordingly with the increasing number of adopters, requests of support (in forums, chats, github issues, etc.) and contributions.

As a tangible plan for the enhancements developed by the core maintainers, but not only by them, you can check the OpenFeature [roadmap](https://github.com/orgs/open-feature/projects/1)
