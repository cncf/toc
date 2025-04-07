**TiKV Graduation Proposal**

TiKV is an open source distributed transactional key-value database. TiKV joined CNCF as a Sandbox project in [August 2018](https://www.cncf.io/blog/2018/08/28/cncf-to-host-tikv-in-the-sandbox/), and was voted by CNCF TOC as an incubation project in [May 2019](https://www.cncf.io/blog/2019/05/21/toc-votes-to-move-tikv-into-cncf-incubator/). The TiKV project has grown significantly over time since joining CNCF, in terms both fostering a healthy ecosystem of maintainers and contributors and production adoptions.

To highlight some of the achievements:

- GitHub Stars: 7400+
- Contributors in the TiKV core repository: 239
- Forks: 1100+
- Releases: 86 releases
- Adopters: 1000+ (including commercial users and community users)
- 4.0 GA released in May, 2020

On behalf of the maintainers team, we believe TiKV is ready for [graduation stage](https://github.com/cncf/toc/blob/master/process/graduation_criteria.md#graduation-stage).

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

- **More diversified (company-wise) core maintainers**:

    - In the incubating proposal, we had "approximately 82% of the commits came from PingCAP, with the remaining 18% coming from companies like Alibaba, Docker, Netease, Tencent, and others". Now, the contribution ratio has been improved - with PingCAP contributing 70% and the remaining 30% coming from multiple other companies like Bilibili, Netease, Spotify, Grafana, and others.

    - With regard to the maintainer structure, we have two new maintainers promoted from our adopters. With more diversity introduced to the maintainership we believe the TiKV project is heading towards healthy and sustained development.


- **More diversified user base**:

    - We have adopters distributed in a broader geographical range, including Europe (dailymotion), East and Southeast Asia (VNG, Bookmyshow, Shoppee, etc), Japan, and the US (not-yet-public adopters can be shared upon request).
    - We have more known adopters deploying TiKV independently without TiDB. Some of the use cases are:
        - JD Cloud (full-service cloud computing platform)  adopts TiKV for its metadata storage of their Object Storage Service (OSS), which was migrated from MySQL. See [Case study: TiKV in JD Cloud](https://www.cncf.io/blog/2019/11/26/case-study-tikv-in-jd-cloud/)
       
        - Yidian Zixun (news aggregator) uses TiKV independently to build performant distributed storage for their user profiling data, by implementing a tikv-proxy that allows TiKV deployment either via tidb-ansible or a modified tidb operator (which they will migrate to)
        
        - Zhihu (quora in China) is building Zetta Table Store, its structured data layer on top of TiKV, to serve its trillion rows of data of feed service. They also hold the record of the largest single cluster, with 252 TiKV nodes. The use case study is working in progress.
    
    - Current adopters are across a wide range of industries including finance, internet, online education, news, e-commerce, entertainment, etc.

- **Roadmap with release target**:

  We have roadmap per release, but previously only on TiDB. Now already published on the [TiKV website](https://tikv.org/docs/dev/roadmap/).

- **Performance test results and regression test between releases**

  We have the following tests in-between releases. A test report will be provided before each release. The related doc can be requested if needed.

    - CI tests
      - Unit tests
      - Integration tests
      - Compatibility testing
    - Feature tests
      - Bugfix tests
      - Ansible tests
      - Regression tests
      - About 200+ cases
    - Chaos test (on both the internal chaos engineering platform and Chaos Mesh, the open-sourced Chaos Engineering solution for Kubernetes )
      - 4 functionality tests
      - More than 5 stability tests
      - 2 compatibility test
    - Jepsen tests
      - Performance tests
      - Sysbench five OLTP tests
      - TPC-C
      - TPC-H, covers 22 SQLs
      - YCSB (workloada, workloadb, workloadc, workloadd, workloadf)

- **A TiKV centric toolset and reduce hard external binary dependencies**:

 TiKV can be easily deployed on Kubernetes via TiDB Operator and monitored by Prometheus, as already has been practiced by some of our adopters like Yidian Zixun.  
 [Updated on June 8] The community has already provided the [design](https://github.com/tikv/rfcs/pull/53) and the initial implementation of [TiKV operator](https://github.com/tikv/tikv-operator), which is a Kubernetes Operator that aims to automate the deployment, management, and operational tasks of TiKV clusters.
 
 

