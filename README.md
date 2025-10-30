# Wagwan RSVP Project

Production-ready RSVP system with a public RSVP page and an admin list. Frontend is Svelte 5 + TailwindCSS, backend is Go with PostgreSQL. Database changes are managed via manual SQL migrations executed with `psql`.

## Tech Stack

- Frontend: Svelte 5, TailwindCSS, Vite, TypeScript
- Backend: Go, Gorilla Mux, CORS
- Database: PostgreSQL 16 (Docker), initialized by `backend/db/init.sql`
- Migrations: manual SQL files in `backend/migrations/` (applied with `psql`)

## Prerequisites

- Node.js 18+ and npm
- Go 1.21+
- Docker + Docker Compose (recommended) or a local Postgres instance
- `psql` CLI in PATH

## 1) Start PostgreSQL

Option A – Docker (recommended)

```bash
docker compose up -d
```

This starts Postgres on port 5432 and auto-seeds schema/data from `backend/db/init.sql`.

Option B – Local Postgres
Create a DB named `eventguests` and user `postgres/postgres` or adjust env vars when running the API.

## 2) Apply Migrations (manual via psql)

Run these in order to ensure your schema matches the code:

```bash
psql -U postgres -d eventguests -f backend/migrations/000001_create_guests_table.up.sql
psql -U postgres -d eventguests -f backend/migrations/000002_add_events_and_guest_notes.up.sql
psql -U postgres -d eventguests -f backend/migrations/000003_add_notes_and_eventid_to_guests.up.sql
psql -U postgres -d eventguests -f backend/migrations/000004_add_plus_ones_and_dietary_restrictions.up.sql
```

Rollback example (latest):

```bash
psql -U postgres -d eventguests -f backend/migrations/000004_add_plus_ones_and_dietary_restrictions.down.sql
```

Notes

- `events` table columns: `title`, `event_date` (aligned with code and docs)
- `guests` table includes: `notes`, `rsvp_date`, `plus_ones`, `dietary_restrictions`, optional `event_id`

## 3) Run the Backend (API)

Default env (override as needed):

```
DB_HOST=localhost  DB_PORT=5432  DB_USER=postgres  DB_PASSWORD=postgres  DB_NAME=eventguests
```

Start the API:

```bash
cd backend
go run main.go
```

The API runs at <http://localhost:8080>. It also serves the static frontend from `../frontend/build`.

## 4) Run the Frontend

Dev mode (Vite, live reload):

```bash
cd frontend
npm install
npm run dev
```

Open the public RSVP at <http://localhost:5173/rsvp>.

Build for the Go server (port 8080):

```bash
cd frontend
npm run build
```

Reload <http://localhost:8080/rsvp>. Repeat `npm run build` whenever UI changes should appear on 8080.

## 5) RSVP via SvelteKit Server Actions

The public RSVP page (`/rsvp`) posts using a SvelteKit server action which proxies to the Go API.

Verify:

- In DevTools → Network, submitting shows a POST to `/rsvp` (same origin), not directly to `http://localhost:8080/api/guests`.
- With JavaScript disabled, submission still works and shows success/error.
- Data appears on the admin page (`/`) and via API `GET /api/guests`.

## Endpoints

- GET `/api/guests?status=pending|attending|declined`
- GET `/api/guests/{id}`
- POST `/api/guests`
- DELETE `/api/guests/{id}`

POST body example:

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "phone": "+1-555-0101",
  "status": "attending",
  "notes": "Vegetarian",
  "plus_ones": 1,
  "dietary_restrictions": "Veg",
  "event_id": 1
}
```

## What to Review (mapped to TASK.md)

- Public RSVP page at `/rsvp`: validation, success/error messages, clear-on-success, Tailwind styling, mobile-first
- Database migration system: SQL files for notes, events + `event_id`, `rsvp_date`/`plus_ones`/`dietary_restrictions` with up/down
- Backend updates: handlers and queries accept/return new fields; status uses `pending|attending|declined`
- Admin page (`/`): can view, filter by status, and delete guests

## Troubleshooting

- 5173 shows latest UI but 8080 doesn’t: run `npm run build` again, restart the Go server, hard-refresh 8080 (Ctrl+Shift+R). Ensure it serves `frontend/build`.
- DB connection errors: confirm Docker is up and `DB_*` envs match.

## Developer

Name: Yash Gautam  
Email: <gautamyash37@gmail.com>  
Date: October 29, 2025  

For full context and decisions, see [SUBMISSION.md](./SUBMISSION.md).
