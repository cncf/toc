# KubeDL 2022 Annual Review

## Table of Contents

## Background
KubeDL is a suite of Kubernetes controllers that enable running machine learning workloads on Kubernetes, such as model training and inferences, configuration tuning for model deployments.
Project website https://kubedl.io/

## DevStats
CNCF [devstats](https://kubedl.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m) website
Docker Pulls: 6.7K
Stars: 344

## Contributors and Maintainers

We currently have 18 contributors and 5 maintainers.

- Qiukai Chen (@SimonCqk ), Maintainer, Alibaba
- Yinghao Yu (@yhalpha ), Maintainer, Alibaba
- Luping Wang (@lwangbm), Maintainer, Alibaba
- Jian He(@jian-he), Maintainer, Alibaba
- Gaoyuan He@(HeGaoYuan), Contributor, Tencent
- Yuzhong Wang (@BourbonWang), Contributor, Chongqing University
- Lingyung Yang (@xxx) Hongkong University of Science and Technology

## Adoption
KubeDL has been used at Alibaba to run machine learning training workloads at very large scale, and has also been used by many external companies. We cannot obtain official adoption statements at this point.
KubeDL has achieved great progress in adopting production usages. Many of the users use the project but have not contributed in community as much. We will aim to engage more community contributors next year.
KubeDL has also a number of contributors from research institutes and universities to innovate on forward-thinking features in ML infra to improve the scheduling and performance efficiency.

## Project Past
KubeDL has added many significant features since last year.
- Accelerate training speed by natively integrating data cache.
- Enhanced KubeDL dashboard for better interactive experiences.
- Miscellaneous bug fixes and improvements.
-
KubeDL has also been under several summer of code projects, such as ASOC(link), xxx, where a number of students are actively participating.

KubeDL/Morphling project is honored to have a paper published at ACM SoCC 2021.

## Current Goals
Project Features

- Support elastic training functionality in KubeDL.
- Enhance fault tolerance of KubeDL training.
- Version control for experiments in Morphling.
- Enable cache sharing across training jobs.

Docs

- Improve the docs more in "How it works", in addition to "How to use".

Community

We focused a lot in enhancing the project quality to production grade, and continuously contribute our best practices in large production to the community. We have many users using KubeDL in production, but have not get many external contributors in community. We had short amount of resources. Next year, we will focus more on building more awareness of the project by writing more blogs, showcasing, and presentation in conferences, and hope to engage more users to contribute.

## Incubation
KubeDL is not yet ready for incubation, though KubeDL has been used heavily at Alibaba in large scale. We would like to see more contributors and maintainers and more adopters.