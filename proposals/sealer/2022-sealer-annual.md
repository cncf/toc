# Sealer 2022 Annual Review

## Background

Sealer provides a new way of distributed application delivery which is reducing the difficulty and complexity by
packaging Kubernetes cluster and all application's dependencies into one Sealer Image.

Sealer joined the CNCF on April 26, 2022. Project website is: [http://sealer.cool/](http://sealer.cool/)

## Include a link to your project’s devstats page.

CNCF [devstats](https://sealer.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m)

- According to devstats, Sealer currently has [76](https://sealer.devstats.cncf.io/d/22/prs-authors-table?orgId=1)
  contributors from [21](https://sealer.devstats.cncf.io/d/5/companies-table?orgId=1) companies.
- [549](https://sealer.devstats.cncf.io/d/15/new-prs-in-repository-groups?orgId=1&var-period=w&var-repogroup_name=All)
  new PRs was opened in last year . On average, there
  are [52.5](https://sealer.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=commits&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now)
  commits per month, contained
  within [43.4](https://sealer.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-2y&to=now)
  merged PRs per month.
- The community has grown since the project entered the CNCF sandbox.
    - Number of contributors: 50 -> **76**
    - Github stars: 1096 -> **1862**
    - Github forks: 236 -> **380**
    - Contributing organizations: 12 -> **21**

## How many maintainers do you have, and which organizations are they from?

We currently have 76 contributors and 9 maintainers.

| Maintainer | Affiliation | GitHub ID | Email |
| --- | --- | --- | --- |
| Haitao Fang | Alibaba Group | [fanux](https://github.com/fanux) | [fhtjob@hotmail.com](mailto:fhtjob@hotmail.com) |
| Hongliang Sun | Alibaba Group | [allencloud](https://github.com/allencloud) | [shlallen1990@gmail.com](mailto:shlallen1990@gmail.com) |
| Jiahui Huai | Boyan | [bxy4543](https://github.com/bxy4543) | [4543bxy@gmail.com](mailto:4543bxy@gmail.com) |
| Lei Zhou | Boyan | [kakaZhou719](https://github.com/kakaZhou719) | [kakazhou719@163.com](mailto:kakazhou719@163.com) |
| Jinrui Cui | Realai | [cuisongliu](https://github.com/cuisongliu) | [cuisongliu@qq.com](mailto:cuisongliu@qq.com) |
| Jiangnan Bao | ZheJiang University | [justadogistaken](https://github.com/justadogistaken) | [baojn1998@163.com](mailto:baojn1998@163.com) |
| Xun Wang | ZhengcaiCloud | [lllwan](https://github.com/lllwan) | [lllwan@vip.qq.com](mailto:lllwan@vip.qq.com) |
| Vince Cui | Alibaba Group | [vincecui](https://github.com/vincecui) | [huaiyou.cyz@alibaba-inc.com](mailto:huaiyou.cyz@alibaba-inc.com) |
| Starnop | Alibaba Group | [starnop](https://github.com/starnop) | [starnop@163.com](mailto:starnop@163.com) |

## What do you know about adoption, and how has this changed since your last review / since you joined Sandbox?

We are currently aware of 7 companies who adopted Sealer into their production environment, and an additional
some [companies](https://github.com/sealerio/sealer/issues/119) who are currently evaluating it. We are in the process
of asking them to be added to the [ADOPTERS](https://github.com/sealerio/sealer/blob/main/Adopters.md) file.

## How has the project performed against its goals since the last review?

Project Features

1. Sealer spec definition stabilization,including Kubefile and Clusterfile.
2. Provide standard OCI format sealer image,for conveniently play with build share run.
3. Provide complete cluster lifecycle management,such as run,scale up,delete,upgrade,rollback.
4. Propose the concept of sealer app, design and implement its definition.
5. Provide richer design and usage document, as well as the version management.

Community

Opening 4 additional project on sealerio organization to stabilize the entire toolchain of the project.

- [sealer-basefs](https://github.com/sealerio/basefs)
- [sealer-plugins](https://github.com/sealerio/plugins)
- [sealer-applications](https://github.com/sealerio/applications)
- [sealer-cluster-spec](https://github.com/sealerio/cluster-spec)

Ecology

The number of [official sealer image](https://hub.docker.com/r/sealerio/kubernetes) downloads from Docker hub exceeds 10k+ over in the past half year.

## What are the current goals of the project?

five main goals for 2023 we will focus on:

- enrich and improve the basic capabilities of the Sealer Image
- improve the delivery capabilities of distributed application
- build more stable and easy-to-use cluster delivery capabilities
- performance improvement
- more contributors and adopters

More information on these goals can be found [here](https://github.com/sealerio/sealer/blob/main/ROADMAP.md)

## How can the CNCF help you achieve your upcoming goals?

There are several vectors we think CNCF can help the project:

- the growth of sealer community. with the influence of CNCF's community, we want to attract more international
  contributors and users for sealer community.
- more opportunities to collaborate with other CNCF projects and even to other foundations.

## Do you think that your project meets the criteria for incubation?

Sealer is not yet ready for incubation since our community ecosystem is not prosperous enough , and we would like to
see more international contributors and adopters in the future.
