# OpenFGA Adopter Interview - Grafana   

**Meeting Time:** 9/18/2025  
**Meeting Record:** TBD  
**Attendees:**  
- **TOC:** Faseela K, TOCa Sponsor of KServe  
- **TOC Shadow:** Ricardo Aravena  
- **Adopter:** Joao Guerreiro, Senior Engineering Manager, Grafana Labs

---

## Organization Intro  

### Can you give us an overview of your organization and what it does?  
Grafana Labs is the company behind Grafana, Loki, Tempo, Mimir, the open and composable observability stack we call LGTM. Our mission is to help every team observe, understand, and act on their data.  
Grafana Cloud is our fully managed observability platform that brings together metrics, logs, traces, profiles, service-level objectives, and incident response. For teams that prefer to run open source on their own, we maintain a range of projects, including Grafana, Loki, Tempo, Mimir, Alloy, and k6.  

---

## Motivation  

### Compared with other products in this space (proprietary and open), what drew you to the project?  
While we maintained our own access control engine within the Grafana mainline, this engine was designed for single-tenant use and proved unsuitable for the new architecture being implemented in Grafana. Consequently, we faced two alternatives: either porting our existing engine and adapting it for multi-tenancy, or seeking an external solution. We opted to pursue an engine maintained by the broader community, as we determined that independently developing such a solution would largely involve re-inventing existing functionalities, making a community-driven approach more advantageous. A further prerequisite was the open-source nature of the access control engine, given its inclusion with open-source Grafana.  

Two main factors drew us to OpenFGA: multi-tenancy and storage.  
We were looking for more than an evaluation engine (such is the case for OPA); we also needed a storage layer for our permissions. That distinction was central to our decision to consider OpenFGA.  

We were also attracted by OpenFGA’s CNCF affiliation, its open source nature, and its governance policy. We would rather align with a CNCF project and community than contribute to another open-source project capable of changing its license at any moment.  

---

## Usage Scenario  

### How long has your organization used the project?  
Our first experiments with OpenFGA dates to February 2024. We began implementing it in our mainline in August 2024.  

### What were the main motivations to adopt the project and which key features do you use today?  
Today, the core features we rely on in OpenFGA are its store implementation, which is unique in the space and supports a per-tenant schema. We also make use of its check and listing capabilities for access evaluation and to enable permission-based search across dashboards as well as all of our other resources.  

OpenFGA's architecture enables us to run it in the cloud as a multi-tenant service, while also supporting embedded deployments in our open source software and on-premises environments.  

In the future, we want to rely on list users to provide a capability that even our old engine was not able to support, reverse permission search, meaning being able to list the users that have access to a certain resource.  

### What is the current level of usage (pre-production, production) and scale?  
We operate three distinct environments: development, staging, and production.  

Both the development and staging environments currently utilize OpenFGA. Our internal production systems are hosted within the staging environment, thus operating OpenFGA in a production capacity for internal use.  

For our external production environment, OpenFGA is deployed within a single cluster in a pre-production capacity. This setup is used to shadow calls and rigorously validate OpenFGA's consistency and performance guarantees before its broader rollout across the remainder of our production infrastructure.  

