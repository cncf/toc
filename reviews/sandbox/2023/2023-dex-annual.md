# Dex 2023 Annual Review

## Highlights from the past year

- Dex implemented RFC 8693 OAuth 2.0 Token Exchange
- OIDC compliance has improved
- We've released 10 minor versions (4 more compared to the previous year).

## Development stats

The Dex devstats page and dashboards can be found
[here](https://dex.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&search=open).

Key highlights:

- The project currently has a total of [400+ contributors](https://dex.devstats.cncf.io/d/22/prs-authors-table?orgId=1) from [300+ companies](https://dex.devstats.cncf.io/d/5/companies-table?orgId=1).
- During the last year [60+ people](https://dex.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Last%20year&var-repogroup_name=All) contributed to Dex from [60+ companies](https://dex.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions).
- According to devstats, the project averages [30 commits](https://dex.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) in [7 merged PRs](https://dex.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now) per month.

## Maintainers

The project has **5** active maintainers from **3** different organizations (Cisco, Red Hat, Flant).
The full list of maintainers can be found [here](https://github.com/dexidp/dex/blob/master/MAINTAINERS).

We have a semi-formal onboarding process and we are looking at the prospect of adding more committers to the project.

## Adoption

The list of known adopters can be found [here](https://github.com/dexidp/dex/blob/master/ADOPTERS.md),
but this is only a small portion of our users.

Prominent CNCF/Open Source projects using Dex are [ArgoCD](https://argo-cd.readthedocs.io/en/stable/) and [sigstore](https://www.sigstore.dev/).

## Performance against past goals

### More frequent releases

We've nearly doubled the number of releases during the past year.

### Formal process for becoming a maintainer

We have a semi-formal process for becoming a maintainer.
Becoming a maintainer today is definitely easier than a couple years ago.

### Get more maintainers and adopters

Adoption measurable improved during the last year and we also have more committers.

We failed to add new maintainers to the project, but we have potential candidates.

### Run tests on ARM architecture

We have an ARM-based build process in place (although we cross compile instead of running tests on native ARM).

### OIDC certification

OIDC spec compliance improved in the last year, but we haven't reached certification yet. We are going to work towards that goal in the coming year.

## Future goals

We have a number of items on our roadmap that we would like to accomplish next:

- Improved (e2e) test coverage
- OIDC certification
- Connector middleware allowing users to inject custom logic into the auth process (eg. use custom claims)
- Custom connectors

The above goals are in no particular order, nor we think that we can necessarily reach all of them in a single year.

## CNCF support

Dex integrates with a number of Identity Providers and it's rather hard to maintain compatibility.
We need access to resources (eg. access to cloud providers) to continuously test compatibility with these providers.

## Incubation

At this point we don't feel that the project meets the incubation criteria.
We would like to grow more in terms of maintainers and known users before considering to move to the next stage.
