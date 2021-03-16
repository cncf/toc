# OpenMetrics proposal for CNCF incubation

## Background

OpenMetrics specifies the de-facto standard for transmitting cloud-native metrics at scale, with support for both text representation and Protocol Buffers. It’s the wire format and data model that Prometheus, Cortex, Thanos and many other open source projects rely on with further more actively migrating to.

At the same time, OpenMetrics is pursuing an [RFC under IETF OPSAWG](https://datatracker.ietf.org/doc/draft-richih-opsawg-openmetrics/).

OpenMetrics was created outside of Prometheus at the request of Dan Kohn and Chris Aniszczyk in 2015 and named OpenMetrics in 2016.

OpenMetrics came out 5th place in the CNCF End User Survey on Observability tools.

OpenMetrics entered the CNCF Sandbox in July 2018.

## Progress since joining the CNCF Sandbox

OpenMetrics carefully evolved and tightened Prometheus exposition format 0.0.4.

Here are some highlights:

 * Prometheus has supported and used OpenMetrics on the wire since 2018.
 * Cortex and Thanos support it.
 * DataDog has reimplemented OpenMetrics on their own without help from OpenMetrics.
 * OpenMetrics has been submitted to IETF OPSAWG.
 * IETF OPSAWG showed very strong interest in quickly releasing OpenMetrics after a presentation at IETF 110.
 * Exemplar support has been merged into Prometheus main and will be part of the next release. This would not ever have happened without OpenMetrics. This directly enables users to correlate metrics with traces at cloud-native scale.

### Transition from Prometheus exposition format 0.0.4 to OpenMetrics 1.0

This is handled transparently by the Prometheus client libraries.

There are two changes, one breaking. They are only an issue if the wrong content type is exposed, e.g. when building metrics by hand, using printf, and serving on webpages directly.

 * Timestamps are in seconds, not milliseconds. This was the only place where Prometheus used microseconds on the wire. As timestamps are not recommended in Prometheus expositions anyway and largely considered a hack/workaround for very few and specific situations, users are very unlikely to hit this.

 * Counters must end in _total. Before, this was only strongly recommended. Client libraries handle this transparently.

### Feature roadmap for 2.x

One highlight feature will be [high-resolution histograms](https://docs.google.com/document/d/1cLNv3aufPZb3fNfaJgdaRBZsInZKKIHo9E6HinJVbpM/edit) in close coordination with Prometheus.


## Incubation 

Please see [due diligence by SIG O11y](https://docs.google.com/document/d/1BN6ABvfbAMhJthH6xNOY5Lmls3jeE3kFbMgF04BzEKE).
