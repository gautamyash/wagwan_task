-- Migration: Add plus_ones, dietary_restrictions, and rsvp_date columns to guests table

ALTER TABLE guests 
ADD COLUMN IF NOT EXISTS plus_ones INT DEFAULT 0,
ADD COLUMN IF NOT EXISTS dietary_restrictions TEXT,
ADD COLUMN IF NOT EXISTS rsvp_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
