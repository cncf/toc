# Fluentd Graduation Application

[Fluentd](https://www.fluentd.org) project has been in incubation for more than <u>two years</u>. The following application links to the required information to become a graduated project plus some additional information about it growth.

### Fluentd Adoption and Ecosystem

In terms of adoption, It's hard to measure the real numbers of deployments of Fluentd, but based in the public statistics considering <u>only</u> our own Docker images, in the last years we see more than <u>48 million pulls</u>.

As a reference of adoption by known companies, we can see that Fluentd plays an important role in the following production environments:

- [Google Cloud Platform Logging Agent](https://cloud.google.com/logging/docs/agent/): Fluentd is packaged and called [google-fluentd](https://github.com/GoogleCloudPlatform/google-fluentd). The official agent for GCP which comes with Stackdriver integration right out of the box.
- [Microsoft Azure Log Analytics](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-data-sources-json): Fluentd is the default agent for OMS to collect metrics and logs from Azure systems, the packaged solution is called [OMS Agent](https://github.com/Microsoft/OMS-Agent-for-Linux).
- [Redhat / OpenShift and Aggregated Logging](https://docs.openshift.com/container-platform/3.10/install_config/aggregate_logging.html): Fluentd is the default solution for aggregated logging through the combination of other components like Elasticsearch and Kibana (known as E__F__K stack)

### Understanding Fluentd and it Ecosystem

Fluentd have a different ecosystem model than other projects, it's composed by a small and stable code base and most of the functionalities are available through plugins which can be: inputs, filters, parsers, buffers and outputs.

Most of contributions to Fluentd happens through the implementation of plugins, where the majority of them (**> 900**) are maintained by different companies and individuals. Since the nature of Fluentd language in Ruby and C, these extensions are available through the common Ruby Gems channels and Fluentd offers a built-in tool to install any plugin available.

note: during the last year we have seen a 50% decline of the number of companies contributing to the code base (Fluentd core), while at the same time a huge increase in the number of contributions to it ecosystem through plugins.

## Progress since joining CNCF as an incubation project on December 2016

- Fluentd **v1.0** released on December 2017 (announced at KubeCon North America)

- Fluentd Code Base

  - **75** Official releases  (today we are at **v1.4.1**)

- **New** Fluentd plugins hosted in our Github organization
  - Prometheus Output: https://github.com/fluent/fluent-plugin-prometheus
  - NodeJS Output: https://github.com/fluent/fluent-logger-node/
  - Kafka Output: https://github.com/fluent/fluent-plugin-kafka
  - Splunk Output: https://github.com/fluent/fluent-plugin-splunk

- Fluentd Plugins published <u>by the community</u>

  - Total as of today: **941**
  - Plugins <u>updated since</u> November 2016 (CNCF Incubation):  **917**

- Community

  - Slack: on January 2017 an average of 11 new members joined, as of July 2018 we have an average of 23 new members per week.

  - Github Stars: **7500** as of today


## To graduate from inception or incubating status, or for a new project to join as a graduated project, a project must meet the incubating stage criteria plus:

### Have committers from at least two organizations.

- [Naotoshi Seo](https://github.com/sonots), [DeNA](https://dena.com/intl/)
- [Okkez](https://github.com/okkez), [Clearcode](https://www.clear-code.com/)
- [Hiroshi Hatake](https://github.com/cosmo0920), [Clearcode](https://www.clear-code.com/)
- [Masahiro Nakagawa](https://github.com/repeatedly),  [Treasure Data](https://www.treasuredata.com/)
- [Satoshi Tagomori](https://github.com/tagomoris), [Treasure Data](https://www.treasuredata.com/)

### Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

https://bestpractices.coreinfrastructure.org/projects/1189

### Adopt the CNCF Code of Conduct.

https://github.com/fluent/fluentd/blob/master/code-of-conduct.md

### Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

https://github.com/fluent/fluentd/blob/master/GOVERNANCE.md

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

https://github.com/fluent/fluentd/blob/master/ADOPTERS.md
