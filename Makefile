.PHONY: image run
.DEFAULT_GOAL := default

APP_NAME ?= Go Prime
APP_VSN ?= 0.0
BUILD ?= `git rev-parse --short HEAD`

help:
	@echo "$(APP_NAME):$(APP_VSN)-$(BUILD)"
	@perl -nle'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

image: ## Build the Docker image
	docker build --no-cache -t goprime .

rund: ## Run the app in Docker in detached mode
	docker run -p 8080:8080 -d goprime

run: ## Run the app in Docker in interactive mode
	docker run -p 8080:8080 -it goprime

default:
	docker build --no-cache -t goprime .
	docker run -p 8080:8080 -it goprime