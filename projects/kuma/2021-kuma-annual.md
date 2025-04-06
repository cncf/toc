# Kuma Sandbox Annual Review

## Table of Contents

- [Background](#background)
- [Alignment with Cloud Native](#alignment-with-cloud-native)
- [Year in Review](#year-in-review)
- [Annual Review Contents](#annual-review-contents)
- [Project Links](#project-links)

## Background

[Kuma](https://kuma.io) is a modern control plane for Service Mesh and Microservices. It can run and be operated natively across both Kubernetes and VM environments, making it easy to adopt by every team in the organization with the goal of accelerating both Kubernetes and Service Mesh adoption.

By bundling the Envoy proxy as the underlying data-plane technology, Kuma can instrument any L4/L7 traffic to secure, observe, route and enhance connectivity between any service or database. It can be used natively in Kubernetes via CRDs, while at the same time providing a RESTful API, a native CLI tool and a built-in GUI that can be used to better integrate Kuma with the rest of the organization.

While Kuma provides easy to use Policy abstractions for most use-cases, Kuma also allows for the configuration of the underlying Envoy data-planes in a more fine-grained manner via the `ProxyTemplate` policy. By doing so, Kuma can be used by both first-time users of Service Mesh, as well as the most experienced ones who want greater control of the underlying networking stack.


**Kuma was accepted as a CNCF Sandbox project on June, 2020.**
- [Sandbox Proposal](https://github.com/cncf/toc/blob/main/proposals/sandbox/kuma.md)
- Kuma falls in the scope of [CNCF Network tag](https://github.com/cncf/tag-network).


## Year in Review

- Since June 2020 Moved from 0.6.0 to 1.4.0
- Version 1.0.0 released 24/11/2020
- 28 releases in total since incubation.
- 1759+ commits; 1.8k+ stars; 400+ forks; 220+ contributors
- Key Features Added: externalService, built in DNS, GUI refresh, IPv6 support, permissive mTLS... 

## Annual Review Contents

### Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

  - [New PRs last 1 year](https://kuma.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1586448000000&to=1622476799000&var-period=w&var-repogroup_name=All)

  - [Commits Repository Groups last 1 year](https://kuma.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=w&var-repogroups=All)

  - Community growth
    - Entered CNCF sandbox in June 2020
    - Continuous momentum
      - Num of Contributors: 18 => **78+**
      - Github Stars: 1400+ => **2616+**
      - Github Forks: 173+ => **297+**

### How many maintainers do you have, and which organizations are they from? [OWNERS.md](https://github.com/kumahq/kuma/blob/master/OWNERS.md)

  Initial maintainers
  * Jakub Dyszkiewicz ([jakubdyszkiewicz](https://github.com/jakubdyszkiewicz)) (jakub.dyszkiewicz@konghq.com) Kong
  * Ilya Lobkov ([lobkovilya](https://github.com/lobkovilya)) (ilya.lobkov@konghq.com) Kong
  * Marco Palladino ([subnetmarco](https://github.com/subnetmarco)) (marco@konghq.com) Kong

  New in the Year

  * Austin Cawley-Edwards([austince](https://github.com/austince)) (austin.cawley@gmail.com) Ververica
  * Bart Smykla ([bartsmykla](https://github.com/bartsmykla)) (bartek@smykla.com) Kong
  * Charly Molter ([lahabana](https://github.com/lahabana)) (charly@molter.io) Kong
  * James Peach ([jpeach](https://github.com/jpeach)) (james.peach@konghq.com) Kong
  * Mike Beaumont ([michaelbeaumont](https://github.com/michaelbeaumont)) (michael.beaumont@konghq.com) Kong
  * Nikolay Nikolaev ([nickolaev](https://github.com/nickolaev)) (nicknickolaev@gmail.com) Juniper Networks 
  * Paul Parkanzky ([parkanzky](https://github.com/parkanzky)) (paul.parkanzky@konghq.com) Kong


### What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

While it's hard to track the global usage (especially without telemetry) we know that Kuma is being used in production across 40+ large enterprise organizations worldwide that decided to engage with Kong (and counting), and has been used by approximately 1000+ organizations worldwide in the last 12 months, including Sabre, American Airlines, and several large financial institutions.

### How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

We've released a total of 28 releases.

At sandbox time goals included:

* Support for L7 HTTP routing policies (in addition to the existing L4 routing policy)
* Support for L7 gRPC routing policies (in addition to the existing L4 routing policy)
* Discussing roadmap for SMI integration (most requested feature on GitHub)
* Introducing pluggable backends for mTLS, logging and tracing policies
* Adding support for multi-cluster Kubernetes deployments
* Support for traffic shadowing and auto-retries
* Support for L7 HTTP traffic permissions (in addition to the existing L4 traffic permission policy)
* Support for L7 gRPC traffic permissions (in addition to the existing L4 traffic permission policy)
* Support for L7 gRPC logging (in addition to the existing L4 traffic log policy)
* GUI wizards for every entity (in addition to Mesh and Dataplane entities)

What we've shipped since:

* Newest GUI with support for all entities
* ExternalService support
* Big performance improvements
* Transparent Proxy for non k8s use-cases
* IPv6 support
* mTLS permissive mode to ease migration to Kuma
* Some policy support for gRPC and HTTP
* DNS Proxy with configurable DNS entries to simplify discovery in the serviceMesh
* Migrated all issues to github issues and setup public triage meeting to provide better visibility to the community.
* Native integration with [prometheus](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#kuma_sd_config) to make Kuma even more cloud native.

### What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

  - Technical
    - Performance improvements and qualification (we want to demonstrate that Kuma scales to O(10K) dataplanes and O(100) zones)
    - Rewrite of policy selector to simplify even more Kuma's UX and unblock some features (some L7 policies are currently blocked by this)
    - Simplify deployment topology for multi-zone setups.

  - Community
    - Further, grow the number of contributors
    - Move to a more predictable versioning policy 
    - Define a security process


### How can the CNCF help you achieve your upcoming goals?

  - We need more speaking and marketing opportunities to help attract more contributors and user adoptions.
  - We would like to reintroduce a way to follow feature usage and installation count as it's hard to track adoption without this.
  - We could use with some help with some technical writing.

### Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

No for the following reasons:

  - We would like to be 2.0 before incubation. 
  - We still need to clarify our release policy and our security process. 

## Project Links
 - [Website](https://kuma.io/)
 - [Github](https://github.com/kumahq/kuma)
 - [Slack](https://chat.kuma.io)
 - [Twitter](https://twitter.com/kumamesh)
 - [Issue Tracker](https://github.com/kumahq/kuma/issues)
