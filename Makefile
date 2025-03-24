.PHONY: help build build-example tests unit-tests unit-tests-cov version
.DEFAULT_GOAL := help

BINARY_NAME=example
BIN_DIR=./bin
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build-example: ## Build example project
	@mkdir -p ./_example/$(BIN_DIR)
	go build -C _example -o $(BIN_DIR)/$(BINARY_NAME)

build: ## Check if code is buildable
	go build ./...

tests: unit-tests ## Run all tests

unit-tests: ## Run unit tests
	go test -v ./...

unit-tests-cov: ## Run unit tests with coverage
	go test -v -coverprofile=coverage.txt ./...

version: ## Create new version. Bump, tag, commit, create tag
	@bump-my-version bump --verbose $(filter-out $@,$(MAKECMDGOALS))

%:
	@:

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
