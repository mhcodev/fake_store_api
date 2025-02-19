# Check if .env exists, otherwise stop execution
ifneq ("$(wildcard .env)", "")
    include .env
else
    $(error ".env file not found. Aborting.")
endif

# Project Variables
APP_NAME=fake-store-api
DOCKER_COMPOSE=docker compose
GO=go
PORT=4000
DB_CONTAINER=postgres-db


# Build and Run the API
.PHONY: docker-build docker-run docker-stop migrate

## Build the Docker image
docker-build:
	@echo "Creating Prometheus credentials..."
	@bash -c 'htpasswd -cb ./prometheus/.htpasswd "$(PROMETHEUS_USERNAME)" "$(PROMETHEUS_PASSWORD)"'
	@echo "Creating Docker images..."
	$(DOCKER_COMPOSE) build

## Run the containerized application
docker-run:
	$(DOCKER_COMPOSE) up -d

## Stop the container
docker-stop:
	$(DOCKER_COMPOSE) down