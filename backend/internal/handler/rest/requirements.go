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

func (h *Handler) GetProgramRequirementsStatus(w http.ResponseWriter, r *http.Request) {
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
