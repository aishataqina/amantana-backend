.PHONY: all build run clean test coverage lint mock swagger help migrate seed

# Variables
BINARY_NAME=amantana
MAIN_FILE=cmd/api/main.go
COVERAGE_FILE=coverage.out

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-w -s"

all: clean build

build:
	$(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME) $(MAIN_FILE)

run:
	$(GORUN) $(MAIN_FILE)

clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f $(COVERAGE_FILE)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=$(COVERAGE_FILE) ./...
	$(GOCMD) tool cover -html=$(COVERAGE_FILE)

lint:
	golangci-lint run

deps:
	$(GOGET) -u ./...
	$(GOMOD) tidy

mock:
	mockgen -source=internal/domain/repository.go -destination=internal/mocks/repository_mock.go

swagger:
	swag init -g $(MAIN_FILE) -o docs

air:
	air

docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p :8080 $(BINARY_NAME)

migrate-up:
	migrate -path database/migration -database "postgresql://postgres:postgres@localhost:5432/amantana_db?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration -database "postgresql://postgres:postgres@localhost:5432/amantana_db?sslmode=disable" -verbose down

migrate:
	$(GORUN) cmd/migrate/main.go migrate

seed:
	$(GORUN) cmd/migrate/main.go seed

help:
	@echo "Make commands:"
	@echo "  all        - Clean and build"
	@echo "  build      - Build the binary"
	@echo "  run        - Run the application"
	@echo "  clean      - Clean build files"
	@echo "  test       - Run tests"
	@echo "  coverage   - Generate test coverage"
	@echo "  lint       - Run linter"
	@echo "  deps       - Update dependencies"
	@echo "  mock       - Generate mock files"
	@echo "  swagger    - Generate swagger docs"
	@echo "  air        - Run with hot reload"
	@echo "  docker-build - Build docker image"
	@echo "  docker-run   - Run docker container"
	@echo "  migrate-up   - Run database migrations up"
	@echo "  migrate-down - Run database migrations down"
	@echo "  migrate    - Run database migrations"
	@echo "  seed       - Run database seeder" 