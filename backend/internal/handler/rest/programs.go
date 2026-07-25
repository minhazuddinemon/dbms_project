package rest

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"dbms-project/internal/db"
)

// HandleListPrograms processes search filters from URL parameters:
// GET /programs?search=engineering&unit=A
func (h *Handler) HandleListPrograms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	search := query.Get("search")
	unit := query.Get("unit") // e.g., 'A', 'B'

	// Use sql.NullString to satisfy sqlc's type requirements
	programs, err := h.Queries.ListPrograms(r.Context(), db.ListProgramsParams{
		CONCAT:   sql.NullString{String: search, Valid: true},
		CONCAT_2: sql.NullString{String: search, Valid: true},
		PUnit:    sql.NullString{String: unit, Valid: true},
		Column4:  sql.NullString{String: unit, Valid: true},
	})

	if err != nil {
		log.Printf("Failed to fetch programs: %v", err)
		http.Error(w, "Failed to fetch programs", http.StatusInternalServerError)
		return
	}

	if programs == nil {
		programs = []db.ListProgramsRow{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(programs)
}

// HandleGetProgramByID fetches details for a single program
// GET /programs/detail?id=1
func (h *Handler) HandleGetProgramByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing program id parameter", http.StatusBadRequest)
		return
	}

	programID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid program id", http.StatusBadRequest)
		return
	}

	program, err := h.Queries.GetProgramByID(r.Context(), int32(programID))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Program not found", http.StatusNotFound)
		} else {
			log.Printf("Failed to fetch program detail: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(program)
}
