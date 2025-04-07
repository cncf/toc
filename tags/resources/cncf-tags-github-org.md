# `cncf-tags` GitHub org

The [`cncf-tags` GitHub org](https://github.com/cncf-tags) serves as a home for TAG-sponsored projects and tools.
This org is intended to provide a vendor-neutral place for TAGs to collaborate on projects
endorsed by and actively worked on by members of a TAG.
The org will _not_ contain projects that can potentially apply to become CNCF projects.

Each repo has a "topic" of the form `tag-<tag-name>`. This ensures that repos are easily searchable per TAG
through repo topics. For example, all repos under tag-runtime can be found at https://github.com/topics/tag-runtime.

## Creating a new repo

To request for a new repo, an issue must be created against the `cncf/toc` repo
using the `Create a repository in the cncf-tags GitHub Org` issue template.

The issue must be approved by:

- TOC liaisons for the respective TAG.
    - If a TAG has 3 TOC liaisons, the issue must be approved by at least 2 liaisons.
- At least one of the TAG leads or any process spelled out in a TAG's charter.

A publicly linkable written decision should be available for all approvals.

Once approved, the TAG leads can open a [ServiceDesk](https://github.com/cncf/servicedesk) ticket
to get the repo created.

## Access for each repo

Repository access will be maintained through `.github/settings.yaml`.
This process will be later migrated over to [CLOWarden](https://github.com/cncf/people#configyaml-configures-cncf-org-repository-access)
once [multi-org support]((https://github.com/cncf/clowarden/issues/43)) is implemented in CLOWarden.

The following should have access:
- TOC liaisons for the respective TAG and TAG leads should have admin access.
- TAG leads may grant additional privileges (write/admin access) to other TAG members.
A publicly linkable written decision should be available for granting such access.

### Archiving a repo

TAG repos may be archived if they are deemed _inactive_. Inactive repos are those that meet any of the following criteria:

- There are no longer any active maintainers for the project and no replacements can be found.
- All PRs or Issues have gone un-addressed for longer than six months.
- There have been no new commits or other changes in more than a year.
- The contents have been folded into another actively maintained project.

**Approval for archiving a repo**

To request archival of a repo, an issue must be created against the `cncf/toc` repo.

The issue must be approved by:
- At least one TOC liaison for the respective TAG
- At least one TAG lead

Once approved, the TAG leads can open a [ServiceDesk](https://github.com/cncf/servicedesk) ticket
to get the repo archived.

### Mandatory files

Each repo, at minium, should have the following files:

- `LICENSE`
- `code-of-conduct.md`
- `README.md`
- `CONTRIBUTING.md`
- `SECURITY.md`

https://github.com/kubernetes/kubernetes-template-project can be referred to for inspiration.
