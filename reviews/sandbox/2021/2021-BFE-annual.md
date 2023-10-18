
# 2021 BFE Annual Review

## Background

[BFE](https://www.bfe-networks.net) is an open-source layer 7 load balancer. BFE was originally created by Baidu in 2012. By the end of 2020, BFE has already processed more than 1 trillion requests per day in Baidu. BFE was officially open sourced in July 2019 and has gained widespread attention. 

**BFE was accepted into CNCF Sandbox On June 25, 2020.**

- [BFE Sandbox Proposal](https://github.com/cncf/toc/pull/453)
- [TOC proposal slide](https://docs.google.com/presentation/d/1vI8cGFbrDWZko22RDKyJF1p3dfMvI_HB-pj8wKiKA_A/edit)


## Annual Review Questions

* Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

  * Commits over the last 6 months: **154**
  * Releases over the last 6 month: **3**
  * Total contributors from various companies: **84**
  * [CNCF DevStats](https://bfe.devstats.cncf.io/)


* How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

  * There are 7 maintainers from 3 different companies.
  * https://github.com/bfenetworks/bfe/blob/develop/MAINTAINERS.md


* What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

  We have made it simple for people to [become listed on our website](https://github.com/bfenetworks/bfe/issues/748). Today, we are aware of the following users: https://github.com/bfenetworks/bfe/blob/develop/ADOPTERS.md

  And more adopters are under confirmation for making their information public.


* How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

  This is our first review, and we didn't state goals in the original Sandbox proposal. In the past year, the project has focused on the following areas:
  * Overhauling the automated test infrastructure, to enable more rapid development from contributors and ensure project quality
  * Adding many key features(gRPC/Fast routing for k8s ingress/Security related modules, etc) to the project 
  * Publishing a book **"[BFE inside](https://github.com/baidu/bfe-book)"** (currently in Chinese) and attracting more users through the increased visibility

  Besides the above metrics we saw, we've also achieved the follow milestones:
  
  * First GA release v1.0.0 on Nov 2020.
  * Participated in Gopher China 2021 and Global Opensource Technology Conference 2021.


* What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

  At the moment, the main focus of the project is enhancing user experience and improving adoption. We are planing to make the components of BFE control plane open source in 2021 and publish English version of  **"BFE inside"** in 2022. We are also planing to support Kubernetes Gateway API in 2021.

  We've also had a number of user requests that we'd like to address, including:

  * Support for HTTP3/TLS1.3
  * Support for traffic mirroring
  * Support for acceleration of RSA based on hardware
  * Better documentation

![bfe-arch](https://raw.githubusercontent.com/bfenetworks/bfe/develop/docs/images/bfe-arch.png)


* How can the CNCF help you achieve your upcoming goals?

  We do wish to get more help from marketing perspective, but it's understandable that the graduated and incubation projects have the priority. We also need some technical writers to help improve documentation.


* Do you think that your project meets the criteria for incubation?

  Yes, we do and the incubation proposal is under preparation.
