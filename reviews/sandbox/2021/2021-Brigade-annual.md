# Brigade Annual Review 2021

## Background

[Brigade](https://brigade.sh) is an event-driven scripting framework that runs on Kubernetes.  The project is currently making strides towards the next major 2.0 release, which brings a number of big improvements and design changes, all while continuing to use Kubernetes as its substrate.

In 2.0, however, one of the main goals is to allow users of the system to be productive *without* needing any Kubernetes experience.  Perhaps equally important, they also do not need any authorization permissions on the Kubernetes cluster itself (as they do in 1.0).  This is accomplished with an entirely new Brigade API, which handles authentication and authorization concerns.  This API server is now what the `brig` CLI and all SDKs communicate with.

Some other notable changes in 2.0 involve a messaging queue system to support higher throughput and performance, a database backend to support log persistence (which brings the added benefit of keeping the footprint on Kubernetes as low as possible, with swift cleanup of pods shortly after an event has reached a terminal state) and new project configuration templates to support quick creation/updates as well as version control tracking.

## Year in Review

After acceptance into the CNCF Sandbox and a successful [1.0](https://cloudblogs.microsoft.com/opensource/2019/03/28/announcing-brigade-1-0-new-kind-of-distributed-application/) release the year prior, this year was all about Brigade 2.0:
 - Unanimous maintainer approval of the [Brigade 2.0 Proposal](https://github.com/brigadecore/brigade/pull/1101)
 - Migration of key components from the [2.0 Prototype](https://github.com/krancour/brignext) into Brigade's main repo via the [v2 branch](https://github.com/brigadecore/brigade/tree/v2), where development proceeded in earnest
 - First alpha release: [v2.0.0-alpha.1](https://blog.brigade.sh/2021-03-05-meet-brigade-2/)
 - Development of SDKs for Brigade 2.0's API (in addition to the [Go SDK](https://github.com/brigadecore/brigade/tree/v2/sdk)): [Brigade SDK for Javascript/Typescript](https://github.com/brigadecore/brigade-sdk-for-js) and [Brigade SDK for Rust](https://github.com/brigadecore/brigade-sdk-for-rust) (still in a prototype stage)


## Annual Review Items

- *Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.*
  - Pulling from the main [Brigade devstats dashboard](https://brigade.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m):
    - Participation from [73 companies](https://brigade.devstats.cncf.io/d/5/companies-table?orgId=1) and counting, including [14](https://brigade.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions) in the last year.
    - Seeing an average of [5 new PRs a week](https://brigade.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1553666400000&to=now) in the last year - similar to the year prior -- but see an increase to [8 new PRs a week](https://brigade.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1553666400000&to=now) in the last 5 months as 2.0 development has accelerated.
    - We have seen [new contributor PRs drop slightly](https://brigade.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now) compared to the [year prior](https://brigade.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-2y&to=now-1y). We believe this is due to the stability and feature-set of 1.0 and are expecting this trend will reverse as adoption and familiarity with Brigade 2.0 increases. 


- *How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)*
    - Brigade currently has [7 core maintainers](https://github.com/brigadecore/brigade/blob/master/CODEOWNERS#L4):
        - Matt Butcher, Microsoft
        - Adam Reese, Microsoft
        - Yusuke Kuoka, Z Lab
        - Radu Matei, Microsoft
        - Luke Philips, Charter Communications
        - Kent Rancourt, Microsoft
        - Vaughn Dice, Microsoft

- *What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)*
  - The following stats as reported in last year's review continue to hold.  As we've been heads-down on Brigade 2.0, we haven't been as focused on increasing adoption of Brigade v1.  We anticipate to shift outreach and new adoption, this time for v2, as we near a final release candidate.
    - According to the CNCF Annual Survey 2019, over 40% of respondents have used Brigade
    - Core Adopters are:
        - Charter
        - Z Lab (Yahoo Japan)
        - iBotta
        - Microsoft
        - Student.com
        - Spotahome.com
        - ThredUp

- *How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)*
  - Here are the goals stated in the 2020 Brigade Annual Review, along with notes/status updates from the past year:
    - Make significant strides towards Brigade 2.0
      - Indeed, this was *the* top priority item in 2020.  As summarized above, we achieved this goal, culminating with the first alpha release.  In addition, milestones have been set for the subsequent [alpha](https://github.com/brigadecore/brigade/milestone/6) and [beta](https://github.com/brigadecore/brigade/milestone/7) releases.
    - Improve notes and awareness for Brigade releases
      - In tandem with the [v2.0.0-alpha.1](https://github.com/brigadecore/brigade/releases/tag/v2.0.0-alpha.1) release, we published an announcement on Brigade's blog: https://blog.brigade.sh/2021-03-05-meet-brigade-2/
    - Increase rate of new contributors
      - We weren't as successful with this goal, having been more focused on 2.0 development.  As the alpha and beta releases go out, however, we will be looking to increase community feedback and contributions.
    - We want to set up Brigade to move into incubation in 2022.
      - We believe we are still on track with this goal, which should segue nicely after Brigade's scheduled 2.0 release later this year.

- *What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?*
  - Ramp up external contributions for Brigade 2.0
  - Increase adoption and community awareness around Brigade 2.0
  - Release Brigade 2.0 final
  - Continue to make progress towards move into incubation in 2022.

- *How can the CNCF help you achieve your upcoming goals?*
    - We believe opportunities like occasional CNCF webinars or blog posts are good for informing the CNCF user base of our project's features and goals.
    - When we get closer to a final 2.0 release candidate, we feel that we need to develop a platform for reaching new contributors. The current anticipated timeline is:
      1. Brigade 2.0.0-alpha.2 will be released April, 2021
      2. Brigade 2.0.0-beta.1 will be released May, 2021
      3. Brigade 2.0.0 final is anticipated to be released by Q4, 2021
    - We are unsure what a sandbox project will be able to ask of CNCF, given the new sandbox definition.

- *Do you think that your project meets the criteria for incubation?*
    - While we believe we have made progress toward that goal, we want to release Brigade 2.0 before entering incubation.