# Variables
APP_NAME := castingdoor
DOCKER_IMAGE := $(APP_NAME):latest
CONTAINER_NAME := $(APP_NAME)-container
PORT := 8080

# Default target
.PHONY: all
all: build

# Build Docker image
.PHONY: build
build:
	@echo "🔨 Building Docker image: $(DOCKER_IMAGE)"
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
.PHONY: run
run:
	@echo "🚀 Running container: $(CONTAINER_NAME)"
	docker run -d -p $(PORT):8080 --name $(CONTAINER_NAME) $(DOCKER_IMAGE)

# Stop container if running
.PHONY: stop
stop:
	@echo "🛑 Stopping container: $(CONTAINER_NAME)"
	-docker stop $(CONTAINER_NAME)

# Clean dangling images
.PHONY: clean
clean:
	@echo "🧹 Cleaning up..."
	docker rmi -f $(DOCKER_IMAGE) || true

# Rebuild image
.PHONY: rebuild
rebuild: clean build

# Tag image for release (example)
.PHONY: tag
tag:
	@echo "🏷️  Tagging image as latest and versioned"
	docker tag $(DOCKER_IMAGE) $(APP_NAME):v1.0.0
