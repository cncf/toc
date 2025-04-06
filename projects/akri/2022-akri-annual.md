# Akri 2022 Annual Review

## Background

Akri is A Kubernetes Resource Interface that exposes heterogeneous leaf devices (such as IP cameras, OPC UA servers, and USB devices) as resources in a Kubernetes cluster. Akri continually detects nodes that have access to these devices and schedules workloads based on them. Simply put: you name it, Akri finds it, you use it.

### Why Akri
At the edge, there are a variety of sensors, controllers, and MCU class devices that are producing data and performing actions. For Kubernetes to be a viable edge computing solution, these heterogeneous “leaf devices” need to be easily utilized by Kubernetes clusters. However, many of these leaf devices are too small to run Kubernetes themselves. Akri discovers them using common protocols (OPC UA, ONVIF, udev, etc.) and registers them as resources using the [device plugin framework](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/).

Connectivity isn't a given at the edge, so Akri handles the dynamic appearance and disappearance of leaf devices. Akri's controller automatically brings down Pods that use devices that are no longer online and redeploys them when devices reconnect.

Additionally, Akri was built to be extensible. Akri currently supports ONVIF, udev, and OPC UA Discovery Handlers, but Discovery Handlers can be added independent of Akri by implementing a simple gRPC interface.

## Sandbox Acceptance

Akri was accepted into the CNCF Sandbox on September 14, 2021.

## Development Metrics

[Akri's devstats dashboard](https://akri.devstats.cncf.io/).

The numbers of Akri contributions, contributors, and issues has remained consistent over the last year, with the number of merged PRs showing an upward trend. On average, Akri averages [88 contributions a month](https://akri.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now) with an average of [6 PRs merged a month](https://akri.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=m&var-metric=mergedprs&var-repogroup_name=All&var-country_name=All&var-company_name=All&var-company=all&from=now-1y&to=now). PR activity has been increasing with its all-time high of 23 merges in one month occurring in April, 2023. Furthermore, the project has seen an increase in the number of repeated contributors and there has been a recent increase in feature development.

One of Akri's goals in joining the CNCF was to expand beyond its Microsoft maintainership. Since becoming a sandbox project, we've added three new maintainers, almost doubling the project's maintainer count, and maintainers now span across three companies. The project has [35 contributors](https://akri.devstats.cncf.io/d/22/prs-authors-table?orgId=1) from at least 14 companies, including Microsoft, SUSE, VMWare, Fermyon Technologies, and more.

Maintainers have stayed engaged in development, with the "PR Time to Engagement" staying consistent with an average of 6.69 hours over the last year.

The community has grown since Akri entered the CNCF Sandbox.

- We have held monthly meetings since September 2021. The meeting notes could be found [here](https://hackmd.io/UUqW3_GgQDimQ5b23rS9rg), and the recordings can be found [on YouTube](https://www.youtube.com/channel/UC90VZDjT8C7ca7zcuFi6oEQ).
- Number of contributors: [10](https://akri.devstats.cncf.io/d/22/prs-authors-table?orgId=1&var-period_name=Before%20joining%20CNCF&var-repogroup_name=All) -> 35
- Number of stars: 630 -> 960
- Number of helm chart installs (measured in Akri agent container downloads): [2,975 of `v0.10.4`](https://github.com/project-akri/akri/pkgs/container/akri%2Fagent/86463296?tag=v0.10.4) (released 2 months ago) and [19,657 of `v0.8.23`](https://github.com/project-akri/akri/pkgs/container/akri%2Fagent/47340592?tag=v0.8.23) (released 8 months ago)
- The [#akri](https://kubernetes.slack.com/messages/akri) channel on Kubernetes has 279 members since its creation upon joining the CNCF

## Maintainers

We now have maintainers from 4 different companies.

- Kate Goldenring (Fermyon) @kate-goldenring
- Edrick Wong (Amazon) @edrick-wong
- Nicolas Belouin (SUSE) @diconico07
- Adithya Jayachandran (Microsoft) @adithyaj
- Johnson Shih (Microsoft) @johnsonshih
- Roaa Sakr (Microsot) @romoh
- Jiri Appl (Microsoft) @jiria
- Brian Fjeldstad (Microsoft) @bfjelds 
- Yu Jin Kim (Microsoft) @yujinkim-msft

## Project Progress and Adoption

The Akri community has released five versions since joining CNCF, with 18 different contributors to these releases in total.

Akri has also since participated in many CNCF talks and hosted booths at the Project Pavilion in KubeCon EU 2022 and 2023. You can view some of our talks here:

- [Wasm Day 2021](https://www.youtube.com/watch?v=8hF9hnUJyCY)
- [KubeCon EU 2022](https://www.youtube.com/watch?v=CMthyqMhuq4)
- [KubeCon NA 2022](https://youtu.be/0Gnh86queVQ)
- [Kubernetes on Edge Day EU 2023](https://youtu.be/hT7bGNKEZk4)

There has also been involvement from Microsoft's AKS Edge Essentials and how it works with Akri to address challenges at the edge. Akri has been mentioned in their docs and various talks which you can find here:

- Akri docs on AKS Edge Essentials
    - [Discovering OPC UA servers](https://learn.microsoft.com/en-us/azure/aks/hybrid/aks-edge-how-to-akri-opc-ua)
    - [Discovering ONVIF cameras](https://learn.microsoft.com/en-us/azure/aks/hybrid/aks-edge-how-to-akri-onvif)
- [WASM I/O 2023](https://youtu.be/L6IhonIUOKU)
- [Arc Jumpstart](https://youtu.be/KLA3N8xbYPE)

## Project Goals

This is the Akri project's first Annual review after joining the CNCF Sandbox.

Since acceptance into the sandbox, Akri has added the following features:

- Short-lived Job brokers
- Passing secrets, configMaps, and volumes to be mounted in brokers
- Ability to mount connectivity information for multiple devices in a Pod
- Expansion of the udev discovery handler to discover multi-node devices
- ...and more. View the full release notes [here](https://github.com/project-akri/akri/releases).

Some feature development goals for 2023 for Akri are:

- Integrated K8s native device secrets management solution
- Support more IoT protocols
- Get Akri to 1.0
- view our roadmap [here](https://hackmd.io/@akri/HkiDQ7rDh).

In the 2023 year, we'd also like to increase participation in monthly community meetings and grow contributor count.


## CNCF Support

The CNCF has provided significant support to Akri since it joined both through Project Pavilion opportunities at Kubecon and help desk support. We would love to collaborate with the CNCF on ways to collaborate with other CNCF projects and retain contributors.

## Incubation Readiness

Akri's feature development has accelerated; however, we want to have a push towards a stable Akri `v1.0` before we apply for incubation in the upcoming months. As a part of this effort, we hope to: 

- Implement production-ready features (list can be found [here](https://hackmd.io/@akri/HJqQWiQs9))
- Revisit the security/threat model and update the scenarios as necessary
- More end-to-end tests for automatically validating releases
