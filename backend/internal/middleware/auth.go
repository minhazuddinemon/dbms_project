package middleware

import (
	"context"
	"net/http"
	"strings"

	"dbms-project/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const StudentIDKey contextKey = "student_id"

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		claims, err := auth.ValidateJWT(parts[1])
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		studentID, ok := claims["student_id"].(float64)
		if !ok {
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), StudentIDKey, int32(studentID))
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Read Authorization Header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 2. Parse and Validate JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("YOUR_SECRET_KEY"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// 3. Extract Claims & Verify Role
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		role, _ := claims["role"].(string)
		if role != "ADMIN" {
			http.Error(w, "Access denied: Admin privileges required", http.StatusForbidden)
			return
		}

		// Verified! Pass to actual handler
		next.ServeHTTP(w, r)
	}
}
