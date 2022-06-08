# cert-manager 2022 Annual Review

cert-manager provides reliable, automated Kubernetes-native management of X.509 certificates, reducing the risk of certificate-related outages and enabling PKI at scale. The project comprises cert-manager itself - which adds Kubernetes Custom Resource Definitions (CRDs) for certificate-related concepts, and controllers to manage them - along with a selection of subprojects which focus on solving other issues relating to certificate management.

The project has remained a prominent tool in the Cloud Native space since its inclusion in the CNCF sandbox this year, fuelled by wide community adoption and strong ambition to keep on growing and improving in the future. We aim for cert-manager to slot in as a de facto standard part of the Kubernetes production control plane.

We're really excited for our first annual CNCF review and we're looking forward to 2022!

_NB: Statistics referring to the last year mean "the year up to 2021-12-02"_

## DevStats

[cert-manager's DevStats page](https://certmanager.devstats.cncf.io/d/8/dashboards?orgId=1&refresh=15m&from=now-1y&to=now-1h) shows a steady rate of contributions throughout the last year.

Importantly for the project, we saw at least [45 new first-time-contributors](https://certmanager.devstats.cncf.io/d/52/new-contributors-table?orgId=1&from=now-1y&to=now&var-repogroup_name=cert-manager&viewPanel=1), only counting the main cert-manager repository.

Among those brand new contributors, cert-manager added 3 official [maintainers](https://github.com/jetstack/cert-manager/blob/71de30931fc54eba59853097b1699ffdc4bbc974/OWNERS) and two contributors who took part in Google Summer of Code 2021: [Arsh Sharma](https://summerofcode.withgoogle.com/archive/2021/projects/4983164438577152/) ([RinkiyaKeDad](https://github.com/RinkiyaKeDad)) and [Tim Ramlot](https://summerofcode.withgoogle.com/archive/2021/projects/5686178510012416/) ([inteon](https://github.com/inteon)). In addition, a previous contributor joined the maintainer team.

The source of commits is varied. A large percentage of activity at any given time is usually from the core maintainers, most of whom work at Jetstack. That said, there are regular contributions from a wide [variety of companies](https://certmanager.devstats.cncf.io/d/4/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=activity&var-repogroup_name=cert-manager&var-companies=All&from=now-1y&to=now).

In addition, the time taken for a contribution to get merged has largely [remained steady](https://certmanager.devstats.cncf.io/d/16/opened-to-merged?orgId=1&var-period=m&var-repogroup_name=cert-manager&from=now-1y&to=now).

## Maintainers

As mentioned previously, the project has added several new maintainers over the past year.

[Irbe Krumina (irbekrm)](https://github.com/irbekrm), [Ashley Davis (SgtCoDFish)](https://github.com/SgtCoDFish), [Joakim Ahrlin (jahrlin)](https://github.com/jahrlin) and [Maël Valais (maelvls)](https://github.com/maelvls) - all at Jetstack - joined the [maintainers](https://github.com/jetstack/cert-manager/blob/71de30931fc54eba59853097b1699ffdc4bbc974/OWNERS) team, bringing the total number of maintainers to 9 with one maintainer at Apple, seven at Jetstack and one at Thomas More University in Belgium.

## Adoption

cert-manager was widely adopted for use in Kubernetes clusters across many different organisations even before it joined the CNCF as a sandbox project. That adoption has only continued to grow, as evidenced by very regular activity on the [`#cert-manager`](https://kubernetes.slack.com/archives/C4NV3DWUC) channel on Kubernetes slack and the statistics presented above.

For many Kubernetes cluster administrators, cert-manager forms part of the de facto control plane they set up when they create a cluster. This carries across into suggestions or requirements by several major cloud native projects that cert-manager be installed alongside the required Kubernetes API components. Examples include [the Kubernetes Cluster API](https://cluster-api.sigs.k8s.io/developer/guide.html#cert-manager), [Knative](https://knative.dev/development/install/serving/installing-cert-manager/) and [Istio](https://istio.io/latest/docs/ops/integrations/certmanager/).

We don't closely track raw numbers of cert-manager users, either on cert-manager's [Artifact Hub package](https://artifacthub.io/packages/helm/cert-manager/cert-manager), visitors to [cert-manager.io](https://cert-manager.io) or on the cert-manager [quay.io registry](https://quay.io/repository/jetstack/cert-manager-controller).

That said, as an example of scale, the main cert-manager controller container image on quay.io reports roughly 1M-1.5M+ "actions" daily, which indicates widespread use and deployment although image caching and mirroring may make that number an underestimate of true usage. In addition, our best estimates are that the cert-manager website averages around 1.5k unique visitors per day, and that our helm charts see around 1M downloads per day, but these numbers could be out of date and again might well be undercounts.

As another rough indication of usage, the number of stars on the main cert-manager repository continues to increase steadily and linearly, passing the 8,000 star milestone in [early November 2021](https://certmanager.devstats.cncf.io/d/3/stars-and-forks-by-repository?orgId=1).

## Subprojects

cert-manager itself isn't the only project under the cert-manager umbrella; we've also been increasing the quantity and quality of subprojects and adding more open-source tooling to make it easier to manage Cloud Native certificates and identities.

These subprojects include:

- [istio-csr](https://github.com/cert-manager/istio-csr): Enables the use of cert-manager for issuing [Istio](https://istio.io/) certificates
- [approver-policy](https://github.com/cert-manager/approver-policy): Adds the `CertificateRequestPolicy` resource to define policy over whether a given certificate should be signed by cert-manager
- [csi-driver](https://github.com/cert-manager/csi-driver): Kubernetes CSI plugin to facilitate seamlessly requesting and mounting certificate key pairs into pods
- [csi-lib](https://github.com/cert-manager/csi-lib): A generalised library for building CSI drivers that interact with cert-manager's `CertificateRequest` resource
- [trust](https://github.com/cert-manager/trust): cert-manager, but for certificate trust bundles rather than for issuing certificates

Importantly, we also received the donation of the [`aws-privateca-issuer`](https://github.com/cert-manager/aws-privateca-issuer) from AWS this year and incorporated this external issuer into the cert-manager organization. We hope to start a trend of companies in the PKI / certificate issuance space working with us to assemble the best possible collection of these addons for end-users to take advantage of.

## Current Goals

In the short term, a major goal is to complete the migration of the [jetstack/cert-manager](https://github.com/jetstack/cert-manager) repo into the [cert-manager organization](https://github.com/cert-manager/).

After that, the guiding star will be to target CNCF [incubation](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage), since we believe cert-manager to be a strong candidate for incubation.

We also aim to grow our support for new features and tools in the cloud native ecosystem; recent and ongoing examples include the [Kubernetes Gateway API](https://gateway-api.sigs.k8s.io/) and the Kubernetes [CertificateSigningRequest](https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/) resource. We're always on the lookout for more in this space.

We're also in the process of:

1. Moving all of our build and test infrastructure over to the CNCF GCP organization
2. Preparing for a new cert-manager release in early 2022
3. Planning more work for the year ahead, with a focus on extensibility through neutral and user-friendly APIs.

## How the CNCF Can Help

As mentioned above, we're migrating our infrastructure to the CNCF organization. Our intention is to move one GCP project - `cert-manager-release` - and to create other projects from scratch. The migration process for the `cert-manager-release` project has been quite slow with a lot of bumps in the road; we'd really like to get that done soon.

Any other assistance would be great; given that cert-manager is so tightly tied to security, the incubation-level security audit would obviously be invaluable. In addition, marketing and project management support is always useful.

As a specific project management example, some aspects of running and maintaining an open source project mirror the joiners / movers / leavers (JML) process which many businesses have in place for onboarding and
offboarding staff members. cert-manager is starting to discover some pain points with these processes internally. For example, when a new maintainer is onboarded, how should we manage cloud access credentials for that
user in a secure and scalable way? These feel like the kinds of problem that might be common to other projects, and any advice or open-source tooling for managing these kinds of JML processes within the project would
be tremendously useful.

Put more generally, we ideally don't want to reinvent the wheel when it comes to project management or "open source JML" and we're sure that there's prior art in this space.

## Incubation

Based on the [graduation criteria](https://github.com/cncf/toc/blob/main/process/graduation_criteria.md#incubating-stage), cert-manager seems to already tick the relevant boxes required for incubation. We believe cert-manager to be a strong candidate for incubation based on the following:

- Document that it is being used successfully in production by at least three independent end users   
  cert-manager is used in production by _many_ more than three independent end users; it shouldn't be hard to find a few case studies if needed. Several major cloud native projects including [Istio](https://istio.io/latest/docs/ops/integrations/certmanager/), the [Kubernetes Cluster API](https://cluster-api.sigs.k8s.io/developer/guide.html#cert-manager) and [Knative](https://knative.dev/development/install/serving/installing-cert-manager/) suggest or require that cert-manager be installed.

- Have a healthy number of committers   
  cert-manager has 9 regular maintainers across 3 organisations. Community contributions from many different organizations are very common and gladly received.

- Demonstrate a substantial ongoing flow of commits and merged contributions   
  Evidence of this can be seen under DevStats above.

- A clear versioning scheme   
  cert-manager has had a clear versioning scheme and a regular release cadence for many months now. Details can be seen on the [cert-manager website](https://cert-manager.io/docs/installation/supported-releases/) under our "Supported Releases" page.

- Clearly documented security processes explaining how to report security issues   
  This was added in the past year, with a clear [security process](https://github.com/jetstack/cert-manager/blob/3191293cb8b475f94aa9f4c843178039c5bb6a48/SECURITY.md) and a dedicated email address (`cert-manager-security@googlegroups.com`).

If there are any concerns about cert-manager's eligibility for incubation we'd be more than happy to address them; we're hoping that this annual review will be the first step on the road towards incubation.
