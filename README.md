# Flight Booking System

A modern API backend for airline flight booking management built with Go, Gin, and PostgreSQL.

## Project Overview

This system provides a RESTful API for managing flight bookings, with functionality to:

- Manage flights, planes, and airports
- Handle flight scheduling, including intermediate stops
- Manage ticket booking and seat selection
- Support multiple ticket classes and seat configurations
- Handle user authentication and authorization
- Generate flight codes automatically
- Support intermediate stops with duration and order
- Calculate ticket prices based on seat class
- Track seat availability and booking status

The project follows a clean architecture pattern with clear separation of concerns:

- **API Layer**: HTTP request/response handling using Gin
- **Service Layer**: Business logic implementation
- **Repository Layer**: Data access logic
- **Models**: Domain entities
- **DTOs**: Data transfer objects for API requests/responses

## Tech Stack

- **Backend**: Go 1.24+
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Configuration**: YAML
- **Logging**: Zap
- **API Documentation**: Swagger
- **Containerization**: Docker and Docker Compose
- **Hot Reload**: Air

## Project Structure

```
.
├── docs/                   # Swagger documentation
├── internal/
│   ├── api/               # API layer (routes, handlers)
│   ├── dto/               # Data Transfer Objects
│   ├── exceptions/        # Error handling
│   ├── middleware/        # HTTP middleware
│   ├── models/            # Domain models
│   ├── repository/        # Data access layer
│   └── service/           # Business logic layer
├── pkg/
│   ├── auth/              # Authentication utilities
│   ├── config/            # Configuration management
│   ├── database/          # Database initialization and utilities
│   ├── init/              # Application initialization
│   ├── logger/            # Logging utilities
│   ├── utils/             # Utility functions
│   └── validator/         # Request validation
├── tests/                 # Test files
├── logs/                  # Application logs
├── compose.yml            # Docker Compose configuration
├── Dockerfile             # Docker configuration
├── go.mod                 # Go module definition
├── go.sum                 # Go dependencies checksum
├── main.go                # Application entry point
├── run.bat                # Windows run script
├── run.sh                 # Unix run script
├── .air.toml              # Air hot reload configuration
└── README.md              # Project documentation
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.24 or later
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Air](https://github.com/cosmtrek/air) for hot reload (optional)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/aprilboiz/FlightBookingSite.git
cd FlightBookingSite
```

2. Install dependencies:

```bash
go mod tidy
```

3. Start the PostgreSQL database:

```bash
docker-compose up -d
```

4. Run the application:

On Windows:
```bash
run.bat
```

On Unix:
```bash
./run.sh
```

Or with hot reload:
```bash
air
```

The server will start at http://localhost:8080.

### API Documentation

The API documentation is available via Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## Configuration

The application is configured through `pkg/config/config.yml`. Key configuration options include:

- Database connection settings
- Server port and host
- Logging configuration

## Development

### Database Seeding

The application automatically seeds the database with sample data when running in development mode. The seed data includes:

- Airports (multiple locations in Vietnam)
- Planes (various Airbus and Boeing models)
- Ticket classes (Economy and Business)
- Seat configurations
- System configuration settings

### Testing

The project includes comprehensive test coverage:

- Unit tests for service and repository layers
- Integration tests using a test database
- Mock implementations for testing

Run tests with:
```bash
go test ./...
```

### Error Handling

The project uses a centralized error handling mechanism with custom error types defined in `internal/exceptions`. All API endpoints are wrapped with error middleware that properly formats error responses.

### Hot Reload

The project uses Air for hot reloading during development. Configuration is in `.air.toml`. Start the application with hot reload:

```bash
air
```
