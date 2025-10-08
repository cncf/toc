# OpenFGA Adopter Interview – Zuplo  

**Meeting Time:** 9/18/2025  
**Meeting Record:** TBD  

**Attendees:**  
- **TOC:** Faseela K, TOC Sponsor of KServe  
- **TOC Shadow:** Ricardo Aravena  
- **Adopter:** Nate Totten, Co-founder & CTO, Zuplo 

## About the Adopter Interview  
The purpose of this interview is to assess the maturity and adoption of the OpenFGA project by its adopters. The content may be shared publicly or kept private as per the adopter's preference. The TOC will record raw notes here, which will then be cleaned up and shared with the adopter for review, correction, and final approval. Once approved, they will be included in OpenFGA’s due diligence document.  

## Organization Overview  

### Can you give us an overview of your organization and what it does?  
Zuplo is a developer-first API management platform that enables teams to build, deploy, and scale APIs efficiently. The platform offers a fully-managed, lightweight solution designed for developers, providing fast deployment, GitOps-friendly workflows, and unlimited preview environments.  

## Motivation  

### Compared with other products in this space (proprietary and open), what drew you to the project?  
Zuplo evaluated several options for implementing fine-grained authorization, including:  
- Axiomatics (AuthZEN-based)  
- Aserto (AuthZEN-based)  
- OktaFGA  
- Building a custom solution in-house  

The decision to adopt OpenFGA was influenced by its open-source nature, allowing for self-hosting and cost-effective scaling. The ability to use PostgreSQL as the backend was particularly advantageous, enabling global data replication to perform authorization at the edge, aligning with Zuplo's requirements.  

## Usage Scenario  

### How long has your organization used the project?  
Zuplo has been utilizing OpenFGA for approximately 18 months, with 12 months in production.  

### What were the main motivations to adopt the project and which key features do you use today?  
As Zuplo expanded its customer base to include larger enterprises, the initial simple authorization mechanisms (e.g., project membership) proved insufficient. The need for Role-Based Access Control (RBAC) became evident. OpenFGA was adopted to implement relationship-based access control, enforcing access to various features within Zuplo's product. Key features in use include:  
- A single authorization model applied across the entire product  
- Management of both user-level access and API key access to product features  
- Single-tenant deployment with one model governing the entire product  

### What is the current level of usage (pre-production, production) and scale?  
OpenFGA is deployed in production across multiple data centers worldwide. The system handles several hundred requests per second, with occasional spikes exceeding 500 requests per second.  

### What version is used and what is your update cadence with the project?  
Zuplo is currently using OpenFGA version 1.8.x. Upgrades are performed every couple of months or as needed for security patches. The upgrade process involves running OpenFGA's tests to ensure the model behaves correctly, along with full integration tests to verify functionality within Zuplo's product. Performance evaluations are conducted using k6 for major upgrades, focusing on end-to-end testing rather than direct benchmarking of OpenFGA.  

### Can you walk me through your experience in adopting or integrating the project? What challenges did you experience?  
Adopting OpenFGA was relatively straightforward. The primary challenge lay in migrating existing systems to utilize OpenFGA. Understanding and applying Zuplo's permission design to OpenFGA's model required learning and adaptation. Early versions of OpenFGA lacked certain features, such as permission listing APIs, which necessitated workarounds. However, subsequent releases addressed these gaps. The absence of caching in initial versions led Zuplo to implement its own caching mechanisms. Once OpenFGA introduced caching, performance improved significantly. Overall, the roadmap was clear, and most initial issues were resolved in later versions.  

### Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?  
Yes, the documentation on the OpenFGA website was instrumental. It was easy to understand and comprehensive. The examples provided, particularly those related to advanced features, were directly applicable to Zuplo's use case. Notably, the Google Drive and GitHub examples were particularly useful.  

### Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.  
Yes, the implementation has streamlined Zuplo's authorization processes. The ability to update and version the authorization model independently of the codebase has facilitated faster development cycles. Testing and refining roles and permissions without code changes has centralized authorization management, allowing teams to move faster and maintain separation of concerns.  

### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.  
Zuplo plans to continue using OpenFGA and periodically submits feature requests. The team has also expressed interest in sharing their experiences by giving talks on how they implement authorization at the edge using OpenFGA.  

## Perception  

### What is your perception in terms of the project’s:  
- **Community openness:** Zuplo has had positive experiences with OpenFGA's community. Communication has been responsive, and feedback is welcomed.  
- **Governance:** Zuplo has observed that OpenFGA is primarily maintained by Okta, which initially raised concerns about vendor neutrality. However, the project's openness and responsiveness have alleviated some of these concerns.  
- **Community growth potential:** While the community appears active, Zuplo notes that there may be opportunities to further foster growth and engagement.  
- **Maintainer diversity and ladder:** The concentration of maintainers from Okta suggests a need for broader contributor diversity to enhance the project's neutrality.  
- **Maintainer response:** Zuplo has experienced timely and helpful responses from maintainers, contributing to a positive experience.  

### How are you participating in the project community?  
Zuplo primarily engages with the OpenFGA community through GitHub issues.  

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use, and did it reach an acceptable outcome?  
Engagement with maintainers, particularly Andreas, has occurred via email, GitHub, and Discord. These interactions have been productive and have led to satisfactory outcomes.  

## Project Strengths  

### In your opinion, what are the overall strengths of the project?  
- A strong corporate sponsor ensures adequate support.  
- A growing and active community contributes to the project's development.  
- Comprehensive documentation and examples facilitate implementation.  
- Responsive maintainers address issues and feedback promptly.  

## Project Improvements  

### Is there something you feel that holds the project back from reaching its ultimate potential?  
Zuplo stated they are pretty much satisfied with everything and have no comments.  

### In your opinion, what can the project do better?  
Zuplo stated they are pretty much satisfied with everything and have no comments.  
