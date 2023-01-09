# Dex - Sandbox

**Name of project:**

Dex

**Description:**

Dex is an identity service that uses OpenID Connect to drive authentication for other apps.

Dex acts as a portal to other identity providers through "connectors." This lets dex defer authentication to LDAP
servers, SAML providers, or established identity providers like GitHub, Google, and Active Directory. Clients write
their authentication logic once to talk to dex, then dex handles the protocols for a given backend.

**Statement on alignment with CNCF mission:**

Dex's main production use is as an auth-N addon in for Kubernetes. Dex runs natively on top of any Kubernetes cluster
and can drive API server authentication through the OpenID Connect plugin. Clients, such as kubectl, can act on behalf
users who can login to the cluster through any identity provider dex supports.

Dex has been used by numerous companies who would like for it to find a neutral home where they can actively
contribute to the long term viability of the project.

**Sponsors from TOC:**

TBD

**Preferred maturity level:**

Sandbox

**License:**

Apache2 + DCO

**Source control:**

Github (https://github.com/dexidp/dex)

**Initial Committers:**

Maintainers:

- Joel Speed (Red Hat)
- Stephan Renatus (Chef)
- Nándor Krácser (Banzai Cloud)
- Márk Sági-Kazár (Banzai Cloud)
- Janos Matyas (Banzai Cloud)
- Rithu John (Lyft)
- Vijay Katam (VMware)
- Mo Khan (VMware)

**Infrastructure requests (CI / CNCF Cluster):**

- GitHub Actions for CI/CD
- Container image repository
- Communication Channels:
  - IM/Slack: TODO-Request
  - Issue tracker: https://github.com/dexidp/dex
  - Mailing List: None
  - Website: None

**Existing sponsorship:**

Not at this time, which is the reason for the request

**Community size:**
The community around Dex has been fairly large over time, but has struggled recently due to its abandonment when CoreOS
was acquired. However there is a very active user community and developers have requested that it find a neutral home
so others can actively engage.

- 4.2k Github Stars
- 890 Forks
- 123 Unique contributors
