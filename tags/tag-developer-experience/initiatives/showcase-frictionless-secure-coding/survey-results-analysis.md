# Survey Results Analysis: Secure Coding Practices in CNCF Projects

Source: `tags/tag-developer-experience/initiatives/showcase-frictionless-secure-coding/survey-results-secure-coding-practices.csv`

> Note: source CSV in this directory is sanitized. Timestamps and direct contact details were removed or redacted. Project names and public technical references were preserved.

## Summary

This survey produced a **small but useful directional dataset**. The strongest signal is that **projects are more constrained by time, clarity, and workflow fit than by philosophical resistance to secure coding**. 

The survey has been communicated extensively in ~20 Slack channels of maintainers and contributors over the course of 8 weeks and **has not received many answers (13) compared to a similar survey TAG DevEx has done around AI within the CNCF that has received 200+ answers** from the same targeted personas. This divergence implies that AI disruptions are currently top-of-mind for CNCF project maintainers and contributors compared to security practices.

The data suggests five headline findings:

1. **Awareness of TAG Security & Compliance guidance is low.** Most respondents rated their familiarity at the low end of the scale.
2. **Projects often have security practices, but they are usually informal.** Only one response clearly indicated an established and documented standard.
3. **The main adoption barrier is maintainer bandwidth.** Time and contributor capacity dominated both structured and free-text answers.
4. **Low-friction wins are mostly automation and repository guardrails.** Dependabot, CodeQL, branch protections, audits, and signing were the clearest success stories.
5. **Friction is concentrated in adoption and remediation, not daily contribution flow.** Most respondents said secure coding practices had not blocked development, but they did report false positives, tooling complexity, and workflow mismatch.

## Methodology and data quality notes

- Raw submissions: **13**
- Responses included in analysis: **12**
- Excluded: **1 likely junk/accidental row** (`sdfsadf`) with only two populated cells
- Sample size is **small**, so findings should be treated as **directional**, not representative
- The sample is **heavily skewed toward the OpenTelemetry ecosystem**
- The CSV exports numeric scale values. These should be interpreted as **higher numbers are more positive / higher familiarity**
- One included response listed the project as `No`; it appears to be a valid response with an invalid project name, so it is treated as **unknown project** rather than dropped

## Respondent profile

### Project mix

Project names were normalized where obvious (for example, `Open Telemetry`, `OpenTelemetry`, `opentelemetry`, `opentelemetry-injector`, and the OpenTelemetry .NET repo links were grouped together).

| Normalized project family | Responses |
| --- | ---: |
| OpenTelemetry ecosystem | 6 |
| TAG DevEx | 1 |
| Kubeflow | 1 |
| Argo Rollouts | 1 |
| Kubescape | 1 |
| Prometheus | 1 |
| Unknown / invalid project name | 1 |

### Contributor roles

| Role answer | Responses |
| --- | ---: |
| Developer + DevOps / CI/CD / infrastructure | 10 |
| No role provided | 2 |

Interpretation: this survey mostly reflects the views of **hands-on contributors and maintainers working in both code and delivery pipelines**, not pure security specialists.

## Detailed findings

### 1) TAG Security & Compliance awareness is low

| Familiarity rating (1 to 5) | Responses |
| --- | ---: |
| 1 | 6 |
| 2 | 4 |
| 3 | 1 |
| 4 | 1 |

- **10/12 (83%)** rated familiarity **1 or 2**
- Average: **1.75 / 4**
- Median: **1.5 / 4**
- Even among low-familiarity respondents, **5/10** still reported at least informal standards

Takeaway: projects appear to be doing some security work, but often without linking it to TAG Security & Compliance guidance.

## 2) Standards mostly exist as informal practice

| Current standards status | Responses |
| --- | ---: |
| Yes, informal / implicit | 6 |
| Yes, established and documented | 1 |
| Not sure | 4 |
| No | 1 |

- **7/12** reported some form of standard
- Only **1/12** reported something clearly established and documented
- Detailed examples often described broader controls: **OWASP practices, `SECURITY.md`, RBAC, branch protection, artifact signing, 2FA**

