### Network Management

Networking is an increasingly important consideration for object storage as well as caches and introduces networking issues that cross layers of the cloud native These manifest in different ways depending on the workload. They may require specialized configuration to set up connectivity within Kubernetes. They also imply ample monitoring to diagnose faults and to assess capacity.

For example, object storage services often apply aggressive rate limiting policies, which lead to either performance throttling or failures at high load. Read and write requests to object storage consume application network bandwidth, hence may produce contention between applications running on the same Kubernetes worker node.

Many design patterns for communication depend on the ability to connect between regions and AZs using private networking services like AWS PrivateLink¹⁹. These design patterns include routing traffic between databases running in separate AZs. They also include design patterns to apply cheap compute to object storage, for example using AWS Direct Connect, which greatly reduces transit costs out of object storage.

---

## References

19. https://docs.aws.amazon.com/vpc/latest/privatelink/what-is-privatelink.html