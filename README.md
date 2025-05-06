# hexagonal-go-architecture
A modular, scalable Go application built with Clean Architecture and Hexagonal principles. This repo demonstrates how to organize your codebase using layers like handler, service, middleware, repository, and builder, suitable for enterprise-grade systems.

# Go Clean Hexagonal Architecture

A sample Go project that demonstrates how to build scalable and maintainable applications using **Clean Architecture** and **Hexagonal principles**.

---

## 📦 Folder Structure

.
├── cmd/ # Application entry points
│ ├── http/ # HTTP server startup
│ └── jobs/ # (Optional) Cron or background jobs
├── internal/ # Core application logic (not accessible outside)
│ ├── builder/ # Builder pattern for dependency injection
│ ├── core/
│ │ ├── constants/ # Shared constants (HTTP codes, error messages)
│ │ ├── model/ # Domain models
│ │ ├── ports/ # Inbound/outbound interfaces (use case boundaries)
│ │ ├── service/ # Business logic implementation
│ │ └── repository/ # Data access interfaces
│ ├── inbound/ # Inbound adapters (HTTP handlers, middleware)
│ ├── outbound/ # Outbound adapters (DB, external APIs, cache)
│ └── utils/ # Helper functions/utilities
├── scripts/ # Project-related scripts (migration, etc.)
├── go.mod
└── go.sum


---

## ✅ Features

- Clear separation of concerns
- Builder pattern for dependency injection
- HTTP handlers and middleware
- Extensible structure for DB and external APIs
- Testable services and interfaces
- Response wrapper (JSON/XML with optional gzip compression)

---

## 🚀 Getting Started

```bash
git clone https://github.com/bhupendra-dudhwal/hexagonal-go-architecture.git
cd hexagonal-go-architecture
go run cmd/http/http.go
