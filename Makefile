.PHONY: build run test lint

build:
	go build -o bin/main ./cmd/app

run: build
	./bin/main

test:
	go test -race ./...

lint:
	golangci-lint run ./...