### Training and its Storage Usage Patterns

Unlike inference workloads, which are predominantly request–response and latency-sensitive, model training is a long-running, distributed, throughput-oriented workload. The primary objective during training is to maintain high utilization of accelerator resources (GPUs/TPUs). From a storage perspective, this translates into sustained, large-scale data movement with predictable read pressure and periodic write amplification events.

Training workloads in cloud-native environments impose distinct requirements on storage systems, particularly in Kubernetes-based deployments where storage is abstracted through CSI drivers and must operate across nodes, zones, and failure domains.

#### Key Storage Usage Patterns in Training

- **Read-Intensive Throughput**: Training involves "shuffling" and streaming petabyte-scale datasets (images, text tokens, or sensor logs) across hundreds of distributed nodes. Storage must support high-bandwidth, parallel access to prevent the "I/O Wait" bottleneck, where GPUs sit idle waiting for the next batch of data.

- **Checkpointing and Resumeability (Write-Burstiness)**: Distributed training is prone to hardware failures or preemption (e.g., using Spot instances). To mitigate loss, the system periodically saves the entire model state (checkpoints) to persistent storage. This creates massive, synchronized write bursts across the cluster, requiring storage backends with high aggregate write performance and metadata stability.

- **Randomized Data Access and Shuffling**: To improve model generalization, training pipelines typically shuffle data and perform pseudo-random access across datasets. This disrupts sequential read optimizations and produces high fan-out random read patterns. Storage systems must tolerate non-sequential access, high IOPS across distributed workers, concurrent metadata lookups. Cloud-native storage solutions often use distributed file systems (parallel FS) or specialized CSI drivers with local NVMe caching to provide the low-latency random seeks required for diverse data loaders.

- **Metadata Overhead**: Training datasets often consist of millions of small files (e.g., individual JPEGs or JSON snippets). This puts immense pressure on the storage metadata server. Architectures often move toward "containerized" data formats like TFRecord, WebDataset, or Parquet to aggregate small files into larger, more manageable chunks, reducing metadata operations.

- **Data Versioning and Lineage**: Reproducibility is a core tenet of MLOps. Storage must support immutable snapshots or integration with versioning tools to ensure that the exact dataset used for a specific training run can be audited or retrained in the future.

- **Scratch Space and Intermediate State**: During the training process, temporary data such as shuffled buffers or gradients are often stored in "scratch" space. This requires high-speed, ephemeral storage, often HostPath or Local Persistent Volumes (PVs), to minimize the latency of intermediate calculations.

## References

1. https://parquet.apache.org/
2. https://www.tensorflow.org/tutorials/load_data/tfrecord
3. https://github.com/webdataset/webdataset
4. https://pytorch.org/
5. https://www.tensorflow.org/
6. https://www.deepspeed.ai/
7. https://horovod.readthedocs.io/
8. https://dvc.org/
9. https://www.kubeflow.org/docs/components/trainer/
10. https://docs.ray.io/en/latest/data/data.html