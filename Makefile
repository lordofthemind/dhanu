# Define variables
APP_NAME = dhanu
BIN_DIR = bin
SRC_DIR = cmd
PKG_DIR = pkgs
INTERNALS_DIR = internals
CONFIG_FILE = config.yaml
TMP_DIR = tmp
BUILD_DIR = build

# Define the Go toolchain
GO = go
GOBUILD = $(GO) build
GOTEST = $(GO) test
GOCLEAN = $(GO) clean

# Default target
all: build

# Build commands
build_win: ## Build the application for Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/win/$(APP_NAME).exe ./main.go

build_lin: ## Build the application for Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/linux/$(APP_NAME) ./main.go

build_mac: ## Build the application for MacOS
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/mac/$(APP_NAME) ./main.go

build: ## Build the application for all platforms and save it in the build folder
	$(MAKE) build_lin
	$(MAKE) build_win
	$(MAKE) build_mac

# Run the application
run: build
	@echo "Running the application..."
	./$(BIN_DIR)/$(APP_NAME)

# Testing commands
test: ## Run tests
	@echo "Running tests..."
	$(GOTEST) ./...

lint: ## Run linter
	@echo "Linting the code..."
	golangci-lint run

# Utility commands
clean: ## Clean the build and the logs
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)/
	rm -rf logs/

tr: ## Generate directory tree
	tree > tree.txt

# Help command
help: ## Show this help message
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1m%-12s\033[0m %s\n\n", "Command", "Description"} /^[a-zA-Z_-]+:.*?##/ { printf "\033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: all build run test lint clean tr help
