# Antrego

Queue management system with a Go backend and web frontend.

## Tech Stack

- **Backend:** Go, Gin, GORM, SQLite
- **API Docs:** Swagger/OpenAPI 2.0

## Getting Started

### Prerequisites

- Go 1.26+
- Node.js (for frontend)

### Backend

```bash
# Copy environment config
cp .env.example .env  # if available, otherwise edit .env directly

# Run the server
go run .
```

The server starts on `http://localhost:8080`.

### Database

Migrations are in `db/migrations/`. The database is SQLite, configured via `.env`:

```
GOOSE_DRIVER=sqlite3
GOOSE_DBSTRING=./db/antrego.db
GOOSE_MIGRATION_DIR=./db/migrations
```

## API

| Method | Path                           | Description              |
|--------|--------------------------------|--------------------------|
| GET    | `/queues`                      | List all queues          |
| POST   | `/queue`                       | Book a queue             |
| GET    | `/queue/:bookingCode`          | Get queue by code        |
| PATCH  | `/queue/:bookingCode/call`     | Mark queue as called     |
| PATCH  | `/queue/:bookingCode/complete` | Mark queue as done       |
| PATCH  | `/queue/:bookingCode/cancel`   | Cancel a queue           |
| GET    | `/swagger/*any`                | Swagger UI               |

### Swagger Docs

Generate or regenerate the spec:

```bash
swag init -g main.go --output docs
```

View the interactive docs at `http://localhost:8080/swagger/index.html` while running.

## Project Structure

```
├── main.go              # Entry point, routes
├── handlers/            # HTTP handlers
├── services/            # Business logic
├── models/              # GORM models
├── dto/                 # Request/response DTOs
├── middleware/          # Gin middleware (error handler)
├── config/              # Database config
├── docs/                # Generated swagger docs
├── db/                  # SQLite DB + migrations
└── apps/web/            # React frontend
```
