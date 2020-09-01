# Cloud Custodian CNCF Sandbox Proposal

## Name of project

Cloud Custodian

## Project description

### Introduction

Cloud Custodian is a YAML DSL based stateless rules engine for cloud audit, management, and governance.

Some of its features:
* Detective controls to ensure adherence to organizational policies
* Supports real-time detection, reporting/notification, and remediation 
* Consistent outputs and telemetry (blob, logs, trace, metrics) across policies with direct outputs to provider native services.
* Minimal installation, usable as a cli query/investigate tool, or as an operations tool in a compliance as code environment.

As organizations move to cloud infrastructure the need to automate governance becomes paramount to realize the benefits of self-service and on demand infrastructure while ensuring compliance with organizational policies. Different teams within an organization will often use distinct tools for their infrastructure provisioning. However, ensuring that the organization as a whole is well-managed requires additional tooling to ensure that governance objectives are met regardless of the choice of provisioning tool. 

Cloud Custodian exists as a policy/governance as code tool for organizations to enable detective controls to ensure that organizational policies are being adhered to. Policy/governance as code is about bringing some of the best practices of software engineering to policy/governance, ie. using version control, code review, testing, and continuous integration and deployment tools. 

Custodian can be used to manage AWS, Azure, and GCP environments by ensuring real time compliance to security policies (like encryption and access requirements), tag policies, and cost management via garbage collection of unused resources and off-hours resource management. Kubernetes support is alpha at the moment (roadmap).

Custodian's goal is to enable automation around operations, governance, and security in a policy/compliance as code setup. Additional Custodian wants to be broadly usable in many environments so is non-opinionated wrt to operator’s choice of compute environment or devops workflow.

Custodian has a non goal at the moment of producing a separate dashboard. Most organizations have tens to thousands of dashboards, so the general philosophy of custodian is that the best dashboard is the one the org is already using, so it will output structured data, metrics, logs that can be ingested and displayed on an existing dashboard. Additionally it maintains integrations with cloud provider dashboards (ie. AWS SSM Ops Center, Security Hub, GCP Cloud Security Command Center, etc).

Another non goal is IAM/RBAC based preventative management, Custodian generally encourages users to utilize the capabilities of the infrastructure provider’s to implement preventive controls within an environment . However custodian can implement detective controls wrt to IAM/RBAC permissions to govern those definitions and assignments. Ie. check that no compute instance should have a role/service principal that allows it to create a user.

Capital One, Amazon, Microsoft, and Stacklet have guided the Cloud Custodian community and are looking to contribute the project to the CNCF initially as a sandbox project.
 
### Intended Use

Custodian enables reusable vocabularies of filters and actions that can be used to implement many semantic policies. For example an action like `stop` a vm instance could be used for offhours savings, tag compliance, or incident response. Some additional examples might be

* Identify and Remediate load balancers or storage buckets not configured for logging
* Turning off development servers/clusters and databases off at night to realize cost savings.
* Finding underutilized resources and sending an email to their creator to reduce the size. 
* Enforcing tag compliance policies on resources.
* Finding resources with embedded access control policies that are setup to give access across the org boundary.
* In response to a security event, used as a remediation tool to snapshot, disk, network isolate a server, change cloud credentials, and install forensics tools.

#### Target Users

Per the users persona documentation found at https://github.com/cncf/sig-security/blob/master/usecases.md, Cloud Custodian automates several of the use cases defined for the Platform Implementor. By making the role of Platform Implementor more efficient, Cloud Custodian’s aligns towards the goal of making an organization’s Operators (Enterprise, Quota, and Network), Administrators (Security and Compliance/Audit), and Developers successful.

#### Operation

Cloud Custodian is designed to operate as a runtime executable with a Command Line Interface (CLI). The CLI allows for both user interaction and for system interaction where a “system” can include, but not be limited too, a CI/CD pipeline. The CLI takes YAML DSL policy files as input and executes the rules defined in those policy files. The Cloud Custodian CLI has an option for policy validation which checks a policy file for errors. The CLI also has a “dry run” option which can query the cloud provider API for resources which match the policy filters, then save that resource information to an output file for review prior to execution.

It is recommended that users of Cloud Custodian setup a vcs repo for their policies and run a CI system to perform validation and dryrun in addition to code review on changed policies.

In addition to the policy files, the CLI has optional input parameters which can scope the runtime execution to a given cloud region, account configuration, and/or IAM role.

Output of Cloud Custodian can be either verbose or quiet and can optionally include metrics on policy execution. The output data can be directed to cloud provider native logging solutions (CloudWatch logs, Stackdriver, or Azure AppInsights) or cloud provider blob storage. In addition to generating output data, Cloud Custodian policy files can be configured to take action on cloud resources which do not conform.

Cloud Custodian can operate in several different modes. The default mode is called “pull” and is used to execute policies directly from the CLI.  Other execution modes are event based which provision a serverless function of the infrastructure provider, and event subscribers

Custodian also includes a safety-belt feature via the max-resources parameter which will prevent actions being taken if the number of filtered resources exceeds a user defined threshold either by cardinality of percent of total.

