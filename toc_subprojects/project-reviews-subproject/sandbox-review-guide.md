# Sandbox Reviewer Guide

> Note!  This is a living document, and is written in a somewhat informal style.
<!-- you can put less-visible notes in a comment, like if I wanted to add "Don't like it?  Rewrite it." -->

This document aims to record best practices (and the reasoning behind them) for CNCF sandbox application reviewers.  This advice is intended to help streamline project reviews, **but it is not binding**.  Use your best judgement when applying this advice, and feel free to ignore or adjust it for specific circumstances.

## Overview

Lots of projects apply to the CNCF Sandbox.  They hope to gain a number of benefits from CNCF membership, including: visibility, neutral governance and IP ownership, access to CNCF expertise, and probably even more.  The TOC approves Sandbox applications using the [sandbox process][sandbox]; the goal of sandbox reviews _is to provide the information the TOC needs to make [a status decision][sandbox-status].

To that end, the primary output of the review process is a fair assessment of the project's [general technical foundations][gtr].  Because the bar for sandbox is intentionally low, only the "Day 0" questions need to be filled out.

Additionally, sandbox reviewers should check the general sandbox application answers, and [flag any concerns for the project and TOC to address](#other-application-concerns).

### Recommended GTR Process:

If the project has already completed the [General Technical Questions][gtr], this process can be simplified to reviewing the project's answers, rather than generating a new technical review document.

If the project has not completed the technical questionnaire, copy the [General Technical Questions document][gtr] into a _new_ https://notes.cncf.io/ document for sharing with co-reviewers or the project.  Note that only the "Day 0" contents are required for sandbox applications.

Iterate on the questionnaire answers until you and any co-reviewer are satisfied that the questionnaire is complete.  If you need information or clarification from the project, prefer public written in the application Issue (best!) or recorded (e.g. TAG presentation) formats over private channels such as DMs or 1:1 video meetings.  Maintaining public channels makes it easier for interested participants and onlookers to understand and engage with the process, and fosters transparency.

When the questionnaire is complete, _COPY_ all of the https://notes.cncf.io contents into a **comment** on the sandbox application issue.

> Why _copy_ the contents back and forth between notes.cncf.io, the issue, and an eventual file in the `cncf/toc` repo?  We're attempting to balance the following:
>
> * Only commit content to the `cncf/toc` repo for projects which have been accepted into the CNCF, to avoid confusion.
> * Make it easy to find (e.g. search) for previous applications, either for returning applications, or for related work.
>    * The CNCF notes documents have access control which makes implementing search very difficult.  They are useful for collaborative editing, but make it difficult to reference and index the contents later.
> * Preserve the record of decisions in a way that is clear and transparent to all participants, including those viewing the results in the future.

### Other Application Concerns

When reviewing a donation from a corporate entity which also intends to sell some portion of the service, ensure that the issue makes clear which code (preferably, as entire repositories) **will** and **will not** be donated.  See [this comment](https://github.com/cncf/sandbox/issues/392#issuecomment-3173237281) for an example of such a clarification.


[sandbox]: https://github.com/cncf/sandbox
[sandbox-status]: https://github.com/cncf/sandbox/blob/main/README.md#the-tocs-sandbox-review-meeting
[gtr]: ./general-technical-questions.md