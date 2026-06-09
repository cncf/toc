# Network Management

Networking is an increasingly important consideration for object storage as well as caches and introduces networking issues that cross layers of the cloud native stack. These manifest in different ways depending on the workload. They may require specialized configuration to set up connectivity within Kubernetes. They also imply ample monitoring to diagnose faults and to assess capacity.

For example, object storage services often apply aggressive rate limiting policies, which lead to either performance throttling or failures at high load. Read and write requests to object storage consume application network bandwidth, hence may produce contention between applications running on the same Kubernetes worker node.

Many design patterns for communication depend on the ability to connect between regions and availability zones using private networking services. These services include AWS PrivateLink¹, Azure Private Link², and Google Cloud Private Service Connect³. These design patterns include routing traffic between databases running in separate availability zones. They also include design patterns to apply compute resources to object storage using dedicated network connections, such as AWS Direct Connect, Azure ExpressRoute, or Google Cloud Interconnect, which can significantly reduce transit costs and improve performance when accessing object storage.

---

## References

1. [AWS PrivateLink](https://docs.aws.amazon.com/vpc/latest/privatelink/what-is-privatelink.html)
2. [Azure Private Link](https://docs.microsoft.com/en-us/azure/private-link/)
3. [Google Cloud Private Service Connect](https://cloud.google.com/vpc/docs/private-service-connect)