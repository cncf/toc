# Application for Graduation: Flux

- [x] Have committers from at least two organizations.

  Yes. The Flux project has committers from at least two organisations.
  - The full list of Flux project maintainers can be found [here](https://github.com/fluxcd/community/blob/main/project/flux-project-maintainers.yaml). The current breakdown (commit 22fd87b, 2022-02-28) is: Weaveworks: 14, CrowdStrike: 1, D2iQ: 1, Independent: 1, Microsoft: 1, NexHealth: 1, Tetrate: 1, Xenit: 1
  - [`flux2` (core) maintainers](https://github.com/fluxcd/flux2/blob/main/MAINTAINERS): Weaveworks: 3, D2iQ, NexHealth, Xenit each: 1.

- [x] Have achieved and maintained a Core Infrastructure Initiative Best Practices Badge.

  Yes. See links for Core Infrastructure Initiative Best Practices Badge.
  - flux2: <https://bestpractices.coreinfrastructure.org/projects/4782>
  - helm-controller: <https://bestpractices.coreinfrastructure.org/en/projects/4784>
  - notification-controller: <https://bestpractices.coreinfrastructure.org/en/projects/4791>
  - image-reflector-controller: <https://bestpractices.coreinfrastructure.org/en/projects/4790>
  - image-automation-controller: <https://bestpractices.coreinfrastructure.org/en/projects/4789>
  - kustomize-controller: <https://bestpractices.coreinfrastructure.org/en/projects/4787>
  - source-controller: <https://bestpractices.coreinfrastructure.org/en/projects/4786>
  - flagger: <https://bestpractices.coreinfrastructure.org/projects/4783>

- [x] Have completed an independent and third party security audit with results published of similar scope and quality as this example which includes all critical vulnerabilities and all critical vulnerabilities need to be addressed before graduation.

  Yes. The Flux project completed the audit and published the [report Nov 2021](https://fluxcd.io/blog/2021-11-10-flux-security-audit/).
  We have addressed all discovered issues:
  - 1 High (assigned a CVE and fixed during the audit)
  - 3 Medium, 13 Low, 5 Informational issues.
  - The entirety of issues can be found in GitHub, here is our project board: <https://github.com/orgs/fluxcd/projects/5> (as of 2021-03-02: 100% DONE).

    Another important piece of feedback in the audit was that the proposal for a general implementation of multi-tenancy in Flux (f/flux2#582) contained many good ideas, but that it would benefit from being broken up into more targeted pieces. We are using our [new RFC process](https://github.com/fluxcd/flux2/tree/main/rfcs) now to [flesh these out individually](https://github.com/fluxcd/flux2/pulls?q=label%3Aarea%2FRFC+) and are looking into working with security experts on integrating this. TBA once the plan has solidified.

- [x] Explicitly define a project governance and committer process. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

  Yes. The Flux project set up [our Governance process](https://fluxcd.io/governance/) about a year ago, and revisited it as a team after the half-year mark. We are [aware](https://github.com/fluxcd/community/issues/106) that the Flux Oversight Committee currently has 3 members from one company and it is a priority to the Flux project to have diverse company representation among its maintainers and contributors. Please note these key points:
  - The oversight committee is a remnant from bootstrapping project governance, with the original maintainers
  - In the near term we expect a change to the companies represented in the committee due to a job move in April, but this will not make much difference in practice;
  - Therefore: we are looking into ways to involve more companies and individuals in an active, steering capacity

  Asking for maintainer permissions, access to any other resources or being added as an adopter or [support-providing company](https://fluxcd.io/support/) is all done transparently and publicly via GitHub.

  In addition to that:
  - We created the <https://github.com/fluxcd-community> organisation to allow for more experimentation within our community.
  - We run regular and public dev meetings
  - We feel we learned lessons from other CNCF projects: community roles, decision making process (+ Plan B), maintainer applications
  - One outcome after 1 year of governance: [clear RFC process](https://github.com/fluxcd/flux2/tree/main/rfcs) for big architectural decisions.

- [x] Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website). For a specification, have a list of adopters for the implementation(s) of the spec.

  Yes. We have <https://fluxcd.io/adopters> where folks can add themselves easily by editing a YAML. We divided the sections into v2, Flagger and v1 users.
  - Sandbox 2019-08-15: 26
  - Incubation 2021-03-10: 78
  - Today 2022-03-07: 128

## About your Project

Project name: Flux

Flux is a set of continuous and progressive delivery solutions for Kubernetes that are open and extensible. Flux provides GitOps for both apps and infrastructure, works with users' existing tools (eg. Jenkins, Helm, etc.),  and works with any Kubernetes and all common Kubernetes tooling (eg. Kubernetes, AKS, EKS, etc.). Flux does Multi-Tenancy (and "Multi-everything"): Flux uses true Kubernetes RBAC via impersonation and supports multiple Git repositories. Multi-cluster infrastructure and apps work out of the box with Cluster API: Flux can use one Kubernetes cluster to manage apps in either the same or other clusters, spin up additional clusters themselves, and manage clusters including lifecycle and fleets.

Flux is trusted by enterprise users in production (eg. SAP, RingCentral, SoulCycle, Starbucks, Under Armour, etc.) and major integrators trust Flux to delivery GitOps to their end customers (eg. Microsoft, AWS, D2iQ, VMware, etc.).

Link to project website with summary of what the project does: <https://fluxcd.io/>

Link to GitHub repos:

- <https://github.com/fluxcd/flux2>
- <https://github.com/fluxcd/image-automation-controller>
- <https://github.com/fluxcd/source-controller>
- <https://github.com/fluxcd/helm-controller>
- <https://github.com/fluxcd/image-reflector-controller>
- <https://github.com/fluxcd/notification-controller>
- <https://github.com/fluxcd/kustomize-controller>
- <https://github.com/fluxcd/flagger>

Link to Incubation doc: <https://github.com/cncf/toc/blob/main/reviews/incubation-flux.md>

Due Diligence (Incubation): <https://docs.google.com/document/d/1Z6yPN9-bWeVGpMrBxXJ3NBTBYZowJ3R93wKHEVmBJ1A/edit>

## Why is the Project ready to Graduate?

The Flux maintainers have designed the project to meet enterprise needs and standards, and the quality of users and integrators attests to this. The quality of the project has been further validated by how Flux fared in the recent, independent  security audit through the CNCF. Flux is designed and maintained with security as a priority. Enterprise-level companies continue to engage in the community slack as well as add themselves to the Adopters page (Eg. SAP and RingCentral are recent ones). Major integrators such as Microsoft and AWS trust Flux for their GitOps offerings to end users. Microsoft's Arc Kubernetes and AWS' EKS Anywhere both leverage Flux to deliver quality GitOps services to their customers. Several others from VMware, D2iQ, and Red Hat have testified to Flux's reliability and they have contributed to Flux as well around this. Contributors include the active community members on slack who help each other and answer questions.

Part of this is because the maintainers and contributors have designed Flux to work closely, sometimes natively, with critical projects such as Helm, Prometheus, and most recently HashiCorp Vault, to meet the needs of the user community. This effort has received validation with statements from the Department of Defense that Flux's seamless integration with Helm is what led them to choose Flux.

Please see highlights below as well as stats on how the Flux community continues to grow.

- **Security: Mature model and community transparency around security**: eg. designing Flux to rely on Kubernetes RBAC, including features such as verifying commits, encrypting secrets, and designing Flux UI as a project outside of core Flux. RFC process (visibility and accountability).
- **CNCF ecosystem**: Flux and Flagger are designed to work seamlessly with projects such as Helm and Prometheus.
- **Platforms / Integrators have trusted Flux as the default technology to provide CD / GitOps to their customers**: Microsoft (Azure Arc Kubernetes), AWS (EKS Anywhere), VMware (Tanzu Application Platform), D2iQ (D2iQ Kubernetes Platform), Weaveworks (Weave GitOps).
- **Flux has been trusted to build software on it**: Gimlet.io, Kraan, Keploy (previously hybridk8s - UI built on Flux and Flagger), Weave GitOps Enterprise
- **Flux supports major deployment sizes by its adopters**:
  - [Department of Defense with 100k devs](https://www.cncf.io/blog/2021/09/30/how-to-get-robust-gitops-the-u-s-department-of-defense-uses-flux-and-helm/)
  - [Deutsche Telekom has 200 managed clusters](https://searchitoperations.techtarget.com/news/252510456/Deutsche-Telekom-preps-Kubernetes-5G-core-with-GitOps)
  - major enterprise companies that cannot be disclosed have shared that they had 50-100 clusters using Flux, and on top of that: biggest vendors integrate Flux in GitOps offerings today
- Governance: Flux maintenance and community processes have been a well-oiled machinery for quite some time - we are very happy where we are today.
- Development health: [Time from Opened to Merged](https://flux.devstats.cncf.io/d/16/opened-to-merged?orgId=1&from=now-5y&to=now) dropped significantly in the past 5 years.
- Transition from v1 to v2 well handled so far. Number of requests for v1 has been ¼ in 2021 and even a smaller percentage in the last 3-6 months.
- Solid roadmap (more security work, multi-tenancy pieces, flagger on top of flux2 controllers), defined by all stakeholders.
Integration with other projects (shared blog posts, presentations at events, collaboration over integrations).

## What has changed since you joined Incubation?

### High-level stats

![Statistics](https://i.imgur.com/vgX0bEi.png)

### Development & Stability

- Flux implements the vast majority of our [roadmap](https://fluxcd.io/roadmap/) and is used as the foundation for GitOps offerings by multiple vendors now.
We made significant headway in our transition from Flux v1 to v2 - it is still in active maintenance mode. Support requests have been 75% about v2 these days of issues/PRs this year. Percentage for the last 3-6 months is even higher.
- The list of features and improvements since Incubation is long: big ones are Server-Side Apply and an ever-growing list of supported Git providers, source types, notification protocols and general tooling integration (git submodules, extended SOPS support, ssh/gpg key formats, etc).
- All TODO items from Incubation have been addressed.
- One exception in process: move [Flagger](https://flagger.app/) docs under <https://fluxcd.io>.
- On feature stability: the public API for Flux is its Kubernetes custom resource APIs; these are all now beta versions, and any further changes will be backward-compatible.
- Flux APIs were marked stable in July 2021: <https://fluxcd.io/blog/2021/07/july-2021-update/#from-now-on-flux-apis-will-be-stable>
- As consumers of Kubernetes and Helm APIs, we are collaborating more closely and have a bigger pool of knowledgeable users and willing testers. One good example is: <https://github.com/fluxcd/helm-controller/issues/345>

### Community

- We made a [big push to make Docs and the website](https://github.com/fluxcd/website/graphs/contributors) more user friendly and fluxcd.io the central hub for information. We had a [recent assessment](https://github.com/cncf/techdocs/pull/89) of our docs and website by the CNCF Tech Docs group and it was quite positive.
- We have more and more active community participation in meetings, or KubeCon booths, folks who help each other out on Slack, etc.
- Steady stream of monthly updates to keep the ecosystem involved. We were able to extend readership and the number of reported items in the posts have grown over time and span more parts of the project.

## How can your project survive its founders and founding companies?

The Flux project members have taken several steps to ensure the project's safe survival past its founders and founding companies.

We moved all discussion, decision-making processes, and team organisation to be public. We started attracting more sub-project maintainers outside of Weaveworks. It has never been easier to participate and be involved.

Several vendors ([Amazon, Microsoft, VMware, D2iQ](https://youtube.com/playlist?list=PLwjBY07V76p5a_KXuAZdGvhWznlj34A5T)) have chosen to base their GitOps offerings on top of Flux. There is also a Flux offering through Red Hat's OperatorHub, which Red Hat is exploring for deeper offerings in the future. We are seeing increasing contributions and engagement in community meetings from them in our project because of this. Companies such as Xenit and DoneOps provide Flux support, Hybridk8s and [Shipa](https://shipa.io/miscellaneous/fluxcd-dashboard/) have built a UI for Flux and Flagger and are building a business around it. D2iQ have been providing increased contributions. Finally, global enterprise companies such as [Deutsche Telekom](https://www.weave.works/blog/case-study-how-deutsche-telekom-automates-edge-infrastructure-using-gitops), [Under Armour](https://youtu.be/VYwfZQc7hxc?t=88), and [Fidelity](https://youtu.be/rT9VM_JSf5A) depend on Flux and have members with technical expertise to contribute if needed.

Flux maintainers are actively involved in the CNCF GitOps WG, and OpenGitOps project, to ensure we hear from the broader ecosystem what is needed and to build ties with other projects.

Looking at the [top 15 contributing companies of Flux (all time)](https://flux.devstats.cncf.io/d/5/companies-table?orgId=1) and [those of the last year](https://flux.devstats.cncf.io/d/5/companies-table?orgId=1&var-period_name=Last%20year&var-metric=contributions), it becomes clear that the majority of contributions outside of Weaveworks have happened in the past year. Particularly by those who built GitOps offerings on top of Flux and/or participate as maintainers: D2iQ, Microsoft, Red Hat, Xenit.

## Who are the best advocates for your project? What do they say?

We have several users from a wide range of companies who have confirmed that they can be interviewed by TOC members. TOC members can request access to [this document](https://docs.google.com/document/d/1TyZmaIkY973roUvnQeMVbVrslTLIfQJYLgCmLiFXgn8/edit) for contact info for users agreeing to be interviewed about their Flux usage. We have gathered their names, contact info, and whether they can be interviewed on or off the record.

### Quotes

#### Adopters / End users

- **[MediaMarktSaturn](https://www.weave.works/blog/mms-finds-confidence-to-implement-progressive-delivery-with-flagger)**: "Adding Flagger to our stack was a really simple task since we were already using Helm charts." - Bernd Stübinger, MediaMarktSaturn
- **[US Department of Defense](https://www.cncf.io/blog/2021/09/30/how-to-get-robust-gitops-the-u-s-department-of-defense-uses-flux-and-helm/)**:
  - For Nicolas M. Chaillan, Chief Software Officer of the U.S. Air Force, "Configuration as code was a shining star. The CNCF Flux project was a clear choice for us at Platform One because it provided fully-baked support for Helm, which provides a robust ecosystem of tools for our end users."
  - **Tom Runyon** (tom@leapfrog.ai | tom@defenseunicorns.com) - Public talk: Orchestrating Software Factories for the Department of Defense
    "We chose flux for several reasons, one of the most impactful is the depends-on field - as we can actually declare how these things roll out without having to worry about sync waves."  
    "The original version of Big Bang was Argo based exclusively and we switched up Flux as the underlying for the native Helm support because there were some charts that did not get rendered correctly in argo-speak and had trouble deploying. Specifically, GitLab had some very strange artifacts with the post install hooks as part of the chart and when argo tried to duplicate all the effort of Helm into their own language, there ended up being some conversion problems which is why I appreciate the native Helm support that Flux provides because we don't have to worry if Flux is doing it right."
- **Anonymous, major financial institution**: "Our platform team supports internal users on different Kubernetes versions and gives cluster admins from different business units permissions across EKS, AKS, and on-prem (with Rancher). We use the open source Kraan project to create this seamless experience in a highly regulated industry by using Kubernetes and CNCF Flux. Because of Flux's helm-controller and source-controller (together with the kraan-controller), our Cluster admins don't have to deal with complex deployment/upgrade logic."
- **Rajarajan Pudupatti SJ, Kraan project maintainer**: "Kraan is an open source project that helps you deploy and manage 'layers' on top of kubernetes. By applying layers on top of K8s clusters, you can build focused platforms on top of K8s e.g ML platforms, Data platform etc. Kraan is powered by the CNCF Flux project and builds on top of Flux's source-controller and helm-controller. Thanks to these controllers, Cluster admins don't have to deal with complex deployment/upgrade logic. For Kraan's creators, relying on Flux's advanced controllers meant that they didn't have to build something from scratch."
- **[Michael Irwin (Virginia Tech)](https://cloud-native.slack.com/team/U010KV76SLT)** "As a platform architect, Flux has been a critical component of our platform. We love the fact that git becomes our source of truth and that we don't need to figure out how to distribute credentials to push changes into the cluster. By simply providing a git repository, we can tell teams "here's a place to define your manifest and we'll take it from there." Teams can choose their own change management practices, use CI pipelines, or make their updates completely manually. Flux has made our platform more flexible to our customers, yet more secure simply by using it."
- **[Josh Callisv (oso.sh)](https://cloud-native.slack.com/archives/CLAJ40HV3/p1638372189247100)**: "I moved over from being a Senior Software Engineer to become a DevOps engineer because of flux and Kubernetes - wouldn't have found the wonderful world of DevOps without it and the job opportunity that required enthusiasm  and knowledge of all thing flux and Kubernetes!"
- **[Laszlo Fogas (gimlet.io)](https://app.slack.com/client/T08PSQ7BQ/D02PZC1400Z/user_profile/U64KTJ493)**: "Flux is a foundational piece of Gimlet, an application operations platform on top of Kubernetes. Flux made working with Kubernetes resources straightforward and transparent. Anyone can inspect the yamls that go on the cluster with tools that they know. It also enabled us to declaratively deliver not just applications, but infrastructure stacks as well. It reliably gets the job done, and provides wonderful opportunities for integration. Thank you Flux team!"

#### Integrators

- **[AWS](https://www.youtube.com/watch?v=FCtImtLNnjk)**: "Thanks to Flux we are able to provide a very smooth and simple user experience to let our users manage their EKS Anywhere clusters in a native way, through a git repository." - Chandler Hoisington & Joey Wang
- **[D2IQ](https://twitter.com/superguenter/status/1448484330955706368)**:
  - "We chose Flux because it's Kubernetes native. Ties into RBAC which makes it more secure. Modular and extendable. Can run per namespace." - Tobi Knaup, CEO & Co-Founder of D2IQ ([Announcement blog post](https://d2iq.com/blog/goodbye-dispatch-hello-fluxcd))
  - "What I love about @fluxcd  is that while adding a thin abstraction layer e.g. on top of Helm it doesn't try to hide the underlying complexity. You can easily reproduce its behavior manually using the Helm or kustomize CLIs." - Max Jonas Werner (D2IQ)
- **Microsoft**
  - **Jason Hansen**: "Azure provides managed GitOps with Flux so that you can easily enable GitOps in your Azure Arc-enabled Kubernetes of Azure Kubernetes Service clusters."
  - **[Brendan Burns (Microsoft and Kubernetes co-creator)](https://youtu.be/c2k16drHqvA)**: "In getting everything into a git repository and ensuring that the only way that configurations flow into your cluster is through something like Flux and the workflows involved in Flux. That's the way that we can actually shift left not just for our code that we're writing, but shift left for the things that we're deploying as well."
- **[Red Hat](https://www.youtube.com/watch?v=8ybCtNVP5To)**: GitOps with Flux + OpenShift - Andrew Block
  - "One of the benefits of GitOps is when you introduce Flux. Flux is that GitOps engine that enables you to achieve GitOps in a Kubernetes environment and coming together with OpenShift you really have this perfect match between the two solutions."
  - "One of the benefits of Flux is really the ability to be that GitOps engine that allows you to deploy."
- **[VMware](https://www.youtube.com/watch?v=qm1ZKsTcxa4)**: GitOps with VMware Tanzu Application Platform - Ben Hale & Leigh Capili
  - "We've gained a lot of utility from using Flux"
  - "The change to "the GitOps Toolkit way of thinking" has been really really great for us."
  - "I am continually impressed by the Flux Maintainers and how they've made something that's this composable that we went and did this novel thing."
  - "We're taking a look at the image update automation area and the foundation exists for us to start thinking about using Flux at this end of our supply chains inside of the Tanzu Application Platform as well."
- **Unofficial quotes in conversation**:
  - "We had some discussions early on about the various options but pretty quickly settled on Flux. [it] came down to "cluster-nativeness" and modularity."
  - "Flux's everything's-a-CRD model really matches with our own thinking..."
  - "The fact that we've been able to adopt Source Controller without taking on the rest of Flux…has been really valuable to us."
  - "In the end the ethos of Flux aligns really well with our own."
- **Mohsin Shafique, eKomi Holding GmbH**
  - [Flexing the Flux - Multi-Cluster, Multi-Cloud - somewhat opinionated take](https://youtu.be/PTsDNC33SFo)
  - "When we started re-platforming our services, what we wanted was a simple tool that doesn't get in the way. So among many of the solutions we evaluated, we thought Flux was the simplest without introducing any new concepts or constructs beyond standard tooling. We're pretty happy with the simplicity of Flux and we'll be continuing to use it in our existing and upcoming Kubernetes deployments."
- **Dan Wessels, solo.io: [GitOps and Cloud Native API Gateways](https://youtu.be/yzE-9qgyJGg)**

## Graduation FAQ

### Assessment of dependencies

- The Flux project uses Go libraries which are well-maintained and well-spread throughout the ecosystem: `{sigs.,}k8s.io`, `helm.sh`, plus relevant SDKs from Google, Azure and AWS to interact with relevant cloud services.
- For multi-arch container image builds, we have a dependency on [`tonistiigi/xx`](https://github.com/tonistiigi/xx). This project is stable (`>1.0.0`), and used by recognized projects and organizations ([[1]](https://github.com/search?q=tonistiigi%2Fxx+org%3Adocker&type=Code), [[2]](https://github.com/go-gost/gost/blob/28644802964ed66ea862d95f4026ecc0dee8a727/Dockerfile#L5), [[3]](https://github.com/search?q=org%3Amoby+tonistiigi%2Fxx&type=code)).
- We realise that some of these APIs have not been marked GA just yet, e.g. Kubernetes core `controller-runtime` or `kustomize`, but a lot of Cloud Native projects rely heavily on them and we have good ties into these communities and figure out issues together when they arise.
- One key design principle for Flux is to integrate as tightly with tooling APIs as possible and never to shell out to CLI tools. For our story around git you might want to refer to e.g. <https://fluxcd.io/blog/2022/03/flux-puts-the-git-into-gitops/>. This is why we use several Go libraries to interact with `age`, `sops`, `git`, `helm`, `ssl`, `gpg` and more. This is to provide the best user experience possible and to support the multitude of different ways of working (Git submodules, GPG signed commits, various SSH key formats, etc.)
- We recommend reading <https://github.com/fluxcd/pkg/tree/main/runtime#goals> as our philosophy towards Kubernetes dependencies and standards.

Understood problem areas:

- The future of SOPS was uncertain for some time, but [Mozilla has confirmed they are committed to continue to maintain it](https://github.com/mozilla/sops/discussions/927#discussioncomment-2183834).
  We are meeting with the SOPS team to see if we can help.
- go-git upstream admits in <https://github.com/go-git/go-git#comparison-with-git> that it is
  > challenging for *go-git* to implement all the features
  
  Till date, this has besides the lack of support for the v2 Git transport protocol (which was solved with libgit2, and is edge-case for many users other than Azure DevOps folks), not posed any issues.
- We had to go great lengths to be able to integrate the most recent `libgit2` with all its dependencies into Flux, but we got the hang of it now, cf <https://github.com/fluxcd/golang-with-libgit2>
  Several other companies and projects depend on [libgit2](https://libgit2.org/), including [AWS](https://github.com/aws/eks-distro-build-tooling), [GitHub](https://github.com/collections/projects-that-power-github), [Microsoft](https://github.com/search?q=org:microsoft+libgit2&type=code) and Git Kraken.

*What If?*

If we should find that e.g. Git libraries should not serve the needs of our users any more, or they go unmaintained, we could fall back to shelling out to `git` as we did in [Flux Legacy](https://github.com/fluxcd/flux) (v1).

Our design is actually also taking this possibility into account, and we can easily introduce new implementations (be it shelling out, or some other Git package that suddenly shows up), by writing a new [GitProvider implementation](https://fluxcd.io/docs/components/source/gitrepositories/#git-implementation).

Theoretically, but much less desirable: We could even fork out some of the projects and maintain them ourselves. Flux requires a small subset of features within Git and therefore, there are sufficient alternatives available.

At present, the need to shell out to `git` does not seem very likely.

### Review of Incubation TODO

During the Incubation process of Flux, the [following TODO items](https://docs.google.com/document/d/1Z6yPN9-bWeVGpMrBxXJ3NBTBYZowJ3R93wKHEVmBJ1A/edit#heading=h.wt1gt34gj6vx) were brought up:

Follow-up after move of /f/flagger (from /weaveworks):

- [x] Move #flagger on Weaveworks Slack to CNCF home
- [x] Replace Flagger logo (formerly Weaveworks logo with ‘flagger’ underneath).
      Discussion here: <https://github.com/fluxcd/flux2/discussions/653>
- [x] Move <https://flagger.app> under CNCF / LinuxFoundation  
      Status: Done: <https://github.com/fluxcd/community/issues/201>

Come up with comprehensive plan re: docs and urls

- [x] move Flux v2 docs from toolkit.f.i to f.i/docs
- [x] Move Flux v1 docs under f.i/docs
- [x] Move Flagger docs under f.i/docs  
      Status: Website has moved. Last step: replace old page with redirects. Tracked here <https://github.com/fluxcd/flagger/issues/1182>

### Flux contributor ladder

We are continuing to put a lot of thought and work into providing a great contributor experience and making it easy for folks to get involved.

Some of the steps we undertook so far:

- Clearly defines roles in our community: <https://github.com/fluxcd/community/blob/main/community-roles.md>
- Run regular meetings ([since Nov 2018](https://www.youtube.com/playlist?list=PLwjBY07V76p5mWNgdINjIiuUiItIeLhIN))
  We are very pleased to have folks from various organisations join these meetings to discuss features they need and would like to see integrated, one particular example is the integration of OCI support, which is of big interest to many.
- Run the weekly [Flux Bug Scrub event](https://www.youtube.com/playlist?list=PLwjBY07V76p6J6z30cBRqS_N0Ka6NhEsY) for 9 months now, which is essentially one-to-one mentoring by a Flux maintainer.
- Broadcast individual ways to get involved: <https://fluxcd.io/blog/2022/04/contributing-to-flux/>.
- [Monthly updates](https://fluxcd.io/tags/monthly-update/) published on all channels, which celebrate the work by everyone and also introduce new maintainers and their background.

### Flux Security Process

Instructions for how to report Flux vulnerabilities confidentially and responsibly can be found prominently on our website <https://fluxcd.io/security/> which is the [SECURITY.md document](https://github.com/fluxcd/.github/blob/main/SECURITY.md) that all of the `fluxcd` repositories share.

That's also the document where the Security Team (and their GPG keys are shared). Past audits and advisories are listed there as well.

Internally, the security team keeps track of the reported vulnerabilities and their status, making sure each reported item is investigated, and if confirmed, fixed. The process also accounts for Coordinated Vulnerability Disclosures as well as Public Security Disclosures.

### Flagger as part of Flux organisation

Flagger was designed to work with Flux and it's a natural fit for the GitOps toolkit. This is already noted briefly in this PR under ["Why is the Project ready to Graduate?"](#why-is-the-project-ready-to-graduate) In short, the plan is to add GitOps Toolkit compatibility to Flagger. Here are some more links. See the top section under [Flagger's Roadmap](https://github.com/fluxcd/flagger/blob/main/README.md#roadmap) in its README, and was added by [this PR](https://github.com/fluxcd/flagger/pull/714) in Oct 2020.

Flagger is intended to remain a sub-project of Flux. There are 133 contributors, 60 with more than one commit, but @stefanprodan has made 1300 commits. The closest 2nd is Takeshi Yoneda, who has made 50 commits, and is the only co-maintainer. Part of the reason for this move is that after Flagger is refactored on top of flux2 controllers, a wider group of existing maintainers can help support it in addition to Stefan.

## Follow-up after meeting with TOC

- [x] <https://github.com/fluxcd/community/pull/199>: fix small inconsistencies in our governance documentation.
- [x] <https://github.com/fluxcd/community/issues/201>: moved <https://flagger.app> domain to LinuxFoundation IT.
- [x] Moved Flagger docs to <https://fluxcd.io/flagger> (<https://github.com/fluxcd/website/pull/903>).  
      Status: After some time we will turn `flagger.app` [into redirects](https://github.com/fluxcd/flagger/issues/1182).
- [x] <https://github.com/fluxcd/community/pull/200>: provide examples for decision making questions, e.g. Flagger moving under Flux organisation, applying for membership, RFC process.
- [ ] <https://github.com/fluxcd/pkg/issues/263>: Update related packages in `fluxcd/pkg`.
- [x] Extend security team: <https://github.com/fluxcd/.github/pull/15>.
- [x] Transform Oversight Committee into transparent Governance part that matches day-to-day practices and has members voted in.  
      Status: [Governance PR](https://github.com/fluxcd/community/pull/203) merged, [follow-up items](https://github.com/fluxcd/community/issues/207) resolved.
- [x] Start process with TAG Security to review post-GA multi-tenancy plans: <https://github.com/cncf/tag-security/issues/896>
