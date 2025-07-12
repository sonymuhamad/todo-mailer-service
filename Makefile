# Variables
APP_NAME := todo-mailer-service
GIT_COMMIT_HASH := $(shell git log --format="%H" -n 1)
IMAGE_NAME := sonymuhamad/$(APP_NAME)
IMAGE_TAG := $(GIT_COMMIT_HASH)

.PHONY: build docker-push

# Build Docker image with commit hash tag
build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Push to Docker Hub or any registry
docker-push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
