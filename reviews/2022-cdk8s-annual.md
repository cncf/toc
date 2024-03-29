## Background

CDK For Kubernetes (cdk8s) is an open-source software development framework for defining Kubernetes 
applications and reusable abstractions. Users leverage familiar programming languages and rich object-oriented APIs to 
create standard Kubernetes manifests which can be applied to any Kubernetes cluster.

CDK8s joined the CNCF on November 10, 2020.

## Annual Review

The following contains the requested content for the annual review, as described [here](https://github.com/cncf/toc/blob/main/process/sandbox-annual-review.md).

### Include a link to your project’s devstats page.

The full devstats page for cdk8s is available here: https://cdk8s.devstats.cncf.io/d/8/dashboards

* In the past year, there were 7 [community contributions](https://cdk8s.devstats.cncf.io/d/56/company-commits-table?orgId=1&from=now-1y&to=now&var-repogroups=All&var-companies=Palo%20Alto%20Networks&var-companies=Cisco&var-companies=GridLine) 
from 3 different users, in 3 different organizations (Palo Alto Networks, GridLine, Cisco). There is an additional [contribution](https://github.com/cdk8s-team/cdk8s-cli/pull/324) 
that doesn’t show up on devstats, probably because the user isn’t associated with any company.

### How many maintainers do you have, and which organizations are they from?

The project currently has 6 maintainers, all from AWS:

* Eli Polonsky [@iliapolo](https://github.com/iliapolo)
* Vinayak Kukreja [@vinayak-kukreja](https://github.com/vinayak-kukreja)
* Sumu Pitchayan [@sumupitchayan](https://github.com/sumupitchayan)
* Evgeny Karasik [@evgenyka](https://github.com/evgenyka)
* Uday Pant [@udaypant](https://github.com/udaypant)
* Nathan Taber [@tabern](https://github.com/tabern)

### What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?

We are currently aware of 7 companies who adopted cdk8s into their production workloads, 
and an additional 9 who are currently evaluating it. We are in the process of asking them 
to be added to the [ADOPTERS](https://github.com/cdk8s-team/cdk8s/blob/master/ADOPTERS.md) file.

We are also seeing signs of increasing adoption by observing [NPM download stats](https://npm-stat.com/charts.html?package=cdk8s&from=2022-01-01&to=2023-02-23). 
They show that over the last year, we have grown from ~12k to ~129k downloads a week for the core library. 
GitHub stars have also increased from 2280 to [3492](https://star-history.com/#cdk8s-team/cdk8s&Date) stars in the course of the last 15 months. 

In addition, the #cdk8s channel in the [cdk.dev](https://cdk.dev/) workspace now contains ~7k members and we 
see consistent engagement and knowledge sharing. The #cdk8s channel in the CNCF workspace has 54 members and is mostly dormant. 
This is expected because [cdk.dev](https://cdk.dev/) existed before #cdk8s was accepted to the CNCF, and is the regarded as 
the main communication channel for all the CDK eco-system.

### How has the project performed against its goals since the last review?

Our goals for the project since joining sandbox were:

1. Stabilize the entire toolchain of the project. This was important for our customers to be able 
to adopt the tool in their production pipelines. We have successfully done so and released a major version 
for all of components of the project: 

    * [cdk8s-cli](https://github.com/cdk8s-team/cdk8s-cli)
    * [cdk8s-core](https://github.com/cdk8s-team/cdk8s-core)
    * [cdk8s-plus](https://github.com/cdk8s-team/cdk8s-plus)

2. Add support for Golang. Users have expressed desire to author cdk8s in applications in Golang, 
it was in-fact our mostly highly voted [issue](https://github.com/cdk8s-team/cdk8s/issues/119) in the repository.
3. Help customers safe guard their Kubernetes workloads by integrating cdk8s into policy compliant pipelines 
and applying policy-as-code tools on the manifests generated by cdk8s. To that end, we created a spec and 
a pluggable mechanism in the cdk8s-cli that allows third-party vendors to implement plugins for their 
tools (e.g Kyverno, Checkov, Kubescape, ...). We have partnered with [Datree](https://www.datree.io/) to create a reference 
implementation of one such plugin, available [here](https://github.com/datreeio/datree-cdk8s).

Official announcements have been made via the AWS blog:

* https://aws.amazon.com/blogs/containers/announcing-the-general-availability-of-cdk8s-and-support-for-go/
* https://aws.amazon.com/blogs/containers/announcing-general-availability-of-cdk8s-plus-and-support-for-manifest-validation/

### What are the current goals of the project?

Our goals for 2023 focus on two themes:

* Make it easier for customers to integrate cdk8s into the existing Kubernetes ecosystem (e.g Helm, ArgoCD, Flux, ...). 
We'll create foundational content on the cdk8s framework, how-to guides and tutorials that show customers 
how to integrate cdk8s into their CI/CD pipelines.
* Make it easier for customers to author complex Kubernetes manifests that leverage cloud services. We will integrate with the AWS CDK and CDK For Terraform to enable applications that contain both Kubernetes and cloud resources.

> More information on these goals can be found [here](https://github.com/cdk8s-team/cdk8s/discussions?discussions_q=label%3Aplanning%2F2023+).

In addition, we would like to cultivate the community, encourage more contributions, and more contributors that 
are versed in the codebase. To that end, we will hold monthly community meetings where we will discuss ongoing 
issues and feature requests for the product. We will focus more on community outreach and educational material in 
the form of videos and joint coding sessions.

### How can the CNCF help you achieve your upcoming goals?

There are several vectors we think CNCF can help the project:

* Help increase cdk8s market presence. Based on our experience, customers who are aware of cdk8s appreciate its 
advantages and enjoy using it. However, the challenge lies in the fact that cdk8s is not widely recognized, 
particularly in larger organizations. While we understand that CNCF does not offer marketing support for sandbox projects, 
we would appreciate some guidance in this regard.
* Our goal is to grow the cdk8s community. If CNCF has any best practices from successful CNCF projects, 
we would like to leverage them for cdk8s. For example, what is discussed during community meetings? how is public 
planning done? how should we present a roadmap? what are some common contribution obstacles we should be aware of?
* Do you have some recommendations on how a project such as cdk8s, i.e client side only, can collect usage data? Currently our 
sources for this are download statistics, but those don’t really map to distinct usage patterns.

### Do you think that your project meets the criteria for incubation?

We believe we might be ready for incubation by the end of the year. We are working on a draft due diligence document 
to start the process and dive deeper into the requirements. Any insight into this would be most appreciated.
