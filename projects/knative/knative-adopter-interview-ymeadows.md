# Knative Adopter Interview - Y Meadows

Interviewee: Adam Rich, VP and co-founder of Y Meadows

Interview Date: June 4, 2025

## Organization Intro

### Can you give us an overview of your organization and what it does?

Y Meadows provides AI-Driven Automation for Business Operations, such as order operations, customer inquiries, and account management.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

We use Knative serving. The main motivation was the scale to 0 functionality. We wanted to use that to help control costs, so that pods only use resources while they are running.

We started in 2020, and at the time there were only a few ways to scale to 0. We think Knative provides a clean way to do it. There are no restrictions on service or framework and it is open source. The project has a strong community presence and is used by GCP for their cloud run products so it is battle tested. It only has a few moving parts and minimum dependencies.

## Usage Scenario

### How long has your organization used the project?

About 5 years. 

### What were the main motivations to adopt the project and which key features do you use today?

We use Knative serving for the scale to zero feature.

### What is the current level of usage (pre-production, production) and scale?

Knative is used in real production flow for all of our clients including some big name customers. It is a core part of our platform. We have used it in production at scale, for years.

### What version is used and what is your update cadence with the project?

We upgraded recently to 1.18. We upgrade about twice a year on average.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

The install guide has decent docs. We had to make some adjustments for high availability and had to make minor config changes as to how we want to handle timeouts.

We ran into some scaling issues later on.  Per the recommendation of maintainers, we switched the network plugin to Contour which helped. We also switched to GKE Data Plane v2 (which is Cilium based). We use a lot of IPs when we scale to a large number of services with Knative, and Cilium handles that well. We can also run on spot instances in GKE to cut cost which works well with Knative.

### Did you find the information in the repo valuable to your implementation? What specifically?

The project docs are critical. They are decent. They have nice reference docs, maintain change docs, and release notes. They also have a high availability section which is important.

### Has your implementation of the project provided measurable value? 

Yes. The main value is that it reduces costs for us.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

We are happy with Knative and plan to continue using it. We plan to explore the internal TLS feature and upgrade to use the operator. We haven’t used Eventing yet, but could be using it in the future.

## Perception

### What is your perception in terms of the project’s:

#### Community openness

The community has been helpful to us when we interact on Slack and GitHub.

#### Governance

The maintainers work together in the open on Slack and GitHub and seem to be maintaining the project well.

#### Community growth potential

I believe that to cost effectively use the cloud you need to be able to autoscale and the ideal of autoscaling is to scale to zero. KNative is an excellent way of achieving that. I think that there is a lot of potential to grow usage and the community.

#### Maintainer diversity and ladder

I am not familiar enough with the maintainers and their backgrounds to comment on this.

#### Maintainer response

The maintainers are responsive and have been helpful to us. They are active on Slack and GitHub.

### How are you participating in the project community?

We have contributed to discussions on Slack and GitHub, providing our perspective as an end user.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

We had communications with [Dave Protasowski](@dprotaso) on CNCF slack, he was helpful and jumped on a call with us - really going above and beyond. We have also looked up GitHub issues for guidance. Most communications were in Slack.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

The project has a solid architecture, scales well and solid documentation. It is transparent to workloads and doesn’t require clients to treat it differently and it is designed to fit the Kubernetes philosophy. 

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

For Knative serving, I think it is really reaching its potential. We would like Kourier to be considered as the default networking plugin instead of an extension. We would love to see TLS internally becomes the default as well. We would like see operator upgrade path from Helm as well.

### In your opinion, what can the project do better?

Same answer as above