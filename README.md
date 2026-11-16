# Kilvia

Kilvia is a logistics coordination platform with an Angular web client and a Go REST API backed by PostgreSQL. The `mobile/` directory contains the Flutter application.

## Local development with Docker

Requirements: Docker Desktop with the Compose plugin.

1. Copy `.env.example` to `.env`.
2. Replace `JWT_SECRET` and the database password with local values.
3. Start the stack:

   ```powershell
   docker compose up --build
   ```

The web app is available at `http://localhost:4200`, the API at `http://localhost:8080`, and the health check at `http://localhost:8080/health`. Compose waits for PostgreSQL, applies `backend/migrations`, then starts the API. Register the first company and administrator from the web app.

The credentials in `.env.example` are for local development only. Do not use them in a deployed environment.

## Running the Angular client separately

Install dependencies and run the development server:

```powershell
npm ci
npm start
```

The Angular dev-server proxy forwards `/api` requests to `http://localhost:8080`. Start PostgreSQL and the API separately before signing in. The API requires `DATABASE_URL` and a `JWT_SECRET` of at least 32 characters; `FRONTEND_ORIGIN` defaults to `http://localhost:4200`.

## API

All application endpoints are versioned under `/api/v1`. Authentication endpoints:

- `POST /api/v1/auth/register` creates a company and its first `ADMIN`.
- `POST /api/v1/auth/login` authenticates a user.
- `GET /api/v1/auth/me` returns the authenticated user.

The API also exposes company, user, driver, vehicle, trip, load, and match read endpoints. Protected endpoints require a bearer token. Role permissions and company isolation are enforced by the backend, not only by the Angular route guards.

## Demo data

Opportunities, matches, and the dashboard summary continue to use Sprint 1 demo data. Company, user, driver, vehicle, trip, and load management screens use the API.

## Tests and builds

```powershell
npm run build
```

With Go 1.24 installed:

```powershell
cd backend
go test ./...
```
