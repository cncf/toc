# ko - Sandbox

**Name of project:**

ko

**Description:**

ko is a simple, fast container image builder for Go applications. It's ideal for use cases where the image contains a single Go application without many dependencies on the OS base image (e.g. no cgo or OS package dependencies). ko builds images by effectively executing go build on a local machine, and as such doesn't require docker. ko is also a good fit for lightweight CI/CD use cases given its multi-platform support and automatic SBOM generation capabilities.

**Statement on alignment with CNCF mission:**

Even before Google’s release as a stand-alone project in March 2019, ko has proven to be indispensable to the Knative community (which now [has been accepted](../incubation/knative.md) as CNCF incubation project), as a simple, fast, and secure container image builder for Go applications. Today the broader open source community and many enterprise development teams depend on ko to improve their developer productivity. The ko project is also increasingly used as a solution for a number of build use-cases, and is being integrated into a variety of third party CI/CD tools.

Open source comes with a responsibility to contribute, sustain, and improve the projects that make our ecosystem better. To support the next phase of community-driven innovation, enable net-new adoption patterns, and to further raise the bar in the container tool industry, ko needs a vendor-neutral foundation.

We believe that CNCF is the right foundation where ko can meet the evolving needs of the community. Being part of CNCF will bring the ko community closer to other Cloud native projects in the ecosystem, including many of the projects which already depend on ko. Being in the same foundation will enable ko to better collaborate and has a potential to increase the velocity of feedback and innovation over time.


**Sponsors from TOC:**

TBD

**Preferred maturity level:**

Sandbox

**License:**

Apache2

**Source control:**

Github (https://github.com/ko-build/ko)

**Initial Committers:**

Maintainers:

- Jason Hall (Chainguard)
- Jon Johnson (Google)

**Infrastructure requests (CI / CNCF Cluster):**

- GitHub Actions for CI/CD
- Container image repository
- Communication Channels:
  - IM/Slack: `#ko-project` channel on the [Kubernetes Slack](https://slack.k8s.io)
  - Issue tracker: https://github.com/ko-build/ko/issues
  - Mailing List: None
  - Website: https://ko.build

**Existing sponsorship:**

None

**Community size:**

The community around ko has been growing staidly since the project's inception. Currently:  

- 5.1k GitHub stars (https://github.com/ko-build/ko/stargazers)
- 64 unique contributors (https://github.com/ko-build/ko/graphs/contributors)
- 280 forks (https://github.com/ko-build/ko/network/members)
- 114 repositories depending on ko (https://github.com/ko-build/ko/network/dependents)

