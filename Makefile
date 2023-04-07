APP_NAME := im2im
SRC_PATH := cmd/im2im.go
BUILD_DIR := ./build
GOBIN := $(shell go env GOPATH)/bin

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_PATH)

.PHONY: run
run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

.PHONY: install
install:
	@echo "Installing dependencies..."
	@go mod download
