# Telepresence Sandbox Annual Review

The goal of the Telepresence project is to accelerate the "developer inner loop" for cloud-native application development on Kubernetes. Telepresence achieves this by creating a dynamic network bridge between a local development environment (e.g., a laptop) and a remote Kubernetes cluster.

Telepresence was presented to the TOC on April 2018. Please see the project's [Sandbox proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/telepresence.adoc).

## Annual Review Questions

### Metrics

| Metric       |  Pre-Sandbox |    Today     |
| ------------ | ------------ | ------------ |
| Docker Pulls | 100K+        | 10M+         |
| Stars        | 700+         | 2700+        |
| Forks        | ~50          | ~190         |

The CNCF provides Telepresence Github metrics through its [devstats site](https://telepresence.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

### Maintainers

There are four maintainers in total that work for Datawire.

### Adoption

The Telepresence project has seen continued growth and adoption since entering the CNCF Sandbox. Some of the known adopters include:

* Avaya
* Bose
* Engel & Volkers (KubeCon EU 2019 presentation: https://www.youtube.com/watch?v=tD0FIlxO1AQ)
* InVision
* Namely
* PTC
* Red Hat
* Timecampus
* ThetaRay

Quote from Timecampus: "It initially took a lot of time for me to build Docker images every time I did some development. Telepresence improved my development cycle dramatically. I just swap the deployment, keep developing and push only when I am done. So easy."

Quote from Avaya: "This tool combined with devspace works magic for our development workflow. Being able to debug directly in any of our installations its pure gold. Keep up the good work!"

There have also been a number of integrations with other software stacks including:

* wad.sh (https://blog.sebastian-daschner.com/entries/reloading-apps-kubernetes-telepresence-wad)
* Ambassador Edge Stack (https://www.getambassador.io/docs/latest/topics/using/edgectl/)
* Kelda (https://kelda.io/docs/tutorials/telepresence/)

In the past 6 months, our monthly download rate has more than doubled to ~2000 installations/month on Linux. (We do not track Mac OS X installations, as they are hosted on GitHub and the data is not retained.)

### Past Goals

This is our first review, and we didn't state goals in the original Sandbox proposal. In the past year, the project has focused on the following areas:

* Overhauling the automated test infrastructure, to enable more rapid development of changes and simplify contributions from non-maintainers
* Introducing a new connection method, Teleproxy, which is designed to be more reliable than other approaches.
* Architectural changes to support incrementally key new features to Go while maintaining existing features in Python. We started switching to Go to take better advantage of the Kubernetes ecosystem libraries.
* Improving robustness in a wide variety of different client (MacOS X, Linux, Windows WSL 2) environments
* Continuing to support the evolving Kubernetes APIs

### Current Goals

In the coming year, we'd like to broaden our contributor base. Telepresence is a technically sophisticated project, and requires in-depth expertise in Kubernetes and understanding the quirks of Mac OS X and Linux networking. Some of the things we're planning include:

* Continued refactoring of the code base in to Go (which also allows us to take better advantage of Kubernetes ecosystem libraries)
* Better documentation of the Telepresence development workflow
* Continued work on the test suite

We've also had a number of user requests that we'd like to address, including:

* Better support for unreliable connectivity to the cluster
* Support for project-specific configuration
* Better documentation

### How can the CNCF help?

We'd welcome help in expanding our user base by having CNCF contributors contribute and publish tutorials and improve our documentation.

### Incubation

We would like to add one non-Datawire maintainer before moving to incubation. From a product adoption and maturity standpoint, we believe we meet the incubation criteria.
