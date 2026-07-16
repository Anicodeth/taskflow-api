# Architecture

TaskFlow follows a small, layered design:

- **cmd/server** wires configuration, storage, and the HTTP router.
- **internal/api** owns routing, handlers, and middleware.
- **internal/store** defines storage interfaces and an in-memory implementation.
- **internal/models** holds pure domain types with no I/O.

The `Store` interfaces keep the HTTP layer independent of persistence, so a
SQL-backed implementation can be added without touching handlers.