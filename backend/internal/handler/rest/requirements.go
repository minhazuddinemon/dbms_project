package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dbms-project/internal/db"
	"dbms-project/internal/middleware"
)

type RequiredFieldStatus struct {
	FieldName  string  `json:"field_name"`
	Value      *string `json:"value"` // Null if not filled yet
	IsProvided bool    `json:"is_provided"`
}

type ProgramRequirementsResponse struct {
	ProgramID      int32                 `json:"program_id"`
	IsReadyToApply bool                  `json:"is_ready_to_apply"`
	RequiredFields []RequiredFieldStatus `json:"required_fields"`
	MissingFields  []string              `json:"missing_fields"`
}

// GetProgramRequirementsStatus checks student completion status for a program's required fields.
// @Summary Get program requirements status
// @Description Checks if the authenticated student has fulfilled all required profile fields for a specific program.
// @Tags Requirements
// @Security BearerAuth
// @Produce json
// @Param program_id query int true "Program ID"
// @Success 200 {object} rest.ProgramRequirementsResponse "Requirements status response"
// @Failure 400 {string} string "Invalid program_id"
// @Failure 401 {string} string "Unauthorized"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Failed to fetch requirements"
// @Router /program/requirements [get]
func (h *Handler) GetProgramRequirementsStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Get Student ID from JWT Context
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 2. Read program_id from Query Params (e.g., /program/requirements?program_id=1)
	programIDStr := r.URL.Query().Get("program_id")
	programID, err := strconv.Atoi(programIDStr)
	if err != nil || programID <= 0 {
		http.Error(w, "Invalid program_id", http.StatusBadRequest)
		return
	}

	// 3. Fetch requirements & existing student values
	rows, err := h.Queries.GetStudentProgramRequirements(r.Context(), db.GetStudentProgramRequirementsParams{
		StudentID: studentID,
		ProgramID: int32(programID),
	})
	if err != nil {
		http.Error(w, "Failed to fetch requirements", http.StatusInternalServerError)
		return
	}

	// 4. Format response
	var fieldsStatus []RequiredFieldStatus
	var missingFields []string
	isReady := true

	for _, row := range rows {
		fieldName := string(row.FieldName)

		var val *string
		if row.FieldValue.Valid {
			v := row.FieldValue.String
			val = &v
		}

		isProvided := val != nil && *val != ""

		if !isProvided {
			isReady = false
			missingFields = append(missingFields, fieldName)
		}

		fieldsStatus = append(fieldsStatus, RequiredFieldStatus{
			FieldName:  fieldName,
			Value:      val,
			IsProvided: isProvided,
		})
	}

	resp := ProgramRequirementsResponse{
		ProgramID:      int32(programID),
		IsReadyToApply: isReady,
		RequiredFields: fieldsStatus,
		MissingFields:  missingFields,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
