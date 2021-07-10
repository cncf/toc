# NATS Graduation Application

NATS was accepted into the CNCF as an incubating project in March 2018. Significant progress has been made in 2018 to improve process and further mature NATS as a project under the guidance of the CNCF.

## The following represents the successful CNCF graduation criterion met by NATS.

### Committers
**Have committers from at least two organizations.**

NATS has committers from [various organizations](https://nats.devstats.cncf.io/d/5/companies-summary?orgId=1&var-period_name=Last%20decade&var-metric=repositories).

* Christoper Watford [@watfordgnf](https://github.com/watfordgnf), GE
* Oleg Shaldybin [@olegshaldibin](https://github.com/olegshaldibin), Google
* Laurent Magnin [@laugimethods](https://github.com/laugimethods), Logimethods
* Paulo Pires [@pires](https://github.com/pires), Mesosphere
* Brian Shannan [@brianshannan-wf](https://github.com/brianshannan-wf), Workiva
* Charlie Strawn [@charliestrawn](https://github.com/charliestrawn), Workiva
* Lev Brouk [@levb](https://github.com/levb), Independent
* R.I. Pienaar [@ripienaar](https://github.com/ripienaar), Independent
* Jaime Piña [@variadico](https://github.com/variadico), Independent
* Various committers from Synadia
  - Derek Collison [@derekcollison](https://github.com/derekcollison)
  - Ivan Kozlovic [@kozlovic](https://github.com/kozlovic)
  - Waldemar Quevedo [@wallyqs](https://github.com/wallyqs)
  - Colin Sullivan [@ColinSullivan1](https://github.com/ColinSullivan1)
  - Alberto Ricart [@aricart](https://github.com/aricart)
  - Stephen Asbury [@sasbury](https://github.com/sasbury)
  
### Best Practices Badge
**Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.**

* https://bestpractices.coreinfrastructure.org/en/projects/1895

### Have completed an independent and third party security audit.

Third party security audit performed by Cure53 at the behest of the CNCF in November 2018. (https://github.com/nats-io/nats-general/blob/master/reports/Cure53_NATS_Audit.pdf)

### Adopt the CNCF Code of Conduct

* https://github.com/nats-io/nats-general/blob/master/CODE-OF-CONDUCT.md

### Governance
**Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.**

* https://github.com/nats-io/nats-general/blob/master/GOVERNANCE.md

## Addendum
This section serves as additional information with regard to the NATS project graduation.

Some concern has been expressed about the long term viability of NATS should the maintainers at Synadia reduce or stop involvement in the project.  We are confident this will not happen, but we’d like to address this concern.
### Maturity

First, NATS is mature, and may be the most mature and stable project in the CNCF, turning 10 years old this year.  NATS was created at VMWare and has been maintained by employees across a number of companies - VMWare, Pivotal, Apcera, and now Synadia.  Over these ten years, NATS has been successfully deployed in production for years by companies such as Pivotal, GE Nuclear, Ericsson, VMWare, and Tinder.
 NATS has been run for years in mission-critical production environments speaking to its stability and market perception of long term viability.  One production use case demonstrating the stability of NATS in the energy sector, is where NATS can only be upgraded when spent fuel rods are being pulled - one of the software engineers behind that use case has become a NATS maintainer.  Furthermore, NATS has been through a thorough security audit and is used in highly secure and hardened environments such as those in the financial and banking industries. 
 
In addition, in these ten years the client protocol has never broken compatibility meaning that clients built ten years ago can connect to the latest version of the NATS server. We strive to maintain a mature, stable project and our users agree.  On the [CNCF project page](https://www.cncf.io/projects/), there is a reference to “Crossing the Chasm” correlating project status and consumer adoption.  With the adoption of NATS by companies like GE and Mastercard, we pose that NATS is being adopted by the “Early Majority” and arguably “Late Majority” consumers, establishing NATS as mature enough to graduate.

### Long Term Viability 

With NATS maturity, the project has demonstrated a solid market fit, with no significant feature gaps in the context of our stated tenets, nor has NATS had any large architectural changes over the last several years until now with NATS 2.0.  These changes are transparent to existing deployments and NATS is still protocol compliant with clients developed nine years ago.

Some of the NATS project core maintainers have supported NATS over the course of their careers through multiple companies, including VMWare, Apcera, Google and Synadia, speaking to a strong belief, dedication and loyalty to the NATS project.  Even in the unlikely event that Synadia NATS maintainers leave, the stability of the project gives plenty of time for new maintainers from companies like Google, IBM or GE to ramp up.
### Maintainers and Contributors

In terms of contributions, the NATS.io project is unique compared to other CNCF projects in that there are 40 actively maintained repositories, with commits from [28 contributing companies](https://nats.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&from=now-5y&to=now) over the last 5 years.  We meet the written requirements of the graduation process, “Have committers from at least two organizations.”
We have a wide variety of maintainers and contributors to the NATS client libraries.  Admittedly, most of the core server code changes are submitted by the Synadia maintainers who have had extensive experience with messaging, some for over 25yrs - diving in to add a feature in the NATS server is not a weekend project for most.  However, we do have maintainers who can pick up and delve into the servers such as R.I. Pienaar or Oleg Shaldybin of Google.  In the past, Oleg has fixed some non-trivial bugs demonstrated [here](https://github.com/nats-io/gnatsd/commits?author=olegshaldibin) and also has experience as a NATS user.

The TOC has raised concerns of the diversity of committers in two repositories for the NATS servers.  By design, NATS clients are relatively easy to develop due to a simple, well documented, and easily parsed protocol.  This has resulted in a plethora of contributions on the client-side with over 50 types of clients.  The NATS servers are different, requiring a deep dive into code that usually requires quite an investment of time - much more than a weekend project.  We’ve had some really good contributions, and ask for them, but more often users often look to the maintainers to create PRs in resolving open issues. For large features, we have an open discussion in Slack and publicly shared design documents opened for community discussion.  We have found this serves the NATS community very well as the core NATS maintainers can add features quickly and with efficacy while the community drives requirements.  This creates a solid and stable project that addresses the needs of our users.  The current process has worked well to date and certainly hasn’t inhibited adoption.

Given NATS maturity, stability, adoption, and proven resilience to survive over three companies we are confident that NATS has a long future ahead of it.  Of course, as we continue to grow the ecosystem, we will welcome additional maintainers.

### Code and Documentation
While not simple, the NATS server code is well commented and well organized as to the rationale of design decisions.  One of many examples of this can be found here [here](https://github.com/nats-io/nats-streaming-server/blob/v0.14.0/server/server.go#L110).  The NATS server protocol is well documented [here](https://docs.nats.io/nats-protocol/nats-protocol) and NATS streaming protocol [here](https://docs.nats.io/developing-with-nats-streaming/protocol).  Clear, concise code and comments allow for quick ramp up time for additional maintainers and for those who have the time to contribute.

### Summary

To summarize, with NATS maturity, stability, and track record of loyal maintainers, we meet the criteria of having maintainers from multiple companies.  With our thorough testing and CI processes, production issues have been almost non-existent.  While the question of future viability of NATS has been asked, we personally have no concerns given the production usage, maturity, and stability, and we feel the current membership of the NATS maintainer team does not pose a risk for the CNCF hosting NATS as a graduated project. 

### Due Diligence

* [NATS Incubation Proposal](https://github.com/cncf/toc/pull/168)
* GitHub Repo Renaming Note wrt DevStats - Some key repositories for the NATS project were renamed when we updated our naming convention in June 2019, e.g. `gnatsd` -> `nats-server` and our client libraries were renamed as nats.`client language abbreviation` e.g. `nats.go`. Because of this rename, the DevStats dashboards for some repositories may stop showing values on the original repo name and continue on the new name.
* Q: What is the general quality of informal communication around the project (slack, github issues, PR reviews, technical blog posts, etc)?
    * NATS Blog (https://nats.io/blog)
    * DevStats - [Open PR Age by Repo](https://nats.devstats.cncf.io/d/19/open-pr-age-by-repository-group?orgId=1)
    * NATS Slack #General channel metrics over the last year
        * ![Slack Metrics](https://github.com/nats-io/nats-general/blob/master/Slack_metrics_2019.png)

