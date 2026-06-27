# Relationship to adjacent work

Runtime Conditions Profiles are not intended to replace workload specifications, package formats, service catalogs, or provisioning systems.

The profile describes what one workload needs at runtime: APIs it calls, datastores it depends on, caches it expects, configuration inputs it consumes, and similar integration requirements. It does not describe how to provision those dependencies, where they run, what credentials are used, etc.

This makes the profile complementary to existing CNCF and cloud-native work:

- Score and Radius can use profile data as input to workload or application modeling, while still owning platform-specific resolution.
- Backstage and API catalogs can use API Conditions to validate declared API usage against cataloged contracts. Similarly, Backstage can be used to store Runtime Conditions Profiles once they have been generated.
- Crossplane, Kratix, Dapr, Microcks, and other platform tools can act on validated profile requirements, but their resource models remain separate.
- CNAB and Porter can package or distribute software artifacts that include profiles, as well as use the profile to determine install, upgrade, and uninstall workflows.

The intended boundary is simple: the Runtime Conditions Profile is a portable requirements document; downstream adapters and platforms turn those requirements into actions.
