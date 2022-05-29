###############################################################################
#
# @description Commands to setup, develop and deploy the project
# @author <tyrkov.zahar@gmail.com>
#
###############################################################################
# Make
###############################################################################

SHELL=/bin/bash
ENVIRONMENT=testing

################################################################
GO_VERSION ?= 1.18


# --- Makefile starts here
.DEFAULT_GOAL := help

# ================== Makefile commands ====================
.PHONY: build dep generate format lint test race clean help


build: ## Build the binary files
	@echo "==> Sourcing .env ..." \
		&& set -a; source .env; set +a \
		&& echo "==> Building binaries..." \
		&& go build -o bin/server cmd/server/server.go \
		&& go build -o bin/client cmd/client/client.go \
		&& echo "==> Done ./bin/ "

dep: ## Get the dependencies
	go get -v -t -d ./...

generate: ## Generate Go code from protobuf
	@mkdir -p gen
	@protoc -I=./pkg/calc --go_out=gen --go_opt=paths=source_relative \
	 --go-grpc_out=gen --go-grpc_opt=paths=source_relative calc.proto
	@echo "==> Generated proto files in ./gen"

format: ## format the files
	@gofmt -w .

lint: ## Lint the files
	@go vet ./...

test: ## Run unittests
	@go test -v ./...

race: ## Run data race detector, and generate coverage file
	@go test -race -short -coverprofile=coverage.txt -covermode=atomic ./...

clean: ## Remove previous build
	@rm -rf bin coverage.txt

help: ## Help target
	@ag '^[a-zA-Z_-]+:.*?## .*$$' --nofilename $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN{FS=": ## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


