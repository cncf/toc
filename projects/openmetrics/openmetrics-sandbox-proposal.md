== OpenMetrics

*Name of project:* OpenMetrics

*Description:*

OpenMetrics refines the Prometheus exposition format into an independent standard.

Prometheus has become the de facto standard in cloud-native metric monitoring, and has active upstream work by competitors.
The ease of implementing this exposition data has led to an explosion in compatible metrics endpoints with 300+ exporters registered, dozens of native integrations, and unknown numbers of internal adoptions.

To allow for even more adoption, OpenMetrics received a lot of additional scrutiny and engineering time from several large players in the cloud-native space.
It also puts the format under a neutral name, allowing more monitoring vendors to adopt it without potential political considerations.

With substantial commitments for adoption, OpenMetrics will enjoy solid support from day 1.
Amongst others, these are:

* Prometheus
* Cloudflare
* GitLab
* Google
* Grafana
* InfluxData
* Oath.com
* RobustPerception
* SpaceNet
* Uber

OpenMetrics was presented at the [CNCF TOC meeting on 2018-06-19](https://docs.google.com/presentation/d/1Ym8fLRCaX43uHPHBRyuRXM62U8m4vXaBXkuUp6tt3js/edit#slide=id.g25ca91f87f_0_0).

*Statement on alignment with CNCF mission:*

Given the CNCF's stated role in "fostering the growth and evolution of the ecosystem" and "making the technology accessible and reliable", we believe OpenMetrics helps with both of these goals.

*Sponsor / Advisor from TOC:* Alexis Richardson, Bryan Cantrill


*Unique identifier:* openmetrics

*Preferred maturity level:* sandbox

*License:* Apache License v2.0

*Source control repositories:* https://github.com/RichiH/OpenMetrics/

*External Dependencies:*

OpenMetrics currently depends on no external software components.

Once the test suite is released, it will depend on Go and Python and some libraries. Proper license hygiene will be ensured.

*Lead:* * Richard Hartmann (SpaceNet)

*Infrastructure requests (CI / CNCF Cluster):* None

*Communication Channels:*

*Issue tracker:* https://github.com/RichiH/OpenMetrics/issues

*Website:* https://www.openmetrics.io

*Release methodology and mechanics:*

Given that this is a format, releases will be slow, deliberate, and forward- and backwards-compatible.

*Social media accounts:* None

*Existing sponsorship*: None

*Community size:*

* 128 stars
* 15 forks
* Commitments by companies with billions of combined yearly turnover
* 6 people on bi-weekly call

*Production usage*: None yet
