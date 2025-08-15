---
title: KNative Adopter Interview - gojek

---

# KNative Adopter Interview - gojek

Interviewee: Roman Wozniak, Head of Engineering, gojek


## Organization Intro

### Can you give us an overview of your organization and what it does?

The merger of Gojek and Tokopedia created the largest technology company in Indonesia, offering ride-hailing, e-commerce, and delivery services to millions of users across Indonesia and Singapore.


## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

In 2019, my team and I worked closely with data scientists at Gojek to help them productize ML deployments and integrate them with the rest of the engineering systems. This approach worked well momentarily, but it didn’t scale as the company grew (hiring more ML engineers to embed them into product streams was expensive and time-consuming). That’s when we decided to build a scalable self-serve platform that would cater to data scientists' needs and offer them easy-to-use interfaces. We had a feature store (hist) and looked at the missing pieces of the lifecycle. We started examining existing platforms and did not find anything that matched our needs and was open-source and easy to use.


## Usage Scenario

### How long has your organization used the project?

We have used Knative since 2020, first as a dependency of KF Serving (now, KServe) and later as an independent component of our DS Platform. 

### What were the main motivations to adopt the project and which key features do you use today?

We currently use Knative Serving. We are looking at eventing, but largely using Serving for HTTP requests. 



### What is the current level of usage (pre-production, production) and scale?

Knative has been used in production since very early in our adoption, and at very large scale. Serving millions of our users with over 100,000 RPS during peak times. 

### What version is used and what is your update cadence with the project?

The update cadence follows latest minus one. Basically everytime a new Knative and kserve release happens, we begin adopting N-1.


### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

We started building an ML model-serving component, but decided to go one level deeper. We looked at open-source components that could allow us to put together various elements and create a product to expose to end users. We first adopted Kserve and built easy-to-use tools and user interfaces on top to develop the platform. Later, we built an integration for ML experimentation based purely on Knative. It provides an API that converts to Knative resources and orchestrates those. 

We did not encounter any major challenges while adopting. Post-COVID, cost optimization was a significant focus to understand how to use Knative more efficiently, specifically regarding right-sizing and tuning. We aimed to explore vertical autoscaling, but Knative primarily supported horizontal scaling at the time. We proposed a feature request to integrate VPA into Knative, but it wasn't finalized. This was not due to the project's unwillingness; rather, the complexity of integrating it into Knative was very high. Additionally, this was early in the Knative project, and Kubernetes also lacked maturity in this area (for instance, an API for in-place PodVerticalScaling was added to Kubernetes v1.33 and is still in beta).


### Did you find the information in the repo valuable to your implementation? What specifically?


The high level direction of the project and the project’s roadmap was very useful to help plan and understand where the project was going.

Github in general was very useful, specifically Github issues and release notes. Following the discussion on issues and pull requests helped understand the state of the project better. 


### Has your implementation of the project provided measurable value? 

Yes.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

We are looking at possible agentic orchestration

## Perception

### What is your perception in terms of the project’s:

#### Community openness

The community feels open. At KubeCon North America we were able to talk to Knative maintainers. 

#### Governance

We have a positive opinion of the governance. 

#### Community growth potential

The community feels pretty mature, don’t feel like worry about the project health. 


#### Maintainer diversity and ladder

We have a positive impression of the maintainer pool.

#### Maintainer response

Response time from the maintainers is good.

### How are you participating in the project community?

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

We have interacted with the community via community calls, slack, and on github. Github is the primary mechanism, followed by community calls, and then slack. Most of the time, we have had an acceptable outcome. The VPA scenario was less acceptable. 

## Project Strengths

### In your opinion, what are the overall strengths of the project?

Scalability is the primary strength. The abstractions defined are good and provided lots of flexibility and allowed us to build a fairly stable API on top of it so we  could focus on the problems we wanted to solve. Another stength is the the stable API. There are not frequent breaking changes. 


## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

Nothing comes to mind.

### In your opinion, what can the project do better?

Nothing is ideal, but the project is good and nothing actionable comes to mind. 
