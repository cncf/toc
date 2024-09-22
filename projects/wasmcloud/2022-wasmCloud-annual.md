# wasmCloud Annual Review

## Devstats

[https://wasmcloud.devstats.cncf.io/d/8/dashboards?orgId=1](https://wasmcloud.devstats.cncf.io/d/8/dashboards?orgId=1)

## Highlights

### Community growth, development, and adoption

- Our stargazers count [nearly doubled](https://wasmcloud.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1&from=1641013200000&to=1672549199000&var-repo_name=wasmCloud) over the course of 2022.
- We see contributions from [over 50 companies](https://wasmcloud.devstats.cncf.io/d/5/companies-table?orgId=1) across the wasmCloud org.
- We respond to pull requests on average in [less than a day](https://wasmcloud.devstats.cncf.io/d/10/pr-time-to-engagement?orgId=1&from=1641013200000&to=1672549199000).

## Maintainers

We have multiple organization maintainers in our [wasmcloud/wasmcloud OWNERS file](https://github.com/wasmCloud/wasmCloud/blob/main/OWNERS), most of which are employed by Cosmonic, as well as additional maintainers for specific projects. [jordan-rash](https://github.com/orgs/wasmCloud/people/jordan-rash) is a maintainer over the majority of the TinyGo and Go libraries and works at Capital One, [mattwilkinsonn](https://github.com/orgs/wasmCloud/people/mattwilkinsonn) is a maintainer on our wash CLI and works at Google, [rvolosatovs](https://github.com/orgs/wasmCloud/people/rvolosatovs) is a maintainer of our latest iteration of the Rust wasmCloud host. We recently accepted the community-led [kasmcloud](https://github.com/wasmCloud/kasmcloud) project as a part of our organization with a new maintainer [Iceber](https://github.com/wasmCloud/kasmcloud/commits?author=Iceber).

## Governance

Please see our [GOVERNANCE.md](https://github.com/wasmCloud/wasmCloud/blob/main/GOVERNANCE.md) document in our main wasmCloud repository.

## Long term planning

wasmCloud officially tracks its high level initiatives and current work in a [GitHub Roadmap](https://github.com/orgs/wasmCloud/projects/7/views/3) which includes both a Now/Next/Later view for large efforts, and a Kanban board for current work. We also track collections of issues in [milestones](https://github.com/wasmCloud/wasmCloud/milestones) when appropriate for larger efforts.

## Development is ongoing and progressive towards versioned releases.

We continually update our [wasmCloud](https://github.com/wasmCloud/wasmCloud/releases) releases according to SEMVER guidelines and plan to continue to do so. Changelogs are automatically generated based on conventional commits and are available for each release. The stability of wasmCloud to push for a v1.0.0 release is a top priority after the stability and release of the WASI developer preview: `wasi-preview-2`.

## Adopters

We recently opened up our [ADOPTERS.md](https://github.com/wasmCloud/wasmCloud/blob/main/ADOPTERS.md) to formally track production use cases, and have met the three required production use cases for incubation. We see increasing adoption of wasmCloud in hobby projects and at the individual developer level in our [community slack](https://slack.wasmcloud.com).

## Goals since previous review

Though I'm not able to find a previous review of wasmCloud, we're happy with the growth and feedback since last year. We see the scope and goals remaining consistent for our next, 2023, annual review.

## Current Goals

- Meet the requirements, apply for, and achieve Incubating status in the CNCF.
- Onboard 3 more maintainers of specific repositories in the project with our recent focus of shifting feature focus to Rust.
- Lead education on WebAssembly components, released in the upcoming WASI preview 2, in order to create better applications for cloud native engineers.
- Support all programming languages that feature proper WebAssembly component support for building wasmCloud actors to broaden our user base, fulfilling our goal to write in whatever language of choice.
- Stabilize and release wasmCloud v1.0.0.

## How can the CNCF help?

- Enabling codespaces and prebuilds on the wasmCloud organization to let contributors get started quicker and contribute faster.
- Help us identify opportunities to continue to make wasmCloud cloud native. For example, our open telemetry integration helped us succeed with built-in tracing, and we would love to collaborate with other projects.
- Continue to steward the WebAssembly community in the CNCF, growing appetite for contributors in the standards space.

## Criteria for Incubation

We meet all incubation requirements and have a willing sponsor for incubating phase. At the last CNCF TAG-Runtime meeting, we presented wasmCloud, our growth, and a demonstration of the project in order to kick off this next stage for the project: [https://youtu.be/j_cUEQttTuw?t=1749](https://youtu.be/j_cUEQttTuw?t=1749). I plan to submit a PR to this repository proposing wasmCloud for incubation in the next few weeks.
