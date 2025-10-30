-- ✅ Create events table (corrected schema)
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    event_date TIMESTAMP NOT NULL,
    location VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- ✅ Insert one sample event
INSERT INTO events (title, description, event_date, location)
VALUES (
        'Sample Event',
        'Initial seeded event',
        CURRENT_TIMESTAMP + INTERVAL '7 days',
        'Downtown Hall'
    ) ON CONFLICT DO NOTHING;
-- ✅ Create guests table (aligned with requirements)
CREATE TABLE IF NOT EXISTS guests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'attending', 'declined')),
    notes TEXT,
    rsvp_date TIMESTAMP DEFAULT NULL,
    plus_ones INT DEFAULT 0,
    dietary_restrictions TEXT,
    event_id INT REFERENCES events(id) ON DELETE
    SET NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- ✅ Insert sample guests linked to the seeded event (event_id = 1)
INSERT INTO guests (
        name,
        email,
        phone,
        status,
        notes,
        rsvp_date,
        plus_ones,
        dietary_restrictions,
        event_id,
        created_at
    )
VALUES (
        'Alice Johnson',
        'alice@example.com',
        '+1-555-0101',
        'attending',
        'Excited to join!',
        NOW() - INTERVAL '5 days',
        0,
        '',
        1,
        NOW() - INTERVAL '5 days'
    ),
    (
        'Bob Smith',
        'bob@example.com',
        '+1-555-0102',
        'pending',
        'Might bring a friend',
        NULL,
        1,
        '',
        1,
        NOW() - INTERVAL '4 days'
    ),
    (
        'Charlie Brown',
        'charlie@example.com',
        '+1-555-0103',
        'attending',
        '',
        NOW() - INTERVAL '3 days',
        0,
        '',
        1,
        NOW() - INTERVAL '3 days'
    ),
    (
        'Diana Prince',
        'diana@example.com',
        '+1-555-0104',
        'declined',
        'Out of town',
        NULL,
        0,
        '',
        1,
        NOW() - INTERVAL '2 days'
    ),
    (
        'Eve Wilson',
        'eve@example.com',
        '+1-555-0105',
        'pending',
        '',
        NULL,
        0,
        '',
        1,
        NOW() - INTERVAL '1 day'
    ),
    (
        'Frank Miller',
        'frank@example.com',
        '+1-555-0106',
        'attending',
        'Vegan meal please',
        NOW(),
        0,
        'Vegan',
        1,
        NOW()
    ),
    (
        'Grace Lee',
        'grace@example.com',
        '+1-555-0107',
        'pending',
        '',
        NULL,
        0,
        '',
        1,
        NOW() - INTERVAL '6 days'
    ),
    (
        'Henry Davis',
        'henry@example.com',
        '+1-555-0108',
        'attending',
        'Will bring +1',
        NOW() - INTERVAL '7 days',
        1,
        '',
        1,
        NOW() - INTERVAL '7 days'
    );