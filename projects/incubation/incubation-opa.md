
# OPA Incubating Stage Review

OPA is currently a sandbox stage project. Please refer to OPA's [sandbox stage proposal](https://github.com/cncf/toc/blob/master/proposals/opa.md)
for details on the sandbox requirements.

Since entering the CNCF Sandbox, OPA has demonstrated growth and progress. 12
releases were published with 480 commits from 41 individuals. Approximately 75%
of the commits came from Styra, 7% Chef, 5% Cisco, and 13% other (compared to
93% Styra and 7% other the year before.) Across the 11 releases, several
important features were added to OPA including improved management APIs (Bundle,
Status, and Decision Logs), initial support for Rego->Wasm compilation,
Rego->SQL and Rego->Elastic Search translation, data mocking, and TLS-based
client authentication (contributed by Chef). In addition, OPA added 25 new
built-in functions for common operations like JWT decoding, date/time, CIDR,
globbing, and regexp. Most of the built-in implementations were contributed by
end-users. Lastly, several integrations were also developed: Envoy ext_authz,
Ceph, Kafka, Minio, Helm, Kong, and Flask.

In January 2019, Styra, Google, Microsoft, and others began jointly developing
and contributing the OPA [Gatekeeper](https://github.com/open-policy-agent/gatekeeper)
sub-project. Gatekeeper integrates OPA with Kubernetes to help admins enforce
admission control policies and audit clusters for existing policy violations.
Gatekeeper also includes a standard library of policies for common use cases
(e.g., registry whitelisting, ingress conflicts, label management, etc.).

Since joining the CNCF, OPA has [formalized a governance policy](https://github.com/open-policy-agent/opa/blob/master/GOVERNANCE.md) and
undergone an [external security audit](https://github.com/open-policy-agent/opa/blob/master/SECURITY_AUDIT.pdf).


# Incubating Stage Criteria

* Document that it is being used successfully in production by at least three independent end users which, in the TOC's judgement, are of adequate quality and scope. Link to the ADOPTERS.md file.
    * [https://github.com/open-policy-agent/opa/blob/master/ADOPTERS.md](https://github.com/open-policy-agent/opa/blob/master/ADOPTERS.md)
* Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.
    * Project maintainers are listed in [https://github.com/open-policy-agent/opa/blob/master/MAINTAINERS.md](https://github.com/open-policy-agent/opa/blob/master/MAINTAINERS.md)
    * Project maintainers are added and removed per the policy in [https://github.com/open-policy-agent/opa/blob/master/GOVERNANCE.md](https://github.com/open-policy-agent/opa/blob/master/GOVERNANCE.md)
* Demonstrate a substantial ongoing flow of commits and merged contributions.
    * Releases:  [https://github.com/open-policy-agent/opa/releases](https://github.com/open-policy-agent/opa/releases)
    * Roadmaps: [Current](https://docs.google.com/presentation/d/11_OBKbx8qK8bx35n3jRtluLnIBXgZwNpcJdEEGDd6YQ/edit#slide=id.g45088576f1_3_10), [Previous](https://docs.google.com/presentation/d/1P74xECzR4gYsmuYREed0DnuN1OdnGsiLwJLj_ovcqk0/edit#slide=id.p)
    * Contributors: [https://github.com/open-policy-agent/opa/graphs/contributors](https://github.com/open-policy-agent/opa/graphs/contributors)
    * Commit activity: [https://github.com/open-policy-agent/opa/graphs/commit-activity](https://github.com/open-policy-agent/opa/graphs/commit-activity)
    * CNCF devstats: [https://opa.devstats.cncf.io/](https://opa.devstats.cncf.io/)
        * Last 30 days activity on GitHub: [https://opa.devstats.cncf.io/d/8/dashboards?refresh=15m&orgId=1&from=now-30d&to=now-1h](https://opa.devstats.cncf.io/d/8/dashboards?refresh=15m&orgId=1&from=now-30d&to=now-1h)
        * Community stats: [https://opa.devstats.cncf.io/d/3/community-stats?orgId=1](https://opa.devstats.cncf.io/d/3/community-stats?orgId=1)

Further details of OPA's growth and progress since entering the sandbox stage as
well as use case details from the OPA community can be found in this [slide deck](https://docs.google.com/presentation/d/1QtQQYquIKrKg05oGPjTcqw6AZCyBU4zBTz-0_FOvF-M/edit#slide=id.g5018366126_0_21).
