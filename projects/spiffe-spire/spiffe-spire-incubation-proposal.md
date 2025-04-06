# SPIFFE/SPIRE Incubation Application

The SPIFFE and SPIRE projects were open-sourced in December 2017. The projects were accepted into the CNCF in March 2018. The projects have grown significantly over time.

The following application links to the required information to become an incubation project. Additionally, a due diligence document has been prepared to assist in the proposal review is accessible at: https://docs.google.com/document/d/1tkN9YgBSLEUszOflWPHO72qedOaUb3iHfAye45dKJT8/


## SPIFFE/SPIRE fulfills all the incubation criteria:
### Document that it is being used successfully in production by at least three independent end users which, in the TOC’s judgement, are of adequate quality and scope.

Known end-users include Bloomberg, Bytedance, Pinterest, Square, Uber, and Yahoo Japan. This is not an exclusive list and it’s likely that other organizations are also SPIFFE/SPIRE users.

The following organizations actively plan to roll out SPIRE over the next 12 months: Github, Twilio.
 
SPIFFE/SPIRE is being used by numerous other companies, both large and small, to build higher layer products and services. This includes Decipher Technology Studios, F5 Networks, HashiCorp, Intel, Google, IBM, Tigera, VMware, and many others.

### Have a healthy number of committers. A committer is defined as someone with the commit bit; i.e., someone who can accept contributions to some or all of the project.

See the current list of SPIFFE owners and SPIRE owners, who are also committers. Currently, SPIFFE has 5 from 5 different organizations. SPIRE has 7 from 2 different organizations.

### Demonstrate a substantial ongoing flow of commits and merged contributions.

The CNCF provides SPIFFE and SPIRE Github metrics through its devstats site: 

https://spiffe.devstats.cncf.io/d/1/activity-repository-groups?orgId=1

https://spire.devstats.cncf.io/d/1/activity-repository-groups?orgId=1

#### SPIFFE spec

|            |         |
|------------|---------|
| Commits    | 298     |
| Committers | 5       |
| Commits    | Avg: 5 Max: 29 |
| Stars      | 586     |



#### SPIRE

|                        |                |
|------------------------|----------------|
| Commits                | 2,892          |
| Committers             | 7              |
| Contributors           | Avg: 7 Max: 50 |
| Contributing Companies | Avg: 3 Max: 27 |
| Stars                  | 552            |
| CII Progress           | Passing        |

### A clear versioning scheme

SPIRE operates on a 90 day major release cadence. SPIRE is currently on version 0.9.3. There are regular releases with an up to date change log.

The SPIFFE specification is currently unversioned.

### Public Reference to the spec

The SPIFFE spec defines a framework and set of standards for identifying and securing communications between web-based services. At its core, SPIFFE is:
A standard defining how services identify themselves to each other. These are called SPIFFE IDs and are implemented as Uniform Resource Identifiers (URIs).
A standard for encoding SPIFFE IDs in a cryptographically-verifiable document called a SPIFFE Verifiable Identity Document or SVIDs.
An API specification for issuing and/or retrieving SVIDs. This is the Workload API.

### Link to 3rd Party Plugins and Integrations

SPIFFE and SPIRE have integrations available with a number of open-source projects. The list includes:

Kubernetes
Docker
Envoy
Istio
NGINX
Ghosttunnel
Knox	
Consul

### Asks to CNCF

Invest in marketing initiatives to educate users and grow awareness of cloud native security, to foster community, and to accelerate production use of SPIFFE/SPIRE.

Developer community support: hangouts, meetups, and events.
Digital marketing: Help with online content, interactive tutorials, webinars, and social
Product marketing: conference promotion, landscape, certification, case studies, AR/PR.
Help conduct a formal third party security audit.  
