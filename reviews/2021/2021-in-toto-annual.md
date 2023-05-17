# in-toto 2021 Annual Review

## Background

in-toto is a framework to secure software supply chains. Today, software supply chains are complex processes that are vulnerable to attack. They comprise of multiple systems working together like version control, build farms, CI/CD, and so on. in-toto provides a specification and a family of tooling that enables the generation of cryptographically verifiable provenance for artifacts in the supply chain.

## DevStats

in-toto's statistics can be found [here](https://intoto.devstats.cncf.io/d/8/dashboards?orgId=1&from=1598918400000&to=1630540799000).

## Maintainers

The in-toto project's maintainers can be found [here](https://github.com/in-toto/in-toto/blob/develop/MAINTAINERS.txt).

Further, its governance model is available [here](https://github.com/in-toto/in-toto/blob/develop/GOVERNANCE.md).

in-toto's primary implementation has five core maintainers from Purdue University, New York University, and Conda. Further, Holger Levsen of Debian and Ofek Lev of Datadog have commit access but cannot craft releases. Ofek Lev helps maintain the git-specific semantics within in-toto, while Holger Levsen helps package in-toto and the [apt-transport](https://github.com/in-toto/apt-transport-in-toto) for Debian releases.

Modifications to the in-toto specification are proposed via in-toto Enhancements (ITEs). The ITE process is detailed in [ITE-1](https://github.com/in-toto/ITE/blob/master/ITE/1/README.adoc). Each ITE goes through a discussion process, and ITE editors are responsible for deciding the feasibility of a proposal, and whether it is right for in-toto. The current ITE editors are:
* Santiago Torres-Arias (Purdue University)
* Trishank Karthik Kuppusamy (Datadog)
* Justin Cappos (New York University)

## Goals / Roadmap

The in-toto project publishes annual roadmaps detailing the goals for the year. Until August 2021, separate roadmaps were published for the reference implementation and for the organization as a whole, but starting from the current period, the project has switched to one roadmap for everything in-toto.

Roadmap reviews are published thrice a year based on a predetermined schedule listed in the roadmap(s). These are the mechanisms the core team uses to update the in-toto community with the latest progress and changes. Typically, about four community meetings are also held in any given year, though we are currently evaluating the feasibility of expanding to more meetings.

Our roadmap reviews for the last year show that we met all our goals except one. We are optimistic that a year from now, we will have similar reports for our current roadmap. This last year, the goal we missed was the CII Best Practices gold badge, but that was because the reference implementation matches the current in-toto specification closely, and has therefore not had significant contributions from people external to or unassociated with the in-toto project.

The last roadmap period ended in July 2021. Our current roadmap applies to the
period beginning August 1st 2021 to July 31st 2022.

### Documents

* [Last organization roadmap](https://github.com/in-toto/docs/blob/master/roadmap-reviews/2021/ROADMAP.md)
* [Organizational roadmap reviews for 2020-2021](https://github.com/in-toto/docs/tree/master/roadmap-reviews/2021)
* [Last implementation roadmap](https://github.com/in-toto/in-toto/blob/develop/roadmap-reviews/2021/ROADMAP.md)
* [Implementation roadmap reviews for 2020-2021](https://github.com/in-toto/in-toto/tree/develop/roadmap-reviews/2021)
* [Current roadmap](https://github.com/in-toto/docs/blob/master/ROADMAP.md)

### Significant Milestones

* We released the first stable version, v1.0, of in-toto's reference implementation written in Python.
* We saw increased ITE activity from external contributors, from companies like [Boxboat](https://github.com/in-toto/ITE/pull/21), [Datadog](https://github.com/in-toto/ITE/blob/master/ITE/3/README.adoc), and [Google](https://github.com/in-toto/ITE/tree/master/ITE/6). This is significant because it reflects the increasing adoption of in-toto in the industry.
* We also observed increased code contributions from the industry, specifically from companies like [Boxboat](https://github.com/in-toto/in-toto-golang/pull/119), [Google](https://github.com/in-toto/in-toto-golang/pulls?q=is%3Apr+is%3Aclosed+author%3Adlorenc+), and [SolarWinds](https://github.com/in-toto/in-toto-golang/pull/101).
* We completed our [second year of Google Summer of Code (GSoC)](https://coda.io/@joy/2021-gsoc-story) successfully.
* in-toto was one of the recommended tools in CNCF's [Software Supply Chain Security Paper](https://github.com/cncf/tag-security/tree/main/supply-chain-security/supply-chain-security-paper).

### Help from CNCF

One of our primary aims is to increase in-toto's visibility and adoption, especially through more integrations with projects in the ecosystem. We are also tracking possible usability improvements that can move these along. The in-toto project would benefit tremendously from support from the CNCF towards achieving these aims.

Some specific usability items we could use help with:

* Improved in-toto layout creation user experience
* Further development of in-toto tooling to better support [SLSA](https://slsa.dev/) and [in-toto Attestations](https://github.com/in-toto/attestation).

## Adoption

* We list adopters on the [in-toto website](https://in-toto.io/integrations/).
* in-toto has inspired the development of "[Argos Notary](https://www.argosnotary.com/docs/00_overview/10_overview/)" by Rabobank.
* Boxboat is implementing SPIFFE, another CNCF project, with in-toto. This integration is documented in [ITE-7](https://github.com/in-toto/ITE/pull/21).
* in-toto is being used by Cloud Native Application Bundles (CNAB), another CNCF project, in [Signy](https://github.com/cnabio/signy).
* Following the completion of GSoC 2021, in-toto link metadata can be generated in [rebuilderd](https://github.com/kpcyrd/rebuilderd).
* In the current implementations and demonstrations, the [in-toto Attestation](https://github.com/in-toto/attestation) project is the metadata vehicle of choice for [SLSA](https://slsa.dev) requirements.
* [Tekton Chains](https://github.com/tektoncd/chains) is another example of in-toto integration. Chains can generate in-toto metadata, capturing information from the pipeline.
* The [Sigstore](https://www.sigstore.dev/) project also supports in-toto metadata, though a tighter integration is under active development.
* The [Qubes OS](https://www.qubes-os.org/) project uses in-toto with our apt-transport within [their reproducible builds](https://www.qubes-os.org/news/2021/02/28/improvements-in-testing-and-building/) setup.

## Incubation

We believe in-toto is ready for incubation. We have had discussions on this matter, and our [incubation review](https://github.com/cncf/toc/pull/393) is currently under discussion.