### Project Design

A Cloud Custodian policy is designed to act on a resource type, corresponding to an underlying provider service feature. A resource type is defined in the codebase via metadata on the api calls to query the resource and using a plugin pattern to manage the differences in resource types which exist between cloud providers. Cloud Custodian users can then craft policies based on this set of resource types in their policy config files. 

Custodian is stateless by design. During policy execution, the runtime executable will perform queries via cloud provider API calls to get the latest-and-greatest information on resource usage. The implication of this stateless design choice is that the Cloud Custodian rules engine is making determinations on adherence to policy using the provider control plane. This stateless design is true regardless of the operational mode (pull, periodic or event based) chosen by the Custodian user.  

Custodian policies are written in simple YAML configuration files that enable users to specify policies on a resource type (EC2, Redshift, CosmosDB, PubSub Topic) and are constructed from a vocabulary of filters and actions. Custodian  integrates with the cloud native serverless capabilities of each provider to provide for real time enforcement of policies with builtin provisioning. Or it can be run as a simple cron job on a server to execute against large existing fleets.

Each policy has a set of generic filters and resource type specific filters that can be combined in various forms (nested arrays of and/or/not), to filter resources to the desired set, which are then processed by a list of actions. The intent is to enable reuse of filters and actions across many different semantic policies.  
 
### Project Compliance

Custodian does not ship/distribute any policies that correspond to standards/control checklists. Numerous examples of various policies exist within the documentation and across third party sample repositories that correspond to various controls.

### Ecosystem
Custodian is primarily used / integrated directly with the cloud provider native ecosystem tooling it’s deployed against, for logs, metrics, and blob storage. It's intended to operate independently of any chosen tooling, giving operators a mechanism to enforce governance regardless of technologies in use by a team.

### Roadmap
The project roadmap is broken out into the roadmap for the core, as well individual infrastructure providers using github projects.
https://github.com/cloud-custodian/cloud-custodian/projects
 
### Appendix

* We have a separate document with more exhaustive threat modeling exercises contributed by Capital One https://docs.google.com/document/d/1S9zQZaT6G1TA3IAx6YNL0f7G938xaFZ-bziszhuxMZg/edit?usp=sharing
 
* Known Issues Over Time - No known security issues have been reported, numerous correctness and safety (not security) issues have been corrected in code review.
 
* CII Best Practices. https://bestpractices.coreinfrastructure.org/en/projects/3402
 
* Case Studies. 
** https://stelligent.com/2017/05/15/cloud-custodian-cleans-up-your-cloud-clutter/
** https://www.cloudsecops.com/aws-security-audit-using-cloud-custodian-for-aws/
** https://aws.amazon.com/blogs/opensource/announcing-cloud-custodian-integration-aws-security-hub/


## Preferred maturity level

Sandbox

## License

Apache 2:
https://github.com/cloud-custodian/cloud-custodian/blob/master/LICENSE.txt

## Source control

The Cloud Custodian project is composed of multiple source code repositories, all of which are hosted on
Github under the `cloud-custodian` organization:
[https://github.com/cloud-custodian](https://github.com/cloud-custodian)

The core repository can be found at
[https://github.com/cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian). 


## Initial committers (how long working on project, companies they represent)

This list of initial committers represents maintainers of repositories within the Cloud Custodian
organization that have write permissions to the respective repositories that they maintain.

| Name | GitHub | First Contribution | Organization |
| ---- | ------ | ------------------ | ------------ |
| Kapil Thangavelu | kapilt | Stacklet |


## Communication channels (slack, irc, mailing lists)

The official communication channel for the project can be found at https://gitter.im/cloud-custodian/cloud-custodian

## Issue tracker

All issues are found in the appropriate repo in our GitHub organization: 
[https://github.com/cloud-custodian](https://github.com/cloud-custodian).

## Website

The official public facing website can be found at [https://cloudcustodian.io/](https://cloudcustodian.io/)


## Community size and any existing sponsorship

The following statistics, when related to code repositories, are for the main Cloud Custodian repo of
[https://github.com/cloud-custodian/cloud-custodian](https://github.com/cloud-custodian/cloud-custodian).

* GitHub Stars: 2,940+
* Releases: 77+
* Commits: 3,292+
* Forks: 856+
* Contributors: 258+ 
* Maintainers: 21 maintainers from 4 organizations


### Who is currently known to be using the project? Are they using it in production and at what scale?

#### Capital One

#### Microsoft

#### Amazon (AWS)


## Project logo in svg format

The Cloud Custodian logo in SVG can be found at
[https://cloudcustodian.io/img/logo_capone_devex_cloud_custodian.svg](https://cloudcustodian.io/img/logo_capone_devex_cloud_custodian.svg)

All graphics and artwork can be found in the primary Cloud Custodian repo:
[https://github.com/cloud-custodian/cloud-custodian/tree/master/docs/logos](https://github.com/cloud-custodian/cloud-custodian/tree/master/docs/logos)
