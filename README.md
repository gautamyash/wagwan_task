# Wagwan RSVP Project

This repository contains my completed submission for the **Wagwan World Full Stack Engineer assignment**.

## 🚀 Overview

- **Frontend**: Svelte 5 + TypeScript + TailwindCSS  
- **Backend**: Go + PostgreSQL  
- **Database**: PostgreSQL (via Docker)
- **Migrations**: manual SQL files via `psql`  

## 🎯 Features

- Public RSVP page (`/rsvp`)  
- Event information display  
- Fully functional RSVP form with validation  
- Connected to backend API  
- Added new fields: plus_ones, dietary_restrictions, rsvp_date  

## ⚙️ Quick Start

1. Backend (API)

   ```bash
   cd backend
   go run main.go
   ```

2. Frontend (Svelte dev server)

   ```bash
   cd frontend
   npm install
   npm run dev
   ```

3. Open
   - Public RSVP: <http://localhost:5173/rsvp>
   - Admin dashboard: <http://localhost:5173/>

## 🗂️ Migrations (manual via psql)

Apply SQL files in `backend/migrations/` using psql. Example for the latest additions:

```bash
psql -U <user> -d <database> -f backend/migrations/000004_add_plus_ones_and_dietary_restrictions.up.sql
```

Rollback example:

```bash
psql -U <user> -d <database> -f backend/migrations/000004_add_plus_ones_and_dietary_restrictions.down.sql
```

Notes:

- Ensure `events` uses columns `title` and `event_date` (aligned with migrations).
- Guests include `notes`, `rsvp_date`, `plus_ones`, `dietary_restrictions`.
- Frontend status values are normalized to backend `attending/pending/declined`.

## 🧠 Developer

**Name:** Yash Gautam  
**Email:** <gautamyash37@gmail.com>  
**Date:** October 29, 2025  

For details, see [SUBMISSION.md](./SUBMISSION.md).
