# 📐 Hexagonal Go Architecture

A modular, scalable Go application designed with **Clean Architecture** and **Hexagonal principles**. This repository demonstrates how to organize a robust, enterprise-ready codebase using clear boundaries between domain logic, infrastructure, and delivery mechanisms.

---

## ✨ Key Highlights

- 🔌 Clean, loosely coupled code structure
- 🧱 Builder pattern for dependency injection
- 🌐 HTTP handlers with middleware support
- 📦 Pluggable outbound adapters (e.g., DB, Cache, APIs)
- 📜 Response utilities (JSON/XML, GZIP)
- 🧪 Testable interfaces and services

---

## 📦 Folder Structure
```
├── cmd/                  # Application entry points
│   ├── http/             # HTTP server bootstrap
│   └── jobs/             # Background workers / CRON jobs
│
├── internal/             # Internal application logic
│   ├── builder/          # App builder (DI, config, logger)
│   ├── core/
│   │   ├── constants/    # HTTP codes, errors, app-level constants
│   │   ├── model/        # Domain models
│   │   ├── ports/        # Inbound/outbound interfaces
│   │   ├── service/      # Business use cases
│   │   └── repository/   # Repository interfaces
│   ├── inbound/          # Adapters: HTTP handlers, middleware
│   ├── outbound/         # Infrastructure: DB, Cache, External APIs
│   └── utils/            # Utility packages
│
├── scripts/              # Automation scripts (migrations, tools)
├── Dockerfile            # Multi-stage build Dockerfile
├── Makefile              # Dev commands (build, run, clean, etc.)
├── go.mod
└── go.sum
```


---

## 🚀 Getting Started

### 🔧 Prerequisites

- Go 1.20+
- Docker
- Make

### 🔨 Running Locally

```bash
git clone https://github.com/bhupendra-dudhwal/hexagonal-go-architecture.git
cd hexagonal-go-architecture
go run cmd/http/http.go
```

### 🐳 Using Docker
```bash
make build      # Build the Docker image
make run        # Run the container
make stop       # Stop the container
make clean      # Remove the Docker image
make rebuild    # Clean and rebuild
make tag        # Add version tag to image
```

---

# 🧰 Available Endpoints
```
| Method | Path    | Description     |
| ------ | ------- | --------------- |
| POST   | `/user` | Create new user |
```

# 🧪 Testing
You can structure unit and integration tests under internal/core/service and internal/core/ports. Add mocks as needed to simulate repository behaviors.

# 📌 Tech Stack
- Go (Golang)
- Gorilla Mux
- Docker + Alpine
- Clean Architecture
- Builder Pattern
- Middleware Architecture

# 🤝 Contributing
- Contributions are welcome! Please:
- Fork the repo
- Create a feature branch
- Raise a pull request

