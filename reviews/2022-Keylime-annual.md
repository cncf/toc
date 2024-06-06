# Community Health

Overall, the Keylime community is relatively small (around 20-25 active
people in the chat, with 170 members total) but growing and welcoming. The
project moves along at a steady pace, without issues and the overall
status is: healthy.

## Contributions

The full devstats for Keylime can be seen [here](https://keylime.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

Here are some noteworthy elements.

In the past year (since October 2021), [29 different people](https://github.com/keylime/keylime/graphs/contributors) have contributed to the main repo, [keylime/keylime](https://github.com/keylime/keylime).

Total number of PRs for all Keylime repos in the past year:

* merged: [362](https://keylime.devstats.cncf.io/d/24/prs-merged-repository-groups?orgId=1)
* total: [396](https://keylime.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&from=now-1y&to=now)

It is worth noting that the median time this year for PRs to be merged is just a few days (3 days, 19 hours).


## Maintainers

Keylime currently has 9 maintainers, from 7 different organisations:
- IBM, MIT, Netflix, Red Hat, Universität Bern, ZTE

Rust-keylime has 7 maintainers, from 1 organisation (Red Hat).

For more details, please refer to
the MAINTAINERS.md file for keylime/keylime ([here](https://github.com/keylime/keylime/blob/master/MAINTAINERS.md))
and for keylime/rust-keylime ([here](https://github.com/keylime/rust-keylime/blob/master/MAINTAINERS)).


# Adoption

## IBM

IBM has continued their use Keylime in production on their IBM Cloud. The
scale continues to grow

This is great news for the project, as the scale of the deployment is both
impressive and a great source of real-world testing, which helps test the
codebase under pressure (and confirms Keylime is scalable).

More details on this deployment:
- [MIT article](https://news.mit.edu/2021/keylime-security-software-deployed-ibm-cloud-0727)
- [CNCF blog article](https://www.cncf.io/blog/2021/07/06/ibm-implements-remote-attestation-on-linux-with-a-hardware-root-of-trust-using-keylime/).

IBM is also investigating using Keylime to secure client controlled hybrid
cloud deployments, where clients can consume IBM Cloud services in their
on-prem environment, and consistently deploy, manage, and control their
workloads under a shared responsibility model.

## Lernstick

The [Lernstick](https://www.digitale-nachhaltigkeit.unibe.ch/services_and_support/lernstick/index_eng.html)
project provides students with a Linux-based operating system on a
(USB) stick allowing them to learn but also take exams on their own
device. Lernstick is using Keylime to guarantee the integrity of the
system being loaded on student's computers. It is currently in a proof
of concept stage, but production usage is getting closer.  They are
planning on rolling out with the new rust agent with a first test run
planned for later this year.  Once in use it will translate to hundreds
of student machines using Keylime to prove to the exam server that they
are runnnig a system that hasn't been tampered with.

Using Keylime in a "bring your own device" (BYOD) context is a novel
and interesting use, which we are excited about.

# Goals


## Achievement of previous goals


### Rewrite Keylime client (Keylime agent) in Rust

This project was a big focus of the team for much of the year and is
now complete. The rust agent is now an officially supported Keylime
agent. It is expected that in Q4 of 2022 it will be designed as THE
official keylime client with the python agent being deprecated.

### Configuration Overhaul

We recently went through a large overhaul of the configuration system for
keylime to simplify, consolidate and make things more consistent. This
involved creating a configuration migration system to help current
deployments make the switch seamlessly. This will help new adopters as
well as people packaging keylime.

### Resiliency and Scaling

We have made a lot of progress in this are and we continue working on
hypothetical and real-life scenarios involving large numbers of attested
machines with realistic attestation workloads and scenarios. These include
customized measured boot attestation, management of IMA policies at
scale, dealing with imperfect (real-world) networking and DB management
scenarios. Our goal is to reduce/minimize the difficulty of _operating_
a keylime system; reducing “false positives” (attestation failures
caused by system errors) and eliminating “false negatives” (keylime
failure in recognizing/warning about compromised systems).


## Current goals

### Tighter Integration with Kubernetes and other Cloud Native projects.

More cloud-native integrations. This would include having published
container images for all components, Helm charts for deployment to
Kubernetes and a potential Kubernetes operator for the server side
components.

### Durable Attestation

We are introducing the concept of permanent and auditable attestation
records — a time series record of a system’s attestation state that
can be used to make provable statements about past attestations even when
the system is offline or no longer exists. The plan is to introduce record
creation and replays as optional features of the existing system. We
foresee a wider integration effort (e.g. with sigstore) to match durable
attestation with durable system policies and inventory records.

### Edge Computing Deployment Models

Keylime currently operates with a "pull" model where after agent
registration the server side components regularly pull attestation data
from the agents. But this pattern won't work well in edge scenarios
where connectivity is not always constant. We would like to have a "push"
model where agents cant periodically push attestation data to the server
side. This will be a significant amount of design and implementation work.

### Confidential Computing

While vendor solutions for attestation of confidential computing enclaves
exist, there is value in extending keylime capabilities to provide
standards-compliant runtime attestation (including durable attestation)
to such solutions. Our current work extends AMD SEV-SNP with a virtual TPM
device provisioned during boot time and enables integration with keylime.

### Extending the scope of runtime attestation to things other than IMA

One such scope is firewall management (i.e. nftables rules). While
the extension of keylime and underlying systems to support nftables is
not difficult, our current work concentrates on formulating a high-level
attestation policy that adequately represents real-life clusters.


## How can the CNCF help with current goals?

Promotion of Keylime and what it can offer to its users is something the
project needs. We are seeing a better understanding of Keylime’s unique
properties and value, but as a project tied to specific hardware, Keylime
is sometimes seen as a solution reserved for a specific niche. TPMs are,
in fact, common, and there are more people that could benefit from using
Keylime than they realise.

To this end, refreshing the website could be very helpful.

More help and training for maintainers and community members on how to
grow and promote the project (eg, training for live events, booths,
etc). could also be helpful. For instance, members of the Keylime
community participated in the virtual and live booths at KubeCon NA 2021
but were unprepared for the event as they did not know exactly what
was expected of them for the event and booths. It’s possible that a
lot of the potential of the event was not realised because of a lack
of preparation.


# Graduating to “incubation”

We believe Keylime meets the criteria for incubation and would be happy to provide the information to support this.
