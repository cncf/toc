# Keycloak CNCF Sandbox Project Submission Proposal
## Description
Keycloak is an open source Identity and access management solution aimed at modern applications and services. It makes it easy to secure applications and services with little to no code. 

Keycloak is based on standard protocols with an aim toward modern use cases and the flexibility to integrate with other solutions and prevent vendor lock in. Supported protocols include: OAuth2, OpenID Connect, User Managed Access 2.0 (UMA) and SAML 2.0. 

### Single-Sign On

Users authenticate with Keycloak rather than to individual applications. This frees developers from dealing with login forms, authenticating users, and storing users. Keycloak provides both single sign-in and single sign-out to different types of applications (client-side, server-side, mobile, etc.) and protocols (OIDC and SAML2) within same realm. It also provides Kerberos bridge to propagate authentication (LDAP or Active Directory) from workstations to web applications via SPNEGO. 

### Token based security

Through OpenID Connect flows applications retrieve access tokens that can be used to invoke services securely without the need to pass user credentials. This enables easy end-to-end user authentication for client-side applications, micro services, service mesh and other modern architectures.
Identity Brokering and Social Login

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

## Value to the Cloud Native Operating Model

Security is a cornerstone of Cloud Native environments and the OAuth2 family of standards (OpenID Connect / UMA) has become a default choice when implementing modern applications, services, and APIs. Keycloak fits perfectly into the cloud landscape as a lightweight and modern solution. Embracing developers with an ease of use and and the rapid velocity of integrating Keycloak into applications, Keycloak also embraces new development models. 

Rapid project growth and adoption has proven it has already become a technology of choice for applications developed in Cloud Native ecosystem. 

## Roadmap
### Improvements around provisioning and maintenance within Cloud Native environments 
* Operator implementation (Operator Framework). WIP already
* Service Broker implementation. WIP already
* Zero downtime upgrades
* Observability
* Storage improvements (reconsidering requirement for traditional RDBMs)
* Runtime improvements (switching to lighter runtime based on Java Microprofile to optimise memory footprint) 

### Core IdP features
* New Account Management Console
* Account Management REST API
* Encrypted storage for credentials
* Authentication improvements
* Additional two factor mechanisms
* Web Authentication (W3C WebAuthN / FIDO2) support
* Step-up authentication
* Misc enhancements

## Current CNCF Ecosystem Integrations:
Keycloak as a project is already present in the wider CNCF or Cloud Native ecosystems

According to recent Keycloak Community Survey:
* 33% deploy Keycloak using Docker
* 16.2% deploy Keycloak using Kubernetes
* 6.7 % deploy Keycloak on AWS
* 4.8% deploy Keycloak using OpenShift

Kubernetes Official Documentation for authentication points to Keycloak as one of possible OpenID Connect Providers to leverage https://kubernetes.io/docs/reference/access-authn-authz/authentication/ 

Keycloak can be used to provide End User Authentication for Istio Service Mesh: https://github.com/kameshsampath/istio-keycloak-demo 

Keycloak contains container images deployable to Docker, Kubernetes and OpenShift as well as example template for Kubernetes and OpenShift and can be used natively in the Kubernetes ecosystem. 

Keycloak integrates with OpenShift https://github.com/keycloak/openshift-integration 

Keycloak provides a standalone adapter called Gatekeeper allowing securing applications via sidecar. Instrumenting security rules without changing application codebase. https://github.com/keycloak/keycloak-gatekeeper 

## Sponsors from TOC
TBD

## Current Project Sponsor: 
Red Hat

## Preferred maturity level
Sandbox 

## Unique identifier
“keycloak”

## License
Apache License v2 (ASL v2)

## Code Repositories
https://github.com/keycloak

## Ticket / Planning
JIRA at https://issues.jboss.org/browse/KEYCLOAK 

## Documentation
https://www.keycloak.org/documentation.html 

