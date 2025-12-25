.PHONY: goose-up goose-down-to goose-create goose-down
goose-up:
	@if [ -z "$(word 1,$(filter-out $@,$(MAKECMDGOALS)))" ]; then \
		echo "Error: Please provide the database driver (e.g., make goose-up postgres)"; \
		exit 1; \
	fi
	@echo "Applying migrations for $(word 1,$(filter-out $@,$(MAKECMDGOALS)))..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh $(word 1,$(filter-out $@,$(MAKECMDGOALS))) up

goose-down-to:
	@if [ -z "$(word 1,$(filter-out $@,$(MAKECMDGOALS)))" ] || [ -z "$(word 2,$(filter-out $@,$(MAKECMDGOALS)))" ]; then \
		echo "Error: Please provide the database driver and version (e.g., make goose-down-to postgres 00001)"; \
		exit 1; \
	fi
	@echo "Rolling back $(word 1,$(filter-out $@,$(MAKECMDGOALS))) to migration $(word 2,$(filter-out $@,$(MAKECMDGOALS)))..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh $(word 1,$(filter-out $@,$(MAKECMDGOALS))) down-to $(word 2,$(filter-out $@,$(MAKECMDGOALS)))

goose-create:
	@if [ -z "$(word 1,$(filter-out $@,$(MAKECMDGOALS)))" ] || [ -z "$(word 2,$(filter-out $@,$(MAKECMDGOALS)))" ]; then \
		echo "Error: Please provide the database driver and migration name (e.g., make goose-create postgres my_migration)"; \
		exit 1; \
	fi
	@echo "Creating migration $(word 2,$(filter-out $@,$(MAKECMDGOALS))) for $(word 1,$(filter-out $@,$(MAKECMDGOALS)))..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh $(word 1,$(filter-out $@,$(MAKECMDGOALS))) create $(word 2,$(filter-out $@,$(MAKECMDGOALS)))

goose-down:
	@if [ -z "$(word 1,$(filter-out $@,$(MAKECMDGOALS)))" ]; then \
		echo "Error: Please provide the database driver (e.g., make goose-down postgres)"; \
		exit 1; \
	fi
	@echo "Rolling the last migration for $(word 1,$(filter-out $@,$(MAKECMDGOALS)))..."
	@chmod +x ./scripts/goose.sh
	@./scripts/goose.sh $(word 1,$(filter-out $@,$(MAKECMDGOALS))) down



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
