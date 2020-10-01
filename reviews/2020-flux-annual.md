# 2020 Flux Annual Review

## Preface

The Flux project aims to provide a complete Continuous Delivery
platform on top of Kubernetes, supporting all the common practices and
tooling in the field, e.g. Helm, kustomize, metrics with Prometheus
and so on.

Flux was the first project to clearly define GitOps practices, which
can be summarised as:

 1. Describing your system declaratively;
 2. Keeping configuration under source control; and
 3. Using software agents to reconcile along
    with alerts, etc., to and ensure correctness.

The original presentation for Flux entering the CNCF Sandbox is
[here](https://docs.google.com/presentation/d/1uuXpXkFvRET2DLuvDqiOto0mgkOc8CquT8IvAbDYbig/edit?usp=sharing).

## What has happened in the last year

### Contributions

Here is the [devstats
page](https://flux.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m)
for Flux. Contributions have been growing fairly steadily over the
year ([devstats
chart](https://flux.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1)).

It's also worth looking at the [commits by
repository](https://flux.devstats.cncf.io/d/2/commits-repository-groups?orgId=1&from=now-1y&to=now&var-period=w&var-repogroups=All)
chart, which shows a change in emphasis. In words --

 * Many commits before landing in Sandbox 2019-08-15 involved
   decoupling from Weave Cloud, splitting out the Helm Operator, etc.
 * In the months before and after CNCF Sandbox we had added features
   and integrations like kustomize, support for encrypted manifests
   with SOPS and lots of other improvements, so in the months after
   that we focused on stabilisation, consolidation, docs, etc.
 * Contributions largely stayed in the same ball-park, with a growing
   proportion of contributions in fluxcd/helm-operator.
 * In April 2020 we started experimentation around a more distributed
   and focused delivery model within the Flux project, called the
   GitOps Toolkit. It uses targeted controllers, for e.g. handling of
   git sources, updating images, kustomization and more. This will
   help us land some long-standing feature requests and be the
   foundation of Flux version 2.

### Maintainers

Our current maintainers are listed in
[fluxcd/flux](https://github.com/fluxcd/flux/blob/master/MAINTAINERS)
and
[fluxcd/helm-operator](https://github.com/fluxcd/helm-operator/blob/master/MAINTAINERS)
-- three out of four working for Weaveworks.

Some of the components of the GitOps Toolkit are seeing bigger and
more steady contributions from people outside Weaveworks and we can
hope to bring some people on board as maintainers of these projects.

### Adoption

We keep a self-selecting list of production users. This was the list
at the time Flux became a CNCF project:

    https://github.com/fluxcd/flux/tree/1.14.0#who-is-using-flux-in-production

.. and this is the list today:

    https://github.com/fluxcd/flux/blob/master/README.md#who-is-using-flux-in-production

(i.e., it has almost tripled). Flux was in the ADOPT sector of the
[CNCF Technology Radar for Continuous
Delivery](https://www.cncf.io/blog/2020/06/12/introducing-the-cncf-technology-radar/),
meaning it was widely adopted and recommended among the CNCF end user
community.

There are some proxy metrics for adoption:

 * the #flux channel in CNCF Slack now has >1100 users;
 * GitHub stars have gone from about 2500 to >5300 in the last year;
 * we believe (based on optional telemetry in fluxd) that the install
   base is on the order of tens of thousands of instances.

### Performance against goals

This is the Flux project's first review, so there are no goals already
on record. However, there are some measures of progress:

 * We have been talking to users and customers a lot. The Flux
   community was instrumental in running the [GitOps
   Days](https://www.gitopsdays.com/) event, for example.
 * We started experimentation in the GitOps Toolkit, designed the APIs
   and wrote the roadmap together with our community. This move [has
   been announced](https://www.weave.works/blog/gitops-with-flux-v2)
   (and [on the mailing
   list](https://lists.cncf.io/g/cncf-flux-dev/message/26)). The
   components aim to be generally re-usable and
   e.g. [Flagger](https://flagger.app/) plans to use the Toolkit’s
   notification-controller.

One thing that did not go well: Weaveworks announced the collaboration
of Flux with the Argo project last year. The engineering teams were in
serious discussions - but at some point it became clear that for the
Argo project, their GitOps Engine was a natural step in their
evolution, but for Flux it would simply make it redundantly more like
ArgoCD (and break backward-compatibility for Flux users). However, the
silver lining was that in seeking our own way forward, the GitOps
Toolkit was born.

## Current goals

The main focus for the team is to build out the GitOps Toolkit, which
will form the basis for Flux v2 and Helm Operator v2.

Once we reach feature parity, we will ask our community of users to
cut over to using v2 (based on the GOTK) within 6 months.

The Flux v2 website has [a
roadmap](https://toolkit.fluxcd.io/roadmap/#the-road-to-flux-v2) for
the above.

### How can the CNCF help us achieve our upcoming goals?

We are struggling somewhat with supporting users on Slack and Github
issues.

 * We recently closed the #flux channel on the Weave Slack as we have
   had the channel in CNCF Slack open for about a year anyway.
 * We are also experimenting with GitHub Discussions for
   [fluxcd/toolkit](https://github.com/fluxcd/toolkit/discussions) --
   this might help redirect some enquiries to a better venue.
 * Still it’d be good to learn more about how other CNCF projects are
   dealing with this by-product of their success.

## Does Flux meet the criteria for incubation?

We think so.

 * Adoption continues to grow, and we see more and more talks, how-to
   guides and integrations around.
 * Cloud vendors and industry partners integrate with Flux and are
   watching GOTK development closely.
 * As Flux is in a very interesting phase of its development right
   now, contributions go to two entirely different branches of Flux -
   this will continue for some while, but we have high hopes for
   development getting a lot easier and more straightforward with the
   GOTK.
