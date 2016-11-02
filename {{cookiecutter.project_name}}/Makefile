.PHONY: image

EXE_NAME := {{ cookiecutter.project_name }}

GO_BUILD = go build
GO_BUILD_FLAGS = -a --installsuffix cgo --ldflags="-s"
GO_BUILD_CMD = CGO_ENABLED=0 $(GO_BUILD) $(GO_BUILD_FLAGS) -o bin/$(EXE_NAME) github.com/{{ cookiecutter.github_account }}/{{ cookiecutter.project_name }}/cmd/{{ cookiecutter.project_name }}

DOCKER_IMAGE_NAME := {{ cookiecutter.project_name }}
DOCKER_REPOSITORY_NAME := {{cookiecutter.docker_repository}}
DOCKER_IMAGE_TAG := $(shell ./scripts/image-tag)

DOCKER_BUILD_CMD = CGO_ENABLED=0 $(GO_BUILD) $(GO_BUILD_FLAGS) -o docker/$(EXE_NAME) github.com/{{ cookiecutter.github_account }}/{{ cookiecutter.project_name }}/cmd/{{ cookiecutter.project_name }}
DOCKER_PACKAGE_CMD = docker build -t $(DOCKER_REPOSITORY_NAME)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) -t $(DOCKER_REPOSITORY_NAME)/$(DOCKER_IMAGE_NAME):latest docker/

build:
	$(GO_BUILD_CMD)

image :
	$(DOCKER_BUILD_CMD)
	$(DOCKER_PACKAGE_CMD)
