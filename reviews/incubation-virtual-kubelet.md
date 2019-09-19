# Virtual kubelet Incubating Stage Review

Virtual kubelet is currently a sandbox stage project.  Please refer to virtual kubelet's [sandbox stage proposal](../proposals/virtualkubelet.adoc) ("inception" at time of acceptance) for details on the sandbox requirements.

In the time since being accepted to the sandbox stage (about 8 months ago), virtual kubelet has demonstrated healthy growth and progress.
Four releases were completed which included v0.9.0, v0.10.0, v0.11.0, v1.0.0. In the last release, virtual kubelet was declared ready for production use. Providers were moved out of tree in v0.11.0 thus each provider is expected to manage their own releases because virtual kubelet core will not release binaries anymore. Virtual kubelet has also standarized on community guidelines and created a process for providers to maintain their respective providers within the virtual kubelet organization. 

A [formalized governance policy](https://github.com/virtual-kubelet/community/blob/master/governance.md) has been approved and instituted for the project.

## Incubating Stage Criteria

To be accepted to incubating stage, a project must meet the sandbox stage requirements plus:

* Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

  * Adopters: [https://github.com/rook/rook/blob/master/ADOPTERS.md](https://github.com/rook/rook/blob/master/ADOPTERS.md)

* Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

  * Maintainers of the project are listed in [https://github.com/virtual-kubelet/community/blob/master/OWNERS.md](https://github.com/virtual-kubelet/community/blob/master/OWNERS.md).

  * Maintainers are added and removed from the project as per the policies outlined in the project governance: [https://github.com/virtual-kubelet/community/blob/master/governance.md](https://github.com/virtual-kubelet/community/blob/master/governance.md).

* Demonstrate a substantial ongoing flow of commits and merged contributions.

  * Releases: [https://github.com/virtual-kubelet/virtual-kubelet/releases](https://github.com/virtual-kubelet/virtual-kubelet/releases)

  * Roadmap in the form of a milestone: [https://github.com/virtual-kubelet/virtual-kubelet/milestone/2](https://github.com/virtual-kubelet/virtual-kubelet/milestone/2)

  * Contributors: [https://github.com/virtual-kubelet/virtual-kubelet/graphs/contributors](https://github.com/virtual-kubelet/virtual-kubelet/graphs/contributors)

  * Commit activity: [https://github.com/virtual-kubelet/virtual-kubelet/graphs/commit-activity](https://github.com/virtual-kubelet/virtual-kubelet/graphs/commit-activity)

  * CNCF DevStats: [https://virtualkubelet.devstats.cncf.io/](https://virtualkubelet.devstats.cncf.io)
    * [Community Stats](https://virtualkubelet.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1)

Further details of Virtual Kubelet's growth and progress since entering the sandbox stage as well as use case details from the Virtual Kubelet community can be found in this [slide deck](https://docs.google.com/presentation/d/1kvZL0Ep3JHzgQsqTAcvNo6FTqIt4XMZi3Ny49bdP2Gw/edit?usp=sharing).
