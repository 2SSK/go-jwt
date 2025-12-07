# JWT Backend

A production-ready Go backend JWT using Echo framework, PostgreSQL, and comprehensive observability features.

## Features

- **Framework**: Echo v4 for high-performance HTTP routing
- **Database**: PostgreSQL with connection pooling and migrations
- **Observability**: Structured logging with Zerolog, health checks, and metrics
- **Configuration**: Environment-based config with validation
- **Middleware**: CORS, rate limiting, request logging, error handling
- **API Documentation**: OpenAPI/Swagger UI
- **Development Tools**: Hot reload, linting, testing setup

## Prerequisites

- Go 1.21+
- PostgreSQL 12+
- Docker (optional, for local development)

## Quick Start

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd jwt
   ```

2. **Set up environment variables**

   ```bash
   cp .env.sample .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**

   ```bash
   go mod download
   ```

4. **Set up database**

   ```bash
   # Ensure PostgreSQL is running
    # Run migrations (only in non-local environments)
    go-task migrations:up
   ```

5. **Run the application**
   ```bash
   go-task run
   ```

The server will start on `http://localhost:8080`.

## Configuration

Configuration is managed through environment variables with the `AUTH_` prefix. Key settings:

- **Server**: Port, timeouts, CORS origins
- **Database**: Connection details, pooling settings
- **Observability**: Logging level, service name, health checks

See `.env.sample` for all available options.

## API Endpoints

- `GET /status` - Health check endpoint
- `GET /docs` - OpenAPI documentation UI
- `GET /static/*` - Static file serving

## Development

### Available Tasks

- `go-task run` - Start the development server
- `go-task build` - Build the application
- `go-task test` - Run tests
- `go-task lint` - Run linters
- `go-task migrations:up` - Run database migrations

### Project Structure

```
├── cmd/
│   └── jwt/
│       └── main.go  # Application entry point
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── observability.go  # Configuration management
│   ├── database/
│   │   ├── migrations/
│   │   │   ├── 001_setup.sql
│   │   │   ├── 002_user.sql
│   │   │   └── 003_create_users.sql
│   │   ├── database.go
│   │   └── migrator.go  # Database setup and migrations
│   ├── errs/
│   │   ├── http.go
│   │   └── types.go  # Custom error types
│   ├── handler/
│   │   ├── auth.go
│   │   ├── base.go
│   │   ├── handlers.go
│   │   ├── health.go
│   │   ├── home.go
│   │   ├── openapi.go
│   │   └── user.go  # HTTP handlers
│   ├── lib/
│   │   ├── authhelper.go
│   │   └── utils.go  # Utility functions
│   ├── logger/
│   │   └── logger.go  # Logging setup
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── context.go
│   │   ├── global.go
│   │   ├── middlewares.go
│   │   ├── rate_limit.go
│   │   └── request_id.go  # Echo middlewares
│   ├── model/
│   │   ├── user/
│   │   │   ├── dto.go
│   │   │   └── user.go
│   │   └── model.go  # Data models
│   ├── repository/
│   │   ├── repositories.go
│   │   └── user.go  # Data access layer
│   ├── router/
│   │   ├── v1/
│   │   │   ├── auth.go
│   │   │   ├── user.go
│   │   │   └── v1.go
│   │   ├── router.go
│   │   └── system.go  # Route definitions
│   ├── server/
│   │   └── server.go  # Server setup
│   ├── service/
│   │   ├── auth.go
│   │   ├── services.go
│   │   └── user.go  # Business logic
│   ├── sqlerr/
│   │   ├── error.go
│   │   └── handler.go  # SQL error handling
│   └── validation/
│       └── utils.go  # Input validation
├── static/
│   ├── openapi.html
│   └── openapi.json  # Static assets (OpenAPI docs)
├── .env.sample  # Environment template
├── .gitignore
├── .golangci.yml
├── go.dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── Taskfile.yml  # Development tasks
```

## Observability

### Logging

- Structured JSON logs in production
- Console-formatted logs in development
- Request/response logging with latency and status
- Error stack traces

### Health Checks

- Database connectivity checks
- Configurable check intervals and timeouts

## Contributing

1. Follow Go best practices and project conventions
2. Add tests for new features
3. Update documentation as needed
4. Run `go-task lint` before committing
