.PHONY: run build test clean

run:
	@go run cmd/main.go  # <-- TAB aqui, não espaços!

build:
	@go build -o bin/estudos-go cmd/main.go

test:
	@go test -v ./...

cover:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
