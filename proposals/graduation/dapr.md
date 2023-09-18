[Dapr](https://dapr.io) Dapr is a vendor-neutral, cloud agnostic and API-driven framework that enables developers to solve hard distributed systems challenges like event-driven messaging, state management at scale, distributed locks,  long running workflows and many others. Underpinning the Dapr APIs are 3 pillars of security, observability and reliability, each one provided out-of-the-box and made configurable.

Dapr integrates with over 120 cloud-based and open-source systems, and uses CNCF projects like OpenTelemetry, gRPC, and SPIFFE at its core.

The project has been [accepted as a CNCF Incubating project](https://www.cncf.io/blog/2021/11/03/dapr-distributed-application-runtime-joins-cncf-incubator/) in November 2021. The [proposal](https://github.com/cncf/toc/blob/main/proposals/incubation/dapr.md) and [due diligence](https://docs.google.com/document/d/1GknuJk8Z26KhvFJiblbMSopWbX27nihpxkEG0_TpkCw/) are provided here for reference.

## Graduation criteria requirements

Dapr meets the graduation stage requirements in the following ways:

### Have committers from at least two organizations.

Dapr makes use of the term "[maintainer](https://github.com/dapr/community/blob/master/community-membership.md#maintainer)" as the noun for a member allowed to accept and merge code contributions. To date, the project has [14 active maintainers](https://github.com/dapr/community/blob/master/MAINTAINERS.md), representing 4 organizations.

### Have achieved and maintained a Open Source Security Foundation Best Practices Badge.

Dapr’s OpenSSF badge can be seen [here](https://www.bestpractices.dev/en/projects/5044).

### Have completed an independent and third party security audit with results published of similar scope and quality as the [Envoy example](https://github.com/envoyproxy/envoy#security-audit), and all critical vulnerabilities need to be addressed before graduation.

A [recent audit was published September 2023](https://www.cncf.io/blog/2023/09/06/dapr-completes-2023-security-audit-increasing-enterprise-confidence/). No critical or high severity issues were found. Auditors concluded that "our overall assessment of Dapr is highly positive. Dapr follows security best practices in both design and implementation. Dapr performed well in this audit demonstrating a strong security posture." In addition, the project [published a fuzzing audit report](https://blog.dapr.io/posts/2023/06/30/dapr-completes-fuzzing-audit/) from June 2023 which has found 3 issues. The auditors concluded they "were impressed with the low count of issues found given the large number of fuzzers created. This is a testament to the well-written and well-maintained codebase of the Dapr project.

### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

The Dapr Steering and Technical Commitee governance document can be found [here](https://github.com/dapr/community/blob/master/steering-and-technical-committee-charter.md).

The Steering and Technical Commitee is in charge of setting the project's technical direction, governance and administation, as well as marketing and relationship with CNCF.

Dapr has [clearly documented community membership roles](https://github.com/dapr/community/blob/master/community-membership.md), giving contributors transparency into the different rights and responsibilities that come with each role and how one can change status from one to the other. The project has [contribution guides](https://docs.dapr.io/contributing/) for the benefit of newcomers, encouraging a diverse set of contributions for all Dapr repositories.

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

Dapr's [Steering and Technical Committee members](https://github.com/dapr/community/blob/master/steering-and-technical-committee-charter.md) are authorized to interact with the CNCF on behalf of the project. The criteria and process for selecting and offboarding maintainers can be found [here](https://github.com/dapr/community/blob/master/community-membership.md#responsibilities-and-privileges-2).

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

A partial list of public adopters can be found [here](https://github.com/dapr/community/blob/master/ADOPTERS.md). CNCF case studies are listed [here](https://www.cncf.io/case-studies?_sft_lf-project=dapr) and several others from the Dapr site can be found [here](https://dapr.io/testimonials/).