## Extensions
Keycloak aims to avoid feature creep keeping the set of core, out-of-the-box capabilities focused on common use-cases. By maintaining a pluggable architecture it also aims to make extensions easy to implement and maintain. A number of such extensions are hosted externally and maintained by the community, such as WS-Trust and CAS protocol integrations. 

https://www.keycloak.org/extensions.html

## External Code Dependencies

* WildFly (https://wildfly.org/) - Apache License 2.0
* Hibernate ORM (http://hibernate.org/) - GNU Lesser General Public License 2.1
* RESTEasy (https://resteasy.github.io/) - Apache License 2.0
* Infinispan (http://infinispan.org/) - Apache License 2.0
* Jackson (https://github.com/FasterXML/jackson) - Apache License 2.0
* ZXing (https://github.com/zxing/zxing) - Apache License 2.0
* OWASP Java HTML Sanitizer (https://www.owasp.org/index.php/OWASP_Java_HTML_Sanitizer_Project) - Apache License 2.0
* XStream (http://x-stream.github.io/) - BSD license
* Apache FreeMarker (http://x-stream.github.io/) - https://freemarker.apache.org/
* Liquibase (https://www.liquibase.org/) - Apache License 2.0
* Twitter4j (http://twitter4j.org/) - Apache License 2.0
* Drools (https://www.drools.org/) - Apache License 2.0

## Maintainers
8 - commiters with push rights to the main repository and actively using it.

## Contributors
https://github.com/keycloak/keycloak/graphs/contributors 

### Statistics from Github
* 287 on main https://github.com/keycloak/keycloak repository 
* 93 on documentation repository https://github.com/keycloak/keycloak-documentation 
* 39 on NodeJS adapter repository https://github.com/keycloak/keycloak-nodejs-connect 
* 25 on Keycloak standalone adapter (GateKeeper) repository https://github.com/keycloak/keycloak-gatekeeper 

## Adopters / Users of Note
https://github.com/keycloak/keycloak/blob/master/ADOPTERS.md

## Community Communication
The majority of communication happens via mailing lists, allowing greater visibility and community involvement.

* User Mailing List: https://lists.jboss.org/mailman/listinfo/keycloak-user 
* Developer Mailing List: https://lists.jboss.org/mailman/listinfo/keycloak-dev 

## Website / Blog
https://keycloak.org 

## Release Cadence:
Every 3 weeks there is a new Keycloak release containing new features, capabilities, incremental improvements, and bugfixes. 

Changelog for the current 4.x stream: https://www.keycloak.org/docs/latest/release_notes/index.html 

## Social media accounts
Only Twitter: https://twitter.com/keycloak 

## Community size and any existing sponsorship
Data from 12th Nov 2018:
* Github stars: 3031
* Github forks: 1620
* Github commits: 11163
* Github contributors: 287 

Developer mailing list (keycloak-dev@jboss.org): ~150 posts/month
User Mailing list (keycloak-users@jboss.org): ~300 posts/month

Keycloak is currently being sponsored by Red Hat with 20+ people working full time on the project. Leveraging JIRA hosted on jira.jboss.org

## Infrastructure requests from CNCF
CI resources. Currently community testing happening in free Travis plan. 

## Handling of security issues
https://www.keycloak.org/security.html 

## Statement on alignment with CNCF mission
The Keycloak team believes that this project aligns well with section 1(c) of the CNCF Charter by providing a standard and simple way to secure Cloud Native applications and services. Out of the box, Keycloak provides an extensive set of features such as user federation, admin console and account management console. Securing applications and services can be done with only a few lines of code through Keycloak adapters that are provided for a range of languages and frameworks. Keycloak also provides both OpenID Connect and SAML 2.0 enabling any application that have support for either to be easily secured.
Keycloak already has wide community support. Bringing Keycloak into the CNCF, the team hopes to continue to expand the list of features, making it even easier to secure different types of applications and reach an even wider community interested in contribution and adoption.

