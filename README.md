# Student CRUD API

A simple RESTful service for managing student records built with Go and SQLite.

---

## 🧾 Overview

This is a lightweight, self‑contained API that lets you create, retrieve and list students. It exists as a learning/demo project and as a starting point for more elaborate Go backend services. The system solves the problem of keeping student data in a small file‑based database and exposes the data over HTTP with a clean, easy‑to‑understand structure.

Use case: educational projects, prototypes or as a template for CRUD services.

---

## 🏗 Architecture

The code follows a simple layered layout loosely inspired by Clean Architecture. There’s a clear direction of dependencies: handlers depend on a `storage.Storage` interface, which has a concrete SQLite implementation in `internal/storage/sqlite`. Configuration is loaded at startup and passed through.

- **cmd/student-api/**: application entry point. Sets up config, database and routes.
- **internal/http/handlers/**: request handling logic, validation and response formatting.
- **internal/storage/**: storage interface and SQLite repository.
- **internal/config/**: configuration loader using cleanenv.
- **internal/types/**: shared domain types (e.g. `Student`).
- **internal/utils/response/**: helpers for JSON responses and error formatting.

Requests enter via `main.go`, are dispatched by `http.ServeMux`, processed by handler functions, and ultimately satisfy the `storage` interface; dependencies point inward.

---

## 🛠 Tech Stack

- **Language:** Go 1.20+
- **HTTP framework:** Built‑in `net/http`
- **Database:** SQLite (file)
- **Driver:** github.com/mattn/go-sqlite3
- **Validation:** github.com/go-playground/validator/v10
- **Config loader:** github.com/ilyakaznacheev/cleanenv
- **Logging:** Go 1.21 `log/slog`
- **Containerization:** none yet (could add Docker)
- **CI/CD:** not configured

---

## 📁 Project Structure

```
cmd/                # application entrypoints
  student-api/main.go
internal/            # private application code
  config/            # configuration loading
  http/handlers/     # request handlers
  storage/           # storage interface & implementations
  types/             # shared domain types
  utils/response/    # JSON helpers
config/              # example YAML config file
```

- `cmd` holds the executable.
- `internal` contains code that may not be imported by other projects.
- there is no `pkg` directory in this simple service.
- configuration is kept under `config`, migrations/docs aren’t needed for the SQLite demo.

---

## ⭐ Features

- CRUD operations for students (create, list, retrieve by ID)
- Request validation and structured JSON responses
- Centralised error formatting
- Simple graceful shutdown with context timeouts

Authentication, authorization and other advanced features are not implemented yet.

---

## 🚀 Setup & Installation

```bash
git clone https://github.com/amandx36/StudentApi-Go.git
cd StudentCrudApiGo

go mod tidy
```

Copy `config/local.yml` (or create your own) and set the `storage_path` and `http_server.address` values. Alternatively set `CONFIG_PATH` env var or supply `-config` flag.

---

## ⚙️ Environment Configuration

The app reads a YAML config file with `cleanenv`. Example fields:

```yaml
env: local
storage_path: ./students.db
http_server:
  address: :8080
```

All values are also exposed via environment variables (`ENV`, etc.). Secrets should be managed externally; there are no credentials in this sample.

---

## ▶️ Running the Application

```bash
CONFIG_PATH=config/local.yml go run ./cmd/student-api
```

The server listens on the address defined in config (default `:8080`).

---

## 🐳 Docker Support

No Dockerfile is currently provided. You can quickly containerise by writing a simple `Dockerfile` and optionally a `docker-compose.yml` for local development.

---

## 📚 API Documentation

No Swagger/OpenAPI spec yet. Endpoints:

- `POST /api/v1/students` – create a student
- `GET /api/v1/students/{id}` – get by ID
- `GET /api/v1/getStudents` – list all students

Request/response bodies are JSON with the `Student` fields.

---

## 🗄 Database Migrations

Schema is created automatically when the service starts. No migration tooling is included.

---

## 🧪 Testing Strategy

No tests are defined. A future version should include unit tests with Go’s `testing` package, table-driven tests for handlers, and mocking of the `storage` interface.

---

## ❗ Error Handling Strategy

Errors are wrapped as `response.Response` structs with a status and message. Validation errors produce human‑readable messages. HTTP status codes are set appropriately in handlers.

---

## 📈 Logging Strategy

Structured logs use `slog` with info/error levels. The current implementation is minimal; correlation IDs are not yet implemented.

---

## 🔒 Security Practices

This demo does not handle authentication. Input is validated via `validator`. CORS, rate limiting, password hashing, JWT policies, etc., are not yet addressed.

---

## ⚙️ Production Readiness

For production you’d place the service behind a reverse proxy (nginx), enable HTTPS, add connection pooling, monitoring/metrics, and remove any hard‑coded settings.

---

## 📈 Scalability Considerations

The service is stateless aside from the SQLite file; horizontal scaling would require moving to a networked database and adding caching/load balancing.

---

## 🌱 Roadmap

- Add authentication/authorization
- Introduce automated tests and CI pipelines
- Add Docker support and Kubernetes manifests
- Replace SQLite with Postgres/MySQL and add migrations
- Implement OpenAPI spec and generate docs

---

## 📜 License

This project is released under the MIT License.

---

_Quality indicators:_ clear architecture, reasonable folder structure, no secrets in code, graceful shutdown, validation, and logging are present. Docker, tests and security features can be added in future versions.
