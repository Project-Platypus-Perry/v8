# Variables
APP_NAME := my-backend
MAIN := cmd/server/main.go
SWAG_OUT := ./docs
GO_BIN := $(shell go env GOPATH)/bin

# Load env vars from .env file
include .env
export $(shell sed 's/=.*//' .env)

# Binary build
.PHONY: build
build:
	go build -o bin/$(APP_NAME) $(MAIN)

# Run app
.PHONY: run
run:
	go run $(MAIN)

# Run with hot reload using air
.PHONY: dev
dev:
	$(GO_BIN)/air

# Swagger generation
.PHONY: swagger
swagger:
	swag init --generalInfo $(MAIN) --output $(SWAG_OUT)

# Linting (you can use golangci-lint or go vet as fallback)
.PHONY: lint
lint:
	golangci-lint run || true

# Tests
.PHONY: test
test:
	go test ./... -v -cover

# Formatting
.PHONY: fmt
fmt:
	go fmt ./...
	goimports -w .

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf bin/

.PHONY: show-env
show-env:
	@echo $(ENV)
	@echo $(DATABASE_URL)

# Create migrations with name
.PHONY: mig-diff 
mig-diff:
	@atlas migrate diff $(name) --env gorm -c file://internal/atlas.hcl

# Create new empty migrations with name
.PHONY: mig-new
mig-new:
	@atlas migrate new $(name) --env gorm -c file://internal/atlas.hcl

# Apply migrations
.PHONY: mig-apply
mig-apply:
	@atlas migrate apply --env gorm -c file://internal/atlas.hcl

# Status of migrations
.PHONY: mig-status
mig-status:
	@atlas migrate status --env gorm -c file://internal/atlas.hcl

# Rollback migrations
.PHONY: mig-rollback
mig-rollback:
	@atlas migrate down --env gorm -c file://internal/atlas.hcl 

# Validate migrations
.PHONY: mig-validate
mig-validate:
	@atlas migrate validate --env gorm -c file://internal/atlas.hcl

# List migrations
.PHONY: mig-list
mig-list:
	@atlas migrate ls --env gorm -c file://internal/atlas.hcl

# Hash migrations
mig-hash:
	@atlas migrate hash --env gorm -c file://internal/atlas.hcl


