# Contributing

Thanks for your interest in TaskFlow API!

## Workflow

1. Fork and create a feature branch (`feat/...`, `fix/...`, `docs/...`).
2. Keep changes small and focused; one logical change per pull request.
3. Run `go build ./...` and `go test ./...` before opening a PR.
4. Write a clear, imperative commit message.

## Code style

- Standard library first; avoid new dependencies without discussion.
- `gofmt` all code.
- Every exported symbol gets a doc comment.