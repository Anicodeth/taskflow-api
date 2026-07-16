# HTTP API

All endpoints accept and return JSON.

## Conventions

- Errors use the envelope `{"error": "message"}`.
- Timestamps are RFC3339 UTC.
- Resource ids are strings.

## Resources

Each resource exposes `GET /api/<name>`, `POST /api/<name>`, and
`GET /api/<name>/{id}`. See the source for request/response shapes.