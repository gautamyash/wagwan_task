## Submission – Wagwan Event RSVP Project

**Candidate Name**: Yash Gautam  
**Date**: October 29, 2025  
**Email**: <gautamyash37@gmail.com>

---

## Time Spent

- RSVP Page Implementation: 9 hours  
- Database Migrations & Schema Fixes: 7 hours  
- Testing & Debugging: 3 hours  
- Documentation & Cleanup: 1 hour  
- **Approximate Total**: 20 hours

---

## What I Built

### Main Feature: Public RSVP Page (`/rsvp`)

I implemented a fully functional public RSVP page that allows users to register for an event, submit their attendance status, and include additional details like plus ones and dietary restrictions.

### Key Features

- Displays event details (title, date, time, location, description)
- Form fields for Name, Email, Phone, RSVP Status
- Added Plus Ones and Dietary Restrictions fields
- Client-side validation for required fields and email format
- Submits data to backend API (`POST /api/guests`)
- Shows success/error messages with smooth UX
- Clears form after successful submission
- Fully responsive and styled with TailwindCSS

### Design Choices

- Minimalist layout for clarity and accessibility
- TailwindCSS for consistent design and fast prototyping
- Simple animations and clear feedback on form submission
- Mobile-first approach for better usability on all devices

---

## Bugs Found & Fixed

### Bug 1: Events Schema Mismatch

- **Problem**: `backend/db/init.sql` created the `events` table using `name` and `date` instead of `title` and `event_date`.
- **Location**: `backend/db/init.sql`
- **Solution**: Updated `init.sql` to use `title` and `event_date` columns to match the `000002` migration.
- **Why**: To ensure consistency between schema and migrations.

### Bug 2: Missing RSVP Fields in Guests Table

- **Problem**: `guests` table lacked `plus_ones`, `dietary_restrictions`, and `rsvp_date` columns.
- **Solution**: Added a new migration file (`000004_add_plus_ones_and_dietary_restrictions.up.sql`) with:

```sql
ALTER TABLE guests 
ADD COLUMN IF NOT EXISTS plus_ones INT DEFAULT 0,
ADD COLUMN IF NOT EXISTS dietary_restrictions TEXT,
ADD COLUMN IF NOT EXISTS rsvp_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
```

- **Why**: To align backend with the new RSVP form requirements.

### Bug 3: API Handler Missing Fields

- **Problem**: `handlers/guests.go` only handled basic fields (name, email, phone, notes).
- **Solution**: Updated the handler to support new fields `plus_ones` and `dietary_restrictions`, including updates to the `CreateGuestRequest` struct and SQL insert query.
- **Why**: So the backend correctly accepts and stores all RSVP data.

### Bug 4: Inconsistent Status Mapping

- **Problem**: Frontend used `yes/maybe/no`, backend expected `attending/pending/declined`.
- **Solution**: Added a conversion map in the POST handler before inserting to the DB.
- **Why**: To ensure data consistency between frontend and backend.

---

## Challenges & Solutions

### Challenge 1: Database Schema Synchronization

- **Problem**: Existing migrations conflicted with `init.sql`, leading to mismatched tables.
- **Solution**: Aligned `init.sql` with the latest migration structure and applied new migrations manually using `psql` to ensure schema consistency.

### Challenge 2: Partial Implementation of RSVP Form

- **Problem**: The initial RSVP form missed `plus_ones` and `dietary_restrictions`.
- **Solution**: Extended the form, updated data binding, and modified the POST payload to include both fields.

### Challenge 3: Migration Execution Errors

- **Problem**: `go run main.go migrate` didn’t apply new columns.
- **Solution**: Executed the migration manually using `psql -f` and verified schema with `\d guests;`.

---

## How to Test My Work

### Setup

1) Clone the repository and navigate to the backend folder.

2) Run database migrations:

```bash
psql -U <user> -d <database> -f backend/migrations/000004_add_plus_ones_and_dietary_restrictions.up.sql
```

3) Start the server:

```bash
cd backend
go run main.go
```

4) Start the frontend (Svelte app):

```bash
cd frontend
npm install
npm run dev
```

### Testing the RSVP Page

1) Visit `http://localhost:5173/rsvp`.
2) Fill in all required fields (Name, Email, Phone).
3) Select an RSVP status and optional plus ones/dietary restrictions.
4) Submit the form.

**Expected behavior**:

- Success message shown after submission.
- Data is inserted into the `guests` table.
- Form clears automatically.

### Testing Bug Fixes

- Verify the `events` table has `title` and `event_date`.
- Run `\d guests;` and confirm the three new columns exist.
- Submit a new RSVP and check that all fields (including new ones) are saved.

---

## What I'd Improve With More Time

- Add duplicate email prevention logic.
- Create a “View My RSVP” lookup feature.
- Add event selection support (for multiple events).
- Improve form animation and transitions.
- Implement RSVP confirmation email notifications.

---

## Additional Notes

- I chose manual SQL migrations for better visibility and control.
- Verified all API routes (`GET`, `POST`, `DELETE`) function correctly post-migration.
- Confirmed backward compatibility — older guest entries remain valid.

---

## Migration System Details

### Tool Chosen: Manual SQL Migrations

I used raw SQL migration files for simplicity and full transparency over schema changes.

### How to Run Migrations

Run the following command for each migration:

```bash
psql -U <user> -d <database> -f backend/migrations/<filename>.sql
```

### Schema Changes Implemented

- Added `notes` to `guests`
- Created `events` table
- Added `rsvp_date`, `plus_ones`, and `dietary_restrictions` to `guests`

### Rollback Strategy

Each migration file can be reverted manually by dropping or altering columns as needed:

```sql
ALTER TABLE guests DROP COLUMN plus_ones;
ALTER TABLE guests DROP COLUMN dietary_restrictions;
ALTER TABLE guests DROP COLUMN rsvp_date;
```

### Performance Considerations

- Indexed foreign key `event_id` for faster event-to-guest lookups.
- Used default values to avoid insert errors for missing fields.

---

## Self-Assessment

| Category      | Confidence |
|---------------|------------|
| Code Quality  | High       |
| Functionality | High       |
| Design        | High       |
| Bug Fixes     | High       |

### What I’m Most Proud Of

Delivering a fully functional, responsive, and bug-free RSVP system with proper schema migrations and a clean, user-friendly interface — all while ensuring backend and frontend consistency.

---

Thank you for reviewing my submission! 🚀
