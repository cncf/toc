# BFE 2022 Annual Review

## Background

[BFE](https://www.bfe-networks.net/) is an open-source layer 7 load balancer. BFE was originally created by Baidu in 2012. By the end of 2020, BFE was processing more than 1 trillion requests per day in Baidu. BFE was officially open sourced in July 2019 and has gained widespread attention.

**BFE was accepted into CNCF Sandbox On June 25, 2020.**

- [BFE Sandbox Proposal](https://github.com/cncf/toc/pull/453)
- [TOC proposal slide](https://docs.google.com/presentation/d/1vI8cGFbrDWZko22RDKyJF1p3dfMvI_HB-pj8wKiKA_A/edit)

## Annual Review Questions

- Include a link to your project’s devstats page. We will be looking for signs of consistent or increasing contribution activity. Please feel free to add commentary to add colour to the numbers and graphs we will see on devstats.

  - [CNCF DevStats](https://bfe.devstats.cncf.io/)
  
- How many maintainers do you have, and which organizations are they from? (Feel free to link to an existing MAINTAINERS file if appropriate.)

  - There are 11 maintainers from 5 different companies.
  - https://github.com/bfenetworks/bfe/blob/develop/MAINTAINERS.md

- What do you know about adoption, and how has this changed since your last review / since you joined Sandbox? If you can list companies that are end users of your project, please do so. (Feel free to link to an existing ADOPTERS file if appropriate.)

  2022 new adoptions includes companies from different industries, including Internet, Media and Entertainment, Manufacturing, Telecom and Banking.

  We listed some of users here: https://github.com/bfenetworks/bfe/blob/develop/ADOPTERS.md . Some other adoption can be found at  [BFE Users](https://github.com/bfenetworks/bfe/issues/748) . We are also working on confirmation of more adopters before we can make their information public.

  And according to CNCF Survey, we also have new users from other continents like North America and Africa.

- How has the project performed against its goals since the last review? (We won't penalize you if your goals changed for good reasons.)

  In 2022, the project has below achievements:

  - [BFE forward engine](https://github.com/bfenetworks/bfe) released 2 new versions,  which introduce HTTP2 fingerprint, improved TCP KeepAlive configurations and other security, performance and CI/CD efficient improvement. 
  - [BFE Ingress](https://github.com/bfenetworks/ingress-bfe) released 1 new version and improved its deployment.
  - The English version of book **"[In-depth Understanding of BFE](https://github.com/baidu/bfe-book)"** is ready and we are now in communications with some publish houses to prepare the publish.
  - An adopters' success story was published on CNCF Blog: [With BFE, China Merchants Bank has improved application availability and operation efficiency](https://www.cncf.io/blog/2023/01/13/with-bfe-china-merchants-bank-has-improved-application-availability-and-operation-efficiency/)
  - Two Internship programs have been finished with success. With mentoring from our senior contributors, two graduate students has successfully added new functions (Redirect and URL Rewrite) to BFE Ingress. During the program, we followed the best practice of software engineering and payed a lot of effort on requirement analysis, system designing, documentation and reviewing besides coding.

- What are the current goals of the project? For example, are you working on major new features? Or are you concentrating on adoption or documentation?

  At the moment, the main focus of the project is enhancing user experience and improving adoption.  We will also working on supporting new topologies and standards. Main goals includes:

  - Publish English version of book **"In-depth Understanding of BFE"**
  - Support Kubernetes Gateway API
  - Support for HTTP3/TLS1.3
  - Support for traffic mirroring
  - Better documentation

- How can the CNCF help you achieve your upcoming goals?

  We do wish to get more help from marketing perspective. We also need some technical writers to help improve documentation.

- Do you think that your project meets the criteria for incubation?

  Yes, we do and the incubation proposal is under preparation.
