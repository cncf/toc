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

### Adopt the CNCF Code of Conduct

* https://github.com/nats-io/gnatsd/blob/master/CODE-OF-CONDUCT.md

### Governance
**Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.**

* https://github.com/nats-io/nats-general/blob/master/GOVERNANCE.md

## Addendum
This section serves as additional information with regard to the NATS project graduation.

Some concern has been expressed about the long term viability of NATS should the maintainers at Synadia reduce or stop involvement in the project.  We are confident this will not happen, but we’d like to address this concern.
### Maturity

First, NATS is mature, and may be the most mature project in the CNCF, turning nine years old this year.  NATS has been run for years in mission-critical production environments speaking to its stability and market perception of long term viability.  One production use case demonstrating the stability of NATS in the energy sector, is where NATS can only be upgraded when spent fuel rods are being pulled - one of the software engineers behind that use case has become a NATS maintainer.  Furthermore, NATS has been through a thorough security audit and is used in highly secure and hardened environments such as those in the financial and banking industries. 

### Long Term Viability 

With NATS maturity, the project has demonstrated a solid market fit, with no significant feature gaps in the context of our stated tenets, nor has NATS had any large architectural changes over the last several years until now with NATS 2.0.  These changes are transparent to existing deployments and NATS is still protocol compliant with clients developed nine years ago.

Some of the NATS project core maintainers have supported NATS over the course of their careers through multiple companies, including VMWare, Apcera, Google and Synadia, speaking to a strong belief, dedication and loyalty to the NATS project.  Even in the unlikely event that Synadia NATS maintainers leave, the stability of the project gives plenty of time for new maintainers from companies like Google, IBM or GE to ramp up.
### Maintainers and Contributors

We have a wide variety of maintainers and contributors to the NATS client libraries.  Admittedly, most of the core server code changes are submitted by the Synadia maintainers who have had extensive experience with messaging, some for over 25yrs - diving in to add a feature in the NATS server is not a weekend project for most.  However, we do have maintainers who can pick up and delve into the servers such as R.I. Pienaar or Oleg Shaldybin of Google.  In the past, Oleg has fixed some non-trivial bugs demonstrated [here](https://github.com/nats-io/gnatsd/commits?author=olegshaldibin) and also has experience as a NATS user.

If you look overall at the number of recent companies and contributors we have [here](https://nats.devstats.cncf.io/d/7/companies-contributing-in-repository-groups?orgId=1&from=now-6M&to=now&var-period=d28&var-repogroup_name=All), you’ll see there is a variety of companies contributing to NATS.

### Code and Documentation
While not simple, the NATS server code is well commented and well organized as to the rationale of design decisions.  One of many examples of this can be found here [here](https://github.com/nats-io/nats-streaming-server/blob/v0.14.0/server/server.go#L110).  The NATS server protocol is well documented [here](https://nats.io/documentation/internals/nats-server-protocol/) and NATS streaming protocol [here](https://nats.io/documentation/streaming/nats-streaming-protocol/).  Clear, concise code and comments allow for quick ramp up time for additional maintainers and for those who have the time to contribute.

### Summary

To summarize, with NATS maturity, stability, and track record of loyal maintainers, we meet the criteria of having maintainers from multiple companies.  With our thorough testing and CI processes, production issues have been almost non-existent.  While the question of future viability of NATS has been asked, we personally have no concerns given the production usage, maturity, and stability, and we feel the current membership of the NATS maintainer team does not pose a risk for the CNCF hosting NATS as a graduated project. 
