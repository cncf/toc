# FabEdge Sandbox Annual Review

## Background

FabEdge is a secure container networking solution based on Kubernetes, focusing on edge computing. It enables cloud-edge,  edge-edge collaboration and solves the problems including complex  configuration management, network isolation, unaware of the underlying  topology, etc. It supports weak network, such as 4/5G, WiFi, etc. The  main use cases are IoT, IoV, smart city, etc.

FabEdge supports the major edge computing frameworks ,like KubeEdge/SuperEdge/OpenYurt.

FabEdge not only supports edge nodes (remote nodes joined  to the cluster via an edge computing framework such as KubeEdge), but  also edge clusters (standalone K8S clusters).

## DevStats

[FabEdge’s DevStats Page](https://fabedge.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m)

## Maintainers

| Maintainer   | GitHub Username                                | Company       |
|--------------|------------------------------------------------|---------------|
| Jianbo Yan | [yanjianbo1983](https://github.com/yanjianbo1983) | BoCloud |

We're are collaborating students from colleges to develop this project. 

## Adoption

* BoCloud

*PS: We have people asked questions but don't have any proof that they're using it in their products.*

## Goals Since Last Review

This is the project’s first annual review, but since FabEdge joined CNCF, we have 4 releases which includes major features like:

* Multi-cluster communication and cross-cluster DNS service.
* Dual-stack network.
* Hole-punching to enable edge nodes to establish VPN tunnels between each other even if NAT network exists.
* Auto-networking on edge nodes in the same LAN. 
* Provide coredns and kube-proxy capability on edge side.

## Current Project Goals

The goal of FabEdge this year is to increase the diversity in both contributors and adoption. Some key works include:

- Provide more documents and videos to help users.
- Improve compatibility with different linux distros, kubernete, k3s and Kubeedge/OpenYurt/SuperEdge.
- Add more unit tests.
- Implement more features like HA solution,  support other VPN softwares or vxlan.
- Support more modes of flannel and calico.

## How the CNCF can help to achieve the upcoming goals

We can manage it ourself for now.

## Incubation

We're not ready for that.