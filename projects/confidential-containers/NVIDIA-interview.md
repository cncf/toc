Interviewer: Lin Sun

Interviewee: Dan Middleton (dmiddleton@nvidia.com)

Date: Jan 23, 2026

## Questions

#### 1. How long has your organization used the project?

Been involved with the project for 2 years. We are maintainers on the project, we are a vendor for the project, along with internal adopters of the project. 
[Confidential computing](https://en.wikipedia.org/wiki/Confidential_computing) is a feature for CPUs and GPUs (carve out a protected memory space, such as creating a virtual machine so you don’t trust other tenants on the same host) , and confidential container (CoCo) is a project which helps end users activate Confidential Computing.

#### 2. What were the main motivations to adopt the project and which key features do you use today?

We have 3 motivations: 
- Contributor: We want the NVIDIA GPU to be enabled for this project, so anyone who wants to use AI can use this security feature provided by the project. 
- Vendor: We are building products that are on top of the CoCo ecosystem.  For example, Redhat openshift is a downstream provider for CoCo.
- Adopter: NVIDIA operates a number of internal and external datacenters. Each has security requirements that benefit from Confidential Computing and we like the cloud-native approach provided by CoCo.

#### 3. Compared with other products and projects in this space (proprietary and open) what drew you to the project?

This is really the main project for cloud-native confidential computing. There are a few other projects but project momentum/feature/contributors drew us. Contributors from major cloud providers, hardware vendors, and independent software vendors.

#### 4. What is the current level of usage (pre-production, production) and scale?

Internally pre-production. Externally for other companies to adopt the NVIDIA version, early access version.  GA is planned for later this year (end of Q1).

#### 5. What version of the project is currently in use and what is your update cadence with the project?

We are on the latest version, 0.18. We try to update pretty close to each formal release of the project. For teams using downstreams like Openshift, they update according to the downstream cadence.

#### 6. Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Because NVIDIA employs some maintainers on the project, it was pretty straightforward. Some challenges involve other CNCF components… making sure we have an agreed version of containerd for example.  These integration points have been improved since.

#### 7. Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

Good docs throughout the repo. In the last year, there is a single website that pulls docs from multiple places.

#### 8. Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

Open community mtg on thurs and we interact live with people there. We also regularly make github issues and the project slack.

#### 9. Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

Fundamental security improvement, faster integrations, reducing manual / custom integrations. Without CoCo there isn’t a consistent cloud-native activation for Confidential Computing. An alternative is running kubelet inside a custom Confidential VM, but this doesn’t satisfy all the security objectives provided by CoCo.

#### 10. If the project were to be archived now or in the future, what level of difficulty would your organization experience to remove it from your environments? If that were to happen, would you fork and maintain the project to keep functionality, step into a maintainership role within the project, or something else?
We’ll continue to maintain the project or fork since we are already maintainers of the project. Because all major hardware vendors support Confidential Computing it is difficult to imagine a future where we abandon the cloud-native stack for it.  

#### 11. Is there something you feel that holds the project back from reaching its ultimate potential?

We want more end users feedback. More visibility always helps.

#### 12. In your opinion, what could the project improve?

Ease of use, almost nothing in k8s is easy to use :) When it comes to security, details matter. The user is supplied with reasonable security defaults but for their specific requirement they have to tweak the configuration. It would be great if the project can make the configuration much easier.

#### 13. What are the overall strengths of the project?

Security strength along with contribution diversity.  In many projects security is an after thought. In CoCo security is a primary requirement. 
Additionally, the variety of hardware, cloud, and software vendors participating as maintainers gives us confidence that the project is well balanced and thoughtfully built. 

#### 14. Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.
We’ll continue to add features to our upcoming products through CoCo, such as multiple GPU support for LLMs.

---
### Maturity Level Survey

The following set of questions goes beyond the scope of the specific project adoption.
Their goal is to benefit from having access to CNCF project adopters and survey their understanding of the CNCF project maturity levels and any reliance on their meaning for decision making. This information should allow the TOC to better document maturity levels targeting adopters.

#### 1. Do you feel you have a good understanding on the meaning of each maturity level for CNCF projects?

Yes I understand and worked with other projects going through the maturity levels.

#### 2. Is there information missing regarding the meaning of each different level?

No.

#### 3. Do you rely on those levels internally in any way, and if yes how?

Indirectly. We are less likely to rely on a project that is early in the maturity level. Our primary decision is based on the project features and implicitly our awareness of alternatives.
