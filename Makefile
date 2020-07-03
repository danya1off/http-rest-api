.PHONY: build
build:
	go build -v ./cmd/restapi

.PHONY: run
run:
	go run ./cmd/restapi/main.go

.PHONY: test
test:
	go test -v -race -timeout 2s ./...

.DEFAULT_GOAL := build