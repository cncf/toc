# Falco Incubation Stage Review Proposal

The goal of the Falco project is to provide cloud native runtime security with a focus on Kubernetes intrusion & anomaly detection. Falco achieves this through the use of eBPF, Linux kernel probes, and the Kubernetes audit logging API to create an event stream of low level system and orchestrator events. This includes network connection info, file opens/writes, process information, and Kubernetes API calls. Falco also polls metadata (labels, resource names, container info, etc) from the Kubernetes API server and the underlying container runtime. 

The Falco engine applies rule sets (or policies) to this event stream to identify potentially malicious behavior, or to simply log activity on a Kubernetes cluster for later auditing or review. The Kubernetes metadata can be incorporated into the rule conditions to allow rules to apply to only particular resources in a Kubernetes cluster. Falco alerts can be sent to a wide range of destinations including logging aggregation services, pub/sub services, serverless functions, and more. This allows Falco users to integrate Falco alerts into their existing incident response and Ops/SRE workflows.

Falco entered the CNCF Sandbox on October 11th, 2018. Please see the project's [Sandbox proposal](https://github.com/cncf/toc/blob/master/proposals/falco.adoc). 

## Project Progress since Sandbox

In short, the Falco project has seen terrific growth and project participation since entering the CNCF Sandbox. 

* All key metrics around contributions and committers have seen significant growth.
* Downloads have seen significant growth.
* The project has seen an increase in integrations, as well as Falco being incorporated into other products (Sumo Logic, Altran).
* The community has seen increased participation and activity.
* Public end users are documented in the projects' [ADOPTERS.md](https://github.com/falcosecurity/falco/blob/dev/ADOPTERS.md). Two end users are speaking at Kubecon NA 2019.
* The project completed a security audit and fixed 13 of 14 security bugs. The audit will be published when the last bug is fixed within the next two releases (~1 month).
* The project successfully participated in the Google Summer of Code.
* The project successfully defined and shipped a roadmap. 

### Metrics

#### Github

