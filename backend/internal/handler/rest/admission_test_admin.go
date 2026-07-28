package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"dbms-project/internal/db"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

type AdmissionTestPayload struct {
	ExamUnit     string `json:"exam_unit"`
	ExamCenter   string `json:"exam_center"`
	ExamDate     string `json:"exam_date"`
	PrereqTestID *int32 `json:"prereq_test_id"`
	ProgramID    int32  `json:"program_id"`
}

func (h *Handler) HandleCreateAdmissionTest(w http.ResponseWriter, r *http.Request) {
	var req AdmissionTestPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	params, err := buildAdmissionTestParams(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Queries.InsertAdmissionTest(r.Context(), params)
	if err != nil {
		status, message := mapAdmissionTestWriteError(err)
		http.Error(w, message, status)
		return
	}

	testID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Admission test created but failed to retrieve ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Admission test created successfully",
		"test_id": testID,
	})
}

func (h *Handler) HandleUpdateAdmissionTest(w http.ResponseWriter, r *http.Request) {
	testIDStr := r.URL.Query().Get("test_id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil || testID <= 0 {
		http.Error(w, "Invalid or missing test_id", http.StatusBadRequest)
		return
	}

	var req AdmissionTestPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if req.PrereqTestID != nil && *req.PrereqTestID == int32(testID) {
		http.Error(w, "prereq_test_id cannot be the same as test_id", http.StatusBadRequest)
		return
	}

	params, err := buildAdmissionTestParams(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Queries.UpdateAdmissionTest(r.Context(), db.UpdateAdmissionTestParams{
		ExamUnit:     params.ExamUnit,
		ExamCenter:   params.ExamCenter,
		ExamDate:     params.ExamDate,
		PrereqTestID: params.PrereqTestID,
		ProgramID:    params.ProgramID,
		TestID:       int32(testID),
	})
	if err != nil {
		status, message := mapAdmissionTestWriteError(err)
		http.Error(w, message, status)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to update admission test", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Admission test not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Admission test updated successfully",
	})
}

func buildAdmissionTestParams(req AdmissionTestPayload) (db.InsertAdmissionTestParams, error) {
	examDateStr := strings.TrimSpace(req.ExamDate)
	if examDateStr == "" {
		return db.InsertAdmissionTestParams{}, errBadRequest("exam_date is required")
	}

	examDate, err := time.Parse("2006-01-02", examDateStr)
	if err != nil {
		return db.InsertAdmissionTestParams{}, errBadRequest("Invalid exam_date format. Use YYYY-MM-DD")
	}

	if req.ProgramID <= 0 {
		return db.InsertAdmissionTestParams{}, errBadRequest("program_id is required and must be greater than 0")
	}

	if req.PrereqTestID != nil && *req.PrereqTestID <= 0 {
		return db.InsertAdmissionTestParams{}, errBadRequest("prereq_test_id must be greater than 0 when provided")
	}

	return db.InsertAdmissionTestParams{
		ExamUnit:     admissionTestToNullString(req.ExamUnit),
		ExamCenter:   admissionTestToNullString(req.ExamCenter),
		ExamDate:     sql.NullTime{Time: examDate, Valid: true},
		PrereqTestID: admissionTestToNullInt32(req.PrereqTestID),
		ProgramID:    sql.NullInt32{Int32: req.ProgramID, Valid: true},
	}, nil
}

func admissionTestToNullString(value string) sql.NullString {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: trimmed, Valid: true}
}

func admissionTestToNullInt32(value *int32) sql.NullInt32 {
	if value == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *value, Valid: true}
}

type badRequestError struct {
	message string
}

func (e badRequestError) Error() string {
	return e.message
}

func errBadRequest(message string) error {
	return badRequestError{message: message}
}

func mapAdmissionTestWriteError(err error) (int, string) {
	mysqlErr, ok := err.(*mysqlDriver.MySQLError)
	if !ok {
		return http.StatusInternalServerError, "Failed to save admission test"
	}

	switch mysqlErr.Number {
	case 1452:
		return http.StatusBadRequest, "Invalid program_id or prereq_test_id"
	default:
		return http.StatusInternalServerError, "Failed to save admission test"
	}
}
