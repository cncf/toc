# Porter 2022 Annual Review

[Porter](https://getporter.org) creates installers for cloud-native applications and their infrastructure.
Package your application, client tools, configuration, and deployment logic into an installer that you can distribute and run with a single command.

## Metrics

| Metric    | Last Review | Today   |
| --------- | ----------- | ------- |
| Stars     | 658         | 907     |
| Forks     | 116         | 150     |
| Downloads | 33,367      | 79,786* |

DevStats: https://porter.devstats.cncf.io  
Download Stats: https://hanadigital.github.io/grev/?user=getporter&repo=porter

\* Downloads are calculated as artifacts downloaded per release. There is not a 1:1 correlation between the download number and the number of users because each release has multiple artifacts that the user may choose to download, such as the Porter client, Porter runtime, mixin client/runtime, etc. One heuristic would be to divide the number by 5 to estimate the number of times a Porter release has been downloaded.

## Maintainers

https://github.com/getporter/porter/blob/main/OWNERS.md

* Carolyn Van Slyck (Microsoft)
* Jeremy Rickard (Microsoft)
* Vaughn Dice (Fermyon)
* Reddy Prasad (Independent)
* Jennifer Davis (Google)
* Yingrong Zhao (Microsoft)
* Brian DeGeeter (F5)
* Steven Gettys (F5)

Since last year, we have added 3 new core maintainers, who are maintainers for the entire getporter organization. We now have 5 active* core maintainers: Carolyn, Yingrong, Vaughn, Brian, and Steven. Our project has lots of smaller repositories for our plugins and mixins which can have additional maintainers. We also have 1 active repository level maintainer for the helm mixin: Mohamed Chorfa (Independent).

\* One of our 2023 goals is to define a process for promoting inactive maintainers to emeritus status.

## Adoption

https://github.com/getporter/porter/blob/main/ADOPTERS.md

* F5 has joined our list of public adopters, and two of their developers are now core maintainers.
* Microsoft recently announced that the Azure Marketplace now allows vendors to distribute their solutions with Porter bundles, a.k.a [Kubernetes Apps Marketplace Offering][k8s-offer]. [Tetrate deploying Istio to Azure Kubernetes Service][tetrate] (AKS) and [Elastic deploying Elastic Cloud to AKS][elastic] are two examples of how vendors are already taking advantage of the integration.
* [EY released a case study][ey] on how they are using Porter to automate rolling out environments for each of their customers which has dramatically saved time and reduced the complexity of deployments. EY Engineering leadership has provided a statement regarding their use of Porter, below:

   > "It used to take a day for the dev team to roll out one to two environments. We’ve gone from that to the ability to parallelize our deployment pipeline to target thousands of instances overnight at the push of a button. The potential cost savings are significant given the scale we're talking about."  
   > Ken Priyadarshi  
   > EY Global Chief Architect

## Goals

Our 2022 goals were to:

* Release v1.0.0, which went [live the first week in October][v1-announcement]! 🎉
* Complete the [Porter Operator]. We have made progress on the operator but due to shifting resources, we now expect to release it early in 2023.
* Promote more contributors to maintainers. We have added 3 core contributors to Porter, two of which are from F5 (Porter was initially created at Microsoft). We are pleased to have a more diverse set of maintainers and companies who are investing in Porter.
* We worked with the CNCF to understand how we could improve our documentation, creating not just reference documentation but also conceptual and task based docs. While we don't have technical writers, our developers have done a great job addressing gaps in our documentation.

For 2023 we hope to:

### See Users Upgrade to v1

A stable v0 Porter release has been available and used in production by a number of large adopters for over a year. Some have already migrated to v1 and we would like to see the rest of our user base upgrade to v1. In the few weeks since releasing v1, we have seen the download counts for v1.0.0 reach over 65% of the download count for our most recent v0 release (v0.38.13).

### Improve Contributor Pipeline

We have added three core contributors this year but plan to maintain the contributor pipeline to account for inevitable churn by:

* **Increase the percentage of new contributors who go on to become episodic contributors.** We need to keep a consistent backlog of curated issues good for new contributors and people who are looking for direction on what to work on next. Reviewers should also engage with a contributor after their pull request is finished to help them find that next issue.
* **Make the path to becoming a maintainer more clear, with improved new maintainer onboarding, so that we can keep our maintainer pipeline healthy**. We have a contributor ladder and some documentation for maintainers that people have used this year to go from contributors to maintainers and we should incorporate frequent questions and feedback into the documents.
* **Define "active" maintainers and the process for a maintainer moving from active to emeritus**. We should recognize that some of our maintainers have contributed a lot in the past but have moved on, while crediting them for their contributions.

### Release the Porter Operator

Many users, especially adopters with large-scale Porter deployments, are eagerly awaiting managing their installations using a GitOps based workflow with the Porter Operator and Flux. We made great progress on the operator in the past year. It works but isn't ready for use in production yet. Our goal is to [release a stable version](https://github.com/getporter/operator/milestone/1) of the Porter Operator once the maintainers for that project are available again to continue development. 

### Support Advanced Dependencies and Workflows

Porter plans to support advanced workflows and dependency scenarios ([PEP003](https://github.com/getporter/proposals/blob/main/pep/003-dependency-namespaces-and-labels.md)), where users can define bundles that have a complex dependency graph. Dependencies may be an interface such as requiring a MySQL database whether it came from a dedicated server, a Helm chart, or a database as a service from a cloud provider. Dependencies may also be resolved to existing installations of bundles, such as a shared dev database or a redis instance for the staging environment.

This allows better reuse of bundles, so that a bundle author can write an bundle for just their software and they don't need to write cloud provider specific bundles, such as "my software on Azure", "my software on Amazon", etc. Instead, they can create a bundle that requires another bundle that matches a specified interface and Porter can satisfy that dependency differently depending on the target environment.

### Improve Bundle Security

Mixins are binary wrappers around common deployment tools, such as Terraform and Helm, that make it easier to use that tool inside a bundle with Porter.
Currently they are distributed as binaries, and embedded inside the bundle's image.

We want to improve bundle security by distributing and executing **mixins as bundles**, essentially breaking up bundles into smaller bundles that we can more securely distribute and execute ([PEP005](https://github.com/getporter/proposals/blob/main/pep/005-mixins-are-bundles.md)). This has a number of advantages such as:
  * Secure distribution mechanism of mixins. This allows us to no longer need to "install" mixins, or trust those installed mixins. Bundles can reference mixins from OCI registries and take advantage of existing security tooling and software to distribute, scan, and sign mixins.
  * Isolate credentials used by a bundle to just the relevant portions of the bundle. For example, only the Azure mixin should have access to your Azure credentials, and only the Helm mixin should have access to your kubeconfig. Credentials applicable to one mixin are not exposed to other unrelated mixins or scripts in the bundle.
  * Achieve significant performance improvements due to improved reuse of cached image layers.

# How can the CNCF help?

We plan to do a pre-recorded CNCF webinar to help get more people familiar with bundles and the types of problems that Porter solves.

# Incubation

We plan to apply for Incubation next year. We believe the Porter project has several important properties that make it a good candidate for incubation, including the following:

* A steady stream of improvements and releases throughout the year with lots more planned ahead.
* Three public adopters of interesting size/scope.
* Maintainers from multiple companies.

We still need to develop a documented security process and prep for a due-diligence review.

[v1-announcement]: https://getporter.org/blog/v1-is-here
[Porter Operator]: https://getporter.org/operator/
[k8s-offer]: https://techcommunity.microsoft.com/t5/apps-on-azure-blog/releasing-kubernetes-apps-offer-in-microsoft-cloud-marketplace/ba-p/3650628
[tetrate]: https://www.tetrate.io/blog/tetrate-istio-distro-on-the-azure-marketplace/
[elastic]: https://www.elastic.co/blog/maximizing-the-experience-of-elastic-cloud-on-kubernetes-with-azure-kubernetes-service
[ey]: https://customers.microsoft.com/en-us/story/1491491165781122603-ey-professionalservices
