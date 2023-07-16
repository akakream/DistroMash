.PHONY: build run test vendor

BUILD_NAME = distromash
BUILD_DIR = $(PWD)/bin

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
