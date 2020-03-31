# Keycloak CNCF Sandbox Project Submission Proposal

This is a resubmission of Keycloa as Sandbox project. Initially proposed in 2018 has been impacted by Sandbox process changes and overal lack of bandwidth in TOC. This resulted in halt of intake of new projects. Now that new process settled and new TOC elections concluded we would like to approach inclusion of Keycloak as CNCF Sandbox project again. 


Authors:  
Bolesław Dawidowicz https://twitter.com/bdawidowicz  
Stian Thorgersen https://github.com/stianst

# Background
Link to TOC PR: This

Link to Presentation: (Oct 2018 TOC presentation - slide 26) https://docs.google.com/presentation/d/1Xt1xNSN8_pGuDLl5H8xEYToFss7VoIm7GBG0e_HrsLc/mobilepresent?slide=id.g3f805096e4_212_0

Link to GitHub project: https://github.com/keycloak

Getting Started / Trying out: https://www.keycloak.org/getting-started

CNCF SIG Security assesment request: https://github.com/cncf/sig-security/issues/372

CNCF SIG Security Self Assesment document: https://docs.google.com/document/d/14IIGliP3BWjdS-0wfOk3l_1AU8kyoSiLUzpPImsz4R0/edit#

## Goal
Keycloak is an open source Identity and access management solution aimed at modern applications and services. It makes it easy to secure applications and services with little to no code. 

Keycloak is based on standard protocols with an aim toward modern use cases and the flexibility to integrate with other solutions and prevent vendor lock in. Supported protocols include: OAuth2, OpenID Connect, User Managed Access 2.0 (UMA) and SAML 2.0. 

### Single-Sign On

Users authenticate with Keycloak rather than to individual applications. This frees developers from dealing with login forms, authenticating users, and storing users. Keycloak provides both single sign-in and single sign-out to different types of applications (client-side, server-side, mobile, etc.) and protocols (OIDC and SAML2) within same realm. It also provides Kerberos bridge to propagate authentication (LDAP or Active Directory) from workstations to web applications via SPNEGO. 

### Token based security

Through OpenID Connect flows applications retrieve access tokens that can be used to invoke services securely without the need to pass user credentials. This enables easy end-to-end user authentication for client-side applications, micro services, service mesh and other modern architectures.

### Identity Brokering and Social Login

Keycloak can authenticate users using any existing external OpenID Connect or SAML 2.0 Identity Providers by acting as a proxy and hiding the involved complexity from applications. Leveraging these mechanisms makes enabling authentication via social providers trivial and quick. 

### User Federation

Keycloak has built-in support to connect to existing LDAP or Active Directory servers, allowing custom provider implementations to connect users in other stores, such as a relational database.

### Client Adapters

Keycloak client adapters make it easy to secure applications and services by providing adapters for a number of platforms and programming languages. Keycloak is built on standard protocols which then allows any existing OpenID Connect Resource Library or SAML 2.0 Service Provider library to integrate with it.

It is also possible to use a proxy adapter to secure applications, which removes the need to modify an application at all. This can be leveraged in Kubernetes deployment by deploying the adapter as a side car. 

### Admin Console

All features and aspects of Keycloak configuration can be managed via an Admin UI, the CLI or a set of REST APIs, this includes managing users, permissions and active SSO sessions. 

### Account Management Console

Through the account management console users can manage their own accounts. They can update their user profile, change passwords, and setup two-factor authentication. Users can also manage sessions and view history for the account. If they have enabled social login or identity brokering, users can also link their accounts with additional providers to allow them to authenticate to the same account with different identity providers.

### Authorization Services

Keycloak provides the ability for more fine-grained authorization for scenarios where traditional RBAC is not enough. By implementing the UMA 2 spec it allows managing resources, permissions and sophisticated policies from the Keycloak admin console.

### SPI and pluggable architecture  

All core components in Keycloak are implemented with pluggable architecture. Introducing new or alternative implementations of particular capabilities or required extensions is easy. There are number of externally hosted and maintained plugins which serve as examples to anyone who would like to write their own or contribute to the project.

### High Availability and Active-Active Cross-Data Center

Keycloak supports traditional high availability by providing clustering. In addition, it supports deployments of connected clusters between geo-separated data centers. Keycloak provides both active-active and active-passive (fallback) switching scenarios between configured data centers. 

## Current Status 

Keycloak is a mature and widely adopted project. Currently on 9.x release since February. Although follows quick every few months major release while keeping backwards compatibility. 

Github and community stats (March 2020):
* Forks: 2.6k
* Stars: 5.6k 
* Controbutors: 377
* Commits: 11.5k
* Website visits: >60k unique users per month
* Developer mailing list: ~100 posts/month
* User mailing list: ~200 posts/month

### Governance and Community
* Governance: https://github.com/keycloak/keycloak/blob/master/GOVERNANCE.md
* Contributing Guide: https://github.com/keycloak/keycloak/blob/master/CONTRIBUTING.md
* Maintainers: https://github.com/keycloak/keycloak/blob/master/MAINTAINERS.md
* Adopters: https://github.com/keycloak/keycloak/blob/master/ADOPTERS.md
* License: ASF 2.0 https://github.com/keycloak/keycloak/blob/master/LICENSE.txt
* Community channels: https://www.keycloak.org/community

## Future Plans / Roadmap
Key high level items
* W3C WebAuthN - initial support already in Keycloak 9
* Admin UI Redesign and reimplementation of Admin UIs in ReactJS. https://groups.google.com/d/msgid/keycloak-dev/188f4f10-6135-4220-a399-96f0a6e94ff9%40googlegroups.com
* Kubernetes Operator - initial release on OperatorHub https://operatorhub.io/operator/keycloak-operator 
* Quarkus - https://quarkus.io - based distribution allowing natively compiled binaries and startup/footprint comparable to golang. 
* New improved storage layer - drop requirement for database. 
* Keycloak.X - number of Cloud Native related improvements https://www.keycloak.org/2019/10/keycloak-x
* FAPI (Financial APIs)
* Token Exchange Service
* Introduce Webhooks as extension mechanism
* Config templates / isolation and realm hierarchy
* Authentication improvments - Adaptive / Risk based Step Up Authentication; Flexible consent authentication flows

