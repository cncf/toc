# Data on Kubernetes - Data Analytics and AI/ML Workloads

**Authors:** Alexa Griffith, Chunxu Tang, Joe Huang, Lu Qiu, Raghu Shankar, Robert Hodges, Shawn Sun, Victor Lu, Wannes Rosiers, Xing Yang

## Overview

The Cloud Native AI whitepaper¹ delves into the interaction of Cloud Native (CN) and Artificial Intelligence (AI) technologies, discusses the current state, the challenges, the opportunities, and the potential solutions in this area.

In this paper, we aim at highlighting the characteristics of data analytics and AI/ML workloads and the patterns and trends in data storage to meet the new challenges.

There are different stages of AI/ML with different utilization patterns. The batch, training/fine tuning workloads have sustained utilization of resources and are long running jobs. On the other hand, the inference workloads have spiky utilization of resources and need immediate access².

Different stages of AI/ML workloads have different requirements for data storage.

Data analytics and AI/ML workloads usually contain a large amount of structured and unstructured data. These workloads need to be highly scalable, highly performant with low latency. They typically need read-only access to data. These characteristics have implications for data storage requirements.

Data warehouses and data lakes are typically used to store large data for analytics workloads. Data warehouses are optimized for Online Analytical Processing (OLAP) as compared to traditional relational databases which are used for Online Transaction Processing (OLTP). Column-oriented databases can be used in data warehouses for efficient query search³. Data warehouses are normally utilized to store structured data. Data lakes, on the other hand, can handle both structured and unstructured data.

Data caching plays an important role in data analytics and AI/ML workloads. It can help improve performance dramatically by placing data close to where it needs to be accessed and processed. It can also help to avoid redundant computation.

In recent years, vector databases have gained popularity due to their capability to do similarity search rather than the exact match used in traditional database searches. This is very important for AI/ML workloads, especially generative AI workloads.

Block, file, and object storage can be used as the underlying storage systems to store data for data analytics and AI/ML workloads. Object storage's ability to allow data to be shared with multiple workloads simultaneously, optimized throughput for parallelised workloads, and highly scalable capacities make it a popular choice for AI/ML workloads as well as data analytics.

Data warehouses and data lakes are centralized repositories. There are trends to decentralize them using the data mesh architecture for building AI/ML models. Another data paradigm is data fabric, an architectural approach and technology framework that addresses data lake challenges by providing a unified and integrated view of data across various sources⁴.

Modern Data Architecture Principles emphasize on data quality, data governance, consistency, data as a shareable asset, and data security and privacy.

## Patterns and Trends in Data Storage

- [Data Warehouses, Data Lakes, and Data Lake Houses](topics/data-lake-houses.md)
- [Data Cache and Data Locality](topics/data-cache-locality.md)
- [Databases](topics/databases.md)
- [Block, File, and Object Storage](topics/block-file-object.md)
- [Network Management](topics/network-management.md)
- [Data Pipelines](topics/data-pipelines.md)
- [Data Mesh and Data Fabric](topics/data-mesh-and-fabric.md)
- [Hardware/Software Co-design](topics/hardware-software-co-design.md)

## Storage in the AI Lifecycle

Different phases of AI workloads have distinct storage requirements and usage patterns. Training workloads are throughput-oriented with sustained data movement, inference workloads are latency-sensitive with spiky traffic patterns, and AI agents require complex state management across iterative reasoning loops.

- [Training and its Storage Usage Patterns](topics/training.md)
- [Inference and its Storage Usage Patterns](topics/inference.md)
- [AI Agent and its Storage Usage Patterns](topics/ai-agent.md)

---

## References

1. https://www.cncf.io/wp-content/uploads/2024/03/cloud_native_ai24_031424a-2.pdf
2. https://sched.co/1YhIO
3. Designing Data Intensive Applications, Martin Kleppmann, O'Reilly Media, 2 May 2017
4. From data warehouse to data fabric: the evolution of data architecture, CNCF Blog, 21 July 2023