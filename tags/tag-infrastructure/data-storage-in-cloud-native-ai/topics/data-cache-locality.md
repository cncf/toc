### Data Cache

With the increasing popularity of data lakes and compute-storage disaggregation, the compute tier is now decoupled from the storage tier in both data analytics and AI/ML applications. Data caching plays a crucial role in accelerating data loading, minimizing data transfer between compute and storage layers, and reducing API calls, an often overlooked cost in object storage access. Data access patterns differ slightly between data analytics and AI/ML workloads.

Modern query optimization strategies, such as dynamic filtering on columnar data files, typically result in small, disparate reads rather than sequentially reading large chunks of data. Conversely, AI/ML workloads lean towards hybrid data access patterns, where random and sequential reads coexist based on the specific AI/ML scenario.

To fit the needs of varying workloads, the data cache needs to be scalable and elastic. A hierarchical caching design, with a local cache on each compute node and a distributed cache between compute and storage layers, can be effective for handling high volumes of data in data and AI applications. Integrating such a unified data cache into different components of the machine learning life cycle, including data preprocessing, feature engineering, model training, and model serving, can potentially improve the overall system performance⁷.

### Data Locality

In the realm of data processing and machine learning, data locality is a crucial factor that significantly impacts performance and cost. The cost considerations include not only data transfer expenses but also the hidden costs associated with low CPU/GPU utilization rates.

Several strategies can be used to achieve data locality, with each offering advantages and challenges:

#### Reading Data Directly from Remote Storage
- **Benefits**: This approach requires minimal setup effort.
- **Challenges**: Every training epoch requires re-reading all data from remote storage. Since multiple epochs are often necessary for better accuracy, this method can lead to significant time spent on data loading rather than training.

#### Copying Data to Local Storage Before Training
- **Benefits**: This method ensures that all data is local, thus gaining all the benefits of data locality.
- **Challenges**: Management can be difficult. Users must manually delete training data after use, and cache space is limited. For very large datasets, the benefits can be limited by local storage capacity.

#### Local Cache Layer for Data Reuse
- **Examples**: Tools like S3FS with built-in local cache and Alluxio Fuse SDK.
- **Benefits**: Reused data remains local, and the cache layer handles data management, eliminating the need for manual supervision.
- **Challenges**: Cache space remains limited, thus for large datasets, the benefits might still be constrained.

#### Distributed Cache Layer
- **Benefits**: A distributed cache system ensures that data is either local or adjacent, offers centralized data management, and provides scalable cache space.
- **Challenges**: Building and maintaining a distributed caching system can be complex and resource-intensive.

Data locality offers significant performance gains and cost savings, particularly for data-intensive applications and machine learning workloads. However, as discussed above, each strategy comes with its own set of trade-offs. Users must carefully consider their specific needs and constraints when selecting an approach to maximize the benefits of data locality.

A notable solution for addressing data locality issues is Fluid⁸, a CNCF sandbox project. Fluid offers significant performance gains and cost savings, particularly for machine learning workloads. It enables connecting to remote storage and supports local and/or distributed caching using Kubernetes-native approaches, significantly simplifying data locality management and accelerating AI workloads. Many of these benefits are enabled by Fluid runtimes, such as Alluxio, an open-source data orchestration and distributed caching system.

---

## References

7. https://www.alluxio.io/blog/data-caching-strategies-for-data-analytics-and-ai-dataai-summit-2023-session-recap/
8. https://fluid-cloudnative.github.io/
