# Cloud Native Technical Papers process

This guide serves to provide technical groups in CNCF a consistent process and mechanism to create, collaborate, and publish technical content in the form of whitepapers that increase the technical accuracy and potency of content provided. It promotes visibility of these activites, fair and opportune community feedback, and rigor in technical review.

This Paper process has been adopted from [CNCF's TAG Security Paper Process](https://github.com/cncf/tag-security/blob/main/publications/paper-process.md) and modified to cover TAGs and WGs under the oversight of the TOC. It is intended to complement existing TAG governance.

## Purpose of Technical Papers in cloud native

Technical papers, commonly referred to as whitepapers, provide the target audience with technical understanding on a given topic. Within the CNCF, technical papers serve one of two primary purposes:
* **Guiding the ecosystem in understanding and reasoning around a technology or technical domain area and any best practices associated with it.** Papers that serve as ecosystem guides inform cloud native projects and community members on a particular topic or series of related topics. The audience for these papers is usually other technologists and engineers that are actively involved in projects so they may incorporate recommendations, concepts, and guidance into their projects. These are may be future leaning, where the ecosystem currently lacks the technology or where the concepts are still emerging and projects may benefit from insights and guidance curated by domain experts.
* **To inform adopters about a particular technology, technical domain, or concept and how they can it addresses a desired use case.** Papers that provide adopters of cloud native projects and the broader industry with a more indepth understanding of a particular technology component such as APIs, technical concepts like sustainability footprints, and sub-domain topics like software supply chain security. These papers may establish decision or technology frameworks that allow adopters to reason about cloud native and open source projects to achieve a particular use case and will convey best practices for their use and integration.

Papers are intended to serve as static sources of information that can frequently be referred back to. When scoped and developed correctly, papers may only need refreshed or updated once a year or as new understandings of the topic emerge. If a particular topic area or technology is rapidly changing, and lacks industry consensus, a paper is not recommended and interested individuals should consider a presentation or blog post on the topic that reflects the pace of change and opportunities for clarity and alignment.

## Initiating a Paper

In order to ensure consistency in development, review, and publication, TAGs and WGs that wish to produce technical papers will follow this process, adapting their groups' particular governance to the necessary areas defined.

### Defining a need

For any individual or group interested in developing a technical paper, there must be a defined need or observable opportunity to close a knowledge gap. When proposing a technical paper, an initial scope must be established to facillitate discussion and interest within the TAG via an Issue. It is expected that the Issue title and content will be adjusted and modified as the TAG and group refine the scope.

### Presenting the paper idea to the TAG

The paper scope and topic should be raised in at least one TAG meeting to solicit volunteers (minimally 4). TAG leadership reserves the right to hold off on moving forward with a paper due to existing TAG commitments, lack of volunteers, or lack of consensus on need. 

### TAG leader sponsorship and cross-TAG awareness

A TAG leader (co-chair or tech lead) must step forward to sponsor the paper. Other TAG leaders may serve as the lead for paper development but the role of sponsor must be independent from the lead for the paper to reduce any bias in the process. The paper's sponsor is responsible for ensuring the development of the paper adheres to this process, is still progressing towards publication, and provides the initial review of the draft before public comment period may open.

The paper sponsor is also expected to coordinate awareness of the paper development and timelines to other TAGs to ensure they are appropriate included and have the opportunity to contribute and participate.

### Refining the scope and audience and participants

Interest in participating is expressed on the Issue by comments of those individuals. Soliciting interest and participation must allow for sufficient time (no less than two weeks) for fair opportunity to gather interest. Interested parties should then meet to define timelines, refine the scope, declare the audience, and determine paper lead.

## Managing paper development

Once the group has met to defien timelines, scope, audience, and determined the lead, the Issue must be updated to provide the following information (initial estimates provided but may be adjusted by the group or TAG leadership as appropriate):

| Milestone | Estimated Time |
| --- | --- |
| Audience, Goals, & refining scope | 1 week |
| Outline | 2 weeks |
| Tasking Assignment | 1 week |
| Content Rough-in | 2-3 weeks |
| Collaborative Review | 2 weeks |
| Executive Summary and content wrap up | 2 weeks |
| Narrative Voice | 1-2 weeks |
| Final Group Review | 1 week |
| Community Review | 2 weeks |
| Public comment adjudication | 2 weeks (simultaneous with review) |
| CNCF publishing engagement | ~2-3 weeks |
| Addition to the repo | 2 weeks |
| Blog post and publishing coordination | 2-3 weeks |

### Milestones Details

#### Audience, Goals, and Refining Scope

Define the following:

- Who is the intended audience?
- What positions do they hold?
- What problem are we addressing?
- What assumptions are we making about the audience or content?
- Document these in the "Introduction > Assumptions" section.

#### Outline

The outline needs crafted and placed in a shared document (google docs, hackmd, or other collaborative documentation tool). It will serve as the basis for paper structur and tasking.

#### Tasking Assignment

Contributors should assign themselves to various outline areas. This allows focus and identifies gaps needing more contributors. Use tools like Google Docs' "assign to me" feature for task assignment. Alternatively, participants may tag themselves in line by typing their name.

#### Content Rough-in

Assignees provide rough content through phrases, paragraphs, or bullets. If the assignee has drafted complete content (full sentences, ideas, well expressed) they should indicate the content as "draft". This informs other collaborators that less effort may be needed to further develop the content.

#### Collaborative Review

Contributors refine raw content into drafts, comment on sections, expand ideas, and provide citations. Use the corresponding Slack channel for larger thematic discussions.

#### Executive Summary and Content Wrap-up

If the paper exceeds five pages, write an executive summary. Finalize content and discussions internal to the group.

#### Narrative Voice

Assign no more than three individuals for this task to ensure a unified voice. Ensure consistent language, phrasing, acronym use, footnotes, and citations. The final result of this unification should ensure the paper reads consistently, like it was written by a single entity.

#### Final Group Review

Collectively review the polished paper before opening it for community review. The group should assign individuals to serve as comment adjudicators. These individuals will be responsible for validating commentary against the scope, assumptions, and audience.

#### Community Review

Lock the paper from editing (except for Adjudicators). Make it public for comments and suggestions. Announce a two week public comment period via the TAG's mailing list as well as the TOC's mailing list.

#### Public Comment Adjudication

Adjudicators resolve community comments and suggestions. Larger discussions should be decisively resolved and documented.

#### Leadership Review and Approval

Two additional reviews are required before the paper may be published

##### TAG Sponsor review and approval

The TAG Sponsor must review and provide their approval on the finalized draft. Their review and approval signifies the paper as ready for the Liaision review and approval.

#### CNCF Publishing Engagement

TAG Leadership sponsor works with CNCF for final edits, PDF conversion, and graphics inclusion. Review drafts before the final version is added to the repo.

#### Addition to the Repo

The paper lead creates a README.md with:

- Title
- About: Brief summary
- Updates: Maintenance intentions
- Markdown: Maintained in markdown
- Contributing updates: Guidance for contributors
- Versioning and publishing: Criteria for updates
- Original design decisions
- Links to files in the repo

### Blog Publishing and Coordination

Coordinate with TAG leadership and CNCF for a blog post to increase visibility. Consider presenting at community events.

## Authorship, Attribution, and Acknowledgements

Papers created by a CNCF TAG are authored by the group, with contributors and reviewers acknowledged. Major contributors are highlighted in the "Acknowledgements" section. Each document should contain "Contributors," "Reviewers," and "Acknowledgements" sections as appropriate.
