# CNCF TOC Principles

_Version 1.0, Nov 27, 2017  
Approved by TOC on: Nov 27, 2017
Approved by GB on: Dec 5, 2017  

[TOC Operating Principles](#toc-operating-principles)  
[We Are Project-Centric](#we-are-project-centric)  
[Projects Are Self-Governing](#projects-are-self-governing)  
[What We're Looking For](#what-were-looking-for)   
[No Kingmakers & One Size Does Not Fit All](#no-kingmakers--one-size-does-not-fit-all)  
[Not a Standards Body](#not-a-standards-body)  
[We Want a Comprehensive Toolchain](#we-want-a-comprehensive-toolchain)    
[Above All We Want To Help Projects](#above-all-we-want-to-help-projects)  

## TOC Operating Principles

Now that CNCF has been active for over a year we want to start writing down what we have learned. Future TOCs can still make changes, but there will at least be documented precedent.

### We Are Project-Centric

_Principle: If it can be on a modern public source code control system, then it can be a project. And we put projects front and center._

CNCF is a home for several kinds of “project” where community collaboration furthers the goals of the CNCF community:
1. Open source software projects, e.g., Prometheus.
1. Projects that develop interface and/or schema specifications (e.g., [CNI](https://github.com/containernetworking/cni), reference implementations, conformance tests, adaptors, etc., in order to facilitate interoperability.
1. Reference materials, such as architectures, stacks, guides, docs.
 
## Projects Are Self-Governing

_Principle: Minimal Viable Governance_

Our expectations around governance and support are all predicated on the notion that a CNCF project works like a typical, modern “community-owned” open source software project, such as a person might discover hosted on GitHub. That means that it has ‘committers’ and shared ownership using source code control, etc. People who want the CNCF to support their ‘thing’ need to make it into a project and support “GitHub-style” communities. (Though please note that CNCF projects don't actually need to live on GitHub.)

In the GitHub era, open projects are able to get a lot “done” without outside help. The CNCF does not want to get in the way of that. This starts with “minimal viable governance”.

- The CNCF, TOC, and GB are available for help if it is asked for.
- But: we do not want to impose bureaucracy on projects because that will slow them down.
- Minimal viable governance also means that the TOC does not step in at a tactical level to overrule project leads’ decisions.
- There are some basics like [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) - see Draft Statement below. Including dealing with problematic leads & maintainers.
- There is a formal & regulated system of [Graduation Criteria](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc) for CNCF Projects
- The TOC/CNCF want the ability to intervene if things go really wrong - i.e., project leads are stuck and cannot fix things.
- Provide a template for new projects, a set of best practices to help jump-start the task of setting up a new project.

### Draft Public Statement for website

The CNCF is committed to helping its member projects succeed, but without dictating or micromanaging how the projects are run. To that end, it requires only minimal viable governance criteria: a [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) with neutral processes for resolving conflicts, a documented governance model that includes a contribution-based process by which contributors can become committers or maintainers, and a clear definition of the top-level project leadership, with which the foundation will engage and inform and from which it may receive requests for funding and support. Once a project has graduated from incubation, new governance requirements cannot be imposed without consent of the project, except where legally required.

## What We're Looking For

We are looking for high-quality, high-velocity projects that fit ‘cloud native’.

_Principle: Great projects already have many ingredients to succeed. First: do no harm._

Identify projects that have a real shot at being a useful tool in the evolving box of cloud native technology. This is a mix of mature and early-stage projects.   Early stage may not have all the criteria we want: diverse contributor bases, formalized governance, interoperability, cloud-native designs, quality bar, etc. 

Some considerations:
- Transparent, consistent technical and governance quality bar for [graduation](https://www.cncf.io/projects/graduation-criteria/) from incubation
- Has users, preferably in production; is a high quality, high-velocity project (for incubation and graduated projects). Inception level projects are targeted at earlier-stage projects to cultivate a community/technology
- Has a committed and excited team that appears to understand the challenges ahead and wishes to meet them
- Has a fundamentally sound design without obvious critical compromises that will inhibit potential widespread adoption
- Is useful for cloud native deployments & ideally, is architected in a cloud native style
- Has an affinity for how CNCF wants to operate
- Charter Section 9(e): _New open source projects initiated in CNCF shall complete a project proposal template adopted by the TOC and be approved by the TOC for inclusion in CNCF. The TOC members shall be afforded sufficient time to discuss and review new project proposals. New project proposals shall include details of the roles in the project, the governance proposed for the project and identify alignment with CNCF’s role and values_

## No Kingmakers & One Size Does Not Fit All

_Preamble:_
1. Many problems in technology have more than one solution. Markets are good at finding them. There may be multiple good OSS projects for solving a problem that are optimized for different edge cases.
1. Often multiple solutions get widespread use, for example, because they are optimized for different constraints. We don’t want to get in the way of that by insisting that one technology is “the answer” for each functional gap that we can identify today. We believe that the market and user community provide a good mechanism for pushing the most appropriate projects forward over time. We want projects to enjoy the support of the CNCF during that process.
1. There is no “one true stack”, cloud native applications cover many different use cases with different needs. Many architectures are reasonable: from 12-factor to microservice, to stateful or data-intensive, to others. And there are many scales from one node to many, from low to high latency, etc. So “one size does not fit all”.

_Principles:_
1. No kingmakers. The TOC picks projects with a real chance of achieving widespread use, and it does not pick a winner in each category. Similar or competitive projects are not excluded for reasons of overlap. 
1. No one stack. The TOC does not pick a “winning stack” - i.e., vertically integrated set of projects as a solution for multiple application problems. Instead by encouraging interop, we hope that a range of patterns & “stacks” will emerge.

Via the “no kingmakers” principle and “what is a project”, the CNCF may support several projects which show how a ‘stack’ is a solution to certain use cases. For example, some stacks might use a container orchestrator. Other stacks might show how to integrate monitoring with other “observability” technologies, for cloud native apps regardless of their orchestration model.

This means that the CNCF is not promoting a single, monolithic stack of technologies. 

CNCF projects may have dependencies on other open source projects that comply with the [CNCF IP policy]( https://github.com/cncf/foundation/blob/master/whitelist-policy.md), which may be outside the CNCF. Widely-used dependencies are preferred; pluggable dependencies fit well with our principles but are not essential.  _Examples: Kubernetes' dependency on etcd, which was not a CNCF project at the time Kubernetes was brought into the CNCF; CNI, which enables pluggable networking dependencies._ 

The CNCF is a badge of quality and velocity. CNCF projects should be on a path to being tools that users can trust, broadly work together, and that meet other cloud native criteria. But the CNCF badge does not mean “this is the standard tool”.

- Overlapping projects are ok, especially where they make significantly different design tradeoffs
- Who to pick:  The market is too young for us to pick winners - sometimes we shall identify several really promising tools that overlap in function. Let's aim for eventual consistency based on real community use, and not create early deadlocks over "which one tool is best".  
- CNCF resources, both time and money, to be extremely limited, so we do need to choose carefully, and therefore should do some reasonable due diligence, including considering alternative projects. 
- Overall the TOC will try to maintain a public roadmap or “backlog” where it sees interesting projects emerging, or space for “RFPs”, and “WGs”. While not a hard and fast document, this will help make dialogue with the community more transparent & efficient.

## Not a Standards Body

_Principle: CNCF promotes interoperability via interfaces that get real-world use_

Users and vendors both want as little friction as possible when it comes to integration. Taking any two projects and putting them into a larger stack, product, platform, solution, etc., is always easier when this is the case. For example, cloud native storage vendors want as few moving parts as possible when it comes to making their products and services work with the various emerging container platforms. At the same time, the world is littered with the untouchable remains of failed ‘standards’ that promised to provide interoperability but did not get traction.

In the CNCF we like projects that are getting traction and may go on to become widespread and popular with end users and the ecosystem. We apply this thinking to the area normally covered by ‘standards’ including specifications for APIs, SPIs, protocols, etc. Where a good interface exists we are happy to use it if our users like it, but we are not compelled to do so. 

We want markets and users to drive interop, not committees. We want to help real-world use happen faster, and foster collaboration. We do not wish to become gated on committees.

We should focus on areas of rough 'de facto' agreement, under the proviso that early markets are also diverse. Possible areas to target:
- CNI - network
- CSI - storage
- CRI - runtime 
- Openmetrics
- CLI - logging
- Spiffe
- (there is demand for some kind of serverless-related event scheme)

### How CNCF works with interface definitions like CNI, vs. standards efforts like OCI

The world has a number of recognized international standards bodies such as IETF and W3C. CNCF is not playing the role of a standards body. By contrast, OCI is a standards body.

CNCF may develop written materials in the style of the current CNI interface document, or in the style of an IETF RFC for example. These CNCF “specification” materials are not “standards”. It is in future possible that an independent and recognized international standards body takes a CNCF document as “upstream” and evolves it into a standard via (e.g.) the IETF process. The CNCF is morally supportive of independent parties doing this, but does not see this work as its own responsibility. For that matter, to date, no such exercise is in process.

In general CNCF specifications will evolve as “living documents” side by side with the CNCF OSS projects that adopt them.   This means they are “dynamic and fast-moving” whereas a typical IETF standard is ”static” and not subject to frequent updates & new releases.

For the avoidance of doubt: Any written “specification” materials in CNCF shall have the same status as CNCF open source software and project technical docs. That is to say that specifications shall be updated and release versioned according to the conventions of a CNCF open source project.

### Important: Principle of Interoperability

CNCF values this highly: that for any given ‘specification’, multiple implementations exist. Those implementations will use the project’s specification as the source of truth.

Moreover, CNCF shall not claim interoperability if there is only one implementation.

### Example: CNI

CNI fits all of the above requirements. There is 1) a specification and it is co-developed with 2) a library (libcni) and 3) ecosystem-contributed plugins. Collectively (1-3) form “the CNI Project”. In the future, CNI might include a compliance test suite which can be run against those implementations.

CNI is a software project, but the centerpiece of that project is the set of interfaces documented in the specification. Those interfaces live in the CNI Project. By the principle of interoperability, the existence of multiple CNI implementations is encouraged. Those implementations will use the interface definition from CNI as the source of truth. 

### Example: OCI

OCI is not in the CNCF. The OCI project operates according to the norms of international standards bodies and has a clear primary goal: to provide a document that describes a standard set of interfaces for a container runtime and label this as 1.0. Nothing else is required.

## We Want a Comprehensive Toolchain

_Principle: users don’t need to look beyond the CNCF for cloud native app tooling_

Grand vision: CNCF should identify, facilitate and promote a *complete toolset* for cloud native applications and stacks at scales from small to large. This enables customers to adopt good tools faster, and be less at risk of confusion and doubt.

We need to flesh out the portfolio of projects needed by users to succeed with cloud-native computing. It's pretty clear where remaining gaps in the project portfolio are. We should try to fill some of those, and at least document the rest. We can make more WGs to help with a few of those. (Note that “detailed architecture and stack” can be a Project)

## Above All We Want To Help Projects

_Principle: Our top priority is helping high-quality high-velocity cloud native open source projects be the main driver of customer adoption and success_
 
We want to be able to say that CNCF is a net positive for big & small projects. Doing so requires more coordination with project leads. 

Project needs may include test automation and CI, cloud resources to test on, clear documentation, per-project marketing & evangelism, roadmaps for interop, and advice from experts on governance and scalability. And we need to make sure project contributors see what value they are getting & are not afraid to ask for help!
