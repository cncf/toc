*Name of project:* Kubernetes Event-Driven Autoscaling (KEDA)

*Description*

Kubernetes Event-Driven Autoscaling (KEDA) is a single-purpose event-driven autoscaler for Kubernetes that can be easily added to your Kubernetes cluster to scale your applications. It aims to make application auto-scaling dead-simple and optimize for cost by supporting scale-to-zero.

KEDA takes away all the scaling infrastructure and manages everything for you, allowing you to scale on 30+ systems or extend it with your scalers (pull-based, push-based, or REST API-driven).

Users only need to create a `ScaledObject` ([docs](https://keda.sh/docs/latest/concepts/scaling-deployments/)) or `ScaledJob` ([docs](https://keda.sh/docs/latest/concepts/scaling-jobs/)) that defines what you want to scale and what triggers you want to use; KEDA will handle all the rest!

![Picture1](https://user-images.githubusercontent.com/4345663/108470255-f9163400-7289-11eb-98fc-6a5f522202e0.png)

It allows you to scale anything; even if it’s a CRD from another product you are using, as long as it implements `/scale` subresource. This allows us to build an ecosystem of products to integrate with and make them more powerful and successful.

To provide production-grade security, we've introduced `TriggerAuthentication` and `ClusterTriggerAuthentication` ([docs](https://keda.sh/docs/latest/concepts/authentication/)) allowing users to move authentication information out of their application into a separate CRD. This not only allows for re-use, but operators can manage the authentication separately in a centralized place and avoid applications running with higher privileges than they need.