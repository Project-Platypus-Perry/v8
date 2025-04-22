# Go Backend Service

A production-ready Go backend service using Echo framework, GORM, and Atlas for database migrations.

## Tech Stack

- **Framework**: [Echo](https://echo.labstack.com/) - High performance, extensible, minimalist Go web framework
- **ORM**: [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- **Database**: PostgreSQL
- **Migration**: [Atlas](https://atlasgo.io/) - Database schema migration tool
- **Logging**: [Zap](https://github.com/uber-go/zap) - Blazing fast, structured logging
- **Validation**: go-playground/validator - Request validation using struct tags
- **Hot Reload**: Air - Live reload for Go apps

## Prerequisites

- Go 1.23.5 or higher
- PostgreSQL
- Make

## Project Setup

1. Clone the repository:

```bash
git clone <repository-url>
cd <project-directory>
```

2. Install dependencies:

```bash
go mod download
```

3. Install development tools:

```bash
# Install air for hot reloading
go install github.com/air-verse/air@latest

# Install Atlas CLI for migrations
curl -sSf https://atlasgo.sh | sh
```

4. Create `.env` file:

```bash
cp .env.example .env
```

Update the following variables in `.env`:

```env
ENV=development
DATABASE_URL=postgres://postgres:postgres@localhost:5432/your_db_name?sslmode=disable
PORT=8080
LOG_LEVEL=debug
```

## Running the Service

There are several ways to run the service:

1. **Standard Run**:

```bash
make run
```

2. **Development Mode with Environment Variables**:

```bash
make dev
```

3. **Build Binary**:

```bash
make build
./bin/my-backend
```

## Database Migrations

This project uses Atlas for database migrations. Here are the common migration commands:

1. **Create a New Migration**:

```bash
# Create a new migration with a name
make mig-new name=add_users_table

# Create a migration based on schema changes
make mig-diff name=add_email_field
```

2. **Apply Migrations**:

```bash
# Apply all pending migrations
make mig-apply
```

3. **Check Migration Status**:

```bash
# View status of all migrations
make mig-status
```

4. **Rollback Migrations**:

```bash
# Rollback the last migration
make mig-rollback
```

5. **List All Migrations**:

```bash
make mig-list
```

6. **Validate Migrations**:

```bash
make mig-validate
```

### Code Structure

```
.
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── app/            # Application setup and DI
│   ├── config/         # Configuration
│   ├── constants/      # Constants and enums
│   ├── db/            # Database connection and migrations
│   ├── handler/       # HTTP handlers
│   ├── middleware/    # HTTP middleware
│   ├── model/        # Data models
│   ├── repository/   # Data access layer
│   ├── router/       # HTTP router setup
│   └── service/      # Business logic
└── pkg/
    └── logger/       # Logging utilities
```

### Making Changes

1. Create a new branch for your feature
2. Make your changes
3. Run tests: `make test`
4. Run linter: `make lint`
5. Format code: `make fmt`
6. Commit your changes
7. Create a pull request

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
