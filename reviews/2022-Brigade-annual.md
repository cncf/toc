# Brigade Annual Review 2022

## Background

[Brigade](https://brigade.sh) is an event-driven and general-purpose scripting
platform that leverages Kubernetes as a workload execution substrate. In the
past year, the project has completed a major pivot from its v1 architecture to
its v2 architecture. The primary impetus and completed objective of that pivot
was to provide a cloud-native scripting platform that _abstracted_ Kubernetes
away from end-users to a sufficient degree that neither Kubernetes expertise nor
direct access to a Kubernetes cluster would represent barriers to productivity.
Only operators executing Brigade installation require any considerable degree of
Kubernetes acumen.

Brigade is a _unique_ product in the cloud-native landscape because it enables
developers to respond to events with workflows that are defined using
_imperative scripting languages_ (JavaScript or TypeScript) _instead of a
declarative language_ (such as YAML). Uniqueness aside, the project maintainers
believe this approach sidesteps significant impedance that is inherent to
modeling workflows declaratively. Apart from this reduced friction, using
JavaScript and TypeScript for this purpose provides many other benefits. Their
enormous popularity means many developers are comfortable working with these
languages. Developers can also take full advantage of all benefits of those
languages, including, but not limited to, syntax highlighting and context help
in their favorite text editor or IDE, and the ability to incorporate third-party
packages into their scripts using common dependency management tools with which
they are already familiar (`npm` or `yarn`). Developers can even extract useful
patterns they discover into NPM modules for reuse.

## Year in Review

* Brigade v2.0.0 was released on Dec 1, 2021. Five minor releases have followed.

* Numerous v2-compatible "peripheral components" have reached general
  availability, including:
    * [Brigade Bitbucket Gateway](https://github.com/brigadecore/brigade-bitbucket-gateway)
    * [Brigade Cron Event Source](https://github.com/brigadecore/brigade-cron-event-source)
    * [Brigade GitHub Gateway](https://github.com/brigadecore/brigade-github-gateway)
    * [Brigade Noisy Neighbor](https://github.com/brigadecore/brigade-noisy-neighbor)

* Many other peripherals remain in a pre-GA state at this time:
    * [Brigade CloudEvents Gateway](https://github.com/brigadecore/brigade-cloudevents-gateway)
    * [Brigade Dashboard](https://github.com/brigadecore/brigade-dashboard)
    * [Brigade Docker Hub Gateway](https://github.com/brigadecore/brigade-dockerhub-gateway)
    * [Brigade Metrics](https://github.com/brigadecore/brigade-metrics)
    * [Brigade Slack Gateway](https://github.com/brigadecore/brigade-slack-gateway)

* Bi-weekly community meetings (previously on hiatus) were revived. To foster a
  diverse and inclusive community, the times of bi-weekly meetings have varied
  (but are always published in advance) to create opportunities for attendance
  by community members in a variety of international timezones. The project has
  attracted modest attention in Brazil and India.

* Content has been abundant:
    * The [static website](https://brigade.sh) was overhauled.
    * [Docs](https://docs.brigade.sh) have been continuously improved -- and
      community-driven efforts to translate them into Portuguese and Hindi are
      in-progress.
    * The [Brigade blog](https://blog.brigade.sh) was launched and maintainers
      frequently post project news.
    * Our [YouTube channel](https://youtube.brigade.sh) was launched to
      aggregate instructional videos, community meeting recordings, and recorded
      conference talks.

* Maintainers have been obsessively focused on improving the project's footing
with respect to the overall industry's well-founded concerns over software
supply chain security. We have gone to lengths to:
    * Verify commits from all contributors.
    * Ensure no components run as `root`.
    * Minimize dependencies and update them frequently.
    * Publish SBOMs for all Docker images.
    * Sign all Docker images.

## Annual Review Items

### Activity

All metrics represent activity over the last year only:

* [Hourly activity](https://brigade.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-1y&to=now-1h)
* [Participation from 7 companies](https://brigade.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions)
* [An average of 22 new PRs a week](https://brigade.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now&var-period=w&var-repogroup_name=All)
* [16 new contributors](https://brigade.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now)

Coloring the above numbers: A predominance of the activity above is attributable
to one maintainer.

Most new contributors have been students who have made extremely modest
contributions, seemingly to satisfy a directive from a professor to contribute
to an OSS project, and then haven't been heard from since.

### Maintainers

Current core maintainers (some specific repositories have additional
maintainers):

* Matt Butcher, Fermyon
* Radu Matei, Fermyon
* Vaughn Dice, Fermyon
* Kent Rancourt, Microsoft

Of these four core maintainers, only one actively contributes code and content.
Two review PRs. Two are completely absent from the day-to-day operations of the
project and retain their positions only for purposes of maintaining what we
consider an adequately sized body of decision makers. No other active
contributors have yet proven themselves capable of taking up the mantel of core
maintainer.

### Adoption

We have little insight into adoption and what little we have is not encouraging.

Our community does not appear to be as active as it once was and this is
reflected by poorly-attended community meetings and a low degree of engagement
on Slack, YouTube, Twitter, and GitHub. Embarrassingly, attempts to survey the
community have yielded zero responses.

Our best measure of adoption (and a poor one at that), is that there are only
2.7k pulls of our API server component from Docker Hub since the V2.0.0 release
six months ago. This does not speak to widespread adoption.

## Future of the Project

_Architecturally_, the project has achieved its goals and delivered exactly what
we set out to build. We continue to believe that Brigade is unique and valuable.
On a technical level, we are immensely proud of what we have accomplished. In
terms of community and adoption, however, the numbers and commentary above
suggest that the project may have arrived at an impasse.

While current trends could possibly be reversed by allocating the correct
(human) resources to the project, this does not seem to any maintainer to be
tenable as none of our employers currently regard Brigade with any strategic
importance and no active contributors have yet proven themselves capable of
taking up the mantel of core maintainer as the existing maintainers continue to
divest.

In light of this, the four current core maintainers unanimously agree that the
only remaining goal for Brigade, as we know it, is a thoughtful and responsible
termination of the project -- i.e. archival.

One current maintainer intends to fork Brigade and continue development as a
personal project. Unencumbered by governance and their employer, this maintainer
wishes to seek out a better product/market fit through a series of experiments
and pivots. If this endeavor is met with any success, the results will proudly
be re-donated to CNCF.

__Amendment:__ Since the initial recommendation to archive the project was made,
a small cohort of contributors has graciously volunteered as new core
maintainers. Per Brigade's own governing docs, none of these individuals meet
the criteria to be considered for this responsibility, but in light of the
extenuating circumstances, with all current maintainers recommending archival,
we believe the TOC should determine whether such a transition is permissible.
