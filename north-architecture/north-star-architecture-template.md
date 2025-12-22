# North Star Architecture — Reusable Template (Publication-Grade)

This document is a **technology-agnostic template** for defining a North Star Architecture: a small, explicit set of principles that guides distributed architectural decisions at scale.

It intentionally avoids prescribing specific tools or platforms. Instead, it encodes **decision constraints**, **accepted trade-offs**, and **review questions** that remain valid across technology cycles.

---

## 🚀 Quick Start (5 minutes)

Don't know where to begin? Follow this:

1. **Read the Purpose section** — fill in your 3 key values
2. **Choose 5 principles** from the table (you can modify later)
3. **List 3 trade-offs** you're already making implicitly
4. **Pick 1 anti-pattern** you see repeatedly

You now have a draft North Star. Use it in your next design review and iterate based on what worked.

**⚠️ Common mistake:** Teams copy this template word-for-word.

Your North Star must reflect YOUR context:
- Your technical constraints
- Your business model
- Your team's maturity
- Your actual architectural debt

A copied North Star is a dead North Star.

---

## How to use this template

You will get the best results if you treat the North Star as an **operational artifact**, not a one-time document:

- reference it in design docs and architecture reviews,
- use it to evaluate trade-offs during incidents and postmortems,
- revise it deliberately on a cadence (for example, every 6–12 months),
- keep it short enough that teams actually remember it.

If teams cannot quote it, it is too long.

---

## 1. Purpose

Describe why this North Star exists. Keep it to 1–3 sentences.

**Purpose (fill in):**
> _To guide distributed architectural decisions while preserving _____________________, _____________________, and _____________________, without sacrificing _____________________._

**Example:**
> To guide distributed architectural decisions while preserving scalability, resilience, and delivery speed, without sacrificing security and systemic coherence.

---

## 2. Core Architectural Principles

Define **5–8 durable, testable, and technology-agnostic principles**. A good principle is not a preference; it is a constraint that affects real decisions.

| Principle | Directive (what this means in practice) |
|---|---|
| Decoupling over local optimization | Prefer reducing cross-team dependencies even at an initial cost |
| Explicit ownership | Every service has a clearly accountable team responsible for build *and* run |
| Failure is expected | Systems must degrade gracefully; "perfect components" are not an assumption |
| Automation by default | Anything repeated becomes automated; production changes are codified |
| Observability is mandatory | Logs, metrics, and traces are part of "done" for production work |
| Security by default | Security and compliance requirements are built in, not gated at the end |

### How to choose YOUR principles

**Bad principle:** "Use microservices" (too specific, technology-dependent)  
**Good principle:** "Prefer loose coupling over local optimization" (durable, testable, survives tool changes)

**Bad principle:** "Be agile" (vague, not actionable)  
**Good principle:** "Optimize for reversibility in the first 18 months" (specific constraint, clear timeline)

Ask yourself: **Would this principle help a team choose between two valid designs?** If not, sharpen it.

### Notes for reviewers
- If a principle cannot be used to reject a design, it is probably too vague.
- If a principle forces the same implementation everywhere, it is probably too specific.

---

## 3. Explicitly Accepted Trade-offs

Architecture is not about avoiding trade-offs; it is about making them explicit.

**Trade-offs we accept (fill in):**
- We accept _____________________ to reduce _____________________.
- We accept _____________________ to improve _____________________.
- We accept _____________________ when _____________________.

**Examples:**
- We accept reduced local efficiency to minimize global coupling.
- We accept added latency to guarantee isolation between domains.
- We accept operational complexity when it materially increases resilience.

---

## 4. Decision-Guiding Questions (Review Checklist)

Use these questions in design reviews, PRDs/ADRs, and architecture proposals.

1. **Alignment:** Does this decision align with our North Star principles?
2. **Coupling:** Does it introduce hard-to-remove coupling (shared databases, hidden dependencies, coordination bottlenecks)?
3. **Resilience:** What happens when a dependency fails or becomes slow? Is degradation acceptable?
4. **Observability:** How will we detect, diagnose, and recover? Are signals defined (SLIs) and measurable?
5. **Security/Compliance:** What control objectives are affected? Are there auditable enforcement mechanisms?
6. **Reversibility:** Can we undo this decision at reasonable cost if assumptions change?
7. **Operational cost:** Who will operate it, and what is the steady-state burden?

If the answers are not clear, the proposal is not ready.

### Example walkthrough for Question 2 (Coupling)

**Proposal:** "Team A wants to read Team B's database directly for reporting."

**Assessment:**
- Does it introduce coupling? **YES** — schema changes in B break A
- Is it hard to remove? **YES** — becomes load-bearing over time
- Alternative? Expose reporting API or async events from B

**Decision:** Reject. Propose event-based approach with explicit contract.

---

## 5. Cloud-Native Guidelines (Optional)

Only include this section if you operate on cloud platforms or Kubernetes-like environments.

- Prefer **stateless services** where possible; make state explicit and externalized.
- Ensure infrastructure is **versioned and reproducible** (Infrastructure as Code).
- Prefer **horizontal scaling** and **graceful degradation** over vertical scaling.
- Make failure modes **observable** and test recovery paths.
- Automate delivery and rollbacks; keep changes small and frequent.

---

## 6. Explicitly Rejected Anti-Patterns

List practices that are unacceptable, even under delivery pressure.

**Anti-patterns we reject (fill in):**
- _____________________
- _____________________
- _____________________

**Examples:**
- Shared databases across independently owned services.
- Manual production configuration changes ("just this once").
- Critical dependencies without observability and defined SLOs.
- Security addressed only at the end of delivery.

---

## 7. Lightweight Governance

Define how this North Star is used and evolved—without creating a gatekeeping process.

**Governance model (fill in):**
- Referenced in: _____________________ (e.g., ADRs, design docs, postmortems)
- Reviewed by: _____________________ (e.g., architecture guild, staff engineers, platform team)
- Review cadence: _____________________ (e.g., every 6–12 months)
- Change process: _____________________ (e.g., RFC with rationale and migration notes)

The North Star should **enable autonomy**, not restrict it.

---

## 8. Success Signals

Define what "working" looks like.

**We expect to see:**

**Within 3 months:**
- Teams reference North Star in 50%+ of design docs
- Fewer "what should we do here?" escalations

**Within 6 months:**
- Measurable reduction in architectural rework
- Improved MTTR in incidents (teams understand system invariants)

**Within 12 months:**
- New engineers can make aligned decisions autonomously
- Platform choices converge naturally without mandates

If none of these improve over time, the North Star is either ignored or too vague.

---

## 9. Choose Your Starting Mode

Adapt the template to your organization's maturity:

### 🌱 Early Stage (0-50 people)
**Focus on:** Reversibility, Simplicity, Learning  
**Use:** 3-4 principles, loose governance  
**Key principle:** "Optimize for learning and change over premature optimization"

### 🌿 Growth Stage (50-200 people)
**Focus on:** Decoupling, Ownership, Observability  
**Use:** 5-6 principles, lightweight reviews  
**Key principle:** "Reduce coordination needs through clear boundaries"

### 🌳 Scale Stage (200+ people)
**Focus on:** Standardization, Resilience, Compliance  
**Use:** 7-8 principles, structured governance  
**Key principle:** "Automate enforcement of non-negotiable constraints"

---

## Summary (Short Version)

> Our North Star Architecture guides distributed decisions through clear, durable, technology-agnostic principles that prioritize decoupling, automation, resilience, observability, and security.

---

## Final note

The star does not build systems.  
But without it, every team follows a different constellation.