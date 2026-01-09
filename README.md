# AI-Based Automated Attendance Verification and Blockchain-Secured Academic Record Management System

This repository implements an AI-assisted automated attendance verification system combined with blockchain-backed academic record management. It consists of a Go backend (API + database + migrations) and a cross-platform mobile app (Expo / React Native) for students and teachers.

Quick overview
- Backend: `backend/` — Go API server, Postgres database, migrations in `migrations/`.
- Internals: `backend/internals/` — configuration, database connection, handlers, models, repositories, routes, and utilities.
- Mobile: `app/` — Expo / React Native app with screens under `(auth)`, `(student)`, and shared components.
- Assets & UI: `assets/`, `components/`, `constants/`, `lib/` — reusable UI and helper code.

Repository layout (important paths)

- `backend/cmd/api/main.go` — main entry for the API server.
- `backend/internals/config/config.go` — configuration loader for env variables.
- `backend/internals/database/postgres.go` — Postgres connection helper.
- `backend/migrations/` — SQL migrations (up/down) for schema changes.
- `app/` — Expo app source (screens, components, hooks).

Requirements

- Go 1.20+ (or as required by `go.mod`).
- PostgreSQL (local or remote) for the backend.
- Node 16+ / npm or Yarn for the Expo mobile app.
- Expo CLI (optional, for development builds): `npm install -g expo-cli` or use `npx expo`.

Backend — quick setup and run

1. Create a Postgres database and user. Example (local):

```bash
# create DB and user (adjust names and password)
sudo -u postgres psql -c "CREATE USER aarcs_user WITH PASSWORD 'secret';"
sudo -u postgres psql -c "CREATE DATABASE aarcs_db OWNER aarcs_user;"
```

2. Set required environment variables (example `.env` or export):

```bash
export DB_HOST=127.0.0.1
export DB_PORT=5432
export DB_USER=aarcs_user
export DB_PASSWORD=secret
export DB_NAME=aarcs_db
export APP_PORT=8080
```

3. Run migrations (the repo includes SQL migration files in `migrations/`). Use your preferred migration tool (e.g., `migrate`). Example with `migrate`:

```bash
# install: https://github.com/golang-migrate/migrate
migrate -path backend/migrations -database "postgres://aarcs_user:secret@localhost:5432/aarcs_db?sslmode=disable" up
```

4. Build and run the API server:

```bash
cd backend
go build ./cmd/api
./api
# or run directly
go run ./cmd/api
```

Mobile (Expo) — quick setup and run

1. Install dependencies and start Expo Metro bundler:

```bash
cd app
npm install    # or yarn
npx expo start
```

2. Open on device or simulator using Expo Go or a native build.

Development notes

- Configuration is read from environment variables via `backend/internals/config`.
- Database interactions are implemented in `internals/repository` and models are in `internals/models`.
- Handlers are under `internals/handlers` (e.g., `student_handler.go`, `teachers_handler.go`).
- If you change database schemas, add corresponding `up`/`down` SQL files to `migrations/` and run the migration tool.

Testing & linting

- Backend: use `go test ./...` to run Go tests (if added).
- Mobile: use Expo's built-in tools and React Native testing libraries as needed.

Further work and tips

- Add a `.env.example` with required environment variables for contributors.
- Consider a `Makefile` or scripts to automate `migrate`, build, and run steps.
- Document APIs in a `docs/` folder or add OpenAPI/Swagger generation in the backend.

Contact

If you want me to expand any section (detailed API usage, full env example, CI setup, or adding a `.env.example`), tell me which part to prioritize and I'll update the README and add supporting files.

