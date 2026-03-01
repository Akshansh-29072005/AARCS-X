# AARCS-X — AI-Based Automated Attendance Verification and Blockchain-Secured Academic Records

This repository contains two deployable applications:
- A Go-based backend API that manages students, teachers, institutes, and future auth flows.
- A cross-platform mobile app built with Flutter for student and faculty interactions.

The goal is to deliver reliable attendance verification and tamper-resistant academic records, with production-grade practices for security, scalability, and observability.

Requirements
- Go 1.20+ (as per [backend/go.mod](backend/go.mod)).
- PostgreSQL 13+.
- Flutter SDK (stable channel, Dart SDK included).
- Platform tooling for your target (`Android Studio` / `Xcode` / `Chrome` for web).

Architecture Overview
- Backend API: Gin framework with modular `internal` packages for configuration, platform concerns (DB/HTTP/server), and domain modules (`students`, `teachers`, `institutes`, `auth`).
- Database: Postgres via `pgx` pool. Schema migrations in [backend/migrations](backend/migrations).
- Mobile App: Flutter app with route-driven screens under [mobile/lib/screens](mobile/lib/screens), shared UI widgets under [mobile/lib/widgets](mobile/lib/widgets), and API/services under [mobile/lib/services](mobile/lib/services).
- Dev tooling: Live reload with Air (see [backend/.air.toml](backend/.air.toml)), example environment in [backend/.env.example](backend/.env.example).

## Why This Architecture?

This backend follows a modular monolith architecture to balance:
- Clear domain boundaries
- Low operational overhead
- Easy future extraction into microservices

Each domain (students, teachers, institutes, auth) owns its API, business logic, and persistence layer, while platform concerns (DB, middleware, logging) are centralized.

### Production Readiness Checklist

- [x] DB connection pooling
- [x] SQL migrations
- [x] Environment-based config
- [ ] JWT authentication
- [ ] Request ID + structured logging
- [ ] Metrics and tracing
- [ ] Rate limiting
- [ ] CI/CD pipeline

## Non-Goals (For Now)

- Kubernetes orchestration
- Microservices split
- Multi-region deployment

These are intentionally deferred to keep focus on correctness, observability, and security first.

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
- [backend/internal/departments](backend/internal/departments): Same structure as `institutes` including (DTO/entity/model/repository/service/handler/routes).
- [backend/internal/auth](backend/internal/auth): Scaffolding for authentication flows (JWT/OAuth). Advantages: clear boundary for security-critical code.
- [backend/migrations](backend/migrations): SQL migration files (`up`/`down`). Advantages: deterministic schema changes, CI-safe database evolution.
- [backend/tmp](backend/tmp): Scratch/build artifacts during dev. Advantages: keeps repo clean by isolating transient files.
- [backend/.air.toml](backend/.air.toml): Air live-reload config for local dev. Advantages: faster iteration; disabled in prod.
- [backend/.env.example](backend/.env.example): Example environment file. Advantages: onboarding support; do not use for secrets in prod.

Mobile — Setup & Run
```bash
cd mobile
flutter pub get
flutter run
```
- Base API URL is defined in [mobile/lib/services/api_client.dart](mobile/lib/services/api_client.dart). Update to your deployed backend (e.g., `https://api.yourapp.com/api/v1`).
- App entry and routing:
	- [mobile/lib/main.dart](mobile/lib/main.dart): app bootstrap (`Provider` + `MaterialApp.router`).
	- [mobile/lib/router.dart](mobile/lib/router.dart): navigation and auth-based route guards via `go_router`.
- Screens:
	- [mobile/lib/screens/auth](mobile/lib/screens/auth): login, register.
	- [mobile/lib/screens/dashboard](mobile/lib/screens/dashboard): dashboard shell.
	- [mobile/lib/screens/students](mobile/lib/screens/students): list and add student.
	- [mobile/lib/screens/teachers](mobile/lib/screens/teachers): list and add teacher.
	- [mobile/lib/screens/institutions](mobile/lib/screens/institutions), [mobile/lib/screens/departments](mobile/lib/screens/departments), [mobile/lib/screens/semesters](mobile/lib/screens/semesters), [mobile/lib/screens/subjects](mobile/lib/screens/subjects).