# Project Scope

## Clear project definition

OpenSource Identity and Access Management Solution for Modern Applications, API and Services. Aiming to make security easy for developers. OpenID Connect Certified and SAML 2.0 compatible Identity Provider. Also User Managed Access (UMA2) implementation. 

Allowing to secure applications with minimum to no code. Providing all typically required authentication, authorization and identity management capabilities for applications. Allowing developers to provide strong security into their applications without implementing typical features like login screens, registration, user management, password policies, figuring out how to securely issue and handle tokens and etc. 

Focusing on modern standards around the OAuth2 ecosystem, while including SAML2 support due to widespread usage. Keeping other technologies (CAS, WS-Fed, etc.) outside of the core codebase and as extensions.  

It is an opinionated solution trying to avoid code and function creep. Aiming to remain fairly lightweight. Delivering on 80/20 principle. Focusing on Cloud Native and modern use cases. 


## Value-add to the CNCF ecosystem

Security is a cornerstone of Cloud Native environments and the OAuth2 family of standards like OpenID Connect has become a default choice when implementing modern applications, services, and APIs. Keycloak fits perfectly into the cloud landscape as a lightweight and modern solution. Embracing developers with an ease of use and and the rapid velocity of integrating Keycloak into applications, Keycloak also embraces new development models. 

Rapid project growth and adoption has proven it has already become a technology of choice for applications developed in Cloud Native ecosystem. 

## Alignment with other CNCF projects

OpenID Connect, OAuth2 and JWT are standards leveraged by all Cloud Native technologies nowadays. 

Kubernetes Official Documentation for authentication points to Keycloak as one of possible OpenID Connect Providers to leverage https://kubernetes.io/docs/reference/access-authn-authz/authentication/ 

OIDC is embraced for end user authentication by service mesh technologies. 

Keycloak provides a standalone adapter called Gatekeeper allowing securing applications via sidecar. Instrumenting security rules without changing application codebase. https://github.com/keycloak/keycloak-gatekeeper 

## Anticipated use cases

Standalone Identity Provider (OpenID Connect, SAML 2.0)

Standalone Authorization Server (OAuth 2.0)

Central Authorization solution (User Managed Access - UMA 2.0). 
* Policy Evaluation Point, custom, chained policies (ABAC, Rules, JavaScript, etc.)

Offloading application from implementing typical capabilities by providing integrations and set of hosted screens and services which can be themed/skinned to look like an integral part of application
* Security (auth/authz)
* Identity and Access Management capabilities (User/Role/Attribute management, password policies, etc.) 
* Authentication / Registration capabilities
* Modern authentication capabilities (W3C WebAuthN, MFA)
* User Self Service (changing profile, password, registering authentication tokens, sessions and consent management)

Shields developer from legacy or custom infrastructure integrations which can be plugged in and leveraged while being hidden from application. Allowing developers to develop using modern protocols (OpenID Connect/OAuth 2.0) while still delivering on traditional integration needs
* Flexible LDAP and Kerberos (SPNEGO) integration
* Custom User Storage implementations
* Social Login Providers (Sign in via FB, Google, Github, Twitter, etc) with flexible attribute mapping
* External Identity Providers (SAML 2.0, OpenID Connect, custom)


## Alignment with SIG Reference Model

CNCF SIG Securiry assesment request: https://github.com/cncf/sig-security/issues/372

CNCF SIG Security Self Assesment: https://docs.google.com/document/d/14IIGliP3BWjdS-0wfOk3l_1AU8kyoSiLUzpPImsz4R0/edit?usp=sharing

## High level architecture

https://docs.google.com/document/d/14IIGliP3BWjdS-0wfOk3l_1AU8kyoSiLUzpPImsz4R0/edit#heading=h.hwn90fa4whrx

# Formal Requirements

Keycloak project is able to meet all requirements. Already has proper governance model, is ready to adopt CNCF Code of Conduct and perform required trademark transfer. Will adhere to the CNCF IP Policy

## Has the TOC been approached for sponsorship
Keycloak has approached Sandbox submission in the past although this has been impacted with process changes. At the point Keycloak submitted CNCF TOC decided to halt intake of new projects and redesign the whole process. This essentially derailed previous submission

Previous PR: https://github.com/cncf/toc/pull/176

## Sponsors from TOC
TBD

## Preferred maturity level
Sandbox 

# Project and Code Quality. Other information
Comprehensive overview of the project has been performed for SIG Security Assesment:

https://github.com/cncf/sig-security/issues/372

https://docs.google.com/document/d/14IIGliP3BWjdS-0wfOk3l_1AU8kyoSiLUzpPImsz4R0/edit?usp=sharing


# Cloud Native
The Keycloak team believes that this project aligns well with section 1(c) of the CNCF Charter by providing a standard and simple way to secure Cloud Native applications and services. Out of the box, Keycloak provides an extensive set of features such as user federation, admin console and account management console. Securing applications and services can be done with only a few lines of code through Keycloak adapters that are provided for a range of languages and frameworks. Keycloak also provides both OpenID Connect and SAML 2.0 enabling any application that have support for either to be easily secured.
Keycloak already has wide community support. Bringing Keycloak into the CNCF, the team hopes to continue to expand the list of features, making it even easier to secure different types of applications and reach an even wider community interested in contribution and adoption.


