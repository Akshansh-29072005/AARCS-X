# AARCS-X — AI-Based Automated Attendance Verification and Blockchain-Secured Academic Records

This repository contains two deployable applications:
- A Go-based backend API that manages students, teachers, institutes, and future auth flows.
- A cross-platform mobile app built with Expo/React Native for student and faculty interactions.

The goal is to deliver reliable attendance verification and tamper-resistant academic records, with production-grade practices for security, scalability, and observability.

Requirements
- Go 1.20+ (as per [backend/go.mod](backend/go.mod)).
- PostgreSQL 13+.
- Node 16+ with npm or Yarn.
- Expo tooling (`npx expo` is sufficient).

Architecture Overview
- Backend API: Gin framework with modular `internal` packages for configuration, platform concerns (DB/HTTP/server), and domain modules (`students`, `teachers`, `institutes`, `auth`).
- Database: Postgres via `pgx` pool. Schema migrations in [backend/migrations](backend/migrations).
- Mobile App: Expo/React Native, modular screens under [mobile/app](mobile/app), shared components/helpers in [mobile/components](mobile/components) and [mobile/lib](mobile/lib).
- Dev tooling: Live reload with Air (see [backend/.air.toml](backend/.air.toml)), example environment in [backend/.env.example](backend/.env.example).

Production Readiness — What’s Implemented vs What To Finalize
- Implemented
	- Health check route at [/](backend/internal/platform/server/routes.go#L9) with DB connectivity indication.
	- Connection pooling and ping validation in [backend/internal/platform/database/postgres.go](backend/internal/platform/database/postgres.go).
	- Environment-driven config in [backend/internal/config/config.go](backend/internal/config/config.go), supports `.env` via `godotenv`.
	- Basic CORS in [backend/cmd/api/main.go](backend/cmd/api/main.go#L34-L47).
	- Students and Teachers modules with create/list endpoints.
- To finalize for production
	- Secure CORS: restrict `AllowOrigins` to your domains; remove `*`.
	- TLS: run behind a reverse proxy (Nginx/Traefik) with HTTPS and HSTS.
	- Secrets: load `DATABASE_URL`, `JWT_SECRET`, and other secrets from a vault (AWS Secrets Manager/SOPS), not `.env`.
	- AuthN/AuthZ: implement JWT/OAuth flows in [backend/internal/auth](backend/internal/auth) and enforce in middleware.
	- Observability: structured logs (Zap/Logrus), metrics (Prometheus), tracing (OpenTelemetry), health/readiness endpoints.
	- Rate limiting & security headers: add middleware for DoS resistance and headers (CSP, X-Frame-Options, etc.).
	- API versioning & docs: provide `/v1` routes and Swagger/OpenAPI spec.
	- CI/CD: build, test, lint, containerize, run migrations, and deploy via GitHub Actions.
	- Performance: DB indexes, pagination, caching (Redis), Gin in release mode.

Backend — Setup & Run
1) Database & environment
- Create a Postgres DB and user:
```bash
sudo -u postgres psql -c "CREATE USER aarcs_user WITH PASSWORD 'secret';"
sudo -u postgres psql -c "CREATE DATABASE aarcs_db OWNER aarcs_user;"
```
- Set environment variables (used by [backend/internal/config/config.go](backend/internal/config/config.go)):
```bash
export DATABASE_URL=postgres://aarcs_user:secret@localhost:5432/aarcs_db?sslmode=disable
export PORT=8000
# Recommended for prod:
# export JWT_SECRET=supersecret
# export ALLOWED_ORIGINS=https://yourapp.com,https://admin.yourapp.com
```

2) Migrations
- Apply SQL migrations from [backend/migrations](backend/migrations) using `golang-migrate`:
```bash
migrate -path backend/migrations -database "$DATABASE_URL" up
```

3) Run the API
```bash
cd backend
go run ./cmd/api
# or
go build ./cmd/api && ./api
```

API Endpoints (current)
- Health: `GET /` — returns status and DB connectivity ([backend/internal/platform/server/routes.go](backend/internal/platform/server/routes.go)).
- Students:
	- `POST /api/students` — create student ([backend/internal/students/routes.go](backend/internal/students/routes.go)).
	- `GET  /api/students` — list students ([backend/internal/students/routes.go](backend/internal/students/routes.go)).
- Teachers:
	- `POST /api/teachers` — create teacher ([backend/internal/teachers/routes.go](backend/internal/teachers/routes.go)).

