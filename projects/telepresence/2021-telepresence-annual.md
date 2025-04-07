# Telepresence Sandbox Annual Review

The goal of the Telepresence project is to accelerate the "developer inner loop" for cloud-native application development on Kubernetes. Telepresence achieves this by creating a dynamic network bridge between a local development environment (e.g., a laptop) and a remote Kubernetes cluster.

Telepresence was presented to the TOC on April 2018. Please see the project's [Sandbox proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/telepresence.adoc).

## Annual Review Questions

### Metrics

| Metric       |  Pre-Sandbox |     2020     |    Today     |
| ------------ | ------------ | ------------ | ------------ |
| Docker Pulls | 100K+        | 10M+         | 10M+         | 
| Stars        | 700+         | 2700+        | 4000         |
| Forks        | ~50          | ~190         | ~300         |

The CNCF provides Telepresence Github metrics through its [devstats site](https://telepresence.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

Reviewing the devstats site above, you can see a significant increase in activity on the telepresence project over the [last six months compared to the last 2 years](https://telepresence.devstats.cncf.io/d/74/contributions-chart?orgId=1&from=now-2y&to=now).


### Maintainers

There are ten maintainers in total that work for Ambassador Labs. We have development team of four that are committing on a daily basis, with the six others committing less frequently as required.

### Adoption

The Telepresence project has seen continued growth and adoption since entering the CNCF Sandbox. Some of the known adopters as mentioned last year include:

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
* Ambassador Cloud/Developer Control Pane (https://www.getambassador.io/)

From a user standpoint, we have been engaging with the Telepresence community and growing the user base. The #telepresence slack channel has almost 5700 members after a little over two years and has become a great way for us to keep in touch with the community and receive feedback. From that feedback, we see that users are loving what Telepresence is capable of and are adopting it for their own use. A sample of comments received from our users:

* Jared Mackey: “Hey guys, I wanted to take the time to write out a big thank you for all you and your team have done so far in regards to TP2 and helping us get going with it. TP1 was already mind blowing but the improvements made in TP2 makes it next level for us here. We’re super stoked to increase adoptions of it across our org.”
* nalbury: “I’ve been using telepresence pretty heavily for awhile now and its been great.”
* Tp: “Telepresence is a great tool on my Mac & want the rest of my team to enjoy using it too. I would be happy to debug anything and help come up with a solution for WSL and Windows.”
* Graham Anderson: “I got things working, had to debug a quirk of SSHFS where it wasn’t playing nice - found a way to work around it for now but I think I’ll probably raise an issue. Otherwise this feels fantastic - amazing product, congrats!”
* Frank Gu: “Hey everyone! T2 is an amazing tool and you guys have done an excellent job reworking it with Golang! We have incorporated T2 into our dev workflow.”
* Nicolas Porta: “We are currently adopting telepresence for developing locally our micro services in k8s, and also starting to configure ArgoCD. We using 2.2 with preview urls and everything. We have around 40 devs, growing every week. Telepresence is really helping our developers to test their code hitting a real backend without having to push the code to the repository and waiting for the CICD pipeline to finish. And also helped us go from a git flow process to a trunk based process.”
* Tomas Soukup: “I am currently testing Telepresence 2. We have been testing Telepresence 1 in the past, but there were some obstacles to really put it in use at company. I like Telepresence 2 design and overall it was much cleaner and nicer to setup and also fits better in our gitops lifecycle.”
* Samuel Loury: “I discovered telepresence last year and was just mind blown. I’m convincing my team to use it now :-p Too bad it is no more python based though. Kinda love python. So just saying keep the great job!”
* Michael Foley: “First off, I want to say how great this product is. Wow!”


### Past Goals

Last year, we stated that we'd like to broaden our contributor base. Telepresence is a technically sophisticated project, and requires in-depth expertise in Kubernetes and understanding the quirks of Mac OS X and Linux networking. Some of the things we planned to address:

* Continued refactoring of the code base in to Go (which also allows us to take better advantage of Kubernetes ecosystem libraries)
* Better documentation of the Telepresence development workflow
* Continued work on the test suite

We also set out to a address a number of user requests, including:

* Better support for unreliable connectivity to the cluster
* Support for project-specific configuration
* Better documentation

Over the past year, we introduced Telepresence 2, a major rearchitecture and refactoring of the original Telepresence code base. Telepresence 2:

* Is written entirely in Go, improving contributor accessibility and Kubernetes compatibility
* Supports third-party integrations to add functionality to Telepresence; we take advantage of this plug-in architecture to display status information and such in Ambassador Cloud
* Updated the test suite and deployment pipeline, enabling a weekly release cadence
* Enabled nightly builds to improve community accessibility
* Addressed issues raised by the community

Throughout our development efforts, we have put an emphasis on improving the user and development documentation to make it easier for new and existing users to adopt Telepresence.

### Current Goals

With the work done in the past year, we believe we have created a strong foundation for the Telepresence project. Our focus this year is to grow adoption of Telepresence by removing the friction and limitations users may experience. To do this our initial areas of focus include:
* Developing a native Windows client and support Windows as a first-class operation system along with macOS and Linux
* Improve deployment in Restricted Environments
* Enable running Telepresence in CI for integration testing
* Improve performance and scalability of Telepresence

Over the past year, we have significantly increased our contribution to Telepresence. The devstats metrics above highlight the increasing number of commits over the past six months. Some of that contribution has been directed towards improving community accessibility. With an improved pipeline in place we would like to increase the number of contributors to Telepresence to accelerate development on Telepresence even more:
* Add non-Ambassador Labs maintainers
* Improve developer focused documentation to enable committers
* Increase outreach to our community and recruit committers

### How can the CNCF help?

We'd welcome help in expanding our user base by having CNCF contributors contribute and publish tutorials and improve our documentation.

### Incubation

We would like to add more non-Ambassador Labs maintainers before moving to incubation. From a product adoption and maturity standpoint, we believe we meet the incubation criteria.
