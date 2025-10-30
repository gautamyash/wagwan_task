package handlers

import (
	"database/sql"
	"encoding/json"
	"event-guest-manager/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GuestHandler struct {
	db *sql.DB
}

func NewGuestHandler(database *sql.DB) *GuestHandler {
	return &GuestHandler{db: database}
}

// ------------------ REQUEST / RESPONSE STRUCTS ------------------
type CreateGuestRequest struct {
	Name                string  `json:"name"`
	Email               string  `json:"email"`
	Phone               string  `json:"phone"`
	Status              string  `json:"status"`
	Notes               *string `json:"notes,omitempty"`
	EventID             int     `json:"event_id"`
	PlusOnes            int     `json:"plus_ones,omitempty"`
	DietaryRestrictions *string `json:"dietary_restrictions,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// ------------------ GET /api/guests ------------------
func (h *GuestHandler) GetGuests(w http.ResponseWriter, r *http.Request) {
	statusFilter := r.URL.Query().Get("status")

	guests, err := db.GetAllGuests(h.db, statusFilter)
	if err != nil {
		fmt.Println("Error fetching guests:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch guests")
		return
	}

	if guests == nil {
		guests = []db.Guest{}
	}

	respondWithJSON(w, http.StatusOK, guests)
}

// ------------------ GET /api/guests/:id ------------------
func (h *GuestHandler) GetGuest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid guest ID")
		return
	}

	guest, err := db.GetGuestByID(h.db, id)
	if err != nil {
		if err.Error() == "guest not found" {
			respondWithError(w, http.StatusNotFound, "Guest not found")
			return
		}
		fmt.Println("Error fetching guest by ID:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch guest")
		return
	}

	respondWithJSON(w, http.StatusOK, guest)
}

// ------------------ POST /api/guests ------------------
func (h *GuestHandler) CreateGuest(w http.ResponseWriter, r *http.Request) {
	var req CreateGuestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Input validation
	if req.Name == "" || req.Email == "" || req.Phone == "" {
		http.Error(w, "Name, email, and phone are required", http.StatusBadRequest)
		return
	}

	// Handle optional fields
	var notes string
	if req.Notes != nil {
		notes = *req.Notes
	}

	var dietaryRestrictions string
	if req.DietaryRestrictions != nil {
		dietaryRestrictions = *req.DietaryRestrictions
	}

	// Create guest in database
	guest, err := db.CreateGuest(h.db, req.Name, req.Email, req.Phone, req.Status, notes, req.EventID, req.PlusOnes, dietaryRestrictions)
	if err != nil {
		fmt.Println("❌ Error creating guest:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create guest")
		return
	}

	fmt.Println("✅ Guest created successfully:", guest.Name)
	respondWithJSON(w, http.StatusCreated, guest)
}

// ------------------ DELETE /api/guests/:id ------------------
func (h *GuestHandler) DeleteGuest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid guest ID")
		return
	}

	err = db.DeleteGuest(h.db, id)
	if err != nil {
		if err.Error() == "guest not found" {
			respondWithError(w, http.StatusNotFound, "Guest not found")
			return
		}
		fmt.Println("Error deleting guest:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete guest")
		return
	}

	fmt.Println("🗑️ Guest deleted:", id)
	w.WriteHeader(http.StatusNoContent)
}

// ------------------ Helpers ------------------
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Failed to encode response"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}