The CNCF provides Falco Github metrics through its [devstats site](https://falco.devstats.cncf.io/d/65/summary-dashboard?orgId=1&from=now-3y&to=now-1w&var-period=m&var-repo_name=Falco&var-repogroup_name=All&var-repogroups=All). The devstats metrics show that the Falco team has maintained a constant flow of commits, and has increased the velocity of commits. This is due in part to Sysdig's sponsorship of the project, providing 4 full time employees for the project. Overall committers increased by 40 to 56, showing that while the sponsorship does influence commit velocity, the project is taking commits from a wider group outside of Sysdig. This is also evidenced by the Contributing Companies metric increasing from a max of 7 to a max of 13.

| GitHub |Pre-Sandbox* (< 10-2018) | Post-Sandbox (>= 10-2018) | | 
|--------|-------------------------|---------------------------|-|
| Commits| 3051 (24 per week)| 2642 (49 per week) |+104% increase |
| Committers | 16 | 58 (+262%) | |
| Contributors | Avg: 7 Max: 25 | Avg: 27 (+285%) Max: 55 (+120%) |28 Day MA |
| Contributing Companies | Avg: 2 Max: 7 | Avg: 7 (+250%) Max: 15 (+114%) | 28 Day MA |
| Stars | 945 | 1835 (+94%) | |
| CII Progress| 0% | 70% to “passing” | |

**Note:** Post Sandbox #'s as of 2019-10-13

\* Pre-Sandbox is a 29 month period (May 2016-Oct 2018)

#### Downloads

Downloads and Docker Hub pulls increased by 179% and 247% respectively. 

| Eco System | Pre-Sandbox* (< 10-2018) | Post-Sandbox (>= 10-2018) ||
|------------|-------|----------|---------|
| Docker Hub Pulls | 1,918,000 | 6,762,966 (+252%) | Current Velocity: ~ 400K/Month |
| Downloads | 110,000 | 270,000 (+145%) | dpkg/rpm |
| Integrations | 9 | 40 (+344%) | 8 CNCF Projects |

**Note:** Post Sandbox #'s as of 2019-10-13

\* Pre-Sandbox is a 29 month period (May 2016-Oct 2018)

### Integrations

Integrations have been a major focus of the Falco team over the last year. Falco's usefulness is magnified when its data can be easily sent to an end user's existing Ops and SRE tooling. The Falco project can count over 40 integrations in a number of categories.

- **Orchestrators:** Kubernetes, Mesosphere
- **Container runtime:** Docker, Containerd, CRI-O, rkt, lxc
- **Logging:** Splunk, Datadog, Elasticsearch, Sumo Logic, Stackdriver Logging, Fluentd, Fluentbit, Kibana, Grafana Loki
- **Pub/Sub:** NATS, AWS SNS, Google Pub/Sub
- **Serverless:** AWS Lambda, Google Functions, Kubeless
- **Other:** gRPC, Prometheus, Helm, Slack, Microsoft Teams

Integrations have also been an easy way for Falco to encourage external contributions and maintainership. Falco community member [Thomas Labarussias](https://github.com/Issif) has contributed and maintains the [Falco Sidekick](https://github.com/falcosecurity/falcosidekick) subproject that includes integrations to 12 services.

One of the other ways Falco is being integrated is by other CN ecosystem members including Falco in their product or service. Significant examples of these integrations include:

* **Sumo Logic:** Sumo Logic integrates Falco as part of its [Continuous Intelligence Platform for Kubernetes](https://www.sumologic.com/solutions/kubernetes/). Sumo Logic also integrates other CNCF projects such as Prometheus, and Fluentd; as well as the OSS project Fluentbit.
* **Atran:** Altran has integrated Falco into its Secure Cloud Native Fabric product along with other OSS projects such as Istio, OSSEC, Kube-bench, Envoy, Cloud Custodian, Clair, Anchore, and more. More information is available via Altran's [whitepaper (pdf)](https://setns.run/altranwp). 

**Note:** Not all integrations are created equal. Some of the above integrations are core changes the Falco project had to make to the Falco code base itself. Examples include Containerd support, CRI-O support, gRPC, and Prometheus. Other integrations are configuring existing tooling to consume Falco alerts. These integrations are either published as documentation, blog posts, or sub projects (such as the Falco [Kubernetes Response Engine](https://github.com/falcosecurity/kubernetes-response-engine)). 

### Community & End Users
#### Community
The Falco project has grown into a vibrant community of end users and project members. The project's Slack team has seen a solid increase in membership and sustained activity. 

| Slack | Pre-Sandbox* (< 10-2018)|Post-Sandbox (>= 10-2018)| |
|-------|-------------------------|-------------------------|-|
|Slack Members|1436|2466 (+71%)||
|Slack WAU|93|221 (+137%)|Weekly Active Users|
|Slack WMPM|34|53 (+55%)|Weekly Members Posting Messages|
|Slack DAU|86 (+152%)|Daily Active Users|
|Slack DMPM|12|23 (+91%)|Daily Members Posting Messages|

**Note:** Post Sandbox #'s as of 2019-10-13

\* Pre-Sandbox is a 29 month period (May 2016-Oct 2018)

**Note:** The slack team is shared by other open source Sysdig projects, but OSS Sysdig accounts for a small % of users and activity.

The Falco team has been holding regular weekly public meetings to provide office hours for end users and community members, as well as to provide a chance for Falco maintainers to plan the upcoming release, and triage issues. More information is available in the project's [community repo](https://github.com/falcosecurity/community).

#### End Users

Falco end users are documented in the project's [ADOPTERS.md](https://github.com/falcosecurity/falco/blob/dev/ADOPTERS.md). Significant users include:

* [Shopify](https://www.shopify.com) - Shopify is the leading multi-channel commerce platform. Merchants use Shopify to design, set up, and manage their stores across multiple sales channels, including mobile, web, social media, marketplaces, brick-and-mortar locations, and pop-up shops. The platform also provides merchants with a powerful back-office and a single view of their business, from payments to shipping. The Shopify platform was engineered for reliability and scale, making enterprise-level technology available to businesses of all sizes. Shopify uses Falco to complement its Host and Network Intrusion Detection Systems.
* [Booz Allen Hamilton](https://www.boozallen.com/) - BAH leverages Falco as part of their Kubernetes environment to verify that work loads behave as they did in their CD DevSecOps pipelines. BAH offers a solution to internal developers to easily build DevSecOps pipelines for projects. This makes it easy for developers to incorporate Security principles early on in the development cycle. In production, Falco is used to verify that the code the developer ships does not violate any of the production security requirements. BAH [is speaking at Kubecon NA 2019](https://kccncna19.sched.com/event/UaWr/building-reusable-devsecops-pipelines-on-a-secure-kubernetes-platform-steven-terrana-booz-allen-hamilton-michael-ducy-sysdig) on their use of Falco.
* [Frame.io](https://frame.io/) - Frame.io is a cloud-based (SaaS) video review and collaboration platform that enables users to securely upload source media, work-in-progress edits, dailies, and more into private workspaces where they can invite their team and clients to collaborate on projects. Understanding what is running on production servers, and the context around why things are running is even more tricky now that we have further abstractions like Docker and Kubernetes. To get this needed visibility into our system, we rely on Falco. Falco's ability to collect raw system calls such as open, connect, exec, along with their arguments offer key insights on what is happening on the production system and became the foundation of our intrusion detection and alerting system. Frame.io is speaking at [Kubecon NA 2019](https://sched.co/UakJ) on their use of Falco.

### CNCF Sponsored Security Audit

The Falco project participated in a CNCF sponsored security audit of the Falco code base. Sysdig seperately sponsored a security audit of Falco's Sysdig OSS dependencies. Both security audits were preformed by [Cure53](https://cure53.de/). The audits found 1 Critical, 4 High, 4 Medium, and 5 Low bugs. 13 of the 14 bugs have been addressed and the audits will be made public when the final bug is fixed. 

### Google Summer of Code

Falco was happy to host a student through the [CNCF's participation](https://summerofcode.withgoogle.com/organizations/5638078861934592/) in the Google Summer of Code (GSoC). [Falco's GSoC project](https://summerofcode.withgoogle.com/projects/#5280508706029568) focused on instrumenting the Falco engine to evaluate Falco's performance. The work performed by GSoC student [Mattia Lavacca](https://github.com/mLavacca) will help Falco improve the engine performance in the future Falco roadmap.

### Features & Roadmap

Falco shipped a number of features and improvements since Sandbox. Below is a partial list.

* eBPF support
* Helm Chart improvements
* CRI-O and Containerd Support
* Performance Improvements
* CVE Response 
* External Security Audit (14 Issues fixed)
* Rules Improvements
* Containerized Builds
* Kubernetes Audit Event support
* MITRE Tagged Rules
* Minimal Images
* Response Engine for AWS & Google Cloud
* Publish Events to SNS or Pub/Sub, Respond with Lambda or Functions
* Falco Operator in Red Hat Operator Hub (operatorhub.io)
* Falco in GKE Marketplace
* Falco Training open sourced
    * https://github.com/falcosecurity/falco-security-workshop
* GRPC Based Alert Output
* Prometheus Exporter


#### Roadmap at Sandbox
The Falco Sandbox proposal included the following roadmap. Shipped and deferred items are noted. 

Short term improvements included:

* Rules Library - Expand the shipped rule set to include rules for commonly deployed applications and CNCF Projects, as well as common compliance rules such as CIS.

  * **SHIPPED** - Container Images/Apps: Nginx, HAProxy, etcd, Java, Node
  * **SHIPPED** - CNCF Projects: Kubernetes, Prometheus, Fluentd, Linkerd
  * **DEFERRED** - CIS Runtime Compliance Rules

Longer term improvements included:

* **SHIPPED** - Prometheus Metrics Exporter - Expose a metrics endpoint to allow collection of metrics by Prometheus. Metrics include # of overall alerts, # of alerts by rule, # of alerts by rule tag.
* **DEFERRED** - Kubernetes networking policy support - Support detecting networking policy violations via the Sysdig kernel module
* Alert Outputs - Add support for additional output destinations to allow Falco to more easily be integrated into a Cloud Native architecture.
  * **SHIPPED** - *Direct webhook support* - Support posting to a generic webhook
  * **SHIPPED** - *Messaging systems* - Support sending messages to a messaging server such as NATS
  * **SHIPPED** - *gRPC* - Support sending to alerts to external systems via gRPC

* **DEFERRED** - Event & Metadata Providers - Support for additional backend providers for the event stream.
* **SHIPPED** - Kubernetes Audit Events - Ingest Kubernetes Audit Events and support rules based on Kubernetes Audit Events.
* **SHIPPED** - Container Runtimes - Support additional container runtime.

* **DEFERRED** - Baselining - Automatic baselining of an application’s “normal” behavior

#### Future Roadmap

* Dropped events & performance improvements
* Improve Prometheus exporter
  * Open Metrics Support
  * Grafana Dashboards 
* Expand gRPC support:
  * Alert Clients
  * Engine services
* Secure by default deployment model
  * Falco
  * Kubeadm, kops, etc
* Improve our cloud native deployment model
  * Falcoctl + Falco API
  * Rules management ([Cloud Native Security Hub](https://github.com/falcosecurity/cloud-native-security-hub))
* [Runtime Policy Interface Specification](https://docs.google.com/document/d/1wKgTodHtYBzUG6hPeg54KeUrbb5m5PHe606g6GmAylM/edit#heading=h.z8wp2cb10kag)

## Incubation Stage Requirements

The Falco project maintainers propose that Falco move to Incubation based on:

* Use in production by 3 end users of note (Shopify, BAH, Frame.io).
* A healthy number of committers and a growing committer base. In addition, a healthy online community.
* Demonstrating a substantial ongoing flow of commits and merged contributions that focused on delivering a defined project roadmap and integrations.
* A clear versioning scheme with dev and stable releases.
