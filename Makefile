.PHONY: help build test lint run clean docker-build docker-run migrate auto_commit

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building application..."
	go build -o bin/hermes ./cmd/web

test: ## Run tests
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Run tests with coverage report
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linters
	@echo "Running linters..."
	golangci-lint run

run: ## Run the application
	@echo "Running application..."
	go run ./cmd/web/main.go

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -f deploy/docker/Dockerfile -t hermes:latest .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -d \
		--name hermes \
		-p 8080:8080 \
		-p 9090:9090 \
		-v $(PWD)/.config.yaml:/app/.config.yaml \
		hermes:latest

docker-stop: ## Stop Docker container
	@echo "Stopping Docker container..."
	docker stop hermes || true
	docker rm hermes || true

migrate: ## Run database migrations (manual)
	@echo "Running migrations..."
	@echo "Migrations are automatically run on startup"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

proto: ## Generate gRPC code from proto files
	@echo "Generating gRPC code..."
	@if command -v protoc > /dev/null; then \
		protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			api/grpc/*.proto; \
	else \
		echo "protoc not found. Install it first."; \
	fi

auto_commit: ## Auto commit: git pull, commit with timestamp, and push
	@echo "Pulling latest changes..."
	@git pull
	@echo "Staging all changes..."
	@git add -A
	@echo "Committing with timestamp..."
	@TIMESTAMP=$$(date +"%Y-%m-%d %H:%M:%S %z"); \
	git commit -m "Auto commit: $$TIMESTAMP" || echo "No changes to commit"
	@echo "Pushing changes..."
	@git push
	@echo "Auto commit completed!"
