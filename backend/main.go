package main

import (
	"database/sql"
	"event-guest-manager/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	// Database connection
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "eventguests")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Wait for DB
	for i := 0; i < 30; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		log.Println("Waiting for database...")
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatal("Database not available:", err)
	}

	log.Println("✅ Connected to database")

	guestHandler := handlers.NewGuestHandler(db)

	router := mux.NewRouter()

	// --- API routes ---
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/guests", guestHandler.GetGuests).Methods("GET")
	api.HandleFunc("/guests", guestHandler.CreateGuest).Methods("POST")
	api.HandleFunc("/guests/{id}", guestHandler.GetGuest).Methods("GET")
	api.HandleFunc("/guests/{id}", guestHandler.DeleteGuest).Methods("DELETE")

	// --- Serve Svelte static frontend ---
	fs := http.FileServer(http.Dir("../frontend/build"))
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../frontend/build" + r.URL.Path
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, "../frontend/build/index.html")
			return
		}
		fs.ServeHTTP(w, r)
	})

	// --- CORS setup ---
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:4173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	port := getEnv("PORT", "8080")
	log.Printf("🚀 Server running at: http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(router)))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
