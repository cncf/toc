# TUF Graduation Application

The Update Framework (TUF) helps developers maintain the security of a software update system, even against attackers that compromise some repository or signing keys. TUF provides a flexible framework and specification that developers can adopt into any software update system.

## TUF fulfills all the incubating and graduation criteria:

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

* TUF is used in production by Cloudflare, Datadog, DigitalOcean, Docker, IBM, Kolide, LEAP, Microsoft, RedHat, VMware, and many others.   Amazon has just released a new Rust implementation of [TUF](https://github.com/awslabs/tough).   [Uptane](https://uptane.github.io/), the automotive variant of TUF, is sold by major automotive vendors and is being adopted by automakers with over 1/3 of the cars on US roads in their forthcoming cars.  Uptane is included as part of [Automotive Grade Linux](https://www.automotivelinux.org/), another Linux Foundation project.  
* While these integrations all use the TUF specification, the implementations vary.  Datadog, OTAinfo, and LEAP use the Python reference implementation, most cloud companies use Notary, Google is integrating their own Go implementation into Fuschia, and most automobiles use custom variants written by different automotive vendors.  
* More detail about adoption, including some of the in progress adoptions, can be found here:  https://theupdateframework.github.io/adoptions.html

Apart from the Docker, LEAP, and AppContainer adoptions (which happened prior to joining the CNCF), these adoptions have all happened while TUF is in incubation.

### Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.  The committers should come from at least two organizations.

As an intentionally minimal security specification / framework, we deliberately do not have a high degree of feature additions in the project.  Effort comes on either the implementations, such as Notary, Amazon's Rust implementation, Google's implementation, or the Python reference implementation, or the specification itself.  As TUF is a specification project, we focus on the specification and the Python reference implementation.

* Python reference implementation / specification (6 committers, 3 organizations)

  * Justin Cappos (NYU), Trishank Karthik Kuppusamy (Datadog / NYU), Lukas Puehringer (NYU), Santiago Torres (NYU), Sebastien Awwad (Conda), Marina Moore (NYU)



### Demonstrate a substantial ongoing flow of commits and merged contributions.

We have had active contributions from an array of contributors across the CNCF landscape, more traditional software projects, and automakers.  One way to see this is via the substantial changes that made their way into the specification.

Changes to the TUF standard largely come in the form of TAPs (TUF Augmentation Proposals).  There are 8 active, non-informational TAPs (TAP 3-6 and 8-11).  

The 5 accepted TAPs have authors from 6 different organizations:
* [TAP3](https://github.com/theupdateframework/taps/blob/master/tap3.md) and [TAP4](https://github.com/theupdateframework/taps/blob/master/tap4.md), by Trishank Karthik Kuppusamy (Datadog / NYU), Sebastien Awwad (Conda), Evan Cordell (CoreOS / RedHat), Vladimir Diaz, Jake Moshenko (CoreOS / RedHat), Justin Cappos (NYU)
* [TAP6](https://github.com/theupdateframework/taps/blob/master/tap6.md) by David Lawrence (Docker), Vladimir Diaz, Justin Cappos (NYU)
* [TAP9](https://github.com/theupdateframework/taps/blob/master/tap9.md) by heartsucker (HERE)
* [TAP10](https://github.com/theupdateframework/taps/blob/master/tap10.md) by Vladimir Diaz

There are 3 TAPs that are currently under consideration (all of which are likely to be accepted) with authors from 5 organizations:
* [TAP5](https://github.com/theupdateframework/taps/blob/master/tap5.md) by Trishank Karthik Kuppusamy (Datadog / NYU), Sebastien Awwad (Conda), Evan Cordell (CoreOS / RedHat), Vladimir Diaz, Jake Moshenko (CoreOS / RedHat), Justin Cappos (NYU)
* [TAP8](https://github.com/theupdateframework/taps/blob/master/tap8.md) by Hannes Mehnert (OCaml / Cambridge University), Justin Cappos (NYU)
* [TAP11](https://github.com/theupdateframework/taps/blob/master/tap11.md) by Marina Moore (NYU), Santiago Torres (NYU), Trishank Kuppusamy (Datadog), Sebastien Awwad (Conda), Justin Cappos (NYU)

There is also a TAPs from the automotive industry dealing with upgrading while changing the TUF metadata version itself that is currently being iterated over.

These demonstrate the breadth of contributions from different adopters and the way in which contributions are merged into the specification.

Each part of TUF has seen activity both from core committers and external parties.  Of course the fewest commits have come in the TAPs and spec itself.   The TAPs have had 575 commits by 10 contributors (5 external, 2 were non-TAP authors).  The specification has had 82 commits by 11 contributors (6 external contributors).  TUF's Python reference implementation has had more than 3.3K commits, by 42 contributors (26 external).   This demonstrates both activity and committer diversity.

### A clear versioning scheme

TUF uses semantic versioning.  We check in a persistent copy of the specification with each release.



### Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

We have both the basic best practices badge and the silver badge.
https://bestpractices.coreinfrastructure.org/projects/1351

### Adopt the CNCF Code of Conduct.

TUF follows the CNCF code of conduct.  https://github.com/theupdateframework/tuf/blob/develop/docs/CODE-OF-CONDUCT.md
This is linked multiple places, including the footer of [TUF's homepage](https://theupdateframework.github.io/).

### Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

These documents may be found here:
* https://github.com/theupdateframework/tuf/blob/develop/docs/GOVERNANCE.md
* https://github.com/theupdateframework/tuf/blob/develop/docs/CONTRIBUTORS.rst

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

Our website has an [adoptions page](https://theupdateframework.github.io/adoptions.html) on it that lists the different projects.  We also have an [ADOPTERS.MD](https://github.com/theupdateframework/tuf/blob/develop/docs/ADOPTERS.md) which contains much of the same information.  [Uptane](https://uptane.github.io/), the automotive version of TUF, also its own [adoptions page](https://uptane.github.io/adoptions.html).

## Security Audits
There are multiple [security audits](https://theupdateframework.github.io/audits.html) of TUF available on the TUF website.

### Alternatives to TUF

The most common alternative to TUF involves using either a signing key on the server (e.g., TLS) or on a server as part of software creation (e.g., GPG signing in a build farm).  Either way, the fundamental difference is that a single key / server compromise can result in an attacker having the ability to install arbitrary code on end user machines.  Existing specifications / proposals like OMA-DM, SUIT, ITU-T X.1373, as well as common use patterns for GPG/PGP/RSA signing and TLS all have this flaw.  

In contrast, TUF is designed with the assumption that the company deploying TUF may not have perfect security.  In other words, a key and/or server will be compromised at some point.  TUF is focused on preventing an attacker from performing harmful attacks while allowing the operator to recover to a secure baseline.```
