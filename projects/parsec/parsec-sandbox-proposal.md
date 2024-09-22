For TOC Consideration: please adopt PARSEC into CNCF as a SandBox project. This proposal is being
made in accordance with the [**stated
process**](https://github.com/cncf/toc/blob/master/process/project_proposals.md) and the
[**proposal template**](https://github.com/cncf/toc/issues/344).

Authors: - Paul Howard, @paulhowardarm, Arm Limited, PARSEC technical lead

# Background

**Link to TOC PR** (this document)

[**Link to
Presentation**](https://drive.google.com/file/d/1GymBrSWWd8B6roayB80L5Gk41X-fDy-e/view?usp=sharing)

[**Link to GitHub project**](https://github.com/parallaxsecond/parsec)

## Project Goal

**PARSEC** is the **P**latform **A**bst**R**action for **SEC**urity. The goal of Parsec is to
provide a common API to secure services in a platform agnostic way.

Parsec aims to define a universal software standard for interacting with secure object storage,
Root-of-Trust services, and cryptography services, creating a common way to interface with functions
that would traditionally have been accessed by more specialised APIs, often requiring specific
knowledge of the platform hardware. The goal is to allow application developers to code against a
common API without requiring this knowledge, and for them to do so in a way that is fluent and
natural in any popular programming language. Parsec contributes to the cloud-native vision by
de-coupling applications and services from these physical platform details.

## Current Status

- **Project releases:** Parsec was made available as a public open source project in October 2019.
   It is a [Rustlang](https://www.rust-lang.org) project, and regular releases are being made to the
   Rust package repository at [crates.io](https://crates.io). For release history, please see the
   package repository links:
   - [parsec-service](https://crates.io/crates/parsec-service)
   - [parsec-interface](https://crates.io/crates/parsec-interface)
   - [tss-esapi](https://crates.io/crates/tss-esapi)
- **Community status:** Parsec has a [community Slack channel](http://dockr.ly/slack) and a
   bi-weekly
   [meeting](https://calendar.google.com/calendar?cid=ZG9ja2VyLmNvbV9xcHAzbzl2aXBhbmE0NGllcmV1MjlvcHZkNEBncm91cC5jYWxlbmRhci5nb29nbGUuY29t).
   The principal GitHub repository has 61 stars, 8 forks and 10 individual contributors including
   some from Arm, Docker and Mirantis.

## Available functionality

- Subset of cryptographic primitives implemented, aimed at supporting use cases such as mTLS
   signatures:
   - Provisioning of asymmetric key pairs
   - Export of public key
   - Signing of digests with private key
- Available back-end integrations (limited to above primitives):
   - Trusted Platform Module (TPM) using the [TPM
      2.0](https://www.trustedcomputinggroup.org/wp-content/uploads/TPM-2.0-A-Brief-Introduction.pdf)
      API.
   - Hardware Security Modules via the
      [PKCS#11](http://docs.oasis-open.org/pkcs11/pkcs11-base/v2.40/os/pkcs11-base-v2.40-os.html)
      cryptographic standard
   - Software implementation based on [Mbed Crypto](https://github.com/ARMmbed/mbed-crypto) for
      simple experimentation

## Future Plans

#### Current Work Items

- Rust client library
- C client library conforming to the Platform Security Architecture (PSA) Crypto API specification
- Proof-of-concept integration efforts with multiple partner organizations
- Integration into a commercial product (initially within Arm)
- Integration into Yocto Linux

#### Backlog Items

- Increase API coverage to the full PSA Crypto specification
- Additional back-end integrations
- Establish design guidelines and patterns for client library development
- Implement client libraries in various popular programming languages.
- Introduce API for secret storage/retrieval
- Review and implement the multi-tenancy design
- Support fully-virtualized environments

# Project Scope

## Clear project definition

The role of Parsec is to provide a common, platform-agnostic API to secure facilities for key
storage and cryptography. The common API closely tracks the [**PSA Crypto
API**](https://github.com/ARM-software/psa-arch-tests/blob/master/api-specs/crypto/v1.0/doc/IHI0086-PSA_Cryptography_API-1.0.0.pdf),
which is well-specified and documented publicly. Parsec has a very extensible architecture, and has
scope to grow in order to support different platforms and different deployment patterns, but its
fundamental vision and value proposition remains the same.

![Parsec Scope Diagram](https://drive.google.com/uc?id=1FoZ_9rdwhSj2mT19Vz64ox8LlqX7hAEf)

## Value-add to the CNCF ecosystem

Parsec enhances the CNCF ecosystem by providing new opportunities for loose coupling between
applications/services and the physical platforms on which they run, whether in data centers or at
the edge.

Parsec has been designed specifically for cloud-native environments where workloads are orchestrated
and containerized such as to be independent of the physical platform, and where multiple distinct
workloads may be running on a single host device, such as a server or an edge gateway.

Parsec's value proposition can be stated as follows:

- **Abstraction** – a common API that is truly agnostic and based on modern cryptographic
   principles
- **Mediation** – security as a microservice, brokering access to the hardware and providing
   isolated key stores in a multi-tenant environment
- **Ergonomics** – a client library ecosystem that brings the API to the fingertips of developers
   in any programming language: “easy to consume, hard to get wrong”
- **Openness** – an open-source project inviting contributions to enhance the ecosystem both
   within the service and among its client libraries

## Alignment with other CNCF projects

**Does the project align and actively collaborate with other CNCF projects?**

The Parsec project is still actively seeking opportunities for alignment and collaboration. However,
some existing CNCF projects can be seen as being highly relevant. Parsec provides security services
to multiple applications on any given physical host. It is able to provide mutually-isolated key
storage facilities to those applications. For this to be successful, there will need to be some
degree of integration between Parsec and the orchestration and identity subsystems of the
environment in which it is deployed. Hence there is a connection of relevance between Parsec and
existing CNCF projects with regards to orchestration and identity, summarised in the image below.

![Parsec CNCF Alignment](https://drive.google.com/uc?id=1GltZ-jzgEIOwKtoSiA58R-_cgiSzhknB)

For further details on how Parsec is expected to be integrated with orchestration and identity
subsystems, please refer to the [**system
architecture**](https://parallaxsecond.github.io/parsec-book/parsec_service/system_architecture.html).

**Does the project require any specific versions of projects (or APIs) to interoperate? (e.g. K8s
API, CSI, CNI, CRI)?**

Parsec does not have any specific dependencies on other CNCF projects, and hence does not require
any specific versions of those projects or their associated APIs.

**Does the project augment or benefit other CNCF projects?**

See section above on how Parsec may be seen to be relevant for other CNCF projects. Parsec's value
as a platform security abstraction could potentially offer benefits to other projects as well.

## Anticipated use cases

- As a **portable Root-of-Trust**, allowing an application or service to sign payloads with a
   private key without the key material being exposed, such as might be needed for an mTLS
   handshake.
- As a provider of **general cryptographic services** to applications in a way that is fluent and
   natural for them to consume, and isolating them from knowledge of the platform hardware.
- As a **brokering service** in situations where multiple applications are competing for the same
   secure resources on the same host device.
- As a provider of a **secret store** facility.

## Alignment with SIG Reference Model

**Does the project align with the SIG CNCF reference model and which capabilities does it
require/provide at each level of the reference model.**

## High level architecture

The core component of Parsec is the **security service** (or **security daemon**). This is a
background process that runs on the host platform and provides connectivity with the secure
facilities of that host. It makes these facilities available to **client applications** in the form
of a standard API based on the [**PSA Crypto API
Specification**](https://github.com/ARM-software/psa-arch-tests/blob/master/api-specs/crypto/v1.0/doc/IHI0086-PSA_Cryptography_API-1.0.0.pdf),
which is exposed to the clients using a custom [**wire
protocol**](https://parallaxsecond.github.io/parsec-book/parsec_client/wire_protocol.html). Clients
do not consume the wire protocol directly, but consume the API using a **client library** in their
preferred programming language, which handles the low-level details of the wire protocol.

The wire protocol is **stream-oriented** but does not prescribe a particular transport. The initial
supported transport is for Unix sockets, but other transports can be built.

The security service is written in **Rust** for its security and safety characteristics.

The security service has pluggable back-end modules known as **providers**. A Parsec provider is the
piece that knows how to talk to the platform-specific secure facilities, such as a hardware secure
element, TPM, HSM or trusted application. The provider model allows Parsec to be extended with
support for hardware and software security on multiple platforms and from multiple vendors.

See the architecture diagram for Parsec below.

![Architecture Diagram](https://drive.google.com/uc?id=1BlIKFN4QgtuJbmQIJcRhwBUfFUDV2tR5)

For further details, please refer to the [**Parsec
Book**](https://parallaxsecond.github.io/parsec-book/).

# Formal Requirements

**Document that the project fulfils the requirements as documented in the CNCF graduation criteria
for sandbox**

- TOC sponsorship pending (see below)
- Parsec will adopt the CNCF code of conduct
- CNCF IP policy will be adopted subject to final review within contributing organizations

**Are there any anticipated issues with any of the criteria ?** None

**Has the TOC been approached for sponsorship? If yes, which TOC members have agreed to sponsor the
project?** There has been no broad approach to the TOC. We have been advised that TOC sponsorship no
longer needs to be sought as a prerequisite of submission in the current process. However, Justin
Cormack has offered to be a sponsor if one is needed.

# CNCF IP Policy

**Becoming a sandbox project requires adoption of the CNCF IP Policy:
https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy**

**Note: there is a grace period after becoming a sandbox period to enable projects to adopt the
policy, however, some prep is required to ensure there are no major blockers.**

**Has the IP policy been reviewed?**

Yes. All source code in the project is Apache 2.0 licensed today, which creates minimum friction for
adopting the CNCF IP policy. There is currently no intention to have a Contributor License Agreement
(CLA) for the project.

**List the repos for the project and their current license**

| Repo Name                                                                  | License                    |
|----------------------------------------------------------------------------|----------------------------|
| [parsec](https://github.com/parallaxsecond/parsec)                         | Apache License Version 2.0 |
| [parsec-interface](https://github.com/parallaxsecond/parsec-interface-rs)  | Apache License Version 2.0 |
| [parsec-client-rust](https://github.com/parallaxsecond/parsec-client-rust) | Apache License Version 2.0 |
| [parsec-client-test](https://github.com/parallaxsecond/parsec-client-test) | Apache License Version 2.0 |
| [parsec-book](https://github.com/parallaxsecond/parsec-book)               | Apache License Version 2.0 |
| [rust-tss-esapi](https://github.com/parallaxsecond/rust-tss-esapi)         | Apache License Version 2.0 |
| [parsec-client-go](https://github.com/parallaxsecond/parsec-client-go)     | Apache License Version 2.0 |
| [parsec-operations](https://github.com/parallaxsecond/parsec-operations)   | Apache License Version 2.0 |
| [rust-psa-crypto](https://github.com/parallaxsecond/rust-psa-crypto)       | Apache License Version 2.0 |

**List any dependent repos (upstream/downstream) that are required to build the project (including
but not limited to libraries, commercial tools, plugins)**

| Repo Name                                             | License                    |
|-------------------------------------------------------|----------------------------|
| [mbed-crypto](https://github.com/ARMmbed/mbed-crypto) | Apache License Version 2.0 |

The Parsec project also uses the following third-party Rust crates:

- serde (MIT and Apache-2.0)
- bindgen (BSD-3-Clause)
- cargo_toml (Apache-2.0)
- toml (MIT and Apache-2.0)
- rand (MIT and Apache-2.0)
- base64 (MIT and Apache-2.0)
- uuid (MIT and Apache-2.0)
- threadpool (MIT and Apache-2.0)
- std-semaphore (MIT and Apache-2.0)
- num_cpus (MIT and Apache-2.0)
- signal-hook (MIT and Apache-2.0)
- sd-notify (MIT and Apache-2.0)
- log (MIT and Apache-2.0)
- env_logger (MIT and Apache-2.0)
- pkcs11 (Apache-2.0)
- picky-asn1-der (MIT and Apache-2.0)
- picky-asn1 (MIT and Apache-2.0)
- bincode (MIT)
- structopt (MIT and Apache-2.0)
- derivative (MIT and Apache-2.0)
- arbitrary (MIT and Apache-2.0)
- libfuzzer-sys (MIT, Apache-2.0 and NCSA)
- flexi_logger (MIT and Apache-2.0)
- lazy_static (MIT and Apache-2.0)
- version (MIT and Apache-2.0)

This list is correct at time of writing. The same list is maintained and kept current in the
[**Parsec README**](https://github.com/parallaxsecond/parsec/blob/master/README.md).

**What actions are required to be compliant with the IP policy?** There needs to be a legal
agreement to transfer the logo ownership, but there are no anticipated issues with this.

# Other Considerations

## **Please note, these are not gating criteria but rather to:**

- **Collect a standard set of information for each project**
- **Provides a point in time capture of the state of the project which makes it easier to track
   progress at future reviews and / or promotion**
- **Help projects to prepare for SIG and TOC presentation**
- **Allow the SIG to review the project and perform due diligence for incubation**
- **Provide the TOC with the information required to accept sponsorship of a project and/or votes**
- **Identify and rectify any significant issues / blockers prior to presenting to the TOC and
   acceptance as a CNCF project**

## Cloud Native

**Does the project meet the definition of Cloud Native? The CNCF charter states:**

“Cloud native technologies empower organizations to build and run scalable applications in modern,
dynamic environments such as public, private, and hybrid clouds. Containers, service meshes,
microservices, immutable infrastructure, and declarative APIs exemplify this approach.

“These techniques enable loosely coupled systems that are resilient, manageable, and observable.
Combined with robust automation, they allow engineers to make high-impact changes frequently and
predictably with minimal toil.”

Parsec meets this definition because it enables new opportunities for loose coupling between
applications and the physical platforms on which they run.

## Project and Code Quality

**Are there any metrics around code quality? Are there good examples of code reviews? Are there
enforced coding standards?**

- The project has a set of written [**contribution
   guidelines**](https://github.com/parallaxsecond/parsec/blob/master/CONTRIBUTING.md)
- Rust code submissions are required to apply the `rustfmt` and `clippy` tools to ensure consistency
   of style.
- Unit tests and integration tests are a stated requirement of code submission.
- There is a [**security threat
   model**](https://parallaxsecond.github.io/parsec-book/threat_model/threat_model.html)
- A good example of code review interactions can be seen in the [**TPM Support Pull
   Request**](https://github.com/parallaxsecond/parsec/pull/75).

**What are the performance goals and results? What performance tradeoffs have been made? What is the
resource cost?**

- Parsec defines an abstraction over functionality provided by the platform. Its general performance
   goal is to add minimal overhead relative to accessing the platform-specific functions directly.
- One of the design trade-offs was to keep the IPC mechanism as lightweight as possible. Parsec has
   a [**custom wire
   protocol**](https://parallaxsecond.github.io/parsec-book/parsec_client/wire_protocol.html), but
   uses [**protobuf**](https://developers.google.com/protocol-buffers) as a standard serialization
   mechanism so that APIs can be used to transact rich/structured data where needed.
- Protobuf was considered alongside [**CBOR**](https://cbor.io),
   [**JSON**](https://www.json.org/json-en.html) and
   [**MessagePack**](https://msgpack.org/index.html). Protobuf was chosen because it is
   schema-oriented, which made it more suited to defining contracts, and has favourable performance
   characteristics.
- The Rust programming language was chosen partly for its predictable performance profile in
   combination with good security and safety features.

**What is the CI/CD system? Are there code coverage metrics? What types of tests exist?**

- There is a [**CI
   build**](https://github.com/parallaxsecond/parsec/actions?query=workflow%3A%22Continuous+Integration%22).
- Test framework including unit tests, integration tests, stress tests and fuzzing. For more
   information, see [**How to Test
   Parsec**](https://parallaxsecond.github.io/parsec-book/parsec_service/test.html).

**Is there documentation?**

There is an extensive body of documentation that is maintained in the [**Parsec
Book**](https://parallaxsecond.github.io/parsec-book/index.html).

**How is it deployed?**

- Parsec can be deployed as a native process (a userspace daemon) on supported platforms
- Parsec has a runtime configuration mechanism
- The integration of Parsec into a suitable Yocto Project layer is under consideration. A suitable
   layer may be [**meta-security**](https://git.yoctoproject.org/cgit/cgit.cgi/meta-security), but
   this has not been finalized
- Parsec can be deployed within a container

**How is it orchestrated?**

The Parsec service abstracts over the security capabilities of the physical platform on which it is
running. Consequently, exactly one instance of the Parsec service needs to be running on each
physical host. Communications between client applications and the Parsec service take place on (and
do not leave) the physical host.

**How will the project benefit from acceptance into the CNCF?**

Parsec's success hinges on the widest possible adoption by industry partners. Companies typically
see open governance as a hugely influential factor in deciding whether to adopt an open-source
technology. For some organizations, the lack of open governance is considered to be a blocker for
adoption. By being accepted into CNCF, Parsec stands a greater chance of being widely adopted, which
in turn should draw more contributions to enhance its ecosystem of back-end integrations and client
libraries.

**Has a security assessment by the security SIG been done? If not, what is the status/progress of
the assessment?**

There has not been a security assessment at time of writing. However, the project does have a
[**published threat
model**](https://parallaxsecond.github.io/parsec-book/threat_model/threat_model.html), and feedback
on the threat model is openly invited.

# Promotion to Incubation

## Open Governance

**How are committers chosen?**

Commits are accepted through PRs, reviewed and moderated by the [**list of
maintainers**](https://github.com/parallaxsecond/parsec/blob/master/MAINTAINERS.toml).

**How are architectural and roadmap decisions made?**

- Some roadmap items are defined and pursued within the sponsoring organization.
- In this initial phase, architectural decisions are collaborative amongst the current three
   maintainer organizations: Arm, Docker and Mirantis. This group of maintainer organizations may
   expand over time.
- Architecture and roadmap decisions are discussed openly on the public Slack channel to include
   community members.
- Roadmap and architecture are also discussed between the maintainer organizations and their
   industry partners.

**How many decision makers are outside the sponsoring organization.**

At time of writing, three maintainers are outside the sponsoring organization.

## Adoption

**Who are the current maintainers?**

The full list of current maintainers is
[**here**](https://github.com/parallaxsecond/parsec/blob/master/MAINTAINERS.toml)

**How long has the project been developed for?**

The Parsec project was initially conceived in April 2019 and was published as open source in October
2019.

**Is there a commercial version of the project or a primary commercial sponsor ?**

Parsec has been conceived as an open-source initiative, and has no commercial versions or associated
services.

**Is the project used in production? If so, please list some of the accounts.**

At time of writing, Parsec has not been used in production, but there is active work to consume the
Parsec service into multiple commercial products.

**Does the project participate in a CNCF User Group?**

Parsec does not participate in any CNCF User Groups today, but the project is actively seeking
further community engagement so this will be considered.

## Vendor Independence

**Is the project reasonably independent from the sponsoring vendor?**

As per the [**list of
maintainers**](https://github.com/parallaxsecond/parsec/blob/master/MAINTAINERS.toml), there are
three separate organizations maintaining the project at this time:

- Docker, Inc.
- Arm Limited.
- Mirantis, Inc.

**Are all communication channels and project resources hosted just for this project or with other
CNCF projects/resources?**

This project has a Slack channel **#parsec** on the [**Docker
Community**](https://docker-community.slack.com) Slack instance.

Communications for this project are independent of those for other CNCF projects.

**Is all code that is part of the project hosted and part of the CNCF managed orgs and repos?**

Yes, the intention is to donate all of the project's existing GitHub repositories to CNCF.

**Are all defaults for upstream reporting either unset or community hosted infrastructure (i.e.
doesn’t point to vendor hosted SaaS control plane or analytics server for usage data)? Is all
project naming independent of vendors?**

All reporting is managed through GitHub PRs and issues. There is no requirement to interact with any
vendor-hosted services either to consume Parsec or to contribute to it.

**Relevant Assets regarding vendor independence**
N/A
