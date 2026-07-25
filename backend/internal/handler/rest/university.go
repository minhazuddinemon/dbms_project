package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"dbms-project/internal/db"
)

type UniversityRequest struct {
	Name     string `json:"name"`
	Website  string `json:"website"`
	Location string `json:"location"`
	LogoURL  string `json:"logo_url"`
}

// POST /admin/university
func (h *Handler) HandleCreateUniversity(w http.ResponseWriter, r *http.Request) {
	var req UniversityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Website == "" {
		http.Error(w, "Name and website are required", http.StatusBadRequest)
		return
	}

	s1 := req.Location
	res1 := sql.NullString{
		String: s1,
	}
	if s1 != "" {
		res1.Valid = true
	}

	s2 := req.Location
	res2 := sql.NullString{
		String: s2,
	}
	if s2 != "" {
		res1.Valid = true
	}

	res, err := h.Queries.InsertUniversity(r.Context(), db.InsertUniversityParams{
		UName:    req.Name,
		Website:  req.Website,
		Location: res1,
		LogoUrl:  res2,
	})
	if err != nil {
		http.Error(w, "Failed to create university", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "University created successfully",
		"u_id":    id,
	})
}

// PUT /admin/university?u_id=1
func (h *Handler) HandleUpdateUniversity(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.URL.Query().Get("u_id")
	uID, err := strconv.Atoi(uIDStr)
	if err != nil || uID <= 0 {
		http.Error(w, "Invalid or missing u_id query parameter", http.StatusBadRequest)
		return
	}

	var req UniversityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	s1 := req.Location
	res1 := sql.NullString{
		String: s1,
	}
	if s1 != "" {
		res1.Valid = true
	}

	s2 := req.Location
	res2 := sql.NullString{
		String: s2,
	}
	if s2 != "" {
		res1.Valid = true
	}

	err = h.Queries.UpdateUniversity(r.Context(), db.UpdateUniversityParams{
		UID:      int32(uID),
		UName:    req.Name,
		Website:  req.Website,
		Location: res1,
		LogoUrl:  res2,
	})
	if err != nil {
		http.Error(w, "Failed to update university", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "University updated successfully",
	})
}

// GET /universities (Public - Frontend List)
func (h *Handler) HandleListUniversities(w http.ResponseWriter, r *http.Request) {
	universities, err := h.Queries.ListUniversities(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve universities", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(universities)
}

// GET /universities/detail?u_id=1 (Public - Frontend Detail View)
func (h *Handler) HandleGetUniversityByID(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.URL.Query().Get("u_id")
	uID, err := strconv.Atoi(uIDStr)
	if err != nil || uID <= 0 {
		http.Error(w, "Invalid or missing u_id query parameter", http.StatusBadRequest)
		return
	}

	university, err := h.Queries.GetUniversityByID(r.Context(), int32(uID))
	if err != nil {
		http.Error(w, "University not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(university)
}

// DELETE /admin/university?u_id=1
func (h *Handler) HandleDeleteUniversity(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.URL.Query().Get("u_id")
	uID, err := strconv.Atoi(uIDStr)
	if err != nil || uID <= 0 {
		http.Error(w, "Invalid or missing u_id query parameter", http.StatusBadRequest)
		return
	}

	err = h.Queries.DeleteUniversity(r.Context(), int32(uID))
	if err != nil {
		http.Error(w, "Failed to delete university", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "University deleted successfully",
	})
}
