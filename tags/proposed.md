# Proposed SIGs

# CNCF Big Data & AI TAG

## TAG Mission Statement

Big data analytics and machine learning are typical workloads on modern computing infrastructures. In the past few years, Kubernetes and many new technologies in the CNCF landscape have completely re-shaped the fundamentals and paradigms on how one builds and runs applications on clouds. However, compared with other areas in the cloud-native landscape, big data and AI applications still require more attention from the community and better native support from the cloud-native ecosystem. With more attention and more technical investment in this area, we believe there might be a great chance to have a new paradigm for building and deploying big data & AI applications in a modern and cloud-native way, with better performance, security, scalability, networking, and inter-operation. Such a new paradigm would benefit the whole community and bring huge potential value to CNCF as well.

There are many opportunities to make big data & AI applications fit better on cloud-native infrastructures. For example, modern data analytical pipelines become increasingly fragmented in terms of the types of workloads involved. While there are already many state-of-the-art computation engines for a single type (SQL, Tensor, Graph, ML, etc) of big data tasks, deploying a complete pipeline end-to-end often involves multiple engines and components working together. And a lot more can be done, for example, reducing the friction and overhead from inter-operation between engines, better fault-tolerance, better orchestration. By working on new technologies based on existing cloud-native building blocks and gaining consensus about the best practice from the community, the fragmented ecosystem could be unified and it would be as straightforward as deploying a web service on Kubernetes to deploy a big data analytical pipeline to modern cloud-native infrastructures.

The bigdata TAG's mission is revisiting current challenges for big data & AI applications in cloud-native environments, figuring out better solutions from the view of computing infrastructures, and reach the consensus of a new cloud-native paradigm for more efficient and user-friendly deploying of big data & AI applications on clouds. The TAG focuses on four key areas:
​
1. revisiting current solutions for big data & AI applications on cloud-native environments and figuring out the problems in areas of networking, security, storage, heterogeneous hardware and AI accelerator, scalability, scheduling, monitoring, operation and maintenance automation, fault-tolerance, etc., and see how the cloud-native community could help on those challenges.
2. inter-operability between the components that target to support big data & AI applications on clouds to mitigate the fragmented situation of the community.
3. the cloud-native paradigm for efficiently developing, deploying, and operating big data & AI applications, and consensus for providing user-friendly interfaces to the end-users in the community.

We need projects that cover areas in cloud-native support for big data and AI applications that are:

1. standards about the interface and improve the cross-system interoperability
2. libraries for data sharing and reusing
3. controllers for data management
4. integration with existing de-facto computation frameworks
5. tools for testing and evaluating

## Responsibilities & Deliverables

### Responsibilities

- identifications of areas of focus, e.g., data caching, life cycle management, commonly-used efficient data layout, scheduling, etc.
- framework and suites for test and evaluation
- technical evaluations to find what is still missing in the cloud-native landscape to make big data and AI applications easier and more efficient
- integration with existing de-facto big data & AI engines
- recommendations of moving big data & AI applications from traditional deployment to cloud-native environment and adopting the new paradigm
- growing CNCF external relationships with the big data community and cloud vendors.

### Deliverables to ToC

- framework documentation for the cloud-native big data & AI computation.
- landscape for the big data computing space
- recommendation for new projects the current problem and where more attention is still needed
- recommendation for existing CNCF projects where there is a bridge with the big data & AI community
- scheduled regular reporting to TOC about the progress and achievements

The benefits in efficiency and scalability that modern cloud-native infrastructure brings to bigdata & AI applications could also be maximized with better native support from the cloud-native community, as bigdata & AI applications follow a different pattern and expose many new challenges that need to be addressed to the runtime environment, such as resource management, the cost of data movement, as well as scheduling policy that take data placement into account. With a new paradigm that focuses on big data & AI applications from CNCF, it is possible to avoid data scientists deploying their application in anti-pattern and less efficient ways.
