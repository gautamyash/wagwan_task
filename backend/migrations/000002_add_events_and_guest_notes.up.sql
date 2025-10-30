-- Create events table
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    event_date TIMESTAMP,
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Add new columns to guests table
ALTER TABLE guests
ADD COLUMN notes TEXT,
    ADD COLUMN event_id INT,
    ADD COLUMN rsvp_date TIMESTAMP,
    ADD COLUMN plus_ones INT DEFAULT 0,
    ADD COLUMN dietary_restrictions TEXT;
-- Add foreign key constraint
ALTER TABLE guests
ADD CONSTRAINT fk_event FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE
SET NULL;