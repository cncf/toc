# Kubernetes Graduation Application

Kubernetes was the project that motivated the creation of the CNCF,
and was its first project.  It has sustained a fast pace of growth of
contributors, contributing organizations, and users, and now operates
at massive scale. The project's governance and community-management
practices continue to evolve and mature as the project grows, but the
[Kubernetes Steering Committee](https://github.com/kubernetes/steering/blob/master/README.md#members)
unanimously believes that Kubernetes fulfills all the incubating and graduation criteria.

1. Document that it is being used successfully in production by at
  least three independent end users which, in the TOC’s judgement, are
  of adequate quality and scope.
  * https://kubernetes.io/case-studies

2. Have a healthy number of committers. A committer is defined as
  someone with the commit bit; i.e., someone who can accept
  contributions to some or all of the project.
  * Kubernetes is so large, with thousands of contributors and nearly 100 repositories,
    that we had to develop our [own mechanism](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md)
    to manage approval permissions. We have hundreds of approvers, listed in more than 4000 
    [OWNERS files across the project](https://github.com/search?utf8=%E2%9C%93&q=org%3Akubernetes+filename%3AOWNERS&type=Code).

3. Demonstrate a substantial ongoing flow of commits and merged contributions
  * [Devstats](https://k8s.devstats.cncf.io/d/000000025/prs-merged-repository-groups?orgId=1&var-period=m&var-repogroups=All) shows that we have thousands of PRs merged per month.

4. Have committers from at least two organizations.
  * Yes, uncountably many, within the steering committee, maintainers, approvers, and reviewers (kubernetes/kubernetes and other repositories).
  * [PR authors by company](https://k8s.devstats.cncf.io/d/000000022/prs-authors-companies-histogram?orgId=1)
  * See also [steering committee diversity requirements](https://github.com/kubernetes/steering)

5. Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.
  * [Badge](https://bestpractices.coreinfrastructure.org/projects/569)

6. Adopt the CNCF Code of Conduct.
  * The CNCF adopted the [Kubernetes Code of Conduct](https://github.com/kubernetes/kubernetes/blob/master/code-of-conduct.md)

7. Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.
  * [Steering Committee](https://github.com/kubernetes/steering)
  * [Contributor ladder](https://github.com/kubernetes/community/blob/master/community-membership.md)
  * [SIG](https://github.com/kubernetes/community/blob/master/sig-list.md) [governance](https://github.com/kubernetes/community/blob/master/governance.md)
