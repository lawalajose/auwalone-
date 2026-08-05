# ADR-0001: Adopt a Monorepo for Platform Development

- **Status:** Accepted
- **Date:** 2026-08-05

## Context

auwalone+ is designed as a platform composed of multiple independently deployable
microservices, including Shipment, Fleet, Tracking, Billing, Notifications, and
Integrations. The platform also includes multiple client applications, shared
libraries, infrastructure-as-code, and architectural documentation.

Although the runtime architecture follows a microservices model, the engineering
team must decide how the source code is organized.

The options considered were:

- **Monorepo** — all applications, services, shared packages, infrastructure, and
  documentation are maintained in a single Git repository.
- **Polyrepo** — each application or service is maintained in its own repository.

This decision concerns **repository organization**, not service architecture.
Services remain independently deployable regardless of repository structure.

---

## Decision

We will maintain the entire auwalone+ platform in a **single monorepo**.

The repository will contain:

- Applications
- Backend services
- Shared packages
- Infrastructure definitions
- Architecture documentation

Each service will continue to have its own deployment pipeline, datastore,
runtime, and ownership boundaries.

---

## Rationale

### 1. Cross-service changes are common

Many features span multiple services and shared packages.

For example, introducing a new shipment field may require updates to:

- Shipment Service
- Tracking Service
- Billing Service
- Shared Types
- API Gateway
- Web Dashboard
- Documentation

A monorepo allows these changes to be implemented, reviewed, tested, and merged
as a single logical change.

---

### 2. Shared libraries evolve with the platform

Packages such as authentication, shared types, logging, and configuration are
consumed by multiple services.

Keeping shared packages in the same repository allows them to evolve alongside
their consumers without introducing version synchronization or internal package
publishing overhead.

---

### 3. Consistent engineering standards

A single repository enables one set of engineering standards across the platform,
including:

- Code formatting
- Linting
- Testing
- Dependency management
- Security scanning
- Code ownership
- CI/CD configuration

Consistency reduces operational complexity and improves maintainability.

---

### 4. Easier refactoring

Platform-wide refactoring becomes significantly simpler.

Changes to event schemas, authentication claims, shared models, or public APIs
can be propagated across all affected services within a single pull request,
allowing reviewers to see the complete impact before merging.

---

### 5. Better developer experience

Developers clone one repository and gain access to the complete platform.

This simplifies:

- Onboarding
- Local development
- Code discovery
- Debugging
- Architectural understanding

A single workspace also improves IDE support for navigation and refactoring.

---

### 6. Infrastructure evolves with the application

Infrastructure definitions, deployment manifests, monitoring configuration, and
architecture documentation are versioned alongside application code.

Operational changes therefore evolve together with the services they support,
improving traceability and reducing configuration drift.

---

### 7. Microservices do not require multiple repositories

The platform follows a microservices architecture, but repository structure is an
independent concern.

Each service continues to:

- Own its own datastore
- Expose its own APIs and events
- Be deployed independently
- Scale independently
- Maintain clear domain ownership

The monorepo affects only how source code is organized, not how services are
built, deployed, or operated.

---

## Consequences

### Positive

- Atomic cross-service changes
- Simplified shared library management
- Consistent tooling and governance
- Easier large-scale refactoring
- Improved developer onboarding
- Better discoverability across the platform
- Infrastructure and documentation remain versioned with application code

### Negative

- Repository size will increase over time.
- CI/CD requires incremental builds, caching, and affected-project detection.
- Ownership boundaries must be enforced through code ownership and architectural
  conventions.
- Teams must avoid introducing unnecessary coupling simply because services
  reside in the same repository.

---

## Alternatives Considered

### Polyrepo

Maintaining one repository per service provides strong repository-level isolation
and allows independent repository lifecycle management.

However, for auwalone+, this approach introduces significant coordination costs:

- Cross-service changes require multiple pull requests.
- Shared libraries require publishing and version management.
- CI/CD configuration is duplicated.
- Documentation becomes fragmented.
- Discoverability across the platform is reduced.

Given the expected frequency of changes spanning multiple business domains, these
trade-offs outweigh the benefits.

---

## Decision Summary

auwalone+ adopts a **monorepo** because it improves developer productivity,
simplifies cross-service evolution, and provides a consistent engineering
experience across the platform.

This decision is independent of the platform's microservices architecture.
Services remain independently deployable, independently scalable, and
independently owned while sharing a common source repository.
