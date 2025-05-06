# Kuma Sandbox Annual Review

## Table of Contents

- [Background](#background)
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

- Since In 2022 Moved from 1.4.0 to 2.2.1
- Version 2.0.0 released 4/11/2022
- 13 releases in 2022 (1 major, 3 minors, 9 patches)
- 4000+ commits; 2.3k+ stars; ~300 forks; 35 new first time contributors
- Key Features Added: eBPF, new policy API, GatewayAPI, zone Egress... 

## Annual Review Contents

### Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

  - [New PRs last 1 year](https://kuma.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1586448000000&to=1622476799000&var-period=w&var-repogroup_name=All)

  - [Commits Repository Groups last 1 year](https://kuma.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=1640991600000&to=1672441200000&var-period=w&var-repogroups=All)

  - [Closed by Repository Groups last 2 years](https://kuma.devstats.cncf.io/d/12/issues-opened-closed-by-repository-group?orgId=1&from=1609455600000&to=1672527599000) -- We can see a strong increase here. Number of commits is stagnating mostly because of the project gaining stability (less big features requiring a lot of code change).

  - [Opened to Merge](https://kuma.devstats.cncf.io/d/16/opened-to-merged?orgId=1&from=now-2y&to=now-1w) -- This remains very low, which indicates strong reactivity to contributions.

  - Community growth
    - Entered CNCF sandbox in June 2020
    - Continuous momentum
      - Num of Contributors: 18 => **80+**
      - Github Stars: 1400+ => **3170+**
      - Github Forks: 173+ => **294+**

### How many maintainers do you have, and which organizations are they from? [OWNERS.md](https://github.com/kumahq/kuma/blob/master/OWNERS.md)

  Initial maintainers
  * Jakub Dyszkiewicz ([jakubdyszkiewicz](https://github.com/jakubdyszkiewicz)) (jakub.dyszkiewicz@konghq.com) Kong
  * Ilya Lobkov ([lobkovilya](https://github.com/lobkovilya)) (ilya.lobkov@konghq.com) Kong
  * Marco Palladino ([subnetmarco](https://github.com/subnetmarco)) (marco@konghq.com) Kong

  New in the Year

  * Mike Beaumont ([michaelbeaumont](https://github.com/michaelbeaumont)) (mjboamail@gmail.com) Kong
  * Philipp Rudloff ([kleingfreund](https://github.com/kleinfreund)) (philipp.rudloff@konghq.com) Kong

  Total maintainers 16

Note that as part of our plans for incubations we'd like to review and move some of the maintainers to an "alumni" group.

### What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

While it's hard to track the global usage we know that Kuma is being used in production across 40+ large enterprise organizations worldwide that decided to engage with Kong (and counting), and has been used by approximately 1000+ organizations worldwide in the last 12 months, including Sabre, American Airlines, Maersk, Koyeb and several large financial institutions.
We're in the process of adding an ADOPTERS.md file soon to better show this.

### How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

The goals of Kuma remain building a simple service mesh that doesn't only work for Kubernetes.

We're still very much aligned with this goal.
In 2022 we've shipped features like:

- ZoneEgress, enabling users to funnel traffic through some hosts, to accommodate restrictive networking setups
- New Policy API making our policies more flexible and easier to use.
- Support for Gateway API
- Native Gateway avoiding the need to run a 3rd party gateway
- Fully rewrote transparent proxy to accommodate more non kubernetes use cases
- Move to a predictable versioning policy which frequent patches especially for security.

### What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

  - Technical
    - Performance improvements and qualification (we want to demonstrate that Kuma scales to O(10K) dataplanes and O(100) zones)
    - Make our new policy API default
    - Simplify deployment topology for multi-zone setups. With the possibility to run Global and Zone in the same kubernetes cluster.

  - Community
    - Further, grow the number of contributors
    - [Review governance](https://github.com/kumahq/kuma/pull/6920)
    - Move to incubation

### How can the CNCF help you achieve your upcoming goals?

  - We would like to move incubation and will need support from the CNCF to work through this process

### Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

Yes we will likely apply for incubation in the coming weeks.

## Project Links
 - [Website](https://kuma.io/)
 - [Github](https://github.com/kumahq/kuma)
 - [Slack](https://chat.kuma.io)
 - [Twitter](https://twitter.com/kumamesh)
 - [Issue Tracker](https://github.com/kumahq/kuma/issues)
