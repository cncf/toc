# Resource and Infrastructure Challenges for AI Workloads

**Authors (in alphabetical order):**

Abhishek Malvankar, Alexander Scammon, Andrey Velichkevich, Ekin Karabulut, Kevin Hannon, Marlow Warnicke, Rajas Kakodkar, Sabrina DiLeva, Victor Lu, Yuan Tang

**Contributors (in alphabetical order):**

Adel Zaalouk, Andrew Shunichi Aikawa, Andrey Velichkevich, Boris Kurktchiev, Brian Redmond, Cathy Zhang, Claudia Misale, Joel Roberts, Johnu George, Josh Halley, Kay Yan, Malini Bhandaru, Michael Yao, Mike SHNg, Naadir Jeewa, Niki Manoledaki, Ricardo Aravena, Ronald Petty, Sachi Desai, Shravan Achar, Sudhanshu Prajapati, Victor Lu, Vijay Rodrigues

# Series Introduction

This paper is the fourth in a five-part series on scheduling for cloud native AI workloads:

1. Understanding AI Workloads for Kubernetes Scheduling: What AI workloads are and why they differ from traditional cloud-native applications

2. Scheduling Fundamentals for AI Workloads: How Kubernetes scheduling works and why additional layers are needed

3. Job Orchestration Challenges for AI Workloads: Gang scheduling, fairness, queues, preemption, priority, and reservation

4. **Resource and Infrastructure Challenges for AI Workloads (this paper):** Topology, GPU sharing, scalability, I/O, fault tolerance, and cost

5. Solutions and Practical Guidance for AI Workload Scheduling: Tools, reference architectures, and real-world use cases

Each paper can be read independently, but together they provide a comprehensive guide to running AI workloads on Kubernetes.

# Executive Summary

Beyond job orchestration, AI workloads face challenges related to the underlying infrastructure: hardware topology affects communication performance, heterogeneous GPU types require careful placement, expensive GPUs should be shared when possible, and long-running jobs must tolerate failures. This paper examines these resource and infrastructure challenges: topology awareness, resource heterogeneity, GPU utilization and sharing, scalability, I/O bottlenecks, fault tolerance and elasticity, and budget constraints.

# Previously

The first three papers in this series established:

* **Paper 1:** AI workloads have fundamentally different characteristics than traditional microservices. They are long-running, tightly coupled, GPU-intensive, and topology-sensitive.

* **Paper 2:** The default Kubernetes scheduler operates at the pod level. Meta-schedulers and queue managers bridge the gap for AI workloads.

* **Paper 3:** Job orchestration challenges include gang scheduling, resource fairness, queue management, preemption, priority scheduling, and resource reservation.

This paper examines challenges related to the underlying resources and infrastructure.

# Resource and Infrastructure Challenges

This section outlines scheduling challenges related to hardware resources and infrastructure—where and how jobs run. While many examples reference training workloads, these challenges apply equally to multi-node inference deployments, such as model-parallel or disaggregated serving architectures.