Takeaway: respondents often interpreted **secure coding** broadly as overall project security hygiene, not only coding standards.

## 3) Security and delivery speed look compatible

| Balance rating | Responses |
| --- | ---: |
| 1 | 1 |
| 2 | 1 |
| 3 | 3 |
| 4 | 4 |

- Answered by **9/12** respondents
- Average: **3.11 / 4**
- Respondents with any standards averaged **3.43 / 4**
- Respondents with **`No`** or **`Not sure`** averaged **2.0 / 4**

Takeaway: some structure seems to reduce perceived tradeoff between security and velocity.

## 4) Biggest pain point is maintainer time

This question allowed multiple selections.

| Pain point | Mentions | Share of respondents who answered the question |
| --- | ---: | ---: |
| Lack of contributor time | 7 | 64% |
| Lack of clear best practices | 4 | 36% |
| Tooling complexity | 4 | 36% |
| High false positives in scans | 3 | 27% |
| Integration with project workflows | 3 | 27% |

Takeaway: friction is layered: **capacity first**, then **clarity**, **tooling quality**, and **workflow fit**.

## 5) Direct blocking is uncommon; noisy tools cause most pain

| Did secure coding slow or block work? | Responses |
| --- | ---: |
| No | 9 |
| Yes | 2 |

Only one respondent gave concrete detail:

> False positives in code scans

Takeaway: biggest pain is not day-to-day coding blockage. It is tool adoption, triage, remediation, and workflow integration.

## 6) Free-text blockers reinforce same pattern

Recurring themes:

- **Bandwidth:** repeated mentions of time pressure, competing priorities, unclear starting point
- **Tool trust:** false positives and blunt criticism of tool quality
- **Workflow mismatch:** incident response and GHSA/private-fix flows do not fit normal CI/review patterns
- **Awareness gap:** one respondent learned about TAG Security & Compliance only from survey itself

Takeaway: blockers are operational, not philosophical.

## 7) Success stories were automation-first

Only **4** respondents shared low-friction wins:

- branch/tag protections and GitHub Action audits
- **CodeQL**
- **Dependabot**
- automated dependency upgrades plus code scans

Takeaway: strongest pattern is **automation + guardrails**, not training programs or heavyweight review processes.

## 8) Advice was pragmatic

Only **3** respondents offered recommendations, but answers aligned:

- plan for security early
- use defense in depth
- keep controls maintainable

Takeaway: respondents want security built in early, but kept lightweight.

## 9) Interview follow-up pool is small

| Willing to participate in an interview | Responses |
| --- | ---: |
| Yes | 1 |
| Maybe | 3 |
| No | 7 |

Only one respondent provided a contact method, and it is redacted in sanitized data. Due to the low amount of interest for joining a working session to discuss it more, these will not be held.

## Cross-cutting insights

- **Practice maturity is ahead of TAG Security & Compliance guidance awareness.**
- **Low-friction security means automation by default.**
- **Real friction is operational overhead:** time, triage, integration, remediation.
- **Noise reduction matters as much as detection.** False positives are recurring DX problem.

## Recommendations for the initiative

Based on this survey, the initiative should prioritize:

1. **A maintainer-oriented quick-start baseline**
   - Start with a short, practical checklist
   - Emphasize what can be adopted in under an hour or a day

2. **Automation-first guidance**
   - Dependabot / dependency automation
   - CodeQL or equivalent scanning
   - Branch protection and release guardrails
   - Artifact signing where appropriate

3. **Guidance on reducing scanner friction**
   - False-positive tuning
   - Triage workflow examples
   - Advice on what to block vs what to warn on

4. **Workflow-native security patterns**
   - GHSA / private security fix workflows
   - CI-compatible incident response practices
   - Review patterns that work in open source projects

5. **Awareness-building for TAG Security & Compliance guidance**
   - The awareness gap is large enough that publishing guidance alone is unlikely to be sufficient
   - The distribution channels need to be rethought to ensure it reaches the intended audience

## Bottom line

This survey does not show strong opposition to secure coding. It shows a more practical problem: **maintainers need secure practices that are easy to discover, easy to enable, low-noise, and compatible with existing workflows without requiring large investment of contributor time**.

