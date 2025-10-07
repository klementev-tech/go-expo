.PHONY:  up lint install-golangci docker-build

docker-build:
	docker build -t klementev-tech/go-expo .

up:
	docker compose up -d

lint:
	golangci-lint run ./...

GOLANGCI_VERSION=v2.5.0

setup:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_VERSION)