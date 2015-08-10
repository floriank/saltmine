VERSION := $(shell git describe --always --long --dirty)
LD_FLAGS := "-X main.Version ${VERSION}"
ROOT_PKG := github.com/floriank/saltmine

.PHONY: all build lint test vet format install run

all: format lint vet test build

format:
	godep go fmt ./...

lint:
	golint ./...

vet:
	godep go vet ./...

test:
	godep go test -v ./...

build:
	godep go build -ldflags ${LD_FLAGS} ${ROOT_PKG}

install:
	godep go install -ldflags ${LD_FLAGS} ${ROOT_PKG}

run:
	./saltmine