- Shared code:
	- [mobile/lib/widgets](mobile/lib/widgets): reusable UI widgets.
	- [mobile/lib/models](mobile/lib/models): typed request/response models.
	- [mobile/lib/services](mobile/lib/services): API integrations by domain.
	- [mobile/lib/providers](mobile/lib/providers): app state (`AuthProvider`).
	- [mobile/lib/theme](mobile/lib/theme): app theming.
- For production builds, use Flutter build targets (`apk`, `appbundle`, `ipa`, `web`) and inject environment-specific API config per build flavor.

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

⚠️ Security Note  
`.env` files are for local development only. Production secrets must be injected via environment variables or a secrets manager.

Try It
```bash
# Start backend
export DATABASE_URL=postgres://aarcs_user:secret@localhost:5432/aarcs_db?sslmode=disable
export PORT=8000
migrate -path backend/migrations -database "$DATABASE_URL" up
cd backend && go run ./cmd/api

# Start mobile
cd mobile && flutter pub get && flutter run
# Update mobile/lib/services/api_client.dart to point to your reachable backend URL
```

Questions or prioritizations
- If you want me to implement JWT auth, Swagger docs, Dockerfile + GitHub Actions, or production CORS/rate limiting now, say which to prioritize and I’ll add the code and CI next.

---

Backend progress update (Feb 2026)

Summary of what’s been built recently
- Migrated the backend to a consistent `/api/v1` route structure across all domains.
- Added new domain modules: Institutions, Departments, Semesters, and Subjects with full create/read flows.
- Expanded the database schema to support core academic structures (institutions → departments → semesters → subjects) plus users, attendance, and assessments.
- Added system health + metrics endpoints with CPU/memory usage and DB readiness checks.
- Introduced layered architecture (repository → service → handler) across new modules for clean separation and testability.

Current API v1 routes (backend)
- System
	- `GET /api/v1/system/health` — health check with DB connectivity ([backend/internal/platform/server/routes.go](backend/internal/platform/server/routes.go)).
	- `GET /api/v1/system/metrics` — CPU/memory + DB readiness ([backend/internal/platform/server/routes.go](backend/internal/platform/server/routes.go)).
- Institutions
	- `POST /api/v1/institutions` — create institution ([backend/internal/institutes/routes.go](backend/internal/institutes/routes.go)).
	- `GET  /api/v1/institutions` — list/filter institutions by `name`, `code` ([backend/internal/institutes/dto.go](backend/internal/institutes/dto.go)).
- Departments
	- `POST /api/v1/departments` — create department ([backend/internal/departments/routes.go](backend/internal/departments/routes.go)).
	- `GET  /api/v1/departments` — list/filter by `name`, `code`, `head_of_department`, `institution_id` ([backend/internal/departments/dto.go](backend/internal/departments/dto.go)).
- Semesters
	- `POST /api/v1/semesters` — create semester ([backend/internal/semesters/routes.go](backend/internal/semesters/routes.go)).
	- `GET  /api/v1/semesters` — list/filter by `number`, `department_id` ([backend/internal/semesters/dto.go](backend/internal/semesters/dto.go)).
- Subjects
	- `POST /api/v1/subjects` — create subject ([backend/internal/subjects/routes.go](backend/internal/subjects/routes.go)).
	- `GET  /api/v1/subjects` — list/filter by `name`, `code`, `semester_id` ([backend/internal/subjects/dto.go](backend/internal/subjects/dto.go)).
- Students
	- `POST /api/v1/students` — create student ([backend/internal/students/routes.go](backend/internal/students/routes.go)).
	- `GET  /api/v1/students` — list/filter by `semester_id`, `department_id`, `institution_id` ([backend/internal/students/dto.go](backend/internal/students/dto.go)).
- Teachers
	- `POST /api/v1/teachers` — create teacher ([backend/internal/teachers/routes.go](backend/internal/teachers/routes.go)).
	- `GET  /api/v1/teachers` — list/filter by `department_id`, `designation` ([backend/internal/teachers/dto.go](backend/internal/teachers/dto.go)).

