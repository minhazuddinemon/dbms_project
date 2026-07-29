package rest

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"dbms-project/internal/db"

	"github.com/golang-jwt/jwt/v5"
)

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) HandleAdminLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	adminEmail := os.Getenv("ADMIN_EMAIL")

	adminPassword := os.Getenv("ADMIN_PASSWORD")

	// Verify Credentials directly
	if req.Email != adminEmail || req.Password != adminPassword {
		http.Error(w, "Invalid admin credentials", http.StatusUnauthorized)
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	// Generate JWT Token with ADMIN role
	claims := jwt.MapClaims{
		"user_id": 0, // System Admin
		"role":    "ADMIN",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
		"role":  "ADMIN",
	})
} // DTO for status update request
type UpdateStatusRequest struct {
	AppID  int32  `json:"app_id"`
	Status string `json:"status"` // e.g., "APPROVED", "REJECTED"
}

// GET /admin/applications?u_id=1
func (h *Handler) GetUniversityApplications(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.URL.Query().Get("u_id")
	uID, err := strconv.Atoi(uIDStr)
	if err != nil || uID <= 0 {
		http.Error(w, "Invalid or missing university ID (u_id)", http.StatusBadRequest)
		return
	}

	apps, err := h.Queries.GetUniversityApplications(r.Context(), int32(uID))
	if err != nil {
		http.Error(w, "Failed to fetch university applications", http.StatusInternalServerError)
		return
	}

	if apps == nil {
		apps = []db.GetUniversityApplicationsRow{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apps)
}

// PUT /admin/applications/status
func (h *Handler) HandleUpdateApplicationStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.AppID <= 0 || req.Status == "" {
		http.Error(w, "app_id and status are required", http.StatusBadRequest)
		return
	}

	err := h.Queries.UpdateApplicationStatus(r.Context(), db.UpdateApplicationStatusParams{
		Status: req.Status,
		AppID:  req.AppID,
	})

	if err != nil {
		http.Error(w, "Failed to update application status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Application status updated to " + req.Status,
	})
}
