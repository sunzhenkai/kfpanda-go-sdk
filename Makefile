.PHONY: mod build test main

mod:
	@go mod tidy
	@go mod download
	@go mod vendor

build:
	@go build ./...

test:
	@go test -v ./...

main:
	@CGO_ENABLED=0 go build -ldflags="-s -w -extldflags=-static" cmd/main/kfpanda-client.go && ./kfpanda-client
