## OpenFGA Adopter Interview - Read AI

**Meeting Time:** 8/9/2025  
**Meeting Record:** TBD  

**Attendees:**  
- **TOC:** Faseela K, TOC Sponsor of KServe  
- **TOC Shadow:** Ricardo Aravena  
- **Adopter:** Andrew Powers, Software Engineering Manager, Read AI

### Organization Intro
### Can you give us an overview of your organization and what it does?  
Read AI is the fastest-growing meeting notetaker and AI assistant in the world, acting as a store of intelligence across meetings, messages, email, documents, and every workplace touchpoint. Based in Seattle and trusted by more than 100,000 organizations and 75% of the Fortune 500, the platform is adding over 1 million new customers every month who value its commitment to transparency and security based on robust, scalable authorization.  

### Motivation
### Compared with other products in this space (proprietary and open), what drew you to the project?  
- Considered alternatives such as Authzed.  
- Previously used a proprietary, organically built authorization solution, but faced performance and scalability issues.  
- Sought a solution with stronger sharing capabilities, standard best practices, and open source flexibility.  
- OpenFGA stood out for:  
  - Clear and detailed documentation.  
  - Based on Google Zanzibar, which aligned well with their desired product functionality.  
  - Maintainers who were approachable, supportive, and responsive.  
  - Ability to self-host.  
  - Cost-effectiveness and performance compared to proprietary offerings.  

### Usage Scenario
### How long has your organization used the project?  
- Began evaluation in February 2023.  
- Running in production since April 28, 2023.  

### What were the main motivations to adopt the project and which key features do you use today?  
- Adoption motivated by the need for scalable, reliable authorization with strong sharing semantics.  
- Using most of the core features.  
- Integration with Postgres for persistence.  

### What is the current level of usage (pre-production, production) and scale?  
- **Production.**  
- Self-hosted OpenFGA service handles **5,200 requests/sec under peak load** with **20ms p99 latency** and **1.8ms average latency**.  
- Data store holds **5,323,283,829 tuples** and is growing daily.  

### What version is used and what is your update cadence with the project?  
- Currently using v1.8.16.  
- Internal cadence: monthly updates.  
- OpenFGA release cadence is faster than Read AI’s adoption cadence, but upgrades have been smooth with no significant backward compatibility issues.  

### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?  
- Aligning internal data models with OpenFGA’s authorization model was the biggest challenge.  
- Migration from proprietary system required careful design and data imports, but subsequent upgrades have been straightforward.  
- Performance at peak load has been excellent; OpenFGA has never been a bottleneck.  

### Did you find the information in the repo or the project docs valuable to your implementation? If so, what did you find useful?  
- Yes, the documentation was key to adoption.  
- Practical examples and detailed modeling guides were especially useful.  
- Documentation on bulk tuple import was referenced during implementation.  

### Has your implementation of the project provided measurable value?  
- Improved confidence in secure data authorization.  
- Adoption of best practices improved internal design decisions.  
- **Cost savings:** reduced compute and hosting costs due to OpenFGA’s optimized, self-hosted performance.  
- **Reliability:** system performance has been stable even at high loads.  

### Do you have any future plans regarding the project?  
- Intend to submit a case study and potentially a conference talk on the adoption experience.  
- Currently not contributing upstream but open to more involvement in the future as the engineering team expands.  

### Perception
### What is your perception in terms of the project’s:  
- **Community openness:** The community is active, approachable, and supportive. Reported issues and feature requests receive timely and constructive responses.  
- **Governance:** Limited visibility into governance structure. While functional, greater transparency would be beneficial.  
- **Community growth potential:** Strong growth observed, with adoption increasing across diverse companies, including large enterprises.  
- **Maintainer diversity:** Currently concentrated within a single company. Broader diversity is needed for long-term sustainability.  
- **Maintainer response:** Maintainers are responsive and effective, with bugs fixed quickly and feature requests handled thoughtfully.  

### How are you participating in the project community?  
- Attend community meetings.  
- Reported bugs that were promptly fixed by maintainers.  
- Actively track project progress and updates.  

### Did you need to engage with the community members or maintainers? If so, what was the context and outcome?  
- Yes, engagement via community channels was productive.  
- Maintainers responded quickly to bugs and feature requests, with high-quality outcomes.  

### Project Strengths
### In your opinion, what are the overall strengths of the project?  
- Excellent documentation (clear, practical, and example-driven).  
- Strong project focus (solves a specific problem very well).  
- High performance and reliability in production.  
- Support and backing from larger companies in the ecosystem.  
- Engaged and responsive community.  

### Project Improvements
### Is there something you feel that holds the project back from reaching its ultimate potential?  
- Maintainer diversity—project governance is still concentrated in a single company.  

### In your opinion, what can the project do better?  
- Broaden the maintainer base and governance to increase sustainability and community trust.
