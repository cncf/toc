# CNCF Networking WG Proposal

## TOC Sponsor

Ken Owens

## Objective

Explore cloud native networking technology and concepts around the container networking interface (CNI).

## Goals and Expected Outcomes

* Recommend CNI be adopted as initial network interface specification focused on connectivity and portability as an official CNCF project.
* Adopt implementations of CNI that have traction in the cloud native ecosystem
* Define cloud native networking patterns
* Define the Policy framework and network services model
* A network plugin author should be able to write one “plugin” (a container) that “just works” across all container orchestration (CO) systems. 
* Enable container orchestrator to present network interfaces to the users in a portable manner that is focused on connectivity initially.
* Support dynamic provisioning and deprovisioning network primitives through this interface.
* Support group of entities that are uniquely addressable that can communicate amongst each other. This could be either an individual container, a machine, or some other network service (e.g. load balancing, firewall, VPN, QoS, Service Discovery). Containers can be conceptually added to or removed from one or* more networks.
* Focused on cloud native application patterns. This includes VM-based, Bare metal based, and FaaS (TBD) based.
* Define policy framework for network isolation

## Non Goals

* Provide or dictate an implementation.
* This includes dictating plugin lifecycle management
* Plugin distribution
* Protocol-level authn/authz
* Plugin discovery
* Not going to make a one network standard for all
* Not going to focus on individual projects per service but rather projects that model network services and patterns not going to be prescriptive but more reference guidelines and patterns

## Interested Parties

* Ken Owens (@kenowens12) [lead]
* Ben Hindman (@benh)
* Alexis Richardson (@monadic)
* Jonathan Boulle (@jonboulle)
* Lee Calcote (@lcalcote)
* Madhu Venugopal
* Jie Yu
* Deepak Bansal
* John Gossman
* Christopher Liljenstolpe (@liljenstolpe)
* Bryan Boreham (@bboreham)
* Minhan Xia (@freehan)
* Daniel Nardo (@dnardo)
* Pengfei Ni (@feiskyer)
* John Belamaric (@johnbelamaric)
* Thomas Graf (@tgraf__)
* Jason Venner (@jvmirdel)
* Doug Davis (@duglin)