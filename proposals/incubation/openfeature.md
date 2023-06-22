# Propose OpenFeature to Incubation Stage

## Introduction

- Name: OpenFeature
- Website: [openfeature.dev](https://openfeature.dev)
- Source Control: [OpenFeature on Github](https://github.com/open-feature)
- Code of Conduct: [OpenFeature Code of conduct](https://github.com/open-feature/community#code-of-conduct)
- License: [Apache 2](https://github.com/backstage/backstage/blob/master/LICENSE)

## Overview

OpenFeature is an open standard for feature flag management, created to support a robust feature flag ecosystem using cloud native technologies. OpenFeature provides a unified API and SDK, and a developer-first, cloud-native implementation, with extensibility for open source and commercial offerings.

## Sponsors from the TOC

The OpenFeature project is currently looking for a TOC sponsor

## Progress Since Sandbox

### Community

**Outreach:** As a project we hold regular community meetings each month, highlighting what is happening in the community with presentations from multiple adopters and discussions on current topics and interests in the community. Community sessions have on average **12 attendees**, are uploaded to a [Project meeting playlist hosted on Youtube](<[https://www.youtube.com/watch?v=3dV3aZo6JN8&list=PLf1KFlSkDLIBmA5TLXn2BzEHmwWzckP8y](https://www.youtube.com/playlist?list=PLiQt8D1ofl8zR0u5kzFjrX_U4rGteyOVa)>) and gained **over 750 views**.

- [OpenFeature Community Hub on Github](https://github.com/open-feature/community)
- [OpenFeature Community on Slack](https://cloud-native.slack.com/archives/C0344AANLA1)

The community statistics are extracted from [CNCF devstats](https://openfeature.devstats.cncf.io/).

### Technical

#### SDK Progress
Since achieveing sandbox we have progressed in the development of our SDKs:
- [.NET](https://github.com/open-feature/dotnet-sdk) released v1.1.0
- [Java](https://github.com/open-feature/java-sdk) released v1.2.0
- [PHP](https://github.com/open-feature/php-sdk) released v2.0.0
- [Python](https://github.com/open-feature/python-sdk) released v0.0.9
- [JavaScript](https://github.com/open-feature/js-sdk) released v1.1.0
- [GO](https://github.com/open-feature/go-sdk) released v1.2.0
- [Ruby](https://github.com/open-feature/ruby-sdk) released v0.1.0

#### Introduction of flagD
[Released v0.3.7](https://github.com/open-feature/flagd)
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
- .NET: Cloudbees, fladg, Go feature flag, LaunchDarkly, Split
- GO: Cloudbees, fladg, Flipt, Go feature flag, Split
- Java: Cloudbees, fladg, Flagsmith, Go feature flag, Split
- JavaScript: Cloudbees, fladg, Go feature flag, LaunchDarkly, PostHog, Split
- PHP: Cloudbees, fladg, Split

## Incubation Stage Requirements

### **_Adhere to CNCF [IP Policy](https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy)_**

As an existing Sandbox project this is already in place

### **_Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope_**
A list of public adopters can be found [here](https://github.com/open-feature/community/blob/main/ADOPTERS.md). There are ~ 10 companies in various stages of adopting, many of whom are non-public in their adoption.
The following are some examples of adopters:

- [Split](https://www.split.io) - Split is excited to announce participation in OpenFeature, an initiative led by Dynatrace and recently submitted to the Cloud Native Computing Foundation (CNCF) for consideration as a sandbox program - Check out the full article [here](https://www.split.io/blog/split-embraces-openfeature/) 
- [Go Feature Flag](https://gofeatureflag.org) - GO Feature Flag believe in OpenSource and offer providers for the new standard for feature flags OpenFeature - check out a more detailed explanation [here](https://gofeatureflag.org/docs/next/openfeature_sdk/concepts)
- [eBay](https://www.ebay.com) - As part of eBay’s tech-led reimagination, we are investing heavily in improving software delivery so that we can launch features more rapidly and safely to our customers. Over the past few months, we have been developing feature flagging capabilities internally, powered by OpenFeature. - read more details [here](https://tech.ebayinc.com/engineering/openfeature-with-contributions-from-ebay-submitted-to-cncfs-sandbox-program/)
- [Schweitzer Engineering Labs](https://selinc.com/)
- [Tapico](https://tapico.io/)
- [Utility Warehouse](https://uw.co.uk/)

### **_Clear versioning scheme & release methodology_**

As outlined by our [Repository requirements](https://github.com/open-feature/.github/blob/main/CONTRIBUTING.md#repository-requirements), OpenFeature artifacts adhere to semantic versioning and include meaningful change logs. The OpenFeature specification includes [Document status](https://github.com/open-feature/spec/tree/main/specification#document-statuses) definitions, which are used to indicate the stability level of each specification section.

### **_Demonstrate a substantial ongoing flow of commits and merged contributions._**
* [Commit Activity](https://openfeature.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-1y&to=now&var-period=w&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all) 
    * In February 2023 there were a total of 691 contributions from 22 contributors. Up from a total of 64 contributions by 6 contributors the year before.
* [New Contributors over past year](https://openfeature.devstats.cncf.io/d/52/new-contributors-table)
    * On average we have 2 new contributors joining each week
* There are on average 30 new PRs on the project per week
* In the last 30 days the average time to close PRs was 14 hours 4 minutes

Full details can be found at [CNCF devstats](https://openfeature.devstats.cncf.io/) dashboards for OpenFeature.

### **_Security_**

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
