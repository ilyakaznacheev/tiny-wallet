GOBASE := $(shell pwd)
GOPATH := $(GOPATH):$(GOBASE)/vendor
GOBIN := $(GOBASE)/bin
VERSION := 0.1.0
TEST_TIMEOUT := 30s

## build: Compile the app into the executable
build:
	@echo ">  Getting dependent packages..."
	@go get
	@echo ">  Building the app..."
	@go build  cmd/tiny-wallet/wallet.go
	@echo ">  Done"

## test: Run unit-tests of the project
test:
	@echo ">  Testing the app..."
	@go test -covermode=count -timeout=$(TEST_TIMEOUT) \
		. \
		./pkg/currency/ \
		./internal/config
	@echo ">  Testing done"

test-coverage:
	@./deployments/ci/test.sh


## run: Run the application
run:
	@echo ">  Running the app..."
	@go run cmd/tiny-wallet/wallet.go
	@echo ">  Done"

## help: Get makefile manual
help: Makefile
	@echo
	@echo Choose command to run:
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: build test test-coverage help