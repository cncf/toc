# Parsec 2021 Annual Review

## Project Summary

Parsec is the Platform Abstraction for Security. Parsec is creating convenient and portable interfaces to hardware security on any platform in any programming language. Parsec is deployed as a local software service (or daemon), which executes in user space on Linux systems. The software service contains back-end provider modules that offer compatibility with various hardware or firmware services for secure storage and cryptography, including Trusted Platform Modules (TPM), Hardware Security Modules (HSM), secure elements or trusted firmware. Applications consume these services using a convenient and portable API. The Parsec API can be made available in any programming language, given the availability of a suitable client library. The end result is that applications are able to consume the best-available hardware security of their host platform, without being tightly coupled to the often complex and specialised APIs and software stacks that are otherwise needed.

The main Parsec repository can be found [here](https://github.com/parallaxsecond/parsec). There are several satellite repositories in the same organisation.

You can read the [Parsec documentation online](https://parallaxsecond.github.io/parsec-book/).

Parsec was [proposed](https://github.com/cncf/toc/issues/442) and [accepted](https://lists.cncf.io/g/cncf-toc/message/4853) as a sandbox project in June 2020.

## Functional Highlights

Some functional development highlights for the reporting period 2020/2021 as follows:

- Introduction of the [Parsec Go client](https://github.com/parallaxsecond/parsec-client-go) through independent community contribution.
- Introduction of the [Parsec Tool](https://github.com/parallaxsecond/parsec-tool), allowing users to access Parsec functions from command-line or scripting contexts.
- Introduction of a new back-end provider module interfacing with [CryptoAuthLib](https://github.com/MicrochipTech/cryptoauthlib) for compatibility with [Microchip Secure Elements](https://www.microchip.com/en-us/products/security-ics/cryptoauthentication-family).
- Introduction of a new back-end provider module for compatibility with firmware trusted services from [trustedfirmware.org](https://www.trustedfirmware.org/projects/trusted-services/).
- Development of multi-tenancy support, including an integration with the [SPIFFE](https://spiffe.io) identity framework.
- Creation of a new, safe Rust abstraction for the [PKCS#11](https://github.com/parallaxsecond/rust-cryptoki) and [PSA Crypto](https://github.com/parallaxsecond/rust-psa-crypto) interfaces.
- Substantial enhancements to the [Rust TPM Library](https://github.com/parallaxsecond/rust-tss-esapi).

## Contribution Metrics

Please refer to the [Parsec DevStats page](https://parsec.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m).

Parsec has enjoyed a modest but increasing contribution rate during its first year as a CNCF sandbox project. The number of contributors has grown, as has the number of contributing organizations.

Across the individual repositories that make up the project, a contribution summary for the last review period is given below:

- The number of individual GitHub repositories has grown from 9 to 14 as new client libraries and back-end provider integrations have been contributed.
- A total of [559 pull requests](https://github.com/pulls?page=1&q=is%3Amerged+is%3Apr+user%3Aparallaxsecond+archived%3Afalse+merged%3A%3E%3D2020-07-01) were successfully merged from 19 separate contributors.
- The number of contributing organisations has grown from three to six. The project has also seen some independent contributions in the community.

## Maintainers

Please refer to the [MAINTAINERS](https://github.com/paulhowardarm/parsec/blob/master/MAINTAINERS.toml) file.

The list of official maintainers covers three organisations: Arm, Docker and Mirantis. These are the individuals who have full admin rights to the project repositories. However, during the reporting period, an additional three contributors were granted the **write** role on some repositories, giving them the permission to merge pull requests.

## Adoption

Parsec has been adopted commercially by [Pelion](https://pelion.com). The Pelion Edge product uses Parsec as a portable means of protecting device identity keys in hardware. See the relevant [release notes](https://developer.pelion.com/docs/device-management-edge/2.3/release-notes/index.html).

Fedora IoT has the [dbus-parsec](https://github.com/fedora-iot/dbus-parsec) project in GitHub, which is the first community open-source project to be consuming Parsec.

The CNCF [Keylime Project](https://keylime.dev) is making use of the Rust TPM library that was initially developed for Parsec.

Parsec has now been packaged for both the [Fedora](https://koji.fedoraproject.org/koji/packageinfo?packageID=32320) and [openSUSE](https://build.opensuse.org/package/show/openSUSE:Factory/parsec) Linux distributions. It is included by default in the IoT variant of the Fedora distribution.

Parsec recipes are available for custom embedded Linux distributions based on [Yocto Project](https://git.yoctoproject.org/cgit/cgit.cgi/meta-security/tree/meta-parsec).

## Events and Publicity

During the reporting period, Parsec has been represented or showcased as follows:

- [Cloud Native Security Day (Europe)](https://youtu.be/bYFQXcPSf0I)
- [Linaro Virtual Connect](https://youtu.be/GqiISmXO_78)
- [Cloud Native Security Day (NA)](https://youtu.be/-I_rCKMyY7Y)
- [Cloud Native Rust Day](https://youtu.be/49cXCDLALYY)
- [KubeCon 2021](https://youtu.be/G-MvFqkVJTI)
- [Yocto Summit 2021](https://www.youtube.com/watch?v=030ZVCTxaIg)
- [Confidential Computing Developer Summit](https://youtu.be/mmTpzRVeSoQ)
- [Arm Software Developer Breakdown](https://www.youtube.com/playlist?list=PLKjl7IFAwc4S7WQqqphCsyy6DPDxJ2Skg)

## Project Goals

### Performance Against Goals For Previous Review Period

The Parsec project has tracked well against the goals that were established when it was proposed for the CNCF sandbox. In particular:

- Rust and C client libraries have been made available as planned. There is also a Go client library, which was developed through community contribution.
- Integration into a commercial product has been achieved (see above).
- Integration into Yocto Linux has been achieved (see above), as well as packaging for Fedora and openSUSE distributions.
- Coverage of the API primitives has increased, as has the number of back-end integrations, which now offer compatibility with secure elements using [CryptoAuthLib](https://github.com/MicrochipTech/cryptoauthlib) and services from [Trusted Firmware](https://www.trustedfirmware.org/projects/trusted-services/).
- The multi-tenancy design was implemented in two variants, allowing client applications to be identified according to their Unix user identity (UID), or alternatively according to a [SPIFFE ID](https://spiffe.io).

The following are still backlog items:

- Support for Parsec in virtualised environments.
- API for arbitrary blob/secret storage and retrieval.

These are still both project goals, although they have not yet been of sufficiently high priority to be addressed by any of the contributors.

### Current Goals

Goals for the coming year include:

- Establishing a formal release process and regular release cadence. Read the proposal [here](https://parallaxsecond.github.io/parsec-book/contributing/release_process.html).
- Establishing a public project roadmap.
- Increasing the number of contributors and contributing organisations.
- Increasing the amount of end-user adoption.
- Meeting the criteria for incubation.
- Functional enhancements as needed in support of the above goals. Support for additional client languages and popular security frameworks will be a likely component of this.

## CNCF Membership

Parsec contributors are grateful to the CNCF for the first year of sandbox membership.

Parsec has benefitted enormously from CNCF membership. CNCF provided invaluable assistance during the onboarding phase, helping with IP and trademark issues. Parsec's presence in the CNCF sandbox has made it more attractive to potential adopters. CNCF have also assisted with website content creation and hosting.

Parsec will continue to look towards the CNCF for guidance in preparing itself for incubation. Assistance with security assessment (if required) would be appreciated, as would recommendations for how to ensure that incubation criteria are being met.

## Readiness for Incubation

Parsec does not yet meet the [stated criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage) for incubation for the following reasons:

- There are not yet three independent adoptions that are in production.
- Security assessment still pending.
- The release process and associated versioning scheme have been defined but not yet exercised in practice.

The following maturity indicators can however be demonstrated:

- Contributions from multiple independent vendor organisations.
- A body of comprehensive and well-maintained [documentation](https://parallaxsecond.github.io/parsec-book/index.html).
- A consistent and increasing rate of contribution.
- A [security page](https://github.com/parallaxsecond/parsec/security), which describes the project's security policy, and provides information about how to report vulnerabilities to the maintenance team.
- A continuous integration process, including fuzz testing and code coverage analysis.
- Use in production by one commercial software product/service.
- A [documented security threat model](https://parallaxsecond.github.io/parsec-book/parsec_security/rust_client_threat_model/threat_model.html).
- A [documented set](https://parallaxsecond.github.io/parsec-book/parsec_service/stability.html) of stability and backwards-compatibility requirements, with appropriate build-time enforcement.
- A [pull request checklist](https://parallaxsecond.github.io/parsec-book/contributing/pr_checklist.html).
- A [CII Best Practices badge](https://bestpractices.coreinfrastructure.org/en/projects/4856).

