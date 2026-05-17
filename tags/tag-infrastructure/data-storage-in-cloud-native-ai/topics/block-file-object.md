### Block, File, and Object Storage

As described in the CNCF Storage Landscape Whitepaper, block storage is best suited for workloads that require availability, low latency performance, and good throughput performance for individual workloads, but it is less suited for workloads that require capacity scaling and sharing data with multiple workloads simultaneously. Block storage can be used accordingly for AI/ML workloads depending on the requirements.

File system based storage is best suited for use cases that need to share data with multiple workloads simultaneously, and need optimized throughput for aggregated workloads. So file system based storage is also a good choice for AI/ML workloads.

The Container Storage Interface (CSI) is a standard for exposing arbitrary block and file storage systems to containerized workloads on Container Orchestration Systems (COs) like Kubernetes¹⁶.

Object storage is best suited for workloads that require availability, large capacities (PB scale), durability, sharing data with multiple workloads simultaneously, and optimized throughput for parallelised workloads¹⁷. This makes object storage an excellent choice for AI/ML workloads. It is also increasingly favored for analytic databases, where it enables separation of compute and storage necessary to support diverse workloads on very large datasets.

Object storage is least suited for workloads that require low latency performance. A cache layer can be used together with object storage to alleviate the latency constraint as well as reduce costs due to excessive API calls. Object storage I/O occurs through API calls, which means that object storage reads do not benefit from the OS page cache. For this reason, workloads like data warehouses that perform large numbers of repeated reads from object storage almost always include local and/or distributed caches.

Container Object Storage Interface (COSI) is a Kubernetes storage project that introduces standard APIs for the provisioning and consuming of object storage in Kubernetes¹⁸.

### FUSE CSI Driver

FUSE (Filesystem in Userspace) is an interface that allows non-privileged users to expose their filesystem to Unix and Unix-like operating systems. In other words, by implementing the interface, a storage system can be translated into a virtual file system. Users are able to access the storage system as if accessing the local filesystem. The FUSE process can be a long-running process inside its own pod and the filesystem can be mounted onto the host machine or VM. Whenever an application pod gets scheduled on the host, it uses the hostPath volume to access the storage system.

However, the FUSE process may not always be favored because different workloads may cause underutilization of resources such as CPU and memory. This is where CSI drivers play their role. Only when the application pods get scheduled, the CSI drivers will get triggered to start the FUSE process, either in a separate pod or in a sidecar mode, that is, injecting another container into the application pod using webhook. Whether the CSI drivers support pod or sidecar depends on the individual implementations.

Major cloud providers have been supporting FUSE CSI Drivers for their object storage services. For example, Azure Blob CSI Driver was released at the end of 2022, GCS Fuse CSI Driver was released in 2023, and AWS mountpoint S3 CSI Driver was released in 2023.

---

## References

16. https://github.com/container-storage-interface/spec/blob/master/spec.md
17. CNCF Storage Landscape Whitepaper
18. https://kubernetes.io/blog/2022/09/02/cosi-kubernetes-object-storage-management/