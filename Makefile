fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v -race ./...

build:
	go build -o bin/app.exe ./cmd/app/main.go

all: fmt lint test build