# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Instruksi Pemeliharaan CLAUDE.md

Selesaikan setiap sesi atau tugas besar dengan memperbarui file ini. Pastikan untuk mencatat progres terbaru dan keputusan teknis yang diambil, serta hapus detail lama yang sudah tidak relevan agar context window tetap efisien.

## Commands

```bash
# Install dependencies
go mod download

# Run the application
go run main.go

# Run database migrations
go run migrations/main.go

# Build binary
go build -o bin/app main.go

# Run tests
go test ./...

# Run a single test
go test ./source/features/<feature>/... -run TestName -v
```

## Architecture

This project uses a **Feature-Based Clean Architecture** with Gin + GORM + PostgreSQL.

### Request Flow

```
HTTP Request
  → Gin Router (source/services/route.go)
  → Middleware chain (RequestID → Auth → LogActivity)
  → Handler (HTTP layer, parses request)
  → Usecase (business logic)
  → Repository (database access via GORM)
  → HTTP Response via httpresputils
```

### Feature Structure

Each feature lives in `source/features/<feature-name>/` and follows this pattern — **copy `exampletocopy/` as a template**:

| File | Role |
|---|---|
| `handler.go` | Defines `Handler` struct + `NewHandler(db)` factory that wires repo → usecase → handler |
| `handler_impl.go` | Implements `Impl(c *gin.Context)` — parses request, calls usecase, writes response |
| `usecase.go` | Defines `Usecase` interface + `usecaseImpl` struct + `injectUsecase()` |
| `usecase_impl.go` | Implements business logic methods |
| `repository.go` | Defines `Repository` interface + `repositoryImpl` struct + `injectRepository()` |
| `repository_impl.go` | Implements database queries via GORM |

`NewHandler(db)` returns a `gin.HandlerFunc` directly — the wiring is internal to the package.

### Registering a New Feature

1. Create `source/features/<name>/` using the exampletocopy pattern
2. Register the route in `source/services/route.go` inside `MountRouters()`

### Middleware

Middleware is in `source/services/middleware/`. The chain in `main.go`:

- `RequestIDMiddleware` — injects/propagates `X-Request-ID` header; stored in context as `"request_id"`
- `AuthMiddleware` — validates JWT Bearer token, sets `"id"`, `"name"`, `"nik"`, `"jabatan"`, `"divisi"`, `"role"` in context
- `AdminMiddleware` — checks `role == "admin"` from context (must run after Auth)
- `LogActivityMiddleware(db)` — creates a `log_activities` row; stores `"log_id"` in context for later enrichment

### Key Packages

**`source/pkg/`**
- `db/postgres.go` — `Database(cfg) (*gorm.DB, error)` — single entry point for DB connection
- `logger/zlog.go` — zerolog wrapper; call `logger.Info()`, `logger.Error()`, etc. globally. `logger.GinZLogger()` is the HTTP request logger middleware.

**`source/common/`**
- `glob_utils/http_resp_utils/` — standardized JSON response helpers: `HttpRespOK`, `HttpRespCreated`, `HttpRespBadRequest`, `HttpResponseUnAuth`, `HttpResponseForbidden`, etc. Every response includes `status`, `app_name`, `app_version`, `timestamp`.
- `glob_utils/jwt_utils/` — `CreateAccessToken()` (24h), `CreateRefreshToken()` (30d), `VerifyToken()`, `GetCurrentUser(c)`, `GetCurrentUserID(c)`
- `glob_utils/hashing_password/` — `HashPassword()` (bcrypt cost 14), `VerifyPassword()`
- `glob_utils/logs_activity/` — `AppendLogsActivity(ctx, db, before, after)` — enriches an existing `log_activities` row with before/after JSON snapshots (call after mutation)
- `models/` — shared GORM models: `UserModel`, `RefreshTokenModel`, `LogActivityModel`

### Models

All models use `uuid.UUID` as primary key with auto-generation in `BeforeCreate`. Soft deletes via `gorm.DeletedAt` on `UserModel`. Add new models to `migrations/main.go`'s `dbConn.AutoMigrate(...)` call.

Roles: `RoleStaff = "staff"`, `RoleAdmin = "admin"` (defined in `source/common/models/user.go`).

### Configuration

Loaded from `.env` via `godotenv` in `source/config/config.go`. Required env vars:

```
APP_NAME, APP_PORT
DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_NAME
JWT_SECRET
```

App version is read from the `version` file at the project root.
