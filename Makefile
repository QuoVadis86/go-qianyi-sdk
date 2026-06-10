# QERP (千易) Go SDK Makefile
# https://github.com/QuoVadis86/go-qianyi-sdk

MODULE := github.com/QuoVadis86/go-qianyi-sdk

# https://github.com/golang/go/wiki/Go-version-tags
GoVersion := $(shell go env GOVERSION 2>/dev/null || echo "unknown")

.PHONY: all
all: lint test build

.PHONY: build
build:
	@echo ">> building $(MODULE)"
	go build ./...

.PHONY: test
test:
	@echo ">> running tests"
	go test -race -count=1 -coverprofile=coverage.out ./...
	@echo ">> test coverage: $$(go tool cover -func=coverage.out | tail -1 | awk '{print $$NF}')"

.PHONY: coverage
coverage: test
	go tool cover -html=coverage.out

.PHONY: lint
lint:
	@echo ">> running golangci-lint"
	golangci-lint run ./... --timeout=5m

.PHONY: vet
vet:
	@echo ">> running go vet"
	go vet ./...

.PHONY: fmt
fmt:
	@echo ">> formatting code"
	go fmt ./...

.PHONY: clean
clean:
	@echo ">> cleaning up"
	rm -f coverage.out
	go clean -i ./...

.PHONY: example
example:
	@echo ">> building example"
	go build -o /dev/null ./example/

.PHONY: ci
ci: fmt vet lint test build

.PHONY: tidy
tidy:
	@echo ">> tidying module"
	go mod tidy
	go mod verify
