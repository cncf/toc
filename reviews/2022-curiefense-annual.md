# Curiefense-2022-Annual Review

## Background

Curiefense is an API-first, DevOps oriented web-defense HTTP-Filter adapter for
Envoy and NGINX. It provides multiple security technologies (WAF,
application-layer DDoS protection, bot management, and more) along with
real-time traffic monitoring and transparency.


## Include a link to your project’s devstats page.

Please refer to Curiefense's [devstat page](https://curiefense.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=1h)

## How many maintainers do you have, and which organisations are they from?

Currently the project has 10 unique maintainers all related to Reblaze across
the main repositories. While the team has grown in numbers, it has yet to
diversify itself. The community is expanding and there has been more engagement
on Slack, GitHub, and other platforms.


## What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so.

The community has started to grow and the team has seen an increased amount of
interest in the project and collaborations. Members of the community have
reached out with new feature requests, and code proposals.

We have also seen an increased amount of MVP/trial environments being created.

Finally, there are 2 companies using it in production. Working on getting
approval for a public reference.


## How has the project performed against its goals since the last review?

This is Curiefense's first annual review. The project has had several major
achievements from a community and technical perspective that are listed below:

- Expanded the list of supported proxies: Now, in addition to Envoy, the
  project can also integrate with Nginx, which is one of the most widely used
  proxies.

- The project has been simplified and shrunken down to ease adoption and
  deployment. This effort allowed the community to think through a good
  deprecation process for major features and services.

- The curiefense-helm code has been split into its own project, providing more
  focus, clarity, and ease of navigation.

- CI stabilization was a major focus for the team. During this period the team
  focused on creating an e2e environment that would match a real production
  environment at a smaller scale.

## What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?


**New Features / Refactoring**

- The curiefense-helm project will be undergoing a major refactoring soon,
  which sparked from conversations with community members over slack.
  This is separate from the effort to split the chart into its own repo.

- Support for new proxies
- New security protection strategies

**General improvements**

- A major refactor to the test suite is in progress. The aim is to make the
  tests agnostic to the environment, to make it simpler, and faster.

**Documentation improvements**

- Improve end user documentation. The team would like to simplify the quick
  start guide so that it would be easier for newcomers to rampup and
  familiarize themselves with the project.


## How can the CNCF help you achieve your upcoming goals?

Curiefense has benefitted a lot from being part of the CNCF. The attention the
project has gotten come not only from the hard work of the team but also
because of the platform that the CNCF has provided.

The same continuos guidance and support is what the team would like to keep
getting from the CNCF.

There are currently no other requirements.

## Do you think that your project meets the criteria for incubation?

Curiefense is not ready to apply for incubation due to limited adoption and the
lack of corporate diversity within its maintainers. The team is working on
improving adoption, which will hopefully result in new maintainers and
organizations joining the community. The team would like improve adoption before
applying for incubation.

