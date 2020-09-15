.DEFAULT_GOAL         := help
REPO                  := github.com/bukhavtsov/museum/
BIN_PATH              ?= ./bin
BACKEND_IMAGE_NAME     ?= backend:latest
BACKEND_CONTAINER_NAME ?= backend
BACKEND_SRC_PATH       ?= ./back-end/cmd/
BACKEND_BIN_PATH       ?= $(BIN_PATH)/backend/backend
BACKEND_DOCKER_PATH    ?= ./docker/backend

DB_IMAGE_NAME     ?= db-museum:latest
DB_CONTAINER_NAME ?= db-museum
DB_SRC_PATH       ?= ./docker/db-museum
DB_DOCKER_PATH    ?= $(DB_SRC_PATH)/Dockerfile

PHONY: help
help: ## makefile targets description
	@echo "Usage:"
	@egrep '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##/#-/' | column -t -s "#"

.PHONY: fmt
fmt: ## automatically formats Go source code
	@echo "Running 'go fmt ...'"
	@go fmt -x "$(REPO)/..."


.PHONY: build
build: fmt ## compile package and dependencies
	@echo "Building backend..."
	CGO_ENABLED=0 go build -o $(BACKEND_BIN_PATH) $(BACKEND_SRC_PATH)

.PHONY: run
run: build ## execute back-end binary
	@echo "Running server..."
	$(BACKEND_BIN_PATH)


.PHONY: image
image: build ## build images from Dockerfile ./docker/back-end/Dockerfile and ./docker/db-museum/Dockerfile
	@echo "Building back-end image..."
	cp $(BACKEND_BIN_PATH) $(BACKEND_DOCKER_PATH)
	@docker build -t $(BACKEND_IMAGE_NAME) $(BACKEND_DOCKER_PATH)
	rm $(BACKEND_DOCKER_PATH)/backend
	@echo "Building db-museum image..."
	@docker build -f $(DB_DOCKER_PATH) -t $(DB_IMAGE_NAME) .