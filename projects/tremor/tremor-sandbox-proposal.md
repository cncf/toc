## Name of project:

Tremor

## Description (what it does, why it is valuable, its origin and history)

[Tremor](https://tremor.rs) is an early stage open source event processing platform designed
primarily for log and metrics processing, distribution and ETL around the needs of Wayfair's
shared logging and metrics production services.

Tremor started as a quality of service enhancement to add proactive rate limiting and reactive
backpressure handling to the Logging and Metrics shared services that run within Wayfair's
on-premise and GCP based cloud environments.

Tremor has since evolved to replace internal uses of logstash, custom gelf tooling and other
processors in our logging stacks, and telegraf and statsd in our metrics stacks.

Tremor is a lot like a system of lego. It supports pluggable sources and sinks, pluggable
processing elements that execute within the confines of a DAG-based processing pipeline.

There are three processing element variants:

* A YAML based declarative pipeline model - deprecated
* A SQL-like query language that compiles to the pipeline model ( replaces the YAML syntax )
* An embedded ETL-oriented scripting language designed around structured pattern matching of
  nested unstructured JSON-like or data.

### Features

* Pluggable onramps (sources) and offramps (sinks) for different methods of transportation
* DAG based pipelines of operators to build flow logic
* A custom scripting language optimized for event processing/ETL on structured data
* A query language designed to describe operator graphs and aggregations over them
* A system of modules, allowing reusable algorithms and processing logic to be assembled into
  libraries for reuse
* A simple autonomic API for live deployments, or static file-based deployment at application start

### Benefits

* An user friendly environment with a focus on hygenic errors and testability.
* A modular scripting and pipeline language designed to allow reusability for logic.
* A code over configuration approach with a minimal number of different operators in favour of a
  scripting language to reduce the surface area and learning curve.
* Support for a JSONesque data model with decoders and encoders to allow representing a variety
  of formats in a generalized form to reduce the number of different processing tools needed in the
  ETL department.

### Origins

Tremor was born out of necessity at Wayfair. As an organization undergoing constant volumetric
growth with regular seasonal peak demand during Cyber5 that have not been seen before we require
our network operations center to be in a position to respond to current events as they occur.

Existing proprietary logging and metrics services, at Wayfairs scale of operations, do not behave
well under extreme operating conditions - back-pressure from overloaded log/metrics visualization
systems can cascade back to transformation and source clusters; errors or bugs in provisioning of
new or changing production services can introduce unwanted logging or metrics that exacerbate
operating conditions; changing business conditions or hard failures of our on-premise or
cloud-based environment may require revising rate limits or reclassifying and redimensionalising
data based on new business objectives; and all of these conditions can occur simultaneously during
peak traffic conditions.

Tremor was designed to bring all of these observability and operational concerns under control
with a simple, well designed experience for developers and operators whilst reducing the number of
different systems under management and simplifying our production operations.

Tremor was open sourced in February 2020, after 6 months in a beta with a select number of
organizations, academic researchers and reviewers so that other organizations processing data
at extreme volumes with time-critical responsiveness and availability concerns can use, contribute
to and collaborate on the project.

## Statement on alignment with CNCF charter mission

Tremor aims to provide a high performance data processing, integration, manipulation and
distribution technology with a well defined model under failure of connected systems -- many of
them from the CNCF ecosystem.

Tremor provides a bendable system of components and processing elements specifically designed to
be plugged into existing systems or systems of systems to inject quality of service, rate limiting,
back-pressure handling and other qualities of service during extreme operating conditions, whilst
reducing the effort of developers and operators to make changes to mission critical production
systems.

## Comparison with similar projects (inside or outside the CNCF), including what differentiates this project

Tremor directly replaces uses of logstash within Wayfair and supports a small subset of Logstash's
filters such as the `KV`, `Grok` and `Dissect` facilities. However, tremor provides a well-defined
embedded scripting language with 1st class support for filtering and elementizing nested data
structures and pluggable support for elementizing embedded micro-formats.

Tremor provides a uniform developer experience for most structured nested data formats with specific
support for SIMD-accelerated processing of JSON data, which happens to be a common format in many
modern systems.

Tremor is fundamentally domain agnostic - although it is used primarily in the observability
domain - specifically in logging and metrics capture, transformation and distribution it is designed
as a general purpose event based system. Tremor supports a subset of features commonly found in
Complex Event Processors and Distributed Event Based Systems such as the Esper Stream Processing
Engine, Spark and others and over time it is expected that its feature-set relevant to CEP and DEBS
systems will evolve, mature and grow.


## Sponsor from TOC:

Chris Aniscyk

## Preferred maturity level:

Sandbox

## License

Apache License 2.0

## Source control
https://github.com/wayfair-tremor/tremor-runtime

## External dependencies (including licenses)

A full listing of external dependencies can be generated via [cargo-license](https://github.com/onur/cargo-license)

For the sake of being transparent, we’ve included the dependencies listed here at the time of writing this document.

* **Apache-2.0 (19)**: tremor-server, window, tremor-kv, tremor-tool, codespan, dissect, grok, tremor-influx, tremor-influx, tremor-api, sketches-ddsketch, tremor-script, tremor-runtime, distance, codespan-reporting, tremor-query, clang-sys, openssl, tremor-pipeline
* **Apache-2.0 AND BSD-2-Clause OR MIT (3)**: crossbeam-channel, crossbeam-queue, crossbeam-channel
* **Apache-2.0 OR Apache-2.0 WITH LLVM-exception OR MIT (1)**: wasi
* **Apache-2.0 OR BSL-1.0 (1)**: ryu
* **Apache-2.0 OR MIT (333)**: vcpkg, stable_deref_trait, semver-parser, cast, fixedbitset, reqwest, getrandom, remove_dir_all, lalrpop, base64, http, idna, cc, gloo-timers, num-bigint, ahash, tungstenite, block-padding, hashbrown, arrayvec, halfbrown, crossbeam-utils, http, fallible-iterator, thread_local, typenum, postgres-types, aes-gcm, concurrent-queue, petgraph, pin-project, geo, tempfile, sha2, time, block-cipher-trait, parking_lot_core, serde_qs, log4rs, parking_lot, lazy_static, rand_chacha, futures-sink, half, futures-timer, ghash, diff, regex-syntax, fake-simd, num-integer, time-macros-impl, unicode-normalization, rustc_version, pretty_assertions, uuid, block-buffer, async-compression, net2, dirs, sha-1, core-foundation-sys, libnghttp2-sys, parking, hmac, itoa, once_cell, rand_isaac, geojson, winapi, cookie, stdweb, futures-executor, failure_derive, blocking, rayon-core, lexical-core, dirs-sys, ascii-canvas, linked-hash-map, async-sse, iovec, num-traits, semver, hashbrown, crypto-mac, fxhash, opaque-debug, crossbeam-epoch, failure, bit-vec, lexical, ctor, humantime, yaml-rust, hkdf, siphasher, bitflags, ena, smallvec, core-foundation, scoped-tls, universal-hash, wasm-bindgen-macro-support, block-buffer, const-random-macro, serde_cbor, postgres-protocol, thiserror-impl, uuid, async-std, time-macros, parking_lot_core, futures, elastic_derive, env_logger, rand_core, crossbeam-deque, syn, futures, waker-fn, uuid, serde_yaml, futures-cpupool, input_buffer, elastic, mio-uds, backtrace, memmap, send_wrapper, object, indexmap, log, futures-core, stringprep, simd-json-derive-int, hmac, match_cfg, thiserror, cache-padded, fnv, ppv-lite86, digest, regex, error-chain, wasm-bindgen-macro, crypto-mac, bstr, rental-impl, rand_core, flate2, lock_api, rand_os, syn, aes, byte-tools, pdqselect, toml, either, futures-task, aesni, geohash, version_check, cexpr, error-chain, surf, async-h1, error-chain, futures-io, chrono, maybe-uninit, socket2, quick-error, rayon, tokio-postgres, opaque-debug, tinytemplate, autocfg, url, cookie, percent-encoding, itertools, stdweb-internal-macros, simd-json, rental, static_assertions, log-mdc, vec_map, base64, base64, unicode-xid, term, winapi-x86_64-pc-windows-gnu, url, serde_derive_internals, digest, pin-utils, const-random, rand_xorshift, standback, peeking_take_while, byte-pool, autocfg, openssl-probe, rand_pcg, hdrhistogram, num_cpus, ahash, futures-util, percent-encoding, gimli, unicode-bidi, crossbeam-utils, kv-log-macro, idna, async-trait, async-task, if_rust_version, unicode-xid, rand_pcg, dirs, proc-macro-nested, shlex, unicode-width, zeroize, sha2, hashbrown, parking_lot, wasm-bindgen-futures, wasm-bindgen-shared, sled, http-client, aead, security-framework, scoped-tls, md5, value-trait, quote, serde_json, rustc-hash, anyhow, proc-macro2, wasm-bindgen, rand, hermit-abi, pin-project-lite, string_cache, bumpalo, polyval, rand_core, arc-swap, rle-decode-fast, serde, stdweb-internal-runtime, pkg-config, ntapi, serde_urlencoded, crc32fast, downcast-rs, derivative, proc-macro-hack, smallvec, serde_derive, utf-8, glob, threadpool, traitobject, rand_hc, synom, rust-argon2, libc, async-native-tls, simd-json-derive, criterion, rand, cpuid-bool, parking_lot, lazycell, wasm-bindgen-backend, libz-sys, encoding_rs, scopeguard, winapi-i686-pc-windows-gnu, rustc-demangle, smol, cmake, xz2, http-types, mime, security-framework-sys, criterion-plot, fs2, pin-project-internal, stdweb-derive, serde_urlencoded, web-sys, miow, generator, foreign-types-shared, rand_chacha, rand, human-panic, femme, httparse, postgres, geo-types, parking_lot_core, time, publicsuffix, rstar, lzma-sys, lock_api, rand_jitter, fastrand, rand_hc, cfg-if, yaml-rust, cookie_store, aes-soft, unicase, thread-id, miow, futures-macro, lalrpop-util, tide, addr2line, dtoa, foreign-types, js-sys, proc-macro-crate, native-tls, quote, futures-channel, bit-set
* **BSD-2-Clause (3)**: cloudabi, cloudabi, arrayref
* **BSD-3-Clause (10)**: instant, fuchsia-zircon-sys, num_enum, snap, num_enum_derive, bindgen, subtle, fuchsia-zircon, sha1, subtle
* **CC0-1.0 (1)**: constant_time_eq
* **ISC (3)**: jumphash, libloading, rdrand
* **MIT (104)**: ws2_32-sys, tokio-fs, ansi_term, ordered-float, rdkafka, fluent_builder, tokio-tcp, string, synstructure, h2, winapi, matches, bytes, libflate_lz77, try_from, try-lock, phf_shared, libflate, redox_syscall, base-x, hostname, sluice, snmalloc-rs, tracing-attributes, tracing-futures, cron, schannel, miniz_oxide, kernel32-sys, difference, clap, float-cmp, tokio-io, tracing, isahc, lz4-sys, tokio-sync, rdkafka-sys, generic-array, output_vt100, nom, tokio-codec, mime_guess, onig, curl, hyper, tokio-udp, enum-ordinalize, tokio-uds, nom, slab, tokio-reactor, memoffset, precomputed-hash, winapi-build, tokio-timer, redox_users, serde-value, lz4, rmp-serde, http-body, data-encoding, want, phf, cidr-utils, tracing-core, tokio, blake2b_simd, loom, infer, tokio-current-thread, route-recognizer, unsafe-any, mio, tokio-util, atty, tokio-executor, generic-array, openssl-sys, new_debug_unreachable, bytes, tokio-threadpool, snmalloc-sys, strsim, nom, which, strsim, lru, tokio, typemap, onig_sys, os_type, mio, tokio-buf, winreg, discard, plotters, async-tungstenite, textwrap, debug-helper, rmp, float-cmp, curl-sys, oorandom
* **MIT OR Unlicense (11)**: regex-automata, memchr, termcolor, same-file, docopt, winapi-util, csv, aho-corasick, walkdir, byteorder, csv-core
* **MPL-2.0 (1)**: wepoll-sys-stjepang
* **N/A (1)**: fuchsia-cprng *(note, these are c-bindings for the fuchsian random number generator which itself is under the BSD license: https://fuchsia.googlesource.com/fuchsia/+/master/garnet/public/rust/fuchsia-cprng/src/lib.rs)*
* **Zlib (2)**: adler32, tinyvec


## Initial committers

* Darach Ennis (@darach, Wayfair)
* Heinz Gies (@Licenser, Wayfair)
* Anup Dhamala (@anupdhml, Wayfair)
* Ernad Halilovic (@ernadh, Wayfair)
  
## Infrastructure requests (CI / CNCF Cluster)

None.

## Communication channels (slack, irc, mailing lists)
* Slack chat: https://chat.tremor.rs/
* Website: https://www.tremor.rs/
* Documentation: https://docs.tremor.rs/
* RFCs: https://rfcs.tremor.rs/

## Issue tracker (GitHub by default)

http://github.com/wayfair-tremor/tremor-runtime/issues

## Website (current version will move to project.cncf.io, see here for guidelines)

http://www.tremor.rs/

## Release methodology and mechanics

Releases are done though GitHub Actions, where docker packages are built and published to the
tremor-project namespace in Docker Hub.

## Social media accounts

* https://twitter.com/TremorDebs

## Community size and any existing sponsorship
_(07, 13, 2020)_

- 180 Stars
- 18 Forks
- 7 Contributors
- 13 Releases

## Adopters:

Not currently formally tracked

## Project logo in svg format
![Tremor SVG logo](https://docs.tremor.rs/img/logo.svg)