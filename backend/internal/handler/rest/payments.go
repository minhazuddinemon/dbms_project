package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

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
		AppID:     req.ApplicationID,
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

	// 3. Validate transaction ID against .env value
	envTxID := os.Getenv("VALID_TX_ID")
	if envTxID != "" && req.TransactionID != envTxID {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	// 4. Verify payment amount matches/covers application fee
	paidAmount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		http.Error(w, "Invalid amount format in request payload", http.StatusBadRequest)
		return
	}

	requiredAmount, err := strconv.ParseFloat(app.ApplicationFee, 64)
	if err != nil {
		requiredAmount = 0.00
	}

	if paidAmount < requiredAmount {
		http.Error(w, "Insufficient payment amount. Required fee: "+app.ApplicationFee, http.StatusBadRequest)
		return
	}

	// 5. Start a database transaction to record payment and create notification atomically
	tx, err := h.DB.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, "Failed to start database transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	qtx := h.Queries.WithTx(tx)

	// Record Payment
	_, err = qtx.CreatePayment(ctx, db.CreatePaymentParams{
		AppID:  req.ApplicationID,
		Amount: req.Amount,
		Method: req.PaymentMethod,
		TxID:   req.TransactionID,
	})
	if err != nil {
		http.Error(w, "Failed to record payment transaction", http.StatusInternalServerError)
		return
	}

	// Update Application Status to PAID
	err = qtx.UpdateApplicationStatus(ctx, db.UpdateApplicationStatusParams{
		Status:    "PAID",
		AppID:     req.ApplicationID,
		StudentID: studentID,
	})
	if err != nil {
		http.Error(w, "Failed to update application status", http.StatusInternalServerError)
		return
	}

	// 6. Notify student of successful application & exam details
	msg := "Your application for program " + app.ProgramName + " is successful!"
	test, err := qtx.GetAdmissionTestByProgramID(ctx, sql.NullInt32{Int32: app.ProgramID, Valid: true})
	if err == nil {
		if test.ExamDate.Valid {
			msg += " The exam will be held on " + test.ExamDate.Time.Format("2006-01-02")
		}
		if test.ExamCenter.Valid {
			msg += " at " + test.ExamCenter.String
		}
	} else if err == sql.ErrNoRows {
		msg += " Exam details will be notified soon."
	}
	msg += "."

	_, err = qtx.CreateNotification(ctx, db.CreateNotificationParams{
		StudentID: studentID,
		Message:   msg,
	})
	if err != nil {
		http.Error(w, "Failed to generate notification: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit payment transaction", http.StatusInternalServerError)
		return
	}

	// 7. Return Confirmation
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":         "success",
		"message":        "Payment received successfully! Application is now complete.",
		"application_id": req.ApplicationID,
		"transaction_id": req.TransactionID,
	})
}
