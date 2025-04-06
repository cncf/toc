# Cloud Native Guide Posts

Guide Posts are intended to outline guiding points that have assisted past cloud native projects as they grew and matured in the cloud native ecosystem. They are NOT requirements for moving between levels, please refer to [the criteria for moving levels](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md) for requirements. 

Projects may leverage these as steps towards increasing contributors, adopters, maturity, etc. and in doing so may find their path towards Incubation and Graduation is more structured, clear, and achievable _however_ this is subject to each project's implementation, design, and community. Even after projects are Graduated, they may refer back to Guide Posts as a mechanism to continue sustaining their project in a manner anticipated by their community and adopters particularly as they experience change and turnover. These are not organized by maturity level, but are presented in an ordering that is commonly experienced as projects mature and move levels.

The community is welcome to submit PRs to improve and add additional Guide Posts based on their experiences in maturing their projects.

## Guide posts

### Governance and maintainers

* A clearly discoverable governance doc exists that covers the basics of decision making, including who approves PRs and how committers/maintainers are chosen
  * Appropriate governance for cloud native projects is critical to their overall success. When projects are early in their maturity, we see simplified governance. As they grow and mature, we expect to see the governance expanded, more readily exercised, and generally shift to accommodate the needs of the community and the project's direction

* Evolves and develops a robust set of governance documentation that defines how the project is managed, and its roadmap is determined, which considers maintainer turnover and organization/affiliation/company diversity. 
  * [TAG Contributor Strategy](https://github.com/cncf/tag-contributor-strategy) has many great [governance templates](https://github.com/cncf/project-template) and plenty of additional templates on [contribute.cncf.io](https://contribute.cncf.io/maintainers/templates/)

* May demonstrate the application, practice, and adjustments of its governance documentation as a result of regular project operations
  * Ensuring documented, transparent, and open decision making is a hallmark of maturing governance for any project. Contributors and community members should have access to sufficient records and decisions to understand what has occurred and why. Blogs, articles, GitHub & Slack discussions, etc. all serve as methods to demonstrate application and continued alignment with a project's governance (so long as they are not private)

* Establish and exercise contributor growth ladder or other community construct to "build the bench" of maintainers or project leaders
  * Such as a Contribution Ladder which considers succession planning and getting maintainers from multiple organizations
  * In addition to defining how contributors may move up the ladder to roles with more responsibility, current project leaders and maintainers should identify potential contributors interested in moving up the ladder and periodically check in to support them
  * Beyond just planning for one, actively exercising a growth ladder can establish a pipeline of contributors that can step forward and distribute work of existing maintainers while passing along institutional knowledge to ensure successful succession when the existing maintainers are ready for their next adventure

* Experiencing successful turnover in leadership that showcases adherence to governance
  * Governance is only beneficial when it is used and exercised, its creation is often down in at a neutral time with no pressure of outcome. It is most needed when problems arise, pressure exists on a specific outcome, and temperaments may create tension and conflict in the project. Establishing an agreed upon process prior to these problems gives the project a "cool head" and clear path forward

* Regularly update governance documentation to accurately reflect changes to processes, staffing, roles, and activities as they occur
  * May re-review as new maintainers come in or as other's move to Emeritus
  * May seek out support from TAG Contributor Strategy to verify governance is clear and addresses potential concerns

* Refreshing long-term planning with a reasonable cadence
  * Long term planning can be a roadmap, vision/feature goals, GitHub project board, etc
  * Long term planning is an open and transparent effort by the project to describe their intended future work and is subject to change
  * Long term planning is leveraged as part of the evaluation for moving levels, it is often leveraged by potential adopters to schedule or plan for adoption if a particular feature is wanted, and also helps community members identify where they can contribute or align their interests

### Contributors and community

* Community meetings or active discussions are held
  * Such as meeting via Zoom, capturing discussion notes, and occurring more than twice a quarter (or community.cncf.io!) 
  * Such as use of GitHub discussions, issues, or a group messaging service (slack, wechat, etc.)
  * Ensuring these are documented on the project README allows new contributors to easily find information

* Community contributions are well guided
  * Such as maintainers and other active contributors providing timely feedback on PRs from community members
  * Constructive and welcoming feedback on PRs and issues allows for a positive first experience with a project and can increase the probability of a contributor returning

* Increasing regular contributors to the project
  * Such as identifying potential contributors from different organizations beyond the current adopters
  * Setting goals to identify contributors that would be wonderful maintainers, and partnering with them to onboard or prepare for that role

* A second organization begins contributing
  * Those contributions may begin path to maintainership
  * Identifying the types of requirements and activities to become and remain as a maintainer is a good start to begin creating a pipeline or ladder for contributors to consider maintainership and a critical step to begin succession planning

* Engaging a [Technical Advisory Group](/tags/README.md) for feedback on the project relevant to the domain the project primarily functions within (Networking, Storage, App Delivery, etc.)
  * TAGs are community members, other maintainers, and domain experts with an interest and experience in a particular domain area. They are an excellent resource for projects to become regularly acquainted with as they often produce educational material for adopters of projects to inform them of best practices and availability of projects and integrations to meet certain use cases. They can serve as a sounding board for ideas, features, and use cases
  * TAGs also support projects in navigating the moving levels process and often provide the TOC with domain specific context about a given project that supports the technical due diligence performed by the TOC

* Engaging with a few Technical Advisory Groups to consider best practices across multiple technical and governance facets
  * Socializing a project with TAGs outside the project's domain area can increase exposure of the project to more potentially interested individuals. This can also identify areas of potential integration which can improve the adopter experience or even provide ideas that allow a project to be competitive in the ecosystem

* Contribution and community documentation clearly defines expectations when contributing (to include non-code contributions)
  * Expectations when contributing could be as simple as providing an overview of what happens after they submit a PR (review, discussion, etc.). They can also be more detailed and outline a variety of mechanisms by which someone can contribute to the project, like triaging issues, updating docs, facilitating community meetings, etc

* Defined a process for managing the conduct of its community, either establishing its own group to do this, leveraging the existing leadership, or deferring to the CNCF
  * Every CNCF project adopts the Code of Conduct, however each project may choose to establish a group dedicated to promoting adherence and embodiment of the code of conduct. Projects may find that a dedicated group in this area can also provide additional support in promoting an inclusive and welcoming community
  * Projects may choose to defer to the CNCF's Code of Conduct Committee for resolving reported conduct issues

### Engineering Principles

* Begins to solidify a versioning schema and release cadence
  * Such as providing a roadmap, long term planner, GitHub project board, or other document which defines project direction or initiatives that may align with releases
  * Adopters may experience releases that begin to establish a pattern of regularity
  * Many projects will initially begin with more frequent releases, sometimes once a month or once a quarter, it just depends on the project. Expect to adjust the frequency of releases as the project continues to mature, becoming more stable and robust with less active development

* Documentation includes getting started guides, operating/administration instruction, security call-outs, and other core elements necessary to ease adoption
  * As project maintainers, we're often very close to the project and know all the ins and outs. Ensuring our project documentation is robust, easy to understand, and makes no assumptions about the expected expertise or skill level of a potential adopter or contributor allows the project to more readily be understood, tested, and adopted. Consider having someone not familiar with the project review project documentation and provide feedback

* Providing information on performance and scalability for deployment options/configurations, autoscaling, or providing benchmarks in these areas for comparison or evaluation
  * Many adopters appreciate performance and scalability information as it allows them to understand more readily what they are getting into. It also gives them a reasonable understanding if their deployment is configured and performing correctly or if there are other issues. It may also support their transition from an existing solution to a more performant and competitive open source project

* Captures use cases to showcase how it can be best used to solve common problems.
  * These may leverage observations and concerns from adopters
  * Understanding how technology can solve a business need is sometimes difficult. Documenting adopter use cases and user stories, creating proof of concepts and demos to demonstrate how a project can support a given use case helps adopters navigate this challenge

* Ensuring all sub-projects are clearly defined in their maturity level, whether they ship with the core project or can be used independently but is subject to the same governance processes
  * Providing statuses of sub-projects and their definitions (alpha, beta, stable, unstable, experimental, etc) can provide adopters with a very clear understanding of what to expect when they adopt or configure a sub-project
  * Distinguishing between which sub-projects are packaged up or included by a project's artifacts allows adopters to understand what features and capabilities are included in the project and enables them to add sub-projects as-needed to suite their needs. It also allows the project to assert what has and hasn't been done or enforced on a sub-project, to remove assumptions on its security posture, trust boundary, and performance capabilities

* Maintains or provides a dependency graph so adopters fully understand the complexity and risk of use
  * Dependency graphs can provide maintainers with insights into new features or bugs introduced by their dependencies which may impact a projects stability, performance, and security posture. They also provide a more comprehensive understanding of what packages a project depends on and how the project is constructed. This also allows adopters to know what they are getting _without installing the software_ to ensure it aligns with their organization's policies for licensing and dependencies

### Security

* The basics of security are underway
  * Such as completion of a [self-assessment](https://github.com/cncf/tag-security/blob/main/assessments/guide/self-assessment.md)

* A robust and mature security posture is being developed
  * Any of the following would assist in achieving this
    * Threat model
    * [Joint review](https://github.com/cncf/tag-security/tree/main/assessments#components-of-the-security-review-package)
    * [OpenSSF Best Practice Badge](https://bestpractices.coreinfrastructure.org/) [Passing level](https://bestpractices.coreinfrastructure.org/en/criteria/0) or [Silver](https://bestpractices.coreinfrastructure.org/en/criteria/1)
      * Note: Completion and maintenance of **an** OpenSSF Best Practice Badge is [a requirement for graduation](../.github/ISSUE_TEMPLATE/template-graduation-application.md#required-3)
      * The OpenSSF Best practices badging program provides a series of best practices for open source projects in increasing levels of maturity. While passing is the current criteria for Graduation projects, the current trends in industry are focusing on more secure projects and demonstrating a project's commitment to security through more advanced security checks and criteria, such as silver or gold best practices badging, can distinguish your project from the next.

### Ecosystem

* Angel adopters assist in stability and production use or deployment
  * An angel adopter is a very early adopter of a cloud native project. They are invaluable to a projects' growth and first experience in interacting with adopters.
  * May be from different industry verticals
  * May have differing use cases
  * While not all projects may experience several angel adopters (characterized by early and active engagement), they may have other early adopters who are less active in the community but do contribute in some way or another (often identifying areas of improvement, bugs, needs, use cases, etc)
  * Angel adopters often are characterized as helping the project reach momentum through:
    * higher level of engagement with the project 
    * ongoing injection of feedback and support via issues and contributions
    * expertise or experience on traditionally difficult challenges for early stage technical development, performance, etc

* Has 3 or more adopters using the project in production
  * Multiple adopters leveraging a project in a production environment shows increased use of a project. It demonstrates the project meets maturity, performance, security, and scalability expectations for a set of adopters who may provide feedback and contributions to continue advancing and refining the project to a point in which more adopters can rely on it

* Maintaining an Adopters file which may record public production users and use cases to benefit other potential adopters and contributors
  * Many adopters of projects look to their industry peers to keep abreast of potential challenges, solutions, and technology in use. Capturing not only by whom, but how your project is being used could attract additional adopters whose needs and challenges aligns with your current ones

* Planning activities to increase project adoption
  * Projects may choose to create focused efforts on driving awareness and increasing opportunities for adoption. This could be in the form of presentations at conferences outside of KubeCon+CloudNativeCon (KCCN), seeking out user stories and publishing them on a blog, and presenting at TAGs

### Other

* Periodically evaluate project health and practices
  * Such as leveraging [CLOMonitor](https://clomonitor.io/)
  * It is often difficult to make time to reflect on your project's health, using existing dashboards and tooling can support you in determining where to focus your governance and community efforts

* TBD - add more!