# Set the default shell to bash
SHELL := /bin/bash

# Set the Go binary name
BINARY := mygpt

# Set the version number
VERSION := 1.0.0

# Set the Go module name
MODULE := github.com/myusername/myapp

# Set the Go flags
GOFLAGS := -ldflags="-X $(MODULE)/version.Version=$(VERSION)"

.PHONY: build clean

build:
	@echo "Building $(BINARY) $(VERSION)..."
	go build $(GOFLAGS) -o $(BINARY)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY)