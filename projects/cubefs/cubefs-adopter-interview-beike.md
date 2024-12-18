# CubeFS Adopter Interview - BEIKE

Interview date: Sept. 11th, Wednesday, 2024
Interviewee: FangRong Lin, Senior Development Engineer, BEIKE; Bin Yan, Senior Development Engineer, BEIKE

## Organization Intro

### Can you give us an overview of your organization and what it does?

Beike is an integrated online and offline platform for housing transactions and services. 
Our website offers users browsing of housing information, house photos, etc., and helps users manage transaction information and contracts for houses.
Our team belongs to Beike's private cloud and provides file storage for internal business teams.
Scenarios involving CubeFS include historical data, logs, and cold backups of databases, as well as AI training/inference scenarios.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

Due to team changes, we are not familiar with the background information of the previous technical selection.
Based on our current usage, CubeFS demonstrates advantages in overall read/write and concurrency performance, which aligns well with our company's business characteristics. The storage team customizes the system according to the needs of the business team to meet their requirements.

## Usage Scenario

### How long has your organization used the project?

Started in early 2020. CubeFS v2.3 was the first version adopted.

### What were the main motivations to adopt the project and which key features do you use today?

We primarily use the triple-replica redundancy feature. CubeFS is easy to scale, which is crucial for us.

We basically collect and evaluate storage growth requirements from all business teams quarterly. However, the growth rate is unpredictable, making it difficult to forecast storage needs in advance. CubeFS excels in this scenario with its robust scalability, ensuring high standards of data stability.

### What is the current level of usage (pre-production, production) and scale?

We are using CubeFS in a large-scale production environment with a total capacity of 30PB (triple-replica redundancy).

### What version is used and what is your update cadence with the project?

We are currently using version 3.3.1 with some downstream modifications. Additionally, some legacy environments are still running on version 2.4. We initially used version 2.4 extensively, but later upgraded to 3.3.1 to meet business requirements.

We don't have a fixed update cadence. But we will upgrade when the community releases particularly attractive features or important bug fixes.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Our business teams were using NFS previously. Compared to NFS, CubeFS has lower write performance for small files. During the migration from NFS to CubeFS, we faced challenges with the decreased write performance.

Additionally, the community had provided a tool for writing HDFS data to CubeFS, but it hasn't been maintained for a long time. When we encountered related scenarios during implementation, we faced some issues and received limited support from the community.

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

The disk throttling feature has been particularly useful.

Our team closely follows the community's annual version plans and examines the features in the development branches. For useful features, we implement them in our own version. For example, we implemented metadata support for RocksDB (the community version caches all metadata in memory) and a disk-based distributed cache (the community version is memory-based).

As for why we have some customized implementations:

1. The community version does not fully meet our performance requirements, so modifications are necessary for implementation.
2. The community's development cycle is relatively long, and our own release cadence does not align well with the community's plan.

After our own implementation is launched, we evaluate the actual performance and functionality improvements. If the results are relatively good, we will discuss with the community and try to contribute back to upstream.

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc

The environment maintenance of CubeFS is relatively convenient.

Previously, Beike did not build its own storage system and mostly adopted procurement methods. After we adopted CubeFS, many business teams used Beike's self-built solution, which can save costs and make it easier to debug and resolve problems.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc

We hope to participate in the community's plans in the future. Currently, our team is relatively small, so we will focus on our actual needs and primarily engage with parts of the community that align well with our goals. We aim to communicate and collaborate with the community to drive progress together.

## Perception

### What is your perception in terms of the project’s

- Community openness
- Governance
- Community growth potential
- Maintainer diversity and ladder
- Maintainer response

The community is quite open and friendly, and questions asked in the WeChat group are well addressed.
However the recent community developments seem slowed down, which is behind the roadmap.
As for community governance, we don’t have much participation.

### How are you participating in the project community?

We mainly participate in discussions in the WeChat group and attend the community's monthly meetings.
For valuable commits, we will contribute back to the community.
We also closely follow articles published by the community.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

Our team leader engages with the community, particularly regarding technical selection, initial cluster sizing, and performance estimation. The primary communication channel used is the WeChat group.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

1. Concurrent read and write operations for multiple files.
2. Efficient utilization of overall bandwidth, which can scale with the size of the cluster.

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

1. The management of the git commit history is not very clean. For example, when we are trying to back port a feature, it is hard to find all commits related.
2. We hope some features in the community can be made pluggable, so they do not affect existing functionality.
3. We hope the community can improve compatibility for upgrading historical versions. For instance, when upgrading from version 2.4 to 3.3.1, we encountered significant compatibility issues. Business teams may be reluctant to upgrade clients, and compatibility between old clients and new backends can become an obstacle.

### In your opinion, what can the project do better?

The community needs to enhance its promotion and outreach efforts, especially regarding usage guidelines.

CubeFS is a relatively simple and user-friendly storage solution in the industry, making it easy to use and customize. However, the number of public users is lower than expected. We suggest the community focus more on increasing its user base.
