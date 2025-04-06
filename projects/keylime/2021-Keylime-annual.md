# Community Health

Overall, the Keylime community is relatively small (around 20 active people in the chat) but active and welcoming. The project moves along at a steady pace, without issues and the overall status is: healthy.


## Contributions

The full devstats for Keylime can be seen [here](https://keylime.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

Here are some noteworthy elements.

In the past year (since September 2020), [26 people](https://github.com/keylime/keylime/graphs/contributors) have contributed to the main repo, [keylime/keylime](https://github.com/keylime/keylime).

Total number of PRs for all Keylime repos in the past year:

* merged: [348](https://keylime.devstats.cncf.io/d/24/prs-merged-repository-groups?orgId=1)
* total: [378](https://keylime.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)

It is worth noting that the median time for PRs to be merged is less than a week (5 days, 10 hours).

There is a weird spike of Github activity on 01-13-2021 that distorts the steady uptick of activity. This spike was caused by a glitch in our CI system that caused a multitude of failures and notifications as we were switching our CI system from TravisCI to GitHub actions. If this spike is ignored, the rest of the activities can be seen as a more steady rising slope.


## Maintainers

Keylime currently has 9 maintainers, from 7 different organisations:
- IBM, MIT, Netflix, Red Hat, Universität Bern, ZTE

Rust-keylime has 7 maintainers, from 1 organisation (Red Hat).

For more details, please refer to the MAINTAINERS.md file for keylime/keylime ([here](https://github.com/keylime/keylime/blob/master/MAINTAINERS.md)) and for keylime/rust-keylime ([here](https://github.com/keylime/rust-keylime/blob/master/MAINTAINERS)).


# Adoption

## IBM

Since last year, IBM has started using Keylime in production, at scale, on their IBM Cloud.

In practice, this means more than 5,000 servers are using Keylime on a daily basis.

This is great news for the project, as the scale of the deployment is both
impressive and a great source of real-world testing, which helps test the
codebase under pressure (and confirms Keylime is scalable).

More details on this deployment:
- [MIT article](https://news.mit.edu/2021/keylime-security-software-deployed-ibm-cloud-0727)
- [CNCF blog article](https://www.cncf.io/blog/2021/07/06/ibm-implements-remote-attestation-on-linux-with-a-hardware-root-of-trust-using-keylime/).

## Lernstick

The
[Lernstick](https://www.digitale-nachhaltigkeit.unibe.ch/services_and_support/lernstick/index_eng.html)
project provides students with a Linux-based operating system on a (USB) stick
allowing them to learn but also take exams on their own device. Lernstick is
using Keylime to guarantee the integrity of the system being loaded on student's
computers. It is currently in a proof of concept stage, but once in use it will
translate to hundreds of student machines using Keylime to prove to the exam
server that they are running a system that hasn't been tampered with.

Using Keylime in a "bring your own device" (BYOD) context is a novel and interesting use, which we are excited about. 

# Goals


## Achievement of previous goals


### Rewrite Keylime client (Keylime agent) in Rust

Good progress has been made, with a first testable version available in July 2021, which led to the discovery and fix of several bugs. Once a crucial upstream bug is fixed, the Rust agent is going to be thoroughly tested by IBM, and a first release should follow.


### vTPM

Work on vTPMs has been postponed pending some potential, necessary, future changes to the Linux kernel, but interest is still high.


### JSON Web Token (JWT) implementation

JWTs are a prerequisite for a multi-tenancy enabled Keylime. The demand for multi-tenancy has not materialised, as no one is running a Keylime in an attestation-as-a-service model (yet). The goal remains on the backlog, in the meantime, the work for multi-tenancy and JWT has been split in multiple parts, some of which has been merged into the codebase.


### Prototype of Keylime integration with Kubernetes

Preparatory work has been done, with further Proof of Concept work planned for the coming months.


## Current goals

Expand reach and adoption of Keylime.

More cloud-native integrations. This would include having published container images for all components, Helm charts for deployment to Kubernetes and a potential Kubernetes operator for the server side components.


## How can the CNCF help with current goals?

Promotion of Keylime and what it can offer to its users is something the project needs. We are seeing a better understanding of Keylime’s unique properties and value, but as a project tied to specific hardware, Keylime is sometimes seen as a solution reserved for a specific niche. TPMs are, in fact, common, and there are more people that could benefit from using Keylime than they realise.

To this end, refreshing the website could be very helpful.

More help and training for maintainers and community members on how to grow and promote the project (eg, training for live events, booths, etc). could also be helpful. For instance, members of the Keylime community participated in the virtual and live booths at KubeCon NA 2021 but were unprepared for the event as they did not know exactly what was expected of them for the event and booths. It’s possible that a lot of the potential of the event was not realised because of a lack of preparation.


# Graduating to “incubation”

We believe Keylime meets the criteria for incubation and would be happy to provide the information to support this.
