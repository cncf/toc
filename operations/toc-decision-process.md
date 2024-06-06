# TOC Decision Making Process

This document outlines a decision-making structure used by the TOC that takes into account feedback from members of the community and strives to find consensus, while avoiding deadlocks.

All discussions take place based on the Operating Model defined in the [charter](https://github.com/cncf/foundation/blob/main/charter.md#6-technical-oversight-committee-toc).

To ensure that the process is as efficient as possible, the amount of notice provided and the amount of consensus sought is driven by the estimation of how complicated an issue or PR is.

## Major Changes

To propose any change to the existing governance, principles or processes that impact the broader CNCF community (including projects, TAGs, etc), the following process shall be followed:

- Open a GitHub issue in the `cncf/toc` repo describing the change.
- Shepherd a public discussion about the topic in a TOC meeting.
- Post the results of the discussion to the GitHub issue and
send a note to the CNCF TOC mailing list.
- Create a pull request to propose the change.
- The PR remains open for a minimum of 2 calendar weeks for the community to review.
- Address any review comments on the PR.
- After resolving all review comments, the PR must remain open
for a _reasonable time_ to allow for any additional comments.
    - At least one calendar week constitutes _reasonable time_ in the above sense.
    - This period resets at every new major change made to the PR.
      Minor changes like typo fixes do not reset the period.
- For the PR to be approved:
    - It needs approval from at least 2 TOC members.
    - TOC decisions are driven by consensus. If the TOC has not reached consensus,
      the PR approval moves to a [vote](#voting).
- The PR can be merged when **all** of the following are true:
    - The two week public comment period has ended
    - All reviews have been addressed
    - If there are any major review comments, the _reasonable time_ period has elapsed
    - There are at least 2 approvals from TOC members _or_ there is a [passing TOC vote](https://github.com/cncf/foundation/blob/main/charter.md#6-technical-oversight-committee-toc)
- Notify the TOC mailing list of the merge.

## Minor Changes

The following process shall be followed for changes that are narrowly scoped. Some examples of these changes are typo fixes or improving the clarity of existing TOC processes as documented where there is no impact to projects, the community, the outcome of processes, or any such changes that modify the intent of their execution.
It is expected these changes are made in good faith to enable velocity for projects by removing confusion and enable increased transparency.

- Open a GitHub issue or directly create a PR to the `cncf/toc` repo to propose the change.
- For the PR to be approved, it needs approval from at least one TOC member _and_ the TOC Chair is to be informed to confirm previous consensus in
applying the minor change.
- If the PR is approved as per above, the PR can be merged.

## Voting

Voting is required for processes like:
- Adding a project to the Sandbox, Incubation or Graduation
- Moving a project between levels
- Adding or removing a TAG Chair or Technical Lead
- Achieving consensus among the TOC

Votes are managed through at least one of the following ways:
- [gitvote](https://github.com/cncf/gitvote)
- over email on the TOC mailing list
- on call in a public TOC meeting

The specific voting details are mentioned in each process' guidelines and the requirements laid out in the [charter](https://github.com/cncf/foundation/blob/main/charter.md#6-technical-oversight-committee-toc).
Processes that pass votes do not require additional approvals and can be directly merged by the CNCF Staff.
