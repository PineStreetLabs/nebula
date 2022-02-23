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
	golangci-lint run -v --exclude-use-default=false --disable-all --enable=golint

.PHONY: build
build:
	go build -o nebula cmd/nebula/*.go

.PHONY: install
install:
	go install -v ./...