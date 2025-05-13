#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status messages
print_status() {
    echo -e "${GREEN}[✓] $1${NC}"
}

print_error() {
    echo -e "${RED}[✗] $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}[!] $1${NC}"
}

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    print_error "Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    print_error "Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go first."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
REQUIRED_VERSION="1.23.5"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    print_error "Go version $REQUIRED_VERSION or higher is required. Current version: $GO_VERSION"
    exit 1
fi

# Start services using docker-compose
print_status "Starting services using Docker Compose..."
docker-compose -f docker/docker-compose.yaml up -d postgres

# Wait for PostgreSQL to be ready
print_status "Waiting for PostgreSQL to be ready..."
sleep 5

# Install required Go tools
print_status "Installing required Go tools..."
go install github.com/air-verse/air@latest
curl -sSf https://atlasgo.sh | sh
go install github.com/swaggo/swag/cmd/swag@latest
go install honnef.co/go/tools/cmd/staticcheck@latest

# Install golangci-lint based on OS
if [[ "$OSTYPE" == "darwin"* ]]; then
    brew install golangci-lint
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
fi

# Download Go dependencies
print_status "Downloading Go dependencies..."
go mod download

# Create .env file if it doesn't exist
if [ ! -f .env ]; then
    print_status "Creating .env file..."
    cat > .env << EOL
ENV=development
DATABASE_URL=postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable
PORT=8080
LOG_LEVEL=debug
EOL
    print_status ".env file created successfully"
else
    print_warning ".env file already exists, skipping creation"
fi

# Run database migrations
print_status "Running database migrations..."
make mig-apply

# Generate API documentation
print_status "Generating API documentation..."
make docs

# Build the application
print_status "Building the application..."
make build

print_status "Setup completed successfully!"
print_status "You can now run the application using:"
echo -e "  ${YELLOW}make run${NC} - for standard run"
echo -e "  ${YELLOW}make dev${NC} - for development mode with hot reload"
echo -e "\nTo stop the services, run:"
echo -e "  ${YELLOW}docker-compose -f docker/docker-compose.yaml down${NC}"
