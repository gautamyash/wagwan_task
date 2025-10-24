# Event Guest Manager - Setup Guide

## Prerequisites

- Go 1.21+
- Node.js 18+
- Docker Desktop
- Git

## Quick Setup

### 1. Start Database
```bash
docker-compose up -d
# Wait 10 seconds for initialization
```

### 2. Start Backend
```bash
cd backend
go mod download
go run main.go
# Runs on http://localhost:8080
```

### 3. Start Frontend
```bash
cd frontend
npm install
npm run dev
# Runs on http://localhost:5173
```

### 4. Verify
Open http://localhost:5173 - you should see a guest list with sample data.

## Tech Stack

**Backend**: Go + PostgreSQL  
**Frontend**: Svelte 5 + TypeScript + TailwindCSS  
**Database**: PostgreSQL 16 (Docker)

## API Endpoints

The backend provides these endpoints:

- `GET /api/guests` - Get all guests (supports ?status= filter)
- `GET /api/guests/:id` - Get single guest
- `POST /api/guests` - Create new guest
- `DELETE /api/guests/:id` - Delete guest

## Troubleshooting

**Database won't start:**
```bash
docker-compose down
docker-compose up -d
```

**Backend can't connect:**
- Make sure Docker is running
- Wait 10-15 seconds after starting docker-compose

**Frontend won't start:**
```bash
rm -rf node_modules .svelte-kit
npm install
npm run dev
```

**Port conflicts:**
- Backend: 8080
- Frontend: 5173
- Database: 5432

Change ports in respective config files if needed.

## Project Structure

```
backend/
  ├── main.go              # Server entry point
  ├── handlers/
  │   └── guests.go        # API handlers
  └── db/
      ├── init.sql         # Database schema
      └── queries.go       # Database functions

frontend/
  ├── src/
  │   ├── lib/
  │   │   ├── api.ts       # API client
  │   │   └── components/  # Reusable components
  │   └── routes/
  │       └── +page.svelte # Main guest list page
  └── package.json
```

## Need Help?

Check existing code for patterns and examples. The application is functional but may have some issues - part of your task is to identify and fix any problems you encounter.

