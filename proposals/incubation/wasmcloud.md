# wasmCloud

## Name of project

wasmCloud

## Project description

wasmCloud helps you build globally distributed applications that reach from the cloud to the edge. Leveraging WebAssembly's security, portability, and performance, the platform helps you compose applications from tiny, independent building blocks loosely coupled to capabilities.

Deciding whether you're building a serverless function, a microservice, a globally distributed app, or a monolith is now a knob you can turn at runtime rather than a one-time decision you're stuck with until your application retires.

wasmCloud is designed around the following core tenets:

- Dead simple distributed applications
- Run anywhere
- Secure by default
- Productivity for both developers and operations

**You should not have to change your design, architecture, or your programming environment as you move from concept to production.**

### Impact and Features

wasmCloud is a cloud native application runtime for distributed WebAssembly applications. Developers write code in their preferred language, compile it to a WebAssembly component, and deploy it to the wasmCloud runtime. The wasmCloud runtime provides a secure, distributed, and portable execution environment for WebAssembly modules.

wasmCloud is designed to be a _better_ way to write, deploy and manage applications. Using WebAssembly for the unit of compute removes the responsibility of the user to consider the target operating system or architecture. Explicit communication protocols to access resources outside of the WebAssembly sandbox provide a secure-by-default environment and keeps developers from creating brittle specific security policies to lock down an application. The use of NATS for inter-service communication lets microservices and monoliths written as wasmCloud applications be deployed anywhere and communicate seamlessly whether they are running on the same machine or on opposite sides of the globe. NATS also provides common service-mesh capabilities without the heavyweight management of a layer 7 mesh like load balancing, immediate failover and service discovery. These concepts are abstracted for the individual developer and are handled automatically (but can be tuned for advanced users.)

wasmCloud integrates directly into modern cloud native environments. At the base level the wasmCloud runtime is a native binary that can run on x86_64 and aarch64 architectures (with more like riscv64 coming soon). We also distribute a multi-platform container image and a helm chart for Kubernetes deployment. Being compatible with but not dependent upon containers and Kubernetes allows wasmCloud to be deployed in a variety of environments from the cloud to the edge, on different clouds and using different orchestrators. wasmCloud supports running its WebAssembly binaries by downloading from OCI compliant registries (AzureCR, GHCR, etc), exporting traces to [OTEL](https://opentelemetry.io/) compatible collectors, and defines its declarative application manifests using the [Open Application Model](https://oam.dev/). All events published from wasmCloud use the CloudEvents format. Soon, wasmCloud will add metrics support using [Prometheus](https://prometheus.io/).

## Preferred maturity level

Incubation

## Sponsor / Advisor from Toc

Ricardo Rocha

## License

All wasmCloud code is licensed under Apache 2: [https://github.com/wasmCloud/wasmCloud/blob/main/LICENSE](https://github.com/wasmCloud/wasmCloud/blob/main/LICENSE)

## Source control

All wasmCloud code is open source under the [https://github.com/wasmcloud/](https://github.com/wasmcloud/) organization, which belongs to CNCF enterprise GitHub.

## Issue tracker

wasmCloud issues are tracked under their specific repositories in the [https://github.com/wasmcloud/](https://github.com/wasmcloud/) organization, most issues from bugs to feature requests to RFCs are filed under [wasmcloud/wasmcloud](https://github.com/wasmCloud/wasmCloud/issues). Issues that represent the high level goals and roadmap of the project are then tracked on our [wasmCloud Roadmap](https://github.com/orgs/wasmCloud/projects/7/views/3).

## Website

Our blogs, community meeting notes and documentation are available on [https://wasmcloud.com/](https://wasmcloud.com/).

## Community size and any existing sponsorship

wasmCloud joined the CNCF as a Sandbox Application Runtime project on July 13th, 2021. Since joining the CNCF as a sandbox project we've seen steady community growth and adoption.

- GitHub organization stars: 2300+
- Contributors: 440+
- Pull Requests: 2000+
- Commits: 9300+

### Company Contributors

wasmCloud has accepted contributions from the following companies (in order of contribution volume):

- Capital One Financial Corporation
- Amazon
- Bytedance Ltd
- wasmCloud
- Lucid
- Google LLC
- Independent
- Microsoft Corporation
- Roku
- Red Hat Inc.
- Stacklet
- Jammin Music
- Critical Stack
- Red Badger Consulting Limited
- DaoCloud Network Technology Co. Ltd.
- NetApp Inc
- Cosmonic
- Conductor
- SnT
- PingCAP
- CapTech Venture
- Abstract Tokenization
- One Agency
- HMBradley
- arvreal
- DeMark
- Machine Sciences
- Sharply
- Spora Solutions
- Vino
- Giant Swarm GmbH
- Leafly
- NielsenIQ
- Change Healthcare
- earth
- Codemancers
- Fugro
- Deutsche Telekom
- Adobe
- Povio Labs
- SecureDocs
- Datadog Inc
- Sourcēvo
- Volvo
- Antgroup
- Enabled Business Solutions Limited
- Liquid Reply
- Shape Security
- Branch Metrics Inc.
- CJ Affiliate
- Fidelity Investments
- Genobank.io
- Heinlein
- Instana Inc.
- Intel Corporation
- Letter
- Nascent LLC
- Northfield
- phpMyAdmin
- Recurly Inc.
- SUSE LLC
- Tinfoil

## Adopters

Our current self-reported adopters are listed on our [wasmCloud adopters document](https://github.com/wasmCloud/wasmCloud/blob/main/ADOPTERS.md). Our adopters run wasmCloud at varying levels of scale. Cosmonic runs wasmCloud in production at scale to host the entire backend of their platform, and has been a major contributor to the project.

## Communication channels (slack, irc, mailing lists)

We primarily communicate on Slack, available publicly at [https://slack.wasmcloud.com](https://slack.wasmcloud.com). We also monitor the [#wasmcloud](https://cloud-native.slack.com/archives/C027YTXEYFL) channel on the CNCF Slack.

We run what we call "wasmCloud Wednesdays" every Wednesday at 1pm EST and have not missed a meeting since we started on 18 Nov 2020. It is on the [CNCF community calendar](https://tockify.com/cncf.public.events/detail/662/1698253200000).

### Appendix

- Adobe published a case study [Better together: A Kubernetes and Wasm case study](https://www.cncf.io/blog/2022/11/17/better-together-a-kubernetes-and-wasm-case-study/) as a CNCF end-user blog detailing their experience using wasmCloud and Kubernetes in tandem for more efficient, secure applications.
- OSTIF + Trail of Bits recently completed a security audit of wasmCloud and found no medium, high or critical vulnerabilities. The full report is available at [https://ostif.org/ostif-has-completed-a-security-audit-of-wasmcloud/](https://ostif.org/ostif-has-completed-a-security-audit-of-wasmcloud/).

## Project logo in svg format

All logos can be found in the wasmCloud branding repository:
[https://github.com/wasmCloud/branding](https://github.com/wasmCloud/branding)

The primary logo used on the CNCF landscape is located at [https://github.com/cncf/landscape/blob/master/hosted_logos/wasmcloud.svg](https://github.com/cncf/landscape/blob/master/hosted_logos/wasmcloud.svg)

## CNCF Due Diligence (DD) Document

The wasmCloud CNCF Due Diligence document is located in a Google document: [https://docs.google.com/document/d/1MVY2mPqoAbHOGSCKwELS8KwbHkW3v8QS0mMSzHCSi00](https://docs.google.com/document/d/1MVY2mPqoAbHOGSCKwELS8KwbHkW3v8QS0mMSzHCSi00)
