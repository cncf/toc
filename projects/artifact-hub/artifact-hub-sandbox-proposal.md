# Artifact Hub Sandbox Proposal

## Project name and description

The [Artifact Hub](https://artifacthub.io/) is a hub for cloud native projects to have their artifacts listed and easily found by end-users. The goal of the Artifact Hub is easy discovery of cloud native artifacts (e.g., Helm charts, Falco rules, etc) and to provide information about the artifacts so that end-users can make decisions about using them.

short name: artifacthub

## Alignment with CNCF charter mission

The [CNCF charter](https://github.com/cncf/foundation/blob/84dc259f50b89d656907f2247d9415d8c10c3efd/charter.md) notes that:

> The Foundation’s mission is to make cloud native computing ubiquitous.

The charter goes on to say that:

> We democratize state-of-the-art patterns to make these innovations accessible for everyone.

The Artifact Hub works to make artifacts accessible. Many artifacts are distributed. They are distributed between different projects, they are distributed between different vendors or other organizations and, they are distributed globally. For end-users who want to use these artifacts with their projects they have to know what to look for and search in various places. Discovering artifacts is difficult and time consuming.

The Artifact Hub provides search and discovery for distributed artifacts. The hub is not a store for the artifacts. Instead, it aggregates metadata about the artifacts and makes them searchable in a single place. It makes access to artifacts simpler for those who are trying to find them.

## Comparison With Similar Projects

There are four classes of similar projects in existence and we will cover each along with how the Artifact Hub is different from them.

### General Search Engines

General search engines (e.g., Google and Bing) can be used to discover distributed artifacts. General search engines provide general content. Artifacts are intermixed with general content. When using search engines one needs to use more descriptive language to find artifacts. Doing a search on "PostgreSQL" won't produce a Helm chart or Falco rules. You need to know to search specifically for those. General search engines also lack metadata to make it easier to figure out what to use.

The Artifact Hub provides a custom and targeted search for users. A search for "PostgreSQL" will return artifacts and display information to help end-users make decisions on what to pick. Faceted search is also available to help reduce options in a manner the end-users can choose.

### Existing Project Hubs

The [Helm Hub](https://hub.helm.sh), [OperatorHub](https://operatorhub.io/), and [Cloud Native Security Hub](https://securityhub.dev/) are examples of existing hubs. Each of these hubs is for a subset of the cloud native artifacts available. The Helm Hub lists Helm charts, the OperatorHub lists operators compatible with the Operator Frameworks Operator Lifecycle Manager (OLM), and the Cloud Native Security Hub lists Falco rules and is experimenting with Open Policy Agent policies.

These hubs work a little differently from each other. The Helm Hub lists distributed Helm charts managed by various organizations (e.g., VMware, AWS, etc). The organizations register a repository and the Helm Hub regularly fetches updates to know about changes, new, or removed charts. It is a distributed discovery mechanism. The OperatorHub lists operators that are listed in a git repository. New operators and versions need to be merged into the repository to be listed. The Cloud Native Security Hub holds the rules for Falco and is maintained by the Falco project via Sysdig.

The Artifact Hub most closely follows the Helm Hub as it pulls in from external locations to make finding them more accessible. This enables organizations to host their own artifacts while making them discoverable. Where it's different from the Helm Hub is that varying artifacts from different projects can be discovered and the architecture is being built in a manner that enables new artifact types to be added. The Artifact Hub also makes self management easier than existing hubs.

### OCI Registries

OCI Registries (e.g., Docker Hub) provide a means to store artifacts, usually container images but more things are being made available, and the experience around them enables discovery of the artifacts.

The Artifact Hub does not host artifacts like container images or Helm charts. The artifacts are hosted by 3rd parties, like the CNCF member organizations who operate those services. The Artifact Hub is not attempting to compete with those. Instead, the Artifact Hub is attempting to make discovering artifacts that are being stored in various services and locations more accessible.

### Package Searches For Other Platforms

Other platforms and languages have their app stores and package discovery services. The Artifact Hub is most similar to [Packagist](https://packagist.org/) and different from services like NPM. NPM stores JavaScript packages along with making them discoverable. Packagist makes packages discoverable while 3rd parties handle the hosting.

## Sponsor from TOC

TBD

## Preferred maturity level

Sandbox

## License

[Apache 2.0 license](https://github.com/cncf/hub/blob/master/LICENSE).

## Source control

https://github.com/artifacthub

## External Dependencies

The following list contains direct dependencies.

| Dependency                                                  | License analysis |
|-------------------------------------------------------------|------------------|
| PostgreSQL                                                  | PostgreSQL       |
| github.com/disintegration/imaging                           | MIT              |
| github.com/domodwyer/mailyak                                | MIT              |
| github.com/go-chi/chi                                       | MIT              |
| github.com/gorilla/securecookie                             | BSD-3-Clause     |
| github.com/h2non/go-is-svg                                  | MIT              |
| github.com/jackc/pgconn                                     | MIT              |
| github.com/jackc/pgx                                        | MIT              |
| github.com/rs/zerolog                                       | MIT              |
| github.com/spf13/viper                                      | MIT              |
| github.com/stretchr/testify                                 | MIT              |
| golang.org/x/crypto                                         | BSD-3-Clause     |
| golang.org/x/time                                           | BSD-3-Clause     |
| gopkg.in/yaml.v2                                            | Apache 2 & MIT   |
| helm.sh/helm                                                | Apache 2         |
| github.com/kulshekhar/ts-jest                               | MIT              |
| github.com/testing-library/jest-dom                         | MIT              |
| github.com/testing-library/react-testing-library            | MIT              |
| github.com/testing-library/react-hooks-testing-library      | MIT              |
| github.com/testing-library/user-event                       | MIT              |
| github.com/DefinitelyTyped/DefinitelyTyped                  | MIT              |
| github.com/moment/moment                                    | MIT              |
| github.com/react-icons/react-icons                          | MIT              |
| github.com/twbs/bootstrap                                   | MIT              |
| github.com/JedWatson/classnames                             | MIT              |
| github.com/prettier/eslint-config-prettier                  | MIT              |
| github.com/prettier/eslint-plugin-prettier                  | MIT              |
| github.com/lydell/eslint-plugin-simple-import-sort          | MIT              |
| github.com/typicode/husky                                   | MIT              |
| github.com/okonet/lint-staged                               | MIT              |
| github.com/lodash/lodash                                    | MIT              |
| github.com/sass/node-sass                                   | MIT              |
| github.com/prettier/prettier                                | MIT              |
| github.com/facebook/react                                   | MIT              |
| github.com/react-icons/react-icons                          | MIT              |
| github.com/rexxars/react-markdown                           | MIT              |
| github.com/ReactTraining/react-router                       | MIT              |
| github.com/facebook/create-react-app                        | MIT              |
| github.com/conorhastings/react-syntax-highlighter           | MIT              |
| github.com/npm/node-semver                                  | ISC              |
| github.com/Microsoft/TypeScript                             | Apache 2         |

## Initial committers

The initial committers are Sergio C. Arteaga and Cynthia S. Garcia who are the primary developers. Matt Farina (of Samsung SDS) also has commit privileges.

## Infrastructure requests

To host the instance of the hub software powering https://artifacthub.io/.

## Communication channels

Communication is currently in CNCF Slack and primarily in the #artifact-hub room.

## Issue tracker

In the [GitHub project issues](https://github.com/artifacthub/hub/issues).

## Website

The [Artifact Hub](https://artifacthub.io/) is a website. The [hub](https://github.com/artifacthub/hub) codebase can be run by other organizations who want to have their own hub (e.g., a corporate version). Documentation for the project lives alongside the code on GitHub.

## Release methodology and mechanics

GitHub Actions are used for CI and run on every pull request in addition to changes landing on branches. The current branches are master, staging, and production. The staging and production branches represent environments (i.e. https://artifacthub.io/ and https://staging.artifacthub.io/).

The current software development process targets web development with staging and production environments.

Hosting is currently in AWS. The code is built into container images, pushed to a registry, and deployed using Helm. These processes are handled using GitHub Actions for CI.


## Social media accounts

None

## Community size and existing sponsorship

The Artifact Hub is in the early stages of development. It currently has 84 stars on GitHub. The project is nearing the point of a wider launch.

The existing sponsorship is the CNCF itself with support from the Helm project. The Helm project needs to do a major update to the Helm Hub and we are looking for the Artifact Hub to be the next generation of hubs so that Helm Hub can be deprecated. Note, deprecation of the Helm Hub will be determined based on the success of the Artifact Hub.

According to the [Sandbox documentation](https://github.com/cncf/toc/blob/master/process/sandbox.md), in the section on what early stage means, the following applies to the Artifact Hub:

> Projects commissioned or sanctioned by the CNCF, including initial code for CNCF WG collaborations, and "experimental" projects.

### Project logo

The current logo is in SVG and can be found on the [Artifact Hub website](https://artifacthub.io/).
