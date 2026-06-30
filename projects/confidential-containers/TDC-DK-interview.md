Interviewees: Nino Wael(NMWA@tdc.dk), Lasse(nmwa@tdc.nuuway.dk), Martin Rasmussen(maras@tdc.dk)

Interviewer: Lin Sun

Date: Jan 29, 2026

## Questions

#### 1. How long has your organization used the project?

A small branch within the company focusing on IT security. We knew the project for a while, and began seriously checking it out at the beginning of 2025.

#### 2. What were the main motivations to adopt the project and which key features do you use today?

Interested in running AI workload securely, we think CoCo would be a good fit. And running confidential workloads in general following the initial AI workloads. 

#### 3. Compared with other products and projects in this space (proprietary and open) what drew you to the project?

Evaluated the Gramine project prior to using CoCo, it was Intel SGX based. What drew us to CoCo is it supports flexible hardware options beyond Intel (Intel TDX). The project contributors/maintainers are from different hardware vendors which is a big plus.

#### 4. What is the current level of usage (pre-production, production) and scale?

Pre production/POC at the moment

#### 5. What version of the project is currently in use and what is your update cadence with the project?

Try to follow the latest release. Version 18 CoCo. We also try to upgrade following the upgrade path.

#### 6. Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

The project has guides/tutorials on basic things. One of the problems is when combining tutorials/features, they don’t work (additional configurations are needed) and end up looking at source code. Docs are sometimes linking to 404 or stale info.

#### 7. Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

We had to check source code to find info sometimes. People are helpful on slack but had to check source code for what I needed to do (small usability wrapper/web service that uses some endpoints in the project which lack docs.)

#### 8. Did you need to engage with the community members or maintainers? If so what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

Yes we did engage with the community. It is mandatory to be in the community to fully understand what is going on. It is an extremely complex area and has a steep learning curve. People are always very helpful.

#### 9. Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

Since it is only POC, the value isn’t measurable. But compared with Gramine, it is a lot easier to use.

#### 10. If the project were to be archived now or in the future, what level of difficulty would your organization experience to remove it from your environments? If that were to happen, would you fork and maintain the project to keep functionality, step into a maintainership role within the project, or something else?

Nino from the team is already a contributor to the project. Depending on if we have customers, there may not be a need to maintain it. We don’t usually have bandwidth for providing maintenance on projects. If we really need it, we could fork and maintain it, assuming the maintenance value is worth the effort compared with the value provided by the project. 

#### 11. Is there something you feel that holds the project back from reaching its ultimate potential?

Docs as discussed earlier.

The project seems to try to cover too many features,  which could be covered by existing OpenSource products, like authentication/authorization where KeyCloak could do it.

Some features divert from the core purpose (like web interface for Trustee administration)…. 

Could use help from/integrate with other projects in CNCF vs developing all features needed in the project

#### 12. In your opinion, what could the project improve?

The compatibility matrix would be really nice.
Documentations - we are not expecting production grade but it could be a lot better.
A lot of code in the project is removed but still in docs… have to research source code to find out. 

Another thing is that some features are “not” supported but still available in the source code, non supported features (Azure CSI wrapper) should be clearly marked as such (ref. compatibility matrix).

#### 13. What are the overall strengths of the project?

Flexible hardware support and diverse contributing companies and members of the project. A lot of big tech companies are involved in the project.

Contributing was not too hard for new contributors.

Single pane to access confidential container features without different vendor docs/downloading different docs and compose them yourself.

#### 14. Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

We do have plans but we are also customer driven. If we can’t convince customers, we’ll be a passive engagement with the project.
