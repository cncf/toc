# Project Description

Every organization has unique policies that affect the entire stack. These policies are vital to long term success because they codify
important requirements around cost, performance, security, legal regulation, and more. At the same time, organizations often rely on
tribal knowledge and documentation to ensure that policies are enforced correctly. While these approaches are known to be error prone,
they exist because systems frequently lack the flexibility and expressiveness required to automate policy enforcement.

The Open Policy Agent (OPA) is a general-purpose policy engine that enables unified, context-aware policy enforcement across the stack.
OPA empowers administrators with greater control and flexibility so that organizations can automate policy enforcement at any layer.

At the core of OPA is a high-level declarative language (and runtime) that allows administrators to enforce policies across multiple
domains such as API authorization, admission control, workload placement, storage, and networking. OPA’s language is purpose-built for 
expressing policy decisions. The language has rich support for processing complex data structures as well as performing search and 
aggregation across context required for policy decisions. The language also provides support for encapsulation and composition so that 
complex policies can be shared and re-used. Finally, the language includes a standard library of built-in functions for performing math 
operations, string manipulation, date/time parsing, and more.

With OPA, policy decisions are decoupled from applications and services so that policy logic can be modified easily and upgraded 
on-the-fly without requiring expensive, time consuming development and release cycles.

OPA provides simple APIs to offload policy decisions from applications and services. Policy decisions are computed by OPA and returned 
to callers as structured data. Callers integrate with OPA by executing policy queries that can include arbitrary input values. For 
example, an API gateway might supply incoming API requests as input and expect boolean values (representing allow/deny decisions) as 
output. On the other hand, a container orchestrator might supply workload resources as input and expect a map of clusters and weights 
to drive workload placement as output. See the appendix for sample policies that cover these use cases.

OPA itself is written in Go and can be integrated as a library, host-level daemon, or sidecar container. OPA provides APIs to load and 
manage policies as well as external data. Finally, OPA provides rich tooling to support the development, testing, and debugging of 
policies.

Since the initial release in July 2016, OPA’s mission has been to provide a powerful building block that enables policy-based control 
across the stack. OPA’s roadmap for the next 12 months includes improvements to the language, integration with Google’s CEL, expansion 
of the standard policy library, as well as continued hardening and performance optimization.

**Sponsor from TOC:** Ken Owens, Brian Grant

**Preferred Maturity Level:** Sandbox

**License:** Apache License v2

# Source Control
https://github.com/open-policy-agent/opa
https://github.com/open-policy-agent/library

# External Dependencies

github.com/ghodss/yaml             MIT License
github.com/gorilla/mux             BSD 3-clause "New" or "Revised" License
github.com/mattn/go-runewidth      MIT License
github.com/olekukonko/tablewriter  MIT License
github.com/peterh/liner            MIT License
github.com/pkg/errors              BSD 2-clause "Simplified" License
github.com/sirupsen/logrus         MIT License
github.com/spf13/cobra             Apache License 2.0
github.com/spf13/pflag             BSD 3-clause "New" or "Revised" License
golang.org/x/crypto/ssh/terminal   BSD 3-clause "New" or "Revised" License
golang.org/x/sys/unix              BSD 3-clause "New" or "Revised" License
gopkg.in/fsnotify.v1               BSD 3-clause "New" or "Revised" License
gopkg.in/yaml.v2                   Apache License 2.0


**Initial Committers:** Torin Sandall and Tim Hinrichs from Styra (since creation), Tristan Swadell from Google (since May 2017)

**Infrastructure Requests:** None initially. CI is currently hosted on Travis and covered by the free tier for open source projects. In 
the future, we would like to leverage CNCF test clusters for system testing integrations built as part of the OPA project.

**Communication Channels:**
Slack: http://slack.openpolicyagent.org

**Issue Tracker:** https://github.com/open-policy-agent/opa/issues

**Website:** http://www.openpolicyagent.org

# Release Methodology and Mechanics