Backend — Folder-by-Folder
- [backend/cmd/api/main.go](backend/cmd/api/main.go): Entry point. Loads config, connects DB, configures CORS, registers routes, starts server.
- [backend/internal/config](backend/internal/config): Centralized configuration loader. Advantages: single source of truth, environment-driven, supports `.env` while favoring real env vars in production.
- [backend/internal/platform/database](backend/internal/platform/database): Postgres connection pool bootstrap. Advantages: encapsulates connection handling, ping checks for early failure, ready for pool tuning in prod.
- [backend/internal/platform/http](backend/internal/platform/http): Standardized JSON responses and validation error formatter. Advantages: consistent API responses, easier client handling, centralized error formatting.
- [backend/internal/platform/server](backend/internal/platform/server): Server-level routes (health, readiness) and future global wiring. Advantages: clear separation of infra routes vs domain routes.
- [backend/internal/platform/middleware](backend/internal/platform/middleware): Placeholder for cross-cutting middleware (auth, rate limit, request ID, compression). Advantages: reusable policies, centralized security and observability.
- [backend/internal/platform/logger](backend/internal/platform/logger): Placeholder for structured logging integrations (Zap/Logrus). Advantages: correlation IDs, leveled logging, JSON output for log aggregation.
- [backend/internal/students](backend/internal/students): Domain module (DTO/entity/model/repository/service/handler/routes). Advantages: clean layering, testable components, easier evolution.
- [backend/internal/teachers](backend/internal/teachers): Same structure as `students`. Advantages: consistency, isolated business logic.
- [backend/internal/institutes](backend/internal/institutes): Scaffolding for institute domain. Advantages: future expansion without touching core infra.
- [backend/internal/auth](backend/internal/auth): Scaffolding for authentication flows (JWT/OAuth). Advantages: clear boundary for security-critical code.
- [backend/migrations](backend/migrations): SQL migration files (`up`/`down`). Advantages: deterministic schema changes, CI-safe database evolution.
- [backend/tmp](backend/tmp): Scratch/build artifacts during dev. Advantages: keeps repo clean by isolating transient files.
- [backend/.air.toml](backend/.air.toml): Air live-reload config for local dev. Advantages: faster iteration; disabled in prod.
- [backend/.env.example](backend/.env.example): Example environment file. Advantages: onboarding support; do not use for secrets in prod.

Mobile — Setup & Run
```bash
cd mobile
npm install
npx expo start
```
- Base API URL is defined in [mobile/lib/api.ts](mobile/lib/api.ts). Update to your deployed backend (e.g., `https://api.yourapp.com/api`).
- Screens:
	- [mobile/app/(auth)](mobile/app/%28auth%29): login, signup.
	- [mobile/app/(student)](mobile/app/%28student%29): dashboard, records, settings.
- Shared code:
	- [mobile/components](mobile/components): UI primitives.
	- [mobile/constants](mobile/constants): themes and tokens.
	- [mobile/context](mobile/context): global app state (e.g., theme).
	- [mobile/hooks](mobile/hooks): reusable hooks.
	- [mobile/lib](mobile/lib): API helpers and mock data.
- For production builds, use EAS Build or native builds, configure OTA updates and environment values via `app.json` and secure runtime config.

Security & Compliance Checklist
- Use `JWT_SECRET` or OAuth provider; store secrets in a vault.
- Validate inputs at handler boundary; sanitize outputs.
- Enforce HTTPS; set HSTS and security headers.
- Implement role-based access control for `student`, `faculty`, `institute`.
- Audit logging for sensitive operations; redact PII in logs.
- Backups and retention for Postgres; define recovery RPO/RTO.

Observability & Reliability
- Logging: add structured JSON logs (correlation IDs, user IDs where appropriate).
- Metrics: expose Prometheus metrics (requests, latencies, DB pool stats).
- Tracing: instrument with OpenTelemetry (HTTP handlers, repository calls).
- Health: add `/healthz` (basic) and `/readyz` (DB reachable, migrations applied).
- Error handling: standardized error model via [backend/internal/platform/http](backend/internal/platform/http).

Performance & Scalability
- Run Gin in release mode (`GIN_MODE=release`).
- Add rate limiting and request size limits.
- Use Redis for caching hot reads (e.g., lists) and session data.
- Paginate list endpoints; add DB indexes for email/phone.
- Consider CQRS for write-heavy modules; async jobs for long-running tasks.

API Versioning & Documentation
- Prefix routes with `/v1` and maintain a changelog.
- Generate Swagger/OpenAPI (e.g., `swaggo/swag`) and host at `/docs`.

CI/CD & Deployment
- CI (GitHub Actions): lint, test, build, run migrations, build Docker image, deploy.
- Containerization: Dockerfile for backend; run behind Nginx/Traefik with TLS.
- Config via environment; secrets via vault. No `.env` in production.
- Zero-downtime deploys: readiness gates + migration strategy.

Environment Variables
- Backend (minimum): `DATABASE_URL`, `PORT`
- Recommended: `JWT_SECRET`, `ALLOWED_ORIGINS`, `LOG_LEVEL`, `GIN_MODE`, `REDIS_URL`

Try It
```bash
# Start backend
export DATABASE_URL=postgres://aarcs_user:secret@localhost:5432/aarcs_db?sslmode=disable
export PORT=8000
migrate -path backend/migrations -database "$DATABASE_URL" up
cd backend && go run ./cmd/api

# Start mobile
cd mobile && npx expo start
# Update mobile/lib/api.ts to point to http://localhost:8000/api (use LAN/URL accessible by device)
```

Questions or prioritizations
- If you want me to implement JWT auth, Swagger docs, Dockerfile + GitHub Actions, or production CORS/rate limiting now, say which to prioritize and I’ll add the code and CI next.


