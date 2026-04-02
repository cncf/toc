# \[WHITEPAPER\] TAG ENV / WG AI 

# Foundations & Framework for Sustainable AI in Cloud Native Environments

<div align="center">

![][image1]

</div>

| TAG / Group Home: [https://tag-runtime.cncf.io/wgs/cnaiwg/](https://tag-runtime.cncf.io/wgs/cnaiwg/)  Authors (listed alphabetically): Adel Zaalouk Andrew Block, Red Hat Marisol Palmero Nimisha Mehta Payam Mohajeri Prateek Kumar Viktor Lu Vincent Caldeira, Red Hat | GitHub Issues: [CNCF TOC Issue 1675](https://github.com/cncf/toc/issues/1675) |
| :---- | :---- |

Table of Contents

[Introduction](#introduction)

[Whitepaper Objectives](#whitepaper-objectives)

[Importance of Sustainable AI in Cloud Native Environments](#importance-of-sustainable-ai-in-cloud-native-environments)

[CNCF’s Role in Sustainable Cloud Native AI](#cncf’s-role-in-sustainable-cloud-native-ai)

[Reference Model for Sustainable AI Design](#reference-model-for-sustainable-ai-design)

[Integrating Sustainability into the AI Lifecycle](#integrating-sustainability-into-the-ai-lifecycle)

[AI System Taxonomy and Sustainability Implications](#ai-system-taxonomy-and-sustainability-implications)

[AI Infrastructure & Deployment Environments](#ai-infrastructure-&-deployment-environments)

[Personas & Responsibilities in Sustainable AI](#personas-&-responsibilities-in-sustainable-ai)

[Organizational Context & Governance](#organizational-context-&-governance)

[Conclusion](#conclusion)

# Introduction {#introduction}

## Whitepaper Objectives {#whitepaper-objectives}

AI adoption delivers clear benefits but can increase environmental impacts through the energy and infrastructure required to train and run models. In cloud native environments, these impacts are strongly influenced by architecture, orchestration, and operational practices. This paper builds on the CNCF Environmental Sustainability TAG[^1] and provides a reference framework for integrating sustainability into cloud native AI, aligned with open source principles. The objectives of this whitepaper are to define a comprehensive reference framework for integrating environmental sustainability throughout the AI lifecycle within cloud-native environments, provide practical guidance and actionable levers across the design, deployment, and operational stages of AI systems, and align sustainability practices with existing open source tools and cloud-native orchestration principles, particularly within the CNCF ecosystem. Readers should be familiar with cloud native basics, AI system concepts, and sustainability fundamentals.[^2]

## Importance of Sustainable AI in Cloud Native Environments {#importance-of-sustainable-ai-in-cloud-native-environments}

AI workloads are becoming a structural driver of energy, water, and infrastructure demand in cloud native platforms. In 2024, global data centers consumed approximately 415 TWh of electricity, representing about 1.5% of total global electricity use, following an average growth rate of around 12% per year over the previous five years.[^3] This growth exceeds overall electricity demand, which increased by about 4.3% year-over-year in 2024, compared to roughly 2.5% in 2023.[^4]

At the same time, the computational intensity of AI systems continues to rise. The Stanford AI Index 2025 reports that training compute for notable AI models now doubles approximately every five months, while the power required for training increases annually. Although hardware efficiency continues to improve, these gains are currently outpaced by increases in model scale, training frequency, and deployment volume.[^5]

The impact is global. The International Energy Agency projects that data-center electricity demand will nearly double to around 945 TWh by 2030, approaching 3% of global electricity consumption, with significant growth across North America, Europe, and Asia-Pacific.3

<div align="center">

![][image2]

</div>

## CNCF’s Role in Sustainable Cloud Native AI {#cncf’s-role-in-sustainable-cloud-native-ai}

As AI adoption accelerates, its energy, carbon, and water impacts are becoming operational sustainability challenges. The CNCF ecosystem provides widely used cloud native building blocks to deploy, observe, and optimize AI workloads across distributed environments, aligned with CNCF’s mission to make cloud native computing ubiquitous.[^6]

Kubernetes, CNCF’s foundational project, is a portable platform for managing containerized workloads and services, including large-scale AI systems.[^7] Real-world case studies show Kubernetes-based platforms being used to scale and adapt AI infrastructure to demand.[^8] Improvements in autoscaling and cluster efficiency can also reduce idle capacity and cost.[^9] The table below summarizes how selected CNCF projects can support sustainable AI use cases in cloud native environments.

Table 1 \- CNCF projects supporting sustainable cloud native AI

| CNCF Project | Sustainability function | AI-related use case |
| :---- | :---- | :---- |
| Kepler [^10] | Power and energy telemetry | Track energy use per container/pod |
| KEDA [^11] | Event-driven autoscaling | Reduce idle compute for inference |
| OpenCost [^12] | Cost transparency | Map cost to energy footprint |
| Volcano [^13] | Batch scheduling for AI | Optimize GPU allocation |
| OpenTelemetry [^14] | Metrics & observability | Monitor sustainability KPIs |
| Kueue [^15] | Job queuing & resource sharing | Efficient resource utilization for AI training |

Taken together, these projects make sustainability actionable within the existing cloud native control plane: measure resource use, expose it as telemetry, and use scaling and scheduling to reduce waste. This enables teams to integrate sustainability KPIs into operational workflows (dashboards, alerts, policy gates) rather than treating them as an external reporting step.[^16]

To embed sustainability in AI operations alongside CNCF projects, the ISO/IEC 21031:2024 (Software Carbon Intensity, SCI)[^17] standard could be adopted. This specification provides a methodology to measure software-generated emissions, accounting for not only operational energy consumption but also the amortized embedded carbon from hardware's full lifecycle (production, transport, reuse, disposal). This allows reducing carbon through less/cleaner energy, responsible hardware production, extended hardware lifespan, or efficient software. An extension is being developed for AI-specific workloads, covering the AI lifecycle and establishing boundaries to fairly attribute emissions between consumers and providers for responsible reporting.[^18]

# Reference Model for Sustainable AI Design {#reference-model-for-sustainable-ai-design}

Designing sustainable and resource-efficient AI requires a shared view of the design dimensions that influence energy, carbon, water, and hardware impacts. The AI Sustainability Reference Model organizes these dimensions to clarify what decisions need to be made, where sustainability interventions apply, and how responsibilities and governance connect to technical choices.

Table 2 \- AI Sustainability Reference Model dimensions

| Dimension | What it covers | Sustainability implications and levers |
| :---- | :---- | :---- |
| AI lifecycle | Stages from data sourcing to end-of-life | Impacts vary by stage; focus on data efficiency, training/inference efficiency, operational monitoring, and decommissioning practices |
| AI system taxonomy | System type and scale (e.g., small ML, RAG, foundation models, agentic) | Larger and more complex systems typically increase compute and serving demand; levers include right-sizing, efficient tuning, distillation, and usage controls |
| Infrastructure and deployment environment | Where the system runs (cloud, on-prem, edge, on-device) and the underlying resources | Energy mix, cooling efficiency, and utilization shape the footprint; levers include efficient placement, orchestration, and carbon/energy-aware scheduling |
| Personas and responsibilities | Roles across build and run (data, ML engineering, platform, ops, app teams) | Clear ownership prevents gaps; levers include role-specific accountability for efficiency, telemetry, capacity management, and product usage patterns |
| Organizational context and governance | Policies, incentives, procurement, measurement, and reporting | Outcomes depend on what is measured and rewarded; levers include KPIs/OKRs, standard metrics, auditability, and governance integrated into delivery workflows |

Together, these dimensions connect sustainability decisions to the AI lifecycle, clarify ownership across teams, and establish the governance mechanisms needed to keep sustainability requirements enforceable and measurable over time.

## Integrating Sustainability into the AI Lifecycle {#integrating-sustainability-into-the-ai-lifecycle}

Sustainability decisions occur throughout the AI lifecycle. The table below summarizes common resource drivers, practical levers, and example metrics for each stage.

Table 3 \- AI lifecycle sustainability mapping: drivers, levers, metrics

| Lifecycle stage | Key resource drivers | Practical sustainability levers | Example metric |
| :---- | :---- | :---- | :---- |
| Data collection and preparation | storage growth, preprocessing compute, repeated dataset copies | deduplicate datasets, tiered storage, efficient data formats, cache-aware pipelines | data processed per kWh |
| Feature, prompt, and retrieval design | embedding generation, retrieval calls, long contexts, repeated prompts | prompt minimization, caching, batching, retrieval tuning, context limits | tokens per request |
| Model selection and training/tuning | GPU/accelerator time, high-utilization clusters, long runs | right-size model, PEFT, distillation, early stopping, efficient dataloading | training kWh per run |
| Evaluation and validation | large test suites, repeated runs, human-in-the-loop cycles | staged evaluation, sampling, automated gates, reuse artifacts | eval cost per release |
| Deployment and serving | always-on replicas, peak provisioning, GPU fragmentation | autoscale to demand, request batching, quantization, model routing | energy per inference |
| Monitoring, maintenance, and retraining | continual telemetry, drift detection, frequent retraining | trigger retraining by signals, incremental updates, lifecycle budgets | retrains per incident |
| End-of-life and decommissioning | unused endpoints, stale models, retained data | retire unused models, delete artifacts, enforce retention policies | stale endpoints removed |

## AI System Taxonomy and Sustainability Implications {#ai-system-taxonomy-and-sustainability-implications}

AI systems differ substantially in how they consume compute, memory, storage, networking, and (in some cases) physical energy. A practical taxonomy helps teams anticipate the dominant resource drivers and select sustainability strategies that reduce waste without compromising reliability or safety.

Table 4 \- AI system types, resource drivers, and sustainability strategies

| AI System Type | Key Resource Drivers | Primary Sustainability Strategy |
| :---- | :---- | :---- |
| Vision Models | GPU-intensive training and inference | Model pruning, quantization |
| Speech Models | Real-time inference and low-latency serving | Streaming-aware inference, right-sized serving |
| LLMs / Generative | Scale-dependent training and high-volume inference | PEFT, distillation, retrieval minimization |
| Robotics / Agentic | Physical \+ digital energy; continuous sensing/actuation | On-device inference, control loop efficiency |
| Forecasting / Planning | Frequent retraining; repeated batch inference | Event-based retraining, incremental updates |

This taxonomy is intended as an operational lens: it highlights where sustainability efforts typically have the highest leverage (model architecture, serving patterns, retraining frequency, or edge vs cloud placement), and supports consistent measurement and governance across heterogeneous AI workloads.

## AI Infrastructure & Deployment Environments {#ai-infrastructure-&-deployment-environments}

Where an AI system runs largely determines its sustainability profile. Deployment choices influence utilization (idle vs busy resources), cooling efficiency, the electricity grid mix, and how much data must move across networks. The same model can therefore have very different operational impacts depending on whether it is served on-device, on-prem, in a public cloud region, or across a hybrid footprint.

Deployment decisions are also constrained by hardware reality. Some organizations operate older accelerators, while others lack accelerator capacity entirely and cannot obtain enough GPUs to meet demand. The AI Index 2025 highlights that cutting-edge AI increasingly requires compute and financial resources that are not available to academia, with leading models predominantly produced by industry.5 This access gap matters for sustainability: many efficiency strategies depend on having the right hardware and then keeping it highly utilized. The table below summarizes common environment trade-offs and the levers typically available in each setting.

Table 5 \- Environment trade-offs, levers, and hardware access

| Environment | Common fit | Dominant sustainability drivers | Primary levers | Hardware access considerations |
| :---- | :---- | :---- | :---- | :---- |
| Public cloud | Elastic training/inference; burst capacity | Region energy mix; overprovisioning; data egress | Right-sizing; autoscaling; placement policies; energy telemetry | Access via rentals; quotas/capacity can constrain peaks |
| Private / on-prem | Data control; predictable workloads | Utilization; cooling efficiency; upgrade cycles | Consolidation; scheduling; instrumentation; capacity management | May have older accelerators or no accelerators available 5 |
| Hybrid / multi-cloud | Mixed constraints; placement flexibility | Telemetry consistency; duplicated data; sprawl | Standard metrics; placement rules; governance controls | Balances constraints; requires strong guardrails |
| Edge / on-device | Low latency; local autonomy | Device power limits; model size; update cadence | Small models; quantization; hardware-aware inference | Reduces reliance on centralized accelerators; limits model scale |

To make sustainability improvements repeatable, teams typically run a simple operational loop:

1. **Measure**: Collect workload-level signals (energy, utilization, carbon proxies).  
2. **Observe**: Dashboards and alerts; trends and anomalies.  
3. **Decide**: Placement, scaling, scheduling, and throttling policies.  
4. **Act**: Autoscale, binpack, batch, or shift workloads.  
5. **Report**: Sustainability KPIs and governance evidence; feed back into measurement.

Even as hardware becomes more efficient, overall demand can still rise due to increased model scale and usage. The AI Index 2025 reports rapid improvements in hardware energy efficiency, while also noting that the power required for training has continued to increase.5 This reinforces why the deployment environment and the control plane matter: sustainability improvements depend on measuring the right signals and then using orchestration and scheduling to reduce idle capacity and unnecessary data movement.

For Kubernetes environments, Kepler-based approaches are one practical path to connect workload operations to energy-aware optimization workflows.16 At the systems level, the IETF GREEN working group provides a standards-oriented framing for energy measurement and control in ICT systems.[^19]

## Personas & Responsibilities in Sustainable AI {#personas-&-responsibilities-in-sustainable-ai}

Sustainable AI is a shared responsibility across teams that build, operate, and govern AI systems. Clear role ownership helps ensure sustainability signals are measured consistently, optimization work is prioritized, and progress is tracked with operational KPIs.

Table 6 \- Personas, responsibilities, and example KPIs

| Persona | Responsibility | Example KPI |
| :---- | :---- | :---- |
| Data Scientist | Optimize dataset and model size; reduce unnecessary features | Compute hours saved |
| ML Engineer | Build efficient training and inference pipelines | GPU utilization rate |
| DevOps / SRE | Operate energy-aware platforms; reduce idle capacity | Carbon per inference |
| Platform Engineer | Provide reusable telemetry and optimization primitives | Percent workloads with energy telemetry |
| Product Owner | Manage demand and feature usage patterns | Inferences per user action |
| Governance Lead | Define incentives, controls, and compliance evidence | Sustainability OKRs met |

This table can be adapted per organization, but each persona should have at least one measurable KPI that links day-to-day decisions to sustainability outcomes.

## Organizational Context & Governance {#organizational-context-&-governance}

Technology choices alone do not ensure sustainable AI outcomes. Incentives, ownership, measurement practices, and policy controls determine whether sustainability is treated as an operational requirement or an optional optimization. The AI Index 2025 notes that AI governance oversight is distributed across functions (no single department dominates), reinforcing the need for clear accountability and an operating model that spans technical and non-technical stakeholders.5 The table below summarizes governance levers that make sustainability measurable, enforceable, and continuously improved.

Table 7 \- Governance levers for sustainable AI operations

| Governance area | What to define | Practical mechanism | Example artifact / KPI |
| :---- | :---- | :---- | :---- |
| Incentives and KPIs | What “good” looks like | OKRs, FinOps/Sustainability reviews | carbon per inference; GPU utilization |
| Measurement and observability | What to measure and how | standard metrics, dashboards, alerts | percent workloads with energy telemetry |
| Workload placement and scheduling | Where and when workloads run | policy-driven placement, scheduling rules | share of batch jobs time-shifted |
| Procurement and capacity strategy | What to buy or rent | performance-per-watt criteria, lifecycle planning | accelerator utilization; refresh cadence |
| Policy, risk, and compliance | Required controls and evidence | documentation, logging, audits | audit trail completeness |
| Continuous improvement | How progress is sustained | recurring reviews, postmortems, education | KPI trend improvements quarter-over-quarter |

Workload placement can be optimized not only by location (region) but also by time. For flexible workloads (e.g., training, batch inference, retraining pipelines), organizations can dynamically shift execution to hours when grid electricity is cleaner or renewable availability is higher, using carbon-intensity forecasts and scheduling policies. This approach is already used in practice in large-scale “carbon-aware computing” initiatives and is increasingly supported by open tooling and forecast APIs.[^20] [^21] [^22]

Governance is also shaped by regulation. The EU AI Act requires providers of certain systems to maintain technical documentation and related compliance evidence, and official guidance highlights documentation elements and lifecycle maintenance expectations.[^23] Regardless of jurisdiction, the control pattern is consistent: define sustainability metrics, measure them continuously, enforce policies through automation where possible, and retain evidence for review and audit.21

# Conclusion {#conclusion}

Sustainable AI in cloud native environments is primarily an operations and governance problem: AI demand is rising, and the infrastructure footprint depends on how systems are deployed, measured, and continuously optimized. Recent analyses show data-centre electricity demand growing rapidly and projected to increase substantially toward 2030, with AI as a major driver.3 At the same time, leading AI systems continue to increase compute requirements, reinforcing the need to treat sustainability as part of the platform control plane rather than an external reporting activity.5

To make sustainability actionable, telemetry must be used to trigger concrete optimization decisions. The CNCF ecosystem provides the primitives to measure and act (energy telemetry, cost signals, autoscaling, scheduling, and observability), but the goal is to close the loop between measurement and change.16 Actionable steps for sustainable cloud native AI:

1. Embed governance and evidence collection  
2. Use telemetry to reduce footprint (not just observe it) and even increase handprint  
3. Make placement a policy decision  
4. Run continuous improvement cycles  
5. Define a small set of operational sustainability KPIs  
6. Instrument and standardize telemetry

The practical outcome of this approach is a measurable loop: telemetry reveals waste (idle accelerators, oversized services, inefficient placement), and platform controls convert those insights into action (autoscaling, scheduling, placement, and time shifting). This makes sustainability improvements repeatable across teams and workloads, while aligning operational behavior with the growing scale of AI and evolving governance expectations.3 5 23  
<div align="center">

![][image3]

</div>

[^1]:  [CNCF Environmental Sustainability TAG](https://tag-env-sustainability.cncf.io/)

[^2]:  [Green Software for Practitioners (LFC131)](https://training.linuxfoundation.org/training/green-software-for-practitioners-lfc131/)

[^3]:  [IEA, Energy and AI: data-centre electricity growth](https://www.iea.org/reports/energy-and-ai/energy-demand-from-ai)

[^4]:  [IEA, Electricity 2025: global demand trends](https://www.iea.org/reports/electricity-2025/demand)

[^5]:  [Stanford HAI, AI Index 2025: compute and energy trends](https://aiindex.stanford.edu/report/)

[^6]:  [CNCF charter: mission statement](https://www.cncf.io/about/who-we-are/)

[^7]:  [Kubernetes overview: manage workloads and services](https://kubernetes.io/docs/concepts/overview/)

[^8]:  [Kubernetes: OpenAI case study](https://kubernetes.io/case-studies/openai/)

[^9]:  [Anthropic: 40% AWS bill reduction](https://www.thestack.technology/aws-anthropic-cloud-bill-eks-karpenter/)

[^10]:  [Kepler: eBPF energy metrics exporter](https://www.cncf.io/projects/kepler/)

[^11]:  [KEDA: event-driven autoscaling](https://keda.sh/)

[^12]:  [OpenCost: Kubernetes cost monitoring](https://opencost.io/docs/)

[^13]:  [Volcano: batch scheduling for Kubernetes](https://www.cncf.io/projects/volcano/)

[^14]:  [OpenTelemetry: traces, metrics, logs](https://opentelemetry.io/docs/what-is-opentelemetry/)

[^15]:  [Kueue Kubernetes-native Job Queueing](https://kueue.sigs.k8s.io/)

[^16]:  [Grafana Labs, Kubernetes, Kepler, and carbon observability tools](https://grafana.com/blog/kubernetes-kepler-and-carbon-footprints-the-latest-tools-and-strategies-to-optimize-observability/)

[^17]:  [ISO/IEC 21031:2024 \- Software Carbon Intensity (SCI) specification](https://www.iso.org/standard/86612.html)

[^18]:  [SCI for AI \- Software Carbon Intensity for Artificial Intelligence | Green Software Foundation](https://sci-for-ai.greensoftware.foundation/)

[^19]:  [IETF Green WG](https://datatracker.ietf.org/group/green/about/) 

[^20]:  [Google: shift compute to cleaner hours/regions](https://blog.google/outreach-initiatives/sustainability/carbon-aware-computing-location/)

[^21]:  [Microsoft: time and location shifting in Azure](https://devblogs.microsoft.com/ise/saving-co2-using-location-and-time-shifting-in-azure/)

[^22]:  [Electricity Maps: carbon intensity history/forecast API](https://app.electricitymaps.com/developer-hub/api/reference)

[^23]:  [European Union, AI Act: risk, transparency, obligations](https://eur-lex.europa.eu/eli/reg/2024/1689/oj)

[image1]: images/image1.png

[image2]: images/image2.png

[image3]: images/image3.png