# Contributing to Auwalone+

Thank you for your interest in contributing to **Auwalone+** — the enterprise-grade logistics
platform that powers shipment management, fleet operations, and real-time delivery tracking
for logistics companies around the world.

This document is the single source of truth for how we build, review, test, and ship software
together. It is intentionally long because Auwalone+ is a production system that moves real
freight for real customers — but it is organized so you only need to read the sections relevant
to what you're working on today.

[![Build Status](https://img.shields.io/github/actions/workflow/status/auwalone-plus/platform/ci.yml?branch=main)](https://github.com/auwalone-plus/platform/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/auwalone-plus/platform)](https://goreportcard.com/report/github.com/auwalone-plus/platform)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
[![Good First Issues](https://img.shields.io/github/issues/auwalone-plus/platform/good%20first%20issue)](https://github.com/auwalone-plus/platform/labels/good%20first%20issue)

> [!NOTE]
> **New here?** Skip straight to [Getting Started](#5-getting-started) to get a local
> environment running in under 15 minutes, then look for issues labeled
> [`good first issue`](https://github.com/auwalone-plus/platform/labels/good%20first%20issue).

---

## Table of Contents

1. [Welcome](#1-welcome)
2. [Code of Conduct](#2-code-of-conduct)
3. [Project Philosophy & Contribution Principles](#3-project-philosophy--contribution-principles)
4. [Ways to Contribute](#4-ways-to-contribute)
5. [Getting Started](#5-getting-started)
   - [5.1 Prerequisites](#51-prerequisites)
   - [5.2 Development Environment Setup](#52-development-environment-setup)
   - [5.3 Installation](#53-installation)
   - [5.4 Running the Project Locally](#54-running-the-project-locally)
   - [5.5 Environment Variables](#55-environment-variables)
6. [Project Structure Overview](#6-project-structure-overview)
7. [Git Workflow & Branching Strategy](#7-git-workflow--branching-strategy)
8. [Commit Message Convention](#8-commit-message-convention)
9. [Pull Request Process](#9-pull-request-process)
10. [Coding Standards](#10-coding-standards)
11. [Naming Conventions](#11-naming-conventions)
12. [Formatting & Linting](#12-formatting--linting)
13. [Testing Requirements](#13-testing-requirements)
14. [API Development Guidelines](#14-api-development-guidelines)
15. [Database Migration Guidelines](#15-database-migration-guidelines)
16. [Documentation Requirements](#16-documentation-requirements)
17. [Security Best Practices](#17-security-best-practices)
18. [Performance Expectations](#18-performance-expectations)
19. [Logging & Observability Standards](#19-logging--observability-standards)
20. [WebSocket Development Guidelines](#20-websocket-development-guidelines)
21. [Message Queue Development Guidelines](#21-message-queue-development-guidelines)
22. [Webhook Implementation Guidelines](#22-webhook-implementation-guidelines)
23. [Feature Flag Policy](#23-feature-flag-policy)
24. [Versioning Strategy](#24-versioning-strategy)
25. [Dependency Management](#25-dependency-management)
26. [CI/CD Expectations](#26-cicd-expectations)
27. [Pull Request Review Checklist](#27-pull-request-review-checklist)
28. [Reporting Bugs](#28-reporting-bugs)
29. [Requesting Features](#29-requesting-features)
30. [Bug Report Template](#30-bug-report-template)
31. [Pull Request Template](#31-pull-request-template)
32. [Architecture Decision Records (ADRs)](#32-architecture-decision-records-adrs)
33. [Release Process](#33-release-process)
34. [Community & Communication Channels](#34-community--communication-channels)
35. [Recognition for Contributors](#35-recognition-for-contributors)
36. [License & Legal Considerations](#36-license--legal-considerations)

---

## 1. Welcome

Auwalone+ is a B2B SaaS platform that logistics companies rely on to manage shipments, vehicle
fleets, drivers, customers, and real-time delivery operations at scale. The platform spans:

| Domain | Description |
|---|---|
| Customer Account Management | Organizations, users, roles, and multi-tenant workspaces |
| Authentication & Authorization | OIDC/JWT auth, RBAC/ABAC, service-to-service identity |
| Billing & Subscriptions | Usage metering, invoicing, plan management |
| Orders & Shipment Lifecycle | Core shipment state machine and public/partner APIs |
| Fleet & Driver Management | Vehicles, drivers, capacity planning, compliance |
| Real-Time GPS Tracking | WebSocket-based live location and ETA updates |
| Driver Notifications | Asynchronous messaging via Kafka |
| Carrier Integrations | Inbound/outbound webhooks with third-party carriers |
| Admin Dashboard | Internal operations and support tooling |
| Reporting & Analytics | Operational and business intelligence |
| Multi-Region Deployment | Active-active and active-passive regional topologies |
| Disaster Recovery | RPO/RTO-driven backup and failover strategy |
| Chaos Engineering | Proactive resilience validation |
| High Availability & Scalability | Horizontal scaling, autoscaling, load shedding |

Whether you're fixing a typo in the docs, shipping a new API endpoint, hardening a Terraform
module, or designing a chaos experiment, **your contribution matters** and this guide will help
you land it successfully.

---

## 2. Code of Conduct

This project and everyone participating in it is governed by the
**[Auwalone+ Code of Conduct](CODE_OF_CONDUCT.md)**, based on the
[Contributor Covenant v2.1](https://www.contributor-covenant.org/version/2/1/code_of_conduct/).

By participating, you agree to uphold this code. Unacceptable behavior may be reported
confidentially to **conduct@auwalone.dev**.

> [!IMPORTANT]
> We are committed to providing a welcoming, harassment-free experience for everyone, regardless
> of experience level, gender identity, sexual orientation, disability, personal appearance,
> body size, race, ethnicity, age, religion, or nationality.

---

## 3. Project Philosophy & Contribution Principles

Auwalone+ moves real freight, bills real money, and coordinates real drivers. Our engineering
culture reflects that responsibility:

| Principle | What it means in practice |
|---|---|
| **Correctness over speed** | A late shipment is bad; an *incorrectly reported* shipment is worse. We favor provably correct code over quick hacks. |
| **Security by default** | Every feature is threat-modeled. Secure-by-default beats "secure it later." |
| **Observability is not optional** | If it isn't logged, metered, and traced, it doesn't exist in production. |
| **Small, reviewable changes** | Large PRs hide bugs. We prefer several small PRs over one sprawling one. |
| **Automate everything** | Linting, testing, security scanning, and deployment are automated — humans review intent, not mechanics. |
| **Documentation as code** | Docs live next to code, are versioned, and are reviewed like code. |
| **Blameless culture** | Incidents and chaos experiments exist to find systemic issues, never to assign blame. |
| **Backward compatibility matters** | We serve external partners and carriers — breaking changes require a deprecation path. |
| **Progressive delivery** | Feature flags, canaries, and staged rollouts are the default deployment pattern, not the exception. |

> [!TIP]
> When in doubt, optimize for the reviewer and the on-call engineer six months from now — not
> for how fast you can merge today.

---

## 4. Ways to Contribute

You don't have to write Go code to contribute meaningfully:

- 🐛 **Bug fixes** — see [Reporting Bugs](#28-reporting-bugs) and issues labeled `bug`
- ✨ **Features** — see [Requesting Features](#29-requesting-features)
- 🧪 **Tests** — increasing coverage on under-tested packages is always welcome
- 📝 **Documentation** — READMEs, ADRs, runbooks, API reference, tutorials
- 🏗️ **Infrastructure** — Terraform modules, Helm charts, ArgoCD applications
- 🔭 **Observability** — Grafana dashboards, alert rules, tracing improvements
- 🧭 **Issue triage** — labeling, reproducing, and reviewing incoming issues
- 🌍 **Community support** — answering questions in Discussions and Slack
- 🎨 **Design/UX** — admin dashboard usability and accessibility improvements

Look for [`good first issue`](https://github.com/auwalone-plus/platform/labels/good%20first%20issue)
and [`help wanted`](https://github.com/auwalone-plus/platform/labels/help%20wanted) labels to find
a place to start.

---

## 5. Getting Started

### 5.1 Prerequisites

| Tool | Minimum Version | Purpose |
|---|---|---|
| [Go](https://go.dev/dl/) | 1.22+ | Primary backend language |
| [Docker](https://docs.docker.com/get-docker/) | 24.x | Local containers |
| [Docker Compose](https://docs.docker.com/compose/) | v2 | Local dev stack |
| [Git](https://git-scm.com/) | 2.40+ | Version control |
| [Make](https://www.gnu.org/software/make/) | any recent | Task runner |
| [Northline CLI](#northline-cli) | 2.x | Auwalone+ developer platform CLI |
| [pre-commit](https://pre-commit.com/) | 3.x | Local git hooks |
| [kubectl](https://kubernetes.io/docs/tasks/tools/) | 1.29+ | Kubernetes contributors only |
| [Helm](https://helm.sh/docs/intro/install/) | 3.14+ | Chart contributors only |
| [Terraform](https://developer.hashicorp.com/terraform/install) | 1.7+ | Infra contributors only |
| [AWS CLI](https://docs.aws.amazon.com/cli/) | v2 | Infra contributors only |
| [buf](https://buf.build/docs/installation) | 1.30+ | Protobuf/gRPC contributors |

#### Northline CLI

**Northline** is Auwalone+'s internal developer platform CLI. It wraps Docker Compose, database
migrations, code generation, and environment diagnostics behind a single, opinionated interface
so that "getting productive" doesn't require memorizing 20 different tools.

```bash
# macOS / Linux
curl -sSL https://get.auwalone.dev/northline | bash

# Verify installation
northline version
```

### 5.2 Development Environment Setup

```bash
# 1. Fork the repository on GitHub, then clone your fork
git clone https://github.com/<your-username>/platform.git auwalone-plus
cd auwalone-plus

# 2. Add the upstream remote
git remote add upstream https://github.com/auwalone-plus/platform.git

# 3. Install git hooks (formatting, linting, DCO sign-off check, secret scanning)
pre-commit install
pre-commit install --hook-type commit-msg

# 4. Copy the example environment file
cp .env.example .env

# 5. Run environment diagnostics
northline doctor
```

`northline doctor` verifies your Go version, Docker daemon, port availability, and required
environment variables, and prints actionable fixes for anything missing.

### 5.3 Installation

```bash
# Download Go module dependencies
go mod download

# Pull and build local service images
northline dev pull
northline dev build
```

### 5.4 Running the Project Locally

```bash
# Start the full local stack: Postgres, Redis, Kafka, Jaeger, Prometheus, Grafana
northline dev up

# Apply database migrations
northline migrate up

# Seed representative sample data (organizations, drivers, vehicles, sample shipments)
northline seed

# Run all services with hot reload
northline dev run --all

# ...or run a single service
northline dev run orders-api
```

| Service | Protocol | Local Port | Notes |
|---|---|---|---|
| Orders API | HTTP/REST | `8080` | Public shipment lifecycle API |
| Admin API | HTTP/REST | `8082` | Internal dashboard backend |
| Billing API | HTTP/REST | `8083` | Subscriptions & invoicing |
| Fleet gRPC Service | gRPC | `9090` | Internal service-to-service |
| Tracking Gateway | WebSocket | `8081` | Live GPS / ETA streaming |
| PostgreSQL | TCP | `5432` | Primary data store |
| Redis | TCP | `6379` | Cache & pub/sub |
| Kafka | TCP | `9092` | Event streaming / notifications |
| Prometheus | HTTP | `9091` | Metrics |
| Grafana | HTTP | `3000` | Dashboards (admin/admin locally) |
| OpenTelemetry Collector | gRPC/HTTP | `4317` / `4318` | Traces |

Once running, visit `http://localhost:8080/healthz` to confirm the API is healthy, and
`http://localhost:3000` for local Grafana dashboards.

```bash
# Tear down the local stack
northline dev down

# Tear down and wipe volumes (fresh start)
northline dev down --volumes
```

### 5.5 Environment Variables

All services read configuration from environment variables (12-factor style). `.env.example`
documents every variable; the most commonly used ones are summarized below.

| Variable | Required | Example | Description |
|---|---|---|---|
| `APP_ENV` | Yes | `local` | `local`, `staging`, `production` |
| `HTTP_PORT` | Yes | `8080` | REST API listen port |
| `GRPC_PORT` | No | `9090` | gRPC listen port |
| `WS_PORT` | No | `8081` | WebSocket gateway port |
| `DATABASE_URL` | Yes | `postgres://user:pass@localhost:5432/auwalone` | Postgres DSN |
| `DATABASE_MAX_CONNS` | No | `25` | Connection pool size |
| `REDIS_URL` | Yes | `redis://localhost:6379/0` | Redis connection string |
| `KAFKA_BROKERS` | Yes | `localhost:9092` | Comma-separated broker list |
| `JWT_PUBLIC_KEY` | Yes | *(PEM string)* | Public key for verifying access tokens |
| `VAULT_ADDR` | Staging/Prod | `https://vault.internal:8200` | HashiCorp Vault address |
| `VAULT_ROLE` | Staging/Prod | `orders-api` | Vault AppRole/K8s auth role |
| `WEBHOOK_SIGNING_SECRET` | Yes | *(secret)* | HMAC key for outbound webhook signing |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | No | `http://localhost:4318` | Trace/metric export target |
| `LOG_LEVEL` | No | `info` | `debug`, `info`, `warn`, `error` |
| `FEATURE_FLAG_PROVIDER` | No | `local` | `local`, `unleash` |

> [!WARNING]
> **Never commit real secrets.** `.env` is git-ignored. In staging and production, all secrets
> are sourced from **HashiCorp Vault** — see [Security Best Practices](#17-security-best-practices).

---

## 6. Project Structure Overview

Auwalone+ is organized as a Go monorepo, with each domain following **Clean Architecture**
(domain → use case → adapter → infrastructure), deployed as independent Kubernetes workloads.

```text
auwalone-plus/
├── cmd/                        # Entry points — one main package per deployable binary
│   ├── orders-api/
│   ├── billing-api/
│   ├── fleet-api/
│   ├── tracking-gateway/       # WebSocket gateway
│   ├── notifications-worker/   # Kafka consumer
│   └── admin-api/
├── internal/
│   ├── orders/
│   │   ├── domain/             # Entities, value objects, domain services (no dependencies)
│   │   ├── usecase/            # Application/business logic, ports (interfaces)
│   │   ├── adapter/
│   │   │   ├── http/           # REST handlers (implements usecase ports)
│   │   │   ├── grpc/           # gRPC handlers
│   │   │   └── repository/     # Postgres/Redis implementations of repository ports
│   │   └── infrastructure/     # Wiring, config, external clients
│   ├── billing/
│   ├── fleet/
│   ├── tracking/
│   ├── auth/
│   └── shared/                 # Cross-cutting: logging, tracing, middleware, errors
├── pkg/                        # Reusable libraries safe for external import
├── api/
│   ├── openapi/                # REST API specs (v1/, v2/) — spec-first design
│   └── proto/                  # gRPC / Protobuf definitions (managed with buf)
├── migrations/                 # SQL migrations, one directory per bounded context
├── deployments/
│   ├── helm/                   # Helm charts per service
│   └── k8s/                    # Raw manifests (rare; prefer Helm)
├── infrastructure/
│   └── terraform/
│       ├── modules/            # Reusable Terraform modules
│       └── environments/       # dev/ staging/ production/ per-region roots
├── scripts/                    # Dev & CI utility scripts
├── test/
│   ├── integration/            # testcontainers-go based integration tests
│   └── e2e/                    # End-to-end suites against a running environment
├── docs/
│   ├── adr/                    # Architecture Decision Records
│   └── runbooks/                # Operational runbooks
├── .github/
│   ├── workflows/               # GitHub Actions CI/CD pipelines
│   ├── ISSUE_TEMPLATE/
│   └── PULL_REQUEST_TEMPLATE.md
├── docker-compose.yml
├── Makefile
├── go.mod
└── CONTRIBUTING.md
```

> [!TIP]
> New to Clean Architecture here? Dependencies always point **inward**:
> `adapter → usecase → domain`. The `domain` package never imports anything from `adapter` or
> `infrastructure`.

---

## 7. Git Workflow & Branching Strategy

Auwalone+ uses a **trunk-based development** model with short-lived branches.

| Branch | Purpose |
|---|---|
| `main` | Always deployable. Protected. All work merges here via PR. |
| `release/x.y` | Cut from `main` for a release; only receives cherry-picked fixes. |
| `hotfix/*` | Urgent production fixes, branched from the affected `release/x.y`. |
| `feature/*`, `fix/*`, `chore/*`, `docs/*`, `refactor/*`, `test/*`, `perf/*`, `ci/*` | Short-lived contributor branches. |

**Branch naming:** `<type>/<ticket-id>-<short-kebab-description>`

```text
feature/AUW-1423-shipment-eta-recalculation
fix/AUW-1590-webhook-signature-mismatch
docs/AUW-1601-update-adr-index
chore/AUW-1622-bump-go-1-23
```

### Workflow

```bash
# 1. Sync with upstream before starting work
git checkout main
git pull upstream main

# 2. Create your branch
git checkout -b feature/AUW-1423-shipment-eta-recalculation

# 3. Commit using Conventional Commits (see below) and sign off (DCO)
git commit -s -m "feat(orders): recalculate ETA on route deviation"

# 4. Keep your branch current via rebase, not merge
git fetch upstream
git rebase upstream/main

# 5. Push and open a Pull Request
git push origin feature/AUW-1423-shipment-eta-recalculation
```

**Rules:**

- ✅ No direct pushes to `main` — it is a protected branch requiring PR + passing checks.
- ✅ Rebase your branch on `main` before requesting review; resolve conflicts locally.
- ✅ We use **squash merge** — your PR becomes a single, clean commit on `main`.
- ❌ Do not merge `main` into your feature branch — rebase instead, to keep history linear.
- ❌ Force-push (`--force-with-lease`) is fine on your own feature branch, never on shared branches.

---

## 8. Commit Message Convention

We follow **[Conventional Commits](https://www.conventionalcommits.org/)**, enforced via
`commitlint` in our pre-commit hook and CI.

```text
<type>(<scope>): <short summary>

[optional body]

[optional footer(s)]
```

| Type | When to use it |
|---|---|
| `feat` | A new feature |
| `fix` | A bug fix |
| `docs` | Documentation-only changes |
| `style` | Formatting, whitespace — no logic change |
| `refactor` | Code change that neither fixes a bug nor adds a feature |
| `perf` | Performance improvement |
| `test` | Adding or correcting tests |
| `build` | Build system or external dependency changes |
| `ci` | CI/CD pipeline changes |
| `chore` | Maintenance tasks not covered above |
| `revert` | Reverts a previous commit |

**Common scopes:** `orders`, `billing`, `fleet`, `tracking`, `auth`, `webhooks`, `notifications`,
`admin`, `helm`, `terraform`, `ci`.

```text
✅ feat(tracking): add reconnect support to WebSocket gateway
✅ fix(billing): correct proration for mid-cycle plan upgrades
✅ refactor(orders): extract shipment state machine into domain service

❌ fixed bug
❌ WIP
❌ updates
```

**Breaking changes** must include a footer:

```text
feat(api): remove deprecated /v1/shipments/legacy endpoint

BREAKING CHANGE: /v1/shipments/legacy is removed. Use /v2/shipments instead.
Closes #1874
```

> [!IMPORTANT]
> Every commit must be **signed off** with `git commit -s` (Developer Certificate of Origin).
> See [License & Legal Considerations](#36-license--legal-considerations).

---

## 9. Pull Request Process

1. **Start from an issue.** Every PR should reference an existing issue (`Closes #1234`). Open one
   first if none exists — this avoids duplicate or unwanted work.
2. **Branch and implement** following [coding standards](#10-coding-standards), with tests and
   docs updated in the same PR.
3. **Run local verification** before pushing:
   ```bash
   make check   # fmt + lint + unit tests + vulnerability scan
   ```
4. **Open the PR** against `main` using the [PR template](#31-pull-request-template). Mark it
   **Draft** if it's a work in progress.
5. **Keep it small.** We target PRs under ~400 changed lines (excluding generated code and
   lockfiles). Larger changes should be split or preceded by an [ADR](#32-architecture-decision-records-adrs).
6. **Ensure CI is green** — all required checks in [CI/CD Expectations](#26-cicd-expectations)
   must pass before review.
7. **Request review** — CODEOWNERS are auto-assigned. Core domains (`orders`, `billing`, `auth`,
   `fleet`) require **2 approvals**, including one from the domain's designated owner. All other
   changes require **1 approval**.
8. **Address feedback** via new commits (don't force-push during active review — it hides diffs).
   Once approved, feel free to rebase and force-push before merge.
9. **Squash merge.** The PR title becomes the squashed commit message and must itself follow
   Conventional Commits.
10. **Delete your branch** after merge (GitHub does this automatically for forks with the setting
    enabled).

> [!TIP]
> Stuck on review for more than a few days? Ping the reviewer in the PR thread or in
> `#auwalone-contributors` on Slack — don't hesitate to ask for status.

---

## 10. Coding Standards

- Write **idiomatic Go** per [Effective Go](https://go.dev/doc/effective_go) and the
  [Google Go Style Guide](https://google.github.io/styleguide/go/).
- Every exported function, type, and package has a doc comment (`godoc`-style).
- **Errors are values.** Wrap with context using `fmt.Errorf("doing X: %w", err)`; never discard
  errors silently. Define sentinel/domain errors (`var ErrShipmentNotFound = errors.New(...)`) in
  the `domain` package.
- **`context.Context` is always the first parameter** of any function that can block, call the
  network, or touch a database. Never store a `Context` in a struct field.
- **Dependency injection via constructors** (`NewOrderService(repo Repository, clock Clock) *OrderService`).
  We do not use a DI framework; wiring happens explicitly in `cmd/*/main.go`.
- **Interfaces are defined by the consumer**, not the implementer (e.g., `usecase` defines the
  `Repository` interface it needs; `adapter/repository` implements it).
- Avoid global mutable state and `init()` side effects beyond simple registration.
- Do not `panic` in library or business logic code — panics are reserved for truly
  unrecoverable programmer errors (and are always recovered at the transport boundary).
- Prefer composition over inheritance-like embedding tricks; keep structs small and focused.
- Concurrency: prefer well-scoped goroutines with explicit lifecycle management
  (`errgroup`, `context` cancellation) over ad-hoc `go func(){}()` with no supervision.

```go
// Example: usecase layer depends only on an interface it defines itself.
package usecase

type ShipmentRepository interface {
    GetByID(ctx context.Context, id ShipmentID) (*domain.Shipment, error)
    Save(ctx context.Context, s *domain.Shipment) error
}

type RecalculateETA struct {
    repo  ShipmentRepository
    clock Clock
}

func (uc *RecalculateETA) Execute(ctx context.Context, id domain.ShipmentID) error {
    shipment, err := uc.repo.GetByID(ctx, id)
    if err != nil {
        return fmt.Errorf("loading shipment %s: %w", id, err)
    }
    if err := shipment.RecalculateETA(uc.clock.Now()); err != nil {
        return fmt.Errorf("recalculating eta: %w", err)
    }
    return uc.repo.Save(ctx, shipment)
}
```

---

## 11. Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Go packages | short, lowercase, no underscores | `tracking`, `billing` |
| Go identifiers | `MixedCaps` / `mixedCaps`, never `snake_case` | `ShipmentID`, `recalculateETA` |
| Interfaces | noun or `-er` suffix describing behavior | `Repository`, `Notifier` |
| Test files | `<file>_test.go`, table-driven test names describe behavior | `TestRecalculateETA_DeviatesFromRoute` |
| Postgres tables | `snake_case`, plural | `shipments`, `driver_locations` |
| Postgres migrations | `NNNN_description.up.sql` / `.down.sql` | `0042_add_shipment_eta.up.sql` |
| REST resources | plural nouns, kebab-case paths | `/v1/shipment-events` |
| gRPC services/RPCs | `PascalCase` service, `PascalCase` RPC | `FleetService.AssignDriver` |
| Kafka topics | `<domain>.<entity>.<event>.v<N>` | `shipment.status.updated.v1` |
| Environment variables | `SCREAMING_SNAKE_CASE` | `DATABASE_URL` |
| Feature flags | `<domain>_<feature>_<qualifier>` | `tracking_multi_leg_eta_v2` |
| Helm releases | `<service>-<environment>` | `orders-api-staging` |

---

## 12. Formatting & Linting

All formatting and linting is enforced locally (pre-commit) and in CI — there should never be a
style debate in code review.

```bash
make fmt     # gofmt + goimports
make lint    # golangci-lint run (errcheck, govet, staticcheck, revive, gosec, unconvert...)
make sql-fmt # SQL migration formatting
terraform fmt -recursive
tflint
helm lint deployments/helm/*
```

Configuration lives in `.golangci.yml`, `.editorconfig`, and `.markdownlint.yml` at the repo root
— do not override these locally without discussion.

> [!TIP]
> Run `pre-commit run --all-files` before your first PR to catch formatting issues across the
> whole repo, not just files you touched.

---

## 13. Testing Requirements

We follow the testing pyramid:

| Layer | Tooling | Runs on | Scope |
|---|---|---|---|
| Unit | `go test`, `testify`, `mockery`-generated mocks | Every commit (pre-commit + CI) | Single function/type, no I/O |
| Integration | `testcontainers-go` (real Postgres/Redis/Kafka in Docker) | Every PR | Real dependencies, single service |
| Contract | OpenAPI schema validation, `buf breaking` | Every PR | API/schema compatibility |
| End-to-End | Go E2E suite against `northline dev up` stack | Nightly + pre-release | Full user journeys |
| Load | [k6](https://k6.io/) | Weekly + before major releases | Throughput/latency under load |
| Chaos | Chaos Mesh experiments | Scheduled game days | Resilience under failure |

**Requirements:**

- New code requires **≥ 80% statement coverage**, enforced via Codecov on the diff, not the
  whole repo.
- All tests are **table-driven** where they cover multiple input cases.
- Run the race detector locally for concurrency-sensitive code: `go test -race ./...`.
- Integration tests must clean up their own containers/state (`testcontainers-go` handles this
  automatically via `Terminate`).
- Flaky tests are treated as bugs — quarantine with `t.Skip()` and a linked issue, never leave
  them silently retried.

```go
func TestRecalculateETA(t *testing.T) {
    tests := []struct {
        name        string
        deviationKM float64
        wantErr     error
    }{
        {name: "no deviation", deviationKM: 0, wantErr: nil},
        {name: "minor deviation recalculates", deviationKM: 2.5, wantErr: nil},
        {name: "route lost returns error", deviationKM: -1, wantErr: domain.ErrRouteUnavailable},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // arrange, act, assert
        })
    }
}
```

---

## 14. API Development Guidelines

Auwalone+ is **spec-first**: the API contract is designed and reviewed *before* implementation.

- **REST:** Author/update the OpenAPI 3.1 spec under `api/openapi/v{n}/` first. Generate server
  interfaces with `oapi-codegen` (`make generate-api`). PRs that hand-write REST handlers without
  a corresponding spec change will be asked to add one.
- **gRPC:** Define/update Protobuf under `api/proto/`, lint and check for breaking changes with
  [`buf`](https://buf.build/):
  ```bash
  buf lint
  buf breaking --against '.git#branch=main'
  ```
- **Versioning:** Breaking changes require a new version prefix (`/v1` → `/v2`); the previous
  version is supported for a minimum deprecation window (see [Versioning Strategy](#24-versioning-strategy)).
- **Pagination:** Cursor-based (`?page_token=` / `next_page_token`), never offset-based, for any
  collection that can grow unbounded (shipments, events, driver locations).
- **Errors:** REST errors follow [RFC 7807 Problem Details](https://www.rfc-editor.org/rfc/rfc7807)
  (`application/problem+json`) with a stable `type` URI per error code.
- **Idempotency:** All mutating endpoints in `orders` and `billing` must accept an
  `Idempotency-Key` header and de-duplicate retried requests.
- **Backward compatibility:** Adding fields is safe; renaming, removing, or changing the type of
  a field is a breaking change requiring a version bump and an ADR.

---

## 15. Database Migration Guidelines

We use [`golang-migrate`](https://github.com/golang-migrate/migrate) for schema migrations and
[`sqlc`](https://sqlc.dev/) for type-safe generated query code.

```bash
# Create a new migration pair
northline migrate new add_shipment_eta_column --domain orders

# Apply migrations locally
northline migrate up

# Test migrations up AND down against an ephemeral database
northline migrate test
```

**Rules:**

- Every migration must have a working `.up.sql` **and** `.down.sql`.
- Follow the **expand–contract pattern** for anything that could break a running instance during
  a rolling deploy: add the new column/table first (expand), deploy code that writes to both,
  backfill, then remove the old column in a *later* migration (contract).
- Never edit a migration that has already been merged to `main` — write a new one.
- Large tables: use `CREATE INDEX CONCURRENTLY` and avoid table locks during peak hours.
- Schema changes to `orders`, `billing`, or `fleet` require review from that domain's
  designated data owner (see `CODEOWNERS`).
- Regenerate `sqlc` code (`make generate-db`) and commit the generated files alongside the
  migration.

> [!WARNING]
> A destructive migration (`DROP COLUMN`, `DROP TABLE`) merged without a deprecation window is
> treated as a production incident, not a code review miss. Always coordinate destructive schema
> changes with the on-call data owner.

---

## 16. Documentation Requirements

- Any PR that changes public behavior (API, CLI, configuration) must update the relevant docs in
  the same PR: `README.md`, `api/openapi/*`, or `docs/`.
- Exported Go identifiers require doc comments; `golangci-lint`'s `revive` rule enforces this for
  public packages.
- Architectural or cross-cutting docs (diagrams) use [Mermaid](https://mermaid.js.org/) fenced
  code blocks so they render natively on GitHub and stay diffable.
- New services require a **runbook** in `docs/runbooks/<service>.md` covering: what it does,
  dashboards, common alerts, and escalation path.
- Documentation changes are reviewed with the same rigor as code — clarity and accuracy are not
  optional.

---

## 17. Security Best Practices

- **No secrets in code, config files, or commit history.** All runtime secrets are sourced from
  **HashiCorp Vault** via dynamic short-lived credentials (database creds, API keys).
- **Parametrized queries only.** Never build SQL via string concatenation, even for
  "internal-only" tooling.
- **AuthN/AuthZ:** Authentication uses OIDC-issued JWTs; authorization is enforced centrally via
  RBAC/ABAC middleware — never re-implement access checks ad hoc inside handlers.
- **Least privilege:** Service accounts, IAM roles, and Vault policies are scoped to the minimum
  required actions and resources.
- **Dependency scanning:** `govulncheck` runs in CI on every PR; Dependabot/Renovate opens
  automated upgrade PRs for vulnerable dependencies.
- **Static analysis:** `gosec` is part of the mandatory `golangci-lint` pipeline.
- **Secret scanning:** `gitleaks` runs pre-commit and in CI to catch accidentally committed
  credentials before they reach `main`.
- **Transport security:** TLS is required for all service-to-service and external traffic;
  mutual TLS (mTLS) is used within the service mesh.
- **Webhook signatures:** All inbound/outbound webhook payloads are HMAC-signed — see
  [Webhook Implementation Guidelines](#22-webhook-implementation-guidelines).
- **Rate limiting** is applied at the API gateway for all public-facing endpoints.
- **Policy as Code:** Infrastructure and Kubernetes changes are validated against Open Policy
  Agent (OPA)/Conftest policies in CI before merge.

> [!IMPORTANT]
> **Do not open a public GitHub issue for a security vulnerability.** Report it privately to
> **security@auwalone.dev** per our [`SECURITY.md`](SECURITY.md) responsible disclosure policy.
> We aim to acknowledge reports within 48 hours.

---

## 18. Performance Expectations

| Metric | Target |
|---|---|
| Orders API — p99 latency (read) | < 150 ms |
| Orders API — p99 latency (write) | < 300 ms |
| Tracking Gateway — location update fan-out latency | < 500 ms end-to-end |
| Billing API — p99 latency | < 400 ms |
| Database query budget per request | ≤ 3 queries, no N+1 |
| Cache hit ratio (hot read paths) | > 90% |

- Profile before optimizing — attach `pprof`/benchmark evidence to PRs claiming a performance
  improvement.
- Use Redis for cache-aside on hot read paths with explicit TTLs; guard against cache stampede
  with request coalescing (`singleflight`) for expensive recomputation.
- All database access goes through pooled connections (`pgxpool` / PgBouncer in production); no
  ad hoc connections per request.
- Any endpoint expected to exceed **50 req/s** in production requires a `k6` load test in
  `test/load/` before merge.
- Kubernetes workloads must define CPU/memory `requests` and `limits` — unbounded resource usage
  will fail Helm chart review.
- Prefer streaming and pagination over large in-memory payloads (e.g., bulk shipment export).

---

## 19. Logging & Observability Standards

Every service must expose `/healthz`, `/readyz`, and `/metrics` (Prometheus format) endpoints.

- **Structured logging** via `zerolog`, JSON output, with mandatory fields: `timestamp`, `level`,
  `service`, `trace_id`, `span_id`, `request_id`. **Never log PII** (customer names, exact
  addresses, payment details) — use redaction helpers in `internal/shared/logging`.
- **Metrics** follow the **RED method** (Rate, Errors, Duration) for every request-handling
  component, and the **USE method** (Utilization, Saturation, Errors) for infrastructure.
  Dashboards are version-controlled as code under `deployments/observability/grafana/`.
- **Tracing** uses OpenTelemetry SDKs with automatic instrumentation for HTTP/gRPC/DB clients,
  plus manual spans around business-critical operations (e.g., `shipment.recalculate_eta`).
  Trace context propagates across HTTP, gRPC, and Kafka message headers.
- **Alerts as code:** Prometheus alerting rules live in `deployments/observability/prometheus/`
  and are reviewed like any other change — no manually-created alerts in the Grafana UI.
- **Log aggregation** is centralized in Loki; use `trace_id` to pivot between logs, metrics, and
  traces during an incident.

---

## 20. WebSocket Development Guidelines

The **Tracking Gateway** streams live GPS and ETA updates to dispatchers and customers.

- **Authentication:** Connections authenticate via a short-lived, single-use token exchanged
  during the WebSocket handshake — never long-lived credentials.
- **Heartbeats:** Server-initiated ping every 30s; a client missing 2 consecutive pongs is
  disconnected.
- **Message envelope:** All messages are versioned JSON:
  ```json
  { "type": "location.updated", "version": 1, "shipment_id": "shp_123", "payload": { "...": "..." } }
  ```
- **Reconnection:** Clients may resume a session with a `last_event_id`; the gateway replays
  missed events from a short-lived buffer (Redis Streams) where available.
- **Fan-out at scale:** Gateway replicas are stateless; broadcast fan-out across replicas uses
  Redis Pub/Sub so a client can be connected to any replica behind the load balancer.
- **Backpressure:** Slow consumers are subject to a bounded send buffer; when exceeded, the
  connection is closed with a clear close code rather than allowing unbounded memory growth.
- **Graceful shutdown:** On deploy, gateways stop accepting new connections, notify clients to
  reconnect elsewhere, and drain existing connections within the pod's termination grace period.
- **Authorization:** Never trust a client-supplied shipment/driver ID without verifying the
  authenticated principal actually owns or is authorized to view that resource.

---

## 21. Message Queue Development Guidelines

Driver notifications and cross-domain events flow through **Apache Kafka**.

- **Topic naming:** `<domain>.<entity>.<event>.v<N>` (e.g., `shipment.status.updated.v1`,
  `driver.notification.dispatched.v1`).
- **Schema management:** Message payloads are Protobuf, registered in the schema registry with
  backward-compatible evolution only (adding optional fields; never repurposing a field number).
- **Partitioning:** Partition key is the aggregate ID that requires ordering (e.g., `shipment_id`)
  so all events for one shipment are processed in order by a single consumer.
- **Delivery semantics:** Kafka gives us **at-least-once** delivery — all consumers **must be
  idempotent** (dedupe via event ID, typically using a Redis or Postgres dedupe table with TTL).
- **Retries & DLQ:** Transient failures retry with exponential backoff; after N attempts, messages
  route to a `<topic>.dlq` dead-letter topic for investigation rather than blocking the partition.
- **Payload size:** Keep messages small (< 1 MB). For large payloads (e.g., bulk manifest files),
  use the **claim-check pattern** — store the object in S3 and reference its key in the message.
- **Consumer lag** is a first-class SLO, monitored via Prometheus and alerted on.

---

## 22. Webhook Implementation Guidelines

Carrier integrations rely on inbound and outbound webhooks.

**Outbound (Auwalone+ → carrier):**

- Every payload is signed with HMAC-SHA256 using a per-carrier secret; the signature and a
  timestamp are sent in headers (`X-Auwalone-Signature`, `X-Auwalone-Timestamp`).
- Requests older than 5 minutes are rejected by receivers to prevent replay attacks.
- Delivery uses exponential backoff with jitter and a circuit breaker per carrier endpoint so one
  degraded carrier cannot exhaust shared retry workers.
- Every event carries a unique `event_id`; carriers are expected to dedupe on it.

**Inbound (carrier → Auwalone+):**

- **Always verify the signature before processing the payload** — reject unsigned or
  invalid-signature requests with `401` before any business logic runs.
- Respond within **5 seconds**; heavy processing is handed off to a Kafka topic and processed
  asynchronously.
- Deduplicate using the sender's event ID before applying side effects.

> [!WARNING]
> Skipping signature verification on inbound webhooks — even temporarily "to test something" — is
> treated as a security incident. Use the sandbox carrier receiver in the local stack instead.

---

## 23. Feature Flag Policy

We use feature flags for progressive delivery of risky or incomplete work.

- Naming: `<domain>_<feature>_<qualifier>` (e.g., `tracking_multi_leg_eta_v2`).
- Every flag has a named **owner** and an **expiry date** recorded in the flag configuration and
  in the introducing PR description.
- Flags that gate a fully-rolled-out feature must be **removed within 30 days** of reaching
  100% rollout — "flag debt" is tracked and reviewed monthly.
- Risky changes should ship behind a **kill-switch flag** that can disable the feature without a
  deploy.
- CI must exercise both the "on" and "off" states for any code path guarded by a flag under test.

---

## 24. Versioning Strategy

- **Public REST APIs:** URL-path versioned (`/v1`, `/v2`). Deprecated versions are supported for
  a minimum of **6 months**, with a `Sunset` HTTP header and advance notice in the changelog.
- **Shared Go libraries (`pkg/`):** [Semantic Versioning 2.0.0](https://semver.org/) via Git tags
  (`pkg/v1.4.2`).
- **Northline CLI:** SemVer, released independently of the platform.
- **Container images:** Tagged with both the Git SHA (for traceability) and, on release, the
  SemVer tag.
- **Helm charts:** Versioned independently of application version, following SemVer, with
  `appVersion` tracking the deployed service version.
- **Changelogs:** Generated automatically from Conventional Commit history on each release.

---

## 25. Dependency Management

- Go module dependencies are managed via `go.mod`/`go.sum`, committed to the repo; minimal
  version selection is respected — do not hand-edit `go.sum`.
- Run `govulncheck ./...` locally before adding a new dependency with known CVEs.
- Dependabot/Renovate opens weekly update PRs; security-relevant updates are prioritized and
  merged within 5 business days of release.
- Container base images use minimal/distroless variants and are pinned by digest in
  multi-stage Dockerfiles.
- Terraform provider versions are pinned in `versions.tf` per environment.
- License compatibility is checked automatically (policy-as-code) — copyleft licenses
  incompatible with our Apache 2.0 distribution require legal sign-off before use.
- Adding a new significant dependency (auth libraries, ORMs, message brokers) should be discussed
  in an issue or [ADR](#32-architecture-decision-records-adrs) first — we favor a small, well
  understood dependency graph.

---

## 26. CI/CD Expectations

All pipelines run on **GitHub Actions**; deployment is **GitOps** via **ArgoCD**.

**Pull Request pipeline (required to pass before merge):**

```yaml
# .github/workflows/ci.yml (excerpt)
jobs:
  lint:
    steps: [checkout, setup-go, golangci-lint, tflint, helm-lint]
  test:
    steps: [checkout, setup-go, unit-tests, integration-tests-testcontainers]
  security:
    steps: [gosec, govulncheck, gitleaks, trivy-image-scan]
  build:
    steps: [docker-build, sbom-generate-syft, push-to-ghcr-with-sha-tag]
```

**Merge-to-main pipeline:** re-runs the full suite, publishes the image tag to the container
registry, and opens an automated PR bumping the image tag in the GitOps environment repo.

**Continuous Delivery:** ArgoCD auto-syncs `dev` on every `main` merge; promotion to `staging` and
`production` requires a manual approval gate and uses **canary/blue-green rollouts** via
Argo Rollouts, with automatic rollback on elevated error rate or latency (Prometheus-based
analysis).

**Required status checks** (branch protection on `main`): `lint`, `test`, `security`, `build`, and
`dco-check`. PRs cannot be merged with any of these failing, or without the required number of
approvals.

---

## 27. Pull Request Review Checklist

Use this checklist before requesting review — reviewers will use the same list.

**Contributor checklist:**

- [ ] PR is linked to an issue (`Closes #____`)
- [ ] Commits are signed off (`git commit -s`) and follow Conventional Commits
- [ ] `make check` passes locally (fmt, lint, unit tests, vuln scan)
- [ ] New/changed behavior has unit and, where applicable, integration tests
- [ ] Public API changes update `api/openapi/` or `api/proto/` and pass `buf breaking`
- [ ] Database schema changes include tested up/down migrations
- [ ] Docs (`README`, runbooks, API reference) updated where behavior changed
- [ ] New metrics/logs/traces added for new user-facing behavior
- [ ] Feature flag added for risky/incomplete changes, with owner and expiry noted
- [ ] No secrets, credentials, or PII in code, logs, or test fixtures
- [ ] PR description explains **why**, not just **what**
- [ ] Diff is reasonably small and focused (~400 lines or split into follow-ups)

**Reviewer checklist:**

- [ ] Does the change match its stated intent and issue?
- [ ] Are edge cases and error paths tested, not just the happy path?
- [ ] Is backward compatibility preserved, or is the breaking change justified and versioned?
- [ ] Are security, performance, and observability requirements met?
- [ ] Is the code understandable to someone unfamiliar with this PR six months from now?

---

## 28. Reporting Bugs

Before opening a new issue:

1. **Search existing issues** (including closed ones) to avoid duplicates.
2. Confirm you're on a supported version.
3. Try to reproduce with the minimal possible setup.

Then open an issue using the **[Bug Report Template](#30-bug-report-template)**, including clear
reproduction steps, expected vs. actual behavior, environment details, and relevant logs
(**redact any secrets or customer data**). Label severity honestly — it helps us triage:

| Label | Meaning |
|---|---|
| `severity: critical` | Production outage or data loss |
| `severity: high` | Major feature broken, no workaround |
| `severity: medium` | Feature degraded, workaround exists |
| `severity: low` | Cosmetic or minor inconvenience |

---

## 29. Requesting Features

- Describe the **problem**, not just a proposed solution — this lets maintainers and the
  community evaluate alternatives.
- Use the **Feature Request** issue template.
- For significant architectural or cross-domain proposals, start a
  [GitHub Discussion](https://github.com/auwalone-plus/platform/discussions) and consider drafting
  an [ADR](#32-architecture-decision-records-adrs) before implementation — this avoids wasted
  effort on large PRs that need architectural rework.
- Maintainers triage feature requests weekly and label them `needs-discussion`, `accepted`, or
  `declined` with rationale.

---

## 30. Bug Report Template

This is the content of `.github/ISSUE_TEMPLATE/bug_report.md`:

```markdown
---
name: Bug report
about: Report a reproducible problem with Auwalone+
labels: bug, needs-triage
---

**Describe the bug**
A clear, concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Go to '...'
2. Call endpoint '...' with payload '...'
3. Observe '...'

**Expected behavior**
What you expected to happen.

**Actual behavior**
What actually happened. Include relevant (redacted) logs, request/response payloads, or
screenshots.

**Environment**
- Auwalone+ version / commit SHA:
- Deployment: local / staging / production
- Service(s) affected:

**Severity**
- [ ] Critical (outage/data loss)
- [ ] High (major feature broken)
- [ ] Medium (workaround exists)
- [ ] Low (cosmetic)

**Additional context**
Anything else relevant (recent deploys, feature flags enabled, etc.)
```

---

## 31. Pull Request Template

This is the content of `.github/PULL_REQUEST_TEMPLATE.md`:

```markdown
## Summary

<!-- What does this PR do and why? -->

Closes #

## Type of change

- [ ] feat
- [ ] fix
- [ ] docs
- [ ] refactor
- [ ] perf
- [ ] test
- [ ] chore / ci / build

## How was this tested?

<!-- Unit / integration / manual steps -->

## Checklist

- [ ] Commits are signed off (DCO) and follow Conventional Commits
- [ ] `make check` passes locally
- [ ] Tests added/updated
- [ ] Docs / API specs updated (if applicable)
- [ ] Database migrations include up/down and were tested (if applicable)
- [ ] Observability (logs/metrics/traces) added for new behavior (if applicable)
- [ ] Feature flag added with owner + expiry (if applicable)
- [ ] No secrets or PII introduced

## Screenshots / additional context

<!-- Optional -->
```

---

## 32. Architecture Decision Records (ADRs)

Significant architectural decisions — new datastores, cross-cutting patterns, major dependency
adoption, breaking API changes — are documented as ADRs in `docs/adr/`, using a lightweight
[MADR](https://adr.github.io/madr/)-style format.

```markdown
# NNNN. Title of the decision

Date: YYYY-MM-DD
Status: proposed | accepted | deprecated | superseded by NNNN

## Context
What problem are we solving? What forces are at play?

## Decision
What are we doing?

## Consequences
What becomes easier or harder as a result? What are the trade-offs?
```

ADRs are proposed via a normal Pull Request and reviewed like code — discussion happens in the
PR thread, and the ADR's `Status` is updated once a decision is reached. An accepted ADR that is
later reversed is marked `superseded by` and links to the new one; we never silently delete
historical ADRs.

---

## 33. Release Process

- **Services** are released continuously — every merge to `main` that passes CI is a release
  candidate, promoted through `dev → staging → production` via ArgoCD with automated and manual
  gates.
- **Northline CLI and shared libraries** follow a scheduled release cadence (typically bi-weekly)
  with SemVer tags and generated changelogs.
- Each release:
  1. Cuts a `release/x.y` branch (for CLI/libraries) or relies on the promoted image SHA
     (for services).
  2. Runs the full regression + smoke test suite against the target environment.
  3. Publishes GitHub Release notes generated from Conventional Commit history.
  4. Is announced in `#auwalone-announcements` on Slack and, for external-facing changes, in the
     public changelog at `docs.auwalone.dev/changelog`.
- **Rollback:** ArgoCD retains the previous synced revision; rollback is a one-command
  `argocd app rollback` plus, if needed, a reverse database migration following the
  expand-contract plan.

---

## 34. Community & Communication Channels

| Channel | Purpose |
|---|---|
| [GitHub Discussions](https://github.com/auwalone-plus/platform/discussions) | Design proposals, Q&A, RFC-style conversations |
| [GitHub Issues](https://github.com/auwalone-plus/platform/issues) | Bugs and feature requests |
| Slack — `#auwalone-contributors` ([invite](https://auwalone.dev/slack)) | Day-to-day contributor chat |
| Community Call | Bi-weekly, Thursdays — agenda posted in Discussions |
| `docs.auwalone.dev` | Public documentation site |
| security@auwalone.dev | Private vulnerability reports |
| conduct@auwalone.dev | Code of Conduct reports |

---

## 35. Recognition for Contributors

We value every contribution, regardless of size:

- All contributors are listed in `CONTRIBUTORS.md`, generated via the
  [all-contributors](https://allcontributors.org/) bot — comment `@all-contributors please add
  @username for code` on any PR/issue to credit someone.
- Notable contributions are highlighted in release notes and in the community call.
- Consistent, high-quality contributors are invited to become **Reviewers**, and from there,
  **Maintainers** with merge rights — our typical path is
  **first-time contributor → regular contributor → reviewer → maintainer**.
- Milestone contributors receive Auwalone+ swag — reach out in `#auwalone-contributors`.

---

## 36. License & Legal Considerations

- Auwalone+ is licensed under the **[Apache License 2.0](LICENSE)**. By contributing, you agree
  that your contributions will be licensed under the same terms.
- We use the **Developer Certificate of Origin (DCO)** instead of a CLA. Every commit must be
  signed off:
  ```bash
  git commit -s -m "feat(orders): add partial delivery support"
  ```
  This certifies you wrote the contribution or otherwise have the right to submit it under the
  project's license. A `dco-check` runs in CI and will block unsigned commits.
- Do not include proprietary, confidential, or third-party licensed code you do not have the
  right to contribute.
- If you incorporate a third-party library, ensure its license is compatible with Apache 2.0 —
  our automated license-compliance check will flag conflicts, but when in doubt, ask in
  `#auwalone-contributors` before opening the PR.

---

Thank you for helping build Auwalone+. We're glad you're here — now go find a
[`good first issue`](https://github.com/auwalone-plus/platform/labels/good%20first%20issue) and
let's ship something. 🚚
