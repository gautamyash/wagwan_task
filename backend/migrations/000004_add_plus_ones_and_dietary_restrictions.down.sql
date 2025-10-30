-- Rollback: Remove plus_ones, dietary_restrictions, and rsvp_date columns from guests table

ALTER TABLE guests 
DROP COLUMN IF EXISTS rsvp_date,
DROP COLUMN IF EXISTS dietary_restrictions,
DROP COLUMN IF EXISTS plus_ones;
