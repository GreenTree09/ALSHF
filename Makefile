# Define Go-related variables
GO = go
GOFMT = $(GO) fmt
GOBUILD = $(GO) build
GOTEST = $(GO) test
GOGET = $(GO) get
GOINSTALL = $(GO) install

# Define project binary name
BINARY_NAME = alshf

# Initialize go.mod file
init:
	$(GO) mod init github.com/GreenTree09/ALSHF

# Install dependencies
deps: init
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

.PHONY: init deps build install test run clean
