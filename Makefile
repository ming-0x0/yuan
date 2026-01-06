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

# Allow Makefile to accept arguments directly
%:
	@:
