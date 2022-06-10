# 2022 Kuberhealthy Annual Review
<center><img src="https://github.com/kuberhealthy/kuberhealthy/blob/master/images/kuberhealthy.png?raw=true"></center><br />

**Project Page**: https://github.com/orgs/kuberhealthy/repositories

**Devstats Dashboard**: https://kuberhealthy.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m

**Docker Hub Images**: https://hub.docker.com/orgs/kuberhealthy/repositories

**Stars**: 1,459

**Forks**: 208

## Active Maintainers 
Currently, I am the primary maintainer (@integrii from @EpicGames (Previously @Comcast)) of Kuberheralthy.  This has been fine due to the relatively low volume of issues and code reviews, but additional contributors should be sought after.  There are several contributors to the project that may be interested in the duties of a maintainer, such as @Hungrylion2019, or @bavarianbidi).

## Contributors

Kuberhealthy has [68 contributors](https://github.com/kuberhealthy/kuberhealthy/graphs/contributors) and growing!

## Adopters
Kuberhealthy has had trouble growing an official adopters list because companies generally require approval to have their names posted in the project.  Additionally, the [ADOPTERS.md](https://github.com/kuberhealthy/kuberhealthy/blob/master/docs/KUBERHEALTHY_ADOPTERS.md) file has been located in the `docs` directory, but is getting [moved to the root for more visibility](https://github.com/kuberhealthy/kuberhealthy/pull/1049). The current official adopters include Adobe, Jenkins X, and Mercedes Benz.  Based on the various GitHub profiles I have seen raising issues and contributing to the project, along with the huge number of image pulls we're getting, I would wager that the actual list of adopters is much, much higher.

On the [Docker Hub Images](https://hub.docker.com/orgs/kuberhealthy/repositories) page, you can see that [one check has over 5 million pulls](https://hub.docker.com/orgs/kuberhealthy/repositories) and [several others have over a million pulls](https://hub.docker.com/repository/docker/kuberhealthy/deployment-check).  The main Kuberhealthy controller image [has more than 500k pulls](https://hub.docker.com/repository/docker/kuberhealthy/kuberhealthy).

Keep in mind that the number of pulls does not reflect the number of actual runs because images are cached on the nodes they run on.

## Performance since last review
While no specific goals for growth were set, Kuberhealthy has continued to ramp up in stars and pull request frequency while maintaining a low issue count.  Currently there are [375 issues](https://github.com/kuberhealthy/kuberhealthy/issues) filed and [688 pull requests](https://github.com/kuberhealthy/kuberhealthy/pulls).  The number of pull requests is inflated some due to GitHub actions rebuilding the helm chart on each change, but the majority of those PRs are from huamns.  

Our top referring site is Google.com, which indicates that people are seeking the project out by need or directly by name the majority of the time, rather than being linked from some specific high volume site or social media post.  GitHub shows 308 unique visitors on the Kuberhealthy home page each month currently.

<img width="448" alt="image" src="https://user-images.githubusercontent.com/98695/172985357-8e7e3a1b-e1cf-482a-b25e-b812ac10f0aa.png">

## Goals
I think it would make sense to increase the number of maintainers on the project, so that I am not the one doing most of the work.  We already have people with the maintainer permissions, but they do not regularly review issues or PRs unless mentioned.  Getting them involved would be healthy for the project, as would finding contributors willing to convert to maintainers from the wider community.  That said, I think we all know how hard it can be to find people willing to proactively donate time to open source software that they are not paid to maintain.

The [milestones](https://github.com/kuberhealthy/kuberhealthy/milestones) for the project are slightly out of date due to a 2.7.x release being cut without issues being juggled to a new version, but I believe the roadmap portrayed within is still a good one.  This road map points towards a web interface, enhanced documentation, and an official reference to our `khcheck`, `khstate`, and `khjob` custom resources.  We obviously need to continue supporting our community, listening to feature requests, and processing code reviews to keep things bug-free and featureful.


## How the CNCF can help
It would be nice if we could have a paid docker hub account.  The existing free account only supports 3 members, and those spots are filled.  More members would mean more people able to cut releases, and that will be important to growth.  Plus, its good to support Docker Hub.

It would also be nice if the CNCF was able to design us an improved logo (ours was basically clip-art) or fund some SWAG for our contributors and official adopters!

If we find ourselves hitting a limit with GitHub actions, some funding to keep that moving would be nice too - but we aren't there yet.


## Incubation
I don't believe Kuberhealthy is ready for incubation yet.  I would like to see the project garner more official adopters, gather some more maintainers, and add a web interface before applying to be promoted.


## Thank You
Having Kuberhealthy adopted into the CNCF has been wonderful for the project!  Thank you for the support!
