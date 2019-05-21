# TiKV Incubating Stage Review Proposal

TiKV is currently a sandbox stage member project. Please refer to TiKV's [sandbox stage proposal](https://github.com/cncf/toc/pull/133)
for details on the previous proposal.

Since entering the CNCF Sandbox in August 2018, TiKV has demonstrated growth in both adoption and community participation. 18 releases were published with 468 commits from 83 individuals. Approximately 82%
of the commits came from PingCAP, with the remaining 18% coming from companies like Alibaba, Docker, Netease, Tencent, and others. These stats are collected from the [Company Summary](https://tikv.devstats.cncf.io/d/5/companies-summary?orgId=1&var-period_name=Last%20year&var-metric=contributions) of TiKV devstats dashboard.

Across these releases, several important features were added to TiKV including:
 - Increase scalability with distributed garbage collection
 - Raft prevote for improved network stability
 - More efficient batch gRPC APIs
 - Multithreaded Raftstore to reduce Raft communications bottleneck
 - HTTP-based monitoring endpoints

TiKV has experienced global adoption growth, including growth in Southeast Asia (Shopee), India (BookMyShow.com), and Japan. (See current list of all known adopters below.) The project has also fostered an ecosystem of other technology built on top of TiKV, including three Redis-on-TiKV projects: [Tidis](https://github.com/yongman/tidis), [Titan](https://github.com/meitu/titan), and [Titea](https://github.com/gengmei-tech/titea), realizing one of TiKV's core visions to become a building block that [simplifies building other systems](https://docs.google.com/presentation/d/1864TEfbwCpbW5kPYGQNAfqAUdc3X83n-_OYigqxfohw/edit#slide=id.g3cb7ac7914_0_28).
 
The TiKV community has [formalized a governance policy](https://github.com/tikv/tikv/blob/master/GOVERNANCE.md), executed its governance process to bring on an additional maintainer, [@sunxiaoguang of Zhihu.com](https://github.com/tikv/tikv/pull/4397) (China's Quora), and
achieved [passing level](https://bestpractices.coreinfrastructure.org/en/projects/2574) for CII best practices.   

## Supporting Evidence for Incubating Stage

 - Current list of [publicly known TiKV adopters](https://github.com/tikv/tikv/blob/master/docs/adopters.md) including Mobike, Xiaomi, Bank of Beijing, Shopee (SEA Group), BookMyShow.com, VIPKID. (Additional not-yet-public adopters can be shared upon request.)
 - Current list of [project maintainers](https://github.com/tikv/tikv/blob/master/MAINTAINERS.md)
 - Project maintainers are added and removed per [TiKV Governance Policy](https://github.com/tikv/tikv/blob/master/MAINTAINERS.mdhttps://github.com/tikv/tikv)
 - Releases: https://github.com/tikv/tikv/releases
 - Current [Roadmap](https://github.com/tikv/tikv/blob/master/docs/ROADMAP.md)
 - Contributors: https://github.com/tikv/tikv/graphs/contributors
 - Commit activity: https://github.com/tikv/tikv/graphs/commit-activity
 - CNCF [devstats dashboard](https://tikv.devstats.cncf.io/d/8/dashboards?refresh=15m&orgId=1) and [Community Stats](https://tikv.devstats.cncf.io/d/3/community-stats?orgId=1)
