package db

import (
	"database/sql"
	"errors"
	"time"
)

// ------------------ Struct ------------------
type Guest struct {
    ID                  int            `json:"id"`
    Name                string         `json:"name"`
    Email               string         `json:"email"`
    Phone               string         `json:"phone"`
    Status              string         `json:"status"`
    Notes               sql.NullString `json:"notes,omitempty"`
    EventID             sql.NullInt64  `json:"event_id,omitempty"`
    PlusOnes            int            `json:"plus_ones"`
    DietaryRestrictions sql.NullString `json:"dietary_restrictions,omitempty"`
    CreatedAt           time.Time      `json:"created_at"`
}

// ------------------ GetAllGuests ------------------
func GetAllGuests(db *sql.DB, status string) ([]Guest, error) {
	var rows *sql.Rows
	var err error

	if status != "" {
		rows, err = db.Query(`
			SELECT id, name, email, phone, status, notes, event_id, plus_ones, dietary_restrictions, created_at
			FROM guests
			WHERE status = $1
			ORDER BY created_at ASC
		`, status)
	} else {
		rows, err = db.Query(`
			SELECT id, name, email, phone, status, notes, event_id, plus_ones, dietary_restrictions, created_at
			FROM guests
			ORDER BY created_at ASC
		`)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []Guest
	for rows.Next() {
		var g Guest
		err := rows.Scan(
			&g.ID, &g.Name, &g.Email, &g.Phone, &g.Status,
			&g.Notes, &g.EventID, &g.PlusOnes, &g.DietaryRestrictions, &g.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		guests = append(guests, g)
	}

	return guests, rows.Err()
}

// ------------------ GetGuestByID ------------------
func GetGuestByID(db *sql.DB, id int) (*Guest, error) {
	var g Guest
	err := db.QueryRow(`
		SELECT id, name, email, phone, status, notes, event_id, plus_ones, dietary_restrictions, created_at
		FROM guests
		WHERE id = $1
	`, id).Scan(
		&g.ID, &g.Name, &g.Email, &g.Phone, &g.Status,
		&g.Notes, &g.EventID, &g.PlusOnes, &g.DietaryRestrictions, &g.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("guest not found")
	}
	if err != nil {
		return nil, err
	}

	return &g, nil
}

// ------------------ CreateGuest ------------------
func CreateGuest(db *sql.DB, name, email, phone, status, notes string, eventID, plusOnes int, dietaryRestrictions string) (*Guest, error) {
	var g Guest
	
	// Convert empty strings to NULL for optional fields
	var notesPtr interface{} = notes
	if notes == "" {
		notesPtr = nil
	}
	
	var dietaryRestrictionsPtr interface{} = dietaryRestrictions
	if dietaryRestrictions == "" {
		dietaryRestrictionsPtr = nil
	}

	// Set default value for plus_ones if not provided
	if plusOnes < 0 {
		plusOnes = 0
	}

	// Handle event_id
	var eventIDPtr interface{} = eventID
	if eventID <= 0 {
		eventIDPtr = nil
	}

	err := db.QueryRow(`
		INSERT INTO guests (name, email, phone, status, notes, event_id, plus_ones, dietary_restrictions, rsvp_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
		RETURNING id, name, email, phone, status, notes, event_id, plus_ones, dietary_restrictions, created_at
	`, name, email, phone, status, notesPtr, eventIDPtr, plusOnes, dietaryRestrictionsPtr).Scan(
		&g.ID, &g.Name, &g.Email, &g.Phone, &g.Status,
		&g.Notes, &g.EventID, &g.PlusOnes, &g.DietaryRestrictions, &g.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &g, nil
}

// ------------------ DeleteGuest ------------------
func DeleteGuest(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM guests WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("guest not found")
	}

	return nil
}
