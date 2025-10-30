// migrate.go
package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/eventguests?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("❌ Migration setup failed: %v", err)
	}

	var direction string
	if len(os.Args) > 1 {
		direction = os.Args[1]
	}

	switch direction {
	case "down":
		err = m.Steps(-1)
	default:
		err = m.Up()
	}

	if err != nil && err.Error() != "no change" {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration applied successfully!")
}
