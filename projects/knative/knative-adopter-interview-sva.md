# Knative Adopter Interview - SVA

Interviewee: Norris Sam Osarenkhoe, Principal Solutions Architect, SVA System Vertrieb Alexander GmbH.

Interview date: June 18, 2025

## Organization Intro


### Can you give us an overview of your organization and what it does?

SVA System Vertrieb Alexander GmbH is one of the leading system integrators in Germany in the fields of holistic IT with more than 3,200 employees at 27 branch offices. SVA focusses on the combination of high quality IT products with the project know-how and flexibility of SVA to achieve optimum solutions. Core subjects are Digital Process Solutions, Datacenter Infrastructure, IT Security, Business Continuity, SAP, Big Data and Analytics, End User Computing and Mainframe. Furthermore, SVA offers professional services around topics such as DevOps, Cloud-Native Software Development, Microsoft, IoT, SAM and many more topics.

## Motivation

### Compared with other products in this space (proprietary and open), what drew you to the project?

Open Standards are a big factor, as well as digital sovereignty. The cloud events spec was important. 2022 - we started evaluating different options, and realizing that hyperscalers were widely adopting the Cloud Events spec. On top of that, Knatives' architecture seems well thought in terms of separation of concerns between its components. E.g. in Eventing the architecture is cleanly cut into roughly three distinct areas: 
 * eventing-core - Implements the general, non-specific Knative Eventing API (Ingress, Routing, Egress)
 * eventing-messaging - Implements a generic Knative Eventing messaging model API, like Broker/Trigger or Channel/Subscription
 * eventing-messaging-bindings - Integrates a specific messaging technology on top of the messaging model API (like Kafka)
We knew if something were to happen at each level we could fix it or at least had the flexibility to swap in another solution. Additionally, Eventing offered integration with well known messaging solutions such as Kafka, NATS and RabbitMQ, which all were relevant for us. Cloud providers using these solutions as well, meant we would easily find specialist if needed.

Knative has a very plug and play architecture, which attracted us. It is very easy to onboard users. Eventing is very easy to explain, and allows building pretty complex use cases. 

We did look at some other projects, but their governance wasn’t clear. The Knative governance stood out to us. It feels like it will be a project that will be around for some time and will have maintainers for some time. 


## Usage Scenario

### How long has your organization used the project?

We started using it in 2022. It appeared to be the right tool for our problem. In order to validate our assumption we started prototyping and did lots of performance testing to better understand failure modes, performance KPIs and the inner workings.

In the summer of 2022, we started said testing and prototyping. In parallel, we started to built up other environments. As of March 2023, we have successfully released our customized Knative-backed event-mesh on our customer project infrastructure. 

2 FTEs were involved in this. 


### What were the main motivations to adopt the project and which key features do you use today?

The main motivations were:

* License 
* Open Standards
* Clean, Extensible architecture via plugins.
* Clean abstraction layer. 

HTTP for serving and eventing was very easy to integrate and adopt. 

Eventing was integrated first, the event driven architecture was used to help decouple legacy architectures. Kafka integration was a pretty crucial one.

Serving is used for more cross-cutting things within the platform, not exposed directly to end users. This is more of an enablement problem.


### What is the current level of usage (pre-production, production) and scale?

It is used in production and handles roughly over a million events daily, with an availability of 99.9% supporting approx. 11 org units.

### What version is used, and what is your update cadence with the project?

We update quarterly, using OpenShift Serverless. OpenShift Serverless aligns with the upstream Knative project, but we use the downstream release for compliance reasons. 


### Can you walk me through what your experience was in either adopting it outright or integrating it with your existing services and applications? What challenges did you experience with the project?

Originally in 2022, the documentation was a bit lacking on the administration side. We needed to look at source code to help understand. 

The operator had some issues around order of deletion of custom resources, and we had to do things like clean up finalizers. 

Another challenge was dependencies on control plane webhooks (of the Knative Operator). The Knative Operator invokes a mutating webhook for part of the reconciler loop check when certain custom resources are to be deleted. If there is a CrashLoopbackOff or these webhooks are unreachable, the operator cleanup can become difficult if you are not deeply knowledgeable of the control plane dependencies.

Adoption on the end user side was pretty easy technically speaking. Mind shift for customers adopting an asynchronous event-driven architecture was more challenging. Knative team doesn’t have experience in highly regulated industries and some testing is lacking there in my opinion. This is apparent when looking at it from an admin perspective working in a usually air gapped environment, with challenges such as access to Container Images, private Certificate Authority or custom certificates. Most challenges in my view really are to be found on the administration side of things. However, it needs to be said, that Knative has done a lot to alleviate some admin issues we had.



### Did you find the information in the repo valuable to your implementation? What specifically?

### Has your implementation of the project provided measurable value? 

The implementation has allowed us to build a solution that is entirely On-Prem, but with different environments. The ease of use and faster integrations has provided value for modernizing legacy applications at our customer site. It has cut down on onboarding quite a bit. We went from 2 weeks to onboard legacy services to one hour.  It has also allowed us to add compliance checking and enforce more compliance in said regulated environments.


### Do you have any future plans regarding the project? More involvement, feature requests, expansion, etc.

We have built some observability tooling around Knative Eventing, and would like to contribute that back. We would also like to create a blueprint for other federal agencies in Germany.


## Perception

### What is your perception in terms of the project’s:

#### Community openness

The community is very open. For the number of people involved, the Knative community achieves a lot. Lots of sub-projects under the umbrella (e.g. knative-extensions).

#### Governance

The project has great governance.

#### Community growth potential

There is a lot of interaction for some use cases. The project probably need more contributions. 

#### Maintainer diversity and ladder

The maintainer diversity could be better. It feels like mostly Red Hat is maintaining the project, based on who is at the booth (at KubeCon) and on slack. However, despite that the available amount of maintainers appears to be stable.

#### Maintainer response

The maintainers are pretty responsive. We feel very safe with Knative.

### How are you participating in the project community?

Via Slack and Github

### Did you need to engage with the community members or maintainers? If so, what was the context of the engagement and did it reach an acceptable outcome?

We have had good experience providing feedback via Red Hat.

## Project Strengths

### In your opinion, what are the overall strengths of the project?

Knative has good open standards and is easy to adopt. It is also very stable and has very good performance, we have not been able to overwhelm it or bring it to its knees and feel very confident in it covering future demand.


## Project Improvements

### Is there something you feel that holds the project back from reaching its ultimate potential?

The project could use more diversity of companies contributing to the Knative Serving side. 


### In your opinion, what can the project do better?

Better docs and guidance for admin scenarios would be better. For example, how to measure event congestion. Nothing talks about how to fine tune Kafka for Knative. More specific docs like that would be useful. It could also have better observability to support administration of deployments.