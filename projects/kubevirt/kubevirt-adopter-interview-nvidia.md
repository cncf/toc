# KubeVirt Adopter Interview - NVIDIA

Interviewee: Alay Patel

Interview Date: July 28, 2026

Interviewer (TOC): Jeremy Rickard

## Organization Intro

### Can you give us an overview of your organization and what it does?

The GeForce NOW Platform is a platform for running online games. Workloads for games are virtual machines. 

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

The scope of projects that solve for virtual machine management is somewhat limited. VMware offers a propietary product, but the organization was looking for an open source solution and wanted to leverage Kubernetes. The workloads must be in virtual machines and for that KubeVirt was perfect. 

## Usage Scenario

### How long has your organization used the project?

KubeVirt has been in use for several years. KubeVirt has been in use since at least version 0.35. It's been a long journey. 

### What were the main motivations to adopt the project and which key features do you use today?

The ability to use Kubernetes to manage virtual machines and being an open source solution were the main motivations. 

### What is the current level of usage (pre-production, production) and scale?

KubeVirt is used in production.

Follow-up clarification from NVIDIA:

> We have about 50-60 clusters. Some of the largest clusters have about ~800 nodes and with 3500 VMs.

### What version is used and what is your update cadence with the project?

Usually lag behind two versions. Currently using 1.7.  One of the largest adopters of KubeVirt and rolling out a new version takes quite a while. 

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Initially, upgrading was a problem. There were issues in KubeVirt that needed to be addressed, and jumping through versions resulted in some API breaks. The API surface is now very stable and has good testing for APIs now. When KubeVirt became v1, adopter was able to put in API round trip tests to help validate updates. 

### Did you find the information in the repo valuable to your implementation? What specifically?

Adopter was looking for documentation on how to integrate GPUs into VMs. There are clear examples in the KubeVirt repository, along with a user guide on the KubeVirt website. 

There is also clear documentation about contributing to the project in the repository that was helpful. 

There are also perf and scale numbers shipped with each relase that has helped qualify each release. 

### Has your implementation of the project provided measurable value?

The biggest value is the fact that the project is open source. Replacing KubeVirt would be challenging, and the APIs the integration with the Kubernetes ecosystem has helped tremendously. 

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

Several members of the team are working on enhancements in KubeVirt. There are 5-6 enhancements being driven by the team. Plan to add more community members and reviewers. Currently one maintainer on the team and would like to grow more. 

## Perception

### What is your perception in terms of the project's:

#### Community openness

Community is open, there are weekly calls. 

#### Governance

There were instances where NVIDIA needed changes and the community supported changes, even though they were not related to other maintaner needs.

#### Community growth potential

Adopter indicated they believe there is still a lot of community growth potential.

#### Maintainer diversity and ladder

The ladder is documented pretty well. The project seems to have a hard time attracting new maintainers, but is actively trying to grow.

#### Maintainer response

Sometimes there is a limit on reviewer bandwidth. 

### How are you participating in the project community?

The adopter is currently a member of the KubeVirt community, and is currently working to help integrate DRA and is active in workgroup device management and SIG Scale. Currently working on becoming a reviewer. 

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

There have been a number of engagements with community members regarding features that were needed. There were engagements across Slack, Github issues/pull requests, and mailing lists and generally always reached an acceptable outcome after discussions.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

Openness of the commnity, the focus on perf and scale, as well as the API surface provided by the project. 

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

Reviewe and maintainer bandwidth are a bottleneck. There is also some work around evolving APIs, similar to how Kubernetes does with API conventions. 

### In your opinion, what can the project do better?

No real suggestions here. The community is receptive to improvements and keeps making actionalble work items. 
