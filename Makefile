BINARY_NAME=lazycommit

build:
	go build -o ${BINARY_NAME} -v

test:
	go test -v ./...

deps:
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run

.PHONY: build test deps fmt lint