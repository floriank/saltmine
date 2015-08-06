VERSION := $(shell git describe --always --long --dirty)
LD_FLAGS := "-X main.Version ${VERSION}"
ROOT_PKG := github.com/floriank/saltmine

.PHONY: all build lint test vet format install run

all: format lint vet test build

format:
	go fmt ./...

lint:
	golint ./...

vet:
	go vet ./...

test:
	go test -race -v ./...

build:
	go build -ldflags ${LD_FLAGS} ${ROOT_PKG}

install:
	go install -ldflags ${LD_FLAGS} ${ROOT_PKG}

run:
	./saltmine
