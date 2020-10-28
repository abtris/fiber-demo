# https://suva.sh/posts/well-documented-makefiles/#simple-makefile
.DEFAULT_GOAL:=help
SHELL:=/bin/bash

.PHONY: help deps build watch run

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

deps:  ## Check && download dependencies
	go mod download
	go mod tidy
	npm install -g nodemon

build: deps ## Build docker image
	docker build .

watch: ## Watch file changes and build
	nodemon --exec go run server.go --signal SIGTERM

run: ## Run
	go run server.go
