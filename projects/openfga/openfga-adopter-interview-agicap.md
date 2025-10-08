# OpenFGA Adopter Interview - Agicap

**Meeting Time:** 9/8/2025  
**Meeting Record:** TBD  

**Attendees:**  
- TOC: Faseela K, TOC Sponsor of KServe  
- TOC Shadow: Ricardo Aravena  
- Adopter: Pauline Jamin, Head of Engineering – Finance and Core, Agicap (pauline.jamin@agicap.com)  

---

## About the Adopter Interview
The intent of the interview is to ascertain the maturity and adoption of the OpenFGA project by adopters of the project. It’s OK to be public or private.

TOC will record raw notes here, then the notes will be cleaned up and shared with you for your review, correction, and final approval. Once approved, they’ll go into OpenFGA’s due diligence document.

---

## Organization Intro

### Can you give us an overview of your organization and what it does?
Agicap is a European fintech company specializing in cash flow management. Its SaaS platform enables small, medium, and large enterprises to gain real-time visibility and control over liquidity.  

Agicap serves over 8,000 customers across industries. Pauline Jamin, Head of Engineering for Finance and Core, is responsible for technical vision of core financial services and leads the team that designed and implemented Agicap’s Identity and Access Management (IAM) system.

---

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?
Need for open-source authorization with:  
- Flexibility and extensibility.  
- Strong, responsive community.  
- On-prem deployment to meet compliance/security requirements.  

Evaluated alternatives (e.g., Oso – not fully open source). OpenFGA was more stable and aligned with requirements.  
- Strong documentation and approachable maintainers.  
- ReBAC (relationship-based access control) model provided more flexibility than RBAC.  

---

## Usage Scenario

### How long has your organization used the project?
We began experimenting with OpenFGA in October 2022. By April 2023, we had deployed it to production as the core authorization service for our SaaS platform. Since then, it has been in continuous production for over two years.

### What were the main motivations to adopt the project and which key features do you use today?
Motivation: A robust and evolvable authorization layer.  

Key features in use:  
- ReBAC model for fine-grained permissions.  
- Conditional access rules.  
- Consistent authorization enforcement across all backend services.  

### What is the current level of usage (pre-production, production) and scale?
OpenFGA has been used in production as our main access management system since April 2023, and it is also integrated across all our pre-production environments, including development, preproduction, and load testing.  

All of our customers—currently over 8,000—rely on OpenFGA via our SaaS platform, as every backend service must validate its authorization rules through our access management service. This means all parts of our backend ecosystem are impacted by OpenFGA.  

We have about **250 RPS in production**. 

### What version is used and what is your update cadence with the project?
We aim to stay on the latest available release of OpenFGA. Typically, we upgrade within a week of a new release. All updates are validated with downstream automated tests, which focus heavily on performance regressions.  

In some cases, we contribute fixes upstream and then adopt the patched release quickly.

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project? 
The migration process was non-trivial, as we moved from an internally built RBAC solution to OpenFGA’s ReBAC model.  

The initial integration was straightforward because our authorization model was simple, but as usage and data volume grew, we encountered performance challenges. To address these, we refactored our authorization model into a flatter structure, which improved both scalability and query performance.  

We also faced a security challenge: OpenFGA’s APIs are not protected by authorization rules out of the box. To mitigate this, we built a secure service facade. All our backend services interact with OpenFGA exclusively through this intermediary, allowing us to enforce security and avoid exposing OpenFGA directly.  

On performance, we sometimes hit bottlenecks, particularly around reads. We suggested enhancements such as read replicas, and many of the performance optimizations we needed were reported upstream as issues. In every case, the maintainers responded quickly; one critical performance bug was fixed within 24 hours of our report.

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?
Yes, we relied heavily on the official documentation throughout our integration process.  

It was especially helpful in the early stages to understand how to design our authorization model and later on for configuration guidance. The documentation is very well structured and clear, making it much easier to get started and to troubleshoot as our use case evolved.

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.
The main value comes from flexibility and efficiency.  

- The ReBAC model allows us to express complex, fine-grained permissions without reinventing our authorization logic.  
- Adding or evolving permissions has become significantly easier, which saves engineering time.  
- Because OpenFGA is self-hosted, we retain control over our infrastructure and achieve cost savings compared with proprietary alternatives.  

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.
Yes. Agicap is already an active contributor.  

- Several engineers from my team and our SRE team regularly open pull requests to fix bugs or improve performance.  
- We also participate in community discussions and conferences.  

Looking ahead, we would like to see features such as **idempotent batch writes** for more efficient relationship inserts.

---

## Perception

### What is your perception in terms of the project’s:

**Community Openness**  
The OpenFGA community is viewed very positively. The engineering team participates in monthly community meetings and contributes through pull requests.  

The CNCF Slack channel is considered a valuable resource, with an active, responsive, and approachable community. The project was initiated by passionate contributors and is growing quickly, with increasing adoption.

**Governance**  
Governance of the project is evolving. Most maintainers are currently from a single company (Okta).  

While contributions are regular, the team does not currently have the bandwidth to step into maintainer roles. Okta’s willingness to take responsibility for stewardship is appreciated and seen as important for the project’s stability.

**Community Growth Potential**  
The OpenFGA community has significant growth potential.  

Contributors and adopters come from multiple countries, and momentum is building as more organizations deploy OpenFGA in production environments. This trajectory suggests strong sustainability.

**Maintainer Diversity and Ladder**  
Maintainer diversity is still limited, with most maintainers concentrated in a single organization.  

While this concentration provides continuity and clear ownership, expanding the maintainer base across more organizations is considered valuable over time.  

Geographically, the community is diverse, but demographic diversity—particularly gender representation—is limited, reflecting broader trends in the tech industry.

**Maintainer Response**  
Maintainers are consistently responsive. Issues reported by the team have been acknowledged and addressed quickly, including a performance-related problem resolved by the upstream team within a single day.  

This responsiveness reinforces confidence in the project’s maturity and reliability.

---

## Community Participation

### How are you participating in the project community?
We contribute via:  
- GitHub issues and PRs  
- CNCF Slack  
- Co-presented talks at **KubeCon** with maintainers:  
  - [KubeCon EU 2024, Paris](https://www.youtube.com/watch?v=dummy-eu2024-link)  
  - [KubeCon NA 2024, Salt Lake City](https://www.youtube.com/watch?v=dummy-na2024-link)  

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?
Yes, regularly.  

We have engaged maintainers through:  
- GitHub issues  
- Public Slack  
- A private CNCF Slack channel for closer collaboration  

All interactions have been positive, and maintainers have consistently been responsive and helpful.

---

## Project Strengths

### In your opinion, what are the overall strengths of the project?
- Robust ReBAC model  
- Reliability at production scale  
- Strong vision from maintainers for the future of access management  
- Engaged and responsive community  
- High-quality documentation  

---

## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential? 
- Lack of **idempotent batch writes** (currently requires manual handling of failures).  
- Maintainer diversity remains limited, though Okta has provided excellent stewardship so far.  

### In your opinion, what can the project do better?
- Highlight adoption stories more prominently.  
- Publish a consolidated **adopter list** on the OpenFGA website to demonstrate maturity and drive further adoption.  
