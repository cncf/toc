Interviewee: Rahul Jadhav(nyrahul@gmail.com)

Interviewer: Lin Sun

Company: AccuKnox

Date: 10th Feb 2026

## Questions

#### 1. How long has your organization used the project?

AccuKnox worked on 5G security solution that provided Intent based security for 5G ORAN and 5G Core deployments.
Here is the [sample policy](https://github.com/5GSEC/nimbus/blob/main/examples/clusterscoped/coco-workload-si-sib.yaml) created 
that used CoCo for ensuring that the workload is installed on hardware enabled security module.

#### 2. What were the main motivations to adopt the project and which key features do you use today?

5G UDR (Unified Data Repository) is a network function that stores sensitive data such as users' subscription keys and has to be 
secured intricately. 5G deployments now readily deploy on public cloud providers and towards this end, the security threat model 
is completely changed i.e., we have to assume that the CSP has root/admin access to the virtual machines. 
AccuKnox depends on CoCo based deployments to secure such workloads.

#### 3. Compared with other products and projects in this space (proprietary and open) what drew you to the project?

CoCo is the only project that has the full TEE framework towards containerized deployments, including cloud APIs adapter that 
make deployment even in cloud hosted env easy.

#### 4. What is the current level of usage (pre-production, production) and scale?

Pre-Production

#### 5. What version of the project is currently in use and what is your update cadence with the project?

We are still using v0.13.0 release of CoCo which is roughly an year old. We haven’t diligently updated the [code base](https://github.com/5GSEC/5g-blueprint-controls).

#### 6. Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

We initially struggled to integrate with the AWS version of the CoCo. However, the [Cloud-api-adapter tooling](https://github.com/confidential-containers/cloud-api-adaptor) made it simpler.

#### 7. Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

The docs were useful, however, we had to reach out to slack for help. Pradeepta and a few other folks helped out earnestly.

#### 8. Did you need to engage with the community members or maintainers? If so, what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

I presented the work we did in 5Gsec to the CoCo community and that's when they advised us to inform as part of adopters list.

#### 9. Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

Measurable value is improved security posture. There is no alternative to TEE/CoCo based execution for certain threat models 
wherein if one cannot trust the cloud platform provider than the only hope left is to deploy using CoCo the sensitive workloads.

#### 10. If the project were to be archived now or in the future, what level of difficulty would your organization experience to remove it from your environments? If that were to happen, would you fork and maintain the project to keep functionality, step into a maintainership role within the project, or something else?

Even though we do not have any hard-bound dependencies after the 5Gsec project, I believe that a certain threat model cannot 
be handled without the use of CoCo.

#### 11. Is there something you feel that holds the project back from reaching its ultimate potential?

The hardware availability and stability around the interfaces. The cost of hardware is still prohibitively high to be deployed in general scenarios.

#### 12. In your opinion, what could the project improve?

Messaging. The project needs to create awareness around the use-cases. People get bogged down with the gory technical details.

#### 13. What are the overall strengths of the project?

It handles a security threat model that cannot be handled otherwise without the use of hardware enabled security and the project 
does wonders for enabling it in the containerized deployment.

#### 14. Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

Yes, we want to make it formally as part of AccuKnox Enterprise offerings. The gov/federal depts use-cases are great to be handled using this project.

---
### Maturity Level Survey


#### 1. Do you feel you have a good understanding on the meaning of each maturity level for CNCF projects?

I believe so (but not very strongly). I have been part of a similar journey for another CNCF project (KubeArmor) for which I am the maintainer. 

#### 2. Is there information missing regarding the meaning of each different level?

I have read through the information regarding what is required for changing from sandbox to incubation maturity. 
Similarly for graduation phase. 

#### 3. Do you rely on those levels internally in any way, and if yes how?

Yes, within AccuKnox we want to use Graduated projects only as far as possible. We leverage sandbox or incubation projects only 
if we have enough competency inhouse to handle the issues.
