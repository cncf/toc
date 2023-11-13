# Kubescape proposal for CNCF Incubation

## About Kubescape

[Kubescape](https://kubescape.io/) is a security posture management tool, designed to identify and resolve security, misconfiguration, and compliance issues in a Kubernetes environment. 

The project includes tools that can be run on a command line, or in a cluster, or integrated into many other popular tools allowing you to scan workload manifests while they are being developed or integrated, or after they are deployed.

Kubescape was the first tool to automate checking against the [NSA hardening guidance](https://kubescape.io/docs/frameworks-and-controls/frameworks/#nsa) and has since added support for other frameworks (including [MITRE ATT&CK®](https://kubescape.io/docs/frameworks-and-controls/frameworks/#mitre) and the [CIS Benchmark](https://kubescape.io/docs/frameworks-and-controls/frameworks/#cis)). 

It also includes comprehensive vulnerability scanning and reporting, allowing you to see the state of vulnerabilities detected in your containers.

## Kubescape in the CNCF Sandbox

Kubescape was proposed to the CNCF Sandbox on 31 October 2022 and accepted on December 13.

Since that date, the maintainers have added substantial new features:

* Support for CIS benchmarks for EKS and AKS
* Vulnerability reachability testing using an eBPF agent
* Compliance and container scan results stored as Kubernetes resources inside the cluster
* Scanning container images for vulnerabilities from the CLI
* Reporting on the vulnerabilities of all the images in a cluster
* A new overview security scan, which helps you set a baseline for cluster security
* Highlighting of high-risk workloads: those that could do the most damage if they are compromised
* Improved display output
* A new capability-based Helm chart
* Per-workload, per-namespace and per-cluster Prometheus metrics
* Alerting through Prometheus Alertmanager

### Community & growth

In our time in the Sandbox, we have: 

* 2500 new GitHub stars (currently 9115)
* 1750 [PRs](https://github.com/pulls?q=is%3Apr+created%3A%3E%3D2022-12-13+user%3Akubescape+) merged
* 270 [issues](https://github.com/issues?q=is%3Aissue+created%3A%3E%3D2022-11-05+user%3Akubescape) raised 
* Mentored 6 students through LFX

### Neutrality

At the time of submission, Kubescape’s creators ARMO operated a commercial service called Kubescape Cloud. As was promised at the time of submission, that service was renamed to “ARMO Platform” to ensure vendor neutrality.

Over the course of our time in the Sandbox, the Kubescape maintainers have separated Kubescape from ARMO:

* Data that was only available in the ARMO Platform is now available in-cluster. 
* The interface between Kubescape and ARMO Platform has been generalized into a provider interface, with documentation on how anyone can build their own provider

## Incubation stage requirements

### Used successfully in production by at least three independent direct adopters

Our download numbers suggest Kubescape is used by thousands of end users, either directly, or as customers of commercial security solutions such as ARMO Platform and Jit.

### Have a healthy number of committers

Kubescape has 9 maintainers across the [main repository](https://github.com/kubescape/kubescape/blob/master/MAINTAINERS.md) and [rule library](https://github.com/kubescape/regolibrary/blob/master/MAINTAINERS.md).  

[71 committers have had PRs merged in the last 12 months](https://kubescape.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=merged_prs&var-repogroup_name=All&var-country_name=All&var-companies=All) and [245 contributors have interacted with the project on GitHub](https://kubescape.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-companies=All).

### Demonstrates a substantial ongoing flow of commits and merged contributions

Please see our DevStats instance, specifically the [merged PRs list for the last 12 months](https://kubescape.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20year&var-metric=merged_prs&var-repogroup_name=All&var-country_name=All&var-companies=All), which shows 1756 PRs merged by 71 people.

### Clear versioning scheme

Individual components of Kubescape are versioned independently:

* [The Kubescape CLI and microservices](https://github.com/kubescape/kubescape/tags)
* [The Helm chart](https://github.com/kubescape/helm-charts/releases)
* [The control library](https://github.com/kubescape/regolibrary/releases)

All use [semantic versioning](https://semver.org/).

### Clearly documented security processes explaining how to report security issues to the project, and describing how the project provides updated releases or patches to resolve security vulnerabilities

Kubescape uses the [GitHub security reporting process](https://github.com/kubescape/kubescape/blob/master/SECURITY.md), with an SLO of 7 days for contact and 90 days for disclosure.

The project aims to release every 2 weeks, and only the latest version is supported.

## Alignment with the CNCF Mission

In order for cloud native computing to be ubiquitous, it must be secure. Kubescape exists solely to improve the security posture of workloads running on Kubernetes; it does not target workloads outside Kubernetes or Cloud Native.

The ultimate goal of the Kubescape project is to provide coverage of the full space of Kubernetes administrative security concerns. 

Kubescape interacts with a number of CNCF projects:

* Kubernetes: data source/operating environment
* Helm: installation method
* Open Policy Agent: rules engine
* Inspektor Gadget: eBPF engine
* Prometheus: metrics export
* OpenTelemetry: telemetry engine

The project also includes integrations for Argo, Backstage and Flux.

## Future plans

The 2024 roadmap includes adding new functionality in the additional areas of network security and runtime/threat detection. Like our existing integrations, this will be a combination of new work and best-of-breed open source engines.