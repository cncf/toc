# Curve 2023 Annual Review

This is the annual review for the [Curve](https://opencurve.io) project for 2023.

## Table of Contents

*   [Curve 2023 Annual Review](#curve-2023-annual-review)
*   [Table of Contents](#table-of-contents)
*   [Background](#background)
*   [DevStats](#devstats)
*   [Maintainers](#maintainers)
*   [Adoption](#adoption)
*   [Project perform](#project-perform)
*   [Project Goals](#project-goals)
*   [Incubation readiness](#incubation-readiness)

## Background

Curve is a modern distributed storage system developed by netease, currently supporting file storage(CurveFS) and block storage(CurveBS). Now it's hosted at CNCF as a sandbox project.

The core application scenarios of CurveBS mainly include:

*   the performance, mixed, capacity cloud disk or persistent volume of virtual machine/container, and remote disks of physical machines
*   high-performance separation of storage and computation architecture: high-performance and low latency architecture based on RDMA+SPDK, supporting the separation deployment structure of various databases such as MySQL and Kafka

The core application scenarios of CurveFS mainly include:

*   the cost-effective storage in AI training scene
*   the hot and cold data automation layered storage in big data scenarios
*   the cost-effective shared file storage on the public cloud: It can be used for business scenarios such as AI, big data, file sharing
*   Hybrid storage: Hot data is stored in the local IDC, cold data is stored in public cloud

Curve Architecture：

<img src="https://github.com/opencurve/curve/blob/21e7dbcf65b36941f42fdb0a858ddbf0c15b1a27/docs/images/Curve-arch.png?raw=true" width="500" height="600" />



## DevStats

Curve DevStats dashboard [here](https://curve.devstats.cncf.io/d/1/activity-repository-groups?orgId=1\&from=now-1y\&to=now\&var-period=d7\&var-repogroups=All) shows an increase in contribution activities. Last year [100+ developers](https://curve.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1\&var-period_name=Last%20year\&var-metric=contributions\&var-repogroup_name=All\&var-country_name=All\&var-companies=All) contribute to Curve.


## Maintainers

*   No new Maintainer at last year and one retired due to work change.
*   3 contributors NOT from NetEase became committers at last year, we hope they can became Maintainers in this year.

| Maintainer  | GitHub Username                              | Company |
| ----------- | -------------------------------------------- | ------- |
| Wangpan     | [@aspirer](https://github.com/aspirer)       | NetEase |
| XiaoCui Li  | [@ilixiaocui](https://github.com/ilixiaocui) | NetEase |
| HongSong Wu | [@wuhongsong](https://github.com/wuhongsong) | NetEase |
| JingLi Chen | [@Wine93](https://github.com/Wine93)         | NetEase |
| Yi Cheng    | [@Cyber-SiKu](https://github.com/Cyber-SiKu) | NetEase |
| HanQing Wu  | [@wu-hanqing](https://github.com/wu-hanqing) | NetEase |

We hope to see more diversity in the maintainer team.

## Adoption

Belows are the NEW companies this year that we know have deployed Curve in their own
infrastructure or integrated Curve as part of their products.

*   Zstack
*   SmartMore 
*   Wanfang SCI.&Tech
*   ZheJiang Yixin Tech
*   Jiangsu Rural Credit Union


## Project perform

This year, we are continuing to add new features to Curve. Some key features added to the project:

*   A brand new Document website, a forum community and a Wechat Official Account for users.
*   Support RDMA+SPDK for CurveBS, provides ~50% lower IO latency than TCP+Ext4 engine
*   Cloud-native database(shared-storage) support base on PolarFS and CurveBS
*   Support storing snapshot in local CurveBS cluster, no longer rely heavily on S3 storage
*   Add hybrid storage engine with SSD/NVMe and HDD for CurveBS
*   Support CurveFS which is POSIX-compatible and mountable by FUSE,  and can store data in Object Storage System by S3-like interface, or store in CurveBS
*   Add memory buffer, local disk, distributed k-v data caches, as well as metadata memory caches for CurveFS
*   Developed a new deployment tool CurveAdm, and refactored a new command line tool for CurveBS and CurveFS
*   Support deploy CurveBS and CurveFS on kubernetes, and also CSI plugin for them

## Project Goals

The goal of Curve this year is to increase the diversity in both contributors and adoption. Some key works include:

*   More timely update about Curve new features based on our roadmap.
*   Improve user documentation, publish more user cases, and provide more help for community user testing.
*   Investigate the needs of community users and design and develop related needs.

## Incubation readiness

Not submitted yet, we still want to further improve the Curve project, and consider submitting after getting more community user cases and more diversity of maintainers.



