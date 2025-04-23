# Flight Booking System

A modern API backend for airline flight booking management built with Go, Gin, and PostgreSQL.

## Project Overview

This system provides a RESTful API for managing flight bookings, with functionality to:

- Manage flights, planes, and airports
- Handle flight scheduling, including intermediate stops
- Manage ticket booking and seat selection

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

## Project Structure

```
.
├── docs/                   # Swagger documentation
├── internal/
│   ├── api/                # API layer (routes, handlers)
│   ├── dto/                # Data Transfer Objects
│   ├── exceptions/         # Error handling
│   ├── middleware/         # HTTP middleware
│   ├── models/             # Domain models
│   ├── repository/         # Data access layer
│   └── service/            # Business logic layer
├── pkg/
│   ├── config/             # Configuration management
│   ├── database/           # Database initialization and utilities
│   ├── logger/             # Logging utilities
│   ├── utils/              # Utility functions
│   └── validator/          # Request validation
├── docker-compose.yml      # Docker Compose configuration
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies checksum
├── main.go                 # Application entry point
└── README.md               # Project documentation
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.24 or later
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

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

```bash
go run main.go
```

The server will start at http://localhost:8080.

### API Documentation

The API documentation is available via Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## Configuration

The application is configured through `pkg/config/config.yml`:

## Development

### Database Seeding

The application automatically seeds the database with sample data when running in development mode. The seed data includes:

- Airports (multiple locations in Vietnam)
- Planes (various Airbus and Boeing models)
- Ticket classes (Economy and Business)
- Seat configurations
- System configuration settings

### Error Handling

The project uses a centralized error handling mechanism with custom error types defined in `internal/exceptions`. All API endpoints are wrapped with error middleware that properly formats error responses.
