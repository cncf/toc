# 2020 Buildpacks Annual
## Background
* [Buildpacks Sandbox Proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/buildpacks.adoc)
* [Original CNCF TOC Meeting Slides](https://docs.google.com/presentation/d/1InW4LbgFb8dSCCgbN7hozNTAFZTsIVBWxm8sv87TjZs/edit?usp=sharing)
* [Cloud Native Buildpacks](https://buildpacks.io) translate source code into container images with a focus on developer productivity, container security, and day-2 operations involving container images at scale.

Notable improvements in the last year include:
* Introduced a [Request for Comments](https://github.com/buildpacks/rfcs) system to encourage collaboration and share knowledge
* Designed a Distribution Spec for sharing and packaging buildpack(s) (https://github.com/buildpacks/spec/blob/master/distribution.md)
* Add Image Introspection with inspect image command and bill of materials provided by buildpacks.
* Reproducible Builds with ["zero out" creation dates](https://github.com/buildpacks/imgutil/pull/19).
* Introduce a [Project Descriptor, project.toml](https://github.com/buildpacks/rfcs/blob/master/text/0019-project-descriptor.md).

## Alignment with Cloud Native
Cloud Native Buildpacks is part of the App Delivery SIG and [presented to the SIG on 12/4/2019](https://docs.google.com/presentation/d/1iM_47YWQWpvQ09fdGYxutbX9eB_PAB0Sl_L9JoEFYEE/edit?usp=sharing). The project cleanly fits into the SIG model in Topic 1 (Application Definition) and Topic 1.5 (Application Packaging). For Topic 1, Buildpacks themselves and a [Project Descriptor](https://github.com/buildpacks/rfcs/blob/master/text/0019-project-descriptor.md) define a way a application is built. At it's core, the project performs source to image builds and "rebasing" Operating System level security patches on existing images which fit into Topic 1.5. Ultimately, Buildpacks "empower organizations build scalable applications in modern, dynamic environments such as public, private, and hybrid clouds" by producing OCI compliant images that can be run on any Cloud Native infrastructure.

## Annual Review Questions
* Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.
  * [CNCF DevStats](https://buildpacks.devstats.cncf.io)
  * Deliverables do not include buildpacks themselves. Instead this project is about fostering an ecosystem where individuals or companies can use the specification and tools provided by this project to create and use them.

* How many maintainers do you have, and which organisations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)
  * There are 6 maintainers in total that work for VMware and Salesforce.

* What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)
  * [ADOPTERS.md](https://github.com/buildpacks/community/blob/master/ADOPTERS.md)

* How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)
  * This is the first time we've done a review (and we didn't state goals as part of our [Sandbox proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/buildpacks.adoc)).
  * For the first year we did have some goals in mind:
    * Grow community to include companies/users beyond Pivotal/Heroku - From the weekly working group meetings and slack we have participation from Google, Microsoft, Dokku, and Deft among others.
    * Use the Buildpacks in production - At the start of the Sandbox period we had 0 users using buildpacks in the wild, but now we have [ADOPTERS.md](https://github.com/buildpacks/community/blob/master/ADOPTERS.md).
    * 1.0 of the Specification - We didn't quite achieve this, but after getting feedback from users, we've made some breaking changes and plan to release 1.0 this year.

* What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?
  * [2020 Roadmap](https://medium.com/buildpacks/cloud-native-buildpacks-2020-roadmap-b7e43876473a)
  * Maturity
    * Finish up half finished features from last year like Buildpackages
    * codify the release cadence/process across the project
    * polish up the tools as we near 1.0.
  * Based off of feedback from KubeCon NA, we want to improve our the experience after `pack` for users wanting to integrate Buildpacks in their own build pipeline.
  * Release 1.0 of the Spec
  * Continue to expand our community of user and contributors

* How can the CNCF help you achieve your upcoming goals?
  * Are there more opportunities to engage potential interested users beyond KubeCon? Sandbox projects more than any project needs help with marketing and outreach.

* Do you think that your project meets the criteria for incubation?
  * Yes, we've already [submitted a proposal](https://github.com/cncf/toc/pull/338) for incubation.
