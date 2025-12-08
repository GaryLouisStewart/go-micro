APP_NAME := go-micro
IMAGE_NAME := $(APP_NAME):latest
PORT := 8080
MAIN_PATH := .cmd/api

.PHONY: all build run test clean podman-build podman-run help


all: build


fmt:
				@echo "Formatting code..."
				$go fmt ./...

vet:
				@echo "Vetting code..."
				@go vet ./...

test:
				@echo "Running tests..."
				@go test -v ./...

build: fmt vet
				@echo "Building binary..."
				@go build -o bin/$(APP_NAME) $(MAIN_PATH)

run: build
				@echo "Building Container image..."
				@podman build -t $(IMAGE_NAME) $(MAIN_PATH)

podman-run:
				@echo "Starting Container on port $(PORT)..."
				@podman run -p $(PORT):$(PORT) --rm --name $(APP_NAME) $(IMAGE_NAME)

podman-stop:
				@echo "Stopping Container..."
				@-podman stop $(APP_NAME)

clean:
				@echo "Cleaning up..."
				@rm -rf bin
				@go clean


help:
				@echo "Usage: make [target]"
				@echo ""
				@echo "	build					 Format, vet and build the binary"
				@echo "	run						 Build and run the application locally"
				@echo "	test					 Run Go tests"
				@echo "	podman-build	 Build the Container image"
				@echo "	podman-run		 Run the application using podman"
				@echo "	clean					 Remove build binaries"

