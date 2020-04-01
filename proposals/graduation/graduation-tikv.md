**TiKV Graduation Proposal**

TiKV is an open source distributed transactional key-value database. TiKV joined CNCF as a Sandbox project in [August 2018](https://www.cncf.io/blog/2018/08/28/cncf-to-host-tikv-in-the-sandbox/), and was voted by CNCF TOC as an incubation project in [May 2019](https://www.cncf.io/blog/2019/05/21/toc-votes-to-move-tikv-into-cncf-incubator/). The TiKV project has grown signifantly over time since joining CNCF, in terms both fostering a healthy ecosystem of maintainers and contributors and production adoptions.

To highlight some of the achievements:

GitHub Stars: 7100+
Contributors in the TiKV core repository: 226
Forks: 1100+
Releases: 86 releases
Adoptors: 500+

On behalf of the maintainers team, we believe TiKV is ready for [graduation stage](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#graduation-stage).

## Graduation State Criteria

### Have committers from at least two organizations.

The [maintainers team](https://github.com/tikv/tikv/blob/master/MAINTAINERS.md) currently consists of 7 members, with a healthy distribution of corporations represented (PingCAP, Zhihu, JD Cloud, and Yidian Zixun).

Committers by company for the [preceding 12 months](https://tikv.devstats.cncf.io/d/4/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=activity&var-repogroup_name=All&var-companies=All) lists the top ten organizations as:

- PingCAP
- Independent
- NetEase
- Alibaba
- CNCF
- Tencent
- Earth
- Bilibili Group
- Meitu
- Tradeshift


### Have achieved and maintained a [Core Infrastructure Initiative Best Practices Badge/(https://bestpractices.coreinfrastructure.org/).

We have achieved passing level criteria for [CII Best Practices](https://bestpractices.coreinfrastructure.org/zh-CN/projects/2574).

### Have completed an independent and third party security audit with results published

A third-party (Cure53) security audit funded by CNCF was performed in the span of February and March of 2020. The full report has been published on the [TiKV website](https://tikv.org/blog/tikv-pass-security-audit/) and archived in the [TiKV repository](https://github.com/tikv/tikv#security-audit).

### Explicitly define a project governance and committer process.

The TiKV project governance is defined in the [community repository](https://github.com/tikv/community/blob/master/GOVERNANCE.md).

### Have a public list of project adopters for at least the primary repo

https://github.com/tikv/tikv/blob/master/docs/adopters.md

## Incubation Details

### * Link to Incubation Due Diligence(DD) Document

[TiKV Due Diligence for CNCF Incubation](https://docs.google.com/document/d/1isLh8e5D0pAgocbj2ht_ZbZ5MORN3fFD-98tCsVu0rM/edit#heading=h.xty7ca1yvruj)

### * Address any concerns or recommendations from the SIG and/or TOC sponsor(s) from the DD Document

- More diversified (company-wise) core maintainers:

    - In the incubating proposal, we had "approximately 82%
of the commits came from PingCAP, with the remaining 18% coming from companies like Alibaba, Docker, Netease, Tencent, and others". Now, the contribution ratio has been improved - with PingCAP contributing 70% and the remaining 30% coming from multiple other companies like Bilibili, Netease, Spotify, Grafana, and others.
    - With regard to the maintainer structure, we have two new maintainers promoted from our adopters. With more diversity introduced to the maintainership we believe the TiKV project is heading towards healthy and sustained development.

- More diversified user base:
    - We have adopters distributed in a more broader geographical range, including Europe (dailymotion), East and Southeast Asia (VNG, Bookmyshow, Shoppee, etc), Japan, and the US (not-yet-public adopters can be shared upon request).
    - We have more (8) known adoptors deploying TiKV independently without TiDB. Some of the case studies are:
       - [Case study: TiKV in JD Cloud](https://www.cncf.io/blog/2019/11/26/case-study-tikv-in-jd-cloud/)
       - [Zhihu's 1.3 Trillion Rows of Data, Milliseconds of Response Time](https://dzone.com/articles/lesson-learned-from-queries-over-13-trillion-rows-1)
    - Current adopters are across a wide range of industries including finance, internet, online education, news, e-commerce, entertainment, etc.