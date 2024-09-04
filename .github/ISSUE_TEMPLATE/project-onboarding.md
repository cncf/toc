---
name: Project onboarding for Sandbox projects
about: Create a checklist of tasks for a project to complete the onboarding process
title: "[SANDBOX PROJECT ONBOARDING] project"
labels: project onboarding, sandbox
assignees: caniszczyk, idvoretskyi, jeefy, krook, mrbobbytables, Cmierly
---

# Welcome to CNCF Project Onboarding

This is an issue created to help onboard your project into the CNCF after the TOC has voted to accept your project into the Sandbox.

We would like your project to complete onboarding within one month of acceptance.

Please track your progress by using "Quote reply" to create your own copy of this checklist in an issue, so that you can update the status as you finish items.

## Review and understand

- [ ] The [project proposal process and requirements](https://github.com/cncf/toc/blob/main/process/README.md).
- [ ] The [services available for your project at the CNCF](https://contribute.cncf.io/resources/project-services/).
- [ ] The [CNCF IP Policy](https://github.com/cncf/foundation/blob/main/charter.md#11-ip-policy).
- [ ] The [trademark guidelines](https://www.linuxfoundation.org/legal/trademark-usage).
- [ ] The [license allowlist](https://github.com/cncf/foundation/blob/main/allowed-third-party-license-policy.md#approved-licenses-for-allowlist).
- [ ] The [online program guidelines](https://github.com/cncf/foundation/blob/main/online-programs-guidelines.md).
- [ ] Optional: [Book time with CNCF staff](http://project-meetings.cncf.io) for any onboarding questions.

## Contribute and transfer

- [ ] Move your project to its own separate neutral GitHub organization. This will make it transferable to the CNCF's GitHub Enterprise account. If it's already in a GHE account, you will need to remove it from that first.
- [ ] Accept the invite to join the CNCF GitHub Enterprise account. We'll then add `thelinuxfoundation` as an organization owner to ensure neutral hosting of your project.
- [ ] Migrate your Slack channels (if any) to the [Kubernetes or CNCF Slack workspace](https://slack.com/help/articles/217872578-Import-data-from-one-Slack-workspace-to-another). CNCF staff can help.
- [ ] Submit a [pull request](https://github.com/cncf/landscape/pulls) to add your project as a Sandbox project to the [Cloud Native Landscape](https://landscape.cncf.io) by updating [landscape.yml](https://github.com/cncf/landscape/blob/master/landscape.yml) following these [instructions](https://github.com/cncf/landscape2/blob/main/docs/config/data.yml).
- [ ] Transfer [your domain to the CNCF](https://jira.linuxfoundation.org/plugins/servlet/theme/portal/2/create/63). The "LF Stakeholder email" is <projects@cncf.io>.
- [ ] Transfer any [trademark and logo assets to the Linux Foundation](https://github.com/cncf/foundation/tree/main/agreements).
- [ ] Submit a [pull request](https://github.com/cncf/artwork) with your artwork.
- [ ] Transfer website analytics to <projects@cncf.io>. CNCF staff can help.

## Update and document

- [ ] Ensure that [DCO](https://github.com/apps/dco) (preferred) or [CLA](https://github.com/cncf/cla) are enabled for all GitHub repositories of the project.
- [ ] Ensure that that the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md) (or your adopted version of it) are explicitly referenced at the project's `README` on GitHub.
- [ ] Ensure LF footer is on your website and [guidelines](https://github.com/cncf/foundation/blob/main/website-guidelines.md) are followed (if your project doesn't have a dedicated website, please adopt those guidelines for the `README` file).
- [ ] Create a maintainer list and add it to the [aggregated CNCF maintainer list](https://maintainers.cncf.io) via pull request.
- [ ] Provide emails for the maintainers to get access to the maintainers mailing list and Service Desk. Email them to <project-onboarding@cncf.io>.
- [ ] Start working on [written, open governance](https://contribute.cncf.io/maintainers/governance/).
- [ ] Start on an [OpenSSF Best Practices Badge](https://www.bestpractices.dev/en).

## CNCF staff tasks

- [ ] Add the project to [DevStats](https://devstats.cncf.io/).
- [ ] Add the project to [CLOmonitor](https://clomonitor.io/).
- [ ] Add the project to [LFX Insights](https://insights.lfx.linuxfoundation.org/). This is done by adding a read-only app to your GitHub organization once it's in CNCF GHE.
- [ ] Add the project to [LFX Project Control Center](https://projectadmin.lfx.linuxfoundation.org/project/a0941000002wBz4AAE).
- [ ] Add a license scanning tool, like [FOSSA](https://fossa.com/) or [Snyk](https://snyk.io/).
- [ ] Invite developers to the [#maintainers-circle](https://app.slack.com/client/T08PSQ7BQ/C014YQ8CDCG) Slack channel.
- [ ] Send a welcome email to confirm maintainer list access.
