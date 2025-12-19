# Todo App

Simple todo application with Go backend, Svelte frontend, and PostgreSQL database.

## Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL 12+

## Setup

### Database

Create a PostgreSQL database:

```bash
createdb todos
```

Or using psql:

```bash
psql -U postgres -c "CREATE DATABASE todos;"
```

### Backend

```bash
cd backend
go mod download
go run main.go
```

The backend runs on port 8080 by default. Set environment variables if needed:

- `DB_HOST` (default: localhost)
- `DB_PORT` (default: 5432)
- `DB_USER` (default: postgres)
- `DB_PASSWORD` (default: postgres)
- `DB_NAME` (default: todos)
- `PORT` (default: 8080)

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend runs on port 5173 by default.

## Production Build

Build the frontend:

```bash
cd frontend
npm run build
```

The backend serves the built frontend from `/` and API endpoints from `/api/*`.

## Railway Deployment

1. Create a new project on Railway
2. Connect your GitHub repository
3. Add a PostgreSQL service
4. Railway automatically provides `DATABASE_URL` environment variable - no manual configuration needed
5. Deploy - Railway will build and deploy using the Dockerfile

Alternatively, you can set individual database variables:
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`

The Dockerfile builds both frontend and backend into a single container.

