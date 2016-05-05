DOCKER_REGISTRY ?= gcr.io
IMAGE_PREFIX    ?= kubernetes-helm
SHORT_NAME      ?= tiller

# go option
GO        ?= go
GOARCH    ?= $(shell go env GOARCH)
GOOS      ?= $(shell go env GOOS)
PKG       := $(shell glide novendor)
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   :=
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
BINARIES  := helm tiller

.PHONY: all
all: build

.PHONY: build
build:
	GOBIN=$(BINDIR) $(GO) install $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' github.com/kubernetes/helm/cmd/...

.PHONY: check-docker
check-docker:
	@if [ -z $$(which docker) ]; then \
	  echo "Missing \`docker\` client which is required for development"; \
	  exit 2; \
	fi

.PHONY: docker-binary
docker-binary: GOOS = linux
docker-binary: GOARCH = amd64
docker-binary: BINDIR = $(CURDIR)/rootfs
docker-binary: GOFLAGS += -a -installsuffix cgo
docker-binary: build

.PHONY: docker-build
docker-build: check-docker docker-binary
	docker build --rm -t ${IMAGE} rootfs
	docker tag -f ${IMAGE} ${MUTABLE_IMAGE}

.PHONY: test
test: build
test: TESTFLAGS += -race -v
test: test-style
test: test-unit

.PHONY: test-unit
test-unit:
	$(GO) test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

.PHONY: test-style
test-style:
	@scripts/validate-go.sh

.PHONY: clean
clean:
	@rm -rf $(BINDIR)

.PHONY: coverage
coverage:
	@scripts/coverage.sh

HAS_GLIDE := $(shell command -v glide;)
HAS_HG := $(shell command -v hg;)
HAS_GIT := $(shell command -v git;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
ifndef HAS_HG
	$(error You must install Mercurial (hg))
endif
ifndef HAS_GIT
	$(error You must install Git)
endif
	glide install

include versioning.mk
