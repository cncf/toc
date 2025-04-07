# etcd Graduation Proposal

etcd is a strongly consistent, distributed key-value store that provides a reliable way to store data that needs to be accessed by a distributed system like Kubernetes. A typical etcd cluster is distributed over 3 to 5 nodes for high availability, while it prioritizes consistency and partition tolerance. Which means, it provides one logical cluster view of many physical servers. So long as a majority of the cluster is up, etcd continues to work even under machine failures. This redundancy provides fault tolerance.

etcd server implements Raft consensus algorithm for data replication. Raft is a leader-based protocol. Data is replicated from leader to follower; follower forwards proposals to leader, and leader decides what to commit or not. Leader persists and replicates an entry once it has been agreed by the quorum of the cluster. The underlying storage layer for Raft log is write-ahead log (WAL). Committed entries are written out to disk, so they can be replayed on restart. etcd uses gRPC for transport layer.

The [etcd project](https://github.com/etcd-io/etcd) was open sourced in July 2014 and was accepted into the CNCF as an incubating project in August 2018. The project has grown significantly over time. Jepsen, a leading company that tests open source distributed systems to check if they fulfill their consistency guarantee, tested and analyzed etcd major release 3.4 (v3.4.3) in January 2020. The [testing results](https://etcd.io/blog/jepsen-343-results/) were favorable showing the etcd project functionality maturity.

## Graduation State Criteria
_**Project should address each graduation criteria listed below**_

### * Have committers from at least two organizations.

The etcd project currently has [10 committers from more than 5 companies](https://github.com/etcd-io/etcd/blob/master/MAINTAINERS) including Alibaba, Amazon, Cockroach Labs, Google, IBM, Indeed, Red Hat, and an individual contributor.

### * Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge/(https://bestpractices.coreinfrastructure.org/).

The etcd project has achieved a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/en/projects/3192).

### * Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

The etcd project has completed a [3rd party](https://www.trailofbits.com/) security review in July 2020. The [review results](https://github.com/etcd-io/etcd/blob/master/security/SECURITY_AUDIT.pdf) are available on the etcd GitHub repository and announced publicly through a [CNCF blog post](https://www.cncf.io/blog/2020/08/05/etcd-security-audit/). There was no major issue found in the core components of etcd. According to the report summary, overall, the etcd codebase represents a mature and heavily adopted product. The reported issues have been addressed and backported to the supported versions of etcd.

### * Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

The etcd project has a [documented governance](https://github.com/etcd-io/etcd/blob/master/GOVERNANCE.md) that outlines various things including roles, responsibilities, how decisions are made, how [maintainers](https://github.com/etcd-io/etcd/blob/master/MAINTAINERS) are elected and removed, and more.

### * Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.

The [ADOPTERS.md](https://github.com/etcd-io/etcd/blob/master/ADOPTERS.md ) file shows the list of etcd project adopters.

## Incubation Details
_**Project should address each area listed below**_

### * Link to Incubation Due Diligence(DD) Document

The etcd project [incubation proposal](https://github.com/cncf/toc/pull/143) has described etcd in great detail. A newly created [Due Diligence](https://docs.google.com/document/d/10IRk__v_nehw-0BpqUnNSY4A8RNP2ztdqKLX2mh0PlU/edit?usp=sharing) document describes etcd in more detail with the recent developments.

### * Address any concerns or recommendations from the SIG and/or TOC sponsor(s) from the DD Document

There are currently no known outstanding concerns or recommendations with the etcd project.
