.PHONY: all build run clean db

# Variables
BINARY_NAME=salesforge-api
CONFIG_FILE=config/config.yaml
DB_CONTAINER_NAME=postgres
TEST_DB_CONTAINER_NAME=postgres-test

# Default target
all: build run

# Build the project
build:
	cd cmd/salesforge-api && go build -o ../../$(BINARY_NAME)

# Run the project
run: db
	./$(BINARY_NAME) -config $(CONFIG_FILE)

# Clean the build
clean:
	rm -f $(BINARY_NAME)

# Start the database
db:
	docker-compose up -d postgres
	@echo "Waiting for database to be ready..."
	@while ! docker exec $(DB_CONTAINER_NAME) pg_isready -U postgres; do sleep 1; done
	@echo "Database is ready."

# Run tests
test: testdb
	@echo "Running tests..."
	go test ./...

# Start the test database
testdb:
	docker-compose up -d testdb
	@echo "Waiting for test database to be ready..."
	@while ! docker exec $(TEST_DB_CONTAINER_NAME) pg_isready -U postgres; do sleep 1; done
	@echo "Test database is ready."
