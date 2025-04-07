# Telepresence Sandbox Annual Review

The goal of the Telepresence project is to accelerate the "developer inner loop" for cloud-native application development on Kubernetes. Telepresence achieves this by creating a dynamic network bridge between a local development environment (e.g., a laptop) and a remote Kubernetes cluster.

Telepresence was presented to the TOC on April 2018. Please see the project's [Sandbox proposal](https://github.com/cncf/toc/blob/master/proposals/sandbox/telepresence.adoc).

## Annual Review Questions

### Metrics

| Metric       |  Pre-Sandbox |     2020     |    2021      |   Today   |
| ------------ | ------------ | ------------ | ------------ | --------- |
| Docker Pulls | 100K+        | 10M+         | 10M+         | 10M+      |
| Stars        | 700+         | 2700+        | 4000         | 5.1k      |
| Forks        | ~50          | ~190         | ~300         | 400+      |

The CNCF provides Telepresence Github metrics through its [devstats site](https://telepresence.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

### Maintainers

There are eight maintainers total. Of these, four commit to the project on a daily basis and are employed by Ambassador Labs.
The remaining four are three Ambassador Labs employees and one external maintainer, all of whom commit less frequently as needed.

### Adoption

The Telepresence project has seen continued growth and adoption since entering the CNCF Sandbox. In addition to the adopters mentioned last year:

* Avaya
* Bose
* Engel & Volkers (KubeCon EU 2019 presentation: https://www.youtube.com/watch?v=tD0FIlxO1AQ)
* InVision
* Namely
* PTC
* Red Hat
* Timecampus
* ThetaRay

We have added new adopters such as:

* Salt Security
* MoneyForward
* NetApp
* Manhattan Associates
* SKU IQ


There have also been a number of integrations with other software stacks including:

* wad.sh (https://blog.sebastian-daschner.com/entries/reloading-apps-kubernetes-telepresence-wad)
* Ambassador Edge Stack (https://www.getambassador.io/docs/latest/topics/using/edgectl/)
* Kelda (https://kelda.io/docs/tutorials/telepresence/)
* Ambassador Cloud/Developer Control Pane (https://www.getambassador.io/)

From a user standpoint, we have been engaging with the Telepresence community and growing the user base. The #telepresence slack channel has around 7700 members now and continues to be a great way for us to keep in touch with the community and receive feedback. We hold a live troubleshooting session through the channel weekly and usually work through 1-2 users issues. We also hold a monthly Maintainers meeting attended by our development team and open to anyone to ask questions, discuss features, etc.

Some feedback we have received through these channels include:

Brooke Weaver "I feel comfortable both using and troubleshooting Telepresence since the docs are thorough and very helpful. I love how straightforward and easy to use it is! I’m a huge fan of the product; it makes my development process so much easier."

Jared Mackey "Hey guys, I wanted to take the time to write out a big thank you for all you and your team have done so far in regards to TP2 and helping us get going with it. TP1 was already mind blowing but the improvements made in TP2 makes it next level for us here. We’re super stoked to increase adoptions of it across our org."

Frank Gu "T2 is an amazing tool and you guys have done an excellent job reworking it with Golang! We have incorporated T2 into our dev workflow."

Ao Chen "I'm excited to find such a wonderful tool as telepresence. Thanks and telepresence is a really good tool!"


### Past Goals

Last year we stated that our primary goal was to grow our userbase by removing barriers to adoption, and presented the following goals:

* Developing a native Windows client and support Windows as a first-class operation system along with macOS and Linux
* Improve deployment in Restricted Environments
* Enable running Telepresence in CI for integration testing
* Improve performance and scalability of Telepresence

The first three of these goals have been delivered into the project, and although scalability and performance remain areas of focus we've delivered considerable improvements, including:
* Limiting the number of API calls made to kubernetes
* Eliminating CPU bottlenecks
* Refactoring our networking code to better deal with concurrent connections



### Current Goals

We feel that the project can offer considerable value for users of kubernetes, but can be difficult to adopt because not every kubernetes stack plays nice with it. Our focus for this year is to make Telepresence easier for organizations, particularly those with complex environments, to adopt. Areas of focus will include:
* Telepresence should always be compatible with service meshes, not just when the happy path is followed.
* Telepresence intercept must properly emulate and/or pass-through all features exposed by cloud providers, particularly cloud credentials (e.g. obtained via IRSA) must be made available to the local workload.
* Telepresence and its traffic agent should not disturb existing network policies and checks. This includes providing network policies for sysadmins with tight firewall requirements, and instructions to configure load balancer sanity checks when the traffic agent is present.
* Telepresence must be easy to deploy in the context of a large cluster with many applications. It should be easy to set up an Argo CD deployment of Telepresence, for example, and cluster admins should have clarity on which config levers they can tweak, as well as the ability to push custer-wide Telepresence configuration to the users' machines.

### How can the CNCF help?

We'd welcome help in expanding our user base by having CNCF contributors contribute and publish tutorials and improve our documentation.
We'd especially welcome contributors who have expertise with complex environments or wish to share how they got Telepresence to work for their use-case and integrated with other tools, particularly cloud native ones.

### Incubation

We would still like to add more non-Ambassador Labs maintainers and better document using Telepresence together with other cloud native tools before moving to incubation. From a product adoption and maturity standpoint, we believe we meet the incubation criteria.
