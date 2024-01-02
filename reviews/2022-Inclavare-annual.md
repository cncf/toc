# Background

[Inclavare Containers project](https://github.com/inclavare-containers/) initially provides a process-level confidential container solution based on Intel SGX Enclave. Later, this project was re-organized from https://github.com/alibaba/inclavare-containers to https://github.com/inclavare-containers/inclavare-containers since 02/2022.

The development efforts on the part of enclave container runtime was moved to [enclave-cc](https://github.com/confidential-containers/enclave-cc) under another CNCF project [Confidential Containers](https://github.com/confidential-containers) since 03/2022. The reason is that Inclavare Containers project can share [some common components](https://github.com/confidential-containers/enclave-cc/blob/main/docs/design.md#architecture) from Confidential Containers project, so this helps to reduce repetitive development efforts and keep active communication with neighbor community to collectively boost the adoption of confidential containers technology.

Inclavare Containers project doesn't just focus on container runtime for Intel SGX Enclave, it also creates several innovative technologies related to confidential container (See the [Goals](#goals) section for the details). Therefore, the most of development efforts on this project are contributed to these sub-projects under the organization https://github.com/inclavare-containers/.

# DevStats

The DevStats statistics can be found [here](https://inclavarecontainers.devstats.cncf.io/d/1/activity-repository-groups?orgId=1&from=1631030400000&to=1662652799000&var-period=d7&var-repogroups=All).

Note: 
- Most of development efforts are contributed to the sub-projects under the organization [Inclavare Containers](https://github.com/inclavare-containers/) and [enclave-cc](https://github.com/confidential-containers/enclave-cc) rather than https://github.com/inclavare-containers/inclavare-containers.
- The statistics for the development efforts on the part of enclave container runtime is moved to [here](https://github.com/confidential-containers/enclave-cc/graphs/contributors?from=2022-03-01&to=2022-11-08&type=c)

# Maintainers

The maintainer list can be found [here](https://github.com/inclavare-containers/inclavare-containers/blob/master/MAINTAINERS.md).

# Adoption

- Inclavare Containers' enclave container runtime is now adopted as Alibaba Cloud Container Service for Kubernetes (ACK), which can be deployed in production and commercial use [TEE-based ACK](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/deploy-confidential-containers-in-tee-based-ack-clusters#section-fld-vni-wnc) in Alibaba Cloud.

- Inclavare Containers' enclave attestation architecture is now adopted in [TEE-based ACK](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/use-confidential-containers-to-implement-remote-attestation-in-tee-based-ack-clusters) to provide the capability of remote attestation.

- Inclavare Containers' enclave attestation architecture is used to support [remote attestation](https://github.com/confidential-containers/attestation-agent/tree/main/src/kbc_modules/eaa_kbc) for Confidential Containers project, and [Verdictd](https://github.com/inclavare-containers/verdictd) from Inclavare Containers project is mentioned in Confidential Containers's [initial official release](https://github.com/confidential-containers/documentation/blob/main/releases/v0.1.0.md#limitations).

- [rats-tls](https://github.com/inclavare-containers/rats-tls) from Inclavare Containers is now supported by [Alibaba Cloud Linux](https://help.aliyun.com/document_detail/411687.html) for SGX remote attestation.

# Goals

Inclavare Container project is targeting to incubate some innovative sub-projects for confidential containers ecosystem.

- Enclave Attestation Architecture (EAA)：support hierarchical and mutual attestation across different TEEs
  - [librats](https://github.com/inclavare-containers/librats): Low level attester and verifier drivers for remote attestations on multiple TEE platforms, e.g, Intel SGX DCAP, AMD SEV, SEV-ES and SEV-SNP, Intel TDX and Hygon CSV.
  - [rats-tls](https://github.com/inclavare-containers/rats-tls): [IETF RATS architecture](https://datatracker.ietf.org/wg/rats/about/) based TLS using librats.
  - [etpm](https://github.com/inclavare-containers/etpm): Protected vTPM backend using TEE approach
  - [attestation evidence broker](https://github.com/inclavare-containers/attestation-evidence-broker): A general attestation-assisted service
  - [eventlog-rs](https://github.com/inclavare-containers/eventlog-rs): Event Log utility for confidential computing

- Software Supply Chain for confidential container: use reproducible / deterministic build to verify the evidence manifest from TEEs
  - [rbi](https://github.com/inclavare-containers/rbi): The scripts to support reproducible build for kata containers.
  - [deterministic builds](https://github.com/inclavare-containers/deterministic-builds): eBPF-based tool for deterministic builds.

- Transparent network traffic gateway over attested channel using TEEs:
  - [enclave network gateway](https://github.com/inclavare-containers/enclave-network-gateway)
 
The developement work for the part of enclave container runtime will continue in [enclave-cc](https://github.com/confidential-containers/enclave-cc).

# Help from CNCF

One of our primary aims is to increase Inclavare Containers' visibility and adoption, together with sharing the view of process-level confidential containers to advance the Confidential Computing.

With the help of CNCF, this project has gained more attention from the community and more developers have invested in the co-construction.

This would help with the targets mentioned in [Goals](#goals).

# Incubation

Now Inclavare Containers is not ready for incubation yet.
