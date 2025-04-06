# Dex 2021 Annual Review

## Background

Dex is an identity service that uses OpenID Connect to drive authentication for other apps.

Dex acts as a portal to other identity providers through "connectors." This lets Dex defer authentication to LDAP servers, SAML providers, or established identity providers like GitHub, Google, and Active Directory. Clients write their authentication logic once to talk to Dex, then Dex handles the protocols for a given backend.

Dex's main production use is as an auth-N addon in for Kubernetes. Dex runs natively on top of any Kubernetes cluster and can drive API server authentication through the OpenID Connect plugin. Clients, such as kubectl, can act on behalf users who can login to the cluster through any identity provider Dex supports.

Dex has been used by numerous companies who would like for it to find a neutral home where they can actively contribute to the long term viability of the project.

### Sandbox Proposal Artifacts

- Sandbox Proposal PR: [PR #379](https://github.com/cncf/toc/pull/379)
- Sandbox Proposal doc: [Proposal
  doc](https://github.com/cncf/toc/blob/main/proposals/dex.md)

Dex was accepted into the CNCF Sandbox on June 23th, 2020.

## Highlights from the past year

- We published a [new website](https://dexidp.io/) and moved over most of the documentation.
  We've also spent a considerable amount of time improving the content.
- A new maintainer came on board.
- Dex has a new [Helm chart](https://github.com/dexidp/helm-charts) to make installation on Kubernetes easier.
- We've released 6 minor versions along with a new gRPC API client.

## Development stats

> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity.
> Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

The Dex devstats page and dashboards can be found
[here](https://dex.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

Key highlights:

- The project currently has a total of [330+ contributors](https://dex.devstats.cncf.io/d/22/prs-authors-table?orgId=1) from [200+ companies](https://dex.devstats.cncf.io/d/5/companies-table?orgId=1).
- During the last year [70+ people](https://dex.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All) contributed to Dex from [70+ companies](https://dex.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions).
- According to devstats, the project averages [41 commits](https://dex.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) in [11 merged PRs](https://dex.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) per month.

## Maintainers

> How many maintainers do you have, and which organisations are they from?
> (Feel free to link to an existing MAINTAINERS file if appropriate.)

The project has **5** active maintainers from **3** different organizations (Cisco, Red Hat, Flant).
The full list of maintainers can be found [here](https://github.com/dexidp/dex/blob/master/MAINTAINERS).

[Maksim Nabokikh](https://github.com/nabokihms) (Flant) joined the project last year after being an active contributor for a long period of time.

Although we don't have a formal process for onboarding new maintainers,
we definitely would like to have one.
The informal requirements are showing continued interest and domain expertise in the project.

## Adoption

> What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?
> If you can list companies that are end users of your project, please do so.
> (Feel free to link to an existing ADOPTERS file if appropriate.)

The list of known adopters can be found [here](https://github.com/dexidp/dex/blob/master/ADOPTERS.md),
but this is only a small portion of our users.

Interestingly, users can be placed in two categories:

- Projects that integrate Dex into their platforms/applications
- Users who run Dex as an OIDC proxy in front of their Identity Providers, so that they can use OIDC in various applications

Most members of the ADOPTERS file belong to the first group,
but based on the amount of issues and PRs that we receive,
there is a vast number of users in the second group as well.

We've opened two communication channels during the last year:

- [dexidp](https://cloud-native.slack.com/archives/C01HF2G1R34) channel on CNCF Slack
- [GitHub Discussions](https://github.com/dexidp/dex/discussions)

Both channels show an increasing number of users based on the questions and feedback that we receive.

## Project goals

> How has the project performed against its goals since the last review?
> (We won't penalize you if your goals changed for good reasons.)

Although this is our first review, we did set some goals when Dex joined the CNCF:

- Improve developer experience, make it easier to contribute to the project
- Set up a website and start improving documentation
- Provide more resources for installing Dex

In the last year, we've managed to deliver each of these goals:

We've formalized issue reporting and PR submission by using issue templates.
Build workflows are more stable than before and cover a lot more cases.
Developer tooling has improved a lot.

We've setup a website and moved over the documentation.
The documentation also improved a lot during the last year.

We've released a new Helm chart for installing Dex on Kubernetes.
It's based on our years of experience in operating Dex,
so even the first version came with quite a few knobs for users to turn.

> What are the current goals of the project?
> For example, are you working on major new features?
> Or are you concentrating on adoption or documentation?

We have a number of items on our roadmap that we'd like to accomplish next:

- More frequent releases
- (Even) better documentation
- Formal process for becoming a maintainer
- Get more maintainers and adopters
- Run tests on ARM architecture (already in progress, thanks to CNCF infra)
- OIDC certification
- Connector middleware allowing users to inject custom logic into the auth process (eg. use custom claims)
- Custom connectors

Our current priority is adoption, documentation and stabilization.
The above goals are in no particular order, nor we think that we can necessarily reach all of them in a single year.

## CNCF support

> How can the CNCF help you achieve your upcoming goals?

In order to move to the next stage, we need more active contributors, maintainers and users ("customers").
This last bit is important, so we can drive our roadmap based on real user needs.
There are requests that keep coming back, so we have some user stories,
but we need to understand better how Dex is used.

The CNCF can help by providing more exposure to the project, connecting the maintainers and the community.
We plan to have a virtual booth in the next KubeCon which is a great opportunity,
but we are open to other suggestions, even some coaching/tips about reaching out to the community.

## Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage)?

No, at this point we don't feel that the project meets the criteria.
We would like to grow more in terms of maintainers and known users before considering to move to the next stage.
