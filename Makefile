BIN=day3

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

GOFILES = $(shell find . -name '*.go')
GOPACKAGES = $(shell go list ./...)

all: dependencies build install

dependencies:
	go mod download

test: dependencies
	@go test -v $(GOPACKAGES)

benchmark: dependencies fmt
	@go test $(GOPACKAGES) -bench=.

fmt:
	gofmt -w .

build:
	go build -o $(GOBIN)/$(BIN) ./cmd/$(BIN)/*.go || exit

build-all: fmt test build

run: build-all
	./bin/$(BIN)

default: build

.PHONY: build project fmt