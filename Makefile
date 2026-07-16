.PHONY: build test run vet tidy

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

vet:
	go vet ./...

run:
	go run ./cmd/server