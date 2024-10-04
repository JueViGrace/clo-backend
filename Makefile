# Build the application
all: build test

build:
	@echo "Building ..."
	
	@go build -o ./bin/main ./cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

run-prod:
	./bin/main

# Create DB container
docker-run:
	@if docker compose up -d 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker compose up -d; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integrations Test for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	rm -f ./bin/main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

sqlc:
	@sqlc generate

migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_MIGRATION_DIR=./sql/schema GOOSE_DBSTRING="host=localhost port=5432 user=jvg25 password=root dbname=bakery_db sslmode=disable search_path=public" goose up  

migrate-down:
	@GOOSE_DRIVER=postgres GOOSE_MIGRATION_DIR=./sql/schema GOOSE_DBSTRING="host=localhost port=5432 user=jvg25 password=root dbname=bakery_db sslmode=disable search_path=public" goose down

.PHONY: all build run test clean watch docker-run docker-down sqlc migrate-up migrate-down itest