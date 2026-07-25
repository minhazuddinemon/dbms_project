package models

// LoginRequest defines the JSON payload for user login
type LoginRequest struct {
	Email    string `json:"email" example:"student@cuet.ac.bd"`
	Password string `json:"password" example:"secret123"`
}

// UserResponse defines the user data returned upon successful auth
type UserResponse struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"Minhaz Uddin"`
	Email string `json:"email" example:"student@cuet.ac.bd"`
	Role  string `json:"role" example:"applicant"`
}

// StandardErrorResponse defines error structure
type StandardErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Invalid credentials"`
}
