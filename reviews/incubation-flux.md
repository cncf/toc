# Flux proposal for CNCF incubation

## Background

 - [Flux proposal for CNCF Sandbox (2019)][flux-sandbox]
 - [Flux proposal for CNCF Sandbox presentation
   (2019)][flux-sandbox-preso]
 - [Flux in CNCF annual review (2020)][flux-annual]

The Flux project aims to provide a complete Continuous Delivery
platform on top of Kubernetes, supporting all the common practices and
tooling in the field, e.g. Helm, kustomize, metrics with Prometheus
and so on.

Flux was the first project to clearly define [GitOps
practices][gitops-defn], which can be summarised as:

 * Describing your system declaratively;
 * Keeping configuration under source control; and,
 * Using software agents to reconcile along with alerts, etc., to and
   ensure correctness.

Flux entered the CNCF Sandbox on August 15th, 2019.

## Progress since joining the CNCF Sandbox

In short, the Flux project has seen terrific growth and project
participation since entering the CNCF Sandbox. The project has matured
in many ways and is looking towards an exciting new phase of the
project.

Here are some highlights:

 * Adoption of Flux has increased significantly. Public end users are
   documented in the [project's README.md][flux-prod-users].
 * The CNCF User Community put Flux alongside Helm in the ADOPT
   category of their [Technology Radar][cncf-cd-radar].
 * Mettle, Fidelity Investments, and Tidepool wrote case-studies with
   the Weaveworks Marketing team (links below).
 * All relevant numbers in terms of users of Flux and Helm Operator
   (v1) have gone up by a factor around 2x to 4x (production users,
   users on Slack, mailing list, etc.)
 * Contributions have gone up in a similar fashion.
 * Releases have happened in a regular cadence. Our v1 projects now
   have more documentation and have received a few more features (more
   below) and many stability improvements.
 * Our community has matured significantly. We have more project
   maintainers on board, and we run weekly public meetings. We
   codified our [governance processes][flux-governance], [contributing
   guidelines][flux-contributing], and [security
   processes][flux-security].

### Transition to Flux version 2

In the last couple of years, the feedback the Flux team got was
overwhelmingly positive, but it also became clear that some of the
most requested features were going to be hard to implement.

This was partly due to Flux being developed at a time when the
ecosystem was more varied and it seemed prudent to hedge with respect
to cluster technology (June 2016), leading to some architectural
dead-ends. And even with respect to Kubernetes, the tooling ecosystem
(controller-runtime, kyaml) we enjoy now was simply not available.

At the end of 2019 we partnered with the Argo CD team at Intuit and
AWS and looked at building a more up-to-date answer to GitOps
together. After thorough deliberation and experimentation Argo and
Flux decided to [go separate ways again][flux-v-argo]. In short:
instead of rebasing Flux on top of core code of Argo, the Flux team
chose to rebuild Flux and Helm Operator from scratch on top of
targeted controllers, so every part used modern tooling (i.e., using
[controller-runtime][controller-runtime]), had a narrow focus and
could be understood, tested and built easily. This also started to
reframe Flux as an organisation to become a Flux family of GitOps
projects.

Timeline for Flux v2:

 * April 2020: Experimentation following the new paradigm started.
 * May & June 2020: first demos of PoC.
 * June & July 2020: Discussion within Flux community and agreement on
   roadmap for Flux v2. Announcement.
 * October 2020: new Flux (in read-only mode) and Helm Operator have
   reached feature parity. Image update functionality is about 30%
   done.
 * ([Prospectively, in Nov 2020][flagger-proposal]) the Flagger
   project joins Flux

The Flux community defined a [roadmap for v2][fluxv2-roadmap]
together. It includes a window for migrating from Flux v1 to Flux v2.

## Requirements for incubation

_Document that it is being used successfully in production by at least
three independent end users which, in the TOC’s judgement, are of
adequate quality and scope._

Many production users have added themselves to the list kept in
[Flux's README][flux-prod-users], among them:

 * [Control Plane](https://control-plane.io/)
 * [IBM Cloudant](https://www.ibm.com/cloud/cloudant)
 * [Rakuten](https://www.rakuten.com/)

For more in-depth descriptions of how Flux is used, we can refer to
case studies put together by Weaveworks:

 * [Mettle leverages GitOps for self-service developer platform][mettle-case-study]
 * [Fidelity Investments Adopts GitOps Automation for Multicloud Kubernetes][fidelity-case-study]
 * [Tidepool streamlines and secures cluster lifecycle with GitOps][tidepool-case-study]

It's also worth noting that Flux appeared in the "ADOPT" segment of
the [CNCF radar for Continuous Delivery][cncf-cd-radar], as judged by
the CNCF end-user community.

_Have a healthy number of committers. A committer is defined as
someone with the commit bit; i.e., someone who can accept
contributions to some or all of the project._

At the time of writing, the Flux project has twelve people who are a
maintainer of one or more subprojects. Many of the maintainers work at
[Weaveworks](https://weave.works/).

Within the Flux v2 subprojects the core team is Hidde, Michael, Stefan
(all Weaveworks), Philip (Xenit), and Sean (AT&T).

_Demonstrate a substantial ongoing flow of commits and merged
contributions._

Different subprojects within Flux have different levels of activity;
the [most active have tens of commits a
week][devstats-commits-weekly].

Across all Flux subprojects, [new PRs are in the tens per
week][devstats-prs-weekly]; and similarly, [merged
PRs][devstats-merges-weekly] are in the tens per week.

_A clear versioning scheme._

Flux v1 is now up to [version 1.21.0][flux1-changelog], and follows
semver semantics. Since 1.0, minor versions have included new
features, while patch versions include only bugfixes and quality of
life improvements; and in all` >=1.0` versions, backward-incompatible
behaviour is guarded by flags.

The Flux v2 subprojects are pre-1.0, and have versioned releases. The
occasional, necessary breaking changes are batched into minor version
releases.

_Specifications must have at least one public reference
implementation._

Not applicable to Flux.


[flux-sandbox]: https://github.com/cncf/toc/pull/232
[flux-sandbox-preso]: https://docs.google.com/presentation/d/1uuXpXkFvRET2DLuvDqiOto0mgkOc8CquT8IvAbDYbig/edit?usp=sharing
[flux-annual]: https://github.com/cncf/toc/pull/511
[gitops-defn]: https://www.gitops.tech/
[flux-prod-users]: https://github.com/fluxcd/flux/blob/master/README.md#who-is-using-flux-in-production
[cncf-cd-radar]: https://www.cncf.io/blog/2020/06/12/introducing-the-cncf-technology-radar/
[flux-governance]: https://github.com/fluxcd/community/pull/13
[flux-contributing]: https://github.com/fluxcd/community/pull/29
[flux-security]: https://github.com/fluxcd/community/pull/31
[flux-v-argo]: https://github.com/argoproj/gitops-engine/blob/master/docs/faq.md
[controller-runtime]: https://github.com/kubernetes-sigs/controller-runtime
[fluxv2-roadmap]: https://toolkit.fluxcd.io/roadmap/
[mettle-case-study]: https://www.weave.works/blog/case-study-mettle-leverages-gitops-for-self-service-developer-platform
[fidelity-case-study]: https://www.weave.works/blog/case-study-fidelity-investments-adopts-gitops-automation-for-multicloud-kubernetes
[tidepool-case-study]: https://www.weave.works/blog/case-study-tidepool-streamlines-and-secures-cluster-lifecycle-with-gitops
[devstats-commits-weekly]: https://flux.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&var-period=w&var-repogroups=All&from=now-90d&to=now
[devstats-prs-weekly]: https://flux.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-90d&to=now
[devstats-merges-weekly]: https://flux.devstats.cncf.io/d/24/prs-merged-repository-groups?orgId=1&var-period=w&var-repogroups=All
[flux1-changelog]: https://github.com/fluxcd/flux/blob/master/CHANGELOG.md#1210-2020-10-01
[flagger-proposal]: https://github.com/fluxcd/community/issues/34
