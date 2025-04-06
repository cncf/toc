# KubeVirt 2020 Annual Review

## Background

KubeVirt is a virtualization management add-on for Kubernetes. It allows a user
to run virtual machines on Kubernetes just like pods. Because KubeVirt is
integrating with the Kubernetes API and resource model, the same set of tools
can be used to manage and monitor virtual machines.

KubeVirt joined the CNCF as a sandbox project in September 2019 (ref.
[KubeVirt sandbox submission](https://github.com/cncf/toc/blob/master/proposals/sandbox/kubevirt.adoc)).

## Develoment metrics

Here are a few highligths from the project's [devstats
page](https://kubevirt.devstats.cncf.io/), looking at the 1 year period after
sandbox inclusion (Sept 2019 - Sept 2020):

 - This year we welcomed [70 new contributors](https://kubevirt.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=1567288800000&to=1598997599000) across our repositories.
 - The number of [active contributors](https://kubevirt.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&from=1567288800000&to=1598997599000&var-period=m&var-repogroup_name=All)
   increased to 152, up from 112.
 - These active contributors come from 27 different companies, up from 16.
 - The
   [activity](https://kubevirt.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&from=1567288800000&to=1598997599000&var-period=d7&var-repogroups=All)
   across repositories shows a clear steady increase.
 
### Maintainers

The KubeVirt project includes several components and there are various
repositories in the KubeVirt github org, each with its own set of core
mainainers:

 - The main [kubevirt](https://github.com/kubevirt/kubevirt) repository
   currently has 13 core maintainers, listed in its
   [OWNERS_ALIASES](https://github.com/kubevirt/kubevirt/blob/master/OWNERS_ALIASES)
   file. At this point, all of them are Red Hat employees.
 - Other repositories for components of the KubeVirt project have their own list
   of maintainers. Usually these maintainers do not overlap with the ones listed
   above. For example, the maintainers for the second and third repositories in
   terms of activity as per the dev stats are:
   - [Hyperconverged Cluster Operator](https://github.com/kubevirt/hyperconverged-cluster-operator/blob/master/OWNERS_ALIASES) has 8 approvers.
   - [Containerized Data Importer](https://github.com/kubevirt/containerized-data-importer/blob/master/OWNERS_ALIASES) has 3 approvers.


### Other metrics

Besides contributor metrics, the following numbers also highlight how the
project has gathered more interest since it joined the CNCF sandbox last year:

 - [~2200 github stars](https://github.com/kubevirt/kubevirt/stargazers), up from 1348 at the time of sandbox inclusion last year.
 - Kubernetes Slack channel members:
   - The general [#virtualization](https://kubernetes.slack.com/archives/C8ED7RKFE) channel now has 700+ members (up from 290 last year).
   - A new channel was created this year, focused on development:
     [#kubevirt-dev](https://kubernetes.slack.com/archives/C0163DT0R8X)
     currently has ~160 members.
 - the traffic on the [kubevirt.io website](https://kubevirt.io) approximately
   doubled in page views and unique visitors.

## Adoption

We have seen the increased interest mentioned in the previous section through
the various community communication channels, with a growth of breadth and depth
of participation from users and contributors.

An informal usage poll was conducted recently through
[twitter](https://twitter.com/kubevirt/status/1319251272772603904) and
[Slack](https://kubernetes.slack.com/archives/C8ED7RKFE/p1603368795347300). Of
the ~100 respondents, a majority expressed that they are _evaluating_ KubeVirt
at the moment. About 20% mentioned a _development or pre-production_ use case,
and around 5% had _production_ workloads.

As highlighted by devstats we have also seen an increase in companies
contributing to KubeVirt. Related to this, the project also has a growing number
of technology vendors who include KubeVirt technology in their offerings:

  - [OpenShift Virtualizaton](https://www.openshift.com/learn/topics/virtualization/)
  - [Gardener](https://gardener.cloud/blog/2020-10/00/)
  - [Kubermatic Virtualization](https://www.kubermatic.com/products/kubevirt/)
  - [Rancher's harvester](https://github.com/rancher/harvester/blob/766abd06561b059c1af623aacc4e505db471ceee/deploy/charts/harvester/README.md)
  - [Google Anthos](https://youtu.be/RE0A3kHT3LA?t=126)

## Project goals

### Performance against last year's goals

Since our inclusion in the CNCF sandbox our focus has been to be a welcoming
community while working to add more core features that are to be expected from a
virtualization management project.

We believe we have achieved good and consistent results through a regular,
predictable cadence of [releases](https://github.com/kubevirt/kubevirt/releases)
([once a month](https://kubevirt.io/blogs/releases.html)) that brought many new
features and improvements.

### Current goals

Our priorities for this next year are:

 - Define and release version 1.0. Work to scope this milestone release already
   started.

 - Improve the communication with end users.

 - Improve the developer experience by making PR contributions easier.

These goals translate to multiple specific work items, including:

 - A documentation revamp, working towards a more use case / workflow oriented
   approach and fully comprehensive documentation. This is one of the
   requirements for version 1.0.

 - CI improvements for an even better development environment

 - A KubeVirt event that includes a contributor oriented working space as
   well as user oriented content.

## CNCF membership

 - How can the CNCF help you achieve your upcoming goals?
 
   - The CNCF has helped the project significantly this year, through presence
     in the maintainer tracks in events and webinars as well as infrastructure
     and support services, and guidance opportunities like the technical writers
     office hours. We would like all these types of project support to continue.

 - Do you think that your project meets the [criteria for
   incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?
 
   - We believe that the project is in a healthy growing status in both user
     adoption and contributors. We are working to properly document existing
     production usage and we expect that the contributor growth will also imply
     a diverse group of maintainers.

   - With the current status and the existing plans we believe that we are on a
     good track towards incubating stage.