* [Topology Awareness](#topology-awareness)
* [Resource Heterogeneity](#resource-heterogeneity)
* [GPU Utilization and Sharing](#gpu-utilization-and-sharing)
* [Scalability](#scalability)
* [I/O Bottlenecks](#io-bottlenecks)
* [Fault Tolerance and Elasticity](#fault-tolerance-and-elasticity)
* [Budget and Cost Constraints](#budget-and-cost-constraints)

## Topology Awareness

Hardware topology significantly impacts AI workload performance. Two different challenges require different solutions:

* **Node-level topology:**  
  * NUMA (Non-Uniform Memory Access). Memory access speed depends on which CPU socket is accessing which memory bank. Placing a workload's threads and memory on the same NUMA node avoids slow cross-socket access.  
  * CPU-GPU affinity. GPUs are attached to specific PCIe buses, which connect to specific CPU sockets. Misaligned placement adds latency.  
  * GPU interconnects within a node. Eight GPUs in a server might be connected via high-speed interconnects or PCIe (slower). A job that needs 4 GPUs performs better if those 4 have high-speed connections.  
  * Unified memory architectures. Some hardware platforms use system-on-chip (SoC) designs where GPU and CPU share a single memory pool. On these architectures, GPU memory consumption directly reduces available system memory. The Kubernetes scheduler treats GPU memory and the memory available to CPU as independent resources, unaware of this coupling. This can lead to over-commitment and system instability (e.g., OOM kills) when GPU workloads consume memory that the scheduler assumed was available for system use.  
* **Cluster-level topology:**  
  * Network fabric. Nodes under the same top-of-rack switch can communicate faster than nodes across the datacenter. For distributed training with frequent synchronization, network topology dominates performance.  
  * Rack placement. Placing all workers in the same rack minimizes network hops but creates a single point of failure.  
  * GPU interconnects across nodes. Technologies like high-bandwidth networking with RDMA provide high-bandwidth, low-latency connections between GPUs on different servers. The scheduler must know which nodes have these connections.   
* **Why this matters:** For communication-heavy workloads like distributed training, placing workers on well-connected nodes can be the difference between a job taking hours vs. days. The default Kubernetes scheduler has limited topology awareness—it knows about NUMA but not about GPU interconnects or network fabric.

## Resource Heterogeneity

* **Problem:** Clusters contain mixed hardware: different GPU models, different generations, different memory sizes. Some nodes have TPUs, some have FPGAs, some have only CPUs.  
* **Solution approach:**  
  * Node pools group nodes with similar hardware. Jobs can target specific pools.  
  * Resource labels tag nodes with their capabilities (gpu-type=high-memory, gpu-generation=current).  
  * Topology hints help the scheduler understand hardware characteristics beyond simple labels.  
* **Challenges:** Jobs may have preferences rather than hard requirements—"prefers current-generation GPUs but can run on previous-generation with reduced batch size." Expressing and handling these preferences adds complexity.

## GPU Utilization and Sharing

* **Problem:** GPUs are expensive, but many workloads don't fully utilize them. An inference service or an interactive session might use 20% of a GPU's compute capacity. Allocating whole GPUs to such workloads wastes resources.  
* **Solution approaches:**  
  * **GPU partitions support** into hardware-isolated instances. A single physical GPU becomes multiple smaller virtual GPUs, each with dedicated memory and compute.  
  * **Time-slicing** shares a GPU between multiple workloads by switching between them rapidly. This improves utilization but adds latency and provides no memory isolation.  
  * **MPS (Multi-Process Service)** allows multiple containers to use a single GPU concurrently, rather than taking turns. This improves utilization for workloads that don't need a full GPU but provides no hardware isolation between processes.  
  * **vGPU** (virtual GPU) provides software-level GPU sharing, typically in virtualized environments.  
  * [**Dynamic Resource Allocation (DRA)**](https://kubernetes.io/docs/concepts/scheduling-eviction/dynamic-resource-allocation/) is a Kubernetes feature that allows flexible resource allocation, including GPU partitioning on demand which need to be implemented via vendor DRA drivers.  
* **Caveats:**  
  * Memory constraints. LLMs consume all available GPU memory. Sharing works for smaller workloads, not for large model inference.

## Scalability

* **Problem:** Large AI training runs require massive scale—hundreds of thousands of GPUs across thousands of nodes. The scheduling system must handle this scale.  
* **Challenges:**  
  * Cluster autoscaling. Tools like Karpenter and the Kubernetes Cluster Autoscaler add nodes when demand exceeds capacity. But adding GPU nodes takes time (minutes), and large jobs may need more nodes than a cloud provider can quickly provision.  
  * Workload autoscaling. Beyond cluster-level scaling, workloads need to scale dynamically. Elastic training and data processing jobs can opportunistically expand beyond quota when resources are idle and shrink back when demand returns, while respecting gang scheduling constraints. For inference, scaling is driven by request load; metrics like tokens-per-second, queue depth, and KV cache utilization may be more relevant than CPU or memory. Disaggregated inference adds complexity, as prefill and decode components may scale independently.  
  * Scheduler throughput. The scheduler must make placement decisions for thousands of pods without becoming a bottleneck. At scale, scheduling overhead matters—every millisecond spent per pod adds up.  
  * Coordination overhead. Gang scheduling at scale (coordinating placement of thousands of pods simultaneously) is harder than scheduling independent pods.

## I/O Bottlenecks

* **Problem:** GPUs are fast, but getting data to them is slow. I/O bottlenecks waste expensive GPU time.  
* **Data loading:** Research from [Google](https://arxiv.org/pdf/2101.12127.pdf) and [Microsoft](https://www.cs.utexas.edu/~vijay/papers/vldb21-datastalls.pdf) found that data loading consumes 10-70% of training execution time. During data loading, GPUs sit idle, waiting for the next batch.  
* **Checkpointing:** Training jobs periodically save model state to storage. Checkpointing a large model can take minutes and requires high storage bandwidth. Frequent checkpointing (for fault tolerance) vs. infrequent checkpointing (for performance) is a tradeoff.  
* **Model loading for inference:** Before serving requests, inference services must load model weights to GPU memory.  The time and method of loading depends on the scenario:  
  * Cold start: Loading from storage (local filesystems or cloud object storage (S3, GCS)) can take minutes for large models. Optimized streaming tools (such as Run:ai Model Streamer, Tensorizer, and fastsafetensors) reduce this to seconds by streaming weights directly into GPU memory.  
  *   
* **Storage performance:** The MLCommons Storage benchmark measures storage system performance for ML workloads. Inadequate storage throughput can bottleneck the entire training pipeline, regardless of how many GPUs are available.

## Fault Tolerance and Elasticity

* **Problem:** Long-running training jobs (days to weeks) will encounter failures. Hardware fails, nodes go down, preemption happens. Without fault tolerance, a failure after 6 days of training means starting over. The frequency of these interruptions is substantial: [Meta reported](https://arxiv.org/pdf/2407.21783) that during a 54-day snapshot of Llama 3 pre-training, they experienced 466 job interruptions: 47 planned (firmware upgrades, configuration updates) and the remainder unplanned hardware or infrastructure failures.  
* **Checkpoint/restart:** Periodically saving model state to durable storage allows resuming from the last checkpoint rather than from scratch. The tradeoff: more frequent checkpoints mean less work lost but more overhead.  
* **Elastic jobs:** Some frameworks (like PyTorch Elastic) support changing the number of workers during training. If a worker fails, training continues with fewer workers (within limits). If resources become available, workers can be added. This requires:  
  * Minimum and maximum worker counts  
  * Ability to redistribute work when workers join/leave  
  * Coordination with the scheduler to add/remove workers gracefully  
  * Role-aware scheduling: the scheduler must understand pod roles (e.g., master vs. worker) and preempt workers before masters to avoid job failure  
* **Handling failures without full restart:** For gang-scheduled jobs, one worker failure typically crashes the entire job. Elastic training relaxes this—the job continues with the surviving workers, and a replacement worker can join later.

## Budget and Cost Constraints

* **Problem:** GPU resources are expensive. Teams have budgets. Without cost controls, a single runaway job can consume a month's cloud budget overnight.  
* **Potential approaches:**  
  * GPU-hour budgets limit how much GPU time a team or job can consume  
  * Cost-aware scheduling considers the cost of different resource allocations (spot vs. on-demand instances, different GPU types)  
  * Automatic job termination when budgets are exhausted  
  * Cost visibility helps teams understand and optimize their spending  
* This challenge is mentioned in comments on the original document but not fully addressed in the current Kubernetes ecosystem. It remains an area for future development.

# Infrastructure Considerations

Scheduling decisions depend on the underlying infrastructure. This section covers datacenter architecture and network topology as they relate to AI workload scheduling.

## GPU Direct Communications

GPUs can communicate without involving the CPU, which reduces latency and increases throughput.

Within a chassis (same server):

* High-speed GPU interconnects provide high-bandwidth, low-latency connections between GPUs. Interconnect bandwidth can exceed hundreds of GB/s between connected GPUs, far faster than PCIe.

Not all GPUs in a server are necessarily connected via high-speed links. In a  GPU dense server only certain pairs have NVLink connections. The scheduler must understand this topology.

Between chassis (across servers):

* InfiniBand provides high-bandwidth, low-latency networking between servers, with RDMA (Remote Direct Memory Access) support for GPU-direct communication.  
* RoCE (RDMA over Converged Ethernet) offers similar capabilities over Ethernet infrastructure.  
* GPUDirect RDMA enables direct data transfer between GPUs on different servers without CPU involvement.

## Network Topology for AI Workloads

For communication-heavy distributed training, network topology can matter more than raw compute speed. Key considerations:

* **Locality matters.** Workers that communicate frequently should be placed close together in the network. "Close" means fewer hops and higher bandwidth—ideally on the same leaf switch.  
* **Bandwidth requirements.** Large-scale distributed training with frequent synchronization can saturate network links. The scheduler should avoid placing communication-heavy jobs in ways that create network bottlenecks.  
* **Failure domains.** Placing all workers in the same rack minimizes network hops but means a rack failure kills the entire job. Spreading workers across racks improves resilience but increases communication latency.

The scheduler must balance these concerns. For latency-sensitive training, co-location may be worth the reduced resilience. For long-running jobs, spreading across failure domains and accepting higher latency may be preferable to risking a full restart.

# What's Next

This paper examined the resource and infrastructure challenges for AI workloads. The final paper in this series, **Solutions and Practical Guidance for AI Workload Scheduling**, catalogs the tools and Kubernetes features that address these challenges, provides a reference table mapping challenges to solutions, and offers practical guidance including real-world use cases.
