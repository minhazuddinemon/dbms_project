package rest

import (
	"encoding/json"
	"net/http"

	"dbms-project/internal/db"
	"dbms-project/internal/middleware"
)

type PaymentRequest struct {
	ApplicationID int32  `json:"application_id"`
	Amount        string `json:"amount"`         // e.g. "500.00"
	PaymentMethod string `json:"payment_method"` // e.g. "bKash", "Nagad", "Card"
	TransactionID string `json:"transaction_id"` // e.g. "TRX12345678"
}

func (h *Handler) HandleProcessPayment(w http.ResponseWriter, r *http.Request) {
	// 1. Authenticate Student
	studentID, ok := r.Context().Value(middleware.StudentIDKey).(int32)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// 2. Verify Application exists and belongs to this student
	app, err := h.Queries.GetApplicationByID(ctx, db.GetApplicationByIDParams{
		AppID:     req.ApplicationID, // Fixed: AppID instead of ApplicationID
		StudentID: studentID,
	})
	if err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	// Prevent double payment
	if app.Status == "PAID" || app.Status == "COMPLETED" {
		http.Error(w, "Application is already paid for", http.StatusBadRequest)
		return
	}

	// 3. Record Payment
	_, err = h.Queries.CreatePayment(ctx, db.CreatePaymentParams{
		AppID:  req.ApplicationID, // Fixed: AppID
		Amount: req.Amount,
		Method: req.PaymentMethod, // Fixed: Method
		TxID:   req.TransactionID, // Fixed: TxID
	})
	if err != nil {
		http.Error(w, "Failed to record payment transaction", http.StatusInternalServerError)
		return
	}

	// 4. Update Application Status to PAID
	err = h.Queries.UpdateApplicationStatus(ctx, db.UpdateApplicationStatusParams{
		Status:    "PAID",
		AppID:     req.ApplicationID, // Fixed: AppID
		StudentID: studentID,
	})
	if err != nil {
		http.Error(w, "Failed to update application status", http.StatusInternalServerError)
		return
	}

	// 5. Return Confirmation
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":         "success",
		"message":        "Payment received successfully! Application is now complete.",
		"application_id": req.ApplicationID,
		"transaction_id": req.TransactionID,
	})
}
