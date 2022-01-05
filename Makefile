GO_PKG_LIST := $(shell go list ./...)

.PHONY: test
test:
	GOFLAGS=-mod=vendor go test -short ${GO_PKG_LIST}

.PHONY: coverage
coverage:
	GOFLAGS=-mod=vendor go test -short ${GO_PKG_LIST} -coverprofile=coverage.out

.PHONY: coverhtml
coverhtml:
	go tool cover -html=coverage.out

.PHONY: staticcheck
staticcheck:
	GOFLAGS=-mod=vendor staticcheck ${GO_PKG_LIST}

.PHONY: lint
lint:
	golint -set_exit_status ${GO_PKG_LIST}

.PHONY: build
build:
	go build -o nebula cmd/nebula/*.go