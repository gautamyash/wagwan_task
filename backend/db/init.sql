-- Create guests table
CREATE TABLE IF NOT EXISTS guests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'attending', 'declined')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create index for faster queries
CREATE INDEX IF NOT EXISTS idx_guests_status ON guests(status);
CREATE INDEX IF NOT EXISTS idx_guests_name ON guests(name);

-- Insert sample data
INSERT INTO guests (name, email, phone, status, created_at) VALUES
    ('Alice Johnson', 'alice@example.com', '+1-555-0101', 'attending', NOW() - INTERVAL '5 days'),
    ('Bob Smith', 'bob@example.com', '+1-555-0102', 'pending', NOW() - INTERVAL '4 days'),
    ('Charlie Brown', 'charlie@example.com', '+1-555-0103', 'attending', NOW() - INTERVAL '3 days'),
    ('Diana Prince', 'diana@example.com', '+1-555-0104', 'declined', NOW() - INTERVAL '2 days'),
    ('Eve Wilson', 'eve@example.com', '+1-555-0105', 'pending', NOW() - INTERVAL '1 day'),
    ('Frank Miller', 'frank@example.com', '+1-555-0106', 'attending', NOW()),
    ('Grace Lee', 'grace@example.com', '+1-555-0107', 'pending', NOW() - INTERVAL '6 days'),
    ('Henry Davis', 'henry@example.com', '+1-555-0108', 'attending', NOW() - INTERVAL '7 days');

