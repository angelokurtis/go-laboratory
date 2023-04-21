ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build: ## Compiles the packages.
	CGO_ENABLED=0 go build -o time-tracker ./cmd/time-tracker

.PHONY: generate
generate: wire ## Generate code.

.PHONY: wire
wire: ## Generate code for dependency injection.
	wire ./...

.PHONY: wsl
wsl: ## Forces the use of empty lines to improve readability.
	wsl -fix ./...

.PHONY: format
format: wsl ## Forces code format.
	gofumpt -l -w .

.PHONY: lint
lint: ## Examines programming errors, bugs, and stylistic inconsistencies.
	golangci-lint run ./...

.PHONY: test
test: generate ## Run tests.
	go test ./... -covermode=count -coverprofile=coverage.out
	go tool cover -func=coverage.out -o=coverage.out
