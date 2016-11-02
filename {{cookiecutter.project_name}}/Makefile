
EXE_NAME := {{ cookiecutter.project_name }}

GO_BUILD := go build
GO_BUILD_VARS := CGO_ENABLED=0 GOOS=linux
GO_BUILD_FLAGS := -a -tags netgo -ldflags="-w"
GO_BUILD_CMD := $(GO_BUILD_VARS) $(GO_BUILD) $(GO_BUILD_FLAGS) -o bin/$(EXE_NAME) github.com/{{ cookiecutter.github_account }}/{{ cookiecutter.project_name }}/cmd/{{ cookiecutter.project_name }}

DOCKER_IMAGE_NAME := {{ cookiecutter.project_name }}
DOCKER_REPOSITORY_NAME := {{cookiecutter.docker_repository}}
DOCKER_IMAGE_TAG := $(shell ./scripts/image-tag)

DOCKER_BUILD_CMD := $(GO_BUILD_VARS) $(GO_BUILD) $(GO_BUILD_FLAGS) -o docker/$(EXE_NAME) github.com/{{ cookiecutter.github_account }}/{{ cookiecutter.project_name }}/cmd/{{ cookiecutter.project_name }}
DOCKER_PACKAGE_CMD := docker build -t $(DOCKER_REPOSITORY_NAME)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) -t $(DOCKER_REPOSITORY_NAME)/$(DOCKER_IMAGE_NAME):latest docker/

build: ## builds project to bin/
	$(GO_BUILD_CMD)

image: ## builds and tags docker image to configured repository
	$(DOCKER_BUILD_CMD)
	$(DOCKER_PACKAGE_CMD)

.PHONY: image

deploy: ## deploys to configured kubernetes instance
	kubectl delete -f k8s/ 2>/dev/null; true
	kubectl create -f k8s/

.PHONY: deploy

# 'help' parses the Makefile and displays the help text
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help
