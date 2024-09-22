# Clusterpedia 2023 Annual Review

This is the annual review for the [Clusterpedia](https://clusterpedia.io) project for 2023.

## Table of Contents

*   [Background](#background)
*   [DevStats](#devstats)
*   [Maintainers](#maintainers)
*   [Adoption](#adoption)
*   [Project perform](#project-perform)
*   [Project Goals](#project-goals)
*   [Incubation readiness](#incubation-readiness)

## Background
This name Clusterpedia is inspired by Wikipedia. It is an encyclopedia of multi-cluster to synchronize, search for, and simply control multi-cluster resources.

Clusterpedia can synchronize resources with multiple clusters and provide more powerful search features on the basis of compatibility with Kubernetes OpenAPI to help you effectively get any multi-cluster resource that you are looking for in a quick and easy way.

And clusterpedia supports exposing resource metrics in addition to the basic functionality of retrieval, like a multi-cluster Kube State Metrics, so users don't need to additionally install the kube-state-metrics component in each cluster.

It also has the following functions and features:
* Supports importing any version of kubernetes above 1.10+, and does not require upgrading with the imported cluster version.
* Very low memory usage and weak network optimization
* No dependency on specific storage components
* Automic conversion of different versions of Kube resources and support for multiple version of resources
* Automatically discover clusters managed by other multi-cloud platforms such as cluster-api, karmada, etc. to reduce manual operation and maintenance costs.

In June 2022, Clusterpedia became a CNCF Sandbox project.

## DevStats
The Clusterpedia devstats page and dashboard are found [here](https://clusterpedia.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-1y&to=now-1h)

In the last year, [36 developers have participated in clusterpedia](https://clusterpedia.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-companies=All), of which [14 are new contributors](https://clusterpedia.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now).

## Maintainers
Last year saw the addition of a maintainer from China Mobile Cloud, who, along with his team, has contributed a number of new features to Clusterpedia and is actively promoting Clusterpedia's growth.

| Maintainer  | GitHub ID        | Affiliation        | Email                            |
|-------------|------------------|--------------------|----------------------------------|
| Calvin Chen | @calvin0327      | DaoCloud           | <120380290@qq.com>               |
| Iceber Gu   | @Iceber          | DaoCloud           | <caiwei95@hotmail.com>           |
| wuyingjun   | @wuyingjun-lucky | China Mobile Cloud | <wuyingjun@cmss.chinamobile.com> |

## Adoption
The following companies use Clusterpedia in their production, infrastructure or products:
* DaoCloud
* JiBu Data
* China Mobile Cloud
* Ping An Technology
* A confidential banking and finance company

## Project perform
In the last year we've released two important versions, [0.6](https://github.com/clusterpedia-io/clusterpedia/releases/tag/v0.6.0) and [0.7](https://github.com/clusterpedia-io/clusterpedia/releases/tag/v0.7.0), as well as four fixed releases.

The following important features are included in the last year:
* Support for customized storage tiers, no longer requiring dependence on specific storage components.
* Improved support for weak network environments.
* In-memory storage layer is supported by default and provides a multi-cluster resource Watch.
* The internal storage layer supports passing sql fragments directly to filter resources.
* The internal storage layer support for sqlite.
* Support for multi-cluster kube state metrics.

In addition to the features added to the clusterpedia repository, we've also added an ElasticSearch-based storage layer plugin.

## Project Goals
We still have a lot of interesting features to develop in the coming year.

For example, we want to have better support for edge scenarios and plan to add an **MQTT-based** way to synchronize resources.

We will also provide log querying and filtering operations to make it easier for users to use.

We will also add more support for synchronization policies, such as namespace filtering or field matching during synchronization.

In [RoadMap](https://github.com/clusterpedia-io/clusterpedia/blob/main/ROADMAP.md), you can see the features we hope and plan to implement.

## Incubation readiness
We have no plans to apply for incubation yet, we still have some important features left.

After we have more maintainers and more contributors, we will consider submitting an application!
