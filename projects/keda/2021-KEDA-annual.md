# KEDA 2021 Annual Review

## Background

**Kubernetes Event-Driven Autoscaling (KEDA)** is a single-purpose event-driven autoscaler for Kubernetes that can be easily added to your Kubernetes cluster to scale your applications. It aims to make application auto-scaling dead-simple and optimize for cost by supporting scale-to-zero.

KEDA takes away all the scaling infrastructure and manages everything for you, allowing you to scale on 30+ systems or extend it with your scalers (pull-based, push-based, or REST API-driven).

Users only need to create a `ScaledObject` ([docs](https://keda.sh/docs/latest/concepts/scaling-deployments/)) or `ScaledJob` ([docs](https://keda.sh/docs/latest/concepts/scaling-jobs/)) that defines what you want to scale and what triggers you want to use; KEDA will handle all the rest!

![Picture1](https://user-images.githubusercontent.com/4345663/108470255-f9163400-7289-11eb-98fc-6a5f522202e0.png)

It allows you to scale anything; even if it’s a CRD from another product you are using, as long as it implements `/scale` subresource. This allows us to build an ecosystem of products to integrate with and make them more powerful and successful.

To provide production-grade security, we've introduced `TriggerAuthentication` and `ClusterTriggerAuthentication` ([docs](https://keda.sh/docs/latest/concepts/authentication/)) allowing users to move authentication information out of their application into a separate CRD. This not only allows for re-use, but operators can manage the authentication separately in a centralized place and avoid applications running with higher privileges than they need.

Although KEDA was originally started by Microsoft and Red Hat in 2019 ([announcement](https://cloudblogs.microsoft.com/opensource/2019/05/06/announcing-keda-kubernetes-event-driven-autoscaling-containers/)), the goal was to fully platform & vendor agnostic allowing us to support any major cloud vendor and products. Because of that, we've decided that it's best to donate KEDA to an open-source foundation where CNCF was the perfect fit since it aligns with the ecosystem and mindset that we have.

KEDA was proposed and accepted as a sandbox project in early 2020. ([GitHub](https://github.com/cncf/toc/issues/335)).

## Development metrics

> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

DevStats can be found [here](https://keda.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

## Maintainers

KEDA currently has 5 maintainers from 3 different companies:

| Maintainer           | GitHub ID                                     | Affiliation |
| -------------------- | --------------------------------------------- | ----------- |
| Zbynek Roubalik      | [zroubalik](https://github.com/zroubalik)     | Red Hat     |
| Tom Kerkhove         | [tomkerkhove](https://github.com/tomkerkhove) | Codit       |
| Jeff Hollan          | [jeffhollan](https://github.com/jeffhollan)   | Microsoft   |
| Anirudh Garg         | [anirudhgarg](https://github.com/anirudhgarg) | Microsoft   |
| Ahmed ElSayed        | [ahmelsayed](https://github.com/ahmelsayed)   | Microsoft   |

But Anirudh [has decided to step down](https://github.com/kedacore/governance/issues/17) and are starting this process now.

The full maintainer list can be found in our [GitHub repository](https://github.com/kedacore/governance/blob/main/MAINTAINERS.md) along with our [governance policy](https://github.com/kedacore/governance/blob/main/GOVERNANCE.md).

## Adoption

Since we have joined as a CNCF Sandbox project we've seen a **big growth in terms of our scalers, community, and users**.

When we initially joined KEDA, we provided around ~15 scalers to our users which were targeting the major cloud providers (AWS, Azure, Google Cloud) as well as popular technologies such as Prometheus and Kafka. **Over the past year, we've grown to more than 30+ scalers** supporting more cloud providers, more technologies, added various features, and providing production-grade security. This effort was not only by our maintainers but mainly by our community who started contributing back more and more.

One of our focus points has been in **growing the community** by getting people to our community standups, listing who is using KEDA provide, and give it a [central place in our documentation](https://keda.sh/community/). This allows people who are considering to use KEDA to get trust in the technology and start using it themselves or learn how they are using it.

We have made it simple for people to [become listed on our website](https://github.com/kedacore/keda#become-a-listed-keda-user). Today, we are aware of the following users:

![EukwAwQXUAMebyQ](https://user-images.githubusercontent.com/4345663/108509240-4790f600-72bd-11eb-9322-58c72f066b17.jpg)

As part of this effort, we are working on a blog post with Alibaba Cloud to highlight how they are using KEDA and how it helps them autoscale applications on [Alibaba Cloud’s Enterprise Distributed Application Service (EDAS)](https://www.alibabacloud.com/product/edas).

As of KEDA 2.0, we allow users to practically scale anything (as long as the CRD supports `/scale` subresource), so we integrate with a variety of other projects. We are currently [considering providing a central place](https://github.com/kedacore/governance/issues/14) for users to learn about all these various integrations so that they can get the most out of the technologies.

In terms of our artifacts on [Docker Hub](https://hub.docker.com/r/kedacore/keda), we have passed 6.5M downloads of our operator and have a steady growth.

Unfortunately, it is hard to measure adoption since Docker Hub only provides limited metrics, so we've built our own dashboard which scrapes the Docker Hub API to understand the growth over time:

![image](https://user-images.githubusercontent.com/4345663/108208162-f51ed080-7128-11eb-8bbe-d05a8bc35b5e.png)

> 💡 Because of the limited metrics, we are moving to GitHub Container Registry which gives a lot more insights. ([GitHub](https://github.com/kedacore/governance/issues/16))

## Project goals

> How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

That said, our main focus was on adding more scalers, making it easier to use KEDA and improve the documentation as well as improve our release process.

We have managed to finalize all of them:

- Our release process is mainly automated and fully documented on [GitHub](https://github.com/kedacore/keda/blob/main/RELEASE-PROCESS.MD)
  - We have different release captains every release to dogfood it and automate where we can
- We've made it easier to use KEDA on a few fronts:
  - Our website is revamped (thanks to CNCF) which is now full version, more usable, provides search, etc.
  - Contributing to the docs is simpler and are using the "Every change is just a PR away", along with PR previews of the changes
  - We have introduced more consistent scaler documentation and configuration so that it's more straight-forward
  - We have grown our samples catalog and are doing more talks on how to use KEDA
- We are constantly adding scaler to every new release, either built by our core team or contributed by the community.
Here is an overview of the growth of our scaler catalog:

![Scaler Growth](https://user-images.githubusercontent.com/4345663/108472973-b5252e00-728d-11eb-8d06-3d7ae07e3b77.png)

> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

We are running our project in a fairly open & flexible model so that we can react to the needs of our community, but we have recently introduced a [high-level roadmap](https://github.com/kedacore/keda/projects/3) on where we are going.

Our major investment area is **HTTP-based autoscaling**. There is a high demand for this and a big gap in our offering, so we are currently building an [experimental add-on](https://github.com/kedacore/http-add-on) to have an early approach and see how we need to evolve.

As part of this effort, we want to rely on an open specification such as [Service Mesh Interface spec](https://smi-spec.io/) which is exactly what we need, but it does not cover Ingress nor Service-to-Service communication. Because of that, we are [proposing to incorporate this in the new Gateway API](https://github.com/kubernetes-sigs/gateway-api/issues/554) from the get-go so that other products can easily build vendor-agnostic integrations.

That said, this seems to be a gap where we should consider building a Traffic Metrics API spec for all things HTTP/gRPC in Kubernetes which covers service mesh, service-to-service & ingress/gateway which is what we'll further contribute to, where possible.

Next to HTTP-based autoscaling, we will continue to focus on making KEDA more powerful by:

- Adding new scalers to our catalog
- Making the current ones more powerful (more authentication options, etc.)
- Support more secret providers for our trigger authentication (ie. Azure Key Vault, etc.)
- Integrate with other projects and are considering providing an overview to our users ([GitHub](https://github.com/kedacore/governance/issues/14))

Last but not least, we will keep on focussing on adoption and keep on growing our community around KEDA.

## CNCF membership

> How can the CNCF help you achieve your upcoming goals?

CNCF allows us to focus on our core business and helps us manage the project itself.

We've consulted the CNCF for legal questions, creative services to design a new website for use, recommendations on tooling, and more.

This is a tremendous help for us and supports us in making autoscaling simpler.

### Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#incubating-stage)?

Yes, we do, and we are planning to propose to graduate later this month/next month and is being tracked on [GitHub](https://github.com/kedacore/governance/issues/2).
