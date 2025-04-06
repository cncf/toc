# Brigade Annual Review 2020

## Background

[Brigade](https://brigade.sh) is an event-driven scripting framework for Kubernetes.  It's built directly on top of the Kubernetes APIs, for use in any conformant cluster.  Javascript is used as the scripting language and Docker images comprise the building-block workers in charge of processing events and executing actions.

The broader [Brigadecore](https://github.com/brigadecore) GitHub organization not only hosts the main [Brigade](https://github.com/brigadecore/brigade) repo, but also the modular components that empower users to configure Brigade to their needs.  This includes:

- [Brigadier](https://github.com/brigadecore/brigadier), the main Javascript library used in Brigade scripts
- The [Brigade-Github-App gateway](https://github.com/brigadecore/brigade-github-app), one of several popular gateways utilized for event creation
- The [Charts](https://github.com/brigadecore/charts) repo containing all of the [Helm](https://helm.sh) charts needed to deploy Brigade
- [Kashti](https://brigadecore/kashti), a UI Dashboard for Brigade
- and many others, all of which feature contributions from open source community members.

**Brigade was accepted the CNCF Sandbox project on March 12, 2019**
- [Brigade Sandbox Proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/brigade.adoc)
- [TOC proposal slides](https://docs.google.com/presentation/d/1K2mI6CCsCLNl4pK4IL80jekq0-DBgcbWW8iPhOsSXaM/edit#slide=id.g5207231d02_0_0)

## Year in Review

This last year was arguably Brigade's most exciting, including:
- CNCF Sandbox project acceptance
- [Brigade v1.0](https://cloudblogs.microsoft.com/opensource/2019/03/28/announcing-brigade-1-0-new-kind-of-distributed-application/), Brigade's first major release 
- Migration of all Brigade projects into its own [Brigadecore](https://github.com/brigadecore) GitHub organization
- Several incremental feature releases, most recently [Brigade v1.4.0](https://github.com/brigadecore/brigade/releases/tag/v1.4.0)
- Maintainers Summit in Boulder, CO to plan Brigade 2.0 
    - We generated a list of Future User Stories: https://hackmd.io/DN152vISS72IWD2Roty7Ag
    - This led to an [early/informal prototype of Brigade 2.0](https://github.com/krancour/brignext) that informed the formal proposal: https://github.com/brigadecore/brigade/pull/1101
- At Helm Summit Amsterdam, [a session featured](https://helmsummit2019.sched.com/event/S8tc/ship-it-faster-safer-cheaper-state-of-the-art-of-gitops-with-helm-yusuke-kuoka-z-lab-corporation) a Brigade and Helm integration project
- At KubeCon San Diego, the Brigade team demonstrated [using Brigade as a Kubernetes Controller](https://github.com/brigadecore/buck).


## Annual Review Items

- *Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.*
  - Pulling from the main [Brigade devstats dashboard](https://brigade.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m):
    - Participation from [50 companies](https://brigade.devstats.cncf.io/d/5/companies-table?orgId=1) and counting
    - [Consistent release cadence](https://brigade.devstats.cncf.io/d/47/github-events?orgId=1&from=1553666400000&to=now) of a feature release every 3-4 months since 1.0
    - Seeing an average of [5 new PRs a week](https://brigade.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=1553666400000&to=now) in the last year
    - Since 1.0, we have seen [new contributor PRs drop slightly](https://brigade.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now) compared to the [year prior](https://brigade.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1522130400000&to=1553666400000). We believe this is due to the stability and feature-set of 1.0 and are expecting this trend will reverse when we open up feature development for Brigade 2.0. 


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
    - This is Brigade's first annual review, so we didn't appear to have official goals set.  However, after reaching the 1.0 milestone, our goals had been to:
        - Maintain a relatively consistent feature release cadence
        - Improve documentation and increase tutorials/examples ([v1.3.0](https://github.com/brigadecore/brigade/releases/tag/v1.3.0) heavily centered around this)
        - Attend to bug reports and security patches in a timely manner 
        - Divert some full-time resources towards 2.0

    - On these counts, we consider ourselves successful

- *What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?*
  - Make significant strides towards Brigade 2.0
  - Improve notes and awareness for Brigade releases
  - Increase rate of new contributors
  - We want to set up Brigade to move into incubation in 2022.

- *How can the CNCF help you achieve your upcoming goals?*
    - We believe opportunities like occasional CNCF webinars or blog posts are good for informing the CNCF user base of our project's features and goals.
    - When the 2.0 planning cycle ends and development begins, we feel that we need to develop a platform for reaching new contributors. The current anticipated timeline is:
      1. Brigade 2.0 development kicks off on July 20, 2020 (unless the proposal needs updates)
      2. Brigade 2.0.0-alpha.1 will be released January, 2021
      3. Tentative date for Brigade 2.0.0 final is June 30, 2021
    - We are unsure what a sandbox project will be able to ask of CNCF, given the new sandbox definition.

- *Do you think that your project meets the criteria for incubation?*
    - While we believe we have made progress toward that goal, we want to release Brigade 2.0 before entering incubation.