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

.PHONY: format
format: goimports-reviser gofumpt wsl ## Reformat Go code.
	$(GOIMPORTS_REVISER) -project-name github.com/angelokurtis/go-laboratory/database-locking -use-cache -rm-unused -set-alias -format ./...
	$(GOFUMPT) -w -extra .
	$(WSL) -fix ./...

.PHONY: generete
generete: wire ## Generate Go code.
	$(WIRE) ./...

WIRE = $(shell pwd)/bin/wire
.PHONY: wire
wire: ## Download wire locally if necessary.
	$(call go-get-tool,$(WIRE), github.com/google/wire/cmd/wire,v0.5.1-0.20220620021424-0675cdc9191c)

GOIMPORTS_REVISER = $(shell pwd)/bin/goimports-reviser
.PHONY: goimports-reviser
goimports-reviser: ## Download goimports-reviser locally if necessary.
	$(call go-get-tool,$(GOIMPORTS_REVISER), github.com/incu6us/goimports-reviser/v3,v3.3.1)

GOFUMPT = $(shell pwd)/bin/gofumpt
.PHONY: gofumpt
gofumpt: ## Download gofumpt locally if necessary.
	$(call go-get-tool,$(GOFUMPT), mvdan.cc/gofumpt,v0.5.0)

WSL = $(shell pwd)/bin/wsl
.PHONY: wsl
wsl: ## Download wsl locally if necessary.
	$(call go-get-tool,$(WSL), github.com/bombsimon/wsl/v4/cmd...,v4.1.0)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
go get -d $(2)@$(3) ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef