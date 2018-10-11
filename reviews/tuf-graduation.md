# TUF Graduation Application

The Update Framework (TUF) helps developers maintain the security of a software update system, even against attackers that compromise the repository or signing keys. TUF provides a flexible framework and specification that developers can adopt into any software update system.

## TUF fulfills all the incubating and graduation criteria:

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

* TUF is used in production by Docker, IBM, VMware, Cloudflare, Kolide, DigitalOcean, RedHat, DataDog, and many others.  [Uptane](https://uptane.github.io/), the automotive variant of TUF, is sold by major automotive vendors and is being adopted by automakers with over 1/3 of the cars on US roads in their forthcoming cars.  Uptane is included as part of [Automotive Grade Linux](https://www.automotivelinux.org/), another Linux Foundation project.
* More detail about adoption, including some of the in progress adoptions, can be found here:  https://theupdateframework.github.io/adoptions.html

### Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

As an intentionally minimal security specification / framework, we intentionally do not have a high degree of feature additions in the project.  Effort either comes on the implementations, such as the Python reference implementation or Notary, or the specification itself.  

* Python reference implementation

TODO

* Notary

TODO

* TUF Specification

TODO

### Demonstrate a substantial ongoing flow of commits and merged contributions.


NEEDS WORK.  From Prometheus
* https://github.com/prometheus/prometheus/graphs/contributors

In all official Prometheus repositories, we have had 850+ unique contributors with a total of 12k+ commits so far.

### Have committers from at least two organizations.

NEEDS WORK.  From Prometheus
We have [17 committers](https://github.com/prometheus/docs/blob/master/content/governance.md#team-members) from ~10 organizations:

* [Ben Kochie](https://github.com/SuperQ) ([GitLab](https://about.gitlab.com/))
* [Björn Rabenstein](https://github.com/beorn7) ([SoundCloud](https://soundcloud.com/))
* [Brian Brazil](https://github.com/brian-brazil) ([Robust Perception](https://www.robustperception.io/))
* [Conor Broderick](https://github.com/Conorbro) ([Robust Perception](https://www.robustperception.io/))
* [Fabian Reinartz](https://github.com/fabxc) ([CoreOS](https://coreos.com/) / [Red Hat](https://www.redhat.com/))
* [Frederic Branczyk](https://github.com/brancz) ([CoreOS](https://coreos.com/) / [Red Hat](https://www.redhat.com/))
* [Goutham Veeramachaneni](https://github.com/Gouthamve) (Independent)
* [Johannes Ziemke](https://github.com/discordianfish) ([Latency.at](https://latency.at/) / Independent)
* [Julius Volz](https://github.com/juliusv) (Independent)
* [Matt Layher](https://github.com/mdlayher) ([DigitalOcean](https://www.digitalocean.com/))
* [Matthias Rampke](https://github.com/matthiasr) ([SoundCloud](https://soundcloud.com/))
* [Max Inden](https://github.com/mxinden) ([CoreOS](https://coreos.com/) / [Red Hat](https://www.redhat.com/))
* [Richard Hartmann](https://github.com/RichiH) ([SpaceNet](https://www.space.net/))
* [Steve Durrheimer](https://github.com/sdurrheimer) ([Netapsys](https://www.netapsys.fr/))
* [Stuart Nelson](https://github.com/stuartnelson3) ([DigitalOcean](https://www.digitalocean.com/))
* [Tobias Schmidt](https://github.com/grobie) ([SoundCloud](https://soundcloud.com/))
* [Tom Wilkie](https://github.com/tomwilkie) ([Kausal](https://kausal.co/))

### Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

We have both the basic best practices badge and the silver badge.
https://bestpractices.coreinfrastructure.org/projects/1351

### Adopt the CNCF Code of Conduct.

TUF follows the CNCF code of conduct.  https://github.com/theupdateframework/tuf/blob/develop/docs/CODE-OF-CONDUCT.md
This is linked multiple places, including the footer of [TUF's homepage](https://theupdateframework.github.io/).

### Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

These documents may be found here:
* https://github.com/theupdateframework/tuf/blob/develop/docs/GOVERNANCE.md
* https://github.com/theupdateframework/tuf/blob/develop/docs/CONTRIBUTORS.rst

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

Our website has an [adoptions page](https://theupdateframework.github.io/adoptions.html) on it that lists the different projects.  We also have an [ADOPTERS.MD](https://github.com/theupdateframework/tuf/blob/develop/docs/ADOPTERS.md) which contains much the same information.

FALCO INFO FOLLOWS

=== Falco CNCF Sandbox Project Proposal
 *Name of Project:* Falco
 *Description:*
 Highly distributed and dynamic architectural patterns such as microservices are proving that traditional models of application and network security alone do not meet today’s current needs. Additionally, the increasing level of regulation being introduced (General Data Protection Regulation, or GDPR, for instance) to any business with a digital presence makes security more important than ever. Organizations must quickly respond to exploits and breaches to minimize financial penalties introduced by such regulation, yet the dynamic nature of modern Cloud Native architectures make it extremely difficult for organizations to keep pace.
 Falco seeks to solve this problem by shortening the security incident detection and response cycle in microservices architectures. Falco provides runtime security for systems running container workloads to detect behavior that is defined as abnormal. Falco can be broken into three areas: 
 *Event & Metadata Providers* - inputs of events to the rules engine.
 * Sysdig Kernel Module - provides a stream of system call events for Linux based systems. 
* Kubernetes API Server - provides metadata for Kubernetes resources such as Namespace, Deployment, Replication Controllers, Pods, and Services. 
* Marathon - provides metadata for Marathon resources.
* Mesos - provides metadata for Mesos resources.
* Docker - provides metadata for containers running under the Docker container runtime.
 *Rules Engine & Condition Syntax* - Falco implements a rules engine that supports the following rule syntax. 
 * https://github.com/draios/falco/wiki/Falco-Rules#conditions[Sysdig Filter Syntax] - Falco supports the Sysdig filter syntax used for filtering system call events from the Sysdig kernel module. This syntax also supports filtering on metadata from sources such as container runtimes, Kubernetes, Mesos, and Marathon. 
 *Notification Outputs* - Falco’s rules engine will send alerts when rule conditions are met. The following output destinations are currently supported.
 * Stdout, Log file, Syslog - These can be aggregated using Fluentd or similar
* Command Execution - Falco can execute a command, passing the alert in via stdin
 For example, by leveraging the Sysdig kernel module’s capabilities of tapping into system calls from the Linux kernel, rules can be written to detect behavior seen as abnormal. Through the system calls, Falco can detect events such as:
 * A Kubernetes Pod running in a Deployment labeled ‘node-frontend’ begins running processes other than ‘node’. 
* A shell is run inside a container
* A container is running in privileged mode, or is mounting a sensitive path like /proc from the host.
* A server process spawns a child process of an unexpected type
* Unexpected read of a sensitive file (like /etc/shadow)
* A non-device file is written to /dev
* A standard system binary (like ls) makes an outbound network connection
 When a rule condition is met, Falco can either log an alert to a file, syslog, stdout, etc, or trigger an external program. This allows an automated system to respond to compromised containers or container hosts. This automated system could stop or kill containers identified as compromised, or mark container hosts as tainted to prevent workloads from being scheduled on the compromised host. 
 *Value to the Cloud Native Operating Model*
 As Cloud Native starts to become the defacto operating model for many organizations, the security of this model is often the first thing many organizations seek to address. The Cloud Native model seeks to empower developers to be able to rapidly package applications and services in containers, then quickly deploy them to platforms such as Kubernetes. This model seeks to remove the traditional points of friction in operations by providing a consistent deployment paradigm and abstraction of the underlying infrastructure. The challenge for many organizations is that applications packaged as containers are often a black box to downstream teams in terms of 1) what is packaged inside the container, and 2) operations any processes might perform once the application is running. 
 Currently there are several prescribed methods for building security into the Cloud Native workflow:
 * *Image Chain of Trust*
** Scan images as part of a deployment process, such as GitOps, to verify their contents and check for known vulnerabilities (for example Anchore or Clair).
** Cryptographically sign images and restrict container runtimes to only run trusted images. (eg Notary)
** Restrict which container registries images can be pulled from.
* *Admittance Control*
** Cryptographically verifiable identities to restrict/allow workloads to run based on a defined policy (eg SPIFFE).
** Leveraging Service Meshes to control what workloads can join a particular service. 
* *Orchestrator/Infra Security*
** Role Based Access Control to restrict access to the orchestrator API services.
** General best practices for securing the orchestrator entry points.
** Network Policy API and CNI Plugins
** Linux Security Module support.
** PodSecurity Policies
* *Runtime Security*
** Detect abnormal behavior inside a workload and take appropriate action, such as telling the orchestrator to kill the workload, thus shortening the security “detect-response” cycle. (eg Falco)
* *Workload Access Control Policies*
** Policies controlling the network activity of workloads and restricting inter-workload communication.
** Policies controlling the API endpoints available to workloads (eg Cilium)
 Each prescribed method provides an additional level of protection, but one method by itself does not provide a complete security solution. Image Chain of Trust for instance is a “point in time” method of providing security. In other words, the container image is considered “secure” when the image scanning process completes successfully, but anytime after that it may become “insecure” once new exploits or vulnerabilities are discovered. 
 Additionally, while container images are considered immutable when built, once a container is created from the image, the process inside the container can modify the container’s instantiation of the root filesystem. Some best-practices suggest starting containers with a read-only root filesystem to prevent this, but this method has its own problems. For instance, the “standard” Node.js image needs to write to the root filesystem to create a number of files (lock files for instance) when node starts. Runtime Security seeks to mitigate this problem by watching what changes may be made once a container is running, and taking action on abnormal behavior. 
 Currently the most of the options for runtime security are limited to proprietary solutions that limits the ability to take advantage of the larger open source software ecosystem. Falco is unique in that its open approach allows for a broader community to define and share rule sets for common security exploits. This open approach also provides the opportunity for a faster response time to newly discovered exploits by providing the ability to share new rules for these exploits as they are discovered.   
 *Falco Roadmap*
 Short term improvements include:
 * *Rules Library* - Expand the shipped rule set to include rules for commonly deployed applications and CNCF Projects, as well as common compliance rules such as CIS.
 ** Container Images/Apps: Nginx, HAProxy, etcd, Java, Node
** CNCF Projects: Kubernetes, Prometheus, Fluentd, Linkerd
** CIS Runtime Compliance Rules
 Longer term improvements include:
 * *Prometheus Metrics Exporter* - Expose a metrics endpoint to allow collection of metrics by Prometheus. Metrics include # of overall alerts, # of alerts by rule, # of alerts by rule tag.
* *Kubernetes networking policy support* - Support detecting networking policy violations via the Sysdig kernel module
* *Alert Output* - Add support for additional output destinations to allow Falco to more easily be integrated into a Cloud Native architecture.
** *Direct webhook support* - Support posting to a generic webhook +
** *Messaging systems* - Support sending messages to a messaging server such as NATS +
** *gRPC* - Support sending to alerts to external systems via gRPC
 * *Event & Metadata Providers* - Support for additional backend providers for the event stream.
* *Kubernetes Audit Events* - Ingest Kubernetes Audit Events and support rules based on Kubernetes Audit Events.  +
* *Container Runtimes* - Support additional container runtime.
 * *Baselining* - Automatic baselining of an application’s “normal” behavior
	
 *Planned Advocacy Work*
 Beyond the engineering work planned, there is also work planned to improve the awareness of Falco in the Cloud Native ecosystem.
 * *Workshops on Falco:* As the project’s main sponsor, Sysdig has been investing in workshops focused on Container Troubleshooting and Container Forensics that include sections on Falco and CNCF projects such as Kubernetes. These workshops will be expanded to include more exercises on writing rules for applications, testing workflow for rule writing, and incorporation of Falco in CD workflows such as GitOps, etc.
* *Documentation Improvements*: Improve documentation with regard to writing rules including out of the box macros, lists, and rules provided by Falco.
* *Documenting Use Cases:* Document existing use cases around using Falco with other projects to deliver a complete end to end solution.
* *Events:* Conference and Meetup presentations to help educate the community on security in the Cloud Native landscape, and to help new community members how to implement Cloud Native based architectures in a secure fashion.    
 *Current CNCF Ecosystem Integrations:*
 *Containerd and rkt*
 Falco can detect containers running in both containerd and rkt container runtimes.
 *Kubernetes*
 Falco can communicate with the Kubernetes API to pull Namespace, Deployment, Service, ReplicaSet, Pod, and Replication controller information such as name and labels. This data can be used to create rule conditions (e.g. k8s.ns.name = mynamspace) as well as used as an outputted field in any generated alerts. 
 A common deployment method for Falco in the Cloud Native landscape is to deploy it as a Daemon Set running in Kubernetes. The Falco project provides releases packaged as containers and provides a Daemon Set example for end users to deploy Falco.
 Docker Hub: https://hub.docker.com/r/sysdig/falco/[https://hub.docker.com/r/sysdig/falco/]
 Kubernetes Daemon Set: https://github.com/draios/falco/tree/dev/integrations/k8s-using-daemonset[https://github.com/draios/falco/tree/dev/integrations/k8s-using-daemonset]
 Helm chart: https://github.com/helm/charts/tree/master/stable/falco[https://github.com/helm/charts/tree/master/stable/falco]
 *Fluentd*
 Falco can also leverage Fluentd from the CNCF ecosystem. Falco alerts can be collected from logs or stdout by Fluentd and the alerts can be aggregated and analyzed. An example of using Falco with Fluentd, Elasticsearch, and Kibana can be found on the Sysdig Blog.
 https://sysdig.com/blog/kubernetes-security-logging-fluentd-falco/[https://sysdig.com/blog/kubernetes-security-logging-fluentd-falco/]
 *NATS*
 A https://github.com/sysdiglabs/falco-nats[proof of concept] was created showing publishing of Falco alerts to a NATS messaging server. These alerts can be subscribed to by various programs to process and take action on alerts. In the proof of concept, Falco alerts published to NATS triggered a Kubeless function to delete an offending Pod.
 *Sponsors from TOC:* Quinton Hoole, Brian Grant
 *Preferred maturity level:* Sandbox
 *Unique identifier:* falco
 *Current Project Sponsor:* https://sysdig.com/opensource/[Sysdig]
 *License:*** **Apache License v 2 (ALv2)
 *Code Repositories:*
Code is currently hosted by Sysdig:
https://github.com/draios/falco[https://github.com/draios/falco]
 The code will move to a vendor netural github organization at:
https://github.com/falcosecurity[https://github.com/falcosecurity]
 *External Code Dependencies* +
External dependencies of Falco are listed below:
 |===
|*Software*|*License*|*Project Page*
 |libb64|Creative Commons|http://libb64.sourceforge.net/[http://libb64.sourceforge.net/]
|curl|MIT/X|https://curl.haxx.se/[https://curl.haxx.se/]
|jq|MIT|https://stedolan.github.io/jq/[https://stedolan.github.io/jq/]
|libyaml|MIT|https://pyyaml.org/wiki/LibYAML[https://pyyaml.org/wiki/LibYAML]
|lpeg|MIT|http://www.inf.puc-rio.br/\~roberto/lpeg/[http://www.inf.puc-rio.br/~roberto/lpeg/]
|luajit|MIT|http://luajit.org/luajit.html[http://luajit.org/luajit.html]
|lyaml|MIT|https://github.com/gvvaughan/lyaml[https://github.com/gvvaughan/lyaml]
|ncurses|MIT?|https://www.gnu.org/software/ncurses/[https://www.gnu.org/software/ncurses/]
|openssl|OpenSSL & SSLeay|https://www.openssl.org/source[https://www.openssl.org/source]
|yamlcpp|MIT|https://github.com/jbeder/yaml-cpp[https://github.com/jbeder/yaml-cpp]
|zlib|zlib|https://www.zlib.net/zlib.html[https://www.zlib.net/zlib.html]
|sysdig|ALv2|https://github.com/draios/sysdig[https://github.com/draios/sysdig]
|tbb|ALv2|https://www.threadingbuildingblocks.org/[https://www.threadingbuildingblocks.org/]
|===
 *Committers:* 16
 *Users of Note:*
 Cloud.gov:
 * https://cloud.gov/docs/apps/experimental/behavior-monitoring/[Dynamic behavior monitoring in Cloud.gov] 
* https://www.youtube.com/watch?v=wFQOXMcZnQg[Detecting tainted apps in Cloud Foundry]
* https://github.com/cloudfoundry-community/falco-boshrelease[falco-boshrelease]
 *Community Communication:* 
Slack is the preferred form of communication. Sysdig runs a Slack team for its open source projects and hosts a #falco channel under that Slack team:
 Slack team: https://sysdig.slack.com[https://sysdig.slack.com] +
Falco Channel: https://sysdig.slack.com/messages/C19S3J21F/[https://sysdig.slack.com/messages/C19S3J21F/]
 *Website/Blog:*
 The website is currently hosted by Sysdig, under the Open Source section of the website: https://sysdig.com/opensource/falco[https://sysdig.com/opensource/falco] 
 Blog posts related to Falco are currently posted to the Sysdig Blog. https://sysdig.com/blog/tag/falco/[https://sysdig.com/blog/tag/falco/]
 The Falco website and blog will be moved to: https://falco.org[https://falco.org]
 *Release Cadence:*
 Minor releases quarterly, Patch releases as frequent needed (Minor and Patch used as defined by https://semver.org/[semantic versioning].)
 *Statement on alignment with CNCF mission:*
 With the number of systems under management increasing at a greater and greater rate, and regulation becoming more common, new approaches are required with regards to security that allows organizations to automatically manage the “detection & response” security cycle. Innovations in Cloud Native technologies allow this automatic approach to security more and more feasible.
 Falco aligns with the CNCF mission statement by:
 * Focusing on containers first: Falco was built with the assumption that containers are the method in which modern applications would be run. Falco has included since its inception the ability to identify containerized processes and apply rules to these processes.
* Enabling the CNCF ecosystem by including Cloud Native best practices: The https://github.com/draios/falco/blob/dev/rules/falco_rules.yaml[default Falco rule set] focuses on container anti-patterns, or rather common mistakes that new users tend to do when deploying a Cloud Native application in containers. While currently these rules focuses on containers and container runtimes, additional rule sets can be written for CNCF projects, and application runtimes in the CNCF Landscape. This work is on the Falco roadmap, and could be easily done wby the broader CNCF community.  
* Falco’s goal is to provide a modular, composable system that allows easy integration with other CNCF projects or open source projects. This idea of composability allows for operators of Cloud Native platforms to easily build systems to manage the security of the platform, while maintaining a high degree of flexibility and maintaining the Cloud Native developer velocity. 
