# Tekton Adopter Interview - Adopter 3

Meeting Time: December 18th
Meeting Record: TBA

Attendees:
	TOC: 
  - Jeremy Rickard
  - Chad Beaudin
	Adopter: 
  - Developer Experience Engineer, Adopter 3
  - Developer Experience Engineer, Adopter 3
  - Developer Experience Team Manager, Adopter 3


## Organization Intro
_Can you give us an overview of your organization and what it does?_

Adopter 3 is a video streaming and advertising platform.


## Motivation

_Compared with other products in this space (proprietary and open), what drew you to the project?_

Tekton was an open source project and the adopters could make changes and contribute if needed. The project was relatively new when the adopters selected it and it integrated directly into Kubernetes. 

Tekton was actually the team's first choice, but embedded in JenkinsX which the adopter was using. 

The adopters migrated to using Tekton directly over time. The adopters were motivated to adopt the project directly because of the ability to have shared pipeline definitions and common intelligence within a single repository and reuse this across multiple projects. This would allow them to make changes in a single place and impact all the projects resuing the pipeline logic.


## Usage Scenario

_How long has your organization used the project?_

The project has been in use for 3-5 years. Initially the project was used in JenkinsX but has been used natively without any JenkinsX features for the last two years.


_What were the main motivations to adopt the project and which key features do you use today?_ 

The adopter is mainly using the Tekton Pipeline subproject, and is heavily using the  `StepAction` and the `pipelineRef` features. This allows them to generically reuse things from their internal catalog. They are not using the Tekton Catalog directly. They are also using the `Matrix` feature in order to increase the parallelism of their CI/CD.


_What is the current level of usage (pre-production, production) and scale?_

Tekton is used in production by the adopter, as well as in staging environments. Tekton handles all pull requests within Adopter 3 and creates 40,000 pods a month in CI/CD pipelines.


_What version is used and what is your update cadence with the project?_

The adopter upgrades one or two times per month, but is generally running one version behind the latest. They are generally confident in adopting new versions and can test before rolling out. 


_Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?_ 

The project initially adopted Tekton via JenkinsX, but JenkinsX hid most of the Tekton capability. They noticed that Tekton was moving faster and had more maintainers and proposed moving to the native Tekton references. The adoption was a challenge because they had to rebuild all the pipelines to use Tekton directly, but moving to Tekton allowed adoption of Tekton features more quickly.

The initial adoption probably took six months, but allowed the organization to consolidate into a shared pipeline through the use of the `pipelineRef` feature in Tekton. They were able to move from updating 100s of pipeline definitions to a single pull request when they needed to make changes.

The main challenge when doing the migration was that they were initially using an older version of Tekton and the documentation was not easy to find anymore at that point. Upgrading to the newer versions was also a challenge because the API versions changed from a beta API to stable API.

_Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?_

The Tekton documentation was very useful. Initially, the adopter had to locate the documentation in the Github repository within the tag for the specific version they were using, as it was not present on the website anymore.


_Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc._

The adopter sees a measurable value through reduced manual activities and toil when making changes and ability to have shared pipeline logic. They are also collecting metrics from pipelines and see improvements in all pipelines as they make changes to the shared catalog and find this beneficial for improving their overall workflow. 


_Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc._

In 2026, the adopter plans to implement a previously submitted Tekton design proposal that was accepted by the project. The adopter plans to continue using Tekton and looks forward to the project joining the CNCF. 

## Perception
_What is your perception in terms of the project’s:
Community openness
Governance
Community growth potential
Maintainer diversity and ladder
Maintainer response_

The adopter views the community as kind and open and open to new ideas and experimentation. Whenever an issue is opened, the community responds positively and this was a reason for moving to Tekton directly. The adopter thinks the governance is good and has no red flags. 

The adopter thinks there is a lot of potential growth in the community, as there are lot of end users still using JenkinsX directly that could benefit from moving to native Tekton. 

The adopter feels that the project has a diverse set of maintainers from a mix of big companies and smaller companies. Maintainers are very responsive when pull requests are opened. They also feel that it is easy to start contributing to the project.  


_How are you participating in the project community?_

The adopter has participated in the community by opening pull requests and contributing Tekton enhancement proposals. 


_Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?_

The adopter has opened pull requests and had engagement with maintainers via Github and Slack. They proposed changes, got feedback, and were able to iterate and merge pull requests. The adopter is now able to use features and changes they have proposed in their production clusters and feel they had an acceptable outcome. 


## Project Strengths
_In your opinion, what are the overall strengths of the project?_

Overall, the ease of integration and ease of customization is a strength. The integration with Kubernetes is another key strength. The openness of project adopt new features is another strength. 


## Project Improvements
_Is there something you feel that holds the project back from reaching its ultimate potential?_ 



_In your opinion, what can the project do better?_

Tekton could provide a better capability to obtain logs. Logs are generally within the pod that ran, and when the pod is removed the logs are no longer available. The adopter is waiting for the ability to store pipeline results outside the cluster. The adopter is considering contributing something to Tekton to help improve this.

The adopter also thinks that the project could provide a better way of discovering when synchronous meetings occur. 



