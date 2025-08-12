---
title: KNative Adopter Interview - IBM

---

# KNative Adopter Interview - IBM

Interviewees:
* Simon Moser, Distinguished Engineer, IBM Cloud Container Services
* Sascha Schwarze, IBM Cloud Code Engine 


## Organization Intro

### Can you give us an overview of your organization and what it does?

IBM Cloud Code Engine is a cloud service partially based on Knative, that aims to provide a single serverless platform. It supports four types of workloads under one platform, all containerized, with the same set of APIs. 


## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

At the time of adopting, Knative was the most promising in the ecosystem. It supported scale to zero and supported functions, and was natively based on Kubernetes. It was not very mature a the time, but covered most of the use cases we had in mind. 


## Usage Scenario

### How long has your organization used the project?

5+ years. Platform development started at the end of 2019 or beginning of 2020.

### What were the main motivations to adopt the project and which key features do you use today?

We primarily use Knative serving. We were mainly motivated to adopt it because it was built natively on Kubernetes. Scale to zero is a key feature we use today. We use some Knative eventing, but mostly serving.


### What is the current level of usage (pre-production, production) and scale?

Knative is used as part of the production cloud service. 

Knative is used on clusters running large number of services. There are four digit number of Knative services per cluster / region. Each region would be an isolated Knative deployment. 

### What version is used and what is your update cadence with the project?

We are using 1.18. We have automation that automatically opens PRs when new version comes out. We sometimes run into issues with the minimum Kubernetes versions, but we try not fall behind too far. We do not really encounter issues w/ Knative upgrades. 

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

We don’t recall any major obstacles while adopting. Some of our challenges have been around integrating Istio due to architectural decisions, which led to larger mesh sizes than anticipated. There were also operational hiccups, but most related to Istio. The build uses ko, and results in a large YAML that cannot be all applied at once. Knative controller also tries to update all deployments when changing versions, which result in a lot of extra traffic to API server. Many of these were scale related. Upstream Knative doesn’t handle scale but IBM carries some custom patches to alleviate some of this.

### Did you find the information in the repo valuable to your implementation? What specifically?

The Knative documentation is quite good. Release notes are what we use the most now. We are not expanding our feature set. We look at issues in Github and review code as needed. 


### Has your implementation of the project provided measurable value? 

IBM Code Engine absolutely benefited from adopting Knative. Knative was a huge help in building the platform. We would not have been able to deliver the platform as quickly without Knative, as we would have had to build features or spend effort stitching together other oss projects to enable this. 

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

We are in a steady state for the moment. For Code Engine, we are focused more on batch workloads and that is outside of Knative capabilities. Investment is focused on areas outside Knative capabilities. We are working with community on bug fixes. 

## Perception

### What is your perception in terms of the project’s:

#### Community openness

Communityu openness is good.

#### Governance

The project has good governance. Meetings and held and records kept, but meetings are perhaps indecisive. 

#### Community growth potential

The project doesn’t have any obvious features missing within serving. 


#### Maintainer diversity and ladder

Contributions to Knative serving seem to have gone down over the last year, and sometimes issues may take longer than expected. 

#### Maintainer response

People are generally responsive on slack. Sometimes maintainers don’t respond as quickly as we want, but this is normal.

### How are you participating in the project community?

We are working on bug fixes and working on general PRs. IBM previously served on steering for Knative. We are not pushing any major changes. 

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

We engage with maintaienrs mostly though github issues. Outcomes are generally acceptable.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

The scale to zero feature and being natively built on Kubernetes.

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

For functions, startup of containers is an issue for real use. Low latency startup is an issue. Eventing is less desirable due to lack of upstream integrations. 

### In your opinion, what can the project do better?

The project could support more Kubernetes versions and possibly be more responsive to issues. 