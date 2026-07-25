package rest

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"dbms-project/internal/auth"
	"dbms-project/internal/db"
	"dbms-project/internal/middleware"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{Queries: queries}
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	DOB       string `json:"dob"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// HandleRegister registers a new student account.
// @Summary Register a new student
// @Description Register a new student account with first name, last name, email, password, and date of birth (YYYY-MM-DD).
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body rest.RegisterRequest true "Student Registration Information"
// @Success 201 {object} map[string]string "Registration successful"
// @Failure 400 {string} string "Invalid request payload or date format"
// @Failure 405 {string} string "Method not allowed"
// @Failure 409 {string} string "Failed to create account (email might already exist)"
// @Failure 500 {string} string "Internal server error"
// @Router /register [post]
func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password hashing failed: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		http.Error(w, "Invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	_, err = h.Queries.CreateStudent(r.Context(), db.CreateStudentParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Dob:       dob,
	})

	if err != nil {
		log.Printf("Failed to create student: %v", err)
		http.Error(w, "Failed to create account (email might already exist)", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Registration successful"}`))
}

// HandleLogin authenticates a student and returns a JWT token.
// @Summary Login a student
// @Description Authenticates student credentials (email and password) and returns a signed JWT authentication token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body rest.LoginRequest true "Student Login Credentials"
// @Success 200 {object} map[string]string "Login successful with token"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 401 {string} string "Invalid email or password"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal server error"
// @Router /login [post]
func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	student, err := h.Queries.GetStudentByEmail(r.Context(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			log.Printf("Database error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.GenerateJWT(student.StudentID, student.Email)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   tokenString,
	})
}

// HandleProfile retrieves protected profile access confirmation for authenticated user.
// @Summary Get authenticated profile status
// @Description Returns confirmation and student ID for the currently authenticated student.
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]any "Access granted to protected profile"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Could not retrieve user identity"
// @Router /profile [get]
func (h *Handler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Could not retrieve user identity", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":    "Access granted to protected profile!",
		"student_id": studentID,
	})
}
