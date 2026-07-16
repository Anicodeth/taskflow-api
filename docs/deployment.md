# Deployment

The service is a single static binary.

```bash
go build -o bin/server ./cmd/server
PORT=8080 ./bin/server
```

Container images are built from the provided `Dockerfile`. The process listens
on `$PORT` and shuts down gracefully on SIGTERM.