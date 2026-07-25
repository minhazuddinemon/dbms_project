package rest

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"dbms-project/internal/db"
)

func (h *Handler) HandleListPrograms(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) HandleGetProgramByID(w http.ResponseWriter, r *http.Request) {
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

type ProgramPayload struct {
	PName        string  `json:"p_name"`
	PUnit        string  `json:"p_unit"`
	TotalSeats   int32   `json:"total_seats"`
	PrevCutmarks float64 `json:"prev_cutmarks"`
	Deadline     string  `json:"deadline"` // Format: YYYY-MM-DD
	UID          int32   `json:"u_id"`
}

// POST /admin/program
func (h *Handler) HandleCreateProgram(w http.ResponseWriter, r *http.Request) {
	var req ProgramPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	deadlineDate, err := time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		http.Error(w, "Invalid deadline format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	s1 := req.PUnit
	res1 := sql.NullString{
		String: s1,
	}
	if s1 != "" {
		res1.Valid = true
	}

	val := req.PrevCutmarks
	str := strconv.FormatFloat(val, 'f', -1, 64)
	res2 := sql.NullString{
		String: str,
	}
	if str != "" {
		res2.Valid = true
	}

	res, err := h.Queries.InsertProgram(r.Context(), db.InsertProgramParams{
		PName:        req.PName,
		PUnit:        res1,
		TotalSeats:   req.TotalSeats,
		PrevCutmarks: res2, // Make sure sqlc generated this as float64/string based on NUMERIC(6,2)
		Deadline:     deadlineDate,
		UID:          req.UID,
	})
	if err != nil {
		http.Error(w, "Failed to create program", http.StatusInternalServerError)
		return
	}

	programID, _ := res.LastInsertId()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message":    "Program created successfully",
		"program_id": programID,
	})
}

// PUT /admin/program?program_id=1
func (h *Handler) HandleUpdateProgram(w http.ResponseWriter, r *http.Request) {
	pIDStr := r.URL.Query().Get("program_id")
	programID, err := strconv.Atoi(pIDStr)
	if err != nil || programID <= 0 {
		http.Error(w, "Invalid or missing program_id", http.StatusBadRequest)
		return
	}

	var req ProgramPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	deadlineDate, err := time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		http.Error(w, "Invalid deadline format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	s1 := req.PUnit
	res1 := sql.NullString{
		String: s1,
	}
	if s1 != "" {
		res1.Valid = true
	}

	val := req.PrevCutmarks
	str := strconv.FormatFloat(val, 'f', -1, 64)
	res2 := sql.NullString{
		String: str,
	}
	if str != "" {
		res2.Valid = true
	}

	err = h.Queries.UpdateProgram(r.Context(), db.UpdateProgramParams{
		ProgramID:    int32(programID),
		PName:        req.PName,
		PUnit:        res1,
		TotalSeats:   req.TotalSeats,
		PrevCutmarks: res2,
		Deadline:     deadlineDate,
		UID:          req.UID,
	})
	if err != nil {
		http.Error(w, "Failed to update program", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Program updated successfully",
	})
}

// DELETE /admin/program?program_id=1
func (h *Handler) HandleDeleteProgram(w http.ResponseWriter, r *http.Request) {
	// Extract program_id from the query string
	pIDStr := r.URL.Query().Get("program_id")
	programID, err := strconv.Atoi(pIDStr)
	if err != nil || programID <= 0 {
		http.Error(w, "Invalid or missing program_id query parameter", http.StatusBadRequest)
		return
	}

	// Execute the delete query
	err = h.Queries.DeleteProgram(r.Context(), int32(programID))
	if err != nil {
		http.Error(w, "Failed to delete program", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Program deleted successfully",
	})
}
