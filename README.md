# Jwt Backend

A production-ready Go backend jwt using Echo framework, PostgreSQL, and comprehensive observability features.

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
   go-task migrate
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
- `go-task migrate` - Run database migrations

### Project Structure

```
├── cmd/jwt/     # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── database/        # Database setup and migrations
│   ├── handler/         # HTTP handlers
│   ├── logger/          # Logging setup
│   ├── middleware/      # Echo middlewares
│   ├── repository/      # Data access layer
│   ├── router/          # Route definitions
│   ├── server/          # Server setup
│   ├── service/         # Business logic
│   └── validation/      # Input validation
├── static/              # Static assets (OpenAPI docs)
├── .env.sample          # Environment template
└── Taskfile.yml         # Development tasks
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
