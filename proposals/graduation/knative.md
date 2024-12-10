[Knative](https://knative.dev) is a set of serverless APIs that provides higher level abstractions to enable developers to be more productive by simplifying and extending the Kubernetes interface for common stateless application patterns.

Knative consists of three APIs – Serving, Eventing and Functions – and includes reference implementations and conformance tests for both the Serving and Eventing components (Functions consists of a set of client-side tools for building containers and interfacing with Serving). These components provide pluggable abstractions over existing HTTP and message routing implementations. A robust ecosystem of plugins and extensions exists in the [knative-extensions GitHub org](https://github.com/knative-extensions) and elsewhere. Developers can use either or both of serving and eventing to accelerate their development and build powerful applications quickly and efficiently.

Knative has been an active project since 2018, with involvement from a diverse group of individual developers and companies, including Google, IBM, Red Hat, VMWare, TriggerMesh, SAP and many more. Knative was accepted to the CNCF on [March 2, 2022](https://github.com/cncf/toc/pull/762).

## Graduation criteria requirements

We believe Knative meets the graduation stage requirements in the following way:

### Have committers from at least two organizations

Knative does not have a “committer” role. Instead, the project has various roles that are defined in project governance’s [roles section](https://github.com/knative/community/blob/main/ROLES.md).

****

As of writing (Jan 16, 2024):

- There are 8 [working groups](https://github.com/knative/community/blob/main/working-groups/WORKING-GROUPS.md), with 13 unique leads with 6 different employer affiliations

- 5 [Steering Committee](https://github.com/knative/community/blob/main/STEERING-COMMITTEE.md#committee-members) members with 4 employer affiliations

- 5 [Technical Oversight Committee](https://github.com/knative/community/blob/main/TECH-OVERSIGHT-COMMITTEE.md#committee-members) with 4 employer affiliations

- 25 reviewers with more than 5 employer affiliations who did at least 50 reviews between Jan 1, 2023 - Dec 31, 2023 [devstats query](https://knative.devstats.cncf.io/d/29/pr-reviews-by-contributor?orgId=1\&var-period=y\&var-repo_name=All\&from=1672520400000\&to=1704056399000\&var-reviewers=%22creydr%22\&var-reviewers=%22ReToCode%22\&var-reviewers=%22pierDipi%22\&var-reviewers=%22dprotaso%22\&var-reviewers=%22Cali0707%22\&var-reviewers=%22matzew%22\&var-reviewers=%22skonto%22\&var-reviewers=%22dsimansk%22\&var-reviewers=%22Leo6Leo%22\&var-reviewers=%22davidhadas%22\&var-reviewers=%22yuzisun%22\&var-reviewers=%22psschwei%22\&var-reviewers=%22nak3%22\&var-reviewers=%22houshengbo%22\&var-reviewers=%22matejvasek%22\&var-reviewers=%22aliok%22\&var-reviewers=%22gabo1208%22\&var-reviewers=%22kvmware%22\&var-reviewers=%22lkingland%22\&var-reviewers=%22prushh%22\&var-reviewers=%22cardil%22\&var-reviewers=%22evankanderson%22\&var-reviewers=%22upodroid%22\&var-reviewers=%22jrangelramos%22\&var-reviewers=%22rhuss%22\&var-reviewers=%22pawarpranav83%22\&var-reviewers=%22mgencur%22\&var-reviewers=%22KauzClay%22\&var-reviewers=%22dan-j%22\&var-reviewers=%22alexagriffith%22\&var-reviewers=%22geekygirldawn%22\&var-reviewers=%22navinag1989%22\&var-reviewers=%22xiangpingjiang%22\&var-reviewers=%22rajibmitra%22\&var-reviewers=%22SD-13%22\&var-reviewers=%22AdamKorcz%22\&var-reviewers=%22ricardozanini%22\&var-reviewers=%22jsanin-vmw%22\&var-reviewers=%22prakrit55%22\&var-reviewers=%22gauron99%22\&var-reviewers=%22vinfinity7%22\&var-reviewers=%22astelmashenko%22\&var-reviewers=%22lance%22\&var-reviewers=%22kauana%22\&var-reviewers=%22arsenetar%22\&var-reviewers=%22andrew-su%22\&var-reviewers=%22md-saif-husain%22\&var-reviewers=%22arkadiusjonczek%22\&var-reviewers=%22nainaz%22\&var-reviewers=%22yanggangtony%22\&var-reviewers=%22sharmaansh21%22\&var-reviewers=%22dolfolife%22\&var-reviewers=%22rahulii%22\&var-reviewers=%22csantanapr%22\&var-reviewers=%22pastequo%22\&var-reviewers=%22pmbanugo%22\&var-reviewers=%22salaboy%22\&var-reviewers=%22professorabhay%22\&var-reviewers=%22octonawish-akcodes%22\&var-reviewers=%22maschmid%22\&var-reviewers=%22MeenuyD%22\&var-reviewers=%22sadath-12%22\&var-reviewers=%22krsna-m%22\&var-reviewers=%22zhouhaibing089%22\&var-reviewers=%22airycanon%22\&var-reviewers=%22puerco%22\&var-reviewers=%22cdalar%22\&var-reviewers=%22mfreeman451%22\&var-reviewers=%22bradbeam%22\&var-reviewers=%22Zazzscoot%22\&var-reviewers=%22ChunyiLyu%22\&var-reviewers=%22zainabhusain227%22\&var-reviewers=%22sandipanpanda%22\&var-reviewers=%22ryutoyasugi%22\&var-reviewers=%22mayank-a-agarwal-db%22\&var-reviewers=%22none%22)

- [82 contributors](https://github.com/knative/community/blob/main/elections/2023-SC/voters.yaml) with more than 5 employer affiliations who are eligible to vote in the [latest election](https://github.com/knative/community/tree/main/elections/2023-SC)

### Have achieved and maintained a Open Source Security Foundation Best Practices Badge.

Knative’s OpenSSF badge is [here](https://bestpractices.coreinfrastructure.org/en/projects/5913).

### Have completed an independent and third party security audit with results published of similar scope and quality as the [Envoy example](https://github.com/envoyproxy/envoy#security-audit), and all critical vulnerabilities need to be addressed before graduation.

Knative completed an independent and third-party audit sponsored by the OSTIF in November of 2023. The third-party auditor was Ada Logics, who engaged from June-October; issues were found and fixed in both the Knative code and in third-party dependencies (mostly denial of service attacks related to large HTTP request reads). The CNCF [blog post](https://www.cncf.io/blog/2023/12/11/knative-completes-third-party-security-audit/) states that there have been 16 issues found including 1 CVE. All issues have been fixed.

In addition to the security audit, an independent third party security fuzzing audit has been completed by Ada Logics. The high level information in CNCF [blog post](https://www.cncf.io/blog/2023/07/14/knative-fuzzing-audit-results/) states that the audit had spanned several months in late 2022 and early 2023.The [report](https://github.com/knative/docs/blob/main/reports/ADA-knative-fuzzing-audit-22-23.pdf) mentions that there was a single issue found. This issue later has been [marked as resolved](https://bugs.chromium.org/p/oss-fuzz/issues/detail?id=56798).

### Explicitly define a project governance and committer process. The committer process should cover the full committer lifecycle including onboarding and offboarding or emeritus criteria. This preferably is laid out in a GOVERNANCE.md file and references an OWNERS.md file showing the current and emeritus committers.

Knative governance does not have a “committer” role. Instead, contributors who have write permissions are called “approvers”.

The lifecycle of the approver status is outlined in the [contributor ladder](https://github.com/knative/community/blob/main/ROLES.md#role-summary) and the responsibilities are described in the project [governance document](https://github.com/knative/community/blob/main/ROLES.md#approver). The governance also defines the criteria for offboarding contributors, including approvers and other roles.

### Explicitly define the criteria, process and offboarding or emeritus conditions for project maintainers; or those who may interact with the CNCF on behalf of the project. The list of maintainers should preferably be stored in a MAINTAINERS.md file and audited at a minimum of an annual cadence.

Knative’s [Steering](https://github.com/knative/community/blob/main/STEERING-COMMITTEE.md#committee-members) and [Technical Oversight Committee](https://github.com/knative/community/blob/main/TECH-OVERSIGHT-COMMITTEE.md#committee-members) members are authorized to interact with the CNCF on behalf of the project and listed publicly in the committee documents.

These committee members hold their positions for 2 years and they are elected with an [election process](https://github.com/knative/community/tree/main/elections). The governance rules define the criteria for offboarding people from all roles, including maintainers.

### Have a public list of project adopters for at least the primary repo (e.g., ADOPTERS.md or logos on the project website).

Knative has a list of [adopters](https://github.com/knative/community/blob/main/ADOPTERS.MD) published publicly. The current list has 29 organizations and is growing regularly.

Furthermore, the community collected [testimonials](https://knative.dev/docs/about/testimonials/) from 4 end-users and conducted 4 [case studies](https://knative.dev/docs/about/case-studies/). There are 3 other [case studies](https://www.cncf.io/case-studies/?_sft_lf-project=knative) published on the CNCF website, which shows the value of Knative not only as a direct usage, but also as an enabler for other CNCF projects such as [OpenFunction](https://github.com/OpenFunction/OpenFunction).

Several other CNCF and Linux Foundation projects leverage functionalities provided by Knative, such as [Tekton](https://tekton.dev/), [KServe](https://kserve.github.io/website/latest/) and [OpenFunction](https://openfunction.dev/).


## Incubation Details

### Link to Incubation Due Diligence(DD) Document

Previous material for applying to incubation:

- [Due Diligence](https://docs.google.com/document/d/1qPMyIBZ1tBk6WpEMPuLtTrjA6lvbrQ7DvCZb22S0llo/edit#heading=h.y6falyj66ofv)
- [TOC Proposal](https://github.com/cncf/toc/pull/762)

### Address any concerns or recommendations from the TAG and/or TOC sponsor(s) from the DD Document

There were no concerns raised during the incubation DD process.

