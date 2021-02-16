# KEDA 2021 Annual Review

## Background

🚧 TODO

- What is KEDA
- Our journey towards CNCF
- Link to proposal

## Develoment metrics

> Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

🚧 TODO

https://keda.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m

## Maintainers

KEDA currently has 5 maintainers from 3 different companies:

| Maintainer           | GitHub ID                                     | Affiliation |
| -------------------- | --------------------------------------------- | ----------- |
| Zbynek Roubalik      | [zroubalik](https://github.com/zroubalik)     | Red Hat     |
| Tom Kerkhove         | [tomkerkhove](https://github.com/tomkerkhove) | Codit       |
| Jeff Hollan          | [jeffhollan](https://github.com/jeffhollan)   | Microsoft   |
| Anirudh Garg         | [anirudhgarg](https://github.com/anirudhgarg) | Microsoft   |
| Ahmed ElSayed        | [ahmelsayed](https://github.com/ahmelsayed)   | Microsoft   |

These have not changed since we have joined CNCF, but this is about to change in the near future.

The full maintainer list can be found in our [GitHub repository](https://github.com/kedacore/governance/blob/main/MAINTAINERS.md) along with our [governance policy](https://github.com/kedacore/governance/blob/main/GOVERNANCE.md).

## Adoption

Since we have joined as a CNCF Sandbox project we've seen a **big growth in terms of our scalers, community and users**.

When we initially joined KEDA, we provided around ~15 scalers to our users which were targetting the major cloud providers (AWS, Azure, Google Cloud) as well as popular technologies such as Prometheus and Kafke. **Over the past year, we've grown to more than 30+ scalers** supporting more cloud providers, more technologies, added various features and providing production-grade security. This effort was not only by our maintainers but mainly by our community who started contributing back more and more.

One of our focus points has been in **growing the community** by getting people to our community standups, listing who is using KEDA provide and give it a [central place in our documentation](https://keda.sh/community/). This allows people who are considering to use KEDA to get trust in the technology and start using it themselves, or learn how they are using it.

We have made it simple for people to [become listed on our website](https://github.com/kedacore/keda#become-a-listed-keda-user). Today, we are aware of the following users:
![image](https://user-images.githubusercontent.com/4345663/108028108-93c50780-702b-11eb-8f2b-69a4b87af968.png)

As part of this effort, we are working on a blog post with Alibaba Cloud to highlight how they are using KEDA and how it helps them autoscale applications on [Alibaba Cloud’s Enterprise Distributed Application Service (EDAS)](https://www.alibabacloud.com/product/edas).

As of KEDA 2.0 we allow users to practically scale anything (as long as the CRD supports `/scale` subresource), so we integrate with a variety of other projects. We are currently [considering to providing a central place](https://github.com/kedacore/governance/issues/14) for users to learn about all these various integrations so that they can get the most out of the technologies.

## Project goals

> How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

🚧 TODO

> What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

🚧 TODO

## CNCF membership

> How can the CNCF help you achieve your upcoming goals?
> 
🚧 TODO

### Incubation

> Do you think that your project meets the [criteria for incubation](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)?

Yes we do, and we are planning to propose to graduate later this month/next month an is being tracked on [GitHub](https://github.com/kedacore/governance/issues/2).
