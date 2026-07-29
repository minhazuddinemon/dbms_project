package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"dbms-project/internal/db"
)

type DepartmentPayload struct {
	DeptName        string `json:"dept_name"`
	DeptDescription string `json:"dept_description"`
	TotalSeats      int32  `json:"total_seats"`
}

type AlbumPayload struct {
	PictureTitle string `json:"picture_title"`
	PictureURL   string `json:"picture_url"`
}

type UniversityRequest struct {
	Name                  string              `json:"name"`
	Website               string              `json:"website"`
	Location              string              `json:"location"`
	LogoURL               string              `json:"logo_url"`
	UniversityDescription string              `json:"university_description"`
	UniversityHistory     string              `json:"university_history"`
	Departments           []DepartmentPayload `json:"departments"`
	Album                 []AlbumPayload      `json:"album"`
}

type DepartmentResponse struct {
	DeptID          int32  `json:"dept_id"`
	DeptName        string `json:"dept_name"`
	DeptDescription string `json:"dept_description"`
	TotalSeats      int32  `json:"total_seats"`
}

type AlbumResponse struct {
	AlbumID      int32  `json:"album_id"`
	PictureTitle string `json:"picture_title"`
	PictureURL   string `json:"picture_url"`
}

type UniversityDetailResponse struct {
	UID         int32                `json:"u_id"`
	UName       string               `json:"u_name"`
	Website     string               `json:"website"`
	Location    string               `json:"location"`
	LogoURL     string               `json:"logo_url"`
	Description string               `json:"university_description"`
	History     string               `json:"university_history"`
	Departments []DepartmentResponse `json:"departments"`
	Album       []AlbumResponse      `json:"album"`
}

func toNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
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

	// Start database transaction
	tx, err := h.DB.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	qtx := h.Queries.WithTx(tx)

	res, err := qtx.InsertUniversity(r.Context(), db.InsertUniversityParams{
		UName:                 req.Name,
		Website:               req.Website,
		Location:              toNullString(req.Location),
		LogoUrl:               toNullString(req.LogoURL),
		UniversityDescription: toNullString(req.UniversityDescription),
		UniversityHistory:     toNullString(req.UniversityHistory),
	})
	if err != nil {
		http.Error(w, "Failed to create university", http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve university ID", http.StatusInternalServerError)
		return
	}
	uID := int32(id)

	// Insert departments
	for _, dept := range req.Departments {
		_, err := qtx.InsertUniversityDepartment(r.Context(), db.InsertUniversityDepartmentParams{
			UID:             uID,
			DeptName:        dept.DeptName,
			DeptDescription: toNullString(dept.DeptDescription),
			TotalSeats:      dept.TotalSeats,
		})
		if err != nil {
			http.Error(w, "Failed to create university department", http.StatusInternalServerError)
			return
		}
	}

	// Insert album pictures
	for _, pic := range req.Album {
		_, err := qtx.InsertUniversityAlbumPicture(r.Context(), db.InsertUniversityAlbumPictureParams{
			UID:          uID,
			PictureTitle: pic.PictureTitle,
			PictureUrl:   pic.PictureURL,
		})
		if err != nil {
			http.Error(w, "Failed to create university album picture", http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "University created successfully",
		"u_id":    uID,
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

	if req.Name == "" || req.Website == "" {
		http.Error(w, "Name and website are required", http.StatusBadRequest)
		return
	}

	// Start database transaction
	tx, err := h.DB.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	qtx := h.Queries.WithTx(tx)

	err = qtx.UpdateUniversity(r.Context(), db.UpdateUniversityParams{
		UID:                   int32(uID),
		UName:                 req.Name,
		Website:               req.Website,
		Location:              toNullString(req.Location),
		LogoUrl:               toNullString(req.LogoURL),
		UniversityDescription: toNullString(req.UniversityDescription),
		UniversityHistory:     toNullString(req.UniversityHistory),
	})
	if err != nil {
		http.Error(w, "Failed to update university", http.StatusInternalServerError)
		return
	}

	// Update departments: delete all existing, then insert new
	err = qtx.DeleteUniversityDepartments(r.Context(), int32(uID))
	if err != nil {
		http.Error(w, "Failed to clear existing university departments", http.StatusInternalServerError)
		return
	}

	for _, dept := range req.Departments {
		_, err := qtx.InsertUniversityDepartment(r.Context(), db.InsertUniversityDepartmentParams{
			UID:             int32(uID),
			DeptName:        dept.DeptName,
			DeptDescription: toNullString(dept.DeptDescription),
			TotalSeats:      dept.TotalSeats,
		})
		if err != nil {
			http.Error(w, "Failed to update university departments", http.StatusInternalServerError)
			return
		}
	}

	// Update album: delete all existing, then insert new
	err = qtx.DeleteUniversityAlbum(r.Context(), int32(uID))
	if err != nil {
		http.Error(w, "Failed to clear existing university album", http.StatusInternalServerError)
		return
	}

	for _, pic := range req.Album {
		_, err := qtx.InsertUniversityAlbumPicture(r.Context(), db.InsertUniversityAlbumPictureParams{
			UID:          int32(uID),
			PictureTitle: pic.PictureTitle,
			PictureUrl:   pic.PictureURL,
		})
		if err != nil {
			http.Error(w, "Failed to update university album", http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "University updated successfully",
	})
}

type UniversityOverviewResponse struct {
	UID              int32  `json:"u_id"`
	UName            string `json:"u_name"`
	Website          string `json:"website"`
	Location         string `json:"location"`
	LogoURL          string `json:"logo_url"`
	ShortDescription string `json:"university_description"`
	ShortHistory     string `json:"university_history"`
}

// GET /universities (Public - Frontend List)
func (h *Handler) HandleListUniversities(w http.ResponseWriter, r *http.Request) {
	universities, err := h.Queries.ListUniversities(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve universities", http.StatusInternalServerError)
		return
	}

	response := make([]UniversityOverviewResponse, 0, len(universities))
	for _, u := range universities {
		desc := nullStringToString(u.UniversityDescription)
		history := nullStringToString(u.UniversityHistory)

		response = append(response, UniversityOverviewResponse{
			UID:              u.UID,
			UName:            extractNickname(u.UName),
			Website:          u.Website,
			Location:         nullStringToString(u.Location),
			LogoURL:          nullStringToString(u.LogoUrl),
			ShortDescription: extractShortDescription(desc),
			ShortHistory:     extractShortHistory(history),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GET /universities/detail?u_id=1 (Public - Frontend Detail View)
func (h *Handler) HandleGetUniversityByID(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.URL.Query().Get("u_id")
	uID, err := strconv.Atoi(uIDStr)
	if err != nil || uID <= 0 {
		http.Error(w, "Invalid or missing u_id query parameter", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	university, err := h.Queries.GetUniversityByID(ctx, int32(uID))
	if err != nil {
		http.Error(w, "University not found", http.StatusNotFound)
		return
	}

	// Fetch departments
	depts, err := h.Queries.GetUniversityDepartments(ctx, int32(uID))
	if err != nil {
		http.Error(w, "Failed to fetch university departments", http.StatusInternalServerError)
		return
	}

	deptResponses := make([]DepartmentResponse, 0, len(depts))
	for _, d := range depts {
		deptResponses = append(deptResponses, DepartmentResponse{
			DeptID:          d.DeptID,
			DeptName:        d.DeptName,
			DeptDescription: nullStringToString(d.DeptDescription),
			TotalSeats:      d.TotalSeats,
		})
	}

	// Fetch album
	albumPics, err := h.Queries.GetUniversityAlbum(ctx, int32(uID))
	if err != nil {
		http.Error(w, "Failed to fetch university album", http.StatusInternalServerError)
		return
	}

	albumResponses := make([]AlbumResponse, 0, len(albumPics))
	for _, p := range albumPics {
		albumResponses = append(albumResponses, AlbumResponse{
			AlbumID:      p.AlbumID,
			PictureTitle: p.PictureTitle,
			PictureURL:   p.PictureUrl,
		})
	}

	resp := UniversityDetailResponse{
		UID:         university.UID,
		UName:       extractNameWithoutNickname(university.UName),
		Website:     university.Website,
		Location:    nullStringToString(university.Location),
		LogoURL:     nullStringToString(university.LogoUrl),
		Description: extractHTMLComponent(nullStringToString(university.UniversityDescription)),
		History:     extractHTMLComponent(nullStringToString(university.UniversityHistory)),
		Departments: deptResponses,
		Album:       albumResponses,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// --- String Parsing Helper Functions ---

func extractNickname(fullName string) string {
	start := strings.Index(fullName, "(")
	end := strings.Index(fullName, ")")
	if start != -1 && end > start {
		return fullName[start+1 : end]
	}
	return fullName
}

func extractNameWithoutNickname(fullName string) string {
	before, _, ok := strings.Cut(fullName, "(")
	if ok {
		name := before
		name = strings.TrimSuffix(name, ".")
		return strings.TrimSpace(name)
	}
	return fullName
}

func extractShortDescription(fullDesc string) string {
	lines := strings.Split(fullDesc, "\n")
	if len(lines) > 0 {
		firstLine := strings.TrimSpace(lines[0])
		if strings.HasPrefix(strings.ToLower(firstLine), "short description:") {
			return strings.TrimSpace(firstLine[18:])
		}
		return firstLine
	}
	return fullDesc
}

func extractShortHistory(fullHistory string) string {
	lines := strings.Split(fullHistory, "\n")
	if len(lines) > 0 {
		firstLine := strings.TrimSpace(lines[0])
		if strings.HasPrefix(strings.ToLower(firstLine), "short history:") {
			return strings.TrimSpace(firstLine[14:])
		}
		return firstLine
	}
	return fullHistory
}

func extractHTMLComponent(fullText string) string {
	idx := strings.Index(fullText, "<div")
	if idx != -1 {
		return strings.TrimSpace(fullText[idx:])
	}
	return fullText
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
