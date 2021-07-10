# in-toto Incubating Stage Review

in-toto is currently a CNCF sandbox project. Please refer to in-toto's initial
[sandbox proposal](../proposals/sandbox/in-toto.adoc) for discussion on in-toto's
alignment with the CNCF and details on sandbox requirements.

In the time since being accepted as a sandbox project, in-toto has demonstrated
healthy growth and progress. The [Python reference implementation](https://github.com/in-toto/in-toto) has had several releases.

* [v1.1.0 is the latest minor release](https://github.com/in-toto/in-toto/releases),
  shipped on April 30th, 2021. It added:
    * SPDX License identifiers and copyright information ([#440](https://github.com/in-toto/in-toto/pull/440))
    * Aditya Sirish (@adityasaky) as a maintainer ([#443](https://github.com/in-toto/in-toto/pull/443))

  Other changes include:
    * PyPI development status from `Beta` to `Production/Stable` ([#447](https://github.com/in-toto/in-toto/pull/447))
    * Santiago Torres-Arias's (@SantiagoTorres) email to reflect Purdue affiliation ([#446](https://github.com/in-toto/in-toto/pull/446))
    * Debian downstream release metadata ([#437](https://github.com/in-toto/in-toto/pull/437))
    * Bump dependency: cryptography ([#442](https://github.com/in-toto/in-toto/pull/442))

  Finally, this release removed:
    * Support for Python 2.7 ([#438](https://github.com/in-toto/in-toto/pull/438))

* [v1.0.1](https://github.com/in-toto/in-toto/releases),
  shipped on March 1st, 2021. This was the final in-toto release that supported Python
  2.7. This release added:
    * Python 3.9 in the CI test matrix ([#419](https://github.com/in-toto/in-toto/pull/419))
    * Logo and other visual enhancements on readthedocs ([#420](https://github.com/in-toto/in-toto/pull/420), [#428](https://github.com/in-toto/in-toto/pull/428))
    * Review of first evaluation period for 2021 roadmap ([#421](https://github.com/in-toto/in-toto/pull/421))

  The changes include:
    * Switch to GitHub Actions for CI ([#432](https://github.com/in-toto/in-toto/pull/432))
    * Switch to only running bandit on Python versions greater than 3.5 ([#416](https://github.com/in-toto/in-toto/pull/416))
    * Debian downstream release metadata ([#418](https://github.com/in-toto/in-toto/pull/418))
    * Bump tested dependencies: cffi ([#415](https://github.com/in-toto/in-toto/pull/415), [#427](https://github.com/in-toto/in-toto/pull/427)), cryptography ([#424](https://github.com/in-toto/in-toto/pull/424), [#429](https://github.com/in-toto/in-toto/pull/429)), securesystemslib ([#430](https://github.com/in-toto/in-toto/pull/430), [#431](https://github.com/in-toto/in-toto/pull/431)), iso8601 ([#423](https://github.com/in-toto/in-toto/pull/423))
    * NOTE: the latest version of cryptography is no longer used on Python 2, as that is not supported.

  This release removed:
    * Dropped support for Python 3.5 ([#419](https://github.com/in-toto/in-toto/pull/419))

* [v1.0.0 is the latest major release](https://github.com/in-toto/in-toto/releases),
  shipped on November 23rd, 2020. This release added:
    * '-P/--password' (prompt) cli argument for in-toto-run/in-toto-record ([#402](https://github.com/in-toto/in-toto/pull/402))
    * in-toto-run link command timeout setting ([#367](https://github.com/in-toto/in-toto/pull/367))
    * API and usage documentation for cryptographic key handling with securesystemslib ([#402](https://github.com/in-toto/in-toto/pull/402), [#408](https://github.com/in-toto/in-toto/pull/408))
    * Artifact recording exclude pattern documentation ([#373](https://github.com/in-toto/in-toto/pull/373), [#405](https://github.com/in-toto/in-toto/pull/405))
    * Test key generation mixin ([#402](https://github.com/in-toto/in-toto/pull/402))
    * 2021 roadmap document ([#381](https://github.com/in-toto/in-toto/pull/381))

  The changes include:
    * Move 'settings' docs to new 'configuration' section and make minor enhancements in structure and content ([#405](https://github.com/in-toto/in-toto/pull/405))
    * Update tested dependencies ([#377](https://github.com/in-toto/in-toto/pull/377), [#383](https://github.com/in-toto/in-toto/pull/383), [#384](https://github.com/in-toto/in-toto/pull/384), [#386](https://github.com/in-toto/in-toto/pull/386), [#389](https://github.com/in-toto/in-toto/pull/389), [#390](https://github.com/in-toto/in-toto/pull/390), [#394](https://github.com/in-toto/in-toto/pull/394), [#397](https://github.com/in-toto/in-toto/pull/397), [#398](https://github.com/in-toto/in-toto/pull/398), [#400](https://github.com/in-toto/in-toto/pull/400), [#404](https://github.com/in-toto/in-toto/pull/404), [#406](https://github.com/in-toto/in-toto/pull/406), [#409](https://github.com/in-toto/in-toto/pull/409), [#410](https://github.com/in-toto/in-toto/pull/410), [#411](https://github.com/in-toto/in-toto/pull/411))
    * Debian downstream release metadata ([#382](https://github.com/in-toto/in-toto/pull/382))

  This release removed:
    * 'util' crypto module in favor of securesystemslib key interfaces ([#402](https://github.com/in-toto/in-toto/pull/402))
    * Obsolete coveralls.io API call in Travis test builds ([#399](https://github.com/in-toto/in-toto/pull/399))

  And it fixed:
    * Minor docs issues ([#396](https://github.com/in-toto/in-toto/pull/396), [#385](https://github.com/in-toto/in-toto/pull/385), [#395](https://github.com/in-toto/in-toto/pull/395))
    * pylint 2.6.0 ([#387](https://github.com/in-toto/in-toto/pull/387)) and lgtm.com ([#379](https://github.com/in-toto/in-toto/pull/379)) warnings

* [v0.5.0](https://github.com/in-toto/in-toto/releases),
  shipped on July 13th, 2020. New features include:
    * Docs: Major CLI and API documentation overhaul and release ([#341](https://github.com/in-toto/in-toto/pull/341), [#369](https://github.com/in-toto/in-toto/pull/369))
    * Bugfix: Use kwargs in in-toto-run to fix lstrip-paths bug ([#340](https://github.com/in-toto/in-toto/pull/340))
    * Feature: Add option to specify target directory for generated metadata ([#364](https://github.com/in-toto/in-toto/pull/364))
    * Tests: Add Python 3.8 to tested versions ([#339](https://github.com/in-toto/in-toto/pull/339))
    * Tests: Add tmp dir and gpg key test mixins ([#345](https://github.com/in-toto/in-toto/pull/345))
    * Tests: Use constant from securesystemslib to detect GPG in tests ([#352](https://github.com/in-toto/in-toto/pull/352))
    * Tests: Enhance test suite feedback on Windows ([#368](https://github.com/in-toto/in-toto/pull/368))
    * Dependencies: Misc updates ([#342](https://github.com/in-toto/in-toto/pull/342), [#346](https://github.com/in-toto/in-toto/pull/346), [#349](https://github.com/in-toto/in-toto/pull/349), [#350](https://github.com/in-toto/in-toto/pull/350), [#353](https://github.com/in-toto/in-toto/pull/353), [#354](https://github.com/in-toto/in-toto/pull/354), [#356](https://github.com/in-toto/in-toto/pull/356), [#358](https://github.com/in-toto/in-toto/pull/358), [#359](https://github.com/in-toto/in-toto/pull/359), [#362](https://github.com/in-toto/in-toto/pull/362), [#363](https://github.com/in-toto/in-toto/pull/363), [#366](https://github.com/in-toto/in-toto/pull/366))

* [v0.4.2](https://github.com/in-toto/in-toto/releases),
  shipped on January 7th, 2020. New features include:
    * Drop custom OpenPGP subpackage and subprocess module and instead use the ones provided by securesystemslib, which are based on the in-toto implementation and receive continued support from a larger community ([#325](https://github.com/in-toto/in-toto/pull/325))
    * Fix a race condition that caused tests to sporadically fail was already fixed in securesystemslib and is now also available to in-toto ([#282](https://github.com/in-toto/in-toto/pull/282), [secure-systems-lab/securesystemslib#186](https://github.com/secure-systems-lab/securesystemslib/pull/186))
    * Add Sphinx boilerplate and update installation instructions ([#298](https://github.com/in-toto/in-toto/pull/298), [#331](https://github.com/in-toto/in-toto/pull/331))
    * Update misc dependencies ([#317](https://github.com/in-toto/in-toto/pull/317), [#318](https://github.com/in-toto/in-toto/pull/318), [#319](https://github.com/in-toto/in-toto/pull/319), [#320](https://github.com/in-toto/in-toto/pull/320), [#322](https://github.com/in-toto/in-toto/pull/322), [#323](https://github.com/in-toto/in-toto/pull/323), [#324](https://github.com/in-toto/in-toto/pull/324), [#326](https://github.com/in-toto/in-toto/pull/326), [#327](https://github.com/in-toto/in-toto/pull/327), [#328](https://github.com/in-toto/in-toto/pull/328), [#333](https://github.com/in-toto/in-toto/pull/333), [#335](https://github.com/in-toto/in-toto/pull/335), [#329](https://github.com/in-toto/in-toto/pull/329))
    * Update downstream debian metadata ([#311](https://github.com/in-toto/in-toto/pull/311), [#334](https://github.com/in-toto/in-toto/pull/334))

* [v0.4.1](https://github.com/in-toto/in-toto/releases),
  was shipped on Oct 14th, 2019. New features include:
    * Update securesystemslib dependency to v0.12.0 ([#299](https://github.com/in-toto/in-toto/pull/299))
    * Add --version option to CLI tools ([#310](https://github.com/in-toto/in-toto/pull/310))
    * Address linter warnings ([#308](https://github.com/in-toto/in-toto/pull/308))
    * Update downstream debian metadata ([#302](https://github.com/in-toto/in-toto/pull/302), [#305](https://github.com/in-toto/in-toto/pull/305), [#309](https://github.com/in-toto/in-toto/pull/309))

* [v0.4.0](https://github.com/in-toto/in-toto/releases),
  shipped on September 9th, 2019. New features include:

  * Full artifact rule spec compliance ([#269](https://github.com/in-toto/in-toto/pull/269), [#280](https://github.com/in-toto/in-toto/pull/280))
  * Enhanced OpenPGP key export and key expiration verification ([#266](https://github.com/in-toto/in-toto/pull/266), [#288](https://github.com/in-toto/in-toto/pull/288))
  * Transitive PyNaCl dependency is now optional ([#291](https://github.com/in-toto/in-toto/pull/291))
  * Improved automatic test coverage analysis ([#295](https://github.com/in-toto/in-toto/pull/295))
  * Static analysis improvements ([279](https://github.com/in-toto/in-toto/pull/279), [#296](https://github.com/in-toto/in-toto/pull/296))
  * Improve upstream release packaging for Debian and Arch Linux ([#279](https://github.com/in-toto/in-toto/pull/279), [#290](https://github.com/in-toto/in-toto/pull/290))

More changes and small improvements are mentioned in the current release
changelog.

Further, the [Go implementation](https://github.com/in-toto/in-toto-golang) has had one
pre-release, [v0.1.0](https://github.com/in-toto/in-toto-golang/releases).

Beyond the current release other improvements to the broader reference
implementation have been achieved.

* A [formalized governance
policy](https://github.com/in-toto/in-toto/blob/develop/GOVERNANCE.md) has been
instituted project-wide. This includes not only the in-toto python reference
implementation, but the specifications, implementations in other languages and
cloud-native tooling.

## Incubating Stage Criteria

In addition to sandbox requirements, a project must meet the following
criteria to become an incubation-stage project:

### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgment, are of adequate quality and scope

  * We document adopters on
    [our website](https://in-toto.io/integrations.html).
  * The in-toto specification has inspired
    the development of [Argos Notary](https://www.argosnotary.com/docs/00_overview/10_overview)
    by Rabobank.
  * Boxboat is currently [integrating](https://github.com/boxboat/in-toto-golang)
    SPIFFE, a CNCF project, with in-toto. This will be presented at
    [Cloud Native Security Day 2021](https://cnsecuritydayeu21.sched.com/event/iRn2).
  * in-toto is also used by Cloud Native Application Bundles (CNAB), another
    CNCF project, in [Signy](https://github.com/cnabio/signy).

### Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project

  * Maintainers of the project are listed in our
    [MAINTAINERS.txt](https://github.com/in-toto/in-toto/blob/develop/MAINTAINERS.txt)
    file. There are currently five core maintainers from Purdue University, New
    York University, and Conda who cut releases.

  * Several other people have commit access, but cannot craft releases. They are:
    * Holger Levsen (Debian)
    * Ofek Lev (Datadog)

  * We have had contributions from people associated with different organizations.
    Some of them, listed in no particular order:
    * Trishank Karthik Kuppusamy (Datadog)
    * Joshua Lock (VMWare)
    * Dan Lorenc (Google)
    * Christian Rebischke (Arch Linux)

  * Maintainers are added and removed from the project as per the policies
    outlined in the project [GOVERNANCE.md](https://github.com/in-toto/in-toto/blob/develop/GOVERNANCE.md)
    file.

  * Finally, in-toto participated in Google Summer of Code (GSOC) 2020 through the CNCF.
    We have received proposals for GSOC 2021 which we are currently evaluating.

### Demonstrate a substantial ongoing flow of commits and merged contributions

  * Releases: There were six [releases](https://github.com/in-toto/in-toto/releases)
    scheduled since the sandbox application as defined on our
    [release schedule](https://github.com/in-toto/in-toto/blob/develop/ROADMAP.md).

  * Roadmap: We have annual roadmaps for
    [the reference implementation](https://github.com/in-toto/in-toto/blob/develop/ROADMAP.md)
    and [in-toto as a whole](https://github.com/in-toto/docs/blob/master/ROADMAP.md).
    Reviews are released at the end of each evaluation period described in the
    roadmap, and they can be found in the repositories for the
    [reference implementation](https://github.com/in-toto/in-toto/tree/develop/roadmap-reviews)
    and [in-toto docs](https://github.com/in-toto/docs/tree/master/roadmap-reviews).

  * in-toto Enhancements (ITEs): We have a [formal process](https://github.com/in-toto/ITE)
    for interested parties to submit new features or describe some aspect or
    integration of the in-toto specification. Since in-toto was accepted into
    the CNCF sandbox, the following ITEs have been proposed:
    * [ITE-2](https://github.com/in-toto/ITE/blob/master/ITE/2/README.adoc) by Trishank Karthik Kuppusamy (Datadog), "Draft"
    * [ITE-3](https://github.com/in-toto/ITE/blob/master/ITE/3/README.adoc) by Trishank Karthik Kuppusamy (Datadog), "Draft"
    * [ITE-4](https://github.com/in-toto/ITE/blob/master/ITE/4/README.adoc) by Santiago Torres-Arias (Purdue University), "Accepted"
    * [ITE-5](https://github.com/in-toto/ITE/blob/master/ITE/5/README.adoc) by Santiago Torres-Arias (Purdue University), "Draft"
    * [ITE-6](https://github.com/in-toto/ITE/pull/15) by Mark Lodato (Google), discussions ongoing

  * Contributors: [https://github.com/in-toto/in-toto/graphs/contributors](https://github.com/in-toto/in-toto/graphs/contributors)

  * Commit activity: [https://github.com/in-toto/in-toto/graphs/commit-activity](https://github.com/in-toto/in-toto/graphs/commit-activity)

  * CNCF DevStats: [https://intoto.devstats.cncf.io/](https://intoto.devstats.cncf.io/)
    * [Last 30 days activity on GitHub](https://intoto.devstats.cncf.io/d/8/dashboards?refresh=15m&orgId=1&from=now-30d&to=now-1h)
    * [Community Stats](https://intoto.devstats.cncf.io/d/3/community-stats?orgId=1&var-period=d7&var-repo_name=goharbor%2Fharbor)

## Security

Given that in-toto is a security product, in-toto's codebase has been written,
designed and tested with security in mind. Some of the actions performed in
order to ensure the quality and security of the codebase, as well as in-toto's
design and specification include:

* Static analysis is performed using [pylint](https://github.com/PyCQA/pylint/) and [bandit](https://bandit.readthedocs.io/en/latest/)
* Dependency vulnerability tracking using [Dependabot](https://dependabot.com/)
* Manual code analysis / review by a Maintainer for each included piece of
  code
* [Security assessment](https://github.com/cncf/sig-security/blob/master/assessments/projects/in-toto/self-assessment.md) by CNCF's SIG-SECURITY
* A peer-reviewed paper describing the threat model, its
  security properties, was published in
  [USENIX Security '19](https://www.usenix.org/conference/usenixsecurity19/presentation/torres-arias)
* in-toto's implementation has received the [CII Silver Criteria Badge](https://bestpractices.coreinfrastructure.org/en/projects/1523)
  for best development practices

A more elaborated description of these security initiatives, as well as a
vulnerability report process is included [here](https://github.com/in-toto/in-toto#security-issues-and-bugs).
