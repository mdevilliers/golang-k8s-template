.PHONY: image

BUILD_IMAGE = golang:1.7.3

IMAGE_NAME := {{ cookiecutter.project_name }}
IMAGE_REPOSITORY_NAME := {{cookiecutter.docker_repository}}
IMAGE_TAG := $(shell ./scripts/image-tag)
EXE_NAME := {{ cookiecutter.project_name }}

GO_BUILD = go build
GO_BUILD_FLAGS = -a --installsuffix cgo --ldflags="-s" -o bin/$(EXE_NAME)
GO_BUILD_CMD = CGO_ENABLED=0 $(GO_BUILD) $(GO_BUILD_FLAGS) {{ cookiecutter.package_path }}/{{ cookiecutter.project_name }}/cmd/{{ cookiecutter.project_name }}

DOCKER_BUILD_CMD = docker run --rm -it -v "$(GOPATH)":/g -v "$(shell pwd)/docker/":/app -e "GOPATH=/g" -w /app $(BUILD_IMAGE) sh -c '$(GO_BUILD_CMD)'
DOCKER_PACKAGE_CMD = docker build -t $(IMAGE_REPOSITORY_NAME)/$(IMAGE_NAME):$(IMAGE_TAG) docker/

build:
	$(GO_BUILD_CMD)

image :
	$(DOCKER_BUILD_CMD)
	$(DOCKER_PACKAGE_CMD)

