package rest

import (
	"encoding/json"
	"net/http"

	"dbms-project/internal/db"
	"dbms-project/internal/middleware"
)

// HandleUpdateProfile upserts student profile fields (e.g. PRESENT_ADDRESS, PHOTO_URL).
// @Summary Update student profile fields
// @Description Upserts dynamic key-value profile fields for the authenticated student.
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body map[string]string true "Profile fields key-value dictionary (e.g. {'PRESENT_ADDRESS': 'Dhaka'})"
// @Success 200 {object} map[string]string "Profile updated successfully"
// @Failure 400 {string} string "Invalid request payload or failed field update"
// @Failure 401 {string} string "Unauthorized"
// @Failure 405 {string} string "Method not allowed"
// @Router /student/profile [post]
func (h *Handler) HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Allow both POST and PUT for updating resources
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Get Student ID from JWT Context
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 2. Decode the dynamic JSON payload
	// Expecting: { "PRESENT_ADDRESS": "Dhaka", "PHOTO_URL": "https://..." }
	var fieldsToUpdate map[string]string
	if err := json.NewDecoder(r.Body).Decode(&fieldsToUpdate); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// 3. Loop through the provided fields and upsert them in the database
	for fieldName, fieldValue := range fieldsToUpdate {
		err := h.Queries.UpsertStudentProfileField(ctx, db.UpsertStudentProfileFieldParams{
			StudentID:  studentID,
			FieldName:  db.StudentProfileInfoFieldName(fieldName), // Explicit type conversion here!
			FieldValue: fieldValue,
		})

		if err != nil {
			http.Error(w, "Failed to update profile field: "+fieldName, http.StatusBadRequest)
			return
		}
	}

	// 4. Return Success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Profile updated successfully!",
	})
}
