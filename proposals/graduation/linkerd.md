# Linkerd Graduation Proposal

[Linkerd](https://linkerd.io) is an ultralight, security-first service mesh for
Kubernetes. Linkerd adds critical security, observability, and reliability
features to your Kubernetes stack with no code change required.

Linkerd features a [control plane](https://github.com/linkerd/linkerd2) built
in Go and a [data plane
"micro-proxy"](https://github.com/linkerd/linkerd2-proxy) written in Rust. The
original [1.x line](https://github.com/linkerd/linkerd) of Linkerd, written in
Scala, is currently in maintenance mode.

Linkerd has 10,000+ GitHub stars and 150+ contributors across the various repos
in the [Linkerd GitHub org](https://github.com/linkerd/), and 5000+ users
registered for the [Linkerd community Slack](https://slack.linkerd.io/).
Linkerd was [open sourced in early
2016](https://linkerd.io/2016/02/18/linkerd-twitter-style-operability-for-microservices/)
and [joined the CNCF as its fifth-ever project in January
2017](https://www.cncf.io/blog/2017/01/23/linkerd-project-joins-cloud-native-computing-foundation/).
It currently powers production systems around the world, including at
Nordstrom, H-E-B, Elkjop, Microsoft, HP, and many more. 

## Graduation State Criteria

### Have committers from at least two organizations.

Currently, all of Linkerd's active maintainers are employed by the same
organization ([Buoyant](https://buoyant.io)). However, in the past, Linkerd has
had maintainers from other organizations.

Throughout its history, Linkerd has always been open to contributors and
maintainers from anywhere in the world. Linkerd has [publicly committed to open
governance](https://linkerd.io/2019/10/03/linkerds-commitment-to-open-governance/)
and features a [steering committee](https://linkerd.io/2021/01/28/announcing-the-linkerd-steering-committee/) comprising end-users.

### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/).

The Linkerd project has achieved a [Core Infrastructure Initiative Best Practices Badge](https://bestpractices.coreinfrastructure.org/en/projects/4629).

### Have completed an independent and third party security audit with results published of similar scope and quality as [this example](https://github.com/envoyproxy/envoy#security-audit) which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

Linkerd completed a CNCF-sponsored [Cure53 security
audit](https://github.com/linkerd/linkerd2/blob/main/SECURITY_AUDIT.pdf) in
June 2019. The audit identified two minor issues in the dashboard, both of
which were promptly fixed.

A second CNCF-sponsored audit by Trail of Bits is scheduled to commence next
month.

A CNCF-sponsored fuzzing audit from ADA Logics is currently underway.

### Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

* Project governance:
  [GOVERNANCE.md](https://github.com/linkerd/linkerd2/blob/main/GOVERNANCE.md)
* Steering committee charter:
  [STEERING.md](https://github.com/linkerd/linkerd2/blob/main/STEERING.md)
* Maintainers and Steering Committee members:
  [MAINTAINERS.md](https://github.com/linkerd/linkerd2/blob/main/MAINTAINERS.md)

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.

See [ADOPTERS.md](https://github.com/linkerd/linkerd2/blob/main/ADOPTERS.md).

## Incubation Details

Linkerd [entered incubation in April 2018](https://github.com/cncf/toc/pull/86).

