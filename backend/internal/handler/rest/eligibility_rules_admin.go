package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"dbms-project/internal/db"
)

type EligibilityRulePayload struct {
	ProgramID int32  `json:"program_id"`
	RuleType  string `json:"rule_type"`
	RuleValue string `json:"rule_value"`
}

func (h *Handler) HandleGetProgramEligibilityRules(w http.ResponseWriter, r *http.Request) {
	programIDStr := r.URL.Query().Get("program_id")
	programID, err := strconv.Atoi(programIDStr)
	if err != nil || programID <= 0 {
		http.Error(w, "Invalid or missing program_id", http.StatusBadRequest)
		return
	}

	rules, err := h.Queries.GetProgramEligibilityRules(r.Context(), int32(programID))
	if err != nil {
		http.Error(w, "Failed to retrieve eligibility rules: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rules == nil {
		rules = []db.ProgramEligibilityRule{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rules)
}

func (h *Handler) HandleUpsertProgramEligibilityRule(w http.ResponseWriter, r *http.Request) {
	var req EligibilityRulePayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	ruleType := strings.TrimSpace(req.RuleType)
	ruleValue := strings.TrimSpace(req.RuleValue)

	if req.ProgramID <= 0 || ruleType == "" || ruleValue == "" {
		http.Error(w, "program_id, rule_type, and rule_value are required", http.StatusBadRequest)
		return
	}

	err := h.Queries.UpsertProgramEligibilityRule(r.Context(), db.UpsertProgramEligibilityRuleParams{
		ProgramID: req.ProgramID,
		RuleType:  ruleType,
		RuleValue: sql.NullString{String: ruleValue, Valid: true},
	})
	if err != nil {
		http.Error(w, "Failed to save eligibility rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Eligibility rule saved successfully",
	})
}

func (h *Handler) HandleDeleteProgramEligibilityRule(w http.ResponseWriter, r *http.Request) {
	programIDStr := r.URL.Query().Get("program_id")
	programID, err := strconv.Atoi(programIDStr)
	if err != nil || programID <= 0 {
		http.Error(w, "Invalid or missing program_id", http.StatusBadRequest)
		return
	}

	ruleType := strings.TrimSpace(r.URL.Query().Get("rule_type"))
	if ruleType == "" {
		http.Error(w, "rule_type query parameter is required", http.StatusBadRequest)
		return
	}

	err = h.Queries.DeleteProgramEligibilityRule(r.Context(), db.DeleteProgramEligibilityRuleParams{
		ProgramID: int32(programID),
		RuleType:  ruleType,
	})
	if err != nil {
		http.Error(w, "Failed to delete eligibility rule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Eligibility rule deleted successfully",
	})
}
