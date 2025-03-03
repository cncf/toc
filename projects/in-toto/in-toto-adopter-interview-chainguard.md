# In-toto Adopter Interview - Chainguard

Interviewee: Billy Lynch, Staff Software Engineer at Chainguard
Interview date: Dec 12, 2024

## Organization Intro

Chainguard is the safe source for open source, so everything we produce includes supply chain metadata like signatures, attestations, SBOMs, etc - in-toto is foundational spec for how we represent much of that data.

## How long has your organization used the project?
Before I joined Chainguard, I was involved with in-toto as an attestation framework in various different projects (SLSA, Sigstore). I met many of the Chainguard founders through these open source projects and when I joined we continued to make heavy use of those projects, and as a result in-toto as well. We are heavy users of in-toto today, mostly the attestation spec.

## What were the main motivations to adopt the project and which key features do you use today?

The key feature is primarily the attestation format (e.g. the envelope and statement specs). In-toto provides a common format to describe attestation metadata like SBOMs and SLSA provenance so it is a nice and important compatibility layer for us.
In-toto also makes it easy for us to develop our own attestation predicates as well - this allows us to provide additional attestation data specific to how our images are built.

## Compared with other products and projects in this space (proprietary and open) what drew you to the project?

We were already working with the adjacent ecosystems (e.g. Sigstore, SLSA), so we wanted to stay in-line with the ecosystems we were already working with - in-toto was a natural choice for us.

## What is the current level of usage (pre-production, production) and scale?

We currently have over 1000 images in our catalog -  we publish attestations for every image we produce which include SLSA provenance, SBOMs, and our own attestation predicate.

## What version of the project is currently in use and what is your update cadence with the project?

Because we were early adopters, we are still using the pre-v1 format, but I don't believe there has been much change from pre-v1 to v1 for the pieces we are using.  If anything this shows that the spec has been fairly stable. :)
We can likely update to v1, but there hasn't been a strong driving force for it at the moment.

## Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

What was really nice about in-toto attestations are the existing specs and examples users can see in the repo. These are really useful for us to see what types of metadata users already exist so we don't need to reinvent the wheel. And if the existing specs don’t fit our needs, we can define our own predicate format but still be compatible with the rest of the in-toto attestation spec.

We've used this for our own images to define an image configuration attestation, which is similar to an SBOM but contains more information about how exactly we built an image. We were able to build this on top of in-toto’s existing format which is very powerful.

Re challenges: Not specific to in-toto, but one of the challenges in the broader ecosystem is predicate fragmentation (e.g. what kind of SBOM predicate to use). I like that in-toto doesn't try to be the one-standard-to-rule-them-all, but gives us basic building blocks to define common concepts and is flexible enough to allow different predicates and easy extensions. There are many adjacent projects working to solve these problems, and it's something we keep an eye on.

## Did you find the information in the repo or the project docs valuable to your implementation? If so, where did you find the information and what specifically was useful?

The spec doc in the attestation repo is very thorough! Love the examples and the detailed spec definitions. Overall the docs are very useful and helpful.

## Did you need to engage with the community members or maintainers? If so what was the context of the engagement, which communication channels did you use and did it reach an acceptable outcome?

Pre-Chainguard, I was helping work on the SLSA build provenance format, which often had discussions with the in-toto community for the best way to define the spec in a way that was complementary to the rest of in-toto.  These conversations often happened on slack (k8s or openssf) or conferences like KubeCon or OSSummit. 
We still work with the maintainers today across various projects in the ecosystem - It's been great to work with all the maintainers, both on in-toto but also adjacent projects as well!

We still work with the maintainers today across various projects in the ecosystem - It's been great to work with all the maintainers, both on in-toto but also adjacent projects as well!

## Has your implementation of the project provided measurable value? Such as reducing manual activities, faster integrations, supported federation/multi-cloud, ease of use, cost savings, etc.

SBOM and attestations have increasingly become table stakes, particularly for users in regulated environments.  Having standard ways to lay out this information with the artifacts we produce is very useful to us. This helps improve the supply chain security for our users by enabling transparency into our artifacts.

## If the project were to be archived now or in the future, what level of difficulty would your organization experience to remove it from your environments? If that were to happen, would you fork and maintain the project to keep functionality, step into a maintainership role within the project, or something else?

Because the spec already exists, even if the maintenance stops, the spec is still there for us to use. The existing spec has been fairly stable for us, and we don't anticipate that changing.

In the event there did need to be a V2 version of the spec, we would likely be interested in being involved to help shape that spec, or at the very least keeping an eye on things and be part of the community.

## Is there something you feel that holds the project back from reaching its ultimate potential?

In many ways, I like where the project is today, at least for attestations. It serves as a common baseline and doesn’t try to be too prescriptive about predicates. 

There are complementary projects that are often intertwined with in-toto such as DSSE (Dead Simple Signing Envelope) and SLSA, but are different enough that I'm fine with them being independent and letting them grow separately. Similar groups of maintainers often contribute across these various projects, so there's been a strong ecosystem behind them making sure they work together well.

## In your opinion, what could the project improve?

More examples in the wild are always appreciated! In the attestation repo, there are folders detailing all the common predicates, these are super interesting to look at and check out how others are using in-toto.

## What are the overall strengths of the project?

In-toto helps define the fundamentals for attestations, but they aren't too prescriptive about exact predicate formats and what data needs to be included. This helps give a common baseline for supply chain metadata, but still gives the flexibility for projects and companies to define their own custom types to fit their needs.

It's also been great to see the cross-industry and academic collaboration for in-toto and other related projects - it's a large community effort.

## Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

Honestly, I'm fairly happy with where things are! I'm sure there will continue to be discussions about various predicate types and the best ways to represent supply chain attestation data, but that isn't necessarily specific to in-toto.

There may be more predicate types we end up using, or creating ourselves in the future!