### What version is used and what is your update cadence with the project?  
- v1.10.0  
- We watch the community calls often and when we identify interesting new features or CVEs are announced we update.  
- Given that OpenFGA is currently in pre-production, it has not been incorporated into our automated update processes. We are presently conducting benchmarks against OpenFGA, which necessitates a stable environment to avoid variable changes.  
- After OpenFGA updates, we have internal integration tests that verify that access logic has not changed. And in our development environment, we run automated load tests to verify changes in performance.  

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?  
One of our main challenges adopting OpenFGA was the disparity in performance between different databases. The MySQL data was not as mature as the Postgres adapter, which forced us to change our backing database from MySQL to Postgres. Our legacy schema was also very complex, which hindered performance. We have now improved this significantly, thanks in part to reasoning about access control as a schema.  

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?  
Yes, [https://openfga.dev/docs/fga](https://openfga.dev/docs/fga) is very complete and we’ve used it a lot (consistency, adoption patterns, …)  
[https://openfga.dev/docs/best-practices/adoption-patterns](https://openfga.dev/docs/best-practices/adoption-patterns) The adoption patterns, even though it came after, is a good reminder of good practices to follow.  
[https://openfga.dev/docs/best-practices/source-of-truth](https://openfga.dev/docs/best-practices/source-of-truth) The source of truth document has actually led us to review some of our data storage and decisions regarding access architecture for the better.  

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.  
Yes. Moving from a single-tenant monolith to a multi-tenant architecture has saved us substantial implementation time. We faced two paths for achieving a multi-tenant setup: migrate our current engine and tables, or pursue an alternative approach. In practice, we have been able to benefit from an open-source community project that offers stronger support than our internal resources would allow. We also expect this approach to yield cost savings, since a centralized multi-tenant architecture generally reduces overhead compared with a dispersed setup.  

The ease of use in changing schema has allowed us to more quickly iterate in finding one that suits our product. And as well, the project has overall decreased the manual activities such as engine tuning that we would regularly do.  

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.  
We have already given a talk with OpenFGA (*From Chaos To Control: Migrating Access Control... Jo Guerreiro & Poovamraj Thanganadar Thiagarajan Kubecon EU 2025*), and we expect to do so again in the future. We are also exploring joint blog posts to highlight best usage practices.  

Our team is heavily invested in the maintenance of the SQLite adapter. Moving forward, we intend to continue reviewing and supporting OpenFGA as its storage options expand, and to contribute to maintaining performance at the storage layer. A specific project we aim to collaborate in is the development of pluggable storage, which would enable non-core storage adapters to function with OpenFGA without requiring rebuilds.  

Additionally, we may explore improvements for observability within OpenFGA.  
We have one maintainer listed as part of the core team of OpenFGA: SECURITY-INSIGHTS.yml and we’ve had 3 other members of our team contribute changes to OpenFGA.  

---

## Perception  

### What is your perception in terms of the project’s:  
- **Community openness**  
- **Governance**  
- **Community growth potential**  
- **Maintainer diversity and ladder**  
- **Maintainer response**  

**Maintainer Diversity and ladder:** OpenFGA is an Okta-heavy project, it is true. But we've found that they’ve gone above and beyond to address it and invite us to contribute. In our own OSS projects such as Grafana, we find it hard to bring external contributors to the authorization and authentication space because of the high complexity of these systems, the maintenance burden they represent, and as well the possible security risks or state actor attacks.  

**Governance:** We follow the roadmap posted by OpenFGA contributors. We interact with the different GitHub issues that interest us. But we also have a monthly maintainer call where we discuss the medium-term direction of the project and what are the different contributor priorities.  

**Community openness:** They respond quickly and sometimes reach out proactively, offering advice or inviting us to try a new feature. The community calls are always well-staffed, and they regularly call for feedback on various proposals.  

**Community Growth Potential:** I see substantial growth potential for the community as this becomes a more common issue for other users. We’ve seen Airbnb and GitHub have also moved toward a ReBAC approach. The growing significance of authorization underscores its increasing importance, and I am confident that OpenFGA stands to gain considerably from the ongoing standardization efforts within this domain.  

### How are you participating in the project community?  
We maintain the SQLite adapter, which leads us to frequent discussions with maintainers in the OpenFGA/CNCF Slack. We take part in community meetings and join the maintainers' monthly call to explore potential directions for the project's future and the challenges it faces.  

GitHub issues/feature requests as well.  

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?  
One of our early engagements involved adding SQLite support so it could be used in Grafana. The maintainers were willing to take on this new code as long as we owned and maintained it. We felt that was fair and proceeded accordingly. They moved quickly on reviews, and we kept the conversation open through Slack as well as by arranging synchronous calls. In the end, the SQLite adapter was merged and is now being used in OpenFGA.  

---

## Project Strengths  

### In your opinion, what are the overall strengths of the project?  
The usage of the ReBAC model and backing from Okta make it a strong candidate to become a standard in authorization.  
The project offers solid documentation, and its structure is designed to support large-scale deployments rather than only experimental or playground deployments.  

---

## Project Improvements  

### Is there something you feel that holds the project back from reaching its ultimate potential?  
I believe the project is currently held back by insufficient production time for many community members to develop best practices and contribute fine-grained improvements at a level that matches the production experience of the current maintainers.  
After delivering the talk at KubeCon EU 2025, I believe the lack of adopter stories on OpenFGA's website makes larger adopters hesitant to try the solution, even though there are many interested, as a lot of larger organizations came to talk with me at the end of the talk to know our experience.  

### In your opinion, what can the project do better?  
I believe the project should pursue additional benchmarking and direct, head-to-head comparisons with other alternatives in the space. This work will help carve out a clear niche and, at the same time, demonstrate that it can support large-scale deployments and run stable in production. Such evidence will strengthen confidence among stakeholders evaluating adoption.  
Improving maintainer diversity in the long run will allow establishing processes that make governance more diverse and targeted to a wider array of use cases.  
