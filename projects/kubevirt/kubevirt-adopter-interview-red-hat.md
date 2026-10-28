# KubeVirt Adopter Interview - Red Hat

Interviewee: Brian Atkisson

Interview Date: Jun 26, 2026

Interviewer (TOC): Mario Fahlandt

## Organization Intro

### Can you give us an overview of your organization and what it does?

Redhat IT, responsible for onpremise and cloud environments from an infrastructure prospective. Responsible for building internal platforms that are used for internal Redhat workloads and customer workloads.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

The ability to use Kubernetes and OpenShift for multiple workload types. 

## Usage Scenario

### How long has your organization used the project?

Production clusters went live in 2023, early POC starting in 2021.

### What were the main motivations to adopt the project and which key features do you use today?

Mid 2010-2020, IT org was largely based on Redhat Virtualization, as well as OpenStack. Later in that period, started using OpenShift more and began using KubeVirt.

The ability to import lots of VMs and bare metal systems into KubeVirt and concurrently run containerized workloads and serverless was very attractive.

This also allows the use of the full Kubernetes ecosystem as part of the virtualization solution, for example storage and backup solutions. Also able to leverage infrastructure as code for fleet management.

Live migration support is very useful and integrates well with Openshift and Kubernetes.

### What is the current level of usage (pre-production, production) and scale?

KubeVirt is used in production. Using around 40 clusters, 500-1000 nodes. The nodes behind KubeVirt are rather large to allow a lot of VM density.

### What version is used and what is your update cadence with the project?

For most important workloads, an update to a new major version once a year. For other clusters, frequent maintenance releases.

Initially needed new features and upgraded more frequently, but basically all the features they need are in KubeVirt.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

### Did you find the information in the repo valuable to your implementation? What specifically?

The upstream documentation is very robust. Everything needed has been found in the KubeVirt documentation, have not needed to look at code to understand anything.

### Has your implementation of the project provided measurable value?

There was value realized in consolidating virtualization teams. There was also efficency in collapsing the infrastructure: less compute, storage, load balancers, etc. It also helped with the VM management story, and customers only have to learn one thing instead of learning OpenStack or OpenShift or Kubernetes APIs directly.

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

KubeVirt will continue to be used. The project implementes most features the adopter needs and is very stable. Looking to see Boot from SAN support, which is currently covered by some enhancement proposals. 

## Perception

The adopter indicated that the maintainers are responsive, but did not have an overall view of the community. 

### How are you participating in the project community?

Upstream bugs and feature requests have been submitted in the past, but have not had to do this in sometime. Sometimes also reach out to KubeVirt maintainers.

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

Feature requests and bug reports were submitted and addressed. 

## Project Strengths

### In your opinion, what are the overall strengths of the project?

The ability to treat virtual machines in the same way as containers and serverless workloads, and the ability to bring the Kubernetes ecosystems to virtual machines is the biggest strength. 

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

None identified

### In your opinion, what can the project do better?

None identified