We currently use numbered releases with the changelog and binaries published to https://github.com/open-policy-agent/opa/releases. 
The release process is partially automated with manual portions assisted by scripts. The current release process is documented here: 
https://github.com/open-policy-agent/opa/blob/master/docs/devel/RELEASE.md. The release schedule is somewhat ad-hoc, aligned around 
large feature boundaries.

**Social Media Accounts:**
Twitter: https://twitter.com/openpolicyagent

# Community Size and any Existing Sponsorship

Adopters:
Netflix
Medallia
Schuberg Phillis
Huawei
More: At least one large financial institution and one large online retailer is testing OPA

Integrations:
Kubernetes (Use cases: federated resource placement, admission control)
Docker (Use cases: Docker engine authorization)
Istio (Use cases: microservice API authorization)
Linkerd (Use cases: microservice API authorization)
OpenSDS (Use cases: storage scheduling)
Terraform (Use cases: risk management on terraform plans)
PAM (Use cases: SSH and sudo authorization)
Cloud Foundry buildpack to enable microservice API authorization

**Sponsors**
https://www.styra.com
https://www.firebase.com (Google)

**Numbers:**
3 active contributors currently (2 from Styra, 1 from Google), with 8 other contributors over past 12 months.
80 stars
49 members on Slack
31 releases

# Statement of Alignment with CNCF Mission

As cloud native technology matures and enterprise adoption increases, the need for policy-based control has become apparent. OPA 
provides a powerful building-block that enables fine-grained, expressive policy enforcement.  As such, we think that OPA would be a 
great for fit for the CNCF

# Benefits to the CNCF

The ecosystem must provide solutions to control who can do what across microservice deployments because legacy approaches to access 
control do not satisfy the requirements of modern environments. OPA provides a purpose-built language and runtime that can be used to 
author and enforce authorization policy. As such, we feel that OPA will complement the CNCF’s portfolio and help accelerate adoption of 
cloud native technology in enterprises. In the longer term, we think that enterprises will benefit from a unified approach to policy 
enforcement can be applied across the stack.

# What does OPA need from the CNCF

OPA needs a well respected, vendor-neutral home that can help serve as a rallying point around policy as code. In addition to increased 
visibility, we hope that inclusion in the CNCF will foster communication between OPA and other projects in the ecosystem. As the project
grows, we would want to leverage the CNCF’s expertise around project governance and community standards as those are fundamental to the 
long term success of the project.

The project does not have any infrastructure requests at this time. CI is currently hosted on Travis and covered by the free tier for 
open source projects. In the future, we would like to leverage CNCF test clusters for system testing integrations built as part of the 
OPA project.

# Appendix A: REST API Authorization Example

This sample shows two simple rules that enforce an authorization policy on an API that serves salary data. In English, the policy says 
that employees can see their own salary and the salary of any of their reports.

allow {
	input.method = "GET"
	input.path = ["salary", employee_id]
	input.user = employee_id
}


allow {
	input.method = "GET"
	input.path = ["salary", employee_id]
	input.user = data.management_chain[employee_id][_]
}


The first rule allows employees to GET their own salary. The rule shows how you can use variables in rules. In that rule, employee_id is
a variable that will be bound to the same value across the last two expressions.

The second rule allow employees to GET the salary of their reports. The rule shows how you can access arbitrary context (e.g., JSON data)
inside the policy. The data may loaded into the policy engine (and cached) or it may be external and fetched dynamically.


# Appendix B: Cluster Placement Example

This sample shows a simple rule that generates a set of clusters that a workload may be deployed to. The workload is provided as input 
to policy. In English, the policy says that workloads must be placed on clusters that satisfy the workload’s jurisdiction requirements.

desired_clusters = {name |
	cluster = data.clusters[name]
	satisfies_jurisdiction(input.deployment, cluster)
}


satisfies_jursidiction(deployment, cluster) {
	deployment.jurisdiction = "europe"
	startswith(cluster.region, "eu")
} else {
	not deployment.jurisdiction
}

This example shows how logic can be composed across rules and functions.
