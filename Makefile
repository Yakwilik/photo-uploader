# Makefile for photo-uploader

# Variables
BINARY_NAME=photo-uploader
BUILD_DIR=build
VERSION?=dev
COMMIT?=unknown
DATE?=unknown

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-s -w -X github.com/khasbulatabdullin/photo-uploader/internal/version.Version=$(VERSION) -X github.com/khasbulatabdullin/photo-uploader/internal/version.Commit=$(COMMIT) -X github.com/khasbulatabdullin/photo-uploader/internal/version.Date=$(DATE)"

.PHONY: all build clean test deps run help version

# Default target
all: clean deps test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/photo-uploader
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for multiple platforms
build-all:
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/photo-uploader
	GOOS=linux GOARCH=386 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-386 ./cmd/photo-uploader
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/photo-uploader
	GOOS=linux GOARCH=arm $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm ./cmd/photo-uploader
	
	# Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/photo-uploader
	GOOS=windows GOARCH=386 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-386.exe ./cmd/photo-uploader
	
	# macOS
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/photo-uploader
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/photo-uploader
	
	@echo "Build complete for all platforms in $(BUILD_DIR)/"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) verify

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	$(GOCMD) run ./cmd/photo-uploader

# Run with custom port
run-port:
	@echo "Running $(BINARY_NAME) on port $(PORT)..."
	$(GOCMD) run ./cmd/photo-uploader --port $(PORT)

# Show version
version:
	@echo "Version information:"
	$(GOCMD) run ./cmd/photo-uploader --version

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete"

# Lint the code
lint:
	@echo "Running linter..."
	golangci-lint run

# Format the code
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

# Vet the code
vet:
	@echo "Vetting code..."
	$(GOCMD) vet ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  all          - Clean, deps, test, and build"
	@echo "  build        - Build the binary"
	@echo "  build-all    - Build for all platforms"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  deps         - Download dependencies"
	@echo "  run          - Run the application"
	@echo "  run-port     - Run with custom port (use PORT=9090)"
	@echo "  version      - Show version information"
	@echo "  install      - Install the binary"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  help         - Show this help"
