# CNCF TOC Principles

_Version 1.0, Nov 27, 2017  
Approved by TOC on: Nov 27, 2017
Approved by GB on: Dec 5, 2017  
Last reviewed and concurred: Jul 27, 2022, [Issue 413](https://github.com/cncf/toc/issues/413)

[TOC Operating Principles](#toc-operating-principles)  
* [We Are Project-Centric](#we-are-project-centric)  
* [Projects Are Self-Governing](#projects-are-self-governing)  
* [What We're Looking For](#what-were-looking-for)   
* [No Kingmakers & One Size Does Not Fit All](#no-kingmakers--one-size-does-not-fit-all)  
* [Not a Standards Body](#not-a-standards-body)  
* [We Want a Comprehensive Toolchain](#we-want-a-comprehensive-toolchain)    
* [Above All We Want To Help Projects](#above-all-we-want-to-help-projects)


[Technical Leadership Principles](#technical-leadership-principles)    
* [Purpose](#purpose)    
* [Principles](#principles)
  * [Be welcoming](#be-welcoming)    
  * [Provide feedback and appreciation appropriately](#provide-feedback-and-appreciation-appropriately)
  * [Always do better](#always-do-better)    
  * [Lead inclusively and empower all voices and opinions to be heard](#lead-inclusively-and-empower-all-voices-and-opinions-to-be-heard)
  * [Foster respectful resolution](#foster-respectful-resolution)    
  * [Facilitate open, transparent, and inclusive decision-making](#facilitate-open-transparent-and-inclusive-decision-making)

## TOC Operating Principles

As part of our continued community engagement, the following list of principles serve to provide an understanding of how the TOC intends to work with projects, effectively serving as the guiding star for the TOC. It also provides insight into what the TOC is looking for in terms of continuity and evolution over time. Since the CNCF's inception we have endeavoured to capture the lessons learned and iterate on these principles to serve the foundation, the community, and the projects.

 1. We are project-centric.
 1. We rely on minimal viable governance to enable community driven success.
 1. Do no harm by enabling robust, well engineered projects and supportive processes to increase the likelihood of success.
 1. No kingmakers. Similar or competitive projects are not excluded for reasons of overlap. 
 1. No single stack. Encourage interoperability for the emergence of a variety of stacks and patterns to serve the community and adopters.
 1. Promote interfaces and defacto implementations over standards for real world use.
 1. Users should not need to look beyond the CNCF for cloud native app tooling.
 1. To help facilitate high-quality high-velocity cloud native open source projects which are the main driver of customer adoption and success.


### We Are Project-Centric

_Principle: If it can be on a modern public source code control system, then it can be a project. And we put projects front and center._

CNCF is a home for several kinds of “projects” where community collaboration furthers the goals of the CNCF community:
1. Open source software projects, e.g., Prometheus.
1. Projects that develop interface and/or schema specifications (e.g., [CNI](https://github.com/containernetworking/cni), reference implementations, conformance tests, adaptors, etc., in order to facilitate interoperability.
1. Reference materials, such as architectures, stacks, guides, docs.
 
## Projects Are Self-Governing

_Principle: We exercise minimal viable governance to enable community driven success._

Our expectations around governance and support are all predicated on the notion that a CNCF project works like a typical, modern “community-owned” open source software project, such as a person might discover hosted on GitHub. That means that it has ‘committers’ and shared ownership using source code control, etc. People who want the CNCF to support their ‘thing’ need to make it into a project and support “GitHub-style” communities. CNCF projects are not required to reside on GitHub to support a GitHub-style community.

In the GitHub era, open projects are able to get a lot “done” without outside help. The CNCF does not want to get in the way of that. This starts with “minimal viable governance”.

- The CNCF, TOC, and GB are available for help if it is asked for.
- But: we do not want to impose bureaucracy on projects because that will slow them down.
- Minimal viable governance also means that the TOC does not step in at a tactical level to overrule project leads’ decisions.
- There are some basics like [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) - see Draft Statement below. Including dealing with problematic leads & maintainers.
- There is a formal & regulated system of [Graduation Criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md) for CNCF Projects
- The TOC/CNCF want the ability to intervene if things go really wrong - i.e., project leads are stuck and cannot fix things.
- Provide a template for new projects, a set of best practices to help jump-start the task of setting up a new project.

### Public Statement

The CNCF is committed to helping its member projects succeed, but without dictating or micromanaging how the projects are run. To that end, it requires only minimal viable governance criteria: a [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) with neutral processes for resolving conflicts, a documented governance model that includes a contribution-based process by which contributors can become committers or maintainers, and a clear definition of the top-level project leadership, with which the foundation will engage and inform and from which it may receive requests for funding and support. Once a project has graduated from incubation, new governance requirements cannot be imposed without consent of the project, except where legally required.

## What We're Looking For

_Principle: Do no harm, by enabling robust, well engineered projects and supportive processes to be leveraged for greatest adoption likelihood._

We are looking for high-quality, high-velocity projects that fit ‘cloud native’.

Identify projects that have a real shot at being a useful tool in the evolving box of cloud native technology. This is a mix of mature and early-stage projects.   Early stage may not have all the criteria we want: diverse contributor bases, formalized governance, interoperability, cloud native designs, quality bar, etc. 

Some considerations:
- Transparent, consistent technical and governance quality bar for [graduation](/process/graduation_criteria.adoc) from incubation
- Has users, preferably in production; is a high quality, high-velocity project (for incubation and graduated projects). Sandbox level projects are targeted at earlier-stage projects to cultivate a community/technology
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
1. No kingmakers. Similar or competitive projects are not excluded for reasons of overlap. 
1. No single stack. Encourage interoperability for the emergence of a variety of stacks and patterns to serve the community and adopters.

The TOC picks projects with a real chance of achieving widespread use, and it does not pick a single winner in each category. Via the “no kingmakers” principle and “what is a project”, the CNCF may support several projects which show how a ‘stack’ is a solution to certain use cases. For example, some stacks might use a container orchestrator. Other stacks might show how to integrate monitoring with other “observability” technologies, for cloud native apps regardless of their orchestration model.

The TOC does not pick a “winning stack” - i.e., vertically integrated set of projects as a solution for multiple application problems. This means that the CNCF is not promoting a single, monolithic stack of technologies.

CNCF projects may have dependencies on other open source projects that comply with the [CNCF IP policy](https://github.com/cncf/foundation/blob/master/allowed-third-party-license-policy.md), which may be outside the CNCF. Widely-used dependencies are preferred; pluggable dependencies fit well with our principles but are not essential.  _Examples: Kubernetes' dependency on etcd, which was not a CNCF project at the time Kubernetes was brought into the CNCF; CNI, which enables pluggable networking dependencies._ 

The CNCF is a badge of quality and velocity. CNCF projects should be on a path to being tools that users can trust, broadly work together, and that meet other cloud native criteria. But the CNCF badge does not mean “this is the standard tool”.

- Overlapping projects are ok, especially where they make significantly different design tradeoffs
- Who to pick:  The market is too young for us to pick winners - sometimes we shall identify several really promising tools that overlap in function. Let's aim for eventual consistency based on real community use, and not create early deadlocks over "which one tool is best".  
- CNCF resources, both time and money, to be extremely limited, so we do need to choose carefully, and therefore should do some reasonable due diligence, including considering alternative projects. 
- Overall the TOC will try to maintain a public roadmap or “backlog” where it sees interesting projects emerging, or space for “RFPs”, and “WGs”. While not a hard and fast document, this will help make dialogue with the community more transparent & efficient.

## Not a Standards Body

_Principle: Promote interfaces and defacto implementations over standards for real world use._

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

CNCF may develop written materials in the style of the current CNI interface document, or in the style of an IETF RFC for example. These CNCF “specification” materials are not “standards”. It is possible in the future that an independent and recognized international standards body takes a CNCF document as “upstream” and evolves it into a standard (e.g. via the IETF process). The CNCF is morally supportive of independent parties doing this, but does not see this work as its own responsibility. For that matter, to date, no such exercise is in process.

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

_Principle: Users shouldn't need to look beyond the CNCF for cloud native app tooling_

Grand vision: CNCF should identify, facilitate and promote a *complete toolset* for cloud native applications and stacks at scales from small to large. This enables customers to adopt good tools faster, and be less at risk of confusion and doubt.

We need to flesh out the portfolio of projects needed by users to succeed with cloud native computing. It's pretty clear where remaining gaps in the project portfolio are. We should try to fill some of those, and at least document the rest. We can make more WGs to help with a few of those. (Note that “detailed architecture and stack” can be a project)

## Above All We Want To Help Projects

_Principle: We help high-quality high-velocity cloud native open source projects which are the main driver of customer adoption and success._
 
We want to be able to say that CNCF is a net positive for big & small projects. Doing so requires more coordination with project leads. 

Project needs may include test automation and CI, cloud resources to test on, clear documentation, per-project marketing & evangelism, roadmaps for interop, and advice from experts on governance and scalability. And we need to make sure project contributors see what value they are getting & are not afraid to ask for help!


## Technical Leadership Principles

### Purpose

The Technical Oversight Committee has adopted these leadership principles to help guide the ecosystem in growth and maturity. These principles define the expectations of TOC members, project maintainers, TAG Chairs and Leads, chairs and leads of any group or committee, and other persons serving in any leadership role in the CNCF community.

#### Guidance and Enforcement

For additional guidance and information on how the Technical Leadership Principles are enforced and how projects can align with the principles, please refer to the [Technical Leadership Principles Guidance in the TOC's docs folder](docs/tech_leadership_principles_guidance.md).

### Principles

Our goal with these is to foster a more inclusive and welcoming community for all, and we expect all leaders in the CNCF community  to participate in accordance with these principles. If a leader is not acting in accordance with these principles, community members should first address it with the other leaders of the project, TAG, committee, group, or body. Project and TAG governance documentation must explain procedures (including who is responsible) for response and resolution if these principles are not being followed. In general, if a leader is not following these principles, there should be a good faith effort to educate, inform, and provide the leader with an opportunity to improve before taking further action such as suspending or removing them from their leadership role. If the behavior that violates these principles persists, or if the other leaders of the applicable project, committee, or group do not take corrective action in response to a reported violation of these principles, projects should escalate the matter to the TOC.

In addition to adhering to these principles, all leaders must adhere to the CNCF Code of Conduct. 

These leadership principles are in addition to the TOC Operating Principles, shown above. 

#### Be welcoming

Make a conscious effort to be welcoming in your interactions. If someone is new, express sincere interest in what brought them to the group. Check in with existing contributors to ensure they continue to feel welcome. Create space for new contributors and existing contributors to be curious and engaged with the work that we do in cloud native by building on common interests and outcomes. 

#### Provide feedback and appreciation appropriately
Provide praise, gratitude, appreciation and encouragement in public. Look for opportunities to say “thank you” and recognize people when they are being helpful or making contributions. Be specific in what was done and describe the impact it had or continues to have.

If you have negative or constructive feedback about an individual or group, provide it in a form that is meaningful and actionable in private, ideally with support from a neutral third-party. Having another individual present that serves as a neutral participant can assist in removing biased power dynamics or asymmetry in the discussion. Constructive feedback should be specific, with a recommended alternative, solution, or suggestion for improvement. If possible, make an effort to find out how the individual or group would like to receive feedback; not everyone is comfortable receiving feedback in person or over a virtual meeting, and ensure they are accepting of a neutral participant to support an engaging feedback conversation while providing a transparent and equitable process. 

Be gracious, courteous, and kind in how you provide feedback and how you receive feedback from others.

#### Always do better

Identify and give opportunities to improve while considering how feedback might be perceived by someone receiving it. Request from others and self-identify areas in your own interactions that could be improved and take steps to make those changes. Make a conscious effort to effect change in your own behaviors and be the model of that effort for others to strive towards.

#### Lead inclusively and empower all voices and opinions to be heard
Actively seek to engage a variety of different voices and opinions as our community is built off a wide group of perspectives, experiences, and backgrounds. We are better as a community when we engage with all of the breadth of cloud native to better understand one another. As leaders, we are responsible for enabling everyone to not only have the opportunity to not only be engaged, but to bring balance to the discussions so no single voice - or perspective - overruns the desires of the collective whole. 

For example, make an extra effort to ask the person who is new or who hasn’t yet spoken up in a meeting if they have any comments to offer. Make conscious effort to facilitate discussions in a manner that allows all voices to be heard and included, not just the loudest or most opinionated voices. Be aware that people who are shy, introverted, from historically underrepresented groups, and from certain cultures may be hesitant to speak up unless you clearly invite them to offer their opinion. Thus, make that extra effort to include people. Be sure to give everyone who wants to express their perspective an opportunity to offer their opinions even if they are not present in meetings. Invite people to continue sharing feedback asynchronously after meetings via Github, slack, or another channel that all participants have access to so they have the opportunity to provide comment.

#### Foster respectful resolution

Handle disagreements respectfully. Seek to resolve conflict and disagreement in a respectful and collaborative manner and encourage others to do the same. Show empathy, care, and concern for others’ feelings, perspectives, and experiences while working towards a resolution that considers the Community, the projects we interact with, the mission of our work, and what can reasonably be accomplished with consensus in accordance with the project's governance and scope. Conflict and disagreement will happen, as leaders we must de-escalate, be the agent of calm, and provide neutrality for all parties to feel welcome in expressing their perspectives while pursuing an acceptable outcome. 

#### Facilitate open, transparent, and inclusive decision-making

Although it is perfectly acceptable to discuss ideas, proposals, and agendas with other participants in private conversations, decisions of any official community group, committee, project, or governing body should be made in an open and transparent manner, in accordance with that body’s governance rules, and the reasoning or basis for such decision should be clearly documented. This not only allows others to understand the why, but also may serve as the basis of future discussions and activities that resulted from that decision and provide the opportunity to change it should circumstances change.
