.PHONY: build run test vendor kill-all

BUILD_NAME = distromash
BUILD_DIR = $(PWD)/bin
PORTS_TO_KILL := 5005 3000 3001 3002

build-all: multiplatform2ipfs-build p2pcomm-build ipdr2-build controller-build

run-all: multiplatform2ipfs-run p2pcomm-run ipdr2-run controller-run
	trap 'kill $(jobs -p)' SIGINT; \
	wait $(jobs -p)

kill-all:
	@for port in $(PORTS_TO_KILL); do \
        echo "Killing processes on port $$port"; \
        kill $$(lsof -t -i :$$port); \
    done

multiplatform2ipfs-build:
	@echo "Building MultiPlatform2IPFS"
	@cd MultiPlatform2IPFS && go build -o ./bin/multiplatform2ipfs

p2pcomm-build:
	@echo "Building P2PComm"
	@cd P2PComm && go build -o ./bin/app 

controller-build:
	@echo "Building Controller"
	@go build -o $(BUILD_DIR)/$(BUILD_NAME)

ipdr2-build:
	@echo "Building IPDR2"
	@cd ipdr2/cmd/ipdr && go build -o ../../bin/ipdr2

multiplatform2ipfs-run:
	@echo "Running MultiPlatform2IPFS"
	@cd MultiPlatform2IPFS && ./bin/multiplatform2ipfs server --port=3002 &

p2pcomm-run:
	@echo "Running P2PComm"
	@cd P2PComm && ./bin/app server --port=3001 --datastore & 

controller-run:
	@echo "Running Controller"
	@$(BUILD_DIR)/$(BUILD_NAME) run

ipdr2-run:
	@echo "Running IPDR2"
	@cd ipdr2 && ./bin/ipdr2 server -p 5005 &


build:
	@go build -o $(BUILD_DIR)/$(BUILD_NAME)

run:
	@$(BUILD_DIR)/$(BUILD_NAME) run

clean:
	rm -rf $(BUILD_DIR)

test:
	go test -v ./... -count=1

vendor:
	@go mod vendor

tidy:
	@go mod tidy

lint:
	golangci-lint run -v ./...

swag:
	swag init
