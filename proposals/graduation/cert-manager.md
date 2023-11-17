# cert-manager Graduation Proposal

We like to say that [cert-manager](https://cert-manager.io/) is the easiest way to automatically manage certificates in Kubernetes and OpenShift clusters.

In short, the project provides Kubernetes Custom Resources which allow users to request and administer X.509 certificates in a Kubernetes-native way. It also includes a variety of subprojects and tooling to solve the various problems associated with managing such important cryptographic credentials.

cert-manager booths have been at three different KubeCon events at the time of writing, and the maintainers have met users from a huge variety of sectors: automotive, banking, retail, gaming, defence, tech and others.

In fact, many users have told us that they consider cert-manager to be an essentially standard piece of infrastructure in any Kubernetes cluster they create. That's reflected in the statistics we see; the main cert-manager repository has over 11,000 GitHub stars and we have over 8,000 members in our community Slack channel on Kubernetes Slack.

cert-manager's container images have so many daily downloads that, at the time of writing, we're unable to load the statistics page on the container registry we use for cert-manager because it crashes from the load!

## History / Incubation Details

cert-manager was created by Jetstack, a Kubernetes start-up which was later acquired by [Venafi](https://venafi.com/). The [first commit](https://github.com/cert-manager/cert-manager/commit/36308ae891a11a7c4f11839f87d98889c95954fa) was made on 2017-07-19, although it was based on previous work which was started in a now-defunct Jetstack project called `kube-lego`.

Development and adoption was rapid, with regular releases leading up to [v1.0.0](https://github.com/cert-manager/cert-manager/releases/tag/v1.0.0) in September 2020, shortly before cert-manager joined the CNCF as a Sandbox project in November 2020.

cert-manager spent roughly two years as a sandbox project before reaching incubating status in October 2022. For more details about the project's history, see our [maintainer track talk](https://youtu.be/yINv8RUkW-E) from KubeCon NA 2022.

## Graduation State Criteria

### Have committers from at least two organizations

We have [committers](https://github.com/cert-manager/community/blob/main/MAINTAINERS.md) from a variety of organisations:

- Apple
- Diagrid
- G-Research
- Tailscale
- Venafi

We also have active community members who are able to review pull requests, but not to merge them. [1][]

### Have achieved and maintained a Core Infrastructure Initiative [Best Practices Badge](https://bestpractices.coreinfrastructure.org/)

cert-manager achieved its [badge](https://www.bestpractices.dev/en/projects/8079) on 2023-11-17.

### In Progress: Have completed an independent and third party security audit with results published which includes all critical vulnerabilities and all critical vulnerabilities addressed before graduation

The audit is currently ongoing and we've already addressed a few issues raised by the audit team. No critical vulnerabilities have been found so far. [2][]

This section will be updated once the audit completes.

### Explicitly define a project governance and committer process

This process is defined in the [cert-manager/community](https://github.com/cert-manager/community/) repo, which includes the [GOVERNANCE.md](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md) file.

Maintainers are tracked in [maintainers.csv](https://github.com/cert-manager/community/blob/main/maintainers.csv).

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project

Our process for onboarding and offboarding maintainers is also defined in our community repo in our [GOVERNANCE.md](https://github.com/cert-manager/community/blob/main/GOVERNANCE.md) document.

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website)

Adopters are tracked (and encouraged to contribute to) our [USERS.md](https://github.com/cert-manager/community/blob/main/USERS.md) file in our community repo.


## Notes

[1]: Otherwise known as "reviewers" in our [governance docs](https://github.com/cert-manager/community/blob/f04069b6e874bbbd0ae15dd057f44329eb2022a9/GOVERNANCE.md#reviewer).
[2]: Examples of fixed issues are [this webhook change](https://github.com/cert-manager/cert-manager/pull/6498) and [this removal of a deprecated function call](https://github.com/cert-manager/cert-manager/pull/6517).
