# Define Go-related variables
GO = go
GOFMT = $(GO) fmt
GOBUILD = $(GO) build
GOTEST = $(GO) test
GOGET = $(GO) get
GOINSTALL = $(GO) install

# Define project binary name
BINARY_NAME = alshf

# Install dependencies
deps:
    $(GO) mod tidy

# Build the project
build: deps
    $(GOBUILD) -o $(BINARY_NAME) cmd/alshf/main.go

# Install the project binary
install: build
    $(GOINSTALL)

# Run tests (if you have any)
test:
    $(GOTEST) ./...

# Run the application (start the service)
run: build
    ./$(BINARY_NAME)

# Clean generated binaries
clean:
    rm -f $(BINARY_NAME)

.PHONY: deps build install test run clean