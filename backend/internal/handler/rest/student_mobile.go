package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"dbms-project/internal/db"
	"dbms-project/internal/middleware"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

type studentMobileRequest struct {
	MobileNo  string `json:"mobile_no"`
	OwnerType string `json:"owner_type"`
}

type updateStudentMobileRequest struct {
	CurrentMobileNo string `json:"current_mobile_no"`
	MobileNo        string `json:"mobile_no"`
	OwnerType       string `json:"owner_type"`
}

func (h *Handler) HandleGetStudentMobiles(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	mobiles, err := h.Queries.GetStudentMobiles(r.Context(), studentID)
	if err != nil {
		http.Error(w, "Failed to fetch student mobiles", http.StatusInternalServerError)
		return
	}

	if mobiles == nil {
		mobiles = []db.StudentMobile{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mobiles)
}

func (h *Handler) HandleAddStudentMobile(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req studentMobileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mobileNo := strings.TrimSpace(req.MobileNo)
	ownerType, valid := parseStudentMobileOwnerType(req.OwnerType)
	if mobileNo == "" || !valid {
		http.Error(w, "mobile_no and a valid owner_type are required", http.StatusBadRequest)
		return
	}

	_, err := h.Queries.InsertStudentMobile(r.Context(), db.InsertStudentMobileParams{
		StudentID: studentID,
		MobileNo:  mobileNo,
		OwnerType: ownerType,
	})
	if err != nil {
		if isDuplicateKeyError(err) {
			http.Error(w, "Mobile number already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to add student mobile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Student mobile added successfully",
		"mobile": map[string]any{
			"student_id": studentID,
			"mobile_no":  mobileNo,
			"owner_type": ownerType,
		},
	})
}

func (h *Handler) HandleUpdateStudentMobile(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req updateStudentMobileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	currentMobileNo := strings.TrimSpace(req.CurrentMobileNo)
	mobileNo := strings.TrimSpace(req.MobileNo)
	ownerType, valid := parseStudentMobileOwnerType(req.OwnerType)
	if currentMobileNo == "" || mobileNo == "" || !valid {
		http.Error(w, "current_mobile_no, mobile_no and a valid owner_type are required", http.StatusBadRequest)
		return
	}

	res, err := h.Queries.UpdateStudentMobile(r.Context(), db.UpdateStudentMobileParams{
		MobileNo:   mobileNo,
		OwnerType:  ownerType,
		StudentID:  studentID,
		MobileNo_2: currentMobileNo,
	})
	if err != nil {
		if isDuplicateKeyError(err) {
			http.Error(w, "Mobile number already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to update student mobile", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to update student mobile", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Student mobile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Student mobile updated successfully",
		"mobile": map[string]any{
			"student_id": studentID,
			"mobile_no":  mobileNo,
			"owner_type": ownerType,
		},
	})
}

func (h *Handler) HandleDeleteStudentMobile(w http.ResponseWriter, r *http.Request) {
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	mobileNo := strings.TrimSpace(r.URL.Query().Get("mobile_no"))
	if mobileNo == "" {
		http.Error(w, "mobile_no query parameter is required", http.StatusBadRequest)
		return
	}

	res, err := h.Queries.DeleteStudentMobile(r.Context(), db.DeleteStudentMobileParams{
		StudentID: studentID,
		MobileNo:  mobileNo,
	})
	if err != nil {
		http.Error(w, "Failed to delete student mobile", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to delete student mobile", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Student mobile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Student mobile deleted successfully",
	})
}

func parseStudentMobileOwnerType(value string) (db.StudentMobileOwnerType, bool) {
	ownerType := db.StudentMobileOwnerType(strings.ToLower(strings.TrimSpace(value)))
	switch ownerType {
	case db.StudentMobileOwnerTypeSelf, db.StudentMobileOwnerTypeMother, db.StudentMobileOwnerTypeFather:
		return ownerType, true
	default:
		return "", false
	}
}

func isDuplicateKeyError(err error) bool {
	mysqlErr, ok := err.(*mysqlDriver.MySQLError)
	return ok && mysqlErr.Number == 1062
}
