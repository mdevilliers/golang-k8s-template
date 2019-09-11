
CMD      := {{cookiecutter.project_name}}
PKG      := github.com/{{cookiecutter.github_account}}/{{cookiecutter.project_name}}
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

EXE_NAME := {{ cookiecutter.project_name }}

# Versioning
GIT_COMMIT ?= $(shell git rev-parse HEAD)
GIT_SHA    ?= $(shell git rev-parse --short HEAD)
GIT_TAG    ?= $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  ?= $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

# Binary Name
BIN_OUTDIR ?= ./build/bin
BIN_NAME   ?= {{cookiecutter.project_name}}-$(shell go env GOOS)-$(shell go env GOARCH)
ifeq ("$(GIT_TAG)","")
	BIN_VERSION = $(GIT_SHA)
endif
BIN_VERSION ?= ${GIT_TAG}


DOCKER_REGISTRY := {{cookiecutter.docker_repository}}

# Docker Tag from Git
DOCKER_IMAGE_TAG  ?= ${GIT_TAG}
ifeq ("$(DOCKER_IMAGE_TAG)","")
	DOCKER_IMAGE_TAG = $(GIT_SHA)
endif

DOCKER_BUILD_CMD := $(GO_BUILD_VARS) $(GO_BUILD) $(GO_BUILD_FLAGS) -o docker/$(EXE_NAME) github.com/{{ cookiecutter.github_account }}/{{ cookiecutter.project_name }}/cmd/{{ cookiecutter.project_name }}
DOCKER_PACKAGE_CMD := docker build -t $(DOCKER_REPOSITORY_NAME)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) -t $(DOCKER_REPOSITORY_NAME)/$(DOCKER_IMAGE_NAME):latest docker/

# LDFlags
# LDFLAGS := -w -s
LDFLAGS += -X $(PKG)/internal/version.Timestamp=$(shell date +%s)
LDFLAGS += -X $(PKG)/internal/version.GitCommit=${GIT_COMMIT}
LDFLAGS += -X $(PKG)/internal/version.GitTreeState=${GIT_DIRTY}
LDFLAGS += -X $(PKG)/internal/version.Version=${BIN_VERSION}

# CGO
CGO ?= 1

# Go Build Flags
GOBUILDFLAGS :=
GOBUILDFLAGS += -o $(BIN_OUTDIR)/$(BIN_NAME)

# Linting
OS := $(shell uname)
GOLANGCI_LINT_VERSION=1.18.0
ifeq ($(OS),Darwin)
	GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-darwin-amd64.tar.gz
else
	GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64.tar.gz
endif

.PHONY: info
info:
	@echo "Version:        ${BIN_VERSION}"
	@echo "Binary Name:    ${BIN_NAME}"
	@echo "Git Tag:        ${GIT_TAG}"
	@echo "Git Commit:     ${GIT_COMMIT}"
	@echo "Git Tree State: ${GIT_DIRTY}"

# Build a statically linked binary
.PHONY: static
static: CGO = 0
static: build

# Build a binary
.PHONY: build
build: CMD = ./cmd/{{cookiecutter.project_name}}
build: GOBUILDFLAGS += -ldflags '$(LDFLAGS)'
build:
	@CGO_ENABLED=$(CGO) go build $(GOBUILDFLAGS) $(CMD)

# Build docker image
.PHONY: image
image:
	docker build \
		--force-rm \
		--build-arg BIN_VERSION=$(BIN_VERSION) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg GIT_SHA=$(GIT_SHA) \
		--build-arg GIT_TAG=$(GIT_TAG) \
		--build-arg GIT_DIRTY=$(GIT_DIRTY) \
		-f ./build/package/Dockerfile \
		-t $(DOCKER_REGISTRY):$(DOCKER_IMAGE_TAG) .

# Run test suite
test:
ifeq ("$(wildcard $(shell which gocov))","")
	go get github.com/axw/gocov/gocov
endif
	gocov test ${PKG_LIST} | gocov report

# deploys to configured kubernetes instance
.PHONY: deploy
deploy:
	kubectl delete -f k8s/ 2>/dev/null; true
	kubectl create -f k8s/

# the linting gods must be obeyed
.PHONY: lint
lint: $(BIN_OUTDIR)/golangci-lint/golangci-lint
	$(BIN_OUTDIR)/golangci-lint/golangci-lint run

$(BIN_OUTDIR)/golangci-lint/golangci-lint:
	curl -OL https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/$(GOLANGCI_LINT_ARCHIVE)
	mkdir -p $(BIN_OUTDIR)/golangci-lint/
	tar -xf $(GOLANGCI_LINT_ARCHIVE) --strip-components=1 -C $(BIN_OUTDIR)/golangci-lint/
	chmod +x $(BIN_OUTDIR)/golangci-lint
	rm -f $(GOLANGCI_LINT_ARCHIVE)

