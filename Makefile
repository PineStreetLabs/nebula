GO_PKG_LIST := $(shell go list ./...)

.PHONY: test
test:
	GOFLAGS=-mod=vendor go test -short ${GO_PKG_LIST}

.PHONY: staticcheck
staticcheck:
	GOFLAGS=-mod=vendor staticcheck ${GO_PKG_LIST}

.PHONY: lint
lint:
	golint -set_exit_status ${GO_PKG_LIST}

.PHONY: build
build:
	go build -o nebula cmd/nebula/*.go