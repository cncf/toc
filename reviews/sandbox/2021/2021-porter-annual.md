# Porter 2021 Annual Review

[Porter](https://porter.sh) creates installers for cloud-native applications and their infrastructure.
Package your application artifact, client tools, configuration and deployment logic together as a versioned bundle that you can distribute, and then install with a single command.

Porter joined the CNCF Sandbox in September 2020.

## Metrics

| Metric | 2020 | Today |
| ------ | ---- | ----- |
| Stars  | 291  | 658   | 
| Forks  | 71   | 116   |
| Downloads | N/A | 33,367* |

\* Downloads since April 2021.

DevStats: https://porter.devstats.cncf.io

## Maintainers

https://github.com/getporter/porter/blob/main/OWNERS.md

* Carolyn Van Slyck (Microsoft)
* Jeremy Rickard (Apple)
* Vaughn Dice (independent)
* Reddy Prasad (independent)
* Jennifer Davis (Google)

We have 3 active maintainers on the core repository right now: Carolyn, Vaughn and Jennifer. Our project has lots of smaller repositories for our plugins and mixins. There we have 2 additional active repo level maintainers: Simon Davies (Microsoft) and Mohamed Chorfa (independent).

We recognize that we need more maintainers and have been encouraging existing contributors to build up experience with the codebase so that they could handle reviewing pull requests as well.

## Adoption

https://github.com/getporter/porter/blob/main/ADOPTERS.md

We have had a significant uptick in adoption in the last 6 months. Microsoft continues to be our largest adopter, and we have added one other official adopter to our list: Thinktecture. Additional companies have privately shared that they are building prototypes on top of Porter, but they aren't yet ready to officially add their names to our adopters file.

We believe that adoption will increase when we release [v1.0.0][milestone] and the Porter Operator. Our v0.38 release is stable but isn't recommended for production.

One promising use case for Porter is building a marketplace for installing third-party software on a cloud provider's platform. Some cloud providers have disclosed that they are prototyping using bundles and the Porter Operator in their marketplace implementations.

## Goals

Over the past year, we have focused on building an adopters list and encouraging known adopters to list their name publicly in our adopters.md file. Once we release v1.0.0, we expect more companies to add their name to the public list.

Our current goals are focused around three initiatives:
* A v1.0.0 release of Porter. We are over half-way to completing our v1 [milestone].
* Completing a Kubernetes operator for Porter. One that is well-suited for automatically managing production installations, especially using a GitOps workflow with Flux. Most, if not all, of our private adopters are waiting for this functionality before going public.
* Promote more contributors to maintainers.

# How can the CNCF help?

We plan on having the CNCF tech docs group evaluate our documentation soon. We have lots of content but need assistance organizing it so that it is easier to locate. We also want to improve the type of documentation available, and reduce gaps, so that our users are more successful.

# Incubation

We are not ready to apply for incubation at this time. We would need more maintainers, and public adopters first.

[milestone]: https://github.com/getporter/porter/milestone/16