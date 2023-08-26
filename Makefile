.PHONY: build run test vendor

BUILD_NAME = distromash
BUILD_DIR = $(PWD)/bin

build-all: multiplatform2ipfs-build p2pcomm-build ipdr2-build controller-build

run-all: multiplatform2ipfs-run p2pcomm-run ipdr2-run controller-run

multiplatform2ipfs-build:
	@echo "Building MultiPlatform2IPFS"
	@cd MultiPlatform2IPFS && go build -o ./bin/multiplatform2ipfs

p2pcomm-build:
	@echo "Building P2PComm"
	@cd ipfs-libp2p-pubsub && go build -o ./bin/app 

controller-build:
	@echo "Building Controller"
	@go build -o $(BUILD_DIR)/$(BUILD_NAME)

ipdr2-build:
	@echo "Building IPDR2"
	@cd ipdr2/cmd/ipdr && go build -o ./bin/ipdr2

multiplatform2ipfs-run:
	@echo "Running MultiPlatform2IPFS"
	@cd MultiPlatform2IPFS && ./bin/multiplatform2ipfs server --port=3002 &

p2pcomm-run:
	@echo "Running P2PComm"
	@cd ipfs-libp2p-pubsub && ./bin/app server --port=3001 --datastore & 

controller-run:
	@echo "Running Controller"
	@$(BUILD_DIR)/$(BUILD_NAME) run

ipdr2-run:
	@echo "Running IPDR2"
	@cd ipdr2 ./bin/ipdr2 server -p 5005 &


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
