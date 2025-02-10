# TaskFlow API

A lightweight, dependency-free task-management REST API written in Go.

TaskFlow is a small backend service for organising work into **projects**, **tasks**,
**tags**, and **comments**. It ships with an in-memory store (great for demos and
tests) behind a `Store` interface, so a database-backed implementation can be dropped
in without touching the HTTP layer.

## Features

- Clean `net/http` router with method + path matching
- In-memory repository behind a storage interface
- JSON request/response helpers with consistent error envelopes
- Middleware: request logging, panic recovery, CORS
- Health and readiness endpoints
- Zero third-party dependencies — standard library only

## Getting started

```bash
go run ./cmd/server
```

The server listens on `:8080` by default. Override with `PORT`:

```bash
PORT=9090 go run ./cmd/server
```

## API

| Method | Path              | Description        |
|--------|-------------------|--------------------|
| GET    | `/healthz`        | Liveness probe     |
| GET    | `/api/tasks`      | List tasks         |
| POST   | `/api/tasks`      | Create a task      |
| GET    | `/api/tasks/{id}` | Fetch a task       |

## Layout

```
cmd/server        # main entrypoint
internal/api       # HTTP router, handlers, middleware
internal/models    # domain types
internal/store     # storage interface + in-memory impl
internal/config    # environment configuration
```

## License

MIT
