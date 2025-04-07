For TOC Consideration: please adopt Keylime into CNCF as a SandBox project.
This proposal is being made in accordance with the [**stated
process**](https://github.com/cncf/toc/blob/master/process/project_proposals.md) and the
[**proposal template**](https://github.com/cncf/toc/issues/344).

Authors: - Luke Hinds, @lukehinds, Red Hat, Keylime Project Lead

# Background

**Link to TOC PR** (this document)

[**Link to
Presentation**](https://docs.google.com/presentation/d/1VPJNRSkD5z2d2kTXzB46xM_s6uwXcfaEvlbquEiFq10/)

[**Link to GitHub project**](https://github.com/keylime)

## Project Goal

Keylime provides an end-to-end solution for establishing and maintaining
cryptographic based trust  of remote workloads, the provisioning of
encrypted payloads and run-time system integrity monitoring.

Keylime's mission is to make TPM Technology easily accessible to developers and
users alike, without the need for a deep understanding of the lower levels of a
TPM's operations.

## Current Status

Keylime was made public as an open source project in 2016.

It is a [Python](https://www.python.org/) and [Rustlang](https://www.rust-lang.org) project

- **Project releases:**

  - Releases are made within the project GitHub repository [keylime](https://github.com/keylime/keylime/releases)
using the [Semantic Versioning](https://semver.org)
- **Community status:**

  - Keylime has a [Gitter community channel](https://gitter.im/keylime-project/community).
  - The principal GitHub repository has 89 stars, 35 forks and 19 individual contributors
   including MIT, Red Hat, IBM, Intel and NetFlix.
- **Licensing**
  - Keylime uses the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) License

## Available functionality

### Remote Attestation

Keylime can remotely monitor a nodes trust, based on measurements captured
using [trusted platform module](https://whatis.techtarget.com/definition/trusted-platform-module-TPM) technology.

Keylime can provide remote:

- Measured System Boot
     - firmware
     - bootloader
     - kernel cmd line options.
     - initramfs
     - modules
- Measured Runtime Integrity
     - Uses Kernel Security Sub System IMA (Integrity Measurement Architecture)
     - IMA provides a system call hook to measure (hash) objects into the TPM
     - Hash is compared with trusted list of expected values

### Application Provisioning
  - If a targeted node **passes** its trust attestation, Keylime can execute a
  delegated auto run style script (which could be a wrapper to call applications
  such as `kubectl` or a deployment framework such as `ansible` or `puppet`
  / `chef`)

### Encrypted payloads
 - If a targeted node **passes** its trust attestation, Keylime will deliver a
 secure payload (e.g. SSH keys, TLS certs, config files with secrets)

### Revocation framework
  - If a targeted node **fails** its trust attestation, a signed event will be
  sent to all other nodes monitored by Keylime, instructing them to perform
  specific actions. This uses an open framework where any action that is
  programmatically possible can be triggered, such as:
   - Revoke TLS Certificates
   - Fail over a system
   - Shutdown network / VPN interfaces
   - Remove host from SSH `known_hosts`
   - Call various APIs

## Future Plans

### vTPM Support

Work is underway to:

- Expose a vTPM in the opencontainers/runc project.
This will allow Keylime to measure within a container.

- Create an IMA namespace in the Linux Kernel. This will then
allow Keylime to remotely attest object integrity state within a
container.

- vTPM provider functionality that will pool quotes from virtual TPMs for
attestation by a hardware TPM.

#### Current Work Items

- JSON Web Token (JWT) implementation
- Rust Port of the Python agent
- Prototype of Keylime integration with Kubernetes
- Prototype of vTPM provider verifier

#### Backlog Items

- vTPM provider verifier
- Graphical User Interface overhaul
- Multi Tenancy

# Project Scope

## Clear project definition

The role of Keylime is to provide a set of APIs to enable Remote Attestation of
a target node.

Node types could be:

* infrastructure: A Bare metal machine that will be provisioned for cloud use
  (such as a Kubernetes cluster)

* Virtual Machine: A VM can be attested for trust prior to deployment of an
  application(s)

* Container: A container is measured for trust prior to instantiation of an
  application(s)

Each node type is measured using a hardware TPM device or a virtualized vTPM.

A user can have Keylime deploy and encrypted payload and run a designated
"auto run" style script, once a machine passes it trust checks.

When a node fails, Keylime can be directed to perform actions to revoke a node.
This revocation  action is customizable by the user.

![Keylime Scope Diagram](https://github.com/lukehinds/file-store/raw/master/keylime-scope.png)

## Value-add to the CNCF ecosystem

Keylime has many synergies with a Cloud Native approach. It was designed from
the outset to operate within a cloud architecture.

Trust is difficult to achieve when a user lacks either physical access to a
machine to ensure it has not been maliciously tampered or from a trust view
limited to a higher abstracted view of a system, such as within a shell
environment of a container or a virtual machine. They have no explicit way of
measuring the trust of the underlying host and its runtime, only an implicit view
can be used to make decisions of trust.

Users also lack a means to monitor the integrity state of the nodes that run
their cloud native application(s). Keylime provides this trust outside of a
software context where measurements are typically stored in memory or disc
(which are in turn susceptible to tampering). Keylime achieves this by
performing a one way hash extend of measurements into a TPM / vTPM.

Using an RSA key pair, with the private key part stored within the TPM and
inaccessible to software, Keylime can 'attest' measurements using the public key
counterpart.

Keylime empowers users to remotely measure a cloud based node for trust using a
"root of trust" based on protected keys stored with a trusted platform module.

## Alignment with other CNCF projects

**Does the project align and actively collaborate with other CNCF projects?**

Keylime has discussed collaboration with the [in-toto](https://in-toto.io/)
community. Plans are in place to have Keylime utilise file digests captured by
in-toto. Keylime will then use the file digests to monitor a running system and
ensure nothing has occurred to change the hash measurement of the file originally
checked in by a developer and measured by [in-toto](https://in-toto.io/).

Users of Keylime have demonstrated Keylime bootstrapping a Kubernetes cluster
and monitoring a etcd cluster. A demonstration of this can be seen on the following
[YouTube video](https://www.youtube.com/watch?v=Qhr_aVBCZPw)

The CNCF sandbox [parsec](https://github.com/parallaxsecond/parsec) project and
Keylime our in discussion to have Keylime use the parsec [TSS crate](https://github.com/keylime/rust-keylime/issues/74#issuecomment-650206630)
and our exploring closer integration of the two projects, as both teams see lots
of synergy between Keylime and Parsec. Keylime orchestrates trust and Parsec
abstracts security for the platform, so both projects have their own specialised
area and do not try to solve the same issue.

We would also like to explore synergies with the Notary v2 project and other
CNCF projects.

**Does the project require any specific versions of projects (or APIs) to interoperate? (e.g. K8s
API, CSI, CNI, CRI)?**

Keylime does not have any specific dependencies on other CNCF projects, and
hence does not require any specific versions of those projects or their
associated APIs.

**Does the project augment or benefit other CNCF projects?**

See section above on how Keylime may be seen to be relevant for other CNCF
projects such as [in-toto](https://in-toto.io/) and [parsec](https://github.com/parallaxsecond/parsec)

## Anticipated use cases

- Attest hardware machine for non tampering.
- Attest virtual machine for non tampering.
- Attest a container for non tampering.

## Alignment with SIG Reference Model

**Does the project align with the SIG CNCF reference model and which capabilities does it
require/provide at each level of the reference model.**

## High level architecture

Keylime also provides frameworks for customised revocation events upon trust and
provision of an encrypted payload.

Keylime can be driven with a CLI application, web front end and a set of RESTful
APIs.

Keylime consists of three main components; The Verifier, Registrar and the Agent

* The Verifier continuously verifies the integrity state of the machine that
the agent is running on. This is typically situated within the users trust
boundary

* The Registrar is a database of all agents registered with Keylime and hosts
the public keys of the TPM vendors. This is typically situated within the users
trust boundary

* The Agent is deployed to the remote machine that is to be measured or
provisioned with secrets stored within an encrypted payload released once
trust is established. This is typically situated outside of the users trust
boundary (another provider, edge or IoT device).

![Keylime Architecture Diagram](https://github.com/lukehinds/file-store/raw/master/keylime-arch.png)

# Formal Requirements

**Document that the project fulfils the requirements as documented in the CNCF graduation criteria
for sandbox**

- Keylime will adopt the CNCF code of conduct
- CNCF IP policy will be adopted subject to final review within contributing organizations

**Are there any anticipated issues with any of the criteria ?**

None

**Has the TOC been approached for sponsorship? If yes, which TOC members have agreed to sponsor the
project?**

There has been no broad approach to the TOC. We have been advised that TOC sponsorship no
longer needs to be sought as a prerequisite of submission in the current process.

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

| Repo Name                                                                               | License                    |
|-----------------------------------------------------------------------------------------|----------------------------|
| [keylime](https://github.com/keylime/keylime)                                           | Apache License Version 2-0 |              |
| [keylime-docs](https://github.com/keylime/keylime-docs)                                 | Apache License Version 2.0 |
| [enhancements](https://github.com/keylime/enhancements)                                 | Apache License Version 2.0 |
| [ansible-keylime-tpm-emulator](https://github.com/keylime/ansible-keylime-tpm-emulator) | Apache License Version 2.0 |
| [keylime-deb](https://github.com/keylime/keylime-deb)                                   | Apache License Version 2.0 |
| [keylime.github.io](https://github.com/keylime/keylime.github.io)                       | Attribution link           |
| [keylime-rpm](https://github.com/keylime/ansible-keylime)                               | Apache License Version 2.0 |
| [meetings](https://github.com/keylime/meetings)                                         | Apache License Version 2.0 |
| [rust-keylime](https://github.com/keylime/rust-keylime)                                 | Apache License Version 2.0 |

**List any dependent repos (upstream/downstream) that are required to build the project (including
but not limited to libraries, commercial tools, plugins)**

Keylime (python) uses the following dependencies

- cryptography>=2.1.4 # BSD/Apache-2.0
- m2crypto>=0.21.1 # MIT
- tornado>=5.0.2 # Apache-2.0
- pyzmq>=14.4 # LGPL+BSD
- pyyaml>=3.11 # MIT
- simplejson>=3.8 # MIT
- requests>=2.6 # Apache-2.0
- sqlalchemy>=1.3 # MIT
- pyjwt>=1.7.1 # MIT
- Werkzeug>=0.14.1

Keylime (rust) uses the following dependencies

- futures = "0.1.24"
- hyper = "0.12"
- rust-ini = "0.12.1"
- tokio-io = "0.1"
- serde_json = "1.0"
- pretty_env_logger = "0.2.0"
- log = "0.4"
- libc = "0.2.43"
- serde = "1.0.80"
- serde_derive = "1.0.80"
- rustc-serialize = "0.3.24"
- tempfile = "3.0.4"
- openssl = "0.10.15"
- hex = "0.3.2"
- base64 = "0.9.3"
- flate2 = "1.0.4"

**What actions are required to be compliant with the IP policy?**

None

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

Keylime meets this definition as it introduces new security protections native to
cloud infrastructure and workload technologies such as containers or virtual
machines. It utilises RESTful interfaces to allow integration with other systems
and seeks to keep all interfaces open.

## Project and Code Quality

**Are there any metrics around code quality? Are there good examples of code reviews? Are there
enforced coding standards?**

- The project has a set of written [**contribution
   guidelines**](https://github.com/keylime/keylime/blob/master/CONTRIBUTORS.md)
- Pull requests are automatically tested using codacy to verify code formatting
- Unit tests and integration tests are performed as CI test on each pull request
- There is a [**security reporting process**](https://github.com/keylime/keylime#report-a-security-vulnerability)
- All pull requests require 1 review before merge.

**What are the performance goals and results? What performance tradeoffs have been made? What is the
resource cost?**

- Keylime uses an asynchronous non blocking I/O networking library to allowing it
  to monitor thousands of instances.
- Keylime comes with its own benchmarking scripts and some early test results
  are available [in this paper](https://github.com/keylime/keylime/blob/master/doc/tci-acm.pdf),
  keylimesection 5.3 "Scalability of Cloud Verifier"

**What is the CI/CD system? Are there code coverage metrics? What types of tests exist?**

- Travis CI runs unit and integration tests against every pull request. An
  example run can be seen on a previous [pull request](https://travis-ci.org/github/keylime/keylime/builds/686611697?utm_source=github_status&utm_medium=notification)

**Is there documentation?**

[Documentation](https://keylime-docs.readthedocs.io/en/latest/) is hosted on
readthedocs.io and is built from the [keylime-docs](https://github.com/keylime/keylime-docs)
repository

**How is it deployed?**

- Keylime can be deployed via the following methods
  - Python setuptools
  - Using Fedora / RHEL `dnf` packaging tool
  - Using the debian `apt` packaging tool
  - Using the `ansible` deployment tool
  - Kubernetes style deployment yaml files are due soon.

**How is it orchestrated?**

A user typically makes use of the `keylime_tenant` CLI application to call
keylimes Restful APIs to request the verifier to measure a node. The node itself
is required to have the agent deployed first.

**How will the project benefit from acceptance into the CNCF?**

Keylime will benefit from CNCF positioning the project as a Remote Attestation
trust project for cloud native technologies. It would also benefit from having
the Governance and support available from being a CNCF project. It will also
help to attract more developers and users to the project and help foster
collaboration with other CNCF projects.

**Has a security assessment by the security SIG been done? If not, what is the status/progress of
the assessment?**

There has not been a security assessment at time of writing, but the SIG Security
said that the project is in scope and there was a growing interest in this space.

# Promotion to Incubation

## Open Governance

**How are committers chosen?**

Commits are accepted through PRs, reviewed and moderated by the [**members of the keylime org**](https://github.com/orgs/keylime/people).

**How are architectural and roadmap decisions made?**

- We have a newly bootstrapped [enhancements proposal repository](https://github.com/keylime/enhancements)
  based on the same process used by the CNCF Kubernetes Project.

**How many decision makers are outside the sponsoring organization.**

At time of writing, all maintainers are outside the sponsoring organization.

## Adoption

**Who are the current maintainers?**

The full list of current maintainers is
[**here**](https://github.com/orgs/keylime/people)

**How long has the project been developed for?**

The keylime project was initially conceived in early 2016 and was published as
an open source in October 2016.

**Is there a commercial version of the project or a primary commercial sponsor ?**

Keylime has been conceived as an open-source initiative, and has no commercial
versions or associated services as of time of writing

**Is the project used in production? If so, please list some of the accounts.**

Keylime is used in the [Mass Open Cloud](https://massopen.cloud).

**Does the project participate in a CNCF User Group?**

keylime does not participate in any CNCF User Groups today, but the project is
actively seeking further community engagement so this will be considered.

## Vendor Independence

**Is the project reasonably independent from the sponsoring vendor?**

As per time of writing, the following organizations have made contributions to
the project.

- MIT - Massachusetts Institute of Technology
- Red Hat
- IBM
- Intel
- Netflix
- Canonical
- ZTE

**Are all communication channels and project resources hosted just for this project or with other
CNCF projects/resources?**

This project currently uses [gitter.im](https://gitter.im/keylime-project/community)
for community chat. We are open to adopt CNCF preferred options such as slack.

**Is all code that is part of the project hosted and part of the CNCF managed orgs and repos?**

Yes, the intention is to donate all of the project's existing GitHub repositories to CNCF.

**Are all defaults for upstream reporting either unset or community hosted infrastructure (i.e.
doesn’t point to vendor hosted SaaS control plane or analytics server for usage data)? Is all
project naming independent of vendors?**

All reporting is managed through GitHub PRs and issues. There is no requirement to interact with any
vendor-hosted services either to consume keylime or to contribute to it.

**Relevant Assets regarding vendor independence**
N/A
