# CubeFS Adopter Interview - JD.com

Interview date: Sept. 6th, 2024
Interviewee: Mevin Zhang, Software Architect, JD.com

## Organization Intro

### Can you give us an overview of your organization and what it does?

We are from JD.com, also known as Jingdong Mall, we primarily serve the retail business.

We use CubeFS as the foundation for our entire storage infrastructure, supporting various unstructured storage scenarios across JD Retail's middleware, online and offline operations, big data, and AI training. We have multiple large resource pools for these purposes.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

We chose CubeFS for its customizability to fit our specific scenarios and its support for large-scale clusters. CubeFS also supports multiple protocols, including POSIX and S3, whereas many other options only support either A or B protocols.

In JD.com's case, different business departments use different protocols. CephFS does not support operating on the same dataset with different protocols simultaneously, which complicates the overall tech stack. Additionally, MooseFS has issues with scalability and stability.

## Usage Scenario

### How long has your organization used the project?

We started using CubeFS in 2018.

### What were the main motivations to adopt the project and which key features do you use today?

We are using most of the major features of CubeFS, including:

- Support for multiple protocols
- Management capabilities for large-scale clusters
- A design that is friendly to both large and small files
- Multi-tenant support

### What is the current level of usage (pre-production, production) and scale?

We have been using CubeFS in the production environment for over six years. We have multiple clusters, with the largest cluster consisting of more than 4,000 servers, utilizing a multi-module hybrid deployment. Each node runs different CubeFS modules. Our total storage capacity exceeds 300TB, and we have over 1 million clients online simultaneously.

### What version is used and what is your update cadence with the project?

We are using the 2020 version of CubeFS from the community.  
Bug fixes are manually back ported without a fixed schedule.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

CubeFS provides general protocols, so no modifications were needed during the application integration process.

However, due to the unique distribution of hardware resources and IDC, and our company's high availability requirements, we had to customize resource scheduling strategies. This included determining how replicas should be distributed and how new resources should be allocated based on the company's IT architecture.

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

The project design docs clearly explained the current key design ideas and the reasons behind. They were very helpful for us to better understand and work on customizing & extending CubeFS.

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc

We adopted CubeFS to build our storage system from scratch, as we previously lacked such a large-scale storage solution.

This implementation reduced the maintenance and resource costs for different business teams.

By utilizing centralized large resource pools, we improved resource utilization. Storage resource usage reached 80-90%, with some environments achieving usage rates up to 98%.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc

We have been actively engaging and collaborating with the community, sharing our solutions and use cases, proposing features, and contributing implementations.

One of the features we are recently discussing is: distributed cache.

We also participate in in-depth discussions and exchanges through community meetings and ad-hoc sessions.

### JD.com was the initiating organization of the CubeFS project, and some of the initial maintainers later on joined other organizations. What motivates you to continue to use CubeFS and engage with the community?

Most crucially, CubeFS fulfills our technical and business requirements, and has been successfully and extensively deployed on a large scale in production since its early stage.

We consider JD.com and the community as a mutually supportive relationship, which motivates us to keep participating and helping the community grow


### For the heavily customized version within JD.com, what are the considerations for the follow-up? E.g. better aligned with the upstream code base, decoupled downstream enhancements

We have had some discussions with the community, but currently, we don’t have enough energy to fully align with the upstream code base. However, in the long term, we hope to better integrate our version with the upstream one. We will contribute some of our features back to the community to help that process.

## Perception

### What is your perception in terms of the project’s

- Community openness
- Governance
- Community growth potential
- Maintainer diversity and ladder
- Maintainer response

The community governance is relatively rational, with a clear ladder and well-defined responsibilities from Contributor to Maintainer.

Maintainers are responsive, actively dealing with GitHub issues and PRs. The community WeChat group is also quite active, with maintainers, other adopters, and contributors all taking part in the discussions.

### How are you participating in the project community?

We contribute by sharing our internal solutions and insights with the community. We also share issues and thoughts when we synchronize differences between the open-source version and JD.com’s internal version, including code and solution-level.

Additionally, JD.com has several project maintainers who participate in community activities such as PR reviews.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

We engaged with community members and maintainers through GitHub and WeChat groups. The communication methods and timeliness met our needs effectively.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

1. Support for extreme scenarios.
2. CubeFS provides functionalities that numerous other projects lack.
3. It fulfills the internal requirements of large enterprises such as JD.com:
  a. Ultra-large scale.
  b. Support for an enormous number of tenants.
  c. Support for multiple protocols (simultaneous operations and data visibility), which is vital for companies with complex internal structures.

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

The items planned in the current roadmap e.g. 1.”Support for distributed caching”, 2. “Support for RDMA devices”, are quite valuable to us.

### In your opinion, what can the project do better?

From a community perspective, the project could:

- Organize and participate in more events to encourage greater involvement in the CubeFS community.
- Promote wider adoption of CubeFS by engaging more users.
