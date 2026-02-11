.PHONY: proto-gen iam-proto blog-proto
# Generate all proto files
proto-gen: iam-proto blog-proto

# Generate IAM service proto files
iam-proto:
	@echo "Generating IAM service proto files..."
	@mkdir -p pkg/proto/iam/v1
	@protoc --go_out=./ \
		--go-grpc_out=./ \
		--grpc-gateway_out=./ \
		--grpc-gateway_opt=paths=source_relative \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		-I ./ \
		-I ./google \
		pkg/proto/iam/auth.proto \
		pkg/proto/iam/user.proto

# Generate Blog service proto files
blog-proto:
	@echo "Generating Blog service proto files..."
	@mkdir -p pkg/proto/blog/v1
	@protoc --go_out=./ \
		--go-grpc_out=./ \
		--grpc-gateway_out=./ \
		--grpc-gateway_opt=paths=source_relative \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		-I ./ \
		-I ./google \
		pkg/proto/blog/blog.proto

.PHONY: goose-up goose-down-to goose-create goose-down
goose-up:
	@echo "Applying migrations..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh up

goose-down-to:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Error: Please provide the migration version (e.g., make goose-down-to 00001)"; \
		exit 1; \
	fi
	@echo "Rolling back to migration $(filter-out $@,$(MAKECMDGOALS))..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh down-to $(filter-out $@,$(MAKECMDGOALS))

goose-create:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Error: Please provide the migration name (e.g., make goose-create my_migration)"; \
		exit 1; \
	fi
	@echo "Creating migration $(filter-out $@,$(MAKECMDGOALS))..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh create $(filter-out $@,$(MAKECMDGOALS))

goose-down:
	@echo "Rolling the last migration..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh down

.PHONY: test coverage
# Run tests
test:
	@go test -v ./...

# Run tests with coverage
coverage:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: lint lint-fix
# Run linter
lint:
	@golangci-lint run ./...

# Run linter and fix issues
lint-fix:
	@golangci-lint run --fix ./...

.PHONY: build
# Build project
build:
	@go build -v ./...

.PHONY: docker-local-up docker-local-down docker-local-rebuild
# Docker Compose commands for local environment
docker-local-up:
	@echo "Starting local environment..."
	@cd deployments/compose && docker-compose -f docker-compose.local.yaml --env-file .env.local up -d

docker-local-down:
	@echo "Stopping local environment..."
	@cd deployments/compose && docker-compose -f docker-compose.local.yaml down

docker-local-rebuild:
	@echo "Rebuilding and starting local environment..."
	@cd deployments/compose && docker-compose -f docker-compose.local.yaml --env-file .env.local up -d --build

# Allow Makefile to accept arguments directly
%:
	@:
