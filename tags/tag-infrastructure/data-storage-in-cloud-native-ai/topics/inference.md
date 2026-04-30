### Inference and its Storage Usage Patterns

While training focuses on the high-throughput ingestion of massive datasets to create a model, inference is the "production" phase where the model is used to make predictions on new data. In a Kubernetes environment, inference workloads are often deployed as microservices (e.g., using KServe, vLLM, Seldon Core, and more). These workloads are characterized by sensitivity to latency, spiky traffic patterns, and the need for rapid model loading during scaling events.

#### Key Storage Usage Patterns in Inference

1. **Model Repository Access (Read-Heavy)**: The primary storage interaction is loading model weights (often several GBs to TBs for LLMs) from a central repository or object store (S3, GCS) into the GPU/CPU memory.
   - Model Caching is a common implementation to store the model locally and reduce service startup time by loading the model directly from the cache.

2. **Model Versioning and Rollouts**: Inference requires strict consistency between the model version and the application code. Storage must support atomic switches between model versions to enable "Blue-Green" or "Canary" deployments without downtime.

3. **Model loading and Artifact access patterns**: Inference workloads use optimized loading techniques to balance storage I/O, memory usage, and startup latency. Model artifacts are commonly memory-mapped to enable on-demand paging rather than fully loading weights at startup, and may be stored as sharded checkpoints to reduce per-node memory requirements at the cost of additional coordination during loading. For very large models, lazy loading further limits initial I/O, while quantized and storage-efficient formats reduce storage footprint, memory consumption, and data transfer overhead during inference.

4. **Feature Stores and Real-time Lookups**: For many inference workloads, such as recommendation or fraud detection, model inputs include features retrieved dynamically from external systems rather than solely from the request payload. These feature stores must support low-latency, high-throughput access, often through in-memory caching or colocated replicas. Storage reliability and consistency are critical to ensure feature values remain aligned with model expectations at serving time.

5. **KV Caching (Key-Value Cache for LLMs)**: In generative AI inference, the "KV Cache" stores the mathematical representations of previous tokens in a conversation to avoid redundant calculations. vLLM provides an open source implementation for this.
   - Prefix Caching (KV Cache): Store pre-computed prompt states, enabling faster responses for repetitive or long-context queries.

6. **Cold Starts, Scaling, and Storage Locality**: Inference workloads exhibit highly variable traffic, often triggering rapid scale-out events. Without local model caching, new replicas may simultaneously fetch large model artifacts from remote storage, increasing startup latency and storage contention. Production systems mitigate this through node pre-warming, warm replica pools, and init containers or sidecars that stage models on local storage. Ensuring storage locality between model data and compute resources is critical for predictable inference latency during scale-out.

#### References

- [Disaggregated Serving - NVIDIA Glossary](https://www.nvidia.com/en-us/glossary/disaggregated-serving/)
- [BentoML - LLM Inference](https://bentoml.com/llm/)
- [KServe - Standardized AI Model Serving](https://kserve.github.io/website/)
- [vLLM - High-throughput and Memory-efficient LLM Inference](https://docs.vllm.ai/)
- [llm-d - Kubernetes-Native Distributed LLM Inference](https://llm-d.ai/)
- [Seldon Core - MLOps Framework for Production ML](https://docs.seldon.io/projects/seldon-core/en/latest/)
- [NVIDIA TensorRT - High-Performance Deep Learning Inference](https://developer.nvidia.com/tensorrt)