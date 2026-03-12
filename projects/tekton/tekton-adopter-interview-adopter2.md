# Tekton Adopter Interview - Adopter 2

Meeting Time: October 20th, 2025.
Meeting Record: TBA

Attendees:
	TOC: 
  - Jeremy Rickard
	Adopter: 
  - DevOps/MLOps Lead, Adopter 2


## Organization Intro
_Can you give us an overview of your organization and what it does?_

The adopter is in the vacation and leisure industry. 


## Motivation

_Compared with other products in this space (proprietary and open), what drew you to the project?_

The adopter was a heavy user of Jenkins and had mature devops practices, but was experiencing issues with scaling as their IT organization was growing. The adopter wanted a Kubernetes native/cloud native CI/CD solution that would provide the ability to scale out and in over time. They wanted to adopt a an open source technology that would allow them to provide declarative pipelines as code in order to shift to a self-service onboarding model.

## Usage Scenario

_How long has your organization used the project?_

The adopter began researching Tekton in 2019 and then during COVID used began to migrate from Jenkins to Tekton in 2020. The adopter is now using Tekton for 100% of CI/CD. 

_What were the main motivations to adopt the project and which key features do you use today?_ 

Essentially use Tekton core features like Tasks and Pipelines. 50 highly reusable Tekton tasks and 45 flavors of CI/CD pipelines. 

Tekton allows creation of CI/CD pipelines with kubectl. This allowed further integration with Backstage as an IdP. This has really streamlined the onboarding time for the adopter, with typical onboarding in an hour and a half for new projects. 

_What is the current level of usage (pre-production, production) and scale?_

Used in development and production, Tekton is a mission critical platform. Average of 70,000 deployments a year. 

_What version is used and what is your update cadence with the project?_

The adopter is using 0.59, which is an older version. The new features were not critical, so they have stuck with the older version.


_Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?_ 

The adopter did not find any significant technical challenges. The main challenge was to shift from a Jenkins UI model to a true gitops model. This introduced some friction and they encountered some challenges with developers learning the new model.


_Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?_

The adopter had to occasionally look to the Github repository for additional information, but generally found the project documentation really good. No complaints with the documentation. 


_Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc._

The adopter indicated that Tekton has provided measurable value, particularly in the ability for teams to self service their own pipelines and a large reduction in support activities that their devops organization needed to engage in. The ability for developers to run their pipelines from within their IDE accelerated the development lifecycle. Tekton also accelerated their adoption of Backstage significantly.

_Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc._

The adopter plans to start extending Tekton to help manage ML pipelines, particularly ML Flow to build an MLOps practice and standardize workflows across different tools. 


## Perception
_What is your perception in terms of the project’s:
Community openness
Governance
Community growth potential
Maintainer diversity and ladder
Maintainer response_

The adopter perceives the Tekton community to be open and very welcoming. The project has invited the adopter to contribute to the project. 

The adopter was not familiar with the governance of the project.

The adopter feels that the maintainers are very responsive compared to other open source communities they have interacted with.

The adopter did not have a clear picture of the maintainer diversity, but felt that the project is very welcoming.

_How are you participating in the project community?_

The adopter mainly participates as a consumer and adopter of the project, but doesn't actively participate in the community. 


_Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?_

The adopter had to occasionally interact with the community and maintainers via Slack when documentation did not cover everything. The adopter found the community and maintainers extremely responsive via Slack. 


## Project Strengths
_In your opinion, what are the overall strengths of the project?_

The integration with Kubernetes is a primary strength. The ability to run the pipelines within Kubernetes on vetted infrastructure simplifies their security model and helps better align with compliance initiatives.  


## Project Improvements
_Is there something you feel that holds the project back from reaching its ultimate potential?_ 

The adopter feels that there could be more awareness of the project and potentially lacks marketing. The adopter thinks that there could be more positioning within the CI/CD space and is optimistic that the project could grow its community after joining the CNCF.


_In your opinion, what can the project do better?_

The adopter indicated that a wider emphasis on a UI and a plugin ecosystem could make the project even more applicable to the space. 





