.PHONY: build run test vendor

BUILD_NAME = distromash
BUILD_DIR = $(PWD)/bin

clean:
	rm -rf $(BUILD_DIR)

build:
	@go build -o $(BUILD_DIR)/$(BUILD_NAME)

run:
	@$(BUILD_DIR)/$(BUILD_NAME) run

test:
	go test -v ./... -count=1

vendor:
	@go mod vendor

lint:
	golangci-lint run -v ./...

swag:
	swag init
