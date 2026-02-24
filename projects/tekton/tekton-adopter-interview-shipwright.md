# Tekton Adopter Interview - Shipwright Project

Meeting Time: October 17th


Attendees:
	TOC: 
  - Jeremy Rickard
	Adopter: 
	- Enrique Encalada, Software Developer (IBM) / Shipwright Maintainer
	- Sascha Schwarze, Software Developer (IBM) / Shipwright Maintainer


## Organization Intro
_Can you give us an overview of your organization and what it does?_

Shipwright is a CNCF Sandbox project. It is a platform for building container images directly on Kubernetes, with the goal of building the container images where they will be used. Shipwright uses a few custom resources to define how to build a container (the tools), what to build (source code), and where it should be pushed to. 


## Motivation

_Compared with other products in this space (proprietary and open), what drew you to the project?_

Tekton was a good fit for Shipwright when the project began. It was chosen because it provided needed capabilities and ran natively in Kubernetes; Tekton was already relevant for Kubernetes-based CI/CD.

## Usage Scenario

_How long has your organization used the project?_

Tekton has been used since the inception of the project in 2020. 

_What were the main motivations to adopt the project and which key features do you use today?_ 

Shipwright needs to execute a series of steps sequentially for builds. The Tekton `TaskRun` executes `Steps` in order, and `Parameters` pass user input to builds. `TaskRun`, `Task`, `Step`, and `Parameters` have been used since the project's start. Tekton `Results` were adopted later to report information in Shipwright build status, such as the Git SHA and vulnerability reports. The Shipwright Trigger capability, which is fairly new, also uses Tekton Custom Tasks.



_What is the current level of usage (pre-production, production) and scale?_

Shipwright is an open source project, so it is hard to quantify scale directly. Shipwright is used in production by IBM Cloud Code Engine and Red Hat OpenShift, which means that Tekton is also being used in production by those vendors.


_What version is used and what is your update cadence with the project?_

Shipwright generally supports Tekton LTS releases and follows Tekton's release cadence. Shipwright updates and tests against both the oldest and newest Tekton releases when they occur.


_Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?_ 

The Shipwright project followed Tekton's evolution from alpha to beta to v1 custom resources and generally did not encounter many difficult upgrades. There were some API changes between versions over the five years of adoption, but they were minimal. The API (Kubernetes resources) was well designed from the beginning. As a platform on top of Tekton, Shipwright found it easy to handle multiple Tekton versions.


_Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?_

The website documentation isn't perfect: core features were well documented initially, but as new features were added the documentation structure could be better organized. Tekton Enhancement Proposals were useful for understanding the rationale behind new features. The Tekton Catalog also provided good examples of how to use the API.


_Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc._

When Shipwright started, adopting Tekton saved several person-months of development by providing core capabilities the project would otherwise have had to rebuild.



_Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc._

Shipwright is still working on the Shipwright Trigger feature, which will build on Tekton custom task capabilities and improve integration. Shipwright is also exploring how Tekton can use Kubernetes' assign pod-level resources feature to better support resource assignment in Shipwright.


## Perception
_What is your perception in terms of the project’s:
Community openness
Governance
Community growth potential
Maintainer diversity and ladder
Maintainer response_

The Shipwright maintainers feel that the Tekton community is very open and welcoming. While Shipwright was also in the Continuous Delivery Foundation (CDF), they felt welcome at Tekton Community Days and other events.

The Shipwright maintainers have the impression that governance is well defined and shows a good level of diversity across subprojects; they feel it does not favor any one company or maintainer. The Shipwright maintainers don't regularly attend the Tekton community calls.

The Shipwright maintainers feel that the core Tekton feature set might be described as "finished." In recent years, they perceive Tekton pipelines have stabilized while several other Tekton projects continue to grow.

The Shipwright maintainers feel there is a progression path for contributors to grow, and mechanisms to transition contributors off the project to make room for new people.

The Shipwright maintainers felt that pull requests were reviewed in a timely manner and questions asked in Slack were always answered. 


_How are you participating in the project community?_

The Shipwright maintainers have interacted with Tekton community members via Slack and GitHub, but have not regularly participated in Tekton community meetings. Tekton community members held Shipwright community meetings to share new features like Tekton Chains.


_Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?_

The Shipwright maintainers have interacted with Tekton community members via Slack and GitHub, generally to ask questions and open pull requests for changes needed by Shipwright, such as bugs they encountered.


## Project Strengths
_In your opinion, what are the overall strengths of the project?_

The Shipwright maintainers identified the maturity and stability of the Tekton API as one of its major strengths. There have not been any major issues upgrading versions of Tekton in at least two years.


## Project Improvements
_Is there something you feel that holds the project back from reaching its ultimate potential?_ 

From a Shipwright perspective, Tekton does everything that Shipwright needs. 


_In your opinion, what can the project do better?_

The project could provide better integration with `kubectl explain`, which would help users learn about custom resources without constantly referring to the website documentation. The website documentation structure could be improved, especially for new adopters of Tekton.



