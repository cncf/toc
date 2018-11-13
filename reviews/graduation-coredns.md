
# CoreDNS Graduation Application

CoreDNS joined CNCF in March 2017 as inception project and graduated as incubation project on February 2018

It has grown significantly over time either as a DNS Server available in Cloud Native environment, either as a key component of Kubernetes project:
- In Fall 2017, CoreDNS was proposed to become a DNS Discovery Service for Kubernetes.
- First version as alpha feature was released in v1.09 of Kubernetes - Dec 2017
- CoreDNS is considered the “default” DNS Discovery Service for Kubernetes v1.13 - to be released in Dec 2018

The following application links to the required information to become a graduated project.

## CoreDNS fulfills all the incubating and graduation criteria:

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

* Bose, Infoblox, Zalando, Tradeshift, Yandex, AdGuard are knows users of CoreDNS. More adopters are listed [here](https://github.com/coredns/coredns/blob/master/ADOPTERS.md).

* On Kubernetes side, from [a recent CNCF Survey](https://docs.google.com/forms/d/e/1FAIpQLSehsvd0n2HKnPEifukNHej4gMBrXe2saJOmxAmPq34dTGMWKA/viewanalytics), 50% deployments of cluster are using CoreDNS as the DNS discovery service


### Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

CoreDNS has currenlty 16 maintainers.
They are listed as ‘approvers’ on our [OWNER document](https://github.com/coredns/coredns/blob/master/OWNERS)

### Demonstrate a substantial ongoing flow of commits and merged contributions.

More than [100+ contributors](https://github.com/coredns/coredns/graphs/contributors) and 12 [releases](https://coredns.devstats.cncf.io/d/3/community-stats?orgId=1&var-period=d&var-repo_name=CoreDNS) since incubation graduation

### Have committers from at least two organizations.

We currently have [16 maintainers](https://github.com/coredns/coredns/blob/master/OWNERS).
Most of them are contributing on their spare time and each considered at its own organization,
4 maintainers are working for 2 distinct companies - Infoblox, Google - that support CoreDNS.

### Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

https://bestpractices.coreinfrastructure.org/en/projects/1250

### Adopt the CNCF Code of Conduct.

https://github.com/coredns/coredns/blob/master/CODE-OF-CONDUCT.md

### Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

https://github.com/coredns/coredns/blob/master/GOVERNANCE.md

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

https://github.com/coredns/coredns/blob/master/ADOPTERS.md
