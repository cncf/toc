# KServe Adopter Interview – Nutanix

**Interview Date:** July 3, 2025  
**Interviewee:** Johnu George

---

## About the Adopter Interview

The purpose of this interview is to assess the maturity and adoption of the KServe project by its users. Notes are recorded by the CNCF Technical Oversight Committee (TOC), refined for clarity, and shared with the adopter for review and approval. Approved notes are included in KServe’s due diligence documentation.

---

## Organization Introduction

### Please provide an overview of your organization and its business focus.  

**A:** Nutanix is a global leader in cloud software, providing organizations with a unified platform for running applications and managing data across multiple clouds. The Nutanix platform is designed to reduce operational complexity, streamline infrastructure management, and allow customers to focus on achieving their business outcomes.

---

## Motivation for Adopting KServe

### Compared with other solutions in the market, what motivated your organization to choose KServe?  

**A:** Nutanix contributors have been engaged with KServe since its early days as a subproject under Kubeflow. We selected KServe over alternative proprietary and open-source solutions due to its flexibility, production-ready features, and Kubernetes-native architecture. KServe supports both predictive and generative models, provides standard inference APIs that abstract the complexity of underlying ML frameworks, and offers advanced capabilities such as accelerator support and autoscaling—features that are highly aligned with our enterprise requirements.

---

## Usage Scenarios

### How long has your organization been using KServe?  

**A:** Nutanix has been using KServe in production for several years, with active contributions to the project dating back to its inception.

### What are the primary use cases and features your organization relies on?  

**A:**  
- Standard inference APIs supporting TensorFlow, PyTorch, and Hugging Face  
- Autoscaling (scale up/down and scale-to-zero)  
- Tight integration with Kubernetes and the broader CNCF ecosystem  
- Scalable controller architecture enabling custom interface plugins  

Our deployments include both predictive and generative AI use cases, such as large language models (LLMs), code review automation, and code generation tools.

### What version are you currently using, and how do you manage upgrades?  

**A:** We are currently using version 0.15. Each new KServe release is evaluated internally over a one- to two-month period before adoption. KServe’s strong commitment to backward compatibility has resulted in consistently smooth upgrade experiences.

### How was the adoption and integration process? Were there any challenges?  

Adoption challenges were minimal due to the quality of KServe’s documentation and our active participation in the community. With three official reviewers and five to six active contributors from Nutanix, we are well-positioned to identify and address potential compatibility issues early.

### Was the project documentation valuable to your implementation?  

**A:** Yes. KServe’s documentation is comprehensive and includes extensive examples, which have been critical for smooth adoption. That said, more examples focused on generative AI would be valuable, as much of the existing content is centered on predictive AI.

### Has the adoption of KServe delivered measurable value to your organization?  

**A:** Yes. KServe’s autoscaling capabilities have significantly reduced manual operational overhead, lowering complexity and improving efficiency.

### What are your future plans for involvement with the project?  

**A:** Nutanix will continue to actively contribute to KServe, with plans to increase contributions over time. We also intend to deepen integration with related CNCF ecosystem projects, such as Envoy, and expand our participation in the broader model inference community.

---

## Perception of the Project

### How do you assess KServe’s community, governance, and growth potential?  

**A:**  
- **Community:** KServe has a vibrant and collaborative community, with healthy contributor diversity.  
- **Governance:** Strong, neutral governance processes are in place, including biannual contributor activity reviews.  
- **Growth Potential:** Significant. We believe that with increased outreach and marketing, KServe’s unique strengths could be more widely recognized in the AI/ML ecosystem.

### How does your organization participate in the community?  

**A:** We regularly attend biweekly community meetings, actively engage on CNCF Slack channels, and encourage all internal contributors to participate in upstream development.

### Have you engaged with maintainers or other community members?  

**A:** Yes. We work closely with maintainers through meetings and Slack-based asynchronous communication, which has consistently yielded productive outcomes.

---

## Project Strengths

- Active, diverse, and collaborative contributor community  
- Strong industry adoption (e.g., Bloomberg has long deployed KServe in production)  
- Transparent and effective governance  
- Proven track record of delivering backward-compatible, production-ready releases  

---

## Opportunities for Improvement

### What could help KServe reach its full potential?  

**A:** Greater marketing and outreach to highlight KServe’s strengths, particularly in the generative AI space, would drive broader adoption. CNCF affiliation can help amplify visibility and increase community engagement.

### What other improvements would you suggest?  

**A:**  
- Expand generative AI documentation and examples  
- Continue emphasizing timely PR reviews in community meetings to avoid bottlenecks  

---

**Reference:** [KServe OWNERS file](https://github.com/kserve/kserve/blob/master/OWNERS)
