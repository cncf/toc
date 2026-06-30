# Buildpacks Adopter Interview - 7SIGNAL, Inc

**Meeting Time:** 1/27/2026 6:30 PM CET  
**Attendees:**  
- **TOC:** Faseela K, TOC Sponsor of Buildpacks  
- **Adopter:** Dan Mikusa, Staff Software Engineer, 7SIGNAL

---

## About the Adopter Interview

The purpose of the CNCF TOC Adopter Interview is to evaluate the maturity, adoption patterns, and operational impact of the Buildpacks project within production environments.

Raw notes are recorded during the interview and refined into this formal version for adopter review, correction, and approval before being included in Buildpacks due diligence documentation.

---

## Organization Introduction

### Can you give us an overview of your organization and what it does?

7SIGNAL is the leader in digital experience optimization, providing comprehensive visibility across wired and wireless networks, endpoints, and applications. In a nutshell, 7SIGNAL is a SaaS company that provides network, wired and wireless, monitoring software.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

We deploy a number of services to AWS ECS. The software for these services is distributed and deployed through containers.

We needed a simple solution to build the containers for our software that did not take up a lot of time to set up or keep up-to-date. Dockerfiles are the typical choice, but they have a lot of pitfalls. Using Cloud-Native Buildpacks was the winning choice because it was quick to set up, requires very little maintenance, and it works consistently across all of our projects.

## Usage Scenario

### How long has your organization used the project?

Approximately 3.25 years. In production for the same duration.

### What were the main motivations to adopt the project and which key features do you use today?

1. Needed to build containers from our software

2. We are a small team and need to move fast. We don’t have time to manage Dockerfiles and ensure things are being done correctly and consistently in them across all of our projects.

### What is the current level of usage (pre-production, production) and scale?

Production. We have about a dozen applications, a mix of Java & static web servers and one Python app. For our apps that are not being actively developed there may be a handful of builds per month, while our actively developed apps may build many times per day. Every build produces a new container image with CNBs and gets deployed to ECS.

### What version is used and what is your update cadence with the project?

We generally run the latest versions. We use Dependabot to keep our CI pipelines up to date with releases from the CNB project, so we are effectively always close to the latest. There is typically a one-to-two week delay before consuming downstream releases, as we first verify them in an internal development environment.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

When we first started with CNB, there was a bit of work to get a build process that we liked. We’re using GitHub Actions, so we created a build and deploy workflow. It builds the application, runs the tests, and then invokes CNB tooling to generate and publish a container image to our private container registry, which is then deployed to AWS ECS.

There was a small amount of customization to pick a builder, buildpacks, and then pass some config through to the buildpacks. We did this once and have been able to use this across all of our projects. The only difference has been selecting different buildpacks based on the language being used in the app (i.e. Java buildpacks, Python buildpacks, or web server buildpacks)

As a way to check the images produced by buildpacks, we also have these pipelines running Trivy against both the application & the image that’s produced. We’re using this as a way to confirm that the images being produced are as up-to-date and CVE free as the application code permits.

In summary, the initial setup was small and we’ve been able to use it for over three years now with almost no updates or changes, so it’s been a good choice and very solid.

We do have a legacy Node.js application that is not using buildpacks and is instead containerized with Dockerfiles, as migrating it has proven to be more complex.

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

Yes, the CNB docs are good for getting started. The project also provides some useful GitHub Actions,<https://github.com/buildpacks/github-actions>, which made setting up our pipelines easier.

We have a legacy application that was converted to an nginx-based containerized setup, and we were able to make use of the Buildpacks documentation even for newcomers.

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

Yes.

1. Takes almost no time to set up new projects and get them publishing containers
2. Reduced maintenance time, as we do not need to manage or update Dockerfiles
3. Better security posture through consistently updated base images and tooling

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

The project is currently meeting our needs, and our plan is to continue using it for container generation both in existing projects and as we bring on new projects. I also contribute upstream to buildpacks/libcnb and participate in conference talks related to buildpacks.

## Perception

### What is your perception in terms of the project’s:

* Community openness
* Governance
* Community growth potential
* Maintainer diversity and ladder
* Maintainer response

- The community is friendly and there are opportunities to engage with the project operations, like Slack and various in-person meetings.

- Governance is clear and documented.

- It is difficult to predict growth precisely, but given how widely containers are used across the industry, the potential market for CNB appears quite large.

- I’ve had a number of interactions with the leadership and they’ve all been positive and inclusive. There have been some opportunities for me to get involved with the community and they were helpful in facilitating that and enabling me to contribute.

- From a vendor-neutrality perspective, I have not personally faced issues, though I think there is room for improvement in the TOC structure.

### How are you participating in the project community?

- I’ve opened issues on GitHub and had discussions there.

- I help with one of the projects,<https://github.com/buildpacks/libcnb/>

-I am on the CNCF Slack and will occasionally ask questions there; responses have generally been timely and helpful.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

Yes. This has included asking questions about the tooling and requesting bug fixes, primarily via Slack. The issues we reported tended to be edge cases, and they were addressed quickly with acceptable outcomes.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

The CNB project is well scoped. It does not try to be everything to every one. It has a target problem that it’s trying to solve, and it does that very well, better than the alternative ways to solve the problem.

The tooling is easy to start using. It does not require clusters of hardware, running well on even small machines.

It has a good ROI, as it takes very little effort to integrate with your build processes and you get a number of helpful benefits (discussed previously).

The project has a published specification so there are options for it to be integrated into different platforms and services, and you can see support for it in a number of places. This makes it easier to consume, but also gives it some name recognition which is helpful when pitching it to your boss.

The project demonstrates strong backward compatibility, ensuring that updates do not break existing workflows or applications.

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

I would like to see continued progress toward stabilizing the specification and APIs. Specifically, I am looking forward to the project graduating and reaching a v1 release. While the project has not broken our use cases, this level of stability is a strong signal of maturity and helps justify adoption, especially in more regulated or conservative organizations. I have been impressed with the quality of the documentation and have found it very useful, and I do not have any specific improvements to suggest in that area.

### In your opinion, what can the project do better?

Nothing