Schema evolution & migrations
- New schema migration: [backend/migrations/20260120185925_create_aarcs_schema.up.sql](backend/migrations/20260120185925_create_aarcs_schema.up.sql)
	- Core entities: `institutions`, `departments`, `semesters`, `subjects`, `students`, `teachers`, `users`.
	- Academic workflow: `teacher_subjects`, `attendance`, `assessments`.
	- Clear FK relationships align with the API domain boundaries.
- Cleanup migration: [backend/migrations/20260120182416_drop_old_student_teacher_tables.up.sql](backend/migrations/20260120182416_drop_old_student_teacher_tables.up.sql) to remove legacy tables.

Backend architecture progress
- The domain modules now follow a consistent pattern: DTOs → Repository → Service → Handler → Routes. This makes the codebase easier to test, scale, and reason about.
- System health is now a first-class concern with real-time metrics in [backend/internal/platform/server/routes.go](backend/internal/platform/server/routes.go), backed by database ping + OS stat sampling.

Environment variables (backend)
- [backend/.env.example](backend/.env.example) currently requires:
	- `DATABASE_URL`
	- `PORT`

Auth module status
- Auth routes are scaffolded in [backend/internal/auth/routes.go](backend/internal/auth/routes.go) with placeholders for register/login/logout and protected `me` endpoints.
<!-- Next production steps recommended based on current progress
- Enable the auth module and add JWT middleware; wire it into the `/api/v1` groups.
- Add migrations for user credential hashing (bcrypt) and role-based access control.
- Add pagination and consistent error response envelopes for list endpoints.
- Expand observability by adding structured logs and request IDs for all routes. -->

---

Backend progress update (Latest snapshot)(24 Feb 2026)

This section reflects the latest backend state and supersedes earlier “auth pending” notes.

What progressed significantly
- Auth is now wired into application startup in [backend/cmd/api/main.go](backend/cmd/api/main.go): repository, service, and handler are initialized and routes are registered.
- JWT secret bootstrapping is active through `utlis.SetJWTSecret(cfg.JWTSecret)` in [backend/cmd/api/main.go](backend/cmd/api/main.go).
- Authentication middleware is implemented in [backend/internal/platform/middleware/auth.go](backend/internal/platform/middleware/auth.go) and validates Bearer tokens, then injects `user_id`, `role`, and `ref_id` into request context.
- Role-based authorization helper exists in [backend/internal/platform/middleware/role.go](backend/internal/platform/middleware/role.go) via `RequireRole(...)`.
- Auth service now supports institution registration + login flows in [backend/internal/auth/service.go](backend/internal/auth/service.go), including password hashing and JWT issuance.

Current auth API status
- Base group: `/api/v1/auth` ([backend/internal/auth/routes.go](backend/internal/auth/routes.go))
- Public
	- `POST /api/v1/auth/register` → Register institution + create auth user + return JWT.
	- `POST /api/v1/auth/login` → Validate credentials + return JWT.
- Protected (requires Bearer token)
	- `GET /api/v1/auth/protected/me` → Returns current token identity (`user_id`, `role`, `ref_id`).

Auth flow (implemented)
1. Registration request reaches `RegisterInstitution` in [backend/internal/auth/handler.go](backend/internal/auth/handler.go).
2. Institution is created using the institution service adapter.
3. Password is hashed using utility functions in [backend/internal/platform/utlis/password.go](backend/internal/platform/utlis/password.go).
4. User record is inserted in `users` through [backend/internal/auth/repository.go](backend/internal/auth/repository.go).
5. JWT is generated via [backend/internal/platform/utlis/jwt.go](backend/internal/platform/utlis/jwt.go).

Config progress
- `JWT_SECRET` is now part of config in [backend/internal/config/config.go](backend/internal/config/config.go).
- Important: [backend/.env.example](backend/.env.example) still lists only `DATABASE_URL` and `PORT`; add `JWT_SECRET` there to match runtime expectations.

Backend maturity highlights
- Domain modules active: `auth`, `institutes`, `departments`, `semesters`, `subjects`, `teachers`, `students`.
- Consistent layered structure across domains (DTO/Repository/Service/Handler/Routes).
- System observability baseline in place via:
	- `GET /api/v1/system/health`
	- `GET /api/v1/system/metrics`




