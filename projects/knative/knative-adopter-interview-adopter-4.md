---
title: Knative Adopter Interview - Adopter 4

---

# Knative Adopter Interview - Adopter 4

## Organization Intro


### Can you give us an overview of your organization and what it does?

Adopter 4 provides a platform for AI workloads.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

Knative was the leader in the space and provided the right set of features and abstractions. Adopter 4 did not want to build something from scratch themselves. 


## Usage Scenario

### How long has your organization used the project?

For about four years.

### What were the main motivations to adopt the project and which key features do you use today?

* Autoscaling (Knative pod autoscaler which takes advantage of user metrics)
* Throughput and latency for metrics
* Scale to zero is particularly useful. 
* Revision management and gradual rollout. 
* Handling complex network configuration. 


### What is the current level of usage (pre-production, production) and scale?

Knative is used in production for Adopter 4's platform. 

Scale varies a lot by customer. Customers use it for both development and production, as well as with both medium and big models, so scale is not that big per customer. There are generally dozens of workloads per customer, but only several replicas. 


### What version is used and what is your update cadence with the project?

Adopter 4 is focusing on 1.18. There are different flavors of the platform. Some flavors are more managed by Adopter 4, and in some the customer is responsible for installing. Generally, they follow a twice-a-year “validation” for the latest versions. 


### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Adoption was pretty straightforward. It works really well most or all of the time. There were no issues that were not easy to spot or related to something that Adopter 4 did wrong. 

### Did you find the information in the repo valuable to your implementation? What specifically?


The docs are generally good, and in some areas excellent. In some areas, they are less organized or missing. For example, there are different ways to configure a service and revision, including per revision and global settings. Most of those settings are well organized, but many of them are scattered across different sections, which makes it hard to get the full picture or to know where to look for a particular feature. For some more advanced topics (for example, the different replica load balancing strategies), the formal docs don’t cover everything and info can only be found by reading code or sometimes markdown files in the repo. Reading code is a common way to address advanced topics.

### Has your implementation of the project provided measurable value? 

Yes, Knative provides relevant features and it’s easy to measure cost savings from autoscaling and scale to zero. Customers can easily measure this. Adopter 4 also realized cost savings by avoiding developing a custom solution. 

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

Not planning to be involved in the community, but will keep using Knative.

Inference is changing a lot. With GenAI and LLMs and distributed inference, Knative is becoming a bit outdated. You either can’t use it or you need to spend a lot of time/effort to make it work. It might not be the right approach for emerging workloads.  


## Perception

### What is your perception in terms of the project’s:

#### Community openness

#### Governance

#### Community growth potential

#### Maintainer diversity and ladder

#### Maintainer response

### How are you participating in the project community?

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

Adopter 4 is generally not familiar with the community. They feel that the project is very mature for what it does and rarely need to engage with the community. They do leverage GitHub, but generally find the answers in documentation or in old GitHub issues. 

## Project Strengths

### In your opinion, what are the overall strengths of the project?

Autoscaling and scale to zero are particularly useful. 

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

There are some conflicts or clashes between Knative and KServe. It might be beneficial for those projects to merge; there is some perception that KServe is more mature/better than Knative. The project could provide better documentation to clarify the differences between the projects and highlight use cases where it excels.

### In your opinion, what can the project do better?

The docs could be better, and Knative could do more to stay relevant with the latest developments in inference serving. Overall, the project is really great for what it does, but needs to take a step forward to evolving workloads. 
