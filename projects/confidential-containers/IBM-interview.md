Interviewer: Lin Sun

Interviewee: Nicolas Mäding (NMAEDING@de.ibm.com)

Date: March 2, 2026

## Questions:

#### 1. How long has your organization used the project?

We started to actively contribute and participated in and promoted the project since it was formed in 2021. We had some internal PoTs and PoCs based on the Confidential Containers Project (https://github.com/confidential-containers)
since then.

#### 2. What were the main motivations to adopt the project and which key features do you use today?

From my point of view this is the most mature CNCF project in this area and has key industry participants contributing and collaborating. We primarily use the s390x and IBM Secure Execution for Linux versions of the project as well as the
Trustee component.

#### 3. Compared with other products and projects in this space (proprietary and open) what drew you to the project?

It has the most references and contributions by various partners and technologies providers, which makes it for me the reference project around Confidential Computing in the CNCF context. There are different products also
regarding non-CNCF based deployments - Kernel-based Virtual Machines (KVM) and Public Cloud (VSIs) - and I believe based on my experience and customer feedback that a community driven project is key for adoption of this technology
in production.

#### 4. What is the current level of usage (pre-production, production) and scale?

In my opinion, production adoptions depend on the availability of additional features like Baremetal support and an established common Trustee flow. This is
work in progress within the project; the TechPreview has been out since 4Q25 and we intend to harden and further improve its design to enable more usecases.

https://www.ibm.com/docs/en/ccro/1.1.x?topic=platform-confidential-computing-linuxone
https://www.ibm.com/docs/en/rhocp-ibm-z?topic=hyper-protect-confidential-container-red-hat-openshift-container-platform-bare-metal-beta

#### 5. What version of the project is currently in use and what is your update cadence with the project?

We are using 0.16.0 and contributing to main. We intend to update the latest version as soon as it becomes available as it is the foundation for our Kubernetes
platform our product is built upon.

#### 6. Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

As we intend to adopt the project for Linux on IBM Z and LinuxONE (s390x architecture), the key challenge from my point of view is the focus on Intel/AMD
(x86 architecture) and adoption by Public Cloud Providers rather than on Premises datacenter.

So different adopting and enabling use cases on different platforms leads to different architectural decisions and adoption of changes. The project from my point of view is too focused on getting the Public Cloud adoption progress and
making decisions focused on this deployment model. I acknowledge that x86 is widespread, but a flexible architecture taking different approaches and deployments into account fosters in my opinion innovation and better solutions
for many use cases to come.

#### 7. Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

The available documentation is sufficient and fine. High-level technology independent architectural pictures are useful in this context.
As this project matures further, thread model documentation will be an important value-add to be made available through the community.

#### 8. Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

We have several team members who are maintainers on the project, so we are in communication with the wider community on a near-daily basis via slack, github and the regular Confidential Containers zoom meetings.

#### 9. Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

Yes, this project provides a common approach and project across different hardware and system architectures, which has values in my opinion for many as they get an ease-of-use solution with this project across their hybrid cloud
datacenter.

#### 10. If the project were to be archived now or in the future, what level of difficulty would your organization experience to remove it from your environments? If that were to happen, would you fork and maintain the project to keep functionality, step into a maintainership role within the project, or something else?

The impact would be significant, and we would need to investigate options, if that would become a risk.

#### 11. Is there something you feel that holds the project back from reaching its ultimate potential?

Not at this point.

#### 12. In your opinion, what could the project improve?

Be more open to other technologies, processor architectures and deployment models.

#### 13. What are the overall strengths of the project?

As stated, the contributions from so many participants across the industry and technical community.

#### 14. Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

We intend to contribute, request features for different solutions, improved user experience, and plan to continue our involvement as products depend on this project and the underlying technology.

---
## Maturity Level Survey

The following set of questions goes beyond the scope of the specific project adoption. Their goal is to benefit from having access to CNCF project adopters and survey their understanding of the CNCF project maturity levels and any reliance on their meaning for decision making. This information should allow the TOC to better document maturity levels targeting adopters.

#### 1. Do you feel you have a good understanding on the meaning of each maturity level for CNCF projects?

Yes.

#### 2. Is there information missing regarding the meaning of each different level?

Not at this point.

#### 3. Do you rely on those levels internally in any way, and if yes how?

Indirectly. The maturity of the projects used are foremost part of our internal risk assessment and are beneficial as this is proof that a given project and/or concept is accepted in the community and industry at large.
