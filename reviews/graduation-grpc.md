# **gRPC Graduation Application**


[gRPC](https://grpc.io/) was released as an open source project in Feb 2015. In August 2016, gRPC reached 1.0 milestone. gRPC joined the CNCF with the 1.0 release. The project and its adoption has grown significantly over time.

*   45,000+ Github stars
*   10,000+ forks
*   900+ contributors

The following application provides the required information for gRPC to be recognized as a graduated project in CNCF as per the graduation criteria listed [here](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc).


## **CNCF Graduation Criteria**


### 
**Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.**



*   gRPC is adopted widely. Some end users are: Netflix, Google, Dropbox, Slack, GO-JEK, Lyft, Uber, Medium, Salesforce, VSCO, Square, Spotify, CoreOS, CockroachLabs and Microsoft
    *   Quotes from some users can be seen [here](https://grpc.io/about/). 
    *   [Public blogs](https://www.google.com/search?q=grpc+adoption&safe=strict) from many adopters can be found on the Internet.
*   In addition to microservice adoption, gRPC is being used to develop API endpoints for products such as Envoy, NGINX, HAProxy, Dgraph.io, CockroachDB and Lightning Network.
*   Several other CNCF projects use gRPC, such as Etcd, Containerd, Vitess and Jaeger.
*   gRPC has a vibrant ecosystem of tooling, reflecting it’s adoption and use.
    *   https://github.com/grpc-ecosystem/awesome-grpc

### 
**Have a healthy number of committers.**


gRPC has a large community of active committers:



*   https://github.com/grpc/grpc/graphs/contributors
*   https://github.com/grpc/grpc-java/graphs/contributors
*   https://github.com/grpc/grpc-go/graphs/contributors
*   https://github.com/grpc/grpc-node/graphs/contributors
*   https://github.com/grpc/grpc-web/graphs/contributors
*   https://github.com/grpc/grpc-swift/graphs/contributors
*   https://github.com/grpc/grpc-dotnet/graphs/contributors
*   https://grpc.devstats.cncf.io/d/48/users-statistics-by-repository-group?orgId=1

### 
**Demonstrate a substantial ongoing flow of commits and merged contributions.**

*   New PRs: https://grpc.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&var-period=d7&var-repogroup_name=grpc
*   Merged PRs: https://grpc.devstats.cncf.io/d/24/prs-merged-repository-groups?orgId=1

### 
**A clear versioning scheme**


gRPC strictly follows semantic versioning scheme. Details are [here](https://github.com/grpc/grpc/blob/master/doc/versioning.md). gRPC versions are released every six weeks. The release schedule is [here](https://github.com/grpc/grpc/blob/57bbe93e404fe59666abd69d46f51a2cee57ef01/doc/grpc_release_schedule.md). Past releases:



*   https://github.com/grpc/grpc/releases (C++, Python, Ruby, C#, PHP, ObjectiveC)
*   https://github.com/grpc/grpc-java/releases (Java)
*   https://github.com/grpc/grpc-go/releases (Go)
*   https://github.com/grpc/grpc-swift/releases (Swift)
*   https://github.com/grpc/grpc-node/releases (Node)
*   https://github.com/grpc/grpc-dotnet/releases (DotNet)

### 
**Specifications must have at least one public reference implementation**


gRPC has public implementations in many languages:



*   https://github.com/grpc/grpc (C++, Python, Ruby, C#, PHP, ObjectiveC)
*   https://github.com/grpc/grpc-java (Java)
*   https://github.com/grpc/grpc-go (Go)
*   https://github.com/grpc/grpc-swift(Swift)
*   https://github.com/grpc/grpc-node (Node)
*   https://github.com/grpc/grpc-dotnet 

### 
**Have committers from at least two organizations.**

*   Committers from Google, Dropox and LinkedIn in [gRPC](https://github.com/grpc/grpc/blob/master/MAINTAINERS.md) repo (C++, Python, Ruby, C#, PHP, ObjectiveC).
*   Committers from Google and Microsoft in [grpc-dotnet](https://github.com/grpc/grpc-dotnet/blob/master/MAINTAINERS.md) repo. 
*   Committers from Apple, Timing, Lyft and Google in [grpc-swift](https://github.com/grpc/grpc-swift/blob/master/MAINTAINERS.md) repo.

### 
**Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.**


https://bestpractices.coreinfrastructure.org/en/projects/2112


### 
**Have completed an independent and third party security audit with results published of similar scope and quality as the following example (including critical vulnerabilities addressed): https://github.com/envoyproxy/envoy#security-audit and all critical vulnerabilities need to be addressed before graduation.**

Third party audit for gRPC C++ implementation in https://github.com/grpc/grpc is underway. We will update as soon as we have the report and all critical vulnerabilities, if any, are addressed.


### 
**Adopt the CNCF Code of Conduct.**

https://github.com/grpc/grpc-community/blob/master/governance.md#code-of-conduct


### 
**Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.**



*   Governance: https://github.com/grpc/grpc-community/blob/master/governance.md
*   Contribution guidelines: https://github.com/grpc/grpc-community/blob/master/CONTRIBUTING.md

### 
**Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).**

*   A small subset of adopters are listed on gRPC website https://grpc.io/about/
*   [Public blogs](https://www.google.com/search?q=grpc+adoption&safe=strict) from many adopters can be found on the Internet.

