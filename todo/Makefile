.PHONY: build test run-daemon run-cli

build:
	go build ./...

test:
	go test ./...

run-daemon:
	go run ./cmd/todod start

run-cli:
	go run ./cmd/todo --help
